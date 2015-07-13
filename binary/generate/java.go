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
	"unicode"
	"unicode/utf8"
)

func (*Templates) JavaFieldName(s string) string {
	r, n := utf8.DecodeRuneInString(s)
	return memberPrefix + string(unicode.ToUpper(r)) + s[n:]
}

type Java struct {
	*File
	Struct          *Struct
	BasePackage     string
	RelativePackage string
	Copyright       string
	Indent          string
	MemberPrefix    string
	ClassPrefix     string
}

func NewJava(file *File) *Java { return &Java{File: file} }
func (file *Java) Run(t *Templates, out string) (bool, error) {
	return t.generate(file, "Java.File", file, out, func(b []byte) []byte {
		s := string(b)
		s = strings.Replace(s, indent, file.Indent, -1)
		s = strings.Replace(s, memberPrefix, file.MemberPrefix, -1)
		s = strings.Replace(s, classPrefix, file.ClassPrefix, -1)
		return []byte(s)
	})
}
