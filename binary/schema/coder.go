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

func EncodeType(e binary.Encoder, t binary.Type, compact bool) {
	switch t := t.(type) {
	case *Primitive:
		e.Uint8(uint8(PrimitiveTag) | (uint8(t.Method) << 4))
		if !compact {
			e.String(t.Name)
		}
	case *Struct:
		e.Uint8(uint8(StructTag))
		e.Entity(t.Entity, compact)
		if !compact {
			e.String(t.Relative)
		}
	case *Pointer:
		e.Uint8(uint8(PointerTag))
		EncodeType(e, t.Type, compact)
	case *Interface:
		e.Uint8(uint8(InterfaceTag))
		if !compact {
			e.String(t.Name)
		}
	case *Variant:
		e.Uint8(uint8(VariantTag))
		if !compact {
			e.String(t.Name)
		}
	case *Any:
		e.Uint8(uint8(AnyTag))
	case *Slice:
		e.Uint8(uint8(SliceTag))
		EncodeType(e, t.ValueType, compact)
		if !compact {
			e.String(t.Alias)
		}
	case *Array:
		e.Uint8(uint8(ArrayTag))
		e.Uint32(t.Size)
		EncodeType(e, t.ValueType, compact)
		if !compact {
			e.String(t.Alias)
		}
	case *Map:
		e.Uint8(uint8(MapTag))
		EncodeType(e, t.KeyType, compact)
		EncodeType(e, t.ValueType, compact)
		if !compact {
			e.String(t.Alias)
		}
	default:
		panic(fmt.Errorf("Encode unknown type %T", t))
	}
}

func DecodeType(d binary.Decoder, compact bool) binary.Type {
	tag := TypeTag(d.Uint8())
	switch tag & 0xf {
	case PrimitiveTag:
		t := &Primitive{}
		t.Method = Method(tag >> 4)
		if !compact {
			t.Name = d.String()
		}
		return t
	case StructTag:
		t := &Struct{}
		t.Entity = d.Entity(compact)
		if !compact {
			t.Relative = d.String()
		}
		return t
	case PointerTag:
		t := &Pointer{}
		t.Type = DecodeType(d, compact)
		return t
	case InterfaceTag:
		t := &Interface{}
		if !compact {
			t.Name = d.String()
		}
		return t
	case VariantTag:
		t := &Variant{}
		if !compact {
			t.Name = d.String()
		}
		return t
	case AnyTag:
		return &Any{}
	case SliceTag:
		t := &Slice{}
		t.ValueType = DecodeType(d, compact)
		if !compact {
			t.Alias = d.String()
		}
		return t
	case ArrayTag:
		t := &Array{}
		t.Size = d.Uint32()
		t.ValueType = DecodeType(d, compact)
		if !compact {
			t.Alias = d.String()
		}
		return t
	case MapTag:
		t := &Map{}
		t.KeyType = DecodeType(d, compact)
		t.ValueType = DecodeType(d, compact)
		if !compact {
			t.Alias = d.String()
		}
		return t
	default:
		panic(fmt.Errorf("Decode unknown type %v", tag))
	}
}

func EncodeEntity(e binary.Encoder, c *binary.Entity, compact bool) {
	e.String(c.Package)
	e.String(c.Identity)
	e.String(c.Version)
	if !compact {
		e.String(c.Display)
		e.ID(c.TypeID)
		e.Bool(c.Exported)
	}
	e.Uint32(uint32(len(c.Fields)))
	for _, f := range c.Fields {
		EncodeType(e, f.Type, compact)
		if !compact {
			e.String(f.Declared)
		}
	}
	if !compact {
		e.Uint32(uint32(len(c.Metadata)))
		for _, m := range c.Metadata {
			e.Object(m)
		}
	}
}

func DecodeEntity(d binary.Decoder, c *binary.Entity, compact bool) {
	c.Package = d.String()
	c.Identity = d.String()
	c.Version = d.String()
	if !compact {
		c.Display = d.String()
		c.TypeID = d.ID()
		c.Exported = d.Bool()
	}
	c.Fields = make(binary.FieldList, d.Uint32())
	for i := range c.Fields {
		c.Fields[i].Type = DecodeType(d, compact)
		if !compact {
			c.Fields[i].Declared = d.String()
		}
	}
	if !compact {
		c.Metadata = make([]binary.Object, d.Uint32())
		for i := range c.Metadata {
			c.Metadata[i] = d.Object()
		}
	}
}

func EncodeConstants(e binary.Encoder, c *ConstantSet) {
	EncodeType(e, c.Type, false)
	e.Uint32(uint32(len(c.Entries)))
	for _, entry := range c.Entries {
		e.String(entry.Name)
		c.Type.EncodeValue(e, entry.Value)
	}
}

func DecodeConstants(d binary.Decoder, c *ConstantSet) {
	c.Type = DecodeType(d, false)
	c.Entries = make([]Constant, d.Uint32())
	for i := range c.Entries {
		c.Entries[i].Name = d.String()
		c.Entries[i].Value = c.Type.DecodeValue(d)
	}
}
