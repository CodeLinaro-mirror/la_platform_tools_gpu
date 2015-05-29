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
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/integration/replay/utils"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/executor"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

func doReplay(t *testing.T, f func(*builder.Builder), handlers executor.PostbackHandlerMap) {
	d, l := database.InMemory(), log.Testing(t)

	mgr := replay.New(d, l)
	device := utils.FindLocalDevice(t, mgr)
	arch := device.Info().Architecture()

	connection, err := device.Connect()
	if err != nil {
		t.Errorf("Failed to connect to '%s': %v", device.Info().Name, err)
		return
	}

	b := builder.New(arch)

	f(b)

	payload, decoder, err := b.Build(l)
	if err != nil {
		t.Errorf("Build failed with error: %v", err)
	}

	err = executor.Execute(payload, decoder, connection, d, l, handlers, arch)
	if err != nil {
		t.Errorf("Executor failed with error: %v", err)
	}
}

func checkPostback(t *testing.T, expected interface{}, then ...func()) func(data interface{}, err error) {
	return func(data interface{}, err error) {
		if err != nil {
			t.Errorf("Postback returned error: %v", err)
		}
		if !reflect.DeepEqual(expected, data) {
			t.Errorf("Postback data was not as expected. Expected: %v. Got: %v", expected, data)
		}
		for _, f := range then {
			f()
		}
	}
}

func TestPostbackString(t *testing.T) {
	expected := "γειά σου κόσμος"

	done := make(chan struct{})

	doReplay(t, func(b *builder.Builder) {
		ptr := b.String(expected)
		b.Post(ptr, uint64(len(expected)), 0, func(d binary.Decoder) (interface{}, error) {
			buf := make([]byte, len(expected))
			return buf, d.Data(buf)
		})
	}, executor.PostbackHandlerMap{
		0: checkPostback(t, []byte(expected), func() { close(done) }),
	})

	<-done
}

func TestMultiPostback(t *testing.T) {
	done := make(chan struct{})

	doReplay(t, func(b *builder.Builder) {
		ptr := b.AllocateTemporaryMemory(8)
		b.Push(value.Bool(false))
		b.Store(ptr)
		b.Post(ptr, 1, 100, func(d binary.Decoder) (interface{}, error) { return d.Bool() })

		b.Push(value.Bool(true))
		b.Store(ptr)
		b.Post(ptr, 1, 200, func(d binary.Decoder) (interface{}, error) { return d.Bool() })

		b.Push(value.F64(123.456))
		b.Store(ptr)
		b.Post(ptr, 8, 300, func(d binary.Decoder) (interface{}, error) { return d.Float64() })
	}, executor.PostbackHandlerMap{
		100: checkPostback(t, false),
		200: checkPostback(t, true),
		300: checkPostback(t, 123.456, func() { close(done) }),
	})

	<-done
}
