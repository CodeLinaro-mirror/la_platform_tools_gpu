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
	info   *service.AtomInfo
	object *schema.Object
}

func (a *Atom) IsEndOfFrame() bool {
	return a.info.IsEndOfFrame
}

func (a *Atom) IsDrawCall() bool {
	return a.info.IsDrawCall
}

func (a *Atom) DocumentationUrl() string {
	return a.info.DocumentationUrl
}

func (a *Atom) DisplayName() string {
	return a.object.Type.Display
}

func (a *Atom) Api() service.ApiId {
	return a.info.Api
}

func (a *Atom) Observations() *atom.Observations {
	for _, f := range a.object.Fields {
		if o, ok := f.(*atom.Observations); ok {
			return o
		}
	}
	return nil
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
func (c *ApplicationContext) DecodeAtoms(stream service.AtomStream, s service.Schema) ([]Atom, error) {
	atomMap := map[binary.Class]*service.AtomInfo{}
	for i := range s.Atoms {
		a := &s.Atoms[i]
		name := a.Name
		var match binary.Class
		c.schemaNamespace.Visit(func(class binary.Class) {
			if class.(*schema.Class).Display == name {
				match = class
			}
		})
		if match == nil {
			return nil, fmt.Errorf("No binary schema entry for %s", name)
		}
		atomMap[match] = a
	}

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
			info, _ := atomMap[value.Class()]
			atoms = append(atoms, Atom{
				info:   info,
				object: value,
			})
		case *objects.Terminator:
			break
		default:
			return nil, fmt.Errorf("Atom was not decoded by schema, %T", value)
		}
	}
	return atoms, nil
}
