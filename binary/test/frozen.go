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

package test

import "android.googlesource.com/platform/tools/gpu/binary"

type X_V1 struct {
	binary.Frozen `name:"X"`
	a             int32
	b             int32
}

type X struct {
	binary.Generate `java:"disable"`
	a               int32
	b               int32
	c               string
}

func (before *X_V1) upgrade(after *X) {
	after.a = before.a
	after.b = before.b
	after.c = "Hello"
}

type Y struct {
	binary.Generate `java:"disable"`
	begin           string
	x               X
	end             string
}
