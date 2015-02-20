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

package store

import (
	"container/list"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
)

type lruEntry struct {
	id    binary.ID
	value interface{}
	size  int
}

type lruCache struct {
	mutex   sync.Mutex
	list    list.List
	entries map[binary.ID]*list.Element
	pending map[binary.ID]*sync.WaitGroup
	size    int
	maxSize int
}

func createLruCache(maxSize int) *lruCache {
	return &lruCache{
		entries: make(map[binary.ID]*list.Element),
		pending: make(map[binary.ID]*sync.WaitGroup),
		maxSize: maxSize,
	}
}

func (c *lruCache) add(e lruEntry) {
	if _, found := c.entries[e.id]; found {
		panic("Attempting to add same id twice to cache")
	}

	c.size += e.size
	for c.size > c.maxSize && c.list.Len() > 0 {
		c.remove(c.list.Back())
	}

	c.entries[e.id] = c.list.PushFront(e)
}

func (c *lruCache) remove(n *list.Element) {
	e := n.Value.(lruEntry)
	c.list.Remove(n)
	c.size -= e.size
	delete(c.entries, e.id)
}

func (c *lruCache) Store(id binary.ID, val interface{}, size int) {
	c.mutex.Lock()
	if e, found := c.entries[id]; found {
		c.remove(e)
	}
	c.add(lruEntry{id: id, value: val, size: size})
	c.mutex.Unlock()
}

type makerFunc func() (interface{}, int, error)

func (c *lruCache) Load(id binary.ID, maker makerFunc) (interface{}, int, error) {
	c.mutex.Lock()
	// Wait for an existing load, if it exists.
	if wg, found := c.pending[id]; found {
		c.mutex.Unlock()
		wg.Wait()
		c.mutex.Lock()
	}
	if n, found := c.entries[id]; found {
		// In cache. Move the entry to the front.
		c.list.MoveToFront(n)
		c.mutex.Unlock()
		e := n.Value.(lruEntry)
		return e.value, e.size, nil
	} else {
		// Not in cache, requires load.
		// Start by creating a pending wait-group
		wg := &sync.WaitGroup{}
		wg.Add(1)
		c.pending[id] = wg
		c.mutex.Unlock()

		// Perform the load
		e := lruEntry{id: id}
		var err error
		e.value, e.size, err = maker()

		// Add the entry to the cache, mark done and remove the wait-group from pending
		c.mutex.Lock()
		if err == nil {
			c.add(e)
		}
		wg.Done()
		delete(c.pending, id)
		c.mutex.Unlock()
		return e.value, e.size, err
	}
}

func (c *lruCache) Contains(id binary.ID) bool {
	c.mutex.Lock()
	_, found := c.entries[id]
	c.mutex.Unlock()
	return found
}
