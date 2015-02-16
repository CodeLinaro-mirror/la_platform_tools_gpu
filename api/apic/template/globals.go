// Copyright (C) 2015 The Android Open Source Project
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
	"path/filepath"
	"strings"
)

var (
	globalList stringSetFlag
)

type globalMap map[string]interface{}

type stringSetFlag []string

func (f *stringSetFlag) String() string    { return strings.Join(f.Strings(), ":") }
func (f *stringSetFlag) Strings() []string { return ([]string)(*f) }

func (f *stringSetFlag) Set(value string) error {
	*f = append(*f, value)
	return nil
}

func initGlobals(f *Functions) {
	apiBase := filepath.Base(f.apiFile)
	f.globals["API"] = strings.TrimSuffix(apiBase, filepath.Ext(apiBase))
	f.globals["OutputDir"] = filepath.Base(f.basePath)
	f.globals["OutputPath"] = f.basePath
	for _, g := range globalList.Strings() {
		v := strings.SplitN(g, "=", 2)
		f.globals[v[0]] = v[1]
	}
}

// Gets or sets a template global variable
// Example:
//  {{Global "CatSays" "Meow"}}
//  The cat says: {{Global "CatSays"}}
func (f *Functions) Global(name string, values ...interface{}) (interface{}, error) {
	switch len(values) {
	case 0:
		value, _ := f.globals[name]
		return value, nil
	case 1:
		f.globals[name] = values[0]
		return "", nil
	default:
		f.globals[name] = values
		return "", nil
	}
}
