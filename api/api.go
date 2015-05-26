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

// Package api holds the main interface to the api language libraries.
// It provides functions for going from api files to abstract syntax trees and
// processed semantic trees.
package api

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"sort"

	"android.googlesource.com/platform/tools/gpu/api/ast"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/parse"
)

// Processor holds the state when resolving multiple api files.
type Processor struct {
	Parsed   map[string]*ast.API
	Resolved map[string]*semantic.API
}

// DefaultProcessor is the Processor used in the package level functions.
// Most applications will not need multiple instances of a Processor, and can
// just use this one.
var DefaultProcessor Processor

// Parse parses the api file with the DefaultProcessor.
// See Processor.Parse for details.
func Parse(apiname string) (*ast.API, parse.ErrorList) {
	return DefaultProcessor.Parse(apiname)
}

// Parse returns an ast that represents the supplied filename.
// It if the file has already been parsed, the cached ast will be returned,
// otherwise it invokes parser.Parse on the content of the supplied file name.
func (p *Processor) Parse(path string) (*ast.API, parse.ErrorList) {
	if api, ok := p.Parsed[path]; ok {
		return api, nil
	}
	info, err := ioutil.ReadFile(path)
	if err != nil {
		return nil, parse.ErrorList{parse.Error{Message: err.Error()}}
	}
	return parser.Parse(string(info))
}

// Resolve resolves the api file with the DefaultProcessor.
// See Processor.Resolve for details.
func Resolve(apiname string, mappings resolver.ASTToSemantic) (*semantic.API, parse.ErrorList) {
	return DefaultProcessor.Resolve(apiname, mappings)
}

// Resolve returns a semantic.API that represents the supplied api file name.
// If the file has already been resolved, the cached semantic tree is returned,
// otherwise the file and all dependant files are parsed using Processor.Parse.
// Recursive calls are made to Resolve for all named imports, and then finally
// the ast and all included ast's are handed to resolver.Resolve to do semantic
// processing.
func (p *Processor) Resolve(apiname string, mappings resolver.ASTToSemantic) (*semantic.API, parse.ErrorList) {
	absname, err := filepath.Abs(apiname)
	if err != nil {
		return nil, parse.ErrorList{parse.Error{Message: err.Error()}}
	}
	if api, ok := p.Resolved[absname]; ok {
		return api, nil
	}
	// Parse all the includes
	wd, name := filepath.Split(absname)
	includes := map[string]*ast.API{}
	errs := p.include(includes, wd, name)
	if len(errs) > 0 {
		return nil, errs
	}
	// Build a sorted list of includes
	names := make(sort.StringSlice, 0, len(includes))
	for name := range includes {
		names = append(names, name)
	}
	names.Sort()
	list := make([]*ast.API, len(names))
	for i, name := range names {
		list[i] = includes[name]
	}
	// Resolve all the imports
	imports := map[string]*semantic.API{}
	for _, api := range list {
		for _, i := range api.Imports {
			if i.Name == nil {
				// unnamed imports have already been included
				continue
			}
			if _, seen := imports[i.Name.Value]; seen {
				return nil, parse.ErrorList{parse.Error{
					Message: fmt.Sprintf("Duplicate import %s", i.Name.Value)},
				}
			}
			api, errs := Resolve(i.Path.Value, mappings)
			if len(errs) > 0 {
				return nil, errs
			}
			imports[i.Name.Value] = api
		}
	}
	// Now resolve the api set as a single unit
	return resolver.Resolve(list, imports, mappings)
}

func (p *Processor) include(includes map[string]*ast.API, wd string, apiname string) parse.ErrorList {
	absname, err := filepath.Abs(filepath.Join(wd, apiname))
	if err != nil {
		return parse.ErrorList{parse.Error{Message: err.Error()}}
	}
	if _, seen := includes[absname]; seen {
		return nil
	}
	api, errs := p.Parse(absname)
	if len(errs) > 0 {
		return errs
	}
	includes[absname] = api
	for _, i := range api.Imports {
		if i.Name != nil {
			// named imports don't get merged
			continue
		}
		errs := p.include(includes, wd, i.Path.Value)
		if len(errs) > 0 {
			return errs
		}
	}
	return nil
}
