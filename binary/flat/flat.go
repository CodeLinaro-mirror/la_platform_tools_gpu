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

package flat

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
)

// Encoder creates a binary.Encoder that writes to the supplied binary.Writer.
func Encoder(writer binary.Writer) binary.Encoder {
	return &encoder{Writer: writer}
}

// Decoder creates a binary.Decoder that reads from the provided binary.Reader.
func Decoder(reader binary.Reader) binary.Decoder {
	return &decoder{
		Reader:    reader,
		Namespace: registry.Global,
	}
}

type encoder struct {
	binary.Writer
}

type decoder struct {
	binary.Reader
	Namespace *registry.Namespace
}

func (e *encoder) ID(id binary.ID) error {
	return e.Data(id[:])
}

func (d *decoder) ID() (binary.ID, error) {
	id := binary.ID{}
	return id, d.Data(id[:])
}

func (d *decoder) SkipID() error {
	return d.Skip(binary.IDSize)
}

func (e *encoder) Value(obj binary.Object) error     { return obj.Class().Encode(e, obj) }
func (d *decoder) Value(obj binary.Object) error     { return obj.Class().DecodeTo(d, obj) }
func (d *decoder) SkipValue(obj binary.Object) error { return obj.Class().Skip(d) }

func (e *encoder) Variant(obj binary.Object) error {
	if obj == nil {
		return e.ID(binary.ID{})
	}
	class := obj.Class()
	if err := e.ID(class.ID()); err != nil {
		return err
	}
	return class.Encode(e, obj)
}

func (d *decoder) Variant() (binary.Object, error) {
	if id, err := d.ID(); err != nil {
		return nil, err
	} else if class := d.Namespace.Lookup(id); class == nil {
		return nil, fmt.Errorf("Unknown type id %v", id)
	} else {
		return class.Decode(d)
	}
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

func (e *encoder) Object(obj binary.Object) error   { return e.Variant(obj) }
func (d *decoder) Object() (binary.Object, error)   { return d.Variant() }
func (d *decoder) SkipObject() (binary.ID, error)   { return d.SkipVariant() }
func (d *decoder) Lookup(id binary.ID) binary.Class { return d.Namespace.Lookup(id) }
