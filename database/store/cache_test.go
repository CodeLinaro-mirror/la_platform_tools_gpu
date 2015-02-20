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
	"bytes"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/log"
)

type storeCallback func(id binary.ID, r binary.Object, d []byte) error
type loadCallback func(id binary.ID, out binary.Object) (size int, err error)
type containsCallback func(id binary.ID) bool
type closeCallback func()

type mockStore struct {
	onStore    storeCallback
	onLoad     loadCallback
	onContains containsCallback
	onClose    closeCallback
}

func (s mockStore) Store(id binary.ID, r binary.Object, d []byte, _ log.Logger) error {
	return s.onStore(id, r, d)
}
func (s mockStore) Load(id binary.ID, _ log.Logger, out binary.Object) (size int, err error) {
	return s.onLoad(id, out)
}
func (s mockStore) Contains(id binary.ID) bool {
	return s.onContains(id)
}
func (s mockStore) Close() {
	s.onClose()
}

func createMockStore() *mockStore {
	return &mockStore{
		onStore:    func(binary.ID, binary.Object, []byte) error { return nil },
		onLoad:     func(binary.ID, binary.Object) (size int, err error) { return -1, nil },
		onContains: func(binary.ID) bool { return false },
		onClose:    func() {},
	}
}

func validateLoad(t *testing.T, cache *cache, id binary.ID, expectedSize int, expectedData *binary.Data, expectedErr error) {
	data := &binary.Data{}
	size, err := cache.Load(id, nil, data)
	if expectedSize != size {
		t.Errorf("invalid size, expected %v got %v", expectedSize, size)
	}
	if !reflect.DeepEqual(expectedData, data) {
		t.Errorf("invalid data, expected %v got %v", expectedData, data)
	}
	if expectedErr != err {
		t.Errorf("invalid error, expected %v got %v", expectedErr, err)
	}
}

func encode(r binary.Object) []byte {
	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	if err := r.Encode(enc); err != nil {
		panic(err)
	}
	return buf.Bytes()
}

func calcSize(r binary.Object) int {
	return len(encode(r))
}

func TestCacheMiss(t *testing.T) {
	id := binary.NewID([]byte("123"))
	data := &binary.Data{1, 2, 3}
	size := calcSize(data)
	loadCallCount := 0

	inner := createMockStore()
	inner.onLoad = func(actualId binary.ID, out binary.Object) (int, error) {
		if id != actualId {
			t.Errorf("invalid id, expected %v got %v", id, actualId)
		}
		loadCallCount++
		CopyResource(out, data)
		return size, nil
	}

	cache := CreateCache(64, inner).(*cache)

	// Nothing should have called inner.Load() yet
	if 0 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 0, loadCallCount)
	}

	validateLoad(t, cache, id, size, data, nil)

	// The cache should have missed, calling inner.Load()
	if 1 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 1, loadCallCount)
	}
}

func TestCacheHit(t *testing.T) {
	id := binary.NewID([]byte("123"))
	data := &binary.Data{1, 2, 3}
	size := calcSize(data)
	loadCallCount := 0

	inner := createMockStore()
	inner.onLoad = func(actualId binary.ID, out binary.Object) (int, error) {
		if id != actualId {
			t.Errorf("invalid id, expected %v got %v", id, actualId)
		}
		loadCallCount++
		CopyResource(out, data)
		return size, nil
	}

	cache := CreateCache(64, inner).(*cache)

	// Warm the cache
	validateLoad(t, cache, id, size, data, nil)

	// The cache should have hit, skipping the call into inner.Load()
	validateLoad(t, cache, id, size, data, nil)
	if 1 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 1, loadCallCount)
	}
}

