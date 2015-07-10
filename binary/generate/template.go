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
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"strconv"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

type Templates struct {
	templates *template.Template
	funcs     template.FuncMap
	active    *template.Template
	writer    io.Writer
	File      interface{}
	counters  map[string]*counter
}

type counter int

func (c *counter) Set(value int) string {
	*c = counter(value)
	return ""
}

func (c *counter) AddLen(value string) string {
	*c += counter(len(value))
	return ""
}

func (c *counter) String() string {
	return fmt.Sprint(*c)
}

func isPublic(s string) bool {
	r, _ := utf8.DecodeRuneInString(s)
	return unicode.IsUpper(r)
}

func installMethods(v reflect.Value, funcs template.FuncMap) {
	ty := v.Type()
	for i := 0; i < ty.NumMethod(); i++ {
		m := ty.Method(i)
		if isPublic(m.Name) {
			funcs[m.Name] = v.Method(i).Interface()
		}
	}
}

func installFields(v reflect.Value, funcs template.FuncMap) {
	ty := v.Type()
	for i := 0; i < ty.NumField(); i++ {
		m := ty.Field(i)
		if isPublic(m.Name) {
			funcs[m.Name] = v.Field(i).Interface
		}
	}
}

func NewTemplates() *Templates {
	f := &Templates{
		templates: template.New("FunctionHolder"),
		funcs:     template.FuncMap{},
		counters:  map[string]*counter{},
	}
	v := reflect.ValueOf(f)
	installMethods(v, f.funcs)
	installFields(v.Elem(), f.funcs)
	f.templates.Funcs(f.funcs)
	template.Must(f.templates.New("go.tmpl").Parse(string(go_tmpl)))
	template.Must(f.templates.New("java.tmpl").Parse(string(java_tmpl)))
	template.Must(f.templates.New("cpp.tmpl").Parse(string(cpp_tmpl)))
	return f
}

type PostProcess func([]byte) []byte

func (t *Templates) generate(f interface{}, name string, arg interface{}, out string, post PostProcess) (bool, error) {
	t.File = f
	defer func() { t.File = nil }()

	b := &bytes.Buffer{}
	if err := t.execute(name, b, arg); err != nil {
		return false, err
	}
	data := post(b.Bytes())
	current, err := ioutil.ReadFile(out)
	if err == nil && bytes.Equal(data, current) {
		return false, nil
	}
	if out == "" {
		return false, nil
	}
	dir, _ := filepath.Split(out)
	if len(dir) > 0 {
		os.MkdirAll(dir, os.ModePerm)
	}
	return true, ioutil.WriteFile(out, data, 0666)
}

func (t *Templates) getTemplate(prefix string, node interface{}) (*template.Template, error) {
	try := []string{}

	switch node := node.(type) {
	case schema.Type:
		try = append(try, fmt.Sprint(prefix, "#", node.Typename()))
		if node.Typename() != node.Basename() {
			try = append(try, fmt.Sprint(prefix, "#", node.Basename()))
		}
	case *variable:
		return t.getTemplate(prefix, node.Type)
	case string:
		try = append(try, prefix+node)
	case schema.Method:
		try = append(try, fmt.Sprint(prefix, "#", node.String()))
	default:
		return nil, fmt.Errorf("Invalid call dispatch type %T", node)
	}
	r := reflect.TypeOf(node)
	// using the reflected typename
	try = append(try, fmt.Sprint(prefix, ".", r.Name()))
	if r.Kind() == reflect.Ptr {
		try = append(try, fmt.Sprint(prefix, ".", r.Elem().Name()))
	}
	// default case is just the prefix
	try = append(try, prefix)
	for _, name := range try {
		if tmpl := t.templates.Lookup(name); tmpl != nil {
			return tmpl, nil
		}
	}
	return nil, fmt.Errorf(`Cannot find templates "%s"`, strings.Join(try, `","`))
}

func (t *Templates) execute(name string, w io.Writer, data interface{}) error {
	oldw := t.writer
	if w != nil {
		t.writer = w
	}
	defer func() { t.writer = oldw }()
	tmpl := t.templates.Lookup(name)
	if tmpl == nil {
		return fmt.Errorf("Cannot find template %s", name)
	}
	return tmpl.Execute(w, data)
}

type variable struct {
	Name string
	Type interface{}
}

func (*Templates) Var(t schema.Type, args ...interface{}) *variable {
	return &variable{
		Name: fmt.Sprint(args...),
		Type: t,
	}
}

func (t *Templates) Call(prefix string, arg interface{}) (string, error) {
	tmpl, err := t.getTemplate(prefix, arg)
	if err != nil {
		return "", err
	}
	return "", tmpl.Execute(t.writer, arg)
}

func (*Templates) Lower(s interface{}) string {
	return strings.ToLower(fmt.Sprint(s))
}

func (*Templates) Upper(s interface{}) string {
	return strings.ToUpper(fmt.Sprint(s))
}

func (*Templates) Contains(test, s interface{}) bool {
	return strings.Contains(fmt.Sprint(s), fmt.Sprint(test))
}

func (*Templates) ToS8(val byte) string {
	return fmt.Sprint(int8(val))
}

func (f *File) Directive(name string, notset interface{}) interface{} {
	d, ok := f.Directives[name]
	if !ok {
		return notset
	}
	if _, isbool := notset.(bool); isbool {
		//coerce the string to bool
		if b, err := strconv.ParseBool(d); err == nil {
			return b
		}
	}
	return d
}

func (t *Templates) Counter(name string) *counter {
	c, ok := t.counters[name]
	if !ok {
		c = new(counter)
		t.counters[name] = c
	}
	return c
}
