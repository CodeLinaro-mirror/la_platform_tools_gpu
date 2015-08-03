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
	"path"
	"strings"

	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

// GoBinary is the struct handed to binary coder generation templates.
type GoBinary struct {
	*Module
	Copyright string
	Imports   Imports
}

// GoService is the struct handed to go rpc service generation templates.
type GoService struct {
	GoBinary          // Needs all the same information as the binary templates
	Service  *Service // The service to generate for.
}

func goFileName(m *Module, prefix string, category string) string {
	name := prefix + "_" + category
	if m.IsTest {
		name += "_test"
	}
	return path.Join(m.Path, name+".go")
}

// Go is called by codergen to prepare and generate go code for a given module.
func Go(m *Module, info copyright.Info, gen chan Generate) error {
	if len(m.Structs) == 0 && len(m.Constants) == 0 {
		return nil
	}
	gen <- Generate{
		Name: "Go.Binary",
		Arg: &GoBinary{
			Module:    m,
			Copyright: copyright.Build("generated_by", info),
		},
		Output: goFileName(m, m.Name, "binary"),
		Indent: "\t",
	}
	for _, s := range m.Services {
		for _, e := range []string{"client", "server", "helpers", "extra"} {
			gen <- Generate{
				Name: "Go." + strings.Title(e),
				Arg: &GoService{
					GoBinary: GoBinary{
						Module:    m,
						Copyright: copyright.Build("generated_by", info),
					},
					Service: s,
				},
				Output: goFileName(m, s.Prefix, e),
				Indent: "\t",
			}
		}
	}
	return nil
}

// Import adds an import to the import set for this template.
func (b *GoBinary) Import(path string) string {
	b.Imports.Add(Import{Name: "", Path: path})
	return ""
}

// ImportOwner adds an import of the package that owns the supplied object.
func (b *GoBinary) ImportOwner(v interface{}) string {
	m, _ := b.ModuleAndName(v)
	if m != b.Module {
		b.Imports.Add(Import{Name: "", Path: m.Import})
	}
	return ""
}
