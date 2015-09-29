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
	"flag"
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/parse"
)

var (
	apiPath  = flag.String("api", "", "Filename of the api file to verify (required)")
	cacheDir = flag.String("cache", "", "Directory for caching downloaded files (required)")
	apiRoot  *semantic.API
)

func main() {
	flag.Parse()
	if *apiPath == "" || *cacheDir == "" {
		flag.PrintDefaults()
		return
	}
	mappings := resolver.ASTToSemantic{}
	api, errs := api.Resolve(*apiPath, mappings)
	if len(errs) > 0 {
		for _, err := range errs {
			fmt.Printf("%v", err.Message)
		}
		return
	}
	apiRoot = api
	reg := DownloadRegistry()
	VerifyApi(reg, GLES2API)
}

func VerifyApi(reg *Registry, api KhronosAPI) {
	VerifyEnum(reg, api, false)
	VerifyEnum(reg, api, true)
	expected := make(map[string]struct{})
	for _, cmd := range reg.Command {
		if reg.GetVersions(api, cmd.Name()) != nil || reg.GetExtensions(api, cmd.Name()) != nil {
			expected[cmd.Name()] = struct{}{}
			VerifyCommand(reg, cmd, api)
		}
	}
	seen := make(map[string]struct{})
	for _, f := range apiRoot.Functions {
		if strings.HasPrefix(f.Name(), "gl") && !strings.HasPrefix(f.Name(), "glX") {
			seen[f.Name()] = struct{}{}
		}
	}
	CompareSets(expected, seen, "")
}

func VerifyEnum(r *Registry, api KhronosAPI, bitfields bool) {
	name := "GLenum"
	if bitfields {
		name = "GLbitfield"
	}
	expected := make(map[string]struct{})
	for _, enums := range r.Enums {
		if (enums.Type == "bitmask") == bitfields && enums.Namespace == "GL" {
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
					expected[fmt.Sprintf("%s = 0x%08X", enum.Name, value)] = struct{}{}
				}
			}
		}
	}
	seen := make(map[string]struct{})
	for _, e := range apiRoot.Enums {
		if e.Name() == name {
			for _, m := range e.Entries {
				seen[fmt.Sprintf("%s = 0x%08X", m.Name(), m.Value)] = struct{}{}
			}
		}
	}
	CompareSets(expected, seen, name+": ")
}

func CompareSets(expected, seen map[string]struct{}, msg_prefix string) {
	for k, _ := range expected {
		if _, found := seen[k]; !found {
			fmt.Printf("%sMissing %s\n", msg_prefix, k)
		}
	}
	for k, _ := range seen {
		if _, found := expected[k]; !found {
			fmt.Printf("%sUnexpected %s\n", msg_prefix, k)
		}
	}
	return
}

var re_const_ptr_pre = regexp.MustCompile(`^const (\w+) \*$`)
var re_const_ptr_post = regexp.MustCompile(`^(.+)\bconst\*$`)

func VerifyType(cmd string, paramIndex int, expected string, seen semantic.Type) bool {
	expected = strings.TrimSpace(expected)
	name := seen.(semantic.NamedNode).Name()
	switch s := seen.(type) {
	case *semantic.Pointer:
		if s.Const {
			if match := re_const_ptr_pre.FindStringSubmatch(expected); match != nil {
				return VerifyType(cmd, paramIndex, match[1], s.To)
			}
			if match := re_const_ptr_post.FindStringSubmatch(expected); match != nil {
				return VerifyType(cmd, paramIndex, match[1], s.To)
			}
		} else {
			if strings.HasSuffix(expected, "*") {
				return VerifyType(cmd, paramIndex, strings.TrimSuffix(expected, "*"), s.To)
			}
		}
	case *semantic.Pseudonym:
		if s.Name() == expected {
			return true
		} else if expected == "GLDEBUGPROCKHR" && s.Name() == "GLDEBUGPROC" {
			return true
		} else {
			return VerifyType(cmd, paramIndex, expected, s.To)
		}
	case *semantic.Enum:
		if s.Name() == expected {
			return true
		}
	case *semantic.Builtin:
		if s.Name() == expected {
			return true
		} else if expected == "const GLchar *" && s.Name() == "string" {
			return true
		}
	}
	fmt.Printf("%s: Param %v: Expected type %s but seen %s (%T)\n", cmd, paramIndex, expected, name, seen)
	return false
}

func VerifyCommand(reg *Registry, cmd *Command, api KhronosAPI) {
	cmdName := cmd.Name()
	versions := reg.GetVersions(api, cmdName)
	extensions := reg.GetExtensions(api, cmdName)

	// Find API function.
	var apiCmd *semantic.Function
	for _, apiFunction := range apiRoot.Functions {
		if apiFunction.Name() == cmdName {
			apiCmd = apiFunction
		}
	}
	if apiCmd == nil {
		return
	}

	// Check documentation strings.
	seen := make(map[string]struct{})
	for _, a := range apiCmd.Annotations {
		if a.Name() == "Doc" || a.Name() == "DrawCall" {
			seen[getSource(a.AST.Node())] = struct{}{}
		}
	}
	expected := make(map[string]struct{})
	for _, version := range versions {
		url, _ := GetCoreManpage(version, cmdName)
		expected[fmt.Sprintf(`@Doc("%s","OpenGL ES %v")`, url, version)] = struct{}{}
	}
	for _, extension := range extensions {
		url, _ := GetExtensionManpage(extension)
		expected[fmt.Sprintf(`@Doc("%s","%v")`, url, extension)] = struct{}{}
	}
	if strings.HasPrefix(cmdName, "glDraw") && !strings.HasPrefix(cmdName, "glDrawBuffers") {
		expected["@DrawCall"] = struct{}{}
	}
	CompareSets(expected, seen, fmt.Sprintf("%s: ", cmdName))

	// Check parameter types.
	if len(cmd.Param) != len(apiCmd.CallParameters()) {
		fmt.Printf("%s: Expected %v parameters but seen %v\n", cmdName, len(cmd.Param), len(apiCmd.CallParameters()))
	} else {
		for i, p := range cmd.Param {
			VerifyType(cmdName, i, p.Type(), apiCmd.FullParameters[i].Type)
		}
	}

	// Check version.
	stmts := apiCmd.Block.AST.Statements
	if len(stmts) == 0 {
		fmt.Printf("%s: Empty method body\n", cmdName)
	} else {
		expected := make(map[string]struct{})
		if versions != nil {
			version := strings.Replace(string(versions[0]), ".", ", ", -1)
			expected[fmt.Sprintf("minRequiredVersion(%v)", version)] = struct{}{}
		}
		if extensions != nil {
			expected[fmt.Sprintf("requiresExtension(%s)", extensions[0])] = struct{}{}
			// TODO: Handle multiple extensions
		}
		seen := map[string]struct{}{getSource(stmts[0].Node()): {}}
		CompareSets(expected, seen, fmt.Sprintf("%s: ", cmdName))
	}
}

func getSource(n parse.Node) string {
	return string(n.Token().Source.Runes[n.Token().Start:n.Token().End])
}
