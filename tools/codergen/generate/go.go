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
	"golang.org/x/tools/imports"
)

// GoBinary is the struct handed to binary coder generation templates.
type GoBinary struct {
	*Module
	Copyright string
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
func Go(m *Module, info copyright.Info, gen Generator) error {
	if len(m.Structs) == 0 && len(m.Constants) == 0 {
		return nil
	}
	pkg := GoBinary{
		Module:    m,
		Copyright: copyright.Build("generated_by", info),
	}
	if err := gen("Go.Binary", pkg, goFileName(m, m.Name, "binary"), reflowGo); err != nil {
		return err
	}
	for _, s := range m.Services {
		arg := GoService{GoBinary: pkg, Service: s}
		for _, e := range []string{"client", "server", "helpers", "extra"} {
			if err := gen("Go."+strings.Title(e), arg, goFileName(m, s.Prefix, e), reflowGo); err != nil {
				return err
			}
		}
	}
	return nil
}

func reflowGo(b []byte) []byte {
	options := &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
	if result, err := imports.Process("", b, options); err != nil {
		return b
	} else {
		return result
	}
}
