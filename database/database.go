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

// Package database implements the persistence layer for the gpu debugger tools.
package database

import (
	"bytes"
	"fmt"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Database is the interface to a resource store.
type Database interface {
	StoreLink(to, id binary.ID, logger log.Logger) error
	StoreRequest(request binary.Object, logger log.Logger) (id binary.ID, err error)
	Store(binary.Object, log.Logger) (binary.ID, error)
	Load(binary.ID, log.Logger, binary.Object) error
	Contains(binary.ID, log.Logger) bool
	Close()
}

// Create builds a new database.
func Create(path string, maxDataCacheSize, metaDataCompactionSize, maxDerivedCacheSize int, builder builder) Database {
	return &database{
		records: map[binary.ID]*record{},
		builder: builder,
	}
}

func InMemory() Database {
	return &database{
		records: map[binary.ID]*record{},
	}
}

type record struct {
	value   binary.Object
	request binary.Object
	link    binary.ID
	err     error
	wait    chan struct{}
}

type database struct {
	mutex   sync.Mutex
	records map[binary.ID]*record
	builder builder
}

func (d *database) StoreLink(to, id binary.ID, logger log.Logger) error {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, got := d.records[id]
	if to != id && !got {
		d.records[id] = &record{link: to}
	}
	return nil
}

func (d *database) StoreRequest(o binary.Object, logger log.Logger) (binary.ID, error) {
	id := hash(o)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, got := d.records[id]
	if !got {
		d.records[id] = &record{request: o}
	}
	return id, nil
}

func (d *database) Store(o binary.Object, logger log.Logger) (binary.ID, error) {
	id := hash(o)
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, got := d.records[id]
	if !got {
		d.records[id] = &record{value: o}
	}
	return id, nil
}

func (d *database) Load(id binary.ID, logger log.Logger, out binary.Object) (err error) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	return d.load(id, logger, out)
}

// load function must be called with a locked mutex
func (d *database) load(id binary.ID, logger log.Logger, out binary.Object) (err error) {
	r, got := d.records[id]
	if !got {
		return fmt.Errorf("Resource '%v' not found", id)
	}
	if r.value != nil {
		// already have a value, copy it to the out and we are done
		store.CopyResource(out, r.value)
		return r.err
	}
	if r.request == nil {
		// not a request or value, must be a link, so load it
		return d.load(r.link, logger, out)
	}
	if r.wait != nil {
		// an in progress request, wait for it
		d.mutex.Unlock()     // unlock before waiting
		defer d.mutex.Lock() // relock after waiting
		<-r.wait
		store.CopyResource(out, r.value)
		return r.err
	}
	// must be a first time access to request
	r.wait = make(chan struct{})
	r.err = func() error { // func for defer scope
		d.mutex.Unlock()     // don't build under the lock
		defer d.mutex.Lock() // relock after build
		return d.builder.BuildResource(r.request, d, logger, out)
	}()
	r.value = out
	close(r.wait)
	return r.err
}

func (d *database) Contains(id binary.ID, logger log.Logger) (res bool) {
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, got := d.records[id]
	return got
}

func (d *database) Close() {}

func hash(o binary.Object) binary.ID {
	b := bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(&b))
	if err := e.Value(o); err != nil {
		panic(err)
	}
	return binary.NewID(b.Bytes())
}
