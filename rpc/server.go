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

// Serve implements the receiving side of a client server rpc pair.
// It listens on the reader for calls, and dispatches them to the supplied handler.
// Any result returned from the handler is then sent back down the writer.
func Serve(r io.Reader, w io.Writer, c io.Closer, mtu int, l log.Logger, handler Handler) {
	multiplexer.New(r, w, c, mtu, l, func(channel io.ReadWriteCloser) {
		// If Close fails, multiplexer already knows, so we ignore the error
		defer channel.Close()

		w := bufio.NewWriterSize(channel, mtu)
		// Flush only fails if channel fails, multiplexer already knows, so we ignore the error
		defer w.Flush()

		d := cyclic.Decoder(vle.Reader(channel))
		e := cyclic.Encoder(vle.Writer(w))

		// Check the RPC header
		var h [4]byte
		if d.Data(h[:]); d.Error() != nil || h != header {
			log.Errorf(l, "%v", ErrInvalidHeader)
			e.Object(ErrInvalidHeader)
			return
		}

		// Decode the call
		val := d.Object()
		if d.Error() != nil {
			log.Errorf(l, "Error decoding call: %v", d.Error())
			e.Object(NewError("Failed to decode call. Reason: %v", d.Error()))
			return
		}

		// Invoke the call
		res := handler(val)

		// Encode the call result
		e.Object(res)
		if err := e.Error(); err != nil {
			log.Errorf(l, "Error encoding result for %+v: %v", val, err)
			e.Object(NewError("Failed to encode call result. Reason: %v", err))
			return
		}
	})
}
