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

package trace

import (
	"flag"
	"fmt"
	"io"
	"os"

	"android.googlesource.com/platform/tools/gpu/gapii"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/tools/verbs"
)

var (
	verb = &verbs.Verb{
		Name:      "info",
		ShortHelp: "Prints information about a gfx trace capture file",
		Run:       doInfo,
	}
)

func init() {
	verbs.Register(verb)
}

type countingReader struct {
	from  io.Reader
	count int
}

func (c *countingReader) Read(buf []byte) (int, error) {
	size, err := c.from.Read(buf)
	c.count += size
	return size, err
}

func doInfo(flags flag.FlagSet) error {
	if flags.NArg() != 1 {
		return verbs.Usage("Exactly one gfx trace file expected, got %d", flags.NArg())
	}
	filename := flags.Arg(0)
	fmt.Println("reading file ", filename)
	fstat, err := os.Stat(filename)
	if err != nil {
		return err
	}
	fmt.Println("total file size ", fstat.Size())
	f, err := os.Open(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	info := os.Stdout
	if verbs.Verbosity > 0 {
		info = nil
	}
	logger := log.Writer(info, os.Stdout, os.Stderr, nil)
	defer log.Close(logger)

	in := &countingReader{from: f}
	list, err := gapii.ReadCapture(fstat.Name(), in, logger)
	fmt.Println("Got ", len(list.Atoms), " atoms from ", in.count, "bytes")
	if err != nil {
		return err
	}
	return err
}
