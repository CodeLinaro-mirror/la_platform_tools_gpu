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

// Writer provides methods for encoding values.
type Writer interface {
	// Data writes the data bytes in their entirety.
	Data([]byte) error
	// Bool encodes a boolean value to the Writer.
	Bool(bool) error
	// Int8 encodes a signed, 8 bit integer value to the Writer.
	Int8(int8) error
	// Uint8 encodes an unsigned, 8 bit integer value to the Writer.
	Uint8(uint8) error
	// Int16 encodes a signed, 16 bit integer value to the Writer.
	Int16(int16) error
	// Uint16 encodes an unsigned, 16 bit integer value to the Writer.
	Uint16(uint16) error
	// Int32 encodes a signed, 32 bit integer value to the Writer.
	Int32(int32) error
	// Uint32 encodes an usigned, 32 bit integer value to the Writer.
	Uint32(uint32) error
	// Float32 encodes a 32 bit floating-point value to the Writer.
	Float32(float32) error
	// Int64 encodes a signed, 64 bit integer value to the Writer.
	Int64(int64) error
	// Uint64 encodes an unsigned, 64 bit integer value to the Encoders's io.Writer.
	Uint64(uint64) error
	// Float64 encodes a 64 bit floating-point value to the Writer.
	Float64(float64) error
	// String encodes a string to the Writer.
	String(string) error
	// If there is an error writing any output, all further writing becomes
	// a no-op. Error() returns the error which stopped writing to the stream.
	// If writing has not stopped it returns nil.
	Error() error
	// Set the error state and stop writing to the stream.
	SetError(error) error
}

// WriteUint writes the unsigned integer v of either 8, 16, 32 or 64 bits to w.
func WriteUint(w Writer, bits int, v uint64) error {
	switch bits {
	case 8:
		return w.Uint8(uint8(v))
	case 16:
		return w.Uint16(uint16(v))
	case 32:
		return w.Uint32(uint32(v))
	case 64:
		return w.Uint64(uint64(v))
	default:
		return fmt.Errorf("Unsupported integer bit count %v", bits)
	}
}

// WriteInt writes the signed integer v of either 8, 16, 32 or 64 bits to w.
func WriteInt(w Writer, bits int, v int64) error {
	switch bits {
	case 8:
		return w.Int8(int8(v))
	case 16:
		return w.Int16(int16(v))
	case 32:
		return w.Int32(int32(v))
	case 64:
		return w.Int64(int64(v))
	default:
		return fmt.Errorf("Unsupported integer bit count %v", bits)
	}
}
