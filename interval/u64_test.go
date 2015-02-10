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

package interval

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func str(l U64SpanList) string {
	s := make([]string, len(l))
	for i, v := range l {
		s[i] = fmt.Sprintf("%d:%d", v.Start, v.End)
	}
	return "[" + strings.Join(s, ",") + "]"
}

func assertEqual(t *testing.T, n int, got U64SpanList, expected U64SpanList) {
	if len(got) != len(expected) {
		t.Fatalf("[%d] wrong length, expected %s got %s", n, str(expected), str(got))
	}
	for i, g := range got {
		if g != expected[i] {
			t.Fatalf("[%d] wrong at %d, expected %s got %s", n, i, str(expected), str(got))
		}
	}
}

func TestMerge(t *testing.T) {
	for n, test := range []struct {
		list     U64SpanList
		with     U64Span
		expected U64SpanList
	}{
		{ // Empty
			U64SpanList{},
			U64Span{0, 0},
			U64SpanList{U64Span{0, 0}},
		},
		{ // Duplicate
			U64SpanList{U64Span{10, 10}},
			U64Span{10, 10},
			U64SpanList{U64Span{10, 10}},
		},
		{ // Zero length duplicate
			U64SpanList{U64Span{10, 0}},
			U64Span{10, 0},
			U64SpanList{U64Span{10, 0}},
		},
		{ // between
			U64SpanList{U64Span{0, 10}, U64Span{40, 50}},
			U64Span{20, 30},
			U64SpanList{U64Span{0, 10}, U64Span{20, 30}, U64Span{40, 50}},
		},
		{ // before
			U64SpanList{U64Span{10, 20}},
			U64Span{0, 5},
			U64SpanList{U64Span{0, 5}, U64Span{10, 20}},
		},
		{ // after
			U64SpanList{U64Span{0, 5}},
			U64Span{10, 20},
			U64SpanList{U64Span{0, 5}, U64Span{10, 20}},
		},
		{ //extend before
			U64SpanList{U64Span{3, 5}},
			U64Span{0, 4},
			U64SpanList{U64Span{0, 5}},
		},
		{ //extend after
			U64SpanList{U64Span{3, 5}},
			U64Span{4, 7},
			U64SpanList{U64Span{3, 7}},
		},
		{ // extend middle before
			U64SpanList{U64Span{0, 2}, U64Span{4, 6}, U64Span{8, 10}},
			U64Span{3, 5},
			U64SpanList{U64Span{0, 2}, U64Span{3, 6}, U64Span{8, 10}},
		},
		{ // extend middle after
			U64SpanList{U64Span{0, 2}, U64Span{4, 6}, U64Span{8, 10}},
			U64Span{5, 7},
			U64SpanList{U64Span{0, 2}, U64Span{4, 7}, U64Span{8, 10}},
		},
		{ // inside start
			U64SpanList{U64Span{10, 20}},
			U64Span{10, 11},
			U64SpanList{U64Span{10, 20}},
		},
		{ // inside end
			U64SpanList{U64Span{10, 20}},
			U64Span{19, 20},
			U64SpanList{U64Span{10, 20}},
		},
		{ // merge first two
			U64SpanList{U64Span{0, 10}, U64Span{20, 30}, U64Span{40, 50}},
			U64Span{5, 25},
			U64SpanList{U64Span{0, 30}, U64Span{40, 50}},
		},
		{ // merge last two
			U64SpanList{U64Span{0, 10}, U64Span{20, 30}, U64Span{40, 50}},
			U64Span{25, 45},
			U64SpanList{U64Span{0, 10}, U64Span{20, 50}},
		},
		{ // merge overlap
			U64SpanList{U64Span{0, 10}, U64Span{20, 30}, U64Span{40, 50}},
			U64Span{5, 45},
			U64SpanList{U64Span{0, 50}},
		},
		{ // merge encompass
			U64SpanList{U64Span{5, 10}, U64Span{20, 30}, U64Span{40, 45}},
			U64Span{0, 50},
			U64SpanList{U64Span{0, 50}},
		},
	} {
		Merge(&test.list, test.with)
		assertEqual(t, n, test.list, test.expected)
	}
}

func TestReplace(t *testing.T) {
	for n, test := range []struct {
		list     U64SpanList
		with     U64Span
		expected U64SpanList
	}{
		{ // Empty
			U64SpanList{},
			U64Span{0, 10},
			U64SpanList{U64Span{0, 10}},
		},
		{ // match
			U64SpanList{U64Span{5, 10}},
			U64Span{5, 10},
			U64SpanList{U64Span{5, 10}},
		},
		{ // split
			U64SpanList{U64Span{5, 25}},
			U64Span{15, 20},
			U64SpanList{U64Span{5, 15}, U64Span{15, 20}, U64Span{20, 25}},
		},
		{ // between
			U64SpanList{U64Span{5, 10}, U64Span{25, 30}},
			U64Span{15, 20},
			U64SpanList{U64Span{5, 10}, U64Span{15, 20}, U64Span{25, 30}},
		},
	} {
		Replace(&test.list, test.with)
		assertEqual(t, n, test.list, test.expected)
	}
}

