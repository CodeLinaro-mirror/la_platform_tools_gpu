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

// Package atexit enables registration of cleanup goroutines to be run at
// program exit or when the process receives an interruption or kill signal.
package atexit

import (
	"os"
	"os/signal"
	"sync"
	"time"
)

var (
	maxTimeout    time.Duration
	exitCallbacks []func()
	mutex         sync.Mutex     // Guards maxTimeout and exitCallbacks.
	sigchan       chan os.Signal = make(chan os.Signal)
)

func init() {
	go func() {
		// Wait for incoming signals to intercept, before calling Exit.
		<-sigchan
		Exit(0)
	}()
}

// Exit calls all registered callbacks once in separate goroutines, waiting for them to complete
// for a duration of at least the maximum timeout value given to Register, then calls os.Exit with
// the given return code.
func Exit(code int) {
	// Disable signal interception.
	signal.Stop(sigchan)

	// Swap exitCallbacks to prevent them from being called twice, and start the timeout counter.
	mutex.Lock()
	callbacks := exitCallbacks
	exitCallbacks = nil
	timeout := time.After(maxTimeout)
	mutex.Unlock()

	// Run all exit callbacks on separate goroutines and signal their completion by closing chan c.
	c := make(chan struct{})
	go func() {
		var wg sync.WaitGroup
		wg.Add(len(callbacks))
		for _, cb := range callbacks {
			callback := cb
			go func() {
				callback()
				wg.Done()
			}()
		}
		wg.Wait()
		close(c)
	}()

	// Return after all callbacks have completed or after the counter times out, whichever comes first.
	select {
	case <-c:
	case <-timeout:
	}

	// Explicitly terminate the current process.
	os.Exit(code)
}

// Register adds the given function f to a list of callbacks that will get called when the current
// process receives an interruption or kill signal, or when Exit is explicitly called. f will be
// given a duration of at least timeout to complete.
func Register(f func(), timeout time.Duration) {
	if f == nil {
		return
	}

	mutex.Lock()
	exitCallbacks = append(exitCallbacks, f)
	if timeout > maxTimeout {
		maxTimeout = timeout
	}
	mutex.Unlock()

	// Enable signal interception, no-op if already enabled.
	// Note: for Unix, these signals translate to SIGINT and SIGKILL.
	signal.Notify(sigchan, os.Interrupt, os.Kill)
}