func TestCacheMultipleLoads(t *testing.T) {
	id := binary.NewID([]byte("123"))
	data := &binary.Data{1, 2, 3}
	size := calcSize(data)
	loadCallCount := 0
	loadSync := make(chan bool)
	loadUnblock := make(chan bool)
	loadComplete := make(chan bool)

	inner := createMockStore()
	inner.onLoad = func(actualId binary.ID, out binary.Object) (int, error) {
		if id != actualId {
			t.Errorf("invalid id, expected %v got %v", id, actualId)
		}
		loadCallCount++
		loadSync <- true
		<-loadUnblock
		CopyResource(out, data)
		return size, nil
	}

	cache := CreateCache(64, inner).(*cache)

	checkLoad := func() {
		validateLoad(t, cache, id, size, data, nil)
		loadComplete <- true
	}

	// Begin a number of loads, but don't let the inner load finish just yet.
	go checkLoad()
	go checkLoad()
	go checkLoad()

	// Sync with the inner load so we can query the load call count
	<-loadSync

	// There should only have been one call to the inner load
	if 1 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 1, loadCallCount)
	}

	// Unblock the inner load
	loadUnblock <- true

	// The three load requests should now finish
	<-loadComplete
	<-loadComplete
	<-loadComplete

	// There still should only have been one call to the inner load
	if 1 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 1, loadCallCount)
	}
}

func TestCacheOverflow(t *testing.T) {
	idA, idB, idC := binary.NewID([]byte("123")), binary.NewID([]byte("456")), binary.NewID([]byte("789"))
	dataA := &binary.Data{1, 2, 3, 4}
	sizeA := calcSize(dataA)
	dataB := &binary.Data{4, 5, 6, 7}
	sizeB := calcSize(dataB)
	dataC := &binary.Data{8, 9, 10, 11}
	sizeC := calcSize(dataC)
	loadCallCount := 0

	inner := createMockStore()
	inner.onLoad = func(id binary.ID, out binary.Object) (int, error) {
		loadCallCount++
		switch id {
		case idA:
			CopyResource(out, dataA)
			return sizeA, nil
		case idB:
			CopyResource(out, dataB)
			return sizeB, nil
		case idC:
			CopyResource(out, dataC)
			return sizeC, nil
		default:
			panic("Unexpected id!")
		}
	}

	// Create a cache that only has room for dataA and dataB
	cache := CreateCache(sizeA+sizeB, inner).(*cache)

	// Warm the cache with A and B
	validateLoad(t, cache, idA, sizeA, dataA, nil)
	validateLoad(t, cache, idB, sizeB, dataB, nil)

	// These should have both cache missed, calling into inner.Load()
	if 2 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 2, loadCallCount)
	}

	// Request load of a third resource, that will overflow the cache
	validateLoad(t, cache, idC, sizeC, dataC, nil)

	if 3 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 3, loadCallCount)
	}

	// The cache should now be holding B and C, with A evicted.
	// Requesting B should not call inner.Load() as it should be cached
	validateLoad(t, cache, idB, sizeB, dataB, nil)
	if 3 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 3, loadCallCount)
	}

	// Requesting C should not call inner.Load() as it should be cached
	validateLoad(t, cache, idC, sizeC, dataC, nil)
	if 3 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 3, loadCallCount)
	}

	// Requesting A should call inner.Load() as it should not be cached
	validateLoad(t, cache, idA, sizeA, dataA, nil)
	if 4 != loadCallCount {
		t.Errorf("invalid load call count, expected %v got %v", 4, loadCallCount)
	}
}

func TestCacheReplace(t *testing.T) {
	id := binary.NewID([]byte("123"))
	dataA := &binary.Data{1, 2, 3}
	dataB := &binary.Data{4, 5, 6, 7, 8, 9}
	sizeB := calcSize(dataB)

	inner := createMockStore()

	cache := CreateCache(64, inner).(*cache)

	// Fill the cache with dataA
	cache.Store(id, dataA, encode(dataA), nil)
	// Replace dataA with dataB
	cache.Store(id, dataB, encode(dataB), nil)

	// We expect dataB to be loaded
	validateLoad(t, cache, id, sizeB, dataB, nil)
}
