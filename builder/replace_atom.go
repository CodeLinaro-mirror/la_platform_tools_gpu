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

// build writes to out the Capture resource resulting from the given ReplaceAtom request.
func (request *ReplaceAtom) build(db database.Database, logger log.Logger, out binary.Object) error {
	atoms, schemaId, err := getAtoms(request.Capture, db, logger)
	if err != nil {
		return err
	}
	if request.Atom >= atom.ID(len(atoms)) {
		return fmt.Errorf("Atom (%d) parameter is out of bounds. [0-%d]", request.Atom, len(atoms))
	}

	atom, err := atom.New(request.Type)
	if err != nil {
		return err
	}

	if err := atom.Decode(cyclic.Decoder(vle.Reader(bytes.NewBuffer(request.Data.Data)))); err != nil {
		return err
	}

	atoms = atoms.Clone()
	atoms[request.Atom] = atom
	newStream, err := service.NewAtomStream(atoms, schemaId)
	if err != nil {
		return err
	}

	newStreamId, err := db.Store(&newStream, logger)
	if err != nil {
		return err
	}

	var capture service.Capture
	if err := db.Load(request.Capture.ID, logger, &capture); err != nil {
		return err
	}
	capture.Atoms = service.AtomStreamId{newStreamId}

	store.CopyResource(out, &capture)
	return nil
}
