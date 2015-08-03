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

package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strings"
)

// Print most of our data to CSV file
func GenerateCsv(reg *Registry, api KhronosAPI) {
	out, err := os.Create(*csvDir + string(os.PathSeparator) + "gles.csv")
	if err != nil {
		panic(err)
	}
	writer := csv.NewWriter(out)
	writer.Write([]string{"command", "version", "paramIndex", "paramType", "paramName", "accepts"})
	for _, cmd := range reg.Command {
		cmdName := cmd.Proto.Name
		for _, version := range reg.GetVersions(api, cmdName) {
			doc := DownloadDoc(version, cmdName)
			if len(doc.Params)-1 /* result */ != len(cmd.Param) {
				panic(fmt.Errorf("%s: param count mismatch: %v vs %v", cmdName, len(doc.Params), len(cmd.Param)))
			}
			printCsvCommand(writer, &cmd, string(version), doc)
		}
		for _, extension := range reg.GetExtensions(api, cmdName) {
			printCsvCommand(writer, &cmd, extension, &CommandDoc{})
		}
	}
	writer.Flush()
	if err := out.Close(); err != nil {
		panic(err)
	}
}

func printCsvCommand(writer *csv.Writer, cmd *Command, definedBy string, doc *CommandDoc) {
	for i := -1; i < len(cmd.Param); i++ {
		param := cmd.ParamsAndResult()[i]
		accepts := ""
		if doc.Params != nil {
			paramDoc := doc.Params[i]
			if paramDoc.Name != param.Name {
				panic(fmt.Errorf("%s: param name mismatch: %s vs %s\n", cmd.Proto.Name, param.Name, paramDoc.Name))
			}
			accepts = strings.Join(paramDoc.Accepts, "|")
		}
		writer.Write([]string{cmd.Proto.Name, definedBy, fmt.Sprintf("%v", i), param.Type(), param.Name, accepts})
	}
}
