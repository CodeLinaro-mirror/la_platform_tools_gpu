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
	"bytes"
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/binary/vle"

	"sort"
)

type byID []*Struct

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].Signature() < a[j].Signature() }

func WriteAllSignatures(w io.Writer, modules Modules) {
	structs := []*Struct{}
	for _, m := range modules {
		structs = append(structs, m.Structs...)
	}
	sort.Sort(byID(structs))
	buf := &bytes.Buffer{}
	e := cyclic.Encoder(vle.Writer(buf))
	total := 0
	largest := 0
	// pre write the entire schema so the lookup table is full
	for _, s := range structs {
		e.Entity(&s.Entity, false)
	}
	all := buf.Len()
	for _, s := range structs {
		start := buf.Len()
		// now encode the entity directly to bypass the table
		schema.EncodeEntity(e, &s.Entity, true)
		size := buf.Len() - start
		total += size
		if largest < size {
			largest = size
		}
		fmt.Fprintln(w)
		fmt.Fprintln(w, s.Name(), ": size", size)
		fmt.Fprintln(w, s.TypeID)
		fmt.Fprintln(w, s.Entity.Signature())
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Schema stats:")
	fmt.Fprintln(w, "Count:", len(structs))
	fmt.Fprintln(w, "All:", all)
	fmt.Fprintln(w, "Total:", total)
	fmt.Fprintln(w, "Average:", total/len(structs))
	fmt.Fprintln(w, "Largest:", largest)
}
