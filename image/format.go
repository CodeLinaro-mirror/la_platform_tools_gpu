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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Format is the interface for an image and/or pixel format.
type Format interface {
	binary.Object

	// Check returns an error if the combination of data, image width and image
	// height is invalid for the given format, otherwise Check returns nil.
	Check(data []byte, width, height int) error

	// Size returns the number of bytes required to hold an image of the specified
	// dimensions in this format. If the size varies based on the image data, then
	// Size returns -1.
	Size(width, height int) int

	// Key returns an object that can be used for equality-testing of the format
	// and can be used as a key in a map. Formats of the same type and parameters
	// will always return equal keys.
	// Formats can be deserialized into new objects so testing equality on the
	// Format object directly is not safe.
	Key() interface{}
}

// FormatCast is automatically called by the generated decoders.
func FormatCast(obj binary.Object) Format {
	return obj.(Format)
}

func checkSize(data []byte, width, height int, bpp int) error {
	expected := width * height * bpp / 8
	actual := len(data)
	if expected != actual {
		return fmt.Errorf("Image data size (0x%x) did not match expected (0x%x) for dimensions %dx%d\n",
			actual, expected, width, height)
	}
	return nil
}

type fmtRGBA struct{ binary.Generate }

func (f *fmtRGBA) Key() interface{}             { return *f }
func (*fmtRGBA) String() string                 { return "RGBA" }
func (*fmtRGBA) Size(w, h int) int              { return w * h * 4 }
func (*fmtRGBA) Check(d []byte, w, h int) error { return checkSize(d, w, h, 32) }

// RGBA returns a format containing an 8-bit red, green, blue and alpha channel
// per pixel.
func RGBA() *fmtRGBA { return &fmtRGBA{} }

// Resize returns a RGBA image resized from srcW x srcH to dstW x dstH.
func (f *fmtRGBA) Resize(data []byte, srcW, srcH, dstW, dstH int) ([]byte, error) {
	f32 := &fmtRGBAF32{}
	data, err := Convert(data, srcW, srcH, f, f32)
	if err != nil {
		return nil, err
	}
	data, err = f32.Resize(data, srcW, srcH, dstW, dstH)
	if err != nil {
		return nil, err
	}
	return Convert(data, dstW, dstH, f32, f)
}
