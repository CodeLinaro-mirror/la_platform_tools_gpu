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

type seenResource struct {
	slice *[]service.ResourceInfo
	index int
}

// GetResources is a Lazy that builds a list of all the resources used by the
// specified capture.
type GetResources struct {
	binary.Generate
	Capture *path.Capture
}

// BuildLazy returns the *service.Resources resulting from the given
// GetResources request.
func (r *GetResources) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	atoms, err := ResolveAtoms(r.Capture.Atoms(), d, l)
	if err != nil {
		return nil, err
	}

	seen := map[interface{}]seenResource{}
	textures := []service.ResourceInfo{}

	var idGen resourceIDGenerator
	state := gfxapi.NewState()
	state.OnResourceCreated = func(r gfxapi.Resource) {
		ty := r.ResourceType()
		info := service.ResourceInfo{
			ID:       idGen.gen(ty),
			Name:     r.ResourceName(),
			Accesses: []uint64{idGen.atomIndex},
		}
		switch ty {
		case gfxapi.TypeTexture:
			seen[r] = seenResource{slice: &textures, index: len(textures)}
			textures = append(textures, info)
		default:
			panic(fmt.Errorf("Unknown resource type %v", ty))
		}
	}
	state.OnResourceAccessed = func(r gfxapi.Resource) {
		s, ok := seen[r]
		if !ok {
			panic(fmt.Errorf("Resource %T %v was accessed at atom %d, but never created",
				r, r, idGen.atomIndex))
		}
		// Update the list of accesses
		info := &(*s.slice)[s.index]
		c := len(info.Accesses)
		if c == 0 || info.Accesses[c-1] != idGen.atomIndex {
			info.Accesses = append(info.Accesses, idGen.atomIndex)
		}
	}

	for i, a := range atoms {
		idGen.begin(uint64(i))
		a.Mutate(state, d, l)
	}

	return &service.Resources{Textures: textures}, nil
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
	var idGen resourceIDGenerator
	var resource gfxapi.Resource
	state := gfxapi.NewState()
	state.OnResourceCreated = func(r gfxapi.Resource) {
		if idGen.gen(r.ResourceType()) == id {
			resource = r
			state.OnResourceCreated = nil // found the resource, no need to keep searching.
		}
	}
	for i, a := range atoms[:r.Path.After.Index+1] {
		idGen.begin(uint64(i))
		a.Mutate(state, d, l)
	}
	if resource != nil {
		return resource.ResourceData(state, d, l)
	}
	return nil, fmt.Errorf("Resource with id %v not found", r.Path.ID)
}

type resourceIDGenerator struct {
	atomIndex uint64                      // current atom index
	resCount  map[gfxapi.ResourceType]int // reset each atom
}

func (g *resourceIDGenerator) begin(atomIndex uint64) {
	g.atomIndex = atomIndex
	g.resCount = nil
}

// gen calculates and returns a unique resource identifier for the resource with
// the given type.
func (g *resourceIDGenerator) gen(resType gfxapi.ResourceType) path.ResourceID {
	if g.resCount == nil {
		g.resCount = map[gfxapi.ResourceType]int{}
	}
	c := g.resCount[resType]
	g.resCount[resType] = c + 1
	id := binary.NewID([]byte(fmt.Sprintf("%d %d %d", g.atomIndex, resType, c)))
	return path.ResourceID(id)
}
