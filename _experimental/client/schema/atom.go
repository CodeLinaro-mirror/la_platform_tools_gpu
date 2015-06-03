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

package schema

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/service"
)

// Atom is a schema-typed Atom value.
type Atom struct {
	Info         service.AtomInfo
	Observations Observations
	Arguments    []interface{}
}

// Pack encodes the Atom to the encoder e.
// The Atom can be decoded using the Unpack method of the atom's AtomInfo.
func (a Atom) Pack(e binary.Encoder) error {
	for _, a := range a.Arguments {
		if err := WriteType(a, e); err != nil {
			return err
		}
	}

	return nil
}
