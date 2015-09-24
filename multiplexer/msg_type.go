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

package multiplexer

import "android.googlesource.com/platform/tools/gpu/binary"

type msgType uint8

func (i msgType) encode(e binary.Encoder) error {
	e.Uint8(uint8(i))
	return e.Error()
}

func (i *msgType) decode(d binary.Decoder) error {
	val := d.Uint8()
	if d.Error() != nil {
		return d.Error()
	}
	*i = msgType(val)
	return nil
}
