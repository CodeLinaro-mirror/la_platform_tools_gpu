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

func (e *encoder) Entity(*binary.Entity, bool) {
	panic(fmt.Errorf("Flat encoders do not support Schema objects"))
}

func (d *decoder) Entity(bool) *binary.Entity {
	panic(fmt.Errorf("Flat decoders do not support Schema objects"))
}

func (e *encoder) Value(obj binary.Object) { obj.Class().Encode(e, obj) }
func (d *decoder) Value(obj binary.Object) { obj.Class().DecodeTo(d, obj) }
func (e *encoder) Variant(obj binary.Object) {
	panic(fmt.Errorf("Flat decoders do not support Schema objects"))
}

func (d *decoder) Variant() binary.Object {
	panic(fmt.Errorf("Flat decoders do not support Schema objects"))
}

func (e *encoder) Object(obj binary.Object)           { e.Variant(obj) }
func (d *decoder) Object() binary.Object              { return d.Variant() }
func (d *decoder) Lookup(*binary.Entity) binary.Class { return nil }
