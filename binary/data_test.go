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
	"testing"
)

func TestDataEncode(t *testing.T) {
	b := &bytes.Buffer{}
	e := NewEncoder(b)
	Data{0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc}.Encode(e)
	expected, got := []byte{
		0x06,
		0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc,
	}, b.Bytes()
	if !bytes.Equal(expected, got) {
		t.Errorf("Encode gave unexpected bytes. Expected: %v, got: %v", expected, got)
	}
}

func TestDataDecode(t *testing.T) {
	d := NewDecoder(bytes.NewBuffer([]byte{
		0x06,
		0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc,
	}))
	expected := []byte{0x10, 0x20, 0x30, 0xaa, 0xbb, 0xcc}
	got := Data{}
	err := got.Decode(d)
	if err != nil {
		t.Errorf("Decode gave unexpected error: %v", err)
	}
	if !bytes.Equal(expected, got) {
		t.Errorf("Decode gave unexpected value. Expected: %v, got: %v", expected, got)
	}
}

func TestDataDecodeError(t *testing.T) {
	d := NewDecoder(errorReader{testError})
	data := Data{}
	err := data.Decode(d)
	if err != testError {
		t.Errorf("Decode gave unexpected error. Expected: %v, got: %v", testError, err)
	}
}
