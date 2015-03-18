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
	"android.googlesource.com/platform/tools/gpu/database/store"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// build writes to out the Binary resource resulting from the given GetState request.
func (request *GetState) build(db database.Database, logger log.Logger, out binary.Object) error {
	atoms, _, err := getAtoms(request.Capture, db, logger)
	if err != nil {
		return err
	}
	if request.After >= atom.ID(len(atoms)) {
		return fmt.Errorf("After (%d) parameter is out of bounds. [0-%d]", request.After, len(atoms))
	}

	contextID := atoms[request.After].ContextID()
	api, err := getAPI(request.Capture, db, logger)
	if err != nil {
		return err
	}

	state := api.InitialState()
	stateMutator := api.StateMutator(state)
	for i, a := range atoms {
		if a.ContextID() == contextID {
			stateMutator.Write(atom.ID(i), a)
		}
		if atom.ID(i) == request.After {
			break
		}
	}

	buf := &bytes.Buffer{}
	enc := cyclic.Encoder(vle.Writer(buf))
	state.Encode(enc)

	store.CopyResource(out, &service.Binary{buf.Bytes()})
	return nil
}
