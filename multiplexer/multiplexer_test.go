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

import (
	"errors"
	"io"
	"testing"

	"android.googlesource.com/platform/tools/gpu/ringbuffer"
)

const mtu = 1024

var testDataPtoQ = []byte("p->q")
var testDataQtoP = []byte("q->p")
var testDataRtoS = []byte("r->s")
var testDataStoR = []byte("s->r")
var testErr = errors.New("Oh noes!")

type errWriter struct{ err error }

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

type writeCheck struct {
	bytes       []byte
	expectedN   int
	expectedErr error
}

func checkWrite(t *testing.T, w io.Writer, check writeCheck) {
	n, err := w.Write(check.bytes)
	if check.expectedN != n {
		t.Errorf("Write() wrote an unexpected number of bytes. Expected: %d, wrote: %d.", check.expectedN, n)
	}
	if check.expectedErr != err {
		t.Errorf("Write() returned unexpected error. Expected: %v, got: %v.", check.expectedErr, err)
	}
}

type readCheck struct {
	expectedBytes []byte
	expectedN     int
	expectedErr   error
}

func checkRead(t *testing.T, r io.Reader, check readCheck) {
	buf := make([]byte, len(check.expectedBytes))
	n, err := io.ReadFull(r, buf)
	if !bytesEqual(check.expectedBytes, buf) {
		t.Errorf("Read() read bytes were not as expected. Expected: %v, got: %v.", check.expectedBytes, buf)
	}
	if check.expectedN != n {
		t.Errorf("Read() read an unexpected number of bytes. Expected: %d, read: %d.", check.expectedN, n)
	}
	if check.expectedErr != err {
		t.Errorf("Read() returned unexpected error. Expected: %v, got: %v.", check.expectedErr, err)
	}
}

func (w errWriter) Write(p []byte) (int, error) {
	if w.err == nil {
		return len(p), nil
	} else {
		return 0, w.err
	}
}

func checkAllChannelsClosed(t *testing.T, m *Multiplexer) {
	m.channelLock.Lock()
	defer m.channelLock.Unlock()
	if len(m.channels) > 0 {
		t.Errorf("Multiplexer test ended with %d unclosed channels!", len(m.channels))
	}
}

func TestSendOpenCloseChannel(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel0, err := multiplexer.OpenChannel()
	if err != nil {
		t.Errorf("Error opening channel 0: %v", err)
	}

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeOpenChannel),
		0, // channel id
	}, 2, nil})

	if err := channel0.Close(); err != nil {
		t.Errorf("Error closing channel p: %v", err)
	}

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeCloseChannel),
		0, // channel id
	}, 2, nil})

	channel1, err := multiplexer.OpenChannel()
	if err != nil {
		t.Errorf("Error opening channel 1: %v", err)
	}

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeOpenChannel),
		1, // channel id
	}, 2, nil})

	if err := channel1.Close(); err != nil {
		t.Errorf("Error closing channel p: %v", err)
	}

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeCloseChannel),
		1, // channel id
	}, 2, nil})
}

func TestRecvOpenCloseChannel(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)

	channelChan := make(chan io.ReadWriteCloser, 1)
	multiplexer := New(inBuf, outBuf, mtu, func(s io.ReadWriteCloser) { channelChan <- s })
	defer checkAllChannelsClosed(t, multiplexer)

	inBuf.Write([]byte{
		byte(msgTypeOpenChannel),
		0, // channel id
	})

	channel0 := <-channelChan

	inBuf.Write([]byte{
		byte(msgTypeOpenChannel),
		1, // channel id
	})

	channel1 := <-channelChan

	inBuf.Write([]byte{
		byte(msgTypeCloseChannel),
		0, // channel id
	})

	inBuf.Write([]byte{
		byte(msgTypeCloseChannel),
		1, // channel id
	})

	checkRead(t, channel0, readCheck{[]byte{0}, 0, io.EOF})
	checkRead(t, channel1, readCheck{[]byte{0}, 0, io.EOF})
}

