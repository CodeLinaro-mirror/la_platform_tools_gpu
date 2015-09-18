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

package do

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"

	"android.googlesource.com/platform/tools/gpu/maker/config"
	"android.googlesource.com/platform/tools/gpu/maker/graph"
)

// ExecAt executes "path" with the specified arguments with the working
// directory set to "wd".
func ExecAt(wd string, verbose int, path string, args ...string) error {
	before := time.Now()
	cmd := exec.Command(path, args...)
	cmd.Env = getEnvVars()
	cmd.Dir = wd

	if verbose > 0 {
		log.Printf("-> %s", strings.Join(cmd.Args, " "))
	}
	if verbose > 2 {
		log.Printf("Working directory: %v", cmd.Dir)
		log.Printf("Environment:")
		for _, v := range cmd.Env {
			log.Printf("• %s", v)
		}
	}

	output, err := cmd.CombinedOutput()
	duration := time.Now().Sub(before)
	switch {
	case err != nil:
		err = fmt.Errorf("Command: \"%v\" %v", strings.Join(cmd.Args, " "), err)
		log.Printf("\n\n%s\n%v", string(output), err)
		return err
	case verbose > 1:
		if msg := string(output); msg != "" {
			log.Printf("\n%s\n--- %s %v succeeded ---", msg, path, duration)
		} else {
			log.Printf("%s %v succeeded", path, duration)
		}
	}
	return nil
}

// Exec builds and returns a new Step that runs the specified external binary
// with the supplied arguments. The newly created Step will be made to depend on
// the binary.
func Exec(binary graph.Entity, args ...string) *graph.Step {
	wd := config.Paths.Root
	return graph.NewStep(func(step *graph.Step) error {
		return ExecAt(wd, config.Verbose, binary.Name(), args...)
	}).DependsOn(binary)
}

type envVar struct {
	key, value string
}

func getEnvVars() []string {
	env := map[string]envVar{}
	for _, pair := range os.Environ() {
		v := strings.LastIndex(pair, "=")
		key, value := pair[:v], pair[v+1:]
		env[strings.ToUpper(key)] = envVar{key, value}
	}

	sep := fmt.Sprintf("%c", filepath.ListSeparator)

	for key, values := range config.EnvVars {
		keyUpper := strings.ToUpper(key)
		if existing, found := env[keyUpper]; found {
			values = append(values, existing.value)
			env[keyUpper] = envVar{existing.key, strings.Join(values, sep)}
		} else {
			env[keyUpper] = envVar{key, strings.Join(values, sep)}
		}
	}

	vars := []string{}
	for _, v := range env {
		vars = append(vars, v.key+"="+v.value)
	}
	return vars
}
