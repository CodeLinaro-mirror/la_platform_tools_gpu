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

// List is a list of atoms.
type List []Atom

// WriteTo writes all atoms in the list to w, inserting EOS atoms after the last
// atom of each context.
func (l *List) WriteTo(w Writer) {
	// Find the last atom index for each context
	last := make(map[ContextID]int)
	for i, a := range *l {
		last[a.ContextID()] = i
	}

	// Write out the atoms, injecting EOS markers for each context.
	nextEosID := ID(len(*l))
	for i, a := range *l {
		w.Write(ID(i), a)
		ctx := a.ContextID()
		if last[ctx] == i {
			w.Write(nextEosID, &EOS{Context: ctx})
			nextEosID++
		}
	}
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
func (l *List) AddAt(a Atom, id ID) {
	*l = append(*l, nil)
	copy((*l)[id+1:], (*l)[id:])
	(*l)[id] = a
}

// Encode encodes the atom list using the specified encoder.
func (l *List) Encode(e binary.Encoder) error {
	if err := e.Uint32(uint32(len(*l))); err != nil {
		return err
	}
	for _, atom := range *l {
		if err := e.Uint16(uint16(atom.TypeID())); err != nil {
			return err
		}
		if err := atom.Encode(e); err != nil {
			return err
		}
	}
	return nil
}

// Encode decodes the atom list using the specified encoder.
func (l *List) Decode(d binary.Decoder) error {
	count, err := d.Uint32()
	if err != nil {
		*l = List{} // Clear the list
		return err
	}
	*l = make(List, count)
	for i := range *l {
		var typeID uint16
		if typeID, err = d.Uint16(); err != nil {
			return err
		}
		atom, err := New(TypeID(typeID))
		if err != nil {
			return err
		}
		if err := atom.Decode(d); err != nil {
			return err
		}
		(*l)[i] = atom
	}
	return nil
}
