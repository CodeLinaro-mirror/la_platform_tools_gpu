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

package transform

import "android.googlesource.com/platform/tools/gpu/atom"

// ContextFilter returns a Transformer that only outputs atoms matching the
// specified context id.
func ContextFilter(context atom.ContextID) atom.Transformer {
	return atom.Transform("ContextFilter", func(id atom.ID, a atom.Atom, out atom.Writer) {
		if a.ContextID() == context {
			out.Write(id, a)
		}
	})
}
