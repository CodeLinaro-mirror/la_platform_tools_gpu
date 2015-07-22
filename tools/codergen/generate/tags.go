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

// Package generate has support for generating encode and decode methods
// for the binary package automatically.
package generate

import (
	"fmt"
	"reflect"
	"strconv"
)

// Tags wraps a go field tag string, and adds method to extract it's values.
type Tags string

// Get returns a single value by name.
func (t Tags) Get(name string) string {
	return reflect.StructTag(t).Get(name)
}

// Flag returns a boolean value by name.
func (t Tags) Flag(name string) bool {
	v := reflect.StructTag(t).Get(name)
	if len(v) == 0 {
		return false
	}
	b, err := strconv.ParseBool(v)
	if err != nil {
		panic(fmt.Errorf("Malformed tag %q in %q: %v", name, t, err))
	}
	return b
}
