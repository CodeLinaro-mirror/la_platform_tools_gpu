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

import "container/list"

type sendMap map[channelId]*list.List

func (m sendMap) pushBack(i sendItem) {
	s := i.channel()
	l, f := m[s]
	if !f {
		l = list.New()
		m[s] = l
	}
	l.PushBack(i)
}

func (m sendMap) pushFront(i sendItem) {
	s := i.channel()
	l, f := m[s]
	if !f {
		l = list.New()
		m[s] = l
	}
	l.PushFront(i)
}

func (m sendMap) popFront(s channelId) sendItem {
	l := m[s]
	f := l.Front()
	l.Remove(f)
	if l.Len() == 0 {
		delete(m, s)
	}
	return f.Value.(sendItem)
}
