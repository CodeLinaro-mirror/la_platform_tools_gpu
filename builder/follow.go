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

package builder

import (
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

// BuildLazy resolves and returns the link on the object represented by the
// specified path.
func (r *Follow) BuildLazy(c interface{}, d database.Database, l log.Logger) (interface{}, error) {
	obj, err := database.Build(&Get{Path: r.Path}, d, l)
	if err != nil {
		return nil, err
	}
	if linker, ok := obj.(path.Linker); ok {
		return linker.Link(r.Path, d, l)
	}
	return nil, nil
}
