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
	"io"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/log"
)

const mtu = 1024

type request struct {
	binary.Generate
	data string
}

type delay struct {
	binary.Generate
	data string
}

type response struct {
	binary.Generate
	data string
}

func create() Client {
	pass := make(chan string, 1)
	sr, cw := io.Pipe()
	cr, sw := io.Pipe()
	Serve(log.Nop{}, sr, sw, mtu, func(call interface{}) binary.Encodable {
		switch o := call.(type) {
		case *request:
			pass <- o.data
			return &response{data: o.data}
		case *delay:
			return &response{data: <-pass}
		default:
			return NewError("Invalid call type %T", o)
		}
	})
	return NewClient(cr, cw, mtu)
}

func simpleRequest(t *testing.T, c Client, v string) {
	r, err := c.Send(&request{data: v})
	if err != nil {
		t.Fatalf("Unexpected error %s from rpc", err)
	}
	if r, ok := r.(*response); !ok {
		t.Fatalf("Unexpected response type %T from rpc", r)
	} else if string(r.data) != v {
		t.Fatalf("expected %s got %s from rpc", v, r.data)
	}
}

func delayRequest(t *testing.T, c Client, send string, expect string) {
	r, err := c.Send(&delay{data: send})
	if err != nil {
		t.Fatalf("Unexpected error %s from rpc", err)
	}
	if r, ok := r.(*response); !ok {
		t.Fatalf("Unexpected response type %T from rpc", r)
	} else if string(r.data) != expect {
		t.Fatalf("expected %s got %s from rpc", expect, r.data)
	}
}

func TestSimpleRpc(t *testing.T) {
	c := create()
	simpleRequest(t, c, "hello")
}

func TestInterleavedRpc(t *testing.T) {
	c := create()
	done := make(chan struct{})
	go func() {
		delayRequest(t, c, "hello", "goodbye")
		close(done)
	}()
	simpleRequest(t, c, "goodbye")
	<-done
}
