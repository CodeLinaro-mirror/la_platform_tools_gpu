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

package image

import (
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Thumbnailer is the interface implemented by types that can be visualized as
// a thumbnail image.
type Thumbnailer interface {
	// Thumbnail returns a thumbnail image info that most closely matches the
	// desired image width w and height h.
	Thumbnail(w, h uint32, d database.Database, l log.Logger) (*Info, error)
}
