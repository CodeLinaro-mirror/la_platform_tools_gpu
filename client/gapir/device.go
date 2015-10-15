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

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/client/process"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/service"
)

const (
	// LocalPort is a legacy fixed port constant for local replay.
	LocalPort = 9284
	// LocalName
	LocalName = "Local machine"
	// AndroidPort is a legacy fixed port forwarded to android devices.
	AndroidPort = 9285
	// AndroidName
	AndroidName = "Android device"
)

var (
	// LogPath is the full filepath of the logfile new instances of gapir should write to.
	LogPath string
)

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
	port   int
}

func (d *deviceBase) ID() binary.ID {
	return d.id
}

func (d *deviceBase) Info() *service.Device {
	return d.device
}

func (d *deviceBase) Connect() (io.ReadWriteCloser, error) {
	return process.Connect(d.port)
}

func NewDevice(name string, port int, logger log.Logger) *deviceBase {
	device := &deviceBase{
		device: &service.Device{
			Name:  name,
			Model: "Unknown",
		},
		port: port,
	}
	err := device.loadInfo()
	if err != nil {
		log.Errorf(logger, "Failed to connect to device %s: %s", name, err)
		return nil
	}
	return device
}

func (d *deviceBase) loadInfo() error {
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

	switch protocolVersion {
	case 1:
		d.device.PointerSize = dec.Uint8()
		d.device.PointerAlignment = dec.Uint8()
		// TODO: Integer size
		// TODO: Endianness
		d.device.MaxMemorySize = dec.Uint64()
		d.device.OS = deviceOS(dec.Uint8()).String()
		d.device.Extensions = dec.String()
		d.device.Renderer = dec.String()
		d.device.Vendor = dec.String()
		d.device.Version = dec.String()
		if dec.Error() != nil {
			return dec.Error()
		}

	default:
		return fmt.Errorf("Unsupported device protocol version: %d", protocolVersion)
	}

	d.id, err = database.Hash(d.device)
	return err
}

func RunLocal(logger log.Logger) *deviceBase {
	args := []string{
		"--port", "0",
		"--nocache",
	}
	if LogPath != "" {
		args = append(args, "--log", LogPath)
	}
	port, err := process.Start("gapir", args...)
	if err != nil {
		log.Errorf(logger, "Failed starting gapir: %s", err)
		return nil
	}
	return NewDevice(LocalName, port, logger)
}
