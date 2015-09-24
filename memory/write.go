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

package memory

import (
	"fmt"
	"reflect"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/device"
)

// Write writes the value v to the writer w. If v is an array or slice, then
// each of the elements will be written, sequentially.
func Write(w binary.Writer, arch device.Architecture, v interface{}) error {
	switch v := v.(type) {
	case Pointer:
		binary.WriteUint(w, arch.PointerSize*8, v.Address)
		return w.Error()
	}

	r := reflect.ValueOf(v)
	t := r.Type()
	switch t.Kind() {
	case reflect.Float32:
		w.Float32(float32(r.Float()))

	case reflect.Float64:
		w.Float64(r.Float())

	case reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		binary.WriteInt(w, t.Bits(), r.Int())

	case reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		binary.WriteUint(w, t.Bits(), r.Uint())

	case reflect.Int:
		binary.WriteInt(w, arch.IntegerSize*8, r.Int())

	case reflect.Uint:
		binary.WriteUint(w, arch.IntegerSize*8, r.Uint())

	case reflect.Array, reflect.Slice:
		for i := 0; i < r.Len(); i++ {
			Write(w, arch, r.Index(i).Interface())
		}

	case reflect.String:
		w.String(r.String())

	case reflect.Bool:
		w.Bool(r.Bool())

	default:
		return fmt.Errorf("Cannot write type: %s", t.Name())
	}
	return w.Error()
}
