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
	"errors"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/value"
	"android.googlesource.com/platform/tools/gpu/service"
)

// timingInfoGpu returns a transform used to measure and retrieve frame and draw call GPU timing information.
// Note: this is experimental and is likely to change in the near future.
func timingInfoGpu(flags service.TimingFlags, out chan<- replay.CallTiming, device *service.Device, db database.Database, logger log.Logger) atom.Transformer {
	timeFrames := flags&service.TimingPerFrame != 0
	timeDrawCalls := flags&service.TimingPerDrawCall != 0

	seen := map[atom.ID]struct{}{}
	timingInfo := service.TimingInfo{}

	// TODO: move TimingInfo construction outside of the predicate.
	predicate := func(id atom.ID, a atom.Atom) (bool, bool) {
		flush := a.Flags().IsEndOfFrame()
		mark := false
		if timeFrames && a.Flags().IsEndOfFrame() {
			if _, ok := seen[id]; !ok {
				timingInfo.PerFrame = append(timingInfo.PerFrame, service.AtomRangeTimer{ToAtomIndex: uint64(id)})
			}
			mark = true
		}
		if timeDrawCalls && a.Flags().IsDrawCall() {
			if _, ok := seen[id]; !ok {
				timingInfo.PerDrawCall = append(timingInfo.PerDrawCall, service.AtomRangeTimer{ToAtomIndex: uint64(id)})
			}
			mark = true
		}
		seen[id] = struct{}{}
		return mark, flush
	}

	ch := make(chan []*timerQuery, 1)
	go func() {
		if queries, ok := <-ch; ok {
			frames, draws := timingInfo.PerFrame, timingInfo.PerDrawCall
			for _, query := range queries {
				// Skip bad results (pending, incomplete). Note: disjoint results are kept.
				if query.state == queryPending || query.state == queryIncomplete {
					continue
				}

				// This logic makes the following assumptions:
				// • Queries are split at the frame and draw call edges.
				// • All frames and draws sit back-to-back, covering the entire atom list.
				// TODO: Refactor to find and use frame/draw overlaps with query intervals.

				// Skip frames and draw calls that don't have any overlapping query values.
				for ; len(frames) > 0 && frames[0].ToAtomIndex < uint64(query.endId); frames = frames[1:] {
				}
				for ; len(draws) > 0 && draws[0].ToAtomIndex < uint64(query.endId); draws = draws[1:] {
				}

				// Attribute delta time to the enclosing frame, assuming no gaps.
				if len(frames) > 0 {
					frames[0].Nanoseconds += query.value
				}
				// Attribute delta time to the enclosing draw call, assuming no gaps.
				if len(draws) > 0 {
					draws[0].Nanoseconds += query.value
				}
			}
			out <- replay.CallTiming{TimingInfo: timingInfo}
		} else {
			out <- replay.CallTiming{Error: errors.New("GPU timer queries not supported.")}
		}
	}()

	return &timingInfoGpuTransform{
		out:       ch,
		device:    device,
		db:        db,
		logger:    logger,
		predicate: predicate,
	}
}

type timingInfoGpuTransform struct {
	device    *service.Device
	db        database.Database
	logger    log.Logger
	predicate func(atom.ID, atom.Atom) (bool, bool)
	out       chan<- []*timerQuery

	// capabilities
	hasFences       bool
	hasValidContext bool
	timerQueryType  gpuTimerQueryType

	// transient state
	disjoint          bool
	availableQueryIds []QueryId
	pendingQueries    []*timerQuery
	queries           []*timerQuery
}

type timerQuery struct {
	state   queryState
	queryId QueryId
	startId atom.ID
	endId   atom.ID
	value   uint64
}

type eventType uint8
type queryState uint8
type gpuTimerQueryType uint8
type disjointCheckType uint8

const (
	newTimerQueriesChunkSize = 32

	rangeStart eventType = iota
	rangeEnd
	intermediateMarker

	gpuTimerQueryNone gpuTimerQueryType = iota
	gpuTimerQueryCore
	gpuTimerQueryDisjointExt
	gpuTimerQueryExt

	queryIncomplete queryState = iota
	queryPending
	querySucceeded
	queryDisjoint

	resetOnly disjointCheckType = iota
	checkAndReset
)

