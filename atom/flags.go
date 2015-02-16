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

package atom

// Flags is a bitfield describing characteristics of an atom.
type Flags uint32

const (
	DrawCall Flags = 1 << iota
	EndOfFrame
)

// IsDrawCall returns true if the atom is a draw call.
func (f Flags) IsDrawCall() bool { return (f & DrawCall) != 0 }

// IsEndOfFrame returns true if the atom represents the end of a frame.
func (f Flags) IsEndOfFrame() bool { return (f & EndOfFrame) != 0 }
