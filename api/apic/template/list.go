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

package template

import (
	"reflect"
	"sort"
)

// Join returns the concatenation of all the passed array or slices into a
// single, sequential list.
func (*Functions) Join(slices ...interface{}) []interface{} {
	out := []interface{}{}
	for _, slice := range slices {
		v := reflect.ValueOf(slice)
		for i, c := 0, v.Len(); i < c; i++ {
			out = append(out, v.Index(i).Interface())
		}
	}
	return out
}

// HasMore returns true if the i'th indexed item in l is not the last.
func (*Functions) HasMore(i int, l interface{}) bool {
	return i < reflect.ValueOf(l).Len()-1
}

// Reverse returns a new list with all the elements of in reversed.
func (*Functions) Reverse(in interface{}) interface{} {
	v := reflect.ValueOf(in)
	c := v.Len()
	out := reflect.MakeSlice(v.Type(), c, c)
	for i := 0; i < c; i++ {
		out.Index(i).Set(v.Index(c - i - 1))
	}
	return out.Interface()
}

// IndexOf returns the index of value in the list array. If value is not found
// in array then IndexOf returns -1.
func (*Functions) IndexOf(array interface{}, value interface{}) int {
	a := reflect.ValueOf(array)
	for i, c := 0, a.Len(); i < c; i++ {
		if a.Index(i).Interface() == value {
			return i
		}
	}
	return -1
}

type sortElem struct {
	key string
	idx int
}

type sortElemList []sortElem

func (l sortElemList) Len() int           { return len(l) }
func (l sortElemList) Less(a, b int) bool { return l[a].key < l[b].key }
func (l sortElemList) Swap(a, b int)      { l[a], l[b] = l[b], l[a] }

// SortBy returns a new list containing the sorted elements of arr. The list is
// sorted by the string values returned by calling the template function
// keyMacro with each element in the list.
func (f *Functions) SortBy(arr interface{}, keyMacro string) (interface{}, error) {
	v := reflect.ValueOf(arr)
	c := v.Len()

	elems := make(sortElemList, c)
	for i := 0; i < c; i++ {
		e := v.Index(i).Interface()
		key, err := f.Macro(keyMacro, e)
		if err != nil {
			return nil, err
		}
		elems[i] = sortElem{key, i}
	}

	sort.Sort(elems)

	out := reflect.MakeSlice(v.Type(), c, c)
	for i, e := range elems {
		out.Index(i).Set(v.Index(e.idx))
	}
	return out.Interface(), nil
}
