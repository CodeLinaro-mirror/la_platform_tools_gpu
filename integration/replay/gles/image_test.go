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

package gles

import (
	"image"
	"image/color"
	"math"
	"testing"
)

func TestCompareImages(t *testing.T) {
	fill := func(i *image.RGBA, c color.RGBA) {
		for p := 0; p < len(i.Pix); p += 4 {
			i.Pix[p+0] = c.R
			i.Pix[p+1] = c.G
			i.Pix[p+2] = c.B
			i.Pix[p+3] = c.A
		}
	}
	for _, test := range []struct {
		a, b color.RGBA
		err  float64
	}{
		{
			a:   color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff},
			b:   color.RGBA{R: 0x00, G: 0x00, B: 0x00, A: 0x00},
			err: 1.0,
		}, {
			a:   color.RGBA{R: 0xff, G: 0xff, B: 0x00, A: 0x00},
			b:   color.RGBA{R: 0x00, G: 0x00, B: 0xff, A: 0xff},
			err: 1.0,
		}, {
			a:   color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0x00},
			b:   color.RGBA{R: 0x00, G: 0xff, B: 0xff, A: 0xff},
			err: 0.5,
		}, {
			a:   color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0x00},
			b:   color.RGBA{R: 0xff, G: 0x00, B: 0xff, A: 0x00},
			err: 0.0,
		},
	} {
		imgA := image.NewRGBA(image.Rect(0, 0, 8, 8))
		imgB := image.NewRGBA(image.Rect(0, 0, 8, 8))
		fill(imgA, test.a)
		fill(imgB, test.b)
		if err := compareImages(t, imgA, imgB); math.Abs(err-test.err) > 0.0001 {
			t.Errorf("CompareImages of %v and %v gave error: %v, expected: %v",
				test.a, test.b, err, test.err)
		}
	}
}
