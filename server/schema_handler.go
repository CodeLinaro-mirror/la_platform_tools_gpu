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

	"android.googlesource.com/platform/tools/gpu/gfxapi/schema"
)

const (
	apiParamName = "api"
)

// ServeHTTP writes to res a human-readable plain text description of the schema for
// the api parameters parsed from the given req query string.
func schemaHandler(res http.ResponseWriter, req *http.Request) {
	s := schema.Schema()

	res.Header().Add("Content-Type", "text/plain;charset=UTF-8")

	fmt.Fprintln(res, "Atoms:")
	for i, e := range s.Atoms {
		fmt.Fprintf(res, "(%v) %+v\n", i, e)
	}

	fmt.Fprintf(res, "Apis: %+v\n", s.Apis)
}
