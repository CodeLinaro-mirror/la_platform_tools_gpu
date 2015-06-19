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
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
)

// BuildLazy returns the *database.Blob holding the converted image for the
// ConvertImage request.
func (r *ConvertImage) BuildLazy(c interface{}, d database.Database, l log.Logger) (binary.Object, error) {
	data, err := database.ResolveBlob(r.Data, d, l)
	if err != nil {
		return nil, err
	}

	data, err = image.Convert(data, r.Width, r.Height, r.FormatFrom, r.FormatTo)
	if err != nil {
		return nil, err
	}

	return &database.Blob{Data: data}, nil
}
