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

// UnpackAtom unpacks and returns an Atom of the AtomInfo type from the decoder d.
func UnpackAtom(d binary.Decoder, i service.AtomInfo) (Atom, error) {
	a := Atom{Info: i}

	for _, p := range i.Parameters {
		arg, err := ReadType(p.Type, d)
		if err != nil {
			return Atom{}, err
		}

		a.Arguments = append(a.Arguments, arg)
	}

	return a, nil
}
