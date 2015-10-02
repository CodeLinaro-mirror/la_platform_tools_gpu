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

package gfxapi

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// Texture2D represents a two-dimensional texture resource.
type Texture2D struct {
	binary.Generate
	Levels []image.Info // The mip-map levels.
}

// Cubemap represents a cube-map texture resource.
type Cubemap struct {
	binary.Generate
	Levels []CubemapLevel // The mip-map levels.
}

// CubemapLevel represents a single mip-map level of a cube-map texture resource.
//
//         .........
//       .  +y   . :
//     .........   :
//     :       :+x :
//     :  +z   : .
//     :.......:
//
type CubemapLevel struct {
	binary.Generate
	NegativeX image.Info
	PositiveX image.Info
	NegativeY image.Info
	PositiveY image.Info
	NegativeZ image.Info
	PositiveZ image.Info
}

func (l *CubemapLevel) faces() [6]*image.Info {
	return [6]*image.Info{
		&l.NegativeX,
		&l.PositiveX,
		&l.NegativeY,
		&l.PositiveY,
		&l.NegativeZ,
		&l.PositiveZ,
	}
}

type imageMatcher struct {
	best          *image.Info
	score         uint32
	width, height uint32
}

func (m *imageMatcher) consider(i image.Info) {
	if m.best == nil {
		m.score = 0xffffffff
	}
	dw, dh := i.Width-m.width, i.Height-m.height
	score := dw*dw + dh*dh
	if m.score > score {
		m.score = score
		m.best = &i
	}
}

// Thumbnail returns the image that most closely matches the desired size.
func (t *Texture2D) Thumbnail(w, h uint32, d database.Database, l log.Logger) (*image.Info, error) {
	m := imageMatcher{width: w, height: h}
	for _, l := range t.Levels {
		m.consider(l)
	}

	return m.best, nil
}

// Convert returns this Texture2D with each mip-level converted to the requested format.
func (t *Texture2D) Convert(p *path.As, d database.Database, l log.Logger) (interface{}, error) {
	out := &Texture2D{
		Levels: make([]image.Info, len(t.Levels)),
	}
	for i, m := range t.Levels {
		if obj, err := m.Convert(p, d, l); err == nil {
			out.Levels[i] = *obj.(*image.Info)
		} else {
			return nil, err
		}
	}
	return out, nil
}

// Thumbnail returns the image that most closely matches the desired size.
func (t *Cubemap) Thumbnail(w, h uint32, d database.Database, l log.Logger) (*image.Info, error) {
	m := imageMatcher{width: w, height: h}

	for _, l := range t.Levels {
		m.consider(l.NegativeX)
		m.consider(l.PositiveX)
		m.consider(l.NegativeY)
		m.consider(l.PositiveY)
		m.consider(l.NegativeZ)
		m.consider(l.PositiveZ)
	}

	return m.best, nil
}

// Convert returns this Cubemap with each mip-level face converted to the requested format.
func (t *Cubemap) Convert(p *path.As, d database.Database, l log.Logger) (interface{}, error) {
	out := &Cubemap{
		Levels: make([]CubemapLevel, len(t.Levels)),
	}
	for i, m := range t.Levels {
		dst, src := out.Levels[i].faces(), m.faces()
		for j := range src {
			if obj, err := src[j].Convert(p, d, l); err == nil {
				*dst[j] = *obj.(*image.Info)
			} else {
				return nil, err
			}
		}
	}
	return out, nil
}
