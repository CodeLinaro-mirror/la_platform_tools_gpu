// Copyright (C) 2014 The Android Open Source Project
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

package generate

import (
	"fmt"
	"io"

	"sort"
)

type byID []*Struct

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].ID().String() < a[j].ID().String() }

func WriteAllSignatures(w io.Writer, modules Modules) {
	structs := []*Struct{}
	for _, m := range modules {
		structs = append(structs, m.Structs...)
	}
	sort.Sort(byID(structs))
	for _, s := range structs {
		fmt.Fprintln(w)
		fmt.Fprintln(w, s.Name)
		fmt.Fprintln(w, s.ID())
		fmt.Fprintln(w, s.Signature())
	}
}
