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

package transform

import (
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
)

type testAtom struct {
	ID        atom.ID
	Type      atom.TypeID
	Context   atom.ContextID
	AtomFlags atom.Flags
}

func (a testAtom) TypeID() atom.TypeID          { return a.Type }
func (a testAtom) ContextID() atom.ContextID    { return a.Context }
func (a testAtom) Flags() atom.Flags            { return a.AtomFlags }
func (a testAtom) Encode(*binary.Encoder) error { return nil }
func (a testAtom) Decode(*binary.Decoder) error { return nil }

type atomAtomIDList []atomAtomID

// list takes a mix of testAtoms and atomAtomIDs and returns a atomAtomIDList.
// testAtoms are transformed into atomAtomIDs by using the ID field as the atom
// id.
func list(atoms ...interface{}) atomAtomIDList {
	l := atomAtomIDList{}
	for _, a := range atoms {
		switch a := a.(type) {
		case testAtom:
			l = append(l, atomAtomID{a, a.ID})
		case atomAtomID:
			l = append(l, a)
		default:
			panic("list only accepts types testAtom or atomAtomID")
		}
	}
	return l
}

func (l *atomAtomIDList) Write(id atom.ID, a atom.Atom) {
	*l = append(*l, atomAtomID{a, id})
}

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

// checkTransform checks that transfomer emits the expected atoms given inputs.
func checkTransform(t *testing.T, transformer atom.Transformer, inputs, expected atomAtomIDList) {
	got := atomAtomIDList{}
	for _, in := range inputs {
		transformer.Transform(in.id, in.atom, &got)
	}

	matched := len(expected) == len(got)

	if matched {
		for i := range expected {
			e, g := expected[i], got[i]
			if e.id != g.id || !reflect.DeepEqual(g.atom, e.atom) {
				matched = false
				break
			}
		}
	}

	if !matched {
		c := max(len(expected), len(got))
		for i := 0; i < c; i++ {
			if i > len(got) {
				t.Errorf("(%d) Expected: %#v Got: <nothing>", i, expected[i])
				continue
			}
			e, g := expected[i], got[i]
			if e.id != g.id || !reflect.DeepEqual(g.atom, e.atom) {
				t.Errorf("(%d) Expected: %#v Got: %#v", i, e, g)
				continue
			}
			t.Logf("(%d) Matched: %#v", i, g)
		}
	}
}
