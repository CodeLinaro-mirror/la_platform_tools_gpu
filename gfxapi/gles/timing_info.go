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
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/service"
)

const transientID = atom.ID(0xffffffffffffffff)

const (
	commandThreadTimer uint8 = iota
	drawCallThreadTimer
	frameThreadTimer
)

// timingInfoTransform is a transform used to postback timing information.
// Note: this is experimental and is likely to change in the near future.
type timingInfoTransform struct {
	timingInfo   service.TimingInfo
	out          chan<- replay.CallTiming
	postback     replay.Postback
	perCommand   bool
	perDrawCall  bool
	perFrame     bool
	timerStartId map[uint8]atom.ID
}

func (t *timingInfoTransform) startTimer(fromId atom.ID, index uint8, out atom.Writer) {
	out.Write(transientID, NewStartTimer(index))
	t.timerStartId[index] = fromId
}

func (t *timingInfoTransform) stopTimer(toID atom.ID, index uint8, mask service.TimingMask, out atom.Writer) {
	fromID := t.timerStartId[index]
	stopTimerId := t.postback(func(data interface{}, err error) {
		if err != nil {
			t.out <- replay.CallTiming{Error: err}
			return
		}
		val := data.(StopTimer_Postback)
		switch mask {
		case service.TimingMaskTimingPerCommand:
			t.timingInfo.PerCommand = append(t.timingInfo.PerCommand, service.AtomTimer{
				AtomId:      uint64(toID),
				Nanoseconds: val.Result,
			})
		case service.TimingMaskTimingPerDrawCall:
			t.timingInfo.PerDrawCall = append(t.timingInfo.PerDrawCall, service.AtomRangeTimer{
				FromAtomId:  uint64(fromID),
				ToAtomId:    uint64(toID),
				Nanoseconds: val.Result,
			})
		case service.TimingMaskTimingPerFrame:
			t.timingInfo.PerFrame = append(t.timingInfo.PerFrame, service.AtomRangeTimer{
				FromAtomId:  uint64(fromID),
				ToAtomId:    uint64(toID),
				Nanoseconds: val.Result,
			})
		}
	})
	out.Write(stopTimerId, NewStopTimer(index, 0))
	delete(t.timerStartId, index)

	switch mask {
	case service.TimingMaskTimingPerFrame:
		out.Write(transientID, NewFlushPostBuffer())
	case service.TimingMaskTimingPerDrawCall:
		if !t.perFrame {
			out.Write(transientID, NewFlushPostBuffer())
		}
	}
}

func (t *timingInfoTransform) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
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

	flags := a.Flags()
	if t.perCommand {
		t.stopTimer(id, commandThreadTimer, service.TimingMaskTimingPerCommand, out)
	}
	if t.perDrawCall && (flags.IsDrawCall() || flags.IsEndOfFrame()) {
		t.stopTimer(id, drawCallThreadTimer, service.TimingMaskTimingPerDrawCall, out)
	}
	if t.perFrame && (flags.IsEndOfFrame()) {
		t.stopTimer(id, frameThreadTimer, service.TimingMaskTimingPerFrame, out)
	}
}

func (t *timingInfoTransform) Flush(out atom.Writer) {
	id := atom.NoID
	if _, drawCallStarted := t.timerStartId[drawCallThreadTimer]; drawCallStarted && t.perDrawCall {
		t.stopTimer(id, drawCallThreadTimer, service.TimingMaskTimingPerDrawCall, out)
	}
	if _, frameStarted := t.timerStartId[frameThreadTimer]; frameStarted && t.perFrame {
		t.stopTimer(id, frameThreadTimer, service.TimingMaskTimingPerFrame, out)
	}

	t.postback(func(interface{}, error) {
		t.out <- replay.CallTiming{TimingInfo: t.timingInfo}
		close(t.out)
	})
}
