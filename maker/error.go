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
	"fmt"
	"log"
	"sync"
)

type errorEntry struct {
	step *Step
	err  error
}

type errors struct {
	mu   sync.Mutex
	list []errorEntry
}

var (
	// Errors is the main error list, it holds all errors that the system has
	// encountered.
	Errors errors
)

// Add a new error to the error list.
// The method is concurrent safe.
func (errs *errors) Add(s *Step, err error) {
	errs.mu.Lock()
	defer errs.mu.Unlock()
	entry := errorEntry{s, err}
	log.Printf("Error: %s", entry)
	errs.list = append(errs.list, entry)
	if s != nil && s.err == nil {
		s.err = err
	}
}

// Returns true if the system has a registered error, and is thus in a failure
// state.
func (errs *errors) Failed() bool {
	return len(errs.list) > 0
}

// Returns the first error that was regisetered.
// This method is concurrent safe.
func (errs *errors) First() errorEntry {
	errs.mu.Lock()
	defer errs.mu.Unlock()
	return errs.list[0]
}

func (e errorEntry) String() string {
	return fmt.Sprintf("%s:%s", e.step, e.err)
}
