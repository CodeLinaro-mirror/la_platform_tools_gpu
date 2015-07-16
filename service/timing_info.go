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

package service

import (
	"sort"

	"android.googlesource.com/platform/tools/gpu/atom"
)

// RangeDuration looks up the timing-record for the range of atoms, r.
// If the TimingInfo holds a timing-record for r then the duration in
// nanoseconds and true is returned.
// If no timing-record for r is contained in the TimingInfo then 0 and false
// is returned.
func (t TimingInfo) RangeDuration(r atom.Range) (uint64, bool) {
	s, e := uint64(r.Start), uint64(r.End)
	for _, group := range [][]AtomRangeTimer{t.PerDrawCall, t.PerFrame} {
		i := sort.Search(len(group), func(i int) bool {
			return group[i].FromAtomID >= s
		})
		if i < len(group) && group[i].FromAtomID == s && group[i].ToAtomID == e-1 {
			return group[i].Nanoseconds, true
		}
	}
	return 0, false
}

// Duration looks up the timing-record for the range of atoms, r.
// If the TimingInfo holds a timing-record for r then the duration in
// nanoseconds and true is returned.
// If no timing-record for r is contained in the TimingInfo then 0 and false
// is returned.
func (t TimingInfo) AtomDuration(id atom.ID) (uint64, bool) {
	a := uint64(id)
	i := sort.Search(len(t.PerCommand), func(i int) bool {
		return t.PerCommand[i].AtomID >= a
	})
	if i < len(t.PerCommand) && t.PerCommand[i].AtomID == a {
		return t.PerCommand[i].Nanoseconds, true
	}
	return 0, false
}
