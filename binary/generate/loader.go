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

// Package generate has support for generating encode and decode methods
// for the binary package automatically.
package generate

import (
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"

	"golang.org/x/tools/go/types"
)

type Loader struct {
	Path    string
	FileSet *token.FileSet
	Context build.Context
	Config  types.Config
	Imports Imports
}

func NewLoader(path string) *Loader {
	l := &Loader{
		Path:    path,
		FileSet: token.NewFileSet(),
		Context: build.Default,
		Config: types.Config{
			IgnoreFuncBodies: true,
			Error:            func(error) {},
			Packages: map[string]*types.Package{
				"unsafe": types.Unsafe,
			},
		},
		Imports: make(Imports),
	}
	l.Config.Import = l.importer
	return l
}

func (l *Loader) Scan(pkg *types.Package) (*File, error) {
	file := &File{
		Package: pkg.Name(),
		Import:  pkg.Path(),
		Imports: make(map[string]struct{}),
	}
	for _, name := range filterStructs(pkg) {
		fromFile := l.FileSet.File(name.Pos()).Name()
		file.Path = filepath.Dir(fromFile)
		s := FromTypename(pkg, name, file.Imports)
		if s != nil {
			file.Structs = append(file.Structs, s)
		}
	}
	return file, nil
}

func (l *Loader) ScanFile(filename, source string) (*File, error) {
	file, err := parser.ParseFile(l.FileSet, filename, source, 0)
	if err != nil {
		return nil, err
	}
	pkg, err := l.Config.Check("", l.FileSet, []*ast.File{file}, nil)
	if err != nil {
		return nil, err
	}
	return l.Scan(pkg)
}

func (l *Loader) ScanPackage(name string) (*File, error) {
	pkg, err := l.importer(l.Config.Packages, name)
	if err != nil {
		return nil, err
	}
	return l.Scan(pkg)
}

func filterStructs(pkg *types.Package) []*types.TypeName {
	result := []*types.TypeName{}
	scope := pkg.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		if n, ok := obj.(*types.TypeName); ok {
			if t, ok := n.Type().(*types.Named); ok {
				if _, ok := t.Underlying().(*types.Struct); ok {
					result = append(result, n)
				}
			}
		}
	}
	return result
}

func (l *Loader) importer(pkgs map[string]*types.Package, name string) (*types.Package, error) {
	pkg, found := pkgs[name]
	if found {
		return pkg, nil
	}
	imp, err := l.Context.Import(name, l.Path, 0)
	if err != nil {
		return pkg, err
	}
	files := []*ast.File{}
	for _, filename := range imp.GoFiles {
		path := filepath.Join(imp.Dir, filename)
		file, err := parser.ParseFile(l.FileSet, path, nil, 0)
		if err != nil {
			return pkg, err
		}
		files = append(files, file)
	}
	pkg, err = l.Config.Check(imp.ImportPath, l.FileSet, files, nil)
	if err != nil {
		return pkg, err
	}
	pkgs[name] = pkg
	return pkg, nil
}
