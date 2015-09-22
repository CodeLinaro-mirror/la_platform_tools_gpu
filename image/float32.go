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
	"math"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
)

type fmtFloat32 struct{ binary.Generate }

func (f *fmtFloat32) Key() interface{}             { return *f }
func (*fmtFloat32) String() string                 { return "Float32" }
func (*fmtFloat32) Size(w, h int) int              { return w * h * 4 }
func (*fmtFloat32) Check(d []byte, w, h int) error { return checkSize(d, w, h, 32) }

// Float32 returns a format containing a single float channel per pixel.
func Float32() Format { return &fmtFloat32{} }

func init() {
	RegisterConverter(Float32(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			r := endian.Reader(bytes.NewBuffer(src), endian.Little)
			dst, i, j := make([]byte, width*height*4), 0, 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					depth, _ := r.Float32()
					d := 0.01 / (1.0 - float64(depth))
					dst[j+0] = byte(math.Cos(d+math.Pi*2.0*0.000)*127.0 + 128.0)
					dst[j+1] = byte(math.Cos(d+math.Pi*2.0*0.333)*127.0 + 128.0)
					dst[j+2] = byte(math.Cos(d+math.Pi*2.0*0.666)*127.0 + 128.0)
					dst[j+3] = byte(0xFF)
					i += 4
					j += 4
				}
			}
			return dst, nil
		})
}
