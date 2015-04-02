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
	"android.googlesource.com/platform/tools/gpu/service"
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

type HGLRC memory.Pointer

func (c *HGLRC) Less(rhs HGLRC) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *HGLRC) Equal(rhs HGLRC) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type HDC memory.Pointer

func (c *HDC) Less(rhs HDC) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *HDC) Equal(rhs HDC) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type BOOL int64

func (c *BOOL) Less(rhs BOOL) bool  { return int64(*c) < int64(rhs) }
func (c *BOOL) Equal(rhs BOOL) bool { return int64(*c) == int64(rhs) }

type ImageOES memory.Pointer

func (c *ImageOES) Less(rhs ImageOES) bool  { return memory.Pointer(*c) < memory.Pointer(rhs) }
func (c *ImageOES) Equal(rhs ImageOES) bool { return memory.Pointer(*c) == memory.Pointer(rhs) }

type BoolArray []bool

func (s BoolArray) Len() int      { return len(s) }
func (s BoolArray) Range() []bool { return s }

type BufferIdArray []BufferId

func (s BufferIdArray) Len() int          { return len(s) }
func (s BufferIdArray) Range() []BufferId { return s }

type DiscardFramebufferAttachmentArray []DiscardFramebufferAttachment

func (s DiscardFramebufferAttachmentArray) Len() int                              { return len(s) }
func (s DiscardFramebufferAttachmentArray) Range() []DiscardFramebufferAttachment { return s }

type F32Array []float32

func (s F32Array) Len() int         { return len(s) }
func (s F32Array) Range() []float32 { return s }

type FramebufferAttachmentArray []FramebufferAttachment

func (s FramebufferAttachmentArray) Len() int                       { return len(s) }
func (s FramebufferAttachmentArray) Range() []FramebufferAttachment { return s }

type FramebufferIdArray []FramebufferId

func (s FramebufferIdArray) Len() int               { return len(s) }
func (s FramebufferIdArray) Range() []FramebufferId { return s }

type QueryIdArray []QueryId

func (s QueryIdArray) Len() int         { return len(s) }
func (s QueryIdArray) Range() []QueryId { return s }

type RenderbufferIdArray []RenderbufferId

func (s RenderbufferIdArray) Len() int                { return len(s) }
func (s RenderbufferIdArray) Range() []RenderbufferId { return s }

type S32Array []int32

func (s S32Array) Len() int       { return len(s) }
func (s S32Array) Range() []int32 { return s }

type ShaderIdArray []ShaderId

func (s ShaderIdArray) Len() int          { return len(s) }
func (s ShaderIdArray) Range() []ShaderId { return s }

type StringArray []string

func (s StringArray) Len() int        { return len(s) }
func (s StringArray) Range() []string { return s }

type TextureIdArray []TextureId

func (s TextureIdArray) Len() int           { return len(s) }
func (s TextureIdArray) Range() []TextureId { return s }

type VertexArrayIdArray []VertexArrayId

func (s VertexArrayIdArray) Len() int               { return len(s) }
func (s VertexArrayIdArray) Range() []VertexArrayId { return s }

type AttributeLocation_stringMap map[string]AttributeLocation

func (m AttributeLocation_stringMap) Get(key string) AttributeLocation {
	return m[key]
}
func (m AttributeLocation_stringMap) Contains(key string) bool {
	_, ok := m[key]
	return ok
}
func (m AttributeLocation_stringMap) Delete(key string) {
	delete(m, key)
}
func (m AttributeLocation_stringMap) Range() []AttributeLocation {
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

type BufferRef_BufferIdMap map[BufferId]*Buffer

func (m BufferRef_BufferIdMap) Get(key BufferId) *Buffer {
	return m[key]
}
func (m BufferRef_BufferIdMap) Contains(key BufferId) bool {
	_, ok := m[key]
	return ok
}
func (m BufferRef_BufferIdMap) Delete(key BufferId) {
	delete(m, key)
}
func (m BufferRef_BufferIdMap) Range() []*Buffer {
	values := make([]*Buffer, 0, len(m))
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

type FramebufferRef_FramebufferIdMap map[FramebufferId]*Framebuffer

func (m FramebufferRef_FramebufferIdMap) Get(key FramebufferId) *Framebuffer {
	return m[key]
}
func (m FramebufferRef_FramebufferIdMap) Contains(key FramebufferId) bool {
	_, ok := m[key]
	return ok
}
func (m FramebufferRef_FramebufferIdMap) Delete(key FramebufferId) {
	delete(m, key)
}
func (m FramebufferRef_FramebufferIdMap) Range() []*Framebuffer {
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

type ProgramRef_ProgramIdMap map[ProgramId]*Program

func (m ProgramRef_ProgramIdMap) Get(key ProgramId) *Program {
	return m[key]
}
func (m ProgramRef_ProgramIdMap) Contains(key ProgramId) bool {
	_, ok := m[key]
	return ok
}
func (m ProgramRef_ProgramIdMap) Delete(key ProgramId) {
	delete(m, key)
}
func (m ProgramRef_ProgramIdMap) Range() []*Program {
	values := make([]*Program, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type QueryRef_QueryIdMap map[QueryId]*Query

func (m QueryRef_QueryIdMap) Get(key QueryId) *Query {
	return m[key]
}
func (m QueryRef_QueryIdMap) Contains(key QueryId) bool {
	_, ok := m[key]
	return ok
}
func (m QueryRef_QueryIdMap) Delete(key QueryId) {
	delete(m, key)
}
func (m QueryRef_QueryIdMap) Range() []*Query {
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

type RenderbufferRef_RenderbufferIdMap map[RenderbufferId]*Renderbuffer

func (m RenderbufferRef_RenderbufferIdMap) Get(key RenderbufferId) *Renderbuffer {
	return m[key]
}
func (m RenderbufferRef_RenderbufferIdMap) Contains(key RenderbufferId) bool {
	_, ok := m[key]
	return ok
}
func (m RenderbufferRef_RenderbufferIdMap) Delete(key RenderbufferId) {
	delete(m, key)
}
func (m RenderbufferRef_RenderbufferIdMap) Range() []*Renderbuffer {
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

type ShaderRef_ShaderIdMap map[ShaderId]*Shader

func (m ShaderRef_ShaderIdMap) Get(key ShaderId) *Shader {
	return m[key]
}
func (m ShaderRef_ShaderIdMap) Contains(key ShaderId) bool {
	_, ok := m[key]
	return ok
}
func (m ShaderRef_ShaderIdMap) Delete(key ShaderId) {
	delete(m, key)
}
func (m ShaderRef_ShaderIdMap) Range() []*Shader {
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

type TextureRef_TextureIdMap map[TextureId]*Texture

func (m TextureRef_TextureIdMap) Get(key TextureId) *Texture {
	return m[key]
}
func (m TextureRef_TextureIdMap) Contains(key TextureId) bool {
	_, ok := m[key]
	return ok
}
func (m TextureRef_TextureIdMap) Delete(key TextureId) {
	delete(m, key)
}
func (m TextureRef_TextureIdMap) Range() []*Texture {
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

type VertexArrayRef_VertexArrayIdMap map[VertexArrayId]*VertexArray

func (m VertexArrayRef_VertexArrayIdMap) Get(key VertexArrayId) *VertexArray {
	return m[key]
}
func (m VertexArrayRef_VertexArrayIdMap) Contains(key VertexArrayId) bool {
	_, ok := m[key]
	return ok
}
func (m VertexArrayRef_VertexArrayIdMap) Delete(key VertexArrayId) {
	delete(m, key)
}
func (m VertexArrayRef_VertexArrayIdMap) Range() []*VertexArray {
	values := make([]*VertexArray, 0, len(m))
	for _, value := range m {
		values = append(values, value)
	}
	return values
}

type VertexAttributeArrayRef_AttributeLocationMap map[AttributeLocation]*VertexAttributeArray

func (m VertexAttributeArrayRef_AttributeLocationMap) Get(key AttributeLocation) *VertexAttributeArray {
	return m[key]
}
func (m VertexAttributeArrayRef_AttributeLocationMap) Contains(key AttributeLocation) bool {
	_, ok := m[key]
	return ok
}
func (m VertexAttributeArrayRef_AttributeLocationMap) Delete(key AttributeLocation) {
	delete(m, key)
}
func (m VertexAttributeArrayRef_AttributeLocationMap) Range() []*VertexAttributeArray {
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
// Init
////////////////////////////////////////////////////////////////////////////////
type Init_In struct {
	binary.Generate
	Width      int32
	Height     int32
	ColorFmt   RenderbufferFormat
	DepthFmt   RenderbufferFormat
	StencilFmt RenderbufferFormat
}
type Init_Out struct {
	binary.Generate
}
type Init struct {
	binary.Generate
	Context atom.ContextID
	In      Init_In
	Out     Init_Out
}

func (c *Init) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "init(",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		c.In.ColorFmt.String(),
		", ",
		c.In.DepthFmt.String(),
		", ",
		c.In.StencilFmt.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *Init) ContextID() atom.ContextID {
	return c.Context
}
func (c *Init) TypeID() atom.TypeID {
	return 0
}
func (c *Init) Flags() atom.Flags {
	return 0
}
func (Init) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// StartTimer
////////////////////////////////////////////////////////////////////////////////
type StartTimer_In struct {
	binary.Generate
	Index uint8
}
type StartTimer_Out struct {
	binary.Generate
}
type StartTimer struct {
	binary.Generate
	Context atom.ContextID
	In      StartTimer_In
	Out     StartTimer_Out
}

func (c *StartTimer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "startTimer(",
		fmt.Sprintf("index:%v", c.In.Index),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *StartTimer) ContextID() atom.ContextID {
	return c.Context
}
func (c *StartTimer) TypeID() atom.TypeID {
	return 1
}
func (c *StartTimer) Flags() atom.Flags {
	return 0
}
func (StartTimer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// StopTimer
////////////////////////////////////////////////////////////////////////////////
type StopTimer_In struct {
	binary.Generate
	Index uint8
}
type StopTimer_Out struct {
	binary.Generate
	Result uint64
}
type StopTimer struct {
	binary.Generate
	Context atom.ContextID
	In      StopTimer_In
	Out     StopTimer_Out
}

func (c *StopTimer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "stopTimer(",
		fmt.Sprintf("index:%v", c.In.Index),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *StopTimer) ContextID() atom.ContextID {
	return c.Context
}
func (c *StopTimer) TypeID() atom.TypeID {
	return 2
}
func (c *StopTimer) Flags() atom.Flags {
	return 0
}
func (StopTimer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// FlushPostBuffer
////////////////////////////////////////////////////////////////////////////////
type FlushPostBuffer_In struct {
	binary.Generate
}
type FlushPostBuffer_Out struct {
	binary.Generate
}
type FlushPostBuffer struct {
	binary.Generate
	Context atom.ContextID
	In      FlushPostBuffer_In
	Out     FlushPostBuffer_Out
}

func (c *FlushPostBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "flushPostBuffer(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *FlushPostBuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *FlushPostBuffer) TypeID() atom.TypeID {
	return 3
}
func (c *FlushPostBuffer) Flags() atom.Flags {
	return 0
}
func (FlushPostBuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// EglCreateContext
////////////////////////////////////////////////////////////////////////////////
type EglCreateContext_In struct {
	binary.Generate
}
type EglCreateContext_Out struct {
	binary.Generate
	Version int32
	Context int32
}
type EglCreateContext struct {
	binary.Generate
	Context atom.ContextID
	In      EglCreateContext_In
	Out     EglCreateContext_Out
}

func (c *EglCreateContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglCreateContext(",
		fmt.Sprintf("version:%v", c.Out.Version),
		", ",
		fmt.Sprintf("context:%v", c.Out.Context),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *EglCreateContext) ContextID() atom.ContextID {
	return c.Context
}
func (c *EglCreateContext) TypeID() atom.TypeID {
	return 4
}
func (c *EglCreateContext) Flags() atom.Flags {
	return 0
}
func (EglCreateContext) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// EglMakeCurrent
////////////////////////////////////////////////////////////////////////////////
type EglMakeCurrent_In struct {
	binary.Generate
	Context int32
}
type EglMakeCurrent_Out struct {
	binary.Generate
}
type EglMakeCurrent struct {
	binary.Generate
	Context atom.ContextID
	In      EglMakeCurrent_In
	Out     EglMakeCurrent_Out
}

func (c *EglMakeCurrent) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglMakeCurrent(",
		fmt.Sprintf("context:%v", c.In.Context),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *EglMakeCurrent) ContextID() atom.ContextID {
	return c.Context
}
func (c *EglMakeCurrent) TypeID() atom.TypeID {
	return 5
}
func (c *EglMakeCurrent) Flags() atom.Flags {
	return 0
}
func (EglMakeCurrent) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// EglSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type EglSwapBuffers_In struct {
	binary.Generate
}
type EglSwapBuffers_Out struct {
	binary.Generate
}
type EglSwapBuffers struct {
	binary.Generate
	Context atom.ContextID
	In      EglSwapBuffers_In
	Out     EglSwapBuffers_Out
}

func (c *EglSwapBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "eglSwapBuffers(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *EglSwapBuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *EglSwapBuffers) TypeID() atom.TypeID {
	return 6
}
func (c *EglSwapBuffers) Flags() atom.Flags {
	return 0 | atom.EndOfFrame
}
func (EglSwapBuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// WglCreateContext
////////////////////////////////////////////////////////////////////////////////
type WglCreateContext_In struct {
	binary.Generate
	Hdc HDC
}
type WglCreateContext_Out struct {
	binary.Generate
	Result HGLRC
}
type WglCreateContext struct {
	binary.Generate
	Context atom.ContextID
	In      WglCreateContext_In
	Out     WglCreateContext_Out
}

func (c *WglCreateContext) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "wglCreateContext(",
		fmt.Sprintf("hdc:%v", c.In.Hdc),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *WglCreateContext) ContextID() atom.ContextID {
	return c.Context
}
func (c *WglCreateContext) TypeID() atom.TypeID {
	return 7
}
func (c *WglCreateContext) Flags() atom.Flags {
	return 0
}
func (WglCreateContext) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// WglMakeCurrent
////////////////////////////////////////////////////////////////////////////////
type WglMakeCurrent_In struct {
	binary.Generate
	Hdc   HDC
	Hglrc HGLRC
}
type WglMakeCurrent_Out struct {
	binary.Generate
	Result BOOL
}
type WglMakeCurrent struct {
	binary.Generate
	Context atom.ContextID
	In      WglMakeCurrent_In
	Out     WglMakeCurrent_Out
}

func (c *WglMakeCurrent) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "wglMakeCurrent(",
		fmt.Sprintf("hdc:%v", c.In.Hdc),
		", ",
		fmt.Sprintf("hglrc:%v", c.In.Hglrc),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *WglMakeCurrent) ContextID() atom.ContextID {
	return c.Context
}
func (c *WglMakeCurrent) TypeID() atom.TypeID {
	return 8
}
func (c *WglMakeCurrent) Flags() atom.Flags {
	return 0
}
func (WglMakeCurrent) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// WglSwapBuffers
////////////////////////////////////////////////////////////////////////////////
type WglSwapBuffers_In struct {
	binary.Generate
	Hdc HDC
}
type WglSwapBuffers_Out struct {
	binary.Generate
}
type WglSwapBuffers struct {
	binary.Generate
	Context atom.ContextID
	In      WglSwapBuffers_In
	Out     WglSwapBuffers_Out
}

func (c *WglSwapBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "wglSwapBuffers(",
		fmt.Sprintf("hdc:%v", c.In.Hdc),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *WglSwapBuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *WglSwapBuffers) TypeID() atom.TypeID {
	return 9
}
func (c *WglSwapBuffers) Flags() atom.Flags {
	return 0 | atom.EndOfFrame
}
func (WglSwapBuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEnableClientState
////////////////////////////////////////////////////////////////////////////////
type GlEnableClientState_In struct {
	binary.Generate
	Type ArrayType
}
type GlEnableClientState_Out struct {
	binary.Generate
}
type GlEnableClientState struct {
	binary.Generate
	Context atom.ContextID
	In      GlEnableClientState_In
	Out     GlEnableClientState_Out
}

func (c *GlEnableClientState) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEnableClientState(",
		c.In.Type.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEnableClientState) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEnableClientState) TypeID() atom.TypeID {
	return 10
}
func (c *GlEnableClientState) Flags() atom.Flags {
	return 0
}
func (GlEnableClientState) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDisableClientState
////////////////////////////////////////////////////////////////////////////////
type GlDisableClientState_In struct {
	binary.Generate
	Type ArrayType
}
type GlDisableClientState_Out struct {
	binary.Generate
}
type GlDisableClientState struct {
	binary.Generate
	Context atom.ContextID
	In      GlDisableClientState_In
	Out     GlDisableClientState_Out
}

func (c *GlDisableClientState) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDisableClientState(",
		c.In.Type.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDisableClientState) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDisableClientState) TypeID() atom.TypeID {
	return 11
}
func (c *GlDisableClientState) Flags() atom.Flags {
	return 0
}
func (GlDisableClientState) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramBinaryOES
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramBinaryOES_In struct {
	binary.Generate
	Program    ProgramId
	BufferSize int32
}
type GlGetProgramBinaryOES_Out struct {
	binary.Generate
	BytesWritten int32
	BinaryFormat uint32
	Binary       memory.Pointer
}
type GlGetProgramBinaryOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetProgramBinaryOES_In
	Out     GlGetProgramBinaryOES_Out
}

func (c *GlGetProgramBinaryOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetProgramBinaryOES(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("buffer_size:%v", c.In.BufferSize),
		", ",
		fmt.Sprintf("bytes_written:%v", c.Out.BytesWritten),
		", ",
		fmt.Sprintf("binary_format:%v", c.Out.BinaryFormat),
		", ",
		fmt.Sprintf("0x%x", c.Out.Binary),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetProgramBinaryOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetProgramBinaryOES) TypeID() atom.TypeID {
	return 12
}
func (c *GlGetProgramBinaryOES) Flags() atom.Flags {
	return 0
}
func (GlGetProgramBinaryOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlProgramBinaryOES
////////////////////////////////////////////////////////////////////////////////
type GlProgramBinaryOES_In struct {
	binary.Generate
	Program      ProgramId
	BinaryFormat uint32
	Binary       memory.Pointer
	BinarySize   int32
}
type GlProgramBinaryOES_Out struct {
	binary.Generate
}
type GlProgramBinaryOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlProgramBinaryOES_In
	Out     GlProgramBinaryOES_Out
}

func (c *GlProgramBinaryOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glProgramBinaryOES(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("binary_format:%v", c.In.BinaryFormat),
		", ",
		fmt.Sprintf("0x%x", c.In.Binary),
		", ",
		fmt.Sprintf("binary_size:%v", c.In.BinarySize),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlProgramBinaryOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlProgramBinaryOES) TypeID() atom.TypeID {
	return 13
}
func (c *GlProgramBinaryOES) Flags() atom.Flags {
	return 0
}
func (GlProgramBinaryOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlStartTilingQCOM
////////////////////////////////////////////////////////////////////////////////
type GlStartTilingQCOM_In struct {
	binary.Generate
	X            int32
	Y            int32
	Width        int32
	Height       int32
	PreserveMask TilePreserveMaskQCOM
}
type GlStartTilingQCOM_Out struct {
	binary.Generate
}
type GlStartTilingQCOM struct {
	binary.Generate
	Context atom.ContextID
	In      GlStartTilingQCOM_In
	Out     GlStartTilingQCOM_Out
}

func (c *GlStartTilingQCOM) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStartTilingQCOM(",
		fmt.Sprintf("x:%v", c.In.X),
		", ",
		fmt.Sprintf("y:%v", c.In.Y),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		c.In.PreserveMask.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStartTilingQCOM) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlStartTilingQCOM) TypeID() atom.TypeID {
	return 14
}
func (c *GlStartTilingQCOM) Flags() atom.Flags {
	return 0
}
func (GlStartTilingQCOM) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEndTilingQCOM
////////////////////////////////////////////////////////////////////////////////
type GlEndTilingQCOM_In struct {
	binary.Generate
	PreserveMask TilePreserveMaskQCOM
}
type GlEndTilingQCOM_Out struct {
	binary.Generate
}
type GlEndTilingQCOM struct {
	binary.Generate
	Context atom.ContextID
	In      GlEndTilingQCOM_In
	Out     GlEndTilingQCOM_Out
}

func (c *GlEndTilingQCOM) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEndTilingQCOM(",
		c.In.PreserveMask.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEndTilingQCOM) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEndTilingQCOM) TypeID() atom.TypeID {
	return 15
}
func (c *GlEndTilingQCOM) Flags() atom.Flags {
	return 0
}
func (GlEndTilingQCOM) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDiscardFramebufferEXT
////////////////////////////////////////////////////////////////////////////////
type GlDiscardFramebufferEXT_In struct {
	binary.Generate
	Target         FramebufferTarget
	NumAttachments int32
	Attachments    DiscardFramebufferAttachmentArray
}
type GlDiscardFramebufferEXT_Out struct {
	binary.Generate
}
type GlDiscardFramebufferEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlDiscardFramebufferEXT_In
	Out     GlDiscardFramebufferEXT_Out
}

func (c *GlDiscardFramebufferEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDiscardFramebufferEXT(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("numAttachments:%v", c.In.NumAttachments),
		", ",
		fmt.Sprintf("%v", c.In.Attachments),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDiscardFramebufferEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDiscardFramebufferEXT) TypeID() atom.TypeID {
	return 16
}
func (c *GlDiscardFramebufferEXT) Flags() atom.Flags {
	return 0
}
func (GlDiscardFramebufferEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlInsertEventMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlInsertEventMarkerEXT_In struct {
	binary.Generate
	Length int32
	Marker string
}
type GlInsertEventMarkerEXT_Out struct {
	binary.Generate
}
type GlInsertEventMarkerEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlInsertEventMarkerEXT_In
	Out     GlInsertEventMarkerEXT_Out
}

func (c *GlInsertEventMarkerEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glInsertEventMarkerEXT(",
		fmt.Sprintf("length:%v", c.In.Length),
		", ",
		fmt.Sprintf("marker:%v", c.In.Marker),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlInsertEventMarkerEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlInsertEventMarkerEXT) TypeID() atom.TypeID {
	return 17
}
func (c *GlInsertEventMarkerEXT) Flags() atom.Flags {
	return 0
}
func (GlInsertEventMarkerEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlPushGroupMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlPushGroupMarkerEXT_In struct {
	binary.Generate
	Length int32
	Marker string
}
type GlPushGroupMarkerEXT_Out struct {
	binary.Generate
}
type GlPushGroupMarkerEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlPushGroupMarkerEXT_In
	Out     GlPushGroupMarkerEXT_Out
}

func (c *GlPushGroupMarkerEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPushGroupMarkerEXT(",
		fmt.Sprintf("length:%v", c.In.Length),
		", ",
		fmt.Sprintf("marker:%v", c.In.Marker),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPushGroupMarkerEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlPushGroupMarkerEXT) TypeID() atom.TypeID {
	return 18
}
func (c *GlPushGroupMarkerEXT) Flags() atom.Flags {
	return 0
}
func (GlPushGroupMarkerEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlPopGroupMarkerEXT
////////////////////////////////////////////////////////////////////////////////
type GlPopGroupMarkerEXT_In struct {
	binary.Generate
}
type GlPopGroupMarkerEXT_Out struct {
	binary.Generate
}
type GlPopGroupMarkerEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlPopGroupMarkerEXT_In
	Out     GlPopGroupMarkerEXT_Out
}

func (c *GlPopGroupMarkerEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPopGroupMarkerEXT(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPopGroupMarkerEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlPopGroupMarkerEXT) TypeID() atom.TypeID {
	return 19
}
func (c *GlPopGroupMarkerEXT) Flags() atom.Flags {
	return 0
}
func (GlPopGroupMarkerEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage1DEXT_In struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
}
type GlTexStorage1DEXT_Out struct {
	binary.Generate
}
type GlTexStorage1DEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlTexStorage1DEXT_In
	Out     GlTexStorage1DEXT_Out
}

func (c *GlTexStorage1DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexStorage1DEXT(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.In.Levels),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexStorage1DEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTexStorage1DEXT) TypeID() atom.TypeID {
	return 20
}
func (c *GlTexStorage1DEXT) Flags() atom.Flags {
	return 0
}
func (GlTexStorage1DEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage2DEXT_In struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
	Height int32
}
type GlTexStorage2DEXT_Out struct {
	binary.Generate
}
type GlTexStorage2DEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlTexStorage2DEXT_In
	Out     GlTexStorage2DEXT_Out
}

func (c *GlTexStorage2DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexStorage2DEXT(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.In.Levels),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexStorage2DEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTexStorage2DEXT) TypeID() atom.TypeID {
	return 21
}
func (c *GlTexStorage2DEXT) Flags() atom.Flags {
	return 0
}
func (GlTexStorage2DEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTexStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTexStorage3DEXT_In struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
	Height int32
	Depth  int32
}
type GlTexStorage3DEXT_Out struct {
	binary.Generate
}
type GlTexStorage3DEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlTexStorage3DEXT_In
	Out     GlTexStorage3DEXT_Out
}

func (c *GlTexStorage3DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexStorage3DEXT(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.In.Levels),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		fmt.Sprintf("depth:%v", c.In.Depth),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexStorage3DEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTexStorage3DEXT) TypeID() atom.TypeID {
	return 22
}
func (c *GlTexStorage3DEXT) Flags() atom.Flags {
	return 0
}
func (GlTexStorage3DEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage1DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage1DEXT_In struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
}
type GlTextureStorage1DEXT_Out struct {
	binary.Generate
}
type GlTextureStorage1DEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlTextureStorage1DEXT_In
	Out     GlTextureStorage1DEXT_Out
}

func (c *GlTextureStorage1DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTextureStorage1DEXT(",
		fmt.Sprintf("texture:%v", c.In.Texture),
		", ",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.In.Levels),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTextureStorage1DEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTextureStorage1DEXT) TypeID() atom.TypeID {
	return 23
}
func (c *GlTextureStorage1DEXT) Flags() atom.Flags {
	return 0
}
func (GlTextureStorage1DEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage2DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage2DEXT_In struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
	Height  int32
}
type GlTextureStorage2DEXT_Out struct {
	binary.Generate
}
type GlTextureStorage2DEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlTextureStorage2DEXT_In
	Out     GlTextureStorage2DEXT_Out
}

func (c *GlTextureStorage2DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTextureStorage2DEXT(",
		fmt.Sprintf("texture:%v", c.In.Texture),
		", ",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.In.Levels),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTextureStorage2DEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTextureStorage2DEXT) TypeID() atom.TypeID {
	return 24
}
func (c *GlTextureStorage2DEXT) Flags() atom.Flags {
	return 0
}
func (GlTextureStorage2DEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTextureStorage3DEXT
////////////////////////////////////////////////////////////////////////////////
type GlTextureStorage3DEXT_In struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
	Height  int32
	Depth   int32
}
type GlTextureStorage3DEXT_Out struct {
	binary.Generate
}
type GlTextureStorage3DEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlTextureStorage3DEXT_In
	Out     GlTextureStorage3DEXT_Out
}

func (c *GlTextureStorage3DEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTextureStorage3DEXT(",
		fmt.Sprintf("texture:%v", c.In.Texture),
		", ",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("levels:%v", c.In.Levels),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		fmt.Sprintf("depth:%v", c.In.Depth),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTextureStorage3DEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTextureStorage3DEXT) TypeID() atom.TypeID {
	return 25
}
func (c *GlTextureStorage3DEXT) Flags() atom.Flags {
	return 0
}
func (GlTextureStorage3DEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenVertexArraysOES
////////////////////////////////////////////////////////////////////////////////
type GlGenVertexArraysOES_In struct {
	binary.Generate
	Count int32
}
type GlGenVertexArraysOES_Out struct {
	binary.Generate
	Arrays VertexArrayIdArray
}
type GlGenVertexArraysOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenVertexArraysOES_In
	Out     GlGenVertexArraysOES_Out
}

func (c *GlGenVertexArraysOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenVertexArraysOES(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.Out.Arrays),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenVertexArraysOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenVertexArraysOES) TypeID() atom.TypeID {
	return 26
}
func (c *GlGenVertexArraysOES) Flags() atom.Flags {
	return 0
}
func (GlGenVertexArraysOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBindVertexArrayOES
////////////////////////////////////////////////////////////////////////////////
type GlBindVertexArrayOES_In struct {
	binary.Generate
	Array VertexArrayId
}
type GlBindVertexArrayOES_Out struct {
	binary.Generate
}
type GlBindVertexArrayOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlBindVertexArrayOES_In
	Out     GlBindVertexArrayOES_Out
}

func (c *GlBindVertexArrayOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindVertexArrayOES(",
		fmt.Sprintf("array:%v", c.In.Array),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindVertexArrayOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBindVertexArrayOES) TypeID() atom.TypeID {
	return 27
}
func (c *GlBindVertexArrayOES) Flags() atom.Flags {
	return 0
}
func (GlBindVertexArrayOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteVertexArraysOES
////////////////////////////////////////////////////////////////////////////////
type GlDeleteVertexArraysOES_In struct {
	binary.Generate
	Count  int32
	Arrays VertexArrayIdArray
}
type GlDeleteVertexArraysOES_Out struct {
	binary.Generate
}
type GlDeleteVertexArraysOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteVertexArraysOES_In
	Out     GlDeleteVertexArraysOES_Out
}

func (c *GlDeleteVertexArraysOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteVertexArraysOES(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Arrays),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteVertexArraysOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteVertexArraysOES) TypeID() atom.TypeID {
	return 28
}
func (c *GlDeleteVertexArraysOES) Flags() atom.Flags {
	return 0
}
func (GlDeleteVertexArraysOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsVertexArrayOES
////////////////////////////////////////////////////////////////////////////////
type GlIsVertexArrayOES_In struct {
	binary.Generate
	Array VertexArrayId
}
type GlIsVertexArrayOES_Out struct {
	binary.Generate
	Result bool
}
type GlIsVertexArrayOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsVertexArrayOES_In
	Out     GlIsVertexArrayOES_Out
}

func (c *GlIsVertexArrayOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsVertexArrayOES(",
		fmt.Sprintf("array:%v", c.In.Array),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsVertexArrayOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsVertexArrayOES) TypeID() atom.TypeID {
	return 29
}
func (c *GlIsVertexArrayOES) Flags() atom.Flags {
	return 0
}
func (GlIsVertexArrayOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetTexture2DOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetTexture2DOES_In struct {
	binary.Generate
	Target ImageTargetTexture
	Image  ImageOES
}
type GlEGLImageTargetTexture2DOES_Out struct {
	binary.Generate
}
type GlEGLImageTargetTexture2DOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlEGLImageTargetTexture2DOES_In
	Out     GlEGLImageTargetTexture2DOES_Out
}

func (c *GlEGLImageTargetTexture2DOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEGLImageTargetTexture2DOES(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("image:%v", c.In.Image),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEGLImageTargetTexture2DOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEGLImageTargetTexture2DOES) TypeID() atom.TypeID {
	return 30
}
func (c *GlEGLImageTargetTexture2DOES) Flags() atom.Flags {
	return 0
}
func (GlEGLImageTargetTexture2DOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEGLImageTargetRenderbufferStorageOES
////////////////////////////////////////////////////////////////////////////////
type GlEGLImageTargetRenderbufferStorageOES_In struct {
	binary.Generate
	Target ImageTargetRenderbufferStorage
	Image  TexturePointer
}
type GlEGLImageTargetRenderbufferStorageOES_Out struct {
	binary.Generate
}
type GlEGLImageTargetRenderbufferStorageOES struct {
	binary.Generate
	Context atom.ContextID
	In      GlEGLImageTargetRenderbufferStorageOES_In
	Out     GlEGLImageTargetRenderbufferStorageOES_Out
}

func (c *GlEGLImageTargetRenderbufferStorageOES) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEGLImageTargetRenderbufferStorageOES(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("image:%v", c.In.Image),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEGLImageTargetRenderbufferStorageOES) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEGLImageTargetRenderbufferStorageOES) TypeID() atom.TypeID {
	return 31
}
func (c *GlEGLImageTargetRenderbufferStorageOES) Flags() atom.Flags {
	return 0
}
func (GlEGLImageTargetRenderbufferStorageOES) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetGraphicsResetStatusEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetGraphicsResetStatusEXT_In struct {
	binary.Generate
}
type GlGetGraphicsResetStatusEXT_Out struct {
	binary.Generate
	Result ResetStatus
}
type GlGetGraphicsResetStatusEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetGraphicsResetStatusEXT_In
	Out     GlGetGraphicsResetStatusEXT_Out
}

func (c *GlGetGraphicsResetStatusEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetGraphicsResetStatusEXT(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlGetGraphicsResetStatusEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetGraphicsResetStatusEXT) TypeID() atom.TypeID {
	return 32
}
func (c *GlGetGraphicsResetStatusEXT) Flags() atom.Flags {
	return 0
}
func (GlGetGraphicsResetStatusEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBindAttribLocation
////////////////////////////////////////////////////////////////////////////////
type GlBindAttribLocation_In struct {
	binary.Generate
	Program  ProgramId
	Location AttributeLocation
	Name     string
}
type GlBindAttribLocation_Out struct {
	binary.Generate
}
type GlBindAttribLocation struct {
	binary.Generate
	Context atom.ContextID
	In      GlBindAttribLocation_In
	Out     GlBindAttribLocation_Out
}

func (c *GlBindAttribLocation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindAttribLocation(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("name:%v", c.In.Name),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindAttribLocation) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBindAttribLocation) TypeID() atom.TypeID {
	return 33
}
func (c *GlBindAttribLocation) Flags() atom.Flags {
	return 0
}
func (GlBindAttribLocation) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBlendFunc
////////////////////////////////////////////////////////////////////////////////
type GlBlendFunc_In struct {
	binary.Generate
	SrcFactor BlendFactor
	DstFactor BlendFactor
}
type GlBlendFunc_Out struct {
	binary.Generate
}
type GlBlendFunc struct {
	binary.Generate
	Context atom.ContextID
	In      GlBlendFunc_In
	Out     GlBlendFunc_Out
}

func (c *GlBlendFunc) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendFunc(",
		c.In.SrcFactor.String(),
		", ",
		c.In.DstFactor.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendFunc) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBlendFunc) TypeID() atom.TypeID {
	return 34
}
func (c *GlBlendFunc) Flags() atom.Flags {
	return 0
}
func (GlBlendFunc) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBlendFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendFuncSeparate_In struct {
	binary.Generate
	SrcFactorRgb   BlendFactor
	DstFactorRgb   BlendFactor
	SrcFactorAlpha BlendFactor
	DstFactorAlpha BlendFactor
}
type GlBlendFuncSeparate_Out struct {
	binary.Generate
}
type GlBlendFuncSeparate struct {
	binary.Generate
	Context atom.ContextID
	In      GlBlendFuncSeparate_In
	Out     GlBlendFuncSeparate_Out
}

func (c *GlBlendFuncSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendFuncSeparate(",
		c.In.SrcFactorRgb.String(),
		", ",
		c.In.DstFactorRgb.String(),
		", ",
		c.In.SrcFactorAlpha.String(),
		", ",
		c.In.DstFactorAlpha.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendFuncSeparate) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBlendFuncSeparate) TypeID() atom.TypeID {
	return 35
}
func (c *GlBlendFuncSeparate) Flags() atom.Flags {
	return 0
}
func (GlBlendFuncSeparate) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquation
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquation_In struct {
	binary.Generate
	Equation BlendEquation
}
type GlBlendEquation_Out struct {
	binary.Generate
}
type GlBlendEquation struct {
	binary.Generate
	Context atom.ContextID
	In      GlBlendEquation_In
	Out     GlBlendEquation_Out
}

func (c *GlBlendEquation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendEquation(",
		c.In.Equation.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendEquation) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBlendEquation) TypeID() atom.TypeID {
	return 36
}
func (c *GlBlendEquation) Flags() atom.Flags {
	return 0
}
func (GlBlendEquation) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBlendEquationSeparate
////////////////////////////////////////////////////////////////////////////////
type GlBlendEquationSeparate_In struct {
	binary.Generate
	Rgb   BlendEquation
	Alpha BlendEquation
}
type GlBlendEquationSeparate_Out struct {
	binary.Generate
}
type GlBlendEquationSeparate struct {
	binary.Generate
	Context atom.ContextID
	In      GlBlendEquationSeparate_In
	Out     GlBlendEquationSeparate_Out
}

func (c *GlBlendEquationSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendEquationSeparate(",
		c.In.Rgb.String(),
		", ",
		c.In.Alpha.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendEquationSeparate) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBlendEquationSeparate) TypeID() atom.TypeID {
	return 37
}
func (c *GlBlendEquationSeparate) Flags() atom.Flags {
	return 0
}
func (GlBlendEquationSeparate) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBlendColor
////////////////////////////////////////////////////////////////////////////////
type GlBlendColor_In struct {
	binary.Generate
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}
type GlBlendColor_Out struct {
	binary.Generate
}
type GlBlendColor struct {
	binary.Generate
	Context atom.ContextID
	In      GlBlendColor_In
	Out     GlBlendColor_Out
}

func (c *GlBlendColor) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlendColor(",
		fmt.Sprintf("red:%v", c.In.Red),
		", ",
		fmt.Sprintf("green:%v", c.In.Green),
		", ",
		fmt.Sprintf("blue:%v", c.In.Blue),
		", ",
		fmt.Sprintf("alpha:%v", c.In.Alpha),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlendColor) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBlendColor) TypeID() atom.TypeID {
	return 38
}
func (c *GlBlendColor) Flags() atom.Flags {
	return 0
}
func (GlBlendColor) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEnableVertexAttribArray
////////////////////////////////////////////////////////////////////////////////
type GlEnableVertexAttribArray_In struct {
	binary.Generate
	Location AttributeLocation
}
type GlEnableVertexAttribArray_Out struct {
	binary.Generate
}
type GlEnableVertexAttribArray struct {
	binary.Generate
	Context atom.ContextID
	In      GlEnableVertexAttribArray_In
	Out     GlEnableVertexAttribArray_Out
}

func (c *GlEnableVertexAttribArray) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEnableVertexAttribArray(",
		fmt.Sprintf("location:%v", c.In.Location),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEnableVertexAttribArray) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEnableVertexAttribArray) TypeID() atom.TypeID {
	return 39
}
func (c *GlEnableVertexAttribArray) Flags() atom.Flags {
	return 0
}
func (GlEnableVertexAttribArray) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDisableVertexAttribArray
////////////////////////////////////////////////////////////////////////////////
type GlDisableVertexAttribArray_In struct {
	binary.Generate
	Location AttributeLocation
}
type GlDisableVertexAttribArray_Out struct {
	binary.Generate
}
type GlDisableVertexAttribArray struct {
	binary.Generate
	Context atom.ContextID
	In      GlDisableVertexAttribArray_In
	Out     GlDisableVertexAttribArray_Out
}

func (c *GlDisableVertexAttribArray) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDisableVertexAttribArray(",
		fmt.Sprintf("location:%v", c.In.Location),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDisableVertexAttribArray) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDisableVertexAttribArray) TypeID() atom.TypeID {
	return 40
}
func (c *GlDisableVertexAttribArray) Flags() atom.Flags {
	return 0
}
func (GlDisableVertexAttribArray) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttribPointer
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttribPointer_In struct {
	binary.Generate
	Location   AttributeLocation
	Size       int32
	Type       VertexAttribType
	Normalized bool
	Stride     int32
	Data       VertexPointer
}
type GlVertexAttribPointer_Out struct {
	binary.Generate
}
type GlVertexAttribPointer struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttribPointer_In
	Out     GlVertexAttribPointer_Out
}

func (c *GlVertexAttribPointer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttribPointer(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("size:%v", c.In.Size),
		", ",
		c.In.Type.String(),
		", ",
		fmt.Sprintf("normalized:%v", c.In.Normalized),
		", ",
		fmt.Sprintf("stride:%v", c.In.Stride),
		", ",
		fmt.Sprintf("data:%v", c.In.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttribPointer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttribPointer) TypeID() atom.TypeID {
	return 41
}
func (c *GlVertexAttribPointer) Flags() atom.Flags {
	return 0
}
func (GlVertexAttribPointer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveAttrib
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveAttrib_In struct {
	binary.Generate
	Program    ProgramId
	Location   AttributeLocation
	BufferSize int32
}
type GlGetActiveAttrib_Out struct {
	binary.Generate
	BufferBytesWritten int32
	VectorCount        int32
	Type               ShaderAttribType
	Name               string
}
type GlGetActiveAttrib struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetActiveAttrib_In
	Out     GlGetActiveAttrib_Out
}

func (c *GlGetActiveAttrib) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetActiveAttrib(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("buffer_size:%v", c.In.BufferSize),
		", ",
		fmt.Sprintf("buffer_bytes_written:%v", c.Out.BufferBytesWritten),
		", ",
		fmt.Sprintf("vector_count:%v", c.Out.VectorCount),
		", ",
		c.Out.Type.String(),
		", ",
		fmt.Sprintf("name:%v", c.Out.Name),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetActiveAttrib) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetActiveAttrib) TypeID() atom.TypeID {
	return 42
}
func (c *GlGetActiveAttrib) Flags() atom.Flags {
	return 0
}
func (GlGetActiveAttrib) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetActiveUniform
////////////////////////////////////////////////////////////////////////////////
type GlGetActiveUniform_In struct {
	binary.Generate
	Program    ProgramId
	Location   int32
	BufferSize int32
}
type GlGetActiveUniform_Out struct {
	binary.Generate
	BufferBytesWritten int32
	Size               int32
	Type               ShaderUniformType
	Name               string
}
type GlGetActiveUniform struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetActiveUniform_In
	Out     GlGetActiveUniform_Out
}

func (c *GlGetActiveUniform) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetActiveUniform(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("buffer_size:%v", c.In.BufferSize),
		", ",
		fmt.Sprintf("buffer_bytes_written:%v", c.Out.BufferBytesWritten),
		", ",
		fmt.Sprintf("size:%v", c.Out.Size),
		", ",
		c.Out.Type.String(),
		", ",
		fmt.Sprintf("name:%v", c.Out.Name),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetActiveUniform) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetActiveUniform) TypeID() atom.TypeID {
	return 43
}
func (c *GlGetActiveUniform) Flags() atom.Flags {
	return 0
}
func (GlGetActiveUniform) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetError
////////////////////////////////////////////////////////////////////////////////
type GlGetError_In struct {
	binary.Generate
}
type GlGetError_Out struct {
	binary.Generate
	Result Error
}
type GlGetError struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetError_In
	Out     GlGetError_Out
}

func (c *GlGetError) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetError(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlGetError) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetError) TypeID() atom.TypeID {
	return 44
}
func (c *GlGetError) Flags() atom.Flags {
	return 0
}
func (GlGetError) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramiv
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramiv_In struct {
	binary.Generate
	Program   ProgramId
	Parameter ProgramParameter
}
type GlGetProgramiv_Out struct {
	binary.Generate
	Value S32Array
}
type GlGetProgramiv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetProgramiv_In
	Out     GlGetProgramiv_Out
}

func (c *GlGetProgramiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetProgramiv(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetProgramiv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetProgramiv) TypeID() atom.TypeID {
	return 45
}
func (c *GlGetProgramiv) Flags() atom.Flags {
	return 0
}
func (GlGetProgramiv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderiv
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderiv_In struct {
	binary.Generate
	Shader    ShaderId
	Parameter ShaderParameter
}
type GlGetShaderiv_Out struct {
	binary.Generate
	Value S32Array
}
type GlGetShaderiv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetShaderiv_In
	Out     GlGetShaderiv_Out
}

func (c *GlGetShaderiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderiv(",
		fmt.Sprintf("shader:%v", c.In.Shader),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderiv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetShaderiv) TypeID() atom.TypeID {
	return 46
}
func (c *GlGetShaderiv) Flags() atom.Flags {
	return 0
}
func (GlGetShaderiv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformLocation
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformLocation_In struct {
	binary.Generate
	Program ProgramId
	Name    string
}
type GlGetUniformLocation_Out struct {
	binary.Generate
	Result UniformLocation
}
type GlGetUniformLocation struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetUniformLocation_In
	Out     GlGetUniformLocation_Out
}

func (c *GlGetUniformLocation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetUniformLocation(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("name:%v", c.In.Name),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlGetUniformLocation) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetUniformLocation) TypeID() atom.TypeID {
	return 47
}
func (c *GlGetUniformLocation) Flags() atom.Flags {
	return 0
}
func (GlGetUniformLocation) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetAttribLocation
////////////////////////////////////////////////////////////////////////////////
type GlGetAttribLocation_In struct {
	binary.Generate
	Program ProgramId
	Name    string
}
type GlGetAttribLocation_Out struct {
	binary.Generate
	Result AttributeLocation
}
type GlGetAttribLocation struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetAttribLocation_In
	Out     GlGetAttribLocation_Out
}

func (c *GlGetAttribLocation) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetAttribLocation(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("name:%v", c.In.Name),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlGetAttribLocation) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetAttribLocation) TypeID() atom.TypeID {
	return 48
}
func (c *GlGetAttribLocation) Flags() atom.Flags {
	return 0
}
func (GlGetAttribLocation) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlPixelStorei
////////////////////////////////////////////////////////////////////////////////
type GlPixelStorei_In struct {
	binary.Generate
	Parameter PixelStoreParameter
	Value     int32
}
type GlPixelStorei_Out struct {
	binary.Generate
}
type GlPixelStorei struct {
	binary.Generate
	Context atom.ContextID
	In      GlPixelStorei_In
	Out     GlPixelStorei_Out
}

func (c *GlPixelStorei) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPixelStorei(",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPixelStorei) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlPixelStorei) TypeID() atom.TypeID {
	return 49
}
func (c *GlPixelStorei) Flags() atom.Flags {
	return 0
}
func (GlPixelStorei) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTexParameteri
////////////////////////////////////////////////////////////////////////////////
type GlTexParameteri_In struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Value     int32
}
type GlTexParameteri_Out struct {
	binary.Generate
}
type GlTexParameteri struct {
	binary.Generate
	Context atom.ContextID
	In      GlTexParameteri_In
	Out     GlTexParameteri_Out
}

func (c *GlTexParameteri) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexParameteri(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexParameteri) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTexParameteri) TypeID() atom.TypeID {
	return 50
}
func (c *GlTexParameteri) Flags() atom.Flags {
	return 0
}
func (GlTexParameteri) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTexParameterf
////////////////////////////////////////////////////////////////////////////////
type GlTexParameterf_In struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Value     float32
}
type GlTexParameterf_Out struct {
	binary.Generate
}
type GlTexParameterf struct {
	binary.Generate
	Context atom.ContextID
	In      GlTexParameterf_In
	Out     GlTexParameterf_Out
}

