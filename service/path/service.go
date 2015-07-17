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

package path

import "android.googlesource.com/platform/tools/gpu/log"

type Service interface {
	// Get resolves and returns the object, value or memory at the path p.
	Get(p Path, l log.Logger) (interface{}, error)

	// Set creates a copy of the objects referenced by p, but with the object, value
	// or memory at p replaced with v. The path returned is identical to p, but with
	// the base changed to refer to the new object.
	Set(p Path, v interface{}, l log.Logger) (Path, error)

	// Follow returns the path to the object that the value at p links to.
	// If the value at p does not link to anything then nil is returned.
	Follow(p Path, l log.Logger) (Path, error)
}
