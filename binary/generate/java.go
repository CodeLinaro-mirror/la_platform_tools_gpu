// Copyright (C) 2014 The Android Open Source Project
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

package generate

import (
	"bytes"
	"fmt"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

const (
	indent       = "»"
	classPrefix  = "⊹"
	memberPrefix = "∍"
)

var (
	javaTemplates = template.Must(template.New("java.tmpl").Funcs(javaFuncs).Parse(java_tmpl))
	javaFuncs     = template.FuncMap{
		"encode": func(name string, t *Type) string {
			return kindDispatch(javaEncodeMap, name, t)
		},
		"decode": func(name string, t *Type) string {
			return kindDispatch(javaDecodeMap, name, t)
		},
		"lower": strings.ToLower,
		"fieldname": func(s string) string {
			r, n := utf8.DecodeRuneInString(s)
			return fmt.Sprintf(memberPrefix+"%s%s", string(unicode.ToUpper(r)), s[n:])
		},
		"toS8": func(val byte) string { return fmt.Sprint(int8(val)) },
		"storage": func(t *Type) string {
			name := t.Name
			if t.Kind == Pointer {
				name = t.SubType.Name
			}
			if result, ok := javaTypeMap[name]; ok {
				return result
			}
			return name
		},
		"id": func(name string) string {
			return fmt.Sprintf(classPrefix+"%s", name)
		},
		"class": func(name string) string {
			if strings.HasPrefix(name, "call") {
				return fmt.Sprintf("Commands."+classPrefix+"%s.Call", name[4:])
			}
			if strings.HasPrefix(name, "result") {
				return fmt.Sprintf("Commands."+classPrefix+"%s.Result", name[6:])
			}
			return fmt.Sprintf(classPrefix+"%s", name)
		},
	}
	javaEncodeMap kindToTemplate
	javaDecodeMap kindToTemplate
	javaFile      *template.Template
	javaTypeMap   = map[string]string{
		"int8":    "byte",
		"uint8":   "byte",
		"int16":   "short",
		"uint16":  "short",
		"int32":   "int",
		"uint32":  "int",
		"int64":   "long",
		"uint64":  "long",
		"float32": "float",
		"float64": "double",
	}
)

func init() {
	javaFile = getTemplate(javaTemplates, "File")
	javaEncodeMap = getTemplateMap(javaTemplates, "Encode")
	javaDecodeMap = getTemplateMap(javaTemplates, "Decode")
}

// JavaFile generates the all the java code for a file with a set of structs.
func JavaFile(file *File) ([]byte, error) {
	f := *file
	if f.MemberPrefix == "" {
		f.MemberPrefix = "m"
	}
	if f.Indent == "" {
		f.Indent = "    "
	}
	b := &bytes.Buffer{}
	err := javaFile.Execute(b, &f)
	s := b.String()
	s = strings.Replace(s, indent, f.Indent, -1)
	s = strings.Replace(s, memberPrefix, f.MemberPrefix, -1)
	s = strings.Replace(s, classPrefix, f.ClassPrefix, -1)
	return []byte(s), err
}
