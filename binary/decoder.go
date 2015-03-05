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

// Decoder provides methods for decoding values to an io.Reader.
type Decoder struct {
	reader  io.Reader
	tmp     [9]byte
	objects map[uint32]interface{}
}

// NewDecoder creates a Decoder that reads from the provided io.Reader.
func NewDecoder(reader io.Reader) *Decoder {
	return &Decoder{reader: reader, objects: map[uint32]interface{}{}}
}

// Read implements the io.Reader interface, delegating to the underlying reader.
func (d *Decoder) Read(p []byte) (int, error) {
	return d.reader.Read(p)
}

// Read a byte array in its entirety.
func (d *Decoder) ReadFull(p []byte) error {
	_, err := io.ReadFull(d.reader, p)
	return err
}

// Bool decodes and returns a boolean value from the Decoder's io.Reader.
func (d *Decoder) Bool() (bool, error) {
	b := d.tmp[:1]
	_, err := io.ReadFull(d.reader, b[:1])
	return b[0] != 0, err
}

// Int8 decodes and returns a signed, 8 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int8() (int8, error) {
	i, err := d.Uint8()
	return int8(i), err
}

// Uint8 decodes and returns an unsigned, 8 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint8() (uint8, error) {
	b := d.tmp[:1]
	_, err := io.ReadFull(d.reader, b[:1])
	return b[0], err
}

// Int16 decodes and returns a signed, 16 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int16() (int16, error) {
	uv, err := d.Uint16()
	v := int16(uv >> 1)
	if uv&1 != 0 {
		v = ^v
	}
	return v, err
}

// Uint16 decodes and returns an unsigned, 16 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint16() (uint16, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:1])
	if err != nil {
		return 0, err
	}
	count := uint(0)
	for tag := d.tmp[0]; tag&0x80 != 0; tag = tag << 1 {
		count += 1
	}
	v := uint16(d.tmp[0] & (byte(0xff) >> count))
	if count == 0 {
		return v, nil
	}
	_, err = io.ReadFull(d.reader, d.tmp[:count])
	if err != nil {
		return 0, err
	}
	for i := uint(0); i < count; i++ {
		v = (v << 8) | uint16(d.tmp[i])
	}
	return v, nil
}

// Int32 decodes and returns a signed, 32 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int32() (int32, error) {
	uv, err := d.Uint32()
	v := int32(uv >> 1)
	if uv&1 != 0 {
		v = ^v
	}
	return v, err
}

// Uint32 decodes and returns an unsigned, 32 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint32() (uint32, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:1])
	if err != nil {
		return 0, err
	}
	count := uint(0)
	for tag := d.tmp[0]; tag&0x80 != 0; tag = tag << 1 {
		count += 1
	}
	v := uint32(d.tmp[0] & (byte(0xff) >> count))
	if count == 0 {
		return v, nil
	}
	_, err = io.ReadFull(d.reader, d.tmp[:count])
	if err != nil {
		return 0, err
	}
	for i := uint(0); i < count; i++ {
		v = (v << 8) | uint32(d.tmp[i])
	}
	return v, nil
}

// Float32 decodes and returns a 32 bit floating-point value from the Decoder's io.Reader.
func (d *Decoder) Float32() (float32, error) {
	bits, err := d.Uint32()
	shuffled := 0 |
		((bits & 0x000000ff) << 24) |
		((bits & 0x0000ff00) << 8) |
		((bits & 0x00ff0000) >> 8) |
		((bits & 0xff000000) >> 24)
	return math.Float32frombits(shuffled), err
}

// Int64 decodes and returns a signed, 64 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int64() (int64, error) {
	uv, err := d.Uint64()
	v := int64(uv >> 1)
	if uv&1 != 0 {
		v = ^v
	}
	return v, err
}

// Uint64 decodes and returns an unsigned, 64 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint64() (uint64, error) {
	_, err := io.ReadFull(d.reader, d.tmp[:1])
	if err != nil {
		return 0, err
	}
	count := uint(0)
	for tag := d.tmp[0]; tag&0x80 != 0; tag = tag << 1 {
		count += 1
	}
	v := uint64(d.tmp[0] & (byte(0xff) >> count))
	if count == 0 {
		return v, nil
	}
	_, err = io.ReadFull(d.reader, d.tmp[:count])
	if err != nil {
		return 0, err
	}
	for i := uint(0); i < count; i++ {
		v = (v << 8) | uint64(d.tmp[i])
	}
	return v, nil
}

// Float64 decodes and returns a 64 bit floating-point value from the Decoder's io.Reader.
func (d *Decoder) Float64() (float64, error) {
	bits, err := d.Uint64()
	shuffled := 0 |
		((bits & 0x00000000000000ff) << 56) |
		((bits & 0x000000000000ff00) << 40) |
		((bits & 0x0000000000ff0000) << 24) |
		((bits & 0x00000000ff000000) << 8) |
		((bits & 0x000000ff00000000) >> 8) |
		((bits & 0x0000ff0000000000) >> 24) |
		((bits & 0x00ff000000000000) >> 40) |
		((bits & 0xff00000000000000) >> 56)
	return math.Float64frombits(shuffled), err
}

// String decodes and returns a string from the Decoder's io.Reader.
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

// Object decodes and returns an Object from the Decoder's io.Reader. Object instances that were
// encoded multiple times will be decoded and returned as a shared, single instance.
// The type id in the stream must have been previously registered with binary.Register.
func (d *Decoder) Object() (interface{}, error) {
	key, err := d.Uint32()
	if err != nil {
		return nil, err
	}

	if key == 0 {
		return nil, nil
	}

	if obj, alreadyDecoded := d.objects[key]; alreadyDecoded {
		return obj, nil
	}

	var id ID
	if err := id.Decode(d); err != nil {
		return nil, err
	}

	obj, err := MakeObject(id)
	if err != nil {
		return nil, err
	}

	if err = obj.Decode(d); err != nil {
		return nil, err
	}

	d.objects[key] = obj
	return obj, nil
}
