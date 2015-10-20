// Copyright (C) 2014 The Android Open Source Project
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

package cyclic

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/errors"
)

// Control represents a control block in a stream.
type Control struct {
	Mode binary.Mode
}

const (
	ControlVersion uint32 = 0
)

func (c *Control) write(e *encoder) {
	e.Uint32(ControlVersion)
	e.Uint32(uint32(c.Mode))
}

func (c *Control) read(d *decoder) {
	version := d.Uint32()
	switch version {
	case 0:
		c.Mode = binary.Mode(d.Uint32())
	default:
		d.SetError(errors.Newf("Invalid control block version %d", version))
	}
}
