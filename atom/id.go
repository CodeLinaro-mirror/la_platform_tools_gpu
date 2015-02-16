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

// Id is the index of an atom in an atom stream.
type Id uint64

// Encode encodes the Id using the specified encoder.
func (i Id) Encode(e *binary.Encoder) error {
	return e.Uint64(uint64(i))
}

// Decode encodes the Id using the specified decoder.
func (i *Id) Decode(d *binary.Decoder) error {
	if val, err := d.Uint64(); err == nil {
		*i = Id(val)
		return nil
	} else {
		return err
	}
}
