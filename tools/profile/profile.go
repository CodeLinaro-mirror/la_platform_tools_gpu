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

// Package profile contains helpers for adding profiling code to applications.
// Including this package will add the command line flag cpuprofile.
package profile

import (
	"flag"
	"log"
	"os"
	"runtime/pprof"
)

var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")

// CPU begins cpu profiling, and returns a function you can call to stop the
// profiling.
// The normal use is to defer the returned function, so putting
//   defer profile.CPU()()
// In the top of main it will profile the application from that point until main
// exits.
// You must have called flag.Parse before calling this function.
func CPU() func() {
	if *cpuprofile == "" {
		return func() {}
	}
	log.Printf("CPU profiling enbled, writing to %s\n", *cpuprofile)
	f, err := os.Create(*cpuprofile)
	if err != nil {
		log.Fatal(err)
	}
	pprof.StartCPUProfile(f)
	return func() {
		pprof.StopCPUProfile()
		f.Close()
		log.Printf("CPU profiling written to %s\n", *cpuprofile)
	}
}
