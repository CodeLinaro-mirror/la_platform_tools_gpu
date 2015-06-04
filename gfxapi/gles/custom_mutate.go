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

package gles

import (
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
)

func (ϟa *Architecture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error {
	bo := endian.Big
	if ϟa.LittleEndian {
		bo = endian.Little
	}
	ϟs.Architecture = device.Architecture{
		PointerAlignment: int(ϟa.PointerAlignment),
		PointerSize:      int(ϟa.PointerSize),
		IntegerSize:      int(ϟa.IntegerSize),
		ByteOrder:        bo,
	}
	return nil
}
