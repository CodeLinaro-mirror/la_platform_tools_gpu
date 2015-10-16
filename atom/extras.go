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

// Package atom provides the fundamental types used to describe a capture stream.
package atom

import "android.googlesource.com/platform/tools/gpu/binary"

// Extra is the interface implemented by atom 'extras' - additional information
// that can be placed inside an atom instance.
type Extra binary.Object

// ExtraCast is automatically called by the generated decoders.
func ExtraCast(obj binary.Object) Extra { return obj }

// Extras is a list of Extra objects.
type Extras []Extra

// Ohservations returns a pointer to the Observations structure in the extras,
// or nil if there are no observations in the extras.
func (e Extras) Observations() *Observations {
	for _, o := range e {
		if o, ok := o.(*Observations); ok {
			return o
		}
	}
	return nil
}

// GetOrAppendObservations returns a pointer to the existing Observations
// structure in the extras, or appends and returns a pointer to a new
// observations structure if the extras does not already contain one.
func (e *Extras) GetOrAppendObservations() *Observations {
	if o := e.Observations(); o != nil {
		return o
	}
	o := &Observations{}
	*e = append(*e, o)
	return o
}
