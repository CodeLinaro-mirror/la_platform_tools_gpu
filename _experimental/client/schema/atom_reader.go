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
	"bytes"
	"fmt"
	"sort"

	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/service"
)

// DecodeAtoms decodes all atoms from the AtomStream stream.
func DecodeAtoms(stream service.AtomStream, schema service.Schema) ([]Atom, error) {
	d := cyclic.Decoder(vle.Reader(bytes.NewReader(stream.Data)))
	// Read all the atoms from the stream
	atoms := []Atom{}
	for i := 0; true; i++ {
		ty, err := d.Uint16()
		if err != nil {
			return nil, fmt.Errorf("(%d) Error reading atom's type: %v", i, err)
		}
		if ty == 0xffff { // EOS
			break
		}

		idx := sort.Search(len(schema.Atoms), func(i int) bool {
			return schema.Atoms[i].Type >= ty
		})
		if idx >= len(schema.Atoms) {
			return nil, fmt.Errorf("(%d) Atom type 0x%x not found in schema!", i, ty)
		}

		atomInfo := schema.Atoms[idx]
		if atomInfo.Type != ty {
			return nil, fmt.Errorf("(%d) Atom type 0x%x not found in schema!", i, ty)
		}

		atom, err := UnpackAtom(d, atomInfo)
		if err != nil {
			return nil, fmt.Errorf("(%d) Error unpacking atom: %v", i, err)
		}

		atoms = append(atoms, atom)
	}
	return atoms, nil
}
