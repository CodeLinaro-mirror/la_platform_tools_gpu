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
	"io"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
)

// Client implements the sending side of a client-server rpc pair.
type Client struct {
	m *multiplexer.Multiplexer
}

// NewClient creates a new rpc client object that sends messages down the
// writer, and waits for responses on the reader.
// The client supports multiple in-flight rpc calls.
// The chunk size used for multiplexing the calls for fairness is specified by mtu.
func NewClient(r io.Reader, w io.Writer, mtu int) Client {
	return Client{
		m: multiplexer.New(r, w, mtu, nil),
	}
}

// Send encodes an rpc call and sends it to the server.
// It blocks until a reply is received or an error indicating there will be no
// reply occurs.
// This method is safe for concurrent use.
func (b Client) Send(call binary.Encodable) (interface{}, error) {
	channel, err := b.m.OpenChannel()
	if err != nil {
		return nil, err
	}
	// We can ignore channel close failures if we already have a complete response
	defer channel.Close()

	w := bufio.NewWriterSize(channel, b.m.MTU())
	d := binary.NewDecoder(channel)
	e := binary.NewEncoder(w)

	// Write the RPC header
	if err := e.Uint32(header); err != nil {
		return nil, err
	}

	// Write the call
	if err := e.Object(call); err != nil {
		return nil, err
	}

	// Flush the bufio writer
	if err := w.Flush(); err != nil {
		return nil, err
	}

	// Wait for and read the response
	res, err := d.Object()
	if err != nil {
		return nil, err
	}

	// Check to see if the response was an error
	if err, b := res.(*Error); b {
		return nil, err
	}

	return res, nil
}