func (c *GlTexParameterf) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexParameterf(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexParameterf) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTexParameterf) TypeID() atom.TypeID {
	return 51
}
func (c *GlTexParameterf) Flags() atom.Flags {
	return 0
}
func (GlTexParameterf) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameteriv_In struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
}
type GlGetTexParameteriv_Out struct {
	binary.Generate
	Values S32Array
}
type GlGetTexParameteriv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetTexParameteriv_In
	Out     GlGetTexParameteriv_Out
}

func (c *GlGetTexParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetTexParameteriv(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetTexParameteriv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetTexParameteriv) TypeID() atom.TypeID {
	return 52
}
func (c *GlGetTexParameteriv) Flags() atom.Flags {
	return 0
}
func (GlGetTexParameteriv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetTexParameterfv
////////////////////////////////////////////////////////////////////////////////
type GlGetTexParameterfv_In struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
}
type GlGetTexParameterfv_Out struct {
	binary.Generate
	Values F32Array
}
type GlGetTexParameterfv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetTexParameterfv_In
	Out     GlGetTexParameterfv_Out
}

func (c *GlGetTexParameterfv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetTexParameterfv(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetTexParameterfv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetTexParameterfv) TypeID() atom.TypeID {
	return 53
}
func (c *GlGetTexParameterfv) Flags() atom.Flags {
	return 0
}
func (GlGetTexParameterfv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1i
////////////////////////////////////////////////////////////////////////////////
type GlUniform1i_In struct {
	binary.Generate
	Location UniformLocation
	Value    int32
}
type GlUniform1i_Out struct {
	binary.Generate
}
type GlUniform1i struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform1i_In
	Out     GlUniform1i_Out
}

func (c *GlUniform1i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1i(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value:%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1i) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform1i) TypeID() atom.TypeID {
	return 54
}
func (c *GlUniform1i) Flags() atom.Flags {
	return 0
}
func (GlUniform1i) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2i
////////////////////////////////////////////////////////////////////////////////
type GlUniform2i_In struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
}
type GlUniform2i_Out struct {
	binary.Generate
}
type GlUniform2i struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform2i_In
	Out     GlUniform2i_Out
}

func (c *GlUniform2i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2i(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2i) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform2i) TypeID() atom.TypeID {
	return 55
}
func (c *GlUniform2i) Flags() atom.Flags {
	return 0
}
func (GlUniform2i) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3i
////////////////////////////////////////////////////////////////////////////////
type GlUniform3i_In struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
	Value2   int32
}
type GlUniform3i_Out struct {
	binary.Generate
}
type GlUniform3i struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform3i_In
	Out     GlUniform3i_Out
}

func (c *GlUniform3i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3i(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.In.Value2),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3i) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform3i) TypeID() atom.TypeID {
	return 56
}
func (c *GlUniform3i) Flags() atom.Flags {
	return 0
}
func (GlUniform3i) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4i
////////////////////////////////////////////////////////////////////////////////
type GlUniform4i_In struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
	Value2   int32
	Value3   int32
}
type GlUniform4i_Out struct {
	binary.Generate
}
type GlUniform4i struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform4i_In
	Out     GlUniform4i_Out
}

func (c *GlUniform4i) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4i(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.In.Value2),
		", ",
		fmt.Sprintf("value3:%v", c.In.Value3),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4i) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform4i) TypeID() atom.TypeID {
	return 57
}
func (c *GlUniform4i) Flags() atom.Flags {
	return 0
}
func (GlUniform4i) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform1iv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
type GlUniform1iv_Out struct {
	binary.Generate
}
type GlUniform1iv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform1iv_In
	Out     GlUniform1iv_Out
}

func (c *GlUniform1iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1iv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1iv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform1iv) TypeID() atom.TypeID {
	return 58
}
func (c *GlUniform1iv) Flags() atom.Flags {
	return 0
}
func (GlUniform1iv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform2iv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
type GlUniform2iv_Out struct {
	binary.Generate
}
type GlUniform2iv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform2iv_In
	Out     GlUniform2iv_Out
}

func (c *GlUniform2iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2iv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2iv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform2iv) TypeID() atom.TypeID {
	return 59
}
func (c *GlUniform2iv) Flags() atom.Flags {
	return 0
}
func (GlUniform2iv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform3iv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
type GlUniform3iv_Out struct {
	binary.Generate
}
type GlUniform3iv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform3iv_In
	Out     GlUniform3iv_Out
}

func (c *GlUniform3iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3iv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3iv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform3iv) TypeID() atom.TypeID {
	return 60
}
func (c *GlUniform3iv) Flags() atom.Flags {
	return 0
}
func (GlUniform3iv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4iv
////////////////////////////////////////////////////////////////////////////////
type GlUniform4iv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
type GlUniform4iv_Out struct {
	binary.Generate
}
type GlUniform4iv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform4iv_In
	Out     GlUniform4iv_Out
}

func (c *GlUniform4iv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4iv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4iv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform4iv) TypeID() atom.TypeID {
	return 61
}
func (c *GlUniform4iv) Flags() atom.Flags {
	return 0
}
func (GlUniform4iv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1f
////////////////////////////////////////////////////////////////////////////////
type GlUniform1f_In struct {
	binary.Generate
	Location UniformLocation
	Value    float32
}
type GlUniform1f_Out struct {
	binary.Generate
}
type GlUniform1f struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform1f_In
	Out     GlUniform1f_Out
}

func (c *GlUniform1f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value:%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform1f) TypeID() atom.TypeID {
	return 62
}
func (c *GlUniform1f) Flags() atom.Flags {
	return 0
}
func (GlUniform1f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2f
////////////////////////////////////////////////////////////////////////////////
type GlUniform2f_In struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
}
type GlUniform2f_Out struct {
	binary.Generate
}
type GlUniform2f struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform2f_In
	Out     GlUniform2f_Out
}

func (c *GlUniform2f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform2f) TypeID() atom.TypeID {
	return 63
}
func (c *GlUniform2f) Flags() atom.Flags {
	return 0
}
func (GlUniform2f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3f
////////////////////////////////////////////////////////////////////////////////
type GlUniform3f_In struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
	Value2   float32
}
type GlUniform3f_Out struct {
	binary.Generate
}
type GlUniform3f struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform3f_In
	Out     GlUniform3f_Out
}

func (c *GlUniform3f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.In.Value2),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform3f) TypeID() atom.TypeID {
	return 64
}
func (c *GlUniform3f) Flags() atom.Flags {
	return 0
}
func (GlUniform3f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4f
////////////////////////////////////////////////////////////////////////////////
type GlUniform4f_In struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
	Value2   float32
	Value3   float32
}
type GlUniform4f_Out struct {
	binary.Generate
}
type GlUniform4f struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform4f_In
	Out     GlUniform4f_Out
}

func (c *GlUniform4f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.In.Value2),
		", ",
		fmt.Sprintf("value3:%v", c.In.Value3),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform4f) TypeID() atom.TypeID {
	return 65
}
func (c *GlUniform4f) Flags() atom.Flags {
	return 0
}
func (GlUniform4f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform1fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform1fv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
type GlUniform1fv_Out struct {
	binary.Generate
}
type GlUniform1fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform1fv_In
	Out     GlUniform1fv_Out
}

func (c *GlUniform1fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform1fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform1fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform1fv) TypeID() atom.TypeID {
	return 66
}
func (c *GlUniform1fv) Flags() atom.Flags {
	return 0
}
func (GlUniform1fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform2fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform2fv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
type GlUniform2fv_Out struct {
	binary.Generate
}
type GlUniform2fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform2fv_In
	Out     GlUniform2fv_Out
}

func (c *GlUniform2fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform2fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform2fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform2fv) TypeID() atom.TypeID {
	return 67
}
func (c *GlUniform2fv) Flags() atom.Flags {
	return 0
}
func (GlUniform2fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform3fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform3fv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
type GlUniform3fv_Out struct {
	binary.Generate
}
type GlUniform3fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform3fv_In
	Out     GlUniform3fv_Out
}

func (c *GlUniform3fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform3fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform3fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform3fv) TypeID() atom.TypeID {
	return 68
}
func (c *GlUniform3fv) Flags() atom.Flags {
	return 0
}
func (GlUniform3fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniform4fv
////////////////////////////////////////////////////////////////////////////////
type GlUniform4fv_In struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
type GlUniform4fv_Out struct {
	binary.Generate
}
type GlUniform4fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniform4fv_In
	Out     GlUniform4fv_Out
}

func (c *GlUniform4fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniform4fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniform4fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniform4fv) TypeID() atom.TypeID {
	return 69
}
func (c *GlUniform4fv) Flags() atom.Flags {
	return 0
}
func (GlUniform4fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix2fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix2fv_In struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}
type GlUniformMatrix2fv_Out struct {
	binary.Generate
}
type GlUniformMatrix2fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniformMatrix2fv_In
	Out     GlUniformMatrix2fv_Out
}

func (c *GlUniformMatrix2fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniformMatrix2fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("transpose:%v", c.In.Transpose),
		", ",
		fmt.Sprintf("%v", c.In.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniformMatrix2fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniformMatrix2fv) TypeID() atom.TypeID {
	return 70
}
func (c *GlUniformMatrix2fv) Flags() atom.Flags {
	return 0
}
func (GlUniformMatrix2fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix3fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix3fv_In struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}
type GlUniformMatrix3fv_Out struct {
	binary.Generate
}
type GlUniformMatrix3fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniformMatrix3fv_In
	Out     GlUniformMatrix3fv_Out
}

func (c *GlUniformMatrix3fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniformMatrix3fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("transpose:%v", c.In.Transpose),
		", ",
		fmt.Sprintf("%v", c.In.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniformMatrix3fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniformMatrix3fv) TypeID() atom.TypeID {
	return 71
}
func (c *GlUniformMatrix3fv) Flags() atom.Flags {
	return 0
}
func (GlUniformMatrix3fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUniformMatrix4fv
////////////////////////////////////////////////////////////////////////////////
type GlUniformMatrix4fv_In struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}
type GlUniformMatrix4fv_Out struct {
	binary.Generate
}
type GlUniformMatrix4fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlUniformMatrix4fv_In
	Out     GlUniformMatrix4fv_Out
}

func (c *GlUniformMatrix4fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUniformMatrix4fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("transpose:%v", c.In.Transpose),
		", ",
		fmt.Sprintf("%v", c.In.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUniformMatrix4fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUniformMatrix4fv) TypeID() atom.TypeID {
	return 72
}
func (c *GlUniformMatrix4fv) Flags() atom.Flags {
	return 0
}
func (GlUniformMatrix4fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformfv
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformfv_In struct {
	binary.Generate
	Program  ProgramId
	Location UniformLocation
	Values   F32Array
}
type GlGetUniformfv_Out struct {
	binary.Generate
}
type GlGetUniformfv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetUniformfv_In
	Out     GlGetUniformfv_Out
}

func (c *GlGetUniformfv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetUniformfv(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("%v", c.In.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetUniformfv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetUniformfv) TypeID() atom.TypeID {
	return 73
}
func (c *GlGetUniformfv) Flags() atom.Flags {
	return 0
}
func (GlGetUniformfv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetUniformiv
////////////////////////////////////////////////////////////////////////////////
type GlGetUniformiv_In struct {
	binary.Generate
	Program  ProgramId
	Location UniformLocation
	Values   S32Array
}
type GlGetUniformiv_Out struct {
	binary.Generate
}
type GlGetUniformiv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetUniformiv_In
	Out     GlGetUniformiv_Out
}

func (c *GlGetUniformiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetUniformiv(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("%v", c.In.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetUniformiv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetUniformiv) TypeID() atom.TypeID {
	return 74
}
func (c *GlGetUniformiv) Flags() atom.Flags {
	return 0
}
func (GlGetUniformiv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib1f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib1f_In struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
}
type GlVertexAttrib1f_Out struct {
	binary.Generate
}
type GlVertexAttrib1f struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib1f_In
	Out     GlVertexAttrib1f_Out
}

func (c *GlVertexAttrib1f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib1f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib1f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib1f) TypeID() atom.TypeID {
	return 75
}
func (c *GlVertexAttrib1f) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib1f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib2f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib2f_In struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
}
type GlVertexAttrib2f_Out struct {
	binary.Generate
}
type GlVertexAttrib2f struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib2f_In
	Out     GlVertexAttrib2f_Out
}

func (c *GlVertexAttrib2f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib2f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib2f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib2f) TypeID() atom.TypeID {
	return 76
}
func (c *GlVertexAttrib2f) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib2f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib3f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib3f_In struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
	Value2   float32
}
type GlVertexAttrib3f_Out struct {
	binary.Generate
}
type GlVertexAttrib3f struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib3f_In
	Out     GlVertexAttrib3f_Out
}

