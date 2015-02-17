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

package transform

import "android.googlesource.com/platform/tools/gpu/atom"

type atomAtomID struct {
	atom atom.Atom
	id   atom.Id
}

// SkipDrawCalls is an implementation of Transformer that skips all draw calls
// that have not been explicitly requested.
type SkipDrawCalls struct {
	requests atom.IdSet
	buffer   []atomAtomID
}

// Draw adds an exception to allow all draw calls up to and including the atom
// with identifier id for the frame holding the atom.
func (t *SkipDrawCalls) Draw(id atom.Id) {
	if t.requests == nil {
		t.requests = make(atom.IdSet)
	}
	t.requests.Add(id)
}

func (t *SkipDrawCalls) Transform(id atom.Id, a atom.Atom, out atom.Writer) {
	t.buffer = append(t.buffer, atomAtomID{a, id})

	if t.requests.Contains(id) {
		t.flush(true, out)
		t.requests.Remove(id)
	} else if _, eos := a.(*atom.EOS); a.Flags().IsEndOfFrame() || eos {
		t.flush(false, out)
	}
}

func (t *SkipDrawCalls) flush(draw bool, out atom.Writer) {
	for _, a := range t.buffer {
		if draw || !a.atom.Flags().IsDrawCall() {
			out.Write(a.id, a.atom)
		}
	}
	t.buffer = t.buffer[:0]
}
