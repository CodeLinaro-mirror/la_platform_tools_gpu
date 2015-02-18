// Copyright (C) 2015 The Android Open Source Project
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

var (
	sampleId = ID{
		0x00, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x00,
		0x00, 0x01, 0x23, 0x45, 0x67, 0x89, 0xab, 0xcd, 0xef, 0x00,
	}
	sampleIdString = "000123456789abcdef00" + "000123456789abcdef00"
)

func TestIDToString(t *testing.T) {
	str := sampleId.String()
	if str != sampleIdString {
		t.Errorf("Expected string: %s, got: %s.", sampleIdString, str)
	}
}

func TestParseID(t *testing.T) {
	id, err := ParseID(sampleIdString)
	if !bytes.Equal(id[:], sampleId[:]) {
		t.Errorf("Expected ID: %+v, got: %+v.", sampleId, id)
	}
	if err != nil {
		t.Errorf("Unexpected error on ParseID(%s): %v", sampleIdString, err)
	}
}

func TestParseTooLongID(t *testing.T) {
	_, err := ParseID(sampleIdString + "00")
	if err == nil {
		t.Errorf("Expected an error when parsing a too long ID.")
	}
}

func TestParseTruncatedID(t *testing.T) {
	_, err := ParseID(sampleIdString[:len(sampleIdString)-2])
	if err == nil {
		t.Errorf("Expected an error when parsing a truncated ID.")
	}
}

func TestParseInvalidID(t *testing.T) {
	_, err := ParseID("abcdefghijklmnopqrst")
	if err == nil {
		t.Errorf("Expected an error when parsing an invalid ID.")
	}
}
