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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

type nullSlice uint64

func (s nullSlice) Get(database.Database, log.Logger) ([]byte, error) {
	return make([]byte, s), nil
}

func (s nullSlice) ResourceID(d database.Database, l log.Logger) (binary.ID, error) {
	return database.Store(make([]byte, s), d, l)
}

func (s nullSlice) Size() uint64 {
	return uint64(s)
}

func (s nullSlice) Slice(r Range) Slice {
	if uint64(r.Last()) > uint64(s) {
		panic(fmt.Errorf("nullSlice(%d).Slice(%v) - out of bounds", s, r))
	}
	return nullSlice(r.Size)
}

func (s nullSlice) ValidRanges() RangeList {
	return RangeList{}
}
