// Copyright (C) 2014 The Android Open Source Project
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

package format

import (
	"bytes"
	"io"
	"text/tabwriter"
	"unicode/utf8"
)

const (
	EOL      = '¶' // outputs a line break
	Space    = '•' // outputs a space even where it would normally be stripped
	Column   = '║' // a column marker used to line up text
	Toggle   = '§' // switches the formatter on or off
	Indent   = '»' // increases the current indent level by 1
	Unindent = '«' // decreases the current indent level by 1
	Flush    = 'ø' // flushes text, resets column detection
)

func New(w io.Writer) *Writer {
	return &Writer{
		out:    tabwriter.NewWriter(w, 1, 2, 1, ' ', tabwriter.StripEscape),
		Indent: "  ",
	}
}

type Writer struct {
	Depth     int
	Indent    string
	Disabled  bool
	out       *tabwriter.Writer
	space     bytes.Buffer
	lastDepth int
	newline   bool
	stripping bool
	runeBuf   [4]byte
}

func (w *Writer) Write(data []byte) (n int, err error) {
	n = len(data)
	for _, r := range string(data) {
		if w.Disabled {
			switch r {
			case Toggle:
				w.Disabled = false
			default:
				err = w.WriteRune(r)
			}
			continue
		}
		switch r {
		case '\n':
			// line endings always stripped, also drops pending witespace
			w.space.Reset()
			w.stripping = true
		case ' ', '\t':
			// whitespace stripped from lines
			if !w.stripping {
				w.space.WriteRune(r)
			}
		case Column:
			w.space.WriteRune('\t')
			w.stripping = true
		case Toggle:
			w.Disabled = true
			w.space.Reset()
		case Indent:
			w.Depth++
		case Unindent:
			w.Depth--
		case Space:
			err = w.WriteRune(' ')
		case EOL:
			w.newline = true
			w.stripping = true
			w.space.Reset()
			err = w.WriteRune('\n')
		case Flush:
			w.Flush()
		default:
			if w.newline {
				w.newline = false
				if w.Depth != w.lastDepth {
					// indentation is different to last real write, so flush the tabwriter
					w.Flush()
					w.lastDepth = w.Depth
				}
				if w.Depth > 0 {
					w.runeBuf[0] = tabwriter.Escape
					w.out.Write(w.runeBuf[:1])
					for i := 0; i < w.Depth; i++ {
						io.WriteString(w.out, w.Indent)
					}
					w.out.Write(w.runeBuf[:1])
				}
			}
			w.stripping = false
			w.space.WriteTo(w.out)
			if err == nil {
				err = w.WriteRune(r)
			}
		}
	}
	return n, err
}

func (w *Writer) WriteRune(r rune) error {
	n := utf8.EncodeRune(w.runeBuf[:], r)
	_, err := w.out.Write(w.runeBuf[:n])
	return err
}

func (w *Writer) Flush() {
	w.out.Flush()
}
