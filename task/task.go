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

// Package task contains types that can be used to create cancelable tasks.
package task

import (
	"errors"
	"sync"
)

var errCanceled = errors.New("Runner canceled")

// hidden is a private type used to prevent external packages from constructing
// CancelSignals.
type hidden struct{}

// CancelSignal is used by Runner to check whether they have been canceled.
type CancelSignal <-chan hidden

// Check stops execution of the Runner if it has been canceled.
// If Check is called and the Runner has been canceled then the deferred
// functions of the Runner will be executed and then the Runner will immediately
// cease execution.
// If the Runner has not been canceled, then Check simply returns and the
// Runable continues to execute.
func (c CancelSignal) Check() {
	select {
	case <-c:
		panic(errCanceled)
	default:
	}
}

// Tasks are the executors of Runners.
// All methods on a Task support concurrent go-routine usage.
type Task struct {
	c chan hidden
	m sync.Mutex
}

// New constructs and returns a new Task.
func New() *Task {
	return &Task{}
}

// Runner is the interface implemented by types that implement the Run method.
type Runner interface {
	// Run is called by a Task to execute lengthy logic on a new go-routine.
	// The Run method should periodically check whether it should continue
	// execution by calling Check() on the CancelSignal.
	Run(CancelSignal)
}

// Run begins execution of the Runner r on a new go-routine.
// If a previous Runner was started by this Task, then it is automatically
// canceled. Note however, that Run does not wait for the previous Runner to
// return before r is started.
// The Runner can be canceled by calling Cancel on this Task.
func (t *Task) Run(r Runner) {
	t.m.Lock()
	defer t.m.Unlock()

	if t.c != nil {
		close(t.c)
	}
	t.c = make(chan hidden)
	go func() {
		defer func() {
			if err := recover(); err != nil {
				if err != errCanceled {
					panic(err)
				}
			}
		}()
		r.Run(t.c)
	}()
}

// Cancel signals to the last Runner started with Run() that it should stop.
// Repeated calls to Cancel() will do nothing.
func (t *Task) Cancel() {
	t.m.Lock()
	defer t.m.Unlock()

	if t.c != nil {
		close(t.c)
		t.c = nil
	}
}
