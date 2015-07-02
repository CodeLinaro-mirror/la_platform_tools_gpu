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

// Pointer is the Type descriptor for pointers.
type Pointer struct {
	binary.Generate
	Type Type // The pointed to type.
}

func (p *Pointer) Basename() string {
	return fmt.Sprintf("*%s", p.Type.Basename())
}

func (p *Pointer) Typename() string {
	return fmt.Sprintf("*%s", p.Type.Typename())
}

func (p *Pointer) String() string {
	return p.Typename()
}

func (p *Pointer) Encode(e binary.Encoder, value interface{}) error {
	if value != nil { // TODO proper nil test needed?
		if err := e.Object(value.(binary.Object)); err != nil {
			return err
		}
	} else if err := e.Object(nil); err != nil {
		return err
	}
	return nil
}

func (p *Pointer) Decode(d binary.Decoder) (interface{}, error) {
	return d.Object()
}

func (p *Pointer) Skip(d binary.Decoder) error {
	_, err := d.SkipObject()
	return err
}
