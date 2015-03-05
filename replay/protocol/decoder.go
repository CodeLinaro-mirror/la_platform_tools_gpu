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

// Decoder provides methods for decoding values from an io.Reader.
type Decoder struct {
	reader    io.Reader
	tmp       [8]byte
	byteOrder binary.ByteOrder
}

// NewDecoder creates a Decoder that reads from the provided io.Reader, with the
// specified endianness.
func NewDecoder(reader io.Reader, byteOrder binary.ByteOrder) *Decoder {
	return &Decoder{reader: reader, byteOrder: byteOrder}
}

// Read implements the io.Reader interface, delegating to the underlying reader.
func (d *Decoder) Read(p []byte) (int, error) {
	return d.reader.Read(p)
}

// Bool decodes and returns a boolean value from the Decoder's io.Reader.
func (d *Decoder) Bool() (bool, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:1])
	return d.tmp[0] != 0, err
}

// Int8 decodes and returns a signed, 8 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int8() (int8, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:1])
	return int8(d.tmp[0]), err
}

// Uint8 decodes and returns an unsigned, 8 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint8() (uint8, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:1])
	return d.tmp[0], err
}

// Int16 decodes and returns a signed, 16 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int16() (int16, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:2])
	return int16(d.byteOrder.Uint16(d.tmp[:])), err
}

// Uint16 decodes and returns an unsigned, 16 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint16() (uint16, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:2])
	return d.byteOrder.Uint16(d.tmp[:]), err
}

// Int32 decodes and returns a signed, 32 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int32() (int32, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:4])
	return int32(d.byteOrder.Uint32(d.tmp[:])), err
}

// Uint32 decodes and returns an unsigned, 32 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint32() (uint32, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:4])
	return d.byteOrder.Uint32(d.tmp[:]), err
}

// Float32 decodes and returns a 32 bit floating-point value from the Decoder's io.Reader.
func (d *Decoder) Float32() (float32, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:4])
	return math.Float32frombits(d.byteOrder.Uint32(d.tmp[:])), err
}

// Int64 decodes and returns a signed, 64 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int64() (int64, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:8])
	return int64(d.byteOrder.Uint64(d.tmp[:])), err
}

// Uint64 decodes and returns an unsigned, 64 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint64() (uint64, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:8])
	return d.byteOrder.Uint64(d.tmp[:]), err
}

// Float64 decodes and returns a 64 bit floating-point value from the Decoder's io.Reader.
func (d *Decoder) Float64() (float64, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:8])
	return math.Float64frombits(d.byteOrder.Uint64(d.tmp[:])), err
}

// String decodes and returns a string (pascal-style) from the Decoder's io.Reader.
func (d *Decoder) String() (string, error) {
	c, err := d.Uint32()
	if err != nil {
		return "", err
	}
	if c > 0 {
		s := make([]byte, c)
		_, err := io.ReadFull(d.reader, s)
		return string(s), err
	} else {
		return "", nil
	}
}

// String decodes and returns a string (c-style) from the Decoder's io.Reader.
func (d *Decoder) CString() (string, error) {
	s := []byte{}
	for {
		c, err := d.Uint8()
		if err != nil {
			return "", err
		}
		if c == 0 {
			break
		}
		s = append(s, c)
	}
	return string(s), nil
}
