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

import (
	"bytes"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// List is a list of atoms.
type List []Atom

// WriteTo writes all atoms in the list and then an EOS atom to w.
func (l *List) WriteTo(w Writer) {
	for i, a := range *l {
		w.Write(Id(i), a)
	}
	w.Write(Id(len(*l)), &EOS{})
}

// Clone makes and returns a shallow copy of the atom list.
func (l *List) Clone() List {
	c := make(List, len(*l))
	copy(c, *l)
	return c
}

// Add adds a to the end of the atom list.
func (l *List) Add(a Atom) {
	*l = append(*l, a)
}

// Add adds a to the list before the atom at id.
func (l *List) AddAt(a Atom, id Id) {
	*l = append(*l, nil)
	copy((*l)[id+1:], (*l)[id:])
	(*l)[id] = a
}

// Encode encodes the atom list using the specified encoder.
func (l *List) Encode(e *binary.Encoder) error {
	atomBuf := &bytes.Buffer{}
	atomEnc := binary.NewEncoder(atomBuf)
	for _, atom := range *l {
		if err := atom.TypeId().Encode(atomEnc); err != nil {
			return err
		}

		if err := atom.Encode(atomEnc); err != nil {
			return err
		}

		atomData := atomBuf.Bytes()
		if err := e.Uint16(uint16(len(atomData) + 2)); err != nil {
			return err
		}

		if _, err := e.Write(atomData); err != nil {
			return err
		}
		atomBuf.Reset()
	}

	e.Uint16(0)
	return nil
}

// Encode decodes the atom list using the specified encoder.
func (l *List) Decode(d *binary.Decoder) error {
	*l = List{} // Clear the list

	for {
		size, err := d.Uint16()
		if err != nil {
			return err
		}

		if size == 0 {
			break
		}

		var typeId TypeId
		if err := typeId.Decode(d); err != nil {
			return err
		}

		atom, err := New(TypeId(typeId))
		if err != nil {
			return err
		}

		if err := atom.Decode(d); err != nil {
			return err
		}

		*l = append(*l, atom)
	}

	return nil
}
