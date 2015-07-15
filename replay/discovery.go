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

	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/service"
)

// discovery is used to find replay devices on the local machine and connected
// Android devices.
type discovery struct {
	sync.Mutex
	devices map[service.DeviceID]Device
	logger  log.Logger
}

func newDiscovery(db database.Database, logger log.Logger) *discovery {
	m := &discovery{
		devices: make(map[service.DeviceID]Device),
		logger:  logger,
	}

	go m.discoverLocalDevices(db)
	go m.discoverAndroidDevices(db)

	return m
}

func (d *discovery) device(id service.DeviceID) Device {
	d.Lock()
	defer d.Unlock()

	return d.devices[id]
}

func (d *discovery) getDevices() []Device {
	d.Lock()
	defer d.Unlock()

	out := make([]Device, 0, len(d.devices))
	for _, d := range d.devices {
		out = append(out, d)
	}
	return out
}

func (m *discovery) discoverAndroidDevices(db database.Database) {
	d := &androidDevice{deviceBase{device: &service.Device{
		Name:  "Android device",
		Model: "Unknown",
	}}}

	if err := loadDeviceConfig(d, db, m.logger); err == nil {
		id, err := database.Store(d.device, db, log.Nop{})
		if err != nil {
			panic(err)
		}
		d.id = service.DeviceID(id)

		m.Lock()
		defer m.Unlock()
		m.devices[d.id] = d
	} else {
		log.Infof(m.logger, "Failed to communicate with Android device '%s': %v", d.Info().Name, err)
	}
}

func (m *discovery) discoverLocalDevices(db database.Database) {
	d := &localDevice{deviceBase{device: &service.Device{
		Name:  "Local machine",
		Model: "Unknown",
	}}}

	if err := loadDeviceConfig(d, db, m.logger); err == nil {
		id, err := database.Store(d.device, db, log.Nop{})
		if err != nil {
			panic(err)
		}
		d.id = service.DeviceID(id)

		m.Lock()
		defer m.Unlock()
		m.devices[d.id] = d
	} else {
		log.Warningf(m.logger, "Failed to communicate with local device '%s': %v", d.Info().Name, err)
	}
}

func loadDeviceConfig(d Device, db database.Database, logger log.Logger) error {
	connection, err := d.Connect()
	if err != nil {
		return err
	}
	defer connection.Close()

	// Endianness has yet to be discovered - we use Little for the hand-shaking.
	enc := flat.Encoder(endian.Writer(connection, endian.Little))
	dec := flat.Decoder(endian.Reader(connection, endian.Little))

	if err := enc.Uint8(uint8(protocol.ConnectionTypeDeviceInfo)); err != nil {
		return err
	}

	protocolVersion, err := dec.Uint32()
	if err != nil {
		return err
	}

	td := d.Info()

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

		// TODO: Integer size
		// TODO: Endianness

		td.MaxMemorySize, err = dec.Uint64()
		if err != nil {
			return err
		}

		var os deviceOS
		if val, err := dec.Uint8(); err == nil {
			os = deviceOS(val)
		} else {
			return err
		}
		td.OS = os.String()

		if td.Extensions, err = dec.String(); err != nil {
			return err
		}
		if td.Renderer, err = dec.String(); err != nil {
			return err
		}
		if td.Vendor, err = dec.String(); err != nil {
			return err
		}
		if td.Version, err = dec.String(); err != nil {
			return err
		}

	default:
		return fmt.Errorf("Unsupported device protocol version: %d", protocolVersion)
	}
	return nil
}
