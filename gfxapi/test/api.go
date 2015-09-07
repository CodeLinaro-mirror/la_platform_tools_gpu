////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package test

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
	"android.googlesource.com/platform/tools/gpu/replay/builder"
	"android.googlesource.com/platform/tools/gpu/replay/protocol"
	"android.googlesource.com/platform/tools/gpu/replay/value"
)

type remapped uint32

// U8ᵖ is a pointer to a uint8 element.
type U8ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewU8ᵖ returns a U8ᵖ that points to addr in the application pool.
func NewU8ᵖ(addr uint64) U8ᵖ {
	return U8ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that U8ᵖ points to.
func (p U8ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the uint8 element at the pointer.
func (p U8ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint8 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the uint8 element at the pointer.
func (p U8ᵖ) Write(value uint8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]uint8{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p U8ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p U8ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p U8ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new U8ˢ from the pointer using start and end indices.
func (p U8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U8ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// U16ᵖ is a pointer to a uint16 element.
type U16ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewU16ᵖ returns a U16ᵖ that points to addr in the application pool.
func NewU16ᵖ(addr uint64) U16ᵖ {
	return U16ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that U16ᵖ points to.
func (p U16ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Read reads and returns the uint16 element at the pointer.
func (p U16ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint16 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the uint16 element at the pointer.
func (p U16ᵖ) Write(value uint16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]uint16{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p U16ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p U16ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p U16ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new U16ˢ from the pointer using start and end indices.
func (p U16ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U16ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U16ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// U32ᵖ is a pointer to a uint32 element.
type U32ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewU32ᵖ returns a U32ᵖ that points to addr in the application pool.
func NewU32ᵖ(addr uint64) U32ᵖ {
	return U32ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that U32ᵖ points to.
func (p U32ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the uint32 element at the pointer.
func (p U32ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint32 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the uint32 element at the pointer.
func (p U32ᵖ) Write(value uint32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]uint32{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p U32ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p U32ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p U32ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new U32ˢ from the pointer using start and end indices.
func (p U32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Intᵖ is a pointer to a int64 element.
type Intᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewIntᵖ returns a Intᵖ that points to addr in the application pool.
func NewIntᵖ(addr uint64) Intᵖ {
	return Intᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Intᵖ points to.
func (p Intᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(ϟs.Architecture.IntegerSize)
}

// Read reads and returns the int64 element at the pointer.
func (p Intᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the int64 element at the pointer.
func (p Intᵖ) Write(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]int64{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Intᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Intᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p Intᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new Intˢ from the pointer using start and end indices.
func (p Intᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Intˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Charᵖ is a pointer to a byte element.
type Charᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCharᵖ returns a Charᵖ that points to addr in the application pool.
func NewCharᵖ(addr uint64) Charᵖ {
	return Charᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Charᵖ points to.
func (p Charᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the byte element at the pointer.
func (p Charᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) byte {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the byte element at the pointer.
func (p Charᵖ) Write(value byte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]byte{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Charᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Charᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p Charᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// StringSlice returns a slice starting at p and ending at the first 0 byte null-terminator.
func (p Charᵖ) StringSlice(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Charˢ {
	i, d := uint64(0), ϟs.MemoryDecoder(ϟs.Memory[p.Pointer.Pool].At(p.Address), ϟd, ϟl)
	for {
		i++
		if b, _ := d.Uint8(); b == 0 {
			return Charˢ(p.Slice(0, i, ϟs))
		}
	}
}

// Slice returns a new Charˢ from the pointer using start and end indices.
func (p Charᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Charᵖᵖ is a pointer to a Charᵖ element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address of an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type Charᵖᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCharᵖᵖ returns a Charᵖᵖ that points to addr in the application pool.
func NewCharᵖᵖ(addr uint64) Charᵖᵖ {
	return Charᵖᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Charᵖᵖ points to.
func (p Charᵖᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pointer.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Read reads and returns the Charᵖ element at the pointer.
func (p Charᵖᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖ {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the Charᵖ element at the pointer.
func (p Charᵖᵖ) Write(value Charᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]Charᵖ{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Charᵖᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Charᵖᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p Charᵖᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new Charᵖˢ from the pointer using start and end indices.
func (p Charᵖᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charᵖˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charᵖˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// S8ᵖ is a pointer to a int8 element.
type S8ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewS8ᵖ returns a S8ᵖ that points to addr in the application pool.
func NewS8ᵖ(addr uint64) S8ᵖ {
	return S8ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that S8ᵖ points to.
func (p S8ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the int8 element at the pointer.
func (p S8ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int8 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the int8 element at the pointer.
func (p S8ᵖ) Write(value int8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]int8{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p S8ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p S8ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p S8ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new S8ˢ from the pointer using start and end indices.
func (p S8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S8ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S8ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// S16ᵖ is a pointer to a int16 element.
type S16ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewS16ᵖ returns a S16ᵖ that points to addr in the application pool.
func NewS16ᵖ(addr uint64) S16ᵖ {
	return S16ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that S16ᵖ points to.
func (p S16ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Read reads and returns the int16 element at the pointer.
func (p S16ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int16 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the int16 element at the pointer.
func (p S16ᵖ) Write(value int16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]int16{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p S16ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p S16ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p S16ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new S16ˢ from the pointer using start and end indices.
func (p S16ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S16ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S16ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// F32ᵖ is a pointer to a float32 element.
type F32ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewF32ᵖ returns a F32ᵖ that points to addr in the application pool.
func NewF32ᵖ(addr uint64) F32ᵖ {
	return F32ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that F32ᵖ points to.
func (p F32ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the float32 element at the pointer.
func (p F32ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float32 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the float32 element at the pointer.
func (p F32ᵖ) Write(value float32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]float32{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p F32ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p F32ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p F32ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new F32ˢ from the pointer using start and end indices.
func (p F32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return F32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// S32ᵖ is a pointer to a int32 element.
type S32ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewS32ᵖ returns a S32ᵖ that points to addr in the application pool.
func NewS32ᵖ(addr uint64) S32ᵖ {
	return S32ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that S32ᵖ points to.
func (p S32ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the int32 element at the pointer.
func (p S32ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int32 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the int32 element at the pointer.
func (p S32ᵖ) Write(value int32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]int32{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p S32ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p S32ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p S32ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new S32ˢ from the pointer using start and end indices.
func (p S32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// F64ᵖ is a pointer to a float64 element.
type F64ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewF64ᵖ returns a F64ᵖ that points to addr in the application pool.
func NewF64ᵖ(addr uint64) F64ᵖ {
	return F64ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that F64ᵖ points to.
func (p F64ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Read reads and returns the float64 element at the pointer.
func (p F64ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) float64 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the float64 element at the pointer.
func (p F64ᵖ) Write(value float64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]float64{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p F64ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p F64ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p F64ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new F64ˢ from the pointer using start and end indices.
func (p F64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return F64ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// U64ᵖ is a pointer to a uint64 element.
type U64ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewU64ᵖ returns a U64ᵖ that points to addr in the application pool.
func NewU64ᵖ(addr uint64) U64ᵖ {
	return U64ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that U64ᵖ points to.
func (p U64ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Read reads and returns the uint64 element at the pointer.
func (p U64ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the uint64 element at the pointer.
func (p U64ᵖ) Write(value uint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]uint64{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p U64ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p U64ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p U64ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new U64ˢ from the pointer using start and end indices.
func (p U64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U64ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// S64ᵖ is a pointer to a int64 element.
type S64ᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewS64ᵖ returns a S64ᵖ that points to addr in the application pool.
func NewS64ᵖ(addr uint64) S64ᵖ {
	return S64ᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that S64ᵖ points to.
func (p S64ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Read reads and returns the int64 element at the pointer.
func (p S64ᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) int64 {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the int64 element at the pointer.
func (p S64ᵖ) Write(value int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]int64{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p S64ᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p S64ᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p S64ᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new S64ˢ from the pointer using start and end indices.
func (p S64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S64ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Boolᵖ is a pointer to a bool element.
type Boolᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewBoolᵖ returns a Boolᵖ that points to addr in the application pool.
func NewBoolᵖ(addr uint64) Boolᵖ {
	return Boolᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Boolᵖ points to.
func (p Boolᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the bool element at the pointer.
func (p Boolᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) bool {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the bool element at the pointer.
func (p Boolᵖ) Write(value bool, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]bool{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Boolᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Boolᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p Boolᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new Boolˢ from the pointer using start and end indices.
func (p Boolᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Boolˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Voidᵖ is a pointer to a void element.
type Voidᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVoidᵖ returns a Voidᵖ that points to addr in the application pool.
func NewVoidᵖ(addr uint64) Voidᵖ {
	return Voidᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Voidᵖ points to.
func (p Voidᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Voidᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Voidᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p Voidᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p Voidᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Remappedᵖ is a pointer to a remapped element.
type Remappedᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewRemappedᵖ returns a Remappedᵖ that points to addr in the application pool.
func NewRemappedᵖ(addr uint64) Remappedᵖ {
	return Remappedᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Remappedᵖ points to.
func (p Remappedᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the remapped element at the pointer.
func (p Remappedᵖ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) remapped {
	return p.Slice(0, 1, ϟs).Read(ϟa, ϟs, ϟd, ϟl, ϟb)[0]
}

// Write writes value to the remapped element at the pointer.
func (p Remappedᵖ) Write(value remapped, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) {
	p.Slice(0, 1, ϟs).Write([]remapped{value}, ϟa, ϟs, ϟd, ϟl, ϟb)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Remappedᵖ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedᵖ {
	p.Slice(0, 1, ϟs).OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Remappedᵖ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedᵖ {
	p.Slice(0, 1, ϟs).OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}
func (p Remappedᵖ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedᵖ {
	p.Slice(0, 1, ϟs).ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
	return p
}

// Slice returns a new Remappedˢ from the pointer using start and end indices.
func (p Remappedᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Remappedˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Remappedˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Boolˢ is a slice of bool.
type Boolˢ struct {
	binary.Generate
	SliceInfo
}

// MakeBoolˢ returns a Boolˢ backed by a new memory pool.
func MakeBoolˢ(count uint64, ϟs *gfxapi.State) Boolˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Boolˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Boolˢ in a new memory pool.
func (s Boolˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Boolˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Boolˢ points to.
func (s Boolˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Boolˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Boolˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Boolˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Boolˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsBoolˢ returns s cast to a Boolˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsBoolˢ(s Slice, ϟs *gfxapi.State) Boolˢ {
	out := Boolˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the bool elements in this Boolˢ.
func (s Boolˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []bool {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]bool, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Bool(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Boolˢ) Write(src []bool, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Bool(bool(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Boolˢ) Copy(src Boolˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s Boolˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Boolˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Boolˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Boolˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Boolˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a Boolᵖ to the i'th element in this Boolˢ.
func (s Boolˢ) Index(i uint64, ϟs *gfxapi.State) Boolᵖ {
	return Boolᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Boolˢ using start and end indices.
func (s Boolˢ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Boolˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Boolˢ slice.
func (s Boolˢ) String() string {
	return fmt.Sprintf("bool(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Charˢ is a slice of byte.
type Charˢ struct {
	binary.Generate
	SliceInfo
}

// MakeCharˢFromString returns a Charˢ backed by a new
// memory pool containing a copy of str.
func MakeCharˢFromString(str string, ϟs *gfxapi.State) Charˢ {
	pool := &memory.Pool{}
	pool.Write(0, memory.Blob([]byte(str)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	return Charˢ{SliceInfo: SliceInfo{Count: uint64(len(str)), Root: memory.Pointer{Pool: id}}}
}

// MakeCharˢ returns a Charˢ backed by a new memory pool.
func MakeCharˢ(count uint64, ϟs *gfxapi.State) Charˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Charˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Charˢ in a new memory pool.
func (s Charˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Charˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Charˢ points to.
func (s Charˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Charˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Charˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Charˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Charˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsCharˢ returns s cast to a Charˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsCharˢ(s Slice, ϟs *gfxapi.State) Charˢ {
	out := Charˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the byte elements in this Charˢ.
func (s Charˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []byte {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]byte, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Uint8(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Charˢ) Write(src []byte, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint8(uint8(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Charˢ) Copy(src Charˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s Charˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Charˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Charˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Charˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a Charᵖ to the i'th element in this Charˢ.
func (s Charˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖ {
	return Charᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Charˢ using start and end indices.
func (s Charˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Charˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Charˢ slice.
func (s Charˢ) String() string {
	return fmt.Sprintf("byte(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Charᵖˢ is a slice of Charᵖ.
type Charᵖˢ struct {
	binary.Generate
	SliceInfo
}

// MakeCharᵖˢ returns a Charᵖˢ backed by a new memory pool.
func MakeCharᵖˢ(count uint64, ϟs *gfxapi.State) Charᵖˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Charᵖˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Charᵖˢ in a new memory pool.
func (s Charᵖˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Charᵖˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Charᵖˢ points to.
func (s Charᵖˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	if s.Root.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Charᵖˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Charᵖˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Charᵖˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Charᵖˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsCharᵖˢ returns s cast to a Charᵖˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsCharᵖˢ(s Slice, ϟs *gfxapi.State) Charᵖˢ {
	out := Charᵖˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Charᵖ elements in this Charᵖˢ.
func (s Charᵖˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []Charᵖ {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Charᵖ, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if s.Root.Pool == memory.ApplicationPool {
			ptr, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			res[i] = NewCharᵖ(ptr)
		} else {
			if err := d.Value(&res[i]); err != nil {
				panic(err)
			}
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Charᵖˢ) Write(src []Charᵖ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if s.Root.Pool == memory.ApplicationPool {
			if err := binary.WriteUint(e, ϟs.Architecture.PointerSize*8, src[i].Address); err != nil {
				panic(err)
			}
		} else {
			if err := e.Value(&src[i]); err != nil {
				panic(err)
			}
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Charᵖˢ) Copy(src Charᵖˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s Charᵖˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	if (dst.Root.Pool == memory.ApplicationPool) != (src.Root.Pool == memory.ApplicationPool) {
		dst.Write(src.Read(ϟa, ϟs, ϟd, ϟl, ϟb), ϟa, ϟs, ϟd, ϟl, ϟb) // Element-wise copy so we can convert u64 <-> Charᵖᵖ
	} else {
		src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
		dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	}
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Charᵖˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step, d := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs)), s.Decoder(ϟs, ϟd, ϟl)
		for i := uint64(0); i < s.Count; i++ {
			v, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			ϟb.Push(NewCharᵖ(v).value())
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Charᵖˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Charᵖˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Charᵖˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a Charᵖᵖ to the i'th element in this Charᵖˢ.
func (s Charᵖˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖᵖ {
	return Charᵖᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Charᵖˢ using start and end indices.
func (s Charᵖˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charᵖˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Charᵖˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Charᵖˢ slice.
func (s Charᵖˢ) String() string {
	return fmt.Sprintf("Charᵖ(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// F32ˢ is a slice of float32.
type F32ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeF32ˢ returns a F32ˢ backed by a new memory pool.
func MakeF32ˢ(count uint64, ϟs *gfxapi.State) F32ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return F32ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the F32ˢ in a new memory pool.
func (s F32ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := F32ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that F32ˢ points to.
func (s F32ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s F32ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s F32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s F32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s F32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsF32ˢ returns s cast to a F32ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsF32ˢ(s Slice, ϟs *gfxapi.State) F32ˢ {
	out := F32ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the float32 elements in this F32ˢ.
func (s F32ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []float32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]float32, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Float32(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s F32ˢ) Write(src []float32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Float32(float32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst F32ˢ) Copy(src F32ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s F32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s F32ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s F32ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s F32ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F32ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a F32ᵖ to the i'th element in this F32ˢ.
func (s F32ˢ) Index(i uint64, ϟs *gfxapi.State) F32ᵖ {
	return F32ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the F32ˢ using start and end indices.
func (s F32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return F32ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the F32ˢ slice.
func (s F32ˢ) String() string {
	return fmt.Sprintf("float32(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// F64ˢ is a slice of float64.
type F64ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeF64ˢ returns a F64ˢ backed by a new memory pool.
func MakeF64ˢ(count uint64, ϟs *gfxapi.State) F64ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return F64ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the F64ˢ in a new memory pool.
func (s F64ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := F64ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that F64ˢ points to.
func (s F64ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s F64ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s F64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s F64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s F64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsF64ˢ returns s cast to a F64ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsF64ˢ(s Slice, ϟs *gfxapi.State) F64ˢ {
	out := F64ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the float64 elements in this F64ˢ.
func (s F64ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []float64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]float64, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Float64(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s F64ˢ) Write(src []float64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Float64(float64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst F64ˢ) Copy(src F64ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s F64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s F64ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s F64ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s F64ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) F64ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a F64ᵖ to the i'th element in this F64ˢ.
func (s F64ˢ) Index(i uint64, ϟs *gfxapi.State) F64ᵖ {
	return F64ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the F64ˢ using start and end indices.
func (s F64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return F64ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the F64ˢ slice.
func (s F64ˢ) String() string {
	return fmt.Sprintf("float64(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Intˢ is a slice of int64.
type Intˢ struct {
	binary.Generate
	SliceInfo
}

// MakeIntˢ returns a Intˢ backed by a new memory pool.
func MakeIntˢ(count uint64, ϟs *gfxapi.State) Intˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Intˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Intˢ in a new memory pool.
func (s Intˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Intˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Intˢ points to.
func (s Intˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(ϟs.Architecture.IntegerSize)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Intˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Intˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Intˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Intˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsIntˢ returns s cast to a Intˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsIntˢ(s Slice, ϟs *gfxapi.State) Intˢ {
	out := Intˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the int64 elements in this Intˢ.
func (s Intˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int64, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := binary.ReadInt(d, ϟs.Architecture.IntegerSize*8); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Intˢ) Write(src []int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := binary.WriteInt(e, ϟs.Architecture.IntegerSize*8, int64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Intˢ) Copy(src Intˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s Intˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Intˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Intˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Intˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Intˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a Intᵖ to the i'th element in this Intˢ.
func (s Intˢ) Index(i uint64, ϟs *gfxapi.State) Intᵖ {
	return Intᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Intˢ using start and end indices.
func (s Intˢ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Intˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Intˢ slice.
func (s Intˢ) String() string {
	return fmt.Sprintf("int64(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Remappedˢ is a slice of remapped.
type Remappedˢ struct {
	binary.Generate
	SliceInfo
}

// MakeRemappedˢ returns a Remappedˢ backed by a new memory pool.
func MakeRemappedˢ(count uint64, ϟs *gfxapi.State) Remappedˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Remappedˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Remappedˢ in a new memory pool.
func (s Remappedˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Remappedˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Remappedˢ points to.
func (s Remappedˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Remappedˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Remappedˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Remappedˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Remappedˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsRemappedˢ returns s cast to a Remappedˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsRemappedˢ(s Slice, ϟs *gfxapi.State) Remappedˢ {
	out := Remappedˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the remapped elements in this Remappedˢ.
func (s Remappedˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []remapped {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]remapped, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Uint32(); err == nil {
			res[i] = remapped(ϟv)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Remappedˢ) Write(src []remapped, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Remappedˢ) Copy(src Remappedˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s Remappedˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Remappedˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ptr, step, d := value.RemappedPointer(s.Base), value.RemappedPointer(s.ElementSize(ϟs)), s.Decoder(ϟs, ϟd, ϟl)
		for i := uint64(0); i < s.Count; i++ {
			var v remapped
			if ϟv, err := d.Uint32(); err == nil {
				v = remapped(ϟv)
			} else {
				panic(err)
			}
			if key, remap := v.remap(ϟa, ϟs); remap {
				loadRemap(ϟb, key, protocol.TypeUint32, v.value(ϟb, ϟa, ϟs))
			} else {
				ϟb.Push(v.value(ϟb, ϟa, ϟs))
			}
			ϟb.Store(ptr)
			ptr += step
		}
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Remappedˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
		size := s.ElementSize(ϟs)
		ptr, step, d := value.RemappedPointer(s.Base), value.RemappedPointer(size), s.Decoder(ϟs, ϟd, ϟl)
		for i := uint64(0); i < s.Count; i++ {
			var v remapped
			if ϟv, err := d.Uint32(); err == nil {
				v = remapped(ϟv)
			} else {
				panic(err)
			}
			if key, remap := v.remap(ϟa, ϟs); remap {
				dst, found := ϟb.Remappings[key]
				if !found {
					dst = ϟb.AllocateMemory(size)
					ϟb.Remappings[key] = dst
				}
				ϟb.Load(protocol.TypeUint32, ptr)
				ϟb.Store(dst)
			}
			ptr += step
		}
	}
	return s
}
func (s Remappedˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Remappedˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a Remappedᵖ to the i'th element in this Remappedˢ.
func (s Remappedˢ) Index(i uint64, ϟs *gfxapi.State) Remappedᵖ {
	return Remappedᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Remappedˢ using start and end indices.
func (s Remappedˢ) Slice(start, end uint64, ϟs *gfxapi.State) Remappedˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Remappedˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Remappedˢ slice.
func (s Remappedˢ) String() string {
	return fmt.Sprintf("remapped(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// S16ˢ is a slice of int16.
type S16ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeS16ˢ returns a S16ˢ backed by a new memory pool.
func MakeS16ˢ(count uint64, ϟs *gfxapi.State) S16ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S16ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the S16ˢ in a new memory pool.
func (s S16ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S16ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that S16ˢ points to.
func (s S16ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S16ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S16ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S16ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S16ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsS16ˢ returns s cast to a S16ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsS16ˢ(s Slice, ϟs *gfxapi.State) S16ˢ {
	out := S16ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the int16 elements in this S16ˢ.
func (s S16ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int16 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int16, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Int16(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s S16ˢ) Write(src []int16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int16(int16(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S16ˢ) Copy(src S16ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s S16ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s S16ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s S16ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s S16ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S16ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a S16ᵖ to the i'th element in this S16ˢ.
func (s S16ˢ) Index(i uint64, ϟs *gfxapi.State) S16ᵖ {
	return S16ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the S16ˢ using start and end indices.
func (s S16ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S16ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S16ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the S16ˢ slice.
func (s S16ˢ) String() string {
	return fmt.Sprintf("int16(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// S32ˢ is a slice of int32.
type S32ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeS32ˢ returns a S32ˢ backed by a new memory pool.
func MakeS32ˢ(count uint64, ϟs *gfxapi.State) S32ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S32ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the S32ˢ in a new memory pool.
func (s S32ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S32ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that S32ˢ points to.
func (s S32ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S32ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsS32ˢ returns s cast to a S32ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsS32ˢ(s Slice, ϟs *gfxapi.State) S32ˢ {
	out := S32ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the int32 elements in this S32ˢ.
func (s S32ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int32, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Int32(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s S32ˢ) Write(src []int32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int32(int32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S32ˢ) Copy(src S32ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s S32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s S32ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s S32ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s S32ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S32ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a S32ᵖ to the i'th element in this S32ˢ.
func (s S32ˢ) Index(i uint64, ϟs *gfxapi.State) S32ᵖ {
	return S32ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the S32ˢ using start and end indices.
func (s S32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S32ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the S32ˢ slice.
func (s S32ˢ) String() string {
	return fmt.Sprintf("int32(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// S64ˢ is a slice of int64.
type S64ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeS64ˢ returns a S64ˢ backed by a new memory pool.
func MakeS64ˢ(count uint64, ϟs *gfxapi.State) S64ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S64ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the S64ˢ in a new memory pool.
func (s S64ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S64ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that S64ˢ points to.
func (s S64ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S64ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsS64ˢ returns s cast to a S64ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsS64ˢ(s Slice, ϟs *gfxapi.State) S64ˢ {
	out := S64ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the int64 elements in this S64ˢ.
func (s S64ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int64, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Int64(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s S64ˢ) Write(src []int64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int64(int64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S64ˢ) Copy(src S64ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s S64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s S64ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s S64ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s S64ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S64ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a S64ᵖ to the i'th element in this S64ˢ.
func (s S64ˢ) Index(i uint64, ϟs *gfxapi.State) S64ᵖ {
	return S64ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the S64ˢ using start and end indices.
func (s S64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S64ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the S64ˢ slice.
func (s S64ˢ) String() string {
	return fmt.Sprintf("int64(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// S8ˢ is a slice of int8.
type S8ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeS8ˢ returns a S8ˢ backed by a new memory pool.
func MakeS8ˢ(count uint64, ϟs *gfxapi.State) S8ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S8ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the S8ˢ in a new memory pool.
func (s S8ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S8ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that S8ˢ points to.
func (s S8ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S8ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S8ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S8ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S8ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsS8ˢ returns s cast to a S8ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsS8ˢ(s Slice, ϟs *gfxapi.State) S8ˢ {
	out := S8ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the int8 elements in this S8ˢ.
func (s S8ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []int8 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int8, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Int8(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s S8ˢ) Write(src []int8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int8(int8(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S8ˢ) Copy(src S8ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s S8ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s S8ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s S8ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s S8ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) S8ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a S8ᵖ to the i'th element in this S8ˢ.
func (s S8ˢ) Index(i uint64, ϟs *gfxapi.State) S8ᵖ {
	return S8ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the S8ˢ using start and end indices.
func (s S8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S8ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S8ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the S8ˢ slice.
func (s S8ˢ) String() string {
	return fmt.Sprintf("int8(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// U16ˢ is a slice of uint16.
type U16ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeU16ˢ returns a U16ˢ backed by a new memory pool.
func MakeU16ˢ(count uint64, ϟs *gfxapi.State) U16ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U16ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the U16ˢ in a new memory pool.
func (s U16ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U16ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that U16ˢ points to.
func (s U16ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U16ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U16ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U16ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U16ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsU16ˢ returns s cast to a U16ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsU16ˢ(s Slice, ϟs *gfxapi.State) U16ˢ {
	out := U16ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the uint16 elements in this U16ˢ.
func (s U16ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint16 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint16, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Uint16(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s U16ˢ) Write(src []uint16, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint16(uint16(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U16ˢ) Copy(src U16ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s U16ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s U16ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s U16ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s U16ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U16ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a U16ᵖ to the i'th element in this U16ˢ.
func (s U16ˢ) Index(i uint64, ϟs *gfxapi.State) U16ᵖ {
	return U16ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the U16ˢ using start and end indices.
func (s U16ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U16ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U16ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the U16ˢ slice.
func (s U16ˢ) String() string {
	return fmt.Sprintf("uint16(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// U32ˢ is a slice of uint32.
type U32ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeU32ˢ returns a U32ˢ backed by a new memory pool.
func MakeU32ˢ(count uint64, ϟs *gfxapi.State) U32ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U32ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the U32ˢ in a new memory pool.
func (s U32ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U32ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that U32ˢ points to.
func (s U32ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U32ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsU32ˢ returns s cast to a U32ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsU32ˢ(s Slice, ϟs *gfxapi.State) U32ˢ {
	out := U32ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the uint32 elements in this U32ˢ.
func (s U32ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint32, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Uint32(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s U32ˢ) Write(src []uint32, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U32ˢ) Copy(src U32ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s U32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s U32ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s U32ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s U32ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U32ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a U32ᵖ to the i'th element in this U32ˢ.
func (s U32ˢ) Index(i uint64, ϟs *gfxapi.State) U32ᵖ {
	return U32ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the U32ˢ using start and end indices.
func (s U32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U32ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the U32ˢ slice.
func (s U32ˢ) String() string {
	return fmt.Sprintf("uint32(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// U64ˢ is a slice of uint64.
type U64ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeU64ˢ returns a U64ˢ backed by a new memory pool.
func MakeU64ˢ(count uint64, ϟs *gfxapi.State) U64ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U64ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the U64ˢ in a new memory pool.
func (s U64ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U64ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that U64ˢ points to.
func (s U64ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U64ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsU64ˢ returns s cast to a U64ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsU64ˢ(s Slice, ϟs *gfxapi.State) U64ˢ {
	out := U64ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the uint64 elements in this U64ˢ.
func (s U64ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint64, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Uint64(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s U64ˢ) Write(src []uint64, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint64(uint64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U64ˢ) Copy(src U64ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s U64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s U64ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s U64ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s U64ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U64ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a U64ᵖ to the i'th element in this U64ˢ.
func (s U64ˢ) Index(i uint64, ϟs *gfxapi.State) U64ᵖ {
	return U64ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the U64ˢ using start and end indices.
func (s U64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U64ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the U64ˢ slice.
func (s U64ˢ) String() string {
	return fmt.Sprintf("uint64(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// U8ˢ is a slice of uint8.
type U8ˢ struct {
	binary.Generate
	SliceInfo
}

// MakeU8ˢ returns a U8ˢ backed by a new memory pool.
func MakeU8ˢ(count uint64, ϟs *gfxapi.State) U8ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U8ˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the U8ˢ in a new memory pool.
func (s U8ˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U8ˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that U8ˢ points to.
func (s U8ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U8ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U8ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U8ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U8ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsU8ˢ returns s cast to a U8ˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsU8ˢ(s Slice, ϟs *gfxapi.State) U8ˢ {
	out := U8ˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the uint8 elements in this U8ˢ.
func (s U8ˢ) Read(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) []uint8 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint8, s.Count)
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	for i := range res {
		if ϟv, err := d.Uint8(); err == nil {
			res[i] = ϟv
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s U8ˢ) Write(src []uint8, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint8(uint8(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U8ˢ) Copy(src U8ˢ, ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) (d, s U8ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟa, ϟs, ϟd, ϟl, ϟb)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s U8ˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s U8ˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s U8ˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) U8ˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a U8ᵖ to the i'th element in this U8ˢ.
func (s U8ˢ) Index(i uint64, ϟs *gfxapi.State) U8ᵖ {
	return U8ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the U8ˢ using start and end indices.
func (s U8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U8ˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the U8ˢ slice.
func (s U8ˢ) String() string {
	return fmt.Sprintf("uint8(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Voidˢ is a slice of void.
type Voidˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVoidˢ returns a Voidˢ backed by a new memory pool.
func MakeVoidˢ(count uint64, ϟs *gfxapi.State) Voidˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Voidˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Voidˢ in a new memory pool.
func (s Voidˢ) Clone(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	s.OnRead(ϟa, ϟs, ϟd, ϟl, ϟb)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Voidˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Voidˢ points to.
func (s Voidˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Voidˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Voidˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Voidˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Voidˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Voidˢ) OnRead(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		s.ReserveMemory(ϟa, ϟs, ϟd, ϟl, ϟb)
		ϟb.Write(s.Range(ϟs), s.ResourceID(ϟs, ϟd, ϟl))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Voidˢ) OnWrite(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		ϟb.ReserveMemory(s.Root.Range(uint64(s.Range(ϟs).End() - s.Root.Address)))
	}
	return s
}
func (s Voidˢ) ReserveMemory(ϟa atom.Atom, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) Voidˢ {
	if ϟb != nil && s.Root.Pool == memory.ApplicationPool {
		rng := s.Range(ϟs)
		ϟb.ReserveMemory(s.Root.Range(uint64(rng.End() - s.Root.Address)))
	}
	return s
}

// Index returns a Voidᵖ to the i'th element in this Voidˢ.
func (s Voidˢ) Index(i uint64, ϟs *gfxapi.State) Voidᵖ {
	return Voidᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Voidˢ using start and end indices.
func (s Voidˢ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("%v.Slice start (%d) is greater than the end (%d)", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Voidˢ slice.
func (s Voidˢ) String() string {
	return fmt.Sprintf("void(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

////////////////////////////////////////////////////////////////////////////////
// CmdClone
////////////////////////////////////////////////////////////////////////////////
type CmdClone struct {
	binary.Generate
	observations atom.Observations
	Src          U8ᵖ
	Cnt          uint32
}

func (a *CmdClone) String() string {
	return fmt.Sprintf("cmdClone(src: %v, cnt: %v)", a.Src, a.Cnt)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdClone pointer is returned so that calls can be chained.
func (a *CmdClone) AddRead(rng memory.Range, id binary.ID) *CmdClone {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdClone pointer is returned so that calls can be chained.
func (a *CmdClone) AddWrite(rng memory.Range, id binary.ID) *CmdClone {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdClone) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdClone) Flags() atom.Flags                { return 0 }
func (a *CmdClone) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdMake
////////////////////////////////////////////////////////////////////////////////
type CmdMake struct {
	binary.Generate
	observations atom.Observations
	Cnt          uint32
}

func (a *CmdMake) String() string {
	return fmt.Sprintf("cmdMake(cnt: %v)", a.Cnt)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdMake pointer is returned so that calls can be chained.
func (a *CmdMake) AddRead(rng memory.Range, id binary.ID) *CmdMake {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdMake pointer is returned so that calls can be chained.
func (a *CmdMake) AddWrite(rng memory.Range, id binary.ID) *CmdMake {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdMake) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdMake) Flags() atom.Flags                { return 0 }
func (a *CmdMake) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdCopy
////////////////////////////////////////////////////////////////////////////////
type CmdCopy struct {
	binary.Generate
	observations atom.Observations
	Src          U8ᵖ
	Cnt          uint32
}

func (a *CmdCopy) String() string {
	return fmt.Sprintf("cmdCopy(src: %v, cnt: %v)", a.Src, a.Cnt)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdCopy pointer is returned so that calls can be chained.
func (a *CmdCopy) AddRead(rng memory.Range, id binary.ID) *CmdCopy {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdCopy pointer is returned so that calls can be chained.
func (a *CmdCopy) AddWrite(rng memory.Range, id binary.ID) *CmdCopy {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdCopy) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdCopy) Flags() atom.Flags                { return 0 }
func (a *CmdCopy) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdCharsliceToString
////////////////////////////////////////////////////////////////////////////////
type CmdCharsliceToString struct {
	binary.Generate
	observations atom.Observations
	S            Charᵖ
	Len          uint32
}

func (a *CmdCharsliceToString) String() string {
	return fmt.Sprintf("cmdCharsliceToString(s: %v, len: %v)", a.S, a.Len)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdCharsliceToString pointer is returned so that calls can be chained.
func (a *CmdCharsliceToString) AddRead(rng memory.Range, id binary.ID) *CmdCharsliceToString {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdCharsliceToString pointer is returned so that calls can be chained.
func (a *CmdCharsliceToString) AddWrite(rng memory.Range, id binary.ID) *CmdCharsliceToString {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdCharsliceToString) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdCharsliceToString) Flags() atom.Flags                { return 0 }
func (a *CmdCharsliceToString) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdCharptrToString
////////////////////////////////////////////////////////////////////////////////
type CmdCharptrToString struct {
	binary.Generate
	observations atom.Observations
	S            Charᵖ
}

func (a *CmdCharptrToString) String() string {
	return fmt.Sprintf("cmdCharptrToString(s: %v)", a.S)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdCharptrToString pointer is returned so that calls can be chained.
func (a *CmdCharptrToString) AddRead(rng memory.Range, id binary.ID) *CmdCharptrToString {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdCharptrToString pointer is returned so that calls can be chained.
func (a *CmdCharptrToString) AddWrite(rng memory.Range, id binary.ID) *CmdCharptrToString {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdCharptrToString) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdCharptrToString) Flags() atom.Flags                { return 0 }
func (a *CmdCharptrToString) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdSliceCasts
////////////////////////////////////////////////////////////////////////////////
type CmdSliceCasts struct {
	binary.Generate
	observations atom.Observations
	S            U16ᵖ
	L            uint32
}

func (a *CmdSliceCasts) String() string {
	return fmt.Sprintf("cmdSliceCasts(s: %v, l: %v)", a.S, a.L)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdSliceCasts pointer is returned so that calls can be chained.
func (a *CmdSliceCasts) AddRead(rng memory.Range, id binary.ID) *CmdSliceCasts {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdSliceCasts pointer is returned so that calls can be chained.
func (a *CmdSliceCasts) AddWrite(rng memory.Range, id binary.ID) *CmdSliceCasts {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdSliceCasts) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdSliceCasts) Flags() atom.Flags                { return 0 }
func (a *CmdSliceCasts) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid
////////////////////////////////////////////////////////////////////////////////
type CmdVoid struct {
	binary.Generate
	observations atom.Observations
}

func (a *CmdVoid) String() string {
	return fmt.Sprintf("cmdVoid()")
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoid pointer is returned so that calls can be chained.
func (a *CmdVoid) AddRead(rng memory.Range, id binary.ID) *CmdVoid {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoid pointer is returned so that calls can be chained.
func (a *CmdVoid) AddWrite(rng memory.Range, id binary.ID) *CmdVoid {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoid) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoid) Flags() atom.Flags                { return 0 }
func (a *CmdVoid) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdUnknownRet
////////////////////////////////////////////////////////////////////////////////
type CmdUnknownRet struct {
	binary.Generate
	observations atom.Observations
	Result       int64
}

func (a *CmdUnknownRet) String() string {
	return fmt.Sprintf("cmdUnknownRet() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdUnknownRet pointer is returned so that calls can be chained.
func (a *CmdUnknownRet) AddRead(rng memory.Range, id binary.ID) *CmdUnknownRet {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdUnknownRet pointer is returned so that calls can be chained.
func (a *CmdUnknownRet) AddWrite(rng memory.Range, id binary.ID) *CmdUnknownRet {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdUnknownRet) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdUnknownRet) Flags() atom.Flags                { return 0 }
func (a *CmdUnknownRet) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdUnknownWritePtr
////////////////////////////////////////////////////////////////////////////////
type CmdUnknownWritePtr struct {
	binary.Generate
	observations atom.Observations
	P            Intᵖ
}

func (a *CmdUnknownWritePtr) String() string {
	return fmt.Sprintf("cmdUnknownWritePtr(p: %v)", a.P)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdUnknownWritePtr pointer is returned so that calls can be chained.
func (a *CmdUnknownWritePtr) AddRead(rng memory.Range, id binary.ID) *CmdUnknownWritePtr {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdUnknownWritePtr pointer is returned so that calls can be chained.
func (a *CmdUnknownWritePtr) AddWrite(rng memory.Range, id binary.ID) *CmdUnknownWritePtr {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdUnknownWritePtr) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdUnknownWritePtr) Flags() atom.Flags                { return 0 }
func (a *CmdUnknownWritePtr) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdUnknownWriteSlice
////////////////////////////////////////////////////////////////////////////////
type CmdUnknownWriteSlice struct {
	binary.Generate
	observations atom.Observations
	A            Intᵖ
}

func (a *CmdUnknownWriteSlice) String() string {
	return fmt.Sprintf("cmdUnknownWriteSlice(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdUnknownWriteSlice pointer is returned so that calls can be chained.
func (a *CmdUnknownWriteSlice) AddRead(rng memory.Range, id binary.ID) *CmdUnknownWriteSlice {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdUnknownWriteSlice pointer is returned so that calls can be chained.
func (a *CmdUnknownWriteSlice) AddWrite(rng memory.Range, id binary.ID) *CmdUnknownWriteSlice {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdUnknownWriteSlice) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdUnknownWriteSlice) Flags() atom.Flags                { return 0 }
func (a *CmdUnknownWriteSlice) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU8 struct {
	binary.Generate
	observations atom.Observations
	A            uint8
}

func (a *CmdVoidU8) String() string {
	return fmt.Sprintf("cmdVoidU8(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidU8 pointer is returned so that calls can be chained.
func (a *CmdVoidU8) AddRead(rng memory.Range, id binary.ID) *CmdVoidU8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidU8 pointer is returned so that calls can be chained.
func (a *CmdVoidU8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidU8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidU8) Flags() atom.Flags                { return 0 }
func (a *CmdVoidU8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS8 struct {
	binary.Generate
	observations atom.Observations
	A            int8
}

func (a *CmdVoidS8) String() string {
	return fmt.Sprintf("cmdVoidS8(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidS8 pointer is returned so that calls can be chained.
func (a *CmdVoidS8) AddRead(rng memory.Range, id binary.ID) *CmdVoidS8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidS8 pointer is returned so that calls can be chained.
func (a *CmdVoidS8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidS8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidS8) Flags() atom.Flags                { return 0 }
func (a *CmdVoidS8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU16 struct {
	binary.Generate
	observations atom.Observations
	A            uint16
}

func (a *CmdVoidU16) String() string {
	return fmt.Sprintf("cmdVoidU16(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidU16 pointer is returned so that calls can be chained.
func (a *CmdVoidU16) AddRead(rng memory.Range, id binary.ID) *CmdVoidU16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidU16 pointer is returned so that calls can be chained.
func (a *CmdVoidU16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidU16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidU16) Flags() atom.Flags                { return 0 }
func (a *CmdVoidU16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS16 struct {
	binary.Generate
	observations atom.Observations
	A            int16
}

func (a *CmdVoidS16) String() string {
	return fmt.Sprintf("cmdVoidS16(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidS16 pointer is returned so that calls can be chained.
func (a *CmdVoidS16) AddRead(rng memory.Range, id binary.ID) *CmdVoidS16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidS16 pointer is returned so that calls can be chained.
func (a *CmdVoidS16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidS16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidS16) Flags() atom.Flags                { return 0 }
func (a *CmdVoidS16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidF32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidF32 struct {
	binary.Generate
	observations atom.Observations
	A            float32
}

func (a *CmdVoidF32) String() string {
	return fmt.Sprintf("cmdVoidF32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidF32 pointer is returned so that calls can be chained.
func (a *CmdVoidF32) AddRead(rng memory.Range, id binary.ID) *CmdVoidF32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidF32 pointer is returned so that calls can be chained.
func (a *CmdVoidF32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidF32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidF32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidF32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidF32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU32 struct {
	binary.Generate
	observations atom.Observations
	A            uint32
}

func (a *CmdVoidU32) String() string {
	return fmt.Sprintf("cmdVoidU32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidU32 pointer is returned so that calls can be chained.
func (a *CmdVoidU32) AddRead(rng memory.Range, id binary.ID) *CmdVoidU32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidU32 pointer is returned so that calls can be chained.
func (a *CmdVoidU32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidU32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidU32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidU32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS32 struct {
	binary.Generate
	observations atom.Observations
	A            int32
}

func (a *CmdVoidS32) String() string {
	return fmt.Sprintf("cmdVoidS32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidS32 pointer is returned so that calls can be chained.
func (a *CmdVoidS32) AddRead(rng memory.Range, id binary.ID) *CmdVoidS32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidS32 pointer is returned so that calls can be chained.
func (a *CmdVoidS32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidS32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidS32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidS32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidF64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidF64 struct {
	binary.Generate
	observations atom.Observations
	A            float64
}

func (a *CmdVoidF64) String() string {
	return fmt.Sprintf("cmdVoidF64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidF64 pointer is returned so that calls can be chained.
func (a *CmdVoidF64) AddRead(rng memory.Range, id binary.ID) *CmdVoidF64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidF64 pointer is returned so that calls can be chained.
func (a *CmdVoidF64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidF64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidF64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidF64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidF64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidU64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidU64 struct {
	binary.Generate
	observations atom.Observations
	A            uint64
}

func (a *CmdVoidU64) String() string {
	return fmt.Sprintf("cmdVoidU64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidU64 pointer is returned so that calls can be chained.
func (a *CmdVoidU64) AddRead(rng memory.Range, id binary.ID) *CmdVoidU64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidU64 pointer is returned so that calls can be chained.
func (a *CmdVoidU64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidU64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidU64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidU64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidU64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidS64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidS64 struct {
	binary.Generate
	observations atom.Observations
	A            int64
}

func (a *CmdVoidS64) String() string {
	return fmt.Sprintf("cmdVoidS64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidS64 pointer is returned so that calls can be chained.
func (a *CmdVoidS64) AddRead(rng memory.Range, id binary.ID) *CmdVoidS64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidS64 pointer is returned so that calls can be chained.
func (a *CmdVoidS64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidS64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidS64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidS64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidS64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidBool
////////////////////////////////////////////////////////////////////////////////
type CmdVoidBool struct {
	binary.Generate
	observations atom.Observations
	A            bool
}

func (a *CmdVoidBool) String() string {
	return fmt.Sprintf("cmdVoidBool(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidBool pointer is returned so that calls can be chained.
func (a *CmdVoidBool) AddRead(rng memory.Range, id binary.ID) *CmdVoidBool {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidBool pointer is returned so that calls can be chained.
func (a *CmdVoidBool) AddWrite(rng memory.Range, id binary.ID) *CmdVoidBool {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidBool) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidBool) Flags() atom.Flags                { return 0 }
func (a *CmdVoidBool) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidString
////////////////////////////////////////////////////////////////////////////////
type CmdVoidString struct {
	binary.Generate
	observations atom.Observations
	A            string
}

func (a *CmdVoidString) String() string {
	return fmt.Sprintf("cmdVoidString(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidString pointer is returned so that calls can be chained.
func (a *CmdVoidString) AddRead(rng memory.Range, id binary.ID) *CmdVoidString {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidString pointer is returned so that calls can be chained.
func (a *CmdVoidString) AddWrite(rng memory.Range, id binary.ID) *CmdVoidString {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidString) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidString) Flags() atom.Flags                { return 0 }
func (a *CmdVoidString) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Strings
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Strings struct {
	binary.Generate
	observations atom.Observations
	A            string
	B            string
	C            string
}

func (a *CmdVoid3Strings) String() string {
	return fmt.Sprintf("cmdVoid3Strings(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoid3Strings pointer is returned so that calls can be chained.
func (a *CmdVoid3Strings) AddRead(rng memory.Range, id binary.ID) *CmdVoid3Strings {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoid3Strings pointer is returned so that calls can be chained.
func (a *CmdVoid3Strings) AddWrite(rng memory.Range, id binary.ID) *CmdVoid3Strings {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoid3Strings) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoid3Strings) Flags() atom.Flags                { return 0 }
func (a *CmdVoid3Strings) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3InArrays
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3InArrays struct {
	binary.Generate
	observations atom.Observations
	A            U8ᵖ
	B            U32ᵖ
	C            Intᵖ
}

func (a *CmdVoid3InArrays) String() string {
	return fmt.Sprintf("cmdVoid3InArrays(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoid3InArrays pointer is returned so that calls can be chained.
func (a *CmdVoid3InArrays) AddRead(rng memory.Range, id binary.ID) *CmdVoid3InArrays {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoid3InArrays pointer is returned so that calls can be chained.
func (a *CmdVoid3InArrays) AddWrite(rng memory.Range, id binary.ID) *CmdVoid3InArrays {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoid3InArrays) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoid3InArrays) Flags() atom.Flags                { return 0 }
func (a *CmdVoid3InArrays) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidInArrayOfPointers
////////////////////////////////////////////////////////////////////////////////
type CmdVoidInArrayOfPointers struct {
	binary.Generate
	observations atom.Observations
	A            Charᵖᵖ
	Count        int32
}

func (a *CmdVoidInArrayOfPointers) String() string {
	return fmt.Sprintf("cmdVoidInArrayOfPointers(a: %v, count: %v)", a.A, a.Count)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidInArrayOfPointers pointer is returned so that calls can be chained.
func (a *CmdVoidInArrayOfPointers) AddRead(rng memory.Range, id binary.ID) *CmdVoidInArrayOfPointers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidInArrayOfPointers pointer is returned so that calls can be chained.
func (a *CmdVoidInArrayOfPointers) AddWrite(rng memory.Range, id binary.ID) *CmdVoidInArrayOfPointers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidInArrayOfPointers) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidInArrayOfPointers) Flags() atom.Flags                { return 0 }
func (a *CmdVoidInArrayOfPointers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadU8 struct {
	binary.Generate
	observations atom.Observations
	A            U8ᵖ
}

func (a *CmdVoidReadU8) String() string {
	return fmt.Sprintf("cmdVoidReadU8(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU8 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU8) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU8 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadU8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadU8) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadU8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadS8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadS8 struct {
	binary.Generate
	observations atom.Observations
	A            S8ᵖ
}

func (a *CmdVoidReadS8) String() string {
	return fmt.Sprintf("cmdVoidReadS8(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS8 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS8) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS8 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadS8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadS8) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadS8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadU16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadU16 struct {
	binary.Generate
	observations atom.Observations
	A            U16ᵖ
}

func (a *CmdVoidReadU16) String() string {
	return fmt.Sprintf("cmdVoidReadU16(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU16 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU16) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU16 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadU16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadU16) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadU16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadS16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadS16 struct {
	binary.Generate
	observations atom.Observations
	A            S16ᵖ
}

func (a *CmdVoidReadS16) String() string {
	return fmt.Sprintf("cmdVoidReadS16(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS16 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS16) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS16 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadS16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadS16) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadS16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadF32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadF32 struct {
	binary.Generate
	observations atom.Observations
	A            F32ᵖ
}

func (a *CmdVoidReadF32) String() string {
	return fmt.Sprintf("cmdVoidReadF32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadF32 pointer is returned so that calls can be chained.
func (a *CmdVoidReadF32) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadF32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadF32 pointer is returned so that calls can be chained.
func (a *CmdVoidReadF32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadF32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadF32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadF32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadF32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadU32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadU32 struct {
	binary.Generate
	observations atom.Observations
	A            U32ᵖ
}

func (a *CmdVoidReadU32) String() string {
	return fmt.Sprintf("cmdVoidReadU32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU32 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU32) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU32 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadU32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadU32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadU32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadS32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadS32 struct {
	binary.Generate
	observations atom.Observations
	A            S32ᵖ
}

func (a *CmdVoidReadS32) String() string {
	return fmt.Sprintf("cmdVoidReadS32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS32 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS32) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS32 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadS32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadS32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadS32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadF64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadF64 struct {
	binary.Generate
	observations atom.Observations
	A            F64ᵖ
}

func (a *CmdVoidReadF64) String() string {
	return fmt.Sprintf("cmdVoidReadF64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadF64 pointer is returned so that calls can be chained.
func (a *CmdVoidReadF64) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadF64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadF64 pointer is returned so that calls can be chained.
func (a *CmdVoidReadF64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadF64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadF64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadF64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadF64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadU64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadU64 struct {
	binary.Generate
	observations atom.Observations
	A            U64ᵖ
}

func (a *CmdVoidReadU64) String() string {
	return fmt.Sprintf("cmdVoidReadU64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU64 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU64) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadU64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadU64 pointer is returned so that calls can be chained.
func (a *CmdVoidReadU64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadU64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadU64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadU64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadU64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadS64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadS64 struct {
	binary.Generate
	observations atom.Observations
	A            S64ᵖ
}

func (a *CmdVoidReadS64) String() string {
	return fmt.Sprintf("cmdVoidReadS64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS64 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS64) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadS64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadS64 pointer is returned so that calls can be chained.
func (a *CmdVoidReadS64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadS64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadS64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadS64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadS64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadBool
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadBool struct {
	binary.Generate
	observations atom.Observations
	A            Boolᵖ
}

func (a *CmdVoidReadBool) String() string {
	return fmt.Sprintf("cmdVoidReadBool(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadBool pointer is returned so that calls can be chained.
func (a *CmdVoidReadBool) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadBool {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadBool pointer is returned so that calls can be chained.
func (a *CmdVoidReadBool) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadBool {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadBool) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadBool) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadBool) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadPtrs
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadPtrs struct {
	binary.Generate
	observations atom.Observations
	A            F32ᵖ
	B            U16ᵖ
	C            Boolᵖ
}

func (a *CmdVoidReadPtrs) String() string {
	return fmt.Sprintf("cmdVoidReadPtrs(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidReadPtrs pointer is returned so that calls can be chained.
func (a *CmdVoidReadPtrs) AddRead(rng memory.Range, id binary.ID) *CmdVoidReadPtrs {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidReadPtrs pointer is returned so that calls can be chained.
func (a *CmdVoidReadPtrs) AddWrite(rng memory.Range, id binary.ID) *CmdVoidReadPtrs {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidReadPtrs) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidReadPtrs) Flags() atom.Flags                { return 0 }
func (a *CmdVoidReadPtrs) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteU8 struct {
	binary.Generate
	observations atom.Observations
	A            U8ᵖ
}

func (a *CmdVoidWriteU8) String() string {
	return fmt.Sprintf("cmdVoidWriteU8(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU8 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU8) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU8 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteU8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteU8) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteU8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteS8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteS8 struct {
	binary.Generate
	observations atom.Observations
	A            S8ᵖ
}

func (a *CmdVoidWriteS8) String() string {
	return fmt.Sprintf("cmdVoidWriteS8(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS8 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS8) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS8 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS8) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteS8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteS8) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteS8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteU16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteU16 struct {
	binary.Generate
	observations atom.Observations
	A            U16ᵖ
}

func (a *CmdVoidWriteU16) String() string {
	return fmt.Sprintf("cmdVoidWriteU16(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU16 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU16) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU16 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteU16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteU16) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteU16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteS16
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteS16 struct {
	binary.Generate
	observations atom.Observations
	A            S16ᵖ
}

func (a *CmdVoidWriteS16) String() string {
	return fmt.Sprintf("cmdVoidWriteS16(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS16 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS16) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS16 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS16) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteS16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteS16) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteS16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteF32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteF32 struct {
	binary.Generate
	observations atom.Observations
	A            F32ᵖ
}

func (a *CmdVoidWriteF32) String() string {
	return fmt.Sprintf("cmdVoidWriteF32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteF32 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteF32) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteF32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteF32 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteF32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteF32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteF32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteF32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteF32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteU32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteU32 struct {
	binary.Generate
	observations atom.Observations
	A            U32ᵖ
}

func (a *CmdVoidWriteU32) String() string {
	return fmt.Sprintf("cmdVoidWriteU32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU32 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU32) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU32 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteU32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteU32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteU32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteS32
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteS32 struct {
	binary.Generate
	observations atom.Observations
	A            S32ᵖ
}

func (a *CmdVoidWriteS32) String() string {
	return fmt.Sprintf("cmdVoidWriteS32(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS32 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS32) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS32 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS32) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteS32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteS32) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteS32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteF64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteF64 struct {
	binary.Generate
	observations atom.Observations
	A            F64ᵖ
}

func (a *CmdVoidWriteF64) String() string {
	return fmt.Sprintf("cmdVoidWriteF64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteF64 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteF64) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteF64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteF64 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteF64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteF64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteF64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteF64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteF64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteU64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteU64 struct {
	binary.Generate
	observations atom.Observations
	A            U64ᵖ
}

func (a *CmdVoidWriteU64) String() string {
	return fmt.Sprintf("cmdVoidWriteU64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU64 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU64) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteU64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteU64 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteU64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteU64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteU64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteU64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteU64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteS64
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteS64 struct {
	binary.Generate
	observations atom.Observations
	A            S64ᵖ
}

func (a *CmdVoidWriteS64) String() string {
	return fmt.Sprintf("cmdVoidWriteS64(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS64 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS64) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteS64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteS64 pointer is returned so that calls can be chained.
func (a *CmdVoidWriteS64) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteS64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteS64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteS64) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteS64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWriteBool
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWriteBool struct {
	binary.Generate
	observations atom.Observations
	A            Boolᵖ
}

func (a *CmdVoidWriteBool) String() string {
	return fmt.Sprintf("cmdVoidWriteBool(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteBool pointer is returned so that calls can be chained.
func (a *CmdVoidWriteBool) AddRead(rng memory.Range, id binary.ID) *CmdVoidWriteBool {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWriteBool pointer is returned so that calls can be chained.
func (a *CmdVoidWriteBool) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWriteBool {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWriteBool) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWriteBool) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWriteBool) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidWritePtrs
////////////////////////////////////////////////////////////////////////////////
type CmdVoidWritePtrs struct {
	binary.Generate
	observations atom.Observations
	A            F32ᵖ
	B            U16ᵖ
	C            Boolᵖ
}

func (a *CmdVoidWritePtrs) String() string {
	return fmt.Sprintf("cmdVoidWritePtrs(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidWritePtrs pointer is returned so that calls can be chained.
func (a *CmdVoidWritePtrs) AddRead(rng memory.Range, id binary.ID) *CmdVoidWritePtrs {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidWritePtrs pointer is returned so that calls can be chained.
func (a *CmdVoidWritePtrs) AddWrite(rng memory.Range, id binary.ID) *CmdVoidWritePtrs {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidWritePtrs) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidWritePtrs) Flags() atom.Flags                { return 0 }
func (a *CmdVoidWritePtrs) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdU8
////////////////////////////////////////////////////////////////////////////////
type CmdU8 struct {
	binary.Generate
	observations atom.Observations
	Result       uint8
}

func (a *CmdU8) String() string {
	return fmt.Sprintf("cmdU8() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdU8 pointer is returned so that calls can be chained.
func (a *CmdU8) AddRead(rng memory.Range, id binary.ID) *CmdU8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdU8 pointer is returned so that calls can be chained.
func (a *CmdU8) AddWrite(rng memory.Range, id binary.ID) *CmdU8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdU8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdU8) Flags() atom.Flags                { return 0 }
func (a *CmdU8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdS8
////////////////////////////////////////////////////////////////////////////////
type CmdS8 struct {
	binary.Generate
	observations atom.Observations
	Result       int8
}

func (a *CmdS8) String() string {
	return fmt.Sprintf("cmdS8() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdS8 pointer is returned so that calls can be chained.
func (a *CmdS8) AddRead(rng memory.Range, id binary.ID) *CmdS8 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdS8 pointer is returned so that calls can be chained.
func (a *CmdS8) AddWrite(rng memory.Range, id binary.ID) *CmdS8 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdS8) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdS8) Flags() atom.Flags                { return 0 }
func (a *CmdS8) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdU16
////////////////////////////////////////////////////////////////////////////////
type CmdU16 struct {
	binary.Generate
	observations atom.Observations
	Result       uint16
}

func (a *CmdU16) String() string {
	return fmt.Sprintf("cmdU16() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdU16 pointer is returned so that calls can be chained.
func (a *CmdU16) AddRead(rng memory.Range, id binary.ID) *CmdU16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdU16 pointer is returned so that calls can be chained.
func (a *CmdU16) AddWrite(rng memory.Range, id binary.ID) *CmdU16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdU16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdU16) Flags() atom.Flags                { return 0 }
func (a *CmdU16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdS16
////////////////////////////////////////////////////////////////////////////////
type CmdS16 struct {
	binary.Generate
	observations atom.Observations
	Result       int16
}

func (a *CmdS16) String() string {
	return fmt.Sprintf("cmdS16() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdS16 pointer is returned so that calls can be chained.
func (a *CmdS16) AddRead(rng memory.Range, id binary.ID) *CmdS16 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdS16 pointer is returned so that calls can be chained.
func (a *CmdS16) AddWrite(rng memory.Range, id binary.ID) *CmdS16 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdS16) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdS16) Flags() atom.Flags                { return 0 }
func (a *CmdS16) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdF32
////////////////////////////////////////////////////////////////////////////////
type CmdF32 struct {
	binary.Generate
	observations atom.Observations
	Result       float32
}

func (a *CmdF32) String() string {
	return fmt.Sprintf("cmdF32() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdF32 pointer is returned so that calls can be chained.
func (a *CmdF32) AddRead(rng memory.Range, id binary.ID) *CmdF32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdF32 pointer is returned so that calls can be chained.
func (a *CmdF32) AddWrite(rng memory.Range, id binary.ID) *CmdF32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdF32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdF32) Flags() atom.Flags                { return 0 }
func (a *CmdF32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdU32
////////////////////////////////////////////////////////////////////////////////
type CmdU32 struct {
	binary.Generate
	observations atom.Observations
	Result       uint32
}

func (a *CmdU32) String() string {
	return fmt.Sprintf("cmdU32() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdU32 pointer is returned so that calls can be chained.
func (a *CmdU32) AddRead(rng memory.Range, id binary.ID) *CmdU32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdU32 pointer is returned so that calls can be chained.
func (a *CmdU32) AddWrite(rng memory.Range, id binary.ID) *CmdU32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdU32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdU32) Flags() atom.Flags                { return 0 }
func (a *CmdU32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdS32
////////////////////////////////////////////////////////////////////////////////
type CmdS32 struct {
	binary.Generate
	observations atom.Observations
	Result       int32
}

func (a *CmdS32) String() string {
	return fmt.Sprintf("cmdS32() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdS32 pointer is returned so that calls can be chained.
func (a *CmdS32) AddRead(rng memory.Range, id binary.ID) *CmdS32 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdS32 pointer is returned so that calls can be chained.
func (a *CmdS32) AddWrite(rng memory.Range, id binary.ID) *CmdS32 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdS32) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdS32) Flags() atom.Flags                { return 0 }
func (a *CmdS32) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdF64
////////////////////////////////////////////////////////////////////////////////
type CmdF64 struct {
	binary.Generate
	observations atom.Observations
	Result       float64
}

func (a *CmdF64) String() string {
	return fmt.Sprintf("cmdF64() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdF64 pointer is returned so that calls can be chained.
func (a *CmdF64) AddRead(rng memory.Range, id binary.ID) *CmdF64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdF64 pointer is returned so that calls can be chained.
func (a *CmdF64) AddWrite(rng memory.Range, id binary.ID) *CmdF64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdF64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdF64) Flags() atom.Flags                { return 0 }
func (a *CmdF64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdU64
////////////////////////////////////////////////////////////////////////////////
type CmdU64 struct {
	binary.Generate
	observations atom.Observations
	Result       uint64
}

func (a *CmdU64) String() string {
	return fmt.Sprintf("cmdU64() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdU64 pointer is returned so that calls can be chained.
func (a *CmdU64) AddRead(rng memory.Range, id binary.ID) *CmdU64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdU64 pointer is returned so that calls can be chained.
func (a *CmdU64) AddWrite(rng memory.Range, id binary.ID) *CmdU64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdU64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdU64) Flags() atom.Flags                { return 0 }
func (a *CmdU64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdS64
////////////////////////////////////////////////////////////////////////////////
type CmdS64 struct {
	binary.Generate
	observations atom.Observations
	Result       int64
}

func (a *CmdS64) String() string {
	return fmt.Sprintf("cmdS64() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdS64 pointer is returned so that calls can be chained.
func (a *CmdS64) AddRead(rng memory.Range, id binary.ID) *CmdS64 {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdS64 pointer is returned so that calls can be chained.
func (a *CmdS64) AddWrite(rng memory.Range, id binary.ID) *CmdS64 {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdS64) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdS64) Flags() atom.Flags                { return 0 }
func (a *CmdS64) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdBool
////////////////////////////////////////////////////////////////////////////////
type CmdBool struct {
	binary.Generate
	observations atom.Observations
	Result       bool
}

func (a *CmdBool) String() string {
	return fmt.Sprintf("cmdBool() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdBool pointer is returned so that calls can be chained.
func (a *CmdBool) AddRead(rng memory.Range, id binary.ID) *CmdBool {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdBool pointer is returned so that calls can be chained.
func (a *CmdBool) AddWrite(rng memory.Range, id binary.ID) *CmdBool {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdBool) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdBool) Flags() atom.Flags                { return 0 }
func (a *CmdBool) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdString
////////////////////////////////////////////////////////////////////////////////
type CmdString struct {
	binary.Generate
	observations atom.Observations
	Result       string
}

func (a *CmdString) String() string {
	return fmt.Sprintf("cmdString() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdString pointer is returned so that calls can be chained.
func (a *CmdString) AddRead(rng memory.Range, id binary.ID) *CmdString {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdString pointer is returned so that calls can be chained.
func (a *CmdString) AddWrite(rng memory.Range, id binary.ID) *CmdString {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdString) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdString) Flags() atom.Flags                { return 0 }
func (a *CmdString) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdPointer
////////////////////////////////////////////////////////////////////////////////
type CmdPointer struct {
	binary.Generate
	observations atom.Observations
	Result       Voidᵖ
}

func (a *CmdPointer) String() string {
	return fmt.Sprintf("cmdPointer() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdPointer pointer is returned so that calls can be chained.
func (a *CmdPointer) AddRead(rng memory.Range, id binary.ID) *CmdPointer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdPointer pointer is returned so that calls can be chained.
func (a *CmdPointer) AddWrite(rng memory.Range, id binary.ID) *CmdPointer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdPointer) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdPointer) Flags() atom.Flags                { return 0 }
func (a *CmdPointer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3Remapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3Remapped struct {
	binary.Generate
	observations atom.Observations
	A            remapped
	B            remapped
	C            remapped
}

func (a *CmdVoid3Remapped) String() string {
	return fmt.Sprintf("cmdVoid3Remapped(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoid3Remapped pointer is returned so that calls can be chained.
func (a *CmdVoid3Remapped) AddRead(rng memory.Range, id binary.ID) *CmdVoid3Remapped {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoid3Remapped pointer is returned so that calls can be chained.
func (a *CmdVoid3Remapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoid3Remapped {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoid3Remapped) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoid3Remapped) Flags() atom.Flags                { return 0 }
func (a *CmdVoid3Remapped) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidInArrayOfRemapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoidInArrayOfRemapped struct {
	binary.Generate
	observations atom.Observations
	A            Remappedᵖ
}

func (a *CmdVoidInArrayOfRemapped) String() string {
	return fmt.Sprintf("cmdVoidInArrayOfRemapped(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidInArrayOfRemapped pointer is returned so that calls can be chained.
func (a *CmdVoidInArrayOfRemapped) AddRead(rng memory.Range, id binary.ID) *CmdVoidInArrayOfRemapped {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidInArrayOfRemapped pointer is returned so that calls can be chained.
func (a *CmdVoidInArrayOfRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoidInArrayOfRemapped {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidInArrayOfRemapped) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidInArrayOfRemapped) Flags() atom.Flags                { return 0 }
func (a *CmdVoidInArrayOfRemapped) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutArrayOfRemapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutArrayOfRemapped struct {
	binary.Generate
	observations atom.Observations
	A            Remappedᵖ
}

func (a *CmdVoidOutArrayOfRemapped) String() string {
	return fmt.Sprintf("cmdVoidOutArrayOfRemapped(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidOutArrayOfRemapped pointer is returned so that calls can be chained.
func (a *CmdVoidOutArrayOfRemapped) AddRead(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfRemapped {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidOutArrayOfRemapped pointer is returned so that calls can be chained.
func (a *CmdVoidOutArrayOfRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfRemapped {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidOutArrayOfRemapped) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidOutArrayOfRemapped) Flags() atom.Flags                { return 0 }
func (a *CmdVoidOutArrayOfRemapped) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidOutArrayOfUnknownRemapped
////////////////////////////////////////////////////////////////////////////////
type CmdVoidOutArrayOfUnknownRemapped struct {
	binary.Generate
	observations atom.Observations
	A            Remappedᵖ
}

func (a *CmdVoidOutArrayOfUnknownRemapped) String() string {
	return fmt.Sprintf("cmdVoidOutArrayOfUnknownRemapped(a: %v)", a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdVoidOutArrayOfUnknownRemapped pointer is returned so that calls can be chained.
func (a *CmdVoidOutArrayOfUnknownRemapped) AddRead(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfUnknownRemapped {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdVoidOutArrayOfUnknownRemapped pointer is returned so that calls can be chained.
func (a *CmdVoidOutArrayOfUnknownRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdVoidOutArrayOfUnknownRemapped {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdVoidOutArrayOfUnknownRemapped) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdVoidOutArrayOfUnknownRemapped) Flags() atom.Flags                { return 0 }
func (a *CmdVoidOutArrayOfUnknownRemapped) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdRemapped
////////////////////////////////////////////////////////////////////////////////
type CmdRemapped struct {
	binary.Generate
	observations atom.Observations
	Result       remapped
}

func (a *CmdRemapped) String() string {
	return fmt.Sprintf("cmdRemapped() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CmdRemapped pointer is returned so that calls can be chained.
func (a *CmdRemapped) AddRead(rng memory.Range, id binary.ID) *CmdRemapped {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CmdRemapped pointer is returned so that calls can be chained.
func (a *CmdRemapped) AddWrite(rng memory.Range, id binary.ID) *CmdRemapped {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CmdRemapped) API() gfxapi.ID                   { return api{}.ID() }
func (c *CmdRemapped) Flags() atom.Flags                { return 0 }
func (a *CmdRemapped) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// class Tester
////////////////////////////////////////////////////////////////////////////////
type Tester struct {
	binary.Generate
	A Imported
	B Included
}

////////////////////////////////////////////////////////////////////////////////
// class Included
////////////////////////////////////////////////////////////////////////////////
type Included struct {
	binary.Generate
	S string
}

////////////////////////////////////////////////////////////////////////////////
// State
////////////////////////////////////////////////////////////////////////////////
type State struct {
	binary.Generate
	U8s  U8ˢ
	U16s U16ˢ
	U32s U32ˢ
	Ints Intˢ
	Str  string
}

func (g *State) Init() {
}
func NewCmdClone(Src memory.Pointer, Cnt uint32) *CmdClone {
	return &CmdClone{Src: U8ᵖ{Pointer: Src}, Cnt: Cnt}
}
func NewCmdMake(Cnt uint32) *CmdMake {
	return &CmdMake{Cnt: Cnt}
}
func NewCmdCopy(Src memory.Pointer, Cnt uint32) *CmdCopy {
	return &CmdCopy{Src: U8ᵖ{Pointer: Src}, Cnt: Cnt}
}
func NewCmdCharsliceToString(S memory.Pointer, Len uint32) *CmdCharsliceToString {
	return &CmdCharsliceToString{S: Charᵖ{Pointer: S}, Len: Len}
}
func NewCmdCharptrToString(S memory.Pointer) *CmdCharptrToString {
	return &CmdCharptrToString{S: Charᵖ{Pointer: S}}
}
func NewCmdSliceCasts(S memory.Pointer, L uint32) *CmdSliceCasts {
	return &CmdSliceCasts{S: U16ᵖ{Pointer: S}, L: L}
}
func NewCmdVoid() *CmdVoid {
	return &CmdVoid{}
}
func NewCmdUnknownRet(Result int64) *CmdUnknownRet {
	return &CmdUnknownRet{Result: Result}
}
func NewCmdUnknownWritePtr(P memory.Pointer) *CmdUnknownWritePtr {
	return &CmdUnknownWritePtr{P: Intᵖ{Pointer: P}}
}
func NewCmdUnknownWriteSlice(A memory.Pointer) *CmdUnknownWriteSlice {
	return &CmdUnknownWriteSlice{A: Intᵖ{Pointer: A}}
}
func NewCmdVoidU8(A uint8) *CmdVoidU8 {
	return &CmdVoidU8{A: A}
}
func NewCmdVoidS8(A int8) *CmdVoidS8 {
	return &CmdVoidS8{A: A}
}
func NewCmdVoidU16(A uint16) *CmdVoidU16 {
	return &CmdVoidU16{A: A}
}
func NewCmdVoidS16(A int16) *CmdVoidS16 {
	return &CmdVoidS16{A: A}
}
func NewCmdVoidF32(A float32) *CmdVoidF32 {
	return &CmdVoidF32{A: A}
}
func NewCmdVoidU32(A uint32) *CmdVoidU32 {
	return &CmdVoidU32{A: A}
}
func NewCmdVoidS32(A int32) *CmdVoidS32 {
	return &CmdVoidS32{A: A}
}
func NewCmdVoidF64(A float64) *CmdVoidF64 {
	return &CmdVoidF64{A: A}
}
func NewCmdVoidU64(A uint64) *CmdVoidU64 {
	return &CmdVoidU64{A: A}
}
func NewCmdVoidS64(A int64) *CmdVoidS64 {
	return &CmdVoidS64{A: A}
}
func NewCmdVoidBool(A bool) *CmdVoidBool {
	return &CmdVoidBool{A: A}
}
func NewCmdVoidString(A string) *CmdVoidString {
	return &CmdVoidString{A: A}
}
func NewCmdVoid3Strings(A string, B string, C string) *CmdVoid3Strings {
	return &CmdVoid3Strings{A: A, B: B, C: C}
}
func NewCmdVoid3InArrays(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoid3InArrays {
	return &CmdVoid3InArrays{A: U8ᵖ{Pointer: A}, B: U32ᵖ{Pointer: B}, C: Intᵖ{Pointer: C}}
}
func NewCmdVoidInArrayOfPointers(A memory.Pointer, Count int32) *CmdVoidInArrayOfPointers {
	return &CmdVoidInArrayOfPointers{A: Charᵖᵖ{Pointer: A}, Count: Count}
}
func NewCmdVoidReadU8(A memory.Pointer) *CmdVoidReadU8 {
	return &CmdVoidReadU8{A: U8ᵖ{Pointer: A}}
}
func NewCmdVoidReadS8(A memory.Pointer) *CmdVoidReadS8 {
	return &CmdVoidReadS8{A: S8ᵖ{Pointer: A}}
}
func NewCmdVoidReadU16(A memory.Pointer) *CmdVoidReadU16 {
	return &CmdVoidReadU16{A: U16ᵖ{Pointer: A}}
}
func NewCmdVoidReadS16(A memory.Pointer) *CmdVoidReadS16 {
	return &CmdVoidReadS16{A: S16ᵖ{Pointer: A}}
}
func NewCmdVoidReadF32(A memory.Pointer) *CmdVoidReadF32 {
	return &CmdVoidReadF32{A: F32ᵖ{Pointer: A}}
}
func NewCmdVoidReadU32(A memory.Pointer) *CmdVoidReadU32 {
	return &CmdVoidReadU32{A: U32ᵖ{Pointer: A}}
}
func NewCmdVoidReadS32(A memory.Pointer) *CmdVoidReadS32 {
	return &CmdVoidReadS32{A: S32ᵖ{Pointer: A}}
}
func NewCmdVoidReadF64(A memory.Pointer) *CmdVoidReadF64 {
	return &CmdVoidReadF64{A: F64ᵖ{Pointer: A}}
}
func NewCmdVoidReadU64(A memory.Pointer) *CmdVoidReadU64 {
	return &CmdVoidReadU64{A: U64ᵖ{Pointer: A}}
}
func NewCmdVoidReadS64(A memory.Pointer) *CmdVoidReadS64 {
	return &CmdVoidReadS64{A: S64ᵖ{Pointer: A}}
}
func NewCmdVoidReadBool(A memory.Pointer) *CmdVoidReadBool {
	return &CmdVoidReadBool{A: Boolᵖ{Pointer: A}}
}
func NewCmdVoidReadPtrs(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoidReadPtrs {
	return &CmdVoidReadPtrs{A: F32ᵖ{Pointer: A}, B: U16ᵖ{Pointer: B}, C: Boolᵖ{Pointer: C}}
}
func NewCmdVoidWriteU8(A memory.Pointer) *CmdVoidWriteU8 {
	return &CmdVoidWriteU8{A: U8ᵖ{Pointer: A}}
}
func NewCmdVoidWriteS8(A memory.Pointer) *CmdVoidWriteS8 {
	return &CmdVoidWriteS8{A: S8ᵖ{Pointer: A}}
}
func NewCmdVoidWriteU16(A memory.Pointer) *CmdVoidWriteU16 {
	return &CmdVoidWriteU16{A: U16ᵖ{Pointer: A}}
}
func NewCmdVoidWriteS16(A memory.Pointer) *CmdVoidWriteS16 {
	return &CmdVoidWriteS16{A: S16ᵖ{Pointer: A}}
}
func NewCmdVoidWriteF32(A memory.Pointer) *CmdVoidWriteF32 {
	return &CmdVoidWriteF32{A: F32ᵖ{Pointer: A}}
}
func NewCmdVoidWriteU32(A memory.Pointer) *CmdVoidWriteU32 {
	return &CmdVoidWriteU32{A: U32ᵖ{Pointer: A}}
}
func NewCmdVoidWriteS32(A memory.Pointer) *CmdVoidWriteS32 {
	return &CmdVoidWriteS32{A: S32ᵖ{Pointer: A}}
}
func NewCmdVoidWriteF64(A memory.Pointer) *CmdVoidWriteF64 {
	return &CmdVoidWriteF64{A: F64ᵖ{Pointer: A}}
}
func NewCmdVoidWriteU64(A memory.Pointer) *CmdVoidWriteU64 {
	return &CmdVoidWriteU64{A: U64ᵖ{Pointer: A}}
}
func NewCmdVoidWriteS64(A memory.Pointer) *CmdVoidWriteS64 {
	return &CmdVoidWriteS64{A: S64ᵖ{Pointer: A}}
}
func NewCmdVoidWriteBool(A memory.Pointer) *CmdVoidWriteBool {
	return &CmdVoidWriteBool{A: Boolᵖ{Pointer: A}}
}
func NewCmdVoidWritePtrs(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoidWritePtrs {
	return &CmdVoidWritePtrs{A: F32ᵖ{Pointer: A}, B: U16ᵖ{Pointer: B}, C: Boolᵖ{Pointer: C}}
}
func NewCmdU8(Result uint8) *CmdU8 {
	return &CmdU8{Result: Result}
}
func NewCmdS8(Result int8) *CmdS8 {
	return &CmdS8{Result: Result}
}
func NewCmdU16(Result uint16) *CmdU16 {
	return &CmdU16{Result: Result}
}
func NewCmdS16(Result int16) *CmdS16 {
	return &CmdS16{Result: Result}
}
func NewCmdF32(Result float32) *CmdF32 {
	return &CmdF32{Result: Result}
}
func NewCmdU32(Result uint32) *CmdU32 {
	return &CmdU32{Result: Result}
}
func NewCmdS32(Result int32) *CmdS32 {
	return &CmdS32{Result: Result}
}
func NewCmdF64(Result float64) *CmdF64 {
	return &CmdF64{Result: Result}
}
func NewCmdU64(Result uint64) *CmdU64 {
	return &CmdU64{Result: Result}
}
func NewCmdS64(Result int64) *CmdS64 {
	return &CmdS64{Result: Result}
}
func NewCmdBool(Result bool) *CmdBool {
	return &CmdBool{Result: Result}
}
func NewCmdString(Result string) *CmdString {
	return &CmdString{Result: Result}
}
func NewCmdPointer(Result memory.Pointer) *CmdPointer {
	return &CmdPointer{Result: Voidᵖ{Pointer: Result}}
}
func NewCmdVoid3Remapped(A remapped, B remapped, C remapped) *CmdVoid3Remapped {
	return &CmdVoid3Remapped{A: A, B: B, C: C}
}
func NewCmdVoidInArrayOfRemapped(A memory.Pointer) *CmdVoidInArrayOfRemapped {
	return &CmdVoidInArrayOfRemapped{A: Remappedᵖ{Pointer: A}}
}
func NewCmdVoidOutArrayOfRemapped(A memory.Pointer) *CmdVoidOutArrayOfRemapped {
	return &CmdVoidOutArrayOfRemapped{A: Remappedᵖ{Pointer: A}}
}
func NewCmdVoidOutArrayOfUnknownRemapped(A memory.Pointer) *CmdVoidOutArrayOfUnknownRemapped {
	return &CmdVoidOutArrayOfUnknownRemapped{A: Remappedᵖ{Pointer: A}}
}
func NewCmdRemapped(Result remapped) *CmdRemapped {
	return &CmdRemapped{Result: Result}
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////
var apiID = gfxapi.ID(binary.NewID([]byte("gfxapi_test")))

type api struct{}

func (api) Name() string {
	return "gfxapi_test"
}
func (api) ID() gfxapi.ID {
	return apiID
}
func (api) GetFramebufferAttachmentSize(state *gfxapi.State, attachment gfxapi.FramebufferAttachment) (width uint32, height uint32, err error) {
	return getState(state).getFramebufferAttachmentSize(attachment)
}
func API() gfxapi.API {
	return api{}
}
func init() {
	gfxapi.Register(API())
}
func min(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}

// SliceInfo is the common data between all slice types.
type SliceInfo struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  uint64         // Address of first element.
	Count uint64         // Number of elements in the slice.
}

// Info returns the SliceInfo. It is used to conform to the Slice interface.
func (s SliceInfo) Info() SliceInfo { return s }

// Slice is the interface implemented by all slice types
type Slice interface {
	// Info returns the SliceInfo of this slice.
	Info() SliceInfo
	// ElementSize returns the size in bytes of a single element in the slice.
	ElementSize(ϟs *gfxapi.State) uint64
	// Range returns the memory range this slice represents in the underlying pool.
	Range(ϟs *gfxapi.State) memory.Range
}
