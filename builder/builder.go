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

// Package builder implements builders for resources from requests
// typically stored in the database, optionally depending on replay outputs.
package builder

import (
	"bytes"
	"fmt"
	"sort"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// The list of captures currently imported.
// TODO: This needs to be moved to persistent storage.
var captures = []*path.Capture{}

// Context is the type that should be passed to the database constructor's
// buildContext parameter.
type Context struct {
	ReplayManager *replay.Manager
}

func encode(v binary.Object) ([]byte, error) {
	buf := &bytes.Buffer{}
	if err := cyclic.Encoder(vle.Writer(buf)).Object(v); err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func decode(data []byte) (binary.Object, error) {
	return cyclic.Decoder(vle.Reader(bytes.NewBuffer(data))).Object()
}

// extractResources returns a new atom list with all the resources extracted
// and placed into the database.
func extractResources(atoms atom.List, d database.Database, l log.Logger) (atom.List, error) {
	out := make(atom.List, 0, len(atoms))
	idmap := map[binary.ID]binary.ID{}
	for _, a := range atoms {
		switch a := a.(type) {
		case *atom.Resource:
			if id, err := database.Store(a.Data, d, l); err != nil {
				return nil, err
			} else {
				idmap[a.ID] = id
			}

		default:
			// Replace resource IDs from identifiers generated at capture time to
			// direct database identifiers. This avoids a database link indirection.
			observations := a.Observations()
			for i, r := range observations.Reads {
				if id, found := idmap[r.ID]; found {
					observations.Reads[i].ID = id
				}
			}
			for i, w := range observations.Writes {
				if id, found := idmap[w.ID]; found {
					observations.Writes[i].ID = id
				}
			}
			out = append(out, a)
		}
	}
	return out, nil
}

// ImportCapture builds a new capture containing atoms, stores it into db and
// returns the new capture identifier.
func ImportCapture(name string, atoms atom.List, d database.Database, l log.Logger) (*path.Capture, error) {
	atoms, err := extractResources(atoms, d, l)
	if err != nil {
		return nil, err
	}

	stream := &service.AtomStream{Atoms: atoms}
	streamID, err := database.Store(stream, d, l)
	if err != nil {
		return nil, err
	}

	// Gather all the APIs used by the capture
	apis := map[gfxapi.ID]struct{}{}
	apiIDs := []service.ApiID{}
	for _, a := range atoms {
		if api := a.API(); api.Valid() {
			if _, found := apis[api]; !found {
				apis[api] = struct{}{}
				apiIDs = append(apiIDs, service.ApiID(api))
			}
		}
	}

	capture := &service.Capture{
		Apis:  apiIDs,
		Name:  name,
		Atoms: service.AtomsID(streamID),
	}

	captureID, err := database.Store(capture, d, l)
	if err != nil {
		return nil, err
	}

	p := &path.Capture{ID: captureID}
	captures = append(captures, p)

	return p, nil
}

// Captures returns all the captures stored by the database by identifier.
func Captures(db database.Database, logger log.Logger) ([]*path.Capture, error) {
	return captures, nil
}

// getAtomFramebufferDimensions returns the framebuffer dimensions after a given
// atom in the given capture and context.
// The first call to getAtomFramebufferDimensions for a given capture/context
// will trigger a computation of the dimensions for all atoms of this
// capture, which will be cached to the database for subsequent calls,
// regardless of the given atom.
func getAtomFramebufferDimensions(after *path.Atom, d database.Database, l log.Logger) (width, height uint32, err error) {
	obj, err := database.Build(&getCaptureFramebufferDimensions{Capture: after.Atoms.Capture}, d, l)
	if err != nil {
		return 0, 0, err
	}
	captureFbDims := obj.(*captureFramebufferDimensions)
	idx := sort.Search(len(captureFbDims.Dimensions), func(x int) bool {
		return uint64(captureFbDims.Dimensions[x].From) > after.Index
	}) - 1

	if idx < 0 {
		return 0, 0, fmt.Errorf("No dimension records found after atom %d. FB dimension records = %d",
			after, len(captureFbDims.Dimensions))
	}

	return captureFbDims.Dimensions[idx].Width, captureFbDims.Dimensions[idx].Height, nil
}

func uniformScale(width, height, maxWidth, maxHeight uint32) (w, h uint32) {
	w, h = width, height
	scaleX, scaleY := float32(w)/float32(maxWidth), float32(h)/float32(maxHeight)
	if scaleX > 1.0 || scaleY > 1.0 {
		if scaleX > scaleY {
			w, h = uint32(float32(w)/scaleX), uint32(float32(h)/scaleX)
		} else {
			w, h = uint32(float32(w)/scaleY), uint32(float32(h)/scaleY)
		}
	}
	return w, h
}

// BuildLazy writes to out the captureFramebufferDimensions resource resulting
// from the given getCaptureFramebufferDimensions request.
func (request *getCaptureFramebufferDimensions) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	var id atom.ID
	defer func() {
		if err := recover(); err != nil {
			panic(fmt.Errorf("Panic at atom %d: %v", id, err))
		}
	}()

	atoms, err := ResolveAtoms(request.Capture.Atoms(), d, l)
	if err != nil {
		return nil, err
	}

	var currentDims *atomFramebufferDimensions

	captureFbDims := &captureFramebufferDimensions{}

	s := gfxapi.NewState()
	for i, a := range atoms {
		id = atom.ID(i)
		if err := a.Mutate(s, d, l); err != nil {
			log.Warningf(l, "Atom %d %v: %v", i, a, err)
		}
		if currentDims == nil || a.Flags().IsDrawCall() || a.Flags().IsEndOfFrame() {
			api := gfxapi.Find(a.API())
			width, height, err := api.GetFramebufferAttachmentSize(s, gfxapi.FramebufferAttachmentColor)
			if err != nil {
				log.Warningf(l, "GetFramebufferAttachmentSize at atom %d %T gave error: %v", i, a, err)
				continue
			}
			if currentDims == nil || width != currentDims.Width || height != currentDims.Height {
				currentDims = &atomFramebufferDimensions{From: id, Width: width, Height: height}
				captureFbDims.Dimensions = append(captureFbDims.Dimensions, *currentDims)
			}
		}
	}
	return captureFbDims, nil
}
