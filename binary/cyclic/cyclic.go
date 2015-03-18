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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
)

// Encoder creates a binary.Encoder that writes to the supplied binary.Writer.
func Encoder(writer binary.Writer) binary.Encoder {
	return &encoder{Writer: writer, objects: map[interface{}]uint32{}}
}

// Decoder creates a binary.Decoder that reads from the provided binary.Reader.
func Decoder(reader binary.Reader) binary.Decoder {
	return &decoder{Reader: reader, objects: map[uint32]interface{}{}}
}

type encoder struct {
	binary.Writer
	objects map[interface{}]uint32
}

type decoder struct {
	binary.Reader
	objects map[uint32]interface{}
}

func (e *encoder) Object(obj binary.Encodable) error {
	if obj == nil {
		return e.Uint32(0)
	}

	key, alreadyEncoded := e.objects[obj]
	if alreadyEncoded {
		return e.Uint32(key)
	}

	id, err := registry.TypeOf(obj)
	if err != nil {
		return err
	}

	key = uint32(len(e.objects) + 1)
	e.objects[obj] = key
	if err := e.Uint32(key); err != nil {
		return err
	}
	if err := id.Encode(e); err != nil {
		return err
	}
	return obj.Encode(e)
}

func (d *decoder) Object() (interface{}, error) {
	key, err := d.Uint32()
	if err != nil {
		return nil, err
	}

	if key == 0 {
		return nil, nil
	}

	if obj, alreadyDecoded := d.objects[key]; alreadyDecoded {
		return obj, nil
	}

	var id binary.ID
	if err := id.Decode(d); err != nil {
		return nil, err
	}

	obj, err := registry.New(id)
	if err != nil {
		return nil, err
	}

	if err = obj.Decode(d); err != nil {
		return nil, err
	}

	d.objects[key] = obj
	return obj, nil
}
