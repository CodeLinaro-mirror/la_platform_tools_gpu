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

package client

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kCommandAdapterItemHeight = 18

type CommandAdapter struct {
	gxui.AdapterBase
	hierarchyTreeNode
}

func CreateCommandAdapter(appCtx *ApplicationContext) *CommandAdapter {
	a := &CommandAdapter{}
	a.ctx = &commandAdapterCtx{appCtx: appCtx}
	return a
}

func (a *CommandAdapter) UpdateAtoms(capture *path.Capture, atoms []atom.Atom, root atom.Group) {
	a.ctx.capture = capture
	a.ctx.atoms = atoms
	a.group = root
	a.DataChanged()
}

func (a *CommandAdapter) UpdateTimings(timings service.TimingInfo) {
	a.ctx.timings = timings
	a.DataChanged()
}

func (a *CommandAdapter) UpdateDevice(device *path.Device) {
	a.ctx.device = device
	a.DataChanged()
}

func (a CommandAdapter) Path(item gxui.AdapterItem) path.Path {
	switch i := item.(type) {
	case nil:
		return nil
	case hierarchyItem:
		return a.ctx.capture.Atoms().Slice(uint64(i.rng.Start), uint64(i.rng.End))
	case atomItem:
		return a.ctx.capture.Atoms().Index(uint64(i.atomID))
	case observationItem:
		return a.ctx.capture.Atoms().Index(uint64(i.atomID)).Field("Observations").ArrayIndex(uint64(i.index))
	default:
		panic(fmt.Errorf("Unknown item type %T", i))
	}
}

func (a CommandAdapter) Item(p path.Path) gxui.AdapterItem {
	if s, _ := path.FindAtomSlice(p); s != nil {
		return hierarchyItem{
			rng: atom.Range{Start: atom.ID(s.Start), End: atom.ID(s.End)},
		}
	}
	if a := path.FindAtom(p); a != nil {
		return atomItem{atomID: atom.ID(a.Index)}
	}
	return nil
}

// gxui.TreeAdapter compliance
func (a CommandAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: kCommandAdapterItemHeight}
}

type commandAdapterCtx struct {
	appCtx  *ApplicationContext
	capture *path.Capture
	device  *path.Device
	atoms   []atom.Atom
	timings service.TimingInfo
}
