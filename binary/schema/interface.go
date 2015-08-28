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

package schema

import "android.googlesource.com/platform/tools/gpu/binary"

// Interface is the Type descriptor for a field who's underlying type is dynamic.
type Interface struct {
	binary.Generate
	Name string // The simple name of the type.
}

func (i *Interface) Basename() string {
	return i.Name
}

func (i *Interface) Typename() string {
	return i.Name
}

func (i *Interface) String() string {
	return i.Name
}

func (i *Interface) Encode(e binary.Encoder, value interface{}) error {
	if value != nil { // TODO proper nil test needed?
		if err := e.Object(value.(binary.Object)); err != nil {
			return err
		}
	} else if err := e.Object(nil); err != nil {
		return err
	}
	return nil
}

func (i *Interface) Decode(d binary.Decoder) (interface{}, error) {
	return d.Object()
}
