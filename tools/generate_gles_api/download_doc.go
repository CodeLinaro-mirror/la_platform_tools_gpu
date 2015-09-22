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
	"encoding/xml"
	"fmt"
	"regexp"
	"sort"
	"strings"
)

// Download documentation from the khronos online man pages
func DownloadDoc(version Version, cmdName string) *CommandDoc {
	url, manPage, paramSectionIndex, paramSectionCount := GetDocUrl(version, cmdName)
	page := &html{}
	if err := xml.Unmarshal(Download(url), page); err != nil {
		panic(err)
	}
	summary := page.FindClass("refnamediv").FindTag("p").Text()
	paramNameTags := page.FindOne(func(h *html) bool {
		return h.Class == "funcprototype-table" && h.FindClass("fsfunc").Text() == cmdName
	}).FindClasses("pdparam")
	paramSections := page.FindAll(func(h *html) bool {
		return h.Class == "refsect1" && strings.HasPrefix(h.FindTag("h2").Text(), "Parameters")
	})
	paramDocs := make(map[string]string) // name -> doc
	if len(paramSections) > 0 {
		if len(paramSections) != paramSectionCount {
			panic(fmt.Errorf("Expected %v parameter sections, but found %v", paramSectionCount, len(paramSections)))
		}
		dlTag := paramSections[paramSectionIndex].FindTag("dl").Children
		for i := 0; i < len(dlTag); i += 2 {
			for _, name := range dlTag[i].FindTags("code") {
				paramDocs[name.Text()] = dlTag[i+1].XML
			}
		}
	}
	accpets, _ := getAcceptedValues(cmdName, "result", version)
	params := map[int]ParamDoc{-1: {"result", summary, accpets}}
	for paramIndex, paramNameTag := range paramNameTags {
		if name := paramNameTag.Text(); name != "void" {
			doc := parseParamDoc(name, paramDocs[name])
			doc.Name = renameParam(cmdName, doc.Name)
			if accpets, ok := getAcceptedValues(cmdName, doc.Name, version); ok {
				doc.Accepts = accpets
			}
			sort.Sort(sort.StringSlice(doc.Accepts))
			params[paramIndex] = doc
		}
	}
	return &CommandDoc{
		ManPage: manPage,
		Url:     url,
		Summary: summary,
		Params:  params,
	}
}

type CommandDoc struct {
	ManPage string
	Url     string
	Summary string
	Params  map[int]ParamDoc
}

type ParamDoc struct {
	Name    string
	Summary string
	Accepts []string
}

type html struct {
	XMLName  xml.Name
	XML      string  `xml:",innerxml"`
	Class    string  `xml:"class,attr"`
	Children []*html `xml:",any"`
}

func (h *html) FindAll(pred func(*html) bool) (res []*html) {
	if pred(h) {
		res = append(res, h)
	}
	for _, child := range h.Children {
		res = append(res, child.FindAll(pred)...)
	}
	return
}

func (h *html) FindOne(pred func(*html) bool) *html {
	res := h.FindAll(pred)
	if len(res) == 0 {
		panic(fmt.Errorf("Could not find matching html element"))
	} else if len(res) > 1 {
		panic(fmt.Errorf("Found multiple matching html elements"))
	} else {
		return res[0]
	}
}

func (h *html) FindTags(tag string) []*html {
	return h.FindAll(func(h *html) bool { return h.XMLName.Local == tag })
}

func (h *html) FindTag(tag string) *html {
	return h.FindOne(func(h *html) bool { return h.XMLName.Local == tag })
}

func (h *html) FindClasses(class string) []*html {
	return h.FindAll(func(h *html) bool { return h.Class == class })
}

func (h *html) FindClass(class string) *html {
	return h.FindOne(func(h *html) bool { return h.Class == class })
}

// Remove html tags to obtain plain text
func htmlToText(page string) string {
	page = strings.Replace(page, "<mml:msup><mml:mn>2</mml:mn>", "2^", -1)
	page = CompileRegexp("<.*?>").ReplaceAllString(page, "")
	page = strings.Replace(page, "&lt;", "<", -1)
	page = strings.Replace(page, "&gt;", ">", -1)
	page = CompileRegexp(`\s+`).ReplaceAllString(page, " ")
	return strings.TrimSpace(page)
}

func (h *html) Text() string {
	return htmlToText(h.XML)
}

// Get the full url for the given command, the page name (eg. "glUniform"), and the relevat section
func GetDocUrl(version Version, cmdName string) (url, manPage string, paramSectionIndex int, paramSectionCount int) {

	// Sometimes the man page is exactly the command name,
	// however sometimes one page documents several command.
	// (overload, enable/disable, begin/end)
	// We use the following tables to find the relevat man page.
	// We also sometimes have to specify which section of the page is relevat.

	type redirect struct {
		oldName           string // regexp
		newName           string // man page
		paramSectionIndex int
		paramSectionCount int // for verification
	}

	var gles11redirects = []*redirect{
		{"glAlphaFunc(x)", "glAlphaFunc", 0, 1},
		{"glClearColor(x)", "glClearColor", 0, 1},
		{"glClearDepth(f|x)", "glClearDepth", 0, 1},
		{"glClipPlane(f|x)", "glClipPlane", 0, 1},
		{"glColor4(f|x|ub)", "glColor", 0, 1},
		{"glDepthRange(f|x)", "glDepthRange", 0, 1},
		{"glDisable", "glEnable", 0, 1},
		{"glDisableClientState", "glEnableClientState", 0, 1},
		{"glFog(f|x)", "glFog", 0, 2},
		{"glFog(f|x)v", "glFog", 1, 2},
		{"glFrustum(f|x)", "glFrustum", 0, 1},
		{"glGetBoolean(v)", "glGet", 0, 1},
		{"glGetClipPlane(f|x)", "glGetClipPlane", 0, 1},
		{"glGetFixed(v)", "glGet", 0, 1},
		{"glGetFloat(v)", "glGet", 0, 1},
		{"glGetInteger(v)", "glGet", 0, 1},
		{"glGetLight(f|x)v", "glGetLight", 0, 1},
		{"glGetMaterial(f|x)v", "glGetMaterial", 0, 1},
		{"glGetTexEnv(f|i|x)v", "glGetTexEnv", 0, 1},
		{"glGetTexParameter(f|i|x)v", "glGetTexParameter", 0, 1},
		{"glLight(f|x)", "glLight", 0, 2},
		{"glLight(f|x)v", "glLight", 1, 2},
		{"glLightModel(f|x)", "glLightModel", 0, 2},
		{"glLightModel(f|x)v", "glLightModel", 1, 2},
		{"glLineWidth(x)", "glLineWidth", 0, 1},
		{"glLoadMatrix(f|x)", "glLoadMatrix", 0, 1},
		{"glMaterial(f|x)", "glMaterial", 0, 2},
		{"glMaterial(f|x)v", "glMaterial", 1, 2},
		{"glMultMatrix(f|x)", "glMultMatrix", 0, 1},
		{"glMultiTexCoord4(f|x)", "glMultiTexCoord", 0, 1},
		{"glNormal3(f|x)", "glNormal", 0, 1},
		{"glOrtho(f|x)", "glOrtho", 0, 1},
		{"glPointParameter(f|x)", "glPointParameter", 0, 2},
		{"glPointParameter(f|x)v", "glPointParameter", 1, 2},
		{"glPointSize(x)", "glPointSize", 0, 1},
		{"glPolygonOffset(x)", "glPolygonOffset", 0, 1},
		{"glPopMatrix", "glPushMatrix", 0, 1},
		{"glRotate(f|x)", "glRotate", 0, 1},
		{"glSampleCoverage(x)", "glSampleCoverage", 0, 1},
		{"glScale(f|x)", "glScale", 0, 1},
		{"glTexEnv(f|i|x)", "glTexEnv", 0, 2},
		{"glTexEnv(f|i|x)v", "glTexEnv", 1, 2},
		{"glTexParameter(f|i|x)", "glTexParameter", 0, 2},
		{"glTexParameter(f|i|x)v", "glTexParameter", 1, 2},
		{"glTranslate(f|x)", "glTranslate", 0, 1},
	}

	var gles20redirects = []*redirect{
		{"glDisable", "glEnable", 1, 2},
		{"glDisableVertexAttribArray", "glEnableVertexAttribArray", 0, 1},
		{"glEnable", "glEnable", 0, 2},
		{"glGetBooleanv", "glGet", 0, 1},
		{"glGetFloatv", "glGet", 0, 1},
		{"glGetIntegerv", "glGet", 0, 1},
		{"glGetTexParameter(f|i)v", "glGetTexParameter", 0, 1},
		{"glGetUniform(f|i)v", "glGetUniform", 0, 1},
		{"glGetVertexAttrib(f|i)v", "glGetVertexAttrib", 0, 1},
		{"glTexParameter(f|i)", "glTexParameter", 0, 2},
		{"glTexParameter(f|i)v", "glTexParameter", 1, 2},
		{"glUniform(1|2|3|4)(f|i|ui)", "glUniform", 0, 3},
		{"glUniform(1|2|3|4)(f|i|ui)v", "glUniform", 1, 3},
		{"glUniformMatrix(2|3|4|2x3|3x2|2x4|4x2|3x4|4x3)fv", "glUniform", 2, 3},
		{"glVertexAttrib(1|2|3|4)f", "glVertexAttrib", 0, 2},
		{"glVertexAttrib(1|2|3|4)fv", "glVertexAttrib", 1, 2},
	}

	var gles30redirects = append([]*redirect{
		{"glBeginQuery", "glBeginQuery", 0, 2},
		{"glClearBuffer(fi|fv|iv|uiv)", "glClearBuffer", 0, 1},
		{"glDisable", "glEnable", 0, 1},
		{"glEnable", "glEnable", 0, 1},
		{"glEndQuery", "glBeginQuery", 1, 2},
		{"glEndTransformFeedback", "glBeginTransformFeedback", 0, 1},
		{"glGetBufferParameter(i)(|64)v", "glGetBufferParameter", 0, 1},
		{"glGetInteger(|64)(|i_)v", "glGet", 0, 1},
		{"glGetSamplerParameter(f|i)v", "glGetSamplerParameter", 0, 1},
		{"glGetString(|i)", "glGetString", 0, 1},
		{"glGetUniform(f|i|ui)v", "glGetUniform", 0, 1},
		{"glGetVertexAttrib(f|Ii|Iui|i)v", "glGetVertexAttrib", 0, 1},
		{"glMapBufferRange", "glMapBufferRange", 0, 2},
		{"glSamplerParameter(f|i)(|v)", "glSamplerParameter", 0, 1},
		{"glTexParameter(f|i)(|v)", "glTexParameter", 0, 1},
		{"glUniform(1|2|3|4)(f|i|ui)(|v)", "glUniform", 0, 1},
		{"glUniformMatrix(2|3|4|2x3|3x2|2x4|4x2|3x4|4x3)fv", "glUniform", 0, 1},
		{"glUnmapBuffer", "glMapBufferRange", 1, 2},
		{"glVertexAttrib(1f|2f|3f|4f|I4i|I4ui)(|v)", "glVertexAttrib", 0, 1},
		{"glVertexAttrib(|I)Pointer", "glVertexAttribPointer", 0, 1},
	}, gles20redirects...)

	var gles31redirects = append([]*redirect{
		{"glCreateShaderProgram(v)", "glCreateShaderProgram", 0, 1},
		{"glGetBoolean(|i_)v", "glGet", 0, 1},
		{"glGetProgramInterface(iv)", "glGetProgramInterface", 0, 1},
		{"glGetProgramPipeline(iv)", "glGetProgramPipeline", 0, 1},
		{"glGetProgramResource(iv)", "glGetProgramResource", 0, 1},
		{"glGetTexLevelParameter(f|i)v", "glGetTexLevelParameter", 0, 1},
		{"glMemoryBarrier(|ByRegion)", "glMemoryBarrier", 0, 1},
		{"glProgramUniform(1|2|3|4)(f|i|ui)(|v)", "glProgramUniform", 0, 1},
		{"glProgramUniformMatrix(2|3|4|2x3|3x2|2x4|4x2|3x4|4x3)fv", "glProgramUniform", 0, 1},
		{"glVertexAttrib(|I)Format", "glVertexAttribFormat", 0, 1},
	}, gles30redirects...)

	var gles32redirects = append([]*redirect{
		{"glBlendEquationi", "glBlendEquation", 0, 1},
		{"glBlendEquationSeparatei", "glBlendEquationSeparate", 0, 1},
		{"glBlendFuncSeparatei", "glBlendFuncSeparate", 0, 1},
		{"glBlendFunci", "glBlendFunc", 0, 1},
		{"glEnable(|i)|glDisable(|i)", "glEnable", 0, 1},
		{"glGetSamplerParameter(f|i|Ii|Iui)v", "glGetSamplerParameter", 0, 1},
		{"glSamplerParameter(f|i|Ii|Iui)(|v)", "glSamplerParameter", 0, 1},
		{"glGetTexParameter(f|i|Ii|Iui)v", "glGetTexParameter", 0, 1},
		{"glTexParameter(f|i|Ii|Iui)v", "glTexParameter", 0, 1},
		{"glIsEnabledi", "glIsEnabled", 0, 1},
		{"glReadnPixels", "glReadPixels", 0, 1},
		{"glColorMaski", "glColorMask", 0, 1},
		{"glGet(|n)Uniform(f|i|ui)v", "glGetUniform", 0, 1},
	}, gles31redirects...)

	var redirects []*redirect
	var urlFormat string
	switch version {
	case "1.1":
		redirects, urlFormat = gles11redirects, "https://www.khronos.org/opengles/sdk/1.1/docs/man/%s.xml"
	case "2.0":
		redirects, urlFormat = gles20redirects, "https://www.khronos.org/opengles/sdk/docs/man/xhtml/%s.xml"
	case "3.0":
		redirects, urlFormat = gles30redirects, "https://www.khronos.org/opengles/sdk/docs/man3/html/%s.xhtml"
	case "3.1":
		redirects, urlFormat = gles31redirects, "https://www.khronos.org/opengles/sdk/docs/man31/html/%s.xhtml"
	case "3.2":
		redirects, urlFormat = gles32redirects, "https://www.khronos.org/opengles/sdk/docs/man32/html/%s.xhtml"
	default:
		panic(fmt.Errorf("Uknown api version: %v", version))
	}

	for _, redir := range redirects {
		if CompileRegexp("^(?:" + redir.oldName + ")$").MatchString(cmdName) {
			return fmt.Sprintf(urlFormat, redir.newName), redir.newName, redir.paramSectionIndex, redir.paramSectionCount
		}
	}
	return fmt.Sprintf(urlFormat, cmdName), cmdName, 0, 1
}

