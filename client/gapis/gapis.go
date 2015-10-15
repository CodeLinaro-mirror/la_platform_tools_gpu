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

// Package gapis provides helper methods and types for communicating with the
// GAPIS service.
package gapis

import (
	"fmt"

	"net"

	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/binary/schema"
	"android.googlesource.com/platform/tools/gpu/client/process"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
	"android.googlesource.com/platform/tools/gpu/service"
)

const mtu = 1024

// Connect attempts to connect to a GAPIS process.
// If port is zero, a new GAPIS server will be started, otherwise a connection will be made to the specified port.
func Connect(port int, logger log.Logger) (service.Service, schema.Message, error) {
	var socket net.Conn
	var err error
	if port == 0 {
		socket, err = process.StartAndConnect("gapis")
	} else {
		socket, err = process.Connect(port)
	}
	if err != nil {
		return nil, schema.Message{}, err
	}
	multiplexer := multiplexer.New(socket, socket, socket, mtu, logger, nil)
	client := service.NewClient(multiplexer, nil)

	message, err := client.GetSchema(logger)
	if err != nil {
		return nil, schema.Message{}, fmt.Errorf("Error resolving schema: %v", err)
	}

	namespace := registry.NewNamespace()
	for _, entity := range message.Entities {
		namespace.Add((*schema.ObjectClass)(entity))
	}

	// Replace the current client with the schema aggregated namespace.
	client = service.NewClient(multiplexer, registry.NewNamespace(registry.Global, namespace))

	return client, message, nil
}
