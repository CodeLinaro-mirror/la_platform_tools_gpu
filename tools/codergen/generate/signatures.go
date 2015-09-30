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

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/binary/vle"

	"sort"
)

type byID []*Struct

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].TypeID.String() < a[j].TypeID.String() }

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
		fmt.Fprintln(w, Signature(&s.Entity))
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Schema stats:")
	fmt.Fprintln(w, "Count:", len(structs))
	fmt.Fprintln(w, "All:", all)
	fmt.Fprintln(w, "Total:", total)
	fmt.Fprintln(w, "Average:", total/len(structs))
	fmt.Fprintln(w, "Largest:", largest)
}

func Signature(e *binary.Entity) string {
	b := &bytes.Buffer{}
	fmt.Fprint(b, e.Package, ".", e.Identity)
	if e.Version != "" {
		fmt.Fprint(b, "@", e.Version)
	}
	fmt.Fprint(b, "{")
	for i, f := range e.Fields {
		if i != 0 {
			fmt.Fprint(b, ",")
		}
		printTag(b, f.Type)
	}
	fmt.Fprint(b, "}")
	return b.String()
}

func printTag(w io.Writer, t binary.Type) {
	switch t := t.(type) {
	case *schema.Primitive:
		fmt.Fprint(w, t.Method)
	case *schema.Struct:
		fmt.Fprint(w, "$")
	case *schema.Pointer:
		fmt.Fprint(w, "*")
		printTag(w, t.Type)
	case *schema.Interface:
		fmt.Fprint(w, "?", t)
	case *schema.Variant:
		fmt.Fprint(w, "&", t)
	case *schema.Any:
		fmt.Fprint(w, "~", t)
	case *schema.Slice:
		fmt.Fprint(w, "[]")
		printTag(w, t.ValueType)
	case *schema.Array:
		fmt.Fprint(w, "[", t.Size, "]")
		printTag(w, t.ValueType)
	case *schema.Map:
		fmt.Fprint(w, "map[")
		printTag(w, t.KeyType)
		fmt.Fprint(w, "]")
		printTag(w, t.ValueType)
	default:
		panic(fmt.Errorf("Unknown type %T generating signature", t))
	}
}
