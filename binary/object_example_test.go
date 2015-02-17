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
)

var ExampleObjectID = binary.NewID([]byte("ExampleObjectId"))

type ExampleObject struct{ Data string }

func (t *ExampleObject) Encode(e *binary.Encoder) error {
	return e.String(t.Data)
}

func (t *ExampleObject) Decode(d *binary.Decoder) error {
	var err error
	t.Data, err = d.String()
	return err
}

func init() {
	binary.Register(ExampleObjectID, &ExampleObject{})
}

// This example shows how to write a type with custom encode and decode
// methods, and send it over a "connection"
func Example_object() {
	// Create a connected input and output stream for example purposes.
	in := io.Reader(&bytes.Buffer{})
	out := in.(io.Writer)

	// Build an encoder and decoder on top of the stream
	e := binary.NewEncoder(out)
	d := binary.NewDecoder(in)

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
