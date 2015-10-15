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
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// discovery is used to find replay devices on the local machine and connected
// Android devices.
type Discovery struct {
	sync.Mutex
	devices []Device
	logger  log.Logger
}

func NewDiscovery(db database.Database, logger log.Logger) *Discovery {
	return &Discovery{logger: logger}
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

func (d *Discovery) DefaultDevice() Device {
	d.Lock()
	defer d.Unlock()
	if len(d.devices) == 0 {
		return nil
	}
	return d.devices[0]
}

func (d *Discovery) AddDevice(dev *deviceBase, db database.Database, l log.Logger) {
	if dev != nil {
		d.Lock()
		defer d.Unlock()
		d.devices = append(d.devices, dev)
		db.Store(dev.ID(), dev.Info(), l)
	}
}
