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

package schema

import "android.googlesource.com/platform/tools/gpu/service"

// Map is a schema-typed Map value.
type Map struct {
	Type     *service.MapInfo // The map type info.
	Elements []MapElement     // The map elements.
}

// MapElement is a single key-value pair element held by a Map.
type MapElement struct {
	Key, Value interface{}
}
