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

package protocol

import (
	"encoding/binary"
	"io"
	"math"
)

// Encoder provides methods for encoding values to an io.Writer.
type Encoder struct {
	writer    io.Writer
	tmp       [8]byte
	byteOrder binary.ByteOrder
}

// NewEncoder creates an Encoder that writes to the supplied stream, with the
// specified endianness.
func NewEncoder(writer io.Writer, byteOrder binary.ByteOrder) *Encoder {
	return &Encoder{writer: writer, byteOrder: byteOrder}
}

// Write implements the io.Writer interface, delegating to the underlying writer.
func (e *Encoder) Write(p []byte) (int, error) {
	return e.writer.Write(p)
}

// Bool encodes a boolean value to the Encoder's io.Writer.
func (e *Encoder) Bool(v bool) error {
	if v {
		e.tmp[0] = 1
	} else {
		e.tmp[0] = 0
	}
	_, err := e.writer.Write(e.tmp[:1])
	return err
}

// Int8 encodes a signed, 8 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int8(v int8) error {
	e.tmp[0] = uint8(v)
	_, err := e.writer.Write(e.tmp[:1])
	return err
}

// Uint8 encodes an unsigned, 8 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint8(v uint8) error {
	e.tmp[0] = v
	_, err := e.writer.Write(e.tmp[:1])
	return err
}

// Int16 encodes a signed, 16 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int16(v int16) error {
	e.byteOrder.PutUint16(e.tmp[:], uint16(v))
	_, err := e.writer.Write(e.tmp[:2])
	return err
}

// Uint16 encodes an unsigned, 16 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint16(v uint16) error {
	e.byteOrder.PutUint16(e.tmp[:], v)
	_, err := e.writer.Write(e.tmp[:2])
	return err
}

// Int32 encodes a signed, 32 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int32(v int32) error {
	e.byteOrder.PutUint32(e.tmp[:], uint32(v))
	_, err := e.writer.Write(e.tmp[:4])
	return err
}

// Uint32 encodes an usigned, 32 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Uint32(v uint32) error {
	e.byteOrder.PutUint32(e.tmp[:], v)
	_, err := e.writer.Write(e.tmp[:4])
	return err
}

// Float32 encodes a 32 bit floating-point value to the Encoder's io.Writer.
func (e *Encoder) Float32(v float32) error {
	e.byteOrder.PutUint32(e.tmp[:], math.Float32bits(v))
	_, err := e.writer.Write(e.tmp[:4])
	return err
}

// Int64 encodes a signed, 64 bit integer value to the Encoder's io.Writer.
func (e *Encoder) Int64(v int64) error {
	e.byteOrder.PutUint64(e.tmp[:], uint64(v))
	_, err := e.writer.Write(e.tmp[:8])
	return err
}

// Uint64 encodes an unsigned, 64 bit integer value to the Encoders's io.Writer.
func (e *Encoder) Uint64(v uint64) error {
	e.byteOrder.PutUint64(e.tmp[:], v)
	_, err := e.writer.Write(e.tmp[:8])
	return err
}

// Float64 encodes a 64 bit floating-point value to the Encoder's io.Writer.
func (e *Encoder) Float64(v float64) error {
	e.byteOrder.PutUint64(e.tmp[:], math.Float64bits(v))
	_, err := e.writer.Write(e.tmp[:8])
	return err
}

// String encodes a string (pascal-style) to the Encoder's io.Writer.
func (e *Encoder) String(v string) error {
	if err := e.Uint32(uint32(len(v))); err != nil {
		return err
	}
	_, err := e.writer.Write([]byte(v))
	return err
}

// String encodes a string (c-style) to the Encoder's io.Writer.
func (e *Encoder) CString(v string) error {
	if _, err := e.writer.Write([]byte(v)); err != nil {
		return err
	}
	return e.Uint8(0)
}
