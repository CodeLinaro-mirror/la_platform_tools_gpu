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

package client

import (
	"image"
	"image/color"

	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

// NewColorTexture returns a gxui.Texture from the rgba-8888 data.
func NewColorTexture(driver gxui.Driver, width, height int, rgba []byte) gxui.Texture {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	img.Pix = rgba
	tex := driver.CreateTexture(img, 1)
	tex.SetFlipY(true)
	return tex
}

// NewDepthTexture returns a gxui.Texture from the depth data.
func NewDepthTexture(driver gxui.Driver, width, height int, depths []byte) gxui.Texture {
	img := image.NewRGBA(image.Rect(0, 0, width, height))
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			r, g, b, a := float32(depths[0]), float32(depths[1])/255.0, float32(depths[2])/65025.0, float32(depths[3])/160581375.0
			depth := (r + g + b + a) / 255.0
			depths = depths[4:]
			d := 0.01 / (1.0 - depth)
			c := color.RGBA{
				R: byte(math.Cosf(d+math.TwoPi*0.000)*127.0 + 128.0),
				G: byte(math.Cosf(d+math.TwoPi*0.333)*127.0 + 128.0),
				B: byte(math.Cosf(d+math.TwoPi*0.666)*127.0 + 128.0),
				A: byte(0xFF),
			}
			img.Set(x, y, c)
		}
	}
	tex := driver.CreateTexture(img, 1)
	tex.SetFlipY(true)
	return tex
}
