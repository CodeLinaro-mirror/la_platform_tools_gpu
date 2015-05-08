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
	feed         chan Request
	context      batcherContext
	persistentDb database.Database
	transientDb  database.Database
	device       Device
	logger       log.Logger
}

func (b *batcher) run() {
	defer b.logger.Flush()

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
			b.logger.Errorf("%v", err)
		}
	}
}

func (b *batcher) send(requests []Request) (err error) {
	postbackHandlers := make(executor.PostbackHandlerMap)

	defer func() {
		if r := recover(); r != nil {
			err = fmt.Errorf("Failed to send replay: %v", r)
		}
		if err != nil {
			// Report errors to all postback handlers.
			for _, h := range postbackHandlers {
				h(nil, err)
			}
		}
	}()

	var c service.Capture
	if err := b.transientDb.Load(b.context.CaptureID.ID, b.logger, &c); err != nil {
		return fmt.Errorf("Failed to load capture (%s): %v", b.context.CaptureID, err)
	}

	var stream service.AtomStream
	if err := b.transientDb.Load(c.Atoms.ID, b.logger, &stream); err != nil {
		return fmt.Errorf("Failed to load atom stream (%s): %v", c.Atoms, err)
	}

	atoms, err := stream.List()
	if err != nil {
		return err
	}

	nextID := atom.ID(0x10000000)
	postback := func(handler PostbackHandler) atom.ID {
		id := nextID
		nextID++
		postbackHandlers[id] = handler
		return id
	}

	td := b.device.Info()

	transforms := b.context.Generator.ReplayTransforms(
		b.context.Context,
		b.context.Config,
		requests,
		postback,
		td,
		b.persistentDb,
		b.logger)

	if config.DebugReplay {
		b.logger.Infof("Replaying %d atoms using transform chain:", len(atoms))
		for i, t := range transforms {
			b.logger.Infof("(%d) %#v", i, t)
		}
	}

	architecture := b.device.Info().Architecture()

	builder := builder.New(architecture)

	transforms.Transform(atoms, &adapter{
		state:   gfxapi.NewState(),
		db:      b.persistentDb,
		logger:  b.logger,
		builder: builder,
	})

	if config.DebugReplay {
		b.logger.Infof("Building payload...")
	}

	payload, decoder, err := builder.Build(b.logger)
	if err != nil {
		return fmt.Errorf("Failed to build replay payload: %v", err)
	}

	connection, err := b.device.Connect()
	if err != nil {
		return fmt.Errorf("Failed to connect to device %v: %v", td.Name, err)
	}
	defer connection.Close()

	if config.DebugReplay {
		b.logger.Infof("Sending payload to %v.", td.Name)
	}

	return executor.Execute(
		payload,
		decoder,
		connection,
		b.persistentDb,
		b.logger,
		postbackHandlers,
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

func (w *adapter) Write(id atom.ID, a atom.Atom) {
	Replay(id, a, w.state, w.db, w.logger, w.builder)
}
