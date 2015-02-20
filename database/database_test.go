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
	"os"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
)

var testRequestTypeId = binary.ID{0x12, 0x34}

type testResource struct {
	Int    int
	String string
	Array  []bool
}

func init() {
	binary.Register(testRequestTypeId, &testRequest{})
}

func (r *testResource) Encode(e *binary.Encoder) error {
	if err := e.Int32(int32(r.Int)); err != nil {
		return err
	}
	if err := e.String(r.String); err != nil {
		return err
	}
	if err := e.Int32(int32(len(r.Array))); err != nil {
		return err
	}
	for _, v := range r.Array {
		if err := e.Bool(bool(v)); err != nil {
			return err
		}
	}
	return nil
}

func (r *testResource) Decode(d *binary.Decoder) error {
	if v, err := d.Int32(); err == nil {
		r.Int = int(v)
	} else {
		return err
	}
	if v, err := d.String(); err == nil {
		r.String = v
	} else {
		return err
	}
	if v, err := d.Int32(); err == nil {
		r.Array = make([]bool, v)
		for i := range r.Array {
			if b, err := d.Bool(); err == nil {
				r.Array[i] = b
			} else {
				return err
			}
		}
	} else {
		return err
	}
	return nil
}

var testResourceA = &testResource{
	Int:    50,
	String: "fifty",
	Array:  []bool{true, true, false, false, true, false},
}

type testRequest struct {
	Id int
}

func (t *testRequest) Encode(e *binary.Encoder) error {
	return e.Int32(int32(t.Id))
}

func (t *testRequest) Decode(d *binary.Decoder) error {
	id, err := d.Int32()
	t.Id = int(id)
	return err
}

func decodeTestRequest(d *binary.Decoder) (binary.Object, error) {
	o := &testRequest{}
	e := o.Decode(d)
	return o, e
}

type testStoreEntry struct {
	resource binary.Object
	data     []byte
}

type testStore struct {
	entries map[binary.ID]testStoreEntry
}

func (s testStore) Store(id binary.ID, r binary.Object, d []byte, _ log.Logger) error {
	s.entries[id] = testStoreEntry{r, d}
	return nil
}

func (s testStore) Load(id binary.ID, _ log.Logger, out binary.Object) (size int, err error) {
	if e, found := s.entries[id]; found {
		store.CopyResource(out, e.resource)
		return len(e.data), nil
	} else {
		return -1, os.ErrNotExist
	}
}
func (s testStore) Contains(id binary.ID) bool {
	_, found := s.entries[id]
	return found
}
func (s testStore) Close() {
}

type testBuilder struct {
	buildResource func(req interface{}, logger log.Logger, out binary.Object) error
}

func (b testBuilder) BuildResource(req interface{}, _ Database, logger log.Logger, out binary.Object) error {
	return b.buildResource(req, logger, out)
}

func (b testBuilder) Version() uint32 {
	return 0
}

func create(builder builder) (db *database, ds, vs, ms *testStore) {
	if builder == nil {
		builder = &testBuilder{}
	}
	ds = &testStore{make(map[binary.ID]testStoreEntry)}
	vs = &testStore{make(map[binary.ID]testStoreEntry)}
	ms = &testStore{make(map[binary.ID]testStoreEntry)}
	db = &database{
		dataStore:    ds,
		derivedStore: vs,
		metaStore:    ms,
		pending:      make(map[binary.ID]*pending),
		builder:      builder,
	}
	return
}

func verifyCounts(t *testing.T, ds, vs, ms *testStore, ed, ev, em int) {
	if len(ds.entries) != ed {
		t.Fatalf("expected 0 entries in data store, got %v", len(ds.entries))
	}
	if len(vs.entries) != ev {
		t.Fatalf("expected 0 entries in derived store, got %v", len(vs.entries))
	}
	if len(ms.entries) != em {
		t.Fatalf("expected 0 entries in meta store, got %v", len(ms.entries))
	}
}

func verifyResource(t *testing.T, expected, got *testResource) {
	if !reflect.DeepEqual(expected, got) {
		t.Fatalf("resource  does not match, expected %v, got %v", expected, got)
	}
}

func TestStoreLink(t *testing.T) {
	fromId, toId := binary.NewID([]byte("123")), binary.NewID([]byte("345"))
	db, ds, vs, ms := create(nil)
	err := db.StoreLink(toId, fromId, log.Nop{})
	if err != nil {
		panic(err)
	}
	verifyCounts(t, ds, vs, ms, 0, 0, 1)
	// Assert the single metadata entry is as expected
	metadata := ms.entries[fromId].resource.(*metadata)
	if metadata.Type != metaTypeLink {
		t.Errorf("Incorrect meta type, expected link got %v", metadata.Type)
	}
	if metadata.LinkTo != toId {
		t.Errorf("Incorrect link id, expected %v got %v", toId, metadata.LinkTo)
	}
}

