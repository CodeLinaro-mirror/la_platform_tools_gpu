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

package binary

import (
	"reflect"
)

// Type is a single Object type definition used by the TypeNamespace.
type Type struct {
	ID  ObjectTypeID  // The type identifier
	New func() Object // A function that will return a default-initialized instance of the object type
}

// TypeNamespace is used by Encoder and Decoder to encode and decode Objects.
// Use NewTypeNamespace to build a new TypeNamespace with the specified Type list.
// Call Encoder.WithNamespace and Decoder.WithNamespace to create Encoders and Decoders that can
// encode and decode objects that have types declared in the TypeNamespace.
type TypeNamespace struct {
	ptrTypeToID map[reflect.Type]ObjectTypeID
	idToNew     map[ObjectTypeID]func() Object
}

func (n TypeNamespace) new(id ObjectTypeID) Object {
	if new, found := n.idToNew[ObjectTypeID(id)]; found {
		return new()
	} else {
		return nil
	}
}

func (n TypeNamespace) idOf(obj Object) (ObjectTypeID, bool) {
	ty := reflect.TypeOf(obj)
	id, found := n.ptrTypeToID[ty]
	return id, found
}

// NewTypeNamespace constructs a new TypeNamespace containing the specified list of Types.
func NewTypeNamespace(types ...Type) TypeNamespace {
	return TypeNamespace{}.Merge(types...)
}

// Merge creates and returns a new TypeNamespace that combines the specified types to types in the
// existing TypeNamespace.
func (n TypeNamespace) Merge(types ...Type) TypeNamespace {
	c := len(n.idToNew) + len(types)
	ptrTypeToID := make(map[reflect.Type]ObjectTypeID, c)
	idToNew := make(map[ObjectTypeID]func() Object, c)

	// Copy existing
	for ty, id := range n.ptrTypeToID {
		ptrTypeToID[ty] = id
	}

	// Copy existing
	for id, new := range n.idToNew {
		idToNew[id] = new
	}

	// Add new
	for _, e := range types {
		id := e.ID
		instance := e.New()
		ptrTy := reflect.TypeOf(instance)
		ptrTypeToID[ptrTy] = id
		idToNew[id] = e.New
	}

	return TypeNamespace{ptrTypeToID, idToNew}
}
