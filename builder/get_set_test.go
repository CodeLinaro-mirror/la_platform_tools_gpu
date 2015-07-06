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

package builder

import (
	"fmt"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

type testStruct struct {
	binary.Generate

	Str string
	Ptr *testStruct
}

type testAtom struct {
	binary.Generate

	api gfxapi.ID

	Str string
	Sli []bool
	Any interface{}
	Ptr *testStruct
	Map map[string]string
}

type testAPI struct{}

func (testAPI) Name() string  { return "foo" }
func (testAPI) ID() gfxapi.ID { return gfxapi.ID{1, 2, 3} }
func (testAPI) GetFramebufferAttachmentSize(state *gfxapi.State, attachment gfxapi.FramebufferAttachment) (uint32, uint32, error) {
	return 0, 0, nil
}

func (a testAtom) API() gfxapi.ID                                          { return a.api }
func (testAtom) Flags() atom.Flags                                         { return 0 }
func (testAtom) Observations() *atom.Observations                          { return &atom.Observations{} }
func (testAtom) Mutate(*gfxapi.State, database.Database, log.Logger) error { return nil }

func newPathTest(t *testing.T, atoms ...atom.Atom) (*path.Capture, database.Database, log.Logger) {
	d, l := database.NewInMemory(nil), log.Testing(t)
	id, err := ImportCapture("test", atoms, d, l)
	if err != nil {
		t.Fatalf("Failed to create capture: %v", err)
	}
	return id.Path(), d, l
}

func TestGet(t *testing.T) {
	gfxapi.Register(testAPI{})
	atomA := &testAtom{
		api: testAPI{}.ID(),
		Str: "aaa",
		Sli: []bool{true, false, true},
		Any: &testStruct{Str: "bbb"},
		Ptr: &testStruct{Str: "ccc", Ptr: &testStruct{Str: "ddd"}},
		Map: map[string]string{"cat": "meow", "dog": "woof"},
	}
	atomB := &testAtom{
		Str: "xxx",
		Sli: []bool{false, true, false},
		Any: &testStruct{Str: "www"},
		Map: map[string]string{"bird": "tweet", "fox": "?"},
	}
	atoms := atom.List{atomA, atomB}
	p, d, l := newPathTest(t, atoms...)

	// Get tests
	for _, test := range []struct {
		path path.Path
		val  interface{}
		err  error
	}{
		{p.Atoms().Index(1), atoms[1], nil},
		{p.Atoms().Index(1).Field("Str"), "xxx", nil},
		{p.Atoms().Index(1).Field("Sli"), []bool{false, true, false}, nil},
		{p.Atoms().Index(1).Field("Any"), &testStruct{Str: "www"}, nil},
		{p.Atoms().Index(0).Field("Ptr"), &testStruct{Str: "ccc", Ptr: &testStruct{Str: "ddd"}}, nil},
		{p.Atoms().Index(1).Field("Sli").ArrayIndex(1), true, nil},
		{p.Atoms().Index(1).Field("Str").ArrayIndex(1), byte('x'), nil},
		{p.Atoms().Index(1).Field("Map").MapIndex("bird"), "tweet", nil},

		// Test invalid paths
		{p.Atoms().Index(5), nil, fmt.Errorf(
			"Atom at Capture(%v).Atoms[5] is out of bounds [0-1]", p.ID)},
		{p.Atoms().Index(1).StateAfter(), nil, fmt.Errorf(
			"Atom at Capture(%v).Atoms[1] has no API", p.ID)},
		{p.Atoms().Index(0).StateAfter(), nil, fmt.Errorf(
			"No state for API '%s' after Capture(%v).Atoms[0]", testAPI{}.Name(), p.ID)},
		{p.Atoms().Index(1).Field("doesnotexist"), nil, fmt.Errorf(
			"Struct of type builder.testAtom at Capture(%v).Atoms[1] does not have field 'doesnotexist'", p.ID)},
		{p.Atoms().Index(1).Field("Ptr").Field("ccc"), nil, fmt.Errorf(
			"Pointer at Capture(%v).Atoms[1].Ptr is nil", p.ID)},
		{p.Atoms().Index(1).Field("Sli").Field("ccc"), nil, fmt.Errorf(
			"Type []bool at Capture(%v).Atoms[1].Sli is not a struct", p.ID)},
		{p.Atoms().Index(1).Field("Sli").ArrayIndex(4), nil, fmt.Errorf(
			"Index at Capture(%v).Atoms[1].Sli[4] is out of bounds [0-2]", p.ID)},
		{p.Atoms().Index(1).Field("Str").ArrayIndex(4), nil, fmt.Errorf(
			"Index at Capture(%v).Atoms[1].Str[4] is out of bounds [0-2]", p.ID)},
		{p.Atoms().Index(1).Field("Ptr").ArrayIndex(4), nil, fmt.Errorf(
			"Type *builder.testStruct at Capture(%v).Atoms[1].Ptr is not an array, slice or string", p.ID)},
		{p.Atoms().Index(1).Field("Map").MapIndex(10), nil, fmt.Errorf(
			"Map at Capture(%v).Atoms[1].Map has key of type string, got type int", p.ID)},
		{p.Atoms().Index(1).Field("Map").MapIndex("rabbit"), nil, fmt.Errorf(
			"Map at Capture(%v).Atoms[1].Map does not contain key rabbit", p.ID)},
		{p.Atoms().Index(1).Field("Ptr").MapIndex("foo"), nil, fmt.Errorf(
			"Type *builder.testStruct at Capture(%v).Atoms[1].Ptr is not a map", p.ID)},
	} {
		got, err := database.Build(&Get{Path: test.path}, d, l)
		if expected := test.err; !reflect.DeepEqual(err, expected) {
			t.Errorf("Get(%s) did not return expected error:\nGot:      %v\nExpected: %v",
				test.path.Path(), err, expected)
		}
		if expected := test.val; !reflect.DeepEqual(got, expected) {
			t.Errorf("Get(%s) was not as expected:\nGot:      %T %v\nExpected: %T %v",
				test.path.Path(), got, got, expected, expected)
		}
	}

	// Set tests
	for _, test := range []struct {
		path path.Path
		val  interface{}
		err  error
	}{
		{path: p.Atoms(), val: &service.AtomStream{Atoms: atom.List{atomB}}},
		{path: p.Atoms().Index(0), val: atomB}, {path: p.Atoms().Index(0).Field("Str"), val: "bbb"},
		{path: p.Atoms().Index(0).Field("Sli"), val: []bool{false, true, false}},
		{path: p.Atoms().Index(0).Field("Any"), val: 0.123},
		{path: p.Atoms().Index(0).Field("Ptr"), val: &testStruct{Str: "ddd"}},
		{path: p.Atoms().Index(0).Field("Ptr").Field("Str"), val: "purr"},

		// Test invalid paths
		{p.Atoms().Index(5), nil, fmt.Errorf(
			"Atom at Capture(%v).Atoms[5] is out of bounds [0-1]", p.ID)},
		{p.Atoms().Index(1).StateAfter(), nil, fmt.Errorf(
			"Atom at Capture(%v).Atoms[1] has no API", p.ID)},
		{p.Atoms().Index(0).StateAfter(), nil, fmt.Errorf(
			"No state for API '%s' after Capture(%v).Atoms[0]", testAPI{}.Name(), p.ID)},
		{p.Atoms().Index(1).Field("doesnotexist"), nil, fmt.Errorf(
			"Struct of type builder.testAtom at Capture(%v).Atoms[1] does not have field 'doesnotexist'", p.ID)},
		{p.Atoms().Index(1).Field("Ptr").Field("ccc"), nil, fmt.Errorf(
			"Pointer at Capture(%v).Atoms[1].Ptr is nil", p.ID)},
		{p.Atoms().Index(1).Field("Sli").Field("ccc"), nil, fmt.Errorf(
			"Type []bool at Capture(%v).Atoms[1].Sli is not a struct", p.ID)},
		{p.Atoms().Index(1).Field("Sli").ArrayIndex(4), nil, fmt.Errorf(
			"Index at Capture(%v).Atoms[1].Sli[4] is out of bounds [0-2]", p.ID)},
		{p.Atoms().Index(1).Field("Str").ArrayIndex(4), nil, fmt.Errorf(
			"Index at Capture(%v).Atoms[1].Str[4] is out of bounds [0-2]", p.ID)},
		{p.Atoms().Index(1).Field("Ptr").ArrayIndex(4), nil, fmt.Errorf(
			"Type *builder.testStruct at Capture(%v).Atoms[1].Ptr is not an array, slice or string", p.ID)},
		{p.Atoms().Index(1).Field("Map").MapIndex(10), nil, fmt.Errorf(
			"Map at Capture(%v).Atoms[1].Map has key of type string, got type int", p.ID)},
		{p.Atoms().Index(1).Field("Map").MapIndex("rabbit"), nil, fmt.Errorf(
			"Map at Capture(%v).Atoms[1].Map does not contain key rabbit", p.ID)},
		{p.Atoms().Index(1).Field("Ptr").MapIndex("foo"), nil, fmt.Errorf(
			"Type *builder.testStruct at Capture(%v).Atoms[1].Ptr is not a map", p.ID)},
	} {
		res, err := database.Build(&Set{Path: test.path, Value: test.val}, d, l)
		if expected := test.err; !reflect.DeepEqual(err, expected) {
			t.Errorf("Set(%s) did not return expected error:\nGot:      %v\nExpected: %v",
				test.path.Path(), err, expected)
		} else if (res == nil) == (err == nil) {
			t.Errorf("Set(%s) returned %T %v and %v.", res, res, err)
		}

		if err == nil {
			path := res.(path.Path)

			// Check the paths have changed
			if reflect.DeepEqual(test.path, p) {
				t.Errorf("Set(%s) returned an unchanged path")
			}

			// Get the changed value
			got, err := database.Build(&Get{Path: path}, d, l)
			if expected := error(nil); err != expected {
				t.Errorf("Get(%s) did not return expected error:\nGot:      %v\nExpected: %v",
					test.path.Path(), err, expected)
			}
			if expected := test.val; !reflect.DeepEqual(got, expected) {
				t.Errorf("Get(%s) was not as expected.\nGot:      %T %v\nExpected: %T %v",
					test.path.Path(), got, got, expected, expected)
			}
		}
	}
}
