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
	"regexp"
	"time"

	"android.googlesource.com/platform/tools/gpu/maker"
)

const (
	// Default is the name of the entity that is built if none are supplied on the
	// command line.
	Default = "default"
)

// Entity represents an object in the build graph.
type Entity interface {
	// Name returns the unique name of this entity.
	// If Name is the empty string, the entity does not appear in the main entity
	// list and cannot be looked up.
	Name() string
	// Timestamp returns the last modified time of the entity, or the zero time if
	// not available.
	Timestamp() time.Time
	// NeedsUpdate returns true if the entity requires it's generating step to run.
	// The supplied timestamp is the newest timestamp of the entities this one
	// depends on, and can be used to decide if an update is neccesary.
	NeedsUpdate(t time.Time) bool
	// Updated is called when a step that modifies this entity completes.
	Updated()
}

var (
	Entities    = map[string]Entity{}
	entityHooks = []func(e Entity){}
)

// EntityHook registers a function that is invoked when a new entity is added
// to the system.
func EntityHook(f func(e Entity)) {
	entityHooks = append(entityHooks, f)
}

// FindEntity tries to look up an entity by name, and returns nil if the entity
// cannot be found.
func FindEntity(name string) Entity {
	if name == "" {
		return nil
	}
	e, _ := Entities[name]
	return e
}

// FindPathEntity tries to look up an entity by name.
// If an exact match cannot be found, then the name is treated as a path, and
// the absolute path is looked up instead.
// If that also does not match, it returns nil.
func FindPathEntity(name string) Entity {
	if e := FindEntity(name); e != nil {
		return e
	}
	if abs, err := maker.OSPath(name); err == nil {
		if e := FindEntity(abs); e != nil {
			return e
		}
	}
	return nil
}

// FindEntities tries to find all entities who's names match the supplied
// string. The string is treated as a case insensitive regular expression.
func FindEntities(pattern string) []Entity {
	re := regexp.MustCompile("(?i)" + pattern)
	matches := []Entity{}
	for name, e := range Entities {
		if re.MatchString(name) {
			matches = append(matches, e)
		}
	}
	return matches
}

// EntityOf tries to find an entity for the supplied value.
// If the value is an entity, it will be returned directly. If the value is a
// string it will use FindPathEntity to look it up.
// It is an error for v to not match an existing entity.
func EntityOf(v interface{}) Entity {
	switch v := v.(type) {
	case string:
		e := FindPathEntity(v)
		if e == nil {
			log.Fatalf("no such entity %s", v)
		}
		return e
	case Entity:
		return v
	default:
		log.Fatalf("cannot get entity for %T", v)
		return nil
	}
}

// AddEntity adds a new entity into the system.
// The entity must not have the same name as an already exixting entity.
func AddEntity(e Entity) {
	name := e.Name()
	if name != "" {
		_, found := Entities[name]
		if found {
			log.Fatalf("Attempt to remap entity name %s", name)
		}
		Entities[name] = e
	}
	for _, h := range entityHooks {
		h(e)
	}
}
