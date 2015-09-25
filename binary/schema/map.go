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
	Alias     string // The alias this array type was given, if present
	KeyType   Type   // The key type used.
	ValueType Type   // The value type stored in the map.
}

func (m *Map) Basename() string {
	return fmt.Sprintf("map[%s]%s", m.KeyType.Basename(), m.ValueType.Basename())
}

func (m *Map) Typename() string {
	if m.Alias != "" {
		return m.Alias
	}
	return fmt.Sprintf("map[%s]%s", m.KeyType.Typename(), m.ValueType.Typename())
}

func (m *Map) String() string {
	return m.Typename()
}

func (m *Map) Encode(e binary.Encoder, value interface{}) {
	v := value.(map[interface{}]interface{})
	e.Uint32(uint32(len(v)))
	for k, o := range v {
		m.KeyType.Encode(e, k)
		m.ValueType.Encode(e, o)
	}
}

func (m *Map) Decode(d binary.Decoder) interface{} {
	count := d.Uint32()
	v := make(map[interface{}]interface{}, count)
	for i := uint32(0); i < count; i++ {
		k := m.KeyType.Decode(d)
		v[k] = m.ValueType.Decode(d)
	}
	return v
}
