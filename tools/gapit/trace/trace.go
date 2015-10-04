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

package trace

import (
	"flag"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strings"
	"time"

	"android.googlesource.com/platform/tools/gpu/adb"
	"android.googlesource.com/platform/tools/gpu/gapii"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/tools/verbs"
)

var (
	verb = &verbs.Verb{
		Name:      "trace",
		ShortHelp: "Captures a gfx trace from an application",
	}
	device   = verb.Flags.String("device", "", "the device to capture on")
	spyport  = verb.Flags.Int("i", 9286, "gapii TCP port to connect to")
	duration = verb.Flags.Duration("d", 0, "duration to trace for")
	output   = verb.Flags.String("out", "", "the file to generate")
	debug    = verb.Flags.Bool("debug", false, "use the debug spy .so")
	local    = verb.Flags.Bool("local", false, "capture a local program instead of using ADB")
	observe  = verb.Flags.String("observe", "", "comma-seperated list of points to capture the framebuffer [frame, draw]")
)

const usage = `gapit: A tool to trace graphics calls on android.
Usage: gapit [--out=file] <activity>
  -help: show this help message
`

func init() {
	verb.Run = doTrace
	verbs.Register(verb)
}

func doTrace(flags flag.FlagSet) error {
	if flags.NArg() != 1 && !*local {
		return verbs.Usage("Invalid number of arguments. Expected 1, got %d", flags.NArg())
	}

	info := os.Stdout
	if verbs.Verbosity > 0 {
		info = nil
	}

	options := gapii.Options{}
	if *observe != "" {
		for _, o := range strings.Split(*observe, ",") {
			o = strings.TrimSpace(o)
			switch o {
			case "draws", "draw", "d":
				options.ObserveFramebufferOnDrawCall = true
			case "frames", "frame", "f":
				options.ObserveFramebufferOnEOF = true
			default:
				return fmt.Errorf("Unknown observation type %s", o)
			}
		}
	}

	logger := log.Writer(info, os.Stdout, os.Stderr, nil)
	defer log.Close(logger)

	if *local {
		return captureLocal(flags, logger, options)
	} else {
		return captureADB(flags, logger, options)
	}
}

func captureLocal(flags flag.FlagSet, logger log.Logger, options gapii.Options) error {
	out := *output
	if out == "" {
		out = "capture.gfxtrace"
	}
	return capture(logger, options, out)
}

func captureADB(flags flag.FlagSet, logger log.Logger, options gapii.Options) error {
	activity := flags.Arg(0)
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
	a, err := getAction(logger, d, activity)
	if err != nil {
		return err
	}

	out := *output
	if out == "" { // No name specified? Use package name.
		name := a.Package.Name
		if i := strings.LastIndex(name, "."); i > 0 { // trim namespace
			name = name[i+1:]
		}
		out = name + ".gfxtrace"
	}

	err = gapii.AdbStart(logger, a, adb.TCPPort(*spyport), *debug)
	if err != nil {
		return err
	}

	return capture(logger, options, out)
}

func capture(logger log.Logger, options gapii.Options, out string) error {
	log.Infof(logger, "Creating file %s", out)
	os.MkdirAll(filepath.Dir(out), 0755)
	file, err := os.Create(out)
	if err != nil {
		return err
	}
	defer file.Close()

	stop := make(chan struct{})
	go func() {
		if d := *duration; d == 0 {
			println("Press enter to stop capturing...")
			os.Stdin.Read([]byte{0})
		} else {
			time.Sleep(d)
		}
		close(stop)
	}()
	_, err = gapii.Capture(logger, *spyport, file, options, stop)
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
	if verbs.Verbosity > 0 {
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
