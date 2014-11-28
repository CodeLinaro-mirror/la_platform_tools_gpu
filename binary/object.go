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

// Used as an object key to define a nil pointer
const objectNil uint16 = ^uint16(0)

// ObjectTypeID is a unique type identifier used by TypeNamespace to construct Objects of the
// correct type. The valid ranges for ObjectTypeID are 0 to 0xfffe inclusive.
type ObjectTypeID uint16

// Object is the interface used by types that can be encoded and decoded using Encoder.Object and
// Decoder.Object, respectively.
type Object interface {
	// Encode the object's data to the Encoder.
	// The implementation must be symmetrical to Decode.
	Encode(e *Encoder) error
	// Decode the object's data from the Decoder.
	// The implementation must be symmetrical to Encode.
	Decode(d *Decoder) error
}