func (c *GlVertexAttrib3f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib3f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.In.Value2),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib3f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib3f) TypeID() atom.TypeID {
	return 77
}
func (c *GlVertexAttrib3f) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib3f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib4f
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib4f_In struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
	Value2   float32
	Value3   float32
}
type GlVertexAttrib4f_Out struct {
	binary.Generate
}
type GlVertexAttrib4f struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib4f_In
	Out     GlVertexAttrib4f_Out
}

func (c *GlVertexAttrib4f) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib4f(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("value0:%v", c.In.Value0),
		", ",
		fmt.Sprintf("value1:%v", c.In.Value1),
		", ",
		fmt.Sprintf("value2:%v", c.In.Value2),
		", ",
		fmt.Sprintf("value3:%v", c.In.Value3),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib4f) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib4f) TypeID() atom.TypeID {
	return 78
}
func (c *GlVertexAttrib4f) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib4f) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib1fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib1fv_In struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
type GlVertexAttrib1fv_Out struct {
	binary.Generate
}
type GlVertexAttrib1fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib1fv_In
	Out     GlVertexAttrib1fv_Out
}

func (c *GlVertexAttrib1fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib1fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib1fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib1fv) TypeID() atom.TypeID {
	return 79
}
func (c *GlVertexAttrib1fv) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib1fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib2fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib2fv_In struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
type GlVertexAttrib2fv_Out struct {
	binary.Generate
}
type GlVertexAttrib2fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib2fv_In
	Out     GlVertexAttrib2fv_Out
}

func (c *GlVertexAttrib2fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib2fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib2fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib2fv) TypeID() atom.TypeID {
	return 80
}
func (c *GlVertexAttrib2fv) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib2fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib3fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib3fv_In struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
type GlVertexAttrib3fv_Out struct {
	binary.Generate
}
type GlVertexAttrib3fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib3fv_In
	Out     GlVertexAttrib3fv_Out
}

func (c *GlVertexAttrib3fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib3fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib3fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib3fv) TypeID() atom.TypeID {
	return 81
}
func (c *GlVertexAttrib3fv) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib3fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlVertexAttrib4fv
////////////////////////////////////////////////////////////////////////////////
type GlVertexAttrib4fv_In struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
type GlVertexAttrib4fv_Out struct {
	binary.Generate
}
type GlVertexAttrib4fv struct {
	binary.Generate
	Context atom.ContextID
	In      GlVertexAttrib4fv_In
	Out     GlVertexAttrib4fv_Out
}

func (c *GlVertexAttrib4fv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glVertexAttrib4fv(",
		fmt.Sprintf("location:%v", c.In.Location),
		", ",
		fmt.Sprintf("%v", c.In.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlVertexAttrib4fv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlVertexAttrib4fv) TypeID() atom.TypeID {
	return 82
}
func (c *GlVertexAttrib4fv) Flags() atom.Flags {
	return 0
}
func (GlVertexAttrib4fv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderPrecisionFormat
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderPrecisionFormat_In struct {
	binary.Generate
	ShaderType    ShaderType
	PrecisionType PrecisionType
}
type GlGetShaderPrecisionFormat_Out struct {
	binary.Generate
	Range     S32Array
	Precision int32
}
type GlGetShaderPrecisionFormat struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetShaderPrecisionFormat_In
	Out     GlGetShaderPrecisionFormat_Out
}

func (c *GlGetShaderPrecisionFormat) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderPrecisionFormat(",
		c.In.ShaderType.String(),
		", ",
		c.In.PrecisionType.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Range),
		", ",
		fmt.Sprintf("precision:%v", c.Out.Precision),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderPrecisionFormat) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetShaderPrecisionFormat) TypeID() atom.TypeID {
	return 83
}
func (c *GlGetShaderPrecisionFormat) Flags() atom.Flags {
	return 0
}
func (GlGetShaderPrecisionFormat) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDepthMask
////////////////////////////////////////////////////////////////////////////////
type GlDepthMask_In struct {
	binary.Generate
	Enabled bool
}
type GlDepthMask_Out struct {
	binary.Generate
}
type GlDepthMask struct {
	binary.Generate
	Context atom.ContextID
	In      GlDepthMask_In
	Out     GlDepthMask_Out
}

func (c *GlDepthMask) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDepthMask(",
		fmt.Sprintf("enabled:%v", c.In.Enabled),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDepthMask) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDepthMask) TypeID() atom.TypeID {
	return 84
}
func (c *GlDepthMask) Flags() atom.Flags {
	return 0
}
func (GlDepthMask) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDepthFunc
////////////////////////////////////////////////////////////////////////////////
type GlDepthFunc_In struct {
	binary.Generate
	Function TestFunction
}
type GlDepthFunc_Out struct {
	binary.Generate
}
type GlDepthFunc struct {
	binary.Generate
	Context atom.ContextID
	In      GlDepthFunc_In
	Out     GlDepthFunc_Out
}

func (c *GlDepthFunc) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDepthFunc(",
		c.In.Function.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDepthFunc) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDepthFunc) TypeID() atom.TypeID {
	return 85
}
func (c *GlDepthFunc) Flags() atom.Flags {
	return 0
}
func (GlDepthFunc) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDepthRangef
////////////////////////////////////////////////////////////////////////////////
type GlDepthRangef_In struct {
	binary.Generate
	Near float32
	Far  float32
}
type GlDepthRangef_Out struct {
	binary.Generate
}
type GlDepthRangef struct {
	binary.Generate
	Context atom.ContextID
	In      GlDepthRangef_In
	Out     GlDepthRangef_Out
}

func (c *GlDepthRangef) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDepthRangef(",
		fmt.Sprintf("near:%v", c.In.Near),
		", ",
		fmt.Sprintf("far:%v", c.In.Far),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDepthRangef) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDepthRangef) TypeID() atom.TypeID {
	return 86
}
func (c *GlDepthRangef) Flags() atom.Flags {
	return 0
}
func (GlDepthRangef) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlColorMask
////////////////////////////////////////////////////////////////////////////////
type GlColorMask_In struct {
	binary.Generate
	Red   bool
	Green bool
	Blue  bool
	Alpha bool
}
type GlColorMask_Out struct {
	binary.Generate
}
type GlColorMask struct {
	binary.Generate
	Context atom.ContextID
	In      GlColorMask_In
	Out     GlColorMask_Out
}

func (c *GlColorMask) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glColorMask(",
		fmt.Sprintf("red:%v", c.In.Red),
		", ",
		fmt.Sprintf("green:%v", c.In.Green),
		", ",
		fmt.Sprintf("blue:%v", c.In.Blue),
		", ",
		fmt.Sprintf("alpha:%v", c.In.Alpha),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlColorMask) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlColorMask) TypeID() atom.TypeID {
	return 87
}
func (c *GlColorMask) Flags() atom.Flags {
	return 0
}
func (GlColorMask) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlStencilMask
////////////////////////////////////////////////////////////////////////////////
type GlStencilMask_In struct {
	binary.Generate
	Mask uint32
}
type GlStencilMask_Out struct {
	binary.Generate
}
type GlStencilMask struct {
	binary.Generate
	Context atom.ContextID
	In      GlStencilMask_In
	Out     GlStencilMask_Out
}

func (c *GlStencilMask) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilMask(",
		fmt.Sprintf("mask:%v", c.In.Mask),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilMask) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlStencilMask) TypeID() atom.TypeID {
	return 88
}
func (c *GlStencilMask) Flags() atom.Flags {
	return 0
}
func (GlStencilMask) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlStencilMaskSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilMaskSeparate_In struct {
	binary.Generate
	Face FaceMode
	Mask uint32
}
type GlStencilMaskSeparate_Out struct {
	binary.Generate
}
type GlStencilMaskSeparate struct {
	binary.Generate
	Context atom.ContextID
	In      GlStencilMaskSeparate_In
	Out     GlStencilMaskSeparate_Out
}

func (c *GlStencilMaskSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilMaskSeparate(",
		c.In.Face.String(),
		", ",
		fmt.Sprintf("mask:%v", c.In.Mask),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilMaskSeparate) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlStencilMaskSeparate) TypeID() atom.TypeID {
	return 89
}
func (c *GlStencilMaskSeparate) Flags() atom.Flags {
	return 0
}
func (GlStencilMaskSeparate) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlStencilFuncSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilFuncSeparate_In struct {
	binary.Generate
	Face           FaceMode
	Function       TestFunction
	ReferenceValue int32
	Mask           int32
}
type GlStencilFuncSeparate_Out struct {
	binary.Generate
}
type GlStencilFuncSeparate struct {
	binary.Generate
	Context atom.ContextID
	In      GlStencilFuncSeparate_In
	Out     GlStencilFuncSeparate_Out
}

func (c *GlStencilFuncSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilFuncSeparate(",
		c.In.Face.String(),
		", ",
		c.In.Function.String(),
		", ",
		fmt.Sprintf("reference_value:%v", c.In.ReferenceValue),
		", ",
		fmt.Sprintf("mask:%v", c.In.Mask),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilFuncSeparate) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlStencilFuncSeparate) TypeID() atom.TypeID {
	return 90
}
func (c *GlStencilFuncSeparate) Flags() atom.Flags {
	return 0
}
func (GlStencilFuncSeparate) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlStencilOpSeparate
////////////////////////////////////////////////////////////////////////////////
type GlStencilOpSeparate_In struct {
	binary.Generate
	Face                 FaceMode
	StencilFail          StencilAction
	StencilPassDepthFail StencilAction
	StencilPassDepthPass StencilAction
}
type GlStencilOpSeparate_Out struct {
	binary.Generate
}
type GlStencilOpSeparate struct {
	binary.Generate
	Context atom.ContextID
	In      GlStencilOpSeparate_In
	Out     GlStencilOpSeparate_Out
}

func (c *GlStencilOpSeparate) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glStencilOpSeparate(",
		c.In.Face.String(),
		", ",
		c.In.StencilFail.String(),
		", ",
		c.In.StencilPassDepthFail.String(),
		", ",
		c.In.StencilPassDepthPass.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlStencilOpSeparate) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlStencilOpSeparate) TypeID() atom.TypeID {
	return 91
}
func (c *GlStencilOpSeparate) Flags() atom.Flags {
	return 0
}
func (GlStencilOpSeparate) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlFrontFace
////////////////////////////////////////////////////////////////////////////////
type GlFrontFace_In struct {
	binary.Generate
	Orientation FaceOrientation
}
type GlFrontFace_Out struct {
	binary.Generate
}
type GlFrontFace struct {
	binary.Generate
	Context atom.ContextID
	In      GlFrontFace_In
	Out     GlFrontFace_Out
}

func (c *GlFrontFace) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFrontFace(",
		c.In.Orientation.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFrontFace) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlFrontFace) TypeID() atom.TypeID {
	return 92
}
func (c *GlFrontFace) Flags() atom.Flags {
	return 0
}
func (GlFrontFace) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlViewport
////////////////////////////////////////////////////////////////////////////////
type GlViewport_In struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
}
type GlViewport_Out struct {
	binary.Generate
}
type GlViewport struct {
	binary.Generate
	Context atom.ContextID
	In      GlViewport_In
	Out     GlViewport_Out
}

func (c *GlViewport) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glViewport(",
		fmt.Sprintf("x:%v", c.In.X),
		", ",
		fmt.Sprintf("y:%v", c.In.Y),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlViewport) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlViewport) TypeID() atom.TypeID {
	return 93
}
func (c *GlViewport) Flags() atom.Flags {
	return 0
}
func (GlViewport) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlScissor
////////////////////////////////////////////////////////////////////////////////
type GlScissor_In struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
}
type GlScissor_Out struct {
	binary.Generate
}
type GlScissor struct {
	binary.Generate
	Context atom.ContextID
	In      GlScissor_In
	Out     GlScissor_Out
}

func (c *GlScissor) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glScissor(",
		fmt.Sprintf("x:%v", c.In.X),
		", ",
		fmt.Sprintf("y:%v", c.In.Y),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlScissor) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlScissor) TypeID() atom.TypeID {
	return 94
}
func (c *GlScissor) Flags() atom.Flags {
	return 0
}
func (GlScissor) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlActiveTexture
////////////////////////////////////////////////////////////////////////////////
type GlActiveTexture_In struct {
	binary.Generate
	Unit TextureUnit
}
type GlActiveTexture_Out struct {
	binary.Generate
}
type GlActiveTexture struct {
	binary.Generate
	Context atom.ContextID
	In      GlActiveTexture_In
	Out     GlActiveTexture_Out
}

func (c *GlActiveTexture) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glActiveTexture(",
		c.In.Unit.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlActiveTexture) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlActiveTexture) TypeID() atom.TypeID {
	return 95
}
func (c *GlActiveTexture) Flags() atom.Flags {
	return 0
}
func (GlActiveTexture) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenTextures
////////////////////////////////////////////////////////////////////////////////
type GlGenTextures_In struct {
	binary.Generate
	Count int32
}
type GlGenTextures_Out struct {
	binary.Generate
	Textures TextureIdArray
}
type GlGenTextures struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenTextures_In
	Out     GlGenTextures_Out
}

func (c *GlGenTextures) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenTextures(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.Out.Textures),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenTextures) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenTextures) TypeID() atom.TypeID {
	return 96
}
func (c *GlGenTextures) Flags() atom.Flags {
	return 0
}
func (GlGenTextures) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteTextures
////////////////////////////////////////////////////////////////////////////////
type GlDeleteTextures_In struct {
	binary.Generate
	Count    int32
	Textures TextureIdArray
}
type GlDeleteTextures_Out struct {
	binary.Generate
}
type GlDeleteTextures struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteTextures_In
	Out     GlDeleteTextures_Out
}

func (c *GlDeleteTextures) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteTextures(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Textures),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteTextures) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteTextures) TypeID() atom.TypeID {
	return 97
}
func (c *GlDeleteTextures) Flags() atom.Flags {
	return 0
}
func (GlDeleteTextures) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsTexture
////////////////////////////////////////////////////////////////////////////////
type GlIsTexture_In struct {
	binary.Generate
	Texture TextureId
}
type GlIsTexture_Out struct {
	binary.Generate
	Result bool
}
type GlIsTexture struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsTexture_In
	Out     GlIsTexture_Out
}

func (c *GlIsTexture) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsTexture(",
		fmt.Sprintf("texture:%v", c.In.Texture),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsTexture) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsTexture) TypeID() atom.TypeID {
	return 98
}
func (c *GlIsTexture) Flags() atom.Flags {
	return 0
}
func (GlIsTexture) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBindTexture
////////////////////////////////////////////////////////////////////////////////
type GlBindTexture_In struct {
	binary.Generate
	Target  TextureTarget
	Texture TextureId
}
type GlBindTexture_Out struct {
	binary.Generate
}
type GlBindTexture struct {
	binary.Generate
	Context atom.ContextID
	In      GlBindTexture_In
	Out     GlBindTexture_Out
}

