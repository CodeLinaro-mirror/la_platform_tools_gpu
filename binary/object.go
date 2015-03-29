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

import "fmt"

type Object interface {
	Encodable
	Decodable
}

// Generate is used to tag structures that need auto generated Encode and Decode
// methods.
// The codergen function searches packages for structs that have this type as an
// anonymous field, and then automatically generates the encoding and decoding
// functionality for those structures. For example, the following struct would
// create the methods needed to encode and decode the Name and Value fields, as
// well as registering a type identifier.
// The embedding will also fully implement the binary.Object interface, but with
// methods that panic. This will get overridden with the generated Encode and
// Decode methods. This is important because it means the package is resolvable
// without the generated code, which means the types can be correctly evaluated
// during the generation process.
//
// type MyNamedValue struct {
//    binary.Generate
//    Name  string
//    Value []byte
// }
type Generate struct{}

func (Generate) Encode(Encoder) error { panic(fmt.Errorf("Missing encode function")) }
func (Generate) Decode(Decoder) error { panic(fmt.Errorf("Missing decode function")) }
