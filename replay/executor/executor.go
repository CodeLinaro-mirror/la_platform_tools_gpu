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

// Package executor contains the Execute function for sending a replay to a device.
package executor

import (
	"bytes"
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/device"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
)

type executor struct {
	payload      protocol.Payload
	decoder      builder.ResponseDecoder
	connection   io.ReadWriteCloser
	database     database.Database
	logger       log.Logger
	architecture device.Architecture
}

// Execute sends the replay payload for execution on the target replay device
// communicating on connection.
// decoder will be used for decoding all postback reponses. Once a postback
// response is decoded, the corresponding handler in the handlers map will be
// called.
func Execute(
	payload protocol.Payload,
	decoder builder.ResponseDecoder,
	connection io.ReadWriteCloser,
	database database.Database,
	logger log.Logger,
	architecture device.Architecture) error {

	return executor{
		payload:      payload,
		decoder:      decoder,
		connection:   connection,
		database:     database,
		logger:       logger,
		architecture: architecture,
	}.execute()
}

func (r executor) execute() error {
	// Encode the payload
	buf := &bytes.Buffer{}
	e := flat.Encoder(endian.Writer(buf, r.architecture.ByteOrder))
	if err := e.Value(&r.payload); err != nil {
		return err
	}
	data := buf.Bytes()

	// Store the payload to the database
	id, err := database.Store(data, r.database, r.logger)
	if err != nil {
		return err
	}

	// Kick the communication handler
	responseR, responseW := io.Pipe()
	comErr := make(chan error)
	go func() {
		comErr <- r.handleReplayCommunication(id, uint32(len(data)), responseW)
	}()

	// Decode and handle postbacks as they are received
	r.decoder(responseR, nil)

	return <-comErr
}

func (r executor) handleReplayCommunication(replayID binary.ID, replaySize uint32, postbacks io.WriteCloser) error {
	connection := r.connection
	defer connection.Close()
	e := flat.Encoder(endian.Writer(connection, r.architecture.ByteOrder))
	d := flat.Decoder(endian.Reader(connection, r.architecture.ByteOrder))

	if err := e.Uint8(uint8(protocol.ConnectionTypeReplay)); err != nil {
		return err
	}

	if err := e.String(replayID.String()); err != nil {
		return err
	}

	if err := e.Uint32(replaySize); err != nil {
		return err
	}

	for {
		msg, err := d.Uint8()
		switch {
		case err == io.EOF:
			return nil
		case err != nil:
			return err
		}

		switch protocol.MessageType(msg) {
		case protocol.MessageTypeGet:
			if err := r.handleGetData(); err != nil {
				return err
			}
		case protocol.MessageTypePost:
			if err := r.handleDataResponse(postbacks); err != nil {
				return err
			}
		default:
			return fmt.Errorf("Unknown message type: %v\n", msg)
		}
	}
}

func (r executor) handleDataResponse(postbacks io.Writer) error {
	d := flat.Decoder(endian.Reader(r.connection, r.architecture.ByteOrder))

	n, err := d.Uint32()
	if err != nil {
		return err
	}

	c, err := io.CopyN(postbacks, r.connection, int64(n))
	if c != int64(n) {
		return err
	}

	return nil
}

func (r executor) handleGetData() error {
	logger := log.Enter(r.logger, "handleGetData")
	d := flat.Decoder(endian.Reader(r.connection, r.architecture.ByteOrder))

	resourceCount, err := d.Uint32()
	if err != nil {
		return err
	}

	resourceIDs := make([]binary.ID, resourceCount)
	for i := range resourceIDs {
		idString, err := d.String()
		if err != nil {
			return err
		}
		resourceIDs[i], err = binary.ParseID(idString)
		if err != nil {
			return err
		}
	}

	for _, rid := range resourceIDs {
		data, err := database.Resolve(rid, r.database, logger)
		if err != nil {
			return err
		}
		if _, err := r.connection.Write(data.([]byte)); err != nil {
			return err
		}
	}

	return nil
}
