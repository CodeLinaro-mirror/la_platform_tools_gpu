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

package log

import (
	"fmt"
	"sync/atomic"
	"time"
)

// FlushRequest is a signal to flush the Logger and is written to the output channel passed
// to Channel whenever Flush() is called. On receiving a FlushRequest, any pending messages
// should be flushed and the FlushRequest should be closed.
type FlushRequest chan struct{}

// Channel is an implementation of Logger interface that writes out an Entry to the specified chan
// for every message, and a FlushRequest when Flush is called.
func Channel(out chan<- interface{}) Logger {
	nextUid := uint32(1)
	return &channel{
		uid:     0,
		nextUid: &nextUid,
		scope:   "",
		out:     out,
	}
}

type channel struct {
	uid     uint32
	nextUid *uint32
	scope   string
	out     chan<- interface{}
}

func (c *channel) Info(msg string, args ...interface{}) {
	c.out <- Entry{
		Kind:      Info,
		Message:   fmt.Sprintf(msg, args...),
		Scope:     c.scope,
		Context:   c.uid,
		Timestamp: time.Now(),
	}
}

func (c *channel) Warning(msg string, args ...interface{}) {
	c.out <- Entry{
		Kind:      Warning,
		Message:   fmt.Sprintf(msg, args...),
		Scope:     c.scope,
		Context:   c.uid,
		Timestamp: time.Now(),
	}
}

func (c *channel) Error(msg string, args ...interface{}) {
	c.out <- Entry{
		Kind:      Error,
		Message:   fmt.Sprintf(msg, args...),
		Scope:     c.scope,
		Context:   c.uid,
		Timestamp: time.Now(),
	}
}

func (c *channel) Enter(name string) Logger {
	return &channel{
		uid:     c.uid,
		nextUid: c.nextUid,
		scope:   c.scope + name + " → ",
		out:     c.out,
	}
}

func (c *channel) Fork() Logger {
	return &channel{
		uid:     atomic.AddUint32(c.nextUid, 1),
		nextUid: c.nextUid,
		scope:   c.scope,
		out:     c.out,
	}
}

func (c *channel) Flush() {
	flush := make(FlushRequest)
	c.out <- flush
	<-flush
}

func (c *channel) Close() {
	c.Flush()
	close(c.out)
}
