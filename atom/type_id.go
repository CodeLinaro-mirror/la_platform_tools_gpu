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

package atom

import "android.googlesource.com/platform/tools/gpu/binary"

// TypeId is an atom type identifier. Each implementation of the Atom interface
// must have a unique type idenitifier. Any changes to the binary format of an
// atom must result in a new type identifier to maintain binary compatability.
type TypeId uint16

func init() {
	Register(TypeInfo{Id: TypeIdEos, New: func() Atom { return &EOS{} }})
}

// Encode encodes the TypeId using the specified encoder.
func (i TypeId) Encode(e *binary.Encoder) error {
	return e.Uint16(uint16(i))
}

// Decode decodes the TypeId using the specified decoder.
func (i *TypeId) Decode(d *binary.Decoder) error {
	if val, err := d.Uint16(); err == nil {
		*i = TypeId(val)
		return nil
	} else {
		return err
	}
}
