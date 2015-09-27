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
	"android.googlesource.com/platform/tools/gpu/binary/registry"
)

var Namespace = registry.NewNamespace()

func init() {
	registry.Global.AddFallbacks(Namespace)
	Namespace.Add((*Entity)(nil).Class())
	Namespace.Add((*ConstantSet)(nil).Class())
}

var (
	binaryIDClass       = binary.ID{0xf1, 0xab, 0xae, 0xcf, 0xc3, 0x23, 0xf8, 0x65, 0xa1, 0xeb, 0xe0, 0x3a, 0xa1, 0xae, 0xb3, 0xab, 0x77, 0xb0, 0x57, 0xef}
	binaryIDConstantSet = binary.ID{0x28, 0x8f, 0x6c, 0x88, 0x31, 0xd1, 0x04, 0x52, 0xb7, 0x5a, 0x25, 0x83, 0x01, 0x4e, 0x9a, 0x7c, 0x53, 0x03, 0x32, 0x9e}
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

func encodeType(e binary.Encoder, t Type) {
	switch t := t.(type) {
	case *Primitive:
		e.Uint8(uint8(PrimitiveTag))
		e.String(t.Name)
		e.Uint8(uint8(t.Method))
	case *Struct:
		e.Uint8(uint8(StructTag))
		e.String(t.Name)
		e.ID(t.ID)
	case *Pointer:
		e.Uint8(uint8(PointerTag))
		encodeType(e, t.Type)
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
		encodeType(e, t.ValueType)
	case *Array:
		e.Uint8(uint8(ArrayTag))
		e.String(t.Alias)
		e.Uint32(t.Size)
		encodeType(e, t.ValueType)
	case *Map:
		e.Uint8(uint8(MapTag))
		e.String(t.Alias)
		encodeType(e, t.KeyType)
		encodeType(e, t.ValueType)
	default:
		panic(fmt.Errorf("Encode unknown type %T", t))
	}
}

func decodeType(d binary.Decoder) Type {
	tag := TypeTag(d.Uint8())
	switch tag {
	case PrimitiveTag:
		t := &Primitive{}
		t.Name = d.String()
		t.Method = Method(d.Uint8())
		return t
	case StructTag:
		t := &Struct{}
		t.Name = d.String()
		t.ID = d.ID()
		return t
	case PointerTag:
		t := &Pointer{}
		t.Type = decodeType(d)
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
		t.ValueType = decodeType(d)
		return t
	case ArrayTag:
		t := &Array{}
		t.Alias = d.String()
		t.Size = d.Uint32()
		t.ValueType = decodeType(d)
		return t
	case MapTag:
		t := &Map{}
		t.Alias = d.String()
		t.KeyType = decodeType(d)
		t.ValueType = decodeType(d)
		return t
	default:
		panic(fmt.Errorf("Decode unknown type %v", tag))
	}
}

func encodeClass(e binary.Encoder, c *Entity) {
	e.ID(c.TypeID)
	e.String(c.Package)
	e.String(c.Name)
	e.String(c.Identity)
	e.String(c.Version)
	e.Bool(c.Exported)
	e.Uint32(uint32(len(c.Fields)))
	for _, f := range c.Fields {
		e.String(f.Declared)
		encodeType(e, f.Type)
	}
	e.Uint32(uint32(len(c.Metadata)))
	for _, m := range c.Metadata {
		e.Object(m)
	}
}

func decodeClass(d binary.Decoder, c *Entity) {
	c.TypeID = d.ID()
	c.Package = d.String()
	c.Name = d.String()
	c.Identity = d.String()
	c.Version = d.String()
	c.Exported = d.Bool()
	c.Fields = make(FieldList, d.Uint32())
	for i := range c.Fields {
		c.Fields[i].Declared = d.String()
		c.Fields[i].Type = decodeType(d)
	}
	c.Metadata = make([]binary.Object, d.Uint32())
	for i := range c.Metadata {
		c.Metadata[i] = d.Object()
	}
}

func encodeConstants(e binary.Encoder, c *ConstantSet) {
	encodeType(e, c.Type)
	e.Uint32(uint32(len(c.Entries)))
	for _, entry := range c.Entries {
		e.String(entry.Name)
		c.Type.Encode(e, entry.Value)
	}
}

func decodeConstants(d binary.Decoder, c *ConstantSet) {
	c.Type = decodeType(d)
	c.Entries = make([]Constant, d.Uint32())
	for i := range c.Entries {
		c.Entries[i].Name = d.String()
		c.Entries[i].Value = c.Type.Decode(d)
	}
}

type binaryClassClass struct{}

func (*Entity) Class() binary.Class          { return (*binaryClassClass)(nil) }
func (*binaryClassClass) ID() binary.ID      { return binaryIDClass }
func (*binaryClassClass) New() binary.Object { return &Entity{} }
func (*binaryClassClass) Encode(e binary.Encoder, obj binary.Object) {
	encodeClass(e, obj.(*Entity))
}
func (*binaryClassClass) Decode(d binary.Decoder) binary.Object {
	c := &Entity{}
	decodeClass(d, c)
	return c
}
func (*binaryClassClass) DecodeTo(d binary.Decoder, obj binary.Object) {
	decodeClass(d, obj.(*Entity))
}

type binaryClassConstantSet struct{}

func (*ConstantSet) Class() binary.Class           { return (*binaryClassConstantSet)(nil) }
func (*binaryClassConstantSet) ID() binary.ID      { return binaryIDConstantSet }
func (*binaryClassConstantSet) New() binary.Object { return &ConstantSet{} }
func (*binaryClassConstantSet) Encode(e binary.Encoder, obj binary.Object) {
	encodeConstants(e, obj.(*ConstantSet))
}
func (*binaryClassConstantSet) Decode(d binary.Decoder) binary.Object {
	c := &ConstantSet{}
	decodeConstants(d, c)
	return c
}
func (*binaryClassConstantSet) DecodeTo(d binary.Decoder, obj binary.Object) {
	decodeConstants(d, obj.(*ConstantSet))
}
