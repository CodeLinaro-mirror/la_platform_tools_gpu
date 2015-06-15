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

package schema

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/service"
)

// ReadType reads the schema type ty from the decoder d.
func ReadType(ty service.TypeInfo, d binary.Decoder) (interface{}, error) {
	var err error
	switch ty.GetKind() {
	case service.TypeKindBool:
		return d.Bool()
	case service.TypeKindS8:
		return d.Int8()
	case service.TypeKindU8:
		return d.Uint8()
	case service.TypeKindS16:
		return d.Int16()
	case service.TypeKindU16:
		return d.Uint16()
	case service.TypeKindS32:
		return d.Int32()
	case service.TypeKindU32:
		return d.Uint32()
	case service.TypeKindF32:
		return d.Float32()
	case service.TypeKindS64:
		return d.Int64()
	case service.TypeKindU64:
		return d.Uint64()
	case service.TypeKindF64:
		return d.Float64()
	case service.TypeKindString:
		return d.String()
	case service.TypeKindEnum:
		value, err := d.Uint32()
		return EnumValue{
			Type:  ty.(*service.EnumInfo),
			Value: value,
		}, err
	case service.TypeKindStruct:
		info := ty.(*service.StructInfo)
		str := Struct{Type: info}
		for _, f := range info.Fields {
			field := Field{Info: f}
			field.Value, err = ReadType(f.Type, d)
			if err != nil {
				return nil, err
			}
			str.Fields = append(str.Fields, field)
		}
		return str, nil
	case service.TypeKindClass:
		info := ty.(*service.ClassInfo)
		class := Class{Type: info}
		for _, f := range info.Fields {
			field := Field{Info: f}
			field.Value, err = ReadType(f.Type, d)
			if err != nil {
				return nil, err
			}
			class.Fields = append(class.Fields, field)
		}
		return class, nil
	case service.TypeKindArray:
		info := ty.(*service.ArrayInfo)
		count, err := d.Uint32()
		if err != nil {
			return nil, err
		}
		array := Array{
			Type:     info,
			Elements: make([]ArrayElement, count),
		}
		for i := range array.Elements {
			array.Elements[i], err = ReadType(info.ElementType, d)
			if err != nil {
				return array, err
			}
		}
		return array, nil
	case service.TypeKindStaticArray:
		info := ty.(*service.StaticArrayInfo)
		staticArray := StaticArray{
			Type:     info,
			Elements: make([]ArrayElement, info.Size),
		}
		for i := range staticArray.Elements {
			staticArray.Elements[i], err = ReadType(info.ElementType, d)
			if err != nil {
				return staticArray, err
			}
		}
		return staticArray, nil
	case service.TypeKindMap:
		info := ty.(*service.MapInfo)
		count, err := d.Uint32()
		if err != nil {
			return nil, err
		}
		m := Map{
			Type:     info,
			Elements: make([]MapElement, count),
		}
		for i := range m.Elements {
			m.Elements[i].Key, err = ReadType(info.KeyType, d)
			if err != nil {
				return m, err
			}
			m.Elements[i].Value, err = ReadType(info.ValueType, d)
			if err != nil {
				return m, err
			}
		}
		return m, nil
	case service.TypeKindPointer:
		addr, err := d.Uint64()
		if err != nil {
			return nil, nil
		}
		_, err = d.Uint32() // Pool - currently not used
		return memory.Pointer(addr), err
	case service.TypeKindMemory:
		// TODO: Memory currently has no payload
		return nil, nil
	case service.TypeKindAny:
		// TODO: Not drop any values on the floor.
		return nil, nil
	case service.TypeKindID:
		return d.ID()
	default:
		return nil, fmt.Errorf("Unknown kind %v", ty.GetKind())
	}
}

// WriteType writes v to the encoder e. The value v must be one of the following
// types:
//   bool
//   int8
//   uint8
//   int16
//   uint16
//   int32
//   uint32
//   float32
//   int64
//   uint64
//   float64
//   string
//   EnumValue
//   Struct
//   Class
//   Array
//   Map
//   memory.Pointer
func WriteType(v interface{}, e binary.Encoder) error {
	switch v := v.(type) {
	case bool:
		return e.Bool(v)
	case int8:
		return e.Int8(v)
	case uint8:
		return e.Uint8(v)
	case int16:
		return e.Int16(v)
	case uint16:
		return e.Uint16(v)
	case int32:
		return e.Int32(v)
	case uint32:
		return e.Uint32(v)
	case float32:
		return e.Float32(v)
	case int64:
		return e.Int64(v)
	case uint64:
		return e.Uint64(v)
	case float64:
		return e.Float64(v)
	case string:
		return e.String(v)
	case EnumValue:
		return e.Uint32(v.Value)
	case Struct:
		for _, f := range v.Fields {
			if err := WriteType(f.Value, e); err != nil {
				return err
			}
		}
	case Class:
		for _, f := range v.Fields {
			if err := WriteType(f.Value, e); err != nil {
				return err
			}
		}
	case Array:
		if err := e.Int32(int32(len(v.Elements))); err != nil {
			return err
		}
		for _, el := range v.Elements {
			if err := WriteType(el, e); err != nil {
				return err
			}
		}
		return nil
	case Map:
		if err := e.Int32(int32(len(v.Elements))); err != nil {
			return err
		}
		for _, el := range v.Elements {
			if err := WriteType(el.Key, e); err != nil {
				return err
			}
			if err := WriteType(el.Value, e); err != nil {
				return err
			}
		}
		return nil
	case memory.Pointer:
		if err := e.Uint64(uint64(v)); err != nil {
			return err
		}
		return e.Uint32(0) // Pool - currently not used
	default:
		return fmt.Errorf("Schema does not support encoding type %T", v)
	}
	return nil
}
