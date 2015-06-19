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
	binary.Generate
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

func (m *Map) Encode(e binary.Encoder, value interface{}) error {
	v := value.(map[interface{}]interface{})
	if err := e.Uint32(uint32(len(v))); err != nil {
		return err
	}
	for k, o := range v {
		if err := m.KeyType.Encode(e, k); err != nil {
			return err
		}
		if err := m.ValueType.Encode(e, o); err != nil {
			return err
		}
	}
	return nil
}

func (m *Map) Decode(d binary.Decoder) (interface{}, error) {
	if count, err := d.Uint32(); err != nil {
		return nil, err
	} else {
		v := make(map[interface{}]interface{}, count)
		for i := uint32(0); i < count; i++ {
			k, err := m.KeyType.Decode(d)
			if err != nil {
				return v, err
			}
			o, err := m.ValueType.Decode(d)
			if err != nil {
				return v, err
			}
			v[k] = o
		}
		return v, nil
	}
}

func (m *Map) Skip(d binary.Decoder) error {
	if count, err := d.Uint32(); err != nil {
		return err
	} else {
		for i := uint32(0); i < count; i++ {
			if err := m.KeyType.Skip(d); err != nil {
				return err
			}
			if err := m.ValueType.Skip(d); err != nil {
				return err
			}
		}
	}
	return nil
}
