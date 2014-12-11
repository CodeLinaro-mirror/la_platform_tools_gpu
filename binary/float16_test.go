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

import "testing"

func TestFloat16To32(t *testing.T) {
	checks := []struct {
		f16 Float16
		f32 float32
	}{
		{0x0000, 0.0},
		{0x3c00, 1.0},
		{0x4000, 2.0},
		{0x4200, 3.0},
		{0x4400, 4.0},
		{0x4500, 5.0},
		{0x3555, 0.333251953125},
		{0xbc00, -1.0},
		{0xc000, -2.0},
		{0xc200, -3.0},
		{0xc400, -4.0},
		{0xc500, -5.0},
		{0xb555, -0.333251953125},
	}
	for _, c := range checks {
		expected, got := c.f32, c.f16.Float32()
		if expected != got {
			t.Errorf("Expansion of float16(0x%x) to float32 gave unexpected value. Expected: %v, got: %v", c.f16, expected, got)
		}
	}
}
