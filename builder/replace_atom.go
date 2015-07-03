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
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy returns a new *service.Capture, with a single atom replaced.
func (request *ReplaceAtom) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	original, err := service.ResolveCapture(request.Capture, d, l)
	if err != nil {
		return nil, err
	}

	atoms, err := loadAtoms(original.Atoms, d, l)
	if err != nil {
		return nil, err
	}

	if request.Atom >= atom.ID(len(atoms)) {
		return nil, fmt.Errorf("Atom (%d) parameter is out of bounds. [0-%d]", request.Atom, len(atoms))
	}

	o, err := decode(request.Data.Data)
	if err != nil {
		return nil, err
	}
	a, ok := o.(atom.Atom)
	if !ok {
		return nil, fmt.Errorf("Atom (%d) was not an atom, got %T.", request.Atom, o)
	}

	atoms = atoms.Clone()
	atoms[request.Atom] = a
	newStream, err := service.NewAtomStream(atoms)
	if err != nil {
		return nil, err
	}

	streamID, err := service.StoreAtomStream(&newStream, d, l)
	if err != nil {
		return nil, err
	}

	reportID, err := getBuildReport(streamID, d, l)
	if err != nil {
		return nil, err
	}

	capture := &service.Capture{
		Apis:   original.Apis,
		Name:   original.Name + "*",
		Atoms:  streamID,
		Report: reportID,
	}

	return capture, nil
}
