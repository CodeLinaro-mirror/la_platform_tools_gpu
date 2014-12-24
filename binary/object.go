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

// TypeID is a unique type identifier used to identify a type.
// It is expected these will be SHA1 hashes of a types signature, such that
// no two types generate the same TypeId, and any change to a types name or
// fields causes it's signature to change.
type TypeID [20]byte

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
	typeToID = map[reflect.Type]TypeID{}
	idToType = map[TypeID]reflect.Type{}
)

// Register adds a new type to the binary encoding system.
func Register(id TypeID, instance Object) {
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
