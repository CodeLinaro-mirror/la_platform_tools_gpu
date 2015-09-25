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
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
)

// FramebufferObservation is an Atom that holds a snapshot of the color-buffer
// of the bound framebuffer at the time of capture. These atoms can be used to
// verify that replay gave the same results as what was captured.
type FramebufferObservation struct {
	binary.Generate `java:"disable"`
	Width, Height   uint32 // Framebuffer dimensions in pixels
	Data            []byte // The RGBA color-buffer data
}

func (a *FramebufferObservation) String() string {
	return fmt.Sprintf("FramebufferObservation %dx%d", a.Width, a.Height)
}

// Atom compliance
func (a *FramebufferObservation) API() gfxapi.ID              { return gfxapi.ID{} }
func (a *FramebufferObservation) Flags() Flags                { return 0 }
func (a *FramebufferObservation) Observations() *Observations { return &Observations{} }
func (a *FramebufferObservation) Mutate(s *gfxapi.State, d database.Database, l log.Logger) error {
	return nil
}

func init() {
	s := schema.Of((*FramebufferObservation)(nil).Class())
	s.Metadata = append(s.Metadata, &Metadata{
		API:              gfxapi.ID{},
		DisplayName:      "FramebufferObservation",
		DrawCall:         false,
		EndOfFrame:       false,
		DocumentationUrl: "[]",
	})
}
