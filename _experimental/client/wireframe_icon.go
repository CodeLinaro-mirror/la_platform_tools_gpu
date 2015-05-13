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

//     0   1   2   3   4   5   6   7   8   9   a   b   c   d   e   f
//   ┌───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┬───┐
// 0 │ 0═╪═══╪═══╪═══╪═══╪═1 │   │   │   │   │ 4═╪═══╪═══╪═══╪═══╪═5 │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// 1 │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// 2 │ ║ │   │   │   │   │ 2═╪═══╪═══╪═══╪═══╪═3 │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼─╫─┤
// 3 │ ║ │   │   │   │   │29═╪═══╪═══╪═══╪═══╪28 │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// 4 │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// 5 │31═╪═══╪18═╪19═╪═══╪30 │   │   │   │   │27═╪═══╪26 │ 7═╪═══╪═6 │
//   ├───┼───┼─╫─┼─╫─┼───┼───┼───┼───┼───┼───┼───┼───┼─╫─┼─╫─┼───┼───┤
// 6 │   │   │ ║ │ ║ │   │   │   │   │   │   │   │   │ ║ │ ║ │   │   │
//   ├───┼───┼─╫─┼─╫─┼───┼───┼───┼───┼───┼───┼───┼───┼─╫─┼─╫─┼───┼───┤
// 7 │   │   │ ║ │ ║ │   │   │   │   │   │   │   │   │ ║ │ ║ │   │   │
//   ├───┼───┼─╫─┼─╫─┼───┼───┼───┼───┼───┼───┼───┼───┼─╫─┼─╫─┼───┼───┤
// 8 │   │   │ ║ │ ║ │   │   │   │   │   │   │   │   │ ║ │ ║ │   │   │
//   ├───┼───┼─╫─┼─╫─┼───┼───┼───┼───┼───┼───┼───┼───┼─╫─┼─╫─┼───┼───┤
// 9 │   │   │ ║ │ ║ │   │   │   │   │   │   │   │   │ ║ │ ║ │   │   │
//   ├───┼───┼─╫─┼─╫─┼───┼───┼───┼───┼───┼───┼───┼───┼─╫─┼─╫─┼───┼───┤
// a │16═╪═══╪17 │20═╪═══╪21 │   │   │   │   │24═╪═══╪25 │ 8═╪═══╪═9 │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// b │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// c │ ║ │   │   │   │   │22═╪═══╪═══╪═══╪═══╪23 │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼───┼─╫─┤
// d │ ║ │   │   │   │   │13═╪═══╪═══╪═══╪═══╪12 │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// e │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │   │   │   │   │ ║ │
//   ├─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┼───┼───┼───┼───┼─╫─┤
// f │15═╪═══╪═══╪═══╪═══╪14 │   │   │   │   │11═╪═══╪═══╪═══╪═══╪10 │
//   └───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┴───┘

var wireframeIconPoly = gxui.Polygon{
	/*  0 */ gxui.PolygonVertex{Position: math.Point{X: 0x0, Y: 0x0}},
	/*  1 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0x0}},
	/*  2 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0x2}},
	/*  3 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0x2}},
	/*  4 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0x0}},
	/*  5 */ gxui.PolygonVertex{Position: math.Point{X: 0xf, Y: 0x0}},
	/*  6 */ gxui.PolygonVertex{Position: math.Point{X: 0xf, Y: 0x5}},
	/*  7 */ gxui.PolygonVertex{Position: math.Point{X: 0xd, Y: 0x5}},
	/*  8 */ gxui.PolygonVertex{Position: math.Point{X: 0xd, Y: 0xa}},
	/*  9 */ gxui.PolygonVertex{Position: math.Point{X: 0xf, Y: 0xa}},
	/* 10 */ gxui.PolygonVertex{Position: math.Point{X: 0xf, Y: 0xf}},
	/* 11 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0xf}},
	/* 12 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0xd}},
	/* 13 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0xd}},
	/* 14 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0xf}},
	/* 15 */ gxui.PolygonVertex{Position: math.Point{X: 0x0, Y: 0xf}},
	/* 16 */ gxui.PolygonVertex{Position: math.Point{X: 0x0, Y: 0xa}},
	/* 17 */ gxui.PolygonVertex{Position: math.Point{X: 0x2, Y: 0xa}},
	/* 18 */ gxui.PolygonVertex{Position: math.Point{X: 0x2, Y: 0x5}},
	/* 19 */ gxui.PolygonVertex{Position: math.Point{X: 0x3, Y: 0x5}},
	/* 20 */ gxui.PolygonVertex{Position: math.Point{X: 0x3, Y: 0xa}},
	/* 21 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0xa}},
	/* 22 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0xc}},
	/* 23 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0xc}},
	/* 24 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0xa}},
	/* 25 */ gxui.PolygonVertex{Position: math.Point{X: 0xc, Y: 0xa}},
	/* 26 */ gxui.PolygonVertex{Position: math.Point{X: 0xc, Y: 0x5}},
	/* 27 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0x5}},
	/* 28 */ gxui.PolygonVertex{Position: math.Point{X: 0xa, Y: 0x3}},
	/* 29 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0x3}},
	/* 30 */ gxui.PolygonVertex{Position: math.Point{X: 0x5, Y: 0x5}},
	/* 31 */ gxui.PolygonVertex{Position: math.Point{X: 0x0, Y: 0x5}},
}
