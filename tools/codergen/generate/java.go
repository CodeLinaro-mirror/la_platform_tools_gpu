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
	"fmt"
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
func (s JavaStruct) Name() string {
	name := s.Tags.Get("java")
	if name == "" {
		name = strings.Title(s.Struct.Name)
	}
	return name
}

type JavaSettings struct {
	*Module
	JavaPackage  string
	Copyright    string
	MemberPrefix string
}

type JavaClass struct {
	JavaSettings
	Struct JavaStruct
}

type JavaService struct {
	JavaSettings
	Service *Service
}

func Java(m *Module, info copyright.Info, gen Generator, path string) error {
	settings := JavaSettings{
		Module:      m,
		JavaPackage: m.Directives["java.package"],
		Copyright:   strings.TrimSpace(copyright.Build("generated_aosp_java", info)),
	}
	settings.MemberPrefix, _ = m.Directives["java.member_prefix"]
	source, _ := m.Directives["java.source"]
	indent, _ := m.Directives["java.indent"]
	reflow := indentor(indent)
	pkgPath := strings.Replace(settings.JavaPackage, ".", "/", -1)
	for _, s := range m.Structs {
		class := JavaClass{JavaSettings: settings, Struct: JavaStruct{s}}
		out := filepath.Join(path, source, pkgPath, class.Struct.Name()+".java")
		if err := gen("Java.File", class, out, reflow); err != nil {
			return err
		}
	}
	for _, s := range m.Services {
		service := JavaService{JavaSettings: settings, Service: s}
		for _, e := range []string{"Client", "ClientImpl"} {
			out := filepath.Join(path, source, pkgPath, service.Service.Name+e+".java")
			if err := gen("Java."+e, service, out, reflow); err != nil {
				return err
			}
		}
	}
	return nil
}

func (settings JavaSettings) FieldName(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return settings.MemberPrefix + string(unicode.ToUpper(r)) + s[n:]
}

// Name returns the Java name to give the type.
func (settings JavaSettings) ClassName(s string) string {
	pkg, name := "", s
	if i := strings.LastIndexAny(s, "."); i >= 0 {
		pkg = s[:i]
		name = s[i+1:]
	}
	name = strings.Title(name)
	if pkg != "" {
		if m := settings.FindImport(pkg); m != nil {
			name = fmt.Sprintf("%v.%s", m.Directive("java.package", "NotJava."+m.Name), name)
		} else {
			name = "Unknown." + pkg + "." + name
		}
	}
	return name
}
