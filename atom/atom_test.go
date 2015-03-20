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

package atom

import "android.googlesource.com/platform/tools/gpu/binary"

const testAtomIDA = TypeID(10)
const testAtomIDB = TypeID(20)
const testAtomIDC = TypeID(30)

type testAtomA struct {
	Context ContextID
	Int32   int32
}

func (testAtomA) TypeID() TypeID          { return testAtomIDA }
func (a *testAtomA) ContextID() ContextID { return a.Context }
func (a *testAtomA) Info() string         { return "" }
func (a *testAtomA) Flags() Flags         { return 0 }
func (a *testAtomA) Encode(e binary.Encoder) error {
	if err := e.Uint32(uint32(a.Context)); err != nil {
		return err
	}
	return e.Int32(a.Int32)
}
func (a *testAtomA) Decode(d binary.Decoder) (err error) {
	if obj, err := d.Uint32(); err != nil {
		return err
	} else {
		a.Context = ContextID(obj)
	}
	a.Int32, err = d.Int32()
	return
}

type testAtomB struct {
	Context ContextID
	Bool    bool
}

func (testAtomB) TypeID() TypeID          { return testAtomIDB }
func (a *testAtomB) ContextID() ContextID { return a.Context }
func (a *testAtomB) Info() string         { return "" }
func (a *testAtomB) Flags() Flags         { return 0 }
func (a *testAtomB) Encode(e binary.Encoder) error {
	if err := e.Uint32(uint32(a.Context)); err != nil {
		return err
	}
	return e.Bool(a.Bool)
}
func (a *testAtomB) Decode(d binary.Decoder) (err error) {
	if obj, err := d.Uint32(); err != nil {
		return err
	} else {
		a.Context = ContextID(obj)
	}
	a.Bool, err = d.Bool()
	return
}

type testAtomC struct {
	Context ContextID
	String  string
}

func (testAtomC) TypeID() TypeID          { return testAtomIDC }
func (a *testAtomC) ContextID() ContextID { return a.Context }
func (a *testAtomC) Info() string         { return "" }
func (a *testAtomC) Flags() Flags         { return 0 }
func (a *testAtomC) Encode(e binary.Encoder) error {
	if err := e.Uint32(uint32(a.Context)); err != nil {
		return err
	}
	return e.String(a.String)
}
func (a *testAtomC) Decode(d binary.Decoder) (err error) {
	if obj, err := d.Uint32(); err != nil {
		return err
	} else {
		a.Context = ContextID(obj)
	}
	a.String, err = d.String()
	return
}

func init() {
	Register(TypeInfo{ID: testAtomIDA, New: func() Atom { return &testAtomA{} }})
	Register(TypeInfo{ID: testAtomIDB, New: func() Atom { return &testAtomB{} }})
	Register(TypeInfo{ID: testAtomIDC, New: func() Atom { return &testAtomC{} }})
}
