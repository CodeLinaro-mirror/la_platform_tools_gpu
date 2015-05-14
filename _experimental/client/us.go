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

package client

import (
	"fmt"

	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type μs int

func (μs μs) String() string {
	switch {
	case μs == 0:
		return "0"
	case μs > 1e6 || -μs > 1e6:
		return fmt.Sprintf("%.1fs", float64(μs)/1e6)
	case μs > 1e3 || -μs > 1e3:
		return fmt.Sprintf("%.1fms", float64(μs)/1e3)
	default:
		return fmt.Sprintf("%dus", μs)
	}
}

var thresholdColorDefault = gxui.Color{R: 0.00, G: 0.53, B: 0.74, A: 1}
var thresholdColorNear = gxui.Color{R: 0.9, G: 0.53, B: 0.2, A: 1}
var thresholdColorOver = gxui.Color{R: 1.0, G: 0.2, B: 0.2, A: 1}

const thresholdNearPercent = 90 // Considered 'near' at 90% threshold value

func (μs μs) ThesholdColor(target μs) gxui.Color {
	if μs > target {
		return thresholdColorOver
	} else {
		w := math.RampSat(float32(μs), float32(target*thresholdNearPercent)/100, float32(target))
		return gxui.Color{
			R: math.Lerpf(thresholdColorDefault.R, thresholdColorNear.R, w),
			G: math.Lerpf(thresholdColorDefault.G, thresholdColorNear.G, w),
			B: math.Lerpf(thresholdColorDefault.B, thresholdColorNear.B, w),
			A: 1,
		}
	}
}
