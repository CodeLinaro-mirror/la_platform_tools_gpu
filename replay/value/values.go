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

package value

import (
	"math"

	"android.googlesource.com/platform/tools/gpu/replay/protocol"
)

// Bool is a Value of type TypeBool.
type Bool bool

// Type returns TypeBool.
func (v Bool) Type() protocol.Type { return protocol.TypeBool }

// Get returns 1 if the Bool is true, otherwise 0.
func (v Bool) Get(PointerResolver) uint64 {
	if v {
		return 1
	} else {
		return 0
	}
}

// U8 is a Value of type TypeUint8.
type U8 uint8

// Type returns TypeUint8.
func (v U8) Type() protocol.Type { return protocol.TypeUint8 }

// Get returns the value zero-extended to a uint64.
func (v U8) Get(PointerResolver) uint64 { return uint64(v) }

// S8 is a Value of type TypeInt8.
type S8 int8

// Type returns TypeInt8.
func (v S8) Type() protocol.Type { return protocol.TypeInt8 }

// Get returns the value sign-extended to a uint64.
func (v S8) Get(PointerResolver) uint64 { return uint64(v) }

// U16 is a Value of type TypeUint16.
type U16 uint16

// Type returns TypeUint16.
func (v U16) Type() protocol.Type { return protocol.TypeUint16 }

// Get returns the value zero-extended to a uint64.
func (v U16) Get(PointerResolver) uint64 { return uint64(v) }

// S16 is a Value of type TypeInt16.
type S16 int16

// Type returns TypeInt16.
func (v S16) Type() protocol.Type { return protocol.TypeInt16 }

// Get returns the value sign-extended to a uint64.
func (v S16) Get(PointerResolver) uint64 { return uint64(v) }

// F32 is a Value of type TypeFloat.
type F32 float32

// Type returns TypeFloat.
func (v F32) Type() protocol.Type { return protocol.TypeFloat }

// Get returns the IEEE 754 representation of the value packed into the low part of a uint64.
func (v F32) Get(PointerResolver) uint64 { return uint64(math.Float32bits(float32(v))) }

// U32 is a Value of type TypeUint32.
type U32 uint32

// Type returns TypeUint32.
func (v U32) Type() protocol.Type { return protocol.TypeUint32 }

// Get returns the value zero-extended to a uint64.
func (v U32) Get(PointerResolver) uint64 { return uint64(v) }

// S32 is a Value of type TypeInt32.
type S32 int32

// Type returns TypeInt32.
func (v S32) Type() protocol.Type { return protocol.TypeInt32 }

// Get returns the value sign-extended to a uint64.
func (v S32) Get(PointerResolver) uint64 { return uint64(v) }

// F64 is a Value of type TypeDouble.
type F64 float64

// Type returns TypeDouble.
func (v F64) Type() protocol.Type { return protocol.TypeDouble }

// Get returns the IEEE 754 representation of the value packed into a uint64.
func (v F64) Get(PointerResolver) uint64 { return math.Float64bits(float64(v)) }

// U64 is a Value of type TypeUint64.
type U64 uint64

// Type returns TypeUint64.
func (v U64) Type() protocol.Type { return protocol.TypeUint64 }

// Get returns the value zero-extended to a uint64.
func (v U64) Get(PointerResolver) uint64 { return uint64(v) }

// S64 is a Value of type TypeInt64.
type S64 int64

// Type returns TypeInt64.
func (v S64) Type() protocol.Type { return protocol.TypeInt64 }

// Get returns the value reinterpreted as a uint64.
func (v S64) Get(PointerResolver) uint64 { return uint64(v) }

// AbsolutePointer is a pointer in the absolute address-space that will not be
// altered before being passed to the protocol.
type AbsolutePointer uint64

// Type returns TypeAbsolutePointer.
func (p AbsolutePointer) Type() protocol.Type { return protocol.TypeAbsolutePointer }

// Get returns the uint64 value of the absolute pointer.
func (p AbsolutePointer) Get(PointerResolver) uint64 { return uint64(p) }

// Offset returns the sum of the pointer with offset.
func (p AbsolutePointer) Offset(offset uint64) Pointer {
	return p + AbsolutePointer(offset)
}

// IsValid returns true for all absolute pointers.
func (p AbsolutePointer) IsValid() bool { return true }

// VolatileCapturePointer is a pointer that was observed at capture time.
// Pointers of this type are remapped to an equivalent volatile address-space
// pointer before being passed to the protocol.
type VolatileCapturePointer uint64

// Type returns TypeVolatilePointer.
func (p VolatileCapturePointer) Type() protocol.Type { return protocol.TypeVolatilePointer }

// Get returns the observed pointer translated to an equivalent volatile
// address-space pointer.
func (p VolatileCapturePointer) Get(r PointerResolver) uint64 {
	return r.TranslateCapturePointer(uint64(p))
}

// Offset returns the sum of the pointer with offset.
func (p VolatileCapturePointer) Offset(offset uint64) Pointer {
	return p + VolatileCapturePointer(offset)
}

// IsValid returns true if the pointer considered valid. Currently this is a
// test for the pointer being greater than 0x10000 as low addresses are likely
// to be a wrong interpretation of the value. This may change in the future.
func (p VolatileCapturePointer) IsValid() bool {
	// Anything very low in applciation address-space is extremely
	// unlikely to be a valid pointer.
	return p > 0x10000
}

// VolatilePointer is a pointer to the volatile address-space.
// Unlike VolatileCapturePointer, there is no remapping.
type VolatilePointer uint64

// Type returns TypeVolatilePointer.
func (p VolatilePointer) Type() protocol.Type { return protocol.TypeVolatilePointer }

// Get returns the uint64 value of the pointer in volatile address-space.
func (p VolatilePointer) Get(PointerResolver) uint64 { return uint64(p) }

// Offset returns the sum of the pointer with offset.
func (p VolatilePointer) Offset(offset uint64) Pointer {
	return p + VolatilePointer(offset)
}

// IsValid returns true.
func (p VolatilePointer) IsValid() bool { return true }

// VolatileTemporaryPointer is a pointer to in temporary address-space.
// The temporary address-space sits within a reserved area of the the volatile
// address space and its offset is calculated dynamically.
type VolatileTemporaryPointer uint64

// Type returns TypeVolatilePointer.
func (p VolatileTemporaryPointer) Type() protocol.Type { return protocol.TypeVolatilePointer }

// Get returns the dynamically calculated offset of the temporary pointer within
// volatile address-space.
func (p VolatileTemporaryPointer) Get(r PointerResolver) uint64 {
	return r.TranslateTemporaryPointer(uint64(p))
}

// Offset returns the sum of the pointer with offset.
func (p VolatileTemporaryPointer) Offset(offset uint64) Pointer {
	return p + VolatileTemporaryPointer(offset)
}

// IsValid returns true.
func (p VolatileTemporaryPointer) IsValid() bool { return true }

// ConstantPointer is a pointer in the constant address-space that will not be
// altered before being passed to the protocol.
type ConstantPointer uint64

// Type returns TypeConstantPointer.
func (p ConstantPointer) Type() protocol.Type { return protocol.TypeConstantPointer }

// Get returns the uint64 value of the pointer in constant address-space.
func (p ConstantPointer) Get(PointerResolver) uint64 { return uint64(p) }

// Offset returns the sum of the pointer with offset.
func (p ConstantPointer) Offset(offset uint64) Pointer {
	return p + ConstantPointer(offset)
}

// IsValid returns true.
func (p ConstantPointer) IsValid() bool { return true }
