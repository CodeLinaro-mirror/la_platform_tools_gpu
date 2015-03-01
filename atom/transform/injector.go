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

// Injector is an implementation of Transformer that can inject atoms into the
// atom stream.
type Injector struct {
	injections map[atom.ID][]atomAtomID
}

// Inject emits the atom a with identifier id after the atom with identifier
// after.
func (t *Injector) Inject(after atom.ID, id atom.ID, a atom.Atom) {
	if t.injections == nil {
		t.injections = make(map[atom.ID][]atomAtomID)
	}
	t.injections[after] = append(t.injections[after], atomAtomID{a, id})
}

func (t *Injector) Transform(id atom.ID, a atom.Atom, out atom.Writer) {
	out.Write(id, a)

	if r, ok := t.injections[id]; ok {
		for _, i := range r {
			out.Write(i.id, i.atom)
		}
		delete(t.injections, id)
	}
}
