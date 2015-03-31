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

type span struct {
	offset int64
	size   int64
}

func (s *span) decode(d binary.Decoder) (err error) {
	s.offset, err = d.Int64()
	if err != nil {
		return err
	}
	s.size, err = d.Int64()
	if err != nil {
		return err
	}
	return nil
}

func (s span) encode(e binary.Encoder) error {
	if err := e.Int64(s.offset); err != nil {
		return err
	}
	return e.Int64(s.size)
}

type record struct {
	id   binary.ID
	span span
}

func (s *record) decode(d binary.Decoder) (err error) {
	if s.id, err = d.ID(); err != nil {
		return err
	}
	return s.span.decode(d)
}

func (r record) encode(e binary.Encoder) error {
	if err := e.ID(r.id); err != nil {
		return err
	}
	return r.span.encode(e)
}

type compactingArchive struct {
	worker         chan<- func()
	workerSync     chan<- func()
	index          *os.File
	data           *os.File
	records        map[binary.ID]span
	waste          int64
	size           int64
	path           string
	compactionSize int64
	insertCh       chan<- keyValue
	compacting     bool
}

const (
	indexFilename           = "index"
	dataFilename            = "data"
	compactingDataFilename  = "compacting_" + dataFilename
	compactingIndexFilename = "compacting_" + indexFilename
	oldIndexFilename        = "old_" + indexFilename
	oldDataFilename         = "old_" + dataFilename
)

func exists(path string) bool {
	_, err := os.Stat(path)
	return err == nil
}

