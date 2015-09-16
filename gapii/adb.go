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

package gapii

import (
	"fmt"
	"os"
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/adb"
	"android.googlesource.com/platform/tools/gpu/log"
)

const preloadPath = "/data/gapii.so"

var abiToSo = map[string]string{
	"armeabi":     "android-arm",
	"armeabi-v7a": "android-arm",
	"arm64-v8a":   "android-arm64",
}

func getSoName(a *adb.Action) (string, error) {
	abi := a.Package.ABI
	if abi == "" {
		abi = a.Package.Device.Abi()
	}
	so, ok := abiToSo[abi]
	if !ok {
		return "", fmt.Errorf("Unknown device abi: %s", abi)
	}
	return so, nil
}

func getSoPath(a *adb.Action, debug bool) (string, error) {
	// TODO: decide how we are going to find the so's
	gopath := filepath.SplitList(os.Getenv("GOPATH"))[0]
	soName, err := getSoName(a)
	if err != nil {
		return "", err
	}
	flavor := "release"
	if debug {
		flavor = "debug"
	}
	return filepath.Join(gopath, "bin", soName, flavor, "libgapii.so"), nil
}

// AdbStart launches an activity on an android device with the gapii tracer
// enabled. Gapii will attempt to connect back on the specified host port to
// write the trace.
func AdbStart(l log.Logger, a *adb.Action, spyport adb.Port, debug bool) error {
	p := a.Package
	d := p.Device
	enforced, err := d.SELinuxEnforcing()
	if err != nil {
		log.Errorf(l, "Failed detecting se linux enforcing state: %s", err)
		return err
	}
	if enforced {
		log.Infof(l, "Disabling SELinux enforcing")
		err := d.SetSELinuxEnforcing(false)
		if err != nil {
			log.Errorf(l, "Failed disabling se linux enforcing: %s", err)
			return err
		}
		defer func() {
			log.Infof(l, "Re-enabling SELinux enforcing")
			err := d.SetSELinuxEnforcing(false)
			if err != nil {
				log.Errorf(l, "Failed re-enabling se linux enforcing: %s", err)
			}
		}()
	}

	gapiiPath, err := getSoPath(a, debug)
	if err != nil {
		log.Errorf(l, "Failed finding gapii: %s", err)
		return err
	}
	log.Infof(l, "Pushing %s to %s", gapiiPath, preloadPath)
	err = d.Push(gapiiPath, preloadPath)
	if err != nil {
		log.Errorf(l, "Failed pushing %s to %s: %s", gapiiPath, preloadPath, err)
		return err
	}

	log.Infof(l, "Setting LD_PRELOAD on %s", p.Name)
	err = p.SetWrapProperties("LD_PRELOAD=" + preloadPath)
	if err != nil {
		log.Errorf(l, "Failed setting LD_PRELOAD: %s", err)
		return err
	}
	defer func() {
		log.Infof(l, "Clearing LD_PRELOAD on %s", p.Name)
		err := p.SetWrapProperties("")
		if err != nil {
			log.Errorf(l, "Failed clearing LD_PRELOAD: %s", err)
			return
		}
	}()

	log.Infof(l, "Forwarding port %v", spyport)
	err = d.Forward(spyport, adb.NamedAbstractSocket("gapii"))
	if err != nil {
		log.Errorf(l, "Failed setting up port forwarding: %s", err)
		return err
	}

	log.Infof(l, "Starting activity %s/%s", p.Name, a.Activity)
	d.StartActivity(*a)
	if err != nil {
		log.Errorf(l, "Failed starting %s/%s: %s", p.Name, a.Activity, err)
		return err
	}
	return nil
}
