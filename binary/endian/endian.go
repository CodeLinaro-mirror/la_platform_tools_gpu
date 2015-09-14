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
	"math"

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
	err       error
}

type writer struct {
	writer    io.Writer
	tmp       [8]byte
	byteOrder ByteOrder
	err       error
}

func (r *reader) earlierError() error {
	return fmt.Errorf("Reading was stopped due to an earlier error: %v", r.err)
}

func (w *writer) earlierError() error {
	return fmt.Errorf("Writing was stopped due to an earlier error: %v", w.err)
}

func (r *reader) Data(p []byte) error {
	if r.err != nil {
		return r.earlierError()
	}
	n, err := io.ReadFull(r.reader, p)
	if err != nil {
		r.err = err
		err = fmt.Errorf("%v after reading %d bytes", err, n)
	}
	return err
}

func (w *writer) Data(data []byte) error {
	if w.err != nil {
		return w.earlierError()
	}
	n, err := w.writer.Write(data)
	if err != nil {
		w.err = err
		return err
	}
	if n != len(data) {
		w.err = io.ErrShortWrite
		return io.ErrShortWrite
	}
	return nil
}

func (r *reader) Bool() (bool, error) {
	if r.err != nil {
		return false, r.earlierError()
	}
	b, err := r.Uint8()
	return b != 0, err
}

func (w *writer) Bool(v bool) error {
	if w.err != nil {
		return w.earlierError()
	}
	if v {
		return w.Uint8(1)
	}
	return w.Uint8(0)
}

func (r *reader) Int8() (int8, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	i, err := r.Uint8()
	return int8(i), err
}

func (w *writer) Int8(v int8) error {
	if w.err != nil {
		return w.earlierError()
	}
	return w.Uint8(uint8(v))
}

func (r *reader) Uint8() (uint8, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	b := r.tmp[:1]
	_, err := io.ReadFull(r.reader, b[:1])
	r.err = err
	return b[0], err
}

func (w *writer) Uint8(v uint8) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.tmp[0] = v
	return w.Data(w.tmp[:1])
}

func (r *reader) Int16() (int16, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:2])
	r.err = err
	return int16(r.byteOrder.Uint16(r.tmp[:])), err
}

func (w *writer) Int16(v int16) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint16(w.tmp[:], uint16(v))
	_, err := w.writer.Write(w.tmp[:2])
	w.err = err
	return err
}

func (r *reader) Uint16() (uint16, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:2])
	r.err = err
	return r.byteOrder.Uint16(r.tmp[:]), err
}

func (w *writer) Uint16(v uint16) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint16(w.tmp[:], v)
	_, err := w.writer.Write(w.tmp[:2])
	w.err = err
	return err
}

func (r *reader) Int32() (int32, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:4])
	r.err = err
	return int32(r.byteOrder.Uint32(r.tmp[:])), err
}

func (w *writer) Int32(v int32) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint32(w.tmp[:], uint32(v))
	_, err := w.writer.Write(w.tmp[:4])
	w.err = err
	return err
}

func (r *reader) Uint32() (uint32, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:4])
	r.err = err
	return r.byteOrder.Uint32(r.tmp[:]), err
}

func (w *writer) Uint32(v uint32) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint32(w.tmp[:], v)
	_, err := w.writer.Write(w.tmp[:4])
	w.err = err
	return err
}

func (r *reader) Int64() (int64, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:8])
	r.err = err
	return int64(r.byteOrder.Uint64(r.tmp[:])), err
}

func (w *writer) Int64(v int64) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint64(w.tmp[:], uint64(v))
	_, err := w.writer.Write(w.tmp[:8])
	w.err = err
	return err
}

func (r *reader) Uint64() (uint64, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:8])
	r.err = err
	return r.byteOrder.Uint64(r.tmp[:]), err
}

func (w *writer) Uint64(v uint64) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint64(w.tmp[:], v)
	_, err := w.writer.Write(w.tmp[:8])
	w.err = err
	return err
}

func (r *reader) Float32() (float32, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:4])
	r.err = err
	return math.Float32frombits(r.byteOrder.Uint32(r.tmp[:])), err
}

func (w *writer) Float32(v float32) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint32(w.tmp[:], math.Float32bits(v))
	_, err := w.writer.Write(w.tmp[:4])
	w.err = err
	return err
}

func (r *reader) Float64() (float64, error) {
	if r.err != nil {
		return 0, r.earlierError()
	}
	_, err := io.ReadFull(r.reader, r.tmp[:8])
	r.err = err
	return math.Float64frombits(r.byteOrder.Uint64(r.tmp[:])), err
}

func (w *writer) Float64(v float64) error {
	if w.err != nil {
		return w.earlierError()
	}
	w.byteOrder.PutUint64(w.tmp[:], math.Float64bits(v))
	_, err := w.writer.Write(w.tmp[:8])
	w.err = err
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

func (w *writer) String(v string) error {
	if w.err != nil {
		return w.earlierError()
	}
	if _, err := w.writer.Write([]byte(v)); err != nil {
		w.err = err
		return err
	}
	return w.Uint8(0)
}

func (w *writer) Error() error {
	return w.err
}

func (r *reader) Error() error {
	return r.err
}

func (r *reader) SetError(err error) error {
	if r.err != nil {
		err = fmt.Errorf("Error %v whilst in error state %v", err, r.err)
	}
	r.err = err
	return err
}

func (w *writer) SetError(err error) error {
	if w.err != nil {
		err = fmt.Errorf("Error %v whilst in error state %v", err, w.err)
	}
	w.err = err
	return err
}
