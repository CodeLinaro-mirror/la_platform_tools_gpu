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

type Vec2i S32ː2ᵃ
type Vec3i S32ː3ᵃ
type Vec4i S32ː4ᵃ
type Vec2f F32ː2ᵃ
type Vec3f F32ː3ᵃ
type Vec4f F32ː4ᵃ
type Mat2f Vec2fː2ᵃ
type Mat3f Vec3fː3ᵃ
type Mat4f Vec4fː4ᵃ
type RenderbufferId uint32
type TextureId uint32
type FramebufferId uint32
type BufferId uint32
type ShaderId uint32
type ProgramId uint32
type VertexArrayId uint32
type QueryId uint32
type UniformLocation int32
type AttributeLocation int32

// IndicesPointer is a pointer to a void element.
type IndicesPointer struct {
	binary.Generate
	memory.Pointer
}

// NewIndicesPointer returns a IndicesPointer that points to addr in the application pool.
func NewIndicesPointer(addr uint64) IndicesPointer {
	return IndicesPointer{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that IndicesPointer points to.
func (p IndicesPointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p IndicesPointer) OnRead(ϟs *gfxapi.State) IndicesPointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p IndicesPointer) OnWrite(ϟs *gfxapi.State) IndicesPointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p IndicesPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// VertexPointer is a pointer to a void element.
type VertexPointer struct {
	binary.Generate
	memory.Pointer
}

// NewVertexPointer returns a VertexPointer that points to addr in the application pool.
func NewVertexPointer(addr uint64) VertexPointer {
	return VertexPointer{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that VertexPointer points to.
func (p VertexPointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p VertexPointer) OnRead(ϟs *gfxapi.State) VertexPointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p VertexPointer) OnWrite(ϟs *gfxapi.State) VertexPointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p VertexPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// TexturePointer is a pointer to a void element.
type TexturePointer struct {
	binary.Generate
	memory.Pointer
}

// NewTexturePointer returns a TexturePointer that points to addr in the application pool.
func NewTexturePointer(addr uint64) TexturePointer {
	return TexturePointer{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that TexturePointer points to.
func (p TexturePointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p TexturePointer) OnRead(ϟs *gfxapi.State) TexturePointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p TexturePointer) OnWrite(ϟs *gfxapi.State) TexturePointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p TexturePointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// BufferDataPointer is a pointer to a void element.
type BufferDataPointer struct {
	binary.Generate
	memory.Pointer
}

// NewBufferDataPointer returns a BufferDataPointer that points to addr in the application pool.
func NewBufferDataPointer(addr uint64) BufferDataPointer {
	return BufferDataPointer{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that BufferDataPointer points to.
func (p BufferDataPointer) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p BufferDataPointer) OnRead(ϟs *gfxapi.State) BufferDataPointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p BufferDataPointer) OnWrite(ϟs *gfxapi.State) BufferDataPointer {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p BufferDataPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

type ContextID uint32
type ThreadID uint64
type EGLBoolean int64
type EGLint int64

// EGLConfig is a pointer to a void element.
type EGLConfig struct {
	binary.Generate
	memory.Pointer
}

// NewEGLConfig returns a EGLConfig that points to addr in the application pool.
func NewEGLConfig(addr uint64) EGLConfig {
	return EGLConfig{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that EGLConfig points to.
func (p EGLConfig) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p EGLConfig) OnRead(ϟs *gfxapi.State) EGLConfig {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p EGLConfig) OnWrite(ϟs *gfxapi.State) EGLConfig {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLConfig) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// EGLContext is a pointer to a void element.
type EGLContext struct {
	binary.Generate
	memory.Pointer
}

// NewEGLContext returns a EGLContext that points to addr in the application pool.
func NewEGLContext(addr uint64) EGLContext {
	return EGLContext{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that EGLContext points to.
func (p EGLContext) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p EGLContext) OnRead(ϟs *gfxapi.State) EGLContext {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p EGLContext) OnWrite(ϟs *gfxapi.State) EGLContext {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLContext) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// EGLDisplay is a pointer to a void element.
type EGLDisplay struct {
	binary.Generate
	memory.Pointer
}

// NewEGLDisplay returns a EGLDisplay that points to addr in the application pool.
func NewEGLDisplay(addr uint64) EGLDisplay {
	return EGLDisplay{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that EGLDisplay points to.
func (p EGLDisplay) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p EGLDisplay) OnRead(ϟs *gfxapi.State) EGLDisplay {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p EGLDisplay) OnWrite(ϟs *gfxapi.State) EGLDisplay {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLDisplay) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// EGLSurface is a pointer to a void element.
type EGLSurface struct {
	binary.Generate
	memory.Pointer
}

// NewEGLSurface returns a EGLSurface that points to addr in the application pool.
func NewEGLSurface(addr uint64) EGLSurface {
	return EGLSurface{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that EGLSurface points to.
func (p EGLSurface) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p EGLSurface) OnRead(ϟs *gfxapi.State) EGLSurface {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p EGLSurface) OnWrite(ϟs *gfxapi.State) EGLSurface {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p EGLSurface) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// GLXContext is a pointer to a void element.
type GLXContext struct {
	binary.Generate
	memory.Pointer
}

// NewGLXContext returns a GLXContext that points to addr in the application pool.
func NewGLXContext(addr uint64) GLXContext {
	return GLXContext{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that GLXContext points to.
func (p GLXContext) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p GLXContext) OnRead(ϟs *gfxapi.State) GLXContext {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p GLXContext) OnWrite(ϟs *gfxapi.State) GLXContext {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p GLXContext) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// GLXDrawable is a pointer to a void element.
type GLXDrawable struct {
	binary.Generate
	memory.Pointer
}

// NewGLXDrawable returns a GLXDrawable that points to addr in the application pool.
func NewGLXDrawable(addr uint64) GLXDrawable {
	return GLXDrawable{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that GLXDrawable points to.
func (p GLXDrawable) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p GLXDrawable) OnRead(ϟs *gfxapi.State) GLXDrawable {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p GLXDrawable) OnWrite(ϟs *gfxapi.State) GLXDrawable {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p GLXDrawable) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

type Bool int64

// HGLRC is a pointer to a void element.
type HGLRC struct {
	binary.Generate
	memory.Pointer
}

// NewHGLRC returns a HGLRC that points to addr in the application pool.
func NewHGLRC(addr uint64) HGLRC {
	return HGLRC{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that HGLRC points to.
func (p HGLRC) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p HGLRC) OnRead(ϟs *gfxapi.State) HGLRC {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p HGLRC) OnWrite(ϟs *gfxapi.State) HGLRC {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p HGLRC) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// HDC is a pointer to a void element.
type HDC struct {
	binary.Generate
	memory.Pointer
}

// NewHDC returns a HDC that points to addr in the application pool.
func NewHDC(addr uint64) HDC {
	return HDC{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that HDC points to.
func (p HDC) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p HDC) OnRead(ϟs *gfxapi.State) HDC {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p HDC) OnWrite(ϟs *gfxapi.State) HDC {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p HDC) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

type BOOL int64
type CGLError int64

// CGLPixelFormatObj is a pointer to a void element.
type CGLPixelFormatObj struct {
	binary.Generate
	memory.Pointer
}

// NewCGLPixelFormatObj returns a CGLPixelFormatObj that points to addr in the application pool.
func NewCGLPixelFormatObj(addr uint64) CGLPixelFormatObj {
	return CGLPixelFormatObj{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that CGLPixelFormatObj points to.
func (p CGLPixelFormatObj) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p CGLPixelFormatObj) OnRead(ϟs *gfxapi.State) CGLPixelFormatObj {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p CGLPixelFormatObj) OnWrite(ϟs *gfxapi.State) CGLPixelFormatObj {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p CGLPixelFormatObj) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// CGLContextObj is a pointer to a void element.
type CGLContextObj struct {
	binary.Generate
	memory.Pointer
}

// NewCGLContextObj returns a CGLContextObj that points to addr in the application pool.
func NewCGLContextObj(addr uint64) CGLContextObj {
	return CGLContextObj{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that CGLContextObj points to.
func (p CGLContextObj) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p CGLContextObj) OnRead(ϟs *gfxapi.State) CGLContextObj {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p CGLContextObj) OnWrite(ϟs *gfxapi.State) CGLContextObj {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p CGLContextObj) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// CGSConnectionID is a pointer to a void element.
type CGSConnectionID struct {
	binary.Generate
	memory.Pointer
}

// NewCGSConnectionID returns a CGSConnectionID that points to addr in the application pool.
func NewCGSConnectionID(addr uint64) CGSConnectionID {
	return CGSConnectionID{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that CGSConnectionID points to.
func (p CGSConnectionID) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p CGSConnectionID) OnRead(ϟs *gfxapi.State) CGSConnectionID {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p CGSConnectionID) OnWrite(ϟs *gfxapi.State) CGSConnectionID {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p CGSConnectionID) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

type CGSWindowID int32
type CGSSurfaceID int32

// ImageOES is a pointer to a void element.
type ImageOES struct {
	binary.Generate
	memory.Pointer
}

// NewImageOES returns a ImageOES that points to addr in the application pool.
func NewImageOES(addr uint64) ImageOES {
	return ImageOES{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that ImageOES points to.
func (p ImageOES) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p ImageOES) OnRead(ϟs *gfxapi.State) ImageOES {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p ImageOES) OnWrite(ϟs *gfxapi.State) ImageOES {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p ImageOES) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

type SyncObject uint64
type GLboolean uint8
type GLbyte int8
type GLubyte uint8
type GLchar byte
type GLshort int16
type GLushort uint16
type GLint int32
type GLuint uint32
type GLint64 int64
type GLuint64 uint64
type GLfixed int32
type GLsizei uint32

// GLsync is a pointer to a __GLsync element.
type GLsync struct {
	binary.Generate
	memory.Pointer
}

// NewGLsync returns a GLsync that points to addr in the application pool.
func NewGLsync(addr uint64) GLsync {
	return GLsync{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that GLsync points to.
func (p GLsync) ElementSize(ϟs *gfxapi.State) uint64 {
	return func() uint64 { panic("Sizeof class is not yet implemented") }()
}

// Read reads and returns the __GLsync element at the pointer.
func (p GLsync) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) __GLsync {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the __GLsync element at the pointer.
func (p GLsync) Write(value __GLsync, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]__GLsync{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p GLsync) OnRead(ϟs *gfxapi.State) GLsync {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p GLsync) OnWrite(ϟs *gfxapi.State) GLsync {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new __GLsyncˢ from the pointer using start and end indices.
func (p GLsync) Slice(start, end uint64, ϟs *gfxapi.State) __GLsyncˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return __GLsyncˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

type GLhalf uint16
type GLfloat float32
type GLclampf float32

// Voidᶜᵖ is a pointer to a void element.
type Voidᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVoidᶜᵖ returns a Voidᶜᵖ that points to addr in the application pool.
func NewVoidᶜᵖ(addr uint64) Voidᶜᵖ {
	return Voidᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Voidᶜᵖ points to.
func (p Voidᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Voidᶜᵖ) OnRead(ϟs *gfxapi.State) Voidᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Voidᶜᵖ) OnWrite(ϟs *gfxapi.State) Voidᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p Voidᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p Voidᵖ) OnRead(ϟs *gfxapi.State) Voidᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Voidᵖ) OnWrite(ϟs *gfxapi.State) Voidᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Voidˢ from the pointer using start and end indices.
func (p Voidᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Voidˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// __GLsyncᵖ is a pointer to a __GLsync element.
type __GLsyncᵖ struct {
	binary.Generate
	memory.Pointer
}

// New__GLsyncᵖ returns a __GLsyncᵖ that points to addr in the application pool.
func New__GLsyncᵖ(addr uint64) __GLsyncᵖ {
	return __GLsyncᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that __GLsyncᵖ points to.
func (p __GLsyncᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return func() uint64 { panic("Sizeof class is not yet implemented") }()
}

// Read reads and returns the __GLsync element at the pointer.
func (p __GLsyncᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) __GLsync {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the __GLsync element at the pointer.
func (p __GLsyncᵖ) Write(value __GLsync, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]__GLsync{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p __GLsyncᵖ) OnRead(ϟs *gfxapi.State) __GLsyncᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p __GLsyncᵖ) OnWrite(ϟs *gfxapi.State) __GLsyncᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new __GLsyncˢ from the pointer using start and end indices.
func (p __GLsyncᵖ) Slice(start, end uint64, ϟs *gfxapi.State) __GLsyncˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return __GLsyncˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

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
func (p U8ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint8 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the uint8 element at the pointer.
func (p U8ᵖ) Write(value uint8, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]uint8{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p U8ᵖ) OnRead(ϟs *gfxapi.State) U8ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p U8ᵖ) OnWrite(ϟs *gfxapi.State) U8ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new U8ˢ from the pointer using start and end indices.
func (p U8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U8ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p Charᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) byte {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the byte element at the pointer.
func (p Charᵖ) Write(value byte, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]byte{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Charᵖ) OnRead(ϟs *gfxapi.State) Charᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Charᵖ) OnWrite(ϟs *gfxapi.State) Charᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// StringSlice returns a slice starting at p and ending at the first 0 byte null-terminator.
// If incNullTerm is true then the null-terminator is included in the slice.
func (p Charᵖ) StringSlice(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, incNullTerm bool) Charˢ {
	i, d := uint64(0), ϟs.MemoryDecoder(ϟs.Memory[p.Pointer.Pool].At(p.Address), ϟd, ϟl)
	for {
		if b, _ := d.Uint8(); b == 0 {
			if incNullTerm {
				return p.Slice(0, i+1, ϟs)
			} else {
				return p.Slice(0, i, ϟs)
			}
		}
		i++
	}
}

// Slice returns a new Charˢ from the pointer using start and end indices.
func (p Charᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// EGLintᵖ is a pointer to a EGLint element.
type EGLintᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewEGLintᵖ returns a EGLintᵖ that points to addr in the application pool.
func NewEGLintᵖ(addr uint64) EGLintᵖ {
	return EGLintᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p EGLintᵖ) OnRead(ϟs *gfxapi.State) EGLintᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p EGLintᵖ) OnWrite(ϟs *gfxapi.State) EGLintᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new EGLintˢ from the pointer using start and end indices.
func (p EGLintᵖ) Slice(start, end uint64, ϟs *gfxapi.State) EGLintˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return EGLintˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p Intᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int64 element at the pointer.
func (p Intᵖ) Write(value int64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int64{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Intᵖ) OnRead(ϟs *gfxapi.State) Intᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Intᵖ) OnWrite(ϟs *gfxapi.State) Intᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Intˢ from the pointer using start and end indices.
func (p Intᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Intˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// CGLContextObjᵖ is a pointer to a CGLContextObj element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address of an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type CGLContextObjᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCGLContextObjᵖ returns a CGLContextObjᵖ that points to addr in the application pool.
func NewCGLContextObjᵖ(addr uint64) CGLContextObjᵖ {
	return CGLContextObjᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that CGLContextObjᵖ points to.
func (p CGLContextObjᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pointer.Pool == memory.ApplicationPool {
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p CGLContextObjᵖ) OnRead(ϟs *gfxapi.State) CGLContextObjᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p CGLContextObjᵖ) OnWrite(ϟs *gfxapi.State) CGLContextObjᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new CGLContextObjˢ from the pointer using start and end indices.
func (p CGLContextObjᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGLContextObjˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return CGLContextObjˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// CGSConnectionIDᵖ is a pointer to a CGSConnectionID element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address of an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type CGSConnectionIDᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCGSConnectionIDᵖ returns a CGSConnectionIDᵖ that points to addr in the application pool.
func NewCGSConnectionIDᵖ(addr uint64) CGSConnectionIDᵖ {
	return CGSConnectionIDᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that CGSConnectionIDᵖ points to.
func (p CGSConnectionIDᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pointer.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Read reads and returns the CGSConnectionID element at the pointer.
func (p CGSConnectionIDᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGSConnectionID {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the CGSConnectionID element at the pointer.
func (p CGSConnectionIDᵖ) Write(value CGSConnectionID, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]CGSConnectionID{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p CGSConnectionIDᵖ) OnRead(ϟs *gfxapi.State) CGSConnectionIDᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p CGSConnectionIDᵖ) OnWrite(ϟs *gfxapi.State) CGSConnectionIDᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new CGSConnectionIDˢ from the pointer using start and end indices.
func (p CGSConnectionIDᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGSConnectionIDˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return CGSConnectionIDˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// CGSWindowIDᵖ is a pointer to a CGSWindowID element.
type CGSWindowIDᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCGSWindowIDᵖ returns a CGSWindowIDᵖ that points to addr in the application pool.
func NewCGSWindowIDᵖ(addr uint64) CGSWindowIDᵖ {
	return CGSWindowIDᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that CGSWindowIDᵖ points to.
func (p CGSWindowIDᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the CGSWindowID element at the pointer.
func (p CGSWindowIDᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGSWindowID {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the CGSWindowID element at the pointer.
func (p CGSWindowIDᵖ) Write(value CGSWindowID, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]CGSWindowID{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p CGSWindowIDᵖ) OnRead(ϟs *gfxapi.State) CGSWindowIDᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p CGSWindowIDᵖ) OnWrite(ϟs *gfxapi.State) CGSWindowIDᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new CGSWindowIDˢ from the pointer using start and end indices.
func (p CGSWindowIDᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGSWindowIDˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return CGSWindowIDˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// CGSSurfaceIDᵖ is a pointer to a CGSSurfaceID element.
type CGSSurfaceIDᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCGSSurfaceIDᵖ returns a CGSSurfaceIDᵖ that points to addr in the application pool.
func NewCGSSurfaceIDᵖ(addr uint64) CGSSurfaceIDᵖ {
	return CGSSurfaceIDᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that CGSSurfaceIDᵖ points to.
func (p CGSSurfaceIDᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the CGSSurfaceID element at the pointer.
func (p CGSSurfaceIDᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGSSurfaceID {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the CGSSurfaceID element at the pointer.
func (p CGSSurfaceIDᵖ) Write(value CGSSurfaceID, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]CGSSurfaceID{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p CGSSurfaceIDᵖ) OnRead(ϟs *gfxapi.State) CGSSurfaceIDᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p CGSSurfaceIDᵖ) OnWrite(ϟs *gfxapi.State) CGSSurfaceIDᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new CGSSurfaceIDˢ from the pointer using start and end indices.
func (p CGSSurfaceIDᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGSSurfaceIDˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return CGSSurfaceIDˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p F64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the float64 element at the pointer.
func (p F64ᵖ) Write(value float64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]float64{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p F64ᵖ) OnRead(ϟs *gfxapi.State) F64ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p F64ᵖ) OnWrite(ϟs *gfxapi.State) F64ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new F64ˢ from the pointer using start and end indices.
func (p F64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return F64ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p S32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int32 element at the pointer.
func (p S32ᵖ) Write(value int32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int32{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p S32ᵖ) OnRead(ϟs *gfxapi.State) S32ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p S32ᵖ) OnWrite(ϟs *gfxapi.State) S32ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new S32ˢ from the pointer using start and end indices.
func (p S32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p U32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the uint32 element at the pointer.
func (p U32ᵖ) Write(value uint32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]uint32{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p U32ᵖ) OnRead(ϟs *gfxapi.State) U32ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p U32ᵖ) OnWrite(ϟs *gfxapi.State) U32ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new U32ˢ from the pointer using start and end indices.
func (p U32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// GLenumᵖ is a pointer to a GLenum element.
type GLenumᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewGLenumᵖ returns a GLenumᵖ that points to addr in the application pool.
func NewGLenumᵖ(addr uint64) GLenumᵖ {
	return GLenumᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that GLenumᵖ points to.
func (p GLenumᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the GLenum element at the pointer.
func (p GLenumᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) GLenum {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the GLenum element at the pointer.
func (p GLenumᵖ) Write(value GLenum, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]GLenum{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p GLenumᵖ) OnRead(ϟs *gfxapi.State) GLenumᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p GLenumᵖ) OnWrite(ϟs *gfxapi.State) GLenumᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new GLenumˢ from the pointer using start and end indices.
func (p GLenumᵖ) Slice(start, end uint64, ϟs *gfxapi.State) GLenumˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return GLenumˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// VertexArrayIdᵖ is a pointer to a VertexArrayId element.
type VertexArrayIdᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVertexArrayIdᵖ returns a VertexArrayIdᵖ that points to addr in the application pool.
func NewVertexArrayIdᵖ(addr uint64) VertexArrayIdᵖ {
	return VertexArrayIdᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p VertexArrayIdᵖ) OnRead(ϟs *gfxapi.State) VertexArrayIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p VertexArrayIdᵖ) OnWrite(ϟs *gfxapi.State) VertexArrayIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new VertexArrayIdˢ from the pointer using start and end indices.
func (p VertexArrayIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return VertexArrayIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// VertexArrayIdᶜᵖ is a pointer to a VertexArrayId element.
type VertexArrayIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVertexArrayIdᶜᵖ returns a VertexArrayIdᶜᵖ that points to addr in the application pool.
func NewVertexArrayIdᶜᵖ(addr uint64) VertexArrayIdᶜᵖ {
	return VertexArrayIdᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that VertexArrayIdᶜᵖ points to.
func (p VertexArrayIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the VertexArrayId element at the pointer.
func (p VertexArrayIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) VertexArrayId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the VertexArrayId element at the pointer.
func (p VertexArrayIdᶜᵖ) Write(value VertexArrayId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]VertexArrayId{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p VertexArrayIdᶜᵖ) OnRead(ϟs *gfxapi.State) VertexArrayIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p VertexArrayIdᶜᵖ) OnWrite(ϟs *gfxapi.State) VertexArrayIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new VertexArrayIdˢ from the pointer using start and end indices.
func (p VertexArrayIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return VertexArrayIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// F32ᶜᵖ is a pointer to a float32 element.
type F32ᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewF32ᶜᵖ returns a F32ᶜᵖ that points to addr in the application pool.
func NewF32ᶜᵖ(addr uint64) F32ᶜᵖ {
	return F32ᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that F32ᶜᵖ points to.
func (p F32ᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the float32 element at the pointer.
func (p F32ᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the float32 element at the pointer.
func (p F32ᶜᵖ) Write(value float32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]float32{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p F32ᶜᵖ) OnRead(ϟs *gfxapi.State) F32ᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p F32ᶜᵖ) OnWrite(ϟs *gfxapi.State) F32ᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new F32ˢ from the pointer using start and end indices.
func (p F32ᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return F32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p F32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the float32 element at the pointer.
func (p F32ᵖ) Write(value float32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]float32{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p F32ᵖ) OnRead(ϟs *gfxapi.State) F32ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p F32ᵖ) OnWrite(ϟs *gfxapi.State) F32ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new F32ˢ from the pointer using start and end indices.
func (p F32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return F32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Vec2iᵖ is a pointer to a Vec2i element.
type Vec2iᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVec2iᵖ returns a Vec2iᵖ that points to addr in the application pool.
func NewVec2iᵖ(addr uint64) Vec2iᵖ {
	return Vec2iᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Vec2iᵖ points to.
func (p Vec2iᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 2
}

// Read reads and returns the Vec2i element at the pointer.
func (p Vec2iᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec2i {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Vec2i element at the pointer.
func (p Vec2iᵖ) Write(value Vec2i, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Vec2i{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Vec2iᵖ) OnRead(ϟs *gfxapi.State) Vec2iᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Vec2iᵖ) OnWrite(ϟs *gfxapi.State) Vec2iᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Vec2iˢ from the pointer using start and end indices.
func (p Vec2iᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2iˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Vec2iˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Vec3iᵖ is a pointer to a Vec3i element.
type Vec3iᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVec3iᵖ returns a Vec3iᵖ that points to addr in the application pool.
func NewVec3iᵖ(addr uint64) Vec3iᵖ {
	return Vec3iᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Vec3iᵖ points to.
func (p Vec3iᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 3
}

// Read reads and returns the Vec3i element at the pointer.
func (p Vec3iᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec3i {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Vec3i element at the pointer.
func (p Vec3iᵖ) Write(value Vec3i, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Vec3i{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Vec3iᵖ) OnRead(ϟs *gfxapi.State) Vec3iᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Vec3iᵖ) OnWrite(ϟs *gfxapi.State) Vec3iᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Vec3iˢ from the pointer using start and end indices.
func (p Vec3iᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3iˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Vec3iˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Vec4iᵖ is a pointer to a Vec4i element.
type Vec4iᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVec4iᵖ returns a Vec4iᵖ that points to addr in the application pool.
func NewVec4iᵖ(addr uint64) Vec4iᵖ {
	return Vec4iᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Vec4iᵖ points to.
func (p Vec4iᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 4
}

// Read reads and returns the Vec4i element at the pointer.
func (p Vec4iᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec4i {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Vec4i element at the pointer.
func (p Vec4iᵖ) Write(value Vec4i, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Vec4i{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Vec4iᵖ) OnRead(ϟs *gfxapi.State) Vec4iᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Vec4iᵖ) OnWrite(ϟs *gfxapi.State) Vec4iᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Vec4iˢ from the pointer using start and end indices.
func (p Vec4iᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4iˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Vec4iˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Vec2fᵖ is a pointer to a Vec2f element.
type Vec2fᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVec2fᵖ returns a Vec2fᵖ that points to addr in the application pool.
func NewVec2fᵖ(addr uint64) Vec2fᵖ {
	return Vec2fᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Vec2fᵖ points to.
func (p Vec2fᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 2
}

// Read reads and returns the Vec2f element at the pointer.
func (p Vec2fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec2f {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Vec2f element at the pointer.
func (p Vec2fᵖ) Write(value Vec2f, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Vec2f{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Vec2fᵖ) OnRead(ϟs *gfxapi.State) Vec2fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Vec2fᵖ) OnWrite(ϟs *gfxapi.State) Vec2fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Vec2fˢ from the pointer using start and end indices.
func (p Vec2fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2fˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Vec2fˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Vec3fᵖ is a pointer to a Vec3f element.
type Vec3fᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVec3fᵖ returns a Vec3fᵖ that points to addr in the application pool.
func NewVec3fᵖ(addr uint64) Vec3fᵖ {
	return Vec3fᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Vec3fᵖ points to.
func (p Vec3fᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 3
}

// Read reads and returns the Vec3f element at the pointer.
func (p Vec3fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec3f {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Vec3f element at the pointer.
func (p Vec3fᵖ) Write(value Vec3f, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Vec3f{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Vec3fᵖ) OnRead(ϟs *gfxapi.State) Vec3fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Vec3fᵖ) OnWrite(ϟs *gfxapi.State) Vec3fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Vec3fˢ from the pointer using start and end indices.
func (p Vec3fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3fˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Vec3fˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Vec4fᵖ is a pointer to a Vec4f element.
type Vec4fᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewVec4fᵖ returns a Vec4fᵖ that points to addr in the application pool.
func NewVec4fᵖ(addr uint64) Vec4fᵖ {
	return Vec4fᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Vec4fᵖ points to.
func (p Vec4fᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 4
}

// Read reads and returns the Vec4f element at the pointer.
func (p Vec4fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec4f {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Vec4f element at the pointer.
func (p Vec4fᵖ) Write(value Vec4f, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Vec4f{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Vec4fᵖ) OnRead(ϟs *gfxapi.State) Vec4fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Vec4fᵖ) OnWrite(ϟs *gfxapi.State) Vec4fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Vec4fˢ from the pointer using start and end indices.
func (p Vec4fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4fˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Vec4fˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Mat2fᵖ is a pointer to a Mat2f element.
type Mat2fᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewMat2fᵖ returns a Mat2fᵖ that points to addr in the application pool.
func NewMat2fᵖ(addr uint64) Mat2fᵖ {
	return Mat2fᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Mat2fᵖ points to.
func (p Mat2fᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 2 * 2
}

// Read reads and returns the Mat2f element at the pointer.
func (p Mat2fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Mat2f {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Mat2f element at the pointer.
func (p Mat2fᵖ) Write(value Mat2f, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Mat2f{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Mat2fᵖ) OnRead(ϟs *gfxapi.State) Mat2fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Mat2fᵖ) OnWrite(ϟs *gfxapi.State) Mat2fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Mat2fˢ from the pointer using start and end indices.
func (p Mat2fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Mat2fˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Mat2fˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Mat3fᵖ is a pointer to a Mat3f element.
type Mat3fᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewMat3fᵖ returns a Mat3fᵖ that points to addr in the application pool.
func NewMat3fᵖ(addr uint64) Mat3fᵖ {
	return Mat3fᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Mat3fᵖ points to.
func (p Mat3fᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 3 * 3
}

// Read reads and returns the Mat3f element at the pointer.
func (p Mat3fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Mat3f {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Mat3f element at the pointer.
func (p Mat3fᵖ) Write(value Mat3f, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Mat3f{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Mat3fᵖ) OnRead(ϟs *gfxapi.State) Mat3fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Mat3fᵖ) OnWrite(ϟs *gfxapi.State) Mat3fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Mat3fˢ from the pointer using start and end indices.
func (p Mat3fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Mat3fˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Mat3fˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Mat4fᵖ is a pointer to a Mat4f element.
type Mat4fᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewMat4fᵖ returns a Mat4fᵖ that points to addr in the application pool.
func NewMat4fᵖ(addr uint64) Mat4fᵖ {
	return Mat4fᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Mat4fᵖ points to.
func (p Mat4fᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 4 * 4
}

// Read reads and returns the Mat4f element at the pointer.
func (p Mat4fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Mat4f {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Mat4f element at the pointer.
func (p Mat4fᵖ) Write(value Mat4f, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Mat4f{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Mat4fᵖ) OnRead(ϟs *gfxapi.State) Mat4fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Mat4fᵖ) OnWrite(ϟs *gfxapi.State) Mat4fᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Mat4fˢ from the pointer using start and end indices.
func (p Mat4fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Mat4fˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Mat4fˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// S32ᶜᵖ is a pointer to a int32 element.
type S32ᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewS32ᶜᵖ returns a S32ᶜᵖ that points to addr in the application pool.
func NewS32ᶜᵖ(addr uint64) S32ᶜᵖ {
	return S32ᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that S32ᶜᵖ points to.
func (p S32ᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the int32 element at the pointer.
func (p S32ᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int32 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int32 element at the pointer.
func (p S32ᶜᵖ) Write(value int32, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int32{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p S32ᶜᵖ) OnRead(ϟs *gfxapi.State) S32ᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p S32ᶜᵖ) OnWrite(ϟs *gfxapi.State) S32ᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new S32ˢ from the pointer using start and end indices.
func (p S32ᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S32ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// TextureIdᵖ is a pointer to a TextureId element.
type TextureIdᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewTextureIdᵖ returns a TextureIdᵖ that points to addr in the application pool.
func NewTextureIdᵖ(addr uint64) TextureIdᵖ {
	return TextureIdᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p TextureIdᵖ) OnRead(ϟs *gfxapi.State) TextureIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p TextureIdᵖ) OnWrite(ϟs *gfxapi.State) TextureIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new TextureIdˢ from the pointer using start and end indices.
func (p TextureIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return TextureIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// TextureIdᶜᵖ is a pointer to a TextureId element.
type TextureIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewTextureIdᶜᵖ returns a TextureIdᶜᵖ that points to addr in the application pool.
func NewTextureIdᶜᵖ(addr uint64) TextureIdᶜᵖ {
	return TextureIdᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that TextureIdᶜᵖ points to.
func (p TextureIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the TextureId element at the pointer.
func (p TextureIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) TextureId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the TextureId element at the pointer.
func (p TextureIdᶜᵖ) Write(value TextureId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]TextureId{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p TextureIdᶜᵖ) OnRead(ϟs *gfxapi.State) TextureIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p TextureIdᶜᵖ) OnWrite(ϟs *gfxapi.State) TextureIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new TextureIdˢ from the pointer using start and end indices.
func (p TextureIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return TextureIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// FramebufferIdᵖ is a pointer to a FramebufferId element.
type FramebufferIdᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewFramebufferIdᵖ returns a FramebufferIdᵖ that points to addr in the application pool.
func NewFramebufferIdᵖ(addr uint64) FramebufferIdᵖ {
	return FramebufferIdᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p FramebufferIdᵖ) OnRead(ϟs *gfxapi.State) FramebufferIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p FramebufferIdᵖ) OnWrite(ϟs *gfxapi.State) FramebufferIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new FramebufferIdˢ from the pointer using start and end indices.
func (p FramebufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return FramebufferIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// FramebufferIdᶜᵖ is a pointer to a FramebufferId element.
type FramebufferIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewFramebufferIdᶜᵖ returns a FramebufferIdᶜᵖ that points to addr in the application pool.
func NewFramebufferIdᶜᵖ(addr uint64) FramebufferIdᶜᵖ {
	return FramebufferIdᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that FramebufferIdᶜᵖ points to.
func (p FramebufferIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the FramebufferId element at the pointer.
func (p FramebufferIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) FramebufferId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the FramebufferId element at the pointer.
func (p FramebufferIdᶜᵖ) Write(value FramebufferId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]FramebufferId{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p FramebufferIdᶜᵖ) OnRead(ϟs *gfxapi.State) FramebufferIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p FramebufferIdᶜᵖ) OnWrite(ϟs *gfxapi.State) FramebufferIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new FramebufferIdˢ from the pointer using start and end indices.
func (p FramebufferIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return FramebufferIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// RenderbufferIdᵖ is a pointer to a RenderbufferId element.
type RenderbufferIdᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewRenderbufferIdᵖ returns a RenderbufferIdᵖ that points to addr in the application pool.
func NewRenderbufferIdᵖ(addr uint64) RenderbufferIdᵖ {
	return RenderbufferIdᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p RenderbufferIdᵖ) OnRead(ϟs *gfxapi.State) RenderbufferIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p RenderbufferIdᵖ) OnWrite(ϟs *gfxapi.State) RenderbufferIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new RenderbufferIdˢ from the pointer using start and end indices.
func (p RenderbufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return RenderbufferIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// RenderbufferIdᶜᵖ is a pointer to a RenderbufferId element.
type RenderbufferIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewRenderbufferIdᶜᵖ returns a RenderbufferIdᶜᵖ that points to addr in the application pool.
func NewRenderbufferIdᶜᵖ(addr uint64) RenderbufferIdᶜᵖ {
	return RenderbufferIdᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that RenderbufferIdᶜᵖ points to.
func (p RenderbufferIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the RenderbufferId element at the pointer.
func (p RenderbufferIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) RenderbufferId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the RenderbufferId element at the pointer.
func (p RenderbufferIdᶜᵖ) Write(value RenderbufferId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]RenderbufferId{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p RenderbufferIdᶜᵖ) OnRead(ϟs *gfxapi.State) RenderbufferIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p RenderbufferIdᶜᵖ) OnWrite(ϟs *gfxapi.State) RenderbufferIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new RenderbufferIdˢ from the pointer using start and end indices.
func (p RenderbufferIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return RenderbufferIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// BufferIdᵖ is a pointer to a BufferId element.
type BufferIdᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewBufferIdᵖ returns a BufferIdᵖ that points to addr in the application pool.
func NewBufferIdᵖ(addr uint64) BufferIdᵖ {
	return BufferIdᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p BufferIdᵖ) OnRead(ϟs *gfxapi.State) BufferIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p BufferIdᵖ) OnWrite(ϟs *gfxapi.State) BufferIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new BufferIdˢ from the pointer using start and end indices.
func (p BufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return BufferIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// BufferIdᶜᵖ is a pointer to a BufferId element.
type BufferIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewBufferIdᶜᵖ returns a BufferIdᶜᵖ that points to addr in the application pool.
func NewBufferIdᶜᵖ(addr uint64) BufferIdᶜᵖ {
	return BufferIdᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that BufferIdᶜᵖ points to.
func (p BufferIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the BufferId element at the pointer.
func (p BufferIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) BufferId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the BufferId element at the pointer.
func (p BufferIdᶜᵖ) Write(value BufferId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]BufferId{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p BufferIdᶜᵖ) OnRead(ϟs *gfxapi.State) BufferIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p BufferIdᶜᵖ) OnWrite(ϟs *gfxapi.State) BufferIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new BufferIdˢ from the pointer using start and end indices.
func (p BufferIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return BufferIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Charᶜᵖ is a pointer to a byte element.
type Charᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCharᶜᵖ returns a Charᶜᵖ that points to addr in the application pool.
func NewCharᶜᵖ(addr uint64) Charᶜᵖ {
	return Charᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Charᶜᵖ points to.
func (p Charᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(1)
}

// Read reads and returns the byte element at the pointer.
func (p Charᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) byte {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the byte element at the pointer.
func (p Charᶜᵖ) Write(value byte, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]byte{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Charᶜᵖ) OnRead(ϟs *gfxapi.State) Charᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Charᶜᵖ) OnWrite(ϟs *gfxapi.State) Charᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// StringSlice returns a slice starting at p and ending at the first 0 byte null-terminator.
// If incNullTerm is true then the null-terminator is included in the slice.
func (p Charᶜᵖ) StringSlice(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, incNullTerm bool) Charˢ {
	i, d := uint64(0), ϟs.MemoryDecoder(ϟs.Memory[p.Pointer.Pool].At(p.Address), ϟd, ϟl)
	for {
		if b, _ := d.Uint8(); b == 0 {
			if incNullTerm {
				return p.Slice(0, i+1, ϟs)
			} else {
				return p.Slice(0, i, ϟs)
			}
		}
		i++
	}
}

// Slice returns a new Charˢ from the pointer using start and end indices.
func (p Charᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Charᶜᵖᶜᵖ is a pointer to a Charᶜᵖ element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address of an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type Charᶜᵖᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCharᶜᵖᶜᵖ returns a Charᶜᵖᶜᵖ that points to addr in the application pool.
func NewCharᶜᵖᶜᵖ(addr uint64) Charᶜᵖᶜᵖ {
	return Charᶜᵖᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Charᶜᵖᶜᵖ points to.
func (p Charᶜᵖᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pointer.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Read reads and returns the Charᶜᵖ element at the pointer.
func (p Charᶜᵖᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Charᶜᵖ {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Charᶜᵖ element at the pointer.
func (p Charᶜᵖᶜᵖ) Write(value Charᶜᵖ, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Charᶜᵖ{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Charᶜᵖᶜᵖ) OnRead(ϟs *gfxapi.State) Charᶜᵖᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Charᶜᵖᶜᵖ) OnWrite(ϟs *gfxapi.State) Charᶜᵖᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Charᶜᵖˢ from the pointer using start and end indices.
func (p Charᶜᵖᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charᶜᵖˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charᶜᵖˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// Charᶜᵖᵖ is a pointer to a Charᶜᵖ element.
// Note: Pointers are stored differently between the application pool and internal pools.
//  * The application pool stores pointers as an address of an architecture-dependant size.
//  * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
//    pool identifier.
type Charᶜᵖᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewCharᶜᵖᵖ returns a Charᶜᵖᵖ that points to addr in the application pool.
func NewCharᶜᵖᵖ(addr uint64) Charᶜᵖᵖ {
	return Charᶜᵖᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that Charᶜᵖᵖ points to.
func (p Charᶜᵖᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	if p.Pointer.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Read reads and returns the Charᶜᵖ element at the pointer.
func (p Charᶜᵖᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Charᶜᵖ {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the Charᶜᵖ element at the pointer.
func (p Charᶜᵖᵖ) Write(value Charᶜᵖ, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]Charᶜᵖ{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Charᶜᵖᵖ) OnRead(ϟs *gfxapi.State) Charᶜᵖᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Charᶜᵖᵖ) OnWrite(ϟs *gfxapi.State) Charᶜᵖᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Charᶜᵖˢ from the pointer using start and end indices.
func (p Charᶜᵖᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charᶜᵖˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Charᶜᵖˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// ShaderIdᶜᵖ is a pointer to a ShaderId element.
type ShaderIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewShaderIdᶜᵖ returns a ShaderIdᶜᵖ that points to addr in the application pool.
func NewShaderIdᶜᵖ(addr uint64) ShaderIdᶜᵖ {
	return ShaderIdᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that ShaderIdᶜᵖ points to.
func (p ShaderIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the ShaderId element at the pointer.
func (p ShaderIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the ShaderId element at the pointer.
func (p ShaderIdᶜᵖ) Write(value ShaderId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]ShaderId{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p ShaderIdᶜᵖ) OnRead(ϟs *gfxapi.State) ShaderIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p ShaderIdᶜᵖ) OnWrite(ϟs *gfxapi.State) ShaderIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new ShaderIdˢ from the pointer using start and end indices.
func (p ShaderIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return ShaderIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// ShaderIdᵖ is a pointer to a ShaderId element.
type ShaderIdᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewShaderIdᵖ returns a ShaderIdᵖ that points to addr in the application pool.
func NewShaderIdᵖ(addr uint64) ShaderIdᵖ {
	return ShaderIdᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p ShaderIdᵖ) OnRead(ϟs *gfxapi.State) ShaderIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p ShaderIdᵖ) OnWrite(ϟs *gfxapi.State) ShaderIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new ShaderIdˢ from the pointer using start and end indices.
func (p ShaderIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return ShaderIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p Boolᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) bool {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the bool element at the pointer.
func (p Boolᵖ) Write(value bool, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]bool{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p Boolᵖ) OnRead(ϟs *gfxapi.State) Boolᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p Boolᵖ) OnWrite(ϟs *gfxapi.State) Boolᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new Boolˢ from the pointer using start and end indices.
func (p Boolᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return Boolˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// GLenumᶜᵖ is a pointer to a GLenum element.
type GLenumᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewGLenumᶜᵖ returns a GLenumᶜᵖ that points to addr in the application pool.
func NewGLenumᶜᵖ(addr uint64) GLenumᶜᵖ {
	return GLenumᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that GLenumᶜᵖ points to.
func (p GLenumᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the GLenum element at the pointer.
func (p GLenumᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) GLenum {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the GLenum element at the pointer.
func (p GLenumᶜᵖ) Write(value GLenum, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]GLenum{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p GLenumᶜᵖ) OnRead(ϟs *gfxapi.State) GLenumᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p GLenumᶜᵖ) OnWrite(ϟs *gfxapi.State) GLenumᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new GLenumˢ from the pointer using start and end indices.
func (p GLenumᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) GLenumˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return GLenumˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// QueryIdᵖ is a pointer to a QueryId element.
type QueryIdᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewQueryIdᵖ returns a QueryIdᵖ that points to addr in the application pool.
func NewQueryIdᵖ(addr uint64) QueryIdᵖ {
	return QueryIdᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
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

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p QueryIdᵖ) OnRead(ϟs *gfxapi.State) QueryIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p QueryIdᵖ) OnWrite(ϟs *gfxapi.State) QueryIdᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new QueryIdˢ from the pointer using start and end indices.
func (p QueryIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return QueryIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

// QueryIdᶜᵖ is a pointer to a QueryId element.
type QueryIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}

// NewQueryIdᶜᵖ returns a QueryIdᶜᵖ that points to addr in the application pool.
func NewQueryIdᶜᵖ(addr uint64) QueryIdᶜᵖ {
	return QueryIdᶜᵖ{Pointer: memory.Pointer{Address: addr, Pool: memory.ApplicationPool}}
}

// ElementSize returns the size in bytes of an element that QueryIdᶜᵖ points to.
func (p QueryIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Read reads and returns the QueryId element at the pointer.
func (p QueryIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) QueryId {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the QueryId element at the pointer.
func (p QueryIdᶜᵖ) Write(value QueryId, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]QueryId{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p QueryIdᶜᵖ) OnRead(ϟs *gfxapi.State) QueryIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p QueryIdᶜᵖ) OnWrite(ϟs *gfxapi.State) QueryIdᶜᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new QueryIdˢ from the pointer using start and end indices.
func (p QueryIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return QueryIdˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p S64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the int64 element at the pointer.
func (p S64ᵖ) Write(value int64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]int64{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p S64ᵖ) OnRead(ϟs *gfxapi.State) S64ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p S64ᵖ) OnWrite(ϟs *gfxapi.State) S64ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new S64ˢ from the pointer using start and end indices.
func (p S64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return S64ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
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
func (p U64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint64 {
	return p.Slice(0, 1, ϟs).Read(ϟs, ϟd, ϟl)[0]
}

// Write writes value to the uint64 element at the pointer.
func (p U64ᵖ) Write(value uint64, ϟs *gfxapi.State) {
	p.Slice(0, 1, ϟs).Write([]uint64{value}, ϟs)
}

// OnRead calls the backing pool's OnRead callback. p is returned so calls can be chained.
func (p U64ᵖ) OnRead(ϟs *gfxapi.State) U64ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnRead; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be chained.
func (p U64ᵖ) OnWrite(ϟs *gfxapi.State) U64ᵖ {
	if f := ϟs.Memory[p.Pointer.Pool].OnWrite; f != nil {
		f(p.Pointer.Range(p.ElementSize(ϟs)))
	}
	return p
}

// Slice returns a new U64ˢ from the pointer using start and end indices.
func (p U64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ {
	if start > end {
		panic(fmt.Errorf("Slice start (%d) is greater than the end (%d)", start, end))
	}
	return U64ˢ{SliceInfo: SliceInfo{Root: p.Pointer, Base: p.Address + start*p.ElementSize(ϟs), Count: end - start}}
}

type S32ː2ᵃ struct {
	binary.Generate
	Elements [2]int32
}
type S32ː3ᵃ struct {
	binary.Generate
	Elements [3]int32
}
type S32ː4ᵃ struct {
	binary.Generate
	Elements [4]int32
}
type F32ː2ᵃ struct {
	binary.Generate
	Elements [2]float32
}
type F32ː3ᵃ struct {
	binary.Generate
	Elements [3]float32
}
type F32ː4ᵃ struct {
	binary.Generate
	Elements [4]float32
}
type Vec2fː2ᵃ struct {
	binary.Generate
	Elements [2]Vec2f
}
type Vec3fː3ᵃ struct {
	binary.Generate
	Elements [3]Vec3f
}
type Vec4fː4ᵃ struct {
	binary.Generate
	Elements [4]Vec4f
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
func (s Boolˢ) Clone(ϟs *gfxapi.State) Boolˢ {
	s.OnRead(ϟs)
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
func (s Boolˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []bool {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]bool, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s Boolˢ) Write(src []bool, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Bool(bool(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Boolˢ) Copy(src Boolˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Boolˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Boolˢ) OnRead(ϟs *gfxapi.State) Boolˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Boolˢ) OnWrite(ϟs *gfxapi.State) Boolˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Boolᵖ to the i'th element in this Boolˢ.
func (s Boolˢ) Index(i uint64, ϟs *gfxapi.State) Boolᵖ {
	return Boolᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Boolˢ using start and end indices.
func (s Boolˢ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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

// BufferIdˢ is a slice of BufferId.
type BufferIdˢ struct {
	binary.Generate
	SliceInfo
}

// MakeBufferIdˢ returns a BufferIdˢ backed by a new memory pool.
func MakeBufferIdˢ(count uint64, ϟs *gfxapi.State) BufferIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return BufferIdˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the BufferIdˢ in a new memory pool.
func (s BufferIdˢ) Clone(ϟs *gfxapi.State) BufferIdˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := BufferIdˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that BufferIdˢ points to.
func (s BufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s BufferIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s BufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s BufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s BufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsBufferIdˢ returns s cast to a BufferIdˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsBufferIdˢ(s Slice, ϟs *gfxapi.State) BufferIdˢ {
	out := BufferIdˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the BufferId elements in this BufferIdˢ.
func (s BufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []BufferId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]BufferId, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s BufferIdˢ) Write(src []BufferId, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst BufferIdˢ) Copy(src BufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s BufferIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s BufferIdˢ) OnRead(ϟs *gfxapi.State) BufferIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s BufferIdˢ) OnWrite(ϟs *gfxapi.State) BufferIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a BufferIdᵖ to the i'th element in this BufferIdˢ.
func (s BufferIdˢ) Index(i uint64, ϟs *gfxapi.State) BufferIdᵖ {
	return BufferIdᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the BufferIdˢ using start and end indices.
func (s BufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return BufferIdˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the BufferIdˢ slice.
func (s BufferIdˢ) String() string {
	return fmt.Sprintf("BufferId(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// CGLContextObjˢ is a slice of CGLContextObj.
type CGLContextObjˢ struct {
	binary.Generate
	SliceInfo
}

// MakeCGLContextObjˢ returns a CGLContextObjˢ backed by a new memory pool.
func MakeCGLContextObjˢ(count uint64, ϟs *gfxapi.State) CGLContextObjˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return CGLContextObjˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the CGLContextObjˢ in a new memory pool.
func (s CGLContextObjˢ) Clone(ϟs *gfxapi.State) CGLContextObjˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := CGLContextObjˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that CGLContextObjˢ points to.
func (s CGLContextObjˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	if s.Root.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Range returns the memory range this slice represents in the underlying pool.
func (s CGLContextObjˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s CGLContextObjˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s CGLContextObjˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s CGLContextObjˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsCGLContextObjˢ returns s cast to a CGLContextObjˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsCGLContextObjˢ(s Slice, ϟs *gfxapi.State) CGLContextObjˢ {
	out := CGLContextObjˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the CGLContextObj elements in this CGLContextObjˢ.
func (s CGLContextObjˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGLContextObj {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]CGLContextObj, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if s.Root.Pool == memory.ApplicationPool {
			ptr, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			res[i] = NewCGLContextObj(ptr)
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
func (s CGLContextObjˢ) Write(src []CGLContextObj, ϟs *gfxapi.State) uint64 {
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
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst CGLContextObjˢ) Copy(src CGLContextObjˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGLContextObjˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	if (dst.Root.Pool == memory.ApplicationPool) != (src.Root.Pool == memory.ApplicationPool) {
		dst.Write(src.Read(ϟs, ϟd, ϟl), ϟs) // Element-wise copy so we can convert u64 <-> CGLContextObjᵖ
	} else {
		src.OnRead(ϟs)
		ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
		dst.OnWrite(ϟs)
	}
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s CGLContextObjˢ) OnRead(ϟs *gfxapi.State) CGLContextObjˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s CGLContextObjˢ) OnWrite(ϟs *gfxapi.State) CGLContextObjˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a CGLContextObjᵖ to the i'th element in this CGLContextObjˢ.
func (s CGLContextObjˢ) Index(i uint64, ϟs *gfxapi.State) CGLContextObjᵖ {
	return CGLContextObjᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the CGLContextObjˢ using start and end indices.
func (s CGLContextObjˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGLContextObjˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return CGLContextObjˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the CGLContextObjˢ slice.
func (s CGLContextObjˢ) String() string {
	return fmt.Sprintf("CGLContextObj(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// CGSConnectionIDˢ is a slice of CGSConnectionID.
type CGSConnectionIDˢ struct {
	binary.Generate
	SliceInfo
}

// MakeCGSConnectionIDˢ returns a CGSConnectionIDˢ backed by a new memory pool.
func MakeCGSConnectionIDˢ(count uint64, ϟs *gfxapi.State) CGSConnectionIDˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return CGSConnectionIDˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the CGSConnectionIDˢ in a new memory pool.
func (s CGSConnectionIDˢ) Clone(ϟs *gfxapi.State) CGSConnectionIDˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := CGSConnectionIDˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that CGSConnectionIDˢ points to.
func (s CGSConnectionIDˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	if s.Root.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Range returns the memory range this slice represents in the underlying pool.
func (s CGSConnectionIDˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s CGSConnectionIDˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s CGSConnectionIDˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s CGSConnectionIDˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsCGSConnectionIDˢ returns s cast to a CGSConnectionIDˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsCGSConnectionIDˢ(s Slice, ϟs *gfxapi.State) CGSConnectionIDˢ {
	out := CGSConnectionIDˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the CGSConnectionID elements in this CGSConnectionIDˢ.
func (s CGSConnectionIDˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGSConnectionID {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]CGSConnectionID, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if s.Root.Pool == memory.ApplicationPool {
			ptr, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			res[i] = NewCGSConnectionID(ptr)
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
func (s CGSConnectionIDˢ) Write(src []CGSConnectionID, ϟs *gfxapi.State) uint64 {
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
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst CGSConnectionIDˢ) Copy(src CGSConnectionIDˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGSConnectionIDˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	if (dst.Root.Pool == memory.ApplicationPool) != (src.Root.Pool == memory.ApplicationPool) {
		dst.Write(src.Read(ϟs, ϟd, ϟl), ϟs) // Element-wise copy so we can convert u64 <-> CGSConnectionIDᵖ
	} else {
		src.OnRead(ϟs)
		ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
		dst.OnWrite(ϟs)
	}
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s CGSConnectionIDˢ) OnRead(ϟs *gfxapi.State) CGSConnectionIDˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s CGSConnectionIDˢ) OnWrite(ϟs *gfxapi.State) CGSConnectionIDˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a CGSConnectionIDᵖ to the i'th element in this CGSConnectionIDˢ.
func (s CGSConnectionIDˢ) Index(i uint64, ϟs *gfxapi.State) CGSConnectionIDᵖ {
	return CGSConnectionIDᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the CGSConnectionIDˢ using start and end indices.
func (s CGSConnectionIDˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGSConnectionIDˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return CGSConnectionIDˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the CGSConnectionIDˢ slice.
func (s CGSConnectionIDˢ) String() string {
	return fmt.Sprintf("CGSConnectionID(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// CGSSurfaceIDˢ is a slice of CGSSurfaceID.
type CGSSurfaceIDˢ struct {
	binary.Generate
	SliceInfo
}

// MakeCGSSurfaceIDˢ returns a CGSSurfaceIDˢ backed by a new memory pool.
func MakeCGSSurfaceIDˢ(count uint64, ϟs *gfxapi.State) CGSSurfaceIDˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return CGSSurfaceIDˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the CGSSurfaceIDˢ in a new memory pool.
func (s CGSSurfaceIDˢ) Clone(ϟs *gfxapi.State) CGSSurfaceIDˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := CGSSurfaceIDˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that CGSSurfaceIDˢ points to.
func (s CGSSurfaceIDˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s CGSSurfaceIDˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s CGSSurfaceIDˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s CGSSurfaceIDˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s CGSSurfaceIDˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsCGSSurfaceIDˢ returns s cast to a CGSSurfaceIDˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsCGSSurfaceIDˢ(s Slice, ϟs *gfxapi.State) CGSSurfaceIDˢ {
	out := CGSSurfaceIDˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the CGSSurfaceID elements in this CGSSurfaceIDˢ.
func (s CGSSurfaceIDˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGSSurfaceID {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]CGSSurfaceID, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if v, err := d.Int32(); err == nil {
			res[i] = CGSSurfaceID(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s CGSSurfaceIDˢ) Write(src []CGSSurfaceID, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int32(int32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst CGSSurfaceIDˢ) Copy(src CGSSurfaceIDˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGSSurfaceIDˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s CGSSurfaceIDˢ) OnRead(ϟs *gfxapi.State) CGSSurfaceIDˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s CGSSurfaceIDˢ) OnWrite(ϟs *gfxapi.State) CGSSurfaceIDˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a CGSSurfaceIDᵖ to the i'th element in this CGSSurfaceIDˢ.
func (s CGSSurfaceIDˢ) Index(i uint64, ϟs *gfxapi.State) CGSSurfaceIDᵖ {
	return CGSSurfaceIDᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the CGSSurfaceIDˢ using start and end indices.
func (s CGSSurfaceIDˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGSSurfaceIDˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return CGSSurfaceIDˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the CGSSurfaceIDˢ slice.
func (s CGSSurfaceIDˢ) String() string {
	return fmt.Sprintf("CGSSurfaceID(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// CGSWindowIDˢ is a slice of CGSWindowID.
type CGSWindowIDˢ struct {
	binary.Generate
	SliceInfo
}

// MakeCGSWindowIDˢ returns a CGSWindowIDˢ backed by a new memory pool.
func MakeCGSWindowIDˢ(count uint64, ϟs *gfxapi.State) CGSWindowIDˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return CGSWindowIDˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the CGSWindowIDˢ in a new memory pool.
func (s CGSWindowIDˢ) Clone(ϟs *gfxapi.State) CGSWindowIDˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := CGSWindowIDˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that CGSWindowIDˢ points to.
func (s CGSWindowIDˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s CGSWindowIDˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s CGSWindowIDˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s CGSWindowIDˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s CGSWindowIDˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsCGSWindowIDˢ returns s cast to a CGSWindowIDˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsCGSWindowIDˢ(s Slice, ϟs *gfxapi.State) CGSWindowIDˢ {
	out := CGSWindowIDˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the CGSWindowID elements in this CGSWindowIDˢ.
func (s CGSWindowIDˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGSWindowID {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]CGSWindowID, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if v, err := d.Int32(); err == nil {
			res[i] = CGSWindowID(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s CGSWindowIDˢ) Write(src []CGSWindowID, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int32(int32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst CGSWindowIDˢ) Copy(src CGSWindowIDˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGSWindowIDˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s CGSWindowIDˢ) OnRead(ϟs *gfxapi.State) CGSWindowIDˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s CGSWindowIDˢ) OnWrite(ϟs *gfxapi.State) CGSWindowIDˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a CGSWindowIDᵖ to the i'th element in this CGSWindowIDˢ.
func (s CGSWindowIDˢ) Index(i uint64, ϟs *gfxapi.State) CGSWindowIDᵖ {
	return CGSWindowIDᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the CGSWindowIDˢ using start and end indices.
func (s CGSWindowIDˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGSWindowIDˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return CGSWindowIDˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the CGSWindowIDˢ slice.
func (s CGSWindowIDˢ) String() string {
	return fmt.Sprintf("CGSWindowID(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
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
func (s Charˢ) Clone(ϟs *gfxapi.State) Charˢ {
	s.OnRead(ϟs)
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
func (s Charˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []byte {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]byte, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s Charˢ) Write(src []byte, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint8(uint8(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Charˢ) Copy(src Charˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Charˢ) OnRead(ϟs *gfxapi.State) Charˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Charˢ) OnWrite(ϟs *gfxapi.State) Charˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Charᵖ to the i'th element in this Charˢ.
func (s Charˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖ {
	return Charᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Charˢ using start and end indices.
func (s Charˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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

// Charᶜᵖˢ is a slice of Charᶜᵖ.
type Charᶜᵖˢ struct {
	binary.Generate
	SliceInfo
}

// MakeCharᶜᵖˢ returns a Charᶜᵖˢ backed by a new memory pool.
func MakeCharᶜᵖˢ(count uint64, ϟs *gfxapi.State) Charᶜᵖˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Charᶜᵖˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Charᶜᵖˢ in a new memory pool.
func (s Charᶜᵖˢ) Clone(ϟs *gfxapi.State) Charᶜᵖˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Charᶜᵖˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Charᶜᵖˢ points to.
func (s Charᶜᵖˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	if s.Root.Pool == memory.ApplicationPool {
		return uint64(ϟs.Architecture.PointerSize)
	} else {
		return 12
	}
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Charᶜᵖˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Charᶜᵖˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Charᶜᵖˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Charᶜᵖˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsCharᶜᵖˢ returns s cast to a Charᶜᵖˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsCharᶜᵖˢ(s Slice, ϟs *gfxapi.State) Charᶜᵖˢ {
	out := Charᶜᵖˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Charᶜᵖ elements in this Charᶜᵖˢ.
func (s Charᶜᵖˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Charᶜᵖ {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Charᶜᵖ, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if s.Root.Pool == memory.ApplicationPool {
			ptr, err := binary.ReadUint(d, ϟs.Architecture.PointerSize*8)
			if err != nil {
				panic(err)
			}
			res[i] = NewCharᶜᵖ(ptr)
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
func (s Charᶜᵖˢ) Write(src []Charᶜᵖ, ϟs *gfxapi.State) uint64 {
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
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Charᶜᵖˢ) Copy(src Charᶜᵖˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charᶜᵖˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	if (dst.Root.Pool == memory.ApplicationPool) != (src.Root.Pool == memory.ApplicationPool) {
		dst.Write(src.Read(ϟs, ϟd, ϟl), ϟs) // Element-wise copy so we can convert u64 <-> Charᶜᵖᵖ
	} else {
		src.OnRead(ϟs)
		ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
		dst.OnWrite(ϟs)
	}
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Charᶜᵖˢ) OnRead(ϟs *gfxapi.State) Charᶜᵖˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Charᶜᵖˢ) OnWrite(ϟs *gfxapi.State) Charᶜᵖˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Charᶜᵖᵖ to the i'th element in this Charᶜᵖˢ.
func (s Charᶜᵖˢ) Index(i uint64, ϟs *gfxapi.State) Charᶜᵖᵖ {
	return Charᶜᵖᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Charᶜᵖˢ using start and end indices.
func (s Charᶜᵖˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charᶜᵖˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Charᶜᵖˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Charᶜᵖˢ slice.
func (s Charᶜᵖˢ) String() string {
	return fmt.Sprintf("Charᶜᵖ(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// EGLintˢ is a slice of EGLint.
type EGLintˢ struct {
	binary.Generate
	SliceInfo
}

// MakeEGLintˢ returns a EGLintˢ backed by a new memory pool.
func MakeEGLintˢ(count uint64, ϟs *gfxapi.State) EGLintˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return EGLintˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the EGLintˢ in a new memory pool.
func (s EGLintˢ) Clone(ϟs *gfxapi.State) EGLintˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := EGLintˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that EGLintˢ points to.
func (s EGLintˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(ϟs.Architecture.IntegerSize)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s EGLintˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s EGLintˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s EGLintˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s EGLintˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsEGLintˢ returns s cast to a EGLintˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsEGLintˢ(s Slice, ϟs *gfxapi.State) EGLintˢ {
	out := EGLintˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the EGLint elements in this EGLintˢ.
func (s EGLintˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []EGLint {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]EGLint, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s EGLintˢ) Write(src []EGLint, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := binary.WriteInt(e, ϟs.Architecture.IntegerSize*8, int64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst EGLintˢ) Copy(src EGLintˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s EGLintˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s EGLintˢ) OnRead(ϟs *gfxapi.State) EGLintˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s EGLintˢ) OnWrite(ϟs *gfxapi.State) EGLintˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a EGLintᵖ to the i'th element in this EGLintˢ.
func (s EGLintˢ) Index(i uint64, ϟs *gfxapi.State) EGLintᵖ {
	return EGLintᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the EGLintˢ using start and end indices.
func (s EGLintˢ) Slice(start, end uint64, ϟs *gfxapi.State) EGLintˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return EGLintˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the EGLintˢ slice.
func (s EGLintˢ) String() string {
	return fmt.Sprintf("EGLint(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
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
func (s F32ˢ) Clone(ϟs *gfxapi.State) F32ˢ {
	s.OnRead(ϟs)
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
func (s F32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]float32, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s F32ˢ) Write(src []float32, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Float32(float32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst F32ˢ) Copy(src F32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s F32ˢ) OnRead(ϟs *gfxapi.State) F32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s F32ˢ) OnWrite(ϟs *gfxapi.State) F32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a F32ᵖ to the i'th element in this F32ˢ.
func (s F32ˢ) Index(i uint64, ϟs *gfxapi.State) F32ᵖ {
	return F32ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the F32ˢ using start and end indices.
func (s F32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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
func (s F64ˢ) Clone(ϟs *gfxapi.State) F64ˢ {
	s.OnRead(ϟs)
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
func (s F64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]float64, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s F64ˢ) Write(src []float64, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Float64(float64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst F64ˢ) Copy(src F64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s F64ˢ) OnRead(ϟs *gfxapi.State) F64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s F64ˢ) OnWrite(ϟs *gfxapi.State) F64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a F64ᵖ to the i'th element in this F64ˢ.
func (s F64ˢ) Index(i uint64, ϟs *gfxapi.State) F64ᵖ {
	return F64ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the F64ˢ using start and end indices.
func (s F64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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

// FramebufferIdˢ is a slice of FramebufferId.
type FramebufferIdˢ struct {
	binary.Generate
	SliceInfo
}

// MakeFramebufferIdˢ returns a FramebufferIdˢ backed by a new memory pool.
func MakeFramebufferIdˢ(count uint64, ϟs *gfxapi.State) FramebufferIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return FramebufferIdˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the FramebufferIdˢ in a new memory pool.
func (s FramebufferIdˢ) Clone(ϟs *gfxapi.State) FramebufferIdˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := FramebufferIdˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that FramebufferIdˢ points to.
func (s FramebufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s FramebufferIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s FramebufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s FramebufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s FramebufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsFramebufferIdˢ returns s cast to a FramebufferIdˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsFramebufferIdˢ(s Slice, ϟs *gfxapi.State) FramebufferIdˢ {
	out := FramebufferIdˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the FramebufferId elements in this FramebufferIdˢ.
func (s FramebufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []FramebufferId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]FramebufferId, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s FramebufferIdˢ) Write(src []FramebufferId, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst FramebufferIdˢ) Copy(src FramebufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s FramebufferIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s FramebufferIdˢ) OnRead(ϟs *gfxapi.State) FramebufferIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s FramebufferIdˢ) OnWrite(ϟs *gfxapi.State) FramebufferIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a FramebufferIdᵖ to the i'th element in this FramebufferIdˢ.
func (s FramebufferIdˢ) Index(i uint64, ϟs *gfxapi.State) FramebufferIdᵖ {
	return FramebufferIdᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the FramebufferIdˢ using start and end indices.
func (s FramebufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return FramebufferIdˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the FramebufferIdˢ slice.
func (s FramebufferIdˢ) String() string {
	return fmt.Sprintf("FramebufferId(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// GLenumˢ is a slice of GLenum.
type GLenumˢ struct {
	binary.Generate
	SliceInfo
}

// MakeGLenumˢ returns a GLenumˢ backed by a new memory pool.
func MakeGLenumˢ(count uint64, ϟs *gfxapi.State) GLenumˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return GLenumˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the GLenumˢ in a new memory pool.
func (s GLenumˢ) Clone(ϟs *gfxapi.State) GLenumˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := GLenumˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that GLenumˢ points to.
func (s GLenumˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s GLenumˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s GLenumˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s GLenumˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s GLenumˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsGLenumˢ returns s cast to a GLenumˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsGLenumˢ(s Slice, ϟs *gfxapi.State) GLenumˢ {
	out := GLenumˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the GLenum elements in this GLenumˢ.
func (s GLenumˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []GLenum {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]GLenum, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if v, err := d.Uint32(); err == nil {
			res[i] = GLenum(v)
		} else {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s GLenumˢ) Write(src []GLenum, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst GLenumˢ) Copy(src GLenumˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s GLenumˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s GLenumˢ) OnRead(ϟs *gfxapi.State) GLenumˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s GLenumˢ) OnWrite(ϟs *gfxapi.State) GLenumˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a GLenumᵖ to the i'th element in this GLenumˢ.
func (s GLenumˢ) Index(i uint64, ϟs *gfxapi.State) GLenumᵖ {
	return GLenumᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the GLenumˢ using start and end indices.
func (s GLenumˢ) Slice(start, end uint64, ϟs *gfxapi.State) GLenumˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return GLenumˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the GLenumˢ slice.
func (s GLenumˢ) String() string {
	return fmt.Sprintf("GLenum(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
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
func (s Intˢ) Clone(ϟs *gfxapi.State) Intˢ {
	s.OnRead(ϟs)
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
func (s Intˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int64, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s Intˢ) Write(src []int64, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := binary.WriteInt(e, ϟs.Architecture.IntegerSize*8, int64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Intˢ) Copy(src Intˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Intˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Intˢ) OnRead(ϟs *gfxapi.State) Intˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Intˢ) OnWrite(ϟs *gfxapi.State) Intˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Intᵖ to the i'th element in this Intˢ.
func (s Intˢ) Index(i uint64, ϟs *gfxapi.State) Intᵖ {
	return Intᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Intˢ using start and end indices.
func (s Intˢ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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

// Mat2fˢ is a slice of Mat2f.
type Mat2fˢ struct {
	binary.Generate
	SliceInfo
}

// MakeMat2fˢ returns a Mat2fˢ backed by a new memory pool.
func MakeMat2fˢ(count uint64, ϟs *gfxapi.State) Mat2fˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Mat2fˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Mat2fˢ in a new memory pool.
func (s Mat2fˢ) Clone(ϟs *gfxapi.State) Mat2fˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Mat2fˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Mat2fˢ points to.
func (s Mat2fˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 2 * 2
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Mat2fˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Mat2fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Mat2fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Mat2fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsMat2fˢ returns s cast to a Mat2fˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsMat2fˢ(s Slice, ϟs *gfxapi.State) Mat2fˢ {
	out := Mat2fˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Mat2f elements in this Mat2fˢ.
func (s Mat2fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Mat2f {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Mat2f, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Mat2fˢ) Write(src []Mat2f, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Mat2fˢ) Copy(src Mat2fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Mat2fˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Mat2fˢ) OnRead(ϟs *gfxapi.State) Mat2fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Mat2fˢ) OnWrite(ϟs *gfxapi.State) Mat2fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Mat2fᵖ to the i'th element in this Mat2fˢ.
func (s Mat2fˢ) Index(i uint64, ϟs *gfxapi.State) Mat2fᵖ {
	return Mat2fᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Mat2fˢ using start and end indices.
func (s Mat2fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Mat2fˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Mat2fˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Mat2fˢ slice.
func (s Mat2fˢ) String() string {
	return fmt.Sprintf("Mat2f(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Mat3fˢ is a slice of Mat3f.
type Mat3fˢ struct {
	binary.Generate
	SliceInfo
}

// MakeMat3fˢ returns a Mat3fˢ backed by a new memory pool.
func MakeMat3fˢ(count uint64, ϟs *gfxapi.State) Mat3fˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Mat3fˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Mat3fˢ in a new memory pool.
func (s Mat3fˢ) Clone(ϟs *gfxapi.State) Mat3fˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Mat3fˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Mat3fˢ points to.
func (s Mat3fˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 3 * 3
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Mat3fˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Mat3fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Mat3fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Mat3fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsMat3fˢ returns s cast to a Mat3fˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsMat3fˢ(s Slice, ϟs *gfxapi.State) Mat3fˢ {
	out := Mat3fˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Mat3f elements in this Mat3fˢ.
func (s Mat3fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Mat3f {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Mat3f, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Mat3fˢ) Write(src []Mat3f, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Mat3fˢ) Copy(src Mat3fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Mat3fˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Mat3fˢ) OnRead(ϟs *gfxapi.State) Mat3fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Mat3fˢ) OnWrite(ϟs *gfxapi.State) Mat3fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Mat3fᵖ to the i'th element in this Mat3fˢ.
func (s Mat3fˢ) Index(i uint64, ϟs *gfxapi.State) Mat3fᵖ {
	return Mat3fᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Mat3fˢ using start and end indices.
func (s Mat3fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Mat3fˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Mat3fˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Mat3fˢ slice.
func (s Mat3fˢ) String() string {
	return fmt.Sprintf("Mat3f(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Mat4fˢ is a slice of Mat4f.
type Mat4fˢ struct {
	binary.Generate
	SliceInfo
}

// MakeMat4fˢ returns a Mat4fˢ backed by a new memory pool.
func MakeMat4fˢ(count uint64, ϟs *gfxapi.State) Mat4fˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Mat4fˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Mat4fˢ in a new memory pool.
func (s Mat4fˢ) Clone(ϟs *gfxapi.State) Mat4fˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Mat4fˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Mat4fˢ points to.
func (s Mat4fˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 4 * 4
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Mat4fˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Mat4fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Mat4fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Mat4fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsMat4fˢ returns s cast to a Mat4fˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsMat4fˢ(s Slice, ϟs *gfxapi.State) Mat4fˢ {
	out := Mat4fˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Mat4f elements in this Mat4fˢ.
func (s Mat4fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Mat4f {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Mat4f, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Mat4fˢ) Write(src []Mat4f, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Mat4fˢ) Copy(src Mat4fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Mat4fˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Mat4fˢ) OnRead(ϟs *gfxapi.State) Mat4fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Mat4fˢ) OnWrite(ϟs *gfxapi.State) Mat4fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Mat4fᵖ to the i'th element in this Mat4fˢ.
func (s Mat4fˢ) Index(i uint64, ϟs *gfxapi.State) Mat4fᵖ {
	return Mat4fᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Mat4fˢ using start and end indices.
func (s Mat4fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Mat4fˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Mat4fˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Mat4fˢ slice.
func (s Mat4fˢ) String() string {
	return fmt.Sprintf("Mat4f(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// QueryIdˢ is a slice of QueryId.
type QueryIdˢ struct {
	binary.Generate
	SliceInfo
}

// MakeQueryIdˢ returns a QueryIdˢ backed by a new memory pool.
func MakeQueryIdˢ(count uint64, ϟs *gfxapi.State) QueryIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return QueryIdˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the QueryIdˢ in a new memory pool.
func (s QueryIdˢ) Clone(ϟs *gfxapi.State) QueryIdˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := QueryIdˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that QueryIdˢ points to.
func (s QueryIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s QueryIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s QueryIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s QueryIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s QueryIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsQueryIdˢ returns s cast to a QueryIdˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsQueryIdˢ(s Slice, ϟs *gfxapi.State) QueryIdˢ {
	out := QueryIdˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the QueryId elements in this QueryIdˢ.
func (s QueryIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []QueryId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]QueryId, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s QueryIdˢ) Write(src []QueryId, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst QueryIdˢ) Copy(src QueryIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s QueryIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s QueryIdˢ) OnRead(ϟs *gfxapi.State) QueryIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s QueryIdˢ) OnWrite(ϟs *gfxapi.State) QueryIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a QueryIdᵖ to the i'th element in this QueryIdˢ.
func (s QueryIdˢ) Index(i uint64, ϟs *gfxapi.State) QueryIdᵖ {
	return QueryIdᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the QueryIdˢ using start and end indices.
func (s QueryIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return QueryIdˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the QueryIdˢ slice.
func (s QueryIdˢ) String() string {
	return fmt.Sprintf("QueryId(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// RenderbufferIdˢ is a slice of RenderbufferId.
type RenderbufferIdˢ struct {
	binary.Generate
	SliceInfo
}

// MakeRenderbufferIdˢ returns a RenderbufferIdˢ backed by a new memory pool.
func MakeRenderbufferIdˢ(count uint64, ϟs *gfxapi.State) RenderbufferIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return RenderbufferIdˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the RenderbufferIdˢ in a new memory pool.
func (s RenderbufferIdˢ) Clone(ϟs *gfxapi.State) RenderbufferIdˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := RenderbufferIdˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that RenderbufferIdˢ points to.
func (s RenderbufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s RenderbufferIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s RenderbufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s RenderbufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s RenderbufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsRenderbufferIdˢ returns s cast to a RenderbufferIdˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsRenderbufferIdˢ(s Slice, ϟs *gfxapi.State) RenderbufferIdˢ {
	out := RenderbufferIdˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the RenderbufferId elements in this RenderbufferIdˢ.
func (s RenderbufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []RenderbufferId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]RenderbufferId, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s RenderbufferIdˢ) Write(src []RenderbufferId, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst RenderbufferIdˢ) Copy(src RenderbufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s RenderbufferIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s RenderbufferIdˢ) OnRead(ϟs *gfxapi.State) RenderbufferIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s RenderbufferIdˢ) OnWrite(ϟs *gfxapi.State) RenderbufferIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a RenderbufferIdᵖ to the i'th element in this RenderbufferIdˢ.
func (s RenderbufferIdˢ) Index(i uint64, ϟs *gfxapi.State) RenderbufferIdᵖ {
	return RenderbufferIdᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the RenderbufferIdˢ using start and end indices.
func (s RenderbufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return RenderbufferIdˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the RenderbufferIdˢ slice.
func (s RenderbufferIdˢ) String() string {
	return fmt.Sprintf("RenderbufferId(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
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
func (s S32ˢ) Clone(ϟs *gfxapi.State) S32ˢ {
	s.OnRead(ϟs)
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
func (s S32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int32, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s S32ˢ) Write(src []int32, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int32(int32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S32ˢ) Copy(src S32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s S32ˢ) OnRead(ϟs *gfxapi.State) S32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s S32ˢ) OnWrite(ϟs *gfxapi.State) S32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a S32ᵖ to the i'th element in this S32ˢ.
func (s S32ˢ) Index(i uint64, ϟs *gfxapi.State) S32ᵖ {
	return S32ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the S32ˢ using start and end indices.
func (s S32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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
func (s S64ˢ) Clone(ϟs *gfxapi.State) S64ˢ {
	s.OnRead(ϟs)
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
func (s S64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]int64, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s S64ˢ) Write(src []int64, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Int64(int64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst S64ˢ) Copy(src S64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s S64ˢ) OnRead(ϟs *gfxapi.State) S64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s S64ˢ) OnWrite(ϟs *gfxapi.State) S64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a S64ᵖ to the i'th element in this S64ˢ.
func (s S64ˢ) Index(i uint64, ϟs *gfxapi.State) S64ᵖ {
	return S64ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the S64ˢ using start and end indices.
func (s S64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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

// ShaderIdˢ is a slice of ShaderId.
type ShaderIdˢ struct {
	binary.Generate
	SliceInfo
}

// MakeShaderIdˢ returns a ShaderIdˢ backed by a new memory pool.
func MakeShaderIdˢ(count uint64, ϟs *gfxapi.State) ShaderIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return ShaderIdˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the ShaderIdˢ in a new memory pool.
func (s ShaderIdˢ) Clone(ϟs *gfxapi.State) ShaderIdˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := ShaderIdˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that ShaderIdˢ points to.
func (s ShaderIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s ShaderIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s ShaderIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s ShaderIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s ShaderIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsShaderIdˢ returns s cast to a ShaderIdˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsShaderIdˢ(s Slice, ϟs *gfxapi.State) ShaderIdˢ {
	out := ShaderIdˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the ShaderId elements in this ShaderIdˢ.
func (s ShaderIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []ShaderId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]ShaderId, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s ShaderIdˢ) Write(src []ShaderId, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst ShaderIdˢ) Copy(src ShaderIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s ShaderIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s ShaderIdˢ) OnRead(ϟs *gfxapi.State) ShaderIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s ShaderIdˢ) OnWrite(ϟs *gfxapi.State) ShaderIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a ShaderIdᵖ to the i'th element in this ShaderIdˢ.
func (s ShaderIdˢ) Index(i uint64, ϟs *gfxapi.State) ShaderIdᵖ {
	return ShaderIdᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the ShaderIdˢ using start and end indices.
func (s ShaderIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return ShaderIdˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the ShaderIdˢ slice.
func (s ShaderIdˢ) String() string {
	return fmt.Sprintf("ShaderId(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// TextureIdˢ is a slice of TextureId.
type TextureIdˢ struct {
	binary.Generate
	SliceInfo
}

// MakeTextureIdˢ returns a TextureIdˢ backed by a new memory pool.
func MakeTextureIdˢ(count uint64, ϟs *gfxapi.State) TextureIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return TextureIdˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the TextureIdˢ in a new memory pool.
func (s TextureIdˢ) Clone(ϟs *gfxapi.State) TextureIdˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := TextureIdˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that TextureIdˢ points to.
func (s TextureIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s TextureIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s TextureIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s TextureIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s TextureIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsTextureIdˢ returns s cast to a TextureIdˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsTextureIdˢ(s Slice, ϟs *gfxapi.State) TextureIdˢ {
	out := TextureIdˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the TextureId elements in this TextureIdˢ.
func (s TextureIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []TextureId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]TextureId, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s TextureIdˢ) Write(src []TextureId, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst TextureIdˢ) Copy(src TextureIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s TextureIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s TextureIdˢ) OnRead(ϟs *gfxapi.State) TextureIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s TextureIdˢ) OnWrite(ϟs *gfxapi.State) TextureIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a TextureIdᵖ to the i'th element in this TextureIdˢ.
func (s TextureIdˢ) Index(i uint64, ϟs *gfxapi.State) TextureIdᵖ {
	return TextureIdᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the TextureIdˢ using start and end indices.
func (s TextureIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return TextureIdˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the TextureIdˢ slice.
func (s TextureIdˢ) String() string {
	return fmt.Sprintf("TextureId(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
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
func (s U32ˢ) Clone(ϟs *gfxapi.State) U32ˢ {
	s.OnRead(ϟs)
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
func (s U32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint32 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint32, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s U32ˢ) Write(src []uint32, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U32ˢ) Copy(src U32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U32ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s U32ˢ) OnRead(ϟs *gfxapi.State) U32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s U32ˢ) OnWrite(ϟs *gfxapi.State) U32ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a U32ᵖ to the i'th element in this U32ˢ.
func (s U32ˢ) Index(i uint64, ϟs *gfxapi.State) U32ᵖ {
	return U32ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the U32ˢ using start and end indices.
func (s U32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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
func (s U64ˢ) Clone(ϟs *gfxapi.State) U64ˢ {
	s.OnRead(ϟs)
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
func (s U64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint64 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint64, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s U64ˢ) Write(src []uint64, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint64(uint64(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U64ˢ) Copy(src U64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U64ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s U64ˢ) OnRead(ϟs *gfxapi.State) U64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s U64ˢ) OnWrite(ϟs *gfxapi.State) U64ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a U64ᵖ to the i'th element in this U64ˢ.
func (s U64ˢ) Index(i uint64, ϟs *gfxapi.State) U64ᵖ {
	return U64ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the U64ˢ using start and end indices.
func (s U64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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
func (s U8ˢ) Clone(ϟs *gfxapi.State) U8ˢ {
	s.OnRead(ϟs)
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
func (s U8ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint8 {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]uint8, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s U8ˢ) Write(src []uint8, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint8(uint8(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst U8ˢ) Copy(src U8ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U8ˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s U8ˢ) OnRead(ϟs *gfxapi.State) U8ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s U8ˢ) OnWrite(ϟs *gfxapi.State) U8ˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a U8ᵖ to the i'th element in this U8ˢ.
func (s U8ˢ) Index(i uint64, ϟs *gfxapi.State) U8ᵖ {
	return U8ᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the U8ˢ using start and end indices.
func (s U8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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

// Vec2fˢ is a slice of Vec2f.
type Vec2fˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVec2fˢ returns a Vec2fˢ backed by a new memory pool.
func MakeVec2fˢ(count uint64, ϟs *gfxapi.State) Vec2fˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Vec2fˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Vec2fˢ in a new memory pool.
func (s Vec2fˢ) Clone(ϟs *gfxapi.State) Vec2fˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Vec2fˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Vec2fˢ points to.
func (s Vec2fˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 2
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Vec2fˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Vec2fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Vec2fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Vec2fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsVec2fˢ returns s cast to a Vec2fˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsVec2fˢ(s Slice, ϟs *gfxapi.State) Vec2fˢ {
	out := Vec2fˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Vec2f elements in this Vec2fˢ.
func (s Vec2fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec2f {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Vec2f, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Vec2fˢ) Write(src []Vec2f, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Vec2fˢ) Copy(src Vec2fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec2fˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Vec2fˢ) OnRead(ϟs *gfxapi.State) Vec2fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Vec2fˢ) OnWrite(ϟs *gfxapi.State) Vec2fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Vec2fᵖ to the i'th element in this Vec2fˢ.
func (s Vec2fˢ) Index(i uint64, ϟs *gfxapi.State) Vec2fᵖ {
	return Vec2fᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Vec2fˢ using start and end indices.
func (s Vec2fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2fˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Vec2fˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Vec2fˢ slice.
func (s Vec2fˢ) String() string {
	return fmt.Sprintf("Vec2f(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Vec2iˢ is a slice of Vec2i.
type Vec2iˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVec2iˢ returns a Vec2iˢ backed by a new memory pool.
func MakeVec2iˢ(count uint64, ϟs *gfxapi.State) Vec2iˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Vec2iˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Vec2iˢ in a new memory pool.
func (s Vec2iˢ) Clone(ϟs *gfxapi.State) Vec2iˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Vec2iˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Vec2iˢ points to.
func (s Vec2iˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 2
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Vec2iˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Vec2iˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Vec2iˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Vec2iˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsVec2iˢ returns s cast to a Vec2iˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsVec2iˢ(s Slice, ϟs *gfxapi.State) Vec2iˢ {
	out := Vec2iˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Vec2i elements in this Vec2iˢ.
func (s Vec2iˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec2i {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Vec2i, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Vec2iˢ) Write(src []Vec2i, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Vec2iˢ) Copy(src Vec2iˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec2iˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Vec2iˢ) OnRead(ϟs *gfxapi.State) Vec2iˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Vec2iˢ) OnWrite(ϟs *gfxapi.State) Vec2iˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Vec2iᵖ to the i'th element in this Vec2iˢ.
func (s Vec2iˢ) Index(i uint64, ϟs *gfxapi.State) Vec2iᵖ {
	return Vec2iᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Vec2iˢ using start and end indices.
func (s Vec2iˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2iˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Vec2iˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Vec2iˢ slice.
func (s Vec2iˢ) String() string {
	return fmt.Sprintf("Vec2i(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Vec3fˢ is a slice of Vec3f.
type Vec3fˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVec3fˢ returns a Vec3fˢ backed by a new memory pool.
func MakeVec3fˢ(count uint64, ϟs *gfxapi.State) Vec3fˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Vec3fˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Vec3fˢ in a new memory pool.
func (s Vec3fˢ) Clone(ϟs *gfxapi.State) Vec3fˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Vec3fˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Vec3fˢ points to.
func (s Vec3fˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 3
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Vec3fˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Vec3fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Vec3fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Vec3fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsVec3fˢ returns s cast to a Vec3fˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsVec3fˢ(s Slice, ϟs *gfxapi.State) Vec3fˢ {
	out := Vec3fˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Vec3f elements in this Vec3fˢ.
func (s Vec3fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec3f {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Vec3f, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Vec3fˢ) Write(src []Vec3f, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Vec3fˢ) Copy(src Vec3fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec3fˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Vec3fˢ) OnRead(ϟs *gfxapi.State) Vec3fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Vec3fˢ) OnWrite(ϟs *gfxapi.State) Vec3fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Vec3fᵖ to the i'th element in this Vec3fˢ.
func (s Vec3fˢ) Index(i uint64, ϟs *gfxapi.State) Vec3fᵖ {
	return Vec3fᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Vec3fˢ using start and end indices.
func (s Vec3fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3fˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Vec3fˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Vec3fˢ slice.
func (s Vec3fˢ) String() string {
	return fmt.Sprintf("Vec3f(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Vec3iˢ is a slice of Vec3i.
type Vec3iˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVec3iˢ returns a Vec3iˢ backed by a new memory pool.
func MakeVec3iˢ(count uint64, ϟs *gfxapi.State) Vec3iˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Vec3iˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Vec3iˢ in a new memory pool.
func (s Vec3iˢ) Clone(ϟs *gfxapi.State) Vec3iˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Vec3iˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Vec3iˢ points to.
func (s Vec3iˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 3
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Vec3iˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Vec3iˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Vec3iˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Vec3iˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsVec3iˢ returns s cast to a Vec3iˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsVec3iˢ(s Slice, ϟs *gfxapi.State) Vec3iˢ {
	out := Vec3iˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Vec3i elements in this Vec3iˢ.
func (s Vec3iˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec3i {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Vec3i, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Vec3iˢ) Write(src []Vec3i, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Vec3iˢ) Copy(src Vec3iˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec3iˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Vec3iˢ) OnRead(ϟs *gfxapi.State) Vec3iˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Vec3iˢ) OnWrite(ϟs *gfxapi.State) Vec3iˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Vec3iᵖ to the i'th element in this Vec3iˢ.
func (s Vec3iˢ) Index(i uint64, ϟs *gfxapi.State) Vec3iᵖ {
	return Vec3iᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Vec3iˢ using start and end indices.
func (s Vec3iˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3iˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Vec3iˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Vec3iˢ slice.
func (s Vec3iˢ) String() string {
	return fmt.Sprintf("Vec3i(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Vec4fˢ is a slice of Vec4f.
type Vec4fˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVec4fˢ returns a Vec4fˢ backed by a new memory pool.
func MakeVec4fˢ(count uint64, ϟs *gfxapi.State) Vec4fˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Vec4fˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Vec4fˢ in a new memory pool.
func (s Vec4fˢ) Clone(ϟs *gfxapi.State) Vec4fˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Vec4fˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Vec4fˢ points to.
func (s Vec4fˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 4
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Vec4fˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Vec4fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Vec4fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Vec4fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsVec4fˢ returns s cast to a Vec4fˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsVec4fˢ(s Slice, ϟs *gfxapi.State) Vec4fˢ {
	out := Vec4fˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Vec4f elements in this Vec4fˢ.
func (s Vec4fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec4f {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Vec4f, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Vec4fˢ) Write(src []Vec4f, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Vec4fˢ) Copy(src Vec4fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec4fˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Vec4fˢ) OnRead(ϟs *gfxapi.State) Vec4fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Vec4fˢ) OnWrite(ϟs *gfxapi.State) Vec4fˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Vec4fᵖ to the i'th element in this Vec4fˢ.
func (s Vec4fˢ) Index(i uint64, ϟs *gfxapi.State) Vec4fᵖ {
	return Vec4fᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Vec4fˢ using start and end indices.
func (s Vec4fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4fˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Vec4fˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Vec4fˢ slice.
func (s Vec4fˢ) String() string {
	return fmt.Sprintf("Vec4f(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// Vec4iˢ is a slice of Vec4i.
type Vec4iˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVec4iˢ returns a Vec4iˢ backed by a new memory pool.
func MakeVec4iˢ(count uint64, ϟs *gfxapi.State) Vec4iˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return Vec4iˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the Vec4iˢ in a new memory pool.
func (s Vec4iˢ) Clone(ϟs *gfxapi.State) Vec4iˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := Vec4iˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that Vec4iˢ points to.
func (s Vec4iˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4) * 4
}

// Range returns the memory range this slice represents in the underlying pool.
func (s Vec4iˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s Vec4iˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s Vec4iˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s Vec4iˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsVec4iˢ returns s cast to a Vec4iˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsVec4iˢ(s Slice, ϟs *gfxapi.State) Vec4iˢ {
	out := Vec4iˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the Vec4i elements in this Vec4iˢ.
func (s Vec4iˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec4i {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]Vec4i, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s Vec4iˢ) Write(src []Vec4i, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst Vec4iˢ) Copy(src Vec4iˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec4iˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s Vec4iˢ) OnRead(ϟs *gfxapi.State) Vec4iˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Vec4iˢ) OnWrite(ϟs *gfxapi.State) Vec4iˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Vec4iᵖ to the i'th element in this Vec4iˢ.
func (s Vec4iˢ) Index(i uint64, ϟs *gfxapi.State) Vec4iᵖ {
	return Vec4iᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Vec4iˢ using start and end indices.
func (s Vec4iˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4iˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return Vec4iˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the Vec4iˢ slice.
func (s Vec4iˢ) String() string {
	return fmt.Sprintf("Vec4i(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
}

// VertexArrayIdˢ is a slice of VertexArrayId.
type VertexArrayIdˢ struct {
	binary.Generate
	SliceInfo
}

// MakeVertexArrayIdˢ returns a VertexArrayIdˢ backed by a new memory pool.
func MakeVertexArrayIdˢ(count uint64, ϟs *gfxapi.State) VertexArrayIdˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return VertexArrayIdˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the VertexArrayIdˢ in a new memory pool.
func (s VertexArrayIdˢ) Clone(ϟs *gfxapi.State) VertexArrayIdˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := VertexArrayIdˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that VertexArrayIdˢ points to.
func (s VertexArrayIdˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return uint64(4)
}

// Range returns the memory range this slice represents in the underlying pool.
func (s VertexArrayIdˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s VertexArrayIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s VertexArrayIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s VertexArrayIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// AsVertexArrayIdˢ returns s cast to a VertexArrayIdˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func AsVertexArrayIdˢ(s Slice, ϟs *gfxapi.State) VertexArrayIdˢ {
	out := VertexArrayIdˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the VertexArrayId elements in this VertexArrayIdˢ.
func (s VertexArrayIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []VertexArrayId {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]VertexArrayId, s.Count)
	s.OnRead(ϟs)
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
// which is the minimum of s.Count and len(src).
func (s VertexArrayIdˢ) Write(src []VertexArrayId, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Uint32(uint32(src[i])); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst VertexArrayIdˢ) Copy(src VertexArrayIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s VertexArrayIdˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s VertexArrayIdˢ) OnRead(ϟs *gfxapi.State) VertexArrayIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s VertexArrayIdˢ) OnWrite(ϟs *gfxapi.State) VertexArrayIdˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a VertexArrayIdᵖ to the i'th element in this VertexArrayIdˢ.
func (s VertexArrayIdˢ) Index(i uint64, ϟs *gfxapi.State) VertexArrayIdᵖ {
	return VertexArrayIdᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the VertexArrayIdˢ using start and end indices.
func (s VertexArrayIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return VertexArrayIdˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the VertexArrayIdˢ slice.
func (s VertexArrayIdˢ) String() string {
	return fmt.Sprintf("VertexArrayId(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
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
func (s Voidˢ) Clone(ϟs *gfxapi.State) Voidˢ {
	s.OnRead(ϟs)
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
func (s Voidˢ) OnRead(ϟs *gfxapi.State) Voidˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s Voidˢ) OnWrite(ϟs *gfxapi.State) Voidˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a Voidᵖ to the i'th element in this Voidˢ.
func (s Voidˢ) Index(i uint64, ϟs *gfxapi.State) Voidᵖ {
	return Voidᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the Voidˢ using start and end indices.
func (s Voidˢ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
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

// __GLsyncˢ is a slice of __GLsync.
type __GLsyncˢ struct {
	binary.Generate
	SliceInfo
}

// Make__GLsyncˢ returns a __GLsyncˢ backed by a new memory pool.
func Make__GLsyncˢ(count uint64, ϟs *gfxapi.State) __GLsyncˢ {
	id := ϟs.NextPoolID
	ϟs.Memory[id] = &memory.Pool{}
	ϟs.NextPoolID++
	return __GLsyncˢ{SliceInfo: SliceInfo{Count: count, Root: memory.Pointer{Pool: id}}}
}

// Clone returns a copy of the __GLsyncˢ in a new memory pool.
func (s __GLsyncˢ) Clone(ϟs *gfxapi.State) __GLsyncˢ {
	s.OnRead(ϟs)
	pool := &memory.Pool{}
	pool.Write(0, ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)))
	id := ϟs.NextPoolID
	ϟs.Memory[id] = pool
	ϟs.NextPoolID++
	dst := __GLsyncˢ{SliceInfo: SliceInfo{Count: s.Count, Root: memory.Pointer{Pool: id}}}
	return dst
}

// ElementSize returns the size in bytes of an element that __GLsyncˢ points to.
func (s __GLsyncˢ) ElementSize(ϟs *gfxapi.State) uint64 {
	return func() uint64 { panic("Sizeof class is not yet implemented") }()
}

// Range returns the memory range this slice represents in the underlying pool.
func (s __GLsyncˢ) Range(ϟs *gfxapi.State) memory.Range {
	return memory.Range{Base: s.Base, Size: s.Count * s.ElementSize(ϟs)}
}

// ResourceID returns an identifier to a resource representing the data of
// this slice.
func (s __GLsyncˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID {
	id, err := ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)).ResourceID(ϟd, ϟl)
	if err != nil {
		panic(err)
	}
	return id
}

// Decoder returns a memory decoder for the slice.
func (s __GLsyncˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder {
	return ϟs.MemoryDecoder(ϟs.Memory[s.Root.Pool].Slice(s.Range(ϟs)), ϟd, ϟl)
}

// Encoder returns a memory encoder for the slice.
func (s __GLsyncˢ) Encoder(ϟs *gfxapi.State) binary.Encoder {
	return ϟs.MemoryEncoder(ϟs.Memory[s.Root.Pool], s.Range(ϟs))
}

// As__GLsyncˢ returns s cast to a __GLsyncˢ.
// The returned slice length will be calculated so that the returned slice is
// no longer (in bytes) than s.
func As__GLsyncˢ(s Slice, ϟs *gfxapi.State) __GLsyncˢ {
	out := __GLsyncˢ{SliceInfo: s.Info()}
	out.Count = (out.Count * s.ElementSize(ϟs)) / out.ElementSize(ϟs)
	return out
}

// Read reads and returns all the __GLsync elements in this __GLsyncˢ.
func (s __GLsyncˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []__GLsync {
	d, res := s.Decoder(ϟs, ϟd, ϟl), make([]__GLsync, s.Count)
	s.OnRead(ϟs)
	for i := range res {
		if err := d.Value(&res[i]); err != nil {
			panic(err)
		}
	}
	return res
}

// Write copies elements from src to this slice. The number of elements copied is returned
// which is the minimum of s.Count and len(src).
func (s __GLsyncˢ) Write(src []__GLsync, ϟs *gfxapi.State) uint64 {
	count := min(s.Count, uint64(len(src)))
	s = s.Slice(0, count, ϟs)
	e := s.Encoder(ϟs)
	for i := uint64(0); i < count; i++ {
		if err := e.Value(&src[i]); err != nil {
			panic(err)
		}
	}
	s.OnWrite(ϟs)
	return count
}

// Copy copies elements from src to this slice.
// The number of elements copied is the minimum of dst.Count and src.Count.
// The slices of this and dst to the copied elements is returned.
func (dst __GLsyncˢ) Copy(src __GLsyncˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s __GLsyncˢ) {
	count := min(dst.Count, src.Count)
	dst, src = dst.Slice(0, count, ϟs), src.Slice(0, count, ϟs)
	src.OnRead(ϟs)
	ϟs.Memory[dst.Root.Pool].Write(dst.Base, ϟs.Memory[src.Root.Pool].Slice(src.Range(ϟs)))
	dst.OnWrite(ϟs)
	return dst, src
}

// OnRead calls the backing pool's OnRead callback. s is returned so calls can be chained.
func (s __GLsyncˢ) OnRead(ϟs *gfxapi.State) __GLsyncˢ {
	if f := ϟs.Memory[s.Root.Pool].OnRead; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be chained.
func (s __GLsyncˢ) OnWrite(ϟs *gfxapi.State) __GLsyncˢ {
	if f := ϟs.Memory[s.Root.Pool].OnWrite; f != nil {
		f(s.Range(ϟs))
	}
	return s
}

// Index returns a __GLsyncᵖ to the i'th element in this __GLsyncˢ.
func (s __GLsyncˢ) Index(i uint64, ϟs *gfxapi.State) __GLsyncᵖ {
	return __GLsyncᵖ{Pointer: memory.Pointer{Address: s.Base + i*s.ElementSize(ϟs), Pool: s.Root.Pool}}
}

// Slice returns a sub-slice from the __GLsyncˢ using start and end indices.
func (s __GLsyncˢ) Slice(start, end uint64, ϟs *gfxapi.State) __GLsyncˢ {
	if start >= end {
		panic(fmt.Errorf("%v.Slice(%d, %d) - start must be less than end", s, start, end))
	}
	if end > s.Count {
		panic(fmt.Errorf("%v.Slice(%d, %d) - out of bounds", s, start, end))
	}
	return __GLsyncˢ{SliceInfo: SliceInfo{Root: s.Root, Base: s.Base + start*s.ElementSize(ϟs), Count: end - start}}
}

// String returns a string description of the __GLsyncˢ slice.
func (s __GLsyncˢ) String() string {
	return fmt.Sprintf("__GLsync(%v@%v)[%d]", s.Base, s.Root.Pool, s.Count)
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

type GLenumːBufferIdᵐ map[GLenum]BufferId

func (m GLenumːBufferIdᵐ) Get(key GLenum) BufferId {
	return m[key]
}
func (m GLenumːBufferIdᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːBufferIdᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːBufferIdᵐ) Range() []BufferId {
	values := make([]BufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːFramebufferAttachmentInfoᵐ map[GLenum]FramebufferAttachmentInfo

func (m GLenumːFramebufferAttachmentInfoᵐ) Get(key GLenum) FramebufferAttachmentInfo {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m GLenumːFramebufferAttachmentInfoᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːFramebufferAttachmentInfoᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːFramebufferAttachmentInfoᵐ) Range() []FramebufferAttachmentInfo {
	values := make([]FramebufferAttachmentInfo, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːFramebufferIdᵐ map[GLenum]FramebufferId

func (m GLenumːFramebufferIdᵐ) Get(key GLenum) FramebufferId {
	return m[key]
}
func (m GLenumːFramebufferIdᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːFramebufferIdᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːFramebufferIdᵐ) Range() []FramebufferId {
	values := make([]FramebufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːGLenumːTextureIdᵐᵐ map[GLenum]GLenumːTextureIdᵐ

func (m GLenumːGLenumːTextureIdᵐᵐ) Get(key GLenum) GLenumːTextureIdᵐ {
	v, ok := m[key]
	if !ok {
		v = make(GLenumːTextureIdᵐ)
	}
	return v
}
func (m GLenumːGLenumːTextureIdᵐᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːGLenumːTextureIdᵐᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːGLenumːTextureIdᵐᵐ) Range() []GLenumːTextureIdᵐ {
	values := make([]GLenumːTextureIdᵐ, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːImageᵐ map[GLenum]Image

func (m GLenumːImageᵐ) Get(key GLenum) Image {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m GLenumːImageᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːImageᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːImageᵐ) Range() []Image {
	values := make([]Image, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːRenderbufferIdᵐ map[GLenum]RenderbufferId

func (m GLenumːRenderbufferIdᵐ) Get(key GLenum) RenderbufferId {
	return m[key]
}
func (m GLenumːRenderbufferIdᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːRenderbufferIdᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːRenderbufferIdᵐ) Range() []RenderbufferId {
	values := make([]RenderbufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːShaderIdᵐ map[GLenum]ShaderId

func (m GLenumːShaderIdᵐ) Get(key GLenum) ShaderId {
	return m[key]
}
func (m GLenumːShaderIdᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːShaderIdᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːShaderIdᵐ) Range() []ShaderId {
	values := make([]ShaderId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːTextureIdᵐ map[GLenum]TextureId

func (m GLenumːTextureIdᵐ) Get(key GLenum) TextureId {
	return m[key]
}
func (m GLenumːTextureIdᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːTextureIdᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːTextureIdᵐ) Range() []TextureId {
	values := make([]TextureId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːboolᵐ map[GLenum]bool

func (m GLenumːboolᵐ) Get(key GLenum) bool {
	return m[key]
}
func (m GLenumːboolᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːboolᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːboolᵐ) Range() []bool {
	values := make([]bool, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːs32ᵐ map[GLenum]int32

func (m GLenumːs32ᵐ) Get(key GLenum) int32 {
	return m[key]
}
func (m GLenumːs32ᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːs32ᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːs32ᵐ) Range() []int32 {
	values := make([]int32, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type GLenumːu32ᵐ map[GLenum]uint32

func (m GLenumːu32ᵐ) Get(key GLenum) uint32 {
	return m[key]
}
func (m GLenumːu32ᵐ) Contains(key GLenum) bool {
	_, ok := m[key]
	return ok
}
func (m GLenumːu32ᵐ) Delete(key GLenum) {
	delete(m, key)
}
func (m GLenumːu32ᵐ) Range() []uint32 {
	values := make([]uint32, 0, len(m))
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
func (c *EglInitialize) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *EglCreateContext) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *EglMakeCurrent) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *EglSwapBuffers) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *EglQuerySurface) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlXCreateContext) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlXCreateNewContext) API() gfxapi.ID                   { return api{}.ID() }
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
	Result       Bool
}

func (a *GlXMakeContextCurrent) String() string {
	return fmt.Sprintf("glXMakeContextCurrent(display: %v, draw: %v, read: %v, ctx: %v) → %v", a.Display, a.Draw, a.Read, a.Ctx, a.Result)
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
func (c *GlXMakeContextCurrent) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlXMakeContextCurrent) Flags() atom.Flags                { return 0 }
func (a *GlXMakeContextCurrent) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlXMakeCurrent
////////////////////////////////////////////////////////////////////////////////
type GlXMakeCurrent struct {
	binary.Generate
	observations atom.Observations
	Display      Voidᵖ
	Drawable     GLXDrawable
	Ctx          GLXContext
	Result       Bool
}

func (a *GlXMakeCurrent) String() string {
	return fmt.Sprintf("glXMakeCurrent(display: %v, drawable: %v, ctx: %v) → %v", a.Display, a.Drawable, a.Ctx, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlXMakeCurrent pointer is returned so that calls can be chained.
func (a *GlXMakeCurrent) AddRead(rng memory.Range, id binary.ID) *GlXMakeCurrent {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlXMakeCurrent pointer is returned so that calls can be chained.
func (a *GlXMakeCurrent) AddWrite(rng memory.Range, id binary.ID) *GlXMakeCurrent {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlXMakeCurrent) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlXMakeCurrent) Flags() atom.Flags                { return 0 }
func (a *GlXMakeCurrent) Observations() *atom.Observations { return &a.observations }

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
func (c *GlXSwapBuffers) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlXSwapBuffers) Flags() atom.Flags                { return 0 | atom.EndOfFrame }
func (a *GlXSwapBuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlXQueryDrawable
////////////////////////////////////////////////////////////////////////////////
type GlXQueryDrawable struct {
	binary.Generate
	observations atom.Observations
	Display      Voidᵖ
	Draw         GLXDrawable
	Attribute    int64
	Value        Intᵖ
	Result       int64
}

func (a *GlXQueryDrawable) String() string {
	return fmt.Sprintf("glXQueryDrawable(display: %v, draw: %v, attribute: %v, value: %v) → %v", a.Display, a.Draw, a.Attribute, a.Value, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlXQueryDrawable pointer is returned so that calls can be chained.
func (a *GlXQueryDrawable) AddRead(rng memory.Range, id binary.ID) *GlXQueryDrawable {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlXQueryDrawable pointer is returned so that calls can be chained.
func (a *GlXQueryDrawable) AddWrite(rng memory.Range, id binary.ID) *GlXQueryDrawable {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlXQueryDrawable) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlXQueryDrawable) Flags() atom.Flags                { return 0 }
func (a *GlXQueryDrawable) Observations() *atom.Observations { return &a.observations }

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
func (c *WglCreateContext) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *WglCreateContextAttribsARB) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *WglMakeCurrent) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *WglSwapBuffers) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *CGLCreateContext) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *CGLSetCurrentContext) API() gfxapi.ID                   { return api{}.ID() }
func (c *CGLSetCurrentContext) Flags() atom.Flags                { return 0 }
func (a *CGLSetCurrentContext) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CGLGetSurface
////////////////////////////////////////////////////////////////////////////////
type CGLGetSurface struct {
	binary.Generate
	observations atom.Observations
	Ctx          CGLContextObj
	Cid          CGSConnectionIDᵖ
	Wid          CGSWindowIDᵖ
	Sid          CGSSurfaceIDᵖ
	Result       int64
}

func (a *CGLGetSurface) String() string {
	return fmt.Sprintf("CGLGetSurface(ctx: %v, cid: %v, wid: %v, sid: %v) → %v", a.Ctx, a.Cid, a.Wid, a.Sid, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CGLGetSurface pointer is returned so that calls can be chained.
func (a *CGLGetSurface) AddRead(rng memory.Range, id binary.ID) *CGLGetSurface {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CGLGetSurface pointer is returned so that calls can be chained.
func (a *CGLGetSurface) AddWrite(rng memory.Range, id binary.ID) *CGLGetSurface {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CGLGetSurface) API() gfxapi.ID                   { return api{}.ID() }
func (c *CGLGetSurface) Flags() atom.Flags                { return 0 }
func (a *CGLGetSurface) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CGSGetSurfaceBounds
////////////////////////////////////////////////////////////////////////////////
type CGSGetSurfaceBounds struct {
	binary.Generate
	observations atom.Observations
	Cid          CGSConnectionID
	Wid          CGSWindowID
	Sid          CGSSurfaceID
	Bounds       F64ᵖ
	Result       int64
}

func (a *CGSGetSurfaceBounds) String() string {
	return fmt.Sprintf("CGSGetSurfaceBounds(cid: %v, wid: %v, sid: %v, bounds: %v) → %v", a.Cid, a.Wid, a.Sid, a.Bounds, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CGSGetSurfaceBounds pointer is returned so that calls can be chained.
func (a *CGSGetSurfaceBounds) AddRead(rng memory.Range, id binary.ID) *CGSGetSurfaceBounds {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CGSGetSurfaceBounds pointer is returned so that calls can be chained.
func (a *CGSGetSurfaceBounds) AddWrite(rng memory.Range, id binary.ID) *CGSGetSurfaceBounds {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CGSGetSurfaceBounds) API() gfxapi.ID                   { return api{}.ID() }
func (c *CGSGetSurfaceBounds) Flags() atom.Flags                { return 0 }
func (a *CGSGetSurfaceBounds) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// CGLFlushDrawable
////////////////////////////////////////////////////////////////////////////////
type CGLFlushDrawable struct {
	binary.Generate
	observations atom.Observations
	Ctx          CGLContextObj
	Result       CGLError
}

func (a *CGLFlushDrawable) String() string {
	return fmt.Sprintf("CGLFlushDrawable(ctx: %v) → %v", a.Ctx, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The CGLFlushDrawable pointer is returned so that calls can be chained.
func (a *CGLFlushDrawable) AddRead(rng memory.Range, id binary.ID) *CGLFlushDrawable {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The CGLFlushDrawable pointer is returned so that calls can be chained.
func (a *CGLFlushDrawable) AddWrite(rng memory.Range, id binary.ID) *CGLFlushDrawable {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *CGLFlushDrawable) API() gfxapi.ID                   { return api{}.ID() }
func (c *CGLFlushDrawable) Flags() atom.Flags                { return 0 | atom.EndOfFrame }
func (a *CGLFlushDrawable) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEnableClientState
////////////////////////////////////////////////////////////////////////////////
type GlEnableClientState struct {
	binary.Generate
	observations atom.Observations
	Type         GLenum
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
func (c *GlEnableClientState) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlEnableClientState) Flags() atom.Flags                { return 0 }
func (a *GlEnableClientState) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDisableClientState
////////////////////////////////////////////////////////////////////////////////
type GlDisableClientState struct {
	binary.Generate
	observations atom.Observations
	Type         GLenum
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
func (c *GlDisableClientState) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGetProgramBinaryOES) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlProgramBinaryOES) API() gfxapi.ID                   { return api{}.ID() }
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
	PreserveMask GLbitfield
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
func (c *GlStartTilingQCOM) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlStartTilingQCOM) Flags() atom.Flags                { return 0 }
func (a *GlStartTilingQCOM) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEndTilingQCOM
////////////////////////////////////////////////////////////////////////////////
type GlEndTilingQCOM struct {
	binary.Generate
	observations atom.Observations
	PreserveMask GLbitfield
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
func (c *GlEndTilingQCOM) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlEndTilingQCOM) Flags() atom.Flags                { return 0 }
func (a *GlEndTilingQCOM) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDiscardFramebufferEXT
////////////////////////////////////////////////////////////////////////////////
type GlDiscardFramebufferEXT struct {
	binary.Generate
	observations   atom.Observations
	Target         GLenum
	NumAttachments int32
	Attachments    GLenumᵖ
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
func (c *GlDiscardFramebufferEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlDiscardFramebufferEXT) Flags() atom.Flags                { return 0 }
func (a *GlDiscardFramebufferEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlInsertEventMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlInsertEventMarkerEXT struct {
	binary.Generate
	observations atom.Observations
	Length       int32
	Marker       Charᵖ
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
func (c *GlInsertEventMarkerEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlInsertEventMarkerEXT) Flags() atom.Flags                { return 0 }
func (a *GlInsertEventMarkerEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlPushGroupMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlPushGroupMarkerEXT struct {
	binary.Generate
	observations atom.Observations
	Length       int32
	Marker       Charᵖ
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
func (c *GlPushGroupMarkerEXT) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlPopGroupMarkerEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlPopGroupMarkerEXT) Flags() atom.Flags                { return 0 }
func (a *GlPopGroupMarkerEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage1DEXT struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Levels       int32
	Format       GLenum
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
func (c *GlTexStorage1DEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTexStorage1DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTexStorage1DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage2DEXT struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Levels       int32
	Format       GLenum
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
func (c *GlTexStorage2DEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTexStorage2DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTexStorage2DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage3DEXT struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Levels       int32
	Format       GLenum
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
func (c *GlTexStorage3DEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTexStorage3DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTexStorage3DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage1DEXT struct {
	binary.Generate
	observations atom.Observations
	Texture      TextureId
	Target       GLenum
	Levels       int32
	Format       GLenum
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
func (c *GlTextureStorage1DEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTextureStorage1DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTextureStorage1DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage2DEXT struct {
	binary.Generate
	observations atom.Observations
	Texture      TextureId
	Target       GLenum
	Levels       int32
	Format       GLenum
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
func (c *GlTextureStorage2DEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTextureStorage2DEXT) Flags() atom.Flags                { return 0 }
func (a *GlTextureStorage2DEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage3DEXT struct {
	binary.Generate
	observations atom.Observations
	Texture      TextureId
	Target       GLenum
	Levels       int32
	Format       GLenum
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
func (c *GlTextureStorage3DEXT) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGenVertexArraysOES) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlBindVertexArrayOES) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindVertexArrayOES) Flags() atom.Flags                { return 0 }
func (a *GlBindVertexArrayOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteVertexArraysOES
////////////////////////////////////////////////////////////////////////////////
type GlDeleteVertexArraysOES struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Arrays       VertexArrayIdᶜᵖ
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
func (c *GlDeleteVertexArraysOES) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsVertexArrayOES) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlIsVertexArrayOES) Flags() atom.Flags                { return 0 }
func (a *GlIsVertexArrayOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetTexture2DOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetTexture2DOES struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlEGLImageTargetTexture2DOES) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlEGLImageTargetTexture2DOES) Flags() atom.Flags                { return 0 }
func (a *GlEGLImageTargetTexture2DOES) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetRenderbufferStorageOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetRenderbufferStorageOES struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlEGLImageTargetRenderbufferStorageOES) API() gfxapi.ID    { return api{}.ID() }
func (c *GlEGLImageTargetRenderbufferStorageOES) Flags() atom.Flags { return 0 }
func (a *GlEGLImageTargetRenderbufferStorageOES) Observations() *atom.Observations {
	return &a.observations
}

////////////////////////////////////////////////////////////////////////////////
// GlGetGraphicsResetStatusEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetGraphicsResetStatusEXT struct {
	binary.Generate
	observations atom.Observations
	Result       GLenum
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
func (c *GlGetGraphicsResetStatusEXT) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlBindAttribLocation) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindAttribLocation) Flags() atom.Flags                { return 0 }
func (a *GlBindAttribLocation) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendFunc
////////////////////////////////////////////////////////////////////////////////
type GlBlendFunc struct {
	binary.Generate
	observations atom.Observations
	SrcFactor    GLenum
	DstFactor    GLenum
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
func (c *GlBlendFunc) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBlendFunc) Flags() atom.Flags                { return 0 }
func (a *GlBlendFunc) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendFuncSeparate struct {
	binary.Generate
	observations   atom.Observations
	SrcFactorRgb   GLenum
	DstFactorRgb   GLenum
	SrcFactorAlpha GLenum
	DstFactorAlpha GLenum
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
func (c *GlBlendFuncSeparate) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBlendFuncSeparate) Flags() atom.Flags                { return 0 }
func (a *GlBlendFuncSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquation
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquation struct {
	binary.Generate
	observations atom.Observations
	Equation     GLenum
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
func (c *GlBlendEquation) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBlendEquation) Flags() atom.Flags                { return 0 }
func (a *GlBlendEquation) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquationSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquationSeparate struct {
	binary.Generate
	observations atom.Observations
	Rgb          GLenum
	Alpha        GLenum
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
func (c *GlBlendEquationSeparate) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlBlendColor) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlEnableVertexAttribArray) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlDisableVertexAttribArray) API() gfxapi.ID                   { return api{}.ID() }
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
	Type         GLenum
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
func (c *GlVertexAttribPointer) API() gfxapi.ID                   { return api{}.ID() }
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
	Type               GLenumᵖ
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
func (c *GlGetActiveAttrib) API() gfxapi.ID                   { return api{}.ID() }
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
	VectorCount        S32ᵖ
	Type               GLenumᵖ
	Name               Charᵖ
}

func (a *GlGetActiveUniform) String() string {
	return fmt.Sprintf("glGetActiveUniform(program: %v, location: %v, buffer_size: %v, buffer_bytes_written: %v, vector_count: %v, type: %v, name: %v)", a.Program, a.Location, a.BufferSize, a.BufferBytesWritten, a.VectorCount, a.Type, a.Name)
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
func (c *GlGetActiveUniform) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetActiveUniform) Flags() atom.Flags                { return 0 }
func (a *GlGetActiveUniform) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetError
////////////////////////////////////////////////////////////////////////////////
type GlGetError struct {
	binary.Generate
	observations atom.Observations
	Result       GLenum
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
func (c *GlGetError) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetError) Flags() atom.Flags                { return 0 }
func (a *GlGetError) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramiv
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramiv struct {
	binary.Generate
	observations atom.Observations
	Program      ProgramId
	Parameter    GLenum
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
func (c *GlGetProgramiv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetProgramiv) Flags() atom.Flags                { return 0 }
func (a *GlGetProgramiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderiv
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderiv struct {
	binary.Generate
	observations atom.Observations
	Shader       ShaderId
	Parameter    GLenum
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
func (c *GlGetShaderiv) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGetUniformLocation) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGetAttribLocation) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetAttribLocation) Flags() atom.Flags                { return 0 }
func (a *GlGetAttribLocation) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlPixelStorei
////////////////////////////////////////////////////////////////////////////////
type GlPixelStorei struct {
	binary.Generate
	observations atom.Observations
	Parameter    GLenum
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
func (c *GlPixelStorei) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlPixelStorei) Flags() atom.Flags                { return 0 }
func (a *GlPixelStorei) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexParameteri
////////////////////////////////////////////////////////////////////////////////
type GlTexParameteri struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
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
func (c *GlTexParameteri) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTexParameteri) Flags() atom.Flags                { return 0 }
func (a *GlTexParameteri) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexParameterf
////////////////////////////////////////////////////////////////////////////////
type GlTexParameterf struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
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
func (c *GlTexParameterf) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTexParameterf) Flags() atom.Flags                { return 0 }
func (a *GlTexParameterf) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameteriv struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
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
func (c *GlGetTexParameteriv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetTexParameteriv) Flags() atom.Flags                { return 0 }
func (a *GlGetTexParameteriv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameterfv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameterfv struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
	Values       F32ᶜᵖ
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
func (c *GlGetTexParameterfv) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform1i) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform2i) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform3i) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform4i) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       S32ᵖ
}

func (a *GlUniform1iv) String() string {
	return fmt.Sprintf("glUniform1iv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform1iv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       S32ᵖ
}

func (a *GlUniform2iv) String() string {
	return fmt.Sprintf("glUniform2iv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform2iv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       S32ᵖ
}

func (a *GlUniform3iv) String() string {
	return fmt.Sprintf("glUniform3iv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform3iv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       S32ᵖ
}

func (a *GlUniform4iv) String() string {
	return fmt.Sprintf("glUniform4iv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform4iv) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform1f) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform2f) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform3f) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUniform4f) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
}

func (a *GlUniform1fv) String() string {
	return fmt.Sprintf("glUniform1fv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform1fv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
}

func (a *GlUniform2fv) String() string {
	return fmt.Sprintf("glUniform2fv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform2fv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
}

func (a *GlUniform3fv) String() string {
	return fmt.Sprintf("glUniform3fv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform3fv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
}

func (a *GlUniform4fv) String() string {
	return fmt.Sprintf("glUniform4fv(location: %v, count: %v, values: %v)", a.Location, a.Count, a.Values)
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
func (c *GlUniform4fv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
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
func (c *GlUniformMatrix2fv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
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
func (c *GlUniformMatrix3fv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
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
func (c *GlUniformMatrix4fv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       F32ᶜᵖ
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
func (c *GlGetUniformfv) API() gfxapi.ID                   { return api{}.ID() }
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
	Values       S32ᶜᵖ
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
func (c *GlGetUniformiv) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlVertexAttrib1f) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlVertexAttrib2f) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlVertexAttrib3f) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlVertexAttrib4f) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlVertexAttrib4f) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib4f) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib1fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib1fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᶜᵖ
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
func (c *GlVertexAttrib1fv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlVertexAttrib1fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib1fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib2fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib2fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᶜᵖ
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
func (c *GlVertexAttrib2fv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlVertexAttrib2fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib2fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib3fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib3fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᶜᵖ
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
func (c *GlVertexAttrib3fv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlVertexAttrib3fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib3fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib4fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib4fv struct {
	binary.Generate
	observations atom.Observations
	Location     AttributeLocation
	Value        F32ᶜᵖ
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
func (c *GlVertexAttrib4fv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlVertexAttrib4fv) Flags() atom.Flags                { return 0 }
func (a *GlVertexAttrib4fv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderPrecisionFormat
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderPrecisionFormat struct {
	binary.Generate
	observations  atom.Observations
	ShaderType    GLenum
	PrecisionType GLenum
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
func (c *GlGetShaderPrecisionFormat) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlDepthMask) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlDepthMask) Flags() atom.Flags                { return 0 }
func (a *GlDepthMask) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDepthFunc
////////////////////////////////////////////////////////////////////////////////
type GlDepthFunc struct {
	binary.Generate
	observations atom.Observations
	Function     GLenum
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
func (c *GlDepthFunc) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlDepthRangef) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlColorMask) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlStencilMask) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlStencilMask) Flags() atom.Flags                { return 0 }
func (a *GlStencilMask) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStencilMaskSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilMaskSeparate struct {
	binary.Generate
	observations atom.Observations
	Face         GLenum
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
func (c *GlStencilMaskSeparate) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlStencilMaskSeparate) Flags() atom.Flags                { return 0 }
func (a *GlStencilMaskSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStencilFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilFuncSeparate struct {
	binary.Generate
	observations   atom.Observations
	Face           GLenum
	Function       GLenum
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
func (c *GlStencilFuncSeparate) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlStencilFuncSeparate) Flags() atom.Flags                { return 0 }
func (a *GlStencilFuncSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlStencilOpSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilOpSeparate struct {
	binary.Generate
	observations         atom.Observations
	Face                 GLenum
	StencilFail          GLenum
	StencilPassDepthFail GLenum
	StencilPassDepthPass GLenum
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
func (c *GlStencilOpSeparate) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlStencilOpSeparate) Flags() atom.Flags                { return 0 }
func (a *GlStencilOpSeparate) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFrontFace
////////////////////////////////////////////////////////////////////////////////
type GlFrontFace struct {
	binary.Generate
	observations atom.Observations
	Orientation  GLenum
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
func (c *GlFrontFace) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlViewport) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlScissor) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlScissor) Flags() atom.Flags                { return 0 }
func (a *GlScissor) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlActiveTexture
////////////////////////////////////////////////////////////////////////////////
type GlActiveTexture struct {
	binary.Generate
	observations atom.Observations
	Unit         GLenum
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
func (c *GlActiveTexture) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGenTextures) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGenTextures) Flags() atom.Flags                { return 0 }
func (a *GlGenTextures) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteTextures
////////////////////////////////////////////////////////////////////////////////
type GlDeleteTextures struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Textures     TextureIdᶜᵖ
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
func (c *GlDeleteTextures) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsTexture) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlIsTexture) Flags() atom.Flags                { return 0 }
func (a *GlIsTexture) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindTexture
////////////////////////////////////////////////////////////////////////////////
type GlBindTexture struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlBindTexture) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindTexture) Flags() atom.Flags                { return 0 }
func (a *GlBindTexture) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexImage2D struct {
	binary.Generate
	observations   atom.Observations
	Target         GLenum
	Level          int32
	InternalFormat GLenum
	Width          int32
	Height         int32
	Border         int32
	Format         GLenum
	Type           GLenum
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
func (c *GlTexImage2D) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTexImage2D) Flags() atom.Flags                { return 0 }
func (a *GlTexImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexSubImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Level        int32
	Xoffset      int32
	Yoffset      int32
	Width        int32
	Height       int32
	Format       GLenum
	Type         GLenum
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
func (c *GlTexSubImage2D) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlTexSubImage2D) Flags() atom.Flags                { return 0 }
func (a *GlTexSubImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Level        int32
	Format       GLenum
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
func (c *GlCopyTexImage2D) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlCopyTexImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCopyTexImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexSubImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlCopyTexSubImage2D) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlCopyTexSubImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCopyTexSubImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Level        int32
	Format       GLenum
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
func (c *GlCompressedTexImage2D) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlCompressedTexImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCompressedTexImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexSubImage2D struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Level        int32
	Xoffset      int32
	Yoffset      int32
	Width        int32
	Height       int32
	Format       GLenum
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
func (c *GlCompressedTexSubImage2D) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlCompressedTexSubImage2D) Flags() atom.Flags                { return 0 }
func (a *GlCompressedTexSubImage2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenerateMipmap
////////////////////////////////////////////////////////////////////////////////
type GlGenerateMipmap struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlGenerateMipmap) API() gfxapi.ID                   { return api{}.ID() }
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
	Format       GLenum
	Type         GLenum
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
func (c *GlReadPixels) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGenFramebuffers) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGenFramebuffers) Flags() atom.Flags                { return 0 }
func (a *GlGenFramebuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindFramebuffer struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlBindFramebuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindFramebuffer) Flags() atom.Flags                { return 0 }
func (a *GlBindFramebuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCheckFramebufferStatus
////////////////////////////////////////////////////////////////////////////////
type GlCheckFramebufferStatus struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Result       GLenum
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
func (c *GlCheckFramebufferStatus) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlCheckFramebufferStatus) Flags() atom.Flags                { return 0 }
func (a *GlCheckFramebufferStatus) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteFramebuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteFramebuffers struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Framebuffers FramebufferIdᶜᵖ
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
func (c *GlDeleteFramebuffers) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsFramebuffer) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGenRenderbuffers) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGenRenderbuffers) Flags() atom.Flags                { return 0 }
func (a *GlGenRenderbuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindRenderbuffer struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlBindRenderbuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindRenderbuffer) Flags() atom.Flags                { return 0 }
func (a *GlBindRenderbuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorage
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorage struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Format       GLenum
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
func (c *GlRenderbufferStorage) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlRenderbufferStorage) Flags() atom.Flags                { return 0 }
func (a *GlRenderbufferStorage) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteRenderbuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteRenderbuffers struct {
	binary.Generate
	observations  atom.Observations
	Count         int32
	Renderbuffers RenderbufferIdᶜᵖ
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
func (c *GlDeleteRenderbuffers) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsRenderbuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlIsRenderbuffer) Flags() atom.Flags                { return 0 }
func (a *GlIsRenderbuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetRenderbufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetRenderbufferParameteriv struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
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
func (c *GlGetRenderbufferParameteriv) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGenBuffers) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGenBuffers) Flags() atom.Flags                { return 0 }
func (a *GlGenBuffers) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindBuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindBuffer struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlBindBuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindBuffer) Flags() atom.Flags                { return 0 }
func (a *GlBindBuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBufferData
////////////////////////////////////////////////////////////////////////////////
type GlBufferData struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Size         int32
	Data         BufferDataPointer
	Usage        GLenum
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
func (c *GlBufferData) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBufferData) Flags() atom.Flags                { return 0 }
func (a *GlBufferData) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBufferSubData
////////////////////////////////////////////////////////////////////////////////
type GlBufferSubData struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Offset       int32
	Size         int32
	Data         BufferDataPointer
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
func (c *GlBufferSubData) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBufferSubData) Flags() atom.Flags                { return 0 }
func (a *GlBufferSubData) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteBuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteBuffers struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Buffers      BufferIdᶜᵖ
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
func (c *GlDeleteBuffers) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsBuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlIsBuffer) Flags() atom.Flags                { return 0 }
func (a *GlIsBuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetBufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetBufferParameteriv struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
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
func (c *GlGetBufferParameteriv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetBufferParameteriv) Flags() atom.Flags                { return 0 }
func (a *GlGetBufferParameteriv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCreateShader
////////////////////////////////////////////////////////////////////////////////
type GlCreateShader struct {
	binary.Generate
	observations atom.Observations
	Type         GLenum
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
func (c *GlCreateShader) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlDeleteShader) API() gfxapi.ID                   { return api{}.ID() }
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
	Source       Charᶜᵖᶜᵖ
	Length       S32ᶜᵖ
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
func (c *GlShaderSource) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlShaderSource) Flags() atom.Flags                { return 0 }
func (a *GlShaderSource) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlShaderBinary
////////////////////////////////////////////////////////////////////////////////
type GlShaderBinary struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Shaders      ShaderIdᶜᵖ
	BinaryFormat uint32
	Binary       Voidᶜᵖ
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
func (c *GlShaderBinary) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGetShaderInfoLog) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGetShaderSource) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlReleaseShaderCompiler) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlCompileShader) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsShader) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlCreateProgram) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlDeleteProgram) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlAttachShader) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlDetachShader) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGetAttachedShaders) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlLinkProgram) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGetProgramInfoLog) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlUseProgram) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsProgram) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlValidateProgram) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlClearColor) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlClearDepthf) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlClearStencil) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlClearStencil) Flags() atom.Flags                { return 0 }
func (a *GlClearStencil) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlClear
////////////////////////////////////////////////////////////////////////////////
type GlClear struct {
	binary.Generate
	observations atom.Observations
	Mask         GLbitfield
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
func (c *GlClear) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlClear) Flags() atom.Flags                { return 0 }
func (a *GlClear) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlCullFace
////////////////////////////////////////////////////////////////////////////////
type GlCullFace struct {
	binary.Generate
	observations atom.Observations
	Mode         GLenum
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
func (c *GlCullFace) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlPolygonOffset) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlLineWidth) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlSampleCoverage) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlSampleCoverage) Flags() atom.Flags                { return 0 }
func (a *GlSampleCoverage) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlHint
////////////////////////////////////////////////////////////////////////////////
type GlHint struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Mode         GLenum
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
func (c *GlHint) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlHint) Flags() atom.Flags                { return 0 }
func (a *GlHint) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferRenderbuffer struct {
	binary.Generate
	observations          atom.Observations
	FramebufferTarget     GLenum
	FramebufferAttachment GLenum
	RenderbufferTarget    GLenum
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
func (c *GlFramebufferRenderbuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlFramebufferRenderbuffer) Flags() atom.Flags                { return 0 }
func (a *GlFramebufferRenderbuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferTexture2D
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferTexture2D struct {
	binary.Generate
	observations          atom.Observations
	FramebufferTarget     GLenum
	FramebufferAttachment GLenum
	TextureTarget         GLenum
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
func (c *GlFramebufferTexture2D) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlFramebufferTexture2D) Flags() atom.Flags                { return 0 }
func (a *GlFramebufferTexture2D) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetFramebufferAttachmentParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetFramebufferAttachmentParameteriv struct {
	binary.Generate
	observations      atom.Observations
	FramebufferTarget GLenum
	Attachment        GLenum
	Parameter         GLenum
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
func (c *GlGetFramebufferAttachmentParameteriv) API() gfxapi.ID    { return api{}.ID() }
func (c *GlGetFramebufferAttachmentParameteriv) Flags() atom.Flags { return 0 }
func (a *GlGetFramebufferAttachmentParameteriv) Observations() *atom.Observations {
	return &a.observations
}

////////////////////////////////////////////////////////////////////////////////
// GlDrawElements
////////////////////////////////////////////////////////////////////////////////
type GlDrawElements struct {
	binary.Generate
	observations atom.Observations
	DrawMode     GLenum
	ElementCount int32
	IndicesType  GLenum
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
func (c *GlDrawElements) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlDrawElements) Flags() atom.Flags                { return 0 | atom.DrawCall }
func (a *GlDrawElements) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDrawArrays
////////////////////////////////////////////////////////////////////////////////
type GlDrawArrays struct {
	binary.Generate
	observations atom.Observations
	DrawMode     GLenum
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
func (c *GlDrawArrays) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlFlush) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlFinish) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlFinish) Flags() atom.Flags                { return 0 }
func (a *GlFinish) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetBooleanv
////////////////////////////////////////////////////////////////////////////////
type GlGetBooleanv struct {
	binary.Generate
	observations atom.Observations
	Param        GLenum
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
func (c *GlGetBooleanv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetBooleanv) Flags() atom.Flags                { return 0 }
func (a *GlGetBooleanv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetFloatv
////////////////////////////////////////////////////////////////////////////////
type GlGetFloatv struct {
	binary.Generate
	observations atom.Observations
	Param        GLenum
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
func (c *GlGetFloatv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetFloatv) Flags() atom.Flags                { return 0 }
func (a *GlGetFloatv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetIntegerv
////////////////////////////////////////////////////////////////////////////////
type GlGetIntegerv struct {
	binary.Generate
	observations atom.Observations
	Param        GLenum
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
func (c *GlGetIntegerv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetIntegerv) Flags() atom.Flags                { return 0 }
func (a *GlGetIntegerv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetString
////////////////////////////////////////////////////////////////////////////////
type GlGetString struct {
	binary.Generate
	observations atom.Observations
	Param        GLenum
	Result       Charᶜᵖ
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
func (c *GlGetString) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetString) Flags() atom.Flags                { return 0 }
func (a *GlGetString) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEnable
////////////////////////////////////////////////////////////////////////////////
type GlEnable struct {
	binary.Generate
	observations atom.Observations
	Capability   GLenum
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
func (c *GlEnable) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlEnable) Flags() atom.Flags                { return 0 }
func (a *GlEnable) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDisable
////////////////////////////////////////////////////////////////////////////////
type GlDisable struct {
	binary.Generate
	observations atom.Observations
	Capability   GLenum
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
func (c *GlDisable) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlDisable) Flags() atom.Flags                { return 0 }
func (a *GlDisable) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlIsEnabled
////////////////////////////////////////////////////////////////////////////////
type GlIsEnabled struct {
	binary.Generate
	observations atom.Observations
	Capability   GLenum
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
func (c *GlIsEnabled) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlIsEnabled) Flags() atom.Flags                { return 0 }
func (a *GlIsEnabled) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlFenceSync
////////////////////////////////////////////////////////////////////////////////
type GlFenceSync struct {
	binary.Generate
	observations atom.Observations
	Condition    SyncCondition
	SyncFlags    SyncFlags
	Result       SyncObject
}

func (a *GlFenceSync) String() string {
	return fmt.Sprintf("glFenceSync(condition: %v, syncFlags: %v) → %v", a.Condition, a.SyncFlags, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlFenceSync pointer is returned so that calls can be chained.
func (a *GlFenceSync) AddRead(rng memory.Range, id binary.ID) *GlFenceSync {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlFenceSync pointer is returned so that calls can be chained.
func (a *GlFenceSync) AddWrite(rng memory.Range, id binary.ID) *GlFenceSync {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlFenceSync) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlFenceSync) Flags() atom.Flags                { return 0 }
func (a *GlFenceSync) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteSync
////////////////////////////////////////////////////////////////////////////////
type GlDeleteSync struct {
	binary.Generate
	observations atom.Observations
	Sync         SyncObject
}

func (a *GlDeleteSync) String() string {
	return fmt.Sprintf("glDeleteSync(sync: %v)", a.Sync)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteSync pointer is returned so that calls can be chained.
func (a *GlDeleteSync) AddRead(rng memory.Range, id binary.ID) *GlDeleteSync {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteSync pointer is returned so that calls can be chained.
func (a *GlDeleteSync) AddWrite(rng memory.Range, id binary.ID) *GlDeleteSync {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteSync) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlDeleteSync) Flags() atom.Flags                { return 0 }
func (a *GlDeleteSync) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlWaitSync
////////////////////////////////////////////////////////////////////////////////
type GlWaitSync struct {
	binary.Generate
	observations atom.Observations
	Sync         SyncObject
	SyncFlags    SyncFlags
	Timeout      uint64
}

func (a *GlWaitSync) String() string {
	return fmt.Sprintf("glWaitSync(sync: %v, syncFlags: %v, timeout: %v)", a.Sync, a.SyncFlags, a.Timeout)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlWaitSync pointer is returned so that calls can be chained.
func (a *GlWaitSync) AddRead(rng memory.Range, id binary.ID) *GlWaitSync {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlWaitSync pointer is returned so that calls can be chained.
func (a *GlWaitSync) AddWrite(rng memory.Range, id binary.ID) *GlWaitSync {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlWaitSync) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlWaitSync) Flags() atom.Flags                { return 0 }
func (a *GlWaitSync) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlClientWaitSync
////////////////////////////////////////////////////////////////////////////////
type GlClientWaitSync struct {
	binary.Generate
	observations atom.Observations
	Sync         SyncObject
	SyncFlags    SyncFlags
	Timeout      uint64
	Result       ClientWaitSyncSignal
}

func (a *GlClientWaitSync) String() string {
	return fmt.Sprintf("glClientWaitSync(sync: %v, syncFlags: %v, timeout: %v) → %v", a.Sync, a.SyncFlags, a.Timeout, a.Result)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlClientWaitSync pointer is returned so that calls can be chained.
func (a *GlClientWaitSync) AddRead(rng memory.Range, id binary.ID) *GlClientWaitSync {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlClientWaitSync pointer is returned so that calls can be chained.
func (a *GlClientWaitSync) AddWrite(rng memory.Range, id binary.ID) *GlClientWaitSync {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlClientWaitSync) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlClientWaitSync) Flags() atom.Flags                { return 0 }
func (a *GlClientWaitSync) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlMapBufferRange
////////////////////////////////////////////////////////////////////////////////
type GlMapBufferRange struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Offset       int32
	Length       int32
	Access       GLbitfield
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
func (c *GlMapBufferRange) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlMapBufferRange) Flags() atom.Flags                { return 0 }
func (a *GlMapBufferRange) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUnmapBuffer
////////////////////////////////////////////////////////////////////////////////
type GlUnmapBuffer struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlUnmapBuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlUnmapBuffer) Flags() atom.Flags                { return 0 }
func (a *GlUnmapBuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlInvalidateFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlInvalidateFramebuffer struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Count        int32
	Attachments  GLenumᶜᵖ
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
func (c *GlInvalidateFramebuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlInvalidateFramebuffer) Flags() atom.Flags                { return 0 }
func (a *GlInvalidateFramebuffer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorageMultisample
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorageMultisample struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Samples      int32
	Format       GLenum
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
func (c *GlRenderbufferStorageMultisample) API() gfxapi.ID                   { return api{}.ID() }
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
	Mask         GLbitfield
	Filter       GLenum
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
func (c *GlBlitFramebuffer) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlGenQueries) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGenQueries) Flags() atom.Flags                { return 0 }
func (a *GlGenQueries) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBeginQuery
////////////////////////////////////////////////////////////////////////////////
type GlBeginQuery struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlBeginQuery) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBeginQuery) Flags() atom.Flags                { return 0 }
func (a *GlBeginQuery) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEndQuery
////////////////////////////////////////////////////////////////////////////////
type GlEndQuery struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlEndQuery) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlEndQuery) Flags() atom.Flags                { return 0 }
func (a *GlEndQuery) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueries
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueries struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Queries      QueryIdᶜᵖ
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
func (c *GlDeleteQueries) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsQuery) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlIsQuery) Flags() atom.Flags                { return 0 }
func (a *GlIsQuery) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryiv struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
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
func (c *GlGetQueryiv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryiv) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuiv struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    GLenum
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
func (c *GlGetQueryObjectuiv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryObjectuiv) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectuiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveUniformBlockName
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveUniformBlockName struct {
	binary.Generate
	observations       atom.Observations
	Program            ProgramId
	UniformBlockIndex  uint32
	BufferSize         int32
	BufferBytesWritten S32ᵖ
	Name               Charᵖ
}

func (a *GlGetActiveUniformBlockName) String() string {
	return fmt.Sprintf("glGetActiveUniformBlockName(program: %v, uniform_block_index: %v, buffer_size: %v, buffer_bytes_written: %v, name: %v)", a.Program, a.UniformBlockIndex, a.BufferSize, a.BufferBytesWritten, a.Name)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniformBlockName pointer is returned so that calls can be chained.
func (a *GlGetActiveUniformBlockName) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockName {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniformBlockName pointer is returned so that calls can be chained.
func (a *GlGetActiveUniformBlockName) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockName {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetActiveUniformBlockName) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetActiveUniformBlockName) Flags() atom.Flags                { return 0 }
func (a *GlGetActiveUniformBlockName) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveUniformBlockiv
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveUniformBlockiv struct {
	binary.Generate
	observations      atom.Observations
	Program           ProgramId
	UniformBlockIndex uint32
	ParameterName     GLenum
	Parameters        S32ᵖ
}

func (a *GlGetActiveUniformBlockiv) String() string {
	return fmt.Sprintf("glGetActiveUniformBlockiv(program: %v, uniform_block_index: %v, parameter_name: %v, parameters: %v)", a.Program, a.UniformBlockIndex, a.ParameterName, a.Parameters)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniformBlockiv pointer is returned so that calls can be chained.
func (a *GlGetActiveUniformBlockiv) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockiv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniformBlockiv pointer is returned so that calls can be chained.
func (a *GlGetActiveUniformBlockiv) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockiv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetActiveUniformBlockiv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetActiveUniformBlockiv) Flags() atom.Flags                { return 0 }
func (a *GlGetActiveUniformBlockiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlUniformBlockBinding
////////////////////////////////////////////////////////////////////////////////
type GlUniformBlockBinding struct {
	binary.Generate
	observations        atom.Observations
	Program             ProgramId
	UniformBlockIndex   uint32
	UniformBlockBinding uint32
}

func (a *GlUniformBlockBinding) String() string {
	return fmt.Sprintf("glUniformBlockBinding(program: %v, uniform_block_index: %v, uniform_block_binding: %v)", a.Program, a.UniformBlockIndex, a.UniformBlockBinding)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlUniformBlockBinding pointer is returned so that calls can be chained.
func (a *GlUniformBlockBinding) AddRead(rng memory.Range, id binary.ID) *GlUniformBlockBinding {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlUniformBlockBinding pointer is returned so that calls can be chained.
func (a *GlUniformBlockBinding) AddWrite(rng memory.Range, id binary.ID) *GlUniformBlockBinding {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlUniformBlockBinding) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlUniformBlockBinding) Flags() atom.Flags                { return 0 }
func (a *GlUniformBlockBinding) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveUniformsiv
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveUniformsiv struct {
	binary.Generate
	observations   atom.Observations
	Program        ProgramId
	UniformCount   uint32
	UniformIndices U32ᵖ
	ParameterName  GLenum
	Parameters     S32ᵖ
}

func (a *GlGetActiveUniformsiv) String() string {
	return fmt.Sprintf("glGetActiveUniformsiv(program: %v, uniform_count: %v, uniform_indices: %v, parameter_name: %v, parameters: %v)", a.Program, a.UniformCount, a.UniformIndices, a.ParameterName, a.Parameters)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniformsiv pointer is returned so that calls can be chained.
func (a *GlGetActiveUniformsiv) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniformsiv {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetActiveUniformsiv pointer is returned so that calls can be chained.
func (a *GlGetActiveUniformsiv) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniformsiv {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetActiveUniformsiv) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetActiveUniformsiv) Flags() atom.Flags                { return 0 }
func (a *GlGetActiveUniformsiv) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindBufferBase
////////////////////////////////////////////////////////////////////////////////
type GlBindBufferBase struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Index        uint32
	Buffer       BufferId
}

func (a *GlBindBufferBase) String() string {
	return fmt.Sprintf("glBindBufferBase(target: %v, index: %v, buffer: %v)", a.Target, a.Index, a.Buffer)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindBufferBase pointer is returned so that calls can be chained.
func (a *GlBindBufferBase) AddRead(rng memory.Range, id binary.ID) *GlBindBufferBase {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindBufferBase pointer is returned so that calls can be chained.
func (a *GlBindBufferBase) AddWrite(rng memory.Range, id binary.ID) *GlBindBufferBase {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindBufferBase) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindBufferBase) Flags() atom.Flags                { return 0 }
func (a *GlBindBufferBase) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGenVertexArrays
////////////////////////////////////////////////////////////////////////////////
type GlGenVertexArrays struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Arrays       VertexArrayIdᵖ
}

func (a *GlGenVertexArrays) String() string {
	return fmt.Sprintf("glGenVertexArrays(count: %v, arrays: %v)", a.Count, a.Arrays)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGenVertexArrays pointer is returned so that calls can be chained.
func (a *GlGenVertexArrays) AddRead(rng memory.Range, id binary.ID) *GlGenVertexArrays {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGenVertexArrays pointer is returned so that calls can be chained.
func (a *GlGenVertexArrays) AddWrite(rng memory.Range, id binary.ID) *GlGenVertexArrays {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGenVertexArrays) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGenVertexArrays) Flags() atom.Flags                { return 0 }
func (a *GlGenVertexArrays) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBindVertexArray
////////////////////////////////////////////////////////////////////////////////
type GlBindVertexArray struct {
	binary.Generate
	observations atom.Observations
	Array        VertexArrayId
}

func (a *GlBindVertexArray) String() string {
	return fmt.Sprintf("glBindVertexArray(array: %v)", a.Array)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlBindVertexArray pointer is returned so that calls can be chained.
func (a *GlBindVertexArray) AddRead(rng memory.Range, id binary.ID) *GlBindVertexArray {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlBindVertexArray pointer is returned so that calls can be chained.
func (a *GlBindVertexArray) AddWrite(rng memory.Range, id binary.ID) *GlBindVertexArray {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlBindVertexArray) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBindVertexArray) Flags() atom.Flags                { return 0 }
func (a *GlBindVertexArray) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteVertexArrays
////////////////////////////////////////////////////////////////////////////////
type GlDeleteVertexArrays struct {
	binary.Generate
	observations atom.Observations
	Count        uint32
	Arrays       VertexArrayIdᶜᵖ
}

func (a *GlDeleteVertexArrays) String() string {
	return fmt.Sprintf("glDeleteVertexArrays(count: %v, arrays: %v)", a.Count, a.Arrays)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlDeleteVertexArrays pointer is returned so that calls can be chained.
func (a *GlDeleteVertexArrays) AddRead(rng memory.Range, id binary.ID) *GlDeleteVertexArrays {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlDeleteVertexArrays pointer is returned so that calls can be chained.
func (a *GlDeleteVertexArrays) AddWrite(rng memory.Range, id binary.ID) *GlDeleteVertexArrays {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlDeleteVertexArrays) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlDeleteVertexArrays) Flags() atom.Flags                { return 0 }
func (a *GlDeleteVertexArrays) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjecti64v
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjecti64v struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    GLenum
	Value        S64ᵖ
}

func (a *GlGetQueryObjecti64v) String() string {
	return fmt.Sprintf("glGetQueryObjecti64v(query: %v, parameter: %v, value: %v)", a.Query, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjecti64v pointer is returned so that calls can be chained.
func (a *GlGetQueryObjecti64v) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjecti64v {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjecti64v pointer is returned so that calls can be chained.
func (a *GlGetQueryObjecti64v) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjecti64v {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryObjecti64v) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryObjecti64v) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjecti64v) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectui64v
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectui64v struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    GLenum
	Value        U64ᵖ
}

func (a *GlGetQueryObjectui64v) String() string {
	return fmt.Sprintf("glGetQueryObjectui64v(query: %v, parameter: %v, value: %v)", a.Query, a.Parameter, a.Value)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectui64v pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectui64v) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectui64v {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The GlGetQueryObjectui64v pointer is returned so that calls can be chained.
func (a *GlGetQueryObjectui64v) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectui64v {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *GlGetQueryObjectui64v) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryObjectui64v) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectui64v) Observations() *atom.Observations { return &a.observations }

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
func (c *GlGenQueriesEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGenQueriesEXT) Flags() atom.Flags                { return 0 }
func (a *GlGenQueriesEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlBeginQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlBeginQueryEXT struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlBeginQueryEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlBeginQueryEXT) Flags() atom.Flags                { return 0 }
func (a *GlBeginQueryEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlEndQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlEndQueryEXT struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
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
func (c *GlEndQueryEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlEndQueryEXT) Flags() atom.Flags                { return 0 }
func (a *GlEndQueryEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueriesEXT
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueriesEXT struct {
	binary.Generate
	observations atom.Observations
	Count        int32
	Queries      QueryIdᶜᵖ
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
func (c *GlDeleteQueriesEXT) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *GlIsQueryEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlIsQueryEXT) Flags() atom.Flags                { return 0 }
func (a *GlIsQueryEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlQueryCounterEXT
////////////////////////////////////////////////////////////////////////////////
type GlQueryCounterEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Target       GLenum
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
func (c *GlQueryCounterEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlQueryCounterEXT) Flags() atom.Flags                { return 0 }
func (a *GlQueryCounterEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryivEXT struct {
	binary.Generate
	observations atom.Observations
	Target       GLenum
	Parameter    GLenum
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
func (c *GlGetQueryivEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryivEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryivEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectivEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    GLenum
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
func (c *GlGetQueryObjectivEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryObjectivEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectivEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuivEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    GLenum
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
func (c *GlGetQueryObjectuivEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryObjectuivEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectuivEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjecti64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjecti64vEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    GLenum
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
func (c *GlGetQueryObjecti64vEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryObjecti64vEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjecti64vEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectui64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectui64vEXT struct {
	binary.Generate
	observations atom.Observations
	Query        QueryId
	Parameter    GLenum
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
func (c *GlGetQueryObjectui64vEXT) API() gfxapi.ID                   { return api{}.ID() }
func (c *GlGetQueryObjectui64vEXT) Flags() atom.Flags                { return 0 }
func (a *GlGetQueryObjectui64vEXT) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// Architecture
////////////////////////////////////////////////////////////////////////////////
type Architecture struct {
	binary.Generate
	observations     atom.Observations
	PointerAlignment uint32
	PointerSize      uint32
	IntegerSize      uint32
	LittleEndian     bool
}

func (a *Architecture) String() string {
	return fmt.Sprintf("architecture(pointer_alignment: %v, pointer_size: %v, integer_size: %v, little_endian: %v)", a.PointerAlignment, a.PointerSize, a.IntegerSize, a.LittleEndian)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The Architecture pointer is returned so that calls can be chained.
func (a *Architecture) AddRead(rng memory.Range, id binary.ID) *Architecture {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The Architecture pointer is returned so that calls can be chained.
func (a *Architecture) AddWrite(rng memory.Range, id binary.ID) *Architecture {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *Architecture) API() gfxapi.ID                   { return api{}.ID() }
func (c *Architecture) Flags() atom.Flags                { return 0 }
func (a *Architecture) Observations() *atom.Observations { return &a.observations }

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
func (c *ReplayCreateRenderer) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *ReplayBindRenderer) API() gfxapi.ID                   { return api{}.ID() }
func (c *ReplayBindRenderer) Flags() atom.Flags                { return 0 }
func (a *ReplayBindRenderer) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// SwitchThread
////////////////////////////////////////////////////////////////////////////////
type SwitchThread struct {
	binary.Generate
	observations atom.Observations
	ThreadID     ThreadID
}

func (a *SwitchThread) String() string {
	return fmt.Sprintf("switchThread(threadID: %v)", a.ThreadID)
}

// AddRead appends a new read observation to the atom of the range rng with
// the data id.
// The SwitchThread pointer is returned so that calls can be chained.
func (a *SwitchThread) AddRead(rng memory.Range, id binary.ID) *SwitchThread {
	a.observations.Reads = append(a.observations.Reads, atom.Observation{Range: rng, ID: id})
	return a
}

// AddWrite appends a new write observation to the atom of the range rng with
// the data id.
// The SwitchThread pointer is returned so that calls can be chained.
func (a *SwitchThread) AddWrite(rng memory.Range, id binary.ID) *SwitchThread {
	a.observations.Writes = append(a.observations.Writes, atom.Observation{Range: rng, ID: id})
	return a
}
func (c *SwitchThread) API() gfxapi.ID                   { return api{}.ID() }
func (c *SwitchThread) Flags() atom.Flags                { return 0 }
func (a *SwitchThread) Observations() *atom.Observations { return &a.observations }

////////////////////////////////////////////////////////////////////////////////
// BackbufferInfo
////////////////////////////////////////////////////////////////////////////////
type BackbufferInfo struct {
	binary.Generate
	observations          atom.Observations
	Width                 int32
	Height                int32
	ColorFmt              GLenum
	DepthFmt              GLenum
	StencilFmt            GLenum
	ResetViewportScissor  bool
	PreserveBuffersOnSwap bool
}

func (a *BackbufferInfo) String() string {
	return fmt.Sprintf("backbufferInfo(width: %v, height: %v, color_fmt: %v, depth_fmt: %v, stencil_fmt: %v, resetViewportScissor: %v, preserveBuffersOnSwap: %v)", a.Width, a.Height, a.ColorFmt, a.DepthFmt, a.StencilFmt, a.ResetViewportScissor, a.PreserveBuffersOnSwap)
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
func (c *BackbufferInfo) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *StartTimer) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *StopTimer) API() gfxapi.ID                   { return api{}.ID() }
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
func (c *FlushPostBuffer) API() gfxapi.ID                   { return api{}.ID() }
func (c *FlushPostBuffer) Flags() atom.Flags                { return 0 }
func (a *FlushPostBuffer) Observations() *atom.Observations { return &a.observations }

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
	Format    GLenum
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
	Format    GLenum
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
	Format        GLenum
	Texture2D     S32ːImageᵐ
	Cubemap       S32ːCubemapLevelᵐ
	MagFilter     GLenum
	MinFilter     GLenum
	WrapS         GLenum
	WrapT         GLenum
	SwizzleR      GLenum
	SwizzleG      GLenum
	SwizzleB      GLenum
	SwizzleA      GLenum
	MaxAnisotropy float32
}

func (c *Texture) Init() {
	c.Texture2D = make(S32ːImageᵐ)
	c.Cubemap = make(S32ːCubemapLevelᵐ)
	c.MagFilter = GLenum_GL_LINEAR
	c.MinFilter = GLenum_GL_NEAREST_MIPMAP_LINEAR
	c.WrapS = GLenum_GL_REPEAT
	c.WrapT = GLenum_GL_REPEAT
	c.SwizzleR = GLenum_GL_RED
	c.SwizzleG = GLenum_GL_GREEN
	c.SwizzleB = GLenum_GL_BLUE
	c.SwizzleA = GLenum_GL_ALPHA
	c.MaxAnisotropy = 1
}
func (c *Texture) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class CubemapLevel
////////////////////////////////////////////////////////////////////////////////
type CubemapLevel struct {
	binary.Generate
	CreatedAt atom.ID
	Faces     GLenumːImageᵐ
}

func (c *CubemapLevel) Init() {
	c.Faces = make(GLenumːImageᵐ)
}
func (c *CubemapLevel) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class FramebufferAttachmentInfo
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachmentInfo struct {
	binary.Generate
	CreatedAt    atom.ID
	Object       uint32
	Type         GLenum
	TextureLevel int32
	CubeMapFace  GLenum
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
	Attachments GLenumːFramebufferAttachmentInfoᵐ
}

func (c *Framebuffer) Init() {
	c.Attachments = make(GLenumːFramebufferAttachmentInfoᵐ)
}
func (c *Framebuffer) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Buffer
////////////////////////////////////////////////////////////////////////////////
type Buffer struct {
	binary.Generate
	CreatedAt     atom.ID
	Data          U8ˢ
	Size          int32
	Usage         GLenum
	MappingAccess GLbitfield
	MappingOffset int32
	MappingData   U8ˢ
}

func (c *Buffer) Init() {
	c.Size = 0
	c.Usage = GLenum_GL_STATIC_DRAW
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
	Type      GLenum
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
	Type        GLenum
}

func (c *VertexAttribute) Init() {
}
func (c *VertexAttribute) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Uniform
////////////////////////////////////////////////////////////////////////////////
type Uniform struct {
	binary.Generate
	CreatedAt atom.ID
	Name      string
	Type      GLenum
	Value     U8ˢ
}

func (c *Uniform) Init() {
}
func (c *Uniform) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Program
////////////////////////////////////////////////////////////////////////////////
type Program struct {
	binary.Generate
	CreatedAt         atom.ID
	Shaders           GLenumːShaderIdᵐ
	Linked            bool
	Binary            U8ˢ
	AttributeBindings StringːAttributeLocationᵐ
	Attributes        S32ːVertexAttributeᵐ
	Uniforms          UniformLocationːUniformᵐ
	InfoLog           Charˢ
}

func (c *Program) Init() {
	c.Shaders = make(GLenumːShaderIdᵐ)
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
	Type       GLenum
	Normalized bool
	Stride     int32
	Buffer     BufferId
	Pointer    VertexPointer
}

func (c *VertexAttributeArray) Init() {
	c.Enabled = false
	c.Size = 4
	c.Type = GLenum_GL_FLOAT
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
	SrcRgbBlendFactor   GLenum
	SrcAlphaBlendFactor GLenum
	DstRgbBlendFactor   GLenum
	DstAlphaBlendFactor GLenum
	BlendEquationRgb    GLenum
	BlendEquationAlpha  GLenum
	BlendColor          Color
}

func (c *BlendState) Init() {
	c.SrcRgbBlendFactor = GLenum_GL_ONE
	c.SrcAlphaBlendFactor = GLenum_GL_ZERO
	c.DstRgbBlendFactor = GLenum_GL_ONE
	c.DstAlphaBlendFactor = GLenum_GL_ZERO
	c.BlendEquationRgb = GLenum_GL_FUNC_ADD
	c.BlendEquationAlpha = GLenum_GL_FUNC_ADD
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
	DepthTestFunction    GLenum
	DepthNear            float32
	DepthFar             float32
	ColorMaskRed         bool
	ColorMaskGreen       bool
	ColorMaskBlue        bool
	ColorMaskAlpha       bool
	StencilMask          GLenumːu32ᵐ
	Viewport             Rect
	Scissor              Rect
	FrontFace            GLenum
	CullFace             GLenum
	LineWidth            float32
	PolygonOffsetFactor  float32
	PolygonOffsetUnits   float32
	SampleCoverageValue  float32
	SampleCoverageInvert bool
}

func (c *RasterizerState) Init() {
	c.DepthMask = true
	c.DepthTestFunction = GLenum_GL_LESS
	c.DepthNear = 0
	c.DepthFar = 1
	c.ColorMaskRed = true
	c.ColorMaskGreen = true
	c.ColorMaskBlue = true
	c.ColorMaskAlpha = true
	c.StencilMask = make(GLenumːu32ᵐ)
	c.Viewport.Init()
	c.Scissor.Init()
	c.FrontFace = GLenum_GL_CCW
	c.CullFace = GLenum_GL_BACK
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
	BoundFramebuffers     GLenumːFramebufferIdᵐ
	BoundRenderbuffers    GLenumːRenderbufferIdᵐ
	BoundBuffers          GLenumːBufferIdᵐ
	BoundProgram          ProgramId
	BoundVertexArray      VertexArrayId
	VertexAttributeArrays AttributeLocationːVertexAttributeArrayʳᵐ
	TextureUnits          GLenumːGLenumːTextureIdᵐᵐ
	ActiveTextureUnit     GLenum
	Capabilities          GLenumːboolᵐ
	GenerateMipmapHint    GLenum
	PixelStorage          GLenumːs32ᵐ
	Instances             Objects
	PreserveBuffersOnSwap bool
}

func (c *Context) Init() {
	c.Blending.Init()
	c.Rasterizing.Init()
	c.Clearing.Init()
	c.BoundFramebuffers = make(GLenumːFramebufferIdᵐ)
	c.BoundRenderbuffers = make(GLenumːRenderbufferIdᵐ)
	c.BoundBuffers = make(GLenumːBufferIdᵐ)
	c.VertexAttributeArrays = make(AttributeLocationːVertexAttributeArrayʳᵐ)
	c.TextureUnits = make(GLenumːGLenumːTextureIdᵐᵐ)
	c.ActiveTextureUnit = GLenum_GL_TEXTURE0
	c.Capabilities = make(GLenumːboolᵐ)
	c.GenerateMipmapHint = GLenum_GL_DONT_CARE
	c.PixelStorage = make(GLenumːs32ᵐ)
	c.Instances.Init()
}
func (c *Context) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class __GLsync
////////////////////////////////////////////////////////////////////////////////
type __GLsync struct {
	binary.Generate
	CreatedAt atom.ID
}

func (c *__GLsync) Init() {
}
func (c *__GLsync) GetCreatedAt() atom.ID { return c.CreatedAt }

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
// enum GLenum
////////////////////////////////////////////////////////////////////////////////
type GLenum uint32

const (
	GLenum_GL_LINE_LOOP                                    = GLenum(2)
	GLenum_GL_LINE_STRIP                                   = GLenum(3)
	GLenum_GL_LINES                                        = GLenum(1)
	GLenum_GL_POINTS                                       = GLenum(0)
	GLenum_GL_TRIANGLE_FAN                                 = GLenum(6)
	GLenum_GL_TRIANGLE_STRIP                               = GLenum(5)
	GLenum_GL_TRIANGLES                                    = GLenum(4)
	GLenum_GL_UNSIGNED_BYTE                                = GLenum(5121)
	GLenum_GL_UNSIGNED_SHORT                               = GLenum(5123)
	GLenum_GL_UNSIGNED_INT                                 = GLenum(5125)
	GLenum_GL_TEXTURE_2D                                   = GLenum(3553)
	GLenum_GL_TEXTURE_CUBE_MAP                             = GLenum(34067)
	GLenum_GL_TEXTURE_EXTERNAL_OES                         = GLenum(36197)
	GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_X                  = GLenum(34070)
	GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y                  = GLenum(34072)
	GLenum_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z                  = GLenum(34074)
	GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_X                  = GLenum(34069)
	GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Y                  = GLenum(34071)
	GLenum_GL_TEXTURE_CUBE_MAP_POSITIVE_Z                  = GLenum(34073)
	GLenum_GL_ALPHA                                        = GLenum(6406)
	GLenum_GL_RGB                                          = GLenum(6407)
	GLenum_GL_RGBA                                         = GLenum(6408)
	GLenum_GL_LUMINANCE                                    = GLenum(6409)
	GLenum_GL_LUMINANCE_ALPHA                              = GLenum(6410)
	GLenum_GL_RED                                          = GLenum(6403)
	GLenum_GL_RED_INTEGER                                  = GLenum(36244)
	GLenum_GL_RG                                           = GLenum(33319)
	GLenum_GL_RG_INTEGER                                   = GLenum(33320)
	GLenum_GL_RGB_INTEGER                                  = GLenum(36248)
	GLenum_GL_RGBA_INTEGER                                 = GLenum(36249)
	GLenum_GL_DEPTH_COMPONENT                              = GLenum(6402)
	GLenum_GL_DEPTH_COMPONENT16                            = GLenum(33189)
	GLenum_GL_DEPTH_STENCIL                                = GLenum(34041)
	GLenum_GL_DEPTH24_STENCIL8                             = GLenum(35056)
	GLenum_GL_RGBA4                                        = GLenum(32854)
	GLenum_GL_RGB5_A1                                      = GLenum(32855)
	GLenum_GL_RGB565                                       = GLenum(36194)
	GLenum_GL_RGBA8                                        = GLenum(32856)
	GLenum_GL_STENCIL_INDEX8                               = GLenum(36168)
	GLenum_GL_HALF_FLOAT_ARB                               = GLenum(5131)
	GLenum_GL_HALF_FLOAT_OES                               = GLenum(36193)
	GLenum_GL_ETC1_RGB8_OES                                = GLenum(36196)
	GLenum_GL_ATC_RGB_AMD                                  = GLenum(35986)
	GLenum_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD                  = GLenum(35987)
	GLenum_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD              = GLenum(34798)
	GLenum_GL_COMPRESSED_RGB_S3TC_DXT1_EXT                 = GLenum(33776)
	GLenum_GL_COMPRESSED_RGBA_S3TC_DXT1_EXT                = GLenum(33777)
	GLenum_GL_COMPRESSED_RGBA_S3TC_DXT3_EXT                = GLenum(33778)
	GLenum_GL_COMPRESSED_RGBA_S3TC_DXT5_EXT                = GLenum(33779)
	GLenum_GL_COMPRESSED_RGBA_ASTC_4x4_KHR                 = GLenum(37808)
	GLenum_GL_COMPRESSED_RGBA_ASTC_5x4_KHR                 = GLenum(37809)
	GLenum_GL_COMPRESSED_RGBA_ASTC_5x5_KHR                 = GLenum(37810)
	GLenum_GL_COMPRESSED_RGBA_ASTC_6x5_KHR                 = GLenum(37811)
	GLenum_GL_COMPRESSED_RGBA_ASTC_6x6_KHR                 = GLenum(37812)
	GLenum_GL_COMPRESSED_RGBA_ASTC_8x5_KHR                 = GLenum(37813)
	GLenum_GL_COMPRESSED_RGBA_ASTC_8x6_KHR                 = GLenum(37814)
	GLenum_GL_COMPRESSED_RGBA_ASTC_8x8_KHR                 = GLenum(37815)
	GLenum_GL_COMPRESSED_RGBA_ASTC_10x5_KHR                = GLenum(37816)
	GLenum_GL_COMPRESSED_RGBA_ASTC_10x6_KHR                = GLenum(37817)
	GLenum_GL_COMPRESSED_RGBA_ASTC_10x8_KHR                = GLenum(37818)
	GLenum_GL_COMPRESSED_RGBA_ASTC_10x10_KHR               = GLenum(37819)
	GLenum_GL_COMPRESSED_RGBA_ASTC_12x10_KHR               = GLenum(37820)
	GLenum_GL_COMPRESSED_RGBA_ASTC_12x12_KHR               = GLenum(37821)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR         = GLenum(37840)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR         = GLenum(37841)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR         = GLenum(37842)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR         = GLenum(37843)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR         = GLenum(37844)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR         = GLenum(37845)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR         = GLenum(37846)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR         = GLenum(37847)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR        = GLenum(37848)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR        = GLenum(37849)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR        = GLenum(37850)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR       = GLenum(37851)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR       = GLenum(37852)
	GLenum_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR       = GLenum(37853)
	GLenum_GL_COMPRESSED_LUMINANCE_LATC1_NV                = GLenum(35952)
	GLenum_GL_COMPRESSED_SIGNED_LUMINANCE_LATC1_NV         = GLenum(35953)
	GLenum_GL_COMPRESSED_LUMINANCE_ALPHA_LATC2_NV          = GLenum(35954)
	GLenum_GL_COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_NV   = GLenum(35955)
	GLenum_GL_FLOAT                                        = GLenum(5126)
	GLenum_GL_UNSIGNED_SHORT_4_4_4_4                       = GLenum(32819)
	GLenum_GL_UNSIGNED_SHORT_5_5_5_1                       = GLenum(32820)
	GLenum_GL_UNSIGNED_SHORT_5_6_5                         = GLenum(33635)
	GLenum_GL_UNSIGNED_INT_24_8                            = GLenum(34042)
	GLenum_GL_COLOR_ATTACHMENT0                            = GLenum(36064)
	GLenum_GL_DEPTH_ATTACHMENT                             = GLenum(36096)
	GLenum_GL_STENCIL_ATTACHMENT                           = GLenum(36128)
	GLenum_GL_NONE                                         = GLenum(0)
	GLenum_GL_RENDERBUFFER                                 = GLenum(36161)
	GLenum_GL_TEXTURE                                      = GLenum(5890)
	GLenum_GL_FRAMEBUFFER                                  = GLenum(36160)
	GLenum_GL_READ_FRAMEBUFFER                             = GLenum(36008)
	GLenum_GL_DRAW_FRAMEBUFFER                             = GLenum(36009)
	GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = GLenum(36048)
	GLenum_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = GLenum(36049)
	GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = GLenum(36050)
	GLenum_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = GLenum(36051)
	GLenum_GL_FRAMEBUFFER_COMPLETE                         = GLenum(36053)
	GLenum_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT            = GLenum(36054)
	GLenum_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT    = GLenum(36055)
	GLenum_GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS            = GLenum(36057)
	GLenum_GL_FRAMEBUFFER_UNSUPPORTED                      = GLenum(36061)
	GLenum_GL_RENDERBUFFER_WIDTH                           = GLenum(36162)
	GLenum_GL_RENDERBUFFER_HEIGHT                          = GLenum(36163)
	GLenum_GL_RENDERBUFFER_INTERNAL_FORMAT                 = GLenum(36164)
	GLenum_GL_RENDERBUFFER_RED_SIZE                        = GLenum(36176)
	GLenum_GL_RENDERBUFFER_GREEN_SIZE                      = GLenum(36177)
	GLenum_GL_RENDERBUFFER_BLUE_SIZE                       = GLenum(36178)
	GLenum_GL_RENDERBUFFER_ALPHA_SIZE                      = GLenum(36179)
	GLenum_GL_RENDERBUFFER_DEPTH_SIZE                      = GLenum(36180)
	GLenum_GL_RENDERBUFFER_STENCIL_SIZE                    = GLenum(36181)
	GLenum_GL_BUFFER_SIZE                                  = GLenum(34660)
	GLenum_GL_BUFFER_USAGE                                 = GLenum(34661)
	GLenum_GL_TEXTURE0                                     = GLenum(33984)
	GLenum_GL_TEXTURE1                                     = GLenum(33985)
	GLenum_GL_TEXTURE2                                     = GLenum(33986)
	GLenum_GL_TEXTURE3                                     = GLenum(33987)
	GLenum_GL_TEXTURE4                                     = GLenum(33988)
	GLenum_GL_TEXTURE5                                     = GLenum(33989)
	GLenum_GL_TEXTURE6                                     = GLenum(33990)
	GLenum_GL_TEXTURE7                                     = GLenum(33991)
	GLenum_GL_TEXTURE8                                     = GLenum(33992)
	GLenum_GL_TEXTURE9                                     = GLenum(33993)
	GLenum_GL_TEXTURE10                                    = GLenum(33994)
	GLenum_GL_TEXTURE11                                    = GLenum(33995)
	GLenum_GL_TEXTURE12                                    = GLenum(33996)
	GLenum_GL_TEXTURE13                                    = GLenum(33997)
	GLenum_GL_TEXTURE14                                    = GLenum(33998)
	GLenum_GL_TEXTURE15                                    = GLenum(33999)
	GLenum_GL_TEXTURE16                                    = GLenum(34000)
	GLenum_GL_TEXTURE17                                    = GLenum(34001)
	GLenum_GL_TEXTURE18                                    = GLenum(34002)
	GLenum_GL_TEXTURE19                                    = GLenum(34003)
	GLenum_GL_TEXTURE20                                    = GLenum(34004)
	GLenum_GL_TEXTURE21                                    = GLenum(34005)
	GLenum_GL_TEXTURE22                                    = GLenum(34006)
	GLenum_GL_TEXTURE23                                    = GLenum(34007)
	GLenum_GL_TEXTURE24                                    = GLenum(34008)
	GLenum_GL_TEXTURE25                                    = GLenum(34009)
	GLenum_GL_TEXTURE26                                    = GLenum(34010)
	GLenum_GL_TEXTURE27                                    = GLenum(34011)
	GLenum_GL_TEXTURE28                                    = GLenum(34012)
	GLenum_GL_TEXTURE29                                    = GLenum(34013)
	GLenum_GL_TEXTURE30                                    = GLenum(34014)
	GLenum_GL_TEXTURE31                                    = GLenum(34015)
	GLenum_GL_DYNAMIC_DRAW                                 = GLenum(35048)
	GLenum_GL_STATIC_DRAW                                  = GLenum(35044)
	GLenum_GL_STREAM_DRAW                                  = GLenum(35040)
	GLenum_GL_VERTEX_SHADER                                = GLenum(35633)
	GLenum_GL_FRAGMENT_SHADER                              = GLenum(35632)
	GLenum_GL_ACTIVE_TEXTURE                               = GLenum(34016)
	GLenum_GL_ALIASED_LINE_WIDTH_RANGE                     = GLenum(33902)
	GLenum_GL_ALIASED_POINT_SIZE_RANGE                     = GLenum(33901)
	GLenum_GL_ALPHA_BITS                                   = GLenum(3413)
	GLenum_GL_ARRAY_BUFFER_BINDING                         = GLenum(34964)
	GLenum_GL_BLEND                                        = GLenum(3042)
	GLenum_GL_BLEND_COLOR                                  = GLenum(32773)
	GLenum_GL_BLEND_DST_ALPHA                              = GLenum(32970)
	GLenum_GL_BLEND_DST_RGB                                = GLenum(32968)
	GLenum_GL_BLEND_EQUATION_ALPHA                         = GLenum(34877)
	GLenum_GL_BLEND_EQUATION_RGB                           = GLenum(32777)
	GLenum_GL_BLEND_SRC_ALPHA                              = GLenum(32971)
	GLenum_GL_BLEND_SRC_RGB                                = GLenum(32969)
	GLenum_GL_BLUE_BITS                                    = GLenum(3412)
	GLenum_GL_COLOR_CLEAR_VALUE                            = GLenum(3106)
	GLenum_GL_COLOR_WRITEMASK                              = GLenum(3107)
	GLenum_GL_COMPRESSED_TEXTURE_FORMATS                   = GLenum(34467)
	GLenum_GL_CULL_FACE                                    = GLenum(2884)
	GLenum_GL_CULL_FACE_MODE                               = GLenum(2885)
	GLenum_GL_CURRENT_PROGRAM                              = GLenum(35725)
	GLenum_GL_DEPTH_BITS                                   = GLenum(3414)
	GLenum_GL_DEPTH_CLEAR_VALUE                            = GLenum(2931)
	GLenum_GL_DEPTH_FUNC                                   = GLenum(2932)
	GLenum_GL_DEPTH_RANGE                                  = GLenum(2928)
	GLenum_GL_DEPTH_TEST                                   = GLenum(2929)
	GLenum_GL_DEPTH_WRITEMASK                              = GLenum(2930)
	GLenum_GL_DITHER                                       = GLenum(3024)
	GLenum_GL_ELEMENT_ARRAY_BUFFER_BINDING                 = GLenum(34965)
	GLenum_GL_FRAMEBUFFER_BINDING                          = GLenum(36006)
	GLenum_GL_FRONT_FACE                                   = GLenum(2886)
	GLenum_GL_GENERATE_MIPMAP_HINT                         = GLenum(33170)
	GLenum_GL_GREEN_BITS                                   = GLenum(3411)
	GLenum_GL_IMPLEMENTATION_COLOR_READ_FORMAT             = GLenum(35739)
	GLenum_GL_IMPLEMENTATION_COLOR_READ_TYPE               = GLenum(35738)
	GLenum_GL_LINE_WIDTH                                   = GLenum(2849)
	GLenum_GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS             = GLenum(35661)
	GLenum_GL_MAX_CUBE_MAP_TEXTURE_SIZE                    = GLenum(34076)
	GLenum_GL_MAX_FRAGMENT_UNIFORM_VECTORS                 = GLenum(36349)
	GLenum_GL_MAX_RENDERBUFFER_SIZE                        = GLenum(34024)
	GLenum_GL_MAX_TEXTURE_IMAGE_UNITS                      = GLenum(34930)
	GLenum_GL_MAX_TEXTURE_SIZE                             = GLenum(3379)
	GLenum_GL_MAX_VARYING_VECTORS                          = GLenum(36348)
	GLenum_GL_MAX_VERTEX_ATTRIBS                           = GLenum(34921)
	GLenum_GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS               = GLenum(35660)
	GLenum_GL_MAX_VERTEX_UNIFORM_VECTORS                   = GLenum(36347)
	GLenum_GL_MAX_VIEWPORT_DIMS                            = GLenum(3386)
	GLenum_GL_NUM_COMPRESSED_TEXTURE_FORMATS               = GLenum(34466)
	GLenum_GL_NUM_SHADER_BINARY_FORMATS                    = GLenum(36345)
	GLenum_GL_PACK_ALIGNMENT                               = GLenum(3333)
	GLenum_GL_POLYGON_OFFSET_FACTOR                        = GLenum(32824)
	GLenum_GL_POLYGON_OFFSET_FILL                          = GLenum(32823)
	GLenum_GL_POLYGON_OFFSET_UNITS                         = GLenum(10752)
	GLenum_GL_RED_BITS                                     = GLenum(3410)
	GLenum_GL_RENDERBUFFER_BINDING                         = GLenum(36007)
	GLenum_GL_SAMPLE_ALPHA_TO_COVERAGE                     = GLenum(32926)
	GLenum_GL_SAMPLE_BUFFERS                               = GLenum(32936)
	GLenum_GL_SAMPLE_COVERAGE                              = GLenum(32928)
	GLenum_GL_SAMPLE_COVERAGE_INVERT                       = GLenum(32939)
	GLenum_GL_SAMPLE_COVERAGE_VALUE                        = GLenum(32938)
	GLenum_GL_SAMPLES                                      = GLenum(32937)
	GLenum_GL_SCISSOR_BOX                                  = GLenum(3088)
	GLenum_GL_SCISSOR_TEST                                 = GLenum(3089)
	GLenum_GL_SHADER_BINARY_FORMATS                        = GLenum(36344)
	GLenum_GL_SHADER_COMPILER                              = GLenum(36346)
	GLenum_GL_STENCIL_BACK_FAIL                            = GLenum(34817)
	GLenum_GL_STENCIL_BACK_FUNC                            = GLenum(34816)
	GLenum_GL_STENCIL_BACK_PASS_DEPTH_FAIL                 = GLenum(34818)
	GLenum_GL_STENCIL_BACK_PASS_DEPTH_PASS                 = GLenum(34819)
	GLenum_GL_STENCIL_BACK_REF                             = GLenum(36003)
	GLenum_GL_STENCIL_BACK_VALUE_MASK                      = GLenum(36004)
	GLenum_GL_STENCIL_BACK_WRITEMASK                       = GLenum(36005)
	GLenum_GL_STENCIL_BITS                                 = GLenum(3415)
	GLenum_GL_STENCIL_CLEAR_VALUE                          = GLenum(2961)
	GLenum_GL_STENCIL_FAIL                                 = GLenum(2964)
	GLenum_GL_STENCIL_FUNC                                 = GLenum(2962)
	GLenum_GL_STENCIL_PASS_DEPTH_FAIL                      = GLenum(2965)
	GLenum_GL_STENCIL_PASS_DEPTH_PASS                      = GLenum(2966)
	GLenum_GL_STENCIL_REF                                  = GLenum(2967)
	GLenum_GL_STENCIL_TEST                                 = GLenum(2960)
	GLenum_GL_STENCIL_VALUE_MASK                           = GLenum(2963)
	GLenum_GL_STENCIL_WRITEMASK                            = GLenum(2968)
	GLenum_GL_SUBPIXEL_BITS                                = GLenum(3408)
	GLenum_GL_TEXTURE_BINDING_2D                           = GLenum(32873)
	GLenum_GL_TEXTURE_BINDING_CUBE_MAP                     = GLenum(34068)
	GLenum_GL_UNPACK_ALIGNMENT                             = GLenum(3317)
	GLenum_GL_VIEWPORT                                     = GLenum(2978)
	GLenum_GL_READ_FRAMEBUFFER_BINDING                     = GLenum(36010)
	GLenum_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT               = GLenum(34047)
	GLenum_GL_GPU_DISJOINT_EXT                             = GLenum(36795)
	GLenum_GL_FRONT                                        = GLenum(1028)
	GLenum_GL_BACK                                         = GLenum(1029)
	GLenum_GL_FRONT_AND_BACK                               = GLenum(1032)
	GLenum_GL_VERTEX_ARRAY                                 = GLenum(32884)
	GLenum_GL_NORMAL_ARRAY                                 = GLenum(32885)
	GLenum_GL_COLOR_ARRAY                                  = GLenum(32886)
	GLenum_GL_TEXTURE_COORD_ARRAY                          = GLenum(32888)
	GLenum_GL_POINT_SIZE_ARRAY_OES                         = GLenum(35740)
	GLenum_GL_EXTENSIONS                                   = GLenum(7939)
	GLenum_GL_RENDERER                                     = GLenum(7937)
	GLenum_GL_VENDOR                                       = GLenum(7936)
	GLenum_GL_VERSION                                      = GLenum(7938)
	GLenum_GL_BYTE                                         = GLenum(5120)
	GLenum_GL_FIXED                                        = GLenum(5132)
	GLenum_GL_SHORT                                        = GLenum(5122)
	GLenum_GL_FLOAT_VEC2                                   = GLenum(35664)
	GLenum_GL_FLOAT_VEC3                                   = GLenum(35665)
	GLenum_GL_FLOAT_VEC4                                   = GLenum(35666)
	GLenum_GL_FLOAT_MAT2                                   = GLenum(35674)
	GLenum_GL_FLOAT_MAT3                                   = GLenum(35675)
	GLenum_GL_FLOAT_MAT4                                   = GLenum(35676)
	GLenum_GL_INT                                          = GLenum(5124)
	GLenum_GL_INT_VEC2                                     = GLenum(35667)
	GLenum_GL_INT_VEC3                                     = GLenum(35668)
	GLenum_GL_INT_VEC4                                     = GLenum(35669)
	GLenum_GL_BOOL                                         = GLenum(35670)
	GLenum_GL_BOOL_VEC2                                    = GLenum(35671)
	GLenum_GL_BOOL_VEC3                                    = GLenum(35672)
	GLenum_GL_BOOL_VEC4                                    = GLenum(35673)
	GLenum_GL_SAMPLER_2D                                   = GLenum(35678)
	GLenum_GL_SAMPLER_CUBE                                 = GLenum(35680)
	GLenum_GL_NO_ERROR                                     = GLenum(0)
	GLenum_GL_INVALID_ENUM                                 = GLenum(1280)
	GLenum_GL_INVALID_VALUE                                = GLenum(1281)
	GLenum_GL_INVALID_OPERATION                            = GLenum(1282)
	GLenum_GL_INVALID_FRAMEBUFFER_OPERATION                = GLenum(1286)
	GLenum_GL_OUT_OF_MEMORY                                = GLenum(1285)
	GLenum_GL_DONT_CARE                                    = GLenum(4352)
	GLenum_GL_FASTEST                                      = GLenum(4353)
	GLenum_GL_NICEST                                       = GLenum(4354)
	GLenum_GL_COLOR_EXT                                    = GLenum(6144)
	GLenum_GL_DEPTH_EXT                                    = GLenum(6145)
	GLenum_GL_STENCIL_EXT                                  = GLenum(6146)
	GLenum_GL_DELETE_STATUS                                = GLenum(35712)
	GLenum_GL_LINK_STATUS                                  = GLenum(35714)
	GLenum_GL_VALIDATE_STATUS                              = GLenum(35715)
	GLenum_GL_INFO_LOG_LENGTH                              = GLenum(35716)
	GLenum_GL_ATTACHED_SHADERS                             = GLenum(35717)
	GLenum_GL_ACTIVE_ATTRIBUTES                            = GLenum(35721)
	GLenum_GL_ACTIVE_ATTRIBUTE_MAX_LENGTH                  = GLenum(35722)
	GLenum_GL_ACTIVE_UNIFORMS                              = GLenum(35718)
	GLenum_GL_ACTIVE_UNIFORM_MAX_LENGTH                    = GLenum(35719)
	GLenum_GL_SHADER_TYPE                                  = GLenum(35663)
	GLenum_GL_COMPILE_STATUS                               = GLenum(35713)
	GLenum_GL_SHADER_SOURCE_LENGTH                         = GLenum(35720)
	GLenum_GL_TEXTURE_MIN_FILTER                           = GLenum(10241)
	GLenum_GL_TEXTURE_MAG_FILTER                           = GLenum(10240)
	GLenum_GL_TEXTURE_WRAP_S                               = GLenum(10242)
	GLenum_GL_TEXTURE_WRAP_T                               = GLenum(10243)
	GLenum_GL_TEXTURE_MAX_ANISOTROPY_EXT                   = GLenum(34046)
	GLenum_GL_TEXTURE_SWIZZLE_R                            = GLenum(36418)
	GLenum_GL_TEXTURE_SWIZZLE_G                            = GLenum(36419)
	GLenum_GL_TEXTURE_SWIZZLE_B                            = GLenum(36420)
	GLenum_GL_TEXTURE_SWIZZLE_A                            = GLenum(36421)
	GLenum_GL_NEAREST                                      = GLenum(9728)
	GLenum_GL_LINEAR                                       = GLenum(9729)
	GLenum_GL_NEAREST_MIPMAP_NEAREST                       = GLenum(9984)
	GLenum_GL_LINEAR_MIPMAP_NEAREST                        = GLenum(9985)
	GLenum_GL_NEAREST_MIPMAP_LINEAR                        = GLenum(9986)
	GLenum_GL_LINEAR_MIPMAP_LINEAR                         = GLenum(9987)
	GLenum_GL_CLAMP_TO_EDGE                                = GLenum(33071)
	GLenum_GL_MIRRORED_REPEAT                              = GLenum(33648)
	GLenum_GL_REPEAT                                       = GLenum(10497)
	GLenum_GL_GREEN                                        = GLenum(6404)
	GLenum_GL_BLUE                                         = GLenum(6405)
	GLenum_GL_ZERO                                         = GLenum(0)
	GLenum_GL_ONE                                          = GLenum(1)
	GLenum_GL_SRC_COLOR                                    = GLenum(768)
	GLenum_GL_ONE_MINUS_SRC_COLOR                          = GLenum(769)
	GLenum_GL_DST_COLOR                                    = GLenum(774)
	GLenum_GL_ONE_MINUS_DST_COLOR                          = GLenum(775)
	GLenum_GL_SRC_ALPHA                                    = GLenum(770)
	GLenum_GL_ONE_MINUS_SRC_ALPHA                          = GLenum(771)
	GLenum_GL_DST_ALPHA                                    = GLenum(772)
	GLenum_GL_ONE_MINUS_DST_ALPHA                          = GLenum(773)
	GLenum_GL_CONSTANT_COLOR                               = GLenum(32769)
	GLenum_GL_ONE_MINUS_CONSTANT_COLOR                     = GLenum(32770)
	GLenum_GL_CONSTANT_ALPHA                               = GLenum(32771)
	GLenum_GL_ONE_MINUS_CONSTANT_ALPHA                     = GLenum(32772)
	GLenum_GL_SRC_ALPHA_SATURATE                           = GLenum(776)
	GLenum_GL_LOW_FLOAT                                    = GLenum(36336)
	GLenum_GL_MEDIUM_FLOAT                                 = GLenum(36337)
	GLenum_GL_HIGH_FLOAT                                   = GLenum(36338)
	GLenum_GL_LOW_INT                                      = GLenum(36339)
	GLenum_GL_MEDIUM_INT                                   = GLenum(36340)
	GLenum_GL_HIGH_INT                                     = GLenum(36341)
	GLenum_GL_NEVER                                        = GLenum(512)
	GLenum_GL_LESS                                         = GLenum(513)
	GLenum_GL_EQUAL                                        = GLenum(514)
	GLenum_GL_LEQUAL                                       = GLenum(515)
	GLenum_GL_GREATER                                      = GLenum(516)
	GLenum_GL_NOTEQUAL                                     = GLenum(517)
	GLenum_GL_GEQUAL                                       = GLenum(518)
	GLenum_GL_ALWAYS                                       = GLenum(519)
	GLenum_GL_KEEP                                         = GLenum(7680)
	GLenum_GL_REPLACE                                      = GLenum(7681)
	GLenum_GL_INCR                                         = GLenum(7682)
	GLenum_GL_INCR_WRAP                                    = GLenum(34055)
	GLenum_GL_DECR                                         = GLenum(7683)
	GLenum_GL_DECR_WRAP                                    = GLenum(34056)
	GLenum_GL_INVERT                                       = GLenum(5386)
	GLenum_GL_CW                                           = GLenum(2304)
	GLenum_GL_CCW                                          = GLenum(2305)
	GLenum_GL_FUNC_ADD                                     = GLenum(32774)
	GLenum_GL_FUNC_SUBTRACT                                = GLenum(32778)
	GLenum_GL_FUNC_REVERSE_SUBTRACT                        = GLenum(32779)
	GLenum_GL_ARRAY_BUFFER                                 = GLenum(34962)
	GLenum_GL_COPY_READ_BUFFER                             = GLenum(36662)
	GLenum_GL_COPY_WRITE_BUFFER                            = GLenum(36663)
	GLenum_GL_ELEMENT_ARRAY_BUFFER                         = GLenum(34963)
	GLenum_GL_PIXEL_PACK_BUFFER                            = GLenum(35051)
	GLenum_GL_PIXEL_UNPACK_BUFFER                          = GLenum(35052)
	GLenum_GL_TRANSFORM_FEEDBACK_BUFFER                    = GLenum(35982)
	GLenum_GL_UNIFORM_BUFFER                               = GLenum(35345)
	GLenum_GL_RENDERBUFFER_OES                             = GLenum(36161)
	GLenum_GL_GUILTY_CONTEXT_RESET_EXT                     = GLenum(33363)
	GLenum_GL_INNOCENT_CONTEXT_RESET_EXT                   = GLenum(33364)
	GLenum_GL_UNKNOWN_CONTEXT_RESET_EXT                    = GLenum(33365)
	GLenum_GL_CURRENT_QUERY                                = GLenum(34917)
	GLenum_GL_QUERY_COUNTER_BITS_EXT                       = GLenum(34916)
	GLenum_GL_QUERY_RESULT                                 = GLenum(34918)
	GLenum_GL_QUERY_RESULT_AVAILABLE                       = GLenum(34919)
	GLenum_GL_ANY_SAMPLES_PASSED                           = GLenum(35887)
	GLenum_GL_ANY_SAMPLES_PASSED_CONSERVATIVE              = GLenum(36202)
	GLenum_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN        = GLenum(35976)
	GLenum_GL_TIME_ELAPSED_EXT                             = GLenum(35007)
	GLenum_GL_TIMESTAMP_EXT                                = GLenum(36392)
	GLenum_GL_UNIFORM_BLOCK_BINDING                        = GLenum(35391)
	GLenum_GL_UNIFORM_BLOCK_DATA_SIZE                      = GLenum(35392)
	GLenum_GL_UNIFORM_BLOCK_NAME_LENGTH                    = GLenum(35393)
	GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS                = GLenum(35394)
	GLenum_GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES         = GLenum(35395)
	GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER    = GLenum(35396)
	GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER  = GLenum(35397)
	GLenum_GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER  = GLenum(35398)
)

////////////////////////////////////////////////////////////////////////////////
// enum GLbitfield
////////////////////////////////////////////////////////////////////////////////
type GLbitfield uint32

const (
	GLbitfield_GL_COLOR_BUFFER_BIT0_QCOM       = GLbitfield(1)
	GLbitfield_GL_COLOR_BUFFER_BIT1_QCOM       = GLbitfield(2)
	GLbitfield_GL_COLOR_BUFFER_BIT2_QCOM       = GLbitfield(4)
	GLbitfield_GL_COLOR_BUFFER_BIT3_QCOM       = GLbitfield(8)
	GLbitfield_GL_COLOR_BUFFER_BIT4_QCOM       = GLbitfield(16)
	GLbitfield_GL_COLOR_BUFFER_BIT5_QCOM       = GLbitfield(32)
	GLbitfield_GL_COLOR_BUFFER_BIT6_QCOM       = GLbitfield(64)
	GLbitfield_GL_COLOR_BUFFER_BIT7_QCOM       = GLbitfield(128)
	GLbitfield_GL_DEPTH_BUFFER_BIT0_QCOM       = GLbitfield(256)
	GLbitfield_GL_DEPTH_BUFFER_BIT1_QCOM       = GLbitfield(512)
	GLbitfield_GL_DEPTH_BUFFER_BIT2_QCOM       = GLbitfield(1024)
	GLbitfield_GL_DEPTH_BUFFER_BIT3_QCOM       = GLbitfield(2048)
	GLbitfield_GL_DEPTH_BUFFER_BIT4_QCOM       = GLbitfield(4096)
	GLbitfield_GL_DEPTH_BUFFER_BIT5_QCOM       = GLbitfield(8192)
	GLbitfield_GL_DEPTH_BUFFER_BIT6_QCOM       = GLbitfield(16384)
	GLbitfield_GL_DEPTH_BUFFER_BIT7_QCOM       = GLbitfield(32768)
	GLbitfield_GL_STENCIL_BUFFER_BIT0_QCOM     = GLbitfield(65536)
	GLbitfield_GL_STENCIL_BUFFER_BIT1_QCOM     = GLbitfield(131072)
	GLbitfield_GL_STENCIL_BUFFER_BIT2_QCOM     = GLbitfield(262144)
	GLbitfield_GL_STENCIL_BUFFER_BIT3_QCOM     = GLbitfield(524288)
	GLbitfield_GL_STENCIL_BUFFER_BIT4_QCOM     = GLbitfield(1048576)
	GLbitfield_GL_STENCIL_BUFFER_BIT5_QCOM     = GLbitfield(2097152)
	GLbitfield_GL_STENCIL_BUFFER_BIT6_QCOM     = GLbitfield(4194304)
	GLbitfield_GL_STENCIL_BUFFER_BIT7_QCOM     = GLbitfield(8388608)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT0_QCOM = GLbitfield(16777216)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT1_QCOM = GLbitfield(33554432)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT2_QCOM = GLbitfield(67108864)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT3_QCOM = GLbitfield(134217728)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT4_QCOM = GLbitfield(268435456)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT5_QCOM = GLbitfield(536870912)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT6_QCOM = GLbitfield(1073741824)
	GLbitfield_GL_MULTISAMPLE_BUFFER_BIT7_QCOM = GLbitfield(2147483648)
	GLbitfield_GL_COLOR_BUFFER_BIT             = GLbitfield(16384)
	GLbitfield_GL_DEPTH_BUFFER_BIT             = GLbitfield(256)
	GLbitfield_GL_STENCIL_BUFFER_BIT           = GLbitfield(1024)
	GLbitfield_GL_MAP_READ_BIT                 = GLbitfield(1)
	GLbitfield_GL_MAP_WRITE_BIT                = GLbitfield(2)
	GLbitfield_GL_MAP_INVALIDATE_RANGE_BIT     = GLbitfield(4)
	GLbitfield_GL_MAP_INVALIDATE_BUFFER_BIT    = GLbitfield(8)
	GLbitfield_GL_MAP_FLUSH_EXPLICIT_BIT       = GLbitfield(16)
	GLbitfield_GL_MAP_UNSYNCHRONIZED_BIT       = GLbitfield(32)
)

////////////////////////////////////////////////////////////////////////////////
// enum SyncCondition
////////////////////////////////////////////////////////////////////////////////
type SyncCondition uint32

const (
	SyncCondition_GL_SYNC_GPU_COMMANDS_COMPLETE = SyncCondition(37143)
)

////////////////////////////////////////////////////////////////////////////////
// enum ClientWaitSyncSignal
////////////////////////////////////////////////////////////////////////////////
type ClientWaitSyncSignal uint32

const (
	ClientWaitSyncSignal_GL_ALREADY_SIGNALED    = ClientWaitSyncSignal(37146)
	ClientWaitSyncSignal_GL_TIMEOUT_EXPIRED     = ClientWaitSyncSignal(37147)
	ClientWaitSyncSignal_GL_CONDITION_SATISFIED = ClientWaitSyncSignal(37148)
	ClientWaitSyncSignal_GL_WAIT_FAILED         = ClientWaitSyncSignal(37149)
)

////////////////////////////////////////////////////////////////////////////////
// enum SyncFlags
////////////////////////////////////////////////////////////////////////////////
type SyncFlags uint32

const (
	SyncFlags_GL_SYNC_FLUSH_COMMANDS_BIT = SyncFlags(1)
)

////////////////////////////////////////////////////////////////////////////////
// State
////////////////////////////////////////////////////////////////////////////////
type State struct {
	binary.Generate
	NextContextID ContextID
	CurrentThread ThreadID
	Contexts      ThreadIDːContextʳᵐ
	EGLContexts   EGLContextːContextʳᵐ
	GLXContexts   GLXContextːContextʳᵐ
	WGLContexts   HGLRCːContextʳᵐ
	CGLContexts   CGLContextObjːContextʳᵐ
}

func (g *State) Init() {
	g.Contexts = make(ThreadIDːContextʳᵐ)
	g.EGLContexts = make(EGLContextːContextʳᵐ)
	g.GLXContexts = make(GLXContextːContextʳᵐ)
	g.WGLContexts = make(HGLRCːContextʳᵐ)
	g.CGLContexts = make(CGLContextObjːContextʳᵐ)
}
func NewEglInitialize(Dpy memory.Pointer, Major memory.Pointer, Minor memory.Pointer, Result EGLBoolean) *EglInitialize {
	return &EglInitialize{Dpy: EGLDisplay{Pointer: Dpy}, Major: EGLintᵖ{Pointer: Major}, Minor: EGLintᵖ{Pointer: Minor}, Result: Result}
}
func NewEglCreateContext(Display memory.Pointer, Config memory.Pointer, Share_context memory.Pointer, Attrib_list memory.Pointer, Result memory.Pointer) *EglCreateContext {
	return &EglCreateContext{Display: EGLDisplay{Pointer: Display}, Config: EGLConfig{Pointer: Config}, ShareContext: EGLContext{Pointer: Share_context}, AttribList: EGLintᵖ{Pointer: Attrib_list}, Result: EGLContext{Pointer: Result}}
}
func NewEglMakeCurrent(Display memory.Pointer, Draw memory.Pointer, Read memory.Pointer, Context memory.Pointer, Result EGLBoolean) *EglMakeCurrent {
	return &EglMakeCurrent{Display: EGLDisplay{Pointer: Display}, Draw: EGLSurface{Pointer: Draw}, Read: EGLSurface{Pointer: Read}, Context: EGLContext{Pointer: Context}, Result: Result}
}
func NewEglSwapBuffers(Display memory.Pointer, Surface memory.Pointer, Result EGLBoolean) *EglSwapBuffers {
	return &EglSwapBuffers{Display: EGLDisplay{Pointer: Display}, Surface: Voidᵖ{Pointer: Surface}, Result: Result}
}
func NewEglQuerySurface(Display memory.Pointer, Surface memory.Pointer, Attribute EGLint, Value memory.Pointer, Result EGLBoolean) *EglQuerySurface {
	return &EglQuerySurface{Display: EGLDisplay{Pointer: Display}, Surface: EGLSurface{Pointer: Surface}, Attribute: Attribute, Value: EGLintᵖ{Pointer: Value}, Result: Result}
}
func NewGlXCreateContext(Dpy memory.Pointer, Vis memory.Pointer, ShareList memory.Pointer, Direct bool, Result memory.Pointer) *GlXCreateContext {
	return &GlXCreateContext{Dpy: Voidᵖ{Pointer: Dpy}, Vis: Voidᵖ{Pointer: Vis}, ShareList: GLXContext{Pointer: ShareList}, Direct: Direct, Result: GLXContext{Pointer: Result}}
}
func NewGlXCreateNewContext(Display memory.Pointer, Fbconfig memory.Pointer, Type uint32, Shared memory.Pointer, Direct bool, Result memory.Pointer) *GlXCreateNewContext {
	return &GlXCreateNewContext{Display: Voidᵖ{Pointer: Display}, Fbconfig: Voidᵖ{Pointer: Fbconfig}, Type: Type, Shared: GLXContext{Pointer: Shared}, Direct: Direct, Result: GLXContext{Pointer: Result}}
}
func NewGlXMakeContextCurrent(Display memory.Pointer, Draw memory.Pointer, Read memory.Pointer, Ctx memory.Pointer, Result Bool) *GlXMakeContextCurrent {
	return &GlXMakeContextCurrent{Display: Voidᵖ{Pointer: Display}, Draw: GLXDrawable{Pointer: Draw}, Read: GLXDrawable{Pointer: Read}, Ctx: GLXContext{Pointer: Ctx}, Result: Result}
}
func NewGlXMakeCurrent(Display memory.Pointer, Drawable memory.Pointer, Ctx memory.Pointer, Result Bool) *GlXMakeCurrent {
	return &GlXMakeCurrent{Display: Voidᵖ{Pointer: Display}, Drawable: GLXDrawable{Pointer: Drawable}, Ctx: GLXContext{Pointer: Ctx}, Result: Result}
}
func NewGlXSwapBuffers(Display memory.Pointer, Drawable memory.Pointer) *GlXSwapBuffers {
	return &GlXSwapBuffers{Display: Voidᵖ{Pointer: Display}, Drawable: GLXDrawable{Pointer: Drawable}}
}
func NewGlXQueryDrawable(Display memory.Pointer, Draw memory.Pointer, Attribute int64, Value memory.Pointer, Result int64) *GlXQueryDrawable {
	return &GlXQueryDrawable{Display: Voidᵖ{Pointer: Display}, Draw: GLXDrawable{Pointer: Draw}, Attribute: Attribute, Value: Intᵖ{Pointer: Value}, Result: Result}
}
func NewWglCreateContext(Hdc memory.Pointer, Result memory.Pointer) *WglCreateContext {
	return &WglCreateContext{Hdc: HDC{Pointer: Hdc}, Result: HGLRC{Pointer: Result}}
}
func NewWglCreateContextAttribsARB(Hdc memory.Pointer, HShareContext memory.Pointer, AttribList memory.Pointer, Result memory.Pointer) *WglCreateContextAttribsARB {
	return &WglCreateContextAttribsARB{Hdc: HDC{Pointer: Hdc}, HShareContext: HGLRC{Pointer: HShareContext}, AttribList: Intᵖ{Pointer: AttribList}, Result: HGLRC{Pointer: Result}}
}
func NewWglMakeCurrent(Hdc memory.Pointer, Hglrc memory.Pointer, Result BOOL) *WglMakeCurrent {
	return &WglMakeCurrent{Hdc: HDC{Pointer: Hdc}, Hglrc: HGLRC{Pointer: Hglrc}, Result: Result}
}
func NewWglSwapBuffers(Hdc memory.Pointer) *WglSwapBuffers {
	return &WglSwapBuffers{Hdc: HDC{Pointer: Hdc}}
}
func NewCGLCreateContext(Pix memory.Pointer, Share memory.Pointer, Ctx memory.Pointer, Result CGLError) *CGLCreateContext {
	return &CGLCreateContext{Pix: CGLPixelFormatObj{Pointer: Pix}, Share: CGLContextObj{Pointer: Share}, Ctx: CGLContextObjᵖ{Pointer: Ctx}, Result: Result}
}
func NewCGLSetCurrentContext(Ctx memory.Pointer, Result CGLError) *CGLSetCurrentContext {
	return &CGLSetCurrentContext{Ctx: CGLContextObj{Pointer: Ctx}, Result: Result}
}
func NewCGLGetSurface(Ctx memory.Pointer, Cid memory.Pointer, Wid memory.Pointer, Sid memory.Pointer, Result int64) *CGLGetSurface {
	return &CGLGetSurface{Ctx: CGLContextObj{Pointer: Ctx}, Cid: CGSConnectionIDᵖ{Pointer: Cid}, Wid: CGSWindowIDᵖ{Pointer: Wid}, Sid: CGSSurfaceIDᵖ{Pointer: Sid}, Result: Result}
}
func NewCGSGetSurfaceBounds(Cid memory.Pointer, Wid CGSWindowID, Sid CGSSurfaceID, Bounds memory.Pointer, Result int64) *CGSGetSurfaceBounds {
	return &CGSGetSurfaceBounds{Cid: CGSConnectionID{Pointer: Cid}, Wid: Wid, Sid: Sid, Bounds: F64ᵖ{Pointer: Bounds}, Result: Result}
}
func NewCGLFlushDrawable(Ctx memory.Pointer, Result CGLError) *CGLFlushDrawable {
	return &CGLFlushDrawable{Ctx: CGLContextObj{Pointer: Ctx}, Result: Result}
}
func NewGlEnableClientState(Type GLenum) *GlEnableClientState {
	return &GlEnableClientState{Type: Type}
}
func NewGlDisableClientState(Type GLenum) *GlDisableClientState {
	return &GlDisableClientState{Type: Type}
}
func NewGlGetProgramBinaryOES(Program ProgramId, Buffer_size int32, Bytes_written memory.Pointer, Binary_format memory.Pointer, Binary memory.Pointer) *GlGetProgramBinaryOES {
	return &GlGetProgramBinaryOES{Program: Program, BufferSize: Buffer_size, BytesWritten: S32ᵖ{Pointer: Bytes_written}, BinaryFormat: U32ᵖ{Pointer: Binary_format}, Binary: Voidᵖ{Pointer: Binary}}
}
func NewGlProgramBinaryOES(Program ProgramId, Binary_format uint32, Binary memory.Pointer, Binary_size int32) *GlProgramBinaryOES {
	return &GlProgramBinaryOES{Program: Program, BinaryFormat: Binary_format, Binary: Voidᵖ{Pointer: Binary}, BinarySize: Binary_size}
}
func NewGlStartTilingQCOM(X int32, Y int32, Width int32, Height int32, PreserveMask GLbitfield) *GlStartTilingQCOM {
	return &GlStartTilingQCOM{X: X, Y: Y, Width: Width, Height: Height, PreserveMask: PreserveMask}
}
func NewGlEndTilingQCOM(Preserve_mask GLbitfield) *GlEndTilingQCOM {
	return &GlEndTilingQCOM{PreserveMask: Preserve_mask}
}
func NewGlDiscardFramebufferEXT(Target GLenum, NumAttachments int32, Attachments memory.Pointer) *GlDiscardFramebufferEXT {
	return &GlDiscardFramebufferEXT{Target: Target, NumAttachments: NumAttachments, Attachments: GLenumᵖ{Pointer: Attachments}}
}
func NewGlInsertEventMarkerEXT(Length int32, Marker memory.Pointer) *GlInsertEventMarkerEXT {
	return &GlInsertEventMarkerEXT{Length: Length, Marker: Charᵖ{Pointer: Marker}}
}
func NewGlPushGroupMarkerEXT(Length int32, Marker memory.Pointer) *GlPushGroupMarkerEXT {
	return &GlPushGroupMarkerEXT{Length: Length, Marker: Charᵖ{Pointer: Marker}}
}
func NewGlPopGroupMarkerEXT() *GlPopGroupMarkerEXT {
	return &GlPopGroupMarkerEXT{}
}
func NewGlTexStorage1DEXT(Target GLenum, Levels int32, Format GLenum, Width int32) *GlTexStorage1DEXT {
	return &GlTexStorage1DEXT{Target: Target, Levels: Levels, Format: Format, Width: Width}
}
func NewGlTexStorage2DEXT(Target GLenum, Levels int32, Format GLenum, Width int32, Height int32) *GlTexStorage2DEXT {
	return &GlTexStorage2DEXT{Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height}
}
func NewGlTexStorage3DEXT(Target GLenum, Levels int32, Format GLenum, Width int32, Height int32, Depth int32) *GlTexStorage3DEXT {
	return &GlTexStorage3DEXT{Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height, Depth: Depth}
}
func NewGlTextureStorage1DEXT(Texture TextureId, Target GLenum, Levels int32, Format GLenum, Width int32) *GlTextureStorage1DEXT {
	return &GlTextureStorage1DEXT{Texture: Texture, Target: Target, Levels: Levels, Format: Format, Width: Width}
}
func NewGlTextureStorage2DEXT(Texture TextureId, Target GLenum, Levels int32, Format GLenum, Width int32, Height int32) *GlTextureStorage2DEXT {
	return &GlTextureStorage2DEXT{Texture: Texture, Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height}
}
func NewGlTextureStorage3DEXT(Texture TextureId, Target GLenum, Levels int32, Format GLenum, Width int32, Height int32, Depth int32) *GlTextureStorage3DEXT {
	return &GlTextureStorage3DEXT{Texture: Texture, Target: Target, Levels: Levels, Format: Format, Width: Width, Height: Height, Depth: Depth}
}
func NewGlGenVertexArraysOES(Count int32, Arrays memory.Pointer) *GlGenVertexArraysOES {
	return &GlGenVertexArraysOES{Count: Count, Arrays: VertexArrayIdᵖ{Pointer: Arrays}}
}
func NewGlBindVertexArrayOES(Array VertexArrayId) *GlBindVertexArrayOES {
	return &GlBindVertexArrayOES{Array: Array}
}
func NewGlDeleteVertexArraysOES(Count int32, Arrays memory.Pointer) *GlDeleteVertexArraysOES {
	return &GlDeleteVertexArraysOES{Count: Count, Arrays: VertexArrayIdᶜᵖ{Pointer: Arrays}}
}
func NewGlIsVertexArrayOES(Array VertexArrayId, Result bool) *GlIsVertexArrayOES {
	return &GlIsVertexArrayOES{Array: Array, Result: Result}
}
func NewGlEGLImageTargetTexture2DOES(Target GLenum, Image memory.Pointer) *GlEGLImageTargetTexture2DOES {
	return &GlEGLImageTargetTexture2DOES{Target: Target, Image: ImageOES{Pointer: Image}}
}
func NewGlEGLImageTargetRenderbufferStorageOES(Target GLenum, Image memory.Pointer) *GlEGLImageTargetRenderbufferStorageOES {
	return &GlEGLImageTargetRenderbufferStorageOES{Target: Target, Image: TexturePointer{Pointer: Image}}
}
func NewGlGetGraphicsResetStatusEXT(Result GLenum) *GlGetGraphicsResetStatusEXT {
	return &GlGetGraphicsResetStatusEXT{Result: Result}
}
func NewGlBindAttribLocation(Program ProgramId, Location AttributeLocation, Name string) *GlBindAttribLocation {
	return &GlBindAttribLocation{Program: Program, Location: Location, Name: Name}
}
func NewGlBlendFunc(Src_factor GLenum, Dst_factor GLenum) *GlBlendFunc {
	return &GlBlendFunc{SrcFactor: Src_factor, DstFactor: Dst_factor}
}
func NewGlBlendFuncSeparate(Src_factor_rgb GLenum, Dst_factor_rgb GLenum, Src_factor_alpha GLenum, Dst_factor_alpha GLenum) *GlBlendFuncSeparate {
	return &GlBlendFuncSeparate{SrcFactorRgb: Src_factor_rgb, DstFactorRgb: Dst_factor_rgb, SrcFactorAlpha: Src_factor_alpha, DstFactorAlpha: Dst_factor_alpha}
}
func NewGlBlendEquation(Equation GLenum) *GlBlendEquation {
	return &GlBlendEquation{Equation: Equation}
}
func NewGlBlendEquationSeparate(Rgb GLenum, Alpha GLenum) *GlBlendEquationSeparate {
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
func NewGlVertexAttribPointer(Location AttributeLocation, Size int32, Type GLenum, Normalized bool, Stride int32, Data memory.Pointer) *GlVertexAttribPointer {
	return &GlVertexAttribPointer{Location: Location, Size: Size, Type: Type, Normalized: Normalized, Stride: Stride, Data: VertexPointer{Pointer: Data}}
}
func NewGlGetActiveAttrib(Program ProgramId, Location AttributeLocation, Buffer_size int32, Buffer_bytes_written memory.Pointer, Vector_count memory.Pointer, Type memory.Pointer, Name memory.Pointer) *GlGetActiveAttrib {
	return &GlGetActiveAttrib{Program: Program, Location: Location, BufferSize: Buffer_size, BufferBytesWritten: S32ᵖ{Pointer: Buffer_bytes_written}, VectorCount: S32ᵖ{Pointer: Vector_count}, Type: GLenumᵖ{Pointer: Type}, Name: Charᵖ{Pointer: Name}}
}
func NewGlGetActiveUniform(Program ProgramId, Location int32, Buffer_size int32, Buffer_bytes_written memory.Pointer, Vector_count memory.Pointer, Type memory.Pointer, Name memory.Pointer) *GlGetActiveUniform {
	return &GlGetActiveUniform{Program: Program, Location: Location, BufferSize: Buffer_size, BufferBytesWritten: S32ᵖ{Pointer: Buffer_bytes_written}, VectorCount: S32ᵖ{Pointer: Vector_count}, Type: GLenumᵖ{Pointer: Type}, Name: Charᵖ{Pointer: Name}}
}
func NewGlGetError(Result GLenum) *GlGetError {
	return &GlGetError{Result: Result}
}
func NewGlGetProgramiv(Program ProgramId, Parameter GLenum, Value memory.Pointer) *GlGetProgramiv {
	return &GlGetProgramiv{Program: Program, Parameter: Parameter, Value: S32ᵖ{Pointer: Value}}
}
func NewGlGetShaderiv(Shader ShaderId, Parameter GLenum, Value memory.Pointer) *GlGetShaderiv {
	return &GlGetShaderiv{Shader: Shader, Parameter: Parameter, Value: S32ᵖ{Pointer: Value}}
}
func NewGlGetUniformLocation(Program ProgramId, Name string, Result UniformLocation) *GlGetUniformLocation {
	return &GlGetUniformLocation{Program: Program, Name: Name, Result: Result}
}
func NewGlGetAttribLocation(Program ProgramId, Name string, Result AttributeLocation) *GlGetAttribLocation {
	return &GlGetAttribLocation{Program: Program, Name: Name, Result: Result}
}
func NewGlPixelStorei(Parameter GLenum, Value int32) *GlPixelStorei {
	return &GlPixelStorei{Parameter: Parameter, Value: Value}
}
func NewGlTexParameteri(Target GLenum, Parameter GLenum, Value int32) *GlTexParameteri {
	return &GlTexParameteri{Target: Target, Parameter: Parameter, Value: Value}
}
func NewGlTexParameterf(Target GLenum, Parameter GLenum, Value float32) *GlTexParameterf {
	return &GlTexParameterf{Target: Target, Parameter: Parameter, Value: Value}
}
func NewGlGetTexParameteriv(Target GLenum, Parameter GLenum, Values memory.Pointer) *GlGetTexParameteriv {
	return &GlGetTexParameteriv{Target: Target, Parameter: Parameter, Values: S32ᵖ{Pointer: Values}}
}
func NewGlGetTexParameterfv(Target GLenum, Parameter GLenum, Values memory.Pointer) *GlGetTexParameterfv {
	return &GlGetTexParameterfv{Target: Target, Parameter: Parameter, Values: F32ᶜᵖ{Pointer: Values}}
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
func NewGlUniform1iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform1iv {
	return &GlUniform1iv{Location: Location, Count: Count, Values: S32ᵖ{Pointer: Values}}
}
func NewGlUniform2iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform2iv {
	return &GlUniform2iv{Location: Location, Count: Count, Values: S32ᵖ{Pointer: Values}}
}
func NewGlUniform3iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform3iv {
	return &GlUniform3iv{Location: Location, Count: Count, Values: S32ᵖ{Pointer: Values}}
}
func NewGlUniform4iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform4iv {
	return &GlUniform4iv{Location: Location, Count: Count, Values: S32ᵖ{Pointer: Values}}
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
func NewGlUniform1fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform1fv {
	return &GlUniform1fv{Location: Location, Count: Count, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlUniform2fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform2fv {
	return &GlUniform2fv{Location: Location, Count: Count, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlUniform3fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform3fv {
	return &GlUniform3fv{Location: Location, Count: Count, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlUniform4fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform4fv {
	return &GlUniform4fv{Location: Location, Count: Count, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlUniformMatrix2fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix2fv {
	return &GlUniformMatrix2fv{Location: Location, Count: Count, Transpose: Transpose, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlUniformMatrix3fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix3fv {
	return &GlUniformMatrix3fv{Location: Location, Count: Count, Transpose: Transpose, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlUniformMatrix4fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix4fv {
	return &GlUniformMatrix4fv{Location: Location, Count: Count, Transpose: Transpose, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlGetUniformfv(Program ProgramId, Location UniformLocation, Values memory.Pointer) *GlGetUniformfv {
	return &GlGetUniformfv{Program: Program, Location: Location, Values: F32ᶜᵖ{Pointer: Values}}
}
func NewGlGetUniformiv(Program ProgramId, Location UniformLocation, Values memory.Pointer) *GlGetUniformiv {
	return &GlGetUniformiv{Program: Program, Location: Location, Values: S32ᶜᵖ{Pointer: Values}}
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
	return &GlVertexAttrib1fv{Location: Location, Value: F32ᶜᵖ{Pointer: Value}}
}
func NewGlVertexAttrib2fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib2fv {
	return &GlVertexAttrib2fv{Location: Location, Value: F32ᶜᵖ{Pointer: Value}}
}
func NewGlVertexAttrib3fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib3fv {
	return &GlVertexAttrib3fv{Location: Location, Value: F32ᶜᵖ{Pointer: Value}}
}
func NewGlVertexAttrib4fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib4fv {
	return &GlVertexAttrib4fv{Location: Location, Value: F32ᶜᵖ{Pointer: Value}}
}
func NewGlGetShaderPrecisionFormat(Shader_type GLenum, Precision_type GLenum, Range memory.Pointer, Precision memory.Pointer) *GlGetShaderPrecisionFormat {
	return &GlGetShaderPrecisionFormat{ShaderType: Shader_type, PrecisionType: Precision_type, Range: S32ᵖ{Pointer: Range}, Precision: S32ᵖ{Pointer: Precision}}
}
func NewGlDepthMask(Enabled bool) *GlDepthMask {
	return &GlDepthMask{Enabled: Enabled}
}
func NewGlDepthFunc(Function GLenum) *GlDepthFunc {
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
func NewGlStencilMaskSeparate(Face GLenum, Mask uint32) *GlStencilMaskSeparate {
	return &GlStencilMaskSeparate{Face: Face, Mask: Mask}
}
func NewGlStencilFuncSeparate(Face GLenum, Function GLenum, Reference_value int32, Mask int32) *GlStencilFuncSeparate {
	return &GlStencilFuncSeparate{Face: Face, Function: Function, ReferenceValue: Reference_value, Mask: Mask}
}
func NewGlStencilOpSeparate(Face GLenum, Stencil_fail GLenum, Stencil_pass_depth_fail GLenum, Stencil_pass_depth_pass GLenum) *GlStencilOpSeparate {
	return &GlStencilOpSeparate{Face: Face, StencilFail: Stencil_fail, StencilPassDepthFail: Stencil_pass_depth_fail, StencilPassDepthPass: Stencil_pass_depth_pass}
}
func NewGlFrontFace(Orientation GLenum) *GlFrontFace {
	return &GlFrontFace{Orientation: Orientation}
}
func NewGlViewport(X int32, Y int32, Width int32, Height int32) *GlViewport {
	return &GlViewport{X: X, Y: Y, Width: Width, Height: Height}
}
func NewGlScissor(X int32, Y int32, Width int32, Height int32) *GlScissor {
	return &GlScissor{X: X, Y: Y, Width: Width, Height: Height}
}
func NewGlActiveTexture(Unit GLenum) *GlActiveTexture {
	return &GlActiveTexture{Unit: Unit}
}
func NewGlGenTextures(Count int32, Textures memory.Pointer) *GlGenTextures {
	return &GlGenTextures{Count: Count, Textures: TextureIdᵖ{Pointer: Textures}}
}
func NewGlDeleteTextures(Count int32, Textures memory.Pointer) *GlDeleteTextures {
	return &GlDeleteTextures{Count: Count, Textures: TextureIdᶜᵖ{Pointer: Textures}}
}
func NewGlIsTexture(Texture TextureId, Result bool) *GlIsTexture {
	return &GlIsTexture{Texture: Texture, Result: Result}
}
func NewGlBindTexture(Target GLenum, Texture TextureId) *GlBindTexture {
	return &GlBindTexture{Target: Target, Texture: Texture}
}
func NewGlTexImage2D(Target GLenum, Level int32, Internal_format GLenum, Width int32, Height int32, Border int32, Format GLenum, Type GLenum, Data memory.Pointer) *GlTexImage2D {
	return &GlTexImage2D{Target: Target, Level: Level, InternalFormat: Internal_format, Width: Width, Height: Height, Border: Border, Format: Format, Type: Type, Data: TexturePointer{Pointer: Data}}
}
func NewGlTexSubImage2D(Target GLenum, Level int32, Xoffset int32, Yoffset int32, Width int32, Height int32, Format GLenum, Type GLenum, Data memory.Pointer) *GlTexSubImage2D {
	return &GlTexSubImage2D{Target: Target, Level: Level, Xoffset: Xoffset, Yoffset: Yoffset, Width: Width, Height: Height, Format: Format, Type: Type, Data: TexturePointer{Pointer: Data}}
}
func NewGlCopyTexImage2D(Target GLenum, Level int32, Format GLenum, X int32, Y int32, Width int32, Height int32, Border int32) *GlCopyTexImage2D {
	return &GlCopyTexImage2D{Target: Target, Level: Level, Format: Format, X: X, Y: Y, Width: Width, Height: Height, Border: Border}
}
func NewGlCopyTexSubImage2D(Target GLenum, Level int32, Xoffset int32, Yoffset int32, X int32, Y int32, Width int32, Height int32) *GlCopyTexSubImage2D {
	return &GlCopyTexSubImage2D{Target: Target, Level: Level, Xoffset: Xoffset, Yoffset: Yoffset, X: X, Y: Y, Width: Width, Height: Height}
}
func NewGlCompressedTexImage2D(Target GLenum, Level int32, Format GLenum, Width int32, Height int32, Border int32, Image_size int32, Data memory.Pointer) *GlCompressedTexImage2D {
	return &GlCompressedTexImage2D{Target: Target, Level: Level, Format: Format, Width: Width, Height: Height, Border: Border, ImageSize: Image_size, Data: TexturePointer{Pointer: Data}}
}
func NewGlCompressedTexSubImage2D(Target GLenum, Level int32, Xoffset int32, Yoffset int32, Width int32, Height int32, Format GLenum, Image_size int32, Data memory.Pointer) *GlCompressedTexSubImage2D {
	return &GlCompressedTexSubImage2D{Target: Target, Level: Level, Xoffset: Xoffset, Yoffset: Yoffset, Width: Width, Height: Height, Format: Format, ImageSize: Image_size, Data: TexturePointer{Pointer: Data}}
}
func NewGlGenerateMipmap(Target GLenum) *GlGenerateMipmap {
	return &GlGenerateMipmap{Target: Target}
}
func NewGlReadPixels(X int32, Y int32, Width int32, Height int32, Format GLenum, Type GLenum, Data memory.Pointer) *GlReadPixels {
	return &GlReadPixels{X: X, Y: Y, Width: Width, Height: Height, Format: Format, Type: Type, Data: Voidᵖ{Pointer: Data}}
}
func NewGlGenFramebuffers(Count int32, Framebuffers memory.Pointer) *GlGenFramebuffers {
	return &GlGenFramebuffers{Count: Count, Framebuffers: FramebufferIdᵖ{Pointer: Framebuffers}}
}
func NewGlBindFramebuffer(Target GLenum, Framebuffer FramebufferId) *GlBindFramebuffer {
	return &GlBindFramebuffer{Target: Target, Framebuffer: Framebuffer}
}
func NewGlCheckFramebufferStatus(Target GLenum, Result GLenum) *GlCheckFramebufferStatus {
	return &GlCheckFramebufferStatus{Target: Target, Result: Result}
}
func NewGlDeleteFramebuffers(Count int32, Framebuffers memory.Pointer) *GlDeleteFramebuffers {
	return &GlDeleteFramebuffers{Count: Count, Framebuffers: FramebufferIdᶜᵖ{Pointer: Framebuffers}}
}
func NewGlIsFramebuffer(Framebuffer FramebufferId, Result bool) *GlIsFramebuffer {
	return &GlIsFramebuffer{Framebuffer: Framebuffer, Result: Result}
}
func NewGlGenRenderbuffers(Count int32, Renderbuffers memory.Pointer) *GlGenRenderbuffers {
	return &GlGenRenderbuffers{Count: Count, Renderbuffers: RenderbufferIdᵖ{Pointer: Renderbuffers}}
}
func NewGlBindRenderbuffer(Target GLenum, Renderbuffer RenderbufferId) *GlBindRenderbuffer {
	return &GlBindRenderbuffer{Target: Target, Renderbuffer: Renderbuffer}
}
func NewGlRenderbufferStorage(Target GLenum, Format GLenum, Width int32, Height int32) *GlRenderbufferStorage {
	return &GlRenderbufferStorage{Target: Target, Format: Format, Width: Width, Height: Height}
}
func NewGlDeleteRenderbuffers(Count int32, Renderbuffers memory.Pointer) *GlDeleteRenderbuffers {
	return &GlDeleteRenderbuffers{Count: Count, Renderbuffers: RenderbufferIdᶜᵖ{Pointer: Renderbuffers}}
}
func NewGlIsRenderbuffer(Renderbuffer RenderbufferId, Result bool) *GlIsRenderbuffer {
	return &GlIsRenderbuffer{Renderbuffer: Renderbuffer, Result: Result}
}
func NewGlGetRenderbufferParameteriv(Target GLenum, Parameter GLenum, Values memory.Pointer) *GlGetRenderbufferParameteriv {
	return &GlGetRenderbufferParameteriv{Target: Target, Parameter: Parameter, Values: S32ᵖ{Pointer: Values}}
}
func NewGlGenBuffers(Count int32, Buffers memory.Pointer) *GlGenBuffers {
	return &GlGenBuffers{Count: Count, Buffers: BufferIdᵖ{Pointer: Buffers}}
}
func NewGlBindBuffer(Target GLenum, Buffer BufferId) *GlBindBuffer {
	return &GlBindBuffer{Target: Target, Buffer: Buffer}
}
func NewGlBufferData(Target GLenum, Size int32, Data memory.Pointer, Usage GLenum) *GlBufferData {
	return &GlBufferData{Target: Target, Size: Size, Data: BufferDataPointer{Pointer: Data}, Usage: Usage}
}
func NewGlBufferSubData(Target GLenum, Offset int32, Size int32, Data memory.Pointer) *GlBufferSubData {
	return &GlBufferSubData{Target: Target, Offset: Offset, Size: Size, Data: BufferDataPointer{Pointer: Data}}
}
func NewGlDeleteBuffers(Count int32, Buffers memory.Pointer) *GlDeleteBuffers {
	return &GlDeleteBuffers{Count: Count, Buffers: BufferIdᶜᵖ{Pointer: Buffers}}
}
func NewGlIsBuffer(Buffer BufferId, Result bool) *GlIsBuffer {
	return &GlIsBuffer{Buffer: Buffer, Result: Result}
}
func NewGlGetBufferParameteriv(Target GLenum, Parameter GLenum, Value memory.Pointer) *GlGetBufferParameteriv {
	return &GlGetBufferParameteriv{Target: Target, Parameter: Parameter, Value: S32ᵖ{Pointer: Value}}
}
func NewGlCreateShader(Type GLenum, Result ShaderId) *GlCreateShader {
	return &GlCreateShader{Type: Type, Result: Result}
}
func NewGlDeleteShader(Shader ShaderId) *GlDeleteShader {
	return &GlDeleteShader{Shader: Shader}
}
func NewGlShaderSource(Shader ShaderId, Count int32, Source memory.Pointer, Length memory.Pointer) *GlShaderSource {
	return &GlShaderSource{Shader: Shader, Count: Count, Source: Charᶜᵖᶜᵖ{Pointer: Source}, Length: S32ᶜᵖ{Pointer: Length}}
}
func NewGlShaderBinary(Count int32, Shaders memory.Pointer, Binary_format uint32, Binary memory.Pointer, Binary_size int32) *GlShaderBinary {
	return &GlShaderBinary{Count: Count, Shaders: ShaderIdᶜᵖ{Pointer: Shaders}, BinaryFormat: Binary_format, Binary: Voidᶜᵖ{Pointer: Binary}, BinarySize: Binary_size}
}
func NewGlGetShaderInfoLog(Shader ShaderId, Buffer_length int32, String_length_written memory.Pointer, Info memory.Pointer) *GlGetShaderInfoLog {
	return &GlGetShaderInfoLog{Shader: Shader, BufferLength: Buffer_length, StringLengthWritten: S32ᵖ{Pointer: String_length_written}, Info: Charᵖ{Pointer: Info}}
}
func NewGlGetShaderSource(Shader ShaderId, Buffer_length int32, String_length_written memory.Pointer, Source memory.Pointer) *GlGetShaderSource {
	return &GlGetShaderSource{Shader: Shader, BufferLength: Buffer_length, StringLengthWritten: S32ᵖ{Pointer: String_length_written}, Source: Charᵖ{Pointer: Source}}
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
	return &GlGetAttachedShaders{Program: Program, BufferLength: Buffer_length, ShadersLengthWritten: S32ᵖ{Pointer: Shaders_length_written}, Shaders: ShaderIdᵖ{Pointer: Shaders}}
}
func NewGlLinkProgram(Program ProgramId) *GlLinkProgram {
	return &GlLinkProgram{Program: Program}
}
func NewGlGetProgramInfoLog(Program ProgramId, Buffer_length int32, String_length_written memory.Pointer, Info memory.Pointer) *GlGetProgramInfoLog {
	return &GlGetProgramInfoLog{Program: Program, BufferLength: Buffer_length, StringLengthWritten: S32ᵖ{Pointer: String_length_written}, Info: Charᵖ{Pointer: Info}}
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
func NewGlClear(Mask GLbitfield) *GlClear {
	return &GlClear{Mask: Mask}
}
func NewGlCullFace(Mode GLenum) *GlCullFace {
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
func NewGlHint(Target GLenum, Mode GLenum) *GlHint {
	return &GlHint{Target: Target, Mode: Mode}
}
func NewGlFramebufferRenderbuffer(Framebuffer_target GLenum, Framebuffer_attachment GLenum, Renderbuffer_target GLenum, Renderbuffer RenderbufferId) *GlFramebufferRenderbuffer {
	return &GlFramebufferRenderbuffer{FramebufferTarget: Framebuffer_target, FramebufferAttachment: Framebuffer_attachment, RenderbufferTarget: Renderbuffer_target, Renderbuffer: Renderbuffer}
}
func NewGlFramebufferTexture2D(Framebuffer_target GLenum, Framebuffer_attachment GLenum, Texture_target GLenum, Texture TextureId, Level int32) *GlFramebufferTexture2D {
	return &GlFramebufferTexture2D{FramebufferTarget: Framebuffer_target, FramebufferAttachment: Framebuffer_attachment, TextureTarget: Texture_target, Texture: Texture, Level: Level}
}
func NewGlGetFramebufferAttachmentParameteriv(Framebuffer_target GLenum, Attachment GLenum, Parameter GLenum, Value memory.Pointer) *GlGetFramebufferAttachmentParameteriv {
	return &GlGetFramebufferAttachmentParameteriv{FramebufferTarget: Framebuffer_target, Attachment: Attachment, Parameter: Parameter, Value: S32ᵖ{Pointer: Value}}
}
func NewGlDrawElements(Draw_mode GLenum, Element_count int32, Indices_type GLenum, Indices memory.Pointer) *GlDrawElements {
	return &GlDrawElements{DrawMode: Draw_mode, ElementCount: Element_count, IndicesType: Indices_type, Indices: IndicesPointer{Pointer: Indices}}
}
func NewGlDrawArrays(Draw_mode GLenum, First_index int32, Index_count int32) *GlDrawArrays {
	return &GlDrawArrays{DrawMode: Draw_mode, FirstIndex: First_index, IndexCount: Index_count}
}
func NewGlFlush() *GlFlush {
	return &GlFlush{}
}
func NewGlFinish() *GlFinish {
	return &GlFinish{}
}
func NewGlGetBooleanv(Param GLenum, Values memory.Pointer) *GlGetBooleanv {
	return &GlGetBooleanv{Param: Param, Values: Boolᵖ{Pointer: Values}}
}
func NewGlGetFloatv(Param GLenum, Values memory.Pointer) *GlGetFloatv {
	return &GlGetFloatv{Param: Param, Values: F32ᵖ{Pointer: Values}}
}
func NewGlGetIntegerv(Param GLenum, Values memory.Pointer) *GlGetIntegerv {
	return &GlGetIntegerv{Param: Param, Values: S32ᵖ{Pointer: Values}}
}
func NewGlGetString(Param GLenum, Result memory.Pointer) *GlGetString {
	return &GlGetString{Param: Param, Result: Charᶜᵖ{Pointer: Result}}
}
func NewGlEnable(Capability GLenum) *GlEnable {
	return &GlEnable{Capability: Capability}
}
func NewGlDisable(Capability GLenum) *GlDisable {
	return &GlDisable{Capability: Capability}
}
func NewGlIsEnabled(Capability GLenum, Result bool) *GlIsEnabled {
	return &GlIsEnabled{Capability: Capability, Result: Result}
}
func NewGlFenceSync(Condition SyncCondition, SyncFlags SyncFlags, Result SyncObject) *GlFenceSync {
	return &GlFenceSync{Condition: Condition, SyncFlags: SyncFlags, Result: Result}
}
func NewGlDeleteSync(Sync SyncObject) *GlDeleteSync {
	return &GlDeleteSync{Sync: Sync}
}
func NewGlWaitSync(Sync SyncObject, SyncFlags SyncFlags, Timeout uint64) *GlWaitSync {
	return &GlWaitSync{Sync: Sync, SyncFlags: SyncFlags, Timeout: Timeout}
}
func NewGlClientWaitSync(Sync SyncObject, SyncFlags SyncFlags, Timeout uint64, Result ClientWaitSyncSignal) *GlClientWaitSync {
	return &GlClientWaitSync{Sync: Sync, SyncFlags: SyncFlags, Timeout: Timeout, Result: Result}
}
func NewGlMapBufferRange(Target GLenum, Offset int32, Length int32, Access GLbitfield, Result memory.Pointer) *GlMapBufferRange {
	return &GlMapBufferRange{Target: Target, Offset: Offset, Length: Length, Access: Access, Result: Voidᵖ{Pointer: Result}}
}
func NewGlUnmapBuffer(Target GLenum) *GlUnmapBuffer {
	return &GlUnmapBuffer{Target: Target}
}
func NewGlInvalidateFramebuffer(Target GLenum, Count int32, Attachments memory.Pointer) *GlInvalidateFramebuffer {
	return &GlInvalidateFramebuffer{Target: Target, Count: Count, Attachments: GLenumᶜᵖ{Pointer: Attachments}}
}
func NewGlRenderbufferStorageMultisample(Target GLenum, Samples int32, Format GLenum, Width int32, Height int32) *GlRenderbufferStorageMultisample {
	return &GlRenderbufferStorageMultisample{Target: Target, Samples: Samples, Format: Format, Width: Width, Height: Height}
}
func NewGlBlitFramebuffer(SrcX0 int32, SrcY0 int32, SrcX1 int32, SrcY1 int32, DstX0 int32, DstY0 int32, DstX1 int32, DstY1 int32, Mask GLbitfield, Filter GLenum) *GlBlitFramebuffer {
	return &GlBlitFramebuffer{SrcX0: SrcX0, SrcY0: SrcY0, SrcX1: SrcX1, SrcY1: SrcY1, DstX0: DstX0, DstY0: DstY0, DstX1: DstX1, DstY1: DstY1, Mask: Mask, Filter: Filter}
}
func NewGlGenQueries(Count int32, Queries memory.Pointer) *GlGenQueries {
	return &GlGenQueries{Count: Count, Queries: QueryIdᵖ{Pointer: Queries}}
}
func NewGlBeginQuery(Target GLenum, Query QueryId) *GlBeginQuery {
	return &GlBeginQuery{Target: Target, Query: Query}
}
func NewGlEndQuery(Target GLenum) *GlEndQuery {
	return &GlEndQuery{Target: Target}
}
func NewGlDeleteQueries(Count int32, Queries memory.Pointer) *GlDeleteQueries {
	return &GlDeleteQueries{Count: Count, Queries: QueryIdᶜᵖ{Pointer: Queries}}
}
func NewGlIsQuery(Query QueryId, Result bool) *GlIsQuery {
	return &GlIsQuery{Query: Query, Result: Result}
}
func NewGlGetQueryiv(Target GLenum, Parameter GLenum, Value memory.Pointer) *GlGetQueryiv {
	return &GlGetQueryiv{Target: Target, Parameter: Parameter, Value: S32ᵖ{Pointer: Value}}
}
func NewGlGetQueryObjectuiv(Query QueryId, Parameter GLenum, Value memory.Pointer) *GlGetQueryObjectuiv {
	return &GlGetQueryObjectuiv{Query: Query, Parameter: Parameter, Value: U32ᵖ{Pointer: Value}}
}
func NewGlGetActiveUniformBlockName(Program ProgramId, Uniform_block_index uint32, Buffer_size int32, Buffer_bytes_written memory.Pointer, Name memory.Pointer) *GlGetActiveUniformBlockName {
	return &GlGetActiveUniformBlockName{Program: Program, UniformBlockIndex: Uniform_block_index, BufferSize: Buffer_size, BufferBytesWritten: S32ᵖ{Pointer: Buffer_bytes_written}, Name: Charᵖ{Pointer: Name}}
}
func NewGlGetActiveUniformBlockiv(Program ProgramId, Uniform_block_index uint32, Parameter_name GLenum, Parameters memory.Pointer) *GlGetActiveUniformBlockiv {
	return &GlGetActiveUniformBlockiv{Program: Program, UniformBlockIndex: Uniform_block_index, ParameterName: Parameter_name, Parameters: S32ᵖ{Pointer: Parameters}}
}
func NewGlUniformBlockBinding(Program ProgramId, Uniform_block_index uint32, Uniform_block_binding uint32) *GlUniformBlockBinding {
	return &GlUniformBlockBinding{Program: Program, UniformBlockIndex: Uniform_block_index, UniformBlockBinding: Uniform_block_binding}
}
func NewGlGetActiveUniformsiv(Program ProgramId, Uniform_count uint32, Uniform_indices memory.Pointer, Parameter_name GLenum, Parameters memory.Pointer) *GlGetActiveUniformsiv {
	return &GlGetActiveUniformsiv{Program: Program, UniformCount: Uniform_count, UniformIndices: U32ᵖ{Pointer: Uniform_indices}, ParameterName: Parameter_name, Parameters: S32ᵖ{Pointer: Parameters}}
}
func NewGlBindBufferBase(Target GLenum, Index uint32, Buffer BufferId) *GlBindBufferBase {
	return &GlBindBufferBase{Target: Target, Index: Index, Buffer: Buffer}
}
func NewGlGenVertexArrays(Count int32, Arrays memory.Pointer) *GlGenVertexArrays {
	return &GlGenVertexArrays{Count: Count, Arrays: VertexArrayIdᵖ{Pointer: Arrays}}
}
func NewGlBindVertexArray(Array VertexArrayId) *GlBindVertexArray {
	return &GlBindVertexArray{Array: Array}
}
func NewGlDeleteVertexArrays(Count uint32, Arrays memory.Pointer) *GlDeleteVertexArrays {
	return &GlDeleteVertexArrays{Count: Count, Arrays: VertexArrayIdᶜᵖ{Pointer: Arrays}}
}
func NewGlGetQueryObjecti64v(Query QueryId, Parameter GLenum, Value memory.Pointer) *GlGetQueryObjecti64v {
	return &GlGetQueryObjecti64v{Query: Query, Parameter: Parameter, Value: S64ᵖ{Pointer: Value}}
}
func NewGlGetQueryObjectui64v(Query QueryId, Parameter GLenum, Value memory.Pointer) *GlGetQueryObjectui64v {
	return &GlGetQueryObjectui64v{Query: Query, Parameter: Parameter, Value: U64ᵖ{Pointer: Value}}
}
func NewGlGenQueriesEXT(Count int32, Queries memory.Pointer) *GlGenQueriesEXT {
	return &GlGenQueriesEXT{Count: Count, Queries: QueryIdᵖ{Pointer: Queries}}
}
func NewGlBeginQueryEXT(Target GLenum, Query QueryId) *GlBeginQueryEXT {
	return &GlBeginQueryEXT{Target: Target, Query: Query}
}
func NewGlEndQueryEXT(Target GLenum) *GlEndQueryEXT {
	return &GlEndQueryEXT{Target: Target}
}
func NewGlDeleteQueriesEXT(Count int32, Queries memory.Pointer) *GlDeleteQueriesEXT {
	return &GlDeleteQueriesEXT{Count: Count, Queries: QueryIdᶜᵖ{Pointer: Queries}}
}
func NewGlIsQueryEXT(Query QueryId, Result bool) *GlIsQueryEXT {
	return &GlIsQueryEXT{Query: Query, Result: Result}
}
func NewGlQueryCounterEXT(Query QueryId, Target GLenum) *GlQueryCounterEXT {
	return &GlQueryCounterEXT{Query: Query, Target: Target}
}
func NewGlGetQueryivEXT(Target GLenum, Parameter GLenum, Value memory.Pointer) *GlGetQueryivEXT {
	return &GlGetQueryivEXT{Target: Target, Parameter: Parameter, Value: S32ᵖ{Pointer: Value}}
}
func NewGlGetQueryObjectivEXT(Query QueryId, Parameter GLenum, Value memory.Pointer) *GlGetQueryObjectivEXT {
	return &GlGetQueryObjectivEXT{Query: Query, Parameter: Parameter, Value: S32ᵖ{Pointer: Value}}
}
func NewGlGetQueryObjectuivEXT(Query QueryId, Parameter GLenum, Value memory.Pointer) *GlGetQueryObjectuivEXT {
	return &GlGetQueryObjectuivEXT{Query: Query, Parameter: Parameter, Value: U32ᵖ{Pointer: Value}}
}
func NewGlGetQueryObjecti64vEXT(Query QueryId, Parameter GLenum, Value memory.Pointer) *GlGetQueryObjecti64vEXT {
	return &GlGetQueryObjecti64vEXT{Query: Query, Parameter: Parameter, Value: S64ᵖ{Pointer: Value}}
}
func NewGlGetQueryObjectui64vEXT(Query QueryId, Parameter GLenum, Value memory.Pointer) *GlGetQueryObjectui64vEXT {
	return &GlGetQueryObjectui64vEXT{Query: Query, Parameter: Parameter, Value: U64ᵖ{Pointer: Value}}
}
func NewArchitecture(Pointer_alignment uint32, Pointer_size uint32, Integer_size uint32, Little_endian bool) *Architecture {
	return &Architecture{PointerAlignment: Pointer_alignment, PointerSize: Pointer_size, IntegerSize: Integer_size, LittleEndian: Little_endian}
}
func NewReplayCreateRenderer(Id uint32) *ReplayCreateRenderer {
	return &ReplayCreateRenderer{Id: Id}
}
func NewReplayBindRenderer(Id uint32) *ReplayBindRenderer {
	return &ReplayBindRenderer{Id: Id}
}
func NewSwitchThread(ThreadID ThreadID) *SwitchThread {
	return &SwitchThread{ThreadID: ThreadID}
}
func NewBackbufferInfo(Width int32, Height int32, Color_fmt GLenum, Depth_fmt GLenum, Stencil_fmt GLenum, ResetViewportScissor bool, PreserveBuffersOnSwap bool) *BackbufferInfo {
	return &BackbufferInfo{Width: Width, Height: Height, ColorFmt: Color_fmt, DepthFmt: Depth_fmt, StencilFmt: Stencil_fmt, ResetViewportScissor: ResetViewportScissor, PreserveBuffersOnSwap: PreserveBuffersOnSwap}
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
}
