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

package maker

import (
	"os/exec"
	"time"
)

// FindTool finds an executable on the host search path, and returns a File
// entity for the tool if found.
func FindTool(name string) Entity {
	path, err := exec.LookPath(name)
	if err != nil {
		return nil
	}
	return File(path)
}

// Oldest returns the oldest of two timestamps.
// Timestamps with a zero value do not count as older.
func Oldest(t1, t2 time.Time) time.Time {
	if !t1.IsZero() && (t2.IsZero() || t1.Before(t2)) {
		return t1
	}
	return t2
}

// Newest returns the newest of two timestamps.
// Timestamps with a zero value do not count as newer.
func Newest(t1, t2 time.Time) time.Time {
	if !t1.IsZero() && (t2.IsZero() || t1.After(t2)) {
		return t1
	}
	return t2
}
