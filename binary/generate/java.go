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
	"unicode"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

var (
	javaTypeMap = map[string]string{
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

func (f *functions) JavaFieldName(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return memberPrefix + string(unicode.ToUpper(r)) + s[n:]
}

func (f *functions) JavaStorage(t schema.Type) string {
	switch t := t.(type) {
	case *schema.Primitive:
		name := t.Name
		if result, ok := javaTypeMap[name]; ok {
			return result
		}
		return name
	case *schema.Struct:
		return t.Name
	case *schema.Interface:
		return t.Name
	case *schema.Pointer:
		return f.JavaStorage(t.Type)
	case *schema.Array:
		panic(fmt.Errorf("Array types not handled"))
	case *schema.Slice:
		panic(fmt.Errorf("Slice types not handled"))
	case *schema.Stream:
		panic(fmt.Errorf("Stream types not handled"))
	case *schema.Map:
		panic(fmt.Errorf("Map types not handled"))
	default:
		panic(fmt.Errorf("Unknown value type %T", t))
	}
}

func (f *functions) JavaID(name string) string {
	return classPrefix + name
}

func (f *functions) JavaClass(name string) string {
	if strings.HasPrefix(name, "call") {
		return "Commands." + classPrefix + name[4:] + ".Call"
	}
	if strings.HasPrefix(name, "result") {
		return "Commands." + classPrefix + name[6:] + ".Result"
	}
	return classPrefix + name
}

// JavaFile generates the all the java code for a file with a set of structs.
func (g *Generator) JavaFile(file *File) ([]byte, error) {
	g.f.prefix = "Java."
	f := *file
	b := &bytes.Buffer{}
	if err := g.f.execute(g.f.prefix+"File", b, &f); err != nil {
		return nil, err
	}
	s := b.String()
	s = strings.Replace(s, indent, f.Indent, -1)
	s = strings.Replace(s, memberPrefix, f.MemberPrefix, -1)
	s = strings.Replace(s, classPrefix, f.ClassPrefix, -1)
	return []byte(s), nil
}
