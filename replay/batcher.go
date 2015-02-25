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
	"time"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
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
	device       device
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
		b.send(requests)
	}
}

func (b *batcher) send(requests []Request) {
	var c service.Capture
	if err := b.transientDb.Load(b.context.CaptureID.ID, log.Nop{}, &c); err != nil {
		b.logger.Warning("Failed to load capture (%s): %v", b.context.CaptureID, err)
		return
	}

	var stream service.AtomStream
	if err := b.transientDb.Load(c.Atoms.ID, log.Nop{}, &stream); err != nil {
		b.logger.Warning("Failed to load atom stream (%s): %v", c.Atoms, err)
		return
	}

	atoms, err := stream.List()
	if err != nil {
		panic(err)
	}

	postbackHandlers := make(postbackHandlerMap)
	nextId := atom.Id(0x10000000)
	postback := func(handler PostbackHandler) atom.Id {
		id := nextId
		nextId++
		postbackHandlers[id] = handler
		return id
	}

	td := b.device.transportDevice()

	transforms := b.context.Generator.ReplayTransforms(
		b.context.Context,
		b.context.Config,
		requests,
		postback,
		td,
		b.persistentDb,
		b.logger)

	b.logger.Info("Replaying using transform chain: ")
	for i, t := range transforms {
		b.logger.Info("(%d) %#v", i, t)
	}

	builder := builder.New(int(td.PointerSize), int(td.PointerAlignment))
	writer := b.context.Generator.ReplayWriter(builder)

	transforms.Transform(atoms, adapter{writer, postbackHandlers})

	payload, decoder := builder.Build(b.logger)

	executor{
		payload:  payload,
		decoder:  decoder,
		device:   b.device,
		database: b.persistentDb,
		logger:   b.logger,
		handlers: postbackHandlers,
	}.execute()
}

// adapter conforms to the the atom Writer interface, forwarding writes to a
// replay Writer.
type adapter struct {
	writer   Writer
	handlers postbackHandlerMap
}

func (w adapter) Write(id atom.Id, a atom.Atom) {
	_, postback := w.handlers[id]
	w.writer.Write(id, a, postback)
}
