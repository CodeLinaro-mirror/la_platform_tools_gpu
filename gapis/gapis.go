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

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary/registry"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
	"android.googlesource.com/platform/tools/gpu/process"
	"android.googlesource.com/platform/tools/gpu/service"
)

const mtu = 1024

// Connect attempts to connect to an existing GAPIS process at the specified
// address, returning the service interface and schema on success. If no GAPIS
// instance can be found, then a new instance will be created.
func Connect(address, data string, logger log.Logger) (service.Service, service.Schema, error) {
	args := []string{
		"--rpc", address,
		"--data", data,
		"--shutdown_on_disconnect",
	}

	socket, err := process.ConnectStartIfNeeded(address, "gapis", args...)
	if err != nil {
		return nil, service.Schema{}, err
	}

	multiplexer := multiplexer.New(socket, socket, socket, mtu, logger, nil)
	client := service.NewClient(multiplexer, nil)

	schema, err := client.GetSchema(logger)
	if err != nil {
		return nil, service.Schema{}, fmt.Errorf("Error resolving schema: %v", err)
	}

	namespace := registry.NewNamespace()
	for _, class := range schema.Classes {
		// Find the atom metadata, if present
		if meta := atom.FindMetadata(class); meta != nil {
			namespace.Add(newAtomClass(class, meta))
		} else {
			namespace.Add(class)
		}
	}

	// Replace the current client with the schema aggregated namespace.
	client = service.NewClient(multiplexer, registry.NewNamespace(registry.Global, namespace))

	return client, schema, nil
}
