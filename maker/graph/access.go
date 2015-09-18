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
	"sort"
	"sync"
)

type lockEntry struct {
	name string
	sync.Mutex
}

var locks = map[string]*sync.Mutex{}

func addLock(name string) {
	_, ok := locks[name]
	if !ok {
		locks[name] = &sync.Mutex{}
	}
}

func lock(s *Step) {
	// sort the names for consistent acquire order
	sort.Strings(s.lockNames)
	// acquire all the mutexes in order
	for _, name := range s.lockNames {
		locks[name].Lock()
	}

}

func unlock(s *Step) {
	// unlock order does not matter
	for _, name := range s.lockNames {
		locks[name].Unlock()
	}
}
