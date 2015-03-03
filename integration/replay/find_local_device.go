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
	"testing"
	"time"

	"android.googlesource.com/platform/tools/gpu/replay"
)

const findLocalDeviceAttempts = 5
const msBetweenFindLocalDeviceAttempts = 500
const localDeviceName = "Local machine" // TODO: Remove hard-coded string.

// findLocalDevice returns the replay Device for the local host. If the local
// host cannot be found then the test fails and nil is returned.
func findLocalDevice(t *testing.T, mgr *replay.Manager) replay.Device {
	for i := 0; i < findLocalDeviceAttempts; i++ {
		for _, d := range mgr.Devices() {
			info := d.Info()
			t.Logf("Found device: '%s'", info.Name)
			if info.Name == localDeviceName {
				return d
			}
		}
		time.Sleep(time.Millisecond * msBetweenFindLocalDeviceAttempts)
	}
	t.Fatalf("Local replay device not found")
	return nil
}
