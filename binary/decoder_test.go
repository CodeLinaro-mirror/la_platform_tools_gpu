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
	"errors"
	"io"
	"reflect"
	"testing"
)

type errorReader struct{ e error }

func (e errorReader) Read([]byte) (int, error) { return 0, e.e }

var testError = errors.New("test error")

var testDecodeObjectsBuffer = []byte{
	0x00, 0x00,
	byte(testObjectIDA & 0xff), byte((testObjectIDA >> 8) & 0xff),
	0x07, 0x00, 0x00, 0x00,
	'O', 'b', 'j', 'e', 'c', 't', 'A',

	0x01, 0x00,
	byte(testObjectIDB & 0xff), byte((testObjectIDB >> 8) & 0xff),
	0x07, 0x00, 0x00, 0x00,
	'O', 'b', 'j', 'e', 'c', 't', 'B',

	0x00, 0x00,

	byte(objectNil & 0xff), byte((objectNil >> 8) & 0xff),
}

func TestDecoderBool(t *testing.T) {
	d := BufferDecoder([]byte{0, 1})
	for i, expected := range []bool{false, true} {
		got, err := d.Bool()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderInt8(t *testing.T) {
	d := BufferDecoder([]byte{0x00, 0x7f, 0x80, 0xff})
	for i, expected := range []int8{0, 127, -128, -1} {
		got, err := d.Int8()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderUint8(t *testing.T) {
	d := BufferDecoder([]byte{0x00, 0x7f, 0x80, 0xff})
	for i, expected := range []uint8{0x00, 0x7f, 0x80, 0xff} {
		got, err := d.Uint8()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderInt16(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00,
		0xff, 0x7f,
		0x00, 0x80,
		0xff, 0xff,
	})
	for i, expected := range []int16{0, 32767, -32768, -1} {
		got, err := d.Int16()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderUint16(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00,
		0xef, 0xbe,
		0xde, 0xc0,
	})
	for i, expected := range []uint16{0, 0xbeef, 0xc0de} {
		got, err := d.Uint16()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderInt32(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x7f,
		0x00, 0x00, 0x00, 0x80,
		0xff, 0xff, 0xff, 0xff,
	})
	for i, expected := range []int32{0, 2147483647, -2147483648, -1} {
		got, err := d.Int32()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderUint32(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00, 0x00, 0x00,
		0x67, 0x45, 0x23, 0x01,
		0xef, 0xcd, 0xab, 0x10,
	})
	for i, expected := range []uint32{0, 0x01234567, 0x10abcdef} {
		got, err := d.Uint32()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderFloat32(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x80, 0x3f,
		0x00, 0x00, 0x81, 0x42,
	})
	for i, expected := range []float32{0, 1, 64.5} {
		got, err := d.Float32()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderFloat64(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x50, 0x40,
	})
	for i, expected := range []float64{0, 1, 64.5} {
		got, err := d.Float64()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderInt64(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	})
	for i, expected := range []int64{0, 9223372036854775807, -9223372036854775808, -1} {
		got, err := d.Int64()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderUint64(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01,
		0x10, 0x32, 0x54, 0x76, 0x98, 0xba, 0xdc, 0xfe,
	})
	for i, expected := range []uint64{0, 0x0123456789abcdef, 0xfedcba9876543210} {
		got, err := d.Uint64()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderString(t *testing.T) {
	d := BufferDecoder([]byte{
		0x05, 0x00, 0x00, 0x00,
		'H', 'e', 'l', 'l', 'o',
		0x00, 0x00, 0x00, 0x00,
		0x05, 0x00, 0x00, 0x00,
		'W', 'o', 'r', 'l', 'd',
		0x15, 0x00, 0x00, 0x00,
		0xe3, 0x81, 0x93, 0xe3, 0x82, 0x93, 0xe3, 0x81, 0xab, 0xe3, 0x81, 0xa1, 0xe3, 0x81, 0xaf, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c,
	})
	for i, expected := range []string{"Hello", "", "World", "こんにちは世界"} {
		got, err := d.String()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderStringError(t *testing.T) {
	d := Decoder{Reader: errorReader{testError}}
	if _, err := d.String(); err != testError {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
}

func TestDecoderCString(t *testing.T) {
	d := BufferDecoder([]byte{
		'H', 'e', 'l', 'l', 'o', 0x00,
		0x00,
		'W', 'o', 'r', 'l', 'd', 0x00,
		0xe3, 0x81, 0x93, 0xe3, 0x82, 0x93, 0xe3, 0x81, 0xab, 0xe3, 0x81, 0xa1, 0xe3, 0x81, 0xaf, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0x00,
	})
	for i, expected := range []string{"Hello", "", "World", "こんにちは世界"} {
		got, err := d.CString()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if expected != got {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}
}

func TestDecoderCStringError(t *testing.T) {
	d := Decoder{Reader: errorReader{testError}}
	if _, err := d.CString(); err != testError {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
}

func TestDecoderData(t *testing.T) {
	d := BufferDecoder([]byte{
		0x06, 0x00, 0x00, 0x00,
		0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc,
	})
	expected := []byte{0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc}
	got, err := d.Data()
	if err != nil {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
	if !bytes.Equal(expected, got) {
		t.Errorf("Decode gave unexpected value. Expected: %v, got: %v", expected, got)
	}
}

func TestDecoderDataError(t *testing.T) {
	d := Decoder{Reader: errorReader{testError}}
	data, err := d.Data()
	if err != testError {
		t.Errorf("Decode gave unexpected error. Expected: %v, got: %v", testError, err)
	}
	if data != nil {
		t.Errorf("Decode gave unexpected value. Expected: %v, got: %v", nil, data)
	}
}

func checkObjectDecode(t *testing.T, d *Decoder, expectedObj Object, expectedErr error) {
	obj, err := d.Object()
	if err != expectedErr {
		t.Errorf("Decode gave unexpected error. Expected, %v, got: %v", expectedErr, err)
	}
	if obj != expectedObj {
		t.Errorf("Decode gave unexpected object. Expected, %v, got: %v", expectedObj, obj)
	}
}

func TestDecoderObject(t *testing.T) {
	d := BufferDecoder(testDecodeObjectsBuffer).WithNamespace(testObjectNamespace)

	var got [4]Object
	var err error

	for i, expected := range []Object{testObjA, testObjB, testObjA, nil} {
		got[i], err = d.Object()
		if err != nil {
			t.Errorf("Decode %d gave unexpected error: %v", i, err)
		}
		if !reflect.DeepEqual(expected, got[i]) {
			t.Errorf("Decode %d gave unexpected value. Expected: %v, got: %v", i, expected, got)
		}
	}

	if got[0] != got[2] {
		t.Errorf("Decode of same object gave difference instance. First: %p Second: %p", got[0], got[2])
	}
}

func TestDecoderObjectUnknownTypeIDError(t *testing.T) {
	d := BufferDecoder([]byte{0x00, 0x00, 0x00, 0x00})
	checkObjectDecode(t, d, nil, UnknownTypeIDError{0})
}

func TestDecoderObjectReadError(t *testing.T) {
	d := Decoder{Reader: errorReader{testError}}
	checkObjectDecode(t, d.WithNamespace(testObjectNamespace), nil, testError)
}

func TestDecoderObjectDecodeError(t *testing.T) {
	d := BufferDecoder([]byte{
		0x00, 0x00,
		byte(testObjectIDA & 0xff), byte((testObjectIDA >> 8) & 0xff),
		0x07, 0x00, 0x00, 0x00,
		'O', 'b', 'j', 'e', // incomplete data
	})
	checkObjectDecode(t, d.WithNamespace(testObjectNamespace), nil, io.ErrUnexpectedEOF)
}

func TestDecoderObjectEOFError(t *testing.T) {
	d := BufferDecoder([]byte{0x00, 0x00})
	checkObjectDecode(t, d.WithNamespace(testObjectNamespace), nil, io.EOF)
}

func TestUnknownTypeIDErrorError(t *testing.T) {
	expected := "Decoder's TypeNamespace did not contain decoded type id 100"
	got := UnknownTypeIDError{100}.Error()
	if expected != got {
		t.Errorf("UnknownTypeIDError.Error() did not return expected result. Expected: %v, got: %v", expected, got)
	}
}

func BenchmarkDecoderObject(b *testing.B) {
	decoders := make([]*Decoder, b.N)
	for i := range decoders {
		decoders[i] = BufferDecoder(testDecodeObjectsBuffer).WithNamespace(testObjectNamespace)
	}
	b.ResetTimer()

	for _, d := range decoders {
		d.Object()
		d.Object()
		d.Object()
		d.Object()
	}
}
