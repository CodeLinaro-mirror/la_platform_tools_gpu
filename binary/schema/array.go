// Copyright (C) 2014 The Android Open Source Project
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

package schema

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Array is the Type descriptor for fixed size buffers of known type.
type Array struct {
	binary.Generate
	Alias     string // The alias this array type was given, if present
	ValueType Type   // The value type stored in the array
	Size      uint32 // The fixed size of the array
}

// Slice is the Type descriptor for dynamically sized buffers of known type,
// encoded with a preceding count.
type Slice struct {
	binary.Generate
	Alias     string // The alias this array type was given, if present
	ValueType Type   // The value type stored in the slice.
}

func (a *Array) Basename() string {
	return fmt.Sprintf("[%d]%s", a.Size, a.ValueType.Basename())
}

func (a *Array) Typename() string {
	if a.Alias != "" {
		return a.Alias
	}
	return fmt.Sprintf("[%d]%s", a.Size, a.ValueType.Typename())
}

func (a *Array) String() string {
	return a.Typename()
}

func (a *Array) Encode(e binary.Encoder, value interface{}) error {
	v := value.([]interface{})
	for i := range v {
		a.ValueType.Encode(e, v[i])
	}
	return nil
}

func (a *Array) Decode(d binary.Decoder) (interface{}, error) {
	var err error
	v := make([]interface{}, a.Size)
	for i := range v {
		if v[i], err = a.ValueType.Decode(d); err != nil {
			return v, err
		}
	}
	return v, nil
}

func (s *Slice) Basename() string {
	return fmt.Sprintf("[]%s", s.ValueType.Basename())
}

func (s *Slice) Typename() string {
	if s.Alias != "" {
		return s.Alias
	}
	return fmt.Sprintf("[]%s", s.ValueType.Typename())
}

func (s *Slice) String() string {
	return s.Typename()
}

func (s *Slice) Encode(e binary.Encoder, value interface{}) error {
	v := value.([]interface{})
	if err := e.Uint32(uint32(len(v))); err != nil {
		return err
	}
	for i := range v {
		s.ValueType.Encode(e, v[i])
	}
	return nil
}

func (s *Slice) Decode(d binary.Decoder) (interface{}, error) {
	size, err := d.Uint32()
	if err != nil {
		return nil, err
	}
	v := make([]interface{}, size)
	for i := range v {
		if v[i], err = s.ValueType.Decode(d); err != nil {
			return v, err
		}
	}
	return v, nil
}
