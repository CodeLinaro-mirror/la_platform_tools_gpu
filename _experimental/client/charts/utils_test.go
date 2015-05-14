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

import "testing"

func TestRound(t *testing.T) {
	check := []struct {
		value    float64
		expected int
	}{
		{-1.1, -1},
		{-0.9, -1},
		{-0.5, -1},
		{-0.1, 0},
		{0.1, 0},
		{0.5, 0},
		{0.9, 1},
		{1.1, 1},
	}

	for _, v := range check {
		got := round(v.value)
		if got != v.expected {
			t.Errorf("round(%v) returned unexpected value. Expected: %v, Got: %v", v.value, v.expected, got)
		}
	}
}

func TestFlooredPOT(t *testing.T) {
	check := []struct {
		value    float64
		expected int
	}{
		{0.5, 1},
		{0, 1},
		{1, 1},
		{2, 2},
		{3, 2},
		{4, 4},
		{5, 4},
	}

	for _, v := range check {
		got := flooredPOT(v.value)
		if got != v.expected {
			t.Errorf("flooredPOT(%v) returned unexpected value. Expected: %v, Got: %v", v.value, v.expected, got)
		}
	}
}

func TestCalculateStepInterval(t *testing.T) {
	check := []struct {
		base, rng, cnt int
		expected       int
	}{
		//  1  2  3  4  5  6  7  8  9  10 11 12 13 14 15 16 17 18
		//     2     4     6     8     10    12    14    16    18
		//           4           8           12          16
		//                       8                       16
		//                                               16
		{1, 18, 18, 1},
		{1, 18, 9, 2},
		{1, 18, 4, 4},
		{1, 18, 2, 8},
		{1, 18, 1, 16},

		//  3  6  9  12 15 18 21 24 27 30 33 36 39 42 45 48 51 54
		//     6     12    18    24    30    36    42    48    54
		//           12          24          36          48
		//                       24                      48
		//                                               48
		{3, 54, 18, 3},
		{3, 54, 9, 6},
		{3, 54, 4, 12},
		{3, 54, 2, 24},
		{3, 54, 1, 48},

		//  5  10 15 20 25 30 35 40 45 50 55 60 65 70 75 80 85 90
		//     10    20    30    40    50    60    70    80    90
		//           20          40          60          80
		//                       40                      80
		//                                               80
		{5, 90, 18, 5},
		{5, 90, 9, 10},
		{5, 90, 4, 20},
		{5, 90, 2, 40},
		{5, 90, 1, 80},
	}

	for _, v := range check {
		got := calculateStepInterval(v.base, v.rng, v.cnt)
		if got != v.expected {
			t.Errorf("calculateStepInterval(%v, %v, %v) returned unexpected value. Expected: %v, Got: %v", v.base, v.rng, v.cnt, v.expected, got)
		}
	}
}
