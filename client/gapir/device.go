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
	"io"
	"net"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/process"
	"android.googlesource.com/platform/tools/gpu/service"
)

// DisableLocalDeviceCache can be used to disable the disk-cache for the local
// device. If true, it is passed as a flag to gapir on spawning. This can be
// used for disabling the cache for tests.
var disableLocalDeviceCache = false

// localReplayBinary is the full path to the local binary.
var localReplayBinary = Replayd

// Port number of the "gapir" running on the local device
var localDevicePort = 9284

// The path to the "gapir" log file for the local device.
var localLogPath = ""

// Disables launching of gapir
var localNoLaunch = false

// ConfigureLocalReplayDevice adjusts the settings for gapir.
func ConfigureLocalReplayDevice(disableCache bool, binary string, port int, logPath string, noLaunch bool) {
	disableLocalDeviceCache = disableCache
	localReplayBinary = binary
	localDevicePort = port
	localLogPath = logPath
	localNoLaunch = noLaunch
}

// deviceOS is an enumerator of operating systems that the replay target may be
// running on.
type deviceOS uint8

// These must be kept in sync with TARGET_OS in cc/common/gapic/target.h
const (
	osLinux   deviceOS = 1
	osOSX     deviceOS = 2
	osWindows deviceOS = 3
	osAndroid deviceOS = 4
)

func (os deviceOS) IsLinux() bool   { return os == osLinux }
func (os deviceOS) IsOSX() bool     { return os == osOSX }
func (os deviceOS) IsWindows() bool { return os == osWindows }
func (os deviceOS) IsAndroid() bool { return os == osAndroid }

func (os deviceOS) String() string {
	switch os {
	case osLinux:
		return "linux"
	case osOSX:
		return "darwin"
	case osWindows:
		return "windows"
	case osAndroid:
		return "android"
	default:
		return fmt.Sprintf("Unknown<%d>", os)
	}
}

// Device is the interface for a discovered replay device.
type Device interface {
	// ID returns the unique identifier for the replay device.
	ID() binary.ID
	// Info returns the service Device describing the replay device.
	Info() *service.Device
	// Connect opens a connection to the replay device.
	Connect() (io.ReadWriteCloser, error)
}

type deviceBase struct {
	id     binary.ID
	device *service.Device
}

func (d deviceBase) ID() binary.ID {
	return d.id
}

func (d deviceBase) Info() *service.Device {
	return d.device
}

type androidDevice struct {
	deviceBase
}

type localDevice struct {
	deviceBase
}

func (androidDevice) Connect() (io.ReadWriteCloser, error) {
	return net.Dial("tcp", "localhost:9285") // TODO: Remove the hardcoded port number.
}

func (localDevice) Connect() (io.ReadWriteCloser, error) {
	endpoint := fmt.Sprintf("localhost:%d", localDevicePort)
	if localNoLaunch {
		return net.Dial("tcp", endpoint)
	}

	args := []string{}
	if disableLocalDeviceCache {
		args = append(args, "--nocache")
	}
	args = append(args, "--port", fmt.Sprintf("%d", localDevicePort))

	if len(localLogPath) > 0 {
		args = append(args, "--log", localLogPath)
	}

	return process.ConnectStartIfNeeded(endpoint, localReplayBinary, args...)
}
