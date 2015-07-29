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

package main

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"time"

	"android.googlesource.com/platform/tools/gpu/adb"
	"android.googlesource.com/platform/tools/gpu/gapii"
	"android.googlesource.com/platform/tools/gpu/log"
)

var (
	verbose  = flag.Bool("v", false, "verbose messages")
	device   = flag.String("device", "", "the device to capture on")
	activity = flag.String("activity", "", "the activity to launch")
	spyport  = flag.Int("i", 9286, "gapii TCP port to connect to")
	duration = flag.Duration("d", 10*time.Second, "duration to trace for")
	output   = flag.String("out", "gapit.gfxtrace", "the file to generate")
	debug    = flag.Bool("debug", false, "use the debug spy .so")
)

const usage = `gapit: A tool to trace graphics calls on android.
Usage: gapit [--out=file] <inputs>...
  -help: show this help message
`

func run() error {
	flag.Usage = func() {
		fmt.Printf(usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	info := os.Stdout
	if *verbose == false {
		info = nil
	}
	logger := log.Writer(info, os.Stdout, os.Stderr, nil)
	defer log.Close(logger)

	d, err := getDevice(logger, *device)
	if err != nil {
		return err
	}
	err = d.Root()
	switch err {
	case nil:
	case adb.ErrDeviceNotRooted:
		return err
	default:
		return fmt.Errorf("Failed to restart ADB as root: %v", err)
	}
	log.Infof(logger, "Device is rooted")
	a, err := getAction(logger, d, *activity)
	if err != nil {
		return err
	}

	log.Infof(logger, "Creating file %s", *output)
	os.MkdirAll(filepath.Dir(*output), 0755)
	file, err := os.Create(*output)
	if err != nil {
		return err
	}
	defer file.Close()

	err = gapii.AdbStart(logger, a, adb.TCPPort(*spyport), *debug)
	if err != nil {
		return err
	}

	stop := make(chan struct{})
	go func() {
		time.Sleep(*duration)
		close(stop)
	}()
	_, err = gapii.Capture(logger, *spyport, file, stop)
	if err != nil {
		return err
	}
	return nil
}

func getDevice(logger log.Logger, pattern string) (*adb.Device, error) {
	devices, err := adb.Devices()
	if err != nil {
		return nil, err
	}
	if len(devices) == 0 {
		return nil, fmt.Errorf("No devices found")
	}
	if *verbose {
		log.Infof(logger, "Device list:")
		for _, test := range devices {
			log.Infof(logger, "    %s", test.Serial)
		}
	}
	matchingDevices := []*adb.Device{}
	if pattern == "" {
		matchingDevices = devices
	} else {
		re := regexp.MustCompile("(?i)" + pattern)
		for _, test := range devices {
			if re.MatchString(test.String()) {
				matchingDevices = append(matchingDevices, test)
			}
		}
	}
	if len(matchingDevices) == 0 {
		return nil, fmt.Errorf("No devices matching %q found", pattern)
	} else if len(matchingDevices) > 1 {
		fmt.Println("Matching devices:")
		for _, test := range matchingDevices {
			fmt.Print("    ")
			fmt.Println(test.Serial)
		}
		return nil, fmt.Errorf("Multiple devices matching %q found", pattern)
	}
	log.Infof(logger, "Tracing on %s", matchingDevices[0])
	return matchingDevices[0], nil
}

func getAction(logger log.Logger, d *adb.Device, pattern string) (*adb.Action, error) {
	re := regexp.MustCompile("(?i)" + pattern)
	packages, err := d.InstalledPackages()
	if err != nil {
		return nil, err
	}
	if len(packages) == 0 {
		return nil, fmt.Errorf("No packages found")
	}
	matchingActions := []*adb.Action{}
	for _, p := range packages {
		for _, action := range p.Actions {
			if re.MatchString(action.String()) {
				matchingActions = append(matchingActions, action)
			}
		}
	}
	if len(matchingActions) == 0 {
		return nil, fmt.Errorf("No actions matching %s found", pattern)
	} else if len(matchingActions) > 1 {
		fmt.Println("Matching actions:")
		for _, test := range matchingActions {
			fmt.Print("    ")
			fmt.Println(test)
		}
		return nil, fmt.Errorf("Multiple actions matching %q found", pattern)
	}
	log.Infof(logger, "Action %s", matchingActions[0])
	return matchingActions[0], nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "gapit failed: %v\n", err)
		os.Exit(1)
	}
}
