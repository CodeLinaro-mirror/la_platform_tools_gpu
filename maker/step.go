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
	"bufio"
	"fmt"
	"log"
	"os"
	"reflect"
	"sync"
	"time"
)

// Step represents an action that creates it's outputs from it's inputs.
// It is the linking entity in the build graph.
// Each entity can be built by only one step, but a step may build many entites,
// and may depend on many entities.
// It is an error to have a cycle in the build graph.
type Step struct {
	inputs  []Entity
	outputs []Entity
	always  bool
	action  func(*Step) error
	once    sync.Once
	done    chan struct{}
	err     error
}

var (
	steps     = map[Entity]*Step{}
	stepHooks = []func(s *Step){}
)

// StepHook adds a function that is invoked each time a step is added or
// modified.
func StepHook(f func(s *Step)) {
	stepHooks = append(stepHooks, f)
}

// Creator looks up the step that builds an entity.
// The entity will be looked up using EntityOf.
func Creator(of interface{}) *Step {
	e := EntityOf(of)
	s, _ := steps[e]
	return s
}

// NewStep creates and returns a new step that runs the supplied action.
func NewStep(a func(*Step) error) *Step {
	s := &Step{
		action: a,
		done:   make(chan struct{}),
	}
	for _, h := range stepHooks {
		h(s)
	}
	return s
}

// Creates adds new outputs to the Step.
// It will complain if any of the outputs already have a Creator.
func (s *Step) Creates(out ...Entity) *Step {
	for _, e := range out {
		if Creator(e) != nil {
			log.Fatalf("%s created by more than one step", e.Name())
		}
		steps[e] = s
	}
	s.outputs = append(s.outputs, out...)
	for _, h := range stepHooks {
		h(s)
	}
	return s
}

// DependsOn adds a new input dependancy to a Step.
// If the Step already depends on an input, it will not be added again.
func (s *Step) DependsOn(in ...interface{}) *Step {
	for _, v := range in {
		e := EntityOf(v)
		if !s.HasInput(e) {
			s.inputs = append(s.inputs, e)
		}
	}
	for _, h := range stepHooks {
		h(s)
	}
	return s
}

// HasInput returns true if the step already has the supplied entity in it's inputs.
func (s *Step) HasInput(e Entity) bool {
	for _, in := range s.inputs {
		if in == e {
			return true
		}
	}
	return false
}

// HasOutput returns true if the step already has the supplied entity in it's outputs.
func (s *Step) HasOutput(e Entity) bool {
	for _, out := range s.outputs {
		if out == e {
			return true
		}
	}
	return false
}

// AlwaysRun marks a step as always needing to run, rather than running only
// when it's outputs need updating.
func (s *Step) AlwaysRun() *Step {
	s.always = true
	return s
}

// String returns the name of the first output if present, for debugging.
func (s *Step) String() string {
	if len(s.outputs) > 0 {
		return s.outputs[0].Name()
	}
	return ""
}

// UseDepsFile reads in a deps file and uses it to populuate the inputs and
// outputs of the Step.
func (s *Step) UseDepsFile(deps Entity) {
	s.Creates(deps)
	depsfile, err := os.Open(deps.Name())
	if err != nil {
		return
	}
	defer depsfile.Close()
	scanner := bufio.NewScanner(depsfile)
	inputs := true
	for scanner.Scan() {
		line := scanner.Text()
		switch {
		case line == "==Inputs==":
			inputs = true
		case line == "==Outputs==":
			inputs = false
		case inputs:
			s.DependsOn(File(line))
		default:
			s.Creates(File(line))
		}
	}
}

// DependsStruct adds all the fields of the supplied struct as input dependacies
// of the step.
// It is an error if the struct has any fields that are not public fields of
// types that implement Entity.
func (s *Step) DependsStruct(v interface{}) {
	val := reflect.ValueOf(v)
	for i := 0; i < val.NumField(); i++ {
		f := val.Field(i)
		s.DependsOn(f.Interface().(Entity))
	}
}

func (s *Step) updateInputs() {
	deps := make([]*Step, 0, len(s.inputs))
	// Bring all inputs up to date in parallel
	for _, e := range s.inputs {
		dep := Creator(e)
		if dep != nil {
			deps = append(deps, dep)
			if Config.DisableParallel {
				dep.start()
			} else {
				go dep.start()
			}
		}
	}
	// Wait for all inputs to be ready
	for _, dep := range deps {
		<-dep.done
		if dep.err != nil && s.err == nil {
			s.err = fmt.Errorf("failed in %s", dep)
		}
	}
}

func (s *Step) shouldRun() bool {
	if s.always {
		return true
	}
	// Find the newest input
	t := time.Time{}
	for _, e := range s.inputs {
		t = Newest(t, e.Timestamp())
	}
	if t.IsZero() {
		// No timestamped inputs, so always run
		return true
	}
	// Ask the outputs if they want an update
	for _, e := range s.outputs {
		if e.NeedsUpdate(t) {
			return true
		}
	}
	return false
}

func (s *Step) run() {
	if s.action == nil {
		return
	}
	if err := s.action(s); err != nil {
		Errors.Add(s, err)
	}
	// Mark all our ouptuts as potentially updated
	for _, e := range s.outputs {
		e.Updated()
	}
}

func (s *Step) start() {
	s.once.Do(func() {
		s.updateInputs()
		if s.err == nil && s.shouldRun() {
			s.run()
		}
		// Signal we are complete
		close(s.done)
	})
}