func TestWriteChannel(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel, _ := multiplexer.OpenChannel()
	defer channel.Close()

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeOpenChannel),
		0, // channel id
	}, 2, nil})

	checkWrite(t, channel, writeCheck{[]byte{'h', 'e', 'l', 'l', 'o'}, 5, nil})

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeData),
		0,                       // channel id
		10,                      // data length
		'h', 'e', 'l', 'l', 'o', // data
	}, 8, nil})
}

func TestOpenChannelError(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), &errWriter{nil}
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	outBuf.err = testErr
	if _, err := multiplexer.OpenChannel(); err != testErr {
		t.Errorf("Expected error: %v, got: %v", testErr, err)
	}
}

func TestCloseChannelError(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), &errWriter{nil}
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel, _ := multiplexer.OpenChannel()

	outBuf.err = testErr
	if err := channel.Close(); err != testErr {
		t.Errorf("Expected error: %v, got: %v", testErr, err)
	}
}

func TestCloseChannelTwice(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel, _ := multiplexer.OpenChannel()
	channel.Close()

	if err := channel.Close(); err != ErrChannelClosed {
		t.Errorf("Unexpected error. Expected: %v, got: %v", ErrChannelClosed, err)
	}
}

func TestWriteChannelError(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), &errWriter{nil}
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel, _ := multiplexer.OpenChannel()
	defer channel.Close()

	outBuf.err = testErr
	checkWrite(t, channel, writeCheck{[]byte{'h', 'e', 'l', 'l', 'o'}, 0, testErr})
}

func TestMultiPacketWriteChannel(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)
	multiplexer := New(inBuf, outBuf, 3, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel, _ := multiplexer.OpenChannel()
	defer channel.Close()

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeOpenChannel),
		0, // channel id
	}, 2, nil})

	checkWrite(t, channel, writeCheck{[]byte{
		'h', 'e', 'l', 'l', 'o', ' ', 'w', 'o', 'r', 'l', 'd',
	}, 11, nil})

	checkRead(t, outBuf, readCheck{[]byte{
		byte(msgTypeData),
		0,             // channel id
		6,             // data length
		'h', 'e', 'l', // data

		byte(msgTypeData),
		0,             // channel id
		6,             // data length
		'l', 'o', ' ', // data

		byte(msgTypeData),
		0,             // channel id
		6,             // data length
		'w', 'o', 'r', // data

		byte(msgTypeData),
		0,        // channel id
		4,        // data length
		'l', 'd', // data
	}, 23, nil})
}

func TestRecvChannel(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)

	channelChan := make(chan io.ReadWriteCloser, 1)
	multiplexer := New(inBuf, outBuf, mtu, func(s io.ReadWriteCloser) { channelChan <- s })
	defer checkAllChannelsClosed(t, multiplexer)

	inBuf.Write([]byte{
		byte(msgTypeOpenChannel),
		0, // channel id
	})

	channel0 := <-channelChan
	defer channel0.Close()

	inBuf.Write([]byte{
		byte(msgTypeData),
		0,                       // channel id
		10,                      // data length
		'h', 'e', 'l', 'l', 'o', // data
	})

	checkRead(t, channel0, readCheck{[]byte{'h', 'e', 'l', 'l', 'o'}, 5, nil})
}

