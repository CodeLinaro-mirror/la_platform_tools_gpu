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

	"android.googlesource.com/platform/tools/gpu/maker/config"
	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

type blockingRunner struct {
	useTimestamp
	abort bool
}

func (r *blockingRunner) UpdateInputs(s *graph.Step) error {
	var result error
	for _, e := range s.Inputs {
		dep := graph.Creator(e)
		if dep != nil {
			dep.Process(r)
			err := dep.Wait()
			if err != nil {
				se := &stepError{step: s, err: err, previous: result}
				fmt.Printf("Error: %s\n", se.String())
				result = se
				if config.StopOnError {
					r.abort = true
					return result
				}
			}
		}
	}
	return result
}

func (r *blockingRunner) Aborting() bool {
	return r.abort
}
