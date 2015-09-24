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

package gapis

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Atom is a wrapper around a schema.Object that describes an atom object.
// Atom conforms to the atom.Atom interface, and provides a number of methods
// for accessing the parameters, return value and observations.
type Atom struct {
	object       *schema.Object
	class        *atomClass
	observations *atom.Observations
	flags        atom.Flags
}

var _ atom.Atom = &Atom{} // Verify that Atom implements atom.Atom.

// API returns the graphics API id this atom belongs to.
func (a *Atom) API() gfxapi.ID {
	return a.class.meta.API
}

// Flags returns the flags of the atom.
func (a *Atom) Flags() atom.Flags {
	return a.flags
}

// Observations returns all the memory observations made by the atom.
func (a *Atom) Observations() *atom.Observations {
	return a.observations
}

// Mutate is not supported by the Atom type, but is exposed in order to comform
// to the atom.Atom interface. Mutate will always return an error.
func (*Atom) Mutate(*gfxapi.State, database.Database, log.Logger) error {
	return fmt.Errorf("Mutate not implemented for client atoms")
}

// ParameterCount returns the number of parameters this atom accepts. This count
// does not include the return value.
func (a *Atom) ParameterCount() int {
	return len(a.class.parameters)
}

// Parameter returns the index'th parameter Field and value.
func (a *Atom) Parameter(index int) (schema.Field, interface{}) {
	index = a.class.parameters[index]
	return a.object.Type.Fields[index], a.object.Fields[index]
}

// SetParameter sets the atom's index'th parameter to the specified value.
func (a *Atom) SetParameter(index int, value interface{}) {
	index = a.class.parameters[index]
	a.object.Fields[index] = value
}

// Result returns the atom's return Field and value. If the atom does not have
// a return value then nil, nil is returned.
func (a *Atom) Result() (*schema.Field, interface{}) {
	if a.class.result < 0 {
		return nil, nil
	}
	return &a.object.Type.Fields[a.class.result], a.object.Fields[a.class.result]
}

// SetResult sets the atom's result to the specified value.
func (a *Atom) SetResult(value interface{}) {
	if a.class.result < 0 {
		panic("Atom has no result")
	}
	a.object.Fields[a.class.result] = value
}

// Class returns the serialize information and functionality for this type.
func (a *Atom) Class() binary.Class {
	return a.class
}

// String returns the string description of the atom and its arguments.
func (a *Atom) String() string {
	params := make([]string, a.ParameterCount())
	for i := range params {
		_, v := a.Parameter(i)
		params[i] = fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%v(%v)", a.class.base.Name, strings.Join(params, ", "))
}

// atomClass is an implementation of binary.Class used for atoms described by
// the schema.
type atomClass struct {
	base         *schema.Class
	meta         *atom.Metadata
	observations int   // index on fields, or -1
	parameters   []int // indices on fields
	result       int   // index on fields, or -1
}

var observationsId = (*atom.Observations)(nil).Class().ID()

func newAtomClass(base *schema.Class, meta *atom.Metadata) *atomClass {
	class := &atomClass{base: base, meta: meta, observations: -1}
	// Find the observations, if present
	for i, f := range base.Fields {
		if s, ok := f.Type.(*schema.Struct); ok {
			if s.ID == observationsId {
				class.observations = i
				continue
			}
		}
		if f.Name() == "Result" {
			class.result = i
			continue
		}
		class.parameters = append(class.parameters, i)
	}
	return class
}

func (c *atomClass) Schema() *schema.Class {
	return c.base
}

func (c *atomClass) ID() binary.ID {
	return c.base.ID()
}

func (c *atomClass) New() binary.Object {
	return &Atom{class: c, object: c.base.New().(*schema.Object)}
}

func (c *atomClass) Encode(e binary.Encoder, object binary.Object) {
	a := object.(*Atom)
	c.base.Encode(e, a.object)
}

func (c *atomClass) Decode(d binary.Decoder) binary.Object {
	a := &Atom{class: c}
	o := c.base.Decode(d)
	if d.Error() != nil {
		return a
	}
	a.object = o.(*schema.Object)
	if c.observations >= 0 {
		if c.observations >= len(a.object.Fields) {
			d.SetError(fmt.Errorf("Missing Observations field in %s", c.base.Name))
			return a
		}
		value := a.object.Fields[c.observations]
		if observations, ok := value.(*atom.Observations); !ok {
			d.SetError(fmt.Errorf("Observations field is of type %T in %s", value, c.base.Name))
			return a
		} else {
			a.observations = observations
		}
	}
	if a.class.meta.DrawCall {
		a.flags |= atom.DrawCall
	}
	if a.class.meta.EndOfFrame {
		a.flags |= atom.EndOfFrame
	}
	return a
}

func (c *atomClass) DecodeTo(d binary.Decoder, object binary.Object) {
	c.base.DecodeTo(d, object)
}
