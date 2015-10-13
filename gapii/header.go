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

package gapii

import (
	"io"

	"android.googlesource.com/platform/tools/gpu/binary/endian"
)

var magic = [4]byte{'s', 'p', 'y', '0'}

const version = 2

// The GAPII header version 2 is defined as:
//
// struct ConnectionHeader {
//   uint8_t  mMagic[4];                     // 's', 'p', 'y', '0'
//   uint32_t mVersion;                      // 2
//   uint32_t mObserveFrameFrequency;        // non-zero == enabled
//   uint32_t mObserveDrawFrequency;         // non-zero == enabled
// };
//
// All fields are encoded little-endian with no compression, regardless of
// architecture. All changes must be kept in sync with:
//   platform/tools/gpu/cc/gapii/connection_header.h

func sendHeader(out io.Writer, options Options) error {
	w := endian.Writer(out, endian.Little)
	for _, m := range magic {
		w.Uint8(m)
	}
	w.Uint32(version)
	w.Uint32(options.ObserveFrameFreqency)
	w.Uint32(options.ObserveDrawFrequency)
	return w.Error()
}
