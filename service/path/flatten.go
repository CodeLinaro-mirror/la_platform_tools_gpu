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

// Flatten returns the path p flattened into a list of path nodes, starting with
// the root and ending with p.
func Flatten(p Path) []Path {
	l := []Path{}
	for p != nil {
		l = append(l, p)
		p = p.Base()
	}
	// l is backwards, reverse.
	for i, c, m := 0, len(l), len(l)/2; i < m; i++ {
		j := c - i - 1
		l[i], l[j] = l[j], l[i]
	}
	return l
}
