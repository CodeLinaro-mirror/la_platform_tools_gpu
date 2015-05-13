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

import (
	"sync/atomic"

	"android.googlesource.com/platform/tools/gpu/binary"
)

type channelId uint32

func remote(i channelId) channelId {
	return ^i
}

func (i *channelId) increment() (old channelId) {
	return channelId(atomic.AddUint32((*uint32)(i), 1) - 1)
}

func (i channelId) encode(e binary.Encoder) error {
	return e.Uint32(uint32(i))
}

func (i *channelId) decode(d binary.Decoder) error {
	if val, err := d.Uint32(); err == nil {
		*i = channelId(val)
		return nil
	} else {
		return err
	}
}
