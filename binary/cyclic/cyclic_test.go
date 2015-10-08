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
	"bytes"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/test"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

const (
	aa = uint32(123)
	bb = uint32(321)
	cc = uint32(456)
)

func EncodeObject(t *testing.T, entry test.Entry, e binary.Encoder, buf *bytes.Buffer) {
	e.SetMode(binary.Compact)
	for i, o := range entry.Values {
		e.Object(o)
		if e.Error() != nil {
			t.Errorf("%v[%v] Object gave unexpected error: %v", entry.Name, i, e.Error())
		}
	}
	test.VerifyData(t, entry, buf)
}

func DecodeObject(t *testing.T, entry test.Entry, d binary.Decoder, reader *bytes.Reader) {
	for i, o := range entry.Values {
		got := d.Object()
		if d.Error() != nil {
			t.Errorf("%v[%v] Object gave unexpected error: %v", entry.Name, i, d.Error())
		} else if !reflect.DeepEqual(o, got) {
			t.Errorf("%v[%v] unexpected object. Expected: %+v, got: %+v", entry.Name, i, o, got)
		}
	}
}

func TestObject(t *testing.T) {
	for _, entry := range []test.Entry{
		{
			Name:   "Nil",
			Values: []binary.Object{nil},
			Data:   []byte{0},
		},
		{
			Name:   "One",
			Values: []binary.Object{test.ObjectA},
			Data: test.Bytes{}.Add(
				0x03, // object sid + encoded
				0x03, // type sid + encoded
			).Add(test.EntityA...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',
			).Data,
		},
		{
			Name:   "Repeat",
			Values: []binary.Object{test.ObjectA, test.ObjectA},
			Data: test.Bytes{}.Add(
				0x03, // object sid + encoded
				0x03, // type sid + encoded
			).Add(test.EntityA...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x02, // repeated object sid
			).Data,
		},
		{
			Name:   "Many",
			Values: []binary.Object{test.ObjectA, test.ObjectB, test.ObjectA, nil},
			Data: test.Bytes{}.Add(
				0x03, // object sid + encoded
				0x03, // type sid + encoded
			).Add(test.EntityA...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'A',

				0x05, // object sid + encoded
				0x05, // type sid + encoded
			).Add(test.EntityB...).Add(
				0x07,
				'O', 'b', 'j', 'e', 'c', 't', 'B',

				0x02, // repeated object sid

				0x00, // nil object sid
			).Data,
		},
	} {
		b := &bytes.Buffer{}
		EncodeObject(t, entry, Encoder(vle.Writer(b)), b)
		r := bytes.NewReader(entry.Data)
		DecodeObject(t, entry, Decoder(vle.Reader(r)), r)
	}
}

func EncodeAndDecode(t *testing.T, obj binary.Object) {
	b := &bytes.Buffer{}
	e := Encoder(vle.Writer(b))
	e.Object(obj)
	if e.Error() != nil {
		t.Fatal(e.Error())
	}
	d := Decoder(vle.Reader(b))
	got := d.Object()
	if d.Error() != nil {
		t.Fatal(d.Error())
	}
	if !reflect.DeepEqual(got, obj) {
		// Note the decoder does not distinguish a nil map or slice from
		// an empty one. So the tests below confirm nil works.
		t.Errorf("Expected %v (%T) %p got %v (%T) %p", obj, obj, obj, got, got, got)
	}
	if len(d.substack.stack) != 0 {
		t.Errorf("Stack is not empty after decoding %v", obj)
	}
}

