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

package binary

var testObjA = &testObjectA{testObjectBase{"ObjectA"}}
var testObjB = &testObjectB{testObjectBase{"ObjectB"}}
var testObjC = &testObjectC{testObjectBase{"ObjectC"}}

var testObjectIDA = ID{0x0A}
var testObjectIDB = ID{0x0B}

type testObjectBase struct{ data string }

func (t *testObjectBase) Encode(e *Encoder) error {
	return e.String(t.data)
}

func (t *testObjectBase) Decode(d *Decoder) error {
	var err error
	t.data, err = d.String()
	return err
}

type testObjectA struct{ testObjectBase }
type testObjectB struct{ testObjectBase }
type testObjectC struct{ testObjectBase }

func init() {
	Register(testObjectIDA, &testObjectA{})
	Register(testObjectIDB, &testObjectB{})
}
