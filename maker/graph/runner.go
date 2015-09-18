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

package graph

// Runner is the interface to something that controls the updating of a maker graph.
type Runner interface {
	// UpdateInputs is responsible for bringing all the inputs of a step.
	// It should return an error if any of the inputs failed, and should not return until either all the inputs are
	// ready or an error has occurred.
	UpdateInputs(s *Step) error
	// IsOutOfDate is called to check whether a step needs to be run.
	// The Runner is allowed to use any means to make this decision, but it is expected that it will involve looking at
	// the inputs to the outputs to determine if they have been changed in a way that invalidates the outputs.
	IsOutOfDate(s *Step) bool
	// Aborting should return true if the runner wishes to cease all processing of the graph. Normally this is because
	// an error has occurred and the runner wishes to terminate early rather than process as much as possible.
	Aborting() bool
}
