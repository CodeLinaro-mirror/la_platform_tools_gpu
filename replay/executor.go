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

package replay

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/vm"
)

const (
	connectionTypeDeviceInfo = 0
	connectionTypeReplay     = 1
)

const (
	messageTypeGet  = 0
	messageTypePost = 1
)

type postbackHandlerMap map[atom.ID]PostbackHandler

// ErrNoPostback is returned when a data for a postback could not be retrieved.
// This can be due to a connection problem or a decode error.
var ErrNoPostback = errors.New("No postback received")

type executor struct {
	payload  vm.Payload
	decoder  builder.ResponseDecoder
	device   device
	database database.Database
	logger   log.Logger
	handlers postbackHandlerMap
}

func (r executor) execute() {
	// Encode the payload
	buf := &bytes.Buffer{}
	e := binary.NewEncoder(buf)
	if err := r.payload.Encode(e); err != nil {
		panic(err)
	}

	// Store the payload to the database
	data := binary.Data(buf.Bytes())
	id, err := r.database.Store(&data, r.logger)
	if err != nil {
		panic(err)
	}

	// Kick the communication handler
	responseR, responseW := io.Pipe()
	go r.handleReplayCommunication(id, uint32(len(data)), responseW)

	// Decode and handle postbacks as they are received
	for postback := range r.decoder(responseR) {
		if handler, found := r.handlers[postback.ID]; found {
			handler(postback.Data, postback.Error)
			delete(r.handlers, postback.ID)
		} else {
			r.logger.Warning("No handler registered for postback id 0x%x (%T)",
				postback.ID, postback.Data)
		}
	}

	// Report missing postbacks as errors
	for _, handler := range r.handlers {
		handler(nil, ErrNoPostback)
	}
}

func (r executor) handleReplayCommunication(replayID binary.ID, replaySize uint32, postbacks io.WriteCloser) {
	logger := r.logger.Enter("handleReplayCommunication")

	connection, err := r.device.connect()
	if err != nil {
		logger.Error("Failed to connect to device %v (%v)", r.device.transportDevice().GetName(), err)
		postbacks.Close()
		return
	}
	defer connection.Close()

	e := binary.NewEncoder(connection)
	d := binary.NewDecoder(connection)

	if err = e.Uint8(connectionTypeReplay); err != nil {
		panic(err)
	}

	if err = e.String(replayID.String()); err != nil {
		panic(err)
	}

	if err = e.Uint32(replaySize); err != nil {
		panic(err)
	}

	for {
		dir, err := d.Uint8()
		if err != nil {
			break
		}

		switch dir {
		case messageTypeGet:
			r.handleGetData(connection)
		case messageTypePost:
			r.handleDataResponse(connection, postbacks)
		default:
			panic(fmt.Sprintf("Unknown request direction: %v\n", dir))
		}
	}
}

func (r executor) handleDataResponse(reader io.Reader, postbacks io.Writer) {
	d := binary.NewDecoder(reader)

	n, err := d.Uint32()
	if err != nil {
		panic(err)
	}

	_, err = io.CopyN(postbacks, reader, int64(n))
	if err != nil {
		panic(err)
	}
}

func (r executor) handleGetData(rw io.ReadWriter) {
	logger := r.logger.Enter("handleGetData")
	d := binary.NewDecoder(rw)

	resourceCount, err := d.Uint32()
	if err != nil {
		panic(err)
	}

	resourceIDs := make([]binary.ID, resourceCount)
	for i := range resourceIDs {
		idString, err := d.String()
		if err != nil {
			panic(err)
		}
		resourceIDs[i], err = binary.ParseID(idString)
		if err != nil {
			panic(err)
		}
		logger.Info("Replay requested resource '%v'", resourceIDs[i])
	}

	for _, rid := range resourceIDs {
		data := binary.Data{}
		err = r.database.Load(rid, logger, &data)
		if err != nil {
			panic(err)
		}
		if _, err := rw.Write(data); err != nil {
			panic(err)
		}
	}
}
