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

// Package validate registers and implements the "validate" apic command.
//
// The validate command analyses the specified API for correctness, reporting errors if any problems
// are found.
package validate

import (
	"flag"
	"fmt"
	"os"

	"android.googlesource.com/platform/tools/gpu/api"
	"android.googlesource.com/platform/tools/gpu/api/apic/commands"
	"android.googlesource.com/platform/tools/gpu/api/resolver"
	"android.googlesource.com/platform/tools/gpu/api/semantic"
	"android.googlesource.com/platform/tools/gpu/parse"
	"android.googlesource.com/platform/tools/gpu/tools/verbs"
)

var (
	verb = &verbs.Verb{
		Name:      "validate",
		ShortHelp: "Validates an api file for correctness",
		Run:       doValidate,
	}
)

func init() {
	verbs.Register(verb)
}

func doValidate(flags flag.FlagSet) error {
	args := flags.Args()
	if len(args) < 1 {
		return verbs.Usage("Missing api file\n")
	}
	mappings := resolver.ASTToSemantic{}
	for _, apiName := range args {
		compiled, errs := api.Resolve(apiName, mappings)
		if err := commands.CheckErrors(apiName, errs); err != nil {
			return err
		}
		verbs.Logf("Validating api file %q\n", apiName)
		errors := Validate(apiName, compiled)
		for _, err := range errors {
			fmt.Fprintf(os.Stderr, "%s\n", err.Error())
		}
		if len(errors) > 0 {
			return errors[0]
		}
	}
	return nil
}

// Validate performs a number of checks on the api file for correctness.
// If any problems are found then they are returned as errors.
func Validate(apiName string, api *semantic.API) []error {
	errs := []error{}
	errs = append(errs, noUnusedTypes(apiName, api)...)
	for _, e := range noUnreachables(api) {
		errs = append(errs, err(apiName, e.At.Token(), e.Message))
	}
	return errs
}

func err(apiName string, at parse.Token, format string, args ...interface{}) error {
	l, c := at.Cursor()
	msg := fmt.Sprintf("%s:%d:%d ", apiName, l, c)
	return fmt.Errorf(msg+format, args...)
}
