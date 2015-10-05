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
	fallbacks []*Namespace
	classes   map[binary.Signature]binary.Class
}

var (
	// Global is the default global Namespace object.
	Global = NewNamespace()
)

// NewNamespace creates a new namespace layered on top of the specified fallback.
func NewNamespace(fallbacks ...*Namespace) *Namespace {
	return &Namespace{
		fallbacks: fallbacks,
		classes:   map[binary.Signature]binary.Class{},
	}
}

// Add a new type to the Namespace.
func (n *Namespace) Add(class binary.Class) {
	if class == nil {
		panic(fmt.Errorf("Attempt to add nil class to registry"))
	}
	entity := class.Schema()
	if entity == nil {
		panic(fmt.Errorf("Class %T has no schema", class))
	}
	signature := class.Schema().Signature()
	if _, found := n.classes[signature]; found {
		panic(fmt.Errorf("Class for %s already present", signature))
	}
	n.classes[signature] = class
}

// AddFallbacks appends new Namespaces to the fallback list of this Namespace.
func (n *Namespace) AddFallbacks(fallbacks ...*Namespace) {
	n.fallbacks = append(n.fallbacks, fallbacks...)
}

// Lookup looks up a Class by the given type id in the Namespace.
// If there is no match, it will return nil.
func (n *Namespace) Lookup(signature binary.Signature) binary.Class {
	if class, found := n.classes[signature]; found {
		return class
	}
	for _, f := range n.fallbacks {
		if class := f.Lookup(signature); class != nil {
			return class
		}
	}
	return nil
}

// Count returns the number of entries reachable through this namespace.
// Because it sums the counts of the namespaces it depends on, this may be
// more than the number of unique keys.
func (n *Namespace) Count() int {
	size := len(n.classes)
	for _, f := range n.fallbacks {
		size += f.Count()
	}
	return size
}

// Visit invokes the visitor for every class object reachable through this
// namespace.
// The visitor maybe be called with the same id more than once if it is present
// in multiple namespaces.
func (n Namespace) Visit(visitor func(binary.Class)) {
	for _, c := range n.classes {
		visitor(c)
	}
	for _, f := range n.fallbacks {
		f.Visit(visitor)
	}
}