func (c *GlBindTexture) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindTexture(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("texture:%v", c.In.Texture),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindTexture) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBindTexture) TypeID() atom.TypeID {
	return 99
}
func (c *GlBindTexture) Flags() atom.Flags {
	return 0
}
func (GlBindTexture) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexImage2D_In struct {
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
type GlTexImage2D_Out struct {
	binary.Generate
}
type GlTexImage2D struct {
	binary.Generate
	Context atom.ContextID
	In      GlTexImage2D_In
	Out     GlTexImage2D_Out
}

func (c *GlTexImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexImage2D(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.In.Level),
		", ",
		c.In.InternalFormat.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		fmt.Sprintf("border:%v", c.In.Border),
		", ",
		c.In.Format.String(),
		", ",
		c.In.Type.String(),
		", ",
		fmt.Sprintf("data:%v", c.In.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexImage2D) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTexImage2D) TypeID() atom.TypeID {
	return 100
}
func (c *GlTexImage2D) Flags() atom.Flags {
	return 0
}
func (GlTexImage2D) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlTexSubImage2D_In struct {
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
type GlTexSubImage2D_Out struct {
	binary.Generate
}
type GlTexSubImage2D struct {
	binary.Generate
	Context atom.ContextID
	In      GlTexSubImage2D_In
	Out     GlTexSubImage2D_Out
}

func (c *GlTexSubImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glTexSubImage2D(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.In.Level),
		", ",
		fmt.Sprintf("xoffset:%v", c.In.Xoffset),
		", ",
		fmt.Sprintf("yoffset:%v", c.In.Yoffset),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		c.In.Format.String(),
		", ",
		c.In.Type.String(),
		", ",
		fmt.Sprintf("data:%v", c.In.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlTexSubImage2D) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlTexSubImage2D) TypeID() atom.TypeID {
	return 101
}
func (c *GlTexSubImage2D) Flags() atom.Flags {
	return 0
}
func (GlTexSubImage2D) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexImage2D_In struct {
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
type GlCopyTexImage2D_Out struct {
	binary.Generate
}
type GlCopyTexImage2D struct {
	binary.Generate
	Context atom.ContextID
	In      GlCopyTexImage2D_In
	Out     GlCopyTexImage2D_Out
}

func (c *GlCopyTexImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCopyTexImage2D(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.In.Level),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("x:%v", c.In.X),
		", ",
		fmt.Sprintf("y:%v", c.In.Y),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		fmt.Sprintf("border:%v", c.In.Border),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCopyTexImage2D) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCopyTexImage2D) TypeID() atom.TypeID {
	return 102
}
func (c *GlCopyTexImage2D) Flags() atom.Flags {
	return 0
}
func (GlCopyTexImage2D) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCopyTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCopyTexSubImage2D_In struct {
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
type GlCopyTexSubImage2D_Out struct {
	binary.Generate
}
type GlCopyTexSubImage2D struct {
	binary.Generate
	Context atom.ContextID
	In      GlCopyTexSubImage2D_In
	Out     GlCopyTexSubImage2D_Out
}

func (c *GlCopyTexSubImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCopyTexSubImage2D(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.In.Level),
		", ",
		fmt.Sprintf("xoffset:%v", c.In.Xoffset),
		", ",
		fmt.Sprintf("yoffset:%v", c.In.Yoffset),
		", ",
		fmt.Sprintf("x:%v", c.In.X),
		", ",
		fmt.Sprintf("y:%v", c.In.Y),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCopyTexSubImage2D) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCopyTexSubImage2D) TypeID() atom.TypeID {
	return 103
}
func (c *GlCopyTexSubImage2D) Flags() atom.Flags {
	return 0
}
func (GlCopyTexSubImage2D) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexImage2D_In struct {
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
type GlCompressedTexImage2D_Out struct {
	binary.Generate
}
type GlCompressedTexImage2D struct {
	binary.Generate
	Context atom.ContextID
	In      GlCompressedTexImage2D_In
	Out     GlCompressedTexImage2D_Out
}

func (c *GlCompressedTexImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCompressedTexImage2D(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.In.Level),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		fmt.Sprintf("border:%v", c.In.Border),
		", ",
		fmt.Sprintf("image_size:%v", c.In.ImageSize),
		", ",
		fmt.Sprintf("data:%v", c.In.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCompressedTexImage2D) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCompressedTexImage2D) TypeID() atom.TypeID {
	return 104
}
func (c *GlCompressedTexImage2D) Flags() atom.Flags {
	return 0
}
func (GlCompressedTexImage2D) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCompressedTexSubImage2D
////////////////////////////////////////////////////////////////////////////////
type GlCompressedTexSubImage2D_In struct {
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
type GlCompressedTexSubImage2D_Out struct {
	binary.Generate
}
type GlCompressedTexSubImage2D struct {
	binary.Generate
	Context atom.ContextID
	In      GlCompressedTexSubImage2D_In
	Out     GlCompressedTexSubImage2D_Out
}

func (c *GlCompressedTexSubImage2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCompressedTexSubImage2D(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("level:%v", c.In.Level),
		", ",
		fmt.Sprintf("xoffset:%v", c.In.Xoffset),
		", ",
		fmt.Sprintf("yoffset:%v", c.In.Yoffset),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("image_size:%v", c.In.ImageSize),
		", ",
		fmt.Sprintf("data:%v", c.In.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCompressedTexSubImage2D) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCompressedTexSubImage2D) TypeID() atom.TypeID {
	return 105
}
func (c *GlCompressedTexSubImage2D) Flags() atom.Flags {
	return 0
}
func (GlCompressedTexSubImage2D) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenerateMipmap
////////////////////////////////////////////////////////////////////////////////
type GlGenerateMipmap_In struct {
	binary.Generate
	Target TextureImageTarget
}
type GlGenerateMipmap_Out struct {
	binary.Generate
}
type GlGenerateMipmap struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenerateMipmap_In
	Out     GlGenerateMipmap_Out
}

func (c *GlGenerateMipmap) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenerateMipmap(",
		c.In.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenerateMipmap) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenerateMipmap) TypeID() atom.TypeID {
	return 106
}
func (c *GlGenerateMipmap) Flags() atom.Flags {
	return 0
}
func (GlGenerateMipmap) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlReadPixels
////////////////////////////////////////////////////////////////////////////////
type GlReadPixels_In struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
	Format BaseTexelFormat
	Type   TexelType
}
type GlReadPixels_Out struct {
	binary.Generate
	Data memory.Pointer
}
type GlReadPixels struct {
	binary.Generate
	Context atom.ContextID
	In      GlReadPixels_In
	Out     GlReadPixels_Out
}

func (c *GlReadPixels) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glReadPixels(",
		fmt.Sprintf("x:%v", c.In.X),
		", ",
		fmt.Sprintf("y:%v", c.In.Y),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
		", ",
		c.In.Format.String(),
		", ",
		c.In.Type.String(),
		", ",
		fmt.Sprintf("0x%x", c.Out.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlReadPixels) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlReadPixels) TypeID() atom.TypeID {
	return 107
}
func (c *GlReadPixels) Flags() atom.Flags {
	return 0
}
func (GlReadPixels) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenFramebuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenFramebuffers_In struct {
	binary.Generate
	Count int32
}
type GlGenFramebuffers_Out struct {
	binary.Generate
	Framebuffers FramebufferIdArray
}
type GlGenFramebuffers struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenFramebuffers_In
	Out     GlGenFramebuffers_Out
}

func (c *GlGenFramebuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenFramebuffers(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.Out.Framebuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenFramebuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenFramebuffers) TypeID() atom.TypeID {
	return 108
}
func (c *GlGenFramebuffers) Flags() atom.Flags {
	return 0
}
func (GlGenFramebuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBindFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindFramebuffer_In struct {
	binary.Generate
	Target      FramebufferTarget
	Framebuffer FramebufferId
}
type GlBindFramebuffer_Out struct {
	binary.Generate
}
type GlBindFramebuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlBindFramebuffer_In
	Out     GlBindFramebuffer_Out
}

func (c *GlBindFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindFramebuffer(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("framebuffer:%v", c.In.Framebuffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindFramebuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBindFramebuffer) TypeID() atom.TypeID {
	return 109
}
func (c *GlBindFramebuffer) Flags() atom.Flags {
	return 0
}
func (GlBindFramebuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCheckFramebufferStatus
////////////////////////////////////////////////////////////////////////////////
type GlCheckFramebufferStatus_In struct {
	binary.Generate
	Target FramebufferTarget
}
type GlCheckFramebufferStatus_Out struct {
	binary.Generate
	Result FramebufferStatus
}
type GlCheckFramebufferStatus struct {
	binary.Generate
	Context atom.ContextID
	In      GlCheckFramebufferStatus_In
	Out     GlCheckFramebufferStatus_Out
}

func (c *GlCheckFramebufferStatus) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCheckFramebufferStatus(",
		c.In.Target.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlCheckFramebufferStatus) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCheckFramebufferStatus) TypeID() atom.TypeID {
	return 110
}
func (c *GlCheckFramebufferStatus) Flags() atom.Flags {
	return 0
}
func (GlCheckFramebufferStatus) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteFramebuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteFramebuffers_In struct {
	binary.Generate
	Count        int32
	Framebuffers FramebufferIdArray
}
type GlDeleteFramebuffers_Out struct {
	binary.Generate
}
type GlDeleteFramebuffers struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteFramebuffers_In
	Out     GlDeleteFramebuffers_Out
}

func (c *GlDeleteFramebuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteFramebuffers(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Framebuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteFramebuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteFramebuffers) TypeID() atom.TypeID {
	return 111
}
func (c *GlDeleteFramebuffers) Flags() atom.Flags {
	return 0
}
func (GlDeleteFramebuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsFramebuffer_In struct {
	binary.Generate
	Framebuffer FramebufferId
}
type GlIsFramebuffer_Out struct {
	binary.Generate
	Result bool
}
type GlIsFramebuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsFramebuffer_In
	Out     GlIsFramebuffer_Out
}

func (c *GlIsFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsFramebuffer(",
		fmt.Sprintf("framebuffer:%v", c.In.Framebuffer),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsFramebuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsFramebuffer) TypeID() atom.TypeID {
	return 112
}
func (c *GlIsFramebuffer) Flags() atom.Flags {
	return 0
}
func (GlIsFramebuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenRenderbuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenRenderbuffers_In struct {
	binary.Generate
	Count int32
}
type GlGenRenderbuffers_Out struct {
	binary.Generate
	Renderbuffers RenderbufferIdArray
}
type GlGenRenderbuffers struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenRenderbuffers_In
	Out     GlGenRenderbuffers_Out
}

func (c *GlGenRenderbuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenRenderbuffers(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.Out.Renderbuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenRenderbuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenRenderbuffers) TypeID() atom.TypeID {
	return 113
}
func (c *GlGenRenderbuffers) Flags() atom.Flags {
	return 0
}
func (GlGenRenderbuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBindRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindRenderbuffer_In struct {
	binary.Generate
	Target       RenderbufferTarget
	Renderbuffer RenderbufferId
}
type GlBindRenderbuffer_Out struct {
	binary.Generate
}
type GlBindRenderbuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlBindRenderbuffer_In
	Out     GlBindRenderbuffer_Out
}

func (c *GlBindRenderbuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindRenderbuffer(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("renderbuffer:%v", c.In.Renderbuffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindRenderbuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBindRenderbuffer) TypeID() atom.TypeID {
	return 114
}
func (c *GlBindRenderbuffer) Flags() atom.Flags {
	return 0
}
func (GlBindRenderbuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorage
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorage_In struct {
	binary.Generate
	Target RenderbufferTarget
	Format RenderbufferFormat
	Width  int32
	Height int32
}
type GlRenderbufferStorage_Out struct {
	binary.Generate
}
type GlRenderbufferStorage struct {
	binary.Generate
	Context atom.ContextID
	In      GlRenderbufferStorage_In
	Out     GlRenderbufferStorage_Out
}

func (c *GlRenderbufferStorage) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glRenderbufferStorage(",
		c.In.Target.String(),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlRenderbufferStorage) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlRenderbufferStorage) TypeID() atom.TypeID {
	return 115
}
func (c *GlRenderbufferStorage) Flags() atom.Flags {
	return 0
}
func (GlRenderbufferStorage) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteRenderbuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteRenderbuffers_In struct {
	binary.Generate
	Count         int32
	Renderbuffers RenderbufferIdArray
}
type GlDeleteRenderbuffers_Out struct {
	binary.Generate
}
type GlDeleteRenderbuffers struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteRenderbuffers_In
	Out     GlDeleteRenderbuffers_Out
}

func (c *GlDeleteRenderbuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteRenderbuffers(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Renderbuffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteRenderbuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteRenderbuffers) TypeID() atom.TypeID {
	return 116
}
func (c *GlDeleteRenderbuffers) Flags() atom.Flags {
	return 0
}
func (GlDeleteRenderbuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsRenderbuffer_In struct {
	binary.Generate
	Renderbuffer RenderbufferId
}
type GlIsRenderbuffer_Out struct {
	binary.Generate
	Result bool
}
type GlIsRenderbuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsRenderbuffer_In
	Out     GlIsRenderbuffer_Out
}

func (c *GlIsRenderbuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsRenderbuffer(",
		fmt.Sprintf("renderbuffer:%v", c.In.Renderbuffer),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsRenderbuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsRenderbuffer) TypeID() atom.TypeID {
	return 117
}
func (c *GlIsRenderbuffer) Flags() atom.Flags {
	return 0
}
func (GlIsRenderbuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetRenderbufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetRenderbufferParameteriv_In struct {
	binary.Generate
	Target    RenderbufferTarget
	Parameter RenderbufferParameter
}
type GlGetRenderbufferParameteriv_Out struct {
	binary.Generate
	Values S32Array
}
type GlGetRenderbufferParameteriv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetRenderbufferParameteriv_In
	Out     GlGetRenderbufferParameteriv_Out
}

func (c *GlGetRenderbufferParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetRenderbufferParameteriv(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetRenderbufferParameteriv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetRenderbufferParameteriv) TypeID() atom.TypeID {
	return 118
}
func (c *GlGetRenderbufferParameteriv) Flags() atom.Flags {
	return 0
}
func (GlGetRenderbufferParameteriv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenBuffers
////////////////////////////////////////////////////////////////////////////////
type GlGenBuffers_In struct {
	binary.Generate
	Count int32
}
type GlGenBuffers_Out struct {
	binary.Generate
	Buffers BufferIdArray
}
type GlGenBuffers struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenBuffers_In
	Out     GlGenBuffers_Out
}

func (c *GlGenBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenBuffers(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.Out.Buffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenBuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenBuffers) TypeID() atom.TypeID {
	return 119
}
func (c *GlGenBuffers) Flags() atom.Flags {
	return 0
}
func (GlGenBuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBindBuffer
////////////////////////////////////////////////////////////////////////////////
type GlBindBuffer_In struct {
	binary.Generate
	Target BufferTarget
	Buffer BufferId
}
type GlBindBuffer_Out struct {
	binary.Generate
}
type GlBindBuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlBindBuffer_In
	Out     GlBindBuffer_Out
}

func (c *GlBindBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBindBuffer(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("buffer:%v", c.In.Buffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBindBuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBindBuffer) TypeID() atom.TypeID {
	return 120
}
func (c *GlBindBuffer) Flags() atom.Flags {
	return 0
}
func (GlBindBuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBufferData
////////////////////////////////////////////////////////////////////////////////
type GlBufferData_In struct {
	binary.Generate
	Target BufferTarget
	Size   int32
	Data   BufferDataPointer
	Usage  BufferUsage
}
type GlBufferData_Out struct {
	binary.Generate
}
type GlBufferData struct {
	binary.Generate
	Context atom.ContextID
	In      GlBufferData_In
	Out     GlBufferData_Out
}

func (c *GlBufferData) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBufferData(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("size:%v", c.In.Size),
		", ",
		fmt.Sprintf("data:%v", c.In.Data),
		", ",
		c.In.Usage.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBufferData) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBufferData) TypeID() atom.TypeID {
	return 121
}
func (c *GlBufferData) Flags() atom.Flags {
	return 0
}
func (GlBufferData) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBufferSubData
////////////////////////////////////////////////////////////////////////////////
type GlBufferSubData_In struct {
	binary.Generate
	Target BufferTarget
	Offset int32
	Size   int32
	Data   memory.Pointer
}
type GlBufferSubData_Out struct {
	binary.Generate
}
type GlBufferSubData struct {
	binary.Generate
	Context atom.ContextID
	In      GlBufferSubData_In
	Out     GlBufferSubData_Out
}

func (c *GlBufferSubData) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBufferSubData(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("offset:%v", c.In.Offset),
		", ",
		fmt.Sprintf("size:%v", c.In.Size),
		", ",
		fmt.Sprintf("0x%x", c.In.Data),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBufferSubData) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBufferSubData) TypeID() atom.TypeID {
	return 122
}
func (c *GlBufferSubData) Flags() atom.Flags {
	return 0
}
func (GlBufferSubData) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteBuffers
////////////////////////////////////////////////////////////////////////////////
type GlDeleteBuffers_In struct {
	binary.Generate
	Count   int32
	Buffers BufferIdArray
}
type GlDeleteBuffers_Out struct {
	binary.Generate
}
type GlDeleteBuffers struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteBuffers_In
	Out     GlDeleteBuffers_Out
}

func (c *GlDeleteBuffers) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteBuffers(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Buffers),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteBuffers) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteBuffers) TypeID() atom.TypeID {
	return 123
}
func (c *GlDeleteBuffers) Flags() atom.Flags {
	return 0
}
func (GlDeleteBuffers) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsBuffer
////////////////////////////////////////////////////////////////////////////////
type GlIsBuffer_In struct {
	binary.Generate
	Buffer BufferId
}
type GlIsBuffer_Out struct {
	binary.Generate
	Result bool
}
type GlIsBuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsBuffer_In
	Out     GlIsBuffer_Out
}

func (c *GlIsBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsBuffer(",
		fmt.Sprintf("buffer:%v", c.In.Buffer),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsBuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsBuffer) TypeID() atom.TypeID {
	return 124
}
func (c *GlIsBuffer) Flags() atom.Flags {
	return 0
}
func (GlIsBuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetBufferParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetBufferParameteriv_In struct {
	binary.Generate
	Target    BufferTarget
	Parameter BufferParameter
}
type GlGetBufferParameteriv_Out struct {
	binary.Generate
	Value int32
}
type GlGetBufferParameteriv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetBufferParameteriv_In
	Out     GlGetBufferParameteriv_Out
}

func (c *GlGetBufferParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetBufferParameteriv(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetBufferParameteriv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetBufferParameteriv) TypeID() atom.TypeID {
	return 125
}
func (c *GlGetBufferParameteriv) Flags() atom.Flags {
	return 0
}
func (GlGetBufferParameteriv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCreateShader
////////////////////////////////////////////////////////////////////////////////
type GlCreateShader_In struct {
	binary.Generate
	Type ShaderType
}
type GlCreateShader_Out struct {
	binary.Generate
	Result ShaderId
}
type GlCreateShader struct {
	binary.Generate
	Context atom.ContextID
	In      GlCreateShader_In
	Out     GlCreateShader_Out
}

func (c *GlCreateShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCreateShader(",
		c.In.Type.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlCreateShader) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCreateShader) TypeID() atom.TypeID {
	return 126
}
func (c *GlCreateShader) Flags() atom.Flags {
	return 0
}
func (GlCreateShader) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteShader
////////////////////////////////////////////////////////////////////////////////
type GlDeleteShader_In struct {
	binary.Generate
	Shader ShaderId
}
type GlDeleteShader_Out struct {
	binary.Generate
}
type GlDeleteShader struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteShader_In
	Out     GlDeleteShader_Out
}

func (c *GlDeleteShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteShader(",
		fmt.Sprintf("shader:%v", c.In.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteShader) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteShader) TypeID() atom.TypeID {
	return 127
}
func (c *GlDeleteShader) Flags() atom.Flags {
	return 0
}
func (GlDeleteShader) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlShaderSource
////////////////////////////////////////////////////////////////////////////////
type GlShaderSource_In struct {
	binary.Generate
	Shader ShaderId
	Count  int32
	Source StringArray
	Length S32Array
}
type GlShaderSource_Out struct {
	binary.Generate
}
type GlShaderSource struct {
	binary.Generate
	Context atom.ContextID
	In      GlShaderSource_In
	Out     GlShaderSource_Out
}

func (c *GlShaderSource) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glShaderSource(",
		fmt.Sprintf("shader:%v", c.In.Shader),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Source),
		", ",
		fmt.Sprintf("%v", c.In.Length),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlShaderSource) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlShaderSource) TypeID() atom.TypeID {
	return 128
}
func (c *GlShaderSource) Flags() atom.Flags {
	return 0
}
func (GlShaderSource) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlShaderBinary
////////////////////////////////////////////////////////////////////////////////
type GlShaderBinary_In struct {
	binary.Generate
	Count        int32
	Shaders      ShaderIdArray
	BinaryFormat uint32
	Binary       memory.Pointer
	BinarySize   int32
}
type GlShaderBinary_Out struct {
	binary.Generate
}
type GlShaderBinary struct {
	binary.Generate
	Context atom.ContextID
	In      GlShaderBinary_In
	Out     GlShaderBinary_Out
}

func (c *GlShaderBinary) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glShaderBinary(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Shaders),
		", ",
		fmt.Sprintf("binary_format:%v", c.In.BinaryFormat),
		", ",
		fmt.Sprintf("0x%x", c.In.Binary),
		", ",
		fmt.Sprintf("binary_size:%v", c.In.BinarySize),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlShaderBinary) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlShaderBinary) TypeID() atom.TypeID {
	return 129
}
func (c *GlShaderBinary) Flags() atom.Flags {
	return 0
}
func (GlShaderBinary) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderInfoLog
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderInfoLog_In struct {
	binary.Generate
	Shader       ShaderId
	BufferLength int32
}
type GlGetShaderInfoLog_Out struct {
	binary.Generate
	StringLengthWritten int32
	Info                string
}
type GlGetShaderInfoLog struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetShaderInfoLog_In
	Out     GlGetShaderInfoLog_Out
}

func (c *GlGetShaderInfoLog) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderInfoLog(",
		fmt.Sprintf("shader:%v", c.In.Shader),
		", ",
		fmt.Sprintf("buffer_length:%v", c.In.BufferLength),
		", ",
		fmt.Sprintf("string_length_written:%v", c.Out.StringLengthWritten),
		", ",
		fmt.Sprintf("info:%v", c.Out.Info),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderInfoLog) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetShaderInfoLog) TypeID() atom.TypeID {
	return 130
}
func (c *GlGetShaderInfoLog) Flags() atom.Flags {
	return 0
}
func (GlGetShaderInfoLog) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetShaderSource
////////////////////////////////////////////////////////////////////////////////
type GlGetShaderSource_In struct {
	binary.Generate
	Shader       ShaderId
	BufferLength int32
}
type GlGetShaderSource_Out struct {
	binary.Generate
	StringLengthWritten int32
	Source              string
}
type GlGetShaderSource struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetShaderSource_In
	Out     GlGetShaderSource_Out
}

func (c *GlGetShaderSource) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetShaderSource(",
		fmt.Sprintf("shader:%v", c.In.Shader),
		", ",
		fmt.Sprintf("buffer_length:%v", c.In.BufferLength),
		", ",
		fmt.Sprintf("string_length_written:%v", c.Out.StringLengthWritten),
		", ",
		fmt.Sprintf("source:%v", c.Out.Source),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetShaderSource) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetShaderSource) TypeID() atom.TypeID {
	return 131
}
func (c *GlGetShaderSource) Flags() atom.Flags {
	return 0
}
func (GlGetShaderSource) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlReleaseShaderCompiler
////////////////////////////////////////////////////////////////////////////////
type GlReleaseShaderCompiler_In struct {
	binary.Generate
}
type GlReleaseShaderCompiler_Out struct {
	binary.Generate
}
type GlReleaseShaderCompiler struct {
	binary.Generate
	Context atom.ContextID
	In      GlReleaseShaderCompiler_In
	Out     GlReleaseShaderCompiler_Out
}

func (c *GlReleaseShaderCompiler) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glReleaseShaderCompiler(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlReleaseShaderCompiler) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlReleaseShaderCompiler) TypeID() atom.TypeID {
	return 132
}
func (c *GlReleaseShaderCompiler) Flags() atom.Flags {
	return 0
}
func (GlReleaseShaderCompiler) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCompileShader
////////////////////////////////////////////////////////////////////////////////
type GlCompileShader_In struct {
	binary.Generate
	Shader ShaderId
}
type GlCompileShader_Out struct {
	binary.Generate
}
type GlCompileShader struct {
	binary.Generate
	Context atom.ContextID
	In      GlCompileShader_In
	Out     GlCompileShader_Out
}

func (c *GlCompileShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCompileShader(",
		fmt.Sprintf("shader:%v", c.In.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCompileShader) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCompileShader) TypeID() atom.TypeID {
	return 133
}
func (c *GlCompileShader) Flags() atom.Flags {
	return 0
}
func (GlCompileShader) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsShader
////////////////////////////////////////////////////////////////////////////////
type GlIsShader_In struct {
	binary.Generate
	Shader ShaderId
}
type GlIsShader_Out struct {
	binary.Generate
	Result bool
}
type GlIsShader struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsShader_In
	Out     GlIsShader_Out
}

func (c *GlIsShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsShader(",
		fmt.Sprintf("shader:%v", c.In.Shader),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsShader) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsShader) TypeID() atom.TypeID {
	return 134
}
func (c *GlIsShader) Flags() atom.Flags {
	return 0
}
func (GlIsShader) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCreateProgram
////////////////////////////////////////////////////////////////////////////////
type GlCreateProgram_In struct {
	binary.Generate
}
type GlCreateProgram_Out struct {
	binary.Generate
	Result ProgramId
}
type GlCreateProgram struct {
	binary.Generate
	Context atom.ContextID
	In      GlCreateProgram_In
	Out     GlCreateProgram_Out
}

func (c *GlCreateProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCreateProgram(")
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlCreateProgram) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCreateProgram) TypeID() atom.TypeID {
	return 135
}
func (c *GlCreateProgram) Flags() atom.Flags {
	return 0
}
func (GlCreateProgram) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteProgram
////////////////////////////////////////////////////////////////////////////////
type GlDeleteProgram_In struct {
	binary.Generate
	Program ProgramId
}
type GlDeleteProgram_Out struct {
	binary.Generate
}
type GlDeleteProgram struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteProgram_In
	Out     GlDeleteProgram_Out
}

func (c *GlDeleteProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteProgram(",
		fmt.Sprintf("program:%v", c.In.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteProgram) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteProgram) TypeID() atom.TypeID {
	return 136
}
func (c *GlDeleteProgram) Flags() atom.Flags {
	return 0
}
func (GlDeleteProgram) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlAttachShader
////////////////////////////////////////////////////////////////////////////////
type GlAttachShader_In struct {
	binary.Generate
	Program ProgramId
	Shader  ShaderId
}
type GlAttachShader_Out struct {
	binary.Generate
}
type GlAttachShader struct {
	binary.Generate
	Context atom.ContextID
	In      GlAttachShader_In
	Out     GlAttachShader_Out
}

func (c *GlAttachShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glAttachShader(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("shader:%v", c.In.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlAttachShader) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlAttachShader) TypeID() atom.TypeID {
	return 137
}
func (c *GlAttachShader) Flags() atom.Flags {
	return 0
}
func (GlAttachShader) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDetachShader
////////////////////////////////////////////////////////////////////////////////
type GlDetachShader_In struct {
	binary.Generate
	Program ProgramId
	Shader  ShaderId
}
type GlDetachShader_Out struct {
	binary.Generate
}
type GlDetachShader struct {
	binary.Generate
	Context atom.ContextID
	In      GlDetachShader_In
	Out     GlDetachShader_Out
}

func (c *GlDetachShader) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDetachShader(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("shader:%v", c.In.Shader),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDetachShader) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDetachShader) TypeID() atom.TypeID {
	return 138
}
func (c *GlDetachShader) Flags() atom.Flags {
	return 0
}
func (GlDetachShader) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetAttachedShaders
////////////////////////////////////////////////////////////////////////////////
type GlGetAttachedShaders_In struct {
	binary.Generate
	Program      ProgramId
	BufferLength int32
}
type GlGetAttachedShaders_Out struct {
	binary.Generate
	ShadersLengthWritten int32
	Shaders              ShaderIdArray
}
type GlGetAttachedShaders struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetAttachedShaders_In
	Out     GlGetAttachedShaders_Out
}

func (c *GlGetAttachedShaders) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetAttachedShaders(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("buffer_length:%v", c.In.BufferLength),
		", ",
		fmt.Sprintf("shaders_length_written:%v", c.Out.ShadersLengthWritten),
		", ",
		fmt.Sprintf("%v", c.Out.Shaders),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetAttachedShaders) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetAttachedShaders) TypeID() atom.TypeID {
	return 139
}
func (c *GlGetAttachedShaders) Flags() atom.Flags {
	return 0
}
func (GlGetAttachedShaders) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlLinkProgram
////////////////////////////////////////////////////////////////////////////////
type GlLinkProgram_In struct {
	binary.Generate
	Program ProgramId
}
type GlLinkProgram_Out struct {
	binary.Generate
}
type GlLinkProgram struct {
	binary.Generate
	Context atom.ContextID
	In      GlLinkProgram_In
	Out     GlLinkProgram_Out
}

func (c *GlLinkProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glLinkProgram(",
		fmt.Sprintf("program:%v", c.In.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlLinkProgram) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlLinkProgram) TypeID() atom.TypeID {
	return 140
}
func (c *GlLinkProgram) Flags() atom.Flags {
	return 0
}
func (GlLinkProgram) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetProgramInfoLog
////////////////////////////////////////////////////////////////////////////////
type GlGetProgramInfoLog_In struct {
	binary.Generate
	Program      ProgramId
	BufferLength int32
}
type GlGetProgramInfoLog_Out struct {
	binary.Generate
	StringLengthWritten int32
	Info                string
}
type GlGetProgramInfoLog struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetProgramInfoLog_In
	Out     GlGetProgramInfoLog_Out
}

func (c *GlGetProgramInfoLog) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetProgramInfoLog(",
		fmt.Sprintf("program:%v", c.In.Program),
		", ",
		fmt.Sprintf("buffer_length:%v", c.In.BufferLength),
		", ",
		fmt.Sprintf("string_length_written:%v", c.Out.StringLengthWritten),
		", ",
		fmt.Sprintf("info:%v", c.Out.Info),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetProgramInfoLog) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetProgramInfoLog) TypeID() atom.TypeID {
	return 141
}
func (c *GlGetProgramInfoLog) Flags() atom.Flags {
	return 0
}
func (GlGetProgramInfoLog) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUseProgram
////////////////////////////////////////////////////////////////////////////////
type GlUseProgram_In struct {
	binary.Generate
	Program ProgramId
}
type GlUseProgram_Out struct {
	binary.Generate
}
type GlUseProgram struct {
	binary.Generate
	Context atom.ContextID
	In      GlUseProgram_In
	Out     GlUseProgram_Out
}

func (c *GlUseProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUseProgram(",
		fmt.Sprintf("program:%v", c.In.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUseProgram) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUseProgram) TypeID() atom.TypeID {
	return 142
}
func (c *GlUseProgram) Flags() atom.Flags {
	return 0
}
func (GlUseProgram) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsProgram
////////////////////////////////////////////////////////////////////////////////
type GlIsProgram_In struct {
	binary.Generate
	Program ProgramId
}
type GlIsProgram_Out struct {
	binary.Generate
	Result bool
}
type GlIsProgram struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsProgram_In
	Out     GlIsProgram_Out
}

func (c *GlIsProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsProgram(",
		fmt.Sprintf("program:%v", c.In.Program),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsProgram) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsProgram) TypeID() atom.TypeID {
	return 143
}
func (c *GlIsProgram) Flags() atom.Flags {
	return 0
}
func (GlIsProgram) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlValidateProgram
////////////////////////////////////////////////////////////////////////////////
type GlValidateProgram_In struct {
	binary.Generate
	Program ProgramId
}
type GlValidateProgram_Out struct {
	binary.Generate
}
type GlValidateProgram struct {
	binary.Generate
	Context atom.ContextID
	In      GlValidateProgram_In
	Out     GlValidateProgram_Out
}

func (c *GlValidateProgram) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glValidateProgram(",
		fmt.Sprintf("program:%v", c.In.Program),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlValidateProgram) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlValidateProgram) TypeID() atom.TypeID {
	return 144
}
func (c *GlValidateProgram) Flags() atom.Flags {
	return 0
}
func (GlValidateProgram) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlClearColor
////////////////////////////////////////////////////////////////////////////////
type GlClearColor_In struct {
	binary.Generate
	R float32
	G float32
	B float32
	A float32
}
type GlClearColor_Out struct {
	binary.Generate
}
type GlClearColor struct {
	binary.Generate
	Context atom.ContextID
	In      GlClearColor_In
	Out     GlClearColor_Out
}

func (c *GlClearColor) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClearColor(",
		fmt.Sprintf("r:%v", c.In.R),
		", ",
		fmt.Sprintf("g:%v", c.In.G),
		", ",
		fmt.Sprintf("b:%v", c.In.B),
		", ",
		fmt.Sprintf("a:%v", c.In.A),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClearColor) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlClearColor) TypeID() atom.TypeID {
	return 145
}
func (c *GlClearColor) Flags() atom.Flags {
	return 0
}
func (GlClearColor) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlClearDepthf
////////////////////////////////////////////////////////////////////////////////
type GlClearDepthf_In struct {
	binary.Generate
	Depth float32
}
type GlClearDepthf_Out struct {
	binary.Generate
}
type GlClearDepthf struct {
	binary.Generate
	Context atom.ContextID
	In      GlClearDepthf_In
	Out     GlClearDepthf_Out
}

func (c *GlClearDepthf) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClearDepthf(",
		fmt.Sprintf("depth:%v", c.In.Depth),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClearDepthf) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlClearDepthf) TypeID() atom.TypeID {
	return 146
}
func (c *GlClearDepthf) Flags() atom.Flags {
	return 0
}
func (GlClearDepthf) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlClearStencil
////////////////////////////////////////////////////////////////////////////////
type GlClearStencil_In struct {
	binary.Generate
	Stencil int32
}
type GlClearStencil_Out struct {
	binary.Generate
}
type GlClearStencil struct {
	binary.Generate
	Context atom.ContextID
	In      GlClearStencil_In
	Out     GlClearStencil_Out
}

func (c *GlClearStencil) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClearStencil(",
		fmt.Sprintf("stencil:%v", c.In.Stencil),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClearStencil) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlClearStencil) TypeID() atom.TypeID {
	return 147
}
func (c *GlClearStencil) Flags() atom.Flags {
	return 0
}
func (GlClearStencil) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlClear
////////////////////////////////////////////////////////////////////////////////
type GlClear_In struct {
	binary.Generate
	Mask ClearMask
}
type GlClear_Out struct {
	binary.Generate
}
type GlClear struct {
	binary.Generate
	Context atom.ContextID
	In      GlClear_In
	Out     GlClear_Out
}

func (c *GlClear) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glClear(",
		c.In.Mask.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlClear) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlClear) TypeID() atom.TypeID {
	return 148
}
func (c *GlClear) Flags() atom.Flags {
	return 0
}
func (GlClear) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlCullFace
////////////////////////////////////////////////////////////////////////////////
type GlCullFace_In struct {
	binary.Generate
	Mode FaceMode
}
type GlCullFace_Out struct {
	binary.Generate
}
type GlCullFace struct {
	binary.Generate
	Context atom.ContextID
	In      GlCullFace_In
	Out     GlCullFace_Out
}

func (c *GlCullFace) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glCullFace(",
		c.In.Mode.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlCullFace) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlCullFace) TypeID() atom.TypeID {
	return 149
}
func (c *GlCullFace) Flags() atom.Flags {
	return 0
}
func (GlCullFace) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlPolygonOffset
////////////////////////////////////////////////////////////////////////////////
type GlPolygonOffset_In struct {
	binary.Generate
	ScaleFactor float32
	Units       float32
}
type GlPolygonOffset_Out struct {
	binary.Generate
}
type GlPolygonOffset struct {
	binary.Generate
	Context atom.ContextID
	In      GlPolygonOffset_In
	Out     GlPolygonOffset_Out
}

func (c *GlPolygonOffset) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glPolygonOffset(",
		fmt.Sprintf("scale_factor:%v", c.In.ScaleFactor),
		", ",
		fmt.Sprintf("units:%v", c.In.Units),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlPolygonOffset) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlPolygonOffset) TypeID() atom.TypeID {
	return 150
}
func (c *GlPolygonOffset) Flags() atom.Flags {
	return 0
}
func (GlPolygonOffset) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlLineWidth
////////////////////////////////////////////////////////////////////////////////
type GlLineWidth_In struct {
	binary.Generate
	Width float32
}
type GlLineWidth_Out struct {
	binary.Generate
}
type GlLineWidth struct {
	binary.Generate
	Context atom.ContextID
	In      GlLineWidth_In
	Out     GlLineWidth_Out
}

func (c *GlLineWidth) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glLineWidth(",
		fmt.Sprintf("width:%v", c.In.Width),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlLineWidth) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlLineWidth) TypeID() atom.TypeID {
	return 151
}
func (c *GlLineWidth) Flags() atom.Flags {
	return 0
}
func (GlLineWidth) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlSampleCoverage
////////////////////////////////////////////////////////////////////////////////
type GlSampleCoverage_In struct {
	binary.Generate
	Value  float32
	Invert bool
}
type GlSampleCoverage_Out struct {
	binary.Generate
}
type GlSampleCoverage struct {
	binary.Generate
	Context atom.ContextID
	In      GlSampleCoverage_In
	Out     GlSampleCoverage_Out
}

func (c *GlSampleCoverage) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glSampleCoverage(",
		fmt.Sprintf("value:%v", c.In.Value),
		", ",
		fmt.Sprintf("invert:%v", c.In.Invert),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlSampleCoverage) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlSampleCoverage) TypeID() atom.TypeID {
	return 152
}
func (c *GlSampleCoverage) Flags() atom.Flags {
	return 0
}
func (GlSampleCoverage) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlHint
////////////////////////////////////////////////////////////////////////////////
type GlHint_In struct {
	binary.Generate
	Target HintTarget
	Mode   HintMode
}
type GlHint_Out struct {
	binary.Generate
}
type GlHint struct {
	binary.Generate
	Context atom.ContextID
	In      GlHint_In
	Out     GlHint_Out
}

func (c *GlHint) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glHint(",
		c.In.Target.String(),
		", ",
		c.In.Mode.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlHint) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlHint) TypeID() atom.TypeID {
	return 153
}
func (c *GlHint) Flags() atom.Flags {
	return 0
}
func (GlHint) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferRenderbuffer
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferRenderbuffer_In struct {
	binary.Generate
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	RenderbufferTarget    RenderbufferTarget
	Renderbuffer          RenderbufferId
}
type GlFramebufferRenderbuffer_Out struct {
	binary.Generate
}
type GlFramebufferRenderbuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlFramebufferRenderbuffer_In
	Out     GlFramebufferRenderbuffer_Out
}

func (c *GlFramebufferRenderbuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFramebufferRenderbuffer(",
		c.In.FramebufferTarget.String(),
		", ",
		c.In.FramebufferAttachment.String(),
		", ",
		c.In.RenderbufferTarget.String(),
		", ",
		fmt.Sprintf("renderbuffer:%v", c.In.Renderbuffer),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFramebufferRenderbuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlFramebufferRenderbuffer) TypeID() atom.TypeID {
	return 154
}
func (c *GlFramebufferRenderbuffer) Flags() atom.Flags {
	return 0
}
func (GlFramebufferRenderbuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlFramebufferTexture2D
////////////////////////////////////////////////////////////////////////////////
type GlFramebufferTexture2D_In struct {
	binary.Generate
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	TextureTarget         TextureImageTarget
	Texture               TextureId
	Level                 int32
}
type GlFramebufferTexture2D_Out struct {
	binary.Generate
}
type GlFramebufferTexture2D struct {
	binary.Generate
	Context atom.ContextID
	In      GlFramebufferTexture2D_In
	Out     GlFramebufferTexture2D_Out
}

func (c *GlFramebufferTexture2D) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFramebufferTexture2D(",
		c.In.FramebufferTarget.String(),
		", ",
		c.In.FramebufferAttachment.String(),
		", ",
		c.In.TextureTarget.String(),
		", ",
		fmt.Sprintf("texture:%v", c.In.Texture),
		", ",
		fmt.Sprintf("level:%v", c.In.Level),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFramebufferTexture2D) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlFramebufferTexture2D) TypeID() atom.TypeID {
	return 155
}
func (c *GlFramebufferTexture2D) Flags() atom.Flags {
	return 0
}
func (GlFramebufferTexture2D) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetFramebufferAttachmentParameteriv
////////////////////////////////////////////////////////////////////////////////
type GlGetFramebufferAttachmentParameteriv_In struct {
	binary.Generate
	FramebufferTarget FramebufferTarget
	Attachment        FramebufferAttachment
	Parameter         FramebufferAttachmentParameter
}
type GlGetFramebufferAttachmentParameteriv_Out struct {
	binary.Generate
	Value S32Array
}
type GlGetFramebufferAttachmentParameteriv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetFramebufferAttachmentParameteriv_In
	Out     GlGetFramebufferAttachmentParameteriv_Out
}

func (c *GlGetFramebufferAttachmentParameteriv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetFramebufferAttachmentParameteriv(",
		c.In.FramebufferTarget.String(),
		", ",
		c.In.Attachment.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetFramebufferAttachmentParameteriv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetFramebufferAttachmentParameteriv) TypeID() atom.TypeID {
	return 156
}
func (c *GlGetFramebufferAttachmentParameteriv) Flags() atom.Flags {
	return 0
}
func (GlGetFramebufferAttachmentParameteriv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDrawElements
////////////////////////////////////////////////////////////////////////////////
type GlDrawElements_In struct {
	binary.Generate
	DrawMode     DrawMode
	ElementCount int32
	IndicesType  IndicesType
	Indices      IndicesPointer
}
type GlDrawElements_Out struct {
	binary.Generate
}
type GlDrawElements struct {
	binary.Generate
	Context atom.ContextID
	In      GlDrawElements_In
	Out     GlDrawElements_Out
}

func (c *GlDrawElements) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDrawElements(",
		c.In.DrawMode.String(),
		", ",
		fmt.Sprintf("element_count:%v", c.In.ElementCount),
		", ",
		c.In.IndicesType.String(),
		", ",
		fmt.Sprintf("indices:%v", c.In.Indices),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDrawElements) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDrawElements) TypeID() atom.TypeID {
	return 157
}
func (c *GlDrawElements) Flags() atom.Flags {
	return 0 | atom.DrawCall
}
func (GlDrawElements) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDrawArrays
////////////////////////////////////////////////////////////////////////////////
type GlDrawArrays_In struct {
	binary.Generate
	DrawMode   DrawMode
	FirstIndex int32
	IndexCount int32
}
type GlDrawArrays_Out struct {
	binary.Generate
}
type GlDrawArrays struct {
	binary.Generate
	Context atom.ContextID
	In      GlDrawArrays_In
	Out     GlDrawArrays_Out
}

func (c *GlDrawArrays) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDrawArrays(",
		c.In.DrawMode.String(),
		", ",
		fmt.Sprintf("first_index:%v", c.In.FirstIndex),
		", ",
		fmt.Sprintf("index_count:%v", c.In.IndexCount),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDrawArrays) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDrawArrays) TypeID() atom.TypeID {
	return 158
}
func (c *GlDrawArrays) Flags() atom.Flags {
	return 0 | atom.DrawCall
}
func (GlDrawArrays) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlFlush
////////////////////////////////////////////////////////////////////////////////
type GlFlush_In struct {
	binary.Generate
}
type GlFlush_Out struct {
	binary.Generate
}
type GlFlush struct {
	binary.Generate
	Context atom.ContextID
	In      GlFlush_In
	Out     GlFlush_Out
}

