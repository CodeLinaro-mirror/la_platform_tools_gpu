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
	// ID decodes a binary.ID from the stream.
	ID() (ID, error)
	// SkipID skips over a binary.ID in the stream.
	SkipID() error
	// Value decodes a Decodable from the stream.
	Value(Decodable) error
	// SkipValue must skip the same data that a call to Value would read.
	// The value may be a typed nil.
	SkipValue(Decodable) error
	// Variant decodes and returns a Decodable from the stream. The type id in the
	// stream must have been previously registered with binary.registry.Add.
	Variant() (interface{}, error)
	// SkipVariant must skip the same data that a call to Variant would read.
	SkipVariant() error
	// Object decodes and returns a Decodable from the stream. Object instances
	// that were encoded multiple times may be decoded and returned as a shared,
	// single instance. The type id in the stream must have been previously
	// registered with binary.registry.Add.
	Object() (interface{}, error)
	// SkipObject must skip the same data that a call to Object would read.
	SkipObject() error
}

// Decodable is the interface for an object that can be read from a Decoder.
type Decodable interface {
	// Decode the object's data from the Decoder.
	// The implementation must be symmetrical to Encode.
	Decode(Decoder) error
	// Skip the object's data from the Decoder.
	// This must skip the same data that Decode would have read, and must be safe
	// to call with a nil receiver.
	Skip(Decoder) error
}
