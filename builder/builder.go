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
	"fmt"
	"sort"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

type builder struct {
	ReplayManager *replay.Manager
}

// New creates a database.builder which can hold a replayManager, potentially required to build request outputs.
func New() *builder {
	return &builder{}
}

// SetReplayManager assigns the given replayManager.
func (b *builder) SetReplayManager(replayManager *replay.Manager) {
	b.ReplayManager = replayManager
}

type replayRequest interface {
	build(*replay.Manager, database.Database, log.Logger, binary.Object) error
}

type standaloneRequest interface {
	build(database.Database, log.Logger, binary.Object) error
}

// Compliance with the database.builder interface.

// BuildResource builds the output of the given request and writes it to the given out.
func (b *builder) BuildResource(request interface{}, db database.Database, logger log.Logger, out binary.Object) error {
	if logger != nil {
		logger.Info("Building resource: %T %v", request, request)
	}

	switch ty := request.(type) {
	case replayRequest:
		return ty.build(b.ReplayManager, db, logger, out)
	case standaloneRequest:
		return ty.build(db, logger, out)
	default:
		return fmt.Errorf("Unknown builder request type: %T", request)
	}
}

// Version returns the builder's version number.
func (b *builder) Version() uint32 {
	return 0
}

func getAPI(captureID service.CaptureId, db database.Database, logger log.Logger) (gfxapi.API, error) {
	var capture service.Capture
	if err := db.Load(captureID.ID, logger, &capture); err != nil {
		return nil, err
	}
	return gfxapi.Find(capture.API), nil
}

func getAtoms(captureID service.CaptureId, db database.Database, logger log.Logger) (atom.List, service.SchemaId, error) {
	var capture service.Capture
	if err := db.Load(captureID.ID, logger, &capture); err != nil {
		return atom.List{}, service.SchemaId{}, err
	}
	var stream service.AtomStream
	if err := db.Load(capture.Atoms.ID, logger, &stream); err != nil {
		return atom.List{}, service.SchemaId{}, err
	}
	atomList, err := stream.List()
	if err != nil {
		return atom.List{}, service.SchemaId{}, err
	}
	return atomList, stream.Schema, nil
}

// getAtomFramebufferDimensions returns the framebuffer dimensions after a given atom in the given capture and context.
// The first call to getAtomFramebufferDimensions for a given capture/context will trigger a computation of the dimensions for
// all atoms of this capture/context, which will be cached to the database for subsequent calls, regardless of the given atom.
func getAtomFramebufferDimensions(captureID service.CaptureId, contextID atom.ContextID, after atom.ID,
	db database.Database, logger log.Logger) (width, height uint32, err error) {
	id, err := db.StoreRequest(&getCaptureFramebufferDimensions{
		Capture: captureID,
		Context: contextID,
	}, logger)
	if err != nil {
		return 0, 0, err
	}

	var captureFbDims captureFramebufferDimensions
	err = db.Load(id, logger, &captureFbDims)
	if err != nil {
		return 0, 0, err
	}

	idx := sort.Search(len(captureFbDims.Dimensions), func(x int) bool {
		return captureFbDims.Dimensions[x].From > after
	}) - 1
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

// build writes to out the captureFramebufferDimensions resource resulting from the given getCaptureFramebufferDimensions request.
func (request *getCaptureFramebufferDimensions) build(db database.Database, logger log.Logger, out binary.Object) error {
	atoms, _, err := getAtoms(request.Capture, db, logger)
	if err != nil {
		return err
	}

	api, err := getAPI(request.Capture, db, logger)
	if err != nil {
		return err
	}

	state := api.InitialState()
	mutator := api.StateMutator(state)

	var captureFbDims captureFramebufferDimensions
	var currentDims *atomFramebufferDimensions

	for i, a := range atoms {
		if a.ContextID() != request.Context {
			continue
		}
		mutator.Write(atom.ID(i), a)
		if currentDims == nil || a.Flags().IsDrawCall() || a.Flags().IsEndOfFrame() {
			width, height, err := state.GetFramebufferAttachmentSize(gfxapi.FramebufferAttachmentColor)
			if err != nil {
				continue
			}
			if currentDims == nil || width != currentDims.Width || height != currentDims.Height {
				currentDims = &atomFramebufferDimensions{From:atom.ID(i), Width:width, Height: height}
				captureFbDims.Dimensions = append(captureFbDims.Dimensions, *currentDims)
			}
		}
	}
	store.CopyResource(out, &captureFbDims)
	return nil
}
