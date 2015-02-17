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
	"fmt"
	"reflect"
)

// Used as an object key to define a nil pointer
const objectNil uint16 = ^uint16(0)

type Encodable interface {
	// Encode the object's data to the Encoder.
	// The implementation must be symmetrical to Decode.
	Encode(e *Encoder) error
}

type Decodable interface {
	// Decode the object's data from the Decoder.
	// The implementation must be symmetrical to Encode.
	Decode(d *Decoder) error
}

type Object interface {
	Encodable
	Decodable
}

var (
	typeToID = map[reflect.Type]ID{}
	idToType = map[ID]reflect.Type{}
)

// Register adds a new type to the binary encoding system.
// The id should be a sha1 has of the types signature, such that
// no two types generate the same ID, and any change to a types name or
// fields causes it's signature to change.
func Register(id ID, instance Object) {
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
	Object Encodable
}

func (e unknownType) Error() string {
	return fmt.Sprintf("Unknown type %T encountered in binary.TypeOf", e.Object)
}

// Given an encodable object return the ID for that type of object.
// If there is no ID for this type of object, return a non-nil error.
func TypeOf(obj Encodable) (ID, error) {
	id, idFound := typeToID[reflect.TypeOf(obj)]
	if !idFound {
		return ID{}, unknownType{obj}
	}
	return id, nil
}

type unknownTypeID ID

func (e unknownTypeID) Error() string {
	return fmt.Sprintf("Unknown type id %v encountered in binary.MakeObject", ID(e))
}

// Given an ID return a zero value instance of the object type.
// If this ID is not for a registered type, return a non-nil error.
func MakeObject(typeId ID) (Decodable, error) {
	t, idFound := idToType[typeId]
	if !idFound {
		return nil, unknownTypeID(typeId)
	}
	return reflect.New(t).Interface().(Decodable), nil
}
