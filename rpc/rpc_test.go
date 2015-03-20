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
	"crypto/sha1"
	"io"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/log"
)

const mtu = 1024

type base string

func (o *base) Encode(e binary.Encoder) error {
	return e.String(string(*o))
}

func (o *base) Decode(d binary.Decoder) error {
	s, err := d.String()
	*o = base(s)
	return err
}

type request struct{ base }
type delay struct{ base }
type response struct{ base }

var requestID = sha1.Sum([]byte("requestID"))
var delayID = sha1.Sum([]byte("delayID"))
var responseID = sha1.Sum([]byte("responseID"))

func init() {
	registry.Add(requestID, &request{})
	registry.Add(delayID, &delay{})
	registry.Add(responseID, &response{})
}

func create() Client {
	pass := make(chan base, 1)
	sr, cw := io.Pipe()
	cr, sw := io.Pipe()
	Serve(log.Nop{}, sr, sw, mtu, func(call interface{}) binary.Encodable {
		switch o := call.(type) {
		case *request:
			pass <- o.base
			return &response{o.base}
		case *delay:
			return &response{<-pass}
		default:
			return NewError("Invalid call type %T", o)
		}
	})
	return NewClient(cr, cw, mtu)
}

func simpleRequest(t *testing.T, c Client, v string) {
	r, err := c.Send(&request{base(v)})
	if err != nil {
		t.Fatalf("Unexpected error %s from rpc", err)
	}
	if r, ok := r.(*response); !ok {
		t.Fatalf("Unexpected response type %T from rpc", r)
	} else if string(r.base) != v {
		t.Fatalf("expected %s got %s from rpc", v, r.base)
	}
}

func delayRequest(t *testing.T, c Client, send string, expect string) {
	r, err := c.Send(&delay{base(send)})
	if err != nil {
		t.Fatalf("Unexpected error %s from rpc", err)
	}
	if r, ok := r.(*response); !ok {
		t.Fatalf("Unexpected response type %T from rpc", r)
	} else if string(r.base) != expect {
		t.Fatalf("expected %s got %s from rpc", expect, r.base)
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
