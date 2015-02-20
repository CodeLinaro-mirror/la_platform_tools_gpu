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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/log"
)

type cache struct {
	lru   *lruCache
	inner Store
}

// CreateCache builds a store that uses a memory bounded lru cache over the top
// of another store.
func CreateCache(maxSize int, inner Store) Store {
	return &cache{createLruCache(maxSize), inner}
}

func (c *cache) Store(id binary.ID, r binary.Object, data []byte, l log.Logger) error {
	c.lru.Store(id, r, len(data))
	return c.inner.Store(id, r, data, l)
}

func (c *cache) Load(id binary.ID, l log.Logger, out binary.Object) (int, error) {
	o, size, err := c.lru.Load(id, func() (interface{}, int, error) {
		s, e := c.inner.Load(id, l, out)
		return out, s, e
	})
	CopyResource(out, o)
	return size, err
}

func (c *cache) Contains(id binary.ID) bool {
	return c.lru.Contains(id) || c.inner.Contains(id)
}

func (c cache) Close() {
	c.inner.Close()
}
