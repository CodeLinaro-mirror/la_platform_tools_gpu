// Copyright (C) 2015 The Android Open Source Project
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

	"android.googlesource.com/platform/tools/gpu/api/semantic"
)

// Unpack throws away type information to work around template system limitations
// When you have a value of an interface type that carries methods, it fails to
// introspect the concrete type for it's members, so the template can't see them.
// The result of Upack no longer has a type, so the concrete type members become
// visible.
func (*Functions) Unpack(v interface{}) interface{} { return v }

// GetAnnotation finds and returns the annotation on ty with the specified name.
// If the annotation cannot be found, or ty does not support annotations then
// GetAnnotation returns nil.
func (*Functions) GetAnnotation(ty interface{}, name string) *semantic.Annotation {
	a, ok := ty.(semantic.Annotated)
	if !ok {
		return nil
	}
	return a.GetAnnotation(name)
}

// Underlying returns the underlying type for ty by recursively traversing the
// pseudonym chain until reaching and returning the first non-pseudoym type.
// If ty is not a pseudonym then it is simply returned.
func (f *Functions) Underlying(ty semantic.Type) semantic.Type {
	for {
		if pseudo, ok := ty.(*semantic.Pseudonym); ok {
			ty = pseudo.To
		} else {
			return ty
		}
	}
}

// AllCommands returns a list of all cmd entries for a given API, regardless
// of whether they are free functions, class methods or pseudonym methods.
func (f *Functions) AllCommands(api interface{}) ([]interface{}, error) {
	switch api := api.(type) {
	case *semantic.API:
		var commands []interface{}
		for _, function := range api.Functions {
			commands = append(commands, function)
		}
		for _, class := range api.Classes {
			for _, method := range class.Methods {
				commands = append(commands, method)
			}
		}
		for _, pseudonym := range api.Pseudonyms {
			for _, method := range pseudonym.Methods {
				commands = append(commands, method)
			}
		}
		return commands, nil
	default:
		return nil, fmt.Errorf("first argument must be of type *semantic.API, was %T", api)
	}
}
