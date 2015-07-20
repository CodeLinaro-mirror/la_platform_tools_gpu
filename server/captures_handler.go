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

	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/service"
)

// CapturesHandler is an HTTP request handler that returns a HTML list of
// captures held in the database.
type capturesHandler struct {
	s service.RPC
	l log.Logger
	c Config
}

// ServeHTTP writes to res a HTML list of captures held in the database.
func (h capturesHandler) ServeHTTP(res http.ResponseWriter, req *http.Request) {
	captures, err := h.s.GetCaptures(h.l)
	if err != nil {
		panic(err)
	}

	res.Header().Add("Content-Type", "text/html")

	fmt.Fprint(res, `<!DOCTYPE html>
<html>
	<head>
	<meta charset="UTF-8">
	<title>Captures</title>
	</head>
	<body>
`)

	for _, p := range captures {
		capture, err := service.GetCapture(p, h.s, h.l)
		if err == nil {
			fmt.Fprintf(res, `			<a href="%s%s?%s=%s">%s</a><br/>`,
				h.c.HttpAddress, atomsRoute, idParamName, p.ID, capture.Name)
		}
	}
	fmt.Fprint(res, "	</body>\n</html>")
}
