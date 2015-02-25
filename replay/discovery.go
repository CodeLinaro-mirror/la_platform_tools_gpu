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
	"runtime"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// discovery is used to find replay devices on the local machine and connected
// Android devices.
type discovery struct {
	sync.Mutex
	devices map[service.DeviceId]device
	logger  log.Logger
}

func newDiscovery(db database.Database, logger log.Logger) *discovery {
	m := &discovery{
		devices: make(map[service.DeviceId]device),
		logger:  logger,
	}

	go m.discoverLocalDevices(db)
	go m.discoverAndroidDevices(db)

	return m
}

func (d *discovery) device(id service.DeviceId) device {
	d.Lock()
	defer d.Unlock()

	return d.devices[id]
}

func (d *discovery) deviceIDs() (ids service.DeviceIdArray) {
	d.Lock()
	defer d.Unlock()

	for d := range d.devices {
		ids = append(ids, d)
	}
	return
}

func (m *discovery) discoverAndroidDevices(db database.Database) {
	d := &androidDevice{deviceBase{device: &service.Device{
		Name:  "Android device",
		Model: "Unknown",
		OS:    "Unknown",
	}}}

	if loadDeviceConfig(d, db, m.logger) == nil {
		d.transportDevice().RequiresShaderPatching = false // HACK FIXME

		id, err := db.Store(d.device, log.Nop{})
		if err != nil {
			panic(err)
		}
		d.id.ID = id

		m.Lock()
		defer m.Unlock()
		m.devices[d.id] = d
	}
}

func (m *discovery) discoverLocalDevices(db database.Database) {
	d := &localDevice{deviceBase{device: &service.Device{
		Name:  "Local machine",
		Model: runtime.GOARCH,
		OS:    runtime.GOOS,
	}}}

	if err := loadDeviceConfig(d, db, m.logger); err == nil {
		d.transportDevice().RequiresShaderPatching = true // HACK FIXME

		id, err := db.Store(d.device, log.Nop{})
		if err != nil {
			panic(err)
		}
		d.id.ID = id

		m.Lock()
		defer m.Unlock()
		m.devices[d.id] = d
	}
}

func loadDeviceConfig(d device, db database.Database, logger log.Logger) (err error) {
	defer func() {
		if err != nil {
			logger.Error("Failed to load device '%s' config: %v", d.transportDevice().Name, err)
		}
	}()

	connection, err := d.connect()
	if err != nil {
		return err
	}
	defer connection.Close()

	enc := binary.NewEncoder(connection)
	dec := binary.NewDecoder(connection)

	if err := enc.Uint8(connectionTypeDeviceInfo); err != nil {
		return err
	}

	protocolVersion, err := dec.Uint32()
	if err != nil {
		return err
	}

	td := d.transportDevice()

	switch protocolVersion {
	case 1:
		td.PointerSize, err = dec.Uint8()
		if err != nil {
			return err
		}

		td.PointerAlignment, err = dec.Uint8()
		if err != nil {
			return err
		}

		td.MaxMemorySize, err = dec.Uint64()
		if err != nil {
			return err
		}

	default:
		return fmt.Errorf("Unsupported device protocol version: %d", protocolVersion)
	}

	return nil
}
