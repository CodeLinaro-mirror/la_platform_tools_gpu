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
	"sort"

	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/image"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/task"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

const (
	resourceThumbnailWidth  = 64
	resourceThumbnailHeight = 64
)

func CreateResourcesPanel(appCtx *ApplicationContext) gxui.Control {
	adapter := &resourcesAdapter{}
	tree := appCtx.theme.CreateTree()
	tree.SetAdapter(adapter)

	atomIndex := ^uint64(0)

	var resources *path.Resources

	t := task.New()
	update := func(r *path.Resources) {
		if !path.Equal(resources, r) {
			resources = r
			context := &resourceAdapterContext{
				appCtx:    appCtx,
				atomIndex: &atomIndex,
				resources: resources,
			}
			t.Run(updateResourcesAdapter{context, adapter})
		}
	}

	appCtx.events.OnSelect(func(p path.Path) {
		if c := path.FindCapture(p); c != nil {
			update(c.Resources())
		}
		if a := lastAtom(p); a != nil {
			atomIndex = a.Index
			adapter.DataChanged(true)
		}
	})

	return tree
}

type updateResourcesAdapter struct {
	context *resourceAdapterContext
	adapter *resourcesAdapter
}

func (t updateResourcesAdapter) Run(c task.CancelSignal) {
	appCtx := t.context.appCtx

	resources, err := appCtx.rpc.LoadResources(t.context.resources)
	if err != nil {
		log.E(appCtx.logger, "LoadResources returned: %v", err)
		return
	}
	log.I(appCtx.logger, "Loaded %d resources", len(resources.Textures2D))
	c.Check()
	textures := resourceGroupNode{name: "Textures"}
	for _, r := range resources.Textures2D {
		textures.nodes = append(textures.nodes, resourceTreeNode{
			context: t.context,
			info:    r,
		})
	}
	c.Check()
	appCtx.Run(func() {
		t.adapter.groups = []resourceGroupNode{textures}
		t.adapter.DataChanged(true)
	})
}

type resourceAdapterContext struct {
	appCtx    *ApplicationContext
	resources *path.Resources
	atomIndex *uint64
}

type resourceTreeNode struct {
	context *resourceAdapterContext
	info    service.ResourceInfo
}

func (n resourceTreeNode) Count() int                     { return 0 }
func (n resourceTreeNode) NodeAt(int) gxui.TreeNode       { return nil }
func (n resourceTreeNode) ItemIndex(gxui.AdapterItem) int { return 0 }
func (n resourceTreeNode) Item() gxui.AdapterItem         { return n.info.ID }
func (n resourceTreeNode) Create(theme gxui.Theme) gxui.Control {
	l := theme.CreateLinearLayout()
	l.SetDirection(gxui.LeftToRight)

	thumb := theme.CreateImage()
	thumb.SetExplicitSize(math.Size{W: resourceThumbnailWidth, H: resourceThumbnailHeight})

	name := theme.CreateLabel()
	name.SetText(n.info.Name)

	desc := theme.CreateLabel()
	desc.SetMultiline(true)

	text := theme.CreateLinearLayout()
	text.SetDirection(gxui.TopToBottom)
	text.AddChild(name)
	text.AddChild(desc)

	usedAt := theme.CreateButton()
	usedAt.SetText("Used at")
	usedAt.OnClick(func(ev gxui.MouseEvent) {
		createContextMenu(theme, n.context.appCtx.dropDownOverlay, ev.WindowPoint, n.info.Accesses, func(item gxui.AdapterItem) {
			atomIndex := item.(uint64)
			p := n.context.resources.Capture.Atoms().Index(atomIndex)
			n.context.appCtx.events.Select(p)
		})
	})

	descriptionTask, thumbnailTask := task.New(), task.New()
	update := func() {
		thumbnailTask.Cancel()
		desc.SetText("")
		thumb.SetTexture(nil)

		// Use atom index of the last access to the resource
		atomIndex := *n.context.atomIndex
		if c := len(n.info.Accesses); c > 0 {
			i := sort.Search(len(n.info.Accesses), func(i int) bool {
				return atomIndex < n.info.Accesses[i]
			}) - 1
			if i > 0 {
				atomIndex = n.info.Accesses[i]
			}
		}

		r := n.context.resources.Capture.Atoms().Index(atomIndex).ResourceAfter(n.info.ID)
		descriptionTask.Run(updateResourceDesc{n.context.appCtx, r, desc})

		t := r.Thumbnail(resourceThumbnailWidth, resourceThumbnailHeight, image.RGBA())
		thumbnailTask.Run(updateImage{n.context.appCtx, t, thumb})
	}
	l.OnAttach(update)
	l.OnDetach(thumbnailTask.Cancel)
	l.OnDetach(descriptionTask.Cancel)

	l.AddChild(thumb)
	l.AddChild(text)
	l.AddChild(usedAt)
	return l
}

