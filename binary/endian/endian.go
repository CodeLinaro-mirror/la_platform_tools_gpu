// Copyright (C) 2014 The Android Open Source Project
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

package endian

import (
	eb "encoding/binary"
	"fmt"
	"io"
	"io/ioutil"
	"math"
	"os"

	"android.googlesource.com/platform/tools/gpu/binary"
)

type ByteOrder eb.ByteOrder

var (
	Little = ByteOrder(eb.LittleEndian)
	Big    = ByteOrder(eb.BigEndian)
)

// Reader creates a binary.Reader that reads from the provided io.Reader, with the
// specified byte order.
func Reader(r io.Reader, byteOrder ByteOrder) binary.Reader {
	return &reader{reader: r, byteOrder: byteOrder}
}

// Writer creates a binary.Writer that writes to the supplied stream, with the
// specified byte order.
func Writer(w io.Writer, byteOrder ByteOrder) binary.Writer {
	return &writer{writer: w, byteOrder: byteOrder}
}

type reader struct {
	reader    io.Reader
	tmp       [8]byte
	byteOrder ByteOrder
}

type writer struct {
	writer    io.Writer
	tmp       [8]byte
	byteOrder ByteOrder
}

func (r *reader) Data(p []byte) error {
	n, err := io.ReadFull(r.reader, p)
	if err != nil {
		err = fmt.Errorf("%v after reading %d bytes", err, n)
	}
	return err
}

func (w *writer) Data(data []byte) error {
	n, err := w.writer.Write(data)
	if err != nil {
		return err
	}
	if n != len(data) {
		return io.ErrShortWrite
	}
	return nil
}

func (r *reader) Skip(count uint32) error {
	if s, ok := r.reader.(io.Seeker); ok {
		_, err := s.Seek(int64(count), os.SEEK_CUR)
		return err
	}
	_, err := io.CopyN(ioutil.Discard, r.reader, int64(count))
	return err
}

func (r *reader) Bool() (bool, error) {
	b, err := r.Uint8()
	return b != 0, err
}

func (w *writer) Bool(v bool) error {
	if v {
		return w.Uint8(1)
	}
	return w.Uint8(0)
}

func (r *reader) Int8() (int8, error) {
	i, err := r.Uint8()
	return int8(i), err
}

func (w *writer) Int8(v int8) error {
	return w.Uint8(uint8(v))
}

func (r *reader) Uint8() (uint8, error) {
	b := r.tmp[:1]
	_, err := io.ReadFull(r.reader, b[:1])
	return b[0], err
}

func (w *writer) Uint8(v uint8) error {
	w.tmp[0] = v
	return w.Data(w.tmp[:1])
}

func (r *reader) Int16() (int16, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:2])
	return int16(r.byteOrder.Uint16(r.tmp[:])), err
}

func (w *writer) Int16(v int16) error {
	w.byteOrder.PutUint16(w.tmp[:], uint16(v))
	_, err := w.writer.Write(w.tmp[:2])
	return err
}

func (r *reader) Uint16() (uint16, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:2])
	return r.byteOrder.Uint16(r.tmp[:]), err
}

func (w *writer) Uint16(v uint16) error {
	w.byteOrder.PutUint16(w.tmp[:], v)
	_, err := w.writer.Write(w.tmp[:2])
	return err
}

func (r *reader) Int32() (int32, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:4])
	return int32(r.byteOrder.Uint32(r.tmp[:])), err
}

func (w *writer) Int32(v int32) error {
	w.byteOrder.PutUint32(w.tmp[:], uint32(v))
	_, err := w.writer.Write(w.tmp[:4])
	return err
}

func (r *reader) Uint32() (uint32, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:4])
	return r.byteOrder.Uint32(r.tmp[:]), err
}

func (w *writer) Uint32(v uint32) error {
	w.byteOrder.PutUint32(w.tmp[:], v)
	_, err := w.writer.Write(w.tmp[:4])
	return err
}

func (r *reader) Int64() (int64, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:8])
	return int64(r.byteOrder.Uint64(r.tmp[:])), err
}

func (w *writer) Int64(v int64) error {
	w.byteOrder.PutUint64(w.tmp[:], uint64(v))
	_, err := w.writer.Write(w.tmp[:8])
	return err
}

func (r *reader) Uint64() (uint64, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:8])
	return r.byteOrder.Uint64(r.tmp[:]), err
}

func (w *writer) Uint64(v uint64) error {
	w.byteOrder.PutUint64(w.tmp[:], v)
	_, err := w.writer.Write(w.tmp[:8])
	return err
}

func (r *reader) Float32() (float32, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:4])
	return math.Float32frombits(r.byteOrder.Uint32(r.tmp[:])), err
}

func (w *writer) Float32(v float32) error {
	w.byteOrder.PutUint32(w.tmp[:], math.Float32bits(v))
	_, err := w.writer.Write(w.tmp[:4])
	return err
}

func (r *reader) Float64() (float64, error) {
	_, err := io.ReadFull(r.reader, r.tmp[:8])
	return math.Float64frombits(r.byteOrder.Uint64(r.tmp[:])), err
}

func (w *writer) Float64(v float64) error {
	w.byteOrder.PutUint64(w.tmp[:], math.Float64bits(v))
	_, err := w.writer.Write(w.tmp[:8])
	return err
}

func (r *reader) String() (string, error) {
	s := []byte{}
	for {
		if c, err := r.Uint8(); err != nil {
			return "", err
		} else if c == 0 {
			break
		} else {
			s = append(s, c)
		}
	}
	return string(s), nil
}

func (r *reader) SkipString() error {
	for {
		if c, err := r.Uint8(); err != nil {
			return err
		} else if c == 0 {
			break
		}
	}
	return nil
}

func (w *writer) String(v string) error {
	if _, err := w.writer.Write([]byte(v)); err != nil {
		return err
	}
	return w.Uint8(0)
}
