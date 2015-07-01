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

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// ErrUnboxable is returned when a non-boxable value type is passed to Box.
type ErrUnboxable struct {
	Value interface{} // The value that could not be encoded.
}

// Error returns the error message.
func (e ErrUnboxable) Error() string {
	return fmt.Sprintf("Value of type %T is not boxable", e.Value)
}

// ErrNotBoxedValue is returned when an Object is passed to Unbox that was not
// previously returned by a call to Box.
type ErrNotBoxedValue struct {
	Object binary.Object // The object that is not a boxed value.
}

// Error returns the error message.
func (e ErrNotBoxedValue) Error() string {
	return fmt.Sprintf("Object of type %T was boxed by a call to Box", e.Object)
}

type object_ struct {
	binary.Generate
	value binary.Object
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

type string_ struct {
	binary.Generate
	value string
}

type any interface {
	unbox() interface{}
}

func (v object_) unbox() interface{}  { return v.value }
func (v bool_) unbox() interface{}    { return v.value }
func (v uint8_) unbox() interface{}   { return v.value }
func (v int8_) unbox() interface{}    { return v.value }
func (v uint16_) unbox() interface{}  { return v.value }
func (v int16_) unbox() interface{}   { return v.value }
func (v float32_) unbox() interface{} { return v.value }
func (v uint32_) unbox() interface{}  { return v.value }
func (v int32_) unbox() interface{}   { return v.value }
func (v float64_) unbox() interface{} { return v.value }
func (v uint64_) unbox() interface{}  { return v.value }
func (v int64_) unbox() interface{}   { return v.value }
func (v string_) unbox() interface{}  { return v.value }

// Box returns v wrapped by a struct implementing binary.Object.
// If v is not boxable then ErrUnboxable is returned.
func Box(v interface{}) (binary.Object, error) {
	switch v := v.(type) {
	case binary.Object:
		return &object_{value: v}, nil
	case bool:
		return &bool_{value: v}, nil
	case uint8:
		return &uint8_{value: v}, nil
	case int8:
		return &int8_{value: v}, nil
	case uint16:
		return &uint16_{value: v}, nil
	case int16:
		return &int16_{value: v}, nil
	case float32:
		return &float32_{value: v}, nil
	case uint32:
		return &uint32_{value: v}, nil
	case int32:
		return &int32_{value: v}, nil
	case float64:
		return &float64_{value: v}, nil
	case uint64:
		return &uint64_{value: v}, nil
	case int64:
		return &int64_{value: v}, nil
	case string:
		return &string_{value: v}, nil
	}
	return nil, ErrUnboxable{Value: v}
}

// Unbox returns the value in o wrapped by a call to Box.
// If o is not a boxed value ErrNotBoxedValue is returned.
func Unbox(o binary.Object) (interface{}, error) {
	if o, any := o.(any); any {
		return o.unbox(), nil
	}
	return nil, ErrNotBoxedValue{Object: o}
}

// Any is the schema Type descriptor for a field who's underlying type requires
// boxing and unboxing. The type is usually declared as an empty interface.
type Any struct {
	binary.Generate
}

func (i *Any) Basename() string {
	return "<any>"
}

func (i *Any) Typename() string {
	return "<any>"
}

func (i *Any) String() string {
	return "<any>"
}

func (i *Any) Encode(e binary.Encoder, value interface{}) error {
	if boxed, err := Box(value); err != nil {
		return err
	} else {
		return e.Variant(boxed)
	}
}

func (i *Any) Decode(d binary.Decoder) (interface{}, error) {
	if boxed, err := d.Variant(); err != nil {
		return nil, err
	} else {
		return Unbox(boxed)
	}
}

func (i *Any) Skip(d binary.Decoder) error {
	_, err := d.SkipVariant()
	return err
}
