# gles
--
    import "android.googlesource.com/platform/tools/gpu/gfxapi/gles"

Package gles implementes the API interface for the OpenGL ES graphics library.

## Usage

```go
const (
	DrawMode_GL_LINE_LOOP      = DrawMode(2)
	DrawMode_GL_LINE_STRIP     = DrawMode(3)
	DrawMode_GL_LINES          = DrawMode(1)
	DrawMode_GL_POINTS         = DrawMode(0)
	DrawMode_GL_TRIANGLE_FAN   = DrawMode(6)
	DrawMode_GL_TRIANGLE_STRIP = DrawMode(5)
	DrawMode_GL_TRIANGLES      = DrawMode(4)
)
```

```go
const (
	IndicesType_GL_UNSIGNED_BYTE  = IndicesType(5121)
	IndicesType_GL_UNSIGNED_SHORT = IndicesType(5123)
	IndicesType_GL_UNSIGNED_INT   = IndicesType(5125)
)
```

```go
const (
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X = CubeMapImageTarget(34070)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = CubeMapImageTarget(34072)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = CubeMapImageTarget(34074)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X = CubeMapImageTarget(34069)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y = CubeMapImageTarget(34071)
	CubeMapImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z = CubeMapImageTarget(34073)
)
```

```go
const (
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_X = TextureImageTarget(34070)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = TextureImageTarget(34072)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = TextureImageTarget(34074)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_X = TextureImageTarget(34069)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Y = TextureImageTarget(34071)
	TextureImageTarget_GL_TEXTURE_CUBE_MAP_POSITIVE_Z = TextureImageTarget(34073)
)
```
CubeMapImageTarget

```go
const (
	BaseTexelFormat_GL_ALPHA = BaseTexelFormat(6406)
	BaseTexelFormat_GL_RGB   = BaseTexelFormat(6407)
	BaseTexelFormat_GL_RGBA  = BaseTexelFormat(6408)
)
```

```go
const (
	TexelFormat_GLES_1_1_GL_LUMINANCE       = TexelFormat_GLES_1_1(6409)
	TexelFormat_GLES_1_1_GL_LUMINANCE_ALPHA = TexelFormat_GLES_1_1(6410)
)
```

```go
const (
	TexelFormat_GLES_1_1_GL_ALPHA = TexelFormat_GLES_1_1(6406)
	TexelFormat_GLES_1_1_GL_RGB   = TexelFormat_GLES_1_1(6407)
	TexelFormat_GLES_1_1_GL_RGBA  = TexelFormat_GLES_1_1(6408)
)
```
BaseTexelFormat

```go
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
```

```go
const (
	TexelFormat_GL_LUMINANCE       = TexelFormat(6409)
	TexelFormat_GL_LUMINANCE_ALPHA = TexelFormat(6410)
)
```
TexelFormat_GLES_1_1

```go
const (
	TexelFormat_GL_ALPHA = TexelFormat(6406)
	TexelFormat_GL_RGB   = TexelFormat(6407)
	TexelFormat_GL_RGBA  = TexelFormat(6408)
)
```
BaseTexelFormat

```go
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
```
TexelFormat_GLES_3_0

```go
const (
	RenderbufferFormat_GL_RGBA4             = RenderbufferFormat(32854)
	RenderbufferFormat_GL_RGB5_A1           = RenderbufferFormat(32855)
	RenderbufferFormat_GL_RGB565            = RenderbufferFormat(36194)
	RenderbufferFormat_GL_RGBA8             = RenderbufferFormat(32856)
	RenderbufferFormat_GL_DEPTH_COMPONENT16 = RenderbufferFormat(33189)
	RenderbufferFormat_GL_STENCIL_INDEX8    = RenderbufferFormat(36168)
)
```

```go
const (
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGB_AMD                     = CompressedTexelFormat_AMD_compressed_ATC_texture(35986)
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD     = CompressedTexelFormat_AMD_compressed_ATC_texture(35987)
	CompressedTexelFormat_AMD_compressed_ATC_texture_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = CompressedTexelFormat_AMD_compressed_ATC_texture(34798)
)
```

```go
const (
	CompressedTexelFormat_GL_ATC_RGB_AMD                     = CompressedTexelFormat(35986)
	CompressedTexelFormat_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD     = CompressedTexelFormat(35987)
	CompressedTexelFormat_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = CompressedTexelFormat(34798)
)
```
CompressedTexelFormat_AMD_compressed_ATC_texture

```go
const (
	ImageTexelFormat_GL_LUMINANCE       = ImageTexelFormat(6409)
	ImageTexelFormat_GL_LUMINANCE_ALPHA = ImageTexelFormat(6410)
)
```
TexelFormat_GLES_1_1

```go
const (
	ImageTexelFormat_GL_ALPHA = ImageTexelFormat(6406)
	ImageTexelFormat_GL_RGB   = ImageTexelFormat(6407)
	ImageTexelFormat_GL_RGBA  = ImageTexelFormat(6408)
)
```
BaseTexelFormat

```go
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
```
TexelFormat_GLES_3_0

```go
const (
	ImageTexelFormat_GL_ATC_RGB_AMD                     = ImageTexelFormat(35986)
	ImageTexelFormat_GL_ATC_RGBA_EXPLICIT_ALPHA_AMD     = ImageTexelFormat(35987)
	ImageTexelFormat_GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = ImageTexelFormat(34798)
)
```
CompressedTexelFormat_AMD_compressed_ATC_texture

```go
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
```

```go
const (
	FramebufferAttachment_GL_COLOR_ATTACHMENT0  = FramebufferAttachment(36064)
	FramebufferAttachment_GL_DEPTH_ATTACHMENT   = FramebufferAttachment(36096)
	FramebufferAttachment_GL_STENCIL_ATTACHMENT = FramebufferAttachment(36128)
)
```

```go
const (
	FramebufferAttachmentType_GL_NONE         = FramebufferAttachmentType(0)
	FramebufferAttachmentType_GL_RENDERBUFFER = FramebufferAttachmentType(36161)
	FramebufferAttachmentType_GL_TEXTURE      = FramebufferAttachmentType(5890)
)
```

```go
const (
	FramebufferTarget_GLES_3_1_GL_READ_FRAMEBUFFER = FramebufferTarget_GLES_3_1(36008)
	FramebufferTarget_GLES_3_1_GL_DRAW_FRAMEBUFFER = FramebufferTarget_GLES_3_1(36009)
)
```

```go
const (
	FramebufferTarget_GL_READ_FRAMEBUFFER = FramebufferTarget(36008)
	FramebufferTarget_GL_DRAW_FRAMEBUFFER = FramebufferTarget(36009)
)
```
FramebufferTarget_GLES_3_1

```go
const (
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE           = FramebufferAttachmentParameter(36048)
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME           = FramebufferAttachmentParameter(36049)
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL         = FramebufferAttachmentParameter(36050)
	FramebufferAttachmentParameter_GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = FramebufferAttachmentParameter(36051)
)
```

```go
const (
	FramebufferStatus_GL_FRAMEBUFFER_COMPLETE                      = FramebufferStatus(36053)
	FramebufferStatus_GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT         = FramebufferStatus(36054)
	FramebufferStatus_GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = FramebufferStatus(36055)
	FramebufferStatus_GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS         = FramebufferStatus(36057)
	FramebufferStatus_GL_FRAMEBUFFER_UNSUPPORTED                   = FramebufferStatus(36061)
)
```

```go
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
```

```go
const (
	BufferParameter_GL_BUFFER_SIZE  = BufferParameter(34660)
	BufferParameter_GL_BUFFER_USAGE = BufferParameter(34661)
)
```

```go
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
```

```go
const (
	BufferUsage_GL_DYNAMIC_DRAW = BufferUsage(35048)
	BufferUsage_GL_STATIC_DRAW  = BufferUsage(35044)
	BufferUsage_GL_STREAM_DRAW  = BufferUsage(35040)
)
```

```go
const (
	ShaderType_GL_VERTEX_SHADER   = ShaderType(35633)
	ShaderType_GL_FRAGMENT_SHADER = ShaderType(35632)
)
```

```go
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
```

```go
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
```
StateVariable_GLES_2_0

```go
const (
	FaceMode_GL_FRONT          = FaceMode(1028)
	FaceMode_GL_BACK           = FaceMode(1029)
	FaceMode_GL_FRONT_AND_BACK = FaceMode(1032)
)
```

```go
const (
	ArrayType_GLES_1_1_GL_VERTEX_ARRAY        = ArrayType_GLES_1_1(32884)
	ArrayType_GLES_1_1_GL_NORMAL_ARRAY        = ArrayType_GLES_1_1(32885)
	ArrayType_GLES_1_1_GL_COLOR_ARRAY         = ArrayType_GLES_1_1(32886)
	ArrayType_GLES_1_1_GL_TEXTURE_COORD_ARRAY = ArrayType_GLES_1_1(32888)
)
```

```go
const (
	ArrayType_GL_VERTEX_ARRAY        = ArrayType(32884)
	ArrayType_GL_NORMAL_ARRAY        = ArrayType(32885)
	ArrayType_GL_COLOR_ARRAY         = ArrayType(32886)
	ArrayType_GL_TEXTURE_COORD_ARRAY = ArrayType(32888)
)
```
ArrayType_GLES_1_1

```go
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
```

```go
const (
	Capability_GL_VERTEX_ARRAY        = Capability(32884)
	Capability_GL_NORMAL_ARRAY        = Capability(32885)
	Capability_GL_COLOR_ARRAY         = Capability(32886)
	Capability_GL_TEXTURE_COORD_ARRAY = Capability(32888)
)
```
ArrayType_GLES_1_1

```go
const (
	StringConstant_GL_EXTENSIONS = StringConstant(7939)
	StringConstant_GL_RENDERER   = StringConstant(7937)
	StringConstant_GL_VENDOR     = StringConstant(7936)
	StringConstant_GL_VERSION    = StringConstant(7938)
)
```

```go
const (
	VertexAttribType_GL_BYTE           = VertexAttribType(5120)
	VertexAttribType_GL_FIXED          = VertexAttribType(5132)
	VertexAttribType_GL_FLOAT          = VertexAttribType(5126)
	VertexAttribType_GL_SHORT          = VertexAttribType(5122)
	VertexAttribType_GL_UNSIGNED_BYTE  = VertexAttribType(5121)
	VertexAttribType_GL_UNSIGNED_SHORT = VertexAttribType(5123)
)
```

```go
const (
	ShaderAttribType_GL_FLOAT      = ShaderAttribType(5126)
	ShaderAttribType_GL_FLOAT_VEC2 = ShaderAttribType(35664)
	ShaderAttribType_GL_FLOAT_VEC3 = ShaderAttribType(35665)
	ShaderAttribType_GL_FLOAT_VEC4 = ShaderAttribType(35666)
	ShaderAttribType_GL_FLOAT_MAT2 = ShaderAttribType(35674)
	ShaderAttribType_GL_FLOAT_MAT3 = ShaderAttribType(35675)
	ShaderAttribType_GL_FLOAT_MAT4 = ShaderAttribType(35676)
)
```

```go
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
```

```go
const (
	Error_GL_NO_ERROR                      = Error(0)
	Error_GL_INVALID_ENUM                  = Error(1280)
	Error_GL_INVALID_VALUE                 = Error(1281)
	Error_GL_INVALID_OPERATION             = Error(1282)
	Error_GL_INVALID_FRAMEBUFFER_OPERATION = Error(1286)
	Error_GL_OUT_OF_MEMORY                 = Error(1285)
)
```

```go
const (
	HintMode_GL_DONT_CARE = HintMode(4352)
	HintMode_GL_FASTEST   = HintMode(4353)
	HintMode_GL_NICEST    = HintMode(4354)
)
```

```go
const (
	DiscardFramebufferAttachment_GL_COLOR_EXT   = DiscardFramebufferAttachment(6144)
	DiscardFramebufferAttachment_GL_DEPTH_EXT   = DiscardFramebufferAttachment(6145)
	DiscardFramebufferAttachment_GL_STENCIL_EXT = DiscardFramebufferAttachment(6146)
)
```

```go
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
```

```go
const (
	ShaderParameter_GL_SHADER_TYPE          = ShaderParameter(35663)
	ShaderParameter_GL_DELETE_STATUS        = ShaderParameter(35712)
	ShaderParameter_GL_COMPILE_STATUS       = ShaderParameter(35713)
	ShaderParameter_GL_INFO_LOG_LENGTH      = ShaderParameter(35716)
	ShaderParameter_GL_SHADER_SOURCE_LENGTH = ShaderParameter(35720)
)
```

```go
const (
	PixelStoreParameter_GL_PACK_ALIGNMENT   = PixelStoreParameter(3333)
	PixelStoreParameter_GL_UNPACK_ALIGNMENT = PixelStoreParameter(3317)
)
```

```go
const (
	TextureParameter_FilterMode_GL_TEXTURE_MIN_FILTER = TextureParameter_FilterMode(10241)
	TextureParameter_FilterMode_GL_TEXTURE_MAG_FILTER = TextureParameter_FilterMode(10240)
)
```

```go
const (
	TextureParameter_WrapMode_GL_TEXTURE_WRAP_S = TextureParameter_WrapMode(10242)
	TextureParameter_WrapMode_GL_TEXTURE_WRAP_T = TextureParameter_WrapMode(10243)
)
```

```go
const (
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_R = TextureParameter_SwizzleMode(36418)
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_G = TextureParameter_SwizzleMode(36419)
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_B = TextureParameter_SwizzleMode(36420)
	TextureParameter_SwizzleMode_GL_TEXTURE_SWIZZLE_A = TextureParameter_SwizzleMode(36421)
)
```

```go
const (
	TextureParameter_GL_TEXTURE_MIN_FILTER = TextureParameter(10241)
	TextureParameter_GL_TEXTURE_MAG_FILTER = TextureParameter(10240)
)
```
TextureParameter_FilterMode

```go
const (
	TextureParameter_GL_TEXTURE_WRAP_S = TextureParameter(10242)
	TextureParameter_GL_TEXTURE_WRAP_T = TextureParameter(10243)
)
```
TextureParameter_WrapMode

```go
const (
	TextureParameter_GL_TEXTURE_SWIZZLE_R = TextureParameter(36418)
	TextureParameter_GL_TEXTURE_SWIZZLE_G = TextureParameter(36419)
	TextureParameter_GL_TEXTURE_SWIZZLE_B = TextureParameter(36420)
	TextureParameter_GL_TEXTURE_SWIZZLE_A = TextureParameter(36421)
)
```
TextureParameter_SwizzleMode

```go
const (
	TextureFilterMode_GL_NEAREST                = TextureFilterMode(9728)
	TextureFilterMode_GL_LINEAR                 = TextureFilterMode(9729)
	TextureFilterMode_GL_NEAREST_MIPMAP_NEAREST = TextureFilterMode(9984)
	TextureFilterMode_GL_LINEAR_MIPMAP_NEAREST  = TextureFilterMode(9985)
	TextureFilterMode_GL_NEAREST_MIPMAP_LINEAR  = TextureFilterMode(9986)
	TextureFilterMode_GL_LINEAR_MIPMAP_LINEAR   = TextureFilterMode(9987)
)
```

```go
const (
	TextureWrapMode_GL_CLAMP_TO_EDGE   = TextureWrapMode(33071)
	TextureWrapMode_GL_MIRRORED_REPEAT = TextureWrapMode(33648)
	TextureWrapMode_GL_REPEAT          = TextureWrapMode(10497)
)
```

```go
const (
	TexelComponent_GL_RED   = TexelComponent(6403)
	TexelComponent_GL_GREEN = TexelComponent(6404)
	TexelComponent_GL_BLUE  = TexelComponent(6405)
	TexelComponent_GL_ALPHA = TexelComponent(6406)
)
```

```go
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
```

```go
const (
	PrecisionType_GL_LOW_FLOAT    = PrecisionType(36336)
	PrecisionType_GL_MEDIUM_FLOAT = PrecisionType(36337)
	PrecisionType_GL_HIGH_FLOAT   = PrecisionType(36338)
	PrecisionType_GL_LOW_INT      = PrecisionType(36339)
	PrecisionType_GL_MEDIUM_INT   = PrecisionType(36340)
	PrecisionType_GL_HIGH_INT     = PrecisionType(36341)
)
```

```go
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
```

```go
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
```

```go
const (
	FaceOrientation_GL_CW  = FaceOrientation(2304)
	FaceOrientation_GL_CCW = FaceOrientation(2305)
)
```

```go
const (
	BlendEquation_GL_FUNC_ADD              = BlendEquation(32774)
	BlendEquation_GL_FUNC_SUBTRACT         = BlendEquation(32778)
	BlendEquation_GL_FUNC_REVERSE_SUBTRACT = BlendEquation(32779)
)
```

```go
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
```

```go
const (
	ResetStatus_GL_NO_ERROR                   = ResetStatus(0)
	ResetStatus_GL_GUILTY_CONTEXT_RESET_EXT   = ResetStatus(33363)
	ResetStatus_GL_INNOCENT_CONTEXT_RESET_EXT = ResetStatus(33364)
	ResetStatus_GL_UNKNOWN_CONTEXT_RESET_EXT  = ResetStatus(33365)
)
```

```go
const (
	TextureKind_UNDEFINED = TextureKind(0)
	TextureKind_TEXTURE2D = TextureKind(1)
	TextureKind_CUBEMAP   = TextureKind(2)
)
```

```go
const (
	QueryObjectParameter_GLES_3_GL_QUERY_RESULT           = QueryObjectParameter_GLES_3(34918)
	QueryObjectParameter_GLES_3_GL_QUERY_RESULT_AVAILABLE = QueryObjectParameter_GLES_3(34919)
)
```

```go
const (
	QueryObjectParameter_GL_QUERY_RESULT           = QueryObjectParameter(34918)
	QueryObjectParameter_GL_QUERY_RESULT_AVAILABLE = QueryObjectParameter(34919)
)
```
QueryObjectParameter_GLES_3

```go
const (
	QueryTarget_GLES_3_GL_ANY_SAMPLES_PASSED                    = QueryTarget_GLES_3(35887)
	QueryTarget_GLES_3_GL_ANY_SAMPLES_PASSED_CONSERVATIVE       = QueryTarget_GLES_3(36202)
	QueryTarget_GLES_3_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = QueryTarget_GLES_3(35976)
)
```

```go
const (
	QueryTarget_EXT_disjoint_timer_query_GL_TIME_ELAPSED_EXT = QueryTarget_EXT_disjoint_timer_query(35007)
	QueryTarget_EXT_disjoint_timer_query_GL_TIMESTAMP_EXT    = QueryTarget_EXT_disjoint_timer_query(36392)
)
```

```go
const (
	QueryTarget_GL_ANY_SAMPLES_PASSED                    = QueryTarget(35887)
	QueryTarget_GL_ANY_SAMPLES_PASSED_CONSERVATIVE       = QueryTarget(36202)
	QueryTarget_GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = QueryTarget(35976)
)
```
QueryTarget_GLES_3

```go
const (
	QueryTarget_GL_TIME_ELAPSED_EXT = QueryTarget(35007)
	QueryTarget_GL_TIMESTAMP_EXT    = QueryTarget(36392)
)
```
QueryTarget_EXT_disjoint_timer_query

```go
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
```

```go
const (
	ClearMask_GL_COLOR_BUFFER_BIT   = ClearMask(16384)
	ClearMask_GL_DEPTH_BUFFER_BIT   = ClearMask(256)
	ClearMask_GL_STENCIL_BUFFER_BIT = ClearMask(1024)
)
```

```go
const (
	MapBufferRangeAccess_GL_MAP_READ_BIT              = MapBufferRangeAccess(1)
	MapBufferRangeAccess_GL_MAP_WRITE_BIT             = MapBufferRangeAccess(2)
	MapBufferRangeAccess_GL_MAP_INVALIDATE_RANGE_BIT  = MapBufferRangeAccess(4)
	MapBufferRangeAccess_GL_MAP_INVALIDATE_BUFFER_BIT = MapBufferRangeAccess(8)
	MapBufferRangeAccess_GL_MAP_FLUSH_EXPLICIT_BIT    = MapBufferRangeAccess(16)
	MapBufferRangeAccess_GL_MAP_UNSYNCHRONIZED_BIT    = MapBufferRangeAccess(32)
)
```

```go
const (
	ArrayType_GL_POINT_SIZE_ARRAY_OES = ArrayType(35740)
)
```
ArrayType_OES_point_size_array

```go
const (
	ArrayType_OES_point_size_array_GL_POINT_SIZE_ARRAY_OES = ArrayType_OES_point_size_array(35740)
)
```

```go
const (
	Capability_GL_POINT_SIZE_ARRAY_OES = Capability(35740)
)
```
ArrayType_OES_point_size_array

```go
const (
	CompressedTexelFormat_GL_ETC1_RGB8_OES = CompressedTexelFormat(36196)
)
```
CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture

```go
const (
	CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture_GL_ETC1_RGB8_OES = CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture(36196)
)
```

```go
const (
	FramebufferTarget_GLES_2_0_GL_FRAMEBUFFER = FramebufferTarget_GLES_2_0(36160)
)
```

```go
const (
	FramebufferTarget_GL_FRAMEBUFFER = FramebufferTarget(36160)
)
```
FramebufferTarget_GLES_2_0

```go
const (
	HintTarget_GL_GENERATE_MIPMAP_HINT = HintTarget(33170)
)
```

```go
const (
	ImageTargetRenderbufferStorage_GL_RENDERBUFFER_OES = ImageTargetRenderbufferStorage(36161)
)
```

```go
const (
	ImageTargetTexture_GL_TEXTURE_2D = ImageTargetTexture(3553)
)
```
ImageTargetTexture_OES_EGL_image

```go
const (
	ImageTargetTexture_GL_TEXTURE_EXTERNAL_OES = ImageTargetTexture(36197)
)
```
ImageTargetTexture_OES_EGL_image_external

```go
const (
	ImageTargetTexture_OES_EGL_image_GL_TEXTURE_2D = ImageTargetTexture_OES_EGL_image(3553)
)
```

```go
const (
	ImageTargetTexture_OES_EGL_image_external_GL_TEXTURE_EXTERNAL_OES = ImageTargetTexture_OES_EGL_image_external(36197)
)
```

```go
const (
	ImageTexelFormat_GL_ETC1_RGB8_OES = ImageTexelFormat(36196)
)
```
CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture

```go
const (
	QueryParameter_EXT_disjoint_timer_query_GL_QUERY_COUNTER_BITS_EXT = QueryParameter_EXT_disjoint_timer_query(34916)
)
```

```go
const (
	QueryParameter_GLES_3_GL_CURRENT_QUERY = QueryParameter_GLES_3(34917)
)
```

```go
const (
	QueryParameter_GL_CURRENT_QUERY = QueryParameter(34917)
)
```
QueryParameter_GLES_3

```go
const (
	QueryParameter_GL_QUERY_COUNTER_BITS_EXT = QueryParameter(34916)
)
```
QueryParameter_EXT_disjoint_timer_query

```go
const (
	RenderbufferTarget_GL_RENDERBUFFER = RenderbufferTarget(36161)
)
```

```go
const (
	StateVariable_EXT_disjoint_timer_query_GL_GPU_DISJOINT_EXT = StateVariable_EXT_disjoint_timer_query(36795)
)
```

```go
const (
	StateVariable_EXT_texture_filter_anisotropic_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = StateVariable_EXT_texture_filter_anisotropic(34047)
)
```

```go
const (
	StateVariable_GLES_3_1_GL_READ_FRAMEBUFFER_BINDING = StateVariable_GLES_3_1(36010)
)
```

```go
const (
	StateVariable_GL_GPU_DISJOINT_EXT = StateVariable(36795)
)
```
StateVariable_EXT_disjoint_timer_query

```go
const (
	StateVariable_GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = StateVariable(34047)
)
```
StateVariable_EXT_texture_filter_anisotropic

```go
const (
	StateVariable_GL_READ_FRAMEBUFFER_BINDING = StateVariable(36010)
)
```
StateVariable_GLES_3_1

```go
const (
	Texture2DImageTarget_GL_TEXTURE_2D = Texture2DImageTarget(3553)
)
```

```go
const (
	TextureImageTarget_GL_TEXTURE_2D = TextureImageTarget(3553)
)
```
Texture2DImageTarget

```go
const (
	TextureParameter_EXT_texture_filter_anisotropic_GL_TEXTURE_MAX_ANISOTROPY_EXT = TextureParameter_EXT_texture_filter_anisotropic(34046)
)
```

```go
const (
	TextureParameter_GL_TEXTURE_MAX_ANISOTROPY_EXT = TextureParameter(34046)
)
```
TextureParameter_EXT_texture_filter_anisotropic

```go
const (
	TextureTarget_GLES_1_1_GL_TEXTURE_2D = TextureTarget_GLES_1_1(3553)
)
```

```go
const (
	TextureTarget_GLES_2_0_GL_TEXTURE_CUBE_MAP = TextureTarget_GLES_2_0(34067)
)
```

```go
const (
	TextureTarget_GL_TEXTURE_2D = TextureTarget(3553)
)
```
TextureTarget_GLES_1_1

```go
const (
	TextureTarget_GL_TEXTURE_CUBE_MAP = TextureTarget(34067)
)
```
TextureTarget_GLES_2_0

```go
const (
	TextureTarget_GL_TEXTURE_EXTERNAL_OES = TextureTarget(36197)
)
```
TextureTarget_OES_EGL_image_external

```go
const (
	TextureTarget_OES_EGL_image_external_GL_TEXTURE_EXTERNAL_OES = TextureTarget_OES_EGL_image_external(36197)
)
```

```go
const (
	Type_ARB_half_float_vertex_GL_ARB_half_float_vertex = Type_ARB_half_float_vertex(5131)
)
```

```go
const (
	Type_OES_vertex_half_float_GL_HALF_FLOAT_OES = Type_OES_vertex_half_float(36193)
)
```

```go
const (
	VertexAttribType_GL_ARB_half_float_vertex = VertexAttribType(5131)
)
```
Type_ARB_half_float_vertex

```go
const (
	VertexAttribType_GL_HALF_FLOAT_OES = VertexAttribType(36193)
)
```
Type_OES_vertex_half_float

#### func  API

```go
func API() gfxapi.API
```

#### type ArrayType

```go
type ArrayType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ArrayType
//////////////////////////////////////////////////////////////////////////////

#### func (ArrayType) String

```go
func (v ArrayType) String() string
```

#### type ArrayType_GLES_1_1

```go
type ArrayType_GLES_1_1 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ArrayType_GLES_1_1
//////////////////////////////////////////////////////////////////////////////

#### func (ArrayType_GLES_1_1) String

```go
func (v ArrayType_GLES_1_1) String() string
```

#### type ArrayType_OES_point_size_array

```go
type ArrayType_OES_point_size_array uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ArrayType_OES_point_size_array
//////////////////////////////////////////////////////////////////////////////

#### func (ArrayType_OES_point_size_array) String

```go
func (v ArrayType_OES_point_size_array) String() string
```

#### type AttributeLocation

```go
type AttributeLocation uint32
```


#### func (*AttributeLocation) Equal

```go
func (c *AttributeLocation) Equal(rhs AttributeLocation) bool
```

#### func (*AttributeLocation) Less

```go
func (c *AttributeLocation) Less(rhs AttributeLocation) bool
```

#### type AttributeLocation_CharBufferMap

```go
type AttributeLocation_CharBufferMap map[string]AttributeLocation
```


#### func (AttributeLocation_CharBufferMap) Contains

```go
func (m AttributeLocation_CharBufferMap) Contains(key string) bool
```

#### func (AttributeLocation_CharBufferMap) Delete

```go
func (m AttributeLocation_CharBufferMap) Delete(key string)
```

#### func (AttributeLocation_CharBufferMap) Get

```go
func (m AttributeLocation_CharBufferMap) Get(key string) AttributeLocation
```

#### func (AttributeLocation_CharBufferMap) Range

```go
func (m AttributeLocation_CharBufferMap) Range() []AttributeLocation
```

#### type BOOL

```go
type BOOL int64
```


#### func (*BOOL) Equal

```go
func (c *BOOL) Equal(rhs BOOL) bool
```

#### func (*BOOL) Less

```go
func (c *BOOL) Less(rhs BOOL) bool
```

#### type BackbufferInfo

```go
type BackbufferInfo struct {
	binary.Generate
	Width                int32
	Height               int32
	ColorFmt             RenderbufferFormat
	DepthFmt             RenderbufferFormat
	StencilFmt           RenderbufferFormat
	ResetViewportScissor bool
}
```

//////////////////////////////////////////////////////////////////////////////
BackbufferInfo
//////////////////////////////////////////////////////////////////////////////

#### func  NewBackbufferInfo

```go
func NewBackbufferInfo(
	pWidth int32,
	pHeight int32,
	pColorFmt RenderbufferFormat,
	pDepthFmt RenderbufferFormat,
	pStencilFmt RenderbufferFormat,
	pResetViewportScissor bool,
) *BackbufferInfo
```

#### func (*BackbufferInfo) API

```go
func (c *BackbufferInfo) API() gfxapi.API
```

#### func (*BackbufferInfo) Class

```go
func (*BackbufferInfo) Class() binary.Class
```

#### func (*BackbufferInfo) Flags

```go
func (c *BackbufferInfo) Flags() atom.Flags
```

#### func (*BackbufferInfo) Mutate

```go
func (ϟa *BackbufferInfo) Mutate(ϟs *gfxapi.State) error
```

#### func (*BackbufferInfo) Replay

```go
func (ϟa *BackbufferInfo) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*BackbufferInfo) String

```go
func (c *BackbufferInfo) String() string
```

#### func (*BackbufferInfo) TypeID

```go
func (c *BackbufferInfo) TypeID() atom.TypeID
```

#### type BaseTexelFormat

```go
type BaseTexelFormat uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BaseTexelFormat
//////////////////////////////////////////////////////////////////////////////

#### func (BaseTexelFormat) String

```go
func (v BaseTexelFormat) String() string
```

#### type BlendEquation

```go
type BlendEquation uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BlendEquation
//////////////////////////////////////////////////////////////////////////////

#### func (BlendEquation) String

```go
func (v BlendEquation) String() string
```

#### type BlendFactor

```go
type BlendFactor uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BlendFactor
//////////////////////////////////////////////////////////////////////////////

#### func (BlendFactor) String

```go
func (v BlendFactor) String() string
```

#### type BlendState

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class BlendState
//////////////////////////////////////////////////////////////////////////////

#### func (*BlendState) Class

```go
func (*BlendState) Class() binary.Class
```

#### func (*BlendState) GetCreatedAt

```go
func (c *BlendState) GetCreatedAt() atom.ID
```

#### func (*BlendState) Init

```go
func (c *BlendState) Init()
```

#### type BoolArray

```go
type BoolArray []bool
```


#### func (BoolArray) Len

```go
func (s BoolArray) Len() int
```

#### func (BoolArray) Range

```go
func (s BoolArray) Range() []bool
```

#### type Bool_CapabilityMap

```go
type Bool_CapabilityMap map[Capability]bool
```


#### func (Bool_CapabilityMap) Contains

```go
func (m Bool_CapabilityMap) Contains(key Capability) bool
```

#### func (Bool_CapabilityMap) Delete

```go
func (m Bool_CapabilityMap) Delete(key Capability)
```

#### func (Bool_CapabilityMap) Get

```go
func (m Bool_CapabilityMap) Get(key Capability) bool
```

#### func (Bool_CapabilityMap) Range

```go
func (m Bool_CapabilityMap) Range() []bool
```

#### type Buffer

```go
type Buffer struct {
	binary.Generate
	CreatedAt atom.ID
	Data      memory.Memory
	Size      int32
	Usage     BufferUsage
}
```

//////////////////////////////////////////////////////////////////////////////
class Buffer
//////////////////////////////////////////////////////////////////////////////

#### func (*Buffer) Class

```go
func (*Buffer) Class() binary.Class
```

#### func (*Buffer) GetCreatedAt

```go
func (c *Buffer) GetCreatedAt() atom.ID
```

#### func (*Buffer) Init

```go
func (c *Buffer) Init()
```

#### type BufferArray

```go
type BufferArray []Buffer
```


#### func (BufferArray) Len

```go
func (s BufferArray) Len() int
```

#### func (BufferArray) Range

```go
func (s BufferArray) Range() []Buffer
```

#### type BufferDataPointer

```go
type BufferDataPointer memory.Pointer
```


#### func (*BufferDataPointer) Equal

```go
func (c *BufferDataPointer) Equal(rhs BufferDataPointer) bool
```

#### func (*BufferDataPointer) Less

```go
func (c *BufferDataPointer) Less(rhs BufferDataPointer) bool
```

#### type BufferId

```go
type BufferId uint32
```


#### func (*BufferId) Equal

```go
func (c *BufferId) Equal(rhs BufferId) bool
```

#### func (*BufferId) Less

```go
func (c *BufferId) Less(rhs BufferId) bool
```

#### type BufferIdArray

```go
type BufferIdArray []BufferId
```


#### func (BufferIdArray) Len

```go
func (s BufferIdArray) Len() int
```

#### func (BufferIdArray) Range

```go
func (s BufferIdArray) Range() []BufferId
```

#### type BufferId_BufferTargetMap

```go
type BufferId_BufferTargetMap map[BufferTarget]BufferId
```


#### func (BufferId_BufferTargetMap) Contains

```go
func (m BufferId_BufferTargetMap) Contains(key BufferTarget) bool
```

#### func (BufferId_BufferTargetMap) Delete

```go
func (m BufferId_BufferTargetMap) Delete(key BufferTarget)
```

#### func (BufferId_BufferTargetMap) Get

```go
func (m BufferId_BufferTargetMap) Get(key BufferTarget) BufferId
```

#### func (BufferId_BufferTargetMap) Range

```go
func (m BufferId_BufferTargetMap) Range() []BufferId
```

#### type BufferParameter

```go
type BufferParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BufferParameter
//////////////////////////////////////////////////////////////////////////////

#### func (BufferParameter) String

```go
func (v BufferParameter) String() string
```

#### type BufferPtr_BufferIdMap

```go
type BufferPtr_BufferIdMap map[BufferId]*Buffer
```


#### func (BufferPtr_BufferIdMap) Contains

```go
func (m BufferPtr_BufferIdMap) Contains(key BufferId) bool
```

#### func (BufferPtr_BufferIdMap) Delete

```go
func (m BufferPtr_BufferIdMap) Delete(key BufferId)
```

#### func (BufferPtr_BufferIdMap) Get

```go
func (m BufferPtr_BufferIdMap) Get(key BufferId) *Buffer
```

#### func (BufferPtr_BufferIdMap) Range

```go
func (m BufferPtr_BufferIdMap) Range() []*Buffer
```

#### type BufferTarget

```go
type BufferTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BufferTarget
//////////////////////////////////////////////////////////////////////////////

#### func (BufferTarget) String

```go
func (v BufferTarget) String() string
```

#### type BufferUsage

```go
type BufferUsage uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BufferUsage
//////////////////////////////////////////////////////////////////////////////

#### func (BufferUsage) String

```go
func (v BufferUsage) String() string
```

#### type CGLContextObj

```go
type CGLContextObj memory.Pointer
```


#### func (*CGLContextObj) Equal

```go
func (c *CGLContextObj) Equal(rhs CGLContextObj) bool
```

#### func (*CGLContextObj) Less

```go
func (c *CGLContextObj) Less(rhs CGLContextObj) bool
```

#### type CGLCreateContext

```go
type CGLCreateContext struct {
	binary.Generate
	Pix    CGLPixelFormatObj
	Share  CGLContextObj
	Ctx    CGLContextObj
	Result CGLError
}
```

//////////////////////////////////////////////////////////////////////////////
CGLCreateContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewCGLCreateContext

```go
func NewCGLCreateContext(
	pPix CGLPixelFormatObj,
	pShare CGLContextObj,
	pCtx CGLContextObj,
	pResult CGLError,
) *CGLCreateContext
```

#### func (*CGLCreateContext) API

```go
func (c *CGLCreateContext) API() gfxapi.API
```

#### func (*CGLCreateContext) Class

```go
func (*CGLCreateContext) Class() binary.Class
```

#### func (*CGLCreateContext) Flags

```go
func (c *CGLCreateContext) Flags() atom.Flags
```

#### func (*CGLCreateContext) Mutate

```go
func (ϟa *CGLCreateContext) Mutate(ϟs *gfxapi.State) error
```

#### func (*CGLCreateContext) Replay

```go
func (ω *CGLCreateContext) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*CGLCreateContext) String

```go
func (c *CGLCreateContext) String() string
```

#### func (*CGLCreateContext) TypeID

```go
func (c *CGLCreateContext) TypeID() atom.TypeID
```

#### type CGLCreateContext_Postback

```go
type CGLCreateContext_Postback struct {
	Ctx    []byte
	Result CGLError
}
```


#### func (*CGLCreateContext_Postback) Decode

```go
func (o *CGLCreateContext_Postback) Decode(ctx_cnt uint64, d binary.Decoder) error
```

#### type CGLError

```go
type CGLError int64
```


#### func (*CGLError) Equal

```go
func (c *CGLError) Equal(rhs CGLError) bool
```

#### func (*CGLError) Less

```go
func (c *CGLError) Less(rhs CGLError) bool
```

#### type CGLPixelFormatObj

```go
type CGLPixelFormatObj memory.Pointer
```


#### func (*CGLPixelFormatObj) Equal

```go
func (c *CGLPixelFormatObj) Equal(rhs CGLPixelFormatObj) bool
```

#### func (*CGLPixelFormatObj) Less

```go
func (c *CGLPixelFormatObj) Less(rhs CGLPixelFormatObj) bool
```

#### type CGLSetCurrentContext

```go
type CGLSetCurrentContext struct {
	binary.Generate
	Ctx    CGLContextObj
	Result CGLError
}
```

//////////////////////////////////////////////////////////////////////////////
CGLSetCurrentContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewCGLSetCurrentContext

```go
func NewCGLSetCurrentContext(
	pCtx CGLContextObj,
	pResult CGLError,
) *CGLSetCurrentContext
```

#### func (*CGLSetCurrentContext) API

```go
func (c *CGLSetCurrentContext) API() gfxapi.API
```

#### func (*CGLSetCurrentContext) Class

```go
func (*CGLSetCurrentContext) Class() binary.Class
```

#### func (*CGLSetCurrentContext) Flags

```go
func (c *CGLSetCurrentContext) Flags() atom.Flags
```

#### func (*CGLSetCurrentContext) Mutate

```go
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State) error
```

#### func (*CGLSetCurrentContext) Replay

```go
func (ω *CGLSetCurrentContext) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*CGLSetCurrentContext) String

```go
func (c *CGLSetCurrentContext) String() string
```

#### func (*CGLSetCurrentContext) TypeID

```go
func (c *CGLSetCurrentContext) TypeID() atom.TypeID
```

#### type CGLSetCurrentContext_Postback

```go
type CGLSetCurrentContext_Postback struct {
	Result CGLError
}
```


#### func (*CGLSetCurrentContext_Postback) Decode

```go
func (o *CGLSetCurrentContext_Postback) Decode(d binary.Decoder) error
```

#### type Capability

```go
type Capability uint32
```

//////////////////////////////////////////////////////////////////////////////
enum Capability
//////////////////////////////////////////////////////////////////////////////

#### func (Capability) String

```go
func (v Capability) String() string
```

#### type CharBufferArray

```go
type CharBufferArray []string
```


#### func (CharBufferArray) Len

```go
func (s CharBufferArray) Len() int
```

#### func (CharBufferArray) Range

```go
func (s CharBufferArray) Range() []string
```

#### type ClearMask

```go
type ClearMask uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ClearMask
//////////////////////////////////////////////////////////////////////////////

#### func (ClearMask) String

```go
func (v ClearMask) String() string
```

#### type ClearState

```go
type ClearState struct {
	binary.Generate
	CreatedAt    atom.ID
	ClearColor   Color
	ClearDepth   float32
	ClearStencil int32
}
```

//////////////////////////////////////////////////////////////////////////////
class ClearState
//////////////////////////////////////////////////////////////////////////////

#### func (*ClearState) Class

```go
func (*ClearState) Class() binary.Class
```

#### func (*ClearState) GetCreatedAt

```go
func (c *ClearState) GetCreatedAt() atom.ID
```

#### func (*ClearState) Init

```go
func (c *ClearState) Init()
```

#### type Color

```go
type Color struct {
	binary.Generate
	CreatedAt atom.ID
	Red       float32
	Green     float32
	Blue      float32
	Alpha     float32
}
```

//////////////////////////////////////////////////////////////////////////////
class Color
//////////////////////////////////////////////////////////////////////////////

#### func (*Color) Class

```go
func (*Color) Class() binary.Class
```

#### func (*Color) GetCreatedAt

```go
func (c *Color) GetCreatedAt() atom.ID
```

#### func (*Color) Init

```go
func (c *Color) Init()
```

#### func (Color) String

```go
func (c Color) String() string
```

#### type CompressedTexelFormat

```go
type CompressedTexelFormat uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat
//////////////////////////////////////////////////////////////////////////////

#### func (CompressedTexelFormat) String

```go
func (v CompressedTexelFormat) String() string
```

#### type CompressedTexelFormat_AMD_compressed_ATC_texture

```go
type CompressedTexelFormat_AMD_compressed_ATC_texture uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat_AMD_compressed_ATC_texture
//////////////////////////////////////////////////////////////////////////////

#### func (CompressedTexelFormat_AMD_compressed_ATC_texture) String

```go
func (v CompressedTexelFormat_AMD_compressed_ATC_texture) String() string
```

#### type CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture

```go
type CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
//////////////////////////////////////////////////////////////////////////////

#### func (CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture) String

```go
func (v CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture) String() string
```

#### type Context

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class Context
//////////////////////////////////////////////////////////////////////////////

#### func (*Context) Class

```go
func (*Context) Class() binary.Class
```

#### func (*Context) GetCreatedAt

```go
func (c *Context) GetCreatedAt() atom.ID
```

#### func (*Context) Init

```go
func (c *Context) Init()
```

#### type ContextArray

```go
type ContextArray []Context
```


#### func (ContextArray) Len

```go
func (s ContextArray) Len() int
```

#### func (ContextArray) Range

```go
func (s ContextArray) Range() []Context
```

#### type ContextID

```go
type ContextID uint32
```


#### func (*ContextID) Equal

```go
func (c *ContextID) Equal(rhs ContextID) bool
```

#### func (*ContextID) Less

```go
func (c *ContextID) Less(rhs ContextID) bool
```

#### type ContextPtr_CGLContextObjMap

```go
type ContextPtr_CGLContextObjMap map[CGLContextObj]*Context
```


#### func (ContextPtr_CGLContextObjMap) Contains

```go
func (m ContextPtr_CGLContextObjMap) Contains(key CGLContextObj) bool
```

#### func (ContextPtr_CGLContextObjMap) Delete

```go
func (m ContextPtr_CGLContextObjMap) Delete(key CGLContextObj)
```

#### func (ContextPtr_CGLContextObjMap) Get

```go
func (m ContextPtr_CGLContextObjMap) Get(key CGLContextObj) *Context
```

#### func (ContextPtr_CGLContextObjMap) Range

```go
func (m ContextPtr_CGLContextObjMap) Range() []*Context
```

#### type ContextPtr_EGLContextMap

```go
type ContextPtr_EGLContextMap map[EGLContext]*Context
```


#### func (ContextPtr_EGLContextMap) Contains

```go
func (m ContextPtr_EGLContextMap) Contains(key EGLContext) bool
```

#### func (ContextPtr_EGLContextMap) Delete

```go
func (m ContextPtr_EGLContextMap) Delete(key EGLContext)
```

#### func (ContextPtr_EGLContextMap) Get

```go
func (m ContextPtr_EGLContextMap) Get(key EGLContext) *Context
```

#### func (ContextPtr_EGLContextMap) Range

```go
func (m ContextPtr_EGLContextMap) Range() []*Context
```

#### type ContextPtr_GLXContextMap

```go
type ContextPtr_GLXContextMap map[GLXContext]*Context
```


#### func (ContextPtr_GLXContextMap) Contains

```go
func (m ContextPtr_GLXContextMap) Contains(key GLXContext) bool
```

#### func (ContextPtr_GLXContextMap) Delete

```go
func (m ContextPtr_GLXContextMap) Delete(key GLXContext)
```

#### func (ContextPtr_GLXContextMap) Get

```go
func (m ContextPtr_GLXContextMap) Get(key GLXContext) *Context
```

#### func (ContextPtr_GLXContextMap) Range

```go
func (m ContextPtr_GLXContextMap) Range() []*Context
```

#### type ContextPtr_HGLRCMap

```go
type ContextPtr_HGLRCMap map[HGLRC]*Context
```


#### func (ContextPtr_HGLRCMap) Contains

```go
func (m ContextPtr_HGLRCMap) Contains(key HGLRC) bool
```

#### func (ContextPtr_HGLRCMap) Delete

```go
func (m ContextPtr_HGLRCMap) Delete(key HGLRC)
```

#### func (ContextPtr_HGLRCMap) Get

```go
func (m ContextPtr_HGLRCMap) Get(key HGLRC) *Context
```

#### func (ContextPtr_HGLRCMap) Range

```go
func (m ContextPtr_HGLRCMap) Range() []*Context
```

#### type ContextPtr_ThreadIDMap

```go
type ContextPtr_ThreadIDMap map[ThreadID]*Context
```


#### func (ContextPtr_ThreadIDMap) Contains

```go
func (m ContextPtr_ThreadIDMap) Contains(key ThreadID) bool
```

#### func (ContextPtr_ThreadIDMap) Delete

```go
func (m ContextPtr_ThreadIDMap) Delete(key ThreadID)
```

#### func (ContextPtr_ThreadIDMap) Get

```go
func (m ContextPtr_ThreadIDMap) Get(key ThreadID) *Context
```

#### func (ContextPtr_ThreadIDMap) Range

```go
func (m ContextPtr_ThreadIDMap) Range() []*Context
```

#### type CubeMapImageTarget

```go
type CubeMapImageTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CubeMapImageTarget
//////////////////////////////////////////////////////////////////////////////

#### func (CubeMapImageTarget) String

```go
func (v CubeMapImageTarget) String() string
```

#### type CubemapLevel

```go
type CubemapLevel struct {
	binary.Generate
	CreatedAt atom.ID
	Faces     Image_CubeMapImageTargetMap
}
```

//////////////////////////////////////////////////////////////////////////////
class CubemapLevel
//////////////////////////////////////////////////////////////////////////////

#### func (*CubemapLevel) Class

```go
func (*CubemapLevel) Class() binary.Class
```

#### func (*CubemapLevel) GetCreatedAt

```go
func (c *CubemapLevel) GetCreatedAt() atom.ID
```

#### func (*CubemapLevel) Init

```go
func (c *CubemapLevel) Init()
```

#### type CubemapLevel_s32Map

```go
type CubemapLevel_s32Map map[int32]CubemapLevel
```


#### func (CubemapLevel_s32Map) Contains

```go
func (m CubemapLevel_s32Map) Contains(key int32) bool
```

#### func (CubemapLevel_s32Map) Delete

```go
func (m CubemapLevel_s32Map) Delete(key int32)
```

#### func (CubemapLevel_s32Map) Get

```go
func (m CubemapLevel_s32Map) Get(key int32) CubemapLevel
```

#### func (CubemapLevel_s32Map) Range

```go
func (m CubemapLevel_s32Map) Range() []CubemapLevel
```

#### type DiscardFramebufferAttachment

```go
type DiscardFramebufferAttachment uint32
```

//////////////////////////////////////////////////////////////////////////////
enum DiscardFramebufferAttachment
//////////////////////////////////////////////////////////////////////////////

#### func (DiscardFramebufferAttachment) String

```go
func (v DiscardFramebufferAttachment) String() string
```

#### type DiscardFramebufferAttachmentArray

```go
type DiscardFramebufferAttachmentArray []DiscardFramebufferAttachment
```


#### func (DiscardFramebufferAttachmentArray) Len

```go
func (s DiscardFramebufferAttachmentArray) Len() int
```

#### func (DiscardFramebufferAttachmentArray) Range

```go
func (s DiscardFramebufferAttachmentArray) Range() []DiscardFramebufferAttachment
```

#### type DrawMode

```go
type DrawMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum DrawMode
//////////////////////////////////////////////////////////////////////////////

#### func (DrawMode) String

```go
func (v DrawMode) String() string
```

#### type EGLBoolean

```go
type EGLBoolean int64
```


#### func (*EGLBoolean) Equal

```go
func (c *EGLBoolean) Equal(rhs EGLBoolean) bool
```

#### func (*EGLBoolean) Less

```go
func (c *EGLBoolean) Less(rhs EGLBoolean) bool
```

#### type EGLConfig

```go
type EGLConfig memory.Pointer
```


#### func (*EGLConfig) Equal

```go
func (c *EGLConfig) Equal(rhs EGLConfig) bool
```

#### func (*EGLConfig) Less

```go
func (c *EGLConfig) Less(rhs EGLConfig) bool
```

#### type EGLContext

```go
type EGLContext memory.Pointer
```


#### func (*EGLContext) Equal

```go
func (c *EGLContext) Equal(rhs EGLContext) bool
```

#### func (*EGLContext) Less

```go
func (c *EGLContext) Less(rhs EGLContext) bool
```

#### type EGLDisplay

```go
type EGLDisplay memory.Pointer
```


#### func (*EGLDisplay) Equal

```go
func (c *EGLDisplay) Equal(rhs EGLDisplay) bool
```

#### func (*EGLDisplay) Less

```go
func (c *EGLDisplay) Less(rhs EGLDisplay) bool
```

#### type EGLSurface

```go
type EGLSurface memory.Pointer
```


#### func (*EGLSurface) Equal

```go
func (c *EGLSurface) Equal(rhs EGLSurface) bool
```

#### func (*EGLSurface) Less

```go
func (c *EGLSurface) Less(rhs EGLSurface) bool
```

#### type EGLint

```go
type EGLint int64
```


#### func (*EGLint) Equal

```go
func (c *EGLint) Equal(rhs EGLint) bool
```

#### func (*EGLint) Less

```go
func (c *EGLint) Less(rhs EGLint) bool
```

#### type EGLintArray

```go
type EGLintArray []EGLint
```


#### func (EGLintArray) Len

```go
func (s EGLintArray) Len() int
```

#### func (EGLintArray) Range

```go
func (s EGLintArray) Range() []EGLint
```

#### type EglCreateContext

```go
type EglCreateContext struct {
	binary.Generate
	Display      EGLDisplay
	Config       EGLConfig
	ShareContext EGLContext
	AttribList   EGLintArray
	Result       EGLContext
}
```

//////////////////////////////////////////////////////////////////////////////
EglCreateContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglCreateContext

```go
func NewEglCreateContext(
	pDisplay EGLDisplay,
	pConfig EGLConfig,
	pShareContext EGLContext,
	pAttribList EGLintArray,
	pResult EGLContext,
) *EglCreateContext
```

#### func (*EglCreateContext) API

```go
func (c *EglCreateContext) API() gfxapi.API
```

#### func (*EglCreateContext) Class

```go
func (*EglCreateContext) Class() binary.Class
```

#### func (*EglCreateContext) Flags

```go
func (c *EglCreateContext) Flags() atom.Flags
```

#### func (*EglCreateContext) Mutate

```go
func (ϟa *EglCreateContext) Mutate(ϟs *gfxapi.State) error
```

#### func (*EglCreateContext) Replay

```go
func (ω *EglCreateContext) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*EglCreateContext) String

```go
func (c *EglCreateContext) String() string
```

#### func (*EglCreateContext) TypeID

```go
func (c *EglCreateContext) TypeID() atom.TypeID
```

#### type EglCreateContext_Postback

```go
type EglCreateContext_Postback struct {
	Result []byte
}
```


#### func (*EglCreateContext_Postback) Decode

```go
func (o *EglCreateContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type EglInitialize

```go
type EglInitialize struct {
	binary.Generate
	Dpy    EGLDisplay
	Major  EGLint
	Minor  EGLint
	Result EGLBoolean
}
```

//////////////////////////////////////////////////////////////////////////////
EglInitialize
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglInitialize

```go
func NewEglInitialize(
	pDpy EGLDisplay,
	pMajor EGLint,
	pMinor EGLint,
	pResult EGLBoolean,
) *EglInitialize
```

#### func (*EglInitialize) API

```go
func (c *EglInitialize) API() gfxapi.API
```

#### func (*EglInitialize) Class

```go
func (*EglInitialize) Class() binary.Class
```

#### func (*EglInitialize) Flags

```go
func (c *EglInitialize) Flags() atom.Flags
```

#### func (*EglInitialize) Mutate

```go
func (ϟa *EglInitialize) Mutate(ϟs *gfxapi.State) error
```

#### func (*EglInitialize) String

```go
func (c *EglInitialize) String() string
```

#### func (*EglInitialize) TypeID

```go
func (c *EglInitialize) TypeID() atom.TypeID
```

#### type EglInitialize_Postback

```go
type EglInitialize_Postback struct {
	Major  EGLint
	Minor  EGLint
	Result EGLBoolean
}
```


#### func (*EglInitialize_Postback) Decode

```go
func (o *EglInitialize_Postback) Decode(d binary.Decoder) error
```

#### type EglMakeCurrent

```go
type EglMakeCurrent struct {
	binary.Generate
	Display EGLDisplay
	Draw    EGLSurface
	Read    EGLSurface
	Context EGLContext
	Result  EGLBoolean
}
```

//////////////////////////////////////////////////////////////////////////////
EglMakeCurrent
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglMakeCurrent

```go
func NewEglMakeCurrent(
	pDisplay EGLDisplay,
	pDraw EGLSurface,
	pRead EGLSurface,
	pContext EGLContext,
	pResult EGLBoolean,
) *EglMakeCurrent
```

#### func (*EglMakeCurrent) API

```go
func (c *EglMakeCurrent) API() gfxapi.API
```

#### func (*EglMakeCurrent) Class

```go
func (*EglMakeCurrent) Class() binary.Class
```

#### func (*EglMakeCurrent) Flags

```go
func (c *EglMakeCurrent) Flags() atom.Flags
```

#### func (*EglMakeCurrent) Mutate

```go
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State) error
```

#### func (*EglMakeCurrent) Replay

```go
func (ω *EglMakeCurrent) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*EglMakeCurrent) String

```go
func (c *EglMakeCurrent) String() string
```

#### func (*EglMakeCurrent) TypeID

```go
func (c *EglMakeCurrent) TypeID() atom.TypeID
```

#### type EglMakeCurrent_Postback

```go
type EglMakeCurrent_Postback struct {
	Result EGLBoolean
}
```


#### func (*EglMakeCurrent_Postback) Decode

```go
func (o *EglMakeCurrent_Postback) Decode(d binary.Decoder) error
```

#### type EglQuerySurface

```go
type EglQuerySurface struct {
	binary.Generate
	Display   EGLDisplay
	Surface   EGLSurface
	Attribute EGLint
	Value     EGLint
	Result    EGLBoolean
}
```

//////////////////////////////////////////////////////////////////////////////
EglQuerySurface
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglQuerySurface

```go
func NewEglQuerySurface(
	pDisplay EGLDisplay,
	pSurface EGLSurface,
	pAttribute EGLint,
	pValue EGLint,
	pResult EGLBoolean,
) *EglQuerySurface
```

#### func (*EglQuerySurface) API

```go
func (c *EglQuerySurface) API() gfxapi.API
```

#### func (*EglQuerySurface) Class

```go
func (*EglQuerySurface) Class() binary.Class
```

#### func (*EglQuerySurface) Flags

```go
func (c *EglQuerySurface) Flags() atom.Flags
```

#### func (*EglQuerySurface) Mutate

```go
func (ϟa *EglQuerySurface) Mutate(ϟs *gfxapi.State) error
```

#### func (*EglQuerySurface) String

```go
func (c *EglQuerySurface) String() string
```

#### func (*EglQuerySurface) TypeID

```go
func (c *EglQuerySurface) TypeID() atom.TypeID
```

#### type EglQuerySurface_Postback

```go
type EglQuerySurface_Postback struct {
	Value  EGLint
	Result EGLBoolean
}
```


#### func (*EglQuerySurface_Postback) Decode

```go
func (o *EglQuerySurface_Postback) Decode(d binary.Decoder) error
```

#### type EglSwapBuffers

```go
type EglSwapBuffers struct {
	binary.Generate
	Display EGLDisplay
	Surface memory.Pointer
	Result  EGLBoolean
}
```

//////////////////////////////////////////////////////////////////////////////
EglSwapBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglSwapBuffers

```go
func NewEglSwapBuffers(
	pDisplay EGLDisplay,
	pSurface memory.Pointer,
	pResult EGLBoolean,
) *EglSwapBuffers
```

#### func (*EglSwapBuffers) API

```go
func (c *EglSwapBuffers) API() gfxapi.API
```

#### func (*EglSwapBuffers) Class

```go
func (*EglSwapBuffers) Class() binary.Class
```

#### func (*EglSwapBuffers) Flags

```go
func (c *EglSwapBuffers) Flags() atom.Flags
```

#### func (*EglSwapBuffers) Mutate

```go
func (ϟa *EglSwapBuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*EglSwapBuffers) String

```go
func (c *EglSwapBuffers) String() string
```

#### func (*EglSwapBuffers) TypeID

```go
func (c *EglSwapBuffers) TypeID() atom.TypeID
```

#### type EglSwapBuffers_Postback

```go
type EglSwapBuffers_Postback struct {
	Result EGLBoolean
}
```


#### func (*EglSwapBuffers_Postback) Decode

```go
func (o *EglSwapBuffers_Postback) Decode(d binary.Decoder) error
```

#### type Error

```go
type Error uint32
```

//////////////////////////////////////////////////////////////////////////////
enum Error
//////////////////////////////////////////////////////////////////////////////

#### func (Error) String

```go
func (v Error) String() string
```

#### type F32Array

```go
type F32Array []float32
```


#### func (F32Array) Len

```go
func (s F32Array) Len() int
```

#### func (F32Array) Range

```go
func (s F32Array) Range() []float32
```

#### type FaceMode

```go
type FaceMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FaceMode
//////////////////////////////////////////////////////////////////////////////

#### func (FaceMode) String

```go
func (v FaceMode) String() string
```

#### type FaceOrientation

```go
type FaceOrientation uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FaceOrientation
//////////////////////////////////////////////////////////////////////////////

#### func (FaceOrientation) String

```go
func (v FaceOrientation) String() string
```

#### type FlushPostBuffer

```go
type FlushPostBuffer struct {
	binary.Generate
}
```

//////////////////////////////////////////////////////////////////////////////
FlushPostBuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewFlushPostBuffer

```go
func NewFlushPostBuffer() *FlushPostBuffer
```

#### func (*FlushPostBuffer) API

```go
func (c *FlushPostBuffer) API() gfxapi.API
```

#### func (*FlushPostBuffer) Class

```go
func (*FlushPostBuffer) Class() binary.Class
```

#### func (*FlushPostBuffer) Flags

```go
func (c *FlushPostBuffer) Flags() atom.Flags
```

#### func (*FlushPostBuffer) Mutate

```go
func (ϟa *FlushPostBuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*FlushPostBuffer) Replay

```go
func (ϟa *FlushPostBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*FlushPostBuffer) String

```go
func (c *FlushPostBuffer) String() string
```

#### func (*FlushPostBuffer) TypeID

```go
func (c *FlushPostBuffer) TypeID() atom.TypeID
```

#### type Framebuffer

```go
type Framebuffer struct {
	binary.Generate
	CreatedAt   atom.ID
	Attachments FramebufferAttachmentInfo_FramebufferAttachmentMap
}
```

//////////////////////////////////////////////////////////////////////////////
class Framebuffer
//////////////////////////////////////////////////////////////////////////////

#### func (*Framebuffer) Class

```go
func (*Framebuffer) Class() binary.Class
```

#### func (*Framebuffer) GetCreatedAt

```go
func (c *Framebuffer) GetCreatedAt() atom.ID
```

#### func (*Framebuffer) Init

```go
func (c *Framebuffer) Init()
```

#### type FramebufferArray

```go
type FramebufferArray []Framebuffer
```


#### func (FramebufferArray) Len

```go
func (s FramebufferArray) Len() int
```

#### func (FramebufferArray) Range

```go
func (s FramebufferArray) Range() []Framebuffer
```

#### type FramebufferAttachable

```go
type FramebufferAttachable interface {
}
```

//////////////////////////////////////////////////////////////////////////////
class FramebufferAttachable
//////////////////////////////////////////////////////////////////////////////

#### type FramebufferAttachment

```go
type FramebufferAttachment uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferAttachment
//////////////////////////////////////////////////////////////////////////////

#### func (FramebufferAttachment) String

```go
func (v FramebufferAttachment) String() string
```

#### type FramebufferAttachmentArray

```go
type FramebufferAttachmentArray []FramebufferAttachment
```


#### func (FramebufferAttachmentArray) Len

```go
func (s FramebufferAttachmentArray) Len() int
```

#### func (FramebufferAttachmentArray) Range

```go
func (s FramebufferAttachmentArray) Range() []FramebufferAttachment
```

#### type FramebufferAttachmentInfo

```go
type FramebufferAttachmentInfo struct {
	binary.Generate
	CreatedAt    atom.ID
	Object       uint32
	Type         FramebufferAttachmentType
	TextureLevel int32
	CubeMapFace  CubeMapImageTarget
}
```

//////////////////////////////////////////////////////////////////////////////
class FramebufferAttachmentInfo
//////////////////////////////////////////////////////////////////////////////

#### func (*FramebufferAttachmentInfo) Class

```go
func (*FramebufferAttachmentInfo) Class() binary.Class
```

#### func (*FramebufferAttachmentInfo) GetCreatedAt

```go
func (c *FramebufferAttachmentInfo) GetCreatedAt() atom.ID
```

#### func (*FramebufferAttachmentInfo) Init

```go
func (c *FramebufferAttachmentInfo) Init()
```

#### type FramebufferAttachmentInfo_FramebufferAttachmentMap

```go
type FramebufferAttachmentInfo_FramebufferAttachmentMap map[FramebufferAttachment]FramebufferAttachmentInfo
```


#### func (FramebufferAttachmentInfo_FramebufferAttachmentMap) Contains

```go
func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Contains(key FramebufferAttachment) bool
```

#### func (FramebufferAttachmentInfo_FramebufferAttachmentMap) Delete

```go
func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Delete(key FramebufferAttachment)
```

#### func (FramebufferAttachmentInfo_FramebufferAttachmentMap) Get

```go
func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Get(key FramebufferAttachment) FramebufferAttachmentInfo
```

#### func (FramebufferAttachmentInfo_FramebufferAttachmentMap) Range

```go
func (m FramebufferAttachmentInfo_FramebufferAttachmentMap) Range() []FramebufferAttachmentInfo
```

#### type FramebufferAttachmentParameter

```go
type FramebufferAttachmentParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferAttachmentParameter
//////////////////////////////////////////////////////////////////////////////

#### func (FramebufferAttachmentParameter) String

```go
func (v FramebufferAttachmentParameter) String() string
```

#### type FramebufferAttachmentType

```go
type FramebufferAttachmentType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferAttachmentType
//////////////////////////////////////////////////////////////////////////////

#### func (FramebufferAttachmentType) String

```go
func (v FramebufferAttachmentType) String() string
```

#### type FramebufferId

```go
type FramebufferId uint32
```


#### func (*FramebufferId) Equal

```go
func (c *FramebufferId) Equal(rhs FramebufferId) bool
```

#### func (*FramebufferId) Less

```go
func (c *FramebufferId) Less(rhs FramebufferId) bool
```

#### type FramebufferIdArray

```go
type FramebufferIdArray []FramebufferId
```


#### func (FramebufferIdArray) Len

```go
func (s FramebufferIdArray) Len() int
```

#### func (FramebufferIdArray) Range

```go
func (s FramebufferIdArray) Range() []FramebufferId
```

#### type FramebufferId_FramebufferTargetMap

```go
type FramebufferId_FramebufferTargetMap map[FramebufferTarget]FramebufferId
```


#### func (FramebufferId_FramebufferTargetMap) Contains

```go
func (m FramebufferId_FramebufferTargetMap) Contains(key FramebufferTarget) bool
```

#### func (FramebufferId_FramebufferTargetMap) Delete

```go
func (m FramebufferId_FramebufferTargetMap) Delete(key FramebufferTarget)
```

#### func (FramebufferId_FramebufferTargetMap) Get

```go
func (m FramebufferId_FramebufferTargetMap) Get(key FramebufferTarget) FramebufferId
```

#### func (FramebufferId_FramebufferTargetMap) Range

```go
func (m FramebufferId_FramebufferTargetMap) Range() []FramebufferId
```

#### type FramebufferPtr_FramebufferIdMap

```go
type FramebufferPtr_FramebufferIdMap map[FramebufferId]*Framebuffer
```


#### func (FramebufferPtr_FramebufferIdMap) Contains

```go
func (m FramebufferPtr_FramebufferIdMap) Contains(key FramebufferId) bool
```

#### func (FramebufferPtr_FramebufferIdMap) Delete

```go
func (m FramebufferPtr_FramebufferIdMap) Delete(key FramebufferId)
```

#### func (FramebufferPtr_FramebufferIdMap) Get

```go
func (m FramebufferPtr_FramebufferIdMap) Get(key FramebufferId) *Framebuffer
```

#### func (FramebufferPtr_FramebufferIdMap) Range

```go
func (m FramebufferPtr_FramebufferIdMap) Range() []*Framebuffer
```

#### type FramebufferStatus

```go
type FramebufferStatus uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferStatus
//////////////////////////////////////////////////////////////////////////////

#### func (FramebufferStatus) String

```go
func (v FramebufferStatus) String() string
```

#### type FramebufferTarget

```go
type FramebufferTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferTarget
//////////////////////////////////////////////////////////////////////////////

#### func (FramebufferTarget) String

```go
func (v FramebufferTarget) String() string
```

#### type FramebufferTarget_GLES_2_0

```go
type FramebufferTarget_GLES_2_0 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferTarget_GLES_2_0
//////////////////////////////////////////////////////////////////////////////

#### func (FramebufferTarget_GLES_2_0) String

```go
func (v FramebufferTarget_GLES_2_0) String() string
```

#### type FramebufferTarget_GLES_3_1

```go
type FramebufferTarget_GLES_3_1 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferTarget_GLES_3_1
//////////////////////////////////////////////////////////////////////////////

#### func (FramebufferTarget_GLES_3_1) String

```go
func (v FramebufferTarget_GLES_3_1) String() string
```

#### type GLXContext

```go
type GLXContext memory.Pointer
```


#### func (*GLXContext) Equal

```go
func (c *GLXContext) Equal(rhs GLXContext) bool
```

#### func (*GLXContext) Less

```go
func (c *GLXContext) Less(rhs GLXContext) bool
```

#### type GLXDrawable

```go
type GLXDrawable memory.Pointer
```


#### func (*GLXDrawable) Equal

```go
func (c *GLXDrawable) Equal(rhs GLXDrawable) bool
```

#### func (*GLXDrawable) Less

```go
func (c *GLXDrawable) Less(rhs GLXDrawable) bool
```

#### type GlActiveTexture

```go
type GlActiveTexture struct {
	binary.Generate
	Unit TextureUnit
}
```

//////////////////////////////////////////////////////////////////////////////
GlActiveTexture
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlActiveTexture

```go
func NewGlActiveTexture(
	pUnit TextureUnit,
) *GlActiveTexture
```

#### func (*GlActiveTexture) API

```go
func (c *GlActiveTexture) API() gfxapi.API
```

#### func (*GlActiveTexture) Class

```go
func (*GlActiveTexture) Class() binary.Class
```

#### func (*GlActiveTexture) Flags

```go
func (c *GlActiveTexture) Flags() atom.Flags
```

#### func (*GlActiveTexture) Mutate

```go
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlActiveTexture) Replay

```go
func (ϟa *GlActiveTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlActiveTexture) String

```go
func (c *GlActiveTexture) String() string
```

#### func (*GlActiveTexture) TypeID

```go
func (c *GlActiveTexture) TypeID() atom.TypeID
```

#### type GlAttachShader

```go
type GlAttachShader struct {
	binary.Generate
	Program ProgramId
	Shader  ShaderId
}
```

//////////////////////////////////////////////////////////////////////////////
GlAttachShader
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlAttachShader

```go
func NewGlAttachShader(
	pProgram ProgramId,
	pShader ShaderId,
) *GlAttachShader
```

#### func (*GlAttachShader) API

```go
func (c *GlAttachShader) API() gfxapi.API
```

#### func (*GlAttachShader) Class

```go
func (*GlAttachShader) Class() binary.Class
```

#### func (*GlAttachShader) Flags

```go
func (c *GlAttachShader) Flags() atom.Flags
```

#### func (*GlAttachShader) Mutate

```go
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlAttachShader) Replay

```go
func (ϟa *GlAttachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlAttachShader) String

```go
func (c *GlAttachShader) String() string
```

#### func (*GlAttachShader) TypeID

```go
func (c *GlAttachShader) TypeID() atom.TypeID
```

#### type GlBeginQuery

```go
type GlBeginQuery struct {
	binary.Generate
	Target QueryTarget
	Query  QueryId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBeginQuery
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBeginQuery

```go
func NewGlBeginQuery(
	pTarget QueryTarget,
	pQuery QueryId,
) *GlBeginQuery
```

#### func (*GlBeginQuery) API

```go
func (c *GlBeginQuery) API() gfxapi.API
```

#### func (*GlBeginQuery) Class

```go
func (*GlBeginQuery) Class() binary.Class
```

#### func (*GlBeginQuery) Flags

```go
func (c *GlBeginQuery) Flags() atom.Flags
```

#### func (*GlBeginQuery) Mutate

```go
func (ϟa *GlBeginQuery) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBeginQuery) Replay

```go
func (ϟa *GlBeginQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBeginQuery) String

```go
func (c *GlBeginQuery) String() string
```

#### func (*GlBeginQuery) TypeID

```go
func (c *GlBeginQuery) TypeID() atom.TypeID
```

#### type GlBeginQueryEXT

```go
type GlBeginQueryEXT struct {
	binary.Generate
	Target QueryTarget
	Query  QueryId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBeginQueryEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBeginQueryEXT

```go
func NewGlBeginQueryEXT(
	pTarget QueryTarget,
	pQuery QueryId,
) *GlBeginQueryEXT
```

#### func (*GlBeginQueryEXT) API

```go
func (c *GlBeginQueryEXT) API() gfxapi.API
```

#### func (*GlBeginQueryEXT) Class

```go
func (*GlBeginQueryEXT) Class() binary.Class
```

#### func (*GlBeginQueryEXT) Flags

```go
func (c *GlBeginQueryEXT) Flags() atom.Flags
```

#### func (*GlBeginQueryEXT) Mutate

```go
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBeginQueryEXT) Replay

```go
func (ϟa *GlBeginQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBeginQueryEXT) String

```go
func (c *GlBeginQueryEXT) String() string
```

#### func (*GlBeginQueryEXT) TypeID

```go
func (c *GlBeginQueryEXT) TypeID() atom.TypeID
```

#### type GlBindAttribLocation

```go
type GlBindAttribLocation struct {
	binary.Generate
	Program  ProgramId
	Location AttributeLocation
	Name     string
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindAttribLocation
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindAttribLocation

```go
func NewGlBindAttribLocation(
	pProgram ProgramId,
	pLocation AttributeLocation,
	pName string,
) *GlBindAttribLocation
```

#### func (*GlBindAttribLocation) API

```go
func (c *GlBindAttribLocation) API() gfxapi.API
```

#### func (*GlBindAttribLocation) Class

```go
func (*GlBindAttribLocation) Class() binary.Class
```

#### func (*GlBindAttribLocation) Flags

```go
func (c *GlBindAttribLocation) Flags() atom.Flags
```

#### func (*GlBindAttribLocation) Mutate

```go
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBindAttribLocation) Replay

```go
func (ϟa *GlBindAttribLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBindAttribLocation) String

```go
func (c *GlBindAttribLocation) String() string
```

#### func (*GlBindAttribLocation) TypeID

```go
func (c *GlBindAttribLocation) TypeID() atom.TypeID
```

#### type GlBindBuffer

```go
type GlBindBuffer struct {
	binary.Generate
	Target BufferTarget
	Buffer BufferId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindBuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindBuffer

```go
func NewGlBindBuffer(
	pTarget BufferTarget,
	pBuffer BufferId,
) *GlBindBuffer
```

#### func (*GlBindBuffer) API

```go
func (c *GlBindBuffer) API() gfxapi.API
```

#### func (*GlBindBuffer) Class

```go
func (*GlBindBuffer) Class() binary.Class
```

#### func (*GlBindBuffer) Flags

```go
func (c *GlBindBuffer) Flags() atom.Flags
```

#### func (*GlBindBuffer) Mutate

```go
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBindBuffer) Replay

```go
func (ϟa *GlBindBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBindBuffer) String

```go
func (c *GlBindBuffer) String() string
```

#### func (*GlBindBuffer) TypeID

```go
func (c *GlBindBuffer) TypeID() atom.TypeID
```

#### type GlBindFramebuffer

```go
type GlBindFramebuffer struct {
	binary.Generate
	Target      FramebufferTarget
	Framebuffer FramebufferId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindFramebuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindFramebuffer

```go
func NewGlBindFramebuffer(
	pTarget FramebufferTarget,
	pFramebuffer FramebufferId,
) *GlBindFramebuffer
```

#### func (*GlBindFramebuffer) API

```go
func (c *GlBindFramebuffer) API() gfxapi.API
```

#### func (*GlBindFramebuffer) Class

```go
func (*GlBindFramebuffer) Class() binary.Class
```

#### func (*GlBindFramebuffer) Flags

```go
func (c *GlBindFramebuffer) Flags() atom.Flags
```

#### func (*GlBindFramebuffer) Mutate

```go
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBindFramebuffer) Replay

```go
func (ϟa *GlBindFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBindFramebuffer) String

```go
func (c *GlBindFramebuffer) String() string
```

#### func (*GlBindFramebuffer) TypeID

```go
func (c *GlBindFramebuffer) TypeID() atom.TypeID
```

#### type GlBindRenderbuffer

```go
type GlBindRenderbuffer struct {
	binary.Generate
	Target       RenderbufferTarget
	Renderbuffer RenderbufferId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindRenderbuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindRenderbuffer

```go
func NewGlBindRenderbuffer(
	pTarget RenderbufferTarget,
	pRenderbuffer RenderbufferId,
) *GlBindRenderbuffer
```

#### func (*GlBindRenderbuffer) API

```go
func (c *GlBindRenderbuffer) API() gfxapi.API
```

#### func (*GlBindRenderbuffer) Class

```go
func (*GlBindRenderbuffer) Class() binary.Class
```

#### func (*GlBindRenderbuffer) Flags

```go
func (c *GlBindRenderbuffer) Flags() atom.Flags
```

#### func (*GlBindRenderbuffer) Mutate

```go
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBindRenderbuffer) Replay

```go
func (ϟa *GlBindRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBindRenderbuffer) String

```go
func (c *GlBindRenderbuffer) String() string
```

#### func (*GlBindRenderbuffer) TypeID

```go
func (c *GlBindRenderbuffer) TypeID() atom.TypeID
```

#### type GlBindTexture

```go
type GlBindTexture struct {
	binary.Generate
	Target  TextureTarget
	Texture TextureId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindTexture
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindTexture

```go
func NewGlBindTexture(
	pTarget TextureTarget,
	pTexture TextureId,
) *GlBindTexture
```

#### func (*GlBindTexture) API

```go
func (c *GlBindTexture) API() gfxapi.API
```

#### func (*GlBindTexture) Class

```go
func (*GlBindTexture) Class() binary.Class
```

#### func (*GlBindTexture) Flags

```go
func (c *GlBindTexture) Flags() atom.Flags
```

#### func (*GlBindTexture) Mutate

```go
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBindTexture) Replay

```go
func (ϟa *GlBindTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBindTexture) String

```go
func (c *GlBindTexture) String() string
```

#### func (*GlBindTexture) TypeID

```go
func (c *GlBindTexture) TypeID() atom.TypeID
```

#### type GlBindVertexArrayOES

```go
type GlBindVertexArrayOES struct {
	binary.Generate
	Array VertexArrayId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindVertexArrayOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindVertexArrayOES

```go
func NewGlBindVertexArrayOES(
	pArray VertexArrayId,
) *GlBindVertexArrayOES
```

#### func (*GlBindVertexArrayOES) API

```go
func (c *GlBindVertexArrayOES) API() gfxapi.API
```

#### func (*GlBindVertexArrayOES) Class

```go
func (*GlBindVertexArrayOES) Class() binary.Class
```

#### func (*GlBindVertexArrayOES) Flags

```go
func (c *GlBindVertexArrayOES) Flags() atom.Flags
```

#### func (*GlBindVertexArrayOES) Mutate

```go
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBindVertexArrayOES) Replay

```go
func (ϟa *GlBindVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBindVertexArrayOES) String

```go
func (c *GlBindVertexArrayOES) String() string
```

#### func (*GlBindVertexArrayOES) TypeID

```go
func (c *GlBindVertexArrayOES) TypeID() atom.TypeID
```

#### type GlBlendColor

```go
type GlBlendColor struct {
	binary.Generate
	Red   float32
	Green float32
	Blue  float32
	Alpha float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlBlendColor
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBlendColor

```go
func NewGlBlendColor(
	pRed float32,
	pGreen float32,
	pBlue float32,
	pAlpha float32,
) *GlBlendColor
```

#### func (*GlBlendColor) API

```go
func (c *GlBlendColor) API() gfxapi.API
```

#### func (*GlBlendColor) Class

```go
func (*GlBlendColor) Class() binary.Class
```

#### func (*GlBlendColor) Flags

```go
func (c *GlBlendColor) Flags() atom.Flags
```

#### func (*GlBlendColor) Mutate

```go
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBlendColor) Replay

```go
func (ϟa *GlBlendColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBlendColor) String

```go
func (c *GlBlendColor) String() string
```

#### func (*GlBlendColor) TypeID

```go
func (c *GlBlendColor) TypeID() atom.TypeID
```

#### type GlBlendEquation

```go
type GlBlendEquation struct {
	binary.Generate
	Equation BlendEquation
}
```

//////////////////////////////////////////////////////////////////////////////
GlBlendEquation
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBlendEquation

```go
func NewGlBlendEquation(
	pEquation BlendEquation,
) *GlBlendEquation
```

#### func (*GlBlendEquation) API

```go
func (c *GlBlendEquation) API() gfxapi.API
```

#### func (*GlBlendEquation) Class

```go
func (*GlBlendEquation) Class() binary.Class
```

#### func (*GlBlendEquation) Flags

```go
func (c *GlBlendEquation) Flags() atom.Flags
```

#### func (*GlBlendEquation) Mutate

```go
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBlendEquation) Replay

```go
func (ϟa *GlBlendEquation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBlendEquation) String

```go
func (c *GlBlendEquation) String() string
```

#### func (*GlBlendEquation) TypeID

```go
func (c *GlBlendEquation) TypeID() atom.TypeID
```

#### type GlBlendEquationSeparate

```go
type GlBlendEquationSeparate struct {
	binary.Generate
	Rgb   BlendEquation
	Alpha BlendEquation
}
```

//////////////////////////////////////////////////////////////////////////////
GlBlendEquationSeparate
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBlendEquationSeparate

```go
func NewGlBlendEquationSeparate(
	pRgb BlendEquation,
	pAlpha BlendEquation,
) *GlBlendEquationSeparate
```

#### func (*GlBlendEquationSeparate) API

```go
func (c *GlBlendEquationSeparate) API() gfxapi.API
```

#### func (*GlBlendEquationSeparate) Class

```go
func (*GlBlendEquationSeparate) Class() binary.Class
```

#### func (*GlBlendEquationSeparate) Flags

```go
func (c *GlBlendEquationSeparate) Flags() atom.Flags
```

#### func (*GlBlendEquationSeparate) Mutate

```go
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBlendEquationSeparate) Replay

```go
func (ϟa *GlBlendEquationSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBlendEquationSeparate) String

```go
func (c *GlBlendEquationSeparate) String() string
```

#### func (*GlBlendEquationSeparate) TypeID

```go
func (c *GlBlendEquationSeparate) TypeID() atom.TypeID
```

#### type GlBlendFunc

```go
type GlBlendFunc struct {
	binary.Generate
	SrcFactor BlendFactor
	DstFactor BlendFactor
}
```

//////////////////////////////////////////////////////////////////////////////
GlBlendFunc
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBlendFunc

```go
func NewGlBlendFunc(
	pSrcFactor BlendFactor,
	pDstFactor BlendFactor,
) *GlBlendFunc
```

#### func (*GlBlendFunc) API

```go
func (c *GlBlendFunc) API() gfxapi.API
```

#### func (*GlBlendFunc) Class

```go
func (*GlBlendFunc) Class() binary.Class
```

#### func (*GlBlendFunc) Flags

```go
func (c *GlBlendFunc) Flags() atom.Flags
```

#### func (*GlBlendFunc) Mutate

```go
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBlendFunc) Replay

```go
func (ϟa *GlBlendFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBlendFunc) String

```go
func (c *GlBlendFunc) String() string
```

#### func (*GlBlendFunc) TypeID

```go
func (c *GlBlendFunc) TypeID() atom.TypeID
```

#### type GlBlendFuncSeparate

```go
type GlBlendFuncSeparate struct {
	binary.Generate
	SrcFactorRgb   BlendFactor
	DstFactorRgb   BlendFactor
	SrcFactorAlpha BlendFactor
	DstFactorAlpha BlendFactor
}
```

//////////////////////////////////////////////////////////////////////////////
GlBlendFuncSeparate
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBlendFuncSeparate

```go
func NewGlBlendFuncSeparate(
	pSrcFactorRgb BlendFactor,
	pDstFactorRgb BlendFactor,
	pSrcFactorAlpha BlendFactor,
	pDstFactorAlpha BlendFactor,
) *GlBlendFuncSeparate
```

#### func (*GlBlendFuncSeparate) API

```go
func (c *GlBlendFuncSeparate) API() gfxapi.API
```

#### func (*GlBlendFuncSeparate) Class

```go
func (*GlBlendFuncSeparate) Class() binary.Class
```

#### func (*GlBlendFuncSeparate) Flags

```go
func (c *GlBlendFuncSeparate) Flags() atom.Flags
```

#### func (*GlBlendFuncSeparate) Mutate

```go
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBlendFuncSeparate) Replay

```go
func (ϟa *GlBlendFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBlendFuncSeparate) String

```go
func (c *GlBlendFuncSeparate) String() string
```

#### func (*GlBlendFuncSeparate) TypeID

```go
func (c *GlBlendFuncSeparate) TypeID() atom.TypeID
```

#### type GlBlitFramebuffer

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlBlitFramebuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBlitFramebuffer

```go
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
) *GlBlitFramebuffer
```

#### func (*GlBlitFramebuffer) API

```go
func (c *GlBlitFramebuffer) API() gfxapi.API
```

#### func (*GlBlitFramebuffer) Class

```go
func (*GlBlitFramebuffer) Class() binary.Class
```

#### func (*GlBlitFramebuffer) Flags

```go
func (c *GlBlitFramebuffer) Flags() atom.Flags
```

#### func (*GlBlitFramebuffer) Mutate

```go
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBlitFramebuffer) Replay

```go
func (ϟa *GlBlitFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBlitFramebuffer) String

```go
func (c *GlBlitFramebuffer) String() string
```

#### func (*GlBlitFramebuffer) TypeID

```go
func (c *GlBlitFramebuffer) TypeID() atom.TypeID
```

#### type GlBufferData

```go
type GlBufferData struct {
	binary.Generate
	Target BufferTarget
	Size   int32
	Data   BufferDataPointer
	Usage  BufferUsage
}
```

//////////////////////////////////////////////////////////////////////////////
GlBufferData
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBufferData

```go
func NewGlBufferData(
	pTarget BufferTarget,
	pSize int32,
	pData BufferDataPointer,
	pUsage BufferUsage,
) *GlBufferData
```

#### func (*GlBufferData) API

```go
func (c *GlBufferData) API() gfxapi.API
```

#### func (*GlBufferData) Class

```go
func (*GlBufferData) Class() binary.Class
```

#### func (*GlBufferData) Flags

```go
func (c *GlBufferData) Flags() atom.Flags
```

#### func (*GlBufferData) Mutate

```go
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBufferData) Replay

```go
func (ϟa *GlBufferData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBufferData) String

```go
func (c *GlBufferData) String() string
```

#### func (*GlBufferData) TypeID

```go
func (c *GlBufferData) TypeID() atom.TypeID
```

#### type GlBufferSubData

```go
type GlBufferSubData struct {
	binary.Generate
	Target BufferTarget
	Offset int32
	Size   int32
	Data   memory.Pointer
}
```

//////////////////////////////////////////////////////////////////////////////
GlBufferSubData
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBufferSubData

```go
func NewGlBufferSubData(
	pTarget BufferTarget,
	pOffset int32,
	pSize int32,
	pData memory.Pointer,
) *GlBufferSubData
```

#### func (*GlBufferSubData) API

```go
func (c *GlBufferSubData) API() gfxapi.API
```

#### func (*GlBufferSubData) Class

```go
func (*GlBufferSubData) Class() binary.Class
```

#### func (*GlBufferSubData) Flags

```go
func (c *GlBufferSubData) Flags() atom.Flags
```

#### func (*GlBufferSubData) Mutate

```go
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlBufferSubData) Replay

```go
func (ϟa *GlBufferSubData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlBufferSubData) String

```go
func (c *GlBufferSubData) String() string
```

#### func (*GlBufferSubData) TypeID

```go
func (c *GlBufferSubData) TypeID() atom.TypeID
```

#### type GlCheckFramebufferStatus

```go
type GlCheckFramebufferStatus struct {
	binary.Generate
	Target FramebufferTarget
	Result FramebufferStatus
}
```

//////////////////////////////////////////////////////////////////////////////
GlCheckFramebufferStatus
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCheckFramebufferStatus

```go
func NewGlCheckFramebufferStatus(
	pTarget FramebufferTarget,
	pResult FramebufferStatus,
) *GlCheckFramebufferStatus
```

#### func (*GlCheckFramebufferStatus) API

```go
func (c *GlCheckFramebufferStatus) API() gfxapi.API
```

#### func (*GlCheckFramebufferStatus) Class

```go
func (*GlCheckFramebufferStatus) Class() binary.Class
```

#### func (*GlCheckFramebufferStatus) Flags

```go
func (c *GlCheckFramebufferStatus) Flags() atom.Flags
```

#### func (*GlCheckFramebufferStatus) Mutate

```go
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCheckFramebufferStatus) Replay

```go
func (ϟa *GlCheckFramebufferStatus) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCheckFramebufferStatus) String

```go
func (c *GlCheckFramebufferStatus) String() string
```

#### func (*GlCheckFramebufferStatus) TypeID

```go
func (c *GlCheckFramebufferStatus) TypeID() atom.TypeID
```

#### type GlCheckFramebufferStatus_Postback

```go
type GlCheckFramebufferStatus_Postback struct {
	Result FramebufferStatus
}
```


#### func (*GlCheckFramebufferStatus_Postback) Decode

```go
func (o *GlCheckFramebufferStatus_Postback) Decode(d binary.Decoder) error
```

#### type GlClear

```go
type GlClear struct {
	binary.Generate
	Mask ClearMask
}
```

//////////////////////////////////////////////////////////////////////////////
GlClear
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlClear

```go
func NewGlClear(
	pMask ClearMask,
) *GlClear
```

#### func (*GlClear) API

```go
func (c *GlClear) API() gfxapi.API
```

#### func (*GlClear) Class

```go
func (*GlClear) Class() binary.Class
```

#### func (*GlClear) Flags

```go
func (c *GlClear) Flags() atom.Flags
```

#### func (*GlClear) Mutate

```go
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlClear) Replay

```go
func (ϟa *GlClear) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlClear) String

```go
func (c *GlClear) String() string
```

#### func (*GlClear) TypeID

```go
func (c *GlClear) TypeID() atom.TypeID
```

#### type GlClearColor

```go
type GlClearColor struct {
	binary.Generate
	R float32
	G float32
	B float32
	A float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlClearColor
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlClearColor

```go
func NewGlClearColor(
	pR float32,
	pG float32,
	pB float32,
	pA float32,
) *GlClearColor
```

#### func (*GlClearColor) API

```go
func (c *GlClearColor) API() gfxapi.API
```

#### func (*GlClearColor) Class

```go
func (*GlClearColor) Class() binary.Class
```

#### func (*GlClearColor) Flags

```go
func (c *GlClearColor) Flags() atom.Flags
```

#### func (*GlClearColor) Mutate

```go
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlClearColor) Replay

```go
func (ϟa *GlClearColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlClearColor) String

```go
func (c *GlClearColor) String() string
```

#### func (*GlClearColor) TypeID

```go
func (c *GlClearColor) TypeID() atom.TypeID
```

#### type GlClearDepthf

```go
type GlClearDepthf struct {
	binary.Generate
	Depth float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlClearDepthf
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlClearDepthf

```go
func NewGlClearDepthf(
	pDepth float32,
) *GlClearDepthf
```

#### func (*GlClearDepthf) API

```go
func (c *GlClearDepthf) API() gfxapi.API
```

#### func (*GlClearDepthf) Class

```go
func (*GlClearDepthf) Class() binary.Class
```

#### func (*GlClearDepthf) Flags

```go
func (c *GlClearDepthf) Flags() atom.Flags
```

#### func (*GlClearDepthf) Mutate

```go
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlClearDepthf) Replay

```go
func (ϟa *GlClearDepthf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlClearDepthf) String

```go
func (c *GlClearDepthf) String() string
```

#### func (*GlClearDepthf) TypeID

```go
func (c *GlClearDepthf) TypeID() atom.TypeID
```

#### type GlClearStencil

```go
type GlClearStencil struct {
	binary.Generate
	Stencil int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlClearStencil
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlClearStencil

```go
func NewGlClearStencil(
	pStencil int32,
) *GlClearStencil
```

#### func (*GlClearStencil) API

```go
func (c *GlClearStencil) API() gfxapi.API
```

#### func (*GlClearStencil) Class

```go
func (*GlClearStencil) Class() binary.Class
```

#### func (*GlClearStencil) Flags

```go
func (c *GlClearStencil) Flags() atom.Flags
```

#### func (*GlClearStencil) Mutate

```go
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlClearStencil) Replay

```go
func (ϟa *GlClearStencil) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlClearStencil) String

```go
func (c *GlClearStencil) String() string
```

#### func (*GlClearStencil) TypeID

```go
func (c *GlClearStencil) TypeID() atom.TypeID
```

#### type GlColorMask

```go
type GlColorMask struct {
	binary.Generate
	Red   bool
	Green bool
	Blue  bool
	Alpha bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlColorMask
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlColorMask

```go
func NewGlColorMask(
	pRed bool,
	pGreen bool,
	pBlue bool,
	pAlpha bool,
) *GlColorMask
```

#### func (*GlColorMask) API

```go
func (c *GlColorMask) API() gfxapi.API
```

#### func (*GlColorMask) Class

```go
func (*GlColorMask) Class() binary.Class
```

#### func (*GlColorMask) Flags

```go
func (c *GlColorMask) Flags() atom.Flags
```

#### func (*GlColorMask) Mutate

```go
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlColorMask) Replay

```go
func (ϟa *GlColorMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlColorMask) String

```go
func (c *GlColorMask) String() string
```

#### func (*GlColorMask) TypeID

```go
func (c *GlColorMask) TypeID() atom.TypeID
```

#### type GlCompileShader

```go
type GlCompileShader struct {
	binary.Generate
	Shader ShaderId
}
```

//////////////////////////////////////////////////////////////////////////////
GlCompileShader
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCompileShader

```go
func NewGlCompileShader(
	pShader ShaderId,
) *GlCompileShader
```

#### func (*GlCompileShader) API

```go
func (c *GlCompileShader) API() gfxapi.API
```

#### func (*GlCompileShader) Class

```go
func (*GlCompileShader) Class() binary.Class
```

#### func (*GlCompileShader) Flags

```go
func (c *GlCompileShader) Flags() atom.Flags
```

#### func (*GlCompileShader) Mutate

```go
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCompileShader) Replay

```go
func (ϟa *GlCompileShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCompileShader) String

```go
func (c *GlCompileShader) String() string
```

#### func (*GlCompileShader) TypeID

```go
func (c *GlCompileShader) TypeID() atom.TypeID
```

#### type GlCompressedTexImage2D

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlCompressedTexImage2D
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCompressedTexImage2D

```go
func NewGlCompressedTexImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pFormat CompressedTexelFormat,
	pWidth int32,
	pHeight int32,
	pBorder int32,
	pImageSize int32,
	pData TexturePointer,
) *GlCompressedTexImage2D
```

#### func (*GlCompressedTexImage2D) API

```go
func (c *GlCompressedTexImage2D) API() gfxapi.API
```

#### func (*GlCompressedTexImage2D) Class

```go
func (*GlCompressedTexImage2D) Class() binary.Class
```

#### func (*GlCompressedTexImage2D) Flags

```go
func (c *GlCompressedTexImage2D) Flags() atom.Flags
```

#### func (*GlCompressedTexImage2D) Mutate

```go
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCompressedTexImage2D) Replay

```go
func (ϟa *GlCompressedTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCompressedTexImage2D) String

```go
func (c *GlCompressedTexImage2D) String() string
```

#### func (*GlCompressedTexImage2D) TypeID

```go
func (c *GlCompressedTexImage2D) TypeID() atom.TypeID
```

#### type GlCompressedTexSubImage2D

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlCompressedTexSubImage2D
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCompressedTexSubImage2D

```go
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
) *GlCompressedTexSubImage2D
```

#### func (*GlCompressedTexSubImage2D) API

```go
func (c *GlCompressedTexSubImage2D) API() gfxapi.API
```

#### func (*GlCompressedTexSubImage2D) Class

```go
func (*GlCompressedTexSubImage2D) Class() binary.Class
```

#### func (*GlCompressedTexSubImage2D) Flags

```go
func (c *GlCompressedTexSubImage2D) Flags() atom.Flags
```

#### func (*GlCompressedTexSubImage2D) Mutate

```go
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCompressedTexSubImage2D) Replay

```go
func (ϟa *GlCompressedTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCompressedTexSubImage2D) String

```go
func (c *GlCompressedTexSubImage2D) String() string
```

#### func (*GlCompressedTexSubImage2D) TypeID

```go
func (c *GlCompressedTexSubImage2D) TypeID() atom.TypeID
```

#### type GlCopyTexImage2D

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlCopyTexImage2D
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCopyTexImage2D

```go
func NewGlCopyTexImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pFormat TexelFormat,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pBorder int32,
) *GlCopyTexImage2D
```

#### func (*GlCopyTexImage2D) API

```go
func (c *GlCopyTexImage2D) API() gfxapi.API
```

#### func (*GlCopyTexImage2D) Class

```go
func (*GlCopyTexImage2D) Class() binary.Class
```

#### func (*GlCopyTexImage2D) Flags

```go
func (c *GlCopyTexImage2D) Flags() atom.Flags
```

#### func (*GlCopyTexImage2D) Mutate

```go
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCopyTexImage2D) Replay

```go
func (ϟa *GlCopyTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCopyTexImage2D) String

```go
func (c *GlCopyTexImage2D) String() string
```

#### func (*GlCopyTexImage2D) TypeID

```go
func (c *GlCopyTexImage2D) TypeID() atom.TypeID
```

#### type GlCopyTexSubImage2D

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlCopyTexSubImage2D
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCopyTexSubImage2D

```go
func NewGlCopyTexSubImage2D(
	pTarget TextureImageTarget,
	pLevel int32,
	pXoffset int32,
	pYoffset int32,
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlCopyTexSubImage2D
```

#### func (*GlCopyTexSubImage2D) API

```go
func (c *GlCopyTexSubImage2D) API() gfxapi.API
```

#### func (*GlCopyTexSubImage2D) Class

```go
func (*GlCopyTexSubImage2D) Class() binary.Class
```

#### func (*GlCopyTexSubImage2D) Flags

```go
func (c *GlCopyTexSubImage2D) Flags() atom.Flags
```

#### func (*GlCopyTexSubImage2D) Mutate

```go
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCopyTexSubImage2D) Replay

```go
func (ϟa *GlCopyTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCopyTexSubImage2D) String

```go
func (c *GlCopyTexSubImage2D) String() string
```

#### func (*GlCopyTexSubImage2D) TypeID

```go
func (c *GlCopyTexSubImage2D) TypeID() atom.TypeID
```

#### type GlCreateProgram

```go
type GlCreateProgram struct {
	binary.Generate
	Result ProgramId
}
```

//////////////////////////////////////////////////////////////////////////////
GlCreateProgram
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCreateProgram

```go
func NewGlCreateProgram(
	pResult ProgramId,
) *GlCreateProgram
```

#### func (*GlCreateProgram) API

```go
func (c *GlCreateProgram) API() gfxapi.API
```

#### func (*GlCreateProgram) Class

```go
func (*GlCreateProgram) Class() binary.Class
```

#### func (*GlCreateProgram) Flags

```go
func (c *GlCreateProgram) Flags() atom.Flags
```

#### func (*GlCreateProgram) Mutate

```go
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCreateProgram) Replay

```go
func (ϟa *GlCreateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCreateProgram) String

```go
func (c *GlCreateProgram) String() string
```

#### func (*GlCreateProgram) TypeID

```go
func (c *GlCreateProgram) TypeID() atom.TypeID
```

#### type GlCreateProgram_Postback

```go
type GlCreateProgram_Postback struct {
	Result ProgramId
}
```


#### func (*GlCreateProgram_Postback) Decode

```go
func (o *GlCreateProgram_Postback) Decode(d binary.Decoder) error
```

#### type GlCreateShader

```go
type GlCreateShader struct {
	binary.Generate
	Type   ShaderType
	Result ShaderId
}
```

//////////////////////////////////////////////////////////////////////////////
GlCreateShader
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCreateShader

```go
func NewGlCreateShader(
	pType ShaderType,
	pResult ShaderId,
) *GlCreateShader
```

#### func (*GlCreateShader) API

```go
func (c *GlCreateShader) API() gfxapi.API
```

#### func (*GlCreateShader) Class

```go
func (*GlCreateShader) Class() binary.Class
```

#### func (*GlCreateShader) Flags

```go
func (c *GlCreateShader) Flags() atom.Flags
```

#### func (*GlCreateShader) Mutate

```go
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCreateShader) Replay

```go
func (ϟa *GlCreateShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCreateShader) String

```go
func (c *GlCreateShader) String() string
```

#### func (*GlCreateShader) TypeID

```go
func (c *GlCreateShader) TypeID() atom.TypeID
```

#### type GlCreateShader_Postback

```go
type GlCreateShader_Postback struct {
	Result ShaderId
}
```


#### func (*GlCreateShader_Postback) Decode

```go
func (o *GlCreateShader_Postback) Decode(d binary.Decoder) error
```

#### type GlCullFace

```go
type GlCullFace struct {
	binary.Generate
	Mode FaceMode
}
```

//////////////////////////////////////////////////////////////////////////////
GlCullFace
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlCullFace

```go
func NewGlCullFace(
	pMode FaceMode,
) *GlCullFace
```

#### func (*GlCullFace) API

```go
func (c *GlCullFace) API() gfxapi.API
```

#### func (*GlCullFace) Class

```go
func (*GlCullFace) Class() binary.Class
```

#### func (*GlCullFace) Flags

```go
func (c *GlCullFace) Flags() atom.Flags
```

#### func (*GlCullFace) Mutate

```go
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlCullFace) Replay

```go
func (ϟa *GlCullFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlCullFace) String

```go
func (c *GlCullFace) String() string
```

#### func (*GlCullFace) TypeID

```go
func (c *GlCullFace) TypeID() atom.TypeID
```

#### type GlDeleteBuffers

```go
type GlDeleteBuffers struct {
	binary.Generate
	Count   int32
	Buffers BufferIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteBuffers

```go
func NewGlDeleteBuffers(
	pCount int32,
	pBuffers BufferIdArray,
) *GlDeleteBuffers
```

#### func (*GlDeleteBuffers) API

```go
func (c *GlDeleteBuffers) API() gfxapi.API
```

#### func (*GlDeleteBuffers) Class

```go
func (*GlDeleteBuffers) Class() binary.Class
```

#### func (*GlDeleteBuffers) Flags

```go
func (c *GlDeleteBuffers) Flags() atom.Flags
```

#### func (*GlDeleteBuffers) Mutate

```go
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteBuffers) Replay

```go
func (ϟa *GlDeleteBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteBuffers) String

```go
func (c *GlDeleteBuffers) String() string
```

#### func (*GlDeleteBuffers) TypeID

```go
func (c *GlDeleteBuffers) TypeID() atom.TypeID
```

#### type GlDeleteFramebuffers

```go
type GlDeleteFramebuffers struct {
	binary.Generate
	Count        int32
	Framebuffers FramebufferIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteFramebuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteFramebuffers

```go
func NewGlDeleteFramebuffers(
	pCount int32,
	pFramebuffers FramebufferIdArray,
) *GlDeleteFramebuffers
```

#### func (*GlDeleteFramebuffers) API

```go
func (c *GlDeleteFramebuffers) API() gfxapi.API
```

#### func (*GlDeleteFramebuffers) Class

```go
func (*GlDeleteFramebuffers) Class() binary.Class
```

#### func (*GlDeleteFramebuffers) Flags

```go
func (c *GlDeleteFramebuffers) Flags() atom.Flags
```

#### func (*GlDeleteFramebuffers) Mutate

```go
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteFramebuffers) Replay

```go
func (ϟa *GlDeleteFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteFramebuffers) String

```go
func (c *GlDeleteFramebuffers) String() string
```

#### func (*GlDeleteFramebuffers) TypeID

```go
func (c *GlDeleteFramebuffers) TypeID() atom.TypeID
```

#### type GlDeleteProgram

```go
type GlDeleteProgram struct {
	binary.Generate
	Program ProgramId
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteProgram
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteProgram

```go
func NewGlDeleteProgram(
	pProgram ProgramId,
) *GlDeleteProgram
```

#### func (*GlDeleteProgram) API

```go
func (c *GlDeleteProgram) API() gfxapi.API
```

#### func (*GlDeleteProgram) Class

```go
func (*GlDeleteProgram) Class() binary.Class
```

#### func (*GlDeleteProgram) Flags

```go
func (c *GlDeleteProgram) Flags() atom.Flags
```

#### func (*GlDeleteProgram) Mutate

```go
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteProgram) Replay

```go
func (ϟa *GlDeleteProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteProgram) String

```go
func (c *GlDeleteProgram) String() string
```

#### func (*GlDeleteProgram) TypeID

```go
func (c *GlDeleteProgram) TypeID() atom.TypeID
```

#### type GlDeleteQueries

```go
type GlDeleteQueries struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteQueries
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteQueries

```go
func NewGlDeleteQueries(
	pCount int32,
	pQueries QueryIdArray,
) *GlDeleteQueries
```

#### func (*GlDeleteQueries) API

```go
func (c *GlDeleteQueries) API() gfxapi.API
```

#### func (*GlDeleteQueries) Class

```go
func (*GlDeleteQueries) Class() binary.Class
```

#### func (*GlDeleteQueries) Flags

```go
func (c *GlDeleteQueries) Flags() atom.Flags
```

#### func (*GlDeleteQueries) Mutate

```go
func (ϟa *GlDeleteQueries) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteQueries) Replay

```go
func (ϟa *GlDeleteQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteQueries) String

```go
func (c *GlDeleteQueries) String() string
```

#### func (*GlDeleteQueries) TypeID

```go
func (c *GlDeleteQueries) TypeID() atom.TypeID
```

#### type GlDeleteQueriesEXT

```go
type GlDeleteQueriesEXT struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteQueriesEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteQueriesEXT

```go
func NewGlDeleteQueriesEXT(
	pCount int32,
	pQueries QueryIdArray,
) *GlDeleteQueriesEXT
```

#### func (*GlDeleteQueriesEXT) API

```go
func (c *GlDeleteQueriesEXT) API() gfxapi.API
```

#### func (*GlDeleteQueriesEXT) Class

```go
func (*GlDeleteQueriesEXT) Class() binary.Class
```

#### func (*GlDeleteQueriesEXT) Flags

```go
func (c *GlDeleteQueriesEXT) Flags() atom.Flags
```

#### func (*GlDeleteQueriesEXT) Mutate

```go
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteQueriesEXT) Replay

```go
func (ϟa *GlDeleteQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteQueriesEXT) String

```go
func (c *GlDeleteQueriesEXT) String() string
```

#### func (*GlDeleteQueriesEXT) TypeID

```go
func (c *GlDeleteQueriesEXT) TypeID() atom.TypeID
```

#### type GlDeleteRenderbuffers

```go
type GlDeleteRenderbuffers struct {
	binary.Generate
	Count         int32
	Renderbuffers RenderbufferIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteRenderbuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteRenderbuffers

```go
func NewGlDeleteRenderbuffers(
	pCount int32,
	pRenderbuffers RenderbufferIdArray,
) *GlDeleteRenderbuffers
```

#### func (*GlDeleteRenderbuffers) API

```go
func (c *GlDeleteRenderbuffers) API() gfxapi.API
```

#### func (*GlDeleteRenderbuffers) Class

```go
func (*GlDeleteRenderbuffers) Class() binary.Class
```

#### func (*GlDeleteRenderbuffers) Flags

```go
func (c *GlDeleteRenderbuffers) Flags() atom.Flags
```

#### func (*GlDeleteRenderbuffers) Mutate

```go
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteRenderbuffers) Replay

```go
func (ϟa *GlDeleteRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteRenderbuffers) String

```go
func (c *GlDeleteRenderbuffers) String() string
```

#### func (*GlDeleteRenderbuffers) TypeID

```go
func (c *GlDeleteRenderbuffers) TypeID() atom.TypeID
```

#### type GlDeleteShader

```go
type GlDeleteShader struct {
	binary.Generate
	Shader ShaderId
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteShader
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteShader

```go
func NewGlDeleteShader(
	pShader ShaderId,
) *GlDeleteShader
```

#### func (*GlDeleteShader) API

```go
func (c *GlDeleteShader) API() gfxapi.API
```

#### func (*GlDeleteShader) Class

```go
func (*GlDeleteShader) Class() binary.Class
```

#### func (*GlDeleteShader) Flags

```go
func (c *GlDeleteShader) Flags() atom.Flags
```

#### func (*GlDeleteShader) Mutate

```go
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteShader) Replay

```go
func (ϟa *GlDeleteShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteShader) String

```go
func (c *GlDeleteShader) String() string
```

#### func (*GlDeleteShader) TypeID

```go
func (c *GlDeleteShader) TypeID() atom.TypeID
```

#### type GlDeleteTextures

```go
type GlDeleteTextures struct {
	binary.Generate
	Count    int32
	Textures TextureIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteTextures
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteTextures

```go
func NewGlDeleteTextures(
	pCount int32,
	pTextures TextureIdArray,
) *GlDeleteTextures
```

#### func (*GlDeleteTextures) API

```go
func (c *GlDeleteTextures) API() gfxapi.API
```

#### func (*GlDeleteTextures) Class

```go
func (*GlDeleteTextures) Class() binary.Class
```

#### func (*GlDeleteTextures) Flags

```go
func (c *GlDeleteTextures) Flags() atom.Flags
```

#### func (*GlDeleteTextures) Mutate

```go
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteTextures) Replay

```go
func (ϟa *GlDeleteTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteTextures) String

```go
func (c *GlDeleteTextures) String() string
```

#### func (*GlDeleteTextures) TypeID

```go
func (c *GlDeleteTextures) TypeID() atom.TypeID
```

#### type GlDeleteVertexArraysOES

```go
type GlDeleteVertexArraysOES struct {
	binary.Generate
	Count  int32
	Arrays VertexArrayIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteVertexArraysOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteVertexArraysOES

```go
func NewGlDeleteVertexArraysOES(
	pCount int32,
	pArrays VertexArrayIdArray,
) *GlDeleteVertexArraysOES
```

#### func (*GlDeleteVertexArraysOES) API

```go
func (c *GlDeleteVertexArraysOES) API() gfxapi.API
```

#### func (*GlDeleteVertexArraysOES) Class

```go
func (*GlDeleteVertexArraysOES) Class() binary.Class
```

#### func (*GlDeleteVertexArraysOES) Flags

```go
func (c *GlDeleteVertexArraysOES) Flags() atom.Flags
```

#### func (*GlDeleteVertexArraysOES) Mutate

```go
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDeleteVertexArraysOES) Replay

```go
func (ϟa *GlDeleteVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDeleteVertexArraysOES) String

```go
func (c *GlDeleteVertexArraysOES) String() string
```

#### func (*GlDeleteVertexArraysOES) TypeID

```go
func (c *GlDeleteVertexArraysOES) TypeID() atom.TypeID
```

#### type GlDepthFunc

```go
type GlDepthFunc struct {
	binary.Generate
	Function TestFunction
}
```

//////////////////////////////////////////////////////////////////////////////
GlDepthFunc
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDepthFunc

```go
func NewGlDepthFunc(
	pFunction TestFunction,
) *GlDepthFunc
```

#### func (*GlDepthFunc) API

```go
func (c *GlDepthFunc) API() gfxapi.API
```

#### func (*GlDepthFunc) Class

```go
func (*GlDepthFunc) Class() binary.Class
```

#### func (*GlDepthFunc) Flags

```go
func (c *GlDepthFunc) Flags() atom.Flags
```

#### func (*GlDepthFunc) Mutate

```go
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDepthFunc) Replay

```go
func (ϟa *GlDepthFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDepthFunc) String

```go
func (c *GlDepthFunc) String() string
```

#### func (*GlDepthFunc) TypeID

```go
func (c *GlDepthFunc) TypeID() atom.TypeID
```

#### type GlDepthMask

```go
type GlDepthMask struct {
	binary.Generate
	Enabled bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlDepthMask
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDepthMask

```go
func NewGlDepthMask(
	pEnabled bool,
) *GlDepthMask
```

#### func (*GlDepthMask) API

```go
func (c *GlDepthMask) API() gfxapi.API
```

#### func (*GlDepthMask) Class

```go
func (*GlDepthMask) Class() binary.Class
```

#### func (*GlDepthMask) Flags

```go
func (c *GlDepthMask) Flags() atom.Flags
```

#### func (*GlDepthMask) Mutate

```go
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDepthMask) Replay

```go
func (ϟa *GlDepthMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDepthMask) String

```go
func (c *GlDepthMask) String() string
```

#### func (*GlDepthMask) TypeID

```go
func (c *GlDepthMask) TypeID() atom.TypeID
```

#### type GlDepthRangef

```go
type GlDepthRangef struct {
	binary.Generate
	Near float32
	Far  float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlDepthRangef
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDepthRangef

```go
func NewGlDepthRangef(
	pNear float32,
	pFar float32,
) *GlDepthRangef
```

#### func (*GlDepthRangef) API

```go
func (c *GlDepthRangef) API() gfxapi.API
```

#### func (*GlDepthRangef) Class

```go
func (*GlDepthRangef) Class() binary.Class
```

#### func (*GlDepthRangef) Flags

```go
func (c *GlDepthRangef) Flags() atom.Flags
```

#### func (*GlDepthRangef) Mutate

```go
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDepthRangef) Replay

```go
func (ϟa *GlDepthRangef) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDepthRangef) String

```go
func (c *GlDepthRangef) String() string
```

#### func (*GlDepthRangef) TypeID

```go
func (c *GlDepthRangef) TypeID() atom.TypeID
```

#### type GlDetachShader

```go
type GlDetachShader struct {
	binary.Generate
	Program ProgramId
	Shader  ShaderId
}
```

//////////////////////////////////////////////////////////////////////////////
GlDetachShader
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDetachShader

```go
func NewGlDetachShader(
	pProgram ProgramId,
	pShader ShaderId,
) *GlDetachShader
```

#### func (*GlDetachShader) API

```go
func (c *GlDetachShader) API() gfxapi.API
```

#### func (*GlDetachShader) Class

```go
func (*GlDetachShader) Class() binary.Class
```

#### func (*GlDetachShader) Flags

```go
func (c *GlDetachShader) Flags() atom.Flags
```

#### func (*GlDetachShader) Mutate

```go
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDetachShader) Replay

```go
func (ϟa *GlDetachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDetachShader) String

```go
func (c *GlDetachShader) String() string
```

#### func (*GlDetachShader) TypeID

```go
func (c *GlDetachShader) TypeID() atom.TypeID
```

#### type GlDisable

```go
type GlDisable struct {
	binary.Generate
	Capability Capability
}
```

//////////////////////////////////////////////////////////////////////////////
GlDisable
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDisable

```go
func NewGlDisable(
	pCapability Capability,
) *GlDisable
```

#### func (*GlDisable) API

```go
func (c *GlDisable) API() gfxapi.API
```

#### func (*GlDisable) Class

```go
func (*GlDisable) Class() binary.Class
```

#### func (*GlDisable) Flags

```go
func (c *GlDisable) Flags() atom.Flags
```

#### func (*GlDisable) Mutate

```go
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDisable) Replay

```go
func (ϟa *GlDisable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDisable) String

```go
func (c *GlDisable) String() string
```

#### func (*GlDisable) TypeID

```go
func (c *GlDisable) TypeID() atom.TypeID
```

#### type GlDisableClientState

```go
type GlDisableClientState struct {
	binary.Generate
	Type ArrayType
}
```

//////////////////////////////////////////////////////////////////////////////
GlDisableClientState
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDisableClientState

```go
func NewGlDisableClientState(
	pType ArrayType,
) *GlDisableClientState
```

#### func (*GlDisableClientState) API

```go
func (c *GlDisableClientState) API() gfxapi.API
```

#### func (*GlDisableClientState) Class

```go
func (*GlDisableClientState) Class() binary.Class
```

#### func (*GlDisableClientState) Flags

```go
func (c *GlDisableClientState) Flags() atom.Flags
```

#### func (*GlDisableClientState) Mutate

```go
func (ϟa *GlDisableClientState) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDisableClientState) Replay

```go
func (ϟa *GlDisableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDisableClientState) String

```go
func (c *GlDisableClientState) String() string
```

#### func (*GlDisableClientState) TypeID

```go
func (c *GlDisableClientState) TypeID() atom.TypeID
```

#### type GlDisableVertexAttribArray

```go
type GlDisableVertexAttribArray struct {
	binary.Generate
	Location AttributeLocation
}
```

//////////////////////////////////////////////////////////////////////////////
GlDisableVertexAttribArray
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDisableVertexAttribArray

```go
func NewGlDisableVertexAttribArray(
	pLocation AttributeLocation,
) *GlDisableVertexAttribArray
```

#### func (*GlDisableVertexAttribArray) API

```go
func (c *GlDisableVertexAttribArray) API() gfxapi.API
```

#### func (*GlDisableVertexAttribArray) Class

```go
func (*GlDisableVertexAttribArray) Class() binary.Class
```

#### func (*GlDisableVertexAttribArray) Flags

```go
func (c *GlDisableVertexAttribArray) Flags() atom.Flags
```

#### func (*GlDisableVertexAttribArray) Mutate

```go
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDisableVertexAttribArray) Replay

```go
func (ϟa *GlDisableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDisableVertexAttribArray) String

```go
func (c *GlDisableVertexAttribArray) String() string
```

#### func (*GlDisableVertexAttribArray) TypeID

```go
func (c *GlDisableVertexAttribArray) TypeID() atom.TypeID
```

#### type GlDiscardFramebufferEXT

```go
type GlDiscardFramebufferEXT struct {
	binary.Generate
	Target         FramebufferTarget
	NumAttachments int32
	Attachments    DiscardFramebufferAttachmentArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlDiscardFramebufferEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDiscardFramebufferEXT

```go
func NewGlDiscardFramebufferEXT(
	pTarget FramebufferTarget,
	pNumAttachments int32,
	pAttachments DiscardFramebufferAttachmentArray,
) *GlDiscardFramebufferEXT
```

#### func (*GlDiscardFramebufferEXT) API

```go
func (c *GlDiscardFramebufferEXT) API() gfxapi.API
```

#### func (*GlDiscardFramebufferEXT) Class

```go
func (*GlDiscardFramebufferEXT) Class() binary.Class
```

#### func (*GlDiscardFramebufferEXT) Flags

```go
func (c *GlDiscardFramebufferEXT) Flags() atom.Flags
```

#### func (*GlDiscardFramebufferEXT) Mutate

```go
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDiscardFramebufferEXT) Replay

```go
func (ϟa *GlDiscardFramebufferEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDiscardFramebufferEXT) String

```go
func (c *GlDiscardFramebufferEXT) String() string
```

#### func (*GlDiscardFramebufferEXT) TypeID

```go
func (c *GlDiscardFramebufferEXT) TypeID() atom.TypeID
```

#### type GlDrawArrays

```go
type GlDrawArrays struct {
	binary.Generate
	DrawMode   DrawMode
	FirstIndex int32
	IndexCount int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlDrawArrays
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDrawArrays

```go
func NewGlDrawArrays(
	pDrawMode DrawMode,
	pFirstIndex int32,
	pIndexCount int32,
) *GlDrawArrays
```

#### func (*GlDrawArrays) API

```go
func (c *GlDrawArrays) API() gfxapi.API
```

#### func (*GlDrawArrays) Class

```go
func (*GlDrawArrays) Class() binary.Class
```

#### func (*GlDrawArrays) Flags

```go
func (c *GlDrawArrays) Flags() atom.Flags
```

#### func (*GlDrawArrays) Mutate

```go
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDrawArrays) Replay

```go
func (ϟa *GlDrawArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDrawArrays) String

```go
func (c *GlDrawArrays) String() string
```

#### func (*GlDrawArrays) TypeID

```go
func (c *GlDrawArrays) TypeID() atom.TypeID
```

#### type GlDrawElements

```go
type GlDrawElements struct {
	binary.Generate
	DrawMode     DrawMode
	ElementCount int32
	IndicesType  IndicesType
	Indices      IndicesPointer
}
```

//////////////////////////////////////////////////////////////////////////////
GlDrawElements
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDrawElements

```go
func NewGlDrawElements(
	pDrawMode DrawMode,
	pElementCount int32,
	pIndicesType IndicesType,
	pIndices IndicesPointer,
) *GlDrawElements
```

#### func (*GlDrawElements) API

```go
func (c *GlDrawElements) API() gfxapi.API
```

#### func (*GlDrawElements) Class

```go
func (*GlDrawElements) Class() binary.Class
```

#### func (*GlDrawElements) Flags

```go
func (c *GlDrawElements) Flags() atom.Flags
```

#### func (*GlDrawElements) Mutate

```go
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlDrawElements) Replay

```go
func (ϟa *GlDrawElements) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlDrawElements) String

```go
func (c *GlDrawElements) String() string
```

#### func (*GlDrawElements) TypeID

```go
func (c *GlDrawElements) TypeID() atom.TypeID
```

#### type GlEGLImageTargetRenderbufferStorageOES

```go
type GlEGLImageTargetRenderbufferStorageOES struct {
	binary.Generate
	Target ImageTargetRenderbufferStorage
	Image  TexturePointer
}
```

//////////////////////////////////////////////////////////////////////////////
GlEGLImageTargetRenderbufferStorageOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEGLImageTargetRenderbufferStorageOES

```go
func NewGlEGLImageTargetRenderbufferStorageOES(
	pTarget ImageTargetRenderbufferStorage,
	pImage TexturePointer,
) *GlEGLImageTargetRenderbufferStorageOES
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) API

```go
func (c *GlEGLImageTargetRenderbufferStorageOES) API() gfxapi.API
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) Class

```go
func (*GlEGLImageTargetRenderbufferStorageOES) Class() binary.Class
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) Flags

```go
func (c *GlEGLImageTargetRenderbufferStorageOES) Flags() atom.Flags
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) Mutate

```go
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) Replay

```go
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) String

```go
func (c *GlEGLImageTargetRenderbufferStorageOES) String() string
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) TypeID

```go
func (c *GlEGLImageTargetRenderbufferStorageOES) TypeID() atom.TypeID
```

#### type GlEGLImageTargetTexture2DOES

```go
type GlEGLImageTargetTexture2DOES struct {
	binary.Generate
	Target ImageTargetTexture
	Image  ImageOES
}
```

//////////////////////////////////////////////////////////////////////////////
GlEGLImageTargetTexture2DOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEGLImageTargetTexture2DOES

```go
func NewGlEGLImageTargetTexture2DOES(
	pTarget ImageTargetTexture,
	pImage ImageOES,
) *GlEGLImageTargetTexture2DOES
```

#### func (*GlEGLImageTargetTexture2DOES) API

```go
func (c *GlEGLImageTargetTexture2DOES) API() gfxapi.API
```

#### func (*GlEGLImageTargetTexture2DOES) Class

```go
func (*GlEGLImageTargetTexture2DOES) Class() binary.Class
```

#### func (*GlEGLImageTargetTexture2DOES) Flags

```go
func (c *GlEGLImageTargetTexture2DOES) Flags() atom.Flags
```

#### func (*GlEGLImageTargetTexture2DOES) Mutate

```go
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEGLImageTargetTexture2DOES) Replay

```go
func (ϟa *GlEGLImageTargetTexture2DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEGLImageTargetTexture2DOES) String

```go
func (c *GlEGLImageTargetTexture2DOES) String() string
```

#### func (*GlEGLImageTargetTexture2DOES) TypeID

```go
func (c *GlEGLImageTargetTexture2DOES) TypeID() atom.TypeID
```

#### type GlEnable

```go
type GlEnable struct {
	binary.Generate
	Capability Capability
}
```

//////////////////////////////////////////////////////////////////////////////
GlEnable
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEnable

```go
func NewGlEnable(
	pCapability Capability,
) *GlEnable
```

#### func (*GlEnable) API

```go
func (c *GlEnable) API() gfxapi.API
```

#### func (*GlEnable) Class

```go
func (*GlEnable) Class() binary.Class
```

#### func (*GlEnable) Flags

```go
func (c *GlEnable) Flags() atom.Flags
```

#### func (*GlEnable) Mutate

```go
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEnable) Replay

```go
func (ϟa *GlEnable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEnable) String

```go
func (c *GlEnable) String() string
```

#### func (*GlEnable) TypeID

```go
func (c *GlEnable) TypeID() atom.TypeID
```

#### type GlEnableClientState

```go
type GlEnableClientState struct {
	binary.Generate
	Type ArrayType
}
```

//////////////////////////////////////////////////////////////////////////////
GlEnableClientState
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEnableClientState

```go
func NewGlEnableClientState(
	pType ArrayType,
) *GlEnableClientState
```

#### func (*GlEnableClientState) API

```go
func (c *GlEnableClientState) API() gfxapi.API
```

#### func (*GlEnableClientState) Class

```go
func (*GlEnableClientState) Class() binary.Class
```

#### func (*GlEnableClientState) Flags

```go
func (c *GlEnableClientState) Flags() atom.Flags
```

#### func (*GlEnableClientState) Mutate

```go
func (ϟa *GlEnableClientState) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEnableClientState) Replay

```go
func (ϟa *GlEnableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEnableClientState) String

```go
func (c *GlEnableClientState) String() string
```

#### func (*GlEnableClientState) TypeID

```go
func (c *GlEnableClientState) TypeID() atom.TypeID
```

#### type GlEnableVertexAttribArray

```go
type GlEnableVertexAttribArray struct {
	binary.Generate
	Location AttributeLocation
}
```

//////////////////////////////////////////////////////////////////////////////
GlEnableVertexAttribArray
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEnableVertexAttribArray

```go
func NewGlEnableVertexAttribArray(
	pLocation AttributeLocation,
) *GlEnableVertexAttribArray
```

#### func (*GlEnableVertexAttribArray) API

```go
func (c *GlEnableVertexAttribArray) API() gfxapi.API
```

#### func (*GlEnableVertexAttribArray) Class

```go
func (*GlEnableVertexAttribArray) Class() binary.Class
```

#### func (*GlEnableVertexAttribArray) Flags

```go
func (c *GlEnableVertexAttribArray) Flags() atom.Flags
```

#### func (*GlEnableVertexAttribArray) Mutate

```go
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEnableVertexAttribArray) Replay

```go
func (ϟa *GlEnableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEnableVertexAttribArray) String

```go
func (c *GlEnableVertexAttribArray) String() string
```

#### func (*GlEnableVertexAttribArray) TypeID

```go
func (c *GlEnableVertexAttribArray) TypeID() atom.TypeID
```

#### type GlEndQuery

```go
type GlEndQuery struct {
	binary.Generate
	Target QueryTarget
}
```

//////////////////////////////////////////////////////////////////////////////
GlEndQuery
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEndQuery

```go
func NewGlEndQuery(
	pTarget QueryTarget,
) *GlEndQuery
```

#### func (*GlEndQuery) API

```go
func (c *GlEndQuery) API() gfxapi.API
```

#### func (*GlEndQuery) Class

```go
func (*GlEndQuery) Class() binary.Class
```

#### func (*GlEndQuery) Flags

```go
func (c *GlEndQuery) Flags() atom.Flags
```

#### func (*GlEndQuery) Mutate

```go
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEndQuery) Replay

```go
func (ϟa *GlEndQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEndQuery) String

```go
func (c *GlEndQuery) String() string
```

#### func (*GlEndQuery) TypeID

```go
func (c *GlEndQuery) TypeID() atom.TypeID
```

#### type GlEndQueryEXT

```go
type GlEndQueryEXT struct {
	binary.Generate
	Target QueryTarget
}
```

//////////////////////////////////////////////////////////////////////////////
GlEndQueryEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEndQueryEXT

```go
func NewGlEndQueryEXT(
	pTarget QueryTarget,
) *GlEndQueryEXT
```

#### func (*GlEndQueryEXT) API

```go
func (c *GlEndQueryEXT) API() gfxapi.API
```

#### func (*GlEndQueryEXT) Class

```go
func (*GlEndQueryEXT) Class() binary.Class
```

#### func (*GlEndQueryEXT) Flags

```go
func (c *GlEndQueryEXT) Flags() atom.Flags
```

#### func (*GlEndQueryEXT) Mutate

```go
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEndQueryEXT) Replay

```go
func (ϟa *GlEndQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEndQueryEXT) String

```go
func (c *GlEndQueryEXT) String() string
```

#### func (*GlEndQueryEXT) TypeID

```go
func (c *GlEndQueryEXT) TypeID() atom.TypeID
```

#### type GlEndTilingQCOM

```go
type GlEndTilingQCOM struct {
	binary.Generate
	PreserveMask TilePreserveMaskQCOM
}
```

//////////////////////////////////////////////////////////////////////////////
GlEndTilingQCOM
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlEndTilingQCOM

```go
func NewGlEndTilingQCOM(
	pPreserveMask TilePreserveMaskQCOM,
) *GlEndTilingQCOM
```

#### func (*GlEndTilingQCOM) API

```go
func (c *GlEndTilingQCOM) API() gfxapi.API
```

#### func (*GlEndTilingQCOM) Class

```go
func (*GlEndTilingQCOM) Class() binary.Class
```

#### func (*GlEndTilingQCOM) Flags

```go
func (c *GlEndTilingQCOM) Flags() atom.Flags
```

#### func (*GlEndTilingQCOM) Mutate

```go
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlEndTilingQCOM) Replay

```go
func (ϟa *GlEndTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlEndTilingQCOM) String

```go
func (c *GlEndTilingQCOM) String() string
```

#### func (*GlEndTilingQCOM) TypeID

```go
func (c *GlEndTilingQCOM) TypeID() atom.TypeID
```

#### type GlFinish

```go
type GlFinish struct {
	binary.Generate
}
```

//////////////////////////////////////////////////////////////////////////////
GlFinish
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlFinish

```go
func NewGlFinish() *GlFinish
```

#### func (*GlFinish) API

```go
func (c *GlFinish) API() gfxapi.API
```

#### func (*GlFinish) Class

```go
func (*GlFinish) Class() binary.Class
```

#### func (*GlFinish) Flags

```go
func (c *GlFinish) Flags() atom.Flags
```

#### func (*GlFinish) Mutate

```go
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlFinish) Replay

```go
func (ϟa *GlFinish) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlFinish) String

```go
func (c *GlFinish) String() string
```

#### func (*GlFinish) TypeID

```go
func (c *GlFinish) TypeID() atom.TypeID
```

#### type GlFlush

```go
type GlFlush struct {
	binary.Generate
}
```

//////////////////////////////////////////////////////////////////////////////
GlFlush
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlFlush

```go
func NewGlFlush() *GlFlush
```

#### func (*GlFlush) API

```go
func (c *GlFlush) API() gfxapi.API
```

#### func (*GlFlush) Class

```go
func (*GlFlush) Class() binary.Class
```

#### func (*GlFlush) Flags

```go
func (c *GlFlush) Flags() atom.Flags
```

#### func (*GlFlush) Mutate

```go
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlFlush) Replay

```go
func (ϟa *GlFlush) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlFlush) String

```go
func (c *GlFlush) String() string
```

#### func (*GlFlush) TypeID

```go
func (c *GlFlush) TypeID() atom.TypeID
```

#### type GlFramebufferRenderbuffer

```go
type GlFramebufferRenderbuffer struct {
	binary.Generate
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	RenderbufferTarget    RenderbufferTarget
	Renderbuffer          RenderbufferId
}
```

//////////////////////////////////////////////////////////////////////////////
GlFramebufferRenderbuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlFramebufferRenderbuffer

```go
func NewGlFramebufferRenderbuffer(
	pFramebufferTarget FramebufferTarget,
	pFramebufferAttachment FramebufferAttachment,
	pRenderbufferTarget RenderbufferTarget,
	pRenderbuffer RenderbufferId,
) *GlFramebufferRenderbuffer
```

#### func (*GlFramebufferRenderbuffer) API

```go
func (c *GlFramebufferRenderbuffer) API() gfxapi.API
```

#### func (*GlFramebufferRenderbuffer) Class

```go
func (*GlFramebufferRenderbuffer) Class() binary.Class
```

#### func (*GlFramebufferRenderbuffer) Flags

```go
func (c *GlFramebufferRenderbuffer) Flags() atom.Flags
```

#### func (*GlFramebufferRenderbuffer) Mutate

```go
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlFramebufferRenderbuffer) Replay

```go
func (ϟa *GlFramebufferRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlFramebufferRenderbuffer) String

```go
func (c *GlFramebufferRenderbuffer) String() string
```

#### func (*GlFramebufferRenderbuffer) TypeID

```go
func (c *GlFramebufferRenderbuffer) TypeID() atom.TypeID
```

#### type GlFramebufferTexture2D

```go
type GlFramebufferTexture2D struct {
	binary.Generate
	FramebufferTarget     FramebufferTarget
	FramebufferAttachment FramebufferAttachment
	TextureTarget         TextureImageTarget
	Texture               TextureId
	Level                 int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlFramebufferTexture2D
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlFramebufferTexture2D

```go
func NewGlFramebufferTexture2D(
	pFramebufferTarget FramebufferTarget,
	pFramebufferAttachment FramebufferAttachment,
	pTextureTarget TextureImageTarget,
	pTexture TextureId,
	pLevel int32,
) *GlFramebufferTexture2D
```

#### func (*GlFramebufferTexture2D) API

```go
func (c *GlFramebufferTexture2D) API() gfxapi.API
```

#### func (*GlFramebufferTexture2D) Class

```go
func (*GlFramebufferTexture2D) Class() binary.Class
```

#### func (*GlFramebufferTexture2D) Flags

```go
func (c *GlFramebufferTexture2D) Flags() atom.Flags
```

#### func (*GlFramebufferTexture2D) Mutate

```go
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlFramebufferTexture2D) Replay

```go
func (ϟa *GlFramebufferTexture2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlFramebufferTexture2D) String

```go
func (c *GlFramebufferTexture2D) String() string
```

#### func (*GlFramebufferTexture2D) TypeID

```go
func (c *GlFramebufferTexture2D) TypeID() atom.TypeID
```

#### type GlFrontFace

```go
type GlFrontFace struct {
	binary.Generate
	Orientation FaceOrientation
}
```

//////////////////////////////////////////////////////////////////////////////
GlFrontFace
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlFrontFace

```go
func NewGlFrontFace(
	pOrientation FaceOrientation,
) *GlFrontFace
```

#### func (*GlFrontFace) API

```go
func (c *GlFrontFace) API() gfxapi.API
```

#### func (*GlFrontFace) Class

```go
func (*GlFrontFace) Class() binary.Class
```

#### func (*GlFrontFace) Flags

```go
func (c *GlFrontFace) Flags() atom.Flags
```

#### func (*GlFrontFace) Mutate

```go
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlFrontFace) Replay

```go
func (ϟa *GlFrontFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlFrontFace) String

```go
func (c *GlFrontFace) String() string
```

#### func (*GlFrontFace) TypeID

```go
func (c *GlFrontFace) TypeID() atom.TypeID
```

#### type GlGenBuffers

```go
type GlGenBuffers struct {
	binary.Generate
	Count   int32
	Buffers BufferIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenBuffers

```go
func NewGlGenBuffers(
	pCount int32,
	pBuffers BufferIdArray,
) *GlGenBuffers
```

#### func (*GlGenBuffers) API

```go
func (c *GlGenBuffers) API() gfxapi.API
```

#### func (*GlGenBuffers) Class

```go
func (*GlGenBuffers) Class() binary.Class
```

#### func (*GlGenBuffers) Flags

```go
func (c *GlGenBuffers) Flags() atom.Flags
```

#### func (*GlGenBuffers) Mutate

```go
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenBuffers) Replay

```go
func (ϟa *GlGenBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenBuffers) String

```go
func (c *GlGenBuffers) String() string
```

#### func (*GlGenBuffers) TypeID

```go
func (c *GlGenBuffers) TypeID() atom.TypeID
```

#### type GlGenBuffers_Postback

```go
type GlGenBuffers_Postback struct {
	Buffers BufferIdArray
}
```


#### func (*GlGenBuffers_Postback) Decode

```go
func (o *GlGenBuffers_Postback) Decode(buffers_cnt uint64, d binary.Decoder) error
```

#### type GlGenFramebuffers

```go
type GlGenFramebuffers struct {
	binary.Generate
	Count        int32
	Framebuffers FramebufferIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenFramebuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenFramebuffers

```go
func NewGlGenFramebuffers(
	pCount int32,
	pFramebuffers FramebufferIdArray,
) *GlGenFramebuffers
```

#### func (*GlGenFramebuffers) API

```go
func (c *GlGenFramebuffers) API() gfxapi.API
```

#### func (*GlGenFramebuffers) Class

```go
func (*GlGenFramebuffers) Class() binary.Class
```

#### func (*GlGenFramebuffers) Flags

```go
func (c *GlGenFramebuffers) Flags() atom.Flags
```

#### func (*GlGenFramebuffers) Mutate

```go
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenFramebuffers) Replay

```go
func (ϟa *GlGenFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenFramebuffers) String

```go
func (c *GlGenFramebuffers) String() string
```

#### func (*GlGenFramebuffers) TypeID

```go
func (c *GlGenFramebuffers) TypeID() atom.TypeID
```

#### type GlGenFramebuffers_Postback

```go
type GlGenFramebuffers_Postback struct {
	Framebuffers FramebufferIdArray
}
```


#### func (*GlGenFramebuffers_Postback) Decode

```go
func (o *GlGenFramebuffers_Postback) Decode(framebuffers_cnt uint64, d binary.Decoder) error
```

#### type GlGenQueries

```go
type GlGenQueries struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenQueries
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenQueries

```go
func NewGlGenQueries(
	pCount int32,
	pQueries QueryIdArray,
) *GlGenQueries
```

#### func (*GlGenQueries) API

```go
func (c *GlGenQueries) API() gfxapi.API
```

#### func (*GlGenQueries) Class

```go
func (*GlGenQueries) Class() binary.Class
```

#### func (*GlGenQueries) Flags

```go
func (c *GlGenQueries) Flags() atom.Flags
```

#### func (*GlGenQueries) Mutate

```go
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenQueries) Replay

```go
func (ϟa *GlGenQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenQueries) String

```go
func (c *GlGenQueries) String() string
```

#### func (*GlGenQueries) TypeID

```go
func (c *GlGenQueries) TypeID() atom.TypeID
```

#### type GlGenQueriesEXT

```go
type GlGenQueriesEXT struct {
	binary.Generate
	Count   int32
	Queries QueryIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenQueriesEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenQueriesEXT

```go
func NewGlGenQueriesEXT(
	pCount int32,
	pQueries QueryIdArray,
) *GlGenQueriesEXT
```

#### func (*GlGenQueriesEXT) API

```go
func (c *GlGenQueriesEXT) API() gfxapi.API
```

#### func (*GlGenQueriesEXT) Class

```go
func (*GlGenQueriesEXT) Class() binary.Class
```

#### func (*GlGenQueriesEXT) Flags

```go
func (c *GlGenQueriesEXT) Flags() atom.Flags
```

#### func (*GlGenQueriesEXT) Mutate

```go
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenQueriesEXT) Replay

```go
func (ϟa *GlGenQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenQueriesEXT) String

```go
func (c *GlGenQueriesEXT) String() string
```

#### func (*GlGenQueriesEXT) TypeID

```go
func (c *GlGenQueriesEXT) TypeID() atom.TypeID
```

#### type GlGenQueriesEXT_Postback

```go
type GlGenQueriesEXT_Postback struct {
	Queries QueryIdArray
}
```


#### func (*GlGenQueriesEXT_Postback) Decode

```go
func (o *GlGenQueriesEXT_Postback) Decode(queries_cnt uint64, d binary.Decoder) error
```

#### type GlGenQueries_Postback

```go
type GlGenQueries_Postback struct {
	Queries QueryIdArray
}
```


#### func (*GlGenQueries_Postback) Decode

```go
func (o *GlGenQueries_Postback) Decode(queries_cnt uint64, d binary.Decoder) error
```

#### type GlGenRenderbuffers

```go
type GlGenRenderbuffers struct {
	binary.Generate
	Count         int32
	Renderbuffers RenderbufferIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenRenderbuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenRenderbuffers

```go
func NewGlGenRenderbuffers(
	pCount int32,
	pRenderbuffers RenderbufferIdArray,
) *GlGenRenderbuffers
```

#### func (*GlGenRenderbuffers) API

```go
func (c *GlGenRenderbuffers) API() gfxapi.API
```

#### func (*GlGenRenderbuffers) Class

```go
func (*GlGenRenderbuffers) Class() binary.Class
```

#### func (*GlGenRenderbuffers) Flags

```go
func (c *GlGenRenderbuffers) Flags() atom.Flags
```

#### func (*GlGenRenderbuffers) Mutate

```go
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenRenderbuffers) Replay

```go
func (ϟa *GlGenRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenRenderbuffers) String

```go
func (c *GlGenRenderbuffers) String() string
```

#### func (*GlGenRenderbuffers) TypeID

```go
func (c *GlGenRenderbuffers) TypeID() atom.TypeID
```

#### type GlGenRenderbuffers_Postback

```go
type GlGenRenderbuffers_Postback struct {
	Renderbuffers RenderbufferIdArray
}
```


#### func (*GlGenRenderbuffers_Postback) Decode

```go
func (o *GlGenRenderbuffers_Postback) Decode(renderbuffers_cnt uint64, d binary.Decoder) error
```

#### type GlGenTextures

```go
type GlGenTextures struct {
	binary.Generate
	Count    int32
	Textures TextureIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenTextures
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenTextures

```go
func NewGlGenTextures(
	pCount int32,
	pTextures TextureIdArray,
) *GlGenTextures
```

#### func (*GlGenTextures) API

```go
func (c *GlGenTextures) API() gfxapi.API
```

#### func (*GlGenTextures) Class

```go
func (*GlGenTextures) Class() binary.Class
```

#### func (*GlGenTextures) Flags

```go
func (c *GlGenTextures) Flags() atom.Flags
```

#### func (*GlGenTextures) Mutate

```go
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenTextures) Replay

```go
func (ϟa *GlGenTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenTextures) String

```go
func (c *GlGenTextures) String() string
```

#### func (*GlGenTextures) TypeID

```go
func (c *GlGenTextures) TypeID() atom.TypeID
```

#### type GlGenTextures_Postback

```go
type GlGenTextures_Postback struct {
	Textures TextureIdArray
}
```


#### func (*GlGenTextures_Postback) Decode

```go
func (o *GlGenTextures_Postback) Decode(textures_cnt uint64, d binary.Decoder) error
```

#### type GlGenVertexArraysOES

```go
type GlGenVertexArraysOES struct {
	binary.Generate
	Count  int32
	Arrays VertexArrayIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenVertexArraysOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenVertexArraysOES

```go
func NewGlGenVertexArraysOES(
	pCount int32,
	pArrays VertexArrayIdArray,
) *GlGenVertexArraysOES
```

#### func (*GlGenVertexArraysOES) API

```go
func (c *GlGenVertexArraysOES) API() gfxapi.API
```

#### func (*GlGenVertexArraysOES) Class

```go
func (*GlGenVertexArraysOES) Class() binary.Class
```

#### func (*GlGenVertexArraysOES) Flags

```go
func (c *GlGenVertexArraysOES) Flags() atom.Flags
```

#### func (*GlGenVertexArraysOES) Mutate

```go
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenVertexArraysOES) Replay

```go
func (ϟa *GlGenVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenVertexArraysOES) String

```go
func (c *GlGenVertexArraysOES) String() string
```

#### func (*GlGenVertexArraysOES) TypeID

```go
func (c *GlGenVertexArraysOES) TypeID() atom.TypeID
```

#### type GlGenVertexArraysOES_Postback

```go
type GlGenVertexArraysOES_Postback struct {
	Arrays VertexArrayIdArray
}
```


#### func (*GlGenVertexArraysOES_Postback) Decode

```go
func (o *GlGenVertexArraysOES_Postback) Decode(arrays_cnt uint64, d binary.Decoder) error
```

#### type GlGenerateMipmap

```go
type GlGenerateMipmap struct {
	binary.Generate
	Target TextureImageTarget
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenerateMipmap
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenerateMipmap

```go
func NewGlGenerateMipmap(
	pTarget TextureImageTarget,
) *GlGenerateMipmap
```

#### func (*GlGenerateMipmap) API

```go
func (c *GlGenerateMipmap) API() gfxapi.API
```

#### func (*GlGenerateMipmap) Class

```go
func (*GlGenerateMipmap) Class() binary.Class
```

#### func (*GlGenerateMipmap) Flags

```go
func (c *GlGenerateMipmap) Flags() atom.Flags
```

#### func (*GlGenerateMipmap) Mutate

```go
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGenerateMipmap) Replay

```go
func (ϟa *GlGenerateMipmap) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGenerateMipmap) String

```go
func (c *GlGenerateMipmap) String() string
```

#### func (*GlGenerateMipmap) TypeID

```go
func (c *GlGenerateMipmap) TypeID() atom.TypeID
```

#### type GlGetActiveAttrib

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlGetActiveAttrib
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetActiveAttrib

```go
func NewGlGetActiveAttrib(
	pProgram ProgramId,
	pLocation AttributeLocation,
	pBufferSize int32,
	pBufferBytesWritten int32,
	pVectorCount int32,
	pType ShaderAttribType,
	pName string,
) *GlGetActiveAttrib
```

#### func (*GlGetActiveAttrib) API

```go
func (c *GlGetActiveAttrib) API() gfxapi.API
```

#### func (*GlGetActiveAttrib) Class

```go
func (*GlGetActiveAttrib) Class() binary.Class
```

#### func (*GlGetActiveAttrib) Flags

```go
func (c *GlGetActiveAttrib) Flags() atom.Flags
```

#### func (*GlGetActiveAttrib) Mutate

```go
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetActiveAttrib) Replay

```go
func (ϟa *GlGetActiveAttrib) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetActiveAttrib) String

```go
func (c *GlGetActiveAttrib) String() string
```

#### func (*GlGetActiveAttrib) TypeID

```go
func (c *GlGetActiveAttrib) TypeID() atom.TypeID
```

#### type GlGetActiveAttrib_Postback

```go
type GlGetActiveAttrib_Postback struct {
	BufferBytesWritten int32
	VectorCount        int32
	Type               ShaderAttribType
	Name               string
}
```


#### func (*GlGetActiveAttrib_Postback) Decode

```go
func (o *GlGetActiveAttrib_Postback) Decode(name_cnt uint64, d binary.Decoder) error
```

#### type GlGetActiveUniform

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlGetActiveUniform
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetActiveUniform

```go
func NewGlGetActiveUniform(
	pProgram ProgramId,
	pLocation int32,
	pBufferSize int32,
	pBufferBytesWritten int32,
	pSize int32,
	pType ShaderUniformType,
	pName string,
) *GlGetActiveUniform
```

#### func (*GlGetActiveUniform) API

```go
func (c *GlGetActiveUniform) API() gfxapi.API
```

#### func (*GlGetActiveUniform) Class

```go
func (*GlGetActiveUniform) Class() binary.Class
```

#### func (*GlGetActiveUniform) Flags

```go
func (c *GlGetActiveUniform) Flags() atom.Flags
```

#### func (*GlGetActiveUniform) Mutate

```go
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetActiveUniform) Replay

```go
func (ϟa *GlGetActiveUniform) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetActiveUniform) String

```go
func (c *GlGetActiveUniform) String() string
```

#### func (*GlGetActiveUniform) TypeID

```go
func (c *GlGetActiveUniform) TypeID() atom.TypeID
```

#### type GlGetActiveUniform_Postback

```go
type GlGetActiveUniform_Postback struct {
	BufferBytesWritten int32
	Size               int32
	Type               ShaderUniformType
	Name               string
}
```


#### func (*GlGetActiveUniform_Postback) Decode

```go
func (o *GlGetActiveUniform_Postback) Decode(name_cnt uint64, d binary.Decoder) error
```

#### type GlGetAttachedShaders

```go
type GlGetAttachedShaders struct {
	binary.Generate
	Program              ProgramId
	BufferLength         int32
	ShadersLengthWritten int32
	Shaders              ShaderIdArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetAttachedShaders
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetAttachedShaders

```go
func NewGlGetAttachedShaders(
	pProgram ProgramId,
	pBufferLength int32,
	pShadersLengthWritten int32,
	pShaders ShaderIdArray,
) *GlGetAttachedShaders
```

#### func (*GlGetAttachedShaders) API

```go
func (c *GlGetAttachedShaders) API() gfxapi.API
```

#### func (*GlGetAttachedShaders) Class

```go
func (*GlGetAttachedShaders) Class() binary.Class
```

#### func (*GlGetAttachedShaders) Flags

```go
func (c *GlGetAttachedShaders) Flags() atom.Flags
```

#### func (*GlGetAttachedShaders) Mutate

```go
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetAttachedShaders) Replay

```go
func (ϟa *GlGetAttachedShaders) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetAttachedShaders) String

```go
func (c *GlGetAttachedShaders) String() string
```

#### func (*GlGetAttachedShaders) TypeID

```go
func (c *GlGetAttachedShaders) TypeID() atom.TypeID
```

#### type GlGetAttachedShaders_Postback

```go
type GlGetAttachedShaders_Postback struct {
	ShadersLengthWritten int32
	Shaders              ShaderIdArray
}
```


#### func (*GlGetAttachedShaders_Postback) Decode

```go
func (o *GlGetAttachedShaders_Postback) Decode(shaders_cnt uint64, d binary.Decoder) error
```

#### type GlGetAttribLocation

```go
type GlGetAttribLocation struct {
	binary.Generate
	Program ProgramId
	Name    string
	Result  AttributeLocation
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetAttribLocation
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetAttribLocation

```go
func NewGlGetAttribLocation(
	pProgram ProgramId,
	pName string,
	pResult AttributeLocation,
) *GlGetAttribLocation
```

#### func (*GlGetAttribLocation) API

```go
func (c *GlGetAttribLocation) API() gfxapi.API
```

#### func (*GlGetAttribLocation) Class

```go
func (*GlGetAttribLocation) Class() binary.Class
```

#### func (*GlGetAttribLocation) Flags

```go
func (c *GlGetAttribLocation) Flags() atom.Flags
```

#### func (*GlGetAttribLocation) Mutate

```go
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetAttribLocation) Replay

```go
func (ω *GlGetAttribLocation) Replay(id atom.ID, s *gfxapi.State, b *builder.Builder, wantOutput bool)
```
AttributeLocations cannot be remapped like UniformLocations as the
VertexAttributeArrays are shared between different programs. Instead, simply
force the location to match what was recorded in the capture using
glBindAttribLocation. TODO: This implementation currently calls glLinkProgram
for every call to glGetAttribLocation!

    This is obviously not ideal, and we should be doing this once at glLinkProgram once the
    spy emits location hinting information.

#### func (*GlGetAttribLocation) String

```go
func (c *GlGetAttribLocation) String() string
```

#### func (*GlGetAttribLocation) TypeID

```go
func (c *GlGetAttribLocation) TypeID() atom.TypeID
```

#### type GlGetAttribLocation_Postback

```go
type GlGetAttribLocation_Postback struct {
	Result AttributeLocation
}
```


#### func (*GlGetAttribLocation_Postback) Decode

```go
func (o *GlGetAttribLocation_Postback) Decode(d binary.Decoder) error
```

#### type GlGetBooleanv

```go
type GlGetBooleanv struct {
	binary.Generate
	Param  StateVariable
	Values BoolArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetBooleanv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetBooleanv

```go
func NewGlGetBooleanv(
	pParam StateVariable,
	pValues BoolArray,
) *GlGetBooleanv
```

#### func (*GlGetBooleanv) API

```go
func (c *GlGetBooleanv) API() gfxapi.API
```

#### func (*GlGetBooleanv) Class

```go
func (*GlGetBooleanv) Class() binary.Class
```

#### func (*GlGetBooleanv) Flags

```go
func (c *GlGetBooleanv) Flags() atom.Flags
```

#### func (*GlGetBooleanv) Mutate

```go
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetBooleanv) Replay

```go
func (ϟa *GlGetBooleanv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetBooleanv) String

```go
func (c *GlGetBooleanv) String() string
```

#### func (*GlGetBooleanv) TypeID

```go
func (c *GlGetBooleanv) TypeID() atom.TypeID
```

#### type GlGetBooleanv_Postback

```go
type GlGetBooleanv_Postback struct {
	Values BoolArray
}
```


#### func (*GlGetBooleanv_Postback) Decode

```go
func (o *GlGetBooleanv_Postback) Decode(values_cnt uint64, d binary.Decoder) error
```

#### type GlGetBufferParameteriv

```go
type GlGetBufferParameteriv struct {
	binary.Generate
	Target    BufferTarget
	Parameter BufferParameter
	Value     int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetBufferParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetBufferParameteriv

```go
func NewGlGetBufferParameteriv(
	pTarget BufferTarget,
	pParameter BufferParameter,
	pValue int32,
) *GlGetBufferParameteriv
```

#### func (*GlGetBufferParameteriv) API

```go
func (c *GlGetBufferParameteriv) API() gfxapi.API
```

#### func (*GlGetBufferParameteriv) Class

```go
func (*GlGetBufferParameteriv) Class() binary.Class
```

#### func (*GlGetBufferParameteriv) Flags

```go
func (c *GlGetBufferParameteriv) Flags() atom.Flags
```

#### func (*GlGetBufferParameteriv) Mutate

```go
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetBufferParameteriv) Replay

```go
func (ϟa *GlGetBufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetBufferParameteriv) String

```go
func (c *GlGetBufferParameteriv) String() string
```

#### func (*GlGetBufferParameteriv) TypeID

```go
func (c *GlGetBufferParameteriv) TypeID() atom.TypeID
```

#### type GlGetBufferParameteriv_Postback

```go
type GlGetBufferParameteriv_Postback struct {
	Value int32
}
```


#### func (*GlGetBufferParameteriv_Postback) Decode

```go
func (o *GlGetBufferParameteriv_Postback) Decode(d binary.Decoder) error
```

#### type GlGetError

```go
type GlGetError struct {
	binary.Generate
	Result Error
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetError
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetError

```go
func NewGlGetError(
	pResult Error,
) *GlGetError
```

#### func (*GlGetError) API

```go
func (c *GlGetError) API() gfxapi.API
```

#### func (*GlGetError) Class

```go
func (*GlGetError) Class() binary.Class
```

#### func (*GlGetError) Flags

```go
func (c *GlGetError) Flags() atom.Flags
```

#### func (*GlGetError) Mutate

```go
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetError) Replay

```go
func (ϟa *GlGetError) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetError) String

```go
func (c *GlGetError) String() string
```

#### func (*GlGetError) TypeID

```go
func (c *GlGetError) TypeID() atom.TypeID
```

#### type GlGetError_Postback

```go
type GlGetError_Postback struct {
	Result Error
}
```


#### func (*GlGetError_Postback) Decode

```go
func (o *GlGetError_Postback) Decode(d binary.Decoder) error
```

#### type GlGetFloatv

```go
type GlGetFloatv struct {
	binary.Generate
	Param  StateVariable
	Values F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetFloatv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetFloatv

```go
func NewGlGetFloatv(
	pParam StateVariable,
	pValues F32Array,
) *GlGetFloatv
```

#### func (*GlGetFloatv) API

```go
func (c *GlGetFloatv) API() gfxapi.API
```

#### func (*GlGetFloatv) Class

```go
func (*GlGetFloatv) Class() binary.Class
```

#### func (*GlGetFloatv) Flags

```go
func (c *GlGetFloatv) Flags() atom.Flags
```

#### func (*GlGetFloatv) Mutate

```go
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetFloatv) Replay

```go
func (ϟa *GlGetFloatv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetFloatv) String

```go
func (c *GlGetFloatv) String() string
```

#### func (*GlGetFloatv) TypeID

```go
func (c *GlGetFloatv) TypeID() atom.TypeID
```

#### type GlGetFloatv_Postback

```go
type GlGetFloatv_Postback struct {
	Values F32Array
}
```


#### func (*GlGetFloatv_Postback) Decode

```go
func (o *GlGetFloatv_Postback) Decode(values_cnt uint64, d binary.Decoder) error
```

#### type GlGetFramebufferAttachmentParameteriv

```go
type GlGetFramebufferAttachmentParameteriv struct {
	binary.Generate
	FramebufferTarget FramebufferTarget
	Attachment        FramebufferAttachment
	Parameter         FramebufferAttachmentParameter
	Value             S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetFramebufferAttachmentParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetFramebufferAttachmentParameteriv

```go
func NewGlGetFramebufferAttachmentParameteriv(
	pFramebufferTarget FramebufferTarget,
	pAttachment FramebufferAttachment,
	pParameter FramebufferAttachmentParameter,
	pValue S32Array,
) *GlGetFramebufferAttachmentParameteriv
```

#### func (*GlGetFramebufferAttachmentParameteriv) API

```go
func (c *GlGetFramebufferAttachmentParameteriv) API() gfxapi.API
```

#### func (*GlGetFramebufferAttachmentParameteriv) Class

```go
func (*GlGetFramebufferAttachmentParameteriv) Class() binary.Class
```

#### func (*GlGetFramebufferAttachmentParameteriv) Flags

```go
func (c *GlGetFramebufferAttachmentParameteriv) Flags() atom.Flags
```

#### func (*GlGetFramebufferAttachmentParameteriv) Mutate

```go
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetFramebufferAttachmentParameteriv) Replay

```go
func (ϟa *GlGetFramebufferAttachmentParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetFramebufferAttachmentParameteriv) String

```go
func (c *GlGetFramebufferAttachmentParameteriv) String() string
```

#### func (*GlGetFramebufferAttachmentParameteriv) TypeID

```go
func (c *GlGetFramebufferAttachmentParameteriv) TypeID() atom.TypeID
```

#### type GlGetFramebufferAttachmentParameteriv_Postback

```go
type GlGetFramebufferAttachmentParameteriv_Postback struct {
	Value S32Array
}
```


#### func (*GlGetFramebufferAttachmentParameteriv_Postback) Decode

```go
func (o *GlGetFramebufferAttachmentParameteriv_Postback) Decode(value_cnt uint64, d binary.Decoder) error
```

#### type GlGetGraphicsResetStatusEXT

```go
type GlGetGraphicsResetStatusEXT struct {
	binary.Generate
	Result ResetStatus
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetGraphicsResetStatusEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetGraphicsResetStatusEXT

```go
func NewGlGetGraphicsResetStatusEXT(
	pResult ResetStatus,
) *GlGetGraphicsResetStatusEXT
```

#### func (*GlGetGraphicsResetStatusEXT) API

```go
func (c *GlGetGraphicsResetStatusEXT) API() gfxapi.API
```

#### func (*GlGetGraphicsResetStatusEXT) Class

```go
func (*GlGetGraphicsResetStatusEXT) Class() binary.Class
```

#### func (*GlGetGraphicsResetStatusEXT) Flags

```go
func (c *GlGetGraphicsResetStatusEXT) Flags() atom.Flags
```

#### func (*GlGetGraphicsResetStatusEXT) Mutate

```go
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetGraphicsResetStatusEXT) Replay

```go
func (ϟa *GlGetGraphicsResetStatusEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetGraphicsResetStatusEXT) String

```go
func (c *GlGetGraphicsResetStatusEXT) String() string
```

#### func (*GlGetGraphicsResetStatusEXT) TypeID

```go
func (c *GlGetGraphicsResetStatusEXT) TypeID() atom.TypeID
```

#### type GlGetGraphicsResetStatusEXT_Postback

```go
type GlGetGraphicsResetStatusEXT_Postback struct {
	Result ResetStatus
}
```


#### func (*GlGetGraphicsResetStatusEXT_Postback) Decode

```go
func (o *GlGetGraphicsResetStatusEXT_Postback) Decode(d binary.Decoder) error
```

#### type GlGetIntegerv

```go
type GlGetIntegerv struct {
	binary.Generate
	Param  StateVariable
	Values S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetIntegerv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetIntegerv

```go
func NewGlGetIntegerv(
	pParam StateVariable,
	pValues S32Array,
) *GlGetIntegerv
```

#### func (*GlGetIntegerv) API

```go
func (c *GlGetIntegerv) API() gfxapi.API
```

#### func (*GlGetIntegerv) Class

```go
func (*GlGetIntegerv) Class() binary.Class
```

#### func (*GlGetIntegerv) Flags

```go
func (c *GlGetIntegerv) Flags() atom.Flags
```

#### func (*GlGetIntegerv) Mutate

```go
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetIntegerv) Replay

```go
func (ϟa *GlGetIntegerv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetIntegerv) String

```go
func (c *GlGetIntegerv) String() string
```

#### func (*GlGetIntegerv) TypeID

```go
func (c *GlGetIntegerv) TypeID() atom.TypeID
```

#### type GlGetIntegerv_Postback

```go
type GlGetIntegerv_Postback struct {
	Values S32Array
}
```


#### func (*GlGetIntegerv_Postback) Decode

```go
func (o *GlGetIntegerv_Postback) Decode(values_cnt uint64, d binary.Decoder) error
```

#### type GlGetProgramBinaryOES

```go
type GlGetProgramBinaryOES struct {
	binary.Generate
	Program      ProgramId
	BufferSize   int32
	BytesWritten int32
	BinaryFormat uint32
	Binary       memory.Pointer
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetProgramBinaryOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetProgramBinaryOES

```go
func NewGlGetProgramBinaryOES(
	pProgram ProgramId,
	pBufferSize int32,
	pBytesWritten int32,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
) *GlGetProgramBinaryOES
```

#### func (*GlGetProgramBinaryOES) API

```go
func (c *GlGetProgramBinaryOES) API() gfxapi.API
```

#### func (*GlGetProgramBinaryOES) Class

```go
func (*GlGetProgramBinaryOES) Class() binary.Class
```

#### func (*GlGetProgramBinaryOES) Flags

```go
func (c *GlGetProgramBinaryOES) Flags() atom.Flags
```

#### func (*GlGetProgramBinaryOES) Mutate

```go
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetProgramBinaryOES) Replay

```go
func (ϟa *GlGetProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetProgramBinaryOES) String

```go
func (c *GlGetProgramBinaryOES) String() string
```

#### func (*GlGetProgramBinaryOES) TypeID

```go
func (c *GlGetProgramBinaryOES) TypeID() atom.TypeID
```

#### type GlGetProgramBinaryOES_Postback

```go
type GlGetProgramBinaryOES_Postback struct {
	BytesWritten int32
	BinaryFormat uint32
	Binary       []byte
}
```


#### func (*GlGetProgramBinaryOES_Postback) Decode

```go
func (o *GlGetProgramBinaryOES_Postback) Decode(binary_cnt uint64, d binary.Decoder) error
```

#### type GlGetProgramInfoLog

```go
type GlGetProgramInfoLog struct {
	binary.Generate
	Program             ProgramId
	BufferLength        int32
	StringLengthWritten int32
	Info                string
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetProgramInfoLog
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetProgramInfoLog

```go
func NewGlGetProgramInfoLog(
	pProgram ProgramId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pInfo string,
) *GlGetProgramInfoLog
```

#### func (*GlGetProgramInfoLog) API

```go
func (c *GlGetProgramInfoLog) API() gfxapi.API
```

#### func (*GlGetProgramInfoLog) Class

```go
func (*GlGetProgramInfoLog) Class() binary.Class
```

#### func (*GlGetProgramInfoLog) Flags

```go
func (c *GlGetProgramInfoLog) Flags() atom.Flags
```

#### func (*GlGetProgramInfoLog) Mutate

```go
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetProgramInfoLog) Replay

```go
func (ϟa *GlGetProgramInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetProgramInfoLog) String

```go
func (c *GlGetProgramInfoLog) String() string
```

#### func (*GlGetProgramInfoLog) TypeID

```go
func (c *GlGetProgramInfoLog) TypeID() atom.TypeID
```

#### type GlGetProgramInfoLog_Postback

```go
type GlGetProgramInfoLog_Postback struct {
	StringLengthWritten int32
	Info                string
}
```


#### func (*GlGetProgramInfoLog_Postback) Decode

```go
func (o *GlGetProgramInfoLog_Postback) Decode(info_cnt uint64, d binary.Decoder) error
```

#### type GlGetProgramiv

```go
type GlGetProgramiv struct {
	binary.Generate
	Program   ProgramId
	Parameter ProgramParameter
	Value     S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetProgramiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetProgramiv

```go
func NewGlGetProgramiv(
	pProgram ProgramId,
	pParameter ProgramParameter,
	pValue S32Array,
) *GlGetProgramiv
```

#### func (*GlGetProgramiv) API

```go
func (c *GlGetProgramiv) API() gfxapi.API
```

#### func (*GlGetProgramiv) Class

```go
func (*GlGetProgramiv) Class() binary.Class
```

#### func (*GlGetProgramiv) Flags

```go
func (c *GlGetProgramiv) Flags() atom.Flags
```

#### func (*GlGetProgramiv) Mutate

```go
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetProgramiv) Replay

```go
func (ϟa *GlGetProgramiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetProgramiv) String

```go
func (c *GlGetProgramiv) String() string
```

#### func (*GlGetProgramiv) TypeID

```go
func (c *GlGetProgramiv) TypeID() atom.TypeID
```

#### type GlGetProgramiv_Postback

```go
type GlGetProgramiv_Postback struct {
	Value S32Array
}
```


#### func (*GlGetProgramiv_Postback) Decode

```go
func (o *GlGetProgramiv_Postback) Decode(value_cnt uint64, d binary.Decoder) error
```

#### type GlGetQueryObjecti64vEXT

```go
type GlGetQueryObjecti64vEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     int64
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjecti64vEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjecti64vEXT

```go
func NewGlGetQueryObjecti64vEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue int64,
) *GlGetQueryObjecti64vEXT
```

#### func (*GlGetQueryObjecti64vEXT) API

```go
func (c *GlGetQueryObjecti64vEXT) API() gfxapi.API
```

#### func (*GlGetQueryObjecti64vEXT) Class

```go
func (*GlGetQueryObjecti64vEXT) Class() binary.Class
```

#### func (*GlGetQueryObjecti64vEXT) Flags

```go
func (c *GlGetQueryObjecti64vEXT) Flags() atom.Flags
```

#### func (*GlGetQueryObjecti64vEXT) Mutate

```go
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetQueryObjecti64vEXT) Replay

```go
func (ϟa *GlGetQueryObjecti64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetQueryObjecti64vEXT) String

```go
func (c *GlGetQueryObjecti64vEXT) String() string
```

#### func (*GlGetQueryObjecti64vEXT) TypeID

```go
func (c *GlGetQueryObjecti64vEXT) TypeID() atom.TypeID
```

#### type GlGetQueryObjecti64vEXT_Postback

```go
type GlGetQueryObjecti64vEXT_Postback struct {
	Value int64
}
```


#### func (*GlGetQueryObjecti64vEXT_Postback) Decode

```go
func (o *GlGetQueryObjecti64vEXT_Postback) Decode(d binary.Decoder) error
```

#### type GlGetQueryObjectivEXT

```go
type GlGetQueryObjectivEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectivEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectivEXT

```go
func NewGlGetQueryObjectivEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue int32,
) *GlGetQueryObjectivEXT
```

#### func (*GlGetQueryObjectivEXT) API

```go
func (c *GlGetQueryObjectivEXT) API() gfxapi.API
```

#### func (*GlGetQueryObjectivEXT) Class

```go
func (*GlGetQueryObjectivEXT) Class() binary.Class
```

#### func (*GlGetQueryObjectivEXT) Flags

```go
func (c *GlGetQueryObjectivEXT) Flags() atom.Flags
```

#### func (*GlGetQueryObjectivEXT) Mutate

```go
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetQueryObjectivEXT) Replay

```go
func (ϟa *GlGetQueryObjectivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetQueryObjectivEXT) String

```go
func (c *GlGetQueryObjectivEXT) String() string
```

#### func (*GlGetQueryObjectivEXT) TypeID

```go
func (c *GlGetQueryObjectivEXT) TypeID() atom.TypeID
```

#### type GlGetQueryObjectivEXT_Postback

```go
type GlGetQueryObjectivEXT_Postback struct {
	Value int32
}
```


#### func (*GlGetQueryObjectivEXT_Postback) Decode

```go
func (o *GlGetQueryObjectivEXT_Postback) Decode(d binary.Decoder) error
```

#### type GlGetQueryObjectui64vEXT

```go
type GlGetQueryObjectui64vEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     uint64
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectui64vEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectui64vEXT

```go
func NewGlGetQueryObjectui64vEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint64,
) *GlGetQueryObjectui64vEXT
```

#### func (*GlGetQueryObjectui64vEXT) API

```go
func (c *GlGetQueryObjectui64vEXT) API() gfxapi.API
```

#### func (*GlGetQueryObjectui64vEXT) Class

```go
func (*GlGetQueryObjectui64vEXT) Class() binary.Class
```

#### func (*GlGetQueryObjectui64vEXT) Flags

```go
func (c *GlGetQueryObjectui64vEXT) Flags() atom.Flags
```

#### func (*GlGetQueryObjectui64vEXT) Mutate

```go
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetQueryObjectui64vEXT) Replay

```go
func (ϟa *GlGetQueryObjectui64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetQueryObjectui64vEXT) String

```go
func (c *GlGetQueryObjectui64vEXT) String() string
```

#### func (*GlGetQueryObjectui64vEXT) TypeID

```go
func (c *GlGetQueryObjectui64vEXT) TypeID() atom.TypeID
```

#### type GlGetQueryObjectui64vEXT_Postback

```go
type GlGetQueryObjectui64vEXT_Postback struct {
	Value uint64
}
```


#### func (*GlGetQueryObjectui64vEXT_Postback) Decode

```go
func (o *GlGetQueryObjectui64vEXT_Postback) Decode(d binary.Decoder) error
```

#### type GlGetQueryObjectuiv

```go
type GlGetQueryObjectuiv struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     uint32
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectuiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectuiv

```go
func NewGlGetQueryObjectuiv(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint32,
) *GlGetQueryObjectuiv
```

#### func (*GlGetQueryObjectuiv) API

```go
func (c *GlGetQueryObjectuiv) API() gfxapi.API
```

#### func (*GlGetQueryObjectuiv) Class

```go
func (*GlGetQueryObjectuiv) Class() binary.Class
```

#### func (*GlGetQueryObjectuiv) Flags

```go
func (c *GlGetQueryObjectuiv) Flags() atom.Flags
```

#### func (*GlGetQueryObjectuiv) Mutate

```go
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetQueryObjectuiv) Replay

```go
func (ϟa *GlGetQueryObjectuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetQueryObjectuiv) String

```go
func (c *GlGetQueryObjectuiv) String() string
```

#### func (*GlGetQueryObjectuiv) TypeID

```go
func (c *GlGetQueryObjectuiv) TypeID() atom.TypeID
```

#### type GlGetQueryObjectuivEXT

```go
type GlGetQueryObjectuivEXT struct {
	binary.Generate
	Query     QueryId
	Parameter QueryObjectParameter
	Value     uint32
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectuivEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectuivEXT

```go
func NewGlGetQueryObjectuivEXT(
	pQuery QueryId,
	pParameter QueryObjectParameter,
	pValue uint32,
) *GlGetQueryObjectuivEXT
```

#### func (*GlGetQueryObjectuivEXT) API

```go
func (c *GlGetQueryObjectuivEXT) API() gfxapi.API
```

#### func (*GlGetQueryObjectuivEXT) Class

```go
func (*GlGetQueryObjectuivEXT) Class() binary.Class
```

#### func (*GlGetQueryObjectuivEXT) Flags

```go
func (c *GlGetQueryObjectuivEXT) Flags() atom.Flags
```

#### func (*GlGetQueryObjectuivEXT) Mutate

```go
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetQueryObjectuivEXT) Replay

```go
func (ϟa *GlGetQueryObjectuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetQueryObjectuivEXT) String

```go
func (c *GlGetQueryObjectuivEXT) String() string
```

#### func (*GlGetQueryObjectuivEXT) TypeID

```go
func (c *GlGetQueryObjectuivEXT) TypeID() atom.TypeID
```

#### type GlGetQueryObjectuivEXT_Postback

```go
type GlGetQueryObjectuivEXT_Postback struct {
	Value uint32
}
```


#### func (*GlGetQueryObjectuivEXT_Postback) Decode

```go
func (o *GlGetQueryObjectuivEXT_Postback) Decode(d binary.Decoder) error
```

#### type GlGetQueryObjectuiv_Postback

```go
type GlGetQueryObjectuiv_Postback struct {
	Value uint32
}
```


#### func (*GlGetQueryObjectuiv_Postback) Decode

```go
func (o *GlGetQueryObjectuiv_Postback) Decode(d binary.Decoder) error
```

#### type GlGetQueryiv

```go
type GlGetQueryiv struct {
	binary.Generate
	Target    QueryTarget
	Parameter QueryParameter
	Value     int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryiv

```go
func NewGlGetQueryiv(
	pTarget QueryTarget,
	pParameter QueryParameter,
	pValue int32,
) *GlGetQueryiv
```

#### func (*GlGetQueryiv) API

```go
func (c *GlGetQueryiv) API() gfxapi.API
```

#### func (*GlGetQueryiv) Class

```go
func (*GlGetQueryiv) Class() binary.Class
```

#### func (*GlGetQueryiv) Flags

```go
func (c *GlGetQueryiv) Flags() atom.Flags
```

#### func (*GlGetQueryiv) Mutate

```go
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetQueryiv) Replay

```go
func (ϟa *GlGetQueryiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetQueryiv) String

```go
func (c *GlGetQueryiv) String() string
```

#### func (*GlGetQueryiv) TypeID

```go
func (c *GlGetQueryiv) TypeID() atom.TypeID
```

#### type GlGetQueryivEXT

```go
type GlGetQueryivEXT struct {
	binary.Generate
	Target    QueryTarget
	Parameter QueryParameter
	Value     int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryivEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryivEXT

```go
func NewGlGetQueryivEXT(
	pTarget QueryTarget,
	pParameter QueryParameter,
	pValue int32,
) *GlGetQueryivEXT
```

#### func (*GlGetQueryivEXT) API

```go
func (c *GlGetQueryivEXT) API() gfxapi.API
```

#### func (*GlGetQueryivEXT) Class

```go
func (*GlGetQueryivEXT) Class() binary.Class
```

#### func (*GlGetQueryivEXT) Flags

```go
func (c *GlGetQueryivEXT) Flags() atom.Flags
```

#### func (*GlGetQueryivEXT) Mutate

```go
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetQueryivEXT) Replay

```go
func (ϟa *GlGetQueryivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetQueryivEXT) String

```go
func (c *GlGetQueryivEXT) String() string
```

#### func (*GlGetQueryivEXT) TypeID

```go
func (c *GlGetQueryivEXT) TypeID() atom.TypeID
```

#### type GlGetQueryivEXT_Postback

```go
type GlGetQueryivEXT_Postback struct {
	Value int32
}
```


#### func (*GlGetQueryivEXT_Postback) Decode

```go
func (o *GlGetQueryivEXT_Postback) Decode(d binary.Decoder) error
```

#### type GlGetQueryiv_Postback

```go
type GlGetQueryiv_Postback struct {
	Value int32
}
```


#### func (*GlGetQueryiv_Postback) Decode

```go
func (o *GlGetQueryiv_Postback) Decode(d binary.Decoder) error
```

#### type GlGetRenderbufferParameteriv

```go
type GlGetRenderbufferParameteriv struct {
	binary.Generate
	Target    RenderbufferTarget
	Parameter RenderbufferParameter
	Values    S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetRenderbufferParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetRenderbufferParameteriv

```go
func NewGlGetRenderbufferParameteriv(
	pTarget RenderbufferTarget,
	pParameter RenderbufferParameter,
	pValues S32Array,
) *GlGetRenderbufferParameteriv
```

#### func (*GlGetRenderbufferParameteriv) API

```go
func (c *GlGetRenderbufferParameteriv) API() gfxapi.API
```

#### func (*GlGetRenderbufferParameteriv) Class

```go
func (*GlGetRenderbufferParameteriv) Class() binary.Class
```

#### func (*GlGetRenderbufferParameteriv) Flags

```go
func (c *GlGetRenderbufferParameteriv) Flags() atom.Flags
```

#### func (*GlGetRenderbufferParameteriv) Mutate

```go
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetRenderbufferParameteriv) Replay

```go
func (ϟa *GlGetRenderbufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetRenderbufferParameteriv) String

```go
func (c *GlGetRenderbufferParameteriv) String() string
```

#### func (*GlGetRenderbufferParameteriv) TypeID

```go
func (c *GlGetRenderbufferParameteriv) TypeID() atom.TypeID
```

#### type GlGetRenderbufferParameteriv_Postback

```go
type GlGetRenderbufferParameteriv_Postback struct {
	Values S32Array
}
```


#### func (*GlGetRenderbufferParameteriv_Postback) Decode

```go
func (o *GlGetRenderbufferParameteriv_Postback) Decode(values_cnt uint64, d binary.Decoder) error
```

#### type GlGetShaderInfoLog

```go
type GlGetShaderInfoLog struct {
	binary.Generate
	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten int32
	Info                string
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderInfoLog
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderInfoLog

```go
func NewGlGetShaderInfoLog(
	pShader ShaderId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pInfo string,
) *GlGetShaderInfoLog
```

#### func (*GlGetShaderInfoLog) API

```go
func (c *GlGetShaderInfoLog) API() gfxapi.API
```

#### func (*GlGetShaderInfoLog) Class

```go
func (*GlGetShaderInfoLog) Class() binary.Class
```

#### func (*GlGetShaderInfoLog) Flags

```go
func (c *GlGetShaderInfoLog) Flags() atom.Flags
```

#### func (*GlGetShaderInfoLog) Mutate

```go
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetShaderInfoLog) Replay

```go
func (ϟa *GlGetShaderInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetShaderInfoLog) String

```go
func (c *GlGetShaderInfoLog) String() string
```

#### func (*GlGetShaderInfoLog) TypeID

```go
func (c *GlGetShaderInfoLog) TypeID() atom.TypeID
```

#### type GlGetShaderInfoLog_Postback

```go
type GlGetShaderInfoLog_Postback struct {
	StringLengthWritten int32
	Info                string
}
```


#### func (*GlGetShaderInfoLog_Postback) Decode

```go
func (o *GlGetShaderInfoLog_Postback) Decode(info_cnt uint64, d binary.Decoder) error
```

#### type GlGetShaderPrecisionFormat

```go
type GlGetShaderPrecisionFormat struct {
	binary.Generate
	ShaderType    ShaderType
	PrecisionType PrecisionType
	Range         S32Array
	Precision     int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderPrecisionFormat
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderPrecisionFormat

```go
func NewGlGetShaderPrecisionFormat(
	pShaderType ShaderType,
	pPrecisionType PrecisionType,
	pRange S32Array,
	pPrecision int32,
) *GlGetShaderPrecisionFormat
```

#### func (*GlGetShaderPrecisionFormat) API

```go
func (c *GlGetShaderPrecisionFormat) API() gfxapi.API
```

#### func (*GlGetShaderPrecisionFormat) Class

```go
func (*GlGetShaderPrecisionFormat) Class() binary.Class
```

#### func (*GlGetShaderPrecisionFormat) Flags

```go
func (c *GlGetShaderPrecisionFormat) Flags() atom.Flags
```

#### func (*GlGetShaderPrecisionFormat) Mutate

```go
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetShaderPrecisionFormat) Replay

```go
func (ϟa *GlGetShaderPrecisionFormat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetShaderPrecisionFormat) String

```go
func (c *GlGetShaderPrecisionFormat) String() string
```

#### func (*GlGetShaderPrecisionFormat) TypeID

```go
func (c *GlGetShaderPrecisionFormat) TypeID() atom.TypeID
```

#### type GlGetShaderPrecisionFormat_Postback

```go
type GlGetShaderPrecisionFormat_Postback struct {
	Range     S32Array
	Precision int32
}
```


#### func (*GlGetShaderPrecisionFormat_Postback) Decode

```go
func (o *GlGetShaderPrecisionFormat_Postback) Decode(range_cnt uint64, d binary.Decoder) error
```

#### type GlGetShaderSource

```go
type GlGetShaderSource struct {
	binary.Generate
	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten int32
	Source              string
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderSource
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderSource

```go
func NewGlGetShaderSource(
	pShader ShaderId,
	pBufferLength int32,
	pStringLengthWritten int32,
	pSource string,
) *GlGetShaderSource
```

#### func (*GlGetShaderSource) API

```go
func (c *GlGetShaderSource) API() gfxapi.API
```

#### func (*GlGetShaderSource) Class

```go
func (*GlGetShaderSource) Class() binary.Class
```

#### func (*GlGetShaderSource) Flags

```go
func (c *GlGetShaderSource) Flags() atom.Flags
```

#### func (*GlGetShaderSource) Mutate

```go
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetShaderSource) Replay

```go
func (ϟa *GlGetShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetShaderSource) String

```go
func (c *GlGetShaderSource) String() string
```

#### func (*GlGetShaderSource) TypeID

```go
func (c *GlGetShaderSource) TypeID() atom.TypeID
```

#### type GlGetShaderSource_Postback

```go
type GlGetShaderSource_Postback struct {
	StringLengthWritten int32
	Source              string
}
```


#### func (*GlGetShaderSource_Postback) Decode

```go
func (o *GlGetShaderSource_Postback) Decode(source_cnt uint64, d binary.Decoder) error
```

#### type GlGetShaderiv

```go
type GlGetShaderiv struct {
	binary.Generate
	Shader    ShaderId
	Parameter ShaderParameter
	Value     S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderiv

```go
func NewGlGetShaderiv(
	pShader ShaderId,
	pParameter ShaderParameter,
	pValue S32Array,
) *GlGetShaderiv
```

#### func (*GlGetShaderiv) API

```go
func (c *GlGetShaderiv) API() gfxapi.API
```

#### func (*GlGetShaderiv) Class

```go
func (*GlGetShaderiv) Class() binary.Class
```

#### func (*GlGetShaderiv) Flags

```go
func (c *GlGetShaderiv) Flags() atom.Flags
```

#### func (*GlGetShaderiv) Mutate

```go
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetShaderiv) Replay

```go
func (ϟa *GlGetShaderiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetShaderiv) String

```go
func (c *GlGetShaderiv) String() string
```

#### func (*GlGetShaderiv) TypeID

```go
func (c *GlGetShaderiv) TypeID() atom.TypeID
```

#### type GlGetShaderiv_Postback

```go
type GlGetShaderiv_Postback struct {
	Value S32Array
}
```


#### func (*GlGetShaderiv_Postback) Decode

```go
func (o *GlGetShaderiv_Postback) Decode(value_cnt uint64, d binary.Decoder) error
```

#### type GlGetString

```go
type GlGetString struct {
	binary.Generate
	Param  StringConstant
	Result string
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetString
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetString

```go
func NewGlGetString(
	pParam StringConstant,
	pResult string,
) *GlGetString
```

#### func (*GlGetString) API

```go
func (c *GlGetString) API() gfxapi.API
```

#### func (*GlGetString) Class

```go
func (*GlGetString) Class() binary.Class
```

#### func (*GlGetString) Flags

```go
func (c *GlGetString) Flags() atom.Flags
```

#### func (*GlGetString) Mutate

```go
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetString) Replay

```go
func (ϟa *GlGetString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetString) String

```go
func (c *GlGetString) String() string
```

#### func (*GlGetString) TypeID

```go
func (c *GlGetString) TypeID() atom.TypeID
```

#### type GlGetString_Postback

```go
type GlGetString_Postback struct {
	Result string
}
```


#### func (*GlGetString_Postback) Decode

```go
func (o *GlGetString_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type GlGetTexParameterfv

```go
type GlGetTexParameterfv struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Values    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetTexParameterfv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetTexParameterfv

```go
func NewGlGetTexParameterfv(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValues F32Array,
) *GlGetTexParameterfv
```

#### func (*GlGetTexParameterfv) API

```go
func (c *GlGetTexParameterfv) API() gfxapi.API
```

#### func (*GlGetTexParameterfv) Class

```go
func (*GlGetTexParameterfv) Class() binary.Class
```

#### func (*GlGetTexParameterfv) Flags

```go
func (c *GlGetTexParameterfv) Flags() atom.Flags
```

#### func (*GlGetTexParameterfv) Mutate

```go
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetTexParameterfv) Replay

```go
func (ϟa *GlGetTexParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetTexParameterfv) String

```go
func (c *GlGetTexParameterfv) String() string
```

#### func (*GlGetTexParameterfv) TypeID

```go
func (c *GlGetTexParameterfv) TypeID() atom.TypeID
```

#### type GlGetTexParameterfv_Postback

```go
type GlGetTexParameterfv_Postback struct {
	Values F32Array
}
```


#### func (*GlGetTexParameterfv_Postback) Decode

```go
func (o *GlGetTexParameterfv_Postback) Decode(values_cnt uint64, d binary.Decoder) error
```

#### type GlGetTexParameteriv

```go
type GlGetTexParameteriv struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Values    S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetTexParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetTexParameteriv

```go
func NewGlGetTexParameteriv(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValues S32Array,
) *GlGetTexParameteriv
```

#### func (*GlGetTexParameteriv) API

```go
func (c *GlGetTexParameteriv) API() gfxapi.API
```

#### func (*GlGetTexParameteriv) Class

```go
func (*GlGetTexParameteriv) Class() binary.Class
```

#### func (*GlGetTexParameteriv) Flags

```go
func (c *GlGetTexParameteriv) Flags() atom.Flags
```

#### func (*GlGetTexParameteriv) Mutate

```go
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetTexParameteriv) Replay

```go
func (ϟa *GlGetTexParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetTexParameteriv) String

```go
func (c *GlGetTexParameteriv) String() string
```

#### func (*GlGetTexParameteriv) TypeID

```go
func (c *GlGetTexParameteriv) TypeID() atom.TypeID
```

#### type GlGetTexParameteriv_Postback

```go
type GlGetTexParameteriv_Postback struct {
	Values S32Array
}
```


#### func (*GlGetTexParameteriv_Postback) Decode

```go
func (o *GlGetTexParameteriv_Postback) Decode(values_cnt uint64, d binary.Decoder) error
```

#### type GlGetUniformLocation

```go
type GlGetUniformLocation struct {
	binary.Generate
	Program ProgramId
	Name    string
	Result  UniformLocation
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetUniformLocation
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetUniformLocation

```go
func NewGlGetUniformLocation(
	pProgram ProgramId,
	pName string,
	pResult UniformLocation,
) *GlGetUniformLocation
```

#### func (*GlGetUniformLocation) API

```go
func (c *GlGetUniformLocation) API() gfxapi.API
```

#### func (*GlGetUniformLocation) Class

```go
func (*GlGetUniformLocation) Class() binary.Class
```

#### func (*GlGetUniformLocation) Flags

```go
func (c *GlGetUniformLocation) Flags() atom.Flags
```

#### func (*GlGetUniformLocation) Mutate

```go
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetUniformLocation) Replay

```go
func (ϟa *GlGetUniformLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetUniformLocation) String

```go
func (c *GlGetUniformLocation) String() string
```

#### func (*GlGetUniformLocation) TypeID

```go
func (c *GlGetUniformLocation) TypeID() atom.TypeID
```

#### type GlGetUniformLocation_Postback

```go
type GlGetUniformLocation_Postback struct {
	Result UniformLocation
}
```


#### func (*GlGetUniformLocation_Postback) Decode

```go
func (o *GlGetUniformLocation_Postback) Decode(d binary.Decoder) error
```

#### type GlGetUniformfv

```go
type GlGetUniformfv struct {
	binary.Generate
	Program  ProgramId
	Location UniformLocation
	Values   F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetUniformfv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetUniformfv

```go
func NewGlGetUniformfv(
	pProgram ProgramId,
	pLocation UniformLocation,
	pValues F32Array,
) *GlGetUniformfv
```

#### func (*GlGetUniformfv) API

```go
func (c *GlGetUniformfv) API() gfxapi.API
```

#### func (*GlGetUniformfv) Class

```go
func (*GlGetUniformfv) Class() binary.Class
```

#### func (*GlGetUniformfv) Flags

```go
func (c *GlGetUniformfv) Flags() atom.Flags
```

#### func (*GlGetUniformfv) Mutate

```go
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetUniformfv) Replay

```go
func (ϟa *GlGetUniformfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetUniformfv) String

```go
func (c *GlGetUniformfv) String() string
```

#### func (*GlGetUniformfv) TypeID

```go
func (c *GlGetUniformfv) TypeID() atom.TypeID
```

#### type GlGetUniformiv

```go
type GlGetUniformiv struct {
	binary.Generate
	Program  ProgramId
	Location UniformLocation
	Values   S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetUniformiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetUniformiv

```go
func NewGlGetUniformiv(
	pProgram ProgramId,
	pLocation UniformLocation,
	pValues S32Array,
) *GlGetUniformiv
```

#### func (*GlGetUniformiv) API

```go
func (c *GlGetUniformiv) API() gfxapi.API
```

#### func (*GlGetUniformiv) Class

```go
func (*GlGetUniformiv) Class() binary.Class
```

#### func (*GlGetUniformiv) Flags

```go
func (c *GlGetUniformiv) Flags() atom.Flags
```

#### func (*GlGetUniformiv) Mutate

```go
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlGetUniformiv) Replay

```go
func (ϟa *GlGetUniformiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlGetUniformiv) String

```go
func (c *GlGetUniformiv) String() string
```

#### func (*GlGetUniformiv) TypeID

```go
func (c *GlGetUniformiv) TypeID() atom.TypeID
```

#### type GlHint

```go
type GlHint struct {
	binary.Generate
	Target HintTarget
	Mode   HintMode
}
```

//////////////////////////////////////////////////////////////////////////////
GlHint
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlHint

```go
func NewGlHint(
	pTarget HintTarget,
	pMode HintMode,
) *GlHint
```

#### func (*GlHint) API

```go
func (c *GlHint) API() gfxapi.API
```

#### func (*GlHint) Class

```go
func (*GlHint) Class() binary.Class
```

#### func (*GlHint) Flags

```go
func (c *GlHint) Flags() atom.Flags
```

#### func (*GlHint) Mutate

```go
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlHint) Replay

```go
func (ϟa *GlHint) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlHint) String

```go
func (c *GlHint) String() string
```

#### func (*GlHint) TypeID

```go
func (c *GlHint) TypeID() atom.TypeID
```

#### type GlInsertEventMarkerEXT

```go
type GlInsertEventMarkerEXT struct {
	binary.Generate
	Length int32
	Marker string
}
```

//////////////////////////////////////////////////////////////////////////////
GlInsertEventMarkerEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlInsertEventMarkerEXT

```go
func NewGlInsertEventMarkerEXT(
	pLength int32,
	pMarker string,
) *GlInsertEventMarkerEXT
```

#### func (*GlInsertEventMarkerEXT) API

```go
func (c *GlInsertEventMarkerEXT) API() gfxapi.API
```

#### func (*GlInsertEventMarkerEXT) Class

```go
func (*GlInsertEventMarkerEXT) Class() binary.Class
```

#### func (*GlInsertEventMarkerEXT) Flags

```go
func (c *GlInsertEventMarkerEXT) Flags() atom.Flags
```

#### func (*GlInsertEventMarkerEXT) Mutate

```go
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlInsertEventMarkerEXT) Replay

```go
func (ϟa *GlInsertEventMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlInsertEventMarkerEXT) String

```go
func (c *GlInsertEventMarkerEXT) String() string
```

#### func (*GlInsertEventMarkerEXT) TypeID

```go
func (c *GlInsertEventMarkerEXT) TypeID() atom.TypeID
```

#### type GlInvalidateFramebuffer

```go
type GlInvalidateFramebuffer struct {
	binary.Generate
	Target      FramebufferTarget
	Count       int32
	Attachments FramebufferAttachmentArray
}
```

//////////////////////////////////////////////////////////////////////////////
GlInvalidateFramebuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlInvalidateFramebuffer

```go
func NewGlInvalidateFramebuffer(
	pTarget FramebufferTarget,
	pCount int32,
	pAttachments FramebufferAttachmentArray,
) *GlInvalidateFramebuffer
```

#### func (*GlInvalidateFramebuffer) API

```go
func (c *GlInvalidateFramebuffer) API() gfxapi.API
```

#### func (*GlInvalidateFramebuffer) Class

```go
func (*GlInvalidateFramebuffer) Class() binary.Class
```

#### func (*GlInvalidateFramebuffer) Flags

```go
func (c *GlInvalidateFramebuffer) Flags() atom.Flags
```

#### func (*GlInvalidateFramebuffer) Mutate

```go
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlInvalidateFramebuffer) Replay

```go
func (ϟa *GlInvalidateFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlInvalidateFramebuffer) String

```go
func (c *GlInvalidateFramebuffer) String() string
```

#### func (*GlInvalidateFramebuffer) TypeID

```go
func (c *GlInvalidateFramebuffer) TypeID() atom.TypeID
```

#### type GlIsBuffer

```go
type GlIsBuffer struct {
	binary.Generate
	Buffer BufferId
	Result bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsBuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsBuffer

```go
func NewGlIsBuffer(
	pBuffer BufferId,
	pResult bool,
) *GlIsBuffer
```

#### func (*GlIsBuffer) API

```go
func (c *GlIsBuffer) API() gfxapi.API
```

#### func (*GlIsBuffer) Class

```go
func (*GlIsBuffer) Class() binary.Class
```

#### func (*GlIsBuffer) Flags

```go
func (c *GlIsBuffer) Flags() atom.Flags
```

#### func (*GlIsBuffer) Mutate

```go
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsBuffer) Replay

```go
func (ϟa *GlIsBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsBuffer) String

```go
func (c *GlIsBuffer) String() string
```

#### func (*GlIsBuffer) TypeID

```go
func (c *GlIsBuffer) TypeID() atom.TypeID
```

#### type GlIsBuffer_Postback

```go
type GlIsBuffer_Postback struct {
	Result bool
}
```


#### func (*GlIsBuffer_Postback) Decode

```go
func (o *GlIsBuffer_Postback) Decode(d binary.Decoder) error
```

#### type GlIsEnabled

```go
type GlIsEnabled struct {
	binary.Generate
	Capability Capability
	Result     bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsEnabled
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsEnabled

```go
func NewGlIsEnabled(
	pCapability Capability,
	pResult bool,
) *GlIsEnabled
```

#### func (*GlIsEnabled) API

```go
func (c *GlIsEnabled) API() gfxapi.API
```

#### func (*GlIsEnabled) Class

```go
func (*GlIsEnabled) Class() binary.Class
```

#### func (*GlIsEnabled) Flags

```go
func (c *GlIsEnabled) Flags() atom.Flags
```

#### func (*GlIsEnabled) Mutate

```go
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsEnabled) Replay

```go
func (ϟa *GlIsEnabled) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsEnabled) String

```go
func (c *GlIsEnabled) String() string
```

#### func (*GlIsEnabled) TypeID

```go
func (c *GlIsEnabled) TypeID() atom.TypeID
```

#### type GlIsEnabled_Postback

```go
type GlIsEnabled_Postback struct {
	Result bool
}
```


#### func (*GlIsEnabled_Postback) Decode

```go
func (o *GlIsEnabled_Postback) Decode(d binary.Decoder) error
```

#### type GlIsFramebuffer

```go
type GlIsFramebuffer struct {
	binary.Generate
	Framebuffer FramebufferId
	Result      bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsFramebuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsFramebuffer

```go
func NewGlIsFramebuffer(
	pFramebuffer FramebufferId,
	pResult bool,
) *GlIsFramebuffer
```

#### func (*GlIsFramebuffer) API

```go
func (c *GlIsFramebuffer) API() gfxapi.API
```

#### func (*GlIsFramebuffer) Class

```go
func (*GlIsFramebuffer) Class() binary.Class
```

#### func (*GlIsFramebuffer) Flags

```go
func (c *GlIsFramebuffer) Flags() atom.Flags
```

#### func (*GlIsFramebuffer) Mutate

```go
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsFramebuffer) Replay

```go
func (ϟa *GlIsFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsFramebuffer) String

```go
func (c *GlIsFramebuffer) String() string
```

#### func (*GlIsFramebuffer) TypeID

```go
func (c *GlIsFramebuffer) TypeID() atom.TypeID
```

#### type GlIsFramebuffer_Postback

```go
type GlIsFramebuffer_Postback struct {
	Result bool
}
```


#### func (*GlIsFramebuffer_Postback) Decode

```go
func (o *GlIsFramebuffer_Postback) Decode(d binary.Decoder) error
```

#### type GlIsProgram

```go
type GlIsProgram struct {
	binary.Generate
	Program ProgramId
	Result  bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsProgram
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsProgram

```go
func NewGlIsProgram(
	pProgram ProgramId,
	pResult bool,
) *GlIsProgram
```

#### func (*GlIsProgram) API

```go
func (c *GlIsProgram) API() gfxapi.API
```

#### func (*GlIsProgram) Class

```go
func (*GlIsProgram) Class() binary.Class
```

#### func (*GlIsProgram) Flags

```go
func (c *GlIsProgram) Flags() atom.Flags
```

#### func (*GlIsProgram) Mutate

```go
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsProgram) Replay

```go
func (ϟa *GlIsProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsProgram) String

```go
func (c *GlIsProgram) String() string
```

#### func (*GlIsProgram) TypeID

```go
func (c *GlIsProgram) TypeID() atom.TypeID
```

#### type GlIsProgram_Postback

```go
type GlIsProgram_Postback struct {
	Result bool
}
```


#### func (*GlIsProgram_Postback) Decode

```go
func (o *GlIsProgram_Postback) Decode(d binary.Decoder) error
```

#### type GlIsQuery

```go
type GlIsQuery struct {
	binary.Generate
	Query  QueryId
	Result bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsQuery
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsQuery

```go
func NewGlIsQuery(
	pQuery QueryId,
	pResult bool,
) *GlIsQuery
```

#### func (*GlIsQuery) API

```go
func (c *GlIsQuery) API() gfxapi.API
```

#### func (*GlIsQuery) Class

```go
func (*GlIsQuery) Class() binary.Class
```

#### func (*GlIsQuery) Flags

```go
func (c *GlIsQuery) Flags() atom.Flags
```

#### func (*GlIsQuery) Mutate

```go
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsQuery) Replay

```go
func (ϟa *GlIsQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsQuery) String

```go
func (c *GlIsQuery) String() string
```

#### func (*GlIsQuery) TypeID

```go
func (c *GlIsQuery) TypeID() atom.TypeID
```

#### type GlIsQueryEXT

```go
type GlIsQueryEXT struct {
	binary.Generate
	Query  QueryId
	Result bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsQueryEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsQueryEXT

```go
func NewGlIsQueryEXT(
	pQuery QueryId,
	pResult bool,
) *GlIsQueryEXT
```

#### func (*GlIsQueryEXT) API

```go
func (c *GlIsQueryEXT) API() gfxapi.API
```

#### func (*GlIsQueryEXT) Class

```go
func (*GlIsQueryEXT) Class() binary.Class
```

#### func (*GlIsQueryEXT) Flags

```go
func (c *GlIsQueryEXT) Flags() atom.Flags
```

#### func (*GlIsQueryEXT) Mutate

```go
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsQueryEXT) Replay

```go
func (ϟa *GlIsQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsQueryEXT) String

```go
func (c *GlIsQueryEXT) String() string
```

#### func (*GlIsQueryEXT) TypeID

```go
func (c *GlIsQueryEXT) TypeID() atom.TypeID
```

#### type GlIsQueryEXT_Postback

```go
type GlIsQueryEXT_Postback struct {
	Result bool
}
```


#### func (*GlIsQueryEXT_Postback) Decode

```go
func (o *GlIsQueryEXT_Postback) Decode(d binary.Decoder) error
```

#### type GlIsQuery_Postback

```go
type GlIsQuery_Postback struct {
	Result bool
}
```


#### func (*GlIsQuery_Postback) Decode

```go
func (o *GlIsQuery_Postback) Decode(d binary.Decoder) error
```

#### type GlIsRenderbuffer

```go
type GlIsRenderbuffer struct {
	binary.Generate
	Renderbuffer RenderbufferId
	Result       bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsRenderbuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsRenderbuffer

```go
func NewGlIsRenderbuffer(
	pRenderbuffer RenderbufferId,
	pResult bool,
) *GlIsRenderbuffer
```

#### func (*GlIsRenderbuffer) API

```go
func (c *GlIsRenderbuffer) API() gfxapi.API
```

#### func (*GlIsRenderbuffer) Class

```go
func (*GlIsRenderbuffer) Class() binary.Class
```

#### func (*GlIsRenderbuffer) Flags

```go
func (c *GlIsRenderbuffer) Flags() atom.Flags
```

#### func (*GlIsRenderbuffer) Mutate

```go
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsRenderbuffer) Replay

```go
func (ϟa *GlIsRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsRenderbuffer) String

```go
func (c *GlIsRenderbuffer) String() string
```

#### func (*GlIsRenderbuffer) TypeID

```go
func (c *GlIsRenderbuffer) TypeID() atom.TypeID
```

#### type GlIsRenderbuffer_Postback

```go
type GlIsRenderbuffer_Postback struct {
	Result bool
}
```


#### func (*GlIsRenderbuffer_Postback) Decode

```go
func (o *GlIsRenderbuffer_Postback) Decode(d binary.Decoder) error
```

#### type GlIsShader

```go
type GlIsShader struct {
	binary.Generate
	Shader ShaderId
	Result bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsShader
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsShader

```go
func NewGlIsShader(
	pShader ShaderId,
	pResult bool,
) *GlIsShader
```

#### func (*GlIsShader) API

```go
func (c *GlIsShader) API() gfxapi.API
```

#### func (*GlIsShader) Class

```go
func (*GlIsShader) Class() binary.Class
```

#### func (*GlIsShader) Flags

```go
func (c *GlIsShader) Flags() atom.Flags
```

#### func (*GlIsShader) Mutate

```go
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsShader) Replay

```go
func (ϟa *GlIsShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsShader) String

```go
func (c *GlIsShader) String() string
```

#### func (*GlIsShader) TypeID

```go
func (c *GlIsShader) TypeID() atom.TypeID
```

#### type GlIsShader_Postback

```go
type GlIsShader_Postback struct {
	Result bool
}
```


#### func (*GlIsShader_Postback) Decode

```go
func (o *GlIsShader_Postback) Decode(d binary.Decoder) error
```

#### type GlIsTexture

```go
type GlIsTexture struct {
	binary.Generate
	Texture TextureId
	Result  bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsTexture
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsTexture

```go
func NewGlIsTexture(
	pTexture TextureId,
	pResult bool,
) *GlIsTexture
```

#### func (*GlIsTexture) API

```go
func (c *GlIsTexture) API() gfxapi.API
```

#### func (*GlIsTexture) Class

```go
func (*GlIsTexture) Class() binary.Class
```

#### func (*GlIsTexture) Flags

```go
func (c *GlIsTexture) Flags() atom.Flags
```

#### func (*GlIsTexture) Mutate

```go
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsTexture) Replay

```go
func (ϟa *GlIsTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsTexture) String

```go
func (c *GlIsTexture) String() string
```

#### func (*GlIsTexture) TypeID

```go
func (c *GlIsTexture) TypeID() atom.TypeID
```

#### type GlIsTexture_Postback

```go
type GlIsTexture_Postback struct {
	Result bool
}
```


#### func (*GlIsTexture_Postback) Decode

```go
func (o *GlIsTexture_Postback) Decode(d binary.Decoder) error
```

#### type GlIsVertexArrayOES

```go
type GlIsVertexArrayOES struct {
	binary.Generate
	Array  VertexArrayId
	Result bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlIsVertexArrayOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlIsVertexArrayOES

```go
func NewGlIsVertexArrayOES(
	pArray VertexArrayId,
	pResult bool,
) *GlIsVertexArrayOES
```

#### func (*GlIsVertexArrayOES) API

```go
func (c *GlIsVertexArrayOES) API() gfxapi.API
```

#### func (*GlIsVertexArrayOES) Class

```go
func (*GlIsVertexArrayOES) Class() binary.Class
```

#### func (*GlIsVertexArrayOES) Flags

```go
func (c *GlIsVertexArrayOES) Flags() atom.Flags
```

#### func (*GlIsVertexArrayOES) Mutate

```go
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlIsVertexArrayOES) Replay

```go
func (ϟa *GlIsVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlIsVertexArrayOES) String

```go
func (c *GlIsVertexArrayOES) String() string
```

#### func (*GlIsVertexArrayOES) TypeID

```go
func (c *GlIsVertexArrayOES) TypeID() atom.TypeID
```

#### type GlIsVertexArrayOES_Postback

```go
type GlIsVertexArrayOES_Postback struct {
	Result bool
}
```


#### func (*GlIsVertexArrayOES_Postback) Decode

```go
func (o *GlIsVertexArrayOES_Postback) Decode(d binary.Decoder) error
```

#### type GlLineWidth

```go
type GlLineWidth struct {
	binary.Generate
	Width float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlLineWidth
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlLineWidth

```go
func NewGlLineWidth(
	pWidth float32,
) *GlLineWidth
```

#### func (*GlLineWidth) API

```go
func (c *GlLineWidth) API() gfxapi.API
```

#### func (*GlLineWidth) Class

```go
func (*GlLineWidth) Class() binary.Class
```

#### func (*GlLineWidth) Flags

```go
func (c *GlLineWidth) Flags() atom.Flags
```

#### func (*GlLineWidth) Mutate

```go
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlLineWidth) Replay

```go
func (ϟa *GlLineWidth) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlLineWidth) String

```go
func (c *GlLineWidth) String() string
```

#### func (*GlLineWidth) TypeID

```go
func (c *GlLineWidth) TypeID() atom.TypeID
```

#### type GlLinkProgram

```go
type GlLinkProgram struct {
	binary.Generate
	Program ProgramId
}
```

//////////////////////////////////////////////////////////////////////////////
GlLinkProgram
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlLinkProgram

```go
func NewGlLinkProgram(
	pProgram ProgramId,
) *GlLinkProgram
```

#### func (*GlLinkProgram) API

```go
func (c *GlLinkProgram) API() gfxapi.API
```

#### func (*GlLinkProgram) Class

```go
func (*GlLinkProgram) Class() binary.Class
```

#### func (*GlLinkProgram) Flags

```go
func (c *GlLinkProgram) Flags() atom.Flags
```

#### func (*GlLinkProgram) Mutate

```go
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlLinkProgram) Replay

```go
func (ϟa *GlLinkProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlLinkProgram) String

```go
func (c *GlLinkProgram) String() string
```

#### func (*GlLinkProgram) TypeID

```go
func (c *GlLinkProgram) TypeID() atom.TypeID
```

#### type GlMapBufferRange

```go
type GlMapBufferRange struct {
	binary.Generate
	Target BufferTarget
	Offset int32
	Length int32
	Access MapBufferRangeAccess
	Result memory.Pointer
}
```

//////////////////////////////////////////////////////////////////////////////
GlMapBufferRange
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlMapBufferRange

```go
func NewGlMapBufferRange(
	pTarget BufferTarget,
	pOffset int32,
	pLength int32,
	pAccess MapBufferRangeAccess,
	pResult memory.Pointer,
) *GlMapBufferRange
```

#### func (*GlMapBufferRange) API

```go
func (c *GlMapBufferRange) API() gfxapi.API
```

#### func (*GlMapBufferRange) Class

```go
func (*GlMapBufferRange) Class() binary.Class
```

#### func (*GlMapBufferRange) Flags

```go
func (c *GlMapBufferRange) Flags() atom.Flags
```

#### func (*GlMapBufferRange) Mutate

```go
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlMapBufferRange) Replay

```go
func (ϟa *GlMapBufferRange) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlMapBufferRange) String

```go
func (c *GlMapBufferRange) String() string
```

#### func (*GlMapBufferRange) TypeID

```go
func (c *GlMapBufferRange) TypeID() atom.TypeID
```

#### type GlMapBufferRange_Postback

```go
type GlMapBufferRange_Postback struct {
	Result []byte
}
```


#### func (*GlMapBufferRange_Postback) Decode

```go
func (o *GlMapBufferRange_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type GlPixelStorei

```go
type GlPixelStorei struct {
	binary.Generate
	Parameter PixelStoreParameter
	Value     int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlPixelStorei
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlPixelStorei

```go
func NewGlPixelStorei(
	pParameter PixelStoreParameter,
	pValue int32,
) *GlPixelStorei
```

#### func (*GlPixelStorei) API

```go
func (c *GlPixelStorei) API() gfxapi.API
```

#### func (*GlPixelStorei) Class

```go
func (*GlPixelStorei) Class() binary.Class
```

#### func (*GlPixelStorei) Flags

```go
func (c *GlPixelStorei) Flags() atom.Flags
```

#### func (*GlPixelStorei) Mutate

```go
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlPixelStorei) Replay

```go
func (ϟa *GlPixelStorei) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlPixelStorei) String

```go
func (c *GlPixelStorei) String() string
```

#### func (*GlPixelStorei) TypeID

```go
func (c *GlPixelStorei) TypeID() atom.TypeID
```

#### type GlPolygonOffset

```go
type GlPolygonOffset struct {
	binary.Generate
	ScaleFactor float32
	Units       float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlPolygonOffset
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlPolygonOffset

```go
func NewGlPolygonOffset(
	pScaleFactor float32,
	pUnits float32,
) *GlPolygonOffset
```

#### func (*GlPolygonOffset) API

```go
func (c *GlPolygonOffset) API() gfxapi.API
```

#### func (*GlPolygonOffset) Class

```go
func (*GlPolygonOffset) Class() binary.Class
```

#### func (*GlPolygonOffset) Flags

```go
func (c *GlPolygonOffset) Flags() atom.Flags
```

#### func (*GlPolygonOffset) Mutate

```go
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlPolygonOffset) Replay

```go
func (ϟa *GlPolygonOffset) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlPolygonOffset) String

```go
func (c *GlPolygonOffset) String() string
```

#### func (*GlPolygonOffset) TypeID

```go
func (c *GlPolygonOffset) TypeID() atom.TypeID
```

#### type GlPopGroupMarkerEXT

```go
type GlPopGroupMarkerEXT struct {
	binary.Generate
}
```

//////////////////////////////////////////////////////////////////////////////
GlPopGroupMarkerEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlPopGroupMarkerEXT

```go
func NewGlPopGroupMarkerEXT() *GlPopGroupMarkerEXT
```

#### func (*GlPopGroupMarkerEXT) API

```go
func (c *GlPopGroupMarkerEXT) API() gfxapi.API
```

#### func (*GlPopGroupMarkerEXT) Class

```go
func (*GlPopGroupMarkerEXT) Class() binary.Class
```

#### func (*GlPopGroupMarkerEXT) Flags

```go
func (c *GlPopGroupMarkerEXT) Flags() atom.Flags
```

#### func (*GlPopGroupMarkerEXT) Mutate

```go
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlPopGroupMarkerEXT) Replay

```go
func (ϟa *GlPopGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlPopGroupMarkerEXT) String

```go
func (c *GlPopGroupMarkerEXT) String() string
```

#### func (*GlPopGroupMarkerEXT) TypeID

```go
func (c *GlPopGroupMarkerEXT) TypeID() atom.TypeID
```

#### type GlProgramBinaryOES

```go
type GlProgramBinaryOES struct {
	binary.Generate
	Program      ProgramId
	BinaryFormat uint32
	Binary       memory.Pointer
	BinarySize   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlProgramBinaryOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlProgramBinaryOES

```go
func NewGlProgramBinaryOES(
	pProgram ProgramId,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
	pBinarySize int32,
) *GlProgramBinaryOES
```

#### func (*GlProgramBinaryOES) API

```go
func (c *GlProgramBinaryOES) API() gfxapi.API
```

#### func (*GlProgramBinaryOES) Class

```go
func (*GlProgramBinaryOES) Class() binary.Class
```

#### func (*GlProgramBinaryOES) Flags

```go
func (c *GlProgramBinaryOES) Flags() atom.Flags
```

#### func (*GlProgramBinaryOES) Mutate

```go
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlProgramBinaryOES) Replay

```go
func (ϟa *GlProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlProgramBinaryOES) String

```go
func (c *GlProgramBinaryOES) String() string
```

#### func (*GlProgramBinaryOES) TypeID

```go
func (c *GlProgramBinaryOES) TypeID() atom.TypeID
```

#### type GlPushGroupMarkerEXT

```go
type GlPushGroupMarkerEXT struct {
	binary.Generate
	Length int32
	Marker string
}
```

//////////////////////////////////////////////////////////////////////////////
GlPushGroupMarkerEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlPushGroupMarkerEXT

```go
func NewGlPushGroupMarkerEXT(
	pLength int32,
	pMarker string,
) *GlPushGroupMarkerEXT
```

#### func (*GlPushGroupMarkerEXT) API

```go
func (c *GlPushGroupMarkerEXT) API() gfxapi.API
```

#### func (*GlPushGroupMarkerEXT) Class

```go
func (*GlPushGroupMarkerEXT) Class() binary.Class
```

#### func (*GlPushGroupMarkerEXT) Flags

```go
func (c *GlPushGroupMarkerEXT) Flags() atom.Flags
```

#### func (*GlPushGroupMarkerEXT) Mutate

```go
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlPushGroupMarkerEXT) Replay

```go
func (ϟa *GlPushGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlPushGroupMarkerEXT) String

```go
func (c *GlPushGroupMarkerEXT) String() string
```

#### func (*GlPushGroupMarkerEXT) TypeID

```go
func (c *GlPushGroupMarkerEXT) TypeID() atom.TypeID
```

#### type GlQueryCounterEXT

```go
type GlQueryCounterEXT struct {
	binary.Generate
	Query  QueryId
	Target QueryTarget
}
```

//////////////////////////////////////////////////////////////////////////////
GlQueryCounterEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlQueryCounterEXT

```go
func NewGlQueryCounterEXT(
	pQuery QueryId,
	pTarget QueryTarget,
) *GlQueryCounterEXT
```

#### func (*GlQueryCounterEXT) API

```go
func (c *GlQueryCounterEXT) API() gfxapi.API
```

#### func (*GlQueryCounterEXT) Class

```go
func (*GlQueryCounterEXT) Class() binary.Class
```

#### func (*GlQueryCounterEXT) Flags

```go
func (c *GlQueryCounterEXT) Flags() atom.Flags
```

#### func (*GlQueryCounterEXT) Mutate

```go
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlQueryCounterEXT) Replay

```go
func (ϟa *GlQueryCounterEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlQueryCounterEXT) String

```go
func (c *GlQueryCounterEXT) String() string
```

#### func (*GlQueryCounterEXT) TypeID

```go
func (c *GlQueryCounterEXT) TypeID() atom.TypeID
```

#### type GlReadPixels

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlReadPixels
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlReadPixels

```go
func NewGlReadPixels(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pFormat BaseTexelFormat,
	pType TexelType,
	pData memory.Pointer,
) *GlReadPixels
```

#### func (*GlReadPixels) API

```go
func (c *GlReadPixels) API() gfxapi.API
```

#### func (*GlReadPixels) Class

```go
func (*GlReadPixels) Class() binary.Class
```

#### func (*GlReadPixels) Flags

```go
func (c *GlReadPixels) Flags() atom.Flags
```

#### func (*GlReadPixels) Mutate

```go
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlReadPixels) Replay

```go
func (ϟa *GlReadPixels) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlReadPixels) String

```go
func (c *GlReadPixels) String() string
```

#### func (*GlReadPixels) TypeID

```go
func (c *GlReadPixels) TypeID() atom.TypeID
```

#### type GlReadPixels_Postback

```go
type GlReadPixels_Postback struct {
	Data []byte
}
```


#### func (*GlReadPixels_Postback) Decode

```go
func (o *GlReadPixels_Postback) Decode(data_cnt uint64, d binary.Decoder) error
```

#### type GlReleaseShaderCompiler

```go
type GlReleaseShaderCompiler struct {
	binary.Generate
}
```

//////////////////////////////////////////////////////////////////////////////
GlReleaseShaderCompiler
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlReleaseShaderCompiler

```go
func NewGlReleaseShaderCompiler() *GlReleaseShaderCompiler
```

#### func (*GlReleaseShaderCompiler) API

```go
func (c *GlReleaseShaderCompiler) API() gfxapi.API
```

#### func (*GlReleaseShaderCompiler) Class

```go
func (*GlReleaseShaderCompiler) Class() binary.Class
```

#### func (*GlReleaseShaderCompiler) Flags

```go
func (c *GlReleaseShaderCompiler) Flags() atom.Flags
```

#### func (*GlReleaseShaderCompiler) Mutate

```go
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlReleaseShaderCompiler) Replay

```go
func (ϟa *GlReleaseShaderCompiler) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlReleaseShaderCompiler) String

```go
func (c *GlReleaseShaderCompiler) String() string
```

#### func (*GlReleaseShaderCompiler) TypeID

```go
func (c *GlReleaseShaderCompiler) TypeID() atom.TypeID
```

#### type GlRenderbufferStorage

```go
type GlRenderbufferStorage struct {
	binary.Generate
	Target RenderbufferTarget
	Format RenderbufferFormat
	Width  int32
	Height int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlRenderbufferStorage
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlRenderbufferStorage

```go
func NewGlRenderbufferStorage(
	pTarget RenderbufferTarget,
	pFormat RenderbufferFormat,
	pWidth int32,
	pHeight int32,
) *GlRenderbufferStorage
```

#### func (*GlRenderbufferStorage) API

```go
func (c *GlRenderbufferStorage) API() gfxapi.API
```

#### func (*GlRenderbufferStorage) Class

```go
func (*GlRenderbufferStorage) Class() binary.Class
```

#### func (*GlRenderbufferStorage) Flags

```go
func (c *GlRenderbufferStorage) Flags() atom.Flags
```

#### func (*GlRenderbufferStorage) Mutate

```go
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlRenderbufferStorage) Replay

```go
func (ϟa *GlRenderbufferStorage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlRenderbufferStorage) String

```go
func (c *GlRenderbufferStorage) String() string
```

#### func (*GlRenderbufferStorage) TypeID

```go
func (c *GlRenderbufferStorage) TypeID() atom.TypeID
```

#### type GlRenderbufferStorageMultisample

```go
type GlRenderbufferStorageMultisample struct {
	binary.Generate
	Target  RenderbufferTarget
	Samples int32
	Format  RenderbufferFormat
	Width   int32
	Height  int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlRenderbufferStorageMultisample
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlRenderbufferStorageMultisample

```go
func NewGlRenderbufferStorageMultisample(
	pTarget RenderbufferTarget,
	pSamples int32,
	pFormat RenderbufferFormat,
	pWidth int32,
	pHeight int32,
) *GlRenderbufferStorageMultisample
```

#### func (*GlRenderbufferStorageMultisample) API

```go
func (c *GlRenderbufferStorageMultisample) API() gfxapi.API
```

#### func (*GlRenderbufferStorageMultisample) Class

```go
func (*GlRenderbufferStorageMultisample) Class() binary.Class
```

#### func (*GlRenderbufferStorageMultisample) Flags

```go
func (c *GlRenderbufferStorageMultisample) Flags() atom.Flags
```

#### func (*GlRenderbufferStorageMultisample) Mutate

```go
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlRenderbufferStorageMultisample) Replay

```go
func (ϟa *GlRenderbufferStorageMultisample) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlRenderbufferStorageMultisample) String

```go
func (c *GlRenderbufferStorageMultisample) String() string
```

#### func (*GlRenderbufferStorageMultisample) TypeID

```go
func (c *GlRenderbufferStorageMultisample) TypeID() atom.TypeID
```

#### type GlSampleCoverage

```go
type GlSampleCoverage struct {
	binary.Generate
	Value  float32
	Invert bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlSampleCoverage
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlSampleCoverage

```go
func NewGlSampleCoverage(
	pValue float32,
	pInvert bool,
) *GlSampleCoverage
```

#### func (*GlSampleCoverage) API

```go
func (c *GlSampleCoverage) API() gfxapi.API
```

#### func (*GlSampleCoverage) Class

```go
func (*GlSampleCoverage) Class() binary.Class
```

#### func (*GlSampleCoverage) Flags

```go
func (c *GlSampleCoverage) Flags() atom.Flags
```

#### func (*GlSampleCoverage) Mutate

```go
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlSampleCoverage) Replay

```go
func (ϟa *GlSampleCoverage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlSampleCoverage) String

```go
func (c *GlSampleCoverage) String() string
```

#### func (*GlSampleCoverage) TypeID

```go
func (c *GlSampleCoverage) TypeID() atom.TypeID
```

#### type GlScissor

```go
type GlScissor struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlScissor
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlScissor

```go
func NewGlScissor(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlScissor
```

#### func (*GlScissor) API

```go
func (c *GlScissor) API() gfxapi.API
```

#### func (*GlScissor) Class

```go
func (*GlScissor) Class() binary.Class
```

#### func (*GlScissor) Flags

```go
func (c *GlScissor) Flags() atom.Flags
```

#### func (*GlScissor) Mutate

```go
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlScissor) Replay

```go
func (ϟa *GlScissor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlScissor) String

```go
func (c *GlScissor) String() string
```

#### func (*GlScissor) TypeID

```go
func (c *GlScissor) TypeID() atom.TypeID
```

#### type GlShaderBinary

```go
type GlShaderBinary struct {
	binary.Generate
	Count        int32
	Shaders      ShaderIdArray
	BinaryFormat uint32
	Binary       memory.Pointer
	BinarySize   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlShaderBinary
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlShaderBinary

```go
func NewGlShaderBinary(
	pCount int32,
	pShaders ShaderIdArray,
	pBinaryFormat uint32,
	pBinary memory.Pointer,
	pBinarySize int32,
) *GlShaderBinary
```

#### func (*GlShaderBinary) API

```go
func (c *GlShaderBinary) API() gfxapi.API
```

#### func (*GlShaderBinary) Class

```go
func (*GlShaderBinary) Class() binary.Class
```

#### func (*GlShaderBinary) Flags

```go
func (c *GlShaderBinary) Flags() atom.Flags
```

#### func (*GlShaderBinary) Mutate

```go
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlShaderBinary) Replay

```go
func (ϟa *GlShaderBinary) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlShaderBinary) String

```go
func (c *GlShaderBinary) String() string
```

#### func (*GlShaderBinary) TypeID

```go
func (c *GlShaderBinary) TypeID() atom.TypeID
```

#### type GlShaderSource

```go
type GlShaderSource struct {
	binary.Generate
	Shader ShaderId
	Count  int32
	Source StringArray
	Length S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlShaderSource
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlShaderSource

```go
func NewGlShaderSource(
	pShader ShaderId,
	pCount int32,
	pSource StringArray,
	pLength S32Array,
) *GlShaderSource
```

#### func (*GlShaderSource) API

```go
func (c *GlShaderSource) API() gfxapi.API
```

#### func (*GlShaderSource) Class

```go
func (*GlShaderSource) Class() binary.Class
```

#### func (*GlShaderSource) Flags

```go
func (c *GlShaderSource) Flags() atom.Flags
```

#### func (*GlShaderSource) Mutate

```go
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlShaderSource) Replay

```go
func (ϟa *GlShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlShaderSource) String

```go
func (c *GlShaderSource) String() string
```

#### func (*GlShaderSource) TypeID

```go
func (c *GlShaderSource) TypeID() atom.TypeID
```

#### type GlStartTilingQCOM

```go
type GlStartTilingQCOM struct {
	binary.Generate
	X            int32
	Y            int32
	Width        int32
	Height       int32
	PreserveMask TilePreserveMaskQCOM
}
```

//////////////////////////////////////////////////////////////////////////////
GlStartTilingQCOM
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlStartTilingQCOM

```go
func NewGlStartTilingQCOM(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
	pPreserveMask TilePreserveMaskQCOM,
) *GlStartTilingQCOM
```

#### func (*GlStartTilingQCOM) API

```go
func (c *GlStartTilingQCOM) API() gfxapi.API
```

#### func (*GlStartTilingQCOM) Class

```go
func (*GlStartTilingQCOM) Class() binary.Class
```

#### func (*GlStartTilingQCOM) Flags

```go
func (c *GlStartTilingQCOM) Flags() atom.Flags
```

#### func (*GlStartTilingQCOM) Mutate

```go
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlStartTilingQCOM) Replay

```go
func (ϟa *GlStartTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlStartTilingQCOM) String

```go
func (c *GlStartTilingQCOM) String() string
```

#### func (*GlStartTilingQCOM) TypeID

```go
func (c *GlStartTilingQCOM) TypeID() atom.TypeID
```

#### type GlStencilFuncSeparate

```go
type GlStencilFuncSeparate struct {
	binary.Generate
	Face           FaceMode
	Function       TestFunction
	ReferenceValue int32
	Mask           int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlStencilFuncSeparate
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlStencilFuncSeparate

```go
func NewGlStencilFuncSeparate(
	pFace FaceMode,
	pFunction TestFunction,
	pReferenceValue int32,
	pMask int32,
) *GlStencilFuncSeparate
```

#### func (*GlStencilFuncSeparate) API

```go
func (c *GlStencilFuncSeparate) API() gfxapi.API
```

#### func (*GlStencilFuncSeparate) Class

```go
func (*GlStencilFuncSeparate) Class() binary.Class
```

#### func (*GlStencilFuncSeparate) Flags

```go
func (c *GlStencilFuncSeparate) Flags() atom.Flags
```

#### func (*GlStencilFuncSeparate) Mutate

```go
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlStencilFuncSeparate) Replay

```go
func (ϟa *GlStencilFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlStencilFuncSeparate) String

```go
func (c *GlStencilFuncSeparate) String() string
```

#### func (*GlStencilFuncSeparate) TypeID

```go
func (c *GlStencilFuncSeparate) TypeID() atom.TypeID
```

#### type GlStencilMask

```go
type GlStencilMask struct {
	binary.Generate
	Mask uint32
}
```

//////////////////////////////////////////////////////////////////////////////
GlStencilMask
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlStencilMask

```go
func NewGlStencilMask(
	pMask uint32,
) *GlStencilMask
```

#### func (*GlStencilMask) API

```go
func (c *GlStencilMask) API() gfxapi.API
```

#### func (*GlStencilMask) Class

```go
func (*GlStencilMask) Class() binary.Class
```

#### func (*GlStencilMask) Flags

```go
func (c *GlStencilMask) Flags() atom.Flags
```

#### func (*GlStencilMask) Mutate

```go
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlStencilMask) Replay

```go
func (ϟa *GlStencilMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlStencilMask) String

```go
func (c *GlStencilMask) String() string
```

#### func (*GlStencilMask) TypeID

```go
func (c *GlStencilMask) TypeID() atom.TypeID
```

#### type GlStencilMaskSeparate

```go
type GlStencilMaskSeparate struct {
	binary.Generate
	Face FaceMode
	Mask uint32
}
```

//////////////////////////////////////////////////////////////////////////////
GlStencilMaskSeparate
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlStencilMaskSeparate

```go
func NewGlStencilMaskSeparate(
	pFace FaceMode,
	pMask uint32,
) *GlStencilMaskSeparate
```

#### func (*GlStencilMaskSeparate) API

```go
func (c *GlStencilMaskSeparate) API() gfxapi.API
```

#### func (*GlStencilMaskSeparate) Class

```go
func (*GlStencilMaskSeparate) Class() binary.Class
```

#### func (*GlStencilMaskSeparate) Flags

```go
func (c *GlStencilMaskSeparate) Flags() atom.Flags
```

#### func (*GlStencilMaskSeparate) Mutate

```go
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlStencilMaskSeparate) Replay

```go
func (ϟa *GlStencilMaskSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlStencilMaskSeparate) String

```go
func (c *GlStencilMaskSeparate) String() string
```

#### func (*GlStencilMaskSeparate) TypeID

```go
func (c *GlStencilMaskSeparate) TypeID() atom.TypeID
```

#### type GlStencilOpSeparate

```go
type GlStencilOpSeparate struct {
	binary.Generate
	Face                 FaceMode
	StencilFail          StencilAction
	StencilPassDepthFail StencilAction
	StencilPassDepthPass StencilAction
}
```

//////////////////////////////////////////////////////////////////////////////
GlStencilOpSeparate
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlStencilOpSeparate

```go
func NewGlStencilOpSeparate(
	pFace FaceMode,
	pStencilFail StencilAction,
	pStencilPassDepthFail StencilAction,
	pStencilPassDepthPass StencilAction,
) *GlStencilOpSeparate
```

#### func (*GlStencilOpSeparate) API

```go
func (c *GlStencilOpSeparate) API() gfxapi.API
```

#### func (*GlStencilOpSeparate) Class

```go
func (*GlStencilOpSeparate) Class() binary.Class
```

#### func (*GlStencilOpSeparate) Flags

```go
func (c *GlStencilOpSeparate) Flags() atom.Flags
```

#### func (*GlStencilOpSeparate) Mutate

```go
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlStencilOpSeparate) Replay

```go
func (ϟa *GlStencilOpSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlStencilOpSeparate) String

```go
func (c *GlStencilOpSeparate) String() string
```

#### func (*GlStencilOpSeparate) TypeID

```go
func (c *GlStencilOpSeparate) TypeID() atom.TypeID
```

#### type GlTexImage2D

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlTexImage2D
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTexImage2D

```go
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
) *GlTexImage2D
```

#### func (*GlTexImage2D) API

```go
func (c *GlTexImage2D) API() gfxapi.API
```

#### func (*GlTexImage2D) Class

```go
func (*GlTexImage2D) Class() binary.Class
```

#### func (*GlTexImage2D) Flags

```go
func (c *GlTexImage2D) Flags() atom.Flags
```

#### func (*GlTexImage2D) Mutate

```go
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTexImage2D) Replay

```go
func (ϟa *GlTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTexImage2D) String

```go
func (c *GlTexImage2D) String() string
```

#### func (*GlTexImage2D) TypeID

```go
func (c *GlTexImage2D) TypeID() atom.TypeID
```

#### type GlTexParameterf

```go
type GlTexParameterf struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Value     float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlTexParameterf
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTexParameterf

```go
func NewGlTexParameterf(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValue float32,
) *GlTexParameterf
```

#### func (*GlTexParameterf) API

```go
func (c *GlTexParameterf) API() gfxapi.API
```

#### func (*GlTexParameterf) Class

```go
func (*GlTexParameterf) Class() binary.Class
```

#### func (*GlTexParameterf) Flags

```go
func (c *GlTexParameterf) Flags() atom.Flags
```

#### func (*GlTexParameterf) Mutate

```go
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTexParameterf) Replay

```go
func (ϟa *GlTexParameterf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTexParameterf) String

```go
func (c *GlTexParameterf) String() string
```

#### func (*GlTexParameterf) TypeID

```go
func (c *GlTexParameterf) TypeID() atom.TypeID
```

#### type GlTexParameteri

```go
type GlTexParameteri struct {
	binary.Generate
	Target    TextureTarget
	Parameter TextureParameter
	Value     int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlTexParameteri
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTexParameteri

```go
func NewGlTexParameteri(
	pTarget TextureTarget,
	pParameter TextureParameter,
	pValue int32,
) *GlTexParameteri
```

#### func (*GlTexParameteri) API

```go
func (c *GlTexParameteri) API() gfxapi.API
```

#### func (*GlTexParameteri) Class

```go
func (*GlTexParameteri) Class() binary.Class
```

#### func (*GlTexParameteri) Flags

```go
func (c *GlTexParameteri) Flags() atom.Flags
```

#### func (*GlTexParameteri) Mutate

```go
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTexParameteri) Replay

```go
func (ϟa *GlTexParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTexParameteri) String

```go
func (c *GlTexParameteri) String() string
```

#### func (*GlTexParameteri) TypeID

```go
func (c *GlTexParameteri) TypeID() atom.TypeID
```

#### type GlTexStorage1DEXT

```go
type GlTexStorage1DEXT struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlTexStorage1DEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTexStorage1DEXT

```go
func NewGlTexStorage1DEXT(
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
) *GlTexStorage1DEXT
```

#### func (*GlTexStorage1DEXT) API

```go
func (c *GlTexStorage1DEXT) API() gfxapi.API
```

#### func (*GlTexStorage1DEXT) Class

```go
func (*GlTexStorage1DEXT) Class() binary.Class
```

#### func (*GlTexStorage1DEXT) Flags

```go
func (c *GlTexStorage1DEXT) Flags() atom.Flags
```

#### func (*GlTexStorage1DEXT) Mutate

```go
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTexStorage1DEXT) Replay

```go
func (ϟa *GlTexStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTexStorage1DEXT) String

```go
func (c *GlTexStorage1DEXT) String() string
```

#### func (*GlTexStorage1DEXT) TypeID

```go
func (c *GlTexStorage1DEXT) TypeID() atom.TypeID
```

#### type GlTexStorage2DEXT

```go
type GlTexStorage2DEXT struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
	Height int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlTexStorage2DEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTexStorage2DEXT

```go
func NewGlTexStorage2DEXT(
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
) *GlTexStorage2DEXT
```

#### func (*GlTexStorage2DEXT) API

```go
func (c *GlTexStorage2DEXT) API() gfxapi.API
```

#### func (*GlTexStorage2DEXT) Class

```go
func (*GlTexStorage2DEXT) Class() binary.Class
```

#### func (*GlTexStorage2DEXT) Flags

```go
func (c *GlTexStorage2DEXT) Flags() atom.Flags
```

#### func (*GlTexStorage2DEXT) Mutate

```go
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTexStorage2DEXT) Replay

```go
func (ϟa *GlTexStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTexStorage2DEXT) String

```go
func (c *GlTexStorage2DEXT) String() string
```

#### func (*GlTexStorage2DEXT) TypeID

```go
func (c *GlTexStorage2DEXT) TypeID() atom.TypeID
```

#### type GlTexStorage3DEXT

```go
type GlTexStorage3DEXT struct {
	binary.Generate
	Target TextureTarget
	Levels int32
	Format TexelFormat
	Width  int32
	Height int32
	Depth  int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlTexStorage3DEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTexStorage3DEXT

```go
func NewGlTexStorage3DEXT(
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
	pDepth int32,
) *GlTexStorage3DEXT
```

#### func (*GlTexStorage3DEXT) API

```go
func (c *GlTexStorage3DEXT) API() gfxapi.API
```

#### func (*GlTexStorage3DEXT) Class

```go
func (*GlTexStorage3DEXT) Class() binary.Class
```

#### func (*GlTexStorage3DEXT) Flags

```go
func (c *GlTexStorage3DEXT) Flags() atom.Flags
```

#### func (*GlTexStorage3DEXT) Mutate

```go
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTexStorage3DEXT) Replay

```go
func (ϟa *GlTexStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTexStorage3DEXT) String

```go
func (c *GlTexStorage3DEXT) String() string
```

#### func (*GlTexStorage3DEXT) TypeID

```go
func (c *GlTexStorage3DEXT) TypeID() atom.TypeID
```

#### type GlTexSubImage2D

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlTexSubImage2D
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTexSubImage2D

```go
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
) *GlTexSubImage2D
```

#### func (*GlTexSubImage2D) API

```go
func (c *GlTexSubImage2D) API() gfxapi.API
```

#### func (*GlTexSubImage2D) Class

```go
func (*GlTexSubImage2D) Class() binary.Class
```

#### func (*GlTexSubImage2D) Flags

```go
func (c *GlTexSubImage2D) Flags() atom.Flags
```

#### func (*GlTexSubImage2D) Mutate

```go
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTexSubImage2D) Replay

```go
func (ϟa *GlTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTexSubImage2D) String

```go
func (c *GlTexSubImage2D) String() string
```

#### func (*GlTexSubImage2D) TypeID

```go
func (c *GlTexSubImage2D) TypeID() atom.TypeID
```

#### type GlTextureStorage1DEXT

```go
type GlTextureStorage1DEXT struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlTextureStorage1DEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTextureStorage1DEXT

```go
func NewGlTextureStorage1DEXT(
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
) *GlTextureStorage1DEXT
```

#### func (*GlTextureStorage1DEXT) API

```go
func (c *GlTextureStorage1DEXT) API() gfxapi.API
```

#### func (*GlTextureStorage1DEXT) Class

```go
func (*GlTextureStorage1DEXT) Class() binary.Class
```

#### func (*GlTextureStorage1DEXT) Flags

```go
func (c *GlTextureStorage1DEXT) Flags() atom.Flags
```

#### func (*GlTextureStorage1DEXT) Mutate

```go
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTextureStorage1DEXT) Replay

```go
func (ϟa *GlTextureStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTextureStorage1DEXT) String

```go
func (c *GlTextureStorage1DEXT) String() string
```

#### func (*GlTextureStorage1DEXT) TypeID

```go
func (c *GlTextureStorage1DEXT) TypeID() atom.TypeID
```

#### type GlTextureStorage2DEXT

```go
type GlTextureStorage2DEXT struct {
	binary.Generate
	Texture TextureId
	Target  TextureTarget
	Levels  int32
	Format  TexelFormat
	Width   int32
	Height  int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlTextureStorage2DEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTextureStorage2DEXT

```go
func NewGlTextureStorage2DEXT(
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
) *GlTextureStorage2DEXT
```

#### func (*GlTextureStorage2DEXT) API

```go
func (c *GlTextureStorage2DEXT) API() gfxapi.API
```

#### func (*GlTextureStorage2DEXT) Class

```go
func (*GlTextureStorage2DEXT) Class() binary.Class
```

#### func (*GlTextureStorage2DEXT) Flags

```go
func (c *GlTextureStorage2DEXT) Flags() atom.Flags
```

#### func (*GlTextureStorage2DEXT) Mutate

```go
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTextureStorage2DEXT) Replay

```go
func (ϟa *GlTextureStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTextureStorage2DEXT) String

```go
func (c *GlTextureStorage2DEXT) String() string
```

#### func (*GlTextureStorage2DEXT) TypeID

```go
func (c *GlTextureStorage2DEXT) TypeID() atom.TypeID
```

#### type GlTextureStorage3DEXT

```go
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
```

//////////////////////////////////////////////////////////////////////////////
GlTextureStorage3DEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlTextureStorage3DEXT

```go
func NewGlTextureStorage3DEXT(
	pTexture TextureId,
	pTarget TextureTarget,
	pLevels int32,
	pFormat TexelFormat,
	pWidth int32,
	pHeight int32,
	pDepth int32,
) *GlTextureStorage3DEXT
```

#### func (*GlTextureStorage3DEXT) API

```go
func (c *GlTextureStorage3DEXT) API() gfxapi.API
```

#### func (*GlTextureStorage3DEXT) Class

```go
func (*GlTextureStorage3DEXT) Class() binary.Class
```

#### func (*GlTextureStorage3DEXT) Flags

```go
func (c *GlTextureStorage3DEXT) Flags() atom.Flags
```

#### func (*GlTextureStorage3DEXT) Mutate

```go
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlTextureStorage3DEXT) Replay

```go
func (ϟa *GlTextureStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlTextureStorage3DEXT) String

```go
func (c *GlTextureStorage3DEXT) String() string
```

#### func (*GlTextureStorage3DEXT) TypeID

```go
func (c *GlTextureStorage3DEXT) TypeID() atom.TypeID
```

#### type GlUniform1f

```go
type GlUniform1f struct {
	binary.Generate
	Location UniformLocation
	Value    float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform1f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform1f

```go
func NewGlUniform1f(
	pLocation UniformLocation,
	pValue float32,
) *GlUniform1f
```

#### func (*GlUniform1f) API

```go
func (c *GlUniform1f) API() gfxapi.API
```

#### func (*GlUniform1f) Class

```go
func (*GlUniform1f) Class() binary.Class
```

#### func (*GlUniform1f) Flags

```go
func (c *GlUniform1f) Flags() atom.Flags
```

#### func (*GlUniform1f) Mutate

```go
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform1f) Replay

```go
func (ϟa *GlUniform1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform1f) String

```go
func (c *GlUniform1f) String() string
```

#### func (*GlUniform1f) TypeID

```go
func (c *GlUniform1f) TypeID() atom.TypeID
```

#### type GlUniform1fv

```go
type GlUniform1fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform1fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform1fv

```go
func NewGlUniform1fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform1fv
```

#### func (*GlUniform1fv) API

```go
func (c *GlUniform1fv) API() gfxapi.API
```

#### func (*GlUniform1fv) Class

```go
func (*GlUniform1fv) Class() binary.Class
```

#### func (*GlUniform1fv) Flags

```go
func (c *GlUniform1fv) Flags() atom.Flags
```

#### func (*GlUniform1fv) Mutate

```go
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform1fv) Replay

```go
func (ϟa *GlUniform1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform1fv) String

```go
func (c *GlUniform1fv) String() string
```

#### func (*GlUniform1fv) TypeID

```go
func (c *GlUniform1fv) TypeID() atom.TypeID
```

#### type GlUniform1i

```go
type GlUniform1i struct {
	binary.Generate
	Location UniformLocation
	Value    int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform1i
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform1i

```go
func NewGlUniform1i(
	pLocation UniformLocation,
	pValue int32,
) *GlUniform1i
```

#### func (*GlUniform1i) API

```go
func (c *GlUniform1i) API() gfxapi.API
```

#### func (*GlUniform1i) Class

```go
func (*GlUniform1i) Class() binary.Class
```

#### func (*GlUniform1i) Flags

```go
func (c *GlUniform1i) Flags() atom.Flags
```

#### func (*GlUniform1i) Mutate

```go
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform1i) Replay

```go
func (ϟa *GlUniform1i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform1i) String

```go
func (c *GlUniform1i) String() string
```

#### func (*GlUniform1i) TypeID

```go
func (c *GlUniform1i) TypeID() atom.TypeID
```

#### type GlUniform1iv

```go
type GlUniform1iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform1iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform1iv

```go
func NewGlUniform1iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform1iv
```

#### func (*GlUniform1iv) API

```go
func (c *GlUniform1iv) API() gfxapi.API
```

#### func (*GlUniform1iv) Class

```go
func (*GlUniform1iv) Class() binary.Class
```

#### func (*GlUniform1iv) Flags

```go
func (c *GlUniform1iv) Flags() atom.Flags
```

#### func (*GlUniform1iv) Mutate

```go
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform1iv) Replay

```go
func (ϟa *GlUniform1iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform1iv) String

```go
func (c *GlUniform1iv) String() string
```

#### func (*GlUniform1iv) TypeID

```go
func (c *GlUniform1iv) TypeID() atom.TypeID
```

#### type GlUniform2f

```go
type GlUniform2f struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform2f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform2f

```go
func NewGlUniform2f(
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
) *GlUniform2f
```

#### func (*GlUniform2f) API

```go
func (c *GlUniform2f) API() gfxapi.API
```

#### func (*GlUniform2f) Class

```go
func (*GlUniform2f) Class() binary.Class
```

#### func (*GlUniform2f) Flags

```go
func (c *GlUniform2f) Flags() atom.Flags
```

#### func (*GlUniform2f) Mutate

```go
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform2f) Replay

```go
func (ϟa *GlUniform2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform2f) String

```go
func (c *GlUniform2f) String() string
```

#### func (*GlUniform2f) TypeID

```go
func (c *GlUniform2f) TypeID() atom.TypeID
```

#### type GlUniform2fv

```go
type GlUniform2fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform2fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform2fv

```go
func NewGlUniform2fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform2fv
```

#### func (*GlUniform2fv) API

```go
func (c *GlUniform2fv) API() gfxapi.API
```

#### func (*GlUniform2fv) Class

```go
func (*GlUniform2fv) Class() binary.Class
```

#### func (*GlUniform2fv) Flags

```go
func (c *GlUniform2fv) Flags() atom.Flags
```

#### func (*GlUniform2fv) Mutate

```go
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform2fv) Replay

```go
func (ϟa *GlUniform2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform2fv) String

```go
func (c *GlUniform2fv) String() string
```

#### func (*GlUniform2fv) TypeID

```go
func (c *GlUniform2fv) TypeID() atom.TypeID
```

#### type GlUniform2i

```go
type GlUniform2i struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform2i
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform2i

```go
func NewGlUniform2i(
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
) *GlUniform2i
```

#### func (*GlUniform2i) API

```go
func (c *GlUniform2i) API() gfxapi.API
```

#### func (*GlUniform2i) Class

```go
func (*GlUniform2i) Class() binary.Class
```

#### func (*GlUniform2i) Flags

```go
func (c *GlUniform2i) Flags() atom.Flags
```

#### func (*GlUniform2i) Mutate

```go
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform2i) Replay

```go
func (ϟa *GlUniform2i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform2i) String

```go
func (c *GlUniform2i) String() string
```

#### func (*GlUniform2i) TypeID

```go
func (c *GlUniform2i) TypeID() atom.TypeID
```

#### type GlUniform2iv

```go
type GlUniform2iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform2iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform2iv

```go
func NewGlUniform2iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform2iv
```

#### func (*GlUniform2iv) API

```go
func (c *GlUniform2iv) API() gfxapi.API
```

#### func (*GlUniform2iv) Class

```go
func (*GlUniform2iv) Class() binary.Class
```

#### func (*GlUniform2iv) Flags

```go
func (c *GlUniform2iv) Flags() atom.Flags
```

#### func (*GlUniform2iv) Mutate

```go
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform2iv) Replay

```go
func (ϟa *GlUniform2iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform2iv) String

```go
func (c *GlUniform2iv) String() string
```

#### func (*GlUniform2iv) TypeID

```go
func (c *GlUniform2iv) TypeID() atom.TypeID
```

#### type GlUniform3f

```go
type GlUniform3f struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
	Value2   float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform3f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform3f

```go
func NewGlUniform3f(
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
) *GlUniform3f
```

#### func (*GlUniform3f) API

```go
func (c *GlUniform3f) API() gfxapi.API
```

#### func (*GlUniform3f) Class

```go
func (*GlUniform3f) Class() binary.Class
```

#### func (*GlUniform3f) Flags

```go
func (c *GlUniform3f) Flags() atom.Flags
```

#### func (*GlUniform3f) Mutate

```go
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform3f) Replay

```go
func (ϟa *GlUniform3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform3f) String

```go
func (c *GlUniform3f) String() string
```

#### func (*GlUniform3f) TypeID

```go
func (c *GlUniform3f) TypeID() atom.TypeID
```

#### type GlUniform3fv

```go
type GlUniform3fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform3fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform3fv

```go
func NewGlUniform3fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform3fv
```

#### func (*GlUniform3fv) API

```go
func (c *GlUniform3fv) API() gfxapi.API
```

#### func (*GlUniform3fv) Class

```go
func (*GlUniform3fv) Class() binary.Class
```

#### func (*GlUniform3fv) Flags

```go
func (c *GlUniform3fv) Flags() atom.Flags
```

#### func (*GlUniform3fv) Mutate

```go
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform3fv) Replay

```go
func (ϟa *GlUniform3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform3fv) String

```go
func (c *GlUniform3fv) String() string
```

#### func (*GlUniform3fv) TypeID

```go
func (c *GlUniform3fv) TypeID() atom.TypeID
```

#### type GlUniform3i

```go
type GlUniform3i struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
	Value2   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform3i
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform3i

```go
func NewGlUniform3i(
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
	pValue2 int32,
) *GlUniform3i
```

#### func (*GlUniform3i) API

```go
func (c *GlUniform3i) API() gfxapi.API
```

#### func (*GlUniform3i) Class

```go
func (*GlUniform3i) Class() binary.Class
```

#### func (*GlUniform3i) Flags

```go
func (c *GlUniform3i) Flags() atom.Flags
```

#### func (*GlUniform3i) Mutate

```go
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform3i) Replay

```go
func (ϟa *GlUniform3i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform3i) String

```go
func (c *GlUniform3i) String() string
```

#### func (*GlUniform3i) TypeID

```go
func (c *GlUniform3i) TypeID() atom.TypeID
```

#### type GlUniform3iv

```go
type GlUniform3iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform3iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform3iv

```go
func NewGlUniform3iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform3iv
```

#### func (*GlUniform3iv) API

```go
func (c *GlUniform3iv) API() gfxapi.API
```

#### func (*GlUniform3iv) Class

```go
func (*GlUniform3iv) Class() binary.Class
```

#### func (*GlUniform3iv) Flags

```go
func (c *GlUniform3iv) Flags() atom.Flags
```

#### func (*GlUniform3iv) Mutate

```go
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform3iv) Replay

```go
func (ϟa *GlUniform3iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform3iv) String

```go
func (c *GlUniform3iv) String() string
```

#### func (*GlUniform3iv) TypeID

```go
func (c *GlUniform3iv) TypeID() atom.TypeID
```

#### type GlUniform4f

```go
type GlUniform4f struct {
	binary.Generate
	Location UniformLocation
	Value0   float32
	Value1   float32
	Value2   float32
	Value3   float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform4f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform4f

```go
func NewGlUniform4f(
	pLocation UniformLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
	pValue3 float32,
) *GlUniform4f
```

#### func (*GlUniform4f) API

```go
func (c *GlUniform4f) API() gfxapi.API
```

#### func (*GlUniform4f) Class

```go
func (*GlUniform4f) Class() binary.Class
```

#### func (*GlUniform4f) Flags

```go
func (c *GlUniform4f) Flags() atom.Flags
```

#### func (*GlUniform4f) Mutate

```go
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform4f) Replay

```go
func (ϟa *GlUniform4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform4f) String

```go
func (c *GlUniform4f) String() string
```

#### func (*GlUniform4f) TypeID

```go
func (c *GlUniform4f) TypeID() atom.TypeID
```

#### type GlUniform4fv

```go
type GlUniform4fv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform4fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform4fv

```go
func NewGlUniform4fv(
	pLocation UniformLocation,
	pCount int32,
	pValue F32Array,
) *GlUniform4fv
```

#### func (*GlUniform4fv) API

```go
func (c *GlUniform4fv) API() gfxapi.API
```

#### func (*GlUniform4fv) Class

```go
func (*GlUniform4fv) Class() binary.Class
```

#### func (*GlUniform4fv) Flags

```go
func (c *GlUniform4fv) Flags() atom.Flags
```

#### func (*GlUniform4fv) Mutate

```go
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform4fv) Replay

```go
func (ϟa *GlUniform4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform4fv) String

```go
func (c *GlUniform4fv) String() string
```

#### func (*GlUniform4fv) TypeID

```go
func (c *GlUniform4fv) TypeID() atom.TypeID
```

#### type GlUniform4i

```go
type GlUniform4i struct {
	binary.Generate
	Location UniformLocation
	Value0   int32
	Value1   int32
	Value2   int32
	Value3   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform4i
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform4i

```go
func NewGlUniform4i(
	pLocation UniformLocation,
	pValue0 int32,
	pValue1 int32,
	pValue2 int32,
	pValue3 int32,
) *GlUniform4i
```

#### func (*GlUniform4i) API

```go
func (c *GlUniform4i) API() gfxapi.API
```

#### func (*GlUniform4i) Class

```go
func (*GlUniform4i) Class() binary.Class
```

#### func (*GlUniform4i) Flags

```go
func (c *GlUniform4i) Flags() atom.Flags
```

#### func (*GlUniform4i) Mutate

```go
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform4i) Replay

```go
func (ϟa *GlUniform4i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform4i) String

```go
func (c *GlUniform4i) String() string
```

#### func (*GlUniform4i) TypeID

```go
func (c *GlUniform4i) TypeID() atom.TypeID
```

#### type GlUniform4iv

```go
type GlUniform4iv struct {
	binary.Generate
	Location UniformLocation
	Count    int32
	Value    S32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform4iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform4iv

```go
func NewGlUniform4iv(
	pLocation UniformLocation,
	pCount int32,
	pValue S32Array,
) *GlUniform4iv
```

#### func (*GlUniform4iv) API

```go
func (c *GlUniform4iv) API() gfxapi.API
```

#### func (*GlUniform4iv) Class

```go
func (*GlUniform4iv) Class() binary.Class
```

#### func (*GlUniform4iv) Flags

```go
func (c *GlUniform4iv) Flags() atom.Flags
```

#### func (*GlUniform4iv) Mutate

```go
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniform4iv) Replay

```go
func (ϟa *GlUniform4iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniform4iv) String

```go
func (c *GlUniform4iv) String() string
```

#### func (*GlUniform4iv) TypeID

```go
func (c *GlUniform4iv) TypeID() atom.TypeID
```

#### type GlUniformMatrix2fv

```go
type GlUniformMatrix2fv struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniformMatrix2fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniformMatrix2fv

```go
func NewGlUniformMatrix2fv(
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix2fv
```

#### func (*GlUniformMatrix2fv) API

```go
func (c *GlUniformMatrix2fv) API() gfxapi.API
```

#### func (*GlUniformMatrix2fv) Class

```go
func (*GlUniformMatrix2fv) Class() binary.Class
```

#### func (*GlUniformMatrix2fv) Flags

```go
func (c *GlUniformMatrix2fv) Flags() atom.Flags
```

#### func (*GlUniformMatrix2fv) Mutate

```go
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniformMatrix2fv) Replay

```go
func (ϟa *GlUniformMatrix2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniformMatrix2fv) String

```go
func (c *GlUniformMatrix2fv) String() string
```

#### func (*GlUniformMatrix2fv) TypeID

```go
func (c *GlUniformMatrix2fv) TypeID() atom.TypeID
```

#### type GlUniformMatrix3fv

```go
type GlUniformMatrix3fv struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniformMatrix3fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniformMatrix3fv

```go
func NewGlUniformMatrix3fv(
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix3fv
```

#### func (*GlUniformMatrix3fv) API

```go
func (c *GlUniformMatrix3fv) API() gfxapi.API
```

#### func (*GlUniformMatrix3fv) Class

```go
func (*GlUniformMatrix3fv) Class() binary.Class
```

#### func (*GlUniformMatrix3fv) Flags

```go
func (c *GlUniformMatrix3fv) Flags() atom.Flags
```

#### func (*GlUniformMatrix3fv) Mutate

```go
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniformMatrix3fv) Replay

```go
func (ϟa *GlUniformMatrix3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniformMatrix3fv) String

```go
func (c *GlUniformMatrix3fv) String() string
```

#### func (*GlUniformMatrix3fv) TypeID

```go
func (c *GlUniformMatrix3fv) TypeID() atom.TypeID
```

#### type GlUniformMatrix4fv

```go
type GlUniformMatrix4fv struct {
	binary.Generate
	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniformMatrix4fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniformMatrix4fv

```go
func NewGlUniformMatrix4fv(
	pLocation UniformLocation,
	pCount int32,
	pTranspose bool,
	pValues F32Array,
) *GlUniformMatrix4fv
```

#### func (*GlUniformMatrix4fv) API

```go
func (c *GlUniformMatrix4fv) API() gfxapi.API
```

#### func (*GlUniformMatrix4fv) Class

```go
func (*GlUniformMatrix4fv) Class() binary.Class
```

#### func (*GlUniformMatrix4fv) Flags

```go
func (c *GlUniformMatrix4fv) Flags() atom.Flags
```

#### func (*GlUniformMatrix4fv) Mutate

```go
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUniformMatrix4fv) Replay

```go
func (ϟa *GlUniformMatrix4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUniformMatrix4fv) String

```go
func (c *GlUniformMatrix4fv) String() string
```

#### func (*GlUniformMatrix4fv) TypeID

```go
func (c *GlUniformMatrix4fv) TypeID() atom.TypeID
```

#### type GlUnmapBuffer

```go
type GlUnmapBuffer struct {
	binary.Generate
	Target BufferTarget
}
```

//////////////////////////////////////////////////////////////////////////////
GlUnmapBuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUnmapBuffer

```go
func NewGlUnmapBuffer(
	pTarget BufferTarget,
) *GlUnmapBuffer
```

#### func (*GlUnmapBuffer) API

```go
func (c *GlUnmapBuffer) API() gfxapi.API
```

#### func (*GlUnmapBuffer) Class

```go
func (*GlUnmapBuffer) Class() binary.Class
```

#### func (*GlUnmapBuffer) Flags

```go
func (c *GlUnmapBuffer) Flags() atom.Flags
```

#### func (*GlUnmapBuffer) Mutate

```go
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUnmapBuffer) Replay

```go
func (ϟa *GlUnmapBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUnmapBuffer) String

```go
func (c *GlUnmapBuffer) String() string
```

#### func (*GlUnmapBuffer) TypeID

```go
func (c *GlUnmapBuffer) TypeID() atom.TypeID
```

#### type GlUseProgram

```go
type GlUseProgram struct {
	binary.Generate
	Program ProgramId
}
```

//////////////////////////////////////////////////////////////////////////////
GlUseProgram
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUseProgram

```go
func NewGlUseProgram(
	pProgram ProgramId,
) *GlUseProgram
```

#### func (*GlUseProgram) API

```go
func (c *GlUseProgram) API() gfxapi.API
```

#### func (*GlUseProgram) Class

```go
func (*GlUseProgram) Class() binary.Class
```

#### func (*GlUseProgram) Flags

```go
func (c *GlUseProgram) Flags() atom.Flags
```

#### func (*GlUseProgram) Mutate

```go
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlUseProgram) Replay

```go
func (ϟa *GlUseProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlUseProgram) String

```go
func (c *GlUseProgram) String() string
```

#### func (*GlUseProgram) TypeID

```go
func (c *GlUseProgram) TypeID() atom.TypeID
```

#### type GlValidateProgram

```go
type GlValidateProgram struct {
	binary.Generate
	Program ProgramId
}
```

//////////////////////////////////////////////////////////////////////////////
GlValidateProgram
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlValidateProgram

```go
func NewGlValidateProgram(
	pProgram ProgramId,
) *GlValidateProgram
```

#### func (*GlValidateProgram) API

```go
func (c *GlValidateProgram) API() gfxapi.API
```

#### func (*GlValidateProgram) Class

```go
func (*GlValidateProgram) Class() binary.Class
```

#### func (*GlValidateProgram) Flags

```go
func (c *GlValidateProgram) Flags() atom.Flags
```

#### func (*GlValidateProgram) Mutate

```go
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlValidateProgram) Replay

```go
func (ϟa *GlValidateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlValidateProgram) String

```go
func (c *GlValidateProgram) String() string
```

#### func (*GlValidateProgram) TypeID

```go
func (c *GlValidateProgram) TypeID() atom.TypeID
```

#### type GlVertexAttrib1f

```go
type GlVertexAttrib1f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib1f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib1f

```go
func NewGlVertexAttrib1f(
	pLocation AttributeLocation,
	pValue0 float32,
) *GlVertexAttrib1f
```

#### func (*GlVertexAttrib1f) API

```go
func (c *GlVertexAttrib1f) API() gfxapi.API
```

#### func (*GlVertexAttrib1f) Class

```go
func (*GlVertexAttrib1f) Class() binary.Class
```

#### func (*GlVertexAttrib1f) Flags

```go
func (c *GlVertexAttrib1f) Flags() atom.Flags
```

#### func (*GlVertexAttrib1f) Mutate

```go
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib1f) Replay

```go
func (ϟa *GlVertexAttrib1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib1f) String

```go
func (c *GlVertexAttrib1f) String() string
```

#### func (*GlVertexAttrib1f) TypeID

```go
func (c *GlVertexAttrib1f) TypeID() atom.TypeID
```

#### type GlVertexAttrib1fv

```go
type GlVertexAttrib1fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib1fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib1fv

```go
func NewGlVertexAttrib1fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib1fv
```

#### func (*GlVertexAttrib1fv) API

```go
func (c *GlVertexAttrib1fv) API() gfxapi.API
```

#### func (*GlVertexAttrib1fv) Class

```go
func (*GlVertexAttrib1fv) Class() binary.Class
```

#### func (*GlVertexAttrib1fv) Flags

```go
func (c *GlVertexAttrib1fv) Flags() atom.Flags
```

#### func (*GlVertexAttrib1fv) Mutate

```go
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib1fv) Replay

```go
func (ϟa *GlVertexAttrib1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib1fv) String

```go
func (c *GlVertexAttrib1fv) String() string
```

#### func (*GlVertexAttrib1fv) TypeID

```go
func (c *GlVertexAttrib1fv) TypeID() atom.TypeID
```

#### type GlVertexAttrib2f

```go
type GlVertexAttrib2f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib2f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib2f

```go
func NewGlVertexAttrib2f(
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
) *GlVertexAttrib2f
```

#### func (*GlVertexAttrib2f) API

```go
func (c *GlVertexAttrib2f) API() gfxapi.API
```

#### func (*GlVertexAttrib2f) Class

```go
func (*GlVertexAttrib2f) Class() binary.Class
```

#### func (*GlVertexAttrib2f) Flags

```go
func (c *GlVertexAttrib2f) Flags() atom.Flags
```

#### func (*GlVertexAttrib2f) Mutate

```go
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib2f) Replay

```go
func (ϟa *GlVertexAttrib2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib2f) String

```go
func (c *GlVertexAttrib2f) String() string
```

#### func (*GlVertexAttrib2f) TypeID

```go
func (c *GlVertexAttrib2f) TypeID() atom.TypeID
```

#### type GlVertexAttrib2fv

```go
type GlVertexAttrib2fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib2fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib2fv

```go
func NewGlVertexAttrib2fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib2fv
```

#### func (*GlVertexAttrib2fv) API

```go
func (c *GlVertexAttrib2fv) API() gfxapi.API
```

#### func (*GlVertexAttrib2fv) Class

```go
func (*GlVertexAttrib2fv) Class() binary.Class
```

#### func (*GlVertexAttrib2fv) Flags

```go
func (c *GlVertexAttrib2fv) Flags() atom.Flags
```

#### func (*GlVertexAttrib2fv) Mutate

```go
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib2fv) Replay

```go
func (ϟa *GlVertexAttrib2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib2fv) String

```go
func (c *GlVertexAttrib2fv) String() string
```

#### func (*GlVertexAttrib2fv) TypeID

```go
func (c *GlVertexAttrib2fv) TypeID() atom.TypeID
```

#### type GlVertexAttrib3f

```go
type GlVertexAttrib3f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
	Value2   float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib3f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib3f

```go
func NewGlVertexAttrib3f(
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
) *GlVertexAttrib3f
```

#### func (*GlVertexAttrib3f) API

```go
func (c *GlVertexAttrib3f) API() gfxapi.API
```

#### func (*GlVertexAttrib3f) Class

```go
func (*GlVertexAttrib3f) Class() binary.Class
```

#### func (*GlVertexAttrib3f) Flags

```go
func (c *GlVertexAttrib3f) Flags() atom.Flags
```

#### func (*GlVertexAttrib3f) Mutate

```go
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib3f) Replay

```go
func (ϟa *GlVertexAttrib3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib3f) String

```go
func (c *GlVertexAttrib3f) String() string
```

#### func (*GlVertexAttrib3f) TypeID

```go
func (c *GlVertexAttrib3f) TypeID() atom.TypeID
```

#### type GlVertexAttrib3fv

```go
type GlVertexAttrib3fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib3fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib3fv

```go
func NewGlVertexAttrib3fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib3fv
```

#### func (*GlVertexAttrib3fv) API

```go
func (c *GlVertexAttrib3fv) API() gfxapi.API
```

#### func (*GlVertexAttrib3fv) Class

```go
func (*GlVertexAttrib3fv) Class() binary.Class
```

#### func (*GlVertexAttrib3fv) Flags

```go
func (c *GlVertexAttrib3fv) Flags() atom.Flags
```

#### func (*GlVertexAttrib3fv) Mutate

```go
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib3fv) Replay

```go
func (ϟa *GlVertexAttrib3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib3fv) String

```go
func (c *GlVertexAttrib3fv) String() string
```

#### func (*GlVertexAttrib3fv) TypeID

```go
func (c *GlVertexAttrib3fv) TypeID() atom.TypeID
```

#### type GlVertexAttrib4f

```go
type GlVertexAttrib4f struct {
	binary.Generate
	Location AttributeLocation
	Value0   float32
	Value1   float32
	Value2   float32
	Value3   float32
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib4f
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib4f

```go
func NewGlVertexAttrib4f(
	pLocation AttributeLocation,
	pValue0 float32,
	pValue1 float32,
	pValue2 float32,
	pValue3 float32,
) *GlVertexAttrib4f
```

#### func (*GlVertexAttrib4f) API

```go
func (c *GlVertexAttrib4f) API() gfxapi.API
```

#### func (*GlVertexAttrib4f) Class

```go
func (*GlVertexAttrib4f) Class() binary.Class
```

#### func (*GlVertexAttrib4f) Flags

```go
func (c *GlVertexAttrib4f) Flags() atom.Flags
```

#### func (*GlVertexAttrib4f) Mutate

```go
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib4f) Replay

```go
func (ϟa *GlVertexAttrib4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib4f) String

```go
func (c *GlVertexAttrib4f) String() string
```

#### func (*GlVertexAttrib4f) TypeID

```go
func (c *GlVertexAttrib4f) TypeID() atom.TypeID
```

#### type GlVertexAttrib4fv

```go
type GlVertexAttrib4fv struct {
	binary.Generate
	Location AttributeLocation
	Value    F32Array
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib4fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib4fv

```go
func NewGlVertexAttrib4fv(
	pLocation AttributeLocation,
	pValue F32Array,
) *GlVertexAttrib4fv
```

#### func (*GlVertexAttrib4fv) API

```go
func (c *GlVertexAttrib4fv) API() gfxapi.API
```

#### func (*GlVertexAttrib4fv) Class

```go
func (*GlVertexAttrib4fv) Class() binary.Class
```

#### func (*GlVertexAttrib4fv) Flags

```go
func (c *GlVertexAttrib4fv) Flags() atom.Flags
```

#### func (*GlVertexAttrib4fv) Mutate

```go
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttrib4fv) Replay

```go
func (ϟa *GlVertexAttrib4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttrib4fv) String

```go
func (c *GlVertexAttrib4fv) String() string
```

#### func (*GlVertexAttrib4fv) TypeID

```go
func (c *GlVertexAttrib4fv) TypeID() atom.TypeID
```

#### type GlVertexAttribPointer

```go
type GlVertexAttribPointer struct {
	binary.Generate
	Location   AttributeLocation
	Size       int32
	Type       VertexAttribType
	Normalized bool
	Stride     int32
	Data       VertexPointer
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttribPointer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttribPointer

```go
func NewGlVertexAttribPointer(
	pLocation AttributeLocation,
	pSize int32,
	pType VertexAttribType,
	pNormalized bool,
	pStride int32,
	pData VertexPointer,
) *GlVertexAttribPointer
```

#### func (*GlVertexAttribPointer) API

```go
func (c *GlVertexAttribPointer) API() gfxapi.API
```

#### func (*GlVertexAttribPointer) Class

```go
func (*GlVertexAttribPointer) Class() binary.Class
```

#### func (*GlVertexAttribPointer) Flags

```go
func (c *GlVertexAttribPointer) Flags() atom.Flags
```

#### func (*GlVertexAttribPointer) Mutate

```go
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlVertexAttribPointer) Replay

```go
func (ϟa *GlVertexAttribPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlVertexAttribPointer) String

```go
func (c *GlVertexAttribPointer) String() string
```

#### func (*GlVertexAttribPointer) TypeID

```go
func (c *GlVertexAttribPointer) TypeID() atom.TypeID
```

#### type GlViewport

```go
type GlViewport struct {
	binary.Generate
	X      int32
	Y      int32
	Width  int32
	Height int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlViewport
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlViewport

```go
func NewGlViewport(
	pX int32,
	pY int32,
	pWidth int32,
	pHeight int32,
) *GlViewport
```

#### func (*GlViewport) API

```go
func (c *GlViewport) API() gfxapi.API
```

#### func (*GlViewport) Class

```go
func (*GlViewport) Class() binary.Class
```

#### func (*GlViewport) Flags

```go
func (c *GlViewport) Flags() atom.Flags
```

#### func (*GlViewport) Mutate

```go
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlViewport) Replay

```go
func (ϟa *GlViewport) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*GlViewport) String

```go
func (c *GlViewport) String() string
```

#### func (*GlViewport) TypeID

```go
func (c *GlViewport) TypeID() atom.TypeID
```

#### type GlXCreateContext

```go
type GlXCreateContext struct {
	binary.Generate
	Dpy       memory.Pointer
	Vis       memory.Pointer
	ShareList GLXContext
	Direct    bool
	Result    GLXContext
}
```

//////////////////////////////////////////////////////////////////////////////
GlXCreateContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXCreateContext

```go
func NewGlXCreateContext(
	pDpy memory.Pointer,
	pVis memory.Pointer,
	pShareList GLXContext,
	pDirect bool,
	pResult GLXContext,
) *GlXCreateContext
```

#### func (*GlXCreateContext) API

```go
func (c *GlXCreateContext) API() gfxapi.API
```

#### func (*GlXCreateContext) Class

```go
func (*GlXCreateContext) Class() binary.Class
```

#### func (*GlXCreateContext) Flags

```go
func (c *GlXCreateContext) Flags() atom.Flags
```

#### func (*GlXCreateContext) Mutate

```go
func (ϟa *GlXCreateContext) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlXCreateContext) Replay

```go
func (ω *GlXCreateContext) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*GlXCreateContext) String

```go
func (c *GlXCreateContext) String() string
```

#### func (*GlXCreateContext) TypeID

```go
func (c *GlXCreateContext) TypeID() atom.TypeID
```

#### type GlXCreateContext_Postback

```go
type GlXCreateContext_Postback struct {
	Result []byte
}
```


#### func (*GlXCreateContext_Postback) Decode

```go
func (o *GlXCreateContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type GlXCreateNewContext

```go
type GlXCreateNewContext struct {
	binary.Generate
	Display  memory.Pointer
	Fbconfig memory.Pointer
	Type     uint32
	Shared   GLXContext
	Direct   bool
	Result   GLXContext
}
```

//////////////////////////////////////////////////////////////////////////////
GlXCreateNewContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXCreateNewContext

```go
func NewGlXCreateNewContext(
	pDisplay memory.Pointer,
	pFbconfig memory.Pointer,
	pType uint32,
	pShared GLXContext,
	pDirect bool,
	pResult GLXContext,
) *GlXCreateNewContext
```

#### func (*GlXCreateNewContext) API

```go
func (c *GlXCreateNewContext) API() gfxapi.API
```

#### func (*GlXCreateNewContext) Class

```go
func (*GlXCreateNewContext) Class() binary.Class
```

#### func (*GlXCreateNewContext) Flags

```go
func (c *GlXCreateNewContext) Flags() atom.Flags
```

#### func (*GlXCreateNewContext) Mutate

```go
func (ϟa *GlXCreateNewContext) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlXCreateNewContext) Replay

```go
func (ω *GlXCreateNewContext) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*GlXCreateNewContext) String

```go
func (c *GlXCreateNewContext) String() string
```

#### func (*GlXCreateNewContext) TypeID

```go
func (c *GlXCreateNewContext) TypeID() atom.TypeID
```

#### type GlXCreateNewContext_Postback

```go
type GlXCreateNewContext_Postback struct {
	Result []byte
}
```


#### func (*GlXCreateNewContext_Postback) Decode

```go
func (o *GlXCreateNewContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type GlXMakeContextCurrent

```go
type GlXMakeContextCurrent struct {
	binary.Generate
	Display memory.Pointer
	Draw    GLXDrawable
	Read    GLXDrawable
	Ctx     GLXContext
}
```

//////////////////////////////////////////////////////////////////////////////
GlXMakeContextCurrent
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXMakeContextCurrent

```go
func NewGlXMakeContextCurrent(
	pDisplay memory.Pointer,
	pDraw GLXDrawable,
	pRead GLXDrawable,
	pCtx GLXContext,
) *GlXMakeContextCurrent
```

#### func (*GlXMakeContextCurrent) API

```go
func (c *GlXMakeContextCurrent) API() gfxapi.API
```

#### func (*GlXMakeContextCurrent) Class

```go
func (*GlXMakeContextCurrent) Class() binary.Class
```

#### func (*GlXMakeContextCurrent) Flags

```go
func (c *GlXMakeContextCurrent) Flags() atom.Flags
```

#### func (*GlXMakeContextCurrent) Mutate

```go
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlXMakeContextCurrent) Replay

```go
func (ω *GlXMakeContextCurrent) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*GlXMakeContextCurrent) String

```go
func (c *GlXMakeContextCurrent) String() string
```

#### func (*GlXMakeContextCurrent) TypeID

```go
func (c *GlXMakeContextCurrent) TypeID() atom.TypeID
```

#### type GlXSwapBuffers

```go
type GlXSwapBuffers struct {
	binary.Generate
	Display  memory.Pointer
	Drawable GLXDrawable
}
```

//////////////////////////////////////////////////////////////////////////////
GlXSwapBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXSwapBuffers

```go
func NewGlXSwapBuffers(
	pDisplay memory.Pointer,
	pDrawable GLXDrawable,
) *GlXSwapBuffers
```

#### func (*GlXSwapBuffers) API

```go
func (c *GlXSwapBuffers) API() gfxapi.API
```

#### func (*GlXSwapBuffers) Class

```go
func (*GlXSwapBuffers) Class() binary.Class
```

#### func (*GlXSwapBuffers) Flags

```go
func (c *GlXSwapBuffers) Flags() atom.Flags
```

#### func (*GlXSwapBuffers) Mutate

```go
func (ϟa *GlXSwapBuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*GlXSwapBuffers) String

```go
func (c *GlXSwapBuffers) String() string
```

#### func (*GlXSwapBuffers) TypeID

```go
func (c *GlXSwapBuffers) TypeID() atom.TypeID
```

#### type Globals

```go
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
```

//////////////////////////////////////////////////////////////////////////////
Globals
//////////////////////////////////////////////////////////////////////////////

#### func (*Globals) Class

```go
func (*Globals) Class() binary.Class
```

#### func (*Globals) Init

```go
func (g *Globals) Init()
```

#### type HDC

```go
type HDC memory.Pointer
```


#### func (*HDC) Equal

```go
func (c *HDC) Equal(rhs HDC) bool
```

#### func (*HDC) Less

```go
func (c *HDC) Less(rhs HDC) bool
```

#### type HGLRC

```go
type HGLRC memory.Pointer
```


#### func (*HGLRC) Equal

```go
func (c *HGLRC) Equal(rhs HGLRC) bool
```

#### func (*HGLRC) Less

```go
func (c *HGLRC) Less(rhs HGLRC) bool
```

#### type HintMode

```go
type HintMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum HintMode
//////////////////////////////////////////////////////////////////////////////

#### func (HintMode) String

```go
func (v HintMode) String() string
```

#### type HintTarget

```go
type HintTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum HintTarget
//////////////////////////////////////////////////////////////////////////////

#### func (HintTarget) String

```go
func (v HintTarget) String() string
```

#### type Image

```go
type Image struct {
	binary.Generate
	CreatedAt atom.ID
	Width     int32
	Height    int32
	Data      memory.Memory
	Size      uint32
	Format    ImageTexelFormat
}
```

//////////////////////////////////////////////////////////////////////////////
class Image
//////////////////////////////////////////////////////////////////////////////

#### func (*Image) Class

```go
func (*Image) Class() binary.Class
```

#### func (*Image) GetCreatedAt

```go
func (c *Image) GetCreatedAt() atom.ID
```

#### func (*Image) Init

```go
func (c *Image) Init()
```

#### type ImageOES

```go
type ImageOES memory.Pointer
```


#### func (*ImageOES) Equal

```go
func (c *ImageOES) Equal(rhs ImageOES) bool
```

#### func (*ImageOES) Less

```go
func (c *ImageOES) Less(rhs ImageOES) bool
```

#### type ImageTargetRenderbufferStorage

```go
type ImageTargetRenderbufferStorage uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ImageTargetRenderbufferStorage
//////////////////////////////////////////////////////////////////////////////

#### func (ImageTargetRenderbufferStorage) String

```go
func (v ImageTargetRenderbufferStorage) String() string
```

#### type ImageTargetTexture

```go
type ImageTargetTexture uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ImageTargetTexture
//////////////////////////////////////////////////////////////////////////////

#### func (ImageTargetTexture) String

```go
func (v ImageTargetTexture) String() string
```

#### type ImageTargetTexture_OES_EGL_image

```go
type ImageTargetTexture_OES_EGL_image uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ImageTargetTexture_OES_EGL_image
//////////////////////////////////////////////////////////////////////////////

#### func (ImageTargetTexture_OES_EGL_image) String

```go
func (v ImageTargetTexture_OES_EGL_image) String() string
```

#### type ImageTargetTexture_OES_EGL_image_external

```go
type ImageTargetTexture_OES_EGL_image_external uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ImageTargetTexture_OES_EGL_image_external
//////////////////////////////////////////////////////////////////////////////

#### func (ImageTargetTexture_OES_EGL_image_external) String

```go
func (v ImageTargetTexture_OES_EGL_image_external) String() string
```

#### type ImageTexelFormat

```go
type ImageTexelFormat uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ImageTexelFormat
//////////////////////////////////////////////////////////////////////////////

#### func (ImageTexelFormat) String

```go
func (v ImageTexelFormat) String() string
```

#### type Image_CubeMapImageTargetMap

```go
type Image_CubeMapImageTargetMap map[CubeMapImageTarget]Image
```


#### func (Image_CubeMapImageTargetMap) Contains

```go
func (m Image_CubeMapImageTargetMap) Contains(key CubeMapImageTarget) bool
```

#### func (Image_CubeMapImageTargetMap) Delete

```go
func (m Image_CubeMapImageTargetMap) Delete(key CubeMapImageTarget)
```

#### func (Image_CubeMapImageTargetMap) Get

```go
func (m Image_CubeMapImageTargetMap) Get(key CubeMapImageTarget) Image
```

#### func (Image_CubeMapImageTargetMap) Range

```go
func (m Image_CubeMapImageTargetMap) Range() []Image
```

#### type Image_s32Map

```go
type Image_s32Map map[int32]Image
```


#### func (Image_s32Map) Contains

```go
func (m Image_s32Map) Contains(key int32) bool
```

#### func (Image_s32Map) Delete

```go
func (m Image_s32Map) Delete(key int32)
```

#### func (Image_s32Map) Get

```go
func (m Image_s32Map) Get(key int32) Image
```

#### func (Image_s32Map) Range

```go
func (m Image_s32Map) Range() []Image
```

#### type IndicesPointer

```go
type IndicesPointer memory.Pointer
```


#### func (*IndicesPointer) Equal

```go
func (c *IndicesPointer) Equal(rhs IndicesPointer) bool
```

#### func (*IndicesPointer) Less

```go
func (c *IndicesPointer) Less(rhs IndicesPointer) bool
```

#### type IndicesType

```go
type IndicesType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum IndicesType
//////////////////////////////////////////////////////////////////////////////

#### func (IndicesType) String

```go
func (v IndicesType) String() string
```

#### type IntArray

```go
type IntArray []int64
```


#### func (IntArray) Len

```go
func (s IntArray) Len() int
```

#### func (IntArray) Range

```go
func (s IntArray) Range() []int64
```

#### type MapBufferRangeAccess

```go
type MapBufferRangeAccess uint32
```

//////////////////////////////////////////////////////////////////////////////
enum MapBufferRangeAccess
//////////////////////////////////////////////////////////////////////////////

#### func (MapBufferRangeAccess) String

```go
func (v MapBufferRangeAccess) String() string
```

#### type Mat2f

```go
type Mat2f struct {
	binary.Generate
	CreatedAt atom.ID
	Col0      Vec2f
	Col1      Vec2f
}
```

//////////////////////////////////////////////////////////////////////////////
class Mat2f
//////////////////////////////////////////////////////////////////////////////

#### func (*Mat2f) Class

```go
func (*Mat2f) Class() binary.Class
```

#### func (*Mat2f) GetCreatedAt

```go
func (c *Mat2f) GetCreatedAt() atom.ID
```

#### func (*Mat2f) Init

```go
func (c *Mat2f) Init()
```

#### func (Mat2f) String

```go
func (m Mat2f) String() string
```

#### type Mat3f

```go
type Mat3f struct {
	binary.Generate
	CreatedAt atom.ID
	Col0      Vec3f
	Col1      Vec3f
	Col2      Vec3f
}
```

//////////////////////////////////////////////////////////////////////////////
class Mat3f
//////////////////////////////////////////////////////////////////////////////

#### func (*Mat3f) Class

```go
func (*Mat3f) Class() binary.Class
```

#### func (*Mat3f) GetCreatedAt

```go
func (c *Mat3f) GetCreatedAt() atom.ID
```

#### func (*Mat3f) Init

```go
func (c *Mat3f) Init()
```

#### func (Mat3f) String

```go
func (m Mat3f) String() string
```

#### type Mat4f

```go
type Mat4f struct {
	binary.Generate
	CreatedAt atom.ID
	Col0      Vec4f
	Col1      Vec4f
	Col2      Vec4f
	Col3      Vec4f
}
```

//////////////////////////////////////////////////////////////////////////////
class Mat4f
//////////////////////////////////////////////////////////////////////////////

#### func (*Mat4f) Class

```go
func (*Mat4f) Class() binary.Class
```

#### func (*Mat4f) GetCreatedAt

```go
func (c *Mat4f) GetCreatedAt() atom.ID
```

#### func (*Mat4f) Init

```go
func (c *Mat4f) Init()
```

#### func (Mat4f) String

```go
func (m Mat4f) String() string
```

#### type Objects

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class Objects
//////////////////////////////////////////////////////////////////////////////

#### func (*Objects) Class

```go
func (*Objects) Class() binary.Class
```

#### func (*Objects) GetCreatedAt

```go
func (c *Objects) GetCreatedAt() atom.ID
```

#### func (*Objects) Init

```go
func (c *Objects) Init()
```

#### type PixelStoreParameter

```go
type PixelStoreParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum PixelStoreParameter
//////////////////////////////////////////////////////////////////////////////

#### func (PixelStoreParameter) String

```go
func (v PixelStoreParameter) String() string
```

#### type PrecisionType

```go
type PrecisionType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum PrecisionType
//////////////////////////////////////////////////////////////////////////////

#### func (PrecisionType) String

```go
func (v PrecisionType) String() string
```

#### type Program

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class Program
//////////////////////////////////////////////////////////////////////////////

#### func (*Program) Class

```go
func (*Program) Class() binary.Class
```

#### func (*Program) GetCreatedAt

```go
func (c *Program) GetCreatedAt() atom.ID
```

#### func (*Program) Init

```go
func (c *Program) Init()
```

#### type ProgramArray

```go
type ProgramArray []Program
```


#### func (ProgramArray) Len

```go
func (s ProgramArray) Len() int
```

#### func (ProgramArray) Range

```go
func (s ProgramArray) Range() []Program
```

#### type ProgramId

```go
type ProgramId uint32
```


#### func (*ProgramId) Equal

```go
func (c *ProgramId) Equal(rhs ProgramId) bool
```

#### func (*ProgramId) Less

```go
func (c *ProgramId) Less(rhs ProgramId) bool
```

#### type ProgramParameter

```go
type ProgramParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ProgramParameter
//////////////////////////////////////////////////////////////////////////////

#### func (ProgramParameter) String

```go
func (v ProgramParameter) String() string
```

#### type ProgramPtr_ProgramIdMap

```go
type ProgramPtr_ProgramIdMap map[ProgramId]*Program
```


#### func (ProgramPtr_ProgramIdMap) Contains

```go
func (m ProgramPtr_ProgramIdMap) Contains(key ProgramId) bool
```

#### func (ProgramPtr_ProgramIdMap) Delete

```go
func (m ProgramPtr_ProgramIdMap) Delete(key ProgramId)
```

#### func (ProgramPtr_ProgramIdMap) Get

```go
func (m ProgramPtr_ProgramIdMap) Get(key ProgramId) *Program
```

#### func (ProgramPtr_ProgramIdMap) Range

```go
func (m ProgramPtr_ProgramIdMap) Range() []*Program
```

#### type Query

```go
type Query struct {
	binary.Generate
	CreatedAt atom.ID
}
```

//////////////////////////////////////////////////////////////////////////////
class Query
//////////////////////////////////////////////////////////////////////////////

#### func (*Query) Class

```go
func (*Query) Class() binary.Class
```

#### func (*Query) GetCreatedAt

```go
func (c *Query) GetCreatedAt() atom.ID
```

#### func (*Query) Init

```go
func (c *Query) Init()
```

#### type QueryArray

```go
type QueryArray []Query
```


#### func (QueryArray) Len

```go
func (s QueryArray) Len() int
```

#### func (QueryArray) Range

```go
func (s QueryArray) Range() []Query
```

#### type QueryId

```go
type QueryId uint32
```


#### func (*QueryId) Equal

```go
func (c *QueryId) Equal(rhs QueryId) bool
```

#### func (*QueryId) Less

```go
func (c *QueryId) Less(rhs QueryId) bool
```

#### type QueryIdArray

```go
type QueryIdArray []QueryId
```


#### func (QueryIdArray) Len

```go
func (s QueryIdArray) Len() int
```

#### func (QueryIdArray) Range

```go
func (s QueryIdArray) Range() []QueryId
```

#### type QueryObjectParameter

```go
type QueryObjectParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryObjectParameter
//////////////////////////////////////////////////////////////////////////////

#### func (QueryObjectParameter) String

```go
func (v QueryObjectParameter) String() string
```

#### type QueryObjectParameter_EXT_disjoint_timer_query

```go
type QueryObjectParameter_EXT_disjoint_timer_query uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryObjectParameter_EXT_disjoint_timer_query
//////////////////////////////////////////////////////////////////////////////

#### func (QueryObjectParameter_EXT_disjoint_timer_query) String

```go
func (v QueryObjectParameter_EXT_disjoint_timer_query) String() string
```

#### type QueryObjectParameter_GLES_3

```go
type QueryObjectParameter_GLES_3 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryObjectParameter_GLES_3
//////////////////////////////////////////////////////////////////////////////

#### func (QueryObjectParameter_GLES_3) String

```go
func (v QueryObjectParameter_GLES_3) String() string
```

#### type QueryParameter

```go
type QueryParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryParameter
//////////////////////////////////////////////////////////////////////////////

#### func (QueryParameter) String

```go
func (v QueryParameter) String() string
```

#### type QueryParameter_EXT_disjoint_timer_query

```go
type QueryParameter_EXT_disjoint_timer_query uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryParameter_EXT_disjoint_timer_query
//////////////////////////////////////////////////////////////////////////////

#### func (QueryParameter_EXT_disjoint_timer_query) String

```go
func (v QueryParameter_EXT_disjoint_timer_query) String() string
```

#### type QueryParameter_GLES_3

```go
type QueryParameter_GLES_3 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryParameter_GLES_3
//////////////////////////////////////////////////////////////////////////////

#### func (QueryParameter_GLES_3) String

```go
func (v QueryParameter_GLES_3) String() string
```

#### type QueryPtr_QueryIdMap

```go
type QueryPtr_QueryIdMap map[QueryId]*Query
```


#### func (QueryPtr_QueryIdMap) Contains

```go
func (m QueryPtr_QueryIdMap) Contains(key QueryId) bool
```

#### func (QueryPtr_QueryIdMap) Delete

```go
func (m QueryPtr_QueryIdMap) Delete(key QueryId)
```

#### func (QueryPtr_QueryIdMap) Get

```go
func (m QueryPtr_QueryIdMap) Get(key QueryId) *Query
```

#### func (QueryPtr_QueryIdMap) Range

```go
func (m QueryPtr_QueryIdMap) Range() []*Query
```

#### type QueryTarget

```go
type QueryTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryTarget
//////////////////////////////////////////////////////////////////////////////

#### func (QueryTarget) String

```go
func (v QueryTarget) String() string
```

#### type QueryTarget_EXT_disjoint_timer_query

```go
type QueryTarget_EXT_disjoint_timer_query uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryTarget_EXT_disjoint_timer_query
//////////////////////////////////////////////////////////////////////////////

#### func (QueryTarget_EXT_disjoint_timer_query) String

```go
func (v QueryTarget_EXT_disjoint_timer_query) String() string
```

#### type QueryTarget_GLES_3

```go
type QueryTarget_GLES_3 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryTarget_GLES_3
//////////////////////////////////////////////////////////////////////////////

#### func (QueryTarget_GLES_3) String

```go
func (v QueryTarget_GLES_3) String() string
```

#### type RasterizerState

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class RasterizerState
//////////////////////////////////////////////////////////////////////////////

#### func (*RasterizerState) Class

```go
func (*RasterizerState) Class() binary.Class
```

#### func (*RasterizerState) GetCreatedAt

```go
func (c *RasterizerState) GetCreatedAt() atom.ID
```

#### func (*RasterizerState) Init

```go
func (c *RasterizerState) Init()
```

#### type Rect

```go
type Rect struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
	Width     int32
	Height    int32
}
```

//////////////////////////////////////////////////////////////////////////////
class Rect
//////////////////////////////////////////////////////////////////////////////

#### func (*Rect) Class

```go
func (*Rect) Class() binary.Class
```

#### func (*Rect) GetCreatedAt

```go
func (c *Rect) GetCreatedAt() atom.ID
```

#### func (*Rect) Init

```go
func (c *Rect) Init()
```

#### type Renderbuffer

```go
type Renderbuffer struct {
	binary.Generate
	CreatedAt atom.ID
	Width     int32
	Height    int32
	Data      memory.Memory
	Format    RenderbufferFormat
}
```

//////////////////////////////////////////////////////////////////////////////
class Renderbuffer
//////////////////////////////////////////////////////////////////////////////

#### func (*Renderbuffer) Class

```go
func (*Renderbuffer) Class() binary.Class
```

#### func (*Renderbuffer) GetCreatedAt

```go
func (c *Renderbuffer) GetCreatedAt() atom.ID
```

#### func (*Renderbuffer) Init

```go
func (c *Renderbuffer) Init()
```

#### type RenderbufferArray

```go
type RenderbufferArray []Renderbuffer
```


#### func (RenderbufferArray) Len

```go
func (s RenderbufferArray) Len() int
```

#### func (RenderbufferArray) Range

```go
func (s RenderbufferArray) Range() []Renderbuffer
```

#### type RenderbufferFormat

```go
type RenderbufferFormat uint32
```

//////////////////////////////////////////////////////////////////////////////
enum RenderbufferFormat
//////////////////////////////////////////////////////////////////////////////

#### func (RenderbufferFormat) String

```go
func (v RenderbufferFormat) String() string
```

#### type RenderbufferId

```go
type RenderbufferId uint32
```


#### func (*RenderbufferId) Equal

```go
func (c *RenderbufferId) Equal(rhs RenderbufferId) bool
```

#### func (*RenderbufferId) Less

```go
func (c *RenderbufferId) Less(rhs RenderbufferId) bool
```

#### type RenderbufferIdArray

```go
type RenderbufferIdArray []RenderbufferId
```


#### func (RenderbufferIdArray) Len

```go
func (s RenderbufferIdArray) Len() int
```

#### func (RenderbufferIdArray) Range

```go
func (s RenderbufferIdArray) Range() []RenderbufferId
```

#### type RenderbufferId_RenderbufferTargetMap

```go
type RenderbufferId_RenderbufferTargetMap map[RenderbufferTarget]RenderbufferId
```


#### func (RenderbufferId_RenderbufferTargetMap) Contains

```go
func (m RenderbufferId_RenderbufferTargetMap) Contains(key RenderbufferTarget) bool
```

#### func (RenderbufferId_RenderbufferTargetMap) Delete

```go
func (m RenderbufferId_RenderbufferTargetMap) Delete(key RenderbufferTarget)
```

#### func (RenderbufferId_RenderbufferTargetMap) Get

```go
func (m RenderbufferId_RenderbufferTargetMap) Get(key RenderbufferTarget) RenderbufferId
```

#### func (RenderbufferId_RenderbufferTargetMap) Range

```go
func (m RenderbufferId_RenderbufferTargetMap) Range() []RenderbufferId
```

#### type RenderbufferParameter

```go
type RenderbufferParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum RenderbufferParameter
//////////////////////////////////////////////////////////////////////////////

#### func (RenderbufferParameter) String

```go
func (v RenderbufferParameter) String() string
```

#### type RenderbufferPtr_RenderbufferIdMap

```go
type RenderbufferPtr_RenderbufferIdMap map[RenderbufferId]*Renderbuffer
```


#### func (RenderbufferPtr_RenderbufferIdMap) Contains

```go
func (m RenderbufferPtr_RenderbufferIdMap) Contains(key RenderbufferId) bool
```

#### func (RenderbufferPtr_RenderbufferIdMap) Delete

```go
func (m RenderbufferPtr_RenderbufferIdMap) Delete(key RenderbufferId)
```

#### func (RenderbufferPtr_RenderbufferIdMap) Get

```go
func (m RenderbufferPtr_RenderbufferIdMap) Get(key RenderbufferId) *Renderbuffer
```

#### func (RenderbufferPtr_RenderbufferIdMap) Range

```go
func (m RenderbufferPtr_RenderbufferIdMap) Range() []*Renderbuffer
```

#### type RenderbufferTarget

```go
type RenderbufferTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum RenderbufferTarget
//////////////////////////////////////////////////////////////////////////////

#### func (RenderbufferTarget) String

```go
func (v RenderbufferTarget) String() string
```

#### type ReplayBindRenderer

```go
type ReplayBindRenderer struct {
	binary.Generate
	Id uint32
}
```

//////////////////////////////////////////////////////////////////////////////
ReplayBindRenderer
//////////////////////////////////////////////////////////////////////////////

#### func  NewReplayBindRenderer

```go
func NewReplayBindRenderer(
	pId uint32,
) *ReplayBindRenderer
```

#### func (*ReplayBindRenderer) API

```go
func (c *ReplayBindRenderer) API() gfxapi.API
```

#### func (*ReplayBindRenderer) Class

```go
func (*ReplayBindRenderer) Class() binary.Class
```

#### func (*ReplayBindRenderer) Flags

```go
func (c *ReplayBindRenderer) Flags() atom.Flags
```

#### func (*ReplayBindRenderer) Mutate

```go
func (ϟa *ReplayBindRenderer) Mutate(ϟs *gfxapi.State) error
```

#### func (*ReplayBindRenderer) Replay

```go
func (ϟa *ReplayBindRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*ReplayBindRenderer) String

```go
func (c *ReplayBindRenderer) String() string
```

#### func (*ReplayBindRenderer) TypeID

```go
func (c *ReplayBindRenderer) TypeID() atom.TypeID
```

#### type ReplayCreateRenderer

```go
type ReplayCreateRenderer struct {
	binary.Generate
	Id uint32
}
```

//////////////////////////////////////////////////////////////////////////////
ReplayCreateRenderer
//////////////////////////////////////////////////////////////////////////////

#### func  NewReplayCreateRenderer

```go
func NewReplayCreateRenderer(
	pId uint32,
) *ReplayCreateRenderer
```

#### func (*ReplayCreateRenderer) API

```go
func (c *ReplayCreateRenderer) API() gfxapi.API
```

#### func (*ReplayCreateRenderer) Class

```go
func (*ReplayCreateRenderer) Class() binary.Class
```

#### func (*ReplayCreateRenderer) Flags

```go
func (c *ReplayCreateRenderer) Flags() atom.Flags
```

#### func (*ReplayCreateRenderer) Mutate

```go
func (ϟa *ReplayCreateRenderer) Mutate(ϟs *gfxapi.State) error
```

#### func (*ReplayCreateRenderer) Replay

```go
func (ϟa *ReplayCreateRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*ReplayCreateRenderer) String

```go
func (c *ReplayCreateRenderer) String() string
```

#### func (*ReplayCreateRenderer) TypeID

```go
func (c *ReplayCreateRenderer) TypeID() atom.TypeID
```

#### type ResetStatus

```go
type ResetStatus uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ResetStatus
//////////////////////////////////////////////////////////////////////////////

#### func (ResetStatus) String

```go
func (v ResetStatus) String() string
```

#### type S32Array

```go
type S32Array []int32
```


#### func (S32Array) Len

```go
func (s S32Array) Len() int
```

#### func (S32Array) Range

```go
func (s S32Array) Range() []int32
```

#### type S32_PixelStoreParameterMap

```go
type S32_PixelStoreParameterMap map[PixelStoreParameter]int32
```


#### func (S32_PixelStoreParameterMap) Contains

```go
func (m S32_PixelStoreParameterMap) Contains(key PixelStoreParameter) bool
```

#### func (S32_PixelStoreParameterMap) Delete

```go
func (m S32_PixelStoreParameterMap) Delete(key PixelStoreParameter)
```

#### func (S32_PixelStoreParameterMap) Get

```go
func (m S32_PixelStoreParameterMap) Get(key PixelStoreParameter) int32
```

#### func (S32_PixelStoreParameterMap) Range

```go
func (m S32_PixelStoreParameterMap) Range() []int32
```

#### type Shader

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class Shader
//////////////////////////////////////////////////////////////////////////////

#### func (*Shader) Class

```go
func (*Shader) Class() binary.Class
```

#### func (*Shader) GetCreatedAt

```go
func (c *Shader) GetCreatedAt() atom.ID
```

#### func (*Shader) Init

```go
func (c *Shader) Init()
```

#### type ShaderArray

```go
type ShaderArray []Shader
```


#### func (ShaderArray) Len

```go
func (s ShaderArray) Len() int
```

#### func (ShaderArray) Range

```go
func (s ShaderArray) Range() []Shader
```

#### type ShaderAttribType

```go
type ShaderAttribType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderAttribType
//////////////////////////////////////////////////////////////////////////////

#### func (ShaderAttribType) String

```go
func (v ShaderAttribType) String() string
```

#### type ShaderId

```go
type ShaderId uint32
```


#### func (*ShaderId) Equal

```go
func (c *ShaderId) Equal(rhs ShaderId) bool
```

#### func (*ShaderId) Less

```go
func (c *ShaderId) Less(rhs ShaderId) bool
```

#### type ShaderIdArray

```go
type ShaderIdArray []ShaderId
```


#### func (ShaderIdArray) Len

```go
func (s ShaderIdArray) Len() int
```

#### func (ShaderIdArray) Range

```go
func (s ShaderIdArray) Range() []ShaderId
```

#### type ShaderId_ShaderTypeMap

```go
type ShaderId_ShaderTypeMap map[ShaderType]ShaderId
```


#### func (ShaderId_ShaderTypeMap) Contains

```go
func (m ShaderId_ShaderTypeMap) Contains(key ShaderType) bool
```

#### func (ShaderId_ShaderTypeMap) Delete

```go
func (m ShaderId_ShaderTypeMap) Delete(key ShaderType)
```

#### func (ShaderId_ShaderTypeMap) Get

```go
func (m ShaderId_ShaderTypeMap) Get(key ShaderType) ShaderId
```

#### func (ShaderId_ShaderTypeMap) Range

```go
func (m ShaderId_ShaderTypeMap) Range() []ShaderId
```

#### type ShaderParameter

```go
type ShaderParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderParameter
//////////////////////////////////////////////////////////////////////////////

#### func (ShaderParameter) String

```go
func (v ShaderParameter) String() string
```

#### type ShaderPtr_ShaderIdMap

```go
type ShaderPtr_ShaderIdMap map[ShaderId]*Shader
```


#### func (ShaderPtr_ShaderIdMap) Contains

```go
func (m ShaderPtr_ShaderIdMap) Contains(key ShaderId) bool
```

#### func (ShaderPtr_ShaderIdMap) Delete

```go
func (m ShaderPtr_ShaderIdMap) Delete(key ShaderId)
```

#### func (ShaderPtr_ShaderIdMap) Get

```go
func (m ShaderPtr_ShaderIdMap) Get(key ShaderId) *Shader
```

#### func (ShaderPtr_ShaderIdMap) Range

```go
func (m ShaderPtr_ShaderIdMap) Range() []*Shader
```

#### type ShaderType

```go
type ShaderType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderType
//////////////////////////////////////////////////////////////////////////////

#### func (ShaderType) String

```go
func (v ShaderType) String() string
```

#### type ShaderUniformType

```go
type ShaderUniformType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderUniformType
//////////////////////////////////////////////////////////////////////////////

#### func (ShaderUniformType) String

```go
func (v ShaderUniformType) String() string
```

#### type StartTimer

```go
type StartTimer struct {
	binary.Generate
	Index uint8
}
```

//////////////////////////////////////////////////////////////////////////////
StartTimer
//////////////////////////////////////////////////////////////////////////////

#### func  NewStartTimer

```go
func NewStartTimer(
	pIndex uint8,
) *StartTimer
```

#### func (*StartTimer) API

```go
func (c *StartTimer) API() gfxapi.API
```

#### func (*StartTimer) Class

```go
func (*StartTimer) Class() binary.Class
```

#### func (*StartTimer) Flags

```go
func (c *StartTimer) Flags() atom.Flags
```

#### func (*StartTimer) Mutate

```go
func (ϟa *StartTimer) Mutate(ϟs *gfxapi.State) error
```

#### func (*StartTimer) Replay

```go
func (ϟa *StartTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*StartTimer) String

```go
func (c *StartTimer) String() string
```

#### func (*StartTimer) TypeID

```go
func (c *StartTimer) TypeID() atom.TypeID
```

#### type State

```go
type State struct {
	Globals
	ValidateOutput bool
}
```


#### type StateVariable

```go
type StateVariable uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StateVariable
//////////////////////////////////////////////////////////////////////////////

#### func (StateVariable) String

```go
func (v StateVariable) String() string
```

#### type StateVariable_EXT_disjoint_timer_query

```go
type StateVariable_EXT_disjoint_timer_query uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StateVariable_EXT_disjoint_timer_query
//////////////////////////////////////////////////////////////////////////////

#### func (StateVariable_EXT_disjoint_timer_query) String

```go
func (v StateVariable_EXT_disjoint_timer_query) String() string
```

#### type StateVariable_EXT_texture_filter_anisotropic

```go
type StateVariable_EXT_texture_filter_anisotropic uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StateVariable_EXT_texture_filter_anisotropic
//////////////////////////////////////////////////////////////////////////////

#### func (StateVariable_EXT_texture_filter_anisotropic) String

```go
func (v StateVariable_EXT_texture_filter_anisotropic) String() string
```

#### type StateVariable_GLES_2_0

```go
type StateVariable_GLES_2_0 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StateVariable_GLES_2_0
//////////////////////////////////////////////////////////////////////////////

#### func (StateVariable_GLES_2_0) String

```go
func (v StateVariable_GLES_2_0) String() string
```

#### type StateVariable_GLES_3_1

```go
type StateVariable_GLES_3_1 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StateVariable_GLES_3_1
//////////////////////////////////////////////////////////////////////////////

#### func (StateVariable_GLES_3_1) String

```go
func (v StateVariable_GLES_3_1) String() string
```

#### type StencilAction

```go
type StencilAction uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StencilAction
//////////////////////////////////////////////////////////////////////////////

#### func (StencilAction) String

```go
func (v StencilAction) String() string
```

#### type StopTimer

```go
type StopTimer struct {
	binary.Generate
	Index  uint8
	Result uint64
}
```

//////////////////////////////////////////////////////////////////////////////
StopTimer
//////////////////////////////////////////////////////////////////////////////

#### func  NewStopTimer

```go
func NewStopTimer(
	pIndex uint8,
	pResult uint64,
) *StopTimer
```

#### func (*StopTimer) API

```go
func (c *StopTimer) API() gfxapi.API
```

#### func (*StopTimer) Class

```go
func (*StopTimer) Class() binary.Class
```

#### func (*StopTimer) Flags

```go
func (c *StopTimer) Flags() atom.Flags
```

#### func (*StopTimer) Mutate

```go
func (ϟa *StopTimer) Mutate(ϟs *gfxapi.State) error
```

#### func (*StopTimer) Replay

```go
func (ϟa *StopTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟb *builder.Builder, postback bool)
```

#### func (*StopTimer) String

```go
func (c *StopTimer) String() string
```

#### func (*StopTimer) TypeID

```go
func (c *StopTimer) TypeID() atom.TypeID
```

#### type StopTimer_Postback

```go
type StopTimer_Postback struct {
	Result uint64
}
```


#### func (*StopTimer_Postback) Decode

```go
func (o *StopTimer_Postback) Decode(d binary.Decoder) error
```

#### type StringArray

```go
type StringArray []string
```


#### func (StringArray) Len

```go
func (s StringArray) Len() int
```

#### func (StringArray) Range

```go
func (s StringArray) Range() []string
```

#### type StringConstant

```go
type StringConstant uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StringConstant
//////////////////////////////////////////////////////////////////////////////

#### func (StringConstant) String

```go
func (v StringConstant) String() string
```

#### type TestFunction

```go
type TestFunction uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TestFunction
//////////////////////////////////////////////////////////////////////////////

#### func (TestFunction) String

```go
func (v TestFunction) String() string
```

#### type TexelComponent

```go
type TexelComponent uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TexelComponent
//////////////////////////////////////////////////////////////////////////////

#### func (TexelComponent) String

```go
func (v TexelComponent) String() string
```

#### type TexelFormat

```go
type TexelFormat uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TexelFormat
//////////////////////////////////////////////////////////////////////////////

#### func (TexelFormat) String

```go
func (v TexelFormat) String() string
```

#### type TexelFormat_GLES_1_1

```go
type TexelFormat_GLES_1_1 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TexelFormat_GLES_1_1
//////////////////////////////////////////////////////////////////////////////

#### func (TexelFormat_GLES_1_1) String

```go
func (v TexelFormat_GLES_1_1) String() string
```

#### type TexelFormat_GLES_3_0

```go
type TexelFormat_GLES_3_0 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TexelFormat_GLES_3_0
//////////////////////////////////////////////////////////////////////////////

#### func (TexelFormat_GLES_3_0) String

```go
func (v TexelFormat_GLES_3_0) String() string
```

#### type TexelType

```go
type TexelType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TexelType
//////////////////////////////////////////////////////////////////////////////

#### func (TexelType) String

```go
func (v TexelType) String() string
```

#### type Texture

```go
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
}
```

//////////////////////////////////////////////////////////////////////////////
class Texture
//////////////////////////////////////////////////////////////////////////////

#### func (*Texture) Class

```go
func (*Texture) Class() binary.Class
```

#### func (*Texture) GetCreatedAt

```go
func (c *Texture) GetCreatedAt() atom.ID
```

#### func (*Texture) Init

```go
func (c *Texture) Init()
```

#### type Texture2DImageTarget

```go
type Texture2DImageTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum Texture2DImageTarget
//////////////////////////////////////////////////////////////////////////////

#### func (Texture2DImageTarget) String

```go
func (v Texture2DImageTarget) String() string
```

#### type TextureArray

```go
type TextureArray []Texture
```


#### func (TextureArray) Len

```go
func (s TextureArray) Len() int
```

#### func (TextureArray) Range

```go
func (s TextureArray) Range() []Texture
```

#### type TextureFilterMode

```go
type TextureFilterMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureFilterMode
//////////////////////////////////////////////////////////////////////////////

#### func (TextureFilterMode) String

```go
func (v TextureFilterMode) String() string
```

#### type TextureId

```go
type TextureId uint32
```


#### func (*TextureId) Equal

```go
func (c *TextureId) Equal(rhs TextureId) bool
```

#### func (*TextureId) Less

```go
func (c *TextureId) Less(rhs TextureId) bool
```

#### type TextureIdArray

```go
type TextureIdArray []TextureId
```


#### func (TextureIdArray) Len

```go
func (s TextureIdArray) Len() int
```

#### func (TextureIdArray) Range

```go
func (s TextureIdArray) Range() []TextureId
```

#### type TextureId_TextureTargetMap

```go
type TextureId_TextureTargetMap map[TextureTarget]TextureId
```


#### func (TextureId_TextureTargetMap) Contains

```go
func (m TextureId_TextureTargetMap) Contains(key TextureTarget) bool
```

#### func (TextureId_TextureTargetMap) Delete

```go
func (m TextureId_TextureTargetMap) Delete(key TextureTarget)
```

#### func (TextureId_TextureTargetMap) Get

```go
func (m TextureId_TextureTargetMap) Get(key TextureTarget) TextureId
```

#### func (TextureId_TextureTargetMap) Range

```go
func (m TextureId_TextureTargetMap) Range() []TextureId
```

#### type TextureId_TextureTargetMap_TextureUnitMap

```go
type TextureId_TextureTargetMap_TextureUnitMap map[TextureUnit]TextureId_TextureTargetMap
```


#### func (TextureId_TextureTargetMap_TextureUnitMap) Contains

```go
func (m TextureId_TextureTargetMap_TextureUnitMap) Contains(key TextureUnit) bool
```

#### func (TextureId_TextureTargetMap_TextureUnitMap) Delete

```go
func (m TextureId_TextureTargetMap_TextureUnitMap) Delete(key TextureUnit)
```

#### func (TextureId_TextureTargetMap_TextureUnitMap) Get

```go
func (m TextureId_TextureTargetMap_TextureUnitMap) Get(key TextureUnit) TextureId_TextureTargetMap
```

#### func (TextureId_TextureTargetMap_TextureUnitMap) Range

```go
func (m TextureId_TextureTargetMap_TextureUnitMap) Range() []TextureId_TextureTargetMap
```

#### type TextureImageTarget

```go
type TextureImageTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureImageTarget
//////////////////////////////////////////////////////////////////////////////

#### func (TextureImageTarget) String

```go
func (v TextureImageTarget) String() string
```

#### type TextureKind

```go
type TextureKind uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureKind
//////////////////////////////////////////////////////////////////////////////

#### func (TextureKind) String

```go
func (v TextureKind) String() string
```

#### type TextureParameter

```go
type TextureParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureParameter
//////////////////////////////////////////////////////////////////////////////

#### func (TextureParameter) String

```go
func (v TextureParameter) String() string
```

#### type TextureParameter_EXT_texture_filter_anisotropic

```go
type TextureParameter_EXT_texture_filter_anisotropic uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureParameter_EXT_texture_filter_anisotropic
//////////////////////////////////////////////////////////////////////////////

#### func (TextureParameter_EXT_texture_filter_anisotropic) String

```go
func (v TextureParameter_EXT_texture_filter_anisotropic) String() string
```

#### type TextureParameter_FilterMode

```go
type TextureParameter_FilterMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureParameter_FilterMode
//////////////////////////////////////////////////////////////////////////////

#### func (TextureParameter_FilterMode) String

```go
func (v TextureParameter_FilterMode) String() string
```

#### type TextureParameter_SwizzleMode

```go
type TextureParameter_SwizzleMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureParameter_SwizzleMode
//////////////////////////////////////////////////////////////////////////////

#### func (TextureParameter_SwizzleMode) String

```go
func (v TextureParameter_SwizzleMode) String() string
```

#### type TextureParameter_WrapMode

```go
type TextureParameter_WrapMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureParameter_WrapMode
//////////////////////////////////////////////////////////////////////////////

#### func (TextureParameter_WrapMode) String

```go
func (v TextureParameter_WrapMode) String() string
```

#### type TexturePointer

```go
type TexturePointer memory.Pointer
```


#### func (*TexturePointer) Equal

```go
func (c *TexturePointer) Equal(rhs TexturePointer) bool
```

#### func (*TexturePointer) Less

```go
func (c *TexturePointer) Less(rhs TexturePointer) bool
```

#### type TexturePtr_TextureIdMap

```go
type TexturePtr_TextureIdMap map[TextureId]*Texture
```


#### func (TexturePtr_TextureIdMap) Contains

```go
func (m TexturePtr_TextureIdMap) Contains(key TextureId) bool
```

#### func (TexturePtr_TextureIdMap) Delete

```go
func (m TexturePtr_TextureIdMap) Delete(key TextureId)
```

#### func (TexturePtr_TextureIdMap) Get

```go
func (m TexturePtr_TextureIdMap) Get(key TextureId) *Texture
```

#### func (TexturePtr_TextureIdMap) Range

```go
func (m TexturePtr_TextureIdMap) Range() []*Texture
```

#### type TextureTarget

```go
type TextureTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureTarget
//////////////////////////////////////////////////////////////////////////////

#### func (TextureTarget) String

```go
func (v TextureTarget) String() string
```

#### type TextureTarget_GLES_1_1

```go
type TextureTarget_GLES_1_1 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureTarget_GLES_1_1
//////////////////////////////////////////////////////////////////////////////

#### func (TextureTarget_GLES_1_1) String

```go
func (v TextureTarget_GLES_1_1) String() string
```

#### type TextureTarget_GLES_2_0

```go
type TextureTarget_GLES_2_0 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureTarget_GLES_2_0
//////////////////////////////////////////////////////////////////////////////

#### func (TextureTarget_GLES_2_0) String

```go
func (v TextureTarget_GLES_2_0) String() string
```

#### type TextureTarget_OES_EGL_image_external

```go
type TextureTarget_OES_EGL_image_external uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureTarget_OES_EGL_image_external
//////////////////////////////////////////////////////////////////////////////

#### func (TextureTarget_OES_EGL_image_external) String

```go
func (v TextureTarget_OES_EGL_image_external) String() string
```

#### type TextureUnit

```go
type TextureUnit uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureUnit
//////////////////////////////////////////////////////////////////////////////

#### func (TextureUnit) String

```go
func (v TextureUnit) String() string
```

#### type TextureWrapMode

```go
type TextureWrapMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureWrapMode
//////////////////////////////////////////////////////////////////////////////

#### func (TextureWrapMode) String

```go
func (v TextureWrapMode) String() string
```

#### type ThreadID

```go
type ThreadID uint32
```


#### func (*ThreadID) Equal

```go
func (c *ThreadID) Equal(rhs ThreadID) bool
```

#### func (*ThreadID) Less

```go
func (c *ThreadID) Less(rhs ThreadID) bool
```

#### type TilePreserveMaskQCOM

```go
type TilePreserveMaskQCOM uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TilePreserveMaskQCOM
//////////////////////////////////////////////////////////////////////////////

#### func (TilePreserveMaskQCOM) String

```go
func (v TilePreserveMaskQCOM) String() string
```

#### type Type_ARB_half_float_vertex

```go
type Type_ARB_half_float_vertex uint32
```

//////////////////////////////////////////////////////////////////////////////
enum Type_ARB_half_float_vertex
//////////////////////////////////////////////////////////////////////////////

#### func (Type_ARB_half_float_vertex) String

```go
func (v Type_ARB_half_float_vertex) String() string
```

#### type Type_OES_vertex_half_float

```go
type Type_OES_vertex_half_float uint32
```

//////////////////////////////////////////////////////////////////////////////
enum Type_OES_vertex_half_float
//////////////////////////////////////////////////////////////////////////////

#### func (Type_OES_vertex_half_float) String

```go
func (v Type_OES_vertex_half_float) String() string
```

#### type U32_FaceModeMap

```go
type U32_FaceModeMap map[FaceMode]uint32
```


#### func (U32_FaceModeMap) Contains

```go
func (m U32_FaceModeMap) Contains(key FaceMode) bool
```

#### func (U32_FaceModeMap) Delete

```go
func (m U32_FaceModeMap) Delete(key FaceMode)
```

#### func (U32_FaceModeMap) Get

```go
func (m U32_FaceModeMap) Get(key FaceMode) uint32
```

#### func (U32_FaceModeMap) Range

```go
func (m U32_FaceModeMap) Range() []uint32
```

#### type Uniform

```go
type Uniform struct {
	binary.Generate
	CreatedAt atom.ID
	Name      string
	Type      ShaderUniformType
	Value     UniformValue
}
```

//////////////////////////////////////////////////////////////////////////////
class Uniform
//////////////////////////////////////////////////////////////////////////////

#### func (*Uniform) Class

```go
func (*Uniform) Class() binary.Class
```

#### func (*Uniform) GetCreatedAt

```go
func (c *Uniform) GetCreatedAt() atom.ID
```

#### func (*Uniform) Init

```go
func (c *Uniform) Init()
```

#### func (Uniform) String

```go
func (u Uniform) String() string
```

#### type UniformLocation

```go
type UniformLocation int32
```


#### func (*UniformLocation) Equal

```go
func (c *UniformLocation) Equal(rhs UniformLocation) bool
```

#### func (*UniformLocation) Less

```go
func (c *UniformLocation) Less(rhs UniformLocation) bool
```

#### type UniformValue

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class UniformValue
//////////////////////////////////////////////////////////////////////////////

#### func (*UniformValue) Class

```go
func (*UniformValue) Class() binary.Class
```

#### func (*UniformValue) GetCreatedAt

```go
func (c *UniformValue) GetCreatedAt() atom.ID
```

#### func (*UniformValue) Init

```go
func (c *UniformValue) Init()
```

#### type Uniform_UniformLocationMap

```go
type Uniform_UniformLocationMap map[UniformLocation]Uniform
```


#### func (Uniform_UniformLocationMap) Contains

```go
func (m Uniform_UniformLocationMap) Contains(key UniformLocation) bool
```

#### func (Uniform_UniformLocationMap) Delete

```go
func (m Uniform_UniformLocationMap) Delete(key UniformLocation)
```

#### func (Uniform_UniformLocationMap) Get

```go
func (m Uniform_UniformLocationMap) Get(key UniformLocation) Uniform
```

#### func (Uniform_UniformLocationMap) Range

```go
func (m Uniform_UniformLocationMap) Range() []Uniform
```

#### type Vec2f

```go
type Vec2f struct {
	binary.Generate
	CreatedAt atom.ID
	X         float32
	Y         float32
}
```

//////////////////////////////////////////////////////////////////////////////
class Vec2f
//////////////////////////////////////////////////////////////////////////////

#### func (*Vec2f) Class

```go
func (*Vec2f) Class() binary.Class
```

#### func (*Vec2f) GetCreatedAt

```go
func (c *Vec2f) GetCreatedAt() atom.ID
```

#### func (*Vec2f) Init

```go
func (c *Vec2f) Init()
```

#### func (Vec2f) String

```go
func (v Vec2f) String() string
```

#### type Vec2i

```go
type Vec2i struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
}
```

//////////////////////////////////////////////////////////////////////////////
class Vec2i
//////////////////////////////////////////////////////////////////////////////

#### func (*Vec2i) Class

```go
func (*Vec2i) Class() binary.Class
```

#### func (*Vec2i) GetCreatedAt

```go
func (c *Vec2i) GetCreatedAt() atom.ID
```

#### func (*Vec2i) Init

```go
func (c *Vec2i) Init()
```

#### func (Vec2i) String

```go
func (v Vec2i) String() string
```

#### type Vec3f

```go
type Vec3f struct {
	binary.Generate
	CreatedAt atom.ID
	X         float32
	Y         float32
	Z         float32
}
```

//////////////////////////////////////////////////////////////////////////////
class Vec3f
//////////////////////////////////////////////////////////////////////////////

#### func (*Vec3f) Class

```go
func (*Vec3f) Class() binary.Class
```

#### func (*Vec3f) GetCreatedAt

```go
func (c *Vec3f) GetCreatedAt() atom.ID
```

#### func (*Vec3f) Init

```go
func (c *Vec3f) Init()
```

#### func (Vec3f) String

```go
func (v Vec3f) String() string
```

#### type Vec3i

```go
type Vec3i struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
	Z         int32
}
```

//////////////////////////////////////////////////////////////////////////////
class Vec3i
//////////////////////////////////////////////////////////////////////////////

#### func (*Vec3i) Class

```go
func (*Vec3i) Class() binary.Class
```

#### func (*Vec3i) GetCreatedAt

```go
func (c *Vec3i) GetCreatedAt() atom.ID
```

#### func (*Vec3i) Init

```go
func (c *Vec3i) Init()
```

#### func (Vec3i) String

```go
func (v Vec3i) String() string
```

#### type Vec4f

```go
type Vec4f struct {
	binary.Generate
	CreatedAt atom.ID
	X         float32
	Y         float32
	Z         float32
	W         float32
}
```

//////////////////////////////////////////////////////////////////////////////
class Vec4f
//////////////////////////////////////////////////////////////////////////////

#### func (*Vec4f) Class

```go
func (*Vec4f) Class() binary.Class
```

#### func (*Vec4f) GetCreatedAt

```go
func (c *Vec4f) GetCreatedAt() atom.ID
```

#### func (*Vec4f) Init

```go
func (c *Vec4f) Init()
```

#### func (Vec4f) String

```go
func (v Vec4f) String() string
```

#### type Vec4i

```go
type Vec4i struct {
	binary.Generate
	CreatedAt atom.ID
	X         int32
	Y         int32
	Z         int32
	W         int32
}
```

//////////////////////////////////////////////////////////////////////////////
class Vec4i
//////////////////////////////////////////////////////////////////////////////

#### func (*Vec4i) Class

```go
func (*Vec4i) Class() binary.Class
```

#### func (*Vec4i) GetCreatedAt

```go
func (c *Vec4i) GetCreatedAt() atom.ID
```

#### func (*Vec4i) Init

```go
func (c *Vec4i) Init()
```

#### func (Vec4i) String

```go
func (v Vec4i) String() string
```

#### type VertexArray

```go
type VertexArray struct {
	binary.Generate
	CreatedAt atom.ID
}
```

//////////////////////////////////////////////////////////////////////////////
class VertexArray
//////////////////////////////////////////////////////////////////////////////

#### func (*VertexArray) Class

```go
func (*VertexArray) Class() binary.Class
```

#### func (*VertexArray) GetCreatedAt

```go
func (c *VertexArray) GetCreatedAt() atom.ID
```

#### func (*VertexArray) Init

```go
func (c *VertexArray) Init()
```

#### type VertexArrayArray

```go
type VertexArrayArray []VertexArray
```


#### func (VertexArrayArray) Len

```go
func (s VertexArrayArray) Len() int
```

#### func (VertexArrayArray) Range

```go
func (s VertexArrayArray) Range() []VertexArray
```

#### type VertexArrayId

```go
type VertexArrayId uint32
```


#### func (*VertexArrayId) Equal

```go
func (c *VertexArrayId) Equal(rhs VertexArrayId) bool
```

#### func (*VertexArrayId) Less

```go
func (c *VertexArrayId) Less(rhs VertexArrayId) bool
```

#### type VertexArrayIdArray

```go
type VertexArrayIdArray []VertexArrayId
```


#### func (VertexArrayIdArray) Len

```go
func (s VertexArrayIdArray) Len() int
```

#### func (VertexArrayIdArray) Range

```go
func (s VertexArrayIdArray) Range() []VertexArrayId
```

#### type VertexArrayPtr_VertexArrayIdMap

```go
type VertexArrayPtr_VertexArrayIdMap map[VertexArrayId]*VertexArray
```


#### func (VertexArrayPtr_VertexArrayIdMap) Contains

```go
func (m VertexArrayPtr_VertexArrayIdMap) Contains(key VertexArrayId) bool
```

#### func (VertexArrayPtr_VertexArrayIdMap) Delete

```go
func (m VertexArrayPtr_VertexArrayIdMap) Delete(key VertexArrayId)
```

#### func (VertexArrayPtr_VertexArrayIdMap) Get

```go
func (m VertexArrayPtr_VertexArrayIdMap) Get(key VertexArrayId) *VertexArray
```

#### func (VertexArrayPtr_VertexArrayIdMap) Range

```go
func (m VertexArrayPtr_VertexArrayIdMap) Range() []*VertexArray
```

#### type VertexAttribType

```go
type VertexAttribType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum VertexAttribType
//////////////////////////////////////////////////////////////////////////////

#### func (VertexAttribType) String

```go
func (v VertexAttribType) String() string
```

#### type VertexAttribute

```go
type VertexAttribute struct {
	binary.Generate
	CreatedAt   atom.ID
	Name        string
	VectorCount int32
	Type        ShaderAttribType
}
```

//////////////////////////////////////////////////////////////////////////////
class VertexAttribute
//////////////////////////////////////////////////////////////////////////////

#### func (*VertexAttribute) Class

```go
func (*VertexAttribute) Class() binary.Class
```

#### func (*VertexAttribute) GetCreatedAt

```go
func (c *VertexAttribute) GetCreatedAt() atom.ID
```

#### func (*VertexAttribute) Init

```go
func (c *VertexAttribute) Init()
```

#### type VertexAttributeArray

```go
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
```

//////////////////////////////////////////////////////////////////////////////
class VertexAttributeArray
//////////////////////////////////////////////////////////////////////////////

#### func (*VertexAttributeArray) Class

```go
func (*VertexAttributeArray) Class() binary.Class
```

#### func (*VertexAttributeArray) GetCreatedAt

```go
func (c *VertexAttributeArray) GetCreatedAt() atom.ID
```

#### func (*VertexAttributeArray) Init

```go
func (c *VertexAttributeArray) Init()
```

#### func (VertexAttributeArray) String

```go
func (a VertexAttributeArray) String() string
```

#### type VertexAttributeArrayArray

```go
type VertexAttributeArrayArray []VertexAttributeArray
```


#### func (VertexAttributeArrayArray) Len

```go
func (s VertexAttributeArrayArray) Len() int
```

#### func (VertexAttributeArrayArray) Range

```go
func (s VertexAttributeArrayArray) Range() []VertexAttributeArray
```

#### type VertexAttributeArrayPtr_AttributeLocationMap

```go
type VertexAttributeArrayPtr_AttributeLocationMap map[AttributeLocation]*VertexAttributeArray
```


#### func (VertexAttributeArrayPtr_AttributeLocationMap) Contains

```go
func (m VertexAttributeArrayPtr_AttributeLocationMap) Contains(key AttributeLocation) bool
```

#### func (VertexAttributeArrayPtr_AttributeLocationMap) Delete

```go
func (m VertexAttributeArrayPtr_AttributeLocationMap) Delete(key AttributeLocation)
```

#### func (VertexAttributeArrayPtr_AttributeLocationMap) Get

```go
func (m VertexAttributeArrayPtr_AttributeLocationMap) Get(key AttributeLocation) *VertexAttributeArray
```

#### func (VertexAttributeArrayPtr_AttributeLocationMap) Range

```go
func (m VertexAttributeArrayPtr_AttributeLocationMap) Range() []*VertexAttributeArray
```

#### type VertexAttribute_s32Map

```go
type VertexAttribute_s32Map map[int32]VertexAttribute
```


#### func (VertexAttribute_s32Map) Contains

```go
func (m VertexAttribute_s32Map) Contains(key int32) bool
```

#### func (VertexAttribute_s32Map) Delete

```go
func (m VertexAttribute_s32Map) Delete(key int32)
```

#### func (VertexAttribute_s32Map) Get

```go
func (m VertexAttribute_s32Map) Get(key int32) VertexAttribute
```

#### func (VertexAttribute_s32Map) Range

```go
func (m VertexAttribute_s32Map) Range() []VertexAttribute
```

#### type VertexPointer

```go
type VertexPointer memory.Pointer
```


#### func (*VertexPointer) Equal

```go
func (c *VertexPointer) Equal(rhs VertexPointer) bool
```

#### func (*VertexPointer) Less

```go
func (c *VertexPointer) Less(rhs VertexPointer) bool
```

#### type WglCreateContext

```go
type WglCreateContext struct {
	binary.Generate
	Hdc    HDC
	Result HGLRC
}
```

//////////////////////////////////////////////////////////////////////////////
WglCreateContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewWglCreateContext

```go
func NewWglCreateContext(
	pHdc HDC,
	pResult HGLRC,
) *WglCreateContext
```

#### func (*WglCreateContext) API

```go
func (c *WglCreateContext) API() gfxapi.API
```

#### func (*WglCreateContext) Class

```go
func (*WglCreateContext) Class() binary.Class
```

#### func (*WglCreateContext) Flags

```go
func (c *WglCreateContext) Flags() atom.Flags
```

#### func (*WglCreateContext) Mutate

```go
func (ϟa *WglCreateContext) Mutate(ϟs *gfxapi.State) error
```

#### func (*WglCreateContext) Replay

```go
func (ω *WglCreateContext) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*WglCreateContext) String

```go
func (c *WglCreateContext) String() string
```

#### func (*WglCreateContext) TypeID

```go
func (c *WglCreateContext) TypeID() atom.TypeID
```

#### type WglCreateContextAttribsARB

```go
type WglCreateContextAttribsARB struct {
	binary.Generate
	Hdc           HDC
	HShareContext HGLRC
	AttribList    IntArray
	Result        HGLRC
}
```

//////////////////////////////////////////////////////////////////////////////
WglCreateContextAttribsARB
//////////////////////////////////////////////////////////////////////////////

#### func  NewWglCreateContextAttribsARB

```go
func NewWglCreateContextAttribsARB(
	pHdc HDC,
	pHShareContext HGLRC,
	pAttribList IntArray,
	pResult HGLRC,
) *WglCreateContextAttribsARB
```

#### func (*WglCreateContextAttribsARB) API

```go
func (c *WglCreateContextAttribsARB) API() gfxapi.API
```

#### func (*WglCreateContextAttribsARB) Class

```go
func (*WglCreateContextAttribsARB) Class() binary.Class
```

#### func (*WglCreateContextAttribsARB) Flags

```go
func (c *WglCreateContextAttribsARB) Flags() atom.Flags
```

#### func (*WglCreateContextAttribsARB) Mutate

```go
func (ϟa *WglCreateContextAttribsARB) Mutate(ϟs *gfxapi.State) error
```

#### func (*WglCreateContextAttribsARB) Replay

```go
func (ω *WglCreateContextAttribsARB) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*WglCreateContextAttribsARB) String

```go
func (c *WglCreateContextAttribsARB) String() string
```

#### func (*WglCreateContextAttribsARB) TypeID

```go
func (c *WglCreateContextAttribsARB) TypeID() atom.TypeID
```

#### type WglCreateContextAttribsARB_Postback

```go
type WglCreateContextAttribsARB_Postback struct {
	Result []byte
}
```


#### func (*WglCreateContextAttribsARB_Postback) Decode

```go
func (o *WglCreateContextAttribsARB_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type WglCreateContext_Postback

```go
type WglCreateContext_Postback struct {
	Result []byte
}
```


#### func (*WglCreateContext_Postback) Decode

```go
func (o *WglCreateContext_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type WglMakeCurrent

```go
type WglMakeCurrent struct {
	binary.Generate
	Hdc    HDC
	Hglrc  HGLRC
	Result BOOL
}
```

//////////////////////////////////////////////////////////////////////////////
WglMakeCurrent
//////////////////////////////////////////////////////////////////////////////

#### func  NewWglMakeCurrent

```go
func NewWglMakeCurrent(
	pHdc HDC,
	pHglrc HGLRC,
	pResult BOOL,
) *WglMakeCurrent
```

#### func (*WglMakeCurrent) API

```go
func (c *WglMakeCurrent) API() gfxapi.API
```

#### func (*WglMakeCurrent) Class

```go
func (*WglMakeCurrent) Class() binary.Class
```

#### func (*WglMakeCurrent) Flags

```go
func (c *WglMakeCurrent) Flags() atom.Flags
```

#### func (*WglMakeCurrent) Mutate

```go
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State) error
```

#### func (*WglMakeCurrent) Replay

```go
func (ω *WglMakeCurrent) Replay(id atom.ID, gs *gfxapi.State, b *builder.Builder, wantOutput bool)
```

#### func (*WglMakeCurrent) String

```go
func (c *WglMakeCurrent) String() string
```

#### func (*WglMakeCurrent) TypeID

```go
func (c *WglMakeCurrent) TypeID() atom.TypeID
```

#### type WglMakeCurrent_Postback

```go
type WglMakeCurrent_Postback struct {
	Result BOOL
}
```


#### func (*WglMakeCurrent_Postback) Decode

```go
func (o *WglMakeCurrent_Postback) Decode(result_cnt uint64, d binary.Decoder) error
```

#### type WglSwapBuffers

```go
type WglSwapBuffers struct {
	binary.Generate
	Hdc HDC
}
```

//////////////////////////////////////////////////////////////////////////////
WglSwapBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewWglSwapBuffers

```go
func NewWglSwapBuffers(
	pHdc HDC,
) *WglSwapBuffers
```

#### func (*WglSwapBuffers) API

```go
func (c *WglSwapBuffers) API() gfxapi.API
```

#### func (*WglSwapBuffers) Class

```go
func (*WglSwapBuffers) Class() binary.Class
```

#### func (*WglSwapBuffers) Flags

```go
func (c *WglSwapBuffers) Flags() atom.Flags
```

#### func (*WglSwapBuffers) Mutate

```go
func (ϟa *WglSwapBuffers) Mutate(ϟs *gfxapi.State) error
```

#### func (*WglSwapBuffers) String

```go
func (c *WglSwapBuffers) String() string
```

#### func (*WglSwapBuffers) TypeID

```go
func (c *WglSwapBuffers) TypeID() atom.TypeID
```
