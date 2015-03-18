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
	// Object encodes an Encodable. The type of obj must have
	// been previously registered with binary.registry.Add.
	Object(obj Encodable) error
}

// Encodable is the interface for an object that can be written to an Encoder.
type Encodable interface {
	// Encode the object's data to the Encoder.
	// The implementation must be symmetrical to Decode.
	Encode(e Encoder) error
}