func TestRecvOnClosedChannel(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)

	channelChan := make(chan io.ReadWriteCloser, 1)
	multiplexer := New(inBuf, outBuf, mtu, func(s io.ReadWriteCloser) { channelChan <- s })
	defer checkAllChannelsClosed(t, multiplexer)

	// Remote opens two channels
	inBuf.Write([]byte{
		byte(msgTypeOpenChannel),
		0, // channel id
	})
	channel0 := <-channelChan

	inBuf.Write([]byte{
		byte(msgTypeOpenChannel),
		1, // channel id
	})
	channel1 := <-channelChan

	// Remote endlessly sends data on both channels
	sendData := make(chan struct{})
	go func() {
		for {
			<-sendData
			inBuf.Write([]byte{
				byte(msgTypeData),
				0,                  // channel id
				8,                  // data length
				'd', 'a', 't', 'a', // data
			})
			inBuf.Write([]byte{
				byte(msgTypeData),
				1,                  // channel id
				8,                  // data length
				'd', 'a', 't', 'a', // data
			})
		}
	}()

	// When remote sends data on both channels
	sendData <- struct{}{}
	// Then local should receive data on both channels
	checkRead(t, channel0, readCheck{[]byte{'d', 'a', 't', 'a'}, 4, nil})
	checkRead(t, channel1, readCheck{[]byte{'d', 'a', 't', 'a'}, 4, nil})

	// When closes channel0...
	channel0.Close()
	// ... and remote continues to send data on both channels
	sendData <- struct{}{}

	// Then data on channel0 should be ignored (EOF)
	checkRead(t, channel0, readCheck{[]byte{0}, 0, io.EOF})
	// And data on channel1 should continue to be received
	checkRead(t, channel1, readCheck{[]byte{'d', 'a', 't', 'a'}, 4, nil})
	sendData <- struct{}{}
	checkRead(t, channel1, readCheck{[]byte{'d', 'a', 't', 'a'}, 4, nil})
	sendData <- struct{}{}
	checkRead(t, channel1, readCheck{[]byte{'d', 'a', 't', 'a'}, 4, nil})

	channel1.Close()
}

func TestLocalClosedChannelReadGivesEOF(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel, _ := multiplexer.OpenChannel()

	if err := channel.Close(); err != nil {
		t.Errorf("Error closing channel: %v", err)
	}

	checkRead(t, channel, readCheck{[]byte{0}, 0, io.EOF})
}

func TestLocalClosedChannelWriteGivesErrChannelClosed(t *testing.T) {
	inBuf, outBuf := ringbuffer.New(1024), ringbuffer.New(1024)
	multiplexer := New(inBuf, outBuf, mtu, nil)
	defer checkAllChannelsClosed(t, multiplexer)

	channel, _ := multiplexer.OpenChannel()
	channel.Close()
	checkWrite(t, channel, writeCheck{[]byte{'h', 'e', 'l', 'l', 'o'}, 0, ErrChannelClosed})
}

func TestChannelStreamReusingWriteBuffer(t *testing.T) {
	bufA, bufB := ringbuffer.New(256), ringbuffer.New(256)

	channelChan := make(chan io.ReadWriteCloser, 1)
	m0 := New(bufA, bufB, mtu, nil)
	m1 := New(bufB, bufA, mtu, func(s io.ReadWriteCloser) { channelChan <- s })
	defer checkAllChannelsClosed(t, m0)
	defer checkAllChannelsClosed(t, m1)

	c0, _ := m0.OpenChannel()
	defer c0.Close()
	c1 := <-channelChan
	defer c1.Close()

	done := make(chan struct{})

	// Emit small chunks of incrementing bytes until there's a write error
	go func() {
		arr := [7]byte{}
		for i := 0; true; i += len(arr) {
			for j := range arr {
				arr[j] = byte(j + i)
			}
			if _, err := c0.Write(arr[:]); err != nil {
				break
			}
		}
		close(done) // Flag that this goroutine has stopped sending
	}()

	for i := 0; i < 10000; i += 5 {
		checkRead(t, c1, readCheck{
			[]byte{byte(i), byte(i + 1), byte(i + 2), byte(i + 3), byte(i + 4)},
			5, nil,
		})
	}
	// Close the communication buffers to stop the sending of data
	bufA.Close()
	bufB.Close()
	<-done
}

