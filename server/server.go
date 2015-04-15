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
	atomsRoute  = "/atoms/"
	schemaRoute = "/schema/"

	maxDataCacheSize       = 2 << 30   // 2 gigabytes
	maxDerivedCacheSize    = 1 << 29   // 0.5 gigabytes
	metaDataCompactionSize = 100 << 20 // 100 metabytes
	mtu                    = 1024
)

// Run listens on the HTTP and RPC TCP ports given in config, initializes the resource database,
// the replay manager and handles HTTP and RPC requests on incoming connections. If not nil, the
// rpcReady channel is closed as soon as incoming RPC requests can start being issued.
func Run(config Config, rpcReady chan<- struct{}) {
	// Create the server logfile.
	logger, err := log.File(config.LogfilePath)
	if err != nil {
		panic(err)
	}
	defer logger.Close()
	fmt.Printf("Server log file created at: %s\n", config.LogfilePath)

	// Initialize the resource database and replay manager for RPC requests.
	b := builder.New()
	database := database.Create(config.DataPath, maxDataCacheSize, metaDataCompactionSize, maxDerivedCacheSize, b)
	replayManager := replay.New(database, logger)
	b.SetReplayManager(replayManager)

	// Setup and run the (blocking) RPC listener on a separate goroutine.
	rpc := &rpcServer{
		Database:      database,
		ReplayManager: replayManager,
	}
	go rpc.ListenAndServe(config.RpcAddress, mtu, logger)

	// If provided, tell the caller chan that the RPC listener is ready.
	if nil != rpcReady {
		close(rpcReady)
	}

	// Setup and run the (blocking) HTTP listener.
	http.Handle(atomsRoute, http.StripPrefix(atomsRoute, atomsHandler{database}))
	http.Handle(schemaRoute, http.StripPrefix(schemaRoute, http.HandlerFunc(schemaHandler)))
	http.ListenAndServe(config.HttpAddress, nil)
}
