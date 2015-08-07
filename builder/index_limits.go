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

package builder

import (
	"bytes"
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// IndexLimits holds the results of CalculateIndexLimits.
type IndexLimits struct {
	binary.Generate
	Min uint32 // Smallest index found in the index buffer.
	Max uint32 // Largest index found in the index buffer.
}

// CalcIndexLimits returns the lowest and highest index contained in the index
// buffer with identifier id. The buffer holds count elements, each of size
// bytes.
func CalcIndexLimits(id binary.ID, count int, size int, littleEndian bool, d database.Database, l log.Logger) (IndexLimits, error) {
	c := &calcIndexLimits{
		indexSize:    size,
		count:        count,
		littleEndian: littleEndian,
		data:         id,
	}
	r, err := database.Build(c, d, l)
	if err != nil {
		return IndexLimits{}, fmt.Errorf("Could not calculate index limits: %v", err)
	}
	return *r.(*IndexLimits), nil
}

type calcIndexLimits struct {
	binary.Generate
	indexSize    int
	count        int
	littleEndian bool
	data         binary.ID
}

func (c calcIndexLimits) BuildLazy(_ interface{}, d database.Database, l log.Logger) (interface{}, error) {
	min, max := ^uint32(0), uint32(0)
	data, err := d.Resolve(c.data, l)
	if err != nil {
		return nil, err
	}
	byteOrder := endian.Little
	if !c.littleEndian {
		byteOrder = endian.Big
	}
	r := endian.Reader(bytes.NewReader(data.([]byte)), byteOrder)

	var decode func() (uint32, error)
	switch c.indexSize {
	case 1:
		decode = func() (uint32, error) {
			v, err := r.Uint8()
			return uint32(v), err
		}

	case 2:
		decode = func() (uint32, error) {
			v, err := r.Uint16()
			return uint32(v), err
		}

	case 4:
		decode = r.Uint32

	default:
		return nil, fmt.Errorf("Unsupported index size %v", c.indexSize)
	}

	for i := 0; i < c.count; i++ {
		v, err := decode()
		if err != nil {
			return nil, err
		}
		if min > v {
			min = v
		}
		if max < v {
			max = v
		}
	}

	return &IndexLimits{Min: min, Max: max}, nil
}
