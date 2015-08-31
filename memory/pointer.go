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
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// Nullptr is a zero-address pointer in the application pool.
var Nullptr = Pointer{Pool: ApplicationPool}

// Pointer is the type representing a memory pointer.
type Pointer struct {
	binary.Generate `java:"MemoryPointer"`
	Address         uint64 // The memory address.
	Pool            PoolID // The memory pool.
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

// Link return the path to the memory pointed-to by p.
func (p Pointer) Link(pth path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if a := path.FindAtom(pth); a != nil {
		return a.MemoryAfter(uint64(p.Pool), p.Address, 0), nil
	}
	return nil, nil
}
