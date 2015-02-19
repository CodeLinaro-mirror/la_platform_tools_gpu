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

import "fmt"

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
