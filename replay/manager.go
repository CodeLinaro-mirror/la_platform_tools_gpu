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

package replay

import (
	"fmt"
	"sync"

	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Manager is used discover replay devices and to send replay requests to those
// discovered devices.
type Manager struct {
	persistentDb database.Database
	transientDb  database.Database
	discovery    *discovery
	batchers     map[batcherContext]*batcher
	mutex        sync.Mutex // guards batchers
	logger       log.Logger
}

func (m *Manager) getBatchStream(ctx batcherContext) (chan<- Request, error) {
	m.mutex.Lock()
	defer m.mutex.Unlock()

	// TODO: This accumulates batchers with running go-routines forever.
	// Rework to free the batcher after execution.
	b, found := m.batchers[ctx]
	if !found {
		device := m.discovery.device(ctx.DeviceID)
		if device == nil {
			return nil, fmt.Errorf("Unknown device %v", ctx.DeviceID)
		}
		b = &batcher{
			context:      ctx,
			feed:         make(chan Request, 8),
			persistentDb: m.persistentDb,
			transientDb:  m.transientDb,
			device:       device,
			logger:       m.logger,
		}
		m.batchers[ctx] = b
		go b.run()
	}
	return b.feed, nil
}

// New returns a new Manager instance using the database db and logger l.
func New(db database.Database, l log.Logger) *Manager {
	return &Manager{
		persistentDb: db,
		transientDb:  database.CreateTransientDatabase(db),
		discovery:    newDiscovery(db, l),
		batchers:     make(map[batcherContext]*batcher),
		logger:       l,
	}
}

// Replay requests that req is to be performed on the device described by ctx,
// using the capture described by ctx. Replay is asynchronous, and the replay
// may take some considerable time before it is executed. Replay requests made
// with configs that have equality (==) will likely be batched into the same
// replay pass.
func (m *Manager) Replay(ctx *Context, cfg Config, req Request, generator Generator) error {
	batch, err := m.getBatchStream(batcherContext{
		Context:   *ctx,
		Generator: generator,
		Config:    cfg,
	})
	if err == nil {
		batch <- req
	}
	return err
}

// DeviceIDs returns the list of devices that have been discovered.
func (m *Manager) Devices() []Device {
	return m.discovery.getDevices()
}
