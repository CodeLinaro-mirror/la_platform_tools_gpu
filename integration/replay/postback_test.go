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

package replay

import (
	"encoding/binary"
	"io"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/integration/replay/utils"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/executor"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func doReplay(t *testing.T, f func(*builder.Builder), handlers executor.PostbackHandlerMap) {
	db, logger := utils.NewInMemoryDatabase(), log.Testing(t)

	mgr := replay.New(db, logger)
	device := utils.FindLocalDevice(t, mgr)

	connection, err := device.Connect()
	if err != nil {
		t.Errorf("Failed to connect to '%s': %v", device.Info().Name, err)
		return
	}

	info := device.Info()
	b := builder.New(int(info.PointerSize), int(info.PointerAlignment), device.ByteOrder())

	f(b)

	payload, decoder := b.Build(logger)
	err = executor.Execute(payload, decoder, connection, db, logger, handlers, binary.LittleEndian)
	if err != nil {
		t.Errorf("Executor failed with error: %v", err)
	}
}

func checkPostback(t *testing.T, expected, got interface{}, err error) {
	if err != nil {
		t.Errorf("Postback returned error: %v", err)
	}
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Postback data was not as expected. Expected: %v. Got: %v", expected, got)
	}
}

func TestPostbackString(t *testing.T) {
	expected := "γειά σου κόσμος"

	done := make(chan struct{})

	doReplay(t, func(b *builder.Builder) {
		ptr := b.String(expected)
		b.Post(ptr, uint64(len(expected)), 0, func(d *protocol.Decoder) (interface{}, error) {
			buf := make([]byte, len(expected))
			_, err := io.ReadFull(d, buf)
			return buf, err
		})
	}, executor.PostbackHandlerMap{
		0: func(data interface{}, err error) {
			checkPostback(t, []byte(expected), data, err)
			close(done)
		},
	})

	<-done
}

func TestMultiPostback(t *testing.T) {
	done := make(chan struct{})

	doReplay(t, func(b *builder.Builder) {
		ptr := b.AllocateTemporaryMemory(8)
		b.Push(value.Bool(false))
		b.Store(ptr)
		b.Post(ptr, 1, 100, func(d *protocol.Decoder) (interface{}, error) { return d.Bool() })

		b.Push(value.Bool(true))
		b.Store(ptr)
		b.Post(ptr, 1, 200, func(d *protocol.Decoder) (interface{}, error) { return d.Bool() })

		b.Push(value.F64(123.456))
		b.Store(ptr)
		b.Post(ptr, 8, 300, func(d *protocol.Decoder) (interface{}, error) { return d.Float64() })
	}, executor.PostbackHandlerMap{
		100: func(data interface{}, err error) { checkPostback(t, false, data, err) },
		200: func(data interface{}, err error) { checkPostback(t, true, data, err) },
		300: func(data interface{}, err error) { checkPostback(t, 123.456, data, err); close(done) },
	})

	<-done
}
