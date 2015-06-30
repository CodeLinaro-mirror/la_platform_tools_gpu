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
	"android.googlesource.com/platform/tools/gpu/binary/pod"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
)

// Encoder creates a binary.Encoder that writes to the supplied binary.Writer.
func Encoder(writer binary.Writer) binary.Encoder {
	return &encoder{
		Writer:  writer,
		objects: map[interface{}]uint32{},
		ids:     map[binary.ID]uint32{},
	}
}

type ref struct {
	v interface{}
	o binary.Object
}

// Decoder creates a binary.Decoder that reads from the provided binary.Reader.
func Decoder(reader binary.Reader) *decoder {
	return &decoder{
		Reader:    reader,
		Namespace: registry.Global,
		objects:   map[uint32]ref{},
		ids:       map[uint32]binary.ID{},
	}
}

type encoder struct {
	binary.Writer
	objects map[interface{}]uint32
	ids     map[binary.ID]uint32
}

type decoder struct {
	binary.Reader
	Namespace *registry.Namespace
	objects   map[uint32]ref
	ids       map[uint32]binary.ID
}

func (e *encoder) ID(id binary.ID) error {
	if sid, found := e.ids[id]; found {
		return e.Uint32(sid << 1)
	} else {
		sid = uint32(len(e.ids)) + 1
		e.ids[id] = sid
		if err := e.Uint32((sid << 1) | 1); err != nil {
			return err
		}
		return e.Data(id[:])
	}
}

func (d *decoder) ID() (binary.ID, error) {
	id := binary.ID{}
	v, err := d.Uint32()
	if err != nil {
		return id, err
	}
	sid := v >> 1
	if (v & 1) != 0 {
		err := d.Data(id[:])
		d.ids[sid] = id
		return id, err
	}
	id, found := d.ids[sid]
	if !found {
		fmt.Errorf("Unknown id sid %v", sid)
	}
	return id, nil
}

func (d *decoder) SkipID() error {
	if v, err := d.Uint32(); err != nil {
		return err
	} else if (v & 1) != 0 {
		return d.Skip(binary.IDSize)
	}
	return nil
}

func (e *encoder) Value(obj binary.Object) error     { return obj.Class().Encode(e, obj) }
func (d *decoder) Value(obj binary.Object) error     { return obj.Class().DecodeTo(d, obj) }
func (d *decoder) SkipValue(obj binary.Object) error { return obj.Class().Skip(d) }

func (e *encoder) Variant(v interface{}) error {
	if v == nil {
		return e.ID(binary.ID{})
	}
	obj := pod.Wrap(v)
	if obj == nil {
		return binary.ErrNotEncodable{Value: v}
	}
	class := obj.Class()
	if err := e.ID(class.ID()); err != nil {
		return err
	}
	return class.Encode(e, obj)
}

func (d *decoder) variant() (interface{}, binary.Object, error) {
	if id, err := d.ID(); err != nil {
		return nil, nil, err
	} else if class := d.Namespace.Lookup(id); class == nil {
		return nil, nil, fmt.Errorf("Unknown type id %v", id)
	} else {
		obj, err := class.Decode(d)
		return pod.Unwrap(obj), obj, err
	}
}

func (d *decoder) Variant() (interface{}, error) {
	v, _, e := d.variant()
	return v, e
}

func (d *decoder) SkipVariant() (binary.ID, error) {
	if id, err := d.ID(); err != nil {
		return id, err
	} else if class := d.Namespace.Lookup(id); class == nil {
		return id, fmt.Errorf("Unknown type id %v", id)
	} else {
		return id, class.Skip(d)
	}
}

func (e *encoder) Object(v interface{}) error {
	if v == nil {
		return e.Uint32(0)
	}
	obj := pod.Wrap(v)
	if obj == nil {
		return binary.ErrNotEncodable{Value: v}
	}
	if sid, found := e.objects[v]; found {
		return e.Uint32(sid << 1)
	} else {
		sid = uint32(len(e.objects)) + 1
		e.objects[v] = sid
		if err := e.Uint32((sid << 1) | 1); err != nil {
			return err
		}
		return e.Variant(v)
	}
}

func (d *decoder) Object() (interface{}, error) {
	i, err := d.Uint32()
	if err != nil || i == 0 {
		return nil, err
	}
	sid := i >> 1
	decode := (i & 1) != 0
	r, found := d.objects[sid]
	switch {
	case found && decode:
		_, err := d.SkipVariant()
		return r.v, err
	case decode:
		v, o, e := d.variant()
		d.objects[sid] = ref{v, o}
		return v, e
	case found:
		return r.v, nil
	default:
		return nil, fmt.Errorf("Unknown object sid %v", sid)
	}
}

func (d *decoder) SkipObject() (binary.ID, error) {
	if i, err := d.Uint32(); err != nil {
		return binary.ID{}, err
	} else if (i & 1) == 0 {
		sid := i >> 1
		if sid == 0 {
			return binary.ID{}, nil
		} else if r, found := d.objects[sid]; !found {
			return binary.ID{}, fmt.Errorf("Unknown object sid %v", sid)
		} else {
			return r.o.Class().ID(), nil
		}
	}
	return d.SkipVariant()
}

func (d *decoder) Lookup(id binary.ID) binary.Class {
	return d.Namespace.Lookup(id)
}
