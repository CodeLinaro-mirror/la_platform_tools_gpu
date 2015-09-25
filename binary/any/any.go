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

// Package any contains Object wrappers for Plain-Old-Data types.
package any

import "android.googlesource.com/platform/tools/gpu/binary"

// binary: java.source = base/rpclib/src/main/java
// binary: java.package = com.android.tools.rpclib.any
// binary: java.indent = "    "
// binary: java.member_prefix = m

type object_ struct {
	binary.Generate `java:"ObjectBox"`
	value           binary.Object
}

type bool_ struct {
	binary.Generate
	value bool
}

type uint8_ struct {
	binary.Generate
	value uint8
}

type int8_ struct {
	binary.Generate
	value int8
}

type uint16_ struct {
	binary.Generate
	value uint16
}

type int16_ struct {
	binary.Generate
	value int16
}

type float32_ struct {
	binary.Generate
	value float32
}

type uint32_ struct {
	binary.Generate
	value uint32
}

type int32_ struct {
	binary.Generate
	value int32
}

type float64_ struct {
	binary.Generate
	value float64
}

type uint64_ struct {
	binary.Generate
	value uint64
}

type int64_ struct {
	binary.Generate
	value int64
}

type int_ struct {
	binary.Generate `java:"disable"`
	value           int
}

type string_ struct {
	binary.Generate `java:"StringBox"`
	value           string
}

func (v object_) Unbox() interface{}  { return v.value }
func (v bool_) Unbox() interface{}    { return v.value }
func (v uint8_) Unbox() interface{}   { return v.value }
func (v int8_) Unbox() interface{}    { return v.value }
func (v uint16_) Unbox() interface{}  { return v.value }
func (v int16_) Unbox() interface{}   { return v.value }
func (v float32_) Unbox() interface{} { return v.value }
func (v uint32_) Unbox() interface{}  { return v.value }
func (v int32_) Unbox() interface{}   { return v.value }
func (v float64_) Unbox() interface{} { return v.value }
func (v uint64_) Unbox() interface{}  { return v.value }
func (v int64_) Unbox() interface{}   { return v.value }
func (v int_) Unbox() interface{}     { return v.value }
func (v string_) Unbox() interface{}  { return v.value }

type objectSlice struct {
	binary.Generate
	value []binary.Object
}

type boolSlice struct {
	binary.Generate
	value []bool
}

type uint8Slice struct {
	binary.Generate
	value []uint8
}

type int8Slice struct {
	binary.Generate
	value []int8
}

type uint16Slice struct {
	binary.Generate
	value []uint16
}

type int16Slice struct {
	binary.Generate
	value []int16
}

type float32Slice struct {
	binary.Generate
	value []float32
}

type uint32Slice struct {
	binary.Generate
	value []uint32
}

type int32Slice struct {
	binary.Generate
	value []int32
}

type float64Slice struct {
	binary.Generate
	value []float64
}

type uint64Slice struct {
	binary.Generate
	value []uint64
}

type int64Slice struct {
	binary.Generate
	value []int64
}

type intSlice struct {
	binary.Generate `java:"disable"`
	value           []int
}

type stringSlice struct {
	binary.Generate
	value []string
}

func (v objectSlice) Unbox() interface{}  { return v.value }
func (v boolSlice) Unbox() interface{}    { return v.value }
func (v uint8Slice) Unbox() interface{}   { return v.value }
func (v int8Slice) Unbox() interface{}    { return v.value }
func (v uint16Slice) Unbox() interface{}  { return v.value }
func (v int16Slice) Unbox() interface{}   { return v.value }
func (v float32Slice) Unbox() interface{} { return v.value }
func (v uint32Slice) Unbox() interface{}  { return v.value }
func (v int32Slice) Unbox() interface{}   { return v.value }
func (v float64Slice) Unbox() interface{} { return v.value }
func (v uint64Slice) Unbox() interface{}  { return v.value }
func (v int64Slice) Unbox() interface{}   { return v.value }
func (v intSlice) Unbox() interface{}     { return v.value }
func (v stringSlice) Unbox() interface{}  { return v.value }

func boxer(v interface{}) binary.Object {
	switch v := v.(type) {
	case binary.Object:
		return &object_{value: v}
	case bool:
		return &bool_{value: v}
	case uint8:
		return &uint8_{value: v}
	case int8:
		return &int8_{value: v}
	case uint16:
		return &uint16_{value: v}
	case int16:
		return &int16_{value: v}
	case float32:
		return &float32_{value: v}
	case uint32:
		return &uint32_{value: v}
	case int32:
		return &int32_{value: v}
	case float64:
		return &float64_{value: v}
	case uint64:
		return &uint64_{value: v}
	case int64:
		return &int64_{value: v}
	case int:
		return &int_{value: v}
	case string:
		return &string_{value: v}

	case []binary.Object:
		return &objectSlice{value: v}
	case []bool:
		return &boolSlice{value: v}
	case []uint8:
		return &uint8Slice{value: v}
	case []int8:
		return &int8Slice{value: v}
	case []uint16:
		return &uint16Slice{value: v}
	case []int16:
		return &int16Slice{value: v}
	case []float32:
		return &float32Slice{value: v}
	case []uint32:
		return &uint32Slice{value: v}
	case []int32:
		return &int32Slice{value: v}
	case []float64:
		return &float64Slice{value: v}
	case []uint64:
		return &uint64Slice{value: v}
	case []int64:
		return &int64Slice{value: v}
	case []int:
		return &intSlice{value: v}
	case []string:
		return &stringSlice{value: v}
	default:
		return nil
	}
}

func init() {
	binary.RegisterBoxer(boxer)
}
