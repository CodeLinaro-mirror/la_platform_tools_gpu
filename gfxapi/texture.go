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
	"errors"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Texture represents a texture resource.
type Texture struct {
	binary.Generate
	Levels []image.Info // The mip-map levels.
}

// Thumbnail returns the image that most closely matches the desired size.
func (t *Texture) Thumbnail(w, h uint32, d database.Database, l log.Logger) (*image.Info, error) {
	if len(t.Levels) == 0 {
		return nil, errors.New("Texture has no levels")
	}

	var best *image.Info
	bestScore := uint32(0xffffffff)

	sqr := func(i uint32) uint32 { return i * i }

	for _, level := range t.Levels {
		score := sqr(level.Width-w) + sqr(level.Height-h)
		if bestScore > score {
			level := level
			best, bestScore = &level, score
		}
	}

	return best, nil
}
