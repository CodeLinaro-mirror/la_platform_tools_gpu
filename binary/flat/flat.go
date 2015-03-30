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

func (e *encoder) Object(obj binary.Encodable) error {
	if obj == nil {
		return binary.ID{}.Encode(e)
	}
	if id, err := registry.TypeOf(obj); err != nil {
		return err
	} else if err := id.Encode(e); err != nil {
		return err
	}
	return obj.Encode(e)
}

func (d *decoder) Object() (interface{}, error) {
	var id binary.ID
	if err := id.Decode(d); err != nil {
		return nil, err
	}
	if obj, err := registry.New(id); err != nil || obj == nil {
		return obj, err
	} else {
		return obj, obj.Decode(d)
	}
}
