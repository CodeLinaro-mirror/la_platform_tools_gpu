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

package template

import (
	"bytes"
	"flag"
	"io/ioutil"
	"os"
	"path/filepath"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/parser"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/tools/copyright"
)

var (
	command = &commands.Command{
		Name:      "template",
		ShortHelp: "Passes the ast to a template for code generation",
	}
	dir    = command.Flags.String("dir", cwd(), "The output directory")
	tracer = command.Flags.String("t", "", "The template function trace expression")
)

func init() {
	command.Flags.Var(&globalList, "G", "A global value setting for the template")
	command.Run = doTemplate
	commands.Register(command)
}

func cwd() string {
	p, _ := os.Getwd()
	return p
}

// Include loads each of the templates and executes their main bodies.
// The filenames are relative to the template doing the include.
func (f *Functions) Include(templates ...string) error {
	original := f.active
	dir := ""
	if original != nil {
		dir = filepath.Dir(original.Name())
	}
	for _, t := range templates {
		if dir != "" {
			t = filepath.Join(dir, t)
		}
		if f.templates.Lookup(t) == nil {
			commands.Log("Reading template %q\n", t)
			tmplData, err := f.loader(t)
			commands.MaybeError(t, err)
			tmpl, err := f.templates.New(t).Parse(string(tmplData))
			commands.MaybeError(t, err)
			f.active = tmpl
			commands.Log("Executing template %q\n", f.active.Name())
			var buf bytes.Buffer
			commands.MaybeError(f.active.Name(), f.active.Execute(&buf, f.api))
		}
	}
	f.active = original
	return nil
}

// Write takes a string and writes it into the specified file.
// The filename is relative to the output directory.
func (f *Functions) Write(fileName string, value string) (string, error) {
	outputPath := filepath.Join(f.basePath, fileName)
	commands.Log("Writing output to %q\n", outputPath)
	return "", ioutil.WriteFile(outputPath, []byte(value), 0666)
}

// Copyright emits the copyright header specified by name with the «Tool» set to tool.
func (f *Functions) Copyright(name string, tool string) (string, error) {
	return copyright.Build(name, copyright.Info{Year: "2015", Tool: tool}), nil
}

func doTemplate(flags flag.FlagSet) {
	args := flags.Args()
	if len(args) < 1 {
		commands.Usage("Missing api file\n")
	}
	apiName := args[0]
	if len(args) < 2 {
		commands.Usage("Missing template file\n")
	}
	mainTemplate := args[1]
	commands.Log("Reading api file %q\n", apiName)
	info, err := ioutil.ReadFile(apiName)
	commands.MaybeError(apiName, err)
	commands.Log("Compiling api file %q\n", apiName)
	parsed, errs := parser.Parse(string(info[:]))
	commands.CheckErrors(apiName, errs)
	compiled, errs := resolver.Resolve(parsed)
	commands.CheckErrors(apiName, errs)
	f := NewFunctions(apiName, compiled, ioutil.ReadFile, nil)
	commands.MaybeError(mainTemplate, f.Include(mainTemplate))
}