func TestStoreDataRequest(t *testing.T) {
	db, ds, vs, ms := create(nil)
	id, err := db.StoreRequest(&testRequest{123}, log.Nop{})
	if err != nil {
		panic(err)
	}
	verifyCounts(t, ds, vs, ms, 0, 0, 1)
	// Assert the single metadata entry is of lazy type
	metadata := ms.entries[id].resource.(*metadata)
	if metadata.Type != metaTypeLazy {
		t.Errorf("Incorrect meta type, expected lazy got %v", metadata.Type)
	}
}

func TestStore(t *testing.T) {
	db, ds, vs, ms := create(nil)
	id, err := db.Store(testResourceA, log.Nop{})
	if err != nil {
		panic(err)
	}
	verifyCounts(t, ds, vs, ms, 1, 0, 1)
	// Assert the single meta entry is as expected
	metadata := ms.entries[id].resource.(*metadata)
	if metadata.Type != metaTypeData {
		t.Errorf("Incorrect meta type, expected data got %v", metadata.Type)
	}
	// Assert the single data entry resource is as expected
	data := ds.entries[id].resource.(*testResource)
	verifyResource(t, testResourceA, data)
	// Assert the single data entry binary data is as expected
	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	testResourceA.Encode(enc)
	if !reflect.DeepEqual(buf.Bytes(), ds.entries[id].data) {
		t.Fatalf("encoded data did not match")
	}
}

func TestLoadData(t *testing.T) {
	db, _, _, _ := create(nil)
	id, _ := db.Store(testResourceA, log.Nop{})

	data := &testResource{}
	err := db.Load(id, log.Nop{}, data)
	if err != nil {
		panic(err)
	}
	// Assert loaded data was as expected
	verifyResource(t, testResourceA, data)
}

func TestLoadDataViaSingleLink(t *testing.T) {
	db, _, _, _ := create(nil)
	dataId, _ := db.Store(testResourceA, log.Nop{})
	linkId := binary.NewID([]byte("123"))
	db.StoreLink(dataId, linkId, log.Nop{})

	data := &testResource{}
	err := db.Load(linkId, log.Nop{}, data)
	if err != nil {
		panic(err)
	}
	// Assert loaded data was as expected
	verifyResource(t, testResourceA, data)
}

func TestLoadDataViaChainedLink(t *testing.T) {
	db, _, _, _ := create(nil)
	dataId, _ := db.Store(testResourceA, log.Nop{})
	linkAId := binary.NewID([]byte("123"))
	linkBId := binary.NewID([]byte("456"))
	linkCId := binary.NewID([]byte("678"))
	db.StoreLink(dataId, linkAId, log.Nop{})
	db.StoreLink(linkAId, linkBId, log.Nop{})
	db.StoreLink(linkBId, linkCId, log.Nop{})

	data := &testResource{}
	err := db.Load(linkCId, log.Nop{}, data)
	if err != nil {
		panic(err)
	}
	// Assert loaded data was as expected
	verifyResource(t, testResourceA, data)
}

func TestLoadDataRequest(t *testing.T) {
	syncBegin := make(chan bool)
	syncEnd := make(chan bool)
	builder := &testBuilder{}
	request := &testRequest{123}
	db, ds, vs, ms := create(builder)

	id, _ := db.StoreRequest(request, log.Nop{})

	builder.buildResource = func(actualRequest interface{}, _ log.Logger, p binary.Object) error {
		<-syncBegin
		if !reflect.DeepEqual(request, actualRequest) {
			t.Fatalf("request did not match")
		}
		store.CopyResource(p, testResourceA)
		syncEnd <- true
		return nil
	}

	data := &testResource{}

	// Assert no entries were created in the derived store before "Load"
	verifyCounts(t, ds, vs, ms, 0, 0, 1)

	loadErr := make(chan error, 1)
	go func() { loadErr <- db.Load(id, log.Nop{}, data) }()

	// Assert that the resource hasn't been set until the builder has processed
	verifyResource(t, &testResource{}, data)

	// Let the builder do its thing
	syncBegin <- true
	<-syncEnd

	// Sync and get the returned error
	err := <-loadErr
	// Assert that the error was as expected
	if err != nil {
		t.Errorf("Unexpected error %s", err)
	}

	verifyCounts(t, ds, vs, ms, 0, 1, 2)

	// Assert loaded data was as expected
	verifyResource(t, testResourceA, data)
}
