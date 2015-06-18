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

package server

import (
	"fmt"
	"net/http"
	"strings"

	"android.googlesource.com/platform/tools/gpu/builder"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

const (
	idParamName   = "id"
	nameParamName = "name"
)

// AtomsHandler is an HTTP request handler that returns a human-readable description
// of the atoms for a given capture and context.
type atomsHandler struct {
	database.Database
}

// ServeHTTP writes to res a human-readable plain text description for each of the
// atoms of the capture/context identified by their respective parameters, parsed
// from the given req query string.
func (h atomsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	logger := log.Nop{}

	captures, err := builder.Captures(h, logger)
	if err != nil {
		panic(err)
	}

	var capture *service.Capture

	if id := req.URL.Query().Get(idParamName); id != "" {
		for _, cid := range captures {
			if strings.EqualFold(cid.ID.String(), id) {
				capture, _ = service.ResolveCapture(h, logger, cid)
				break
			}
		}
	} else if name := req.URL.Query().Get(nameParamName); name != "" {
		for _, id := range captures {
			capture, err := service.ResolveCapture(h, logger, id)
			if err == nil {
				if strings.EqualFold(capture.Name, name) {
					break
				}
			}
		}
	}

	if capture.Name == "" {
		http.NotFound(res, req)
		return
	}

	stream, err := service.ResolveAtomStream(h, log.Nop{}, capture.Atoms)
	if err != nil {
		panic(err)
	}

	atoms, err := stream.List()
	if err != nil {
		panic(err)
	}

	res.Header().Add("Content-Type", "text/plain;charset=UTF-8")

	for i, a := range atoms {
		fmt.Fprintf(res, "%.6d %s\n", i, a)
	}
}
