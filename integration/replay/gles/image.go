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
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"testing"
)

const referenceImageDir = "reference"

// storeReferenceImage replaces the reference image with img.
func storeReferenceImage(t *testing.T, name string, img image.Image) {
	data := &bytes.Buffer{}
	if err := png.Encode(data, img); err != nil {
		t.Errorf("Failed to encode reference image %s: %v", name, err)
		return
	}
	path := filepath.Join(referenceImageDir, name+".png")
	err := ioutil.WriteFile(path, data.Bytes(), 0666)
	if err != nil {
		t.Errorf("Failed to store reference image %s: %v", name, err)
	}
}

// loadReferenceImage loads the reference image with the specified name.
func loadReferenceImage(t *testing.T, name string) image.Image {
	path := filepath.Join(referenceImageDir, name+".png")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Errorf("Failed to load reference image %s: %v", name, err)
		return nil
	}
	img, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		t.Errorf("Failed to decode reference image %s: %v", name, err)
		return nil
	}
	return img
}

func sqr(f float64) float64 { return f * f }

// compareImages returns the normalized square error between the two images.
// A return value of 0 denotes identical images, a return value of 1 denotes
// a complete mismatch (black vs white).
func compareImages(t *testing.T, expected, got image.Image) float64 {
	s := expected.Bounds().Size()
	if s != got.Bounds().Size() {
		t.Errorf("Image dimensions are not as expected. Expected: %v, Got: %v", s, got.Bounds().Size())
		return 1
	}

	sqrErr := float64(0)
	for y := 0; y < s.Y; y++ {
		for x := 0; x < s.X; x++ {
			e := expected.At(x, y)
			g := got.At(x, y)
			er, eg, eb, ea := e.RGBA()
			gr, gg, gb, ga := g.RGBA()
			sqrErr += sqr((float64(er) - float64(gr)) / 0xffff)
			sqrErr += sqr((float64(eg) - float64(gg)) / 0xffff)
			sqrErr += sqr((float64(eb) - float64(gb)) / 0xffff)
			sqrErr += sqr((float64(ea) - float64(ga)) / 0xffff)
		}
	}
	return sqrErr / float64(4*s.Y*s.X)
}
