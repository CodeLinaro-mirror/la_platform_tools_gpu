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
package main

import (
	"flag"
	"fmt"
	"os"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/gapis"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

const usage = `report: A tool to check a capture replays without issues.
Usage: report <capture>
  -help: show this help message
`

var (
	gapisAddr = flag.String("gapis", "localhost:6700", "gapis tcp host:port to connect to")
	dataPath  = flag.String("data", "data", "Path to the server's data folder")
)

func run() error {
	flag.Usage = func() {
		fmt.Printf(usage)
		flag.PrintDefaults()
	}
	flag.Parse()

	if flag.NArg() != 1 {
		flag.Usage()
		return fmt.Errorf("Invalid number of arguments. Expected 1, got %d", flag.NArg())
	}

	capture := flag.Arg(0)

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

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "report failed: %v\n", err)
		os.Exit(1)
	}
}
