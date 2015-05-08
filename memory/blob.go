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

package memory

import (
	"bytes"
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/log"
)

type blob struct {
	data []byte
	id   binary.ID
}

func (r blob) Get(db database.Database, logger log.Logger) ([]byte, error) {
	return r.data, nil
}

func (r blob) ResourceID(d database.Database, l log.Logger) (binary.ID, error) {
	return r.id, nil
}

func (r blob) Slice(rng Range) Slice {
	return Blob(r.data[rng.First() : rng.Last()+1])
}

func (r blob) Size() uint64 {
	return uint64(len(r.data))
}

func (r blob) String() string {
	return fmt.Sprintf("Blob[% x]", r.data)
}

// Blob returns a read-only Slice that wraps data.
func Blob(data []byte) Slice {
	return blob{data: data, id: binary.NewID(data)}
}

// Data returns a read-only Slice that contains the encoding of data.
func Data(arch device.Architecture, data ...interface{}) Slice {
	buf := &bytes.Buffer{}
	w := endian.Writer(buf, arch.ByteOrder)
	for _, d := range data {
		if err := Write(w, arch, d); err != nil {
			panic(err)
		}
	}
	return Blob(buf.Bytes())
}
