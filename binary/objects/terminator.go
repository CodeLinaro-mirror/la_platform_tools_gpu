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

package objects

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
)

// Terminator is an object with no payload who's purpose is to mark the end of a
// an object stream.
type Terminator struct{}

func init() {
	registry.Global.Add((*Terminator)(nil).Class())
}

var (
	TerminatorID = binary.ID{0x01}
)

type binaryClassTerminator struct{}

func (*Terminator) Class() binary.Class                                         { return (*binaryClassTerminator)(nil) }
func (*binaryClassTerminator) ID() binary.ID                                    { return TerminatorID }
func (*binaryClassTerminator) New() binary.Object                               { return (*Terminator)(nil) }
func (*binaryClassTerminator) Encode(e binary.Encoder, obj binary.Object) error { return nil }
func (*binaryClassTerminator) Decode(d binary.Decoder) (binary.Object, error) {
	return (*Terminator)(nil), nil
}
func (*binaryClassTerminator) DecodeTo(d binary.Decoder, obj binary.Object) error { return nil }
func (*binaryClassTerminator) Skip(d binary.Decoder) error                        { return nil }
