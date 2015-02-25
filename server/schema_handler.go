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

	"android.googlesource.com/platform/tools/gpu/gfxapi"
)

const (
	apiParamName = "api"
)

// ServeHTTP writes to res a human-readable plain text description of the schema for
// the api parameters parsed from the given req query string.
func schemaHandler(res http.ResponseWriter, req *http.Request) {
	apiName := req.URL.Query().Get("api")
	api := gfxapi.Find(apiName)
	if api == nil {
		http.NotFound(res, req)
		return
	}

	schema := api.Schema()

	res.Header().Add("Content-Type", "text/plain;charset=UTF-8")

	fmt.Fprintln(res, "Arrays:")
	for i, e := range schema.Arrays {
		fmt.Fprintf(res, "(%v) %#v\n", i, e)
	}

	fmt.Fprintln(res, "Maps:")
	for i, e := range schema.Maps {
		fmt.Fprintf(res, "(%v) %#v\n", i, e)
	}

	fmt.Fprintln(res, "Enums:")
	for i, e := range schema.Enums {
		fmt.Fprintf(res, "(%v) %#v\n", i, e)
	}

	fmt.Fprintln(res, "Structs:")
	for i, e := range schema.Structs {
		fmt.Fprintf(res, "(%v) %#v\n", i, e)
	}

	fmt.Fprintln(res, "Classes:")
	for i, e := range schema.Classes {
		fmt.Fprintf(res, "(%v) %#v\n", i, e)
	}

	fmt.Fprintln(res, "Atoms:")
	for i, e := range schema.Atoms {
		fmt.Fprintf(res, "(%v) %#v\n", i, e)
	}

	fmt.Fprintf(res, "State: %#v\n", schema.State)
}
