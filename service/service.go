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

// Package service is the definition of the RPC GPU debugger service exposed by the server.
//
// It is not the actual implementation of the service functionality.
package service

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// binary: java.source = adt/idea/android/src
// binary: java.package = com.android.tools.idea.editors.gfxtrace.rpc
// binary: java.indent = "  "
// binary: java.member_prefix = my
// binary: service = RPC

// Clone returns a new AtomStream holding a shallow-copy of the Atoms slice.
func (s *AtomStream) Clone() *AtomStream {
	atoms := make([]atom.Atom, len(s.Atoms))
	copy(atoms, s.Atoms)
	return &AtomStream{Atoms: atoms}
}

// ResolveBlob resolves a binary.ID safely casts the result to a []byte.
func ResolveBlob(id binary.ID, d database.Database, l log.Logger) ([]byte, error) {
	if v, err := database.Resolve(id, d, l); err != nil {
		return nil, err
	} else if r, ok := v.([]byte); !ok {
		return nil, fmt.Errorf("ID %s gave %T, expected []byte", id, v)
	} else {
		return r, nil
	}
}

// GetBlob calls s.Get with p and then safely casts the result to a []byte.
func GetBlob(p *path.Blob, s RPC, l log.Logger) ([]byte, error) {
	if v, err := s.Get(p, l); err != nil {
		return nil, err
	} else if r, ok := v.([]byte); !ok {
		return nil, fmt.Errorf("path %s gave %T, expected []byte", p, v)
	} else {
		return r, nil
	}
}
