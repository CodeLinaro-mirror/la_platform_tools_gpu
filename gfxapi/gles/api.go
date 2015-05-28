////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/database"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/log"
	"android.googlesource.com/platform/tools/gpu/memory"
)

type RenderbufferId uint32
type TextureId uint32
type FramebufferId uint32
type BufferId uint32
type ShaderId uint32
type ProgramId uint32
type VertexArrayId uint32
type QueryId uint32
type UniformLocation int32
type AttributeLocation uint32

// IndicesPointer is a pointer to a void element.
type IndicesPointer struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewIndicesPointer returns a IndicesPointer that points to addr in the application pool.
func NewIndicesPointer(addr memory.Pointer) IndicesPointer {
	return IndicesPointer{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p IndicesPointer) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that IndicesPointer points to.
func (p IndicesPointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p IndicesPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the IndicesPointer pointer.
func (p IndicesPointer) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// VertexPointer is a pointer to a void element.
type VertexPointer struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewVertexPointer returns a VertexPointer that points to addr in the application pool.
func NewVertexPointer(addr memory.Pointer) VertexPointer {
	return VertexPointer{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p VertexPointer) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that VertexPointer points to.
func (p VertexPointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p VertexPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the VertexPointer pointer.
func (p VertexPointer) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// TexturePointer is a pointer to a void element.
type TexturePointer struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewTexturePointer returns a TexturePointer that points to addr in the application pool.
func NewTexturePointer(addr memory.Pointer) TexturePointer {
	return TexturePointer{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p TexturePointer) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that TexturePointer points to.
func (p TexturePointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p TexturePointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the TexturePointer pointer.
func (p TexturePointer) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// BufferDataPointer is a pointer to a void element.
type BufferDataPointer struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewBufferDataPointer returns a BufferDataPointer that points to addr in the application pool.
func NewBufferDataPointer(addr memory.Pointer) BufferDataPointer {
	return BufferDataPointer{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p BufferDataPointer) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that BufferDataPointer points to.
func (p BufferDataPointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p BufferDataPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the BufferDataPointer pointer.
func (p BufferDataPointer) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

type ContextID uint32
type ThreadID uint32
type EGLBoolean int64
type EGLint int64

// EGLConfig is a pointer to a void element.
type EGLConfig struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewEGLConfig returns a EGLConfig that points to addr in the application pool.
func NewEGLConfig(addr memory.Pointer) EGLConfig {
	return EGLConfig{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p EGLConfig) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that EGLConfig points to.
func (p EGLConfig) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLConfig) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the EGLConfig pointer.
func (p EGLConfig) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// EGLContext is a pointer to a void element.
type EGLContext struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewEGLContext returns a EGLContext that points to addr in the application pool.
func NewEGLContext(addr memory.Pointer) EGLContext {
	return EGLContext{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p EGLContext) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that EGLContext points to.
func (p EGLContext) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLContext) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the EGLContext pointer.
func (p EGLContext) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// EGLDisplay is a pointer to a void element.
type EGLDisplay struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewEGLDisplay returns a EGLDisplay that points to addr in the application pool.
func NewEGLDisplay(addr memory.Pointer) EGLDisplay {
	return EGLDisplay{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p EGLDisplay) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that EGLDisplay points to.
func (p EGLDisplay) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLDisplay) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the EGLDisplay pointer.
func (p EGLDisplay) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// EGLSurface is a pointer to a void element.
type EGLSurface struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewEGLSurface returns a EGLSurface that points to addr in the application pool.
func NewEGLSurface(addr memory.Pointer) EGLSurface {
	return EGLSurface{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p EGLSurface) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that EGLSurface points to.
func (p EGLSurface) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLSurface) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the EGLSurface pointer.
func (p EGLSurface) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// GLXContext is a pointer to a void element.
type GLXContext struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewGLXContext returns a GLXContext that points to addr in the application pool.
func NewGLXContext(addr memory.Pointer) GLXContext {
	return GLXContext{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p GLXContext) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that GLXContext points to.
func (p GLXContext) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p GLXContext) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the GLXContext pointer.
func (p GLXContext) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// GLXDrawable is a pointer to a void element.
type GLXDrawable struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewGLXDrawable returns a GLXDrawable that points to addr in the application pool.
func NewGLXDrawable(addr memory.Pointer) GLXDrawable {
	return GLXDrawable{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p GLXDrawable) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that GLXDrawable points to.
func (p GLXDrawable) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p GLXDrawable) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the GLXDrawable pointer.
func (p GLXDrawable) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// HGLRC is a pointer to a void element.
type HGLRC struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewHGLRC returns a HGLRC that points to addr in the application pool.
func NewHGLRC(addr memory.Pointer) HGLRC {
	return HGLRC{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p HGLRC) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that HGLRC points to.
func (p HGLRC) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p HGLRC) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the HGLRC pointer.
func (p HGLRC) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// HDC is a pointer to a void element.
type HDC struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewHDC returns a HDC that points to addr in the application pool.
func NewHDC(addr memory.Pointer) HDC {
	return HDC{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p HDC) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that HDC points to.
func (p HDC) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p HDC) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the HDC pointer.
func (p HDC) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

type BOOL int64
type CGLError int64

// CGLPixelFormatObj is a pointer to a void element.
type CGLPixelFormatObj struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewCGLPixelFormatObj returns a CGLPixelFormatObj that points to addr in the application pool.
func NewCGLPixelFormatObj(addr memory.Pointer) CGLPixelFormatObj {
	return CGLPixelFormatObj{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p CGLPixelFormatObj) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that CGLPixelFormatObj points to.
func (p CGLPixelFormatObj) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p CGLPixelFormatObj) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the CGLPixelFormatObj pointer.
func (p CGLPixelFormatObj) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// CGLContextObj is a pointer to a void element.
type CGLContextObj struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewCGLContextObj returns a CGLContextObj that points to addr in the application pool.
func NewCGLContextObj(addr memory.Pointer) CGLContextObj {
	return CGLContextObj{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p CGLContextObj) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that CGLContextObj points to.
func (p CGLContextObj) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p CGLContextObj) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the CGLContextObj pointer.
func (p CGLContextObj) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// ImageOES is a pointer to a void element.
type ImageOES struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewImageOES returns a ImageOES that points to addr in the application pool.
func NewImageOES(addr memory.Pointer) ImageOES {
	return ImageOES{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p ImageOES) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that ImageOES points to.
func (p ImageOES) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p ImageOES) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the ImageOES pointer.
func (p ImageOES) String() string {
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

// EGLintᵖ is a pointer to a EGLint element.
type EGLintᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewEGLintᵖ returns a EGLintᵖ that points to addr in the application pool.
func NewEGLintᵖ(addr memory.Pointer) EGLintᵖ {
	return EGLintᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p EGLintᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that EGLintᵖ points to.
func (p EGLintᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(ϟs.Architecture.IntegerSize)
}

// Read reads and returns the EGLint element at the pointer.
func (p EGLintᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) EGLint {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the EGLint element at the pointer.
func (p EGLintᵖ) Write(value EGLint, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]EGLint{value}, ϟs)
}

// Slice returns a new EGLintˢ from the pointer using start and end indices.
func (p EGLintᵖ) Slice(start, end uint64, ϟs *gfxapi.State) EGLintˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return EGLintˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the EGLintᵖ pointer.
func (p EGLintᵖ) String() string {
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

// CGLContextObjᵖ is a pointer to a CGLContextObj element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type CGLContextObjᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewCGLContextObjᵖ returns a CGLContextObjᵖ that points to addr in the application pool.
func NewCGLContextObjᵖ(addr memory.Pointer) CGLContextObjᵖ {
	return CGLContextObjᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p CGLContextObjᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that CGLContextObjᵖ points to.
func (p CGLContextObjᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Read reads and returns the CGLContextObj element at the pointer.
func (p CGLContextObjᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGLContextObj {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the CGLContextObj element at the pointer.
func (p CGLContextObjᵖ) Write(value CGLContextObj, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]CGLContextObj{value}, ϟs)
}

// Slice returns a new CGLContextObjˢ from the pointer using start and end indices.
func (p CGLContextObjᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGLContextObjˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return CGLContextObjˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the CGLContextObjᵖ pointer.
func (p CGLContextObjᵖ) String() string {
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

// DiscardFramebufferAttachmentᵖ is a pointer to a DiscardFramebufferAttachment element.
type DiscardFramebufferAttachmentᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewDiscardFramebufferAttachmentᵖ returns a DiscardFramebufferAttachmentᵖ that points to addr in the application pool.
func NewDiscardFramebufferAttachmentᵖ(addr memory.Pointer) DiscardFramebufferAttachmentᵖ {
	return DiscardFramebufferAttachmentᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p DiscardFramebufferAttachmentᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that DiscardFramebufferAttachmentᵖ points to.
func (p DiscardFramebufferAttachmentᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the DiscardFramebufferAttachment element at the pointer.
func (p DiscardFramebufferAttachmentᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) DiscardFramebufferAttachment {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the DiscardFramebufferAttachment element at the pointer.
func (p DiscardFramebufferAttachmentᵖ) Write(value DiscardFramebufferAttachment, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]DiscardFramebufferAttachment{value}, ϟs)
}

// Slice returns a new DiscardFramebufferAttachmentˢ from the pointer using start and end indices.
func (p DiscardFramebufferAttachmentᵖ) Slice(start, end uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return DiscardFramebufferAttachmentˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the DiscardFramebufferAttachmentᵖ pointer.
func (p DiscardFramebufferAttachmentᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// VertexArrayIdᵖ is a pointer to a VertexArrayId element.
type VertexArrayIdᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewVertexArrayIdᵖ returns a VertexArrayIdᵖ that points to addr in the application pool.
func NewVertexArrayIdᵖ(addr memory.Pointer) VertexArrayIdᵖ {
	return VertexArrayIdᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p VertexArrayIdᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that VertexArrayIdᵖ points to.
func (p VertexArrayIdᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the VertexArrayId element at the pointer.
func (p VertexArrayIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) VertexArrayId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the VertexArrayId element at the pointer.
func (p VertexArrayIdᵖ) Write(value VertexArrayId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]VertexArrayId{value}, ϟs)
}

// Slice returns a new VertexArrayIdˢ from the pointer using start and end indices.
func (p VertexArrayIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return VertexArrayIdˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the VertexArrayIdᵖ pointer.
func (p VertexArrayIdᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// ShaderAttribTypeᵖ is a pointer to a ShaderAttribType element.
type ShaderAttribTypeᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewShaderAttribTypeᵖ returns a ShaderAttribTypeᵖ that points to addr in the application pool.
func NewShaderAttribTypeᵖ(addr memory.Pointer) ShaderAttribTypeᵖ {
	return ShaderAttribTypeᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p ShaderAttribTypeᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that ShaderAttribTypeᵖ points to.
func (p ShaderAttribTypeᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the ShaderAttribType element at the pointer.
func (p ShaderAttribTypeᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderAttribType {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the ShaderAttribType element at the pointer.
func (p ShaderAttribTypeᵖ) Write(value ShaderAttribType, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]ShaderAttribType{value}, ϟs)
}

// Slice returns a new ShaderAttribTypeˢ from the pointer using start and end indices.
func (p ShaderAttribTypeᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderAttribTypeˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return ShaderAttribTypeˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the ShaderAttribTypeᵖ pointer.
func (p ShaderAttribTypeᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// ShaderUniformTypeᵖ is a pointer to a ShaderUniformType element.
type ShaderUniformTypeᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewShaderUniformTypeᵖ returns a ShaderUniformTypeᵖ that points to addr in the application pool.
func NewShaderUniformTypeᵖ(addr memory.Pointer) ShaderUniformTypeᵖ {
	return ShaderUniformTypeᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p ShaderUniformTypeᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that ShaderUniformTypeᵖ points to.
func (p ShaderUniformTypeᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the ShaderUniformType element at the pointer.
func (p ShaderUniformTypeᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderUniformType {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the ShaderUniformType element at the pointer.
func (p ShaderUniformTypeᵖ) Write(value ShaderUniformType, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]ShaderUniformType{value}, ϟs)
}

// Slice returns a new ShaderUniformTypeˢ from the pointer using start and end indices.
func (p ShaderUniformTypeᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderUniformTypeˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return ShaderUniformTypeˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the ShaderUniformTypeᵖ pointer.
func (p ShaderUniformTypeᵖ) String() string {
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

// TextureIdᵖ is a pointer to a TextureId element.
type TextureIdᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewTextureIdᵖ returns a TextureIdᵖ that points to addr in the application pool.
func NewTextureIdᵖ(addr memory.Pointer) TextureIdᵖ {
	return TextureIdᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p TextureIdᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that TextureIdᵖ points to.
func (p TextureIdᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the TextureId element at the pointer.
func (p TextureIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) TextureId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the TextureId element at the pointer.
func (p TextureIdᵖ) Write(value TextureId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]TextureId{value}, ϟs)
}

// Slice returns a new TextureIdˢ from the pointer using start and end indices.
func (p TextureIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return TextureIdˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the TextureIdᵖ pointer.
func (p TextureIdᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// FramebufferIdᵖ is a pointer to a FramebufferId element.
type FramebufferIdᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewFramebufferIdᵖ returns a FramebufferIdᵖ that points to addr in the application pool.
func NewFramebufferIdᵖ(addr memory.Pointer) FramebufferIdᵖ {
	return FramebufferIdᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p FramebufferIdᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that FramebufferIdᵖ points to.
func (p FramebufferIdᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the FramebufferId element at the pointer.
func (p FramebufferIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) FramebufferId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the FramebufferId element at the pointer.
func (p FramebufferIdᵖ) Write(value FramebufferId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]FramebufferId{value}, ϟs)
}

// Slice returns a new FramebufferIdˢ from the pointer using start and end indices.
func (p FramebufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return FramebufferIdˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the FramebufferIdᵖ pointer.
func (p FramebufferIdᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// RenderbufferIdᵖ is a pointer to a RenderbufferId element.
type RenderbufferIdᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewRenderbufferIdᵖ returns a RenderbufferIdᵖ that points to addr in the application pool.
func NewRenderbufferIdᵖ(addr memory.Pointer) RenderbufferIdᵖ {
	return RenderbufferIdᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p RenderbufferIdᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that RenderbufferIdᵖ points to.
func (p RenderbufferIdᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the RenderbufferId element at the pointer.
func (p RenderbufferIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) RenderbufferId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the RenderbufferId element at the pointer.
func (p RenderbufferIdᵖ) Write(value RenderbufferId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]RenderbufferId{value}, ϟs)
}

// Slice returns a new RenderbufferIdˢ from the pointer using start and end indices.
func (p RenderbufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return RenderbufferIdˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the RenderbufferIdᵖ pointer.
func (p RenderbufferIdᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// BufferIdᵖ is a pointer to a BufferId element.
type BufferIdᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewBufferIdᵖ returns a BufferIdᵖ that points to addr in the application pool.
func NewBufferIdᵖ(addr memory.Pointer) BufferIdᵖ {
	return BufferIdᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p BufferIdᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that BufferIdᵖ points to.
func (p BufferIdᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the BufferId element at the pointer.
func (p BufferIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) BufferId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the BufferId element at the pointer.
func (p BufferIdᵖ) Write(value BufferId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]BufferId{value}, ϟs)
}

// Slice returns a new BufferIdˢ from the pointer using start and end indices.
func (p BufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return BufferIdˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the BufferIdᵖ pointer.
func (p BufferIdᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// Charᵖᵖ is a pointer to a Charᵖ element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type Charᵖᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewCharᵖᵖ returns a Charᵖᵖ that points to addr in the application pool.
func NewCharᵖᵖ(addr memory.Pointer) Charᵖᵖ {
	return Charᵖᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p Charᵖᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that Charᵖᵖ points to.
func (p Charᵖᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Read reads and returns the Charᵖ element at the pointer.
func (p Charᵖᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Charᵖ {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Charᵖ element at the pointer.
func (p Charᵖᵖ) Write(value Charᵖ, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Charᵖ{value}, ϟs)
}

// Slice returns a new Charᵖˢ from the pointer using start and end indices.
func (p Charᵖᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charᵖˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charᵖˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the Charᵖᵖ pointer.
func (p Charᵖᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// ShaderIdᵖ is a pointer to a ShaderId element.
type ShaderIdᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewShaderIdᵖ returns a ShaderIdᵖ that points to addr in the application pool.
func NewShaderIdᵖ(addr memory.Pointer) ShaderIdᵖ {
	return ShaderIdᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p ShaderIdᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that ShaderIdᵖ points to.
func (p ShaderIdᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the ShaderId element at the pointer.
func (p ShaderIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the ShaderId element at the pointer.
func (p ShaderIdᵖ) Write(value ShaderId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]ShaderId{value}, ϟs)
}

// Slice returns a new ShaderIdˢ from the pointer using start and end indices.
func (p ShaderIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return ShaderIdˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the ShaderIdᵖ pointer.
func (p ShaderIdᵖ) String() string {
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

// FramebufferAttachmentᵖ is a pointer to a FramebufferAttachment element.
type FramebufferAttachmentᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewFramebufferAttachmentᵖ returns a FramebufferAttachmentᵖ that points to addr in the application pool.
func NewFramebufferAttachmentᵖ(addr memory.Pointer) FramebufferAttachmentᵖ {
	return FramebufferAttachmentᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p FramebufferAttachmentᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that FramebufferAttachmentᵖ points to.
func (p FramebufferAttachmentᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the FramebufferAttachment element at the pointer.
func (p FramebufferAttachmentᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) FramebufferAttachment {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the FramebufferAttachment element at the pointer.
func (p FramebufferAttachmentᵖ) Write(value FramebufferAttachment, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]FramebufferAttachment{value}, ϟs)
}

// Slice returns a new FramebufferAttachmentˢ from the pointer using start and end indices.
func (p FramebufferAttachmentᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferAttachmentˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return FramebufferAttachmentˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the FramebufferAttachmentᵖ pointer.
func (p FramebufferAttachmentᵖ) String() string {
	return fmt.Sprintf("%v@%v", p.Address, p.Pool)
}

// QueryIdᵖ is a pointer to a QueryId element.
type QueryIdᵖ struct {
	binary.Generate
	Address memory.Pointer
	Pool    memory.PoolID
}

// NewQueryIdᵖ returns a QueryIdᵖ that points to addr in the application pool.
func NewQueryIdᵖ(addr memory.Pointer) QueryIdᵖ {
	return QueryIdᵖ{Address: addr, Pool: memory.ApplicationPool}
}

// Unbounded returns a new unbounded memory.Slice based at p.
func (p QueryIdᵖ) Unbounded(ϟs *gfxapi.State) memory.Slice {
	return ϟs.Memory[p.Pool].At(p.Address)
}

// ElementSize returns the size in bytes of an element that QueryIdᵖ points to.
func (p QueryIdᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the QueryId element at the pointer.
func (p QueryIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) QueryId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the QueryId element at the pointer.
func (p QueryIdᵖ) Write(value QueryId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]QueryId{value}, ϟs)
}

// Slice returns a new QueryIdˢ from the pointer using start and end indices.
func (p QueryIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return QueryIdˢ{Root: p.Address, Base: p.Address.Offset(start * p.ElementSize(ϟs)), Count: end - start, Pool: p.Pool}
}

// String returns a string description of the QueryIdᵖ pointer.
func (p QueryIdᵖ) String() string {
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

type BufferIdˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeBufferIdˢ returns a BufferIdˢ backed by a new memory pool.
func MakeBufferIdˢ(count uint64, ϟs *gfxapi.State) BufferIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return BufferIdˢ{Count: count, Pool: id}
}

// Clone returns a copy of the BufferIdˢ in a new memory pool.
func (s BufferIdˢ) Clone(ϟs *gfxapi.State) BufferIdˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := BufferIdˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that BufferIdˢ points to.
func (s BufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s BufferIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s BufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s BufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s BufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the BufferId elements in this BufferIdˢ.
func (s BufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []BufferId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]BufferId, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = BufferId(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst BufferIdˢ) Write(src []BufferId, ϟs *gfxapi.State) uint64 {
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
func (dst BufferIdˢ) Copy(src BufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s BufferIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a BufferIdᵖ to the i'th element in this BufferIdˢ.
func (s BufferIdˢ) Index(i uint64, ϟs *gfxapi.State) BufferIdᵖ {
	return BufferIdᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the BufferIdˢ using start and end indices.
func (s BufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return BufferIdˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the BufferIdˢ slice.
func (s BufferIdˢ) String() string {
	return fmt.Sprintf("BufferId(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type CGLContextObjˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeCGLContextObjˢ returns a CGLContextObjˢ backed by a new memory pool.
func MakeCGLContextObjˢ(count uint64, ϟs *gfxapi.State) CGLContextObjˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return CGLContextObjˢ{Count: count, Pool: id}
}

// Clone returns a copy of the CGLContextObjˢ in a new memory pool.
func (s CGLContextObjˢ) Clone(ϟs *gfxapi.State) CGLContextObjˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := CGLContextObjˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that CGLContextObjˢ points to.
func (s CGLContextObjˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	if s.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Range returns the memory range this slice represents in the underlying pool.
func (s CGLContextObjˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s CGLContextObjˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s CGLContextObjˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s CGLContextObjˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the CGLContextObj elements in this CGLContextObjˢ.
func (s CGLContextObjˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGLContextObj {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]CGLContextObj, s.Count)
	for i := range res {
		if s.Pool == memory.ApplicationPool {
			ptr, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			res[i] = NewCGLContextObj(memory.Pointer(ptr))
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
func (dst CGLContextObjˢ) Write(src []CGLContextObj, ϟs *gfxapi.State) uint64 {
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
func (dst CGLContextObjˢ) Copy(src CGLContextObjˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGLContextObjˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	if (dst.Pool == memory.ApplicationPool) != (src.Pool == memory.ApplicationPool) {
		dst.Write(src.Read(ϟs, ϟd, ϟl), ϟs)
	} else {
		ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	}
	return dst, src
}

// Index returns a CGLContextObjᵖ to the i'th element in this CGLContextObjˢ.
func (s CGLContextObjˢ) Index(i uint64, ϟs *gfxapi.State) CGLContextObjᵖ {
	return CGLContextObjᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the CGLContextObjˢ using start and end indices.
func (s CGLContextObjˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGLContextObjˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return CGLContextObjˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the CGLContextObjˢ slice.
func (s CGLContextObjˢ) String() string {
	return fmt.Sprintf("CGLContextObj(%v@%v)[%d]", s.Base, s.Pool, s.Count)
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

type Charᵖˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeCharᵖˢ returns a Charᵖˢ backed by a new memory pool.
func MakeCharᵖˢ(count uint64, ϟs *gfxapi.State) Charᵖˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Charᵖˢ{Count: count, Pool: id}
}

// Clone returns a copy of the Charᵖˢ in a new memory pool.
func (s Charᵖˢ) Clone(ϟs *gfxapi.State) Charᵖˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Charᵖˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that Charᵖˢ points to.
func (s Charᵖˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	if s.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Charᵖˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Charᵖˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Charᵖˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Charᵖˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the Charᵖ elements in this Charᵖˢ.
func (s Charᵖˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Charᵖ {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Charᵖ, s.Count)
	for i := range res {
		if s.Pool == memory.ApplicationPool {
			ptr, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			res[i] = NewCharᵖ(memory.Pointer(ptr))
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
func (dst Charᵖˢ) Write(src []Charᵖ, ϟs *gfxapi.State) uint64 {
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
func (dst Charᵖˢ) Copy(src Charᵖˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charᵖˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	if (dst.Pool == memory.ApplicationPool) != (src.Pool == memory.ApplicationPool) {
		dst.Write(src.Read(ϟs, ϟd, ϟl), ϟs)
	} else {
		ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	}
	return dst, src
}

// Index returns a Charᵖᵖ to the i'th element in this Charᵖˢ.
func (s Charᵖˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖᵖ {
	return Charᵖᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the Charᵖˢ using start and end indices.
func (s Charᵖˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charᵖˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Charᵖˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the Charᵖˢ slice.
func (s Charᵖˢ) String() string {
	return fmt.Sprintf("Charᵖ(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type DiscardFramebufferAttachmentˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeDiscardFramebufferAttachmentˢ returns a DiscardFramebufferAttachmentˢ backed by a new memory pool.
func MakeDiscardFramebufferAttachmentˢ(count uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return DiscardFramebufferAttachmentˢ{Count: count, Pool: id}
}

// Clone returns a copy of the DiscardFramebufferAttachmentˢ in a new memory pool.
func (s DiscardFramebufferAttachmentˢ) Clone(ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := DiscardFramebufferAttachmentˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that DiscardFramebufferAttachmentˢ points to.
func (s DiscardFramebufferAttachmentˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s DiscardFramebufferAttachmentˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s DiscardFramebufferAttachmentˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s DiscardFramebufferAttachmentˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s DiscardFramebufferAttachmentˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the DiscardFramebufferAttachment elements in this DiscardFramebufferAttachmentˢ.
func (s DiscardFramebufferAttachmentˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []DiscardFramebufferAttachment {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]DiscardFramebufferAttachment, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = DiscardFramebufferAttachment(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst DiscardFramebufferAttachmentˢ) Write(src []DiscardFramebufferAttachment, ϟs *gfxapi.State) uint64 {
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
func (dst DiscardFramebufferAttachmentˢ) Copy(src DiscardFramebufferAttachmentˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s DiscardFramebufferAttachmentˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a DiscardFramebufferAttachmentᵖ to the i'th element in this DiscardFramebufferAttachmentˢ.
func (s DiscardFramebufferAttachmentˢ) Index(i uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentᵖ {
	return DiscardFramebufferAttachmentᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the DiscardFramebufferAttachmentˢ using start and end indices.
func (s DiscardFramebufferAttachmentˢ) Slice(start, end uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return DiscardFramebufferAttachmentˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the DiscardFramebufferAttachmentˢ slice.
func (s DiscardFramebufferAttachmentˢ) String() string {
	return fmt.Sprintf("DiscardFramebufferAttachment(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type EGLintˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeEGLintˢ returns a EGLintˢ backed by a new memory pool.
func MakeEGLintˢ(count uint64, ϟs *gfxapi.State) EGLintˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return EGLintˢ{Count: count, Pool: id}
}

// Clone returns a copy of the EGLintˢ in a new memory pool.
func (s EGLintˢ) Clone(ϟs *gfxapi.State) EGLintˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := EGLintˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that EGLintˢ points to.
func (s EGLintˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(ϟs.Architecture.IntegerSize)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s EGLintˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s EGLintˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s EGLintˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s EGLintˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the EGLint elements in this EGLintˢ.
func (s EGLintˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []EGLint {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]EGLint, s.Count)
	for i := range res {
		if v, err := binary.ReadInt(d, ϟs.Architecture.IntegerSize*8); err == nil {
			res[i] = EGLint(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst EGLintˢ) Write(src []EGLint, ϟs *gfxapi.State) uint64 {
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
func (dst EGLintˢ) Copy(src EGLintˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s EGLintˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a EGLintᵖ to the i'th element in this EGLintˢ.
func (s EGLintˢ) Index(i uint64, ϟs *gfxapi.State) EGLintᵖ {
	return EGLintᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the EGLintˢ using start and end indices.
func (s EGLintˢ) Slice(start, end uint64, ϟs *gfxapi.State) EGLintˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return EGLintˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the EGLintˢ slice.
func (s EGLintˢ) String() string {
	return fmt.Sprintf("EGLint(%v@%v)[%d]", s.Base, s.Pool, s.Count)
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

type FramebufferAttachmentˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeFramebufferAttachmentˢ returns a FramebufferAttachmentˢ backed by a new memory pool.
func MakeFramebufferAttachmentˢ(count uint64, ϟs *gfxapi.State) FramebufferAttachmentˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return FramebufferAttachmentˢ{Count: count, Pool: id}
}

// Clone returns a copy of the FramebufferAttachmentˢ in a new memory pool.
func (s FramebufferAttachmentˢ) Clone(ϟs *gfxapi.State) FramebufferAttachmentˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := FramebufferAttachmentˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that FramebufferAttachmentˢ points to.
func (s FramebufferAttachmentˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s FramebufferAttachmentˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s FramebufferAttachmentˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s FramebufferAttachmentˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s FramebufferAttachmentˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the FramebufferAttachment elements in this FramebufferAttachmentˢ.
func (s FramebufferAttachmentˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []FramebufferAttachment {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]FramebufferAttachment, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = FramebufferAttachment(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst FramebufferAttachmentˢ) Write(src []FramebufferAttachment, ϟs *gfxapi.State) uint64 {
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
func (dst FramebufferAttachmentˢ) Copy(src FramebufferAttachmentˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s FramebufferAttachmentˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a FramebufferAttachmentᵖ to the i'th element in this FramebufferAttachmentˢ.
func (s FramebufferAttachmentˢ) Index(i uint64, ϟs *gfxapi.State) FramebufferAttachmentᵖ {
	return FramebufferAttachmentᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the FramebufferAttachmentˢ using start and end indices.
func (s FramebufferAttachmentˢ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferAttachmentˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return FramebufferAttachmentˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the FramebufferAttachmentˢ slice.
func (s FramebufferAttachmentˢ) String() string {
	return fmt.Sprintf("FramebufferAttachment(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type FramebufferIdˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeFramebufferIdˢ returns a FramebufferIdˢ backed by a new memory pool.
func MakeFramebufferIdˢ(count uint64, ϟs *gfxapi.State) FramebufferIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return FramebufferIdˢ{Count: count, Pool: id}
}

// Clone returns a copy of the FramebufferIdˢ in a new memory pool.
func (s FramebufferIdˢ) Clone(ϟs *gfxapi.State) FramebufferIdˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := FramebufferIdˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that FramebufferIdˢ points to.
func (s FramebufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s FramebufferIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s FramebufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s FramebufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s FramebufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the FramebufferId elements in this FramebufferIdˢ.
func (s FramebufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []FramebufferId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]FramebufferId, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = FramebufferId(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst FramebufferIdˢ) Write(src []FramebufferId, ϟs *gfxapi.State) uint64 {
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
func (dst FramebufferIdˢ) Copy(src FramebufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s FramebufferIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a FramebufferIdᵖ to the i'th element in this FramebufferIdˢ.
func (s FramebufferIdˢ) Index(i uint64, ϟs *gfxapi.State) FramebufferIdᵖ {
	return FramebufferIdᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the FramebufferIdˢ using start and end indices.
func (s FramebufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return FramebufferIdˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the FramebufferIdˢ slice.
func (s FramebufferIdˢ) String() string {
	return fmt.Sprintf("FramebufferId(%v@%v)[%d]", s.Base, s.Pool, s.Count)
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

type QueryIdˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeQueryIdˢ returns a QueryIdˢ backed by a new memory pool.
func MakeQueryIdˢ(count uint64, ϟs *gfxapi.State) QueryIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return QueryIdˢ{Count: count, Pool: id}
}

// Clone returns a copy of the QueryIdˢ in a new memory pool.
func (s QueryIdˢ) Clone(ϟs *gfxapi.State) QueryIdˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := QueryIdˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that QueryIdˢ points to.
func (s QueryIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s QueryIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s QueryIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s QueryIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s QueryIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the QueryId elements in this QueryIdˢ.
func (s QueryIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []QueryId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]QueryId, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = QueryId(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst QueryIdˢ) Write(src []QueryId, ϟs *gfxapi.State) uint64 {
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
func (dst QueryIdˢ) Copy(src QueryIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s QueryIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a QueryIdᵖ to the i'th element in this QueryIdˢ.
func (s QueryIdˢ) Index(i uint64, ϟs *gfxapi.State) QueryIdᵖ {
	return QueryIdᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the QueryIdˢ using start and end indices.
func (s QueryIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return QueryIdˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the QueryIdˢ slice.
func (s QueryIdˢ) String() string {
	return fmt.Sprintf("QueryId(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type RenderbufferIdˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeRenderbufferIdˢ returns a RenderbufferIdˢ backed by a new memory pool.
func MakeRenderbufferIdˢ(count uint64, ϟs *gfxapi.State) RenderbufferIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return RenderbufferIdˢ{Count: count, Pool: id}
}

// Clone returns a copy of the RenderbufferIdˢ in a new memory pool.
func (s RenderbufferIdˢ) Clone(ϟs *gfxapi.State) RenderbufferIdˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := RenderbufferIdˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that RenderbufferIdˢ points to.
func (s RenderbufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s RenderbufferIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s RenderbufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s RenderbufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s RenderbufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the RenderbufferId elements in this RenderbufferIdˢ.
func (s RenderbufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []RenderbufferId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]RenderbufferId, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = RenderbufferId(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst RenderbufferIdˢ) Write(src []RenderbufferId, ϟs *gfxapi.State) uint64 {
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
func (dst RenderbufferIdˢ) Copy(src RenderbufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s RenderbufferIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a RenderbufferIdᵖ to the i'th element in this RenderbufferIdˢ.
func (s RenderbufferIdˢ) Index(i uint64, ϟs *gfxapi.State) RenderbufferIdᵖ {
	return RenderbufferIdᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the RenderbufferIdˢ using start and end indices.
func (s RenderbufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return RenderbufferIdˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the RenderbufferIdˢ slice.
func (s RenderbufferIdˢ) String() string {
	return fmt.Sprintf("RenderbufferId(%v@%v)[%d]", s.Base, s.Pool, s.Count)
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

type ShaderAttribTypeˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeShaderAttribTypeˢ returns a ShaderAttribTypeˢ backed by a new memory pool.
func MakeShaderAttribTypeˢ(count uint64, ϟs *gfxapi.State) ShaderAttribTypeˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return ShaderAttribTypeˢ{Count: count, Pool: id}
}

// Clone returns a copy of the ShaderAttribTypeˢ in a new memory pool.
func (s ShaderAttribTypeˢ) Clone(ϟs *gfxapi.State) ShaderAttribTypeˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := ShaderAttribTypeˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that ShaderAttribTypeˢ points to.
func (s ShaderAttribTypeˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s ShaderAttribTypeˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s ShaderAttribTypeˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s ShaderAttribTypeˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s ShaderAttribTypeˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the ShaderAttribType elements in this ShaderAttribTypeˢ.
func (s ShaderAttribTypeˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []ShaderAttribType {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]ShaderAttribType, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = ShaderAttribType(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst ShaderAttribTypeˢ) Write(src []ShaderAttribType, ϟs *gfxapi.State) uint64 {
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
func (dst ShaderAttribTypeˢ) Copy(src ShaderAttribTypeˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s ShaderAttribTypeˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a ShaderAttribTypeᵖ to the i'th element in this ShaderAttribTypeˢ.
func (s ShaderAttribTypeˢ) Index(i uint64, ϟs *gfxapi.State) ShaderAttribTypeᵖ {
	return ShaderAttribTypeᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the ShaderAttribTypeˢ using start and end indices.
func (s ShaderAttribTypeˢ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderAttribTypeˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return ShaderAttribTypeˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the ShaderAttribTypeˢ slice.
func (s ShaderAttribTypeˢ) String() string {
	return fmt.Sprintf("ShaderAttribType(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type ShaderIdˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeShaderIdˢ returns a ShaderIdˢ backed by a new memory pool.
func MakeShaderIdˢ(count uint64, ϟs *gfxapi.State) ShaderIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return ShaderIdˢ{Count: count, Pool: id}
}

// Clone returns a copy of the ShaderIdˢ in a new memory pool.
func (s ShaderIdˢ) Clone(ϟs *gfxapi.State) ShaderIdˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := ShaderIdˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that ShaderIdˢ points to.
func (s ShaderIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s ShaderIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s ShaderIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s ShaderIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s ShaderIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the ShaderId elements in this ShaderIdˢ.
func (s ShaderIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []ShaderId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]ShaderId, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = ShaderId(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst ShaderIdˢ) Write(src []ShaderId, ϟs *gfxapi.State) uint64 {
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
func (dst ShaderIdˢ) Copy(src ShaderIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s ShaderIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a ShaderIdᵖ to the i'th element in this ShaderIdˢ.
func (s ShaderIdˢ) Index(i uint64, ϟs *gfxapi.State) ShaderIdᵖ {
	return ShaderIdᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the ShaderIdˢ using start and end indices.
func (s ShaderIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return ShaderIdˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the ShaderIdˢ slice.
func (s ShaderIdˢ) String() string {
	return fmt.Sprintf("ShaderId(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type ShaderUniformTypeˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeShaderUniformTypeˢ returns a ShaderUniformTypeˢ backed by a new memory pool.
func MakeShaderUniformTypeˢ(count uint64, ϟs *gfxapi.State) ShaderUniformTypeˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return ShaderUniformTypeˢ{Count: count, Pool: id}
}

// Clone returns a copy of the ShaderUniformTypeˢ in a new memory pool.
func (s ShaderUniformTypeˢ) Clone(ϟs *gfxapi.State) ShaderUniformTypeˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := ShaderUniformTypeˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that ShaderUniformTypeˢ points to.
func (s ShaderUniformTypeˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s ShaderUniformTypeˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s ShaderUniformTypeˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s ShaderUniformTypeˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s ShaderUniformTypeˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the ShaderUniformType elements in this ShaderUniformTypeˢ.
func (s ShaderUniformTypeˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []ShaderUniformType {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]ShaderUniformType, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = ShaderUniformType(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst ShaderUniformTypeˢ) Write(src []ShaderUniformType, ϟs *gfxapi.State) uint64 {
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
func (dst ShaderUniformTypeˢ) Copy(src ShaderUniformTypeˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s ShaderUniformTypeˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a ShaderUniformTypeᵖ to the i'th element in this ShaderUniformTypeˢ.
func (s ShaderUniformTypeˢ) Index(i uint64, ϟs *gfxapi.State) ShaderUniformTypeᵖ {
	return ShaderUniformTypeᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the ShaderUniformTypeˢ using start and end indices.
func (s ShaderUniformTypeˢ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderUniformTypeˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return ShaderUniformTypeˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the ShaderUniformTypeˢ slice.
func (s ShaderUniformTypeˢ) String() string {
	return fmt.Sprintf("ShaderUniformType(%v@%v)[%d]", s.Base, s.Pool, s.Count)
}

type TextureIdˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeTextureIdˢ returns a TextureIdˢ backed by a new memory pool.
func MakeTextureIdˢ(count uint64, ϟs *gfxapi.State) TextureIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return TextureIdˢ{Count: count, Pool: id}
}

// Clone returns a copy of the TextureIdˢ in a new memory pool.
func (s TextureIdˢ) Clone(ϟs *gfxapi.State) TextureIdˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := TextureIdˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that TextureIdˢ points to.
func (s TextureIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s TextureIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s TextureIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s TextureIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s TextureIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the TextureId elements in this TextureIdˢ.
func (s TextureIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []TextureId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]TextureId, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = TextureId(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst TextureIdˢ) Write(src []TextureId, ϟs *gfxapi.State) uint64 {
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
func (dst TextureIdˢ) Copy(src TextureIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s TextureIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a TextureIdᵖ to the i'th element in this TextureIdˢ.
func (s TextureIdˢ) Index(i uint64, ϟs *gfxapi.State) TextureIdᵖ {
	return TextureIdᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the TextureIdˢ using start and end indices.
func (s TextureIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return TextureIdˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the TextureIdˢ slice.
func (s TextureIdˢ) String() string {
	return fmt.Sprintf("TextureId(%v@%v)[%d]", s.Base, s.Pool, s.Count)
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

type VertexArrayIdˢ struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  memory.Pointer // Address of first element.
	Count uint64         // Number of elements in the slice.
	Pool  memory.PoolID  // Pool that holds the element data.
}

// MakeVertexArrayIdˢ returns a VertexArrayIdˢ backed by a new memory pool.
func MakeVertexArrayIdˢ(count uint64, ϟs *gfxapi.State) VertexArrayIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return VertexArrayIdˢ{Count: count, Pool: id}
}

// Clone returns a copy of the VertexArrayIdˢ in a new memory pool.
func (s VertexArrayIdˢ) Clone(ϟs *gfxapi.State) VertexArrayIdˢ {
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := VertexArrayIdˢ{Count: s.Count, Pool: id}
	return dst
}

// ElementSize returns the size in bytes of an element that VertexArrayIdˢ points to.
func (s VertexArrayIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s VertexArrayIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return s.Base.Range(s.Count * s.ElementSize(ϟs))
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s VertexArrayIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s VertexArrayIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s VertexArrayIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Pool], s.Range(ϟs))
}

// Read reads and returns all the VertexArrayId elements in this VertexArrayIdˢ.
func (s VertexArrayIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []VertexArrayId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]VertexArrayId, s.Count)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = VertexArrayId(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of dst.Count and src.Count.
func (dst VertexArrayIdˢ) Write(src []VertexArrayId, ϟs *gfxapi.State) uint64 {
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
func (dst VertexArrayIdˢ) Copy(src VertexArrayIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s VertexArrayIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	ϟs.Memory[dst.Pool].Write(dst.Base, ϟs.Memory[src.Pool].Slice(src.Range(ϟs)))
	return dst, src
}

// Index returns a VertexArrayIdᵖ to the i'th element in this VertexArrayIdˢ.
func (s VertexArrayIdˢ) Index(i uint64, ϟs *gfxapi.State) VertexArrayIdᵖ {
	return VertexArrayIdᵖ{Address: s.Base.Offset(i * s.ElementSize(ϟs)), Pool: s.Pool}
}

// Slice returns a sub-slice from the VertexArrayIdˢ using start and end indices.
func (s VertexArrayIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return VertexArrayIdˢ{Root: s.Root, Base: s.Base.Offset(start * s.ElementSize(ϟs)), Count: end - start, Pool: s.Pool}
}

// String returns a string description of the VertexArrayIdˢ slice.
func (s VertexArrayIdˢ) String() string {
	return fmt.Sprintf("VertexArrayId(%v@%v)[%d]", s.Base, s.Pool, s.Count)
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

type AttributeLocationːVertexAttributeArrayʳᵐ map[AttributeLocation](*VertexAttributeArray)

func (m AttributeLocationːVertexAttributeArrayʳᵐ) Get(key AttributeLocation) *VertexAttributeArray {
	return m[key]
}
func (m AttributeLocationːVertexAttributeArrayʳᵐ) Contains(key AttributeLocation) bool {
	_, ok := m[key]
	return ok
}
func (m AttributeLocationːVertexAttributeArrayʳᵐ) Delete(key AttributeLocation) {
	delete(m, key)
}
func (m AttributeLocationːVertexAttributeArrayʳᵐ) Range() [](*VertexAttributeArray) {
	values := make([](*VertexAttributeArray), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type BufferIdːBufferʳᵐ map[BufferId](*Buffer)

func (m BufferIdːBufferʳᵐ) Get(key BufferId) *Buffer {
	return m[key]
}
func (m BufferIdːBufferʳᵐ) Contains(key BufferId) bool {
	_, ok := m[key]
	return ok
}
func (m BufferIdːBufferʳᵐ) Delete(key BufferId) {
	delete(m, key)
}
func (m BufferIdːBufferʳᵐ) Range() [](*Buffer) {
	values := make([](*Buffer), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type BufferTargetːBufferIdᵐ map[BufferTarget]BufferId

func (m BufferTargetːBufferIdᵐ) Get(key BufferTarget) BufferId {
	return m[key]
}
func (m BufferTargetːBufferIdᵐ) Contains(key BufferTarget) bool {
	_, ok := m[key]
	return ok
}
func (m BufferTargetːBufferIdᵐ) Delete(key BufferTarget) {
	delete(m, key)
}
func (m BufferTargetːBufferIdᵐ) Range() []BufferId {
	values := make([]BufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type CGLContextObjːContextʳᵐ map[CGLContextObj](*Context)

func (m CGLContextObjːContextʳᵐ) Get(key CGLContextObj) *Context {
	return m[key]
}
func (m CGLContextObjːContextʳᵐ) Contains(key CGLContextObj) bool {
	_, ok := m[key]
	return ok
}
func (m CGLContextObjːContextʳᵐ) Delete(key CGLContextObj) {
	delete(m, key)
}
func (m CGLContextObjːContextʳᵐ) Range() [](*Context) {
	values := make([](*Context), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type Capabilityːboolᵐ map[Capability]bool

func (m Capabilityːboolᵐ) Get(key Capability) bool {
	return m[key]
}
func (m Capabilityːboolᵐ) Contains(key Capability) bool {
	_, ok := m[key]
	return ok
}
func (m Capabilityːboolᵐ) Delete(key Capability) {
	delete(m, key)
}
func (m Capabilityːboolᵐ) Range() []bool {
	values := make([]bool, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type CubeMapImageTargetːImageᵐ map[CubeMapImageTarget]Image

func (m CubeMapImageTargetːImageᵐ) Get(key CubeMapImageTarget) Image {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m CubeMapImageTargetːImageᵐ) Contains(key CubeMapImageTarget) bool {
	_, ok := m[key]
	return ok
}
func (m CubeMapImageTargetːImageᵐ) Delete(key CubeMapImageTarget) {
	delete(m, key)
}
func (m CubeMapImageTargetːImageᵐ) Range() []Image {
	values := make([]Image, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type EGLContextːContextʳᵐ map[EGLContext](*Context)

func (m EGLContextːContextʳᵐ) Get(key EGLContext) *Context {
	return m[key]
}
func (m EGLContextːContextʳᵐ) Contains(key EGLContext) bool {
	_, ok := m[key]
	return ok
}
func (m EGLContextːContextʳᵐ) Delete(key EGLContext) {
	delete(m, key)
}
func (m EGLContextːContextʳᵐ) Range() [](*Context) {
	values := make([](*Context), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type FaceModeːu32ᵐ map[FaceMode]uint32

func (m FaceModeːu32ᵐ) Get(key FaceMode) uint32 {
	return m[key]
}
func (m FaceModeːu32ᵐ) Contains(key FaceMode) bool {
	_, ok := m[key]
	return ok
}
func (m FaceModeːu32ᵐ) Delete(key FaceMode) {
	delete(m, key)
}
func (m FaceModeːu32ᵐ) Range() []uint32 {
	values := make([]uint32, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type FramebufferAttachmentːFramebufferAttachmentInfoᵐ map[FramebufferAttachment]FramebufferAttachmentInfo

func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Get(key FramebufferAttachment) FramebufferAttachmentInfo {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Contains(key FramebufferAttachment) bool {
	_, ok := m[key]
	return ok
}
func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Delete(key FramebufferAttachment) {
	delete(m, key)
}
func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Range() []FramebufferAttachmentInfo {
	values := make([]FramebufferAttachmentInfo, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type FramebufferIdːFramebufferʳᵐ map[FramebufferId](*Framebuffer)

func (m FramebufferIdːFramebufferʳᵐ) Get(key FramebufferId) *Framebuffer {
	return m[key]
}
func (m FramebufferIdːFramebufferʳᵐ) Contains(key FramebufferId) bool {
	_, ok := m[key]
	return ok
}
func (m FramebufferIdːFramebufferʳᵐ) Delete(key FramebufferId) {
	delete(m, key)
}
func (m FramebufferIdːFramebufferʳᵐ) Range() [](*Framebuffer) {
	values := make([](*Framebuffer), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type FramebufferTargetːFramebufferIdᵐ map[FramebufferTarget]FramebufferId

func (m FramebufferTargetːFramebufferIdᵐ) Get(key FramebufferTarget) FramebufferId {
	return m[key]
}
func (m FramebufferTargetːFramebufferIdᵐ) Contains(key FramebufferTarget) bool {
	_, ok := m[key]
	return ok
}
func (m FramebufferTargetːFramebufferIdᵐ) Delete(key FramebufferTarget) {
	delete(m, key)
}
func (m FramebufferTargetːFramebufferIdᵐ) Range() []FramebufferId {
	values := make([]FramebufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLXContextːContextʳᵐ map[GLXContext](*Context)

func (m GLXContextːContextʳᵐ) Get(key GLXContext) *Context {
	return m[key]
}
func (m GLXContextːContextʳᵐ) Contains(key GLXContext) bool {
	_, ok := m[key]
	return ok
}
func (m GLXContextːContextʳᵐ) Delete(key GLXContext) {
	delete(m, key)
}
func (m GLXContextːContextʳᵐ) Range() [](*Context) {
	values := make([](*Context), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type HGLRCːContextʳᵐ map[HGLRC](*Context)

func (m HGLRCːContextʳᵐ) Get(key HGLRC) *Context {
	return m[key]
}
func (m HGLRCːContextʳᵐ) Contains(key HGLRC) bool {
	_, ok := m[key]
	return ok
}
func (m HGLRCːContextʳᵐ) Delete(key HGLRC) {
	delete(m, key)
}
func (m HGLRCːContextʳᵐ) Range() [](*Context) {
	values := make([](*Context), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type PixelStoreParameterːs32ᵐ map[PixelStoreParameter]int32

func (m PixelStoreParameterːs32ᵐ) Get(key PixelStoreParameter) int32 {
	return m[key]
}
func (m PixelStoreParameterːs32ᵐ) Contains(key PixelStoreParameter) bool {
	_, ok := m[key]
	return ok
}
func (m PixelStoreParameterːs32ᵐ) Delete(key PixelStoreParameter) {
	delete(m, key)
}
func (m PixelStoreParameterːs32ᵐ) Range() []int32 {
	values := make([]int32, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ProgramIdːProgramʳᵐ map[ProgramId](*Program)

func (m ProgramIdːProgramʳᵐ) Get(key ProgramId) *Program {
	return m[key]
}
func (m ProgramIdːProgramʳᵐ) Contains(key ProgramId) bool {
	_, ok := m[key]
	return ok
}
func (m ProgramIdːProgramʳᵐ) Delete(key ProgramId) {
	delete(m, key)
}
func (m ProgramIdːProgramʳᵐ) Range() [](*Program) {
	values := make([](*Program), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type QueryIdːQueryʳᵐ map[QueryId](*Query)

func (m QueryIdːQueryʳᵐ) Get(key QueryId) *Query {
	return m[key]
}
func (m QueryIdːQueryʳᵐ) Contains(key QueryId) bool {
	_, ok := m[key]
	return ok
}
func (m QueryIdːQueryʳᵐ) Delete(key QueryId) {
	delete(m, key)
}
func (m QueryIdːQueryʳᵐ) Range() [](*Query) {
	values := make([](*Query), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type RenderbufferIdːRenderbufferʳᵐ map[RenderbufferId](*Renderbuffer)

func (m RenderbufferIdːRenderbufferʳᵐ) Get(key RenderbufferId) *Renderbuffer {
	return m[key]
}
func (m RenderbufferIdːRenderbufferʳᵐ) Contains(key RenderbufferId) bool {
	_, ok := m[key]
	return ok
}
func (m RenderbufferIdːRenderbufferʳᵐ) Delete(key RenderbufferId) {
	delete(m, key)
}
func (m RenderbufferIdːRenderbufferʳᵐ) Range() [](*Renderbuffer) {
	values := make([](*Renderbuffer), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type RenderbufferTargetːRenderbufferIdᵐ map[RenderbufferTarget]RenderbufferId

func (m RenderbufferTargetːRenderbufferIdᵐ) Get(key RenderbufferTarget) RenderbufferId {
	return m[key]
}
func (m RenderbufferTargetːRenderbufferIdᵐ) Contains(key RenderbufferTarget) bool {
	_, ok := m[key]
	return ok
}
func (m RenderbufferTargetːRenderbufferIdᵐ) Delete(key RenderbufferTarget) {
	delete(m, key)
}
func (m RenderbufferTargetːRenderbufferIdᵐ) Range() []RenderbufferId {
	values := make([]RenderbufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type S32ːCubemapLevelᵐ map[int32]CubemapLevel

func (m S32ːCubemapLevelᵐ) Get(key int32) CubemapLevel {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m S32ːCubemapLevelᵐ) Contains(key int32) bool {
	_, ok := m[key]
	return ok
}
func (m S32ːCubemapLevelᵐ) Delete(key int32) {
	delete(m, key)
}
func (m S32ːCubemapLevelᵐ) Range() []CubemapLevel {
	values := make([]CubemapLevel, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type S32ːImageᵐ map[int32]Image

func (m S32ːImageᵐ) Get(key int32) Image {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m S32ːImageᵐ) Contains(key int32) bool {
	_, ok := m[key]
	return ok
}
func (m S32ːImageᵐ) Delete(key int32) {
	delete(m, key)
}
func (m S32ːImageᵐ) Range() []Image {
	values := make([]Image, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type S32ːVertexAttributeᵐ map[int32]VertexAttribute

func (m S32ːVertexAttributeᵐ) Get(key int32) VertexAttribute {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m S32ːVertexAttributeᵐ) Contains(key int32) bool {
	_, ok := m[key]
	return ok
}
func (m S32ːVertexAttributeᵐ) Delete(key int32) {
	delete(m, key)
}
func (m S32ːVertexAttributeᵐ) Range() []VertexAttribute {
	values := make([]VertexAttribute, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ShaderIdːShaderʳᵐ map[ShaderId](*Shader)

func (m ShaderIdːShaderʳᵐ) Get(key ShaderId) *Shader {
	return m[key]
}
func (m ShaderIdːShaderʳᵐ) Contains(key ShaderId) bool {
	_, ok := m[key]
	return ok
}
func (m ShaderIdːShaderʳᵐ) Delete(key ShaderId) {
	delete(m, key)
}
func (m ShaderIdːShaderʳᵐ) Range() [](*Shader) {
	values := make([](*Shader), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ShaderTypeːShaderIdᵐ map[ShaderType]ShaderId

func (m ShaderTypeːShaderIdᵐ) Get(key ShaderType) ShaderId {
	return m[key]
}
func (m ShaderTypeːShaderIdᵐ) Contains(key ShaderType) bool {
	_, ok := m[key]
	return ok
}
func (m ShaderTypeːShaderIdᵐ) Delete(key ShaderType) {
	delete(m, key)
}
func (m ShaderTypeːShaderIdᵐ) Range() []ShaderId {
	values := make([]ShaderId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type StringːAttributeLocationᵐ map[string]AttributeLocation

func (m StringːAttributeLocationᵐ) Get(key string) AttributeLocation {
	return m[key]
}
func (m StringːAttributeLocationᵐ) Contains(key string) bool {
	_, ok := m[key]
	return ok
}
func (m StringːAttributeLocationᵐ) Delete(key string) {
	delete(m, key)
}
func (m StringːAttributeLocationᵐ) Range() []AttributeLocation {
	values := make([]AttributeLocation, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type TextureIdːTextureʳᵐ map[TextureId](*Texture)

func (m TextureIdːTextureʳᵐ) Get(key TextureId) *Texture {
	return m[key]
}
func (m TextureIdːTextureʳᵐ) Contains(key TextureId) bool {
	_, ok := m[key]
	return ok
}
func (m TextureIdːTextureʳᵐ) Delete(key TextureId) {
	delete(m, key)
}
func (m TextureIdːTextureʳᵐ) Range() [](*Texture) {
	values := make([](*Texture), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type TextureTargetːTextureIdᵐ map[TextureTarget]TextureId

func (m TextureTargetːTextureIdᵐ) Get(key TextureTarget) TextureId {
	return m[key]
}
func (m TextureTargetːTextureIdᵐ) Contains(key TextureTarget) bool {
	_, ok := m[key]
	return ok
}
func (m TextureTargetːTextureIdᵐ) Delete(key TextureTarget) {
	delete(m, key)
}
func (m TextureTargetːTextureIdᵐ) Range() []TextureId {
	values := make([]TextureId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type TextureUnitːTextureTargetːTextureIdᵐᵐ map[TextureUnit]TextureTargetːTextureIdᵐ

func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Get(key TextureUnit) TextureTargetːTextureIdᵐ {
	v, ok := m[key]
	if !ok {
		v = make(TextureTargetːTextureIdᵐ)
	}
	return v
}
func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Contains(key TextureUnit) bool {
	_, ok := m[key]
	return ok
}
func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Delete(key TextureUnit) {
	delete(m, key)
}
func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Range() []TextureTargetːTextureIdᵐ {
	values := make([]TextureTargetːTextureIdᵐ, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ThreadIDːContextʳᵐ map[ThreadID](*Context)

func (m ThreadIDːContextʳᵐ) Get(key ThreadID) *Context {
	return m[key]
}
func (m ThreadIDːContextʳᵐ) Contains(key ThreadID) bool {
	_, ok := m[key]
	return ok
}
func (m ThreadIDːContextʳᵐ) Delete(key ThreadID) {
	delete(m, key)
}
func (m ThreadIDːContextʳᵐ) Range() [](*Context) {
	values := make([](*Context), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type UniformLocationːUniformᵐ map[UniformLocation]Uniform

func (m UniformLocationːUniformᵐ) Get(key UniformLocation) Uniform {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m UniformLocationːUniformᵐ) Contains(key UniformLocation) bool {
	_, ok := m[key]
	return ok
}
func (m UniformLocationːUniformᵐ) Delete(key UniformLocation) {
	delete(m, key)
}
func (m UniformLocationːUniformᵐ) Range() []Uniform {
	values := make([]Uniform, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type VertexArrayIdːVertexArrayʳᵐ map[VertexArrayId](*VertexArray)

func (m VertexArrayIdːVertexArrayʳᵐ) Get(key VertexArrayId) *VertexArray {
	return m[key]
}
func (m VertexArrayIdːVertexArrayʳᵐ) Contains(key VertexArrayId) bool {
	_, ok := m[key]
	return ok
}
func (m VertexArrayIdːVertexArrayʳᵐ) Delete(key VertexArrayId) {
	delete(m, key)
}
func (m VertexArrayIdːVertexArrayʳᵐ) Range() [](*VertexArray) {
	values := make([](*VertexArray), 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

////////////////////////////////////////////////////////////////////////////////
// ReplayCreateRenderer
////////////////////////////////////////////////////////////////////////////////
type ReplayCreateRenderer struct {
	binary.Generate
	observations atom.Observations
	Id           uint32
}

func (a *ReplayCreateRenderer) String() string {
	return fmt.Sprintf("replayCreateRenderer(id: %v)", a.Id)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The ReplayCreateRenderer pointer is returned so that calls can be chained.
func (a *ReplayCreateRenderer) AddRead(rng memory.Range, id binary.ID) *ReplayCreateRenderer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The ReplayCreateRenderer pointer is returned so that calls can be chained.
func (a *ReplayCreateRenderer) AddWrite(rng memory.Range, id binary.ID) *ReplayCreateRenderer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *ReplayCreateRenderer) API() gfxapi.API                  { return api{} }
func (c *ReplayCreateRenderer) TypeID() atom.TypeID              { return 0 }
func (c *ReplayCreateRenderer) Flags() atom.Flags                { return 0 }
func (a *ReplayCreateRenderer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// ReplayBindRenderer
////////////////////////////////////////////////////////////////////////////////
type ReplayBindRenderer struct {
	binary.Generate
	observations atom.Observations
	Id           uint32
}

func (a *ReplayBindRenderer) String() string {
	return fmt.Sprintf("replayBindRenderer(id: %v)", a.Id)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The ReplayBindRenderer pointer is returned so that calls can be chained.
func (a *ReplayBindRenderer) AddRead(rng memory.Range, id binary.ID) *ReplayBindRenderer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The ReplayBindRenderer pointer is returned so that calls can be chained.
func (a *ReplayBindRenderer) AddWrite(rng memory.Range, id binary.ID) *ReplayBindRenderer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *ReplayBindRenderer) API() gfxapi.API                  { return api{} }
func (c *ReplayBindRenderer) TypeID() atom.TypeID              { return 1 }
func (c *ReplayBindRenderer) Flags() atom.Flags                { return 0 }
func (a *ReplayBindRenderer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// BackbufferInfo
////////////////////////////////////////////////////////////////////////////////
type BackbufferInfo struct {
	binary.Generate
	observations         atom.Observations
	Width                int32
	Height               int32
	ColorFmt             RenderbufferFormat
	DepthFmt             RenderbufferFormat
	StencilFmt           RenderbufferFormat
	ResetViewportScissor bool
}

func (a *BackbufferInfo) String() string {
	return fmt.Sprintf("backbufferInfo(width: %v, height: %v, color_fmt: %v, depth_fmt: %v, stencil_fmt: %v, resetViewportScissor: %v)", a.Width, a.Height, a.ColorFmt, a.DepthFmt, a.StencilFmt, a.ResetViewportScissor)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The BackbufferInfo pointer is returned so that calls can be chained.
func (a *BackbufferInfo) AddRead(rng memory.Range, id binary.ID) *BackbufferInfo {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The BackbufferInfo pointer is returned so that calls can be chained.
func (a *BackbufferInfo) AddWrite(rng memory.Range, id binary.ID) *BackbufferInfo {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *BackbufferInfo) API() gfxapi.API                  { return api{} }
func (c *BackbufferInfo) TypeID() atom.TypeID              { return 2 }
func (c *BackbufferInfo) Flags() atom.Flags                { return 0 }
func (a *BackbufferInfo) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// StartTimer
////////////////////////////////////////////////////////////////////////////////
type StartTimer struct {
	binary.Generate
	observations atom.Observations
	Index        uint8
}

func (a *StartTimer) String() string {
	return fmt.Sprintf("startTimer(index: %v)", a.Index)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The StartTimer pointer is returned so that calls can be chained.
func (a *StartTimer) AddRead(rng memory.Range, id binary.ID) *StartTimer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The StartTimer pointer is returned so that calls can be chained.
func (a *StartTimer) AddWrite(rng memory.Range, id binary.ID) *StartTimer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *StartTimer) API() gfxapi.API                  { return api{} }
func (c *StartTimer) TypeID() atom.TypeID              { return 3 }
func (c *StartTimer) Flags() atom.Flags                { return 0 }
func (a *StartTimer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// StopTimer
////////////////////////////////////////////////////////////////////////////////
type StopTimer struct {
	binary.Generate
	observations atom.Observations
	Index        uint8
	Result       uint64
}

func (a *StopTimer) String() string {
	return fmt.Sprintf("stopTimer(index: %v) → %v", a.Index, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The StopTimer pointer is returned so that calls can be chained.
func (a *StopTimer) AddRead(rng memory.Range, id binary.ID) *StopTimer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The StopTimer pointer is returned so that calls can be chained.
func (a *StopTimer) AddWrite(rng memory.Range, id binary.ID) *StopTimer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *StopTimer) API() gfxapi.API                  { return api{} }
func (c *StopTimer) TypeID() atom.TypeID              { return 4 }
func (c *StopTimer) Flags() atom.Flags                { return 0 }
func (a *StopTimer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// FlushPostBuffer
////////////////////////////////////////////////////////////////////////////////
type FlushPostBuffer struct {
	binary.Generate
	observations atom.Observations
}

func (a *FlushPostBuffer) String() string {
	return fmt.Sprintf("flushPostBuffer()")
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The FlushPostBuffer pointer is returned so that calls can be chained.
func (a *FlushPostBuffer) AddRead(rng memory.Range, id binary.ID) *FlushPostBuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The FlushPostBuffer pointer is returned so that calls can be chained.
func (a *FlushPostBuffer) AddWrite(rng memory.Range, id binary.ID) *FlushPostBuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *FlushPostBuffer) API() gfxapi.API                  { return api{} }
func (c *FlushPostBuffer) TypeID() atom.TypeID              { return 5 }
func (c *FlushPostBuffer) Flags() atom.Flags                { return 0 }
func (a *FlushPostBuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// EglInitialize
////////////////////////////////////////////////////////////////////////////////
type EglInitialize struct {
	binary.Generate
	observations atom.Observations
	Dpy          EGLDisplay
	Major        EGLintᵖ
	Minor        EGLintᵖ
	Result       EGLBoolean
}

func (a *EglInitialize) String() string {
	return fmt.Sprintf("eglInitialize(dpy: %v, major: %v, minor: %v) → %v", a.Dpy, a.Major, a.Minor, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The EglInitialize pointer is returned so that calls can be chained.
func (a *EglInitialize) AddRead(rng memory.Range, id binary.ID) *EglInitialize {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The EglInitialize pointer is returned so that calls can be chained.
func (a *EglInitialize) AddWrite(rng memory.Range, id binary.ID) *EglInitialize {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *EglInitialize) API() gfxapi.API                  { return api{} }
func (c *EglInitialize) TypeID() atom.TypeID              { return 6 }
func (c *EglInitialize) Flags() atom.Flags                { return 0 }
func (a *EglInitialize) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// EglCreateContext
////////////////////////////////////////////////////////////////////////////////
type EglCreateContext struct {
	binary.Generate
	observations atom.Observations
	Display      EGLDisplay
	Config       EGLConfig
	ShareContext EGLContext
	AttribList   EGLintᵖ
	Result       EGLContext
}

func (a *EglCreateContext) String() string {
	return fmt.Sprintf("eglCreateContext(display: %v, config: %v, share_context: %v, attrib_list: %v) → %v", a.Display, a.Config, a.ShareContext, a.AttribList, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The EglCreateContext pointer is returned so that calls can be chained.
func (a *EglCreateContext) AddRead(rng memory.Range, id binary.ID) *EglCreateContext {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The EglCreateContext pointer is returned so that calls can be chained.
func (a *EglCreateContext) AddWrite(rng memory.Range, id binary.ID) *EglCreateContext {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *EglCreateContext) API() gfxapi.API                  { return api{} }
func (c *EglCreateContext) TypeID() atom.TypeID              { return 7 }
func (c *EglCreateContext) Flags() atom.Flags                { return 0 }
func (a *EglCreateContext) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// EglMakeCurrent
////////////////////////////////////////////////////////////////////////////////
type EglMakeCurrent struct {
	binary.Generate
	observations atom.Observations
	Display      EGLDisplay
	Draw         EGLSurface
	Read         EGLSurface
	Context      EGLContext
	Result       EGLBoolean
}

func (a *EglMakeCurrent) String() string {
	return fmt.Sprintf("eglMakeCurrent(display: %v, draw: %v, read: %v, context: %v) → %v", a.Display, a.Draw, a.Read, a.Context, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The EglMakeCurrent pointer is returned so that calls can be chained.
func (a *EglMakeCurrent) AddRead(rng memory.Range, id binary.ID) *EglMakeCurrent {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The EglMakeCurrent pointer is returned so that calls can be chained.
func (a *EglMakeCurrent) AddWrite(rng memory.Range, id binary.ID) *EglMakeCurrent {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *EglMakeCurrent) API() gfxapi.API                  { return api{} }
func (c *EglMakeCurrent) TypeID() atom.TypeID              { return 8 }
func (c *EglMakeCurrent) Flags() atom.Flags                { return 0 }
func (a *EglMakeCurrent) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// EglSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type EglSwapBuffers struct {
	binary.Generate
	observations atom.Observations
	Display      EGLDisplay
	Surface      Voidᵖ
	Result       EGLBoolean
}

func (a *EglSwapBuffers) String() string {
	return fmt.Sprintf("eglSwapBuffers(display: %v, surface: %v) → %v", a.Display, a.Surface, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The EglSwapBuffers pointer is returned so that calls can be chained.
func (a *EglSwapBuffers) AddRead(rng memory.Range, id binary.ID) *EglSwapBuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The EglSwapBuffers pointer is returned so that calls can be chained.
func (a *EglSwapBuffers) AddWrite(rng memory.Range, id binary.ID) *EglSwapBuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *EglSwapBuffers) API() gfxapi.API                  { return api{} }
func (c *EglSwapBuffers) TypeID() atom.TypeID              { return 9 }
func (c *EglSwapBuffers) Flags() atom.Flags                { return 0 | atom.EndOfFrame }
func (a *EglSwapBuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// EglQuerySurface
////////////////////////////////////////////////////////////////////////////////
type EglQuerySurface struct {
	binary.Generate
	observations atom.Observations
	Display      EGLDisplay
	Surface      EGLSurface
	Attribute    EGLint
	Value        EGLintᵖ
	Result       EGLBoolean
}

func (a *EglQuerySurface) String() string {
	return fmt.Sprintf("eglQuerySurface(display: %v, surface: %v, attribute: %v, value: %v) → %v", a.Display, a.Surface, a.Attribute, a.Value, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The EglQuerySurface pointer is returned so that calls can be chained.
func (a *EglQuerySurface) AddRead(rng memory.Range, id binary.ID) *EglQuerySurface {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The EglQuerySurface pointer is returned so that calls can be chained.
func (a *EglQuerySurface) AddWrite(rng memory.Range, id binary.ID) *EglQuerySurface {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *EglQuerySurface) API() gfxapi.API                  { return api{} }
func (c *EglQuerySurface) TypeID() atom.TypeID              { return 10 }
func (c *EglQuerySurface) Flags() atom.Flags                { return 0 }
func (a *EglQuerySurface) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlXCreateContext
////////////////////////////////////////////////////////////////////////////////
type GlXCreateContext struct {
	binary.Generate
	observations atom.Observations
	Dpy          Voidᵖ
	Vis          Voidᵖ
	ShareList    GLXContext
	Direct       bool
	Result       GLXContext
}

func (a *GlXCreateContext) String() string {
	return fmt.Sprintf("glXCreateContext(dpy: %v, vis: %v, shareList: %v, direct: %v) → %v", a.Dpy, a.Vis, a.ShareList, a.Direct, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlXCreateContext pointer is returned so that calls can be chained.
func (a *GlXCreateContext) AddRead(rng memory.Range, id binary.ID) *GlXCreateContext {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlXCreateContext pointer is returned so that calls can be chained.
func (a *GlXCreateContext) AddWrite(rng memory.Range, id binary.ID) *GlXCreateContext {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlXCreateContext) API() gfxapi.API                  { return api{} }
func (c *GlXCreateContext) TypeID() atom.TypeID              { return 11 }
func (c *GlXCreateContext) Flags() atom.Flags                { return 0 }
func (a *GlXCreateContext) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlXCreateNewContext
////////////////////////////////////////////////////////////////////////////////
type GlXCreateNewContext struct {
	binary.Generate
	observations atom.Observations
	Display      Voidᵖ
	Fbconfig     Voidᵖ
	Type         uint32
	Shared       GLXContext
	Direct       bool
	Result       GLXContext
}

func (a *GlXCreateNewContext) String() string {
	return fmt.Sprintf("glXCreateNewContext(display: %v, fbconfig: %v, type: %v, shared: %v, direct: %v) → %v", a.Display, a.Fbconfig, a.Type, a.Shared, a.Direct, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlXCreateNewContext pointer is returned so that calls can be chained.
func (a *GlXCreateNewContext) AddRead(rng memory.Range, id binary.ID) *GlXCreateNewContext {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlXCreateNewContext pointer is returned so that calls can be chained.
func (a *GlXCreateNewContext) AddWrite(rng memory.Range, id binary.ID) *GlXCreateNewContext {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlXCreateNewContext) API() gfxapi.API                  { return api{} }
func (c *GlXCreateNewContext) TypeID() atom.TypeID              { return 12 }
func (c *GlXCreateNewContext) Flags() atom.Flags                { return 0 }
func (a *GlXCreateNewContext) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlXMakeContextCurrent
////////////////////////////////////////////////////////////////////////////////
type GlXMakeContextCurrent struct {
	binary.Generate
	observations atom.Observations
	Display      Voidᵖ
	Draw         GLXDrawable
	Read         GLXDrawable
	Ctx          GLXContext
}

func (a *GlXMakeContextCurrent) String() string {
	return fmt.Sprintf("glXMakeContextCurrent(display: %v, draw: %v, read: %v, ctx: %v)", a.Display, a.Draw, a.Read, a.Ctx)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlXMakeContextCurrent pointer is returned so that calls can be chained.
func (a *GlXMakeContextCurrent) AddRead(rng memory.Range, id binary.ID) *GlXMakeContextCurrent {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlXMakeContextCurrent pointer is returned so that calls can be chained.
func (a *GlXMakeContextCurrent) AddWrite(rng memory.Range, id binary.ID) *GlXMakeContextCurrent {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlXMakeContextCurrent) API() gfxapi.API                  { return api{} }
func (c *GlXMakeContextCurrent) TypeID() atom.TypeID              { return 13 }
func (c *GlXMakeContextCurrent) Flags() atom.Flags                { return 0 }
func (a *GlXMakeContextCurrent) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlXSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type GlXSwapBuffers struct {
	binary.Generate
	observations atom.Observations
	Display      Voidᵖ
	Drawable     GLXDrawable
}

func (a *GlXSwapBuffers) String() string {
	return fmt.Sprintf("glXSwapBuffers(display: %v, drawable: %v)", a.Display, a.Drawable)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlXSwapBuffers pointer is returned so that calls can be chained.
func (a *GlXSwapBuffers) AddRead(rng memory.Range, id binary.ID) *GlXSwapBuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlXSwapBuffers pointer is returned so that calls can be chained.
func (a *GlXSwapBuffers) AddWrite(rng memory.Range, id binary.ID) *GlXSwapBuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlXSwapBuffers) API() gfxapi.API                  { return api{} }
func (c *GlXSwapBuffers) TypeID() atom.TypeID              { return 14 }
func (c *GlXSwapBuffers) Flags() atom.Flags                { return 0 | atom.EndOfFrame }
func (a *GlXSwapBuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// WglCreateContext
////////////////////////////////////////////////////////////////////////////////
type WglCreateContext struct {
	binary.Generate
	observations atom.Observations
	Hdc          HDC
	Result       HGLRC
}

func (a *WglCreateContext) String() string {
	return fmt.Sprintf("wglCreateContext(hdc: %v) → %v", a.Hdc, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The WglCreateContext pointer is returned so that calls can be chained.
func (a *WglCreateContext) AddRead(rng memory.Range, id binary.ID) *WglCreateContext {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The WglCreateContext pointer is returned so that calls can be chained.
func (a *WglCreateContext) AddWrite(rng memory.Range, id binary.ID) *WglCreateContext {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *WglCreateContext) API() gfxapi.API                  { return api{} }
func (c *WglCreateContext) TypeID() atom.TypeID              { return 15 }
func (c *WglCreateContext) Flags() atom.Flags                { return 0 }
func (a *WglCreateContext) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// WglCreateContextAttribsARB
////////////////////////////////////////////////////////////////////////////////
type WglCreateContextAttribsARB struct {
	binary.Generate
	observations  atom.Observations
	Hdc           HDC
	HShareContext HGLRC
	AttribList    Intᵖ
	Result        HGLRC
}

func (a *WglCreateContextAttribsARB) String() string {
	return fmt.Sprintf("wglCreateContextAttribsARB(hdc: %v, hShareContext: %v, attribList: %v) → %v", a.Hdc, a.HShareContext, a.AttribList, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The WglCreateContextAttribsARB pointer is returned so that calls can be chained.
func (a *WglCreateContextAttribsARB) AddRead(rng memory.Range, id binary.ID) *WglCreateContextAttribsARB {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The WglCreateContextAttribsARB pointer is returned so that calls can be chained.
func (a *WglCreateContextAttribsARB) AddWrite(rng memory.Range, id binary.ID) *WglCreateContextAttribsARB {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *WglCreateContextAttribsARB) API() gfxapi.API                  { return api{} }
func (c *WglCreateContextAttribsARB) TypeID() atom.TypeID              { return 16 }
func (c *WglCreateContextAttribsARB) Flags() atom.Flags                { return 0 }
func (a *WglCreateContextAttribsARB) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// WglMakeCurrent
////////////////////////////////////////////////////////////////////////////////
type WglMakeCurrent struct {
	binary.Generate
	observations atom.Observations
	Hdc          HDC
	Hglrc        HGLRC
	Result       BOOL
}

func (a *WglMakeCurrent) String() string {
	return fmt.Sprintf("wglMakeCurrent(hdc: %v, hglrc: %v) → %v", a.Hdc, a.Hglrc, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The WglMakeCurrent pointer is returned so that calls can be chained.
func (a *WglMakeCurrent) AddRead(rng memory.Range, id binary.ID) *WglMakeCurrent {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The WglMakeCurrent pointer is returned so that calls can be chained.
func (a *WglMakeCurrent) AddWrite(rng memory.Range, id binary.ID) *WglMakeCurrent {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *WglMakeCurrent) API() gfxapi.API                  { return api{} }
func (c *WglMakeCurrent) TypeID() atom.TypeID              { return 17 }
func (c *WglMakeCurrent) Flags() atom.Flags                { return 0 }
func (a *WglMakeCurrent) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// WglSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type WglSwapBuffers struct {
	binary.Generate
	observations atom.Observations
	Hdc          HDC
}

func (a *WglSwapBuffers) String() string {
	return fmt.Sprintf("wglSwapBuffers(hdc: %v)", a.Hdc)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The WglSwapBuffers pointer is returned so that calls can be chained.
func (a *WglSwapBuffers) AddRead(rng memory.Range, id binary.ID) *WglSwapBuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The WglSwapBuffers pointer is returned so that calls can be chained.
func (a *WglSwapBuffers) AddWrite(rng memory.Range, id binary.ID) *WglSwapBuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *WglSwapBuffers) API() gfxapi.API                  { return api{} }
func (c *WglSwapBuffers) TypeID() atom.TypeID              { return 18 }
func (c *WglSwapBuffers) Flags() atom.Flags                { return 0 | atom.EndOfFrame }
func (a *WglSwapBuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CGLCreateContext
////////////////////////////////////////////////////////////////////////////////
type CGLCreateContext struct {
	binary.Generate
	observations atom.Observations
	Pix          CGLPixelFormatObj
	Share        CGLContextObj
	Ctx          CGLContextObjᵖ
	Result       CGLError
}

func (a *CGLCreateContext) String() string {
	return fmt.Sprintf("CGLCreateContext(pix: %v, share: %v, ctx: %v) → %v", a.Pix, a.Share, a.Ctx, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CGLCreateContext pointer is returned so that calls can be chained.
func (a *CGLCreateContext) AddRead(rng memory.Range, id binary.ID) *CGLCreateContext {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CGLCreateContext pointer is returned so that calls can be chained.
func (a *CGLCreateContext) AddWrite(rng memory.Range, id binary.ID) *CGLCreateContext {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CGLCreateContext) API() gfxapi.API                  { return api{} }
func (c *CGLCreateContext) TypeID() atom.TypeID              { return 19 }
func (c *CGLCreateContext) Flags() atom.Flags                { return 0 }
func (a *CGLCreateContext) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CGLSetCurrentContext
////////////////////////////////////////////////////////////////////////////////
type CGLSetCurrentContext struct {
	binary.Generate
	observations atom.Observations
	Ctx          CGLContextObj
	Result       CGLError
}

func (a *CGLSetCurrentContext) String() string {
	return fmt.Sprintf("CGLSetCurrentContext(ctx: %v) → %v", a.Ctx, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CGLSetCurrentContext pointer is returned so that calls can be chained.
func (a *CGLSetCurrentContext) AddRead(rng memory.Range, id binary.ID) *CGLSetCurrentContext {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CGLSetCurrentContext pointer is returned so that calls can be chained.
func (a *CGLSetCurrentContext) AddWrite(rng memory.Range, id binary.ID) *CGLSetCurrentContext {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CGLSetCurrentContext) API() gfxapi.API                  { return api{} }
func (c *CGLSetCurrentContext) TypeID() atom.TypeID              { return 20 }
func (c *CGLSetCurrentContext) Flags() atom.Flags                { return 0 }
func (a *CGLSetCurrentContext) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEnableClientState
////////////////////////////////////////////////////////////////////////////////
type GlEnableClientState struct {
	binary.Generate
	observations atom.Observations
	Type         ArrayType
}

func (a *GlEnableClientState) String() string {
	return fmt.Sprintf("glEnableClientState(type: %v)", a.Type)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEnableClientState pointer is returned so that calls can be chained.
func (a *GlEnableClientState) AddRead(rng memory.Range, id binary.ID) *GlEnableClientState {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEnableClientState pointer is returned so that calls can be chained.
func (a *GlEnableClientState) AddWrite(rng memory.Range, id binary.ID) *GlEnableClientState {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEnableClientState) API() gfxapi.API                  { return api{} }
func (c *GlEnableClientState) TypeID() atom.TypeID              { return 21 }
func (c *GlEnableClientState) Flags() atom.Flags                { return 0 }
func (a *GlEnableClientState) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDisableClientState
////////////////////////////////////////////////////////////////////////////////
type GlDisableClientState struct {
	binary.Generate
	observations atom.Observations
	Type         ArrayType
}

func (a *GlDisableClientState) String() string {
	return fmt.Sprintf("glDisableClientState(type: %v)", a.Type)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDisableClientState pointer is returned so that calls can be chained.
func (a *GlDisableClientState) AddRead(rng memory.Range, id binary.ID) *GlDisableClientState {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDisableClientState pointer is returned so that calls can be chained.
func (a *GlDisableClientState) AddWrite(rng memory.Range, id binary.ID) *GlDisableClientState {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDisableClientState) API() gfxapi.API                  { return api{} }
func (c *GlDisableClientState) TypeID() atom.TypeID              { return 22 }
func (c *GlDisableClientState) Flags() atom.Flags                { return 0 }
func (a *GlDisableClientState) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramBinaryOES
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramBinaryOES struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	BufferSize   int32
	BytesWritten S32ᵖ
	BinaryFormat U32ᵖ
	Binary       Voidᵖ
}

func (a *GlGetProgramBinaryOES) String() string {
	return fmt.Sprintf("glGetProgramBinaryOES(program: %v, buffer_size: %v, bytes_written: %v, binary_format: %v, binary: %v)", a.Program, a.BufferSize, a.BytesWritten, a.BinaryFormat, a.Binary)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetProgramBinaryOES pointer is returned so that calls can be chained.
func (a *GlGetProgramBinaryOES) AddRead(rng memory.Range, id binary.ID) *GlGetProgramBinaryOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetProgramBinaryOES pointer is returned so that calls can be chained.
func (a *GlGetProgramBinaryOES) AddWrite(rng memory.Range, id binary.ID) *GlGetProgramBinaryOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetProgramBinaryOES) API() gfxapi.API                  { return api{} }
func (c *GlGetProgramBinaryOES) TypeID() atom.TypeID              { return 23 }
func (c *GlGetProgramBinaryOES) Flags() atom.Flags                { return 0 }
func (a *GlGetProgramBinaryOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlProgramBinaryOES
////////////////////////////////////////////////////////////////////////////////
type GlProgramBinaryOES struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	BinaryFormat uint32
	Binary       Voidᵖ
	BinarySize   int32
}

func (a *GlProgramBinaryOES) String() string {
	return fmt.Sprintf("glProgramBinaryOES(program: %v, binary_format: %v, binary: %v, binary_size: %v)", a.Program, a.BinaryFormat, a.Binary, a.BinarySize)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlProgramBinaryOES pointer is returned so that calls can be chained.
func (a *GlProgramBinaryOES) AddRead(rng memory.Range, id binary.ID) *GlProgramBinaryOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlProgramBinaryOES pointer is returned so that calls can be chained.
func (a *GlProgramBinaryOES) AddWrite(rng memory.Range, id binary.ID) *GlProgramBinaryOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlProgramBinaryOES) API() gfxapi.API                  { return api{} }
func (c *GlProgramBinaryOES) TypeID() atom.TypeID              { return 24 }
func (c *GlProgramBinaryOES) Flags() atom.Flags                { return 0 }
func (a *GlProgramBinaryOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStartTilingQCOM
////////////////////////////////////////////////////////////////////////////////
type GlStartTilingQCOM struct {
	binary.Generate
	observations atom.Observations
	X            int32
	Y            int32
	Width        int32
	Height       int32
	PreserveMask TilePreserveMaskQCOM
}

func (a *GlStartTilingQCOM) String() string {
	return fmt.Sprintf("glStartTilingQCOM(x: %v, y: %v, width: %v, height: %v, preserveMask: %v)", a.X, a.Y, a.Width, a.Height, a.PreserveMask)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlStartTilingQCOM pointer is returned so that calls can be chained.
func (a *GlStartTilingQCOM) AddRead(rng memory.Range, id binary.ID) *GlStartTilingQCOM {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlStartTilingQCOM pointer is returned so that calls can be chained.
func (a *GlStartTilingQCOM) AddWrite(rng memory.Range, id binary.ID) *GlStartTilingQCOM {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlStartTilingQCOM) API() gfxapi.API                  { return api{} }
func (c *GlStartTilingQCOM) TypeID() atom.TypeID              { return 25 }
func (c *GlStartTilingQCOM) Flags() atom.Flags                { return 0 }
func (a *GlStartTilingQCOM) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEndTilingQCOM
////////////////////////////////////////////////////////////////////////////////
type GlEndTilingQCOM struct {
	binary.Generate
	observations atom.Observations
	PreserveMask TilePreserveMaskQCOM
}

func (a *GlEndTilingQCOM) String() string {
	return fmt.Sprintf("glEndTilingQCOM(preserve_mask: %v)", a.PreserveMask)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEndTilingQCOM pointer is returned so that calls can be chained.
func (a *GlEndTilingQCOM) AddRead(rng memory.Range, id binary.ID) *GlEndTilingQCOM {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEndTilingQCOM pointer is returned so that calls can be chained.
func (a *GlEndTilingQCOM) AddWrite(rng memory.Range, id binary.ID) *GlEndTilingQCOM {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEndTilingQCOM) API() gfxapi.API                  { return api{} }
func (c *GlEndTilingQCOM) TypeID() atom.TypeID              { return 26 }
func (c *GlEndTilingQCOM) Flags() atom.Flags                { return 0 }
func (a *GlEndTilingQCOM) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDiscardFramebufferEXT
////////////////////////////////////////////////////////////////////////////////
type GlDiscardFramebufferEXT struct {
	binary.Generate
	observations   atom.Observations
	Target         FramebufferTarget
	NumAttachments int32
	Attachments    DiscardFramebufferAttachmentᵖ
}

func (a *GlDiscardFramebufferEXT) String() string {
	return fmt.Sprintf("glDiscardFramebufferEXT(target: %v, numAttachments: %v, attachments: %v)", a.Target, a.NumAttachments, a.Attachments)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDiscardFramebufferEXT pointer is returned so that calls can be chained.
func (a *GlDiscardFramebufferEXT) AddRead(rng memory.Range, id binary.ID) *GlDiscardFramebufferEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDiscardFramebufferEXT pointer is returned so that calls can be chained.
func (a *GlDiscardFramebufferEXT) AddWrite(rng memory.Range, id binary.ID) *GlDiscardFramebufferEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDiscardFramebufferEXT) API() gfxapi.API                  { return api{} }
func (c *GlDiscardFramebufferEXT) TypeID() atom.TypeID              { return 27 }
func (c *GlDiscardFramebufferEXT) Flags() atom.Flags                { return 0 }
func (a *GlDiscardFramebufferEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlInsertEventMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlInsertEventMarkerEXT struct {
	binary.Generate
	observations atom.Observations
	Length       int32
	Marker       string
}

func (a *GlInsertEventMarkerEXT) String() string {
	return fmt.Sprintf("glInsertEventMarkerEXT(length: %v, marker: %v)", a.Length, a.Marker)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlInsertEventMarkerEXT pointer is returned so that calls can be chained.
func (a *GlInsertEventMarkerEXT) AddRead(rng memory.Range, id binary.ID) *GlInsertEventMarkerEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlInsertEventMarkerEXT pointer is returned so that calls can be chained.
func (a *GlInsertEventMarkerEXT) AddWrite(rng memory.Range, id binary.ID) *GlInsertEventMarkerEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlInsertEventMarkerEXT) API() gfxapi.API                  { return api{} }
func (c *GlInsertEventMarkerEXT) TypeID() atom.TypeID              { return 28 }
func (c *GlInsertEventMarkerEXT) Flags() atom.Flags                { return 0 }
func (a *GlInsertEventMarkerEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlPushGroupMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlPushGroupMarkerEXT struct {
	binary.Generate
	observations atom.Observations
	Length       int32
	Marker       string
}

func (a *GlPushGroupMarkerEXT) String() string {
	return fmt.Sprintf("glPushGroupMarkerEXT(length: %v, marker: %v)", a.Length, a.Marker)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlPushGroupMarkerEXT pointer is returned so that calls can be chained.
func (a *GlPushGroupMarkerEXT) AddRead(rng memory.Range, id binary.ID) *GlPushGroupMarkerEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlPushGroupMarkerEXT pointer is returned so that calls can be chained.
func (a *GlPushGroupMarkerEXT) AddWrite(rng memory.Range, id binary.ID) *GlPushGroupMarkerEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlPushGroupMarkerEXT) API() gfxapi.API                  { return api{} }
func (c *GlPushGroupMarkerEXT) TypeID() atom.TypeID              { return 29 }
func (c *GlPushGroupMarkerEXT) Flags() atom.Flags                { return 0 }
func (a *GlPushGroupMarkerEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlPopGroupMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlPopGroupMarkerEXT struct {
	binary.Generate
	observations atom.Observations
}

func (a *GlPopGroupMarkerEXT) String() string {
	return fmt.Sprintf("glPopGroupMarkerEXT()")
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlPopGroupMarkerEXT pointer is returned so that calls can be chained.
func (a *GlPopGroupMarkerEXT) AddRead(rng memory.Range, id binary.ID) *GlPopGroupMarkerEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlPopGroupMarkerEXT pointer is returned so that calls can be chained.
func (a *GlPopGroupMarkerEXT) AddWrite(rng memory.Range, id binary.ID) *GlPopGroupMarkerEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlPopGroupMarkerEXT) API() gfxapi.API                  { return api{} }
func (c *GlPopGroupMarkerEXT) TypeID() atom.TypeID              { return 30 }
func (c *GlPopGroupMarkerEXT) Flags() atom.Flags                { return 0 }
func (a *GlPopGroupMarkerEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage1DEXT struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Levels       int32
	Format       TexelFormat
	Width        int32
}

func (a *GlTexStorage1DEXT) String() string {
	return fmt.Sprintf("glTexStorage1DEXT(target: %v, levels: %v, format: %v, width: %v)", a.Target, a.Levels, a.Format, a.Width)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTexStorage1DEXT pointer is returned so that calls can be chained.
func (a *GlTexStorage1DEXT) AddRead(rng memory.Range, id binary.ID) *GlTexStorage1DEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTexStorage1DEXT pointer is returned so that calls can be chained.
func (a *GlTexStorage1DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTexStorage1DEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTexStorage1DEXT) API() gfxapi.API                  { return api{} }
func (c *GlTexStorage1DEXT) TypeID() atom.TypeID              { return 31 }
func (c *GlTexStorage1DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTexStorage1DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage2DEXT struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Levels       int32
	Format       TexelFormat
	Width        int32
	Height       int32
}

func (a *GlTexStorage2DEXT) String() string {
	return fmt.Sprintf("glTexStorage2DEXT(target: %v, levels: %v, format: %v, width: %v, height: %v)", a.Target, a.Levels, a.Format, a.Width, a.Height)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTexStorage2DEXT pointer is returned so that calls can be chained.
func (a *GlTexStorage2DEXT) AddRead(rng memory.Range, id binary.ID) *GlTexStorage2DEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTexStorage2DEXT pointer is returned so that calls can be chained.
func (a *GlTexStorage2DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTexStorage2DEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTexStorage2DEXT) API() gfxapi.API                  { return api{} }
func (c *GlTexStorage2DEXT) TypeID() atom.TypeID              { return 32 }
func (c *GlTexStorage2DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTexStorage2DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage3DEXT struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Levels       int32
	Format       TexelFormat
	Width        int32
	Height       int32
	Depth        int32
}

func (a *GlTexStorage3DEXT) String() string {
	return fmt.Sprintf("glTexStorage3DEXT(target: %v, levels: %v, format: %v, width: %v, height: %v, depth: %v)", a.Target, a.Levels, a.Format, a.Width, a.Height, a.Depth)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTexStorage3DEXT pointer is returned so that calls can be chained.
func (a *GlTexStorage3DEXT) AddRead(rng memory.Range, id binary.ID) *GlTexStorage3DEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTexStorage3DEXT pointer is returned so that calls can be chained.
func (a *GlTexStorage3DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTexStorage3DEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTexStorage3DEXT) API() gfxapi.API                  { return api{} }
func (c *GlTexStorage3DEXT) TypeID() atom.TypeID              { return 33 }
func (c *GlTexStorage3DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTexStorage3DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage1DEXT struct {
	binary.Generate
	observations atom.Observations
	Texture      TextureId
	Target       TextureTarget
	Levels       int32
	Format       TexelFormat
	Width        int32
}

func (a *GlTextureStorage1DEXT) String() string {
	return fmt.Sprintf("glTextureStorage1DEXT(texture: %v, target: %v, levels: %v, format: %v, width: %v)", a.Texture, a.Target, a.Levels, a.Format, a.Width)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTextureStorage1DEXT pointer is returned so that calls can be chained.
func (a *GlTextureStorage1DEXT) AddRead(rng memory.Range, id binary.ID) *GlTextureStorage1DEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTextureStorage1DEXT pointer is returned so that calls can be chained.
func (a *GlTextureStorage1DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTextureStorage1DEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTextureStorage1DEXT) API() gfxapi.API                  { return api{} }
func (c *GlTextureStorage1DEXT) TypeID() atom.TypeID              { return 34 }
func (c *GlTextureStorage1DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTextureStorage1DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage2DEXT struct {
	binary.Generate
	observations atom.Observations
	Texture      TextureId
	Target       TextureTarget
	Levels       int32
	Format       TexelFormat
	Width        int32
	Height       int32
}

func (a *GlTextureStorage2DEXT) String() string {
	return fmt.Sprintf("glTextureStorage2DEXT(texture: %v, target: %v, levels: %v, format: %v, width: %v, height: %v)", a.Texture, a.Target, a.Levels, a.Format, a.Width, a.Height)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTextureStorage2DEXT pointer is returned so that calls can be chained.
func (a *GlTextureStorage2DEXT) AddRead(rng memory.Range, id binary.ID) *GlTextureStorage2DEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTextureStorage2DEXT pointer is returned so that calls can be chained.
func (a *GlTextureStorage2DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTextureStorage2DEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTextureStorage2DEXT) API() gfxapi.API                  { return api{} }
func (c *GlTextureStorage2DEXT) TypeID() atom.TypeID              { return 35 }
func (c *GlTextureStorage2DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTextureStorage2DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage3DEXT struct {
	binary.Generate
	observations atom.Observations
	Texture      TextureId
	Target       TextureTarget
	Levels       int32
	Format       TexelFormat
	Width        int32
	Height       int32
	Depth        int32
}

func (a *GlTextureStorage3DEXT) String() string {
	return fmt.Sprintf("glTextureStorage3DEXT(texture: %v, target: %v, levels: %v, format: %v, width: %v, height: %v, depth: %v)", a.Texture, a.Target, a.Levels, a.Format, a.Width, a.Height, a.Depth)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTextureStorage3DEXT pointer is returned so that calls can be chained.
func (a *GlTextureStorage3DEXT) AddRead(rng memory.Range, id binary.ID) *GlTextureStorage3DEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTextureStorage3DEXT pointer is returned so that calls can be chained.
func (a *GlTextureStorage3DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTextureStorage3DEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTextureStorage3DEXT) API() gfxapi.API                  { return api{} }
func (c *GlTextureStorage3DEXT) TypeID() atom.TypeID              { return 36 }
func (c *GlTextureStorage3DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTextureStorage3DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenVertexArraysOES
////////////////////////////////////////////////////////////////////////////////
type GlGenVertexArraysOES struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Arrays       VertexArrayIdᵖ
}

func (a *GlGenVertexArraysOES) String() string {
	return fmt.Sprintf("glGenVertexArraysOES(count: %v, arrays: %v)", a.Count, a.Arrays)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenVertexArraysOES pointer is returned so that calls can be chained.
func (a *GlGenVertexArraysOES) AddRead(rng memory.Range, id binary.ID) *GlGenVertexArraysOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenVertexArraysOES pointer is returned so that calls can be chained.
func (a *GlGenVertexArraysOES) AddWrite(rng memory.Range, id binary.ID) *GlGenVertexArraysOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenVertexArraysOES) API() gfxapi.API                  { return api{} }
func (c *GlGenVertexArraysOES) TypeID() atom.TypeID              { return 37 }
func (c *GlGenVertexArraysOES) Flags() atom.Flags                { return 0 }
func (a *GlGenVertexArraysOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindVertexArrayOES
////////////////////////////////////////////////////////////////////////////////
type GlBindVertexArrayOES struct {
	binary.Generate
	observations atom.Observations
	Array        VertexArrayId
}

func (a *GlBindVertexArrayOES) String() string {
	return fmt.Sprintf("glBindVertexArrayOES(array: %v)", a.Array)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindVertexArrayOES pointer is returned so that calls can be chained.
func (a *GlBindVertexArrayOES) AddRead(rng memory.Range, id binary.ID) *GlBindVertexArrayOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindVertexArrayOES pointer is returned so that calls can be chained.
func (a *GlBindVertexArrayOES) AddWrite(rng memory.Range, id binary.ID) *GlBindVertexArrayOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindVertexArrayOES) API() gfxapi.API                  { return api{} }
func (c *GlBindVertexArrayOES) TypeID() atom.TypeID              { return 38 }
func (c *GlBindVertexArrayOES) Flags() atom.Flags                { return 0 }
func (a *GlBindVertexArrayOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteVertexArraysOES
////////////////////////////////////////////////////////////////////////////////
type GlDeleteVertexArraysOES struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Arrays       VertexArrayIdᵖ
}

func (a *GlDeleteVertexArraysOES) String() string {
	return fmt.Sprintf("glDeleteVertexArraysOES(count: %v, arrays: %v)", a.Count, a.Arrays)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteVertexArraysOES pointer is returned so that calls can be chained.
func (a *GlDeleteVertexArraysOES) AddRead(rng memory.Range, id binary.ID) *GlDeleteVertexArraysOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteVertexArraysOES pointer is returned so that calls can be chained.
func (a *GlDeleteVertexArraysOES) AddWrite(rng memory.Range, id binary.ID) *GlDeleteVertexArraysOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteVertexArraysOES) API() gfxapi.API                  { return api{} }
func (c *GlDeleteVertexArraysOES) TypeID() atom.TypeID              { return 39 }
func (c *GlDeleteVertexArraysOES) Flags() atom.Flags                { return 0 }
func (a *GlDeleteVertexArraysOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsVertexArrayOES
////////////////////////////////////////////////////////////////////////////////
type GlIsVertexArrayOES struct {
	binary.Generate
	observations atom.Observations
	Array        VertexArrayId
	Result       bool
}

func (a *GlIsVertexArrayOES) String() string {
	return fmt.Sprintf("glIsVertexArrayOES(array: %v) → %v", a.Array, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsVertexArrayOES pointer is returned so that calls can be chained.
func (a *GlIsVertexArrayOES) AddRead(rng memory.Range, id binary.ID) *GlIsVertexArrayOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsVertexArrayOES pointer is returned so that calls can be chained.
func (a *GlIsVertexArrayOES) AddWrite(rng memory.Range, id binary.ID) *GlIsVertexArrayOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsVertexArrayOES) API() gfxapi.API                  { return api{} }
func (c *GlIsVertexArrayOES) TypeID() atom.TypeID              { return 40 }
func (c *GlIsVertexArrayOES) Flags() atom.Flags                { return 0 }
func (a *GlIsVertexArrayOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetTexture2DOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetTexture2DOES struct {
	binary.Generate
	observations atom.Observations
	Target       ImageTargetTexture
	Image        ImageOES
}

func (a *GlEGLImageTargetTexture2DOES) String() string {
	return fmt.Sprintf("glEGLImageTargetTexture2DOES(target: %v, image: %v)", a.Target, a.Image)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEGLImageTargetTexture2DOES pointer is returned so that calls can be chained.
func (a *GlEGLImageTargetTexture2DOES) AddRead(rng memory.Range, id binary.ID) *GlEGLImageTargetTexture2DOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEGLImageTargetTexture2DOES pointer is returned so that calls can be chained.
func (a *GlEGLImageTargetTexture2DOES) AddWrite(rng memory.Range, id binary.ID) *GlEGLImageTargetTexture2DOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEGLImageTargetTexture2DOES) API() gfxapi.API                  { return api{} }
func (c *GlEGLImageTargetTexture2DOES) TypeID() atom.TypeID              { return 41 }
func (c *GlEGLImageTargetTexture2DOES) Flags() atom.Flags                { return 0 }
func (a *GlEGLImageTargetTexture2DOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetRenderbufferStorageOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetRenderbufferStorageOES struct {
	binary.Generate
	observations atom.Observations
	Target       ImageTargetRenderbufferStorage
	Image        TexturePointer
}

func (a *GlEGLImageTargetRenderbufferStorageOES) String() string {
	return fmt.Sprintf("glEGLImageTargetRenderbufferStorageOES(target: %v, image: %v)", a.Target, a.Image)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEGLImageTargetRenderbufferStorageOES pointer is returned so that calls can be chained.
func (a *GlEGLImageTargetRenderbufferStorageOES) AddRead(rng memory.Range, id binary.ID) *GlEGLImageTargetRenderbufferStorageOES {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEGLImageTargetRenderbufferStorageOES pointer is returned so that calls can be chained.
func (a *GlEGLImageTargetRenderbufferStorageOES) AddWrite(rng memory.Range, id binary.ID) *GlEGLImageTargetRenderbufferStorageOES {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEGLImageTargetRenderbufferStorageOES) API() gfxapi.API     { return api{} }
func (c *GlEGLImageTargetRenderbufferStorageOES) TypeID() atom.TypeID { return 42 }
func (c *GlEGLImageTargetRenderbufferStorageOES) Flags() atom.Flags   { return 0 }
func (a *GlEGLImageTargetRenderbufferStorageOES) Observations() *atom.Observations {
	return &a.observations
}

////////////////////////////////////////////////////////////////////////////////
// GlGetGraphicsResetStatusEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetGraphicsResetStatusEXT struct {
	binary.Generate
	observations atom.Observations
	Result       ResetStatus
}

func (a *GlGetGraphicsResetStatusEXT) String() string {
	return fmt.Sprintf("glGetGraphicsResetStatusEXT() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetGraphicsResetStatusEXT pointer is returned so that calls can be chained.
func (a *GlGetGraphicsResetStatusEXT) AddRead(rng memory.Range, id binary.ID) *GlGetGraphicsResetStatusEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetGraphicsResetStatusEXT pointer is returned so that calls can be chained.
func (a *GlGetGraphicsResetStatusEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetGraphicsResetStatusEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetGraphicsResetStatusEXT) API() gfxapi.API                  { return api{} }
func (c *GlGetGraphicsResetStatusEXT) TypeID() atom.TypeID              { return 43 }
func (c *GlGetGraphicsResetStatusEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetGraphicsResetStatusEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindAttribLocation
////////////////////////////////////////////////////////////////////////////////
type GlBindAttribLocation struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Location     AttributeLocation
	Name         string
}

func (a *GlBindAttribLocation) String() string {
	return fmt.Sprintf("glBindAttribLocation(program: %v, location: %v, name: %v)", a.Program, a.Location, a.Name)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindAttribLocation pointer is returned so that calls can be chained.
func (a *GlBindAttribLocation) AddRead(rng memory.Range, id binary.ID) *GlBindAttribLocation {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindAttribLocation pointer is returned so that calls can be chained.
func (a *GlBindAttribLocation) AddWrite(rng memory.Range, id binary.ID) *GlBindAttribLocation {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindAttribLocation) API() gfxapi.API                  { return api{} }
func (c *GlBindAttribLocation) TypeID() atom.TypeID              { return 44 }
func (c *GlBindAttribLocation) Flags() atom.Flags                { return 0 }
func (a *GlBindAttribLocation) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendFunc
////////////////////////////////////////////////////////////////////////////////
type GlBlendFunc struct {
	binary.Generate
	observations atom.Observations
	SrcFactor    BlendFactor
	DstFactor    BlendFactor
}

func (a *GlBlendFunc) String() string {
	return fmt.Sprintf("glBlendFunc(src_factor: %v, dst_factor: %v)", a.SrcFactor, a.DstFactor)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBlendFunc pointer is returned so that calls can be chained.
func (a *GlBlendFunc) AddRead(rng memory.Range, id binary.ID) *GlBlendFunc {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBlendFunc pointer is returned so that calls can be chained.
func (a *GlBlendFunc) AddWrite(rng memory.Range, id binary.ID) *GlBlendFunc {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBlendFunc) API() gfxapi.API                  { return api{} }
func (c *GlBlendFunc) TypeID() atom.TypeID              { return 45 }
func (c *GlBlendFunc) Flags() atom.Flags                { return 0 }
func (a *GlBlendFunc) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendFuncSeparate struct {
	binary.Generate
	observations   atom.Observations
	SrcFactorRgb   BlendFactor
	DstFactorRgb   BlendFactor
	SrcFactorAlpha BlendFactor
	DstFactorAlpha BlendFactor
}

func (a *GlBlendFuncSeparate) String() string {
	return fmt.Sprintf("glBlendFuncSeparate(src_factor_rgb: %v, dst_factor_rgb: %v, src_factor_alpha: %v, dst_factor_alpha: %v)", a.SrcFactorRgb, a.DstFactorRgb, a.SrcFactorAlpha, a.DstFactorAlpha)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBlendFuncSeparate pointer is returned so that calls can be chained.
func (a *GlBlendFuncSeparate) AddRead(rng memory.Range, id binary.ID) *GlBlendFuncSeparate {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBlendFuncSeparate pointer is returned so that calls can be chained.
func (a *GlBlendFuncSeparate) AddWrite(rng memory.Range, id binary.ID) *GlBlendFuncSeparate {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBlendFuncSeparate) API() gfxapi.API                  { return api{} }
func (c *GlBlendFuncSeparate) TypeID() atom.TypeID              { return 46 }
func (c *GlBlendFuncSeparate) Flags() atom.Flags                { return 0 }
func (a *GlBlendFuncSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquation
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquation struct {
	binary.Generate
	observations atom.Observations
	Equation     BlendEquation
}

func (a *GlBlendEquation) String() string {
	return fmt.Sprintf("glBlendEquation(equation: %v)", a.Equation)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBlendEquation pointer is returned so that calls can be chained.
func (a *GlBlendEquation) AddRead(rng memory.Range, id binary.ID) *GlBlendEquation {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBlendEquation pointer is returned so that calls can be chained.
func (a *GlBlendEquation) AddWrite(rng memory.Range, id binary.ID) *GlBlendEquation {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBlendEquation) API() gfxapi.API                  { return api{} }
func (c *GlBlendEquation) TypeID() atom.TypeID              { return 47 }
func (c *GlBlendEquation) Flags() atom.Flags                { return 0 }
func (a *GlBlendEquation) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquationSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquationSeparate struct {
	binary.Generate
	observations atom.Observations
	Rgb          BlendEquation
	Alpha        BlendEquation
}

func (a *GlBlendEquationSeparate) String() string {
	return fmt.Sprintf("glBlendEquationSeparate(rgb: %v, alpha: %v)", a.Rgb, a.Alpha)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBlendEquationSeparate pointer is returned so that calls can be chained.
func (a *GlBlendEquationSeparate) AddRead(rng memory.Range, id binary.ID) *GlBlendEquationSeparate {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBlendEquationSeparate pointer is returned so that calls can be chained.
func (a *GlBlendEquationSeparate) AddWrite(rng memory.Range, id binary.ID) *GlBlendEquationSeparate {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBlendEquationSeparate) API() gfxapi.API                  { return api{} }
func (c *GlBlendEquationSeparate) TypeID() atom.TypeID              { return 48 }
func (c *GlBlendEquationSeparate) Flags() atom.Flags                { return 0 }
func (a *GlBlendEquationSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendColor
////////////////////////////////////////////////////////////////////////////////
type GlBlendColor struct {
	binary.Generate
	observations atom.Observations
	Red          float32
	Green        float32
	Blue         float32
	Alpha        float32
}

func (a *GlBlendColor) String() string {
	return fmt.Sprintf("glBlendColor(red: %v, green: %v, blue: %v, alpha: %v)", a.Red, a.Green, a.Blue, a.Alpha)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBlendColor pointer is returned so that calls can be chained.
func (a *GlBlendColor) AddRead(rng memory.Range, id binary.ID) *GlBlendColor {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBlendColor pointer is returned so that calls can be chained.
func (a *GlBlendColor) AddWrite(rng memory.Range, id binary.ID) *GlBlendColor {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBlendColor) API() gfxapi.API                  { return api{} }
func (c *GlBlendColor) TypeID() atom.TypeID              { return 49 }
func (c *GlBlendColor) Flags() atom.Flags                { return 0 }
func (a *GlBlendColor) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEnableVertexAttribArray
////////////////////////////////////////////////////////////////////////////////
type GlEnableVertexAttribArray struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
}

func (a *GlEnableVertexAttribArray) String() string {
	return fmt.Sprintf("glEnableVertexAttribArray(location: %v)", a.Location)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEnableVertexAttribArray pointer is returned so that calls can be chained.
func (a *GlEnableVertexAttribArray) AddRead(rng memory.Range, id binary.ID) *GlEnableVertexAttribArray {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEnableVertexAttribArray pointer is returned so that calls can be chained.
func (a *GlEnableVertexAttribArray) AddWrite(rng memory.Range, id binary.ID) *GlEnableVertexAttribArray {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEnableVertexAttribArray) API() gfxapi.API                  { return api{} }
func (c *GlEnableVertexAttribArray) TypeID() atom.TypeID              { return 50 }
func (c *GlEnableVertexAttribArray) Flags() atom.Flags                { return 0 }
func (a *GlEnableVertexAttribArray) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDisableVertexAttribArray
////////////////////////////////////////////////////////////////////////////////
type GlDisableVertexAttribArray struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
}

func (a *GlDisableVertexAttribArray) String() string {
	return fmt.Sprintf("glDisableVertexAttribArray(location: %v)", a.Location)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDisableVertexAttribArray pointer is returned so that calls can be chained.
func (a *GlDisableVertexAttribArray) AddRead(rng memory.Range, id binary.ID) *GlDisableVertexAttribArray {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDisableVertexAttribArray pointer is returned so that calls can be chained.
func (a *GlDisableVertexAttribArray) AddWrite(rng memory.Range, id binary.ID) *GlDisableVertexAttribArray {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDisableVertexAttribArray) API() gfxapi.API                  { return api{} }
func (c *GlDisableVertexAttribArray) TypeID() atom.TypeID              { return 51 }
func (c *GlDisableVertexAttribArray) Flags() atom.Flags                { return 0 }
func (a *GlDisableVertexAttribArray) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttribPointer
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttribPointer struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Size         int32
	Type         VertexAttribType
	Normalized   bool
	Stride       int32
	Data         VertexPointer
}

func (a *GlVertexAttribPointer) String() string {
	return fmt.Sprintf("glVertexAttribPointer(location: %v, size: %v, type: %v, normalized: %v, stride: %v, data: %v)", a.Location, a.Size, a.Type, a.Normalized, a.Stride, a.Data)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttribPointer pointer is returned so that calls can be chained.
func (a *GlVertexAttribPointer) AddRead(rng memory.Range, id binary.ID) *GlVertexAttribPointer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttribPointer pointer is returned so that calls can be chained.
func (a *GlVertexAttribPointer) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttribPointer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttribPointer) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttribPointer) TypeID() atom.TypeID              { return 52 }
func (c *GlVertexAttribPointer) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttribPointer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveAttrib
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveAttrib struct {
	binary.Generate
	observations       atom.Observations
	Program            ProgramId
	Location           AttributeLocation
	BufferSize         int32
	BufferBytesWritten S32ᵖ
	VectorCount        S32ᵖ
	Type               ShaderAttribTypeᵖ
	Name               Charᵖ
}

func (a *GlGetActiveAttrib) String() string {
	return fmt.Sprintf("glGetActiveAttrib(program: %v, location: %v, buffer_size: %v, buffer_bytes_written: %v, vector_count: %v, type: %v, name: %v)", a.Program, a.Location, a.BufferSize, a.BufferBytesWritten, a.VectorCount, a.Type, a.Name)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetActiveAttrib pointer is returned so that calls can be chained.
func (a *GlGetActiveAttrib) AddRead(rng memory.Range, id binary.ID) *GlGetActiveAttrib {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetActiveAttrib pointer is returned so that calls can be chained.
func (a *GlGetActiveAttrib) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveAttrib {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetActiveAttrib) API() gfxapi.API                  { return api{} }
func (c *GlGetActiveAttrib) TypeID() atom.TypeID              { return 53 }
func (c *GlGetActiveAttrib) Flags() atom.Flags                { return 0 }
func (a *GlGetActiveAttrib) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveUniform
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveUniform struct {
	binary.Generate
	observations       atom.Observations
	Program            ProgramId
	Location           int32
	BufferSize         int32
	BufferBytesWritten S32ᵖ
	Size               S32ᵖ
	Type               ShaderUniformTypeᵖ
	Name               Charᵖ
}

func (a *GlGetActiveUniform) String() string {
	return fmt.Sprintf("glGetActiveUniform(program: %v, location: %v, buffer_size: %v, buffer_bytes_written: %v, size: %v, type: %v, name: %v)", a.Program, a.Location, a.BufferSize, a.BufferBytesWritten, a.Size, a.Type, a.Name)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniform pointer is returned so that calls can be chained.
func (a *GlGetActiveUniform) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniform {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniform pointer is returned so that calls can be chained.
func (a *GlGetActiveUniform) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniform {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetActiveUniform) API() gfxapi.API                  { return api{} }
func (c *GlGetActiveUniform) TypeID() atom.TypeID              { return 54 }
func (c *GlGetActiveUniform) Flags() atom.Flags                { return 0 }
func (a *GlGetActiveUniform) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetError
////////////////////////////////////////////////////////////////////////////////
type GlGetError struct {
	binary.Generate
	observations atom.Observations
	Result       Error
}

func (a *GlGetError) String() string {
	return fmt.Sprintf("glGetError() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetError pointer is returned so that calls can be chained.
func (a *GlGetError) AddRead(rng memory.Range, id binary.ID) *GlGetError {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetError pointer is returned so that calls can be chained.
func (a *GlGetError) AddWrite(rng memory.Range, id binary.ID) *GlGetError {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetError) API() gfxapi.API                  { return api{} }
func (c *GlGetError) TypeID() atom.TypeID              { return 55 }
func (c *GlGetError) Flags() atom.Flags                { return 0 }
func (a *GlGetError) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramiv
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramiv struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Parameter    ProgramParameter
	Value        S32ᵖ
}

func (a *GlGetProgramiv) String() string {
	return fmt.Sprintf("glGetProgramiv(program: %v, parameter: %v, value: %v)", a.Program, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetProgramiv pointer is returned so that calls can be chained.
func (a *GlGetProgramiv) AddRead(rng memory.Range, id binary.ID) *GlGetProgramiv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetProgramiv pointer is returned so that calls can be chained.
func (a *GlGetProgramiv) AddWrite(rng memory.Range, id binary.ID) *GlGetProgramiv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetProgramiv) API() gfxapi.API                  { return api{} }
func (c *GlGetProgramiv) TypeID() atom.TypeID              { return 56 }
func (c *GlGetProgramiv) Flags() atom.Flags                { return 0 }
func (a *GlGetProgramiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderiv
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderiv struct {
	binary.Generate
	observations atom.Observations
	Shader       ShaderId
	Parameter    ShaderParameter
	Value        S32ᵖ
}

func (a *GlGetShaderiv) String() string {
	return fmt.Sprintf("glGetShaderiv(shader: %v, parameter: %v, value: %v)", a.Shader, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetShaderiv pointer is returned so that calls can be chained.
func (a *GlGetShaderiv) AddRead(rng memory.Range, id binary.ID) *GlGetShaderiv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetShaderiv pointer is returned so that calls can be chained.
func (a *GlGetShaderiv) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderiv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetShaderiv) API() gfxapi.API                  { return api{} }
func (c *GlGetShaderiv) TypeID() atom.TypeID              { return 57 }
func (c *GlGetShaderiv) Flags() atom.Flags                { return 0 }
func (a *GlGetShaderiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformLocation
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformLocation struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Name         string
	Result       UniformLocation
}

func (a *GlGetUniformLocation) String() string {
	return fmt.Sprintf("glGetUniformLocation(program: %v, name: %v) → %v", a.Program, a.Name, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetUniformLocation pointer is returned so that calls can be chained.
func (a *GlGetUniformLocation) AddRead(rng memory.Range, id binary.ID) *GlGetUniformLocation {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetUniformLocation pointer is returned so that calls can be chained.
func (a *GlGetUniformLocation) AddWrite(rng memory.Range, id binary.ID) *GlGetUniformLocation {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetUniformLocation) API() gfxapi.API                  { return api{} }
func (c *GlGetUniformLocation) TypeID() atom.TypeID              { return 58 }
func (c *GlGetUniformLocation) Flags() atom.Flags                { return 0 }
func (a *GlGetUniformLocation) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetAttribLocation
////////////////////////////////////////////////////////////////////////////////
type GlGetAttribLocation struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Name         string
	Result       AttributeLocation
}

func (a *GlGetAttribLocation) String() string {
	return fmt.Sprintf("glGetAttribLocation(program: %v, name: %v) → %v", a.Program, a.Name, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetAttribLocation pointer is returned so that calls can be chained.
func (a *GlGetAttribLocation) AddRead(rng memory.Range, id binary.ID) *GlGetAttribLocation {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetAttribLocation pointer is returned so that calls can be chained.
func (a *GlGetAttribLocation) AddWrite(rng memory.Range, id binary.ID) *GlGetAttribLocation {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetAttribLocation) API() gfxapi.API                  { return api{} }
func (c *GlGetAttribLocation) TypeID() atom.TypeID              { return 59 }
func (c *GlGetAttribLocation) Flags() atom.Flags                { return 0 }
func (a *GlGetAttribLocation) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlPixelStorei
////////////////////////////////////////////////////////////////////////////////
type GlPixelStorei struct {
	binary.Generate
	observations atom.Observations
	Parameter    PixelStoreParameter
	Value        int32
}

func (a *GlPixelStorei) String() string {
	return fmt.Sprintf("glPixelStorei(parameter: %v, value: %v)", a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlPixelStorei pointer is returned so that calls can be chained.
func (a *GlPixelStorei) AddRead(rng memory.Range, id binary.ID) *GlPixelStorei {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlPixelStorei pointer is returned so that calls can be chained.
func (a *GlPixelStorei) AddWrite(rng memory.Range, id binary.ID) *GlPixelStorei {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlPixelStorei) API() gfxapi.API                  { return api{} }
func (c *GlPixelStorei) TypeID() atom.TypeID              { return 60 }
func (c *GlPixelStorei) Flags() atom.Flags                { return 0 }
func (a *GlPixelStorei) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexParameteri
////////////////////////////////////////////////////////////////////////////////
type GlTexParameteri struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Parameter    TextureParameter
	Value        int32
}

func (a *GlTexParameteri) String() string {
	return fmt.Sprintf("glTexParameteri(target: %v, parameter: %v, value: %v)", a.Target, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTexParameteri pointer is returned so that calls can be chained.
func (a *GlTexParameteri) AddRead(rng memory.Range, id binary.ID) *GlTexParameteri {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTexParameteri pointer is returned so that calls can be chained.
func (a *GlTexParameteri) AddWrite(rng memory.Range, id binary.ID) *GlTexParameteri {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTexParameteri) API() gfxapi.API                  { return api{} }
func (c *GlTexParameteri) TypeID() atom.TypeID              { return 61 }
func (c *GlTexParameteri) Flags() atom.Flags                { return 0 }
func (a *GlTexParameteri) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexParameterf
////////////////////////////////////////////////////////////////////////////////
type GlTexParameterf struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Parameter    TextureParameter
	Value        float32
}

func (a *GlTexParameterf) String() string {
	return fmt.Sprintf("glTexParameterf(target: %v, parameter: %v, value: %v)", a.Target, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTexParameterf pointer is returned so that calls can be chained.
func (a *GlTexParameterf) AddRead(rng memory.Range, id binary.ID) *GlTexParameterf {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTexParameterf pointer is returned so that calls can be chained.
func (a *GlTexParameterf) AddWrite(rng memory.Range, id binary.ID) *GlTexParameterf {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTexParameterf) API() gfxapi.API                  { return api{} }
func (c *GlTexParameterf) TypeID() atom.TypeID              { return 62 }
func (c *GlTexParameterf) Flags() atom.Flags                { return 0 }
func (a *GlTexParameterf) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameteriv struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Parameter    TextureParameter
	Values       S32ᵖ
}

func (a *GlGetTexParameteriv) String() string {
	return fmt.Sprintf("glGetTexParameteriv(target: %v, parameter: %v, values: %v)", a.Target, a.Parameter, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetTexParameteriv pointer is returned so that calls can be chained.
func (a *GlGetTexParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetTexParameteriv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetTexParameteriv pointer is returned so that calls can be chained.
func (a *GlGetTexParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetTexParameteriv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetTexParameteriv) API() gfxapi.API                  { return api{} }
func (c *GlGetTexParameteriv) TypeID() atom.TypeID              { return 63 }
func (c *GlGetTexParameteriv) Flags() atom.Flags                { return 0 }
func (a *GlGetTexParameteriv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameterfv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameterfv struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Parameter    TextureParameter
	Values       F32ᵖ
}

func (a *GlGetTexParameterfv) String() string {
	return fmt.Sprintf("glGetTexParameterfv(target: %v, parameter: %v, values: %v)", a.Target, a.Parameter, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetTexParameterfv pointer is returned so that calls can be chained.
func (a *GlGetTexParameterfv) AddRead(rng memory.Range, id binary.ID) *GlGetTexParameterfv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetTexParameterfv pointer is returned so that calls can be chained.
func (a *GlGetTexParameterfv) AddWrite(rng memory.Range, id binary.ID) *GlGetTexParameterfv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetTexParameterfv) API() gfxapi.API                  { return api{} }
func (c *GlGetTexParameterfv) TypeID() atom.TypeID              { return 64 }
func (c *GlGetTexParameterfv) Flags() atom.Flags                { return 0 }
func (a *GlGetTexParameterfv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1i
////////////////////////////////////////////////////////////////////////////////
type GlUniform1i struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value        int32
}

func (a *GlUniform1i) String() string {
	return fmt.Sprintf("glUniform1i(location: %v, value: %v)", a.Location, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform1i pointer is returned so that calls can be chained.
func (a *GlUniform1i) AddRead(rng memory.Range, id binary.ID) *GlUniform1i {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform1i pointer is returned so that calls can be chained.
func (a *GlUniform1i) AddWrite(rng memory.Range, id binary.ID) *GlUniform1i {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform1i) API() gfxapi.API                  { return api{} }
func (c *GlUniform1i) TypeID() atom.TypeID              { return 65 }
func (c *GlUniform1i) Flags() atom.Flags                { return 0 }
func (a *GlUniform1i) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2i
////////////////////////////////////////////////////////////////////////////////
type GlUniform2i struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value0       int32
	Value1       int32
}

func (a *GlUniform2i) String() string {
	return fmt.Sprintf("glUniform2i(location: %v, value0: %v, value1: %v)", a.Location, a.Value0, a.Value1)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform2i pointer is returned so that calls can be chained.
func (a *GlUniform2i) AddRead(rng memory.Range, id binary.ID) *GlUniform2i {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform2i pointer is returned so that calls can be chained.
func (a *GlUniform2i) AddWrite(rng memory.Range, id binary.ID) *GlUniform2i {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform2i) API() gfxapi.API                  { return api{} }
func (c *GlUniform2i) TypeID() atom.TypeID              { return 66 }
func (c *GlUniform2i) Flags() atom.Flags                { return 0 }
func (a *GlUniform2i) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3i
////////////////////////////////////////////////////////////////////////////////
type GlUniform3i struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value0       int32
	Value1       int32
	Value2       int32
}

func (a *GlUniform3i) String() string {
	return fmt.Sprintf("glUniform3i(location: %v, value0: %v, value1: %v, value2: %v)", a.Location, a.Value0, a.Value1, a.Value2)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform3i pointer is returned so that calls can be chained.
func (a *GlUniform3i) AddRead(rng memory.Range, id binary.ID) *GlUniform3i {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform3i pointer is returned so that calls can be chained.
func (a *GlUniform3i) AddWrite(rng memory.Range, id binary.ID) *GlUniform3i {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform3i) API() gfxapi.API                  { return api{} }
func (c *GlUniform3i) TypeID() atom.TypeID              { return 67 }
func (c *GlUniform3i) Flags() atom.Flags                { return 0 }
func (a *GlUniform3i) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4i
////////////////////////////////////////////////////////////////////////////////
type GlUniform4i struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value0       int32
	Value1       int32
	Value2       int32
	Value3       int32
}

func (a *GlUniform4i) String() string {
	return fmt.Sprintf("glUniform4i(location: %v, value0: %v, value1: %v, value2: %v, value3: %v)", a.Location, a.Value0, a.Value1, a.Value2, a.Value3)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform4i pointer is returned so that calls can be chained.
func (a *GlUniform4i) AddRead(rng memory.Range, id binary.ID) *GlUniform4i {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform4i pointer is returned so that calls can be chained.
func (a *GlUniform4i) AddWrite(rng memory.Range, id binary.ID) *GlUniform4i {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform4i) API() gfxapi.API                  { return api{} }
func (c *GlUniform4i) TypeID() atom.TypeID              { return 68 }
func (c *GlUniform4i) Flags() atom.Flags                { return 0 }
func (a *GlUniform4i) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform1iv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        S32ᵖ
}

func (a *GlUniform1iv) String() string {
	return fmt.Sprintf("glUniform1iv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform1iv pointer is returned so that calls can be chained.
func (a *GlUniform1iv) AddRead(rng memory.Range, id binary.ID) *GlUniform1iv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform1iv pointer is returned so that calls can be chained.
func (a *GlUniform1iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform1iv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform1iv) API() gfxapi.API                  { return api{} }
func (c *GlUniform1iv) TypeID() atom.TypeID              { return 69 }
func (c *GlUniform1iv) Flags() atom.Flags                { return 0 }
func (a *GlUniform1iv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform2iv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        S32ᵖ
}

func (a *GlUniform2iv) String() string {
	return fmt.Sprintf("glUniform2iv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform2iv pointer is returned so that calls can be chained.
func (a *GlUniform2iv) AddRead(rng memory.Range, id binary.ID) *GlUniform2iv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform2iv pointer is returned so that calls can be chained.
func (a *GlUniform2iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform2iv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform2iv) API() gfxapi.API                  { return api{} }
func (c *GlUniform2iv) TypeID() atom.TypeID              { return 70 }
func (c *GlUniform2iv) Flags() atom.Flags                { return 0 }
func (a *GlUniform2iv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform3iv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        S32ᵖ
}

func (a *GlUniform3iv) String() string {
	return fmt.Sprintf("glUniform3iv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform3iv pointer is returned so that calls can be chained.
func (a *GlUniform3iv) AddRead(rng memory.Range, id binary.ID) *GlUniform3iv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform3iv pointer is returned so that calls can be chained.
func (a *GlUniform3iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform3iv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform3iv) API() gfxapi.API                  { return api{} }
func (c *GlUniform3iv) TypeID() atom.TypeID              { return 71 }
func (c *GlUniform3iv) Flags() atom.Flags                { return 0 }
func (a *GlUniform3iv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform4iv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        S32ᵖ
}

func (a *GlUniform4iv) String() string {
	return fmt.Sprintf("glUniform4iv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform4iv pointer is returned so that calls can be chained.
func (a *GlUniform4iv) AddRead(rng memory.Range, id binary.ID) *GlUniform4iv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform4iv pointer is returned so that calls can be chained.
func (a *GlUniform4iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform4iv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform4iv) API() gfxapi.API                  { return api{} }
func (c *GlUniform4iv) TypeID() atom.TypeID              { return 72 }
func (c *GlUniform4iv) Flags() atom.Flags                { return 0 }
func (a *GlUniform4iv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1f
////////////////////////////////////////////////////////////////////////////////
type GlUniform1f struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value        float32
}

func (a *GlUniform1f) String() string {
	return fmt.Sprintf("glUniform1f(location: %v, value: %v)", a.Location, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform1f pointer is returned so that calls can be chained.
func (a *GlUniform1f) AddRead(rng memory.Range, id binary.ID) *GlUniform1f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform1f pointer is returned so that calls can be chained.
func (a *GlUniform1f) AddWrite(rng memory.Range, id binary.ID) *GlUniform1f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform1f) API() gfxapi.API                  { return api{} }
func (c *GlUniform1f) TypeID() atom.TypeID              { return 73 }
func (c *GlUniform1f) Flags() atom.Flags                { return 0 }
func (a *GlUniform1f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2f
////////////////////////////////////////////////////////////////////////////////
type GlUniform2f struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value0       float32
	Value1       float32
}

func (a *GlUniform2f) String() string {
	return fmt.Sprintf("glUniform2f(location: %v, value0: %v, value1: %v)", a.Location, a.Value0, a.Value1)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform2f pointer is returned so that calls can be chained.
func (a *GlUniform2f) AddRead(rng memory.Range, id binary.ID) *GlUniform2f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform2f pointer is returned so that calls can be chained.
func (a *GlUniform2f) AddWrite(rng memory.Range, id binary.ID) *GlUniform2f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform2f) API() gfxapi.API                  { return api{} }
func (c *GlUniform2f) TypeID() atom.TypeID              { return 74 }
func (c *GlUniform2f) Flags() atom.Flags                { return 0 }
func (a *GlUniform2f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3f
////////////////////////////////////////////////////////////////////////////////
type GlUniform3f struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value0       float32
	Value1       float32
	Value2       float32
}

func (a *GlUniform3f) String() string {
	return fmt.Sprintf("glUniform3f(location: %v, value0: %v, value1: %v, value2: %v)", a.Location, a.Value0, a.Value1, a.Value2)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform3f pointer is returned so that calls can be chained.
func (a *GlUniform3f) AddRead(rng memory.Range, id binary.ID) *GlUniform3f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform3f pointer is returned so that calls can be chained.
func (a *GlUniform3f) AddWrite(rng memory.Range, id binary.ID) *GlUniform3f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform3f) API() gfxapi.API                  { return api{} }
func (c *GlUniform3f) TypeID() atom.TypeID              { return 75 }
func (c *GlUniform3f) Flags() atom.Flags                { return 0 }
func (a *GlUniform3f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4f
////////////////////////////////////////////////////////////////////////////////
type GlUniform4f struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Value0       float32
	Value1       float32
	Value2       float32
	Value3       float32
}

func (a *GlUniform4f) String() string {
	return fmt.Sprintf("glUniform4f(location: %v, value0: %v, value1: %v, value2: %v, value3: %v)", a.Location, a.Value0, a.Value1, a.Value2, a.Value3)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform4f pointer is returned so that calls can be chained.
func (a *GlUniform4f) AddRead(rng memory.Range, id binary.ID) *GlUniform4f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform4f pointer is returned so that calls can be chained.
func (a *GlUniform4f) AddWrite(rng memory.Range, id binary.ID) *GlUniform4f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform4f) API() gfxapi.API                  { return api{} }
func (c *GlUniform4f) TypeID() atom.TypeID              { return 76 }
func (c *GlUniform4f) Flags() atom.Flags                { return 0 }
func (a *GlUniform4f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform1fv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        F32ᵖ
}

func (a *GlUniform1fv) String() string {
	return fmt.Sprintf("glUniform1fv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform1fv pointer is returned so that calls can be chained.
func (a *GlUniform1fv) AddRead(rng memory.Range, id binary.ID) *GlUniform1fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform1fv pointer is returned so that calls can be chained.
func (a *GlUniform1fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform1fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform1fv) API() gfxapi.API                  { return api{} }
func (c *GlUniform1fv) TypeID() atom.TypeID              { return 77 }
func (c *GlUniform1fv) Flags() atom.Flags                { return 0 }
func (a *GlUniform1fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform2fv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        F32ᵖ
}

func (a *GlUniform2fv) String() string {
	return fmt.Sprintf("glUniform2fv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform2fv pointer is returned so that calls can be chained.
func (a *GlUniform2fv) AddRead(rng memory.Range, id binary.ID) *GlUniform2fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform2fv pointer is returned so that calls can be chained.
func (a *GlUniform2fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform2fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform2fv) API() gfxapi.API                  { return api{} }
func (c *GlUniform2fv) TypeID() atom.TypeID              { return 78 }
func (c *GlUniform2fv) Flags() atom.Flags                { return 0 }
func (a *GlUniform2fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform3fv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        F32ᵖ
}

func (a *GlUniform3fv) String() string {
	return fmt.Sprintf("glUniform3fv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform3fv pointer is returned so that calls can be chained.
func (a *GlUniform3fv) AddRead(rng memory.Range, id binary.ID) *GlUniform3fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform3fv pointer is returned so that calls can be chained.
func (a *GlUniform3fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform3fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform3fv) API() gfxapi.API                  { return api{} }
func (c *GlUniform3fv) TypeID() atom.TypeID              { return 79 }
func (c *GlUniform3fv) Flags() atom.Flags                { return 0 }
func (a *GlUniform3fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform4fv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Value        F32ᵖ
}

func (a *GlUniform4fv) String() string {
	return fmt.Sprintf("glUniform4fv(location: %v, count: %v, value: %v)", a.Location, a.Count, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniform4fv pointer is returned so that calls can be chained.
func (a *GlUniform4fv) AddRead(rng memory.Range, id binary.ID) *GlUniform4fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniform4fv pointer is returned so that calls can be chained.
func (a *GlUniform4fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform4fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniform4fv) API() gfxapi.API                  { return api{} }
func (c *GlUniform4fv) TypeID() atom.TypeID              { return 80 }
func (c *GlUniform4fv) Flags() atom.Flags                { return 0 }
func (a *GlUniform4fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix2fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix2fv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Transpose    bool
	Values       F32ᵖ
}

func (a *GlUniformMatrix2fv) String() string {
	return fmt.Sprintf("glUniformMatrix2fv(location: %v, count: %v, transpose: %v, values: %v)", a.Location, a.Count, a.Transpose, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniformMatrix2fv pointer is returned so that calls can be chained.
func (a *GlUniformMatrix2fv) AddRead(rng memory.Range, id binary.ID) *GlUniformMatrix2fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniformMatrix2fv pointer is returned so that calls can be chained.
func (a *GlUniformMatrix2fv) AddWrite(rng memory.Range, id binary.ID) *GlUniformMatrix2fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniformMatrix2fv) API() gfxapi.API                  { return api{} }
func (c *GlUniformMatrix2fv) TypeID() atom.TypeID              { return 81 }
func (c *GlUniformMatrix2fv) Flags() atom.Flags                { return 0 }
func (a *GlUniformMatrix2fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix3fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix3fv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Transpose    bool
	Values       F32ᵖ
}

func (a *GlUniformMatrix3fv) String() string {
	return fmt.Sprintf("glUniformMatrix3fv(location: %v, count: %v, transpose: %v, values: %v)", a.Location, a.Count, a.Transpose, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniformMatrix3fv pointer is returned so that calls can be chained.
func (a *GlUniformMatrix3fv) AddRead(rng memory.Range, id binary.ID) *GlUniformMatrix3fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniformMatrix3fv pointer is returned so that calls can be chained.
func (a *GlUniformMatrix3fv) AddWrite(rng memory.Range, id binary.ID) *GlUniformMatrix3fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniformMatrix3fv) API() gfxapi.API                  { return api{} }
func (c *GlUniformMatrix3fv) TypeID() atom.TypeID              { return 82 }
func (c *GlUniformMatrix3fv) Flags() atom.Flags                { return 0 }
func (a *GlUniformMatrix3fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix4fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix4fv struct {
	binary.Generate
	observations atom.Observations
	Location     UniformLocation
	Count        int32
	Transpose    bool
	Values       F32ᵖ
}

func (a *GlUniformMatrix4fv) String() string {
	return fmt.Sprintf("glUniformMatrix4fv(location: %v, count: %v, transpose: %v, values: %v)", a.Location, a.Count, a.Transpose, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniformMatrix4fv pointer is returned so that calls can be chained.
func (a *GlUniformMatrix4fv) AddRead(rng memory.Range, id binary.ID) *GlUniformMatrix4fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniformMatrix4fv pointer is returned so that calls can be chained.
func (a *GlUniformMatrix4fv) AddWrite(rng memory.Range, id binary.ID) *GlUniformMatrix4fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniformMatrix4fv) API() gfxapi.API                  { return api{} }
func (c *GlUniformMatrix4fv) TypeID() atom.TypeID              { return 83 }
func (c *GlUniformMatrix4fv) Flags() atom.Flags                { return 0 }
func (a *GlUniformMatrix4fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformfv
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformfv struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Location     UniformLocation
	Values       F32ᵖ
}

func (a *GlGetUniformfv) String() string {
	return fmt.Sprintf("glGetUniformfv(program: %v, location: %v, values: %v)", a.Program, a.Location, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetUniformfv pointer is returned so that calls can be chained.
func (a *GlGetUniformfv) AddRead(rng memory.Range, id binary.ID) *GlGetUniformfv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetUniformfv pointer is returned so that calls can be chained.
func (a *GlGetUniformfv) AddWrite(rng memory.Range, id binary.ID) *GlGetUniformfv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetUniformfv) API() gfxapi.API                  { return api{} }
func (c *GlGetUniformfv) TypeID() atom.TypeID              { return 84 }
func (c *GlGetUniformfv) Flags() atom.Flags                { return 0 }
func (a *GlGetUniformfv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformiv
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformiv struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Location     UniformLocation
	Values       S32ᵖ
}

func (a *GlGetUniformiv) String() string {
	return fmt.Sprintf("glGetUniformiv(program: %v, location: %v, values: %v)", a.Program, a.Location, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetUniformiv pointer is returned so that calls can be chained.
func (a *GlGetUniformiv) AddRead(rng memory.Range, id binary.ID) *GlGetUniformiv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetUniformiv pointer is returned so that calls can be chained.
func (a *GlGetUniformiv) AddWrite(rng memory.Range, id binary.ID) *GlGetUniformiv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetUniformiv) API() gfxapi.API                  { return api{} }
func (c *GlGetUniformiv) TypeID() atom.TypeID              { return 85 }
func (c *GlGetUniformiv) Flags() atom.Flags                { return 0 }
func (a *GlGetUniformiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib1f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib1f struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value0       float32
}

func (a *GlVertexAttrib1f) String() string {
	return fmt.Sprintf("glVertexAttrib1f(location: %v, value0: %v)", a.Location, a.Value0)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib1f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib1f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib1f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib1f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib1f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib1f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib1f) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib1f) TypeID() atom.TypeID              { return 86 }
func (c *GlVertexAttrib1f) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib1f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib2f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib2f struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value0       float32
	Value1       float32
}

func (a *GlVertexAttrib2f) String() string {
	return fmt.Sprintf("glVertexAttrib2f(location: %v, value0: %v, value1: %v)", a.Location, a.Value0, a.Value1)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib2f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib2f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib2f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib2f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib2f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib2f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib2f) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib2f) TypeID() atom.TypeID              { return 87 }
func (c *GlVertexAttrib2f) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib2f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib3f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib3f struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value0       float32
	Value1       float32
	Value2       float32
}

func (a *GlVertexAttrib3f) String() string {
	return fmt.Sprintf("glVertexAttrib3f(location: %v, value0: %v, value1: %v, value2: %v)", a.Location, a.Value0, a.Value1, a.Value2)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib3f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib3f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib3f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib3f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib3f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib3f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib3f) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib3f) TypeID() atom.TypeID              { return 88 }
func (c *GlVertexAttrib3f) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib3f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib4f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib4f struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value0       float32
	Value1       float32
	Value2       float32
	Value3       float32
}

func (a *GlVertexAttrib4f) String() string {
	return fmt.Sprintf("glVertexAttrib4f(location: %v, value0: %v, value1: %v, value2: %v, value3: %v)", a.Location, a.Value0, a.Value1, a.Value2, a.Value3)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib4f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib4f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib4f {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib4f pointer is returned so that calls can be chained.
func (a *GlVertexAttrib4f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib4f {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib4f) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib4f) TypeID() atom.TypeID              { return 89 }
func (c *GlVertexAttrib4f) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib4f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib1fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib1fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᵖ
}

func (a *GlVertexAttrib1fv) String() string {
	return fmt.Sprintf("glVertexAttrib1fv(location: %v, value: %v)", a.Location, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib1fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib1fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib1fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib1fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib1fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib1fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib1fv) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib1fv) TypeID() atom.TypeID              { return 90 }
func (c *GlVertexAttrib1fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib1fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib2fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib2fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᵖ
}

func (a *GlVertexAttrib2fv) String() string {
	return fmt.Sprintf("glVertexAttrib2fv(location: %v, value: %v)", a.Location, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib2fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib2fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib2fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib2fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib2fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib2fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib2fv) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib2fv) TypeID() atom.TypeID              { return 91 }
func (c *GlVertexAttrib2fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib2fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib3fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib3fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᵖ
}

func (a *GlVertexAttrib3fv) String() string {
	return fmt.Sprintf("glVertexAttrib3fv(location: %v, value: %v)", a.Location, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib3fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib3fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib3fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib3fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib3fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib3fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib3fv) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib3fv) TypeID() atom.TypeID              { return 92 }
func (c *GlVertexAttrib3fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib3fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib4fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib4fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᵖ
}

func (a *GlVertexAttrib4fv) String() string {
	return fmt.Sprintf("glVertexAttrib4fv(location: %v, value: %v)", a.Location, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib4fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib4fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib4fv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlVertexAttrib4fv pointer is returned so that calls can be chained.
func (a *GlVertexAttrib4fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib4fv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlVertexAttrib4fv) API() gfxapi.API                  { return api{} }
func (c *GlVertexAttrib4fv) TypeID() atom.TypeID              { return 93 }
func (c *GlVertexAttrib4fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib4fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderPrecisionFormat
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderPrecisionFormat struct {
	binary.Generate
	observations  atom.Observations
	ShaderType    ShaderType
	PrecisionType PrecisionType
	Range         S32ᵖ
	Precision     S32ᵖ
}

func (a *GlGetShaderPrecisionFormat) String() string {
	return fmt.Sprintf("glGetShaderPrecisionFormat(shader_type: %v, precision_type: %v, range: %v, precision: %v)", a.ShaderType, a.PrecisionType, a.Range, a.Precision)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetShaderPrecisionFormat pointer is returned so that calls can be chained.
func (a *GlGetShaderPrecisionFormat) AddRead(rng memory.Range, id binary.ID) *GlGetShaderPrecisionFormat {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetShaderPrecisionFormat pointer is returned so that calls can be chained.
func (a *GlGetShaderPrecisionFormat) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderPrecisionFormat {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetShaderPrecisionFormat) API() gfxapi.API                  { return api{} }
func (c *GlGetShaderPrecisionFormat) TypeID() atom.TypeID              { return 94 }
func (c *GlGetShaderPrecisionFormat) Flags() atom.Flags                { return 0 }
func (a *GlGetShaderPrecisionFormat) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDepthMask
////////////////////////////////////////////////////////////////////////////////
type GlDepthMask struct {
	binary.Generate
	observations atom.Observations
	Enabled      bool
}

func (a *GlDepthMask) String() string {
	return fmt.Sprintf("glDepthMask(enabled: %v)", a.Enabled)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDepthMask pointer is returned so that calls can be chained.
func (a *GlDepthMask) AddRead(rng memory.Range, id binary.ID) *GlDepthMask {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDepthMask pointer is returned so that calls can be chained.
func (a *GlDepthMask) AddWrite(rng memory.Range, id binary.ID) *GlDepthMask {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDepthMask) API() gfxapi.API                  { return api{} }
func (c *GlDepthMask) TypeID() atom.TypeID              { return 95 }
func (c *GlDepthMask) Flags() atom.Flags                { return 0 }
func (a *GlDepthMask) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDepthFunc
////////////////////////////////////////////////////////////////////////////////
type GlDepthFunc struct {
	binary.Generate
	observations atom.Observations
	Function     TestFunction
}

func (a *GlDepthFunc) String() string {
	return fmt.Sprintf("glDepthFunc(function: %v)", a.Function)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDepthFunc pointer is returned so that calls can be chained.
func (a *GlDepthFunc) AddRead(rng memory.Range, id binary.ID) *GlDepthFunc {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDepthFunc pointer is returned so that calls can be chained.
func (a *GlDepthFunc) AddWrite(rng memory.Range, id binary.ID) *GlDepthFunc {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDepthFunc) API() gfxapi.API                  { return api{} }
func (c *GlDepthFunc) TypeID() atom.TypeID              { return 96 }
func (c *GlDepthFunc) Flags() atom.Flags                { return 0 }
func (a *GlDepthFunc) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDepthRangef
////////////////////////////////////////////////////////////////////////////////
type GlDepthRangef struct {
	binary.Generate
	observations atom.Observations
	Near         float32
	Far          float32
}

func (a *GlDepthRangef) String() string {
	return fmt.Sprintf("glDepthRangef(near: %v, far: %v)", a.Near, a.Far)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDepthRangef pointer is returned so that calls can be chained.
func (a *GlDepthRangef) AddRead(rng memory.Range, id binary.ID) *GlDepthRangef {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDepthRangef pointer is returned so that calls can be chained.
func (a *GlDepthRangef) AddWrite(rng memory.Range, id binary.ID) *GlDepthRangef {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDepthRangef) API() gfxapi.API                  { return api{} }
func (c *GlDepthRangef) TypeID() atom.TypeID              { return 97 }
func (c *GlDepthRangef) Flags() atom.Flags                { return 0 }
func (a *GlDepthRangef) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlColorMask
////////////////////////////////////////////////////////////////////////////////
type GlColorMask struct {
	binary.Generate
	observations atom.Observations
	Red          bool
	Green        bool
	Blue         bool
	Alpha        bool
}

func (a *GlColorMask) String() string {
	return fmt.Sprintf("glColorMask(red: %v, green: %v, blue: %v, alpha: %v)", a.Red, a.Green, a.Blue, a.Alpha)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlColorMask pointer is returned so that calls can be chained.
func (a *GlColorMask) AddRead(rng memory.Range, id binary.ID) *GlColorMask {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlColorMask pointer is returned so that calls can be chained.
func (a *GlColorMask) AddWrite(rng memory.Range, id binary.ID) *GlColorMask {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlColorMask) API() gfxapi.API                  { return api{} }
func (c *GlColorMask) TypeID() atom.TypeID              { return 98 }
func (c *GlColorMask) Flags() atom.Flags                { return 0 }
func (a *GlColorMask) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStencilMask
////////////////////////////////////////////////////////////////////////////////
type GlStencilMask struct {
	binary.Generate
	observations atom.Observations
	Mask         uint32
}

func (a *GlStencilMask) String() string {
	return fmt.Sprintf("glStencilMask(mask: %v)", a.Mask)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlStencilMask pointer is returned so that calls can be chained.
func (a *GlStencilMask) AddRead(rng memory.Range, id binary.ID) *GlStencilMask {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlStencilMask pointer is returned so that calls can be chained.
func (a *GlStencilMask) AddWrite(rng memory.Range, id binary.ID) *GlStencilMask {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlStencilMask) API() gfxapi.API                  { return api{} }
func (c *GlStencilMask) TypeID() atom.TypeID              { return 99 }
func (c *GlStencilMask) Flags() atom.Flags                { return 0 }
func (a *GlStencilMask) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStencilMaskSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilMaskSeparate struct {
	binary.Generate
	observations atom.Observations
	Face         FaceMode
	Mask         uint32
}

func (a *GlStencilMaskSeparate) String() string {
	return fmt.Sprintf("glStencilMaskSeparate(face: %v, mask: %v)", a.Face, a.Mask)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlStencilMaskSeparate pointer is returned so that calls can be chained.
func (a *GlStencilMaskSeparate) AddRead(rng memory.Range, id binary.ID) *GlStencilMaskSeparate {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlStencilMaskSeparate pointer is returned so that calls can be chained.
func (a *GlStencilMaskSeparate) AddWrite(rng memory.Range, id binary.ID) *GlStencilMaskSeparate {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlStencilMaskSeparate) API() gfxapi.API                  { return api{} }
func (c *GlStencilMaskSeparate) TypeID() atom.TypeID              { return 100 }
func (c *GlStencilMaskSeparate) Flags() atom.Flags                { return 0 }
func (a *GlStencilMaskSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStencilFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilFuncSeparate struct {
	binary.Generate
	observations   atom.Observations
	Face           FaceMode
	Function       TestFunction
	ReferenceValue int32
	Mask           int32
}

func (a *GlStencilFuncSeparate) String() string {
	return fmt.Sprintf("glStencilFuncSeparate(face: %v, function: %v, reference_value: %v, mask: %v)", a.Face, a.Function, a.ReferenceValue, a.Mask)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlStencilFuncSeparate pointer is returned so that calls can be chained.
func (a *GlStencilFuncSeparate) AddRead(rng memory.Range, id binary.ID) *GlStencilFuncSeparate {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlStencilFuncSeparate pointer is returned so that calls can be chained.
func (a *GlStencilFuncSeparate) AddWrite(rng memory.Range, id binary.ID) *GlStencilFuncSeparate {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlStencilFuncSeparate) API() gfxapi.API                  { return api{} }
func (c *GlStencilFuncSeparate) TypeID() atom.TypeID              { return 101 }
func (c *GlStencilFuncSeparate) Flags() atom.Flags                { return 0 }
func (a *GlStencilFuncSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStencilOpSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilOpSeparate struct {
	binary.Generate
	observations         atom.Observations
	Face                 FaceMode
	StencilFail          StencilAction
	StencilPassDepthFail StencilAction
	StencilPassDepthPass StencilAction
}

func (a *GlStencilOpSeparate) String() string {
	return fmt.Sprintf("glStencilOpSeparate(face: %v, stencil_fail: %v, stencil_pass_depth_fail: %v, stencil_pass_depth_pass: %v)", a.Face, a.StencilFail, a.StencilPassDepthFail, a.StencilPassDepthPass)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlStencilOpSeparate pointer is returned so that calls can be chained.
func (a *GlStencilOpSeparate) AddRead(rng memory.Range, id binary.ID) *GlStencilOpSeparate {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlStencilOpSeparate pointer is returned so that calls can be chained.
func (a *GlStencilOpSeparate) AddWrite(rng memory.Range, id binary.ID) *GlStencilOpSeparate {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlStencilOpSeparate) API() gfxapi.API                  { return api{} }
func (c *GlStencilOpSeparate) TypeID() atom.TypeID              { return 102 }
func (c *GlStencilOpSeparate) Flags() atom.Flags                { return 0 }
func (a *GlStencilOpSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFrontFace
////////////////////////////////////////////////////////////////////////////////
type GlFrontFace struct {
	binary.Generate
	observations atom.Observations
	Orientation  FaceOrientation
}

func (a *GlFrontFace) String() string {
	return fmt.Sprintf("glFrontFace(orientation: %v)", a.Orientation)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlFrontFace pointer is returned so that calls can be chained.
func (a *GlFrontFace) AddRead(rng memory.Range, id binary.ID) *GlFrontFace {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlFrontFace pointer is returned so that calls can be chained.
func (a *GlFrontFace) AddWrite(rng memory.Range, id binary.ID) *GlFrontFace {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlFrontFace) API() gfxapi.API                  { return api{} }
func (c *GlFrontFace) TypeID() atom.TypeID              { return 103 }
func (c *GlFrontFace) Flags() atom.Flags                { return 0 }
func (a *GlFrontFace) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlViewport
////////////////////////////////////////////////////////////////////////////////
type GlViewport struct {
	binary.Generate
	observations atom.Observations
	X            int32
	Y            int32
	Width        int32
	Height       int32
}

func (a *GlViewport) String() string {
	return fmt.Sprintf("glViewport(x: %v, y: %v, width: %v, height: %v)", a.X, a.Y, a.Width, a.Height)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlViewport pointer is returned so that calls can be chained.
func (a *GlViewport) AddRead(rng memory.Range, id binary.ID) *GlViewport {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlViewport pointer is returned so that calls can be chained.
func (a *GlViewport) AddWrite(rng memory.Range, id binary.ID) *GlViewport {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlViewport) API() gfxapi.API                  { return api{} }
func (c *GlViewport) TypeID() atom.TypeID              { return 104 }
func (c *GlViewport) Flags() atom.Flags                { return 0 }
func (a *GlViewport) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlScissor
////////////////////////////////////////////////////////////////////////////////
type GlScissor struct {
	binary.Generate
	observations atom.Observations
	X            int32
	Y            int32
	Width        int32
	Height       int32
}

func (a *GlScissor) String() string {
	return fmt.Sprintf("glScissor(x: %v, y: %v, width: %v, height: %v)", a.X, a.Y, a.Width, a.Height)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlScissor pointer is returned so that calls can be chained.
func (a *GlScissor) AddRead(rng memory.Range, id binary.ID) *GlScissor {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlScissor pointer is returned so that calls can be chained.
func (a *GlScissor) AddWrite(rng memory.Range, id binary.ID) *GlScissor {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlScissor) API() gfxapi.API                  { return api{} }
func (c *GlScissor) TypeID() atom.TypeID              { return 105 }
func (c *GlScissor) Flags() atom.Flags                { return 0 }
func (a *GlScissor) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlActiveTexture
////////////////////////////////////////////////////////////////////////////////
type GlActiveTexture struct {
	binary.Generate
	observations atom.Observations
	Unit         TextureUnit
}

func (a *GlActiveTexture) String() string {
	return fmt.Sprintf("glActiveTexture(unit: %v)", a.Unit)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlActiveTexture pointer is returned so that calls can be chained.
func (a *GlActiveTexture) AddRead(rng memory.Range, id binary.ID) *GlActiveTexture {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlActiveTexture pointer is returned so that calls can be chained.
func (a *GlActiveTexture) AddWrite(rng memory.Range, id binary.ID) *GlActiveTexture {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlActiveTexture) API() gfxapi.API                  { return api{} }
func (c *GlActiveTexture) TypeID() atom.TypeID              { return 106 }
func (c *GlActiveTexture) Flags() atom.Flags                { return 0 }
func (a *GlActiveTexture) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenTextures
////////////////////////////////////////////////////////////////////////////////
type GlGenTextures struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Textures     TextureIdᵖ
}

func (a *GlGenTextures) String() string {
	return fmt.Sprintf("glGenTextures(count: %v, textures: %v)", a.Count, a.Textures)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenTextures pointer is returned so that calls can be chained.
func (a *GlGenTextures) AddRead(rng memory.Range, id binary.ID) *GlGenTextures {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenTextures pointer is returned so that calls can be chained.
func (a *GlGenTextures) AddWrite(rng memory.Range, id binary.ID) *GlGenTextures {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenTextures) API() gfxapi.API                  { return api{} }
func (c *GlGenTextures) TypeID() atom.TypeID              { return 107 }
func (c *GlGenTextures) Flags() atom.Flags                { return 0 }
func (a *GlGenTextures) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteTextures
////////////////////////////////////////////////////////////////////////////////
type GlDeleteTextures struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Textures     TextureIdᵖ
}

func (a *GlDeleteTextures) String() string {
	return fmt.Sprintf("glDeleteTextures(count: %v, textures: %v)", a.Count, a.Textures)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteTextures pointer is returned so that calls can be chained.
func (a *GlDeleteTextures) AddRead(rng memory.Range, id binary.ID) *GlDeleteTextures {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteTextures pointer is returned so that calls can be chained.
func (a *GlDeleteTextures) AddWrite(rng memory.Range, id binary.ID) *GlDeleteTextures {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteTextures) API() gfxapi.API                  { return api{} }
func (c *GlDeleteTextures) TypeID() atom.TypeID              { return 108 }
func (c *GlDeleteTextures) Flags() atom.Flags                { return 0 }
func (a *GlDeleteTextures) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsTexture
////////////////////////////////////////////////////////////////////////////////
type GlIsTexture struct {
	binary.Generate
	observations atom.Observations
	Texture      TextureId
	Result       bool
}

func (a *GlIsTexture) String() string {
	return fmt.Sprintf("glIsTexture(texture: %v) → %v", a.Texture, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsTexture pointer is returned so that calls can be chained.
func (a *GlIsTexture) AddRead(rng memory.Range, id binary.ID) *GlIsTexture {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsTexture pointer is returned so that calls can be chained.
func (a *GlIsTexture) AddWrite(rng memory.Range, id binary.ID) *GlIsTexture {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsTexture) API() gfxapi.API                  { return api{} }
func (c *GlIsTexture) TypeID() atom.TypeID              { return 109 }
func (c *GlIsTexture) Flags() atom.Flags                { return 0 }
func (a *GlIsTexture) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindTexture
////////////////////////////////////////////////////////////////////////////////
type GlBindTexture struct {
	binary.Generate
	observations atom.Observations
	Target       TextureTarget
	Texture      TextureId
}

func (a *GlBindTexture) String() string {
	return fmt.Sprintf("glBindTexture(target: %v, texture: %v)", a.Target, a.Texture)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindTexture pointer is returned so that calls can be chained.
func (a *GlBindTexture) AddRead(rng memory.Range, id binary.ID) *GlBindTexture {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindTexture pointer is returned so that calls can be chained.
func (a *GlBindTexture) AddWrite(rng memory.Range, id binary.ID) *GlBindTexture {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindTexture) API() gfxapi.API                  { return api{} }
func (c *GlBindTexture) TypeID() atom.TypeID              { return 110 }
func (c *GlBindTexture) Flags() atom.Flags                { return 0 }
func (a *GlBindTexture) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexImage2D struct {
	binary.Generate
	observations   atom.Observations
	Target         TextureImageTarget
	Level          int32
	InternalFormat TexelFormat
	Width          int32
	Height         int32
	Border         int32
	Format         TexelFormat
	Type           TexelType
	Data           TexturePointer
}

func (a *GlTexImage2D) String() string {
	return fmt.Sprintf("glTexImage2D(target: %v, level: %v, internal_format: %v, width: %v, height: %v, border: %v, format: %v, type: %v, data: %v)", a.Target, a.Level, a.InternalFormat, a.Width, a.Height, a.Border, a.Format, a.Type, a.Data)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTexImage2D pointer is returned so that calls can be chained.
func (a *GlTexImage2D) AddRead(rng memory.Range, id binary.ID) *GlTexImage2D {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTexImage2D pointer is returned so that calls can be chained.
func (a *GlTexImage2D) AddWrite(rng memory.Range, id binary.ID) *GlTexImage2D {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTexImage2D) API() gfxapi.API                  { return api{} }
func (c *GlTexImage2D) TypeID() atom.TypeID              { return 111 }
func (c *GlTexImage2D) Flags() atom.Flags                { return 0 }
func (a *GlTexImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexSubImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       TextureImageTarget
	Level        int32
	Xoffset      int32
	Yoffset      int32
	Width        int32
	Height       int32
	Format       TexelFormat
	Type         TexelType
	Data         TexturePointer
}

func (a *GlTexSubImage2D) String() string {
	return fmt.Sprintf("glTexSubImage2D(target: %v, level: %v, xoffset: %v, yoffset: %v, width: %v, height: %v, format: %v, type: %v, data: %v)", a.Target, a.Level, a.Xoffset, a.Yoffset, a.Width, a.Height, a.Format, a.Type, a.Data)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlTexSubImage2D pointer is returned so that calls can be chained.
func (a *GlTexSubImage2D) AddRead(rng memory.Range, id binary.ID) *GlTexSubImage2D {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlTexSubImage2D pointer is returned so that calls can be chained.
func (a *GlTexSubImage2D) AddWrite(rng memory.Range, id binary.ID) *GlTexSubImage2D {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlTexSubImage2D) API() gfxapi.API                  { return api{} }
func (c *GlTexSubImage2D) TypeID() atom.TypeID              { return 112 }
func (c *GlTexSubImage2D) Flags() atom.Flags                { return 0 }
func (a *GlTexSubImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       TextureImageTarget
	Level        int32
	Format       TexelFormat
	X            int32
	Y            int32
	Width        int32
	Height       int32
	Border       int32
}

func (a *GlCopyTexImage2D) String() string {
	return fmt.Sprintf("glCopyTexImage2D(target: %v, level: %v, format: %v, x: %v, y: %v, width: %v, height: %v, border: %v)", a.Target, a.Level, a.Format, a.X, a.Y, a.Width, a.Height, a.Border)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCopyTexImage2D pointer is returned so that calls can be chained.
func (a *GlCopyTexImage2D) AddRead(rng memory.Range, id binary.ID) *GlCopyTexImage2D {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCopyTexImage2D pointer is returned so that calls can be chained.
func (a *GlCopyTexImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCopyTexImage2D {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCopyTexImage2D) API() gfxapi.API                  { return api{} }
func (c *GlCopyTexImage2D) TypeID() atom.TypeID              { return 113 }
func (c *GlCopyTexImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCopyTexImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexSubImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       TextureImageTarget
	Level        int32
	Xoffset      int32
	Yoffset      int32
	X            int32
	Y            int32
	Width        int32
	Height       int32
}

func (a *GlCopyTexSubImage2D) String() string {
	return fmt.Sprintf("glCopyTexSubImage2D(target: %v, level: %v, xoffset: %v, yoffset: %v, x: %v, y: %v, width: %v, height: %v)", a.Target, a.Level, a.Xoffset, a.Yoffset, a.X, a.Y, a.Width, a.Height)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCopyTexSubImage2D pointer is returned so that calls can be chained.
func (a *GlCopyTexSubImage2D) AddRead(rng memory.Range, id binary.ID) *GlCopyTexSubImage2D {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCopyTexSubImage2D pointer is returned so that calls can be chained.
func (a *GlCopyTexSubImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCopyTexSubImage2D {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCopyTexSubImage2D) API() gfxapi.API                  { return api{} }
func (c *GlCopyTexSubImage2D) TypeID() atom.TypeID              { return 114 }
func (c *GlCopyTexSubImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCopyTexSubImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       TextureImageTarget
	Level        int32
	Format       CompressedTexelFormat
	Width        int32
	Height       int32
	Border       int32
	ImageSize    int32
	Data         TexturePointer
}

func (a *GlCompressedTexImage2D) String() string {
	return fmt.Sprintf("glCompressedTexImage2D(target: %v, level: %v, format: %v, width: %v, height: %v, border: %v, image_size: %v, data: %v)", a.Target, a.Level, a.Format, a.Width, a.Height, a.Border, a.ImageSize, a.Data)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCompressedTexImage2D pointer is returned so that calls can be chained.
func (a *GlCompressedTexImage2D) AddRead(rng memory.Range, id binary.ID) *GlCompressedTexImage2D {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCompressedTexImage2D pointer is returned so that calls can be chained.
func (a *GlCompressedTexImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCompressedTexImage2D {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCompressedTexImage2D) API() gfxapi.API                  { return api{} }
func (c *GlCompressedTexImage2D) TypeID() atom.TypeID              { return 115 }
func (c *GlCompressedTexImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCompressedTexImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexSubImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       TextureImageTarget
	Level        int32
	Xoffset      int32
	Yoffset      int32
	Width        int32
	Height       int32
	Format       CompressedTexelFormat
	ImageSize    int32
	Data         TexturePointer
}

func (a *GlCompressedTexSubImage2D) String() string {
	return fmt.Sprintf("glCompressedTexSubImage2D(target: %v, level: %v, xoffset: %v, yoffset: %v, width: %v, height: %v, format: %v, image_size: %v, data: %v)", a.Target, a.Level, a.Xoffset, a.Yoffset, a.Width, a.Height, a.Format, a.ImageSize, a.Data)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCompressedTexSubImage2D pointer is returned so that calls can be chained.
func (a *GlCompressedTexSubImage2D) AddRead(rng memory.Range, id binary.ID) *GlCompressedTexSubImage2D {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCompressedTexSubImage2D pointer is returned so that calls can be chained.
func (a *GlCompressedTexSubImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCompressedTexSubImage2D {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCompressedTexSubImage2D) API() gfxapi.API                  { return api{} }
func (c *GlCompressedTexSubImage2D) TypeID() atom.TypeID              { return 116 }
func (c *GlCompressedTexSubImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCompressedTexSubImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenerateMipmap
////////////////////////////////////////////////////////////////////////////////
type GlGenerateMipmap struct {
	binary.Generate
	observations atom.Observations
	Target       TextureImageTarget
}

func (a *GlGenerateMipmap) String() string {
	return fmt.Sprintf("glGenerateMipmap(target: %v)", a.Target)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenerateMipmap pointer is returned so that calls can be chained.
func (a *GlGenerateMipmap) AddRead(rng memory.Range, id binary.ID) *GlGenerateMipmap {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenerateMipmap pointer is returned so that calls can be chained.
func (a *GlGenerateMipmap) AddWrite(rng memory.Range, id binary.ID) *GlGenerateMipmap {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenerateMipmap) API() gfxapi.API                  { return api{} }
func (c *GlGenerateMipmap) TypeID() atom.TypeID              { return 117 }
func (c *GlGenerateMipmap) Flags() atom.Flags                { return 0 }
func (a *GlGenerateMipmap) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlReadPixels
////////////////////////////////////////////////////////////////////////////////
type GlReadPixels struct {
	binary.Generate
	observations atom.Observations
	X            int32
	Y            int32
	Width        int32
	Height       int32
	Format       BaseTexelFormat
	Type         TexelType
	Data         Voidᵖ
}

func (a *GlReadPixels) String() string {
	return fmt.Sprintf("glReadPixels(x: %v, y: %v, width: %v, height: %v, format: %v, type: %v, data: %v)", a.X, a.Y, a.Width, a.Height, a.Format, a.Type, a.Data)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlReadPixels pointer is returned so that calls can be chained.
func (a *GlReadPixels) AddRead(rng memory.Range, id binary.ID) *GlReadPixels {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlReadPixels pointer is returned so that calls can be chained.
func (a *GlReadPixels) AddWrite(rng memory.Range, id binary.ID) *GlReadPixels {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlReadPixels) API() gfxapi.API                  { return api{} }
func (c *GlReadPixels) TypeID() atom.TypeID              { return 118 }
func (c *GlReadPixels) Flags() atom.Flags                { return 0 }
func (a *GlReadPixels) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenFramebuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenFramebuffers struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Framebuffers FramebufferIdᵖ
}

func (a *GlGenFramebuffers) String() string {
	return fmt.Sprintf("glGenFramebuffers(count: %v, framebuffers: %v)", a.Count, a.Framebuffers)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenFramebuffers pointer is returned so that calls can be chained.
func (a *GlGenFramebuffers) AddRead(rng memory.Range, id binary.ID) *GlGenFramebuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenFramebuffers pointer is returned so that calls can be chained.
func (a *GlGenFramebuffers) AddWrite(rng memory.Range, id binary.ID) *GlGenFramebuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenFramebuffers) API() gfxapi.API                  { return api{} }
func (c *GlGenFramebuffers) TypeID() atom.TypeID              { return 119 }
func (c *GlGenFramebuffers) Flags() atom.Flags                { return 0 }
func (a *GlGenFramebuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindFramebuffer struct {
	binary.Generate
	observations atom.Observations
	Target       FramebufferTarget
	Framebuffer  FramebufferId
}

func (a *GlBindFramebuffer) String() string {
	return fmt.Sprintf("glBindFramebuffer(target: %v, framebuffer: %v)", a.Target, a.Framebuffer)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindFramebuffer pointer is returned so that calls can be chained.
func (a *GlBindFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlBindFramebuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindFramebuffer pointer is returned so that calls can be chained.
func (a *GlBindFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlBindFramebuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindFramebuffer) API() gfxapi.API                  { return api{} }
func (c *GlBindFramebuffer) TypeID() atom.TypeID              { return 120 }
func (c *GlBindFramebuffer) Flags() atom.Flags                { return 0 }
func (a *GlBindFramebuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCheckFramebufferStatus
////////////////////////////////////////////////////////////////////////////////
type GlCheckFramebufferStatus struct {
	binary.Generate
	observations atom.Observations
	Target       FramebufferTarget
	Result       FramebufferStatus
}

func (a *GlCheckFramebufferStatus) String() string {
	return fmt.Sprintf("glCheckFramebufferStatus(target: %v) → %v", a.Target, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCheckFramebufferStatus pointer is returned so that calls can be chained.
func (a *GlCheckFramebufferStatus) AddRead(rng memory.Range, id binary.ID) *GlCheckFramebufferStatus {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCheckFramebufferStatus pointer is returned so that calls can be chained.
func (a *GlCheckFramebufferStatus) AddWrite(rng memory.Range, id binary.ID) *GlCheckFramebufferStatus {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCheckFramebufferStatus) API() gfxapi.API                  { return api{} }
func (c *GlCheckFramebufferStatus) TypeID() atom.TypeID              { return 121 }
func (c *GlCheckFramebufferStatus) Flags() atom.Flags                { return 0 }
func (a *GlCheckFramebufferStatus) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteFramebuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteFramebuffers struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Framebuffers FramebufferIdᵖ
}

func (a *GlDeleteFramebuffers) String() string {
	return fmt.Sprintf("glDeleteFramebuffers(count: %v, framebuffers: %v)", a.Count, a.Framebuffers)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteFramebuffers pointer is returned so that calls can be chained.
func (a *GlDeleteFramebuffers) AddRead(rng memory.Range, id binary.ID) *GlDeleteFramebuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteFramebuffers pointer is returned so that calls can be chained.
func (a *GlDeleteFramebuffers) AddWrite(rng memory.Range, id binary.ID) *GlDeleteFramebuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteFramebuffers) API() gfxapi.API                  { return api{} }
func (c *GlDeleteFramebuffers) TypeID() atom.TypeID              { return 122 }
func (c *GlDeleteFramebuffers) Flags() atom.Flags                { return 0 }
func (a *GlDeleteFramebuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsFramebuffer struct {
	binary.Generate
	observations atom.Observations
	Framebuffer  FramebufferId
	Result       bool
}

func (a *GlIsFramebuffer) String() string {
	return fmt.Sprintf("glIsFramebuffer(framebuffer: %v) → %v", a.Framebuffer, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsFramebuffer pointer is returned so that calls can be chained.
func (a *GlIsFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlIsFramebuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsFramebuffer pointer is returned so that calls can be chained.
func (a *GlIsFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlIsFramebuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsFramebuffer) API() gfxapi.API                  { return api{} }
func (c *GlIsFramebuffer) TypeID() atom.TypeID              { return 123 }
func (c *GlIsFramebuffer) Flags() atom.Flags                { return 0 }
func (a *GlIsFramebuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenRenderbuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenRenderbuffers struct {
	binary.Generate
	observations  atom.Observations
	Count         int32
	Renderbuffers RenderbufferIdᵖ
}

func (a *GlGenRenderbuffers) String() string {
	return fmt.Sprintf("glGenRenderbuffers(count: %v, renderbuffers: %v)", a.Count, a.Renderbuffers)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenRenderbuffers pointer is returned so that calls can be chained.
func (a *GlGenRenderbuffers) AddRead(rng memory.Range, id binary.ID) *GlGenRenderbuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenRenderbuffers pointer is returned so that calls can be chained.
func (a *GlGenRenderbuffers) AddWrite(rng memory.Range, id binary.ID) *GlGenRenderbuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenRenderbuffers) API() gfxapi.API                  { return api{} }
func (c *GlGenRenderbuffers) TypeID() atom.TypeID              { return 124 }
func (c *GlGenRenderbuffers) Flags() atom.Flags                { return 0 }
func (a *GlGenRenderbuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindRenderbuffer struct {
	binary.Generate
	observations atom.Observations
	Target       RenderbufferTarget
	Renderbuffer RenderbufferId
}

func (a *GlBindRenderbuffer) String() string {
	return fmt.Sprintf("glBindRenderbuffer(target: %v, renderbuffer: %v)", a.Target, a.Renderbuffer)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindRenderbuffer pointer is returned so that calls can be chained.
func (a *GlBindRenderbuffer) AddRead(rng memory.Range, id binary.ID) *GlBindRenderbuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindRenderbuffer pointer is returned so that calls can be chained.
func (a *GlBindRenderbuffer) AddWrite(rng memory.Range, id binary.ID) *GlBindRenderbuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindRenderbuffer) API() gfxapi.API                  { return api{} }
func (c *GlBindRenderbuffer) TypeID() atom.TypeID              { return 125 }
func (c *GlBindRenderbuffer) Flags() atom.Flags                { return 0 }
func (a *GlBindRenderbuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorage
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorage struct {
	binary.Generate
	observations atom.Observations
	Target       RenderbufferTarget
	Format       RenderbufferFormat
	Width        int32
	Height       int32
}

func (a *GlRenderbufferStorage) String() string {
	return fmt.Sprintf("glRenderbufferStorage(target: %v, format: %v, width: %v, height: %v)", a.Target, a.Format, a.Width, a.Height)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlRenderbufferStorage pointer is returned so that calls can be chained.
func (a *GlRenderbufferStorage) AddRead(rng memory.Range, id binary.ID) *GlRenderbufferStorage {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlRenderbufferStorage pointer is returned so that calls can be chained.
func (a *GlRenderbufferStorage) AddWrite(rng memory.Range, id binary.ID) *GlRenderbufferStorage {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlRenderbufferStorage) API() gfxapi.API                  { return api{} }
func (c *GlRenderbufferStorage) TypeID() atom.TypeID              { return 126 }
func (c *GlRenderbufferStorage) Flags() atom.Flags                { return 0 }
func (a *GlRenderbufferStorage) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteRenderbuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteRenderbuffers struct {
	binary.Generate
	observations  atom.Observations
	Count         int32
	Renderbuffers RenderbufferIdᵖ
}

func (a *GlDeleteRenderbuffers) String() string {
	return fmt.Sprintf("glDeleteRenderbuffers(count: %v, renderbuffers: %v)", a.Count, a.Renderbuffers)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteRenderbuffers pointer is returned so that calls can be chained.
func (a *GlDeleteRenderbuffers) AddRead(rng memory.Range, id binary.ID) *GlDeleteRenderbuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteRenderbuffers pointer is returned so that calls can be chained.
func (a *GlDeleteRenderbuffers) AddWrite(rng memory.Range, id binary.ID) *GlDeleteRenderbuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteRenderbuffers) API() gfxapi.API                  { return api{} }
func (c *GlDeleteRenderbuffers) TypeID() atom.TypeID              { return 127 }
func (c *GlDeleteRenderbuffers) Flags() atom.Flags                { return 0 }
func (a *GlDeleteRenderbuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsRenderbuffer struct {
	binary.Generate
	observations atom.Observations
	Renderbuffer RenderbufferId
	Result       bool
}

func (a *GlIsRenderbuffer) String() string {
	return fmt.Sprintf("glIsRenderbuffer(renderbuffer: %v) → %v", a.Renderbuffer, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsRenderbuffer pointer is returned so that calls can be chained.
func (a *GlIsRenderbuffer) AddRead(rng memory.Range, id binary.ID) *GlIsRenderbuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsRenderbuffer pointer is returned so that calls can be chained.
func (a *GlIsRenderbuffer) AddWrite(rng memory.Range, id binary.ID) *GlIsRenderbuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsRenderbuffer) API() gfxapi.API                  { return api{} }
func (c *GlIsRenderbuffer) TypeID() atom.TypeID              { return 128 }
func (c *GlIsRenderbuffer) Flags() atom.Flags                { return 0 }
func (a *GlIsRenderbuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetRenderbufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetRenderbufferParameteriv struct {
	binary.Generate
	observations atom.Observations
	Target       RenderbufferTarget
	Parameter    RenderbufferParameter
	Values       S32ᵖ
}

func (a *GlGetRenderbufferParameteriv) String() string {
	return fmt.Sprintf("glGetRenderbufferParameteriv(target: %v, parameter: %v, values: %v)", a.Target, a.Parameter, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetRenderbufferParameteriv pointer is returned so that calls can be chained.
func (a *GlGetRenderbufferParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetRenderbufferParameteriv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetRenderbufferParameteriv pointer is returned so that calls can be chained.
func (a *GlGetRenderbufferParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetRenderbufferParameteriv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetRenderbufferParameteriv) API() gfxapi.API                  { return api{} }
func (c *GlGetRenderbufferParameteriv) TypeID() atom.TypeID              { return 129 }
func (c *GlGetRenderbufferParameteriv) Flags() atom.Flags                { return 0 }
func (a *GlGetRenderbufferParameteriv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenBuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenBuffers struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Buffers      BufferIdᵖ
}

func (a *GlGenBuffers) String() string {
	return fmt.Sprintf("glGenBuffers(count: %v, buffers: %v)", a.Count, a.Buffers)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenBuffers pointer is returned so that calls can be chained.
func (a *GlGenBuffers) AddRead(rng memory.Range, id binary.ID) *GlGenBuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenBuffers pointer is returned so that calls can be chained.
func (a *GlGenBuffers) AddWrite(rng memory.Range, id binary.ID) *GlGenBuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenBuffers) API() gfxapi.API                  { return api{} }
func (c *GlGenBuffers) TypeID() atom.TypeID              { return 130 }
func (c *GlGenBuffers) Flags() atom.Flags                { return 0 }
func (a *GlGenBuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindBuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindBuffer struct {
	binary.Generate
	observations atom.Observations
	Target       BufferTarget
	Buffer       BufferId
}

func (a *GlBindBuffer) String() string {
	return fmt.Sprintf("glBindBuffer(target: %v, buffer: %v)", a.Target, a.Buffer)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindBuffer pointer is returned so that calls can be chained.
func (a *GlBindBuffer) AddRead(rng memory.Range, id binary.ID) *GlBindBuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindBuffer pointer is returned so that calls can be chained.
func (a *GlBindBuffer) AddWrite(rng memory.Range, id binary.ID) *GlBindBuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindBuffer) API() gfxapi.API                  { return api{} }
func (c *GlBindBuffer) TypeID() atom.TypeID              { return 131 }
func (c *GlBindBuffer) Flags() atom.Flags                { return 0 }
func (a *GlBindBuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBufferData
////////////////////////////////////////////////////////////////////////////////
type GlBufferData struct {
	binary.Generate
	observations atom.Observations
	Target       BufferTarget
	Size         int32
	Data         BufferDataPointer
	Usage        BufferUsage
}

func (a *GlBufferData) String() string {
	return fmt.Sprintf("glBufferData(target: %v, size: %v, data: %v, usage: %v)", a.Target, a.Size, a.Data, a.Usage)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBufferData pointer is returned so that calls can be chained.
func (a *GlBufferData) AddRead(rng memory.Range, id binary.ID) *GlBufferData {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBufferData pointer is returned so that calls can be chained.
func (a *GlBufferData) AddWrite(rng memory.Range, id binary.ID) *GlBufferData {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBufferData) API() gfxapi.API                  { return api{} }
func (c *GlBufferData) TypeID() atom.TypeID              { return 132 }
func (c *GlBufferData) Flags() atom.Flags                { return 0 }
func (a *GlBufferData) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBufferSubData
////////////////////////////////////////////////////////////////////////////////
type GlBufferSubData struct {
	binary.Generate
	observations atom.Observations
	Target       BufferTarget
	Offset       int32
	Size         int32
	Data         Voidᵖ
}

func (a *GlBufferSubData) String() string {
	return fmt.Sprintf("glBufferSubData(target: %v, offset: %v, size: %v, data: %v)", a.Target, a.Offset, a.Size, a.Data)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBufferSubData pointer is returned so that calls can be chained.
func (a *GlBufferSubData) AddRead(rng memory.Range, id binary.ID) *GlBufferSubData {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBufferSubData pointer is returned so that calls can be chained.
func (a *GlBufferSubData) AddWrite(rng memory.Range, id binary.ID) *GlBufferSubData {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBufferSubData) API() gfxapi.API                  { return api{} }
func (c *GlBufferSubData) TypeID() atom.TypeID              { return 133 }
func (c *GlBufferSubData) Flags() atom.Flags                { return 0 }
func (a *GlBufferSubData) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteBuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteBuffers struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Buffers      BufferIdᵖ
}

func (a *GlDeleteBuffers) String() string {
	return fmt.Sprintf("glDeleteBuffers(count: %v, buffers: %v)", a.Count, a.Buffers)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteBuffers pointer is returned so that calls can be chained.
func (a *GlDeleteBuffers) AddRead(rng memory.Range, id binary.ID) *GlDeleteBuffers {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteBuffers pointer is returned so that calls can be chained.
func (a *GlDeleteBuffers) AddWrite(rng memory.Range, id binary.ID) *GlDeleteBuffers {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteBuffers) API() gfxapi.API                  { return api{} }
func (c *GlDeleteBuffers) TypeID() atom.TypeID              { return 134 }
func (c *GlDeleteBuffers) Flags() atom.Flags                { return 0 }
func (a *GlDeleteBuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsBuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsBuffer struct {
	binary.Generate
	observations atom.Observations
	Buffer       BufferId
	Result       bool
}

func (a *GlIsBuffer) String() string {
	return fmt.Sprintf("glIsBuffer(buffer: %v) → %v", a.Buffer, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsBuffer pointer is returned so that calls can be chained.
func (a *GlIsBuffer) AddRead(rng memory.Range, id binary.ID) *GlIsBuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsBuffer pointer is returned so that calls can be chained.
func (a *GlIsBuffer) AddWrite(rng memory.Range, id binary.ID) *GlIsBuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsBuffer) API() gfxapi.API                  { return api{} }
func (c *GlIsBuffer) TypeID() atom.TypeID              { return 135 }
func (c *GlIsBuffer) Flags() atom.Flags                { return 0 }
func (a *GlIsBuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetBufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetBufferParameteriv struct {
	binary.Generate
	observations atom.Observations
	Target       BufferTarget
	Parameter    BufferParameter
	Value        S32ᵖ
}

func (a *GlGetBufferParameteriv) String() string {
	return fmt.Sprintf("glGetBufferParameteriv(target: %v, parameter: %v, value: %v)", a.Target, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetBufferParameteriv pointer is returned so that calls can be chained.
func (a *GlGetBufferParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetBufferParameteriv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetBufferParameteriv pointer is returned so that calls can be chained.
func (a *GlGetBufferParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetBufferParameteriv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetBufferParameteriv) API() gfxapi.API                  { return api{} }
func (c *GlGetBufferParameteriv) TypeID() atom.TypeID              { return 136 }
func (c *GlGetBufferParameteriv) Flags() atom.Flags                { return 0 }
func (a *GlGetBufferParameteriv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCreateShader
////////////////////////////////////////////////////////////////////////////////
type GlCreateShader struct {
	binary.Generate
	observations atom.Observations
	Type         ShaderType
	Result       ShaderId
}

func (a *GlCreateShader) String() string {
	return fmt.Sprintf("glCreateShader(type: %v) → %v", a.Type, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCreateShader pointer is returned so that calls can be chained.
func (a *GlCreateShader) AddRead(rng memory.Range, id binary.ID) *GlCreateShader {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCreateShader pointer is returned so that calls can be chained.
func (a *GlCreateShader) AddWrite(rng memory.Range, id binary.ID) *GlCreateShader {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCreateShader) API() gfxapi.API                  { return api{} }
func (c *GlCreateShader) TypeID() atom.TypeID              { return 137 }
func (c *GlCreateShader) Flags() atom.Flags                { return 0 }
func (a *GlCreateShader) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteShader
////////////////////////////////////////////////////////////////////////////////
type GlDeleteShader struct {
	binary.Generate
	observations atom.Observations
	Shader       ShaderId
}

func (a *GlDeleteShader) String() string {
	return fmt.Sprintf("glDeleteShader(shader: %v)", a.Shader)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteShader pointer is returned so that calls can be chained.
func (a *GlDeleteShader) AddRead(rng memory.Range, id binary.ID) *GlDeleteShader {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteShader pointer is returned so that calls can be chained.
func (a *GlDeleteShader) AddWrite(rng memory.Range, id binary.ID) *GlDeleteShader {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteShader) API() gfxapi.API                  { return api{} }
func (c *GlDeleteShader) TypeID() atom.TypeID              { return 138 }
func (c *GlDeleteShader) Flags() atom.Flags                { return 0 }
func (a *GlDeleteShader) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlShaderSource
////////////////////////////////////////////////////////////////////////////////
type GlShaderSource struct {
	binary.Generate
	observations atom.Observations
	Shader       ShaderId
	Count        int32
	Source       Charᵖᵖ
	Length       S32ᵖ
}

func (a *GlShaderSource) String() string {
	return fmt.Sprintf("glShaderSource(shader: %v, count: %v, source: %v, length: %v)", a.Shader, a.Count, a.Source, a.Length)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlShaderSource pointer is returned so that calls can be chained.
func (a *GlShaderSource) AddRead(rng memory.Range, id binary.ID) *GlShaderSource {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlShaderSource pointer is returned so that calls can be chained.
func (a *GlShaderSource) AddWrite(rng memory.Range, id binary.ID) *GlShaderSource {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlShaderSource) API() gfxapi.API                  { return api{} }
func (c *GlShaderSource) TypeID() atom.TypeID              { return 139 }
func (c *GlShaderSource) Flags() atom.Flags                { return 0 }
func (a *GlShaderSource) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlShaderBinary
////////////////////////////////////////////////////////////////////////////////
type GlShaderBinary struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Shaders      ShaderIdᵖ
	BinaryFormat uint32
	Binary       Voidᵖ
	BinarySize   int32
}

func (a *GlShaderBinary) String() string {
	return fmt.Sprintf("glShaderBinary(count: %v, shaders: %v, binary_format: %v, binary: %v, binary_size: %v)", a.Count, a.Shaders, a.BinaryFormat, a.Binary, a.BinarySize)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlShaderBinary pointer is returned so that calls can be chained.
func (a *GlShaderBinary) AddRead(rng memory.Range, id binary.ID) *GlShaderBinary {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlShaderBinary pointer is returned so that calls can be chained.
func (a *GlShaderBinary) AddWrite(rng memory.Range, id binary.ID) *GlShaderBinary {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlShaderBinary) API() gfxapi.API                  { return api{} }
func (c *GlShaderBinary) TypeID() atom.TypeID              { return 140 }
func (c *GlShaderBinary) Flags() atom.Flags                { return 0 }
func (a *GlShaderBinary) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderInfoLog
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderInfoLog struct {
	binary.Generate
	observations        atom.Observations
	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten S32ᵖ
	Info                Charᵖ
}

func (a *GlGetShaderInfoLog) String() string {
	return fmt.Sprintf("glGetShaderInfoLog(shader: %v, buffer_length: %v, string_length_written: %v, info: %v)", a.Shader, a.BufferLength, a.StringLengthWritten, a.Info)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetShaderInfoLog pointer is returned so that calls can be chained.
func (a *GlGetShaderInfoLog) AddRead(rng memory.Range, id binary.ID) *GlGetShaderInfoLog {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetShaderInfoLog pointer is returned so that calls can be chained.
func (a *GlGetShaderInfoLog) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderInfoLog {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetShaderInfoLog) API() gfxapi.API                  { return api{} }
func (c *GlGetShaderInfoLog) TypeID() atom.TypeID              { return 141 }
func (c *GlGetShaderInfoLog) Flags() atom.Flags                { return 0 }
func (a *GlGetShaderInfoLog) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderSource
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderSource struct {
	binary.Generate
	observations        atom.Observations
	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten S32ᵖ
	Source              Charᵖ
}

func (a *GlGetShaderSource) String() string {
	return fmt.Sprintf("glGetShaderSource(shader: %v, buffer_length: %v, string_length_written: %v, source: %v)", a.Shader, a.BufferLength, a.StringLengthWritten, a.Source)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetShaderSource pointer is returned so that calls can be chained.
func (a *GlGetShaderSource) AddRead(rng memory.Range, id binary.ID) *GlGetShaderSource {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetShaderSource pointer is returned so that calls can be chained.
func (a *GlGetShaderSource) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderSource {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetShaderSource) API() gfxapi.API                  { return api{} }
func (c *GlGetShaderSource) TypeID() atom.TypeID              { return 142 }
func (c *GlGetShaderSource) Flags() atom.Flags                { return 0 }
func (a *GlGetShaderSource) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlReleaseShaderCompiler
////////////////////////////////////////////////////////////////////////////////
type GlReleaseShaderCompiler struct {
	binary.Generate
	observations atom.Observations
}

func (a *GlReleaseShaderCompiler) String() string {
	return fmt.Sprintf("glReleaseShaderCompiler()")
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlReleaseShaderCompiler pointer is returned so that calls can be chained.
func (a *GlReleaseShaderCompiler) AddRead(rng memory.Range, id binary.ID) *GlReleaseShaderCompiler {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlReleaseShaderCompiler pointer is returned so that calls can be chained.
func (a *GlReleaseShaderCompiler) AddWrite(rng memory.Range, id binary.ID) *GlReleaseShaderCompiler {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlReleaseShaderCompiler) API() gfxapi.API                  { return api{} }
func (c *GlReleaseShaderCompiler) TypeID() atom.TypeID              { return 143 }
func (c *GlReleaseShaderCompiler) Flags() atom.Flags                { return 0 }
func (a *GlReleaseShaderCompiler) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCompileShader
////////////////////////////////////////////////////////////////////////////////
type GlCompileShader struct {
	binary.Generate
	observations atom.Observations
	Shader       ShaderId
}

func (a *GlCompileShader) String() string {
	return fmt.Sprintf("glCompileShader(shader: %v)", a.Shader)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCompileShader pointer is returned so that calls can be chained.
func (a *GlCompileShader) AddRead(rng memory.Range, id binary.ID) *GlCompileShader {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCompileShader pointer is returned so that calls can be chained.
func (a *GlCompileShader) AddWrite(rng memory.Range, id binary.ID) *GlCompileShader {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCompileShader) API() gfxapi.API                  { return api{} }
func (c *GlCompileShader) TypeID() atom.TypeID              { return 144 }
func (c *GlCompileShader) Flags() atom.Flags                { return 0 }
func (a *GlCompileShader) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsShader
////////////////////////////////////////////////////////////////////////////////
type GlIsShader struct {
	binary.Generate
	observations atom.Observations
	Shader       ShaderId
	Result       bool
}

func (a *GlIsShader) String() string {
	return fmt.Sprintf("glIsShader(shader: %v) → %v", a.Shader, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsShader pointer is returned so that calls can be chained.
func (a *GlIsShader) AddRead(rng memory.Range, id binary.ID) *GlIsShader {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsShader pointer is returned so that calls can be chained.
func (a *GlIsShader) AddWrite(rng memory.Range, id binary.ID) *GlIsShader {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsShader) API() gfxapi.API                  { return api{} }
func (c *GlIsShader) TypeID() atom.TypeID              { return 145 }
func (c *GlIsShader) Flags() atom.Flags                { return 0 }
func (a *GlIsShader) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCreateProgram
////////////////////////////////////////////////////////////////////////////////
type GlCreateProgram struct {
	binary.Generate
	observations atom.Observations
	Result       ProgramId
}

func (a *GlCreateProgram) String() string {
	return fmt.Sprintf("glCreateProgram() → %v", a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCreateProgram pointer is returned so that calls can be chained.
func (a *GlCreateProgram) AddRead(rng memory.Range, id binary.ID) *GlCreateProgram {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCreateProgram pointer is returned so that calls can be chained.
func (a *GlCreateProgram) AddWrite(rng memory.Range, id binary.ID) *GlCreateProgram {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCreateProgram) API() gfxapi.API                  { return api{} }
func (c *GlCreateProgram) TypeID() atom.TypeID              { return 146 }
func (c *GlCreateProgram) Flags() atom.Flags                { return 0 }
func (a *GlCreateProgram) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteProgram
////////////////////////////////////////////////////////////////////////////////
type GlDeleteProgram struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
}

func (a *GlDeleteProgram) String() string {
	return fmt.Sprintf("glDeleteProgram(program: %v)", a.Program)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteProgram pointer is returned so that calls can be chained.
func (a *GlDeleteProgram) AddRead(rng memory.Range, id binary.ID) *GlDeleteProgram {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteProgram pointer is returned so that calls can be chained.
func (a *GlDeleteProgram) AddWrite(rng memory.Range, id binary.ID) *GlDeleteProgram {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteProgram) API() gfxapi.API                  { return api{} }
func (c *GlDeleteProgram) TypeID() atom.TypeID              { return 147 }
func (c *GlDeleteProgram) Flags() atom.Flags                { return 0 }
func (a *GlDeleteProgram) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlAttachShader
////////////////////////////////////////////////////////////////////////////////
type GlAttachShader struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Shader       ShaderId
}

func (a *GlAttachShader) String() string {
	return fmt.Sprintf("glAttachShader(program: %v, shader: %v)", a.Program, a.Shader)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlAttachShader pointer is returned so that calls can be chained.
func (a *GlAttachShader) AddRead(rng memory.Range, id binary.ID) *GlAttachShader {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlAttachShader pointer is returned so that calls can be chained.
func (a *GlAttachShader) AddWrite(rng memory.Range, id binary.ID) *GlAttachShader {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlAttachShader) API() gfxapi.API                  { return api{} }
func (c *GlAttachShader) TypeID() atom.TypeID              { return 148 }
func (c *GlAttachShader) Flags() atom.Flags                { return 0 }
func (a *GlAttachShader) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDetachShader
////////////////////////////////////////////////////////////////////////////////
type GlDetachShader struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Shader       ShaderId
}

func (a *GlDetachShader) String() string {
	return fmt.Sprintf("glDetachShader(program: %v, shader: %v)", a.Program, a.Shader)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDetachShader pointer is returned so that calls can be chained.
func (a *GlDetachShader) AddRead(rng memory.Range, id binary.ID) *GlDetachShader {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDetachShader pointer is returned so that calls can be chained.
func (a *GlDetachShader) AddWrite(rng memory.Range, id binary.ID) *GlDetachShader {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDetachShader) API() gfxapi.API                  { return api{} }
func (c *GlDetachShader) TypeID() atom.TypeID              { return 149 }
func (c *GlDetachShader) Flags() atom.Flags                { return 0 }
func (a *GlDetachShader) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetAttachedShaders
////////////////////////////////////////////////////////////////////////////////
type GlGetAttachedShaders struct {
	binary.Generate
	observations         atom.Observations
	Program              ProgramId
	BufferLength         int32
	ShadersLengthWritten S32ᵖ
	Shaders              ShaderIdᵖ
}

func (a *GlGetAttachedShaders) String() string {
	return fmt.Sprintf("glGetAttachedShaders(program: %v, buffer_length: %v, shaders_length_written: %v, shaders: %v)", a.Program, a.BufferLength, a.ShadersLengthWritten, a.Shaders)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetAttachedShaders pointer is returned so that calls can be chained.
func (a *GlGetAttachedShaders) AddRead(rng memory.Range, id binary.ID) *GlGetAttachedShaders {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetAttachedShaders pointer is returned so that calls can be chained.
func (a *GlGetAttachedShaders) AddWrite(rng memory.Range, id binary.ID) *GlGetAttachedShaders {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetAttachedShaders) API() gfxapi.API                  { return api{} }
func (c *GlGetAttachedShaders) TypeID() atom.TypeID              { return 150 }
func (c *GlGetAttachedShaders) Flags() atom.Flags                { return 0 }
func (a *GlGetAttachedShaders) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlLinkProgram
////////////////////////////////////////////////////////////////////////////////
type GlLinkProgram struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
}

func (a *GlLinkProgram) String() string {
	return fmt.Sprintf("glLinkProgram(program: %v)", a.Program)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlLinkProgram pointer is returned so that calls can be chained.
func (a *GlLinkProgram) AddRead(rng memory.Range, id binary.ID) *GlLinkProgram {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlLinkProgram pointer is returned so that calls can be chained.
func (a *GlLinkProgram) AddWrite(rng memory.Range, id binary.ID) *GlLinkProgram {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlLinkProgram) API() gfxapi.API                  { return api{} }
func (c *GlLinkProgram) TypeID() atom.TypeID              { return 151 }
func (c *GlLinkProgram) Flags() atom.Flags                { return 0 }
func (a *GlLinkProgram) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramInfoLog
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramInfoLog struct {
	binary.Generate
	observations        atom.Observations
	Program             ProgramId
	BufferLength        int32
	StringLengthWritten S32ᵖ
	Info                Charᵖ
}

func (a *GlGetProgramInfoLog) String() string {
	return fmt.Sprintf("glGetProgramInfoLog(program: %v, buffer_length: %v, string_length_written: %v, info: %v)", a.Program, a.BufferLength, a.StringLengthWritten, a.Info)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetProgramInfoLog pointer is returned so that calls can be chained.
func (a *GlGetProgramInfoLog) AddRead(rng memory.Range, id binary.ID) *GlGetProgramInfoLog {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetProgramInfoLog pointer is returned so that calls can be chained.
func (a *GlGetProgramInfoLog) AddWrite(rng memory.Range, id binary.ID) *GlGetProgramInfoLog {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetProgramInfoLog) API() gfxapi.API                  { return api{} }
func (c *GlGetProgramInfoLog) TypeID() atom.TypeID              { return 152 }
func (c *GlGetProgramInfoLog) Flags() atom.Flags                { return 0 }
func (a *GlGetProgramInfoLog) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUseProgram
////////////////////////////////////////////////////////////////////////////////
type GlUseProgram struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
}

func (a *GlUseProgram) String() string {
	return fmt.Sprintf("glUseProgram(program: %v)", a.Program)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUseProgram pointer is returned so that calls can be chained.
func (a *GlUseProgram) AddRead(rng memory.Range, id binary.ID) *GlUseProgram {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUseProgram pointer is returned so that calls can be chained.
func (a *GlUseProgram) AddWrite(rng memory.Range, id binary.ID) *GlUseProgram {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUseProgram) API() gfxapi.API                  { return api{} }
func (c *GlUseProgram) TypeID() atom.TypeID              { return 153 }
func (c *GlUseProgram) Flags() atom.Flags                { return 0 }
func (a *GlUseProgram) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsProgram
////////////////////////////////////////////////////////////////////////////////
type GlIsProgram struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Result       bool
}

func (a *GlIsProgram) String() string {
	return fmt.Sprintf("glIsProgram(program: %v) → %v", a.Program, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsProgram pointer is returned so that calls can be chained.
func (a *GlIsProgram) AddRead(rng memory.Range, id binary.ID) *GlIsProgram {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsProgram pointer is returned so that calls can be chained.
func (a *GlIsProgram) AddWrite(rng memory.Range, id binary.ID) *GlIsProgram {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsProgram) API() gfxapi.API                  { return api{} }
func (c *GlIsProgram) TypeID() atom.TypeID              { return 154 }
func (c *GlIsProgram) Flags() atom.Flags                { return 0 }
func (a *GlIsProgram) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlValidateProgram
////////////////////////////////////////////////////////////////////////////////
type GlValidateProgram struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
}

func (a *GlValidateProgram) String() string {
	return fmt.Sprintf("glValidateProgram(program: %v)", a.Program)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlValidateProgram pointer is returned so that calls can be chained.
func (a *GlValidateProgram) AddRead(rng memory.Range, id binary.ID) *GlValidateProgram {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlValidateProgram pointer is returned so that calls can be chained.
func (a *GlValidateProgram) AddWrite(rng memory.Range, id binary.ID) *GlValidateProgram {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlValidateProgram) API() gfxapi.API                  { return api{} }
func (c *GlValidateProgram) TypeID() atom.TypeID              { return 155 }
func (c *GlValidateProgram) Flags() atom.Flags                { return 0 }
func (a *GlValidateProgram) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlClearColor
////////////////////////////////////////////////////////////////////////////////
type GlClearColor struct {
	binary.Generate
	observations atom.Observations
	R            float32
	G            float32
	B            float32
	A            float32
}

func (a *GlClearColor) String() string {
	return fmt.Sprintf("glClearColor(r: %v, g: %v, b: %v, a: %v)", a.R, a.G, a.B, a.A)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlClearColor pointer is returned so that calls can be chained.
func (a *GlClearColor) AddRead(rng memory.Range, id binary.ID) *GlClearColor {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlClearColor pointer is returned so that calls can be chained.
func (a *GlClearColor) AddWrite(rng memory.Range, id binary.ID) *GlClearColor {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlClearColor) API() gfxapi.API                  { return api{} }
func (c *GlClearColor) TypeID() atom.TypeID              { return 156 }
func (c *GlClearColor) Flags() atom.Flags                { return 0 }
func (a *GlClearColor) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlClearDepthf
////////////////////////////////////////////////////////////////////////////////
type GlClearDepthf struct {
	binary.Generate
	observations atom.Observations
	Depth        float32
}

func (a *GlClearDepthf) String() string {
	return fmt.Sprintf("glClearDepthf(depth: %v)", a.Depth)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlClearDepthf pointer is returned so that calls can be chained.
func (a *GlClearDepthf) AddRead(rng memory.Range, id binary.ID) *GlClearDepthf {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlClearDepthf pointer is returned so that calls can be chained.
func (a *GlClearDepthf) AddWrite(rng memory.Range, id binary.ID) *GlClearDepthf {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlClearDepthf) API() gfxapi.API                  { return api{} }
func (c *GlClearDepthf) TypeID() atom.TypeID              { return 157 }
func (c *GlClearDepthf) Flags() atom.Flags                { return 0 }
func (a *GlClearDepthf) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlClearStencil
////////////////////////////////////////////////////////////////////////////////
type GlClearStencil struct {
	binary.Generate
	observations atom.Observations
	Stencil      int32
}

func (a *GlClearStencil) String() string {
	return fmt.Sprintf("glClearStencil(stencil: %v)", a.Stencil)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlClearStencil pointer is returned so that calls can be chained.
func (a *GlClearStencil) AddRead(rng memory.Range, id binary.ID) *GlClearStencil {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlClearStencil pointer is returned so that calls can be chained.
func (a *GlClearStencil) AddWrite(rng memory.Range, id binary.ID) *GlClearStencil {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlClearStencil) API() gfxapi.API                  { return api{} }
func (c *GlClearStencil) TypeID() atom.TypeID              { return 158 }
func (c *GlClearStencil) Flags() atom.Flags                { return 0 }
func (a *GlClearStencil) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlClear
////////////////////////////////////////////////////////////////////////////////
type GlClear struct {
	binary.Generate
	observations atom.Observations
	Mask         ClearMask
}

func (a *GlClear) String() string {
	return fmt.Sprintf("glClear(mask: %v)", a.Mask)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlClear pointer is returned so that calls can be chained.
func (a *GlClear) AddRead(rng memory.Range, id binary.ID) *GlClear {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlClear pointer is returned so that calls can be chained.
func (a *GlClear) AddWrite(rng memory.Range, id binary.ID) *GlClear {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlClear) API() gfxapi.API                  { return api{} }
func (c *GlClear) TypeID() atom.TypeID              { return 159 }
func (c *GlClear) Flags() atom.Flags                { return 0 }
func (a *GlClear) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCullFace
////////////////////////////////////////////////////////////////////////////////
type GlCullFace struct {
	binary.Generate
	observations atom.Observations
	Mode         FaceMode
}

func (a *GlCullFace) String() string {
	return fmt.Sprintf("glCullFace(mode: %v)", a.Mode)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlCullFace pointer is returned so that calls can be chained.
func (a *GlCullFace) AddRead(rng memory.Range, id binary.ID) *GlCullFace {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlCullFace pointer is returned so that calls can be chained.
func (a *GlCullFace) AddWrite(rng memory.Range, id binary.ID) *GlCullFace {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlCullFace) API() gfxapi.API                  { return api{} }
func (c *GlCullFace) TypeID() atom.TypeID              { return 160 }
func (c *GlCullFace) Flags() atom.Flags                { return 0 }
func (a *GlCullFace) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlPolygonOffset
////////////////////////////////////////////////////////////////////////////////
type GlPolygonOffset struct {
	binary.Generate
	observations atom.Observations
	ScaleFactor  float32
	Units        float32
}

func (a *GlPolygonOffset) String() string {
	return fmt.Sprintf("glPolygonOffset(scale_factor: %v, units: %v)", a.ScaleFactor, a.Units)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlPolygonOffset pointer is returned so that calls can be chained.
func (a *GlPolygonOffset) AddRead(rng memory.Range, id binary.ID) *GlPolygonOffset {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlPolygonOffset pointer is returned so that calls can be chained.
func (a *GlPolygonOffset) AddWrite(rng memory.Range, id binary.ID) *GlPolygonOffset {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlPolygonOffset) API() gfxapi.API                  { return api{} }
func (c *GlPolygonOffset) TypeID() atom.TypeID              { return 161 }
func (c *GlPolygonOffset) Flags() atom.Flags                { return 0 }
func (a *GlPolygonOffset) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlLineWidth
////////////////////////////////////////////////////////////////////////////////
type GlLineWidth struct {
	binary.Generate
	observations atom.Observations
	Width        float32
}

func (a *GlLineWidth) String() string {
	return fmt.Sprintf("glLineWidth(width: %v)", a.Width)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlLineWidth pointer is returned so that calls can be chained.
func (a *GlLineWidth) AddRead(rng memory.Range, id binary.ID) *GlLineWidth {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlLineWidth pointer is returned so that calls can be chained.
func (a *GlLineWidth) AddWrite(rng memory.Range, id binary.ID) *GlLineWidth {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlLineWidth) API() gfxapi.API                  { return api{} }
func (c *GlLineWidth) TypeID() atom.TypeID              { return 162 }
func (c *GlLineWidth) Flags() atom.Flags                { return 0 }
func (a *GlLineWidth) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlSampleCoverage
////////////////////////////////////////////////////////////////////////////////
type GlSampleCoverage struct {
	binary.Generate
	observations atom.Observations
	Value        float32
	Invert       bool
}

func (a *GlSampleCoverage) String() string {
	return fmt.Sprintf("glSampleCoverage(value: %v, invert: %v)", a.Value, a.Invert)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlSampleCoverage pointer is returned so that calls can be chained.
func (a *GlSampleCoverage) AddRead(rng memory.Range, id binary.ID) *GlSampleCoverage {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlSampleCoverage pointer is returned so that calls can be chained.
func (a *GlSampleCoverage) AddWrite(rng memory.Range, id binary.ID) *GlSampleCoverage {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlSampleCoverage) API() gfxapi.API                  { return api{} }
func (c *GlSampleCoverage) TypeID() atom.TypeID              { return 163 }
func (c *GlSampleCoverage) Flags() atom.Flags                { return 0 }
func (a *GlSampleCoverage) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlHint
////////////////////////////////////////////////////////////////////////////////
type GlHint struct {
	binary.Generate
	observations atom.Observations
	Target       HintTarget
	Mode         HintMode
}

func (a *GlHint) String() string {
	return fmt.Sprintf("glHint(target: %v, mode: %v)", a.Target, a.Mode)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlHint pointer is returned so that calls can be chained.
func (a *GlHint) AddRead(rng memory.Range, id binary.ID) *GlHint {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlHint pointer is returned so that calls can be chained.
func (a *GlHint) AddWrite(rng memory.Range, id binary.ID) *GlHint {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlHint) API() gfxapi.API                  { return api{} }
func (c *GlHint) TypeID() atom.TypeID              { return 164 }
func (c *GlHint) Flags() atom.Flags                { return 0 }
func (a *GlHint) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferRenderbuffer struct {
	binary.Generate
	observations          atom.Observations
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	RenderbufferTarget    RenderbufferTarget
	Renderbuffer          RenderbufferId
}

func (a *GlFramebufferRenderbuffer) String() string {
	return fmt.Sprintf("glFramebufferRenderbuffer(framebuffer_target: %v, framebuffer_attachment: %v, renderbuffer_target: %v, renderbuffer: %v)", a.FramebufferTarget, a.FramebufferAttachment, a.RenderbufferTarget, a.Renderbuffer)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlFramebufferRenderbuffer pointer is returned so that calls can be chained.
func (a *GlFramebufferRenderbuffer) AddRead(rng memory.Range, id binary.ID) *GlFramebufferRenderbuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlFramebufferRenderbuffer pointer is returned so that calls can be chained.
func (a *GlFramebufferRenderbuffer) AddWrite(rng memory.Range, id binary.ID) *GlFramebufferRenderbuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlFramebufferRenderbuffer) API() gfxapi.API                  { return api{} }
func (c *GlFramebufferRenderbuffer) TypeID() atom.TypeID              { return 165 }
func (c *GlFramebufferRenderbuffer) Flags() atom.Flags                { return 0 }
func (a *GlFramebufferRenderbuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferTexture2D
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferTexture2D struct {
	binary.Generate
	observations          atom.Observations
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	TextureTarget         TextureImageTarget
	Texture               TextureId
	Level                 int32
}

func (a *GlFramebufferTexture2D) String() string {
	return fmt.Sprintf("glFramebufferTexture2D(framebuffer_target: %v, framebuffer_attachment: %v, texture_target: %v, texture: %v, level: %v)", a.FramebufferTarget, a.FramebufferAttachment, a.TextureTarget, a.Texture, a.Level)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlFramebufferTexture2D pointer is returned so that calls can be chained.
func (a *GlFramebufferTexture2D) AddRead(rng memory.Range, id binary.ID) *GlFramebufferTexture2D {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlFramebufferTexture2D pointer is returned so that calls can be chained.
func (a *GlFramebufferTexture2D) AddWrite(rng memory.Range, id binary.ID) *GlFramebufferTexture2D {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlFramebufferTexture2D) API() gfxapi.API                  { return api{} }
func (c *GlFramebufferTexture2D) TypeID() atom.TypeID              { return 166 }
func (c *GlFramebufferTexture2D) Flags() atom.Flags                { return 0 }
func (a *GlFramebufferTexture2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetFramebufferAttachmentParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetFramebufferAttachmentParameteriv struct {
	binary.Generate
	observations      atom.Observations
	FramebufferTarget FramebufferTarget
	Attachment        FramebufferAttachment
	Parameter         FramebufferAttachmentParameter
	Value             S32ᵖ
}

func (a *GlGetFramebufferAttachmentParameteriv) String() string {
	return fmt.Sprintf("glGetFramebufferAttachmentParameteriv(framebuffer_target: %v, attachment: %v, parameter: %v, value: %v)", a.FramebufferTarget, a.Attachment, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetFramebufferAttachmentParameteriv pointer is returned so that calls can be chained.
func (a *GlGetFramebufferAttachmentParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetFramebufferAttachmentParameteriv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetFramebufferAttachmentParameteriv pointer is returned so that calls can be chained.
func (a *GlGetFramebufferAttachmentParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetFramebufferAttachmentParameteriv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetFramebufferAttachmentParameteriv) API() gfxapi.API     { return api{} }
func (c *GlGetFramebufferAttachmentParameteriv) TypeID() atom.TypeID { return 167 }
func (c *GlGetFramebufferAttachmentParameteriv) Flags() atom.Flags   { return 0 }
func (a *GlGetFramebufferAttachmentParameteriv) Observations() *atom.Observations {
	return &a.observations
}

////////////////////////////////////////////////////////////////////////////////
// GlDrawElements
////////////////////////////////////////////////////////////////////////////////
type GlDrawElements struct {
	binary.Generate
	observations atom.Observations
	DrawMode     DrawMode
	ElementCount int32
	IndicesType  IndicesType
	Indices      IndicesPointer
}

func (a *GlDrawElements) String() string {
	return fmt.Sprintf("glDrawElements(draw_mode: %v, element_count: %v, indices_type: %v, indices: %v)", a.DrawMode, a.ElementCount, a.IndicesType, a.Indices)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDrawElements pointer is returned so that calls can be chained.
func (a *GlDrawElements) AddRead(rng memory.Range, id binary.ID) *GlDrawElements {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDrawElements pointer is returned so that calls can be chained.
func (a *GlDrawElements) AddWrite(rng memory.Range, id binary.ID) *GlDrawElements {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDrawElements) API() gfxapi.API                  { return api{} }
func (c *GlDrawElements) TypeID() atom.TypeID              { return 168 }
func (c *GlDrawElements) Flags() atom.Flags                { return 0 | atom.DrawCall }
func (a *GlDrawElements) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDrawArrays
////////////////////////////////////////////////////////////////////////////////
type GlDrawArrays struct {
	binary.Generate
	observations atom.Observations
	DrawMode     DrawMode
	FirstIndex   int32
	IndexCount   int32
}

func (a *GlDrawArrays) String() string {
	return fmt.Sprintf("glDrawArrays(draw_mode: %v, first_index: %v, index_count: %v)", a.DrawMode, a.FirstIndex, a.IndexCount)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDrawArrays pointer is returned so that calls can be chained.
func (a *GlDrawArrays) AddRead(rng memory.Range, id binary.ID) *GlDrawArrays {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDrawArrays pointer is returned so that calls can be chained.
func (a *GlDrawArrays) AddWrite(rng memory.Range, id binary.ID) *GlDrawArrays {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDrawArrays) API() gfxapi.API                  { return api{} }
func (c *GlDrawArrays) TypeID() atom.TypeID              { return 169 }
func (c *GlDrawArrays) Flags() atom.Flags                { return 0 | atom.DrawCall }
func (a *GlDrawArrays) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFlush
////////////////////////////////////////////////////////////////////////////////
type GlFlush struct {
	binary.Generate
	observations atom.Observations
}

func (a *GlFlush) String() string {
	return fmt.Sprintf("glFlush()")
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlFlush pointer is returned so that calls can be chained.
func (a *GlFlush) AddRead(rng memory.Range, id binary.ID) *GlFlush {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlFlush pointer is returned so that calls can be chained.
func (a *GlFlush) AddWrite(rng memory.Range, id binary.ID) *GlFlush {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlFlush) API() gfxapi.API                  { return api{} }
func (c *GlFlush) TypeID() atom.TypeID              { return 170 }
func (c *GlFlush) Flags() atom.Flags                { return 0 }
func (a *GlFlush) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFinish
////////////////////////////////////////////////////////////////////////////////
type GlFinish struct {
	binary.Generate
	observations atom.Observations
}

func (a *GlFinish) String() string {
	return fmt.Sprintf("glFinish()")
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlFinish pointer is returned so that calls can be chained.
func (a *GlFinish) AddRead(rng memory.Range, id binary.ID) *GlFinish {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlFinish pointer is returned so that calls can be chained.
func (a *GlFinish) AddWrite(rng memory.Range, id binary.ID) *GlFinish {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlFinish) API() gfxapi.API                  { return api{} }
func (c *GlFinish) TypeID() atom.TypeID              { return 171 }
func (c *GlFinish) Flags() atom.Flags                { return 0 }
func (a *GlFinish) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetBooleanv
////////////////////////////////////////////////////////////////////////////////
type GlGetBooleanv struct {
	binary.Generate
	observations atom.Observations
	Param        StateVariable
	Values       Boolᵖ
}

func (a *GlGetBooleanv) String() string {
	return fmt.Sprintf("glGetBooleanv(param: %v, values: %v)", a.Param, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetBooleanv pointer is returned so that calls can be chained.
func (a *GlGetBooleanv) AddRead(rng memory.Range, id binary.ID) *GlGetBooleanv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetBooleanv pointer is returned so that calls can be chained.
func (a *GlGetBooleanv) AddWrite(rng memory.Range, id binary.ID) *GlGetBooleanv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetBooleanv) API() gfxapi.API                  { return api{} }
func (c *GlGetBooleanv) TypeID() atom.TypeID              { return 172 }
func (c *GlGetBooleanv) Flags() atom.Flags                { return 0 }
func (a *GlGetBooleanv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetFloatv
////////////////////////////////////////////////////////////////////////////////
type GlGetFloatv struct {
	binary.Generate
	observations atom.Observations
	Param        StateVariable
	Values       F32ᵖ
}

func (a *GlGetFloatv) String() string {
	return fmt.Sprintf("glGetFloatv(param: %v, values: %v)", a.Param, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetFloatv pointer is returned so that calls can be chained.
func (a *GlGetFloatv) AddRead(rng memory.Range, id binary.ID) *GlGetFloatv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetFloatv pointer is returned so that calls can be chained.
func (a *GlGetFloatv) AddWrite(rng memory.Range, id binary.ID) *GlGetFloatv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetFloatv) API() gfxapi.API                  { return api{} }
func (c *GlGetFloatv) TypeID() atom.TypeID              { return 173 }
func (c *GlGetFloatv) Flags() atom.Flags                { return 0 }
func (a *GlGetFloatv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetIntegerv
////////////////////////////////////////////////////////////////////////////////
type GlGetIntegerv struct {
	binary.Generate
	observations atom.Observations
	Param        StateVariable
	Values       S32ᵖ
}

func (a *GlGetIntegerv) String() string {
	return fmt.Sprintf("glGetIntegerv(param: %v, values: %v)", a.Param, a.Values)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetIntegerv pointer is returned so that calls can be chained.
func (a *GlGetIntegerv) AddRead(rng memory.Range, id binary.ID) *GlGetIntegerv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetIntegerv pointer is returned so that calls can be chained.
func (a *GlGetIntegerv) AddWrite(rng memory.Range, id binary.ID) *GlGetIntegerv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetIntegerv) API() gfxapi.API                  { return api{} }
func (c *GlGetIntegerv) TypeID() atom.TypeID              { return 174 }
func (c *GlGetIntegerv) Flags() atom.Flags                { return 0 }
func (a *GlGetIntegerv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetString
////////////////////////////////////////////////////////////////////////////////
type GlGetString struct {
	binary.Generate
	observations atom.Observations
	Param        StringConstant
	Result       Charᵖ
}

func (a *GlGetString) String() string {
	return fmt.Sprintf("glGetString(param: %v) → %v", a.Param, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetString pointer is returned so that calls can be chained.
func (a *GlGetString) AddRead(rng memory.Range, id binary.ID) *GlGetString {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetString pointer is returned so that calls can be chained.
func (a *GlGetString) AddWrite(rng memory.Range, id binary.ID) *GlGetString {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetString) API() gfxapi.API                  { return api{} }
func (c *GlGetString) TypeID() atom.TypeID              { return 175 }
func (c *GlGetString) Flags() atom.Flags                { return 0 }
func (a *GlGetString) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEnable
////////////////////////////////////////////////////////////////////////////////
type GlEnable struct {
	binary.Generate
	observations atom.Observations
	Capability   Capability
}

func (a *GlEnable) String() string {
	return fmt.Sprintf("glEnable(capability: %v)", a.Capability)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEnable pointer is returned so that calls can be chained.
func (a *GlEnable) AddRead(rng memory.Range, id binary.ID) *GlEnable {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEnable pointer is returned so that calls can be chained.
func (a *GlEnable) AddWrite(rng memory.Range, id binary.ID) *GlEnable {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEnable) API() gfxapi.API                  { return api{} }
func (c *GlEnable) TypeID() atom.TypeID              { return 176 }
func (c *GlEnable) Flags() atom.Flags                { return 0 }
func (a *GlEnable) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDisable
////////////////////////////////////////////////////////////////////////////////
type GlDisable struct {
	binary.Generate
	observations atom.Observations
	Capability   Capability
}

func (a *GlDisable) String() string {
	return fmt.Sprintf("glDisable(capability: %v)", a.Capability)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDisable pointer is returned so that calls can be chained.
func (a *GlDisable) AddRead(rng memory.Range, id binary.ID) *GlDisable {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDisable pointer is returned so that calls can be chained.
func (a *GlDisable) AddWrite(rng memory.Range, id binary.ID) *GlDisable {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDisable) API() gfxapi.API                  { return api{} }
func (c *GlDisable) TypeID() atom.TypeID              { return 177 }
func (c *GlDisable) Flags() atom.Flags                { return 0 }
func (a *GlDisable) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsEnabled
////////////////////////////////////////////////////////////////////////////////
type GlIsEnabled struct {
	binary.Generate
	observations atom.Observations
	Capability   Capability
	Result       bool
}

func (a *GlIsEnabled) String() string {
	return fmt.Sprintf("glIsEnabled(capability: %v) → %v", a.Capability, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsEnabled pointer is returned so that calls can be chained.
func (a *GlIsEnabled) AddRead(rng memory.Range, id binary.ID) *GlIsEnabled {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsEnabled pointer is returned so that calls can be chained.
func (a *GlIsEnabled) AddWrite(rng memory.Range, id binary.ID) *GlIsEnabled {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsEnabled) API() gfxapi.API                  { return api{} }
func (c *GlIsEnabled) TypeID() atom.TypeID              { return 178 }
func (c *GlIsEnabled) Flags() atom.Flags                { return 0 }
func (a *GlIsEnabled) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlMapBufferRange
////////////////////////////////////////////////////////////////////////////////
type GlMapBufferRange struct {
	binary.Generate
	observations atom.Observations
	Target       BufferTarget
	Offset       int32
	Length       int32
	Access       MapBufferRangeAccess
	Result       Voidᵖ
}

func (a *GlMapBufferRange) String() string {
	return fmt.Sprintf("glMapBufferRange(target: %v, offset: %v, length: %v, access: %v) → %v", a.Target, a.Offset, a.Length, a.Access, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlMapBufferRange pointer is returned so that calls can be chained.
func (a *GlMapBufferRange) AddRead(rng memory.Range, id binary.ID) *GlMapBufferRange {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlMapBufferRange pointer is returned so that calls can be chained.
func (a *GlMapBufferRange) AddWrite(rng memory.Range, id binary.ID) *GlMapBufferRange {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlMapBufferRange) API() gfxapi.API                  { return api{} }
func (c *GlMapBufferRange) TypeID() atom.TypeID              { return 179 }
func (c *GlMapBufferRange) Flags() atom.Flags                { return 0 }
func (a *GlMapBufferRange) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUnmapBuffer
////////////////////////////////////////////////////////////////////////////////
type GlUnmapBuffer struct {
	binary.Generate
	observations atom.Observations
	Target       BufferTarget
}

func (a *GlUnmapBuffer) String() string {
	return fmt.Sprintf("glUnmapBuffer(target: %v)", a.Target)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUnmapBuffer pointer is returned so that calls can be chained.
func (a *GlUnmapBuffer) AddRead(rng memory.Range, id binary.ID) *GlUnmapBuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUnmapBuffer pointer is returned so that calls can be chained.
func (a *GlUnmapBuffer) AddWrite(rng memory.Range, id binary.ID) *GlUnmapBuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUnmapBuffer) API() gfxapi.API                  { return api{} }
func (c *GlUnmapBuffer) TypeID() atom.TypeID              { return 180 }
func (c *GlUnmapBuffer) Flags() atom.Flags                { return 0 }
func (a *GlUnmapBuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlInvalidateFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlInvalidateFramebuffer struct {
	binary.Generate
	observations atom.Observations
	Target       FramebufferTarget
	Count        int32
	Attachments  FramebufferAttachmentᵖ
}

func (a *GlInvalidateFramebuffer) String() string {
	return fmt.Sprintf("glInvalidateFramebuffer(target: %v, count: %v, attachments: %v)", a.Target, a.Count, a.Attachments)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlInvalidateFramebuffer pointer is returned so that calls can be chained.
func (a *GlInvalidateFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlInvalidateFramebuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlInvalidateFramebuffer pointer is returned so that calls can be chained.
func (a *GlInvalidateFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlInvalidateFramebuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlInvalidateFramebuffer) API() gfxapi.API                  { return api{} }
func (c *GlInvalidateFramebuffer) TypeID() atom.TypeID              { return 181 }
func (c *GlInvalidateFramebuffer) Flags() atom.Flags                { return 0 }
func (a *GlInvalidateFramebuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorageMultisample
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorageMultisample struct {
	binary.Generate
	observations atom.Observations
	Target       RenderbufferTarget
	Samples      int32
	Format       RenderbufferFormat
	Width        int32
	Height       int32
}

func (a *GlRenderbufferStorageMultisample) String() string {
	return fmt.Sprintf("glRenderbufferStorageMultisample(target: %v, samples: %v, format: %v, width: %v, height: %v)", a.Target, a.Samples, a.Format, a.Width, a.Height)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlRenderbufferStorageMultisample pointer is returned so that calls can be chained.
func (a *GlRenderbufferStorageMultisample) AddRead(rng memory.Range, id binary.ID) *GlRenderbufferStorageMultisample {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlRenderbufferStorageMultisample pointer is returned so that calls can be chained.
func (a *GlRenderbufferStorageMultisample) AddWrite(rng memory.Range, id binary.ID) *GlRenderbufferStorageMultisample {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlRenderbufferStorageMultisample) API() gfxapi.API                  { return api{} }
func (c *GlRenderbufferStorageMultisample) TypeID() atom.TypeID              { return 182 }
func (c *GlRenderbufferStorageMultisample) Flags() atom.Flags                { return 0 }
func (a *GlRenderbufferStorageMultisample) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlitFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlBlitFramebuffer struct {
	binary.Generate
	observations atom.Observations
	SrcX0        int32
	SrcY0        int32
	SrcX1        int32
	SrcY1        int32
	DstX0        int32
	DstY0        int32
	DstX1        int32
	DstY1        int32
	Mask         ClearMask
	Filter       TextureFilterMode
}

func (a *GlBlitFramebuffer) String() string {
	return fmt.Sprintf("glBlitFramebuffer(srcX0: %v, srcY0: %v, srcX1: %v, srcY1: %v, dstX0: %v, dstY0: %v, dstX1: %v, dstY1: %v, mask: %v, filter: %v)", a.SrcX0, a.SrcY0, a.SrcX1, a.SrcY1, a.DstX0, a.DstY0, a.DstX1, a.DstY1, a.Mask, a.Filter)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBlitFramebuffer pointer is returned so that calls can be chained.
func (a *GlBlitFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlBlitFramebuffer {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBlitFramebuffer pointer is returned so that calls can be chained.
func (a *GlBlitFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlBlitFramebuffer {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBlitFramebuffer) API() gfxapi.API                  { return api{} }
func (c *GlBlitFramebuffer) TypeID() atom.TypeID              { return 183 }
func (c *GlBlitFramebuffer) Flags() atom.Flags                { return 0 }
func (a *GlBlitFramebuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenQueries
////////////////////////////////////////////////////////////////////////////////
type GlGenQueries struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Queries      QueryIdᵖ
}

func (a *GlGenQueries) String() string {
	return fmt.Sprintf("glGenQueries(count: %v, queries: %v)", a.Count, a.Queries)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenQueries pointer is returned so that calls can be chained.
func (a *GlGenQueries) AddRead(rng memory.Range, id binary.ID) *GlGenQueries {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenQueries pointer is returned so that calls can be chained.
func (a *GlGenQueries) AddWrite(rng memory.Range, id binary.ID) *GlGenQueries {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenQueries) API() gfxapi.API                  { return api{} }
func (c *GlGenQueries) TypeID() atom.TypeID              { return 184 }
func (c *GlGenQueries) Flags() atom.Flags                { return 0 }
func (a *GlGenQueries) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBeginQuery
////////////////////////////////////////////////////////////////////////////////
type GlBeginQuery struct {
	binary.Generate
	observations atom.Observations
	Target       QueryTarget
	Query        QueryId
}

func (a *GlBeginQuery) String() string {
	return fmt.Sprintf("glBeginQuery(target: %v, query: %v)", a.Target, a.Query)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBeginQuery pointer is returned so that calls can be chained.
func (a *GlBeginQuery) AddRead(rng memory.Range, id binary.ID) *GlBeginQuery {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBeginQuery pointer is returned so that calls can be chained.
func (a *GlBeginQuery) AddWrite(rng memory.Range, id binary.ID) *GlBeginQuery {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBeginQuery) API() gfxapi.API                  { return api{} }
func (c *GlBeginQuery) TypeID() atom.TypeID              { return 185 }
func (c *GlBeginQuery) Flags() atom.Flags                { return 0 }
func (a *GlBeginQuery) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEndQuery
////////////////////////////////////////////////////////////////////////////////
type GlEndQuery struct {
	binary.Generate
	observations atom.Observations
	Target       QueryTarget
}

func (a *GlEndQuery) String() string {
	return fmt.Sprintf("glEndQuery(target: %v)", a.Target)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEndQuery pointer is returned so that calls can be chained.
func (a *GlEndQuery) AddRead(rng memory.Range, id binary.ID) *GlEndQuery {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEndQuery pointer is returned so that calls can be chained.
func (a *GlEndQuery) AddWrite(rng memory.Range, id binary.ID) *GlEndQuery {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEndQuery) API() gfxapi.API                  { return api{} }
func (c *GlEndQuery) TypeID() atom.TypeID              { return 186 }
func (c *GlEndQuery) Flags() atom.Flags                { return 0 }
func (a *GlEndQuery) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueries
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueries struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Queries      QueryIdᵖ
}

func (a *GlDeleteQueries) String() string {
	return fmt.Sprintf("glDeleteQueries(count: %v, queries: %v)", a.Count, a.Queries)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteQueries pointer is returned so that calls can be chained.
func (a *GlDeleteQueries) AddRead(rng memory.Range, id binary.ID) *GlDeleteQueries {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteQueries pointer is returned so that calls can be chained.
func (a *GlDeleteQueries) AddWrite(rng memory.Range, id binary.ID) *GlDeleteQueries {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteQueries) API() gfxapi.API                  { return api{} }
func (c *GlDeleteQueries) TypeID() atom.TypeID              { return 187 }
func (c *GlDeleteQueries) Flags() atom.Flags                { return 0 }
func (a *GlDeleteQueries) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsQuery
////////////////////////////////////////////////////////////////////////////////
type GlIsQuery struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Result       bool
}

func (a *GlIsQuery) String() string {
	return fmt.Sprintf("glIsQuery(query: %v) → %v", a.Query, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsQuery pointer is returned so that calls can be chained.
func (a *GlIsQuery) AddRead(rng memory.Range, id binary.ID) *GlIsQuery {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsQuery pointer is returned so that calls can be chained.
func (a *GlIsQuery) AddWrite(rng memory.Range, id binary.ID) *GlIsQuery {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsQuery) API() gfxapi.API                  { return api{} }
func (c *GlIsQuery) TypeID() atom.TypeID              { return 188 }
func (c *GlIsQuery) Flags() atom.Flags                { return 0 }
func (a *GlIsQuery) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryiv struct {
	binary.Generate
	observations atom.Observations
	Target       QueryTarget
	Parameter    QueryParameter
	Value        S32ᵖ
}

func (a *GlGetQueryiv) String() string {
	return fmt.Sprintf("glGetQueryiv(target: %v, parameter: %v, value: %v)", a.Target, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryiv pointer is returned so that calls can be chained.
func (a *GlGetQueryiv) AddRead(rng memory.Range, id binary.ID) *GlGetQueryiv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryiv pointer is returned so that calls can be chained.
func (a *GlGetQueryiv) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryiv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryiv) API() gfxapi.API                  { return api{} }
func (c *GlGetQueryiv) TypeID() atom.TypeID              { return 189 }
func (c *GlGetQueryiv) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuiv struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    QueryObjectParameter
	Value        U32ᵖ
}

func (a *GlGetQueryObjectuiv) String() string {
	return fmt.Sprintf("glGetQueryObjectuiv(query: %v, parameter: %v, value: %v)", a.Query, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectuiv pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectuiv) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectuiv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectuiv pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectuiv) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectuiv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryObjectuiv) API() gfxapi.API                  { return api{} }
func (c *GlGetQueryObjectuiv) TypeID() atom.TypeID              { return 190 }
func (c *GlGetQueryObjectuiv) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectuiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenQueriesEXT
////////////////////////////////////////////////////////////////////////////////
type GlGenQueriesEXT struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Queries      QueryIdᵖ
}

func (a *GlGenQueriesEXT) String() string {
	return fmt.Sprintf("glGenQueriesEXT(count: %v, queries: %v)", a.Count, a.Queries)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenQueriesEXT pointer is returned so that calls can be chained.
func (a *GlGenQueriesEXT) AddRead(rng memory.Range, id binary.ID) *GlGenQueriesEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenQueriesEXT pointer is returned so that calls can be chained.
func (a *GlGenQueriesEXT) AddWrite(rng memory.Range, id binary.ID) *GlGenQueriesEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenQueriesEXT) API() gfxapi.API                  { return api{} }
func (c *GlGenQueriesEXT) TypeID() atom.TypeID              { return 191 }
func (c *GlGenQueriesEXT) Flags() atom.Flags                { return 0 }
func (a *GlGenQueriesEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBeginQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlBeginQueryEXT struct {
	binary.Generate
	observations atom.Observations
	Target       QueryTarget
	Query        QueryId
}

func (a *GlBeginQueryEXT) String() string {
	return fmt.Sprintf("glBeginQueryEXT(target: %v, query: %v)", a.Target, a.Query)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBeginQueryEXT pointer is returned so that calls can be chained.
func (a *GlBeginQueryEXT) AddRead(rng memory.Range, id binary.ID) *GlBeginQueryEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBeginQueryEXT pointer is returned so that calls can be chained.
func (a *GlBeginQueryEXT) AddWrite(rng memory.Range, id binary.ID) *GlBeginQueryEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBeginQueryEXT) API() gfxapi.API                  { return api{} }
func (c *GlBeginQueryEXT) TypeID() atom.TypeID              { return 192 }
func (c *GlBeginQueryEXT) Flags() atom.Flags                { return 0 }
func (a *GlBeginQueryEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEndQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlEndQueryEXT struct {
	binary.Generate
	observations atom.Observations
	Target       QueryTarget
}

func (a *GlEndQueryEXT) String() string {
	return fmt.Sprintf("glEndQueryEXT(target: %v)", a.Target)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlEndQueryEXT pointer is returned so that calls can be chained.
func (a *GlEndQueryEXT) AddRead(rng memory.Range, id binary.ID) *GlEndQueryEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlEndQueryEXT pointer is returned so that calls can be chained.
func (a *GlEndQueryEXT) AddWrite(rng memory.Range, id binary.ID) *GlEndQueryEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlEndQueryEXT) API() gfxapi.API                  { return api{} }
func (c *GlEndQueryEXT) TypeID() atom.TypeID              { return 193 }
func (c *GlEndQueryEXT) Flags() atom.Flags                { return 0 }
func (a *GlEndQueryEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueriesEXT
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueriesEXT struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Queries      QueryIdᵖ
}

func (a *GlDeleteQueriesEXT) String() string {
	return fmt.Sprintf("glDeleteQueriesEXT(count: %v, queries: %v)", a.Count, a.Queries)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteQueriesEXT pointer is returned so that calls can be chained.
func (a *GlDeleteQueriesEXT) AddRead(rng memory.Range, id binary.ID) *GlDeleteQueriesEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteQueriesEXT pointer is returned so that calls can be chained.
func (a *GlDeleteQueriesEXT) AddWrite(rng memory.Range, id binary.ID) *GlDeleteQueriesEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteQueriesEXT) API() gfxapi.API                  { return api{} }
func (c *GlDeleteQueriesEXT) TypeID() atom.TypeID              { return 194 }
func (c *GlDeleteQueriesEXT) Flags() atom.Flags                { return 0 }
func (a *GlDeleteQueriesEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlIsQueryEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Result       bool
}

func (a *GlIsQueryEXT) String() string {
	return fmt.Sprintf("glIsQueryEXT(query: %v) → %v", a.Query, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlIsQueryEXT pointer is returned so that calls can be chained.
func (a *GlIsQueryEXT) AddRead(rng memory.Range, id binary.ID) *GlIsQueryEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlIsQueryEXT pointer is returned so that calls can be chained.
func (a *GlIsQueryEXT) AddWrite(rng memory.Range, id binary.ID) *GlIsQueryEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlIsQueryEXT) API() gfxapi.API                  { return api{} }
func (c *GlIsQueryEXT) TypeID() atom.TypeID              { return 195 }
func (c *GlIsQueryEXT) Flags() atom.Flags                { return 0 }
func (a *GlIsQueryEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlQueryCounterEXT
////////////////////////////////////////////////////////////////////////////////
type GlQueryCounterEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Target       QueryTarget
}

func (a *GlQueryCounterEXT) String() string {
	return fmt.Sprintf("glQueryCounterEXT(query: %v, target: %v)", a.Query, a.Target)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlQueryCounterEXT pointer is returned so that calls can be chained.
func (a *GlQueryCounterEXT) AddRead(rng memory.Range, id binary.ID) *GlQueryCounterEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlQueryCounterEXT pointer is returned so that calls can be chained.
func (a *GlQueryCounterEXT) AddWrite(rng memory.Range, id binary.ID) *GlQueryCounterEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlQueryCounterEXT) API() gfxapi.API                  { return api{} }
func (c *GlQueryCounterEXT) TypeID() atom.TypeID              { return 196 }
func (c *GlQueryCounterEXT) Flags() atom.Flags                { return 0 }
func (a *GlQueryCounterEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryivEXT struct {
	binary.Generate
	observations atom.Observations
	Target       QueryTarget
	Parameter    QueryParameter
	Value        S32ᵖ
}

func (a *GlGetQueryivEXT) String() string {
	return fmt.Sprintf("glGetQueryivEXT(target: %v, parameter: %v, value: %v)", a.Target, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryivEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryivEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryivEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryivEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryivEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryivEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryivEXT) API() gfxapi.API                  { return api{} }
func (c *GlGetQueryivEXT) TypeID() atom.TypeID              { return 197 }
func (c *GlGetQueryivEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryivEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectivEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    QueryObjectParameter
	Value        S32ᵖ
}

func (a *GlGetQueryObjectivEXT) String() string {
	return fmt.Sprintf("glGetQueryObjectivEXT(query: %v, parameter: %v, value: %v)", a.Query, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectivEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectivEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectivEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectivEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectivEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectivEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryObjectivEXT) API() gfxapi.API                  { return api{} }
func (c *GlGetQueryObjectivEXT) TypeID() atom.TypeID              { return 198 }
func (c *GlGetQueryObjectivEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectivEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuivEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    QueryObjectParameter
	Value        U32ᵖ
}

func (a *GlGetQueryObjectuivEXT) String() string {
	return fmt.Sprintf("glGetQueryObjectuivEXT(query: %v, parameter: %v, value: %v)", a.Query, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectuivEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectuivEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectuivEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectuivEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectuivEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectuivEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryObjectuivEXT) API() gfxapi.API                  { return api{} }
func (c *GlGetQueryObjectuivEXT) TypeID() atom.TypeID              { return 199 }
func (c *GlGetQueryObjectuivEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectuivEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjecti64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjecti64vEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    QueryObjectParameter
	Value        S64ᵖ
}

func (a *GlGetQueryObjecti64vEXT) String() string {
	return fmt.Sprintf("glGetQueryObjecti64vEXT(query: %v, parameter: %v, value: %v)", a.Query, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjecti64vEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjecti64vEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjecti64vEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjecti64vEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjecti64vEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjecti64vEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryObjecti64vEXT) API() gfxapi.API                  { return api{} }
func (c *GlGetQueryObjecti64vEXT) TypeID() atom.TypeID              { return 200 }
func (c *GlGetQueryObjecti64vEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjecti64vEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectui64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectui64vEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    QueryObjectParameter
	Value        U64ᵖ
}

func (a *GlGetQueryObjectui64vEXT) String() string {
	return fmt.Sprintf("glGetQueryObjectui64vEXT(query: %v, parameter: %v, value: %v)", a.Query, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectui64vEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectui64vEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectui64vEXT {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectui64vEXT pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectui64vEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectui64vEXT {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryObjectui64vEXT) API() gfxapi.API                  { return api{} }
func (c *GlGetQueryObjectui64vEXT) TypeID() atom.TypeID              { return 201 }
func (c *GlGetQueryObjectui64vEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectui64vEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// class Color
////////////////////////////////////////////////////////////////////////////////
type Color struct {
	binary.Generate
	CreatedAt atom.ID
	Red       float32
	Green     float32
	Blue      float32
	Alpha     float32
}

func (c *Color) Init() {
}
func (c *Color) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Rect
////////////////////////////////////////////////////////////////////////////////
type Rect struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
	Width     int32
	Height    int32
}

func (c *Rect) Init() {
}
func (c *Rect) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Image
////////////////////////////////////////////////////////////////////////////////
type Image struct {
	binary.Generate
	CreatedAt atom.ID
	Width     int32
	Height    int32
	Data      U8ˢ
	Size      uint32
	Format    ImageTexelFormat
}

func (c *Image) Init() {
}
func (c *Image) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Renderbuffer
////////////////////////////////////////////////////////////////////////////////
type Renderbuffer struct {
	binary.Generate
	CreatedAt atom.ID
	Width     int32
	Height    int32
	Data      U8ˢ
	Format    RenderbufferFormat
}

func (c *Renderbuffer) Init() {
}
func (c *Renderbuffer) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Texture
////////////////////////////////////////////////////////////////////////////////
type Texture struct {
	binary.Generate
	CreatedAt     atom.ID
	Kind          TextureKind
	Format        ImageTexelFormat
	Texture2D     S32ːImageᵐ
	Cubemap       S32ːCubemapLevelᵐ
	MagFilter     TextureFilterMode
	MinFilter     TextureFilterMode
	WrapS         TextureWrapMode
	WrapT         TextureWrapMode
	SwizzleR      TexelComponent
	SwizzleG      TexelComponent
	SwizzleB      TexelComponent
	SwizzleA      TexelComponent
	MaxAnisotropy float32
}

func (c *Texture) Init() {
	c.Texture2D = make(S32ːImageᵐ)
	c.Cubemap = make(S32ːCubemapLevelᵐ)
	c.MagFilter = TextureFilterMode_GL_LINEAR
	c.MinFilter = TextureFilterMode_GL_NEAREST_MIPMAP_LINEAR
	c.WrapS = TextureWrapMode_GL_REPEAT
	c.WrapT = TextureWrapMode_GL_REPEAT
	c.SwizzleR = TexelComponent_GL_RED
	c.SwizzleG = TexelComponent_GL_GREEN
	c.SwizzleB = TexelComponent_GL_BLUE
	c.SwizzleA = TexelComponent_GL_ALPHA
	c.MaxAnisotropy = 1
}
func (c *Texture) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class CubemapLevel
////////////////////////////////////////////////////////////////////////////////
type CubemapLevel struct {
	binary.Generate
	CreatedAt atom.ID
	Faces     CubeMapImageTargetːImageᵐ
}

func (c *CubemapLevel) Init() {
	c.Faces = make(CubeMapImageTargetːImageᵐ)
}
func (c *CubemapLevel) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class FramebufferAttachmentInfo
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachmentInfo struct {
	binary.Generate
	CreatedAt    atom.ID
	Object       uint32
	Type         FramebufferAttachmentType
	TextureLevel int32
	CubeMapFace  CubeMapImageTarget
}

func (c *FramebufferAttachmentInfo) Init() {
}
func (c *FramebufferAttachmentInfo) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Framebuffer
////////////////////////////////////////////////////////////////////////////////
type Framebuffer struct {
	binary.Generate
	CreatedAt   atom.ID
	Attachments FramebufferAttachmentːFramebufferAttachmentInfoᵐ
}

func (c *Framebuffer) Init() {
	c.Attachments = make(FramebufferAttachmentːFramebufferAttachmentInfoᵐ)
}
func (c *Framebuffer) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Buffer
////////////////////////////////////////////////////////////////////////////////
type Buffer struct {
	binary.Generate
	CreatedAt atom.ID
	Data      U8ˢ
	Size      int32
	Usage     BufferUsage
}

func (c *Buffer) Init() {
	c.Size = 0
	c.Usage = BufferUsage_GL_STATIC_DRAW
}
func (c *Buffer) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Shader
////////////////////////////////////////////////////////////////////////////////
type Shader struct {
	binary.Generate
	CreatedAt atom.ID
	Binary    U8ˢ
	Compiled  bool
	Deletable bool
	InfoLog   Charˢ
	Source    string
	Type      ShaderType
}

func (c *Shader) Init() {
	c.Compiled = false
	c.Deletable = false
}
func (c *Shader) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class VertexAttribute
////////////////////////////////////////////////////////////////////////////////
type VertexAttribute struct {
	binary.Generate
	CreatedAt   atom.ID
	Name        Charˢ
	VectorCount int32
	Type        ShaderAttribType
}

func (c *VertexAttribute) Init() {
}
func (c *VertexAttribute) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Vec2i
////////////////////////////////////////////////////////////////////////////////
type Vec2i struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
}

func (c *Vec2i) Init() {
}
func (c *Vec2i) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Vec3i
////////////////////////////////////////////////////////////////////////////////
type Vec3i struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
	Z         int32
}

func (c *Vec3i) Init() {
}
func (c *Vec3i) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Vec4i
////////////////////////////////////////////////////////////////////////////////
type Vec4i struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
	Z         int32
	W         int32
}

func (c *Vec4i) Init() {
}
func (c *Vec4i) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Vec2f
////////////////////////////////////////////////////////////////////////////////
type Vec2f struct {
	binary.Generate
	CreatedAt atom.ID
	X         float32
	Y         float32
}

func (c *Vec2f) Init() {
}
func (c *Vec2f) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Vec3f
////////////////////////////////////////////////////////////////////////////////
type Vec3f struct {
	binary.Generate
	CreatedAt atom.ID
	X         float32
	Y         float32
	Z         float32
}

func (c *Vec3f) Init() {
}
func (c *Vec3f) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Vec4f
////////////////////////////////////////////////////////////////////////////////
type Vec4f struct {
	binary.Generate
	CreatedAt atom.ID
	X         float32
	Y         float32
	Z         float32
	W         float32
}

func (c *Vec4f) Init() {
}
func (c *Vec4f) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Mat2f
////////////////////////////////////////////////////////////////////////////////
type Mat2f struct {
	binary.Generate
	CreatedAt atom.ID
	Col0      Vec2f
	Col1      Vec2f
}

func (c *Mat2f) Init() {
	c.Col0.Init()
	c.Col1.Init()
}
func (c *Mat2f) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Mat3f
////////////////////////////////////////////////////////////////////////////////
type Mat3f struct {
	binary.Generate
	CreatedAt atom.ID
	Col0      Vec3f
	Col1      Vec3f
	Col2      Vec3f
}

func (c *Mat3f) Init() {
	c.Col0.Init()
	c.Col1.Init()
	c.Col2.Init()
}
func (c *Mat3f) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Mat4f
////////////////////////////////////////////////////////////////////////////////
type Mat4f struct {
	binary.Generate
	CreatedAt atom.ID
	Col0      Vec4f
	Col1      Vec4f
	Col2      Vec4f
	Col3      Vec4f
}

func (c *Mat4f) Init() {
	c.Col0.Init()
	c.Col1.Init()
	c.Col2.Init()
	c.Col3.Init()
}
func (c *Mat4f) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class UniformValue
////////////////////////////////////////////////////////////////////////////////
type UniformValue struct {
	binary.Generate
	CreatedAt atom.ID
	F32       float32
	Vec2f     Vec2f
	Vec3f     Vec3f
	Vec4f     Vec4f
	S32       int32
	Vec2i     Vec2i
	Vec3i     Vec3i
	Vec4i     Vec4i
	Mat2f     Mat2f
	Mat3f     Mat3f
	Mat4f     Mat4f
}

func (c *UniformValue) Init() {
	c.Vec2f.Init()
	c.Vec3f.Init()
	c.Vec4f.Init()
	c.Vec2i.Init()
	c.Vec3i.Init()
	c.Vec4i.Init()
	c.Mat2f.Init()
	c.Mat3f.Init()
	c.Mat4f.Init()
}
func (c *UniformValue) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Uniform
////////////////////////////////////////////////////////////////////////////////
type Uniform struct {
	binary.Generate
	CreatedAt atom.ID
	Name      string
	Type      ShaderUniformType
	Value     UniformValue
}

func (c *Uniform) Init() {
	c.Value.Init()
}
func (c *Uniform) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Program
////////////////////////////////////////////////////////////////////////////////
type Program struct {
	binary.Generate
	CreatedAt         atom.ID
	Shaders           ShaderTypeːShaderIdᵐ
	Linked            bool
	Binary            U8ˢ
	AttributeBindings StringːAttributeLocationᵐ
	Attributes        S32ːVertexAttributeᵐ
	Uniforms          UniformLocationːUniformᵐ
	InfoLog           Charˢ
}

func (c *Program) Init() {
	c.Shaders = make(ShaderTypeːShaderIdᵐ)
	c.AttributeBindings = make(StringːAttributeLocationᵐ)
	c.Attributes = make(S32ːVertexAttributeᵐ)
	c.Uniforms = make(UniformLocationːUniformᵐ)
}
func (c *Program) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class VertexArray
////////////////////////////////////////////////////////////////////////////////
type VertexArray struct {
	binary.Generate
	CreatedAt atom.ID
}

func (c *VertexArray) Init() {
}
func (c *VertexArray) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class VertexAttributeArray
////////////////////////////////////////////////////////////////////////////////
type VertexAttributeArray struct {
	binary.Generate
	CreatedAt  atom.ID
	Enabled    bool
	Size       uint32
	Type       VertexAttribType
	Normalized bool
	Stride     int32
	Buffer     BufferId
	Pointer    VertexPointer
}

func (c *VertexAttributeArray) Init() {
	c.Enabled = false
	c.Size = 4
	c.Type = VertexAttribType_GL_FLOAT
	c.Normalized = false
	c.Stride = 0
	c.Buffer = 0
}
func (c *VertexAttributeArray) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Query
////////////////////////////////////////////////////////////////////////////////
type Query struct {
	binary.Generate
	CreatedAt atom.ID
}

func (c *Query) Init() {
}
func (c *Query) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class BlendState
////////////////////////////////////////////////////////////////////////////////
type BlendState struct {
	binary.Generate
	CreatedAt           atom.ID
	SrcRgbBlendFactor   BlendFactor
	SrcAlphaBlendFactor BlendFactor
	DstRgbBlendFactor   BlendFactor
	DstAlphaBlendFactor BlendFactor
	BlendEquationRgb    BlendEquation
	BlendEquationAlpha  BlendEquation
	BlendColor          Color
}

func (c *BlendState) Init() {
	c.SrcRgbBlendFactor = BlendFactor_GL_ONE
	c.SrcAlphaBlendFactor = BlendFactor_GL_ZERO
	c.DstRgbBlendFactor = BlendFactor_GL_ONE
	c.DstAlphaBlendFactor = BlendFactor_GL_ZERO
	c.BlendEquationRgb = BlendEquation_GL_FUNC_ADD
	c.BlendEquationAlpha = BlendEquation_GL_FUNC_ADD
	c.BlendColor.Init()
}
func (c *BlendState) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class RasterizerState
////////////////////////////////////////////////////////////////////////////////
type RasterizerState struct {
	binary.Generate
	CreatedAt            atom.ID
	DepthMask            bool
	DepthTestFunction    TestFunction
	DepthNear            float32
	DepthFar             float32
	ColorMaskRed         bool
	ColorMaskGreen       bool
	ColorMaskBlue        bool
	ColorMaskAlpha       bool
	StencilMask          FaceModeːu32ᵐ
	Viewport             Rect
	Scissor              Rect
	FrontFace            FaceOrientation
	CullFace             FaceMode
	LineWidth            float32
	PolygonOffsetFactor  float32
	PolygonOffsetUnits   float32
	SampleCoverageValue  float32
	SampleCoverageInvert bool
}

func (c *RasterizerState) Init() {
	c.DepthMask = true
	c.DepthTestFunction = TestFunction_GL_LESS
	c.DepthNear = 0
	c.DepthFar = 1
	c.ColorMaskRed = true
	c.ColorMaskGreen = true
	c.ColorMaskBlue = true
	c.ColorMaskAlpha = true
	c.StencilMask = make(FaceModeːu32ᵐ)
	c.Viewport.Init()
	c.Scissor.Init()
	c.FrontFace = FaceOrientation_GL_CCW
	c.CullFace = FaceMode_GL_BACK
	c.LineWidth = 1
	c.SampleCoverageValue = 1
}
func (c *RasterizerState) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class ClearState
////////////////////////////////////////////////////////////////////////////////
type ClearState struct {
	binary.Generate
	CreatedAt    atom.ID
	ClearColor   Color
	ClearDepth   float32
	ClearStencil int32
}

func (c *ClearState) Init() {
	c.ClearColor.Init()
	c.ClearDepth = 1
}
func (c *ClearState) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Objects
////////////////////////////////////////////////////////////////////////////////
type Objects struct {
	binary.Generate
	CreatedAt     atom.ID
	Renderbuffers RenderbufferIdːRenderbufferʳᵐ
	Textures      TextureIdːTextureʳᵐ
	Framebuffers  FramebufferIdːFramebufferʳᵐ
	Buffers       BufferIdːBufferʳᵐ
	Shaders       ShaderIdːShaderʳᵐ
	Programs      ProgramIdːProgramʳᵐ
	VertexArrays  VertexArrayIdːVertexArrayʳᵐ
	Queries       QueryIdːQueryʳᵐ
}

func (c *Objects) Init() {
	c.Renderbuffers = make(RenderbufferIdːRenderbufferʳᵐ)
	c.Textures = make(TextureIdːTextureʳᵐ)
	c.Framebuffers = make(FramebufferIdːFramebufferʳᵐ)
	c.Buffers = make(BufferIdːBufferʳᵐ)
	c.Shaders = make(ShaderIdːShaderʳᵐ)
	c.Programs = make(ProgramIdːProgramʳᵐ)
	c.VertexArrays = make(VertexArrayIdːVertexArrayʳᵐ)
	c.Queries = make(QueryIdːQueryʳᵐ)
}
func (c *Objects) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Context
////////////////////////////////////////////////////////////////////////////////
type Context struct {
	binary.Generate
	CreatedAt             atom.ID
	Identifier            ContextID
	Blending              BlendState
	Rasterizing           RasterizerState
	Clearing              ClearState
	BoundFramebuffers     FramebufferTargetːFramebufferIdᵐ
	BoundRenderbuffers    RenderbufferTargetːRenderbufferIdᵐ
	BoundBuffers          BufferTargetːBufferIdᵐ
	BoundProgram          ProgramId
	BoundVertexArray      VertexArrayId
	VertexAttributeArrays AttributeLocationːVertexAttributeArrayʳᵐ
	TextureUnits          TextureUnitːTextureTargetːTextureIdᵐᵐ
	ActiveTextureUnit     TextureUnit
	Capabilities          Capabilityːboolᵐ
	GenerateMipmapHint    HintMode
	PixelStorage          PixelStoreParameterːs32ᵐ
	Instances             Objects
}

func (c *Context) Init() {
	c.Blending.Init()
	c.Rasterizing.Init()
	c.Clearing.Init()
	c.BoundFramebuffers = make(FramebufferTargetːFramebufferIdᵐ)
	c.BoundRenderbuffers = make(RenderbufferTargetːRenderbufferIdᵐ)
	c.BoundBuffers = make(BufferTargetːBufferIdᵐ)
	c.VertexAttributeArrays = make(AttributeLocationːVertexAttributeArrayʳᵐ)
	c.TextureUnits = make(TextureUnitːTextureTargetːTextureIdᵐᵐ)
	c.ActiveTextureUnit = TextureUnit_GL_TEXTURE0
	c.Capabilities = make(Capabilityːboolᵐ)
	c.GenerateMipmapHint = HintMode_GL_DONT_CARE
	c.PixelStorage = make(PixelStoreParameterːs32ᵐ)
	c.Instances.Init()
}
func (c *Context) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// enum DrawMode
////////////////////////////////////////////////////////////////////////////////
type DrawMode uint32

const (
	DrawMode_GL_LINE_LOOP      = DrawMode(2)
	DrawMode_GL_LINE_STRIP     = DrawMode(3)
	DrawMode_GL_LINES          = DrawMode(1)
	DrawMode_GL_POINTS         = DrawMode(0)
	DrawMode_GL_TRIANGLE_FAN   = DrawMode(6)
	DrawMode_GL_TRIANGLE_STRIP = DrawMode(5)
	DrawMode_GL_TRIANGLES      = DrawMode(4)
)

////////////////////////////////////////////////////////////////////////////////
// enum IndicesType
////////////////////////////////////////////////////////////////////////////////
type IndicesType uint32

const (
	IndicesType_GL_UNSIGNED_BYTE  = IndicesType(5121)
	IndicesType_GL_UNSIGNED_SHORT = IndicesType(5123)
	IndicesType_GL_UNSIGNED_INT   = IndicesType(5125)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureTarget_GLES_1_1
////////////////////////////////////////////////////////////////////////////////
type TextureTarget_GLES_1_1 uint32

const (
	TextureTarget_GLES_1_1_GL_TEXTURE_2D = TextureTarget_GLES_1_1(3553)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureTarget_GLES_2_0
////////////////////////////////////////////////////////////////////////////////
type TextureTarget_GLES_2_0 uint32

const (
	TextureTarget_GLES_2_0_GL_TEXTURE_CUBE_MAP = TextureTarget_GLES_2_0(34067)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureTarget_OES_EGL_image_external
////////////////////////////////////////////////////////////////////////////////
type TextureTarget_OES_EGL_image_external uint32

const (
	TextureTarget_OES_EGL_image_external_GL_TEXTURE_EXTERNAL_OES = TextureTarget_OES_EGL_image_external(36197)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureTarget
////////////////////////////////////////////////////////////////////////////////
type TextureTarget uint32

const ()

// TextureTarget_GLES_1_1
const (
	TextureTarget_GL_TEXTURE_2D = TextureTarget(3553)
)

// TextureTarget_GLES_2_0
const (
	TextureTarget_GL_TEXTURE_CUBE_MAP = TextureTarget(34067)
)

// TextureTarget_OES_EGL_image_external
const (
	TextureTarget_GL_TEXTURE_EXTERNAL_OES = TextureTarget(36197)
)

////////////////////////////////////////////////////////////////////////////////
// enum CubeMapImageTarget
////////////////////////////////////////////////////////////////////////////////
type CubeMapImageTarget uint32

const (
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X = CubeMapImageTarget(34070)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = CubeMapImageTarget(34072)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = CubeMapImageTarget(34074)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X = CubeMapImageTarget(34069)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y = CubeMapImageTarget(34071)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z = CubeMapImageTarget(34073)
)

////////////////////////////////////////////////////////////////////////////////
// enum Texture2DImageTarget
////////////////////////////////////////////////////////////////////////////////
type Texture2DImageTarget uint32

const (
	Texture2DImageTarget_GL_TEXTURE_2D = Texture2DImageTarget(3553)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureImageTarget
////////////////////////////////////////////////////////////////////////////////
type TextureImageTarget uint32

const ()

// CubeMapImageTarget
const (
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X = TextureImageTarget(34070)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = TextureImageTarget(34072)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = TextureImageTarget(34074)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X = TextureImageTarget(34069)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y = TextureImageTarget(34071)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z = TextureImageTarget(34073)
)

// Texture2DImageTarget
const (
	TextureImageTarget_GL_TEXTURE_2D = TextureImageTarget(3553)
)

////////////////////////////////////////////////////////////////////////////////
// enum BaseTexelFormat
////////////////////////////////////////////////////////////////////////////////
type BaseTexelFormat uint32

const (
	BaseTexelFormat_GL_ALPHA = BaseTexelFormat(6406)
	BaseTexelFormat_GL_RGB   = BaseTexelFormat(6407)
	BaseTexelFormat_GL_RGBA  = BaseTexelFormat(6408)
)

////////////////////////////////////////////////////////////////////////////////
// enum TexelFormat_GLES_1_1
////////////////////////////////////////////////////////////////////////////////
type TexelFormat_GLES_1_1 uint32

const (
	TexelFormat_GLES_1_1_GL_LUMINANCE       = TexelFormat_GLES_1_1(6409)
	TexelFormat_GLES_1_1_GL_LUMINANCE_ALPHA = TexelFormat_GLES_1_1(6410)
)

// BaseTexelFormat
const (
	TexelFormat_GLES_1_1_GL_ALPHA = TexelFormat_GLES_1_1(6406)
	TexelFormat_GLES_1_1_GL_RGB   = TexelFormat_GLES_1_1(6407)
	TexelFormat_GLES_1_1_GL_RGBA  = TexelFormat_GLES_1_1(6408)
)

////////////////////////////////////////////////////////////////////////////////
// enum TexelFormat_GLES_3_0
////////////////////////////////////////////////////////////////////////////////
type TexelFormat_GLES_3_0 uint32

const (
	TexelFormat_GLES_3_0_GL_RED               = TexelFormat_GLES_3_0(6403)
	TexelFormat_GLES_3_0_GL_RED_INTEGER       = TexelFormat_GLES_3_0(36244)
	TexelFormat_GLES_3_0_GL_RG                = TexelFormat_GLES_3_0(33319)
	TexelFormat_GLES_3_0_GL_RG_INTEGER        = TexelFormat_GLES_3_0(33320)
	TexelFormat_GLES_3_0_GL_RGB_INTEGER       = TexelFormat_GLES_3_0(36248)
	TexelFormat_GLES_3_0_GL_RGBA_INTEGER      = TexelFormat_GLES_3_0(36249)
	TexelFormat_GLES_3_0_GL_DEPTH_COMPONENT   = TexelFormat_GLES_3_0(6402)
	TexelFormat_GLES_3_0_GL_DEPTH_COMPONENT16 = TexelFormat_GLES_3_0(33189)
	TexelFormat_GLES_3_0_GL_DEPTH_STENCIL     = TexelFormat_GLES_3_0(34041)
	TexelFormat_GLES_3_0_GL_DEPTH24_STENCIL8  = TexelFormat_GLES_3_0(35056)
)

////////////////////////////////////////////////////////////////////////////////
// enum TexelFormat
////////////////////////////////////////////////////////////////////////////////
type TexelFormat uint32

const ()

// TexelFormat_GLES_1_1
const (
	TexelFormat_GL_LUMINANCE       = TexelFormat(6409)
	TexelFormat_GL_LUMINANCE_ALPHA = TexelFormat(6410)
)

// BaseTexelFormat
const (
	TexelFormat_GL_ALPHA = TexelFormat(6406)
	TexelFormat_GL_RGB   = TexelFormat(6407)
	TexelFormat_GL_RGBA  = TexelFormat(6408)
)

// TexelFormat_GLES_3_0
const (
	TexelFormat_GL_RED               = TexelFormat(6403)
	TexelFormat_GL_RED_INTEGER       = TexelFormat(36244)
	TexelFormat_GL_RG                = TexelFormat(33319)
	TexelFormat_GL_RG_INTEGER        = TexelFormat(33320)
	TexelFormat_GL_RGB_INTEGER       = TexelFormat(36248)
	TexelFormat_GL_RGBA_INTEGER      = TexelFormat(36249)
	TexelFormat_GL_DEPTH_COMPONENT   = TexelFormat(6402)
	TexelFormat_GL_DEPTH_COMPONENT16 = TexelFormat(33189)
	TexelFormat_GL_DEPTH_STENCIL     = TexelFormat(34041)
	TexelFormat_GL_DEPTH24_STENCIL8  = TexelFormat(35056)
)

////////////////////////////////////////////////////////////////////////////////
// enum RenderbufferFormat
////////////////////////////////////////////////////////////////////////////////
type RenderbufferFormat uint32

const (
	RenderbufferFormat_GL_RGBA4             = RenderbufferFormat(32854)
	RenderbufferFormat_GL_RGB5_A1           = RenderbufferFormat(32855)
	RenderbufferFormat_GL_RGB565            = RenderbufferFormat(36194)
	RenderbufferFormat_GL_RGBA8             = RenderbufferFormat(32856)
	RenderbufferFormat_GL_DEPTH_COMPONENT16 = RenderbufferFormat(33189)
	RenderbufferFormat_GL_STENCIL_INDEX8    = RenderbufferFormat(36168)
)

////////////////////////////////////////////////////////////////////////////////
// enum Type_ARB_half_float_vertex
////////////////////////////////////////////////////////////////////////////////
type Type_ARB_half_float_vertex uint32

const (
	Type_ARB_half_float_vertex_GL_ARB_half_float_vertex = Type_ARB_half_float_vertex(5131)
)

////////////////////////////////////////////////////////////////////////////////
// enum Type_OES_vertex_half_float
////////////////////////////////////////////////////////////////////////////////
type Type_OES_vertex_half_float uint32

const (
	Type_OES_vertex_half_float_GL_HALF_FLOAT_OES = Type_OES_vertex_half_float(36193)
)

////////////////////////////////////////////////////////////////////////////////
// enum CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
////////////////////////////////////////////////////////////////////////////////
type CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture uint32

const (
	CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture_GL_ETC1_RGB8_OES = CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture(36196)
)

////////////////////////////////////////////////////////////////////////////////
// enum CompressedTexelFormat_AMD_compressed_ATC_texture
////////////////////////////////////////////////////////////////////////////////
type CompressedTexelFormat_AMD_compressed_ATC_texture uint32

const (
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGB_AMD                     = CompressedTexelFormat_AMD_compressed_ATC_texture(35986)
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD     = CompressedTexelFormat_AMD_compressed_ATC_texture(35987)
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = CompressedTexelFormat_AMD_compressed_ATC_texture(34798)
)

////////////////////////////////////////////////////////////////////////////////
// enum CompressedTexelFormat
////////////////////////////////////////////////////////////////////////////////
type CompressedTexelFormat uint32

const ()

// CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
const (
	CompressedTexelFormat_GL_ETC1_RGB8_OES = CompressedTexelFormat(36196)
)

// CompressedTexelFormat_AMD_compressed_ATC_texture
const (
	CompressedTexelFormat_GL_ATC_RGB_AMD                     = CompressedTexelFormat(35986)
	CompressedTexelFormat_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD     = CompressedTexelFormat(35987)
	CompressedTexelFormat_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = CompressedTexelFormat(34798)
)

////////////////////////////////////////////////////////////////////////////////
// enum ImageTexelFormat
////////////////////////////////////////////////////////////////////////////////
type ImageTexelFormat uint32

const ()

// TexelFormat
const ()

// TexelFormat_GLES_1_1
const (
	ImageTexelFormat_GL_LUMINANCE       = ImageTexelFormat(6409)
	ImageTexelFormat_GL_LUMINANCE_ALPHA = ImageTexelFormat(6410)
)

// BaseTexelFormat
const (
	ImageTexelFormat_GL_ALPHA = ImageTexelFormat(6406)
	ImageTexelFormat_GL_RGB   = ImageTexelFormat(6407)
	ImageTexelFormat_GL_RGBA  = ImageTexelFormat(6408)
)

// TexelFormat_GLES_3_0
const (
	ImageTexelFormat_GL_RED               = ImageTexelFormat(6403)
	ImageTexelFormat_GL_RED_INTEGER       = ImageTexelFormat(36244)
	ImageTexelFormat_GL_RG                = ImageTexelFormat(33319)
	ImageTexelFormat_GL_RG_INTEGER        = ImageTexelFormat(33320)
	ImageTexelFormat_GL_RGB_INTEGER       = ImageTexelFormat(36248)
	ImageTexelFormat_GL_RGBA_INTEGER      = ImageTexelFormat(36249)
	ImageTexelFormat_GL_DEPTH_COMPONENT   = ImageTexelFormat(6402)
	ImageTexelFormat_GL_DEPTH_COMPONENT16 = ImageTexelFormat(33189)
	ImageTexelFormat_GL_DEPTH_STENCIL     = ImageTexelFormat(34041)
	ImageTexelFormat_GL_DEPTH24_STENCIL8  = ImageTexelFormat(35056)
)

// CompressedTexelFormat
const ()

// CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
const (
	ImageTexelFormat_GL_ETC1_RGB8_OES = ImageTexelFormat(36196)
)

// CompressedTexelFormat_AMD_compressed_ATC_texture
const (
	ImageTexelFormat_GL_ATC_RGB_AMD                     = ImageTexelFormat(35986)
	ImageTexelFormat_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD     = ImageTexelFormat(35987)
	ImageTexelFormat_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = ImageTexelFormat(34798)
)

////////////////////////////////////////////////////////////////////////////////
// enum TexelType
////////////////////////////////////////////////////////////////////////////////
type TexelType uint32

const (
	TexelType_GL_UNSIGNED_BYTE          = TexelType(5121)
	TexelType_GL_UNSIGNED_SHORT         = TexelType(5123)
	TexelType_GL_UNSIGNED_INT           = TexelType(5125)
	TexelType_GL_FLOAT                  = TexelType(5126)
	TexelType_GL_UNSIGNED_SHORT_4_4_4_4 = TexelType(32819)
	TexelType_GL_UNSIGNED_SHORT_5_5_5_1 = TexelType(32820)
	TexelType_GL_UNSIGNED_SHORT_5_6_5   = TexelType(33635)
	TexelType_GL_UNSIGNED_INT_24_8      = TexelType(34042)
)

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferAttachment
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachment uint32

const (
	FramebufferAttachment_GL_COLOR_ATTACHMENT0  = FramebufferAttachment(36064)
	FramebufferAttachment_GL_DEPTH_ATTACHMENT   = FramebufferAttachment(36096)
	FramebufferAttachment_GL_STENCIL_ATTACHMENT = FramebufferAttachment(36128)
)

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferAttachmentType
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachmentType uint32

const (
	FramebufferAttachmentType_GL_NONE         = FramebufferAttachmentType(0)
	FramebufferAttachmentType_GL_RENDERBUFFER = FramebufferAttachmentType(36161)
	FramebufferAttachmentType_GL_TEXTURE      = FramebufferAttachmentType(5890)
)

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferTarget_GLES_2_0
////////////////////////////////////////////////////////////////////////////////
type FramebufferTarget_GLES_2_0 uint32

const (
	FramebufferTarget_GLES_2_0_GL_FRAMEBUFFER = FramebufferTarget_GLES_2_0(36160)
)

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferTarget_GLES_3_1
////////////////////////////////////////////////////////////////////////////////
type FramebufferTarget_GLES_3_1 uint32

const (
	FramebufferTarget_GLES_3_1_GL_READ_FRAMEBUFFER = FramebufferTarget_GLES_3_1(36008)
	FramebufferTarget_GLES_3_1_GL_DRAW_FRAMEBUFFER = FramebufferTarget_GLES_3_1(36009)
)

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferTarget
////////////////////////////////////////////////////////////////////////////////
type FramebufferTarget uint32

const ()

// FramebufferTarget_GLES_2_0
const (
	FramebufferTarget_GL_FRAMEBUFFER = FramebufferTarget(36160)
)

// FramebufferTarget_GLES_3_1
const (
	FramebufferTarget_GL_READ_FRAMEBUFFER = FramebufferTarget(36008)
	FramebufferTarget_GL_DRAW_FRAMEBUFFER = FramebufferTarget(36009)
)

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferAttachmentParameter
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachmentParameter uint32

const (
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = FramebufferAttachmentParameter(36048)
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = FramebufferAttachmentParameter(36049)
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = FramebufferAttachmentParameter(36050)
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = FramebufferAttachmentParameter(36051)
)

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferStatus
////////////////////////////////////////////////////////////////////////////////
type FramebufferStatus uint32

const (
	FramebufferStatus_GL_FRAMEBUFFER_COMPLETE                      = FramebufferStatus(36053)
	FramebufferStatus_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT         = FramebufferStatus(36054)
	FramebufferStatus_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = FramebufferStatus(36055)
	FramebufferStatus_GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS         = FramebufferStatus(36057)
	FramebufferStatus_GL_FRAMEBUFFER_UNSUPPORTED                   = FramebufferStatus(36061)
)

////////////////////////////////////////////////////////////////////////////////
// enum RenderbufferTarget
////////////////////////////////////////////////////////////////////////////////
type RenderbufferTarget uint32

const (
	RenderbufferTarget_GL_RENDERBUFFER = RenderbufferTarget(36161)
)

////////////////////////////////////////////////////////////////////////////////
// enum RenderbufferParameter
////////////////////////////////////////////////////////////////////////////////
type RenderbufferParameter uint32

const (
	RenderbufferParameter_GL_RENDERBUFFER_WIDTH           = RenderbufferParameter(36162)
	RenderbufferParameter_GL_RENDERBUFFER_HEIGHT          = RenderbufferParameter(36163)
	RenderbufferParameter_GL_RENDERBUFFER_INTERNAL_FORMAT = RenderbufferParameter(36164)
	RenderbufferParameter_GL_RENDERBUFFER_RED_SIZE        = RenderbufferParameter(36176)
	RenderbufferParameter_GL_RENDERBUFFER_GREEN_SIZE      = RenderbufferParameter(36177)
	RenderbufferParameter_GL_RENDERBUFFER_BLUE_SIZE       = RenderbufferParameter(36178)
	RenderbufferParameter_GL_RENDERBUFFER_ALPHA_SIZE      = RenderbufferParameter(36179)
	RenderbufferParameter_GL_RENDERBUFFER_DEPTH_SIZE      = RenderbufferParameter(36180)
	RenderbufferParameter_GL_RENDERBUFFER_STENCIL_SIZE    = RenderbufferParameter(36181)
)

////////////////////////////////////////////////////////////////////////////////
// enum BufferParameter
////////////////////////////////////////////////////////////////////////////////
type BufferParameter uint32

const (
	BufferParameter_GL_BUFFER_SIZE  = BufferParameter(34660)
	BufferParameter_GL_BUFFER_USAGE = BufferParameter(34661)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureUnit
////////////////////////////////////////////////////////////////////////////////
type TextureUnit uint32

const (
	TextureUnit_GL_TEXTURE0  = TextureUnit(33984)
	TextureUnit_GL_TEXTURE1  = TextureUnit(33985)
	TextureUnit_GL_TEXTURE2  = TextureUnit(33986)
	TextureUnit_GL_TEXTURE3  = TextureUnit(33987)
	TextureUnit_GL_TEXTURE4  = TextureUnit(33988)
	TextureUnit_GL_TEXTURE5  = TextureUnit(33989)
	TextureUnit_GL_TEXTURE6  = TextureUnit(33990)
	TextureUnit_GL_TEXTURE7  = TextureUnit(33991)
	TextureUnit_GL_TEXTURE8  = TextureUnit(33992)
	TextureUnit_GL_TEXTURE9  = TextureUnit(33993)
	TextureUnit_GL_TEXTURE10 = TextureUnit(33994)
	TextureUnit_GL_TEXTURE11 = TextureUnit(33995)
	TextureUnit_GL_TEXTURE12 = TextureUnit(33996)
	TextureUnit_GL_TEXTURE13 = TextureUnit(33997)
	TextureUnit_GL_TEXTURE14 = TextureUnit(33998)
	TextureUnit_GL_TEXTURE15 = TextureUnit(33999)
	TextureUnit_GL_TEXTURE16 = TextureUnit(34000)
	TextureUnit_GL_TEXTURE17 = TextureUnit(34001)
	TextureUnit_GL_TEXTURE18 = TextureUnit(34002)
	TextureUnit_GL_TEXTURE19 = TextureUnit(34003)
	TextureUnit_GL_TEXTURE20 = TextureUnit(34004)
	TextureUnit_GL_TEXTURE21 = TextureUnit(34005)
	TextureUnit_GL_TEXTURE22 = TextureUnit(34006)
	TextureUnit_GL_TEXTURE23 = TextureUnit(34007)
	TextureUnit_GL_TEXTURE24 = TextureUnit(34008)
	TextureUnit_GL_TEXTURE25 = TextureUnit(34009)
	TextureUnit_GL_TEXTURE26 = TextureUnit(34010)
	TextureUnit_GL_TEXTURE27 = TextureUnit(34011)
	TextureUnit_GL_TEXTURE28 = TextureUnit(34012)
	TextureUnit_GL_TEXTURE29 = TextureUnit(34013)
	TextureUnit_GL_TEXTURE30 = TextureUnit(34014)
	TextureUnit_GL_TEXTURE31 = TextureUnit(34015)
)

////////////////////////////////////////////////////////////////////////////////
// enum BufferUsage
////////////////////////////////////////////////////////////////////////////////
type BufferUsage uint32

const (
	BufferUsage_GL_DYNAMIC_DRAW = BufferUsage(35048)
	BufferUsage_GL_STATIC_DRAW  = BufferUsage(35044)
	BufferUsage_GL_STREAM_DRAW  = BufferUsage(35040)
)

////////////////////////////////////////////////////////////////////////////////
// enum ShaderType
////////////////////////////////////////////////////////////////////////////////
type ShaderType uint32

const (
	ShaderType_GL_VERTEX_SHADER   = ShaderType(35633)
	ShaderType_GL_FRAGMENT_SHADER = ShaderType(35632)
)

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable_GLES_2_0
////////////////////////////////////////////////////////////////////////////////
type StateVariable_GLES_2_0 uint32

const (
	StateVariable_GLES_2_0_GL_ACTIVE_TEXTURE                   = StateVariable_GLES_2_0(34016)
	StateVariable_GLES_2_0_GL_ALIASED_LINE_WIDTH_RANGE         = StateVariable_GLES_2_0(33902)
	StateVariable_GLES_2_0_GL_ALIASED_POINT_SIZE_RANGE         = StateVariable_GLES_2_0(33901)
	StateVariable_GLES_2_0_GL_ALPHA_BITS                       = StateVariable_GLES_2_0(3413)
	StateVariable_GLES_2_0_GL_ARRAY_BUFFER_BINDING             = StateVariable_GLES_2_0(34964)
	StateVariable_GLES_2_0_GL_BLEND                            = StateVariable_GLES_2_0(3042)
	StateVariable_GLES_2_0_GL_BLEND_COLOR                      = StateVariable_GLES_2_0(32773)
	StateVariable_GLES_2_0_GL_BLEND_DST_ALPHA                  = StateVariable_GLES_2_0(32970)
	StateVariable_GLES_2_0_GL_BLEND_DST_RGB                    = StateVariable_GLES_2_0(32968)
	StateVariable_GLES_2_0_GL_BLEND_EQUATION_ALPHA             = StateVariable_GLES_2_0(34877)
	StateVariable_GLES_2_0_GL_BLEND_EQUATION_RGB               = StateVariable_GLES_2_0(32777)
	StateVariable_GLES_2_0_GL_BLEND_SRC_ALPHA                  = StateVariable_GLES_2_0(32971)
	StateVariable_GLES_2_0_GL_BLEND_SRC_RGB                    = StateVariable_GLES_2_0(32969)
	StateVariable_GLES_2_0_GL_BLUE_BITS                        = StateVariable_GLES_2_0(3412)
	StateVariable_GLES_2_0_GL_COLOR_CLEAR_VALUE                = StateVariable_GLES_2_0(3106)
	StateVariable_GLES_2_0_GL_COLOR_WRITEMASK                  = StateVariable_GLES_2_0(3107)
	StateVariable_GLES_2_0_GL_COMPRESSED_TEXTURE_FORMATS       = StateVariable_GLES_2_0(34467)
	StateVariable_GLES_2_0_GL_CULL_FACE                        = StateVariable_GLES_2_0(2884)
	StateVariable_GLES_2_0_GL_CULL_FACE_MODE                   = StateVariable_GLES_2_0(2885)
	StateVariable_GLES_2_0_GL_CURRENT_PROGRAM                  = StateVariable_GLES_2_0(35725)
	StateVariable_GLES_2_0_GL_DEPTH_BITS                       = StateVariable_GLES_2_0(3414)
	StateVariable_GLES_2_0_GL_DEPTH_CLEAR_VALUE                = StateVariable_GLES_2_0(2931)
	StateVariable_GLES_2_0_GL_DEPTH_FUNC                       = StateVariable_GLES_2_0(2932)
	StateVariable_GLES_2_0_GL_DEPTH_RANGE                      = StateVariable_GLES_2_0(2928)
	StateVariable_GLES_2_0_GL_DEPTH_TEST                       = StateVariable_GLES_2_0(2929)
	StateVariable_GLES_2_0_GL_DEPTH_WRITEMASK                  = StateVariable_GLES_2_0(2930)
	StateVariable_GLES_2_0_GL_DITHER                           = StateVariable_GLES_2_0(3024)
	StateVariable_GLES_2_0_GL_ELEMENT_ARRAY_BUFFER_BINDING     = StateVariable_GLES_2_0(34965)
	StateVariable_GLES_2_0_GL_FRAMEBUFFER_BINDING              = StateVariable_GLES_2_0(36006)
	StateVariable_GLES_2_0_GL_FRONT_FACE                       = StateVariable_GLES_2_0(2886)
	StateVariable_GLES_2_0_GL_GENERATE_MIPMAP_HINT             = StateVariable_GLES_2_0(33170)
	StateVariable_GLES_2_0_GL_GREEN_BITS                       = StateVariable_GLES_2_0(3411)
	StateVariable_GLES_2_0_GL_IMPLEMENTATION_COLOR_READ_FORMAT = StateVariable_GLES_2_0(35739)
	StateVariable_GLES_2_0_GL_IMPLEMENTATION_COLOR_READ_TYPE   = StateVariable_GLES_2_0(35738)
	StateVariable_GLES_2_0_GL_LINE_WIDTH                       = StateVariable_GLES_2_0(2849)
	StateVariable_GLES_2_0_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS = StateVariable_GLES_2_0(35661)
	StateVariable_GLES_2_0_GL_MAX_CUBE_MAP_TEXTURE_SIZE        = StateVariable_GLES_2_0(34076)
	StateVariable_GLES_2_0_GL_MAX_FRAGMENT_UNIFORM_VECTORS     = StateVariable_GLES_2_0(36349)
	StateVariable_GLES_2_0_GL_MAX_RENDERBUFFER_SIZE            = StateVariable_GLES_2_0(34024)
	StateVariable_GLES_2_0_GL_MAX_TEXTURE_IMAGE_UNITS          = StateVariable_GLES_2_0(34930)
	StateVariable_GLES_2_0_GL_MAX_TEXTURE_SIZE                 = StateVariable_GLES_2_0(3379)
	StateVariable_GLES_2_0_GL_MAX_VARYING_VECTORS              = StateVariable_GLES_2_0(36348)
	StateVariable_GLES_2_0_GL_MAX_VERTEX_ATTRIBS               = StateVariable_GLES_2_0(34921)
	StateVariable_GLES_2_0_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS   = StateVariable_GLES_2_0(35660)
	StateVariable_GLES_2_0_GL_MAX_VERTEX_UNIFORM_VECTORS       = StateVariable_GLES_2_0(36347)
	StateVariable_GLES_2_0_GL_MAX_VIEWPORT_DIMS                = StateVariable_GLES_2_0(3386)
	StateVariable_GLES_2_0_GL_NUM_COMPRESSED_TEXTURE_FORMATS   = StateVariable_GLES_2_0(34466)
	StateVariable_GLES_2_0_GL_NUM_SHADER_BINARY_FORMATS        = StateVariable_GLES_2_0(36345)
	StateVariable_GLES_2_0_GL_PACK_ALIGNMENT                   = StateVariable_GLES_2_0(3333)
	StateVariable_GLES_2_0_GL_POLYGON_OFFSET_FACTOR            = StateVariable_GLES_2_0(32824)
	StateVariable_GLES_2_0_GL_POLYGON_OFFSET_FILL              = StateVariable_GLES_2_0(32823)
	StateVariable_GLES_2_0_GL_POLYGON_OFFSET_UNITS             = StateVariable_GLES_2_0(10752)
	StateVariable_GLES_2_0_GL_RED_BITS                         = StateVariable_GLES_2_0(3410)
	StateVariable_GLES_2_0_GL_RENDERBUFFER_BINDING             = StateVariable_GLES_2_0(36007)
	StateVariable_GLES_2_0_GL_SAMPLE_ALPHA_TO_COVERAGE         = StateVariable_GLES_2_0(32926)
	StateVariable_GLES_2_0_GL_SAMPLE_BUFFERS                   = StateVariable_GLES_2_0(32936)
	StateVariable_GLES_2_0_GL_SAMPLE_COVERAGE                  = StateVariable_GLES_2_0(32928)
	StateVariable_GLES_2_0_GL_SAMPLE_COVERAGE_INVERT           = StateVariable_GLES_2_0(32939)
	StateVariable_GLES_2_0_GL_SAMPLE_COVERAGE_VALUE            = StateVariable_GLES_2_0(32938)
	StateVariable_GLES_2_0_GL_SAMPLES                          = StateVariable_GLES_2_0(32937)
	StateVariable_GLES_2_0_GL_SCISSOR_BOX                      = StateVariable_GLES_2_0(3088)
	StateVariable_GLES_2_0_GL_SCISSOR_TEST                     = StateVariable_GLES_2_0(3089)
	StateVariable_GLES_2_0_GL_SHADER_BINARY_FORMATS            = StateVariable_GLES_2_0(36344)
	StateVariable_GLES_2_0_GL_SHADER_COMPILER                  = StateVariable_GLES_2_0(36346)
	StateVariable_GLES_2_0_GL_STENCIL_BACK_FAIL                = StateVariable_GLES_2_0(34817)
	StateVariable_GLES_2_0_GL_STENCIL_BACK_FUNC                = StateVariable_GLES_2_0(34816)
	StateVariable_GLES_2_0_GL_STENCIL_BACK_PASS_DEPTH_FAIL     = StateVariable_GLES_2_0(34818)
	StateVariable_GLES_2_0_GL_STENCIL_BACK_PASS_DEPTH_PASS     = StateVariable_GLES_2_0(34819)
	StateVariable_GLES_2_0_GL_STENCIL_BACK_REF                 = StateVariable_GLES_2_0(36003)
	StateVariable_GLES_2_0_GL_STENCIL_BACK_VALUE_MASK          = StateVariable_GLES_2_0(36004)
	StateVariable_GLES_2_0_GL_STENCIL_BACK_WRITEMASK           = StateVariable_GLES_2_0(36005)
	StateVariable_GLES_2_0_GL_STENCIL_BITS                     = StateVariable_GLES_2_0(3415)
	StateVariable_GLES_2_0_GL_STENCIL_CLEAR_VALUE              = StateVariable_GLES_2_0(2961)
	StateVariable_GLES_2_0_GL_STENCIL_FAIL                     = StateVariable_GLES_2_0(2964)
	StateVariable_GLES_2_0_GL_STENCIL_FUNC                     = StateVariable_GLES_2_0(2962)
	StateVariable_GLES_2_0_GL_STENCIL_PASS_DEPTH_FAIL          = StateVariable_GLES_2_0(2965)
	StateVariable_GLES_2_0_GL_STENCIL_PASS_DEPTH_PASS          = StateVariable_GLES_2_0(2966)
	StateVariable_GLES_2_0_GL_STENCIL_REF                      = StateVariable_GLES_2_0(2967)
	StateVariable_GLES_2_0_GL_STENCIL_TEST                     = StateVariable_GLES_2_0(2960)
	StateVariable_GLES_2_0_GL_STENCIL_VALUE_MASK               = StateVariable_GLES_2_0(2963)
	StateVariable_GLES_2_0_GL_STENCIL_WRITEMASK                = StateVariable_GLES_2_0(2968)
	StateVariable_GLES_2_0_GL_SUBPIXEL_BITS                    = StateVariable_GLES_2_0(3408)
	StateVariable_GLES_2_0_GL_TEXTURE_BINDING_2D               = StateVariable_GLES_2_0(32873)
	StateVariable_GLES_2_0_GL_TEXTURE_BINDING_CUBE_MAP         = StateVariable_GLES_2_0(34068)
	StateVariable_GLES_2_0_GL_UNPACK_ALIGNMENT                 = StateVariable_GLES_2_0(3317)
	StateVariable_GLES_2_0_GL_VIEWPORT                         = StateVariable_GLES_2_0(2978)
)

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable_GLES_3_1
////////////////////////////////////////////////////////////////////////////////
type StateVariable_GLES_3_1 uint32

const (
	StateVariable_GLES_3_1_GL_READ_FRAMEBUFFER_BINDING = StateVariable_GLES_3_1(36010)
)

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable_EXT_texture_filter_anisotropic
////////////////////////////////////////////////////////////////////////////////
type StateVariable_EXT_texture_filter_anisotropic uint32

const (
	StateVariable_EXT_texture_filter_anisotropic_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = StateVariable_EXT_texture_filter_anisotropic(34047)
)

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type StateVariable_EXT_disjoint_timer_query uint32

const (
	StateVariable_EXT_disjoint_timer_query_GL_GPU_DISJOINT_EXT = StateVariable_EXT_disjoint_timer_query(36795)
)

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable
////////////////////////////////////////////////////////////////////////////////
type StateVariable uint32

const ()

// StateVariable_GLES_2_0
const (
	StateVariable_GL_ACTIVE_TEXTURE                   = StateVariable(34016)
	StateVariable_GL_ALIASED_LINE_WIDTH_RANGE         = StateVariable(33902)
	StateVariable_GL_ALIASED_POINT_SIZE_RANGE         = StateVariable(33901)
	StateVariable_GL_ALPHA_BITS                       = StateVariable(3413)
	StateVariable_GL_ARRAY_BUFFER_BINDING             = StateVariable(34964)
	StateVariable_GL_BLEND                            = StateVariable(3042)
	StateVariable_GL_BLEND_COLOR                      = StateVariable(32773)
	StateVariable_GL_BLEND_DST_ALPHA                  = StateVariable(32970)
	StateVariable_GL_BLEND_DST_RGB                    = StateVariable(32968)
	StateVariable_GL_BLEND_EQUATION_ALPHA             = StateVariable(34877)
	StateVariable_GL_BLEND_EQUATION_RGB               = StateVariable(32777)
	StateVariable_GL_BLEND_SRC_ALPHA                  = StateVariable(32971)
	StateVariable_GL_BLEND_SRC_RGB                    = StateVariable(32969)
	StateVariable_GL_BLUE_BITS                        = StateVariable(3412)
	StateVariable_GL_COLOR_CLEAR_VALUE                = StateVariable(3106)
	StateVariable_GL_COLOR_WRITEMASK                  = StateVariable(3107)
	StateVariable_GL_COMPRESSED_TEXTURE_FORMATS       = StateVariable(34467)
	StateVariable_GL_CULL_FACE                        = StateVariable(2884)
	StateVariable_GL_CULL_FACE_MODE                   = StateVariable(2885)
	StateVariable_GL_CURRENT_PROGRAM                  = StateVariable(35725)
	StateVariable_GL_DEPTH_BITS                       = StateVariable(3414)
	StateVariable_GL_DEPTH_CLEAR_VALUE                = StateVariable(2931)
	StateVariable_GL_DEPTH_FUNC                       = StateVariable(2932)
	StateVariable_GL_DEPTH_RANGE                      = StateVariable(2928)
	StateVariable_GL_DEPTH_TEST                       = StateVariable(2929)
	StateVariable_GL_DEPTH_WRITEMASK                  = StateVariable(2930)
	StateVariable_GL_DITHER                           = StateVariable(3024)
	StateVariable_GL_ELEMENT_ARRAY_BUFFER_BINDING     = StateVariable(34965)
	StateVariable_GL_FRAMEBUFFER_BINDING              = StateVariable(36006)
	StateVariable_GL_FRONT_FACE                       = StateVariable(2886)
	StateVariable_GL_GENERATE_MIPMAP_HINT             = StateVariable(33170)
	StateVariable_GL_GREEN_BITS                       = StateVariable(3411)
	StateVariable_GL_IMPLEMENTATION_COLOR_READ_FORMAT = StateVariable(35739)
	StateVariable_GL_IMPLEMENTATION_COLOR_READ_TYPE   = StateVariable(35738)
	StateVariable_GL_LINE_WIDTH                       = StateVariable(2849)
	StateVariable_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS = StateVariable(35661)
	StateVariable_GL_MAX_CUBE_MAP_TEXTURE_SIZE        = StateVariable(34076)
	StateVariable_GL_MAX_FRAGMENT_UNIFORM_VECTORS     = StateVariable(36349)
	StateVariable_GL_MAX_RENDERBUFFER_SIZE            = StateVariable(34024)
	StateVariable_GL_MAX_TEXTURE_IMAGE_UNITS          = StateVariable(34930)
	StateVariable_GL_MAX_TEXTURE_SIZE                 = StateVariable(3379)
	StateVariable_GL_MAX_VARYING_VECTORS              = StateVariable(36348)
	StateVariable_GL_MAX_VERTEX_ATTRIBS               = StateVariable(34921)
	StateVariable_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS   = StateVariable(35660)
	StateVariable_GL_MAX_VERTEX_UNIFORM_VECTORS       = StateVariable(36347)
	StateVariable_GL_MAX_VIEWPORT_DIMS                = StateVariable(3386)
	StateVariable_GL_NUM_COMPRESSED_TEXTURE_FORMATS   = StateVariable(34466)
	StateVariable_GL_NUM_SHADER_BINARY_FORMATS        = StateVariable(36345)
	StateVariable_GL_PACK_ALIGNMENT                   = StateVariable(3333)
	StateVariable_GL_POLYGON_OFFSET_FACTOR            = StateVariable(32824)
	StateVariable_GL_POLYGON_OFFSET_FILL              = StateVariable(32823)
	StateVariable_GL_POLYGON_OFFSET_UNITS             = StateVariable(10752)
	StateVariable_GL_RED_BITS                         = StateVariable(3410)
	StateVariable_GL_RENDERBUFFER_BINDING             = StateVariable(36007)
	StateVariable_GL_SAMPLE_ALPHA_TO_COVERAGE         = StateVariable(32926)
	StateVariable_GL_SAMPLE_BUFFERS                   = StateVariable(32936)
	StateVariable_GL_SAMPLE_COVERAGE                  = StateVariable(32928)
	StateVariable_GL_SAMPLE_COVERAGE_INVERT           = StateVariable(32939)
	StateVariable_GL_SAMPLE_COVERAGE_VALUE            = StateVariable(32938)
	StateVariable_GL_SAMPLES                          = StateVariable(32937)
	StateVariable_GL_SCISSOR_BOX                      = StateVariable(3088)
	StateVariable_GL_SCISSOR_TEST                     = StateVariable(3089)
	StateVariable_GL_SHADER_BINARY_FORMATS            = StateVariable(36344)
	StateVariable_GL_SHADER_COMPILER                  = StateVariable(36346)
	StateVariable_GL_STENCIL_BACK_FAIL                = StateVariable(34817)
	StateVariable_GL_STENCIL_BACK_FUNC                = StateVariable(34816)
	StateVariable_GL_STENCIL_BACK_PASS_DEPTH_FAIL     = StateVariable(34818)
	StateVariable_GL_STENCIL_BACK_PASS_DEPTH_PASS     = StateVariable(34819)
	StateVariable_GL_STENCIL_BACK_REF                 = StateVariable(36003)
	StateVariable_GL_STENCIL_BACK_VALUE_MASK          = StateVariable(36004)
	StateVariable_GL_STENCIL_BACK_WRITEMASK           = StateVariable(36005)
	StateVariable_GL_STENCIL_BITS                     = StateVariable(3415)
	StateVariable_GL_STENCIL_CLEAR_VALUE              = StateVariable(2961)
	StateVariable_GL_STENCIL_FAIL                     = StateVariable(2964)
	StateVariable_GL_STENCIL_FUNC                     = StateVariable(2962)
	StateVariable_GL_STENCIL_PASS_DEPTH_FAIL          = StateVariable(2965)
	StateVariable_GL_STENCIL_PASS_DEPTH_PASS          = StateVariable(2966)
	StateVariable_GL_STENCIL_REF                      = StateVariable(2967)
	StateVariable_GL_STENCIL_TEST                     = StateVariable(2960)
	StateVariable_GL_STENCIL_VALUE_MASK               = StateVariable(2963)
	StateVariable_GL_STENCIL_WRITEMASK                = StateVariable(2968)
	StateVariable_GL_SUBPIXEL_BITS                    = StateVariable(3408)
	StateVariable_GL_TEXTURE_BINDING_2D               = StateVariable(32873)
	StateVariable_GL_TEXTURE_BINDING_CUBE_MAP         = StateVariable(34068)
	StateVariable_GL_UNPACK_ALIGNMENT                 = StateVariable(3317)
	StateVariable_GL_VIEWPORT                         = StateVariable(2978)
)

// StateVariable_GLES_3_1
const (
	StateVariable_GL_READ_FRAMEBUFFER_BINDING = StateVariable(36010)
)

// StateVariable_EXT_texture_filter_anisotropic
const (
	StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = StateVariable(34047)
)

// StateVariable_EXT_disjoint_timer_query
const (
	StateVariable_GL_GPU_DISJOINT_EXT = StateVariable(36795)
)

////////////////////////////////////////////////////////////////////////////////
// enum FaceMode
////////////////////////////////////////////////////////////////////////////////
type FaceMode uint32

const (
	FaceMode_GL_FRONT          = FaceMode(1028)
	FaceMode_GL_BACK           = FaceMode(1029)
	FaceMode_GL_FRONT_AND_BACK = FaceMode(1032)
)

////////////////////////////////////////////////////////////////////////////////
// enum ArrayType_GLES_1_1
////////////////////////////////////////////////////////////////////////////////
type ArrayType_GLES_1_1 uint32

const (
	ArrayType_GLES_1_1_GL_VERTEX_ARRAY        = ArrayType_GLES_1_1(32884)
	ArrayType_GLES_1_1_GL_NORMAL_ARRAY        = ArrayType_GLES_1_1(32885)
	ArrayType_GLES_1_1_GL_COLOR_ARRAY         = ArrayType_GLES_1_1(32886)
	ArrayType_GLES_1_1_GL_TEXTURE_COORD_ARRAY = ArrayType_GLES_1_1(32888)
)

////////////////////////////////////////////////////////////////////////////////
// enum ArrayType_OES_point_size_array
////////////////////////////////////////////////////////////////////////////////
type ArrayType_OES_point_size_array uint32

const (
	ArrayType_OES_point_size_array_GL_POINT_SIZE_ARRAY_OES = ArrayType_OES_point_size_array(35740)
)

////////////////////////////////////////////////////////////////////////////////
// enum ArrayType
////////////////////////////////////////////////////////////////////////////////
type ArrayType uint32

const ()

// ArrayType_GLES_1_1
const (
	ArrayType_GL_VERTEX_ARRAY        = ArrayType(32884)
	ArrayType_GL_NORMAL_ARRAY        = ArrayType(32885)
	ArrayType_GL_COLOR_ARRAY         = ArrayType(32886)
	ArrayType_GL_TEXTURE_COORD_ARRAY = ArrayType(32888)
)

// ArrayType_OES_point_size_array
const (
	ArrayType_GL_POINT_SIZE_ARRAY_OES = ArrayType(35740)
)

////////////////////////////////////////////////////////////////////////////////
// enum Capability
////////////////////////////////////////////////////////////////////////////////
type Capability uint32

const (
	Capability_GL_BLEND                    = Capability(3042)
	Capability_GL_CULL_FACE                = Capability(2884)
	Capability_GL_DEPTH_TEST               = Capability(2929)
	Capability_GL_DITHER                   = Capability(3024)
	Capability_GL_POLYGON_OFFSET_FILL      = Capability(32823)
	Capability_GL_SAMPLE_ALPHA_TO_COVERAGE = Capability(32926)
	Capability_GL_SAMPLE_COVERAGE          = Capability(32928)
	Capability_GL_SCISSOR_TEST             = Capability(3089)
	Capability_GL_STENCIL_TEST             = Capability(2960)
)

// ArrayType
const ()

// ArrayType_GLES_1_1
const (
	Capability_GL_VERTEX_ARRAY        = Capability(32884)
	Capability_GL_NORMAL_ARRAY        = Capability(32885)
	Capability_GL_COLOR_ARRAY         = Capability(32886)
	Capability_GL_TEXTURE_COORD_ARRAY = Capability(32888)
)

// ArrayType_OES_point_size_array
const (
	Capability_GL_POINT_SIZE_ARRAY_OES = Capability(35740)
)

////////////////////////////////////////////////////////////////////////////////
// enum StringConstant
////////////////////////////////////////////////////////////////////////////////
type StringConstant uint32

const (
	StringConstant_GL_EXTENSIONS = StringConstant(7939)
	StringConstant_GL_RENDERER   = StringConstant(7937)
	StringConstant_GL_VENDOR     = StringConstant(7936)
	StringConstant_GL_VERSION    = StringConstant(7938)
)

////////////////////////////////////////////////////////////////////////////////
// enum VertexAttribType
////////////////////////////////////////////////////////////////////////////////
type VertexAttribType uint32

const (
	VertexAttribType_GL_BYTE           = VertexAttribType(5120)
	VertexAttribType_GL_FIXED          = VertexAttribType(5132)
	VertexAttribType_GL_FLOAT          = VertexAttribType(5126)
	VertexAttribType_GL_SHORT          = VertexAttribType(5122)
	VertexAttribType_GL_UNSIGNED_BYTE  = VertexAttribType(5121)
	VertexAttribType_GL_UNSIGNED_SHORT = VertexAttribType(5123)
)

// Type_OES_vertex_half_float
const (
	VertexAttribType_GL_HALF_FLOAT_OES = VertexAttribType(36193)
)

// Type_ARB_half_float_vertex
const (
	VertexAttribType_GL_ARB_half_float_vertex = VertexAttribType(5131)
)

////////////////////////////////////////////////////////////////////////////////
// enum ShaderAttribType
////////////////////////////////////////////////////////////////////////////////
type ShaderAttribType uint32

const (
	ShaderAttribType_GL_FLOAT      = ShaderAttribType(5126)
	ShaderAttribType_GL_FLOAT_VEC2 = ShaderAttribType(35664)
	ShaderAttribType_GL_FLOAT_VEC3 = ShaderAttribType(35665)
	ShaderAttribType_GL_FLOAT_VEC4 = ShaderAttribType(35666)
	ShaderAttribType_GL_FLOAT_MAT2 = ShaderAttribType(35674)
	ShaderAttribType_GL_FLOAT_MAT3 = ShaderAttribType(35675)
	ShaderAttribType_GL_FLOAT_MAT4 = ShaderAttribType(35676)
)

////////////////////////////////////////////////////////////////////////////////
// enum ShaderUniformType
////////////////////////////////////////////////////////////////////////////////
type ShaderUniformType uint32

const (
	ShaderUniformType_GL_FLOAT        = ShaderUniformType(5126)
	ShaderUniformType_GL_FLOAT_VEC2   = ShaderUniformType(35664)
	ShaderUniformType_GL_FLOAT_VEC3   = ShaderUniformType(35665)
	ShaderUniformType_GL_FLOAT_VEC4   = ShaderUniformType(35666)
	ShaderUniformType_GL_INT          = ShaderUniformType(5124)
	ShaderUniformType_GL_INT_VEC2     = ShaderUniformType(35667)
	ShaderUniformType_GL_INT_VEC3     = ShaderUniformType(35668)
	ShaderUniformType_GL_INT_VEC4     = ShaderUniformType(35669)
	ShaderUniformType_GL_BOOL         = ShaderUniformType(35670)
	ShaderUniformType_GL_BOOL_VEC2    = ShaderUniformType(35671)
	ShaderUniformType_GL_BOOL_VEC3    = ShaderUniformType(35672)
	ShaderUniformType_GL_BOOL_VEC4    = ShaderUniformType(35673)
	ShaderUniformType_GL_FLOAT_MAT2   = ShaderUniformType(35674)
	ShaderUniformType_GL_FLOAT_MAT3   = ShaderUniformType(35675)
	ShaderUniformType_GL_FLOAT_MAT4   = ShaderUniformType(35676)
	ShaderUniformType_GL_SAMPLER_2D   = ShaderUniformType(35678)
	ShaderUniformType_GL_SAMPLER_CUBE = ShaderUniformType(35680)
)

////////////////////////////////////////////////////////////////////////////////
// enum Error
////////////////////////////////////////////////////////////////////////////////
type Error uint32

const (
	Error_GL_NO_ERROR                      = Error(0)
	Error_GL_INVALID_ENUM                  = Error(1280)
	Error_GL_INVALID_VALUE                 = Error(1281)
	Error_GL_INVALID_OPERATION             = Error(1282)
	Error_GL_INVALID_FRAMEBUFFER_OPERATION = Error(1286)
	Error_GL_OUT_OF_MEMORY                 = Error(1285)
)

////////////////////////////////////////////////////////////////////////////////
// enum HintTarget
////////////////////////////////////////////////////////////////////////////////
type HintTarget uint32

const (
	HintTarget_GL_GENERATE_MIPMAP_HINT = HintTarget(33170)
)

////////////////////////////////////////////////////////////////////////////////
// enum HintMode
////////////////////////////////////////////////////////////////////////////////
type HintMode uint32

const (
	HintMode_GL_DONT_CARE = HintMode(4352)
	HintMode_GL_FASTEST   = HintMode(4353)
	HintMode_GL_NICEST    = HintMode(4354)
)

////////////////////////////////////////////////////////////////////////////////
// enum DiscardFramebufferAttachment
////////////////////////////////////////////////////////////////////////////////
type DiscardFramebufferAttachment uint32

const (
	DiscardFramebufferAttachment_GL_COLOR_EXT   = DiscardFramebufferAttachment(6144)
	DiscardFramebufferAttachment_GL_DEPTH_EXT   = DiscardFramebufferAttachment(6145)
	DiscardFramebufferAttachment_GL_STENCIL_EXT = DiscardFramebufferAttachment(6146)
)

////////////////////////////////////////////////////////////////////////////////
// enum ProgramParameter
////////////////////////////////////////////////////////////////////////////////
type ProgramParameter uint32

const (
	ProgramParameter_GL_DELETE_STATUS               = ProgramParameter(35712)
	ProgramParameter_GL_LINK_STATUS                 = ProgramParameter(35714)
	ProgramParameter_GL_VALIDATE_STATUS             = ProgramParameter(35715)
	ProgramParameter_GL_INFO_LOG_LENGTH             = ProgramParameter(35716)
	ProgramParameter_GL_ATTACHED_SHADERS            = ProgramParameter(35717)
	ProgramParameter_GL_ACTIVE_ATTRIBUTES           = ProgramParameter(35721)
	ProgramParameter_GL_ACTIVE_ATTRIBUTE_MAX_LENGTH = ProgramParameter(35722)
	ProgramParameter_GL_ACTIVE_UNIFORMS             = ProgramParameter(35718)
	ProgramParameter_GL_ACTIVE_UNIFORM_MAX_LENGTH   = ProgramParameter(35719)
)

////////////////////////////////////////////////////////////////////////////////
// enum ShaderParameter
////////////////////////////////////////////////////////////////////////////////
type ShaderParameter uint32

const (
	ShaderParameter_GL_SHADER_TYPE          = ShaderParameter(35663)
	ShaderParameter_GL_DELETE_STATUS        = ShaderParameter(35712)
	ShaderParameter_GL_COMPILE_STATUS       = ShaderParameter(35713)
	ShaderParameter_GL_INFO_LOG_LENGTH      = ShaderParameter(35716)
	ShaderParameter_GL_SHADER_SOURCE_LENGTH = ShaderParameter(35720)
)

////////////////////////////////////////////////////////////////////////////////
// enum PixelStoreParameter
////////////////////////////////////////////////////////////////////////////////
type PixelStoreParameter uint32

const (
	PixelStoreParameter_GL_PACK_ALIGNMENT   = PixelStoreParameter(3333)
	PixelStoreParameter_GL_UNPACK_ALIGNMENT = PixelStoreParameter(3317)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter_FilterMode
////////////////////////////////////////////////////////////////////////////////
type TextureParameter_FilterMode uint32

const (
	TextureParameter_FilterMode_GL_TEXTURE_MIN_FILTER = TextureParameter_FilterMode(10241)
	TextureParameter_FilterMode_GL_TEXTURE_MAG_FILTER = TextureParameter_FilterMode(10240)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter_WrapMode
////////////////////////////////////////////////////////////////////////////////
type TextureParameter_WrapMode uint32

const (
	TextureParameter_WrapMode_GL_TEXTURE_WRAP_S = TextureParameter_WrapMode(10242)
	TextureParameter_WrapMode_GL_TEXTURE_WRAP_T = TextureParameter_WrapMode(10243)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter_EXT_texture_filter_anisotropic
////////////////////////////////////////////////////////////////////////////////
type TextureParameter_EXT_texture_filter_anisotropic uint32

const (
	TextureParameter_EXT_texture_filter_anisotropic_GL_TEXTURE_MAX_ANISOTROPY_EXT = TextureParameter_EXT_texture_filter_anisotropic(34046)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter_SwizzleMode
////////////////////////////////////////////////////////////////////////////////
type TextureParameter_SwizzleMode uint32

const (
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_R = TextureParameter_SwizzleMode(36418)
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_G = TextureParameter_SwizzleMode(36419)
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_B = TextureParameter_SwizzleMode(36420)
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_A = TextureParameter_SwizzleMode(36421)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter
////////////////////////////////////////////////////////////////////////////////
type TextureParameter uint32

const ()

// TextureParameter_FilterMode
const (
	TextureParameter_GL_TEXTURE_MIN_FILTER = TextureParameter(10241)
	TextureParameter_GL_TEXTURE_MAG_FILTER = TextureParameter(10240)
)

// TextureParameter_WrapMode
const (
	TextureParameter_GL_TEXTURE_WRAP_S = TextureParameter(10242)
	TextureParameter_GL_TEXTURE_WRAP_T = TextureParameter(10243)
)

// TextureParameter_SwizzleMode
const (
	TextureParameter_GL_TEXTURE_SWIZZLE_R = TextureParameter(36418)
	TextureParameter_GL_TEXTURE_SWIZZLE_G = TextureParameter(36419)
	TextureParameter_GL_TEXTURE_SWIZZLE_B = TextureParameter(36420)
	TextureParameter_GL_TEXTURE_SWIZZLE_A = TextureParameter(36421)
)

// TextureParameter_EXT_texture_filter_anisotropic
const (
	TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT = TextureParameter(34046)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureFilterMode
////////////////////////////////////////////////////////////////////////////////
type TextureFilterMode uint32

const (
	TextureFilterMode_GL_NEAREST                = TextureFilterMode(9728)
	TextureFilterMode_GL_LINEAR                 = TextureFilterMode(9729)
	TextureFilterMode_GL_NEAREST_MIPMAP_NEAREST = TextureFilterMode(9984)
	TextureFilterMode_GL_LINEAR_MIPMAP_NEAREST  = TextureFilterMode(9985)
	TextureFilterMode_GL_NEAREST_MIPMAP_LINEAR  = TextureFilterMode(9986)
	TextureFilterMode_GL_LINEAR_MIPMAP_LINEAR   = TextureFilterMode(9987)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureWrapMode
////////////////////////////////////////////////////////////////////////////////
type TextureWrapMode uint32

const (
	TextureWrapMode_GL_CLAMP_TO_EDGE   = TextureWrapMode(33071)
	TextureWrapMode_GL_MIRRORED_REPEAT = TextureWrapMode(33648)
	TextureWrapMode_GL_REPEAT          = TextureWrapMode(10497)
)

////////////////////////////////////////////////////////////////////////////////
// enum TexelComponent
////////////////////////////////////////////////////////////////////////////////
type TexelComponent uint32

const (
	TexelComponent_GL_RED   = TexelComponent(6403)
	TexelComponent_GL_GREEN = TexelComponent(6404)
	TexelComponent_GL_BLUE  = TexelComponent(6405)
	TexelComponent_GL_ALPHA = TexelComponent(6406)
)

////////////////////////////////////////////////////////////////////////////////
// enum BlendFactor
////////////////////////////////////////////////////////////////////////////////
type BlendFactor uint32

const (
	BlendFactor_GL_ZERO                     = BlendFactor(0)
	BlendFactor_GL_ONE                      = BlendFactor(1)
	BlendFactor_GL_SRC_COLOR                = BlendFactor(768)
	BlendFactor_GL_ONE_MINUS_SRC_COLOR      = BlendFactor(769)
	BlendFactor_GL_DST_COLOR                = BlendFactor(774)
	BlendFactor_GL_ONE_MINUS_DST_COLOR      = BlendFactor(775)
	BlendFactor_GL_SRC_ALPHA                = BlendFactor(770)
	BlendFactor_GL_ONE_MINUS_SRC_ALPHA      = BlendFactor(771)
	BlendFactor_GL_DST_ALPHA                = BlendFactor(772)
	BlendFactor_GL_ONE_MINUS_DST_ALPHA      = BlendFactor(773)
	BlendFactor_GL_CONSTANT_COLOR           = BlendFactor(32769)
	BlendFactor_GL_ONE_MINUS_CONSTANT_COLOR = BlendFactor(32770)
	BlendFactor_GL_CONSTANT_ALPHA           = BlendFactor(32771)
	BlendFactor_GL_ONE_MINUS_CONSTANT_ALPHA = BlendFactor(32772)
	BlendFactor_GL_SRC_ALPHA_SATURATE       = BlendFactor(776)
)

////////////////////////////////////////////////////////////////////////////////
// enum PrecisionType
////////////////////////////////////////////////////////////////////////////////
type PrecisionType uint32

const (
	PrecisionType_GL_LOW_FLOAT    = PrecisionType(36336)
	PrecisionType_GL_MEDIUM_FLOAT = PrecisionType(36337)
	PrecisionType_GL_HIGH_FLOAT   = PrecisionType(36338)
	PrecisionType_GL_LOW_INT      = PrecisionType(36339)
	PrecisionType_GL_MEDIUM_INT   = PrecisionType(36340)
	PrecisionType_GL_HIGH_INT     = PrecisionType(36341)
)

////////////////////////////////////////////////////////////////////////////////
// enum TestFunction
////////////////////////////////////////////////////////////////////////////////
type TestFunction uint32

const (
	TestFunction_GL_NEVER    = TestFunction(512)
	TestFunction_GL_LESS     = TestFunction(513)
	TestFunction_GL_EQUAL    = TestFunction(514)
	TestFunction_GL_LEQUAL   = TestFunction(515)
	TestFunction_GL_GREATER  = TestFunction(516)
	TestFunction_GL_NOTEQUAL = TestFunction(517)
	TestFunction_GL_GEQUAL   = TestFunction(518)
	TestFunction_GL_ALWAYS   = TestFunction(519)
)

////////////////////////////////////////////////////////////////////////////////
// enum StencilAction
////////////////////////////////////////////////////////////////////////////////
type StencilAction uint32

const (
	StencilAction_GL_KEEP      = StencilAction(7680)
	StencilAction_GL_ZERO      = StencilAction(0)
	StencilAction_GL_REPLACE   = StencilAction(7681)
	StencilAction_GL_INCR      = StencilAction(7682)
	StencilAction_GL_INCR_WRAP = StencilAction(34055)
	StencilAction_GL_DECR      = StencilAction(7683)
	StencilAction_GL_DECR_WRAP = StencilAction(34056)
	StencilAction_GL_INVERT    = StencilAction(5386)
)

////////////////////////////////////////////////////////////////////////////////
// enum FaceOrientation
////////////////////////////////////////////////////////////////////////////////
type FaceOrientation uint32

const (
	FaceOrientation_GL_CW  = FaceOrientation(2304)
	FaceOrientation_GL_CCW = FaceOrientation(2305)
)

////////////////////////////////////////////////////////////////////////////////
// enum BlendEquation
////////////////////////////////////////////////////////////////////////////////
type BlendEquation uint32

const (
	BlendEquation_GL_FUNC_ADD              = BlendEquation(32774)
	BlendEquation_GL_FUNC_SUBTRACT         = BlendEquation(32778)
	BlendEquation_GL_FUNC_REVERSE_SUBTRACT = BlendEquation(32779)
)

////////////////////////////////////////////////////////////////////////////////
// enum BufferTarget
////////////////////////////////////////////////////////////////////////////////
type BufferTarget uint32

const (
	BufferTarget_GL_ARRAY_BUFFER              = BufferTarget(34962)
	BufferTarget_GL_COPY_READ_BUFFER          = BufferTarget(36662)
	BufferTarget_GL_COPY_WRITE_BUFFER         = BufferTarget(36663)
	BufferTarget_GL_ELEMENT_ARRAY_BUFFER      = BufferTarget(34963)
	BufferTarget_GL_PIXEL_PACK_BUFFER         = BufferTarget(35051)
	BufferTarget_GL_PIXEL_UNPACK_BUFFER       = BufferTarget(35052)
	BufferTarget_GL_TRANSFORM_FEEDBACK_BUFFER = BufferTarget(35982)
	BufferTarget_GL_UNIFORM_BUFFER            = BufferTarget(35345)
)

////////////////////////////////////////////////////////////////////////////////
// enum ImageTargetTexture_OES_EGL_image
////////////////////////////////////////////////////////////////////////////////
type ImageTargetTexture_OES_EGL_image uint32

const (
	ImageTargetTexture_OES_EGL_image_GL_TEXTURE_2D = ImageTargetTexture_OES_EGL_image(3553)
)

////////////////////////////////////////////////////////////////////////////////
// enum ImageTargetTexture_OES_EGL_image_external
////////////////////////////////////////////////////////////////////////////////
type ImageTargetTexture_OES_EGL_image_external uint32

const (
	ImageTargetTexture_OES_EGL_image_external_GL_TEXTURE_EXTERNAL_OES = ImageTargetTexture_OES_EGL_image_external(36197)
)

////////////////////////////////////////////////////////////////////////////////
// enum ImageTargetTexture
////////////////////////////////////////////////////////////////////////////////
type ImageTargetTexture uint32

const ()

// ImageTargetTexture_OES_EGL_image
const (
	ImageTargetTexture_GL_TEXTURE_2D = ImageTargetTexture(3553)
)

// ImageTargetTexture_OES_EGL_image_external
const (
	ImageTargetTexture_GL_TEXTURE_EXTERNAL_OES = ImageTargetTexture(36197)
)

////////////////////////////////////////////////////////////////////////////////
// enum ImageTargetRenderbufferStorage
////////////////////////////////////////////////////////////////////////////////
type ImageTargetRenderbufferStorage uint32

const (
	ImageTargetRenderbufferStorage_GL_RENDERBUFFER_OES = ImageTargetRenderbufferStorage(36161)
)

////////////////////////////////////////////////////////////////////////////////
// enum ResetStatus
////////////////////////////////////////////////////////////////////////////////
type ResetStatus uint32

const (
	ResetStatus_GL_NO_ERROR                   = ResetStatus(0)
	ResetStatus_GL_GUILTY_CONTEXT_RESET_EXT   = ResetStatus(33363)
	ResetStatus_GL_INNOCENT_CONTEXT_RESET_EXT = ResetStatus(33364)
	ResetStatus_GL_UNKNOWN_CONTEXT_RESET_EXT  = ResetStatus(33365)
)

////////////////////////////////////////////////////////////////////////////////
// enum TextureKind
////////////////////////////////////////////////////////////////////////////////
type TextureKind uint32

const (
	TextureKind_UNDEFINED = TextureKind(0)
	TextureKind_TEXTURE2D = TextureKind(1)
	TextureKind_CUBEMAP   = TextureKind(2)
)

////////////////////////////////////////////////////////////////////////////////
// enum QueryParameter_GLES_3
////////////////////////////////////////////////////////////////////////////////
type QueryParameter_GLES_3 uint32

const (
	QueryParameter_GLES_3_GL_CURRENT_QUERY = QueryParameter_GLES_3(34917)
)

////////////////////////////////////////////////////////////////////////////////
// enum QueryParameter_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type QueryParameter_EXT_disjoint_timer_query uint32

const (
	QueryParameter_EXT_disjoint_timer_query_GL_QUERY_COUNTER_BITS_EXT = QueryParameter_EXT_disjoint_timer_query(34916)
)

////////////////////////////////////////////////////////////////////////////////
// enum QueryParameter
////////////////////////////////////////////////////////////////////////////////
type QueryParameter uint32

const ()

// QueryParameter_GLES_3
const (
	QueryParameter_GL_CURRENT_QUERY = QueryParameter(34917)
)

// QueryParameter_EXT_disjoint_timer_query
const (
	QueryParameter_GL_QUERY_COUNTER_BITS_EXT = QueryParameter(34916)
)

////////////////////////////////////////////////////////////////////////////////
// enum QueryObjectParameter_GLES_3
////////////////////////////////////////////////////////////////////////////////
type QueryObjectParameter_GLES_3 uint32

const (
	QueryObjectParameter_GLES_3_GL_QUERY_RESULT           = QueryObjectParameter_GLES_3(34918)
	QueryObjectParameter_GLES_3_GL_QUERY_RESULT_AVAILABLE = QueryObjectParameter_GLES_3(34919)
)

////////////////////////////////////////////////////////////////////////////////
// enum QueryObjectParameter_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type QueryObjectParameter_EXT_disjoint_timer_query uint32

const ()

////////////////////////////////////////////////////////////////////////////////
// enum QueryObjectParameter
////////////////////////////////////////////////////////////////////////////////
type QueryObjectParameter uint32

const ()

// QueryObjectParameter_GLES_3
const (
	QueryObjectParameter_GL_QUERY_RESULT           = QueryObjectParameter(34918)
	QueryObjectParameter_GL_QUERY_RESULT_AVAILABLE = QueryObjectParameter(34919)
)

// QueryObjectParameter_EXT_disjoint_timer_query
const ()

////////////////////////////////////////////////////////////////////////////////
// enum QueryTarget_GLES_3
////////////////////////////////////////////////////////////////////////////////
type QueryTarget_GLES_3 uint32

const (
	QueryTarget_GLES_3_GL_ANY_SAMPLES_PASSED                    = QueryTarget_GLES_3(35887)
	QueryTarget_GLES_3_GL_ANY_SAMPLES_PASSED_CONSERVATIVE       = QueryTarget_GLES_3(36202)
	QueryTarget_GLES_3_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = QueryTarget_GLES_3(35976)
)

////////////////////////////////////////////////////////////////////////////////
// enum QueryTarget_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type QueryTarget_EXT_disjoint_timer_query uint32

const (
	QueryTarget_EXT_disjoint_timer_query_GL_TIME_ELAPSED_EXT = QueryTarget_EXT_disjoint_timer_query(35007)
	QueryTarget_EXT_disjoint_timer_query_GL_TIMESTAMP_EXT    = QueryTarget_EXT_disjoint_timer_query(36392)
)

////////////////////////////////////////////////////////////////////////////////
// enum QueryTarget
////////////////////////////////////////////////////////////////////////////////
type QueryTarget uint32

const ()

// QueryTarget_GLES_3
const (
	QueryTarget_GL_ANY_SAMPLES_PASSED                    = QueryTarget(35887)
	QueryTarget_GL_ANY_SAMPLES_PASSED_CONSERVATIVE       = QueryTarget(36202)
	QueryTarget_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = QueryTarget(35976)
)

// QueryTarget_EXT_disjoint_timer_query
const (
	QueryTarget_GL_TIME_ELAPSED_EXT = QueryTarget(35007)
	QueryTarget_GL_TIMESTAMP_EXT    = QueryTarget(36392)
)

////////////////////////////////////////////////////////////////////////////////
// enum TilePreserveMaskQCOM
////////////////////////////////////////////////////////////////////////////////
type TilePreserveMaskQCOM uint32

const (
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT0_QCOM       = TilePreserveMaskQCOM(1)
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT1_QCOM       = TilePreserveMaskQCOM(2)
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT2_QCOM       = TilePreserveMaskQCOM(4)
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT3_QCOM       = TilePreserveMaskQCOM(8)
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT4_QCOM       = TilePreserveMaskQCOM(16)
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT5_QCOM       = TilePreserveMaskQCOM(32)
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT6_QCOM       = TilePreserveMaskQCOM(64)
	TilePreserveMaskQCOM_GL_COLOR_BUFFER_BIT7_QCOM       = TilePreserveMaskQCOM(128)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT0_QCOM       = TilePreserveMaskQCOM(256)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT1_QCOM       = TilePreserveMaskQCOM(512)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT2_QCOM       = TilePreserveMaskQCOM(1024)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT3_QCOM       = TilePreserveMaskQCOM(2048)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT4_QCOM       = TilePreserveMaskQCOM(4096)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT5_QCOM       = TilePreserveMaskQCOM(8192)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT6_QCOM       = TilePreserveMaskQCOM(16384)
	TilePreserveMaskQCOM_GL_DEPTH_BUFFER_BIT7_QCOM       = TilePreserveMaskQCOM(32768)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT0_QCOM     = TilePreserveMaskQCOM(65536)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT1_QCOM     = TilePreserveMaskQCOM(131072)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT2_QCOM     = TilePreserveMaskQCOM(262144)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT3_QCOM     = TilePreserveMaskQCOM(524288)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT4_QCOM     = TilePreserveMaskQCOM(1048576)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT5_QCOM     = TilePreserveMaskQCOM(2097152)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT6_QCOM     = TilePreserveMaskQCOM(4194304)
	TilePreserveMaskQCOM_GL_STENCIL_BUFFER_BIT7_QCOM     = TilePreserveMaskQCOM(8388608)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT0_QCOM = TilePreserveMaskQCOM(16777216)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT1_QCOM = TilePreserveMaskQCOM(33554432)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT2_QCOM = TilePreserveMaskQCOM(67108864)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT3_QCOM = TilePreserveMaskQCOM(134217728)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT4_QCOM = TilePreserveMaskQCOM(268435456)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT5_QCOM = TilePreserveMaskQCOM(536870912)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT6_QCOM = TilePreserveMaskQCOM(1073741824)
	TilePreserveMaskQCOM_GL_MULTISAMPLE_BUFFER_BIT7_QCOM = TilePreserveMaskQCOM(2147483648)
)

////////////////////////////////////////////////////////////////////////////////
// enum ClearMask
////////////////////////////////////////////////////////////////////////////////
type ClearMask uint32

const (
	ClearMask_GL_COLOR_BUFFER_BIT   = ClearMask(16384)
	ClearMask_GL_DEPTH_BUFFER_BIT   = ClearMask(256)
	ClearMask_GL_STENCIL_BUFFER_BIT = ClearMask(1024)
)

////////////////////////////////////////////////////////////////////////////////
// enum MapBufferRangeAccess
////////////////////////////////////////////////////////////////////////////////
type MapBufferRangeAccess uint32

const (
	MapBufferRangeAccess_GL_MAP_READ_BIT              = MapBufferRangeAccess(1)
	MapBufferRangeAccess_GL_MAP_WRITE_BIT             = MapBufferRangeAccess(2)
	MapBufferRangeAccess_GL_MAP_INVALIDATE_RANGE_BIT  = MapBufferRangeAccess(4)
	MapBufferRangeAccess_GL_MAP_INVALIDATE_BUFFER_BIT = MapBufferRangeAccess(8)
	MapBufferRangeAccess_GL_MAP_FLUSH_EXPLICIT_BIT    = MapBufferRangeAccess(16)
	MapBufferRangeAccess_GL_MAP_UNSYNCHRONIZED_BIT    = MapBufferRangeAccess(32)
)

////////////////////////////////////////////////////////////////////////////////
// Globals
////////////////////////////////////////////////////////////////////////////////
type Globals struct {
	binary.Generate
	NextContextID ContextID
	CurrentThread ThreadID
	Contexts      ThreadIDːContextʳᵐ
	EGLContexts   EGLContextːContextʳᵐ
	GLXContexts   GLXContextːContextʳᵐ
	WGLContexts   HGLRCːContextʳᵐ
	CGLContexts   CGLContextObjːContextʳᵐ
}

func (g *Globals) Init() {
	g.Contexts = make(ThreadIDːContextʳᵐ)
	g.EGLContexts = make(EGLContextːContextʳᵐ)
	g.GLXContexts = make(GLXContextːContextʳᵐ)
	g.WGLContexts = make(HGLRCːContextʳᵐ)
	g.CGLContexts = make(CGLContextObjːContextʳᵐ)
}
func NewReplayCreateRenderer(Id uint32) *ReplayCreateRenderer {
	return &ReplayCreateRenderer{Id: Id}
}
func NewReplayBindRenderer(Id uint32) *ReplayBindRenderer {
	return &ReplayBindRenderer{Id: Id}
}
func NewBackbufferInfo(Width int32, Height int32, Color_fmt RenderbufferFormat, Depth_fmt RenderbufferFormat, Stencil_fmt RenderbufferFormat, ResetViewportScissor bool) *BackbufferInfo {
	return &BackbufferInfo{Width: Width, Height: Height, ColorFmt: Color_fmt, DepthFmt: Depth_fmt, StencilFmt: Stencil_fmt, ResetViewportScissor: ResetViewportScissor}
}
func NewStartTimer(Index uint8) *StartTimer {
	return &StartTimer{Index: Index}
}
func NewStopTimer(Index uint8, Result uint64) *StopTimer {
	return &StopTimer{Index: Index, Result: Result}
}
func NewFlushPostBuffer() *FlushPostBuffer {
	return &FlushPostBuffer{}
}
func NewEglInitialize(Dpy memory.Pointer, Major memory.Pointer, Minor memory.Pointer, Result EGLBoolean) *EglInitialize {
	return &EglInitialize{Dpy: NewEGLDisplay(Dpy), Major: NewEGLintᵖ(Major), Minor: NewEGLintᵖ(Minor), Result: Result}
}
func NewEglCreateContext(Display memory.Pointer, Config memory.Pointer, Share_context memory.Pointer, Attrib_list memory.Pointer, Result memory.Pointer) *EglCreateContext {
	return &EglCreateContext{Display: NewEGLDisplay(Display), Config: NewEGLConfig(Config), ShareContext: NewEGLContext(Share_context), AttribList: NewEGLintᵖ(Attrib_list), Result: NewEGLContext(Result)}
}
func NewEglMakeCurrent(Display memory.Pointer, Draw memory.Pointer, Read memory.Pointer, Context memory.Pointer, Result EGLBoolean) *EglMakeCurrent {
	return &EglMakeCurrent{Display: NewEGLDisplay(Display), Draw: NewEGLSurface(Draw), Read: NewEGLSurface(Read), Context: NewEGLContext(Context), Result: Result}
}
func NewEglSwapBuffers(Display memory.Pointer, Surface memory.Pointer, Result EGLBoolean) *EglSwapBuffers {
	return &EglSwapBuffers{Display: NewEGLDisplay(Display), Surface: NewVoidᵖ(Surface), Result: Result}
}
func NewEglQuerySurface(Display memory.Pointer, Surface memory.Pointer, Attribute EGLint, Value memory.Pointer, Result EGLBoolean) *EglQuerySurface {
	return &EglQuerySurface{Display: NewEGLDisplay(Display), Surface: NewEGLSurface(Surface), Attribute: Attribute, Value: NewEGLintᵖ(Value), Result: Result}
}
func NewGlXCreateContext(Dpy memory.Pointer, Vis memory.Pointer, ShareList memory.Pointer, Direct bool, Result memory.Pointer) *GlXCreateContext {
	return &GlXCreateContext{Dpy: NewVoidᵖ(Dpy), Vis: NewVoidᵖ(Vis), ShareList: NewGLXContext(ShareList), Direct: Direct, Result: NewGLXContext(Result)}
}
func NewGlXCreateNewContext(Display memory.Pointer, Fbconfig memory.Pointer, Type uint32, Shared memory.Pointer, Direct bool, Result memory.Pointer) *GlXCreateNewContext {
	return &GlXCreateNewContext{Display: NewVoidᵖ(Display), Fbconfig: NewVoidᵖ(Fbconfig), Type: Type, Shared: NewGLXContext(Shared), Direct: Direct, Result: NewGLXContext(Result)}
}
func NewGlXMakeContextCurrent(Display memory.Pointer, Draw memory.Pointer, Read memory.Pointer, Ctx memory.Pointer) *GlXMakeContextCurrent {
	return &GlXMakeContextCurrent{Display: NewVoidᵖ(Display), Draw: NewGLXDrawable(Draw), Read: NewGLXDrawable(Read), Ctx: NewGLXContext(Ctx)}
}
func NewGlXSwapBuffers(Display memory.Pointer, Drawable memory.Pointer) *GlXSwapBuffers {
	return &GlXSwapBuffers{Display: NewVoidᵖ(Display), Drawable: NewGLXDrawable(Drawable)}
}
func NewWglCreateContext(Hdc memory.Pointer, Result memory.Pointer) *WglCreateContext {
	return &WglCreateContext{Hdc: NewHDC(Hdc), Result: NewHGLRC(Result)}
}
func NewWglCreateContextAttribsARB(Hdc memory.Pointer, HShareContext memory.Pointer, AttribList memory.Pointer, Result memory.Pointer) *WglCreateContextAttribsARB {
	return &WglCreateContextAttribsARB{Hdc: NewHDC(Hdc), HShareContext: NewHGLRC(HShareContext), AttribList: NewIntᵖ(AttribList), Result: NewHGLRC(Result)}
}
func NewWglMakeCurrent(Hdc memory.Pointer, Hglrc memory.Pointer, Result BOOL) *WglMakeCurrent {
	return &WglMakeCurrent{Hdc: NewHDC(Hdc), Hglrc: NewHGLRC(Hglrc), Result: Result}
}
func NewWglSwapBuffers(Hdc memory.Pointer) *WglSwapBuffers {
	return &WglSwapBuffers{Hdc: NewHDC(Hdc)}
}
func NewCGLCreateContext(Pix memory.Pointer, Share memory.Pointer, Ctx memory.Pointer, Result CGLError) *CGLCreateContext {
	return &CGLCreateContext{Pix: NewCGLPixelFormatObj(Pix), Share: NewCGLContextObj(Share), Ctx: NewCGLContextObjᵖ(Ctx), Result: Result}
}
func NewCGLSetCurrentContext(Ctx memory.Pointer, Result CGLError) *CGLSetCurrentContext {
	return &CGLSetCurrentContext{Ctx: NewCGLContextObj(Ctx), Result: Result}
}
func NewGlEnableClientState(Type ArrayType) *GlEnableClientState {
	return &GlEnableClientState{Type: Type}
}
func NewGlDisableClientState(Type ArrayType) *GlDisableClientState {
	return &GlDisableClientState{Type: Type}
}
func NewGlGetProgramBinaryOES(Program ProgramId, Buffer_size int32, Bytes_written memory.Pointer, Binary_format memory.Pointer, Binary memory.Pointer) *GlGetProgramBinaryOES {
	return &GlGetProgramBinaryOES{Program: Program, BufferSize: Buffer_size, BytesWritten: NewS32ᵖ(Bytes_written), BinaryFormat: NewU32ᵖ(Binary_format), Binary: NewVoidᵖ(Binary)}
}
func NewGlProgramBinaryOES(Program ProgramId, Binary_format uint32, Binary memory.Pointer, Binary_size int32) *GlProgramBinaryOES {
	return &GlProgramBinaryOES{Program: Program, BinaryFormat: Binary_format, Binary: NewVoidᵖ(Binary), BinarySize: Binary_size}
}
func NewGlStartTilingQCOM(X int32, Y int32, Width int32, Height int32, PreserveMask TilePreserveMaskQCOM) *GlStartTilingQCOM {
	return &GlStartTilingQCOM{X: X, Y: Y, Width: Width, Height: Height, PreserveMask: PreserveMask}
}
func NewGlEndTilingQCOM(Preserve_mask TilePreserveMaskQCOM) *GlEndTilingQCOM {
	return &GlEndTilingQCOM{PreserveMask: Preserve_mask}
}
func NewGlDiscardFramebufferEXT(Target FramebufferTarget, NumAttachments int32, Attachments memory.Pointer) *GlDiscardFramebufferEXT {
	return &GlDiscardFramebufferEXT{Target: Target, NumAttachments: NumAttachments, Attachments: NewDiscardFramebufferAttachmentᵖ(Attachments)}
}
func NewGlInsertEventMarkerEXT(Length int32, Marker string) *GlInsertEventMarkerEXT {
	return &GlInsertEventMarkerEXT{Length: Length, Marker: Marker}
}
func NewGlPushGroupMarkerEXT(Length int32, Marker string) *GlPushGroupMarkerEXT {
	return &GlPushGroupMarkerEXT{Length: Length, Marker: Marker}
}
func NewGlPopGroupMarkerEXT() *GlPopGroupMarkerEXT {
	return &GlPopGroupMarkerEXT{}
}
func NewGlTexStorage1DEXT(Target TextureTarget, Levels int32, Format TexelFormat, Width int32) *GlTexStorage1DEXT {
	return &GlTexStorage1DEXT{Target: Target, Levels: Levels, Format: Format, Width: Width}
}
func NewGlTexStorage2DEXT(Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32) *GlTexStorage2DEXT {
	return &GlTexStorage2DEXT{Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height}
}
func NewGlTexStorage3DEXT(Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32, Depth int32) *GlTexStorage3DEXT {
	return &GlTexStorage3DEXT{Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height, Depth: Depth}
}
func NewGlTextureStorage1DEXT(Texture TextureId, Target TextureTarget, Levels int32, Format TexelFormat, Width int32) *GlTextureStorage1DEXT {
	return &GlTextureStorage1DEXT{Texture: Texture, Target: Target, Levels: Levels, Format: Format, Width: Width}
}
func NewGlTextureStorage2DEXT(Texture TextureId, Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32) *GlTextureStorage2DEXT {
	return &GlTextureStorage2DEXT{Texture: Texture, Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height}
}
func NewGlTextureStorage3DEXT(Texture TextureId, Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32, Depth int32) *GlTextureStorage3DEXT {
	return &GlTextureStorage3DEXT{Texture: Texture, Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height, Depth: Depth}
}
func NewGlGenVertexArraysOES(Count int32, Arrays memory.Pointer) *GlGenVertexArraysOES {
	return &GlGenVertexArraysOES{Count: Count, Arrays: NewVertexArrayIdᵖ(Arrays)}
}
func NewGlBindVertexArrayOES(Array VertexArrayId) *GlBindVertexArrayOES {
	return &GlBindVertexArrayOES{Array: Array}
}
func NewGlDeleteVertexArraysOES(Count int32, Arrays memory.Pointer) *GlDeleteVertexArraysOES {
	return &GlDeleteVertexArraysOES{Count: Count, Arrays: NewVertexArrayIdᵖ(Arrays)}
}
func NewGlIsVertexArrayOES(Array VertexArrayId, Result bool) *GlIsVertexArrayOES {
	return &GlIsVertexArrayOES{Array: Array, Result: Result}
}
func NewGlEGLImageTargetTexture2DOES(Target ImageTargetTexture, Image memory.Pointer) *GlEGLImageTargetTexture2DOES {
	return &GlEGLImageTargetTexture2DOES{Target: Target, Image: NewImageOES(Image)}
}
func NewGlEGLImageTargetRenderbufferStorageOES(Target ImageTargetRenderbufferStorage, Image memory.Pointer) *GlEGLImageTargetRenderbufferStorageOES {
	return &GlEGLImageTargetRenderbufferStorageOES{Target: Target, Image: NewTexturePointer(Image)}
}
func NewGlGetGraphicsResetStatusEXT(Result ResetStatus) *GlGetGraphicsResetStatusEXT {
	return &GlGetGraphicsResetStatusEXT{Result: Result}
}
func NewGlBindAttribLocation(Program ProgramId, Location AttributeLocation, Name string) *GlBindAttribLocation {
	return &GlBindAttribLocation{Program: Program, Location: Location, Name: Name}
}
func NewGlBlendFunc(Src_factor BlendFactor, Dst_factor BlendFactor) *GlBlendFunc {
	return &GlBlendFunc{SrcFactor: Src_factor, DstFactor: Dst_factor}
}
func NewGlBlendFuncSeparate(Src_factor_rgb BlendFactor, Dst_factor_rgb BlendFactor, Src_factor_alpha BlendFactor, Dst_factor_alpha BlendFactor) *GlBlendFuncSeparate {
	return &GlBlendFuncSeparate{SrcFactorRgb: Src_factor_rgb, DstFactorRgb: Dst_factor_rgb, SrcFactorAlpha: Src_factor_alpha, DstFactorAlpha: Dst_factor_alpha}
}
func NewGlBlendEquation(Equation BlendEquation) *GlBlendEquation {
	return &GlBlendEquation{Equation: Equation}
}
func NewGlBlendEquationSeparate(Rgb BlendEquation, Alpha BlendEquation) *GlBlendEquationSeparate {
	return &GlBlendEquationSeparate{Rgb: Rgb, Alpha: Alpha}
}
func NewGlBlendColor(Red float32, Green float32, Blue float32, Alpha float32) *GlBlendColor {
	return &GlBlendColor{Red: Red, Green: Green, Blue: Blue, Alpha: Alpha}
}
func NewGlEnableVertexAttribArray(Location AttributeLocation) *GlEnableVertexAttribArray {
	return &GlEnableVertexAttribArray{Location: Location}
}
func NewGlDisableVertexAttribArray(Location AttributeLocation) *GlDisableVertexAttribArray {
	return &GlDisableVertexAttribArray{Location: Location}
}
func NewGlVertexAttribPointer(Location AttributeLocation, Size int32, Type VertexAttribType, Normalized bool, Stride int32, Data memory.Pointer) *GlVertexAttribPointer {
	return &GlVertexAttribPointer{Location: Location, Size: Size, Type: Type, Normalized: Normalized, Stride: Stride, Data: NewVertexPointer(Data)}
}
func NewGlGetActiveAttrib(Program ProgramId, Location AttributeLocation, Buffer_size int32, Buffer_bytes_written memory.Pointer, Vector_count memory.Pointer, Type memory.Pointer, Name memory.Pointer) *GlGetActiveAttrib {
	return &GlGetActiveAttrib{Program: Program, Location: Location, BufferSize: Buffer_size, BufferBytesWritten: NewS32ᵖ(Buffer_bytes_written), VectorCount: NewS32ᵖ(Vector_count), Type: NewShaderAttribTypeᵖ(Type), Name: NewCharᵖ(Name)}
}
func NewGlGetActiveUniform(Program ProgramId, Location int32, Buffer_size int32, Buffer_bytes_written memory.Pointer, Size memory.Pointer, Type memory.Pointer, Name memory.Pointer) *GlGetActiveUniform {
	return &GlGetActiveUniform{Program: Program, Location: Location, BufferSize: Buffer_size, BufferBytesWritten: NewS32ᵖ(Buffer_bytes_written), Size: NewS32ᵖ(Size), Type: NewShaderUniformTypeᵖ(Type), Name: NewCharᵖ(Name)}
}
func NewGlGetError(Result Error) *GlGetError {
	return &GlGetError{Result: Result}
}
func NewGlGetProgramiv(Program ProgramId, Parameter ProgramParameter, Value memory.Pointer) *GlGetProgramiv {
	return &GlGetProgramiv{Program: Program, Parameter: Parameter, Value: NewS32ᵖ(Value)}
}
func NewGlGetShaderiv(Shader ShaderId, Parameter ShaderParameter, Value memory.Pointer) *GlGetShaderiv {
	return &GlGetShaderiv{Shader: Shader, Parameter: Parameter, Value: NewS32ᵖ(Value)}
}
func NewGlGetUniformLocation(Program ProgramId, Name string, Result UniformLocation) *GlGetUniformLocation {
	return &GlGetUniformLocation{Program: Program, Name: Name, Result: Result}
}
func NewGlGetAttribLocation(Program ProgramId, Name string, Result AttributeLocation) *GlGetAttribLocation {
	return &GlGetAttribLocation{Program: Program, Name: Name, Result: Result}
}
func NewGlPixelStorei(Parameter PixelStoreParameter, Value int32) *GlPixelStorei {
	return &GlPixelStorei{Parameter: Parameter, Value: Value}
}
func NewGlTexParameteri(Target TextureTarget, Parameter TextureParameter, Value int32) *GlTexParameteri {
	return &GlTexParameteri{Target: Target, Parameter: Parameter, Value: Value}
}
func NewGlTexParameterf(Target TextureTarget, Parameter TextureParameter, Value float32) *GlTexParameterf {
	return &GlTexParameterf{Target: Target, Parameter: Parameter, Value: Value}
}
func NewGlGetTexParameteriv(Target TextureTarget, Parameter TextureParameter, Values memory.Pointer) *GlGetTexParameteriv {
	return &GlGetTexParameteriv{Target: Target, Parameter: Parameter, Values: NewS32ᵖ(Values)}
}
func NewGlGetTexParameterfv(Target TextureTarget, Parameter TextureParameter, Values memory.Pointer) *GlGetTexParameterfv {
	return &GlGetTexParameterfv{Target: Target, Parameter: Parameter, Values: NewF32ᵖ(Values)}
}
func NewGlUniform1i(Location UniformLocation, Value int32) *GlUniform1i {
	return &GlUniform1i{Location: Location, Value: Value}
}
func NewGlUniform2i(Location UniformLocation, Value0 int32, Value1 int32) *GlUniform2i {
	return &GlUniform2i{Location: Location, Value0: Value0, Value1: Value1}
}
func NewGlUniform3i(Location UniformLocation, Value0 int32, Value1 int32, Value2 int32) *GlUniform3i {
	return &GlUniform3i{Location: Location, Value0: Value0, Value1: Value1, Value2: Value2}
}
func NewGlUniform4i(Location UniformLocation, Value0 int32, Value1 int32, Value2 int32, Value3 int32) *GlUniform4i {
	return &GlUniform4i{Location: Location, Value0: Value0, Value1: Value1, Value2: Value2, Value3: Value3}
}
func NewGlUniform1iv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform1iv {
	return &GlUniform1iv{Location: Location, Count: Count, Value: NewS32ᵖ(Value)}
}
func NewGlUniform2iv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform2iv {
	return &GlUniform2iv{Location: Location, Count: Count, Value: NewS32ᵖ(Value)}
}
func NewGlUniform3iv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform3iv {
	return &GlUniform3iv{Location: Location, Count: Count, Value: NewS32ᵖ(Value)}
}
func NewGlUniform4iv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform4iv {
	return &GlUniform4iv{Location: Location, Count: Count, Value: NewS32ᵖ(Value)}
}
func NewGlUniform1f(Location UniformLocation, Value float32) *GlUniform1f {
	return &GlUniform1f{Location: Location, Value: Value}
}
func NewGlUniform2f(Location UniformLocation, Value0 float32, Value1 float32) *GlUniform2f {
	return &GlUniform2f{Location: Location, Value0: Value0, Value1: Value1}
}
func NewGlUniform3f(Location UniformLocation, Value0 float32, Value1 float32, Value2 float32) *GlUniform3f {
	return &GlUniform3f{Location: Location, Value0: Value0, Value1: Value1, Value2: Value2}
}
func NewGlUniform4f(Location UniformLocation, Value0 float32, Value1 float32, Value2 float32, Value3 float32) *GlUniform4f {
	return &GlUniform4f{Location: Location, Value0: Value0, Value1: Value1, Value2: Value2, Value3: Value3}
}
func NewGlUniform1fv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform1fv {
	return &GlUniform1fv{Location: Location, Count: Count, Value: NewF32ᵖ(Value)}
}
func NewGlUniform2fv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform2fv {
	return &GlUniform2fv{Location: Location, Count: Count, Value: NewF32ᵖ(Value)}
}
func NewGlUniform3fv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform3fv {
	return &GlUniform3fv{Location: Location, Count: Count, Value: NewF32ᵖ(Value)}
}
func NewGlUniform4fv(Location UniformLocation, Count int32, Value memory.Pointer) *GlUniform4fv {
	return &GlUniform4fv{Location: Location, Count: Count, Value: NewF32ᵖ(Value)}
}
func NewGlUniformMatrix2fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix2fv {
	return &GlUniformMatrix2fv{Location: Location, Count: Count, Transpose: Transpose, Values: NewF32ᵖ(Values)}
}
func NewGlUniformMatrix3fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix3fv {
	return &GlUniformMatrix3fv{Location: Location, Count: Count, Transpose: Transpose, Values: NewF32ᵖ(Values)}
}
func NewGlUniformMatrix4fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix4fv {
	return &GlUniformMatrix4fv{Location: Location, Count: Count, Transpose: Transpose, Values: NewF32ᵖ(Values)}
}
func NewGlGetUniformfv(Program ProgramId, Location UniformLocation, Values memory.Pointer) *GlGetUniformfv {
	return &GlGetUniformfv{Program: Program, Location: Location, Values: NewF32ᵖ(Values)}
}
func NewGlGetUniformiv(Program ProgramId, Location UniformLocation, Values memory.Pointer) *GlGetUniformiv {
	return &GlGetUniformiv{Program: Program, Location: Location, Values: NewS32ᵖ(Values)}
}
func NewGlVertexAttrib1f(Location AttributeLocation, Value0 float32) *GlVertexAttrib1f {
	return &GlVertexAttrib1f{Location: Location, Value0: Value0}
}
func NewGlVertexAttrib2f(Location AttributeLocation, Value0 float32, Value1 float32) *GlVertexAttrib2f {
	return &GlVertexAttrib2f{Location: Location, Value0: Value0, Value1: Value1}
}
func NewGlVertexAttrib3f(Location AttributeLocation, Value0 float32, Value1 float32, Value2 float32) *GlVertexAttrib3f {
	return &GlVertexAttrib3f{Location: Location, Value0: Value0, Value1: Value1, Value2: Value2}
}
func NewGlVertexAttrib4f(Location AttributeLocation, Value0 float32, Value1 float32, Value2 float32, Value3 float32) *GlVertexAttrib4f {
	return &GlVertexAttrib4f{Location: Location, Value0: Value0, Value1: Value1, Value2: Value2, Value3: Value3}
}
func NewGlVertexAttrib1fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib1fv {
	return &GlVertexAttrib1fv{Location: Location, Value: NewF32ᵖ(Value)}
}
func NewGlVertexAttrib2fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib2fv {
	return &GlVertexAttrib2fv{Location: Location, Value: NewF32ᵖ(Value)}
}
func NewGlVertexAttrib3fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib3fv {
	return &GlVertexAttrib3fv{Location: Location, Value: NewF32ᵖ(Value)}
}
func NewGlVertexAttrib4fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib4fv {
	return &GlVertexAttrib4fv{Location: Location, Value: NewF32ᵖ(Value)}
}
func NewGlGetShaderPrecisionFormat(Shader_type ShaderType, Precision_type PrecisionType, Range memory.Pointer, Precision memory.Pointer) *GlGetShaderPrecisionFormat {
	return &GlGetShaderPrecisionFormat{ShaderType: Shader_type, PrecisionType: Precision_type, Range: NewS32ᵖ(Range), Precision: NewS32ᵖ(Precision)}
}
func NewGlDepthMask(Enabled bool) *GlDepthMask {
	return &GlDepthMask{Enabled: Enabled}
}
func NewGlDepthFunc(Function TestFunction) *GlDepthFunc {
	return &GlDepthFunc{Function: Function}
}
func NewGlDepthRangef(Near float32, Far float32) *GlDepthRangef {
	return &GlDepthRangef{Near: Near, Far: Far}
}
func NewGlColorMask(Red bool, Green bool, Blue bool, Alpha bool) *GlColorMask {
	return &GlColorMask{Red: Red, Green: Green, Blue: Blue, Alpha: Alpha}
}
func NewGlStencilMask(Mask uint32) *GlStencilMask {
	return &GlStencilMask{Mask: Mask}
}
func NewGlStencilMaskSeparate(Face FaceMode, Mask uint32) *GlStencilMaskSeparate {
	return &GlStencilMaskSeparate{Face: Face, Mask: Mask}
}
func NewGlStencilFuncSeparate(Face FaceMode, Function TestFunction, Reference_value int32, Mask int32) *GlStencilFuncSeparate {
	return &GlStencilFuncSeparate{Face: Face, Function: Function, ReferenceValue: Reference_value, Mask: Mask}
}
func NewGlStencilOpSeparate(Face FaceMode, Stencil_fail StencilAction, Stencil_pass_depth_fail StencilAction, Stencil_pass_depth_pass StencilAction) *GlStencilOpSeparate {
	return &GlStencilOpSeparate{Face: Face, StencilFail: Stencil_fail, StencilPassDepthFail: Stencil_pass_depth_fail, StencilPassDepthPass: Stencil_pass_depth_pass}
}
func NewGlFrontFace(Orientation FaceOrientation) *GlFrontFace {
	return &GlFrontFace{Orientation: Orientation}
}
func NewGlViewport(X int32, Y int32, Width int32, Height int32) *GlViewport {
	return &GlViewport{X: X, Y: Y, Width: Width, Height: Height}
}
func NewGlScissor(X int32, Y int32, Width int32, Height int32) *GlScissor {
	return &GlScissor{X: X, Y: Y, Width: Width, Height: Height}
}
func NewGlActiveTexture(Unit TextureUnit) *GlActiveTexture {
	return &GlActiveTexture{Unit: Unit}
}
func NewGlGenTextures(Count int32, Textures memory.Pointer) *GlGenTextures {
	return &GlGenTextures{Count: Count, Textures: NewTextureIdᵖ(Textures)}
}
func NewGlDeleteTextures(Count int32, Textures memory.Pointer) *GlDeleteTextures {
	return &GlDeleteTextures{Count: Count, Textures: NewTextureIdᵖ(Textures)}
}
func NewGlIsTexture(Texture TextureId, Result bool) *GlIsTexture {
	return &GlIsTexture{Texture: Texture, Result: Result}
}
func NewGlBindTexture(Target TextureTarget, Texture TextureId) *GlBindTexture {
	return &GlBindTexture{Target: Target, Texture: Texture}
}
func NewGlTexImage2D(Target TextureImageTarget, Level int32, Internal_format TexelFormat, Width int32, Height int32, Border int32, Format TexelFormat, Type TexelType, Data memory.Pointer) *GlTexImage2D {
	return &GlTexImage2D{Target: Target, Level: Level, InternalFormat: Internal_format, Width: Width, Height: Height, Border: Border, Format: Format, Type: Type, Data: NewTexturePointer(Data)}
}
func NewGlTexSubImage2D(Target TextureImageTarget, Level int32, Xoffset int32, Yoffset int32, Width int32, Height int32, Format TexelFormat, Type TexelType, Data memory.Pointer) *GlTexSubImage2D {
	return &GlTexSubImage2D{Target: Target, Level: Level, Xoffset: Xoffset, Yoffset: Yoffset, Width: Width, Height: Height, Format: Format, Type: Type, Data: NewTexturePointer(Data)}
}
func NewGlCopyTexImage2D(Target TextureImageTarget, Level int32, Format TexelFormat, X int32, Y int32, Width int32, Height int32, Border int32) *GlCopyTexImage2D {
	return &GlCopyTexImage2D{Target: Target, Level: Level, Format: Format, X: X, Y: Y, Width: Width, Height: Height, Border: Border}
}
func NewGlCopyTexSubImage2D(Target TextureImageTarget, Level int32, Xoffset int32, Yoffset int32, X int32, Y int32, Width int32, Height int32) *GlCopyTexSubImage2D {
	return &GlCopyTexSubImage2D{Target: Target, Level: Level, Xoffset: Xoffset, Yoffset: Yoffset, X: X, Y: Y, Width: Width, Height: Height}
}
func NewGlCompressedTexImage2D(Target TextureImageTarget, Level int32, Format CompressedTexelFormat, Width int32, Height int32, Border int32, Image_size int32, Data memory.Pointer) *GlCompressedTexImage2D {
	return &GlCompressedTexImage2D{Target: Target, Level: Level, Format: Format, Width: Width, Height: Height, Border: Border, ImageSize: Image_size, Data: NewTexturePointer(Data)}
}
func NewGlCompressedTexSubImage2D(Target TextureImageTarget, Level int32, Xoffset int32, Yoffset int32, Width int32, Height int32, Format CompressedTexelFormat, Image_size int32, Data memory.Pointer) *GlCompressedTexSubImage2D {
	return &GlCompressedTexSubImage2D{Target: Target, Level: Level, Xoffset: Xoffset, Yoffset: Yoffset, Width: Width, Height: Height, Format: Format, ImageSize: Image_size, Data: NewTexturePointer(Data)}
}
func NewGlGenerateMipmap(Target TextureImageTarget) *GlGenerateMipmap {
	return &GlGenerateMipmap{Target: Target}
}
func NewGlReadPixels(X int32, Y int32, Width int32, Height int32, Format BaseTexelFormat, Type TexelType, Data memory.Pointer) *GlReadPixels {
	return &GlReadPixels{X: X, Y: Y, Width: Width, Height: Height, Format: Format, Type: Type, Data: NewVoidᵖ(Data)}
}
func NewGlGenFramebuffers(Count int32, Framebuffers memory.Pointer) *GlGenFramebuffers {
	return &GlGenFramebuffers{Count: Count, Framebuffers: NewFramebufferIdᵖ(Framebuffers)}
}
func NewGlBindFramebuffer(Target FramebufferTarget, Framebuffer FramebufferId) *GlBindFramebuffer {
	return &GlBindFramebuffer{Target: Target, Framebuffer: Framebuffer}
}
func NewGlCheckFramebufferStatus(Target FramebufferTarget, Result FramebufferStatus) *GlCheckFramebufferStatus {
	return &GlCheckFramebufferStatus{Target: Target, Result: Result}
}
func NewGlDeleteFramebuffers(Count int32, Framebuffers memory.Pointer) *GlDeleteFramebuffers {
	return &GlDeleteFramebuffers{Count: Count, Framebuffers: NewFramebufferIdᵖ(Framebuffers)}
}
func NewGlIsFramebuffer(Framebuffer FramebufferId, Result bool) *GlIsFramebuffer {
	return &GlIsFramebuffer{Framebuffer: Framebuffer, Result: Result}
}
func NewGlGenRenderbuffers(Count int32, Renderbuffers memory.Pointer) *GlGenRenderbuffers {
	return &GlGenRenderbuffers{Count: Count, Renderbuffers: NewRenderbufferIdᵖ(Renderbuffers)}
}
func NewGlBindRenderbuffer(Target RenderbufferTarget, Renderbuffer RenderbufferId) *GlBindRenderbuffer {
	return &GlBindRenderbuffer{Target: Target, Renderbuffer: Renderbuffer}
}
func NewGlRenderbufferStorage(Target RenderbufferTarget, Format RenderbufferFormat, Width int32, Height int32) *GlRenderbufferStorage {
	return &GlRenderbufferStorage{Target: Target, Format: Format, Width: Width, Height: Height}
}
func NewGlDeleteRenderbuffers(Count int32, Renderbuffers memory.Pointer) *GlDeleteRenderbuffers {
	return &GlDeleteRenderbuffers{Count: Count, Renderbuffers: NewRenderbufferIdᵖ(Renderbuffers)}
}
func NewGlIsRenderbuffer(Renderbuffer RenderbufferId, Result bool) *GlIsRenderbuffer {
	return &GlIsRenderbuffer{Renderbuffer: Renderbuffer, Result: Result}
}
func NewGlGetRenderbufferParameteriv(Target RenderbufferTarget, Parameter RenderbufferParameter, Values memory.Pointer) *GlGetRenderbufferParameteriv {
	return &GlGetRenderbufferParameteriv{Target: Target, Parameter: Parameter, Values: NewS32ᵖ(Values)}
}
func NewGlGenBuffers(Count int32, Buffers memory.Pointer) *GlGenBuffers {
	return &GlGenBuffers{Count: Count, Buffers: NewBufferIdᵖ(Buffers)}
}
func NewGlBindBuffer(Target BufferTarget, Buffer BufferId) *GlBindBuffer {
	return &GlBindBuffer{Target: Target, Buffer: Buffer}
}
func NewGlBufferData(Target BufferTarget, Size int32, Data memory.Pointer, Usage BufferUsage) *GlBufferData {
	return &GlBufferData{Target: Target, Size: Size, Data: NewBufferDataPointer(Data), Usage: Usage}
}
func NewGlBufferSubData(Target BufferTarget, Offset int32, Size int32, Data memory.Pointer) *GlBufferSubData {
	return &GlBufferSubData{Target: Target, Offset: Offset, Size: Size, Data: NewVoidᵖ(Data)}
}
func NewGlDeleteBuffers(Count int32, Buffers memory.Pointer) *GlDeleteBuffers {
	return &GlDeleteBuffers{Count: Count, Buffers: NewBufferIdᵖ(Buffers)}
}
func NewGlIsBuffer(Buffer BufferId, Result bool) *GlIsBuffer {
	return &GlIsBuffer{Buffer: Buffer, Result: Result}
}
func NewGlGetBufferParameteriv(Target BufferTarget, Parameter BufferParameter, Value memory.Pointer) *GlGetBufferParameteriv {
	return &GlGetBufferParameteriv{Target: Target, Parameter: Parameter, Value: NewS32ᵖ(Value)}
}
func NewGlCreateShader(Type ShaderType, Result ShaderId) *GlCreateShader {
	return &GlCreateShader{Type: Type, Result: Result}
}
func NewGlDeleteShader(Shader ShaderId) *GlDeleteShader {
	return &GlDeleteShader{Shader: Shader}
}
func NewGlShaderSource(Shader ShaderId, Count int32, Source memory.Pointer, Length memory.Pointer) *GlShaderSource {
	return &GlShaderSource{Shader: Shader, Count: Count, Source: NewCharᵖᵖ(Source), Length: NewS32ᵖ(Length)}
}
func NewGlShaderBinary(Count int32, Shaders memory.Pointer, Binary_format uint32, Binary memory.Pointer, Binary_size int32) *GlShaderBinary {
	return &GlShaderBinary{Count: Count, Shaders: NewShaderIdᵖ(Shaders), BinaryFormat: Binary_format, Binary: NewVoidᵖ(Binary), BinarySize: Binary_size}
}
func NewGlGetShaderInfoLog(Shader ShaderId, Buffer_length int32, String_length_written memory.Pointer, Info memory.Pointer) *GlGetShaderInfoLog {
	return &GlGetShaderInfoLog{Shader: Shader, BufferLength: Buffer_length, StringLengthWritten: NewS32ᵖ(String_length_written), Info: NewCharᵖ(Info)}
}
func NewGlGetShaderSource(Shader ShaderId, Buffer_length int32, String_length_written memory.Pointer, Source memory.Pointer) *GlGetShaderSource {
	return &GlGetShaderSource{Shader: Shader, BufferLength: Buffer_length, StringLengthWritten: NewS32ᵖ(String_length_written), Source: NewCharᵖ(Source)}
}
func NewGlReleaseShaderCompiler() *GlReleaseShaderCompiler {
	return &GlReleaseShaderCompiler{}
}
func NewGlCompileShader(Shader ShaderId) *GlCompileShader {
	return &GlCompileShader{Shader: Shader}
}
func NewGlIsShader(Shader ShaderId, Result bool) *GlIsShader {
	return &GlIsShader{Shader: Shader, Result: Result}
}
func NewGlCreateProgram(Result ProgramId) *GlCreateProgram {
	return &GlCreateProgram{Result: Result}
}
func NewGlDeleteProgram(Program ProgramId) *GlDeleteProgram {
	return &GlDeleteProgram{Program: Program}
}
func NewGlAttachShader(Program ProgramId, Shader ShaderId) *GlAttachShader {
	return &GlAttachShader{Program: Program, Shader: Shader}
}
func NewGlDetachShader(Program ProgramId, Shader ShaderId) *GlDetachShader {
	return &GlDetachShader{Program: Program, Shader: Shader}
}
func NewGlGetAttachedShaders(Program ProgramId, Buffer_length int32, Shaders_length_written memory.Pointer, Shaders memory.Pointer) *GlGetAttachedShaders {
	return &GlGetAttachedShaders{Program: Program, BufferLength: Buffer_length, ShadersLengthWritten: NewS32ᵖ(Shaders_length_written), Shaders: NewShaderIdᵖ(Shaders)}
}
func NewGlLinkProgram(Program ProgramId) *GlLinkProgram {
	return &GlLinkProgram{Program: Program}
}
func NewGlGetProgramInfoLog(Program ProgramId, Buffer_length int32, String_length_written memory.Pointer, Info memory.Pointer) *GlGetProgramInfoLog {
	return &GlGetProgramInfoLog{Program: Program, BufferLength: Buffer_length, StringLengthWritten: NewS32ᵖ(String_length_written), Info: NewCharᵖ(Info)}
}
func NewGlUseProgram(Program ProgramId) *GlUseProgram {
	return &GlUseProgram{Program: Program}
}
func NewGlIsProgram(Program ProgramId, Result bool) *GlIsProgram {
	return &GlIsProgram{Program: Program, Result: Result}
}
func NewGlValidateProgram(Program ProgramId) *GlValidateProgram {
	return &GlValidateProgram{Program: Program}
}
func NewGlClearColor(R float32, G float32, B float32, A float32) *GlClearColor {
	return &GlClearColor{R: R, G: G, B: B, A: A}
}
func NewGlClearDepthf(Depth float32) *GlClearDepthf {
	return &GlClearDepthf{Depth: Depth}
}
func NewGlClearStencil(Stencil int32) *GlClearStencil {
	return &GlClearStencil{Stencil: Stencil}
}
func NewGlClear(Mask ClearMask) *GlClear {
	return &GlClear{Mask: Mask}
}
func NewGlCullFace(Mode FaceMode) *GlCullFace {
	return &GlCullFace{Mode: Mode}
}
func NewGlPolygonOffset(Scale_factor float32, Units float32) *GlPolygonOffset {
	return &GlPolygonOffset{ScaleFactor: Scale_factor, Units: Units}
}
func NewGlLineWidth(Width float32) *GlLineWidth {
	return &GlLineWidth{Width: Width}
}
func NewGlSampleCoverage(Value float32, Invert bool) *GlSampleCoverage {
	return &GlSampleCoverage{Value: Value, Invert: Invert}
}
func NewGlHint(Target HintTarget, Mode HintMode) *GlHint {
	return &GlHint{Target: Target, Mode: Mode}
}
func NewGlFramebufferRenderbuffer(Framebuffer_target FramebufferTarget, Framebuffer_attachment FramebufferAttachment, Renderbuffer_target RenderbufferTarget, Renderbuffer RenderbufferId) *GlFramebufferRenderbuffer {
	return &GlFramebufferRenderbuffer{FramebufferTarget: Framebuffer_target, FramebufferAttachment: Framebuffer_attachment, RenderbufferTarget: Renderbuffer_target, Renderbuffer: Renderbuffer}
}
func NewGlFramebufferTexture2D(Framebuffer_target FramebufferTarget, Framebuffer_attachment FramebufferAttachment, Texture_target TextureImageTarget, Texture TextureId, Level int32) *GlFramebufferTexture2D {
	return &GlFramebufferTexture2D{FramebufferTarget: Framebuffer_target, FramebufferAttachment: Framebuffer_attachment, TextureTarget: Texture_target, Texture: Texture, Level: Level}
}
func NewGlGetFramebufferAttachmentParameteriv(Framebuffer_target FramebufferTarget, Attachment FramebufferAttachment, Parameter FramebufferAttachmentParameter, Value memory.Pointer) *GlGetFramebufferAttachmentParameteriv {
	return &GlGetFramebufferAttachmentParameteriv{FramebufferTarget: Framebuffer_target, Attachment: Attachment, Parameter: Parameter, Value: NewS32ᵖ(Value)}
}
func NewGlDrawElements(Draw_mode DrawMode, Element_count int32, Indices_type IndicesType, Indices memory.Pointer) *GlDrawElements {
	return &GlDrawElements{DrawMode: Draw_mode, ElementCount: Element_count, IndicesType: Indices_type, Indices: NewIndicesPointer(Indices)}
}
func NewGlDrawArrays(Draw_mode DrawMode, First_index int32, Index_count int32) *GlDrawArrays {
	return &GlDrawArrays{DrawMode: Draw_mode, FirstIndex: First_index, IndexCount: Index_count}
}
func NewGlFlush() *GlFlush {
	return &GlFlush{}
}
func NewGlFinish() *GlFinish {
	return &GlFinish{}
}
func NewGlGetBooleanv(Param StateVariable, Values memory.Pointer) *GlGetBooleanv {
	return &GlGetBooleanv{Param: Param, Values: NewBoolᵖ(Values)}
}
func NewGlGetFloatv(Param StateVariable, Values memory.Pointer) *GlGetFloatv {
	return &GlGetFloatv{Param: Param, Values: NewF32ᵖ(Values)}
}
func NewGlGetIntegerv(Param StateVariable, Values memory.Pointer) *GlGetIntegerv {
	return &GlGetIntegerv{Param: Param, Values: NewS32ᵖ(Values)}
}
func NewGlGetString(Param StringConstant, Result memory.Pointer) *GlGetString {
	return &GlGetString{Param: Param, Result: NewCharᵖ(Result)}
}
func NewGlEnable(Capability Capability) *GlEnable {
	return &GlEnable{Capability: Capability}
}
func NewGlDisable(Capability Capability) *GlDisable {
	return &GlDisable{Capability: Capability}
}
func NewGlIsEnabled(Capability Capability, Result bool) *GlIsEnabled {
	return &GlIsEnabled{Capability: Capability, Result: Result}
}
func NewGlMapBufferRange(Target BufferTarget, Offset int32, Length int32, Access MapBufferRangeAccess, Result memory.Pointer) *GlMapBufferRange {
	return &GlMapBufferRange{Target: Target, Offset: Offset, Length: Length, Access: Access, Result: NewVoidᵖ(Result)}
}
func NewGlUnmapBuffer(Target BufferTarget) *GlUnmapBuffer {
	return &GlUnmapBuffer{Target: Target}
}
func NewGlInvalidateFramebuffer(Target FramebufferTarget, Count int32, Attachments memory.Pointer) *GlInvalidateFramebuffer {
	return &GlInvalidateFramebuffer{Target: Target, Count: Count, Attachments: NewFramebufferAttachmentᵖ(Attachments)}
}
func NewGlRenderbufferStorageMultisample(Target RenderbufferTarget, Samples int32, Format RenderbufferFormat, Width int32, Height int32) *GlRenderbufferStorageMultisample {
	return &GlRenderbufferStorageMultisample{Target: Target, Samples: Samples, Format: Format, Width: Width, Height: Height}
}
func NewGlBlitFramebuffer(SrcX0 int32, SrcY0 int32, SrcX1 int32, SrcY1 int32, DstX0 int32, DstY0 int32, DstX1 int32, DstY1 int32, Mask ClearMask, Filter TextureFilterMode) *GlBlitFramebuffer {
	return &GlBlitFramebuffer{SrcX0: SrcX0, SrcY0: SrcY0, SrcX1: SrcX1, SrcY1: SrcY1, DstX0: DstX0, DstY0: DstY0, DstX1: DstX1, DstY1: DstY1, Mask: Mask, Filter: Filter}
}
func NewGlGenQueries(Count int32, Queries memory.Pointer) *GlGenQueries {
	return &GlGenQueries{Count: Count, Queries: NewQueryIdᵖ(Queries)}
}
func NewGlBeginQuery(Target QueryTarget, Query QueryId) *GlBeginQuery {
	return &GlBeginQuery{Target: Target, Query: Query}
}
func NewGlEndQuery(Target QueryTarget) *GlEndQuery {
	return &GlEndQuery{Target: Target}
}
func NewGlDeleteQueries(Count int32, Queries memory.Pointer) *GlDeleteQueries {
	return &GlDeleteQueries{Count: Count, Queries: NewQueryIdᵖ(Queries)}
}
func NewGlIsQuery(Query QueryId, Result bool) *GlIsQuery {
	return &GlIsQuery{Query: Query, Result: Result}
}
func NewGlGetQueryiv(Target QueryTarget, Parameter QueryParameter, Value memory.Pointer) *GlGetQueryiv {
	return &GlGetQueryiv{Target: Target, Parameter: Parameter, Value: NewS32ᵖ(Value)}
}
func NewGlGetQueryObjectuiv(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectuiv {
	return &GlGetQueryObjectuiv{Query: Query, Parameter: Parameter, Value: NewU32ᵖ(Value)}
}
func NewGlGenQueriesEXT(Count int32, Queries memory.Pointer) *GlGenQueriesEXT {
	return &GlGenQueriesEXT{Count: Count, Queries: NewQueryIdᵖ(Queries)}
}
func NewGlBeginQueryEXT(Target QueryTarget, Query QueryId) *GlBeginQueryEXT {
	return &GlBeginQueryEXT{Target: Target, Query: Query}
}
func NewGlEndQueryEXT(Target QueryTarget) *GlEndQueryEXT {
	return &GlEndQueryEXT{Target: Target}
}
func NewGlDeleteQueriesEXT(Count int32, Queries memory.Pointer) *GlDeleteQueriesEXT {
	return &GlDeleteQueriesEXT{Count: Count, Queries: NewQueryIdᵖ(Queries)}
}
func NewGlIsQueryEXT(Query QueryId, Result bool) *GlIsQueryEXT {
	return &GlIsQueryEXT{Query: Query, Result: Result}
}
func NewGlQueryCounterEXT(Query QueryId, Target QueryTarget) *GlQueryCounterEXT {
	return &GlQueryCounterEXT{Query: Query, Target: Target}
}
func NewGlGetQueryivEXT(Target QueryTarget, Parameter QueryParameter, Value memory.Pointer) *GlGetQueryivEXT {
	return &GlGetQueryivEXT{Target: Target, Parameter: Parameter, Value: NewS32ᵖ(Value)}
}
func NewGlGetQueryObjectivEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectivEXT {
	return &GlGetQueryObjectivEXT{Query: Query, Parameter: Parameter, Value: NewS32ᵖ(Value)}
}
func NewGlGetQueryObjectuivEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectuivEXT {
	return &GlGetQueryObjectuivEXT{Query: Query, Parameter: Parameter, Value: NewU32ᵖ(Value)}
}
func NewGlGetQueryObjecti64vEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjecti64vEXT {
	return &GlGetQueryObjecti64vEXT{Query: Query, Parameter: Parameter, Value: NewS64ᵖ(Value)}
}
func NewGlGetQueryObjectui64vEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectui64vEXT {
	return &GlGetQueryObjectui64vEXT{Query: Query, Parameter: Parameter, Value: NewU64ᵖ(Value)}
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////
var apiID = gfxapi.ID(binary.NewID([]byte("gles")))

type api struct{}

func (api) Name() string {
	return "gles"
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
		Name: "replayCreateRenderer",
		Docs: "[]",
		ID:   0,
		New:  func() atom.Atom { return &ReplayCreateRenderer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "replayBindRenderer",
		Docs: "[]",
		ID:   1,
		New:  func() atom.Atom { return &ReplayBindRenderer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "backbufferInfo",
		Docs: "[]",
		ID:   2,
		New:  func() atom.Atom { return &BackbufferInfo{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "startTimer",
		Docs: "[]",
		ID:   3,
		New:  func() atom.Atom { return &StartTimer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "stopTimer",
		Docs: "[]",
		ID:   4,
		New:  func() atom.Atom { return &StopTimer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "flushPostBuffer",
		Docs: "[]",
		ID:   5,
		New:  func() atom.Atom { return &FlushPostBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "eglInitialize",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglInitialize.xhtml]",
		ID:   6,
		New:  func() atom.Atom { return &EglInitialize{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "eglCreateContext",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateContext.xhtml]",
		ID:   7,
		New:  func() atom.Atom { return &EglCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "eglMakeCurrent",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglMakeCurrent.xhtml]",
		ID:   8,
		New:  func() atom.Atom { return &EglMakeCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "eglSwapBuffers",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapBuffers.xhtml]",
		ID:   9,
		New:  func() atom.Atom { return &EglSwapBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "eglQuerySurface",
		Docs: "[]",
		ID:   10,
		New:  func() atom.Atom { return &EglQuerySurface{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glXCreateContext",
		Docs: "[]",
		ID:   11,
		New:  func() atom.Atom { return &GlXCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glXCreateNewContext",
		Docs: "[]",
		ID:   12,
		New:  func() atom.Atom { return &GlXCreateNewContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glXMakeContextCurrent",
		Docs: "[]",
		ID:   13,
		New:  func() atom.Atom { return &GlXMakeContextCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glXSwapBuffers",
		Docs: "[]",
		ID:   14,
		New:  func() atom.Atom { return &GlXSwapBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "wglCreateContext",
		Docs: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374379(v=vs.85).aspx]",
		ID:   15,
		New:  func() atom.Atom { return &WglCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "wglCreateContextAttribsARB",
		Docs: "[http://www.opengl.org/registry/specs/ARB/wgl_create_context.txt]",
		ID:   16,
		New:  func() atom.Atom { return &WglCreateContextAttribsARB{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "wglMakeCurrent",
		Docs: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374387(v=vs.85).aspx]",
		ID:   17,
		New:  func() atom.Atom { return &WglMakeCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "wglSwapBuffers",
		Docs: "[http://msdn.microsoft.com/en-us/library/dd369060(v=vs.85)]",
		ID:   18,
		New:  func() atom.Atom { return &WglSwapBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CGLCreateContext",
		Docs: "[http://developer.apple.com/library/mac/documentation/GraphicsImaging/Reference/CGL_OpenGL/index.html#//apple_ref/c/func/CGLCreateContext]",
		ID:   19,
		New:  func() atom.Atom { return &CGLCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "CGLSetCurrentContext",
		Docs: "[]",
		ID:   20,
		New:  func() atom.Atom { return &CGLSetCurrentContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEnableClientState",
		Docs: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
		ID:   21,
		New:  func() atom.Atom { return &GlEnableClientState{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDisableClientState",
		Docs: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
		ID:   22,
		New:  func() atom.Atom { return &GlDisableClientState{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetProgramBinaryOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
		ID:   23,
		New:  func() atom.Atom { return &GlGetProgramBinaryOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glProgramBinaryOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
		ID:   24,
		New:  func() atom.Atom { return &GlProgramBinaryOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glStartTilingQCOM",
		Docs: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
		ID:   25,
		New:  func() atom.Atom { return &GlStartTilingQCOM{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEndTilingQCOM",
		Docs: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
		ID:   26,
		New:  func() atom.Atom { return &GlEndTilingQCOM{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDiscardFramebufferEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_discard_framebuffer.txt]",
		ID:   27,
		New:  func() atom.Atom { return &GlDiscardFramebufferEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glInsertEventMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   28,
		New:  func() atom.Atom { return &GlInsertEventMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glPushGroupMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   29,
		New:  func() atom.Atom { return &GlPushGroupMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glPopGroupMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   30,
		New:  func() atom.Atom { return &GlPopGroupMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTexStorage1DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   31,
		New:  func() atom.Atom { return &GlTexStorage1DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTexStorage2DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   32,
		New:  func() atom.Atom { return &GlTexStorage2DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTexStorage3DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   33,
		New:  func() atom.Atom { return &GlTexStorage3DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTextureStorage1DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   34,
		New:  func() atom.Atom { return &GlTextureStorage1DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTextureStorage2DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   35,
		New:  func() atom.Atom { return &GlTextureStorage2DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTextureStorage3DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   36,
		New:  func() atom.Atom { return &GlTextureStorage3DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenVertexArraysOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   37,
		New:  func() atom.Atom { return &GlGenVertexArraysOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBindVertexArrayOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   38,
		New:  func() atom.Atom { return &GlBindVertexArrayOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteVertexArraysOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   39,
		New:  func() atom.Atom { return &GlDeleteVertexArraysOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsVertexArrayOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   40,
		New:  func() atom.Atom { return &GlIsVertexArrayOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEGLImageTargetTexture2DOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
		ID:   41,
		New:  func() atom.Atom { return &GlEGLImageTargetTexture2DOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEGLImageTargetRenderbufferStorageOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
		ID:   42,
		New:  func() atom.Atom { return &GlEGLImageTargetRenderbufferStorageOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetGraphicsResetStatusEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_robustness.txt]",
		ID:   43,
		New:  func() atom.Atom { return &GlGetGraphicsResetStatusEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBindAttribLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindAttribLocation.xml]",
		ID:   44,
		New:  func() atom.Atom { return &GlBindAttribLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBlendFunc",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFunc.xml]",
		ID:   45,
		New:  func() atom.Atom { return &GlBlendFunc{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBlendFuncSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFuncSeparate.xml]",
		ID:   46,
		New:  func() atom.Atom { return &GlBlendFuncSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBlendEquation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquation.xml]",
		ID:   47,
		New:  func() atom.Atom { return &GlBlendEquation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBlendEquationSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquationSeparate.xml]",
		ID:   48,
		New:  func() atom.Atom { return &GlBlendEquationSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBlendColor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendColor.xml]",
		ID:   49,
		New:  func() atom.Atom { return &GlBlendColor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEnableVertexAttribArray",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml]",
		ID:   50,
		New:  func() atom.Atom { return &GlEnableVertexAttribArray{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDisableVertexAttribArray",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml]",
		ID:   51,
		New:  func() atom.Atom { return &GlDisableVertexAttribArray{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttribPointer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttribPointer.xml]",
		ID:   52,
		New:  func() atom.Atom { return &GlVertexAttribPointer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetActiveAttrib",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveAttrib.xml]",
		ID:   53,
		New:  func() atom.Atom { return &GlGetActiveAttrib{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetActiveUniform",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveUniform.xml]",
		ID:   54,
		New:  func() atom.Atom { return &GlGetActiveUniform{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetError",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetError.xml]",
		ID:   55,
		New:  func() atom.Atom { return &GlGetError{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetProgramiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgram.xml]",
		ID:   56,
		New:  func() atom.Atom { return &GlGetProgramiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetShaderiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderiv.xml]",
		ID:   57,
		New:  func() atom.Atom { return &GlGetShaderiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetUniformLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniformLocation.xml]",
		ID:   58,
		New:  func() atom.Atom { return &GlGetUniformLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetAttribLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttribLocation.xml]",
		ID:   59,
		New:  func() atom.Atom { return &GlGetAttribLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glPixelStorei",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPixelStorei.xml]",
		ID:   60,
		New:  func() atom.Atom { return &GlPixelStorei{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTexParameteri",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
		ID:   61,
		New:  func() atom.Atom { return &GlTexParameteri{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTexParameterf",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
		ID:   62,
		New:  func() atom.Atom { return &GlTexParameterf{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetTexParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
		ID:   63,
		New:  func() atom.Atom { return &GlGetTexParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetTexParameterfv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
		ID:   64,
		New:  func() atom.Atom { return &GlGetTexParameterfv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform1i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   65,
		New:  func() atom.Atom { return &GlUniform1i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform2i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   66,
		New:  func() atom.Atom { return &GlUniform2i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform3i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   67,
		New:  func() atom.Atom { return &GlUniform3i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform4i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   68,
		New:  func() atom.Atom { return &GlUniform4i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform1iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   69,
		New:  func() atom.Atom { return &GlUniform1iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform2iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   70,
		New:  func() atom.Atom { return &GlUniform2iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform3iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   71,
		New:  func() atom.Atom { return &GlUniform3iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform4iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   72,
		New:  func() atom.Atom { return &GlUniform4iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform1f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   73,
		New:  func() atom.Atom { return &GlUniform1f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform2f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   74,
		New:  func() atom.Atom { return &GlUniform2f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform3f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   75,
		New:  func() atom.Atom { return &GlUniform3f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform4f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   76,
		New:  func() atom.Atom { return &GlUniform4f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform1fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   77,
		New:  func() atom.Atom { return &GlUniform1fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   78,
		New:  func() atom.Atom { return &GlUniform2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   79,
		New:  func() atom.Atom { return &GlUniform3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniform4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   80,
		New:  func() atom.Atom { return &GlUniform4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniformMatrix2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   81,
		New:  func() atom.Atom { return &GlUniformMatrix2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniformMatrix3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   82,
		New:  func() atom.Atom { return &GlUniformMatrix3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUniformMatrix4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   83,
		New:  func() atom.Atom { return &GlUniformMatrix4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetUniformfv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
		ID:   84,
		New:  func() atom.Atom { return &GlGetUniformfv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetUniformiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
		ID:   85,
		New:  func() atom.Atom { return &GlGetUniformiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib1f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   86,
		New:  func() atom.Atom { return &GlVertexAttrib1f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib2f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   87,
		New:  func() atom.Atom { return &GlVertexAttrib2f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib3f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   88,
		New:  func() atom.Atom { return &GlVertexAttrib3f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib4f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   89,
		New:  func() atom.Atom { return &GlVertexAttrib4f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib1fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   90,
		New:  func() atom.Atom { return &GlVertexAttrib1fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   91,
		New:  func() atom.Atom { return &GlVertexAttrib2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   92,
		New:  func() atom.Atom { return &GlVertexAttrib3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glVertexAttrib4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   93,
		New:  func() atom.Atom { return &GlVertexAttrib4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetShaderPrecisionFormat",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml]",
		ID:   94,
		New:  func() atom.Atom { return &GlGetShaderPrecisionFormat{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDepthMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthMask.xml]",
		ID:   95,
		New:  func() atom.Atom { return &GlDepthMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDepthFunc",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthFunc.xml]",
		ID:   96,
		New:  func() atom.Atom { return &GlDepthFunc{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDepthRangef",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthRangef.xml]",
		ID:   97,
		New:  func() atom.Atom { return &GlDepthRangef{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glColorMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glColorMask.xml]",
		ID:   98,
		New:  func() atom.Atom { return &GlColorMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glStencilMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMask.xml]",
		ID:   99,
		New:  func() atom.Atom { return &GlStencilMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glStencilMaskSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMaskSeparate.xml]",
		ID:   100,
		New:  func() atom.Atom { return &GlStencilMaskSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glStencilFuncSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilFuncSeparate.xml]",
		ID:   101,
		New:  func() atom.Atom { return &GlStencilFuncSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glStencilOpSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilOpSeparate.xml]",
		ID:   102,
		New:  func() atom.Atom { return &GlStencilOpSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glFrontFace",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFrontFace.xml]",
		ID:   103,
		New:  func() atom.Atom { return &GlFrontFace{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glViewport",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glViewport.xml]",
		ID:   104,
		New:  func() atom.Atom { return &GlViewport{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glScissor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glScissor.xml]",
		ID:   105,
		New:  func() atom.Atom { return &GlScissor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glActiveTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glActiveTexture.xml]",
		ID:   106,
		New:  func() atom.Atom { return &GlActiveTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenTextures",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenTextures.xml]",
		ID:   107,
		New:  func() atom.Atom { return &GlGenTextures{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteTextures",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteTextures.xml]",
		ID:   108,
		New:  func() atom.Atom { return &GlDeleteTextures{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsTexture.xml]",
		ID:   109,
		New:  func() atom.Atom { return &GlIsTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBindTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindTexture.xml]",
		ID:   110,
		New:  func() atom.Atom { return &GlBindTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexImage2D.xml]",
		ID:   111,
		New:  func() atom.Atom { return &GlTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexSubImage2D.xml]",
		ID:   112,
		New:  func() atom.Atom { return &GlTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCopyTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexImage2D.xml]",
		ID:   113,
		New:  func() atom.Atom { return &GlCopyTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCopyTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml]",
		ID:   114,
		New:  func() atom.Atom { return &GlCopyTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCompressedTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexImage2D.xml]",
		ID:   115,
		New:  func() atom.Atom { return &GlCompressedTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCompressedTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml]",
		ID:   116,
		New:  func() atom.Atom { return &GlCompressedTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenerateMipmap",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenerateMipmap.xml]",
		ID:   117,
		New:  func() atom.Atom { return &GlGenerateMipmap{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glReadPixels",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReadPixels.xml]",
		ID:   118,
		New:  func() atom.Atom { return &GlReadPixels{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenFramebuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenFramebuffers.xml]",
		ID:   119,
		New:  func() atom.Atom { return &GlGenFramebuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBindFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindFramebuffer.xml]",
		ID:   120,
		New:  func() atom.Atom { return &GlBindFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCheckFramebufferStatus",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml]",
		ID:   121,
		New:  func() atom.Atom { return &GlCheckFramebufferStatus{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteFramebuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteFramebuffers.xml]",
		ID:   122,
		New:  func() atom.Atom { return &GlDeleteFramebuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsFramebuffer.xml]",
		ID:   123,
		New:  func() atom.Atom { return &GlIsFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenRenderbuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenRenderbuffers.xml]",
		ID:   124,
		New:  func() atom.Atom { return &GlGenRenderbuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBindRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindRenderbuffer.xml]",
		ID:   125,
		New:  func() atom.Atom { return &GlBindRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glRenderbufferStorage",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glRenderbufferStorage.xml]",
		ID:   126,
		New:  func() atom.Atom { return &GlRenderbufferStorage{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteRenderbuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml]",
		ID:   127,
		New:  func() atom.Atom { return &GlDeleteRenderbuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsRenderbuffer.xml]",
		ID:   128,
		New:  func() atom.Atom { return &GlIsRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetRenderbufferParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml]",
		ID:   129,
		New:  func() atom.Atom { return &GlGetRenderbufferParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenBuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenBuffers.xml]",
		ID:   130,
		New:  func() atom.Atom { return &GlGenBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBindBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindBuffer.xml]",
		ID:   131,
		New:  func() atom.Atom { return &GlBindBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBufferData",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferData.xml]",
		ID:   132,
		New:  func() atom.Atom { return &GlBufferData{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBufferSubData",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferSubData.xml]",
		ID:   133,
		New:  func() atom.Atom { return &GlBufferSubData{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteBuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteBuffers.xml]",
		ID:   134,
		New:  func() atom.Atom { return &GlDeleteBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsBuffer.xml]",
		ID:   135,
		New:  func() atom.Atom { return &GlIsBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetBufferParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetBufferParameteriv.xml]",
		ID:   136,
		New:  func() atom.Atom { return &GlGetBufferParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCreateShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateShader.xml]",
		ID:   137,
		New:  func() atom.Atom { return &GlCreateShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteShader.xml]",
		ID:   138,
		New:  func() atom.Atom { return &GlDeleteShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glShaderSource",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderSource.xml]",
		ID:   139,
		New:  func() atom.Atom { return &GlShaderSource{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glShaderBinary",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderBinary.xml]",
		ID:   140,
		New:  func() atom.Atom { return &GlShaderBinary{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetShaderInfoLog",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderInfoLog.xml]",
		ID:   141,
		New:  func() atom.Atom { return &GlGetShaderInfoLog{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetShaderSource",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderSource.xml]",
		ID:   142,
		New:  func() atom.Atom { return &GlGetShaderSource{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glReleaseShaderCompiler",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml]",
		ID:   143,
		New:  func() atom.Atom { return &GlReleaseShaderCompiler{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCompileShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompileShader.xml]",
		ID:   144,
		New:  func() atom.Atom { return &GlCompileShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsShader.xml]",
		ID:   145,
		New:  func() atom.Atom { return &GlIsShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCreateProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateProgram.xml]",
		ID:   146,
		New:  func() atom.Atom { return &GlCreateProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteProgram.xml]",
		ID:   147,
		New:  func() atom.Atom { return &GlDeleteProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glAttachShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glAttachShader.xml]",
		ID:   148,
		New:  func() atom.Atom { return &GlAttachShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDetachShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDetachShader.xml]",
		ID:   149,
		New:  func() atom.Atom { return &GlDetachShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetAttachedShaders",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttachedShaders.xml]",
		ID:   150,
		New:  func() atom.Atom { return &GlGetAttachedShaders{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glLinkProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLinkProgram.xml]",
		ID:   151,
		New:  func() atom.Atom { return &GlLinkProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetProgramInfoLog",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgramInfoLog.xml]",
		ID:   152,
		New:  func() atom.Atom { return &GlGetProgramInfoLog{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUseProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUseProgram.xml]",
		ID:   153,
		New:  func() atom.Atom { return &GlUseProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsProgram.xml]",
		ID:   154,
		New:  func() atom.Atom { return &GlIsProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glValidateProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glValidateProgram.xml]",
		ID:   155,
		New:  func() atom.Atom { return &GlValidateProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glClearColor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearColor.xml]",
		ID:   156,
		New:  func() atom.Atom { return &GlClearColor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glClearDepthf",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearDepthf.xml]",
		ID:   157,
		New:  func() atom.Atom { return &GlClearDepthf{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glClearStencil",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearStencil.xml]",
		ID:   158,
		New:  func() atom.Atom { return &GlClearStencil{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glClear",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClear.xml]",
		ID:   159,
		New:  func() atom.Atom { return &GlClear{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glCullFace",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCullFace.xml]",
		ID:   160,
		New:  func() atom.Atom { return &GlCullFace{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glPolygonOffset",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPolygonOffset.xml]",
		ID:   161,
		New:  func() atom.Atom { return &GlPolygonOffset{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glLineWidth",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLineWidth.xml]",
		ID:   162,
		New:  func() atom.Atom { return &GlLineWidth{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glSampleCoverage",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glSampleCoverage.xml]",
		ID:   163,
		New:  func() atom.Atom { return &GlSampleCoverage{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glHint",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glHint.xml]",
		ID:   164,
		New:  func() atom.Atom { return &GlHint{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glFramebufferRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml]",
		ID:   165,
		New:  func() atom.Atom { return &GlFramebufferRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glFramebufferTexture2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferTexture2D.xml]",
		ID:   166,
		New:  func() atom.Atom { return &GlFramebufferTexture2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetFramebufferAttachmentParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml]",
		ID:   167,
		New:  func() atom.Atom { return &GlGetFramebufferAttachmentParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDrawElements",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawElements.xml]",
		ID:   168,
		New:  func() atom.Atom { return &GlDrawElements{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDrawArrays",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawArrays.xml]",
		ID:   169,
		New:  func() atom.Atom { return &GlDrawArrays{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glFlush",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFlush.xml]",
		ID:   170,
		New:  func() atom.Atom { return &GlFlush{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glFinish",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFinish.xml]",
		ID:   171,
		New:  func() atom.Atom { return &GlFinish{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetBooleanv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   172,
		New:  func() atom.Atom { return &GlGetBooleanv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetFloatv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   173,
		New:  func() atom.Atom { return &GlGetFloatv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetIntegerv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   174,
		New:  func() atom.Atom { return &GlGetIntegerv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetString",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetString.xml]",
		ID:   175,
		New:  func() atom.Atom { return &GlGetString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEnable",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnable.xml]",
		ID:   176,
		New:  func() atom.Atom { return &GlEnable{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDisable",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisable.xml]",
		ID:   177,
		New:  func() atom.Atom { return &GlDisable{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsEnabled",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsEnabled.xml]",
		ID:   178,
		New:  func() atom.Atom { return &GlIsEnabled{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glMapBufferRange",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
		ID:   179,
		New:  func() atom.Atom { return &GlMapBufferRange{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glUnmapBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
		ID:   180,
		New:  func() atom.Atom { return &GlUnmapBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glInvalidateFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glInvalidateFramebuffer.xhtml]",
		ID:   181,
		New:  func() atom.Atom { return &GlInvalidateFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glRenderbufferStorageMultisample",
		Docs: "[http://www.opengl.org/registry/specs/EXT/framebuffer_multisample.txt]",
		ID:   182,
		New:  func() atom.Atom { return &GlRenderbufferStorageMultisample{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBlitFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBlitFramebuffer.xhtml]",
		ID:   183,
		New:  func() atom.Atom { return &GlBlitFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenQueries",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenQueries.xhtml]",
		ID:   184,
		New:  func() atom.Atom { return &GlGenQueries{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBeginQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBeginQuery.xhtml]",
		ID:   185,
		New:  func() atom.Atom { return &GlBeginQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEndQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glEndQuery.xhtml]",
		ID:   186,
		New:  func() atom.Atom { return &GlEndQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteQueries",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteQueries.xhtml]",
		ID:   187,
		New:  func() atom.Atom { return &GlDeleteQueries{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glIsQuery.xhtml]",
		ID:   188,
		New:  func() atom.Atom { return &GlIsQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetQueryiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryiv.xhtml]",
		ID:   189,
		New:  func() atom.Atom { return &GlGetQueryiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetQueryObjectuiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryObjectuiv.xhtml]",
		ID:   190,
		New:  func() atom.Atom { return &GlGetQueryObjectuiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGenQueriesEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   191,
		New:  func() atom.Atom { return &GlGenQueriesEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glBeginQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   192,
		New:  func() atom.Atom { return &GlBeginQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glEndQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   193,
		New:  func() atom.Atom { return &GlEndQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glDeleteQueriesEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   194,
		New:  func() atom.Atom { return &GlDeleteQueriesEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glIsQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   195,
		New:  func() atom.Atom { return &GlIsQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glQueryCounterEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   196,
		New:  func() atom.Atom { return &GlQueryCounterEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetQueryivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   197,
		New:  func() atom.Atom { return &GlGetQueryivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetQueryObjectivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   198,
		New:  func() atom.Atom { return &GlGetQueryObjectivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetQueryObjectuivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   199,
		New:  func() atom.Atom { return &GlGetQueryObjectuivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetQueryObjecti64vEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   200,
		New:  func() atom.Atom { return &GlGetQueryObjecti64vEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "glGetQueryObjectui64vEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   201,
		New:  func() atom.Atom { return &GlGetQueryObjectui64vEXT{} },
	})
}
func min(a, b uint64) uint64 {
	if a < b {
		return a
	} else {
		return b
	}
}
