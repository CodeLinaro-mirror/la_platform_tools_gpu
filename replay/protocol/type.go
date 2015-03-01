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

package protocol

import "fmt"

// Type is one of the primitive types supported by the replay virtual machine.
type Type uint32

const (
	TypeBool            Type = 0  // A boolean type.
	TypeInt8            Type = 1  // A signed 8-bit integer type.
	TypeInt16           Type = 2  // A signed 16-bit integer type.
	TypeInt32           Type = 3  // A signed 32-bit integer type.
	TypeInt64           Type = 4  // A signed 64-bit integer type.
	TypeUint8           Type = 5  // An unsigned 8-bit integer type.
	TypeUint16          Type = 6  // An unsigned 16-bit integer type.
	TypeUint32          Type = 7  // An unsigned 32-bit integer type.
	TypeUint64          Type = 8  // An unsigned 64-bit integer type.
	TypeFloat           Type = 9  // A 32-bit floating-point number type.
	TypeDouble          Type = 10 // A 64-bit floating-point number type.
	TypeAbsolutePointer Type = 11 // A pointer type that is not remapped by the protocol.
	TypeConstantPointer Type = 12 // A pointer into the constant buffer space.
	TypeVolatilePointer Type = 13 // A pointer into the volatile buffer space.

	TypeVoid Type = 0xffffffff // A non-existant type. Not handled by the protocol.
)

// Size returns the size in bytes of the type. pointerSize is the size in bytes
// of a pointer for the target architecture.
func (t Type) Size(pointerSize int) int {
	switch t {
	case TypeBool:
		return 1
	case TypeInt8:
		return 1
	case TypeInt16:
		return 2
	case TypeInt32:
		return 4
	case TypeInt64:
		return 8
	case TypeUint8:
		return 1
	case TypeUint16:
		return 2
	case TypeUint32:
		return 4
	case TypeUint64:
		return 8
	case TypeFloat:
		return 4
	case TypeDouble:
		return 8
	case TypeAbsolutePointer:
		return pointerSize
	case TypeConstantPointer:
		return pointerSize
	case TypeVolatilePointer:
		return pointerSize
	case TypeVoid:
		return 0
	default:
		panic(fmt.Errorf("Unknown ValueType %v", t))
	}
}

// String returns the human-readable name of the type.
func (t Type) String() string {
	switch t {
	case TypeBool:
		return "bool"
	case TypeInt8:
		return "int8"
	case TypeInt16:
		return "int16"
	case TypeInt32:
		return "int32"
	case TypeInt64:
		return "int64"
	case TypeUint8:
		return "uint8"
	case TypeUint16:
		return "uint16"
	case TypeUint32:
		return "uint32"
	case TypeUint64:
		return "uint64"
	case TypeFloat:
		return "float"
	case TypeDouble:
		return "double"
	case TypeAbsolutePointer:
		return "absolute-pointer"
	case TypeConstantPointer:
		return "constant-pointer"
	case TypeVolatilePointer:
		return "volatile-pointer"
	case TypeVoid:
		return "void"
	default:
		panic(fmt.Errorf("Unknown ValueType %d", uint32(t)))
	}
}
