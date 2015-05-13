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

package schema

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/service"
)

// Array is a schema-typed Array value.
type Array struct {
	Type     *service.ArrayInfo // The array type info.
	Elements []ArrayElement     // The array elements.
}

// StaticArray is a schema-typed StaticArray value.
type StaticArray struct {
	Type     *service.StaticArrayInfo // The static array type info.
	Elements []ArrayElement           // The static array elements.
}

// ArrayElement is a single element held by an Array.
type ArrayElement interface{}

func (a Array) String() string {
	parts := make([]string, len(a.Elements))
	for i, e := range a.Elements {
		parts[i] = fmt.Sprintf("%v", e)
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
}

func (a StaticArray) String() string {
	parts := make([]string, len(a.Elements))
	for i, e := range a.Elements {
		parts[i] = fmt.Sprintf("%v", e)
	}
	return fmt.Sprintf("[%s]", strings.Join(parts, ", "))
}
