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
	"encoding/binary"
	"fmt"
	"io"
	"net"
	"os"
	"os/exec"
	"time"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/service"
)

// DisableLocalDeviceCache can be used to disable the disk-cache for the local
// device. If true, it is passed as a flag to replayd on spawning. This can be
// used for disabling the cache for tests.
var DisableLocalDeviceCache = false

// deviceOS is an enumerator of operating systems that the replay target may be
// running on.
type deviceOS uint8

// These must be kept in sync with TARGET_OS in cc/replayd/src/Target.h
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
	// ID returns the identifier for the replay device.
	ID() service.DeviceId
	// Info returns the service Device describing the replay device.
	Info() *service.Device
	// Connect opens a connection to the replay device.
	Connect() (io.ReadWriteCloser, error)
	// ByteOrder returns a byte ordering object for the replay device.
	ByteOrder() binary.ByteOrder
}

type deviceBase struct {
	id     service.DeviceId
	device *service.Device
}

func (d deviceBase) ID() service.DeviceId {
	return d.id
}

func (d deviceBase) Info() *service.Device {
	return d.device
}

func (d deviceBase) ByteOrder() binary.ByteOrder {
	// TODO: vary the endian based on the device info
	return binary.LittleEndian
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
	endpoint := "localhost:9284" // TODO: Remove the hardcoded port number.
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		if err := spawnChild(replayd); err != nil {
			return nil, err
		}
		for i := 0; i < 10; i++ {
			conn, err = net.Dial("tcp", endpoint)
			if err == nil {
				return conn, err
			}
			time.Sleep(100 * time.Millisecond)
		}
	}
	return conn, err
}

func spawnChild(path string) error {
	null, err := os.OpenFile(os.DevNull, os.O_RDWR, 0)
	if err != nil {
		return err
	}
	path, err = exec.LookPath(path)
	if err != nil {
		return err
	}
	args := []string{path}
	if DisableLocalDeviceCache {
		args = append(args, "--nocache")
	}
	proc, err := os.StartProcess(path, args, &os.ProcAttr{
		Files: []*os.File{null, null, null},
	})
	if err != nil {
		return err
	}

	// Kill the child process on parent exit.
	atexit.Register(func() {
		proc.Kill()
		proc.Wait()
	}, time.Second)

	return nil
}
