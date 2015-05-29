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
)

type remapped uint32

// U8ᵖ is a pointer to a uint8 element.
type U8ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewU8ᵖ returns a U8ᵖ that points to addr in the application pool.
func NewU8ᵖ(addr memory.Pointer) U8ᵖ {
	return U8ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p U8ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that U8ᵖ points to.
func (p U8ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the uint8 element at the pointer.
func (p U8ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint8 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the uint8 element at the pointer.
func (p U8ᵖ) Write(value uint8, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]uint8{value}, ϟs)
}

// Slice returns a new U8ˢ from the pointer using start and end indices.
func (p U8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U8ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the U8ᵖ pointer.
func (p U8ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// Charᵖ is a pointer to a byte element.
type Charᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewCharᵖ returns a Charᵖ that points to addr in the application pool.
func NewCharᵖ(addr memory.Pointer) Charᵖ {
	return Charᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p Charᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that Charᵖ points to.
func (p Charᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the byte element at the pointer.
func (p Charᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) byte {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the byte element at the pointer.
func (p Charᵖ) Write(value byte, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]byte{value}, ϟs)
}

// Slice returns a new Charˢ from the pointer using start and end indices.
func (p Charᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the Charᵖ pointer.
func (p Charᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// Intᵖ is a pointer to a int64 element.
type Intᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewIntᵖ returns a Intᵖ that points to addr in the application pool.
func NewIntᵖ(addr memory.Pointer) Intᵖ {
	return Intᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p Intᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that Intᵖ points to.
func (p Intᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(ϟs.Architecture.IntegerSize)
}

// Read reads and returns the int64 element at the pointer.
func (p Intᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int64 element at the pointer.
func (p Intᵖ) Write(value int64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int64{value}, ϟs)
}

// Slice returns a new Intˢ from the pointer using start and end indices.
func (p Intᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Intˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the Intᵖ pointer.
func (p Intᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// U8ᵖᵖ is a pointer to a U8ᵖ element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type U8ᵖᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewU8ᵖᵖ returns a U8ᵖᵖ that points to addr in the application pool.
func NewU8ᵖᵖ(addr memory.Pointer) U8ᵖᵖ {
	return U8ᵖᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p U8ᵖᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that U8ᵖᵖ points to.
func (p U8ᵖᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Read reads and returns the U8ᵖ element at the pointer.
func (p U8ᵖᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) U8ᵖ {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the U8ᵖ element at the pointer.
func (p U8ᵖᵖ) Write(value U8ᵖ, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]U8ᵖ{value}, ϟs)
}

// Slice returns a new U8ᵖˢ from the pointer using start and end indices.
func (p U8ᵖᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U8ᵖˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U8ᵖˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the U8ᵖᵖ pointer.
func (p U8ᵖᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// S8ᵖ is a pointer to a int8 element.
type S8ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewS8ᵖ returns a S8ᵖ that points to addr in the application pool.
func NewS8ᵖ(addr memory.Pointer) S8ᵖ {
	return S8ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p S8ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that S8ᵖ points to.
func (p S8ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the int8 element at the pointer.
func (p S8ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int8 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int8 element at the pointer.
func (p S8ᵖ) Write(value int8, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int8{value}, ϟs)
}

// Slice returns a new S8ˢ from the pointer using start and end indices.
func (p S8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S8ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S8ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the S8ᵖ pointer.
func (p S8ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// U16ᵖ is a pointer to a uint16 element.
type U16ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewU16ᵖ returns a U16ᵖ that points to addr in the application pool.
func NewU16ᵖ(addr memory.Pointer) U16ᵖ {
	return U16ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p U16ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that U16ᵖ points to.
func (p U16ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Read reads and returns the uint16 element at the pointer.
func (p U16ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint16 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the uint16 element at the pointer.
func (p U16ᵖ) Write(value uint16, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]uint16{value}, ϟs)
}

// Slice returns a new U16ˢ from the pointer using start and end indices.
func (p U16ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U16ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U16ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the U16ᵖ pointer.
func (p U16ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// S16ᵖ is a pointer to a int16 element.
type S16ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewS16ᵖ returns a S16ᵖ that points to addr in the application pool.
func NewS16ᵖ(addr memory.Pointer) S16ᵖ {
	return S16ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p S16ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that S16ᵖ points to.
func (p S16ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Read reads and returns the int16 element at the pointer.
func (p S16ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int16 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int16 element at the pointer.
func (p S16ᵖ) Write(value int16, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int16{value}, ϟs)
}

// Slice returns a new S16ˢ from the pointer using start and end indices.
func (p S16ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S16ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S16ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the S16ᵖ pointer.
func (p S16ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// F32ᵖ is a pointer to a float32 element.
type F32ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewF32ᵖ returns a F32ᵖ that points to addr in the application pool.
func NewF32ᵖ(addr memory.Pointer) F32ᵖ {
	return F32ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p F32ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that F32ᵖ points to.
func (p F32ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the float32 element at the pointer.
func (p F32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the float32 element at the pointer.
func (p F32ᵖ) Write(value float32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]float32{value}, ϟs)
}

// Slice returns a new F32ˢ from the pointer using start and end indices.
func (p F32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return F32ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the F32ᵖ pointer.
func (p F32ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// U32ᵖ is a pointer to a uint32 element.
type U32ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewU32ᵖ returns a U32ᵖ that points to addr in the application pool.
func NewU32ᵖ(addr memory.Pointer) U32ᵖ {
	return U32ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p U32ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that U32ᵖ points to.
func (p U32ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the uint32 element at the pointer.
func (p U32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the uint32 element at the pointer.
func (p U32ᵖ) Write(value uint32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]uint32{value}, ϟs)
}

// Slice returns a new U32ˢ from the pointer using start and end indices.
func (p U32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U32ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the U32ᵖ pointer.
func (p U32ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// S32ᵖ is a pointer to a int32 element.
type S32ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewS32ᵖ returns a S32ᵖ that points to addr in the application pool.
func NewS32ᵖ(addr memory.Pointer) S32ᵖ {
	return S32ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p S32ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that S32ᵖ points to.
func (p S32ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the int32 element at the pointer.
func (p S32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int32 element at the pointer.
func (p S32ᵖ) Write(value int32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int32{value}, ϟs)
}

// Slice returns a new S32ˢ from the pointer using start and end indices.
func (p S32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S32ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the S32ᵖ pointer.
func (p S32ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// F64ᵖ is a pointer to a float64 element.
type F64ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewF64ᵖ returns a F64ᵖ that points to addr in the application pool.
func NewF64ᵖ(addr memory.Pointer) F64ᵖ {
	return F64ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p F64ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that F64ᵖ points to.
func (p F64ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Read reads and returns the float64 element at the pointer.
func (p F64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the float64 element at the pointer.
func (p F64ᵖ) Write(value float64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]float64{value}, ϟs)
}

// Slice returns a new F64ˢ from the pointer using start and end indices.
func (p F64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return F64ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the F64ᵖ pointer.
func (p F64ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// U64ᵖ is a pointer to a uint64 element.
type U64ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewU64ᵖ returns a U64ᵖ that points to addr in the application pool.
func NewU64ᵖ(addr memory.Pointer) U64ᵖ {
	return U64ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p U64ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that U64ᵖ points to.
func (p U64ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Read reads and returns the uint64 element at the pointer.
func (p U64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the uint64 element at the pointer.
func (p U64ᵖ) Write(value uint64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]uint64{value}, ϟs)
}

// Slice returns a new U64ˢ from the pointer using start and end indices.
func (p U64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U64ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the U64ᵖ pointer.
func (p U64ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// S64ᵖ is a pointer to a int64 element.
type S64ᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewS64ᵖ returns a S64ᵖ that points to addr in the application pool.
func NewS64ᵖ(addr memory.Pointer) S64ᵖ {
	return S64ᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p S64ᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that S64ᵖ points to.
func (p S64ᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Read reads and returns the int64 element at the pointer.
func (p S64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int64 element at the pointer.
func (p S64ᵖ) Write(value int64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int64{value}, ϟs)
}

// Slice returns a new S64ˢ from the pointer using start and end indices.
func (p S64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S64ˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the S64ᵖ pointer.
func (p S64ᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// Boolᵖ is a pointer to a bool element.
type Boolᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewBoolᵖ returns a Boolᵖ that points to addr in the application pool.
func NewBoolᵖ(addr memory.Pointer) Boolᵖ {
	return Boolᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p Boolᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that Boolᵖ points to.
func (p Boolᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the bool element at the pointer.
func (p Boolᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) bool {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the bool element at the pointer.
func (p Boolᵖ) Write(value bool, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]bool{value}, ϟs)
}

// Slice returns a new Boolˢ from the pointer using start and end indices.
func (p Boolᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Boolˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the Boolᵖ pointer.
func (p Boolᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// Voidᵖ is a pointer to a void element.
type Voidᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewVoidᵖ returns a Voidᵖ that points to addr in the application pool.
func NewVoidᵖ(addr memory.Pointer) Voidᵖ {
	return Voidᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p Voidᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that Voidᵖ points to.
func (p Voidᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p Voidᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the Voidᵖ pointer.
func (p Voidᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// Remappedᵖ is a pointer to a remapped element.
type Remappedᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewRemappedᵖ returns a Remappedᵖ that points to addr in the application pool.
func NewRemappedᵖ(addr memory.Pointer) Remappedᵖ {
	return Remappedᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p Remappedᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that Remappedᵖ points to.
func (p Remappedᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the remapped element at the pointer.
func (p Remappedᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) remapped {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the remapped element at the pointer.
func (p Remappedᵖ) Write(value remapped, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]remapped{value}, ϟs)
}

// Slice returns a new Remappedˢ from the pointer using start and end indices.
func (p Remappedᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Remappedˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Remappedˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the Remappedᵖ pointer.
func (p Remappedᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

type Boolˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeBoolˢ returns a Boolˢ backed by a new memory pool.
func MakeBoolˢ(count uint64, ϟs *gfxapi.State) Boolˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Boolˢ{Count: count, Pool: id}
}

// Clone returns a copy of the Boolˢ in a new memory pool.
func (s Boolˢ) Clone(ϟs *gfxapi.State) Boolˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Boolˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that Boolˢ points to.
func (s Boolˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Boolˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Boolˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Boolˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Boolˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the bool elements in this Boolˢ.
func (s Boolˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []bool {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]bool, s.Count)
	for i := range res {
		if v, err := d.Bool(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst Boolˢ) Write(src []bool, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Bool(bool(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Boolˢ) Copy(src Boolˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Boolˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a Boolᵖ to the i'th element in this Boolˢ.
func (s Boolˢ) Index(i uint64, ϟs *gfxapi.State) Boolᵖ {
	return Boolᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the Boolˢ using start and end indices.
func (s Boolˢ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Boolˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the Boolˢ slice.
func (s Boolˢ) String() string {
	return fmt.Sprintf("bool(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type Charˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeCharˢFromString returns a Charˢ backed by a new
// memory pool containing a copy of str.
func MakeCharˢFromString(str string, ϟs *gfxapi.State) Charˢ {
	pool := &memory.Pool{}
	pool.Write(0, memory.Blob([]byte(str)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	return Charˢ{Count: uint64(len(str)), Pool: id}
}

// MakeCharˢ returns a Charˢ backed by a new memory pool.
func MakeCharˢ(count uint64, ϟs *gfxapi.State) Charˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Charˢ{Count: count, Pool: id}
}

// Clone returns a copy of the Charˢ in a new memory pool.
func (s Charˢ) Clone(ϟs *gfxapi.State) Charˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Charˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that Charˢ points to.
func (s Charˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Charˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Charˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Charˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Charˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the byte elements in this Charˢ.
func (s Charˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []byte {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]byte, s.Count)
	for i := range res {
		if v, err := d.Uint8(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst Charˢ) Write(src []byte, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint8(uint8(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Charˢ) Copy(src Charˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a Charᵖ to the i'th element in this Charˢ.
func (s Charˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖ {
	return Charᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the Charˢ using start and end indices.
func (s Charˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Charˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the Charˢ slice.
func (s Charˢ) String() string {
	return fmt.Sprintf("byte(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type F32ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeF32ˢ returns a F32ˢ backed by a new memory pool.
func MakeF32ˢ(count uint64, ϟs *gfxapi.State) F32ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return F32ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the F32ˢ in a new memory pool.
func (s F32ˢ) Clone(ϟs *gfxapi.State) F32ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := F32ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that F32ˢ points to.
func (s F32ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s F32ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s F32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s F32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s F32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the float32 elements in this F32ˢ.
func (s F32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]float32, s.Count)
	for i := range res {
		if v, err := d.Float32(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst F32ˢ) Write(src []float32, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Float32(float32(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst F32ˢ) Copy(src F32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a F32ᵖ to the i'th element in this F32ˢ.
func (s F32ˢ) Index(i uint64, ϟs *gfxapi.State) F32ᵖ {
	return F32ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the F32ˢ using start and end indices.
func (s F32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return F32ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the F32ˢ slice.
func (s F32ˢ) String() string {
	return fmt.Sprintf("float32(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type F64ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeF64ˢ returns a F64ˢ backed by a new memory pool.
func MakeF64ˢ(count uint64, ϟs *gfxapi.State) F64ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return F64ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the F64ˢ in a new memory pool.
func (s F64ˢ) Clone(ϟs *gfxapi.State) F64ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := F64ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that F64ˢ points to.
func (s F64ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s F64ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s F64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s F64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s F64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the float64 elements in this F64ˢ.
func (s F64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]float64, s.Count)
	for i := range res {
		if v, err := d.Float64(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst F64ˢ) Write(src []float64, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Float64(float64(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst F64ˢ) Copy(src F64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a F64ᵖ to the i'th element in this F64ˢ.
func (s F64ˢ) Index(i uint64, ϟs *gfxapi.State) F64ᵖ {
	return F64ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the F64ˢ using start and end indices.
func (s F64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return F64ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the F64ˢ slice.
func (s F64ˢ) String() string {
	return fmt.Sprintf("float64(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type Intˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeIntˢ returns a Intˢ backed by a new memory pool.
func MakeIntˢ(count uint64, ϟs *gfxapi.State) Intˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Intˢ{Count: count, Pool: id}
}

// Clone returns a copy of the Intˢ in a new memory pool.
func (s Intˢ) Clone(ϟs *gfxapi.State) Intˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Intˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that Intˢ points to.
func (s Intˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(ϟs.Architecture.IntegerSize)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Intˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Intˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Intˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Intˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the int64 elements in this Intˢ.
func (s Intˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int64, s.Count)
	for i := range res {
		if v, err := binary.ReadInt(d, ϟs.Architecture.IntegerSize*8); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst Intˢ) Write(src []int64, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := binary.WriteInt(e, ϟs.Architecture.IntegerSize*8, int64(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Intˢ) Copy(src Intˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Intˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a Intᵖ to the i'th element in this Intˢ.
func (s Intˢ) Index(i uint64, ϟs *gfxapi.State) Intᵖ {
	return Intᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the Intˢ using start and end indices.
func (s Intˢ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Intˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the Intˢ slice.
func (s Intˢ) String() string {
	return fmt.Sprintf("int64(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type Remappedˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeRemappedˢ returns a Remappedˢ backed by a new memory pool.
func MakeRemappedˢ(count uint64, ϟs *gfxapi.State) Remappedˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Remappedˢ{Count: count, Pool: id}
}

// Clone returns a copy of the Remappedˢ in a new memory pool.
func (s Remappedˢ) Clone(ϟs *gfxapi.State) Remappedˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Remappedˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that Remappedˢ points to.
func (s Remappedˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Remappedˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Remappedˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Remappedˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Remappedˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the remapped elements in this Remappedˢ.
func (s Remappedˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []remapped {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]remapped, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = remapped(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst Remappedˢ) Write(src []remapped, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Remappedˢ) Copy(src Remappedˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Remappedˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a Remappedᵖ to the i'th element in this Remappedˢ.
func (s Remappedˢ) Index(i uint64, ϟs *gfxapi.State) Remappedᵖ {
	return Remappedᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the Remappedˢ using start and end indices.
func (s Remappedˢ) Slice(start, end uint64, ϟs *gfxapi.State) Remappedˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Remappedˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the Remappedˢ slice.
func (s Remappedˢ) String() string {
	return fmt.Sprintf("remapped(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type S16ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeS16ˢ returns a S16ˢ backed by a new memory pool.
func MakeS16ˢ(count uint64, ϟs *gfxapi.State) S16ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S16ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the S16ˢ in a new memory pool.
func (s S16ˢ) Clone(ϟs *gfxapi.State) S16ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S16ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that S16ˢ points to.
func (s S16ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S16ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S16ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S16ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S16ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the int16 elements in this S16ˢ.
func (s S16ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int16 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int16, s.Count)
	for i := range res {
		if v, err := d.Int16(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst S16ˢ) Write(src []int16, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int16(int16(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S16ˢ) Copy(src S16ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S16ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a S16ᵖ to the i'th element in this S16ˢ.
func (s S16ˢ) Index(i uint64, ϟs *gfxapi.State) S16ᵖ {
	return S16ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the S16ˢ using start and end indices.
func (s S16ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S16ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S16ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the S16ˢ slice.
func (s S16ˢ) String() string {
	return fmt.Sprintf("int16(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type S32ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeS32ˢ returns a S32ˢ backed by a new memory pool.
func MakeS32ˢ(count uint64, ϟs *gfxapi.State) S32ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S32ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the S32ˢ in a new memory pool.
func (s S32ˢ) Clone(ϟs *gfxapi.State) S32ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S32ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that S32ˢ points to.
func (s S32ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S32ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the int32 elements in this S32ˢ.
func (s S32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int32, s.Count)
	for i := range res {
		if v, err := d.Int32(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst S32ˢ) Write(src []int32, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int32(int32(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S32ˢ) Copy(src S32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a S32ᵖ to the i'th element in this S32ˢ.
func (s S32ˢ) Index(i uint64, ϟs *gfxapi.State) S32ᵖ {
	return S32ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the S32ˢ using start and end indices.
func (s S32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S32ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the S32ˢ slice.
func (s S32ˢ) String() string {
	return fmt.Sprintf("int32(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type S64ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeS64ˢ returns a S64ˢ backed by a new memory pool.
func MakeS64ˢ(count uint64, ϟs *gfxapi.State) S64ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S64ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the S64ˢ in a new memory pool.
func (s S64ˢ) Clone(ϟs *gfxapi.State) S64ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S64ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that S64ˢ points to.
func (s S64ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S64ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the int64 elements in this S64ˢ.
func (s S64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int64, s.Count)
	for i := range res {
		if v, err := d.Int64(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst S64ˢ) Write(src []int64, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int64(int64(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S64ˢ) Copy(src S64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a S64ᵖ to the i'th element in this S64ˢ.
func (s S64ˢ) Index(i uint64, ϟs *gfxapi.State) S64ᵖ {
	return S64ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the S64ˢ using start and end indices.
func (s S64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S64ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the S64ˢ slice.
func (s S64ˢ) String() string {
	return fmt.Sprintf("int64(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type S8ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeS8ˢ returns a S8ˢ backed by a new memory pool.
func MakeS8ˢ(count uint64, ϟs *gfxapi.State) S8ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return S8ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the S8ˢ in a new memory pool.
func (s S8ˢ) Clone(ϟs *gfxapi.State) S8ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := S8ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that S8ˢ points to.
func (s S8ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s S8ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s S8ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s S8ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s S8ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the int8 elements in this S8ˢ.
func (s S8ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int8 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int8, s.Count)
	for i := range res {
		if v, err := d.Int8(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst S8ˢ) Write(src []int8, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int8(int8(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S8ˢ) Copy(src S8ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S8ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a S8ᵖ to the i'th element in this S8ˢ.
func (s S8ˢ) Index(i uint64, ϟs *gfxapi.State) S8ᵖ {
	return S8ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the S8ˢ using start and end indices.
func (s S8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S8ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return S8ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the S8ˢ slice.
func (s S8ˢ) String() string {
	return fmt.Sprintf("int8(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type U16ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeU16ˢ returns a U16ˢ backed by a new memory pool.
func MakeU16ˢ(count uint64, ϟs *gfxapi.State) U16ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U16ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the U16ˢ in a new memory pool.
func (s U16ˢ) Clone(ϟs *gfxapi.State) U16ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U16ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that U16ˢ points to.
func (s U16ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(2)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U16ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U16ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U16ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U16ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the uint16 elements in this U16ˢ.
func (s U16ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint16 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint16, s.Count)
	for i := range res {
		if v, err := d.Uint16(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst U16ˢ) Write(src []uint16, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint16(uint16(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U16ˢ) Copy(src U16ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U16ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a U16ᵖ to the i'th element in this U16ˢ.
func (s U16ˢ) Index(i uint64, ϟs *gfxapi.State) U16ᵖ {
	return U16ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the U16ˢ using start and end indices.
func (s U16ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U16ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U16ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the U16ˢ slice.
func (s U16ˢ) String() string {
	return fmt.Sprintf("uint16(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type U32ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeU32ˢ returns a U32ˢ backed by a new memory pool.
func MakeU32ˢ(count uint64, ϟs *gfxapi.State) U32ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U32ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the U32ˢ in a new memory pool.
func (s U32ˢ) Clone(ϟs *gfxapi.State) U32ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U32ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that U32ˢ points to.
func (s U32ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U32ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the uint32 elements in this U32ˢ.
func (s U32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint32, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst U32ˢ) Write(src []uint32, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U32ˢ) Copy(src U32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a U32ᵖ to the i'th element in this U32ˢ.
func (s U32ˢ) Index(i uint64, ϟs *gfxapi.State) U32ᵖ {
	return U32ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the U32ˢ using start and end indices.
func (s U32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U32ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the U32ˢ slice.
func (s U32ˢ) String() string {
	return fmt.Sprintf("uint32(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type U64ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeU64ˢ returns a U64ˢ backed by a new memory pool.
func MakeU64ˢ(count uint64, ϟs *gfxapi.State) U64ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U64ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the U64ˢ in a new memory pool.
func (s U64ˢ) Clone(ϟs *gfxapi.State) U64ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U64ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that U64ˢ points to.
func (s U64ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(8)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U64ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the uint64 elements in this U64ˢ.
func (s U64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint64, s.Count)
	for i := range res {
		if v, err := d.Uint64(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst U64ˢ) Write(src []uint64, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint64(uint64(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U64ˢ) Copy(src U64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a U64ᵖ to the i'th element in this U64ˢ.
func (s U64ˢ) Index(i uint64, ϟs *gfxapi.State) U64ᵖ {
	return U64ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the U64ˢ using start and end indices.
func (s U64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U64ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the U64ˢ slice.
func (s U64ˢ) String() string {
	return fmt.Sprintf("uint64(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type U8ˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeU8ˢ returns a U8ˢ backed by a new memory pool.
func MakeU8ˢ(count uint64, ϟs *gfxapi.State) U8ˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U8ˢ{Count: count, Pool: id}
}

// Clone returns a copy of the U8ˢ in a new memory pool.
func (s U8ˢ) Clone(ϟs *gfxapi.State) U8ˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U8ˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that U8ˢ points to.
func (s U8ˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U8ˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U8ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U8ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U8ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the uint8 elements in this U8ˢ.
func (s U8ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint8 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint8, s.Count)
	for i := range res {
		if v, err := d.Uint8(); err == nil {
			res[i] = v
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst U8ˢ) Write(src []uint8, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint8(uint8(src[i])); err != nil {
			panic(err)
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U8ˢ) Copy(src U8ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U8ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a U8ᵖ to the i'th element in this U8ˢ.
func (s U8ˢ) Index(i uint64, ϟs *gfxapi.State) U8ᵖ {
	return U8ᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the U8ˢ using start and end indices.
func (s U8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U8ˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the U8ˢ slice.
func (s U8ˢ) String() string {
	return fmt.Sprintf("uint8(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type U8ᵖˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeU8ᵖˢ returns a U8ᵖˢ backed by a new memory pool.
func MakeU8ᵖˢ(count uint64, ϟs *gfxapi.State) U8ᵖˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return U8ᵖˢ{Count: count, Pool: id}
}

// Clone returns a copy of the U8ᵖˢ in a new memory pool.
func (s U8ᵖˢ) Clone(ϟs *gfxapi.State) U8ᵖˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := U8ᵖˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that U8ᵖˢ points to.
func (s U8ᵖˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	if s.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Range returns the memory range this slice represents in the underlying pool.
func (s U8ᵖˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s U8ᵖˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s U8ᵖˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s U8ᵖˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the U8ᵖ elements in this U8ᵖˢ.
func (s U8ᵖˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []U8ᵖ {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]U8ᵖ, s.Count)
	for i := range res {
		if s.Pool == memory.ApplicationPool {
			ptr, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			res[i] = NewU8ᵖ(memory.Pointer(ptr))
		} else {
			if err := d.Value(&res[i]); err != nil {
				panic(err)
			}
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst U8ᵖˢ) Write(src []U8ᵖ, ϟs *gfxapi.State) uint64 {
	count := min(dst.Count, uint64(len(src)))
	dst = dst.Slice(0, count, ϟs)
	e := dst.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if dst.Pool == memory.ApplicationPool {
			if err := binary.WriteUint(e, ϟs.Architecture.PointerSize*8, uint64(src[i].Address)); err != nil {
				panic(err)
			}
		} else {
			if err := e.Value(&src[i]); err != nil {
				panic(err)
			}
		}
	}
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U8ᵖˢ) Copy(src U8ᵖˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U8ᵖˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	if (dst.Pool == memory.ApplicationPool) != (src.Pool == memory.ApplicationPool) {
		dst.Write(src.Read(ϟs, ϟd, ϟl), ϟs)
	} else {
		ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	}
	return dst, src
}

// Index returns a U8ᵖᵖ to the i'th element in this U8ᵖˢ.
func (s U8ᵖˢ) Index(i uint64, ϟs *gfxapi.State) U8ᵖᵖ {
	return U8ᵖᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the U8ᵖˢ using start and end indices.
func (s U8ᵖˢ) Slice(start, end uint64, ϟs *gfxapi.State) U8ᵖˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return U8ᵖˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the U8ᵖˢ slice.
func (s U8ᵖˢ) String() string {
	return fmt.Sprintf("U8ᵖ(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type Voidˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeVoidˢ returns a Voidˢ backed by a new memory pool.
func MakeVoidˢ(count uint64, ϟs *gfxapi.State) Voidˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Voidˢ{Count: count, Pool: id}
}

// Clone returns a copy of the Voidˢ in a new memory pool.
func (s Voidˢ) Clone(ϟs *gfxapi.State) Voidˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Voidˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that Voidˢ points to.
func (s Voidˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Voidˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Voidˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Voidˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Voidˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Index returns a Voidᵖ to the i'th element in this Voidˢ.
func (s Voidˢ) Index(i uint64, ϟs *gfxapi.State) Voidᵖ {
	return Voidᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the Voidˢ using start and end indices.
func (s Voidˢ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Voidˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the Voidˢ slice.
func (s Voidˢ) String() string {
	return fmt.Sprintf("void(%v@%v)[%d]", s.Base, s.Pool, s.Count)
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
	return fmt.Sprintf("cmd_clone(src: %v, cnt: %v)", a.Src, a.Cnt)
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
func (c *CmdClone) API() gfxapi.API                  { return api{} }
func (c *CmdClone) TypeID() atom.TypeID              { return 0 }
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
	return fmt.Sprintf("cmd_make(cnt: %v)", a.Cnt)
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
func (c *CmdMake) API() gfxapi.API                  { return api{} }
func (c *CmdMake) TypeID() atom.TypeID              { return 1 }
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
	return fmt.Sprintf("cmd_copy(src: %v, cnt: %v)", a.Src, a.Cnt)
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
func (c *CmdCopy) API() gfxapi.API                  { return api{} }
func (c *CmdCopy) TypeID() atom.TypeID              { return 2 }
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
	return fmt.Sprintf("cmd_charslice_to_string(s: %v, len: %v)", a.S, a.Len)
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
func (c *CmdCharsliceToString) API() gfxapi.API                  { return api{} }
func (c *CmdCharsliceToString) TypeID() atom.TypeID              { return 3 }
func (c *CmdCharsliceToString) Flags() atom.Flags                { return 0 }
func (a *CmdCharsliceToString) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid
////////////////////////////////////////////////////////////////////////////////
type CmdVoid struct {
	binary.Generate
	observations atom.Observations
}

func (a *CmdVoid) String() string {
	return fmt.Sprintf("cmd_void()")
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
func (c *CmdVoid) API() gfxapi.API                  { return api{} }
func (c *CmdVoid) TypeID() atom.TypeID              { return 4 }
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
	return fmt.Sprintf("cmd_unknown_ret() → %v", a.Result)
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
func (c *CmdUnknownRet) API() gfxapi.API                  { return api{} }
func (c *CmdUnknownRet) TypeID() atom.TypeID              { return 5 }
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
	return fmt.Sprintf("cmd_unknown_write_ptr(p: %v)", a.P)
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
func (c *CmdUnknownWritePtr) API() gfxapi.API                  { return api{} }
func (c *CmdUnknownWritePtr) TypeID() atom.TypeID              { return 6 }
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
	return fmt.Sprintf("cmd_unknown_write_slice(a: %v)", a.A)
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
func (c *CmdUnknownWriteSlice) API() gfxapi.API                  { return api{} }
func (c *CmdUnknownWriteSlice) TypeID() atom.TypeID              { return 7 }
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
	return fmt.Sprintf("cmd_void_u8(a: %v)", a.A)
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
func (c *CmdVoidU8) API() gfxapi.API                  { return api{} }
func (c *CmdVoidU8) TypeID() atom.TypeID              { return 8 }
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
	return fmt.Sprintf("cmd_void_s8(a: %v)", a.A)
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
func (c *CmdVoidS8) API() gfxapi.API                  { return api{} }
func (c *CmdVoidS8) TypeID() atom.TypeID              { return 9 }
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
	return fmt.Sprintf("cmd_void_u16(a: %v)", a.A)
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
func (c *CmdVoidU16) API() gfxapi.API                  { return api{} }
func (c *CmdVoidU16) TypeID() atom.TypeID              { return 10 }
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
	return fmt.Sprintf("cmd_void_s16(a: %v)", a.A)
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
func (c *CmdVoidS16) API() gfxapi.API                  { return api{} }
func (c *CmdVoidS16) TypeID() atom.TypeID              { return 11 }
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
	return fmt.Sprintf("cmd_void_f32(a: %v)", a.A)
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
func (c *CmdVoidF32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidF32) TypeID() atom.TypeID              { return 12 }
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
	return fmt.Sprintf("cmd_void_u32(a: %v)", a.A)
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
func (c *CmdVoidU32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidU32) TypeID() atom.TypeID              { return 13 }
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
	return fmt.Sprintf("cmd_void_s32(a: %v)", a.A)
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
func (c *CmdVoidS32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidS32) TypeID() atom.TypeID              { return 14 }
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
	return fmt.Sprintf("cmd_void_f64(a: %v)", a.A)
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
func (c *CmdVoidF64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidF64) TypeID() atom.TypeID              { return 15 }
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
	return fmt.Sprintf("cmd_void_u64(a: %v)", a.A)
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
func (c *CmdVoidU64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidU64) TypeID() atom.TypeID              { return 16 }
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
	return fmt.Sprintf("cmd_void_s64(a: %v)", a.A)
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
func (c *CmdVoidS64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidS64) TypeID() atom.TypeID              { return 17 }
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
	return fmt.Sprintf("cmd_void_bool(a: %v)", a.A)
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
func (c *CmdVoidBool) API() gfxapi.API                  { return api{} }
func (c *CmdVoidBool) TypeID() atom.TypeID              { return 18 }
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
	return fmt.Sprintf("cmd_void_string(a: %v)", a.A)
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
func (c *CmdVoidString) API() gfxapi.API                  { return api{} }
func (c *CmdVoidString) TypeID() atom.TypeID              { return 19 }
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
	return fmt.Sprintf("cmd_void_3_strings(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
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
func (c *CmdVoid3Strings) API() gfxapi.API                  { return api{} }
func (c *CmdVoid3Strings) TypeID() atom.TypeID              { return 20 }
func (c *CmdVoid3Strings) Flags() atom.Flags                { return 0 }
func (a *CmdVoid3Strings) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoid3InArrays
////////////////////////////////////////////////////////////////////////////////
type CmdVoid3InArrays struct {
	binary.Generate
	observations atom.Observations
	A            U8ᵖ
	B            U8ᵖᵖ
	C            Intᵖ
}

func (a *CmdVoid3InArrays) String() string {
	return fmt.Sprintf("cmd_void_3_in_arrays(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
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
func (c *CmdVoid3InArrays) API() gfxapi.API                  { return api{} }
func (c *CmdVoid3InArrays) TypeID() atom.TypeID              { return 21 }
func (c *CmdVoid3InArrays) Flags() atom.Flags                { return 0 }
func (a *CmdVoid3InArrays) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CmdVoidReadU8
////////////////////////////////////////////////////////////////////////////////
type CmdVoidReadU8 struct {
	binary.Generate
	observations atom.Observations
	A            U8ᵖ
}

func (a *CmdVoidReadU8) String() string {
	return fmt.Sprintf("cmd_void_read_u8(a: %v)", a.A)
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
func (c *CmdVoidReadU8) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadU8) TypeID() atom.TypeID              { return 22 }
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
	return fmt.Sprintf("cmd_void_read_s8(a: %v)", a.A)
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
func (c *CmdVoidReadS8) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadS8) TypeID() atom.TypeID              { return 23 }
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
	return fmt.Sprintf("cmd_void_read_u16(a: %v)", a.A)
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
func (c *CmdVoidReadU16) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadU16) TypeID() atom.TypeID              { return 24 }
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
	return fmt.Sprintf("cmd_void_read_s16(a: %v)", a.A)
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
func (c *CmdVoidReadS16) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadS16) TypeID() atom.TypeID              { return 25 }
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
	return fmt.Sprintf("cmd_void_read_f32(a: %v)", a.A)
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
func (c *CmdVoidReadF32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadF32) TypeID() atom.TypeID              { return 26 }
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
	return fmt.Sprintf("cmd_void_read_u32(a: %v)", a.A)
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
func (c *CmdVoidReadU32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadU32) TypeID() atom.TypeID              { return 27 }
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
	return fmt.Sprintf("cmd_void_read_s32(a: %v)", a.A)
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
func (c *CmdVoidReadS32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadS32) TypeID() atom.TypeID              { return 28 }
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
	return fmt.Sprintf("cmd_void_read_f64(a: %v)", a.A)
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
func (c *CmdVoidReadF64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadF64) TypeID() atom.TypeID              { return 29 }
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
	return fmt.Sprintf("cmd_void_read_u64(a: %v)", a.A)
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
func (c *CmdVoidReadU64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadU64) TypeID() atom.TypeID              { return 30 }
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
	return fmt.Sprintf("cmd_void_read_s64(a: %v)", a.A)
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
func (c *CmdVoidReadS64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadS64) TypeID() atom.TypeID              { return 31 }
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
	return fmt.Sprintf("cmd_void_read_bool(a: %v)", a.A)
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
func (c *CmdVoidReadBool) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadBool) TypeID() atom.TypeID              { return 32 }
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
	return fmt.Sprintf("cmd_void_read_ptrs(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
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
func (c *CmdVoidReadPtrs) API() gfxapi.API                  { return api{} }
func (c *CmdVoidReadPtrs) TypeID() atom.TypeID              { return 33 }
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
	return fmt.Sprintf("cmd_void_write_u8(a: %v)", a.A)
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
func (c *CmdVoidWriteU8) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteU8) TypeID() atom.TypeID              { return 34 }
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
	return fmt.Sprintf("cmd_void_write_s8(a: %v)", a.A)
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
func (c *CmdVoidWriteS8) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteS8) TypeID() atom.TypeID              { return 35 }
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
	return fmt.Sprintf("cmd_void_write_u16(a: %v)", a.A)
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
func (c *CmdVoidWriteU16) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteU16) TypeID() atom.TypeID              { return 36 }
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
	return fmt.Sprintf("cmd_void_write_s16(a: %v)", a.A)
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
func (c *CmdVoidWriteS16) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteS16) TypeID() atom.TypeID              { return 37 }
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
	return fmt.Sprintf("cmd_void_write_f32(a: %v)", a.A)
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
func (c *CmdVoidWriteF32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteF32) TypeID() atom.TypeID              { return 38 }
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
	return fmt.Sprintf("cmd_void_write_u32(a: %v)", a.A)
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
func (c *CmdVoidWriteU32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteU32) TypeID() atom.TypeID              { return 39 }
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
	return fmt.Sprintf("cmd_void_write_s32(a: %v)", a.A)
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
func (c *CmdVoidWriteS32) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteS32) TypeID() atom.TypeID              { return 40 }
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
	return fmt.Sprintf("cmd_void_write_f64(a: %v)", a.A)
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
func (c *CmdVoidWriteF64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteF64) TypeID() atom.TypeID              { return 41 }
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
	return fmt.Sprintf("cmd_void_write_u64(a: %v)", a.A)
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
func (c *CmdVoidWriteU64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteU64) TypeID() atom.TypeID              { return 42 }
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
	return fmt.Sprintf("cmd_void_write_s64(a: %v)", a.A)
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
func (c *CmdVoidWriteS64) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteS64) TypeID() atom.TypeID              { return 43 }
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
	return fmt.Sprintf("cmd_void_write_bool(a: %v)", a.A)
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
func (c *CmdVoidWriteBool) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWriteBool) TypeID() atom.TypeID              { return 44 }
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
	return fmt.Sprintf("cmd_void_write_ptrs(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
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
func (c *CmdVoidWritePtrs) API() gfxapi.API                  { return api{} }
func (c *CmdVoidWritePtrs) TypeID() atom.TypeID              { return 45 }
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
	return fmt.Sprintf("cmd_u8() → %v", a.Result)
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
func (c *CmdU8) API() gfxapi.API                  { return api{} }
func (c *CmdU8) TypeID() atom.TypeID              { return 46 }
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
	return fmt.Sprintf("cmd_s8() → %v", a.Result)
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
func (c *CmdS8) API() gfxapi.API                  { return api{} }
func (c *CmdS8) TypeID() atom.TypeID              { return 47 }
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
	return fmt.Sprintf("cmd_u16() → %v", a.Result)
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
func (c *CmdU16) API() gfxapi.API                  { return api{} }
func (c *CmdU16) TypeID() atom.TypeID              { return 48 }
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
	return fmt.Sprintf("cmd_s16() → %v", a.Result)
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
func (c *CmdS16) API() gfxapi.API                  { return api{} }
func (c *CmdS16) TypeID() atom.TypeID              { return 49 }
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
	return fmt.Sprintf("cmd_f32() → %v", a.Result)
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
func (c *CmdF32) API() gfxapi.API                  { return api{} }
func (c *CmdF32) TypeID() atom.TypeID              { return 50 }
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
	return fmt.Sprintf("cmd_u32() → %v", a.Result)
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
func (c *CmdU32) API() gfxapi.API                  { return api{} }
func (c *CmdU32) TypeID() atom.TypeID              { return 51 }
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
	return fmt.Sprintf("cmd_s32() → %v", a.Result)
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
func (c *CmdS32) API() gfxapi.API                  { return api{} }
func (c *CmdS32) TypeID() atom.TypeID              { return 52 }
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
	return fmt.Sprintf("cmd_f64() → %v", a.Result)
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
func (c *CmdF64) API() gfxapi.API                  { return api{} }
func (c *CmdF64) TypeID() atom.TypeID              { return 53 }
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
	return fmt.Sprintf("cmd_u64() → %v", a.Result)
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
func (c *CmdU64) API() gfxapi.API                  { return api{} }
func (c *CmdU64) TypeID() atom.TypeID              { return 54 }
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
	return fmt.Sprintf("cmd_s64() → %v", a.Result)
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
func (c *CmdS64) API() gfxapi.API                  { return api{} }
func (c *CmdS64) TypeID() atom.TypeID              { return 55 }
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
	return fmt.Sprintf("cmd_bool() → %v", a.Result)
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
func (c *CmdBool) API() gfxapi.API                  { return api{} }
func (c *CmdBool) TypeID() atom.TypeID              { return 56 }
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
	return fmt.Sprintf("cmd_string() → %v", a.Result)
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
func (c *CmdString) API() gfxapi.API                  { return api{} }
func (c *CmdString) TypeID() atom.TypeID              { return 57 }
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
	return fmt.Sprintf("cmd_pointer() → %v", a.Result)
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
func (c *CmdPointer) API() gfxapi.API                  { return api{} }
func (c *CmdPointer) TypeID() atom.TypeID              { return 58 }
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
	return fmt.Sprintf("cmd_void_3_remapped(a: %v, b: %v, c: %v)", a.A, a.B, a.C)
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
func (c *CmdVoid3Remapped) API() gfxapi.API                  { return api{} }
func (c *CmdVoid3Remapped) TypeID() atom.TypeID              { return 59 }
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
	return fmt.Sprintf("cmd_void_in_array_of_remapped(a: %v)", a.A)
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
func (c *CmdVoidInArrayOfRemapped) API() gfxapi.API                  { return api{} }
func (c *CmdVoidInArrayOfRemapped) TypeID() atom.TypeID              { return 60 }
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
	return fmt.Sprintf("cmd_void_out_array_of_remapped(a: %v)", a.A)
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
func (c *CmdVoidOutArrayOfRemapped) API() gfxapi.API                  { return api{} }
func (c *CmdVoidOutArrayOfRemapped) TypeID() atom.TypeID              { return 61 }
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
	return fmt.Sprintf("cmd_void_out_array_of_unknown_remapped(a: %v)", a.A)
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
func (c *CmdVoidOutArrayOfUnknownRemapped) API() gfxapi.API                  { return api{} }
func (c *CmdVoidOutArrayOfUnknownRemapped) TypeID() atom.TypeID              { return 62 }
func (c *CmdVoidOutArrayOfUnknownRemapped) Flags() atom.Flags                { return 0 }
func (a *CmdVoidOutArrayOfUnknownRemapped) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// class Tester
////////////////////////////////////////////////////////////////////////////////
type Tester struct {
	binary.Generate
	CreatedAt atom.ID
	A         Imported
	B         Included
}

func (c *Tester) Init() {
	c.A.Init()
	c.B.Init()
}
func (c *Tester) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Included
////////////////////////////////////////////////////////////////////////////////
type Included struct {
	binary.Generate
	CreatedAt atom.ID
	S         string
}

func (c *Included) Init() {
}
func (c *Included) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// Globals
////////////////////////////////////////////////////////////////////////////////
type Globals struct {
	binary.Generate
	Buf U8ˢ
	Str string
}

func (g *Globals) Init() {
}
func NewCmdClone(Src memory.Pointer, Cnt uint32) *CmdClone {
	return &CmdClone{Src: NewU8ᵖ(Src), Cnt: Cnt}
}
func NewCmdMake(Cnt uint32) *CmdMake {
	return &CmdMake{Cnt: Cnt}
}
func NewCmdCopy(Src memory.Pointer, Cnt uint32) *CmdCopy {
	return &CmdCopy{Src: NewU8ᵖ(Src), Cnt: Cnt}
}
func NewCmdCharsliceToString(S memory.Pointer, Len uint32) *CmdCharsliceToString {
	return &CmdCharsliceToString{S: NewCharᵖ(S), Len: Len}
}
func NewCmdVoid() *CmdVoid {
	return &CmdVoid{}
}
func NewCmdUnknownRet(Result int64) *CmdUnknownRet {
	return &CmdUnknownRet{Result: Result}
}
func NewCmdUnknownWritePtr(P memory.Pointer) *CmdUnknownWritePtr {
	return &CmdUnknownWritePtr{P: NewIntᵖ(P)}
}
func NewCmdUnknownWriteSlice(A memory.Pointer) *CmdUnknownWriteSlice {
	return &CmdUnknownWriteSlice{A: NewIntᵖ(A)}
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
	return &CmdVoid3InArrays{A: NewU8ᵖ(A), B: NewU8ᵖᵖ(B), C: NewIntᵖ(C)}
}
func NewCmdVoidReadU8(A memory.Pointer) *CmdVoidReadU8 {
	return &CmdVoidReadU8{A: NewU8ᵖ(A)}
}
func NewCmdVoidReadS8(A memory.Pointer) *CmdVoidReadS8 {
	return &CmdVoidReadS8{A: NewS8ᵖ(A)}
}
func NewCmdVoidReadU16(A memory.Pointer) *CmdVoidReadU16 {
	return &CmdVoidReadU16{A: NewU16ᵖ(A)}
}
func NewCmdVoidReadS16(A memory.Pointer) *CmdVoidReadS16 {
	return &CmdVoidReadS16{A: NewS16ᵖ(A)}
}
func NewCmdVoidReadF32(A memory.Pointer) *CmdVoidReadF32 {
	return &CmdVoidReadF32{A: NewF32ᵖ(A)}
}
func NewCmdVoidReadU32(A memory.Pointer) *CmdVoidReadU32 {
	return &CmdVoidReadU32{A: NewU32ᵖ(A)}
}
func NewCmdVoidReadS32(A memory.Pointer) *CmdVoidReadS32 {
	return &CmdVoidReadS32{A: NewS32ᵖ(A)}
}
func NewCmdVoidReadF64(A memory.Pointer) *CmdVoidReadF64 {
	return &CmdVoidReadF64{A: NewF64ᵖ(A)}
}
func NewCmdVoidReadU64(A memory.Pointer) *CmdVoidReadU64 {
	return &CmdVoidReadU64{A: NewU64ᵖ(A)}
}
func NewCmdVoidReadS64(A memory.Pointer) *CmdVoidReadS64 {
	return &CmdVoidReadS64{A: NewS64ᵖ(A)}
}
func NewCmdVoidReadBool(A memory.Pointer) *CmdVoidReadBool {
	return &CmdVoidReadBool{A: NewBoolᵖ(A)}
}
func NewCmdVoidReadPtrs(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoidReadPtrs {
	return &CmdVoidReadPtrs{A: NewF32ᵖ(A), B: NewU16ᵖ(B), C: NewBoolᵖ(C)}
}
func NewCmdVoidWriteU8(A memory.Pointer) *CmdVoidWriteU8 {
	return &CmdVoidWriteU8{A: NewU8ᵖ(A)}
}
func NewCmdVoidWriteS8(A memory.Pointer) *CmdVoidWriteS8 {
	return &CmdVoidWriteS8{A: NewS8ᵖ(A)}
}
func NewCmdVoidWriteU16(A memory.Pointer) *CmdVoidWriteU16 {
	return &CmdVoidWriteU16{A: NewU16ᵖ(A)}
}
func NewCmdVoidWriteS16(A memory.Pointer) *CmdVoidWriteS16 {
	return &CmdVoidWriteS16{A: NewS16ᵖ(A)}
}
func NewCmdVoidWriteF32(A memory.Pointer) *CmdVoidWriteF32 {
	return &CmdVoidWriteF32{A: NewF32ᵖ(A)}
}
func NewCmdVoidWriteU32(A memory.Pointer) *CmdVoidWriteU32 {
	return &CmdVoidWriteU32{A: NewU32ᵖ(A)}
}
func NewCmdVoidWriteS32(A memory.Pointer) *CmdVoidWriteS32 {
	return &CmdVoidWriteS32{A: NewS32ᵖ(A)}
}
func NewCmdVoidWriteF64(A memory.Pointer) *CmdVoidWriteF64 {
	return &CmdVoidWriteF64{A: NewF64ᵖ(A)}
}
func NewCmdVoidWriteU64(A memory.Pointer) *CmdVoidWriteU64 {
	return &CmdVoidWriteU64{A: NewU64ᵖ(A)}
}
func NewCmdVoidWriteS64(A memory.Pointer) *CmdVoidWriteS64 {
	return &CmdVoidWriteS64{A: NewS64ᵖ(A)}
}
func NewCmdVoidWriteBool(A memory.Pointer) *CmdVoidWriteBool {
	return &CmdVoidWriteBool{A: NewBoolᵖ(A)}
}
func NewCmdVoidWritePtrs(A memory.Pointer, B memory.Pointer, C memory.Pointer) *CmdVoidWritePtrs {
	return &CmdVoidWritePtrs{A: NewF32ᵖ(A), B: NewU16ᵖ(B), C: NewBoolᵖ(C)}
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
	return &CmdPointer{Result: NewVoidᵖ(Result)}
}
func NewCmdVoid3Remapped(A remapped, B remapped, C remapped) *CmdVoid3Remapped {
	return &CmdVoid3Remapped{A: A, B: B, C: C}
}
func NewCmdVoidInArrayOfRemapped(A memory.Pointer) *CmdVoidInArrayOfRemapped {
	return &CmdVoidInArrayOfRemapped{A: NewRemappedᵖ(A)}
}
func NewCmdVoidOutArrayOfRemapped(A memory.Pointer) *CmdVoidOutArrayOfRemapped {
	return &CmdVoidOutArrayOfRemapped{A: NewRemappedᵖ(A)}
}
func NewCmdVoidOutArrayOfUnknownRemapped(A memory.Pointer) *CmdVoidOutArrayOfUnknownRemapped {
	return &CmdVoidOutArrayOfUnknownRemapped{A: NewRemappedᵖ(A)}
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
	atom.Register(atom.TypeInfo{
		Name: "cmd_clone",
		Docs: "[]",
		ID:   0,
		New:  func() atom.Atom { return &CmdClone{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_make",
		Docs: "[]",
		ID:   1,
		New:  func() atom.Atom { return &CmdMake{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_copy",
		Docs: "[]",
		ID:   2,
		New:  func() atom.Atom { return &CmdCopy{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_charslice_to_string",
		Docs: "[]",
		ID:   3,
		New:  func() atom.Atom { return &CmdCharsliceToString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void",
		Docs: "[]",
		ID:   4,
		New:  func() atom.Atom { return &CmdVoid{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_unknown_ret",
		Docs: "[]",
		ID:   5,
		New:  func() atom.Atom { return &CmdUnknownRet{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_unknown_write_ptr",
		Docs: "[]",
		ID:   6,
		New:  func() atom.Atom { return &CmdUnknownWritePtr{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_unknown_write_slice",
		Docs: "[]",
		ID:   7,
		New:  func() atom.Atom { return &CmdUnknownWriteSlice{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_u8",
		Docs: "[]",
		ID:   8,
		New:  func() atom.Atom { return &CmdVoidU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_s8",
		Docs: "[]",
		ID:   9,
		New:  func() atom.Atom { return &CmdVoidS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_u16",
		Docs: "[]",
		ID:   10,
		New:  func() atom.Atom { return &CmdVoidU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_s16",
		Docs: "[]",
		ID:   11,
		New:  func() atom.Atom { return &CmdVoidS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_f32",
		Docs: "[]",
		ID:   12,
		New:  func() atom.Atom { return &CmdVoidF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_u32",
		Docs: "[]",
		ID:   13,
		New:  func() atom.Atom { return &CmdVoidU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_s32",
		Docs: "[]",
		ID:   14,
		New:  func() atom.Atom { return &CmdVoidS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_f64",
		Docs: "[]",
		ID:   15,
		New:  func() atom.Atom { return &CmdVoidF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_u64",
		Docs: "[]",
		ID:   16,
		New:  func() atom.Atom { return &CmdVoidU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_s64",
		Docs: "[]",
		ID:   17,
		New:  func() atom.Atom { return &CmdVoidS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_bool",
		Docs: "[]",
		ID:   18,
		New:  func() atom.Atom { return &CmdVoidBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_string",
		Docs: "[]",
		ID:   19,
		New:  func() atom.Atom { return &CmdVoidString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_3_strings",
		Docs: "[]",
		ID:   20,
		New:  func() atom.Atom { return &CmdVoid3Strings{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_3_in_arrays",
		Docs: "[]",
		ID:   21,
		New:  func() atom.Atom { return &CmdVoid3InArrays{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_u8",
		Docs: "[]",
		ID:   22,
		New:  func() atom.Atom { return &CmdVoidReadU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_s8",
		Docs: "[]",
		ID:   23,
		New:  func() atom.Atom { return &CmdVoidReadS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_u16",
		Docs: "[]",
		ID:   24,
		New:  func() atom.Atom { return &CmdVoidReadU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_s16",
		Docs: "[]",
		ID:   25,
		New:  func() atom.Atom { return &CmdVoidReadS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_f32",
		Docs: "[]",
		ID:   26,
		New:  func() atom.Atom { return &CmdVoidReadF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_u32",
		Docs: "[]",
		ID:   27,
		New:  func() atom.Atom { return &CmdVoidReadU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_s32",
		Docs: "[]",
		ID:   28,
		New:  func() atom.Atom { return &CmdVoidReadS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_f64",
		Docs: "[]",
		ID:   29,
		New:  func() atom.Atom { return &CmdVoidReadF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_u64",
		Docs: "[]",
		ID:   30,
		New:  func() atom.Atom { return &CmdVoidReadU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_s64",
		Docs: "[]",
		ID:   31,
		New:  func() atom.Atom { return &CmdVoidReadS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_bool",
		Docs: "[]",
		ID:   32,
		New:  func() atom.Atom { return &CmdVoidReadBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_read_ptrs",
		Docs: "[]",
		ID:   33,
		New:  func() atom.Atom { return &CmdVoidReadPtrs{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_u8",
		Docs: "[]",
		ID:   34,
		New:  func() atom.Atom { return &CmdVoidWriteU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_s8",
		Docs: "[]",
		ID:   35,
		New:  func() atom.Atom { return &CmdVoidWriteS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_u16",
		Docs: "[]",
		ID:   36,
		New:  func() atom.Atom { return &CmdVoidWriteU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_s16",
		Docs: "[]",
		ID:   37,
		New:  func() atom.Atom { return &CmdVoidWriteS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_f32",
		Docs: "[]",
		ID:   38,
		New:  func() atom.Atom { return &CmdVoidWriteF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_u32",
		Docs: "[]",
		ID:   39,
		New:  func() atom.Atom { return &CmdVoidWriteU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_s32",
		Docs: "[]",
		ID:   40,
		New:  func() atom.Atom { return &CmdVoidWriteS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_f64",
		Docs: "[]",
		ID:   41,
		New:  func() atom.Atom { return &CmdVoidWriteF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_u64",
		Docs: "[]",
		ID:   42,
		New:  func() atom.Atom { return &CmdVoidWriteU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_s64",
		Docs: "[]",
		ID:   43,
		New:  func() atom.Atom { return &CmdVoidWriteS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_bool",
		Docs: "[]",
		ID:   44,
		New:  func() atom.Atom { return &CmdVoidWriteBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_write_ptrs",
		Docs: "[]",
		ID:   45,
		New:  func() atom.Atom { return &CmdVoidWritePtrs{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_u8",
		Docs: "[]",
		ID:   46,
		New:  func() atom.Atom { return &CmdU8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_s8",
		Docs: "[]",
		ID:   47,
		New:  func() atom.Atom { return &CmdS8{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_u16",
		Docs: "[]",
		ID:   48,
		New:  func() atom.Atom { return &CmdU16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_s16",
		Docs: "[]",
		ID:   49,
		New:  func() atom.Atom { return &CmdS16{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_f32",
		Docs: "[]",
		ID:   50,
		New:  func() atom.Atom { return &CmdF32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_u32",
		Docs: "[]",
		ID:   51,
		New:  func() atom.Atom { return &CmdU32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_s32",
		Docs: "[]",
		ID:   52,
		New:  func() atom.Atom { return &CmdS32{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_f64",
		Docs: "[]",
		ID:   53,
		New:  func() atom.Atom { return &CmdF64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_u64",
		Docs: "[]",
		ID:   54,
		New:  func() atom.Atom { return &CmdU64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_s64",
		Docs: "[]",
		ID:   55,
		New:  func() atom.Atom { return &CmdS64{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_bool",
		Docs: "[]",
		ID:   56,
		New:  func() atom.Atom { return &CmdBool{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_string",
		Docs: "[]",
		ID:   57,
		New:  func() atom.Atom { return &CmdString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_pointer",
		Docs: "[]",
		ID:   58,
		New:  func() atom.Atom { return &CmdPointer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_3_remapped",
		Docs: "[]",
		ID:   59,
		New:  func() atom.Atom { return &CmdVoid3Remapped{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_in_array_of_remapped",
		Docs: "[]",
		ID:   60,
		New:  func() atom.Atom { return &CmdVoidInArrayOfRemapped{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_out_array_of_remapped",
		Docs: "[]",
		ID:   61,
		New:  func() atom.Atom { return &CmdVoidOutArrayOfRemapped{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "cmd_void_out_array_of_unknown_remapped",
		Docs: "[]",
		ID:   62,
		New:  func() atom.Atom { return &CmdVoidOutArrayOfUnknownRemapped{} },
	})
}
func min(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}
