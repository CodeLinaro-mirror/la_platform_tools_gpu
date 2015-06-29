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

package log

import "sync"

// Splitter is an implementation of the Logger interface that delegates all method calls to
// all logs passed to Add.
type Splitter struct {
	mutex     sync.Mutex
	listeners []Logger
}

// Add adds l to the list of loggers that the Splitter will delegate calls to.
func (s *Splitter) Add(l Logger) {
	s.mutex.Lock()
	s.listeners = append(s.listeners, l)
	s.mutex.Unlock()
}

// Logf will call Logf with the same arguments on all logs passed to Add.
func (s *Splitter) Log(severity Severity, msg string, args ...interface{}) {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, l := range s.listeners {
		l.Log(severity, msg, args...)
	}
}

// Enter will call Enter with the same argument on all logs passed to Add.
func (s *Splitter) Enter(name string) Logger {
	n := Splitter{}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, l := range s.listeners {
		n.listeners = append(n.listeners, l.Enter(name))
	}
	return &n
}

// Fork will call Fork on all logs passed to Add.
func (s *Splitter) Fork() Logger {
	n := Splitter{}
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, l := range s.listeners {
		n.listeners = append(n.listeners, l.Fork())
	}
	return &n
}

// Flush will call Flush on all logs passed to Add.
func (s *Splitter) Flush() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, l := range s.listeners {
		l.Flush()
	}
}

// Close will call Close on all logs passed to Add.
func (s *Splitter) Close() {
	s.mutex.Lock()
	defer s.mutex.Unlock()
	for _, l := range s.listeners {
		l.Close()
	}
}
