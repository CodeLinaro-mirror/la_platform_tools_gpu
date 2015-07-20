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

// Package server implements the rpc gpu debugger service, queriable by the
// clients, along with some helpers exposed via an http listener.
package server

import (
	"fmt"
	"net/http"
	_ "net/http/pprof"

	"android.googlesource.com/platform/tools/gpu/atexit"
	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay"
)

type Config struct {
	HttpAddress string
	RpcAddress  string
	DataPath    string
	LogfilePath string
}

const (
	atomsRoute    = "/atoms/"
	capturesRoute = "/captures/"

	mtu = 1024
)

// Run listens on the HTTP and RPC TCP ports given in config, initializes the resource database,
// the replay manager and handles HTTP and RPC requests on incoming connections.
func Run(config Config) {
	// Create the server logfile.
	logger, err := log.File(config.LogfilePath)
	if err != nil {
		panic(err)
	}
	defer log.Close(logger)
	fmt.Printf("Server log file created at: %s\n", config.LogfilePath)

	// Initialize the resource database and replay manager for RPC requests.
	b := &builder.Context{}
	database := database.NewInMemory(b)
	replayManager := replay.New(database, logger)
	b.ReplayManager = replayManager

	// Setup the RPC listener.
	rpc := &rpcServer{
		Database:      database,
		ReplayManager: replayManager,
	}

	// Setup and run the (blocking) HTTP listener on a separate goroutine.
	go func() {
		http.Handle(atomsRoute, http.StripPrefix(atomsRoute, atomsHandler{rpc}))
		http.Handle(capturesRoute, http.StripPrefix(capturesRoute, capturesHandler{rpc, logger, config}))
		if err := http.ListenAndServe(config.HttpAddress, nil); err != nil {
			log.Errorf(logger, "HTTP server shutdown with error %v", err)
		}
		log.Infof(logger, "HTTP server shutdown")
	}()

	// Run the blocking RPC listener
	if err := rpc.ListenAndServe(config.RpcAddress, mtu, logger); err != nil {
		log.Errorf(logger, "RPC Server shutdown with error %v", err)
		atexit.Exit(1)
	}

	log.Infof(logger, "RPC Server shutdown")
}
