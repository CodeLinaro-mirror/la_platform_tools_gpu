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
	binary.Generate
	Name string    // The simple name of the type.
	ID   binary.ID // The unique type identifier for the Object.
}

func (s *Struct) Basename() string {
	return s.Name
}

func (s *Struct) Typename() string {
	return s.Name
}

func (s *Struct) String() string {
	return s.Name
}

func (s *Struct) Encode(e binary.Encoder, value interface{}) error {
	return e.Value(value.(binary.Object))
}

func (s *Struct) Decode(d binary.Decoder) (interface{}, error) {
	class := d.Lookup(s.ID)
	if class == nil {
		return nil, fmt.Errorf("Unknown type id %v for %s", s.ID, s)
	}
	o := class.New()
	if o == nil {
		return nil, fmt.Errorf("Nil object built by class for %s : %T", s, class)
	}
	return o, d.Value(o)
}

func (s *Struct) Skip(d binary.Decoder) error {
	class := d.Lookup(s.ID)
	if class == nil {
		return fmt.Errorf("Unknown type id %v for %s", s.ID, s)
	}
	o := class.New()
	if o == nil {
		return fmt.Errorf("Nil object built by class for %s : %T", s, class)
	}
	return d.SkipValue(o)
}
