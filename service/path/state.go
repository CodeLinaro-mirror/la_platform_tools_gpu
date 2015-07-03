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

// State is a path that refers to the driver state immediately after an atom.
type State struct {
	binary.Generate
	After *Atom // The path to the atom the state immediately follows.
}

// Path implements the Path interface.
func (n *State) Path() string {
	return fmt.Sprintf("State(After: %v)", n.After)
}

// Field implements the Value interface.
func (n *State) Field(name string) Value {
	return &Field{Struct: n, Name: name}
}

// ArrayIndex implements the Value interface.
func (n *State) ArrayIndex(index uint64) Value {
	return &ArrayIndex{Array: n, Index: index}
}

// MapIndex implements the Value interface.
func (n *State) MapIndex(key interface{}) Value {
	return &MapIndex{Map: n, Key: key}
}
