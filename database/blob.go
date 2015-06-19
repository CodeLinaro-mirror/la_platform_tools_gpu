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

package database

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Blob is an encodable wrapper for a byte array, used for storing raw data
// in databases.
type Blob struct {
	binary.Generate
	Data []byte
}

// StoreBlob stores the byte slice data inside a Blob to the database d.
func StoreBlob(data []byte, d Database, l log.Logger) (binary.ID, error) {
	return Store(d, &Blob{Data: data}, l)
}

// Resolve blob loads a Blob from the database, returning the byte slice.
func ResolveBlob(id binary.ID, d Database, l log.Logger) ([]byte, error) {
	b := Blob{}
	err := Load(d, id, l, &b)
	if err != nil {
		return nil, err
	}
	return b.Data, nil
}
