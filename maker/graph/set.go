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

package graph

import (
	"fmt"
	"path/filepath"
)

// Set manages an entity set.
type Set []Entity

// Append returns a new FileSet with files appended to fs.
// It does not attempt to suppress duplicates.
func (s Set) Append(entities ...Entity) Set {
	return append(s, entities...)
}

// Filter removes entries for which the predicate returns true. The order of entries is not changed.
func (s Set) Filter(predicate func(entry Entity) bool) Set {
	out := []Entity{}
	for _, entry := range s {
		if !predicate(entry) {
			out = append(out, entry)
		}
	}
	fmt.Println(out)
	return out
}

// Unique prunes the entity set down by removing duplicates. The order of entries remains stable.
func (s Set) Unique() Set {
	seen := map[string]struct{}{}
	return s.Filter(func(entry Entity) bool {
		name := entry.Name()
		_, found := seen[name]
		if !found {
			seen[name] = struct{}{}
		}
		return found
	})
}

func matches(name string, pattern string) bool {
	matched, err := filepath.Match(pattern, name)
	if err != nil {
		panic(err)
	}
	return matched
}

// Filter returns the list of entities in this Set that match any pattern in patterns.
// The pattern is a glob, see filepath.Match for details.
func (s Set) Include(patterns ...string) Set {
	return s.Filter(func(entry Entity) bool {
		name := filepath.Base(entry.Name())
		for _, pattern := range patterns {
			if matches(name, pattern) {
				return false
			}
		}
		return true
	})
}

// Exclude returns the list of entities in this Set that does not match any pattern in patterns.
// The pattern is a glob, see filepath.Match for details.
func (s Set) Exclude(patterns ...string) Set {
	return s.Filter(func(entry Entity) bool {
		name := filepath.Base(entry.Name())
		for _, pattern := range patterns {
			if matches(name, pattern) {
				return true
			}
		}
		return false
	})
}
