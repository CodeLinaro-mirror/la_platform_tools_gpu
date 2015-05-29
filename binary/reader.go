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

import "fmt"

// Reader provides methods for decoding values.
type Reader interface {
	// Data reads the data bytes in their entirety.
	Data([]byte) error
	// Skip jumps past count bytes.
	Skip(count uint32) error
	// Bool decodes and returns a boolean value from the Reader.
	Bool() (bool, error)
	// Int8 decodes and returns a signed, 8 bit integer value from the Reader.
	Int8() (int8, error)
	// Uint8 decodes and returns an unsigned, 8 bit integer value from the Reader.
	Uint8() (uint8, error)
	// Int16 decodes and returns a signed, 16 bit integer value from the Reader.
	Int16() (int16, error)
	// Uint16 decodes and returns an unsigned, 16 bit integer value from the Reader.
	Uint16() (uint16, error)
	// Int32 decodes and returns a signed, 32 bit integer value from the Reader.
	Int32() (int32, error)
	// Uint32 decodes and returns an unsigned, 32 bit integer value from the Reader.
	Uint32() (uint32, error)
	// Float32 decodes and returns a 32 bit floating-point value from the Reader.
	Float32() (float32, error)
	// Int64 decodes and returns a signed, 64 bit integer value from the Reader.
	Int64() (int64, error)
	// Uint64 decodes and returns an unsigned, 64 bit integer value from the Reader.
	Uint64() (uint64, error)
	// Float64 decodes and returns a 64 bit floating-point value from the Reader.
	Float64() (float64, error)
	// String decodes and returns a string from the Reader.
	String() (string, error)
	// SkipString skips over a single string from the Reader.
	SkipString() error
}

// ReadUint reads an unsigned integer of either 8, 16, 32 or 64 bits from r,
// returning the result as a uint64.
func ReadUint(r Reader, bits int) (uint64, error) {
	switch bits {
	case 8:
		v, err := r.Uint8()
		return uint64(v), err
	case 16:
		v, err := r.Uint16()
		return uint64(v), err
	case 32:
		v, err := r.Uint32()
		return uint64(v), err
	case 64:
		v, err := r.Uint64()
		return v, err
	default:
		return 0, fmt.Errorf("Unsupported integer bit count %v", bits)
	}
}

// ReadInt reads a signed integer of either 8, 16, 32 or 64 bits from r,
// returning the result as a int64.
func ReadInt(r Reader, bits int) (int64, error) {
	switch bits {
	case 8:
		v, err := r.Int8()
		return int64(v), err
	case 16:
		v, err := r.Int16()
		return int64(v), err
	case 32:
		v, err := r.Int32()
		return int64(v), err
	case 64:
		v, err := r.Int64()
		return v, err
	default:
		return 0, fmt.Errorf("Unsupported integer bit count %v", bits)
	}
}
