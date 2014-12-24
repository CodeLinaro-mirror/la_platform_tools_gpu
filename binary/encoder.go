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

import (
	"fmt"
	"io"
	"math"
	"reflect"
)

// Encoder provides methods for encoding values to an io.Writer.
type Encoder struct {
	writer  io.Writer
	objects map[interface{}]uint16
	tmp     [8]byte
}

// NewEncoder creates an Encoder that writes to the supplied stream.
func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{writer: writer, objects: map[interface{}]uint16{}}
}

func (e *Encoder) write(data []byte) error {
	n, err := e.writer.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return io.ErrShortWrite
	}
	return nil
}

// Write implements the io.Writer interface, delegating to the underlying writer.
func (e *Encoder) Write(p []byte) (int, error) {
	return e.writer.Write(p)
}

// TypeID encodes a type id to the Encoder's io.Writer.
func (e *Encoder) TypeID(v TypeID) error {
	return e.write(v[:])
}

// Bool encodes a boolean value to the Encoder's io.Writer.
func (e *Encoder) Bool(v bool) error {
	if v {
		e.tmp[0] = 1
	} else {
		e.tmp[0] = 0
	}
	return e.write(e.tmp[:1])
}

// Int8 encodes a signed, 8 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int8(v int8) error {
	return e.Uint8(uint8(v))
}

// Uint8 encodes an unsigned, 8 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint8(v uint8) error {
	e.tmp[0] = v
	return e.write(e.tmp[:1])
}

// Int16 encodes a signed, 16 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int16(v int16) error {
	return e.Uint16(uint16(v))
}

// Uint16 encodes an unsigned, 16 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint16(v uint16) error {
	e.tmp[0] = byte(v)
	e.tmp[1] = byte(v >> 8)
	return e.write(e.tmp[:2])
}

// Int32 encodes a signed, 32 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int32(v int32) error {
	return e.Uint32(uint32(v))
}

// Uint32 encodes an usigned, 32 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint32(v uint32) error {
	e.tmp[0] = byte(v)
	e.tmp[1] = byte(v >> 8)
	e.tmp[2] = byte(v >> 16)
	e.tmp[3] = byte(v >> 24)
	return e.write(e.tmp[:4])
}

// Float32 encodes a 32 bit floating-point value to the Encoder's io.Writer.
func (e *Encoder) Float32(v float32) error {
	return e.Uint32(math.Float32bits(v))
}

// Int64 encodes a signed, 64 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int64(v int64) error {
	return e.Uint64(uint64(v))
}

// Uint64 encodes an unsigned, 64 bit integer value to the Encoders's io.Writer.
func (e *Encoder) Uint64(v uint64) error {
	e.tmp[0] = byte(v)
	e.tmp[1] = byte(v >> 8)
	e.tmp[2] = byte(v >> 16)
	e.tmp[3] = byte(v >> 24)
	e.tmp[4] = byte(v >> 32)
	e.tmp[5] = byte(v >> 40)
	e.tmp[6] = byte(v >> 48)
	e.tmp[7] = byte(v >> 56)
	return e.write(e.tmp[:8])
}

// Float64 encodes a 64 bit floating-point value to the Encoder's io.Writer.
func (e *Encoder) Float64(v float64) error {
	return e.Uint64(math.Float64bits(v))
}

// String encodes a string (pascal-style) to the Encoder's io.Writer.
func (e *Encoder) String(v string) error {
	if err := e.Uint32(uint32(len(v))); err != nil {
		return err
	}
	return e.write([]byte(v))
}

// String encodes a string (c-style) to the Encoder's io.Writer.
func (e *Encoder) CString(v string) error {
	if err := e.write([]byte(v)); err != nil {
		return err
	}
	return e.Uint8(0)
}

// Data encodes a sequence of bytes to the Encoder's io.Writer.
func (e *Encoder) Data(data []byte) error {
	if err := e.Uint32(uint32(len(data))); err != nil {
		return err
	}
	return e.write(data)
}

type unknownType struct {
	Object Encodable
}

func (e unknownType) Error() string {
	return fmt.Sprintf("Unknown type %T encountered in binary.Encoder", e.Object)
}

// Object encodes an Encodable to the Encoder's io.Writer. If Object is called repeatedly with the
// same argument (i.e. the argument has identical dynamic types and equal dynamic values), then the
// argument will only be encoded with the first call, and later encodings will reference the first
// encoding.
// The type of obj must have been previously registered with binary.Register.
func (e *Encoder) Object(obj Encodable) error {
	if obj == nil {
		return e.Uint16(objectNil)
	}

	key, alreadyEncoded := e.objects[obj]
	if alreadyEncoded {
		return e.Uint16(key)
	}

	id, idFound := typeToID[reflect.TypeOf(obj)]
	if !idFound {
		return unknownType{obj}
	}

	key = uint16(len(e.objects))
	e.objects[obj] = key
	e.Uint16(key)
	e.TypeID(id)
	return obj.Encode(e)
}
