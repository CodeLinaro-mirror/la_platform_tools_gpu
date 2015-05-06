////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

package gles

import (
	"fmt"
	"strings"

	"android.googlesource.com/platform/tools/gpu/atom"
	"android.googlesource.com/platform/tools/gpu/binary"
	"android.googlesource.com/platform/tools/gpu/gfxapi"
	"android.googlesource.com/platform/tools/gpu/memory"
)

type RenderbufferId uint32

func (c *RenderbufferId) Less(rhs RenderbufferId) bool  { return uint32(*c) < uint32(rhs) }
func (c *RenderbufferId) Equal(rhs RenderbufferId) bool { return uint32(*c) == uint32(rhs) }

type TextureId uint32

func (c *TextureId) Less(rhs TextureId) bool  { return uint32(*c) < uint32(rhs) }
func (c *TextureId) Equal(rhs TextureId) bool { return uint32(*c) == uint32(rhs) }

type FramebufferId uint32

func (c *FramebufferId) Less(rhs FramebufferId) bool  { return uint32(*c) < uint32(rhs) }
func (c *FramebufferId) Equal(rhs FramebufferId) bool { return uint32(*c) == uint32(rhs) }

type BufferId uint32

func (c *BufferId) Less(rhs BufferId) bool  { return uint32(*c) < uint32(rhs) }
func (c *BufferId) Equal(rhs BufferId) bool { return uint32(*c) == uint32(rhs) }

type ShaderId uint32

func (c *ShaderId) Less(rhs ShaderId) bool  { return uint32(*c) < uint32(rhs) }
func (c *ShaderId) Equal(rhs ShaderId) bool { return uint32(*c) == uint32(rhs) }

type ProgramId uint32

func (c *ProgramId) Less(rhs ProgramId) bool  { return uint32(*c) < uint32(rhs) }
func (c *ProgramId) Equal(rhs ProgramId) bool { return uint32(*c) == uint32(rhs) }

type VertexArrayId uint32

func (c *VertexArrayId) Less(rhs VertexArrayId) bool  { return uint32(*c) < uint32(rhs) }
func (c *VertexArrayId) Equal(rhs VertexArrayId) bool { return uint32(*c) == uint32(rhs) }

type QueryId uint32

func (c *QueryId) Less(rhs QueryId) bool  { return uint32(*c) < uint32(rhs) }
func (c *QueryId) Equal(rhs QueryId) bool { return uint32(*c) == uint32(rhs) }

type UniformLocation int32

func (c *UniformLocation) Less(rhs UniformLocation) bool  { return int32(*c) < int32(rhs) }
func (c *UniformLocation) Equal(rhs UniformLocation) bool { return int32(*c) == int32(rhs) }

type AttributeLocation uint32

func (c *AttributeLocation) Less(rhs AttributeLocation) bool  { return uint32(*c) < uint32(rhs) }
func (c *AttributeLocation) Equal(rhs AttributeLocation) bool { return uint32(*c) == uint32(rhs) }

type IndicesPointer memory.Pointer

func (c *IndicesPointer) Less(rhs IndicesPointer) bool {
	return memory.Pointer(*c) < memory.Pointer(rhs)
}
func (c *IndicesPointer) Equal(rhs IndicesPointer) bool {
	return memory.Pointer(*c) == memory.Pointer(rhs)
}

type VertexPointer memory.Pointer

func (c *VertexPointer) Less(rhs VertexPointer) bool { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *VertexPointer) Equal(rhs VertexPointer) bool {
	return memory.Pointer(*c) == memory.Pointer(rhs)
}

type TexturePointer memory.Pointer

func (c *TexturePointer) Less(rhs TexturePointer) bool {
	return memory.Pointer(*c) < memory.Pointer(rhs)
}
func (c *TexturePointer) Equal(rhs TexturePointer) bool {
	return memory.Pointer(*c) == memory.Pointer(rhs)
}

type BufferDataPointer memory.Pointer

func (c *BufferDataPointer) Less(rhs BufferDataPointer) bool {
	return memory.Pointer(*c) < memory.Pointer(rhs)
}
func (c *BufferDataPointer) Equal(rhs BufferDataPointer) bool {
	return memory.Pointer(*c) == memory.Pointer(rhs)
}

type ContextID uint32

func (c *ContextID) Less(rhs ContextID) bool  { return uint32(*c) < uint32(rhs) }
func (c *ContextID) Equal(rhs ContextID) bool { return uint32(*c) == uint32(rhs) }

type ThreadID uint32

func (c *ThreadID) Less(rhs ThreadID) bool  { return uint32(*c) < uint32(rhs) }
func (c *ThreadID) Equal(rhs ThreadID) bool { return uint32(*c) == uint32(rhs) }

type EGLBoolean int64

func (c *EGLBoolean) Less(rhs EGLBoolean) bool  { return int64(*c) < int64(rhs) }
func (c *EGLBoolean) Equal(rhs EGLBoolean) bool { return int64(*c) == int64(rhs) }

type EGLint int64

func (c *EGLint) Less(rhs EGLint) bool  { return int64(*c) < int64(rhs) }
func (c *EGLint) Equal(rhs EGLint) bool { return int64(*c) == int64(rhs) }

type EGLConfig memory.Pointer

func (c *EGLConfig) Less(rhs EGLConfig) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *EGLConfig) Equal(rhs EGLConfig) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type EGLContext memory.Pointer

func (c *EGLContext) Less(rhs EGLContext) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *EGLContext) Equal(rhs EGLContext) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type EGLDisplay memory.Pointer

func (c *EGLDisplay) Less(rhs EGLDisplay) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *EGLDisplay) Equal(rhs EGLDisplay) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type EGLSurface memory.Pointer

func (c *EGLSurface) Less(rhs EGLSurface) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *EGLSurface) Equal(rhs EGLSurface) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type GLXContext memory.Pointer

func (c *GLXContext) Less(rhs GLXContext) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *GLXContext) Equal(rhs GLXContext) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type GLXDrawable memory.Pointer

func (c *GLXDrawable) Less(rhs GLXDrawable) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *GLXDrawable) Equal(rhs GLXDrawable) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type HGLRC memory.Pointer

func (c *HGLRC) Less(rhs HGLRC) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *HGLRC) Equal(rhs HGLRC) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type HDC memory.Pointer

func (c *HDC) Less(rhs HDC) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *HDC) Equal(rhs HDC) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type BOOL int64

func (c *BOOL) Less(rhs BOOL) bool  { return int64(*c) < int64(rhs) }
func (c *BOOL) Equal(rhs BOOL) bool { return int64(*c) == int64(rhs) }

type CGLError int64

func (c *CGLError) Less(rhs CGLError) bool  { return int64(*c) < int64(rhs) }
func (c *CGLError) Equal(rhs CGLError) bool { return int64(*c) == int64(rhs) }

type CGLPixelFormatObj memory.Pointer

func (c *CGLPixelFormatObj) Less(rhs CGLPixelFormatObj) bool {
	return memory.Pointer(*c) < memory.Pointer(rhs)
}
func (c *CGLPixelFormatObj) Equal(rhs CGLPixelFormatObj) bool {
	return memory.Pointer(*c) == memory.Pointer(rhs)
}

type CGLContextObj memory.Pointer

func (c *CGLContextObj) Less(rhs CGLContextObj) bool { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *CGLContextObj) Equal(rhs CGLContextObj) bool {
	return memory.Pointer(*c) == memory.Pointer(rhs)
}

type ImageOES memory.Pointer

func (c *ImageOES) Less(rhs ImageOES) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *ImageOES) Equal(rhs ImageOES) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type BoolArray []bool

func (s BoolArray) Len() int      { return len(s) }
func (s BoolArray) Range() []bool { return s }

type BufferArray []Buffer

func (s BufferArray) Len() int        { return len(s) }
func (s BufferArray) Range() []Buffer { return s }

type BufferIdArray []BufferId

func (s BufferIdArray) Len() int          { return len(s) }
func (s BufferIdArray) Range() []BufferId { return s }

type CharBufferArray []string

func (s CharBufferArray) Len() int        { return len(s) }
func (s CharBufferArray) Range() []string { return s }

type ContextArray []Context

func (s ContextArray) Len() int         { return len(s) }
func (s ContextArray) Range() []Context { return s }

type DiscardFramebufferAttachmentArray []DiscardFramebufferAttachment

func (s DiscardFramebufferAttachmentArray) Len() int                              { return len(s) }
func (s DiscardFramebufferAttachmentArray) Range() []DiscardFramebufferAttachment { return s }

type EGLintArray []EGLint

func (s EGLintArray) Len() int        { return len(s) }
func (s EGLintArray) Range() []EGLint { return s }

type F32Array []float32

func (s F32Array) Len() int         { return len(s) }
func (s F32Array) Range() []float32 { return s }

type FramebufferArray []Framebuffer

func (s FramebufferArray) Len() int             { return len(s) }
func (s FramebufferArray) Range() []Framebuffer { return s }

type FramebufferAttachmentArray []FramebufferAttachment

func (s FramebufferAttachmentArray) Len() int                       { return len(s) }
func (s FramebufferAttachmentArray) Range() []FramebufferAttachment { return s }

type FramebufferIdArray []FramebufferId

func (s FramebufferIdArray) Len() int               { return len(s) }
func (s FramebufferIdArray) Range() []FramebufferId { return s }

type IntArray []int64

func (s IntArray) Len() int       { return len(s) }
func (s IntArray) Range() []int64 { return s }

type ProgramArray []Program

func (s ProgramArray) Len() int         { return len(s) }
func (s ProgramArray) Range() []Program { return s }

type QueryArray []Query

func (s QueryArray) Len() int       { return len(s) }
func (s QueryArray) Range() []Query { return s }

type QueryIdArray []QueryId

func (s QueryIdArray) Len() int         { return len(s) }
func (s QueryIdArray) Range() []QueryId { return s }

type RenderbufferArray []Renderbuffer

func (s RenderbufferArray) Len() int              { return len(s) }
func (s RenderbufferArray) Range() []Renderbuffer { return s }

type RenderbufferIdArray []RenderbufferId

func (s RenderbufferIdArray) Len() int                { return len(s) }
func (s RenderbufferIdArray) Range() []RenderbufferId { return s }

type S32Array []int32

func (s S32Array) Len() int       { return len(s) }
func (s S32Array) Range() []int32 { return s }

type ShaderArray []Shader

func (s ShaderArray) Len() int        { return len(s) }
func (s ShaderArray) Range() []Shader { return s }

type ShaderIdArray []ShaderId

func (s ShaderIdArray) Len() int          { return len(s) }
func (s ShaderIdArray) Range() []ShaderId { return s }

type StringArray []string

func (s StringArray) Len() int        { return len(s) }
func (s StringArray) Range() []string { return s }

type TextureArray []Texture

func (s TextureArray) Len() int         { return len(s) }
func (s TextureArray) Range() []Texture { return s }

type TextureIdArray []TextureId

func (s TextureIdArray) Len() int           { return len(s) }
func (s TextureIdArray) Range() []TextureId { return s }

type VertexArrayArray []VertexArray

func (s VertexArrayArray) Len() int             { return len(s) }
func (s VertexArrayArray) Range() []VertexArray { return s }

type VertexArrayIdArray []VertexArrayId

func (s VertexArrayIdArray) Len() int               { return len(s) }
func (s VertexArrayIdArray) Range() []VertexArrayId { return s }

type VertexAttributeArrayArray []VertexAttributeArray

func (s VertexAttributeArrayArray) Len() int                      { return len(s) }
func (s VertexAttributeArrayArray) Range() []VertexAttributeArray { return s }

type AttributeLocation_CharBufferMap map[string]AttributeLocation

