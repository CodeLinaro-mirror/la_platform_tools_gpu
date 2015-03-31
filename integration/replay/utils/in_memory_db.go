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

package utils

import (
	"bytes"
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
)

type inMemoryDatabase struct {
	links   map[binary.ID]binary.ID
	entries map[binary.ID]binary.Object
}

// NewInMemoryDatabase returns a partial implementation of Database, keeping
// all entries in-memory.
func NewInMemoryDatabase() database.Database {
	return &inMemoryDatabase{
		links:   make(map[binary.ID]binary.ID),
		entries: make(map[binary.ID]binary.Object),
	}
}

func (d *inMemoryDatabase) StoreLink(to, id binary.ID, _ log.Logger) error {
	d.links[id] = to
	return nil
}

func (d *inMemoryDatabase) StoreRequest(request binary.Object, _ log.Logger) (binary.ID, error) {
	panic("StoreRequest() not supported by the inMemoryDatabase")
	return binary.ID{}, nil
}

func (d *inMemoryDatabase) Store(o binary.Object, _ log.Logger) (binary.ID, error) {
	id := hash(o)
	d.entries[id] = o
	return id, nil
}

func (d *inMemoryDatabase) Load(id binary.ID, l log.Logger, out binary.Object) error {
	if link, found := d.links[id]; found {
		return d.Load(link, l, out)
	}
	if entry, found := d.entries[id]; found {
		store.CopyResource(out, entry)
		return nil
	}
	return fmt.Errorf("Resource '%v' not found", id)
}

func (d *inMemoryDatabase) Contains(id binary.ID, _ log.Logger) bool {
	if _, found := d.links[id]; found {
		return true
	}
	if _, found := d.entries[id]; found {
		return true
	}
	return false
}

func (d *inMemoryDatabase) Captures() (map[string]binary.ID, error) {
	panic("Captures() not supported by the inMemoryDatabase")
}

func (d *inMemoryDatabase) Close() {}

func hash(o binary.Object) binary.ID {
	b := bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(&b))
	if err := e.Value(o); err != nil {
		panic(err)
	}
	return binary.NewID(b.Bytes())
}