func GetExtensionUrl(extension string) string {
	match := CompileRegexp(`^GL_(([A-Z]+)_\w+)$`).FindStringSubmatch(extension)
	vendor := match[2]
	page := match[1]
	// most of the time the page name includes the vendor prefix, but sometimes it does not
	for _, rename := range []struct{ oldPage, newPage string }{
		{"AMD_performance_monitor", "performance_monitor"},
		{"APPLE_rgb_422", "rgb_422"},
		{"EXT_blend_minmax", "blend_minmax"},
		{"EXT_draw_instanced", "draw_instanced"},
		{"EXT_multi_draw_arrays", "multi_draw_arrays"},
		{"EXT_post_depth_coverage", "post_depth_coverage"},
		{"EXT_raster_multisample", "raster_multisample"},
		{"EXT_shader_integer_mix", "shader_integer_mix"},
		{"EXT_texture_compression_dxt1", "texture_compression_dxt1"},
		{"EXT_texture_compression_s3tc", "texture_compression_s3tc"},
		{"EXT_texture_filter_anisotropic", "texture_filter_anisotropic"},
		{"EXT_texture_filter_minmax", "texture_filter_minmax"},
		{"EXT_texture_lod_bias", "texture_lod_bias"},
		{"EXT_texture_sRGB_decode", "texture_sRGB_decode"},
		{"FJ_shader_binary_GCCSO", "shader_binary_GCCSO"},
		{"INTEL_framebuffer_CMAA", "framebuffer_CMAA"},
		{"INTEL_performance_query", "performance_query"},
		{"KHR_blend_equation_advanced", "blend_equation_advanced"},
		{"KHR_context_flush_control", "context_flush_control"},
		{"KHR_debug", "debug"},
		{"KHR_robust_buffer_access_behavior", "robust_buffer_access_behavior"},
		{"KHR_robustness", "robustness"},
		{"KHR_texture_compression_astc_hdr", "texture_compression_astc_hdr"},
		{"NV_bindless_texture", "bindless_texture"},
		{"NV_blend_equation_advanced", "blend_equation_advanced"},
		{"NV_conditional_render", "conditional_render"},
		{"NV_conservative_raster", "conservative_raster"},
		{"NV_coverage_sample", "EGL_NV_coverage_sample"},
		{"NV_depth_nonlinear", "EGL_NV_depth_nonlinear"},
		{"NV_draw_texture", "draw_texture"},
		{"NV_EGL_NV_coverage_sample", "EGL_NV_coverage_sample"},
		{"NV_fence", "fence"},
		{"NV_fill_rectangle", "fill_rectangle"},
		{"NV_fragment_coverage_to_color", "fragment_coverage_to_color"},
		{"NV_fragment_shader_interlock", "fragment_shader_interlock"},
		{"NV_framebuffer_mixed_samples", "framebuffer_mixed_samples"},
		{"NV_geometry_shader_passthrough", "geometry_shader_passthrough"},
		{"NV_internalformat_sample_query", "internalformat_sample_query"},
		{"NV_path_rendering", "path_rendering"},
		{"NV_path_rendering_shared_edge", "path_rendering_shared_edge"},
		{"NV_sample_locations", "sample_locations"},
		{"NV_sample_mask_override_coverage", "sample_mask_override_coverage"},
		{"NV_viewport_array2", "viewport_array2"},
		{"OVR_multiview", "multiview"},
		{"OVR_multiview2", "multiview2"},
	} {
		if page == rename.oldPage {
			page = rename.newPage
		}
	}
	url := fmt.Sprintf("https://www.khronos.org/registry/gles/extensions/%s/%s.txt", vendor, page)
	Download(url)
	return url
}

// More often than not, the accpted GLenum values can be extracted using regexp.
func parseParamDoc(paramName, doc string) ParamDoc {
	// Remove noise
	for _, tag := range []string{"<span>", "</span>", "<p>", "</p>"} {
		doc = strings.Replace(doc, tag, "", -1)
	}

	// Regexp for enum value.  Represented in phrases as $
	exprConst := `(?:<code class="constant">(\w+)</code>(?:,|\.|\s|and|or)*)`

	// Remove the default value doc
	phrases := []string{
		"The initial value is $.",
		"The intial value is $.", // typo in docs
	}
	for _, phrase := range phrases {
		// Must follow sentence; relaxed whitespace; must be end of doc
		expr := `\.\s*` + strings.Replace(regexp.QuoteMeta(phrase), ` `, `\s+`, -1) + `\s*$`
		expr = strings.Replace(expr, `\$`, exprConst, -1)
		if loc := CompileRegexp(expr).FindStringIndex(doc); loc != nil {
			doc = doc[:loc[0]+1]
		}
	}

	// List of the common phrases which the man pages use
	paramNameTag := `<em class="parameter"><code>` + paramName + `</code></em>`
	phrases = []string{
		"$ are accepted",
		"$ is accepted",
		"Accepted symbolic names are $",
		"Accepted values are $",
		"Can be either $",
		"Eight symbolic constants are accepted: $",
		"Eight symbolic constants are valid: $",
		"Eight tokens are valid: $",
		"It must be $",
		"May be $",
		"May be either $",
		"Must be $",
		"Must be either $",
		"Must be one of $",
		"Must be one of the following symbolic constants: $",
		"Six symbolic constants are accepted: $",
		"Symbolic constants $ are accepted",
		"The allowed flags are $",
		"The following symbolic constants are accepted: $",
		"The following symbolic values are accepted: $",
		"The following symbols are accepted: $",
		"The symbolic constant must be $",
		"The symbolic constant must be one of $",
		"The symbolc constants $ is accepted",
		"The three masks are $",
		"These values are accepted: $",
		"These values are accepted: $",
		"Three symbolic constants are valid: $",
		"Valid masks are $",
		"Which can be one of the following: $",
		paramNameTag + " can be one of the following: $",
		paramNameTag + " may be $",
		paramNameTag + " must be $",
		paramNameTag + " must be either $",
		paramNameTag + " must be one of $",
	}
	var enumValues []string
	for _, phrase := range phrases {
		// Must follow sentence (dot); relaxed whitespace; must be end of doc
		expr := `\.\s*` + strings.Replace(regexp.QuoteMeta(phrase), ` `, `\s+`, -1) + `(?:\.|,|\s)*$`
		expr = strings.Replace(expr, `\$`, exprConst+`*`, -1)
		if loc := CompileRegexp(expr).FindStringIndex(doc); loc != nil {
			for _, match := range CompileRegexp(exprConst).FindAllStringSubmatch(doc[loc[0]:loc[1]], -1) {
				enumValues = append(enumValues, match[1])
			}
			// remove the enum list from the doc
			doc = doc[:loc[0]+1]
			break
		}
	}
	return ParamDoc{
		Name:    paramName,
		Summary: htmlToText(doc),
		Accepts: enumValues,
	}
}

// rename parameters so that they match names in the khronos registry
// this used to be important, but it is used just for verification now
func renameParam(cmdName, paramName string) string {
	var renames = []struct {
		cmdName string // regexp
		newName string
		oldName string
	}{
		{"glBlendFuncSeparate", "dfactorAlpha", "dstAlpha"},
		{"glBlendFuncSeparate", "dfactorRGB", "dstRGB"},
		{"glBlendFuncSeparate", "sfactorAlpha", "srcAlpha"},
		{"glBlendFuncSeparate", "sfactorRGB", "srcRGB"},
		{"glBlendFunc(|i)", "sfactor", "src"},
		{"glBlendFunc(|i)", "dfactor", "dst"},
		{"glClearBuffer(fi|fv|iv|uiv)", "drawbuffer", "drawBuffer"},
		{"glClearDepthf", "d", "depth"},
		{"glClipPlanef", "eqn", "equation"},
		{"glClipPlanef", "p", "plane"},
		{"glCopyBufferSubData", "readOffset", "readoffset"},
		{"glCopyBufferSubData", "readTarget", "readtarget"},
		{"glCopyBufferSubData", "writeOffset", "writeoffset"},
		{"glCopyBufferSubData", "writeTarget", "writetarget"},
		{"glCreateShader", "type", "shaderType"},
		{"glDeleteSamplers", "count", "n"},
		{"glDepthRange(f|x)", "f", "far"},
		{"glDepthRange(f|x)", "f", "farVal"},
		{"glDepthRange(f|x)", "n", "near"},
		{"glDepthRange(f|x)", "n", "nearVal"},
		{"glDrawArraysInstanced", "instancecount", "primcount"},
		{"glDrawElementsInstanced", "instancecount", "primcount"},
		{"glFogxv", "param", "params"},
		{"glFrustum(f|x)", "b", "bottom"},
		{"glFrustum(f|x)", "f", "far"},
		{"glFrustum(f|x)", "l", "left"},
		{"glFrustum(f|x)", "n", "near"},
		{"glFrustum(f|x)", "r", "right"},
		{"glFrustum(f|x)", "t", "top"},
		{"glGenSamplers", "count", "n"},
		{"glGetBooleanv", "data", "params"},
		{"glGetBufferParameteri(64|)v", "params", "data"},
		{"glGetBufferParameteri(64|)v", "pname", "value"},
		{"glGetFloatv", "data", "params"},
		{"glGetIntegerv", "data", "params"},
		{"glGetProgramBinary", "bufSize", "bufsize"},
		{"glGetProgramInfoLog", "bufSize", "maxLength"},
		{"glGetShaderInfoLog", "bufSize", "maxLength"},
		{"glGetShaderPrecisionFormat", "precisiontype", "precisionType"},
		{"glGetShaderPrecisionFormat", "shadertype", "shaderType"},
		{"glIsSampler", "sampler", "id"},
		{"glLightModelxv", "param", "params"},
		{"glMaterialxv", "param", "params"},
		{"glMultiTexCoord4x", "texture", "target"},
		{"glOrtho(f|x)", "b", "bottom"},
		{"glOrtho(f|x)", "f", "far"},
		{"glOrtho(f|x)", "l", "left"},
		{"glOrtho(f|x)", "n", "near"},
		{"glOrtho(f|x)", "r", "right"},
		{"glOrtho(f|x)", "t", "top"},
		{"glReadPixels", "pixels", "data"},
		{"glSamplerParameter(fv|iv)", "param", "params"},
		{"glShaderBinary", "binaryformat", "binaryFormat"},
		{"glShaderBinary", "count", "n"},
		{"glStencilOp", "fail", "sfail"},
		{"glStencilOp", "zfail", "dpfail"},
		{"glStencilOp", "zpass", "dppass"},
		{"glTex(Sub|)Image(2D|3D)", "pixels", "data"},
		{"glTexImage(2D|3D)", "internalformat", "internalFormat"},
		{"glVertexAttrib(1f|2f|3f|4f|I4i|I4ui)", "w", "v3"},
		{"glVertexAttrib(1f|2f|3f|4f|I4i|I4ui)", "x", "v0"},
		{"glVertexAttrib(1f|2f|3f|4f|I4i|I4ui)", "y", "v1"},
		{"glVertexAttrib(1f|2f|3f|4f|I4i|I4ui)", "z", "v2"},
		{"glBlendFunci", "src", "sfactor"},
		{"glBlendFunci", "dst", "dfactor"},
		{"glColorMaski", "index", "buf"},
		{"glColorMaski", "r", "red"},
		{"glColorMaski", "g", "green"},
		{"glColorMaski", "b", "blue"},
		{"glColorMaski", "a", "alpha"},
		{"glDebugMessageInsert", "buf", "message"},
		{"glDisablei", "target", "cap"},
		{"glDrawElementsInstancedBaseVertex", "instancecount", "primcount"},
		{"glEnablei", "target", "cap"},
		{"glGetObjectLabel", "bufSize", "bifSize"},
		{"glGetObjectPtrLabel", "bufSize", "bifSize"},
		{"glIsEnabledi", "target", "cap"},
		{"glSamplerParameterIiv", "param", "params"},
		{"glSamplerParameterIuiv", "param", "params"},
		{"glTexBuffer", "internalformat", "internalFormat"},
		{"glTexBufferRange", "internalformat", "internalFormat"},
	}
	for _, rename := range renames {
		if CompileRegexp("^(?:" + rename.cmdName + ")$").MatchString(cmdName) {
			if rename.oldName == paramName {
				return rename.newName
			}
		}
	}
	return paramName
}

func listContains(commaSeparatedList, value string) bool {
	if strings.Contains(commaSeparatedList, value) {
		for _, v := range strings.Split(commaSeparatedList, ",") {
			if v == value {
				return true
			}
		}
	}
	return false
}

