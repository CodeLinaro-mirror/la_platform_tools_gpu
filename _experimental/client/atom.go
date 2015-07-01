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
		case *schema.Object:
			i = len(atoms)
			atoms = append(atoms, Atom{
				object: value,
			})
			a := &atoms[i]
			// Find the observations, if present
			for _, f := range a.object.Fields {
				if o, ok := f.(*atom.Observations); ok {
					a.observations = o
					break
				}
			}
			// Find the atom metadata, if present
			a.meta = atom.FindMetadata(a.object.Type)
			if a.meta == nil {
				return nil, fmt.Errorf("(%d) Atom was missing metadata", i)
			}
		case *objects.Terminator:
			break
		default:
			return nil, fmt.Errorf("(%d) Atom was not decoded by schema got, %T", i, value)
		}
	}
	return atoms, nil
}
