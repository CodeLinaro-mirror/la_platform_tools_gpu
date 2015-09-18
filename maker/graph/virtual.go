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

import (
	"log"
	"time"
)

// Virtual returns an Entity that represents a virtual object in the build graph.
// If name already represents a virtual entity, it will be returned, otherwise a
// new one will be built.
// This is used as the output for commands that have no file outputs.
func Virtual(name string) *virtual {
	e := FindEntity(name)
	if e != nil {
		l, is := e.(*virtual)
		if !is {
			log.Fatalf("%s is not a virtual entity", name)
		}
		return l
	}
	l := &virtual{name: name}
	AddEntity(l)
	return l
}

// List returns the Step for a virtual entity.
// If the virtual entity does not exist, it will be created, and if it does not
// have a Creator, then a new Step with no action will be created.
// This is used to build lists of entities that want to be updated together.
func List(name string) *Step {
	l := Virtual(name)
	s := Creator(l)
	if s == nil {
		s = NewStep(nil).Creates(l)
	}
	return s
}

// IsVirtual returns true if the supplied entity is a virtual type.
func IsVirtual(e Entity) bool {
	_, is := e.(*virtual)
	return is
}

type virtual struct {
	name string
}

// Name returns the name that was supplied when the virtual entity was built.
func (l *virtual) Name() string { return l.name }

// String returns the Name of the virtual for debugging.
func (l *virtual) String() string { return l.name }

// Timestamp return the zero time, virtual entitis are not modified.
func (l *virtual) Timestamp() time.Time { return time.Time{} }

// NeedsUpdate returns true, actions that build virtual objects should always be run.
func (l *virtual) NeedsUpdate(t time.Time) bool { return true }

// Updated does nothing.
func (l *virtual) Updated() {}
