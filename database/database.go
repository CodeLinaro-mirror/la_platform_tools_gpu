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

// Package database implements the persistence layer for the gpu debugger tools.
package database

import (
	"crypto/sha1"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Database is the interface to a resource store.
type Database interface {
	Store(binary.ID, binary.Object, log.Logger) error
	Resolve(binary.ID, log.Logger) (binary.Object, error)
	Contains(binary.ID, log.Logger) bool
}

// Store is a helper that stores an object to the database with the id
// calculated by the Hash function.
func Store(d Database, obj binary.Object, l log.Logger) (binary.ID, error) {
	id, err := Hash(obj)
	if err != nil {
		return id, err
	}
	return id, d.Store(id, obj, l)
}

// Hash returns a unique binary.ID based on the contents of the object.
// Two objects of identical content will return the same ID, and the
// probability of two objects with different content generating the same ID
// will be ignorable.
// Objects with a graph structure are allowed.
// Only members that would be encoded using a binary.Encoder are considered.
func Hash(o binary.Object) (binary.ID, error) {
	id := binary.ID{}
	h := sha1.New()
	e := cyclic.Encoder(vle.Writer(h))
	if err := e.Object(o); err != nil {
		return id, err
	}
	copy(id[:], h.Sum(nil))
	return id, nil
}
