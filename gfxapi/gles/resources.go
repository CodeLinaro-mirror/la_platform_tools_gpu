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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// ResourceName returns the UI name for the resource.
func (t *Texture) ResourceName() string {
	return fmt.Sprintf("Texture<%d>", t.ID)
}

// ResourceType returns the type of this resource.
func (t *Texture) ResourceType() gfxapi.ResourceType {
	return gfxapi.TypeTexture
}

// ResourceData returns the resource data given the current state.
func (t *Texture) ResourceData(s *gfxapi.State, d database.Database, l log.Logger) (interface{}, error) {
	l = log.Enter(l, "Texture.Resource()")
	switch t.Kind {
	case TextureKind_TEXTURE2D:
		levels := make([]image.Info, len(t.Texture2D))
		for levelIdx, level := range t.Texture2D {
			levels[levelIdx] = image.Info{
				Format: imageFormat(level.Format),
				Width:  uint32(level.Width),
				Height: uint32(level.Height),
				Data:   &path.Blob{ID: level.Data.ResourceID(s, d, l)},
			}
		}
		return &gfxapi.Texture{Levels: levels}, nil

	default:
		return nil, fmt.Errorf("Unsupported texture kind %v", t.Kind)
	}
}
