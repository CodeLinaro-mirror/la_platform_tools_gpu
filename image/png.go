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
	"fmt"
	"image"
	"image/color"
	"image/png"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
)

type fmtPNG struct{ binary.Generate }

func (f *fmtPNG) Key() interface{}             { return *f }
func (*fmtPNG) String() string                 { return "PNG" }
func (*fmtPNG) Size(w, h int) int              { return -1 }
func (*fmtPNG) Check(d []byte, w, h int) error { return nil }

// PNG returns a format representing the the texture compression format with the
// same name.
func PNG() Format { return &fmtPNG{} }

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
	RegisterConverter(PNG(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			img, err := png.Decode(bytes.NewReader(src))
			if err != nil {
				return nil, err
			}
			if w := img.Bounds().Dx(); width != w {
				return nil, fmt.Errorf("PNG width was not as expected. Got: %v, expected: %v", w, width)
			}
			if h := img.Bounds().Dy(); height != h {
				return nil, fmt.Errorf("PNG width was not as expected. Got: %v, expected: %v", h, height)
			}

			var f Format
			buf := &bytes.Buffer{}
			e := endian.Writer(buf, endian.Little)

			switch img.ColorModel() {
			case color.RGBA64Model:
				return nil, fmt.Errorf("Unsupported color model 'RGBA64'")
			case color.RGBAModel, color.NRGBAModel:
				f = RGBA()
				for y := 0; y < height; y++ {
					for x := 0; x < width; x++ {
						r, g, b, a := img.At(x, y).RGBA()
						e.Uint8(uint8(r >> 8))
						e.Uint8(uint8(g >> 8))
						e.Uint8(uint8(b >> 8))
						e.Uint8(uint8(a >> 8))
					}
				}
			case color.NRGBA64Model:
				return nil, fmt.Errorf("Unsupported color model 'NRGBA64'")
			case color.AlphaModel:
				return nil, fmt.Errorf("Unsupported color model 'Alpha'")
			case color.Alpha16Model:
				return nil, fmt.Errorf("Unsupported color model 'Alpha16'")
			case color.GrayModel:
				return nil, fmt.Errorf("Unsupported color model 'Gray'")
			case color.Gray16Model:
				f = Float32()
				for y := 0; y < height; y++ {
					for x := 0; x < width; x++ {
						r, _, _, _ := img.At(x, y).RGBA()
						e.Float32(float32(r) / 0xffff)
					}
				}
			default:
				return nil, fmt.Errorf("Unrecognised color model")
			}
			return Convert(buf.Bytes(), width, height, f, RGBA())
		})
}
