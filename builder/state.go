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
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// BuildLazy returns the *service.Binary resulting from the given GetState request.
func (r *GetState) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	capture, err := service.ResolveCapture(r.Capture, d, l)
	if err != nil {
		return nil, err
	}

	api := gfxapi.Find(gfxapi.ID(r.API.ID))
	if api == nil {
		return nil, fmt.Errorf("Unknown graphics API '%v'", r.API.ID)
	}

	atoms, err := loadAtoms(capture.Atoms, d, l)
	if err != nil {
		return nil, err
	}

	if r.After >= atom.ID(len(atoms)) {
		return nil, fmt.Errorf("After (%d) parameter is out of bounds. [0-%d]", r.After, len(atoms))
	}

	s := gfxapi.NewState()
	for _, a := range atoms[:r.After] {
		a.Mutate(s, d, l)
	}

	res, found := s.APIs[api]
	if !found {
		return nil, fmt.Errorf("No state for API '%v'", api.Name())
	}

	return res, nil
}
