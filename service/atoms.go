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

package service

import (
	"bytes"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
)

// NewAtomStream creates a fully-encoded AtomStream from the atom list and
// schema.
func NewAtomStream(list atom.List, schema SchemaId) (AtomStream, error) {
	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	if err := list.Encode(enc); err != nil {
		return AtomStream{}, err
	}

	return AtomStream{
		Data:   U8Array(buf.Bytes()),
		Schema: schema,
	}, nil
}

// List decodes and returns the AtomStream decoded to an atom list.
func (s AtomStream) List() (atom.List, error) {
	list := atom.List{}
	if err := list.Decode(binary.NewDecoder(bytes.NewBuffer(s.Data))); err != nil {
		return atom.List{}, err
	}
	return list, nil
}

// Pack packs the atom Group o into the RPC-friendly AtomGroup structure.
func (g *AtomGroup) Pack(o atom.Group) {
	g.Name = o.Name
	g.Range.Pack(o.Range)
	g.SubGroups = make(AtomGroupArray, len(o.SubGroups))
	for i := range g.SubGroups {
		g.SubGroups[i].Pack(o.SubGroups[i])
	}
}

// Unpack unpacks the RPC-friendly AtomGroup structure into the atom Group o.
func (g AtomGroup) Unpack(o *atom.Group) {
	o.Name = g.Name
	g.Range.Unpack(&o.Range)
	o.SubGroups = make(atom.GroupList, len(g.SubGroups))
	for i := range o.SubGroups {
		g.SubGroups[i].Unpack(&o.SubGroups[i])
	}
}

// Pack packs the atom Range o into the RPC-friendly AtomRange structure.
func (r *AtomRange) Pack(o atom.Range) {
	r.First = uint64(o.First())
	r.Count = uint64(o.Length())
}

// Unpack unpacks the RPC-friendly AtomRange structure into the atom Range o.
func (r AtomRange) Unpack(o *atom.Range) {
	(*o) = atom.Range{Start: atom.Id(r.First), End: atom.Id(r.First + r.Count)}
}