// CreateCompactingArchive returns a Store that attempts to reclaim space by
// dropping old resources that can be rebuilt on demand, and compacting the
// remaining resources to reclaim space.
func CreateCompactingArchive(path string) Store {
	os.MkdirAll(path, 0755)

	iFn := filepath.Join(path, indexFilename)
	dFn := filepath.Join(path, dataFilename)
	ciFn := filepath.Join(path, compactingIndexFilename)
	cdFn := filepath.Join(path, compactingDataFilename)
	oiFn := filepath.Join(path, oldIndexFilename)
	odFn := filepath.Join(path, oldDataFilename)

	// As flip can not be atomic. We follow these are the rules need to recover
	// in the case that we discover an incomplete flip.

	// If the index file is present then data is in either data or data_old
	// If not then if the data file is present the index is compacting_index
	// If neither exist then the data is in index_old and data_old
	// Otherwise create the files

	index, err := os.OpenFile(iFn, os.O_APPEND|os.O_RDWR, 0666)
	if err != nil && !os.IsNotExist(err) {
		if err != nil {
			panic(err)
		}
	}

	var data *os.File
	if err == nil {
		// The index file is present. The data is in either data or data_old
		data, err = os.OpenFile(dFn, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			if !os.IsNotExist(err) {
				if err != nil {
					panic(err)
				}
			}
			// The data does not exist, try data_old
			err = os.Rename(odFn, dFn)
			if err != nil {
				panic(err)
			}

			data, err = os.OpenFile(dFn, os.O_APPEND|os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
		}
		// Both files are open
	} else {
		// The index is not present, check to see if the data file is.
		data, err = os.OpenFile(dFn, os.O_APPEND|os.O_RDWR, 0666)
		if err != nil {
			if !os.IsNotExist(err) {
				if err != nil {
					panic(err)
				}
			}
			// Neither index nor data file exist. Check for index_old and data_old.
			if exists(odFn) && exists(oiFn) {
				// Both index_old and data_old exist. Rename them, index file first.
				err = os.Rename(oiFn, iFn)
				if err != nil {
					panic(err)
				}

				err = os.Rename(odFn, dFn)
				if err != nil {
					panic(err)
				}

				index, err = os.OpenFile(iFn, os.O_APPEND|os.O_RDWR, 0666)
				if err != nil {
					panic(err)
				}

				data, err = os.OpenFile(dFn, os.O_APPEND|os.O_RDWR, 0666)
				if err != nil {
					panic(err)
				}
			} else {
				// Create both files
				index, err = os.OpenFile(iFn, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
				if err != nil {
					panic(err)
				}

				data, err = os.OpenFile(dFn, os.O_CREATE|os.O_APPEND|os.O_RDWR, 0666)
				if err != nil {
					panic(err)
				}
			}
			// Both files are open
		} else {
			// The data file exists. Move the index from compacting_index.
			err = os.Rename(ciFn, iFn)
			if err != nil {
				panic(err)
			}

			index, err = os.OpenFile(iFn, os.O_APPEND|os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
			// Both files are open
		}
	}
	// Both files are open

	// Tidy up
	os.Remove(oiFn)
	os.Remove(odFn)
	os.Remove(ciFn)
	os.Remove(cdFn)

	waste := int64(0)
	size := int64(0)
	records := make(map[binary.ID]span)

	// Use a buffered reader to quickly read the archive records
	d := cyclic.Decoder(vle.Reader(bufio.NewReaderSize(index, 256<<10)))
	for {
		var r record
		if err := r.decode(d); err != io.EOF {
			if err != nil {
				panic(err)
			}
			if prevSpan, ok := records[r.id]; ok {
				// Archive contains multiple records with the same id.
				waste += prevSpan.size
			}
			if r.span.size != 0 {
				size += r.span.size
				records[r.id] = r.span
			}
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

	return &compactingArchive{
		worker:     worker,
		workerSync: workerSync,
		index:      index,
		data:       data,
		records:    records,
		waste:      waste,
		size:       size,
		path:       path,
		insertCh:   nil,
		compacting: false,
	}
}

// Store compliance
func (s *compactingArchive) Store(id binary.ID, _ binary.Object, data []byte, logger log.Logger) (err error) {
	if logger != nil {
		logger = logger.Enter("Archive.Store")
		logger.Info("id: %s size: %d", id, len(data))
	}
	s.worker <- func() {
		if prevRecord, ok := s.records[id]; ok {
			// Archive contains multiple records with the same id.
			s.waste += prevRecord.size
		}

		offset, err := s.data.Seek(0, os.SEEK_END)
		if err != nil {
			panic(err)
		}

		n, err := s.data.Write(data)
		if n != len(data) {
			panic(fmt.Errorf("Byte count written was not as expected (%d/%d), error: %v.", n, len(data), err))
		}
		if err != nil {
			panic(err)
		}

		record := record{
			id: id,
			span: span{
				offset: offset,
				size:   int64(len(data)),
			},
		}

		e := cyclic.Encoder(vle.Writer(s.index))
		record.encode(e)

		s.records[id] = record.span

		if s.size > s.compactionSize && !s.compacting && s.waste*3 > s.compactionSize {
			if logger != nil {
				logger.Info("Archive starting a compaction. Size %v Waste %v",
					s.size, s.waste)
			}
			s.compaction(logger)
		}

		// If compaction is in-flight send the record to the compaction goroutine.
		if s.insertCh != nil {
			kv := keyValue{id: id, buffer: data}
			s.insertCh <- kv
			if logger != nil {
				logger.Info("Insert sent to compactor. Size %v", len(data))
			}
		}
	}
	return nil
}

func (s *compactingArchive) Load(id binary.ID, logger log.Logger, out binary.Object) (size int, err error) {
	if logger != nil {
		logger = logger.Enter("Archive.Load")
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

		span, found := s.records[id]
		if !found {
			size, err = 0, fmt.Errorf("Resource %v not found", id)
			return
		}

		data := make([]byte, span.size)
		if n, e := s.data.ReadAt(data, span.offset); n != len(data) {
			size, err = 0, e
			return
		}

		d := cyclic.Decoder(vle.Reader(bytes.NewBuffer(data)))
		size, err = len(data), d.Value(out)
		return
	}

	<-done
	return
}

func (s *compactingArchive) Contains(id binary.ID) (found bool) {
	done := make(chan struct{})
	s.worker <- func() {
		_, found = s.records[id]
		close(done)
	}
	<-done
	return
}

func (s *compactingArchive) Close() {
	done := make(chan struct{})
	s.worker <- func() {
		s.index.Close()
		s.data.Close()
		close(done)
		close(s.worker)
		close(s.workerSync)
	}
	<-done
}

func (s *compactingArchive) Delete(id binary.ID, logger log.Logger) {
	done := make(chan struct{})
	s.worker <- func() {
		if prev, found := s.records[id]; found {
			delete(s.records, id)

			record := record{
				id: id,
				span: span{
					offset: 0,
					size:   0,
				},
			}
			s.waste += int64(prev.size)

			e := cyclic.Encoder(vle.Writer(s.index))
			record.encode(e)

			// If compaction is in-flight send the record to the compaction goroutine.
			if s.insertCh != nil {
				empty := make([]byte, 0)
				kv := keyValue{id: id, buffer: empty}
				s.insertCh <- kv
				if logger != nil {
					logger.Info("Delete sent to compactor.")
				}
			}
		}
		close(done)
	}
	<-done
	return
}

func (s *compactingArchive) compaction(l log.Logger) error {
	if s.compacting {
		return fmt.Errorf("compaction() call when compaction in progress")
	}
	s.compacting = true

	cdFn := filepath.Join(s.path, compactingDataFilename)
	dcompacting, err := os.OpenFile(cdFn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		s.compacting = false
		return err
	}

	ciFn := filepath.Join(s.path, compactingIndexFilename)
	icompacting, err := os.OpenFile(ciFn, os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0666)
	if err != nil {
		s.compacting = false
		return err
	}

	// We compact the file whilst allowing concurrent updates. Synchronous
	// operations need to happen on the worker and therefore need to be
	// relatively fast, but IO is allowed. The bulk of the compacting IO happens
	// in a separate goroutine. This is what happens:
	// - In worker: create a channel for concurrent updates.
	// - In worker: make an in-memory copy of the index in an array.
	// - In worker: start compactor go routine.
	// - In compactor: load each record from the data file and write to compacting_data.
	// - In compactor: write span to compacting_index.
	// - In compactor: whilst outputting records allow updates from channel.
	// When finished outputting:
	// - In compactor: synchronously send the completion function to the worker.
	// - In compactor: continue to process inserts until synchronous send is possible.
	// - In worker: process any remaining inserts.
	// - In worker: flip: close both
	// - In worker: flip: close both files, rename the originals "_old".
	//       Rename the new files to replace them. Get rid of the originals.
	// - In worker: put everything back to normal.

	// Create a channel for concurrent updates.
	insertCh := make(chan keyValue, 64)
	s.insertCh = insertCh

	// Make an in-memory copy of the records. By copying into an array we can manage
	// about 20k records per millisecond. Note archives expected to have large
	// data files rather than large indexes. Consider maintaining an offset ordered
	// array to avoid
	// the need for this copy.

	beforeCopyRecords := time.Now()
	recordsCopyArray := make([]record, len(s.records))
	i := 0
	for id, span := range s.records {
		recordsCopyArray[i] = record{id, span}
		i++
	}
	afterCopyRecords := time.Now()

	// Start compactor go routine.
	go func() {
		// Output records using buffered IO.
		iwriter := bufio.NewWriterSize(icompacting, 256<<10)
		dwriter := bufio.NewWriterSize(dcompacting, 256<<10)

		ie := cyclic.Encoder(vle.Writer(iwriter))

		size := int64(0)
		waste := int64(0)

		numInserts := 0
		sizeInserts := 0

		beforeWriteRecords := time.Now()
		recordsWritten := make(map[binary.ID]int64)

		didFirstInsertPath := false
		didSecondInsertPath := false
		didThirdInsertPath := false

		offset := int64(0)

		insertFun := func(insert keyValue) {
			if prevSize, ok := recordsWritten[insert.id]; ok {
				// As new records can be added we need to keep track of waste
				waste += int64(prevSize)
			}

			bufSize := int64(0)
			if len(insert.buffer) != 0 {
				bufSize = int64(len(insert.buffer))
				n, err := dwriter.Write(insert.buffer)
				if err != nil {
					panic(err)
				}
				if n != len(insert.buffer) {
					panic(fmt.Errorf("Byte count written was not as expected (%d/%d), error: %v.", n, len(insert.buffer), err))
				}
			}

			record := record{
				id: insert.id,
				span: span{
					offset: offset,
					size:   bufSize,
				},
			}
			offset += record.span.size

			err = record.encode(ie)
			if err != nil {
				panic(err)
			}

			size += bufSize
			recordsWritten[insert.id] = bufSize

			numInserts++
			sizeInserts += len(insert.buffer)
		}

		cleanupFun := func() {
			err := dcompacting.Close()
			if err != nil && l != nil {
				l.Warning("Failed to close compacting data file")
			}
			err = icompacting.Close()
			if err != nil && l != nil {
				l.Warning("Failed to close compacting index file")
			}
			err = os.Remove(cdFn)
			if err != nil && l != nil {
				l.Warning("Failed to remove compacting data file")
			}
			err = os.Remove(ciFn)
			if err != nil && l != nil {
				l.Warning("Failed to remove compacting index file")
			}
		}

		for _, record := range recordsCopyArray {
			data := make([]byte, record.span.size)
			if record.span.size != 0 {
				if n, err := s.data.ReadAt(data, record.span.offset); n != len(data) {
					if err != nil {
						panic(err)
					}
				}
			}

			insertFun(keyValue{id: record.id, buffer: data})

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
			err := dwriter.Flush()
			if err != nil {
				panic(err)
			}
			err = iwriter.Flush()
			if err != nil {
				panic(err)
			}
			afterFlush := time.Now()

			// Flip: close both files, rename the originals "_old".
			// Rename the new files to replace them. Get rid of the originals.
			err = icompacting.Close()
			if err != nil {
				panic(err)
			}

			err = dcompacting.Close()
			if err != nil {
				panic(err)
			}

			afterCatchup := time.Now()

			dFn := filepath.Join(s.path, dataFilename)
			iFn := filepath.Join(s.path, indexFilename)
			odFn := filepath.Join(s.path, oldDataFilename)
			oiFn := filepath.Join(s.path, oldIndexFilename)

			// If the index file is present then data is in either data or data_old
			// If not then if the data file is present the index is compacting_index
			// If neither exist then the data is in index_old and data_old
			err = os.Rename(dFn, odFn)
			if err != nil {
				panic(err)
			}
			err = os.Rename(iFn, oiFn)
			if err != nil {
				panic(err)
			}
			err = os.Rename(cdFn, dFn)
			if err != nil {
				panic(err)
			}
			err = os.Rename(ciFn, iFn)
			if err != nil {
				panic(err)
			}
			err = os.Remove(odFn)
			if err != nil {
				panic(err)
			}
			err = os.Remove(oiFn)
			if err != nil {
				panic(err)
			}

			data, err := os.OpenFile(dFn, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}
			index, err := os.OpenFile(iFn, os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
			if err != nil {
				panic(err)
			}

			// Put everything back to normal.
			s.data = data
			s.index = index
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