func TestNested(t *testing.T) {
	leaf := test.Leaf{A: aa}
	EncodeAndDecode(t, &leaf)
	anon := test.Anonymous{Leaf: test.Leaf{A: aa}}
	EncodeAndDecode(t, &anon)
	contains := test.Contains{LeafField: test.Leaf{A: aa}}
	EncodeAndDecode(t, &contains)
	leafSlice := []test.Leaf{test.Leaf{A: aa}, test.Leaf{A: bb}}
	slice := test.Slice{Leaves: leafSlice}
	EncodeAndDecode(t, &slice)
	emptySlice := test.Slice{}
	EncodeAndDecode(t, &emptySlice)
	emptyMapKey := test.MapKey{}
	EncodeAndDecode(t, &emptyMapKey)
	mapKey := test.MapKey{M: map[test.Leaf]uint32{test.Leaf{A: aa}: bb}}
	EncodeAndDecode(t, &mapKey)
	mapKeyValue := test.MapKeyValue{
		M: map[test.Leaf]test.Leaf{test.Leaf{A: aa}: test.Leaf{A: bb}}}
	EncodeAndDecode(t, &mapKeyValue)
	sliceInMap := test.SliceInMap{M: map[uint32][]test.Leaf{aa: leafSlice}}
	EncodeAndDecode(t, &sliceInMap)
	emptySliceInMap := test.SliceInMap{M: map[uint32][]test.Leaf{bb: nil}}
	EncodeAndDecode(t, &emptySliceInMap)
	sliceOfSlices := test.SliceOfSlices{
		Slice: [][]test.Leaf{leafSlice, []test.Leaf{test.Leaf{A: cc}}, nil, leafSlice}}
	EncodeAndDecode(t, &sliceOfSlices)
	mapInSlice := test.MapInSlice{
		Slice: []map[uint32]uint32{{cc: bb}, {aa: bb, bb: cc}}}
	EncodeAndDecode(t, &mapInSlice)
	mapOfMaps := test.MapOfMaps{M: map[uint32]map[test.Leaf]test.Leaf{
		aa: {leaf: leaf}, cc: {test.Leaf{A: bb}: test.Leaf{A: aa}}}}
	EncodeAndDecode(t, &mapOfMaps)
	leafArray := [3]test.Leaf{test.Leaf{A: aa}, test.Leaf{A: bb}, test.Leaf{A: cc}}
	array := test.Array{Leaves: leafArray}
	EncodeAndDecode(t, &array)
	mapInArray :=
		test.MapInArray{Array: [2]map[uint32]uint32{map[uint32]uint32{
			cc: bb, aa: cc}, map[uint32]uint32{bb: cc}}}
	EncodeAndDecode(t, &mapInArray)
	arrayInMap := test.ArrayInMap{M: map[uint32][3]test.Leaf{bb: leafArray}}
	EncodeAndDecode(t, &arrayInMap)
	arrayOfArrays := test.ArrayOfArrays{
		Array: [2][3]test.Leaf{leafArray, leafArray}}
	EncodeAndDecode(t, &arrayOfArrays)
	ca := test.Contains{LeafField: test.Leaf{A: aa}}
	cb := test.Contains{LeafField: test.Leaf{A: bb}}
	cc := test.Contains{LeafField: test.Leaf{A: cc}}

	complex := test.Complex{
		SliceMapArray: []map[test.Contains][3]test.Contains{
			{ca: {cb, cc, ca}},
		},
		SliceArrayMap: [][3]map[test.Contains]test.Contains{
			{{ca: cb}, {cb: cc, ca: cc}, {cb: cc}},
		},
		ArraySliceMap: [3][]map[test.Contains]test.Contains{
			{{ca: cb}, {cb: cc, ca: cc}, {cb: cc}},
			{{cb: cc, ca: cc}, {ca: cb}, {cb: cc}},
			{{ca: cb}, {cb: cc, ca: cc}, {cb: cc}},
		},
		ArrayMapSlice: [3]map[test.Contains][]test.Contains{
			{ca: {ca, cb, cc}},
			{cb: {cc, cb}},
			{cc: {cb, ca}},
		},
		MapArraySlice: map[test.Contains][3][]test.Contains{
			ca: {{ca, cb, cc}, {ca, cb}, {cc}},
		},
		MapSliceArray: map[test.Contains][][3]test.Contains{
			ca: {{ca, cb, cc}, {cc, cb, ca}},
		},
	}
	EncodeAndDecode(t, &complex)
}

func TestUnknownTypeError(t *testing.T) {
	d := Decoder(vle.Reader(bytes.NewBuffer([]byte{
		0x03, // object sid + encoded
		0x03, // type sid + encoded
		0x0C, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x09,
		'B', 'a', 'd', 'O', 'b', 'j', 'e', 'c', 't',
	})))
	d.AllowDynamic = false
	d.Object()
	if d.Error() == nil {
		t.Errorf("Expected error decoding unknown type")
	}
}
