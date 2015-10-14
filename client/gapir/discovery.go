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

package gapir

import (
	"fmt"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/service"
)

// discovery is used to find replay devices on the local machine and connected
// Android devices.
type Discovery struct {
	sync.Mutex
	devices []Device
	logger  log.Logger
}

func NewDiscovery(db database.Database, logger log.Logger) *Discovery {
	m := &Discovery{logger: logger}
	go m.discoverLocalDevices(db)
	go m.discoverAndroidDevices(db)

	return m
}

func (d *Discovery) Device(id binary.ID) Device {
	d.Lock()
	defer d.Unlock()
	for _, d := range d.devices {
		if d.ID() == id {
			return d
		}
	}
	return nil
}

func (d *Discovery) Devices() []Device {
	d.Lock()
	defer d.Unlock()

	return d.devices
}

func (d *Discovery) discoverAndroidDevices(db database.Database) {
	dev := &androidDevice{deviceBase{device: &service.Device{
		Name:  "Android device",
		Model: "Unknown",
	}}}

	if err := loadDeviceConfig(dev, db, d.logger); err == nil {
		id, err := database.Store(dev.device, db, log.Nop{})
		if err != nil {
			panic(err)
		}
		dev.id = id

		d.Lock()
		defer d.Unlock()
		d.devices = append(d.devices, dev)
	} else {
		log.Infof(d.logger, "Failed to communicate with Android device '%s': %v", dev.Info().Name, err)
	}
}

func (d *Discovery) discoverLocalDevices(db database.Database) {
	dev := &localDevice{deviceBase{device: &service.Device{
		Name:  "Local machine",
		Model: "Unknown",
	}}}

	if err := loadDeviceConfig(dev, db, d.logger); err == nil {
		id, err := database.Store(dev.device, db, log.Nop{})
		if err != nil {
			panic(err)
		}
		dev.id = id

		d.Lock()
		defer d.Unlock()
		d.devices = append(d.devices, dev)
	} else {
		log.Warningf(d.logger, "Failed to communicate with local device '%s': %v", dev.Info().Name, err)
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

	if enc.Uint8(uint8(protocol.ConnectionTypeDeviceInfo)); enc.Error() != nil {
		return enc.Error()
	}

	protocolVersion := dec.Uint32()
	if dec.Error() != nil {
		return dec.Error()
	}

	td := d.Info()

	switch protocolVersion {
	case 1:
		td.PointerSize = dec.Uint8()
		td.PointerAlignment = dec.Uint8()
		// TODO: Integer size
		// TODO: Endianness
		td.MaxMemorySize = dec.Uint64()
		td.OS = deviceOS(dec.Uint8()).String()
		td.Extensions = dec.String()
		td.Renderer = dec.String()
		td.Vendor = dec.String()
		td.Version = dec.String()
		if dec.Error() != nil {
			return dec.Error()
		}

	default:
		return fmt.Errorf("Unsupported device protocol version: %d", protocolVersion)
	}
	return nil
}
