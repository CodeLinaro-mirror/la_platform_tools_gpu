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

package cyclic

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

type substack struct {
	stack []*binary.Entity
}

func (s *substack) Push(ent *binary.Entity) {
	s.stack = append(s.stack, ent)
}

func (s *substack) PushSubspace(ent *binary.Entity) {
	ents := ent.Subspace()
	for i := len(ents) - 1; i >= 0; i-- {
		subEntity := ents[i]
		s.Push(subEntity)
	}
}

func (s *substack) Pop() (*binary.Entity, error) {
	if len(s.stack) == 0 {
		return nil, fmt.Errorf("Pop on empty subtype Entity stack")
	}
	head := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return head, nil
}

func (s *substack) Dup(count uint32) error {
	top, err := s.Pop()
	if err != nil {
		return err
	}
	for i := uint32(0); i < count; i++ {
		s.Push(top)
	}
	return nil
}
