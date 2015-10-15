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
	"flag"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/client/gapir"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/executor"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

var (
	port   = flag.Int("gapir", 0, "The port to connect to gapir on, 0 means start new instance")
	d      = database.NewInMemory(nil)
	device gapir.Device
)

func TestMain(m *testing.M) {
	flag.Parse()
	if *port > 0 {
		device = gapir.NewDevice(gapir.LocalName, *port, log.Std())
	} else {
		device = gapir.RunLocal(log.Std())
	}
	code := m.Run()
	atexit.Exit(code)
}

func doReplay(t *testing.T, f func(*builder.Builder)) {
	l := log.Testing(t)

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
			defer close(done)
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			data := make([]byte, len(expected))
			d.Data(data)
			err = d.Error()
			if err != nil {
				t.Errorf("Postback returned error: %v", err)
				return err
			}
			if expected != string(data) {
				t.Errorf("Postback data was not as expected. Expected: %v. Got: %v", expected, data)
			}
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
			data := d.Bool()
			err = d.Error()
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
			data := d.Bool()
			err = d.Error()
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
			data := d.Float64()
			err = d.Error()
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
