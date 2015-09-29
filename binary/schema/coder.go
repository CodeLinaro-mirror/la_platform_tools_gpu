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

	"android.googlesource.com/platform/tools/gpu/binary"
)

// TypeTag denotes the schema type that follows.
// Each tag corresponts to an implementation of the Type interface.
type TypeTag uint8

const (
	PrimitiveTag TypeTag = iota
	StructTag
	PointerTag
	InterfaceTag
	VariantTag
	AnyTag
	SliceTag
	ArrayTag
	MapTag
)

func EncodeType(e binary.Encoder, t binary.Type) {
	switch t := t.(type) {
	case *Primitive:
		e.Uint8(uint8(PrimitiveTag))
		e.String(t.Name)
		e.Uint8(uint8(t.Method))
	case *Struct:
		e.Uint8(uint8(StructTag))
		e.Entity(t.Entity)
	case *Pointer:
		e.Uint8(uint8(PointerTag))
		EncodeType(e, t.Type)
	case *Interface:
		e.Uint8(uint8(InterfaceTag))
		e.String(t.Name)
	case *Variant:
		e.Uint8(uint8(VariantTag))
		e.String(t.Name)
	case *Any:
		e.Uint8(uint8(AnyTag))
	case *Slice:
		e.Uint8(uint8(SliceTag))
		e.String(t.Alias)
		EncodeType(e, t.ValueType)
	case *Array:
		e.Uint8(uint8(ArrayTag))
		e.String(t.Alias)
		e.Uint32(t.Size)
		EncodeType(e, t.ValueType)
	case *Map:
		e.Uint8(uint8(MapTag))
		e.String(t.Alias)
		EncodeType(e, t.KeyType)
		EncodeType(e, t.ValueType)
	default:
		panic(fmt.Errorf("Encode unknown type %T", t))
	}
}

func DecodeType(d binary.Decoder) binary.Type {
	tag := TypeTag(d.Uint8())
	switch tag {
	case PrimitiveTag:
		t := &Primitive{}
		t.Name = d.String()
		t.Method = Method(d.Uint8())
		return t
	case StructTag:
		t := &Struct{}
		t.Entity = d.Entity()
		return t
	case PointerTag:
		t := &Pointer{}
		t.Type = DecodeType(d)
		return t
	case InterfaceTag:
		t := &Interface{}
		t.Name = d.String()
		return t
	case VariantTag:
		t := &Variant{}
		t.Name = d.String()
		return t
	case AnyTag:
		return &Any{}
	case SliceTag:
		t := &Slice{}
		t.Alias = d.String()
		t.ValueType = DecodeType(d)
		return t
	case ArrayTag:
		t := &Array{}
		t.Alias = d.String()
		t.Size = d.Uint32()
		t.ValueType = DecodeType(d)
		return t
	case MapTag:
		t := &Map{}
		t.Alias = d.String()
		t.KeyType = DecodeType(d)
		t.ValueType = DecodeType(d)
		return t
	default:
		panic(fmt.Errorf("Decode unknown type %v", tag))
	}
}

func EncodeEntity(e binary.Encoder, c *binary.Entity) {
	e.ID(c.TypeID)
	e.String(c.Package)
	e.String(c.Name)
	e.String(c.Identity)
	e.String(c.Version)
	e.Bool(c.Exported)
	e.Uint32(uint32(len(c.Fields)))
	for _, f := range c.Fields {
		e.String(f.Declared)
		EncodeType(e, f.Type)
	}
	e.Uint32(uint32(len(c.Metadata)))
	for _, m := range c.Metadata {
		e.Object(m)
	}
}

func DecodeEntity(d binary.Decoder, c *binary.Entity) {
	c.TypeID = d.ID()
	c.Package = d.String()
	c.Name = d.String()
	c.Identity = d.String()
	c.Version = d.String()
	c.Exported = d.Bool()
	c.Fields = make(binary.FieldList, d.Uint32())
	for i := range c.Fields {
		c.Fields[i].Declared = d.String()
		c.Fields[i].Type = DecodeType(d)
	}
	c.Metadata = make([]binary.Object, d.Uint32())
	for i := range c.Metadata {
		c.Metadata[i] = d.Object()
	}
}

func EncodeConstants(e binary.Encoder, c *ConstantSet) {
	EncodeType(e, c.Type)
	e.Uint32(uint32(len(c.Entries)))
	for _, entry := range c.Entries {
		e.String(entry.Name)
		c.Type.EncodeValue(e, entry.Value)
	}
}

func DecodeConstants(d binary.Decoder, c *ConstantSet) {
	c.Type = DecodeType(d)
	c.Entries = make([]Constant, d.Uint32())
	for i := range c.Entries {
		c.Entries[i].Name = d.String()
		c.Entries[i].Value = c.Type.DecodeValue(d)
	}
}
