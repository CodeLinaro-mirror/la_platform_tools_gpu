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
	"crypto/sha1"
	"fmt"
	"reflect"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Database is the interface to a resource store.
type Database interface {
	StoreLink(to, id binary.ID, logger log.Logger) error
	Store(binary.Object, log.Logger) (binary.ID, error)
	Load(binary.ID, log.Logger, binary.Object) error
	Contains(binary.ID, log.Logger) bool
	Close()
}

func StoreLink(d Database, to, id binary.ID, l log.Logger) error {
	return d.StoreLink(to, id, l)
}

func StoreRequest(d Database, obj binary.Object, l log.Logger) (binary.ID, error) {
	return d.Store(obj, l)
}

func Store(d Database, obj binary.Object, l log.Logger) (binary.ID, error) {
	return d.Store(obj, l)
}

func Load(d Database, id binary.ID, l log.Logger, out binary.Object) error {
	return d.Load(id, l, out)
}

// NewInMemory builds a new in memory database.
func NewInMemory(buildContext interface{}) Database {
	return &database{
		records:      map[binary.ID]*record{},
		buildContext: buildContext,
	}
}

type record struct {
	value binary.Object
	lazy  Lazy
	link  binary.ID
	err   error
	wait  chan struct{}
}

type database struct {
	mutex        sync.Mutex
	records      map[binary.ID]*record
	buildContext interface{} // The build context, user-defined.
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

func (d *database) Store(o binary.Object, logger log.Logger) (binary.ID, error) {
	id, err := Hash(o)
	if err != nil {
		return id, err
	}
	d.mutex.Lock()
	defer d.mutex.Unlock()
	_, got := d.records[id]
	if !got {
		if lazy, islazy := o.(Lazy); islazy {
			d.records[id] = &record{lazy: lazy}
		} else {
			d.records[id] = &record{value: o}
		}
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
		CopyResource(out, r.value)
		return r.err
	}
	if r.lazy == nil {
		// not a request or value, must be a link, so load it
		return d.load(r.link, logger, out)
	}
	if r.wait != nil {
		// an in progress request, wait for it
		d.mutex.Unlock()     // unlock before waiting
		defer d.mutex.Lock() // relock after waiting
		<-r.wait
		CopyResource(out, r.value)
		return r.err
	}
	// must be a first time access to request
	r.wait = make(chan struct{})
	r.err = func() error { // func for defer scope
		d.mutex.Unlock()     // don't build under the lock
		defer d.mutex.Lock() // relock after build

		built, err := r.lazy.BuildLazy(d.buildContext, d, logger)
		CopyResource(out, built)
		return err
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

// Hash returns a unique binary.ID based on the contents of the object.
// Two objects of identical content will return the same ID, and the
// probability of two objects with different content generating the same ID
// will be ignorable.
// Objects with a graph structure are allowed.
// Only members that would be encoded using a binary.Encoder are considered.
func Hash(o binary.Object) (binary.ID, error) {
	id := binary.ID{}
	h := sha1.New()
	e := cyclic.Encoder(vle.Writer(h))
	if err := e.Value(o); err != nil {
		return id, err
	}
	copy(id[:], h.Sum(nil))
	return id, nil
}

// CopyResource assigns the value object to the variable out points to
func CopyResource(out interface{}, value interface{}) {
	o := reflect.ValueOf(out).Elem()
	v := reflect.ValueOf(value).Elem()
	o.Set(v)
}
