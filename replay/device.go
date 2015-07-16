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
	"io"
	"net"
	"os"
	"os/exec"
	"time"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// DisableLocalDeviceCache can be used to disable the disk-cache for the local
// device. If true, it is passed as a flag to replayd on spawning. This can be
// used for disabling the cache for tests.
var disableLocalDeviceCache = false

// localReplayBinary is the full path to the local binary.
var localReplayBinary = Replayd

// Port number of the "replayd" running on the local device
var localDevicePort = 9284

// Allow adjusting the settings for replayd.
func ConfigureLocalReplayDevice(disableCache bool, binary string, port int) {
	disableLocalDeviceCache = disableCache
	localReplayBinary = binary
	localDevicePort = port
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
	// Path returns the path for the replay device.
	Path() *path.Device
	// Info returns the service Device describing the replay device.
	Info() *service.Device
	// Connect opens a connection to the replay device.
	Connect() (io.ReadWriteCloser, error)
}

type deviceBase struct {
	path   *path.Device
	device *service.Device
}

func (d deviceBase) Path() *path.Device {
	return d.path
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
	conn, err := net.Dial("tcp", endpoint)
	if err != nil {
		if err := spawnChild(localReplayBinary); err != nil {
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
	if disableLocalDeviceCache {
		args = append(args, "--nocache")
	}
	args = append(args, "--port")
	args = append(args, fmt.Sprintf("%d", localDevicePort))
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
