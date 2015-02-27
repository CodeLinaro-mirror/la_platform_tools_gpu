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

import "fmt"

var registry = map[TypeID]TypeInfo{}

// Register registers the atom type ty with the atom registry. If another atom
// is already registered with the same type identifer then Register will panic.
func Register(ty TypeInfo) {
	if _, dup := registry[ty.ID]; dup {
		panic(fmt.Errorf("Duplicate atom type id 0x%x registered", ty.ID))
	}
	registry[ty.ID] = ty
}

// New builds a new instance of the atom with type identifier id. The type must
// have previously been registered with Register.
func New(id TypeID) (Atom, error) {
	if ty, ok := registry[id]; ok {
		return ty.New(), nil
	} else {
		return nil, fmt.Errorf("Atom type id 0x%x not registered", id)
	}
}
