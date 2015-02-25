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

// Package semantic holds the set of types used in the abstract semantic graph
// representation of the api language.
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
	Register(func() { wg.Done() }, callbackTimeout)
	Register(func() { wg.Done() }, callbackTimeout)
	Register(func() { wg.Done() }, callbackTimeout)

	Exit()

	// Check that all callbacks have been called, or timeout after a second.
	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(callbackTimeout):
		t.Errorf("Expected all callbacks to have been called, but timed out.")
	}
}

func TestCallbackCalledOnInterruption(t *testing.T) {
	var wg sync.WaitGroup
	wg.Add(3)
	Register(func() { wg.Done() }, callbackTimeout)
	Register(func() { wg.Done() }, callbackTimeout)
	Register(func() { wg.Done() }, callbackTimeout)

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

	// Wait for signal interception, or timeout after a second.
	select {
	case <-interrupted:
	case <-time.After(interruptTimeout):
		t.Errorf("Unexpected signal interception timeout after one second.")
	}

	// Check that all callbacks have been called, or timeout after a second.
	done := make(chan bool, 1)
	go func() {
		wg.Wait()
		close(done)
	}()
	select {
	case <-done:
	case <-time.After(callbackTimeout):
		t.Errorf("Expected all callbacks to have been called, but timed out.")
	}
}
