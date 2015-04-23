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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
)

const TypeIDResource TypeID = 0xfffd

func init() {
	Register(TypeInfo{ID: TypeIDResource, New: func() Atom { return &Resource{} }})
}

// Resource is an Atom that embeds a blob of memory into the stream. These atoms
// are typically only used for .gfxtrace files as they are stripped from the
// stream on import and placed into the database.
type Resource struct {
	binary.Generate
	ResourceID binary.ID // The resource identifier holding the memory that was observed.
	Data       []byte    // The resource data
}

func (a *Resource) String() string {
	return fmt.Sprintf("ResID: %s - 0x%x bytes", a.ResourceID, len(a.Data))
}

// Atom compliance
func (a *Resource) API() gfxapi.API {
	return nil
}

func (a *Resource) TypeID() TypeID {
	return TypeIDResource
}

func (a *Resource) Flags() Flags {
	return 0
}

func (a *Resource) Mutate(s *gfxapi.State) error {
	return nil
}