// If the enum values can not be automatically extracted from the documentation,
// we use the hand-written table blow. (which can also be used for overrides)
func getAcceptedValues(cmdName, paramName string, apiVersion Version) ([]string, bool) {

	var GL_TEXTUREi = "GL_TEXTURE0,GL_TEXTURE1,GL_TEXTURE2,GL_TEXTURE3,GL_TEXTURE4,GL_TEXTURE5,GL_TEXTURE6,GL_TEXTURE7,GL_TEXTURE8,GL_TEXTURE9,GL_TEXTURE10,GL_TEXTURE11,GL_TEXTURE12,GL_TEXTURE13,GL_TEXTURE14,GL_TEXTURE15,GL_TEXTURE16,GL_TEXTURE17,GL_TEXTURE18,GL_TEXTURE19,GL_TEXTURE20,GL_TEXTURE21,GL_TEXTURE22,GL_TEXTURE23,GL_TEXTURE24,GL_TEXTURE25,GL_TEXTURE26,GL_TEXTURE27,GL_TEXTURE28,GL_TEXTURE29,GL_TEXTURE30,GL_TEXTURE31"
	var GL_CLIP_PLANEi = "GL_CLIP_PLANE0,GL_CLIP_PLANE1,GL_CLIP_PLANE2,GL_CLIP_PLANE3,GL_CLIP_PLANE4,GL_CLIP_PLANE5"
	var GL_LIGHTi = "GL_LIGHT0,GL_LIGHT1,GL_LIGHT2,GL_LIGHT3,GL_LIGHT4,GL_LIGHT5,GL_LIGHT6,GL_LIGHT7"
	var GL_COLOR_ATTACHMENTi = "GL_COLOR_ATTACHMENT0,GL_COLOR_ATTACHMENT1,GL_COLOR_ATTACHMENT2,GL_COLOR_ATTACHMENT3,GL_COLOR_ATTACHMENT4,GL_COLOR_ATTACHMENT5,GL_COLOR_ATTACHMENT6,GL_COLOR_ATTACHMENT7,GL_COLOR_ATTACHMENT8,GL_COLOR_ATTACHMENT9,GL_COLOR_ATTACHMENT10,GL_COLOR_ATTACHMENT11,GL_COLOR_ATTACHMENT12,GL_COLOR_ATTACHMENT13,GL_COLOR_ATTACHMENT14,GL_COLOR_ATTACHMENT15"

	// TODO: When is GL_SRC_ALPHA_SATURATE allowed?
	// TODO: GL_SAMPLE_MASK missing from glIsEnabled
	// TODO: When is GL_BACK,GL_DEPTH,GL_STENCIL valid for attachments?
	// TODO: glGetProgramResourceiv - is GL_TRANSFORM_FEEDBACK_BUFFER supported or not?
	// TODO: glGetActiveAttrib - accepts value of GL_IMPLEMENTATION_COLOR_READ_FORMAT_OES
	// TODO: Check glCopyImageSubData
	// TODO: glReadPixels; it seems to remove values
	// TODO: GL_FRAMEBUFFER_BINDING not allowed in 3.0 glGet
	// TODO: glGetInternalformativ

	var enumMetadata = []struct {
		// all field are comma separated values
		cmdName        string
		paramName      string
		apiVersion     string
		acceptedValues string
	}{
		{"glActiveTexture", "texture", "1.1,2.0,3.0,3.1,3.2", GL_TEXTUREi},
		{"glBeginTransformFeedback", "primitiveMode", "3.0,3.1,3.2", "GL_POINTS,GL_LINES,GL_TRIANGLES"},
		{"glBindBuffer", "target", "3.2", "GL_ARRAY_BUFFER,GL_ATOMIC_COUNTER_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_DISPATCH_INDIRECT_BUFFER,GL_DRAW_INDIRECT_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_SHADER_STORAGE_BUFFER,GL_TEXTURE_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glBindFramebuffer", "target", "3.0,3.1,3.2", "GL_DRAW_FRAMEBUFFER,GL_READ_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glBindImageTexture", "access", "3.1,3.2", "GL_READ_ONLY,GL_WRITE_ONLY,GL_READ_WRITE"},
		{"glBindImageTexture", "format", "3.1,3.2", "GL_RGBA32F,GL_RGBA16F,GL_R32F,GL_RGBA32UI,GL_RGBA16UI,GL_RGBA8UI,GL_R32UI,GL_RGBA32I,GL_RGBA16I,GL_RGBA8I,GL_R32I,GL_RGBA8,GL_RGBA8_SNORM"},
		{"glBlendFunc,glBlendFunci,glBlendFuncSeparate,glBlendFuncSeparatei", "dfactor,dfactorRGB,dfactorAlpha,sfactor,sfactorRGB,sfactorAlpha,dst,dstRGB,dstAlpha,src,srcRGB,srcAlpha", "2.0,3.0,3.1,3.2", "GL_ZERO,GL_ONE,GL_SRC_COLOR,GL_ONE_MINUS_SRC_COLOR,GL_DST_COLOR,GL_ONE_MINUS_DST_COLOR,GL_SRC_ALPHA,GL_ONE_MINUS_SRC_ALPHA,GL_DST_ALPHA,GL_ONE_MINUS_DST_ALPHA,GL_CONSTANT_COLOR,GL_ONE_MINUS_CONSTANT_COLOR,GL_CONSTANT_ALPHA,GL_ONE_MINUS_CONSTANT_ALPHA,GL_SRC_ALPHA_SATURATE"},
		{"glBufferData,glBufferSubData", "target", "3.2", "GL_ARRAY_BUFFER,GL_ATOMIC_COUNTER_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_DISPATCH_INDIRECT_BUFFER,GL_DRAW_INDIRECT_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_SHADER_STORAGE_BUFFER,GL_TEXTURE_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glCheckFramebufferStatus", "result", "2.0", "GL_FRAMEBUFFER_COMPLETE,GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT,GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS,GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT,GL_FRAMEBUFFER_UNSUPPORTED"},
		{"glCheckFramebufferStatus", "result", "3.0,3.1", "GL_FRAMEBUFFER_COMPLETE,GL_FRAMEBUFFER_UNDEFINED,GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT,GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT,GL_FRAMEBUFFER_UNSUPPORTED,GL_FRAMEBUFFER_INCOMPLETE_MULTISAMPLE"},
		{"glCheckFramebufferStatus", "target", "3.0,3.1,3.2", "GL_DRAW_FRAMEBUFFER,GL_READ_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glClearBufferfi", "buffer", "3.0,3.1,3.2", "GL_DEPTH_STENCIL"},
		{"glClearBufferfv", "buffer", "3.0,3.1,3.2", "GL_COLOR,GL_DEPTH"},
		{"glClearBufferiv", "buffer", "3.0,3.1,3.2", "GL_COLOR,GL_STENCIL"},
		{"glClearBufferuiv", "buffer", "3.0,3.1,3.2", "GL_COLOR"},
		{"glClientActiveTexture", "texture", "1.1", GL_TEXTUREi},
		{"glClientWaitSync", "result", "3.0,3.1", "GL_ALREADY_SIGNALED,GL_TIMEOUT_EXPIRED,GL_CONDITION_SATISFIED,GL_WAIT_FAILED"},
		{"glClipPlanef", "p", "1.1", GL_CLIP_PLANEi},
		{"glClipPlanex", "plane", "1.1", GL_CLIP_PLANEi},
		{"glColorPointer", "type", "1.1", "GL_UNSIGNED_BYTE,GL_FIXED,GL_FLOAT"},                                                                   // TODO
		{"glCompressedTexImage2D,glCompressedTexImage3D,glCompressedTexSubImage2D,glCompressedTexSubImage3D", "internalformat,format", "2.0", ""}, // provided by extensions only
		{"glCompressedTexImage2D,glCompressedTexImage3D,glCompressedTexSubImage2D,glCompressedTexSubImage3D", "internalformat,format", "3.0,3.1", "GL_COMPRESSED_R11_EAC,GL_COMPRESSED_SIGNED_R11_EAC,GL_COMPRESSED_RG11_EAC,GL_COMPRESSED_SIGNED_RG11_EAC,GL_COMPRESSED_RGB8_ETC2,GL_COMPRESSED_SRGB8_ETC2,GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_RGBA8_ETC2_EAC,GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC"},
		{"glCompressedTexImage2D,glCompressedTexImage3D,glCompressedTexSubImage2D,glCompressedTexSubImage3D", "internalformat,format", "3.2", "GL_COMPRESSED_R11_EAC,GL_COMPRESSED_SIGNED_R11_EAC,GL_COMPRESSED_RG11_EAC,GL_COMPRESSED_SIGNED_RG11_EAC,GL_COMPRESSED_RGB8_ETC2,GL_COMPRESSED_SRGB8_ETC2,GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_RGBA8_ETC2_EAC,GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC,GL_COMPRESSED_RGBA_ASTC_4x4,GL_COMPRESSED_RGBA_ASTC_5x4,GL_COMPRESSED_RGBA_ASTC_5x5,GL_COMPRESSED_RGBA_ASTC_6x5,GL_COMPRESSED_RGBA_ASTC_6x6,GL_COMPRESSED_RGBA_ASTC_8x5,GL_COMPRESSED_RGBA_ASTC_8x6,GL_COMPRESSED_RGBA_ASTC_8x8,GL_COMPRESSED_RGBA_ASTC_10x5,GL_COMPRESSED_RGBA_ASTC_10x6,GL_COMPRESSED_RGBA_ASTC_10x8,GL_COMPRESSED_RGBA_ASTC_10x10,GL_COMPRESSED_RGBA_ASTC_12x10,GL_COMPRESSED_RGBA_ASTC_12x12,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12"},
		{"glCompressedTexSubImage2D", "format", "1.1", "GL_PALETTE4_RGB8_OES,GL_PALETTE4_RGBA8_OES,GL_PALETTE4_R5_G6_B5_OES,GL_PALETTE4_RGBA4_OES,GL_PALETTE4_RGB5_A1_OES,GL_PALETTE8_RGB8_OES,GL_PALETTE8_RGBA8_OES,GL_PALETTE8_R5_G6_B5_OES,GL_PALETTE8_RGBA4_OES,GL_PALETTE8_RGB5_A1_OES"},
		{"glCompressedTexSubImage2D", "format", "2.0", ""},
		{"glCopyBufferSubData", "readTarget,writeTarget", "3.0,3.1", "GL_ARRAY_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glCopyBufferSubData", "readTarget,writeTarget", "3.2", "GL_ARRAY_BUFFER,GL_ATOMIC_COUNTER_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_DISPATCH_INDIRECT_BUFFER,GL_DRAW_INDIRECT_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_SHADER_STORAGE_BUFFER,GL_TEXTURE_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glCopyImageSubData", "srcTarget,dstTarget", "3.2", "GL_RENDERBUFFER,GL_TEXTURE_2D,GL_TEXTURE_2D_MULTISAMPLE,GL_TEXTURE_2D_MULTISAMPLE_ARRAY,GL_TEXTURE_3D,GL_TEXTURE_2D_ARRAY,GL_TEXTURE_CUBE_MAP,GL_TEXTURE_CUBE_MAP_ARRAY"},
		{"glCreateShader", "type", "3.1", "GL_VERTEX_SHADER,GL_FRAGMENT_SHADER,GL_COMPUTE_SHADER"},
		{"glCreateShaderProgramv", "type", "3.1", "GL_VERTEX_SHADER,GL_FRAGMENT_SHADER,GL_COMPUTE_SHADER"},                                                                     // TODO
		{"glCreateShaderProgramv", "type", "3.2", "GL_COMPUTE_SHADER,GL_VERTEX_SHADER,GL_TESS_CONTROL_SHADER,GL_TESS_EVALUATION_SHADER,GL_GEOMETRY_SHADER,GL_FRAGMENT_SHADER"}, // TODO
		{"glDebugMessageControl", "source", "3.2", "GL_DEBUG_SOURCE_API,GL_DEBUG_SOURCE_WINDOW_SYSTEM_,GL_DEBUG_SOURCE_SHADER_COMPILER,GL_DEBUG_SOURCE_THIRD_PARTY,GL_DEBUG_SOURCE_APPLICATION,GL_DEBUG_SOURCE_OTHER,GL_DONT_CARE"},
		{"glDebugMessageControl", "type", "3.2", "GL_DEBUG_TYPE_ERROR,GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR,GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR,GL_DEBUG_TYPE_PORTABILITY,GL_DEBUG_TYPE_PERFORMANCE,GL_DEBUG_TYPE_MARKER,GL_DEBUG_TYPE_PUSH_GROUP,GL_DEBUG_TYPE_POP_GROUP,GL_DEBUG_TYPE_OTHER,GL_DONT_CARE"},
		{"glDebugMessageControl", "severity", "3.2", "GL_DEBUG_SEVERITY_LOW,GL_DEBUG_SEVERITY_MEDIUM,GL_DEBUG_SEVERITY_HIGH,GL_DEBUG_SEVERITY_NOTIFICATION,GL_DONT_CARE"},
		{"glDebugMessageInsert", "source", "3.2", "GL_DEBUG_SOURCE_APPLICATION,GL_DEBUG_SOURCE_THIRD_PARTY"},
		{"glDebugMessageInsert", "type", "3.2", "GL_DEBUG_TYPE_ERROR,GL_DEBUG_TYPE_DEPRECATED_BEHAVIOR,GL_DEBUG_TYPE_UNDEFINED_BEHAVIOR,GL_DEBUG_TYPE_PORTABILITY,GL_DEBUG_TYPE_PERFORMANCE,GL_DEBUG_TYPE_MARKER,GL_DEBUG_TYPE_PUSH_GROUP,GL_DEBUG_TYPE_POP_GROUP,GL_DEBUG_TYPE_OTHER"},
		{"glDebugMessageInsert", "severity", "3.2", "GL_DEBUG_SEVERITY_LOW,GL_DEBUG_SEVERITY_MEDIUM,GL_DEBUG_SEVERITY_HIGH,GL_DEBUG_SEVERITY_NOTIFICATION"},
		{"glDrawBuffers", "bufs", "3.0,3.1", "GL_NONE,GL_BACK," + GL_COLOR_ATTACHMENTi},
		{"glDrawElementsIndirect", "type", "3.1,3.2", "GL_UNSIGNED_BYTE,GL_UNSIGNED_SHORT,GL_UNSIGNED_INT"},
		{"glEnable,glDisable", "cap", "1.1", "GL_ALPHA_TEST,GL_BLEND,GL_COLOR_LOGIC_OP,GL_COLOR_MATERIAL,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_FOG,GL_LIGHTING,GL_LINE_SMOOTH,GL_MULTISAMPLE,GL_NORMALIZE,GL_POINT_SMOOTH,GL_POINT_SPRITE_OES,GL_POLYGON_OFFSET_FILL,GL_RESCALE_NORMAL,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_ALPHA_TO_ONE,GL_SAMPLE_COVERAGE,GL_SCISSOR_TEST,GL_STENCIL_TEST,GL_TEXTURE_2D," + GL_CLIP_PLANEi + "," + GL_LIGHTi},
		{"glEnable,glDisable", "cap", "2.0", "GL_BLEND,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_POLYGON_OFFSET_FILL,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_COVERAGE,GL_SCISSOR_TEST,GL_STENCIL_TEST"},
		{"glEnable,glDisable", "cap", "3.0", "GL_BLEND,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_POLYGON_OFFSET_FILL,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_RASTERIZER_DISCARD,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_COVERAGE,GL_SCISSOR_TEST,GL_STENCIL_TEST"},
		{"glEnable,glDisable", "cap", "3.1", "GL_BLEND,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_POLYGON_OFFSET_FILL,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_RASTERIZER_DISCARD,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_COVERAGE,GL_SAMPLE_MASK,GL_SCISSOR_TEST,GL_STENCIL_TEST"},
		{"glEnable,glDisable,glEnablei,glDisablei,glIsEnabled,glIsEnabledi", "cap,target", "3.2", "GL_BLEND,GL_CULL_FACE,GL_DEBUG_OUTPUT,GL_DEBUG_OUTPUT_SYNCHRONOUS,GL_DEPTH_TEST,GL_DITHER,GL_POLYGON_OFFSET_FILL,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_RASTERIZER_DISCARD,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_COVERAGE,GL_SAMPLE_MASK,GL_SCISSOR_TEST,GL_STENCIL_TEST"},
		{"glFlushMappedBufferRange", "target", "3.2", "GL_ARRAY_BUFFER,GL_ATOMIC_COUNTER_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_DISPATCH_INDIRECT_BUFFER,GL_DRAW_INDIRECT_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_SHADER_STORAGE_BUFFER,GL_TEXTURE_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glFramebufferParameteri", "pname", "3.1", "GL_FRAMEBUFFER_DEFAULT_WIDTH,GL_FRAMEBUFFER_DEFAULT_HEIGHT,GL_FRAMEBUFFER_DEFAULT_SAMPLES,GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS"},
		{"glFramebufferParameteri", "pname", "3.2", "GL_FRAMEBUFFER_DEFAULT_WIDTH,GL_FRAMEBUFFER_DEFAULT_HEIGHT,GL_FRAMEBUFFER_DEFAULT_SAMPLES,GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS,GL_FRAMEBUFFER_DEFAULT_LAYERS"},
		{"glFramebufferParameteri", "target", "3.1,3.2", "GL_READ_FRAMEBUFFER,GL_DRAW_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glFramebufferRenderbuffer", "attachment", "3.0,3.1,3.2", "GL_DEPTH_ATTACHMENT,GL_STENCIL_ATTACHMENT,GL_DEPTH_STENCIL_ATTACHMENT," + GL_COLOR_ATTACHMENTi},
		{"glFramebufferRenderbuffer", "renderbuffertarget", "3.0,3.1,3.2", "GL_RENDERBUFFER"},
		{"glFramebufferRenderbuffer", "target", "3.0,3.1,3.2", "GL_DRAW_FRAMEBUFFER,GL_READ_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glFramebufferTexture,glFramebufferTexture2D", "attachment", "3.0,3.1,3.2", "GL_DEPTH_ATTACHMENT,GL_STENCIL_ATTACHMENT,GL_DEPTH_STENCIL_ATTACHMENT," + GL_COLOR_ATTACHMENTi},
		{"glFramebufferTexture,glFramebufferTexture2D", "target", "3.0,3.1,3.2", "GL_DRAW_FRAMEBUFFER,GL_READ_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glFramebufferTexture2D", "textarget", "3.0", "GL_TEXTURE_2D,GL_TEXTURE_CUBE_MAP_POSITIVE_X,GL_TEXTURE_CUBE_MAP_NEGATIVE_X,GL_TEXTURE_CUBE_MAP_POSITIVE_Y,GL_TEXTURE_CUBE_MAP_NEGATIVE_Y,GL_TEXTURE_CUBE_MAP_POSITIVE_Z,GL_TEXTURE_CUBE_MAP_NEGATIVE_Z"},
		{"glFramebufferTexture2D", "textarget", "3.1,3.2", "GL_TEXTURE_2D,GL_TEXTURE_CUBE_MAP_POSITIVE_X,GL_TEXTURE_CUBE_MAP_NEGATIVE_X,GL_TEXTURE_CUBE_MAP_POSITIVE_Y,GL_TEXTURE_CUBE_MAP_NEGATIVE_Y,GL_TEXTURE_CUBE_MAP_POSITIVE_Z,GL_TEXTURE_CUBE_MAP_NEGATIVE_Z,GL_TEXTURE_2D_MULTISAMPLE"},
		{"glFramebufferTextureLayer", "attachment", "3.0,3.1,3.2", "GL_DEPTH_ATTACHMENT,GL_STENCIL_ATTACHMENT,GL_DEPTH_STENCIL_ATTACHMENT," + GL_COLOR_ATTACHMENTi},
		{"glFramebufferTextureLayer", "target", "3.0,3.1,3.2", "GL_DRAW_FRAMEBUFFER,GL_READ_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glGetActiveAttrib", "type", "2.0", "GL_FLOAT,GL_FLOAT_VEC2,GL_FLOAT_VEC3,GL_FLOAT_VEC4,GL_FLOAT_MAT2,GL_FLOAT_MAT3,GL_FLOAT_MAT4"},
		{"glGetActiveAttrib,glGetTransformFeedbackVarying", "type", "3.0,3.1", "GL_FLOAT,GL_FLOAT_VEC2,GL_FLOAT_VEC3,GL_FLOAT_VEC4,GL_FLOAT_MAT2,GL_FLOAT_MAT3,GL_FLOAT_MAT4,GL_FLOAT_MAT2x3,GL_FLOAT_MAT2x4,GL_FLOAT_MAT3x2,GL_FLOAT_MAT3x4,GL_FLOAT_MAT4x2,GL_FLOAT_MAT4x3,GL_INT,GL_INT_VEC2,GL_INT_VEC3,GL_INT_VEC4,GL_UNSIGNED_INT,GL_UNSIGNED_INT_VEC2,GL_UNSIGNED_INT_VEC3,GL_UNSIGNED_INT_VEC4"},
		{"glGetActiveUniform", "type", "2.0", "GL_FLOAT,GL_FLOAT_VEC2,GL_FLOAT_VEC3,GL_FLOAT_VEC4,GL_INT,GL_INT_VEC2,GL_INT_VEC3,GL_INT_VEC4,GL_BOOL,GL_BOOL_VEC2,GL_BOOL_VEC3,GL_BOOL_VEC4,GL_FLOAT_MAT2,GL_FLOAT_MAT3,GL_FLOAT_MAT4,GL_SAMPLER_2D,GL_SAMPLER_CUBE"},
		{"glGetActiveUniform", "type", "3.0,3.1", "GL_FLOAT,GL_FLOAT_VEC2,GL_FLOAT_VEC3,GL_FLOAT_VEC4,GL_INT,GL_INT_VEC2,GL_INT_VEC3,GL_INT_VEC4,GL_UNSIGNED_INT,GL_UNSIGNED_INT_VEC2,GL_UNSIGNED_INT_VEC3,GL_UNSIGNED_INT_VEC4,GL_BOOL,GL_BOOL_VEC2,GL_BOOL_VEC3,GL_BOOL_VEC4,GL_FLOAT_MAT2,GL_FLOAT_MAT3,GL_FLOAT_MAT4,GL_FLOAT_MAT2x3,GL_FLOAT_MAT2x4,GL_FLOAT_MAT3x2,GL_FLOAT_MAT3x4,GL_FLOAT_MAT4x2,GL_FLOAT_MAT4x3,GL_SAMPLER_2D,GL_SAMPLER_3D,GL_SAMPLER_CUBE,GL_SAMPLER_2D_SHADOW,GL_SAMPLER_2D_ARRAY,GL_SAMPLER_2D_ARRAY_SHADOW,GL_SAMPLER_CUBE_SHADOW,GL_INT_SAMPLER_2D,GL_INT_SAMPLER_3D,GL_INT_SAMPLER_CUBE,GL_INT_SAMPLER_2D_ARRAY,GL_UNSIGNED_INT_SAMPLER_2D,GL_UNSIGNED_INT_SAMPLER_3D,GL_UNSIGNED_INT_SAMPLER_CUBE,GL_UNSIGNED_INT_SAMPLER_2D_ARRAY"},
		{"glGetActiveUniformBlockiv", "pname", "3.0,3.1,3.2", "GL_UNIFORM_BLOCK_BINDING,GL_UNIFORM_BLOCK_DATA_SIZE,GL_UNIFORM_BLOCK_NAME_LENGTH,GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS,GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES,GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER,GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER"},
		{"glGetActiveUniformsiv", "pname", "3.0,3.1,3.2", "GL_UNIFORM_TYPE,GL_UNIFORM_SIZE,GL_UNIFORM_NAME_LENGTH,GL_UNIFORM_BLOCK_INDEX,GL_UNIFORM_OFFSET,GL_UNIFORM_ARRAY_STRIDE,GL_UNIFORM_MATRIX_STRIDE,GL_UNIFORM_IS_ROW_MAJOR"},
		{"glGetBooleanv,glGetBooleani_v,glGetFloatv,glGetInteger64i_v,glGetInteger64v,glGetIntegeri_v,glGetIntegerv", "pname,target", "3.0", "GL_ACTIVE_TEXTURE,GL_ALIASED_LINE_WIDTH_RANGE,GL_ALIASED_POINT_SIZE_RANGE,GL_ALPHA_BITS,GL_ARRAY_BUFFER_BINDING,GL_BLEND,GL_BLEND_COLOR,GL_BLEND_DST_ALPHA,GL_BLEND_DST_RGB,GL_BLEND_EQUATION_ALPHA,GL_BLEND_EQUATION_RGB,GL_BLEND_SRC_ALPHA,GL_BLEND_SRC_RGB,GL_BLUE_BITS,GL_COLOR_CLEAR_VALUE,GL_COLOR_WRITEMASK,GL_COMPRESSED_TEXTURE_FORMATS,GL_COPY_READ_BUFFER_BINDING,GL_COPY_WRITE_BUFFER_BINDING,GL_CULL_FACE,GL_CULL_FACE_MODE,GL_CURRENT_PROGRAM,GL_DEPTH_BITS,GL_DEPTH_CLEAR_VALUE,GL_DEPTH_FUNC,GL_DEPTH_RANGE,GL_DEPTH_TEST,GL_DEPTH_WRITEMASK,GL_DITHER,GL_DRAW_BUFFER,GL_DRAW_FRAMEBUFFER_BINDING,GL_ELEMENT_ARRAY_BUFFER_BINDING,GL_FRAGMENT_SHADER_DERIVATIVE_HINT,GL_FRONT_FACE,GL_GENERATE_MIPMAP_HINT,GL_GREEN_BITS,GL_IMPLEMENTATION_COLOR_READ_FORMAT,GL_IMPLEMENTATION_COLOR_READ_TYPE,GL_LINE_WIDTH,GL_MAJOR_VERSION,GL_MAX_3D_TEXTURE_SIZE,GL_MAX_ARRAY_TEXTURE_LAYERS,GL_MAX_COLOR_ATTACHMENTS,GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS,GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS,GL_MAX_COMBINED_UNIFORM_BLOCKS,GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS,GL_MAX_CUBE_MAP_TEXTURE_SIZE,GL_MAX_DRAW_BUFFERS,GL_MAX_ELEMENT_INDEX,GL_MAX_ELEMENTS_INDICES,GL_MAX_ELEMENTS_VERTICES,GL_MAX_FRAGMENT_INPUT_COMPONENTS,GL_MAX_FRAGMENT_UNIFORM_BLOCKS,GL_MAX_FRAGMENT_UNIFORM_COMPONENTS,GL_MAX_FRAGMENT_UNIFORM_VECTORS,GL_MAX_PROGRAM_TEXEL_OFFSET,GL_MAX_RENDERBUFFER_SIZE,GL_MAX_SAMPLES,GL_MAX_SERVER_WAIT_TIMEOUT,GL_MAX_TEXTURE_IMAGE_UNITS,GL_MAX_TEXTURE_LOD_BIAS,GL_MAX_TEXTURE_SIZE,GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS,GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS,GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS,GL_MAX_UNIFORM_BLOCK_SIZE,GL_MAX_UNIFORM_BUFFER_BINDINGS,GL_MAX_VARYING_COMPONENTS,GL_MAX_VARYING_VECTORS,GL_MAX_VERTEX_ATTRIBS,GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS,GL_MAX_VERTEX_OUTPUT_COMPONENTS,GL_MAX_VERTEX_UNIFORM_BLOCKS,GL_MAX_VERTEX_UNIFORM_COMPONENTS,GL_MAX_VERTEX_UNIFORM_VECTORS,GL_MAX_VIEWPORT_DIMS,GL_MIN_PROGRAM_TEXEL_OFFSET,GL_MINOR_VERSION,GL_NUM_COMPRESSED_TEXTURE_FORMATS,GL_NUM_EXTENSIONS,GL_NUM_PROGRAM_BINARY_FORMATS,GL_NUM_SHADER_BINARY_FORMATS,GL_PACK_ALIGNMENT,GL_PACK_ROW_LENGTH,GL_PACK_SKIP_PIXELS,GL_PACK_SKIP_ROWS,GL_PIXEL_PACK_BUFFER_BINDING,GL_PIXEL_UNPACK_BUFFER_BINDING,GL_POLYGON_OFFSET_FACTOR,GL_POLYGON_OFFSET_FILL,GL_POLYGON_OFFSET_UNITS,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_PROGRAM_BINARY_FORMATS,GL_RASTERIZER_DISCARD,GL_READ_BUFFER,GL_READ_FRAMEBUFFER_BINDING,GL_RED_BITS,GL_RENDERBUFFER_BINDING,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_BUFFERS,GL_SAMPLE_COVERAGE,GL_SAMPLE_COVERAGE_INVERT,GL_SAMPLE_COVERAGE_VALUE,GL_SAMPLER_BINDING,GL_SAMPLES,GL_SCISSOR_BOX,GL_SCISSOR_TEST,GL_SHADER_BINARY_FORMATS,GL_SHADER_COMPILER,GL_STENCIL_BACK_FAIL,GL_STENCIL_BACK_FUNC,GL_STENCIL_BACK_PASS_DEPTH_FAIL,GL_STENCIL_BACK_PASS_DEPTH_PASS,GL_STENCIL_BACK_REF,GL_STENCIL_BACK_VALUE_MASK,GL_STENCIL_BACK_WRITEMASK,GL_STENCIL_BITS,GL_STENCIL_CLEAR_VALUE,GL_STENCIL_FAIL,GL_STENCIL_FUNC,GL_STENCIL_PASS_DEPTH_FAIL,GL_STENCIL_PASS_DEPTH_PASS,GL_STENCIL_REF,GL_STENCIL_TEST,GL_STENCIL_VALUE_MASK,GL_STENCIL_WRITEMASK,GL_SUBPIXEL_BITS,GL_TEXTURE_BINDING_2D,GL_TEXTURE_BINDING_2D_ARRAY,GL_TEXTURE_BINDING_3D,GL_TEXTURE_BINDING_CUBE_MAP,GL_TRANSFORM_FEEDBACK_BINDING,GL_TRANSFORM_FEEDBACK_ACTIVE,GL_TRANSFORM_FEEDBACK_BUFFER_BINDING,GL_TRANSFORM_FEEDBACK_PAUSED,GL_TRANSFORM_FEEDBACK_BUFFER_SIZE,GL_TRANSFORM_FEEDBACK_BUFFER_START,GL_UNIFORM_BUFFER_BINDING,GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT,GL_UNIFORM_BUFFER_SIZE,GL_UNIFORM_BUFFER_START,GL_UNPACK_ALIGNMENT,GL_UNPACK_IMAGE_HEIGHT,GL_UNPACK_ROW_LENGTH,GL_UNPACK_SKIP_IMAGES,GL_UNPACK_SKIP_PIXELS,GL_UNPACK_SKIP_ROWS,GL_VERTEX_ARRAY_BINDING,GL_VIEWPORT,GL_FRAMEBUFFER_BINDING"},
		{"glGetBooleanv,glGetBooleani_v,glGetFloatv,glGetInteger64i_v,glGetInteger64v,glGetIntegeri_v,glGetIntegerv", "pname,target", "3.1", "GL_ACTIVE_TEXTURE,GL_ALIASED_LINE_WIDTH_RANGE,GL_ALIASED_POINT_SIZE_RANGE,GL_ALPHA_BITS,GL_ARRAY_BUFFER_BINDING,GL_BLEND,GL_BLEND_COLOR,GL_BLEND_DST_ALPHA,GL_BLEND_DST_RGB,GL_BLEND_EQUATION_ALPHA,GL_BLEND_EQUATION_RGB,GL_BLEND_SRC_ALPHA,GL_BLEND_SRC_RGB,GL_BLUE_BITS,GL_COLOR_CLEAR_VALUE,GL_COLOR_WRITEMASK,GL_COMPRESSED_TEXTURE_FORMATS,GL_COPY_READ_BUFFER_BINDING,GL_COPY_WRITE_BUFFER_BINDING,GL_CULL_FACE,GL_CULL_FACE_MODE,GL_CURRENT_PROGRAM,GL_DEPTH_BITS,GL_DEPTH_CLEAR_VALUE,GL_DEPTH_FUNC,GL_DEPTH_RANGE,GL_DEPTH_TEST,GL_DEPTH_WRITEMASK,GL_DISPATCH_INDIRECT_BUFFER_BINDING,GL_DITHER,GL_DRAW_BUFFER,GL_DRAW_FRAMEBUFFER_BINDING,GL_ELEMENT_ARRAY_BUFFER_BINDING,GL_FRAGMENT_SHADER_DERIVATIVE_HINT,GL_FRONT_FACE,GL_GENERATE_MIPMAP_HINT,GL_GREEN_BITS,GL_IMAGE_BINDING_LAYERED,GL_IMPLEMENTATION_COLOR_READ_FORMAT,GL_IMPLEMENTATION_COLOR_READ_TYPE,GL_LINE_WIDTH,GL_MAJOR_VERSION,GL_MAX_3D_TEXTURE_SIZE,GL_MAX_ARRAY_TEXTURE_LAYERS,GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS,GL_MAX_COLOR_ATTACHMENTS,GL_MAX_COLOR_TEXTURE_SAMPLES,GL_MAX_COMBINED_ATOMIC_COUNTERS,GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS,GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS,GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS,GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS,GL_MAX_COMBINED_UNIFORM_BLOCKS,GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS,GL_MAX_COMPUTE_ATOMIC_COUNTERS,GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS,GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS,GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS,GL_MAX_COMPUTE_UNIFORM_BLOCKS,GL_MAX_COMPUTE_UNIFORM_COMPONENTS,GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS,GL_MAX_COMPUTE_WORK_GROUP_COUNT,GL_MAX_COMPUTE_WORK_GROUP_SIZE,GL_MAX_CUBE_MAP_TEXTURE_SIZE,GL_MAX_DRAW_BUFFERS,GL_MAX_ELEMENT_INDEX,GL_MAX_ELEMENTS_INDICES,GL_MAX_ELEMENTS_VERTICES,GL_MAX_FRAGMENT_ATOMIC_COUNTERS,GL_MAX_FRAGMENT_INPUT_COMPONENTS,GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS,GL_MAX_FRAGMENT_UNIFORM_BLOCKS,GL_MAX_FRAGMENT_UNIFORM_COMPONENTS,GL_MAX_FRAGMENT_UNIFORM_VECTORS,GL_MAX_FRAMEBUFFER_HEIGHT,GL_MAX_FRAMEBUFFER_SAMPLES,GL_MAX_FRAMEBUFFER_WIDTH,GL_MAX_INTEGER_SAMPLES,GL_MAX_PROGRAM_TEXEL_OFFSET,GL_MAX_RENDERBUFFER_SIZE,GL_MAX_SAMPLE_MASK_WORDS,GL_MAX_SAMPLES,GL_MAX_SERVER_WAIT_TIMEOUT,GL_MAX_SHADER_STORAGE_BLOCK_SIZE,GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS,GL_MAX_TEXTURE_IMAGE_UNITS,GL_MAX_TEXTURE_LOD_BIAS,GL_MAX_TEXTURE_SIZE,GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS,GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS,GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS,GL_MAX_UNIFORM_BLOCK_SIZE,GL_MAX_UNIFORM_BUFFER_BINDINGS,GL_MAX_UNIFORM_LOCATIONS,GL_MAX_VARYING_COMPONENTS,GL_MAX_VARYING_VECTORS,GL_MAX_VERTEX_ATOMIC_COUNTERS,GL_MAX_VERTEX_ATTRIB_BINDINGS,GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET,GL_MAX_VERTEX_ATTRIBS,GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS,GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS,GL_MAX_VERTEX_OUTPUT_COMPONENTS,GL_MAX_VERTEX_UNIFORM_BLOCKS,GL_MAX_VERTEX_UNIFORM_COMPONENTS,GL_MAX_VERTEX_UNIFORM_VECTORS,GL_MAX_VIEWPORT_DIMS,GL_MIN_PROGRAM_TEXEL_OFFSET,GL_MINOR_VERSION,GL_NUM_COMPRESSED_TEXTURE_FORMATS,GL_NUM_EXTENSIONS,GL_NUM_PROGRAM_BINARY_FORMATS,GL_NUM_SHADER_BINARY_FORMATS,GL_PACK_ALIGNMENT,GL_PACK_ROW_LENGTH,GL_PACK_SKIP_PIXELS,GL_PACK_SKIP_ROWS,GL_PIXEL_PACK_BUFFER_BINDING,GL_PIXEL_UNPACK_BUFFER_BINDING,GL_POLYGON_OFFSET_FACTOR,GL_POLYGON_OFFSET_FILL,GL_POLYGON_OFFSET_UNITS,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_PROGRAM_BINARY_FORMATS,GL_PROGRAM_PIPELINE_BINDING,GL_RASTERIZER_DISCARD,GL_READ_BUFFER,GL_READ_FRAMEBUFFER_BINDING,GL_RED_BITS,GL_RENDERBUFFER_BINDING,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_BUFFERS,GL_SAMPLE_COVERAGE,GL_SAMPLE_COVERAGE_INVERT,GL_SAMPLE_COVERAGE_VALUE,GL_SAMPLER_BINDING,GL_SAMPLES,GL_SCISSOR_BOX,GL_SCISSOR_TEST,GL_SHADER_BINARY_FORMATS,GL_SHADER_COMPILER,GL_SHADER_STORAGE_BUFFER_BINDING,GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT,GL_SHADER_STORAGE_BUFFER_SIZE,GL_SHADER_STORAGE_BUFFER_START,GL_STENCIL_BACK_FAIL,GL_STENCIL_BACK_FUNC,GL_STENCIL_BACK_PASS_DEPTH_FAIL,GL_STENCIL_BACK_PASS_DEPTH_PASS,GL_STENCIL_BACK_REF,GL_STENCIL_BACK_VALUE_MASK,GL_STENCIL_BACK_WRITEMASK,GL_STENCIL_BITS,GL_STENCIL_CLEAR_VALUE,GL_STENCIL_FAIL,GL_STENCIL_FUNC,GL_STENCIL_PASS_DEPTH_FAIL,GL_STENCIL_PASS_DEPTH_PASS,GL_STENCIL_REF,GL_STENCIL_TEST,GL_STENCIL_VALUE_MASK,GL_STENCIL_WRITEMASK,GL_SUBPIXEL_BITS,GL_TEXTURE_BINDING_2D,GL_TEXTURE_BINDING_2D_ARRAY,GL_TEXTURE_BINDING_3D,GL_TEXTURE_BINDING_CUBE_MAP,GL_TEXTURE_BINDING_2D_MULTISAMPLE,GL_TRANSFORM_FEEDBACK_BINDING,GL_TRANSFORM_FEEDBACK_ACTIVE,GL_TRANSFORM_FEEDBACK_BUFFER_BINDING,GL_TRANSFORM_FEEDBACK_PAUSED,GL_TRANSFORM_FEEDBACK_BUFFER_SIZE,GL_TRANSFORM_FEEDBACK_BUFFER_START,GL_UNIFORM_BUFFER_BINDING,GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT,GL_UNIFORM_BUFFER_SIZE,GL_UNIFORM_BUFFER_START,GL_UNPACK_ALIGNMENT,GL_UNPACK_IMAGE_HEIGHT,GL_UNPACK_ROW_LENGTH,GL_UNPACK_SKIP_IMAGES,GL_UNPACK_SKIP_PIXELS,GL_UNPACK_SKIP_ROWS,GL_VERTEX_ARRAY_BINDING,GL_VERTEX_BINDING_DIVISOR,GL_VERTEX_BINDING_OFFSET,GL_VERTEX_BINDING_STRIDE,GL_VIEWPORT,GL_FRAMEBUFFER_BINDING"},
		{"glGetBooleanv,glGetFloatv,glGetIntegerv", "pname", "2.0", "GL_ACTIVE_TEXTURE,GL_ALIASED_LINE_WIDTH_RANGE,GL_ALIASED_POINT_SIZE_RANGE,GL_ALPHA_BITS,GL_ARRAY_BUFFER_BINDING,GL_BLEND,GL_BLEND_COLOR,GL_BLEND_DST_ALPHA,GL_BLEND_DST_RGB,GL_BLEND_EQUATION_ALPHA,GL_BLEND_EQUATION_RGB,GL_BLEND_SRC_ALPHA,GL_BLEND_SRC_RGB,GL_BLUE_BITS,GL_COLOR_CLEAR_VALUE,GL_COLOR_WRITEMASK,GL_COMPRESSED_TEXTURE_FORMATS,GL_CULL_FACE,GL_CULL_FACE_MODE,GL_CURRENT_PROGRAM,GL_DEPTH_BITS,GL_DEPTH_CLEAR_VALUE,GL_DEPTH_FUNC,GL_DEPTH_RANGE,GL_DEPTH_TEST,GL_DEPTH_WRITEMASK,GL_DITHER,GL_ELEMENT_ARRAY_BUFFER_BINDING,GL_FRAMEBUFFER_BINDING,GL_FRONT_FACE,GL_GENERATE_MIPMAP_HINT,GL_GREEN_BITS,GL_IMPLEMENTATION_COLOR_READ_FORMAT,GL_IMPLEMENTATION_COLOR_READ_TYPE,GL_LINE_WIDTH,GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS,GL_MAX_CUBE_MAP_TEXTURE_SIZE,GL_MAX_FRAGMENT_UNIFORM_VECTORS,GL_MAX_RENDERBUFFER_SIZE,GL_MAX_TEXTURE_IMAGE_UNITS,GL_MAX_TEXTURE_SIZE,GL_MAX_VARYING_VECTORS,GL_MAX_VERTEX_ATTRIBS,GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS,GL_MAX_VERTEX_UNIFORM_VECTORS,GL_MAX_VIEWPORT_DIMS,GL_NUM_COMPRESSED_TEXTURE_FORMATS,GL_NUM_SHADER_BINARY_FORMATS,GL_PACK_ALIGNMENT,GL_POLYGON_OFFSET_FACTOR,GL_POLYGON_OFFSET_FILL,GL_POLYGON_OFFSET_UNITS,GL_RED_BITS,GL_RENDERBUFFER_BINDING,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_BUFFERS,GL_SAMPLE_COVERAGE,GL_SAMPLE_COVERAGE_INVERT,GL_SAMPLE_COVERAGE_VALUE,GL_SAMPLES,GL_SCISSOR_BOX,GL_SCISSOR_TEST,GL_SHADER_BINARY_FORMATS,GL_SHADER_COMPILER,GL_STENCIL_BACK_FAIL,GL_STENCIL_BACK_FUNC,GL_STENCIL_BACK_PASS_DEPTH_FAIL,GL_STENCIL_BACK_PASS_DEPTH_PASS,GL_STENCIL_BACK_REF,GL_STENCIL_BACK_VALUE_MASK,GL_STENCIL_BACK_WRITEMASK,GL_STENCIL_BITS,GL_STENCIL_CLEAR_VALUE,GL_STENCIL_FAIL,GL_STENCIL_FUNC,GL_STENCIL_PASS_DEPTH_FAIL,GL_STENCIL_PASS_DEPTH_PASS,GL_STENCIL_REF,GL_STENCIL_TEST,GL_STENCIL_VALUE_MASK,GL_STENCIL_WRITEMASK,GL_SUBPIXEL_BITS,GL_TEXTURE_BINDING_2D,GL_TEXTURE_BINDING_CUBE_MAP,GL_UNPACK_ALIGNMENT,GL_VIEWPORT"},
		{"glGetBooleanv,glGetFloatv,glGetIntegerv,glGetInteger64v,glGetBooleani_v,glGetIntegeri_v,glGetInteger64i_v", "pname,target", "3.2", "GL_ACTIVE_TEXTURE,GL_ALIASED_LINE_WIDTH_RANGE,GL_ALIASED_POINT_SIZE_RANGE,GL_ALPHA_BITS,GL_ARRAY_BUFFER_BINDING,GL_BLEND,GL_BLEND_COLOR,GL_BLEND_DST_ALPHA,GL_BLEND_DST_RGB,GL_BLEND_EQUATION_ALPHA,GL_BLEND_EQUATION_RGB,GL_BLEND_SRC_ALPHA,GL_BLEND_SRC_RGB,GL_BLUE_BITS,GL_COLOR_CLEAR_VALUE,GL_COLOR_WRITEMASK,GL_COMPRESSED_TEXTURE_FORMATS,GL_CONTEXT_FLAGS,GL_CONTEXT_ROBUST_ACCESS,GL_COPY_READ_BUFFER_BINDING,GL_COPY_WRITE_BUFFER_BINDING,GL_CULL_FACE,GL_CULL_FACE_MODE,GL_CURRENT_PROGRAM,GL_DEBUG_GROUP_STACK_DEPTH,GL_DEBUG_LOGGED_MESSAGES,GL_DEBUG_NEXT_LOGGED_MESSAGE_LENGTH,GL_DEPTH_BITS,GL_DEPTH_CLEAR_VALUE,GL_DEPTH_FUNC,GL_DEPTH_RANGE,GL_DEPTH_TEST,GL_DEPTH_WRITEMASK,GL_DISPATCH_INDIRECT_BUFFER_BINDING,GL_DITHER,GL_DRAW_BUFFER,GL_DRAW_FRAMEBUFFER_BINDING,GL_ELEMENT_ARRAY_BUFFER_BINDING,GL_FRAGMENT_INTERPOLATION_OFFSET_BITS,GL_FRAGMENT_SHADER_DERIVATIVE_HINT,GL_FRONT_FACE,GL_GENERATE_MIPMAP_HINT,GL_GREEN_BITS,GL_IMAGE_BINDING_LAYERED,GL_IMPLEMENTATION_COLOR_READ_FORMAT,GL_IMPLEMENTATION_COLOR_READ_TYPE,GL_LAYER_PROVOKING_VERTEX,GL_LINE_WIDTH,GL_MAJOR_VERSION,GL_MAX_3D_TEXTURE_SIZE,GL_MAX_ARRAY_TEXTURE_LAYERS,GL_MAX_ATOMIC_COUNTER_BUFFER_BINDINGS,GL_MAX_COLOR_ATTACHMENTS,GL_MAX_COLOR_TEXTURE_SAMPLES,GL_MAX_COMBINED_ATOMIC_COUNTERS,GL_MAX_COMBINED_COMPUTE_UNIFORM_COMPONENTS,GL_MAX_COMBINED_FRAGMENT_UNIFORM_COMPONENTS,GL_MAX_COMBINED_GEOMETRY_UNIFORM_COMPONENTS,GL_MAX_COMBINED_SHADER_STORAGE_BLOCKS,GL_MAX_COMBINED_TESS_CONTROL_UNIFORM_COMPONENTS,GL_MAX_COMBINED_TESS_EVALUATION_UNIFORM_COMPONENTS,GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS,GL_MAX_COMBINED_UNIFORM_BLOCKS,GL_MAX_COMBINED_VERTEX_UNIFORM_COMPONENTS,GL_MAX_COMPUTE_ATOMIC_COUNTERS,GL_MAX_COMPUTE_ATOMIC_COUNTER_BUFFERS,GL_MAX_COMPUTE_IMAGE_UNIFORMS,GL_MAX_COMPUTE_SHADER_STORAGE_BLOCKS,GL_MAX_COMPUTE_TEXTURE_IMAGE_UNITS,GL_MAX_COMPUTE_UNIFORM_BLOCKS,GL_MAX_COMPUTE_UNIFORM_COMPONENTS,GL_MAX_COMPUTE_WORK_GROUP_INVOCATIONS,GL_MAX_COMPUTE_WORK_GROUP_COUNT,GL_MAX_COMPUTE_WORK_GROUP_SIZE,GL_MAX_CUBE_MAP_TEXTURE_SIZE,GL_MAX_DEBUG_GROUP_STACK_DEPTH,GL_MAX_DEBUG_LOGGED_MESSAGES,GL_MAX_DEBUG_MESSAGE_LENGTH,GL_MAX_DRAW_BUFFERS,GL_MAX_ELEMENT_INDEX,GL_MAX_ELEMENTS_INDICES,GL_MAX_ELEMENTS_VERTICES,GL_MAX_FRAGMENT_ATOMIC_COUNTERS,GL_MAX_FRAGMENT_ATOMIC_COUNTER_BUFFERS,GL_MAX_FRAGMENT_IMAGE_UNIFORMS,GL_MAX_FRAGMENT_INPUT_COMPONENTS,GL_MAX_FRAGMENT_INTERPOLATION_OFFSET,GL_MAX_FRAGMENT_SHADER_STORAGE_BLOCKS,GL_MAX_FRAGMENT_UNIFORM_BLOCKS,GL_MAX_FRAGMENT_UNIFORM_COMPONENTS,GL_MAX_FRAGMENT_UNIFORM_VECTORS,GL_MAX_FRAMEBUFFER_HEIGHT,GL_MAX_FRAMEBUFFER_LAYERS,GL_MAX_FRAMEBUFFER_SAMPLES,GL_MAX_FRAMEBUFFER_WIDTH,GL_MAX_GEOMETRY_ATOMIC_COUNTER_BUFFERS,GL_MAX_GEOMETRY_ATOMIC_COUNTERS,GL_MAX_GEOMETRY_IMAGE_UNIFORMS,GL_MAX_GEOMETRY_INPUT_COMPONENTS,GL_MAX_GEOMETRY_OUTPUT_COMPONENTS,GL_MAX_GEOMETRY_OUTPUT_VERTICES,GL_MAX_GEOMETRY_SHADER_STORAGE_BLOCKS,GL_MAX_GEOMETRY_SHADER_INVOCATIONS,GL_MAX_GEOMETRY_TEXTURE_IMAGE_UNITS,GL_MAX_GEOMETRY_TOTAL_OUTPUT_COMPONENTS,GL_MAX_GEOMETRY_UNIFORM_BLOCKS,GL_MAX_GEOMETRY_UNIFORM_COMPONENTS,GL_MAX_INTEGER_SAMPLES,GL_MAX_LABEL_LENGTH,GL_MAX_PROGRAM_TEXEL_OFFSET,GL_MAX_RENDERBUFFER_SIZE,GL_MAX_SAMPLE_MASK_WORDS,GL_MAX_SAMPLES,GL_MAX_SERVER_WAIT_TIMEOUT,GL_MAX_SHADER_STORAGE_BLOCK_SIZE,GL_MAX_SHADER_STORAGE_BUFFER_BINDINGS,GL_MAX_TESS_CONTROL_ATOMIC_COUNTER_BUFFERS,GL_MAX_TESS_CONTROL_ATOMIC_COUNTERS,GL_MAX_TESS_CONTROL_IMAGE_UNIFORMS,GL_MAX_TESS_CONTROL_INPUT_COMPONENTS,GL_MAX_TESS_CONTROL_OUTPUT_COMPONENTS,GL_MAX_TESS_CONTROL_SHADER_STORAGE_BLOCKS,GL_MAX_TESS_CONTROL_TEXTURE_IMAGE_UNITS,GL_MAX_TESS_CONTROL_TOTAL_OUTPUT_COMPONENTS,GL_MAX_TESS_CONTROL_UNIFORM_BLOCKS,GL_MAX_TESS_CONTROL_UNIFORM_COMPONENTS,GL_MAX_TESS_EVALUATION_ATOMIC_COUNTER_BUFFERS,GL_MAX_TESS_EVALUATION_ATOMIC_COUNTERS,GL_MAX_TESS_EVALUATION_IMAGE_UNIFORMS,GL_MAX_TESS_EVALUATION_INPUT_COMPONENTS,GL_MAX_TESS_EVALUATION_OUTPUT_COMPONENTS,GL_MAX_TESS_EVALUATION_SHADER_STORAGE_BLOCKS,GL_MAX_TESS_EVALUATION_TEXTURE_IMAGE_UNITS,GL_MAX_TESS_EVALUATION_UNIFORM_BLOCKS,GL_MAX_TESS_EVALUATION_UNIFORM_COMPONENTS,GL_MAX_TESS_GEN_LEVEL,GL_MAX_TESS_PATCH_COMPONENTS,GL_MAX_TEXTURE_BUFFER_SIZE,GL_MAX_TEXTURE_IMAGE_UNITS,GL_MAX_TEXTURE_LOD_BIAS,GL_MAX_TEXTURE_SIZE,GL_MAX_TRANSFORM_FEEDBACK_INTERLEAVED_COMPONENTS,GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_ATTRIBS,GL_MAX_TRANSFORM_FEEDBACK_SEPARATE_COMPONENTS,GL_MAX_UNIFORM_BLOCK_SIZE,GL_MAX_UNIFORM_BUFFER_BINDINGS,GL_MAX_UNIFORM_LOCATIONS,GL_MAX_VARYING_COMPONENTS,GL_MAX_VARYING_VECTORS,GL_MAX_VERTEX_ATOMIC_COUNTERS,GL_MAX_VERTEX_ATTRIB_BINDINGS,GL_MAX_VERTEX_ATTRIB_RELATIVE_OFFSET,GL_MAX_VERTEX_ATTRIBS,GL_MAX_VERTEX_IMAGE_UNIFORMS,GL_MAX_VERTEX_SHADER_STORAGE_BLOCKS,GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS,GL_MAX_VERTEX_OUTPUT_COMPONENTS,GL_MAX_VERTEX_UNIFORM_BLOCKS,GL_MAX_VERTEX_UNIFORM_COMPONENTS,GL_MAX_VERTEX_UNIFORM_VECTORS,GL_MAX_VIEWPORT_DIMS,GL_MIN_FRAGMENT_INTERPOLATION_OFFSET,GL_MIN_PROGRAM_TEXEL_OFFSET,GL_MIN_SAMPLE_SHADING_VALUE,GL_MINOR_VERSION,GL_MULTISAMPLE_LINE_WIDTH_RANGE,GL_NUM_COMPRESSED_TEXTURE_FORMATS,GL_NUM_EXTENSIONS,GL_NUM_PROGRAM_BINARY_FORMATS,GL_NUM_SHADER_BINARY_FORMATS,GL_PACK_ALIGNMENT,GL_PACK_ROW_LENGTH,GL_PACK_SKIP_PIXELS,GL_PACK_SKIP_ROWS,GL_PATCH_VERTICES,GL_PIXEL_PACK_BUFFER_BINDING,GL_PIXEL_UNPACK_BUFFER_BINDING,GL_POLYGON_OFFSET_FACTOR,GL_POLYGON_OFFSET_FILL,GL_POLYGON_OFFSET_UNITS,GL_PRIMITIVE_BOUNDING_BOX,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_PROGRAM_BINARY_FORMATS,GL_PROGRAM_PIPELINE_BINDING,GL_RASTERIZER_DISCARD,GL_READ_BUFFER,GL_READ_FRAMEBUFFER_BINDING,GL_RED_BITS,GL_RENDERBUFFER_BINDING,GL_RESET_NOTIFICATION_STRATEGY,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_BUFFERS,GL_SAMPLE_COVERAGE,GL_SAMPLE_COVERAGE_INVERT,GL_SAMPLE_COVERAGE_VALUE,GL_SAMPLE_SHADING,GL_SAMPLER_BINDING,GL_SAMPLES,GL_SCISSOR_BOX,GL_SCISSOR_TEST,GL_SHADER_BINARY_FORMATS,GL_SHADER_COMPILER,GL_SHADER_STORAGE_BUFFER_BINDING,GL_SHADER_STORAGE_BUFFER_OFFSET_ALIGNMENT,GL_SHADER_STORAGE_BUFFER_SIZE,GL_SHADER_STORAGE_BUFFER_START,GL_STENCIL_BACK_FAIL,GL_STENCIL_BACK_FUNC,GL_STENCIL_BACK_PASS_DEPTH_FAIL,GL_STENCIL_BACK_PASS_DEPTH_PASS,GL_STENCIL_BACK_REF,GL_STENCIL_BACK_VALUE_MASK,GL_STENCIL_BACK_WRITEMASK,GL_STENCIL_BITS,GL_STENCIL_CLEAR_VALUE,GL_STENCIL_FAIL,GL_STENCIL_FUNC,GL_STENCIL_PASS_DEPTH_FAIL,GL_STENCIL_PASS_DEPTH_PASS,GL_STENCIL_REF,GL_STENCIL_TEST,GL_STENCIL_VALUE_MASK,GL_STENCIL_WRITEMASK,GL_SUBPIXEL_BITS,GL_TEXTURE_BINDING_2D,GL_TEXTURE_BINDING_2D_ARRAY,GL_TEXTURE_BINDING_3D,GL_TEXTURE_BINDING_BUFFER,GL_TEXTURE_BINDING_CUBE_MAP,GL_TEXTURE_BINDING_2D_MULTISAMPLE,GL_TEXTURE_BINDING_2D_MULTISAMPLE_ARRAY,GL_TEXTURE_BINDING_CUBE_MAP_ARRAY,GL_TEXTURE_BUFFER_BINDING,GL_TEXTURE_BUFFER_OFFSET_ALIGNMENT,GL_TRANSFORM_FEEDBACK_BINDING,GL_TRANSFORM_FEEDBACK_ACTIVE,GL_TRANSFORM_FEEDBACK_BUFFER_BINDING,GL_TRANSFORM_FEEDBACK_PAUSED,GL_TRANSFORM_FEEDBACK_BUFFER_SIZE,GL_TRANSFORM_FEEDBACK_BUFFER_START,GL_UNIFORM_BUFFER_BINDING,GL_UNIFORM_BUFFER_OFFSET_ALIGNMENT,GL_UNIFORM_BUFFER_SIZE,GL_UNIFORM_BUFFER_START,GL_UNPACK_ALIGNMENT,GL_UNPACK_IMAGE_HEIGHT,GL_UNPACK_ROW_LENGTH,GL_UNPACK_SKIP_IMAGES,GL_UNPACK_SKIP_PIXELS,GL_UNPACK_SKIP_ROWS,GL_VERTEX_ARRAY_BINDING,GL_VERTEX_BINDING_DIVISOR,GL_VERTEX_BINDING_OFFSET,GL_VERTEX_BINDING_STRIDE,GL_VIEWPORT,GL_FRAMEBUFFER_BINDING"},
		{"glGetBooleanv,glGetFloatv,glGetIntegerv,glGetFixedv", "pname", "1.1", "GL_ACTIVE_TEXTURE,GL_ALIASED_POINT_SIZE_RANGE,GL_ALIASED_LINE_WIDTH_RANGE,GL_ALPHA_BITS,GL_ALPHA_TEST,GL_ALPHA_TEST_FUNC,GL_ALPHA_TEST_REF,GL_ARRAY_BUFFER_BINDING,GL_BLEND,GL_BLEND_DST,GL_BLEND_SRC,GL_BLUE_BITS,GL_CLIENT_ACTIVE_TEXTURE,GL_COLOR_ARRAY,GL_COLOR_ARRAY_BUFFER_BINDING,GL_COLOR_ARRAY_SIZE,GL_COLOR_ARRAY_STRIDE,GL_COLOR_ARRAY_TYPE,GL_COLOR_CLEAR_VALUE,GL_COLOR_LOGIC_OP,GL_COLOR_MATERIAL,GL_COLOR_WRITEMASK,GL_COMPRESSED_TEXTURE_FORMATS,GL_CULL_FACE,GL_CULL_FACE_MODE,GL_CURRENT_COLOR,GL_CURRENT_NORMAL,GL_CURRENT_TEXTURE_COORDS,GL_DEPTH_BITS,GL_DEPTH_CLEAR_VALUE,GL_DEPTH_FUNC,GL_DEPTH_RANGE,GL_DEPTH_TEST,GL_DEPTH_WRITEMASK,GL_ELEMENT_ARRAY_BUFFER_BINDING,GL_FOG,GL_FOG_COLOR,GL_FOG_DENSITY,GL_FOG_END,GL_FOG_HINT,GL_FOG_MODE,GL_FOG_START,GL_FRONT_FACE,GL_GREEN_BITS,GL_IMPLEMENTATION_COLOR_READ_FORMAT_OES,GL_IMPLEMENTATION_COLOR_READ_TYPE_OES,GL_LIGHT_MODEL_AMBIENT,GL_LIGHT_MODEL_TWO_SIDE,GL_LIGHTING,GL_LINE_SMOOTH,GL_LINE_SMOOTH_HINT,GL_LINE_WIDTH,GL_LOGIC_OP_MODE,GL_MATRIX_MODE,GL_MAX_CLIP_PLANES,GL_MAX_LIGHTS,GL_MAX_MODELVIEW_STACK_DEPTH,GL_MAX_PROJECTION_STACK_DEPTH,GL_MAX_TEXTURE_SIZE,GL_MAX_TEXTURE_STACK_DEPTH,GL_MAX_TEXTURE_UNITS,GL_MAX_VIEWPORT_DIMS,GL_MODELVIEW_MATRIX,GL_MODELVIEW_STACK_DEPTH,GL_MULTISAMPLE,GL_NORMAL_ARRAY,GL_NORMAL_ARRAY_BUFFER_BINDING,GL_NORMAL_ARRAY_STRIDE,GL_NORMAL_ARRAY_TYPE,GL_NORMALIZE,GL_NUM_COMPRESSED_TEXTURE_FORMATS,GL_PACK_ALIGNMENT,GL_PERSPECTIVE_CORRECTION_HINT,GL_POINT_DISTANCE_ATTENUATION,GL_POINT_FADE_THRESHOLD_SIZE,GL_POINT_SIZE,GL_POINT_SIZE_ARRAY_BUFFER_BINDING_OES,GL_POINT_SIZE_ARRAY_OES,GL_POINT_SIZE_ARRAY_STRIDE_OES,GL_POINT_SIZE_ARRAY_TYPE_OES,GL_POINT_SIZE_MAX,GL_POINT_SIZE_MIN,GL_POINT_SMOOTH,GL_POINT_SMOOTH_HINT,GL_POINT_SPRITE_OES,GL_POLYGON_OFFSET_FACTOR,GL_POLYGON_OFFSET_FILL,GL_POLYGON_OFFSET_UNITS,GL_PROJECTION_MATRIX,GL_PROJECTION_STACK_DEPTH,GL_RED_BITS,GL_RESCALE_NORMAL,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_ALPHA_TO_ONE,GL_SAMPLE_BUFFERS,GL_SAMPLE_COVERAGE,GL_SAMPLE_COVERAGE_INVERT,GL_SAMPLE_COVERAGE_VALUE,GL_SAMPLES,GL_SCISSOR_BOX,GL_SCISSOR_TEST,GL_SHADE_MODEL,GL_SMOOTH_LINE_WIDTH_RANGE,GL_SMOOTH_POINT_SIZE_RANGE,GL_STENCIL_BITS,GL_STENCIL_CLEAR_VALUE,GL_STENCIL_FAIL,GL_STENCIL_FUNC,GL_STENCIL_PASS_DEPTH_FAIL,GL_STENCIL_PASS_DEPTH_PASS,GL_STENCIL_REF,GL_STENCIL_TEST,GL_STENCIL_VALUE_MASK,GL_STENCIL_WRITEMASK,GL_SUBPIXEL_BITS,GL_TEXTURE_2D,GL_TEXTURE_BINDING_2D,GL_TEXTURE_COORD_ARRAY,GL_TEXTURE_COORD_ARRAY_BUFFER_BINDING,GL_TEXTURE_COORD_ARRAY_SIZE,GL_TEXTURE_COORD_ARRAY_STRIDE,GL_TEXTURE_COORD_ARRAY_TYPE,GL_TEXTURE_MATRIX,GL_TEXTURE_STACK_DEPTH,GL_UNPACK_ALIGNMENT,GL_VIEWPORT,GL_VERTEX_ARRAY,GL_VERTEX_ARRAY_BUFFER_BINDING,GL_VERTEX_ARRAY_SIZE,GL_VERTEX_ARRAY_STRIDE,GL_VERTEX_ARRAY_TYPE," + GL_CLIP_PLANEi + "," + GL_LIGHTi},
		{"glGetBufferParameteriv,glGetBufferParameteri64v,glGetBufferPointerv", "target", "3.2", "GL_ARRAY_BUFFER,GL_ATOMIC_COUNTER_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_DISPATCH_INDIRECT_BUFFER,GL_DRAW_INDIRECT_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_SHADER_STORAGE_BUFFER,GL_TEXTURE_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glGetClipPlanef,glGetClipPlanex", "plane", "1.1", GL_CLIP_PLANEi},
		{"glGetError", "result", "1.1", "GL_NO_ERROR,GL_INVALID_ENUM,GL_INVALID_VALUE,GL_INVALID_OPERATION,GL_STACK_OVERFLOW,GL_STACK_UNDERFLOW,GL_OUT_OF_MEMORY"},
		{"glGetError", "result", "2.0,3.0,3.1", "GL_NO_ERROR,GL_INVALID_ENUM,GL_INVALID_VALUE,GL_INVALID_OPERATION,GL_INVALID_FRAMEBUFFER_OPERATION,GL_OUT_OF_MEMORY"},
		{"glGetFramebufferAttachmentParameteriv", "attachment", "3.0,3.1,3.2", "GL_BACK,GL_DEPTH,GL_STENCIL,GL_DEPTH_ATTACHMENT,GL_STENCIL_ATTACHMENT,GL_DEPTH_STENCIL_ATTACHMENT," + GL_COLOR_ATTACHMENTi},
		{"glGetFramebufferAttachmentParameteriv", "pname", "3.0,3.1", "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE,GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE,GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE,GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE,GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE,GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE,GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE,GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE,GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING,GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME,GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL,GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE,GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER"},
		{"glGetFramebufferAttachmentParameteriv", "pname", "3.2", "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE,GL_FRAMEBUFFER_ATTACHMENT_RED_SIZE,GL_FRAMEBUFFER_ATTACHMENT_GREEN_SIZE,GL_FRAMEBUFFER_ATTACHMENT_BLUE_SIZE,GL_FRAMEBUFFER_ATTACHMENT_ALPHA_SIZE,GL_FRAMEBUFFER_ATTACHMENT_DEPTH_SIZE,GL_FRAMEBUFFER_ATTACHMENT_STENCIL_SIZE,GL_FRAMEBUFFER_ATTACHMENT_COMPONENT_TYPE,GL_FRAMEBUFFER_ATTACHMENT_COLOR_ENCODING,GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME,GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL,GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE,GL_FRAMEBUFFER_ATTACHMENT_LAYERED,GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LAYER"},
		{"glGetFramebufferAttachmentParameteriv", "target", "3.0,3.1,3.2", "GL_DRAW_FRAMEBUFFER,GL_READ_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glGetFramebufferParameteriv", "pname", "3.1", "GL_FRAMEBUFFER_DEFAULT_WIDTH,GL_FRAMEBUFFER_DEFAULT_HEIGHT,GL_FRAMEBUFFER_DEFAULT_SAMPLES,GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS"},
		{"glGetFramebufferParameteriv", "pname", "3.2", "GL_FRAMEBUFFER_DEFAULT_WIDTH,GL_FRAMEBUFFER_DEFAULT_HEIGHT,GL_FRAMEBUFFER_DEFAULT_SAMPLES,GL_FRAMEBUFFER_DEFAULT_FIXED_SAMPLE_LOCATIONS,GL_FRAMEBUFFER_DEFAULT_LAYERS"},
		{"glGetFramebufferParameteriv", "target", "3.1,3.2", "GL_READ_FRAMEBUFFER,GL_DRAW_FRAMEBUFFER,GL_FRAMEBUFFER"},
		{"glGetInternalformativ", "internalformat", "3.0,3.1", "GL_R8,GL_R8_SNORM,GL_R16F,GL_R32F,GL_R8UI,GL_R8I,GL_R16UI,GL_R16I,GL_R32UI,GL_R32I,GL_RG8,GL_RG8_SNORM,GL_RG16F,GL_RG32F,GL_RG8UI,GL_RG8I,GL_RG16UI,GL_RG16I,GL_RG32UI,GL_RG32I,GL_RGB8,GL_SRGB8,GL_RGB565,GL_RGB8_SNORM,GL_R11F_G11F_B10F,GL_RGB9_E5,GL_RGB16F,GL_RGB32F,GL_RGB8UI,GL_RGB8I,GL_RGB16UI,GL_RGB16I,GL_RGB32UI,GL_RGB32I,GL_RGBA8,GL_SRGB8_ALPHA8,GL_RGBA8_SNORM,GL_RGB5_A1,GL_RGBA4,GL_RGB10_A2,GL_RGBA16F,GL_RGBA32F,GL_RGBA8UI,GL_RGBA8I,GL_RGB10_A2UI,GL_RGBA16UI,GL_RGBA16I,GL_RGBA32I,GL_RGBA32UI,GL_DEPTH_COMPONENT16,GL_DEPTH_COMPONENT24,GL_DEPTH_COMPONENT32F,GL_DEPTH24_STENCIL8,GL_DEPTH32F_STENCIL8"},
		{"glGetInternalformativ", "pname", "3.0,3.1,3.2", "GL_NUM_SAMPLE_COUNTS,GL_SAMPLES"},
		{"glGetLightfv,glGetLightxv", "light", "1.1", GL_LIGHTi},
		{"glGetMaterialfv,glGetMaterialxv", "face", "1.1", "GL_FRONT,GL_BACK"},
		{"glGetProgramInterfaceiv", "pname", "3.1,3.2", "GL_ACTIVE_RESOURCES,GL_MAX_NAME_LENGTH,GL_MAX_NUM_ACTIVE_VARIABLES"},
		{"glGetProgramInterfaceiv", "programInterface", "3.1,3.2", "GL_UNIFORM,GL_UNIFORM_BLOCK,GL_ATOMIC_COUNTER_BUFFER,GL_PROGRAM_INPUT,GL_PROGRAM_OUTPUT,GL_TRANSFORM_FEEDBACK_VARYING,GL_BUFFER_VARIABLE,GL_SHADER_STORAGE_BLOCK,GL_TRANSFORM_FEEDBACK_BUFFER"},
		{"glGetProgramPipelineiv", "pname", "3.1", "GL_ACTIVE_PROGRAM,GL_VERTEX_SHADER,GL_FRAGMENT_SHADER,GL_COMPUTE_SHADER,GL_INFO_LOG_LENGTH,GL_VALIDATE_STATUS"},
		{"glGetProgramPipelineiv", "pname", "3.2", "GL_ACTIVE_PROGRAM,GL_VERTEX_SHADER,GL_FRAGMENT_SHADER,GL_TESS_CONTROL_SHADER,GL_TESS_EVALUATION_SHADER,GL_GEOMETRY_SHADER,GL_COMPUTE_SHADER,GL_INFO_LOG_LENGTH,GL_VALIDATE_STATUS"},
		{"glGetProgramResourceIndex", "programInterface", "3.1,3.2", "GL_UNIFORM,GL_UNIFORM_BLOCK,GL_PROGRAM_INPUT,GL_PROGRAM_OUTPUT,GL_TRANSFORM_FEEDBACK_VARYING,GL_BUFFER_VARIABLE,GL_SHADER_STORAGE_BLOCK"},
		{"glGetProgramResourceLocation", "programInterface", "3.1,3.2", "GL_UNIFORM,GL_PROGRAM_INPUT,GL_PROGRAM_OUTPUT,GL_TRANSFORM_FEEDBACK_BUFFER"},
		{"glGetProgramResourceName", "programInterface", "3.1,3.2", "GL_UNIFORM,GL_UNIFORM_BLOCK,GL_PROGRAM_INPUT,GL_PROGRAM_OUTPUT,GL_TRANSFORM_FEEDBACK_VARYING,GL_BUFFER_VARIABLE,GL_SHADER_STORAGE_BLOCK"},
		{"glGetProgramResourceiv", "programInterface", "3.1,3.2", "GL_ATOMIC_COUNTER_BUFFER,GL_BUFFER_VARIABLE,GL_PROGRAM_INPUT,GL_PROGRAM_OUTPUT,GL_SHADER_STORAGE_BLOCK,GL_TRANSFORM_FEEDBACK_VARYING,GL_UNIFORM,GL_UNIFORM_BLOCK,GL_TRANSFORM_FEEDBACK_BUFFER"},
		{"glGetProgramResourceiv", "props", "3.1", "GL_NAME_LENGTH,GL_TYPE,GL_ARRAY_SIZE,GL_OFFSET,GL_BLOCK_INDEX,GL_ARRAY_STRIDE,GL_MATRIX_STRIDE,GL_IS_ROW_MAJOR,GL_ATOMIC_COUNTER_BUFFER_INDEX,GL_BUFFER_BINDING,GL_BUFFER_DATA_SIZE,GL_NUM_ACTIVE_VARIABLES,GL_ACTIVE_VARIABLES,GL_REFERENCED_BY_VERTEX_SHADER,GL_REFERENCED_BY_FRAGMENT_SHADER,GL_REFERENCED_BY_COMPUTE_SHADER,GL_TOP_LEVEL_ARRAY_SIZE,GL_TOP_LEVEL_ARRAY_STRIDE,GL_LOCATION,GL_LOCATION_INDEX,GL_LOCATION_COMPONENT"},
		{"glGetRenderbufferParameteriv", "pname", "3.0,3.1,3.2", "GL_RENDERBUFFER_WIDTH,GL_RENDERBUFFER_HEIGHT,GL_RENDERBUFFER_INTERNAL_FORMAT,GL_RENDERBUFFER_RED_SIZE,GL_RENDERBUFFER_GREEN_SIZE,GL_RENDERBUFFER_BLUE_SIZE,GL_RENDERBUFFER_ALPHA_SIZE,GL_RENDERBUFFER_DEPTH_SIZE,GL_RENDERBUFFER_STENCIL_SIZE,GL_RENDERBUFFER_SAMPLES"},
		{"glGetShaderPrecisionFormat", "precisiontype", "3.0,3.1,3.2", "GL_LOW_FLOAT,GL_MEDIUM_FLOAT,GL_HIGH_FLOAT,GL_LOW_INT,GL_MEDIUM_INT,GL_HIGH_INT"},
		{"glGetString", "name", "1.1", "GL_VENDOR,GL_RENDERER,GL_VERSION,GL_EXTENSIONS"},
		{"glGetString", "name", "2.0", "GL_VENDOR,GL_RENDERER,GL_VERSION,GL_SHADING_LANGUAGE_VERSION,GL_EXTENSIONS"},
		{"glGetString", "name", "3.0,3.1,3.2", "GL_EXTENSIONS,GL_RENDERER,GL_SHADING_LANGUAGE_VERSION,GL_VENDOR,GL_VERSION"},
		{"glGetStringi", "name", "3.0,3.1,3.2", "GL_EXTENSIONS"},
		{"glGetSynciv", "pname", "3.0,3.1,3.2", "GL_OBJECT_TYPE,GL_SYNC_STATUS,GL_SYNC_CONDITION,GL_SYNC_FLAGS"},
		{"glGetTexLevelParameterfv,glGetTexLevelParameteriv", "target", "3.1", "GL_TEXTURE_2D,GL_TEXTURE_3D,GL_TEXTURE_2D_ARRAY,GL_TEXTURE_2D_MULTISAMPLE,GL_TEXTURE_CUBE_MAP_POSITIVE_X,GL_TEXTURE_CUBE_MAP_NEGATIVE_X,GL_TEXTURE_CUBE_MAP_POSITIVE_Y,GL_TEXTURE_CUBE_MAP_NEGATIVE_Y,GL_TEXTURE_CUBE_MAP_POSITIVE_Z,GL_TEXTURE_CUBE_MAP_NEGATIVE_Z"},
		{"glGetTexLevelParameterfv,glGetTexLevelParameteriv", "target", "3.2", "GL_TEXTURE_2D,GL_TEXTURE_3D,GL_TEXTURE_2D_ARRAY,GL_TEXTURE_2D_MULTISAMPLE,GL_TEXTURE_2D_MULTISAMPLE_ARRAY,GL_TEXTURE_CUBE_MAP_POSITIVE_X,GL_TEXTURE_CUBE_MAP_NEGATIVE_X,GL_TEXTURE_CUBE_MAP_POSITIVE_Y,GL_TEXTURE_CUBE_MAP_NEGATIVE_Y,GL_TEXTURE_CUBE_MAP_POSITIVE_Z,GL_TEXTURE_CUBE_MAP_NEGATIVE_Z,GL_TEXTURE_CUBE_MAP_ARRAY,GL_TEXTURE_BUFFER"},
		{"glGetTexParameterfv,glGetTexParameteriv,glGetTexParameterxv", "target", "1.1", "GL_TEXTURE_2D"},
		{"glInvalidateFramebuffer,glInvalidateSubFramebuffer", "attachments", "3.0,3.1", "GL_DEPTH_ATTACHMENT,GL_STENCIL_ATTACHMENT,GL_DEPTH_STENCIL_ATTACHMENT," + GL_COLOR_ATTACHMENTi},
		{"glIsEnabled", "cap", "1.1", "GL_ALPHA_TEST,GL_BLEND,GL_COLOR_ARRAY,GL_COLOR_LOGIC_OP,GL_COLOR_MATERIAL,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_FOG,GL_LIGHTING,GL_LINE_SMOOTH,GL_MULTISAMPLE,GL_NORMAL_ARRAY,GL_NORMALIZE,GL_POINT_SIZE_ARRAY_OES,GL_POINT_SMOOTH,GL_POINT_SPRITE_OES,GL_POLYGON_OFFSET_FILL,GL_RESCALE_NORMAL,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_ALPHA_TO_ONE,GL_SAMPLE_COVERAGE,GL_SCISSOR_TEST,GL_STENCIL_TEST,GL_TEXTURE_2D,GL_TEXTURE_COORD_ARRAY,GL_VERTEX_ARRAY," + GL_CLIP_PLANEi + "," + GL_LIGHTi},
		{"glIsEnabled", "cap", "2.0", "GL_BLEND,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_POLYGON_OFFSET_FILL,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_COVERAGE,GL_SCISSOR_TEST,GL_STENCIL_TEST"},
		{"glIsEnabled", "cap", "3.0", "GL_BLEND,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_POLYGON_OFFSET_FILL,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_RASTERIZER_DISCARD,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_COVERAGE,GL_SCISSOR_TEST,GL_STENCIL_TEST"},
		{"glIsEnabled", "cap", "3.1", "GL_BLEND,GL_CULL_FACE,GL_DEPTH_TEST,GL_DITHER,GL_POLYGON_OFFSET_FILL,GL_PRIMITIVE_RESTART_FIXED_INDEX,GL_RASTERIZER_DISCARD,GL_SAMPLE_ALPHA_TO_COVERAGE,GL_SAMPLE_COVERAGE,GL_SCISSOR_TEST,GL_STENCIL_TEST"},
		{"glLightf,glLightfv,glLightx,glLightxv", "light", "1.1", GL_LIGHTi},
		{"glMapBufferRange", "access", "3.0,3.1", "GL_MAP_READ_BIT,GL_MAP_WRITE_BIT,GL_MAP_INVALIDATE_RANGE_BIT,GL_MAP_INVALIDATE_BUFFER_BIT,GL_MAP_FLUSH_EXPLICIT_BIT,GL_MAP_UNSYNCHRONIZED_BIT"},
		{"glMapBufferRange,glUnmapBuffer", "target", "3.0", "GL_ARRAY_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glMapBufferRange,glUnmapBuffer", "target", "3.1", "GL_ARRAY_BUFFER,GL_ATOMIC_COUNTER_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_DISPATCH_INDIRECT_BUFFER,GL_DRAW_INDIRECT_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_SHADER_STORAGE_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glMapBufferRange,glUnmapBuffer", "target", "3.2", "GL_ARRAY_BUFFER,GL_ATOMIC_COUNTER_BUFFER,GL_COPY_READ_BUFFER,GL_COPY_WRITE_BUFFER,GL_DISPATCH_INDIRECT_BUFFER,GL_DRAW_INDIRECT_BUFFER,GL_ELEMENT_ARRAY_BUFFER,GL_PIXEL_PACK_BUFFER,GL_PIXEL_UNPACK_BUFFER,GL_SHADER_STORAGE_BUFFER,GL_TEXTURE_BUFFER,GL_TRANSFORM_FEEDBACK_BUFFER,GL_UNIFORM_BUFFER"},
		{"glMemoryBarrier,glMemoryBarrierByRegion", "barriers", "3.1", "GL_VERTEX_ATTRIB_ARRAY_BARRIER_BIT,GL_ELEMENT_ARRAY_BARRIER_BIT,GL_UNIFORM_BARRIER_BIT,GL_TEXTURE_FETCH_BARRIER_BIT,GL_SHADER_IMAGE_ACCESS_BARRIER_BIT,GL_COMMAND_BARRIER_BIT,GL_PIXEL_BUFFER_BARRIER_BIT,GL_TEXTURE_UPDATE_BARRIER_BIT,GL_BUFFER_UPDATE_BARRIER_BIT,GL_FRAMEBUFFER_BARRIER_BIT,GL_TRANSFORM_FEEDBACK_BARRIER_BIT,GL_ATOMIC_COUNTER_BARRIER_BIT,GL_SHADER_STORAGE_BARRIER_BIT,GL_ALL_BARRIER_BITS"},
		{"glMultiTexCoord4f", "target", "1.1", GL_TEXTUREi},
		{"glMultiTexCoord4x", "texture", "1.1", GL_TEXTUREi},
		{"glObjectLabel,glGetObjectLabel", "identifier", "3.2", "GL_BUFFER,GL_SHADER,GL_PROGRAM,GL_VERTEX_ARRAY,GL_QUERY,GL_PROGRAM_PIPELINE,GL_TRANSFORM_FEEDBACK,GL_SAMPLER,GL_TEXTURE,GL_RENDERBUFFER,GL_FRAMEBUFFER"},
		{"glPixelStorei", "pname", "1.1,2.0", "GL_PACK_ALIGNMENT,GL_UNPACK_ALIGNMENT"},
		{"glPixelStorei", "pname", "3.0,3.1,3.2", "GL_PACK_ROW_LENGTH,GL_PACK_IMAGE_HEIGHT,GL_PACK_SKIP_PIXELS,GL_PACK_SKIP_ROWS,GL_PACK_SKIP_IMAGES,GL_PACK_ALIGNMENT,GL_UNPACK_ROW_LENGTH,GL_UNPACK_IMAGE_HEIGHT,GL_UNPACK_SKIP_PIXELS,GL_UNPACK_SKIP_ROWS,GL_UNPACK_SKIP_IMAGES,GL_UNPACK_ALIGNMENT"},
		{"glProgramBinary", "binaryFormat", "3.0,3.1,3.2", ""}, // implementation dependent
		{"glProgramParameteri", "pname", "3.0", "GL_PROGRAM_BINARY_RETRIEVABLE_HINT"},
		{"glProgramParameteri", "pname", "3.1,3.2", "GL_PROGRAM_BINARY_RETRIEVABLE_HINT,GL_PROGRAM_SEPARABLE"},
		{"glPushDebugGroup", "source", "3.2", "GL_DEBUG_SOURCE_APPLICATION,GL_DEBUG_SOURCE_THIRD_PARTY"},
		{"glReadBuffer", "src", "3.0,3.1,3.2", "GL_BACK,GL_NONE," + GL_COLOR_ATTACHMENTi},
		{"glReadPixels", "format", "1.1", "GL_ALPHA,GL_RGB,GL_RGBA"},
		{"glReadPixels", "type", "1.1", "GL_UNSIGNED_BYTE,GL_UNSIGNED_SHORT_5_6_5,GL_UNSIGNED_SHORT_4_4_4_4,GL_UNSIGNED_SHORT_5_5_5_1"},
		{"glReadPixels,glReadnPixels", "format", "2.0,3.0,3.1,3.2", "GL_RGBA,GL_RGBA_INTEGER"},
		{"glReadPixels,glReadnPixels", "type", "2.0,3.0", "GL_UNSIGNED_BYTE,GL_UNSIGNED_INT,GL_INT,GL_FLOAT"},
		{"glReadPixels,glReadnPixels", "type", "2.0,3.1,3.2", "GL_UNSIGNED_BYTE,GL_UNSIGNED_INT,GL_UNSIGNED_INT_2_10_10_10_REV,GL_INT,GL_FLOAT"},
		{"glRenderbufferStorage,glRenderbufferStorageMultisample", "internalformat", "3.0,3.1,3.2", "GL_R8,GL_R8UI,GL_R8I,GL_R16UI,GL_R16I,GL_R32UI,GL_R32I,GL_RG8,GL_RG8UI,GL_RG8I,GL_RG16UI,GL_RG16I,GL_RG32UI,GL_RG32I,GL_RGB8,GL_RGB565,GL_RGBA8,GL_SRGB8_ALPHA8,GL_RGB5_A1,GL_RGBA4,GL_RGB10_A2,GL_RGBA8UI,GL_RGBA8I,GL_RGB10_A2UI,GL_RGBA16UI,GL_RGBA16I,GL_RGBA32I,GL_RGBA32UI,GL_DEPTH_COMPONENT16,GL_DEPTH_COMPONENT24,GL_DEPTH_COMPONENT32F,GL_DEPTH24_STENCIL8,GL_DEPTH32F_STENCIL8,GL_STENCIL_INDEX8"},
		{"glRenderbufferStorage,glRenderbufferStorageMultisample", "target", "3.0,3.1,3.2", "GL_RENDERBUFFER"},

		{"glSamplerParameterf,glSamplerParameteri", "pname", "3.2", "GL_TEXTURE_WRAP_S,GL_TEXTURE_WRAP_T,GL_TEXTURE_WRAP_R,GL_TEXTURE_MIN_FILTER,GL_TEXTURE_MAG_FILTER,GL_TEXTURE_MIN_LOD,GL_TEXTURE_MAX_LOD,GL_TEXTURE_COMPARE_MODE,GL_TEXTURE_COMPARE_FUNC"},
		{"glSamplerParameterfv,glSamplerParameteriv,glSamplerParameterIiv,glSamplerParameterIuiv", "pname", "3.2", "GL_TEXTURE_WRAP_S,GL_TEXTURE_WRAP_T,GL_TEXTURE_WRAP_R,GL_TEXTURE_MIN_FILTER,GL_TEXTURE_MAG_FILTER,GL_TEXTURE_MIN_LOD,GL_TEXTURE_MAX_LOD,GL_TEXTURE_COMPARE_MODE,GL_TEXTURE_COMPARE_FUNC,GL_TEXTURE_BORDER_COLOR"},

		{"glShaderBinary", "binaryformat", "2.0,3.0,3.1,3.2", ""}, // OpenGL ES defines no specific binary formats
		{"glStencilOp", "fail,zfail,zpass", "1.1", "GL_KEEP,GL_ZERO,GL_REPLACE,GL_INCR,GL_DECR,GL_INVERT"},
		{"glStencilOp,glStencilOpSeparate", "fail,zfail,zpass,sfail,dpfail,dppass", "2.0,3.0,3.1,3.2", "GL_KEEP,GL_ZERO,GL_REPLACE,GL_INCR,GL_INCR_WRAP,GL_DECR,GL_DECR_WRAP,GL_INVERT"},
		{"glTexBuffer,glTexBufferRange", "internalformat", "3.2", "GL_R8,GL_R16,GL_R16F,GL_R32F,GL_R8I,GL_R16I,GL_R32I,GL_R8UI,GL_R16UI,GL_R32UI,GL_RG8,GL_RG16,GL_RG16F,GL_RG32F,GL_RG8I,GL_RG16I,GL_RG32I,GL_RG8UI,GL_RG16UI,GL_RG32UI,GL_RGB32F,GL_RGB32I,GL_RGB32UI,GL_RGBA8,GL_RGBA16,GL_RGBA16F,GL_RGBA32F,GL_RGBA8I,GL_RGBA16I,GL_RGBA32I,GL_RGBA8UI,GL_RGBA16UI,GL_RGBA32UI"},
		{"glTexCoordPointer,glVertexPointer,glNormalPointer", "type", "1.1", "GL_BYTE,GL_SHORT,GL_FIXED,GL_FLOAT"}, // TODO: common profile?
		{"glTexParameterf,glTexParameterfv,glTexParameteri,glTexParameteriv", "target", "1.1", "GL_TEXTURE_2D"},
		{"glTexParameterf,glTexParameterfv,glTexParameteri,glTexParameteriv", "target", "2.0", "GL_TEXTURE_2D,GL_TEXTURE_CUBE_MAP"},
		{"glTexParameterf,glTexParameterfv,glTexParameteri,glTexParameteriv", "target", "3.0", "GL_TEXTURE_2D,GL_TEXTURE_3D,GL_TEXTURE_2D_ARRAY,GL_TEXTURE_CUBE_MAP"},
		{"glTexParameterf,glTexParameterfv,glTexParameteri,glTexParameteriv", "target", "3.1", "GL_TEXTURE_2D,GL_TEXTURE_3D,GL_TEXTURE_2D_ARRAY,GL_TEXTURE_2D_MULTISAMPLE,GL_TEXTURE_CUBE_MAP"},
		{"glTexParameterx,glTexParameterxv", "target", "1.1", "GL_TEXTURE_2D"},
		{"glTexParameterf,glTexParameteri,glTexParameterfv,glTexParameteriv,glTexParameterIiv,glTexParameterIuiv", "target", "3.2", "GL_TEXTURE_2D,GL_TEXTURE_3D,GL_TEXTURE_2D_ARRAY,GL_TEXTURE_2D_MULTISAMPLE,GL_TEXTURE_2D_MULTISAMPLE_ARRAY,GL_TEXTURE_CUBE_MAP,GL_TEXTURE_CUBE_MAP_ARRAY"},
		{"glTexParameterf,glTexParameteri", "pname", "3.2", "GL_DEPTH_STENCIL_TEXTURE_MODE,GL_TEXTURE_BASE_LEVEL,GL_TEXTURE_COMPARE_FUNC,GL_TEXTURE_COMPARE_MODE,GL_TEXTURE_MIN_FILTER,GL_TEXTURE_MAG_FILTER,GL_TEXTURE_MIN_LOD,GL_TEXTURE_MAX_LOD,GL_TEXTURE_MAX_LEVEL,GL_TEXTURE_SWIZZLE_R,GL_TEXTURE_SWIZZLE_G,GL_TEXTURE_SWIZZLE_B,GL_TEXTURE_SWIZZLE_A,GL_TEXTURE_WRAP_S,GL_TEXTURE_WRAP_T,GL_TEXTURE_WRAP_R"},
		{"glTexParameterfv,glTexParameteriv,glTexParameterIiv,glTexParameterIuiv", "pname", "3.2", "GL_DEPTH_STENCIL_TEXTURE_MODE,GL_TEXTURE_BASE_LEVEL,GL_TEXTURE_COMPARE_FUNC,GL_TEXTURE_COMPARE_MODE,GL_TEXTURE_MIN_FILTER,GL_TEXTURE_MAG_FILTER,GL_TEXTURE_MIN_LOD,GL_TEXTURE_MAX_LOD,GL_TEXTURE_MAX_LEVEL,GL_TEXTURE_SWIZZLE_R,GL_TEXTURE_SWIZZLE_G,GL_TEXTURE_SWIZZLE_B,GL_TEXTURE_SWIZZLE_A,GL_TEXTURE_WRAP_S,GL_TEXTURE_WRAP_T,GL_TEXTURE_WRAP_R,GL_TEXTURE_BORDER_COLOR"},
		{"glTexStorage2D,glTexStorage3D", "internalformat", "3.0,3.1", "GL_R8,GL_R8_SNORM,GL_R16F,GL_R32F,GL_R8UI,GL_R8I,GL_R16UI,GL_R16I,GL_R32UI,GL_R32I,GL_RG8,GL_RG8_SNORM,GL_RG16F,GL_RG32F,GL_RG8UI,GL_RG8I,GL_RG16UI,GL_RG16I,GL_RG32UI,GL_RG32I,GL_RGB8,GL_SRGB8,GL_RGB565,GL_RGB8_SNORM,GL_R11F_G11F_B10F,GL_RGB9_E5,GL_RGB16F,GL_RGB32F,GL_RGB8UI,GL_RGB8I,GL_RGB16UI,GL_RGB16I,GL_RGB32UI,GL_RGB32I,GL_RGBA8,GL_SRGB8_ALPHA8,GL_RGBA8_SNORM,GL_RGB5_A1,GL_RGBA4,GL_RGB10_A2,GL_RGBA16F,GL_RGBA32F,GL_RGBA8UI,GL_RGBA8I,GL_RGB10_A2UI,GL_RGBA16UI,GL_RGBA16I,GL_RGBA32I,GL_RGBA32UI,GL_DEPTH_COMPONENT16,GL_DEPTH_COMPONENT24,GL_DEPTH_COMPONENT32F,GL_DEPTH24_STENCIL8,GL_DEPTH32F_STENCIL8,GL_COMPRESSED_R11_EAC,GL_COMPRESSED_SIGNED_R11_EAC,GL_COMPRESSED_RG11_EAC,GL_COMPRESSED_SIGNED_RG11_EAC,GL_COMPRESSED_RGB8_ETC2,GL_COMPRESSED_SRGB8_ETC2,GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_RGBA8_ETC2_EAC,GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC"},
		{"glTexStorage2DMultisample", "internalformat", "3.1", "GL_R8,GL_R8_SNORM,GL_R16F,GL_R32F,GL_R8UI,GL_R8I,GL_R16UI,GL_R16I,GL_R32UI,GL_R32I,GL_RG8,GL_RG8_SNORM,GL_RG16F,GL_RG32F,GL_RG8UI,GL_RG8I,GL_RG16UI,GL_RG16I,GL_RG32UI,GL_RG32I,GL_RGB8,GL_SRGB8,GL_RGB565,GL_RGB8_SNORM,GL_R11F_G11F_B10F,GL_RGB9_E5,GL_RGB16F,GL_RGB32F,GL_RGB8UI,GL_RGB8I,GL_RGB16UI,GL_RGB16I,GL_RGB32UI,GL_RGB32I,GL_RGBA8,GL_SRGB8_ALPHA8,GL_RGBA8_SNORM,GL_RGB5_A1,GL_RGBA4,GL_RGB10_A2,GL_RGBA16F,GL_RGBA32F,GL_RGBA8UI,GL_RGBA8I,GL_RGB10_A2UI,GL_RGBA16UI,GL_RGBA16I,GL_RGBA32I,GL_RGBA32UI,GL_DEPTH_COMPONENT16,GL_DEPTH_COMPONENT24,GL_DEPTH_COMPONENT32F,GL_DEPTH24_STENCIL8,GL_DEPTH32F_STENCIL8"},
		{"glTexStorage2D,glTexStorage3D", "internalformat", "3.2", "GL_R8,GL_R8_SNORM,GL_R16F,GL_R32F,GL_R8UI,GL_R8I,GL_R16UI,GL_R16I,GL_R32UI,GL_R32I,GL_RG8,GL_RG8_SNORM,GL_RG16F,GL_RG32F,GL_RG8UI,GL_RG8I,GL_RG16UI,GL_RG16I,GL_RG32UI,GL_RG32I,GL_RGB8,GL_SRGB8,GL_RGB565,GL_RGB8_SNORM,GL_R11F_G11F_B10F,GL_RGB9_E5,GL_RGB16F,GL_RGB32F,GL_RGB8UI,GL_RGB8I,GL_RGB16UI,GL_RGB16I,GL_RGB32UI,GL_RGB32I,GL_RGBA8,GL_SRGB8_ALPHA8,GL_RGBA8_SNORM,GL_RGB5_A1,GL_RGBA4,GL_RGB10_A2,GL_RGBA16F,GL_RGBA32F,GL_RGBA8UI,GL_RGBA8I,GL_RGB10_A2UI,GL_RGBA16UI,GL_RGBA16I,GL_RGBA32I,GL_RGBA32UI,GL_DEPTH_COMPONENT16,GL_DEPTH_COMPONENT24,GL_DEPTH_COMPONENT32F,GL_DEPTH24_STENCIL8,GL_DEPTH32F_STENCIL8,GL_STENCIL_INDEX8,GL_COMPRESSED_R11_EAC,GL_COMPRESSED_SIGNED_R11_EAC,GL_COMPRESSED_RG11_EAC,GL_COMPRESSED_SIGNED_RG11_EAC,GL_COMPRESSED_RGB8_ETC2,GL_COMPRESSED_SRGB8_ETC2,GL_COMPRESSED_RGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_SRGB8_PUNCHTHROUGH_ALPHA1_ETC2,GL_COMPRESSED_RGBA8_ETC2_EAC,GL_COMPRESSED_SRGB8_ALPHA8_ETC2_EAC,GL_COMPRESSED_RGBA_ASTC_4x4,GL_COMPRESSED_RGBA_ASTC_5x4,GL_COMPRESSED_RGBA_ASTC_5x5,GL_COMPRESSED_RGBA_ASTC_6x5,GL_COMPRESSED_RGBA_ASTC_6x6,GL_COMPRESSED_RGBA_ASTC_8x5,GL_COMPRESSED_RGBA_ASTC_8x6,GL_COMPRESSED_RGBA_ASTC_8x8,GL_COMPRESSED_RGBA_ASTC_10x5,GL_COMPRESSED_RGBA_ASTC_10x6,GL_COMPRESSED_RGBA_ASTC_10x8,GL_COMPRESSED_RGBA_ASTC_10x10,GL_COMPRESSED_RGBA_ASTC_12x10,GL_COMPRESSED_RGBA_ASTC_12x12,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10,GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12"},
		{"glTexStorage2DMultisample,glTexStorage3DMultisample,glGetInternalformativ", "internalformat", "3.2", "GL_R8,GL_R8_SNORM,GL_R16F,GL_R32F,GL_R8UI,GL_R8I,GL_R16UI,GL_R16I,GL_R32UI,GL_R32I,GL_RG8,GL_RG8_SNORM,GL_RG16F,GL_RG32F,GL_RG8UI,GL_RG8I,GL_RG16UI,GL_RG16I,GL_RG32UI,GL_RG32I,GL_RGB8,GL_SRGB8,GL_RGB565,GL_RGB8_SNORM,GL_R11F_G11F_B10F,GL_RGB9_E5,GL_RGB16F,GL_RGB32F,GL_RGB8UI,GL_RGB8I,GL_RGB16UI,GL_RGB16I,GL_RGB32UI,GL_RGB32I,GL_RGBA8,GL_SRGB8_ALPHA8,GL_RGBA8_SNORM,GL_RGB5_A1,GL_RGBA4,GL_RGB10_A2,GL_RGBA16F,GL_RGBA32F,GL_RGBA8UI,GL_RGBA8I,GL_RGB10_A2UI,GL_RGBA16UI,GL_RGBA16I,GL_RGBA32I,GL_RGBA32UI,GL_DEPTH_COMPONENT16,GL_DEPTH_COMPONENT24,GL_DEPTH_COMPONENT32F,GL_DEPTH24_STENCIL8,GL_DEPTH32F_STENCIL8,GL_STENCIL_INDEX8"},
		{"glUseProgramStages", "stages", "3.1", "GL_VERTEX_SHADER_BIT,GL_FRAGMENT_SHADER_BIT,GL_COMPUTE_SHADER_BIT,GL_ALL_SHADER_BITS"},
		{"glVertexAttribFormat,glVertexAttribIFormat", "type", "3.1,3.2", "GL_BYTE,GL_SHORT,GL_INT,GL_FIXED,GL_FLOAT,GL_HALF_FLOAT,GL_UNSIGNED_BYTE,GL_UNSIGNED_SHORT,GL_UNSIGNED_INT,GL_INT_2_10_10_10_REV,GL_UNSIGNED_INT_2_10_10_10_REV"},
		{"glVertexAttribIPointer", "type", "3.0,3.1,3.2", "GL_BYTE,GL_UNSIGNED_BYTE,GL_SHORT,GL_UNSIGNED_SHORT,GL_INT,GL_UNSIGNED_INT"},
		{"glVertexAttribPointer", "type", "3.0,3.1,3.2", "GL_BYTE,GL_UNSIGNED_BYTE,GL_SHORT,GL_UNSIGNED_SHORT,GL_INT,GL_UNSIGNED_INT,GL_HALF_FLOAT,GL_FLOAT,GL_FIXED,GL_INT_2_10_10_10_REV,GL_UNSIGNED_INT_2_10_10_10_REV"},
	}
	for _, row := range enumMetadata {
		if listContains(row.cmdName, cmdName) &&
			listContains(row.paramName, paramName) &&
			listContains(row.apiVersion, string(apiVersion)) {
			if row.acceptedValues != "" {
				return strings.Split(row.acceptedValues, ","), true
			} else {
				return []string{}, true
			}
		}
	}
	return nil, false
}
