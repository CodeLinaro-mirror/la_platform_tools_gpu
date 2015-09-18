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
	"sync/atomic"

	"android.googlesource.com/platform/tools/gpu/maker/config"
	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

type parallelRunner struct {
	useTimestamp
	abort int32
}

func (r *parallelRunner) UpdateInputs(s *graph.Step) error {
	deps := make([]*graph.Step, 0, len(s.Inputs))
	// Bring all inputs up to date in parallel
	for _, e := range s.Inputs {
		dep := graph.Creator(e)
		if dep != nil {
			deps = append(deps, dep)
			go dep.Process(r)
		}
	}
	// Wait for all inputs to be ready
	var result error
	for _, dep := range deps {
		err := dep.Wait()
		if err != nil {
			se := &stepError{step: s, err: err, previous: result}
			fmt.Printf("Error: %s\n", se.String())
			result = se
			if config.StopOnError {
				atomic.StoreInt32(&r.abort, 1)
			}
		}
	}
	return result
}

func (r *parallelRunner) Aborting() bool {
	return atomic.LoadInt32(&r.abort) > 0
}
