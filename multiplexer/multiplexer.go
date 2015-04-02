/*
 * Copyright 2015, The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *     http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

//go:generate codergen -go

// Package multiplexer provides multiple data-stream multiplexing over a single binary data-stream.
package multiplexer

import (
	"errors"
	"io"
	"sync"

	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
)

var ErrChannelClosed = errors.New("Channel closed")

const sendChanSize = 256

// Multiplexer provides multiple data-stream multiplexing over a single binary data-stream.
// Channels can be opened and closed by either endpoint, and data can be sent on an open channel in
// either direction.
// The muliplexer attempts to evenly distribute bandwidth between channels.
type Multiplexer struct {
	in                    io.Reader
	out                   io.Writer
	mtu                   int
	channelOpenedCallback func(io.ReadWriteCloser)
	channels              map[channelId]*channel
	channelLock           *sync.Mutex
	sender                sender
	nextChannelId         channelId
}

func (m *Multiplexer) createChannel(id channelId) *channel {
	m.channelLock.Lock()
	defer m.channelLock.Unlock()

	if len(m.channels) == 0 {
		m.sender.begin(sendChanSize, m.mtu, m.out)
	}
	s := newChannel(id, m)
	m.channels[id] = s
	return s
}

func (m *Multiplexer) closeChannel(id channelId, sendMsg bool) error {
	m.channelLock.Lock()
	defer m.channelLock.Unlock()

	if channel, found := m.channels[id]; found {
		var err error
		if sendMsg {
			err = m.sender.sendCloseChannel(id)
		}
		channel.closeInput()
		delete(m.channels, id)
		if len(m.channels) == 0 {
			m.sender.end()
		}
		return err
	} else {
		// Attempting to close an unknown channel.
		// This can happen when both ends close simultaneously.
		return ErrChannelClosed
	}
}

func (m *Multiplexer) writeChannel(id channelId, data []byte) (n int, err error) {
	m.channelLock.Lock()
	_, found := m.channels[id]
	defer m.channelLock.Unlock()

	if found {
		return m.sender.sendData(id, data)
	} else {
		return 0, ErrChannelClosed
	}
}

func (m *Multiplexer) recv() {
	defer m.closeAllChannels()
	d := cyclic.Decoder(vle.Reader(m.in))
	for {
		var ty msgType
		if err := ty.decode(d); err != nil {
			return
		}
		switch ty {
		case msgTypeOpenChannel:
			msg := &msgOpenChannel{}
			if err := d.Value(msg); err != nil {
				return
			}
			s := m.createChannel(remote(msg.channelId))
			go m.channelOpenedCallback(s)

		case msgTypeCloseChannel:
			msg := &msgCloseChannel{}
			if err := d.Value(msg); err != nil {
				return
			}
			m.closeChannel(remote(msg.channelId), false)

		case msgTypeData:
			msg := &msgData{}
			if err := d.Value(msg); err != nil {
				return
			}
			id := remote(msg.c)
			if channel, found := m.channels[id]; found {
				channel.receive(msg.d)
			} else {
				// Likely this channel was closed this side, and we're receiving data
				// that should be dropped on the floor.
			}
		}
	}
}

func (m *Multiplexer) closeAllChannels() {
	m.channelLock.Lock()
	defer m.channelLock.Unlock()

	if len(m.channels) > 0 {
		for _, c := range m.channels {
			c.inWriter.Close()
		}
		m.sender.end()
	}
	m.channels = nil
}

// OpenChannel opens a new channel for communication. This will invoke the the channel-opened
// callback on the remote endpoint.
func (m *Multiplexer) OpenChannel() (io.ReadWriteCloser, error) {
	id := m.nextChannelId.increment()
	s := m.createChannel(id)
	if err := m.sender.sendOpenChannel(id); err == nil {
		return s, nil
	} else {
		m.closeChannel(id, false)
		return nil, err
	}
}

// MTU returns the maximum transmission unit size the multiplexer was created with.
func (m *Multiplexer) MTU() int { return m.mtu }

// New creates and returns a new Multiplexer using the specified reader and writer for
// communication. mtu defines the maximum size of each packet of data. channelOpenedCallback will be
// called for each channel that was opened by the remote endpoint.
func New(in io.Reader, out io.Writer, mtu int, channelOpenedCallback func(io.ReadWriteCloser)) *Multiplexer {
	m := &Multiplexer{
		in:  in,
		out: out,
		mtu: mtu,
		channelOpenedCallback: channelOpenedCallback,
		channels:              make(map[channelId]*channel),
		channelLock:           &sync.Mutex{},
	}
	go m.recv()
	return m
}
