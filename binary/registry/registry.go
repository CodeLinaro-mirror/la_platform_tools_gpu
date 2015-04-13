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

// Namespace represents a mapping of type identifiers to their Class.
type Namespace struct {
	parent  *Namespace
	classes map[binary.ID]binary.Class
}

var (
	// Global is the default global Namespace object.
	Global = NewNamespace(nil)
)

// NewNamespace creates a new namespace layered on top of the specified parent.
func NewNamespace(parent *Namespace) *Namespace {
	return &Namespace{
		parent:  parent,
		classes: map[binary.ID]binary.Class{},
	}
}

// Add a new type to the global Namespace.
func Add(class binary.Class) {
	Global.Add(class)
}

// Add a new type to the Namespace.
func (n Namespace) Add(class binary.Class) {
	id := class.ID()
	if old, found := n.classes[id]; found {
		panic(fmt.Errorf("Id %x for %s already as type %s", id, class, old))
	}
	n.classes[id] = class
}

// Lookup looks up a Class by the given type id.
// If there is no match, it will return nil.
func Lookup(id binary.ID) binary.Class {
	return Global.Lookup(id)
}

// Lookup looks up a Class by the given type id in the Namespace.
// If there is no match, it will return nil.
func (n Namespace) Lookup(id binary.ID) binary.Class {
	if class, found := n.classes[id]; found {
		return class
	}
	if n.parent != nil {
		return n.parent.Lookup(id)
	}
	return nil
}
