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
	"regexp"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary/schema"

	"golang.org/x/tools/go/exact"
	"golang.org/x/tools/go/gcimporter"
	"golang.org/x/tools/go/types"
)

// Source holds a file a filename content pair as consumed by go/parser.ParseFile.
type Source struct {
	Filename   string            // The filename for this source
	Content    interface{}       // The content of this source, see ParseFiles for details.
	AST        *ast.File         // The parsed syntax tree
	Directives map[string]string // the set of comment overrides
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
	dir.Module.Sources = append(dir.Module.Sources, Source{Filename: filename, Content: source})
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
		l.scan(dir, &dir.Module, false)
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
		l.scan(dir, &dir.Test, true)
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
		dir.Module.Sources = append(dir.Module.Sources, Source{Filename: filepath.Join(dir.Dir, filename)})
	}
	if dir.Scan {
		for _, filename := range imp.TestGoFiles {
			dir.Test.Sources = append(dir.Test.Sources, Source{Filename: filepath.Join(dir.Dir, filename)})
		}
	}
}

var directive = regexp.MustCompile(`binary: *([^= ]+) *(= *(.+))? *`)

func (l *Loader) parse(module *Module) error {
	for i := range module.Sources {
		src := &module.Sources[i]
		file, err := parser.ParseFile(l.FileSet, src.Filename, src.Content, parser.ParseComments)
		if err != nil {
			return err
		}
		src.AST = file
		module.Files = append(module.Files, file)
		src.Directives = make(map[string]string)
		for _, group := range file.Comments {
			for _, comment := range group.List {
				if matches := directive.FindStringSubmatch(comment.Text); len(matches) >= 1 {
					k := matches[1]
					v := matches[3]
					if matches[2] == "" {
						v = "true"
					}
					src.Directives[k] = v
				}
			}
		}
	}
	return nil
}

func (l *Loader) typeCheck(dir *Directory, module *Module) error {
	t, err := l.config.Check(dir.ImportPath, l.FileSet, module.Files, nil)
	module.Types = t
	return err
}

func (l *Loader) scan(dir *Directory, module *Module, isTest bool) error {
	module.Output = &File{
		Package:    dir.Name,
		Path:       dir.Dir,
		Import:     dir.ImportPath,
		Imports:    make(map[string]struct{}),
		Directives: make(map[string]string),
		IsTest:     isTest,
	}
	scope := module.Types.Scope()
	for _, name := range scope.Names() {
		obj := scope.Lookup(name)
		f := l.FileSet.File(obj.Pos())
		filename := f.Name()
		var source *Source
		for i := range module.Sources {
			if module.Sources[i].Filename == filename {
				source = &module.Sources[i]
				break
			}
		}
		if source == nil {
			continue
		}
		if n, ok := obj.(*types.TypeName); ok {
			if t, ok := n.Type().(*types.Named); ok {
				if _, ok := t.Underlying().(*types.Struct); ok {
					if s := FromTypename(module.Types, n, module.Output.Imports); s != nil {
						module.Output.Structs = append(module.Output.Structs, s)
					}
				}
			}
		}
		if c, ok := obj.(*types.Const); ok && c.Exported() {
			if t, ok := c.Type().(*types.Named); ok {
				if t.Obj().Pkg() == obj.Pkg() {
					if _, ok := c.Type().Underlying().(*types.Basic); ok {
						l.addConst(module, source, c)
					}
				}
			}
		}
	}
	for _, s := range module.Sources {
		for k, v := range s.Directives {
			module.Output.Directives[k] = v
		}
	}
	return nil
}

func (l *Loader) addConst(module *Module, source *Source, c *types.Const) {
	t := fromType(module.Types, c.Type(), "", module.Output.Imports)
	name := c.Name()
	directive := fmt.Sprintf("%s#%s", t, name)
	if d, found := source.Directives[directive]; found {
		name = d
	} else {
		name = strings.TrimPrefix(name, t.String())
		name = strings.Trim(name, "_")
	}
	if p, ok := t.(*schema.Primitive); ok {
		switch p.Method {
		case schema.Int8:
			v, _ := exact.Int64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int8(v),
			})
		case schema.Uint8:
			v, _ := exact.Uint64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint8(v),
			})
		case schema.Int16:
			v, _ := exact.Int64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int16(v),
			})
		case schema.Uint16:
			v, _ := exact.Uint64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint16(v),
			})
		case schema.Int32:
			v, _ := exact.Int64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int32(v),
			})
		case schema.Uint32:
			v, _ := exact.Uint64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint32(v),
			})
		case schema.Int64:
			v, _ := exact.Int64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: v,
			})
		case schema.Uint64:
			v, _ := exact.Uint64Val(c.Val())
			module.Output.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: v,
			})
		}
	}
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
