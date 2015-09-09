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

package database

import (
	"fmt"
	"reflect"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/config"
	"android.googlesource.com/platform/tools/gpu/log"
)

// NewInMemory builds a new in memory database.
func NewInMemory(buildContext interface{}) Database {
	return &memory{
		records:      map[binary.ID]*record{},
		buildContext: buildContext,
	}
}

type record struct {
	value interface{}
	err   error
	wait  chan struct{}
}

type memory struct {
	mutex        sync.Mutex
	records      map[binary.ID]*record
	buildContext interface{} // The build context, user-defined.
}

// Implements Database
func (d *memory) Store(id binary.ID, v interface{}, logger log.Logger) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.store(id, v, logger)
}

// store function must be called with a locked mutex
func (d *memory) store(id binary.ID, v interface{}, logger log.Logger) error {
	if v == nil {
		panic("Store nil in database (that is bad)")
	}
	r, got := d.records[id]
	if !got {
		d.records[id] = &record{value: v}
	} else if config.DebugDatabaseVerify {
		if !reflect.DeepEqual(v, r.value) {
			return fmt.Errorf("Duplicate object id %v", id)
		}
	}
	return nil
}

// Implements Database
func (d *memory) Resolve(id binary.ID, logger log.Logger) (interface{}, error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.resolve(id, logger)
}

// load function must be called with a locked mutex
func (d *memory) resolve(id binary.ID, logger log.Logger) (interface{}, error) {
	r, got := d.records[id]
	if !got {
		return nil, fmt.Errorf("Resource '%v' not found", id)
	}
	if r.err != nil {
		return r.value, r.err
	}
	lazy, islazy := r.value.(Lazy)
	if !islazy {
		return r.value, r.err
	}
	// transform the id to get the lazy result id
	lazyID := LazyOutputID(id)
	if r.wait != nil {
		// an in progress request, wait for it
		d.mutex.Unlock() // unlock before waiting
		<-r.wait
		d.mutex.Lock() // relock after waiting
		return d.resolve(lazyID, logger)
	}
	// must be a first time access to request
	r.wait = make(chan struct{})
	value, err := func() (interface{}, error) { // func for defer scope
		d.mutex.Unlock()     // don't build under the lock
		defer d.mutex.Lock() // relock after build
		return lazy.BuildLazy(d.buildContext, d, logger)
	}()
	if err == nil {
		err = d.store(lazyID, value, logger)
	}
	r.err = err
	close(r.wait)
	return value, r.err
}

// Implements Database
func (d *memory) Contains(id binary.ID, logger log.Logger) (res bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, got := d.records[id]
	return got
}
