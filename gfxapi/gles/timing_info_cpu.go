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
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/value"
	"android.googlesource.com/platform/tools/gpu/service"
)

func timingInfo(flags service.TimingFlags, out chan<- replay.CallTiming, device *service.Device, db database.Database, logger log.Logger) atom.Transformer {
	if flags&service.TimingGPU != 0 {
		return timingInfoGpu(flags, out, device, db, logger)
	}
	return &timingInfoCpuTransform{
		out:          out,
		perCommand:   flags&service.TimingPerCommand != 0,
		perDrawCall:  flags&service.TimingPerDrawCall != 0,
		perFrame:     flags&service.TimingPerFrame != 0,
		timerStartId: map[uint8]atom.ID{},
	}
}

const (
	commandThreadTimer uint8 = iota
	drawCallThreadTimer
	frameThreadTimer
)

// timingInfoCpuTransform is a transform used to postback timing information.
// Note: this is experimental and is likely to change in the near future.
type timingInfoCpuTransform struct {
	timingInfo   service.TimingInfo
	out          chan<- replay.CallTiming
	perCommand   bool
	perDrawCall  bool
	perFrame     bool
	timerStartId map[uint8]atom.ID
}

func (t *timingInfoCpuTransform) startTimer(fromId atom.ID, index uint8, out atom.Writer) {
	out.Write(atom.NoID, NewStartTimer(index))
	t.timerStartId[index] = fromId
}

func (t *timingInfoCpuTransform) stopTimer(toID atom.ID, index uint8, flags service.TimingFlags, out atom.Writer) {
	fromID := t.timerStartId[index]
	delete(t.timerStartId, index)

	out.Write(toID, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		NewStopTimer(index, 0).Replay(i, s, d, l, b) // returns a uint64 on the stack
		b.Store(value.VolatileTemporaryPointer(0))
		b.Post(value.VolatileTemporaryPointer(0), 8, func(d binary.Decoder, err error) error {
			var nanoseconds uint64
			if err == nil {
				nanoseconds = d.Uint64()
				err = d.Error()
			}
			if err != nil {
				t.out <- replay.CallTiming{Error: err}
				return err
			}

			switch flags {
			case service.TimingPerCommand:
				t.timingInfo.PerCommand = append(t.timingInfo.PerCommand, service.AtomTimer{
					AtomIndex:   uint64(toID),
					Nanoseconds: nanoseconds,
				})
			case service.TimingPerDrawCall:
				t.timingInfo.PerDrawCall = append(t.timingInfo.PerDrawCall, service.AtomRangeTimer{
					FromAtomIndex: uint64(fromID),
					ToAtomIndex:   uint64(toID),
					Nanoseconds:   nanoseconds,
				})
			case service.TimingPerFrame:
				t.timingInfo.PerFrame = append(t.timingInfo.PerFrame, service.AtomRangeTimer{
					FromAtomIndex: uint64(fromID),
					ToAtomIndex:   uint64(toID),
					Nanoseconds:   nanoseconds,
				})
			}
			return err
		})
		return nil
	}))

	switch flags {
	case service.TimingPerFrame:
		out.Write(toID, NewFlushPostBuffer())
	case service.TimingPerDrawCall:
		if !t.perFrame {
			out.Write(toID, NewFlushPostBuffer())
		}
	}
}

func (t *timingInfoCpuTransform) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	if _, frameStarted := t.timerStartId[frameThreadTimer]; t.perFrame && !frameStarted {
		t.startTimer(id, frameThreadTimer, out)
	}
	if _, drawCallStarted := t.timerStartId[drawCallThreadTimer]; t.perDrawCall && !drawCallStarted {
		t.startTimer(id, drawCallThreadTimer, out)
	}
	if t.perCommand {
		t.startTimer(id, commandThreadTimer, out)
	}

	out.Write(id, a)

	atomFlags := a.Flags()
	if t.perCommand {
		t.stopTimer(id, commandThreadTimer, service.TimingPerCommand, out)
	}
	if t.perDrawCall && (atomFlags.IsDrawCall() || atomFlags.IsEndOfFrame()) {
		t.stopTimer(id, drawCallThreadTimer, service.TimingPerDrawCall, out)
	}
	if t.perFrame && atomFlags.IsEndOfFrame() {
		t.stopTimer(id, frameThreadTimer, service.TimingPerFrame, out)
	}
}

func (t *timingInfoCpuTransform) Flush(out atom.Writer) {
	id := atom.NoID
	if _, drawCallStarted := t.timerStartId[drawCallThreadTimer]; drawCallStarted && t.perDrawCall {
		t.stopTimer(id, drawCallThreadTimer, service.TimingPerDrawCall, out)
	}
	if _, frameStarted := t.timerStartId[frameThreadTimer]; frameStarted && t.perFrame {
		t.stopTimer(id, frameThreadTimer, service.TimingPerFrame, out)
	}

	out.Write(id, replay.Custom(func(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error {
		b.Post(value.VolatileTemporaryPointer(0), 0, func(d binary.Decoder, err error) error {
			if err == nil {
				t.out <- replay.CallTiming{TimingInfo: t.timingInfo}
			} else {
				t.out <- replay.CallTiming{Error: err}
			}
			return err
		})
		return nil
	}))
}
