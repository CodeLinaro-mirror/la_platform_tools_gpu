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

// Package process has helper code for managing server processes and connections to them.
package process

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	"bufio"
	"regexp"

	"strconv"

	"android.googlesource.com/platform/tools/gpu/atexit"
)

var portPattern = regexp.MustCompile(`^Bound on port '(\d+)'$`)

func Start(name string, args ...string) (int, error) {
	cmd := exec.Command(name, args...)
	cmd.Stderr = os.Stderr
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return 0, err
	}
	if err := cmd.Start(); err != nil {
		return 0, err
	}
	terminated := make(chan struct{})
	// Register an atexit handler to kill the process
	atexit.Register(func() {
		select {
		case <-terminated:
			// Process already terminated
			break
		default:
			cmd.Process.Kill()
		}
	}, time.Second)
	// Now we search the output for the line that tells us the port
	prefix := fmt.Sprint(name, ":")
	lines := bufio.NewScanner(stdout)
	port := ""
	for port == "" && lines.Scan() {
		line := lines.Text()
		fmt.Fprintln(os.Stdout, prefix, line)
		match := portPattern.FindStringSubmatch(line)
		if match != nil {
			port = match[1]
		}
	}
	if err := lines.Err(); err != nil {
		close(terminated)
		return 0, err
	}
	if port == "" {
		close(terminated)
		return 0, fmt.Errorf("%s: did not start", name)
	}
	// We need to keep copying the stdout now
	go func() {
		defer close(terminated)
		for lines.Scan() {
			fmt.Fprintln(os.Stdout, prefix, lines.Text())
		}
		if err := lines.Err(); err != nil {
			fmt.Fprintln(os.Stderr, prefix, lines.Err())
		}
	}()
	return strconv.Atoi(port)
}

func Connect(port int) (net.Conn, error) {
	return net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
}

func StartAndConnect(name string, args ...string) (net.Conn, error) {
	if port, err := Start(name, args...); err != nil {
		return nil, err
	} else {
		return Connect(port)
	}
}
