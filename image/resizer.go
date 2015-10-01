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
)

// Resizer is the interface implemented by formats that can resize image data.
type Resizer interface {
	// Resize returns an image resized from srcW x srcH to dstW x dstH.
	Resize(data []byte, srcW, srcH, dstW, dstH int) ([]byte, error)
}

// LazyResizer is a lazy request to resize an image.
type LazyResizer struct {
	binary.Generate `java:"disable"`
	Data            binary.ID
	Format          Format
	SrcWidth        uint32
	SrcHeight       uint32
	DstWidth        uint32
	DstHeight       uint32
}

// BuildLazy returns the byte array holding the resized image for the
// LazyResizer request.
func (r *LazyResizer) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	data, err := database.Resolve(r.Data, d, l)
	if err != nil {
		return nil, err
	}

	resizer, ok := r.Format.(Resizer)
	if !ok {
		return nil, fmt.Errorf("Image format %v does not support resizing", r.Format)
	}

	return resizer.Resize(data.([]byte),
		int(r.SrcWidth), int(r.SrcHeight),
		int(r.DstWidth), int(r.DstHeight))
}
