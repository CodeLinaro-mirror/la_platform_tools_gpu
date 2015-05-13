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
)

type PixelType interface {
	Name() string
	SizeBytes() int
	Read(buffer []byte) gxui.Color
}

type RGB888 int
type RGBA8888 int

func (v RGB888) Name() string {
	return "rgb888"
}

func (v RGB888) SizeBytes() int {
	return 3
}

func (v RGB888) Read(data []byte) gxui.Color {
	return gxui.Color{
		R: float32(data[0]) / 255.0,
		G: float32(data[1]) / 255.0,
		B: float32(data[2]) / 255.0,
		A: 1.0,
	}
}

func (v RGBA8888) Name() string {
	return "rgba8888"
}

func (v RGBA8888) SizeBytes() int {
	return 4
}

func (v RGBA8888) Read(data []byte) gxui.Color {
	return gxui.Color{
		R: float32(data[0]) / 255.0,
		G: float32(data[1]) / 255.0,
		B: float32(data[2]) / 255.0,
		A: float32(data[3]) / 255.0,
	}
}
