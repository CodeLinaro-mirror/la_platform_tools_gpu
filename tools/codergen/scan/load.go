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

package scan

import (
	"fmt"
	"go/parser"
	"io/ioutil"
	"path/filepath"
	"regexp"

	"android.googlesource.com/platform/tools/gpu/tools/copyright"

	"golang.org/x/tools/go/types"
)

const Tool = "codergen"

func (m *Module) addSource(filename, content string) error {
	if content == "" {
		file, err := ioutil.ReadFile(filename)
		if err != nil {
			return err
		}
		content = string(file)
	}

	for _, re := range copyright.Generated {
		match := re.FindStringSubmatch(content)
		if len(match) > 1 {
			for _, test := range match[1:] {
				if test == Tool {
					return nil
				}
			}
		}
	}

	m.Sources = append(m.Sources, Source{Filename: filename, Content: content})
	return nil
}

func (s *Scanner) load(dir *Directory) {
	dir.loaded = true
	imp, err := s.context.Import(dir.ImportPath, s.Path, 0)
	if err != nil {
		return
	}
	dir.Name = imp.Name
	dir.ImportPath = imp.ImportPath
	dir.Dir = imp.Dir
	for _, filename := range imp.GoFiles {
		dir.Module.addSource(filepath.Join(dir.Dir, filename), "")
	}
	if dir.Scan {
		for _, filename := range imp.TestGoFiles {
			dir.Test.addSource(filepath.Join(dir.Dir, filename), "")
		}
	}
}

var directive = regexp.MustCompile(`binary: *([^= ]+) *(= *(.+))? *`)

func (s *Scanner) parse(module *Module) error {
	for i := range module.Sources {
		src := &module.Sources[i]
		file, err := parser.ParseFile(s.FileSet, src.Filename, src.Content, parser.ParseComments)
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

func (s *Scanner) importer(pkgs map[string]*types.Package, importPath string) (*types.Package, error) {
	if importPath == "unsafe" {
		pkgs[importPath] = types.Unsafe
		return types.Unsafe, nil
	}
	dir := s.GetDir(importPath)
	if err := s.process(dir); err != nil {
		return nil, err
	}
	if dir.Module.Types == nil {
		return nil, fmt.Errorf("Cyclic import on %s", dir.ImportPath)
	}
	pkgs[importPath] = dir.Module.Types
	return dir.Module.Types, nil
}
