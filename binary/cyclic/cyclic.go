// Copyright (C) 2014 The Android Open Source Project
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

package cyclic

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/registry"

	// Force the any package to be included so the boxers are registered.
	_ "android.googlesource.com/platform/tools/gpu/binary/any"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

// Encoder creates a binary.Encoder that writes to the supplied binary.Writer.
func Encoder(writer binary.Writer) binary.Encoder {
	return &encoder{
		Writer:   writer,
		entities: map[*binary.Entity]uint32{},
		objects:  map[binary.Object]uint32{},
		ids:      map[binary.ID]uint32{},
	}
}

// Decoder creates a binary.Decoder that reads from the provided binary.Reader.
func Decoder(reader binary.Reader) *decoder {
	return &decoder{
		Reader:    reader,
		Namespace: registry.Global,
		entities:  map[uint32]*binary.Entity{},
		objects:   map[uint32]binary.Object{},
		ids:       map[uint32]binary.ID{},
	}
}

type encoder struct {
	binary.Writer
	entities map[*binary.Entity]uint32
	objects  map[binary.Object]uint32
	ids      map[binary.ID]uint32
}

type decoder struct {
	binary.Reader
	Namespace *registry.Namespace
	entities  map[uint32]*binary.Entity
	objects   map[uint32]binary.Object
	ids       map[uint32]binary.ID
}

func (e *encoder) ID(id binary.ID) {
	if sid, found := e.ids[id]; found {
		e.Uint32(sid << 1)
	} else {
		sid = uint32(len(e.ids)) + 1
		e.ids[id] = sid
		e.Uint32((sid << 1) | 1)
		e.Data(id[:])
	}
}

func (d *decoder) ID() binary.ID {
	id := binary.ID{}
	v := d.Uint32()
	sid := v >> 1
	if (v & 1) != 0 {
		d.Data(id[:])
		d.ids[sid] = id
		return id
	}
	id, found := d.ids[sid]
	if !found {
		d.SetError(fmt.Errorf("Unknown id sid %v", sid))
	}
	return id
}

func (e *encoder) Entity(s *binary.Entity, compact bool) {
	if s == nil {
		e.Uint32(0)
		return
	}
	if sid, found := e.entities[s]; found {
		e.Uint32(sid << 1)
	} else {
		sid = uint32(len(e.entities)) + 1
		e.entities[s] = sid
		e.Uint32((sid << 1) | 1)
		schema.EncodeEntity(e, s, compact)
	}
}

func (d *decoder) Entity(compact bool) *binary.Entity {
	v := d.Uint32()
	if v == 0 {
		return nil
	}
	sid := v >> 1
	if (v & 1) != 0 {
		s := &binary.Entity{}
		d.entities[sid] = s
		schema.DecodeEntity(d, s, compact)
		return s
	}
	s, found := d.entities[sid]
	if !found {
		d.SetError(fmt.Errorf("Unknown entity sid %v", sid))
	}
	return s
}

func (e *encoder) Value(obj binary.Object) { obj.Class().Encode(e, obj) }
func (d *decoder) Value(obj binary.Object) { obj.Class().DecodeTo(d, obj) }
func (e *encoder) Variant(obj binary.Object) {
	if obj == nil {
		e.Entity(nil, true)
		return
	}
	class := obj.Class()
	e.Entity(class.Schema(), true)
	class.Encode(e, obj)
}

func (d *decoder) Variant() binary.Object {
	entity := d.Entity(true)
	if class := d.Lookup(entity); class == nil {
		d.SetError(fmt.Errorf("Unknown type %v", entity))
		return nil
	} else {
		return class.Decode(d)
	}
}

func (e *encoder) Object(obj binary.Object) {
	if obj == nil {
		e.Uint32(0)
		return
	}
	if sid, found := e.objects[obj]; found {
		e.Uint32(sid << 1)
	} else {
		sid = uint32(len(e.objects)) + 1
		e.objects[obj] = sid
		e.Uint32((sid << 1) | 1)
		e.Variant(obj)
	}
}

func (d *decoder) Object() binary.Object {
	v := d.Uint32()
	if v == 0 {
		return nil
	}
	sid := v >> 1
	decode := (v & 1) != 0
	o, found := d.objects[sid]
	switch {
	case found && decode:
		// TODO consider whether we want to reintroduce some skipping ability
		// just for this
		d.Variant()
	case decode:
		o = d.Variant()
		d.objects[sid] = o
	case found:
	default:
		d.SetError(fmt.Errorf("Unknown object sid %v", sid))
	}
	return o
}

func (d *decoder) Lookup(entity *binary.Entity) binary.Class {
	return d.Namespace.Lookup(entity.Signature())
}
