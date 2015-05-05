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
	"log"
	"path/filepath"
	"time"
)

// Entity represents an object in the build graph.
type Entity interface {
	// Name returns the unique name of this entity.
	// If Name is the empty string, the entity does not appear in the main entity
	// list and cannot be looked up.
	Name() string
	// Exists returns true if the entity currently exists.
	Exists() bool
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
	entities    = map[string]Entity{}
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
	e, _ := entities[name]
	return e
}

// EntityOf tries to find an entity for the supplied value.
// If the value is an entity, it will be returned directly. If the value is a
// string that matches an existing entity, that entity will be returned.
// If it is a relative path that when resolved mathes an existing file entity,
// then the file is returned.
func EntityOf(v interface{}) Entity {
	switch v := v.(type) {
	case string:
		if e := FindEntity(v); e != nil {
			return e
		}
		if abs, err := filepath.Abs(v); err == nil {
			if e := FindEntity(abs); e != nil {
				return e
			}
		}
		log.Fatalf("no such entity %s", v)
		return nil
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
		_, found := entities[name]
		if found {
			log.Fatalf("Attempt to remap entity name %s", name)
		}
		entities[name] = e
	}
	for _, h := range entityHooks {
		h(e)
	}
}
