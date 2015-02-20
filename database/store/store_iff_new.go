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

type storeIfNew struct {
	inner Store
}

// CreateStoreIfNew wraps a store with a helper that discards attempts to store
// a resource with a duplicate id to one already in the store.
func CreateStoreIfNew(inner Store) Store {
	return &storeIfNew{
		inner: inner,
	}
}

func (s storeIfNew) Store(id binary.ID, r binary.Object, data []byte, l log.Logger) error {
	if !s.inner.Contains(id) {
		return s.inner.Store(id, r, data, l)
	}
	return nil
}

func (s storeIfNew) Load(id binary.ID, l log.Logger, out binary.Object) (size int, err error) {
	return s.inner.Load(id, l, out)
}

func (s storeIfNew) Contains(id binary.ID) bool {
	return s.inner.Contains(id)
}

func (s storeIfNew) Close() {
	s.inner.Close()
}
