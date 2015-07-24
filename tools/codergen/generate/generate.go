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
	"path"
	"strconv"
	"strings"

	"golang.org/x/tools/go/types"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/scan"
	"android.googlesource.com/platform/tools/gpu/tools/codergen/template"
)

// Generator is the func handed in to language specific code generation functions.
// They will call the generator once per output file they want to produce.
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

// Modules holds a the list of all modules in a single scan.
type Modules []*Module

// Module represents a go package. In normal go layout, there will be at most
// two modules per directory, one normal module and one test module.
type Module struct {
	Modules    *Modules          // The set of modules this module belongs to.
	Source     *scan.Module      // The source module this was generated from.
	Name       string            // The name of the module.
	Import     string            // The import path of this module.
	IsTest     bool              // Wether this module is for a test package.
	Path       string            // The directory name this module was scanned from.
	Directives map[string]string // The set of codergen directives encountered in the files.
	Structs    []*Struct         // The structs encountered.
	Constants  schema.Constants  // All the const declarations and their types.
	Services   []*Service        // The service interfaces discovered.
	Imports    Imports           // The set of package imports encountered.
}

// Import represents a go import declaration.
type Import struct {
	Name string // The name if present.
	Path string // The full import path.
}

// Imports represetns a list of Import declarations.
type Imports []Import

// Directive looks up a directive by name, and returns notset if the directive
// is not found.
// It will make an attempt to coerce the return type to match that of notset if
// it is a bool.
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

// Add adds a new import to the import list.
func (i *Imports) Add(v Import) {
	*i = append(*i, v)
}

// FindName returns the import that matches the supplied name, or an empty
// import if not present. Test the returned .Path to detect this.
func (i Imports) FindName(name string) Import {
	for _, e := range i {
		if e.Name == name {
			return e
		}
	}
	return Import{}
}

// FindPath finds the import for the specified import path, or an empty
// import if not present. Test the returned .Path to detect this.
func (i Imports) FindPath(path string) Import {
	for _, e := range i {
		if e.Path == path {
			return e
		}
	}
	return Import{}
}

// FindImport searches the modules imports for the specified name, and then
// searches the parent module set for the module that matches the import path
// found.
// It will return nil if either the name is not valid or the module cannot be
// found.
func (m *Module) FindImport(name string) *Module {
	path := m.Imports.FindName(name).Path
	if path == "" {
		return nil
	}
	for _, o := range *m.Modules {
		if o.Import == path {
			return o
		}
	}
	return nil
}

// From processes scanned source code to produce the module set it represents.
func From(scanner *scan.Scanner) (Modules, error) {
	result := Modules{}
	for _, dir := range scanner.Directories {
		if !dir.Scan {
			continue
		}
		if m, err := convert(scanner, &dir.Module, false); err != nil {
			return nil, err
		} else if m != nil {
			m.Modules = &result
			result = append(result, m)
		}
		if m, err := convert(scanner, &dir.Test, true); err != nil {
			return nil, err
		} else if m != nil {
			m.Modules = &result
			result = append(result, m)
		}
	}
	return result, nil
}

// convert processes a single module from a scan set.
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
		Import:     path.Clean(src.Directory.ImportPath),
		Directives: directives,
		IsTest:     isTest,
	}
	b := findBinaryObject(m.Source.Types)
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
					m.addStruct(n, b)
				}
				if _, ok := t.Underlying().(*types.Interface); ok {
					if err := m.addService(n, b); err != nil {
						return nil, err
					}
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
	m.finaliseStructs()
	m.finaliseConstants()
	m.finaliseServices()
	return m, nil
}
