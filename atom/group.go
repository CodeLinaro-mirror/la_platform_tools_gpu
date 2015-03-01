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
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/interval"
)

// Group represents a named, contiguous span of atoms with support for sparse
// sub-groups. Groups are ideal for expressing nested hierarchies of atoms.
//
// Groups have the concept of items. An item is either an immediate sub-group,
// or an atom identifier that is within this group's span but outside of any
// sub-group.
//
// For example a Group spanning the atom identifier range [0 - 9] with two
// sub-groups spanning [2 - 4] and [7 - 8] would have the following tree of
// items:
//
//  Group
//    │
//    ├─── Item[0] ─── Atom[0]
//    │
//    ├─── Item[1] ─── Atom[1]
//    │
//    ├─── Item[2] ─── Sub-group 0
//    │                   │
//    │                   ├─── Item[0] ─── Atom[2]
//    │                   │
//    │                   ├─── Item[1] ─── Atom[3]
//    │                   │
//    │                   └─── Item[2] ─── Atom[4]
//    │
//    ├─── Item[3] ─── Atom[5]
//    │
//    ├─── Item[4] ─── Atom[6]
//    │
//    ├─── Item[5] ─── Sub-group 1
//    │                   │
//    │                   ├─── Item[0] ─── Atom[7]
//    │                   │
//    │                   └─── Item[1] ─── Atom[8]
//    │
//    └─── Item[6] ─── Atom[9]
//
type Group struct {
	Name      string    // Name of this group.
	Range     Range     // The range of atoms this group (and sub-groups) represents.
	SubGroups GroupList // All sub-groups of this group.
}

func (g Group) info(depth int) string {
	str := fmt.Sprintf("%s%s %s",
		strings.Repeat("  ", depth), g.Range.String(), g.Name)
	if len(g.SubGroups) > 0 {
		str += "\n" + g.SubGroups.info(depth+1)
	}
	return str
}

// String returns a string representing the group's name, range and sub-groups.
func (g Group) String() string {
	return g.info(0)
}

// Count returns the number of immediate items this group contains.
func (g Group) Count() uint64 {
	count := g.Range.Length()
	for _, sg := range g.SubGroups {
		count -= sg.Range.Length()
		count++ // For the group itself
	}
	return count
}

// Index returns the item with the specified index. If the item refers directly
// to an atom identifier then the atom identifier is returned in baseAtomID and
// subgroup is assigned nil.
// If the item is a sub-group then baseAtomID is returned as the lowest atom
// identifier found in the sub-group and subgroup is assigned the sub-group
// pointer.
func (g Group) Index(index uint64) (baseAtomID ID, subgroup *Group) {
	base := g.Range.First()
	for i := range g.SubGroups {
		sg := &g.SubGroups[i]
		if base+ID(index) < sg.Range.First() {
			break
		}
		index -= uint64(sg.Range.First() - base)
		if index == 0 {
			return sg.Range.First(), sg
		}
		index--
		base = sg.Range.Last() + 1
	}
	return base + ID(index), nil
}

// IndexOf returns the item index that refers directly to, or contains the given
// atom identifer.
func (g Group) IndexOf(atomID ID) uint64 {
	index := uint64(0)
	base := g.Range.First()
	for _, sg := range g.SubGroups {
		if atomID < sg.Range.First() {
			break
		}
		index += uint64(sg.Range.First() - base)
		base = sg.Range.Last() + 1
		if atomID <= sg.Range.Last() {
			return index
		}
		index++
	}
	return index + uint64(atomID-base)
}

// Insert adjusts the spans of this group and all subgroups for an insertion
// of count elements at atomID.
func (g *Group) Insert(atomID ID, count int) {
	s, e := g.Range.Range()
	if s >= atomID {
		s += ID(count)
	}
	if e > atomID {
		e += ID(count)
	}
	g.Range = Range{s, e}
	i := interval.Search(&g.SubGroups, func(test interval.U64Span) bool {
		return uint64(atomID) < test.End
	})
	for i < len(g.SubGroups) {
		sg := g.SubGroups[i]
		sg.Insert(atomID, count)
		g.SubGroups[i] = sg
		i++
	}
}

// Encode encodes the Group using the specified encoder.
func (g *Group) Encode(e *binary.Encoder) error {
	if err := e.String(g.Name); err != nil {
		return err
	}
	if err := g.Range.Encode(e); err != nil {
		return err
	}
	if err := g.SubGroups.Encode(e); err != nil {
		return err
	}
	return nil
}

// Decode decodes the Group using the specified decoder.
func (g *Group) Decode(d *binary.Decoder) error {
	if name, err := d.String(); err == nil {
		g.Name = name
	} else {
		return err
	}
	if err := g.Range.Decode(d); err != nil {
		return err
	}
	if err := g.SubGroups.Decode(d); err != nil {
		return err
	}
	return nil
}
