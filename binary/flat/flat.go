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
func Encoder(writer binary.Writer) *encoder {
	return &encoder{Writer: writer}
}

// Decoder creates a binary.Decoder that reads from the provided binary.Reader.
func Decoder(reader binary.Reader) *decoder {
	return &decoder{Reader: reader}
}

type encoder struct {
	binary.Writer
}

type decoder struct {
	binary.Reader
}

func (e *encoder) Value(obj binary.Object) {
	obj.Class().Encode(e, obj)
}

func (e *encoder) Struct(obj binary.Object) {
	// Permitted until the code generation is smarter.
	e.Value(obj)
}

func (d *decoder) Value(obj binary.Object) {
	obj.Class().DecodeTo(d, obj)
}

func (d *decoder) Struct(obj binary.Object) {
	// Permitted until the code generation is smarter.
	d.Value(obj)
}

func (e *encoder) Entity(*binary.Entity, bool) {
	panic(fmt.Errorf("Flat encoders do not support Schema objects"))
}

func (d *decoder) Entity(bool) *binary.Entity {
	panic(fmt.Errorf("Flat decoders do not support Schema objects"))
}

func (e *encoder) Variant(obj binary.Object) {
	panic("e.Variant() called on flat decoder")
}

func (d *decoder) Variant() binary.Object {
	panic("d.Variant() called on flat decoder")
}

func (e *encoder) Object(obj binary.Object) {
	panic("e.Object() called on flat encoder")
}
func (d *decoder) Object() binary.Object {
	panic("d.Object() called on flat decoder")
}

func (d *decoder) Lookup(ent *binary.Entity) binary.UpgradeDecoder {
	panic("d.Lookup() called on flat decoder")
}

func (d *decoder) Count() uint32 {
	return d.Uint32()
}
