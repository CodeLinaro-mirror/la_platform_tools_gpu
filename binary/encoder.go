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
	"io"
	"math"
)

// Encoder provides methods for encoding values to an io.Writer.
type Encoder struct {
	writer  io.Writer
	objects map[interface{}]uint32
	tmp     [9]byte
}

// NewEncoder creates an Encoder that writes to the supplied stream.
func NewEncoder(writer io.Writer) *Encoder {
	return &Encoder{writer: writer, objects: map[interface{}]uint32{}}
}

// Write implements the io.Writer interface, delegating to the underlying writer.
func (e *Encoder) Write(p []byte) (int, error) {
	return e.writer.Write(p)
}

// WriteFull writes the data bytes in their entirety.
func (e *Encoder) WriteFull(data []byte) error {
	n, err := e.writer.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return io.ErrShortWrite
	}
	return nil
}

// Bool encodes a boolean value to the Encoder's io.Writer.
func (e *Encoder) Bool(v bool) error {
	if v {
		e.tmp[0] = 1
	} else {
		e.tmp[0] = 0
	}
	return e.WriteFull(e.tmp[:1])
}

// Int8 encodes a signed, 8 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int8(v int8) error {
	return e.Uint8(uint8(v))
}

// Uint8 encodes an unsigned, 8 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint8(v uint8) error {
	e.tmp[0] = v
	return e.WriteFull(e.tmp[:1])
}

// Int16 encodes a signed, 16 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int16(v int16) error {
	uv := uint16(v) << 1
	if v < 0 {
		uv = ^uv
	}
	return e.Uint16(uv)
}

// Uint16 encodes an unsigned, 16 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint16(v uint16) error {
	space := uint16(0x7f)
	tag := byte(0)
	for o := 8; true; o-- {
		if v <= space {
			e.tmp[o] = byte(v) | byte(tag)
			return e.WriteFull(e.tmp[o:])
		}
		e.tmp[o] = byte(v)
		v = v >> 8
		space >>= 1
		tag = (tag >> 1) | 0x80
	}
	panic("Cannot get here")
}

// Int32 encodes a signed, 32 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int32(v int32) error {
	uv := uint32(v) << 1
	if v < 0 {
		uv = ^uv
	}
	return e.Uint32(uv)
}

// Uint32 encodes an usigned, 32 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint32(v uint32) error {
	space := uint32(0x7f)
	tag := byte(0)
	for o := 8; true; o-- {
		if v <= space {
			e.tmp[o] = byte(v) | byte(tag)
			return e.WriteFull(e.tmp[o:])
		}
		e.tmp[o] = byte(v)
		v = v >> 8
		space >>= 1
		tag = (tag >> 1) | 0x80
	}
	panic("Cannot get here")
}

// Float32 encodes a 32 bit floating-point value to the Encoder's io.Writer.
func (e *Encoder) Float32(v float32) error {
	bits := math.Float32bits(v)
	shuffled := 0 |
		((bits & 0x000000ff) << 24) |
		((bits & 0x0000ff00) << 8) |
		((bits & 0x00ff0000) >> 8) |
		((bits & 0xff000000) >> 24)
	return e.Uint32(shuffled)
}

// Int64 encodes a signed, 64 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int64(v int64) error {
	uv := uint64(v) << 1
	if v < 0 {
		uv = ^uv
	}
	return e.Uint64(uv)
}

// Uint64 encodes an unsigned, 64 bit integer value to the Encoders's io.Writer.
func (e *Encoder) Uint64(v uint64) error {
	space := uint64(0x7f)
	tag := byte(0)
	for o := 8; true; o-- {
		if v <= space {
			e.tmp[o] = byte(v) | byte(tag)
			return e.WriteFull(e.tmp[o:])
		}
		e.tmp[o] = byte(v)
		v = v >> 8
		space >>= 1
		tag = (tag >> 1) | 0x80
	}
	panic("Cannot get here")
}

// Float64 encodes a 64 bit floating-point value to the Encoder's io.Writer.
func (e *Encoder) Float64(v float64) error {
	bits := math.Float64bits(v)
	shuffled := 0 |
		((bits & 0x00000000000000ff) << 56) |
		((bits & 0x000000000000ff00) << 40) |
		((bits & 0x0000000000ff0000) << 24) |
		((bits & 0x00000000ff000000) << 8) |
		((bits & 0x000000ff00000000) >> 8) |
		((bits & 0x0000ff0000000000) >> 24) |
		((bits & 0x00ff000000000000) >> 40) |
		((bits & 0xff00000000000000) >> 56)
	return e.Uint64(shuffled)
}

// String encodes a string to the Encoder's io.Writer.
func (e *Encoder) String(v string) error {
	if err := e.Uint32(uint32(len(v))); err != nil {
		return err
	}
	return e.WriteFull([]byte(v))
}

// Object encodes an Encodable to the Encoder's io.Writer. If Object is called repeatedly with the
// same argument (i.e. the argument has identical dynamic types and equal dynamic values), then the
// argument will only be encoded with the first call, and later encodings will reference the first
// encoding.
// The type of obj must have been previously registered with binary.Register.
func (e *Encoder) Object(obj Encodable) error {
	if obj == nil {
		return e.Uint32(0)
	}

	key, alreadyEncoded := e.objects[obj]
	if alreadyEncoded {
		return e.Uint32(key)
	}

	id, err := TypeOf(obj)
	if err != nil {
		return err
	}

	key = uint32(len(e.objects) + 1)
	e.objects[obj] = key
	if err := e.Uint32(key); err != nil {
		return err
	}
	if err := id.Encode(e); err != nil {
		return err
	}
	return obj.Encode(e)
}
