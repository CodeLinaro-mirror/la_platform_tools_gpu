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

package gapii

import (
	"fmt"
	"io"
	"math"
	"net"
	"time"

	"android.googlesource.com/platform/tools/gpu/log"
)

const CaptureTag = "GapiiTraceFile_V1.1"

type Options struct {
	// If true, then a framebuffer-observation will be made after every end-of-frame.
	ObserveFramebufferOnEOF bool
	// If true, then a framebuffer-observation will be made after every draw call.
	ObserveFramebufferOnDrawCall bool
}

func closed(s chan struct{}) bool {
	select {
	case <-s:
		return true
	default:
		return false
	}
}

const sizeGap = 1024 * 1024 * 5
const timeGap = time.Second

type siSize int64

var formats = []string{
	"%.0fB",
	"%.2fKB",
	"%.2fMB",
	"%.2fGB",
	"%.2fTB",
	"%.2fPB",
	"%.2fEB",
}

func (s siSize) String() string {
	if s == 0 {
		return "0.0B"
	}
	size := float64(s)
	e := math.Floor(math.Log(size) / math.Log(1000))
	f := formats[int(e)]
	v := math.Floor(size/math.Pow(1000, e)*10+0.5) / 10
	return fmt.Sprintf(f, v)
}

func capture(logger log.Logger, port int, w io.Writer, o Options, stop chan struct{}) (int64, error) {
	if closed(stop) {
		return 0, nil
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return 0, nil // Treat failure-to-connect as target-not-ready instead of an error.
	}
	defer conn.Close()
	if err := sendHeader(conn, o); err != nil {
		return 0, err
	}

	var count, nextSize siSize
	startTime := time.Now()
	nextTime := startTime
	for {
		if closed(stop) {
			log.Infof(logger, "Stop: %v", count)
			break
		}
		now := time.Now()
		conn.SetReadDeadline(now.Add(time.Millisecond * 100)) // Allow for stop event and UI refreshes.
		n, err := io.CopyN(w, conn, 1024*64)
		count += siSize(n)
		if err == io.EOF {
			log.Infof(logger, "EOF: %v", count)
			break
		}
		if err != nil {
			err, isnet := err.(net.Error)
			if !isnet || (!err.Temporary() && !err.Timeout()) {
				log.Infof(logger, "Connection error: %v", err)
				return int64(count), err
			}
		}
		if count > nextSize || now.After(nextTime) {
			nextSize = count + sizeGap
			nextTime = now.Add(timeGap)
			delta := time.Duration(int64(now.Sub(startTime)/time.Millisecond)) * time.Millisecond
			log.Infof(logger, "Capturing: %v in %v", count, delta)
		}
	}
	return int64(count), nil
}

// Capture opens up the specified port and then waits for a capture to be
// delivered using the specified capture options.
// It copies the capture into the supplied writer.
func Capture(logger log.Logger, port int, w io.Writer, options Options, stop chan struct{}) (int64, error) {
	log.Infof(logger, "Waiting for connection to localhost:%d...", port)
	for {
		count, err := capture(logger, port, w, options, stop)
		if err != nil {
			return count, err
		}
		if count != 0 {
			return count, nil
		}
		// ADB has an annoying tendancy to insta-close forwarded sockets when
		// there's no application waiting for the connection. Treat this as
		// another waiting-for-connection case.
		select {
		case <-stop:
			log.Infof(logger, "Aborted.")
			return 0, nil
		case <-time.After(500 * time.Millisecond):
			log.Infof(logger, "Retry...")
		}
	}
}
