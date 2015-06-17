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
//
// Check returns an error if the combination of data, image width and image
// height is invalid for the given format, otherwise Check returns nil.
type Format interface {
	binary.Object

	Check(data []byte, width, height int) error
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

func (*fmtRGBA) String() string                 { return "RGBA" }
func (*fmtRGBA) Check(d []byte, w, h int) error { return checkSize(d, w, h, 32) }

// RGBA returns a format containing an 8-bit red, green, blue and alpha channel
// per pixel.
func RGBA() Format { return &fmtRGBA{} }
