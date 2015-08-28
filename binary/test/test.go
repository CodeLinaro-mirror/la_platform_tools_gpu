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

package test

import (
	"bytes"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// binary: java.source = base/rpclib/src/test/java
// binary: java.package = com.android.tools.rpclib.binary
// binary: java.indent = "    "
// binary: java.member_prefix = m

type TypeA struct {
	binary.Generate `id:"TypeAID"`
	Data            string
}

type TypeB struct {
	binary.Generate `id:"TypeBID"`
	Data            string
}

type BadType struct {
	binary.Generate `disable:"true"`
	Data            string
}

var ObjectA = &TypeA{Data: "ObjectA"}
var ObjectB = &TypeB{Data: "ObjectB"}
var BadObject = &BadType{Data: "BadObject"}

type Entry struct {
	Name   string
	Values []binary.Object
	Data   []byte
}

func VerifyData(t *testing.T, entry Entry, got *bytes.Buffer) {
	if !bytes.Equal(entry.Data, got.Bytes()) {
		t.Errorf(`%v gave unexpected bytes.
Expected: %# x
Got:      %# x`, entry.Name, entry.Data, got.Bytes())
	}
}

func EncodeValue(t *testing.T, entry Entry, e binary.Encoder, buf *bytes.Buffer) {
	for i, o := range entry.Values {
		if err := e.Value(o); err != nil {
			t.Errorf("%v[%v] Value gave unexpected error: %v", entry.Name, i, err)
		}
	}
	VerifyData(t, entry, buf)
}

func EncodeObject(t *testing.T, entry Entry, e binary.Encoder, buf *bytes.Buffer) {
	for i, o := range entry.Values {
		if err := e.Object(o); err != nil {
			t.Errorf("%v[%v] Object gave unexpected error: %v", entry.Name, i, err)
		}
	}
	VerifyData(t, entry, buf)
}

func DecodeValue(t *testing.T, entry Entry, d binary.Decoder, reader *bytes.Reader) {
	for i, o := range entry.Values {
		got := reflect.New(reflect.TypeOf(o).Elem()).Interface().(binary.Object)
		if err := d.Value(got); err != nil {
			t.Errorf("%v[%v] Value gave unexpected error: %v", entry.Name, i, err)
		} else if !reflect.DeepEqual(o, got) {
			t.Errorf("%v[%v] unexpected object. Expected: %+v, got: %+v", entry.Name, i, o, got)
		}
	}
}

func DecodeObject(t *testing.T, entry Entry, d binary.Decoder, reader *bytes.Reader) {
	for i, o := range entry.Values {
		if got, err := d.Object(); err != nil {
			t.Errorf("%v[%v] Object gave unexpected error: %v", entry.Name, i, err)
		} else if !reflect.DeepEqual(o, got) {
			t.Errorf("%v[%v] unexpected object. Expected: %+v, got: %+v", entry.Name, i, o, got)
		}
	}
}
