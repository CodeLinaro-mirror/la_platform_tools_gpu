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

package binary_test

import (
	"bytes"
	"fmt"
	"io"
	"log"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

type ExampleObject struct{ Data string }
type ExampleClass struct{}

var ExampleID = binary.ID{0x01, 0x02, 0x03, 0x04, 0x05, 0x06, 0x07, 0x08, 0x09, 0x0A, 0x0B, 0x0C, 0x0D, 0x0E, 0x0F, 0x10, 0x11, 0x12, 0x13, 0x14}

func (*ExampleObject) Class() binary.Class {
	return (*ExampleClass)(nil)
}

func (*ExampleClass) ID() binary.ID {
	return ExampleID
}

func (*ExampleClass) Encode(e binary.Encoder, obj binary.Object) error {
	o := obj.(*ExampleObject)
	return e.String(o.Data)
}

func (*ExampleClass) Decode(d binary.Decoder) (binary.Object, error) {
	o := &ExampleObject{}
	var err error
	o.Data, err = d.String()
	return o, err
}

func (*ExampleClass) DecodeTo(d binary.Decoder, obj binary.Object) error {
	o := obj.(*ExampleObject)
	var err error
	o.Data, err = d.String()
	return err
}

func (*ExampleClass) Skip(d binary.Decoder) error {
	return d.SkipString()
}

func init() {
	registry.Add((*ExampleObject)(nil).Class())
}

// This example shows how to write a type with custom encode and decode
// methods, and send it over a "connection"
func Example_object() {
	// Create a connected input and output stream for example purposes.
	in := io.Reader(&bytes.Buffer{})
	out := in.(io.Writer)

	// Build an encoder and decoder on top of the stream
	e := cyclic.Encoder(vle.Writer(out))
	d := cyclic.Decoder(vle.Reader(in))

	// Encode an object onto the stream
	if err := e.Object(&ExampleObject{"MyObject"}); err != nil {
		log.Fatalf("Encode gave unexpected error: %v", err)
	}

	// Read the object back
	if o, err := d.Object(); err == nil {
		fmt.Printf("read %q\n", o.(*ExampleObject).Data)
	} else {
		log.Fatalf("Decode gave unexpected error: %v", err)
	}

	// Output:
	// read "MyObject"
}
