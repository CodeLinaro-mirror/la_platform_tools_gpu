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

package utils

import (
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

func storeSchema(t *testing.T, api gfxapi.API, db database.Database, l log.Logger) service.SchemaId {
	schema := api.Schema()
	id, err := db.Store(&schema, l)
	if err != nil {
		t.Fatalf("Failed to store schema: %v", err)
	}
	return service.SchemaId{id}
}

func storeAtoms(t *testing.T, atoms atom.List, api gfxapi.API, db database.Database, l log.Logger) service.AtomStreamId {
	stream, err := service.NewAtomStream(atoms, storeSchema(t, api, db, l))
	if err != nil {
		t.Fatalf("Failed to build atom stream: %v", err)
	}
	id, err := db.Store(&stream, l)
	if err != nil {
		t.Fatalf("Failed to store atom stream: %v", err)
	}
	return service.AtomStreamId{id}
}

// StoreCapture encodes and writes the atom list to the database, returning an
// identifier to the newly constructed and stored Capture.
func StoreCapture(t *testing.T, atoms atom.List, api gfxapi.API, db database.Database, l log.Logger) service.CaptureId {
	capture := service.Capture{
		Name:  "test-capture",
		API:   api.Name(),
		Atoms: storeAtoms(t, atoms, api, db, l),
	}
	id, err := db.Store(&capture, l)
	if err != nil {
		t.Fatalf("Failed to store test capture: %v", err)
	}
	return service.CaptureId{id}
}
