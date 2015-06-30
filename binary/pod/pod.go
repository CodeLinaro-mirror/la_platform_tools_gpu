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

// Package pod contains Object wrappers for Plain-Old-Data types.
package pod

import "android.googlesource.com/platform/tools/gpu/binary"

// Bool wraps a bool value into a binary.Object.
type Bool struct {
	binary.Generate
	Value bool
}

// Uint8 wraps a uint8 value into a binary.Object.
type Uint8 struct {
	binary.Generate
	Value uint8
}

// Int8 wraps a int8 value into a binary.Object.
type Int8 struct {
	binary.Generate
	Value int8
}

// Uint16 wraps a uint16 value into a binary.Object.
type Uint16 struct {
	binary.Generate
	Value uint16
}

// Int16 wraps a int16 value into a binary.Object.
type Int16 struct {
	binary.Generate
	Value int16
}

// Float32 wraps a float32 value into a binary.Object.
type Float32 struct {
	binary.Generate
	Value float32
}

// Uint32 wraps a uint32 value into a binary.Object.
type Uint32 struct {
	binary.Generate
	Value uint32
}

// Int32 wraps a int32 value into a binary.Object.
type Int32 struct {
	binary.Generate
	Value int32
}

// Float64 wraps a float64 value into a binary.Object.
type Float64 struct {
	binary.Generate
	Value float64
}

// Uint64 wraps a uint64 value into a binary.Object.
type Uint64 struct {
	binary.Generate
	Value uint64
}

// Int64 wraps a int64 value into a binary.Object.
type Int64 struct {
	binary.Generate
	Value int64
}

// String wraps a string value into a binary.Object.
type String struct {
	binary.Generate
	Value string
}

// Wrap returns v wrapped by a struct implementing binary.Object.
// If v already conforms to binary.Object then v is returned.
// If v is not a POD type, then the function returns nil.
func Wrap(v interface{}) binary.Object {
	switch v := v.(type) {
	case binary.Object:
		return v
	case bool:
		return &Bool{Value: v}
	case uint8:
		return &Uint8{Value: v}
	case int8:
		return &Int8{Value: v}
	case uint16:
		return &Uint16{Value: v}
	case int16:
		return &Int16{Value: v}
	case float32:
		return &Float32{Value: v}
	case uint32:
		return &Uint32{Value: v}
	case int32:
		return &Int32{Value: v}
	case float64:
		return &Float64{Value: v}
	case uint64:
		return &Uint64{Value: v}
	case int64:
		return &Int64{Value: v}
	case string:
		return &String{Value: v}
	}
	return nil
}

// Unwrap returns the POD value wrapped in the binary.Object.
// If v is not a POD type, then the function returns o unaltered.
func Unwrap(o binary.Object) interface{} {
	switch o := o.(type) {
	case *Bool:
		return o.Value
	case *Uint8:
		return o.Value
	case *Int8:
		return o.Value
	case *Uint16:
		return o.Value
	case *Int16:
		return o.Value
	case *Float32:
		return o.Value
	case *Uint32:
		return o.Value
	case *Int32:
		return o.Value
	case *Float64:
		return o.Value
	case *Uint64:
		return o.Value
	case *Int64:
		return o.Value
	case *String:
		return o.Value
	}
	return o
}
