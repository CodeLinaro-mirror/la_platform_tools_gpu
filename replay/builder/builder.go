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
	"fmt"
	"io"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/binary/endian"
	"android.googlesource.com/platform/tools/gpu/binary/flat"
	"android.googlesource.com/platform/tools/gpu/config"
	"android.googlesource.com/platform/tools/gpu/device"
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

type marker struct {
	instruction int     // first instruction index for this marker
	atom        atom.ID // the atom identifier
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
	mappedMemory    memory.RangeList
	instructions    []asm.Instruction
	decoders        []idPostDecoder
	stack           []stackItem
	architecture    device.Architecture

	// Remappings is a map of a arbitrary keys to pointers. Typically, this is
	// used as a map of observed values to values that are only known at replay
	// execution time, such as driver generated handles.
	// The Remappings field is not accessed by the Builder and can be used in any
	// way the developer requires.
	Remappings map[interface{}]value.Pointer
}

// New returns a newly constructed Builder configured to replay on a target
// with the specified Architecture.
func New(architecture device.Architecture) *Builder {
	return &Builder{
		constantMemory:  newConstantEncoder(architecture),
		heap:            allocator{alignment: uint64(architecture.PointerAlignment)},
		temp:            allocator{alignment: uint64(architecture.PointerAlignment)},
		resourceIDToIdx: map[binary.ID]uint32{},
		resources:       []protocol.ResourceInfo{},
		mappedMemory:    memory.RangeList{},
		instructions:    []asm.Instruction{},
		architecture:    architecture,
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

// Architecture returns the architecture for the target replay device.
func (b *Builder) Architecture() device.Architecture {
	return b.architecture
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
	alignment := uint64(b.architecture.PointerAlignment)
	ptrs = make([]value.Pointer, len(sizes))
	for _, s := range sizes {
		size = align(size, alignment)
		size += s
	}
	base := b.AllocateTemporaryMemory(size)
	offset := uint64(0)
	for i, s := range sizes {
		ptrs[i] = base.Offset(offset)
		offset += align(s, alignment)
	}
	return ptrs, size
}

// BeginAtom should be called before building any replay instructions.
func (b *Builder) BeginAtom(id atom.ID) {
	if id > 0x3ffffff {
		id = 0x3ffffff // Labels have 26 bit values.
	}
	b.instructions = append(b.instructions, asm.Label{
		Value: uint32(id),
	})
}

// EndAtom should be called after emitting the commands to replay a single atom.
// EndAtom frees all allocated temporary memory and clears the stack.
func (b *Builder) EndAtom() {
	b.temp.reset()
	c := len(b.stack)
	// Change calls that push an unused return value to discard the value.
	for _, i := range b.stack {
		if call, ok := b.instructions[i.idx].(asm.Call); ok && call.PushReturn {
			call.PushReturn = false
			b.instructions[i.idx] = call
			c--
		}
	}
	if c > 0 {
		b.instructions = append(b.instructions, asm.Pop{Count: uint32(c)})
	}
	b.stack = b.stack[:0]
}

// Buffer returns a pointer to a block of memory in holding the count number of
// previously pushed values.
// If all the values are constant, then the buffer will be held in the constant
// address-space, otherwise the buffer will be built in the temporary
// address-space.
func (b *Builder) Buffer(count int) value.Pointer {
	pointerSize := b.architecture.PointerSize
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

		size += ty.Size(pointerSize)
	}

	if dynamic {
		// At least one of the values was not from a Push()
		// Build the buffer in temporary memory at replay time.
		buf := b.AllocateTemporaryMemory(uint64(size))
		offset := size
		for i := 0; i < count; i++ {
			e := b.stack[len(b.stack)-1]
			offset -= e.ty.Size(pointerSize)
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

// Call will invoke the function f, popping all parameter values previously
// pushed to the stack with Push, starting with the first parameter. If f has
// a non-void return type, after invoking the function the return value of the
// function will be pushed on to the stack.
func (b *Builder) Call(f FunctionInfo) {
	for i := 0; i < f.Parameters; i++ {
		b.popStack()
	}
	push := f.ReturnType != protocol.TypeVoid
	if push {
		b.pushStack(f.ReturnType)
	}
	b.instructions = append(b.instructions, asm.Call{
		PushReturn: push,
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

// MapMemory adds rng as a memory range that needs allocating for replay.
func (b *Builder) MapMemory(rng memory.Range) {
	interval.Merge(&b.mappedMemory, rng.Span(), true)
}

// Write fills the memory range in capture address-space rng with the data
// of resourceID.
func (b *Builder) Write(rng memory.Range, resourceID binary.ID) {
	idx, found := b.resourceIDToIdx[resourceID]
	if !found {
		idx = uint32(len(b.resources))
		b.resourceIDToIdx[resourceID] = idx
		b.resources = append(b.resources, protocol.ResourceInfo{
			ID:   resourceID.String(),
			Size: uint32(rng.Size),
		})
	}
	b.instructions = append(b.instructions, asm.Resource{
		Index:       idx,
		Destination: rng.Base,
	})
	b.MapMemory(rng)
}

// Build compiles the replay instructions, returning a Payload that can be
// sent to the replay virtual-machine and a ResponseDecoder for interpreting
// the responses.
func (b *Builder) Build(logger log.Logger) (protocol.Payload, ResponseDecoder, error) {
	logger = logger.Enter("Build")
	if config.DebugReplayBuilder {
		logger.Infof("Instruction count: %d", len(b.instructions))
	}

	vml := b.layoutVolatileMemory(logger)

	byteOrder := b.architecture.ByteOrder

	opcodes := &bytes.Buffer{}
	e := flat.Encoder(endian.Writer(opcodes, byteOrder))
	id := uint32(0)
	for _, i := range b.instructions {
		if label, ok := i.(asm.Label); ok {
			id = label.Value
		}
		if err := i.Encode(vml, e); err != nil {
			err = fmt.Errorf("Encode %T failed for atom with id %v: %v", i, id, err)
			return protocol.Payload{}, nil, err
		}
	}

	payload := protocol.Payload{
		StackSize:          uint32(512), // TODO: Calculate stack size
		VolatileMemorySize: uint32(vml.size),
		Constants:          b.constantMemory.data,
		Resources:          b.resources,
		Opcodes:            opcodes.Bytes(),
	}

	if config.DebugReplayBuilder {
		logger.Infof("Stack size:           0x%x", payload.StackSize)
		logger.Infof("Volatile memory size: 0x%x", payload.VolatileMemorySize)
		logger.Infof("Constant memory size: 0x%x", len(payload.Constants))
		logger.Infof("Opcodes size:         0x%x", len(payload.Opcodes))
		logger.Infof("Resource count:         %d", len(payload.Resources))
	}

	responseDecoder := func(r io.Reader) <-chan Postback {
		d := flat.Decoder(endian.Reader(r, byteOrder))
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
	return payload, responseDecoder, nil
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
	remappedBases := make([]uint64, len(b.mappedMemory))
	for i, m := range b.mappedMemory {
		remappedBases[i] = remapBase + remapped.alloc(m.Size)
	}

	size := remapBase + remapped.size
	vml := &volatileMemoryLayout{
		mappedMemory:  b.mappedMemory,
		tempBase:      tempBase,
		remappedBases: remappedBases,
		size:          size,
	}

	if config.DebugReplayBuilder {
		logger.Infof("Volatile memory layout: [0x%x, 0x%x]", 0, size-1)
		logger.Infof("  Heap:      [0x%x, 0x%x]", 0, tempBase-1)
		logger.Infof("  Temporary: [0x%x, 0x%x]", tempBase, remapBase-1)
		logger.Infof("  Remapped:  [0x%x, 0x%x]", remapBase, size-1)
		for _, m := range b.mappedMemory {
			logger.Infof("    Block:   %v", m)
		}
	}

	return vml
}

type volatileMemoryLayout struct {
	mappedMemory  memory.RangeList // Mapped memory ranges.
	tempBase      uint64           // Base address of the temp space.
	remappedBases []uint64         // Base address for each remapped entry in mappedMemory.
	size          uint64           // Total size of volatile memory.
}

// TranslateTemporaryPointer implements the PointerResolver interface method in
// the replay/value package.
func (l volatileMemoryLayout) TranslateTemporaryPointer(offset uint64) (uint64, error) {
	return l.tempBase + offset, nil
}

// TranslateCapturePointer implements the PointerResolver interface method in
// the replay/value package.
func (l volatileMemoryLayout) TranslateCapturePointer(offset uint64) (uint64, error) {
	bufferIdx := interval.IndexOf(&l.mappedMemory, offset)
	if bufferIdx < 0 {
		return 0, fmt.Errorf("Pointer 0x%x was not found in the observed memory ranges", offset)
	}
	bufferStart := l.mappedMemory[bufferIdx].First()
	return l.remappedBases[bufferIdx] + offset - uint64(bufferStart), nil
}
