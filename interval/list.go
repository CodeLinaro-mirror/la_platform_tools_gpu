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

// package interval implements algorithms that operate on lists of intervals
package interval

// List is the interface to an object that can be used as an interval list by
// the algorithms in this package.
type List interface {
	// Length returns the number of elements in the list
	Length() int
	// GetSpan returns the span for the element at index in the list
	GetSpan(index int) U64Span
	// SetSpan sets the span for the element at index in the list
	SetSpan(index int, span U64Span)
	// Copy count list entries
	Copy(to, from, count int)
	// Resize adjusts the length of the array
	Resize(length int)
}

// Predicate is used as the condition for a Search
type Predicate func(test U64Span) bool

// Contains returns true if the value is found inside on of the intervals.
func Contains(l List, value uint64) bool {
	return findSpanFor(l, value) >= 0
}

// IndexOf returns the index of the span the value is a part of, or -1 if not found
func IndexOf(l List, value uint64) int {
	return findSpanFor(l, value)
}

// Merge adds a span to the list, merging it with existing spans if it overlaps
// them, and returns the index of that span.
func Merge(l List, span U64Span) int {
	return merge(l, span)
}

// Replace cuts the span out of any existing intervals, and then adds a new interval,
// and returns it's index.
func Replace(l List, span U64Span) int {
	index, newSpan := cut(l, span, true)
	l.SetSpan(index, newSpan)
	return index
}

// Remove strips the specified span from the list, cutting it out of any
// overlapping intervals
func Remove(l List, span U64Span) {
	cut(l, span, false)
}

// Intersect finds the intervals from the list that overlap with the specified span.
func Intersect(l List, span U64Span) (first, count int) {
	s := intersection{}
	s.intersect(l, span)
	return s.lowIndex, s.overlap
}

// Search finds the first interval in the list that the supplied predicate returns
// true for. If no interval matches the predicate, it returns the length of the list.
func Search(l List, t Predicate) int {
	return search(l, t)
}
