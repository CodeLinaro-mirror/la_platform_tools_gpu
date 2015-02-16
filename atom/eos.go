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

import "android.googlesource.com/platform/tools/gpu/binary"

// TypeIdEos is the EOS's unique type identifier.
const TypeIdEos TypeId = 0xffff

// EOS is used to indicate that there will be no more atoms in the stream with
// the EOS's context identifier after the EOS atom.
type EOS struct {
	Context ContextId
}

func (c *EOS) Name() string                 { return "EOS" }
func (c *EOS) Docs() string                 { return "" }
func (c *EOS) TypeId() TypeId               { return TypeIdEos }
func (c *EOS) ContextId() ContextId         { return c.Context }
func (c *EOS) String() string               { return "EOS" }
func (c *EOS) Flags() Flags                 { return 0 }
func (c *EOS) Encode(*binary.Encoder) error { return nil }
func (c *EOS) Decode(*binary.Decoder) error { return nil }
