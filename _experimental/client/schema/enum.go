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

package schema

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/service"
)

// EnumValue is a schema-typed Enum value. Unlike EnumEntry, EnumValue can hold
// values outside of the acceptable enum values.
type EnumValue struct {
	Type  *service.EnumInfo // The enum this value belongs to.
	Value uint32            // The numerical value.
}

// AllEnumEntries returns the flattened list of enum entries for the given
// EnumInfo and all enums it extends.
func AllEnumEntries(enum *service.EnumInfo) service.EnumEntryArray {
	entries := append(service.EnumEntryArray{}, enum.Entries...)
	for _, base := range enum.Extends {
		entries = append(entries, AllEnumEntries(base)...)
	}
	return entries
}

func (e EnumValue) String() string {
	entries := AllEnumEntries(e.Type)
	idx := IndexOfValue(entries, e.Value)
	if idx >= 0 {
		return entries[idx].Name
	} else {
		return fmt.Sprintf("%s<%d>", e.Type.Name, e.Value)
	}
}

// IndexOfValue returns the index of the enum entry with the specified value, or
// -1 if no value exists in the array.
func IndexOfValue(l service.EnumEntryArray, value uint32) int {
	for i, e := range l {
		if e.Value == value {
			return i
		}
	}
	return -1
}
