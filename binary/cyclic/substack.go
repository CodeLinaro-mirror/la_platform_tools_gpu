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
	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

type substack struct {
	stack []binary.Type
}

func (s *substack) pushStruct(ent *binary.Entity) {
	s.pushSubTypes(ent.Subspace().SubTypes)
}

func entityForStruct(t binary.Type) *binary.Entity {
	if s, ok := t.(*schema.Struct); ok {
		return s.Entity
	}
	return nil
}

func (s *substack) pushExpectStruct(t binary.Type) *binary.Entity {
	entity := entityForStruct(t)
	if entity == nil {
		return nil
	}
	s.pushStruct(entity)
	return entity
}

func (s *substack) pushSubTypes(t binary.TypeList) {
	for i := len(t) - 1; i >= 0; i-- {
		s.stack = append(s.stack, t[i])
	}
}

func (s *substack) pushCount(count uint32) error {
	t, err := s.popType()
	if err != nil {
		return err
	}
	sub := t.Subspace()
	if sub == nil || !sub.Counted {
		return fmt.Errorf(
			"Decoding counted collection, found non-counted type %s", t)
	}
	if len(sub.SubTypes) == 0 {
		return nil
	}
	for j := uint32(0); j < count; j++ {
		s.pushSubTypes(sub.SubTypes)
	}
	return nil
}

func (s *substack) popType() (binary.Type, error) {
	if len(s.stack) == 0 {
		return nil, fmt.Errorf("Pop on empty subtype Entity stack")
	}
	head := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return head, nil
}
