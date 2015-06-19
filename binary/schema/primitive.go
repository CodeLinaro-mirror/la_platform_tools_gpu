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

package schema

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Method denotes the encoding/decoding method a primitive type will use.
type Method int

// binary: Method.String = false

const (
	ID Method = iota
	Bool
	Int8
	Uint8
	Int16
	Uint16
	Int32
	Uint32
	Int64
	Uint64
	Float32
	Float64
	String
)

var (
	methodToString = map[Method]string{
		ID:      "ID",
		Bool:    "Bool",
		Int8:    "Int8",
		Uint8:   "Uint8",
		Int16:   "Int16",
		Uint16:  "Uint16",
		Int32:   "Int32",
		Uint32:  "Uint32",
		Int64:   "Int64",
		Uint64:  "Uint64",
		Float32: "Float32",
		Float64: "Float64",
		String:  "String",
	}
	stringToMethod = map[string]Method{}
	methodToBase   = map[Method]string{
		ID:      "binary.ID",
		Bool:    "bool",
		Int8:    "int8",
		Uint8:   "uint8",
		Int16:   "int16",
		Uint16:  "uint16",
		Int32:   "int32",
		Uint32:  "uint32",
		Int64:   "int64",
		Uint64:  "uint64",
		Float32: "float32",
		Float64: "float64",
		String:  "string",
	}
)

func init() {
	for m, s := range methodToString {
		stringToMethod[s] = m
	}
}

// Primitive is the kind for primitive types with corresponding direct methods on
// Encoder and Decoder
type Primitive struct {
	binary.Generate
	Name   string // The simple name of the type.
	Method Method // The enocde/decode method to use.
}

// Native returns the go native type name for this primitive.
func (p *Primitive) Native() string {
	switch p.Method {
	case ID:
		return "binary.ID"
	default:
		return strings.ToLower(p.Method.String())
	}
}

func (p *Primitive) Basename() string {
	return methodToBase[p.Method]
}

func (p *Primitive) Typename() string {
	return p.Name
}

func (p *Primitive) String() string {
	return p.Name
}

func (p *Primitive) Encode(e binary.Encoder, value interface{}) error {
	switch p.Method {
	case ID:
		return e.ID(value.(binary.ID))
	case Bool:
		return e.Bool(value.(bool))
	case Int8:
		return e.Int8(value.(int8))
	case Uint8:
		return e.Uint8(value.(uint8))
	case Int16:
		return e.Int16(value.(int16))
	case Uint16:
		return e.Uint16(value.(uint16))
	case Int32:
		return e.Int32(value.(int32))
	case Uint32:
		return e.Uint32(value.(uint32))
	case Int64:
		return e.Int64(value.(int64))
	case Uint64:
		return e.Uint64(value.(uint64))
	case Float32:
		return e.Float32(value.(float32))
	case Float64:
		return e.Float64(value.(float64))
	case String:
		return e.String(value.(string))
	default:
		return fmt.Errorf("Unknown encode method %q", p.Method)
	}
}

func (p *Primitive) Decode(d binary.Decoder) (interface{}, error) {
	switch p.Method {
	case ID:
		return d.ID()
	case Bool:
		return d.Bool()
	case Int8:
		return d.Int8()
	case Uint8:
		return d.Uint8()
	case Int16:
		return d.Int16()
	case Uint16:
		return d.Uint16()
	case Int32:
		return d.Int32()
	case Uint32:
		return d.Uint32()
	case Int64:
		return d.Int64()
	case Uint64:
		return d.Uint64()
	case Float32:
		return d.Float32()
	case Float64:
		return d.Float64()
	case String:
		return d.String()
	default:
		return nil, fmt.Errorf("Unknown decode method %q", p.Method)
	}
}

// Implements binary.Class
func (p *Primitive) Skip(d binary.Decoder) error {
	var err error
	switch p.Method {
	case ID:
		return d.SkipID()
	case Bool:
		_, err = d.Bool()
	case Int8:
		_, err = d.Int8()
	case Uint8:
		_, err = d.Uint8()
	case Int16:
		_, err = d.Int16()
	case Uint16:
		_, err = d.Uint16()
	case Int32:
		_, err = d.Int32()
	case Uint32:
		_, err = d.Uint32()
	case Int64:
		_, err = d.Int64()
	case Uint64:
		_, err = d.Uint64()
	case Float32:
		_, err = d.Float32()
	case Float64:
		_, err = d.Float64()
		return err
	case String:
		return d.SkipString()
	default:
		err = fmt.Errorf("Unknown skip method %q", p.Method)
	}
	return err
}

// This will convert a string to a Method, or return an error if the string was
// not a valid method name.
func ParseMethod(s string) (Method, error) {
	if m, ok := stringToMethod[s]; ok {
		return m, nil
	}
	return 0, fmt.Errorf("Invalid Method name %s", s)
}

func (m Method) String() string {
	if s, ok := methodToString[m]; ok {
		return s
	}
	return fmt.Sprintf("Method(%d)", m)
}

// Skippable returns true if the method has a complimentary skip method on the
// decoder interface. If this is not true, the normal decoding method is used
// during skipping.
func (m Method) Skippable() bool {
	return m == ID || m == String
}
