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

package builder

import (
	"bytes"
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// build writes to out the Binary resource resulting from the given GetState request.
func (request *GetState) build(d database.Database, l log.Logger, out binary.Object) error {
	capture, err := LoadCapture(request.Capture, d, l)
	if err != nil {
		return err
	}

	atoms, err := loadAtoms(capture.Atoms, d, l)
	if err != nil {
		return err
	}

	if request.After >= atom.ID(len(atoms)) {
		return fmt.Errorf("After (%d) parameter is out of bounds. [0-%d]", request.After, len(atoms))
	}

	s := gfxapi.NewState()
	for _, a := range atoms[:request.After] {
		if err := a.Mutate(s, d, l); err != nil {
			return err
		}
	}

	buf := &bytes.Buffer{}
	if err := cyclic.Encoder(vle.Writer(buf)).Value(s); err != nil {
		return err
	}

	database.CopyResource(out, &service.Binary{Data: buf.Bytes()})
	return nil
}
