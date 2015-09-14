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

package vle

import (
	"fmt"
	"io"
	"math"

	"android.googlesource.com/platform/tools/gpu/binary"
)

// Reader creates a binary.Reader that reads from the provided io.Reader.
func Reader(r io.Reader) binary.Reader {
	return &reader{reader: r}
}

// Writer creates a binary.Writer that writes to the supplied io.Writer.
func Writer(w io.Writer) binary.Writer {
	return &writer{writer: w}
}

type reader struct {
	reader io.Reader
	tmp    [9]byte
	err    error
}

type writer struct {
	writer io.Writer
	tmp    [9]byte
	err    error
}

func shuffle32(v uint32) uint32 {
	return 0 |
		((v & 0x000000ff) << 24) |
		((v & 0x0000ff00) << 8) |
		((v & 0x00ff0000) >> 8) |
		((v & 0xff000000) >> 24)
}

func shuffle64(v uint64) uint64 {
	return 0 |
		((v & 0x00000000000000ff) << 56) |
		((v & 0x000000000000ff00) << 40) |
		((v & 0x0000000000ff0000) << 24) |
		((v & 0x00000000ff000000) << 8) |
		((v & 0x000000ff00000000) >> 8) |
		((v & 0x0000ff0000000000) >> 24) |
		((v & 0x00ff000000000000) >> 40) |
		((v & 0xff00000000000000) >> 56)
}

func (r *reader) intv() (int64, error) {
	uv, err := r.uintv()
	v := int64(uv >> 1)
	if uv&1 != 0 {
		v = ^v
	}
	return v, err
}

func (w *writer) intv(v int64) error {
	uv := uint64(v) << 1
	if v < 0 {
		uv = ^uv
	}
	return w.uintv(uv)
}

func (r *reader) uintv() (uint64, error) {
	tag, err := r.Uint8()
	if err != nil {
		return 0, err
	}
	count := uint(0)
	for ; ((0x80 >> count) & tag) != 0; count++ {
	}
	v := uint64(tag & (byte(0xff) >> count))
	if count == 0 {
		return v, nil
	}
	if err := r.Data(r.tmp[:count]); err != nil {
		return 0, err
	}
	for i := uint(0); i < count; i++ {
		v = (v << 8) | uint64(r.tmp[i])
	}
	return v, nil
}

func (w *writer) uintv(v uint64) error {
	space := uint64(0x7f)
	tag := byte(0)
	for o := 8; true; o-- {
		if v <= space {
			w.tmp[o] = byte(v) | byte(tag)
			return w.Data(w.tmp[o:])
		}
		w.tmp[o] = byte(v)
		v >>= 8
		space >>= 1
		tag = (tag >> 1) | 0x80
	}
	panic("Cannot get here")
}

func (r *reader) Data(p []byte) error {
	if r.err != nil {
		return fmt.Errorf("Reading was stopped due to an earlier error: %v", r.err)
	}
	_, err := io.ReadFull(r.reader, p)
	r.err = err
	return err
}

func (w *writer) Data(data []byte) error {
	if w.err != nil {
		return fmt.Errorf("Writing was stopped due to an earlier error: %v", w.err)
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
	if r.err != nil {
		return 0, fmt.Errorf("Reading was stopped due to an earlier error: %v", r.err)
	}
	b := r.tmp[:1]
	_, err := io.ReadFull(r.reader, b[:1])
	r.err = err
	return b[0], err
}

func (w *writer) Uint8(v uint8) error {
	w.tmp[0] = v
	return w.Data(w.tmp[:1])
}

func (r *reader) Int16() (int16, error)   { v, err := r.intv(); return int16(v), err }
func (w *writer) Int16(v int16) error     { return w.intv(int64(v)) }
func (r *reader) Uint16() (uint16, error) { v, err := r.uintv(); return uint16(v), err }
func (w *writer) Uint16(v uint16) error   { return w.uintv(uint64(v)) }
func (r *reader) Int32() (int32, error)   { v, err := r.intv(); return int32(v), err }
func (w *writer) Int32(v int32) error     { return w.intv(int64(v)) }
func (r *reader) Uint32() (uint32, error) { v, err := r.uintv(); return uint32(v), err }
func (w *writer) Uint32(v uint32) error   { return w.uintv(uint64(v)) }
func (r *reader) Int64() (int64, error)   { return r.intv() }
func (w *writer) Int64(v int64) error     { return w.intv(v) }
func (r *reader) Uint64() (uint64, error) { return r.uintv() }
func (w *writer) Uint64(v uint64) error   { return w.uintv(v) }

func (r *reader) Float32() (float32, error) {
	bits, err := r.Uint32()
	return math.Float32frombits(shuffle32(bits)), err
}

func (w *writer) Float32(v float32) error {
	return w.Uint32(shuffle32(math.Float32bits(v)))
}

func (r *reader) Float64() (float64, error) {
	bits, err := r.Uint64()
	return math.Float64frombits(shuffle64(bits)), err
}

func (w *writer) Float64(v float64) error {
	return w.Uint64(shuffle64(math.Float64bits(v)))
}

func (r *reader) String() (string, error) {
	c, err := r.Uint32()
	if err != nil || c == 0 {
		return "", err
	}
	s := make([]byte, c)
	err = r.Data(s)
	return string(s), err
}

func (w *writer) String(v string) error {
	if err := w.Uint32(uint32(len(v))); err != nil {
		return err
	}
	return w.Data([]byte(v))
}

func (r *reader) Error() error {
	return r.err
}

func (w *writer) Error() error {
	return w.err
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
