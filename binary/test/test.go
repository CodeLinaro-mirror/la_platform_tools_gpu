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
	"android.googlesource.com/platform/tools/gpu/binary/objects"
)

type TypeA struct {
	binary.Generate
	data string
}

type TypeB struct {
	binary.Generate
	data string
}

type BadType struct {
	binary.Generate `disable:"true"`
	data            string
}

var ObjectA = &TypeA{data: "ObjectA"}
var ObjectB = &TypeB{data: "ObjectB"}
var BadObject = &BadType{data: "BadObject"}

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
	offsets := make([]int, len(entry.Values))
	for i, o := range entry.Values {
		got := reflect.New(reflect.TypeOf(o).Elem()).Interface().(binary.Object)
		if err := d.Value(got); err != nil {
			t.Errorf("%v[%v] Value gave unexpected error: %v", entry.Name, i, err)
		} else if !reflect.DeepEqual(o, got) {
			t.Errorf("%v[%v] unexpected object. Expected: %+v, got: %+v", entry.Name, i, o, got)
		}
		offsets[i] = reader.Len()
	}
	// Reset to beginning so we can verify skip offsets
	reader.Seek(0, 0)
	for i, o := range entry.Values {
		got := reflect.Zero(reflect.TypeOf(o)).Interface().(binary.Object)
		if err := d.SkipValue(got); err != nil {
			t.Errorf("%v[%v] SkipValue gave unexpected error: %v", entry.Name, i, err)
		}
		if offsets[i] != reader.Len() {
			t.Errorf("%v[%v] bad skip. Expected: %v, got: %v", entry.Name, i, offsets[i], reader.Len())
		}
	}
}

func DecodeObject(t *testing.T, entry Entry, d binary.Decoder, reader *bytes.Reader) {
	offsets := make([]int, len(entry.Values))
	for i, o := range entry.Values {
		if got, err := d.Object(); err != nil {
			t.Errorf("%v[%v] Object gave unexpected error: %v", entry.Name, i, err)
		} else if !reflect.DeepEqual(o, got) {
			t.Errorf("%v[%v] unexpected object. Expected: %+v, got: %+v", entry.Name, i, o, got)
		}
		offsets[i] = reader.Len()
	}
	// Reset to beginning so we can verify skip offsets
	reader.Seek(0, 0)
	for i, v := range entry.Values {
		var ty binary.Class = objects.NilClass
		if v != nil {
			ty = v.Class()
		}
		if id, err := d.SkipObject(); err != nil {
			t.Errorf("%v[%v] SkipObject gave unexpected error: %v", entry.Name, i, err)
		} else if ty.ID() != id {
			t.Errorf("%v[%v] SkipObject gave unexpected type: expected %v got %v", entry.Name, i, ty.ID, id)
		}
		if offsets[i] != reader.Len() {
			t.Errorf("%v[%v] bad skip. Expected: %v, got: %v", entry.Name, i, offsets[i], reader.Len())
		}
	}
}
