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

import "android.googlesource.com/platform/tools/gpu/binary"

type Constants []interface{}

type Int8Constants struct {
	binary.Generate
	Type   Type           // The type of the constant.
	Values []Int8Constant // The constant values
}

type Int8Constant struct {
	binary.Generate
	Name  string
	Value int8
}

func (s Int8Constants) Len() int           { return len(s.Values) }
func (s Int8Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Int8Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Int8Constant) Add(c *Constants, t Type) {
	var s *Int8Constants
	for _, e := range *c {
		if e, ok := e.(*Int8Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Int8Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}

type Uint8Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Uint8Constant // The constant values
}

type Uint8Constant struct {
	binary.Generate
	Name  string
	Value uint8
}

func (s Uint8Constants) Len() int           { return len(s.Values) }
func (s Uint8Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Uint8Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Uint8Constant) Add(c *Constants, t Type) {
	var s *Uint8Constants
	for _, e := range *c {
		if e, ok := e.(*Uint8Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Uint8Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}

type Int16Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Int16Constant // The constant values
}

type Int16Constant struct {
	binary.Generate
	Name  string
	Value int16
}

func (s Int16Constants) Len() int           { return len(s.Values) }
func (s Int16Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Int16Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Int16Constant) Add(c *Constants, t Type) {
	var s *Int16Constants
	for _, e := range *c {
		if e, ok := e.(*Int16Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Int16Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}

type Uint16Constants struct {
	binary.Generate
	Type   Type             // The type of the constant.
	Values []Uint16Constant // The constant values
}

type Uint16Constant struct {
	binary.Generate
	Name  string
	Value uint16
}

func (s Uint16Constants) Len() int           { return len(s.Values) }
func (s Uint16Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Uint16Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Uint16Constant) Add(c *Constants, t Type) {
	var s *Uint16Constants
	for _, e := range *c {
		if e, ok := e.(*Uint16Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Uint16Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}

type Int32Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Int32Constant // The constant values
}

type Int32Constant struct {
	binary.Generate
	Name  string
	Value int32
}

func (s Int32Constants) Len() int           { return len(s.Values) }
func (s Int32Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Int32Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Int32Constant) Add(c *Constants, t Type) {
	var s *Int32Constants
	for _, e := range *c {
		if e, ok := e.(*Int32Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Int32Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}

type Uint32Constants struct {
	binary.Generate
	Type   Type             // The type of the constant.
	Values []Uint32Constant // The constant values
}

type Uint32Constant struct {
	binary.Generate
	Name  string
	Value uint32
}

func (s Uint32Constants) Len() int           { return len(s.Values) }
func (s Uint32Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Uint32Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Uint32Constant) Add(c *Constants, t Type) {
	var s *Uint32Constants
	for _, e := range *c {
		if e, ok := e.(*Uint32Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Uint32Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}

type Int64Constants struct {
	binary.Generate
	Type   Type            // The type of the constant.
	Values []Int64Constant // The constant values
}

type Int64Constant struct {
	binary.Generate
	Name  string
	Value int64
}

func (s Int64Constants) Len() int           { return len(s.Values) }
func (s Int64Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Int64Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Int64Constant) Add(c *Constants, t Type) {
	var s *Int64Constants
	for _, e := range *c {
		if e, ok := e.(*Int64Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Int64Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}

type Uint64Constants struct {
	binary.Generate
	Type   Type             // The type of the constant.
	Values []Uint64Constant // The constant values
}

type Uint64Constant struct {
	binary.Generate
	Name  string
	Value uint64
}

func (s Uint64Constants) Len() int           { return len(s.Values) }
func (s Uint64Constants) Swap(i, j int)      { s.Values[i], s.Values[j] = s.Values[j], s.Values[i] }
func (s Uint64Constants) Less(i, j int) bool { return s.Values[i].Value < s.Values[j].Value }

func (v Uint64Constant) Add(c *Constants, t Type) {
	var s *Uint64Constants
	for _, e := range *c {
		if e, ok := e.(*Uint64Constants); ok {
			if e.Type.String() == t.String() {
				s = e
				break
			}
		}
	}
	if s == nil {
		s = &Uint64Constants{Type: t}
		*c = append(*c, s)
	}
	s.Values = append(s.Values, v)
}
