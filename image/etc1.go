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

package image

import "android.googlesource.com/platform/tools/gpu/binary"

type fmtETC1_RGB8 struct{ binary.Generate }

func (f *fmtETC1_RGB8) Key() interface{} { return *f }
func (*fmtETC1_RGB8) String() string     { return "ETC1_RGB8" }
func (*fmtETC1_RGB8) Size(w, h int) int {
	return (max(alignup(w, 4), 4) * max(alignup(h, 4), 4)) / 2
}
func (*fmtETC1_RGB8) Check(d []byte, w, h int) error {
	return checkSize(d, max(alignup(w, 4), 4), max(alignup(h, 4), 4), 4)
}

// ETC1_RGB8 returns a format representing the texture compression format
// with the same name.
func ETC1_RGB8() Format { return &fmtETC1_RGB8{} }

func init() {
	RegisterConverter(ETC1_RGB8(), RGBA(), func(src []byte, width, height int) ([]byte, error) {
		return Convert(src, width, height, ETC2_RGB8(), RGBA())
	})
}
