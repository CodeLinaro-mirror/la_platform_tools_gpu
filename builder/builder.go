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
	"android.googlesource.com/platform/tools/gpu/gfxapi/schema"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

// The root database ID to a captures object, storing the identifiers to all
// captures that are held in the database.
var captureIdsDatabaseID = binary.NewID([]byte("ALL CAPTURES VERSION 1.0"))

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
		logger.Infof("Building resource: %T %v", request, request)
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

// extractResources returns a new atom list with all the resources extracted
// and placed into the database.
func extractResources(atoms atom.List, db database.Database, logger log.Logger) (atom.List, error) {
	out := make(atom.List, 0, len(atoms))
	idmap := map[binary.ID]binary.ID{}
	for _, a := range atoms {
		switch a := a.(type) {
		case *atom.Resource:
			blob := store.Blob{Data: a.Data}
			if id, err := db.Store(&blob, logger); err != nil {
				return nil, err
			} else {
				idmap[a.ResourceID] = id
			}

		case *atom.Observation:
			id := a.ResourceID
			if remapped, ok := idmap[id]; ok { // Avoid a db link and remap ids here.
				id = remapped
			}
			out = append(out, &atom.Observation{
				Range:      a.Range,
				ResourceID: id,
			})

		default:
			out = append(out, a)
		}
	}
	return out, nil
}

// ImportCapture builds a new capture containing atoms, stores it into db and
// returns the new capture identifier.
func ImportCapture(name string, atoms atom.List, db database.Database, logger log.Logger) (service.CaptureId, error) {
	atoms, err := extractResources(atoms, db, logger)
	if err != nil {
		return service.CaptureId{}, err
	}

	stream, err := service.NewAtomStream(atoms)
	if err != nil {
		return service.CaptureId{}, err
	}

	streamID, err := db.Store(&stream, logger)
	if err != nil {
		return service.CaptureId{}, err
	}

	schema := schema.Schema()
	schemaID, err := db.Store(&schema, logger)
	if err != nil {
		return service.CaptureId{}, err
	}

	// Gather all the APIs used by the capture
	apis := map[gfxapi.API]struct{}{}
	apiIDs := service.ApiIdArray{}
	for _, a := range atoms {
		if api := a.API(); api != nil {
			if _, found := apis[api]; !found {
				apis[api] = struct{}{}
				apiIDs = append(apiIDs, service.ApiId{ID: binary.ID(api.ID())})
			}
		}
	}

	capture := service.Capture{
		Apis:   apiIDs,
		Name:   name,
		Atoms:  service.AtomStreamId{ID: streamID},
		Schema: service.SchemaId{ID: schemaID},
	}

	id, err := db.Store(&capture, logger)
	if err != nil {
		return service.CaptureId{}, err
	}

	// TODO: This is a read-modify-write operation with no safty for concurrent
	// writes! A mutex in this function would provided limited safety as the
	// database could be modifying the captures list elsewhere. We really need
	// atomic write support in the database. b/19889089.

	// Load the list of captures stored in the database
	ids, _ := Captures(db, logger)

	for _, i := range ids {
		if i.ID == id {
			// Capture already imported
			return service.CaptureId{ID: id}, nil
		}
	}

	// Add the capture into the list of captures stored by the database.
	c := captures{ids: ids}
	c.ids = append(c.ids, service.CaptureId{ID: id})

	record, err := db.Store(&c, logger)
	if err != nil {
		return service.CaptureId{}, err
	}
	if err := db.StoreLink(record, captureIdsDatabaseID, logger); err != nil {
		return service.CaptureId{}, err
	}

	return service.CaptureId{ID: id}, nil
}

// Captures returns all the captures stored by the database.
func Captures(db database.Database, logger log.Logger) (service.CaptureIdArray, error) {
	var c captures
	err := db.Load(captureIdsDatabaseID, logger, &c)
	return c.ids, err
}

func loadCapture(captureID service.CaptureId, db database.Database, logger log.Logger) (service.Capture, error) {
	var capture service.Capture
	if err := db.Load(captureID.ID, logger, &capture); err != nil {
		return service.Capture{}, err
	}
	return capture, nil
}

func loadAtoms(streamID service.AtomStreamId, db database.Database, logger log.Logger) (atom.List, error) {
	var stream service.AtomStream
	if err := db.Load(streamID.ID, logger, &stream); err != nil {
		return atom.List{}, err
	}
	atomList, err := stream.List()
	if err != nil {
		return atom.List{}, err
	}
	return atomList, nil
}

// getAtomFramebufferDimensions returns the framebuffer dimensions after a given atom in the given capture and context.
// The first call to getAtomFramebufferDimensions for a given capture/context will trigger a computation of the dimensions for
// all atoms of this capture/context, which will be cached to the database for subsequent calls, regardless of the given atom.
func getAtomFramebufferDimensions(captureID service.CaptureId, after atom.ID, db database.Database, logger log.Logger) (width, height uint32, err error) {
	id, err := db.StoreRequest(&getCaptureFramebufferDimensions{Capture: captureID}, logger)
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

// build writes to out the captureFramebufferDimensions resource resulting from the given getCaptureFramebufferDimensions request.
func (request *getCaptureFramebufferDimensions) build(db database.Database, logger log.Logger, out binary.Object) error {
	capture, err := loadCapture(request.Capture, db, logger)
	if err != nil {
		return err
	}

	atoms, err := loadAtoms(capture.Atoms, db, logger)
	if err != nil {
		return err
	}

	var captureFbDims captureFramebufferDimensions
	var currentDims *atomFramebufferDimensions

	s := &gfxapi.State{}
	for i, a := range atoms {
		if err := a.Mutate(s); err != nil {
			return err
		}
		if currentDims == nil || a.Flags().IsDrawCall() || a.Flags().IsEndOfFrame() {
			api := a.API()
			width, height, err := api.GetFramebufferAttachmentSize(s, gfxapi.FramebufferAttachmentColor)
			if err != nil {
				logger.Warningf("GetFramebufferAttachmentSize at atom %d %T gave error: %v", i, a, err)
				continue
			}
			if currentDims == nil || width != currentDims.Width || height != currentDims.Height {
				currentDims = &atomFramebufferDimensions{From: atom.ID(i), Width: width, Height: height}
				captureFbDims.Dimensions = append(captureFbDims.Dimensions, *currentDims)
			}
		}
	}
	store.CopyResource(out, &captureFbDims)
	return nil
}
