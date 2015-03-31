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
	"crypto/sha1"
	"encoding/hex"
	"fmt"
)

// IDSize is the size of an ID.
const IDSize = 20

// ID is a codeable unique identifier.
type ID [IDSize]byte

// Valid returns true if the id is not the default value.
func (id ID) Valid() bool {
	return id != ID{}
}

func (id ID) Format(f fmt.State, c rune) {
	fmt.Fprintf(f, "%x", id[:])
}

func (id ID) String() string {
	return hex.EncodeToString(id[:])
}

// ParseID parses lowercase string s as a 20 byte hex-encoded ID.
func ParseID(s string) (id ID, err error) {
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return
	}
	if len(bytes) != len(id) {
		err = fmt.Errorf("Invalid ID size: got %d, expected %d", len(bytes), len(id))
		return
	}
	copy(id[:], bytes)
	return
}

// Create a new ID that is the sha1 hash of the supplied data.
func NewID(data ...[]byte) ID {
	h := sha1.New()
	for _, d := range data {
		h.Write(d)
	}
	id := ID{}
	copy(id[:], h.Sum(nil))
	return id
}
