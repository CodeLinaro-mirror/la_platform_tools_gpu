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
	"strings"
	"testing"
)

type errorReader struct{ e error }

func (e errorReader) Read([]byte) (int, error) { return 0, e.e }

var testError = errors.New("test error")

var invalidTypeId = ID{0x01}

var testDecodeObjectsBuffer = []byte{
	0x00, 0x00,
	0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x07, 0x00, 0x00, 0x00,
	'O', 'b', 'j', 'e', 'c', 't', 'A',

	0x01, 0x00,
	0x0B, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	0x07, 0x00, 0x00, 0x00,
	'O', 'b', 'j', 'e', 'c', 't', 'B',

	0x00, 0x00,

	byte(objectNil & 0xff), byte((objectNil >> 8) & 0xff),
}

func TestDecoderBool(t *testing.T) {
	d := NewDecoder(bytes.NewBuffer(boolBytes))
	for i, expected := range boolValues {
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
	d := NewDecoder(bytes.NewBuffer(int8Bytes))
	for i, expected := range int8Values {
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
	d := NewDecoder(bytes.NewBuffer(uint8Bytes))
	for i, expected := range uint8Values {
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
	d := NewDecoder(bytes.NewBuffer(int16Bytes))
	for i, expected := range int16Values {
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
	d := NewDecoder(bytes.NewBuffer(uint16Bytes))
	for i, expected := range uint16Values {
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
	d := NewDecoder(bytes.NewBuffer(int32Bytes))
	for i, expected := range int32Values {
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
	d := NewDecoder(bytes.NewBuffer(uint32Bytes))
	for i, expected := range uint32Values {
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
	d := NewDecoder(bytes.NewBuffer(float32Bytes))
	for i, expected := range float32Values {
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
	d := NewDecoder(bytes.NewBuffer(float64Bytes))
	for i, expected := range float64Values {
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
	d := NewDecoder(bytes.NewBuffer(int64Bytes))
	for i, expected := range int64Values {
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
	d := NewDecoder(bytes.NewBuffer(uint64Bytes))
	for i, expected := range uint64Values {
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
	d := NewDecoder(bytes.NewBuffer(stringBytes))
	for i, expected := range stringValues {
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
	d := NewDecoder(errorReader{testError})
	if _, err := d.String(); err != testError {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
}

func TestDecoderCString(t *testing.T) {
	d := NewDecoder(bytes.NewBuffer(cStringBytes))
	for i, expected := range stringValues {
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
	d := NewDecoder(errorReader{testError})
	if _, err := d.CString(); err != testError {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
}

func checkObjectDecode(t *testing.T, d *Decoder, expectedObj interface{}, expectedErr error) {
	obj, err := d.Object()
	if err != expectedErr {
		t.Errorf("Decode gave unexpected error. Expected: %v, got: %v", expectedErr, err)
	}
	if obj != expectedObj {
		t.Errorf("Decode gave unexpected object. Expected: %v, got: %v", expectedObj, obj)
	}
}

func TestDecoderObject(t *testing.T) {
	d := NewDecoder(bytes.NewBuffer(objectBytes))

	var got [4]interface{}
	var err error

	for i, expected := range objectValues {
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
	d := NewDecoder(bytes.NewBuffer([]byte{
		0x00, 0x00,
		0x01, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00,
	}))
	checkObjectDecode(t, d, nil, unknownTypeID(invalidTypeId))
}

func TestDecoderObjectReadError(t *testing.T) {
	d := NewDecoder(errorReader{testError})
	checkObjectDecode(t, d, nil, testError)
}

func TestDecoderObjectDecodeError(t *testing.T) {
	d := NewDecoder(bytes.NewBuffer([]byte{
		0x00, 0x00,
		0x0A, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x07, 0x00, 0x00, 0x00,
		'O', 'b', 'j', 'e', // incomplete data
	}))
	checkObjectDecode(t, d, nil, io.ErrUnexpectedEOF)
}

func TestDecoderObjectEOFError(t *testing.T) {
	d := NewDecoder(bytes.NewBuffer([]byte{0x00, 0x00}))
	checkObjectDecode(t, d, nil, io.EOF)
}

func TestUnknownTypeIDErrorError(t *testing.T) {
	expected := "Unknown type id 0100000000000000000000000000000000000000"
	got := unknownTypeID(invalidTypeId).Error()
	if !strings.Contains(got, expected) {
		t.Errorf("Error() did not return expected result. Expected: %v, got: %v", expected, got)
	}
}

func BenchmarkDecoderObject(b *testing.B) {
	decoders := make([]*Decoder, b.N)
	for i := range decoders {
		decoders[i] = NewDecoder(bytes.NewBuffer(testDecodeObjectsBuffer))
	}
	b.ResetTimer()

	for _, d := range decoders {
		d.Object()
		d.Object()
		d.Object()
		d.Object()
	}
}
