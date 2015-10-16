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

	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
	"android.googlesource.com/platform/tools/gpu/service/path"
)

const (
	idParamName   = "id"
	nameParamName = "name"
)

// AtomsHandler is an HTTP request handler that returns a human-readable description
// of the atoms for a given capture and context.
type atomsHandler struct {
	s service.Service
}

// ServeHTTP writes to res a human-readable plain text description for each of the
// atoms of the capture/context identified by their respective parameters, parsed
// from the given req query string.
func (h atomsHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	logger := log.Nop{}

	captures, err := h.s.GetCaptures(logger)
	if err != nil {
		panic(err)
	}

	var capture *path.Capture

	if id := req.URL.Query().Get(idParamName); id != "" {
		for _, p := range captures {
			if strings.EqualFold(p.ID.String(), id) {
				capture = p
				break
			}
		}
	} else if name := req.URL.Query().Get(nameParamName); name != "" {
		for _, p := range captures {
			if c, err := service.GetCapture(p, h.s, logger); err != nil {
				panic(err)
			} else if strings.EqualFold(c.Name, name) {
				capture = p
				break
			}
		}
	}

	list, err := service.GetAtomList(capture.Atoms(), h.s, logger)
	if err != nil {
		panic(err)
	}

	res.Header().Add("Content-Type", "text/plain;charset=UTF-8")

	for i, a := range list.Atoms {
		var f string
		if o := a.Extras().Observations(); o != nil {
			if len(o.Reads) > 0 {
				f += " R"
			}
			if len(o.Writes) > 0 {
				f += " W"
			}
		}
		fmt.Fprintf(res, "%.6d %s %s\n", i, a, f)
	}
}
