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
	"fmt"
	"reflect"
	"testing"

	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/multiplexer"
	"android.googlesource.com/platform/tools/gpu/ringbuffer"
)

var testStruct = Struct{
	String: "This is a struct",
	U32:    42,
	Enum:   EnumOne,
}
var testResourceId = ResourceId{ID: binary.NewID([]byte("Test resource id"))}
var testResource = Resource{Int: 10, Float: 20, String: "30"}
var testSingleListNode = &ListNode{Name: "Single ListNode"}
var testListNodeChain = &ListNode{Name: "ListNodeA", Next: &ListNode{Name: "ListNodeB", Next: &ListNode{Name: "ListNodeC"}}}
var testListNodeChainArray = []*ListNode{
	testListNodeChain, testListNodeChain, testListNodeChain,
}

type server struct {
	calls []string
	err   error
}

func (s *server) Add(a uint32, b uint32, l log.Logger) (uint32, error) {
	s.calls = append(s.calls, fmt.Sprintf("Add(%d, %d)", a, b))
	return a + b, s.err
}

func (s *server) EnumToString(e Enum, l log.Logger) (string, error) {
	s.calls = append(s.calls, fmt.Sprintf("EnumToString(%v)", e))
	return fmt.Sprint(e), s.err
}

func (s *server) GetStruct(l log.Logger) (Struct, error) {
	s.calls = append(s.calls, "GetStruct()")
	return testStruct, s.err
}

func (s *server) SetStruct(str Struct, l log.Logger) error {
	s.calls = append(s.calls, fmt.Sprintf("SetStruct(%v)", str))
	return s.err
}

func (s *server) GetResource(l log.Logger) (ResourceId, error) {
	s.calls = append(s.calls, "GetResource()")
	return testResourceId, s.err
}

func (s *server) ResolveResource(id ResourceId, l log.Logger) (Resource, error) {
	s.calls = append(s.calls, fmt.Sprintf("ResolveResource(%v)", id))
	return testResource, s.err
}

func (s *server) UseResource(r ResourceId, l log.Logger) error {
	s.calls = append(s.calls, fmt.Sprintf("UseResource(%v)", r))
	return s.err
}

func (s *server) GetSingleListNode(l log.Logger) (*ListNode, error) {
	s.calls = append(s.calls, "GetSingleListNode()")
	return testSingleListNode, s.err
}

func (s *server) GetListNodeChain(l log.Logger) (*ListNode, error) {
	s.calls = append(s.calls, "GetListNodeChain()")
	return testListNodeChain, s.err
}

func (s *server) GetListNodeChainArray(l log.Logger) ([]*ListNode, error) {
	s.calls = append(s.calls, "GetListNodeChainArray()")
	return testListNodeChainArray, s.err
}

func create(t *testing.T) (RPC, *server) {
	l := log.Enter(log.Testing(t), "Server")
	mtu := 64
	s2c, c2s := ringbuffer.New(64), ringbuffer.New(64)
	server := &server{}
	client := NewClient(multiplexer.New(s2c, c2s, mtu, nil), nil)
	BindServer(c2s, s2c, mtu, l, server)
	return client, server
}

func verifyCalls(t *testing.T, s *server, err error, calls ...string) {
	if err != nil {
		t.Fatalf("Expected success, got error: %v", err)
	}
	if len(calls) != len(s.calls) {
		t.Fatalf("Expected %d calls, got %d", len(calls), len(s.calls))
	}
	for i, expected := range calls {
		got := s.calls[i]
		if expected != got {
			t.Errorf("Call %d wrong, expected %s got %s", i, expected, got)
		}
	}
}

func verifyResult(t *testing.T, expected interface{}, got interface{}) {
	if !reflect.DeepEqual(expected, got) {
		t.Errorf("Expected result %v got %v", expected, got)
	}
}

func TestInitialConditions(t *testing.T) {
	_, server := create(t)
	verifyCalls(t, server, nil)
}

func TestCallAdd(t *testing.T) {
	client, server := create(t)
	res, err := client.Add(1, 2, log.Testing(t))
	verifyCalls(t, server, err, "Add(1, 2)")
	verifyResult(t, uint32(3), res)
}

func TestCallEnumToString(t *testing.T) {
	client, server := create(t)
	res, err := client.EnumToString(EnumOne, log.Testing(t))
	verifyCalls(t, server, err, fmt.Sprintf("EnumToString(%v)", EnumOne))
	verifyResult(t, "One", res)
}

func TestCallGetStruct(t *testing.T) {
	client, server := create(t)
	res, err := client.GetStruct(log.Testing(t))
	verifyCalls(t, server, err, "GetStruct()")
	verifyResult(t, testStruct, res)
}

func TestCallSetStruct(t *testing.T) {
	client, server := create(t)
	client.SetStruct(testStruct, log.Testing(t))
	verifyCalls(t, server, nil, fmt.Sprintf("SetStruct(%v)", testStruct))
}

func TestCallGetResource(t *testing.T) {
	client, server := create(t)
	res, err := client.GetResource(log.Testing(t))
	verifyCalls(t, server, err, "GetResource()")
	verifyResult(t, testResourceId, res)
}

func TestCallResolveResource(t *testing.T) {
	client, server := create(t)
	res, err := client.ResolveResource(testResourceId, log.Testing(t))
	verifyCalls(t, server, err, fmt.Sprintf("ResolveResource(%v)", testResourceId))
	verifyResult(t, testResource, res)
}

func TestCallUseResource(t *testing.T) {
	client, server := create(t)
	client.UseResource(testResourceId, log.Testing(t))
	verifyCalls(t, server, nil, fmt.Sprintf("UseResource(%v)", testResourceId))
}

func TestCallGetSingleListNode(t *testing.T) {
	client, server := create(t)
	res, err := client.GetSingleListNode(log.Testing(t))
	verifyCalls(t, server, err, "GetSingleListNode()")
	verifyResult(t, testSingleListNode, res)
}

func TestCallGetListNodeChain(t *testing.T) {
	client, server := create(t)
	res, err := client.GetListNodeChain(log.Testing(t))
	verifyCalls(t, server, err, "GetListNodeChain()")
	verifyResult(t, testListNodeChain, res)
}

func TestCallGetListNodeChainArray(t *testing.T) {
	client, server := create(t)
	res, err := client.GetListNodeChainArray(log.Testing(t))
	verifyCalls(t, server, err, "GetListNodeChainArray()")
	verifyResult(t, testListNodeChainArray, res)
	if !reflect.DeepEqual(res[0], res[1]) || !reflect.DeepEqual(res[1], res[2]) {
		t.Errorf("Array values differ")
	}
}

// TODO: Test errors
