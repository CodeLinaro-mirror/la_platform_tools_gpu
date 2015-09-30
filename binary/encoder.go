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

// Encoder extends Writer with additional methods for encoding objects.
type Encoder interface {
	Writer
	// ID writes a binary.ID to the stream.
	ID(ID)
	// Entity writes a binary.Entity to the stream.
	// See Entity for details of what is included in the compact form.
	Entity(e *Entity, compact bool)
	// Object encodes an Object with no type preamble and no sharing.
	Value(obj Object)
	// Variant encodes an Object with no sharing. The type of obj must have
	// been previously registered with binary.registry.Add.
	Variant(obj Object)
	// Object encodes an Object, optionally encoding objects only on the first
	// time it sees them. The type of obj must have been previously registered
	// with binary.registry.Add.
	Object(obj Object)
}
