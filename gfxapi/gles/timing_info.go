package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
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
	out          chan<- gfxapi.CallTiming
	postback     replay.Postback
	perCommand   bool
	perDrawCall  bool
	perFrame     bool
	timerStartId map[uint8]atom.ID
}

func (t *timingInfoTransform) startTimer(cid atom.ContextID, fromId atom.ID, index uint8, out atom.Writer) {
	out.Write(transientID, NewStartTimer(cid, index))
	t.timerStartId[index] = fromId
}

func (t *timingInfoTransform) stopTimer(cid atom.ContextID, toID atom.ID, index uint8, mask service.TimingMask, out atom.Writer) {
	fromID := t.timerStartId[index]
	stopTimerId := t.postback(func(data interface{}, err error) {
		if err != nil {
			t.out <- gfxapi.CallTiming{Error: err}
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
	out.Write(stopTimerId, NewStopTimer(cid, index, 0))
	delete(t.timerStartId, index)

	switch mask {
	case service.TimingMaskTimingPerFrame:
		out.Write(transientID, NewFlushPostBuffer(cid))
	case service.TimingMaskTimingPerDrawCall:
		if !t.perFrame {
			out.Write(transientID, NewFlushPostBuffer(cid))
		}
	}
}

func (t *timingInfoTransform) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	cid := a.ContextID()
	switch a := a.(type) {
	case *Init:
		out.Write(id, a)

	case *atom.EOS:
		if _, drawCallStarted := t.timerStartId[drawCallThreadTimer]; drawCallStarted && t.perDrawCall {
			t.stopTimer(cid, id, drawCallThreadTimer, service.TimingMaskTimingPerDrawCall, out)
		}
		if _, frameStarted := t.timerStartId[frameThreadTimer]; frameStarted && t.perFrame {
			t.stopTimer(cid, id, frameThreadTimer, service.TimingMaskTimingPerFrame, out)
		}

		id := t.postback(func(interface{}, error) {
			t.out <- gfxapi.CallTiming{TimingInfo: t.timingInfo}
			close(t.out)
		})

		out.Write(id, a)

	default:
		if _, frameStarted := t.timerStartId[frameThreadTimer]; t.perFrame && !frameStarted {
			t.startTimer(cid, id, frameThreadTimer, out)
		}
		if _, drawCallStarted := t.timerStartId[drawCallThreadTimer]; t.perDrawCall && !drawCallStarted {
			t.startTimer(cid, id, drawCallThreadTimer, out)
		}
		if t.perCommand {
			t.startTimer(cid, id, commandThreadTimer, out)
		}

		out.Write(id, a)

		flags := a.Flags()
		if t.perCommand {
			t.stopTimer(cid, id, commandThreadTimer, service.TimingMaskTimingPerCommand, out)
		}
		if t.perDrawCall && (flags.IsDrawCall() || flags.IsEndOfFrame()) {
			t.stopTimer(cid, id, drawCallThreadTimer, service.TimingMaskTimingPerDrawCall, out)
		}
		if t.perFrame && (flags.IsEndOfFrame()) {
			t.stopTimer(cid, id, frameThreadTimer, service.TimingMaskTimingPerFrame, out)
		}
	}
}
