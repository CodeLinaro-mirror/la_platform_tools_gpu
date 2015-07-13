package client

import (
	"github.com/google/gxui"

	"android.googlesource.com/platform/tools/gpu/service/path"
)

type Events struct {
	onSelect gxui.Event
}

func (e *Events) Init() {
	e.onSelect = gxui.CreateEvent(func(path.Path) {})
}

func (e *Events) OnSelect(f func(p path.Path)) gxui.EventSubscription {
	return e.onSelect.Listen(f)
}

func (e *Events) Select(p path.Path) {
	e.onSelect.Fire(p)
}