func createContextMenu(theme gxui.Theme, overlay gxui.BubbleOverlay, target math.Point, items interface{}, onClick func(gxui.AdapterItem)) {
	adapter := gxui.CreateDefaultAdapter()
	adapter.SetItems(items)

	list := theme.CreateList()
	list.SetAdapter(adapter)
	list.OnSelectionChanged(onClick)
	list.OnClick(func(gxui.MouseEvent) { overlay.Hide() })
	list.OnLostFocus(func() { overlay.Hide() })

	overlay.Show(list, target)

	gxui.SetFocus(list)
}

type resourceGroupNode struct {
	name  string
	nodes []resourceTreeNode
}

func (n resourceGroupNode) Count() int                 { return len(n.nodes) }
func (n resourceGroupNode) NodeAt(i int) gxui.TreeNode { return n.nodes[i] }
func (n resourceGroupNode) ItemIndex(item gxui.AdapterItem) int {
	for i, r := range n.nodes {
		if r.Item() == item {
			return i
		}
	}
	return -1
}
func (n resourceGroupNode) Item() gxui.AdapterItem { return n.name }
func (n resourceGroupNode) Create(theme gxui.Theme) gxui.Control {
	name := theme.CreateLabel()
	name.SetText(n.name)
	return name
}

type resourcesAdapter struct {
	gxui.AdapterBase
	groups []resourceGroupNode
}

func (a resourcesAdapter) Count() int                 { return len(a.groups) }
func (a resourcesAdapter) NodeAt(i int) gxui.TreeNode { return a.groups[i] }
func (a resourcesAdapter) ItemIndex(item gxui.AdapterItem) int {
	for i, r := range a.groups {
		if r.Item() == item {
			return i
		}
		if r.ItemIndex(item) >= 0 {
			return i
		}
	}
	return -1
}

func (a resourcesAdapter) Size(gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: resourceThumbnailHeight}
}

type updateImage struct {
	context *ApplicationContext
	path    *path.Thumbnail
	image   gxui.Image
}

func (t updateImage) Run(c task.CancelSignal) {
	if img, err := t.context.rpc.LoadImageInfo(t.path); err == nil {
		c.Check()
		if img.Format != image.RGBA() {
			img, err = t.context.rpc.LoadImageInfo(t.path.As(image.RGBA()))
			if err != nil {
				return
			}
			c.Check()
		}
		if data, err := t.context.rpc.LoadBlob(img.Data); err == nil {
			t.context.Run(func() {
				tex := NewTexture(t.context.theme.Driver(), int(img.Width), int(img.Height), data)
				tex.SetFlipY(false)
				t.image.SetTexture(tex)
				t.context.toolTipController.AddToolTip(t.image, 0.7, func(math.Point) gxui.Control {
					large := t.context.theme.CreateImage()
					large.SetTexture(tex)
					large.SetAspectMode(gxui.AspectCorrectLetterbox)
					return large
				})
			})
		}
	}
}

type updateResourceDesc struct {
	context *ApplicationContext
	path    *path.Resource
	text    gxui.Label
}

func (t updateResourceDesc) Run(c task.CancelSignal) {
	if r, err := t.context.rpc.LoadResource(t.path); err == nil {
		c.Check()
		switch r := r.(type) {
		case *gfxapi.Texture2D:
			s := fmt.Sprintf("%v %dx%d\n%d mips",
				r.Levels[0].Format, r.Levels[0].Width, r.Levels[0].Height, len(r.Levels))
			t.context.Run(func() {
				t.text.SetText(s)
			})
		}
	}
}
