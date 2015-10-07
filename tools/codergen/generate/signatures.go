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

	"android.googlesource.com/platform/tools/gpu/binary"
)

type byID []*Struct

func (a byID) Len() int           { return len(a) }
func (a byID) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }
func (a byID) Less(i, j int) bool { return a[i].Signature() < a[j].Signature() }

type stream struct {
	name    string
	b       bytes.Buffer
	e       binary.Encoder
	d       binary.Decoder
	compact bool
	size    int
}

func newStream(compact bool) *stream {
	s := &stream{compact: compact}
	if compact {
		s.name = "compact"
	} else {
		s.name = "full"
	}
	s.e = cyclic.Encoder(vle.Writer(&s.b))
	s.d = cyclic.Decoder(vle.Reader(&s.b))
	return s
}

func (s *stream) test(e *binary.Entity) {
	s.e.Entity(e, s.compact)
	s.size += s.b.Len()
	if s.e.Error() != nil {
		panic(fmt.Errorf("Failed encoding %s entity for %q, %v", s.name, e.Signature(), s.e.Error()))
	}
	got := s.d.Entity(s.compact)
	if got == nil || s.d.Error() != nil {
		panic(fmt.Errorf("Failed reading %s entity for %q, %v", s.name, e.Signature(), s.d.Error()))
	}
	if e.Signature() != got.Signature() {
		panic(fmt.Errorf("Signature of %s entity did not match, expected %q got %q", s.name, e.Signature(), got.Signature()))
	}
	if !s.compact {
		es := fmt.Sprint(e)
		gots := fmt.Sprint(got)
		if es != gots {
			panic(fmt.Errorf("Full encoding did not match, expected %#v got %#v", es, gots))
		}
	}
}

func WriteAllSignatures(w io.Writer, modules Modules) {
	structs := []*Struct{}
	for _, m := range modules {
		structs = append(structs, m.Structs...)
	}
	sort.Sort(byID(structs))
	full := newStream(false)
	compact := newStream(true)
	total := 0
	largest := 0
	// pre write the entire schema so the lookup table is full, and verify the encode/decode behaviour while doing it
	for _, s := range structs {
		full.test(&s.Entity)
		compact.test(&s.Entity)
	}
	for _, s := range structs {
		start := compact.b.Len()
		// now encode the entity directly to bypass the table
		schema.EncodeEntity(compact.e, &s.Entity, true)
		size := compact.b.Len() - start + 2
		total += size
		if largest < size {
			largest = size
		}
		fmt.Fprintln(w)
		fmt.Fprintln(w, s.Name(), ": size", size)
		fmt.Fprintln(w, s.Entity.Signature())
	}
	fmt.Fprintln(w)
	fmt.Fprintln(w, "Schema stats:")
	fmt.Fprintln(w, "Count:", len(structs))
	fmt.Fprintln(w, "Total:", total)
	fmt.Fprintln(w, "Compact:", compact.size)
	fmt.Fprintln(w, "Full:", full.size)
	fmt.Fprintln(w, "Average:", total/len(structs))
	fmt.Fprintln(w, "Largest:", largest)
}
