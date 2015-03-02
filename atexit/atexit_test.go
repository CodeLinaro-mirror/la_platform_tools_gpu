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

package atexit

import (
	"os"
	"os/signal"
	"sync"
	"testing"
	"time"
)

const (
	callbackTimeout  = time.Second
	interruptTimeout = time.Second
)

func TestCallbackCalledOnExit(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	Register(func() { wg.Done(); select {} }, callbackTimeout)
	Register(func() { wg.Done(); select {} }, callbackTimeout)
	Register(func() { wg.Done(); select {} }, callbackTimeout)

	go Exit(0)

	// Close the done channel after all callbacks have been called.
	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		close(done)
	}()

	// Note: we're only waiting for half the timeout time, otherwise os.Exit()
	// will kick in and our test won't have a chance to complete. This is
	// shorter than the worst-case expected behavior but our registered callbacks
	// aren't doing anything before closing their WaitGroup item.
	select {
	case <-done:
	case <-time.After(callbackTimeout / 2):
		t.Errorf("Unexpected callbacks timeout.")
	}
}

func TestCallbackCalledOnInterruption(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	Register(func() { wg.Done(); select {} }, callbackTimeout)
	Register(func() { wg.Done(); select {} }, callbackTimeout)
	Register(func() { wg.Done(); select {} }, callbackTimeout)

	// Setup interruption signal interception.
	interrupted := make(chan bool, 1)
	go func() {
		sigchan := make(chan os.Signal, 1)
		signal.Notify(sigchan, os.Interrupt)
		<-sigchan
		close(interrupted)
	}()

	// Send an interruption signal.
	proc, err := os.FindProcess(os.Getpid())
	if err != nil {
		t.Errorf("Unexpected error: %v.", err)
	}
	proc.Signal(os.Interrupt)

	// Wait for signal interception, or timeout after interruptTimeout.
	select {
	case <-interrupted:
	case <-time.After(interruptTimeout):
		t.Errorf("Unexpected signal timeout.")
	}

	// Close the done channel after all callbacks have been called.
	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		close(done)
	}()

	// Note: we're only waiting for half the timeout time, otherwise os.Exit()
	// will kick in and our test won't have a chance to complete. This is
	// shorter than the worst-case expected behavior but our registered callbacks
	// aren't doing anything before closing their WaitGroup item.
	select {
	case <-done:
	case <-time.After(callbackTimeout / 2):
		t.Errorf("Unexpected callbacks timeout.")
	}
}
