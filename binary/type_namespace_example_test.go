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

type Cat struct {
	// fields...
}

func (*Cat) Encode(e *Encoder) error { /* ...encode fields... */ return nil }
func (*Cat) Decode(d *Decoder) error { /* ...decode fields... */ return nil }

type Dog struct {
	// fields...
}

func (*Dog) Encode(e *Encoder) error { /* ...encode fields... */ return nil }
func (*Dog) Decode(d *Decoder) error { /* ...decode fields... */ return nil }

// Constructs and decodes a Cat using the decoder
func DecodeCat(d *Decoder) (Object, error) {
	o := &Cat{}
	err := o.Decode(d)
	return o, err
}

// Constructs and decodes a Dog using the decoder
func DecodeDog(d *Decoder) (Object, error) {
	o := &Dog{}
	err := o.Decode(d)
	return o, err
}

var AnimalNamespace = NewTypeNamespace(
	Type{ID: 0x1234, New: func() Object { return &Cat{} }},
	Type{ID: 0x4567, New: func() Object { return &Dog{} }},
)

func ExampleNewTypeNamespace() TypeNamespace {
	return AnimalNamespace
}
