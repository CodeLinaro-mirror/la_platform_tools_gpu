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
