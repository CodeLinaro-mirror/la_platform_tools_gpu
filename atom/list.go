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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/config"
)

// List is a list of atoms.
type List []Atom

// WriteTo writes all atoms in the list to w, terminating with a single EOS
// atom.
func (l *List) WriteTo(w Writer) {
	for i, a := range *l {
		w.Write(ID(i), a)
	}
}

// Clone makes and returns a shallow copy of the atom list.
func (l *List) Clone() List {
	c := make(List, len(*l))
	copy(c, *l)
	return c
}

// Add appends a to the end of the atom list, returning the id of the last added
// atom.
func (l *List) Add(a ...Atom) ID {
	*l = append(*l, a...)
	return ID(len(*l) - 1)
}

// Add adds a to the list before the atom at id.
func (l *List) AddAt(a Atom, id ID) {
	*l = append(*l, nil)
	copy((*l)[id+1:], (*l)[id:])
	(*l)[id] = a
}

// Encode encodes the atom list using the specified encoder.
func (l *List) Encode(e binary.Encoder) error {
	for _, atom := range *l {
		if err := e.Uint16(uint16(atom.TypeID())); err != nil {
			return err
		}
		if err := e.Value(atom); err != nil {
			return err
		}
	}
	if err := e.Uint16(uint16(TypeIDEos)); err != nil {
		return err
	}
	return nil
}

// Encode decodes the atom list using the specified encoder.
func (l *List) Decode(d binary.Decoder) error {
	*l = List{}
	if config.DebugAtom {
		fmt.Printf("atom.List.Decode:\n")
	}
	for true {
		if config.DebugAtom {
			fmt.Printf("(%d) ", len(*l))
		}
		typeID, err := d.Uint16()
		if err != nil {
			return err
		}
		if config.DebugAtom {
			fmt.Printf("type-id: 0x%x ", typeID)
		}
		if TypeID(typeID) == TypeIDEos {
			break
		}
		atom, err := New(TypeID(typeID))
		if err != nil {
			return err
		}
		if config.DebugAtom {
			fmt.Printf("type: %T ", atom)
		}
		if err := d.Value(atom); err != nil {
			if config.DebugAtom {
				fmt.Printf("-- errored: %v\n", err)
			}
			return err
		}
		if config.DebugAtom {
			fmt.Printf("-- decoded\n")
		}
		(*l) = append(*l, atom)
	}
	return nil
}