func openTwoChannelPairs(t *testing.T) (ma, mb *Multiplexer, p, q, r, s io.ReadWriteCloser) {
	//       ┌─────→ writerA ─────┆─────→ readerB ─────┐
	//       │                    ┆                    │
	//       │    ┌─────→ ─────→ ─┆───→ ─────→ ───┐    │
	//       │    p˚              ┆               q    │
	//       │    └──── ←───── ←──┆── ←───── ←────┘    │
	//       │                    ┆                    │
	// MultiplexerA               ┆              MultiplexerB
	//       │                    ┆                    │
	//       │    ┌─────→ ─────→ ─┆───→ ─────→ ───┐    │
	//       │    r               ┆               s˚   │
	//       │    └──── ←───── ←──┆── ←───── ←────┘    │
	//       │                    ┆                    │
	//       └───── readerA ←─────┆────── writerB ←────┘
	//
	// ˚ = locally created
	var err error
	bufAtoB, bufBtoA := ringbuffer.New(1024), ringbuffer.New(1024)

	qChan, rChan := make(chan io.ReadWriteCloser, 1), make(chan io.ReadWriteCloser, 1)
	multiplexerA := New(bufBtoA, bufAtoB, mtu, func(s io.ReadWriteCloser) { rChan <- s })
	multiplexerB := New(bufAtoB, bufBtoA, mtu, func(s io.ReadWriteCloser) { qChan <- s })

	p, err = multiplexerA.OpenChannel()
	if err != nil {
		t.Errorf("Error opening channel p: %v", err)
	}

	s, err = multiplexerB.OpenChannel()
	if err != nil {
		t.Errorf("Error opening channel s: %v", err)
	}

	q, r = <-qChan, <-rChan
	return multiplexerA, multiplexerB, p, q, r, s
}

func TestPairedReadWrites(t *testing.T) {
	ma, mb, p, q, r, s := openTwoChannelPairs(t)
	defer func() {
		p.Close()
		q.Close()
		r.Close()
		s.Close()
		checkAllChannelsClosed(t, ma)
		checkAllChannelsClosed(t, mb)
	}()

	checkWrite(t, p, writeCheck{testDataPtoQ, len(testDataPtoQ), nil})
	checkWrite(t, q, writeCheck{testDataQtoP, len(testDataQtoP), nil})
	checkWrite(t, r, writeCheck{testDataRtoS, len(testDataRtoS), nil})
	checkWrite(t, s, writeCheck{testDataStoR, len(testDataStoR), nil})

	checkRead(t, q, readCheck{testDataPtoQ, len(testDataPtoQ), nil})
	checkRead(t, p, readCheck{testDataQtoP, len(testDataQtoP), nil})
	checkRead(t, s, readCheck{testDataRtoS, len(testDataRtoS), nil})
	checkRead(t, r, readCheck{testDataStoR, len(testDataStoR), nil})
}

func TestPairedChannelCloseReadGivesEOF(t *testing.T) {
	ma, mb, p, q, r, s := openTwoChannelPairs(t)
	defer checkAllChannelsClosed(t, ma)
	defer checkAllChannelsClosed(t, mb)

	// Check closing of local channel gives EOF reads
	if err := p.Close(); err != nil {
		t.Errorf("Error closing channel p: %v", p)
	}

	checkRead(t, p, readCheck{[]byte{0}, 0, io.EOF})
	checkRead(t, q, readCheck{[]byte{0}, 0, io.EOF})

	// Check closing of remote channel gives EOF reads
	if err := r.Close(); err != nil {
		t.Errorf("Error closing channel r: %v", r)
	}

	checkRead(t, r, readCheck{[]byte{0}, 0, io.EOF})
	checkRead(t, s, readCheck{[]byte{0}, 0, io.EOF})
}

func TestPairedChannelSimultaneousCloseReadGivesEOF(t *testing.T) {
	ma, mb, p, q, r, s := openTwoChannelPairs(t)
	defer checkAllChannelsClosed(t, ma)
	defer checkAllChannelsClosed(t, mb)

	p.Close()
	q.Close()
	r.Close()
	s.Close()
	checkRead(t, r, readCheck{[]byte{0}, 0, io.EOF})
	checkRead(t, s, readCheck{[]byte{0}, 0, io.EOF})
}
