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

package generate

import (
	"bytes"
	"io"
	"text/template"

	"android.googlesource.com/platform/tools/gpu/tools/copyright"

	"golang.org/x/tools/imports"
)

var (
	goTemplates = template.Must(template.New("go.tmpl").Funcs(goFuncs).Parse(go_tmpl))
	goFuncs     = template.FuncMap{
		"encode": func(name string, t *Type) string {
			return kindDispatch(goEncodeMap, name, t)
		},
		"decode": func(name string, t *Type) string {
			return kindDispatch(goDecodeMap, name, t)
		},
		"header": func(tool string) string {
			return copyright.Build("generated_by", copyright.Info{Tool: tool})
		},
	}
	goEncodeMap kindToTemplate
	goDecodeMap kindToTemplate
	goFile      *template.Template
	goRegister  *template.Template
	goEncoder   *template.Template
	goDecoder   *template.Template
)

func init() {
	goFile = getTemplate(goTemplates, "File")
	goRegister = getTemplate(goTemplates, "Register")
	goEncoder = getTemplate(goTemplates, "Encoder")
	goDecoder = getTemplate(goTemplates, "Decoder")
	goEncodeMap = getTemplateMap(goTemplates, "Encode")
	goDecodeMap = getTemplateMap(goTemplates, "Decode")
}

// GoFile generates the all the go code for a file with a set of structs.
func GoFile(file *File) ([]byte, error) {
	b := &bytes.Buffer{}
	if err := goFile.Execute(b, file); err != nil {
		return nil, err
	}
	result, err := imports.Process("", b.Bytes(), nil)
	if err != nil {
		return b.Bytes(), nil
	}
	return result, nil
}

// GoRegister generates the go code to register a Struct with the system,
// writing it to an io.Writer.
func GoRegister(w io.Writer, s *Struct) error {
	return goRegister.Execute(w, s)
}

// GoEncoder generates an encoder for a Struct, writing it to an io.Writer.
func GoEncoder(w io.Writer, s *Struct) error {
	return goEncoder.Execute(w, s)
}

// GoEncoder generates a decoder for a Struct, writing it to an io.Writer.
func GoDecoder(w io.Writer, s *Struct) error {
	return goDecoder.Execute(w, s)
}
