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
	"os"
	"path/filepath"
	"runtime"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/client/gapir"
	"android.googlesource.com/platform/tools/gpu/server"
)

var (
	http                 = flag.String("http", "localhost:8080", "TCP host:port of the server's HTTP listener")
	rpc                  = flag.String("rpc", "localhost:6700", "TCP host:port of the server's RPC listener")
	logsPath             = flag.String("logs", "logs", "Directory to place log files")
	localDevicePort      = flag.Int("local_gapir_port", gapir.LocalPort, "Port number of the \"gapir\" running on the local device")
	shutdownOnDisconnect = flag.Bool("shutdown_on_disconnect", false, "Shutdown server when no connections remain")
)

func main() {
	defer atexit.Exit(0)

	flag.Parse()

	if os.Getenv("GOMAXPROCS") == "" {
		runtime.GOMAXPROCS(runtime.NumCPU())
	}

	gapisLogPath, _ := filepath.Abs(filepath.Join(*logsPath, "gapis.log"))
	gapir.LogPath, _ = filepath.Abs(filepath.Join(*logsPath, "gapir.log"))
	server.Run(server.Config{
		HttpAddress:          *http,
		RpcAddress:           *rpc,
		LogfilePath:          gapisLogPath,
		ShutdownOnDisconnect: *shutdownOnDisconnect,
		LocalPort:            *localDevicePort,
	})
}
