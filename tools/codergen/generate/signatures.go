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
func (a byID) Less(i, j int) bool { return a[i].ID().String() < a[j].ID().String() }

func WriteAllSignatures(w io.Writer, modules Modules) {
	structs := []*Struct{}
	for _, m := range modules {
		structs = append(structs, m.Structs...)
	}
	sort.Sort(byID(structs))
	one := &bytes.Buffer{}
	all := &bytes.Buffer{}
	allEnc := cyclic.Encoder(vle.Writer(all))
	total := 0
	largest := 0
	for _, s := range structs {
		one.Reset()
		e := cyclic.Encoder(vle.Writer(one))
		// pre-encode the field types, so we get an accurate measure
		for _, f := range s.Fields {
			schema.EncodeType(e, f.Type)
		}
		start := one.Len()
		// now encode the schema itself and measure the difference
		s.EncodeEntity(e)
		s.EncodeEntity(allEnc)
		size := one.Len() - start
		total += size
		if largest < size {
			largest = size
		}
		fmt.Fprintln(w)
		fmt.Fprintln(w, s.Name, ": size", size)
		fmt.Fprintln(w, s.ID())
		fmt.Fprintln(w, s.Signature())
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Schema stats:")
	fmt.Fprintln(w, "Count:", len(structs))
	fmt.Fprintln(w, "All:", all.Len())
	fmt.Fprintln(w, "Total:", total)
	fmt.Fprintln(w, "Average:", total/len(structs))
	fmt.Fprintln(w, "Largest:", largest)
}
