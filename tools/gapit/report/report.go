// Copyright (C) 2014 The Android Open Source Project
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

// The report generates and displays a report for the given capture file.
package report

import (
	"flag"
	"fmt"
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/client/gapis"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
	"android.googlesource.com/platform/tools/gpu/tools/verbs"
)

var (
	verb = &verbs.Verb{
		Name:      "report",
		ShortHelp: "Check a capture replays without issues",
	}
	gapisAddr = verb.Flags.String("gapis", "localhost:6700", "gapis tcp host:port to connect to")
	dataPath  = verb.Flags.String("data", "data", "Path to the server's data folder")
)

func init() {
	verb.Run = doReport
	verbs.Register(verb)
}

func doReport(flags flag.FlagSet) error {
	if flags.NArg() != 1 {
		return verbs.Usage("Exactly one gfx trace file expected, got %d", flags.NArg())
	}

	capture, err := filepath.Abs(flags.Arg(0))
	if err != nil {
		return fmt.Errorf("Could not find capture file '%s': %v", flags.Arg(0), err)
	}

	logger := log.Std()
	defer log.Close(logger)

	client, _, err := gapis.Connect(*gapisAddr, *dataPath, logger)
	if err != nil {
		return fmt.Errorf("Failed to connect to the GAPIS server: %v", err)
	}

	devices, err := client.GetDevices(logger)
	if err != nil {
		return fmt.Errorf("Failed query list of devices: %v", err)
	}

	var devicePath *path.Device

	// TODO: Let the user control the device used.
	if len(devices) > 0 {
		devicePath = devices[0]
	} else {
		log.W(logger, "No replay device found")
	}

	capturePath, err := client.LoadCapture(capture, logger)
	if err != nil {
		return fmt.Errorf("Failed to load the capture file '%v': %v", capture, err)
	}

	boxedAtoms, err := client.Get(capturePath.Atoms(), logger)
	if err != nil {
		return fmt.Errorf("Failed to acquire the capture's atoms: %v", err)
	}
	atoms := boxedAtoms.(*atom.List).Atoms

	boxedReport, err := client.Get(capturePath.Report(devicePath), logger)
	if err != nil {
		return fmt.Errorf("Failed to acquire the capture's report: %v", err)
	}

	report := boxedReport.(*service.Report)
	for _, e := range report.Items {
		if e.Atom != uint64(atom.NoID) {
			log.Log(logger, e.Severity, "(%d) %v %v", e.Atom, atoms[e.Atom], e.Message)
		} else {
			log.Log(logger, e.Severity, e.Message)
		}
	}

	if len(report.Items) == 0 {
		log.I(logger, "No issues found")
	}

	return nil
}
