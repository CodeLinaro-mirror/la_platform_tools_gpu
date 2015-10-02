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

// Struct is the Type descriptor for an binary.Object typed value.
type Struct struct {
	Relative string         // The relative go import name of the type, only useful to codergen.
	Entity   *binary.Entity // The schema entity this is a field of.
}

func (s *Struct) Representation() string {
	return fmt.Sprintf("%r", s)
}

func (s *Struct) String() string {
	return fmt.Sprint(s)
}

// Format implements the fmt.Formatter interface
func (s *Struct) Format(f fmt.State, c rune) {
	switch c {
	case 'z': // Private format specifier, supports Entity.Signature
		fmt.Fprint(f, "$")
	case 'r': // Private format specifier, supports Type.Representation
		fmt.Fprint(f, s.Entity.Name())
	default:
		if s.Relative != "" {
			fmt.Fprint(f, s.Relative)
		} else {
			fmt.Fprint(f, s.Entity.Name())
		}
	}
}

func (s *Struct) EncodeValue(e binary.Encoder, value interface{}) {
	e.Value(value.(binary.Object))
}

func (s *Struct) DecodeValue(d binary.Decoder) interface{} {
	class := d.Lookup(s.Entity.TypeID)
	if class == nil {
		d.SetError(fmt.Errorf("Unknown type id %v for %s", s.Entity.TypeID, s))
	}
	o := class.New()
	if o == nil {
		d.SetError(fmt.Errorf("Nil object built by class for %s : %T", s, class))
	}
	d.Value(o)
	return o
}
