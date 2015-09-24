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

package gles

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

// findIssues is an atom transform that detects issues when replaying the
// stream of atoms. Any issues that are found are written to all the chans in
// the slice out. Once the last issue is sent (if any) all the chans in out are
// closed.
type findIssues struct {
	state    *gfxapi.State
	database database.Database
	logger   log.Logger
	out      []chan<- replay.Issue
}

func newFindIssues(d database.Database, l log.Logger) *findIssues {
	return &findIssues{
		state:    gfxapi.NewState(),
		database: d,
		logger:   log.Enter(l, "findIssues"),
	}
}

// reportTo adds the chan c to the list of issue listeners.
func (t *findIssues) reportTo(c chan<- replay.Issue) { t.out = append(t.out, c) }

func (t *findIssues) onIssue(i atom.ID, s log.Severity, e error) {
	issue := replay.Issue{Atom: i, Severity: s, Error: e}
	for _, o := range t.out {
		o <- issue
	}
}

func (t *findIssues) Transform(i atom.ID, a atom.Atom, out atom.Writer) {
	if err := a.Mutate(t.state, t.database, t.logger); err != nil {
		t.onIssue(i, log.Error, err)
	}
	out.Write(i, a)
	// Check the result of glGetError after every command.
	out.Write(atom.NoID, replay.Custom(func(_ atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		ptr := b.AllocateTemporaryMemory(4)
		b.Call(funcInfoGlGetError)
		b.Store(ptr)
		b.Post(ptr, 4, builder.Postback(func(d binary.Decoder, err error) error {
			if err != nil {
				return err
			}
			v, err := d.Uint32()
			if err != nil {
				t.onIssue(i, log.Error, fmt.Errorf("Failed to decode glGetError postback: %v", err))
				return err
			}
			if e := GLenum(v); e != GLenum_GL_NO_ERROR {
				t.onIssue(i, log.Error, fmt.Errorf("glGetError() returned %s", e))
			}
			return nil
		}))
		return nil
	}))

	if a, ok := a.(*atom.FramebufferObservation); ok {
		// Check that the framebuffer matches the FramebufferObservation's image.
		w, h, _, err := getState(t.state).getFramebufferAttachmentSizeAndFmt(gfxapi.FramebufferAttachmentColor)
		if err != nil {
			t.onIssue(i, log.Error, fmt.Errorf("Failed to resolve framebuffer dimensions: %v", err))
			return
		}
		if a.Width != w || a.Height != h {
			t.onIssue(i, log.Error, fmt.Errorf("Framebuffer dimensions were not as expected. Expected: %dx%d, Got: %dx%d",
				a.Width, a.Height, w, h))
			return
		}
		postColorData(t.state, int32(w), int32(h), out, func(img replay.Image) {
			if img.Error != nil {
				t.onIssue(i, log.Error, fmt.Errorf("Failed to fetch framebuffer color: %v", img.Error))
				return
			}
			expected := &image.Image{Width: a.Width, Height: a.Height, Data: a.Data, Format: image.RGBA()}
			diff, err := image.Difference(img.Image, expected)
			if err != nil {
				t.onIssue(i, log.Error, fmt.Errorf("Could not compare FramebufferObservation: %v", err))
				return
			}

			const theshold = 0.01 // TODO: Add as option to query.

			if diff > theshold {
				t.onIssue(i, log.Error, fmt.Errorf("FramebufferObservation did not match replayed framebuffer. Difference: %v%%",
					diff*100))
			}
		})
	}
}

func (t *findIssues) Flush(out atom.Writer) {
	out.Write(atom.NoID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		b.Post(value.AbsolutePointer(0x0), 0, func(d binary.Decoder, err error) error {
			for _, c := range t.out {
				close(c)
			}
			return err
		})
		return nil
	}))
}
