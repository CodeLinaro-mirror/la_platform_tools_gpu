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

package test

import (
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/log"
)

type RPC interface {
	Add(a uint32, b uint32, l log.Logger) (uint32, error)
	EnumToString(e Enum, l log.Logger) (string, error)
	GetListNodeChain(l log.Logger) (*ListNode, error)
	GetListNodeChainArray(l log.Logger) ([]*ListNode, error)
	GetResource(l log.Logger) (ResourceId, error)
	GetSingleListNode(l log.Logger) (*ListNode, error)
	GetStruct(l log.Logger) (Struct, error)
	ResolveResource(r ResourceId, l log.Logger) (Resource, error)
	SetStruct(s Struct, l log.Logger) error
	UseResource(r ResourceId, l log.Logger) error
}

type Resolver struct {
	Database database.Database
}

type ResourceId struct {
	binary.Generate `handle:"Resource"`
	ID              binary.ID
}

// Enum Enum
type Enum int

const (
	EnumOne   Enum = 1
	EnumTwo   Enum = 2
	EnumThree Enum = 3
)

// Interface Base
type Base interface {
	binary.Object
	GetName() string
}

// Class Resource
type Resource struct {
	binary.Generate
	Int    uint32
	Float  float32
	String string
}

// Class Struct
type Struct struct {
	binary.Generate
	String string
	U32    uint32
	Enum   Enum
}

// Class ListNode
type ListNode struct {
	binary.Generate
	Name string
	Next *ListNode
}