var (
	syncId           = GLsync(0x0FFFFFFF0FFFFFFF)
	transientPointer = memory.Tmp
)

func (t *timingInfoGpuTransform) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	out.Write(id, a)

	// Skip atoms until we have a valid context for issuing timing queries.
	if !t.hasValidContext {
		// Wait for the first ContextInfo atom, indicating the existence of a valid context.
		if _, ok := a.(*ContextInfo); ok {
			t.hasValidContext = true
			t.detectTimerSupport()

			// Append a query for the first frame/drawcall (no previous EndOfFrame).
			t.appendQuery(0, rangeStart, out)
		}
		return
	}

	mark, flush := t.predicate(id, a)
	if flush {
		t.appendQuery(id, rangeEnd, out)
		t.appendQuery(id, rangeStart, out)
	} else if mark {
		t.appendQuery(id, intermediateMarker, out)
	}
}

func (t *timingInfoGpuTransform) Flush(out atom.Writer) {
	// No timer query mechanism supported.
	if t.timerQueryType == gpuTimerQueryNone {
		close(t.out)
		return
	}

	// Account for all pending timers into the last drawcall and frame counters.
	t.appendQuery(atom.NoID, rangeEnd, out)

	// Synchronize on a final empty post before aggregating and returning query results.
	out.Write(atom.NoID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		b.Post(value.RemappedPointer(transientPointer.Address), 0, func(d binary.Decoder, err error) error {
			t.out <- t.queries
			close(t.out)
			return err
		})
		return nil
	}))
}

func (t *timingInfoGpuTransform) generateNewQueries(count int, out atom.Writer) {
	offset := len(t.availableQueryIds) + len(t.pendingQueries)
	newQueries := make([]QueryId, count)
	for i := 0; i < count; i++ {
		newQueries[i] = QueryId(0xF0FF0000 + offset + i)
	}
	out.Write(atom.NoID,
		NewGlGenQueries(GLsizei(count), transientPointer).
			AddRead(atom.Data(t.device.Architecture(), t.db, t.logger, transientPointer, newQueries)))
	t.availableQueryIds = append(t.availableQueryIds, newQueries...)
}

func (t *timingInfoGpuTransform) appendQuery(id atom.ID, et eventType, out atom.Writer) {
	if t.timerQueryType == gpuTimerQueryNone {
		return
	}

	beginQuery := func(id atom.ID) {
		if len(t.availableQueryIds) == 0 {
			t.generateNewQueries(newTimerQueriesChunkSize, out)
		}
		queryId := t.availableQueryIds[0]
		t.availableQueryIds = t.availableQueryIds[1:]

		query := &timerQuery{queryId: queryId, startId: id, state: queryIncomplete}
		t.queries = append(t.queries, query)
		t.pendingQueries = append(t.pendingQueries, query)
		out.Write(atom.NoID, NewGlBeginQuery(GLenum_GL_TIME_ELAPSED_EXT, queryId))
	}

	endQuery := func(id atom.ID) {
		query := t.pendingQueries[len(t.pendingQueries)-1]
		query.endId = id
		query.state = queryPending
		out.Write(atom.NoID, NewGlEndQuery(GLenum_GL_TIME_ELAPSED_EXT))
	}

	switch et {
	case rangeStart:
		t.checkDisjoint(resetOnly, out)
		beginQuery(id)

	case intermediateMarker:
		endQuery(id)
		beginQuery(id)

	case rangeEnd:
		endQuery(id)
		t.retrievePendingQueries(out)
	}
}

// checkDisjoint retrieves the GPU disjoint timer flag, optionally saving its status to t.
func (t *timingInfoGpuTransform) checkDisjoint(checkType disjointCheckType, out atom.Writer) {
	out.Write(atom.NoID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		NewGlGetIntegerv(GLenum_GL_GPU_DISJOINT_EXT, transientPointer).Replay(i, s, d, l, b)
		b.Post(value.RemappedPointer(transientPointer.Address), 4, func(d binary.Decoder, err error) error {
			if err != nil {
				return err
			}
			var value int32
			if value, err = d.Int32(); err == nil {
				switch checkType {
				case resetOnly:
					t.disjoint = false
				case checkAndReset:
					t.disjoint = value != 0
				}
			}
			return err
		})
		return nil
	}))
}

