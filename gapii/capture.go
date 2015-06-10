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

func closed(s chan struct{}) bool {
	select {
	case <-s:
		return true
	default:
		return false
	}
}

const messageGap = 1024 * 128

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
	size := float64(s)
	e := math.Floor(math.Log(size) / math.Log(1000))
	f := formats[int(e)]
	v := math.Floor(size/math.Pow(1000, e)*10+0.5) / 10
	return fmt.Sprintf(f, v)
}

func capture(logger log.Logger, port int, w io.Writer, stop chan struct{}) (int64, error) {
	if closed(stop) {
		return 0, nil
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		return 0, err
	}
	defer conn.Close()
	var count, nextMessage siSize
	for {
		if closed(stop) {
			logger.Infof("Stop: %v", count)
			break
		}
		conn.SetReadDeadline(time.Now().Add(time.Millisecond * 100)) // Allow for stop event and UI refreshes.
		n, err := io.CopyN(w, conn, 1024*64)
		count += siSize(n)
		if err == io.EOF {
			logger.Infof("EOF: %v", count)
			break
		}
		if err != nil {
			err, isnet := err.(net.Error)
			if !isnet || (!err.Temporary() && !err.Timeout()) {
				logger.Infof("Connection error: %v", err)
				return int64(count), err
			}
		}
		if count > nextMessage {
			nextMessage = count + messageGap
			logger.Infof("Capturing: %v", count)
		}
	}
	return int64(count), nil
}

// Capture opens up the specified port and then waits for a capture to be
// delivered.
// It copies the capture into the supplied writer.
func Capture(logger log.Logger, port int, w io.Writer, stop chan struct{}) (int64, error) {
	logger.Infof("Waiting for connection to localhost:%d...", port)
	for {
		count, err := capture(logger, port, w, stop)
		if err != nil {
			return count, err
		}
		if count != 0 {
			return count, nil
		}
		// ADB has an annoying tendancy to insta-close forwarded sockets when
		// there's no application waiting for the connection. Treat this as
		// another waiting-for-connection case.
		logger.Infof("Pausing...")
		select {
		case <-stop:
			logger.Infof("Aborted.")
			return 0, nil
		case <-time.After(500 * time.Millisecond):
			logger.Infof("Retry...")
		}
	}
}
