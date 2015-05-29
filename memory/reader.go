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
	"io"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Reader returns a binary reader for the specified Slice.
func Reader(s Slice, d database.Database, l log.Logger) io.Reader {
	return &reader{s, d, l, 0}
}

type reader struct {
	s Slice
	d database.Database
	l log.Logger
	o Pointer
}

func (r *reader) Read(dst []byte) (n int, err error) {
	src, err := r.s.Slice(Range{Base: r.o, Size: uint64(len(dst))}).Get(r.d, r.l)
	n = copy(dst, src)
	r.o += Pointer(n)
	return n, err
}
