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

	"golang.org/x/tools/imports"
)

// GoFile generates the all the go code for a file with a set of structs.
func (t *Templates) GoFile(file *File) ([]byte, error) {
	b := &bytes.Buffer{}
	t.File = file
	defer func() { t.File = nil }()
	if err := t.execute("Go.File", b, file); err != nil {
		return nil, err
	}
	options := &imports.Options{
		TabWidth:  8,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
	result, err := imports.Process("", b.Bytes(), options)
	if err != nil {
		return b.Bytes(), nil
	}
	return result, nil
}
