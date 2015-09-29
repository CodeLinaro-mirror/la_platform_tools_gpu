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
	Type binary.Type // The pointed to type.
}

func (p *Pointer) Representation() string {
	return fmt.Sprintf("*%s", p.Type.Representation())
}

func (p *Pointer) String() string {
	return fmt.Sprintf("*%s", p.Type)
}

func (p *Pointer) EncodeValue(e binary.Encoder, value interface{}) {
	if value != nil { // TODO proper nil test needed?
		e.Object(value.(binary.Object))
	} else {
		e.Object(nil)
	}
}

func (p *Pointer) DecodeValue(d binary.Decoder) interface{} {
	return d.Object()
}
