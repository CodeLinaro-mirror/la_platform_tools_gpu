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
	"net/http"
	_ "net/http/pprof"
	"os"
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/_experimental/client"
	"android.googlesource.com/platform/tools/gpu/atexit"
)

var (
	data         = flag.String("data", "data", "path to data")
	gapis        = flag.String("gapis", "localhost:6700", "gapis tcp host:port to connect to")
	capture      = flag.String("capture", "", "name of the capture to load")
	gxuiDebug    = flag.Bool("gxuidebug", false, "enable GXUI debug")
	replayDevice = flag.String("device", "Local machine", "name of the device to replay from")
)

const usage = `gapid: A user interface for the android gpu debugger.
Usage: gapid [--data=dir] [--gapis=address]
  -help: show this help message
`

func run() error {
	defer atexit.Exit(0)
	flag.Usage = func() {
		fmt.Printf(usage)
		flag.PrintDefaults()
	}
	flag.Parse()
	dataAbsPath, err := filepath.Abs(*data)
	if err != nil {
		return err
	}

	go http.ListenAndServe("localhost:6060", nil) // Profiling

	config := client.Config{
		DataPath:       dataAbsPath,
		Gapis:          *gapis,
		GXUIDebug:      *gxuiDebug,
		InitialCapture: *capture,
		ReplayDevice:   *replayDevice,
	}
	if config.InitialCapture == "" {
		config.InitialCapture = os.Getenv("LOAD_CAPTURE")
	}
	client.Run(config)
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "gapid failed: %v\n", err)
		os.Exit(1)
	}
}
