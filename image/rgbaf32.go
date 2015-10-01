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

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
)

type fmtRGBAF32 struct{ binary.Generate }

func (f *fmtRGBAF32) Key() interface{}             { return *f }
func (*fmtRGBAF32) String() string                 { return "RGBAF32" }
func (*fmtRGBAF32) Size(w, h int) int              { return w * h * 4 * 4 }
func (*fmtRGBAF32) Check(d []byte, w, h int) error { return checkSize(d, w, h, 128) }

// RGBA32 returns a format containing a red, green, blue and alpha 32-bit
// floating point channel per pixel.
func RGBAF32() *fmtRGBAF32 { return &fmtRGBAF32{} }

func f32ToByte(f float32) byte {
	switch {
	case f < 0:
		return 0xff
	case f > 1:
		return 0xff
	default:
		return byte(f * 0xff)
	}
}

func init() {
	RegisterConverter(RGBAF32(), RGBA(),
		func(src []byte, width, height int) ([]byte, error) {
			r := endian.Reader(bytes.NewBuffer(src), endian.Little)
			dst, i := make([]byte, width*height*4), 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					dst[i+0] = f32ToByte(r.Float32())
					dst[i+1] = f32ToByte(r.Float32())
					dst[i+2] = f32ToByte(r.Float32())
					dst[i+3] = f32ToByte(r.Float32())
					i += 4
				}
			}
			return dst, nil
		})
	RegisterConverter(RGBA(), RGBAF32(),
		func(src []byte, width, height int) ([]byte, error) {
			dst := make([]byte, width*height*4*4)
			w, i := endian.Writer(bytes.NewBuffer(dst[:0]), endian.Little), 0
			for y := 0; y < height; y++ {
				for x := 0; x < width; x++ {
					w.Float32(float32(src[i+0]) / 0xff)
					w.Float32(float32(src[i+1]) / 0xff)
					w.Float32(float32(src[i+2]) / 0xff)
					w.Float32(float32(src[i+3]) / 0xff)
					i += 4
				}
			}
			return dst, nil
		})
}

type rgbaF32 struct {
	r, g, b, a float32
}

// Resize returns a RGBAF32 image resized from srcW x srcH to dstW x dstH.
// The algorithm uses pixel-pair averaging to down-sample (if required) the
// image to no greater than twice the width or height than the target
// dimensions, then uses a bilinear interpolator to calculate the final image
// at the requested size.
func (fmtRGBAF32) Resize(data []byte, srcW, srcH, dstW, dstH int) ([]byte, error) {
	r := endian.Reader(bytes.NewReader(data), endian.Little)
	bufA, bufB := make([]rgbaF32, srcW*srcH), make([]rgbaF32, srcW*srcH)
	for i := range bufA {
		bufA[i] = rgbaF32{r.Float32(), r.Float32(), r.Float32(), r.Float32()}
	}

	dst, src := bufB, bufA

	for dstW*2 <= srcW { // Horizontal 2x downsample
		newW := srcW / 2
		for y := 0; y < srcH; y++ {
			i := newW * y
			a, b := srcW*y, srcW*y+1
			for x := 0; x < srcW/2; x++ {
				dst[i].r = (src[a].r + src[b].r) * 0.5
				dst[i].g = (src[a].g + src[b].g) * 0.5
				dst[i].b = (src[a].b + src[b].b) * 0.5
				dst[i].a = (src[a].a + src[b].a) * 0.5
				i, a, b = i+1, a+2, b+2
			}
		}
		dst, src, srcW = src, dst, newW
	}

	for dstH*2 <= srcH { // Vertical 2x downsample
		newH := srcH / 2
		for y := 0; y < newH; y++ {
			i := srcW * y
			a, b := i*2, i*2+srcW
			for x := 0; x < srcW; x++ {
				dst[i].r = (src[a].r + src[b].r) * 0.5
				dst[i].g = (src[a].g + src[b].g) * 0.5
				dst[i].b = (src[a].b + src[b].b) * 0.5
				dst[i].a = (src[a].a + src[b].a) * 0.5
				i, a, b = i+1, a+1, b+1
			}
		}
		dst, src, srcH = src, dst, newH
	}

	out := make([]byte, dstW*dstH*4*4)
	w := endian.Writer(bytes.NewBuffer(out[:0]), endian.Little)
	if srcW == dstW && srcH == dstH {
		for i, c := 0, dstW*dstH; i < c; i++ {
			w.Float32(src[i].r)
			w.Float32(src[i].g)
			w.Float32(src[i].b)
			w.Float32(src[i].a)
		}
	} else {
		// bi-linear filtering
		sx := float32(srcW-1) / float32(dstW-1)
		sy := float32(srcH-1) / float32(dstH-1)
		for y := 0; y < dstH; y++ {
			fy := float32(y) * sy
			iy := int(fy)
			dy, y0, y1 := fy-float32(iy), iy, min(iy+1, srcH-1)
			for x := 0; x < dstW; x++ {
				fx := float32(x) * sx
				ix := int(fx)
				dx, x0, x1 := fx-float32(ix), ix, min(ix+1, srcW-1)

				a, b := src[x0+y0*srcW], src[x1+y0*srcW]
				c, d := src[x0+y1*srcW], src[x1+y1*srcW]

				p := rgbaF32{a.r + (b.r-a.r)*dx, a.g + (b.g-a.g)*dx, a.b + (b.b-a.b)*dx, a.a + (b.a-a.a)*dx}
				q := rgbaF32{c.r + (d.r-c.r)*dx, c.g + (d.g-c.g)*dx, c.b + (d.b-c.b)*dx, c.a + (d.a-c.a)*dx}
				r := rgbaF32{p.r + (q.r-p.r)*dy, p.g + (q.g-p.g)*dy, p.b + (q.b-p.b)*dy, p.a + (q.a-p.a)*dy}

				w.Float32(r.r)
				w.Float32(r.g)
				w.Float32(r.b)
				w.Float32(r.a)
			}
		}
	}
	return out, nil
}
