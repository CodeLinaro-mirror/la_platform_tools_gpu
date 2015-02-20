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
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
)

const extension = ".capture"

// Database is the interface to a resource store.
type Database interface {
	StoreLink(to, id binary.ID, logger log.Logger) error
	StoreRequest(request binary.Object, logger log.Logger) (id binary.ID, err error)
	Store(binary.Object, log.Logger) (binary.ID, error)
	Load(binary.ID, log.Logger, binary.Object) error
	Contains(binary.ID, log.Logger) bool
	Captures() (map[string]binary.ID, error)
	Close()
}

// Create builds a new database.
func Create(path string, maxDataCacheSize, metaDataCompactionSize, maxDerivedCacheSize int, builder builder) Database {
	var dataStore store.Store
	dataStore = store.CreateUnboundedArchive(filepath.Join(path, "resource", "data"))
	dataStore = store.CreateStoreIfNew(dataStore)
	dataStore = store.CreateCache(maxDataCacheSize, dataStore)

	var derivedStore store.Store
	derivedStore = store.CreateUnboundedArchive(filepath.Join(path, "resource", "derived"))
	derivedStore = store.CreateStoreIfNew(derivedStore)
	derivedStore = store.CreateCache(maxDerivedCacheSize, derivedStore)

	var metaStore store.Store
	metaStore = store.CreateSmallArchive(filepath.Join(path, "resource", "meta"), metaDataCompactionSize)

	return &database{
		path:         path,
		dataStore:    dataStore,
		metaStore:    metaStore,
		derivedStore: derivedStore,
		pending:      make(map[binary.ID]*pending),
		builder:      builder,
	}
}

type pending struct {
	res binary.Object
	err error
	wg  sync.WaitGroup
}

type database struct {
	path         string
	dataStore    store.Store
	metaStore    store.Store
	derivedStore store.Store
	mutex        sync.Mutex // guards against pending
	pending      map[binary.ID]*pending
	builder      builder
}

func (d *database) loadMetadata(id binary.ID, logger log.Logger, metadata *metadata) error {
	logger = logger.Enter("Database.loadMetadata")
	_, err := d.metaStore.Load(id, logger, metadata)
	return err
}

func (d *database) loadMetadataIfExists(id binary.ID, logger log.Logger, metadata *metadata) error {
	if d.metaStore.Contains(id) {
		return d.loadMetadata(id, logger, metadata)
	}
	return nil
}

func (d *database) storeMetadata(id binary.ID, metadata *metadata, logger log.Logger) error {
	logger = logger.Enter("Database.storeMetadata")
	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	if err := metadata.Encode(enc); err != nil {
		return err
	}
	return d.metaStore.Store(id, metadata, buf.Bytes(), logger)
}

func (d *database) SetBuilder(b builder) {
	d.builder = b
}

func (d *database) StoreLink(to, id binary.ID, logger log.Logger) (err error) {
	logger = logger.Enter("Database.StoreLink")
	logger.Info("(to: %v, id: %v)", to, id)
	defer func() { logger.Info("↪ err: %v", err) }()

	if to == id {
		return nil // Link to itself, ignore
	}

	metadata := &metadata{}
	d.loadMetadataIfExists(id, logger, metadata)
	metadata.Type = metaTypeLink
	metadata.LinkTo = to
	return d.storeMetadata(id, metadata, logger)
}

func (d *database) store(r binary.Object, logger log.Logger, metaType metaType, storeToUse store.Store) (id binary.ID, err error) {
	// Encode the resource
	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	if err := r.Encode(enc); err != nil {
		return binary.ID{}, err
	}

	// Calculate the resource hash, this is the resource id
	data := buf.Bytes()
	id = binary.NewID(data)
	if err := storeToUse.Store(id, r, data, logger); err != nil {
		return binary.ID{}, err
	}

	// Store the metadata for the resource
	metadata := &metadata{}
	d.loadMetadataIfExists(id, logger, metadata)
	metadata.Type = metaType
	metadata.LinkTo = binary.ID{}
	if err := d.storeMetadata(id, metadata, logger); err != nil {
		return binary.ID{}, err
	}

	return id, nil
}

func (d *database) storeDerived(r binary.Object, logger log.Logger) (id binary.ID, err error) {
	logger = logger.Enter("Database.storeDerived")
	defer func() { logger.Info("↪ id: %v, err: %v", id, err) }()
	return d.store(r, logger, metaTypeDerived, d.derivedStore)
}

