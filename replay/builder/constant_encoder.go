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

package builder

import (
	"bytes"
	eb "encoding/binary"
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

type constantEncoder struct {
	encoder     *protocol.Encoder
	buffer      *bytes.Buffer
	constantMap map[binary.ID]uint64
	data        []byte
	alignment   uint64
}

func newConstantEncoder(alignment int, byteOrder eb.ByteOrder) *constantEncoder {
	buffer := &bytes.Buffer{}
	encoder := protocol.NewEncoder(buffer, byteOrder)
	return &constantEncoder{
		encoder:     encoder,
		buffer:      buffer,
		constantMap: make(map[binary.ID]uint64),
		alignment:   uint64(alignment),
	}
}

func (e *constantEncoder) writeValues(v ...value.Value) value.Pointer {
	if len(v) == 0 {
		panic("Cannot write an empty list of values!")
	}

	e.begin()
	for _, v := range v {
		switch v := v.(type) {
		case value.Bool:
			e.encoder.Bool(bool(v))
		case value.U8:
			e.encoder.Uint8(uint8(v))
		case value.S8:
			e.encoder.Int8(int8(v))
		case value.U16:
			e.encoder.Uint16(uint16(v))
		case value.S16:
			e.encoder.Int16(int16(v))
		case value.F32:
			e.encoder.Float32(float32(v))
		case value.U32:
			e.encoder.Uint32(uint32(v))
		case value.S32:
			e.encoder.Int32(int32(v))
		case value.F64:
			e.encoder.Float64(float64(v))
		case value.U64:
			e.encoder.Uint64(uint64(v))
		case value.S64:
			e.encoder.Int64(int64(v))
		default:
			panic(fmt.Errorf("Cannot write Value %T to constant memory", v))
		}
	}
	return e.finish()
}

func (e *constantEncoder) writeString(s string) value.Pointer {
	e.begin()
	e.encoder.CString(s)
	return e.finish()
}

func (e *constantEncoder) begin() {
	e.buffer.Reset()
}

func (e *constantEncoder) finish() value.Pointer {
	data := e.buffer.Bytes()
	hash := binary.NewID(data)
	offset, found := e.constantMap[hash]
	if !found {
		offset = uint64(len(e.data))
		if offset%e.alignment != 0 {
			padding := e.alignment - offset%e.alignment
			e.data = append(e.data, make([]byte, padding)...)
			offset += padding
		}
		e.data = append(e.data, data...)
		e.constantMap[hash] = offset
	}
	return value.ConstantPointer(offset)
}
