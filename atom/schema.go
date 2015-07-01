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

package atom

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

// Metadata is the meta information about an atom type that is added to the
// binary schema class for the atom.
type Metadata struct {
	binary.Generate
	Api              binary.ID // The api this atom belongs to.
	Flags            Flags     // The atom flags for this type.
	DocumentationUrl string    // A url for documentation about this atom.
}

// Finds the atom metadata for the given schema class.
// Returns nil if the class was not for an atom.
func FindMetadata(class *schema.Class) *Metadata {
	for _, m := range class.Metadata {
		if meta, ok := m.(*Metadata); ok {
			return meta
		}
	}
	return nil
}
