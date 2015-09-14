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
	// If there is an error reading any input, all further reading returns the
	// zero value of the type read. Error() returns the error which stopped
	// reading from the stream. If reading has not stopped it returns nil.
	Error() error
	// Set the error state and stop reading from the stream.
	SetError(error) error
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

// ReadBool decodes and returns a boolean value from the Reader.
func ReadBool(r Reader) bool {
	v, _ := r.Bool()
	return v
}

// ReadInt8 decodes and returns a signed, 8 bit integer value from the Reader.
func ReadInt8(r Reader) int8 {
	v, _ := r.Int8()
	return v
}

// ReadUint8 decodes and returns an unsigned, 8 bit integer value from the Reader.
func ReadUint8(r Reader) uint8 {
	v, _ := r.Uint8()
	return v
}

// ReadInt16 decodes and returns a signed, 16 bit integer value from the Reader.
func ReadInt16(r Reader) int16 {
	v, _ := r.Int16()
	return v
}

// ReadUint16 decodes and returns an unsigned, 16 bit integer value from the Reader.
func ReadUint16(r Reader) uint16 {
	v, _ := r.Uint16()
	return v
}

// ReadInt32 decodes and returns a signed, 32 bit integer value from the Reader.
func ReadInt32(r Reader) int32 {
	v, _ := r.Int32()
	return v
}

// ReadUint32 decodes and returns an unsigned, 32 bit integer value from the Reader.
func ReadUint32(r Reader) uint32 {
	v, _ := r.Uint32()
	return v
}

// ReadFloat32 decodes and returns a 32 bit floating-point value from the Reader.
func ReadFloat32(r Reader) float32 {
	v, _ := r.Float32()
	return v
}

// ReadInt64 decodes and returns a signed, 64 bit integer value from the Reader.
func ReadInt64(r Reader) int64 {
	v, _ := r.Int64()
	return v
}

// ReadUint64 decodes and returns an unsigned, 64 bit integer value from the Reader.
func ReadUint64(r Reader) uint64 {
	v, _ := r.Uint64()
	return v
}

// ReadFloat64 decodes and returns a 64 bit floating-point value from the Reader.
func ReadFloat64(r Reader) float64 {
	v, _ := r.Float64()
	return v
}

// ReadString decodes and returns a string from the Reader.
func ReadString(r Reader) string {
	v, _ := r.String()
	return v
}
