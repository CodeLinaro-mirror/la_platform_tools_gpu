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

package adb

import (
	"errors"
	"fmt"
	"strings"
)

// DeviceState represents the last queried state of an Android device.
type DeviceState int

const (
	Offline DeviceState = iota
	Online
	Unauthorized
)

// Device represents an attached Android device.
type Device struct {
	Serial string
	State  DeviceState
}

// Command returns a new Cmd that will run the command with the specified name
// and arguments on this device.
func (d *Device) Command(path string, args ...string) *Cmd {
	return &Cmd{
		Path:   path,
		Args:   args,
		Device: d,
	}
}

// SELinuxEnforcing returns true if the device is currently in a
// SELinux enforcing mode, or false if the device is currently in a SELinux
// permissive mode.
func (d *Device) SELinuxEnforcing() (bool, error) {
	res, err := d.Command("getenforce").Call()
	return strings.Contains(strings.ToLower(res), "enforcing"), err
}

// SetSELinuxEnforcing changes the SELinux-enforcing mode.
func (d *Device) SetSELinuxEnforcing(enforce bool) error {
	if enforce {
		return d.Command("setenforce", "1").Run()
	} else {
		return d.Command("setenforce", "0").Run()
	}
}

// StartActivity launches the specified action.
func (d *Device) StartActivity(a Action) error {
	return d.Command("am", "start",
		//		"-W", // Wait for launch to complete
		"-S", // Force-stop the target app before starting the activity
		"-a", a.Name,
		"-n", a.Component).Run()
}

// String returns a string representing the device.
func (d *Device) String() string {
	return fmt.Sprintf("Device<%s>", d.Serial)
}

// Devices returns the list of serial numbers of all the attached Android
// devices.
func Devices() ([]*Device, error) {
	if adb == "" {
		return nil, ErrADBNotFound
	}
	if out, err := run("devices"); err == nil {
		return parseDevices(out)
	} else {
		return nil, err
	}
}

func parseDevices(out string) ([]*Device, error) {
	a := strings.SplitAfter(out, "List of devices attached")
	if len(a) != 2 {
		return nil, errors.New("Device list not returned")
	}
	lines := strings.Split(a[1], "\n")
	devices := make([]*Device, 0, len(lines))
	for _, line := range lines {
		fields := strings.Fields(line)
		switch len(fields) {
		case 0:
			continue
		case 2:
			state := DeviceState(0)
			if err := state.parse(fields[1]); err != nil {
				return nil, err
			}
			device := &Device{
				Serial: fields[0],
				State:  state,
			}
			devices = append(devices, device)
		default:
			return nil, errors.New("Could not parse device list")
		}
	}
	return devices, nil
}

func (s *DeviceState) parse(str string) error {
	switch str {
	case "offline":
		*s = Offline
		return nil
	case "unauthorized":
		*s = Unauthorized
		return nil
	case "device":
		*s = Online
		return nil
	}
	return fmt.Errorf("Unknown device state '%s'", str)
}
