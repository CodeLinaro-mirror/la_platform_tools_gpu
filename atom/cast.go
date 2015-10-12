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
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Dynamic is a wrapper around a schema.Object that describes an atom object.
// Dynamic conforms to the Atom interface, and provides a number of methods
// for accessing the parameters, return value and observations.
type Dynamic struct {
	object       *schema.Object
	info         *atomInfo
	observations *Observations
	flags        Flags
}

var _ Atom = &Dynamic{} // Verify that Dynamic implements Atom.

// API returns the graphics API id this atom belongs to.
func (a *Dynamic) API() gfxapi.ID {
	return a.info.meta.API
}

// Flags returns the flags of the atom.
func (a *Dynamic) Flags() Flags {
	return a.flags
}

// Observations returns all the memory observations made by the atom.
func (a *Dynamic) Observations() *Observations {
	return a.observations
}

// Mutate is not supported by the Dynamic type, but is exposed in order to comform
// to the Atom interface. Mutate will always return an error.
func (*Dynamic) Mutate(*gfxapi.State, database.Database, log.Logger) error {
	return fmt.Errorf("Mutate not implemented for client atoms")
}

// ParameterCount returns the number of parameters this atom accepts. This count
// does not include the return value.
func (a *Dynamic) ParameterCount() int {
	return len(a.info.parameters)
}

// Parameter returns the index'th parameter Field and value.
func (a *Dynamic) Parameter(index int) (binary.Field, interface{}) {
	index = a.info.parameters[index]
	return a.object.Type.Fields[index], a.object.Fields[index]
}

// SetParameter sets the atom's index'th parameter to the specified value.
func (a *Dynamic) SetParameter(index int, value interface{}) {
	index = a.info.parameters[index]
	a.object.Fields[index] = value
}

// Result returns the atom's return Field and value. If the atom does not have
// a return value then nil, nil is returned.
func (a *Dynamic) Result() (*binary.Field, interface{}) {
	if a.info.result < 0 {
		return nil, nil
	}
	return &a.object.Type.Fields[a.info.result], a.object.Fields[a.info.result]
}

// SetResult sets the atom's result to the specified value.
func (a *Dynamic) SetResult(value interface{}) {
	if a.info.result < 0 {
		panic("Atom has no result")
	}
	a.object.Fields[a.info.result] = value
}

// Class returns the serialize information and functionality for this type.
func (a *Dynamic) Class() binary.Class {
	return a.object.Class()
}

// String returns the string description of the atom and its arguments.
func (a *Dynamic) String() string {
	params := make([]string, a.ParameterCount())
	for i := range params {
		_, v := a.Parameter(i)
		params[i] = fmt.Sprintf("%v", v)
	}
	return fmt.Sprintf("%v(%v)", a.Class().Schema().Name(), strings.Join(params, ", "))
}

// atomInfo is a cache of precalculated atom information from the schema.
type atomInfo struct {
	meta         *Metadata
	observations int   // index on fields, or -1
	parameters   []int // indices on fields
	result       int   // index on fields, or -1
}

var observationsSig = (*Observations)(nil).Class().Schema().Signature()

func newAtomInfo(entity *binary.Entity) *atomInfo {
	meta := FindMetadata(entity)
	class := &atomInfo{meta: meta, observations: -1}
	// Find the observations, if present
	for i, f := range entity.Fields {
		if s, ok := f.Type.(*schema.Struct); ok {
			if s.Entity.Signature() == observationsSig {
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

var (
	wrappers = map[binary.Signature]*atomInfo{}
)

func Wrap(obj *schema.Object) (Atom, error) {
	entity := obj.Class().Schema()
	info := wrappers[entity.Signature()]
	if info == nil {
		info = newAtomInfo(entity)
		wrappers[entity.Signature()] = info
	}
	a := &Dynamic{info: info, object: obj}
	if info.observations >= 0 {
		if info.observations >= len(a.object.Fields) {
			return a, fmt.Errorf("Missing Observations field in %s", entity.Name())
		}
		value := a.object.Fields[info.observations]
		if observations, ok := value.(*Observations); !ok {
			return a, fmt.Errorf("Observations field is of type %T in %s", value, entity.Name())
		} else {
			a.observations = observations
		}
	}
	if info.meta != nil {
		if info.meta.DrawCall {
			a.flags |= DrawCall
		}
		if info.meta.EndOfFrame {
			a.flags |= EndOfFrame
		}
	}
	return a, nil
}
