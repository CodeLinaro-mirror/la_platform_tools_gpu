// Copyright (C) 2014 The Android Open Source Project
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

package schema

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Map is the Type descriptor for key/value stores.
type Map struct {
	Alias     string      // The alias this array type was given, if present
	KeyType   binary.Type // The key type used.
	ValueType binary.Type // The value type stored in the map.
}

func (m *Map) Representation() string {
	return fmt.Sprintf("%r", m)
}

func (m *Map) String() string {
	return fmt.Sprint(m)
}

func (m *Map) Format(f fmt.State, c rune) {
	switch c {
	case 'z':
		fmt.Fprintf(f, "map[%z]%z", m.KeyType, m.ValueType)
	case 'r':
		fmt.Fprintf(f, "map[%r]%r", m.KeyType, m.ValueType)
	default:
		if m.Alias != "" {
			fmt.Fprint(f, m.Alias)
		} else {
			fmt.Fprintf(f, "map[%v]%v", m.KeyType, m.ValueType)
		}
	}
}

func (m *Map) EncodeValue(e binary.Encoder, value interface{}) {
	v := value.(map[interface{}]interface{})
	e.Uint32(uint32(len(v)))
	for k, o := range v {
		m.KeyType.EncodeValue(e, k)
		m.ValueType.EncodeValue(e, o)
	}
}

func (m *Map) DecodeValue(d binary.Decoder) interface{} {
	count := d.Uint32()
	v := make(map[interface{}]interface{}, count)
	for i := uint32(0); i < count; i++ {
		k := m.KeyType.DecodeValue(d)
		v[k] = m.ValueType.DecodeValue(d)
	}
	return v
}
