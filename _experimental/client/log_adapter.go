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
	"android.googlesource.com/platform/tools/gpu/log"
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

type LogAdapter struct {
	gxui.AdapterBase
	logger     log.Logger
	maxLength  int
	entries    []log.Entry
	head, tail int
	tailIndex  int
}

func CreateLogAdapter(maxLength int, callOnUI func(func()) bool) *LogAdapter {
	a := &LogAdapter{
		maxLength: maxLength,
	}

	c := make(chan interface{}, 64)
	go func() {
		for t := range c {
			switch t := t.(type) {
			case log.Entry:
				callOnUI(func() { a.add(t) })
			case log.FlushRequest:
				close(t)
			}
		}
	}()
	a.logger = log.Channel(c)

	return a
}

func (a *LogAdapter) add(entry log.Entry) {
	if len(a.entries) < a.maxLength {
		a.entries = append(a.entries, entry)
	} else {
		a.entries[a.head] = entry
		a.tail = (a.tail + 1) % a.maxLength
		a.tailIndex++
	}
	a.head = (a.head + 1) % a.maxLength

	a.DataChanged(false)
}

func (a *LogAdapter) Logger() log.Logger {
	return a.logger
}

func (a *LogAdapter) Entry(index int) log.Entry {
	return a.entries[(a.tail+index)%a.maxLength]
}

func (a *LogAdapter) Clear() {
	a.entries = a.entries[:0]
	a.head = 0
	a.tail = 0
	a.DataReplaced()
}

func (a *LogAdapter) Count() int {
	return len(a.entries)
}

func (a *LogAdapter) ItemAt(index int) gxui.AdapterItem {
	return a.tailIndex + index
}

func (a *LogAdapter) ItemIndex(item gxui.AdapterItem) int {
	return item.(int) - a.tailIndex
}

func (a *LogAdapter) Size(theme gxui.Theme) math.Size {
	return math.Size{W: math.MaxSize.W, H: theme.DefaultFont().GlyphMaxSize().H}
}

func (a *LogAdapter) Create(theme gxui.Theme, index int) gxui.Control {
	entry := a.Entry(index)
	l := theme.CreateLabel()
	l.SetText(entry.String())
	if entry.Severity <= log.Error {
		l.SetColor(gxui.Red)
	} else {
		l.SetColor(gxui.Yellow)
	}
	return l
}
