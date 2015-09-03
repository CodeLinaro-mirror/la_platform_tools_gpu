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

// Converter is used to convert the the image formed from the parameters data,
// width and height into another format. If the conversion succeeds then the
// converted image data is returned, otherwise an error is returned.
type Converter func(data []byte, width int, height int) ([]byte, error)

type srcDstFmt struct{ src, dst Format }

var registeredConverters = make(map[srcDstFmt]Converter)

// RegisterConverter registers the Converter for converting from src to dst
// formats. If a converter already exists for converting from src to dst, then
// this function panics.
func RegisterConverter(src, dst Format, c Converter) {
	key := srcDstFmt{src, dst}
	if _, found := registeredConverters[key]; found {
		panic(fmt.Errorf("Converter from %s to %s already registered", src, dst))
	}
	registeredConverters[key] = c
}

// Convert uses the registered Converters to convert the image formed from the
// parameters data, width and height from srcFmt to dstFmt. If the conversion
// succeeds then the converted image data is returned, otherwise an error is
// returned.
// If no direct converter has been registered to convert from srcFmt to dstFmt,
// then Convert may try converting via an intermediate format.
func Convert(data []byte, width int, height int, srcFmt Format, dstFmt Format) ([]byte, error) {
	if srcFmt == dstFmt {
		return data, nil // No conversion required.
	}

	if err := srcFmt.Check(data, width, height); err != nil {
		return nil, fmt.Errorf("Source data of format %s is invalid: %s", srcFmt, err)
	}

	if conv, found := registeredConverters[srcDstFmt{srcFmt, dstFmt}]; found {
		return conv(data, width, height)
	}

	// Try going via RGBA
	if convA, found := registeredConverters[srcDstFmt{srcFmt, RGBA()}]; found {
		if convB, found := registeredConverters[srcDstFmt{RGBA(), dstFmt}]; found {
			if data, err := convA(data, width, height); err != nil {
				return convB(data, width, height)
			}
		}
	}

	return nil, fmt.Errorf("No converter registered that can convert from format '%s' to '%s'\n",
		srcFmt, dstFmt)
}

// LazyConverter is a lazy request to decode a compressed texture.
type LazyConverter struct {
	binary.Generate
	Data       binary.ID
	Width      uint32
	Height     uint32
	FormatFrom Format
	FormatTo   Format

	// Number of bytes between lines in the source image.
	// If 0 then lines are contiguous.
	StrideFrom int
}

// BuildLazy returns the byte array holding the converted image for the
// ConvertLazy request.
func (r *LazyConverter) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	data, err := database.Resolve(r.Data, d, l)
	if err != nil {
		return nil, err
	}
	rowLength := r.FormatFrom.Size(int(r.Width), 1)
	if r.StrideFrom != 0 && r.StrideFrom != rowLength {
		// Remove any padding from the source image
		packed := make([]byte, r.FormatFrom.Size(int(r.Width), int(r.Height)))
		src, dst := data.([]byte), packed
		for y := 0; y < int(r.Height); y++ {
			copy(dst, src[:rowLength])
			dst, src = dst[rowLength:], src[r.StrideFrom:]
		}
		data = packed
	}

	data, err = Convert(data.([]byte), int(r.Width), int(r.Height), r.FormatFrom, r.FormatTo)
	if err != nil {
		return nil, err
	}

	return data, nil
}
