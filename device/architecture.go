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

package device

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary/endian"
)

// Architecture holds architecture information about a device.
type Architecture struct {
	PointerAlignment int              // The alignment in bytes of a pointer type.
	PointerSize      int              // The size in bytes of a pointer type.
	IntegerSize      int              // The size in bytes of a int or unsigned int.
	ByteOrder        endian.ByteOrder // The byte ordering for the target.
}

func (a Architecture) String() string {
	return fmt.Sprintf("Architecture{ alignof(void*): %d, sizeof(void*): %d, sizeof(int): %d, order: %v }",
		a.PointerAlignment, a.PointerSize, a.IntegerSize, a.ByteOrder)
}
