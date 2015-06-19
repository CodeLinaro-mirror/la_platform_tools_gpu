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
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
)

// Handler is the signature for a function that handles incoming rpc calls.
type Handler func(interface{}) binary.Object

// Server implements the receiving side of a client server rpc pair.
// It listens on the reader for calls, and dispatches them to the supplied handler.
// Any result returned from the handler is then sent back down the writer.
func Serve(r io.Reader, w io.Writer, mtu int, l log.Logger, handler Handler) {
	multiplexer.New(r, w, mtu, func(channel io.ReadWriteCloser) {
		// If Close fails, multiplexer already knows, so we ignore the error
		defer channel.Close()

		w := bufio.NewWriterSize(channel, mtu)
		// Flush only fails if channel fails, multiplexer already knows, so we ignore the error
		defer w.Flush()

		d := cyclic.Decoder(vle.Reader(channel))
		e := cyclic.Encoder(vle.Writer(w))

		// Check the RPC header
		var h [4]byte
		if err := d.Data(h[:]); err != nil || h != header {
			l.Errorf("%v", ErrInvalidHeader)
			e.Object(ErrInvalidHeader)
			return
		}

		// Decode the call
		val, err := d.Object()
		if err != nil {
			l.Errorf("Error decoding call: %v", err)
			e.Object(NewError("Failed to decode call. Reason: %v", err))
			return
		}

		// Invoke the call
		res := handler(val)

		// Encode the call result
		if err := e.Object(res); err != nil {
			l.Errorf("Error encoding result for %T: %v", val, err)
			e.Object(NewError("Failed to encode call result. Reason: %v", err))
			return
		}
	})
}
