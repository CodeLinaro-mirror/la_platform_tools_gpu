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
	"github.com/google/gxui"
	"github.com/google/gxui/math"
)

var (
	LINE_NUMBER_COLOR        gxui.Color = gxui.ColorFromHex(0xFF1CAFFF)
	CODE_COLOR               gxui.Color = gxui.ColorFromHex(0xFFFFE8BB)
	COMMAND_COLOR            gxui.Color = gxui.ColorFromHex(0xFFFB9868)
	MEMORY_OBSERVATION_COLOR gxui.Color = gxui.ColorFromHex(0xFFA1CF8A)
	INACTIVE_COLOR           gxui.Color = gxui.ColorFromHex(0xFF505050)
	CONSTANT_COLOR           gxui.Color = gxui.ColorFromHex(0xFFDDB7FF)
	READ_MEMORY_COLOR        gxui.Color = gxui.Green
	WRITE_MEMORY_COLOR       gxui.Color = gxui.Red
	STALE_MEMORY_COLOR       gxui.Color = gxui.ColorFromHex(0xFF8A7753)
)

func CreateLabel(t gxui.Theme, s string, c gxui.Color, active bool) gxui.Label {
	l := t.CreateLabel()
	l.SetText(s)
	if active {
		l.SetColor(c)
	} else {
		l.SetColor(INACTIVE_COLOR)
	}
	l.SetMargin(math.Spacing{})
	return l
}
