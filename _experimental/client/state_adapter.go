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

package client

import (
	"fmt"
	"reflect"
	"sort"
	"strconv"
	"strings"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const kStateAdapterNodeHeight = 18

func createControls(appCtx *ApplicationContext, name string, p path.Path, t binary.Type, v interface{}) gxui.Control {
	layout := appCtx.theme.CreateLinearLayout()
	layout.SetDirection(gxui.LeftToRight)

	addLabel := func(format string, args ...interface{}) gxui.Label {
		label := appCtx.theme.CreateLabel()
		label.SetText(fmt.Sprintf(format, args...))
		label.SetMargin(math.ZeroSpacing)
		label.SetMultiline(false)
		layout.AddChild(label)
		return label
	}

	label := addLabel("%s: ", name)
	appCtx.toolTipController.AddToolTip(label, 0.7, func(math.Point) gxui.Control {
		return createLabel(appCtx, stableStatePath(p), gxui.White)
	})

	if v != nil {
		if c := createField(appCtx, p, t, v); c != nil {
			layout.AddChild(c)
		}
	}

	return layout
}

type StateAdapterNode struct {
	appCtx   *ApplicationContext
	name     string
	ty       binary.Type
	value    interface{}
	path     path.Value
	item     string
	children StateAdapterNodeList
	parent   *StateAdapterNode
}

type StateAdapterNodeList []*StateAdapterNode

func (l StateAdapterNodeList) Len() int           { return len(l) }
func (l StateAdapterNodeList) Less(a, b int) bool { return l[a].name < l[b].name }
func (l StateAdapterNodeList) Swap(a, b int)      { l[a], l[b] = l[b], l[a] }

func (n *StateAdapterNode) add(name string, ty binary.Type, value interface{}, path path.Value) {
	n.children = append(n.children, &StateAdapterNode{
		appCtx: n.appCtx,
		name:   name,
		ty:     ty,
		value:  value,
		path:   path,
		item:   stableStatePath(path),
		parent: n,
	})
}

func stableStatePath(p path.Path) string {
	if s := path.FindState(p); s != nil {
		return strings.TrimPrefix(p.Path(), s.Path())
	}
	return p.Path()
}

func (n *StateAdapterNode) init() {
	n.children = nil
	if o, ok := n.value.(*schema.Object); ok {
		for i := range o.Fields {
			name := o.Type.Fields[i].Name()
			n.add(name, o.Type.Fields[i].Type, o.Fields[i], n.path.Field(name))
		}
	} else {
		name := func(ty binary.Type, v interface{}) string {
			if c := findConstant(findConstants(ty, n.appCtx), v); c.Value != nil {
				return c.Name
			}
			return fmt.Sprintf("%v", v)
		}

		switch ty := n.ty.(type) {
		case *schema.Array:
			v := reflect.ValueOf(n.value)
			for i, c := 0, v.Len(); i < c; i++ {
				v := v.Index(i)
				n.add(strconv.Itoa(i), ty.ValueType, v.Interface(), n.path.ArrayIndex(uint64(i)))
			}

		case *schema.Map:
			v := reflect.ValueOf(n.value)
			for _, k := range v.MapKeys() {
				v := v.MapIndex(k)
				k := k.Interface()
				n.add(name(ty.KeyType, k), ty.ValueType, v.Interface(), n.path.MapIndex(k))
			}
			sort.Sort(n.children)
		}
	}

	for _, c := range n.children {
		c.init()
	}
}

func (n *StateAdapterNode) Count() int {
	return len(n.children)
}

func (n *StateAdapterNode) NodeAt(index int) gxui.TreeNode {
	return n.children[index]
}

func (n *StateAdapterNode) ItemIndex(item gxui.AdapterItem) int {
	// Brute-force search
	for i, c := range n.children {
		if c.item == item {
			return i
		}
		if c.ItemIndex(item) >= 0 {
			return i
		}
	}
	return -1
}

func (n *StateAdapterNode) Item() gxui.AdapterItem {
	return n.item
}

func (n *StateAdapterNode) Create(t gxui.Theme) gxui.Control {
	if len(n.children) > 0 {
		return createControls(n.appCtx, n.name, n.path, n.ty, nil)
	} else {
		return createControls(n.appCtx, n.name, n.path, n.ty, n.value)
	}
}

type StateAdapter struct {
	StateAdapterNode
	gxui.AdapterBase
}

func (r *StateAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: kStateAdapterNodeHeight}
}

func NewStateAdapter(appCtx *ApplicationContext) *StateAdapter {
	return &StateAdapter{
		StateAdapterNode: StateAdapterNode{appCtx: appCtx},
	}
}

func (a *StateAdapter) Update(value interface{}, path *path.State) {
	v := value.(*schema.Object)
	a.value, a.path = v, path
	a.init()
	a.DataChanged(true)
}
