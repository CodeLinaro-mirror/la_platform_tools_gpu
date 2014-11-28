// Copyright (C) 2014 The Android Open Source Project
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

package binary

import "unsafe"

// Float16 represents a 16-bit floating point number, containing a single sign bit, 5 exponent bits
// and 10 fractional bits:
//
//    MSB                                                                         LSB
//   ╔════╦════╤════╤════╤════╤════╦════╤════╤════╤════╤════╤════╤════╤════╤════╤════╗
//   ║Sign║ E₄ │ E₃ │ E₂ │ E₁ │ E₀ ║ F₉ │ F₈ │ F₇ │ F₆ │ F₅ │ F₄ │ F₃ │ F₂ │ F₁ │ F₀ ║
//   ╚════╩════╧════╧════╧════╧════╩════╧════╧════╧════╧════╧════╧════╧════╧════╧════╝
//   Where E is the exponent bits and F is the fractional bits.
//
// This floating-point number is similar to IEEE 754-2008, but Float16 does not support NaNs nor
// ±Infs.
type Float16 uint16

// Float32 returns the Float16 value expanded to a float32
func (f Float16) Float32() float32 {
	u32 := expandF16ToF32(f)
	ptr := unsafe.Pointer(&u32)
	f32 := *(*float32)(ptr)
	return f32
}

func expandF16ToF32(in Float16) uint32 {
	sign := uint32(in&0x8000) << 16
	nonSign := uint32(in&0x7fff) << 13
	exp := uint32(in & 0x7c00)

	nonSign += 0x38000000

	if exp == 0 {
		nonSign = 0
	}

	return sign | nonSign
}
