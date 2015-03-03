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

package binary

import (
	"bytes"
	"reflect"

	"math/rand"
	"testing"
)

func prepare(a interface{}) (*Encoder, *Decoder) {
	rand.Seed(1)
	s := reflect.ValueOf(a)
	maxSize := 0
	for i := 0; i < s.Len(); i++ {
		e := s.Index(i)
		switch e.Kind() {
		case reflect.Int8:
			e.SetInt(rand.Int63())
			maxSize = 2
		case reflect.Uint8:
			e.SetUint(uint64(rand.Int63()))
			maxSize = 2
		case reflect.Int16:
			e.SetInt(rand.Int63())
			maxSize = 3
		case reflect.Uint16:
			e.SetUint(uint64(rand.Int63()))
			maxSize = 3
		case reflect.Int32:
			e.SetInt(rand.Int63())
			maxSize = 5
		case reflect.Uint32:
			e.SetUint(uint64(rand.Int63()))
			maxSize = 5
		case reflect.Int64:
			e.SetInt(int64((uint64(rand.Uint32()) << 32) | uint64(rand.Uint32())))
			maxSize = 9
		case reflect.Uint64:
			e.SetUint((uint64(rand.Uint32()) << 32) | uint64(rand.Uint32()))
			maxSize = 9
		}
	}
	// build a big enough buffer, and wrap it in coders
	buf := &bytes.Buffer{}
	buf.Grow(s.Len() * maxSize)
	return NewEncoder(buf), NewDecoder(buf)
}

func TestRoundTripInt8(t *testing.T) {
	values := make([]int8, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Int8(v)
	}
	for _, v := range values {
		got, _ := d.Int8()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func TestRoundTripUint8(t *testing.T) {
	values := make([]uint8, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Uint8(v)
	}
	for _, v := range values {
		got, _ := d.Uint8()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func TestRoundTripInt16(t *testing.T) {
	values := make([]int16, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Int16(v)
	}
	for _, v := range values {
		got, _ := d.Int16()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func TestRoundTripUint16i(t *testing.T) {
	values := make([]uint16, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Uint16(v)
	}
	for _, v := range values {
		got, _ := d.Uint16()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func TestRoundTripInt32(t *testing.T) {
	values := make([]int32, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Int32(v)
	}
	for _, v := range values {
		got, _ := d.Int32()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func TestRoundTripUint32(t *testing.T) {
	values := make([]uint32, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Uint32(v)
	}
	for _, v := range values {
		got, _ := d.Uint32()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func TestRoundTripInt64(t *testing.T) {
	values := make([]int64, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Int64(v)
	}
	for _, v := range values {
		got, _ := d.Int64()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func TestRoundTripUint64(t *testing.T) {
	values := make([]uint64, 1000)
	e, d := prepare(values)
	for _, v := range values {
		e.Uint64(v)
	}
	for _, v := range values {
		got, _ := d.Uint64()
		if v != got {
			t.Errorf("Round trip gave wrong value. Expected: %x, got: %x", v, got)
		}
	}
}

func BenchmarkInt8(b *testing.B) {
	values := make([]int8, 1000000)
	e, d := prepare(values)
	b.ResetTimer()
	for _, v := range values {
		e.Int8(v)
	}
	for _ = range values {
		d.Int8()
	}
}
