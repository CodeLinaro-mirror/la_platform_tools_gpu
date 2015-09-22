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
	"fmt"
	"image/color"
	"image/png"
	"io/ioutil"
	"path/filepath"
	"testing"

	goimg "image"

	"android.googlesource.com/platform/tools/gpu/binary/endian"
	gpuimg "android.googlesource.com/platform/tools/gpu/image"
)

const referenceImageDir = "reference"

// storeReferenceImage replaces the reference image with img.
func storeReferenceImage(t *testing.T, name string, img *gpuimg.Image) {
	data := &bytes.Buffer{}
	i, err := toGoImage(img)
	if err != nil {
		t.Fatalf("Failed to convert GPU image to Go image: %v", err)
	}
	if err := png.Encode(data, i); err != nil {
		t.Fatalf("Failed to encode reference image %s: %v", name, err)
	}
	path := filepath.Join(referenceImageDir, name+".png")
	if err := ioutil.WriteFile(path, data.Bytes(), 0666); err != nil {
		t.Fatalf("Failed to store reference image %s: %v", name, err)
	}
}

// loadReferenceImage loads the reference image with the specified name.
func loadReferenceImage(t *testing.T, name string) *gpuimg.Image {
	path := filepath.Join(referenceImageDir, name+".png")
	data, err := ioutil.ReadFile(path)
	if err != nil {
		t.Fatalf("Failed to load reference image %s: %v", name, err)
	}
	img, err := png.Decode(bytes.NewBuffer(data))
	if err != nil {
		t.Fatalf("Failed to decode reference image %s: %v", name, err)
	}
	out, err := toGPUImage(img)
	if err != nil {
		t.Fatalf("Failed to convert Go image to GPU image: %v", err)
	}
	return out
}

func toGoImage(in *gpuimg.Image) (goimg.Image, error) {
	rect := goimg.Rect(0, 0, int(in.Width), int(in.Height))
	switch in.Format.Key() {
	case gpuimg.RGBA().Key():
		out := goimg.NewNRGBA(rect)
		out.Pix = in.Data
		return out, nil

	case gpuimg.Float32().Key():
		buf := &bytes.Buffer{}
		src := endian.Reader(bytes.NewReader(in.Data), endian.Little)
		dst := endian.Writer(buf, endian.Big) // Yes. Big-endian. Really.

		for i, c := 0, int(in.Width*in.Height); i < c; i++ {
			v, _ := src.Float32()
			dst.Uint16(uint16(v * 0xffff))
		}

		out := goimg.NewGray16(rect)
		out.Pix = buf.Bytes()
		return out, nil

	default:
		return nil, fmt.Errorf("Unsupported format %v", in.Format)
	}
}

func toGPUImage(in goimg.Image) (*gpuimg.Image, error) {
	w, h := in.Bounds().Dx(), in.Bounds().Dy()
	out := &gpuimg.Image{Width: uint32(w), Height: uint32(h)}
	buf := &bytes.Buffer{}
	e := endian.Writer(buf, endian.Little)

	switch in.ColorModel() {
	case color.RGBA64Model:
		return nil, fmt.Errorf("Unsupported color model 'RGBA64'")
	case color.RGBAModel, color.NRGBAModel:
		out.Format = gpuimg.RGBA()
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r, g, b, a := in.At(x, y).RGBA()
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
		out.Format = gpuimg.Float32()
		for y := 0; y < h; y++ {
			for x := 0; x < w; x++ {
				r, _, _, _ := in.At(x, y).RGBA()
				e.Float32(float32(r) / 0xffff)
			}
		}
	default:
		return nil, fmt.Errorf("Unrecognised color model")
	}

	out.Data = buf.Bytes()
	return out, nil
}
