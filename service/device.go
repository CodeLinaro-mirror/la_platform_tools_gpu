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

package service

import (
	"fmt"
	"regexp"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/device"
)

type DeviceID binary.ID

// Device describes replay target avaliable to the server.
type Device struct {
	binary.Generate  `handle:"DeviceID"`
	Name             string // The name of the device. e.g. "Bob's phone"
	Model            string // The model of the device. e.g. "Nexus 5"
	OS               string // The operating system of the device. e.g. "Android 5.0"
	PointerSize      uint8  // Size in bytes of a pointer on the device's architecture.
	PointerAlignment uint8  // Alignment in bytes of a pointer on the device's architecture.
	MaxMemorySize    uint64 // The total amount of contiguous memory pre-allocated for replay.
	Extensions       string // Renderer extensions list. e.g. "GL_KHR_debug GL_EXT_sRGB [...]".
	Renderer         string // Renderer name. e.g. "Adreno (TM) 320".
	Vendor           string // Renderer vendor name. e.g. "Qualcomm".
	Version          string // Renderer version. e.g. "OpenGL ES 3.0 V@53.0 AU@  (CL@)".
}

// Architecture return's the device's architecture.
func (d Device) Architecture() device.Architecture {
	return device.Architecture{
		PointerAlignment: int(d.PointerAlignment),
		PointerSize:      int(d.PointerSize),
		IntegerSize:      int(d.PointerSize), // TODO: Resolve
		ByteOrder:        endian.Little,      // TODO: Resolve
	}
}

func (d Device) HasExtension(extension string) bool {
	if re, err := regexp.Compile(fmt.Sprintf(`\b%s\b`, extension)); err == nil {
		return re.MatchString(d.Extensions)
	}
	return false
}
