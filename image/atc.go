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

type fmtATC_RGB_AMD struct{ binary.Generate }

func (f *fmtATC_RGB_AMD) Key() interface{} { return *f }
func (*fmtATC_RGB_AMD) String() string     { return "ATC_RGB_AMD" }
func (*fmtATC_RGB_AMD) Size(w, h int) int  { return (max(w, 4) * max(h, 4)) / 2 }
func (*fmtATC_RGB_AMD) Check(d []byte, w, h int) error {
	return checkSize(d, max(w, 4), max(h, 4), 4)
}

type fmtATC_RGBA_EXPLICIT_ALPHA_AMD struct{ binary.Generate }

func (f *fmtATC_RGBA_EXPLICIT_ALPHA_AMD) Key() interface{} { return *f }
func (*fmtATC_RGBA_EXPLICIT_ALPHA_AMD) String() string     { return "ATC_RGBA_EXPLICIT_ALPHA_AMD" }
func (*fmtATC_RGBA_EXPLICIT_ALPHA_AMD) Size(w, h int) int  { return (max(w, 4) * max(h, 4)) }
func (*fmtATC_RGBA_EXPLICIT_ALPHA_AMD) Check(d []byte, w, h int) error {
	return checkSize(d, max(w, 4), max(h, 4), 8)
}

// ATC_RGB_AMD returns a format representing the texture compression format with
// the same name.
func ATC_RGB_AMD() Format { return &fmtATC_RGB_AMD{} }

// ATC_RGBA_EXPLICIT_ALPHA_AMD returns a format representing the texture
// compression format with the same name.
func ATC_RGBA_EXPLICIT_ALPHA_AMD() Format { return &fmtATC_RGBA_EXPLICIT_ALPHA_AMD{} }

func init() {
	RegisterConverter(ATC_RGB_AMD(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			dst, j := make([]byte, width*height*4), 0

			block_width := max(width/4, 1)
			block_height := max(height/4, 1)

			bs := binary.BitStream{Data: src}

			c := [4][3]uint32{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			}

			for by := 0; by < block_height; by++ {
				for bx := 0; bx < block_width; bx++ {
					c[0][2] = bs.Read(5) << 3
					c[0][1] = bs.Read(5) << 3
					c[0][0] = bs.Read(5) << 3
					alt := bs.ReadBit() != 0
					c[3][2] = bs.Read(5) << 3
					c[3][1] = bs.Read(6) << 2
					c[3][0] = bs.Read(5) << 3
					for i := 0; i < 3; i++ {
						if alt {
							c[2][i] = c[0][i]
							c[1][i] = c[0][i] - c[3][i]/4
							c[0][i] = 0
						} else {
							c[1][i] = (c[0][i]*2 + c[3][i]*1) / 3
							c[2][i] = (c[0][i]*1 + c[3][i]*2) / 3
						}
					}
					for y := by * 4; y < (by+1)*4; y++ {
						for x := bx * 4; x < (bx+1)*4; x++ {
							idx := bs.Read(2)
							if x < width && y < height {
								j = 4 * (y*width + x)
								dst[j+0] = uint8(c[idx][0])
								dst[j+1] = uint8(c[idx][1])
								dst[j+2] = uint8(c[idx][2])
								dst[j+3] = 255
							}
						}
					}
				}
			}

			return dst, nil
		})

	RegisterConverter(ATC_RGBA_EXPLICIT_ALPHA_AMD(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			dst, j := make([]byte, width*height*4), 0

			block_width := max(width/4, 1)
			block_height := max(height/4, 1)

			bs := binary.BitStream{Data: src}

			a := [16]uint8{
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
				0, 0, 0, 0,
			}

			c := [4][3]uint32{
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
				{0, 0, 0},
			}

			for by := 0; by < block_height; by++ {
				for bx := 0; bx < block_width; bx++ {
					for i := 0; i < 16; i++ {
						a[i] = uint8(bs.Read(4) << 4)
					}
					c[0][2] = bs.Read(5) << 3
					c[0][1] = bs.Read(5) << 3
					c[0][0] = bs.Read(5) << 3
					bs.ReadBit()
					c[3][2] = bs.Read(5) << 3
					c[3][1] = bs.Read(6) << 2
					c[3][0] = bs.Read(5) << 3
					for i := 0; i < 3; i++ {
						c[1][i] = (c[0][i]*2 + c[3][i]*1) / 3
						c[2][i] = (c[0][i]*1 + c[3][i]*2) / 3
					}
					p := 0
					for y := by * 4; y < (by+1)*4; y++ {
						for x := bx * 4; x < (bx+1)*4; x++ {
							idx := bs.Read(2)
							if x < width && y < height {
								j = 4 * (y*width + x)
								dst[j+0] = uint8(c[idx][0])
								dst[j+1] = uint8(c[idx][1])
								dst[j+2] = uint8(c[idx][2])
								dst[j+3] = a[p]
								p++
							}
						}
					}
				}
			}

			return dst, nil
		})
}
