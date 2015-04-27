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
	"fmt"
	"go/ast"
	"go/build"
	"go/parser"
	"go/token"
	"path/filepath"

	"golang.org/x/tools/go/gcimporter"
	"golang.org/x/tools/go/types"
)

// Source holds a file a filename content pair as consumed by go/parser.ParseFile.
type Source struct {
	Filename string      // The filename for this source
	Content  interface{} // The content of this source, see ParseFiles for details.
}

// Module represents a resolvable module. Under normal go layout conditions a
// directory has one module that represents the files that are considered
// when the directory is imported, and a second one that also includes the test
// files. They must be considered separately because otherwise you can get
// import cycles.
type Module struct {
	Sources   []Source       // The set of sources to parse
	Files     []*ast.File    // The parsed files included
	Types     *types.Package // The resolved type information
	Output    *File          // The prepared structures to generate for
	processed bool
}

type Directory struct {
	Name       string // The package name (as used in package declarations)
	ImportPath string // The full import path (as used in import statements)
	Dir        string // The actual directory in which the files live
	Scan       bool   // Whether to scan this directory for structs
	Module     Module // The main module data for this directory
	Test       Module // The test module data for this directory
	loaded     bool
}

type Loader struct {
	Path        string                // The base path of file scanning
	ForceSource bool                  // If true, forces source only imports
	Directories map[string]*Directory // The set of directories considered
	FileSet     *token.FileSet        // The parser file set
	context     build.Context
	config      types.Config
}

// NewLoader creates a new go source loader.
func NewLoader(path string, forceSource bool) *Loader {
	l := &Loader{
		Path:        path,
		ForceSource: forceSource,
		Directories: map[string]*Directory{},
		FileSet:     token.NewFileSet(),
		context:     build.Default,
		config: types.Config{
			IgnoreFuncBodies: true,
			Error:            func(error) {},
			Packages:         map[string]*types.Package{},
		},
	}
	l.config.Import = l.importer
	return l
}

// GetDir returns the Directory that matches the supplied import path.
// It will add a new one if needed.
func (l *Loader) GetDir(importPath string) *Directory {
	dir, ok := l.Directories[importPath]
	if !ok {
		dir = &Directory{
			ImportPath: importPath,
		}
		l.Directories[importPath] = dir
	}
	return dir
}

// ScanFile adds a fake package with the file as it's only source.
func (l *Loader) ScanFile(filename, source string) {
	dir := l.GetDir(filename)
	dir.Scan = true
	dir.loaded = true
	dir.Module.Sources = append(dir.Module.Sources, Source{filename, source})
}

// ScanPackage marks the directory specified by the import path as needing to be
// scanned for binary structures.
func (l *Loader) ScanPackage(importPath string) {
	l.GetDir(importPath).Scan = true
}

// Process generates output data for all packages that have been marked as
// needing to be scanned.
func (l *Loader) Process() error {
	dirs := make([]*Directory, 0, len(l.Directories))
	for _, dir := range l.Directories {
		if dir.Scan {
			dirs = append(dirs, dir)
		}
	}
	for _, dir := range dirs {
		if err := l.process(dir); err != nil {
			return err
		}
	}
	return nil
}

func (l *Loader) process(dir *Directory) error {
	if !dir.Scan && !dir.Module.processed && !l.ForceSource {
		t, err := gcimporter.Import(l.config.Packages, dir.ImportPath)
		if err == nil {
			dir.Module.processed = true
			dir.Module.Types = t
		}
	}
	if !dir.loaded {
		l.load(dir)
	}
	if !dir.Module.processed {
		dir.Module.processed = true
		if err := l.parse(&dir.Module); err != nil {
			return err
		}
		if err := l.typeCheck(dir, &dir.Module); err != nil {
			return err
		}
		l.config.Packages[dir.ImportPath] = dir.Module.Types
		l.scan(dir, &dir.Module)
	}
	if dir.Scan && !dir.Test.processed && len(dir.Test.Sources) > 0 {
		dir.Test.processed = true
		dir.Test.Files = dir.Module.Files
		if err := l.parse(&dir.Test); err != nil {
			return err
		}
		if err := l.typeCheck(dir, &dir.Test); err != nil {
			return err
		}
		l.scan(dir, &dir.Test)
		dir.Test.Output.IsTest = true
	}
	return nil
}

func (l *Loader) load(dir *Directory) {
	dir.loaded = true
	imp, err := l.context.Import(dir.ImportPath, l.Path, 0)
	if err != nil {
		return
	}
	dir.Name = imp.Name
	dir.ImportPath = imp.ImportPath
	dir.Dir = imp.Dir
	for _, filename := range imp.GoFiles {
		dir.Module.Sources = append(dir.Module.Sources, Source{filepath.Join(dir.Dir, filename), nil})
	}
	if dir.Scan {
		for _, filename := range imp.TestGoFiles {
			dir.Test.Sources = append(dir.Test.Sources, Source{filepath.Join(dir.Dir, filename), nil})
		}
	}
}

func (l *Loader) parse(module *Module) error {
	for _, src := range module.Sources {
		file, err := parser.ParseFile(l.FileSet, src.Filename, src.Content, 0)
		if err != nil {
			return err
		}
		module.Files = append(module.Files, file)
	}
	return nil
}

func (l *Loader) typeCheck(dir *Directory, module *Module) error {
	t, err := l.config.Check(dir.ImportPath, l.FileSet, module.Files, nil)
	module.Types = t
	return err
}

func (l *Loader) scan(dir *Directory, module *Module) error {
	module.Output = &File{
		Package: dir.Name,
		Path:    dir.Dir,
		Import:  dir.ImportPath,
		Imports: make(map[string]struct{}),
	}
	for _, name := range filterStructs(module.Types) {
		filename := l.FileSet.File(name.Pos()).Name()
		found := false
		for _, f := range module.Sources {
			if f.Filename == filename {
				found = true
				break
			}
		}
		if !found {
			continue
		}
		s := FromTypename(module.Types, name, module.Output.Imports)
		if s != nil {
			module.Output.Structs = append(module.Output.Structs, s)
		}
	}
	return nil
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

func (l *Loader) importer(pkgs map[string]*types.Package, importPath string) (*types.Package, error) {
	if importPath == "unsafe" {
		pkgs[importPath] = types.Unsafe
		return types.Unsafe, nil
	}
	dir := l.GetDir(importPath)
	if err := l.process(dir); err != nil {
		return nil, err
	}
	if dir.Module.Types == nil {
		return nil, fmt.Errorf("Cyclic import on %s", dir.ImportPath)
	}
	pkgs[importPath] = dir.Module.Types
	return dir.Module.Types, nil
}
