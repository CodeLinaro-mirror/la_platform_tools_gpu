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
	"strings"

	"android.googlesource.com/platform/tools/gpu/api/resolver"
)

type Cpp struct {
	*File
	Namespace string
	Copyright string
	Indent    string
}

func (t *Templates) CppName(n string) string {
	n = strings.Replace(n, ".", "::", -1)
	n = strings.Replace(n, resolver.ConstSuffix+resolver.PointerSuffix, "__CP", -1)
	n = strings.Replace(n, resolver.PointerSuffix, "__P", -1)
	n = strings.Replace(n, resolver.SliceSuffix, "__S", -1)
	n = strings.Replace(n, resolver.ArraySuffix, "__A", -1)
	n = strings.Replace(n, resolver.TypeInfix, "__", -1)
	return n
}

func NewCpp(file *File) *Cpp { return &Cpp{File: file, Indent: "    "} }
func (file *Cpp) Run(t *Templates, out string) (bool, error) {
	return t.generate(file.File, "Cpp.File", file, out, func(b []byte) []byte {
		return []byte(strings.Replace(string(b), indent, file.Indent, -1))
	})
}
