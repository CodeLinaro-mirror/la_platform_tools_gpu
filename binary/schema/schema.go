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

// Package schema implements rtti for the binary system.
package schema

import (
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
)

// binary: Schema = false
// binary: java.source = base/rpclib/src/main/java
// binary: java.package = com.android.tools.rpclib.schema
// binary: java.indent = "    "
// binary: java.member_prefix = m

// Type represents the common interface to all type objects in the schema.
type Type interface {
	String() string
	Encode(e binary.Encoder, value interface{})
	Decode(d binary.Decoder) interface{}
	Typename() string
	Basename() string
}

type schema interface {
	Schema() *Entity
}

// Returns the schema class for a binary class, if it has one.
func Of(class binary.Class) *Entity {
	if s, ok := class.(*Entity); ok {
		return s
	}
	if s, ok := class.(schema); ok {
		return s.Schema()
	}
	return nil
}

// Lookup looks up a Class by the given type id.
// If there is no match, it will return nil.
func Lookup(id binary.ID) *Entity {
	return Of(registry.Global.Lookup(id))
}

func printTag(w io.Writer, t Type) {
	switch t := t.(type) {
	case *Primitive:
		fmt.Fprint(w, t.Method)
	case *Struct:
		fmt.Fprint(w, "$")
	case *Pointer:
		fmt.Fprint(w, "*")
		printTag(w, t.Type)
	case *Interface:
		fmt.Fprint(w, "?", t)
	case *Variant:
		fmt.Fprint(w, "&", t)
	case *Any:
		fmt.Fprint(w, "~", t)
	case *Slice:
		fmt.Fprint(w, "[]")
		printTag(w, t.ValueType)
	case *Array:
		fmt.Fprint(w, "[", t.Size, "]")
		printTag(w, t.ValueType)
	case *Map:
		fmt.Fprint(w, "map[")
		printTag(w, t.KeyType)
		fmt.Fprint(w, "]")
		printTag(w, t.ValueType)
	default:
		panic(fmt.Errorf("Unknown type %T generating signature", t))
	}
}
