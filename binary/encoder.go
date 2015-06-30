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
	ID(ID) error
	// Object encodes an Object with no type preamble and no sharing.
	Value(obj Object) error
	// Variant encodes a POD value or an Object with no sharing.
	// If v is an Object then its type must have been previously registered with
	// binary.registry.Add.
	// If v is not a POD type nor a Object then ErrNotEncodable is returned.
	Variant(v interface{}) error
	// Object encodes a POD value or an Object, optionally encoding objects only
	// on the first time it sees them.
	// If v is an Object then its type must have been previously registered with
	// binary.registry.Add.
	// If v is not a POD type nor a Object then ErrNotEncodable is returned.
	Object(v interface{}) error
}
