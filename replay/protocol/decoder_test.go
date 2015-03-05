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
	"errors"
	"testing"
)

type errorReader struct{ e error }

func (e errorReader) Read([]byte) (int, error) { return 0, e.e }

var testError = errors.New("test error")

func TestDecoderBool(t *testing.T) {
	d := NewDecoder(bytes.NewBuffer(boolBytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(int8Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(uint8Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(int16Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(uint16Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(int32Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(uint32Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(float32Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(float64Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(int64Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(uint64Bytes), binary.LittleEndian)
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
	d := NewDecoder(bytes.NewBuffer(stringBytes), binary.LittleEndian)
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
	d := NewDecoder(errorReader{testError}, binary.LittleEndian)
	if _, err := d.String(); err != testError {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
}

func TestDecoderCStringError(t *testing.T) {
	d := NewDecoder(errorReader{testError}, binary.LittleEndian)
	if _, err := d.CString(); err != testError {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
}
