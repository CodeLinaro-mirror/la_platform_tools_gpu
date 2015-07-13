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

import "golang.org/x/tools/imports"

type Go struct {
	*File
	Copyright string
}

func NewGo(file *File) *Go { return &Go{File: file} }

func (file *Go) Run(t *Templates, out string) (bool, error) {
	return t.generate(file, "Go.File", file, out, func(b []byte) []byte {
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
	})
}
