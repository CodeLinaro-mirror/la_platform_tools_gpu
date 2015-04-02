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
	"bytes"
	"fmt"
	"sync"
	"time"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Transient database wrapper. The wrapper allows transient stores to
// be stored in RAM rather than disk.  A resource stored transiently
// must be loaded exactly once. After being loaded it is dropped from
// the store and another load will be an error. If the transient
// resource is never loaded it will eventually be purged.

// The transient database wrapper is used by the replay system where
// opcode streams are encoded as large byte streams which are normally
// used once with a few milliseconds of creation.

// How much to keep in RAM before trying to purge
const purgeSizeBytes = 100 << 10 // 100MB

// How many dropped timestamps to keep before trying to purge
const purgeDroppedCount = 100 // 100

// How many nanoseconds before it is safe to purge an item
const keepTimeNanoSeconds = int64(30) * int64(1000000000) // 30 seconds

func CreateTransientDatabase(inner Database) Database {
	return &transient{inner: inner, entries: make(map[binary.ID][]byte), droppedAt: make(map[binary.ID]time.Time), storedAt: make(map[binary.ID]time.Time)}
}

type transient struct {
	inner     Database                // database for non-transient resources
	mutex     sync.Mutex              // held when modifying the maps or size
	entries   map[binary.ID][]byte    // the transient resource data
	droppedAt map[binary.ID]time.Time // the time the transient resource was dropped from the store
	storedAt  map[binary.ID]time.Time // the time the transient resource was stored in the store
	size      int                     // the total size of all the transient resources
}

func (t *transient) StoreLink(to, id binary.ID, logger log.Logger) error {
	return fmt.Errorf("Can not StoreLink in a Transient Database")
}

func (t *transient) StoreRequest(request binary.Object, logger log.Logger) (id binary.ID, err error) {
	return binary.ID{}, fmt.Errorf("Can not StoreRequest in a Transient Database")
}

// Drop a resource from the store. The t.mutex is already held.
func (t *transient) dropLocked(id binary.ID, logger log.Logger) {
	if entry, ok := t.entries[id]; ok {
		t.size -= len(entry)
		delete(t.entries, id)
		// Keep track of the fact that this id has been dropped.
		now := time.Now()
		t.droppedAt[id] = now
	}

	delete(t.storedAt, id)
}

// Store a resource in the transient store. If the store gets large purge old info.
func (t *transient) Store(r binary.Object, logger log.Logger) (id binary.ID, err error) {
	logger = logger.Enter("Transient.Store")
	defer func() { logger.Info("↪ id: %v, err: %v total_size: %v", id, err, t.size) }()

	buf := &bytes.Buffer{}
	enc := cyclic.Encoder(vle.Writer(buf))
	if err := enc.Value(r); err != nil {
		return binary.ID{}, err
	}

	data := buf.Bytes()
	id = binary.NewID(data)

	t.mutex.Lock()
	now := time.Now()
	// Purge any old items from the transient store, if it is large.
	if t.size+len(data) > purgeSizeBytes {
		for id, storedTime := range t.storedAt {
			if now.Sub(storedTime).Nanoseconds() > keepTimeNanoSeconds {
				t.dropLocked(id, logger)
				if t.size <= purgeSizeBytes {
					break
				}
			}
		}
	}

	// Purge any old dropped markers from the store, if there are many.
	if len(t.droppedAt) > purgeDroppedCount {
		for id, droppedTime := range t.droppedAt {
			if now.Sub(droppedTime).Nanoseconds() > keepTimeNanoSeconds {
				delete(t.droppedAt, id)
				if len(t.droppedAt) <= purgeDroppedCount {
					break
				}
			}
		}
	}

	t.entries[id] = data
	t.storedAt[id] = now
	t.size += len(data)
	delete(t.droppedAt, id)

	t.mutex.Unlock()
	return id, nil
}

// Load a resource from the transient store. If not present in the transient store look in the
// underlying store. If in the transient store, drop it and return it.
func (t *transient) Load(id binary.ID, logger log.Logger, out binary.Object) (err error) {
	logger = logger.Enter("Transient.Load")
	logger.Info("(id: %v total_size: %v)", id, t.size)
	defer func() { logger.Info("↪ err: %v total_size: %v", err, t.size) }()

	t.mutex.Lock()
	if entry, ok := t.entries[id]; ok {
		t.mutex.Unlock()
		d := cyclic.Decoder(vle.Reader(bytes.NewBuffer(entry)))
		err := d.Value(out)
		t.mutex.Lock()
		t.dropLocked(id, logger)
		t.mutex.Unlock()
		return err
	} else {
		if t, ok := t.droppedAt[id]; ok {
			return fmt.Errorf("Id %v was dropped at %v", id, t)
		}
		t.mutex.Unlock()
		return t.inner.Load(id, logger, out)
	}
}

func (t *transient) Contains(id binary.ID, logger log.Logger) bool {
	_, found := t.entries[id]
	return found || t.inner.Contains(id, logger)
}

func (t *transient) Captures() (map[string]binary.ID, error) {
	return nil, nil
}

func (t *transient) Close() {
	t.inner.Close()
}
