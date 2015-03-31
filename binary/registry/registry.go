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

package registry

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

var (
	ids = map[binary.ID]binary.Class{
		binary.NilClass.ID(): binary.NilClass,
	}
)

// Add a new type to the binary encoding system.
func Add(class binary.Class) {
	id := class.ID()
	if old, found := ids[id]; found {
		panic(fmt.Errorf("Id %x for %s already as type %s", id, class, old))
	}
	ids[id] = class
}

// Lookup looks up a Class by the given type id.
// If there is no match, it will return nil.
func Lookup(id binary.ID) binary.Class {
	return ids[id]
}
