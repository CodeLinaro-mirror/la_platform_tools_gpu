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
	"io"
	"io/ioutil"
	"testing"
)

type errorWriter struct{ e error }

func (e errorWriter) Write([]byte) (int, error) { return 0, e.e }

type shortWriter struct{}

func (e shortWriter) Write([]byte) (int, error) { return 0, nil }

func checkEncoderGivesError(t *testing.T, e *Encoder, expectedError error) {
	errs := []struct {
		name string
		err  error
	}{
		{"Bool", e.Bool(false)},
		{"Int8", e.Int8(42)},
		{"Uint8", e.Uint8(42)},
		{"Int16", e.Int16(42)},
		{"Uint16", e.Uint16(42)},
		{"Int32", e.Int32(42)},
		{"Uint32", e.Uint32(42)},
		{"Int64", e.Int64(42)},
		{"Uint64", e.Uint64(42)},
		{"String", e.String("string")},
		{"CString", e.CString("string")},
		{"Data", e.Data([]byte{1, 2, 3, 4})},
		{"Object", e.Object(testObjA)},
	}
	for _, e := range errs {
		if e.err != expectedError {
			t.Errorf("Encoding %s gave unexpected error. Expected: %v, got: %v", e.name, expectedError, e.err)
		}
	}
}

func TestEncoderNoError(t *testing.T) {
	e := NewEncoder(ioutil.Discard)
	checkEncoderGivesError(t, e, nil)
}

func TestEncoderShortWrite(t *testing.T) {
	e := NewEncoder(shortWriter{})
	checkEncoderGivesError(t, e, io.ErrShortWrite)
}

func TestEncoderTestError(t *testing.T) {
	e := NewEncoder(errorWriter{testError})
	checkEncoderGivesError(t, e, testError)
}

func TestEncoderBool(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Bool(false)
	e.Bool(true)
	expected, got := []byte{0, 1}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderInt8(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Int8(0)
	e.Int8(127)
	e.Int8(-128)
	e.Int8(-1)
	expected, got := []byte{0x00, 0x7f, 0x80, 0xff}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderUint8(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Uint8(0x00)
	e.Uint8(0x7f)
	e.Uint8(0x80)
	e.Uint8(0xff)
	expected, got := []byte{0x00, 0x7f, 0x80, 0xff}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderInt16(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Int16(0)
	e.Int16(32767)
	e.Int16(-32768)
	e.Int16(-1)
	expected, got := []byte{
		0x00, 0x00,
		0xff, 0x7f,
		0x00, 0x80,
		0xff, 0xff,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderUint16(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Uint16(0)
	e.Uint16(0xbeef)
	e.Uint16(0xc0de)
	expected, got := []byte{
		0x00, 0x00,
		0xef, 0xbe,
		0xde, 0xc0,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderInt32(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Int32(0)
	e.Int32(2147483647)
	e.Int32(-2147483648)
	e.Int32(-1)
	expected, got := []byte{
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0x7f,
		0x00, 0x00, 0x00, 0x80,
		0xff, 0xff, 0xff, 0xff,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderUint32(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Uint32(0)
	e.Uint32(0x01234567)
	e.Uint32(0x10abcdef)
	expected, got := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x67, 0x45, 0x23, 0x01,
		0xef, 0xcd, 0xab, 0x10,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderFloat32(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Float32(0)
	e.Float32(1)
	e.Float32(64.5)
	expected, got := []byte{
		0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x80, 0x3f,
		0x00, 0x00, 0x81, 0x42,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderInt64(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Int64(0)
	e.Int64(9223372036854775807)
	e.Int64(-9223372036854775808)
	e.Int64(-1)
	expected, got := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x80,
		0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderUint64(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Uint64(0)
	e.Uint64(0x0123456789abcdef)
	e.Uint64(0xfedcba9876543210)
	expected, got := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0xef, 0xcd, 0xab, 0x89, 0x67, 0x45, 0x23, 0x01,
		0x10, 0x32, 0x54, 0x76, 0x98, 0xba, 0xdc, 0xfe,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderFloat64(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Float64(0)
	e.Float64(1)
	e.Float64(64.5)

	expected, got := []byte{
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0xf0, 0x3f,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x20, 0x50, 0x40,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderString(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.String("Hello")
	e.String("")
	e.String("World")
	e.String("こんにちは世界")

	expected, got := []byte{
		0x05, 0x00, 0x00, 0x00,
		'H', 'e', 'l', 'l', 'o',
		0x00, 0x00, 0x00, 0x00,
		0x05, 0x00, 0x00, 0x00,
		'W', 'o', 'r', 'l', 'd',
		0x15, 0x00, 0x00, 0x00,
		0xe3, 0x81, 0x93, 0xe3, 0x82, 0x93, 0xe3, 0x81, 0xab, 0xe3, 0x81, 0xa1, 0xe3, 0x81, 0xaf, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderCString(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.CString("Hello")
	e.CString("")
	e.CString("World")
	e.CString("こんにちは世界")

	expected, got := []byte{
		'H', 'e', 'l', 'l', 'o', 0x00,
		0x00,
		'W', 'o', 'r', 'l', 'd', 0x00,
		0xe3, 0x81, 0x93, 0xe3, 0x82, 0x93, 0xe3, 0x81, 0xab, 0xe3, 0x81, 0xa1, 0xe3, 0x81, 0xaf, 0xe4, 0xb8, 0x96, 0xe7, 0x95, 0x8c, 0x00,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderData(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	e.Data([]byte{0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc})
	expected, got := []byte{
		0x06, 0x00, 0x00, 0x00,
		0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderObject(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	if err := e.Object(testObjA); err != nil {
		t.Errorf("Encode gave unexpected error: %v", err)
	}
	if err := e.Object(testObjB); err != nil {
		t.Errorf("Encode gave unexpected error: %v", err)
	}
	if err := e.Object(testObjA); err != nil {
		t.Errorf("Encode gave unexpected error: %v", err)
	}
	if err := e.Object(nil); err != nil {
		t.Errorf("Encode gave unexpected error: %v", err)
	}
	expected, got := []byte{
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
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestEncoderObjectUnknownTypeError(t *testing.T) {
	e := NewEncoder(ioutil.Discard)
	expected := unknownType{testObjC}
	if err := e.Object(testObjC); err != expected {
		t.Errorf("Encode gave unexpected error. Expected: %v, got: %v", expected, err)
	}
}

func TestUnknownTypeErrorError(t *testing.T) {
	expected := "Unknown type *binary.testObjectC encountered in binary.Encoder"
	got := unknownType{&testObjectC{}}.Error()
	if expected != got {
		t.Errorf("Error() did not return expected result. Expected: %v, got: %v", expected, got)
	}
}

func BenchmarkEncoderObject(b *testing.B) {
	encoders := make([]*Encoder, b.N)
	for i := range encoders {
		encoders[i] = NewEncoder(ioutil.Discard)
	}
	b.ResetTimer()

	for _, e := range encoders {
		e.Object(testObjA)
		e.Object(testObjB)
		e.Object(testObjA)
		e.Object(testObjB)
	}
}
