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

// Package rpc implements a remote procedure call system.
//
// RPC uses the binary package for serialization and the multiplexer package to
// merge in-flight requests onto a single stream.
package rpc

// binary: java.source = base/rpclib/src/main/java
// binary: java.package = com.android.tools.rpclib.rpccore
// binary: java.member_prefix = m

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/binary"
)

var header = [4]byte{'r', 'p', 'c', '0'}

// ErrInvalidHeader is returned when either client or server detects an
// incorrectly formed rpc header.
var ErrInvalidHeader = NewError("Invalid RPC header")

// NewError is used to create new rpc error objects with the specified human readable message.
func NewError(msg string, args ...interface{}) *Error {
	return &Error{message: fmt.Sprintf(msg, args...)}
}

// Error is an implementation of error that can be sent over the wire.
type Error struct {
	binary.Generate `java:"RpcError"`
	message         string
}

func (e *Error) Error() string { return e.message }
