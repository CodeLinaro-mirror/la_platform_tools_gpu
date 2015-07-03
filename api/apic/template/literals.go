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

package template

// Based on http://graphics.stanford.edu/~seander/bithacks.html#ZerosOnRightMultLookup
var deBruijnBitPositionSequence = [32]int{
	0, 1, 28, 2, 29, 14, 24, 3, 30, 22, 20, 15, 25, 17, 4, 8,
	31, 27, 13, 23, 21, 19, 16, 7, 26, 12, 18, 6, 11, 5, 10, 9,
}

// Bitpos returns the position of the only non-zero bit of v.
// If none or more than one bits are non-zero in v, returns -1.
func (*Functions) Bitpos(v uint32) int {
	if v == 0 || v != v&-v {
		return -1
	}
	return deBruijnBitPositionSequence[v&-v*0x077CB531>>27]
}
