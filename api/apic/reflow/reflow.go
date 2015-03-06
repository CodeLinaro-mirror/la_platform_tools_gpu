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

package reflow

import (
	"bufio"
	"flag"
	"io/ioutil"
	"os"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
)

var (
	command = &commands.Command{
		Name:      "reflow",
		ShortHelp: "Reflows an api file for smart formatting",
	}
)

func init() {
	command.Run = doReflow
	commands.Register(command)
}

func doReflow(flags flag.FlagSet) {
	args := flags.Args()
	if len(args) < 1 {
		commands.Usage("Missing api file\n")
	}
	for _, apiName := range args {
		info, err := ioutil.ReadFile(apiName)
		commands.MaybeError(apiName, err)
		commands.Log("Compiling api file %q\n", apiName)
		parsed, errs := parser.Parse(string(info[:]))
		commands.CheckErrors(apiName, errs)
		compiled, errs, _ := resolver.Resolve(parsed)
		commands.CheckErrors(apiName, errs)
		commands.Log("Reflowing api %s\n", apiName)

		commands.Log("Writing output to %q\n", apiName)
		w, err := os.Create(apiName)
		commands.MaybeError(apiName, err)
		out := bufio.NewWriter(w)
		err = compiled.AST.CST.WriteTo(out)
		commands.MaybeError(apiName, err)
		err = out.Flush()
		commands.MaybeError(apiName, err)
		err = w.Close()
		commands.MaybeError(apiName, err)
	}
}
