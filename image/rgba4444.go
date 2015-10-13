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

type fmtRGBA4444 struct{ binary.Generate }

func (f *fmtRGBA4444) Key() interface{}             { return *f }
func (*fmtRGBA4444) String() string                 { return "RGB4444" }
func (*fmtRGBA4444) Size(w, h int) int              { return w * h * 2 }
func (*fmtRGBA4444) Check(d []byte, w, h int) error { return checkSize(d, w, h, 16) }

// RGBA4444 returns a format containing a 4-bit red, green, blue and alpha
// channel per-pixel.
func RGBA4444() Format { return &fmtRGBA4444{} }

func init() {
	RegisterConverter(RGBA4444(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			dst, j := make([]byte, width*height*4), 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					rgb := (uint16(src[1]) << 8) | uint16(src[0])
					src = src[2:]
					r := uint8((float32(rgb&0xF000) / 0xF000) * 255)
					g := uint8((float32(rgb&0x0F00) / 0x0F00) * 255)
					b := uint8((float32(rgb&0x00F0) / 0x00F0) * 255)
					a := uint8((float32(rgb&0x000F) / 0x000F) * 255)
					dst[j+0], dst[j+1], dst[j+2], dst[j+3] = r, g, b, a
					j += 4
				}
			}

			return dst, nil
		})
}
