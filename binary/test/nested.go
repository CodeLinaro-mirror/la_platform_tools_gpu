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

package test

import "android.googlesource.com/platform/tools/gpu/binary"

type Leaf struct {
	binary.Generate `java:"disable"`
	A               uint32
}

type Anonymous struct {
	binary.Generate `java:"disable"`
	Leaf
}

type Contains struct {
	binary.Generate `java:"disable"`
	LeafField       Leaf
}

type Array struct {
	binary.Generate `java:"disable"`
	Leaves          [3]Leaf
}

type Slice struct {
	binary.Generate `java:"disable"`
	Leaves          []Leaf
}

type MapKey struct {
	binary.Generate `java:"disable"`
	M               map[Leaf]uint32
}

type MapValue struct {
	binary.Generate `java:"disable"`
	M               map[uint32]Leaf
}

type MapKeyValue struct {
	binary.Generate `java:"disable"`
	M               map[Leaf]Leaf
}

type ArrayInMap struct {
	binary.Generate `java:"disable"`
	M               map[uint32][3]Leaf
}

type SliceInMap struct {
	binary.Generate `java:"disable"`
	M               map[uint32][]Leaf
}

type MapInSlice struct {
	binary.Generate `java:"disable"`
	Slice           []map[uint32]uint32
}

type MapInArray struct {
	binary.Generate `java:"disable"`
	Array           [2]map[uint32]uint32
}

type MapOfMaps struct {
	binary.Generate `java:"disable"`
	M               map[uint32]map[Leaf]Leaf
}

type ArrayOfArrays struct {
	binary.Generate `java:"disable"`
	Array           [2][3]Leaf
}

type SliceOfSlices struct {
	binary.Generate `java:"disable"`
	Slice           [][]Leaf
}

type Complex struct {
	binary.Generate `java:"disable"`
	SliceMapArray   []map[Contains][3]Contains
	SliceArrayMap   [][3]map[Contains]Contains
	ArraySliceMap   [3][]map[Contains]Contains
	ArrayMapSlice   [3]map[Contains][]Contains
	MapArraySlice   map[Contains][3][]Contains
	MapSliceArray   map[Contains][][3]Contains
}
