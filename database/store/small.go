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

package store

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"time"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/log"
)

// Each record in the small archive is the id and a data buffer.
type keyValue struct {
	binary.Generate
	id     binary.ID
	buffer []byte
}

type smallArchive struct {
	worker         chan<- func()
	workerSync     chan<- func()
	data           *os.File
	records        map[binary.ID][]byte
	waste          int
	size           int
	path           string
	compactionSize int
	insertCh       chan<- *keyValue
	compacting     bool
}

const compactingFilename = "compacting"
const mainFilename = "small_data"

// CreateSmallArchive creates a Store optimized for small objects.
// When the resources are small enough the index entry of a regular archive
// uses more space than the data itself. The small archive stores the
// data unsorted in a single file. New records with a new value are
// appended to the data file. All current records are stored in RAM
// (this uses less RAM than storing all the index entries for the
// equivalent archive in RAM).
func CreateSmallArchive(path string, compactionSize int) Store {
	os.MkdirAll(path, 0755)

	data, err := os.OpenFile(filepath.Join(path, mainFilename), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	err = os.Remove(filepath.Join(path, compactingFilename))
	if err != nil && !os.IsNotExist(err) {
		if err != nil {
			panic(err)
		}
	}

	waste := 0
	size := 0
	records := make(map[binary.ID][]byte)

	// Use a buffered reader to quickly read the archive records
	reader := vle.Reader(bufio.NewReaderSize(data, 256<<10))
	for {
		r := &keyValue{}
		if err := cyclic.Decoder(reader).Value(r); err != io.EOF {
			if err != nil {
				panic(err)
			}
			if prevRecord, ok := records[r.id]; ok {
				// Archive contains multiple records with the same id.
				waste += len(prevRecord)
			}
			size += len(r.buffer)
			records[r.id] = r.buffer
		} else {
			break
		}
	}

	worker := make(chan func(), 64)
	workerSync := make(chan func())
	go func() {
		for worker != nil || workerSync != nil {
			select {
			case f, ok := <-worker:
				if ok {
					f()
				} else {
					worker = nil
				}
			case f, ok := <-workerSync:
				if ok {
					f()
				} else {
					workerSync = nil
				}
			}
		}
	}()

	return &smallArchive{
		worker:         worker,
		workerSync:     workerSync,
		data:           data,
		records:        records,
		waste:          waste,
		size:           size,
		path:           path,
		compactionSize: compactionSize,
		insertCh:       nil,
		compacting:     false,
	}
}

func (s *smallArchive) compaction(l log.Logger) error {
	if s.compacting {
		return fmt.Errorf("compaction() call when compaction in progress")
	}
	s.compacting = true

	fn := filepath.Join(s.path, compactingFilename)
	compacting, err := os.OpenFile(fn, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0666)
	if err != nil {
		s.compacting = false
		return err
	}

	// We compact the file whilst allowing concurrent updates. Synchronous
	// operations need to happen on the worker and therefore need to be
	// relatively fast, but IO is allowed. The bulk of the compacting IO happens
	// in a separate goroutine. This is what happens:
	// - In worker: create a channel for concurrent updates.
	// - In worker: make an in-memory copy of the records in an array.
	// - In worker: start compactor go routine.
	// - In compactor: output records using buffered IO.
	// - In compactor: whilst outputting records allow updates from channel.
	// When finished outputting:
	// - In compactor: synchronously send the completion function to the worker.
	// - In compactor: continue to process inserts until synchronous send is possible.
	// - In worker: process any remaining inserts.
	// - In worker: flip: close the compacting log, rename it and open it.
	// - In worker: put everything back to normal.

	// Create a channel for concurrent updates.
	insertCh := make(chan *keyValue, 64)
	s.insertCh = insertCh

	// Make an in-memory copy of the records. By copying into an array we can manage
	// about 20k records per millisecond.

	beforeCopyRecords := time.Now()
	recordsCopyArray := make([]*keyValue, len(s.records))
	i := 0
	for id, data := range s.records {
		recordsCopyArray[i] = &keyValue{id: id, buffer: data}
		i++
	}
	afterCopyRecords := time.Now()

	// Start compactor go routine.
	go func() {
		// Output records using buffered IO.
		writer := bufio.NewWriterSize(compacting, 256<<10)
		w := vle.Writer(writer)

		size := 0
		waste := 0

		numInserts := 0
		sizeInserts := 0

		beforeWriteRecords := time.Now()
		recordsWritten := make(map[binary.ID]int)

		didFirstInsertPath := false
		didSecondInsertPath := false
		didThirdInsertPath := false

		insertFun := func(insert *keyValue) {
			if prevSize, ok := recordsWritten[insert.id]; ok {
				// As new records can be added we need to keep track of waste
				waste += prevSize
			}

			err := cyclic.Encoder(w).Value(insert)
			if err != nil {
				panic(err)
			}

			size += len(insert.buffer)
			recordsWritten[insert.id] = len(insert.buffer)

			numInserts++
			sizeInserts += len(insert.buffer)
		}

		cleanupFun := func() {
			err := compacting.Close()
			if err != nil && l != nil {
				l.Warning("Failed to close compacting file")
			}
			err = os.Remove(fn)
			if err != nil && l != nil {
				l.Warning("Failed to remove compacting file")
			}
		}

		for _, record := range recordsCopyArray {
			err := cyclic.Encoder(w).Value(record)
			if err != nil {
				panic(err)
			}

			size += len(record.buffer)
			recordsWritten[record.id] = len(record.buffer)

			// Whilst outputting records allow updates from channel.

			// In the case of a rapid sustained insertion this loop may never
			// terminate. However, there are non-buffered writes done in the worker
			// thread which should be slower, making this unlikely. Also if there
			// is rapid sustained insertion it will fill the drive at some point
			// regardless of compaction.
		loop:
			for {
				select {
				case insert, ok := <-insertCh:
					if !ok {
						cleanupFun()
						return
					}

					insertFun(insert)
					didFirstInsertPath = true
				default:
					// nothing to insert right now
					break loop
				}
			}
		}

		// All the records are written, but more inserts may still arrive.

		afterWriteRecords := time.Now()

		completionFun := func() {
			// Process any remaining inserts.
			startSyncPart := time.Now()

			close(insertCh)
			s.insertCh = nil

			// The other end of the channel is now closed, but there may still
			// be messages to process.
			for insert := range insertCh {
				insertFun(insert)
				didSecondInsertPath = true
			}

			beforeFlush := time.Now()
			err := writer.Flush()
			if err != nil {
				panic(err)
			}
			afterFlush := time.Now()

			// Flip: close the compacting log, rename it and open it.
			err = compacting.Close()
			if err != nil {
				panic(err)
			}

			afterCatchup := time.Now()

			mainFn := filepath.Join(s.path, mainFilename)
			// TODO figure out whether rename is truely atomic on all supported OS.
			err = os.Rename(fn, mainFn)
			if err != nil {
				panic(err)
			}

			data, err := os.OpenFile(mainFn, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}

			// Put everything back to normal.
			s.data = data
			s.waste = waste
			s.size = size
			s.compacting = false

			afterAll := time.Now()

			if l != nil {
				l = l.Enter("Compaction")
				l.Info("After compacting Waste %v Size %v", s.waste, s.size)
				l.Info("Write records to disk %v", afterWriteRecords.Sub(beforeWriteRecords))
				l.Info("Num concurrent inserts %v size of inserts %v", numInserts, sizeInserts)
				l.Info("Sync with worker %v", startSyncPart.Sub(afterWriteRecords))
				l.Info("Catchup time %v", afterCatchup.Sub(startSyncPart))
				l.Info("Flush to disk %v", afterFlush.Sub(beforeFlush))
				l.Info("Flip and open %v", afterAll.Sub(afterCatchup))
				l.Info("Total sync part %v", afterAll.Sub(startSyncPart))
				l.Info("Total time %v", afterAll.Sub(beforeWriteRecords))

				l.Info("didFirstInsertPath %v", didFirstInsertPath)
				l.Info("didSecondInsertPath %v", didSecondInsertPath)
				l.Info("didThirdInsertPath %v", didThirdInsertPath)
			}
		}

		// Note this code executes before the completionFun above.
	loop2:
		for {
			select {
			case s.workerSync <- completionFun:
				break loop2
			// Keep draining the insertCh. As the completionFun is run using a
			// synchronous channel we know that no further insertions can be added
			// and that therefore deadlock can not happen.
			case insert, ok := <-insertCh:
				if !ok {
					cleanupFun()
					return
				}
				insertFun(insert)
				didThirdInsertPath = true
			}
		}
	}()

	if l != nil {
		l.Info("Compacting copy %v records: %v", len(recordsCopyArray), afterCopyRecords.Sub(beforeCopyRecords))
	}

	return err
}

// Store compliance
func (s *smallArchive) Store(id binary.ID, _ binary.Object, data []byte, logger log.Logger) error {
	if logger != nil {
		logger = logger.Enter("SmallArchive.Store")
		logger.Info("id: %s size: %d waste %v total_size %v", id, len(data), s.waste, s.size)
		if len(data) > (1 << 16) {
			logger.Warning("Store of non-small record in SmallArchive: %v bytes", len(data))
		}
	}

	s.worker <- func() {
		if prevRecord, ok := s.records[id]; ok {
			if bytes.Equal(prevRecord, data) {
				// No need to do anything, nothing changed
				return
			}

			// Archive contains multiple records with the same id.
			s.waste += len(prevRecord)
		}

		record := &keyValue{
			id:     id,
			buffer: data,
		}

		err := cyclic.Encoder(vle.Writer(s.data)).Value(record)
		if err != nil {
			panic(err)
		}

		s.size += len(data)
		s.records[id] = data

		if s.size > s.compactionSize && !s.compacting && s.waste*3 > s.compactionSize {
			if logger != nil {
				logger.Info("Small archive starting a compaction. Size %v Waste %v",
					s.size, s.waste)
			}
			s.compaction(logger)
		}

		// If compaction is in-flight send the record to the compaction goroutine.
		if s.insertCh != nil {
			s.insertCh <- record
		}
	}

	return nil
}

func (s *smallArchive) Load(id binary.ID, logger log.Logger, out binary.Object) (size int, err error) {
	if logger != nil {
		logger = logger.Enter("SmallArchive.Load")
		logger.Info("Loading from archive: %s", id)
		defer func() {
			logger.Info("↪ size: %d, err: %v", size, err)
			if err := recover(); err != nil {
				logger.Error("Panic when loading %v: %v", id, err)
				panic(err)
			}
		}()
	}

	done := make(chan struct{})
	s.worker <- func() {
		defer close(done)

		rec, found := s.records[id]
		if !found {
			size, err = 0, fmt.Errorf("Resource %v not found", id)
			return
		}

		d := cyclic.Decoder(vle.Reader(bytes.NewBuffer(rec)))
		size, err = len(rec), d.Value(out)
		return
	}
	<-done
	return
}

func (s *smallArchive) Contains(id binary.ID) (found bool) {
	done := make(chan struct{})
	s.worker <- func() {
		_, found = s.records[id]
		close(done)
	}
	<-done
	return
}

func (s *smallArchive) Close() {
	done := make(chan struct{})
	s.worker <- func() {
		s.data.Close()
		if s.insertCh != nil {
			close(s.insertCh)
			s.insertCh = nil
			s.compacting = false
		}
		close(done)
		close(s.worker)
		close(s.workerSync)
	}
	<-done
}
