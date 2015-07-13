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
	"path/filepath"
	"strings"
	"unicode"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

type JavaStruct struct {
	*Struct
}

// Name returns the Java name to give the type.
func (s *JavaStruct) Name() string {
	name := s.Tags.Get("java")
	if name == "" {
		name = s.Struct.Name
	}
	return name
}

type JavaClass struct {
	*Module
	Struct       JavaStruct
	JavaPackage  string
	Copyright    string
	MemberPrefix string
}

func Java(m *Module, info copyright.Info, gen Generator, path string) error {
	class := &JavaClass{
		Module:      m,
		JavaPackage: m.Directives["java.package"],
		Copyright:   strings.TrimSpace(copyright.Build("generated_aosp_java", info)),
	}
	class.MemberPrefix, _ = m.Directives["java.member_prefix"]
	source, _ := m.Directives["java.source"]
	indent, _ := m.Directives["java.indent"]
	reflow := indentor(indent)
	pkgPath := strings.Replace(class.JavaPackage, ".", "/", -1)
	for _, s := range m.Structs {
		class.Struct.Struct = s
		out := filepath.Join(path, source, pkgPath, class.Struct.Name()+".java")
		if err := gen("Java.File", class, out, reflow); err != nil {
			return err
		}
	}
	return nil
}

func (class *JavaClass) FieldName(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return class.MemberPrefix + string(unicode.ToUpper(r)) + s[n:]
}
