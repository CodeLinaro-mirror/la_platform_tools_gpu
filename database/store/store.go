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

// Package store implements the storage layers of the database system.
package store

import (
	"reflect"
)

// CopyResource assigns the value object to the variable out points to
func CopyResource(out interface{}, value interface{}) {
	o := reflect.ValueOf(out).Elem()
	v := reflect.ValueOf(value).Elem()
	o.Set(v)
}
