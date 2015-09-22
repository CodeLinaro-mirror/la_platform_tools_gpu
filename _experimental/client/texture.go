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

	"github.com/google/gxui"
)

// NewTexture returns a gxui.Texture from the rgba-8888 data.
func NewTexture(driver gxui.Driver, width, height int, rgba []byte) gxui.Texture {
	img := image.NewNRGBA(image.Rect(0, 0, width, height))
	img.Pix = rgba
	tex := driver.CreateTexture(img, 1)
	tex.SetFlipY(true)
	return tex
}
