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

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// GetResources is a Lazy that builds a list of all the resources used by the
// specified capture.
type GetResources struct {
	binary.Generate
	Capture *path.Capture
}

type trackedResource struct {
	resource gfxapi.Resource
	id       path.ResourceID
	name     string
	accesses []uint64
}

func genResourceID(createdAt uint64, name string) path.ResourceID {
	return path.ResourceID(binary.NewID([]byte(fmt.Sprintf("%d %s", createdAt, name))))
}

// BuildLazy returns the *service.Resources resulting from the given
// GetResources request.
func (r *GetResources) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	atoms, err := ResolveAtoms(r.Capture.Atoms(), d, l)
	if err != nil {
		return nil, err
	}

	resources := []trackedResource{}
	seen := map[gfxapi.Resource]int{}

	var currentAtomIndex uint64

	state := gfxapi.NewState()
	state.OnResourceCreated = func(r gfxapi.Resource) {
		name := r.ResourceName()
		seen[r] = len(seen)
		resources = append(resources, trackedResource{
			resource: r,
			id:       genResourceID(currentAtomIndex, name),
			name:     name,
		})
	}
	state.OnResourceAccessed = func(r gfxapi.Resource) {
		if index, ok := seen[r]; ok { // Update the list of accesses
			c := len(resources[index].accesses)
			if c == 0 || resources[index].accesses[c-1] != currentAtomIndex {
				resources[index].accesses = append(resources[index].accesses, currentAtomIndex)
			}
		}
	}
	for i, a := range atoms {
		currentAtomIndex = uint64(i)
		a.Mutate(state, d, l)
	}

	out := &service.Resources{}
	for _, r := range resources {
		ty := r.resource.ResourceType()
		info := service.ResourceInfo{
			ID:       r.id,
			Name:     r.name,
			Accesses: r.accesses,
		}
		switch ty {
		case gfxapi.TypeUnknown:
			// We can't do anything with these objects.
		case gfxapi.TypeTexture1D:
			out.Textures1D = append(out.Textures1D, info)
		case gfxapi.TypeTexture2D:
			out.Textures2D = append(out.Textures2D, info)
		case gfxapi.TypeTexture3D:
			out.Textures3D = append(out.Textures3D, info)
		case gfxapi.TypeCubemap:
			out.Cubemaps = append(out.Cubemaps, info)
		default:
			panic(fmt.Errorf("Unknown resource type %v", ty))
		}
	}

	return out, nil
}

// GetResourceData is a Lazy that retrieves a resource's data by path.
type GetResourceData struct {
	binary.Generate
	Path *path.Resource
}

// BuildLazy returns the a requested Resource's data.
func (r *GetResourceData) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	atoms, err := ResolveAtoms(r.Path.After.Atoms, d, l)
	if err != nil {
		return nil, err
	}
	id := r.Path.ID
	var currentAtomIndex uint64
	var resource gfxapi.Resource
	state := gfxapi.NewState()
	state.OnResourceCreated = func(r gfxapi.Resource) {
		if genResourceID(currentAtomIndex, r.ResourceName()) == id {
			resource = r
			state.OnResourceCreated = nil // found the resource, no need to keep searching.
		}
	}
	for i, a := range atoms[:r.Path.After.Index+1] {
		currentAtomIndex = uint64(i)
		a.Mutate(state, d, l)
	}
	if resource != nil {
		return resource.ResourceData(state, d, l)
	}
	return nil, fmt.Errorf("Resource with id %v not found", r.Path.ID)
}
