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

package schema

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Object is an instance of a Class.
type Object struct {
	class  *Class
	Fields []interface{}
}

// Class implements binary.Object using the schema system to do the encoding and
// decoding of fields.
func (o *Object) Class() binary.Class {
	return o.class
}

// Struct is the Type descriptor for an binary.Object typed value.
type Struct struct {
	binary.Generate
	Name string    // The simple name of the type.
	ID   binary.ID // The unique type identifier for the Object.
}

// Interface is the Type descriptor for a field who's underlying type is dynamic.
type Interface struct {
	binary.Generate
	Name string // The simple name of the type.
}

// Pointer is the Type descriptor for pointers.
type Pointer struct {
	binary.Generate
	Type Type // The pointed to type.
}

func (s *Struct) String() string {
	return s.Name
}

func (s *Struct) Encode(e binary.Encoder, value interface{}) error {
	return e.Value(value.(binary.Object))
}

func (s *Struct) Decode(d binary.Decoder) (interface{}, error) {
	o := &Object{class: Lookup(s.ID)}
	return o, d.Value(o)
}

func (s *Struct) Skip(d binary.Decoder) error {
	o := &Object{class: Lookup(s.ID)}
	return d.SkipValue(o)
}

func (i *Interface) String() string {
	return i.Name
}

func (i *Interface) Encode(e binary.Encoder, value interface{}) error {
	if value != nil { // TODO proper nil test needed?
		if err := e.Object(value.(binary.Object)); err != nil {
			return err
		}
	} else if err := e.Object(nil); err != nil {
		return err
	}
	return nil
}

func (i *Interface) Decode(d binary.Decoder) (interface{}, error) {
	return d.Object()
}

func (i *Interface) Skip(d binary.Decoder) error {
	_, err := d.SkipObject()
	return err
}

func (p *Pointer) String() string {
	return fmt.Sprintf("*%s", p.Type)
}

func (p *Pointer) Encode(e binary.Encoder, value interface{}) error {
	if value != nil { // TODO proper nil test needed?
		if err := e.Object(value.(binary.Object)); err != nil {
			return err
		}
	} else if err := e.Object(nil); err != nil {
		return err
	}
	return nil
}

func (p *Pointer) Decode(d binary.Decoder) (interface{}, error) {
	return d.Object()
}

func (p *Pointer) Skip(d binary.Decoder) error {
	_, err := d.SkipObject()
	return err
}
