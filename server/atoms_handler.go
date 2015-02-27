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
	"strconv"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

const (
	captureParamName = "capture"
	contextParamName = "ctx"
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
	captures, err := h.Captures()
	if err != nil {
		panic(err)
	}

	captureName := req.URL.Query().Get(captureParamName)
	captureID, found := captures[captureName]
	if !found {
		http.NotFound(res, req)
		return
	}

	var c service.Capture
	if err := h.Load(captureID, log.Nop{}, &c); err != nil {
		panic(err)
	}

	var stream service.AtomStream
	if err := h.Load(c.Atoms.ID, log.Nop{}, &stream); err != nil {
		panic(err)
	}

	atoms, err := stream.List()
	if err != nil {
		panic(err)
	}

	var contextID atom.ContextID
	if ctx, err := strconv.ParseInt(req.URL.Query().Get(contextParamName), 10, 32); err == nil {
		contextID = atom.ContextID(int(ctx))
	}

	res.Header().Add("Content-Type", "text/plain;charset=UTF-8")

	for i, a := range atoms {
		if a.ContextID() == contextID {
			fmt.Fprintf(res, "%.6d %s\n", i, a)
		}
	}
}
