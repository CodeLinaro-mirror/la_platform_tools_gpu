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

package binary

// Subspace defines a mapping from type Entitys to sub-types Entitys. This is
// used to decode composition hierarchies.

type EntityList []*Entity

// EncodeSubspace recusively encode entities needed in the definition of 'ent'
func EncodeSubspace(e Encoder, ent *Entity) {
	ents := ent.Subspace()
	for _, s := range ents {
		e.Entity(s, true) // TODO
		if e.Error() != nil {
			return
		}
	}
}

// DecodeSubspace recursively decodes sub-types Entities needed in the
// definition of 'ent'.
func DecodeSubspace(d Decoder, ent *Entity) error {
	subs := ent.Subspace()
	for i, _ := range subs {
		if subEntity := d.Entity(true); d.Error() != nil { // TODO
			return d.Error()
		} else {
			subs[i] = subEntity
		}
	}
	return nil
}
