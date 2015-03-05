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
	"bytes"
	"encoding/binary"
	"io/ioutil"
	"testing"
)

type errorWriter struct{ e error }

func (e errorWriter) Write([]byte) (int, error) { return 0, e.e }

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
	}
	for _, e := range errs {
		if e.err != expectedError {
			t.Errorf("Encoding %s gave unexpected error. Expected: %v, got: %v", e.name, expectedError, e.err)
		}
	}
}

func TestEncoderNoError(t *testing.T) {
	e := NewEncoder(ioutil.Discard, binary.LittleEndian)
	checkEncoderGivesError(t, e, nil)
}

func TestEncoderTestError(t *testing.T) {
	e := NewEncoder(errorWriter{testError}, binary.LittleEndian)
	checkEncoderGivesError(t, e, testError)
}

func compareBytes(t *testing.T, got []byte, expected []byte) {
	if !bytes.Equal(expected, got) {
		t.Errorf(`Encode gave unexpected bytes.
Expected: %# x
Got:      %# x`, expected, got)
	}
}

func TestEncoderBool(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range boolValues {
		e.Bool(v)
	}
	compareBytes(t, b.Bytes(), boolBytes)
}

func TestEncoderInt8(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range int8Values {
		e.Int8(v)
	}
	compareBytes(t, b.Bytes(), int8Bytes)
}

func TestEncoderUint8(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range uint8Values {
		e.Uint8(v)
	}
	compareBytes(t, b.Bytes(), uint8Bytes)
}

func TestEncoderInt16(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range int16Values {
		e.Int16(v)
	}
	compareBytes(t, b.Bytes(), int16Bytes)
}

func TestEncoderUint16(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range uint16Values {
		e.Uint16(v)
	}
	compareBytes(t, b.Bytes(), uint16Bytes)
}

func TestEncoderInt32(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range int32Values {
		e.Int32(v)
	}
	compareBytes(t, b.Bytes(), int32Bytes)
}

func TestEncoderUint32(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range uint32Values {
		e.Uint32(v)
	}
	compareBytes(t, b.Bytes(), uint32Bytes)
}

func TestEncoderFloat32(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range float32Values {
		e.Float32(v)
	}
	compareBytes(t, b.Bytes(), float32Bytes)
}

func TestEncoderInt64(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range int64Values {
		e.Int64(v)
	}
	compareBytes(t, b.Bytes(), int64Bytes)
}

func TestEncoderUint64(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range uint64Values {
		e.Uint64(v)
	}
	compareBytes(t, b.Bytes(), uint64Bytes)
}

func TestEncoderFloat64(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, v := range float64Values {
		e.Float64(v)
	}
	compareBytes(t, b.Bytes(), float64Bytes)
}

func TestEncoderString(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, s := range stringValues {
		e.String(s)
	}
	compareBytes(t, b.Bytes(), stringBytes)
}

func TestEncoderCString(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b, binary.LittleEndian)
	for _, s := range stringValues {
		e.CString(s)
	}
	compareBytes(t, b.Bytes(), cStringBytes)
}
