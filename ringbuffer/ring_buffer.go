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

// Package ringbuffer implements an in-memory circular buffer conforming to the io.ReadWriteCloser
// interface.
package ringbuffer

import (
	"io"
	"sync"
)

type ringBuffer struct {
	buf      []byte
	tail     int
	count    int
	closed   bool
	lock     *sync.Mutex
	notFull  *sync.Cond
	notEmpty *sync.Cond
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}

func ringRead(offset int, ring []byte, out []byte) {
	if offset+len(out) < len(ring) {
		copy(out, ring[offset:])
	} else {
		copy(out[copy(out, ring[offset:]):], ring)
	}
}

func ringWrite(offset int, ring []byte, in []byte) {
	if offset+len(in) < len(ring) {
		copy(ring[offset:], in)
	} else {
		copy(ring, in[copy(ring[offset:], in):])
	}
}

func (b *ringBuffer) Read(p []byte) (n int, err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	for {
		if b.count > 0 {
			n = min(len(p), b.count)
			ringRead(b.tail, b.buf, p[:n])
			b.tail = (b.tail + n) % len(b.buf)
			b.count -= n
			b.notFull.Broadcast()
			if b.count == 0 && b.closed {
				return n, io.EOF
			} else {
				return n, nil
			}
		} else {
			if b.closed {
				return 0, io.EOF
			}
			b.notEmpty.Wait()
		}
	}
}

func (b *ringBuffer) Write(p []byte) (n int, err error) {
	b.lock.Lock()
	defer b.lock.Unlock()

	for len(p) > 0 {
		if b.closed {
			return n, io.ErrClosedPipe
		}
		space := len(b.buf) - b.count
		if space > 0 {
			head := (b.tail + b.count) % len(b.buf)
			c := min(len(p), space)
			ringWrite(head, b.buf, p[:c])
			b.count, p, n = b.count+c, p[c:], n+c
			b.notEmpty.Broadcast()
		} else {
			b.notFull.Wait()
		}
	}
	return n, nil
}

func (b *ringBuffer) Close() error {
	b.lock.Lock()
	defer b.lock.Unlock()

	b.closed = true
	b.notEmpty.Broadcast()
	b.notFull.Broadcast()

	return nil
}

// New constructs a new ring-buffer with the specified capacity in bytes.
//
// Writes to the ring-buffer will block until all the bytes are written into the buffer or the
// buffer is closed (whichever comes first.)
// Reads from the ring-buffer will block until a single byte is read or the buffer is closed
// (whichever comes first.)
// It is safe to call Read and Write in parallel with each other or with Close.
// If the ring-buffer is closed while a read or write is in progress then io.ErrClosedPipe will
// be returned by the read / write function.
func New(capacity int) io.ReadWriteCloser {
	lock := &sync.Mutex{}
	return &ringBuffer{
		lock:     lock,
		notFull:  sync.NewCond(lock),
		notEmpty: sync.NewCond(lock),
		buf:      make([]byte, capacity),
	}
}