func (m AttributeLocation_CharBufferMap) Get(key string) AttributeLocation {
	return m[key]
}
func (m AttributeLocation_CharBufferMap) Contains(key string) bool {
	_, ok := m[key]
	return ok
}
func (m AttributeLocation_CharBufferMap) Delete(key string) {
	delete(m, key)
}
func (m AttributeLocation_CharBufferMap) Range() []AttributeLocation {
	values := make([]AttributeLocation, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type Bool_CapabilityMap map[Capability]bool

func (m Bool_CapabilityMap) Get(key Capability) bool {
	return m[key]
}
func (m Bool_CapabilityMap) Contains(key Capability) bool {
	_, ok := m[key]
	return ok
}
func (m Bool_CapabilityMap) Delete(key Capability) {
	delete(m, key)
}
func (m Bool_CapabilityMap) Range() []bool {
	values := make([]bool, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type BufferId_BufferTargetMap map[BufferTarget]BufferId

func (m BufferId_BufferTargetMap) Get(key BufferTarget) BufferId {
	return m[key]
}
func (m BufferId_BufferTargetMap) Contains(key BufferTarget) bool {
	_, ok := m[key]
	return ok
}
func (m BufferId_BufferTargetMap) Delete(key BufferTarget) {
	delete(m, key)
}
func (m BufferId_BufferTargetMap) Range() []BufferId {
	values := make([]BufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type BufferPtr_BufferIdMap map[BufferId]*Buffer

func (m BufferPtr_BufferIdMap) Get(key BufferId) *Buffer {
	return m[key]
}
func (m BufferPtr_BufferIdMap) Contains(key BufferId) bool {
	_, ok := m[key]
	return ok
}
func (m BufferPtr_BufferIdMap) Delete(key BufferId) {
	delete(m, key)
}
func (m BufferPtr_BufferIdMap) Range() []*Buffer {
	values := make([]*Buffer, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ContextPtr_CGLContextObjMap map[CGLContextObj]*Context

func (m ContextPtr_CGLContextObjMap) Get(key CGLContextObj) *Context {
	return m[key]
}
func (m ContextPtr_CGLContextObjMap) Contains(key CGLContextObj) bool {
	_, ok := m[key]
	return ok
}
func (m ContextPtr_CGLContextObjMap) Delete(key CGLContextObj) {
	delete(m, key)
}
func (m ContextPtr_CGLContextObjMap) Range() []*Context {
	values := make([]*Context, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ContextPtr_EGLContextMap map[EGLContext]*Context

func (m ContextPtr_EGLContextMap) Get(key EGLContext) *Context {
	return m[key]
}
func (m ContextPtr_EGLContextMap) Contains(key EGLContext) bool {
	_, ok := m[key]
	return ok
}
func (m ContextPtr_EGLContextMap) Delete(key EGLContext) {
	delete(m, key)
}
func (m ContextPtr_EGLContextMap) Range() []*Context {
	values := make([]*Context, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ContextPtr_GLXContextMap map[GLXContext]*Context

func (m ContextPtr_GLXContextMap) Get(key GLXContext) *Context {
	return m[key]
}
func (m ContextPtr_GLXContextMap) Contains(key GLXContext) bool {
	_, ok := m[key]
	return ok
}
func (m ContextPtr_GLXContextMap) Delete(key GLXContext) {
	delete(m, key)
}
func (m ContextPtr_GLXContextMap) Range() []*Context {
	values := make([]*Context, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ContextPtr_HGLRCMap map[HGLRC]*Context

func (m ContextPtr_HGLRCMap) Get(key HGLRC) *Context {
	return m[key]
}
func (m ContextPtr_HGLRCMap) Contains(key HGLRC) bool {
	_, ok := m[key]
	return ok
}
func (m ContextPtr_HGLRCMap) Delete(key HGLRC) {
	delete(m, key)
}
func (m ContextPtr_HGLRCMap) Range() []*Context {
	values := make([]*Context, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ContextPtr_ThreadIDMap map[ThreadID]*Context

func (m ContextPtr_ThreadIDMap) Get(key ThreadID) *Context {
	return m[key]
}
func (m ContextPtr_ThreadIDMap) Contains(key ThreadID) bool {
	_, ok := m[key]
	return ok
}
func (m ContextPtr_ThreadIDMap) Delete(key ThreadID) {
	delete(m, key)
}
func (m ContextPtr_ThreadIDMap) Range() []*Context {
	values := make([]*Context, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type CubemapLevel_s32Map map[int32]CubemapLevel

func (m CubemapLevel_s32Map) Get(key int32) CubemapLevel {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m CubemapLevel_s32Map) Contains(key int32) bool {
	_, ok := m[key]
	return ok
}
func (m CubemapLevel_s32Map) Delete(key int32) {
	delete(m, key)
}
func (m CubemapLevel_s32Map) Range() []CubemapLevel {
	values := make([]CubemapLevel, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type FramebufferAttachmentInfo_FramebufferAttachmentMap map[FramebufferAttachment]FramebufferAttachmentInfo

func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Get(key FramebufferAttachment) FramebufferAttachmentInfo {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Contains(key FramebufferAttachment) bool {
	_, ok := m[key]
	return ok
}
func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Delete(key FramebufferAttachment) {
	delete(m, key)
}
func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Range() []FramebufferAttachmentInfo {
	values := make([]FramebufferAttachmentInfo, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type FramebufferId_FramebufferTargetMap map[FramebufferTarget]FramebufferId

func (m FramebufferId_FramebufferTargetMap) Get(key FramebufferTarget) FramebufferId {
	return m[key]
}
func (m FramebufferId_FramebufferTargetMap) Contains(key FramebufferTarget) bool {
	_, ok := m[key]
	return ok
}
func (m FramebufferId_FramebufferTargetMap) Delete(key FramebufferTarget) {
	delete(m, key)
}
func (m FramebufferId_FramebufferTargetMap) Range() []FramebufferId {
	values := make([]FramebufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type FramebufferPtr_FramebufferIdMap map[FramebufferId]*Framebuffer

func (m FramebufferPtr_FramebufferIdMap) Get(key FramebufferId) *Framebuffer {
	return m[key]
}
func (m FramebufferPtr_FramebufferIdMap) Contains(key FramebufferId) bool {
	_, ok := m[key]
	return ok
}
func (m FramebufferPtr_FramebufferIdMap) Delete(key FramebufferId) {
	delete(m, key)
}
func (m FramebufferPtr_FramebufferIdMap) Range() []*Framebuffer {
	values := make([]*Framebuffer, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type Image_CubeMapImageTargetMap map[CubeMapImageTarget]Image

func (m Image_CubeMapImageTargetMap) Get(key CubeMapImageTarget) Image {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m Image_CubeMapImageTargetMap) Contains(key CubeMapImageTarget) bool {
	_, ok := m[key]
	return ok
}
func (m Image_CubeMapImageTargetMap) Delete(key CubeMapImageTarget) {
	delete(m, key)
}
func (m Image_CubeMapImageTargetMap) Range() []Image {
	values := make([]Image, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type Image_s32Map map[int32]Image

func (m Image_s32Map) Get(key int32) Image {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m Image_s32Map) Contains(key int32) bool {
	_, ok := m[key]
	return ok
}
func (m Image_s32Map) Delete(key int32) {
	delete(m, key)
}
func (m Image_s32Map) Range() []Image {
	values := make([]Image, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ProgramPtr_ProgramIdMap map[ProgramId]*Program

func (m ProgramPtr_ProgramIdMap) Get(key ProgramId) *Program {
	return m[key]
}
func (m ProgramPtr_ProgramIdMap) Contains(key ProgramId) bool {
	_, ok := m[key]
	return ok
}
func (m ProgramPtr_ProgramIdMap) Delete(key ProgramId) {
	delete(m, key)
}
func (m ProgramPtr_ProgramIdMap) Range() []*Program {
	values := make([]*Program, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type QueryPtr_QueryIdMap map[QueryId]*Query

func (m QueryPtr_QueryIdMap) Get(key QueryId) *Query {
	return m[key]
}
func (m QueryPtr_QueryIdMap) Contains(key QueryId) bool {
	_, ok := m[key]
	return ok
}
func (m QueryPtr_QueryIdMap) Delete(key QueryId) {
	delete(m, key)
}
func (m QueryPtr_QueryIdMap) Range() []*Query {
	values := make([]*Query, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type RenderbufferId_RenderbufferTargetMap map[RenderbufferTarget]RenderbufferId

func (m RenderbufferId_RenderbufferTargetMap) Get(key RenderbufferTarget) RenderbufferId {
	return m[key]
}
func (m RenderbufferId_RenderbufferTargetMap) Contains(key RenderbufferTarget) bool {
	_, ok := m[key]
	return ok
}
func (m RenderbufferId_RenderbufferTargetMap) Delete(key RenderbufferTarget) {
	delete(m, key)
}
func (m RenderbufferId_RenderbufferTargetMap) Range() []RenderbufferId {
	values := make([]RenderbufferId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type RenderbufferPtr_RenderbufferIdMap map[RenderbufferId]*Renderbuffer

func (m RenderbufferPtr_RenderbufferIdMap) Get(key RenderbufferId) *Renderbuffer {
	return m[key]
}
func (m RenderbufferPtr_RenderbufferIdMap) Contains(key RenderbufferId) bool {
	_, ok := m[key]
	return ok
}
func (m RenderbufferPtr_RenderbufferIdMap) Delete(key RenderbufferId) {
	delete(m, key)
}
func (m RenderbufferPtr_RenderbufferIdMap) Range() []*Renderbuffer {
	values := make([]*Renderbuffer, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type S32_PixelStoreParameterMap map[PixelStoreParameter]int32

func (m S32_PixelStoreParameterMap) Get(key PixelStoreParameter) int32 {
	return m[key]
}
func (m S32_PixelStoreParameterMap) Contains(key PixelStoreParameter) bool {
	_, ok := m[key]
	return ok
}
func (m S32_PixelStoreParameterMap) Delete(key PixelStoreParameter) {
	delete(m, key)
}
func (m S32_PixelStoreParameterMap) Range() []int32 {
	values := make([]int32, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ShaderId_ShaderTypeMap map[ShaderType]ShaderId

func (m ShaderId_ShaderTypeMap) Get(key ShaderType) ShaderId {
	return m[key]
}
func (m ShaderId_ShaderTypeMap) Contains(key ShaderType) bool {
	_, ok := m[key]
	return ok
}
func (m ShaderId_ShaderTypeMap) Delete(key ShaderType) {
	delete(m, key)
}
func (m ShaderId_ShaderTypeMap) Range() []ShaderId {
	values := make([]ShaderId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type ShaderPtr_ShaderIdMap map[ShaderId]*Shader

func (m ShaderPtr_ShaderIdMap) Get(key ShaderId) *Shader {
	return m[key]
}
func (m ShaderPtr_ShaderIdMap) Contains(key ShaderId) bool {
	_, ok := m[key]
	return ok
}
func (m ShaderPtr_ShaderIdMap) Delete(key ShaderId) {
	delete(m, key)
}
func (m ShaderPtr_ShaderIdMap) Range() []*Shader {
	values := make([]*Shader, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type TextureId_TextureTargetMap map[TextureTarget]TextureId

func (m TextureId_TextureTargetMap) Get(key TextureTarget) TextureId {
	return m[key]
}
func (m TextureId_TextureTargetMap) Contains(key TextureTarget) bool {
	_, ok := m[key]
	return ok
}
func (m TextureId_TextureTargetMap) Delete(key TextureTarget) {
	delete(m, key)
}
func (m TextureId_TextureTargetMap) Range() []TextureId {
	values := make([]TextureId, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type TextureId_TextureTargetMap_TextureUnitMap map[TextureUnit]TextureId_TextureTargetMap

func (m TextureId_TextureTargetMap_TextureUnitMap) Get(key TextureUnit) TextureId_TextureTargetMap {
	v, ok := m[key]
	if !ok {
		v = make(TextureId_TextureTargetMap)
	}
	return v
}
func (m TextureId_TextureTargetMap_TextureUnitMap) Contains(key TextureUnit) bool {
	_, ok := m[key]
	return ok
}
func (m TextureId_TextureTargetMap_TextureUnitMap) Delete(key TextureUnit) {
	delete(m, key)
}
func (m TextureId_TextureTargetMap_TextureUnitMap) Range() []TextureId_TextureTargetMap {
	values := make([]TextureId_TextureTargetMap, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type TexturePtr_TextureIdMap map[TextureId]*Texture

func (m TexturePtr_TextureIdMap) Get(key TextureId) *Texture {
	return m[key]
}
func (m TexturePtr_TextureIdMap) Contains(key TextureId) bool {
	_, ok := m[key]
	return ok
}
func (m TexturePtr_TextureIdMap) Delete(key TextureId) {
	delete(m, key)
}
func (m TexturePtr_TextureIdMap) Range() []*Texture {
	values := make([]*Texture, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type U32_FaceModeMap map[FaceMode]uint32

func (m U32_FaceModeMap) Get(key FaceMode) uint32 {
	return m[key]
}
func (m U32_FaceModeMap) Contains(key FaceMode) bool {
	_, ok := m[key]
	return ok
}
func (m U32_FaceModeMap) Delete(key FaceMode) {
	delete(m, key)
}
func (m U32_FaceModeMap) Range() []uint32 {
	values := make([]uint32, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type Uniform_UniformLocationMap map[UniformLocation]Uniform

func (m Uniform_UniformLocationMap) Get(key UniformLocation) Uniform {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m Uniform_UniformLocationMap) Contains(key UniformLocation) bool {
	_, ok := m[key]
	return ok
}
func (m Uniform_UniformLocationMap) Delete(key UniformLocation) {
	delete(m, key)
}
func (m Uniform_UniformLocationMap) Range() []Uniform {
	values := make([]Uniform, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type VertexArrayPtr_VertexArrayIdMap map[VertexArrayId]*VertexArray

func (m VertexArrayPtr_VertexArrayIdMap) Get(key VertexArrayId) *VertexArray {
	return m[key]
}
func (m VertexArrayPtr_VertexArrayIdMap) Contains(key VertexArrayId) bool {
	_, ok := m[key]
	return ok
}
func (m VertexArrayPtr_VertexArrayIdMap) Delete(key VertexArrayId) {
	delete(m, key)
}
func (m VertexArrayPtr_VertexArrayIdMap) Range() []*VertexArray {
	values := make([]*VertexArray, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type VertexAttributeArrayPtr_AttributeLocationMap map[AttributeLocation]*VertexAttributeArray

func (m VertexAttributeArrayPtr_AttributeLocationMap) Get(key AttributeLocation) *VertexAttributeArray {
	return m[key]
}
func (m VertexAttributeArrayPtr_AttributeLocationMap) Contains(key AttributeLocation) bool {
	_, ok := m[key]
	return ok
}
func (m VertexAttributeArrayPtr_AttributeLocationMap) Delete(key AttributeLocation) {
	delete(m, key)
}
func (m VertexAttributeArrayPtr_AttributeLocationMap) Range() []*VertexAttributeArray {
	values := make([]*VertexAttributeArray, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type VertexAttribute_s32Map map[int32]VertexAttribute

func (m VertexAttribute_s32Map) Get(key int32) VertexAttribute {
	v, ok := m[key]
	if !ok {
		v.Init()
	}
	return v
}
func (m VertexAttribute_s32Map) Contains(key int32) bool {
	_, ok := m[key]
	return ok
}
func (m VertexAttribute_s32Map) Delete(key int32) {
	delete(m, key)
}
func (m VertexAttribute_s32Map) Range() []VertexAttribute {
	values := make([]VertexAttribute, 0, len(m))
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
	Id uint32
}

func (c *ReplayCreateRenderer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "replayCreateRenderer(",
		fmt.Sprintf("id:%v", c.Id),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *ReplayCreateRenderer) API() gfxapi.API {
	return api{}
}
func (c *ReplayCreateRenderer) TypeID() atom.TypeID {
	return 0
}
func (c *ReplayCreateRenderer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// ReplayBindRenderer
////////////////////////////////////////////////////////////////////////////////
type ReplayBindRenderer struct {
	binary.Generate
	Id uint32
}

func (c *ReplayBindRenderer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "replayBindRenderer(",
		fmt.Sprintf("id:%v", c.Id),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *ReplayBindRenderer) API() gfxapi.API {
	return api{}
}
func (c *ReplayBindRenderer) TypeID() atom.TypeID {
	return 1
}
func (c *ReplayBindRenderer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// BackbufferInfo
////////////////////////////////////////////////////////////////////////////////
type BackbufferInfo struct {
	binary.Generate
	Width                int32
	Height               int32
	ColorFmt             RenderbufferFormat
	DepthFmt             RenderbufferFormat
	StencilFmt           RenderbufferFormat
	ResetViewportScissor bool
}

func (c *BackbufferInfo) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "backbufferInfo(",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		c.ColorFmt.String(),
		", ",
		c.DepthFmt.String(),
		", ",
		c.StencilFmt.String(),
		", ",
		fmt.Sprintf("resetViewportScissor:%v", c.ResetViewportScissor),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *BackbufferInfo) API() gfxapi.API {
	return api{}
}
func (c *BackbufferInfo) TypeID() atom.TypeID {
	return 2
}
func (c *BackbufferInfo) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// StartTimer
////////////////////////////////////////////////////////////////////////////////
type StartTimer struct {
	binary.Generate
	Index uint8
}

func (c *StartTimer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "startTimer(",
		fmt.Sprintf("index:%v", c.Index),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *StartTimer) API() gfxapi.API {
	return api{}
}
func (c *StartTimer) TypeID() atom.TypeID {
	return 3
}
func (c *StartTimer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// StopTimer
////////////////////////////////////////////////////////////////////////////////
type StopTimer struct {
	binary.Generate
	Index  uint8
	Result uint64
}

func (c *StopTimer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "stopTimer(",
		fmt.Sprintf("index:%v", c.Index),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *StopTimer) API() gfxapi.API {
	return api{}
}
func (c *StopTimer) TypeID() atom.TypeID {
	return 4
}
func (c *StopTimer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// FlushPostBuffer
////////////////////////////////////////////////////////////////////////////////
type FlushPostBuffer struct {
	binary.Generate
}

func (c *FlushPostBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "flushPostBuffer(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *FlushPostBuffer) API() gfxapi.API {
	return api{}
}
func (c *FlushPostBuffer) TypeID() atom.TypeID {
	return 5
}
func (c *FlushPostBuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// EglInitialize
////////////////////////////////////////////////////////////////////////////////
type EglInitialize struct {
	binary.Generate
	Dpy    EGLDisplay
	Major  EGLint
	Minor  EGLint
	Result EGLBoolean
}

func (c *EglInitialize) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglInitialize(",
		fmt.Sprintf("dpy:%v", c.Dpy),
		", ",
		fmt.Sprintf("major:%v", c.Major),
		", ",
		fmt.Sprintf("minor:%v", c.Minor),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *EglInitialize) API() gfxapi.API {
	return api{}
}
func (c *EglInitialize) TypeID() atom.TypeID {
	return 6
}
func (c *EglInitialize) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// EglCreateContext
////////////////////////////////////////////////////////////////////////////////
type EglCreateContext struct {
	binary.Generate
	Display      EGLDisplay
	Config       EGLConfig
	ShareContext EGLContext
	AttribList   EGLintArray
	Result       EGLContext
}

func (c *EglCreateContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglCreateContext(",
		fmt.Sprintf("display:%v", c.Display),
		", ",
		fmt.Sprintf("config:%v", c.Config),
		", ",
		fmt.Sprintf("share_context:%v", c.ShareContext),
		", ",
		fmt.Sprintf("%v", c.AttribList),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *EglCreateContext) API() gfxapi.API {
	return api{}
}
func (c *EglCreateContext) TypeID() atom.TypeID {
	return 7
}
func (c *EglCreateContext) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// EglMakeCurrent
////////////////////////////////////////////////////////////////////////////////
type EglMakeCurrent struct {
	binary.Generate
	Display EGLDisplay
	Draw    EGLSurface
	Read    EGLSurface
	Context EGLContext
	Result  EGLBoolean
}

func (c *EglMakeCurrent) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglMakeCurrent(",
		fmt.Sprintf("display:%v", c.Display),
		", ",
		fmt.Sprintf("draw:%v", c.Draw),
		", ",
		fmt.Sprintf("read:%v", c.Read),
		", ",
		fmt.Sprintf("context:%v", c.Context),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *EglMakeCurrent) API() gfxapi.API {
	return api{}
}
func (c *EglMakeCurrent) TypeID() atom.TypeID {
	return 8
}
func (c *EglMakeCurrent) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// EglSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type EglSwapBuffers struct {
	binary.Generate
	Display EGLDisplay
	Surface memory.Pointer
	Result  EGLBoolean
}

func (c *EglSwapBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglSwapBuffers(",
		fmt.Sprintf("display:%v", c.Display),
		", ",
		fmt.Sprintf("0x%x", c.Surface),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *EglSwapBuffers) API() gfxapi.API {
	return api{}
}
func (c *EglSwapBuffers) TypeID() atom.TypeID {
	return 9
}
func (c *EglSwapBuffers) Flags() atom.Flags {
	return 0 | atom.EndOfFrame
}

////////////////////////////////////////////////////////////////////////////////
// EglQuerySurface
////////////////////////////////////////////////////////////////////////////////
type EglQuerySurface struct {
	binary.Generate
	Display   EGLDisplay
	Surface   EGLSurface
	Attribute EGLint
	Value     EGLint
	Result    EGLBoolean
}

func (c *EglQuerySurface) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglQuerySurface(",
		fmt.Sprintf("display:%v", c.Display),
		", ",
		fmt.Sprintf("surface:%v", c.Surface),
		", ",
		fmt.Sprintf("attribute:%v", c.Attribute),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *EglQuerySurface) API() gfxapi.API {
	return api{}
}
func (c *EglQuerySurface) TypeID() atom.TypeID {
	return 10
}
func (c *EglQuerySurface) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlXCreateContext
////////////////////////////////////////////////////////////////////////////////
type GlXCreateContext struct {
	binary.Generate
	Dpy       memory.Pointer
	Vis       memory.Pointer
	ShareList GLXContext
	Direct    bool
	Result    GLXContext
}

func (c *GlXCreateContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glXCreateContext(",
		fmt.Sprintf("0x%x", c.Dpy),
		", ",
		fmt.Sprintf("0x%x", c.Vis),
		", ",
		fmt.Sprintf("shareList:%v", c.ShareList),
		", ",
		fmt.Sprintf("direct:%v", c.Direct),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlXCreateContext) API() gfxapi.API {
	return api{}
}
func (c *GlXCreateContext) TypeID() atom.TypeID {
	return 11
}
func (c *GlXCreateContext) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlXCreateNewContext
////////////////////////////////////////////////////////////////////////////////
type GlXCreateNewContext struct {
	binary.Generate
	Display  memory.Pointer
	Fbconfig memory.Pointer
	Type     uint32
	Shared   GLXContext
	Direct   bool
	Result   GLXContext
}

func (c *GlXCreateNewContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glXCreateNewContext(",
		fmt.Sprintf("0x%x", c.Display),
		", ",
		fmt.Sprintf("0x%x", c.Fbconfig),
		", ",
		fmt.Sprintf("type:%v", c.Type),
		", ",
		fmt.Sprintf("shared:%v", c.Shared),
		", ",
		fmt.Sprintf("direct:%v", c.Direct),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlXCreateNewContext) API() gfxapi.API {
	return api{}
}
func (c *GlXCreateNewContext) TypeID() atom.TypeID {
	return 12
}
func (c *GlXCreateNewContext) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlXMakeContextCurrent
////////////////////////////////////////////////////////////////////////////////
type GlXMakeContextCurrent struct {
	binary.Generate
	Display memory.Pointer
	Draw    GLXDrawable
	Read    GLXDrawable
	Ctx     GLXContext
}

func (c *GlXMakeContextCurrent) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glXMakeContextCurrent(",
		fmt.Sprintf("0x%x", c.Display),
		", ",
		fmt.Sprintf("draw:%v", c.Draw),
		", ",
		fmt.Sprintf("read:%v", c.Read),
		", ",
		fmt.Sprintf("ctx:%v", c.Ctx),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlXMakeContextCurrent) API() gfxapi.API {
	return api{}
}
func (c *GlXMakeContextCurrent) TypeID() atom.TypeID {
	return 13
}
func (c *GlXMakeContextCurrent) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlXSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type GlXSwapBuffers struct {
	binary.Generate
	Display  memory.Pointer
	Drawable GLXDrawable
}

func (c *GlXSwapBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glXSwapBuffers(",
		fmt.Sprintf("0x%x", c.Display),
		", ",
		fmt.Sprintf("drawable:%v", c.Drawable),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlXSwapBuffers) API() gfxapi.API {
	return api{}
}
func (c *GlXSwapBuffers) TypeID() atom.TypeID {
	return 14
}
func (c *GlXSwapBuffers) Flags() atom.Flags {
	return 0 | atom.EndOfFrame
}

////////////////////////////////////////////////////////////////////////////////
// WglCreateContext
////////////////////////////////////////////////////////////////////////////////
type WglCreateContext struct {
	binary.Generate
	Hdc    HDC
	Result HGLRC
}

func (c *WglCreateContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "wglCreateContext(",
		fmt.Sprintf("hdc:%v", c.Hdc),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *WglCreateContext) API() gfxapi.API {
	return api{}
}
func (c *WglCreateContext) TypeID() atom.TypeID {
	return 15
}
func (c *WglCreateContext) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// WglCreateContextAttribsARB
////////////////////////////////////////////////////////////////////////////////
type WglCreateContextAttribsARB struct {
	binary.Generate
	Hdc           HDC
	HShareContext HGLRC
	AttribList    IntArray
	Result        HGLRC
}

func (c *WglCreateContextAttribsARB) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "wglCreateContextAttribsARB(",
		fmt.Sprintf("hdc:%v", c.Hdc),
		", ",
		fmt.Sprintf("hShareContext:%v", c.HShareContext),
		", ",
		fmt.Sprintf("%v", c.AttribList),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *WglCreateContextAttribsARB) API() gfxapi.API {
	return api{}
}
func (c *WglCreateContextAttribsARB) TypeID() atom.TypeID {
	return 16
}
func (c *WglCreateContextAttribsARB) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// WglMakeCurrent
////////////////////////////////////////////////////////////////////////////////
type WglMakeCurrent struct {
	binary.Generate
	Hdc    HDC
	Hglrc  HGLRC
	Result BOOL
}

func (c *WglMakeCurrent) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "wglMakeCurrent(",
		fmt.Sprintf("hdc:%v", c.Hdc),
		", ",
		fmt.Sprintf("hglrc:%v", c.Hglrc),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *WglMakeCurrent) API() gfxapi.API {
	return api{}
}
func (c *WglMakeCurrent) TypeID() atom.TypeID {
	return 17
}
func (c *WglMakeCurrent) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// WglSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type WglSwapBuffers struct {
	binary.Generate
	Hdc HDC
}

func (c *WglSwapBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "wglSwapBuffers(",
		fmt.Sprintf("hdc:%v", c.Hdc),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *WglSwapBuffers) API() gfxapi.API {
	return api{}
}
func (c *WglSwapBuffers) TypeID() atom.TypeID {
	return 18
}
func (c *WglSwapBuffers) Flags() atom.Flags {
	return 0 | atom.EndOfFrame
}

////////////////////////////////////////////////////////////////////////////////
// CGLCreateContext
////////////////////////////////////////////////////////////////////////////////
type CGLCreateContext struct {
	binary.Generate
	Pix    CGLPixelFormatObj
	Share  CGLContextObj
	Ctx    CGLContextObj
	Result CGLError
}

func (c *CGLCreateContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "CGLCreateContext(",
		fmt.Sprintf("pix:%v", c.Pix),
		", ",
		fmt.Sprintf("share:%v", c.Share),
		", ",
		fmt.Sprintf("ctx:%v", c.Ctx),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CGLCreateContext) API() gfxapi.API {
	return api{}
}
func (c *CGLCreateContext) TypeID() atom.TypeID {
	return 19
}
func (c *CGLCreateContext) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// CGLSetCurrentContext
////////////////////////////////////////////////////////////////////////////////
type CGLSetCurrentContext struct {
	binary.Generate
	Ctx    CGLContextObj
	Result CGLError
}

func (c *CGLSetCurrentContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "CGLSetCurrentContext(",
		fmt.Sprintf("ctx:%v", c.Ctx),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *CGLSetCurrentContext) API() gfxapi.API {
	return api{}
}
func (c *CGLSetCurrentContext) TypeID() atom.TypeID {
	return 20
}
func (c *CGLSetCurrentContext) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEnableClientState
////////////////////////////////////////////////////////////////////////////////
type GlEnableClientState struct {
	binary.Generate
	Type ArrayType
}

func (c *GlEnableClientState) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEnableClientState(",
		c.Type.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEnableClientState) API() gfxapi.API {
	return api{}
}
func (c *GlEnableClientState) TypeID() atom.TypeID {
	return 21
}
func (c *GlEnableClientState) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDisableClientState
////////////////////////////////////////////////////////////////////////////////
type GlDisableClientState struct {
	binary.Generate
	Type ArrayType
}

func (c *GlDisableClientState) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDisableClientState(",
		c.Type.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDisableClientState) API() gfxapi.API {
	return api{}
}
func (c *GlDisableClientState) TypeID() atom.TypeID {
	return 22
}
func (c *GlDisableClientState) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramBinaryOES
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramBinaryOES struct {
	binary.Generate
	Program      ProgramId
	BufferSize   int32
	BytesWritten int32
	BinaryFormat uint32
	Binary       memory.Pointer
}

func (c *GlGetProgramBinaryOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetProgramBinaryOES(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("buffer_size:%v", c.BufferSize),
		", ",
		fmt.Sprintf("bytes_written:%v", c.BytesWritten),
		", ",
		fmt.Sprintf("binary_format:%v", c.BinaryFormat),
		", ",
		fmt.Sprintf("0x%x", c.Binary),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetProgramBinaryOES) API() gfxapi.API {
	return api{}
}
func (c *GlGetProgramBinaryOES) TypeID() atom.TypeID {
	return 23
}
func (c *GlGetProgramBinaryOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlProgramBinaryOES
////////////////////////////////////////////////////////////////////////////////
type GlProgramBinaryOES struct {
	binary.Generate
	Program      ProgramId
	BinaryFormat uint32
	Binary       memory.Pointer
	BinarySize   int32
}

func (c *GlProgramBinaryOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glProgramBinaryOES(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("binary_format:%v", c.BinaryFormat),
		", ",
		fmt.Sprintf("0x%x", c.Binary),
		", ",
		fmt.Sprintf("binary_size:%v", c.BinarySize),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlProgramBinaryOES) API() gfxapi.API {
	return api{}
}
func (c *GlProgramBinaryOES) TypeID() atom.TypeID {
	return 24
}
func (c *GlProgramBinaryOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlStartTilingQCOM
////////////////////////////////////////////////////////////////////////////////
type GlStartTilingQCOM struct {
	binary.Generate
	X            int32
	Y            int32
	Width        int32
	Height       int32
	PreserveMask TilePreserveMaskQCOM
}

func (c *GlStartTilingQCOM) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStartTilingQCOM(",
		fmt.Sprintf("x:%v", c.X),
		", ",
		fmt.Sprintf("y:%v", c.Y),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		c.PreserveMask.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStartTilingQCOM) API() gfxapi.API {
	return api{}
}
func (c *GlStartTilingQCOM) TypeID() atom.TypeID {
	return 25
}
func (c *GlStartTilingQCOM) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEndTilingQCOM
////////////////////////////////////////////////////////////////////////////////
type GlEndTilingQCOM struct {
	binary.Generate
	PreserveMask TilePreserveMaskQCOM
}

func (c *GlEndTilingQCOM) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEndTilingQCOM(",
		c.PreserveMask.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEndTilingQCOM) API() gfxapi.API {
	return api{}
}
func (c *GlEndTilingQCOM) TypeID() atom.TypeID {
	return 26
}
func (c *GlEndTilingQCOM) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDiscardFramebufferEXT
////////////////////////////////////////////////////////////////////////////////
type GlDiscardFramebufferEXT struct {
	binary.Generate
	Target         FramebufferTarget
	NumAttachments int32
	Attachments    DiscardFramebufferAttachmentArray
}

func (c *GlDiscardFramebufferEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDiscardFramebufferEXT(",
		c.Target.String(),
		", ",
		fmt.Sprintf("numAttachments:%v", c.NumAttachments),
		", ",
		fmt.Sprintf("%v", c.Attachments),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDiscardFramebufferEXT) API() gfxapi.API {
	return api{}
}
func (c *GlDiscardFramebufferEXT) TypeID() atom.TypeID {
	return 27
}
func (c *GlDiscardFramebufferEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlInsertEventMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlInsertEventMarkerEXT struct {
	binary.Generate
	Length int32
	Marker string
}

func (c *GlInsertEventMarkerEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glInsertEventMarkerEXT(",
		fmt.Sprintf("length:%v", c.Length),
		", ",
		fmt.Sprintf("marker:%v", c.Marker),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlInsertEventMarkerEXT) API() gfxapi.API {
	return api{}
}
func (c *GlInsertEventMarkerEXT) TypeID() atom.TypeID {
	return 28
}
func (c *GlInsertEventMarkerEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlPushGroupMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlPushGroupMarkerEXT struct {
	binary.Generate
	Length int32
	Marker string
}

func (c *GlPushGroupMarkerEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPushGroupMarkerEXT(",
		fmt.Sprintf("length:%v", c.Length),
		", ",
		fmt.Sprintf("marker:%v", c.Marker),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPushGroupMarkerEXT) API() gfxapi.API {
	return api{}
}
func (c *GlPushGroupMarkerEXT) TypeID() atom.TypeID {
	return 29
}
func (c *GlPushGroupMarkerEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlPopGroupMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlPopGroupMarkerEXT struct {
	binary.Generate
}

func (c *GlPopGroupMarkerEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPopGroupMarkerEXT(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPopGroupMarkerEXT) API() gfxapi.API {
	return api{}
}
func (c *GlPopGroupMarkerEXT) TypeID() atom.TypeID {
	return 30
}
func (c *GlPopGroupMarkerEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage1DEXT struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
}

func (c *GlTexStorage1DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexStorage1DEXT(",
		c.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.Levels),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexStorage1DEXT) API() gfxapi.API {
	return api{}
}
func (c *GlTexStorage1DEXT) TypeID() atom.TypeID {
	return 31
}
func (c *GlTexStorage1DEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage2DEXT struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
	Height int32
}

func (c *GlTexStorage2DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexStorage2DEXT(",
		c.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.Levels),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexStorage2DEXT) API() gfxapi.API {
	return api{}
}
func (c *GlTexStorage2DEXT) TypeID() atom.TypeID {
	return 32
}
func (c *GlTexStorage2DEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage3DEXT struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
	Height int32
	Depth  int32
}

func (c *GlTexStorage3DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexStorage3DEXT(",
		c.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.Levels),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		fmt.Sprintf("depth:%v", c.Depth),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexStorage3DEXT) API() gfxapi.API {
	return api{}
}
func (c *GlTexStorage3DEXT) TypeID() atom.TypeID {
	return 33
}
func (c *GlTexStorage3DEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage1DEXT struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
}

func (c *GlTextureStorage1DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTextureStorage1DEXT(",
		fmt.Sprintf("texture:%v", c.Texture),
		", ",
		c.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.Levels),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTextureStorage1DEXT) API() gfxapi.API {
	return api{}
}
func (c *GlTextureStorage1DEXT) TypeID() atom.TypeID {
	return 34
}
func (c *GlTextureStorage1DEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage2DEXT struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
	Height  int32
}

func (c *GlTextureStorage2DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTextureStorage2DEXT(",
		fmt.Sprintf("texture:%v", c.Texture),
		", ",
		c.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.Levels),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTextureStorage2DEXT) API() gfxapi.API {
	return api{}
}
func (c *GlTextureStorage2DEXT) TypeID() atom.TypeID {
	return 35
}
func (c *GlTextureStorage2DEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage3DEXT struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
	Height  int32
	Depth   int32
}

func (c *GlTextureStorage3DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTextureStorage3DEXT(",
		fmt.Sprintf("texture:%v", c.Texture),
		", ",
		c.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.Levels),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		fmt.Sprintf("depth:%v", c.Depth),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTextureStorage3DEXT) API() gfxapi.API {
	return api{}
}
func (c *GlTextureStorage3DEXT) TypeID() atom.TypeID {
	return 36
}
func (c *GlTextureStorage3DEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenVertexArraysOES
////////////////////////////////////////////////////////////////////////////////
type GlGenVertexArraysOES struct {
	binary.Generate
	Count  int32
	Arrays VertexArrayIdArray
}

func (c *GlGenVertexArraysOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenVertexArraysOES(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Arrays),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenVertexArraysOES) API() gfxapi.API {
	return api{}
}
func (c *GlGenVertexArraysOES) TypeID() atom.TypeID {
	return 37
}
func (c *GlGenVertexArraysOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBindVertexArrayOES
////////////////////////////////////////////////////////////////////////////////
type GlBindVertexArrayOES struct {
	binary.Generate
	Array VertexArrayId
}

func (c *GlBindVertexArrayOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindVertexArrayOES(",
		fmt.Sprintf("array:%v", c.Array),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindVertexArrayOES) API() gfxapi.API {
	return api{}
}
func (c *GlBindVertexArrayOES) TypeID() atom.TypeID {
	return 38
}
func (c *GlBindVertexArrayOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteVertexArraysOES
////////////////////////////////////////////////////////////////////////////////
type GlDeleteVertexArraysOES struct {
	binary.Generate
	Count  int32
	Arrays VertexArrayIdArray
}

func (c *GlDeleteVertexArraysOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteVertexArraysOES(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Arrays),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteVertexArraysOES) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteVertexArraysOES) TypeID() atom.TypeID {
	return 39
}
func (c *GlDeleteVertexArraysOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsVertexArrayOES
////////////////////////////////////////////////////////////////////////////////
type GlIsVertexArrayOES struct {
	binary.Generate
	Array  VertexArrayId
	Result bool
}

func (c *GlIsVertexArrayOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsVertexArrayOES(",
		fmt.Sprintf("array:%v", c.Array),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsVertexArrayOES) API() gfxapi.API {
	return api{}
}
func (c *GlIsVertexArrayOES) TypeID() atom.TypeID {
	return 40
}
func (c *GlIsVertexArrayOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetTexture2DOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetTexture2DOES struct {
	binary.Generate
	Target ImageTargetTexture
	Image  ImageOES
}

func (c *GlEGLImageTargetTexture2DOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEGLImageTargetTexture2DOES(",
		c.Target.String(),
		", ",
		fmt.Sprintf("image:%v", c.Image),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEGLImageTargetTexture2DOES) API() gfxapi.API {
	return api{}
}
func (c *GlEGLImageTargetTexture2DOES) TypeID() atom.TypeID {
	return 41
}
func (c *GlEGLImageTargetTexture2DOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetRenderbufferStorageOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetRenderbufferStorageOES struct {
	binary.Generate
	Target ImageTargetRenderbufferStorage
	Image  TexturePointer
}

func (c *GlEGLImageTargetRenderbufferStorageOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEGLImageTargetRenderbufferStorageOES(",
		c.Target.String(),
		", ",
		fmt.Sprintf("image:%v", c.Image),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEGLImageTargetRenderbufferStorageOES) API() gfxapi.API {
	return api{}
}
func (c *GlEGLImageTargetRenderbufferStorageOES) TypeID() atom.TypeID {
	return 42
}
func (c *GlEGLImageTargetRenderbufferStorageOES) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetGraphicsResetStatusEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetGraphicsResetStatusEXT struct {
	binary.Generate
	Result ResetStatus
}

func (c *GlGetGraphicsResetStatusEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetGraphicsResetStatusEXT(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlGetGraphicsResetStatusEXT) API() gfxapi.API {
	return api{}
}
func (c *GlGetGraphicsResetStatusEXT) TypeID() atom.TypeID {
	return 43
}
func (c *GlGetGraphicsResetStatusEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBindAttribLocation
////////////////////////////////////////////////////////////////////////////////
type GlBindAttribLocation struct {
	binary.Generate
	Program  ProgramId
	Location AttributeLocation
	Name     string
}

func (c *GlBindAttribLocation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindAttribLocation(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("name:%v", c.Name),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindAttribLocation) API() gfxapi.API {
	return api{}
}
func (c *GlBindAttribLocation) TypeID() atom.TypeID {
	return 44
}
func (c *GlBindAttribLocation) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBlendFunc
////////////////////////////////////////////////////////////////////////////////
type GlBlendFunc struct {
	binary.Generate
	SrcFactor BlendFactor
	DstFactor BlendFactor
}

func (c *GlBlendFunc) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendFunc(",
		c.SrcFactor.String(),
		", ",
		c.DstFactor.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendFunc) API() gfxapi.API {
	return api{}
}
func (c *GlBlendFunc) TypeID() atom.TypeID {
	return 45
}
func (c *GlBlendFunc) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBlendFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendFuncSeparate struct {
	binary.Generate
	SrcFactorRgb   BlendFactor
	DstFactorRgb   BlendFactor
	SrcFactorAlpha BlendFactor
	DstFactorAlpha BlendFactor
}

func (c *GlBlendFuncSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendFuncSeparate(",
		c.SrcFactorRgb.String(),
		", ",
		c.DstFactorRgb.String(),
		", ",
		c.SrcFactorAlpha.String(),
		", ",
		c.DstFactorAlpha.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendFuncSeparate) API() gfxapi.API {
	return api{}
}
func (c *GlBlendFuncSeparate) TypeID() atom.TypeID {
	return 46
}
func (c *GlBlendFuncSeparate) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquation
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquation struct {
	binary.Generate
	Equation BlendEquation
}

func (c *GlBlendEquation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendEquation(",
		c.Equation.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendEquation) API() gfxapi.API {
	return api{}
}
func (c *GlBlendEquation) TypeID() atom.TypeID {
	return 47
}
func (c *GlBlendEquation) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquationSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquationSeparate struct {
	binary.Generate
	Rgb   BlendEquation
	Alpha BlendEquation
}

func (c *GlBlendEquationSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendEquationSeparate(",
		c.Rgb.String(),
		", ",
		c.Alpha.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendEquationSeparate) API() gfxapi.API {
	return api{}
}
func (c *GlBlendEquationSeparate) TypeID() atom.TypeID {
	return 48
}
func (c *GlBlendEquationSeparate) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBlendColor
////////////////////////////////////////////////////////////////////////////////
type GlBlendColor struct {
	binary.Generate
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}

func (c *GlBlendColor) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendColor(",
		fmt.Sprintf("red:%v", c.Red),
		", ",
		fmt.Sprintf("green:%v", c.Green),
		", ",
		fmt.Sprintf("blue:%v", c.Blue),
		", ",
		fmt.Sprintf("alpha:%v", c.Alpha),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendColor) API() gfxapi.API {
	return api{}
}
func (c *GlBlendColor) TypeID() atom.TypeID {
	return 49
}
func (c *GlBlendColor) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEnableVertexAttribArray
////////////////////////////////////////////////////////////////////////////////
type GlEnableVertexAttribArray struct {
	binary.Generate
	Location AttributeLocation
}

func (c *GlEnableVertexAttribArray) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEnableVertexAttribArray(",
		fmt.Sprintf("location:%v", c.Location),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEnableVertexAttribArray) API() gfxapi.API {
	return api{}
}
func (c *GlEnableVertexAttribArray) TypeID() atom.TypeID {
	return 50
}
func (c *GlEnableVertexAttribArray) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDisableVertexAttribArray
////////////////////////////////////////////////////////////////////////////////
type GlDisableVertexAttribArray struct {
	binary.Generate
	Location AttributeLocation
}

func (c *GlDisableVertexAttribArray) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDisableVertexAttribArray(",
		fmt.Sprintf("location:%v", c.Location),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDisableVertexAttribArray) API() gfxapi.API {
	return api{}
}
func (c *GlDisableVertexAttribArray) TypeID() atom.TypeID {
	return 51
}
func (c *GlDisableVertexAttribArray) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttribPointer
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttribPointer struct {
	binary.Generate
	Location   AttributeLocation
	Size       int32
	Type       VertexAttribType
	Normalized bool
	Stride     int32
	Data       VertexPointer
}

func (c *GlVertexAttribPointer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttribPointer(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("size:%v", c.Size),
		", ",
		c.Type.String(),
		", ",
		fmt.Sprintf("normalized:%v", c.Normalized),
		", ",
		fmt.Sprintf("stride:%v", c.Stride),
		", ",
		fmt.Sprintf("data:%v", c.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttribPointer) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttribPointer) TypeID() atom.TypeID {
	return 52
}
func (c *GlVertexAttribPointer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveAttrib
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveAttrib struct {
	binary.Generate
	Program            ProgramId
	Location           AttributeLocation
	BufferSize         int32
	BufferBytesWritten int32
	VectorCount        int32
	Type               ShaderAttribType
	Name               string
}

func (c *GlGetActiveAttrib) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetActiveAttrib(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("buffer_size:%v", c.BufferSize),
		", ",
		fmt.Sprintf("buffer_bytes_written:%v", c.BufferBytesWritten),
		", ",
		fmt.Sprintf("vector_count:%v", c.VectorCount),
		", ",
		c.Type.String(),
		", ",
		fmt.Sprintf("name:%v", c.Name),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetActiveAttrib) API() gfxapi.API {
	return api{}
}
func (c *GlGetActiveAttrib) TypeID() atom.TypeID {
	return 53
}
func (c *GlGetActiveAttrib) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveUniform
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveUniform struct {
	binary.Generate
	Program            ProgramId
	Location           int32
	BufferSize         int32
	BufferBytesWritten int32
	Size               int32
	Type               ShaderUniformType
	Name               string
}

func (c *GlGetActiveUniform) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetActiveUniform(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("buffer_size:%v", c.BufferSize),
		", ",
		fmt.Sprintf("buffer_bytes_written:%v", c.BufferBytesWritten),
		", ",
		fmt.Sprintf("size:%v", c.Size),
		", ",
		c.Type.String(),
		", ",
		fmt.Sprintf("name:%v", c.Name),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetActiveUniform) API() gfxapi.API {
	return api{}
}
func (c *GlGetActiveUniform) TypeID() atom.TypeID {
	return 54
}
func (c *GlGetActiveUniform) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetError
////////////////////////////////////////////////////////////////////////////////
type GlGetError struct {
	binary.Generate
	Result Error
}

func (c *GlGetError) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetError(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlGetError) API() gfxapi.API {
	return api{}
}
func (c *GlGetError) TypeID() atom.TypeID {
	return 55
}
func (c *GlGetError) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramiv
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramiv struct {
	binary.Generate
	Program   ProgramId
	Parameter ProgramParameter
	Value     S32Array
}

func (c *GlGetProgramiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetProgramiv(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetProgramiv) API() gfxapi.API {
	return api{}
}
func (c *GlGetProgramiv) TypeID() atom.TypeID {
	return 56
}
func (c *GlGetProgramiv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderiv
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderiv struct {
	binary.Generate
	Shader    ShaderId
	Parameter ShaderParameter
	Value     S32Array
}

func (c *GlGetShaderiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderiv(",
		fmt.Sprintf("shader:%v", c.Shader),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderiv) API() gfxapi.API {
	return api{}
}
func (c *GlGetShaderiv) TypeID() atom.TypeID {
	return 57
}
func (c *GlGetShaderiv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformLocation
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformLocation struct {
	binary.Generate
	Program ProgramId
	Name    string
	Result  UniformLocation
}

func (c *GlGetUniformLocation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetUniformLocation(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("name:%v", c.Name),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlGetUniformLocation) API() gfxapi.API {
	return api{}
}
func (c *GlGetUniformLocation) TypeID() atom.TypeID {
	return 58
}
func (c *GlGetUniformLocation) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetAttribLocation
////////////////////////////////////////////////////////////////////////////////
type GlGetAttribLocation struct {
	binary.Generate
	Program ProgramId
	Name    string
	Result  AttributeLocation
}

func (c *GlGetAttribLocation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetAttribLocation(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("name:%v", c.Name),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlGetAttribLocation) API() gfxapi.API {
	return api{}
}
func (c *GlGetAttribLocation) TypeID() atom.TypeID {
	return 59
}
func (c *GlGetAttribLocation) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlPixelStorei
////////////////////////////////////////////////////////////////////////////////
type GlPixelStorei struct {
	binary.Generate
	Parameter PixelStoreParameter
	Value     int32
}

func (c *GlPixelStorei) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPixelStorei(",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPixelStorei) API() gfxapi.API {
	return api{}
}
func (c *GlPixelStorei) TypeID() atom.TypeID {
	return 60
}
func (c *GlPixelStorei) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTexParameteri
////////////////////////////////////////////////////////////////////////////////
type GlTexParameteri struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Value     int32
}

func (c *GlTexParameteri) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexParameteri(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexParameteri) API() gfxapi.API {
	return api{}
}
func (c *GlTexParameteri) TypeID() atom.TypeID {
	return 61
}
func (c *GlTexParameteri) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTexParameterf
////////////////////////////////////////////////////////////////////////////////
type GlTexParameterf struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Value     float32
}

func (c *GlTexParameterf) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexParameterf(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexParameterf) API() gfxapi.API {
	return api{}
}
func (c *GlTexParameterf) TypeID() atom.TypeID {
	return 62
}
func (c *GlTexParameterf) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameteriv struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Values    S32Array
}

func (c *GlGetTexParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetTexParameteriv(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetTexParameteriv) API() gfxapi.API {
	return api{}
}
func (c *GlGetTexParameteriv) TypeID() atom.TypeID {
	return 63
}
func (c *GlGetTexParameteriv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameterfv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameterfv struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Values    F32Array
}

func (c *GlGetTexParameterfv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetTexParameterfv(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetTexParameterfv) API() gfxapi.API {
	return api{}
}
func (c *GlGetTexParameterfv) TypeID() atom.TypeID {
	return 64
}
func (c *GlGetTexParameterfv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform1i
////////////////////////////////////////////////////////////////////////////////
type GlUniform1i struct {
	binary.Generate
	Location UniformLocation
	Value    int32
}

func (c *GlUniform1i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1i(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1i) API() gfxapi.API {
	return api{}
}
func (c *GlUniform1i) TypeID() atom.TypeID {
	return 65
}
func (c *GlUniform1i) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform2i
////////////////////////////////////////////////////////////////////////////////
type GlUniform2i struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
}

func (c *GlUniform2i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2i(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2i) API() gfxapi.API {
	return api{}
}
func (c *GlUniform2i) TypeID() atom.TypeID {
	return 66
}
func (c *GlUniform2i) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform3i
////////////////////////////////////////////////////////////////////////////////
type GlUniform3i struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
	Value2   int32
}

func (c *GlUniform3i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3i(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.Value2),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3i) API() gfxapi.API {
	return api{}
}
func (c *GlUniform3i) TypeID() atom.TypeID {
	return 67
}
func (c *GlUniform3i) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform4i
////////////////////////////////////////////////////////////////////////////////
type GlUniform4i struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
	Value2   int32
	Value3   int32
}

func (c *GlUniform4i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4i(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.Value2),
		", ",
		fmt.Sprintf("value3:%v", c.Value3),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4i) API() gfxapi.API {
	return api{}
}
func (c *GlUniform4i) TypeID() atom.TypeID {
	return 68
}
func (c *GlUniform4i) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform1iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform1iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}

func (c *GlUniform1iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1iv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1iv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform1iv) TypeID() atom.TypeID {
	return 69
}
func (c *GlUniform1iv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform2iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform2iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}

func (c *GlUniform2iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2iv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2iv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform2iv) TypeID() atom.TypeID {
	return 70
}
func (c *GlUniform2iv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform3iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform3iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}

func (c *GlUniform3iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3iv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3iv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform3iv) TypeID() atom.TypeID {
	return 71
}
func (c *GlUniform3iv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform4iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform4iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}

func (c *GlUniform4iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4iv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4iv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform4iv) TypeID() atom.TypeID {
	return 72
}
func (c *GlUniform4iv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform1f
////////////////////////////////////////////////////////////////////////////////
type GlUniform1f struct {
	binary.Generate
	Location UniformLocation
	Value    float32
}

func (c *GlUniform1f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1f) API() gfxapi.API {
	return api{}
}
func (c *GlUniform1f) TypeID() atom.TypeID {
	return 73
}
func (c *GlUniform1f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform2f
////////////////////////////////////////////////////////////////////////////////
type GlUniform2f struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
}

func (c *GlUniform2f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2f) API() gfxapi.API {
	return api{}
}
func (c *GlUniform2f) TypeID() atom.TypeID {
	return 74
}
func (c *GlUniform2f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform3f
////////////////////////////////////////////////////////////////////////////////
type GlUniform3f struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
	Value2   float32
}

func (c *GlUniform3f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.Value2),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3f) API() gfxapi.API {
	return api{}
}
func (c *GlUniform3f) TypeID() atom.TypeID {
	return 75
}
func (c *GlUniform3f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform4f
////////////////////////////////////////////////////////////////////////////////
type GlUniform4f struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
	Value2   float32
	Value3   float32
}

func (c *GlUniform4f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.Value2),
		", ",
		fmt.Sprintf("value3:%v", c.Value3),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4f) API() gfxapi.API {
	return api{}
}
func (c *GlUniform4f) TypeID() atom.TypeID {
	return 76
}
func (c *GlUniform4f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform1fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform1fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}

func (c *GlUniform1fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1fv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform1fv) TypeID() atom.TypeID {
	return 77
}
func (c *GlUniform1fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform2fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform2fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}

func (c *GlUniform2fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2fv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform2fv) TypeID() atom.TypeID {
	return 78
}
func (c *GlUniform2fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform3fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform3fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}

func (c *GlUniform3fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3fv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform3fv) TypeID() atom.TypeID {
	return 79
}
func (c *GlUniform3fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniform4fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform4fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}

func (c *GlUniform4fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4fv) API() gfxapi.API {
	return api{}
}
func (c *GlUniform4fv) TypeID() atom.TypeID {
	return 80
}
func (c *GlUniform4fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix2fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix2fv struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}

func (c *GlUniformMatrix2fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniformMatrix2fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("transpose:%v", c.Transpose),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniformMatrix2fv) API() gfxapi.API {
	return api{}
}
func (c *GlUniformMatrix2fv) TypeID() atom.TypeID {
	return 81
}
func (c *GlUniformMatrix2fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix3fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix3fv struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}

func (c *GlUniformMatrix3fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniformMatrix3fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("transpose:%v", c.Transpose),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniformMatrix3fv) API() gfxapi.API {
	return api{}
}
func (c *GlUniformMatrix3fv) TypeID() atom.TypeID {
	return 82
}
func (c *GlUniformMatrix3fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix4fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix4fv struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}

func (c *GlUniformMatrix4fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniformMatrix4fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("transpose:%v", c.Transpose),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniformMatrix4fv) API() gfxapi.API {
	return api{}
}
func (c *GlUniformMatrix4fv) TypeID() atom.TypeID {
	return 83
}
func (c *GlUniformMatrix4fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformfv
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformfv struct {
	binary.Generate
	Program  ProgramId
	Location UniformLocation
	Values   F32Array
}

func (c *GlGetUniformfv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetUniformfv(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetUniformfv) API() gfxapi.API {
	return api{}
}
func (c *GlGetUniformfv) TypeID() atom.TypeID {
	return 84
}
func (c *GlGetUniformfv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformiv
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformiv struct {
	binary.Generate
	Program  ProgramId
	Location UniformLocation
	Values   S32Array
}

func (c *GlGetUniformiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetUniformiv(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetUniformiv) API() gfxapi.API {
	return api{}
}
func (c *GlGetUniformiv) TypeID() atom.TypeID {
	return 85
}
func (c *GlGetUniformiv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib1f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib1f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
}

func (c *GlVertexAttrib1f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib1f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib1f) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib1f) TypeID() atom.TypeID {
	return 86
}
func (c *GlVertexAttrib1f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib2f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib2f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
}

func (c *GlVertexAttrib2f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib2f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib2f) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib2f) TypeID() atom.TypeID {
	return 87
}
func (c *GlVertexAttrib2f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib3f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib3f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
	Value2   float32
}

func (c *GlVertexAttrib3f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib3f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.Value2),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib3f) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib3f) TypeID() atom.TypeID {
	return 88
}
func (c *GlVertexAttrib3f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib4f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib4f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
	Value2   float32
	Value3   float32
}

func (c *GlVertexAttrib4f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib4f(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("value0:%v", c.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.Value2),
		", ",
		fmt.Sprintf("value3:%v", c.Value3),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib4f) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib4f) TypeID() atom.TypeID {
	return 89
}
func (c *GlVertexAttrib4f) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib1fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib1fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}

func (c *GlVertexAttrib1fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib1fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib1fv) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib1fv) TypeID() atom.TypeID {
	return 90
}
func (c *GlVertexAttrib1fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib2fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib2fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}

func (c *GlVertexAttrib2fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib2fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib2fv) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib2fv) TypeID() atom.TypeID {
	return 91
}
func (c *GlVertexAttrib2fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib3fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib3fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}

func (c *GlVertexAttrib3fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib3fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib3fv) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib3fv) TypeID() atom.TypeID {
	return 92
}
func (c *GlVertexAttrib3fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib4fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib4fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}

func (c *GlVertexAttrib4fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib4fv(",
		fmt.Sprintf("location:%v", c.Location),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib4fv) API() gfxapi.API {
	return api{}
}
func (c *GlVertexAttrib4fv) TypeID() atom.TypeID {
	return 93
}
func (c *GlVertexAttrib4fv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderPrecisionFormat
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderPrecisionFormat struct {
	binary.Generate
	ShaderType    ShaderType
	PrecisionType PrecisionType
	Range         S32Array
	Precision     int32
}

func (c *GlGetShaderPrecisionFormat) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderPrecisionFormat(",
		c.ShaderType.String(),
		", ",
		c.PrecisionType.String(),
		", ",
		fmt.Sprintf("%v", c.Range),
		", ",
		fmt.Sprintf("precision:%v", c.Precision),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderPrecisionFormat) API() gfxapi.API {
	return api{}
}
func (c *GlGetShaderPrecisionFormat) TypeID() atom.TypeID {
	return 94
}
func (c *GlGetShaderPrecisionFormat) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDepthMask
////////////////////////////////////////////////////////////////////////////////
type GlDepthMask struct {
	binary.Generate
	Enabled bool
}

func (c *GlDepthMask) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDepthMask(",
		fmt.Sprintf("enabled:%v", c.Enabled),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDepthMask) API() gfxapi.API {
	return api{}
}
func (c *GlDepthMask) TypeID() atom.TypeID {
	return 95
}
func (c *GlDepthMask) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDepthFunc
////////////////////////////////////////////////////////////////////////////////
type GlDepthFunc struct {
	binary.Generate
	Function TestFunction
}

func (c *GlDepthFunc) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDepthFunc(",
		c.Function.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDepthFunc) API() gfxapi.API {
	return api{}
}
func (c *GlDepthFunc) TypeID() atom.TypeID {
	return 96
}
func (c *GlDepthFunc) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDepthRangef
////////////////////////////////////////////////////////////////////////////////
type GlDepthRangef struct {
	binary.Generate
	Near float32
	Far  float32
}

func (c *GlDepthRangef) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDepthRangef(",
		fmt.Sprintf("near:%v", c.Near),
		", ",
		fmt.Sprintf("far:%v", c.Far),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDepthRangef) API() gfxapi.API {
	return api{}
}
func (c *GlDepthRangef) TypeID() atom.TypeID {
	return 97
}
func (c *GlDepthRangef) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlColorMask
////////////////////////////////////////////////////////////////////////////////
type GlColorMask struct {
	binary.Generate
	Red   bool
	Green bool
	Blue  bool
	Alpha bool
}

func (c *GlColorMask) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glColorMask(",
		fmt.Sprintf("red:%v", c.Red),
		", ",
		fmt.Sprintf("green:%v", c.Green),
		", ",
		fmt.Sprintf("blue:%v", c.Blue),
		", ",
		fmt.Sprintf("alpha:%v", c.Alpha),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlColorMask) API() gfxapi.API {
	return api{}
}
func (c *GlColorMask) TypeID() atom.TypeID {
	return 98
}
func (c *GlColorMask) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlStencilMask
////////////////////////////////////////////////////////////////////////////////
type GlStencilMask struct {
	binary.Generate
	Mask uint32
}

func (c *GlStencilMask) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilMask(",
		fmt.Sprintf("mask:%v", c.Mask),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilMask) API() gfxapi.API {
	return api{}
}
func (c *GlStencilMask) TypeID() atom.TypeID {
	return 99
}
func (c *GlStencilMask) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlStencilMaskSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilMaskSeparate struct {
	binary.Generate
	Face FaceMode
	Mask uint32
}

func (c *GlStencilMaskSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilMaskSeparate(",
		c.Face.String(),
		", ",
		fmt.Sprintf("mask:%v", c.Mask),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilMaskSeparate) API() gfxapi.API {
	return api{}
}
func (c *GlStencilMaskSeparate) TypeID() atom.TypeID {
	return 100
}
func (c *GlStencilMaskSeparate) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlStencilFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilFuncSeparate struct {
	binary.Generate
	Face           FaceMode
	Function       TestFunction
	ReferenceValue int32
	Mask           int32
}

func (c *GlStencilFuncSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilFuncSeparate(",
		c.Face.String(),
		", ",
		c.Function.String(),
		", ",
		fmt.Sprintf("reference_value:%v", c.ReferenceValue),
		", ",
		fmt.Sprintf("mask:%v", c.Mask),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilFuncSeparate) API() gfxapi.API {
	return api{}
}
func (c *GlStencilFuncSeparate) TypeID() atom.TypeID {
	return 101
}
func (c *GlStencilFuncSeparate) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlStencilOpSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilOpSeparate struct {
	binary.Generate
	Face                 FaceMode
	StencilFail          StencilAction
	StencilPassDepthFail StencilAction
	StencilPassDepthPass StencilAction
}

func (c *GlStencilOpSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilOpSeparate(",
		c.Face.String(),
		", ",
		c.StencilFail.String(),
		", ",
		c.StencilPassDepthFail.String(),
		", ",
		c.StencilPassDepthPass.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilOpSeparate) API() gfxapi.API {
	return api{}
}
func (c *GlStencilOpSeparate) TypeID() atom.TypeID {
	return 102
}
func (c *GlStencilOpSeparate) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlFrontFace
////////////////////////////////////////////////////////////////////////////////
type GlFrontFace struct {
	binary.Generate
	Orientation FaceOrientation
}

func (c *GlFrontFace) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFrontFace(",
		c.Orientation.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFrontFace) API() gfxapi.API {
	return api{}
}
func (c *GlFrontFace) TypeID() atom.TypeID {
	return 103
}
func (c *GlFrontFace) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlViewport
////////////////////////////////////////////////////////////////////////////////
type GlViewport struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func (c *GlViewport) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glViewport(",
		fmt.Sprintf("x:%v", c.X),
		", ",
		fmt.Sprintf("y:%v", c.Y),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlViewport) API() gfxapi.API {
	return api{}
}
func (c *GlViewport) TypeID() atom.TypeID {
	return 104
}
func (c *GlViewport) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlScissor
////////////////////////////////////////////////////////////////////////////////
type GlScissor struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
}

func (c *GlScissor) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glScissor(",
		fmt.Sprintf("x:%v", c.X),
		", ",
		fmt.Sprintf("y:%v", c.Y),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlScissor) API() gfxapi.API {
	return api{}
}
func (c *GlScissor) TypeID() atom.TypeID {
	return 105
}
func (c *GlScissor) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlActiveTexture
////////////////////////////////////////////////////////////////////////////////
type GlActiveTexture struct {
	binary.Generate
	Unit TextureUnit
}

func (c *GlActiveTexture) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glActiveTexture(",
		c.Unit.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlActiveTexture) API() gfxapi.API {
	return api{}
}
func (c *GlActiveTexture) TypeID() atom.TypeID {
	return 106
}
func (c *GlActiveTexture) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenTextures
////////////////////////////////////////////////////////////////////////////////
type GlGenTextures struct {
	binary.Generate
	Count    int32
	Textures TextureIdArray
}

func (c *GlGenTextures) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenTextures(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Textures),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenTextures) API() gfxapi.API {
	return api{}
}
func (c *GlGenTextures) TypeID() atom.TypeID {
	return 107
}
func (c *GlGenTextures) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteTextures
////////////////////////////////////////////////////////////////////////////////
type GlDeleteTextures struct {
	binary.Generate
	Count    int32
	Textures TextureIdArray
}

func (c *GlDeleteTextures) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteTextures(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Textures),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteTextures) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteTextures) TypeID() atom.TypeID {
	return 108
}
func (c *GlDeleteTextures) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsTexture
////////////////////////////////////////////////////////////////////////////////
type GlIsTexture struct {
	binary.Generate
	Texture TextureId
	Result  bool
}

func (c *GlIsTexture) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsTexture(",
		fmt.Sprintf("texture:%v", c.Texture),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsTexture) API() gfxapi.API {
	return api{}
}
func (c *GlIsTexture) TypeID() atom.TypeID {
	return 109
}
func (c *GlIsTexture) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBindTexture
////////////////////////////////////////////////////////////////////////////////
type GlBindTexture struct {
	binary.Generate
	Target  TextureTarget
	Texture TextureId
}

func (c *GlBindTexture) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindTexture(",
		c.Target.String(),
		", ",
		fmt.Sprintf("texture:%v", c.Texture),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindTexture) API() gfxapi.API {
	return api{}
}
func (c *GlBindTexture) TypeID() atom.TypeID {
	return 110
}
func (c *GlBindTexture) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexImage2D struct {
	binary.Generate
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

func (c *GlTexImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexImage2D(",
		c.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.Level),
		", ",
		c.InternalFormat.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		fmt.Sprintf("border:%v", c.Border),
		", ",
		c.Format.String(),
		", ",
		c.Type.String(),
		", ",
		fmt.Sprintf("data:%v", c.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexImage2D) API() gfxapi.API {
	return api{}
}
func (c *GlTexImage2D) TypeID() atom.TypeID {
	return 111
}
func (c *GlTexImage2D) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexSubImage2D struct {
	binary.Generate
	Target  TextureImageTarget
	Level   int32
	Xoffset int32
	Yoffset int32
	Width   int32
	Height  int32
	Format  TexelFormat
	Type    TexelType
	Data    TexturePointer
}

func (c *GlTexSubImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexSubImage2D(",
		c.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.Level),
		", ",
		fmt.Sprintf("xoffset:%v", c.Xoffset),
		", ",
		fmt.Sprintf("yoffset:%v", c.Yoffset),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		c.Format.String(),
		", ",
		c.Type.String(),
		", ",
		fmt.Sprintf("data:%v", c.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexSubImage2D) API() gfxapi.API {
	return api{}
}
func (c *GlTexSubImage2D) TypeID() atom.TypeID {
	return 112
}
func (c *GlTexSubImage2D) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexImage2D struct {
	binary.Generate
	Target TextureImageTarget
	Level  int32
	Format TexelFormat
	X      int32
	Y      int32
	Width  int32
	Height int32
	Border int32
}

func (c *GlCopyTexImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCopyTexImage2D(",
		c.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.Level),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("x:%v", c.X),
		", ",
		fmt.Sprintf("y:%v", c.Y),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		fmt.Sprintf("border:%v", c.Border),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCopyTexImage2D) API() gfxapi.API {
	return api{}
}
func (c *GlCopyTexImage2D) TypeID() atom.TypeID {
	return 113
}
func (c *GlCopyTexImage2D) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexSubImage2D struct {
	binary.Generate
	Target  TextureImageTarget
	Level   int32
	Xoffset int32
	Yoffset int32
	X       int32
	Y       int32
	Width   int32
	Height  int32
}

func (c *GlCopyTexSubImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCopyTexSubImage2D(",
		c.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.Level),
		", ",
		fmt.Sprintf("xoffset:%v", c.Xoffset),
		", ",
		fmt.Sprintf("yoffset:%v", c.Yoffset),
		", ",
		fmt.Sprintf("x:%v", c.X),
		", ",
		fmt.Sprintf("y:%v", c.Y),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCopyTexSubImage2D) API() gfxapi.API {
	return api{}
}
func (c *GlCopyTexSubImage2D) TypeID() atom.TypeID {
	return 114
}
func (c *GlCopyTexSubImage2D) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexImage2D struct {
	binary.Generate
	Target    TextureImageTarget
	Level     int32
	Format    CompressedTexelFormat
	Width     int32
	Height    int32
	Border    int32
	ImageSize int32
	Data      TexturePointer
}

func (c *GlCompressedTexImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCompressedTexImage2D(",
		c.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.Level),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		fmt.Sprintf("border:%v", c.Border),
		", ",
		fmt.Sprintf("image_size:%v", c.ImageSize),
		", ",
		fmt.Sprintf("data:%v", c.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCompressedTexImage2D) API() gfxapi.API {
	return api{}
}
func (c *GlCompressedTexImage2D) TypeID() atom.TypeID {
	return 115
}
func (c *GlCompressedTexImage2D) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexSubImage2D struct {
	binary.Generate
	Target    TextureImageTarget
	Level     int32
	Xoffset   int32
	Yoffset   int32
	Width     int32
	Height    int32
	Format    CompressedTexelFormat
	ImageSize int32
	Data      TexturePointer
}

func (c *GlCompressedTexSubImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCompressedTexSubImage2D(",
		c.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.Level),
		", ",
		fmt.Sprintf("xoffset:%v", c.Xoffset),
		", ",
		fmt.Sprintf("yoffset:%v", c.Yoffset),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("image_size:%v", c.ImageSize),
		", ",
		fmt.Sprintf("data:%v", c.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCompressedTexSubImage2D) API() gfxapi.API {
	return api{}
}
func (c *GlCompressedTexSubImage2D) TypeID() atom.TypeID {
	return 116
}
func (c *GlCompressedTexSubImage2D) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenerateMipmap
////////////////////////////////////////////////////////////////////////////////
type GlGenerateMipmap struct {
	binary.Generate
	Target TextureImageTarget
}

func (c *GlGenerateMipmap) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenerateMipmap(",
		c.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenerateMipmap) API() gfxapi.API {
	return api{}
}
func (c *GlGenerateMipmap) TypeID() atom.TypeID {
	return 117
}
func (c *GlGenerateMipmap) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlReadPixels
////////////////////////////////////////////////////////////////////////////////
type GlReadPixels struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
	Format BaseTexelFormat
	Type   TexelType
	Data   memory.Pointer
}

func (c *GlReadPixels) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glReadPixels(",
		fmt.Sprintf("x:%v", c.X),
		", ",
		fmt.Sprintf("y:%v", c.Y),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
		", ",
		c.Format.String(),
		", ",
		c.Type.String(),
		", ",
		fmt.Sprintf("0x%x", c.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlReadPixels) API() gfxapi.API {
	return api{}
}
func (c *GlReadPixels) TypeID() atom.TypeID {
	return 118
}
func (c *GlReadPixels) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenFramebuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenFramebuffers struct {
	binary.Generate
	Count        int32
	Framebuffers FramebufferIdArray
}

func (c *GlGenFramebuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenFramebuffers(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Framebuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenFramebuffers) API() gfxapi.API {
	return api{}
}
func (c *GlGenFramebuffers) TypeID() atom.TypeID {
	return 119
}
func (c *GlGenFramebuffers) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBindFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindFramebuffer struct {
	binary.Generate
	Target      FramebufferTarget
	Framebuffer FramebufferId
}

func (c *GlBindFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindFramebuffer(",
		c.Target.String(),
		", ",
		fmt.Sprintf("framebuffer:%v", c.Framebuffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindFramebuffer) API() gfxapi.API {
	return api{}
}
func (c *GlBindFramebuffer) TypeID() atom.TypeID {
	return 120
}
func (c *GlBindFramebuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCheckFramebufferStatus
////////////////////////////////////////////////////////////////////////////////
type GlCheckFramebufferStatus struct {
	binary.Generate
	Target FramebufferTarget
	Result FramebufferStatus
}

func (c *GlCheckFramebufferStatus) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCheckFramebufferStatus(",
		c.Target.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlCheckFramebufferStatus) API() gfxapi.API {
	return api{}
}
func (c *GlCheckFramebufferStatus) TypeID() atom.TypeID {
	return 121
}
func (c *GlCheckFramebufferStatus) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteFramebuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteFramebuffers struct {
	binary.Generate
	Count        int32
	Framebuffers FramebufferIdArray
}

func (c *GlDeleteFramebuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteFramebuffers(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Framebuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteFramebuffers) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteFramebuffers) TypeID() atom.TypeID {
	return 122
}
func (c *GlDeleteFramebuffers) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsFramebuffer struct {
	binary.Generate
	Framebuffer FramebufferId
	Result      bool
}

func (c *GlIsFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsFramebuffer(",
		fmt.Sprintf("framebuffer:%v", c.Framebuffer),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsFramebuffer) API() gfxapi.API {
	return api{}
}
func (c *GlIsFramebuffer) TypeID() atom.TypeID {
	return 123
}
func (c *GlIsFramebuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenRenderbuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenRenderbuffers struct {
	binary.Generate
	Count         int32
	Renderbuffers RenderbufferIdArray
}

func (c *GlGenRenderbuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenRenderbuffers(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Renderbuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenRenderbuffers) API() gfxapi.API {
	return api{}
}
func (c *GlGenRenderbuffers) TypeID() atom.TypeID {
	return 124
}
func (c *GlGenRenderbuffers) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBindRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindRenderbuffer struct {
	binary.Generate
	Target       RenderbufferTarget
	Renderbuffer RenderbufferId
}

func (c *GlBindRenderbuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindRenderbuffer(",
		c.Target.String(),
		", ",
		fmt.Sprintf("renderbuffer:%v", c.Renderbuffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindRenderbuffer) API() gfxapi.API {
	return api{}
}
func (c *GlBindRenderbuffer) TypeID() atom.TypeID {
	return 125
}
func (c *GlBindRenderbuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorage
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorage struct {
	binary.Generate
	Target RenderbufferTarget
	Format RenderbufferFormat
	Width  int32
	Height int32
}

func (c *GlRenderbufferStorage) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glRenderbufferStorage(",
		c.Target.String(),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlRenderbufferStorage) API() gfxapi.API {
	return api{}
}
func (c *GlRenderbufferStorage) TypeID() atom.TypeID {
	return 126
}
func (c *GlRenderbufferStorage) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteRenderbuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteRenderbuffers struct {
	binary.Generate
	Count         int32
	Renderbuffers RenderbufferIdArray
}

func (c *GlDeleteRenderbuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteRenderbuffers(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Renderbuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteRenderbuffers) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteRenderbuffers) TypeID() atom.TypeID {
	return 127
}
func (c *GlDeleteRenderbuffers) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsRenderbuffer struct {
	binary.Generate
	Renderbuffer RenderbufferId
	Result       bool
}

func (c *GlIsRenderbuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsRenderbuffer(",
		fmt.Sprintf("renderbuffer:%v", c.Renderbuffer),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsRenderbuffer) API() gfxapi.API {
	return api{}
}
func (c *GlIsRenderbuffer) TypeID() atom.TypeID {
	return 128
}
func (c *GlIsRenderbuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetRenderbufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetRenderbufferParameteriv struct {
	binary.Generate
	Target    RenderbufferTarget
	Parameter RenderbufferParameter
	Values    S32Array
}

func (c *GlGetRenderbufferParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetRenderbufferParameteriv(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetRenderbufferParameteriv) API() gfxapi.API {
	return api{}
}
func (c *GlGetRenderbufferParameteriv) TypeID() atom.TypeID {
	return 129
}
func (c *GlGetRenderbufferParameteriv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenBuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenBuffers struct {
	binary.Generate
	Count   int32
	Buffers BufferIdArray
}

func (c *GlGenBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenBuffers(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Buffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenBuffers) API() gfxapi.API {
	return api{}
}
func (c *GlGenBuffers) TypeID() atom.TypeID {
	return 130
}
func (c *GlGenBuffers) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBindBuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindBuffer struct {
	binary.Generate
	Target BufferTarget
	Buffer BufferId
}

func (c *GlBindBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindBuffer(",
		c.Target.String(),
		", ",
		fmt.Sprintf("buffer:%v", c.Buffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindBuffer) API() gfxapi.API {
	return api{}
}
func (c *GlBindBuffer) TypeID() atom.TypeID {
	return 131
}
func (c *GlBindBuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBufferData
////////////////////////////////////////////////////////////////////////////////
type GlBufferData struct {
	binary.Generate
	Target BufferTarget
	Size   int32
	Data   BufferDataPointer
	Usage  BufferUsage
}

func (c *GlBufferData) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBufferData(",
		c.Target.String(),
		", ",
		fmt.Sprintf("size:%v", c.Size),
		", ",
		fmt.Sprintf("data:%v", c.Data),
		", ",
		c.Usage.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBufferData) API() gfxapi.API {
	return api{}
}
func (c *GlBufferData) TypeID() atom.TypeID {
	return 132
}
func (c *GlBufferData) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBufferSubData
////////////////////////////////////////////////////////////////////////////////
type GlBufferSubData struct {
	binary.Generate
	Target BufferTarget
	Offset int32
	Size   int32
	Data   memory.Pointer
}

func (c *GlBufferSubData) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBufferSubData(",
		c.Target.String(),
		", ",
		fmt.Sprintf("offset:%v", c.Offset),
		", ",
		fmt.Sprintf("size:%v", c.Size),
		", ",
		fmt.Sprintf("0x%x", c.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBufferSubData) API() gfxapi.API {
	return api{}
}
func (c *GlBufferSubData) TypeID() atom.TypeID {
	return 133
}
func (c *GlBufferSubData) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteBuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteBuffers struct {
	binary.Generate
	Count   int32
	Buffers BufferIdArray
}

func (c *GlDeleteBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteBuffers(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Buffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteBuffers) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteBuffers) TypeID() atom.TypeID {
	return 134
}
func (c *GlDeleteBuffers) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsBuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsBuffer struct {
	binary.Generate
	Buffer BufferId
	Result bool
}

func (c *GlIsBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsBuffer(",
		fmt.Sprintf("buffer:%v", c.Buffer),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsBuffer) API() gfxapi.API {
	return api{}
}
func (c *GlIsBuffer) TypeID() atom.TypeID {
	return 135
}
func (c *GlIsBuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetBufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetBufferParameteriv struct {
	binary.Generate
	Target    BufferTarget
	Parameter BufferParameter
	Value     int32
}

func (c *GlGetBufferParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetBufferParameteriv(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetBufferParameteriv) API() gfxapi.API {
	return api{}
}
func (c *GlGetBufferParameteriv) TypeID() atom.TypeID {
	return 136
}
func (c *GlGetBufferParameteriv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCreateShader
////////////////////////////////////////////////////////////////////////////////
type GlCreateShader struct {
	binary.Generate
	Type   ShaderType
	Result ShaderId
}

func (c *GlCreateShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCreateShader(",
		c.Type.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlCreateShader) API() gfxapi.API {
	return api{}
}
func (c *GlCreateShader) TypeID() atom.TypeID {
	return 137
}
func (c *GlCreateShader) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteShader
////////////////////////////////////////////////////////////////////////////////
type GlDeleteShader struct {
	binary.Generate
	Shader ShaderId
}

func (c *GlDeleteShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteShader(",
		fmt.Sprintf("shader:%v", c.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteShader) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteShader) TypeID() atom.TypeID {
	return 138
}
func (c *GlDeleteShader) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlShaderSource
////////////////////////////////////////////////////////////////////////////////
type GlShaderSource struct {
	binary.Generate
	Shader ShaderId
	Count  int32
	Source StringArray
	Length S32Array
}

func (c *GlShaderSource) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glShaderSource(",
		fmt.Sprintf("shader:%v", c.Shader),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Source),
		", ",
		fmt.Sprintf("%v", c.Length),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlShaderSource) API() gfxapi.API {
	return api{}
}
func (c *GlShaderSource) TypeID() atom.TypeID {
	return 139
}
func (c *GlShaderSource) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlShaderBinary
////////////////////////////////////////////////////////////////////////////////
type GlShaderBinary struct {
	binary.Generate
	Count        int32
	Shaders      ShaderIdArray
	BinaryFormat uint32
	Binary       memory.Pointer
	BinarySize   int32
}

func (c *GlShaderBinary) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glShaderBinary(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Shaders),
		", ",
		fmt.Sprintf("binary_format:%v", c.BinaryFormat),
		", ",
		fmt.Sprintf("0x%x", c.Binary),
		", ",
		fmt.Sprintf("binary_size:%v", c.BinarySize),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlShaderBinary) API() gfxapi.API {
	return api{}
}
func (c *GlShaderBinary) TypeID() atom.TypeID {
	return 140
}
func (c *GlShaderBinary) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderInfoLog
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderInfoLog struct {
	binary.Generate
	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten int32
	Info                string
}

func (c *GlGetShaderInfoLog) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderInfoLog(",
		fmt.Sprintf("shader:%v", c.Shader),
		", ",
		fmt.Sprintf("buffer_length:%v", c.BufferLength),
		", ",
		fmt.Sprintf("string_length_written:%v", c.StringLengthWritten),
		", ",
		fmt.Sprintf("info:%v", c.Info),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderInfoLog) API() gfxapi.API {
	return api{}
}
func (c *GlGetShaderInfoLog) TypeID() atom.TypeID {
	return 141
}
func (c *GlGetShaderInfoLog) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderSource
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderSource struct {
	binary.Generate
	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten int32
	Source              string
}

func (c *GlGetShaderSource) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderSource(",
		fmt.Sprintf("shader:%v", c.Shader),
		", ",
		fmt.Sprintf("buffer_length:%v", c.BufferLength),
		", ",
		fmt.Sprintf("string_length_written:%v", c.StringLengthWritten),
		", ",
		fmt.Sprintf("source:%v", c.Source),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderSource) API() gfxapi.API {
	return api{}
}
func (c *GlGetShaderSource) TypeID() atom.TypeID {
	return 142
}
func (c *GlGetShaderSource) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlReleaseShaderCompiler
////////////////////////////////////////////////////////////////////////////////
type GlReleaseShaderCompiler struct {
	binary.Generate
}

func (c *GlReleaseShaderCompiler) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glReleaseShaderCompiler(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlReleaseShaderCompiler) API() gfxapi.API {
	return api{}
}
func (c *GlReleaseShaderCompiler) TypeID() atom.TypeID {
	return 143
}
func (c *GlReleaseShaderCompiler) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCompileShader
////////////////////////////////////////////////////////////////////////////////
type GlCompileShader struct {
	binary.Generate
	Shader ShaderId
}

func (c *GlCompileShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCompileShader(",
		fmt.Sprintf("shader:%v", c.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCompileShader) API() gfxapi.API {
	return api{}
}
func (c *GlCompileShader) TypeID() atom.TypeID {
	return 144
}
func (c *GlCompileShader) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsShader
////////////////////////////////////////////////////////////////////////////////
type GlIsShader struct {
	binary.Generate
	Shader ShaderId
	Result bool
}

func (c *GlIsShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsShader(",
		fmt.Sprintf("shader:%v", c.Shader),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsShader) API() gfxapi.API {
	return api{}
}
func (c *GlIsShader) TypeID() atom.TypeID {
	return 145
}
func (c *GlIsShader) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCreateProgram
////////////////////////////////////////////////////////////////////////////////
type GlCreateProgram struct {
	binary.Generate
	Result ProgramId
}

func (c *GlCreateProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCreateProgram(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlCreateProgram) API() gfxapi.API {
	return api{}
}
func (c *GlCreateProgram) TypeID() atom.TypeID {
	return 146
}
func (c *GlCreateProgram) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteProgram
////////////////////////////////////////////////////////////////////////////////
type GlDeleteProgram struct {
	binary.Generate
	Program ProgramId
}

func (c *GlDeleteProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteProgram(",
		fmt.Sprintf("program:%v", c.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteProgram) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteProgram) TypeID() atom.TypeID {
	return 147
}
func (c *GlDeleteProgram) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlAttachShader
////////////////////////////////////////////////////////////////////////////////
type GlAttachShader struct {
	binary.Generate
	Program ProgramId
	Shader  ShaderId
}

func (c *GlAttachShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glAttachShader(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("shader:%v", c.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlAttachShader) API() gfxapi.API {
	return api{}
}
func (c *GlAttachShader) TypeID() atom.TypeID {
	return 148
}
func (c *GlAttachShader) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDetachShader
////////////////////////////////////////////////////////////////////////////////
type GlDetachShader struct {
	binary.Generate
	Program ProgramId
	Shader  ShaderId
}

func (c *GlDetachShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDetachShader(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("shader:%v", c.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDetachShader) API() gfxapi.API {
	return api{}
}
func (c *GlDetachShader) TypeID() atom.TypeID {
	return 149
}
func (c *GlDetachShader) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetAttachedShaders
////////////////////////////////////////////////////////////////////////////////
type GlGetAttachedShaders struct {
	binary.Generate
	Program              ProgramId
	BufferLength         int32
	ShadersLengthWritten int32
	Shaders              ShaderIdArray
}

func (c *GlGetAttachedShaders) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetAttachedShaders(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("buffer_length:%v", c.BufferLength),
		", ",
		fmt.Sprintf("shaders_length_written:%v", c.ShadersLengthWritten),
		", ",
		fmt.Sprintf("%v", c.Shaders),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetAttachedShaders) API() gfxapi.API {
	return api{}
}
func (c *GlGetAttachedShaders) TypeID() atom.TypeID {
	return 150
}
func (c *GlGetAttachedShaders) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlLinkProgram
////////////////////////////////////////////////////////////////////////////////
type GlLinkProgram struct {
	binary.Generate
	Program ProgramId
}

func (c *GlLinkProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glLinkProgram(",
		fmt.Sprintf("program:%v", c.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlLinkProgram) API() gfxapi.API {
	return api{}
}
func (c *GlLinkProgram) TypeID() atom.TypeID {
	return 151
}
func (c *GlLinkProgram) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramInfoLog
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramInfoLog struct {
	binary.Generate
	Program             ProgramId
	BufferLength        int32
	StringLengthWritten int32
	Info                string
}

func (c *GlGetProgramInfoLog) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetProgramInfoLog(",
		fmt.Sprintf("program:%v", c.Program),
		", ",
		fmt.Sprintf("buffer_length:%v", c.BufferLength),
		", ",
		fmt.Sprintf("string_length_written:%v", c.StringLengthWritten),
		", ",
		fmt.Sprintf("info:%v", c.Info),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetProgramInfoLog) API() gfxapi.API {
	return api{}
}
func (c *GlGetProgramInfoLog) TypeID() atom.TypeID {
	return 152
}
func (c *GlGetProgramInfoLog) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUseProgram
////////////////////////////////////////////////////////////////////////////////
type GlUseProgram struct {
	binary.Generate
	Program ProgramId
}

func (c *GlUseProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUseProgram(",
		fmt.Sprintf("program:%v", c.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUseProgram) API() gfxapi.API {
	return api{}
}
func (c *GlUseProgram) TypeID() atom.TypeID {
	return 153
}
func (c *GlUseProgram) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsProgram
////////////////////////////////////////////////////////////////////////////////
type GlIsProgram struct {
	binary.Generate
	Program ProgramId
	Result  bool
}

func (c *GlIsProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsProgram(",
		fmt.Sprintf("program:%v", c.Program),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsProgram) API() gfxapi.API {
	return api{}
}
func (c *GlIsProgram) TypeID() atom.TypeID {
	return 154
}
func (c *GlIsProgram) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlValidateProgram
////////////////////////////////////////////////////////////////////////////////
type GlValidateProgram struct {
	binary.Generate
	Program ProgramId
}

func (c *GlValidateProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glValidateProgram(",
		fmt.Sprintf("program:%v", c.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlValidateProgram) API() gfxapi.API {
	return api{}
}
func (c *GlValidateProgram) TypeID() atom.TypeID {
	return 155
}
func (c *GlValidateProgram) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlClearColor
////////////////////////////////////////////////////////////////////////////////
type GlClearColor struct {
	binary.Generate
	R float32
	G float32
	B float32
	A float32
}

func (c *GlClearColor) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClearColor(",
		fmt.Sprintf("r:%v", c.R),
		", ",
		fmt.Sprintf("g:%v", c.G),
		", ",
		fmt.Sprintf("b:%v", c.B),
		", ",
		fmt.Sprintf("a:%v", c.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClearColor) API() gfxapi.API {
	return api{}
}
func (c *GlClearColor) TypeID() atom.TypeID {
	return 156
}
func (c *GlClearColor) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlClearDepthf
////////////////////////////////////////////////////////////////////////////////
type GlClearDepthf struct {
	binary.Generate
	Depth float32
}

func (c *GlClearDepthf) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClearDepthf(",
		fmt.Sprintf("depth:%v", c.Depth),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClearDepthf) API() gfxapi.API {
	return api{}
}
func (c *GlClearDepthf) TypeID() atom.TypeID {
	return 157
}
func (c *GlClearDepthf) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlClearStencil
////////////////////////////////////////////////////////////////////////////////
type GlClearStencil struct {
	binary.Generate
	Stencil int32
}

func (c *GlClearStencil) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClearStencil(",
		fmt.Sprintf("stencil:%v", c.Stencil),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClearStencil) API() gfxapi.API {
	return api{}
}
func (c *GlClearStencil) TypeID() atom.TypeID {
	return 158
}
func (c *GlClearStencil) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlClear
////////////////////////////////////////////////////////////////////////////////
type GlClear struct {
	binary.Generate
	Mask ClearMask
}

func (c *GlClear) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClear(",
		c.Mask.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClear) API() gfxapi.API {
	return api{}
}
func (c *GlClear) TypeID() atom.TypeID {
	return 159
}
func (c *GlClear) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlCullFace
////////////////////////////////////////////////////////////////////////////////
type GlCullFace struct {
	binary.Generate
	Mode FaceMode
}

func (c *GlCullFace) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCullFace(",
		c.Mode.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCullFace) API() gfxapi.API {
	return api{}
}
func (c *GlCullFace) TypeID() atom.TypeID {
	return 160
}
func (c *GlCullFace) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlPolygonOffset
////////////////////////////////////////////////////////////////////////////////
type GlPolygonOffset struct {
	binary.Generate
	ScaleFactor float32
	Units       float32
}

func (c *GlPolygonOffset) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPolygonOffset(",
		fmt.Sprintf("scale_factor:%v", c.ScaleFactor),
		", ",
		fmt.Sprintf("units:%v", c.Units),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPolygonOffset) API() gfxapi.API {
	return api{}
}
func (c *GlPolygonOffset) TypeID() atom.TypeID {
	return 161
}
func (c *GlPolygonOffset) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlLineWidth
////////////////////////////////////////////////////////////////////////////////
type GlLineWidth struct {
	binary.Generate
	Width float32
}

func (c *GlLineWidth) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glLineWidth(",
		fmt.Sprintf("width:%v", c.Width),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlLineWidth) API() gfxapi.API {
	return api{}
}
func (c *GlLineWidth) TypeID() atom.TypeID {
	return 162
}
func (c *GlLineWidth) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlSampleCoverage
////////////////////////////////////////////////////////////////////////////////
type GlSampleCoverage struct {
	binary.Generate
	Value  float32
	Invert bool
}

func (c *GlSampleCoverage) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glSampleCoverage(",
		fmt.Sprintf("value:%v", c.Value),
		", ",
		fmt.Sprintf("invert:%v", c.Invert),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlSampleCoverage) API() gfxapi.API {
	return api{}
}
func (c *GlSampleCoverage) TypeID() atom.TypeID {
	return 163
}
func (c *GlSampleCoverage) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlHint
////////////////////////////////////////////////////////////////////////////////
type GlHint struct {
	binary.Generate
	Target HintTarget
	Mode   HintMode
}

func (c *GlHint) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glHint(",
		c.Target.String(),
		", ",
		c.Mode.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlHint) API() gfxapi.API {
	return api{}
}
func (c *GlHint) TypeID() atom.TypeID {
	return 164
}
func (c *GlHint) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferRenderbuffer struct {
	binary.Generate
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	RenderbufferTarget    RenderbufferTarget
	Renderbuffer          RenderbufferId
}

func (c *GlFramebufferRenderbuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFramebufferRenderbuffer(",
		c.FramebufferTarget.String(),
		", ",
		c.FramebufferAttachment.String(),
		", ",
		c.RenderbufferTarget.String(),
		", ",
		fmt.Sprintf("renderbuffer:%v", c.Renderbuffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFramebufferRenderbuffer) API() gfxapi.API {
	return api{}
}
func (c *GlFramebufferRenderbuffer) TypeID() atom.TypeID {
	return 165
}
func (c *GlFramebufferRenderbuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferTexture2D
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferTexture2D struct {
	binary.Generate
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	TextureTarget         TextureImageTarget
	Texture               TextureId
	Level                 int32
}

func (c *GlFramebufferTexture2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFramebufferTexture2D(",
		c.FramebufferTarget.String(),
		", ",
		c.FramebufferAttachment.String(),
		", ",
		c.TextureTarget.String(),
		", ",
		fmt.Sprintf("texture:%v", c.Texture),
		", ",
		fmt.Sprintf("level:%v", c.Level),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFramebufferTexture2D) API() gfxapi.API {
	return api{}
}
func (c *GlFramebufferTexture2D) TypeID() atom.TypeID {
	return 166
}
func (c *GlFramebufferTexture2D) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetFramebufferAttachmentParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetFramebufferAttachmentParameteriv struct {
	binary.Generate
	FramebufferTarget FramebufferTarget
	Attachment        FramebufferAttachment
	Parameter         FramebufferAttachmentParameter
	Value             S32Array
}

func (c *GlGetFramebufferAttachmentParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetFramebufferAttachmentParameteriv(",
		c.FramebufferTarget.String(),
		", ",
		c.Attachment.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetFramebufferAttachmentParameteriv) API() gfxapi.API {
	return api{}
}
func (c *GlGetFramebufferAttachmentParameteriv) TypeID() atom.TypeID {
	return 167
}
func (c *GlGetFramebufferAttachmentParameteriv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDrawElements
////////////////////////////////////////////////////////////////////////////////
type GlDrawElements struct {
	binary.Generate
	DrawMode     DrawMode
	ElementCount int32
	IndicesType  IndicesType
	Indices      IndicesPointer
}

func (c *GlDrawElements) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDrawElements(",
		c.DrawMode.String(),
		", ",
		fmt.Sprintf("element_count:%v", c.ElementCount),
		", ",
		c.IndicesType.String(),
		", ",
		fmt.Sprintf("indices:%v", c.Indices),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDrawElements) API() gfxapi.API {
	return api{}
}
func (c *GlDrawElements) TypeID() atom.TypeID {
	return 168
}
func (c *GlDrawElements) Flags() atom.Flags {
	return 0 | atom.DrawCall
}

////////////////////////////////////////////////////////////////////////////////
// GlDrawArrays
////////////////////////////////////////////////////////////////////////////////
type GlDrawArrays struct {
	binary.Generate
	DrawMode   DrawMode
	FirstIndex int32
	IndexCount int32
}

func (c *GlDrawArrays) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDrawArrays(",
		c.DrawMode.String(),
		", ",
		fmt.Sprintf("first_index:%v", c.FirstIndex),
		", ",
		fmt.Sprintf("index_count:%v", c.IndexCount),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDrawArrays) API() gfxapi.API {
	return api{}
}
func (c *GlDrawArrays) TypeID() atom.TypeID {
	return 169
}
func (c *GlDrawArrays) Flags() atom.Flags {
	return 0 | atom.DrawCall
}

////////////////////////////////////////////////////////////////////////////////
// GlFlush
////////////////////////////////////////////////////////////////////////////////
type GlFlush struct {
	binary.Generate
}

func (c *GlFlush) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFlush(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFlush) API() gfxapi.API {
	return api{}
}
func (c *GlFlush) TypeID() atom.TypeID {
	return 170
}
func (c *GlFlush) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlFinish
////////////////////////////////////////////////////////////////////////////////
type GlFinish struct {
	binary.Generate
}

func (c *GlFinish) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFinish(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFinish) API() gfxapi.API {
	return api{}
}
func (c *GlFinish) TypeID() atom.TypeID {
	return 171
}
func (c *GlFinish) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetBooleanv
////////////////////////////////////////////////////////////////////////////////
type GlGetBooleanv struct {
	binary.Generate
	Param  StateVariable
	Values BoolArray
}

func (c *GlGetBooleanv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetBooleanv(",
		c.Param.String(),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetBooleanv) API() gfxapi.API {
	return api{}
}
func (c *GlGetBooleanv) TypeID() atom.TypeID {
	return 172
}
func (c *GlGetBooleanv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetFloatv
////////////////////////////////////////////////////////////////////////////////
type GlGetFloatv struct {
	binary.Generate
	Param  StateVariable
	Values F32Array
}

func (c *GlGetFloatv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetFloatv(",
		c.Param.String(),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetFloatv) API() gfxapi.API {
	return api{}
}
func (c *GlGetFloatv) TypeID() atom.TypeID {
	return 173
}
func (c *GlGetFloatv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetIntegerv
////////////////////////////////////////////////////////////////////////////////
type GlGetIntegerv struct {
	binary.Generate
	Param  StateVariable
	Values S32Array
}

func (c *GlGetIntegerv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetIntegerv(",
		c.Param.String(),
		", ",
		fmt.Sprintf("%v", c.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetIntegerv) API() gfxapi.API {
	return api{}
}
func (c *GlGetIntegerv) TypeID() atom.TypeID {
	return 174
}
func (c *GlGetIntegerv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetString
////////////////////////////////////////////////////////////////////////////////
type GlGetString struct {
	binary.Generate
	Param  StringConstant
	Result string
}

func (c *GlGetString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetString(",
		c.Param.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlGetString) API() gfxapi.API {
	return api{}
}
func (c *GlGetString) TypeID() atom.TypeID {
	return 175
}
func (c *GlGetString) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEnable
////////////////////////////////////////////////////////////////////////////////
type GlEnable struct {
	binary.Generate
	Capability Capability
}

func (c *GlEnable) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEnable(",
		c.Capability.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEnable) API() gfxapi.API {
	return api{}
}
func (c *GlEnable) TypeID() atom.TypeID {
	return 176
}
func (c *GlEnable) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDisable
////////////////////////////////////////////////////////////////////////////////
type GlDisable struct {
	binary.Generate
	Capability Capability
}

func (c *GlDisable) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDisable(",
		c.Capability.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDisable) API() gfxapi.API {
	return api{}
}
func (c *GlDisable) TypeID() atom.TypeID {
	return 177
}
func (c *GlDisable) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsEnabled
////////////////////////////////////////////////////////////////////////////////
type GlIsEnabled struct {
	binary.Generate
	Capability Capability
	Result     bool
}

func (c *GlIsEnabled) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsEnabled(",
		c.Capability.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsEnabled) API() gfxapi.API {
	return api{}
}
func (c *GlIsEnabled) TypeID() atom.TypeID {
	return 178
}
func (c *GlIsEnabled) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlMapBufferRange
////////////////////////////////////////////////////////////////////////////////
type GlMapBufferRange struct {
	binary.Generate
	Target BufferTarget
	Offset int32
	Length int32
	Access MapBufferRangeAccess
	Result memory.Pointer
}

func (c *GlMapBufferRange) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glMapBufferRange(",
		c.Target.String(),
		", ",
		fmt.Sprintf("offset:%v", c.Offset),
		", ",
		fmt.Sprintf("length:%v", c.Length),
		", ",
		c.Access.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlMapBufferRange) API() gfxapi.API {
	return api{}
}
func (c *GlMapBufferRange) TypeID() atom.TypeID {
	return 179
}
func (c *GlMapBufferRange) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlUnmapBuffer
////////////////////////////////////////////////////////////////////////////////
type GlUnmapBuffer struct {
	binary.Generate
	Target BufferTarget
}

func (c *GlUnmapBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUnmapBuffer(",
		c.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUnmapBuffer) API() gfxapi.API {
	return api{}
}
func (c *GlUnmapBuffer) TypeID() atom.TypeID {
	return 180
}
func (c *GlUnmapBuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlInvalidateFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlInvalidateFramebuffer struct {
	binary.Generate
	Target      FramebufferTarget
	Count       int32
	Attachments FramebufferAttachmentArray
}

func (c *GlInvalidateFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glInvalidateFramebuffer(",
		c.Target.String(),
		", ",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Attachments),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlInvalidateFramebuffer) API() gfxapi.API {
	return api{}
}
func (c *GlInvalidateFramebuffer) TypeID() atom.TypeID {
	return 181
}
func (c *GlInvalidateFramebuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorageMultisample
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorageMultisample struct {
	binary.Generate
	Target  RenderbufferTarget
	Samples int32
	Format  RenderbufferFormat
	Width   int32
	Height  int32
}

func (c *GlRenderbufferStorageMultisample) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glRenderbufferStorageMultisample(",
		c.Target.String(),
		", ",
		fmt.Sprintf("samples:%v", c.Samples),
		", ",
		c.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.Width),
		", ",
		fmt.Sprintf("height:%v", c.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlRenderbufferStorageMultisample) API() gfxapi.API {
	return api{}
}
func (c *GlRenderbufferStorageMultisample) TypeID() atom.TypeID {
	return 182
}
func (c *GlRenderbufferStorageMultisample) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBlitFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlBlitFramebuffer struct {
	binary.Generate
	SrcX0  int32
	SrcY0  int32
	SrcX1  int32
	SrcY1  int32
	DstX0  int32
	DstY0  int32
	DstX1  int32
	DstY1  int32
	Mask   ClearMask
	Filter TextureFilterMode
}

func (c *GlBlitFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlitFramebuffer(",
		fmt.Sprintf("srcX0:%v", c.SrcX0),
		", ",
		fmt.Sprintf("srcY0:%v", c.SrcY0),
		", ",
		fmt.Sprintf("srcX1:%v", c.SrcX1),
		", ",
		fmt.Sprintf("srcY1:%v", c.SrcY1),
		", ",
		fmt.Sprintf("dstX0:%v", c.DstX0),
		", ",
		fmt.Sprintf("dstY0:%v", c.DstY0),
		", ",
		fmt.Sprintf("dstX1:%v", c.DstX1),
		", ",
		fmt.Sprintf("dstY1:%v", c.DstY1),
		", ",
		c.Mask.String(),
		", ",
		c.Filter.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlitFramebuffer) API() gfxapi.API {
	return api{}
}
func (c *GlBlitFramebuffer) TypeID() atom.TypeID {
	return 183
}
func (c *GlBlitFramebuffer) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenQueries
////////////////////////////////////////////////////////////////////////////////
type GlGenQueries struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}

func (c *GlGenQueries) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenQueries(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenQueries) API() gfxapi.API {
	return api{}
}
func (c *GlGenQueries) TypeID() atom.TypeID {
	return 184
}
func (c *GlGenQueries) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBeginQuery
////////////////////////////////////////////////////////////////////////////////
type GlBeginQuery struct {
	binary.Generate
	Target QueryTarget
	Query  QueryId
}

func (c *GlBeginQuery) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBeginQuery(",
		c.Target.String(),
		", ",
		fmt.Sprintf("query:%v", c.Query),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBeginQuery) API() gfxapi.API {
	return api{}
}
func (c *GlBeginQuery) TypeID() atom.TypeID {
	return 185
}
func (c *GlBeginQuery) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEndQuery
////////////////////////////////////////////////////////////////////////////////
type GlEndQuery struct {
	binary.Generate
	Target QueryTarget
}

func (c *GlEndQuery) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEndQuery(",
		c.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEndQuery) API() gfxapi.API {
	return api{}
}
func (c *GlEndQuery) TypeID() atom.TypeID {
	return 186
}
func (c *GlEndQuery) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueries
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueries struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}

func (c *GlDeleteQueries) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteQueries(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteQueries) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteQueries) TypeID() atom.TypeID {
	return 187
}
func (c *GlDeleteQueries) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsQuery
////////////////////////////////////////////////////////////////////////////////
type GlIsQuery struct {
	binary.Generate
	Query  QueryId
	Result bool
}

func (c *GlIsQuery) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsQuery(",
		fmt.Sprintf("query:%v", c.Query),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsQuery) API() gfxapi.API {
	return api{}
}
func (c *GlIsQuery) TypeID() atom.TypeID {
	return 188
}
func (c *GlIsQuery) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryiv struct {
	binary.Generate
	Target    QueryTarget
	Parameter QueryParameter
	Value     int32
}

func (c *GlGetQueryiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryiv(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryiv) API() gfxapi.API {
	return api{}
}
func (c *GlGetQueryiv) TypeID() atom.TypeID {
	return 189
}
func (c *GlGetQueryiv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuiv struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     uint32
}

func (c *GlGetQueryObjectuiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectuiv(",
		fmt.Sprintf("query:%v", c.Query),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectuiv) API() gfxapi.API {
	return api{}
}
func (c *GlGetQueryObjectuiv) TypeID() atom.TypeID {
	return 190
}
func (c *GlGetQueryObjectuiv) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGenQueriesEXT
////////////////////////////////////////////////////////////////////////////////
type GlGenQueriesEXT struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}

func (c *GlGenQueriesEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenQueriesEXT(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenQueriesEXT) API() gfxapi.API {
	return api{}
}
func (c *GlGenQueriesEXT) TypeID() atom.TypeID {
	return 191
}
func (c *GlGenQueriesEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlBeginQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlBeginQueryEXT struct {
	binary.Generate
	Target QueryTarget
	Query  QueryId
}

func (c *GlBeginQueryEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBeginQueryEXT(",
		c.Target.String(),
		", ",
		fmt.Sprintf("query:%v", c.Query),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBeginQueryEXT) API() gfxapi.API {
	return api{}
}
func (c *GlBeginQueryEXT) TypeID() atom.TypeID {
	return 192
}
func (c *GlBeginQueryEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlEndQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlEndQueryEXT struct {
	binary.Generate
	Target QueryTarget
}

func (c *GlEndQueryEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEndQueryEXT(",
		c.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEndQueryEXT) API() gfxapi.API {
	return api{}
}
func (c *GlEndQueryEXT) TypeID() atom.TypeID {
	return 193
}
func (c *GlEndQueryEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueriesEXT
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueriesEXT struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}

func (c *GlDeleteQueriesEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteQueriesEXT(",
		fmt.Sprintf("count:%v", c.Count),
		", ",
		fmt.Sprintf("%v", c.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteQueriesEXT) API() gfxapi.API {
	return api{}
}
func (c *GlDeleteQueriesEXT) TypeID() atom.TypeID {
	return 194
}
func (c *GlDeleteQueriesEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlIsQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlIsQueryEXT struct {
	binary.Generate
	Query  QueryId
	Result bool
}

func (c *GlIsQueryEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsQueryEXT(",
		fmt.Sprintf("query:%v", c.Query),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Result))
	return strings.Join(parts, "")
}
func (c *GlIsQueryEXT) API() gfxapi.API {
	return api{}
}
func (c *GlIsQueryEXT) TypeID() atom.TypeID {
	return 195
}
func (c *GlIsQueryEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlQueryCounterEXT
////////////////////////////////////////////////////////////////////////////////
type GlQueryCounterEXT struct {
	binary.Generate
	Query  QueryId
	Target QueryTarget
}

func (c *GlQueryCounterEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glQueryCounterEXT(",
		fmt.Sprintf("query:%v", c.Query),
		", ",
		c.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlQueryCounterEXT) API() gfxapi.API {
	return api{}
}
func (c *GlQueryCounterEXT) TypeID() atom.TypeID {
	return 196
}
func (c *GlQueryCounterEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryivEXT struct {
	binary.Generate
	Target    QueryTarget
	Parameter QueryParameter
	Value     int32
}

func (c *GlGetQueryivEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryivEXT(",
		c.Target.String(),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryivEXT) API() gfxapi.API {
	return api{}
}
func (c *GlGetQueryivEXT) TypeID() atom.TypeID {
	return 197
}
func (c *GlGetQueryivEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectivEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     int32
}

func (c *GlGetQueryObjectivEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectivEXT(",
		fmt.Sprintf("query:%v", c.Query),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectivEXT) API() gfxapi.API {
	return api{}
}
func (c *GlGetQueryObjectivEXT) TypeID() atom.TypeID {
	return 198
}
func (c *GlGetQueryObjectivEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuivEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     uint32
}

func (c *GlGetQueryObjectuivEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectuivEXT(",
		fmt.Sprintf("query:%v", c.Query),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectuivEXT) API() gfxapi.API {
	return api{}
}
func (c *GlGetQueryObjectuivEXT) TypeID() atom.TypeID {
	return 199
}
func (c *GlGetQueryObjectuivEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjecti64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjecti64vEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     int64
}

func (c *GlGetQueryObjecti64vEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjecti64vEXT(",
		fmt.Sprintf("query:%v", c.Query),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjecti64vEXT) API() gfxapi.API {
	return api{}
}
func (c *GlGetQueryObjecti64vEXT) TypeID() atom.TypeID {
	return 200
}
func (c *GlGetQueryObjecti64vEXT) Flags() atom.Flags {
	return 0
}

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectui64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectui64vEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     uint64
}

func (c *GlGetQueryObjectui64vEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectui64vEXT(",
		fmt.Sprintf("query:%v", c.Query),
		", ",
		c.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectui64vEXT) API() gfxapi.API {
	return api{}
}
func (c *GlGetQueryObjectui64vEXT) TypeID() atom.TypeID {
	return 201
}
func (c *GlGetQueryObjectui64vEXT) Flags() atom.Flags {
	return 0
}

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
	Data      memory.Memory
	Size      uint32
	Format    ImageTexelFormat
}

func (c *Image) Init() {
}
func (c *Image) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class FramebufferAttachable
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachable interface {
}

////////////////////////////////////////////////////////////////////////////////
// class Renderbuffer
////////////////////////////////////////////////////////////////////////////////
type Renderbuffer struct {
	binary.Generate
	CreatedAt atom.ID
	Width     int32
	Height    int32
	Data      memory.Memory
	Format    RenderbufferFormat
	// FramebufferAttachable
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
	Texture2D     Image_s32Map
	Cubemap       CubemapLevel_s32Map
	MagFilter     TextureFilterMode
	MinFilter     TextureFilterMode
	WrapS         TextureWrapMode
	WrapT         TextureWrapMode
	SwizzleR      TexelComponent
	SwizzleG      TexelComponent
	SwizzleB      TexelComponent
	SwizzleA      TexelComponent
	MaxAnisotropy float32
	// FramebufferAttachable
}

func (c *Texture) Init() {
	c.Texture2D = make(Image_s32Map)
	c.Cubemap = make(CubemapLevel_s32Map)
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
	Faces     Image_CubeMapImageTargetMap
}

func (c *CubemapLevel) Init() {
	c.Faces = make(Image_CubeMapImageTargetMap)
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
	Attachments FramebufferAttachmentInfo_FramebufferAttachmentMap
}

func (c *Framebuffer) Init() {
	c.Attachments = make(FramebufferAttachmentInfo_FramebufferAttachmentMap)
}
func (c *Framebuffer) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Buffer
////////////////////////////////////////////////////////////////////////////////
type Buffer struct {
	binary.Generate
	CreatedAt atom.ID
	Data      memory.Memory
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
	Binary    memory.Memory
	Compiled  bool
	Deletable bool
	InfoLog   string
	Source    []string
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
	Name        string
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
	Shaders           ShaderId_ShaderTypeMap
	Linked            bool
	Binary            memory.Memory
	AttributeBindings AttributeLocation_CharBufferMap
	Attributes        VertexAttribute_s32Map
	Uniforms          Uniform_UniformLocationMap
	InfoLog           string
}

func (c *Program) Init() {
	c.Shaders = make(ShaderId_ShaderTypeMap)
	c.AttributeBindings = make(AttributeLocation_CharBufferMap)
	c.Attributes = make(VertexAttribute_s32Map)
	c.Uniforms = make(Uniform_UniformLocationMap)
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
	Pointer    memory.Pointer
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
	StencilMask          U32_FaceModeMap
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
	c.StencilMask = make(U32_FaceModeMap)
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
	Renderbuffers RenderbufferPtr_RenderbufferIdMap
	Textures      TexturePtr_TextureIdMap
	Framebuffers  FramebufferPtr_FramebufferIdMap
	Buffers       BufferPtr_BufferIdMap
	Shaders       ShaderPtr_ShaderIdMap
	Programs      ProgramPtr_ProgramIdMap
	VertexArrays  VertexArrayPtr_VertexArrayIdMap
	Queries       QueryPtr_QueryIdMap
}

func (c *Objects) Init() {
	c.Renderbuffers = make(RenderbufferPtr_RenderbufferIdMap)
	c.Textures = make(TexturePtr_TextureIdMap)
	c.Framebuffers = make(FramebufferPtr_FramebufferIdMap)
	c.Buffers = make(BufferPtr_BufferIdMap)
	c.Shaders = make(ShaderPtr_ShaderIdMap)
	c.Programs = make(ProgramPtr_ProgramIdMap)
	c.VertexArrays = make(VertexArrayPtr_VertexArrayIdMap)
	c.Queries = make(QueryPtr_QueryIdMap)
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
	BoundFramebuffers     FramebufferId_FramebufferTargetMap
	BoundRenderbuffers    RenderbufferId_RenderbufferTargetMap
	BoundBuffers          BufferId_BufferTargetMap
	BoundProgram          ProgramId
	BoundVertexArray      VertexArrayId
	VertexAttributeArrays VertexAttributeArrayPtr_AttributeLocationMap
	TextureUnits          TextureId_TextureTargetMap_TextureUnitMap
	ActiveTextureUnit     TextureUnit
	Capabilities          Bool_CapabilityMap
	GenerateMipmapHint    HintMode
	PixelStorage          S32_PixelStoreParameterMap
	Instances             Objects
}

func (c *Context) Init() {
	c.Blending.Init()
	c.Rasterizing.Init()
	c.Clearing.Init()
	c.BoundFramebuffers = make(FramebufferId_FramebufferTargetMap)
	c.BoundRenderbuffers = make(RenderbufferId_RenderbufferTargetMap)
	c.BoundBuffers = make(BufferId_BufferTargetMap)
	c.VertexAttributeArrays = make(VertexAttributeArrayPtr_AttributeLocationMap)
	c.TextureUnits = make(TextureId_TextureTargetMap_TextureUnitMap)
	c.ActiveTextureUnit = TextureUnit_GL_TEXTURE0
	c.Capabilities = make(Bool_CapabilityMap)
	c.GenerateMipmapHint = HintMode_GL_DONT_CARE
	c.PixelStorage = make(S32_PixelStoreParameterMap)
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

func (v DrawMode) String() string {
	switch v {
	case 2:
		return "GL_LINE_LOOP"
	case 3:
		return "GL_LINE_STRIP"
	case 1:
		return "GL_LINES"
	case 0:
		return "GL_POINTS"
	case 6:
		return "GL_TRIANGLE_FAN"
	case 5:
		return "GL_TRIANGLE_STRIP"
	case 4:
		return "GL_TRIANGLES"
	default:
		return fmt.Sprintf("DrawMode<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum IndicesType
////////////////////////////////////////////////////////////////////////////////
type IndicesType uint32

const (
	IndicesType_GL_UNSIGNED_BYTE  = IndicesType(5121)
	IndicesType_GL_UNSIGNED_SHORT = IndicesType(5123)
	IndicesType_GL_UNSIGNED_INT   = IndicesType(5125)
)

func (v IndicesType) String() string {
	switch v {
	case 5121:
		return "GL_UNSIGNED_BYTE"
	case 5123:
		return "GL_UNSIGNED_SHORT"
	case 5125:
		return "GL_UNSIGNED_INT"
	default:
		return fmt.Sprintf("IndicesType<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureTarget_GLES_1_1
////////////////////////////////////////////////////////////////////////////////
type TextureTarget_GLES_1_1 uint32

const (
	TextureTarget_GLES_1_1_GL_TEXTURE_2D = TextureTarget_GLES_1_1(3553)
)

func (v TextureTarget_GLES_1_1) String() string {
	switch v {
	case 3553:
		return "GL_TEXTURE_2D"
	default:
		return fmt.Sprintf("TextureTarget_GLES_1_1<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureTarget_GLES_2_0
////////////////////////////////////////////////////////////////////////////////
type TextureTarget_GLES_2_0 uint32

const (
	TextureTarget_GLES_2_0_GL_TEXTURE_CUBE_MAP = TextureTarget_GLES_2_0(34067)
)

func (v TextureTarget_GLES_2_0) String() string {
	switch v {
	case 34067:
		return "GL_TEXTURE_CUBE_MAP"
	default:
		return fmt.Sprintf("TextureTarget_GLES_2_0<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureTarget_OES_EGL_image_external
////////////////////////////////////////////////////////////////////////////////
type TextureTarget_OES_EGL_image_external uint32

const (
	TextureTarget_OES_EGL_image_external_GL_TEXTURE_EXTERNAL_OES = TextureTarget_OES_EGL_image_external(36197)
)

func (v TextureTarget_OES_EGL_image_external) String() string {
	switch v {
	case 36197:
		return "GL_TEXTURE_EXTERNAL_OES"
	default:
		return fmt.Sprintf("TextureTarget_OES_EGL_image_external<0x%.4x>", uint32(v))
	}
}

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

func (v TextureTarget) String() string {
	switch v {
	// TextureTarget_GLES_1_1
	case 3553:
		return "GL_TEXTURE_2D"
	// TextureTarget_GLES_2_0
	case 34067:
		return "GL_TEXTURE_CUBE_MAP"
	// TextureTarget_OES_EGL_image_external
	case 36197:
		return "GL_TEXTURE_EXTERNAL_OES"
	default:
		return fmt.Sprintf("TextureTarget<0x%.4x>", uint32(v))
	}
}

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

func (v CubeMapImageTarget) String() string {
	switch v {
	case 34070:
		return "GL_TEXTURE_CUBE_MAP_NEGATIVE_X"
	case 34072:
		return "GL_TEXTURE_CUBE_MAP_NEGATIVE_Y"
	case 34074:
		return "GL_TEXTURE_CUBE_MAP_NEGATIVE_Z"
	case 34069:
		return "GL_TEXTURE_CUBE_MAP_POSITIVE_X"
	case 34071:
		return "GL_TEXTURE_CUBE_MAP_POSITIVE_Y"
	case 34073:
		return "GL_TEXTURE_CUBE_MAP_POSITIVE_Z"
	default:
		return fmt.Sprintf("CubeMapImageTarget<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum Texture2DImageTarget
////////////////////////////////////////////////////////////////////////////////
type Texture2DImageTarget uint32

const (
	Texture2DImageTarget_GL_TEXTURE_2D = Texture2DImageTarget(3553)
)

func (v Texture2DImageTarget) String() string {
	switch v {
	case 3553:
		return "GL_TEXTURE_2D"
	default:
		return fmt.Sprintf("Texture2DImageTarget<0x%.4x>", uint32(v))
	}
}

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

func (v TextureImageTarget) String() string {
	switch v {
	// CubeMapImageTarget
	case 34070:
		return "GL_TEXTURE_CUBE_MAP_NEGATIVE_X"
	case 34072:
		return "GL_TEXTURE_CUBE_MAP_NEGATIVE_Y"
	case 34074:
		return "GL_TEXTURE_CUBE_MAP_NEGATIVE_Z"
	case 34069:
		return "GL_TEXTURE_CUBE_MAP_POSITIVE_X"
	case 34071:
		return "GL_TEXTURE_CUBE_MAP_POSITIVE_Y"
	case 34073:
		return "GL_TEXTURE_CUBE_MAP_POSITIVE_Z"
	// Texture2DImageTarget
	case 3553:
		return "GL_TEXTURE_2D"
	default:
		return fmt.Sprintf("TextureImageTarget<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum BaseTexelFormat
////////////////////////////////////////////////////////////////////////////////
type BaseTexelFormat uint32

const (
	BaseTexelFormat_GL_ALPHA = BaseTexelFormat(6406)
	BaseTexelFormat_GL_RGB   = BaseTexelFormat(6407)
	BaseTexelFormat_GL_RGBA  = BaseTexelFormat(6408)
)

func (v BaseTexelFormat) String() string {
	switch v {
	case 6406:
		return "GL_ALPHA"
	case 6407:
		return "GL_RGB"
	case 6408:
		return "GL_RGBA"
	default:
		return fmt.Sprintf("BaseTexelFormat<0x%.4x>", uint32(v))
	}
}

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

func (v TexelFormat_GLES_1_1) String() string {
	switch v {
	case 6409:
		return "GL_LUMINANCE"
	case 6410:
		return "GL_LUMINANCE_ALPHA"
	// BaseTexelFormat
	case 6406:
		return "GL_ALPHA"
	case 6407:
		return "GL_RGB"
	case 6408:
		return "GL_RGBA"
	default:
		return fmt.Sprintf("TexelFormat_GLES_1_1<0x%.4x>", uint32(v))
	}
}

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

func (v TexelFormat_GLES_3_0) String() string {
	switch v {
	case 6403:
		return "GL_RED"
	case 36244:
		return "GL_RED_INTEGER"
	case 33319:
		return "GL_RG"
	case 33320:
		return "GL_RG_INTEGER"
	case 36248:
		return "GL_RGB_INTEGER"
	case 36249:
		return "GL_RGBA_INTEGER"
	case 6402:
		return "GL_DEPTH_COMPONENT"
	case 33189:
		return "GL_DEPTH_COMPONENT16"
	case 34041:
		return "GL_DEPTH_STENCIL"
	case 35056:
		return "GL_DEPTH24_STENCIL8"
	default:
		return fmt.Sprintf("TexelFormat_GLES_3_0<0x%.4x>", uint32(v))
	}
}

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

func (v TexelFormat) String() string {
	switch v {
	// TexelFormat_GLES_1_1
	case 6409:
		return "GL_LUMINANCE"
	case 6410:
		return "GL_LUMINANCE_ALPHA"
	// BaseTexelFormat
	case 6406:
		return "GL_ALPHA"
	case 6407:
		return "GL_RGB"
	case 6408:
		return "GL_RGBA"
	// TexelFormat_GLES_3_0
	case 6403:
		return "GL_RED"
	case 36244:
		return "GL_RED_INTEGER"
	case 33319:
		return "GL_RG"
	case 33320:
		return "GL_RG_INTEGER"
	case 36248:
		return "GL_RGB_INTEGER"
	case 36249:
		return "GL_RGBA_INTEGER"
	case 6402:
		return "GL_DEPTH_COMPONENT"
	case 33189:
		return "GL_DEPTH_COMPONENT16"
	case 34041:
		return "GL_DEPTH_STENCIL"
	case 35056:
		return "GL_DEPTH24_STENCIL8"
	default:
		return fmt.Sprintf("TexelFormat<0x%.4x>", uint32(v))
	}
}

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

func (v RenderbufferFormat) String() string {
	switch v {
	case 32854:
		return "GL_RGBA4"
	case 32855:
		return "GL_RGB5_A1"
	case 36194:
		return "GL_RGB565"
	case 32856:
		return "GL_RGBA8"
	case 33189:
		return "GL_DEPTH_COMPONENT16"
	case 36168:
		return "GL_STENCIL_INDEX8"
	default:
		return fmt.Sprintf("RenderbufferFormat<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum Type_ARB_half_float_vertex
////////////////////////////////////////////////////////////////////////////////
type Type_ARB_half_float_vertex uint32

const (
	Type_ARB_half_float_vertex_GL_ARB_half_float_vertex = Type_ARB_half_float_vertex(5131)
)

func (v Type_ARB_half_float_vertex) String() string {
	switch v {
	case 5131:
		return "GL_ARB_half_float_vertex"
	default:
		return fmt.Sprintf("Type_ARB_half_float_vertex<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum Type_OES_vertex_half_float
////////////////////////////////////////////////////////////////////////////////
type Type_OES_vertex_half_float uint32

const (
	Type_OES_vertex_half_float_GL_HALF_FLOAT_OES = Type_OES_vertex_half_float(36193)
)

func (v Type_OES_vertex_half_float) String() string {
	switch v {
	case 36193:
		return "GL_HALF_FLOAT_OES"
	default:
		return fmt.Sprintf("Type_OES_vertex_half_float<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
////////////////////////////////////////////////////////////////////////////////
type CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture uint32

const (
	CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture_GL_ETC1_RGB8_OES = CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture(36196)
)

func (v CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture) String() string {
	switch v {
	case 36196:
		return "GL_ETC1_RGB8_OES"
	default:
		return fmt.Sprintf("CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum CompressedTexelFormat_AMD_compressed_ATC_texture
////////////////////////////////////////////////////////////////////////////////
type CompressedTexelFormat_AMD_compressed_ATC_texture uint32

const (
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGB_AMD                     = CompressedTexelFormat_AMD_compressed_ATC_texture(35986)
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD     = CompressedTexelFormat_AMD_compressed_ATC_texture(35987)
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = CompressedTexelFormat_AMD_compressed_ATC_texture(34798)
)

func (v CompressedTexelFormat_AMD_compressed_ATC_texture) String() string {
	switch v {
	case 35986:
		return "GL_ATC_RGB_AMD"
	case 35987:
		return "GL_ATC_RGBA_EXPLICIT_ALPHA_AMD"
	case 34798:
		return "GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD"
	default:
		return fmt.Sprintf("CompressedTexelFormat_AMD_compressed_ATC_texture<0x%.4x>", uint32(v))
	}
}

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

func (v CompressedTexelFormat) String() string {
	switch v {
	// CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
	case 36196:
		return "GL_ETC1_RGB8_OES"
	// CompressedTexelFormat_AMD_compressed_ATC_texture
	case 35986:
		return "GL_ATC_RGB_AMD"
	case 35987:
		return "GL_ATC_RGBA_EXPLICIT_ALPHA_AMD"
	case 34798:
		return "GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD"
	default:
		return fmt.Sprintf("CompressedTexelFormat<0x%.4x>", uint32(v))
	}
}

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

func (v ImageTexelFormat) String() string {
	switch v {
	// TexelFormat
	// TexelFormat_GLES_1_1
	case 6409:
		return "GL_LUMINANCE"
	case 6410:
		return "GL_LUMINANCE_ALPHA"
	// BaseTexelFormat
	case 6406:
		return "GL_ALPHA"
	case 6407:
		return "GL_RGB"
	case 6408:
		return "GL_RGBA"
	// TexelFormat_GLES_3_0
	case 6403:
		return "GL_RED"
	case 36244:
		return "GL_RED_INTEGER"
	case 33319:
		return "GL_RG"
	case 33320:
		return "GL_RG_INTEGER"
	case 36248:
		return "GL_RGB_INTEGER"
	case 36249:
		return "GL_RGBA_INTEGER"
	case 6402:
		return "GL_DEPTH_COMPONENT"
	case 33189:
		return "GL_DEPTH_COMPONENT16"
	case 34041:
		return "GL_DEPTH_STENCIL"
	case 35056:
		return "GL_DEPTH24_STENCIL8"
	// CompressedTexelFormat
	// CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
	case 36196:
		return "GL_ETC1_RGB8_OES"
	// CompressedTexelFormat_AMD_compressed_ATC_texture
	case 35986:
		return "GL_ATC_RGB_AMD"
	case 35987:
		return "GL_ATC_RGBA_EXPLICIT_ALPHA_AMD"
	case 34798:
		return "GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD"
	default:
		return fmt.Sprintf("ImageTexelFormat<0x%.4x>", uint32(v))
	}
}

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

func (v TexelType) String() string {
	switch v {
	case 5121:
		return "GL_UNSIGNED_BYTE"
	case 5123:
		return "GL_UNSIGNED_SHORT"
	case 5125:
		return "GL_UNSIGNED_INT"
	case 5126:
		return "GL_FLOAT"
	case 32819:
		return "GL_UNSIGNED_SHORT_4_4_4_4"
	case 32820:
		return "GL_UNSIGNED_SHORT_5_5_5_1"
	case 33635:
		return "GL_UNSIGNED_SHORT_5_6_5"
	case 34042:
		return "GL_UNSIGNED_INT_24_8"
	default:
		return fmt.Sprintf("TexelType<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferAttachment
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachment uint32

const (
	FramebufferAttachment_GL_COLOR_ATTACHMENT0  = FramebufferAttachment(36064)
	FramebufferAttachment_GL_DEPTH_ATTACHMENT   = FramebufferAttachment(36096)
	FramebufferAttachment_GL_STENCIL_ATTACHMENT = FramebufferAttachment(36128)
)

func (v FramebufferAttachment) String() string {
	switch v {
	case 36064:
		return "GL_COLOR_ATTACHMENT0"
	case 36096:
		return "GL_DEPTH_ATTACHMENT"
	case 36128:
		return "GL_STENCIL_ATTACHMENT"
	default:
		return fmt.Sprintf("FramebufferAttachment<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferAttachmentType
////////////////////////////////////////////////////////////////////////////////
type FramebufferAttachmentType uint32

const (
	FramebufferAttachmentType_GL_NONE         = FramebufferAttachmentType(0)
	FramebufferAttachmentType_GL_RENDERBUFFER = FramebufferAttachmentType(36161)
	FramebufferAttachmentType_GL_TEXTURE      = FramebufferAttachmentType(5890)
)

func (v FramebufferAttachmentType) String() string {
	switch v {
	case 0:
		return "GL_NONE"
	case 36161:
		return "GL_RENDERBUFFER"
	case 5890:
		return "GL_TEXTURE"
	default:
		return fmt.Sprintf("FramebufferAttachmentType<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferTarget_GLES_2_0
////////////////////////////////////////////////////////////////////////////////
type FramebufferTarget_GLES_2_0 uint32

const (
	FramebufferTarget_GLES_2_0_GL_FRAMEBUFFER = FramebufferTarget_GLES_2_0(36160)
)

func (v FramebufferTarget_GLES_2_0) String() string {
	switch v {
	case 36160:
		return "GL_FRAMEBUFFER"
	default:
		return fmt.Sprintf("FramebufferTarget_GLES_2_0<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum FramebufferTarget_GLES_3_1
////////////////////////////////////////////////////////////////////////////////
type FramebufferTarget_GLES_3_1 uint32

const (
	FramebufferTarget_GLES_3_1_GL_READ_FRAMEBUFFER = FramebufferTarget_GLES_3_1(36008)
	FramebufferTarget_GLES_3_1_GL_DRAW_FRAMEBUFFER = FramebufferTarget_GLES_3_1(36009)
)

func (v FramebufferTarget_GLES_3_1) String() string {
	switch v {
	case 36008:
		return "GL_READ_FRAMEBUFFER"
	case 36009:
		return "GL_DRAW_FRAMEBUFFER"
	default:
		return fmt.Sprintf("FramebufferTarget_GLES_3_1<0x%.4x>", uint32(v))
	}
}

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

func (v FramebufferTarget) String() string {
	switch v {
	// FramebufferTarget_GLES_2_0
	case 36160:
		return "GL_FRAMEBUFFER"
	// FramebufferTarget_GLES_3_1
	case 36008:
		return "GL_READ_FRAMEBUFFER"
	case 36009:
		return "GL_DRAW_FRAMEBUFFER"
	default:
		return fmt.Sprintf("FramebufferTarget<0x%.4x>", uint32(v))
	}
}

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

func (v FramebufferAttachmentParameter) String() string {
	switch v {
	case 36048:
		return "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE"
	case 36049:
		return "GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME"
	case 36050:
		return "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL"
	case 36051:
		return "GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE"
	default:
		return fmt.Sprintf("FramebufferAttachmentParameter<0x%.4x>", uint32(v))
	}
}

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

func (v FramebufferStatus) String() string {
	switch v {
	case 36053:
		return "GL_FRAMEBUFFER_COMPLETE"
	case 36054:
		return "GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT"
	case 36055:
		return "GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT"
	case 36057:
		return "GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS"
	case 36061:
		return "GL_FRAMEBUFFER_UNSUPPORTED"
	default:
		return fmt.Sprintf("FramebufferStatus<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum RenderbufferTarget
////////////////////////////////////////////////////////////////////////////////
type RenderbufferTarget uint32

const (
	RenderbufferTarget_GL_RENDERBUFFER = RenderbufferTarget(36161)
)

func (v RenderbufferTarget) String() string {
	switch v {
	case 36161:
		return "GL_RENDERBUFFER"
	default:
		return fmt.Sprintf("RenderbufferTarget<0x%.4x>", uint32(v))
	}
}

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

func (v RenderbufferParameter) String() string {
	switch v {
	case 36162:
		return "GL_RENDERBUFFER_WIDTH"
	case 36163:
		return "GL_RENDERBUFFER_HEIGHT"
	case 36164:
		return "GL_RENDERBUFFER_INTERNAL_FORMAT"
	case 36176:
		return "GL_RENDERBUFFER_RED_SIZE"
	case 36177:
		return "GL_RENDERBUFFER_GREEN_SIZE"
	case 36178:
		return "GL_RENDERBUFFER_BLUE_SIZE"
	case 36179:
		return "GL_RENDERBUFFER_ALPHA_SIZE"
	case 36180:
		return "GL_RENDERBUFFER_DEPTH_SIZE"
	case 36181:
		return "GL_RENDERBUFFER_STENCIL_SIZE"
	default:
		return fmt.Sprintf("RenderbufferParameter<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum BufferParameter
////////////////////////////////////////////////////////////////////////////////
type BufferParameter uint32

const (
	BufferParameter_GL_BUFFER_SIZE  = BufferParameter(34660)
	BufferParameter_GL_BUFFER_USAGE = BufferParameter(34661)
)

func (v BufferParameter) String() string {
	switch v {
	case 34660:
		return "GL_BUFFER_SIZE"
	case 34661:
		return "GL_BUFFER_USAGE"
	default:
		return fmt.Sprintf("BufferParameter<0x%.4x>", uint32(v))
	}
}

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

func (v TextureUnit) String() string {
	switch v {
	case 33984:
		return "GL_TEXTURE0"
	case 33985:
		return "GL_TEXTURE1"
	case 33986:
		return "GL_TEXTURE2"
	case 33987:
		return "GL_TEXTURE3"
	case 33988:
		return "GL_TEXTURE4"
	case 33989:
		return "GL_TEXTURE5"
	case 33990:
		return "GL_TEXTURE6"
	case 33991:
		return "GL_TEXTURE7"
	case 33992:
		return "GL_TEXTURE8"
	case 33993:
		return "GL_TEXTURE9"
	case 33994:
		return "GL_TEXTURE10"
	case 33995:
		return "GL_TEXTURE11"
	case 33996:
		return "GL_TEXTURE12"
	case 33997:
		return "GL_TEXTURE13"
	case 33998:
		return "GL_TEXTURE14"
	case 33999:
		return "GL_TEXTURE15"
	case 34000:
		return "GL_TEXTURE16"
	case 34001:
		return "GL_TEXTURE17"
	case 34002:
		return "GL_TEXTURE18"
	case 34003:
		return "GL_TEXTURE19"
	case 34004:
		return "GL_TEXTURE20"
	case 34005:
		return "GL_TEXTURE21"
	case 34006:
		return "GL_TEXTURE22"
	case 34007:
		return "GL_TEXTURE23"
	case 34008:
		return "GL_TEXTURE24"
	case 34009:
		return "GL_TEXTURE25"
	case 34010:
		return "GL_TEXTURE26"
	case 34011:
		return "GL_TEXTURE27"
	case 34012:
		return "GL_TEXTURE28"
	case 34013:
		return "GL_TEXTURE29"
	case 34014:
		return "GL_TEXTURE30"
	case 34015:
		return "GL_TEXTURE31"
	default:
		return fmt.Sprintf("TextureUnit<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum BufferUsage
////////////////////////////////////////////////////////////////////////////////
type BufferUsage uint32

const (
	BufferUsage_GL_DYNAMIC_DRAW = BufferUsage(35048)
	BufferUsage_GL_STATIC_DRAW  = BufferUsage(35044)
	BufferUsage_GL_STREAM_DRAW  = BufferUsage(35040)
)

func (v BufferUsage) String() string {
	switch v {
	case 35048:
		return "GL_DYNAMIC_DRAW"
	case 35044:
		return "GL_STATIC_DRAW"
	case 35040:
		return "GL_STREAM_DRAW"
	default:
		return fmt.Sprintf("BufferUsage<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum ShaderType
////////////////////////////////////////////////////////////////////////////////
type ShaderType uint32

const (
	ShaderType_GL_VERTEX_SHADER   = ShaderType(35633)
	ShaderType_GL_FRAGMENT_SHADER = ShaderType(35632)
)

func (v ShaderType) String() string {
	switch v {
	case 35633:
		return "GL_VERTEX_SHADER"
	case 35632:
		return "GL_FRAGMENT_SHADER"
	default:
		return fmt.Sprintf("ShaderType<0x%.4x>", uint32(v))
	}
}

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

func (v StateVariable_GLES_2_0) String() string {
	switch v {
	case 34016:
		return "GL_ACTIVE_TEXTURE"
	case 33902:
		return "GL_ALIASED_LINE_WIDTH_RANGE"
	case 33901:
		return "GL_ALIASED_POINT_SIZE_RANGE"
	case 3413:
		return "GL_ALPHA_BITS"
	case 34964:
		return "GL_ARRAY_BUFFER_BINDING"
	case 3042:
		return "GL_BLEND"
	case 32773:
		return "GL_BLEND_COLOR"
	case 32970:
		return "GL_BLEND_DST_ALPHA"
	case 32968:
		return "GL_BLEND_DST_RGB"
	case 34877:
		return "GL_BLEND_EQUATION_ALPHA"
	case 32777:
		return "GL_BLEND_EQUATION_RGB"
	case 32971:
		return "GL_BLEND_SRC_ALPHA"
	case 32969:
		return "GL_BLEND_SRC_RGB"
	case 3412:
		return "GL_BLUE_BITS"
	case 3106:
		return "GL_COLOR_CLEAR_VALUE"
	case 3107:
		return "GL_COLOR_WRITEMASK"
	case 34467:
		return "GL_COMPRESSED_TEXTURE_FORMATS"
	case 2884:
		return "GL_CULL_FACE"
	case 2885:
		return "GL_CULL_FACE_MODE"
	case 35725:
		return "GL_CURRENT_PROGRAM"
	case 3414:
		return "GL_DEPTH_BITS"
	case 2931:
		return "GL_DEPTH_CLEAR_VALUE"
	case 2932:
		return "GL_DEPTH_FUNC"
	case 2928:
		return "GL_DEPTH_RANGE"
	case 2929:
		return "GL_DEPTH_TEST"
	case 2930:
		return "GL_DEPTH_WRITEMASK"
	case 3024:
		return "GL_DITHER"
	case 34965:
		return "GL_ELEMENT_ARRAY_BUFFER_BINDING"
	case 36006:
		return "GL_FRAMEBUFFER_BINDING"
	case 2886:
		return "GL_FRONT_FACE"
	case 33170:
		return "GL_GENERATE_MIPMAP_HINT"
	case 3411:
		return "GL_GREEN_BITS"
	case 35739:
		return "GL_IMPLEMENTATION_COLOR_READ_FORMAT"
	case 35738:
		return "GL_IMPLEMENTATION_COLOR_READ_TYPE"
	case 2849:
		return "GL_LINE_WIDTH"
	case 35661:
		return "GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS"
	case 34076:
		return "GL_MAX_CUBE_MAP_TEXTURE_SIZE"
	case 36349:
		return "GL_MAX_FRAGMENT_UNIFORM_VECTORS"
	case 34024:
		return "GL_MAX_RENDERBUFFER_SIZE"
	case 34930:
		return "GL_MAX_TEXTURE_IMAGE_UNITS"
	case 3379:
		return "GL_MAX_TEXTURE_SIZE"
	case 36348:
		return "GL_MAX_VARYING_VECTORS"
	case 34921:
		return "GL_MAX_VERTEX_ATTRIBS"
	case 35660:
		return "GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS"
	case 36347:
		return "GL_MAX_VERTEX_UNIFORM_VECTORS"
	case 3386:
		return "GL_MAX_VIEWPORT_DIMS"
	case 34466:
		return "GL_NUM_COMPRESSED_TEXTURE_FORMATS"
	case 36345:
		return "GL_NUM_SHADER_BINARY_FORMATS"
	case 3333:
		return "GL_PACK_ALIGNMENT"
	case 32824:
		return "GL_POLYGON_OFFSET_FACTOR"
	case 32823:
		return "GL_POLYGON_OFFSET_FILL"
	case 10752:
		return "GL_POLYGON_OFFSET_UNITS"
	case 3410:
		return "GL_RED_BITS"
	case 36007:
		return "GL_RENDERBUFFER_BINDING"
	case 32926:
		return "GL_SAMPLE_ALPHA_TO_COVERAGE"
	case 32936:
		return "GL_SAMPLE_BUFFERS"
	case 32928:
		return "GL_SAMPLE_COVERAGE"
	case 32939:
		return "GL_SAMPLE_COVERAGE_INVERT"
	case 32938:
		return "GL_SAMPLE_COVERAGE_VALUE"
	case 32937:
		return "GL_SAMPLES"
	case 3088:
		return "GL_SCISSOR_BOX"
	case 3089:
		return "GL_SCISSOR_TEST"
	case 36344:
		return "GL_SHADER_BINARY_FORMATS"
	case 36346:
		return "GL_SHADER_COMPILER"
	case 34817:
		return "GL_STENCIL_BACK_FAIL"
	case 34816:
		return "GL_STENCIL_BACK_FUNC"
	case 34818:
		return "GL_STENCIL_BACK_PASS_DEPTH_FAIL"
	case 34819:
		return "GL_STENCIL_BACK_PASS_DEPTH_PASS"
	case 36003:
		return "GL_STENCIL_BACK_REF"
	case 36004:
		return "GL_STENCIL_BACK_VALUE_MASK"
	case 36005:
		return "GL_STENCIL_BACK_WRITEMASK"
	case 3415:
		return "GL_STENCIL_BITS"
	case 2961:
		return "GL_STENCIL_CLEAR_VALUE"
	case 2964:
		return "GL_STENCIL_FAIL"
	case 2962:
		return "GL_STENCIL_FUNC"
	case 2965:
		return "GL_STENCIL_PASS_DEPTH_FAIL"
	case 2966:
		return "GL_STENCIL_PASS_DEPTH_PASS"
	case 2967:
		return "GL_STENCIL_REF"
	case 2960:
		return "GL_STENCIL_TEST"
	case 2963:
		return "GL_STENCIL_VALUE_MASK"
	case 2968:
		return "GL_STENCIL_WRITEMASK"
	case 3408:
		return "GL_SUBPIXEL_BITS"
	case 32873:
		return "GL_TEXTURE_BINDING_2D"
	case 34068:
		return "GL_TEXTURE_BINDING_CUBE_MAP"
	case 3317:
		return "GL_UNPACK_ALIGNMENT"
	case 2978:
		return "GL_VIEWPORT"
	default:
		return fmt.Sprintf("StateVariable_GLES_2_0<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable_GLES_3_1
////////////////////////////////////////////////////////////////////////////////
type StateVariable_GLES_3_1 uint32

const (
	StateVariable_GLES_3_1_GL_READ_FRAMEBUFFER_BINDING = StateVariable_GLES_3_1(36010)
)

func (v StateVariable_GLES_3_1) String() string {
	switch v {
	case 36010:
		return "GL_READ_FRAMEBUFFER_BINDING"
	default:
		return fmt.Sprintf("StateVariable_GLES_3_1<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable_EXT_texture_filter_anisotropic
////////////////////////////////////////////////////////////////////////////////
type StateVariable_EXT_texture_filter_anisotropic uint32

const (
	StateVariable_EXT_texture_filter_anisotropic_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = StateVariable_EXT_texture_filter_anisotropic(34047)
)

func (v StateVariable_EXT_texture_filter_anisotropic) String() string {
	switch v {
	case 34047:
		return "GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT"
	default:
		return fmt.Sprintf("StateVariable_EXT_texture_filter_anisotropic<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum StateVariable_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type StateVariable_EXT_disjoint_timer_query uint32

const (
	StateVariable_EXT_disjoint_timer_query_GL_GPU_DISJOINT_EXT = StateVariable_EXT_disjoint_timer_query(36795)
)

func (v StateVariable_EXT_disjoint_timer_query) String() string {
	switch v {
	case 36795:
		return "GL_GPU_DISJOINT_EXT"
	default:
		return fmt.Sprintf("StateVariable_EXT_disjoint_timer_query<0x%.4x>", uint32(v))
	}
}

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

func (v StateVariable) String() string {
	switch v {
	// StateVariable_GLES_2_0
	case 34016:
		return "GL_ACTIVE_TEXTURE"
	case 33902:
		return "GL_ALIASED_LINE_WIDTH_RANGE"
	case 33901:
		return "GL_ALIASED_POINT_SIZE_RANGE"
	case 3413:
		return "GL_ALPHA_BITS"
	case 34964:
		return "GL_ARRAY_BUFFER_BINDING"
	case 3042:
		return "GL_BLEND"
	case 32773:
		return "GL_BLEND_COLOR"
	case 32970:
		return "GL_BLEND_DST_ALPHA"
	case 32968:
		return "GL_BLEND_DST_RGB"
	case 34877:
		return "GL_BLEND_EQUATION_ALPHA"
	case 32777:
		return "GL_BLEND_EQUATION_RGB"
	case 32971:
		return "GL_BLEND_SRC_ALPHA"
	case 32969:
		return "GL_BLEND_SRC_RGB"
	case 3412:
		return "GL_BLUE_BITS"
	case 3106:
		return "GL_COLOR_CLEAR_VALUE"
	case 3107:
		return "GL_COLOR_WRITEMASK"
	case 34467:
		return "GL_COMPRESSED_TEXTURE_FORMATS"
	case 2884:
		return "GL_CULL_FACE"
	case 2885:
		return "GL_CULL_FACE_MODE"
	case 35725:
		return "GL_CURRENT_PROGRAM"
	case 3414:
		return "GL_DEPTH_BITS"
	case 2931:
		return "GL_DEPTH_CLEAR_VALUE"
	case 2932:
		return "GL_DEPTH_FUNC"
	case 2928:
		return "GL_DEPTH_RANGE"
	case 2929:
		return "GL_DEPTH_TEST"
	case 2930:
		return "GL_DEPTH_WRITEMASK"
	case 3024:
		return "GL_DITHER"
	case 34965:
		return "GL_ELEMENT_ARRAY_BUFFER_BINDING"
	case 36006:
		return "GL_FRAMEBUFFER_BINDING"
	case 2886:
		return "GL_FRONT_FACE"
	case 33170:
		return "GL_GENERATE_MIPMAP_HINT"
	case 3411:
		return "GL_GREEN_BITS"
	case 35739:
		return "GL_IMPLEMENTATION_COLOR_READ_FORMAT"
	case 35738:
		return "GL_IMPLEMENTATION_COLOR_READ_TYPE"
	case 2849:
		return "GL_LINE_WIDTH"
	case 35661:
		return "GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS"
	case 34076:
		return "GL_MAX_CUBE_MAP_TEXTURE_SIZE"
	case 36349:
		return "GL_MAX_FRAGMENT_UNIFORM_VECTORS"
	case 34024:
		return "GL_MAX_RENDERBUFFER_SIZE"
	case 34930:
		return "GL_MAX_TEXTURE_IMAGE_UNITS"
	case 3379:
		return "GL_MAX_TEXTURE_SIZE"
	case 36348:
		return "GL_MAX_VARYING_VECTORS"
	case 34921:
		return "GL_MAX_VERTEX_ATTRIBS"
	case 35660:
		return "GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS"
	case 36347:
		return "GL_MAX_VERTEX_UNIFORM_VECTORS"
	case 3386:
		return "GL_MAX_VIEWPORT_DIMS"
	case 34466:
		return "GL_NUM_COMPRESSED_TEXTURE_FORMATS"
	case 36345:
		return "GL_NUM_SHADER_BINARY_FORMATS"
	case 3333:
		return "GL_PACK_ALIGNMENT"
	case 32824:
		return "GL_POLYGON_OFFSET_FACTOR"
	case 32823:
		return "GL_POLYGON_OFFSET_FILL"
	case 10752:
		return "GL_POLYGON_OFFSET_UNITS"
	case 3410:
		return "GL_RED_BITS"
	case 36007:
		return "GL_RENDERBUFFER_BINDING"
	case 32926:
		return "GL_SAMPLE_ALPHA_TO_COVERAGE"
	case 32936:
		return "GL_SAMPLE_BUFFERS"
	case 32928:
		return "GL_SAMPLE_COVERAGE"
	case 32939:
		return "GL_SAMPLE_COVERAGE_INVERT"
	case 32938:
		return "GL_SAMPLE_COVERAGE_VALUE"
	case 32937:
		return "GL_SAMPLES"
	case 3088:
		return "GL_SCISSOR_BOX"
	case 3089:
		return "GL_SCISSOR_TEST"
	case 36344:
		return "GL_SHADER_BINARY_FORMATS"
	case 36346:
		return "GL_SHADER_COMPILER"
	case 34817:
		return "GL_STENCIL_BACK_FAIL"
	case 34816:
		return "GL_STENCIL_BACK_FUNC"
	case 34818:
		return "GL_STENCIL_BACK_PASS_DEPTH_FAIL"
	case 34819:
		return "GL_STENCIL_BACK_PASS_DEPTH_PASS"
	case 36003:
		return "GL_STENCIL_BACK_REF"
	case 36004:
		return "GL_STENCIL_BACK_VALUE_MASK"
	case 36005:
		return "GL_STENCIL_BACK_WRITEMASK"
	case 3415:
		return "GL_STENCIL_BITS"
	case 2961:
		return "GL_STENCIL_CLEAR_VALUE"
	case 2964:
		return "GL_STENCIL_FAIL"
	case 2962:
		return "GL_STENCIL_FUNC"
	case 2965:
		return "GL_STENCIL_PASS_DEPTH_FAIL"
	case 2966:
		return "GL_STENCIL_PASS_DEPTH_PASS"
	case 2967:
		return "GL_STENCIL_REF"
	case 2960:
		return "GL_STENCIL_TEST"
	case 2963:
		return "GL_STENCIL_VALUE_MASK"
	case 2968:
		return "GL_STENCIL_WRITEMASK"
	case 3408:
		return "GL_SUBPIXEL_BITS"
	case 32873:
		return "GL_TEXTURE_BINDING_2D"
	case 34068:
		return "GL_TEXTURE_BINDING_CUBE_MAP"
	case 3317:
		return "GL_UNPACK_ALIGNMENT"
	case 2978:
		return "GL_VIEWPORT"
	// StateVariable_GLES_3_1
	case 36010:
		return "GL_READ_FRAMEBUFFER_BINDING"
	// StateVariable_EXT_texture_filter_anisotropic
	case 34047:
		return "GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT"
	// StateVariable_EXT_disjoint_timer_query
	case 36795:
		return "GL_GPU_DISJOINT_EXT"
	default:
		return fmt.Sprintf("StateVariable<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum FaceMode
////////////////////////////////////////////////////////////////////////////////
type FaceMode uint32

const (
	FaceMode_GL_FRONT          = FaceMode(1028)
	FaceMode_GL_BACK           = FaceMode(1029)
	FaceMode_GL_FRONT_AND_BACK = FaceMode(1032)
)

func (v FaceMode) String() string {
	switch v {
	case 1028:
		return "GL_FRONT"
	case 1029:
		return "GL_BACK"
	case 1032:
		return "GL_FRONT_AND_BACK"
	default:
		return fmt.Sprintf("FaceMode<0x%.4x>", uint32(v))
	}
}

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

func (v ArrayType_GLES_1_1) String() string {
	switch v {
	case 32884:
		return "GL_VERTEX_ARRAY"
	case 32885:
		return "GL_NORMAL_ARRAY"
	case 32886:
		return "GL_COLOR_ARRAY"
	case 32888:
		return "GL_TEXTURE_COORD_ARRAY"
	default:
		return fmt.Sprintf("ArrayType_GLES_1_1<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum ArrayType_OES_point_size_array
////////////////////////////////////////////////////////////////////////////////
type ArrayType_OES_point_size_array uint32

const (
	ArrayType_OES_point_size_array_GL_POINT_SIZE_ARRAY_OES = ArrayType_OES_point_size_array(35740)
)

func (v ArrayType_OES_point_size_array) String() string {
	switch v {
	case 35740:
		return "GL_POINT_SIZE_ARRAY_OES"
	default:
		return fmt.Sprintf("ArrayType_OES_point_size_array<0x%.4x>", uint32(v))
	}
}

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

func (v ArrayType) String() string {
	switch v {
	// ArrayType_GLES_1_1
	case 32884:
		return "GL_VERTEX_ARRAY"
	case 32885:
		return "GL_NORMAL_ARRAY"
	case 32886:
		return "GL_COLOR_ARRAY"
	case 32888:
		return "GL_TEXTURE_COORD_ARRAY"
	// ArrayType_OES_point_size_array
	case 35740:
		return "GL_POINT_SIZE_ARRAY_OES"
	default:
		return fmt.Sprintf("ArrayType<0x%.4x>", uint32(v))
	}
}

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

func (v Capability) String() string {
	switch v {
	case 3042:
		return "GL_BLEND"
	case 2884:
		return "GL_CULL_FACE"
	case 2929:
		return "GL_DEPTH_TEST"
	case 3024:
		return "GL_DITHER"
	case 32823:
		return "GL_POLYGON_OFFSET_FILL"
	case 32926:
		return "GL_SAMPLE_ALPHA_TO_COVERAGE"
	case 32928:
		return "GL_SAMPLE_COVERAGE"
	case 3089:
		return "GL_SCISSOR_TEST"
	case 2960:
		return "GL_STENCIL_TEST"
	// ArrayType
	// ArrayType_GLES_1_1
	case 32884:
		return "GL_VERTEX_ARRAY"
	case 32885:
		return "GL_NORMAL_ARRAY"
	case 32886:
		return "GL_COLOR_ARRAY"
	case 32888:
		return "GL_TEXTURE_COORD_ARRAY"
	// ArrayType_OES_point_size_array
	case 35740:
		return "GL_POINT_SIZE_ARRAY_OES"
	default:
		return fmt.Sprintf("Capability<0x%.4x>", uint32(v))
	}
}

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

func (v StringConstant) String() string {
	switch v {
	case 7939:
		return "GL_EXTENSIONS"
	case 7937:
		return "GL_RENDERER"
	case 7936:
		return "GL_VENDOR"
	case 7938:
		return "GL_VERSION"
	default:
		return fmt.Sprintf("StringConstant<0x%.4x>", uint32(v))
	}
}

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

func (v VertexAttribType) String() string {
	switch v {
	case 5120:
		return "GL_BYTE"
	case 5132:
		return "GL_FIXED"
	case 5126:
		return "GL_FLOAT"
	case 5122:
		return "GL_SHORT"
	case 5121:
		return "GL_UNSIGNED_BYTE"
	case 5123:
		return "GL_UNSIGNED_SHORT"
	// Type_OES_vertex_half_float
	case 36193:
		return "GL_HALF_FLOAT_OES"
	// Type_ARB_half_float_vertex
	case 5131:
		return "GL_ARB_half_float_vertex"
	default:
		return fmt.Sprintf("VertexAttribType<0x%.4x>", uint32(v))
	}
}

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

func (v ShaderAttribType) String() string {
	switch v {
	case 5126:
		return "GL_FLOAT"
	case 35664:
		return "GL_FLOAT_VEC2"
	case 35665:
		return "GL_FLOAT_VEC3"
	case 35666:
		return "GL_FLOAT_VEC4"
	case 35674:
		return "GL_FLOAT_MAT2"
	case 35675:
		return "GL_FLOAT_MAT3"
	case 35676:
		return "GL_FLOAT_MAT4"
	default:
		return fmt.Sprintf("ShaderAttribType<0x%.4x>", uint32(v))
	}
}

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

func (v ShaderUniformType) String() string {
	switch v {
	case 5126:
		return "GL_FLOAT"
	case 35664:
		return "GL_FLOAT_VEC2"
	case 35665:
		return "GL_FLOAT_VEC3"
	case 35666:
		return "GL_FLOAT_VEC4"
	case 5124:
		return "GL_INT"
	case 35667:
		return "GL_INT_VEC2"
	case 35668:
		return "GL_INT_VEC3"
	case 35669:
		return "GL_INT_VEC4"
	case 35670:
		return "GL_BOOL"
	case 35671:
		return "GL_BOOL_VEC2"
	case 35672:
		return "GL_BOOL_VEC3"
	case 35673:
		return "GL_BOOL_VEC4"
	case 35674:
		return "GL_FLOAT_MAT2"
	case 35675:
		return "GL_FLOAT_MAT3"
	case 35676:
		return "GL_FLOAT_MAT4"
	case 35678:
		return "GL_SAMPLER_2D"
	case 35680:
		return "GL_SAMPLER_CUBE"
	default:
		return fmt.Sprintf("ShaderUniformType<0x%.4x>", uint32(v))
	}
}

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

func (v Error) String() string {
	switch v {
	case 0:
		return "GL_NO_ERROR"
	case 1280:
		return "GL_INVALID_ENUM"
	case 1281:
		return "GL_INVALID_VALUE"
	case 1282:
		return "GL_INVALID_OPERATION"
	case 1286:
		return "GL_INVALID_FRAMEBUFFER_OPERATION"
	case 1285:
		return "GL_OUT_OF_MEMORY"
	default:
		return fmt.Sprintf("Error<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum HintTarget
////////////////////////////////////////////////////////////////////////////////
type HintTarget uint32

const (
	HintTarget_GL_GENERATE_MIPMAP_HINT = HintTarget(33170)
)

func (v HintTarget) String() string {
	switch v {
	case 33170:
		return "GL_GENERATE_MIPMAP_HINT"
	default:
		return fmt.Sprintf("HintTarget<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum HintMode
////////////////////////////////////////////////////////////////////////////////
type HintMode uint32

const (
	HintMode_GL_DONT_CARE = HintMode(4352)
	HintMode_GL_FASTEST   = HintMode(4353)
	HintMode_GL_NICEST    = HintMode(4354)
)

func (v HintMode) String() string {
	switch v {
	case 4352:
		return "GL_DONT_CARE"
	case 4353:
		return "GL_FASTEST"
	case 4354:
		return "GL_NICEST"
	default:
		return fmt.Sprintf("HintMode<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum DiscardFramebufferAttachment
////////////////////////////////////////////////////////////////////////////////
type DiscardFramebufferAttachment uint32

const (
	DiscardFramebufferAttachment_GL_COLOR_EXT   = DiscardFramebufferAttachment(6144)
	DiscardFramebufferAttachment_GL_DEPTH_EXT   = DiscardFramebufferAttachment(6145)
	DiscardFramebufferAttachment_GL_STENCIL_EXT = DiscardFramebufferAttachment(6146)
)

func (v DiscardFramebufferAttachment) String() string {
	switch v {
	case 6144:
		return "GL_COLOR_EXT"
	case 6145:
		return "GL_DEPTH_EXT"
	case 6146:
		return "GL_STENCIL_EXT"
	default:
		return fmt.Sprintf("DiscardFramebufferAttachment<0x%.4x>", uint32(v))
	}
}

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

func (v ProgramParameter) String() string {
	switch v {
	case 35712:
		return "GL_DELETE_STATUS"
	case 35714:
		return "GL_LINK_STATUS"
	case 35715:
		return "GL_VALIDATE_STATUS"
	case 35716:
		return "GL_INFO_LOG_LENGTH"
	case 35717:
		return "GL_ATTACHED_SHADERS"
	case 35721:
		return "GL_ACTIVE_ATTRIBUTES"
	case 35722:
		return "GL_ACTIVE_ATTRIBUTE_MAX_LENGTH"
	case 35718:
		return "GL_ACTIVE_UNIFORMS"
	case 35719:
		return "GL_ACTIVE_UNIFORM_MAX_LENGTH"
	default:
		return fmt.Sprintf("ProgramParameter<0x%.4x>", uint32(v))
	}
}

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

func (v ShaderParameter) String() string {
	switch v {
	case 35663:
		return "GL_SHADER_TYPE"
	case 35712:
		return "GL_DELETE_STATUS"
	case 35713:
		return "GL_COMPILE_STATUS"
	case 35716:
		return "GL_INFO_LOG_LENGTH"
	case 35720:
		return "GL_SHADER_SOURCE_LENGTH"
	default:
		return fmt.Sprintf("ShaderParameter<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum PixelStoreParameter
////////////////////////////////////////////////////////////////////////////////
type PixelStoreParameter uint32

const (
	PixelStoreParameter_GL_PACK_ALIGNMENT   = PixelStoreParameter(3333)
	PixelStoreParameter_GL_UNPACK_ALIGNMENT = PixelStoreParameter(3317)
)

func (v PixelStoreParameter) String() string {
	switch v {
	case 3333:
		return "GL_PACK_ALIGNMENT"
	case 3317:
		return "GL_UNPACK_ALIGNMENT"
	default:
		return fmt.Sprintf("PixelStoreParameter<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter_FilterMode
////////////////////////////////////////////////////////////////////////////////
type TextureParameter_FilterMode uint32

const (
	TextureParameter_FilterMode_GL_TEXTURE_MIN_FILTER = TextureParameter_FilterMode(10241)
	TextureParameter_FilterMode_GL_TEXTURE_MAG_FILTER = TextureParameter_FilterMode(10240)
)

func (v TextureParameter_FilterMode) String() string {
	switch v {
	case 10241:
		return "GL_TEXTURE_MIN_FILTER"
	case 10240:
		return "GL_TEXTURE_MAG_FILTER"
	default:
		return fmt.Sprintf("TextureParameter_FilterMode<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter_WrapMode
////////////////////////////////////////////////////////////////////////////////
type TextureParameter_WrapMode uint32

const (
	TextureParameter_WrapMode_GL_TEXTURE_WRAP_S = TextureParameter_WrapMode(10242)
	TextureParameter_WrapMode_GL_TEXTURE_WRAP_T = TextureParameter_WrapMode(10243)
)

func (v TextureParameter_WrapMode) String() string {
	switch v {
	case 10242:
		return "GL_TEXTURE_WRAP_S"
	case 10243:
		return "GL_TEXTURE_WRAP_T"
	default:
		return fmt.Sprintf("TextureParameter_WrapMode<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureParameter_EXT_texture_filter_anisotropic
////////////////////////////////////////////////////////////////////////////////
type TextureParameter_EXT_texture_filter_anisotropic uint32

const (
	TextureParameter_EXT_texture_filter_anisotropic_GL_TEXTURE_MAX_ANISOTROPY_EXT = TextureParameter_EXT_texture_filter_anisotropic(34046)
)

func (v TextureParameter_EXT_texture_filter_anisotropic) String() string {
	switch v {
	case 34046:
		return "GL_TEXTURE_MAX_ANISOTROPY_EXT"
	default:
		return fmt.Sprintf("TextureParameter_EXT_texture_filter_anisotropic<0x%.4x>", uint32(v))
	}
}

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

func (v TextureParameter_SwizzleMode) String() string {
	switch v {
	case 36418:
		return "GL_TEXTURE_SWIZZLE_R"
	case 36419:
		return "GL_TEXTURE_SWIZZLE_G"
	case 36420:
		return "GL_TEXTURE_SWIZZLE_B"
	case 36421:
		return "GL_TEXTURE_SWIZZLE_A"
	default:
		return fmt.Sprintf("TextureParameter_SwizzleMode<0x%.4x>", uint32(v))
	}
}

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

func (v TextureParameter) String() string {
	switch v {
	// TextureParameter_FilterMode
	case 10241:
		return "GL_TEXTURE_MIN_FILTER"
	case 10240:
		return "GL_TEXTURE_MAG_FILTER"
	// TextureParameter_WrapMode
	case 10242:
		return "GL_TEXTURE_WRAP_S"
	case 10243:
		return "GL_TEXTURE_WRAP_T"
	// TextureParameter_SwizzleMode
	case 36418:
		return "GL_TEXTURE_SWIZZLE_R"
	case 36419:
		return "GL_TEXTURE_SWIZZLE_G"
	case 36420:
		return "GL_TEXTURE_SWIZZLE_B"
	case 36421:
		return "GL_TEXTURE_SWIZZLE_A"
	// TextureParameter_EXT_texture_filter_anisotropic
	case 34046:
		return "GL_TEXTURE_MAX_ANISOTROPY_EXT"
	default:
		return fmt.Sprintf("TextureParameter<0x%.4x>", uint32(v))
	}
}

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

func (v TextureFilterMode) String() string {
	switch v {
	case 9728:
		return "GL_NEAREST"
	case 9729:
		return "GL_LINEAR"
	case 9984:
		return "GL_NEAREST_MIPMAP_NEAREST"
	case 9985:
		return "GL_LINEAR_MIPMAP_NEAREST"
	case 9986:
		return "GL_NEAREST_MIPMAP_LINEAR"
	case 9987:
		return "GL_LINEAR_MIPMAP_LINEAR"
	default:
		return fmt.Sprintf("TextureFilterMode<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureWrapMode
////////////////////////////////////////////////////////////////////////////////
type TextureWrapMode uint32

const (
	TextureWrapMode_GL_CLAMP_TO_EDGE   = TextureWrapMode(33071)
	TextureWrapMode_GL_MIRRORED_REPEAT = TextureWrapMode(33648)
	TextureWrapMode_GL_REPEAT          = TextureWrapMode(10497)
)

func (v TextureWrapMode) String() string {
	switch v {
	case 33071:
		return "GL_CLAMP_TO_EDGE"
	case 33648:
		return "GL_MIRRORED_REPEAT"
	case 10497:
		return "GL_REPEAT"
	default:
		return fmt.Sprintf("TextureWrapMode<0x%.4x>", uint32(v))
	}
}

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

func (v TexelComponent) String() string {
	switch v {
	case 6403:
		return "GL_RED"
	case 6404:
		return "GL_GREEN"
	case 6405:
		return "GL_BLUE"
	case 6406:
		return "GL_ALPHA"
	default:
		return fmt.Sprintf("TexelComponent<0x%.4x>", uint32(v))
	}
}

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

func (v BlendFactor) String() string {
	switch v {
	case 0:
		return "GL_ZERO"
	case 1:
		return "GL_ONE"
	case 768:
		return "GL_SRC_COLOR"
	case 769:
		return "GL_ONE_MINUS_SRC_COLOR"
	case 774:
		return "GL_DST_COLOR"
	case 775:
		return "GL_ONE_MINUS_DST_COLOR"
	case 770:
		return "GL_SRC_ALPHA"
	case 771:
		return "GL_ONE_MINUS_SRC_ALPHA"
	case 772:
		return "GL_DST_ALPHA"
	case 773:
		return "GL_ONE_MINUS_DST_ALPHA"
	case 32769:
		return "GL_CONSTANT_COLOR"
	case 32770:
		return "GL_ONE_MINUS_CONSTANT_COLOR"
	case 32771:
		return "GL_CONSTANT_ALPHA"
	case 32772:
		return "GL_ONE_MINUS_CONSTANT_ALPHA"
	case 776:
		return "GL_SRC_ALPHA_SATURATE"
	default:
		return fmt.Sprintf("BlendFactor<0x%.4x>", uint32(v))
	}
}

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

func (v PrecisionType) String() string {
	switch v {
	case 36336:
		return "GL_LOW_FLOAT"
	case 36337:
		return "GL_MEDIUM_FLOAT"
	case 36338:
		return "GL_HIGH_FLOAT"
	case 36339:
		return "GL_LOW_INT"
	case 36340:
		return "GL_MEDIUM_INT"
	case 36341:
		return "GL_HIGH_INT"
	default:
		return fmt.Sprintf("PrecisionType<0x%.4x>", uint32(v))
	}
}

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

func (v TestFunction) String() string {
	switch v {
	case 512:
		return "GL_NEVER"
	case 513:
		return "GL_LESS"
	case 514:
		return "GL_EQUAL"
	case 515:
		return "GL_LEQUAL"
	case 516:
		return "GL_GREATER"
	case 517:
		return "GL_NOTEQUAL"
	case 518:
		return "GL_GEQUAL"
	case 519:
		return "GL_ALWAYS"
	default:
		return fmt.Sprintf("TestFunction<0x%.4x>", uint32(v))
	}
}

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

func (v StencilAction) String() string {
	switch v {
	case 7680:
		return "GL_KEEP"
	case 0:
		return "GL_ZERO"
	case 7681:
		return "GL_REPLACE"
	case 7682:
		return "GL_INCR"
	case 34055:
		return "GL_INCR_WRAP"
	case 7683:
		return "GL_DECR"
	case 34056:
		return "GL_DECR_WRAP"
	case 5386:
		return "GL_INVERT"
	default:
		return fmt.Sprintf("StencilAction<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum FaceOrientation
////////////////////////////////////////////////////////////////////////////////
type FaceOrientation uint32

const (
	FaceOrientation_GL_CW  = FaceOrientation(2304)
	FaceOrientation_GL_CCW = FaceOrientation(2305)
)

func (v FaceOrientation) String() string {
	switch v {
	case 2304:
		return "GL_CW"
	case 2305:
		return "GL_CCW"
	default:
		return fmt.Sprintf("FaceOrientation<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum BlendEquation
////////////////////////////////////////////////////////////////////////////////
type BlendEquation uint32

const (
	BlendEquation_GL_FUNC_ADD              = BlendEquation(32774)
	BlendEquation_GL_FUNC_SUBTRACT         = BlendEquation(32778)
	BlendEquation_GL_FUNC_REVERSE_SUBTRACT = BlendEquation(32779)
)

func (v BlendEquation) String() string {
	switch v {
	case 32774:
		return "GL_FUNC_ADD"
	case 32778:
		return "GL_FUNC_SUBTRACT"
	case 32779:
		return "GL_FUNC_REVERSE_SUBTRACT"
	default:
		return fmt.Sprintf("BlendEquation<0x%.4x>", uint32(v))
	}
}

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

func (v BufferTarget) String() string {
	switch v {
	case 34962:
		return "GL_ARRAY_BUFFER"
	case 36662:
		return "GL_COPY_READ_BUFFER"
	case 36663:
		return "GL_COPY_WRITE_BUFFER"
	case 34963:
		return "GL_ELEMENT_ARRAY_BUFFER"
	case 35051:
		return "GL_PIXEL_PACK_BUFFER"
	case 35052:
		return "GL_PIXEL_UNPACK_BUFFER"
	case 35982:
		return "GL_TRANSFORM_FEEDBACK_BUFFER"
	case 35345:
		return "GL_UNIFORM_BUFFER"
	default:
		return fmt.Sprintf("BufferTarget<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum ImageTargetTexture_OES_EGL_image
////////////////////////////////////////////////////////////////////////////////
type ImageTargetTexture_OES_EGL_image uint32

const (
	ImageTargetTexture_OES_EGL_image_GL_TEXTURE_2D = ImageTargetTexture_OES_EGL_image(3553)
)

func (v ImageTargetTexture_OES_EGL_image) String() string {
	switch v {
	case 3553:
		return "GL_TEXTURE_2D"
	default:
		return fmt.Sprintf("ImageTargetTexture_OES_EGL_image<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum ImageTargetTexture_OES_EGL_image_external
////////////////////////////////////////////////////////////////////////////////
type ImageTargetTexture_OES_EGL_image_external uint32

const (
	ImageTargetTexture_OES_EGL_image_external_GL_TEXTURE_EXTERNAL_OES = ImageTargetTexture_OES_EGL_image_external(36197)
)

func (v ImageTargetTexture_OES_EGL_image_external) String() string {
	switch v {
	case 36197:
		return "GL_TEXTURE_EXTERNAL_OES"
	default:
		return fmt.Sprintf("ImageTargetTexture_OES_EGL_image_external<0x%.4x>", uint32(v))
	}
}

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

func (v ImageTargetTexture) String() string {
	switch v {
	// ImageTargetTexture_OES_EGL_image
	case 3553:
		return "GL_TEXTURE_2D"
	// ImageTargetTexture_OES_EGL_image_external
	case 36197:
		return "GL_TEXTURE_EXTERNAL_OES"
	default:
		return fmt.Sprintf("ImageTargetTexture<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum ImageTargetRenderbufferStorage
////////////////////////////////////////////////////////////////////////////////
type ImageTargetRenderbufferStorage uint32

const (
	ImageTargetRenderbufferStorage_GL_RENDERBUFFER_OES = ImageTargetRenderbufferStorage(36161)
)

func (v ImageTargetRenderbufferStorage) String() string {
	switch v {
	case 36161:
		return "GL_RENDERBUFFER_OES"
	default:
		return fmt.Sprintf("ImageTargetRenderbufferStorage<0x%.4x>", uint32(v))
	}
}

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

func (v ResetStatus) String() string {
	switch v {
	case 0:
		return "GL_NO_ERROR"
	case 33363:
		return "GL_GUILTY_CONTEXT_RESET_EXT"
	case 33364:
		return "GL_INNOCENT_CONTEXT_RESET_EXT"
	case 33365:
		return "GL_UNKNOWN_CONTEXT_RESET_EXT"
	default:
		return fmt.Sprintf("ResetStatus<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum TextureKind
////////////////////////////////////////////////////////////////////////////////
type TextureKind uint32

const (
	TextureKind_UNDEFINED = TextureKind(0)
	TextureKind_TEXTURE2D = TextureKind(1)
	TextureKind_CUBEMAP   = TextureKind(2)
)

func (v TextureKind) String() string {
	switch v {
	case 0:
		return "UNDEFINED"
	case 1:
		return "TEXTURE2D"
	case 2:
		return "CUBEMAP"
	default:
		return fmt.Sprintf("TextureKind<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum QueryParameter_GLES_3
////////////////////////////////////////////////////////////////////////////////
type QueryParameter_GLES_3 uint32

const (
	QueryParameter_GLES_3_GL_CURRENT_QUERY = QueryParameter_GLES_3(34917)
)

func (v QueryParameter_GLES_3) String() string {
	switch v {
	case 34917:
		return "GL_CURRENT_QUERY"
	default:
		return fmt.Sprintf("QueryParameter_GLES_3<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum QueryParameter_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type QueryParameter_EXT_disjoint_timer_query uint32

const (
	QueryParameter_EXT_disjoint_timer_query_GL_QUERY_COUNTER_BITS_EXT = QueryParameter_EXT_disjoint_timer_query(34916)
)

func (v QueryParameter_EXT_disjoint_timer_query) String() string {
	switch v {
	case 34916:
		return "GL_QUERY_COUNTER_BITS_EXT"
	default:
		return fmt.Sprintf("QueryParameter_EXT_disjoint_timer_query<0x%.4x>", uint32(v))
	}
}

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

func (v QueryParameter) String() string {
	switch v {
	// QueryParameter_GLES_3
	case 34917:
		return "GL_CURRENT_QUERY"
	// QueryParameter_EXT_disjoint_timer_query
	case 34916:
		return "GL_QUERY_COUNTER_BITS_EXT"
	default:
		return fmt.Sprintf("QueryParameter<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum QueryObjectParameter_GLES_3
////////////////////////////////////////////////////////////////////////////////
type QueryObjectParameter_GLES_3 uint32

const (
	QueryObjectParameter_GLES_3_GL_QUERY_RESULT           = QueryObjectParameter_GLES_3(34918)
	QueryObjectParameter_GLES_3_GL_QUERY_RESULT_AVAILABLE = QueryObjectParameter_GLES_3(34919)
)

func (v QueryObjectParameter_GLES_3) String() string {
	switch v {
	case 34918:
		return "GL_QUERY_RESULT"
	case 34919:
		return "GL_QUERY_RESULT_AVAILABLE"
	default:
		return fmt.Sprintf("QueryObjectParameter_GLES_3<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum QueryObjectParameter_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type QueryObjectParameter_EXT_disjoint_timer_query uint32

const ()

func (v QueryObjectParameter_EXT_disjoint_timer_query) String() string {
	switch v {
	default:
		return fmt.Sprintf("QueryObjectParameter_EXT_disjoint_timer_query<0x%.4x>", uint32(v))
	}
}

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

func (v QueryObjectParameter) String() string {
	switch v {
	// QueryObjectParameter_GLES_3
	case 34918:
		return "GL_QUERY_RESULT"
	case 34919:
		return "GL_QUERY_RESULT_AVAILABLE"
	// QueryObjectParameter_EXT_disjoint_timer_query
	default:
		return fmt.Sprintf("QueryObjectParameter<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum QueryTarget_GLES_3
////////////////////////////////////////////////////////////////////////////////
type QueryTarget_GLES_3 uint32

const (
	QueryTarget_GLES_3_GL_ANY_SAMPLES_PASSED                    = QueryTarget_GLES_3(35887)
	QueryTarget_GLES_3_GL_ANY_SAMPLES_PASSED_CONSERVATIVE       = QueryTarget_GLES_3(36202)
	QueryTarget_GLES_3_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = QueryTarget_GLES_3(35976)
)

func (v QueryTarget_GLES_3) String() string {
	switch v {
	case 35887:
		return "GL_ANY_SAMPLES_PASSED"
	case 36202:
		return "GL_ANY_SAMPLES_PASSED_CONSERVATIVE"
	case 35976:
		return "GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN"
	default:
		return fmt.Sprintf("QueryTarget_GLES_3<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum QueryTarget_EXT_disjoint_timer_query
////////////////////////////////////////////////////////////////////////////////
type QueryTarget_EXT_disjoint_timer_query uint32

const (
	QueryTarget_EXT_disjoint_timer_query_GL_TIME_ELAPSED_EXT = QueryTarget_EXT_disjoint_timer_query(35007)
	QueryTarget_EXT_disjoint_timer_query_GL_TIMESTAMP_EXT    = QueryTarget_EXT_disjoint_timer_query(36392)
)

func (v QueryTarget_EXT_disjoint_timer_query) String() string {
	switch v {
	case 35007:
		return "GL_TIME_ELAPSED_EXT"
	case 36392:
		return "GL_TIMESTAMP_EXT"
	default:
		return fmt.Sprintf("QueryTarget_EXT_disjoint_timer_query<0x%.4x>", uint32(v))
	}
}

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

func (v QueryTarget) String() string {
	switch v {
	// QueryTarget_GLES_3
	case 35887:
		return "GL_ANY_SAMPLES_PASSED"
	case 36202:
		return "GL_ANY_SAMPLES_PASSED_CONSERVATIVE"
	case 35976:
		return "GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN"
	// QueryTarget_EXT_disjoint_timer_query
	case 35007:
		return "GL_TIME_ELAPSED_EXT"
	case 36392:
		return "GL_TIMESTAMP_EXT"
	default:
		return fmt.Sprintf("QueryTarget<0x%.4x>", uint32(v))
	}
}

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

func (v TilePreserveMaskQCOM) String() string {
	switch v {
	case 1:
		return "GL_COLOR_BUFFER_BIT0_QCOM"
	case 2:
		return "GL_COLOR_BUFFER_BIT1_QCOM"
	case 4:
		return "GL_COLOR_BUFFER_BIT2_QCOM"
	case 8:
		return "GL_COLOR_BUFFER_BIT3_QCOM"
	case 16:
		return "GL_COLOR_BUFFER_BIT4_QCOM"
	case 32:
		return "GL_COLOR_BUFFER_BIT5_QCOM"
	case 64:
		return "GL_COLOR_BUFFER_BIT6_QCOM"
	case 128:
		return "GL_COLOR_BUFFER_BIT7_QCOM"
	case 256:
		return "GL_DEPTH_BUFFER_BIT0_QCOM"
	case 512:
		return "GL_DEPTH_BUFFER_BIT1_QCOM"
	case 1024:
		return "GL_DEPTH_BUFFER_BIT2_QCOM"
	case 2048:
		return "GL_DEPTH_BUFFER_BIT3_QCOM"
	case 4096:
		return "GL_DEPTH_BUFFER_BIT4_QCOM"
	case 8192:
		return "GL_DEPTH_BUFFER_BIT5_QCOM"
	case 16384:
		return "GL_DEPTH_BUFFER_BIT6_QCOM"
	case 32768:
		return "GL_DEPTH_BUFFER_BIT7_QCOM"
	case 65536:
		return "GL_STENCIL_BUFFER_BIT0_QCOM"
	case 131072:
		return "GL_STENCIL_BUFFER_BIT1_QCOM"
	case 262144:
		return "GL_STENCIL_BUFFER_BIT2_QCOM"
	case 524288:
		return "GL_STENCIL_BUFFER_BIT3_QCOM"
	case 1048576:
		return "GL_STENCIL_BUFFER_BIT4_QCOM"
	case 2097152:
		return "GL_STENCIL_BUFFER_BIT5_QCOM"
	case 4194304:
		return "GL_STENCIL_BUFFER_BIT6_QCOM"
	case 8388608:
		return "GL_STENCIL_BUFFER_BIT7_QCOM"
	case 16777216:
		return "GL_MULTISAMPLE_BUFFER_BIT0_QCOM"
	case 33554432:
		return "GL_MULTISAMPLE_BUFFER_BIT1_QCOM"
	case 67108864:
		return "GL_MULTISAMPLE_BUFFER_BIT2_QCOM"
	case 134217728:
		return "GL_MULTISAMPLE_BUFFER_BIT3_QCOM"
	case 268435456:
		return "GL_MULTISAMPLE_BUFFER_BIT4_QCOM"
	case 536870912:
		return "GL_MULTISAMPLE_BUFFER_BIT5_QCOM"
	case 1073741824:
		return "GL_MULTISAMPLE_BUFFER_BIT6_QCOM"
	case 2147483648:
		return "GL_MULTISAMPLE_BUFFER_BIT7_QCOM"
	default:
		return fmt.Sprintf("TilePreserveMaskQCOM<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// enum ClearMask
////////////////////////////////////////////////////////////////////////////////
type ClearMask uint32

const (
	ClearMask_GL_COLOR_BUFFER_BIT   = ClearMask(16384)
	ClearMask_GL_DEPTH_BUFFER_BIT   = ClearMask(256)
	ClearMask_GL_STENCIL_BUFFER_BIT = ClearMask(1024)
)

func (v ClearMask) String() string {
	switch v {
	case 16384:
		return "GL_COLOR_BUFFER_BIT"
	case 256:
		return "GL_DEPTH_BUFFER_BIT"
	case 1024:
		return "GL_STENCIL_BUFFER_BIT"
	default:
		return fmt.Sprintf("ClearMask<0x%.4x>", uint32(v))
	}
}

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

func (v MapBufferRangeAccess) String() string {
	switch v {
	case 1:
		return "GL_MAP_READ_BIT"
	case 2:
		return "GL_MAP_WRITE_BIT"
	case 4:
		return "GL_MAP_INVALIDATE_RANGE_BIT"
	case 8:
		return "GL_MAP_INVALIDATE_BUFFER_BIT"
	case 16:
		return "GL_MAP_FLUSH_EXPLICIT_BIT"
	case 32:
		return "GL_MAP_UNSYNCHRONIZED_BIT"
	default:
		return fmt.Sprintf("MapBufferRangeAccess<0x%.4x>", uint32(v))
	}
}

////////////////////////////////////////////////////////////////////////////////
// Globals
////////////////////////////////////////////////////////////////////////////////
type Globals struct {
	binary.Generate
	NextContextID ContextID
	CurrentThread ThreadID
	Contexts      ContextPtr_ThreadIDMap
	EGLContexts   ContextPtr_EGLContextMap
	GLXContexts   ContextPtr_GLXContextMap
	WGLContexts   ContextPtr_HGLRCMap
	CGLContexts   ContextPtr_CGLContextObjMap
}

func (g *Globals) Init() {
	g.Contexts = make(ContextPtr_ThreadIDMap)
	g.EGLContexts = make(ContextPtr_EGLContextMap)
	g.GLXContexts = make(ContextPtr_GLXContextMap)
	g.WGLContexts = make(ContextPtr_HGLRCMap)
	g.CGLContexts = make(ContextPtr_CGLContextObjMap)
}
func NewReplayCreateRenderer(
	pId uint32,
) *ReplayCreateRenderer {
	return &ReplayCreateRenderer{
		Id: pId}
}
func NewReplayBindRenderer(
	pId uint32,
) *ReplayBindRenderer {
	return &ReplayBindRenderer{
		Id: pId}
}
func NewBackbufferInfo(
	pWidth int32,
	pHeight int32,
	pColorFmt RenderbufferFormat,
	pDepthFmt RenderbufferFormat,
	pStencilFmt RenderbufferFormat,
	pResetViewportScissor bool,
) *BackbufferInfo {
	return &BackbufferInfo{
		Width: pWidth, Height: pHeight, ColorFmt: pColorFmt, DepthFmt: pDepthFmt, StencilFmt: pStencilFmt, ResetViewportScissor: pResetViewportScissor}
}
func NewStartTimer(
	pIndex uint8,
) *StartTimer {
	return &StartTimer{
		Index: pIndex}
}
func NewStopTimer(
	pIndex uint8,
	pResult uint64,
) *StopTimer {
	return &StopTimer{
		Index: pIndex, Result: pResult}
}
func NewFlushPostBuffer() *FlushPostBuffer {
	return &FlushPostBuffer{}
}
func NewEglInitialize(
	pDpy EGLDisplay,
	pMajor EGLint,
	pMinor EGLint,
	pResult EGLBoolean,
) *EglInitialize {
	return &EglInitialize{
		Dpy: pDpy, Major: pMajor, Minor: pMinor, Result: pResult}
}
func NewEglCreateContext(
	pDisplay EGLDisplay,
	pConfig EGLConfig,
	pShareContext EGLContext,
	pAttribList EGLintArray,
	pResult EGLContext,
) *EglCreateContext {
	return &EglCreateContext{
		Display: pDisplay, Config: pConfig, ShareContext: pShareContext, AttribList: pAttribList, Result: pResult}
}
func NewEglMakeCurrent(
	pDisplay EGLDisplay,
	pDraw EGLSurface,
	pRead EGLSurface,
	pContext EGLContext,
	pResult EGLBoolean,
) *EglMakeCurrent {
	return &EglMakeCurrent{
		Display: pDisplay, Draw: pDraw, Read: pRead, Context: pContext, Result: pResult}
}
func NewEglSwapBuffers(
	pDisplay EGLDisplay,
	pSurface memory.Pointer,
	pResult EGLBoolean,
) *EglSwapBuffers {
	return &EglSwapBuffers{
		Display: pDisplay, Surface: pSurface, Result: pResult}
}
func NewEglQuerySurface(
	pDisplay EGLDisplay,
	pSurface EGLSurface,
	pAttribute EGLint,
	pValue EGLint,
	pResult EGLBoolean,
) *EglQuerySurface {
	return &EglQuerySurface{
		Display: pDisplay, Surface: pSurface, Attribute: pAttribute, Value: pValue, Result: pResult}
}
func NewGlXCreateContext(
	pDpy memory.Pointer,
	pVis memory.Pointer,
	pShareList GLXContext,
	pDirect bool,
	pResult GLXContext,
) *GlXCreateContext {
	return &GlXCreateContext{
		Dpy: pDpy, Vis: pVis, ShareList: pShareList, Direct: pDirect, Result: pResult}
}
func NewGlXCreateNewContext(
	pDisplay memory.Pointer,
	pFbconfig memory.Pointer,
	pType uint32,
	pShared GLXContext,
	pDirect bool,
	pResult GLXContext,
) *GlXCreateNewContext {
	return &GlXCreateNewContext{
		Display: pDisplay, Fbconfig: pFbconfig, Type: pType, Shared: pShared, Direct: pDirect, Result: pResult}
}
func NewGlXMakeContextCurrent(
	pDisplay memory.Pointer,
	pDraw GLXDrawable,
	pRead GLXDrawable,
	pCtx GLXContext,
) *GlXMakeContextCurrent {
	return &GlXMakeContextCurrent{
		Display: pDisplay, Draw: pDraw, Read: pRead, Ctx: pCtx}
}
func NewGlXSwapBuffers(
	pDisplay memory.Pointer,
	pDrawable GLXDrawable,
) *GlXSwapBuffers {
	return &GlXSwapBuffers{
		Display: pDisplay, Drawable: pDrawable}
}
func NewWglCreateContext(
	pHdc HDC,
	pResult HGLRC,
) *WglCreateContext {
	return &WglCreateContext{
		Hdc: pHdc, Result: pResult}
}
func NewWglCreateContextAttribsARB(
	pHdc HDC,
	pHShareContext HGLRC,
	pAttribList IntArray,
	pResult HGLRC,
) *WglCreateContextAttribsARB {
	return &WglCreateContextAttribsARB{
		Hdc: pHdc, HShareContext: pHShareContext, AttribList: pAttribList, Result: pResult}
}
func NewWglMakeCurrent(
	pHdc HDC,
	pHglrc HGLRC,
	pResult BOOL,
) *WglMakeCurrent {
	return &WglMakeCurrent{
		Hdc: pHdc, Hglrc: pHglrc, Result: pResult}
}
func NewWglSwapBuffers(
	pHdc HDC,
) *WglSwapBuffers {
	return &WglSwapBuffers{
		Hdc: pHdc}
}
func NewCGLCreateContext(
	pPix CGLPixelFormatObj,
	pShare CGLContextObj,
	pCtx CGLContextObj,
	pResult CGLError,
) *CGLCreateContext {
	return &CGLCreateContext{
		Pix: pPix, Share: pShare, Ctx: pCtx, Result: pResult}
}
func NewCGLSetCurrentContext(
	pCtx CGLContextObj,
	pResult CGLError,
) *CGLSetCurrentContext {
	return &CGLSetCurrentContext{
		Ctx: pCtx, Result: pResult}
}
func NewGlEnableClientState(
	pType ArrayType,
) *GlEnableClientState {
	return &GlEnableClientState{
		Type: pType}
}
func NewGlDisableClientState(
	pType ArrayType,
) *GlDisableClientState {
	return &GlDisableClientState{
		Type: pType}
}
func NewGlGetProgramBinaryOES(
	pProgram ProgramId,
	pBufferSize int32,
	pBytesWritten int32,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
) *GlGetProgramBinaryOES {
	return &GlGetProgramBinaryOES{
		Program: pProgram, BufferSize: pBufferSize, BytesWritten: pBytesWritten, BinaryFormat: pBinaryFormat, Binary: pBinary}
}
func NewGlProgramBinaryOES(
	pProgram ProgramId,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
	pBinarySize int32,
) *GlProgramBinaryOES {
	return &GlProgramBinaryOES{
		Program: pProgram, BinaryFormat: pBinaryFormat, Binary: pBinary, BinarySize: pBinarySize}
}
func NewGlStartTilingQCOM(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pPreserveMask TilePreserveMaskQCOM,
) *GlStartTilingQCOM {
	return &GlStartTilingQCOM{
		X: pX, Y: pY, Width: pWidth, Height: pHeight, PreserveMask: pPreserveMask}
}
func NewGlEndTilingQCOM(
	pPreserveMask TilePreserveMaskQCOM,
) *GlEndTilingQCOM {
	return &GlEndTilingQCOM{
		PreserveMask: pPreserveMask}
}
func NewGlDiscardFramebufferEXT(
	pTarget FramebufferTarget,
	pNumAttachments int32,
	pAttachments DiscardFramebufferAttachmentArray,
) *GlDiscardFramebufferEXT {
	return &GlDiscardFramebufferEXT{
		Target: pTarget, NumAttachments: pNumAttachments, Attachments: pAttachments}
}
func NewGlInsertEventMarkerEXT(
	pLength int32,
	pMarker string,
) *GlInsertEventMarkerEXT {
	return &GlInsertEventMarkerEXT{
		Length: pLength, Marker: pMarker}
}
func NewGlPushGroupMarkerEXT(
	pLength int32,
	pMarker string,
) *GlPushGroupMarkerEXT {
	return &GlPushGroupMarkerEXT{
		Length: pLength, Marker: pMarker}
}
func NewGlPopGroupMarkerEXT() *GlPopGroupMarkerEXT {
	return &GlPopGroupMarkerEXT{}
}
func NewGlTexStorage1DEXT(
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
) *GlTexStorage1DEXT {
	return &GlTexStorage1DEXT{
		Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth}
}
func NewGlTexStorage2DEXT(
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
) *GlTexStorage2DEXT {
	return &GlTexStorage2DEXT{
		Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight}
}
func NewGlTexStorage3DEXT(
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
	pDepth int32,
) *GlTexStorage3DEXT {
	return &GlTexStorage3DEXT{
		Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight, Depth: pDepth}
}
func NewGlTextureStorage1DEXT(
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
) *GlTextureStorage1DEXT {
	return &GlTextureStorage1DEXT{
		Texture: pTexture, Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth}
}
func NewGlTextureStorage2DEXT(
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
) *GlTextureStorage2DEXT {
	return &GlTextureStorage2DEXT{
		Texture: pTexture, Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight}
}
func NewGlTextureStorage3DEXT(
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
	pDepth int32,
) *GlTextureStorage3DEXT {
	return &GlTextureStorage3DEXT{
		Texture: pTexture, Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight, Depth: pDepth}
}
func NewGlGenVertexArraysOES(
	pCount int32,
	pArrays VertexArrayIdArray,
) *GlGenVertexArraysOES {
	return &GlGenVertexArraysOES{
		Count: pCount, Arrays: pArrays}
}
func NewGlBindVertexArrayOES(
	pArray VertexArrayId,
) *GlBindVertexArrayOES {
	return &GlBindVertexArrayOES{
		Array: pArray}
}
func NewGlDeleteVertexArraysOES(
	pCount int32,
	pArrays VertexArrayIdArray,
) *GlDeleteVertexArraysOES {
	return &GlDeleteVertexArraysOES{
		Count: pCount, Arrays: pArrays}
}
func NewGlIsVertexArrayOES(
	pArray VertexArrayId,
	pResult bool,
) *GlIsVertexArrayOES {
	return &GlIsVertexArrayOES{
		Array: pArray, Result: pResult}
}
func NewGlEGLImageTargetTexture2DOES(
	pTarget ImageTargetTexture,
	pImage ImageOES,
) *GlEGLImageTargetTexture2DOES {
	return &GlEGLImageTargetTexture2DOES{
		Target: pTarget, Image: pImage}
}
func NewGlEGLImageTargetRenderbufferStorageOES(
	pTarget ImageTargetRenderbufferStorage,
	pImage TexturePointer,
) *GlEGLImageTargetRenderbufferStorageOES {
	return &GlEGLImageTargetRenderbufferStorageOES{
		Target: pTarget, Image: pImage}
}
func NewGlGetGraphicsResetStatusEXT(
	pResult ResetStatus,
) *GlGetGraphicsResetStatusEXT {
	return &GlGetGraphicsResetStatusEXT{
		Result: pResult}
}
func NewGlBindAttribLocation(
	pProgram ProgramId,
	pLocation AttributeLocation,
	pName string,
) *GlBindAttribLocation {
	return &GlBindAttribLocation{
		Program: pProgram, Location: pLocation, Name: pName}
}
func NewGlBlendFunc(
	pSrcFactor BlendFactor,
	pDstFactor BlendFactor,
) *GlBlendFunc {
	return &GlBlendFunc{
		SrcFactor: pSrcFactor, DstFactor: pDstFactor}
}
func NewGlBlendFuncSeparate(
	pSrcFactorRgb BlendFactor,
	pDstFactorRgb BlendFactor,
	pSrcFactorAlpha BlendFactor,
	pDstFactorAlpha BlendFactor,
) *GlBlendFuncSeparate {
	return &GlBlendFuncSeparate{
		SrcFactorRgb: pSrcFactorRgb, DstFactorRgb: pDstFactorRgb, SrcFactorAlpha: pSrcFactorAlpha, DstFactorAlpha: pDstFactorAlpha}
}
func NewGlBlendEquation(
	pEquation BlendEquation,
) *GlBlendEquation {
	return &GlBlendEquation{
		Equation: pEquation}
}
func NewGlBlendEquationSeparate(
	pRgb BlendEquation,
	pAlpha BlendEquation,
) *GlBlendEquationSeparate {
	return &GlBlendEquationSeparate{
		Rgb: pRgb, Alpha: pAlpha}
}
func NewGlBlendColor(
	pRed float32,
	pGreen float32,
	pBlue float32,
	pAlpha float32,
) *GlBlendColor {
	return &GlBlendColor{
		Red: pRed, Green: pGreen, Blue: pBlue, Alpha: pAlpha}
}
func NewGlEnableVertexAttribArray(
	pLocation AttributeLocation,
) *GlEnableVertexAttribArray {
	return &GlEnableVertexAttribArray{
		Location: pLocation}
}
func NewGlDisableVertexAttribArray(
	pLocation AttributeLocation,
) *GlDisableVertexAttribArray {
	return &GlDisableVertexAttribArray{
		Location: pLocation}
}
func NewGlVertexAttribPointer(
	pLocation AttributeLocation,
	pSize int32,
	pType VertexAttribType,
	pNormalized bool,
	pStride int32,
	pData VertexPointer,
) *GlVertexAttribPointer {
	return &GlVertexAttribPointer{
		Location: pLocation, Size: pSize, Type: pType, Normalized: pNormalized, Stride: pStride, Data: pData}
}
func NewGlGetActiveAttrib(
	pProgram ProgramId,
	pLocation AttributeLocation,
	pBufferSize int32,
	pBufferBytesWritten int32,
	pVectorCount int32,
	pType ShaderAttribType,
	pName string,
) *GlGetActiveAttrib {
	return &GlGetActiveAttrib{
		Program: pProgram, Location: pLocation, BufferSize: pBufferSize, BufferBytesWritten: pBufferBytesWritten, VectorCount: pVectorCount, Type: pType, Name: pName}
}
func NewGlGetActiveUniform(
	pProgram ProgramId,
	pLocation int32,
	pBufferSize int32,
	pBufferBytesWritten int32,
	pSize int32,
	pType ShaderUniformType,
	pName string,
) *GlGetActiveUniform {
	return &GlGetActiveUniform{
		Program: pProgram, Location: pLocation, BufferSize: pBufferSize, BufferBytesWritten: pBufferBytesWritten, Size: pSize, Type: pType, Name: pName}
}
func NewGlGetError(
	pResult Error,
) *GlGetError {
	return &GlGetError{
		Result: pResult}
}
func NewGlGetProgramiv(
	pProgram ProgramId,
	pParameter ProgramParameter,
	pValue S32Array,
) *GlGetProgramiv {
	return &GlGetProgramiv{
		Program: pProgram, Parameter: pParameter, Value: pValue}
}
func NewGlGetShaderiv(
	pShader ShaderId,
	pParameter ShaderParameter,
	pValue S32Array,
) *GlGetShaderiv {
	return &GlGetShaderiv{
		Shader: pShader, Parameter: pParameter, Value: pValue}
}
func NewGlGetUniformLocation(
	pProgram ProgramId,
	pName string,
	pResult UniformLocation,
) *GlGetUniformLocation {
	return &GlGetUniformLocation{
		Program: pProgram, Name: pName, Result: pResult}
}
func NewGlGetAttribLocation(
	pProgram ProgramId,
	pName string,
	pResult AttributeLocation,
) *GlGetAttribLocation {
	return &GlGetAttribLocation{
		Program: pProgram, Name: pName, Result: pResult}
}
func NewGlPixelStorei(
	pParameter PixelStoreParameter,
	pValue int32,
) *GlPixelStorei {
	return &GlPixelStorei{
		Parameter: pParameter, Value: pValue}
}
func NewGlTexParameteri(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValue int32,
) *GlTexParameteri {
	return &GlTexParameteri{
		Target: pTarget, Parameter: pParameter, Value: pValue}
}
func NewGlTexParameterf(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValue float32,
) *GlTexParameterf {
	return &GlTexParameterf{
		Target: pTarget, Parameter: pParameter, Value: pValue}
}
func NewGlGetTexParameteriv(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValues S32Array,
) *GlGetTexParameteriv {
	return &GlGetTexParameteriv{
		Target: pTarget, Parameter: pParameter, Values: pValues}
}
func NewGlGetTexParameterfv(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValues F32Array,
) *GlGetTexParameterfv {
	return &GlGetTexParameterfv{
		Target: pTarget, Parameter: pParameter, Values: pValues}
}
func NewGlUniform1i(
	pLocation UniformLocation,
	pValue int32,
) *GlUniform1i {
	return &GlUniform1i{
		Location: pLocation, Value: pValue}
}
func NewGlUniform2i(
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
) *GlUniform2i {
	return &GlUniform2i{
		Location: pLocation, Value0: pValue0, Value1: pValue1}
}
func NewGlUniform3i(
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
	pValue2 int32,
) *GlUniform3i {
	return &GlUniform3i{
		Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2}
}
func NewGlUniform4i(
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
	pValue2 int32,
	pValue3 int32,
) *GlUniform4i {
	return &GlUniform4i{
		Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2, Value3: pValue3}
}
func NewGlUniform1iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform1iv {
	return &GlUniform1iv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniform2iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform2iv {
	return &GlUniform2iv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniform3iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform3iv {
	return &GlUniform3iv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniform4iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform4iv {
	return &GlUniform4iv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniform1f(
	pLocation UniformLocation,
	pValue float32,
) *GlUniform1f {
	return &GlUniform1f{
		Location: pLocation, Value: pValue}
}
func NewGlUniform2f(
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
) *GlUniform2f {
	return &GlUniform2f{
		Location: pLocation, Value0: pValue0, Value1: pValue1}
}
func NewGlUniform3f(
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
) *GlUniform3f {
	return &GlUniform3f{
		Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2}
}
func NewGlUniform4f(
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
	pValue3 float32,
) *GlUniform4f {
	return &GlUniform4f{
		Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2, Value3: pValue3}
}
func NewGlUniform1fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform1fv {
	return &GlUniform1fv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniform2fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform2fv {
	return &GlUniform2fv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniform3fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform3fv {
	return &GlUniform3fv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniform4fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform4fv {
	return &GlUniform4fv{
		Location: pLocation, Count: pCount, Value: pValue}
}
func NewGlUniformMatrix2fv(
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix2fv {
	return &GlUniformMatrix2fv{
		Location: pLocation, Count: pCount, Transpose: pTranspose, Values: pValues}
}
func NewGlUniformMatrix3fv(
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix3fv {
	return &GlUniformMatrix3fv{
		Location: pLocation, Count: pCount, Transpose: pTranspose, Values: pValues}
}
func NewGlUniformMatrix4fv(
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix4fv {
	return &GlUniformMatrix4fv{
		Location: pLocation, Count: pCount, Transpose: pTranspose, Values: pValues}
}
func NewGlGetUniformfv(
	pProgram ProgramId,
	pLocation UniformLocation,
	pValues F32Array,
) *GlGetUniformfv {
	return &GlGetUniformfv{
		Program: pProgram, Location: pLocation, Values: pValues}
}
func NewGlGetUniformiv(
	pProgram ProgramId,
	pLocation UniformLocation,
	pValues S32Array,
) *GlGetUniformiv {
	return &GlGetUniformiv{
		Program: pProgram, Location: pLocation, Values: pValues}
}
func NewGlVertexAttrib1f(
	pLocation AttributeLocation,
	pValue0 float32,
) *GlVertexAttrib1f {
	return &GlVertexAttrib1f{
		Location: pLocation, Value0: pValue0}
}
func NewGlVertexAttrib2f(
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
) *GlVertexAttrib2f {
	return &GlVertexAttrib2f{
		Location: pLocation, Value0: pValue0, Value1: pValue1}
}
func NewGlVertexAttrib3f(
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
) *GlVertexAttrib3f {
	return &GlVertexAttrib3f{
		Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2}
}
func NewGlVertexAttrib4f(
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
	pValue3 float32,
) *GlVertexAttrib4f {
	return &GlVertexAttrib4f{
		Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2, Value3: pValue3}
}
func NewGlVertexAttrib1fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib1fv {
	return &GlVertexAttrib1fv{
		Location: pLocation, Value: pValue}
}
func NewGlVertexAttrib2fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib2fv {
	return &GlVertexAttrib2fv{
		Location: pLocation, Value: pValue}
}
func NewGlVertexAttrib3fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib3fv {
	return &GlVertexAttrib3fv{
		Location: pLocation, Value: pValue}
}
func NewGlVertexAttrib4fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib4fv {
	return &GlVertexAttrib4fv{
		Location: pLocation, Value: pValue}
}
func NewGlGetShaderPrecisionFormat(
	pShaderType ShaderType,
	pPrecisionType PrecisionType,
	pRange S32Array,
	pPrecision int32,
) *GlGetShaderPrecisionFormat {
	return &GlGetShaderPrecisionFormat{
		ShaderType: pShaderType, PrecisionType: pPrecisionType, Range: pRange, Precision: pPrecision}
}
func NewGlDepthMask(
	pEnabled bool,
) *GlDepthMask {
	return &GlDepthMask{
		Enabled: pEnabled}
}
func NewGlDepthFunc(
	pFunction TestFunction,
) *GlDepthFunc {
	return &GlDepthFunc{
		Function: pFunction}
}
func NewGlDepthRangef(
	pNear float32,
	pFar float32,
) *GlDepthRangef {
	return &GlDepthRangef{
		Near: pNear, Far: pFar}
}
func NewGlColorMask(
	pRed bool,
	pGreen bool,
	pBlue bool,
	pAlpha bool,
) *GlColorMask {
	return &GlColorMask{
		Red: pRed, Green: pGreen, Blue: pBlue, Alpha: pAlpha}
}
func NewGlStencilMask(
	pMask uint32,
) *GlStencilMask {
	return &GlStencilMask{
		Mask: pMask}
}
func NewGlStencilMaskSeparate(
	pFace FaceMode,
	pMask uint32,
) *GlStencilMaskSeparate {
	return &GlStencilMaskSeparate{
		Face: pFace, Mask: pMask}
}
func NewGlStencilFuncSeparate(
	pFace FaceMode,
	pFunction TestFunction,
	pReferenceValue int32,
	pMask int32,
) *GlStencilFuncSeparate {
	return &GlStencilFuncSeparate{
		Face: pFace, Function: pFunction, ReferenceValue: pReferenceValue, Mask: pMask}
}
func NewGlStencilOpSeparate(
	pFace FaceMode,
	pStencilFail StencilAction,
	pStencilPassDepthFail StencilAction,
	pStencilPassDepthPass StencilAction,
) *GlStencilOpSeparate {
	return &GlStencilOpSeparate{
		Face: pFace, StencilFail: pStencilFail, StencilPassDepthFail: pStencilPassDepthFail, StencilPassDepthPass: pStencilPassDepthPass}
}
func NewGlFrontFace(
	pOrientation FaceOrientation,
) *GlFrontFace {
	return &GlFrontFace{
		Orientation: pOrientation}
}
func NewGlViewport(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlViewport {
	return &GlViewport{
		X: pX, Y: pY, Width: pWidth, Height: pHeight}
}
func NewGlScissor(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlScissor {
	return &GlScissor{
		X: pX, Y: pY, Width: pWidth, Height: pHeight}
}
func NewGlActiveTexture(
	pUnit TextureUnit,
) *GlActiveTexture {
	return &GlActiveTexture{
		Unit: pUnit}
}
func NewGlGenTextures(
	pCount int32,
	pTextures TextureIdArray,
) *GlGenTextures {
	return &GlGenTextures{
		Count: pCount, Textures: pTextures}
}
func NewGlDeleteTextures(
	pCount int32,
	pTextures TextureIdArray,
) *GlDeleteTextures {
	return &GlDeleteTextures{
		Count: pCount, Textures: pTextures}
}
func NewGlIsTexture(
	pTexture TextureId,
	pResult bool,
) *GlIsTexture {
	return &GlIsTexture{
		Texture: pTexture, Result: pResult}
}
func NewGlBindTexture(
	pTarget TextureTarget,
	pTexture TextureId,
) *GlBindTexture {
	return &GlBindTexture{
		Target: pTarget, Texture: pTexture}
}
func NewGlTexImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pInternalFormat TexelFormat,
	pWidth int32,
	pHeight int32,
	pBorder int32,
	pFormat TexelFormat,
	pType TexelType,
	pData TexturePointer,
) *GlTexImage2D {
	return &GlTexImage2D{
		Target: pTarget, Level: pLevel, InternalFormat: pInternalFormat, Width: pWidth, Height: pHeight, Border: pBorder, Format: pFormat, Type: pType, Data: pData}
}
func NewGlTexSubImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pXoffset int32,
	pYoffset int32,
	pWidth int32,
	pHeight int32,
	pFormat TexelFormat,
	pType TexelType,
	pData TexturePointer,
) *GlTexSubImage2D {
	return &GlTexSubImage2D{
		Target: pTarget, Level: pLevel, Xoffset: pXoffset, Yoffset: pYoffset, Width: pWidth, Height: pHeight, Format: pFormat, Type: pType, Data: pData}
}
func NewGlCopyTexImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pFormat TexelFormat,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pBorder int32,
) *GlCopyTexImage2D {
	return &GlCopyTexImage2D{
		Target: pTarget, Level: pLevel, Format: pFormat, X: pX, Y: pY, Width: pWidth, Height: pHeight, Border: pBorder}
}
func NewGlCopyTexSubImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pXoffset int32,
	pYoffset int32,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlCopyTexSubImage2D {
	return &GlCopyTexSubImage2D{
		Target: pTarget, Level: pLevel, Xoffset: pXoffset, Yoffset: pYoffset, X: pX, Y: pY, Width: pWidth, Height: pHeight}
}
func NewGlCompressedTexImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pFormat CompressedTexelFormat,
	pWidth int32,
	pHeight int32,
	pBorder int32,
	pImageSize int32,
	pData TexturePointer,
) *GlCompressedTexImage2D {
	return &GlCompressedTexImage2D{
		Target: pTarget, Level: pLevel, Format: pFormat, Width: pWidth, Height: pHeight, Border: pBorder, ImageSize: pImageSize, Data: pData}
}
func NewGlCompressedTexSubImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pXoffset int32,
	pYoffset int32,
	pWidth int32,
	pHeight int32,
	pFormat CompressedTexelFormat,
	pImageSize int32,
	pData TexturePointer,
) *GlCompressedTexSubImage2D {
	return &GlCompressedTexSubImage2D{
		Target: pTarget, Level: pLevel, Xoffset: pXoffset, Yoffset: pYoffset, Width: pWidth, Height: pHeight, Format: pFormat, ImageSize: pImageSize, Data: pData}
}
func NewGlGenerateMipmap(
	pTarget TextureImageTarget,
) *GlGenerateMipmap {
	return &GlGenerateMipmap{
		Target: pTarget}
}
func NewGlReadPixels(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pFormat BaseTexelFormat,
	pType TexelType,
	pData memory.Pointer,
) *GlReadPixels {
	return &GlReadPixels{
		X: pX, Y: pY, Width: pWidth, Height: pHeight, Format: pFormat, Type: pType, Data: pData}
}
func NewGlGenFramebuffers(
	pCount int32,
	pFramebuffers FramebufferIdArray,
) *GlGenFramebuffers {
	return &GlGenFramebuffers{
		Count: pCount, Framebuffers: pFramebuffers}
}
func NewGlBindFramebuffer(
	pTarget FramebufferTarget,
	pFramebuffer FramebufferId,
) *GlBindFramebuffer {
	return &GlBindFramebuffer{
		Target: pTarget, Framebuffer: pFramebuffer}
}
func NewGlCheckFramebufferStatus(
	pTarget FramebufferTarget,
	pResult FramebufferStatus,
) *GlCheckFramebufferStatus {
	return &GlCheckFramebufferStatus{
		Target: pTarget, Result: pResult}
}
func NewGlDeleteFramebuffers(
	pCount int32,
	pFramebuffers FramebufferIdArray,
) *GlDeleteFramebuffers {
	return &GlDeleteFramebuffers{
		Count: pCount, Framebuffers: pFramebuffers}
}
func NewGlIsFramebuffer(
	pFramebuffer FramebufferId,
	pResult bool,
) *GlIsFramebuffer {
	return &GlIsFramebuffer{
		Framebuffer: pFramebuffer, Result: pResult}
}
func NewGlGenRenderbuffers(
	pCount int32,
	pRenderbuffers RenderbufferIdArray,
) *GlGenRenderbuffers {
	return &GlGenRenderbuffers{
		Count: pCount, Renderbuffers: pRenderbuffers}
}
func NewGlBindRenderbuffer(
	pTarget RenderbufferTarget,
	pRenderbuffer RenderbufferId,
) *GlBindRenderbuffer {
	return &GlBindRenderbuffer{
		Target: pTarget, Renderbuffer: pRenderbuffer}
}
func NewGlRenderbufferStorage(
	pTarget RenderbufferTarget,
	pFormat RenderbufferFormat,
	pWidth int32,
	pHeight int32,
) *GlRenderbufferStorage {
	return &GlRenderbufferStorage{
		Target: pTarget, Format: pFormat, Width: pWidth, Height: pHeight}
}
func NewGlDeleteRenderbuffers(
	pCount int32,
	pRenderbuffers RenderbufferIdArray,
) *GlDeleteRenderbuffers {
	return &GlDeleteRenderbuffers{
		Count: pCount, Renderbuffers: pRenderbuffers}
}
func NewGlIsRenderbuffer(
	pRenderbuffer RenderbufferId,
	pResult bool,
) *GlIsRenderbuffer {
	return &GlIsRenderbuffer{
		Renderbuffer: pRenderbuffer, Result: pResult}
}
func NewGlGetRenderbufferParameteriv(
	pTarget RenderbufferTarget,
	pParameter RenderbufferParameter,
	pValues S32Array,
) *GlGetRenderbufferParameteriv {
	return &GlGetRenderbufferParameteriv{
		Target: pTarget, Parameter: pParameter, Values: pValues}
}
func NewGlGenBuffers(
	pCount int32,
	pBuffers BufferIdArray,
) *GlGenBuffers {
	return &GlGenBuffers{
		Count: pCount, Buffers: pBuffers}
}
func NewGlBindBuffer(
	pTarget BufferTarget,
	pBuffer BufferId,
) *GlBindBuffer {
	return &GlBindBuffer{
		Target: pTarget, Buffer: pBuffer}
}
func NewGlBufferData(
	pTarget BufferTarget,
	pSize int32,
	pData BufferDataPointer,
	pUsage BufferUsage,
) *GlBufferData {
	return &GlBufferData{
		Target: pTarget, Size: pSize, Data: pData, Usage: pUsage}
}
func NewGlBufferSubData(
	pTarget BufferTarget,
	pOffset int32,
	pSize int32,
	pData memory.Pointer,
) *GlBufferSubData {
	return &GlBufferSubData{
		Target: pTarget, Offset: pOffset, Size: pSize, Data: pData}
}
func NewGlDeleteBuffers(
	pCount int32,
	pBuffers BufferIdArray,
) *GlDeleteBuffers {
	return &GlDeleteBuffers{
		Count: pCount, Buffers: pBuffers}
}
func NewGlIsBuffer(
	pBuffer BufferId,
	pResult bool,
) *GlIsBuffer {
	return &GlIsBuffer{
		Buffer: pBuffer, Result: pResult}
}
func NewGlGetBufferParameteriv(
	pTarget BufferTarget,
	pParameter BufferParameter,
	pValue int32,
) *GlGetBufferParameteriv {
	return &GlGetBufferParameteriv{
		Target: pTarget, Parameter: pParameter, Value: pValue}
}
func NewGlCreateShader(
	pType ShaderType,
	pResult ShaderId,
) *GlCreateShader {
	return &GlCreateShader{
		Type: pType, Result: pResult}
}
func NewGlDeleteShader(
	pShader ShaderId,
) *GlDeleteShader {
	return &GlDeleteShader{
		Shader: pShader}
}
func NewGlShaderSource(
	pShader ShaderId,
	pCount int32,
	pSource StringArray,
	pLength S32Array,
) *GlShaderSource {
	return &GlShaderSource{
		Shader: pShader, Count: pCount, Source: pSource, Length: pLength}
}
func NewGlShaderBinary(
	pCount int32,
	pShaders ShaderIdArray,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
	pBinarySize int32,
) *GlShaderBinary {
	return &GlShaderBinary{
		Count: pCount, Shaders: pShaders, BinaryFormat: pBinaryFormat, Binary: pBinary, BinarySize: pBinarySize}
}
func NewGlGetShaderInfoLog(
	pShader ShaderId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pInfo string,
) *GlGetShaderInfoLog {
	return &GlGetShaderInfoLog{
		Shader: pShader, BufferLength: pBufferLength, StringLengthWritten: pStringLengthWritten, Info: pInfo}
}
func NewGlGetShaderSource(
	pShader ShaderId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pSource string,
) *GlGetShaderSource {
	return &GlGetShaderSource{
		Shader: pShader, BufferLength: pBufferLength, StringLengthWritten: pStringLengthWritten, Source: pSource}
}
func NewGlReleaseShaderCompiler() *GlReleaseShaderCompiler {
	return &GlReleaseShaderCompiler{}
}
func NewGlCompileShader(
	pShader ShaderId,
) *GlCompileShader {
	return &GlCompileShader{
		Shader: pShader}
}
func NewGlIsShader(
	pShader ShaderId,
	pResult bool,
) *GlIsShader {
	return &GlIsShader{
		Shader: pShader, Result: pResult}
}
func NewGlCreateProgram(
	pResult ProgramId,
) *GlCreateProgram {
	return &GlCreateProgram{
		Result: pResult}
}
func NewGlDeleteProgram(
	pProgram ProgramId,
) *GlDeleteProgram {
	return &GlDeleteProgram{
		Program: pProgram}
}
func NewGlAttachShader(
	pProgram ProgramId,
	pShader ShaderId,
) *GlAttachShader {
	return &GlAttachShader{
		Program: pProgram, Shader: pShader}
}
func NewGlDetachShader(
	pProgram ProgramId,
	pShader ShaderId,
) *GlDetachShader {
	return &GlDetachShader{
		Program: pProgram, Shader: pShader}
}
func NewGlGetAttachedShaders(
	pProgram ProgramId,
	pBufferLength int32,
	pShadersLengthWritten int32,
	pShaders ShaderIdArray,
) *GlGetAttachedShaders {
	return &GlGetAttachedShaders{
		Program: pProgram, BufferLength: pBufferLength, ShadersLengthWritten: pShadersLengthWritten, Shaders: pShaders}
}
func NewGlLinkProgram(
	pProgram ProgramId,
) *GlLinkProgram {
	return &GlLinkProgram{
		Program: pProgram}
}
func NewGlGetProgramInfoLog(
	pProgram ProgramId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pInfo string,
) *GlGetProgramInfoLog {
	return &GlGetProgramInfoLog{
		Program: pProgram, BufferLength: pBufferLength, StringLengthWritten: pStringLengthWritten, Info: pInfo}
}
func NewGlUseProgram(
	pProgram ProgramId,
) *GlUseProgram {
	return &GlUseProgram{
		Program: pProgram}
}
func NewGlIsProgram(
	pProgram ProgramId,
	pResult bool,
) *GlIsProgram {
	return &GlIsProgram{
		Program: pProgram, Result: pResult}
}
func NewGlValidateProgram(
	pProgram ProgramId,
) *GlValidateProgram {
	return &GlValidateProgram{
		Program: pProgram}
}
func NewGlClearColor(
	pR float32,
	pG float32,
	pB float32,
	pA float32,
) *GlClearColor {
	return &GlClearColor{
		R: pR, G: pG, B: pB, A: pA}
}
func NewGlClearDepthf(
	pDepth float32,
) *GlClearDepthf {
	return &GlClearDepthf{
		Depth: pDepth}
}
func NewGlClearStencil(
	pStencil int32,
) *GlClearStencil {
	return &GlClearStencil{
		Stencil: pStencil}
}
func NewGlClear(
	pMask ClearMask,
) *GlClear {
	return &GlClear{
		Mask: pMask}
}
func NewGlCullFace(
	pMode FaceMode,
) *GlCullFace {
	return &GlCullFace{
		Mode: pMode}
}
func NewGlPolygonOffset(
	pScaleFactor float32,
	pUnits float32,
) *GlPolygonOffset {
	return &GlPolygonOffset{
		ScaleFactor: pScaleFactor, Units: pUnits}
}
func NewGlLineWidth(
	pWidth float32,
) *GlLineWidth {
	return &GlLineWidth{
		Width: pWidth}
}
func NewGlSampleCoverage(
	pValue float32,
	pInvert bool,
) *GlSampleCoverage {
	return &GlSampleCoverage{
		Value: pValue, Invert: pInvert}
}
func NewGlHint(
	pTarget HintTarget,
	pMode HintMode,
) *GlHint {
	return &GlHint{
		Target: pTarget, Mode: pMode}
}
func NewGlFramebufferRenderbuffer(
	pFramebufferTarget FramebufferTarget,
	pFramebufferAttachment FramebufferAttachment,
	pRenderbufferTarget RenderbufferTarget,
	pRenderbuffer RenderbufferId,
) *GlFramebufferRenderbuffer {
	return &GlFramebufferRenderbuffer{
		FramebufferTarget: pFramebufferTarget, FramebufferAttachment: pFramebufferAttachment, RenderbufferTarget: pRenderbufferTarget, Renderbuffer: pRenderbuffer}
}
func NewGlFramebufferTexture2D(
	pFramebufferTarget FramebufferTarget,
	pFramebufferAttachment FramebufferAttachment,
	pTextureTarget TextureImageTarget,
	pTexture TextureId,
	pLevel int32,
) *GlFramebufferTexture2D {
	return &GlFramebufferTexture2D{
		FramebufferTarget: pFramebufferTarget, FramebufferAttachment: pFramebufferAttachment, TextureTarget: pTextureTarget, Texture: pTexture, Level: pLevel}
}
func NewGlGetFramebufferAttachmentParameteriv(
	pFramebufferTarget FramebufferTarget,
	pAttachment FramebufferAttachment,
	pParameter FramebufferAttachmentParameter,
	pValue S32Array,
) *GlGetFramebufferAttachmentParameteriv {
	return &GlGetFramebufferAttachmentParameteriv{
		FramebufferTarget: pFramebufferTarget, Attachment: pAttachment, Parameter: pParameter, Value: pValue}
}
func NewGlDrawElements(
	pDrawMode DrawMode,
	pElementCount int32,
	pIndicesType IndicesType,
	pIndices IndicesPointer,
) *GlDrawElements {
	return &GlDrawElements{
		DrawMode: pDrawMode, ElementCount: pElementCount, IndicesType: pIndicesType, Indices: pIndices}
}
func NewGlDrawArrays(
	pDrawMode DrawMode,
	pFirstIndex int32,
	pIndexCount int32,
) *GlDrawArrays {
	return &GlDrawArrays{
		DrawMode: pDrawMode, FirstIndex: pFirstIndex, IndexCount: pIndexCount}
}
func NewGlFlush() *GlFlush {
	return &GlFlush{}
}
func NewGlFinish() *GlFinish {
	return &GlFinish{}
}
func NewGlGetBooleanv(
	pParam StateVariable,
	pValues BoolArray,
) *GlGetBooleanv {
	return &GlGetBooleanv{
		Param: pParam, Values: pValues}
}
func NewGlGetFloatv(
	pParam StateVariable,
	pValues F32Array,
) *GlGetFloatv {
	return &GlGetFloatv{
		Param: pParam, Values: pValues}
}
func NewGlGetIntegerv(
	pParam StateVariable,
	pValues S32Array,
) *GlGetIntegerv {
	return &GlGetIntegerv{
		Param: pParam, Values: pValues}
}
func NewGlGetString(
	pParam StringConstant,
	pResult string,
) *GlGetString {
	return &GlGetString{
		Param: pParam, Result: pResult}
}
func NewGlEnable(
	pCapability Capability,
) *GlEnable {
	return &GlEnable{
		Capability: pCapability}
}
func NewGlDisable(
	pCapability Capability,
) *GlDisable {
	return &GlDisable{
		Capability: pCapability}
}
func NewGlIsEnabled(
	pCapability Capability,
	pResult bool,
) *GlIsEnabled {
	return &GlIsEnabled{
		Capability: pCapability, Result: pResult}
}
func NewGlMapBufferRange(
	pTarget BufferTarget,
	pOffset int32,
	pLength int32,
	pAccess MapBufferRangeAccess,
	pResult memory.Pointer,
) *GlMapBufferRange {
	return &GlMapBufferRange{
		Target: pTarget, Offset: pOffset, Length: pLength, Access: pAccess, Result: pResult}
}
func NewGlUnmapBuffer(
	pTarget BufferTarget,
) *GlUnmapBuffer {
	return &GlUnmapBuffer{
		Target: pTarget}
}
func NewGlInvalidateFramebuffer(
	pTarget FramebufferTarget,
	pCount int32,
	pAttachments FramebufferAttachmentArray,
) *GlInvalidateFramebuffer {
	return &GlInvalidateFramebuffer{
		Target: pTarget, Count: pCount, Attachments: pAttachments}
}
func NewGlRenderbufferStorageMultisample(
	pTarget RenderbufferTarget,
	pSamples int32,
	pFormat RenderbufferFormat,
	pWidth int32,
	pHeight int32,
) *GlRenderbufferStorageMultisample {
	return &GlRenderbufferStorageMultisample{
		Target: pTarget, Samples: pSamples, Format: pFormat, Width: pWidth, Height: pHeight}
}
func NewGlBlitFramebuffer(
	pSrcX0 int32,
	pSrcY0 int32,
	pSrcX1 int32,
	pSrcY1 int32,
	pDstX0 int32,
	pDstY0 int32,
	pDstX1 int32,
	pDstY1 int32,
	pMask ClearMask,
	pFilter TextureFilterMode,
) *GlBlitFramebuffer {
	return &GlBlitFramebuffer{
		SrcX0: pSrcX0, SrcY0: pSrcY0, SrcX1: pSrcX1, SrcY1: pSrcY1, DstX0: pDstX0, DstY0: pDstY0, DstX1: pDstX1, DstY1: pDstY1, Mask: pMask, Filter: pFilter}
}
func NewGlGenQueries(
	pCount int32,
	pQueries QueryIdArray,
) *GlGenQueries {
	return &GlGenQueries{
		Count: pCount, Queries: pQueries}
}
func NewGlBeginQuery(
	pTarget QueryTarget,
	pQuery QueryId,
) *GlBeginQuery {
	return &GlBeginQuery{
		Target: pTarget, Query: pQuery}
}
func NewGlEndQuery(
	pTarget QueryTarget,
) *GlEndQuery {
	return &GlEndQuery{
		Target: pTarget}
}
func NewGlDeleteQueries(
	pCount int32,
	pQueries QueryIdArray,
) *GlDeleteQueries {
	return &GlDeleteQueries{
		Count: pCount, Queries: pQueries}
}
func NewGlIsQuery(
	pQuery QueryId,
	pResult bool,
) *GlIsQuery {
	return &GlIsQuery{
		Query: pQuery, Result: pResult}
}
func NewGlGetQueryiv(
	pTarget QueryTarget,
	pParameter QueryParameter,
	pValue int32,
) *GlGetQueryiv {
	return &GlGetQueryiv{
		Target: pTarget, Parameter: pParameter, Value: pValue}
}
func NewGlGetQueryObjectuiv(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint32,
) *GlGetQueryObjectuiv {
	return &GlGetQueryObjectuiv{
		Query: pQuery, Parameter: pParameter, Value: pValue}
}
func NewGlGenQueriesEXT(
	pCount int32,
	pQueries QueryIdArray,
) *GlGenQueriesEXT {
	return &GlGenQueriesEXT{
		Count: pCount, Queries: pQueries}
}
func NewGlBeginQueryEXT(
	pTarget QueryTarget,
	pQuery QueryId,
) *GlBeginQueryEXT {
	return &GlBeginQueryEXT{
		Target: pTarget, Query: pQuery}
}
func NewGlEndQueryEXT(
	pTarget QueryTarget,
) *GlEndQueryEXT {
	return &GlEndQueryEXT{
		Target: pTarget}
}
func NewGlDeleteQueriesEXT(
	pCount int32,
	pQueries QueryIdArray,
) *GlDeleteQueriesEXT {
	return &GlDeleteQueriesEXT{
		Count: pCount, Queries: pQueries}
}
func NewGlIsQueryEXT(
	pQuery QueryId,
	pResult bool,
) *GlIsQueryEXT {
	return &GlIsQueryEXT{
		Query: pQuery, Result: pResult}
}
func NewGlQueryCounterEXT(
	pQuery QueryId,
	pTarget QueryTarget,
) *GlQueryCounterEXT {
	return &GlQueryCounterEXT{
		Query: pQuery, Target: pTarget}
}
func NewGlGetQueryivEXT(
	pTarget QueryTarget,
	pParameter QueryParameter,
	pValue int32,
) *GlGetQueryivEXT {
	return &GlGetQueryivEXT{
		Target: pTarget, Parameter: pParameter, Value: pValue}
}
func NewGlGetQueryObjectivEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue int32,
) *GlGetQueryObjectivEXT {
	return &GlGetQueryObjectivEXT{
		Query: pQuery, Parameter: pParameter, Value: pValue}
}
func NewGlGetQueryObjectuivEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint32,
) *GlGetQueryObjectuivEXT {
	return &GlGetQueryObjectuivEXT{
		Query: pQuery, Parameter: pParameter, Value: pValue}
}
func NewGlGetQueryObjecti64vEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue int64,
) *GlGetQueryObjecti64vEXT {
	return &GlGetQueryObjecti64vEXT{
		Query: pQuery, Parameter: pParameter, Value: pValue}
}
func NewGlGetQueryObjectui64vEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint64,
) *GlGetQueryObjectui64vEXT {
	return &GlGetQueryObjectui64vEXT{
		Query: pQuery, Parameter: pParameter, Value: pValue}
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
		Name: "ReplayCreateRenderer",
		Docs: "[]",
		ID:   0,
		New:  func() atom.Atom { return &ReplayCreateRenderer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "ReplayBindRenderer",
		Docs: "[]",
		ID:   1,
		New:  func() atom.Atom { return &ReplayBindRenderer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "BackbufferInfo",
		Docs: "[]",
		ID:   2,
		New:  func() atom.Atom { return &BackbufferInfo{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "StartTimer",
		Docs: "[]",
		ID:   3,
		New:  func() atom.Atom { return &StartTimer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "StopTimer",
		Docs: "[]",
		ID:   4,
		New:  func() atom.Atom { return &StopTimer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "FlushPostBuffer",
		Docs: "[]",
		ID:   5,
		New:  func() atom.Atom { return &FlushPostBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglInitialize",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglInitialize.xhtml]",
		ID:   6,
		New:  func() atom.Atom { return &EglInitialize{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglCreateContext",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateContext.xhtml]",
		ID:   7,
		New:  func() atom.Atom { return &EglCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglMakeCurrent",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglMakeCurrent.xhtml]",
		ID:   8,
		New:  func() atom.Atom { return &EglMakeCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglSwapBuffers",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapBuffers.xhtml]",
		ID:   9,
		New:  func() atom.Atom { return &EglSwapBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglQuerySurface",
		Docs: "[]",
		ID:   10,
		New:  func() atom.Atom { return &EglQuerySurface{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlXCreateContext",
		Docs: "[]",
		ID:   11,
		New:  func() atom.Atom { return &GlXCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlXCreateNewContext",
		Docs: "[]",
		ID:   12,
		New:  func() atom.Atom { return &GlXCreateNewContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlXMakeContextCurrent",
		Docs: "[]",
		ID:   13,
		New:  func() atom.Atom { return &GlXMakeContextCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlXSwapBuffers",
		Docs: "[]",
		ID:   14,
		New:  func() atom.Atom { return &GlXSwapBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "WglCreateContext",
		Docs: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374379(v=vs.85).aspx]",
		ID:   15,
		New:  func() atom.Atom { return &WglCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "WglCreateContextAttribsARB",
		Docs: "[http://www.opengl.org/registry/specs/ARB/wgl_create_context.txt]",
		ID:   16,
		New:  func() atom.Atom { return &WglCreateContextAttribsARB{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "WglMakeCurrent",
		Docs: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374387(v=vs.85).aspx]",
		ID:   17,
		New:  func() atom.Atom { return &WglMakeCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "WglSwapBuffers",
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
		Name: "GlEnableClientState",
		Docs: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
		ID:   21,
		New:  func() atom.Atom { return &GlEnableClientState{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDisableClientState",
		Docs: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
		ID:   22,
		New:  func() atom.Atom { return &GlDisableClientState{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetProgramBinaryOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
		ID:   23,
		New:  func() atom.Atom { return &GlGetProgramBinaryOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlProgramBinaryOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
		ID:   24,
		New:  func() atom.Atom { return &GlProgramBinaryOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStartTilingQCOM",
		Docs: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
		ID:   25,
		New:  func() atom.Atom { return &GlStartTilingQCOM{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEndTilingQCOM",
		Docs: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
		ID:   26,
		New:  func() atom.Atom { return &GlEndTilingQCOM{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDiscardFramebufferEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_discard_framebuffer.txt]",
		ID:   27,
		New:  func() atom.Atom { return &GlDiscardFramebufferEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlInsertEventMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   28,
		New:  func() atom.Atom { return &GlInsertEventMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPushGroupMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   29,
		New:  func() atom.Atom { return &GlPushGroupMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPopGroupMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   30,
		New:  func() atom.Atom { return &GlPopGroupMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexStorage1DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   31,
		New:  func() atom.Atom { return &GlTexStorage1DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexStorage2DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   32,
		New:  func() atom.Atom { return &GlTexStorage2DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexStorage3DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   33,
		New:  func() atom.Atom { return &GlTexStorage3DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTextureStorage1DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   34,
		New:  func() atom.Atom { return &GlTextureStorage1DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTextureStorage2DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   35,
		New:  func() atom.Atom { return &GlTextureStorage2DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTextureStorage3DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   36,
		New:  func() atom.Atom { return &GlTextureStorage3DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenVertexArraysOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   37,
		New:  func() atom.Atom { return &GlGenVertexArraysOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindVertexArrayOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   38,
		New:  func() atom.Atom { return &GlBindVertexArrayOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteVertexArraysOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   39,
		New:  func() atom.Atom { return &GlDeleteVertexArraysOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsVertexArrayOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   40,
		New:  func() atom.Atom { return &GlIsVertexArrayOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEGLImageTargetTexture2DOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
		ID:   41,
		New:  func() atom.Atom { return &GlEGLImageTargetTexture2DOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEGLImageTargetRenderbufferStorageOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
		ID:   42,
		New:  func() atom.Atom { return &GlEGLImageTargetRenderbufferStorageOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetGraphicsResetStatusEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_robustness.txt]",
		ID:   43,
		New:  func() atom.Atom { return &GlGetGraphicsResetStatusEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindAttribLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindAttribLocation.xml]",
		ID:   44,
		New:  func() atom.Atom { return &GlBindAttribLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendFunc",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFunc.xml]",
		ID:   45,
		New:  func() atom.Atom { return &GlBlendFunc{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendFuncSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFuncSeparate.xml]",
		ID:   46,
		New:  func() atom.Atom { return &GlBlendFuncSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendEquation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquation.xml]",
		ID:   47,
		New:  func() atom.Atom { return &GlBlendEquation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendEquationSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquationSeparate.xml]",
		ID:   48,
		New:  func() atom.Atom { return &GlBlendEquationSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendColor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendColor.xml]",
		ID:   49,
		New:  func() atom.Atom { return &GlBlendColor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEnableVertexAttribArray",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml]",
		ID:   50,
		New:  func() atom.Atom { return &GlEnableVertexAttribArray{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDisableVertexAttribArray",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml]",
		ID:   51,
		New:  func() atom.Atom { return &GlDisableVertexAttribArray{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttribPointer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttribPointer.xml]",
		ID:   52,
		New:  func() atom.Atom { return &GlVertexAttribPointer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetActiveAttrib",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveAttrib.xml]",
		ID:   53,
		New:  func() atom.Atom { return &GlGetActiveAttrib{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetActiveUniform",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveUniform.xml]",
		ID:   54,
		New:  func() atom.Atom { return &GlGetActiveUniform{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetError",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetError.xml]",
		ID:   55,
		New:  func() atom.Atom { return &GlGetError{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetProgramiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgram.xml]",
		ID:   56,
		New:  func() atom.Atom { return &GlGetProgramiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderiv.xml]",
		ID:   57,
		New:  func() atom.Atom { return &GlGetShaderiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetUniformLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniformLocation.xml]",
		ID:   58,
		New:  func() atom.Atom { return &GlGetUniformLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetAttribLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttribLocation.xml]",
		ID:   59,
		New:  func() atom.Atom { return &GlGetAttribLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPixelStorei",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPixelStorei.xml]",
		ID:   60,
		New:  func() atom.Atom { return &GlPixelStorei{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexParameteri",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
		ID:   61,
		New:  func() atom.Atom { return &GlTexParameteri{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexParameterf",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
		ID:   62,
		New:  func() atom.Atom { return &GlTexParameterf{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetTexParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
		ID:   63,
		New:  func() atom.Atom { return &GlGetTexParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetTexParameterfv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
		ID:   64,
		New:  func() atom.Atom { return &GlGetTexParameterfv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   65,
		New:  func() atom.Atom { return &GlUniform1i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   66,
		New:  func() atom.Atom { return &GlUniform2i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   67,
		New:  func() atom.Atom { return &GlUniform3i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   68,
		New:  func() atom.Atom { return &GlUniform4i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   69,
		New:  func() atom.Atom { return &GlUniform1iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   70,
		New:  func() atom.Atom { return &GlUniform2iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   71,
		New:  func() atom.Atom { return &GlUniform3iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   72,
		New:  func() atom.Atom { return &GlUniform4iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   73,
		New:  func() atom.Atom { return &GlUniform1f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   74,
		New:  func() atom.Atom { return &GlUniform2f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   75,
		New:  func() atom.Atom { return &GlUniform3f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   76,
		New:  func() atom.Atom { return &GlUniform4f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   77,
		New:  func() atom.Atom { return &GlUniform1fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   78,
		New:  func() atom.Atom { return &GlUniform2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   79,
		New:  func() atom.Atom { return &GlUniform3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   80,
		New:  func() atom.Atom { return &GlUniform4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniformMatrix2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   81,
		New:  func() atom.Atom { return &GlUniformMatrix2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniformMatrix3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   82,
		New:  func() atom.Atom { return &GlUniformMatrix3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniformMatrix4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   83,
		New:  func() atom.Atom { return &GlUniformMatrix4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetUniformfv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
		ID:   84,
		New:  func() atom.Atom { return &GlGetUniformfv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetUniformiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
		ID:   85,
		New:  func() atom.Atom { return &GlGetUniformiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib1f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   86,
		New:  func() atom.Atom { return &GlVertexAttrib1f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib2f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   87,
		New:  func() atom.Atom { return &GlVertexAttrib2f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib3f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   88,
		New:  func() atom.Atom { return &GlVertexAttrib3f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib4f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   89,
		New:  func() atom.Atom { return &GlVertexAttrib4f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib1fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   90,
		New:  func() atom.Atom { return &GlVertexAttrib1fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   91,
		New:  func() atom.Atom { return &GlVertexAttrib2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   92,
		New:  func() atom.Atom { return &GlVertexAttrib3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   93,
		New:  func() atom.Atom { return &GlVertexAttrib4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderPrecisionFormat",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml]",
		ID:   94,
		New:  func() atom.Atom { return &GlGetShaderPrecisionFormat{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDepthMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthMask.xml]",
		ID:   95,
		New:  func() atom.Atom { return &GlDepthMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDepthFunc",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthFunc.xml]",
		ID:   96,
		New:  func() atom.Atom { return &GlDepthFunc{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDepthRangef",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthRangef.xml]",
		ID:   97,
		New:  func() atom.Atom { return &GlDepthRangef{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlColorMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glColorMask.xml]",
		ID:   98,
		New:  func() atom.Atom { return &GlColorMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMask.xml]",
		ID:   99,
		New:  func() atom.Atom { return &GlStencilMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilMaskSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMaskSeparate.xml]",
		ID:   100,
		New:  func() atom.Atom { return &GlStencilMaskSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilFuncSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilFuncSeparate.xml]",
		ID:   101,
		New:  func() atom.Atom { return &GlStencilFuncSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilOpSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilOpSeparate.xml]",
		ID:   102,
		New:  func() atom.Atom { return &GlStencilOpSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFrontFace",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFrontFace.xml]",
		ID:   103,
		New:  func() atom.Atom { return &GlFrontFace{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlViewport",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glViewport.xml]",
		ID:   104,
		New:  func() atom.Atom { return &GlViewport{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlScissor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glScissor.xml]",
		ID:   105,
		New:  func() atom.Atom { return &GlScissor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlActiveTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glActiveTexture.xml]",
		ID:   106,
		New:  func() atom.Atom { return &GlActiveTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenTextures",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenTextures.xml]",
		ID:   107,
		New:  func() atom.Atom { return &GlGenTextures{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteTextures",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteTextures.xml]",
		ID:   108,
		New:  func() atom.Atom { return &GlDeleteTextures{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsTexture.xml]",
		ID:   109,
		New:  func() atom.Atom { return &GlIsTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindTexture.xml]",
		ID:   110,
		New:  func() atom.Atom { return &GlBindTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexImage2D.xml]",
		ID:   111,
		New:  func() atom.Atom { return &GlTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexSubImage2D.xml]",
		ID:   112,
		New:  func() atom.Atom { return &GlTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCopyTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexImage2D.xml]",
		ID:   113,
		New:  func() atom.Atom { return &GlCopyTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCopyTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml]",
		ID:   114,
		New:  func() atom.Atom { return &GlCopyTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCompressedTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexImage2D.xml]",
		ID:   115,
		New:  func() atom.Atom { return &GlCompressedTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCompressedTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml]",
		ID:   116,
		New:  func() atom.Atom { return &GlCompressedTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenerateMipmap",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenerateMipmap.xml]",
		ID:   117,
		New:  func() atom.Atom { return &GlGenerateMipmap{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlReadPixels",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReadPixels.xml]",
		ID:   118,
		New:  func() atom.Atom { return &GlReadPixels{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenFramebuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenFramebuffers.xml]",
		ID:   119,
		New:  func() atom.Atom { return &GlGenFramebuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindFramebuffer.xml]",
		ID:   120,
		New:  func() atom.Atom { return &GlBindFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCheckFramebufferStatus",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml]",
		ID:   121,
		New:  func() atom.Atom { return &GlCheckFramebufferStatus{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteFramebuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteFramebuffers.xml]",
		ID:   122,
		New:  func() atom.Atom { return &GlDeleteFramebuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsFramebuffer.xml]",
		ID:   123,
		New:  func() atom.Atom { return &GlIsFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenRenderbuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenRenderbuffers.xml]",
		ID:   124,
		New:  func() atom.Atom { return &GlGenRenderbuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindRenderbuffer.xml]",
		ID:   125,
		New:  func() atom.Atom { return &GlBindRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlRenderbufferStorage",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glRenderbufferStorage.xml]",
		ID:   126,
		New:  func() atom.Atom { return &GlRenderbufferStorage{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteRenderbuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml]",
		ID:   127,
		New:  func() atom.Atom { return &GlDeleteRenderbuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsRenderbuffer.xml]",
		ID:   128,
		New:  func() atom.Atom { return &GlIsRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetRenderbufferParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml]",
		ID:   129,
		New:  func() atom.Atom { return &GlGetRenderbufferParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenBuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenBuffers.xml]",
		ID:   130,
		New:  func() atom.Atom { return &GlGenBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindBuffer.xml]",
		ID:   131,
		New:  func() atom.Atom { return &GlBindBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBufferData",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferData.xml]",
		ID:   132,
		New:  func() atom.Atom { return &GlBufferData{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBufferSubData",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferSubData.xml]",
		ID:   133,
		New:  func() atom.Atom { return &GlBufferSubData{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteBuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteBuffers.xml]",
		ID:   134,
		New:  func() atom.Atom { return &GlDeleteBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsBuffer.xml]",
		ID:   135,
		New:  func() atom.Atom { return &GlIsBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetBufferParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetBufferParameteriv.xml]",
		ID:   136,
		New:  func() atom.Atom { return &GlGetBufferParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCreateShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateShader.xml]",
		ID:   137,
		New:  func() atom.Atom { return &GlCreateShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteShader.xml]",
		ID:   138,
		New:  func() atom.Atom { return &GlDeleteShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlShaderSource",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderSource.xml]",
		ID:   139,
		New:  func() atom.Atom { return &GlShaderSource{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlShaderBinary",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderBinary.xml]",
		ID:   140,
		New:  func() atom.Atom { return &GlShaderBinary{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderInfoLog",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderInfoLog.xml]",
		ID:   141,
		New:  func() atom.Atom { return &GlGetShaderInfoLog{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderSource",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderSource.xml]",
		ID:   142,
		New:  func() atom.Atom { return &GlGetShaderSource{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlReleaseShaderCompiler",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml]",
		ID:   143,
		New:  func() atom.Atom { return &GlReleaseShaderCompiler{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCompileShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompileShader.xml]",
		ID:   144,
		New:  func() atom.Atom { return &GlCompileShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsShader.xml]",
		ID:   145,
		New:  func() atom.Atom { return &GlIsShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCreateProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateProgram.xml]",
		ID:   146,
		New:  func() atom.Atom { return &GlCreateProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteProgram.xml]",
		ID:   147,
		New:  func() atom.Atom { return &GlDeleteProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlAttachShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glAttachShader.xml]",
		ID:   148,
		New:  func() atom.Atom { return &GlAttachShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDetachShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDetachShader.xml]",
		ID:   149,
		New:  func() atom.Atom { return &GlDetachShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetAttachedShaders",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttachedShaders.xml]",
		ID:   150,
		New:  func() atom.Atom { return &GlGetAttachedShaders{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlLinkProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLinkProgram.xml]",
		ID:   151,
		New:  func() atom.Atom { return &GlLinkProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetProgramInfoLog",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgramInfoLog.xml]",
		ID:   152,
		New:  func() atom.Atom { return &GlGetProgramInfoLog{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUseProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUseProgram.xml]",
		ID:   153,
		New:  func() atom.Atom { return &GlUseProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsProgram.xml]",
		ID:   154,
		New:  func() atom.Atom { return &GlIsProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlValidateProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glValidateProgram.xml]",
		ID:   155,
		New:  func() atom.Atom { return &GlValidateProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClearColor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearColor.xml]",
		ID:   156,
		New:  func() atom.Atom { return &GlClearColor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClearDepthf",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearDepthf.xml]",
		ID:   157,
		New:  func() atom.Atom { return &GlClearDepthf{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClearStencil",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearStencil.xml]",
		ID:   158,
		New:  func() atom.Atom { return &GlClearStencil{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClear",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClear.xml]",
		ID:   159,
		New:  func() atom.Atom { return &GlClear{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCullFace",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCullFace.xml]",
		ID:   160,
		New:  func() atom.Atom { return &GlCullFace{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPolygonOffset",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPolygonOffset.xml]",
		ID:   161,
		New:  func() atom.Atom { return &GlPolygonOffset{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlLineWidth",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLineWidth.xml]",
		ID:   162,
		New:  func() atom.Atom { return &GlLineWidth{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlSampleCoverage",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glSampleCoverage.xml]",
		ID:   163,
		New:  func() atom.Atom { return &GlSampleCoverage{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlHint",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glHint.xml]",
		ID:   164,
		New:  func() atom.Atom { return &GlHint{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFramebufferRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml]",
		ID:   165,
		New:  func() atom.Atom { return &GlFramebufferRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFramebufferTexture2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferTexture2D.xml]",
		ID:   166,
		New:  func() atom.Atom { return &GlFramebufferTexture2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetFramebufferAttachmentParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml]",
		ID:   167,
		New:  func() atom.Atom { return &GlGetFramebufferAttachmentParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDrawElements",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawElements.xml]",
		ID:   168,
		New:  func() atom.Atom { return &GlDrawElements{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDrawArrays",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawArrays.xml]",
		ID:   169,
		New:  func() atom.Atom { return &GlDrawArrays{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFlush",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFlush.xml]",
		ID:   170,
		New:  func() atom.Atom { return &GlFlush{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFinish",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFinish.xml]",
		ID:   171,
		New:  func() atom.Atom { return &GlFinish{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetBooleanv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   172,
		New:  func() atom.Atom { return &GlGetBooleanv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetFloatv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   173,
		New:  func() atom.Atom { return &GlGetFloatv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetIntegerv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   174,
		New:  func() atom.Atom { return &GlGetIntegerv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetString",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetString.xml]",
		ID:   175,
		New:  func() atom.Atom { return &GlGetString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEnable",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnable.xml]",
		ID:   176,
		New:  func() atom.Atom { return &GlEnable{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDisable",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisable.xml]",
		ID:   177,
		New:  func() atom.Atom { return &GlDisable{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsEnabled",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsEnabled.xml]",
		ID:   178,
		New:  func() atom.Atom { return &GlIsEnabled{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlMapBufferRange",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
		ID:   179,
		New:  func() atom.Atom { return &GlMapBufferRange{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUnmapBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
		ID:   180,
		New:  func() atom.Atom { return &GlUnmapBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlInvalidateFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glInvalidateFramebuffer.xhtml]",
		ID:   181,
		New:  func() atom.Atom { return &GlInvalidateFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlRenderbufferStorageMultisample",
		Docs: "[http://www.opengl.org/registry/specs/EXT/framebuffer_multisample.txt]",
		ID:   182,
		New:  func() atom.Atom { return &GlRenderbufferStorageMultisample{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlitFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBlitFramebuffer.xhtml]",
		ID:   183,
		New:  func() atom.Atom { return &GlBlitFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenQueries",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenQueries.xhtml]",
		ID:   184,
		New:  func() atom.Atom { return &GlGenQueries{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBeginQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBeginQuery.xhtml]",
		ID:   185,
		New:  func() atom.Atom { return &GlBeginQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEndQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glEndQuery.xhtml]",
		ID:   186,
		New:  func() atom.Atom { return &GlEndQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteQueries",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteQueries.xhtml]",
		ID:   187,
		New:  func() atom.Atom { return &GlDeleteQueries{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glIsQuery.xhtml]",
		ID:   188,
		New:  func() atom.Atom { return &GlIsQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryiv.xhtml]",
		ID:   189,
		New:  func() atom.Atom { return &GlGetQueryiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectuiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryObjectuiv.xhtml]",
		ID:   190,
		New:  func() atom.Atom { return &GlGetQueryObjectuiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenQueriesEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   191,
		New:  func() atom.Atom { return &GlGenQueriesEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBeginQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   192,
		New:  func() atom.Atom { return &GlBeginQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEndQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   193,
		New:  func() atom.Atom { return &GlEndQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteQueriesEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   194,
		New:  func() atom.Atom { return &GlDeleteQueriesEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   195,
		New:  func() atom.Atom { return &GlIsQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlQueryCounterEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   196,
		New:  func() atom.Atom { return &GlQueryCounterEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   197,
		New:  func() atom.Atom { return &GlGetQueryivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   198,
		New:  func() atom.Atom { return &GlGetQueryObjectivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectuivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   199,
		New:  func() atom.Atom { return &GlGetQueryObjectuivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjecti64vEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   200,
		New:  func() atom.Atom { return &GlGetQueryObjecti64vEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectui64vEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   201,
		New:  func() atom.Atom { return &GlGetQueryObjectui64vEXT{} },
	})
}
