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
	"fmt"
	"time"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/client/gapir"
	"android.googlesource.com/platform/tools/gpu/config"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/executor"
	"android.googlesource.com/platform/tools/gpu/service"
)

const maxBatchDelay = 250 * time.Millisecond

type batcherContext struct {
	Context
	Generator Generator
	Config    Config
}

type batcher struct {
	feed     chan Request
	context  batcherContext
	database database.Database
	device   gapir.Device
	logger   log.Logger
}

func (b *batcher) run() {
	// Gather all the batchEntries that are added to feed within maxBatchDelay.
	for r := range b.feed {
		requests := []Request{r}
		timeout := time.After(maxBatchDelay)
	inner:
		for {
			select {
			case r, ok := <-b.feed:
				if !ok {
					break inner
				}
				requests = append(requests, r)
			case <-timeout:
				break inner
			}
		}
		// Batch formed. Trigger the replay.
		if err := b.send(requests); err != nil {
			log.Errorf(b.logger, "%v", err)
		}
	}
}

func (b *batcher) send(requests []Request) (err error) {
	log.Infof(b.logger, "Replaying on device: %s", b.device.Info().Name)

	c, err := service.ResolveCapture(b.context.Capture, b.database, b.logger)
	if err != nil {
		return fmt.Errorf("Failed to load capture (%s): %v", b.context.Capture, err)
	}

	list, err := service.ResolveAtomList(c.Atoms, b.database, b.logger)
	if err != nil {
		return fmt.Errorf("Failed to load atom stream (%s): %v", c.Atoms, err)
	}

	device := b.device.Info()

	architecture := b.device.Info().Architecture()

	builder := builder.New(architecture)

	out := &adapter{
		state:   gfxapi.NewState(),
		db:      b.database,
		logger:  b.logger,
		builder: builder,
	}

	if err := b.context.Generator.Replay(
		b.context.Context,
		b.context.Config,
		requests,
		device,
		*list,
		out,
		b.database,
		b.logger); err != nil {

		// TODO: Inform the request that the replay failed.
		log.Errorf(b.logger, "Replay returned error: %v", err)
		return err
	}

	if config.DebugReplay {
		log.Infof(b.logger, "Building payload...")
	}

	payload, decoder, err := builder.Build(b.logger)
	if err != nil {
		return fmt.Errorf("Failed to build replay payload: %v", err)
	}

	defer func() {
		caught := recover()
		if err == nil && caught != nil {
			err, _ = caught.(error)
			if err == nil {
				// If we are panicing, we always want an error to send.
				err = fmt.Errorf("%s", caught)
			}
		}
		if err != nil {
			// An error was returned or thrown after the replay postbacks were requested.
			// Inform each postback handler that they're not going to get data,
			// to avoid chans blocking forever.
			decoder(nil, err)
		}
		if caught != nil {
			panic(caught)
		}
	}()

	connection, err := b.device.Connect()
	if err != nil {
		return fmt.Errorf("Failed to connect to device %v: %v", device.Name, err)
	}
	defer connection.Close()

	if config.DebugReplay {
		log.Infof(b.logger, "Sending payload to %v.", device.Name)
	}

	if Events.OnReplay != nil {
		Events.OnReplay(b.device, b.context.Context, b.context.Config, requests)
	}

	return executor.Execute(
		payload,
		decoder,
		connection,
		b.database,
		b.logger,
		architecture,
	)
}

// adapter conforms to the the atom Writer interface, performing replay writes
// on each atom.
type adapter struct {
	state   *gfxapi.State
	db      database.Database
	logger  log.Logger
	builder *builder.Builder
}

func (w *adapter) Write(i atom.ID, a atom.Atom) {
	w.builder.BeginAtom(i)
	if err := Replay(i, a, w.state, w.db, w.logger, w.builder); err == nil {
		w.builder.CommitAtom()
	} else {
		w.builder.RevertAtom(err)
	}
}
