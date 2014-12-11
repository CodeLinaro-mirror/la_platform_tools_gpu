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
	"bytes"
	"fmt"
	"io"
	"math"
)

// UnknownTypeIDError is returned by Decoder.Object when decoding an Object which has a type that
// does not exist in the Decoder's TypeNamespace.
type UnknownTypeIDError struct {
	ID ObjectTypeID // The type id that is not in the Decoder's TypeNamespace
}

func (e UnknownTypeIDError) Error() string {
	return fmt.Sprintf("Decoder's TypeNamespace did not contain decoded type id %v", e.ID)
}

// Decoder provides methods for decoding values to an io.Reader.
type Decoder struct {
	Reader    io.Reader
	tmp       [8]byte
	objects   map[uint16]Object
	namespace TypeNamespace
}

// BufferDecoder creates a Decoder that reads from the provided byte slice.
func BufferDecoder(buf []byte) *Decoder {
	return &Decoder{Reader: bytes.NewBuffer(buf)}
}

// WithNamespace returns a new Decoder that will use the specified TypeNamespace for constructing
// encoded objects. Any previously assigned TypeNamespace will not be used by the returned Decoder.
func (d *Decoder) WithNamespace(namespace TypeNamespace) *Decoder {
	if d.objects == nil {
		d.objects = make(map[uint16]Object)
	}
	return &Decoder{
		Reader:    d.Reader,
		objects:   d.objects,
		namespace: namespace,
	}
}

// Bool decodes and returns a boolean value from the Decoder's io.Reader.
func (d *Decoder) Bool() (bool, error) {
	b := d.tmp[:1]
	_, err := io.ReadFull(d.Reader, b[:1])
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
	_, err := io.ReadFull(d.Reader, b[:1])
	return b[0], err
}

// Int16 decodes and returns a signed, 16 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int16() (int16, error) {
	i, err := d.Uint16()
	return int16(i), err
}

// Uint16 decodes and returns an unsigned, 16 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint16() (uint16, error) {
	b := d.tmp[:2]
	_, err := io.ReadFull(d.Reader, b)
	return uint16(b[0]) | uint16(b[1])<<8, err
}

// Int32 decodes and returns a signed, 32 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int32() (int32, error) {
	i, err := d.Uint32()
	return int32(i), err
}

// Uint32 decodes and returns an unsigned, 32 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint32() (uint32, error) {
	b := d.tmp[:4]
	_, err := io.ReadFull(d.Reader, b[:])
	return uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24, err
}

// Float32 decodes and returns a 32 bit floating-point value from the Decoder's io.Reader.
func (d *Decoder) Float32() (float32, error) {
	i, err := d.Uint32()
	return math.Float32frombits(i), err
}

// Int64 decodes and returns a signed, 64 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Int64() (int64, error) {
	i, err := d.Uint64()
	return int64(i), err
}

// Uint64 decodes and returns an unsigned, 64 bit integer value from the Decoder's io.Reader.
func (d *Decoder) Uint64() (uint64, error) {
	b := d.tmp[:8]
	_, err := io.ReadFull(d.Reader, b[:])
	return uint64(b[0]) |
		uint64(b[1])<<8 |
		uint64(b[2])<<16 |
		uint64(b[3])<<24 |
		uint64(b[4])<<32 |
		uint64(b[5])<<40 |
		uint64(b[6])<<48 |
		uint64(b[7])<<56, err
}

// Float64 decodes and returns a 64 bit floating-point value from the Decoder's io.Reader.
func (d *Decoder) Float64() (float64, error) {
	i, err := d.Uint64()
	return math.Float64frombits(i), err
}

// String decodes and returns a string (pascal-style) from the Decoder's io.Reader.
func (d *Decoder) String() (string, error) {
	c, err := d.Uint32()
	if err != nil {
		return "", err
	}
	if c > 0 {
		s := make([]byte, c)
		_, err := io.ReadFull(d.Reader, s)
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

// Data decodes and returns a sequence of bytes from the Decoder's io.Reader.
func (d *Decoder) Data() ([]byte, error) {
	c, err := d.Uint32()
	if err != nil {
		return nil, err
	}

	buf := make([]byte, c)
	_, err = io.ReadFull(d.Reader, buf)

	return buf, err
}

// Object decodes and returns an Object from the Decoder's io.Reader. Object instances that were
// encoded multiple times will be decoded and returned as a shared, single instance.
// The Decoder must be bound to the same TypeNamespace used to encode the object, otherwise the
// decoder may return an incorrectly decoded object, or may return an UnknownTypeIDError.
func (d *Decoder) Object() (Object, error) {
	key, err := d.Uint16()
	if err != nil {
		return nil, err
	}

	if key == objectNil {
		return nil, nil
	}

	obj, alreadyDecoded := d.objects[key]
	if alreadyDecoded {
		return obj, nil
	}

	ty, err := d.Uint16()
	if err != nil {
		return nil, err
	}

	if obj = d.namespace.new(ObjectTypeID(ty)); obj == nil {
		return nil, UnknownTypeIDError{ObjectTypeID(ty)}
	}

	if err = obj.Decode(d); err != nil {
		return nil, err
	}

	d.objects[key] = obj
	return obj, nil
}
