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

func doReplay(t *testing.T, f func(*builder.Builder)) {
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

	err = executor.Execute(payload, decoder, connection, d, l, arch)
	if err != nil {
		t.Errorf("Executor failed with error: %v", err)
	}
}

func TestPostbackString(t *testing.T) {
	expected := "γειά σου κόσμος"

	done := make(chan struct{})

	doReplay(t, func(b *builder.Builder) {
		ptr := b.String(expected)
		b.Post(ptr, uint64(len(expected)), func(d binary.Decoder, err error) error {
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			data := make([]byte, len(expected))
			err = d.Data(data)
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			if expected != string(data) {
				t.Errorf("Postback data was not as expected. Expected: %v. Got: %v", expected, data)
			}
			close(done)
			return err
		})
	})

	<-done
}

func TestMultiPostback(t *testing.T) {
	done := make(chan struct{})

	doReplay(t, func(b *builder.Builder) {
		ptr := b.AllocateTemporaryMemory(8)
		b.Push(value.Bool(false))
		b.Store(ptr)
		b.Post(ptr, 1, func(d binary.Decoder, err error) error {
			expected := false
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			data, err := d.Bool()
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			if !reflect.DeepEqual(expected, data) {
				t.Errorf("Postback data was not as expected. Expected: %v. Got: %v", expected, data)
			}
			return err
		})

		b.Push(value.Bool(true))
		b.Store(ptr)
		b.Post(ptr, 1, func(d binary.Decoder, err error) error {
			expected := true
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			data, err := d.Bool()
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			if !reflect.DeepEqual(expected, data) {
				t.Errorf("Postback data was not as expected. Expected: %v. Got: %v", expected, data)
			}
			return err
		})

		b.Push(value.F64(123.456))
		b.Store(ptr)
		b.Post(ptr, 8, func(d binary.Decoder, err error) error {
			expected := float64(123.456)
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			data, err := d.Float64()
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			if !reflect.DeepEqual(expected, data) {
				t.Errorf("Postback data was not as expected. Expected: %v. Got: %v", expected, data)
			}
			close(done)
			return err
		})
	})

	<-done
}
