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

// Package generate has support processing loaded go code, finding the items
// that require generated code, and converting them to a form the templates can
// easily consume.
package generate

import (
	"fmt"
	"strings"

	"golang.org/x/tools/go/exact"
	"golang.org/x/tools/go/types"

	"android.googlesource.com/platform/tools/gpu/binary/schema"
)

func (m *Module) addConst(c *types.Const) {
	t := fromType(m.Source.Types, c.Type(), "", m.Imports, nil)
	name := c.Name()
	directive := fmt.Sprintf("%s#%s", t, name)
	if d, found := m.Directives[directive]; found {
		name = d
	} else {
		name = strings.TrimPrefix(name, t.String())
		name = strings.Trim(name, "_")
	}
	if p, ok := t.(*schema.Primitive); ok {
		switch p.Method {
		case schema.Int8:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int8(v),
			})
		case schema.Uint8:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint8(v),
			})
		case schema.Int16:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int16(v),
			})
		case schema.Uint16:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint16(v),
			})
		case schema.Int32:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: int32(v),
			})
		case schema.Uint32:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: uint32(v),
			})
		case schema.Int64:
			v, _ := exact.Int64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: v,
			})
		case schema.Uint64:
			v, _ := exact.Uint64Val(c.Val())
			m.Constants.Add(t, schema.Constant{
				Name:  name,
				Value: v,
			})
		}
	}
}
