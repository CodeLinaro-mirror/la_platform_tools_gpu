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

type fmtRGBA5551 struct{ binary.Generate }

func (f *fmtRGBA5551) Key() interface{}             { return *f }
func (*fmtRGBA5551) String() string                 { return "RGB5551" }
func (*fmtRGBA5551) Size(w, h int) int              { return w * h * 2 }
func (*fmtRGBA5551) Check(d []byte, w, h int) error { return checkSize(d, w, h, 16) }

// RGBA5551 returns a format containing an 5-bit red, green and blue
// channels and 1-bit of alpha per pixel.
func RGBA5551() Format { return &fmtRGBA5551{} }

func init() {
	RegisterConverter(RGBA5551(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			dst, j := make([]byte, width*height*4), 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					rgba := (uint16(src[1]) << 8) | uint16(src[0])
					src = src[2:]
					r := uint8((float32(rgba&0xF800) / 0xF800) * 255)
					g := uint8((float32(rgba&0x07C0) / 0x07C0) * 255)
					b := uint8((float32(rgba&0x003E) / 0x003E) * 255)
					a := uint8((rgba & 0x0001) * 255)
					dst[j+0], dst[j+1], dst[j+2], dst[j+3] = r, g, b, a
					j += 4
				}
			}

			return dst, nil
		})
}