func (c *GlFlush) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFlush(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFlush) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlFlush) TypeID() atom.TypeID {
	return 159
}
func (c *GlFlush) Flags() atom.Flags {
	return 0
}
func (GlFlush) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlFinish
////////////////////////////////////////////////////////////////////////////////
type GlFinish_In struct {
	binary.Generate
}
type GlFinish_Out struct {
	binary.Generate
}
type GlFinish struct {
	binary.Generate
	Context atom.ContextID
	In      GlFinish_In
	Out     GlFinish_Out
}

func (c *GlFinish) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glFinish(")
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlFinish) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlFinish) TypeID() atom.TypeID {
	return 160
}
func (c *GlFinish) Flags() atom.Flags {
	return 0
}
func (GlFinish) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetBooleanv
////////////////////////////////////////////////////////////////////////////////
type GlGetBooleanv_In struct {
	binary.Generate
	Param StateVariable
}
type GlGetBooleanv_Out struct {
	binary.Generate
	Values BoolArray
}
type GlGetBooleanv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetBooleanv_In
	Out     GlGetBooleanv_Out
}

func (c *GlGetBooleanv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetBooleanv(",
		c.In.Param.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetBooleanv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetBooleanv) TypeID() atom.TypeID {
	return 161
}
func (c *GlGetBooleanv) Flags() atom.Flags {
	return 0
}
func (GlGetBooleanv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetFloatv
////////////////////////////////////////////////////////////////////////////////
type GlGetFloatv_In struct {
	binary.Generate
	Param StateVariable
}
type GlGetFloatv_Out struct {
	binary.Generate
	Values F32Array
}
type GlGetFloatv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetFloatv_In
	Out     GlGetFloatv_Out
}

func (c *GlGetFloatv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetFloatv(",
		c.In.Param.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetFloatv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetFloatv) TypeID() atom.TypeID {
	return 162
}
func (c *GlGetFloatv) Flags() atom.Flags {
	return 0
}
func (GlGetFloatv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetIntegerv
////////////////////////////////////////////////////////////////////////////////
type GlGetIntegerv_In struct {
	binary.Generate
	Param StateVariable
}
type GlGetIntegerv_Out struct {
	binary.Generate
	Values S32Array
}
type GlGetIntegerv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetIntegerv_In
	Out     GlGetIntegerv_Out
}

func (c *GlGetIntegerv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetIntegerv(",
		c.In.Param.String(),
		", ",
		fmt.Sprintf("%v", c.Out.Values),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetIntegerv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetIntegerv) TypeID() atom.TypeID {
	return 163
}
func (c *GlGetIntegerv) Flags() atom.Flags {
	return 0
}
func (GlGetIntegerv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetString
////////////////////////////////////////////////////////////////////////////////
type GlGetString_In struct {
	binary.Generate
	Param StringConstant
}
type GlGetString_Out struct {
	binary.Generate
	Result string
}
type GlGetString struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetString_In
	Out     GlGetString_Out
}

func (c *GlGetString) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetString(",
		c.In.Param.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlGetString) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetString) TypeID() atom.TypeID {
	return 164
}
func (c *GlGetString) Flags() atom.Flags {
	return 0
}
func (GlGetString) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEnable
////////////////////////////////////////////////////////////////////////////////
type GlEnable_In struct {
	binary.Generate
	Capability Capability
}
type GlEnable_Out struct {
	binary.Generate
}
type GlEnable struct {
	binary.Generate
	Context atom.ContextID
	In      GlEnable_In
	Out     GlEnable_Out
}

func (c *GlEnable) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEnable(",
		c.In.Capability.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEnable) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEnable) TypeID() atom.TypeID {
	return 165
}
func (c *GlEnable) Flags() atom.Flags {
	return 0
}
func (GlEnable) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDisable
////////////////////////////////////////////////////////////////////////////////
type GlDisable_In struct {
	binary.Generate
	Capability Capability
}
type GlDisable_Out struct {
	binary.Generate
}
type GlDisable struct {
	binary.Generate
	Context atom.ContextID
	In      GlDisable_In
	Out     GlDisable_Out
}

func (c *GlDisable) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDisable(",
		c.In.Capability.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDisable) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDisable) TypeID() atom.TypeID {
	return 166
}
func (c *GlDisable) Flags() atom.Flags {
	return 0
}
func (GlDisable) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsEnabled
////////////////////////////////////////////////////////////////////////////////
type GlIsEnabled_In struct {
	binary.Generate
	Capability Capability
}
type GlIsEnabled_Out struct {
	binary.Generate
	Result bool
}
type GlIsEnabled struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsEnabled_In
	Out     GlIsEnabled_Out
}

func (c *GlIsEnabled) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsEnabled(",
		c.In.Capability.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsEnabled) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsEnabled) TypeID() atom.TypeID {
	return 167
}
func (c *GlIsEnabled) Flags() atom.Flags {
	return 0
}
func (GlIsEnabled) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlMapBufferRange
////////////////////////////////////////////////////////////////////////////////
type GlMapBufferRange_In struct {
	binary.Generate
	Target MapBufferTarget
	Offset int32
	Length int32
	Access MapBufferRangeAccess
}
type GlMapBufferRange_Out struct {
	binary.Generate
	Result memory.Pointer
}
type GlMapBufferRange struct {
	binary.Generate
	Context atom.ContextID
	In      GlMapBufferRange_In
	Out     GlMapBufferRange_Out
}

func (c *GlMapBufferRange) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glMapBufferRange(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("offset:%v", c.In.Offset),
		", ",
		fmt.Sprintf("length:%v", c.In.Length),
		", ",
		c.In.Access.String(),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlMapBufferRange) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlMapBufferRange) TypeID() atom.TypeID {
	return 168
}
func (c *GlMapBufferRange) Flags() atom.Flags {
	return 0
}
func (GlMapBufferRange) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlUnmapBuffer
////////////////////////////////////////////////////////////////////////////////
type GlUnmapBuffer_In struct {
	binary.Generate
	Target MapBufferTarget
}
type GlUnmapBuffer_Out struct {
	binary.Generate
}
type GlUnmapBuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlUnmapBuffer_In
	Out     GlUnmapBuffer_Out
}

func (c *GlUnmapBuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glUnmapBuffer(",
		c.In.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlUnmapBuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlUnmapBuffer) TypeID() atom.TypeID {
	return 169
}
func (c *GlUnmapBuffer) Flags() atom.Flags {
	return 0
}
func (GlUnmapBuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlInvalidateFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlInvalidateFramebuffer_In struct {
	binary.Generate
	Target      FramebufferTarget
	Count       int32
	Attachments FramebufferAttachmentArray
}
type GlInvalidateFramebuffer_Out struct {
	binary.Generate
}
type GlInvalidateFramebuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlInvalidateFramebuffer_In
	Out     GlInvalidateFramebuffer_Out
}

func (c *GlInvalidateFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glInvalidateFramebuffer(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Attachments),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlInvalidateFramebuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlInvalidateFramebuffer) TypeID() atom.TypeID {
	return 170
}
func (c *GlInvalidateFramebuffer) Flags() atom.Flags {
	return 0
}
func (GlInvalidateFramebuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlRenderbufferStorageMultisample
////////////////////////////////////////////////////////////////////////////////
type GlRenderbufferStorageMultisample_In struct {
	binary.Generate
	Target  RenderbufferTarget
	Samples int32
	Format  RenderbufferFormat
	Width   int32
	Height  int32
}
type GlRenderbufferStorageMultisample_Out struct {
	binary.Generate
}
type GlRenderbufferStorageMultisample struct {
	binary.Generate
	Context atom.ContextID
	In      GlRenderbufferStorageMultisample_In
	Out     GlRenderbufferStorageMultisample_Out
}

func (c *GlRenderbufferStorageMultisample) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glRenderbufferStorageMultisample(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("samples:%v", c.In.Samples),
		", ",
		c.In.Format.String(),
		", ",
		fmt.Sprintf("width:%v", c.In.Width),
		", ",
		fmt.Sprintf("height:%v", c.In.Height),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlRenderbufferStorageMultisample) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlRenderbufferStorageMultisample) TypeID() atom.TypeID {
	return 171
}
func (c *GlRenderbufferStorageMultisample) Flags() atom.Flags {
	return 0
}
func (GlRenderbufferStorageMultisample) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBlitFramebuffer
////////////////////////////////////////////////////////////////////////////////
type GlBlitFramebuffer_In struct {
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
type GlBlitFramebuffer_Out struct {
	binary.Generate
}
type GlBlitFramebuffer struct {
	binary.Generate
	Context atom.ContextID
	In      GlBlitFramebuffer_In
	Out     GlBlitFramebuffer_Out
}

func (c *GlBlitFramebuffer) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBlitFramebuffer(",
		fmt.Sprintf("srcX0:%v", c.In.SrcX0),
		", ",
		fmt.Sprintf("srcY0:%v", c.In.SrcY0),
		", ",
		fmt.Sprintf("srcX1:%v", c.In.SrcX1),
		", ",
		fmt.Sprintf("srcY1:%v", c.In.SrcY1),
		", ",
		fmt.Sprintf("dstX0:%v", c.In.DstX0),
		", ",
		fmt.Sprintf("dstY0:%v", c.In.DstY0),
		", ",
		fmt.Sprintf("dstX1:%v", c.In.DstX1),
		", ",
		fmt.Sprintf("dstY1:%v", c.In.DstY1),
		", ",
		c.In.Mask.String(),
		", ",
		c.In.Filter.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBlitFramebuffer) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBlitFramebuffer) TypeID() atom.TypeID {
	return 172
}
func (c *GlBlitFramebuffer) Flags() atom.Flags {
	return 0
}
func (GlBlitFramebuffer) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenQueries
////////////////////////////////////////////////////////////////////////////////
type GlGenQueries_In struct {
	binary.Generate
	Count int32
}
type GlGenQueries_Out struct {
	binary.Generate
	Queries QueryIdArray
}
type GlGenQueries struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenQueries_In
	Out     GlGenQueries_Out
}

func (c *GlGenQueries) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenQueries(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.Out.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenQueries) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenQueries) TypeID() atom.TypeID {
	return 173
}
func (c *GlGenQueries) Flags() atom.Flags {
	return 0
}
func (GlGenQueries) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBeginQuery
////////////////////////////////////////////////////////////////////////////////
type GlBeginQuery_In struct {
	binary.Generate
	Target QueryTarget
	Query  QueryId
}
type GlBeginQuery_Out struct {
	binary.Generate
}
type GlBeginQuery struct {
	binary.Generate
	Context atom.ContextID
	In      GlBeginQuery_In
	Out     GlBeginQuery_Out
}

func (c *GlBeginQuery) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBeginQuery(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("query:%v", c.In.Query),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBeginQuery) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBeginQuery) TypeID() atom.TypeID {
	return 174
}
func (c *GlBeginQuery) Flags() atom.Flags {
	return 0
}
func (GlBeginQuery) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEndQuery
////////////////////////////////////////////////////////////////////////////////
type GlEndQuery_In struct {
	binary.Generate
	Target QueryTarget
}
type GlEndQuery_Out struct {
	binary.Generate
}
type GlEndQuery struct {
	binary.Generate
	Context atom.ContextID
	In      GlEndQuery_In
	Out     GlEndQuery_Out
}

func (c *GlEndQuery) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEndQuery(",
		c.In.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEndQuery) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEndQuery) TypeID() atom.TypeID {
	return 175
}
func (c *GlEndQuery) Flags() atom.Flags {
	return 0
}
func (GlEndQuery) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueries
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueries_In struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}
type GlDeleteQueries_Out struct {
	binary.Generate
}
type GlDeleteQueries struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteQueries_In
	Out     GlDeleteQueries_Out
}

func (c *GlDeleteQueries) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteQueries(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteQueries) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteQueries) TypeID() atom.TypeID {
	return 176
}
func (c *GlDeleteQueries) Flags() atom.Flags {
	return 0
}
func (GlDeleteQueries) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsQuery
////////////////////////////////////////////////////////////////////////////////
type GlIsQuery_In struct {
	binary.Generate
	Query QueryId
}
type GlIsQuery_Out struct {
	binary.Generate
	Result bool
}
type GlIsQuery struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsQuery_In
	Out     GlIsQuery_Out
}

func (c *GlIsQuery) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsQuery(",
		fmt.Sprintf("query:%v", c.In.Query),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsQuery) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsQuery) TypeID() atom.TypeID {
	return 177
}
func (c *GlIsQuery) Flags() atom.Flags {
	return 0
}
func (GlIsQuery) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryiv_In struct {
	binary.Generate
	Target    QueryTarget
	Parameter QueryParameter
}
type GlGetQueryiv_Out struct {
	binary.Generate
	Value int32
}
type GlGetQueryiv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetQueryiv_In
	Out     GlGetQueryiv_Out
}

func (c *GlGetQueryiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryiv(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryiv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetQueryiv) TypeID() atom.TypeID {
	return 178
}
func (c *GlGetQueryiv) Flags() atom.Flags {
	return 0
}
func (GlGetQueryiv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuiv
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuiv_In struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
}
type GlGetQueryObjectuiv_Out struct {
	binary.Generate
	Value uint32
}
type GlGetQueryObjectuiv struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetQueryObjectuiv_In
	Out     GlGetQueryObjectuiv_Out
}

func (c *GlGetQueryObjectuiv) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectuiv(",
		fmt.Sprintf("query:%v", c.In.Query),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectuiv) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetQueryObjectuiv) TypeID() atom.TypeID {
	return 179
}
func (c *GlGetQueryObjectuiv) Flags() atom.Flags {
	return 0
}
func (GlGetQueryObjectuiv) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGenQueriesEXT
////////////////////////////////////////////////////////////////////////////////
type GlGenQueriesEXT_In struct {
	binary.Generate
	Count int32
}
type GlGenQueriesEXT_Out struct {
	binary.Generate
	Queries QueryIdArray
}
type GlGenQueriesEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlGenQueriesEXT_In
	Out     GlGenQueriesEXT_Out
}

