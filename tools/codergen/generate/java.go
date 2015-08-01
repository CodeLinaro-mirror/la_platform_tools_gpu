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

	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

// JavaSettings holds the information and methods common to all java generators.
type JavaSettings struct {
	*Module                          // The module we are generating for.
	JavaPackage  string              // The java package name for this module.
	Copyright    string              // The copyright header to put on generatated files.
	MemberPrefix string              // The prefix to give all field members of classes.
	Imported     map[string]struct{} // The set of imports generated for the file.
}

// JavaClass is the struct handed to java binary coder generation templates.
type JavaClass struct {
	JavaSettings
	Struct *Struct
}

// JavaService is the struct handed to java rpc service generation templates.
type JavaService struct {
	JavaSettings
	Service *Service
}

// JavaEnum is the struct handed to java enum generation templates.
type JavaEnum struct {
	JavaSettings
	schema.ConstantSet
}

// Java is called by codergen to prepare and generate java code for a given module.
func Java(m *Module, info copyright.Info, gen chan Generate, path string) {
	settings := JavaSettings{
		Module:      m,
		JavaPackage: m.Directives["java.package"],
		Copyright:   strings.TrimSpace(copyright.Build("generated_aosp_java", info)),
	}
	settings.MemberPrefix, _ = m.Directives["java.member_prefix"]
	source, _ := m.Directives["java.source"]
	indent, _ := m.Directives["java.indent"]
	pkgPath := strings.Replace(settings.JavaPackage, ".", "/", -1)
	for _, s := range m.Structs {
		gen <- Generate{
			Name:   "Java.File",
			Arg:    JavaClass{JavaSettings: settings, Struct: s},
			Output: filepath.Join(path, source, pkgPath, settings.ClassName(s.Name)+".java"),
			Reflow: indentor(indent),
		}
	}
	for _, s := range m.Constants {
		gen <- Generate{
			Name:   "Java.Enum",
			Arg:    JavaEnum{JavaSettings: settings, ConstantSet: s},
			Output: filepath.Join(path, source, pkgPath, settings.ClassName(s.Type.String())+".java"),
			Reflow: indentor(indent),
		}
	}
	for _, s := range m.Services {
		for _, e := range []string{"Client", "ClientImpl"} {
			gen <- Generate{
				Name:   "Java." + e,
				Arg:    JavaService{JavaSettings: settings, Service: s},
				Output: filepath.Join(path, source, pkgPath, s.Name+e+".java"),
				Reflow: indentor(indent),
			}
		}
	}
}

// FieldName converts from a go struct field name to the correct java member name.
func (settings JavaSettings) FieldName(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return settings.MemberPrefix + string(unicode.ToUpper(r)) + s[n:]
}

// Getter converts from a go struct field name to the correct java getter name.
func (settings JavaSettings) Getter(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return "get" + string(unicode.ToUpper(r)) + s[n:]
}

// Setter converts from a go struct field name to the correct java getter name.
func (settings JavaSettings) Setter(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return "set" + string(unicode.ToUpper(r)) + s[n:]
}

// returns the module if found, the extracted type name and the modified java name
func (settings JavaSettings) moduleAndName(v interface{}) (*Module, string, string) {
	name := ""
	switch v := v.(type) {
	case string:
		name = v
	case *Struct:
		name = v.Name
	case Call:
		name = v.Struct.Name
	case Result:
		name = v.Struct.Name
	case *schema.Struct:
		name = v.Name
	case *schema.Interface:
		name = v.Name
	case *schema.Primitive:
		name = v.Name
	default:
		panic(fmt.Errorf("Invalid type %T to moduleAndName", v))
	}
	pkg, name := "", name
	if i := strings.LastIndexAny(name, "."); i >= 0 {
		pkg = name[:i]
		name = name[i+1:]
	}
	java := strings.Title(name)
	java = strings.Title(strings.Replace(name, "_", "", -1))
	m := settings.Module
	if pkg != "" && pkg != "binary" {
		m = settings.FindImport(pkg)
		if m == nil {
			m = &Module{Name: pkg}
		}
	}
	return m, name, java
}

func (settings JavaSettings) findClass(v interface{}) (string, string) {
	m, original, name := settings.moduleAndName(v)
	for _, t := range m.Structs {
		if t.Name == original {
			if n := fmt.Sprint(t.Tags.Get("java")); n != "" {
				name = n
			} else {
				name += fmt.Sprint(m.Directive("java.class_suffix", ""))
			}
		}
	}
	if m == settings.Module {
		return "", name
	}
	return fmt.Sprint(m.Directive("java.package", "NotJava."+m.Name)), name
}

// Import returns the fully qualified name for the type, if not previously encountered.
func (settings JavaSettings) Import(v interface{}) string {
	pkg, name := settings.findClass(v)
	if pkg == "" {
		return "" // Not an import
	}
	fullname := pkg + "." + name
	if settings.Imported == nil {
		settings.Imported = map[string]struct{}{}
	} else if _, ok := settings.Imported[fullname]; ok {
		return "" // Already imported
	}
	settings.Imported[fullname] = struct{}{}
	return fullname
}

// ClassName returns the Java name to give the class type.
func (settings JavaSettings) ClassName(v interface{}) string {
	_, name := settings.findClass(v)
	return name
}

// InterfaceName returns the Java name to give the interface type.
func (settings JavaSettings) InterfaceName(v interface{}) string {
	_, _, name := settings.moduleAndName(v)
	return name
}
