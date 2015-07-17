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
	"testing"

	"android.googlesource.com/platform/tools/gpu/database"
)

func TestLazyInterfaceCompliance(t *testing.T) {
	// Interface compliance tests
	_ = []database.Lazy{
		(*BuildReport)(nil),
		(*ConvertImage)(nil),
		(*Get)(nil),
		(*GetFramebufferColor)(nil),
		(*GetFramebufferDepth)(nil),
		(*GetHierarchy)(nil),
		(*GetState)(nil),
		(*GetTimingInfo)(nil),
		(*PrerenderFramebuffers)(nil),
		(*RenderFramebufferColor)(nil),
		(*RenderFramebufferDepth)(nil),
		(*Set)(nil),

		(*getCaptureFramebufferDimensions)(nil),
	}
}
