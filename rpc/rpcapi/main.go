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
	"os"
	"strings"

	"android.googlesource.com/platform/tools/gpu/api"
	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/rpc/generate"
)

var (
	dir    = flag.String("dir", cwd(), "The output directory")
	golang = flag.Bool("go", false, "enable go generation")
	java   = flag.Bool("java", false, "enable java generation")
)

func cwd() string {
	p, _ := os.Getwd()
	return p
}

func run() error {
	flag.Parse()
	if len(flag.Args()) < 1 {
		return fmt.Errorf("Missing api file")
	}
	if !(*golang || *java) {
		return fmt.Errorf("Specify languages to build")
	}
	if c := commands.Filter("template"); len(c) != 1 {
		return fmt.Errorf("Could not find template command")
	} else {
		c[0].Flags.Set("dir", *dir)
	}
	apiName := flag.Args()[0]
	mappings := resolver.ASTToSemantic{}
	compiled, errs := api.Resolve(apiName, mappings)
	if len(errs) > 0 {
		return errs
	}
	f := generate.Init(apiName, compiled)
	if len(errs) > 0 {
		return errs[0]
	}
	if *golang {
		generate.Go(f)
	}
	if *java {
		pkg := "unknown"
		i := strings.LastIndex(*dir, "/com/")
		if i >= 0 {
			pkg = strings.Replace((*dir)[i+1:], "/", ".", -1)
		}
		f.Global("package", pkg)
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