func (c *GlGenQueriesEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGenQueriesEXT(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.Out.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGenQueriesEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGenQueriesEXT) TypeID() atom.TypeID {
	return 180
}
func (c *GlGenQueriesEXT) Flags() atom.Flags {
	return 0
}
func (GlGenQueriesEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlBeginQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlBeginQueryEXT_In struct {
	binary.Generate
	Target QueryTarget
	Query  QueryId
}
type GlBeginQueryEXT_Out struct {
	binary.Generate
}
type GlBeginQueryEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlBeginQueryEXT_In
	Out     GlBeginQueryEXT_Out
}

func (c *GlBeginQueryEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glBeginQueryEXT(",
		c.In.Target.String(),
		", ",
		fmt.Sprintf("query:%v", c.In.Query),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlBeginQueryEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlBeginQueryEXT) TypeID() atom.TypeID {
	return 181
}
func (c *GlBeginQueryEXT) Flags() atom.Flags {
	return 0
}
func (GlBeginQueryEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlEndQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlEndQueryEXT_In struct {
	binary.Generate
	Target QueryTarget
}
type GlEndQueryEXT_Out struct {
	binary.Generate
}
type GlEndQueryEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlEndQueryEXT_In
	Out     GlEndQueryEXT_Out
}

func (c *GlEndQueryEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glEndQueryEXT(",
		c.In.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlEndQueryEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlEndQueryEXT) TypeID() atom.TypeID {
	return 182
}
func (c *GlEndQueryEXT) Flags() atom.Flags {
	return 0
}
func (GlEndQueryEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlDeleteQueriesEXT
////////////////////////////////////////////////////////////////////////////////
type GlDeleteQueriesEXT_In struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}
type GlDeleteQueriesEXT_Out struct {
	binary.Generate
}
type GlDeleteQueriesEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlDeleteQueriesEXT_In
	Out     GlDeleteQueriesEXT_Out
}

func (c *GlDeleteQueriesEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glDeleteQueriesEXT(",
		fmt.Sprintf("count:%v", c.In.Count),
		", ",
		fmt.Sprintf("%v", c.In.Queries),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlDeleteQueriesEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlDeleteQueriesEXT) TypeID() atom.TypeID {
	return 183
}
func (c *GlDeleteQueriesEXT) Flags() atom.Flags {
	return 0
}
func (GlDeleteQueriesEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlIsQueryEXT
////////////////////////////////////////////////////////////////////////////////
type GlIsQueryEXT_In struct {
	binary.Generate
	Query QueryId
}
type GlIsQueryEXT_Out struct {
	binary.Generate
	Result bool
}
type GlIsQueryEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlIsQueryEXT_In
	Out     GlIsQueryEXT_Out
}

func (c *GlIsQueryEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glIsQueryEXT(",
		fmt.Sprintf("query:%v", c.In.Query),
	)
	parts = append(parts, ")")
	parts = append(parts, fmt.Sprintf(" → %v", c.Out.Result))
	return strings.Join(parts, "")
}
func (c *GlIsQueryEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlIsQueryEXT) TypeID() atom.TypeID {
	return 184
}
func (c *GlIsQueryEXT) Flags() atom.Flags {
	return 0
}
func (GlIsQueryEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlQueryCounterEXT
////////////////////////////////////////////////////////////////////////////////
type GlQueryCounterEXT_In struct {
	binary.Generate
	Query  QueryId
	Target QueryTarget
}
type GlQueryCounterEXT_Out struct {
	binary.Generate
}
type GlQueryCounterEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlQueryCounterEXT_In
	Out     GlQueryCounterEXT_Out
}

func (c *GlQueryCounterEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glQueryCounterEXT(",
		fmt.Sprintf("query:%v", c.In.Query),
		", ",
		c.In.Target.String(),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlQueryCounterEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlQueryCounterEXT) TypeID() atom.TypeID {
	return 185
}
func (c *GlQueryCounterEXT) Flags() atom.Flags {
	return 0
}
func (GlQueryCounterEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryivEXT_In struct {
	binary.Generate
	Target    QueryTarget
	Parameter QueryParameter
}
type GlGetQueryivEXT_Out struct {
	binary.Generate
	Value int32
}
type GlGetQueryivEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetQueryivEXT_In
	Out     GlGetQueryivEXT_Out
}

func (c *GlGetQueryivEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryivEXT(",
		c.In.Target.String(),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryivEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetQueryivEXT) TypeID() atom.TypeID {
	return 186
}
func (c *GlGetQueryivEXT) Flags() atom.Flags {
	return 0
}
func (GlGetQueryivEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectivEXT_In struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
}
type GlGetQueryObjectivEXT_Out struct {
	binary.Generate
	Value int32
}
type GlGetQueryObjectivEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetQueryObjectivEXT_In
	Out     GlGetQueryObjectivEXT_Out
}

func (c *GlGetQueryObjectivEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectivEXT(",
		fmt.Sprintf("query:%v", c.In.Query),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectivEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetQueryObjectivEXT) TypeID() atom.TypeID {
	return 187
}
func (c *GlGetQueryObjectivEXT) Flags() atom.Flags {
	return 0
}
func (GlGetQueryObjectivEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectuivEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectuivEXT_In struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
}
type GlGetQueryObjectuivEXT_Out struct {
	binary.Generate
	Value uint32
}
type GlGetQueryObjectuivEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetQueryObjectuivEXT_In
	Out     GlGetQueryObjectuivEXT_Out
}

func (c *GlGetQueryObjectuivEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectuivEXT(",
		fmt.Sprintf("query:%v", c.In.Query),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectuivEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetQueryObjectuivEXT) TypeID() atom.TypeID {
	return 188
}
func (c *GlGetQueryObjectuivEXT) Flags() atom.Flags {
	return 0
}
func (GlGetQueryObjectuivEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjecti64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjecti64vEXT_In struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
}
type GlGetQueryObjecti64vEXT_Out struct {
	binary.Generate
	Value int64
}
type GlGetQueryObjecti64vEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetQueryObjecti64vEXT_In
	Out     GlGetQueryObjecti64vEXT_Out
}

func (c *GlGetQueryObjecti64vEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjecti64vEXT(",
		fmt.Sprintf("query:%v", c.In.Query),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjecti64vEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetQueryObjecti64vEXT) TypeID() atom.TypeID {
	return 189
}
func (c *GlGetQueryObjecti64vEXT) Flags() atom.Flags {
	return 0
}
func (GlGetQueryObjecti64vEXT) API() gfxapi.API { return API() }

////////////////////////////////////////////////////////////////////////////////
// GlGetQueryObjectui64vEXT
////////////////////////////////////////////////////////////////////////////////
type GlGetQueryObjectui64vEXT_In struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
}
type GlGetQueryObjectui64vEXT_Out struct {
	binary.Generate
	Value uint64
}
type GlGetQueryObjectui64vEXT struct {
	binary.Generate
	Context atom.ContextID
	In      GlGetQueryObjectui64vEXT_In
	Out     GlGetQueryObjectui64vEXT_Out
}

func (c *GlGetQueryObjectui64vEXT) String() string {
	parts := make([]string, 0, 32)
	parts = append(parts, "glGetQueryObjectui64vEXT(",
		fmt.Sprintf("query:%v", c.In.Query),
		", ",
		c.In.Parameter.String(),
		", ",
		fmt.Sprintf("value:%v", c.Out.Value),
	)
	parts = append(parts, ")")
	return strings.Join(parts, "")
}
func (c *GlGetQueryObjectui64vEXT) ContextID() atom.ContextID {
	return c.Context
}
func (c *GlGetQueryObjectui64vEXT) TypeID() atom.TypeID {
	return 190
}
func (c *GlGetQueryObjectui64vEXT) Flags() atom.Flags {
	return 0
}
func (GlGetQueryObjectui64vEXT) API() gfxapi.API { return API() }

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
	Size      int32
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
	Source    StringArray
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
	AttributeBindings AttributeLocation_stringMap
	Attributes        VertexAttribute_s32Map
	Uniforms          Uniform_UniformLocationMap
	InfoLog           string
}

func (c *Program) Init() {
	c.Shaders = make(ShaderId_ShaderTypeMap)
	c.AttributeBindings = make(AttributeLocation_stringMap)
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
	Size       int32
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
// class InternalState
////////////////////////////////////////////////////////////////////////////////
type InternalState struct {
	binary.Generate
	CreatedAt       atom.ID
	NilBuffer       BufferId
	NilTexture      TextureId
	NilRenderbuffer RenderbufferId
	Backbuffer      FramebufferId
}

func (c *InternalState) Init() {
	c.NilBuffer = 0
	c.NilTexture = 0
	c.NilRenderbuffer = 0
	c.Backbuffer = 0
}
func (c *InternalState) GetCreatedAt() atom.ID { return c.CreatedAt }

////////////////////////////////////////////////////////////////////////////////
// class Objects
////////////////////////////////////////////////////////////////////////////////
type Objects struct {
	binary.Generate
	CreatedAt     atom.ID
	Renderbuffers RenderbufferRef_RenderbufferIdMap
	Textures      TextureRef_TextureIdMap
	Framebuffers  FramebufferRef_FramebufferIdMap
	Buffers       BufferRef_BufferIdMap
	Shaders       ShaderRef_ShaderIdMap
	Programs      ProgramRef_ProgramIdMap
	VertexArrays  VertexArrayRef_VertexArrayIdMap
	Queries       QueryRef_QueryIdMap
}

func (c *Objects) Init() {
	c.Renderbuffers = make(RenderbufferRef_RenderbufferIdMap)
	c.Textures = make(TextureRef_TextureIdMap)
	c.Framebuffers = make(FramebufferRef_FramebufferIdMap)
	c.Buffers = make(BufferRef_BufferIdMap)
	c.Shaders = make(ShaderRef_ShaderIdMap)
	c.Programs = make(ProgramRef_ProgramIdMap)
	c.VertexArrays = make(VertexArrayRef_VertexArrayIdMap)
	c.Queries = make(QueryRef_QueryIdMap)
}
func (c *Objects) GetCreatedAt() atom.ID { return c.CreatedAt }

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
// enum BufferTarget
////////////////////////////////////////////////////////////////////////////////
type BufferTarget uint32

const (
	BufferTarget_GL_ARRAY_BUFFER         = BufferTarget(34962)
	BufferTarget_GL_ELEMENT_ARRAY_BUFFER = BufferTarget(34963)
)

func (v BufferTarget) String() string {
	switch v {
	case 34962:
		return "GL_ARRAY_BUFFER"
	case 34963:
		return "GL_ELEMENT_ARRAY_BUFFER"
	default:
		return fmt.Sprintf("BufferTarget<0x%.4x>", uint32(v))
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
// enum MapBufferTarget
////////////////////////////////////////////////////////////////////////////////
type MapBufferTarget uint32

const (
	MapBufferTarget_GL_ARRAY_BUFFER              = MapBufferTarget(34962)
	MapBufferTarget_GL_COPY_READ_BUFFER          = MapBufferTarget(36662)
	MapBufferTarget_GL_COPY_WRITE_BUFFER         = MapBufferTarget(36663)
	MapBufferTarget_GL_ELEMENT_ARRAY_BUFFER      = MapBufferTarget(34963)
	MapBufferTarget_GL_PIXEL_PACK_BUFFER         = MapBufferTarget(35051)
	MapBufferTarget_GL_PIXEL_UNPACK_BUFFER       = MapBufferTarget(35052)
	MapBufferTarget_GL_TRANSFORM_FEEDBACK_BUFFER = MapBufferTarget(35982)
	MapBufferTarget_GL_UNIFORM_BUFFER            = MapBufferTarget(35345)
)

func (v MapBufferTarget) String() string {
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
		return fmt.Sprintf("MapBufferTarget<0x%.4x>", uint32(v))
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
	Blending              BlendState
	Rasterizing           RasterizerState
	Clearing              ClearState
	BoundFramebuffers     FramebufferId_FramebufferTargetMap
	BoundRenderbuffers    RenderbufferId_RenderbufferTargetMap
	BoundBuffers          BufferId_BufferTargetMap
	BoundProgram          ProgramId
	BoundVertexArray      VertexArrayId
	VertexAttributeArrays VertexAttributeArrayRef_AttributeLocationMap
	TextureUnits          TextureId_TextureTargetMap_TextureUnitMap
	ActiveTextureUnit     TextureUnit
	Capabilities          Bool_CapabilityMap
	Internals             InternalState
	GenerateMipmapHint    HintMode
	PixelStorage          S32_PixelStoreParameterMap
	Instances             Objects
}

func (g *Globals) Init() {
	g.Blending.Init()
	g.Rasterizing.Init()
	g.Clearing.Init()
	g.BoundFramebuffers = make(FramebufferId_FramebufferTargetMap)
	g.BoundRenderbuffers = make(RenderbufferId_RenderbufferTargetMap)
	g.BoundBuffers = make(BufferId_BufferTargetMap)
	g.VertexAttributeArrays = make(VertexAttributeArrayRef_AttributeLocationMap)
	g.TextureUnits = make(TextureId_TextureTargetMap_TextureUnitMap)
	g.ActiveTextureUnit = TextureUnit_GL_TEXTURE0
	g.Capabilities = make(Bool_CapabilityMap)
	g.Internals.Init()
	g.GenerateMipmapHint = HintMode_GL_DONT_CARE
	g.PixelStorage = make(S32_PixelStoreParameterMap)
	g.Instances.Init()
}
func NewInit(
	contextID atom.ContextID,
	pWidth int32,
	pHeight int32,
	pColorFmt RenderbufferFormat,
	pDepthFmt RenderbufferFormat,
	pStencilFmt RenderbufferFormat,
) *Init {
	return &Init{
		Context: contextID,
		In:      Init_In{Width: pWidth, Height: pHeight, ColorFmt: pColorFmt, DepthFmt: pDepthFmt, StencilFmt: pStencilFmt},
		Out:     Init_Out{},
	}
}
func NewStartTimer(
	contextID atom.ContextID,
	pIndex uint8,
) *StartTimer {
	return &StartTimer{
		Context: contextID,
		In:      StartTimer_In{Index: pIndex},
		Out:     StartTimer_Out{},
	}
}
func NewStopTimer(
	contextID atom.ContextID,
	pIndex uint8,
	pResult uint64,
) *StopTimer {
	return &StopTimer{
		Context: contextID,
		In:      StopTimer_In{Index: pIndex},
		Out:     StopTimer_Out{Result: pResult},
	}
}
func NewFlushPostBuffer(
	contextID atom.ContextID,
) *FlushPostBuffer {
	return &FlushPostBuffer{
		Context: contextID,
		In:      FlushPostBuffer_In{},
		Out:     FlushPostBuffer_Out{},
	}
}
func NewEglCreateContext(
	contextID atom.ContextID,
	pVersion int32,
	pContext int32,
) *EglCreateContext {
	return &EglCreateContext{
		Context: contextID,
		In:      EglCreateContext_In{},
		Out:     EglCreateContext_Out{Version: pVersion, Context: pContext},
	}
}
func NewEglMakeCurrent(
	contextID atom.ContextID,
	pContext int32,
) *EglMakeCurrent {
	return &EglMakeCurrent{
		Context: contextID,
		In:      EglMakeCurrent_In{Context: pContext},
		Out:     EglMakeCurrent_Out{},
	}
}
func NewEglSwapBuffers(
	contextID atom.ContextID,
) *EglSwapBuffers {
	return &EglSwapBuffers{
		Context: contextID,
		In:      EglSwapBuffers_In{},
		Out:     EglSwapBuffers_Out{},
	}
}
func NewWglCreateContext(
	contextID atom.ContextID,
	pHdc HDC,
	pResult HGLRC,
) *WglCreateContext {
	return &WglCreateContext{
		Context: contextID,
		In:      WglCreateContext_In{Hdc: pHdc},
		Out:     WglCreateContext_Out{Result: pResult},
	}
}
func NewWglMakeCurrent(
	contextID atom.ContextID,
	pHdc HDC,
	pHglrc HGLRC,
	pResult BOOL,
) *WglMakeCurrent {
	return &WglMakeCurrent{
		Context: contextID,
		In:      WglMakeCurrent_In{Hdc: pHdc, Hglrc: pHglrc},
		Out:     WglMakeCurrent_Out{Result: pResult},
	}
}
func NewWglSwapBuffers(
	contextID atom.ContextID,
	pHdc HDC,
) *WglSwapBuffers {
	return &WglSwapBuffers{
		Context: contextID,
		In:      WglSwapBuffers_In{Hdc: pHdc},
		Out:     WglSwapBuffers_Out{},
	}
}
func NewGlEnableClientState(
	contextID atom.ContextID,
	pType ArrayType,
) *GlEnableClientState {
	return &GlEnableClientState{
		Context: contextID,
		In:      GlEnableClientState_In{Type: pType},
		Out:     GlEnableClientState_Out{},
	}
}
func NewGlDisableClientState(
	contextID atom.ContextID,
	pType ArrayType,
) *GlDisableClientState {
	return &GlDisableClientState{
		Context: contextID,
		In:      GlDisableClientState_In{Type: pType},
		Out:     GlDisableClientState_Out{},
	}
}
func NewGlGetProgramBinaryOES(
	contextID atom.ContextID,
	pProgram ProgramId,
	pBufferSize int32,
	pBytesWritten int32,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
) *GlGetProgramBinaryOES {
	return &GlGetProgramBinaryOES{
		Context: contextID,
		In:      GlGetProgramBinaryOES_In{Program: pProgram, BufferSize: pBufferSize},
		Out:     GlGetProgramBinaryOES_Out{BytesWritten: pBytesWritten, BinaryFormat: pBinaryFormat, Binary: pBinary},
	}
}
func NewGlProgramBinaryOES(
	contextID atom.ContextID,
	pProgram ProgramId,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
	pBinarySize int32,
) *GlProgramBinaryOES {
	return &GlProgramBinaryOES{
		Context: contextID,
		In:      GlProgramBinaryOES_In{Program: pProgram, BinaryFormat: pBinaryFormat, Binary: pBinary, BinarySize: pBinarySize},
		Out:     GlProgramBinaryOES_Out{},
	}
}
func NewGlStartTilingQCOM(
	contextID atom.ContextID,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pPreserveMask TilePreserveMaskQCOM,
) *GlStartTilingQCOM {
	return &GlStartTilingQCOM{
		Context: contextID,
		In:      GlStartTilingQCOM_In{X: pX, Y: pY, Width: pWidth, Height: pHeight, PreserveMask: pPreserveMask},
		Out:     GlStartTilingQCOM_Out{},
	}
}
func NewGlEndTilingQCOM(
	contextID atom.ContextID,
	pPreserveMask TilePreserveMaskQCOM,
) *GlEndTilingQCOM {
	return &GlEndTilingQCOM{
		Context: contextID,
		In:      GlEndTilingQCOM_In{PreserveMask: pPreserveMask},
		Out:     GlEndTilingQCOM_Out{},
	}
}
func NewGlDiscardFramebufferEXT(
	contextID atom.ContextID,
	pTarget FramebufferTarget,
	pNumAttachments int32,
	pAttachments DiscardFramebufferAttachmentArray,
) *GlDiscardFramebufferEXT {
	return &GlDiscardFramebufferEXT{
		Context: contextID,
		In:      GlDiscardFramebufferEXT_In{Target: pTarget, NumAttachments: pNumAttachments, Attachments: pAttachments},
		Out:     GlDiscardFramebufferEXT_Out{},
	}
}
func NewGlInsertEventMarkerEXT(
	contextID atom.ContextID,
	pLength int32,
	pMarker string,
) *GlInsertEventMarkerEXT {
	return &GlInsertEventMarkerEXT{
		Context: contextID,
		In:      GlInsertEventMarkerEXT_In{Length: pLength, Marker: pMarker},
		Out:     GlInsertEventMarkerEXT_Out{},
	}
}
func NewGlPushGroupMarkerEXT(
	contextID atom.ContextID,
	pLength int32,
	pMarker string,
) *GlPushGroupMarkerEXT {
	return &GlPushGroupMarkerEXT{
		Context: contextID,
		In:      GlPushGroupMarkerEXT_In{Length: pLength, Marker: pMarker},
		Out:     GlPushGroupMarkerEXT_Out{},
	}
}
func NewGlPopGroupMarkerEXT(
	contextID atom.ContextID,
) *GlPopGroupMarkerEXT {
	return &GlPopGroupMarkerEXT{
		Context: contextID,
		In:      GlPopGroupMarkerEXT_In{},
		Out:     GlPopGroupMarkerEXT_Out{},
	}
}
func NewGlTexStorage1DEXT(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
) *GlTexStorage1DEXT {
	return &GlTexStorage1DEXT{
		Context: contextID,
		In:      GlTexStorage1DEXT_In{Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth},
		Out:     GlTexStorage1DEXT_Out{},
	}
}
func NewGlTexStorage2DEXT(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
) *GlTexStorage2DEXT {
	return &GlTexStorage2DEXT{
		Context: contextID,
		In:      GlTexStorage2DEXT_In{Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight},
		Out:     GlTexStorage2DEXT_Out{},
	}
}
func NewGlTexStorage3DEXT(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
	pDepth int32,
) *GlTexStorage3DEXT {
	return &GlTexStorage3DEXT{
		Context: contextID,
		In:      GlTexStorage3DEXT_In{Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight, Depth: pDepth},
		Out:     GlTexStorage3DEXT_Out{},
	}
}
func NewGlTextureStorage1DEXT(
	contextID atom.ContextID,
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
) *GlTextureStorage1DEXT {
	return &GlTextureStorage1DEXT{
		Context: contextID,
		In:      GlTextureStorage1DEXT_In{Texture: pTexture, Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth},
		Out:     GlTextureStorage1DEXT_Out{},
	}
}
func NewGlTextureStorage2DEXT(
	contextID atom.ContextID,
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
) *GlTextureStorage2DEXT {
	return &GlTextureStorage2DEXT{
		Context: contextID,
		In:      GlTextureStorage2DEXT_In{Texture: pTexture, Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight},
		Out:     GlTextureStorage2DEXT_Out{},
	}
}
func NewGlTextureStorage3DEXT(
	contextID atom.ContextID,
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
	pDepth int32,
) *GlTextureStorage3DEXT {
	return &GlTextureStorage3DEXT{
		Context: contextID,
		In:      GlTextureStorage3DEXT_In{Texture: pTexture, Target: pTarget, Levels: pLevels, Format: pFormat, Width: pWidth, Height: pHeight, Depth: pDepth},
		Out:     GlTextureStorage3DEXT_Out{},
	}
}
func NewGlGenVertexArraysOES(
	contextID atom.ContextID,
	pCount int32,
	pArrays VertexArrayIdArray,
) *GlGenVertexArraysOES {
	return &GlGenVertexArraysOES{
		Context: contextID,
		In:      GlGenVertexArraysOES_In{Count: pCount},
		Out:     GlGenVertexArraysOES_Out{Arrays: pArrays},
	}
}
func NewGlBindVertexArrayOES(
	contextID atom.ContextID,
	pArray VertexArrayId,
) *GlBindVertexArrayOES {
	return &GlBindVertexArrayOES{
		Context: contextID,
		In:      GlBindVertexArrayOES_In{Array: pArray},
		Out:     GlBindVertexArrayOES_Out{},
	}
}
func NewGlDeleteVertexArraysOES(
	contextID atom.ContextID,
	pCount int32,
	pArrays VertexArrayIdArray,
) *GlDeleteVertexArraysOES {
	return &GlDeleteVertexArraysOES{
		Context: contextID,
		In:      GlDeleteVertexArraysOES_In{Count: pCount, Arrays: pArrays},
		Out:     GlDeleteVertexArraysOES_Out{},
	}
}
func NewGlIsVertexArrayOES(
	contextID atom.ContextID,
	pArray VertexArrayId,
	pResult bool,
) *GlIsVertexArrayOES {
	return &GlIsVertexArrayOES{
		Context: contextID,
		In:      GlIsVertexArrayOES_In{Array: pArray},
		Out:     GlIsVertexArrayOES_Out{Result: pResult},
	}
}
func NewGlEGLImageTargetTexture2DOES(
	contextID atom.ContextID,
	pTarget ImageTargetTexture,
	pImage ImageOES,
) *GlEGLImageTargetTexture2DOES {
	return &GlEGLImageTargetTexture2DOES{
		Context: contextID,
		In:      GlEGLImageTargetTexture2DOES_In{Target: pTarget, Image: pImage},
		Out:     GlEGLImageTargetTexture2DOES_Out{},
	}
}
func NewGlEGLImageTargetRenderbufferStorageOES(
	contextID atom.ContextID,
	pTarget ImageTargetRenderbufferStorage,
	pImage TexturePointer,
) *GlEGLImageTargetRenderbufferStorageOES {
	return &GlEGLImageTargetRenderbufferStorageOES{
		Context: contextID,
		In:      GlEGLImageTargetRenderbufferStorageOES_In{Target: pTarget, Image: pImage},
		Out:     GlEGLImageTargetRenderbufferStorageOES_Out{},
	}
}
func NewGlGetGraphicsResetStatusEXT(
	contextID atom.ContextID,
	pResult ResetStatus,
) *GlGetGraphicsResetStatusEXT {
	return &GlGetGraphicsResetStatusEXT{
		Context: contextID,
		In:      GlGetGraphicsResetStatusEXT_In{},
		Out:     GlGetGraphicsResetStatusEXT_Out{Result: pResult},
	}
}
func NewGlBindAttribLocation(
	contextID atom.ContextID,
	pProgram ProgramId,
	pLocation AttributeLocation,
	pName string,
) *GlBindAttribLocation {
	return &GlBindAttribLocation{
		Context: contextID,
		In:      GlBindAttribLocation_In{Program: pProgram, Location: pLocation, Name: pName},
		Out:     GlBindAttribLocation_Out{},
	}
}
func NewGlBlendFunc(
	contextID atom.ContextID,
	pSrcFactor BlendFactor,
	pDstFactor BlendFactor,
) *GlBlendFunc {
	return &GlBlendFunc{
		Context: contextID,
		In:      GlBlendFunc_In{SrcFactor: pSrcFactor, DstFactor: pDstFactor},
		Out:     GlBlendFunc_Out{},
	}
}
func NewGlBlendFuncSeparate(
	contextID atom.ContextID,
	pSrcFactorRgb BlendFactor,
	pDstFactorRgb BlendFactor,
	pSrcFactorAlpha BlendFactor,
	pDstFactorAlpha BlendFactor,
) *GlBlendFuncSeparate {
	return &GlBlendFuncSeparate{
		Context: contextID,
		In:      GlBlendFuncSeparate_In{SrcFactorRgb: pSrcFactorRgb, DstFactorRgb: pDstFactorRgb, SrcFactorAlpha: pSrcFactorAlpha, DstFactorAlpha: pDstFactorAlpha},
		Out:     GlBlendFuncSeparate_Out{},
	}
}
func NewGlBlendEquation(
	contextID atom.ContextID,
	pEquation BlendEquation,
) *GlBlendEquation {
	return &GlBlendEquation{
		Context: contextID,
		In:      GlBlendEquation_In{Equation: pEquation},
		Out:     GlBlendEquation_Out{},
	}
}
func NewGlBlendEquationSeparate(
	contextID atom.ContextID,
	pRgb BlendEquation,
	pAlpha BlendEquation,
) *GlBlendEquationSeparate {
	return &GlBlendEquationSeparate{
		Context: contextID,
		In:      GlBlendEquationSeparate_In{Rgb: pRgb, Alpha: pAlpha},
		Out:     GlBlendEquationSeparate_Out{},
	}
}
func NewGlBlendColor(
	contextID atom.ContextID,
	pRed float32,
	pGreen float32,
	pBlue float32,
	pAlpha float32,
) *GlBlendColor {
	return &GlBlendColor{
		Context: contextID,
		In:      GlBlendColor_In{Red: pRed, Green: pGreen, Blue: pBlue, Alpha: pAlpha},
		Out:     GlBlendColor_Out{},
	}
}
func NewGlEnableVertexAttribArray(
	contextID atom.ContextID,
	pLocation AttributeLocation,
) *GlEnableVertexAttribArray {
	return &GlEnableVertexAttribArray{
		Context: contextID,
		In:      GlEnableVertexAttribArray_In{Location: pLocation},
		Out:     GlEnableVertexAttribArray_Out{},
	}
}
func NewGlDisableVertexAttribArray(
	contextID atom.ContextID,
	pLocation AttributeLocation,
) *GlDisableVertexAttribArray {
	return &GlDisableVertexAttribArray{
		Context: contextID,
		In:      GlDisableVertexAttribArray_In{Location: pLocation},
		Out:     GlDisableVertexAttribArray_Out{},
	}
}
func NewGlVertexAttribPointer(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pSize int32,
	pType VertexAttribType,
	pNormalized bool,
	pStride int32,
	pData VertexPointer,
) *GlVertexAttribPointer {
	return &GlVertexAttribPointer{
		Context: contextID,
		In:      GlVertexAttribPointer_In{Location: pLocation, Size: pSize, Type: pType, Normalized: pNormalized, Stride: pStride, Data: pData},
		Out:     GlVertexAttribPointer_Out{},
	}
}
func NewGlGetActiveAttrib(
	contextID atom.ContextID,
	pProgram ProgramId,
	pLocation AttributeLocation,
	pBufferSize int32,
	pBufferBytesWritten int32,
	pVectorCount int32,
	pType ShaderAttribType,
	pName string,
) *GlGetActiveAttrib {
	return &GlGetActiveAttrib{
		Context: contextID,
		In:      GlGetActiveAttrib_In{Program: pProgram, Location: pLocation, BufferSize: pBufferSize},
		Out:     GlGetActiveAttrib_Out{BufferBytesWritten: pBufferBytesWritten, VectorCount: pVectorCount, Type: pType, Name: pName},
	}
}
func NewGlGetActiveUniform(
	contextID atom.ContextID,
	pProgram ProgramId,
	pLocation int32,
	pBufferSize int32,
	pBufferBytesWritten int32,
	pSize int32,
	pType ShaderUniformType,
	pName string,
) *GlGetActiveUniform {
	return &GlGetActiveUniform{
		Context: contextID,
		In:      GlGetActiveUniform_In{Program: pProgram, Location: pLocation, BufferSize: pBufferSize},
		Out:     GlGetActiveUniform_Out{BufferBytesWritten: pBufferBytesWritten, Size: pSize, Type: pType, Name: pName},
	}
}
func NewGlGetError(
	contextID atom.ContextID,
	pResult Error,
) *GlGetError {
	return &GlGetError{
		Context: contextID,
		In:      GlGetError_In{},
		Out:     GlGetError_Out{Result: pResult},
	}
}
func NewGlGetProgramiv(
	contextID atom.ContextID,
	pProgram ProgramId,
	pParameter ProgramParameter,
	pValue S32Array,
) *GlGetProgramiv {
	return &GlGetProgramiv{
		Context: contextID,
		In:      GlGetProgramiv_In{Program: pProgram, Parameter: pParameter},
		Out:     GlGetProgramiv_Out{Value: pValue},
	}
}
func NewGlGetShaderiv(
	contextID atom.ContextID,
	pShader ShaderId,
	pParameter ShaderParameter,
	pValue S32Array,
) *GlGetShaderiv {
	return &GlGetShaderiv{
		Context: contextID,
		In:      GlGetShaderiv_In{Shader: pShader, Parameter: pParameter},
		Out:     GlGetShaderiv_Out{Value: pValue},
	}
}
func NewGlGetUniformLocation(
	contextID atom.ContextID,
	pProgram ProgramId,
	pName string,
	pResult UniformLocation,
) *GlGetUniformLocation {
	return &GlGetUniformLocation{
		Context: contextID,
		In:      GlGetUniformLocation_In{Program: pProgram, Name: pName},
		Out:     GlGetUniformLocation_Out{Result: pResult},
	}
}
func NewGlGetAttribLocation(
	contextID atom.ContextID,
	pProgram ProgramId,
	pName string,
	pResult AttributeLocation,
) *GlGetAttribLocation {
	return &GlGetAttribLocation{
		Context: contextID,
		In:      GlGetAttribLocation_In{Program: pProgram, Name: pName},
		Out:     GlGetAttribLocation_Out{Result: pResult},
	}
}
func NewGlPixelStorei(
	contextID atom.ContextID,
	pParameter PixelStoreParameter,
	pValue int32,
) *GlPixelStorei {
	return &GlPixelStorei{
		Context: contextID,
		In:      GlPixelStorei_In{Parameter: pParameter, Value: pValue},
		Out:     GlPixelStorei_Out{},
	}
}
func NewGlTexParameteri(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValue int32,
) *GlTexParameteri {
	return &GlTexParameteri{
		Context: contextID,
		In:      GlTexParameteri_In{Target: pTarget, Parameter: pParameter, Value: pValue},
		Out:     GlTexParameteri_Out{},
	}
}
func NewGlTexParameterf(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValue float32,
) *GlTexParameterf {
	return &GlTexParameterf{
		Context: contextID,
		In:      GlTexParameterf_In{Target: pTarget, Parameter: pParameter, Value: pValue},
		Out:     GlTexParameterf_Out{},
	}
}
func NewGlGetTexParameteriv(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValues S32Array,
) *GlGetTexParameteriv {
	return &GlGetTexParameteriv{
		Context: contextID,
		In:      GlGetTexParameteriv_In{Target: pTarget, Parameter: pParameter},
		Out:     GlGetTexParameteriv_Out{Values: pValues},
	}
}
func NewGlGetTexParameterfv(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValues F32Array,
) *GlGetTexParameterfv {
	return &GlGetTexParameterfv{
		Context: contextID,
		In:      GlGetTexParameterfv_In{Target: pTarget, Parameter: pParameter},
		Out:     GlGetTexParameterfv_Out{Values: pValues},
	}
}
func NewGlUniform1i(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue int32,
) *GlUniform1i {
	return &GlUniform1i{
		Context: contextID,
		In:      GlUniform1i_In{Location: pLocation, Value: pValue},
		Out:     GlUniform1i_Out{},
	}
}
func NewGlUniform2i(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
) *GlUniform2i {
	return &GlUniform2i{
		Context: contextID,
		In:      GlUniform2i_In{Location: pLocation, Value0: pValue0, Value1: pValue1},
		Out:     GlUniform2i_Out{},
	}
}
func NewGlUniform3i(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
	pValue2 int32,
) *GlUniform3i {
	return &GlUniform3i{
		Context: contextID,
		In:      GlUniform3i_In{Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2},
		Out:     GlUniform3i_Out{},
	}
}
func NewGlUniform4i(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
	pValue2 int32,
	pValue3 int32,
) *GlUniform4i {
	return &GlUniform4i{
		Context: contextID,
		In:      GlUniform4i_In{Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2, Value3: pValue3},
		Out:     GlUniform4i_Out{},
	}
}
func NewGlUniform1iv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform1iv {
	return &GlUniform1iv{
		Context: contextID,
		In:      GlUniform1iv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform1iv_Out{},
	}
}
func NewGlUniform2iv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform2iv {
	return &GlUniform2iv{
		Context: contextID,
		In:      GlUniform2iv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform2iv_Out{},
	}
}
func NewGlUniform3iv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform3iv {
	return &GlUniform3iv{
		Context: contextID,
		In:      GlUniform3iv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform3iv_Out{},
	}
}
func NewGlUniform4iv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform4iv {
	return &GlUniform4iv{
		Context: contextID,
		In:      GlUniform4iv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform4iv_Out{},
	}
}
func NewGlUniform1f(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue float32,
) *GlUniform1f {
	return &GlUniform1f{
		Context: contextID,
		In:      GlUniform1f_In{Location: pLocation, Value: pValue},
		Out:     GlUniform1f_Out{},
	}
}
func NewGlUniform2f(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
) *GlUniform2f {
	return &GlUniform2f{
		Context: contextID,
		In:      GlUniform2f_In{Location: pLocation, Value0: pValue0, Value1: pValue1},
		Out:     GlUniform2f_Out{},
	}
}
func NewGlUniform3f(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
) *GlUniform3f {
	return &GlUniform3f{
		Context: contextID,
		In:      GlUniform3f_In{Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2},
		Out:     GlUniform3f_Out{},
	}
}
func NewGlUniform4f(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
	pValue3 float32,
) *GlUniform4f {
	return &GlUniform4f{
		Context: contextID,
		In:      GlUniform4f_In{Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2, Value3: pValue3},
		Out:     GlUniform4f_Out{},
	}
}
func NewGlUniform1fv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform1fv {
	return &GlUniform1fv{
		Context: contextID,
		In:      GlUniform1fv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform1fv_Out{},
	}
}
func NewGlUniform2fv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform2fv {
	return &GlUniform2fv{
		Context: contextID,
		In:      GlUniform2fv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform2fv_Out{},
	}
}
func NewGlUniform3fv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform3fv {
	return &GlUniform3fv{
		Context: contextID,
		In:      GlUniform3fv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform3fv_Out{},
	}
}
func NewGlUniform4fv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform4fv {
	return &GlUniform4fv{
		Context: contextID,
		In:      GlUniform4fv_In{Location: pLocation, Count: pCount, Value: pValue},
		Out:     GlUniform4fv_Out{},
	}
}
func NewGlUniformMatrix2fv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix2fv {
	return &GlUniformMatrix2fv{
		Context: contextID,
		In:      GlUniformMatrix2fv_In{Location: pLocation, Count: pCount, Transpose: pTranspose, Values: pValues},
		Out:     GlUniformMatrix2fv_Out{},
	}
}
func NewGlUniformMatrix3fv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix3fv {
	return &GlUniformMatrix3fv{
		Context: contextID,
		In:      GlUniformMatrix3fv_In{Location: pLocation, Count: pCount, Transpose: pTranspose, Values: pValues},
		Out:     GlUniformMatrix3fv_Out{},
	}
}
func NewGlUniformMatrix4fv(
	contextID atom.ContextID,
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix4fv {
	return &GlUniformMatrix4fv{
		Context: contextID,
		In:      GlUniformMatrix4fv_In{Location: pLocation, Count: pCount, Transpose: pTranspose, Values: pValues},
		Out:     GlUniformMatrix4fv_Out{},
	}
}
func NewGlGetUniformfv(
	contextID atom.ContextID,
	pProgram ProgramId,
	pLocation UniformLocation,
	pValues F32Array,
) *GlGetUniformfv {
	return &GlGetUniformfv{
		Context: contextID,
		In:      GlGetUniformfv_In{Program: pProgram, Location: pLocation, Values: pValues},
		Out:     GlGetUniformfv_Out{},
	}
}
func NewGlGetUniformiv(
	contextID atom.ContextID,
	pProgram ProgramId,
	pLocation UniformLocation,
	pValues S32Array,
) *GlGetUniformiv {
	return &GlGetUniformiv{
		Context: contextID,
		In:      GlGetUniformiv_In{Program: pProgram, Location: pLocation, Values: pValues},
		Out:     GlGetUniformiv_Out{},
	}
}
func NewGlVertexAttrib1f(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue0 float32,
) *GlVertexAttrib1f {
	return &GlVertexAttrib1f{
		Context: contextID,
		In:      GlVertexAttrib1f_In{Location: pLocation, Value0: pValue0},
		Out:     GlVertexAttrib1f_Out{},
	}
}
func NewGlVertexAttrib2f(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
) *GlVertexAttrib2f {
	return &GlVertexAttrib2f{
		Context: contextID,
		In:      GlVertexAttrib2f_In{Location: pLocation, Value0: pValue0, Value1: pValue1},
		Out:     GlVertexAttrib2f_Out{},
	}
}
func NewGlVertexAttrib3f(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
) *GlVertexAttrib3f {
	return &GlVertexAttrib3f{
		Context: contextID,
		In:      GlVertexAttrib3f_In{Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2},
		Out:     GlVertexAttrib3f_Out{},
	}
}
func NewGlVertexAttrib4f(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
	pValue3 float32,
) *GlVertexAttrib4f {
	return &GlVertexAttrib4f{
		Context: contextID,
		In:      GlVertexAttrib4f_In{Location: pLocation, Value0: pValue0, Value1: pValue1, Value2: pValue2, Value3: pValue3},
		Out:     GlVertexAttrib4f_Out{},
	}
}
func NewGlVertexAttrib1fv(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib1fv {
	return &GlVertexAttrib1fv{
		Context: contextID,
		In:      GlVertexAttrib1fv_In{Location: pLocation, Value: pValue},
		Out:     GlVertexAttrib1fv_Out{},
	}
}
func NewGlVertexAttrib2fv(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib2fv {
	return &GlVertexAttrib2fv{
		Context: contextID,
		In:      GlVertexAttrib2fv_In{Location: pLocation, Value: pValue},
		Out:     GlVertexAttrib2fv_Out{},
	}
}
func NewGlVertexAttrib3fv(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib3fv {
	return &GlVertexAttrib3fv{
		Context: contextID,
		In:      GlVertexAttrib3fv_In{Location: pLocation, Value: pValue},
		Out:     GlVertexAttrib3fv_Out{},
	}
}
func NewGlVertexAttrib4fv(
	contextID atom.ContextID,
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib4fv {
	return &GlVertexAttrib4fv{
		Context: contextID,
		In:      GlVertexAttrib4fv_In{Location: pLocation, Value: pValue},
		Out:     GlVertexAttrib4fv_Out{},
	}
}
func NewGlGetShaderPrecisionFormat(
	contextID atom.ContextID,
	pShaderType ShaderType,
	pPrecisionType PrecisionType,
	pRange S32Array,
	pPrecision int32,
) *GlGetShaderPrecisionFormat {
	return &GlGetShaderPrecisionFormat{
		Context: contextID,
		In:      GlGetShaderPrecisionFormat_In{ShaderType: pShaderType, PrecisionType: pPrecisionType},
		Out:     GlGetShaderPrecisionFormat_Out{Range: pRange, Precision: pPrecision},
	}
}
func NewGlDepthMask(
	contextID atom.ContextID,
	pEnabled bool,
) *GlDepthMask {
	return &GlDepthMask{
		Context: contextID,
		In:      GlDepthMask_In{Enabled: pEnabled},
		Out:     GlDepthMask_Out{},
	}
}
func NewGlDepthFunc(
	contextID atom.ContextID,
	pFunction TestFunction,
) *GlDepthFunc {
	return &GlDepthFunc{
		Context: contextID,
		In:      GlDepthFunc_In{Function: pFunction},
		Out:     GlDepthFunc_Out{},
	}
}
func NewGlDepthRangef(
	contextID atom.ContextID,
	pNear float32,
	pFar float32,
) *GlDepthRangef {
	return &GlDepthRangef{
		Context: contextID,
		In:      GlDepthRangef_In{Near: pNear, Far: pFar},
		Out:     GlDepthRangef_Out{},
	}
}
func NewGlColorMask(
	contextID atom.ContextID,
	pRed bool,
	pGreen bool,
	pBlue bool,
	pAlpha bool,
) *GlColorMask {
	return &GlColorMask{
		Context: contextID,
		In:      GlColorMask_In{Red: pRed, Green: pGreen, Blue: pBlue, Alpha: pAlpha},
		Out:     GlColorMask_Out{},
	}
}
func NewGlStencilMask(
	contextID atom.ContextID,
	pMask uint32,
) *GlStencilMask {
	return &GlStencilMask{
		Context: contextID,
		In:      GlStencilMask_In{Mask: pMask},
		Out:     GlStencilMask_Out{},
	}
}
func NewGlStencilMaskSeparate(
	contextID atom.ContextID,
	pFace FaceMode,
	pMask uint32,
) *GlStencilMaskSeparate {
	return &GlStencilMaskSeparate{
		Context: contextID,
		In:      GlStencilMaskSeparate_In{Face: pFace, Mask: pMask},
		Out:     GlStencilMaskSeparate_Out{},
	}
}
func NewGlStencilFuncSeparate(
	contextID atom.ContextID,
	pFace FaceMode,
	pFunction TestFunction,
	pReferenceValue int32,
	pMask int32,
) *GlStencilFuncSeparate {
	return &GlStencilFuncSeparate{
		Context: contextID,
		In:      GlStencilFuncSeparate_In{Face: pFace, Function: pFunction, ReferenceValue: pReferenceValue, Mask: pMask},
		Out:     GlStencilFuncSeparate_Out{},
	}
}
func NewGlStencilOpSeparate(
	contextID atom.ContextID,
	pFace FaceMode,
	pStencilFail StencilAction,
	pStencilPassDepthFail StencilAction,
	pStencilPassDepthPass StencilAction,
) *GlStencilOpSeparate {
	return &GlStencilOpSeparate{
		Context: contextID,
		In:      GlStencilOpSeparate_In{Face: pFace, StencilFail: pStencilFail, StencilPassDepthFail: pStencilPassDepthFail, StencilPassDepthPass: pStencilPassDepthPass},
		Out:     GlStencilOpSeparate_Out{},
	}
}
func NewGlFrontFace(
	contextID atom.ContextID,
	pOrientation FaceOrientation,
) *GlFrontFace {
	return &GlFrontFace{
		Context: contextID,
		In:      GlFrontFace_In{Orientation: pOrientation},
		Out:     GlFrontFace_Out{},
	}
}
func NewGlViewport(
	contextID atom.ContextID,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlViewport {
	return &GlViewport{
		Context: contextID,
		In:      GlViewport_In{X: pX, Y: pY, Width: pWidth, Height: pHeight},
		Out:     GlViewport_Out{},
	}
}
func NewGlScissor(
	contextID atom.ContextID,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlScissor {
	return &GlScissor{
		Context: contextID,
		In:      GlScissor_In{X: pX, Y: pY, Width: pWidth, Height: pHeight},
		Out:     GlScissor_Out{},
	}
}
func NewGlActiveTexture(
	contextID atom.ContextID,
	pUnit TextureUnit,
) *GlActiveTexture {
	return &GlActiveTexture{
		Context: contextID,
		In:      GlActiveTexture_In{Unit: pUnit},
		Out:     GlActiveTexture_Out{},
	}
}
func NewGlGenTextures(
	contextID atom.ContextID,
	pCount int32,
	pTextures TextureIdArray,
) *GlGenTextures {
	return &GlGenTextures{
		Context: contextID,
		In:      GlGenTextures_In{Count: pCount},
		Out:     GlGenTextures_Out{Textures: pTextures},
	}
}
func NewGlDeleteTextures(
	contextID atom.ContextID,
	pCount int32,
	pTextures TextureIdArray,
) *GlDeleteTextures {
	return &GlDeleteTextures{
		Context: contextID,
		In:      GlDeleteTextures_In{Count: pCount, Textures: pTextures},
		Out:     GlDeleteTextures_Out{},
	}
}
func NewGlIsTexture(
	contextID atom.ContextID,
	pTexture TextureId,
	pResult bool,
) *GlIsTexture {
	return &GlIsTexture{
		Context: contextID,
		In:      GlIsTexture_In{Texture: pTexture},
		Out:     GlIsTexture_Out{Result: pResult},
	}
}
func NewGlBindTexture(
	contextID atom.ContextID,
	pTarget TextureTarget,
	pTexture TextureId,
) *GlBindTexture {
	return &GlBindTexture{
		Context: contextID,
		In:      GlBindTexture_In{Target: pTarget, Texture: pTexture},
		Out:     GlBindTexture_Out{},
	}
}
func NewGlTexImage2D(
	contextID atom.ContextID,
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
		Context: contextID,
		In:      GlTexImage2D_In{Target: pTarget, Level: pLevel, InternalFormat: pInternalFormat, Width: pWidth, Height: pHeight, Border: pBorder, Format: pFormat, Type: pType, Data: pData},
		Out:     GlTexImage2D_Out{},
	}
}
func NewGlTexSubImage2D(
	contextID atom.ContextID,
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
		Context: contextID,
		In:      GlTexSubImage2D_In{Target: pTarget, Level: pLevel, Xoffset: pXoffset, Yoffset: pYoffset, Width: pWidth, Height: pHeight, Format: pFormat, Type: pType, Data: pData},
		Out:     GlTexSubImage2D_Out{},
	}
}
func NewGlCopyTexImage2D(
	contextID atom.ContextID,
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
		Context: contextID,
		In:      GlCopyTexImage2D_In{Target: pTarget, Level: pLevel, Format: pFormat, X: pX, Y: pY, Width: pWidth, Height: pHeight, Border: pBorder},
		Out:     GlCopyTexImage2D_Out{},
	}
}
func NewGlCopyTexSubImage2D(
	contextID atom.ContextID,
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
		Context: contextID,
		In:      GlCopyTexSubImage2D_In{Target: pTarget, Level: pLevel, Xoffset: pXoffset, Yoffset: pYoffset, X: pX, Y: pY, Width: pWidth, Height: pHeight},
		Out:     GlCopyTexSubImage2D_Out{},
	}
}
func NewGlCompressedTexImage2D(
	contextID atom.ContextID,
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
		Context: contextID,
		In:      GlCompressedTexImage2D_In{Target: pTarget, Level: pLevel, Format: pFormat, Width: pWidth, Height: pHeight, Border: pBorder, ImageSize: pImageSize, Data: pData},
		Out:     GlCompressedTexImage2D_Out{},
	}
}
func NewGlCompressedTexSubImage2D(
	contextID atom.ContextID,
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
		Context: contextID,
		In:      GlCompressedTexSubImage2D_In{Target: pTarget, Level: pLevel, Xoffset: pXoffset, Yoffset: pYoffset, Width: pWidth, Height: pHeight, Format: pFormat, ImageSize: pImageSize, Data: pData},
		Out:     GlCompressedTexSubImage2D_Out{},
	}
}
func NewGlGenerateMipmap(
	contextID atom.ContextID,
	pTarget TextureImageTarget,
) *GlGenerateMipmap {
	return &GlGenerateMipmap{
		Context: contextID,
		In:      GlGenerateMipmap_In{Target: pTarget},
		Out:     GlGenerateMipmap_Out{},
	}
}
func NewGlReadPixels(
	contextID atom.ContextID,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pFormat BaseTexelFormat,
	pType TexelType,
	pData memory.Pointer,
) *GlReadPixels {
	return &GlReadPixels{
		Context: contextID,
		In:      GlReadPixels_In{X: pX, Y: pY, Width: pWidth, Height: pHeight, Format: pFormat, Type: pType},
		Out:     GlReadPixels_Out{Data: pData},
	}
}
func NewGlGenFramebuffers(
	contextID atom.ContextID,
	pCount int32,
	pFramebuffers FramebufferIdArray,
) *GlGenFramebuffers {
	return &GlGenFramebuffers{
		Context: contextID,
		In:      GlGenFramebuffers_In{Count: pCount},
		Out:     GlGenFramebuffers_Out{Framebuffers: pFramebuffers},
	}
}
func NewGlBindFramebuffer(
	contextID atom.ContextID,
	pTarget FramebufferTarget,
	pFramebuffer FramebufferId,
) *GlBindFramebuffer {
	return &GlBindFramebuffer{
		Context: contextID,
		In:      GlBindFramebuffer_In{Target: pTarget, Framebuffer: pFramebuffer},
		Out:     GlBindFramebuffer_Out{},
	}
}
func NewGlCheckFramebufferStatus(
	contextID atom.ContextID,
	pTarget FramebufferTarget,
	pResult FramebufferStatus,
) *GlCheckFramebufferStatus {
	return &GlCheckFramebufferStatus{
		Context: contextID,
		In:      GlCheckFramebufferStatus_In{Target: pTarget},
		Out:     GlCheckFramebufferStatus_Out{Result: pResult},
	}
}
func NewGlDeleteFramebuffers(
	contextID atom.ContextID,
	pCount int32,
	pFramebuffers FramebufferIdArray,
) *GlDeleteFramebuffers {
	return &GlDeleteFramebuffers{
		Context: contextID,
		In:      GlDeleteFramebuffers_In{Count: pCount, Framebuffers: pFramebuffers},
		Out:     GlDeleteFramebuffers_Out{},
	}
}
func NewGlIsFramebuffer(
	contextID atom.ContextID,
	pFramebuffer FramebufferId,
	pResult bool,
) *GlIsFramebuffer {
	return &GlIsFramebuffer{
		Context: contextID,
		In:      GlIsFramebuffer_In{Framebuffer: pFramebuffer},
		Out:     GlIsFramebuffer_Out{Result: pResult},
	}
}
func NewGlGenRenderbuffers(
	contextID atom.ContextID,
	pCount int32,
	pRenderbuffers RenderbufferIdArray,
) *GlGenRenderbuffers {
	return &GlGenRenderbuffers{
		Context: contextID,
		In:      GlGenRenderbuffers_In{Count: pCount},
		Out:     GlGenRenderbuffers_Out{Renderbuffers: pRenderbuffers},
	}
}
func NewGlBindRenderbuffer(
	contextID atom.ContextID,
	pTarget RenderbufferTarget,
	pRenderbuffer RenderbufferId,
) *GlBindRenderbuffer {
	return &GlBindRenderbuffer{
		Context: contextID,
		In:      GlBindRenderbuffer_In{Target: pTarget, Renderbuffer: pRenderbuffer},
		Out:     GlBindRenderbuffer_Out{},
	}
}
func NewGlRenderbufferStorage(
	contextID atom.ContextID,
	pTarget RenderbufferTarget,
	pFormat RenderbufferFormat,
	pWidth int32,
	pHeight int32,
) *GlRenderbufferStorage {
	return &GlRenderbufferStorage{
		Context: contextID,
		In:      GlRenderbufferStorage_In{Target: pTarget, Format: pFormat, Width: pWidth, Height: pHeight},
		Out:     GlRenderbufferStorage_Out{},
	}
}
func NewGlDeleteRenderbuffers(
	contextID atom.ContextID,
	pCount int32,
	pRenderbuffers RenderbufferIdArray,
) *GlDeleteRenderbuffers {
	return &GlDeleteRenderbuffers{
		Context: contextID,
		In:      GlDeleteRenderbuffers_In{Count: pCount, Renderbuffers: pRenderbuffers},
		Out:     GlDeleteRenderbuffers_Out{},
	}
}
func NewGlIsRenderbuffer(
	contextID atom.ContextID,
	pRenderbuffer RenderbufferId,
	pResult bool,
) *GlIsRenderbuffer {
	return &GlIsRenderbuffer{
		Context: contextID,
		In:      GlIsRenderbuffer_In{Renderbuffer: pRenderbuffer},
		Out:     GlIsRenderbuffer_Out{Result: pResult},
	}
}
func NewGlGetRenderbufferParameteriv(
	contextID atom.ContextID,
	pTarget RenderbufferTarget,
	pParameter RenderbufferParameter,
	pValues S32Array,
) *GlGetRenderbufferParameteriv {
	return &GlGetRenderbufferParameteriv{
		Context: contextID,
		In:      GlGetRenderbufferParameteriv_In{Target: pTarget, Parameter: pParameter},
		Out:     GlGetRenderbufferParameteriv_Out{Values: pValues},
	}
}
func NewGlGenBuffers(
	contextID atom.ContextID,
	pCount int32,
	pBuffers BufferIdArray,
) *GlGenBuffers {
	return &GlGenBuffers{
		Context: contextID,
		In:      GlGenBuffers_In{Count: pCount},
		Out:     GlGenBuffers_Out{Buffers: pBuffers},
	}
}
func NewGlBindBuffer(
	contextID atom.ContextID,
	pTarget BufferTarget,
	pBuffer BufferId,
) *GlBindBuffer {
	return &GlBindBuffer{
		Context: contextID,
		In:      GlBindBuffer_In{Target: pTarget, Buffer: pBuffer},
		Out:     GlBindBuffer_Out{},
	}
}
func NewGlBufferData(
	contextID atom.ContextID,
	pTarget BufferTarget,
	pSize int32,
	pData BufferDataPointer,
	pUsage BufferUsage,
) *GlBufferData {
	return &GlBufferData{
		Context: contextID,
		In:      GlBufferData_In{Target: pTarget, Size: pSize, Data: pData, Usage: pUsage},
		Out:     GlBufferData_Out{},
	}
}
func NewGlBufferSubData(
	contextID atom.ContextID,
	pTarget BufferTarget,
	pOffset int32,
	pSize int32,
	pData memory.Pointer,
) *GlBufferSubData {
	return &GlBufferSubData{
		Context: contextID,
		In:      GlBufferSubData_In{Target: pTarget, Offset: pOffset, Size: pSize, Data: pData},
		Out:     GlBufferSubData_Out{},
	}
}
func NewGlDeleteBuffers(
	contextID atom.ContextID,
	pCount int32,
	pBuffers BufferIdArray,
) *GlDeleteBuffers {
	return &GlDeleteBuffers{
		Context: contextID,
		In:      GlDeleteBuffers_In{Count: pCount, Buffers: pBuffers},
		Out:     GlDeleteBuffers_Out{},
	}
}
func NewGlIsBuffer(
	contextID atom.ContextID,
	pBuffer BufferId,
	pResult bool,
) *GlIsBuffer {
	return &GlIsBuffer{
		Context: contextID,
		In:      GlIsBuffer_In{Buffer: pBuffer},
		Out:     GlIsBuffer_Out{Result: pResult},
	}
}
func NewGlGetBufferParameteriv(
	contextID atom.ContextID,
	pTarget BufferTarget,
	pParameter BufferParameter,
	pValue int32,
) *GlGetBufferParameteriv {
	return &GlGetBufferParameteriv{
		Context: contextID,
		In:      GlGetBufferParameteriv_In{Target: pTarget, Parameter: pParameter},
		Out:     GlGetBufferParameteriv_Out{Value: pValue},
	}
}
func NewGlCreateShader(
	contextID atom.ContextID,
	pType ShaderType,
	pResult ShaderId,
) *GlCreateShader {
	return &GlCreateShader{
		Context: contextID,
		In:      GlCreateShader_In{Type: pType},
		Out:     GlCreateShader_Out{Result: pResult},
	}
}
func NewGlDeleteShader(
	contextID atom.ContextID,
	pShader ShaderId,
) *GlDeleteShader {
	return &GlDeleteShader{
		Context: contextID,
		In:      GlDeleteShader_In{Shader: pShader},
		Out:     GlDeleteShader_Out{},
	}
}
func NewGlShaderSource(
	contextID atom.ContextID,
	pShader ShaderId,
	pCount int32,
	pSource StringArray,
	pLength S32Array,
) *GlShaderSource {
	return &GlShaderSource{
		Context: contextID,
		In:      GlShaderSource_In{Shader: pShader, Count: pCount, Source: pSource, Length: pLength},
		Out:     GlShaderSource_Out{},
	}
}
func NewGlShaderBinary(
	contextID atom.ContextID,
	pCount int32,
	pShaders ShaderIdArray,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
	pBinarySize int32,
) *GlShaderBinary {
	return &GlShaderBinary{
		Context: contextID,
		In:      GlShaderBinary_In{Count: pCount, Shaders: pShaders, BinaryFormat: pBinaryFormat, Binary: pBinary, BinarySize: pBinarySize},
		Out:     GlShaderBinary_Out{},
	}
}
func NewGlGetShaderInfoLog(
	contextID atom.ContextID,
	pShader ShaderId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pInfo string,
) *GlGetShaderInfoLog {
	return &GlGetShaderInfoLog{
		Context: contextID,
		In:      GlGetShaderInfoLog_In{Shader: pShader, BufferLength: pBufferLength},
		Out:     GlGetShaderInfoLog_Out{StringLengthWritten: pStringLengthWritten, Info: pInfo},
	}
}
func NewGlGetShaderSource(
	contextID atom.ContextID,
	pShader ShaderId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pSource string,
) *GlGetShaderSource {
	return &GlGetShaderSource{
		Context: contextID,
		In:      GlGetShaderSource_In{Shader: pShader, BufferLength: pBufferLength},
		Out:     GlGetShaderSource_Out{StringLengthWritten: pStringLengthWritten, Source: pSource},
	}
}
func NewGlReleaseShaderCompiler(
	contextID atom.ContextID,
) *GlReleaseShaderCompiler {
	return &GlReleaseShaderCompiler{
		Context: contextID,
		In:      GlReleaseShaderCompiler_In{},
		Out:     GlReleaseShaderCompiler_Out{},
	}
}
func NewGlCompileShader(
	contextID atom.ContextID,
	pShader ShaderId,
) *GlCompileShader {
	return &GlCompileShader{
		Context: contextID,
		In:      GlCompileShader_In{Shader: pShader},
		Out:     GlCompileShader_Out{},
	}
}
func NewGlIsShader(
	contextID atom.ContextID,
	pShader ShaderId,
	pResult bool,
) *GlIsShader {
	return &GlIsShader{
		Context: contextID,
		In:      GlIsShader_In{Shader: pShader},
		Out:     GlIsShader_Out{Result: pResult},
	}
}
func NewGlCreateProgram(
	contextID atom.ContextID,
	pResult ProgramId,
) *GlCreateProgram {
	return &GlCreateProgram{
		Context: contextID,
		In:      GlCreateProgram_In{},
		Out:     GlCreateProgram_Out{Result: pResult},
	}
}
func NewGlDeleteProgram(
	contextID atom.ContextID,
	pProgram ProgramId,
) *GlDeleteProgram {
	return &GlDeleteProgram{
		Context: contextID,
		In:      GlDeleteProgram_In{Program: pProgram},
		Out:     GlDeleteProgram_Out{},
	}
}
func NewGlAttachShader(
	contextID atom.ContextID,
	pProgram ProgramId,
	pShader ShaderId,
) *GlAttachShader {
	return &GlAttachShader{
		Context: contextID,
		In:      GlAttachShader_In{Program: pProgram, Shader: pShader},
		Out:     GlAttachShader_Out{},
	}
}
func NewGlDetachShader(
	contextID atom.ContextID,
	pProgram ProgramId,
	pShader ShaderId,
) *GlDetachShader {
	return &GlDetachShader{
		Context: contextID,
		In:      GlDetachShader_In{Program: pProgram, Shader: pShader},
		Out:     GlDetachShader_Out{},
	}
}
func NewGlGetAttachedShaders(
	contextID atom.ContextID,
	pProgram ProgramId,
	pBufferLength int32,
	pShadersLengthWritten int32,
	pShaders ShaderIdArray,
) *GlGetAttachedShaders {
	return &GlGetAttachedShaders{
		Context: contextID,
		In:      GlGetAttachedShaders_In{Program: pProgram, BufferLength: pBufferLength},
		Out:     GlGetAttachedShaders_Out{ShadersLengthWritten: pShadersLengthWritten, Shaders: pShaders},
	}
}
func NewGlLinkProgram(
	contextID atom.ContextID,
	pProgram ProgramId,
) *GlLinkProgram {
	return &GlLinkProgram{
		Context: contextID,
		In:      GlLinkProgram_In{Program: pProgram},
		Out:     GlLinkProgram_Out{},
	}
}
func NewGlGetProgramInfoLog(
	contextID atom.ContextID,
	pProgram ProgramId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pInfo string,
) *GlGetProgramInfoLog {
	return &GlGetProgramInfoLog{
		Context: contextID,
		In:      GlGetProgramInfoLog_In{Program: pProgram, BufferLength: pBufferLength},
		Out:     GlGetProgramInfoLog_Out{StringLengthWritten: pStringLengthWritten, Info: pInfo},
	}
}
func NewGlUseProgram(
	contextID atom.ContextID,
	pProgram ProgramId,
) *GlUseProgram {
	return &GlUseProgram{
		Context: contextID,
		In:      GlUseProgram_In{Program: pProgram},
		Out:     GlUseProgram_Out{},
	}
}
func NewGlIsProgram(
	contextID atom.ContextID,
	pProgram ProgramId,
	pResult bool,
) *GlIsProgram {
	return &GlIsProgram{
		Context: contextID,
		In:      GlIsProgram_In{Program: pProgram},
		Out:     GlIsProgram_Out{Result: pResult},
	}
}
func NewGlValidateProgram(
	contextID atom.ContextID,
	pProgram ProgramId,
) *GlValidateProgram {
	return &GlValidateProgram{
		Context: contextID,
		In:      GlValidateProgram_In{Program: pProgram},
		Out:     GlValidateProgram_Out{},
	}
}
func NewGlClearColor(
	contextID atom.ContextID,
	pR float32,
	pG float32,
	pB float32,
	pA float32,
) *GlClearColor {
	return &GlClearColor{
		Context: contextID,
		In:      GlClearColor_In{R: pR, G: pG, B: pB, A: pA},
		Out:     GlClearColor_Out{},
	}
}
func NewGlClearDepthf(
	contextID atom.ContextID,
	pDepth float32,
) *GlClearDepthf {
	return &GlClearDepthf{
		Context: contextID,
		In:      GlClearDepthf_In{Depth: pDepth},
		Out:     GlClearDepthf_Out{},
	}
}
func NewGlClearStencil(
	contextID atom.ContextID,
	pStencil int32,
) *GlClearStencil {
	return &GlClearStencil{
		Context: contextID,
		In:      GlClearStencil_In{Stencil: pStencil},
		Out:     GlClearStencil_Out{},
	}
}
func NewGlClear(
	contextID atom.ContextID,
	pMask ClearMask,
) *GlClear {
	return &GlClear{
		Context: contextID,
		In:      GlClear_In{Mask: pMask},
		Out:     GlClear_Out{},
	}
}
func NewGlCullFace(
	contextID atom.ContextID,
	pMode FaceMode,
) *GlCullFace {
	return &GlCullFace{
		Context: contextID,
		In:      GlCullFace_In{Mode: pMode},
		Out:     GlCullFace_Out{},
	}
}
func NewGlPolygonOffset(
	contextID atom.ContextID,
	pScaleFactor float32,
	pUnits float32,
) *GlPolygonOffset {
	return &GlPolygonOffset{
		Context: contextID,
		In:      GlPolygonOffset_In{ScaleFactor: pScaleFactor, Units: pUnits},
		Out:     GlPolygonOffset_Out{},
	}
}
func NewGlLineWidth(
	contextID atom.ContextID,
	pWidth float32,
) *GlLineWidth {
	return &GlLineWidth{
		Context: contextID,
		In:      GlLineWidth_In{Width: pWidth},
		Out:     GlLineWidth_Out{},
	}
}
func NewGlSampleCoverage(
	contextID atom.ContextID,
	pValue float32,
	pInvert bool,
) *GlSampleCoverage {
	return &GlSampleCoverage{
		Context: contextID,
		In:      GlSampleCoverage_In{Value: pValue, Invert: pInvert},
		Out:     GlSampleCoverage_Out{},
	}
}
func NewGlHint(
	contextID atom.ContextID,
	pTarget HintTarget,
	pMode HintMode,
) *GlHint {
	return &GlHint{
		Context: contextID,
		In:      GlHint_In{Target: pTarget, Mode: pMode},
		Out:     GlHint_Out{},
	}
}
func NewGlFramebufferRenderbuffer(
	contextID atom.ContextID,
	pFramebufferTarget FramebufferTarget,
	pFramebufferAttachment FramebufferAttachment,
	pRenderbufferTarget RenderbufferTarget,
	pRenderbuffer RenderbufferId,
) *GlFramebufferRenderbuffer {
	return &GlFramebufferRenderbuffer{
		Context: contextID,
		In:      GlFramebufferRenderbuffer_In{FramebufferTarget: pFramebufferTarget, FramebufferAttachment: pFramebufferAttachment, RenderbufferTarget: pRenderbufferTarget, Renderbuffer: pRenderbuffer},
		Out:     GlFramebufferRenderbuffer_Out{},
	}
}
func NewGlFramebufferTexture2D(
	contextID atom.ContextID,
	pFramebufferTarget FramebufferTarget,
	pFramebufferAttachment FramebufferAttachment,
	pTextureTarget TextureImageTarget,
	pTexture TextureId,
	pLevel int32,
) *GlFramebufferTexture2D {
	return &GlFramebufferTexture2D{
		Context: contextID,
		In:      GlFramebufferTexture2D_In{FramebufferTarget: pFramebufferTarget, FramebufferAttachment: pFramebufferAttachment, TextureTarget: pTextureTarget, Texture: pTexture, Level: pLevel},
		Out:     GlFramebufferTexture2D_Out{},
	}
}
func NewGlGetFramebufferAttachmentParameteriv(
	contextID atom.ContextID,
	pFramebufferTarget FramebufferTarget,
	pAttachment FramebufferAttachment,
	pParameter FramebufferAttachmentParameter,
	pValue S32Array,
) *GlGetFramebufferAttachmentParameteriv {
	return &GlGetFramebufferAttachmentParameteriv{
		Context: contextID,
		In:      GlGetFramebufferAttachmentParameteriv_In{FramebufferTarget: pFramebufferTarget, Attachment: pAttachment, Parameter: pParameter},
		Out:     GlGetFramebufferAttachmentParameteriv_Out{Value: pValue},
	}
}
func NewGlDrawElements(
	contextID atom.ContextID,
	pDrawMode DrawMode,
	pElementCount int32,
	pIndicesType IndicesType,
	pIndices IndicesPointer,
) *GlDrawElements {
	return &GlDrawElements{
		Context: contextID,
		In:      GlDrawElements_In{DrawMode: pDrawMode, ElementCount: pElementCount, IndicesType: pIndicesType, Indices: pIndices},
		Out:     GlDrawElements_Out{},
	}
}
func NewGlDrawArrays(
	contextID atom.ContextID,
	pDrawMode DrawMode,
	pFirstIndex int32,
	pIndexCount int32,
) *GlDrawArrays {
	return &GlDrawArrays{
		Context: contextID,
		In:      GlDrawArrays_In{DrawMode: pDrawMode, FirstIndex: pFirstIndex, IndexCount: pIndexCount},
		Out:     GlDrawArrays_Out{},
	}
}
func NewGlFlush(
	contextID atom.ContextID,
) *GlFlush {
	return &GlFlush{
		Context: contextID,
		In:      GlFlush_In{},
		Out:     GlFlush_Out{},
	}
}
func NewGlFinish(
	contextID atom.ContextID,
) *GlFinish {
	return &GlFinish{
		Context: contextID,
		In:      GlFinish_In{},
		Out:     GlFinish_Out{},
	}
}
func NewGlGetBooleanv(
	contextID atom.ContextID,
	pParam StateVariable,
	pValues BoolArray,
) *GlGetBooleanv {
	return &GlGetBooleanv{
		Context: contextID,
		In:      GlGetBooleanv_In{Param: pParam},
		Out:     GlGetBooleanv_Out{Values: pValues},
	}
}
func NewGlGetFloatv(
	contextID atom.ContextID,
	pParam StateVariable,
	pValues F32Array,
) *GlGetFloatv {
	return &GlGetFloatv{
		Context: contextID,
		In:      GlGetFloatv_In{Param: pParam},
		Out:     GlGetFloatv_Out{Values: pValues},
	}
}
func NewGlGetIntegerv(
	contextID atom.ContextID,
	pParam StateVariable,
	pValues S32Array,
) *GlGetIntegerv {
	return &GlGetIntegerv{
		Context: contextID,
		In:      GlGetIntegerv_In{Param: pParam},
		Out:     GlGetIntegerv_Out{Values: pValues},
	}
}
func NewGlGetString(
	contextID atom.ContextID,
	pParam StringConstant,
	pResult string,
) *GlGetString {
	return &GlGetString{
		Context: contextID,
		In:      GlGetString_In{Param: pParam},
		Out:     GlGetString_Out{Result: pResult},
	}
}
func NewGlEnable(
	contextID atom.ContextID,
	pCapability Capability,
) *GlEnable {
	return &GlEnable{
		Context: contextID,
		In:      GlEnable_In{Capability: pCapability},
		Out:     GlEnable_Out{},
	}
}
func NewGlDisable(
	contextID atom.ContextID,
	pCapability Capability,
) *GlDisable {
	return &GlDisable{
		Context: contextID,
		In:      GlDisable_In{Capability: pCapability},
		Out:     GlDisable_Out{},
	}
}
func NewGlIsEnabled(
	contextID atom.ContextID,
	pCapability Capability,
	pResult bool,
) *GlIsEnabled {
	return &GlIsEnabled{
		Context: contextID,
		In:      GlIsEnabled_In{Capability: pCapability},
		Out:     GlIsEnabled_Out{Result: pResult},
	}
}
func NewGlMapBufferRange(
	contextID atom.ContextID,
	pTarget MapBufferTarget,
	pOffset int32,
	pLength int32,
	pAccess MapBufferRangeAccess,
	pResult memory.Pointer,
) *GlMapBufferRange {
	return &GlMapBufferRange{
		Context: contextID,
		In:      GlMapBufferRange_In{Target: pTarget, Offset: pOffset, Length: pLength, Access: pAccess},
		Out:     GlMapBufferRange_Out{Result: pResult},
	}
}
func NewGlUnmapBuffer(
	contextID atom.ContextID,
	pTarget MapBufferTarget,
) *GlUnmapBuffer {
	return &GlUnmapBuffer{
		Context: contextID,
		In:      GlUnmapBuffer_In{Target: pTarget},
		Out:     GlUnmapBuffer_Out{},
	}
}
func NewGlInvalidateFramebuffer(
	contextID atom.ContextID,
	pTarget FramebufferTarget,
	pCount int32,
	pAttachments FramebufferAttachmentArray,
) *GlInvalidateFramebuffer {
	return &GlInvalidateFramebuffer{
		Context: contextID,
		In:      GlInvalidateFramebuffer_In{Target: pTarget, Count: pCount, Attachments: pAttachments},
		Out:     GlInvalidateFramebuffer_Out{},
	}
}
func NewGlRenderbufferStorageMultisample(
	contextID atom.ContextID,
	pTarget RenderbufferTarget,
	pSamples int32,
	pFormat RenderbufferFormat,
	pWidth int32,
	pHeight int32,
) *GlRenderbufferStorageMultisample {
	return &GlRenderbufferStorageMultisample{
		Context: contextID,
		In:      GlRenderbufferStorageMultisample_In{Target: pTarget, Samples: pSamples, Format: pFormat, Width: pWidth, Height: pHeight},
		Out:     GlRenderbufferStorageMultisample_Out{},
	}
}
func NewGlBlitFramebuffer(
	contextID atom.ContextID,
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
		Context: contextID,
		In:      GlBlitFramebuffer_In{SrcX0: pSrcX0, SrcY0: pSrcY0, SrcX1: pSrcX1, SrcY1: pSrcY1, DstX0: pDstX0, DstY0: pDstY0, DstX1: pDstX1, DstY1: pDstY1, Mask: pMask, Filter: pFilter},
		Out:     GlBlitFramebuffer_Out{},
	}
}
func NewGlGenQueries(
	contextID atom.ContextID,
	pCount int32,
	pQueries QueryIdArray,
) *GlGenQueries {
	return &GlGenQueries{
		Context: contextID,
		In:      GlGenQueries_In{Count: pCount},
		Out:     GlGenQueries_Out{Queries: pQueries},
	}
}
func NewGlBeginQuery(
	contextID atom.ContextID,
	pTarget QueryTarget,
	pQuery QueryId,
) *GlBeginQuery {
	return &GlBeginQuery{
		Context: contextID,
		In:      GlBeginQuery_In{Target: pTarget, Query: pQuery},
		Out:     GlBeginQuery_Out{},
	}
}
func NewGlEndQuery(
	contextID atom.ContextID,
	pTarget QueryTarget,
) *GlEndQuery {
	return &GlEndQuery{
		Context: contextID,
		In:      GlEndQuery_In{Target: pTarget},
		Out:     GlEndQuery_Out{},
	}
}
func NewGlDeleteQueries(
	contextID atom.ContextID,
	pCount int32,
	pQueries QueryIdArray,
) *GlDeleteQueries {
	return &GlDeleteQueries{
		Context: contextID,
		In:      GlDeleteQueries_In{Count: pCount, Queries: pQueries},
		Out:     GlDeleteQueries_Out{},
	}
}
func NewGlIsQuery(
	contextID atom.ContextID,
	pQuery QueryId,
	pResult bool,
) *GlIsQuery {
	return &GlIsQuery{
		Context: contextID,
		In:      GlIsQuery_In{Query: pQuery},
		Out:     GlIsQuery_Out{Result: pResult},
	}
}
func NewGlGetQueryiv(
	contextID atom.ContextID,
	pTarget QueryTarget,
	pParameter QueryParameter,
	pValue int32,
) *GlGetQueryiv {
	return &GlGetQueryiv{
		Context: contextID,
		In:      GlGetQueryiv_In{Target: pTarget, Parameter: pParameter},
		Out:     GlGetQueryiv_Out{Value: pValue},
	}
}
func NewGlGetQueryObjectuiv(
	contextID atom.ContextID,
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint32,
) *GlGetQueryObjectuiv {
	return &GlGetQueryObjectuiv{
		Context: contextID,
		In:      GlGetQueryObjectuiv_In{Query: pQuery, Parameter: pParameter},
		Out:     GlGetQueryObjectuiv_Out{Value: pValue},
	}
}
func NewGlGenQueriesEXT(
	contextID atom.ContextID,
	pCount int32,
	pQueries QueryIdArray,
) *GlGenQueriesEXT {
	return &GlGenQueriesEXT{
		Context: contextID,
		In:      GlGenQueriesEXT_In{Count: pCount},
		Out:     GlGenQueriesEXT_Out{Queries: pQueries},
	}
}
func NewGlBeginQueryEXT(
	contextID atom.ContextID,
	pTarget QueryTarget,
	pQuery QueryId,
) *GlBeginQueryEXT {
	return &GlBeginQueryEXT{
		Context: contextID,
		In:      GlBeginQueryEXT_In{Target: pTarget, Query: pQuery},
		Out:     GlBeginQueryEXT_Out{},
	}
}
func NewGlEndQueryEXT(
	contextID atom.ContextID,
	pTarget QueryTarget,
) *GlEndQueryEXT {
	return &GlEndQueryEXT{
		Context: contextID,
		In:      GlEndQueryEXT_In{Target: pTarget},
		Out:     GlEndQueryEXT_Out{},
	}
}
func NewGlDeleteQueriesEXT(
	contextID atom.ContextID,
	pCount int32,
	pQueries QueryIdArray,
) *GlDeleteQueriesEXT {
	return &GlDeleteQueriesEXT{
		Context: contextID,
		In:      GlDeleteQueriesEXT_In{Count: pCount, Queries: pQueries},
		Out:     GlDeleteQueriesEXT_Out{},
	}
}
func NewGlIsQueryEXT(
	contextID atom.ContextID,
	pQuery QueryId,
	pResult bool,
) *GlIsQueryEXT {
	return &GlIsQueryEXT{
		Context: contextID,
		In:      GlIsQueryEXT_In{Query: pQuery},
		Out:     GlIsQueryEXT_Out{Result: pResult},
	}
}
func NewGlQueryCounterEXT(
	contextID atom.ContextID,
	pQuery QueryId,
	pTarget QueryTarget,
) *GlQueryCounterEXT {
	return &GlQueryCounterEXT{
		Context: contextID,
		In:      GlQueryCounterEXT_In{Query: pQuery, Target: pTarget},
		Out:     GlQueryCounterEXT_Out{},
	}
}
func NewGlGetQueryivEXT(
	contextID atom.ContextID,
	pTarget QueryTarget,
	pParameter QueryParameter,
	pValue int32,
) *GlGetQueryivEXT {
	return &GlGetQueryivEXT{
		Context: contextID,
		In:      GlGetQueryivEXT_In{Target: pTarget, Parameter: pParameter},
		Out:     GlGetQueryivEXT_Out{Value: pValue},
	}
}
func NewGlGetQueryObjectivEXT(
	contextID atom.ContextID,
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue int32,
) *GlGetQueryObjectivEXT {
	return &GlGetQueryObjectivEXT{
		Context: contextID,
		In:      GlGetQueryObjectivEXT_In{Query: pQuery, Parameter: pParameter},
		Out:     GlGetQueryObjectivEXT_Out{Value: pValue},
	}
}
func NewGlGetQueryObjectuivEXT(
	contextID atom.ContextID,
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint32,
) *GlGetQueryObjectuivEXT {
	return &GlGetQueryObjectuivEXT{
		Context: contextID,
		In:      GlGetQueryObjectuivEXT_In{Query: pQuery, Parameter: pParameter},
		Out:     GlGetQueryObjectuivEXT_Out{Value: pValue},
	}
}
func NewGlGetQueryObjecti64vEXT(
	contextID atom.ContextID,
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue int64,
) *GlGetQueryObjecti64vEXT {
	return &GlGetQueryObjecti64vEXT{
		Context: contextID,
		In:      GlGetQueryObjecti64vEXT_In{Query: pQuery, Parameter: pParameter},
		Out:     GlGetQueryObjecti64vEXT_Out{Value: pValue},
	}
}
func NewGlGetQueryObjectui64vEXT(
	contextID atom.ContextID,
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint64,
) *GlGetQueryObjectui64vEXT {
	return &GlGetQueryObjectui64vEXT{
		Context: contextID,
		In:      GlGetQueryObjectui64vEXT_In{Query: pQuery, Parameter: pParameter},
		Out:     GlGetQueryObjectui64vEXT_Out{Value: pValue},
	}
}

////////////////////////////////////////////////////////////////////////////////
// API
////////////////////////////////////////////////////////////////////////////////
type api struct{}

func (api) Name() string {
	return "gles"
}
func (api) ID() service.ApiId {
	return service.ApiId{ID: binary.NewID([]byte("gles"))}
}
func API() gfxapi.API {
	return api{}
}
func init() {
	gfxapi.Register(API())
	atom.Register(atom.TypeInfo{
		Name: "Init",
		Docs: "[]",
		ID:   0,
		New:  func() atom.Atom { return &Init{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "StartTimer",
		Docs: "[]",
		ID:   1,
		New:  func() atom.Atom { return &StartTimer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "StopTimer",
		Docs: "[]",
		ID:   2,
		New:  func() atom.Atom { return &StopTimer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "FlushPostBuffer",
		Docs: "[]",
		ID:   3,
		New:  func() atom.Atom { return &FlushPostBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglCreateContext",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglCreateContext.xhtml]",
		ID:   4,
		New:  func() atom.Atom { return &EglCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglMakeCurrent",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglMakeCurrent.xhtml]",
		ID:   5,
		New:  func() atom.Atom { return &EglMakeCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "EglSwapBuffers",
		Docs: "[http://www.khronos.org/registry/egl/sdk/docs/man/html/eglSwapBuffers.xhtml]",
		ID:   6,
		New:  func() atom.Atom { return &EglSwapBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "WglCreateContext",
		Docs: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374379%28v=vs.85%29.aspx]",
		ID:   7,
		New:  func() atom.Atom { return &WglCreateContext{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "WglMakeCurrent",
		Docs: "[http://msdn.microsoft.com/en-us/library/windows/desktop/dd374387%28v=vs.85%29.aspx]",
		ID:   8,
		New:  func() atom.Atom { return &WglMakeCurrent{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "WglSwapBuffers",
		Docs: "[http://msdn.microsoft.com/en-us/library/dd369060%28v=vs.85%29]",
		ID:   9,
		New:  func() atom.Atom { return &WglSwapBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEnableClientState",
		Docs: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
		ID:   10,
		New:  func() atom.Atom { return &GlEnableClientState{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDisableClientState",
		Docs: "[http://www.khronos.org/opengles/sdk/1.1/docs/man/glEnableClientState.xml]",
		ID:   11,
		New:  func() atom.Atom { return &GlDisableClientState{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetProgramBinaryOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
		ID:   12,
		New:  func() atom.Atom { return &GlGetProgramBinaryOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlProgramBinaryOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_get_program_binary.txt]",
		ID:   13,
		New:  func() atom.Atom { return &GlProgramBinaryOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStartTilingQCOM",
		Docs: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
		ID:   14,
		New:  func() atom.Atom { return &GlStartTilingQCOM{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEndTilingQCOM",
		Docs: "[http://www.khronos.org/registry/gles/extensions/QCOM/QCOM_tiled_rendering.txt]",
		ID:   15,
		New:  func() atom.Atom { return &GlEndTilingQCOM{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDiscardFramebufferEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_discard_framebuffer.txt]",
		ID:   16,
		New:  func() atom.Atom { return &GlDiscardFramebufferEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlInsertEventMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   17,
		New:  func() atom.Atom { return &GlInsertEventMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPushGroupMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   18,
		New:  func() atom.Atom { return &GlPushGroupMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPopGroupMarkerEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_debug_marker.txt]",
		ID:   19,
		New:  func() atom.Atom { return &GlPopGroupMarkerEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexStorage1DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   20,
		New:  func() atom.Atom { return &GlTexStorage1DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexStorage2DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   21,
		New:  func() atom.Atom { return &GlTexStorage2DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexStorage3DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   22,
		New:  func() atom.Atom { return &GlTexStorage3DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTextureStorage1DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   23,
		New:  func() atom.Atom { return &GlTextureStorage1DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTextureStorage2DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   24,
		New:  func() atom.Atom { return &GlTextureStorage2DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTextureStorage3DEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_texture_storage.txt]",
		ID:   25,
		New:  func() atom.Atom { return &GlTextureStorage3DEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenVertexArraysOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   26,
		New:  func() atom.Atom { return &GlGenVertexArraysOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindVertexArrayOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   27,
		New:  func() atom.Atom { return &GlBindVertexArrayOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteVertexArraysOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   28,
		New:  func() atom.Atom { return &GlDeleteVertexArraysOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsVertexArrayOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_vertex_array_object.txt]",
		ID:   29,
		New:  func() atom.Atom { return &GlIsVertexArrayOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEGLImageTargetTexture2DOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
		ID:   30,
		New:  func() atom.Atom { return &GlEGLImageTargetTexture2DOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEGLImageTargetRenderbufferStorageOES",
		Docs: "[http://www.khronos.org/registry/gles/extensions/OES/OES_EGL_image.txt]",
		ID:   31,
		New:  func() atom.Atom { return &GlEGLImageTargetRenderbufferStorageOES{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetGraphicsResetStatusEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_robustness.txt]",
		ID:   32,
		New:  func() atom.Atom { return &GlGetGraphicsResetStatusEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindAttribLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindAttribLocation.xml]",
		ID:   33,
		New:  func() atom.Atom { return &GlBindAttribLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendFunc",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFunc.xml]",
		ID:   34,
		New:  func() atom.Atom { return &GlBlendFunc{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendFuncSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendFuncSeparate.xml]",
		ID:   35,
		New:  func() atom.Atom { return &GlBlendFuncSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendEquation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquation.xml]",
		ID:   36,
		New:  func() atom.Atom { return &GlBlendEquation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendEquationSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendEquationSeparate.xml]",
		ID:   37,
		New:  func() atom.Atom { return &GlBlendEquationSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlendColor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBlendColor.xml]",
		ID:   38,
		New:  func() atom.Atom { return &GlBlendColor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEnableVertexAttribArray",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnableVertexAttribArray.xml]",
		ID:   39,
		New:  func() atom.Atom { return &GlEnableVertexAttribArray{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDisableVertexAttribArray",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisableVertexAttribArray.xml]",
		ID:   40,
		New:  func() atom.Atom { return &GlDisableVertexAttribArray{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttribPointer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttribPointer.xml]",
		ID:   41,
		New:  func() atom.Atom { return &GlVertexAttribPointer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetActiveAttrib",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveAttrib.xml]",
		ID:   42,
		New:  func() atom.Atom { return &GlGetActiveAttrib{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetActiveUniform",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetActiveUniform.xml]",
		ID:   43,
		New:  func() atom.Atom { return &GlGetActiveUniform{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetError",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetError.xml]",
		ID:   44,
		New:  func() atom.Atom { return &GlGetError{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetProgramiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgram.xml]",
		ID:   45,
		New:  func() atom.Atom { return &GlGetProgramiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderiv.xml]",
		ID:   46,
		New:  func() atom.Atom { return &GlGetShaderiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetUniformLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniformLocation.xml]",
		ID:   47,
		New:  func() atom.Atom { return &GlGetUniformLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetAttribLocation",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttribLocation.xml]",
		ID:   48,
		New:  func() atom.Atom { return &GlGetAttribLocation{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPixelStorei",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPixelStorei.xml]",
		ID:   49,
		New:  func() atom.Atom { return &GlPixelStorei{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexParameteri",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
		ID:   50,
		New:  func() atom.Atom { return &GlTexParameteri{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexParameterf",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexParameter.xml]",
		ID:   51,
		New:  func() atom.Atom { return &GlTexParameterf{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetTexParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
		ID:   52,
		New:  func() atom.Atom { return &GlGetTexParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetTexParameterfv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetTexParameter.xml]",
		ID:   53,
		New:  func() atom.Atom { return &GlGetTexParameterfv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   54,
		New:  func() atom.Atom { return &GlUniform1i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   55,
		New:  func() atom.Atom { return &GlUniform2i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   56,
		New:  func() atom.Atom { return &GlUniform3i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4i",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   57,
		New:  func() atom.Atom { return &GlUniform4i{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   58,
		New:  func() atom.Atom { return &GlUniform1iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   59,
		New:  func() atom.Atom { return &GlUniform2iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   60,
		New:  func() atom.Atom { return &GlUniform3iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4iv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   61,
		New:  func() atom.Atom { return &GlUniform4iv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   62,
		New:  func() atom.Atom { return &GlUniform1f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   63,
		New:  func() atom.Atom { return &GlUniform2f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   64,
		New:  func() atom.Atom { return &GlUniform3f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   65,
		New:  func() atom.Atom { return &GlUniform4f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform1fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   66,
		New:  func() atom.Atom { return &GlUniform1fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   67,
		New:  func() atom.Atom { return &GlUniform2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   68,
		New:  func() atom.Atom { return &GlUniform3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniform4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   69,
		New:  func() atom.Atom { return &GlUniform4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniformMatrix2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   70,
		New:  func() atom.Atom { return &GlUniformMatrix2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniformMatrix3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   71,
		New:  func() atom.Atom { return &GlUniformMatrix3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUniformMatrix4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUniform.xml]",
		ID:   72,
		New:  func() atom.Atom { return &GlUniformMatrix4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetUniformfv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
		ID:   73,
		New:  func() atom.Atom { return &GlGetUniformfv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetUniformiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetUniform.xml]",
		ID:   74,
		New:  func() atom.Atom { return &GlGetUniformiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib1f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   75,
		New:  func() atom.Atom { return &GlVertexAttrib1f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib2f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   76,
		New:  func() atom.Atom { return &GlVertexAttrib2f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib3f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   77,
		New:  func() atom.Atom { return &GlVertexAttrib3f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib4f",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   78,
		New:  func() atom.Atom { return &GlVertexAttrib4f{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib1fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   79,
		New:  func() atom.Atom { return &GlVertexAttrib1fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib2fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   80,
		New:  func() atom.Atom { return &GlVertexAttrib2fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib3fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   81,
		New:  func() atom.Atom { return &GlVertexAttrib3fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlVertexAttrib4fv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glVertexAttrib.xml]",
		ID:   82,
		New:  func() atom.Atom { return &GlVertexAttrib4fv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderPrecisionFormat",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderPrecisionFormat.xml]",
		ID:   83,
		New:  func() atom.Atom { return &GlGetShaderPrecisionFormat{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDepthMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthMask.xml]",
		ID:   84,
		New:  func() atom.Atom { return &GlDepthMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDepthFunc",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthFunc.xml]",
		ID:   85,
		New:  func() atom.Atom { return &GlDepthFunc{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDepthRangef",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDepthRangef.xml]",
		ID:   86,
		New:  func() atom.Atom { return &GlDepthRangef{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlColorMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glColorMask.xml]",
		ID:   87,
		New:  func() atom.Atom { return &GlColorMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilMask",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMask.xml]",
		ID:   88,
		New:  func() atom.Atom { return &GlStencilMask{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilMaskSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilMaskSeparate.xml]",
		ID:   89,
		New:  func() atom.Atom { return &GlStencilMaskSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilFuncSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilFuncSeparate.xml]",
		ID:   90,
		New:  func() atom.Atom { return &GlStencilFuncSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlStencilOpSeparate",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glStencilOpSeparate.xml]",
		ID:   91,
		New:  func() atom.Atom { return &GlStencilOpSeparate{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFrontFace",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFrontFace.xml]",
		ID:   92,
		New:  func() atom.Atom { return &GlFrontFace{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlViewport",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glViewport.xml]",
		ID:   93,
		New:  func() atom.Atom { return &GlViewport{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlScissor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glScissor.xml]",
		ID:   94,
		New:  func() atom.Atom { return &GlScissor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlActiveTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glActiveTexture.xml]",
		ID:   95,
		New:  func() atom.Atom { return &GlActiveTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenTextures",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenTextures.xml]",
		ID:   96,
		New:  func() atom.Atom { return &GlGenTextures{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteTextures",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteTextures.xml]",
		ID:   97,
		New:  func() atom.Atom { return &GlDeleteTextures{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsTexture.xml]",
		ID:   98,
		New:  func() atom.Atom { return &GlIsTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindTexture",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindTexture.xml]",
		ID:   99,
		New:  func() atom.Atom { return &GlBindTexture{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexImage2D.xml]",
		ID:   100,
		New:  func() atom.Atom { return &GlTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glTexSubImage2D.xml]",
		ID:   101,
		New:  func() atom.Atom { return &GlTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCopyTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexImage2D.xml]",
		ID:   102,
		New:  func() atom.Atom { return &GlCopyTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCopyTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCopyTexSubImage2D.xml]",
		ID:   103,
		New:  func() atom.Atom { return &GlCopyTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCompressedTexImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexImage2D.xml]",
		ID:   104,
		New:  func() atom.Atom { return &GlCompressedTexImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCompressedTexSubImage2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompressedTexSubImage2D.xml]",
		ID:   105,
		New:  func() atom.Atom { return &GlCompressedTexSubImage2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenerateMipmap",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenerateMipmap.xml]",
		ID:   106,
		New:  func() atom.Atom { return &GlGenerateMipmap{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlReadPixels",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReadPixels.xml]",
		ID:   107,
		New:  func() atom.Atom { return &GlReadPixels{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenFramebuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenFramebuffers.xml]",
		ID:   108,
		New:  func() atom.Atom { return &GlGenFramebuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindFramebuffer.xml]",
		ID:   109,
		New:  func() atom.Atom { return &GlBindFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCheckFramebufferStatus",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCheckFramebufferStatus.xml]",
		ID:   110,
		New:  func() atom.Atom { return &GlCheckFramebufferStatus{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteFramebuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteFramebuffers.xml]",
		ID:   111,
		New:  func() atom.Atom { return &GlDeleteFramebuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsFramebuffer.xml]",
		ID:   112,
		New:  func() atom.Atom { return &GlIsFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenRenderbuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenRenderbuffers.xml]",
		ID:   113,
		New:  func() atom.Atom { return &GlGenRenderbuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindRenderbuffer.xml]",
		ID:   114,
		New:  func() atom.Atom { return &GlBindRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlRenderbufferStorage",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glRenderbufferStorage.xml]",
		ID:   115,
		New:  func() atom.Atom { return &GlRenderbufferStorage{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteRenderbuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteRenderbuffers.xml]",
		ID:   116,
		New:  func() atom.Atom { return &GlDeleteRenderbuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsRenderbuffer.xml]",
		ID:   117,
		New:  func() atom.Atom { return &GlIsRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetRenderbufferParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetRenderbufferParameteriv.xml]",
		ID:   118,
		New:  func() atom.Atom { return &GlGetRenderbufferParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenBuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGenBuffers.xml]",
		ID:   119,
		New:  func() atom.Atom { return &GlGenBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBindBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBindBuffer.xml]",
		ID:   120,
		New:  func() atom.Atom { return &GlBindBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBufferData",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferData.xml]",
		ID:   121,
		New:  func() atom.Atom { return &GlBufferData{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBufferSubData",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glBufferSubData.xml]",
		ID:   122,
		New:  func() atom.Atom { return &GlBufferSubData{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteBuffers",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteBuffers.xml]",
		ID:   123,
		New:  func() atom.Atom { return &GlDeleteBuffers{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsBuffer.xml]",
		ID:   124,
		New:  func() atom.Atom { return &GlIsBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetBufferParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetBufferParameteriv.xml]",
		ID:   125,
		New:  func() atom.Atom { return &GlGetBufferParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCreateShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateShader.xml]",
		ID:   126,
		New:  func() atom.Atom { return &GlCreateShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteShader.xml]",
		ID:   127,
		New:  func() atom.Atom { return &GlDeleteShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlShaderSource",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderSource.xml]",
		ID:   128,
		New:  func() atom.Atom { return &GlShaderSource{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlShaderBinary",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glShaderBinary.xml]",
		ID:   129,
		New:  func() atom.Atom { return &GlShaderBinary{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderInfoLog",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderInfoLog.xml]",
		ID:   130,
		New:  func() atom.Atom { return &GlGetShaderInfoLog{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetShaderSource",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetShaderSource.xml]",
		ID:   131,
		New:  func() atom.Atom { return &GlGetShaderSource{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlReleaseShaderCompiler",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glReleaseShaderCompiler.xml]",
		ID:   132,
		New:  func() atom.Atom { return &GlReleaseShaderCompiler{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCompileShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCompileShader.xml]",
		ID:   133,
		New:  func() atom.Atom { return &GlCompileShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsShader.xml]",
		ID:   134,
		New:  func() atom.Atom { return &GlIsShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCreateProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCreateProgram.xml]",
		ID:   135,
		New:  func() atom.Atom { return &GlCreateProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDeleteProgram.xml]",
		ID:   136,
		New:  func() atom.Atom { return &GlDeleteProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlAttachShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glAttachShader.xml]",
		ID:   137,
		New:  func() atom.Atom { return &GlAttachShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDetachShader",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDetachShader.xml]",
		ID:   138,
		New:  func() atom.Atom { return &GlDetachShader{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetAttachedShaders",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetAttachedShaders.xml]",
		ID:   139,
		New:  func() atom.Atom { return &GlGetAttachedShaders{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlLinkProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLinkProgram.xml]",
		ID:   140,
		New:  func() atom.Atom { return &GlLinkProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetProgramInfoLog",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetProgramInfoLog.xml]",
		ID:   141,
		New:  func() atom.Atom { return &GlGetProgramInfoLog{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUseProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glUseProgram.xml]",
		ID:   142,
		New:  func() atom.Atom { return &GlUseProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsProgram.xml]",
		ID:   143,
		New:  func() atom.Atom { return &GlIsProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlValidateProgram",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glValidateProgram.xml]",
		ID:   144,
		New:  func() atom.Atom { return &GlValidateProgram{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClearColor",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearColor.xml]",
		ID:   145,
		New:  func() atom.Atom { return &GlClearColor{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClearDepthf",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearDepthf.xml]",
		ID:   146,
		New:  func() atom.Atom { return &GlClearDepthf{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClearStencil",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClearStencil.xml]",
		ID:   147,
		New:  func() atom.Atom { return &GlClearStencil{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlClear",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glClear.xml]",
		ID:   148,
		New:  func() atom.Atom { return &GlClear{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlCullFace",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glCullFace.xml]",
		ID:   149,
		New:  func() atom.Atom { return &GlCullFace{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlPolygonOffset",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glPolygonOffset.xml]",
		ID:   150,
		New:  func() atom.Atom { return &GlPolygonOffset{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlLineWidth",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glLineWidth.xml]",
		ID:   151,
		New:  func() atom.Atom { return &GlLineWidth{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlSampleCoverage",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glSampleCoverage.xml]",
		ID:   152,
		New:  func() atom.Atom { return &GlSampleCoverage{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlHint",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glHint.xml]",
		ID:   153,
		New:  func() atom.Atom { return &GlHint{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFramebufferRenderbuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferRenderbuffer.xml]",
		ID:   154,
		New:  func() atom.Atom { return &GlFramebufferRenderbuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFramebufferTexture2D",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFramebufferTexture2D.xml]",
		ID:   155,
		New:  func() atom.Atom { return &GlFramebufferTexture2D{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetFramebufferAttachmentParameteriv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetFramebufferAttachmentParameteriv.xml]",
		ID:   156,
		New:  func() atom.Atom { return &GlGetFramebufferAttachmentParameteriv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDrawElements",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawElements.xml]",
		ID:   157,
		New:  func() atom.Atom { return &GlDrawElements{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDrawArrays",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDrawArrays.xml]",
		ID:   158,
		New:  func() atom.Atom { return &GlDrawArrays{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFlush",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFlush.xml]",
		ID:   159,
		New:  func() atom.Atom { return &GlFlush{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlFinish",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glFinish.xml]",
		ID:   160,
		New:  func() atom.Atom { return &GlFinish{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetBooleanv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   161,
		New:  func() atom.Atom { return &GlGetBooleanv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetFloatv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   162,
		New:  func() atom.Atom { return &GlGetFloatv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetIntegerv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGet.xml]",
		ID:   163,
		New:  func() atom.Atom { return &GlGetIntegerv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetString",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glGetString.xml]",
		ID:   164,
		New:  func() atom.Atom { return &GlGetString{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEnable",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glEnable.xml]",
		ID:   165,
		New:  func() atom.Atom { return &GlEnable{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDisable",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glDisable.xml]",
		ID:   166,
		New:  func() atom.Atom { return &GlDisable{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsEnabled",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man/xhtml/glIsEnabled.xml]",
		ID:   167,
		New:  func() atom.Atom { return &GlIsEnabled{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlMapBufferRange",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
		ID:   168,
		New:  func() atom.Atom { return &GlMapBufferRange{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlUnmapBuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glMapBufferRange.xhtml]",
		ID:   169,
		New:  func() atom.Atom { return &GlUnmapBuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlInvalidateFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glInvalidateFramebuffer.xhtml]",
		ID:   170,
		New:  func() atom.Atom { return &GlInvalidateFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlRenderbufferStorageMultisample",
		Docs: "[http://www.opengl.org/registry/specs/EXT/framebuffer_multisample.txt]",
		ID:   171,
		New:  func() atom.Atom { return &GlRenderbufferStorageMultisample{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBlitFramebuffer",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBlitFramebuffer.xhtml]",
		ID:   172,
		New:  func() atom.Atom { return &GlBlitFramebuffer{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenQueries",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGenQueries.xhtml]",
		ID:   173,
		New:  func() atom.Atom { return &GlGenQueries{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBeginQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glBeginQuery.xhtml]",
		ID:   174,
		New:  func() atom.Atom { return &GlBeginQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEndQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glEndQuery.xhtml]",
		ID:   175,
		New:  func() atom.Atom { return &GlEndQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteQueries",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glDeleteQueries.xhtml]",
		ID:   176,
		New:  func() atom.Atom { return &GlDeleteQueries{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsQuery",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glIsQuery.xhtml]",
		ID:   177,
		New:  func() atom.Atom { return &GlIsQuery{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryiv.xhtml]",
		ID:   178,
		New:  func() atom.Atom { return &GlGetQueryiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectuiv",
		Docs: "[http://www.khronos.org/opengles/sdk/docs/man3/html/glGetQueryObjectuiv.xhtml]",
		ID:   179,
		New:  func() atom.Atom { return &GlGetQueryObjectuiv{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGenQueriesEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   180,
		New:  func() atom.Atom { return &GlGenQueriesEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlBeginQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   181,
		New:  func() atom.Atom { return &GlBeginQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlEndQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   182,
		New:  func() atom.Atom { return &GlEndQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlDeleteQueriesEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   183,
		New:  func() atom.Atom { return &GlDeleteQueriesEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlIsQueryEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   184,
		New:  func() atom.Atom { return &GlIsQueryEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlQueryCounterEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   185,
		New:  func() atom.Atom { return &GlQueryCounterEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   186,
		New:  func() atom.Atom { return &GlGetQueryivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   187,
		New:  func() atom.Atom { return &GlGetQueryObjectivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectuivEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   188,
		New:  func() atom.Atom { return &GlGetQueryObjectuivEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjecti64vEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   189,
		New:  func() atom.Atom { return &GlGetQueryObjecti64vEXT{} },
	})
	atom.Register(atom.TypeInfo{
		Name: "GlGetQueryObjectui64vEXT",
		Docs: "[http://www.khronos.org/registry/gles/extensions/EXT/EXT_disjoint_timer_query.txt]",
		ID:   190,
		New:  func() atom.Atom { return &GlGetQueryObjectui64vEXT{} },
	})
}
