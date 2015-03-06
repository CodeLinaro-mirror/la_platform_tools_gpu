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

// The rpcapi command generates RPC client, server and transport code
// automatically from an API description.
package main

import (
	"flag"
	"fmt"
	"io/ioutil"
	"os"

	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/rpc/generate"
)

var (
	golang = flag.Bool("go", false, "enable go generation")
	java   = flag.Bool("java", false, "enable java generation")
)

func run() error {
	flag.Parse()
	if len(flag.Args()) < 1 {
		return fmt.Errorf("Missing api file")
	}
	if !(*golang || *java) {
		return fmt.Errorf("Specify languages to build")
	}
	apiName := flag.Args()[0]
	info, err := ioutil.ReadFile(apiName)
	if err != nil {
		return err
	}
	parsed, errs := parser.Parse(string(info[:]))
	if err != nil {
		return err
	}
	compiled, errs, _ := resolver.Resolve(parsed)
	if len(errs) > 0 {
		return errs[0]
	}
	f := generate.Init(apiName, compiled)
	if len(errs) > 0 {
		return errs[0]
	}
	if *golang {
		generate.Go(f)
	}
	if *java {
		generate.Java(f)
	}
	return nil
}

func main() {
	if err := run(); err != nil {
		fmt.Fprintf(os.Stderr, "rpcgen failed: %v\n", err)
		os.Exit(1)
	}
}
