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
	"io"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

type sender struct {
	c    chan sendItem
	done chan struct{}
}

type sendCloseChannel struct {
	c channelId
	r chan<- error
}

type sendOpenChannel struct {
	c channelId
	r chan<- error
}

type sendData struct {
	c channelId
	d []byte
	r chan<- sendDataResult
	n int // bytes written
}

type sendDataResult struct {
	e error
	n int
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func (m sendCloseChannel) channel() channelId { return m.c }
func (m sendOpenChannel) channel() channelId  { return m.c }
func (m sendData) channel() channelId         { return m.c }

type sendItem interface {
	channel() channelId
}

func encodeOpenChannel(e binary.Encoder, s channelId) error {
	if err := msgTypeOpenChannel.encode(e); err != nil {
		return err
	}
	return e.Value(&msgOpenChannel{channelId: s})
}

func encodeCloseChannel(e binary.Encoder, s channelId) error {
	if err := msgTypeCloseChannel.encode(e); err != nil {
		return err
	}
	return e.Value(&msgCloseChannel{channelId: s})
}

func encodeData(e binary.Encoder, s channelId, d []byte) error {
	if err := msgTypeData.encode(e); err != nil {
		return err
	}
	return e.Value(&msgData{c: s, d: d})
}

func (s sender) sendCloseChannel(i channelId) error {
	r := make(chan error)
	s.c <- sendCloseChannel{i, r}
	return <-r
}

func (s sender) sendOpenChannel(i channelId) error {
	r := make(chan error)
	s.c <- sendOpenChannel{i, r}
	return <-r
}

func (s sender) sendData(i channelId, d []byte) (int, error) {
	r := make(chan sendDataResult)
	s.c <- sendData{i, d, r, 0}
	res := <-r
	return res.n, res.e
}

func (s *sender) begin(bufSize, mtu int, out io.Writer) {
	c, done := make(chan sendItem, bufSize), make(chan struct{})
	s.c, s.done = c, done

	go func() {
		close(done)
		// We get away with using a single encoder between multiple channels because we
		// do not use Encoder.Object() which is the only method that has state.
		e := cyclic.Encoder(vle.Writer(out))
		m := sendMap{}
		for {
			if len(m) == 0 {
				// If there's nothing being worked on, block until we have something.
				if item, ok := <-c; ok {
					m.pushBack(item)
				} else {
					return // Done
				}
			} else {
				// If we're busy, grab more work only if there's something there.
				select {
				case item := <-c:
					m.pushBack(item)
				default:
				}
			}

			for s := range m { // Send something on each of the pending channels...
				switch i := m.popFront(s).(type) {
				case sendOpenChannel:
					i.r <- encodeOpenChannel(e, i.c)
				case sendCloseChannel:
					i.r <- encodeCloseChannel(e, i.c)
				case sendData:
					n := min(len(i.d), mtu)
					err := encodeData(e, i.c, i.d[:n])
					if err == nil {
						i.d, i.n = i.d[n:], i.n+n
						if len(i.d) > 0 {
							m.pushFront(i) // more to send later
							continue
						}
					}
					i.r <- sendDataResult{err, i.n}
				}
			}
		}
	}()
}

func (s *sender) end() {
	close(s.c)
	<-s.done
	s.c, s.done = nil, nil
}
