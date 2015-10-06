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

// substack is a stack of type objects. Only types which need decoder
// support for nested sub-structures are added to the stack. The
// substack is an implementation detail of the cyclic decoder.
type substack struct {
	stack []binary.Type
}

// pushStruct pushes the sub-types needed to decode a struct described
// the the schema object 'ent'.
func (s *substack) pushStruct(ent *binary.Entity) {
	s.pushSubTypes(ent.Subspace().SubTypes)
}

// entityForStruct if 't' is a schema object for a struct type return
// the schema entity for that struct.
func entityForStruct(t binary.Type) *binary.Entity {
	if s, ok := t.(*schema.Struct); ok {
		return s.Entity
	}
	return nil
}

// pushExpectStruct pushes the sub-types needed to decode a struct
// described by the type 't'. If 't' is not a struct type push nothing
// and return nil, otherwise return the schema entity for that struct.
func (s *substack) pushExpectStruct(t binary.Type) *binary.Entity {
	entity := entityForStruct(t)
	if entity == nil {
		return nil
	}
	s.pushStruct(entity)
	return entity
}

// pushSubTypes pushes the sub-types needed to decode a value of type 't'.
func (s *substack) pushSubTypes(t binary.TypeList) {
	for i := len(t) - 1; i >= 0; i-- {
		s.stack = append(s.stack, t[i])
	}
}

// pushCount, pops the top type from the stack and pushes any sub-types
// of that type count times. This is used to decode a collection of size
// count. If the top item on the stack is not a collection (slice, array, map)
// then an error is returned.
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

// popType pops the type which is on the top of the stack. An error
// is returned if the stack is empty.
func (s *substack) popType() (binary.Type, error) {
	if len(s.stack) == 0 {
		return nil, fmt.Errorf("Pop on empty subtype Entity stack")
	}
	head := s.stack[len(s.stack)-1]
	s.stack = s.stack[:len(s.stack)-1]
	return head, nil
}