func (t *timingInfoGpuTransform) retrievePendingQueries(out atom.Writer) {
	// Wait for all commands to complete and check the disjoint timer flag before retrieving results.
	if t.timerQueryType == gpuTimerQueryDisjointExt {
		if t.hasFences {
			out.Write(atom.NoID, NewGlFenceSync(GLenum_GL_SYNC_GPU_COMMANDS_COMPLETE, 0, syncId))
			out.Write(atom.NoID, NewGlClientWaitSync(syncId, GLbitfield_GL_SYNC_FLUSH_COMMANDS_BIT, 0xFFFFFFFFFFFFFFFF, 0))
			out.Write(atom.NoID, NewGlDeleteSync(syncId))
		} else {
			out.Write(atom.NoID, NewGlReadPixels(0, 0, 1, 1, GLenum_GL_RGBA, GLenum_GL_UNSIGNED_BYTE, transientPointer))
		}
		t.checkDisjoint(checkAndReset, out)
	}

	// Retrieve result availability and value from all pending queries.
	for _, query := range t.pendingQueries {
		query := query

		switch t.timerQueryType {
		case gpuTimerQueryCore, gpuTimerQueryExt:
			// GL results retrieval is synchronous, mark as always succesful.
			query.state = querySucceeded

		case gpuTimerQueryDisjointExt:
			// GLES results retrieval is asynchronous, depends on result availability.
			out.Write(atom.NoID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
				NewGlGetQueryObjectivEXT(query.queryId, GLenum_GL_QUERY_RESULT_AVAILABLE, transientPointer).Replay(i, s, d, l, b)
				b.Post(value.RemappedPointer(transientPointer.Address), 4, func(d binary.Decoder, err error) error {
					if err != nil {
						return err
					}
					var value int32
					if value, err = d.Int32(); err == nil && value != 0 {
						if t.disjoint {
							query.state = queryDisjoint
						} else {
							query.state = querySucceeded
						}
					}
					return err
				})
				return nil
			}))
		}

		out.Write(atom.NoID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
			switch t.timerQueryType {
			case gpuTimerQueryCore:
				NewGlGetQueryObjectui64v(query.queryId, GLenum_GL_QUERY_RESULT, transientPointer).Replay(i, s, d, l, b)
			case gpuTimerQueryExt, gpuTimerQueryDisjointExt:
				NewGlGetQueryObjectui64vEXT(query.queryId, GLenum_GL_QUERY_RESULT, transientPointer).Replay(i, s, d, l, b)
			default:
				return nil
			}
			b.Post(value.RemappedPointer(transientPointer.Address), 8, func(d binary.Decoder, err error) error {
				if err != nil {
					return err
				}
				var value uint64
				if value, err = d.Uint64(); err == nil {
					query.value = value
				}
				return err
			})
			return nil
		}))

		// Recycle the completed query IDs.
		t.availableQueryIds = append(t.availableQueryIds, query.queryId)
	}

	// Clear pending queries so query objects can be reused.
	t.pendingQueries = t.pendingQueries[:0]
}

func (t *timingInfoGpuTransform) detectTimerSupport() {
	t.timerQueryType = gpuTimerQueryNone

	if v, err := ParseVersion(t.device.Version); err == nil {
		if v.IsES {
			if t.device.HasExtension("GL_EXT_disjoint_timer_query") {
				t.timerQueryType = gpuTimerQueryDisjointExt
			}
			if v.Major >= 3 { // OpenGL ES 3.0+
				t.hasFences = true
			}
		} else {
			if (v.Major == 3 && v.Minor >= 3) || v.Major > 3 { // OpenGL 3.3+
				t.timerQueryType = gpuTimerQueryCore
			} else if t.device.HasExtension("GL_EXT_timer_query") {
				t.timerQueryType = gpuTimerQueryExt
			}
			if (v.Major == 3 && v.Minor >= 2) || v.Major > 3 { // OpenGL 3.2+
				t.hasFences = true
			}
		}
	}
}
