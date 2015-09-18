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

package run

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

type dumper map[*graph.Step]struct{}

func (d dumper) dump(s *graph.Step, seen []*graph.Step) {
	if s == nil {
		fmt.Println()
		return
	}
	fmt.Printf(" [%d]", len(s.Inputs))
	if s.Disabled() {
		fmt.Println(" - disabled")
		return
	}
	if _, done := d[s]; done {
		fmt.Println(" - already seen")
		for i := range seen {
			if seen[i] == s {
				err := "Error: Cyclic dependency chain found:\n"
				for i := range seen {
					err += fmt.Sprintf("  [%d]: %v\n", i, seen[i])
				}
				panic(err)
			}
		}
		return
	}
	fmt.Println()
	d[s] = struct{}{}
	for i, e := range s.Inputs {
		for i := 0; i < len(seen); i++ {
			fmt.Print("  ")
		}
		fmt.Printf("(%d) %s", i+1, e)
		d.dump(graph.Creator(e), append(seen, s))
	}
}
