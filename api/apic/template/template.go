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
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime/debug"
	"text/template"
	"unicode/utf8"

	"android.googlesource.com/platform/tools/gpu/api"
	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
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
	deps   = command.Flags.String("deps", "", "The dependancies file to generate")
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

var (
	inputs  []string
	outputs []string
)

func inputDep(name string) {
	path, _ := filepath.Abs(name)
	inputs = append(inputs, path)
}

func outputDep(name string) {
	path, _ := filepath.Abs(name)
	outputs = append(outputs, path)
}

func writeDeps() error {
	if len(*deps) == 0 {
		return nil
	}
	commands.Logf("Write deps to %v\n", *deps)
	file, err := os.Create(*deps)
	if err != nil {
		return err
	}
	fmt.Fprintln(file, "==Inputs==")
	for _, entry := range inputs {
		fmt.Fprintln(file, entry)
	}
	fmt.Fprintln(file, "==Outputs==")
	for _, entry := range outputs {
		fmt.Fprintln(file, entry)
	}
	return file.Close()
}

// Note isTrimSpace is only testing the Latin1 spaces.
func isTrimSpace(b byte) bool {
	switch b {
	case ' ', '\n':
		return true
	}
	return false
}

// trimWriter, writes to the underlying io.Writer, but with leading
// and trailing spaces trimmed from the output. Only current trailing
// spaces are saved between calls to Write().
type trimWriter struct {
	out     io.Writer
	atStart bool   // at the start we throw away leading spaces.
	spaces  []byte // spaces which are currently trailing.
}

func newTrimWriter(out io.Writer) io.Writer {
	return &trimWriter{out: out, atStart: true}
}

func (t *trimWriter) Write(buf []byte) (int, error) {
	l := len(buf)
	if l == 0 {
		return 0, nil
	}

	begin := 0 // index of the first byte to output
	if t.atStart {
		// Skip over leading spaces
		// Find the start of the interesting content.
		for ; begin < len(buf); begin++ {
			b := buf[begin]
			// If the character is in Latin1, it is safe to treat it is a byte
			if b >= utf8.RuneSelf || !isTrimSpace(b) {
				t.atStart = false
				break
			}
		}

		if t.atStart {
			// The whole buffer is leading spaces
			return l, nil
		}
	}

	// Find the end of the interesting content (remove trailing spaces).
	end := len(buf) // index one beyond the end of the interesting content
	for ; end > begin; end-- {
		b := buf[end-1]
		// If the character is in Latin1, it is safe to treat it is a byte
		if b >= utf8.RuneSelf || !isTrimSpace(b) {
			break
		}
	}

	if begin == end {
		// The whole buffer is trailing spaces
		t.spaces = append(t.spaces, buf...)
		return l, nil
	}

	// The buffer has some content to output.
	// First output any trailing spaces from the previous call
	if len(t.spaces) != 0 {
		if ws, err := t.out.Write(t.spaces); err != nil || ws != len(t.spaces) {
			return ws, err
		}
		// We are done with previous trailing spaces
		t.spaces = nil
	}

	// Output the content.
	if ws, err := t.out.Write(buf[begin:end]); err != nil || ws != end-begin {
		return ws, err
	}

	if end != len(buf) {
		// Save any trailing spaces
		t.spaces = append(t.spaces, buf[end:]...)
	}

	return l, nil
}

func (f *Functions) execute(active *template.Template, writer io.Writer, data interface{}) (err error) {
	olda := f.active
	oldw := f.writer
	f.active = active
	if writer != nil {
		f.writer = writer
	}
	f.writer = newTrimWriter(f.writer)
	defer func() {
		if r := recover(); r != nil {
			// There doesn't appear to be a clean way to get both the panic stack
			// and the template stack. This is the closest I can figure.
			err = fmt.Errorf("panic executing template %v: %v %v", f.active.Name(), r, string(debug.Stack()))
		}

		f.active = olda
		f.writer = oldw
	}()
	return f.active.Execute(f.writer, data)
}

// Include loads each of the templates and executes their main bodies.
// The filenames are relative to the template doing the include.
func (f *Functions) Include(templates ...string) error {
	dir := ""
	if f.active != nil {
		dir = filepath.Dir(f.active.Name())
	}
	for _, t := range templates {
		if dir != "" {
			t = filepath.Join(dir, t)
		}
		if f.templates.Lookup(t) == nil {
			commands.Logf("Reading template %q\n", t)
			inputDep(t)
			tmplData, err := f.loader(t)
			commands.MaybeError(t, err)
			tmpl, err := f.templates.New(t).Parse(string(tmplData))
			commands.MaybeError(t, err)
			commands.Logf("Executing template %q\n", tmpl.Name())
			var buf bytes.Buffer
			commands.MaybeError(tmpl.Name(), f.execute(tmpl, &buf, f.api))
		}
	}
	return nil
}

// Write takes a string and writes it into the specified file.
// The filename is relative to the output directory.
func (f *Functions) Write(fileName string, value string) (string, error) {
	outputPath := filepath.Join(f.basePath, fileName)
	commands.Logf("Writing output to %q\n", outputPath)
	outputDep(outputPath)

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
	commands.Logf("Reading api file %q\n", apiName)
	inputDep(apiName)

	commands.Logf("Compiling api file %q\n", apiName)
	mappings := resolver.ASTToSemantic{}
	compiled, errs := api.Resolve(apiName, mappings)
	commands.CheckErrors(apiName, errs)
	f := NewFunctions(apiName, compiled, ioutil.ReadFile, nil)
	commands.MaybeError(mainTemplate, f.Include(mainTemplate))
	writeDeps()
}
