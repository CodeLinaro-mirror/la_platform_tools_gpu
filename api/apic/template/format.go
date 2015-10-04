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
	"fmt"
	"io/ioutil"
	"os/exec"

	"golang.org/x/tools/imports"

	"android.googlesource.com/platform/tools/gpu/tools/verbs"
)

const (
	disable    = '⋖'
	enable     = '⋗'
	indent     = '»'
	unindent   = '«'
	suppress   = '§'
	newline    = '¶'
	whitespace = '•'
)

// Reflow does the primitive reflow, but no language specific handling.
func (f *Functions) Reflow(indentSize int, value string) (string, error) {
	verbs.Logf("Reflowing string\n")
	result, err := reflow(value, indentSize)
	if err != nil {
		return "", fmt.Errorf("%s : %s", f.active.Name(), err)
	}
	return string(result), nil
}

const goIndent = 2 // Required by go style guide

// GoFmt reflows the string as if it were go code using the standard go fmt library.
func (f *Functions) GoFmt(value string) (string, error) {
	verbs.Logf("Reflowing go code\n")
	result, err := reflow(value, goIndent)
	if err != nil {
		return "", fmt.Errorf("%s : %s", f.active.Name(), err)
	}
	opt := &imports.Options{
		TabWidth:  goIndent,
		TabIndent: true,
		Comments:  true,
		Fragment:  true,
	}
	formatted, err := imports.Process(f.active.Name(), result, opt)
	if err != nil {
		verbs.Logf("GoFmt failed with %s\n", err)
		return string(result), nil
	}
	return string(formatted), nil
}

// Format reflows the string using an external command.
func (f *Functions) Format(command stringList, value string) (string, error) {
	if len(command) == 0 {
		return "", fmt.Errorf("%s : Invalid Format command", f.active.Name())
	}
	binary := command[0]
	verbs.Logf("Reflowing code with %s\n", binary)
	// indent level is arbitrary, because we expect the external formatter to redo it anyway
	result, err := reflow(value, 4)
	if err != nil {
		return "", fmt.Errorf("%s : %s", f.active.Name(), err)
	}

	_, err = exec.LookPath(binary)
	if err != nil {
		verbs.Logf("Could not find external formatter %s\n", binary)
		return string(result), nil
	}
	cmd := exec.Command(binary, command[1:]...)
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		verbs.Logf("Reformat out pipe failed: %s\n", err)
		return string(result), nil
	}
	stdin, err := cmd.StdinPipe()
	if err != nil {
		verbs.Logf("Reformat in pipe failed: %s\n", err)
		return string(result), nil
	}
	err = cmd.Start()
	if err != nil {
		verbs.Logf("Reformat start failed: %s\n", err)
		return string(result), nil
	}
	_, err = stdin.Write(result)
	if err != nil {
		verbs.Logf("Reformat write failed: %s\n", err)
		return string(result), nil
	}
	err = stdin.Close()
	if err != nil {
		verbs.Logf("Reformat close failed: %s\n", err)
		return string(result), nil
	}
	formatted, err := ioutil.ReadAll(stdout)
	if err != nil {
		verbs.Logf("Reformat read: %s\n", err)
		return string(result), nil
	}
	err = cmd.Wait()
	if err != nil {
		verbs.Logf("Reformat wait failed: %s\n", err)
		return string(result), nil
	}
	return string(formatted), nil
}

func panicWrite(buf *bytes.Buffer, r rune) {
	_, err := buf.WriteRune(r)
	if err != nil {
		panic(err)
	}
}

func reflow(in string, indentSize int) ([]byte, error) {
	depth := 0
	wasNewline := false
	suppressing := true
	join := false
	enabled := true
	buf := &bytes.Buffer{}
	flushPending := func() {
		if wasNewline && !suppressing {
			// write the indent
			panicWrite(buf, '\n')
			for i := 0; i < depth*indentSize; i++ {
				panicWrite(buf, ' ')
			}
		}
		suppressing = false
		wasNewline = false
		join = false
	}
	for _, ch := range in {
		if !enabled {
			if ch == enable {
				enabled = true
			} else {
				panicWrite(buf, ch)
			}
		} else {
			switch ch {
			case disable:
				flushPending()
				enabled = false
				ch = 0
			case whitespace:
				ch = ' '
			case suppress:
				suppressing = true
				ch = 0
			case newline:
				panicWrite(buf, '\n')
				fallthrough
			case '\n', '\r':
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
	}
	return buf.Bytes(), nil
}
