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

// Package builder contains the Builder type to build replay payloads.
package builder

import (
	"bytes"
	eb "encoding/binary"
	"errors"
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/interval"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay/asm"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

type stackItem struct {
	ty  protocol.Type // Type of the item.
	idx int           // Index of the op that generated this.
}

type idPostDecoder struct {
	id atom.ID
	pd PostDecoder
}

// Postback holds the information for a single atom's postback data.
type Postback struct {
	ID    atom.ID     // The associated atom for this Postback.
	Data  interface{} // The postback data. Nil if Error is non-nil.
	Error error       // Error raised decoding the postback, or nil if there was no error.
}

// ResponseDecoder decodes all postback responses from the replay virtual
// machine, writing each to the returned chan. The chan will be closed once all
// postbacks have been read, or after the first Postback with a non-nil Error.
type ResponseDecoder func(r io.Reader) <-chan Postback

// PostDecoder decodes a single atom's postback, returning the postback data or
// an error. The PostDecoder must decode all the data that was issued in the
// Post call before returning.
type PostDecoder func(binary.Decoder) (interface{}, error)

// Builder is used to build the Payload to send to the replay virtual machine.
// The builder has a number of methods for mutating the virtual machine stack,
// invoking functions and posting back data.
type Builder struct {
	constantMemory  *constantEncoder
	heap, temp      allocator
	resourceIDToIdx map[binary.ID]uint32
	resources       []protocol.ResourceInfo
	observedRanges  memory.RangeList
	instructions    []asm.Instruction
	decoders        []idPostDecoder
	stack           []stackItem
	ptrSize         int
	ptrAlignment    int
	byteOrder       eb.ByteOrder

	// Remappings is a map of a arbitrary keys to pointers. Typically, this is
	// used as a map of observed values to values that are only known at replay
	// execution time, such as driver generated handles.
	// The Remappings field is not accessed by the Builder and can be used in any
	// way the developer requires.
	Remappings map[interface{}]value.Pointer
}

// New returns a newly constructed Builder configured to replay on a target
// architecture that has a pointer size of ptrSize bytes and and alignment of
// ptrAlignment bytes, with byte ordering that matches byteOrder.
func New(ptrSize, ptrAlignment int, byteOrder eb.ByteOrder) *Builder {
	return &Builder{
		constantMemory:  newConstantEncoder(ptrAlignment, byteOrder),
		heap:            allocator{alignment: uint64(ptrAlignment)},
		temp:            allocator{alignment: uint64(ptrAlignment)},
		resourceIDToIdx: map[binary.ID]uint32{},
		resources:       []protocol.ResourceInfo{},
		observedRanges:  memory.RangeList{},
		instructions:    []asm.Instruction{},
		ptrSize:         ptrSize,
		ptrAlignment:    ptrAlignment,
		byteOrder:       byteOrder,
		Remappings:      make(map[interface{}]value.Pointer),
	}
}

func (b *Builder) pushStack(t protocol.Type) {
	b.stack = append(b.stack, stackItem{t, len(b.instructions)})
}

func (b *Builder) popStack() {
	b.stack = b.stack[:len(b.stack)-1]
}

func (b *Builder) removeInstruction(at int) {
	if at == len(b.instructions)-1 {
		b.instructions = b.instructions[:at]
	} else {
		b.instructions[at] = asm.Nop{}
	}
}

// PointerSize returns the size of a pointer in bytes for the target replay
// architecture.
func (b *Builder) PointerSize() int {
	return b.ptrSize
}

// PointerAlignment returns the required alignment of a pointer in bytes for the
// replay target architecture.
func (b *Builder) PointerAlignment() int {
	return b.ptrAlignment
}

// AllocateMemory allocates and returns a pointer to a block of memory in the
// volatile address-space big enough to hold size bytes. The memory will be
// allocated for the entire replay duration and cannot be freed.
func (b *Builder) AllocateMemory(size uint64) value.Pointer {
	return value.VolatilePointer(b.heap.alloc(size))
}

// AllocateTemporaryMemory allocates and returns a pointer to a block of memory
// in the temporary volatile address-space big enough to hold size bytes. The
// memory block will be freed on the next call to EndAtom, upon which reading or
// writing to this memory will result in undefined behavior.
func (b *Builder) AllocateTemporaryMemory(size uint64) value.Pointer {
	return value.VolatileTemporaryPointer(b.temp.alloc(size))
}

// AllocateTemporaryMemoryChunks allocates a contiguous block of memory in the
// temporary volatile address-space big enough to hold all the specified chunks
// sizes, in sequential order. AllocateTemporaryMemoryChunks returns a pointer
// to each of the allocated chunks and the size of the entire allocation. The
// allocation block will be freed on the next call to EndAtom, upon which
// reading or writing to this memory will result in undefined behavior.
func (b *Builder) AllocateTemporaryMemoryChunks(sizes []uint64) (ptrs []value.Pointer, size uint64) {
	ptrs = make([]value.Pointer, len(sizes))
	for _, s := range sizes {
		size = align(size, uint64(b.PointerAlignment()))
		size += s
	}
	base := b.AllocateTemporaryMemory(size)
	offset := uint64(0)
	for i, s := range sizes {
		ptrs[i] = base.Offset(offset)
		offset += align(s, uint64(b.PointerAlignment()))
	}
	return ptrs, size
}

// EndAtom should be called after emitting the commands to replay a single atom.
// EndAtom frees all allocated temporary memory, and asserts the stack should be
// empty.
func (b *Builder) EndAtom() {
	b.temp.reset()
	if len(b.stack) > 0 {
		msg := "Stack not empty at end of function:\n"
		for i, e := range b.stack {
			msg += fmt.Sprintf("(%d) %s %#v", i, e.ty.String(), b.instructions[e.idx])
		}
		panic(errors.New(msg))
	}
}

// Buffer returns a pointer to a block of memory in holding the count number of
// previously pushed values.
// If all the values are constant, then the buffer will be held in the constant
// address-space, otherwise the buffer will be built in the temporary
// address-space.
func (b *Builder) Buffer(count int) value.Pointer {
	dynamic := false
	size := 0

	// Examine the stack to see where these values came from
	for i := 0; i < count; i++ {
		e := b.stack[len(b.stack)-i-1]
		op := b.instructions[e.idx]
		ty := e.ty
		if _, isPush := op.(asm.Push); !isPush {
			// Values that have made their way on to the stack from non-constant
			// sources cannot be put in the constant buffer.
			dynamic = true
		}
		switch ty {
		case protocol.TypeConstantPointer, protocol.TypeVolatilePointer:
			// Pointers cannot be put into the constant buffer as they are remapped
			// by the VM
			dynamic = true
		}

		size += ty.Size(b.PointerSize())
	}

	if dynamic {
		// At least one of the values was not from a Push()
		// Build the buffer in temporary memory at replay time.
		buf := b.AllocateTemporaryMemory(uint64(size))
		offset := size
		for i := 0; i < count; i++ {
			e := b.stack[len(b.stack)-1]
			offset -= e.ty.Size(b.PointerSize())
			b.Store(buf.Offset(uint64(offset)))
		}
		return buf
	} else {
		// All the values are constant.
		// Move the pushed values into a constant memory buffer.
		values := make([]value.Value, count)
		for i := 0; i < count; i++ {
			e := b.stack[len(b.stack)-1]
			values[count-i-1] = b.instructions[e.idx].(asm.Push).Value
			b.removeInstruction(e.idx)
			b.popStack()
		}
		return b.constantMemory.writeValues(values...)
	}
}

// String returns a pointer to a block of memory in the constant address-space
// holding the string s. The string will be stored with a null-terminating byte.
func (b *Builder) String(s string) value.Pointer {
	return b.constantMemory.writeString(s)
}

// CallPush will invoke the function f, popping all parameter values previously
// pushed to the stack with Push, starting with the first parameter. After
// invoking the function the return value of the function will be pushed on to
// the stack. If f has a void return type then CallPush will panic.
func (b *Builder) CallPush(f FunctionInfo) {
	if f.ReturnType == protocol.TypeVoid {
		panic("CallPush called with a void returning function")
	}
	for i := 0; i < f.Parameters; i++ {
		b.popStack()
	}
	b.pushStack(f.ReturnType)
	b.instructions = append(b.instructions, asm.Call{
		PushReturn: true,
		FunctionID: f.ID,
	})
}

// CallNoPush will invoke the function f, popping all parameter values
// previously pushed to the stack with Push, starting with the first parameter.
// Unlike CallPush the return value of the function will not be pushed on to the
// stack and functions with void return types are accepted.
func (b *Builder) CallNoPush(f FunctionInfo) {
	for i := 0; i < f.Parameters; i++ {
		b.popStack()
	}
	b.instructions = append(b.instructions, asm.Call{
		PushReturn: false,
		FunctionID: f.ID,
	})
}

// Copy pops the target address and then the source address from the top of the
// stack, and then copies Count bytes from source to target.
func (b *Builder) Copy(size uint64) {
	b.popStack()
	b.popStack()
	b.instructions = append(b.instructions, asm.Copy{
		Count: size,
	})
}

// Clone makes a copy of the n-th element from the top of the stack and pushes
// the copy to the top of the stack.
func (b *Builder) Clone(index int) {
	b.pushStack(b.stack[len(b.stack)-1-index].ty)
	b.instructions = append(b.instructions, asm.Clone{
		Index: index,
	})
}

// Load loads the value of type ty from addr and then pushes the loaded value to
// the top of the stack.
func (b *Builder) Load(ty protocol.Type, addr value.Pointer) {
	if !addr.IsValid() {
		panic(fmt.Errorf("Pointer address %v is not valid", addr))
	}
	b.pushStack(ty)
	b.instructions = append(b.instructions, asm.Load{
		DataType: ty,
		Source:   addr,
	})
}

// Store pops the value from the top of the stack and writes the value to addr.
func (b *Builder) Store(addr value.Pointer) {
	if !addr.IsValid() {
		panic(fmt.Errorf("Pointer address %v is not valid", addr))
	}
	b.popStack()
	b.instructions = append(b.instructions, asm.Store{
		Destination: addr,
	})
}

// Strcpy pops the source address then the target address from the top of the
// stack, and then copies at most maxCount-1 bytes from source to target. If
// maxCount is greater than the source string length, then the target will be
// padded with 0s. The destination buffer will always be 0-terminated.
func (b *Builder) Strcpy(maxCount uint64) {
	b.popStack()
	b.popStack()
	b.instructions = append(b.instructions, asm.Strcpy{
		MaxCount: maxCount,
	})
}

// Post posts size bytes from addr to the decoder d. The decoder d must consume
// all size bytes before returning; failure to do this will corrupt all
// subsequent postbacks.
func (b *Builder) Post(addr value.Pointer, size uint64, id atom.ID, d PostDecoder) {
	if !addr.IsValid() {
		panic(fmt.Errorf("Pointer address %v is not valid", addr))
	}
	b.instructions = append(b.instructions, asm.Post{
		Source: addr,
		Size:   size,
	})
	b.decoders = append(b.decoders, idPostDecoder{id, d})
}

// Push pushes val to the top of the stack.
func (b *Builder) Push(val value.Value) {
	if pv, ok := val.(value.Pointer); ok && !pv.IsValid() {
		panic(fmt.Errorf("PointerValue %v is not valid", val))
	}
	b.pushStack(val.Type())
	b.instructions = append(b.instructions, asm.Push{
		Value: val,
	})
}

// Pop removes the top count values from the top of the stack.
func (b *Builder) Pop(count uint32) {
	for i := uint32(0); i < count; i++ {
		b.popStack()
	}
	b.instructions = append(b.instructions, asm.Pop{
		Count: count,
	})
}

// Observation fills the memory range in capture address-space rng with the data
// of resourceID.
func (b *Builder) Observation(rng memory.Range, resourceID binary.ID) {
	idx, found := b.resourceIDToIdx[resourceID]
	if !found {
		idx = uint32(len(b.resources))
		b.resourceIDToIdx[resourceID] = idx
		b.resources = append(b.resources, protocol.ResourceInfo{
			ID:   resourceID,
			Size: uint32(rng.Size),
		})
	}
	b.instructions = append(b.instructions, asm.Resource{
		Index:       idx,
		Destination: rng.Base,
	})
	interval.Merge(&b.observedRanges, rng.Span())
}

// Build compiles the replay instructions, returning a Payload that can be
// sent to the replay virtual-machine and a ResponseDecoder for interpreting
// the responses.
func (b *Builder) Build(logger log.Logger) (protocol.Payload, ResponseDecoder) {
	logger = logger.Enter("Build")
	vml := b.layoutVolatileMemory(logger)

	opcodes := &bytes.Buffer{}
	e := flat.Encoder(endian.Writer(opcodes, b.byteOrder))
	for _, i := range b.instructions {
		i.Encode(vml, e)
	}

	payload := protocol.Payload{
		StackSize:          uint32(512), // TODO: Calculate stack size
		VolatileMemorySize: uint32(vml.size),
		Constants:          b.constantMemory.data,
		Resources:          b.resources,
		Opcodes:            opcodes.Bytes(),
	}

	logger.Info("Stack size:           0x%x", payload.StackSize)
	logger.Info("Volatile memory size: 0x%x", payload.VolatileMemorySize)
	logger.Info("Constant memory size: 0x%x", len(payload.Constants))
	logger.Info("Opcodes size:         0x%x", len(payload.Opcodes))
	logger.Info("Resource count:         %d", len(payload.Resources))

	responseDecoder := func(r io.Reader) <-chan Postback {
		d := flat.Decoder(endian.Reader(r, b.byteOrder))
		c := make(chan Postback, 8)
		go func() {
			for _, p := range b.decoders {
				data, err := p.pd(d)
				c <- Postback{
					ID:    p.id,
					Data:  data,
					Error: err,
				}
				if err != nil {
					break
				}
			}
			close(c)
		}()
		return c
	}
	return payload, responseDecoder
}

func (b *Builder) layoutVolatileMemory(logger log.Logger) *volatileMemoryLayout {
	// Volatile memory layout:
	//
	//  low ┌──────────────────┐
	//      │       heap       │
	//      ├──────────────────┤
	//      │       temp       │
	//      ├──────────────────┤
	//      │ remapped range 0 │
	//      ├──────────────────┤
	//      ├──────────────────┤
	//      │ remapped range N │
	// high └──────────────────┘

	tempBase := b.heap.size
	remapBase := tempBase + b.temp.size

	remapped := allocator{alignment: b.heap.alignment}
	remappedBases := make([]uint64, len(b.observedRanges))
	for i, m := range b.observedRanges {
		remappedBases[i] = remapBase + remapped.alloc(m.Size)
	}

	size := remapBase + remapped.size
	vml := &volatileMemoryLayout{
		observedRanges: b.observedRanges,
		tempBase:       tempBase,
		remappedBases:  remappedBases,
		size:           size,
	}

	logger.Info("Volatile memory layout: [0x%x, 0x%x]", 0, size-1)
	logger.Info("  Heap:      [0x%x, 0x%x]", 0, tempBase-1)
	logger.Info("  Temporary: [0x%x, 0x%x]", tempBase, remapBase-1)
	logger.Info("  Remapped:  [0x%x, 0x%x]", remapBase, size-1)
	return vml
}

type volatileMemoryLayout struct {
	observedRanges memory.RangeList // Observed captured memory ranges.

	tempBase      uint64   // Base address of the temp space.
	remappedBases []uint64 // Base address for each remapped entry in observedRanges.
	size          uint64   // Total size of volatile memory.
}

// TranslateTemporaryPointer implements the PointerResolver interface method in
// the replay/value package.
func (l volatileMemoryLayout) TranslateTemporaryPointer(offset uint64) uint64 {
	return l.tempBase + offset
}

// TranslateCapturePointer implements the PointerResolver interface method in
// the replay/value package.
func (l volatileMemoryLayout) TranslateCapturePointer(offset uint64) uint64 {
	bufferIdx := interval.IndexOf(&l.observedRanges, offset)
	if bufferIdx < 0 {
		panic(fmt.Errorf("Pointer 0x%x was not found in the observed memory ranges", offset))
	}
	bufferStart := l.observedRanges[bufferIdx].First()
	return l.remappedBases[bufferIdx] + offset - uint64(bufferStart)
}
