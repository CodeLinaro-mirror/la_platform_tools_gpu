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

package client

import (
	"fmt"
	"unsafe"
)

type DataType interface {
	Name() string
	SizeBytes() int
	String() string
	Unknown() string
	Read(buffer []byte) DataType
}

type U8 uint8

func (v U8) Name() string                { return "u8" }
func (v U8) SizeBytes() int              { return 1 }
func (v U8) String() string              { return fmt.Sprintf("%.2x", uint8(v)) }
func (v U8) Unknown() string             { return "??" }
func (v U8) Read(buffer []byte) DataType { return *(*U8)(unsafe.Pointer(&buffer[0])) }

type S8 int8

func (v S8) Name() string                { return "s8" }
func (v S8) SizeBytes() int              { return 1 }
func (v S8) String() string              { return fmt.Sprintf("% .2x", int8(v)) }
func (v S8) Unknown() string             { return "???" }
func (v S8) Read(buffer []byte) DataType { return *(*S8)(unsafe.Pointer(&buffer[0])) }

type U16 uint16

func (v U16) Name() string                { return "u16" }
func (v U16) SizeBytes() int              { return 2 }
func (v U16) String() string              { return fmt.Sprintf("%.4x", uint16(v)) }
func (v U16) Unknown() string             { return "????" }
func (v U16) Read(buffer []byte) DataType { return *(*U16)(unsafe.Pointer(&buffer[0])) }

type S16 int16

func (v S16) Name() string                { return "s16" }
func (v S16) SizeBytes() int              { return 2 }
func (v S16) String() string              { return fmt.Sprintf("% .4x", int16(v)) }
func (v S16) Unknown() string             { return "?????" }
func (v S16) Read(buffer []byte) DataType { return *(*S16)(unsafe.Pointer(&buffer[0])) }

type F32 float32

func (v F32) Name() string                { return "f32" }
func (v F32) SizeBytes() int              { return 4 }
func (v F32) String() string              { return fmt.Sprintf("% 8f", float32(v)) }
func (v F32) Unknown() string             { return "?????????" }
func (v F32) Read(buffer []byte) DataType { return *(*F32)(unsafe.Pointer(&buffer[0])) }

type U32 uint32

func (v U32) Name() string                { return "u32" }
func (v U32) SizeBytes() int              { return 4 }
func (v U32) String() string              { return fmt.Sprintf("%.8x", uint32(v)) }
func (v U32) Unknown() string             { return "????????" }
func (v U32) Read(buffer []byte) DataType { return *(*U32)(unsafe.Pointer(&buffer[0])) }

type S32 int32

func (v S32) Name() string                { return "s32" }
func (v S32) SizeBytes() int              { return 4 }
func (v S32) String() string              { return fmt.Sprintf("% .8x", int32(v)) }
func (v S32) Unknown() string             { return "?????????" }
func (v S32) Read(buffer []byte) DataType { return *(*S32)(unsafe.Pointer(&buffer[0])) }

type F64 float64

func (v F64) Name() string                { return "f64" }
func (v F64) SizeBytes() int              { return 8 }
func (v F64) String() string              { return fmt.Sprintf("% 8f", float64(v)) }
func (v F64) Unknown() string             { return "?????????" }
func (v F64) Read(buffer []byte) DataType { return *(*F64)(unsafe.Pointer(&buffer[0])) }

type U64 uint64

func (v U64) Name() string                { return "u64" }
func (v U64) SizeBytes() int              { return 8 }
func (v U64) String() string              { return fmt.Sprintf("%.16x", uint64(v)) }
func (v U64) Unknown() string             { return "????????????????" }
func (v U64) Read(buffer []byte) DataType { return *(*U32)(unsafe.Pointer(&buffer[0])) }

type S64 int64

func (v S64) Name() string                { return "s64" }
func (v S64) SizeBytes() int              { return 8 }
func (v S64) String() string              { return fmt.Sprintf("% .16x", int64(v)) }
func (v S64) Unknown() string             { return "?????????????????" }
func (v S64) Read(buffer []byte) DataType { return *(*S32)(unsafe.Pointer(&buffer[0])) }

type ASCII uint8

func (v ASCII) Name() string                { return "ASCII" }
func (v ASCII) SizeBytes() int              { return 1 }
func (v ASCII) String() string              { return string(rune(v)) }
func (v ASCII) Unknown() string             { return "?" }
func (v ASCII) Read(buffer []byte) DataType { return *(*ASCII)(unsafe.Pointer(&buffer[0])) }
