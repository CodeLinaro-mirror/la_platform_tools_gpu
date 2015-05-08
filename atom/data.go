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

package atom

import (
	"bytes"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

// Data encodes and stores the value v to the database d, returning the
// memory range and new resource identifier. Data can be used to as a helper
// to AddRead and AddWrite methods on atoms.
func Data(a device.Architecture, d database.Database, l log.Logger, at memory.Pointer, v ...interface{}) (memory.Range, binary.ID) {
	buf := &bytes.Buffer{}
	w := endian.Writer(buf, a.ByteOrder)
	if err := memory.Write(w, a, v); err != nil {
		panic(err)
	}
	id, err := d.Store(&store.Blob{Data: buf.Bytes()}, l)
	if err != nil {
		panic(err)
	}
	return at.Range(uint64(len(buf.Bytes()))), id
}
