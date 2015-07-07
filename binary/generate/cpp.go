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
	"bytes"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api/resolver"
)

func (t *Templates) CppName(n string) string {
	n = strings.Replace(n, ".", "::", -1)
	n = strings.Replace(n, resolver.ConstSuffix+resolver.PointerSuffix, "__CP", -1)
	n = strings.Replace(n, resolver.PointerSuffix, "__P", -1)
	n = strings.Replace(n, resolver.SliceSuffix, "__S", -1)
	n = strings.Replace(n, resolver.ArraySuffix, "__A", -1)
	n = strings.Replace(n, resolver.TypeInfix, "__", -1)
	return n
}

// CppFile generates the all the cpp code for a file with a set of structs.
func (t *Templates) CppFile(file *File) ([]byte, error) {
	f := *file
	if f.Indent == "" {
		f.Indent = "    "
	}
	b := &bytes.Buffer{}
	if err := t.execute("Cpp.File", b, &f); err != nil {
		return nil, err
	}
	s := b.String()
	s = strings.Replace(s, indent, f.Indent, -1)
	return []byte(s), nil
}
