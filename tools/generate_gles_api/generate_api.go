// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package main

import (
	"bytes"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

const indent = "  "

type CommandSlice []*Command

func (cmds CommandSlice) Len() int { return len(cmds) }
func (cmds CommandSlice) Less(i, j int) bool {
	return cmds[i].Name() < cmds[j].Name()
}
func (cmds CommandSlice) Swap(i, j int) {
	cmd := cmds[i]
	cmds[i] = cmds[j]
	cmds[j] = cmd
}

// Generate the "gles.api" file from combined online and hand written data.
func GenerateApi(reg *Registry, api KhronosAPI) {
	files := map[string]*os.File{}
	printEnums(createFile(files, "glenum"), reg, api, false)
	printEnums(createFile(files, "glbitfield"), reg, api, true)
	commands := CommandSlice(reg.Command)
	sort.Sort(commands)
	for _, cmd := range commands {
		versions := reg.GetVersions(api, cmd.Name())
		extensions := reg.GetExtensions(api, cmd.Name())
		if versions != nil {
			printCommand(createFile(files, getCategory(cmd.Name())), reg, cmd, api)
		}
		if extensions != nil {
			if isInAndroidExtensionPack(extensions) {
				printCommand(createFile(files, "android_extension_pack"), reg, cmd, api)
			} else {
				printCommand(createFile(files, "extensions"), reg, cmd, api)
			}
		}
	}
	for _, file := range files {
		if err := file.Close(); err != nil {
			panic(err)
		}
	}
}

func createFile(files map[string]*os.File, basename string) *os.File {
	filename := *apiDir + string(os.PathSeparator) + basename + ".api"
	if file, ok := files[filename]; ok {
		return file
	}
	file, err := os.Create(filename)
	if err != nil {
		panic(err)
	}
	files[filename] = file
	fmt.Fprint(file, `// Copyright (C) 2015 The Android Open Source Project
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//      http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

`)
	return file
}

func getCategory(cmdName string) string {
	var table = []struct {
		category string
		cmdName  string // Regexp
	}{
		{"asynchronous_queries", ".*(Query|Queries).*"},
		{"synchronization", ".*Sync.*"},
		{"programs_and_shaders", ".*(Program|Shader|Compute|Uniform|MemoryBarrier|GetFragDataLocation|glValidateProgram|AttribLocation|GetActiveAttrib).*"},
		{"framebuffer", ".*(Framebuffer|Renderbuffer).*|glReadBuffer|glRead(|n)Pixels|gl(Color|Depth|Stencil)Mask|glClear.*|glInvalidate.*"},
		{"textures_and_samplers", ".*(Tex|Image|Sampler|PixelStorei|GenerateMipmap).*"},
		{"transform_feedback", "gl(.*TransformFeedback.*)"},
		{"vertex_arrays", ".*Vertex(Attrib|Buffer|Binding|Array).*"},
		{"draw_commands", "glDraw.*|glPatchParameteri|glPrimitiveBoundingBox"},
		{"rasterization", "gl(GetMultisamplefv|LineWidth|FrontFace|CullFace|PolygonOffset|DepthRangef|Viewport|MinSampleShading)"},
		{"fragment_operations", "glScissor|glSample.*|glStencil.*|glDepthFunc|glBlend.*"},
		{"state_queries", "gl(GetBoolean|GetInteger|GetFloat|IsEnabled|GetString|GetInternalformativ).*"},
		{"other", "glEnable|glDisable|glIsEnabled|glHint|glGetError|glFlush|glFinish|glGetGraphicsResetStatus"},
		{"debug", "(glDebug.*)|(glGetDebug.*)|gl(|Get)Object(|Ptr)Label|glGetPointerv|gl(Push|Pop)DebugGroup"},
		{"buffer_objects", ".*Buffer.*"},
	}
	for _, row := range table {
		if CompileRegexp("^(?:" + row.cmdName + ")").MatchString(cmdName) {
			return row.category
		}
	}
	panic(fmt.Errorf("Can not find category of %v", cmdName))
}

func isInAndroidExtensionPack(extensions []string) (result bool) {
	aep := map[string]bool{
		"GL_KHR_debug":                                true,
		"GL_KHR_texture_compression_astc_ldr":         true,
		"GL_KHR_blend_equation_advanced":              true,
		"GL_OES_sample_shading":                       true,
		"GL_OES_sample_variables":                     true,
		"GL_OES_shader_image_atomic":                  true,
		"GL_OES_shader_multisample_interpolation":     true,
		"GL_OES_texture_stencil8":                     true,
		"GL_OES_texture_storage_multisample_2d_array": true,
		"GL_EXT_copy_image":                           true,
		"GL_EXT_draw_buffers_indexed":                 true,
		"GL_EXT_geometry_shader":                      true,
		"GL_EXT_gpu_shader5":                          true,
		"GL_EXT_primitive_bounding_box":               true,
		"GL_EXT_shader_io_blocks":                     true,
		"GL_EXT_tessellation_shader":                  true,
		"GL_EXT_texture_border_clamp":                 true,
		"GL_EXT_texture_buffer":                       true,
		"GL_EXT_texture_cube_map_array":               true,
		"GL_EXT_texture_sRGB_decode":                  true,
	}
	for _, extension := range extensions {
		if aep[extension] {
			return true
		}
	}
	return false
}

func wrap(text string, width int, linePrefix string) string {
	if len(text) > width {
		if idx := strings.LastIndex(text[0:width], " "); idx != -1 {
			return text[0:idx] + "\n" + wrap(linePrefix+text[idx+1:], width, linePrefix)
		}
	}
	return text
}

func printEnums(out io.Writer, r *Registry, api KhronosAPI, bitfields bool) {
	if bitfields {
		fmt.Fprintf(out, "bitfield GLbitfield {\n")
	} else {
		fmt.Fprintf(out, "enum GLenum {\n")
	}
	for _, enums := range r.Enums {
		if (enums.Type == "bitmask") == bitfields && enums.Namespace == "GL" && len(enums.Enum) > 0 {
			for _, enum := range enums.Enum {
				// The following 64bit values are not proper GLenum values.
				if enum.Name == "GL_TIMEOUT_IGNORED" || enum.Name == "GL_TIMEOUT_IGNORED_APPLE" {
					continue
				}
				if enum.API == "" || enum.API == api {
					var value uint32
					if v, err := strconv.ParseUint(enum.Value, 0, 32); err == nil {
						value = uint32(v)
					} else if v, err := strconv.ParseInt(enum.Value, 0, 32); err == nil {
						value = uint32(v)
					} else {
						panic(fmt.Errorf("Failed to parse enum value %v", enum.Value))
					}
					fmt.Fprintf(out, indent+"%-64s = 0x%08X,\n", enum.Name, value)
				}
			}
		}
	}
	fmt.Fprintf(out, "}\n\n")
}

func renameType(cmdName string, paramIndex int, paramType, paramName string) string {
	if paramType == "const void *" {
		if oldType, ok := oldParamType(cmdName, paramIndex); ok {
			if strings.Contains(oldType, "Pointer") {
				return oldType
			}
		}
	}
	if paramType == "const GLchar *" {
		if oldType, ok := oldParamType(cmdName, paramIndex); ok {
			if oldType == "string" {
				return oldType
			}
		}
	}

	// Find the root GL type, ignoring any prefix or sufix
	loc := CompileRegexp(`\bGL\w+\b`).FindStringIndex(paramType)
	if loc == nil {
		return paramType
	}
	glType := paramType[loc[0]:loc[1]]

	// Type substitution table
	table := []struct {
		oldType   string
		cmdName   string // regexp
		paramName string // regexp
		newType   string
	}{
		{"GLint", ".*Uniform.*", "location|result", "UniformLocation"},
		{"GLuint", ".*", "program(|s)", "ProgramId"},
		{"GLuint", ".*", "texture(|s)", "TextureId"},
		{"GLuint", ".*", "uniformBlockIndex", "UniformBlockId"},
		{"GLuint", ".*(Query|Queries).*", "id(|s)", "QueryId"},
		{"GLuint", ".*Buffer.*", "buffer(|s)", "BufferId"},
		{"GLuint", ".*Framebuffer.*", "framebuffer(|s)", "FramebufferId"},
		{"GLuint", ".*Program.*", "pipeline(|s)", "PipelineId"},
		{"GLuint", ".*Renderbuffer.*", "renderbuffer(|s)", "RenderbufferId"},
		{"GLuint", ".*Sampler.*", "sampler(|s)", "SamplerId"},
		{"GLuint", ".*Shader.*", "shader(|s)", "ShaderId"},
		{"GLuint", ".*TransformFeedback.*", "id(|s)", "TransformFeedbackId"},
		{"GLuint", ".*VertexArray.*", "array(|s)", "VertexArrayId"},
		{"GLuint", "glCreateProgram", "result", "ProgramId"},
		{"GLuint", "glCreateShader", "result", "ShaderId"},
		{"GLuint", "glCreateShaderProgramv", "result", "ProgramId"},
		{"GLuint", "glGetUniformBlockIndex", "result", "UniformBlockId"},
		{"GLuint", ".*Attrib.*", "result|index", "AttributeLocation"},
	}
	for _, row := range table {
		if row.oldType == glType &&
			CompileRegexp("^(?:"+row.paramName+")$").MatchString(paramName) &&
			CompileRegexp("^(?:"+row.cmdName+")$").MatchString(cmdName) {
			// Replace the GL type, but keep the original prefix and suffix
			return paramType[:loc[0]] + row.newType + paramType[loc[1]:]
		}
	}
	return paramType
}

func printVersionChecks(out io.Writer, cmd *Command, versions []Version, extensions []string) {
	for _, extension := range extensions {
		fmt.Fprintf(out, indent+"requiresExtension(%s)\n", extension)
	}
	if len(extensions) > 1 {
		fmt.Fprintf(out, indent+"// TODO: Multiple extensions\n")
	}
	if versions != nil {
		fmt.Fprintf(out, indent+"minRequiredVersion(%v)\n", strings.Replace(string(versions[0]), ".", ", ", -1))
	}
}

func printChecks(out io.Writer, cmd *Command, versions []Version) {
	if versions == nil {
		return
	}
	for paramIndex, param := range cmd.Param {
		paramName := oldParamName(cmd.Name(), paramIndex, param.Name)
		if param.Type() == "GLenum" {
			fmt.Fprintf(out, indent+"switch (%s) {\n", paramName)
			seen := map[string]Version{}
			for _, version := range versions {
				doc := DownloadDoc(version, cmd.Name())
				var newCases []string
				accepts := doc.Params[paramIndex].Accepts
				if accepts == nil {
					fmt.Printf("%s(%s): Missing accepted GLenum values for %s\n", cmd.Name(), version, paramName)
					continue
				}
				for _, value := range accepts {
					if _, ok := seen[value]; !ok {
						newCases = append(newCases, value)
					}
					seen[value] = version
				}
				// Check that the set of accepted values is increasing from version to version.
				for k, v := range seen {
					if v != version {
						fmt.Printf("%s(%s): %s seems to be removed\n", cmd.Name(), version, k)
					}
				}
				if newCases != nil {
					line := fmt.Sprintf(indent+indent+"case %s: {\n", strings.Join(newCases, ", "))
					fmt.Fprint(out, wrap(line, 100, indent+indent+indent+indent))
					if version != versions[0] {
						fmt.Fprintf(out, indent+indent+indent+"minRequiredVersion(%v)\n", strings.Replace(string(version), ".", ", ", -1))
					} else {
						fmt.Fprintf(out, indent+indent+indent+"// version %v\n", version)
					}
					fmt.Fprint(out, indent+indent+"}\n")
				}
			}
			fmt.Fprint(out, indent+indent+"default: {\n")
			fmt.Fprintf(out, indent+indent+indent+"glErrorInvalidEnum(%s)\n", paramName)
			fmt.Fprint(out, indent+indent+"}\n")
			fmt.Fprintf(out, indent+"}\n")
		}
		if param.Type() == "GLbitfield" {
			bits := []string{}
			seen := map[string]Version{} // name -> min_version
			for _, version := range versions {
				doc := DownloadDoc(version, cmd.Name())
				for _, value := range doc.Params[paramIndex].Accepts {
					if _, ok := seen[value]; !ok {
						bits = append(bits, value)
						seen[value] = version
					}
				}
			}
			if len(bits) > 0 {
				fmt.Fprintf(out, indent+"supportsBits(%s, %s)\n", paramName, strings.Join(bits, "|"))
				for _, bit := range bits {
					version := seen[bit]
					fmt.Fprintf(out, indent+"if (%s in %s) {\n", bit, paramName)
					if version != versions[0] {
						fmt.Fprintf(out, indent+indent+"minRequiredVersion(%v)\n", strings.Replace(string(version), ".", ", ", -1))
					}
					fmt.Fprintf(out, indent+"}\n")
				}
			}
		}
	}
}

func printCommand(out io.Writer, reg *Registry, cmd *Command, api KhronosAPI) {
	versions := reg.GetVersions(api, cmd.Name())
	extensions := reg.GetExtensions(api, cmd.Name())
	if versions != nil && extensions != nil {
		panic(fmt.Errorf("%s is both core feature and extension", cmd.Name()))
	}

	// Function signature
	params := &bytes.Buffer{}
	leadWidth := params.Len()
	maxTypeWidth := 0
	paramTypes := []string{}
	for paramIndex, param := range cmd.Param {
		paramType := renameType(cmd.Name(), paramIndex, param.Type(), param.Name)
		paramTypes = append(paramTypes, paramType)
		if typeWidth := len(paramType); typeWidth > maxTypeWidth {
			maxTypeWidth = typeWidth
		}
	}
	for i, param := range cmd.Param {
		if i > 0 {
			fmt.Fprintf(params, ",\n%s", strings.Repeat(" ", leadWidth))
		}
		format := fmt.Sprintf("%%-%vs %%s", maxTypeWidth)
		fmt.Fprintf(params, format, paramTypes[i], oldParamName(cmd.Name(), i, param.Name))
		if oldType, ok := oldParamType(cmd.Name(), i); ok {
			newType := strings.Replace(paramTypes[i], " *", "*", -1)
			if newType != oldType {
				// fmt.Printf("Changed type of %s:%s from %s to %s\n", cmd.Name(), param.Name, oldType, newType)
			}
		}
	}
	returnType := renameType(cmd.Name(), -1, cmd.Proto.Type(), "result")

	// Try to find the old code
	oldCode := ""
	if oldApi != nil {
		for _, apiFunction := range oldApi.Functions {
			apiCmd := apiFunction.AST
			if apiCmd.Name.Value == cmd.Name() {
				token := apiCmd.Block.CST.Token()
				oldCode = string(token.Source.Runes[token.Start+1 : token.End-1])
			}
		}
	}

	// See if we should generate macro
	macroName := ""
	if cmd.Alias.Name != "" {
		if reg.GetVersions(GLES2API, cmd.Alias.Name) != nil {
			macroName = cmd.Alias.Name // this cmd is extension
		}
	}
	if macroName == "" {
		for _, c := range reg.Command {
			if c.Alias.Name == cmd.Name() {
				if reg.GetExtensions(GLES2API, c.Name()) != nil {
					macroName = cmd.Name() // this cmd is core version of extension
				}
			}
		}
	}
	if strings.HasPrefix(macroName, "gl") {
		macroName = macroName[2:]
	}
	if macroName != "" && oldCode != "" {
		if strings.HasPrefix(oldCode, "\n  minRequiredVersion(") {
			oldCode = oldCode[strings.IndexRune(oldCode, ')')+1:]
		}
	}

	// Print comments
	for _, version := range versions {
		url, _, _, _ := GetDocUrl(version, cmd.Name())
		fmt.Fprintf(out, `@Doc("%s","OpenGL ES %v")`+"\n", url, version)
	}
	for _, extension := range extensions {
		url := GetExtensionUrl(extension)
		fmt.Fprintf(out, `@Doc("%s","%v")`+"\n", url, extension)
	}

	longSig := fmt.Sprintf("cmd %s %s(%s)", returnType, cmd.Name(), string(params.Bytes()))
	shortSig := CompileRegexp(`\s+`).ReplaceAllString(longSig, " ")
	if len(shortSig)+2 <= 100 {
		fmt.Fprintf(out, "%s ", shortSig)
	} else {
		fmt.Fprintf(out, "%s ", longSig)
	}

	// Generate checks
	checks := ""
	{
		out := &bytes.Buffer{}
		printVersionChecks(out, cmd, versions, extensions)
		printChecks(out, cmd, versions)
		checks = string(out.Bytes())
	}

	// Print the method body
	fmt.Fprintf(out, "{")
	if macroName != "" {
		fmt.Fprint(out, "\n"+indent)
		printVersionChecks(out, cmd, versions, extensions)
		if cmd.Proto.Type() != "void" {
			fmt.Fprint(out, "return ")
		}
		fmt.Fprintf(out, "%s(", macroName)
		for i, param := range cmd.Param {
			if i > 0 {
				fmt.Fprintf(out, ",")
			}
			fmt.Fprintf(out, "%s", oldParamName(cmd.Name(), i, param.Name))
		}
		fmt.Fprint(out, ")\n")
	} else if true || strings.HasPrefix(oldCode, "\n"+checks) {
		// The checks in the existing code match what we expect so keep the old code.
		fmt.Fprintf(out, "%s", oldCode)
	} else {
		// Something is different.  Print both and sort it out manually.
		fmt.Fprintf(out, "\n%s", checks)
		fmt.Fprintf(out, indent+"// TODO\n")
		if cmd.Proto.Type() != "void" {
			fmt.Fprint(out, indent+"return ?\n")
		}
		fmt.Fprintf(out, "%s", oldCode)
	}
	fmt.Fprintf(out, "}")
	fmt.Fprintf(out, "\n\n")

	// Print macro
	if macroName != "" && versions != nil {
		shortSig := CompileRegexp(`\s+`).ReplaceAllString(string(params.Bytes()), " ")
		fmt.Fprintf(out, "macro %s %s(%s) {%s}\n\n", returnType, macroName, shortSig, oldCode)
	}
}

// Rename parameters so that they match our old api file
func oldParamName(cmdName string, paramIndex int, paramName string) string {
	if oldApi != nil {
		for _, apiFunction := range oldApi.Functions {
			cmd := apiFunction.AST
			if cmd.Name.Value == cmdName {
				return cmd.Parameters[paramIndex].Name.Value
			}
		}
	}
	return paramName
}

func oldParamType(cmdName string, paramIndex int) (string, bool) {
	if oldApi != nil {
		for _, apiFunction := range oldApi.Functions {
			cmd := apiFunction.AST
			if cmd.Name.Value == cmdName {
				token := cmd.Parameters[paramIndex].Type.Node().Token()
				return string(token.Source.Runes[token.Start:token.End]), true
			}
		}
	}
	return "", false
}
