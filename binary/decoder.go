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

// Decoder extends Reader with additional methods for decoding objects.
type Decoder interface {
	Reader
	// Entity supports reading a binary.Entity from the stream.
	// See Entity for details of what is included in the compact form.
	Entity(compact bool) *Entity
	// Value decodes an Object from the stream.
	Value(Object)
	// Struct decodes an Object from the stream. The type identifier is specified.
	Struct(t Type, obj Object) error
	// Pop and subtype ID from the decoder stack.
	PopType() (Type, error)
	// Variant decodes and returns an Object from the stream. The Class in the
	// stream must have been previously registered with binary.registry.Add.
	Variant() Object
	// Object decodes and returns an Object from the stream. Object instances
	// that were encoded multiple times may be decoded and returned as a shared,
	// single instance. The Class in the stream must have been previously
	// registered with binary.registry.Add.
	Object() Object
	// Lookup the upgrade decoder for decoding this type of entity.
	Lookup(*Entity) UpgradeDecoder
	// Decode a collection count from the stream.
	Count() uint32
}
