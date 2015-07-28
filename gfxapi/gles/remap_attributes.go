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

import "android.googlesource.com/platform/tools/gpu/atom"

// remapAttributes returns an atom transform that remaps attribute locations by replacing glGetAttribLocation
// calls with glBindAttribLocation, then lazily injects glLinkProgram before glUseProgram when necessary.
func remapAttributes() atom.Transformer {
	pendingLink := map[ProgramId]struct{}{}

	return atom.Transform("RemapAttributes", func(i atom.ID, a atom.Atom, out atom.Writer) {
		switch a := a.(type) {
		case *GlGetAttribLocation:
			if a.Result != -1 {
				pendingLink[a.Program] = struct{}{} // Mark program as requiring linking before use.
				out.Write(i, NewGlBindAttribLocation(a.Program, AttributeLocation(a.Result), a.Name))
			}

		case *GlLinkProgram:
			delete(pendingLink, a.Program)
			out.Write(i, a)

		case *GlUseProgram:
			if _, pending := pendingLink[a.Program]; pending {
				out.Write(i, NewGlLinkProgram(a.Program))
				delete(pendingLink, a.Program)
			}
			out.Write(i, a)

		default:
			out.Write(i, a)
		}
	})
}
