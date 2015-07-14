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

	"android.googlesource.com/platform/tools/gpu/tools/copyright"
	"golang.org/x/tools/imports"
)

type GoBinary struct {
	*Module
	Copyright string
}

type GoService struct {
	GoBinary
	Service *Service
}

func goFileName(m *Module, prefix string, category string) string {
	name := prefix + "_" + category
	if m.IsTest {
		name += "_test"
	}
	return path.Join(m.Path, name+".go")
}

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
		if err := gen("Go.Client", arg, goFileName(m, s.Prefix, "client"), reflowGo); err != nil {
			return err
		}
		if err := gen("Go.Server", arg, goFileName(m, s.Prefix, "server"), reflowGo); err != nil {
			return err
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
