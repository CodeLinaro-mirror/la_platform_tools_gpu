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

package atom

// Transformer is the interface that wraps the basic Transform method.
type Transformer interface {
	// Transform takes a given atom and identifier and Writes out a new atom and
	// identifier to the output atom Writer. Transform must not modify the atom in
	// any way.
	Transform(id ID, atom Atom, output Writer)
	// Flush is called at the end of an atom stream to cause Transformers that
	// cache atoms to send any they have stored into the output.
	Flush(output Writer)
}
