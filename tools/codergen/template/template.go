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

package template

import (
	"bytes"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"regexp"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/format"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/generate"
)

// Templates manages the loaded templates and executes them on demand.
// It exposes all it's public methods and fields to the templates.
type Templates struct {
	templates *template.Template
	funcs     template.FuncMap
	active    *template.Template
	writer    io.Writer
	File      interface{}
	counters  map[string]*counter
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

// New constructs and returns a new template set.
// This parses all embedded templates automatically.
func New() *Templates {
	f := &Templates{
		templates: template.New("FunctionHolder"),
		funcs: template.FuncMap{
			// fake builtin functions
			"add": func(a, b int) int { return a + b },
			"sub": func(a, b int) int { return a - b },
		},
		counters: map[string]*counter{},
	}
	v := reflect.ValueOf(f)
	installMethods(v, f.funcs)
	installFields(v.Elem(), f.funcs)
	f.templates.Funcs(f.funcs)
	for name, content := range embedded {
		template.Must(f.templates.New(name).Parse(content))
	}
	return f
}

var (
	sectionMarker = "<<<%s:%s>>>"
	sectionStart  = "Start"
	sectionEnd    = "End"
	section       = regexp.MustCompile(fmt.Sprintf(sectionMarker, "(.+)", "(.+)"))
)

// Generate is an implementation of generate.Generator
func (t *Templates) Generate(g generate.Generate) (bool, error) {
	t.File = g.Arg
	defer func() { t.File = nil }()
	old, _ := ioutil.ReadFile(g.Output)
	buf := &bytes.Buffer{}
	out := format.New(buf)
	out.Indent = g.Indent
	matches := section.FindAllSubmatchIndex(old, -1)
	if len(matches) > 0 {
		last := 0
		tmpl := ""
		// we are doing a partial update...
		for _, match := range matches {
			mode := string(old[match[2]:match[3]])
			name := string(old[match[4]:match[5]])
			switch mode {
			case sectionStart:
				if tmpl != "" {
					return false, fmt.Errorf("Overlapping template %s found starting %s", tmpl, name)
				}
				// section start, write the prefix
				buf.Write(old[last:match[1]])
				// now run the template
				tmpl = name
				if err := t.execute(tmpl, out, g.Arg); err != nil {
					return false, err
				}
				out.Flush()
			case sectionEnd:
				// section end marker, check it matches
				if name != tmpl {
					return false, fmt.Errorf("Invalid end %s found, expected %s", name, tmpl)
				}
				// set the markers ready for the next write
				tmpl = ""
				last = match[0]
			default:
				return false, fmt.Errorf("Invalid section marker %s:%s", mode, name)
			}
		}
		if tmpl != "" {
			return false, fmt.Errorf("Unclosed template %s found", tmpl)
		}
		// write the prefix
		buf.Write(old[last:])
	} else {
		if err := t.execute(g.Name, out, g.Arg); err != nil {
			return false, err
		}
		out.Flush()
	}
	data := buf.Bytes()
	if g.Output == "" || bytes.Equal(data, old) {
		return false, nil
	}
	dir, _ := filepath.Split(g.Output)
	if len(dir) > 0 {
		os.MkdirAll(dir, os.ModePerm)
	}
	return true, ioutil.WriteFile(g.Output, data, 0666)
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
	case nil:
	default:
		return nil, fmt.Errorf("Invalid call dispatch type %T", node)
	}
	if node != nil {
		r := reflect.TypeOf(node)
		// using the reflected typename
		if r.Name() != "" {
			try = append(try, fmt.Sprint(prefix, ".", r.Name()))
		}
		if r.Kind() == reflect.Ptr {
			try = append(try, fmt.Sprint(prefix, ".", r.Elem().Name()))
		}
	} else {
		try = append(try, fmt.Sprint(prefix, ".nil"))
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
