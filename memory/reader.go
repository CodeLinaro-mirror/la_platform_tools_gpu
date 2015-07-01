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
	"bufio"
	"io"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Maximum number of bytes to buffer
const maxReadBuffer = 1024

// Reader returns a binary reader for the specified Slice.
func Reader(s Slice, d database.Database, l log.Logger) io.Reader {
	c := s.Size()
	return bufio.NewReaderSize(&reader{s, d, l, c, 0}, int(min(c, maxReadBuffer)))
}

type reader struct {
	s Slice
	d database.Database
	l log.Logger
	r uint64
	o uint64
}

func (r *reader) Read(dst []byte) (n int, err error) {
	if r.r == 0 {
		return 0, io.EOF
	}
	c := min(uint64(len(dst)), r.r)
	src, err := r.s.Slice(Range{Base: r.o, Size: c}).Get(r.d, r.l)
	r.o += c
	r.r -= c
	return copy(dst, src), err
}
