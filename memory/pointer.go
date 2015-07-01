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

package memory

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Nullptr is a zero-address pointer in the application pool.
var Nullptr = Pointer{Pool: ApplicationPool}

// Pointer is the type representing a memory pointer.
type Pointer struct {
	binary.Generate
	Address uint64 // The memory address.
	Pool    PoolID // The memory pool.
}

// Offset returns the pointer offset by n bytes.
func (p Pointer) Offset(n uint64) Pointer {
	return Pointer{Address: p.Address + n, Pool: p.Pool}
}

// Range returns a Range of size s with the base of this pointer.
func (p Pointer) Range(s uint64) Range {
	return Range{Base: p.Address, Size: s}
}

func (p Pointer) String() string {
	return fmt.Sprintf("0x%.16x@%d", uint64(p.Address), p.Pool)
}
