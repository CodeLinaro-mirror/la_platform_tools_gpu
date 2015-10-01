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

type fmtRGB565 struct{ binary.Generate }

func (f *fmtRGB565) Key() interface{}             { return *f }
func (*fmtRGB565) String() string                 { return "RGB565" }
func (*fmtRGB565) Size(w, h int) int              { return w * h * 2 }
func (*fmtRGB565) Check(d []byte, w, h int) error { return checkSize(d, w, h, 16) }

// RGB565 returns a format containing an 5-bit red, 6-bit green and 5-bit blue
// channels per pixel.
func RGB565() Format { return &fmtRGB565{} }

func init() {
	RegisterConverter(RGB565(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			dst, i, j := make([]byte, width*height*4), 0, 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					rgb := (uint16(src[0]) << 8) | uint16(src[1])
					r := uint8((rgb & 0xF800) >> 11)
					g := uint8((rgb & 0x07E0) >> 5)
					b := uint8((rgb & 0x001F))
					dst[j+0], dst[j+1], dst[j+2], dst[j+3] = r, g, b, 255
					i += 2
					j += 4
				}
			}

			return dst, nil
		})
}
