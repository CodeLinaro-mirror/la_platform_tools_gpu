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

import (
	"math"
	"testing"

	"android.googlesource.com/platform/tools/gpu/service/path"
)

var _ = path.Converter((*Info)(nil))

func TestDifference(t *testing.T) {
	fill := func(w, h uint32, r, g, b, a byte) *Image {
		data := make([]byte, w*h*4)
		for p := 0; p < len(data); p += 4 {
			data[p+0] = r
			data[p+1] = g
			data[p+2] = b
			data[p+3] = a
		}
		return &Image{
			Width:  w,
			Height: h,
			Data:   data,
			Format: RGBA(),
		}
	}
	for _, test := range []struct {
		name string
		a, b *Image
		diff float64
	}{
		{
			name: "white vs black",
			a:    fill(8, 8, 0xff, 0xff, 0xff, 0xff),
			b:    fill(8, 8, 0x00, 0x00, 0x00, 0x00),
			diff: 1.0,
		}, {
			name: "transparent-yellow vs blue",
			a:    fill(8, 8, 0xff, 0xff, 0x00, 0x00),
			b:    fill(8, 8, 0x00, 0x00, 0xff, 0xff),
			diff: 1.0,
		}, {
			name: "transparent-white vs cyan",
			a:    fill(8, 8, 0xff, 0xff, 0xff, 0x00),
			b:    fill(8, 8, 0x00, 0xff, 0xff, 0xff),
			diff: 0.5,
		}, {
			name: "transparent-purple vs transparent-purple",
			a:    fill(8, 8, 0xff, 0x00, 0xff, 0x00),
			b:    fill(8, 8, 0xff, 0x00, 0xff, 0x00),
			diff: 0.0,
		},
	} {
		diff, err := Difference(test.a, test.b)
		if err != nil {
			t.Errorf("Difference of %v returned error: %v", test.name, err)
			continue
		}

		if math.Abs(diff-test.diff) > 0.0000001 {
			t.Errorf("Difference of %v gave value: %v, expected: %v",
				test.name, diff, test.diff)
		}
	}
}
