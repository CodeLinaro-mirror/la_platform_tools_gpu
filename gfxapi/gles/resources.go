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

// IsResource returns true if this instance should be considered as a resource.
func (t *Texture) IsResource() bool {
	return t.ID != 0
}

// ResourceName returns the UI name for the resource.
func (t *Texture) ResourceName() string {
	return fmt.Sprintf("Texture<%d>", t.ID)
}

// ResourceType returns the type of this resource.
func (t *Texture) ResourceType() gfxapi.ResourceType {
	switch t.Kind {
	case TextureKind_TEXTURE2D:
		return gfxapi.TypeTexture2D
	case TextureKind_CUBEMAP:
		return gfxapi.TypeCubemap
	default:
		return gfxapi.TypeUnknown
	}
}

// ResourceData returns the resource data given the current state.
func (t *Texture) ResourceData(s *gfxapi.State, d database.Database, l log.Logger) (interface{}, error) {
	l = log.Enter(l, "Texture.Resource()")
	switch t.Kind {
	case TextureKind_UNDEFINED:
		return nil, nil

	case TextureKind_TEXTURE2D:
		levels := make([]image.Info, len(t.Texture2D))
		for i, level := range t.Texture2D {
			levels[i] = image.Info{
				Format: imageFormat(level.TexelFormat, level.TexelType),
				Width:  uint32(level.Width),
				Height: uint32(level.Height),
				Data:   &path.Blob{ID: level.Data.ResourceID(s, d, l)},
			}
		}
		return &gfxapi.Texture2D{Levels: levels}, nil

	case TextureKind_CUBEMAP:
		levels := make([]gfxapi.CubemapLevel, len(t.Cubemap))
		for i, level := range t.Cubemap {
			for j, face := range level.Faces {
				img := image.Info{
					Format: imageFormat(face.TexelFormat, face.TexelType),
					Width:  uint32(face.Width),
					Height: uint32(face.Height),
					Data:   &path.Blob{ID: face.Data.ResourceID(s, d, l)},
				}
				switch j {
				case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X:
					levels[i].NegativeX = img
				case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X:
					levels[i].PositiveX = img
				case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y:
					levels[i].NegativeY = img
				case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y:
					levels[i].PositiveY = img
				case GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z:
					levels[i].NegativeZ = img
				case GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z:
					levels[i].PositiveZ = img
				}
			}
		}
		return &gfxapi.Cubemap{Levels: levels}, nil

	default:
		return nil, fmt.Errorf("Unsupported texture kind %v", t.Kind)
	}
}
