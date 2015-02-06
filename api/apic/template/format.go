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
	"io/ioutil"
	"os/exec"
	"strings"

	"golang.org/x/tools/imports"

	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
)

const (
	indent     = '»'
	unindent   = '«'
	suppress   = '§'
	newline    = '¶'
	whitespace = '•'
)

// SetIndentSize sets the number of whitespace characters used for a single
// indentation. The default is 2.
func (*Functions) SetIndentSize(i int) string {
	*indentSize = i
	return ""
}

func reformat(outputPath string, in string) []byte {
	commands.Log("Reflowing output for %s\n", outputPath)
	result, err := reflow(in)
	commands.MaybeError(outputPath, err)
	if *formatEnable {
		if strings.HasSuffix(outputPath, ".go") {
			opt := &imports.Options{
				TabWidth:  2,
				TabIndent: true,
				Comments:  true,
				Fragment:  true,
			}
			formatted, err := imports.Process(outputPath, result, opt)
			if err == nil {
				result = formatted
			} else {
				commands.Log("Reflow failed with %s for %s\n", err, outputPath)
			}
		} else if strings.HasSuffix(outputPath, ".h") || strings.HasSuffix(outputPath, ".cpp") {
			formatted, err := clangFormat(result)
			if err == nil {
				result = formatted
			} else {
				commands.Log("Reflow failed with %s for %s\n", err, outputPath)
			}
		}
	}
	return result
}

func panicWrite(buf *bytes.Buffer, r rune) {
	_, err := buf.WriteRune(r)
	if err != nil {
		panic(err)
	}
}

func reflow(in string) ([]byte, error) {
	depth := 0
	wasNewline := false
	suppressing := true
	join := false
	buf := &bytes.Buffer{}
	flushPending := func() {
		if wasNewline && !suppressing {
			// write the indent
			panicWrite(buf, '\n')
			for i := 0; i < depth*(*indentSize); i++ {
				panicWrite(buf, ' ')
			}
		}
		suppressing = false
		wasNewline = false
		join = false
	}
	for _, ch := range in {
		switch ch {
		case whitespace:
			ch = ' '
		case suppress:
			suppressing = true
			ch = 0
		case newline:
			panicWrite(buf, '\n')
			fallthrough
		case '\n':
			if !join {
				wasNewline = true
			}
			ch = 0
		case '\t', ' ':
			if wasNewline {
				ch = 0
			}
		case indent:
			ch = 0
			depth += 1
		case '{', '[':
			flushPending()
			depth += 1
		case unindent:
			ch = 0
			fallthrough
		case '}', ']':
			depth -= 1
		}

		if ch != 0 {
			flushPending()
			panicWrite(buf, ch)
		}
	}
	return buf.Bytes(), nil
}

func clangFormat(data []byte) ([]byte, error) {
	_, err := exec.LookPath("clang-format")
	if err != nil {
		return data, nil
	}

	// Style specifier to mostly match Android style
	style := "{BasedOnStyle: Google, AccessModifierOffset: -4, ColumnLimit: 100, ContinuationIndentWidth: 8, IndentWidth: 4}"
	cmd := exec.Command("clang-format", "-style", style)

	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}

	err = cmd.Start()
	if err != nil {
		return nil, err
	}

	_, err = stdin.Write(data)
	if err != nil {
		return nil, err
	}

	err = stdin.Close()
	if err != nil {
		return nil, err
	}

	formatted, err := ioutil.ReadAll(stdout)
	if err != nil {
		return nil, err
	}

	err = cmd.Wait()
	if err != nil {
		return nil, err
	}

	return formatted, nil
}
