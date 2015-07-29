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
	"path/filepath"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

// CppNamespace is the struct handed to the Cpp.File template.
type CppNamespace struct {
	*Module          // The go module to generate cpp binary coders for.
	Namespace string // The name to use for the cpp namespace itself.
	Copyright string // The copyright header to put on the file.
}

// Called by codergen to prepare and generate cpp code for a given module.
func Cpp(m *Module, info copyright.Info, gen chan Generate, path string) {
	namespace := m.Directives["cpp"]
	gen <- Generate{
		Name: "Cpp.File",
		Arg: CppNamespace{
			Module:    m,
			Namespace: namespace,
			Copyright: strings.TrimSpace(copyright.Build("generated_by", info)),
		},
		Output: filepath.Join(path, namespace+".h"),
		Reflow: indentor("    "),
	}
}

// Converts a typename to cpp form by replacing the unicode characters.
func (CppNamespace) TypeName(n string) string {
	n = strings.Replace(n, ".", "::", -1)
	n = strings.Replace(n, resolver.ConstSuffix+resolver.PointerSuffix, "__CP", -1)
	n = strings.Replace(n, resolver.PointerSuffix, "__P", -1)
	n = strings.Replace(n, resolver.SliceSuffix, "__S", -1)
	n = strings.Replace(n, resolver.ArraySuffix, "__A", -1)
	n = strings.Replace(n, resolver.TypeInfix, "__", -1)
	return n
}
