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
	"fmt"
	"strings"
	"text/template"
	"unicode"
	"unicode/utf8"
)

var (
	javaTemplates = template.Must(template.New("").Funcs(javaFuncs).Parse(`
{{define "Enum"}}{{.Name}}Enum{{end}}
{{define "ID"}}{{.Name}}ID{{end}}
{{define "IDBytes"}}{{.Name}}IDBytes{{end}}

{{define "Encoder"}}
  public static void encode(Encoder e, {{class .Name}} o) throws IOException {
{{range .Fields}}    {{encode (print "o." (fieldname .Name)) .Type}}
{{end}}  }{{end}}

{{define "EncodeNative"}}e.{{lower .Type.Method}}({{.Name}});{{end}}
{{define "EncodeRemap"}}{{.Name}}.encode(e);{{end}}
{{define "EncodeCodeable"}}{{.Name}}.encode(e);{{end}}
{{define "EncodeObject"}}e.object({{.Name}});{{end}}

{{define "EncodeArray"}}e.int32({{.Name}}.length);
    for (int i = 0; i < {{.Name}}.length; i++) {
      {{encode (print .Name "[i]") .Type.SubType}}
    }{{end}}

{{define "Decoder"}}
  public static void decode(Decoder d, {{class .Name}} o) throws IOException {
{{range .Fields}}    {{decode (print "o." (fieldname .Name)) .Type}}
{{end}}  }{{end}}

{{define "DecodeNative"}}{{.Name}} = d.{{lower .Type.Method}}();{{end}}
{{define "DecodeRemap"}}{{.Name}} = {{.Type.Name}}.decode(d);{{end}}
{{define "DecodeCodeable"}}{{.Name}} = new {{.Type.Name}}(d);{{end}}
{{define "DecodeObject"}}{{.Name}} = ({{storage .Type}})d.object();{{end}}

{{define "DecodeArray"}}{{.Name}} = new {{storage .Type.SubType}}[d.int32()];
    for (int i = 0; i < {{.Name}}.length; i++) {
      {{decode (print .Name "[i]") .Type.SubType}}
    }{{end}}

{{define "File"}}/*
 * Copyright (C) 2015 The Android Open Source Project
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 *      http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 *
 * THIS WILL BE REMOVED ONCE THE CODE GENERATOR IS INTEGRATED INTO THE BUILD.
 */
package {{.Package}};

import com.android.tools.rpclib.binary.BinaryObject;
import com.android.tools.rpclib.binary.BinaryObjectCreator;
import com.android.tools.rpclib.binary.Decoder;
import com.android.tools.rpclib.binary.Encoder;
import com.android.tools.rpclib.binary.ObjectTypeID;
import java.io.IOException;

class ObjectFactory {
  public enum Entries implements BinaryObjectCreator {{"{"}}{{range .Structs}}
    {{template "Enum" .}} {
      @Override public BinaryObject create() {
        return new {{class .Name}}();
      }
    },{{end}}
  }
{{range .Structs}}
  public static byte[] {{template "IDBytes" .}} = { {{range .ID}}{{toS8 .}}, {{end}}{{"}"}};{{end}}
{{range .Structs}}
  public static ObjectTypeID {{template "ID" .}} = new ObjectTypeID({{template "IDBytes" .}});{{end}}

  static {{"{"}}{{range .Structs}}
    ObjectTypeID.register({{template "ID" .}}, Entries.{{template "Enum" .}});{{end}}
  }{{range .Structs}}
{{template "Encoder" .}}
{{template "Decoder" .}}{{end}}
}
{{end}}
	`))
	javaFuncs = template.FuncMap{
		"encode": func(name string, t *Type) string {
			return kindDispatch(javaEncodeMap, name, t)
		},
		"decode": func(name string, t *Type) string {
			return kindDispatch(javaDecodeMap, name, t)
		},
		"lower": strings.ToLower,
		"fieldname": func(s string) string {
			r, n := utf8.DecodeRuneInString(s)
			return fmt.Sprintf("my%s%s", string(unicode.ToUpper(r)), s[n:])
		},
		"toS8": func(val byte) string { return fmt.Sprint(int8(val)) },
		"storage": func(t *Type) string {
			name := t.Name
			if t.Kind == Pointer {
				name = t.SubType.Name
			}
			if result, ok := javaTypeMap[name]; ok {
				return result
			}
			return name
		},
		"class": func(name string) string {
			if strings.HasPrefix(name, "call") {
				return fmt.Sprintf("Commands.%s.Call", name[4:])
			}
			if strings.HasPrefix(name, "result") {
				return fmt.Sprintf("Commands.%s.Result", name[6:])
			}
			return name
		},
	}
	javaEncodeMap kindToTemplate
	javaDecodeMap kindToTemplate
	javaFile      *template.Template
	javaTypeMap   = map[string]string{
		"int8":    "byte",
		"uint8":   "byte",
		"int16":   "short",
		"uint16":  "short",
		"int32":   "int",
		"uint32":  "int",
		"int64":   "long",
		"uint64":  "long",
		"float32": "float",
		"float64": "double",
	}
)

func init() {
	javaFile = getTemplate(javaTemplates, "File")
	javaEncodeMap = kindToTemplate{
		Native:    getTemplate(javaTemplates, "EncodeNative"),
		Remap:     getTemplate(javaTemplates, "EncodeRemap"),
		Codeable:  getTemplate(javaTemplates, "EncodeCodeable"),
		Pointer:   getTemplate(javaTemplates, "EncodeObject"),
		Interface: getTemplate(javaTemplates, "EncodeObject"),
		Array:     getTemplate(javaTemplates, "EncodeArray"),
	}
	javaDecodeMap = kindToTemplate{
		Native:    getTemplate(javaTemplates, "DecodeNative"),
		Remap:     getTemplate(javaTemplates, "DecodeRemap"),
		Codeable:  getTemplate(javaTemplates, "DecodeCodeable"),
		Pointer:   getTemplate(javaTemplates, "DecodeObject"),
		Interface: getTemplate(javaTemplates, "DecodeObject"),
		Array:     getTemplate(javaTemplates, "DecodeArray"),
	}
}

// JavaFile generates the all the java code for a file with a set of structs.
func JavaFile(file *File) ([]byte, error) {
	b := &bytes.Buffer{}
	err := javaFile.Execute(b, file)
	return b.Bytes(), err
}
