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

package gles

import (
	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// instances returns the path to the Instances field of the currently bound
// context at p.
func instances(p path.Path, d database.Database, l log.Logger) (path.Value, error) {
	if atomp := path.FindAtom(p); atomp != nil {
		statei, err := database.Build(&builder.GetState{After: atomp}, d, l)
		if err != nil {
			return nil, err
		}
		state := statei.(*State)
		return atomp.StateAfter().
			Field("Contexts").
			MapIndex(int64(state.CurrentThread)).
			Field("Instances"), nil
	}
	return nil, nil
}

func (o AttributeLocation) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("VertexAttributeArrays").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o BufferId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("Buffers").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o FramebufferId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("Framebuffers").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o ProgramId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("Programs").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o QueryId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("Queries").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o RenderbufferId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("Renderbuffers").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o ShaderId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("Shaders").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o TextureId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("Textures").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}

func (o UniformLocation) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if atomp := path.FindAtom(p); atomp != nil {
		atomi, err := database.Build(&builder.Get{Path: atomp}, d, l)
		if err != nil {
			return nil, err
		}
		statei, err := database.Build(&builder.GetState{After: atomp}, d, l)
		if err != nil {
			return nil, err
		}
		state := statei.(*State)
		atom := atomi.(atom.Atom)
		var program ProgramId
		switch atom := atom.(type) {
		case *GlGetActiveUniform:
			program = atom.Program
		case *GlGetUniformLocation:
			program = atom.Program
		default:
			context := state.getContext()
			program = context.BoundProgram
		}
		return atomp.StateAfter().
			Field("Contexts").
			MapIndex(int64(state.CurrentThread)).
			Field("Instances").
			Field("Programs").
			MapIndex(int64(program)).
			Field("Uniforms").
			MapIndex(int64(o)), nil
	} else {
		return nil, nil
	}
}

func (o VertexArrayId) Link(p path.Path, d database.Database, l log.Logger) (path.Path, error) {
	if i, err := instances(p, d, l); err == nil {
		return i.Field("VertexArrays").MapIndex(int64(o)), nil
	} else {
		return nil, err
	}
}
