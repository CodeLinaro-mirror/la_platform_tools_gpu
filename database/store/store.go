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

//go:generate codergen -go

// Package store implements the storage layers of the database system.
package store

import (
	"reflect"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Store is the interface to a database storage layer.
type Store interface {
	// Stores the resource r with the key id. d holds r as binary encoded.
	Store(id binary.ID, r binary.Object, d []byte, l log.Logger) error
	// Load puts the resource with key id into the out parameter
	Load(id binary.ID, l log.Logger, out binary.Object) (size int, err error)
	// Contains returns true if this store contains the resource identified by id.
	Contains(id binary.ID) bool
	// Close shuts down the store, it is an error to call any other method after this one.
	Close()
}

// CopyResource assigns the value object to the variable out points to
func CopyResource(out interface{}, value interface{}) {
	o := reflect.ValueOf(out).Elem()
	v := reflect.ValueOf(value).Elem()
	o.Set(v)
}
