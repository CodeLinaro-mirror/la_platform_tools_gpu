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

type GoPackage struct {
	*Module
	Copyright string
}

func Go(m *Module, info copyright.Info, gen Generator) error {
	if len(m.Structs) == 0 && len(m.Constants) == 0 {
		return nil
	}
	pkg := &GoPackage{
		Module:    m,
		Copyright: copyright.Build("generated_by", info),
	}
	out := m.Name + "_binary.go"
	if m.IsTest {
		out = m.Name + "_binary_test.go"
	}
	out = path.Join(m.Path, out)
	return gen("Go.File", pkg, out, reflowGo)
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
