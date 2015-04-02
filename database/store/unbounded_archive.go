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

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/cyclic"
	"android.googlesource.com/platform/tools/gpu/binary/vle"
	"android.googlesource.com/platform/tools/gpu/log"
)

type unboundedArchive struct {
	worker  chan<- func()
	index   *os.File
	data    *os.File
	records map[binary.ID]span
}

// CreateUnboundedArchive returns a store that appends new resources, and never
// attempts to reclaim obsolete ones.
func CreateUnboundedArchive(path string) Store {
	os.MkdirAll(path, 0755)

	index, err := os.OpenFile(filepath.Join(path, "index"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	data, err := os.OpenFile(filepath.Join(path, "data"), os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(err)
	}

	records := make(map[binary.ID]span)

	// Use a buffered reader to quickly read the active records
	d := cyclic.Decoder(vle.Reader(bufio.NewReaderSize(index, 256<<10)))
	for {
		var r record
		if err := r.decode(d); err != io.EOF {
			if err != nil {
				panic(err)
			}
			records[r.id] = r.span
		} else {
			break
		}
	}

	worker := make(chan func(), 64)
	go func() {
		for f := range worker {
			f()
		}
	}()

	return &unboundedArchive{
		worker:  worker,
		index:   index,
		data:    data,
		records: records,
	}
}

// Store compliance
func (s unboundedArchive) Store(id binary.ID, _ binary.Object, data []byte, logger log.Logger) (err error) {
	if logger != nil {
		logger = logger.Enter("unboundedArchive.Store")
		logger.Info("id: %s size: %d", id, len(data))
	}
	s.worker <- func() {
		offset, err := s.data.Seek(0, os.SEEK_END)
		if err != nil {
			panic(err)
		}

		n, err := s.data.Write(data)
		if err != nil {
			panic(err)
		}
		if n != len(data) {
			panic(fmt.Errorf("Byte count written was not as expected. (%d/%d)", n, len(data)))
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
	}
	return nil
}

func (s unboundedArchive) Load(id binary.ID, logger log.Logger, out binary.Object) (size int, err error) {
	if logger != nil {
		logger = logger.Enter("unboundedArchive.Load")
		logger.Info("Loading from unboundedArchive: %s", id)
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

func (s unboundedArchive) Contains(id binary.ID) (found bool) {
	done := make(chan struct{})
	s.worker <- func() {
		_, found = s.records[id]
		close(done)
	}
	<-done
	return
}

func (s unboundedArchive) Close() {
	done := make(chan struct{})
	s.worker <- func() {
		s.index.Close()
		s.data.Close()
		close(done)
		close(s.worker)
	}
	<-done
}
