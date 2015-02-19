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
	"bytes"
	"image"
	"image/color"
	"image/png"
)

type fmtPNG struct{}

func (fmtPNG) String() string                 { return "PNG" }
func (fmtPNG) Check(d []byte, w, h int) error { return nil }

// PNG returns a format representing the the texture compression format with the
// same name.
func PNG() Format { return fmtPNG{} }

func init() {
	RegisterConverter(RGBA(), PNG(),
		func(src []byte, width, height int) ([]byte, error) {
			img := image.NewRGBA(image.Rect(0, 0, width, height))
			i := 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					img.Set(x, y, color.RGBA{src[i+0], src[i+1], src[i+2], src[i+3]})
					i += 4
				}
			}

			buffer := bytes.Buffer{}
			png.Encode(&buffer, img)
			return buffer.Bytes(), nil
		})
}
