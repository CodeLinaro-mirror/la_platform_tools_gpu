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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// Path returns a path.Capture representing the capture with this identifier.
func (c CaptureID) Path() *path.Capture {
	return &path.Capture{ID: binary.ID(c)}
}

// Path returns a path.Device representing the device with this identifier.
func (c DeviceID) Path() *path.Device {
	return &path.Device{ID: binary.ID(c)}
}
