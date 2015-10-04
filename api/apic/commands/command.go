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

package commands

import (
	"fmt"
	"os"

	"android.googlesource.com/platform/tools/gpu/parse"
)

const (
	maxErrors = 10
)

// CheckErrors will, if len(errs) > 0, print each of the error messages for the
// specified api and then terminate the program. If errs is zero length,
// CheckErrors does nothing.
func CheckErrors(apiName string, errs parse.ErrorList) error {
	if len(errs) == 0 {
		return nil
	}
	if len(errs) > maxErrors {
		errs = errs[:maxErrors]
	}
	for _, e := range errs {
		if e.At != nil {
			filename := e.At.Token().Source.Filename
			line, column := e.At.Token().Cursor()
			fmt.Fprintf(os.Stderr, "%s:%v:%v: %s\n", filename, line, column, e.Message)
		} else {
			fmt.Fprintf(os.Stderr, "%s: %s\n", apiName, e.Message)
		}
	}
	if len(errs) > maxErrors {
		fmt.Fprintf(os.Stderr, "And %d more errors\n", len(errs)-maxErrors)
	}
	fmt.Fprintf(os.Stderr, "Stack of first error:\n%s\n", errs[0].Stack)
	return errs
}
