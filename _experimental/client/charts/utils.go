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

package charts

import "math"

func sum(vals []int) int {
	v := 0
	for _, x := range vals {
		v += x
	}
	return v
}

func round(v float64) int {
	if v < 0 {
		return int(v - 0.5)
	} else {
		return int(v + 0.4999999)
	}
}

func flooredPOT(v float64) int {
	log2 := math.Log2(v)
	if log2 >= 0 {
		return int(1 << uint(log2))
	} else {
		return 1
	}
}

func calculateStepInterval(multiple, valueRange, targetCount int) int {
	return multiple * flooredPOT(float64(valueRange)/float64(targetCount*multiple))
}

func quantize(valueRange, targetCount int, multiples ...int) int {
	best := 0
	bestErr := math.MaxInt32
	for _, m := range multiples {
		step := calculateStepInterval(m, valueRange, targetCount)
		err := int(valueRange) - int(step*targetCount)
		if err < bestErr {
			bestErr = err
			best = step
		}
	}
	return best
}
