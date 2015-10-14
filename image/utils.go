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

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func toByte(i int) byte {
	if i < 0 {
		return 0
	}
	if i > 255 {
		return 255
	}
	return byte(i)
}

func expand4to8(v uint64) uint64 {
	v &= 0xF
	return (v << 4) | v
}

func expand5to8(v uint64) uint64 {
	v &= 0x1F
	return (v << 3) | (v >> 2)
}

func expand6to8(v uint64) uint64 {
	v &= 0x3F
	return (v << 2) | (v >> 4)
}

func expand7to8(v uint64) uint64 {
	v &= 0x7F
	return (v << 1) | (v >> 6)
}

func alignup(v int, alignment int) int {
	return alignment * ((v + (alignment - 1)) / alignment)
}
