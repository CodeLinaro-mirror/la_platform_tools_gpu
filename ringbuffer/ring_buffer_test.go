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

package ringbuffer

import (
	"io"
	"testing"
)

type ioCheck struct {
	bytes []byte
	n     int
	err   error
}

func bytesEqual(a, b []byte) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func checkWrite(t *testing.T, w io.Writer, check ioCheck) {
	n, err := w.Write(check.bytes)
	if check.n != n {
		t.Errorf("Write() wrote an unexpected number of bytes. Expected: %d, wrote: %d.", check.n, n)
	}
	if check.err != err {
		t.Errorf("Write() returned unexpected error. Expected: %v, got: %v.", check.err, err)
	}
}

func checkRead(t *testing.T, r io.Reader, check ioCheck) {
	buf := make([]byte, len(check.bytes))
	n, err := io.ReadFull(r, buf)
	if !bytesEqual(check.bytes, buf) {
		t.Errorf("Read() read bytes were not as expected. Expected: %v, got: %v.", check.bytes, buf)
	}
	if check.n != n {
		t.Errorf("Read() read an unexpected number of bytes. Expected: %d, read: %d.", check.n, n)
	}
	if check.err != err {
		t.Errorf("Read() returned unexpected error. Expected: %v, got: %v.", check.err, err)
	}
}

func TestRingBufferWriteSimple(t *testing.T) {
	checks := [][]ioCheck{
		{
			{[]byte{1, 2, 3, 4, 5}, 5, nil},
		}, {
			{[]byte{1, 2, 3}, 3, nil},
			{[]byte{4, 5}, 2, nil},
		}, {
			{[]byte{1, 2, 3, 4}, 4, nil},
			{[]byte{5}, 1, nil},
		},
	}

	for _, pass := range checks {
		b := New(5)
		for _, check := range pass {
			checkWrite(t, b, check)
		}
	}
}

func TestRingBufferWriteWhenFull(t *testing.T) {
	b := New(5)

	// Fill a ring buffer
	b.Write([]byte{1, 2, 3, 4, 5})

	// Further writes must block until there's a read to free space
	go checkWrite(t, b, ioCheck{[]byte{6, 7, 8, 9}, 4, nil})

	checkRead(t, b, ioCheck{[]byte{1, 2, 3, 4}, 4, nil})
}

func TestRingBufferWriteWhenClosed(t *testing.T) {
	b := New(5)

	b.Close()
	checkWrite(t, b, ioCheck{[]byte{1, 2, 3, 4, 5}, 0, io.ErrClosedPipe})
}

func TestRingBufferWriteWhenFullAndClosed(t *testing.T) {
	b := New(5)

	checkWrite(t, b, ioCheck{[]byte{1, 2, 3, 4, 5}, 5, nil})

	b.Close()
	checkWrite(t, b, ioCheck{[]byte{1, 2, 3, 4, 5}, 0, io.ErrClosedPipe})
}

func TestRingBufferReadWhenEmpty(t *testing.T) {
	b := New(5)

	// Reads must block until there's something to read
	go checkRead(t, b, ioCheck{[]byte{1, 2, 3, 4}, 4, nil})

	checkWrite(t, b, ioCheck{[]byte{1, 2, 3, 4}, 4, nil})
}

func TestRingBufferReadWhenClosed(t *testing.T) {
	b := New(5)

	b.Close()
	checkRead(t, b, ioCheck{[]byte{0, 0, 0, 0}, 0, io.EOF})
}

func TestRingBufferReadThenClose(t *testing.T) {
	b := New(5)
	go func() {
		checkWrite(t, b, ioCheck{[]byte{1, 2, 3, 4}, 4, nil})
		b.Close()
	}()

	checkRead(t, b, ioCheck{[]byte{1, 2, 3, 4}, 4, nil})
}

func TestRingBufferIncompleteReadThenClose(t *testing.T) {
	b := New(5)
	go func() {
		checkWrite(t, b, ioCheck{[]byte{1, 2}, 2, nil})
		b.Close()
	}()

	checkRead(t, b, ioCheck{[]byte{1, 2, 0, 0}, 2, io.ErrUnexpectedEOF})
}

func TestRingBufferLargeWrite(t *testing.T) {
	b := New(5)
	go func() {
		checkWrite(t, b, ioCheck{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15, nil})
		b.Close()
	}()

	checkRead(t, b, ioCheck{[]byte{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}, 15, nil})
}

func TestRingBufferReadWrite(t *testing.T) {
	b := New(5)
	go func() {
		buf := [7]byte{}
		for i := 0; i < 1000; i++ {
			for j := range buf {
				buf[j] = byte(i*len(buf) + j)
			}
			checkWrite(t, b, ioCheck{buf[:], len(buf), nil})
		}
	}()

	buf := [4]byte{}
	for i := 0; i < 1750; i++ {
		for j := range buf {
			buf[j] = byte(i*len(buf) + j)
		}
		checkRead(t, b, ioCheck{buf[:], len(buf), nil})
	}
}
