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

// EarlyTerminator is an implementation of Transformer that will consume all
// atoms (except for the EOS atom) once all the atoms passed to Add have passed
// through the transformer.
type EarlyTerminator struct {
	requests atom.IDSet
}

// Add adds the atom with identifier id to the set of atoms that must be seen
// before the EarlyTerminator will consume all atoms (excluding the EOS atom).
func (t *EarlyTerminator) Add(id atom.ID) {
	if t.requests == nil {
		t.requests = make(atom.IDSet)
	}
	t.requests.Add(id)
}

func (t *EarlyTerminator) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	if len(t.requests) == 0 {
		// Seen all the atoms we want, ignore remaining ones
		return
	}
	out.Write(id, a)
	t.requests.Remove(id)
}

func (t *EarlyTerminator) Flush(out atom.Writer) {}
