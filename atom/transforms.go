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

// Transforms is a list of Transformer objects.
type Transforms []Transformer

// Transform sequentially transforms the atoms by each of the transformers in
// the list, before writing the final output to the output atom Writer.
func (l Transforms) Transform(atoms List, out Writer) {
	chain := out
	for i := len(l) - 1; i >= 0; i-- {
		chain = processWriter{l[i], chain}
	}
	atoms.WriteTo(chain)
	for p, ok := chain.(processWriter); ok; p, ok = chain.(processWriter) {
		chain = p.o
		p.t.Flush(chain)
	}
}

// Add is a convenience function for appending the list of Transformers t to the
// end of the Transforms list.
func (l *Transforms) Add(t ...Transformer) {
	*l = append(*l, t...)
}

// Transform is a helper for building simple Transformers that are implemented
// by function f. name is used to identify the transform when logging.
func Transform(name string, f func(id ID, atom Atom, output Writer)) Transformer {
	return transform{name, f}
}

type transform struct {
	N string                                // Transform name. Used for debugging.
	F func(id ID, atom Atom, output Writer) // The transform function.
}

func (t transform) Transform(id ID, atom Atom, output Writer) {
	t.F(id, atom, output)
}

func (t transform) Flush(output Writer) {}

type processWriter struct {
	t Transformer
	o Writer
}

func (p processWriter) Write(id ID, a Atom) {
	p.t.Transform(id, a, p.o)
}
