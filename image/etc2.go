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

type fmtETC2_RGB8 struct{ binary.Generate }

func (f *fmtETC2_RGB8) Key() interface{} { return *f }
func (*fmtETC2_RGB8) String() string     { return "ETC2_RGB8" }
func (*fmtETC2_RGB8) Size(w, h int) int {
	return (max(alignup(w, 4), 4) * max(alignup(h, 4), 4)) / 2
}
func (*fmtETC2_RGB8) Check(d []byte, w, h int) error {
	return checkSize(d, max(alignup(w, 4), 4), max(alignup(h, 4), 4), 4)
}

// ETC2_RGB8 returns a format representing the texture compression format
// with the same name.
func ETC2_RGB8() Format { return &fmtETC2_RGB8{} }

func init() {
	RegisterConverter(ETC2_RGB8(), RGBA(), func(src []byte, width, height int) ([]byte, error) {
		dst := make([]byte, width*height*4)

		block_width := max((width+3)/4, 1)
		block_height := max((height+3)/4, 1)

		const (
			R = 0
			G = 1
			B = 2
		)
		c := [4][3]int{}
		codes := [2][4]int{}
		modTbl0 := [8][4]int{ // mode 0 (ETC1)
			{2, 8, -2, -8},
			{5, 17, -5, -17},
			{9, 29, -9, -29},
			{13, 42, -13, -42},
			{18, 60, -18, -60},
			{24, 80, -24, -80},
			{33, 106, -33, -106},
			{47, 183, -47, -183},
		}
		modTbl1 := [8]int{3, 6, 11, 16, 23, 32, 41, 64}
		diffTbl := [8]int{0, 1, 2, 3, -4, -3, -2, -1}
		flipTbl := [2][16]int{
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

				mode := uint(0)
				for i := uint(0); i < 3; i++ {
					if diff == 0 {
						// ┏━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━┳━━┳━━┓
						// ┃    C₀     ┃    C₁     ┃    C₀     ┃    C₁     ┃    C₀     ┃    C₁     ┃   C₀   ┃   C₁   ┃df┃fp┃
						// ┃    Red    ┃    Red    ┃   Green   ┃   Green   ┃    Blue   ┃   Blue    ┃  Table ┃  Table ┃  ┃  ┃
						// ┣━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━╋━━┯━━┯━━╋━━╋━━┫
						// ┃₆₃│₆₂│₆₁│₆₀┃₅₉│₅₈│₅₇│₅₆┃₅₅│₅₄│₅₃│₅₂┃₅₁│₅₀│₄₉│₄₈┃₄₇│₄₆│₄₅│₄₄┃₄₃│₄₂│₄₁│₄₀┃₃₉│₃₈│₃₇┃₃₆│₃₅│₃₄┃₃₃┃₃₂┃
						// ┖──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┸──┴──┴──┸──┸──┚
						a := (u64 >> (60 - i*8)) & 15
						b := (u64 >> (56 - i*8)) & 15
						c[0][i] = int((a << 4) | a)
						c[1][i] = int((b << 4) | b)
					} else {
						// ┏━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━┳━━━━━━━━┳━━━━━━━━┳━━┳━━┓
						// ┃      C₀      ┃ Delta  ┃      C₀      ┃ Delta  ┃      C₀      ┃ Delta  ┃   C₀   ┃   C₁   ┃df┃fp┃
						// ┃      Red     ┃ Red    ┃     Green    ┃ Green  ┃     Blue     ┃  Blue  ┃  Table ┃  Table ┃  ┃  ┃
						// ┣━━┯━━┯━━┯━━┯━━╋━━┯━━┯━━╋━━┯━━┯━━┯━━┯━━╋━━┯━━┯━━╋━━┯━━┯━━┯━━┯━━╋━━┯━━┯━━╋━━┯━━┯━━╋━━┯━━┯━━╋━━╋━━┫
						// ┃₆₃│₆₂│₆₁│₆₀│₅₉┃₅₈│₅₇│₅₆┃₅₅│₅₄│₅₃│₅₂│₅₁┃₅₀│₄₉│₄₈┃₄₇│₄₆│₄₅│₄₄│₄₃┃₄₂│₄₁│₄₀┃₃₉│₃₈│₃₇┃₃₆│₃₅│₃₄┃₃₃┃₃₂┃
						// ┖──┴──┴──┴──┴──┸──┴──┴──┸──┴──┴──┴──┴──┸──┴──┴──┸──┴──┴──┴──┴──┸──┴──┴──┸──┴──┴──┸──┴──┴──┸──┸──┚
						a := (u64 >> (59 - i*8)) & 31
						b := (u64 >> (56 - i*8)) & 7
						d := int(a) + diffTbl[b]
						if d < 0 || d > 31 {
							mode = i + 1
							break
						}
						c[0][i] = int((a << 3) | (a >> 2))
						c[1][i] = int((d << 3) | (d >> 2))
					}
				}

				switch mode {
				case 0: // ETC1
					codes[0] = modTbl0[(u64>>37)&7]
					codes[1] = modTbl0[(u64>>34)&7]

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
								dst[k+2] = toByte(base[2] + shift)
								dst[k+1] = toByte(base[1] + shift)
								dst[k+0] = toByte(base[0] + shift)
								dst[k+3] = 255
							}
							i++
						}
					}
				case 1: // T-mode
					// ┏━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┓
					// ┃        ┃  R₀ ┏━━┓     ┃    G₀     ┃    B₀     ┃    R₂     ┃    G₂     ┃    B₂     ┃     ┏━━┓  ┃
					// ┣━━┯━━┯━━╋━━┯━━╋━━╋━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━╋━━╋━━┫
					// ┃₆₃│₆₂│₆₁┃₆₀│₅₉┃₅₈┃₅₇│₅₆┃₅₅│₅₄│₅₃│₅₂┃₅₁│₅₀│₄₉│₄₈┃₄₇│₄₆│₄₅│₄₄┃₄₃│₄₂│₄₁│₄₀┃₃₉│₃₈│₃₇│₃₆┃₃₅│₃₄┃₃₃┃₃₂┃
					// ┖──┴──┴──┸──┴──┸──┸──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┸──┸──┚

					// Load colors
					c[0][R] = int(expand4to8(((u64 >> 57) & 12) | (u64>>56)&3))
					c[0][G] = int(expand4to8(u64 >> 52))
					c[0][B] = int(expand4to8(u64 >> 48))
					c[2][R] = int(expand4to8(u64 >> 44))
					c[2][G] = int(expand4to8(u64 >> 40))
					c[2][B] = int(expand4to8(u64 >> 36))

					// Load intensity modifier
					modIdx := ((u64 >> 33) & 6) | ((u64 >> 32) & 1)
					mod := modTbl1[modIdx]
					// Calculate C₁ and c₃
					for i := 0; i < 3; i++ {
						c[1][i] = c[2][i] + mod
						c[3][i] = c[2][i] - mod
					}
					// ┏━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┓
					// ┃  x₀ ┃  x₀ ┃  x₀ ┃  x₀ ┃  x₁ ┃  x₁ ┃  x₁ ┃  x₁ ┃  x₂ ┃  x₂ ┃  x₂ ┃  x₂ ┃  x₃ ┃  x₃ ┃  x₃ ┃  x₃ ┃
					// ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃
					// ┣━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━┫
					// ┃₃₁│₃₀┃₂₉│₂₈┃₂₇│₂₆┃₂₅│₂₄┃₂₃│₂₂┃₂₁│₂₀┃₁₉│₁₈┃₁₇│₁₆┃₁₅│₁₄┃₁₃│₁₂┃₁₁│₁₀┃ ₉│ ₈┃ ₇│ ₆┃ ₅│ ₄┃ ₃│ ₂┃ ₁│ ₀┃
					// ┖──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┚
					// Use 2-bit indices to lookup texel colors
					i := uint(0)
					for x := bx * 4; x < (bx+1)*4; x++ {
						for y := by * 4; y < (by+1)*4; y++ {
							if x < width && y < height {
								k := 4 * (y*width + x)
								idx := ((u64 >> i) & 1) | ((u64 >> (15 + i)) & 2)
								dst[k+0] = toByte(c[idx][0])
								dst[k+1] = toByte(c[idx][1])
								dst[k+2] = toByte(c[idx][2])
								dst[k+3] = 255
							}
							i++
						}
					}
				case 2: // H-mode
					// ┏━━┳━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━━━━┳━━━━━━━━┓
					// ┃  ┃    R₀     ┃   G₀   ┏━━━━━━━━┓  ┃  ┏━━┓   B₀   ┃    R₂     ┃    G₂     ┃    B₂     ┃md┏━━┓  ┃
					// ┣━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━╋━━┯━━┯━━╋━━╋━━╋━━╋━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━╋━━╋━━╋━━┫
					// ┃₆₃┃₆₂│₆₁│₆₀│₅₉┃₅₈│₅₇│₅₆┃₅₅│₅₄│₅₃┃₅₂┃₅₁┃₅₀┃₄₉│₄₈│₄₇┃₄₆│₄₅│₄₄│₄₃┃₄₂│₄₁│₄₀│₃₉┃₃₈│₃₇│₃₆│₃₅┃₃₄┃₃₃┃₃₂┃
					// ┖──┸──┴──┴──┴──┸──┴──┴──┸──┴──┴──┸──┸──┸──┸──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┴──┴──┴──┸──┸──┸──┚

					// Load colors
					c[0][R] = int(expand4to8(u64 >> 59))
					c[0][G] = int(expand4to8(((u64 >> 55) & 14) | ((u64 >> 52) & 1)))
					c[0][B] = int(expand4to8(((u64 >> 48) & 8) | ((u64 >> 47) & 7)))
					c[2][R] = int(expand4to8(u64 >> 43))
					c[2][G] = int(expand4to8(u64 >> 39))
					c[2][B] = int(expand4to8(u64 >> 35))

					// Load intensity modifier
					modIdx := ((u64 >> 32) & 4) | ((u64 >> 31) & 2)
					// LSB of modIdx is 1 if:
					if (c[0][R]<<16)+(c[0][G]<<8)+c[0][B] >= (c[2][R]<<16)+(c[2][G]<<8)+c[2][B] {
						modIdx++
					}
					mod := modTbl1[modIdx]
					// Calculate C₁ and c₃
					for i := 0; i < 3; i++ {
						c[0][i], c[1][i] = c[0][i]+mod, c[0][i]-mod
						c[2][i], c[3][i] = c[2][i]+mod, c[2][i]-mod
					}
					// ┏━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┳━━━━━┓
					// ┃  x₀ ┃  x₀ ┃  x₀ ┃  x₀ ┃  x₁ ┃  x₁ ┃  x₁ ┃  x₁ ┃  x₂ ┃  x₂ ┃  x₂ ┃  x₂ ┃  x₃ ┃  x₃ ┃  x₃ ┃  x₃ ┃
					// ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃  y₀ ┃  y₁ ┃  y₂ ┃  y₃ ┃
					// ┣━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━╋━━┯━━┫
					// ┃₃₁│₃₀┃₂₉│₂₈┃₂₇│₂₆┃₂₅│₂₄┃₂₃│₂₂┃₂₁│₂₀┃₁₉│₁₈┃₁₇│₁₆┃₁₅│₁₄┃₁₃│₁₂┃₁₁│₁₀┃ ₉│ ₈┃ ₇│ ₆┃ ₅│ ₄┃ ₃│ ₂┃ ₁│ ₀┃
					// ┖──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┸──┴──┚
					// Use 2-bit indices to lookup texel colors
					i := uint(0)
					for x := bx * 4; x < (bx+1)*4; x++ {
						for y := by * 4; y < (by+1)*4; y++ {
							if x < width && y < height {
								k := 4 * (y*width + x)
								idx := ((u64 >> i) & 1) | ((u64 >> (15 + i)) & 2)
								dst[k+0] = toByte(c[idx][0])
								dst[k+1] = toByte(c[idx][1])
								dst[k+2] = toByte(c[idx][2])
								dst[k+3] = 255
							}
							i++
						}
					}
				case 3: // planar-mode
					// ┏━━┳━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━┓
					// ┃  ┃    R₀           ┃  ┏━━┓           G₀    ┃  ┏━━━━━━━━┓  B₀ ┏━━┓        ┃    R₁        ┏━━┓  ┃
					// ┣━━╋━━┯━━┯━━┯━━┯━━┯━━╋━━╋━━╋━━┯━━┯━━┯━━┯━━┯━━╋━━╋━━┯━━┯━━╋━━┯━━╋━━╋━━┯━━┯━━╋━━┯━━┯━━┯━━┯━━╋━━╋━━┫
					// ┃₆₃┃₆₂│₆₁│₆₀│₅₉│₅₈│₅₇┃₅₆┃₅₅┃₅₄│₅₃│₅₂│₅₁│₅₀│₄₉┃₄₈┃₄₇│₄₆│₄₅┃₄₄│₄₃┃₄₂┃₄₁│₄₀│₃₉┃₃₈│₃₇│₃₆│₃₅│₃₄┃₃₃┃₃₂┃
					// ┖──┸──┴──┴──┴──┴──┴──┸──┸──┸──┴──┴──┴──┴──┴──┸──┸──┴──┴──┸──┴──┸──┸──┴──┴──┸──┴──┴──┴──┴──┸──┸──┚
					//
					// ┏━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━━━━┳━━━━━━━━━━━━━━━━━┓
					// ┃         G₁         ┃        B₁       ┃        R₂       ┃         G₂         ┃       B₂        ┃
					// ┣━━┯━━┯━━┯━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━┯━━┯━━┯━━╋━━┯━━┯━━┯━━┯━━┯━━┫
					// ┃₃₁│₃₀│₂₉│₂₈│₂₇│₂₆│₂₅┃₂₄│₂₃│₂₂│₂₁│₂₀│₁₉┃₁₈│₁₇│₁₆│₁₅│₁₄│₁₃┃₁₂│₁₁│₁₀│ ₉│ ₈│ ₇│ ₆┃ ₅│ ₄│ ₃│ ₂│ ₁│ ₀┃
					// ┖──┴──┴──┴──┴──┴──┴──┸──┴──┴──┴──┴──┴──┸──┴──┴──┴──┴──┴──┸──┴──┴──┴──┴──┴──┴──┸──┴──┴──┴──┴──┴──┚

					// Load colors
					c[0][R] = int(expand6to8(u64 >> 57))
					c[0][G] = int(expand7to8(((u64 >> 50) & 64) | ((u64 >> 49) & 63)))
					c[0][B] = int(expand6to8(((u64 >> 43) & 32) | ((u64 >> 40) & 24) | ((u64 >> 39) & 7)))

					c[1][R] = int(expand6to8(((u64 >> 33) & 62) | ((u64 >> 32) & 1)))
					c[1][G] = int(expand7to8(u64 >> 25))
					c[1][B] = int(expand6to8(u64 >> 19))

					c[2][R] = int(expand6to8(u64 >> 13))
					c[2][G] = int(expand7to8(u64 >> 6))
					c[2][B] = int(expand6to8(u64))

					for dx := 0; dx < 4; dx++ {
						x := bx*4 + dx
						for dy := 0; dy < 4; dy++ {
							y := by*4 + dy
							if x < width && y < height {
								k := 4 * (y*width + x)
								dst[k+0] = toByte((dx*(c[1][R]-c[0][R]) + dy*(c[2][R]-c[0][R]) + 4*c[0][R] + 2) >> 2)
								dst[k+1] = toByte((dx*(c[1][G]-c[0][G]) + dy*(c[2][G]-c[0][G]) + 4*c[0][G] + 2) >> 2)
								dst[k+2] = toByte((dx*(c[1][B]-c[0][B]) + dy*(c[2][B]-c[0][B]) + 4*c[0][B] + 2) >> 2)
								dst[k+3] = 255
							}
						}
					}
				}
			}
		}

		return dst, nil
	})
}
