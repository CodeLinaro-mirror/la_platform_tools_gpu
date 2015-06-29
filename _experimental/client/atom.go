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

package client

import (
	"bytes"
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/objects"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/service"
)

type Atom struct {
	object       *schema.Object
	meta         *atom.Metadata
	observations *atom.Observations
}

// Flags returns the flags of the atom.
func (a *Atom) Flags() atom.Flags {
	return a.meta.Flags
}

func (a *Atom) Observations() *atom.Observations {
	return a.observations
}

func (a *Atom) DisplayName() string {
	return a.object.Type.Display
}

func (a *Atom) Api() service.ApiId {
	return service.ApiId{ID: a.meta.Api}
}

func (a *Atom) DocumentationUrl() string {
	return a.meta.DocumentationUrl
}

func (a *Atom) FieldCount() int {
	return len(a.object.Fields)
}

func (a *Atom) Field(index int) (schema.Field, interface{}) {
	return a.object.Type.Fields[index], a.object.Fields[index]
}

func (a *Atom) SetField(index int, value interface{}) {
	a.object.Fields[index] = value
}

func (a Atom) Class() binary.Class {
	return a.object.Class()
}

type AtomClass struct {
	base         *schema.Class
	meta         *atom.Metadata
	observations int
}

var observationsId = (*atom.Observations)(nil).Class().ID()

func NewAtomClass(base *schema.Class, meta *atom.Metadata) *AtomClass {
	class := &AtomClass{base: base, meta: meta, observations: -1}
	// Find the observations, if present
	for i, f := range base.Fields {
		if s, ok := f.Type.(*schema.Struct); ok {
			if s.ID == observationsId {
				class.observations = i
				break
			}
		}
	}
	return class
}

func (c *AtomClass) ID() binary.ID {
	return c.base.ID()
}

func (c *AtomClass) New() binary.Object {
	return &Atom{object: c.base.New().(*schema.Object)}
}

func (c *AtomClass) Encode(e binary.Encoder, object binary.Object) error {
	a := object.(*Atom)
	return c.base.Encode(e, a.object)
}

func (c *AtomClass) Decode(d binary.Decoder) (binary.Object, error) {
	a := Atom{}
	o, err := c.base.Decode(d)
	if err != nil {
		return a, err
	}
	a.object = o.(*schema.Object)
	a.meta = c.meta
	if c.observations >= 0 {
		if c.observations >= len(a.object.Fields) {
			return a, fmt.Errorf("Missing Observations field in %s", c.base.Name)
		}
		value := a.object.Fields[c.observations]
		if observations, ok := value.(*atom.Observations); !ok {
			return a, fmt.Errorf("Observations field is of type %T in %s", value, c.base.Name)
		} else {
			a.observations = observations
		}
	}
	return a, nil
}

func (c *AtomClass) DecodeTo(d binary.Decoder, object binary.Object) error {
	return c.base.DecodeTo(d, object)
}

func (c *AtomClass) Skip(d binary.Decoder) error { return c.base.Skip(d) }

// DecodeAtoms decodes all atoms from the AtomStream stream.
func (c *ApplicationContext) DecodeAtoms(stream service.AtomStream) ([]Atom, error) {
	d := cyclic.Decoder(vle.Reader(bytes.NewReader(stream.Data)))
	d.Namespace = c.namespace
	// Read all the atoms from the stream
	atoms := []Atom{}
	for i := 0; true; i++ {
		value, err := d.Object()
		if err != nil {
			if err == io.EOF {
				break
			}
			return nil, fmt.Errorf("(%d) Error decoding atom: %v", i, err)
		}
		switch value := value.(type) {
		case Atom:
			atoms = append(atoms, value)
		case *objects.Terminator:
			break
		default:
			return nil, fmt.Errorf("(%d) Atom was not decoded by schema got, %T", i, value)
		}
	}
	return atoms, nil
}
