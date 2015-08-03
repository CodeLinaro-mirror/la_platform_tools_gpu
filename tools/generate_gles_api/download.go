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

package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
)

// Download the given URL, optionally using disk cache
func Download(url string) []byte {
	filename := *cacheDir + string(os.PathSeparator) + CompileRegexp(`[^\w\.]+`).ReplaceAllString(url, "-")
	if *cacheDir != "" {
		if bytes, err := ioutil.ReadFile(filename); err == nil {
			return bytes
		}
	}
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	if resp.StatusCode != 200 {
		panic(fmt.Errorf("%s: %s", url, resp.Status))
	}
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		panic(err)
	}
	resp.Body.Close()
	if *cacheDir != "" {
		os.MkdirAll(*cacheDir, 0750)
		ioutil.WriteFile(filename, bytes, 0666)
	}
	return bytes
}
