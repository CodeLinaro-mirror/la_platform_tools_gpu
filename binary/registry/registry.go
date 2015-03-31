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
	"reflect"

	"android.googlesource.com/platform/tools/gpu/binary"
)

var (
	typeToID = map[reflect.Type]binary.ID{}
	idToType = map[binary.ID]reflect.Type{}
)

// Add a new type to the binary encoding system.
// The id should be a sha1 has of the types signature, such that
// no two types generate the same ID, and any change to a types name or
// fields causes it's signature to change.
func Add(id binary.ID, instance binary.Object) {
	t := reflect.TypeOf(instance)
	if oldId, found := typeToID[t]; found {
		panic(fmt.Errorf("Type %s as %x already has id %x", t, id, oldId))
	}
	typeToID[t] = id
	if oldType, found := idToType[id]; found {
		panic(fmt.Errorf("Id %x for %s already as type %s", id, t, oldType))
	}
	idToType[id] = t.Elem()
}

type unknownType struct {
	Object interface{}
}

func (e unknownType) Error() string {
	return fmt.Sprintf("Unknown type %T encountered in binary.TypeOf", e.Object)
}

// TypeOf returns the registered type ID for the obj.
// If there is no ID for this type of object, TypeOf returns a non-nil error.
func TypeOf(obj interface{}) (binary.ID, error) {
	id, idFound := typeToID[reflect.TypeOf(obj)]
	if !idFound {
		return binary.ID{}, unknownType{obj}
	}
	return id, nil
}

type unknownTypeID binary.ID

func (e unknownTypeID) Error() string {
	return fmt.Sprintf("Unknown type id %v encountered in binary.MakeObject", binary.ID(e))
}

// New creates a zero value instance for the given type ID.
// If this ID is not for a registered type, New returns a non-nil error.
func New(typeId binary.ID) (binary.Decodable, error) {
	if typeId == (binary.ID{}) {
		return nil, nil
	}
	t, idFound := idToType[typeId]
	if !idFound {
		return nil, unknownTypeID(typeId)
	}
	return reflect.New(t).Interface().(binary.Decodable), nil
}

// Nil returns a typed nil interface for the given type ID.
// This allows calling of methods that accept a nil pointer, eg Skip.
func Nil(typeId binary.ID) (binary.Decodable, error) {
	if typeId == (binary.ID{}) {
		return nil, nil
	}
	t, idFound := idToType[typeId]
	if !idFound {
		return nil, unknownTypeID(typeId)
	}
	return reflect.Zero(reflect.PtrTo(t)).Interface().(binary.Decodable), nil
}
