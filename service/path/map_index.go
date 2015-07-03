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

package path

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// MapIndex is a path that refers to a single value in a map.
type MapIndex struct {
	binary.Generate
	Map Value       // The path to the map containing the value.
	Key interface{} // The key to the value in the map.
}

// Path implements the Path interface.
func (n *MapIndex) Path() string {
	return fmt.Sprintf("%v[%v]", n.Map, n.Key)
}

// Field implements the Value interface.
func (n *MapIndex) Field(name string) Value {
	return &Field{Struct: n, Name: name}
}

// ArrayIndex implements the Value interface.
func (n *MapIndex) ArrayIndex(index uint64) Value {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex implements the Value interface.
func (n *MapIndex) MapIndex(key interface{}) Value {
	return &MapIndex{Map: n, Key: key}
}
