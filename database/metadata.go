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

package database

import "android.googlesource.com/platform/tools/gpu/binary"

type metaType int

const (
	// The resource has not been built, but can be by using the request data.
	metaTypeLazy metaType = iota

	// This resource id is a link to the resource with the LinkTo identifier.
	metaTypeLink

	// This resource id points directly to data.
	metaTypeData

	// This resource id is data for a derived object.
	metaTypeDerived
)

func (t metaType) String() string {
	switch t {
	case metaTypeLazy:
		return "Lazy"
	case metaTypeLink:
		return "Link"
	case metaTypeData:
		return "Data"
	case metaTypeDerived:
		return "Derived"
	default:
		return "Unknown"
	}
}

type metadata struct {
	binary.Generate
	Type    metaType
	LinkTo  binary.ID
	Request binary.Object
}
