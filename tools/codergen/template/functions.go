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

package template

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
)

type variable struct {
	Name string
	Type interface{}
}

func (*Templates) Var(t binary.Type, args ...interface{}) *variable {
	return &variable{
		Name: fmt.Sprint(args...),
		Type: t,
	}
}

func (t *Templates) Call(prefix string, arg interface{}) (string, error) {
	tmpl, err := t.getTemplate(prefix, arg)
	if err != nil {
		return "", err
	}
	return "", tmpl.Execute(t.writer, arg)
}

func (*Templates) Lower(s interface{}) string {
	return strings.ToLower(fmt.Sprint(s))
}

func (*Templates) Upper(s interface{}) string {
	return strings.ToUpper(fmt.Sprint(s))
}

func (*Templates) Contains(test, s interface{}) bool {
	return strings.Contains(fmt.Sprint(s), fmt.Sprint(test))
}

func (*Templates) ToS8(val byte) string {
	return fmt.Sprint(int8(val))
}
