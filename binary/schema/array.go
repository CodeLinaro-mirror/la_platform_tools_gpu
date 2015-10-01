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
	Alias     string      // The alias this array type was given, if present
	ValueType binary.Type // The value type stored in the array
	Size      uint32      // The fixed size of the array
}

// Slice is the Type descriptor for dynamically sized buffers of known type,
// encoded with a preceding count.
type Slice struct {
	Alias     string      // The alias this array type was given, if present
	ValueType binary.Type // The value type stored in the slice.
}

func (a *Array) Representation() string {
	return fmt.Sprintf("%r", a)
}

func (a *Array) String() string {
	return fmt.Sprint(a)
}

func (a *Array) Format(f fmt.State, c rune) {
	switch c {
	case 'z':
		fmt.Fprintf(f, "[%d]%z", a.Size, a.ValueType)
	case 'r':
		fmt.Fprintf(f, "[%d]%r", a.Size, a.ValueType)
	default:
		if a.Alias != "" {
			fmt.Fprint(f, a.Alias)
		} else {
			fmt.Fprintf(f, "[%d]%v", a.Size, a.ValueType)
		}
	}
}

func (a *Array) EncodeValue(e binary.Encoder, value interface{}) {
	v := value.([]interface{})
	for i := range v {
		a.ValueType.EncodeValue(e, v[i])
	}
}

func (a *Array) DecodeValue(d binary.Decoder) interface{} {
	v := make([]interface{}, a.Size)
	for i := range v {
		v[i] = a.ValueType.DecodeValue(d)
	}
	return v
}

func (s *Slice) Representation() string {
	return fmt.Sprintf("%r", s)
}

func (s *Slice) String() string {
	return fmt.Sprint(s)
}

func (s *Slice) Format(f fmt.State, c rune) {
	switch c {
	case 'z':
		fmt.Fprintf(f, "[]%z", s.ValueType)
	case 'r':
		fmt.Fprintf(f, "[]%r", s.ValueType)
	default:
		if s.Alias != "" {
			fmt.Fprint(f, s.Alias)
		} else {
			fmt.Fprintf(f, "[]%v", s.ValueType)
		}
	}
}

func (s *Slice) EncodeValue(e binary.Encoder, value interface{}) {
	v := value.([]interface{})
	e.Uint32(uint32(len(v)))
	for i := range v {
		s.ValueType.EncodeValue(e, v[i])
	}
}

func (s *Slice) DecodeValue(d binary.Decoder) interface{} {
	size := d.Uint32()
	v := make([]interface{}, size)
	for i := range v {
		v[i] = s.ValueType.DecodeValue(d)
	}
	return v
}