func TestRemove(t *testing.T) {
	for n, test := range []struct {
		list     U64SpanList
		with     U64Span
		expected U64SpanList
	}{
		{ // empty
			U64SpanList{},
			U64Span{0, 0},
			U64SpanList{},
		},
		{ // match
			U64SpanList{U64Span{3, 5}},
			U64Span{3, 5},
			U64SpanList{},
		},
		{ // before
			U64SpanList{U64Span{3, 5}},
			U64Span{0, 3},
			U64SpanList{U64Span{3, 5}},
		},
		{ // after
			U64SpanList{U64Span{3, 6}},
			U64Span{6, 7},
			U64SpanList{U64Span{3, 6}},
		},
		{ // trim front
			U64SpanList{U64Span{10, 20}},
			U64Span{5, 15},
			U64SpanList{U64Span{15, 20}},
		},
		{ // split
			U64SpanList{U64Span{10, 20}, U64Span{30, 40}, U64Span{50, 60}},
			U64Span{32, 38},
			U64SpanList{U64Span{10, 20}, U64Span{30, 32}, U64Span{38, 40}, U64Span{50, 60}},
		},
		{ // trim 2
			U64SpanList{U64Span{10, 20}, U64Span{30, 40}, U64Span{50, 60}},
			U64Span{35, 55},
			U64SpanList{U64Span{10, 20}, U64Span{30, 35}, U64Span{55, 60}},
		},
		{ //
			U64SpanList{U64Span{10, 20}, U64Span{30, 40}, U64Span{50, 60}},
			U64Span{0, 100},
			U64SpanList{},
		},
	} {
		Remove(&test.list, test.with)
		assertEqual(t, n, test.list, test.expected)
	}
}

func TestIntersect(t *testing.T) {
	for n, test := range []struct {
		list   U64SpanList
		with   U64Span
		offset int
		count  int
	}{
		{ // empty
			U64SpanList{},
			U64Span{0, 10},
			0, 0,
		},
		{ // match
			U64SpanList{U64Span{10, 20}},
			U64Span{10, 20},
			0, 1,
		},
		{ // start
			U64SpanList{U64Span{10, 20}},
			U64Span{10, 15},
			0, 1,
		},
		{ // end
			U64SpanList{U64Span{10, 20}},
			U64Span{15, 20},
			0, 1,
		},
		{ // middle
			U64SpanList{U64Span{10, 20}},
			U64Span{12, 18},
			0, 1,
		},
		{ // overlap 3
			U64SpanList{U64Span{10, 20}, U64Span{30, 40}, U64Span{50, 60}},
			U64Span{15, 55},
			0, 3,
		},
		{ // first 2
			U64SpanList{U64Span{10, 20}, U64Span{30, 40}, U64Span{50, 60}},
			U64Span{10, 35},
			0, 2,
		},
		{ // last 2
			U64SpanList{U64Span{10, 20}, U64Span{30, 40}, U64Span{50, 60}},
			U64Span{35, 60},
			1, 2,
		},
	} {
		o, c := Intersect(&test.list, test.with)
		if o != test.offset {
			t.Fatalf("[%d] wrong offset, expected %s got %s", n, test.offset, o)
		}
		if c != test.count {
			t.Fatalf("[%d] wrong count, expected %s got %s", n, test.count, c)
		}
	}
}

func TestU64SpanListIndexOf(t *testing.T) {
	l := U64SpanList{U64Span{10, 20}, U64Span{30, 40}, U64Span{50, 60}}
	for n, test := range []struct {
		value uint64
		index int
	}{
		{0, -1},
		{9, -1},
		{10, 0},
		{15, 0},
		{19, 0},
		{20, -1},
		{32, 1},
		{59, 2},
	} {
		got := IndexOf(&l, test.value)
		if got != test.index {
			t.Fatalf("[%d] wrong index, expected %d got %d looking for %d in %s", n, test.index, got, test.value, str(l))
		}
	}
}

const maxIntervalValue = 100000
const maxIntervalRange = 10000

type iteration struct{ merge, replace U64Span }

func buildRands(b *testing.B) []iteration {
	b.StopTimer()
	defer b.StartTimer()
	iterations := make([]iteration, b.N)
	rand.Seed(1)
	for i := range iterations {
		iterations[i].merge = U64Range{
			uint64(rand.Intn(maxIntervalValue)),
			uint64(rand.Intn(maxIntervalValue))}.Span()
		iterations[i].replace = U64Range{
			uint64(rand.Intn(maxIntervalValue)),
			uint64(rand.Intn(maxIntervalValue))}.Span()
	}
	return iterations
}

func BenchmarkGeneral(b *testing.B) {
	iterations := buildRands(b)
	l := U64SpanList{}
	for _, iter := range iterations {
		Merge(&l, iter.merge)
		Replace(&l, iter.replace)
	}
}
