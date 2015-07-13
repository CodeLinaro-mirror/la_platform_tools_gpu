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

package task

import "testing"

type runnable struct {
	out chan<- int
}

func (r runnable) Run(c CancelSignal) {
	defer close(r.out)
	r.out <- 0
	c.Check()
	r.out <- 1
	c.Check()
	r.out <- 2
	c.Check()
	r.out <- 3
}

func TestRun(t *testing.T) {
	for i, test := range []struct {
		expected []int
		cancelAt int
	}{
		{
			expected: []int{0, 1, 0},
			cancelAt: 1,
		}, {
			expected: []int{0, 1, 2, 0},
			cancelAt: 2,
		}, {
			expected: []int{0, 1, 2, 3, 0},
			cancelAt: 3,
		}, {
			expected: []int{0, 1, 2, 3, 0},
			cancelAt: 99,
		},
	} {
		out := make(chan int)
		task := New()
		task.Run(runnable{out: out})
		for j, expected := range test.expected {
			if j == test.cancelAt {
				task.Cancel()
			}
			got := <-out
			if expected != got {
				t.Errorf("(%d.%d) Expected %d, got %d", i, j, expected, got)
			}
			if got == 0 {
				break
			}
		}
	}
}
