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

//go:generate codergen -go

package flat

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
)

// Encoder creates a binary.Encoder that writes to the supplied binary.Writer.
func Encoder(writer binary.Writer) binary.Encoder {
	return &encoder{Writer: writer}
}

// Decoder creates a binary.Decoder that reads from the provided binary.Reader.
func Decoder(reader binary.Reader) binary.Decoder {
	return &decoder{Reader: reader}
}

type encoder struct {
	binary.Writer
}

type decoder struct {
	binary.Reader
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

func (e *encoder) Value(obj binary.Encodable) error     { return obj.Encode(e) }
func (d *decoder) Value(obj binary.Decodable) error     { return obj.Decode(d) }
func (d *decoder) SkipValue(obj binary.Decodable) error { return obj.Skip(d) }

func (e *encoder) Variant(obj binary.Encodable) error {
	if obj == nil {
		return e.ID(binary.ID{})
	}
	if id, err := registry.TypeOf(obj); err != nil {
		return err
	} else if err := e.ID(id); err != nil {
		return err
	}
	return obj.Encode(e)
}

func (d *decoder) Variant() (interface{}, error) {
	if id, err := d.ID(); err != nil {
		return nil, err
	} else if obj, err := registry.New(id); err != nil || obj == nil {
		return obj, err
	} else {
		return obj, obj.Decode(d)
	}
}

func (d *decoder) SkipVariant() error {
	if id, err := d.ID(); err != nil {
		return err
	} else if obj, err := registry.Nil(id); err != nil || obj == nil {
		return err
	} else {
		return obj.Skip(d)
	}
}

func (e *encoder) Object(obj binary.Encodable) error { return e.Variant(obj) }
func (d *decoder) Object() (interface{}, error)      { return d.Variant() }
func (d *decoder) SkipObject() error                 { return d.SkipVariant() }
