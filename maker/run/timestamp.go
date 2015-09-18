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

package run

import (
	"time"

	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

func newest(t1, t2 time.Time) time.Time {
	if !t1.IsZero() && (t2.IsZero() || t1.After(t2)) {
		return t1
	}
	return t2
}

type useTimestamp struct{}

func (useTimestamp) IsOutOfDate(s *graph.Step) bool {
	// Find the newest input
	t := time.Time{}
	for _, e := range s.Inputs {
		t = newest(t, e.Timestamp())
	}
	if t.IsZero() {
		// No timestamped inputs, so always run
		return true
	}
	// Ask the outputs if they want an update
	for _, e := range s.Outputs {
		if e.NeedsUpdate(t) {
			return true
		}
	}
	return false
}
