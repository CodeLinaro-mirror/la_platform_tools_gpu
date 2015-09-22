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

package test

import (
	"bytes"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

const (
	aa    = 10
	bb    = 20
	cc    = "Hello"
	begin = "Begin"
	end   = "End"
)

func encodeX1(t *testing.T, e binary.Encoder) {
	e.Int32(aa)
	e.Int32(bb)
}

func encodeX1toUpgrade(t *testing.T, e binary.Encoder) {
	x1 := X_V1{}
	// TODO (should we allow encoding of a schema of a frozen class).
	e.Entity(x1.Class().Schema(), true)
	encodeX1(t, e)
}

func encodeYcontainsX1(t *testing.T, e binary.Encoder) {
	// Build a schema object for a Y containing X1.
	schemaY1 := &binary.Entity{
		Package:  "test",
		Identity: "Y",
		Fields: []binary.Field{
			{Declared: "begin", Type: &schema.Primitive{Name: "string", Method: schema.String}},
			{Declared: "x", Type: &schema.Struct{Entity: (*X_V1)(nil).Class().Schema()}},
			{Declared: "end", Type: &schema.Primitive{Name: "string", Method: schema.String}},
		},
	}
	e.Entity(schemaY1, true)
	e.String(begin)
	encodeX1(t, e)
	e.String(end)
}

func decoderForX(t *testing.T) binary.Decoder {
	buf := bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(&buf))
	encodeX1toUpgrade(t, e)
	if err := e.Error(); err != nil {
		t.Fatal(err)
	}
	return cyclic.Decoder(vle.Reader(&buf))
}

func decoderForY(t *testing.T) binary.Decoder {
	buf := bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(&buf))
	encodeYcontainsX1(t, e)
	if err := e.Error(); err != nil {
		t.Fatal(err)
	}
	return cyclic.Decoder(vle.Reader(&buf))
}

func TestEncodeX1(t *testing.T) {
	x1 := X_V1{a: aa, b: bb}
	buf := bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(&buf))
	e.Variant(&x1)
	if err := e.Error(); err == nil {
		t.Errorf("Expected error encoding X_V1")
	}
}

func TestDecodeX(t *testing.T) {
	d := decoderForX(t)
	o := d.Variant()
	if err := d.Error(); err != nil {
		t.Fatalf("Error decoding X: %v", err)
	}
	x := o.(*X)
	expect := X{a: aa, b: bb, c: cc}
	if *x != expect {
		t.Errorf("Got %v Expected %v", x, expect)
	}
}

func TestDecodeY(t *testing.T) {
	d := decoderForY(t)
	o := d.Variant()
	if err := d.Error(); err != nil {
		t.Fatalf("Error decoding Y: %v", err)
	}
	y := o.(*Y)
	expect := Y{begin: begin, x: X{a: aa, b: bb, c: cc}, end: end}
	if *y != expect {
		t.Errorf("Got %v Expected %v", y, expect)
	}
}
