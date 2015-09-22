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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// Image is a two-dimensional bitmap.
type Image struct {
	binary.Generate
	Format Format // The format of the image.
	Width  uint32 // The width of the image in pixels.
	Height uint32 // The height of the image in pixels.
	Data   []byte // The pixel data.
}

// Info describes a two-dimensional bitmap.
type Info struct {
	binary.Generate `java:"ImageInfo"`
	Format          Format     // The format of the image.
	Width           uint32     // The width of the image in pixels.
	Height          uint32     // The height of the image in pixels.
	Data            *path.Blob // The path to the pixel data of the image.
}

// Convert returns this image Info converted to format specified by p.
func (i Info) Convert(p *path.As, d database.Database, l log.Logger) (interface{}, error) {
	if f, ok := p.Type.(Format); ok {
		id, err := database.Store(&LazyConverter{
			Data:       i.Data.ID,
			Width:      i.Width,
			Height:     i.Height,
			FormatFrom: i.Format,
			FormatTo:   f,
		}, d, l)
		if err != nil {
			return nil, fmt.Errorf("Failed to convert ImageInfo at %s to format %T: %v",
				p.Path(), f, err)
		}
		return &Info{
			Format: f,
			Width:  i.Width,
			Height: i.Height,
			Data:   &path.Blob{ID: id},
		}, nil
	}
	return nil, fmt.Errorf("Cannot convert ImageInfo at %s to type %T",
		p.Path(), p.Type)
}

// Convert returns the Image converted to the format to.
func (i *Image) Convert(to Format) (*Image, error) {
	data, err := Convert(i.Data, int(i.Width), int(i.Height), i.Format, to)
	if err != nil {
		return nil, err
	}
	return &Image{Data: data, Width: i.Width, Height: i.Height, Format: to}, nil
}

// Difference returns the normalized square error between the two images.
// A return value of 0 denotes identical images, a return value of 1 denotes
// a complete mismatch (black vs white).
func Difference(a, b *Image) (float64, error) {
	if a.Width != b.Width || a.Height != b.Height {
		return 1, fmt.Errorf("Image dimensions are not identical. %dx%d vs %dx%d",
			a.Width, a.Height, b.Width, b.Height)
	}

	a, err := a.Convert(RGBA())
	if err != nil {
		return 1, err
	}
	b, err = b.Convert(RGBA())
	if err != nil {
		return 1, err
	}

	p, q := a.Data, b.Data
	sqrErr := float64(0)
	c := a.Width * a.Height * 4
	for i := uint32(0); i < c; i++ {
		sqrErr += sqr((float64(p[i]) - float64(q[i])) / 0xff)
	}
	return sqrErr / float64(c), nil
}

func sqr(f float64) float64 { return f * f }
