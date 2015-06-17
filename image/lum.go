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

type fmtLuminance struct{ binary.Generate }

func (*fmtLuminance) String() string                 { return "Luminance" }
func (*fmtLuminance) Check(d []byte, w, h int) error { return checkSize(d, w, h, 8) }

// Luminance returns a format containing a single 8-bit luminance channel per
// pixel.
func Luminance() Format { return &fmtLuminance{} }

func init() {
	RegisterConverter(Luminance(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			dst, i, j := make([]byte, width*height*4), 0, 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					dst[j+0], dst[j+1], dst[j+2], dst[j+3] = src[i], src[i], src[i], 255
					i++
					j += 4
				}
			}

			return dst, nil
		})
}