func (d *database) Store(r binary.Object, logger log.Logger) (id binary.ID, err error) {
	logger = logger.Enter("Database.Store")
	defer func() { logger.Info("↪ id: %v, err: %v", id, err) }()
	return d.store(r, logger, metaTypeData, d.dataStore)
}

func (d *database) StoreRequest(request binary.Object, logger log.Logger) (id binary.ID, err error) {
	logger = logger.Enter("Database.StoreRequest")
	logger.Info("(%+v)", request)
	defer func() { logger.Info("↪ id: %v, err: %v", id, err) }()

	buf := &bytes.Buffer{}
	enc := binary.NewEncoder(buf)
	if err := enc.Object(request); err != nil {
		return binary.ID{}, err
	}
	requestData := buf.Bytes()

	version := d.builder.Version()
	requestId := binary.NewID(requestData, []byte{
		byte(version),
		byte(version >> 8),
		byte(version >> 16),
		byte(version >> 24),
	})
	metadata := &metadata{}
	err = d.loadMetadataIfExists(requestId, logger, metadata)
	if err == nil && metadata.Type == metaTypeLink {
		// TODO: Check request data matches
		logger.Info("Resource already built")
		return requestId, nil
	}

	// Store the pending state and request data in the metadata
	metadata.Type = metaTypeLazy
	metadata.Request = request
	if err = d.storeMetadata(requestId, metadata, logger); err != nil {
		return binary.ID{}, err
	}

	return requestId, nil
}

func (d *database) Load(id binary.ID, logger log.Logger, out binary.Object) (err error) {
	logger = logger.Enter("Database.Load")
	logger.Info("(id: %v)", id)
	defer func() { logger.Info("↪ err: %v", err) }()

	d.mutex.Lock()
	if pending, found := d.pending[id]; found {
		d.mutex.Unlock()
		pending.wg.Wait()
		store.CopyResource(out, pending.res)
		return pending.err
	}

	p := &pending{}
	p.wg.Add(1)
	d.pending[id] = p
	d.mutex.Unlock()

	defer func() {
		p.err = err
		p.res = out
		p.wg.Done()

		d.mutex.Lock()
		delete(d.pending, id)
		d.mutex.Unlock()
	}()

	metadata := &metadata{}
	if err := d.loadMetadata(id, logger, metadata); err != nil {
		return err
	}
	logger.Info("Metadata: %+v", metadata)

	switch metadata.Type {
	case metaTypeLink:
		return d.Load(metadata.LinkTo, logger, out)
	case metaTypeLazy:
		logger.Info("Recreating resource")

		// Begin building of the resource
		if err := d.builder.BuildResource(metadata.Request, d, logger, out); err != nil {
			return err
		}

		// Store the resource
		resourceId, err := d.storeDerived(out, logger)
		if err != nil {
			return err
		}

		// Update the metadata from Lazy to Link
		if d.StoreLink(resourceId, id, logger); err != nil {
			return err
		}

		return nil
	case metaTypeData:
		_, err := d.dataStore.Load(id, logger, out)
		return err
	case metaTypeDerived:
		_, err := d.derivedStore.Load(id, logger, out)
		return err
	default:
		err := fmt.Errorf("Unknown metadata type %v", metadata.Type)
		logger.Error("%v", err)
		return err
	}
}

func (d *database) Contains(id binary.ID, logger log.Logger) (res bool) {
	logger = logger.Enter("Database.Contains")
	logger.Info("(id: %v)", id)
	defer func() { logger.Info("↪ %v", res) }()
	return d.metaStore.Contains(id)
}

func (d *database) Close() {
	d.metaStore.Close()
	d.dataStore.Close()
}

func (d *database) Captures() (map[string]binary.ID, error) {
	result := make(map[string]binary.ID)
	files, err := ioutil.ReadDir(d.path)
	if err != nil {
		return result, err
	}
	for _, file := range files {
		filename := file.Name()
		if filepath.Ext(filename) == extension {
			name := strings.TrimSuffix(filename, extension)
			path := filepath.Join(d.path, filename)
			if r, err := os.Open(path); err != nil {
				return result, err
			} else {
				defer r.Close()
				d := binary.NewDecoder(r)
				id := binary.ID{}
				if err := id.Decode(d); err != nil {
					return result, err
				} else {
					result[name] = id
				}
			}
		}
	}
	return result, nil
}
