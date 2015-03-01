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
// See the License for the specific language governing permissions ands
// limitations under the License.

//go:generate codergen -go=payload_code.go payload.go

package protocol

import "android.googlesource.com/platform/tools/gpu/binary"

// Hack to make codergen use the binary.Data Encode/Decode methods
// rather than encoding and decoding a byte at a time.
// BUG: b/19474821
type Data struct {
	binary.Data
}

// ResourceInfo describes a resource used by a Payload.
type ResourceInfo struct {
	ID   string // The resource identifier as a string.
	Size uint32 // The size in bytes of the resource.
}

// Payload contains all the information to perform a replay. The encoded form
// is what is passed to the replay system.
type Payload struct {
	StackSize          uint32         // Maximum number of values.
	VolatileMemorySize uint32         // In bytes.
	Constants          Data           // The constant buffer.
	Resources          []ResourceInfo // Resources used by this replay payload.
	Opcodes            Data           // The encoded list of opcodes.
}
