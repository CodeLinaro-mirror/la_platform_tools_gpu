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

type fmtETC1_RGB8_OES struct{}

func (fmtETC1_RGB8_OES) String() string { return "ETC1_RGB8_OES" }
func (fmtETC1_RGB8_OES) Check(d []byte, w, h int) error {
	return checkSize(d, max(w, 4), max(h, 4), 4)
}

// ETC1_RGB8_OES returns a format representing the texture compression format
// with the same name.
func ETC1_RGB8_OES() Format { return fmtETC1_RGB8_OES{} }

func init() {
	RegisterConverter(ETC1_RGB8_OES(), RGBA(), func(src []byte, width, height int) ([]byte, error) {
		dst := make([]byte, width*height*4)

		block_width := max(width/4, 1)
		block_height := max(height/4, 1)

		c := [2][3]int32{}
		codes := [2][4]int32{}
		modTbl := [8][4]int32{
			{2, 8, -2, -8},
			{5, 17, -5, -17},
			{9, 29, -9, -29},
			{13, 42, -13, -42},
			{18, 60, -18, -60},
			{24, 80, -24, -80},
			{33, 106, -33, -106},
			{47, 183, -47, -183},
		}
		diffTbl := [8]int32{0, 1, 2, 3, -4, -3, -2, -1}
		flipTbl := [2][16]int32{
			{0, 0, 0, 0, 0, 0, 0, 0, 1, 1, 1, 1, 1, 1, 1, 1},
			{0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1, 0, 0, 1, 1},
		}
		for by := 0; by < block_height; by++ {
			for bx := 0; bx < block_width; bx++ {
				u64 := uint64(src[0])
				u64 = (u64 << 8) | uint64(src[1])
				u64 = (u64 << 8) | uint64(src[2])
				u64 = (u64 << 8) | uint64(src[3])
				u64 = (u64 << 8) | uint64(src[4])
				u64 = (u64 << 8) | uint64(src[5])
				u64 = (u64 << 8) | uint64(src[6])
				u64 = (u64 << 8) | uint64(src[7])
				src = src[8:]
				flip := (u64 >> 32) & 1
				diff := (u64 >> 33) & 1

				codes[0] = modTbl[(u64>>37)&7]
				codes[1] = modTbl[(u64>>34)&7]

				for i := uint(0); i < 3; i++ {
					if diff == 0 {
						a := (u64 >> (44 + i*8)) & 15
						b := (u64 >> (40 + i*8)) & 15
						c[0][i] = int32((a << 4) | a)
						c[1][i] = int32((b << 4) | b)
					} else {
						a := (u64 >> (43 + i*8)) & 31
						b := (u64 >> (40 + i*8)) & 7
						d := int32(int32(a) + diffTbl[b])
						c[0][i] = int32((a << 3) | (a >> 2))
						c[1][i] = int32((d << 3) | (d >> 2))
					}
				}

				blockTbl := flipTbl[flip]

				i := uint(0)
				for x := bx * 4; x < (bx+1)*4; x++ {
					for y := by * 4; y < (by+1)*4; y++ {
						if x < width && y < height {
							block := blockTbl[i]
							base := c[block]
							k := 4 * (y*width + x)
							idx := ((u64 >> i) & 1) | ((u64 >> (15 + i)) & 2)
							shift := codes[block][idx]
							dst[k+0] = toByte(base[2] + shift)
							dst[k+1] = toByte(base[1] + shift)
							dst[k+2] = toByte(base[0] + shift)
							dst[k+3] = 255
							i++
						}
					}
				}
			}
		}

		return dst, nil
	})
}
