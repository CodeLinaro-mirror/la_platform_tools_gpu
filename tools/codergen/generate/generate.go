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

// Package generate has support processing loaded go code, finding the items
// that require generated code, and converting them to a form the templates can
// easily consume.
package generate

import (
	"fmt"
	"sort"
	"strconv"
	"strings"

	"golang.org/x/tools/go/exact"
	"golang.org/x/tools/go/types"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/scan"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/template"
)

type Generator func(name string, arg interface{}, output string, reflow template.PostProcess) error

const (
	indentRune = "»"
)

func indentor(indent string) template.PostProcess {
	indent = strings.Trim(indent, `"`)
	if indent == "" {
		indent = "    "
	}
	return func(b []byte) []byte {
		return []byte(strings.Replace(string(b), indentRune, indent, -1))
	}
}

type Module struct {
	Source     *scan.Module
	Name       string
	Import     string
	IsTest     bool
	Path       string
	Directives map[string]string
	Structs    []*Struct
	Constants  schema.Constants
	Imports    Imports
}

type Imports map[string]struct{}

func (m *Module) Directive(name string, notset interface{}) interface{} {
	d, ok := m.Directives[name]
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

func From(scanner *scan.Scanner) ([]*Module, error) {
	result := []*Module{}
	for _, dir := range scanner.Directories {
		if !dir.Scan {
			continue
		}
		if m, err := convert(scanner, &dir.Module, false); err != nil {
			return nil, err
		} else if m != nil {
			result = append(result, m)
		}
		if m, err := convert(scanner, &dir.Test, true); err != nil {
			return nil, err
		} else if m != nil {
			result = append(result, m)
		}
	}
	return result, nil
}

func convert(scanner *scan.Scanner, src *scan.Module, isTest bool) (*Module, error) {
	if src.Types == nil {
		return nil, nil
	}
	directives := map[string]string{}
	for _, file := range src.Sources {
		for k, v := range file.Directives {
			directives[k] = v
		}
	}
	if _, ignored := directives["ignore"]; ignored {
		return nil, nil
	}
	m := &Module{
		Source:     src,
		Name:       src.Directory.Name,
		Path:       src.Directory.Dir,
		Import:     src.Directory.ImportPath,
		Imports:    make(map[string]struct{}),
		Directives: directives,
		IsTest:     isTest,
	}
	scope := src.Types.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		f := scanner.FileSet.File(obj.Pos())
		filename := f.Name()
		var source *scan.Source
		for i := range src.Sources {
			if src.Sources[i].Filename == filename {
				source = &src.Sources[i]
				break
			}
		}
		if source == nil {
			continue
		}
		if n, ok := obj.(*types.TypeName); ok {
			if t, ok := n.Type().(*types.Named); ok {
				if _, ok := t.Underlying().(*types.Struct); ok {
					m.addStruct(n)
				}
			}
		}
		if c, ok := obj.(*types.Const); ok && c.Exported() {
			if t, ok := c.Type().(*types.Named); ok {
				if t.Obj().Pkg() == obj.Pkg() {
					if _, ok := c.Type().Underlying().(*types.Basic); ok {
						m.addConst(c)
					}
				}
			}
		}
	}
	sortStructs(m.Structs)
	sort.Sort(&m.Constants)
	for i := range m.Constants {
		sort.Sort(&m.Constants[i])
	}
	return m, nil
}

func (m *Module) addStruct(n *types.TypeName) {
	if s := NewStruct(m.Source.Types, n, m.Imports); s != nil {
		m.Structs = append(m.Structs, s)
	}
}

func (m *Module) addConst(c *types.Const) {
	t := fromType(m.Source.Types, c.Type(), "", m.Imports, nil)
	name := c.Name()
	directive := fmt.Sprintf("%s#%s", t, name)
	if d, found := m.Directives[directive]; found {
		name = d
	} else {
		name = strings.TrimPrefix(name, t.String())
		name = strings.Trim(name, "_")
	}
	if p, ok := t.(*schema.Primitive); ok {
		switch p.Method {
		case schema.Int8:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int8(v),
			})
		case schema.Uint8:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint8(v),
			})
		case schema.Int16:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int16(v),
			})
		case schema.Uint16:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint16(v),
			})
		case schema.Int32:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int32(v),
			})
		case schema.Uint32:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint32(v),
			})
		case schema.Int64:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: v,
			})
		case schema.Uint64:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: v,
			})
		}
	}
}
