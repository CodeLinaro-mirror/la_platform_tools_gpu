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

package rpc

import (
	"bufio"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
)

// Client implements the sending side of a client-server rpc pair.
type Client struct {
	m *multiplexer.Multiplexer
	n *registry.Namespace
}

// NewClient creates a new rpc client object that uses the multiplexer m for
// communication the namespace n for decoding objects. If n is nil then the
// global namespace is used.
func NewClient(m *multiplexer.Multiplexer, n *registry.Namespace) Client {
	return Client{
		m: m,
		n: n,
	}
}

// Multiplexer returns the multiplexer used for communication to the server.
func (c Client) Multiplexer() *multiplexer.Multiplexer {
	return c.m
}

// Namespace returns the custom namespace used for decoding responses from the
// server, or nil if no custom namespace has been specified.
func (c Client) Namespace() *registry.Namespace {
	return c.n
}

// Send encodes an rpc call and sends it to the server.
// It blocks until a reply is received or an error indicating there will be no
// reply occurs.
// This method is safe for concurrent use.
func (c Client) Send(call binary.Object) (interface{}, error) {
	channel, err := c.m.OpenChannel()
	if err != nil {
		return nil, err
	}
	// We can ignore channel close failures if we already have a complete response
	defer channel.Close()

	w := bufio.NewWriterSize(channel, c.m.MTU())
	d := cyclic.Decoder(vle.Reader(channel))
	e := cyclic.Encoder(vle.Writer(w))

	if c.n != nil {
		d.Namespace = c.n // Use custom decoding namespace.
	}

	// Write the RPC header
	if e.Data(header[:]); d.Error() != nil {
		return nil, d.Error()
	}

	// Write the call
	if e.Object(call); d.Error() != nil {
		return nil, d.Error()
	}

	// Flush the bufio writer
	if err := w.Flush(); err != nil {
		return nil, err
	}

	// Wait for and read the response
	res := d.Object()
	if d.Error() != nil {
		return nil, d.Error()
	}

	// Check to see if the response was an error
	if err, b := res.(*Error); b {
		return nil, err
	}

	return res, nil
}
