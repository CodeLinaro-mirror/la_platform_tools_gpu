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
	"io"
	"reflect"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

type functions struct {
	templates *template.Template
	funcs     template.FuncMap
	active    *template.Template
	prefix    string
	schema    string
	writer    io.Writer
}

func newFunctions() *functions {
	f := &functions{
		templates: template.New("FunctionHolder"),
		funcs:     template.FuncMap{},
	}
	v := reflect.ValueOf(f)
	t := v.Type()
	for i := 0; i < t.NumMethod(); i++ {
		m := t.Method(i)
		r, _ := utf8.DecodeRuneInString(m.Name)
		if unicode.IsUpper(r) {
			c := v.MethodByName(m.Name)
			f.funcs[m.Name] = c.Interface()
		}
	}
	f.templates.Funcs(f.funcs)
	template.Must(f.templates.New("go.tmpl").Parse(string(go_tmpl)))
	template.Must(f.templates.New("java.tmpl").Parse(string(java_tmpl)))
	return f
}

func (f *functions) getTemplate(action string, t schema.Type) *template.Template {
	kind := reflect.TypeOf(t).Elem().Name()
	name := fmt.Sprint(f.prefix, action, kind)
	result := f.templates.Lookup(name)
	if result == nil {
		panic(fmt.Errorf("Could not find template %s", name))
	}
	return result
}

func (f *functions) execute(name string, w io.Writer, data interface{}) error {
	oldw := f.writer
	if w != nil {
		f.writer = w
	}
	defer func() { f.writer = oldw }()
	t := f.templates.Lookup(name)
	if t == nil {
		return fmt.Errorf("Cannot find template %s", name)
	}
	return t.Execute(w, data)
}

type field struct {
	Name string
	Type interface{}
}

func (f *functions) Encode(name string, t schema.Type) (string, error) {
	return "", f.getTemplate("Encode", t).Execute(f.writer, field{name, t})
}

func (f *functions) Decode(name string, t schema.Type) (string, error) {
	return "", f.getTemplate("Decode", t).Execute(f.writer, field{name, t})
}

func (f *functions) Skip(name string, t schema.Type) (string, error) {
	return "", f.getTemplate("Skip", t).Execute(f.writer, field{name, t})
}

func (f *functions) Schema(t schema.Type) (string, error) {
	return "", f.getTemplate("Schema", t).Execute(f.writer, t)
}

func (f *functions) SchemaPrefix() string {
	return f.schema
}

func (f *functions) Header(tool string) (string, error) {
	_, err := io.WriteString(f.writer, copyright.Build("generated_by", copyright.Info{Tool: tool}))
	return "", err
}

func (f *functions) Lower(s string) string {
	return strings.ToLower(s)
}

func (f *functions) ToS8(val byte) string {
	return fmt.Sprint(int8(val))
}
