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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

type stepError struct {
	step     *graph.Step
	err      error
	previous error
}

func (e *stepError) Error() string {
	if e == nil {
		return ""
	}
	if e.previous != nil {
		return e.previous.Error()
	}
	return e.String()
}

func (e *stepError) String() string {
	name := e.step.String()
	if name == "" {
		return fmt.Sprintf("%s", e.err)
	}
	return fmt.Sprintf("%s:%s", name, e.err)
}
