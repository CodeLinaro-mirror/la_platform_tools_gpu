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

package process

import (
	"fmt"
	"net"
	"os"
	"os/exec"
	"time"

	"android.googlesource.com/platform/tools/gpu/atexit"
)

// Number of attempts to connect to make before giving up.
const connectRetries = 20

// Amount of time to backoff before making another attempt to connect.
const connectBackoffTime = 100 * time.Millisecond

// StartProcess starts a new process of the executable at path with args.
// stderr, stdout and stdin are not inherited by the new process.
// When the process terminates procDone is called back.
func StartProcess(procDone func(*os.ProcessState, error), path string, args ...string) (*os.Process, error) {
	// try to run the server ourselves
	null, err := os.OpenFile(os.DevNull, os.O_RDWR, 0)
	if err != nil {
		return nil, err
	}

	path, err = exec.LookPath(path)
	if err != nil {
		return nil, err
	}

	procAttr := &os.ProcAttr{Files: []*os.File{null, null, null}}
	proc, err := os.StartProcess(path, args, procAttr)
	if err != nil {
		return nil, err
	}

	go func() {
		// Call wait so that the child process does not become a zombie.
		procDone(proc.Wait())
	}()

	return proc, nil
}

// ConnectStartIfNeeded will connect to the server process at addr, if it fails
// to connect and hostname is "localhost", then it will start a new process
// of the executable at path with args and then try again to connect.
func ConnectStartIfNeeded(addr string, path string, args ...string) (net.Conn, error) {
	conn, err := net.Dial("tcp", addr)
	if err == nil {
		// connected
		return conn, nil
	}

	// connection failed, was it localhost?
	if host, _, addrErr := net.SplitHostPort(addr); addrErr != nil {
		return nil, addrErr
	} else if host != "localhost" {
		return nil, err
	}

	terminated := make(chan struct{})
	done := func(state *os.ProcessState, err error) {
		close(terminated)
	}
	args = append([]string{path}, args...)
	proc, err := StartProcess(done, path, args...)
	if err != nil {
		return nil, fmt.Errorf("Failed to start process %v for connection %v: %v",
			path, addr, err)
	}

	// Register an atexit handler to kill the process
	atexit.Register(func() {
		select {
		case <-terminated:
			// Process already terminated
			break
		default:
			proc.Kill()
		}
	}, time.Second)

	for i := 0; i < connectRetries; i++ {
		conn, err = net.Dial("tcp", addr)
		if err == nil {
			// connected
			return conn, nil
		}

		if i+1 == connectRetries {
			break
		}

		// Backoff before retrying
		select {
		case <-time.After(connectBackoffTime):
			break
		case <-terminated:
			return nil, fmt.Errorf(
				"Spawned process %v terminated before we connected", path)
		}
	}

	if killErr := proc.Kill(); killErr != nil {
		return nil, fmt.Errorf(
			"Failed to connect to %v in time: %v and kill %v failed with %v", addr, err, path, killErr)
	}
	return nil, fmt.Errorf("Failed to connect to %v in time: %v", path, err)
}
