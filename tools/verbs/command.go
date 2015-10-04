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

package verbs

import (
	"flag"
	"fmt"
	"strings"
)

// Verb holds information about a runnable api command.
type Verb struct {
	Name      string                         // The name of the command
	Run       func(flags flag.FlagSet) error // the action for the command
	ShortHelp string                         // Help for how to use the command
	Flags     flag.FlagSet                   // The command line flags it accepts
}

var (
	verbs = []*Verb{}
)

// Register adds a new verb to the supported set, it will panic if a
// duplicate name is encountered.
func Register(v *Verb) {
	if len(Filter(v.Name)) != 0 {
		panic(fmt.Errorf("Duplicate verb name %s", v.Name))
	}
	verbs = append(verbs, v)
}

// Filter returns the filtered list of verbs who's names match the specified
// prefix.
func Filter(prefix string) (result []*Verb) {
	for _, v := range verbs {
		if strings.HasPrefix(v.Name, prefix) {
			result = append(result, v)
		}
	}
	return result
}
