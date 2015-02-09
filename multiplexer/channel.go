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

package multiplexer

import "io"

type channelEvents interface {
	closeChannel(channelId channelId, sendMsg bool) error
	writeChannel(channelId, []byte) (int, error)
}

type channel struct {
	id       channelId
	inReader io.ReadCloser
	inWriter io.WriteCloser
	events   channelEvents
}

func newChannel(id channelId, events channelEvents) *channel {
	inReader, inWriter := io.Pipe()
	return &channel{
		id:       id,
		inReader: inReader,
		inWriter: inWriter,
		events:   events,
	}
}

func (c *channel) closeInput() error {
	return c.inWriter.Close()
}

func (c *channel) receive(p []byte) error {
	_, err := c.inWriter.Write(p)
	return err
}

func (c *channel) Close() error {
	return c.events.closeChannel(c.id, true)
}

func (c *channel) Read(p []byte) (n int, err error) {
	return c.inReader.Read(p)
}

func (c *channel) Write(p []byte) (n int, err error) {
	return c.events.writeChannel(c.id, p)
}
