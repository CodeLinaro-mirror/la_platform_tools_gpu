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
	"testing"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

func storeAtoms(t *testing.T, a *atom.List, d database.Database, l log.Logger) service.AtomsID {
	id, err := database.Store(a, d, l)
	if err != nil {
		t.Fatalf("Failed to store atom stream: %v", err)
	}
	return service.AtomsID(id)
}

// StoreCapture encodes and writes the atom list to the database, returning an
// identifier to the newly constructed and stored Capture.
func StoreCapture(t *testing.T, a *atom.List, d database.Database, l log.Logger) *path.Capture {
	capture := &service.Capture{
		Name:  "test-capture",
		Atoms: storeAtoms(t, a, d, l),
	}
	id, err := database.Store(capture, d, l)
	if err != nil {
		t.Fatalf("Failed to store test capture: %v", err)
	}
	return &path.Capture{ID: id}
}
