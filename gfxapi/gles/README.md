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
	CompressedTexelFormat_EXT_texture_compression_dxt1_GL_COMPRESSED_RGB_S3TC_DXT1_EXT  = CompressedTexelFormat_EXT_texture_compression_dxt1(33776)
	CompressedTexelFormat_EXT_texture_compression_dxt1_GL_COMPRESSED_RGBA_S3TC_DXT1_EXT = CompressedTexelFormat_EXT_texture_compression_dxt1(33777)
)
```

```go
const (
	CompressedTexelFormat_EXT_texture_compression_s3tc_GL_COMPRESSED_RGBA_S3TC_DXT3_EXT = CompressedTexelFormat_EXT_texture_compression_s3tc(33778)
	CompressedTexelFormat_EXT_texture_compression_s3tc_GL_COMPRESSED_RGBA_S3TC_DXT5_EXT = CompressedTexelFormat_EXT_texture_compression_s3tc(33779)
)
```

```go
const (
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_4x4_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37808)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_5x4_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37809)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_5x5_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37810)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_6x5_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37811)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_6x6_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37812)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_8x5_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37813)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_8x6_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37814)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_8x8_KHR           = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37815)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_10x5_KHR          = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37816)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_10x6_KHR          = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37817)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_10x8_KHR          = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37818)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_10x10_KHR         = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37819)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_12x10_KHR         = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37820)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_RGBA_ASTC_12x12_KHR         = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37821)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37840)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37841)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37842)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37843)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37844)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37845)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37846)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR   = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37847)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR  = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37848)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR  = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37849)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR  = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37850)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37851)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37852)
	CompressedTexelFormat_KHR_texture_compression_astc_ldr_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR = CompressedTexelFormat_KHR_texture_compression_astc_ldr(37853)
)
```

```go
const (
	CompressedTexelFormat_NV_texture_compression_latc_GL_COMPRESSED_LUMINANCE_LATC1_NV              = CompressedTexelFormat_NV_texture_compression_latc(35952)
	CompressedTexelFormat_NV_texture_compression_latc_GL_COMPRESSED_SIGNED_LUMINANCE_LATC1_NV       = CompressedTexelFormat_NV_texture_compression_latc(35953)
	CompressedTexelFormat_NV_texture_compression_latc_GL_COMPRESSED_LUMINANCE_ALPHA_LATC2_NV        = CompressedTexelFormat_NV_texture_compression_latc(35954)
	CompressedTexelFormat_NV_texture_compression_latc_GL_COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_NV = CompressedTexelFormat_NV_texture_compression_latc(35955)
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
	CompressedTexelFormat_GL_COMPRESSED_RGB_S3TC_DXT1_EXT  = CompressedTexelFormat(33776)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_S3TC_DXT1_EXT = CompressedTexelFormat(33777)
)
```
CompressedTexelFormat_EXT_texture_compression_dxt1

```go
const (
	CompressedTexelFormat_GL_COMPRESSED_RGBA_S3TC_DXT3_EXT = CompressedTexelFormat(33778)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_S3TC_DXT5_EXT = CompressedTexelFormat(33779)
)
```
CompressedTexelFormat_EXT_texture_compression_s3tc

```go
const (
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_4x4_KHR           = CompressedTexelFormat(37808)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_5x4_KHR           = CompressedTexelFormat(37809)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_5x5_KHR           = CompressedTexelFormat(37810)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_6x5_KHR           = CompressedTexelFormat(37811)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_6x6_KHR           = CompressedTexelFormat(37812)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_8x5_KHR           = CompressedTexelFormat(37813)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_8x6_KHR           = CompressedTexelFormat(37814)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_8x8_KHR           = CompressedTexelFormat(37815)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x5_KHR          = CompressedTexelFormat(37816)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x6_KHR          = CompressedTexelFormat(37817)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x8_KHR          = CompressedTexelFormat(37818)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x10_KHR         = CompressedTexelFormat(37819)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_12x10_KHR         = CompressedTexelFormat(37820)
	CompressedTexelFormat_GL_COMPRESSED_RGBA_ASTC_12x12_KHR         = CompressedTexelFormat(37821)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR   = CompressedTexelFormat(37840)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR   = CompressedTexelFormat(37841)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR   = CompressedTexelFormat(37842)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR   = CompressedTexelFormat(37843)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR   = CompressedTexelFormat(37844)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR   = CompressedTexelFormat(37845)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR   = CompressedTexelFormat(37846)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR   = CompressedTexelFormat(37847)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR  = CompressedTexelFormat(37848)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR  = CompressedTexelFormat(37849)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR  = CompressedTexelFormat(37850)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR = CompressedTexelFormat(37851)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR = CompressedTexelFormat(37852)
	CompressedTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR = CompressedTexelFormat(37853)
)
```
CompressedTexelFormat_KHR_texture_compression_astc_ldr

```go
const (
	CompressedTexelFormat_GL_COMPRESSED_LUMINANCE_LATC1_NV              = CompressedTexelFormat(35952)
	CompressedTexelFormat_GL_COMPRESSED_SIGNED_LUMINANCE_LATC1_NV       = CompressedTexelFormat(35953)
	CompressedTexelFormat_GL_COMPRESSED_LUMINANCE_ALPHA_LATC2_NV        = CompressedTexelFormat(35954)
	CompressedTexelFormat_GL_COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_NV = CompressedTexelFormat(35955)
)
```
CompressedTexelFormat_NV_texture_compression_latc

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
	ImageTexelFormat_GL_COMPRESSED_RGB_S3TC_DXT1_EXT  = ImageTexelFormat(33776)
	ImageTexelFormat_GL_COMPRESSED_RGBA_S3TC_DXT1_EXT = ImageTexelFormat(33777)
)
```
CompressedTexelFormat_EXT_texture_compression_dxt1

```go
const (
	ImageTexelFormat_GL_COMPRESSED_RGBA_S3TC_DXT3_EXT = ImageTexelFormat(33778)
	ImageTexelFormat_GL_COMPRESSED_RGBA_S3TC_DXT5_EXT = ImageTexelFormat(33779)
)
```
CompressedTexelFormat_EXT_texture_compression_s3tc

```go
const (
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_4x4_KHR           = ImageTexelFormat(37808)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_5x4_KHR           = ImageTexelFormat(37809)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_5x5_KHR           = ImageTexelFormat(37810)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_6x5_KHR           = ImageTexelFormat(37811)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_6x6_KHR           = ImageTexelFormat(37812)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_8x5_KHR           = ImageTexelFormat(37813)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_8x6_KHR           = ImageTexelFormat(37814)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_8x8_KHR           = ImageTexelFormat(37815)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x5_KHR          = ImageTexelFormat(37816)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x6_KHR          = ImageTexelFormat(37817)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x8_KHR          = ImageTexelFormat(37818)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_10x10_KHR         = ImageTexelFormat(37819)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_12x10_KHR         = ImageTexelFormat(37820)
	ImageTexelFormat_GL_COMPRESSED_RGBA_ASTC_12x12_KHR         = ImageTexelFormat(37821)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_4x4_KHR   = ImageTexelFormat(37840)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x4_KHR   = ImageTexelFormat(37841)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_5x5_KHR   = ImageTexelFormat(37842)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x5_KHR   = ImageTexelFormat(37843)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_6x6_KHR   = ImageTexelFormat(37844)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x5_KHR   = ImageTexelFormat(37845)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x6_KHR   = ImageTexelFormat(37846)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_8x8_KHR   = ImageTexelFormat(37847)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x5_KHR  = ImageTexelFormat(37848)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x6_KHR  = ImageTexelFormat(37849)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x8_KHR  = ImageTexelFormat(37850)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_10x10_KHR = ImageTexelFormat(37851)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x10_KHR = ImageTexelFormat(37852)
	ImageTexelFormat_GL_COMPRESSED_SRGB8_ALPHA8_ASTC_12x12_KHR = ImageTexelFormat(37853)
)
```
CompressedTexelFormat_KHR_texture_compression_astc_ldr

```go
const (
	ImageTexelFormat_GL_COMPRESSED_LUMINANCE_LATC1_NV              = ImageTexelFormat(35952)
	ImageTexelFormat_GL_COMPRESSED_SIGNED_LUMINANCE_LATC1_NV       = ImageTexelFormat(35953)
	ImageTexelFormat_GL_COMPRESSED_LUMINANCE_ALPHA_LATC2_NV        = ImageTexelFormat(35954)
	ImageTexelFormat_GL_COMPRESSED_SIGNED_LUMINANCE_ALPHA_LATC2_NV = ImageTexelFormat(35955)
)
```
CompressedTexelFormat_NV_texture_compression_latc

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
	UniformBlockParameter_GL_UNIFORM_BLOCK_BINDING                       = UniformBlockParameter(35391)
	UniformBlockParameter_GL_UNIFORM_BLOCK_DATA_SIZE                     = UniformBlockParameter(35392)
	UniformBlockParameter_GL_UNIFORM_BLOCK_NAME_LENGTH                   = UniformBlockParameter(35393)
	UniformBlockParameter_GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS               = UniformBlockParameter(35394)
	UniformBlockParameter_GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES        = UniformBlockParameter(35395)
	UniformBlockParameter_GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER   = UniformBlockParameter(35396)
	UniformBlockParameter_GL_UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER = UniformBlockParameter(35397)
	UniformBlockParameter_GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER = UniformBlockParameter(35398)
)
```

```go
const (
	IndexedBufferTarget_GL_TRANSFORM_FEEDBACK_BUFFER = IndexedBufferTarget(35982)
	IndexedBufferTarget_GL_UNIFORM_BUFFER            = IndexedBufferTarget(35345)
)
```

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
	ClientWaitSyncSignal_GL_ALREADY_SIGNALED    = ClientWaitSyncSignal(37146)
	ClientWaitSyncSignal_GL_TIMEOUT_EXPIRED     = ClientWaitSyncSignal(37147)
	ClientWaitSyncSignal_GL_CONDITION_SATISFIED = ClientWaitSyncSignal(37148)
	ClientWaitSyncSignal_GL_WAIT_FAILED         = ClientWaitSyncSignal(37149)
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
	SyncCondition_GL_SYNC_GPU_COMMANDS_COMPLETE = SyncCondition(37143)
)
```

```go
const (
	SyncFlags_GL_SYNC_FLUSH_COMMANDS_BIT = SyncFlags(1)
)
```

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
	Type_ARB_half_float_vertex_GL_HALF_FLOAT_ARB = Type_ARB_half_float_vertex(5131)
)
```

```go
const (
	Type_OES_vertex_half_float_GL_HALF_FLOAT_OES = Type_OES_vertex_half_float(36193)
)
```

```go
const (
	VertexAttribType_GL_HALF_FLOAT_ARB = VertexAttribType(5131)
)
```
Type_ARB_half_float_vertex

```go
const (
	VertexAttribType_GL_HALF_FLOAT_OES = VertexAttribType(36193)
)
```
Type_OES_vertex_half_float

```go
var ConstantValues schema.Constants
```

```go
var Namespace = registry.NewNamespace()
```

#### func  API

```go
func API() gfxapi.API
```

#### func  NewProgram

```go
func NewProgram(a device.Architecture, d database.Database, l log.Logger,
	vertexShaderID, fragmentShaderID ShaderId, programID ProgramId,
	vertexShaderSource, fragmentShaderSource string) []atom.Atom
```
NewProgram returns the atoms to create a shader program with compiled vertex and
fragment shaders. The returned program is not linked.

#### type Architecture

```go
type Architecture struct {
	binary.Generate

	PointerAlignment uint32
	PointerSize      uint32
	IntegerSize      uint32
	LittleEndian     bool
}
```

//////////////////////////////////////////////////////////////////////////////
Architecture
//////////////////////////////////////////////////////////////////////////////

#### func  NewArchitecture

```go
func NewArchitecture(Pointer_alignment uint32, Pointer_size uint32, Integer_size uint32, Little_endian bool) *Architecture
```

#### func (*Architecture) API

```go
func (c *Architecture) API() gfxapi.ID
```

#### func (*Architecture) AddRead

```go
func (a *Architecture) AddRead(rng memory.Range, id binary.ID) *Architecture
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The Architecture pointer is returned so that calls can be chained.

#### func (*Architecture) AddWrite

```go
func (a *Architecture) AddWrite(rng memory.Range, id binary.ID) *Architecture
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The Architecture pointer is returned so that calls can be chained.

#### func (*Architecture) Class

```go
func (*Architecture) Class() binary.Class
```

#### func (*Architecture) Flags

```go
func (c *Architecture) Flags() atom.Flags
```

#### func (*Architecture) Mutate

```go
func (ϟa *Architecture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*Architecture) Observations

```go
func (a *Architecture) Observations() *atom.Observations
```

#### func (*Architecture) String

```go
func (a *Architecture) String() string
```

#### type ArrayType

```go
type ArrayType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ArrayType
//////////////////////////////////////////////////////////////////////////////

#### func (*ArrayType) Parse

```go
func (v *ArrayType) Parse(s string) error
```

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

#### func (*ArrayType_GLES_1_1) Parse

```go
func (v *ArrayType_GLES_1_1) Parse(s string) error
```

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

#### func (*ArrayType_OES_point_size_array) Parse

```go
func (v *ArrayType_OES_point_size_array) Parse(s string) error
```

#### func (ArrayType_OES_point_size_array) String

```go
func (v ArrayType_OES_point_size_array) String() string
```

#### type AttributeLocation

```go
type AttributeLocation int32
```


#### type AttributeLocationːVertexAttributeArrayʳᵐ

```go
type AttributeLocationːVertexAttributeArrayʳᵐ map[AttributeLocation](*VertexAttributeArray)
```


#### func (AttributeLocationːVertexAttributeArrayʳᵐ) Contains

```go
func (m AttributeLocationːVertexAttributeArrayʳᵐ) Contains(key AttributeLocation) bool
```

#### func (AttributeLocationːVertexAttributeArrayʳᵐ) Delete

```go
func (m AttributeLocationːVertexAttributeArrayʳᵐ) Delete(key AttributeLocation)
```

#### func (AttributeLocationːVertexAttributeArrayʳᵐ) Get

```go
func (m AttributeLocationːVertexAttributeArrayʳᵐ) Get(key AttributeLocation) *VertexAttributeArray
```

#### func (AttributeLocationːVertexAttributeArrayʳᵐ) Range

```go
func (m AttributeLocationːVertexAttributeArrayʳᵐ) Range() [](*VertexAttributeArray)
```

#### type BOOL

```go
type BOOL int64
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
func NewBackbufferInfo(Width int32, Height int32, Color_fmt RenderbufferFormat, Depth_fmt RenderbufferFormat, Stencil_fmt RenderbufferFormat, ResetViewportScissor bool) *BackbufferInfo
```

#### func (*BackbufferInfo) API

```go
func (c *BackbufferInfo) API() gfxapi.ID
```

#### func (*BackbufferInfo) AddRead

```go
func (a *BackbufferInfo) AddRead(rng memory.Range, id binary.ID) *BackbufferInfo
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The BackbufferInfo pointer is returned so that calls can be chained.

#### func (*BackbufferInfo) AddWrite

```go
func (a *BackbufferInfo) AddWrite(rng memory.Range, id binary.ID) *BackbufferInfo
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The BackbufferInfo pointer is returned so that calls can be chained.

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
func (ϟa *BackbufferInfo) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*BackbufferInfo) Observations

```go
func (a *BackbufferInfo) Observations() *atom.Observations
```

#### func (*BackbufferInfo) Replay

```go
func (ϟa *BackbufferInfo) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*BackbufferInfo) String

```go
func (a *BackbufferInfo) String() string
```

#### type BaseTexelFormat

```go
type BaseTexelFormat uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BaseTexelFormat
//////////////////////////////////////////////////////////////////////////////

#### func (*BaseTexelFormat) Parse

```go
func (v *BaseTexelFormat) Parse(s string) error
```

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

#### func (*BlendEquation) Parse

```go
func (v *BlendEquation) Parse(s string) error
```

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

#### func (*BlendFactor) Parse

```go
func (v *BlendFactor) Parse(s string) error
```

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

#### type Bool

```go
type Bool int64
```


#### type Boolˢ

```go
type Boolˢ struct {
	binary.Generate
	SliceInfo
}
```

Boolˢ is a slice of bool.

#### func  AsBoolˢ

```go
func AsBoolˢ(s Slice, ϟs *gfxapi.State) Boolˢ
```
AsBoolˢ returns s cast to a Boolˢ. The returned slice length will be calculated
so that the returned slice is no longer (in bytes) than s.

#### func  MakeBoolˢ

```go
func MakeBoolˢ(count uint64, ϟs *gfxapi.State) Boolˢ
```
MakeBoolˢ returns a Boolˢ backed by a new memory pool.

#### func (*Boolˢ) Class

```go
func (*Boolˢ) Class() binary.Class
```

#### func (Boolˢ) Clone

```go
func (s Boolˢ) Clone(ϟs *gfxapi.State) Boolˢ
```
Clone returns a copy of the Boolˢ in a new memory pool.

#### func (Boolˢ) Copy

```go
func (dst Boolˢ) Copy(src Boolˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Boolˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Boolˢ) Decoder

```go
func (s Boolˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Boolˢ) ElementSize

```go
func (s Boolˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Boolˢ points to.

#### func (Boolˢ) Encoder

```go
func (s Boolˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Boolˢ) Index

```go
func (s Boolˢ) Index(i uint64, ϟs *gfxapi.State) Boolᵖ
```
Index returns a Boolᵖ to the i'th element in this Boolˢ.

#### func (Boolˢ) OnRead

```go
func (s Boolˢ) OnRead(ϟs *gfxapi.State) Boolˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Boolˢ) OnWrite

```go
func (s Boolˢ) OnWrite(ϟs *gfxapi.State) Boolˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Boolˢ) Range

```go
func (s Boolˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Boolˢ) Read

```go
func (s Boolˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []bool
```
Read reads and returns all the bool elements in this Boolˢ.

#### func (Boolˢ) ResourceID

```go
func (s Boolˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Boolˢ) Slice

```go
func (s Boolˢ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ
```
Slice returns a sub-slice from the Boolˢ using start and end indices.

#### func (Boolˢ) String

```go
func (s Boolˢ) String() string
```
String returns a string description of the Boolˢ slice.

#### func (Boolˢ) Write

```go
func (s Boolˢ) Write(src []bool, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Boolᵖ

```go
type Boolᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Boolᵖ is a pointer to a bool element.

#### func  NewBoolᵖ

```go
func NewBoolᵖ(addr uint64) Boolᵖ
```
NewBoolᵖ returns a Boolᵖ that points to addr in the application pool.

#### func (*Boolᵖ) Class

```go
func (*Boolᵖ) Class() binary.Class
```

#### func (Boolᵖ) ElementSize

```go
func (p Boolᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Boolᵖ points to.

#### func (Boolᵖ) OnRead

```go
func (p Boolᵖ) OnRead(ϟs *gfxapi.State) Boolᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Boolᵖ) OnWrite

```go
func (p Boolᵖ) OnWrite(ϟs *gfxapi.State) Boolᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Boolᵖ) Read

```go
func (p Boolᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) bool
```
Read reads and returns the bool element at the pointer.

#### func (Boolᵖ) Slice

```go
func (p Boolᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Boolˢ
```
Slice returns a new Boolˢ from the pointer using start and end indices.

#### func (Boolᵖ) Write

```go
func (p Boolᵖ) Write(value bool, ϟs *gfxapi.State)
```
Write writes value to the bool element at the pointer.

#### type Buffer

```go
type Buffer struct {
	binary.Generate
	CreatedAt atom.ID
	Data      U8ˢ
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

#### type BufferDataPointer

```go
type BufferDataPointer struct {
	binary.Generate
	memory.Pointer
}
```

BufferDataPointer is a pointer to a void element.

#### func  NewBufferDataPointer

```go
func NewBufferDataPointer(addr uint64) BufferDataPointer
```
NewBufferDataPointer returns a BufferDataPointer that points to addr in the
application pool.

#### func (*BufferDataPointer) Class

```go
func (*BufferDataPointer) Class() binary.Class
```

#### func (BufferDataPointer) ElementSize

```go
func (p BufferDataPointer) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that BufferDataPointer
points to.

#### func (BufferDataPointer) OnRead

```go
func (p BufferDataPointer) OnRead(ϟs *gfxapi.State) BufferDataPointer
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (BufferDataPointer) OnWrite

```go
func (p BufferDataPointer) OnWrite(ϟs *gfxapi.State) BufferDataPointer
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (BufferDataPointer) Slice

```go
func (p BufferDataPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type BufferId

```go
type BufferId uint32
```


#### type BufferIdːBufferʳᵐ

```go
type BufferIdːBufferʳᵐ map[BufferId](*Buffer)
```


#### func (BufferIdːBufferʳᵐ) Contains

```go
func (m BufferIdːBufferʳᵐ) Contains(key BufferId) bool
```

#### func (BufferIdːBufferʳᵐ) Delete

```go
func (m BufferIdːBufferʳᵐ) Delete(key BufferId)
```

#### func (BufferIdːBufferʳᵐ) Get

```go
func (m BufferIdːBufferʳᵐ) Get(key BufferId) *Buffer
```

#### func (BufferIdːBufferʳᵐ) Range

```go
func (m BufferIdːBufferʳᵐ) Range() [](*Buffer)
```

#### type BufferIdˢ

```go
type BufferIdˢ struct {
	binary.Generate
	SliceInfo
}
```

BufferIdˢ is a slice of BufferId.

#### func  AsBufferIdˢ

```go
func AsBufferIdˢ(s Slice, ϟs *gfxapi.State) BufferIdˢ
```
AsBufferIdˢ returns s cast to a BufferIdˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeBufferIdˢ

```go
func MakeBufferIdˢ(count uint64, ϟs *gfxapi.State) BufferIdˢ
```
MakeBufferIdˢ returns a BufferIdˢ backed by a new memory pool.

#### func (*BufferIdˢ) Class

```go
func (*BufferIdˢ) Class() binary.Class
```

#### func (BufferIdˢ) Clone

```go
func (s BufferIdˢ) Clone(ϟs *gfxapi.State) BufferIdˢ
```
Clone returns a copy of the BufferIdˢ in a new memory pool.

#### func (BufferIdˢ) Copy

```go
func (dst BufferIdˢ) Copy(src BufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s BufferIdˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (BufferIdˢ) Decoder

```go
func (s BufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (BufferIdˢ) ElementSize

```go
func (s BufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that BufferIdˢ points to.

#### func (BufferIdˢ) Encoder

```go
func (s BufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (BufferIdˢ) Index

```go
func (s BufferIdˢ) Index(i uint64, ϟs *gfxapi.State) BufferIdᵖ
```
Index returns a BufferIdᵖ to the i'th element in this BufferIdˢ.

#### func (BufferIdˢ) OnRead

```go
func (s BufferIdˢ) OnRead(ϟs *gfxapi.State) BufferIdˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (BufferIdˢ) OnWrite

```go
func (s BufferIdˢ) OnWrite(ϟs *gfxapi.State) BufferIdˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (BufferIdˢ) Range

```go
func (s BufferIdˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (BufferIdˢ) Read

```go
func (s BufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []BufferId
```
Read reads and returns all the BufferId elements in this BufferIdˢ.

#### func (BufferIdˢ) ResourceID

```go
func (s BufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (BufferIdˢ) Slice

```go
func (s BufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ
```
Slice returns a sub-slice from the BufferIdˢ using start and end indices.

#### func (BufferIdˢ) String

```go
func (s BufferIdˢ) String() string
```
String returns a string description of the BufferIdˢ slice.

#### func (BufferIdˢ) Write

```go
func (s BufferIdˢ) Write(src []BufferId, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type BufferIdᵖ

```go
type BufferIdᵖ struct {
	binary.Generate
	memory.Pointer
}
```

BufferIdᵖ is a pointer to a BufferId element.

#### func  NewBufferIdᵖ

```go
func NewBufferIdᵖ(addr uint64) BufferIdᵖ
```
NewBufferIdᵖ returns a BufferIdᵖ that points to addr in the application pool.

#### func (*BufferIdᵖ) Class

```go
func (*BufferIdᵖ) Class() binary.Class
```

#### func (BufferIdᵖ) ElementSize

```go
func (p BufferIdᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that BufferIdᵖ points to.

#### func (BufferIdᵖ) OnRead

```go
func (p BufferIdᵖ) OnRead(ϟs *gfxapi.State) BufferIdᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (BufferIdᵖ) OnWrite

```go
func (p BufferIdᵖ) OnWrite(ϟs *gfxapi.State) BufferIdᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (BufferIdᵖ) Read

```go
func (p BufferIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) BufferId
```
Read reads and returns the BufferId element at the pointer.

#### func (BufferIdᵖ) Slice

```go
func (p BufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ
```
Slice returns a new BufferIdˢ from the pointer using start and end indices.

#### func (BufferIdᵖ) Write

```go
func (p BufferIdᵖ) Write(value BufferId, ϟs *gfxapi.State)
```
Write writes value to the BufferId element at the pointer.

#### type BufferIdᶜᵖ

```go
type BufferIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

BufferIdᶜᵖ is a pointer to a BufferId element.

#### func  NewBufferIdᶜᵖ

```go
func NewBufferIdᶜᵖ(addr uint64) BufferIdᶜᵖ
```
NewBufferIdᶜᵖ returns a BufferIdᶜᵖ that points to addr in the application pool.

#### func (*BufferIdᶜᵖ) Class

```go
func (*BufferIdᶜᵖ) Class() binary.Class
```

#### func (BufferIdᶜᵖ) ElementSize

```go
func (p BufferIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that BufferIdᶜᵖ points to.

#### func (BufferIdᶜᵖ) OnRead

```go
func (p BufferIdᶜᵖ) OnRead(ϟs *gfxapi.State) BufferIdᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (BufferIdᶜᵖ) OnWrite

```go
func (p BufferIdᶜᵖ) OnWrite(ϟs *gfxapi.State) BufferIdᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (BufferIdᶜᵖ) Read

```go
func (p BufferIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) BufferId
```
Read reads and returns the BufferId element at the pointer.

#### func (BufferIdᶜᵖ) Slice

```go
func (p BufferIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) BufferIdˢ
```
Slice returns a new BufferIdˢ from the pointer using start and end indices.

#### func (BufferIdᶜᵖ) Write

```go
func (p BufferIdᶜᵖ) Write(value BufferId, ϟs *gfxapi.State)
```
Write writes value to the BufferId element at the pointer.

#### type BufferParameter

```go
type BufferParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BufferParameter
//////////////////////////////////////////////////////////////////////////////

#### func (*BufferParameter) Parse

```go
func (v *BufferParameter) Parse(s string) error
```

#### func (BufferParameter) String

```go
func (v BufferParameter) String() string
```

#### type BufferTarget

```go
type BufferTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BufferTarget
//////////////////////////////////////////////////////////////////////////////

#### func (*BufferTarget) Parse

```go
func (v *BufferTarget) Parse(s string) error
```

#### func (BufferTarget) String

```go
func (v BufferTarget) String() string
```

#### type BufferTargetːBufferIdᵐ

```go
type BufferTargetːBufferIdᵐ map[BufferTarget]BufferId
```


#### func (BufferTargetːBufferIdᵐ) Contains

```go
func (m BufferTargetːBufferIdᵐ) Contains(key BufferTarget) bool
```

#### func (BufferTargetːBufferIdᵐ) Delete

```go
func (m BufferTargetːBufferIdᵐ) Delete(key BufferTarget)
```

#### func (BufferTargetːBufferIdᵐ) Get

```go
func (m BufferTargetːBufferIdᵐ) Get(key BufferTarget) BufferId
```

#### func (BufferTargetːBufferIdᵐ) Range

```go
func (m BufferTargetːBufferIdᵐ) Range() []BufferId
```

#### type BufferUsage

```go
type BufferUsage uint32
```

//////////////////////////////////////////////////////////////////////////////
enum BufferUsage
//////////////////////////////////////////////////////////////////////////////

#### func (*BufferUsage) Parse

```go
func (v *BufferUsage) Parse(s string) error
```

#### func (BufferUsage) String

```go
func (v BufferUsage) String() string
```

#### type CGLContextObj

```go
type CGLContextObj struct {
	binary.Generate
	memory.Pointer
}
```

CGLContextObj is a pointer to a void element.

#### func  NewCGLContextObj

```go
func NewCGLContextObj(addr uint64) CGLContextObj
```
NewCGLContextObj returns a CGLContextObj that points to addr in the application
pool.

#### func (*CGLContextObj) Class

```go
func (*CGLContextObj) Class() binary.Class
```

#### func (CGLContextObj) ElementSize

```go
func (p CGLContextObj) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGLContextObj points
to.

#### func (CGLContextObj) OnRead

```go
func (p CGLContextObj) OnRead(ϟs *gfxapi.State) CGLContextObj
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (CGLContextObj) OnWrite

```go
func (p CGLContextObj) OnWrite(ϟs *gfxapi.State) CGLContextObj
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (CGLContextObj) Slice

```go
func (p CGLContextObj) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type CGLContextObjːContextʳᵐ

```go
type CGLContextObjːContextʳᵐ map[CGLContextObj](*Context)
```


#### func (CGLContextObjːContextʳᵐ) Contains

```go
func (m CGLContextObjːContextʳᵐ) Contains(key CGLContextObj) bool
```

#### func (CGLContextObjːContextʳᵐ) Delete

```go
func (m CGLContextObjːContextʳᵐ) Delete(key CGLContextObj)
```

#### func (CGLContextObjːContextʳᵐ) Get

```go
func (m CGLContextObjːContextʳᵐ) Get(key CGLContextObj) *Context
```

#### func (CGLContextObjːContextʳᵐ) Range

```go
func (m CGLContextObjːContextʳᵐ) Range() [](*Context)
```

#### type CGLContextObjˢ

```go
type CGLContextObjˢ struct {
	binary.Generate
	SliceInfo
}
```

CGLContextObjˢ is a slice of CGLContextObj.

#### func  AsCGLContextObjˢ

```go
func AsCGLContextObjˢ(s Slice, ϟs *gfxapi.State) CGLContextObjˢ
```
AsCGLContextObjˢ returns s cast to a CGLContextObjˢ. The returned slice length
will be calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeCGLContextObjˢ

```go
func MakeCGLContextObjˢ(count uint64, ϟs *gfxapi.State) CGLContextObjˢ
```
MakeCGLContextObjˢ returns a CGLContextObjˢ backed by a new memory pool.

#### func (*CGLContextObjˢ) Class

```go
func (*CGLContextObjˢ) Class() binary.Class
```

#### func (CGLContextObjˢ) Clone

```go
func (s CGLContextObjˢ) Clone(ϟs *gfxapi.State) CGLContextObjˢ
```
Clone returns a copy of the CGLContextObjˢ in a new memory pool.

#### func (CGLContextObjˢ) Copy

```go
func (dst CGLContextObjˢ) Copy(src CGLContextObjˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGLContextObjˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (CGLContextObjˢ) Decoder

```go
func (s CGLContextObjˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (CGLContextObjˢ) ElementSize

```go
func (s CGLContextObjˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGLContextObjˢ points
to.

#### func (CGLContextObjˢ) Encoder

```go
func (s CGLContextObjˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (CGLContextObjˢ) Index

```go
func (s CGLContextObjˢ) Index(i uint64, ϟs *gfxapi.State) CGLContextObjᵖ
```
Index returns a CGLContextObjᵖ to the i'th element in this CGLContextObjˢ.

#### func (CGLContextObjˢ) OnRead

```go
func (s CGLContextObjˢ) OnRead(ϟs *gfxapi.State) CGLContextObjˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (CGLContextObjˢ) OnWrite

```go
func (s CGLContextObjˢ) OnWrite(ϟs *gfxapi.State) CGLContextObjˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (CGLContextObjˢ) Range

```go
func (s CGLContextObjˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (CGLContextObjˢ) Read

```go
func (s CGLContextObjˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGLContextObj
```
Read reads and returns all the CGLContextObj elements in this CGLContextObjˢ.

#### func (CGLContextObjˢ) ResourceID

```go
func (s CGLContextObjˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (CGLContextObjˢ) Slice

```go
func (s CGLContextObjˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGLContextObjˢ
```
Slice returns a sub-slice from the CGLContextObjˢ using start and end indices.

#### func (CGLContextObjˢ) String

```go
func (s CGLContextObjˢ) String() string
```
String returns a string description of the CGLContextObjˢ slice.

#### func (CGLContextObjˢ) Write

```go
func (s CGLContextObjˢ) Write(src []CGLContextObj, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type CGLContextObjᵖ

```go
type CGLContextObjᵖ struct {
	binary.Generate
	memory.Pointer
}
```

CGLContextObjᵖ is a pointer to a CGLContextObj element. Note: Pointers are
stored differently between the application pool and internal pools.

    * The application pool stores pointers as an address of an architecture-dependant size.
    * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
      pool identifier.

#### func  NewCGLContextObjᵖ

```go
func NewCGLContextObjᵖ(addr uint64) CGLContextObjᵖ
```
NewCGLContextObjᵖ returns a CGLContextObjᵖ that points to addr in the
application pool.

#### func (*CGLContextObjᵖ) Class

```go
func (*CGLContextObjᵖ) Class() binary.Class
```

#### func (CGLContextObjᵖ) ElementSize

```go
func (p CGLContextObjᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGLContextObjᵖ points
to.

#### func (CGLContextObjᵖ) OnRead

```go
func (p CGLContextObjᵖ) OnRead(ϟs *gfxapi.State) CGLContextObjᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (CGLContextObjᵖ) OnWrite

```go
func (p CGLContextObjᵖ) OnWrite(ϟs *gfxapi.State) CGLContextObjᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (CGLContextObjᵖ) Read

```go
func (p CGLContextObjᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGLContextObj
```
Read reads and returns the CGLContextObj element at the pointer.

#### func (CGLContextObjᵖ) Slice

```go
func (p CGLContextObjᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGLContextObjˢ
```
Slice returns a new CGLContextObjˢ from the pointer using start and end indices.

#### func (CGLContextObjᵖ) Write

```go
func (p CGLContextObjᵖ) Write(value CGLContextObj, ϟs *gfxapi.State)
```
Write writes value to the CGLContextObj element at the pointer.

#### type CGLCreateContext

```go
type CGLCreateContext struct {
	binary.Generate

	Pix    CGLPixelFormatObj
	Share  CGLContextObj
	Ctx    CGLContextObjᵖ
	Result CGLError
}
```

//////////////////////////////////////////////////////////////////////////////
CGLCreateContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewCGLCreateContext

```go
func NewCGLCreateContext(Pix memory.Pointer, Share memory.Pointer, Ctx memory.Pointer, Result CGLError) *CGLCreateContext
```

#### func (*CGLCreateContext) API

```go
func (c *CGLCreateContext) API() gfxapi.ID
```

#### func (*CGLCreateContext) AddRead

```go
func (a *CGLCreateContext) AddRead(rng memory.Range, id binary.ID) *CGLCreateContext
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CGLCreateContext pointer is returned so that calls can be chained.

#### func (*CGLCreateContext) AddWrite

```go
func (a *CGLCreateContext) AddWrite(rng memory.Range, id binary.ID) *CGLCreateContext
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CGLCreateContext pointer is returned so that calls can be chained.

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
func (ϟa *CGLCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CGLCreateContext) Observations

```go
func (a *CGLCreateContext) Observations() *atom.Observations
```

#### func (*CGLCreateContext) Replay

```go
func (ω *CGLCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*CGLCreateContext) String

```go
func (a *CGLCreateContext) String() string
```

#### type CGLError

```go
type CGLError int64
```


#### type CGLFlushDrawable

```go
type CGLFlushDrawable struct {
	binary.Generate

	Ctx    CGLContextObj
	Result CGLError
}
```

//////////////////////////////////////////////////////////////////////////////
CGLFlushDrawable
//////////////////////////////////////////////////////////////////////////////

#### func  NewCGLFlushDrawable

```go
func NewCGLFlushDrawable(Ctx memory.Pointer, Result CGLError) *CGLFlushDrawable
```

#### func (*CGLFlushDrawable) API

```go
func (c *CGLFlushDrawable) API() gfxapi.ID
```

#### func (*CGLFlushDrawable) AddRead

```go
func (a *CGLFlushDrawable) AddRead(rng memory.Range, id binary.ID) *CGLFlushDrawable
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CGLFlushDrawable pointer is returned so that calls can be chained.

#### func (*CGLFlushDrawable) AddWrite

```go
func (a *CGLFlushDrawable) AddWrite(rng memory.Range, id binary.ID) *CGLFlushDrawable
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CGLFlushDrawable pointer is returned so that calls can be chained.

#### func (*CGLFlushDrawable) Class

```go
func (*CGLFlushDrawable) Class() binary.Class
```

#### func (*CGLFlushDrawable) Flags

```go
func (c *CGLFlushDrawable) Flags() atom.Flags
```

#### func (*CGLFlushDrawable) Mutate

```go
func (ϟa *CGLFlushDrawable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CGLFlushDrawable) Observations

```go
func (a *CGLFlushDrawable) Observations() *atom.Observations
```

#### func (*CGLFlushDrawable) String

```go
func (a *CGLFlushDrawable) String() string
```

#### type CGLGetSurface

```go
type CGLGetSurface struct {
	binary.Generate

	Ctx    CGLContextObj
	Cid    CGSConnectionIDᵖ
	Wid    CGSWindowIDᵖ
	Sid    CGSSurfaceIDᵖ
	Result int64
}
```

//////////////////////////////////////////////////////////////////////////////
CGLGetSurface
//////////////////////////////////////////////////////////////////////////////

#### func  NewCGLGetSurface

```go
func NewCGLGetSurface(Ctx memory.Pointer, Cid memory.Pointer, Wid memory.Pointer, Sid memory.Pointer, Result int64) *CGLGetSurface
```

#### func (*CGLGetSurface) API

```go
func (c *CGLGetSurface) API() gfxapi.ID
```

#### func (*CGLGetSurface) AddRead

```go
func (a *CGLGetSurface) AddRead(rng memory.Range, id binary.ID) *CGLGetSurface
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CGLGetSurface pointer is returned so that calls can be chained.

#### func (*CGLGetSurface) AddWrite

```go
func (a *CGLGetSurface) AddWrite(rng memory.Range, id binary.ID) *CGLGetSurface
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CGLGetSurface pointer is returned so that calls can be chained.

#### func (*CGLGetSurface) Class

```go
func (*CGLGetSurface) Class() binary.Class
```

#### func (*CGLGetSurface) Flags

```go
func (c *CGLGetSurface) Flags() atom.Flags
```

#### func (*CGLGetSurface) Mutate

```go
func (ϟa *CGLGetSurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CGLGetSurface) Observations

```go
func (a *CGLGetSurface) Observations() *atom.Observations
```

#### func (*CGLGetSurface) String

```go
func (a *CGLGetSurface) String() string
```

#### type CGLPixelFormatObj

```go
type CGLPixelFormatObj struct {
	binary.Generate
	memory.Pointer
}
```

CGLPixelFormatObj is a pointer to a void element.

#### func  NewCGLPixelFormatObj

```go
func NewCGLPixelFormatObj(addr uint64) CGLPixelFormatObj
```
NewCGLPixelFormatObj returns a CGLPixelFormatObj that points to addr in the
application pool.

#### func (*CGLPixelFormatObj) Class

```go
func (*CGLPixelFormatObj) Class() binary.Class
```

#### func (CGLPixelFormatObj) ElementSize

```go
func (p CGLPixelFormatObj) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGLPixelFormatObj
points to.

#### func (CGLPixelFormatObj) OnRead

```go
func (p CGLPixelFormatObj) OnRead(ϟs *gfxapi.State) CGLPixelFormatObj
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (CGLPixelFormatObj) OnWrite

```go
func (p CGLPixelFormatObj) OnWrite(ϟs *gfxapi.State) CGLPixelFormatObj
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (CGLPixelFormatObj) Slice

```go
func (p CGLPixelFormatObj) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

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
func NewCGLSetCurrentContext(Ctx memory.Pointer, Result CGLError) *CGLSetCurrentContext
```

#### func (*CGLSetCurrentContext) API

```go
func (c *CGLSetCurrentContext) API() gfxapi.ID
```

#### func (*CGLSetCurrentContext) AddRead

```go
func (a *CGLSetCurrentContext) AddRead(rng memory.Range, id binary.ID) *CGLSetCurrentContext
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CGLSetCurrentContext pointer is returned so that calls can be
chained.

#### func (*CGLSetCurrentContext) AddWrite

```go
func (a *CGLSetCurrentContext) AddWrite(rng memory.Range, id binary.ID) *CGLSetCurrentContext
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CGLSetCurrentContext pointer is returned so that calls can be
chained.

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
func (ϟa *CGLSetCurrentContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CGLSetCurrentContext) Observations

```go
func (a *CGLSetCurrentContext) Observations() *atom.Observations
```

#### func (*CGLSetCurrentContext) Replay

```go
func (ω *CGLSetCurrentContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*CGLSetCurrentContext) String

```go
func (a *CGLSetCurrentContext) String() string
```

#### type CGSConnectionID

```go
type CGSConnectionID struct {
	binary.Generate
	memory.Pointer
}
```

CGSConnectionID is a pointer to a void element.

#### func  NewCGSConnectionID

```go
func NewCGSConnectionID(addr uint64) CGSConnectionID
```
NewCGSConnectionID returns a CGSConnectionID that points to addr in the
application pool.

#### func (*CGSConnectionID) Class

```go
func (*CGSConnectionID) Class() binary.Class
```

#### func (CGSConnectionID) ElementSize

```go
func (p CGSConnectionID) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGSConnectionID points
to.

#### func (CGSConnectionID) OnRead

```go
func (p CGSConnectionID) OnRead(ϟs *gfxapi.State) CGSConnectionID
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (CGSConnectionID) OnWrite

```go
func (p CGSConnectionID) OnWrite(ϟs *gfxapi.State) CGSConnectionID
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (CGSConnectionID) Slice

```go
func (p CGSConnectionID) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type CGSConnectionIDˢ

```go
type CGSConnectionIDˢ struct {
	binary.Generate
	SliceInfo
}
```

CGSConnectionIDˢ is a slice of CGSConnectionID.

#### func  AsCGSConnectionIDˢ

```go
func AsCGSConnectionIDˢ(s Slice, ϟs *gfxapi.State) CGSConnectionIDˢ
```
AsCGSConnectionIDˢ returns s cast to a CGSConnectionIDˢ. The returned slice
length will be calculated so that the returned slice is no longer (in bytes)
than s.

#### func  MakeCGSConnectionIDˢ

```go
func MakeCGSConnectionIDˢ(count uint64, ϟs *gfxapi.State) CGSConnectionIDˢ
```
MakeCGSConnectionIDˢ returns a CGSConnectionIDˢ backed by a new memory pool.

#### func (*CGSConnectionIDˢ) Class

```go
func (*CGSConnectionIDˢ) Class() binary.Class
```

#### func (CGSConnectionIDˢ) Clone

```go
func (s CGSConnectionIDˢ) Clone(ϟs *gfxapi.State) CGSConnectionIDˢ
```
Clone returns a copy of the CGSConnectionIDˢ in a new memory pool.

#### func (CGSConnectionIDˢ) Copy

```go
func (dst CGSConnectionIDˢ) Copy(src CGSConnectionIDˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGSConnectionIDˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (CGSConnectionIDˢ) Decoder

```go
func (s CGSConnectionIDˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (CGSConnectionIDˢ) ElementSize

```go
func (s CGSConnectionIDˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGSConnectionIDˢ points
to.

#### func (CGSConnectionIDˢ) Encoder

```go
func (s CGSConnectionIDˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (CGSConnectionIDˢ) Index

```go
func (s CGSConnectionIDˢ) Index(i uint64, ϟs *gfxapi.State) CGSConnectionIDᵖ
```
Index returns a CGSConnectionIDᵖ to the i'th element in this CGSConnectionIDˢ.

#### func (CGSConnectionIDˢ) OnRead

```go
func (s CGSConnectionIDˢ) OnRead(ϟs *gfxapi.State) CGSConnectionIDˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (CGSConnectionIDˢ) OnWrite

```go
func (s CGSConnectionIDˢ) OnWrite(ϟs *gfxapi.State) CGSConnectionIDˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (CGSConnectionIDˢ) Range

```go
func (s CGSConnectionIDˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (CGSConnectionIDˢ) Read

```go
func (s CGSConnectionIDˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGSConnectionID
```
Read reads and returns all the CGSConnectionID elements in this
CGSConnectionIDˢ.

#### func (CGSConnectionIDˢ) ResourceID

```go
func (s CGSConnectionIDˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (CGSConnectionIDˢ) Slice

```go
func (s CGSConnectionIDˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGSConnectionIDˢ
```
Slice returns a sub-slice from the CGSConnectionIDˢ using start and end indices.

#### func (CGSConnectionIDˢ) String

```go
func (s CGSConnectionIDˢ) String() string
```
String returns a string description of the CGSConnectionIDˢ slice.

#### func (CGSConnectionIDˢ) Write

```go
func (s CGSConnectionIDˢ) Write(src []CGSConnectionID, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type CGSConnectionIDᵖ

```go
type CGSConnectionIDᵖ struct {
	binary.Generate
	memory.Pointer
}
```

CGSConnectionIDᵖ is a pointer to a CGSConnectionID element. Note: Pointers are
stored differently between the application pool and internal pools.

    * The application pool stores pointers as an address of an architecture-dependant size.
    * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
      pool identifier.

#### func  NewCGSConnectionIDᵖ

```go
func NewCGSConnectionIDᵖ(addr uint64) CGSConnectionIDᵖ
```
NewCGSConnectionIDᵖ returns a CGSConnectionIDᵖ that points to addr in the
application pool.

#### func (*CGSConnectionIDᵖ) Class

```go
func (*CGSConnectionIDᵖ) Class() binary.Class
```

#### func (CGSConnectionIDᵖ) ElementSize

```go
func (p CGSConnectionIDᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGSConnectionIDᵖ points
to.

#### func (CGSConnectionIDᵖ) OnRead

```go
func (p CGSConnectionIDᵖ) OnRead(ϟs *gfxapi.State) CGSConnectionIDᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (CGSConnectionIDᵖ) OnWrite

```go
func (p CGSConnectionIDᵖ) OnWrite(ϟs *gfxapi.State) CGSConnectionIDᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (CGSConnectionIDᵖ) Read

```go
func (p CGSConnectionIDᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGSConnectionID
```
Read reads and returns the CGSConnectionID element at the pointer.

#### func (CGSConnectionIDᵖ) Slice

```go
func (p CGSConnectionIDᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGSConnectionIDˢ
```
Slice returns a new CGSConnectionIDˢ from the pointer using start and end
indices.

#### func (CGSConnectionIDᵖ) Write

```go
func (p CGSConnectionIDᵖ) Write(value CGSConnectionID, ϟs *gfxapi.State)
```
Write writes value to the CGSConnectionID element at the pointer.

#### type CGSGetSurfaceBounds

```go
type CGSGetSurfaceBounds struct {
	binary.Generate

	Cid    CGSConnectionID
	Wid    CGSWindowID
	Sid    CGSSurfaceID
	Bounds F64ᵖ
	Result int64
}
```

//////////////////////////////////////////////////////////////////////////////
CGSGetSurfaceBounds
//////////////////////////////////////////////////////////////////////////////

#### func  NewCGSGetSurfaceBounds

```go
func NewCGSGetSurfaceBounds(Cid memory.Pointer, Wid CGSWindowID, Sid CGSSurfaceID, Bounds memory.Pointer, Result int64) *CGSGetSurfaceBounds
```

#### func (*CGSGetSurfaceBounds) API

```go
func (c *CGSGetSurfaceBounds) API() gfxapi.ID
```

#### func (*CGSGetSurfaceBounds) AddRead

```go
func (a *CGSGetSurfaceBounds) AddRead(rng memory.Range, id binary.ID) *CGSGetSurfaceBounds
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The CGSGetSurfaceBounds pointer is returned so that calls can be
chained.

#### func (*CGSGetSurfaceBounds) AddWrite

```go
func (a *CGSGetSurfaceBounds) AddWrite(rng memory.Range, id binary.ID) *CGSGetSurfaceBounds
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The CGSGetSurfaceBounds pointer is returned so that calls can be
chained.

#### func (*CGSGetSurfaceBounds) Class

```go
func (*CGSGetSurfaceBounds) Class() binary.Class
```

#### func (*CGSGetSurfaceBounds) Flags

```go
func (c *CGSGetSurfaceBounds) Flags() atom.Flags
```

#### func (*CGSGetSurfaceBounds) Mutate

```go
func (ϟa *CGSGetSurfaceBounds) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*CGSGetSurfaceBounds) Observations

```go
func (a *CGSGetSurfaceBounds) Observations() *atom.Observations
```

#### func (*CGSGetSurfaceBounds) String

```go
func (a *CGSGetSurfaceBounds) String() string
```

#### type CGSSurfaceID

```go
type CGSSurfaceID int32
```


#### type CGSSurfaceIDˢ

```go
type CGSSurfaceIDˢ struct {
	binary.Generate
	SliceInfo
}
```

CGSSurfaceIDˢ is a slice of CGSSurfaceID.

#### func  AsCGSSurfaceIDˢ

```go
func AsCGSSurfaceIDˢ(s Slice, ϟs *gfxapi.State) CGSSurfaceIDˢ
```
AsCGSSurfaceIDˢ returns s cast to a CGSSurfaceIDˢ. The returned slice length
will be calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeCGSSurfaceIDˢ

```go
func MakeCGSSurfaceIDˢ(count uint64, ϟs *gfxapi.State) CGSSurfaceIDˢ
```
MakeCGSSurfaceIDˢ returns a CGSSurfaceIDˢ backed by a new memory pool.

#### func (*CGSSurfaceIDˢ) Class

```go
func (*CGSSurfaceIDˢ) Class() binary.Class
```

#### func (CGSSurfaceIDˢ) Clone

```go
func (s CGSSurfaceIDˢ) Clone(ϟs *gfxapi.State) CGSSurfaceIDˢ
```
Clone returns a copy of the CGSSurfaceIDˢ in a new memory pool.

#### func (CGSSurfaceIDˢ) Copy

```go
func (dst CGSSurfaceIDˢ) Copy(src CGSSurfaceIDˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGSSurfaceIDˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (CGSSurfaceIDˢ) Decoder

```go
func (s CGSSurfaceIDˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (CGSSurfaceIDˢ) ElementSize

```go
func (s CGSSurfaceIDˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGSSurfaceIDˢ points
to.

#### func (CGSSurfaceIDˢ) Encoder

```go
func (s CGSSurfaceIDˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (CGSSurfaceIDˢ) Index

```go
func (s CGSSurfaceIDˢ) Index(i uint64, ϟs *gfxapi.State) CGSSurfaceIDᵖ
```
Index returns a CGSSurfaceIDᵖ to the i'th element in this CGSSurfaceIDˢ.

#### func (CGSSurfaceIDˢ) OnRead

```go
func (s CGSSurfaceIDˢ) OnRead(ϟs *gfxapi.State) CGSSurfaceIDˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (CGSSurfaceIDˢ) OnWrite

```go
func (s CGSSurfaceIDˢ) OnWrite(ϟs *gfxapi.State) CGSSurfaceIDˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (CGSSurfaceIDˢ) Range

```go
func (s CGSSurfaceIDˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (CGSSurfaceIDˢ) Read

```go
func (s CGSSurfaceIDˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGSSurfaceID
```
Read reads and returns all the CGSSurfaceID elements in this CGSSurfaceIDˢ.

#### func (CGSSurfaceIDˢ) ResourceID

```go
func (s CGSSurfaceIDˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (CGSSurfaceIDˢ) Slice

```go
func (s CGSSurfaceIDˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGSSurfaceIDˢ
```
Slice returns a sub-slice from the CGSSurfaceIDˢ using start and end indices.

#### func (CGSSurfaceIDˢ) String

```go
func (s CGSSurfaceIDˢ) String() string
```
String returns a string description of the CGSSurfaceIDˢ slice.

#### func (CGSSurfaceIDˢ) Write

```go
func (s CGSSurfaceIDˢ) Write(src []CGSSurfaceID, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type CGSSurfaceIDᵖ

```go
type CGSSurfaceIDᵖ struct {
	binary.Generate
	memory.Pointer
}
```

CGSSurfaceIDᵖ is a pointer to a CGSSurfaceID element.

#### func  NewCGSSurfaceIDᵖ

```go
func NewCGSSurfaceIDᵖ(addr uint64) CGSSurfaceIDᵖ
```
NewCGSSurfaceIDᵖ returns a CGSSurfaceIDᵖ that points to addr in the application
pool.

#### func (*CGSSurfaceIDᵖ) Class

```go
func (*CGSSurfaceIDᵖ) Class() binary.Class
```

#### func (CGSSurfaceIDᵖ) ElementSize

```go
func (p CGSSurfaceIDᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGSSurfaceIDᵖ points
to.

#### func (CGSSurfaceIDᵖ) OnRead

```go
func (p CGSSurfaceIDᵖ) OnRead(ϟs *gfxapi.State) CGSSurfaceIDᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (CGSSurfaceIDᵖ) OnWrite

```go
func (p CGSSurfaceIDᵖ) OnWrite(ϟs *gfxapi.State) CGSSurfaceIDᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (CGSSurfaceIDᵖ) Read

```go
func (p CGSSurfaceIDᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGSSurfaceID
```
Read reads and returns the CGSSurfaceID element at the pointer.

#### func (CGSSurfaceIDᵖ) Slice

```go
func (p CGSSurfaceIDᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGSSurfaceIDˢ
```
Slice returns a new CGSSurfaceIDˢ from the pointer using start and end indices.

#### func (CGSSurfaceIDᵖ) Write

```go
func (p CGSSurfaceIDᵖ) Write(value CGSSurfaceID, ϟs *gfxapi.State)
```
Write writes value to the CGSSurfaceID element at the pointer.

#### type CGSWindowID

```go
type CGSWindowID int32
```


#### type CGSWindowIDˢ

```go
type CGSWindowIDˢ struct {
	binary.Generate
	SliceInfo
}
```

CGSWindowIDˢ is a slice of CGSWindowID.

#### func  AsCGSWindowIDˢ

```go
func AsCGSWindowIDˢ(s Slice, ϟs *gfxapi.State) CGSWindowIDˢ
```
AsCGSWindowIDˢ returns s cast to a CGSWindowIDˢ. The returned slice length will
be calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeCGSWindowIDˢ

```go
func MakeCGSWindowIDˢ(count uint64, ϟs *gfxapi.State) CGSWindowIDˢ
```
MakeCGSWindowIDˢ returns a CGSWindowIDˢ backed by a new memory pool.

#### func (*CGSWindowIDˢ) Class

```go
func (*CGSWindowIDˢ) Class() binary.Class
```

#### func (CGSWindowIDˢ) Clone

```go
func (s CGSWindowIDˢ) Clone(ϟs *gfxapi.State) CGSWindowIDˢ
```
Clone returns a copy of the CGSWindowIDˢ in a new memory pool.

#### func (CGSWindowIDˢ) Copy

```go
func (dst CGSWindowIDˢ) Copy(src CGSWindowIDˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s CGSWindowIDˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (CGSWindowIDˢ) Decoder

```go
func (s CGSWindowIDˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (CGSWindowIDˢ) ElementSize

```go
func (s CGSWindowIDˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGSWindowIDˢ points to.

#### func (CGSWindowIDˢ) Encoder

```go
func (s CGSWindowIDˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (CGSWindowIDˢ) Index

```go
func (s CGSWindowIDˢ) Index(i uint64, ϟs *gfxapi.State) CGSWindowIDᵖ
```
Index returns a CGSWindowIDᵖ to the i'th element in this CGSWindowIDˢ.

#### func (CGSWindowIDˢ) OnRead

```go
func (s CGSWindowIDˢ) OnRead(ϟs *gfxapi.State) CGSWindowIDˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (CGSWindowIDˢ) OnWrite

```go
func (s CGSWindowIDˢ) OnWrite(ϟs *gfxapi.State) CGSWindowIDˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (CGSWindowIDˢ) Range

```go
func (s CGSWindowIDˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (CGSWindowIDˢ) Read

```go
func (s CGSWindowIDˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []CGSWindowID
```
Read reads and returns all the CGSWindowID elements in this CGSWindowIDˢ.

#### func (CGSWindowIDˢ) ResourceID

```go
func (s CGSWindowIDˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (CGSWindowIDˢ) Slice

```go
func (s CGSWindowIDˢ) Slice(start, end uint64, ϟs *gfxapi.State) CGSWindowIDˢ
```
Slice returns a sub-slice from the CGSWindowIDˢ using start and end indices.

#### func (CGSWindowIDˢ) String

```go
func (s CGSWindowIDˢ) String() string
```
String returns a string description of the CGSWindowIDˢ slice.

#### func (CGSWindowIDˢ) Write

```go
func (s CGSWindowIDˢ) Write(src []CGSWindowID, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type CGSWindowIDᵖ

```go
type CGSWindowIDᵖ struct {
	binary.Generate
	memory.Pointer
}
```

CGSWindowIDᵖ is a pointer to a CGSWindowID element.

#### func  NewCGSWindowIDᵖ

```go
func NewCGSWindowIDᵖ(addr uint64) CGSWindowIDᵖ
```
NewCGSWindowIDᵖ returns a CGSWindowIDᵖ that points to addr in the application
pool.

#### func (*CGSWindowIDᵖ) Class

```go
func (*CGSWindowIDᵖ) Class() binary.Class
```

#### func (CGSWindowIDᵖ) ElementSize

```go
func (p CGSWindowIDᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that CGSWindowIDᵖ points to.

#### func (CGSWindowIDᵖ) OnRead

```go
func (p CGSWindowIDᵖ) OnRead(ϟs *gfxapi.State) CGSWindowIDᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (CGSWindowIDᵖ) OnWrite

```go
func (p CGSWindowIDᵖ) OnWrite(ϟs *gfxapi.State) CGSWindowIDᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (CGSWindowIDᵖ) Read

```go
func (p CGSWindowIDᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) CGSWindowID
```
Read reads and returns the CGSWindowID element at the pointer.

#### func (CGSWindowIDᵖ) Slice

```go
func (p CGSWindowIDᵖ) Slice(start, end uint64, ϟs *gfxapi.State) CGSWindowIDˢ
```
Slice returns a new CGSWindowIDˢ from the pointer using start and end indices.

#### func (CGSWindowIDᵖ) Write

```go
func (p CGSWindowIDᵖ) Write(value CGSWindowID, ϟs *gfxapi.State)
```
Write writes value to the CGSWindowID element at the pointer.

#### type Capability

```go
type Capability uint32
```

//////////////////////////////////////////////////////////////////////////////
enum Capability
//////////////////////////////////////////////////////////////////////////////

#### func (*Capability) Parse

```go
func (v *Capability) Parse(s string) error
```

#### func (Capability) String

```go
func (v Capability) String() string
```

#### type Capabilityːboolᵐ

```go
type Capabilityːboolᵐ map[Capability]bool
```


#### func (Capabilityːboolᵐ) Contains

```go
func (m Capabilityːboolᵐ) Contains(key Capability) bool
```

#### func (Capabilityːboolᵐ) Delete

```go
func (m Capabilityːboolᵐ) Delete(key Capability)
```

#### func (Capabilityːboolᵐ) Get

```go
func (m Capabilityːboolᵐ) Get(key Capability) bool
```

#### func (Capabilityːboolᵐ) Range

```go
func (m Capabilityːboolᵐ) Range() []bool
```

#### type Charˢ

```go
type Charˢ struct {
	binary.Generate
	SliceInfo
}
```

Charˢ is a slice of byte.

#### func  AsCharˢ

```go
func AsCharˢ(s Slice, ϟs *gfxapi.State) Charˢ
```
AsCharˢ returns s cast to a Charˢ. The returned slice length will be calculated
so that the returned slice is no longer (in bytes) than s.

#### func  MakeCharˢ

```go
func MakeCharˢ(count uint64, ϟs *gfxapi.State) Charˢ
```
MakeCharˢ returns a Charˢ backed by a new memory pool.

#### func  MakeCharˢFromString

```go
func MakeCharˢFromString(str string, ϟs *gfxapi.State) Charˢ
```
MakeCharˢFromString returns a Charˢ backed by a new memory pool containing a
copy of str.

#### func (*Charˢ) Class

```go
func (*Charˢ) Class() binary.Class
```

#### func (Charˢ) Clone

```go
func (s Charˢ) Clone(ϟs *gfxapi.State) Charˢ
```
Clone returns a copy of the Charˢ in a new memory pool.

#### func (Charˢ) Copy

```go
func (dst Charˢ) Copy(src Charˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Charˢ) Decoder

```go
func (s Charˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Charˢ) ElementSize

```go
func (s Charˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charˢ points to.

#### func (Charˢ) Encoder

```go
func (s Charˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Charˢ) Index

```go
func (s Charˢ) Index(i uint64, ϟs *gfxapi.State) Charᵖ
```
Index returns a Charᵖ to the i'th element in this Charˢ.

#### func (Charˢ) OnRead

```go
func (s Charˢ) OnRead(ϟs *gfxapi.State) Charˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Charˢ) OnWrite

```go
func (s Charˢ) OnWrite(ϟs *gfxapi.State) Charˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Charˢ) Range

```go
func (s Charˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Charˢ) Read

```go
func (s Charˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []byte
```
Read reads and returns all the byte elements in this Charˢ.

#### func (Charˢ) ResourceID

```go
func (s Charˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Charˢ) Slice

```go
func (s Charˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ
```
Slice returns a sub-slice from the Charˢ using start and end indices.

#### func (Charˢ) String

```go
func (s Charˢ) String() string
```
String returns a string description of the Charˢ slice.

#### func (Charˢ) Write

```go
func (s Charˢ) Write(src []byte, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Charᵖ

```go
type Charᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Charᵖ is a pointer to a byte element.

#### func  NewCharᵖ

```go
func NewCharᵖ(addr uint64) Charᵖ
```
NewCharᵖ returns a Charᵖ that points to addr in the application pool.

#### func (*Charᵖ) Class

```go
func (*Charᵖ) Class() binary.Class
```

#### func (Charᵖ) ElementSize

```go
func (p Charᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᵖ points to.

#### func (Charᵖ) OnRead

```go
func (p Charᵖ) OnRead(ϟs *gfxapi.State) Charᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Charᵖ) OnWrite

```go
func (p Charᵖ) OnWrite(ϟs *gfxapi.State) Charᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Charᵖ) Read

```go
func (p Charᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) byte
```
Read reads and returns the byte element at the pointer.

#### func (Charᵖ) Slice

```go
func (p Charᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ
```
Slice returns a new Charˢ from the pointer using start and end indices.

#### func (Charᵖ) StringSlice

```go
func (p Charᵖ) StringSlice(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, incNullTerm bool) Charˢ
```
StringSlice returns a slice starting at p and ending at the first 0 byte
null-terminator. If incNullTerm is true then the null-terminator is included in
the slice.

#### func (Charᵖ) Write

```go
func (p Charᵖ) Write(value byte, ϟs *gfxapi.State)
```
Write writes value to the byte element at the pointer.

#### type Charᶜᵖ

```go
type Charᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Charᶜᵖ is a pointer to a byte element.

#### func  NewCharᶜᵖ

```go
func NewCharᶜᵖ(addr uint64) Charᶜᵖ
```
NewCharᶜᵖ returns a Charᶜᵖ that points to addr in the application pool.

#### func (*Charᶜᵖ) Class

```go
func (*Charᶜᵖ) Class() binary.Class
```

#### func (Charᶜᵖ) ElementSize

```go
func (p Charᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᶜᵖ points to.

#### func (Charᶜᵖ) OnRead

```go
func (p Charᶜᵖ) OnRead(ϟs *gfxapi.State) Charᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Charᶜᵖ) OnWrite

```go
func (p Charᶜᵖ) OnWrite(ϟs *gfxapi.State) Charᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Charᶜᵖ) Read

```go
func (p Charᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) byte
```
Read reads and returns the byte element at the pointer.

#### func (Charᶜᵖ) Slice

```go
func (p Charᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charˢ
```
Slice returns a new Charˢ from the pointer using start and end indices.

#### func (Charᶜᵖ) StringSlice

```go
func (p Charᶜᵖ) StringSlice(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, incNullTerm bool) Charˢ
```
StringSlice returns a slice starting at p and ending at the first 0 byte
null-terminator. If incNullTerm is true then the null-terminator is included in
the slice.

#### func (Charᶜᵖ) Write

```go
func (p Charᶜᵖ) Write(value byte, ϟs *gfxapi.State)
```
Write writes value to the byte element at the pointer.

#### type Charᶜᵖˢ

```go
type Charᶜᵖˢ struct {
	binary.Generate
	SliceInfo
}
```

Charᶜᵖˢ is a slice of Charᶜᵖ.

#### func  AsCharᶜᵖˢ

```go
func AsCharᶜᵖˢ(s Slice, ϟs *gfxapi.State) Charᶜᵖˢ
```
AsCharᶜᵖˢ returns s cast to a Charᶜᵖˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeCharᶜᵖˢ

```go
func MakeCharᶜᵖˢ(count uint64, ϟs *gfxapi.State) Charᶜᵖˢ
```
MakeCharᶜᵖˢ returns a Charᶜᵖˢ backed by a new memory pool.

#### func (*Charᶜᵖˢ) Class

```go
func (*Charᶜᵖˢ) Class() binary.Class
```

#### func (Charᶜᵖˢ) Clone

```go
func (s Charᶜᵖˢ) Clone(ϟs *gfxapi.State) Charᶜᵖˢ
```
Clone returns a copy of the Charᶜᵖˢ in a new memory pool.

#### func (Charᶜᵖˢ) Copy

```go
func (dst Charᶜᵖˢ) Copy(src Charᶜᵖˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Charᶜᵖˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Charᶜᵖˢ) Decoder

```go
func (s Charᶜᵖˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Charᶜᵖˢ) ElementSize

```go
func (s Charᶜᵖˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᶜᵖˢ points to.

#### func (Charᶜᵖˢ) Encoder

```go
func (s Charᶜᵖˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Charᶜᵖˢ) Index

```go
func (s Charᶜᵖˢ) Index(i uint64, ϟs *gfxapi.State) Charᶜᵖᵖ
```
Index returns a Charᶜᵖᵖ to the i'th element in this Charᶜᵖˢ.

#### func (Charᶜᵖˢ) OnRead

```go
func (s Charᶜᵖˢ) OnRead(ϟs *gfxapi.State) Charᶜᵖˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Charᶜᵖˢ) OnWrite

```go
func (s Charᶜᵖˢ) OnWrite(ϟs *gfxapi.State) Charᶜᵖˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Charᶜᵖˢ) Range

```go
func (s Charᶜᵖˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Charᶜᵖˢ) Read

```go
func (s Charᶜᵖˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Charᶜᵖ
```
Read reads and returns all the Charᶜᵖ elements in this Charᶜᵖˢ.

#### func (Charᶜᵖˢ) ResourceID

```go
func (s Charᶜᵖˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Charᶜᵖˢ) Slice

```go
func (s Charᶜᵖˢ) Slice(start, end uint64, ϟs *gfxapi.State) Charᶜᵖˢ
```
Slice returns a sub-slice from the Charᶜᵖˢ using start and end indices.

#### func (Charᶜᵖˢ) String

```go
func (s Charᶜᵖˢ) String() string
```
String returns a string description of the Charᶜᵖˢ slice.

#### func (Charᶜᵖˢ) Write

```go
func (s Charᶜᵖˢ) Write(src []Charᶜᵖ, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Charᶜᵖᵖ

```go
type Charᶜᵖᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Charᶜᵖᵖ is a pointer to a Charᶜᵖ element. Note: Pointers are stored differently
between the application pool and internal pools.

    * The application pool stores pointers as an address of an architecture-dependant size.
    * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
      pool identifier.

#### func  NewCharᶜᵖᵖ

```go
func NewCharᶜᵖᵖ(addr uint64) Charᶜᵖᵖ
```
NewCharᶜᵖᵖ returns a Charᶜᵖᵖ that points to addr in the application pool.

#### func (*Charᶜᵖᵖ) Class

```go
func (*Charᶜᵖᵖ) Class() binary.Class
```

#### func (Charᶜᵖᵖ) ElementSize

```go
func (p Charᶜᵖᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᶜᵖᵖ points to.

#### func (Charᶜᵖᵖ) OnRead

```go
func (p Charᶜᵖᵖ) OnRead(ϟs *gfxapi.State) Charᶜᵖᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Charᶜᵖᵖ) OnWrite

```go
func (p Charᶜᵖᵖ) OnWrite(ϟs *gfxapi.State) Charᶜᵖᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Charᶜᵖᵖ) Read

```go
func (p Charᶜᵖᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Charᶜᵖ
```
Read reads and returns the Charᶜᵖ element at the pointer.

#### func (Charᶜᵖᵖ) Slice

```go
func (p Charᶜᵖᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charᶜᵖˢ
```
Slice returns a new Charᶜᵖˢ from the pointer using start and end indices.

#### func (Charᶜᵖᵖ) Write

```go
func (p Charᶜᵖᵖ) Write(value Charᶜᵖ, ϟs *gfxapi.State)
```
Write writes value to the Charᶜᵖ element at the pointer.

#### type Charᶜᵖᶜᵖ

```go
type Charᶜᵖᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Charᶜᵖᶜᵖ is a pointer to a Charᶜᵖ element. Note: Pointers are stored differently
between the application pool and internal pools.

    * The application pool stores pointers as an address of an architecture-dependant size.
    * Internal pools store pointers as an 64-bit unsigned address and a 32-bit unsigned
      pool identifier.

#### func  NewCharᶜᵖᶜᵖ

```go
func NewCharᶜᵖᶜᵖ(addr uint64) Charᶜᵖᶜᵖ
```
NewCharᶜᵖᶜᵖ returns a Charᶜᵖᶜᵖ that points to addr in the application pool.

#### func (*Charᶜᵖᶜᵖ) Class

```go
func (*Charᶜᵖᶜᵖ) Class() binary.Class
```

#### func (Charᶜᵖᶜᵖ) ElementSize

```go
func (p Charᶜᵖᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Charᶜᵖᶜᵖ points to.

#### func (Charᶜᵖᶜᵖ) OnRead

```go
func (p Charᶜᵖᶜᵖ) OnRead(ϟs *gfxapi.State) Charᶜᵖᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Charᶜᵖᶜᵖ) OnWrite

```go
func (p Charᶜᵖᶜᵖ) OnWrite(ϟs *gfxapi.State) Charᶜᵖᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Charᶜᵖᶜᵖ) Read

```go
func (p Charᶜᵖᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Charᶜᵖ
```
Read reads and returns the Charᶜᵖ element at the pointer.

#### func (Charᶜᵖᶜᵖ) Slice

```go
func (p Charᶜᵖᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Charᶜᵖˢ
```
Slice returns a new Charᶜᵖˢ from the pointer using start and end indices.

#### func (Charᶜᵖᶜᵖ) Write

```go
func (p Charᶜᵖᶜᵖ) Write(value Charᶜᵖ, ϟs *gfxapi.State)
```
Write writes value to the Charᶜᵖ element at the pointer.

#### type ClearMask

```go
type ClearMask uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ClearMask
//////////////////////////////////////////////////////////////////////////////

#### func (*ClearMask) Parse

```go
func (v *ClearMask) Parse(s string) error
```

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

#### type ClientWaitSyncSignal

```go
type ClientWaitSyncSignal uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ClientWaitSyncSignal
//////////////////////////////////////////////////////////////////////////////

#### func (*ClientWaitSyncSignal) Parse

```go
func (v *ClientWaitSyncSignal) Parse(s string) error
```

#### func (ClientWaitSyncSignal) String

```go
func (v ClientWaitSyncSignal) String() string
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

#### func (*CompressedTexelFormat) Parse

```go
func (v *CompressedTexelFormat) Parse(s string) error
```

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

#### func (*CompressedTexelFormat_AMD_compressed_ATC_texture) Parse

```go
func (v *CompressedTexelFormat_AMD_compressed_ATC_texture) Parse(s string) error
```

#### func (CompressedTexelFormat_AMD_compressed_ATC_texture) String

```go
func (v CompressedTexelFormat_AMD_compressed_ATC_texture) String() string
```

#### type CompressedTexelFormat_EXT_texture_compression_dxt1

```go
type CompressedTexelFormat_EXT_texture_compression_dxt1 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat_EXT_texture_compression_dxt1
//////////////////////////////////////////////////////////////////////////////

#### func (*CompressedTexelFormat_EXT_texture_compression_dxt1) Parse

```go
func (v *CompressedTexelFormat_EXT_texture_compression_dxt1) Parse(s string) error
```

#### func (CompressedTexelFormat_EXT_texture_compression_dxt1) String

```go
func (v CompressedTexelFormat_EXT_texture_compression_dxt1) String() string
```

#### type CompressedTexelFormat_EXT_texture_compression_s3tc

```go
type CompressedTexelFormat_EXT_texture_compression_s3tc uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat_EXT_texture_compression_s3tc
//////////////////////////////////////////////////////////////////////////////

#### func (*CompressedTexelFormat_EXT_texture_compression_s3tc) Parse

```go
func (v *CompressedTexelFormat_EXT_texture_compression_s3tc) Parse(s string) error
```

#### func (CompressedTexelFormat_EXT_texture_compression_s3tc) String

```go
func (v CompressedTexelFormat_EXT_texture_compression_s3tc) String() string
```

#### type CompressedTexelFormat_KHR_texture_compression_astc_ldr

```go
type CompressedTexelFormat_KHR_texture_compression_astc_ldr uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat_KHR_texture_compression_astc_ldr
//////////////////////////////////////////////////////////////////////////////

#### func (*CompressedTexelFormat_KHR_texture_compression_astc_ldr) Parse

```go
func (v *CompressedTexelFormat_KHR_texture_compression_astc_ldr) Parse(s string) error
```

#### func (CompressedTexelFormat_KHR_texture_compression_astc_ldr) String

```go
func (v CompressedTexelFormat_KHR_texture_compression_astc_ldr) String() string
```

#### type CompressedTexelFormat_NV_texture_compression_latc

```go
type CompressedTexelFormat_NV_texture_compression_latc uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat_NV_texture_compression_latc
//////////////////////////////////////////////////////////////////////////////

#### func (*CompressedTexelFormat_NV_texture_compression_latc) Parse

```go
func (v *CompressedTexelFormat_NV_texture_compression_latc) Parse(s string) error
```

#### func (CompressedTexelFormat_NV_texture_compression_latc) String

```go
func (v CompressedTexelFormat_NV_texture_compression_latc) String() string
```

#### type CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture

```go
type CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture
//////////////////////////////////////////////////////////////////////////////

#### func (*CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture) Parse

```go
func (v *CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture) Parse(s string) error
```

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

#### type ContextID

```go
type ContextID uint32
```


#### type CubeMapImageTarget

```go
type CubeMapImageTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum CubeMapImageTarget
//////////////////////////////////////////////////////////////////////////////

#### func (*CubeMapImageTarget) Parse

```go
func (v *CubeMapImageTarget) Parse(s string) error
```

#### func (CubeMapImageTarget) String

```go
func (v CubeMapImageTarget) String() string
```

#### type CubeMapImageTargetːImageᵐ

```go
type CubeMapImageTargetːImageᵐ map[CubeMapImageTarget]Image
```


#### func (CubeMapImageTargetːImageᵐ) Contains

```go
func (m CubeMapImageTargetːImageᵐ) Contains(key CubeMapImageTarget) bool
```

#### func (CubeMapImageTargetːImageᵐ) Delete

```go
func (m CubeMapImageTargetːImageᵐ) Delete(key CubeMapImageTarget)
```

#### func (CubeMapImageTargetːImageᵐ) Get

```go
func (m CubeMapImageTargetːImageᵐ) Get(key CubeMapImageTarget) Image
```

#### func (CubeMapImageTargetːImageᵐ) Range

```go
func (m CubeMapImageTargetːImageᵐ) Range() []Image
```

#### type CubemapLevel

```go
type CubemapLevel struct {
	binary.Generate
	CreatedAt atom.ID
	Faces     CubeMapImageTargetːImageᵐ
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

#### type DiscardFramebufferAttachment

```go
type DiscardFramebufferAttachment uint32
```

//////////////////////////////////////////////////////////////////////////////
enum DiscardFramebufferAttachment
//////////////////////////////////////////////////////////////////////////////

#### func (*DiscardFramebufferAttachment) Parse

```go
func (v *DiscardFramebufferAttachment) Parse(s string) error
```

#### func (DiscardFramebufferAttachment) String

```go
func (v DiscardFramebufferAttachment) String() string
```

#### type DiscardFramebufferAttachmentˢ

```go
type DiscardFramebufferAttachmentˢ struct {
	binary.Generate
	SliceInfo
}
```

DiscardFramebufferAttachmentˢ is a slice of DiscardFramebufferAttachment.

#### func  AsDiscardFramebufferAttachmentˢ

```go
func AsDiscardFramebufferAttachmentˢ(s Slice, ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ
```
AsDiscardFramebufferAttachmentˢ returns s cast to a
DiscardFramebufferAttachmentˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeDiscardFramebufferAttachmentˢ

```go
func MakeDiscardFramebufferAttachmentˢ(count uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ
```
MakeDiscardFramebufferAttachmentˢ returns a DiscardFramebufferAttachmentˢ backed
by a new memory pool.

#### func (*DiscardFramebufferAttachmentˢ) Class

```go
func (*DiscardFramebufferAttachmentˢ) Class() binary.Class
```

#### func (DiscardFramebufferAttachmentˢ) Clone

```go
func (s DiscardFramebufferAttachmentˢ) Clone(ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ
```
Clone returns a copy of the DiscardFramebufferAttachmentˢ in a new memory pool.

#### func (DiscardFramebufferAttachmentˢ) Copy

```go
func (dst DiscardFramebufferAttachmentˢ) Copy(src DiscardFramebufferAttachmentˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s DiscardFramebufferAttachmentˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (DiscardFramebufferAttachmentˢ) Decoder

```go
func (s DiscardFramebufferAttachmentˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (DiscardFramebufferAttachmentˢ) ElementSize

```go
func (s DiscardFramebufferAttachmentˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that
DiscardFramebufferAttachmentˢ points to.

#### func (DiscardFramebufferAttachmentˢ) Encoder

```go
func (s DiscardFramebufferAttachmentˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (DiscardFramebufferAttachmentˢ) Index

```go
func (s DiscardFramebufferAttachmentˢ) Index(i uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentᵖ
```
Index returns a DiscardFramebufferAttachmentᵖ to the i'th element in this
DiscardFramebufferAttachmentˢ.

#### func (DiscardFramebufferAttachmentˢ) OnRead

```go
func (s DiscardFramebufferAttachmentˢ) OnRead(ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (DiscardFramebufferAttachmentˢ) OnWrite

```go
func (s DiscardFramebufferAttachmentˢ) OnWrite(ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (DiscardFramebufferAttachmentˢ) Range

```go
func (s DiscardFramebufferAttachmentˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (DiscardFramebufferAttachmentˢ) Read

```go
func (s DiscardFramebufferAttachmentˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []DiscardFramebufferAttachment
```
Read reads and returns all the DiscardFramebufferAttachment elements in this
DiscardFramebufferAttachmentˢ.

#### func (DiscardFramebufferAttachmentˢ) ResourceID

```go
func (s DiscardFramebufferAttachmentˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (DiscardFramebufferAttachmentˢ) Slice

```go
func (s DiscardFramebufferAttachmentˢ) Slice(start, end uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ
```
Slice returns a sub-slice from the DiscardFramebufferAttachmentˢ using start and
end indices.

#### func (DiscardFramebufferAttachmentˢ) String

```go
func (s DiscardFramebufferAttachmentˢ) String() string
```
String returns a string description of the DiscardFramebufferAttachmentˢ slice.

#### func (DiscardFramebufferAttachmentˢ) Write

```go
func (s DiscardFramebufferAttachmentˢ) Write(src []DiscardFramebufferAttachment, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type DiscardFramebufferAttachmentᵖ

```go
type DiscardFramebufferAttachmentᵖ struct {
	binary.Generate
	memory.Pointer
}
```

DiscardFramebufferAttachmentᵖ is a pointer to a DiscardFramebufferAttachment
element.

#### func  NewDiscardFramebufferAttachmentᵖ

```go
func NewDiscardFramebufferAttachmentᵖ(addr uint64) DiscardFramebufferAttachmentᵖ
```
NewDiscardFramebufferAttachmentᵖ returns a DiscardFramebufferAttachmentᵖ that
points to addr in the application pool.

#### func (*DiscardFramebufferAttachmentᵖ) Class

```go
func (*DiscardFramebufferAttachmentᵖ) Class() binary.Class
```

#### func (DiscardFramebufferAttachmentᵖ) ElementSize

```go
func (p DiscardFramebufferAttachmentᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that
DiscardFramebufferAttachmentᵖ points to.

#### func (DiscardFramebufferAttachmentᵖ) OnRead

```go
func (p DiscardFramebufferAttachmentᵖ) OnRead(ϟs *gfxapi.State) DiscardFramebufferAttachmentᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (DiscardFramebufferAttachmentᵖ) OnWrite

```go
func (p DiscardFramebufferAttachmentᵖ) OnWrite(ϟs *gfxapi.State) DiscardFramebufferAttachmentᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (DiscardFramebufferAttachmentᵖ) Read

```go
func (p DiscardFramebufferAttachmentᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) DiscardFramebufferAttachment
```
Read reads and returns the DiscardFramebufferAttachment element at the pointer.

#### func (DiscardFramebufferAttachmentᵖ) Slice

```go
func (p DiscardFramebufferAttachmentᵖ) Slice(start, end uint64, ϟs *gfxapi.State) DiscardFramebufferAttachmentˢ
```
Slice returns a new DiscardFramebufferAttachmentˢ from the pointer using start
and end indices.

#### func (DiscardFramebufferAttachmentᵖ) Write

```go
func (p DiscardFramebufferAttachmentᵖ) Write(value DiscardFramebufferAttachment, ϟs *gfxapi.State)
```
Write writes value to the DiscardFramebufferAttachment element at the pointer.

#### type DrawMode

```go
type DrawMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum DrawMode
//////////////////////////////////////////////////////////////////////////////

#### func (*DrawMode) Parse

```go
func (v *DrawMode) Parse(s string) error
```

#### func (DrawMode) String

```go
func (v DrawMode) String() string
```

#### type EGLBoolean

```go
type EGLBoolean int64
```


#### type EGLConfig

```go
type EGLConfig struct {
	binary.Generate
	memory.Pointer
}
```

EGLConfig is a pointer to a void element.

#### func  NewEGLConfig

```go
func NewEGLConfig(addr uint64) EGLConfig
```
NewEGLConfig returns a EGLConfig that points to addr in the application pool.

#### func (*EGLConfig) Class

```go
func (*EGLConfig) Class() binary.Class
```

#### func (EGLConfig) ElementSize

```go
func (p EGLConfig) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that EGLConfig points to.

#### func (EGLConfig) OnRead

```go
func (p EGLConfig) OnRead(ϟs *gfxapi.State) EGLConfig
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (EGLConfig) OnWrite

```go
func (p EGLConfig) OnWrite(ϟs *gfxapi.State) EGLConfig
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (EGLConfig) Slice

```go
func (p EGLConfig) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type EGLContext

```go
type EGLContext struct {
	binary.Generate
	memory.Pointer
}
```

EGLContext is a pointer to a void element.

#### func  NewEGLContext

```go
func NewEGLContext(addr uint64) EGLContext
```
NewEGLContext returns a EGLContext that points to addr in the application pool.

#### func (*EGLContext) Class

```go
func (*EGLContext) Class() binary.Class
```

#### func (EGLContext) ElementSize

```go
func (p EGLContext) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that EGLContext points to.

#### func (EGLContext) OnRead

```go
func (p EGLContext) OnRead(ϟs *gfxapi.State) EGLContext
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (EGLContext) OnWrite

```go
func (p EGLContext) OnWrite(ϟs *gfxapi.State) EGLContext
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (EGLContext) Slice

```go
func (p EGLContext) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type EGLContextːContextʳᵐ

```go
type EGLContextːContextʳᵐ map[EGLContext](*Context)
```


#### func (EGLContextːContextʳᵐ) Contains

```go
func (m EGLContextːContextʳᵐ) Contains(key EGLContext) bool
```

#### func (EGLContextːContextʳᵐ) Delete

```go
func (m EGLContextːContextʳᵐ) Delete(key EGLContext)
```

#### func (EGLContextːContextʳᵐ) Get

```go
func (m EGLContextːContextʳᵐ) Get(key EGLContext) *Context
```

#### func (EGLContextːContextʳᵐ) Range

```go
func (m EGLContextːContextʳᵐ) Range() [](*Context)
```

#### type EGLDisplay

```go
type EGLDisplay struct {
	binary.Generate
	memory.Pointer
}
```

EGLDisplay is a pointer to a void element.

#### func  NewEGLDisplay

```go
func NewEGLDisplay(addr uint64) EGLDisplay
```
NewEGLDisplay returns a EGLDisplay that points to addr in the application pool.

#### func (*EGLDisplay) Class

```go
func (*EGLDisplay) Class() binary.Class
```

#### func (EGLDisplay) ElementSize

```go
func (p EGLDisplay) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that EGLDisplay points to.

#### func (EGLDisplay) OnRead

```go
func (p EGLDisplay) OnRead(ϟs *gfxapi.State) EGLDisplay
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (EGLDisplay) OnWrite

```go
func (p EGLDisplay) OnWrite(ϟs *gfxapi.State) EGLDisplay
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (EGLDisplay) Slice

```go
func (p EGLDisplay) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type EGLSurface

```go
type EGLSurface struct {
	binary.Generate
	memory.Pointer
}
```

EGLSurface is a pointer to a void element.

#### func  NewEGLSurface

```go
func NewEGLSurface(addr uint64) EGLSurface
```
NewEGLSurface returns a EGLSurface that points to addr in the application pool.

#### func (*EGLSurface) Class

```go
func (*EGLSurface) Class() binary.Class
```

#### func (EGLSurface) ElementSize

```go
func (p EGLSurface) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that EGLSurface points to.

#### func (EGLSurface) OnRead

```go
func (p EGLSurface) OnRead(ϟs *gfxapi.State) EGLSurface
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (EGLSurface) OnWrite

```go
func (p EGLSurface) OnWrite(ϟs *gfxapi.State) EGLSurface
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (EGLSurface) Slice

```go
func (p EGLSurface) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type EGLint

```go
type EGLint int64
```


#### type EGLintˢ

```go
type EGLintˢ struct {
	binary.Generate
	SliceInfo
}
```

EGLintˢ is a slice of EGLint.

#### func  AsEGLintˢ

```go
func AsEGLintˢ(s Slice, ϟs *gfxapi.State) EGLintˢ
```
AsEGLintˢ returns s cast to a EGLintˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeEGLintˢ

```go
func MakeEGLintˢ(count uint64, ϟs *gfxapi.State) EGLintˢ
```
MakeEGLintˢ returns a EGLintˢ backed by a new memory pool.

#### func (*EGLintˢ) Class

```go
func (*EGLintˢ) Class() binary.Class
```

#### func (EGLintˢ) Clone

```go
func (s EGLintˢ) Clone(ϟs *gfxapi.State) EGLintˢ
```
Clone returns a copy of the EGLintˢ in a new memory pool.

#### func (EGLintˢ) Copy

```go
func (dst EGLintˢ) Copy(src EGLintˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s EGLintˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (EGLintˢ) Decoder

```go
func (s EGLintˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (EGLintˢ) ElementSize

```go
func (s EGLintˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that EGLintˢ points to.

#### func (EGLintˢ) Encoder

```go
func (s EGLintˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (EGLintˢ) Index

```go
func (s EGLintˢ) Index(i uint64, ϟs *gfxapi.State) EGLintᵖ
```
Index returns a EGLintᵖ to the i'th element in this EGLintˢ.

#### func (EGLintˢ) OnRead

```go
func (s EGLintˢ) OnRead(ϟs *gfxapi.State) EGLintˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (EGLintˢ) OnWrite

```go
func (s EGLintˢ) OnWrite(ϟs *gfxapi.State) EGLintˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (EGLintˢ) Range

```go
func (s EGLintˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (EGLintˢ) Read

```go
func (s EGLintˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []EGLint
```
Read reads and returns all the EGLint elements in this EGLintˢ.

#### func (EGLintˢ) ResourceID

```go
func (s EGLintˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (EGLintˢ) Slice

```go
func (s EGLintˢ) Slice(start, end uint64, ϟs *gfxapi.State) EGLintˢ
```
Slice returns a sub-slice from the EGLintˢ using start and end indices.

#### func (EGLintˢ) String

```go
func (s EGLintˢ) String() string
```
String returns a string description of the EGLintˢ slice.

#### func (EGLintˢ) Write

```go
func (s EGLintˢ) Write(src []EGLint, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type EGLintᵖ

```go
type EGLintᵖ struct {
	binary.Generate
	memory.Pointer
}
```

EGLintᵖ is a pointer to a EGLint element.

#### func  NewEGLintᵖ

```go
func NewEGLintᵖ(addr uint64) EGLintᵖ
```
NewEGLintᵖ returns a EGLintᵖ that points to addr in the application pool.

#### func (*EGLintᵖ) Class

```go
func (*EGLintᵖ) Class() binary.Class
```

#### func (EGLintᵖ) ElementSize

```go
func (p EGLintᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that EGLintᵖ points to.

#### func (EGLintᵖ) OnRead

```go
func (p EGLintᵖ) OnRead(ϟs *gfxapi.State) EGLintᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (EGLintᵖ) OnWrite

```go
func (p EGLintᵖ) OnWrite(ϟs *gfxapi.State) EGLintᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (EGLintᵖ) Read

```go
func (p EGLintᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) EGLint
```
Read reads and returns the EGLint element at the pointer.

#### func (EGLintᵖ) Slice

```go
func (p EGLintᵖ) Slice(start, end uint64, ϟs *gfxapi.State) EGLintˢ
```
Slice returns a new EGLintˢ from the pointer using start and end indices.

#### func (EGLintᵖ) Write

```go
func (p EGLintᵖ) Write(value EGLint, ϟs *gfxapi.State)
```
Write writes value to the EGLint element at the pointer.

#### type EglCreateContext

```go
type EglCreateContext struct {
	binary.Generate

	Display      EGLDisplay
	Config       EGLConfig
	ShareContext EGLContext
	AttribList   EGLintᵖ
	Result       EGLContext
}
```

//////////////////////////////////////////////////////////////////////////////
EglCreateContext
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglCreateContext

```go
func NewEglCreateContext(Display memory.Pointer, Config memory.Pointer, Share_context memory.Pointer, Attrib_list memory.Pointer, Result memory.Pointer) *EglCreateContext
```

#### func (*EglCreateContext) API

```go
func (c *EglCreateContext) API() gfxapi.ID
```

#### func (*EglCreateContext) AddRead

```go
func (a *EglCreateContext) AddRead(rng memory.Range, id binary.ID) *EglCreateContext
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The EglCreateContext pointer is returned so that calls can be chained.

#### func (*EglCreateContext) AddWrite

```go
func (a *EglCreateContext) AddWrite(rng memory.Range, id binary.ID) *EglCreateContext
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The EglCreateContext pointer is returned so that calls can be chained.

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
func (ϟa *EglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*EglCreateContext) Observations

```go
func (a *EglCreateContext) Observations() *atom.Observations
```

#### func (*EglCreateContext) Replay

```go
func (ω *EglCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*EglCreateContext) String

```go
func (a *EglCreateContext) String() string
```

#### type EglInitialize

```go
type EglInitialize struct {
	binary.Generate

	Dpy    EGLDisplay
	Major  EGLintᵖ
	Minor  EGLintᵖ
	Result EGLBoolean
}
```

//////////////////////////////////////////////////////////////////////////////
EglInitialize
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglInitialize

```go
func NewEglInitialize(Dpy memory.Pointer, Major memory.Pointer, Minor memory.Pointer, Result EGLBoolean) *EglInitialize
```

#### func (*EglInitialize) API

```go
func (c *EglInitialize) API() gfxapi.ID
```

#### func (*EglInitialize) AddRead

```go
func (a *EglInitialize) AddRead(rng memory.Range, id binary.ID) *EglInitialize
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The EglInitialize pointer is returned so that calls can be chained.

#### func (*EglInitialize) AddWrite

```go
func (a *EglInitialize) AddWrite(rng memory.Range, id binary.ID) *EglInitialize
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The EglInitialize pointer is returned so that calls can be chained.

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
func (ϟa *EglInitialize) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*EglInitialize) Observations

```go
func (a *EglInitialize) Observations() *atom.Observations
```

#### func (*EglInitialize) String

```go
func (a *EglInitialize) String() string
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
func NewEglMakeCurrent(Display memory.Pointer, Draw memory.Pointer, Read memory.Pointer, Context memory.Pointer, Result EGLBoolean) *EglMakeCurrent
```

#### func (*EglMakeCurrent) API

```go
func (c *EglMakeCurrent) API() gfxapi.ID
```

#### func (*EglMakeCurrent) AddRead

```go
func (a *EglMakeCurrent) AddRead(rng memory.Range, id binary.ID) *EglMakeCurrent
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The EglMakeCurrent pointer is returned so that calls can be chained.

#### func (*EglMakeCurrent) AddWrite

```go
func (a *EglMakeCurrent) AddWrite(rng memory.Range, id binary.ID) *EglMakeCurrent
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The EglMakeCurrent pointer is returned so that calls can be chained.

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
func (ϟa *EglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*EglMakeCurrent) Observations

```go
func (a *EglMakeCurrent) Observations() *atom.Observations
```

#### func (*EglMakeCurrent) Replay

```go
func (ω *EglMakeCurrent) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*EglMakeCurrent) String

```go
func (a *EglMakeCurrent) String() string
```

#### type EglQuerySurface

```go
type EglQuerySurface struct {
	binary.Generate

	Display   EGLDisplay
	Surface   EGLSurface
	Attribute EGLint
	Value     EGLintᵖ
	Result    EGLBoolean
}
```

//////////////////////////////////////////////////////////////////////////////
EglQuerySurface
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglQuerySurface

```go
func NewEglQuerySurface(Display memory.Pointer, Surface memory.Pointer, Attribute EGLint, Value memory.Pointer, Result EGLBoolean) *EglQuerySurface
```

#### func (*EglQuerySurface) API

```go
func (c *EglQuerySurface) API() gfxapi.ID
```

#### func (*EglQuerySurface) AddRead

```go
func (a *EglQuerySurface) AddRead(rng memory.Range, id binary.ID) *EglQuerySurface
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The EglQuerySurface pointer is returned so that calls can be chained.

#### func (*EglQuerySurface) AddWrite

```go
func (a *EglQuerySurface) AddWrite(rng memory.Range, id binary.ID) *EglQuerySurface
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The EglQuerySurface pointer is returned so that calls can be chained.

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
func (ϟa *EglQuerySurface) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*EglQuerySurface) Observations

```go
func (a *EglQuerySurface) Observations() *atom.Observations
```

#### func (*EglQuerySurface) String

```go
func (a *EglQuerySurface) String() string
```

#### type EglSwapBuffers

```go
type EglSwapBuffers struct {
	binary.Generate

	Display EGLDisplay
	Surface Voidᵖ
	Result  EGLBoolean
}
```

//////////////////////////////////////////////////////////////////////////////
EglSwapBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewEglSwapBuffers

```go
func NewEglSwapBuffers(Display memory.Pointer, Surface memory.Pointer, Result EGLBoolean) *EglSwapBuffers
```

#### func (*EglSwapBuffers) API

```go
func (c *EglSwapBuffers) API() gfxapi.ID
```

#### func (*EglSwapBuffers) AddRead

```go
func (a *EglSwapBuffers) AddRead(rng memory.Range, id binary.ID) *EglSwapBuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The EglSwapBuffers pointer is returned so that calls can be chained.

#### func (*EglSwapBuffers) AddWrite

```go
func (a *EglSwapBuffers) AddWrite(rng memory.Range, id binary.ID) *EglSwapBuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The EglSwapBuffers pointer is returned so that calls can be chained.

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
func (ϟa *EglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*EglSwapBuffers) Observations

```go
func (a *EglSwapBuffers) Observations() *atom.Observations
```

#### func (*EglSwapBuffers) String

```go
func (a *EglSwapBuffers) String() string
```

#### type Error

```go
type Error uint32
```

//////////////////////////////////////////////////////////////////////////////
enum Error
//////////////////////////////////////////////////////////////////////////////

#### func (*Error) Parse

```go
func (v *Error) Parse(s string) error
```

#### func (Error) String

```go
func (v Error) String() string
```

#### type F32ː2ᵃ

```go
type F32ː2ᵃ struct {
	binary.Generate
	Elements [2]float32
}
```


#### func (*F32ː2ᵃ) Class

```go
func (*F32ː2ᵃ) Class() binary.Class
```

#### type F32ː3ᵃ

```go
type F32ː3ᵃ struct {
	binary.Generate
	Elements [3]float32
}
```


#### func (*F32ː3ᵃ) Class

```go
func (*F32ː3ᵃ) Class() binary.Class
```

#### type F32ː4ᵃ

```go
type F32ː4ᵃ struct {
	binary.Generate
	Elements [4]float32
}
```


#### func (*F32ː4ᵃ) Class

```go
func (*F32ː4ᵃ) Class() binary.Class
```

#### type F32ˢ

```go
type F32ˢ struct {
	binary.Generate
	SliceInfo
}
```

F32ˢ is a slice of float32.

#### func  AsF32ˢ

```go
func AsF32ˢ(s Slice, ϟs *gfxapi.State) F32ˢ
```
AsF32ˢ returns s cast to a F32ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeF32ˢ

```go
func MakeF32ˢ(count uint64, ϟs *gfxapi.State) F32ˢ
```
MakeF32ˢ returns a F32ˢ backed by a new memory pool.

#### func (*F32ˢ) Class

```go
func (*F32ˢ) Class() binary.Class
```

#### func (F32ˢ) Clone

```go
func (s F32ˢ) Clone(ϟs *gfxapi.State) F32ˢ
```
Clone returns a copy of the F32ˢ in a new memory pool.

#### func (F32ˢ) Copy

```go
func (dst F32ˢ) Copy(src F32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F32ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (F32ˢ) Decoder

```go
func (s F32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (F32ˢ) ElementSize

```go
func (s F32ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F32ˢ points to.

#### func (F32ˢ) Encoder

```go
func (s F32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (F32ˢ) Index

```go
func (s F32ˢ) Index(i uint64, ϟs *gfxapi.State) F32ᵖ
```
Index returns a F32ᵖ to the i'th element in this F32ˢ.

#### func (F32ˢ) OnRead

```go
func (s F32ˢ) OnRead(ϟs *gfxapi.State) F32ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (F32ˢ) OnWrite

```go
func (s F32ˢ) OnWrite(ϟs *gfxapi.State) F32ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (F32ˢ) Range

```go
func (s F32ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (F32ˢ) Read

```go
func (s F32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float32
```
Read reads and returns all the float32 elements in this F32ˢ.

#### func (F32ˢ) ResourceID

```go
func (s F32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (F32ˢ) Slice

```go
func (s F32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ
```
Slice returns a sub-slice from the F32ˢ using start and end indices.

#### func (F32ˢ) String

```go
func (s F32ˢ) String() string
```
String returns a string description of the F32ˢ slice.

#### func (F32ˢ) Write

```go
func (s F32ˢ) Write(src []float32, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type F32ᵖ

```go
type F32ᵖ struct {
	binary.Generate
	memory.Pointer
}
```

F32ᵖ is a pointer to a float32 element.

#### func  NewF32ᵖ

```go
func NewF32ᵖ(addr uint64) F32ᵖ
```
NewF32ᵖ returns a F32ᵖ that points to addr in the application pool.

#### func (*F32ᵖ) Class

```go
func (*F32ᵖ) Class() binary.Class
```

#### func (F32ᵖ) ElementSize

```go
func (p F32ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F32ᵖ points to.

#### func (F32ᵖ) OnRead

```go
func (p F32ᵖ) OnRead(ϟs *gfxapi.State) F32ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (F32ᵖ) OnWrite

```go
func (p F32ᵖ) OnWrite(ϟs *gfxapi.State) F32ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (F32ᵖ) Read

```go
func (p F32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float32
```
Read reads and returns the float32 element at the pointer.

#### func (F32ᵖ) Slice

```go
func (p F32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ
```
Slice returns a new F32ˢ from the pointer using start and end indices.

#### func (F32ᵖ) Write

```go
func (p F32ᵖ) Write(value float32, ϟs *gfxapi.State)
```
Write writes value to the float32 element at the pointer.

#### type F32ᶜᵖ

```go
type F32ᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

F32ᶜᵖ is a pointer to a float32 element.

#### func  NewF32ᶜᵖ

```go
func NewF32ᶜᵖ(addr uint64) F32ᶜᵖ
```
NewF32ᶜᵖ returns a F32ᶜᵖ that points to addr in the application pool.

#### func (*F32ᶜᵖ) Class

```go
func (*F32ᶜᵖ) Class() binary.Class
```

#### func (F32ᶜᵖ) ElementSize

```go
func (p F32ᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F32ᶜᵖ points to.

#### func (F32ᶜᵖ) OnRead

```go
func (p F32ᶜᵖ) OnRead(ϟs *gfxapi.State) F32ᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (F32ᶜᵖ) OnWrite

```go
func (p F32ᶜᵖ) OnWrite(ϟs *gfxapi.State) F32ᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (F32ᶜᵖ) Read

```go
func (p F32ᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float32
```
Read reads and returns the float32 element at the pointer.

#### func (F32ᶜᵖ) Slice

```go
func (p F32ᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F32ˢ
```
Slice returns a new F32ˢ from the pointer using start and end indices.

#### func (F32ᶜᵖ) Write

```go
func (p F32ᶜᵖ) Write(value float32, ϟs *gfxapi.State)
```
Write writes value to the float32 element at the pointer.

#### type F64ˢ

```go
type F64ˢ struct {
	binary.Generate
	SliceInfo
}
```

F64ˢ is a slice of float64.

#### func  AsF64ˢ

```go
func AsF64ˢ(s Slice, ϟs *gfxapi.State) F64ˢ
```
AsF64ˢ returns s cast to a F64ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeF64ˢ

```go
func MakeF64ˢ(count uint64, ϟs *gfxapi.State) F64ˢ
```
MakeF64ˢ returns a F64ˢ backed by a new memory pool.

#### func (*F64ˢ) Class

```go
func (*F64ˢ) Class() binary.Class
```

#### func (F64ˢ) Clone

```go
func (s F64ˢ) Clone(ϟs *gfxapi.State) F64ˢ
```
Clone returns a copy of the F64ˢ in a new memory pool.

#### func (F64ˢ) Copy

```go
func (dst F64ˢ) Copy(src F64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s F64ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (F64ˢ) Decoder

```go
func (s F64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (F64ˢ) ElementSize

```go
func (s F64ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F64ˢ points to.

#### func (F64ˢ) Encoder

```go
func (s F64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (F64ˢ) Index

```go
func (s F64ˢ) Index(i uint64, ϟs *gfxapi.State) F64ᵖ
```
Index returns a F64ᵖ to the i'th element in this F64ˢ.

#### func (F64ˢ) OnRead

```go
func (s F64ˢ) OnRead(ϟs *gfxapi.State) F64ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (F64ˢ) OnWrite

```go
func (s F64ˢ) OnWrite(ϟs *gfxapi.State) F64ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (F64ˢ) Range

```go
func (s F64ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (F64ˢ) Read

```go
func (s F64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []float64
```
Read reads and returns all the float64 elements in this F64ˢ.

#### func (F64ˢ) ResourceID

```go
func (s F64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (F64ˢ) Slice

```go
func (s F64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ
```
Slice returns a sub-slice from the F64ˢ using start and end indices.

#### func (F64ˢ) String

```go
func (s F64ˢ) String() string
```
String returns a string description of the F64ˢ slice.

#### func (F64ˢ) Write

```go
func (s F64ˢ) Write(src []float64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type F64ᵖ

```go
type F64ᵖ struct {
	binary.Generate
	memory.Pointer
}
```

F64ᵖ is a pointer to a float64 element.

#### func  NewF64ᵖ

```go
func NewF64ᵖ(addr uint64) F64ᵖ
```
NewF64ᵖ returns a F64ᵖ that points to addr in the application pool.

#### func (*F64ᵖ) Class

```go
func (*F64ᵖ) Class() binary.Class
```

#### func (F64ᵖ) ElementSize

```go
func (p F64ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that F64ᵖ points to.

#### func (F64ᵖ) OnRead

```go
func (p F64ᵖ) OnRead(ϟs *gfxapi.State) F64ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (F64ᵖ) OnWrite

```go
func (p F64ᵖ) OnWrite(ϟs *gfxapi.State) F64ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (F64ᵖ) Read

```go
func (p F64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) float64
```
Read reads and returns the float64 element at the pointer.

#### func (F64ᵖ) Slice

```go
func (p F64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) F64ˢ
```
Slice returns a new F64ˢ from the pointer using start and end indices.

#### func (F64ᵖ) Write

```go
func (p F64ᵖ) Write(value float64, ϟs *gfxapi.State)
```
Write writes value to the float64 element at the pointer.

#### type FaceMode

```go
type FaceMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FaceMode
//////////////////////////////////////////////////////////////////////////////

#### func (*FaceMode) Parse

```go
func (v *FaceMode) Parse(s string) error
```

#### func (FaceMode) String

```go
func (v FaceMode) String() string
```

#### type FaceModeːu32ᵐ

```go
type FaceModeːu32ᵐ map[FaceMode]uint32
```


#### func (FaceModeːu32ᵐ) Contains

```go
func (m FaceModeːu32ᵐ) Contains(key FaceMode) bool
```

#### func (FaceModeːu32ᵐ) Delete

```go
func (m FaceModeːu32ᵐ) Delete(key FaceMode)
```

#### func (FaceModeːu32ᵐ) Get

```go
func (m FaceModeːu32ᵐ) Get(key FaceMode) uint32
```

#### func (FaceModeːu32ᵐ) Range

```go
func (m FaceModeːu32ᵐ) Range() []uint32
```

#### type FaceOrientation

```go
type FaceOrientation uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FaceOrientation
//////////////////////////////////////////////////////////////////////////////

#### func (*FaceOrientation) Parse

```go
func (v *FaceOrientation) Parse(s string) error
```

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
func (c *FlushPostBuffer) API() gfxapi.ID
```

#### func (*FlushPostBuffer) AddRead

```go
func (a *FlushPostBuffer) AddRead(rng memory.Range, id binary.ID) *FlushPostBuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The FlushPostBuffer pointer is returned so that calls can be chained.

#### func (*FlushPostBuffer) AddWrite

```go
func (a *FlushPostBuffer) AddWrite(rng memory.Range, id binary.ID) *FlushPostBuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The FlushPostBuffer pointer is returned so that calls can be chained.

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
func (ϟa *FlushPostBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*FlushPostBuffer) Observations

```go
func (a *FlushPostBuffer) Observations() *atom.Observations
```

#### func (*FlushPostBuffer) Replay

```go
func (ϟa *FlushPostBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*FlushPostBuffer) String

```go
func (a *FlushPostBuffer) String() string
```

#### type Framebuffer

```go
type Framebuffer struct {
	binary.Generate
	CreatedAt   atom.ID
	Attachments FramebufferAttachmentːFramebufferAttachmentInfoᵐ
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

#### type FramebufferAttachment

```go
type FramebufferAttachment uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferAttachment
//////////////////////////////////////////////////////////////////////////////

#### func (*FramebufferAttachment) Parse

```go
func (v *FramebufferAttachment) Parse(s string) error
```

#### func (FramebufferAttachment) String

```go
func (v FramebufferAttachment) String() string
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

#### type FramebufferAttachmentParameter

```go
type FramebufferAttachmentParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferAttachmentParameter
//////////////////////////////////////////////////////////////////////////////

#### func (*FramebufferAttachmentParameter) Parse

```go
func (v *FramebufferAttachmentParameter) Parse(s string) error
```

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

#### func (*FramebufferAttachmentType) Parse

```go
func (v *FramebufferAttachmentType) Parse(s string) error
```

#### func (FramebufferAttachmentType) String

```go
func (v FramebufferAttachmentType) String() string
```

#### type FramebufferAttachmentːFramebufferAttachmentInfoᵐ

```go
type FramebufferAttachmentːFramebufferAttachmentInfoᵐ map[FramebufferAttachment]FramebufferAttachmentInfo
```


#### func (FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Contains

```go
func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Contains(key FramebufferAttachment) bool
```

#### func (FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Delete

```go
func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Delete(key FramebufferAttachment)
```

#### func (FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Get

```go
func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Get(key FramebufferAttachment) FramebufferAttachmentInfo
```

#### func (FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Range

```go
func (m FramebufferAttachmentːFramebufferAttachmentInfoᵐ) Range() []FramebufferAttachmentInfo
```

#### type FramebufferAttachmentˢ

```go
type FramebufferAttachmentˢ struct {
	binary.Generate
	SliceInfo
}
```

FramebufferAttachmentˢ is a slice of FramebufferAttachment.

#### func  AsFramebufferAttachmentˢ

```go
func AsFramebufferAttachmentˢ(s Slice, ϟs *gfxapi.State) FramebufferAttachmentˢ
```
AsFramebufferAttachmentˢ returns s cast to a FramebufferAttachmentˢ. The
returned slice length will be calculated so that the returned slice is no longer
(in bytes) than s.

#### func  MakeFramebufferAttachmentˢ

```go
func MakeFramebufferAttachmentˢ(count uint64, ϟs *gfxapi.State) FramebufferAttachmentˢ
```
MakeFramebufferAttachmentˢ returns a FramebufferAttachmentˢ backed by a new
memory pool.

#### func (*FramebufferAttachmentˢ) Class

```go
func (*FramebufferAttachmentˢ) Class() binary.Class
```

#### func (FramebufferAttachmentˢ) Clone

```go
func (s FramebufferAttachmentˢ) Clone(ϟs *gfxapi.State) FramebufferAttachmentˢ
```
Clone returns a copy of the FramebufferAttachmentˢ in a new memory pool.

#### func (FramebufferAttachmentˢ) Copy

```go
func (dst FramebufferAttachmentˢ) Copy(src FramebufferAttachmentˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s FramebufferAttachmentˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (FramebufferAttachmentˢ) Decoder

```go
func (s FramebufferAttachmentˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (FramebufferAttachmentˢ) ElementSize

```go
func (s FramebufferAttachmentˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that FramebufferAttachmentˢ
points to.

#### func (FramebufferAttachmentˢ) Encoder

```go
func (s FramebufferAttachmentˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (FramebufferAttachmentˢ) Index

```go
func (s FramebufferAttachmentˢ) Index(i uint64, ϟs *gfxapi.State) FramebufferAttachmentᵖ
```
Index returns a FramebufferAttachmentᵖ to the i'th element in this
FramebufferAttachmentˢ.

#### func (FramebufferAttachmentˢ) OnRead

```go
func (s FramebufferAttachmentˢ) OnRead(ϟs *gfxapi.State) FramebufferAttachmentˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (FramebufferAttachmentˢ) OnWrite

```go
func (s FramebufferAttachmentˢ) OnWrite(ϟs *gfxapi.State) FramebufferAttachmentˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (FramebufferAttachmentˢ) Range

```go
func (s FramebufferAttachmentˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (FramebufferAttachmentˢ) Read

```go
func (s FramebufferAttachmentˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []FramebufferAttachment
```
Read reads and returns all the FramebufferAttachment elements in this
FramebufferAttachmentˢ.

#### func (FramebufferAttachmentˢ) ResourceID

```go
func (s FramebufferAttachmentˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (FramebufferAttachmentˢ) Slice

```go
func (s FramebufferAttachmentˢ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferAttachmentˢ
```
Slice returns a sub-slice from the FramebufferAttachmentˢ using start and end
indices.

#### func (FramebufferAttachmentˢ) String

```go
func (s FramebufferAttachmentˢ) String() string
```
String returns a string description of the FramebufferAttachmentˢ slice.

#### func (FramebufferAttachmentˢ) Write

```go
func (s FramebufferAttachmentˢ) Write(src []FramebufferAttachment, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type FramebufferAttachmentᵖ

```go
type FramebufferAttachmentᵖ struct {
	binary.Generate
	memory.Pointer
}
```

FramebufferAttachmentᵖ is a pointer to a FramebufferAttachment element.

#### func  NewFramebufferAttachmentᵖ

```go
func NewFramebufferAttachmentᵖ(addr uint64) FramebufferAttachmentᵖ
```
NewFramebufferAttachmentᵖ returns a FramebufferAttachmentᵖ that points to addr
in the application pool.

#### func (*FramebufferAttachmentᵖ) Class

```go
func (*FramebufferAttachmentᵖ) Class() binary.Class
```

#### func (FramebufferAttachmentᵖ) ElementSize

```go
func (p FramebufferAttachmentᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that FramebufferAttachmentᵖ
points to.

#### func (FramebufferAttachmentᵖ) OnRead

```go
func (p FramebufferAttachmentᵖ) OnRead(ϟs *gfxapi.State) FramebufferAttachmentᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (FramebufferAttachmentᵖ) OnWrite

```go
func (p FramebufferAttachmentᵖ) OnWrite(ϟs *gfxapi.State) FramebufferAttachmentᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (FramebufferAttachmentᵖ) Read

```go
func (p FramebufferAttachmentᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) FramebufferAttachment
```
Read reads and returns the FramebufferAttachment element at the pointer.

#### func (FramebufferAttachmentᵖ) Slice

```go
func (p FramebufferAttachmentᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferAttachmentˢ
```
Slice returns a new FramebufferAttachmentˢ from the pointer using start and end
indices.

#### func (FramebufferAttachmentᵖ) Write

```go
func (p FramebufferAttachmentᵖ) Write(value FramebufferAttachment, ϟs *gfxapi.State)
```
Write writes value to the FramebufferAttachment element at the pointer.

#### type FramebufferAttachmentᶜᵖ

```go
type FramebufferAttachmentᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

FramebufferAttachmentᶜᵖ is a pointer to a FramebufferAttachment element.

#### func  NewFramebufferAttachmentᶜᵖ

```go
func NewFramebufferAttachmentᶜᵖ(addr uint64) FramebufferAttachmentᶜᵖ
```
NewFramebufferAttachmentᶜᵖ returns a FramebufferAttachmentᶜᵖ that points to addr
in the application pool.

#### func (*FramebufferAttachmentᶜᵖ) Class

```go
func (*FramebufferAttachmentᶜᵖ) Class() binary.Class
```

#### func (FramebufferAttachmentᶜᵖ) ElementSize

```go
func (p FramebufferAttachmentᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that FramebufferAttachmentᶜᵖ
points to.

#### func (FramebufferAttachmentᶜᵖ) OnRead

```go
func (p FramebufferAttachmentᶜᵖ) OnRead(ϟs *gfxapi.State) FramebufferAttachmentᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (FramebufferAttachmentᶜᵖ) OnWrite

```go
func (p FramebufferAttachmentᶜᵖ) OnWrite(ϟs *gfxapi.State) FramebufferAttachmentᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (FramebufferAttachmentᶜᵖ) Read

```go
func (p FramebufferAttachmentᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) FramebufferAttachment
```
Read reads and returns the FramebufferAttachment element at the pointer.

#### func (FramebufferAttachmentᶜᵖ) Slice

```go
func (p FramebufferAttachmentᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferAttachmentˢ
```
Slice returns a new FramebufferAttachmentˢ from the pointer using start and end
indices.

#### func (FramebufferAttachmentᶜᵖ) Write

```go
func (p FramebufferAttachmentᶜᵖ) Write(value FramebufferAttachment, ϟs *gfxapi.State)
```
Write writes value to the FramebufferAttachment element at the pointer.

#### type FramebufferId

```go
type FramebufferId uint32
```


#### type FramebufferIdːFramebufferʳᵐ

```go
type FramebufferIdːFramebufferʳᵐ map[FramebufferId](*Framebuffer)
```


#### func (FramebufferIdːFramebufferʳᵐ) Contains

```go
func (m FramebufferIdːFramebufferʳᵐ) Contains(key FramebufferId) bool
```

#### func (FramebufferIdːFramebufferʳᵐ) Delete

```go
func (m FramebufferIdːFramebufferʳᵐ) Delete(key FramebufferId)
```

#### func (FramebufferIdːFramebufferʳᵐ) Get

```go
func (m FramebufferIdːFramebufferʳᵐ) Get(key FramebufferId) *Framebuffer
```

#### func (FramebufferIdːFramebufferʳᵐ) Range

```go
func (m FramebufferIdːFramebufferʳᵐ) Range() [](*Framebuffer)
```

#### type FramebufferIdˢ

```go
type FramebufferIdˢ struct {
	binary.Generate
	SliceInfo
}
```

FramebufferIdˢ is a slice of FramebufferId.

#### func  AsFramebufferIdˢ

```go
func AsFramebufferIdˢ(s Slice, ϟs *gfxapi.State) FramebufferIdˢ
```
AsFramebufferIdˢ returns s cast to a FramebufferIdˢ. The returned slice length
will be calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeFramebufferIdˢ

```go
func MakeFramebufferIdˢ(count uint64, ϟs *gfxapi.State) FramebufferIdˢ
```
MakeFramebufferIdˢ returns a FramebufferIdˢ backed by a new memory pool.

#### func (*FramebufferIdˢ) Class

```go
func (*FramebufferIdˢ) Class() binary.Class
```

#### func (FramebufferIdˢ) Clone

```go
func (s FramebufferIdˢ) Clone(ϟs *gfxapi.State) FramebufferIdˢ
```
Clone returns a copy of the FramebufferIdˢ in a new memory pool.

#### func (FramebufferIdˢ) Copy

```go
func (dst FramebufferIdˢ) Copy(src FramebufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s FramebufferIdˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (FramebufferIdˢ) Decoder

```go
func (s FramebufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (FramebufferIdˢ) ElementSize

```go
func (s FramebufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that FramebufferIdˢ points
to.

#### func (FramebufferIdˢ) Encoder

```go
func (s FramebufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (FramebufferIdˢ) Index

```go
func (s FramebufferIdˢ) Index(i uint64, ϟs *gfxapi.State) FramebufferIdᵖ
```
Index returns a FramebufferIdᵖ to the i'th element in this FramebufferIdˢ.

#### func (FramebufferIdˢ) OnRead

```go
func (s FramebufferIdˢ) OnRead(ϟs *gfxapi.State) FramebufferIdˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (FramebufferIdˢ) OnWrite

```go
func (s FramebufferIdˢ) OnWrite(ϟs *gfxapi.State) FramebufferIdˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (FramebufferIdˢ) Range

```go
func (s FramebufferIdˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (FramebufferIdˢ) Read

```go
func (s FramebufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []FramebufferId
```
Read reads and returns all the FramebufferId elements in this FramebufferIdˢ.

#### func (FramebufferIdˢ) ResourceID

```go
func (s FramebufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (FramebufferIdˢ) Slice

```go
func (s FramebufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ
```
Slice returns a sub-slice from the FramebufferIdˢ using start and end indices.

#### func (FramebufferIdˢ) String

```go
func (s FramebufferIdˢ) String() string
```
String returns a string description of the FramebufferIdˢ slice.

#### func (FramebufferIdˢ) Write

```go
func (s FramebufferIdˢ) Write(src []FramebufferId, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type FramebufferIdᵖ

```go
type FramebufferIdᵖ struct {
	binary.Generate
	memory.Pointer
}
```

FramebufferIdᵖ is a pointer to a FramebufferId element.

#### func  NewFramebufferIdᵖ

```go
func NewFramebufferIdᵖ(addr uint64) FramebufferIdᵖ
```
NewFramebufferIdᵖ returns a FramebufferIdᵖ that points to addr in the
application pool.

#### func (*FramebufferIdᵖ) Class

```go
func (*FramebufferIdᵖ) Class() binary.Class
```

#### func (FramebufferIdᵖ) ElementSize

```go
func (p FramebufferIdᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that FramebufferIdᵖ points
to.

#### func (FramebufferIdᵖ) OnRead

```go
func (p FramebufferIdᵖ) OnRead(ϟs *gfxapi.State) FramebufferIdᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (FramebufferIdᵖ) OnWrite

```go
func (p FramebufferIdᵖ) OnWrite(ϟs *gfxapi.State) FramebufferIdᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (FramebufferIdᵖ) Read

```go
func (p FramebufferIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) FramebufferId
```
Read reads and returns the FramebufferId element at the pointer.

#### func (FramebufferIdᵖ) Slice

```go
func (p FramebufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ
```
Slice returns a new FramebufferIdˢ from the pointer using start and end indices.

#### func (FramebufferIdᵖ) Write

```go
func (p FramebufferIdᵖ) Write(value FramebufferId, ϟs *gfxapi.State)
```
Write writes value to the FramebufferId element at the pointer.

#### type FramebufferIdᶜᵖ

```go
type FramebufferIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

FramebufferIdᶜᵖ is a pointer to a FramebufferId element.

#### func  NewFramebufferIdᶜᵖ

```go
func NewFramebufferIdᶜᵖ(addr uint64) FramebufferIdᶜᵖ
```
NewFramebufferIdᶜᵖ returns a FramebufferIdᶜᵖ that points to addr in the
application pool.

#### func (*FramebufferIdᶜᵖ) Class

```go
func (*FramebufferIdᶜᵖ) Class() binary.Class
```

#### func (FramebufferIdᶜᵖ) ElementSize

```go
func (p FramebufferIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that FramebufferIdᶜᵖ points
to.

#### func (FramebufferIdᶜᵖ) OnRead

```go
func (p FramebufferIdᶜᵖ) OnRead(ϟs *gfxapi.State) FramebufferIdᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (FramebufferIdᶜᵖ) OnWrite

```go
func (p FramebufferIdᶜᵖ) OnWrite(ϟs *gfxapi.State) FramebufferIdᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (FramebufferIdᶜᵖ) Read

```go
func (p FramebufferIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) FramebufferId
```
Read reads and returns the FramebufferId element at the pointer.

#### func (FramebufferIdᶜᵖ) Slice

```go
func (p FramebufferIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) FramebufferIdˢ
```
Slice returns a new FramebufferIdˢ from the pointer using start and end indices.

#### func (FramebufferIdᶜᵖ) Write

```go
func (p FramebufferIdᶜᵖ) Write(value FramebufferId, ϟs *gfxapi.State)
```
Write writes value to the FramebufferId element at the pointer.

#### type FramebufferStatus

```go
type FramebufferStatus uint32
```

//////////////////////////////////////////////////////////////////////////////
enum FramebufferStatus
//////////////////////////////////////////////////////////////////////////////

#### func (*FramebufferStatus) Parse

```go
func (v *FramebufferStatus) Parse(s string) error
```

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

#### func (*FramebufferTarget) Parse

```go
func (v *FramebufferTarget) Parse(s string) error
```

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

#### func (*FramebufferTarget_GLES_2_0) Parse

```go
func (v *FramebufferTarget_GLES_2_0) Parse(s string) error
```

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

#### func (*FramebufferTarget_GLES_3_1) Parse

```go
func (v *FramebufferTarget_GLES_3_1) Parse(s string) error
```

#### func (FramebufferTarget_GLES_3_1) String

```go
func (v FramebufferTarget_GLES_3_1) String() string
```

#### type FramebufferTargetːFramebufferIdᵐ

```go
type FramebufferTargetːFramebufferIdᵐ map[FramebufferTarget]FramebufferId
```


#### func (FramebufferTargetːFramebufferIdᵐ) Contains

```go
func (m FramebufferTargetːFramebufferIdᵐ) Contains(key FramebufferTarget) bool
```

#### func (FramebufferTargetːFramebufferIdᵐ) Delete

```go
func (m FramebufferTargetːFramebufferIdᵐ) Delete(key FramebufferTarget)
```

#### func (FramebufferTargetːFramebufferIdᵐ) Get

```go
func (m FramebufferTargetːFramebufferIdᵐ) Get(key FramebufferTarget) FramebufferId
```

#### func (FramebufferTargetːFramebufferIdᵐ) Range

```go
func (m FramebufferTargetːFramebufferIdᵐ) Range() []FramebufferId
```

#### type GLXContext

```go
type GLXContext struct {
	binary.Generate
	memory.Pointer
}
```

GLXContext is a pointer to a void element.

#### func  NewGLXContext

```go
func NewGLXContext(addr uint64) GLXContext
```
NewGLXContext returns a GLXContext that points to addr in the application pool.

#### func (*GLXContext) Class

```go
func (*GLXContext) Class() binary.Class
```

#### func (GLXContext) ElementSize

```go
func (p GLXContext) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that GLXContext points to.

#### func (GLXContext) OnRead

```go
func (p GLXContext) OnRead(ϟs *gfxapi.State) GLXContext
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (GLXContext) OnWrite

```go
func (p GLXContext) OnWrite(ϟs *gfxapi.State) GLXContext
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (GLXContext) Slice

```go
func (p GLXContext) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type GLXContextːContextʳᵐ

```go
type GLXContextːContextʳᵐ map[GLXContext](*Context)
```


#### func (GLXContextːContextʳᵐ) Contains

```go
func (m GLXContextːContextʳᵐ) Contains(key GLXContext) bool
```

#### func (GLXContextːContextʳᵐ) Delete

```go
func (m GLXContextːContextʳᵐ) Delete(key GLXContext)
```

#### func (GLXContextːContextʳᵐ) Get

```go
func (m GLXContextːContextʳᵐ) Get(key GLXContext) *Context
```

#### func (GLXContextːContextʳᵐ) Range

```go
func (m GLXContextːContextʳᵐ) Range() [](*Context)
```

#### type GLXDrawable

```go
type GLXDrawable struct {
	binary.Generate
	memory.Pointer
}
```

GLXDrawable is a pointer to a void element.

#### func  NewGLXDrawable

```go
func NewGLXDrawable(addr uint64) GLXDrawable
```
NewGLXDrawable returns a GLXDrawable that points to addr in the application
pool.

#### func (*GLXDrawable) Class

```go
func (*GLXDrawable) Class() binary.Class
```

#### func (GLXDrawable) ElementSize

```go
func (p GLXDrawable) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that GLXDrawable points to.

#### func (GLXDrawable) OnRead

```go
func (p GLXDrawable) OnRead(ϟs *gfxapi.State) GLXDrawable
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (GLXDrawable) OnWrite

```go
func (p GLXDrawable) OnWrite(ϟs *gfxapi.State) GLXDrawable
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (GLXDrawable) Slice

```go
func (p GLXDrawable) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

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
func NewGlActiveTexture(Unit TextureUnit) *GlActiveTexture
```

#### func (*GlActiveTexture) API

```go
func (c *GlActiveTexture) API() gfxapi.ID
```

#### func (*GlActiveTexture) AddRead

```go
func (a *GlActiveTexture) AddRead(rng memory.Range, id binary.ID) *GlActiveTexture
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlActiveTexture pointer is returned so that calls can be chained.

#### func (*GlActiveTexture) AddWrite

```go
func (a *GlActiveTexture) AddWrite(rng memory.Range, id binary.ID) *GlActiveTexture
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlActiveTexture pointer is returned so that calls can be chained.

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
func (ϟa *GlActiveTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlActiveTexture) Observations

```go
func (a *GlActiveTexture) Observations() *atom.Observations
```

#### func (*GlActiveTexture) Replay

```go
func (ϟa *GlActiveTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlActiveTexture) String

```go
func (a *GlActiveTexture) String() string
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
func NewGlAttachShader(Program ProgramId, Shader ShaderId) *GlAttachShader
```

#### func (*GlAttachShader) API

```go
func (c *GlAttachShader) API() gfxapi.ID
```

#### func (*GlAttachShader) AddRead

```go
func (a *GlAttachShader) AddRead(rng memory.Range, id binary.ID) *GlAttachShader
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlAttachShader pointer is returned so that calls can be chained.

#### func (*GlAttachShader) AddWrite

```go
func (a *GlAttachShader) AddWrite(rng memory.Range, id binary.ID) *GlAttachShader
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlAttachShader pointer is returned so that calls can be chained.

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
func (ϟa *GlAttachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlAttachShader) Observations

```go
func (a *GlAttachShader) Observations() *atom.Observations
```

#### func (*GlAttachShader) Replay

```go
func (ϟa *GlAttachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlAttachShader) String

```go
func (a *GlAttachShader) String() string
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
func NewGlBeginQuery(Target QueryTarget, Query QueryId) *GlBeginQuery
```

#### func (*GlBeginQuery) API

```go
func (c *GlBeginQuery) API() gfxapi.ID
```

#### func (*GlBeginQuery) AddRead

```go
func (a *GlBeginQuery) AddRead(rng memory.Range, id binary.ID) *GlBeginQuery
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBeginQuery pointer is returned so that calls can be chained.

#### func (*GlBeginQuery) AddWrite

```go
func (a *GlBeginQuery) AddWrite(rng memory.Range, id binary.ID) *GlBeginQuery
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBeginQuery pointer is returned so that calls can be chained.

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
func (ϟa *GlBeginQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBeginQuery) Observations

```go
func (a *GlBeginQuery) Observations() *atom.Observations
```

#### func (*GlBeginQuery) Replay

```go
func (ϟa *GlBeginQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBeginQuery) String

```go
func (a *GlBeginQuery) String() string
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
func NewGlBeginQueryEXT(Target QueryTarget, Query QueryId) *GlBeginQueryEXT
```

#### func (*GlBeginQueryEXT) API

```go
func (c *GlBeginQueryEXT) API() gfxapi.ID
```

#### func (*GlBeginQueryEXT) AddRead

```go
func (a *GlBeginQueryEXT) AddRead(rng memory.Range, id binary.ID) *GlBeginQueryEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBeginQueryEXT pointer is returned so that calls can be chained.

#### func (*GlBeginQueryEXT) AddWrite

```go
func (a *GlBeginQueryEXT) AddWrite(rng memory.Range, id binary.ID) *GlBeginQueryEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBeginQueryEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlBeginQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBeginQueryEXT) Observations

```go
func (a *GlBeginQueryEXT) Observations() *atom.Observations
```

#### func (*GlBeginQueryEXT) Replay

```go
func (ϟa *GlBeginQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBeginQueryEXT) String

```go
func (a *GlBeginQueryEXT) String() string
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
func NewGlBindAttribLocation(Program ProgramId, Location AttributeLocation, Name string) *GlBindAttribLocation
```

#### func (*GlBindAttribLocation) API

```go
func (c *GlBindAttribLocation) API() gfxapi.ID
```

#### func (*GlBindAttribLocation) AddRead

```go
func (a *GlBindAttribLocation) AddRead(rng memory.Range, id binary.ID) *GlBindAttribLocation
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindAttribLocation pointer is returned so that calls can be
chained.

#### func (*GlBindAttribLocation) AddWrite

```go
func (a *GlBindAttribLocation) AddWrite(rng memory.Range, id binary.ID) *GlBindAttribLocation
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindAttribLocation pointer is returned so that calls can be
chained.

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
func (ϟa *GlBindAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindAttribLocation) Observations

```go
func (a *GlBindAttribLocation) Observations() *atom.Observations
```

#### func (*GlBindAttribLocation) Replay

```go
func (ϟa *GlBindAttribLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindAttribLocation) String

```go
func (a *GlBindAttribLocation) String() string
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
func NewGlBindBuffer(Target BufferTarget, Buffer BufferId) *GlBindBuffer
```

#### func (*GlBindBuffer) API

```go
func (c *GlBindBuffer) API() gfxapi.ID
```

#### func (*GlBindBuffer) AddRead

```go
func (a *GlBindBuffer) AddRead(rng memory.Range, id binary.ID) *GlBindBuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindBuffer pointer is returned so that calls can be chained.

#### func (*GlBindBuffer) AddWrite

```go
func (a *GlBindBuffer) AddWrite(rng memory.Range, id binary.ID) *GlBindBuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindBuffer pointer is returned so that calls can be chained.

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
func (ϟa *GlBindBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindBuffer) Observations

```go
func (a *GlBindBuffer) Observations() *atom.Observations
```

#### func (*GlBindBuffer) Replay

```go
func (ϟa *GlBindBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindBuffer) String

```go
func (a *GlBindBuffer) String() string
```

#### type GlBindBufferBase

```go
type GlBindBufferBase struct {
	binary.Generate

	Target IndexedBufferTarget
	Index  uint32
	Buffer BufferId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindBufferBase
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindBufferBase

```go
func NewGlBindBufferBase(Target IndexedBufferTarget, Index uint32, Buffer BufferId) *GlBindBufferBase
```

#### func (*GlBindBufferBase) API

```go
func (c *GlBindBufferBase) API() gfxapi.ID
```

#### func (*GlBindBufferBase) AddRead

```go
func (a *GlBindBufferBase) AddRead(rng memory.Range, id binary.ID) *GlBindBufferBase
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindBufferBase pointer is returned so that calls can be chained.

#### func (*GlBindBufferBase) AddWrite

```go
func (a *GlBindBufferBase) AddWrite(rng memory.Range, id binary.ID) *GlBindBufferBase
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindBufferBase pointer is returned so that calls can be chained.

#### func (*GlBindBufferBase) Class

```go
func (*GlBindBufferBase) Class() binary.Class
```

#### func (*GlBindBufferBase) Flags

```go
func (c *GlBindBufferBase) Flags() atom.Flags
```

#### func (*GlBindBufferBase) Mutate

```go
func (ϟa *GlBindBufferBase) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindBufferBase) Observations

```go
func (a *GlBindBufferBase) Observations() *atom.Observations
```

#### func (*GlBindBufferBase) Replay

```go
func (ϟa *GlBindBufferBase) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindBufferBase) String

```go
func (a *GlBindBufferBase) String() string
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
func NewGlBindFramebuffer(Target FramebufferTarget, Framebuffer FramebufferId) *GlBindFramebuffer
```

#### func (*GlBindFramebuffer) API

```go
func (c *GlBindFramebuffer) API() gfxapi.ID
```

#### func (*GlBindFramebuffer) AddRead

```go
func (a *GlBindFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlBindFramebuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindFramebuffer pointer is returned so that calls can be chained.

#### func (*GlBindFramebuffer) AddWrite

```go
func (a *GlBindFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlBindFramebuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindFramebuffer pointer is returned so that calls can be chained.

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
func (ϟa *GlBindFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindFramebuffer) Observations

```go
func (a *GlBindFramebuffer) Observations() *atom.Observations
```

#### func (*GlBindFramebuffer) Replay

```go
func (ϟa *GlBindFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindFramebuffer) String

```go
func (a *GlBindFramebuffer) String() string
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
func NewGlBindRenderbuffer(Target RenderbufferTarget, Renderbuffer RenderbufferId) *GlBindRenderbuffer
```

#### func (*GlBindRenderbuffer) API

```go
func (c *GlBindRenderbuffer) API() gfxapi.ID
```

#### func (*GlBindRenderbuffer) AddRead

```go
func (a *GlBindRenderbuffer) AddRead(rng memory.Range, id binary.ID) *GlBindRenderbuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindRenderbuffer pointer is returned so that calls can be
chained.

#### func (*GlBindRenderbuffer) AddWrite

```go
func (a *GlBindRenderbuffer) AddWrite(rng memory.Range, id binary.ID) *GlBindRenderbuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindRenderbuffer pointer is returned so that calls can be
chained.

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
func (ϟa *GlBindRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindRenderbuffer) Observations

```go
func (a *GlBindRenderbuffer) Observations() *atom.Observations
```

#### func (*GlBindRenderbuffer) Replay

```go
func (ϟa *GlBindRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindRenderbuffer) String

```go
func (a *GlBindRenderbuffer) String() string
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
func NewGlBindTexture(Target TextureTarget, Texture TextureId) *GlBindTexture
```

#### func (*GlBindTexture) API

```go
func (c *GlBindTexture) API() gfxapi.ID
```

#### func (*GlBindTexture) AddRead

```go
func (a *GlBindTexture) AddRead(rng memory.Range, id binary.ID) *GlBindTexture
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindTexture pointer is returned so that calls can be chained.

#### func (*GlBindTexture) AddWrite

```go
func (a *GlBindTexture) AddWrite(rng memory.Range, id binary.ID) *GlBindTexture
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindTexture pointer is returned so that calls can be chained.

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
func (ϟa *GlBindTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindTexture) Observations

```go
func (a *GlBindTexture) Observations() *atom.Observations
```

#### func (*GlBindTexture) Replay

```go
func (ϟa *GlBindTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindTexture) String

```go
func (a *GlBindTexture) String() string
```

#### type GlBindVertexArray

```go
type GlBindVertexArray struct {
	binary.Generate

	Array VertexArrayId
}
```

//////////////////////////////////////////////////////////////////////////////
GlBindVertexArray
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBindVertexArray

```go
func NewGlBindVertexArray(Array VertexArrayId) *GlBindVertexArray
```

#### func (*GlBindVertexArray) API

```go
func (c *GlBindVertexArray) API() gfxapi.ID
```

#### func (*GlBindVertexArray) AddRead

```go
func (a *GlBindVertexArray) AddRead(rng memory.Range, id binary.ID) *GlBindVertexArray
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindVertexArray pointer is returned so that calls can be chained.

#### func (*GlBindVertexArray) AddWrite

```go
func (a *GlBindVertexArray) AddWrite(rng memory.Range, id binary.ID) *GlBindVertexArray
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindVertexArray pointer is returned so that calls can be chained.

#### func (*GlBindVertexArray) Class

```go
func (*GlBindVertexArray) Class() binary.Class
```

#### func (*GlBindVertexArray) Flags

```go
func (c *GlBindVertexArray) Flags() atom.Flags
```

#### func (*GlBindVertexArray) Mutate

```go
func (ϟa *GlBindVertexArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindVertexArray) Observations

```go
func (a *GlBindVertexArray) Observations() *atom.Observations
```

#### func (*GlBindVertexArray) Replay

```go
func (ϟa *GlBindVertexArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindVertexArray) String

```go
func (a *GlBindVertexArray) String() string
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
func NewGlBindVertexArrayOES(Array VertexArrayId) *GlBindVertexArrayOES
```

#### func (*GlBindVertexArrayOES) API

```go
func (c *GlBindVertexArrayOES) API() gfxapi.ID
```

#### func (*GlBindVertexArrayOES) AddRead

```go
func (a *GlBindVertexArrayOES) AddRead(rng memory.Range, id binary.ID) *GlBindVertexArrayOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBindVertexArrayOES pointer is returned so that calls can be
chained.

#### func (*GlBindVertexArrayOES) AddWrite

```go
func (a *GlBindVertexArrayOES) AddWrite(rng memory.Range, id binary.ID) *GlBindVertexArrayOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBindVertexArrayOES pointer is returned so that calls can be
chained.

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
func (ϟa *GlBindVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBindVertexArrayOES) Observations

```go
func (a *GlBindVertexArrayOES) Observations() *atom.Observations
```

#### func (*GlBindVertexArrayOES) Replay

```go
func (ϟa *GlBindVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBindVertexArrayOES) String

```go
func (a *GlBindVertexArrayOES) String() string
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
func NewGlBlendColor(Red float32, Green float32, Blue float32, Alpha float32) *GlBlendColor
```

#### func (*GlBlendColor) API

```go
func (c *GlBlendColor) API() gfxapi.ID
```

#### func (*GlBlendColor) AddRead

```go
func (a *GlBlendColor) AddRead(rng memory.Range, id binary.ID) *GlBlendColor
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBlendColor pointer is returned so that calls can be chained.

#### func (*GlBlendColor) AddWrite

```go
func (a *GlBlendColor) AddWrite(rng memory.Range, id binary.ID) *GlBlendColor
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBlendColor pointer is returned so that calls can be chained.

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
func (ϟa *GlBlendColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBlendColor) Observations

```go
func (a *GlBlendColor) Observations() *atom.Observations
```

#### func (*GlBlendColor) Replay

```go
func (ϟa *GlBlendColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBlendColor) String

```go
func (a *GlBlendColor) String() string
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
func NewGlBlendEquation(Equation BlendEquation) *GlBlendEquation
```

#### func (*GlBlendEquation) API

```go
func (c *GlBlendEquation) API() gfxapi.ID
```

#### func (*GlBlendEquation) AddRead

```go
func (a *GlBlendEquation) AddRead(rng memory.Range, id binary.ID) *GlBlendEquation
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBlendEquation pointer is returned so that calls can be chained.

#### func (*GlBlendEquation) AddWrite

```go
func (a *GlBlendEquation) AddWrite(rng memory.Range, id binary.ID) *GlBlendEquation
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBlendEquation pointer is returned so that calls can be chained.

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
func (ϟa *GlBlendEquation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBlendEquation) Observations

```go
func (a *GlBlendEquation) Observations() *atom.Observations
```

#### func (*GlBlendEquation) Replay

```go
func (ϟa *GlBlendEquation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBlendEquation) String

```go
func (a *GlBlendEquation) String() string
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
func NewGlBlendEquationSeparate(Rgb BlendEquation, Alpha BlendEquation) *GlBlendEquationSeparate
```

#### func (*GlBlendEquationSeparate) API

```go
func (c *GlBlendEquationSeparate) API() gfxapi.ID
```

#### func (*GlBlendEquationSeparate) AddRead

```go
func (a *GlBlendEquationSeparate) AddRead(rng memory.Range, id binary.ID) *GlBlendEquationSeparate
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBlendEquationSeparate pointer is returned so that calls can be
chained.

#### func (*GlBlendEquationSeparate) AddWrite

```go
func (a *GlBlendEquationSeparate) AddWrite(rng memory.Range, id binary.ID) *GlBlendEquationSeparate
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBlendEquationSeparate pointer is returned so that calls can be
chained.

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
func (ϟa *GlBlendEquationSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBlendEquationSeparate) Observations

```go
func (a *GlBlendEquationSeparate) Observations() *atom.Observations
```

#### func (*GlBlendEquationSeparate) Replay

```go
func (ϟa *GlBlendEquationSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBlendEquationSeparate) String

```go
func (a *GlBlendEquationSeparate) String() string
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
func NewGlBlendFunc(Src_factor BlendFactor, Dst_factor BlendFactor) *GlBlendFunc
```

#### func (*GlBlendFunc) API

```go
func (c *GlBlendFunc) API() gfxapi.ID
```

#### func (*GlBlendFunc) AddRead

```go
func (a *GlBlendFunc) AddRead(rng memory.Range, id binary.ID) *GlBlendFunc
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBlendFunc pointer is returned so that calls can be chained.

#### func (*GlBlendFunc) AddWrite

```go
func (a *GlBlendFunc) AddWrite(rng memory.Range, id binary.ID) *GlBlendFunc
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBlendFunc pointer is returned so that calls can be chained.

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
func (ϟa *GlBlendFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBlendFunc) Observations

```go
func (a *GlBlendFunc) Observations() *atom.Observations
```

#### func (*GlBlendFunc) Replay

```go
func (ϟa *GlBlendFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBlendFunc) String

```go
func (a *GlBlendFunc) String() string
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
func NewGlBlendFuncSeparate(Src_factor_rgb BlendFactor, Dst_factor_rgb BlendFactor, Src_factor_alpha BlendFactor, Dst_factor_alpha BlendFactor) *GlBlendFuncSeparate
```

#### func (*GlBlendFuncSeparate) API

```go
func (c *GlBlendFuncSeparate) API() gfxapi.ID
```

#### func (*GlBlendFuncSeparate) AddRead

```go
func (a *GlBlendFuncSeparate) AddRead(rng memory.Range, id binary.ID) *GlBlendFuncSeparate
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBlendFuncSeparate pointer is returned so that calls can be
chained.

#### func (*GlBlendFuncSeparate) AddWrite

```go
func (a *GlBlendFuncSeparate) AddWrite(rng memory.Range, id binary.ID) *GlBlendFuncSeparate
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBlendFuncSeparate pointer is returned so that calls can be
chained.

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
func (ϟa *GlBlendFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBlendFuncSeparate) Observations

```go
func (a *GlBlendFuncSeparate) Observations() *atom.Observations
```

#### func (*GlBlendFuncSeparate) Replay

```go
func (ϟa *GlBlendFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBlendFuncSeparate) String

```go
func (a *GlBlendFuncSeparate) String() string
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
func NewGlBlitFramebuffer(SrcX0 int32, SrcY0 int32, SrcX1 int32, SrcY1 int32, DstX0 int32, DstY0 int32, DstX1 int32, DstY1 int32, Mask ClearMask, Filter TextureFilterMode) *GlBlitFramebuffer
```

#### func (*GlBlitFramebuffer) API

```go
func (c *GlBlitFramebuffer) API() gfxapi.ID
```

#### func (*GlBlitFramebuffer) AddRead

```go
func (a *GlBlitFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlBlitFramebuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBlitFramebuffer pointer is returned so that calls can be chained.

#### func (*GlBlitFramebuffer) AddWrite

```go
func (a *GlBlitFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlBlitFramebuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBlitFramebuffer pointer is returned so that calls can be chained.

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
func (ϟa *GlBlitFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBlitFramebuffer) Observations

```go
func (a *GlBlitFramebuffer) Observations() *atom.Observations
```

#### func (*GlBlitFramebuffer) Replay

```go
func (ϟa *GlBlitFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBlitFramebuffer) String

```go
func (a *GlBlitFramebuffer) String() string
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
func NewGlBufferData(Target BufferTarget, Size int32, Data memory.Pointer, Usage BufferUsage) *GlBufferData
```

#### func (*GlBufferData) API

```go
func (c *GlBufferData) API() gfxapi.ID
```

#### func (*GlBufferData) AddRead

```go
func (a *GlBufferData) AddRead(rng memory.Range, id binary.ID) *GlBufferData
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBufferData pointer is returned so that calls can be chained.

#### func (*GlBufferData) AddWrite

```go
func (a *GlBufferData) AddWrite(rng memory.Range, id binary.ID) *GlBufferData
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBufferData pointer is returned so that calls can be chained.

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
func (ϟa *GlBufferData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBufferData) Observations

```go
func (a *GlBufferData) Observations() *atom.Observations
```

#### func (*GlBufferData) Replay

```go
func (ϟa *GlBufferData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBufferData) String

```go
func (a *GlBufferData) String() string
```

#### type GlBufferSubData

```go
type GlBufferSubData struct {
	binary.Generate

	Target BufferTarget
	Offset int32
	Size   int32
	Data   BufferDataPointer
}
```

//////////////////////////////////////////////////////////////////////////////
GlBufferSubData
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlBufferSubData

```go
func NewGlBufferSubData(Target BufferTarget, Offset int32, Size int32, Data memory.Pointer) *GlBufferSubData
```

#### func (*GlBufferSubData) API

```go
func (c *GlBufferSubData) API() gfxapi.ID
```

#### func (*GlBufferSubData) AddRead

```go
func (a *GlBufferSubData) AddRead(rng memory.Range, id binary.ID) *GlBufferSubData
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlBufferSubData pointer is returned so that calls can be chained.

#### func (*GlBufferSubData) AddWrite

```go
func (a *GlBufferSubData) AddWrite(rng memory.Range, id binary.ID) *GlBufferSubData
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlBufferSubData pointer is returned so that calls can be chained.

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
func (ϟa *GlBufferSubData) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlBufferSubData) Observations

```go
func (a *GlBufferSubData) Observations() *atom.Observations
```

#### func (*GlBufferSubData) Replay

```go
func (ϟa *GlBufferSubData) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlBufferSubData) String

```go
func (a *GlBufferSubData) String() string
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
func NewGlCheckFramebufferStatus(Target FramebufferTarget, Result FramebufferStatus) *GlCheckFramebufferStatus
```

#### func (*GlCheckFramebufferStatus) API

```go
func (c *GlCheckFramebufferStatus) API() gfxapi.ID
```

#### func (*GlCheckFramebufferStatus) AddRead

```go
func (a *GlCheckFramebufferStatus) AddRead(rng memory.Range, id binary.ID) *GlCheckFramebufferStatus
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCheckFramebufferStatus pointer is returned so that calls can be
chained.

#### func (*GlCheckFramebufferStatus) AddWrite

```go
func (a *GlCheckFramebufferStatus) AddWrite(rng memory.Range, id binary.ID) *GlCheckFramebufferStatus
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCheckFramebufferStatus pointer is returned so that calls can be
chained.

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
func (ϟa *GlCheckFramebufferStatus) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCheckFramebufferStatus) Observations

```go
func (a *GlCheckFramebufferStatus) Observations() *atom.Observations
```

#### func (*GlCheckFramebufferStatus) Replay

```go
func (ϟa *GlCheckFramebufferStatus) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCheckFramebufferStatus) String

```go
func (a *GlCheckFramebufferStatus) String() string
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
func NewGlClear(Mask ClearMask) *GlClear
```

#### func (*GlClear) API

```go
func (c *GlClear) API() gfxapi.ID
```

#### func (*GlClear) AddRead

```go
func (a *GlClear) AddRead(rng memory.Range, id binary.ID) *GlClear
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlClear pointer is returned so that calls can be chained.

#### func (*GlClear) AddWrite

```go
func (a *GlClear) AddWrite(rng memory.Range, id binary.ID) *GlClear
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlClear pointer is returned so that calls can be chained.

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
func (ϟa *GlClear) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlClear) Observations

```go
func (a *GlClear) Observations() *atom.Observations
```

#### func (*GlClear) Replay

```go
func (ϟa *GlClear) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlClear) String

```go
func (a *GlClear) String() string
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
func NewGlClearColor(R float32, G float32, B float32, A float32) *GlClearColor
```

#### func (*GlClearColor) API

```go
func (c *GlClearColor) API() gfxapi.ID
```

#### func (*GlClearColor) AddRead

```go
func (a *GlClearColor) AddRead(rng memory.Range, id binary.ID) *GlClearColor
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlClearColor pointer is returned so that calls can be chained.

#### func (*GlClearColor) AddWrite

```go
func (a *GlClearColor) AddWrite(rng memory.Range, id binary.ID) *GlClearColor
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlClearColor pointer is returned so that calls can be chained.

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
func (ϟa *GlClearColor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlClearColor) Observations

```go
func (a *GlClearColor) Observations() *atom.Observations
```

#### func (*GlClearColor) Replay

```go
func (ϟa *GlClearColor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlClearColor) String

```go
func (a *GlClearColor) String() string
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
func NewGlClearDepthf(Depth float32) *GlClearDepthf
```

#### func (*GlClearDepthf) API

```go
func (c *GlClearDepthf) API() gfxapi.ID
```

#### func (*GlClearDepthf) AddRead

```go
func (a *GlClearDepthf) AddRead(rng memory.Range, id binary.ID) *GlClearDepthf
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlClearDepthf pointer is returned so that calls can be chained.

#### func (*GlClearDepthf) AddWrite

```go
func (a *GlClearDepthf) AddWrite(rng memory.Range, id binary.ID) *GlClearDepthf
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlClearDepthf pointer is returned so that calls can be chained.

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
func (ϟa *GlClearDepthf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlClearDepthf) Observations

```go
func (a *GlClearDepthf) Observations() *atom.Observations
```

#### func (*GlClearDepthf) Replay

```go
func (ϟa *GlClearDepthf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlClearDepthf) String

```go
func (a *GlClearDepthf) String() string
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
func NewGlClearStencil(Stencil int32) *GlClearStencil
```

#### func (*GlClearStencil) API

```go
func (c *GlClearStencil) API() gfxapi.ID
```

#### func (*GlClearStencil) AddRead

```go
func (a *GlClearStencil) AddRead(rng memory.Range, id binary.ID) *GlClearStencil
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlClearStencil pointer is returned so that calls can be chained.

#### func (*GlClearStencil) AddWrite

```go
func (a *GlClearStencil) AddWrite(rng memory.Range, id binary.ID) *GlClearStencil
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlClearStencil pointer is returned so that calls can be chained.

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
func (ϟa *GlClearStencil) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlClearStencil) Observations

```go
func (a *GlClearStencil) Observations() *atom.Observations
```

#### func (*GlClearStencil) Replay

```go
func (ϟa *GlClearStencil) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlClearStencil) String

```go
func (a *GlClearStencil) String() string
```

#### type GlClientWaitSync

```go
type GlClientWaitSync struct {
	binary.Generate

	Sync      SyncObject
	SyncFlags SyncFlags
	Timeout   uint64
	Result    ClientWaitSyncSignal
}
```

//////////////////////////////////////////////////////////////////////////////
GlClientWaitSync
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlClientWaitSync

```go
func NewGlClientWaitSync(Sync SyncObject, SyncFlags SyncFlags, Timeout uint64, Result ClientWaitSyncSignal) *GlClientWaitSync
```

#### func (*GlClientWaitSync) API

```go
func (c *GlClientWaitSync) API() gfxapi.ID
```

#### func (*GlClientWaitSync) AddRead

```go
func (a *GlClientWaitSync) AddRead(rng memory.Range, id binary.ID) *GlClientWaitSync
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlClientWaitSync pointer is returned so that calls can be chained.

#### func (*GlClientWaitSync) AddWrite

```go
func (a *GlClientWaitSync) AddWrite(rng memory.Range, id binary.ID) *GlClientWaitSync
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlClientWaitSync pointer is returned so that calls can be chained.

#### func (*GlClientWaitSync) Class

```go
func (*GlClientWaitSync) Class() binary.Class
```

#### func (*GlClientWaitSync) Flags

```go
func (c *GlClientWaitSync) Flags() atom.Flags
```

#### func (*GlClientWaitSync) Mutate

```go
func (ϟa *GlClientWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlClientWaitSync) Observations

```go
func (a *GlClientWaitSync) Observations() *atom.Observations
```

#### func (*GlClientWaitSync) Replay

```go
func (ϟa *GlClientWaitSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlClientWaitSync) String

```go
func (a *GlClientWaitSync) String() string
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
func NewGlColorMask(Red bool, Green bool, Blue bool, Alpha bool) *GlColorMask
```

#### func (*GlColorMask) API

```go
func (c *GlColorMask) API() gfxapi.ID
```

#### func (*GlColorMask) AddRead

```go
func (a *GlColorMask) AddRead(rng memory.Range, id binary.ID) *GlColorMask
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlColorMask pointer is returned so that calls can be chained.

#### func (*GlColorMask) AddWrite

```go
func (a *GlColorMask) AddWrite(rng memory.Range, id binary.ID) *GlColorMask
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlColorMask pointer is returned so that calls can be chained.

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
func (ϟa *GlColorMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlColorMask) Observations

```go
func (a *GlColorMask) Observations() *atom.Observations
```

#### func (*GlColorMask) Replay

```go
func (ϟa *GlColorMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlColorMask) String

```go
func (a *GlColorMask) String() string
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
func NewGlCompileShader(Shader ShaderId) *GlCompileShader
```

#### func (*GlCompileShader) API

```go
func (c *GlCompileShader) API() gfxapi.ID
```

#### func (*GlCompileShader) AddRead

```go
func (a *GlCompileShader) AddRead(rng memory.Range, id binary.ID) *GlCompileShader
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCompileShader pointer is returned so that calls can be chained.

#### func (*GlCompileShader) AddWrite

```go
func (a *GlCompileShader) AddWrite(rng memory.Range, id binary.ID) *GlCompileShader
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCompileShader pointer is returned so that calls can be chained.

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
func (ϟa *GlCompileShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCompileShader) Observations

```go
func (a *GlCompileShader) Observations() *atom.Observations
```

#### func (*GlCompileShader) Replay

```go
func (ϟa *GlCompileShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCompileShader) String

```go
func (a *GlCompileShader) String() string
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
func NewGlCompressedTexImage2D(Target TextureImageTarget, Level int32, Format CompressedTexelFormat, Width int32, Height int32, Border int32, Image_size int32, Data memory.Pointer) *GlCompressedTexImage2D
```

#### func (*GlCompressedTexImage2D) API

```go
func (c *GlCompressedTexImage2D) API() gfxapi.ID
```

#### func (*GlCompressedTexImage2D) AddRead

```go
func (a *GlCompressedTexImage2D) AddRead(rng memory.Range, id binary.ID) *GlCompressedTexImage2D
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCompressedTexImage2D pointer is returned so that calls can be
chained.

#### func (*GlCompressedTexImage2D) AddWrite

```go
func (a *GlCompressedTexImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCompressedTexImage2D
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCompressedTexImage2D pointer is returned so that calls can be
chained.

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
func (ϟa *GlCompressedTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCompressedTexImage2D) Observations

```go
func (a *GlCompressedTexImage2D) Observations() *atom.Observations
```

#### func (*GlCompressedTexImage2D) Replay

```go
func (ϟa *GlCompressedTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCompressedTexImage2D) String

```go
func (a *GlCompressedTexImage2D) String() string
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
func NewGlCompressedTexSubImage2D(Target TextureImageTarget, Level int32, Xoffset int32, Yoffset int32, Width int32, Height int32, Format CompressedTexelFormat, Image_size int32, Data memory.Pointer) *GlCompressedTexSubImage2D
```

#### func (*GlCompressedTexSubImage2D) API

```go
func (c *GlCompressedTexSubImage2D) API() gfxapi.ID
```

#### func (*GlCompressedTexSubImage2D) AddRead

```go
func (a *GlCompressedTexSubImage2D) AddRead(rng memory.Range, id binary.ID) *GlCompressedTexSubImage2D
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCompressedTexSubImage2D pointer is returned so that calls can be
chained.

#### func (*GlCompressedTexSubImage2D) AddWrite

```go
func (a *GlCompressedTexSubImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCompressedTexSubImage2D
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCompressedTexSubImage2D pointer is returned so that calls can be
chained.

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
func (ϟa *GlCompressedTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCompressedTexSubImage2D) Observations

```go
func (a *GlCompressedTexSubImage2D) Observations() *atom.Observations
```

#### func (*GlCompressedTexSubImage2D) Replay

```go
func (ϟa *GlCompressedTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCompressedTexSubImage2D) String

```go
func (a *GlCompressedTexSubImage2D) String() string
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
func NewGlCopyTexImage2D(Target TextureImageTarget, Level int32, Format TexelFormat, X int32, Y int32, Width int32, Height int32, Border int32) *GlCopyTexImage2D
```

#### func (*GlCopyTexImage2D) API

```go
func (c *GlCopyTexImage2D) API() gfxapi.ID
```

#### func (*GlCopyTexImage2D) AddRead

```go
func (a *GlCopyTexImage2D) AddRead(rng memory.Range, id binary.ID) *GlCopyTexImage2D
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCopyTexImage2D pointer is returned so that calls can be chained.

#### func (*GlCopyTexImage2D) AddWrite

```go
func (a *GlCopyTexImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCopyTexImage2D
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCopyTexImage2D pointer is returned so that calls can be chained.

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
func (ϟa *GlCopyTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCopyTexImage2D) Observations

```go
func (a *GlCopyTexImage2D) Observations() *atom.Observations
```

#### func (*GlCopyTexImage2D) Replay

```go
func (ϟa *GlCopyTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCopyTexImage2D) String

```go
func (a *GlCopyTexImage2D) String() string
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
func NewGlCopyTexSubImage2D(Target TextureImageTarget, Level int32, Xoffset int32, Yoffset int32, X int32, Y int32, Width int32, Height int32) *GlCopyTexSubImage2D
```

#### func (*GlCopyTexSubImage2D) API

```go
func (c *GlCopyTexSubImage2D) API() gfxapi.ID
```

#### func (*GlCopyTexSubImage2D) AddRead

```go
func (a *GlCopyTexSubImage2D) AddRead(rng memory.Range, id binary.ID) *GlCopyTexSubImage2D
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCopyTexSubImage2D pointer is returned so that calls can be
chained.

#### func (*GlCopyTexSubImage2D) AddWrite

```go
func (a *GlCopyTexSubImage2D) AddWrite(rng memory.Range, id binary.ID) *GlCopyTexSubImage2D
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCopyTexSubImage2D pointer is returned so that calls can be
chained.

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
func (ϟa *GlCopyTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCopyTexSubImage2D) Observations

```go
func (a *GlCopyTexSubImage2D) Observations() *atom.Observations
```

#### func (*GlCopyTexSubImage2D) Replay

```go
func (ϟa *GlCopyTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCopyTexSubImage2D) String

```go
func (a *GlCopyTexSubImage2D) String() string
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
func NewGlCreateProgram(Result ProgramId) *GlCreateProgram
```

#### func (*GlCreateProgram) API

```go
func (c *GlCreateProgram) API() gfxapi.ID
```

#### func (*GlCreateProgram) AddRead

```go
func (a *GlCreateProgram) AddRead(rng memory.Range, id binary.ID) *GlCreateProgram
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCreateProgram pointer is returned so that calls can be chained.

#### func (*GlCreateProgram) AddWrite

```go
func (a *GlCreateProgram) AddWrite(rng memory.Range, id binary.ID) *GlCreateProgram
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCreateProgram pointer is returned so that calls can be chained.

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
func (ϟa *GlCreateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCreateProgram) Observations

```go
func (a *GlCreateProgram) Observations() *atom.Observations
```

#### func (*GlCreateProgram) Replay

```go
func (ϟa *GlCreateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCreateProgram) String

```go
func (a *GlCreateProgram) String() string
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
func NewGlCreateShader(Type ShaderType, Result ShaderId) *GlCreateShader
```

#### func (*GlCreateShader) API

```go
func (c *GlCreateShader) API() gfxapi.ID
```

#### func (*GlCreateShader) AddRead

```go
func (a *GlCreateShader) AddRead(rng memory.Range, id binary.ID) *GlCreateShader
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCreateShader pointer is returned so that calls can be chained.

#### func (*GlCreateShader) AddWrite

```go
func (a *GlCreateShader) AddWrite(rng memory.Range, id binary.ID) *GlCreateShader
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCreateShader pointer is returned so that calls can be chained.

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
func (ϟa *GlCreateShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCreateShader) Observations

```go
func (a *GlCreateShader) Observations() *atom.Observations
```

#### func (*GlCreateShader) Replay

```go
func (ϟa *GlCreateShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCreateShader) String

```go
func (a *GlCreateShader) String() string
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
func NewGlCullFace(Mode FaceMode) *GlCullFace
```

#### func (*GlCullFace) API

```go
func (c *GlCullFace) API() gfxapi.ID
```

#### func (*GlCullFace) AddRead

```go
func (a *GlCullFace) AddRead(rng memory.Range, id binary.ID) *GlCullFace
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlCullFace pointer is returned so that calls can be chained.

#### func (*GlCullFace) AddWrite

```go
func (a *GlCullFace) AddWrite(rng memory.Range, id binary.ID) *GlCullFace
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlCullFace pointer is returned so that calls can be chained.

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
func (ϟa *GlCullFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlCullFace) Observations

```go
func (a *GlCullFace) Observations() *atom.Observations
```

#### func (*GlCullFace) Replay

```go
func (ϟa *GlCullFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlCullFace) String

```go
func (a *GlCullFace) String() string
```

#### type GlDeleteBuffers

```go
type GlDeleteBuffers struct {
	binary.Generate

	Count   int32
	Buffers BufferIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteBuffers

```go
func NewGlDeleteBuffers(Count int32, Buffers memory.Pointer) *GlDeleteBuffers
```

#### func (*GlDeleteBuffers) API

```go
func (c *GlDeleteBuffers) API() gfxapi.ID
```

#### func (*GlDeleteBuffers) AddRead

```go
func (a *GlDeleteBuffers) AddRead(rng memory.Range, id binary.ID) *GlDeleteBuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteBuffers pointer is returned so that calls can be chained.

#### func (*GlDeleteBuffers) AddWrite

```go
func (a *GlDeleteBuffers) AddWrite(rng memory.Range, id binary.ID) *GlDeleteBuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteBuffers pointer is returned so that calls can be chained.

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
func (ϟa *GlDeleteBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteBuffers) Observations

```go
func (a *GlDeleteBuffers) Observations() *atom.Observations
```

#### func (*GlDeleteBuffers) Replay

```go
func (ϟa *GlDeleteBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteBuffers) String

```go
func (a *GlDeleteBuffers) String() string
```

#### type GlDeleteFramebuffers

```go
type GlDeleteFramebuffers struct {
	binary.Generate

	Count        int32
	Framebuffers FramebufferIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteFramebuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteFramebuffers

```go
func NewGlDeleteFramebuffers(Count int32, Framebuffers memory.Pointer) *GlDeleteFramebuffers
```

#### func (*GlDeleteFramebuffers) API

```go
func (c *GlDeleteFramebuffers) API() gfxapi.ID
```

#### func (*GlDeleteFramebuffers) AddRead

```go
func (a *GlDeleteFramebuffers) AddRead(rng memory.Range, id binary.ID) *GlDeleteFramebuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteFramebuffers pointer is returned so that calls can be
chained.

#### func (*GlDeleteFramebuffers) AddWrite

```go
func (a *GlDeleteFramebuffers) AddWrite(rng memory.Range, id binary.ID) *GlDeleteFramebuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteFramebuffers pointer is returned so that calls can be
chained.

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
func (ϟa *GlDeleteFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteFramebuffers) Observations

```go
func (a *GlDeleteFramebuffers) Observations() *atom.Observations
```

#### func (*GlDeleteFramebuffers) Replay

```go
func (ϟa *GlDeleteFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteFramebuffers) String

```go
func (a *GlDeleteFramebuffers) String() string
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
func NewGlDeleteProgram(Program ProgramId) *GlDeleteProgram
```

#### func (*GlDeleteProgram) API

```go
func (c *GlDeleteProgram) API() gfxapi.ID
```

#### func (*GlDeleteProgram) AddRead

```go
func (a *GlDeleteProgram) AddRead(rng memory.Range, id binary.ID) *GlDeleteProgram
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteProgram pointer is returned so that calls can be chained.

#### func (*GlDeleteProgram) AddWrite

```go
func (a *GlDeleteProgram) AddWrite(rng memory.Range, id binary.ID) *GlDeleteProgram
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteProgram pointer is returned so that calls can be chained.

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
func (ϟa *GlDeleteProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteProgram) Observations

```go
func (a *GlDeleteProgram) Observations() *atom.Observations
```

#### func (*GlDeleteProgram) Replay

```go
func (ϟa *GlDeleteProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteProgram) String

```go
func (a *GlDeleteProgram) String() string
```

#### type GlDeleteQueries

```go
type GlDeleteQueries struct {
	binary.Generate

	Count   int32
	Queries QueryIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteQueries
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteQueries

```go
func NewGlDeleteQueries(Count int32, Queries memory.Pointer) *GlDeleteQueries
```

#### func (*GlDeleteQueries) API

```go
func (c *GlDeleteQueries) API() gfxapi.ID
```

#### func (*GlDeleteQueries) AddRead

```go
func (a *GlDeleteQueries) AddRead(rng memory.Range, id binary.ID) *GlDeleteQueries
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteQueries pointer is returned so that calls can be chained.

#### func (*GlDeleteQueries) AddWrite

```go
func (a *GlDeleteQueries) AddWrite(rng memory.Range, id binary.ID) *GlDeleteQueries
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteQueries pointer is returned so that calls can be chained.

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
func (ϟa *GlDeleteQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteQueries) Observations

```go
func (a *GlDeleteQueries) Observations() *atom.Observations
```

#### func (*GlDeleteQueries) Replay

```go
func (ϟa *GlDeleteQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteQueries) String

```go
func (a *GlDeleteQueries) String() string
```

#### type GlDeleteQueriesEXT

```go
type GlDeleteQueriesEXT struct {
	binary.Generate

	Count   int32
	Queries QueryIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteQueriesEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteQueriesEXT

```go
func NewGlDeleteQueriesEXT(Count int32, Queries memory.Pointer) *GlDeleteQueriesEXT
```

#### func (*GlDeleteQueriesEXT) API

```go
func (c *GlDeleteQueriesEXT) API() gfxapi.ID
```

#### func (*GlDeleteQueriesEXT) AddRead

```go
func (a *GlDeleteQueriesEXT) AddRead(rng memory.Range, id binary.ID) *GlDeleteQueriesEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteQueriesEXT pointer is returned so that calls can be
chained.

#### func (*GlDeleteQueriesEXT) AddWrite

```go
func (a *GlDeleteQueriesEXT) AddWrite(rng memory.Range, id binary.ID) *GlDeleteQueriesEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteQueriesEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlDeleteQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteQueriesEXT) Observations

```go
func (a *GlDeleteQueriesEXT) Observations() *atom.Observations
```

#### func (*GlDeleteQueriesEXT) Replay

```go
func (ϟa *GlDeleteQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteQueriesEXT) String

```go
func (a *GlDeleteQueriesEXT) String() string
```

#### type GlDeleteRenderbuffers

```go
type GlDeleteRenderbuffers struct {
	binary.Generate

	Count         int32
	Renderbuffers RenderbufferIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteRenderbuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteRenderbuffers

```go
func NewGlDeleteRenderbuffers(Count int32, Renderbuffers memory.Pointer) *GlDeleteRenderbuffers
```

#### func (*GlDeleteRenderbuffers) API

```go
func (c *GlDeleteRenderbuffers) API() gfxapi.ID
```

#### func (*GlDeleteRenderbuffers) AddRead

```go
func (a *GlDeleteRenderbuffers) AddRead(rng memory.Range, id binary.ID) *GlDeleteRenderbuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteRenderbuffers pointer is returned so that calls can be
chained.

#### func (*GlDeleteRenderbuffers) AddWrite

```go
func (a *GlDeleteRenderbuffers) AddWrite(rng memory.Range, id binary.ID) *GlDeleteRenderbuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteRenderbuffers pointer is returned so that calls can be
chained.

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
func (ϟa *GlDeleteRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteRenderbuffers) Observations

```go
func (a *GlDeleteRenderbuffers) Observations() *atom.Observations
```

#### func (*GlDeleteRenderbuffers) Replay

```go
func (ϟa *GlDeleteRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteRenderbuffers) String

```go
func (a *GlDeleteRenderbuffers) String() string
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
func NewGlDeleteShader(Shader ShaderId) *GlDeleteShader
```

#### func (*GlDeleteShader) API

```go
func (c *GlDeleteShader) API() gfxapi.ID
```

#### func (*GlDeleteShader) AddRead

```go
func (a *GlDeleteShader) AddRead(rng memory.Range, id binary.ID) *GlDeleteShader
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteShader pointer is returned so that calls can be chained.

#### func (*GlDeleteShader) AddWrite

```go
func (a *GlDeleteShader) AddWrite(rng memory.Range, id binary.ID) *GlDeleteShader
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteShader pointer is returned so that calls can be chained.

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
func (ϟa *GlDeleteShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteShader) Observations

```go
func (a *GlDeleteShader) Observations() *atom.Observations
```

#### func (*GlDeleteShader) Replay

```go
func (ϟa *GlDeleteShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteShader) String

```go
func (a *GlDeleteShader) String() string
```

#### type GlDeleteSync

```go
type GlDeleteSync struct {
	binary.Generate

	Sync SyncObject
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteSync
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteSync

```go
func NewGlDeleteSync(Sync SyncObject) *GlDeleteSync
```

#### func (*GlDeleteSync) API

```go
func (c *GlDeleteSync) API() gfxapi.ID
```

#### func (*GlDeleteSync) AddRead

```go
func (a *GlDeleteSync) AddRead(rng memory.Range, id binary.ID) *GlDeleteSync
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteSync pointer is returned so that calls can be chained.

#### func (*GlDeleteSync) AddWrite

```go
func (a *GlDeleteSync) AddWrite(rng memory.Range, id binary.ID) *GlDeleteSync
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteSync pointer is returned so that calls can be chained.

#### func (*GlDeleteSync) Class

```go
func (*GlDeleteSync) Class() binary.Class
```

#### func (*GlDeleteSync) Flags

```go
func (c *GlDeleteSync) Flags() atom.Flags
```

#### func (*GlDeleteSync) Mutate

```go
func (ϟa *GlDeleteSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteSync) Observations

```go
func (a *GlDeleteSync) Observations() *atom.Observations
```

#### func (*GlDeleteSync) Replay

```go
func (ϟa *GlDeleteSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteSync) String

```go
func (a *GlDeleteSync) String() string
```

#### type GlDeleteTextures

```go
type GlDeleteTextures struct {
	binary.Generate

	Count    int32
	Textures TextureIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteTextures
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteTextures

```go
func NewGlDeleteTextures(Count int32, Textures memory.Pointer) *GlDeleteTextures
```

#### func (*GlDeleteTextures) API

```go
func (c *GlDeleteTextures) API() gfxapi.ID
```

#### func (*GlDeleteTextures) AddRead

```go
func (a *GlDeleteTextures) AddRead(rng memory.Range, id binary.ID) *GlDeleteTextures
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteTextures pointer is returned so that calls can be chained.

#### func (*GlDeleteTextures) AddWrite

```go
func (a *GlDeleteTextures) AddWrite(rng memory.Range, id binary.ID) *GlDeleteTextures
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteTextures pointer is returned so that calls can be chained.

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
func (ϟa *GlDeleteTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteTextures) Observations

```go
func (a *GlDeleteTextures) Observations() *atom.Observations
```

#### func (*GlDeleteTextures) Replay

```go
func (ϟa *GlDeleteTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteTextures) String

```go
func (a *GlDeleteTextures) String() string
```

#### type GlDeleteVertexArrays

```go
type GlDeleteVertexArrays struct {
	binary.Generate

	Count  uint32
	Arrays VertexArrayIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteVertexArrays
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteVertexArrays

```go
func NewGlDeleteVertexArrays(Count uint32, Arrays memory.Pointer) *GlDeleteVertexArrays
```

#### func (*GlDeleteVertexArrays) API

```go
func (c *GlDeleteVertexArrays) API() gfxapi.ID
```

#### func (*GlDeleteVertexArrays) AddRead

```go
func (a *GlDeleteVertexArrays) AddRead(rng memory.Range, id binary.ID) *GlDeleteVertexArrays
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteVertexArrays pointer is returned so that calls can be
chained.

#### func (*GlDeleteVertexArrays) AddWrite

```go
func (a *GlDeleteVertexArrays) AddWrite(rng memory.Range, id binary.ID) *GlDeleteVertexArrays
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteVertexArrays pointer is returned so that calls can be
chained.

#### func (*GlDeleteVertexArrays) Class

```go
func (*GlDeleteVertexArrays) Class() binary.Class
```

#### func (*GlDeleteVertexArrays) Flags

```go
func (c *GlDeleteVertexArrays) Flags() atom.Flags
```

#### func (*GlDeleteVertexArrays) Mutate

```go
func (ϟa *GlDeleteVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteVertexArrays) Observations

```go
func (a *GlDeleteVertexArrays) Observations() *atom.Observations
```

#### func (*GlDeleteVertexArrays) Replay

```go
func (ϟa *GlDeleteVertexArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteVertexArrays) String

```go
func (a *GlDeleteVertexArrays) String() string
```

#### type GlDeleteVertexArraysOES

```go
type GlDeleteVertexArraysOES struct {
	binary.Generate

	Count  int32
	Arrays VertexArrayIdᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDeleteVertexArraysOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDeleteVertexArraysOES

```go
func NewGlDeleteVertexArraysOES(Count int32, Arrays memory.Pointer) *GlDeleteVertexArraysOES
```

#### func (*GlDeleteVertexArraysOES) API

```go
func (c *GlDeleteVertexArraysOES) API() gfxapi.ID
```

#### func (*GlDeleteVertexArraysOES) AddRead

```go
func (a *GlDeleteVertexArraysOES) AddRead(rng memory.Range, id binary.ID) *GlDeleteVertexArraysOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDeleteVertexArraysOES pointer is returned so that calls can be
chained.

#### func (*GlDeleteVertexArraysOES) AddWrite

```go
func (a *GlDeleteVertexArraysOES) AddWrite(rng memory.Range, id binary.ID) *GlDeleteVertexArraysOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDeleteVertexArraysOES pointer is returned so that calls can be
chained.

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
func (ϟa *GlDeleteVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDeleteVertexArraysOES) Observations

```go
func (a *GlDeleteVertexArraysOES) Observations() *atom.Observations
```

#### func (*GlDeleteVertexArraysOES) Replay

```go
func (ϟa *GlDeleteVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDeleteVertexArraysOES) String

```go
func (a *GlDeleteVertexArraysOES) String() string
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
func NewGlDepthFunc(Function TestFunction) *GlDepthFunc
```

#### func (*GlDepthFunc) API

```go
func (c *GlDepthFunc) API() gfxapi.ID
```

#### func (*GlDepthFunc) AddRead

```go
func (a *GlDepthFunc) AddRead(rng memory.Range, id binary.ID) *GlDepthFunc
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDepthFunc pointer is returned so that calls can be chained.

#### func (*GlDepthFunc) AddWrite

```go
func (a *GlDepthFunc) AddWrite(rng memory.Range, id binary.ID) *GlDepthFunc
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDepthFunc pointer is returned so that calls can be chained.

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
func (ϟa *GlDepthFunc) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDepthFunc) Observations

```go
func (a *GlDepthFunc) Observations() *atom.Observations
```

#### func (*GlDepthFunc) Replay

```go
func (ϟa *GlDepthFunc) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDepthFunc) String

```go
func (a *GlDepthFunc) String() string
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
func NewGlDepthMask(Enabled bool) *GlDepthMask
```

#### func (*GlDepthMask) API

```go
func (c *GlDepthMask) API() gfxapi.ID
```

#### func (*GlDepthMask) AddRead

```go
func (a *GlDepthMask) AddRead(rng memory.Range, id binary.ID) *GlDepthMask
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDepthMask pointer is returned so that calls can be chained.

#### func (*GlDepthMask) AddWrite

```go
func (a *GlDepthMask) AddWrite(rng memory.Range, id binary.ID) *GlDepthMask
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDepthMask pointer is returned so that calls can be chained.

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
func (ϟa *GlDepthMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDepthMask) Observations

```go
func (a *GlDepthMask) Observations() *atom.Observations
```

#### func (*GlDepthMask) Replay

```go
func (ϟa *GlDepthMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDepthMask) String

```go
func (a *GlDepthMask) String() string
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
func NewGlDepthRangef(Near float32, Far float32) *GlDepthRangef
```

#### func (*GlDepthRangef) API

```go
func (c *GlDepthRangef) API() gfxapi.ID
```

#### func (*GlDepthRangef) AddRead

```go
func (a *GlDepthRangef) AddRead(rng memory.Range, id binary.ID) *GlDepthRangef
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDepthRangef pointer is returned so that calls can be chained.

#### func (*GlDepthRangef) AddWrite

```go
func (a *GlDepthRangef) AddWrite(rng memory.Range, id binary.ID) *GlDepthRangef
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDepthRangef pointer is returned so that calls can be chained.

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
func (ϟa *GlDepthRangef) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDepthRangef) Observations

```go
func (a *GlDepthRangef) Observations() *atom.Observations
```

#### func (*GlDepthRangef) Replay

```go
func (ϟa *GlDepthRangef) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDepthRangef) String

```go
func (a *GlDepthRangef) String() string
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
func NewGlDetachShader(Program ProgramId, Shader ShaderId) *GlDetachShader
```

#### func (*GlDetachShader) API

```go
func (c *GlDetachShader) API() gfxapi.ID
```

#### func (*GlDetachShader) AddRead

```go
func (a *GlDetachShader) AddRead(rng memory.Range, id binary.ID) *GlDetachShader
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDetachShader pointer is returned so that calls can be chained.

#### func (*GlDetachShader) AddWrite

```go
func (a *GlDetachShader) AddWrite(rng memory.Range, id binary.ID) *GlDetachShader
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDetachShader pointer is returned so that calls can be chained.

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
func (ϟa *GlDetachShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDetachShader) Observations

```go
func (a *GlDetachShader) Observations() *atom.Observations
```

#### func (*GlDetachShader) Replay

```go
func (ϟa *GlDetachShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDetachShader) String

```go
func (a *GlDetachShader) String() string
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
func NewGlDisable(Capability Capability) *GlDisable
```

#### func (*GlDisable) API

```go
func (c *GlDisable) API() gfxapi.ID
```

#### func (*GlDisable) AddRead

```go
func (a *GlDisable) AddRead(rng memory.Range, id binary.ID) *GlDisable
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDisable pointer is returned so that calls can be chained.

#### func (*GlDisable) AddWrite

```go
func (a *GlDisable) AddWrite(rng memory.Range, id binary.ID) *GlDisable
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDisable pointer is returned so that calls can be chained.

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
func (ϟa *GlDisable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDisable) Observations

```go
func (a *GlDisable) Observations() *atom.Observations
```

#### func (*GlDisable) Replay

```go
func (ϟa *GlDisable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDisable) String

```go
func (a *GlDisable) String() string
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
func NewGlDisableClientState(Type ArrayType) *GlDisableClientState
```

#### func (*GlDisableClientState) API

```go
func (c *GlDisableClientState) API() gfxapi.ID
```

#### func (*GlDisableClientState) AddRead

```go
func (a *GlDisableClientState) AddRead(rng memory.Range, id binary.ID) *GlDisableClientState
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDisableClientState pointer is returned so that calls can be
chained.

#### func (*GlDisableClientState) AddWrite

```go
func (a *GlDisableClientState) AddWrite(rng memory.Range, id binary.ID) *GlDisableClientState
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDisableClientState pointer is returned so that calls can be
chained.

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
func (ϟa *GlDisableClientState) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDisableClientState) Observations

```go
func (a *GlDisableClientState) Observations() *atom.Observations
```

#### func (*GlDisableClientState) Replay

```go
func (ϟa *GlDisableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDisableClientState) String

```go
func (a *GlDisableClientState) String() string
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
func NewGlDisableVertexAttribArray(Location AttributeLocation) *GlDisableVertexAttribArray
```

#### func (*GlDisableVertexAttribArray) API

```go
func (c *GlDisableVertexAttribArray) API() gfxapi.ID
```

#### func (*GlDisableVertexAttribArray) AddRead

```go
func (a *GlDisableVertexAttribArray) AddRead(rng memory.Range, id binary.ID) *GlDisableVertexAttribArray
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDisableVertexAttribArray pointer is returned so that calls can be
chained.

#### func (*GlDisableVertexAttribArray) AddWrite

```go
func (a *GlDisableVertexAttribArray) AddWrite(rng memory.Range, id binary.ID) *GlDisableVertexAttribArray
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDisableVertexAttribArray pointer is returned so that calls can be
chained.

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
func (ϟa *GlDisableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDisableVertexAttribArray) Observations

```go
func (a *GlDisableVertexAttribArray) Observations() *atom.Observations
```

#### func (*GlDisableVertexAttribArray) Replay

```go
func (ϟa *GlDisableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDisableVertexAttribArray) String

```go
func (a *GlDisableVertexAttribArray) String() string
```

#### type GlDiscardFramebufferEXT

```go
type GlDiscardFramebufferEXT struct {
	binary.Generate

	Target         FramebufferTarget
	NumAttachments int32
	Attachments    DiscardFramebufferAttachmentᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlDiscardFramebufferEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlDiscardFramebufferEXT

```go
func NewGlDiscardFramebufferEXT(Target FramebufferTarget, NumAttachments int32, Attachments memory.Pointer) *GlDiscardFramebufferEXT
```

#### func (*GlDiscardFramebufferEXT) API

```go
func (c *GlDiscardFramebufferEXT) API() gfxapi.ID
```

#### func (*GlDiscardFramebufferEXT) AddRead

```go
func (a *GlDiscardFramebufferEXT) AddRead(rng memory.Range, id binary.ID) *GlDiscardFramebufferEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDiscardFramebufferEXT pointer is returned so that calls can be
chained.

#### func (*GlDiscardFramebufferEXT) AddWrite

```go
func (a *GlDiscardFramebufferEXT) AddWrite(rng memory.Range, id binary.ID) *GlDiscardFramebufferEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDiscardFramebufferEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlDiscardFramebufferEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDiscardFramebufferEXT) Observations

```go
func (a *GlDiscardFramebufferEXT) Observations() *atom.Observations
```

#### func (*GlDiscardFramebufferEXT) Replay

```go
func (ϟa *GlDiscardFramebufferEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDiscardFramebufferEXT) String

```go
func (a *GlDiscardFramebufferEXT) String() string
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
func NewGlDrawArrays(Draw_mode DrawMode, First_index int32, Index_count int32) *GlDrawArrays
```

#### func (*GlDrawArrays) API

```go
func (c *GlDrawArrays) API() gfxapi.ID
```

#### func (*GlDrawArrays) AddRead

```go
func (a *GlDrawArrays) AddRead(rng memory.Range, id binary.ID) *GlDrawArrays
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDrawArrays pointer is returned so that calls can be chained.

#### func (*GlDrawArrays) AddWrite

```go
func (a *GlDrawArrays) AddWrite(rng memory.Range, id binary.ID) *GlDrawArrays
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDrawArrays pointer is returned so that calls can be chained.

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
func (ϟa *GlDrawArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDrawArrays) Observations

```go
func (a *GlDrawArrays) Observations() *atom.Observations
```

#### func (*GlDrawArrays) Replay

```go
func (ϟa *GlDrawArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDrawArrays) String

```go
func (a *GlDrawArrays) String() string
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
func NewGlDrawElements(Draw_mode DrawMode, Element_count int32, Indices_type IndicesType, Indices memory.Pointer) *GlDrawElements
```

#### func (*GlDrawElements) API

```go
func (c *GlDrawElements) API() gfxapi.ID
```

#### func (*GlDrawElements) AddRead

```go
func (a *GlDrawElements) AddRead(rng memory.Range, id binary.ID) *GlDrawElements
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlDrawElements pointer is returned so that calls can be chained.

#### func (*GlDrawElements) AddWrite

```go
func (a *GlDrawElements) AddWrite(rng memory.Range, id binary.ID) *GlDrawElements
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlDrawElements pointer is returned so that calls can be chained.

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
func (ϟa *GlDrawElements) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlDrawElements) Observations

```go
func (a *GlDrawElements) Observations() *atom.Observations
```

#### func (*GlDrawElements) Replay

```go
func (ϟa *GlDrawElements) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlDrawElements) String

```go
func (a *GlDrawElements) String() string
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
func NewGlEGLImageTargetRenderbufferStorageOES(Target ImageTargetRenderbufferStorage, Image memory.Pointer) *GlEGLImageTargetRenderbufferStorageOES
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) API

```go
func (c *GlEGLImageTargetRenderbufferStorageOES) API() gfxapi.ID
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) AddRead

```go
func (a *GlEGLImageTargetRenderbufferStorageOES) AddRead(rng memory.Range, id binary.ID) *GlEGLImageTargetRenderbufferStorageOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEGLImageTargetRenderbufferStorageOES pointer is returned so that
calls can be chained.

#### func (*GlEGLImageTargetRenderbufferStorageOES) AddWrite

```go
func (a *GlEGLImageTargetRenderbufferStorageOES) AddWrite(rng memory.Range, id binary.ID) *GlEGLImageTargetRenderbufferStorageOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEGLImageTargetRenderbufferStorageOES pointer is returned so that
calls can be chained.

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
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) Observations

```go
func (a *GlEGLImageTargetRenderbufferStorageOES) Observations() *atom.Observations
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) Replay

```go
func (ϟa *GlEGLImageTargetRenderbufferStorageOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEGLImageTargetRenderbufferStorageOES) String

```go
func (a *GlEGLImageTargetRenderbufferStorageOES) String() string
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
func NewGlEGLImageTargetTexture2DOES(Target ImageTargetTexture, Image memory.Pointer) *GlEGLImageTargetTexture2DOES
```

#### func (*GlEGLImageTargetTexture2DOES) API

```go
func (c *GlEGLImageTargetTexture2DOES) API() gfxapi.ID
```

#### func (*GlEGLImageTargetTexture2DOES) AddRead

```go
func (a *GlEGLImageTargetTexture2DOES) AddRead(rng memory.Range, id binary.ID) *GlEGLImageTargetTexture2DOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEGLImageTargetTexture2DOES pointer is returned so that calls can
be chained.

#### func (*GlEGLImageTargetTexture2DOES) AddWrite

```go
func (a *GlEGLImageTargetTexture2DOES) AddWrite(rng memory.Range, id binary.ID) *GlEGLImageTargetTexture2DOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEGLImageTargetTexture2DOES pointer is returned so that calls can
be chained.

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
func (ϟa *GlEGLImageTargetTexture2DOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEGLImageTargetTexture2DOES) Observations

```go
func (a *GlEGLImageTargetTexture2DOES) Observations() *atom.Observations
```

#### func (*GlEGLImageTargetTexture2DOES) Replay

```go
func (ϟa *GlEGLImageTargetTexture2DOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEGLImageTargetTexture2DOES) String

```go
func (a *GlEGLImageTargetTexture2DOES) String() string
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
func NewGlEnable(Capability Capability) *GlEnable
```

#### func (*GlEnable) API

```go
func (c *GlEnable) API() gfxapi.ID
```

#### func (*GlEnable) AddRead

```go
func (a *GlEnable) AddRead(rng memory.Range, id binary.ID) *GlEnable
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEnable pointer is returned so that calls can be chained.

#### func (*GlEnable) AddWrite

```go
func (a *GlEnable) AddWrite(rng memory.Range, id binary.ID) *GlEnable
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEnable pointer is returned so that calls can be chained.

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
func (ϟa *GlEnable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEnable) Observations

```go
func (a *GlEnable) Observations() *atom.Observations
```

#### func (*GlEnable) Replay

```go
func (ϟa *GlEnable) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEnable) String

```go
func (a *GlEnable) String() string
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
func NewGlEnableClientState(Type ArrayType) *GlEnableClientState
```

#### func (*GlEnableClientState) API

```go
func (c *GlEnableClientState) API() gfxapi.ID
```

#### func (*GlEnableClientState) AddRead

```go
func (a *GlEnableClientState) AddRead(rng memory.Range, id binary.ID) *GlEnableClientState
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEnableClientState pointer is returned so that calls can be
chained.

#### func (*GlEnableClientState) AddWrite

```go
func (a *GlEnableClientState) AddWrite(rng memory.Range, id binary.ID) *GlEnableClientState
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEnableClientState pointer is returned so that calls can be
chained.

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
func (ϟa *GlEnableClientState) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEnableClientState) Observations

```go
func (a *GlEnableClientState) Observations() *atom.Observations
```

#### func (*GlEnableClientState) Replay

```go
func (ϟa *GlEnableClientState) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEnableClientState) String

```go
func (a *GlEnableClientState) String() string
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
func NewGlEnableVertexAttribArray(Location AttributeLocation) *GlEnableVertexAttribArray
```

#### func (*GlEnableVertexAttribArray) API

```go
func (c *GlEnableVertexAttribArray) API() gfxapi.ID
```

#### func (*GlEnableVertexAttribArray) AddRead

```go
func (a *GlEnableVertexAttribArray) AddRead(rng memory.Range, id binary.ID) *GlEnableVertexAttribArray
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEnableVertexAttribArray pointer is returned so that calls can be
chained.

#### func (*GlEnableVertexAttribArray) AddWrite

```go
func (a *GlEnableVertexAttribArray) AddWrite(rng memory.Range, id binary.ID) *GlEnableVertexAttribArray
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEnableVertexAttribArray pointer is returned so that calls can be
chained.

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
func (ϟa *GlEnableVertexAttribArray) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEnableVertexAttribArray) Observations

```go
func (a *GlEnableVertexAttribArray) Observations() *atom.Observations
```

#### func (*GlEnableVertexAttribArray) Replay

```go
func (ϟa *GlEnableVertexAttribArray) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEnableVertexAttribArray) String

```go
func (a *GlEnableVertexAttribArray) String() string
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
func NewGlEndQuery(Target QueryTarget) *GlEndQuery
```

#### func (*GlEndQuery) API

```go
func (c *GlEndQuery) API() gfxapi.ID
```

#### func (*GlEndQuery) AddRead

```go
func (a *GlEndQuery) AddRead(rng memory.Range, id binary.ID) *GlEndQuery
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEndQuery pointer is returned so that calls can be chained.

#### func (*GlEndQuery) AddWrite

```go
func (a *GlEndQuery) AddWrite(rng memory.Range, id binary.ID) *GlEndQuery
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEndQuery pointer is returned so that calls can be chained.

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
func (ϟa *GlEndQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEndQuery) Observations

```go
func (a *GlEndQuery) Observations() *atom.Observations
```

#### func (*GlEndQuery) Replay

```go
func (ϟa *GlEndQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEndQuery) String

```go
func (a *GlEndQuery) String() string
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
func NewGlEndQueryEXT(Target QueryTarget) *GlEndQueryEXT
```

#### func (*GlEndQueryEXT) API

```go
func (c *GlEndQueryEXT) API() gfxapi.ID
```

#### func (*GlEndQueryEXT) AddRead

```go
func (a *GlEndQueryEXT) AddRead(rng memory.Range, id binary.ID) *GlEndQueryEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEndQueryEXT pointer is returned so that calls can be chained.

#### func (*GlEndQueryEXT) AddWrite

```go
func (a *GlEndQueryEXT) AddWrite(rng memory.Range, id binary.ID) *GlEndQueryEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEndQueryEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlEndQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEndQueryEXT) Observations

```go
func (a *GlEndQueryEXT) Observations() *atom.Observations
```

#### func (*GlEndQueryEXT) Replay

```go
func (ϟa *GlEndQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEndQueryEXT) String

```go
func (a *GlEndQueryEXT) String() string
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
func NewGlEndTilingQCOM(Preserve_mask TilePreserveMaskQCOM) *GlEndTilingQCOM
```

#### func (*GlEndTilingQCOM) API

```go
func (c *GlEndTilingQCOM) API() gfxapi.ID
```

#### func (*GlEndTilingQCOM) AddRead

```go
func (a *GlEndTilingQCOM) AddRead(rng memory.Range, id binary.ID) *GlEndTilingQCOM
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlEndTilingQCOM pointer is returned so that calls can be chained.

#### func (*GlEndTilingQCOM) AddWrite

```go
func (a *GlEndTilingQCOM) AddWrite(rng memory.Range, id binary.ID) *GlEndTilingQCOM
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlEndTilingQCOM pointer is returned so that calls can be chained.

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
func (ϟa *GlEndTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlEndTilingQCOM) Observations

```go
func (a *GlEndTilingQCOM) Observations() *atom.Observations
```

#### func (*GlEndTilingQCOM) Replay

```go
func (ϟa *GlEndTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlEndTilingQCOM) String

```go
func (a *GlEndTilingQCOM) String() string
```

#### type GlFenceSync

```go
type GlFenceSync struct {
	binary.Generate

	Condition SyncCondition
	SyncFlags SyncFlags
	Result    SyncObject
}
```

//////////////////////////////////////////////////////////////////////////////
GlFenceSync
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlFenceSync

```go
func NewGlFenceSync(Condition SyncCondition, SyncFlags SyncFlags, Result SyncObject) *GlFenceSync
```

#### func (*GlFenceSync) API

```go
func (c *GlFenceSync) API() gfxapi.ID
```

#### func (*GlFenceSync) AddRead

```go
func (a *GlFenceSync) AddRead(rng memory.Range, id binary.ID) *GlFenceSync
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlFenceSync pointer is returned so that calls can be chained.

#### func (*GlFenceSync) AddWrite

```go
func (a *GlFenceSync) AddWrite(rng memory.Range, id binary.ID) *GlFenceSync
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlFenceSync pointer is returned so that calls can be chained.

#### func (*GlFenceSync) Class

```go
func (*GlFenceSync) Class() binary.Class
```

#### func (*GlFenceSync) Flags

```go
func (c *GlFenceSync) Flags() atom.Flags
```

#### func (*GlFenceSync) Mutate

```go
func (ϟa *GlFenceSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlFenceSync) Observations

```go
func (a *GlFenceSync) Observations() *atom.Observations
```

#### func (*GlFenceSync) Replay

```go
func (ϟa *GlFenceSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlFenceSync) String

```go
func (a *GlFenceSync) String() string
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
func (c *GlFinish) API() gfxapi.ID
```

#### func (*GlFinish) AddRead

```go
func (a *GlFinish) AddRead(rng memory.Range, id binary.ID) *GlFinish
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlFinish pointer is returned so that calls can be chained.

#### func (*GlFinish) AddWrite

```go
func (a *GlFinish) AddWrite(rng memory.Range, id binary.ID) *GlFinish
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlFinish pointer is returned so that calls can be chained.

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
func (ϟa *GlFinish) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlFinish) Observations

```go
func (a *GlFinish) Observations() *atom.Observations
```

#### func (*GlFinish) Replay

```go
func (ϟa *GlFinish) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlFinish) String

```go
func (a *GlFinish) String() string
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
func (c *GlFlush) API() gfxapi.ID
```

#### func (*GlFlush) AddRead

```go
func (a *GlFlush) AddRead(rng memory.Range, id binary.ID) *GlFlush
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlFlush pointer is returned so that calls can be chained.

#### func (*GlFlush) AddWrite

```go
func (a *GlFlush) AddWrite(rng memory.Range, id binary.ID) *GlFlush
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlFlush pointer is returned so that calls can be chained.

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
func (ϟa *GlFlush) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlFlush) Observations

```go
func (a *GlFlush) Observations() *atom.Observations
```

#### func (*GlFlush) Replay

```go
func (ϟa *GlFlush) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlFlush) String

```go
func (a *GlFlush) String() string
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
func NewGlFramebufferRenderbuffer(Framebuffer_target FramebufferTarget, Framebuffer_attachment FramebufferAttachment, Renderbuffer_target RenderbufferTarget, Renderbuffer RenderbufferId) *GlFramebufferRenderbuffer
```

#### func (*GlFramebufferRenderbuffer) API

```go
func (c *GlFramebufferRenderbuffer) API() gfxapi.ID
```

#### func (*GlFramebufferRenderbuffer) AddRead

```go
func (a *GlFramebufferRenderbuffer) AddRead(rng memory.Range, id binary.ID) *GlFramebufferRenderbuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlFramebufferRenderbuffer pointer is returned so that calls can be
chained.

#### func (*GlFramebufferRenderbuffer) AddWrite

```go
func (a *GlFramebufferRenderbuffer) AddWrite(rng memory.Range, id binary.ID) *GlFramebufferRenderbuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlFramebufferRenderbuffer pointer is returned so that calls can be
chained.

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
func (ϟa *GlFramebufferRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlFramebufferRenderbuffer) Observations

```go
func (a *GlFramebufferRenderbuffer) Observations() *atom.Observations
```

#### func (*GlFramebufferRenderbuffer) Replay

```go
func (ϟa *GlFramebufferRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlFramebufferRenderbuffer) String

```go
func (a *GlFramebufferRenderbuffer) String() string
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
func NewGlFramebufferTexture2D(Framebuffer_target FramebufferTarget, Framebuffer_attachment FramebufferAttachment, Texture_target TextureImageTarget, Texture TextureId, Level int32) *GlFramebufferTexture2D
```

#### func (*GlFramebufferTexture2D) API

```go
func (c *GlFramebufferTexture2D) API() gfxapi.ID
```

#### func (*GlFramebufferTexture2D) AddRead

```go
func (a *GlFramebufferTexture2D) AddRead(rng memory.Range, id binary.ID) *GlFramebufferTexture2D
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlFramebufferTexture2D pointer is returned so that calls can be
chained.

#### func (*GlFramebufferTexture2D) AddWrite

```go
func (a *GlFramebufferTexture2D) AddWrite(rng memory.Range, id binary.ID) *GlFramebufferTexture2D
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlFramebufferTexture2D pointer is returned so that calls can be
chained.

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
func (ϟa *GlFramebufferTexture2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlFramebufferTexture2D) Observations

```go
func (a *GlFramebufferTexture2D) Observations() *atom.Observations
```

#### func (*GlFramebufferTexture2D) Replay

```go
func (ϟa *GlFramebufferTexture2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlFramebufferTexture2D) String

```go
func (a *GlFramebufferTexture2D) String() string
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
func NewGlFrontFace(Orientation FaceOrientation) *GlFrontFace
```

#### func (*GlFrontFace) API

```go
func (c *GlFrontFace) API() gfxapi.ID
```

#### func (*GlFrontFace) AddRead

```go
func (a *GlFrontFace) AddRead(rng memory.Range, id binary.ID) *GlFrontFace
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlFrontFace pointer is returned so that calls can be chained.

#### func (*GlFrontFace) AddWrite

```go
func (a *GlFrontFace) AddWrite(rng memory.Range, id binary.ID) *GlFrontFace
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlFrontFace pointer is returned so that calls can be chained.

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
func (ϟa *GlFrontFace) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlFrontFace) Observations

```go
func (a *GlFrontFace) Observations() *atom.Observations
```

#### func (*GlFrontFace) Replay

```go
func (ϟa *GlFrontFace) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlFrontFace) String

```go
func (a *GlFrontFace) String() string
```

#### type GlGenBuffers

```go
type GlGenBuffers struct {
	binary.Generate

	Count   int32
	Buffers BufferIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenBuffers

```go
func NewGlGenBuffers(Count int32, Buffers memory.Pointer) *GlGenBuffers
```

#### func (*GlGenBuffers) API

```go
func (c *GlGenBuffers) API() gfxapi.ID
```

#### func (*GlGenBuffers) AddRead

```go
func (a *GlGenBuffers) AddRead(rng memory.Range, id binary.ID) *GlGenBuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenBuffers pointer is returned so that calls can be chained.

#### func (*GlGenBuffers) AddWrite

```go
func (a *GlGenBuffers) AddWrite(rng memory.Range, id binary.ID) *GlGenBuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenBuffers pointer is returned so that calls can be chained.

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
func (ϟa *GlGenBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenBuffers) Observations

```go
func (a *GlGenBuffers) Observations() *atom.Observations
```

#### func (*GlGenBuffers) Replay

```go
func (ϟa *GlGenBuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenBuffers) String

```go
func (a *GlGenBuffers) String() string
```

#### type GlGenFramebuffers

```go
type GlGenFramebuffers struct {
	binary.Generate

	Count        int32
	Framebuffers FramebufferIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenFramebuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenFramebuffers

```go
func NewGlGenFramebuffers(Count int32, Framebuffers memory.Pointer) *GlGenFramebuffers
```

#### func (*GlGenFramebuffers) API

```go
func (c *GlGenFramebuffers) API() gfxapi.ID
```

#### func (*GlGenFramebuffers) AddRead

```go
func (a *GlGenFramebuffers) AddRead(rng memory.Range, id binary.ID) *GlGenFramebuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenFramebuffers pointer is returned so that calls can be chained.

#### func (*GlGenFramebuffers) AddWrite

```go
func (a *GlGenFramebuffers) AddWrite(rng memory.Range, id binary.ID) *GlGenFramebuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenFramebuffers pointer is returned so that calls can be chained.

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
func (ϟa *GlGenFramebuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenFramebuffers) Observations

```go
func (a *GlGenFramebuffers) Observations() *atom.Observations
```

#### func (*GlGenFramebuffers) Replay

```go
func (ϟa *GlGenFramebuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenFramebuffers) String

```go
func (a *GlGenFramebuffers) String() string
```

#### type GlGenQueries

```go
type GlGenQueries struct {
	binary.Generate

	Count   int32
	Queries QueryIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenQueries
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenQueries

```go
func NewGlGenQueries(Count int32, Queries memory.Pointer) *GlGenQueries
```

#### func (*GlGenQueries) API

```go
func (c *GlGenQueries) API() gfxapi.ID
```

#### func (*GlGenQueries) AddRead

```go
func (a *GlGenQueries) AddRead(rng memory.Range, id binary.ID) *GlGenQueries
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenQueries pointer is returned so that calls can be chained.

#### func (*GlGenQueries) AddWrite

```go
func (a *GlGenQueries) AddWrite(rng memory.Range, id binary.ID) *GlGenQueries
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenQueries pointer is returned so that calls can be chained.

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
func (ϟa *GlGenQueries) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenQueries) Observations

```go
func (a *GlGenQueries) Observations() *atom.Observations
```

#### func (*GlGenQueries) Replay

```go
func (ϟa *GlGenQueries) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenQueries) String

```go
func (a *GlGenQueries) String() string
```

#### type GlGenQueriesEXT

```go
type GlGenQueriesEXT struct {
	binary.Generate

	Count   int32
	Queries QueryIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenQueriesEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenQueriesEXT

```go
func NewGlGenQueriesEXT(Count int32, Queries memory.Pointer) *GlGenQueriesEXT
```

#### func (*GlGenQueriesEXT) API

```go
func (c *GlGenQueriesEXT) API() gfxapi.ID
```

#### func (*GlGenQueriesEXT) AddRead

```go
func (a *GlGenQueriesEXT) AddRead(rng memory.Range, id binary.ID) *GlGenQueriesEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenQueriesEXT pointer is returned so that calls can be chained.

#### func (*GlGenQueriesEXT) AddWrite

```go
func (a *GlGenQueriesEXT) AddWrite(rng memory.Range, id binary.ID) *GlGenQueriesEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenQueriesEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlGenQueriesEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenQueriesEXT) Observations

```go
func (a *GlGenQueriesEXT) Observations() *atom.Observations
```

#### func (*GlGenQueriesEXT) Replay

```go
func (ϟa *GlGenQueriesEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenQueriesEXT) String

```go
func (a *GlGenQueriesEXT) String() string
```

#### type GlGenRenderbuffers

```go
type GlGenRenderbuffers struct {
	binary.Generate

	Count         int32
	Renderbuffers RenderbufferIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenRenderbuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenRenderbuffers

```go
func NewGlGenRenderbuffers(Count int32, Renderbuffers memory.Pointer) *GlGenRenderbuffers
```

#### func (*GlGenRenderbuffers) API

```go
func (c *GlGenRenderbuffers) API() gfxapi.ID
```

#### func (*GlGenRenderbuffers) AddRead

```go
func (a *GlGenRenderbuffers) AddRead(rng memory.Range, id binary.ID) *GlGenRenderbuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenRenderbuffers pointer is returned so that calls can be
chained.

#### func (*GlGenRenderbuffers) AddWrite

```go
func (a *GlGenRenderbuffers) AddWrite(rng memory.Range, id binary.ID) *GlGenRenderbuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenRenderbuffers pointer is returned so that calls can be
chained.

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
func (ϟa *GlGenRenderbuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenRenderbuffers) Observations

```go
func (a *GlGenRenderbuffers) Observations() *atom.Observations
```

#### func (*GlGenRenderbuffers) Replay

```go
func (ϟa *GlGenRenderbuffers) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenRenderbuffers) String

```go
func (a *GlGenRenderbuffers) String() string
```

#### type GlGenTextures

```go
type GlGenTextures struct {
	binary.Generate

	Count    int32
	Textures TextureIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenTextures
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenTextures

```go
func NewGlGenTextures(Count int32, Textures memory.Pointer) *GlGenTextures
```

#### func (*GlGenTextures) API

```go
func (c *GlGenTextures) API() gfxapi.ID
```

#### func (*GlGenTextures) AddRead

```go
func (a *GlGenTextures) AddRead(rng memory.Range, id binary.ID) *GlGenTextures
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenTextures pointer is returned so that calls can be chained.

#### func (*GlGenTextures) AddWrite

```go
func (a *GlGenTextures) AddWrite(rng memory.Range, id binary.ID) *GlGenTextures
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenTextures pointer is returned so that calls can be chained.

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
func (ϟa *GlGenTextures) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenTextures) Observations

```go
func (a *GlGenTextures) Observations() *atom.Observations
```

#### func (*GlGenTextures) Replay

```go
func (ϟa *GlGenTextures) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenTextures) String

```go
func (a *GlGenTextures) String() string
```

#### type GlGenVertexArrays

```go
type GlGenVertexArrays struct {
	binary.Generate

	Count  int32
	Arrays VertexArrayIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenVertexArrays
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenVertexArrays

```go
func NewGlGenVertexArrays(Count int32, Arrays memory.Pointer) *GlGenVertexArrays
```

#### func (*GlGenVertexArrays) API

```go
func (c *GlGenVertexArrays) API() gfxapi.ID
```

#### func (*GlGenVertexArrays) AddRead

```go
func (a *GlGenVertexArrays) AddRead(rng memory.Range, id binary.ID) *GlGenVertexArrays
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenVertexArrays pointer is returned so that calls can be chained.

#### func (*GlGenVertexArrays) AddWrite

```go
func (a *GlGenVertexArrays) AddWrite(rng memory.Range, id binary.ID) *GlGenVertexArrays
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenVertexArrays pointer is returned so that calls can be chained.

#### func (*GlGenVertexArrays) Class

```go
func (*GlGenVertexArrays) Class() binary.Class
```

#### func (*GlGenVertexArrays) Flags

```go
func (c *GlGenVertexArrays) Flags() atom.Flags
```

#### func (*GlGenVertexArrays) Mutate

```go
func (ϟa *GlGenVertexArrays) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenVertexArrays) Observations

```go
func (a *GlGenVertexArrays) Observations() *atom.Observations
```

#### func (*GlGenVertexArrays) Replay

```go
func (ϟa *GlGenVertexArrays) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenVertexArrays) String

```go
func (a *GlGenVertexArrays) String() string
```

#### type GlGenVertexArraysOES

```go
type GlGenVertexArraysOES struct {
	binary.Generate

	Count  int32
	Arrays VertexArrayIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGenVertexArraysOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGenVertexArraysOES

```go
func NewGlGenVertexArraysOES(Count int32, Arrays memory.Pointer) *GlGenVertexArraysOES
```

#### func (*GlGenVertexArraysOES) API

```go
func (c *GlGenVertexArraysOES) API() gfxapi.ID
```

#### func (*GlGenVertexArraysOES) AddRead

```go
func (a *GlGenVertexArraysOES) AddRead(rng memory.Range, id binary.ID) *GlGenVertexArraysOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenVertexArraysOES pointer is returned so that calls can be
chained.

#### func (*GlGenVertexArraysOES) AddWrite

```go
func (a *GlGenVertexArraysOES) AddWrite(rng memory.Range, id binary.ID) *GlGenVertexArraysOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenVertexArraysOES pointer is returned so that calls can be
chained.

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
func (ϟa *GlGenVertexArraysOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenVertexArraysOES) Observations

```go
func (a *GlGenVertexArraysOES) Observations() *atom.Observations
```

#### func (*GlGenVertexArraysOES) Replay

```go
func (ϟa *GlGenVertexArraysOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenVertexArraysOES) String

```go
func (a *GlGenVertexArraysOES) String() string
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
func NewGlGenerateMipmap(Target TextureImageTarget) *GlGenerateMipmap
```

#### func (*GlGenerateMipmap) API

```go
func (c *GlGenerateMipmap) API() gfxapi.ID
```

#### func (*GlGenerateMipmap) AddRead

```go
func (a *GlGenerateMipmap) AddRead(rng memory.Range, id binary.ID) *GlGenerateMipmap
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGenerateMipmap pointer is returned so that calls can be chained.

#### func (*GlGenerateMipmap) AddWrite

```go
func (a *GlGenerateMipmap) AddWrite(rng memory.Range, id binary.ID) *GlGenerateMipmap
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGenerateMipmap pointer is returned so that calls can be chained.

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
func (ϟa *GlGenerateMipmap) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGenerateMipmap) Observations

```go
func (a *GlGenerateMipmap) Observations() *atom.Observations
```

#### func (*GlGenerateMipmap) Replay

```go
func (ϟa *GlGenerateMipmap) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGenerateMipmap) String

```go
func (a *GlGenerateMipmap) String() string
```

#### type GlGetActiveAttrib

```go
type GlGetActiveAttrib struct {
	binary.Generate

	Program            ProgramId
	Location           AttributeLocation
	BufferSize         int32
	BufferBytesWritten S32ᵖ
	VectorCount        S32ᵖ
	Type               ShaderAttribTypeᵖ
	Name               Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetActiveAttrib
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetActiveAttrib

```go
func NewGlGetActiveAttrib(Program ProgramId, Location AttributeLocation, Buffer_size int32, Buffer_bytes_written memory.Pointer, Vector_count memory.Pointer, Type memory.Pointer, Name memory.Pointer) *GlGetActiveAttrib
```

#### func (*GlGetActiveAttrib) API

```go
func (c *GlGetActiveAttrib) API() gfxapi.ID
```

#### func (*GlGetActiveAttrib) AddRead

```go
func (a *GlGetActiveAttrib) AddRead(rng memory.Range, id binary.ID) *GlGetActiveAttrib
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetActiveAttrib pointer is returned so that calls can be chained.

#### func (*GlGetActiveAttrib) AddWrite

```go
func (a *GlGetActiveAttrib) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveAttrib
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetActiveAttrib pointer is returned so that calls can be chained.

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
func (ϟa *GlGetActiveAttrib) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetActiveAttrib) Observations

```go
func (a *GlGetActiveAttrib) Observations() *atom.Observations
```

#### func (*GlGetActiveAttrib) Replay

```go
func (ϟa *GlGetActiveAttrib) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetActiveAttrib) String

```go
func (a *GlGetActiveAttrib) String() string
```

#### type GlGetActiveUniform

```go
type GlGetActiveUniform struct {
	binary.Generate

	Program            ProgramId
	Location           int32
	BufferSize         int32
	BufferBytesWritten S32ᵖ
	VectorCount        S32ᵖ
	Type               ShaderUniformTypeᵖ
	Name               Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetActiveUniform
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetActiveUniform

```go
func NewGlGetActiveUniform(Program ProgramId, Location int32, Buffer_size int32, Buffer_bytes_written memory.Pointer, Vector_count memory.Pointer, Type memory.Pointer, Name memory.Pointer) *GlGetActiveUniform
```

#### func (*GlGetActiveUniform) API

```go
func (c *GlGetActiveUniform) API() gfxapi.ID
```

#### func (*GlGetActiveUniform) AddRead

```go
func (a *GlGetActiveUniform) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniform
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetActiveUniform pointer is returned so that calls can be
chained.

#### func (*GlGetActiveUniform) AddWrite

```go
func (a *GlGetActiveUniform) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniform
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetActiveUniform pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetActiveUniform) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetActiveUniform) Observations

```go
func (a *GlGetActiveUniform) Observations() *atom.Observations
```

#### func (*GlGetActiveUniform) Replay

```go
func (ϟa *GlGetActiveUniform) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetActiveUniform) String

```go
func (a *GlGetActiveUniform) String() string
```

#### type GlGetActiveUniformBlockName

```go
type GlGetActiveUniformBlockName struct {
	binary.Generate

	Program            ProgramId
	UniformBlockIndex  uint32
	BufferSize         int32
	BufferBytesWritten S32ᵖ
	Name               Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetActiveUniformBlockName
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetActiveUniformBlockName

```go
func NewGlGetActiveUniformBlockName(Program ProgramId, Uniform_block_index uint32, Buffer_size int32, Buffer_bytes_written memory.Pointer, Name memory.Pointer) *GlGetActiveUniformBlockName
```

#### func (*GlGetActiveUniformBlockName) API

```go
func (c *GlGetActiveUniformBlockName) API() gfxapi.ID
```

#### func (*GlGetActiveUniformBlockName) AddRead

```go
func (a *GlGetActiveUniformBlockName) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockName
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetActiveUniformBlockName pointer is returned so that calls can
be chained.

#### func (*GlGetActiveUniformBlockName) AddWrite

```go
func (a *GlGetActiveUniformBlockName) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockName
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetActiveUniformBlockName pointer is returned so that calls can
be chained.

#### func (*GlGetActiveUniformBlockName) Class

```go
func (*GlGetActiveUniformBlockName) Class() binary.Class
```

#### func (*GlGetActiveUniformBlockName) Flags

```go
func (c *GlGetActiveUniformBlockName) Flags() atom.Flags
```

#### func (*GlGetActiveUniformBlockName) Mutate

```go
func (ϟa *GlGetActiveUniformBlockName) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetActiveUniformBlockName) Observations

```go
func (a *GlGetActiveUniformBlockName) Observations() *atom.Observations
```

#### func (*GlGetActiveUniformBlockName) Replay

```go
func (ϟa *GlGetActiveUniformBlockName) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetActiveUniformBlockName) String

```go
func (a *GlGetActiveUniformBlockName) String() string
```

#### type GlGetActiveUniformBlockiv

```go
type GlGetActiveUniformBlockiv struct {
	binary.Generate

	Program           ProgramId
	UniformBlockIndex uint32
	ParameterName     UniformBlockParameter
	Parameters        S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetActiveUniformBlockiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetActiveUniformBlockiv

```go
func NewGlGetActiveUniformBlockiv(Program ProgramId, Uniform_block_index uint32, Parameter_name UniformBlockParameter, Parameters memory.Pointer) *GlGetActiveUniformBlockiv
```

#### func (*GlGetActiveUniformBlockiv) API

```go
func (c *GlGetActiveUniformBlockiv) API() gfxapi.ID
```

#### func (*GlGetActiveUniformBlockiv) AddRead

```go
func (a *GlGetActiveUniformBlockiv) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockiv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetActiveUniformBlockiv pointer is returned so that calls can be
chained.

#### func (*GlGetActiveUniformBlockiv) AddWrite

```go
func (a *GlGetActiveUniformBlockiv) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniformBlockiv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetActiveUniformBlockiv pointer is returned so that calls can be
chained.

#### func (*GlGetActiveUniformBlockiv) Class

```go
func (*GlGetActiveUniformBlockiv) Class() binary.Class
```

#### func (*GlGetActiveUniformBlockiv) Flags

```go
func (c *GlGetActiveUniformBlockiv) Flags() atom.Flags
```

#### func (*GlGetActiveUniformBlockiv) Mutate

```go
func (ϟa *GlGetActiveUniformBlockiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetActiveUniformBlockiv) Observations

```go
func (a *GlGetActiveUniformBlockiv) Observations() *atom.Observations
```

#### func (*GlGetActiveUniformBlockiv) Replay

```go
func (ϟa *GlGetActiveUniformBlockiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetActiveUniformBlockiv) String

```go
func (a *GlGetActiveUniformBlockiv) String() string
```

#### type GlGetActiveUniformsiv

```go
type GlGetActiveUniformsiv struct {
	binary.Generate

	Program        ProgramId
	UniformCount   uint32
	UniformIndices U32ᵖ
	ParameterName  UniformBlockParameter
	Parameters     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetActiveUniformsiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetActiveUniformsiv

```go
func NewGlGetActiveUniformsiv(Program ProgramId, Uniform_count uint32, Uniform_indices memory.Pointer, Parameter_name UniformBlockParameter, Parameters memory.Pointer) *GlGetActiveUniformsiv
```

#### func (*GlGetActiveUniformsiv) API

```go
func (c *GlGetActiveUniformsiv) API() gfxapi.ID
```

#### func (*GlGetActiveUniformsiv) AddRead

```go
func (a *GlGetActiveUniformsiv) AddRead(rng memory.Range, id binary.ID) *GlGetActiveUniformsiv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetActiveUniformsiv pointer is returned so that calls can be
chained.

#### func (*GlGetActiveUniformsiv) AddWrite

```go
func (a *GlGetActiveUniformsiv) AddWrite(rng memory.Range, id binary.ID) *GlGetActiveUniformsiv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetActiveUniformsiv pointer is returned so that calls can be
chained.

#### func (*GlGetActiveUniformsiv) Class

```go
func (*GlGetActiveUniformsiv) Class() binary.Class
```

#### func (*GlGetActiveUniformsiv) Flags

```go
func (c *GlGetActiveUniformsiv) Flags() atom.Flags
```

#### func (*GlGetActiveUniformsiv) Mutate

```go
func (ϟa *GlGetActiveUniformsiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetActiveUniformsiv) Observations

```go
func (a *GlGetActiveUniformsiv) Observations() *atom.Observations
```

#### func (*GlGetActiveUniformsiv) Replay

```go
func (ϟa *GlGetActiveUniformsiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetActiveUniformsiv) String

```go
func (a *GlGetActiveUniformsiv) String() string
```

#### type GlGetAttachedShaders

```go
type GlGetAttachedShaders struct {
	binary.Generate

	Program              ProgramId
	BufferLength         int32
	ShadersLengthWritten S32ᵖ
	Shaders              ShaderIdᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetAttachedShaders
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetAttachedShaders

```go
func NewGlGetAttachedShaders(Program ProgramId, Buffer_length int32, Shaders_length_written memory.Pointer, Shaders memory.Pointer) *GlGetAttachedShaders
```

#### func (*GlGetAttachedShaders) API

```go
func (c *GlGetAttachedShaders) API() gfxapi.ID
```

#### func (*GlGetAttachedShaders) AddRead

```go
func (a *GlGetAttachedShaders) AddRead(rng memory.Range, id binary.ID) *GlGetAttachedShaders
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetAttachedShaders pointer is returned so that calls can be
chained.

#### func (*GlGetAttachedShaders) AddWrite

```go
func (a *GlGetAttachedShaders) AddWrite(rng memory.Range, id binary.ID) *GlGetAttachedShaders
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetAttachedShaders pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetAttachedShaders) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetAttachedShaders) Observations

```go
func (a *GlGetAttachedShaders) Observations() *atom.Observations
```

#### func (*GlGetAttachedShaders) Replay

```go
func (ϟa *GlGetAttachedShaders) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetAttachedShaders) String

```go
func (a *GlGetAttachedShaders) String() string
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
func NewGlGetAttribLocation(Program ProgramId, Name string, Result AttributeLocation) *GlGetAttribLocation
```

#### func (*GlGetAttribLocation) API

```go
func (c *GlGetAttribLocation) API() gfxapi.ID
```

#### func (*GlGetAttribLocation) AddRead

```go
func (a *GlGetAttribLocation) AddRead(rng memory.Range, id binary.ID) *GlGetAttribLocation
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetAttribLocation pointer is returned so that calls can be
chained.

#### func (*GlGetAttribLocation) AddWrite

```go
func (a *GlGetAttribLocation) AddWrite(rng memory.Range, id binary.ID) *GlGetAttribLocation
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetAttribLocation pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetAttribLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetAttribLocation) Observations

```go
func (a *GlGetAttribLocation) Observations() *atom.Observations
```

#### func (*GlGetAttribLocation) Replay

```go
func (ω *GlGetAttribLocation) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
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
func (a *GlGetAttribLocation) String() string
```

#### type GlGetBooleanv

```go
type GlGetBooleanv struct {
	binary.Generate

	Param  StateVariable
	Values Boolᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetBooleanv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetBooleanv

```go
func NewGlGetBooleanv(Param StateVariable, Values memory.Pointer) *GlGetBooleanv
```

#### func (*GlGetBooleanv) API

```go
func (c *GlGetBooleanv) API() gfxapi.ID
```

#### func (*GlGetBooleanv) AddRead

```go
func (a *GlGetBooleanv) AddRead(rng memory.Range, id binary.ID) *GlGetBooleanv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetBooleanv pointer is returned so that calls can be chained.

#### func (*GlGetBooleanv) AddWrite

```go
func (a *GlGetBooleanv) AddWrite(rng memory.Range, id binary.ID) *GlGetBooleanv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetBooleanv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetBooleanv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetBooleanv) Observations

```go
func (a *GlGetBooleanv) Observations() *atom.Observations
```

#### func (*GlGetBooleanv) Replay

```go
func (ϟa *GlGetBooleanv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetBooleanv) String

```go
func (a *GlGetBooleanv) String() string
```

#### type GlGetBufferParameteriv

```go
type GlGetBufferParameteriv struct {
	binary.Generate

	Target    BufferTarget
	Parameter BufferParameter
	Value     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetBufferParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetBufferParameteriv

```go
func NewGlGetBufferParameteriv(Target BufferTarget, Parameter BufferParameter, Value memory.Pointer) *GlGetBufferParameteriv
```

#### func (*GlGetBufferParameteriv) API

```go
func (c *GlGetBufferParameteriv) API() gfxapi.ID
```

#### func (*GlGetBufferParameteriv) AddRead

```go
func (a *GlGetBufferParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetBufferParameteriv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetBufferParameteriv pointer is returned so that calls can be
chained.

#### func (*GlGetBufferParameteriv) AddWrite

```go
func (a *GlGetBufferParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetBufferParameteriv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetBufferParameteriv pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetBufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetBufferParameteriv) Observations

```go
func (a *GlGetBufferParameteriv) Observations() *atom.Observations
```

#### func (*GlGetBufferParameteriv) Replay

```go
func (ϟa *GlGetBufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetBufferParameteriv) String

```go
func (a *GlGetBufferParameteriv) String() string
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
func NewGlGetError(Result Error) *GlGetError
```

#### func (*GlGetError) API

```go
func (c *GlGetError) API() gfxapi.ID
```

#### func (*GlGetError) AddRead

```go
func (a *GlGetError) AddRead(rng memory.Range, id binary.ID) *GlGetError
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetError pointer is returned so that calls can be chained.

#### func (*GlGetError) AddWrite

```go
func (a *GlGetError) AddWrite(rng memory.Range, id binary.ID) *GlGetError
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetError pointer is returned so that calls can be chained.

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
func (ϟa *GlGetError) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetError) Observations

```go
func (a *GlGetError) Observations() *atom.Observations
```

#### func (*GlGetError) Replay

```go
func (ϟa *GlGetError) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetError) String

```go
func (a *GlGetError) String() string
```

#### type GlGetFloatv

```go
type GlGetFloatv struct {
	binary.Generate

	Param  StateVariable
	Values F32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetFloatv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetFloatv

```go
func NewGlGetFloatv(Param StateVariable, Values memory.Pointer) *GlGetFloatv
```

#### func (*GlGetFloatv) API

```go
func (c *GlGetFloatv) API() gfxapi.ID
```

#### func (*GlGetFloatv) AddRead

```go
func (a *GlGetFloatv) AddRead(rng memory.Range, id binary.ID) *GlGetFloatv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetFloatv pointer is returned so that calls can be chained.

#### func (*GlGetFloatv) AddWrite

```go
func (a *GlGetFloatv) AddWrite(rng memory.Range, id binary.ID) *GlGetFloatv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetFloatv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetFloatv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetFloatv) Observations

```go
func (a *GlGetFloatv) Observations() *atom.Observations
```

#### func (*GlGetFloatv) Replay

```go
func (ϟa *GlGetFloatv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetFloatv) String

```go
func (a *GlGetFloatv) String() string
```

#### type GlGetFramebufferAttachmentParameteriv

```go
type GlGetFramebufferAttachmentParameteriv struct {
	binary.Generate

	FramebufferTarget FramebufferTarget
	Attachment        FramebufferAttachment
	Parameter         FramebufferAttachmentParameter
	Value             S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetFramebufferAttachmentParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetFramebufferAttachmentParameteriv

```go
func NewGlGetFramebufferAttachmentParameteriv(Framebuffer_target FramebufferTarget, Attachment FramebufferAttachment, Parameter FramebufferAttachmentParameter, Value memory.Pointer) *GlGetFramebufferAttachmentParameteriv
```

#### func (*GlGetFramebufferAttachmentParameteriv) API

```go
func (c *GlGetFramebufferAttachmentParameteriv) API() gfxapi.ID
```

#### func (*GlGetFramebufferAttachmentParameteriv) AddRead

```go
func (a *GlGetFramebufferAttachmentParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetFramebufferAttachmentParameteriv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetFramebufferAttachmentParameteriv pointer is returned so that
calls can be chained.

#### func (*GlGetFramebufferAttachmentParameteriv) AddWrite

```go
func (a *GlGetFramebufferAttachmentParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetFramebufferAttachmentParameteriv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetFramebufferAttachmentParameteriv pointer is returned so that
calls can be chained.

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
func (ϟa *GlGetFramebufferAttachmentParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetFramebufferAttachmentParameteriv) Observations

```go
func (a *GlGetFramebufferAttachmentParameteriv) Observations() *atom.Observations
```

#### func (*GlGetFramebufferAttachmentParameteriv) Replay

```go
func (ϟa *GlGetFramebufferAttachmentParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetFramebufferAttachmentParameteriv) String

```go
func (a *GlGetFramebufferAttachmentParameteriv) String() string
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
func NewGlGetGraphicsResetStatusEXT(Result ResetStatus) *GlGetGraphicsResetStatusEXT
```

#### func (*GlGetGraphicsResetStatusEXT) API

```go
func (c *GlGetGraphicsResetStatusEXT) API() gfxapi.ID
```

#### func (*GlGetGraphicsResetStatusEXT) AddRead

```go
func (a *GlGetGraphicsResetStatusEXT) AddRead(rng memory.Range, id binary.ID) *GlGetGraphicsResetStatusEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetGraphicsResetStatusEXT pointer is returned so that calls can
be chained.

#### func (*GlGetGraphicsResetStatusEXT) AddWrite

```go
func (a *GlGetGraphicsResetStatusEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetGraphicsResetStatusEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetGraphicsResetStatusEXT pointer is returned so that calls can
be chained.

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
func (ϟa *GlGetGraphicsResetStatusEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetGraphicsResetStatusEXT) Observations

```go
func (a *GlGetGraphicsResetStatusEXT) Observations() *atom.Observations
```

#### func (*GlGetGraphicsResetStatusEXT) Replay

```go
func (ϟa *GlGetGraphicsResetStatusEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetGraphicsResetStatusEXT) String

```go
func (a *GlGetGraphicsResetStatusEXT) String() string
```

#### type GlGetIntegerv

```go
type GlGetIntegerv struct {
	binary.Generate

	Param  StateVariable
	Values S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetIntegerv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetIntegerv

```go
func NewGlGetIntegerv(Param StateVariable, Values memory.Pointer) *GlGetIntegerv
```

#### func (*GlGetIntegerv) API

```go
func (c *GlGetIntegerv) API() gfxapi.ID
```

#### func (*GlGetIntegerv) AddRead

```go
func (a *GlGetIntegerv) AddRead(rng memory.Range, id binary.ID) *GlGetIntegerv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetIntegerv pointer is returned so that calls can be chained.

#### func (*GlGetIntegerv) AddWrite

```go
func (a *GlGetIntegerv) AddWrite(rng memory.Range, id binary.ID) *GlGetIntegerv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetIntegerv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetIntegerv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetIntegerv) Observations

```go
func (a *GlGetIntegerv) Observations() *atom.Observations
```

#### func (*GlGetIntegerv) Replay

```go
func (ϟa *GlGetIntegerv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetIntegerv) String

```go
func (a *GlGetIntegerv) String() string
```

#### type GlGetProgramBinaryOES

```go
type GlGetProgramBinaryOES struct {
	binary.Generate

	Program      ProgramId
	BufferSize   int32
	BytesWritten S32ᵖ
	BinaryFormat U32ᵖ
	Binary       Voidᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetProgramBinaryOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetProgramBinaryOES

```go
func NewGlGetProgramBinaryOES(Program ProgramId, Buffer_size int32, Bytes_written memory.Pointer, Binary_format memory.Pointer, Binary memory.Pointer) *GlGetProgramBinaryOES
```

#### func (*GlGetProgramBinaryOES) API

```go
func (c *GlGetProgramBinaryOES) API() gfxapi.ID
```

#### func (*GlGetProgramBinaryOES) AddRead

```go
func (a *GlGetProgramBinaryOES) AddRead(rng memory.Range, id binary.ID) *GlGetProgramBinaryOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetProgramBinaryOES pointer is returned so that calls can be
chained.

#### func (*GlGetProgramBinaryOES) AddWrite

```go
func (a *GlGetProgramBinaryOES) AddWrite(rng memory.Range, id binary.ID) *GlGetProgramBinaryOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetProgramBinaryOES pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetProgramBinaryOES) Observations

```go
func (a *GlGetProgramBinaryOES) Observations() *atom.Observations
```

#### func (*GlGetProgramBinaryOES) Replay

```go
func (ϟa *GlGetProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetProgramBinaryOES) String

```go
func (a *GlGetProgramBinaryOES) String() string
```

#### type GlGetProgramInfoLog

```go
type GlGetProgramInfoLog struct {
	binary.Generate

	Program             ProgramId
	BufferLength        int32
	StringLengthWritten S32ᵖ
	Info                Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetProgramInfoLog
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetProgramInfoLog

```go
func NewGlGetProgramInfoLog(Program ProgramId, Buffer_length int32, String_length_written memory.Pointer, Info memory.Pointer) *GlGetProgramInfoLog
```

#### func (*GlGetProgramInfoLog) API

```go
func (c *GlGetProgramInfoLog) API() gfxapi.ID
```

#### func (*GlGetProgramInfoLog) AddRead

```go
func (a *GlGetProgramInfoLog) AddRead(rng memory.Range, id binary.ID) *GlGetProgramInfoLog
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetProgramInfoLog pointer is returned so that calls can be
chained.

#### func (*GlGetProgramInfoLog) AddWrite

```go
func (a *GlGetProgramInfoLog) AddWrite(rng memory.Range, id binary.ID) *GlGetProgramInfoLog
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetProgramInfoLog pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetProgramInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetProgramInfoLog) Observations

```go
func (a *GlGetProgramInfoLog) Observations() *atom.Observations
```

#### func (*GlGetProgramInfoLog) Replay

```go
func (ϟa *GlGetProgramInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetProgramInfoLog) String

```go
func (a *GlGetProgramInfoLog) String() string
```

#### type GlGetProgramiv

```go
type GlGetProgramiv struct {
	binary.Generate

	Program   ProgramId
	Parameter ProgramParameter
	Value     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetProgramiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetProgramiv

```go
func NewGlGetProgramiv(Program ProgramId, Parameter ProgramParameter, Value memory.Pointer) *GlGetProgramiv
```

#### func (*GlGetProgramiv) API

```go
func (c *GlGetProgramiv) API() gfxapi.ID
```

#### func (*GlGetProgramiv) AddRead

```go
func (a *GlGetProgramiv) AddRead(rng memory.Range, id binary.ID) *GlGetProgramiv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetProgramiv pointer is returned so that calls can be chained.

#### func (*GlGetProgramiv) AddWrite

```go
func (a *GlGetProgramiv) AddWrite(rng memory.Range, id binary.ID) *GlGetProgramiv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetProgramiv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetProgramiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetProgramiv) Observations

```go
func (a *GlGetProgramiv) Observations() *atom.Observations
```

#### func (*GlGetProgramiv) Replay

```go
func (ϟa *GlGetProgramiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetProgramiv) String

```go
func (a *GlGetProgramiv) String() string
```

#### type GlGetQueryObjecti64v

```go
type GlGetQueryObjecti64v struct {
	binary.Generate

	Query     QueryId
	Parameter QueryObjectParameter
	Value     S64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjecti64v
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjecti64v

```go
func NewGlGetQueryObjecti64v(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjecti64v
```

#### func (*GlGetQueryObjecti64v) API

```go
func (c *GlGetQueryObjecti64v) API() gfxapi.ID
```

#### func (*GlGetQueryObjecti64v) AddRead

```go
func (a *GlGetQueryObjecti64v) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjecti64v
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryObjecti64v pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjecti64v) AddWrite

```go
func (a *GlGetQueryObjecti64v) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjecti64v
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryObjecti64v pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjecti64v) Class

```go
func (*GlGetQueryObjecti64v) Class() binary.Class
```

#### func (*GlGetQueryObjecti64v) Flags

```go
func (c *GlGetQueryObjecti64v) Flags() atom.Flags
```

#### func (*GlGetQueryObjecti64v) Mutate

```go
func (ϟa *GlGetQueryObjecti64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryObjecti64v) Observations

```go
func (a *GlGetQueryObjecti64v) Observations() *atom.Observations
```

#### func (*GlGetQueryObjecti64v) Replay

```go
func (ϟa *GlGetQueryObjecti64v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryObjecti64v) String

```go
func (a *GlGetQueryObjecti64v) String() string
```

#### type GlGetQueryObjecti64vEXT

```go
type GlGetQueryObjecti64vEXT struct {
	binary.Generate

	Query     QueryId
	Parameter QueryObjectParameter
	Value     S64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjecti64vEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjecti64vEXT

```go
func NewGlGetQueryObjecti64vEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjecti64vEXT
```

#### func (*GlGetQueryObjecti64vEXT) API

```go
func (c *GlGetQueryObjecti64vEXT) API() gfxapi.ID
```

#### func (*GlGetQueryObjecti64vEXT) AddRead

```go
func (a *GlGetQueryObjecti64vEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjecti64vEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryObjecti64vEXT pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjecti64vEXT) AddWrite

```go
func (a *GlGetQueryObjecti64vEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjecti64vEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryObjecti64vEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetQueryObjecti64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryObjecti64vEXT) Observations

```go
func (a *GlGetQueryObjecti64vEXT) Observations() *atom.Observations
```

#### func (*GlGetQueryObjecti64vEXT) Replay

```go
func (ϟa *GlGetQueryObjecti64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryObjecti64vEXT) String

```go
func (a *GlGetQueryObjecti64vEXT) String() string
```

#### type GlGetQueryObjectivEXT

```go
type GlGetQueryObjectivEXT struct {
	binary.Generate

	Query     QueryId
	Parameter QueryObjectParameter
	Value     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectivEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectivEXT

```go
func NewGlGetQueryObjectivEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectivEXT
```

#### func (*GlGetQueryObjectivEXT) API

```go
func (c *GlGetQueryObjectivEXT) API() gfxapi.ID
```

#### func (*GlGetQueryObjectivEXT) AddRead

```go
func (a *GlGetQueryObjectivEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectivEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryObjectivEXT pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjectivEXT) AddWrite

```go
func (a *GlGetQueryObjectivEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectivEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryObjectivEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetQueryObjectivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryObjectivEXT) Observations

```go
func (a *GlGetQueryObjectivEXT) Observations() *atom.Observations
```

#### func (*GlGetQueryObjectivEXT) Replay

```go
func (ϟa *GlGetQueryObjectivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryObjectivEXT) String

```go
func (a *GlGetQueryObjectivEXT) String() string
```

#### type GlGetQueryObjectui64v

```go
type GlGetQueryObjectui64v struct {
	binary.Generate

	Query     QueryId
	Parameter QueryObjectParameter
	Value     U64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectui64v
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectui64v

```go
func NewGlGetQueryObjectui64v(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectui64v
```

#### func (*GlGetQueryObjectui64v) API

```go
func (c *GlGetQueryObjectui64v) API() gfxapi.ID
```

#### func (*GlGetQueryObjectui64v) AddRead

```go
func (a *GlGetQueryObjectui64v) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectui64v
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryObjectui64v pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjectui64v) AddWrite

```go
func (a *GlGetQueryObjectui64v) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectui64v
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryObjectui64v pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjectui64v) Class

```go
func (*GlGetQueryObjectui64v) Class() binary.Class
```

#### func (*GlGetQueryObjectui64v) Flags

```go
func (c *GlGetQueryObjectui64v) Flags() atom.Flags
```

#### func (*GlGetQueryObjectui64v) Mutate

```go
func (ϟa *GlGetQueryObjectui64v) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryObjectui64v) Observations

```go
func (a *GlGetQueryObjectui64v) Observations() *atom.Observations
```

#### func (*GlGetQueryObjectui64v) Replay

```go
func (ϟa *GlGetQueryObjectui64v) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryObjectui64v) String

```go
func (a *GlGetQueryObjectui64v) String() string
```

#### type GlGetQueryObjectui64vEXT

```go
type GlGetQueryObjectui64vEXT struct {
	binary.Generate

	Query     QueryId
	Parameter QueryObjectParameter
	Value     U64ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectui64vEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectui64vEXT

```go
func NewGlGetQueryObjectui64vEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectui64vEXT
```

#### func (*GlGetQueryObjectui64vEXT) API

```go
func (c *GlGetQueryObjectui64vEXT) API() gfxapi.ID
```

#### func (*GlGetQueryObjectui64vEXT) AddRead

```go
func (a *GlGetQueryObjectui64vEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectui64vEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryObjectui64vEXT pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjectui64vEXT) AddWrite

```go
func (a *GlGetQueryObjectui64vEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectui64vEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryObjectui64vEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetQueryObjectui64vEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryObjectui64vEXT) Observations

```go
func (a *GlGetQueryObjectui64vEXT) Observations() *atom.Observations
```

#### func (*GlGetQueryObjectui64vEXT) Replay

```go
func (ϟa *GlGetQueryObjectui64vEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryObjectui64vEXT) String

```go
func (a *GlGetQueryObjectui64vEXT) String() string
```

#### type GlGetQueryObjectuiv

```go
type GlGetQueryObjectuiv struct {
	binary.Generate

	Query     QueryId
	Parameter QueryObjectParameter
	Value     U32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectuiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectuiv

```go
func NewGlGetQueryObjectuiv(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectuiv
```

#### func (*GlGetQueryObjectuiv) API

```go
func (c *GlGetQueryObjectuiv) API() gfxapi.ID
```

#### func (*GlGetQueryObjectuiv) AddRead

```go
func (a *GlGetQueryObjectuiv) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectuiv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryObjectuiv pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjectuiv) AddWrite

```go
func (a *GlGetQueryObjectuiv) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectuiv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryObjectuiv pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetQueryObjectuiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryObjectuiv) Observations

```go
func (a *GlGetQueryObjectuiv) Observations() *atom.Observations
```

#### func (*GlGetQueryObjectuiv) Replay

```go
func (ϟa *GlGetQueryObjectuiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryObjectuiv) String

```go
func (a *GlGetQueryObjectuiv) String() string
```

#### type GlGetQueryObjectuivEXT

```go
type GlGetQueryObjectuivEXT struct {
	binary.Generate

	Query     QueryId
	Parameter QueryObjectParameter
	Value     U32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryObjectuivEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryObjectuivEXT

```go
func NewGlGetQueryObjectuivEXT(Query QueryId, Parameter QueryObjectParameter, Value memory.Pointer) *GlGetQueryObjectuivEXT
```

#### func (*GlGetQueryObjectuivEXT) API

```go
func (c *GlGetQueryObjectuivEXT) API() gfxapi.ID
```

#### func (*GlGetQueryObjectuivEXT) AddRead

```go
func (a *GlGetQueryObjectuivEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryObjectuivEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryObjectuivEXT pointer is returned so that calls can be
chained.

#### func (*GlGetQueryObjectuivEXT) AddWrite

```go
func (a *GlGetQueryObjectuivEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryObjectuivEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryObjectuivEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetQueryObjectuivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryObjectuivEXT) Observations

```go
func (a *GlGetQueryObjectuivEXT) Observations() *atom.Observations
```

#### func (*GlGetQueryObjectuivEXT) Replay

```go
func (ϟa *GlGetQueryObjectuivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryObjectuivEXT) String

```go
func (a *GlGetQueryObjectuivEXT) String() string
```

#### type GlGetQueryiv

```go
type GlGetQueryiv struct {
	binary.Generate

	Target    QueryTarget
	Parameter QueryParameter
	Value     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryiv

```go
func NewGlGetQueryiv(Target QueryTarget, Parameter QueryParameter, Value memory.Pointer) *GlGetQueryiv
```

#### func (*GlGetQueryiv) API

```go
func (c *GlGetQueryiv) API() gfxapi.ID
```

#### func (*GlGetQueryiv) AddRead

```go
func (a *GlGetQueryiv) AddRead(rng memory.Range, id binary.ID) *GlGetQueryiv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryiv pointer is returned so that calls can be chained.

#### func (*GlGetQueryiv) AddWrite

```go
func (a *GlGetQueryiv) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryiv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryiv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetQueryiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryiv) Observations

```go
func (a *GlGetQueryiv) Observations() *atom.Observations
```

#### func (*GlGetQueryiv) Replay

```go
func (ϟa *GlGetQueryiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryiv) String

```go
func (a *GlGetQueryiv) String() string
```

#### type GlGetQueryivEXT

```go
type GlGetQueryivEXT struct {
	binary.Generate

	Target    QueryTarget
	Parameter QueryParameter
	Value     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetQueryivEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetQueryivEXT

```go
func NewGlGetQueryivEXT(Target QueryTarget, Parameter QueryParameter, Value memory.Pointer) *GlGetQueryivEXT
```

#### func (*GlGetQueryivEXT) API

```go
func (c *GlGetQueryivEXT) API() gfxapi.ID
```

#### func (*GlGetQueryivEXT) AddRead

```go
func (a *GlGetQueryivEXT) AddRead(rng memory.Range, id binary.ID) *GlGetQueryivEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetQueryivEXT pointer is returned so that calls can be chained.

#### func (*GlGetQueryivEXT) AddWrite

```go
func (a *GlGetQueryivEXT) AddWrite(rng memory.Range, id binary.ID) *GlGetQueryivEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetQueryivEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlGetQueryivEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetQueryivEXT) Observations

```go
func (a *GlGetQueryivEXT) Observations() *atom.Observations
```

#### func (*GlGetQueryivEXT) Replay

```go
func (ϟa *GlGetQueryivEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetQueryivEXT) String

```go
func (a *GlGetQueryivEXT) String() string
```

#### type GlGetRenderbufferParameteriv

```go
type GlGetRenderbufferParameteriv struct {
	binary.Generate

	Target    RenderbufferTarget
	Parameter RenderbufferParameter
	Values    S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetRenderbufferParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetRenderbufferParameteriv

```go
func NewGlGetRenderbufferParameteriv(Target RenderbufferTarget, Parameter RenderbufferParameter, Values memory.Pointer) *GlGetRenderbufferParameteriv
```

#### func (*GlGetRenderbufferParameteriv) API

```go
func (c *GlGetRenderbufferParameteriv) API() gfxapi.ID
```

#### func (*GlGetRenderbufferParameteriv) AddRead

```go
func (a *GlGetRenderbufferParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetRenderbufferParameteriv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetRenderbufferParameteriv pointer is returned so that calls can
be chained.

#### func (*GlGetRenderbufferParameteriv) AddWrite

```go
func (a *GlGetRenderbufferParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetRenderbufferParameteriv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetRenderbufferParameteriv pointer is returned so that calls can
be chained.

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
func (ϟa *GlGetRenderbufferParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetRenderbufferParameteriv) Observations

```go
func (a *GlGetRenderbufferParameteriv) Observations() *atom.Observations
```

#### func (*GlGetRenderbufferParameteriv) Replay

```go
func (ϟa *GlGetRenderbufferParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetRenderbufferParameteriv) String

```go
func (a *GlGetRenderbufferParameteriv) String() string
```

#### type GlGetShaderInfoLog

```go
type GlGetShaderInfoLog struct {
	binary.Generate

	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten S32ᵖ
	Info                Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderInfoLog
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderInfoLog

```go
func NewGlGetShaderInfoLog(Shader ShaderId, Buffer_length int32, String_length_written memory.Pointer, Info memory.Pointer) *GlGetShaderInfoLog
```

#### func (*GlGetShaderInfoLog) API

```go
func (c *GlGetShaderInfoLog) API() gfxapi.ID
```

#### func (*GlGetShaderInfoLog) AddRead

```go
func (a *GlGetShaderInfoLog) AddRead(rng memory.Range, id binary.ID) *GlGetShaderInfoLog
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetShaderInfoLog pointer is returned so that calls can be
chained.

#### func (*GlGetShaderInfoLog) AddWrite

```go
func (a *GlGetShaderInfoLog) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderInfoLog
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetShaderInfoLog pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetShaderInfoLog) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetShaderInfoLog) Observations

```go
func (a *GlGetShaderInfoLog) Observations() *atom.Observations
```

#### func (*GlGetShaderInfoLog) Replay

```go
func (ϟa *GlGetShaderInfoLog) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetShaderInfoLog) String

```go
func (a *GlGetShaderInfoLog) String() string
```

#### type GlGetShaderPrecisionFormat

```go
type GlGetShaderPrecisionFormat struct {
	binary.Generate

	ShaderType    ShaderType
	PrecisionType PrecisionType
	Range         S32ᵖ
	Precision     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderPrecisionFormat
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderPrecisionFormat

```go
func NewGlGetShaderPrecisionFormat(Shader_type ShaderType, Precision_type PrecisionType, Range memory.Pointer, Precision memory.Pointer) *GlGetShaderPrecisionFormat
```

#### func (*GlGetShaderPrecisionFormat) API

```go
func (c *GlGetShaderPrecisionFormat) API() gfxapi.ID
```

#### func (*GlGetShaderPrecisionFormat) AddRead

```go
func (a *GlGetShaderPrecisionFormat) AddRead(rng memory.Range, id binary.ID) *GlGetShaderPrecisionFormat
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetShaderPrecisionFormat pointer is returned so that calls can be
chained.

#### func (*GlGetShaderPrecisionFormat) AddWrite

```go
func (a *GlGetShaderPrecisionFormat) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderPrecisionFormat
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetShaderPrecisionFormat pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetShaderPrecisionFormat) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetShaderPrecisionFormat) Observations

```go
func (a *GlGetShaderPrecisionFormat) Observations() *atom.Observations
```

#### func (*GlGetShaderPrecisionFormat) Replay

```go
func (ϟa *GlGetShaderPrecisionFormat) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetShaderPrecisionFormat) String

```go
func (a *GlGetShaderPrecisionFormat) String() string
```

#### type GlGetShaderSource

```go
type GlGetShaderSource struct {
	binary.Generate

	Shader              ShaderId
	BufferLength        int32
	StringLengthWritten S32ᵖ
	Source              Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderSource
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderSource

```go
func NewGlGetShaderSource(Shader ShaderId, Buffer_length int32, String_length_written memory.Pointer, Source memory.Pointer) *GlGetShaderSource
```

#### func (*GlGetShaderSource) API

```go
func (c *GlGetShaderSource) API() gfxapi.ID
```

#### func (*GlGetShaderSource) AddRead

```go
func (a *GlGetShaderSource) AddRead(rng memory.Range, id binary.ID) *GlGetShaderSource
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetShaderSource pointer is returned so that calls can be chained.

#### func (*GlGetShaderSource) AddWrite

```go
func (a *GlGetShaderSource) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderSource
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetShaderSource pointer is returned so that calls can be chained.

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
func (ϟa *GlGetShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetShaderSource) Observations

```go
func (a *GlGetShaderSource) Observations() *atom.Observations
```

#### func (*GlGetShaderSource) Replay

```go
func (ϟa *GlGetShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetShaderSource) String

```go
func (a *GlGetShaderSource) String() string
```

#### type GlGetShaderiv

```go
type GlGetShaderiv struct {
	binary.Generate

	Shader    ShaderId
	Parameter ShaderParameter
	Value     S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetShaderiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetShaderiv

```go
func NewGlGetShaderiv(Shader ShaderId, Parameter ShaderParameter, Value memory.Pointer) *GlGetShaderiv
```

#### func (*GlGetShaderiv) API

```go
func (c *GlGetShaderiv) API() gfxapi.ID
```

#### func (*GlGetShaderiv) AddRead

```go
func (a *GlGetShaderiv) AddRead(rng memory.Range, id binary.ID) *GlGetShaderiv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetShaderiv pointer is returned so that calls can be chained.

#### func (*GlGetShaderiv) AddWrite

```go
func (a *GlGetShaderiv) AddWrite(rng memory.Range, id binary.ID) *GlGetShaderiv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetShaderiv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetShaderiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetShaderiv) Observations

```go
func (a *GlGetShaderiv) Observations() *atom.Observations
```

#### func (*GlGetShaderiv) Replay

```go
func (ϟa *GlGetShaderiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetShaderiv) String

```go
func (a *GlGetShaderiv) String() string
```

#### type GlGetString

```go
type GlGetString struct {
	binary.Generate

	Param  StringConstant
	Result Charᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetString
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetString

```go
func NewGlGetString(Param StringConstant, Result memory.Pointer) *GlGetString
```

#### func (*GlGetString) API

```go
func (c *GlGetString) API() gfxapi.ID
```

#### func (*GlGetString) AddRead

```go
func (a *GlGetString) AddRead(rng memory.Range, id binary.ID) *GlGetString
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetString pointer is returned so that calls can be chained.

#### func (*GlGetString) AddWrite

```go
func (a *GlGetString) AddWrite(rng memory.Range, id binary.ID) *GlGetString
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetString pointer is returned so that calls can be chained.

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
func (ϟa *GlGetString) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetString) Observations

```go
func (a *GlGetString) Observations() *atom.Observations
```

#### func (*GlGetString) Replay

```go
func (ϟa *GlGetString) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetString) String

```go
func (a *GlGetString) String() string
```

#### type GlGetTexParameterfv

```go
type GlGetTexParameterfv struct {
	binary.Generate

	Target    TextureTarget
	Parameter TextureParameter
	Values    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetTexParameterfv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetTexParameterfv

```go
func NewGlGetTexParameterfv(Target TextureTarget, Parameter TextureParameter, Values memory.Pointer) *GlGetTexParameterfv
```

#### func (*GlGetTexParameterfv) API

```go
func (c *GlGetTexParameterfv) API() gfxapi.ID
```

#### func (*GlGetTexParameterfv) AddRead

```go
func (a *GlGetTexParameterfv) AddRead(rng memory.Range, id binary.ID) *GlGetTexParameterfv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetTexParameterfv pointer is returned so that calls can be
chained.

#### func (*GlGetTexParameterfv) AddWrite

```go
func (a *GlGetTexParameterfv) AddWrite(rng memory.Range, id binary.ID) *GlGetTexParameterfv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetTexParameterfv pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetTexParameterfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetTexParameterfv) Observations

```go
func (a *GlGetTexParameterfv) Observations() *atom.Observations
```

#### func (*GlGetTexParameterfv) Replay

```go
func (ϟa *GlGetTexParameterfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetTexParameterfv) String

```go
func (a *GlGetTexParameterfv) String() string
```

#### type GlGetTexParameteriv

```go
type GlGetTexParameteriv struct {
	binary.Generate

	Target    TextureTarget
	Parameter TextureParameter
	Values    S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetTexParameteriv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetTexParameteriv

```go
func NewGlGetTexParameteriv(Target TextureTarget, Parameter TextureParameter, Values memory.Pointer) *GlGetTexParameteriv
```

#### func (*GlGetTexParameteriv) API

```go
func (c *GlGetTexParameteriv) API() gfxapi.ID
```

#### func (*GlGetTexParameteriv) AddRead

```go
func (a *GlGetTexParameteriv) AddRead(rng memory.Range, id binary.ID) *GlGetTexParameteriv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetTexParameteriv pointer is returned so that calls can be
chained.

#### func (*GlGetTexParameteriv) AddWrite

```go
func (a *GlGetTexParameteriv) AddWrite(rng memory.Range, id binary.ID) *GlGetTexParameteriv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetTexParameteriv pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetTexParameteriv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetTexParameteriv) Observations

```go
func (a *GlGetTexParameteriv) Observations() *atom.Observations
```

#### func (*GlGetTexParameteriv) Replay

```go
func (ϟa *GlGetTexParameteriv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetTexParameteriv) String

```go
func (a *GlGetTexParameteriv) String() string
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
func NewGlGetUniformLocation(Program ProgramId, Name string, Result UniformLocation) *GlGetUniformLocation
```

#### func (*GlGetUniformLocation) API

```go
func (c *GlGetUniformLocation) API() gfxapi.ID
```

#### func (*GlGetUniformLocation) AddRead

```go
func (a *GlGetUniformLocation) AddRead(rng memory.Range, id binary.ID) *GlGetUniformLocation
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetUniformLocation pointer is returned so that calls can be
chained.

#### func (*GlGetUniformLocation) AddWrite

```go
func (a *GlGetUniformLocation) AddWrite(rng memory.Range, id binary.ID) *GlGetUniformLocation
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetUniformLocation pointer is returned so that calls can be
chained.

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
func (ϟa *GlGetUniformLocation) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetUniformLocation) Observations

```go
func (a *GlGetUniformLocation) Observations() *atom.Observations
```

#### func (*GlGetUniformLocation) Replay

```go
func (ϟa *GlGetUniformLocation) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetUniformLocation) String

```go
func (a *GlGetUniformLocation) String() string
```

#### type GlGetUniformfv

```go
type GlGetUniformfv struct {
	binary.Generate

	Program  ProgramId
	Location UniformLocation
	Values   F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetUniformfv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetUniformfv

```go
func NewGlGetUniformfv(Program ProgramId, Location UniformLocation, Values memory.Pointer) *GlGetUniformfv
```

#### func (*GlGetUniformfv) API

```go
func (c *GlGetUniformfv) API() gfxapi.ID
```

#### func (*GlGetUniformfv) AddRead

```go
func (a *GlGetUniformfv) AddRead(rng memory.Range, id binary.ID) *GlGetUniformfv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetUniformfv pointer is returned so that calls can be chained.

#### func (*GlGetUniformfv) AddWrite

```go
func (a *GlGetUniformfv) AddWrite(rng memory.Range, id binary.ID) *GlGetUniformfv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetUniformfv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetUniformfv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetUniformfv) Observations

```go
func (a *GlGetUniformfv) Observations() *atom.Observations
```

#### func (*GlGetUniformfv) Replay

```go
func (ϟa *GlGetUniformfv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetUniformfv) String

```go
func (a *GlGetUniformfv) String() string
```

#### type GlGetUniformiv

```go
type GlGetUniformiv struct {
	binary.Generate

	Program  ProgramId
	Location UniformLocation
	Values   S32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlGetUniformiv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlGetUniformiv

```go
func NewGlGetUniformiv(Program ProgramId, Location UniformLocation, Values memory.Pointer) *GlGetUniformiv
```

#### func (*GlGetUniformiv) API

```go
func (c *GlGetUniformiv) API() gfxapi.ID
```

#### func (*GlGetUniformiv) AddRead

```go
func (a *GlGetUniformiv) AddRead(rng memory.Range, id binary.ID) *GlGetUniformiv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlGetUniformiv pointer is returned so that calls can be chained.

#### func (*GlGetUniformiv) AddWrite

```go
func (a *GlGetUniformiv) AddWrite(rng memory.Range, id binary.ID) *GlGetUniformiv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlGetUniformiv pointer is returned so that calls can be chained.

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
func (ϟa *GlGetUniformiv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlGetUniformiv) Observations

```go
func (a *GlGetUniformiv) Observations() *atom.Observations
```

#### func (*GlGetUniformiv) Replay

```go
func (ϟa *GlGetUniformiv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlGetUniformiv) String

```go
func (a *GlGetUniformiv) String() string
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
func NewGlHint(Target HintTarget, Mode HintMode) *GlHint
```

#### func (*GlHint) API

```go
func (c *GlHint) API() gfxapi.ID
```

#### func (*GlHint) AddRead

```go
func (a *GlHint) AddRead(rng memory.Range, id binary.ID) *GlHint
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlHint pointer is returned so that calls can be chained.

#### func (*GlHint) AddWrite

```go
func (a *GlHint) AddWrite(rng memory.Range, id binary.ID) *GlHint
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlHint pointer is returned so that calls can be chained.

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
func (ϟa *GlHint) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlHint) Observations

```go
func (a *GlHint) Observations() *atom.Observations
```

#### func (*GlHint) Replay

```go
func (ϟa *GlHint) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlHint) String

```go
func (a *GlHint) String() string
```

#### type GlInsertEventMarkerEXT

```go
type GlInsertEventMarkerEXT struct {
	binary.Generate

	Length int32
	Marker Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlInsertEventMarkerEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlInsertEventMarkerEXT

```go
func NewGlInsertEventMarkerEXT(Length int32, Marker memory.Pointer) *GlInsertEventMarkerEXT
```

#### func (*GlInsertEventMarkerEXT) API

```go
func (c *GlInsertEventMarkerEXT) API() gfxapi.ID
```

#### func (*GlInsertEventMarkerEXT) AddRead

```go
func (a *GlInsertEventMarkerEXT) AddRead(rng memory.Range, id binary.ID) *GlInsertEventMarkerEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlInsertEventMarkerEXT pointer is returned so that calls can be
chained.

#### func (*GlInsertEventMarkerEXT) AddWrite

```go
func (a *GlInsertEventMarkerEXT) AddWrite(rng memory.Range, id binary.ID) *GlInsertEventMarkerEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlInsertEventMarkerEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlInsertEventMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlInsertEventMarkerEXT) Observations

```go
func (a *GlInsertEventMarkerEXT) Observations() *atom.Observations
```

#### func (*GlInsertEventMarkerEXT) Replay

```go
func (ϟa *GlInsertEventMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlInsertEventMarkerEXT) String

```go
func (a *GlInsertEventMarkerEXT) String() string
```

#### type GlInvalidateFramebuffer

```go
type GlInvalidateFramebuffer struct {
	binary.Generate

	Target      FramebufferTarget
	Count       int32
	Attachments FramebufferAttachmentᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlInvalidateFramebuffer
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlInvalidateFramebuffer

```go
func NewGlInvalidateFramebuffer(Target FramebufferTarget, Count int32, Attachments memory.Pointer) *GlInvalidateFramebuffer
```

#### func (*GlInvalidateFramebuffer) API

```go
func (c *GlInvalidateFramebuffer) API() gfxapi.ID
```

#### func (*GlInvalidateFramebuffer) AddRead

```go
func (a *GlInvalidateFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlInvalidateFramebuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlInvalidateFramebuffer pointer is returned so that calls can be
chained.

#### func (*GlInvalidateFramebuffer) AddWrite

```go
func (a *GlInvalidateFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlInvalidateFramebuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlInvalidateFramebuffer pointer is returned so that calls can be
chained.

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
func (ϟa *GlInvalidateFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlInvalidateFramebuffer) Observations

```go
func (a *GlInvalidateFramebuffer) Observations() *atom.Observations
```

#### func (*GlInvalidateFramebuffer) Replay

```go
func (ϟa *GlInvalidateFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlInvalidateFramebuffer) String

```go
func (a *GlInvalidateFramebuffer) String() string
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
func NewGlIsBuffer(Buffer BufferId, Result bool) *GlIsBuffer
```

#### func (*GlIsBuffer) API

```go
func (c *GlIsBuffer) API() gfxapi.ID
```

#### func (*GlIsBuffer) AddRead

```go
func (a *GlIsBuffer) AddRead(rng memory.Range, id binary.ID) *GlIsBuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsBuffer pointer is returned so that calls can be chained.

#### func (*GlIsBuffer) AddWrite

```go
func (a *GlIsBuffer) AddWrite(rng memory.Range, id binary.ID) *GlIsBuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsBuffer pointer is returned so that calls can be chained.

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
func (ϟa *GlIsBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsBuffer) Observations

```go
func (a *GlIsBuffer) Observations() *atom.Observations
```

#### func (*GlIsBuffer) Replay

```go
func (ϟa *GlIsBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsBuffer) String

```go
func (a *GlIsBuffer) String() string
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
func NewGlIsEnabled(Capability Capability, Result bool) *GlIsEnabled
```

#### func (*GlIsEnabled) API

```go
func (c *GlIsEnabled) API() gfxapi.ID
```

#### func (*GlIsEnabled) AddRead

```go
func (a *GlIsEnabled) AddRead(rng memory.Range, id binary.ID) *GlIsEnabled
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsEnabled pointer is returned so that calls can be chained.

#### func (*GlIsEnabled) AddWrite

```go
func (a *GlIsEnabled) AddWrite(rng memory.Range, id binary.ID) *GlIsEnabled
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsEnabled pointer is returned so that calls can be chained.

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
func (ϟa *GlIsEnabled) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsEnabled) Observations

```go
func (a *GlIsEnabled) Observations() *atom.Observations
```

#### func (*GlIsEnabled) Replay

```go
func (ϟa *GlIsEnabled) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsEnabled) String

```go
func (a *GlIsEnabled) String() string
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
func NewGlIsFramebuffer(Framebuffer FramebufferId, Result bool) *GlIsFramebuffer
```

#### func (*GlIsFramebuffer) API

```go
func (c *GlIsFramebuffer) API() gfxapi.ID
```

#### func (*GlIsFramebuffer) AddRead

```go
func (a *GlIsFramebuffer) AddRead(rng memory.Range, id binary.ID) *GlIsFramebuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsFramebuffer pointer is returned so that calls can be chained.

#### func (*GlIsFramebuffer) AddWrite

```go
func (a *GlIsFramebuffer) AddWrite(rng memory.Range, id binary.ID) *GlIsFramebuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsFramebuffer pointer is returned so that calls can be chained.

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
func (ϟa *GlIsFramebuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsFramebuffer) Observations

```go
func (a *GlIsFramebuffer) Observations() *atom.Observations
```

#### func (*GlIsFramebuffer) Replay

```go
func (ϟa *GlIsFramebuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsFramebuffer) String

```go
func (a *GlIsFramebuffer) String() string
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
func NewGlIsProgram(Program ProgramId, Result bool) *GlIsProgram
```

#### func (*GlIsProgram) API

```go
func (c *GlIsProgram) API() gfxapi.ID
```

#### func (*GlIsProgram) AddRead

```go
func (a *GlIsProgram) AddRead(rng memory.Range, id binary.ID) *GlIsProgram
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsProgram pointer is returned so that calls can be chained.

#### func (*GlIsProgram) AddWrite

```go
func (a *GlIsProgram) AddWrite(rng memory.Range, id binary.ID) *GlIsProgram
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsProgram pointer is returned so that calls can be chained.

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
func (ϟa *GlIsProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsProgram) Observations

```go
func (a *GlIsProgram) Observations() *atom.Observations
```

#### func (*GlIsProgram) Replay

```go
func (ϟa *GlIsProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsProgram) String

```go
func (a *GlIsProgram) String() string
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
func NewGlIsQuery(Query QueryId, Result bool) *GlIsQuery
```

#### func (*GlIsQuery) API

```go
func (c *GlIsQuery) API() gfxapi.ID
```

#### func (*GlIsQuery) AddRead

```go
func (a *GlIsQuery) AddRead(rng memory.Range, id binary.ID) *GlIsQuery
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsQuery pointer is returned so that calls can be chained.

#### func (*GlIsQuery) AddWrite

```go
func (a *GlIsQuery) AddWrite(rng memory.Range, id binary.ID) *GlIsQuery
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsQuery pointer is returned so that calls can be chained.

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
func (ϟa *GlIsQuery) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsQuery) Observations

```go
func (a *GlIsQuery) Observations() *atom.Observations
```

#### func (*GlIsQuery) Replay

```go
func (ϟa *GlIsQuery) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsQuery) String

```go
func (a *GlIsQuery) String() string
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
func NewGlIsQueryEXT(Query QueryId, Result bool) *GlIsQueryEXT
```

#### func (*GlIsQueryEXT) API

```go
func (c *GlIsQueryEXT) API() gfxapi.ID
```

#### func (*GlIsQueryEXT) AddRead

```go
func (a *GlIsQueryEXT) AddRead(rng memory.Range, id binary.ID) *GlIsQueryEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsQueryEXT pointer is returned so that calls can be chained.

#### func (*GlIsQueryEXT) AddWrite

```go
func (a *GlIsQueryEXT) AddWrite(rng memory.Range, id binary.ID) *GlIsQueryEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsQueryEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlIsQueryEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsQueryEXT) Observations

```go
func (a *GlIsQueryEXT) Observations() *atom.Observations
```

#### func (*GlIsQueryEXT) Replay

```go
func (ϟa *GlIsQueryEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsQueryEXT) String

```go
func (a *GlIsQueryEXT) String() string
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
func NewGlIsRenderbuffer(Renderbuffer RenderbufferId, Result bool) *GlIsRenderbuffer
```

#### func (*GlIsRenderbuffer) API

```go
func (c *GlIsRenderbuffer) API() gfxapi.ID
```

#### func (*GlIsRenderbuffer) AddRead

```go
func (a *GlIsRenderbuffer) AddRead(rng memory.Range, id binary.ID) *GlIsRenderbuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsRenderbuffer pointer is returned so that calls can be chained.

#### func (*GlIsRenderbuffer) AddWrite

```go
func (a *GlIsRenderbuffer) AddWrite(rng memory.Range, id binary.ID) *GlIsRenderbuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsRenderbuffer pointer is returned so that calls can be chained.

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
func (ϟa *GlIsRenderbuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsRenderbuffer) Observations

```go
func (a *GlIsRenderbuffer) Observations() *atom.Observations
```

#### func (*GlIsRenderbuffer) Replay

```go
func (ϟa *GlIsRenderbuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsRenderbuffer) String

```go
func (a *GlIsRenderbuffer) String() string
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
func NewGlIsShader(Shader ShaderId, Result bool) *GlIsShader
```

#### func (*GlIsShader) API

```go
func (c *GlIsShader) API() gfxapi.ID
```

#### func (*GlIsShader) AddRead

```go
func (a *GlIsShader) AddRead(rng memory.Range, id binary.ID) *GlIsShader
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsShader pointer is returned so that calls can be chained.

#### func (*GlIsShader) AddWrite

```go
func (a *GlIsShader) AddWrite(rng memory.Range, id binary.ID) *GlIsShader
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsShader pointer is returned so that calls can be chained.

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
func (ϟa *GlIsShader) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsShader) Observations

```go
func (a *GlIsShader) Observations() *atom.Observations
```

#### func (*GlIsShader) Replay

```go
func (ϟa *GlIsShader) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsShader) String

```go
func (a *GlIsShader) String() string
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
func NewGlIsTexture(Texture TextureId, Result bool) *GlIsTexture
```

#### func (*GlIsTexture) API

```go
func (c *GlIsTexture) API() gfxapi.ID
```

#### func (*GlIsTexture) AddRead

```go
func (a *GlIsTexture) AddRead(rng memory.Range, id binary.ID) *GlIsTexture
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsTexture pointer is returned so that calls can be chained.

#### func (*GlIsTexture) AddWrite

```go
func (a *GlIsTexture) AddWrite(rng memory.Range, id binary.ID) *GlIsTexture
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsTexture pointer is returned so that calls can be chained.

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
func (ϟa *GlIsTexture) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsTexture) Observations

```go
func (a *GlIsTexture) Observations() *atom.Observations
```

#### func (*GlIsTexture) Replay

```go
func (ϟa *GlIsTexture) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsTexture) String

```go
func (a *GlIsTexture) String() string
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
func NewGlIsVertexArrayOES(Array VertexArrayId, Result bool) *GlIsVertexArrayOES
```

#### func (*GlIsVertexArrayOES) API

```go
func (c *GlIsVertexArrayOES) API() gfxapi.ID
```

#### func (*GlIsVertexArrayOES) AddRead

```go
func (a *GlIsVertexArrayOES) AddRead(rng memory.Range, id binary.ID) *GlIsVertexArrayOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlIsVertexArrayOES pointer is returned so that calls can be
chained.

#### func (*GlIsVertexArrayOES) AddWrite

```go
func (a *GlIsVertexArrayOES) AddWrite(rng memory.Range, id binary.ID) *GlIsVertexArrayOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlIsVertexArrayOES pointer is returned so that calls can be
chained.

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
func (ϟa *GlIsVertexArrayOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlIsVertexArrayOES) Observations

```go
func (a *GlIsVertexArrayOES) Observations() *atom.Observations
```

#### func (*GlIsVertexArrayOES) Replay

```go
func (ϟa *GlIsVertexArrayOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlIsVertexArrayOES) String

```go
func (a *GlIsVertexArrayOES) String() string
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
func NewGlLineWidth(Width float32) *GlLineWidth
```

#### func (*GlLineWidth) API

```go
func (c *GlLineWidth) API() gfxapi.ID
```

#### func (*GlLineWidth) AddRead

```go
func (a *GlLineWidth) AddRead(rng memory.Range, id binary.ID) *GlLineWidth
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlLineWidth pointer is returned so that calls can be chained.

#### func (*GlLineWidth) AddWrite

```go
func (a *GlLineWidth) AddWrite(rng memory.Range, id binary.ID) *GlLineWidth
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlLineWidth pointer is returned so that calls can be chained.

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
func (ϟa *GlLineWidth) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlLineWidth) Observations

```go
func (a *GlLineWidth) Observations() *atom.Observations
```

#### func (*GlLineWidth) Replay

```go
func (ϟa *GlLineWidth) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlLineWidth) String

```go
func (a *GlLineWidth) String() string
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
func NewGlLinkProgram(Program ProgramId) *GlLinkProgram
```

#### func (*GlLinkProgram) API

```go
func (c *GlLinkProgram) API() gfxapi.ID
```

#### func (*GlLinkProgram) AddRead

```go
func (a *GlLinkProgram) AddRead(rng memory.Range, id binary.ID) *GlLinkProgram
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlLinkProgram pointer is returned so that calls can be chained.

#### func (*GlLinkProgram) AddWrite

```go
func (a *GlLinkProgram) AddWrite(rng memory.Range, id binary.ID) *GlLinkProgram
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlLinkProgram pointer is returned so that calls can be chained.

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
func (ϟa *GlLinkProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlLinkProgram) Observations

```go
func (a *GlLinkProgram) Observations() *atom.Observations
```

#### func (*GlLinkProgram) Replay

```go
func (ϟa *GlLinkProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlLinkProgram) String

```go
func (a *GlLinkProgram) String() string
```

#### type GlMapBufferRange

```go
type GlMapBufferRange struct {
	binary.Generate

	Target BufferTarget
	Offset int32
	Length int32
	Access MapBufferRangeAccess
	Result Voidᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlMapBufferRange
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlMapBufferRange

```go
func NewGlMapBufferRange(Target BufferTarget, Offset int32, Length int32, Access MapBufferRangeAccess, Result memory.Pointer) *GlMapBufferRange
```

#### func (*GlMapBufferRange) API

```go
func (c *GlMapBufferRange) API() gfxapi.ID
```

#### func (*GlMapBufferRange) AddRead

```go
func (a *GlMapBufferRange) AddRead(rng memory.Range, id binary.ID) *GlMapBufferRange
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlMapBufferRange pointer is returned so that calls can be chained.

#### func (*GlMapBufferRange) AddWrite

```go
func (a *GlMapBufferRange) AddWrite(rng memory.Range, id binary.ID) *GlMapBufferRange
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlMapBufferRange pointer is returned so that calls can be chained.

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
func (ϟa *GlMapBufferRange) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlMapBufferRange) Observations

```go
func (a *GlMapBufferRange) Observations() *atom.Observations
```

#### func (*GlMapBufferRange) Replay

```go
func (ϟa *GlMapBufferRange) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlMapBufferRange) String

```go
func (a *GlMapBufferRange) String() string
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
func NewGlPixelStorei(Parameter PixelStoreParameter, Value int32) *GlPixelStorei
```

#### func (*GlPixelStorei) API

```go
func (c *GlPixelStorei) API() gfxapi.ID
```

#### func (*GlPixelStorei) AddRead

```go
func (a *GlPixelStorei) AddRead(rng memory.Range, id binary.ID) *GlPixelStorei
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlPixelStorei pointer is returned so that calls can be chained.

#### func (*GlPixelStorei) AddWrite

```go
func (a *GlPixelStorei) AddWrite(rng memory.Range, id binary.ID) *GlPixelStorei
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlPixelStorei pointer is returned so that calls can be chained.

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
func (ϟa *GlPixelStorei) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlPixelStorei) Observations

```go
func (a *GlPixelStorei) Observations() *atom.Observations
```

#### func (*GlPixelStorei) Replay

```go
func (ϟa *GlPixelStorei) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlPixelStorei) String

```go
func (a *GlPixelStorei) String() string
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
func NewGlPolygonOffset(Scale_factor float32, Units float32) *GlPolygonOffset
```

#### func (*GlPolygonOffset) API

```go
func (c *GlPolygonOffset) API() gfxapi.ID
```

#### func (*GlPolygonOffset) AddRead

```go
func (a *GlPolygonOffset) AddRead(rng memory.Range, id binary.ID) *GlPolygonOffset
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlPolygonOffset pointer is returned so that calls can be chained.

#### func (*GlPolygonOffset) AddWrite

```go
func (a *GlPolygonOffset) AddWrite(rng memory.Range, id binary.ID) *GlPolygonOffset
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlPolygonOffset pointer is returned so that calls can be chained.

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
func (ϟa *GlPolygonOffset) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlPolygonOffset) Observations

```go
func (a *GlPolygonOffset) Observations() *atom.Observations
```

#### func (*GlPolygonOffset) Replay

```go
func (ϟa *GlPolygonOffset) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlPolygonOffset) String

```go
func (a *GlPolygonOffset) String() string
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
func (c *GlPopGroupMarkerEXT) API() gfxapi.ID
```

#### func (*GlPopGroupMarkerEXT) AddRead

```go
func (a *GlPopGroupMarkerEXT) AddRead(rng memory.Range, id binary.ID) *GlPopGroupMarkerEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlPopGroupMarkerEXT pointer is returned so that calls can be
chained.

#### func (*GlPopGroupMarkerEXT) AddWrite

```go
func (a *GlPopGroupMarkerEXT) AddWrite(rng memory.Range, id binary.ID) *GlPopGroupMarkerEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlPopGroupMarkerEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlPopGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlPopGroupMarkerEXT) Observations

```go
func (a *GlPopGroupMarkerEXT) Observations() *atom.Observations
```

#### func (*GlPopGroupMarkerEXT) Replay

```go
func (ϟa *GlPopGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlPopGroupMarkerEXT) String

```go
func (a *GlPopGroupMarkerEXT) String() string
```

#### type GlProgramBinaryOES

```go
type GlProgramBinaryOES struct {
	binary.Generate

	Program      ProgramId
	BinaryFormat uint32
	Binary       Voidᵖ
	BinarySize   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlProgramBinaryOES
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlProgramBinaryOES

```go
func NewGlProgramBinaryOES(Program ProgramId, Binary_format uint32, Binary memory.Pointer, Binary_size int32) *GlProgramBinaryOES
```

#### func (*GlProgramBinaryOES) API

```go
func (c *GlProgramBinaryOES) API() gfxapi.ID
```

#### func (*GlProgramBinaryOES) AddRead

```go
func (a *GlProgramBinaryOES) AddRead(rng memory.Range, id binary.ID) *GlProgramBinaryOES
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlProgramBinaryOES pointer is returned so that calls can be
chained.

#### func (*GlProgramBinaryOES) AddWrite

```go
func (a *GlProgramBinaryOES) AddWrite(rng memory.Range, id binary.ID) *GlProgramBinaryOES
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlProgramBinaryOES pointer is returned so that calls can be
chained.

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
func (ϟa *GlProgramBinaryOES) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlProgramBinaryOES) Observations

```go
func (a *GlProgramBinaryOES) Observations() *atom.Observations
```

#### func (*GlProgramBinaryOES) Replay

```go
func (ϟa *GlProgramBinaryOES) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlProgramBinaryOES) String

```go
func (a *GlProgramBinaryOES) String() string
```

#### type GlPushGroupMarkerEXT

```go
type GlPushGroupMarkerEXT struct {
	binary.Generate

	Length int32
	Marker Charᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlPushGroupMarkerEXT
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlPushGroupMarkerEXT

```go
func NewGlPushGroupMarkerEXT(Length int32, Marker memory.Pointer) *GlPushGroupMarkerEXT
```

#### func (*GlPushGroupMarkerEXT) API

```go
func (c *GlPushGroupMarkerEXT) API() gfxapi.ID
```

#### func (*GlPushGroupMarkerEXT) AddRead

```go
func (a *GlPushGroupMarkerEXT) AddRead(rng memory.Range, id binary.ID) *GlPushGroupMarkerEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlPushGroupMarkerEXT pointer is returned so that calls can be
chained.

#### func (*GlPushGroupMarkerEXT) AddWrite

```go
func (a *GlPushGroupMarkerEXT) AddWrite(rng memory.Range, id binary.ID) *GlPushGroupMarkerEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlPushGroupMarkerEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlPushGroupMarkerEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlPushGroupMarkerEXT) Observations

```go
func (a *GlPushGroupMarkerEXT) Observations() *atom.Observations
```

#### func (*GlPushGroupMarkerEXT) Replay

```go
func (ϟa *GlPushGroupMarkerEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlPushGroupMarkerEXT) String

```go
func (a *GlPushGroupMarkerEXT) String() string
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
func NewGlQueryCounterEXT(Query QueryId, Target QueryTarget) *GlQueryCounterEXT
```

#### func (*GlQueryCounterEXT) API

```go
func (c *GlQueryCounterEXT) API() gfxapi.ID
```

#### func (*GlQueryCounterEXT) AddRead

```go
func (a *GlQueryCounterEXT) AddRead(rng memory.Range, id binary.ID) *GlQueryCounterEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlQueryCounterEXT pointer is returned so that calls can be chained.

#### func (*GlQueryCounterEXT) AddWrite

```go
func (a *GlQueryCounterEXT) AddWrite(rng memory.Range, id binary.ID) *GlQueryCounterEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlQueryCounterEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlQueryCounterEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlQueryCounterEXT) Observations

```go
func (a *GlQueryCounterEXT) Observations() *atom.Observations
```

#### func (*GlQueryCounterEXT) Replay

```go
func (ϟa *GlQueryCounterEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlQueryCounterEXT) String

```go
func (a *GlQueryCounterEXT) String() string
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
	Data   Voidᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlReadPixels
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlReadPixels

```go
func NewGlReadPixels(X int32, Y int32, Width int32, Height int32, Format BaseTexelFormat, Type TexelType, Data memory.Pointer) *GlReadPixels
```

#### func (*GlReadPixels) API

```go
func (c *GlReadPixels) API() gfxapi.ID
```

#### func (*GlReadPixels) AddRead

```go
func (a *GlReadPixels) AddRead(rng memory.Range, id binary.ID) *GlReadPixels
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlReadPixels pointer is returned so that calls can be chained.

#### func (*GlReadPixels) AddWrite

```go
func (a *GlReadPixels) AddWrite(rng memory.Range, id binary.ID) *GlReadPixels
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlReadPixels pointer is returned so that calls can be chained.

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
func (ϟa *GlReadPixels) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlReadPixels) Observations

```go
func (a *GlReadPixels) Observations() *atom.Observations
```

#### func (*GlReadPixels) Replay

```go
func (ϟa *GlReadPixels) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlReadPixels) String

```go
func (a *GlReadPixels) String() string
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
func (c *GlReleaseShaderCompiler) API() gfxapi.ID
```

#### func (*GlReleaseShaderCompiler) AddRead

```go
func (a *GlReleaseShaderCompiler) AddRead(rng memory.Range, id binary.ID) *GlReleaseShaderCompiler
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlReleaseShaderCompiler pointer is returned so that calls can be
chained.

#### func (*GlReleaseShaderCompiler) AddWrite

```go
func (a *GlReleaseShaderCompiler) AddWrite(rng memory.Range, id binary.ID) *GlReleaseShaderCompiler
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlReleaseShaderCompiler pointer is returned so that calls can be
chained.

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
func (ϟa *GlReleaseShaderCompiler) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlReleaseShaderCompiler) Observations

```go
func (a *GlReleaseShaderCompiler) Observations() *atom.Observations
```

#### func (*GlReleaseShaderCompiler) Replay

```go
func (ϟa *GlReleaseShaderCompiler) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlReleaseShaderCompiler) String

```go
func (a *GlReleaseShaderCompiler) String() string
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
func NewGlRenderbufferStorage(Target RenderbufferTarget, Format RenderbufferFormat, Width int32, Height int32) *GlRenderbufferStorage
```

#### func (*GlRenderbufferStorage) API

```go
func (c *GlRenderbufferStorage) API() gfxapi.ID
```

#### func (*GlRenderbufferStorage) AddRead

```go
func (a *GlRenderbufferStorage) AddRead(rng memory.Range, id binary.ID) *GlRenderbufferStorage
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlRenderbufferStorage pointer is returned so that calls can be
chained.

#### func (*GlRenderbufferStorage) AddWrite

```go
func (a *GlRenderbufferStorage) AddWrite(rng memory.Range, id binary.ID) *GlRenderbufferStorage
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlRenderbufferStorage pointer is returned so that calls can be
chained.

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
func (ϟa *GlRenderbufferStorage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlRenderbufferStorage) Observations

```go
func (a *GlRenderbufferStorage) Observations() *atom.Observations
```

#### func (*GlRenderbufferStorage) Replay

```go
func (ϟa *GlRenderbufferStorage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlRenderbufferStorage) String

```go
func (a *GlRenderbufferStorage) String() string
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
func NewGlRenderbufferStorageMultisample(Target RenderbufferTarget, Samples int32, Format RenderbufferFormat, Width int32, Height int32) *GlRenderbufferStorageMultisample
```

#### func (*GlRenderbufferStorageMultisample) API

```go
func (c *GlRenderbufferStorageMultisample) API() gfxapi.ID
```

#### func (*GlRenderbufferStorageMultisample) AddRead

```go
func (a *GlRenderbufferStorageMultisample) AddRead(rng memory.Range, id binary.ID) *GlRenderbufferStorageMultisample
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlRenderbufferStorageMultisample pointer is returned so that calls
can be chained.

#### func (*GlRenderbufferStorageMultisample) AddWrite

```go
func (a *GlRenderbufferStorageMultisample) AddWrite(rng memory.Range, id binary.ID) *GlRenderbufferStorageMultisample
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlRenderbufferStorageMultisample pointer is returned so that calls
can be chained.

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
func (ϟa *GlRenderbufferStorageMultisample) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlRenderbufferStorageMultisample) Observations

```go
func (a *GlRenderbufferStorageMultisample) Observations() *atom.Observations
```

#### func (*GlRenderbufferStorageMultisample) Replay

```go
func (ϟa *GlRenderbufferStorageMultisample) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlRenderbufferStorageMultisample) String

```go
func (a *GlRenderbufferStorageMultisample) String() string
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
func NewGlSampleCoverage(Value float32, Invert bool) *GlSampleCoverage
```

#### func (*GlSampleCoverage) API

```go
func (c *GlSampleCoverage) API() gfxapi.ID
```

#### func (*GlSampleCoverage) AddRead

```go
func (a *GlSampleCoverage) AddRead(rng memory.Range, id binary.ID) *GlSampleCoverage
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlSampleCoverage pointer is returned so that calls can be chained.

#### func (*GlSampleCoverage) AddWrite

```go
func (a *GlSampleCoverage) AddWrite(rng memory.Range, id binary.ID) *GlSampleCoverage
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlSampleCoverage pointer is returned so that calls can be chained.

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
func (ϟa *GlSampleCoverage) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlSampleCoverage) Observations

```go
func (a *GlSampleCoverage) Observations() *atom.Observations
```

#### func (*GlSampleCoverage) Replay

```go
func (ϟa *GlSampleCoverage) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlSampleCoverage) String

```go
func (a *GlSampleCoverage) String() string
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
func NewGlScissor(X int32, Y int32, Width int32, Height int32) *GlScissor
```

#### func (*GlScissor) API

```go
func (c *GlScissor) API() gfxapi.ID
```

#### func (*GlScissor) AddRead

```go
func (a *GlScissor) AddRead(rng memory.Range, id binary.ID) *GlScissor
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlScissor pointer is returned so that calls can be chained.

#### func (*GlScissor) AddWrite

```go
func (a *GlScissor) AddWrite(rng memory.Range, id binary.ID) *GlScissor
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlScissor pointer is returned so that calls can be chained.

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
func (ϟa *GlScissor) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlScissor) Observations

```go
func (a *GlScissor) Observations() *atom.Observations
```

#### func (*GlScissor) Replay

```go
func (ϟa *GlScissor) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlScissor) String

```go
func (a *GlScissor) String() string
```

#### type GlShaderBinary

```go
type GlShaderBinary struct {
	binary.Generate

	Count        int32
	Shaders      ShaderIdᶜᵖ
	BinaryFormat uint32
	Binary       Voidᶜᵖ
	BinarySize   int32
}
```

//////////////////////////////////////////////////////////////////////////////
GlShaderBinary
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlShaderBinary

```go
func NewGlShaderBinary(Count int32, Shaders memory.Pointer, Binary_format uint32, Binary memory.Pointer, Binary_size int32) *GlShaderBinary
```

#### func (*GlShaderBinary) API

```go
func (c *GlShaderBinary) API() gfxapi.ID
```

#### func (*GlShaderBinary) AddRead

```go
func (a *GlShaderBinary) AddRead(rng memory.Range, id binary.ID) *GlShaderBinary
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlShaderBinary pointer is returned so that calls can be chained.

#### func (*GlShaderBinary) AddWrite

```go
func (a *GlShaderBinary) AddWrite(rng memory.Range, id binary.ID) *GlShaderBinary
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlShaderBinary pointer is returned so that calls can be chained.

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
func (ϟa *GlShaderBinary) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlShaderBinary) Observations

```go
func (a *GlShaderBinary) Observations() *atom.Observations
```

#### func (*GlShaderBinary) Replay

```go
func (ϟa *GlShaderBinary) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlShaderBinary) String

```go
func (a *GlShaderBinary) String() string
```

#### type GlShaderSource

```go
type GlShaderSource struct {
	binary.Generate

	Shader ShaderId
	Count  int32
	Source Charᶜᵖᶜᵖ
	Length S32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlShaderSource
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlShaderSource

```go
func NewGlShaderSource(Shader ShaderId, Count int32, Source memory.Pointer, Length memory.Pointer) *GlShaderSource
```

#### func (*GlShaderSource) API

```go
func (c *GlShaderSource) API() gfxapi.ID
```

#### func (*GlShaderSource) AddRead

```go
func (a *GlShaderSource) AddRead(rng memory.Range, id binary.ID) *GlShaderSource
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlShaderSource pointer is returned so that calls can be chained.

#### func (*GlShaderSource) AddWrite

```go
func (a *GlShaderSource) AddWrite(rng memory.Range, id binary.ID) *GlShaderSource
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlShaderSource pointer is returned so that calls can be chained.

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
func (ϟa *GlShaderSource) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlShaderSource) Observations

```go
func (a *GlShaderSource) Observations() *atom.Observations
```

#### func (*GlShaderSource) Replay

```go
func (ϟa *GlShaderSource) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlShaderSource) String

```go
func (a *GlShaderSource) String() string
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
func NewGlStartTilingQCOM(X int32, Y int32, Width int32, Height int32, PreserveMask TilePreserveMaskQCOM) *GlStartTilingQCOM
```

#### func (*GlStartTilingQCOM) API

```go
func (c *GlStartTilingQCOM) API() gfxapi.ID
```

#### func (*GlStartTilingQCOM) AddRead

```go
func (a *GlStartTilingQCOM) AddRead(rng memory.Range, id binary.ID) *GlStartTilingQCOM
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlStartTilingQCOM pointer is returned so that calls can be chained.

#### func (*GlStartTilingQCOM) AddWrite

```go
func (a *GlStartTilingQCOM) AddWrite(rng memory.Range, id binary.ID) *GlStartTilingQCOM
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlStartTilingQCOM pointer is returned so that calls can be chained.

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
func (ϟa *GlStartTilingQCOM) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlStartTilingQCOM) Observations

```go
func (a *GlStartTilingQCOM) Observations() *atom.Observations
```

#### func (*GlStartTilingQCOM) Replay

```go
func (ϟa *GlStartTilingQCOM) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlStartTilingQCOM) String

```go
func (a *GlStartTilingQCOM) String() string
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
func NewGlStencilFuncSeparate(Face FaceMode, Function TestFunction, Reference_value int32, Mask int32) *GlStencilFuncSeparate
```

#### func (*GlStencilFuncSeparate) API

```go
func (c *GlStencilFuncSeparate) API() gfxapi.ID
```

#### func (*GlStencilFuncSeparate) AddRead

```go
func (a *GlStencilFuncSeparate) AddRead(rng memory.Range, id binary.ID) *GlStencilFuncSeparate
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlStencilFuncSeparate pointer is returned so that calls can be
chained.

#### func (*GlStencilFuncSeparate) AddWrite

```go
func (a *GlStencilFuncSeparate) AddWrite(rng memory.Range, id binary.ID) *GlStencilFuncSeparate
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlStencilFuncSeparate pointer is returned so that calls can be
chained.

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
func (ϟa *GlStencilFuncSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlStencilFuncSeparate) Observations

```go
func (a *GlStencilFuncSeparate) Observations() *atom.Observations
```

#### func (*GlStencilFuncSeparate) Replay

```go
func (ϟa *GlStencilFuncSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlStencilFuncSeparate) String

```go
func (a *GlStencilFuncSeparate) String() string
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
func NewGlStencilMask(Mask uint32) *GlStencilMask
```

#### func (*GlStencilMask) API

```go
func (c *GlStencilMask) API() gfxapi.ID
```

#### func (*GlStencilMask) AddRead

```go
func (a *GlStencilMask) AddRead(rng memory.Range, id binary.ID) *GlStencilMask
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlStencilMask pointer is returned so that calls can be chained.

#### func (*GlStencilMask) AddWrite

```go
func (a *GlStencilMask) AddWrite(rng memory.Range, id binary.ID) *GlStencilMask
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlStencilMask pointer is returned so that calls can be chained.

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
func (ϟa *GlStencilMask) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlStencilMask) Observations

```go
func (a *GlStencilMask) Observations() *atom.Observations
```

#### func (*GlStencilMask) Replay

```go
func (ϟa *GlStencilMask) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlStencilMask) String

```go
func (a *GlStencilMask) String() string
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
func NewGlStencilMaskSeparate(Face FaceMode, Mask uint32) *GlStencilMaskSeparate
```

#### func (*GlStencilMaskSeparate) API

```go
func (c *GlStencilMaskSeparate) API() gfxapi.ID
```

#### func (*GlStencilMaskSeparate) AddRead

```go
func (a *GlStencilMaskSeparate) AddRead(rng memory.Range, id binary.ID) *GlStencilMaskSeparate
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlStencilMaskSeparate pointer is returned so that calls can be
chained.

#### func (*GlStencilMaskSeparate) AddWrite

```go
func (a *GlStencilMaskSeparate) AddWrite(rng memory.Range, id binary.ID) *GlStencilMaskSeparate
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlStencilMaskSeparate pointer is returned so that calls can be
chained.

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
func (ϟa *GlStencilMaskSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlStencilMaskSeparate) Observations

```go
func (a *GlStencilMaskSeparate) Observations() *atom.Observations
```

#### func (*GlStencilMaskSeparate) Replay

```go
func (ϟa *GlStencilMaskSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlStencilMaskSeparate) String

```go
func (a *GlStencilMaskSeparate) String() string
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
func NewGlStencilOpSeparate(Face FaceMode, Stencil_fail StencilAction, Stencil_pass_depth_fail StencilAction, Stencil_pass_depth_pass StencilAction) *GlStencilOpSeparate
```

#### func (*GlStencilOpSeparate) API

```go
func (c *GlStencilOpSeparate) API() gfxapi.ID
```

#### func (*GlStencilOpSeparate) AddRead

```go
func (a *GlStencilOpSeparate) AddRead(rng memory.Range, id binary.ID) *GlStencilOpSeparate
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlStencilOpSeparate pointer is returned so that calls can be
chained.

#### func (*GlStencilOpSeparate) AddWrite

```go
func (a *GlStencilOpSeparate) AddWrite(rng memory.Range, id binary.ID) *GlStencilOpSeparate
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlStencilOpSeparate pointer is returned so that calls can be
chained.

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
func (ϟa *GlStencilOpSeparate) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlStencilOpSeparate) Observations

```go
func (a *GlStencilOpSeparate) Observations() *atom.Observations
```

#### func (*GlStencilOpSeparate) Replay

```go
func (ϟa *GlStencilOpSeparate) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlStencilOpSeparate) String

```go
func (a *GlStencilOpSeparate) String() string
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
func NewGlTexImage2D(Target TextureImageTarget, Level int32, Internal_format TexelFormat, Width int32, Height int32, Border int32, Format TexelFormat, Type TexelType, Data memory.Pointer) *GlTexImage2D
```

#### func (*GlTexImage2D) API

```go
func (c *GlTexImage2D) API() gfxapi.ID
```

#### func (*GlTexImage2D) AddRead

```go
func (a *GlTexImage2D) AddRead(rng memory.Range, id binary.ID) *GlTexImage2D
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTexImage2D pointer is returned so that calls can be chained.

#### func (*GlTexImage2D) AddWrite

```go
func (a *GlTexImage2D) AddWrite(rng memory.Range, id binary.ID) *GlTexImage2D
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTexImage2D pointer is returned so that calls can be chained.

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
func (ϟa *GlTexImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTexImage2D) Observations

```go
func (a *GlTexImage2D) Observations() *atom.Observations
```

#### func (*GlTexImage2D) Replay

```go
func (ϟa *GlTexImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTexImage2D) String

```go
func (a *GlTexImage2D) String() string
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
func NewGlTexParameterf(Target TextureTarget, Parameter TextureParameter, Value float32) *GlTexParameterf
```

#### func (*GlTexParameterf) API

```go
func (c *GlTexParameterf) API() gfxapi.ID
```

#### func (*GlTexParameterf) AddRead

```go
func (a *GlTexParameterf) AddRead(rng memory.Range, id binary.ID) *GlTexParameterf
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTexParameterf pointer is returned so that calls can be chained.

#### func (*GlTexParameterf) AddWrite

```go
func (a *GlTexParameterf) AddWrite(rng memory.Range, id binary.ID) *GlTexParameterf
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTexParameterf pointer is returned so that calls can be chained.

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
func (ϟa *GlTexParameterf) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTexParameterf) Observations

```go
func (a *GlTexParameterf) Observations() *atom.Observations
```

#### func (*GlTexParameterf) Replay

```go
func (ϟa *GlTexParameterf) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTexParameterf) String

```go
func (a *GlTexParameterf) String() string
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
func NewGlTexParameteri(Target TextureTarget, Parameter TextureParameter, Value int32) *GlTexParameteri
```

#### func (*GlTexParameteri) API

```go
func (c *GlTexParameteri) API() gfxapi.ID
```

#### func (*GlTexParameteri) AddRead

```go
func (a *GlTexParameteri) AddRead(rng memory.Range, id binary.ID) *GlTexParameteri
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTexParameteri pointer is returned so that calls can be chained.

#### func (*GlTexParameteri) AddWrite

```go
func (a *GlTexParameteri) AddWrite(rng memory.Range, id binary.ID) *GlTexParameteri
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTexParameteri pointer is returned so that calls can be chained.

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
func (ϟa *GlTexParameteri) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTexParameteri) Observations

```go
func (a *GlTexParameteri) Observations() *atom.Observations
```

#### func (*GlTexParameteri) Replay

```go
func (ϟa *GlTexParameteri) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTexParameteri) String

```go
func (a *GlTexParameteri) String() string
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
func NewGlTexStorage1DEXT(Target TextureTarget, Levels int32, Format TexelFormat, Width int32) *GlTexStorage1DEXT
```

#### func (*GlTexStorage1DEXT) API

```go
func (c *GlTexStorage1DEXT) API() gfxapi.ID
```

#### func (*GlTexStorage1DEXT) AddRead

```go
func (a *GlTexStorage1DEXT) AddRead(rng memory.Range, id binary.ID) *GlTexStorage1DEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTexStorage1DEXT pointer is returned so that calls can be chained.

#### func (*GlTexStorage1DEXT) AddWrite

```go
func (a *GlTexStorage1DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTexStorage1DEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTexStorage1DEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlTexStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTexStorage1DEXT) Observations

```go
func (a *GlTexStorage1DEXT) Observations() *atom.Observations
```

#### func (*GlTexStorage1DEXT) Replay

```go
func (ϟa *GlTexStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTexStorage1DEXT) String

```go
func (a *GlTexStorage1DEXT) String() string
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
func NewGlTexStorage2DEXT(Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32) *GlTexStorage2DEXT
```

#### func (*GlTexStorage2DEXT) API

```go
func (c *GlTexStorage2DEXT) API() gfxapi.ID
```

#### func (*GlTexStorage2DEXT) AddRead

```go
func (a *GlTexStorage2DEXT) AddRead(rng memory.Range, id binary.ID) *GlTexStorage2DEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTexStorage2DEXT pointer is returned so that calls can be chained.

#### func (*GlTexStorage2DEXT) AddWrite

```go
func (a *GlTexStorage2DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTexStorage2DEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTexStorage2DEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlTexStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTexStorage2DEXT) Observations

```go
func (a *GlTexStorage2DEXT) Observations() *atom.Observations
```

#### func (*GlTexStorage2DEXT) Replay

```go
func (ϟa *GlTexStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTexStorage2DEXT) String

```go
func (a *GlTexStorage2DEXT) String() string
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
func NewGlTexStorage3DEXT(Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32, Depth int32) *GlTexStorage3DEXT
```

#### func (*GlTexStorage3DEXT) API

```go
func (c *GlTexStorage3DEXT) API() gfxapi.ID
```

#### func (*GlTexStorage3DEXT) AddRead

```go
func (a *GlTexStorage3DEXT) AddRead(rng memory.Range, id binary.ID) *GlTexStorage3DEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTexStorage3DEXT pointer is returned so that calls can be chained.

#### func (*GlTexStorage3DEXT) AddWrite

```go
func (a *GlTexStorage3DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTexStorage3DEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTexStorage3DEXT pointer is returned so that calls can be chained.

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
func (ϟa *GlTexStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTexStorage3DEXT) Observations

```go
func (a *GlTexStorage3DEXT) Observations() *atom.Observations
```

#### func (*GlTexStorage3DEXT) Replay

```go
func (ϟa *GlTexStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTexStorage3DEXT) String

```go
func (a *GlTexStorage3DEXT) String() string
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
func NewGlTexSubImage2D(Target TextureImageTarget, Level int32, Xoffset int32, Yoffset int32, Width int32, Height int32, Format TexelFormat, Type TexelType, Data memory.Pointer) *GlTexSubImage2D
```

#### func (*GlTexSubImage2D) API

```go
func (c *GlTexSubImage2D) API() gfxapi.ID
```

#### func (*GlTexSubImage2D) AddRead

```go
func (a *GlTexSubImage2D) AddRead(rng memory.Range, id binary.ID) *GlTexSubImage2D
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTexSubImage2D pointer is returned so that calls can be chained.

#### func (*GlTexSubImage2D) AddWrite

```go
func (a *GlTexSubImage2D) AddWrite(rng memory.Range, id binary.ID) *GlTexSubImage2D
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTexSubImage2D pointer is returned so that calls can be chained.

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
func (ϟa *GlTexSubImage2D) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTexSubImage2D) Observations

```go
func (a *GlTexSubImage2D) Observations() *atom.Observations
```

#### func (*GlTexSubImage2D) Replay

```go
func (ϟa *GlTexSubImage2D) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTexSubImage2D) String

```go
func (a *GlTexSubImage2D) String() string
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
func NewGlTextureStorage1DEXT(Texture TextureId, Target TextureTarget, Levels int32, Format TexelFormat, Width int32) *GlTextureStorage1DEXT
```

#### func (*GlTextureStorage1DEXT) API

```go
func (c *GlTextureStorage1DEXT) API() gfxapi.ID
```

#### func (*GlTextureStorage1DEXT) AddRead

```go
func (a *GlTextureStorage1DEXT) AddRead(rng memory.Range, id binary.ID) *GlTextureStorage1DEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTextureStorage1DEXT pointer is returned so that calls can be
chained.

#### func (*GlTextureStorage1DEXT) AddWrite

```go
func (a *GlTextureStorage1DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTextureStorage1DEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTextureStorage1DEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlTextureStorage1DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTextureStorage1DEXT) Observations

```go
func (a *GlTextureStorage1DEXT) Observations() *atom.Observations
```

#### func (*GlTextureStorage1DEXT) Replay

```go
func (ϟa *GlTextureStorage1DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTextureStorage1DEXT) String

```go
func (a *GlTextureStorage1DEXT) String() string
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
func NewGlTextureStorage2DEXT(Texture TextureId, Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32) *GlTextureStorage2DEXT
```

#### func (*GlTextureStorage2DEXT) API

```go
func (c *GlTextureStorage2DEXT) API() gfxapi.ID
```

#### func (*GlTextureStorage2DEXT) AddRead

```go
func (a *GlTextureStorage2DEXT) AddRead(rng memory.Range, id binary.ID) *GlTextureStorage2DEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTextureStorage2DEXT pointer is returned so that calls can be
chained.

#### func (*GlTextureStorage2DEXT) AddWrite

```go
func (a *GlTextureStorage2DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTextureStorage2DEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTextureStorage2DEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlTextureStorage2DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTextureStorage2DEXT) Observations

```go
func (a *GlTextureStorage2DEXT) Observations() *atom.Observations
```

#### func (*GlTextureStorage2DEXT) Replay

```go
func (ϟa *GlTextureStorage2DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTextureStorage2DEXT) String

```go
func (a *GlTextureStorage2DEXT) String() string
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
func NewGlTextureStorage3DEXT(Texture TextureId, Target TextureTarget, Levels int32, Format TexelFormat, Width int32, Height int32, Depth int32) *GlTextureStorage3DEXT
```

#### func (*GlTextureStorage3DEXT) API

```go
func (c *GlTextureStorage3DEXT) API() gfxapi.ID
```

#### func (*GlTextureStorage3DEXT) AddRead

```go
func (a *GlTextureStorage3DEXT) AddRead(rng memory.Range, id binary.ID) *GlTextureStorage3DEXT
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlTextureStorage3DEXT pointer is returned so that calls can be
chained.

#### func (*GlTextureStorage3DEXT) AddWrite

```go
func (a *GlTextureStorage3DEXT) AddWrite(rng memory.Range, id binary.ID) *GlTextureStorage3DEXT
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlTextureStorage3DEXT pointer is returned so that calls can be
chained.

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
func (ϟa *GlTextureStorage3DEXT) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlTextureStorage3DEXT) Observations

```go
func (a *GlTextureStorage3DEXT) Observations() *atom.Observations
```

#### func (*GlTextureStorage3DEXT) Replay

```go
func (ϟa *GlTextureStorage3DEXT) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlTextureStorage3DEXT) String

```go
func (a *GlTextureStorage3DEXT) String() string
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
func NewGlUniform1f(Location UniformLocation, Value float32) *GlUniform1f
```

#### func (*GlUniform1f) API

```go
func (c *GlUniform1f) API() gfxapi.ID
```

#### func (*GlUniform1f) AddRead

```go
func (a *GlUniform1f) AddRead(rng memory.Range, id binary.ID) *GlUniform1f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform1f pointer is returned so that calls can be chained.

#### func (*GlUniform1f) AddWrite

```go
func (a *GlUniform1f) AddWrite(rng memory.Range, id binary.ID) *GlUniform1f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform1f pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform1f) Observations

```go
func (a *GlUniform1f) Observations() *atom.Observations
```

#### func (*GlUniform1f) Replay

```go
func (ϟa *GlUniform1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform1f) String

```go
func (a *GlUniform1f) String() string
```

#### type GlUniform1fv

```go
type GlUniform1fv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform1fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform1fv

```go
func NewGlUniform1fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform1fv
```

#### func (*GlUniform1fv) API

```go
func (c *GlUniform1fv) API() gfxapi.ID
```

#### func (*GlUniform1fv) AddRead

```go
func (a *GlUniform1fv) AddRead(rng memory.Range, id binary.ID) *GlUniform1fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform1fv pointer is returned so that calls can be chained.

#### func (*GlUniform1fv) AddWrite

```go
func (a *GlUniform1fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform1fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform1fv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform1fv) Observations

```go
func (a *GlUniform1fv) Observations() *atom.Observations
```

#### func (*GlUniform1fv) Replay

```go
func (ϟa *GlUniform1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform1fv) String

```go
func (a *GlUniform1fv) String() string
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
func NewGlUniform1i(Location UniformLocation, Value int32) *GlUniform1i
```

#### func (*GlUniform1i) API

```go
func (c *GlUniform1i) API() gfxapi.ID
```

#### func (*GlUniform1i) AddRead

```go
func (a *GlUniform1i) AddRead(rng memory.Range, id binary.ID) *GlUniform1i
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform1i pointer is returned so that calls can be chained.

#### func (*GlUniform1i) AddWrite

```go
func (a *GlUniform1i) AddWrite(rng memory.Range, id binary.ID) *GlUniform1i
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform1i pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform1i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform1i) Observations

```go
func (a *GlUniform1i) Observations() *atom.Observations
```

#### func (*GlUniform1i) Replay

```go
func (ϟa *GlUniform1i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform1i) String

```go
func (a *GlUniform1i) String() string
```

#### type GlUniform1iv

```go
type GlUniform1iv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform1iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform1iv

```go
func NewGlUniform1iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform1iv
```

#### func (*GlUniform1iv) API

```go
func (c *GlUniform1iv) API() gfxapi.ID
```

#### func (*GlUniform1iv) AddRead

```go
func (a *GlUniform1iv) AddRead(rng memory.Range, id binary.ID) *GlUniform1iv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform1iv pointer is returned so that calls can be chained.

#### func (*GlUniform1iv) AddWrite

```go
func (a *GlUniform1iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform1iv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform1iv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform1iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform1iv) Observations

```go
func (a *GlUniform1iv) Observations() *atom.Observations
```

#### func (*GlUniform1iv) Replay

```go
func (ϟa *GlUniform1iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform1iv) String

```go
func (a *GlUniform1iv) String() string
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
func NewGlUniform2f(Location UniformLocation, Value0 float32, Value1 float32) *GlUniform2f
```

#### func (*GlUniform2f) API

```go
func (c *GlUniform2f) API() gfxapi.ID
```

#### func (*GlUniform2f) AddRead

```go
func (a *GlUniform2f) AddRead(rng memory.Range, id binary.ID) *GlUniform2f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform2f pointer is returned so that calls can be chained.

#### func (*GlUniform2f) AddWrite

```go
func (a *GlUniform2f) AddWrite(rng memory.Range, id binary.ID) *GlUniform2f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform2f pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform2f) Observations

```go
func (a *GlUniform2f) Observations() *atom.Observations
```

#### func (*GlUniform2f) Replay

```go
func (ϟa *GlUniform2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform2f) String

```go
func (a *GlUniform2f) String() string
```

#### type GlUniform2fv

```go
type GlUniform2fv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform2fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform2fv

```go
func NewGlUniform2fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform2fv
```

#### func (*GlUniform2fv) API

```go
func (c *GlUniform2fv) API() gfxapi.ID
```

#### func (*GlUniform2fv) AddRead

```go
func (a *GlUniform2fv) AddRead(rng memory.Range, id binary.ID) *GlUniform2fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform2fv pointer is returned so that calls can be chained.

#### func (*GlUniform2fv) AddWrite

```go
func (a *GlUniform2fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform2fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform2fv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform2fv) Observations

```go
func (a *GlUniform2fv) Observations() *atom.Observations
```

#### func (*GlUniform2fv) Replay

```go
func (ϟa *GlUniform2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform2fv) String

```go
func (a *GlUniform2fv) String() string
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
func NewGlUniform2i(Location UniformLocation, Value0 int32, Value1 int32) *GlUniform2i
```

#### func (*GlUniform2i) API

```go
func (c *GlUniform2i) API() gfxapi.ID
```

#### func (*GlUniform2i) AddRead

```go
func (a *GlUniform2i) AddRead(rng memory.Range, id binary.ID) *GlUniform2i
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform2i pointer is returned so that calls can be chained.

#### func (*GlUniform2i) AddWrite

```go
func (a *GlUniform2i) AddWrite(rng memory.Range, id binary.ID) *GlUniform2i
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform2i pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform2i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform2i) Observations

```go
func (a *GlUniform2i) Observations() *atom.Observations
```

#### func (*GlUniform2i) Replay

```go
func (ϟa *GlUniform2i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform2i) String

```go
func (a *GlUniform2i) String() string
```

#### type GlUniform2iv

```go
type GlUniform2iv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform2iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform2iv

```go
func NewGlUniform2iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform2iv
```

#### func (*GlUniform2iv) API

```go
func (c *GlUniform2iv) API() gfxapi.ID
```

#### func (*GlUniform2iv) AddRead

```go
func (a *GlUniform2iv) AddRead(rng memory.Range, id binary.ID) *GlUniform2iv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform2iv pointer is returned so that calls can be chained.

#### func (*GlUniform2iv) AddWrite

```go
func (a *GlUniform2iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform2iv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform2iv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform2iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform2iv) Observations

```go
func (a *GlUniform2iv) Observations() *atom.Observations
```

#### func (*GlUniform2iv) Replay

```go
func (ϟa *GlUniform2iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform2iv) String

```go
func (a *GlUniform2iv) String() string
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
func NewGlUniform3f(Location UniformLocation, Value0 float32, Value1 float32, Value2 float32) *GlUniform3f
```

#### func (*GlUniform3f) API

```go
func (c *GlUniform3f) API() gfxapi.ID
```

#### func (*GlUniform3f) AddRead

```go
func (a *GlUniform3f) AddRead(rng memory.Range, id binary.ID) *GlUniform3f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform3f pointer is returned so that calls can be chained.

#### func (*GlUniform3f) AddWrite

```go
func (a *GlUniform3f) AddWrite(rng memory.Range, id binary.ID) *GlUniform3f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform3f pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform3f) Observations

```go
func (a *GlUniform3f) Observations() *atom.Observations
```

#### func (*GlUniform3f) Replay

```go
func (ϟa *GlUniform3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform3f) String

```go
func (a *GlUniform3f) String() string
```

#### type GlUniform3fv

```go
type GlUniform3fv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform3fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform3fv

```go
func NewGlUniform3fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform3fv
```

#### func (*GlUniform3fv) API

```go
func (c *GlUniform3fv) API() gfxapi.ID
```

#### func (*GlUniform3fv) AddRead

```go
func (a *GlUniform3fv) AddRead(rng memory.Range, id binary.ID) *GlUniform3fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform3fv pointer is returned so that calls can be chained.

#### func (*GlUniform3fv) AddWrite

```go
func (a *GlUniform3fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform3fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform3fv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform3fv) Observations

```go
func (a *GlUniform3fv) Observations() *atom.Observations
```

#### func (*GlUniform3fv) Replay

```go
func (ϟa *GlUniform3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform3fv) String

```go
func (a *GlUniform3fv) String() string
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
func NewGlUniform3i(Location UniformLocation, Value0 int32, Value1 int32, Value2 int32) *GlUniform3i
```

#### func (*GlUniform3i) API

```go
func (c *GlUniform3i) API() gfxapi.ID
```

#### func (*GlUniform3i) AddRead

```go
func (a *GlUniform3i) AddRead(rng memory.Range, id binary.ID) *GlUniform3i
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform3i pointer is returned so that calls can be chained.

#### func (*GlUniform3i) AddWrite

```go
func (a *GlUniform3i) AddWrite(rng memory.Range, id binary.ID) *GlUniform3i
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform3i pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform3i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform3i) Observations

```go
func (a *GlUniform3i) Observations() *atom.Observations
```

#### func (*GlUniform3i) Replay

```go
func (ϟa *GlUniform3i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform3i) String

```go
func (a *GlUniform3i) String() string
```

#### type GlUniform3iv

```go
type GlUniform3iv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform3iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform3iv

```go
func NewGlUniform3iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform3iv
```

#### func (*GlUniform3iv) API

```go
func (c *GlUniform3iv) API() gfxapi.ID
```

#### func (*GlUniform3iv) AddRead

```go
func (a *GlUniform3iv) AddRead(rng memory.Range, id binary.ID) *GlUniform3iv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform3iv pointer is returned so that calls can be chained.

#### func (*GlUniform3iv) AddWrite

```go
func (a *GlUniform3iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform3iv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform3iv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform3iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform3iv) Observations

```go
func (a *GlUniform3iv) Observations() *atom.Observations
```

#### func (*GlUniform3iv) Replay

```go
func (ϟa *GlUniform3iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform3iv) String

```go
func (a *GlUniform3iv) String() string
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
func NewGlUniform4f(Location UniformLocation, Value0 float32, Value1 float32, Value2 float32, Value3 float32) *GlUniform4f
```

#### func (*GlUniform4f) API

```go
func (c *GlUniform4f) API() gfxapi.ID
```

#### func (*GlUniform4f) AddRead

```go
func (a *GlUniform4f) AddRead(rng memory.Range, id binary.ID) *GlUniform4f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform4f pointer is returned so that calls can be chained.

#### func (*GlUniform4f) AddWrite

```go
func (a *GlUniform4f) AddWrite(rng memory.Range, id binary.ID) *GlUniform4f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform4f pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform4f) Observations

```go
func (a *GlUniform4f) Observations() *atom.Observations
```

#### func (*GlUniform4f) Replay

```go
func (ϟa *GlUniform4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform4f) String

```go
func (a *GlUniform4f) String() string
```

#### type GlUniform4fv

```go
type GlUniform4fv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform4fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform4fv

```go
func NewGlUniform4fv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform4fv
```

#### func (*GlUniform4fv) API

```go
func (c *GlUniform4fv) API() gfxapi.ID
```

#### func (*GlUniform4fv) AddRead

```go
func (a *GlUniform4fv) AddRead(rng memory.Range, id binary.ID) *GlUniform4fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform4fv pointer is returned so that calls can be chained.

#### func (*GlUniform4fv) AddWrite

```go
func (a *GlUniform4fv) AddWrite(rng memory.Range, id binary.ID) *GlUniform4fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform4fv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform4fv) Observations

```go
func (a *GlUniform4fv) Observations() *atom.Observations
```

#### func (*GlUniform4fv) Replay

```go
func (ϟa *GlUniform4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform4fv) String

```go
func (a *GlUniform4fv) String() string
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
func NewGlUniform4i(Location UniformLocation, Value0 int32, Value1 int32, Value2 int32, Value3 int32) *GlUniform4i
```

#### func (*GlUniform4i) API

```go
func (c *GlUniform4i) API() gfxapi.ID
```

#### func (*GlUniform4i) AddRead

```go
func (a *GlUniform4i) AddRead(rng memory.Range, id binary.ID) *GlUniform4i
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform4i pointer is returned so that calls can be chained.

#### func (*GlUniform4i) AddWrite

```go
func (a *GlUniform4i) AddWrite(rng memory.Range, id binary.ID) *GlUniform4i
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform4i pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform4i) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform4i) Observations

```go
func (a *GlUniform4i) Observations() *atom.Observations
```

#### func (*GlUniform4i) Replay

```go
func (ϟa *GlUniform4i) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform4i) String

```go
func (a *GlUniform4i) String() string
```

#### type GlUniform4iv

```go
type GlUniform4iv struct {
	binary.Generate

	Location UniformLocation
	Count    int32
	Values   S32ᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniform4iv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniform4iv

```go
func NewGlUniform4iv(Location UniformLocation, Count int32, Values memory.Pointer) *GlUniform4iv
```

#### func (*GlUniform4iv) API

```go
func (c *GlUniform4iv) API() gfxapi.ID
```

#### func (*GlUniform4iv) AddRead

```go
func (a *GlUniform4iv) AddRead(rng memory.Range, id binary.ID) *GlUniform4iv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniform4iv pointer is returned so that calls can be chained.

#### func (*GlUniform4iv) AddWrite

```go
func (a *GlUniform4iv) AddWrite(rng memory.Range, id binary.ID) *GlUniform4iv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniform4iv pointer is returned so that calls can be chained.

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
func (ϟa *GlUniform4iv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniform4iv) Observations

```go
func (a *GlUniform4iv) Observations() *atom.Observations
```

#### func (*GlUniform4iv) Replay

```go
func (ϟa *GlUniform4iv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniform4iv) String

```go
func (a *GlUniform4iv) String() string
```

#### type GlUniformBlockBinding

```go
type GlUniformBlockBinding struct {
	binary.Generate

	Program             ProgramId
	UniformBlockIndex   uint32
	UniformBlockBinding uint32
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniformBlockBinding
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniformBlockBinding

```go
func NewGlUniformBlockBinding(Program ProgramId, Uniform_block_index uint32, Uniform_block_binding uint32) *GlUniformBlockBinding
```

#### func (*GlUniformBlockBinding) API

```go
func (c *GlUniformBlockBinding) API() gfxapi.ID
```

#### func (*GlUniformBlockBinding) AddRead

```go
func (a *GlUniformBlockBinding) AddRead(rng memory.Range, id binary.ID) *GlUniformBlockBinding
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniformBlockBinding pointer is returned so that calls can be
chained.

#### func (*GlUniformBlockBinding) AddWrite

```go
func (a *GlUniformBlockBinding) AddWrite(rng memory.Range, id binary.ID) *GlUniformBlockBinding
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniformBlockBinding pointer is returned so that calls can be
chained.

#### func (*GlUniformBlockBinding) Class

```go
func (*GlUniformBlockBinding) Class() binary.Class
```

#### func (*GlUniformBlockBinding) Flags

```go
func (c *GlUniformBlockBinding) Flags() atom.Flags
```

#### func (*GlUniformBlockBinding) Mutate

```go
func (ϟa *GlUniformBlockBinding) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniformBlockBinding) Observations

```go
func (a *GlUniformBlockBinding) Observations() *atom.Observations
```

#### func (*GlUniformBlockBinding) Replay

```go
func (ϟa *GlUniformBlockBinding) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniformBlockBinding) String

```go
func (a *GlUniformBlockBinding) String() string
```

#### type GlUniformMatrix2fv

```go
type GlUniformMatrix2fv struct {
	binary.Generate

	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniformMatrix2fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniformMatrix2fv

```go
func NewGlUniformMatrix2fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix2fv
```

#### func (*GlUniformMatrix2fv) API

```go
func (c *GlUniformMatrix2fv) API() gfxapi.ID
```

#### func (*GlUniformMatrix2fv) AddRead

```go
func (a *GlUniformMatrix2fv) AddRead(rng memory.Range, id binary.ID) *GlUniformMatrix2fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniformMatrix2fv pointer is returned so that calls can be
chained.

#### func (*GlUniformMatrix2fv) AddWrite

```go
func (a *GlUniformMatrix2fv) AddWrite(rng memory.Range, id binary.ID) *GlUniformMatrix2fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniformMatrix2fv pointer is returned so that calls can be
chained.

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
func (ϟa *GlUniformMatrix2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniformMatrix2fv) Observations

```go
func (a *GlUniformMatrix2fv) Observations() *atom.Observations
```

#### func (*GlUniformMatrix2fv) Replay

```go
func (ϟa *GlUniformMatrix2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniformMatrix2fv) String

```go
func (a *GlUniformMatrix2fv) String() string
```

#### type GlUniformMatrix3fv

```go
type GlUniformMatrix3fv struct {
	binary.Generate

	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniformMatrix3fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniformMatrix3fv

```go
func NewGlUniformMatrix3fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix3fv
```

#### func (*GlUniformMatrix3fv) API

```go
func (c *GlUniformMatrix3fv) API() gfxapi.ID
```

#### func (*GlUniformMatrix3fv) AddRead

```go
func (a *GlUniformMatrix3fv) AddRead(rng memory.Range, id binary.ID) *GlUniformMatrix3fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniformMatrix3fv pointer is returned so that calls can be
chained.

#### func (*GlUniformMatrix3fv) AddWrite

```go
func (a *GlUniformMatrix3fv) AddWrite(rng memory.Range, id binary.ID) *GlUniformMatrix3fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniformMatrix3fv pointer is returned so that calls can be
chained.

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
func (ϟa *GlUniformMatrix3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniformMatrix3fv) Observations

```go
func (a *GlUniformMatrix3fv) Observations() *atom.Observations
```

#### func (*GlUniformMatrix3fv) Replay

```go
func (ϟa *GlUniformMatrix3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniformMatrix3fv) String

```go
func (a *GlUniformMatrix3fv) String() string
```

#### type GlUniformMatrix4fv

```go
type GlUniformMatrix4fv struct {
	binary.Generate

	Location  UniformLocation
	Count     int32
	Transpose bool
	Values    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlUniformMatrix4fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlUniformMatrix4fv

```go
func NewGlUniformMatrix4fv(Location UniformLocation, Count int32, Transpose bool, Values memory.Pointer) *GlUniformMatrix4fv
```

#### func (*GlUniformMatrix4fv) API

```go
func (c *GlUniformMatrix4fv) API() gfxapi.ID
```

#### func (*GlUniformMatrix4fv) AddRead

```go
func (a *GlUniformMatrix4fv) AddRead(rng memory.Range, id binary.ID) *GlUniformMatrix4fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUniformMatrix4fv pointer is returned so that calls can be
chained.

#### func (*GlUniformMatrix4fv) AddWrite

```go
func (a *GlUniformMatrix4fv) AddWrite(rng memory.Range, id binary.ID) *GlUniformMatrix4fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUniformMatrix4fv pointer is returned so that calls can be
chained.

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
func (ϟa *GlUniformMatrix4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUniformMatrix4fv) Observations

```go
func (a *GlUniformMatrix4fv) Observations() *atom.Observations
```

#### func (*GlUniformMatrix4fv) Replay

```go
func (ϟa *GlUniformMatrix4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUniformMatrix4fv) String

```go
func (a *GlUniformMatrix4fv) String() string
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
func NewGlUnmapBuffer(Target BufferTarget) *GlUnmapBuffer
```

#### func (*GlUnmapBuffer) API

```go
func (c *GlUnmapBuffer) API() gfxapi.ID
```

#### func (*GlUnmapBuffer) AddRead

```go
func (a *GlUnmapBuffer) AddRead(rng memory.Range, id binary.ID) *GlUnmapBuffer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUnmapBuffer pointer is returned so that calls can be chained.

#### func (*GlUnmapBuffer) AddWrite

```go
func (a *GlUnmapBuffer) AddWrite(rng memory.Range, id binary.ID) *GlUnmapBuffer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUnmapBuffer pointer is returned so that calls can be chained.

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
func (ϟa *GlUnmapBuffer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUnmapBuffer) Observations

```go
func (a *GlUnmapBuffer) Observations() *atom.Observations
```

#### func (*GlUnmapBuffer) Replay

```go
func (ϟa *GlUnmapBuffer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUnmapBuffer) String

```go
func (a *GlUnmapBuffer) String() string
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
func NewGlUseProgram(Program ProgramId) *GlUseProgram
```

#### func (*GlUseProgram) API

```go
func (c *GlUseProgram) API() gfxapi.ID
```

#### func (*GlUseProgram) AddRead

```go
func (a *GlUseProgram) AddRead(rng memory.Range, id binary.ID) *GlUseProgram
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlUseProgram pointer is returned so that calls can be chained.

#### func (*GlUseProgram) AddWrite

```go
func (a *GlUseProgram) AddWrite(rng memory.Range, id binary.ID) *GlUseProgram
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlUseProgram pointer is returned so that calls can be chained.

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
func (ϟa *GlUseProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlUseProgram) Observations

```go
func (a *GlUseProgram) Observations() *atom.Observations
```

#### func (*GlUseProgram) Replay

```go
func (ϟa *GlUseProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlUseProgram) String

```go
func (a *GlUseProgram) String() string
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
func NewGlValidateProgram(Program ProgramId) *GlValidateProgram
```

#### func (*GlValidateProgram) API

```go
func (c *GlValidateProgram) API() gfxapi.ID
```

#### func (*GlValidateProgram) AddRead

```go
func (a *GlValidateProgram) AddRead(rng memory.Range, id binary.ID) *GlValidateProgram
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlValidateProgram pointer is returned so that calls can be chained.

#### func (*GlValidateProgram) AddWrite

```go
func (a *GlValidateProgram) AddWrite(rng memory.Range, id binary.ID) *GlValidateProgram
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlValidateProgram pointer is returned so that calls can be chained.

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
func (ϟa *GlValidateProgram) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlValidateProgram) Observations

```go
func (a *GlValidateProgram) Observations() *atom.Observations
```

#### func (*GlValidateProgram) Replay

```go
func (ϟa *GlValidateProgram) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlValidateProgram) String

```go
func (a *GlValidateProgram) String() string
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
func NewGlVertexAttrib1f(Location AttributeLocation, Value0 float32) *GlVertexAttrib1f
```

#### func (*GlVertexAttrib1f) API

```go
func (c *GlVertexAttrib1f) API() gfxapi.ID
```

#### func (*GlVertexAttrib1f) AddRead

```go
func (a *GlVertexAttrib1f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib1f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib1f pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib1f) AddWrite

```go
func (a *GlVertexAttrib1f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib1f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib1f pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib1f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib1f) Observations

```go
func (a *GlVertexAttrib1f) Observations() *atom.Observations
```

#### func (*GlVertexAttrib1f) Replay

```go
func (ϟa *GlVertexAttrib1f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib1f) String

```go
func (a *GlVertexAttrib1f) String() string
```

#### type GlVertexAttrib1fv

```go
type GlVertexAttrib1fv struct {
	binary.Generate

	Location AttributeLocation
	Value    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib1fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib1fv

```go
func NewGlVertexAttrib1fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib1fv
```

#### func (*GlVertexAttrib1fv) API

```go
func (c *GlVertexAttrib1fv) API() gfxapi.ID
```

#### func (*GlVertexAttrib1fv) AddRead

```go
func (a *GlVertexAttrib1fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib1fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib1fv pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib1fv) AddWrite

```go
func (a *GlVertexAttrib1fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib1fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib1fv pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib1fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib1fv) Observations

```go
func (a *GlVertexAttrib1fv) Observations() *atom.Observations
```

#### func (*GlVertexAttrib1fv) Replay

```go
func (ϟa *GlVertexAttrib1fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib1fv) String

```go
func (a *GlVertexAttrib1fv) String() string
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
func NewGlVertexAttrib2f(Location AttributeLocation, Value0 float32, Value1 float32) *GlVertexAttrib2f
```

#### func (*GlVertexAttrib2f) API

```go
func (c *GlVertexAttrib2f) API() gfxapi.ID
```

#### func (*GlVertexAttrib2f) AddRead

```go
func (a *GlVertexAttrib2f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib2f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib2f pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib2f) AddWrite

```go
func (a *GlVertexAttrib2f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib2f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib2f pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib2f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib2f) Observations

```go
func (a *GlVertexAttrib2f) Observations() *atom.Observations
```

#### func (*GlVertexAttrib2f) Replay

```go
func (ϟa *GlVertexAttrib2f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib2f) String

```go
func (a *GlVertexAttrib2f) String() string
```

#### type GlVertexAttrib2fv

```go
type GlVertexAttrib2fv struct {
	binary.Generate

	Location AttributeLocation
	Value    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib2fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib2fv

```go
func NewGlVertexAttrib2fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib2fv
```

#### func (*GlVertexAttrib2fv) API

```go
func (c *GlVertexAttrib2fv) API() gfxapi.ID
```

#### func (*GlVertexAttrib2fv) AddRead

```go
func (a *GlVertexAttrib2fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib2fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib2fv pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib2fv) AddWrite

```go
func (a *GlVertexAttrib2fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib2fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib2fv pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib2fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib2fv) Observations

```go
func (a *GlVertexAttrib2fv) Observations() *atom.Observations
```

#### func (*GlVertexAttrib2fv) Replay

```go
func (ϟa *GlVertexAttrib2fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib2fv) String

```go
func (a *GlVertexAttrib2fv) String() string
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
func NewGlVertexAttrib3f(Location AttributeLocation, Value0 float32, Value1 float32, Value2 float32) *GlVertexAttrib3f
```

#### func (*GlVertexAttrib3f) API

```go
func (c *GlVertexAttrib3f) API() gfxapi.ID
```

#### func (*GlVertexAttrib3f) AddRead

```go
func (a *GlVertexAttrib3f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib3f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib3f pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib3f) AddWrite

```go
func (a *GlVertexAttrib3f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib3f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib3f pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib3f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib3f) Observations

```go
func (a *GlVertexAttrib3f) Observations() *atom.Observations
```

#### func (*GlVertexAttrib3f) Replay

```go
func (ϟa *GlVertexAttrib3f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib3f) String

```go
func (a *GlVertexAttrib3f) String() string
```

#### type GlVertexAttrib3fv

```go
type GlVertexAttrib3fv struct {
	binary.Generate

	Location AttributeLocation
	Value    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib3fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib3fv

```go
func NewGlVertexAttrib3fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib3fv
```

#### func (*GlVertexAttrib3fv) API

```go
func (c *GlVertexAttrib3fv) API() gfxapi.ID
```

#### func (*GlVertexAttrib3fv) AddRead

```go
func (a *GlVertexAttrib3fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib3fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib3fv pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib3fv) AddWrite

```go
func (a *GlVertexAttrib3fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib3fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib3fv pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib3fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib3fv) Observations

```go
func (a *GlVertexAttrib3fv) Observations() *atom.Observations
```

#### func (*GlVertexAttrib3fv) Replay

```go
func (ϟa *GlVertexAttrib3fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib3fv) String

```go
func (a *GlVertexAttrib3fv) String() string
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
func NewGlVertexAttrib4f(Location AttributeLocation, Value0 float32, Value1 float32, Value2 float32, Value3 float32) *GlVertexAttrib4f
```

#### func (*GlVertexAttrib4f) API

```go
func (c *GlVertexAttrib4f) API() gfxapi.ID
```

#### func (*GlVertexAttrib4f) AddRead

```go
func (a *GlVertexAttrib4f) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib4f
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib4f pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib4f) AddWrite

```go
func (a *GlVertexAttrib4f) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib4f
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib4f pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib4f) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib4f) Observations

```go
func (a *GlVertexAttrib4f) Observations() *atom.Observations
```

#### func (*GlVertexAttrib4f) Replay

```go
func (ϟa *GlVertexAttrib4f) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib4f) String

```go
func (a *GlVertexAttrib4f) String() string
```

#### type GlVertexAttrib4fv

```go
type GlVertexAttrib4fv struct {
	binary.Generate

	Location AttributeLocation
	Value    F32ᶜᵖ
}
```

//////////////////////////////////////////////////////////////////////////////
GlVertexAttrib4fv
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlVertexAttrib4fv

```go
func NewGlVertexAttrib4fv(Location AttributeLocation, Value memory.Pointer) *GlVertexAttrib4fv
```

#### func (*GlVertexAttrib4fv) API

```go
func (c *GlVertexAttrib4fv) API() gfxapi.ID
```

#### func (*GlVertexAttrib4fv) AddRead

```go
func (a *GlVertexAttrib4fv) AddRead(rng memory.Range, id binary.ID) *GlVertexAttrib4fv
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttrib4fv pointer is returned so that calls can be chained.

#### func (*GlVertexAttrib4fv) AddWrite

```go
func (a *GlVertexAttrib4fv) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttrib4fv
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttrib4fv pointer is returned so that calls can be chained.

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
func (ϟa *GlVertexAttrib4fv) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttrib4fv) Observations

```go
func (a *GlVertexAttrib4fv) Observations() *atom.Observations
```

#### func (*GlVertexAttrib4fv) Replay

```go
func (ϟa *GlVertexAttrib4fv) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttrib4fv) String

```go
func (a *GlVertexAttrib4fv) String() string
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
func NewGlVertexAttribPointer(Location AttributeLocation, Size int32, Type VertexAttribType, Normalized bool, Stride int32, Data memory.Pointer) *GlVertexAttribPointer
```

#### func (*GlVertexAttribPointer) API

```go
func (c *GlVertexAttribPointer) API() gfxapi.ID
```

#### func (*GlVertexAttribPointer) AddRead

```go
func (a *GlVertexAttribPointer) AddRead(rng memory.Range, id binary.ID) *GlVertexAttribPointer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlVertexAttribPointer pointer is returned so that calls can be
chained.

#### func (*GlVertexAttribPointer) AddWrite

```go
func (a *GlVertexAttribPointer) AddWrite(rng memory.Range, id binary.ID) *GlVertexAttribPointer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlVertexAttribPointer pointer is returned so that calls can be
chained.

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
func (ϟa *GlVertexAttribPointer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlVertexAttribPointer) Observations

```go
func (a *GlVertexAttribPointer) Observations() *atom.Observations
```

#### func (*GlVertexAttribPointer) Replay

```go
func (ϟa *GlVertexAttribPointer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlVertexAttribPointer) String

```go
func (a *GlVertexAttribPointer) String() string
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
func NewGlViewport(X int32, Y int32, Width int32, Height int32) *GlViewport
```

#### func (*GlViewport) API

```go
func (c *GlViewport) API() gfxapi.ID
```

#### func (*GlViewport) AddRead

```go
func (a *GlViewport) AddRead(rng memory.Range, id binary.ID) *GlViewport
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlViewport pointer is returned so that calls can be chained.

#### func (*GlViewport) AddWrite

```go
func (a *GlViewport) AddWrite(rng memory.Range, id binary.ID) *GlViewport
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlViewport pointer is returned so that calls can be chained.

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
func (ϟa *GlViewport) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlViewport) Observations

```go
func (a *GlViewport) Observations() *atom.Observations
```

#### func (*GlViewport) Replay

```go
func (ϟa *GlViewport) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlViewport) String

```go
func (a *GlViewport) String() string
```

#### type GlWaitSync

```go
type GlWaitSync struct {
	binary.Generate

	Sync      SyncObject
	SyncFlags SyncFlags
	Timeout   uint64
}
```

//////////////////////////////////////////////////////////////////////////////
GlWaitSync
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlWaitSync

```go
func NewGlWaitSync(Sync SyncObject, SyncFlags SyncFlags, Timeout uint64) *GlWaitSync
```

#### func (*GlWaitSync) API

```go
func (c *GlWaitSync) API() gfxapi.ID
```

#### func (*GlWaitSync) AddRead

```go
func (a *GlWaitSync) AddRead(rng memory.Range, id binary.ID) *GlWaitSync
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlWaitSync pointer is returned so that calls can be chained.

#### func (*GlWaitSync) AddWrite

```go
func (a *GlWaitSync) AddWrite(rng memory.Range, id binary.ID) *GlWaitSync
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlWaitSync pointer is returned so that calls can be chained.

#### func (*GlWaitSync) Class

```go
func (*GlWaitSync) Class() binary.Class
```

#### func (*GlWaitSync) Flags

```go
func (c *GlWaitSync) Flags() atom.Flags
```

#### func (*GlWaitSync) Mutate

```go
func (ϟa *GlWaitSync) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlWaitSync) Observations

```go
func (a *GlWaitSync) Observations() *atom.Observations
```

#### func (*GlWaitSync) Replay

```go
func (ϟa *GlWaitSync) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*GlWaitSync) String

```go
func (a *GlWaitSync) String() string
```

#### type GlXCreateContext

```go
type GlXCreateContext struct {
	binary.Generate

	Dpy       Voidᵖ
	Vis       Voidᵖ
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
func NewGlXCreateContext(Dpy memory.Pointer, Vis memory.Pointer, ShareList memory.Pointer, Direct bool, Result memory.Pointer) *GlXCreateContext
```

#### func (*GlXCreateContext) API

```go
func (c *GlXCreateContext) API() gfxapi.ID
```

#### func (*GlXCreateContext) AddRead

```go
func (a *GlXCreateContext) AddRead(rng memory.Range, id binary.ID) *GlXCreateContext
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlXCreateContext pointer is returned so that calls can be chained.

#### func (*GlXCreateContext) AddWrite

```go
func (a *GlXCreateContext) AddWrite(rng memory.Range, id binary.ID) *GlXCreateContext
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlXCreateContext pointer is returned so that calls can be chained.

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
func (ϟa *GlXCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlXCreateContext) Observations

```go
func (a *GlXCreateContext) Observations() *atom.Observations
```

#### func (*GlXCreateContext) Replay

```go
func (ω *GlXCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*GlXCreateContext) String

```go
func (a *GlXCreateContext) String() string
```

#### type GlXCreateNewContext

```go
type GlXCreateNewContext struct {
	binary.Generate

	Display  Voidᵖ
	Fbconfig Voidᵖ
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
func NewGlXCreateNewContext(Display memory.Pointer, Fbconfig memory.Pointer, Type uint32, Shared memory.Pointer, Direct bool, Result memory.Pointer) *GlXCreateNewContext
```

#### func (*GlXCreateNewContext) API

```go
func (c *GlXCreateNewContext) API() gfxapi.ID
```

#### func (*GlXCreateNewContext) AddRead

```go
func (a *GlXCreateNewContext) AddRead(rng memory.Range, id binary.ID) *GlXCreateNewContext
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlXCreateNewContext pointer is returned so that calls can be
chained.

#### func (*GlXCreateNewContext) AddWrite

```go
func (a *GlXCreateNewContext) AddWrite(rng memory.Range, id binary.ID) *GlXCreateNewContext
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlXCreateNewContext pointer is returned so that calls can be
chained.

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
func (ϟa *GlXCreateNewContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlXCreateNewContext) Observations

```go
func (a *GlXCreateNewContext) Observations() *atom.Observations
```

#### func (*GlXCreateNewContext) Replay

```go
func (ω *GlXCreateNewContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*GlXCreateNewContext) String

```go
func (a *GlXCreateNewContext) String() string
```

#### type GlXMakeContextCurrent

```go
type GlXMakeContextCurrent struct {
	binary.Generate

	Display Voidᵖ
	Draw    GLXDrawable
	Read    GLXDrawable
	Ctx     GLXContext
	Result  Bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlXMakeContextCurrent
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXMakeContextCurrent

```go
func NewGlXMakeContextCurrent(Display memory.Pointer, Draw memory.Pointer, Read memory.Pointer, Ctx memory.Pointer, Result Bool) *GlXMakeContextCurrent
```

#### func (*GlXMakeContextCurrent) API

```go
func (c *GlXMakeContextCurrent) API() gfxapi.ID
```

#### func (*GlXMakeContextCurrent) AddRead

```go
func (a *GlXMakeContextCurrent) AddRead(rng memory.Range, id binary.ID) *GlXMakeContextCurrent
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlXMakeContextCurrent pointer is returned so that calls can be
chained.

#### func (*GlXMakeContextCurrent) AddWrite

```go
func (a *GlXMakeContextCurrent) AddWrite(rng memory.Range, id binary.ID) *GlXMakeContextCurrent
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlXMakeContextCurrent pointer is returned so that calls can be
chained.

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
func (ϟa *GlXMakeContextCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlXMakeContextCurrent) Observations

```go
func (a *GlXMakeContextCurrent) Observations() *atom.Observations
```

#### func (*GlXMakeContextCurrent) Replay

```go
func (ω *GlXMakeContextCurrent) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*GlXMakeContextCurrent) String

```go
func (a *GlXMakeContextCurrent) String() string
```

#### type GlXMakeCurrent

```go
type GlXMakeCurrent struct {
	binary.Generate

	Display  Voidᵖ
	Drawable GLXDrawable
	Ctx      GLXContext
	Result   Bool
}
```

//////////////////////////////////////////////////////////////////////////////
GlXMakeCurrent
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXMakeCurrent

```go
func NewGlXMakeCurrent(Display memory.Pointer, Drawable memory.Pointer, Ctx memory.Pointer, Result Bool) *GlXMakeCurrent
```

#### func (*GlXMakeCurrent) API

```go
func (c *GlXMakeCurrent) API() gfxapi.ID
```

#### func (*GlXMakeCurrent) AddRead

```go
func (a *GlXMakeCurrent) AddRead(rng memory.Range, id binary.ID) *GlXMakeCurrent
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlXMakeCurrent pointer is returned so that calls can be chained.

#### func (*GlXMakeCurrent) AddWrite

```go
func (a *GlXMakeCurrent) AddWrite(rng memory.Range, id binary.ID) *GlXMakeCurrent
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlXMakeCurrent pointer is returned so that calls can be chained.

#### func (*GlXMakeCurrent) Class

```go
func (*GlXMakeCurrent) Class() binary.Class
```

#### func (*GlXMakeCurrent) Flags

```go
func (c *GlXMakeCurrent) Flags() atom.Flags
```

#### func (*GlXMakeCurrent) Mutate

```go
func (ϟa *GlXMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlXMakeCurrent) Observations

```go
func (a *GlXMakeCurrent) Observations() *atom.Observations
```

#### func (*GlXMakeCurrent) String

```go
func (a *GlXMakeCurrent) String() string
```

#### type GlXQueryDrawable

```go
type GlXQueryDrawable struct {
	binary.Generate

	Display   Voidᵖ
	Draw      GLXDrawable
	Attribute int64
	Value     Intᵖ
	Result    int64
}
```

//////////////////////////////////////////////////////////////////////////////
GlXQueryDrawable
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXQueryDrawable

```go
func NewGlXQueryDrawable(Display memory.Pointer, Draw memory.Pointer, Attribute int64, Value memory.Pointer, Result int64) *GlXQueryDrawable
```

#### func (*GlXQueryDrawable) API

```go
func (c *GlXQueryDrawable) API() gfxapi.ID
```

#### func (*GlXQueryDrawable) AddRead

```go
func (a *GlXQueryDrawable) AddRead(rng memory.Range, id binary.ID) *GlXQueryDrawable
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlXQueryDrawable pointer is returned so that calls can be chained.

#### func (*GlXQueryDrawable) AddWrite

```go
func (a *GlXQueryDrawable) AddWrite(rng memory.Range, id binary.ID) *GlXQueryDrawable
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlXQueryDrawable pointer is returned so that calls can be chained.

#### func (*GlXQueryDrawable) Class

```go
func (*GlXQueryDrawable) Class() binary.Class
```

#### func (*GlXQueryDrawable) Flags

```go
func (c *GlXQueryDrawable) Flags() atom.Flags
```

#### func (*GlXQueryDrawable) Mutate

```go
func (ϟa *GlXQueryDrawable) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlXQueryDrawable) Observations

```go
func (a *GlXQueryDrawable) Observations() *atom.Observations
```

#### func (*GlXQueryDrawable) String

```go
func (a *GlXQueryDrawable) String() string
```

#### type GlXSwapBuffers

```go
type GlXSwapBuffers struct {
	binary.Generate

	Display  Voidᵖ
	Drawable GLXDrawable
}
```

//////////////////////////////////////////////////////////////////////////////
GlXSwapBuffers
//////////////////////////////////////////////////////////////////////////////

#### func  NewGlXSwapBuffers

```go
func NewGlXSwapBuffers(Display memory.Pointer, Drawable memory.Pointer) *GlXSwapBuffers
```

#### func (*GlXSwapBuffers) API

```go
func (c *GlXSwapBuffers) API() gfxapi.ID
```

#### func (*GlXSwapBuffers) AddRead

```go
func (a *GlXSwapBuffers) AddRead(rng memory.Range, id binary.ID) *GlXSwapBuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The GlXSwapBuffers pointer is returned so that calls can be chained.

#### func (*GlXSwapBuffers) AddWrite

```go
func (a *GlXSwapBuffers) AddWrite(rng memory.Range, id binary.ID) *GlXSwapBuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The GlXSwapBuffers pointer is returned so that calls can be chained.

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
func (ϟa *GlXSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*GlXSwapBuffers) Observations

```go
func (a *GlXSwapBuffers) Observations() *atom.Observations
```

#### func (*GlXSwapBuffers) String

```go
func (a *GlXSwapBuffers) String() string
```

#### type HDC

```go
type HDC struct {
	binary.Generate
	memory.Pointer
}
```

HDC is a pointer to a void element.

#### func  NewHDC

```go
func NewHDC(addr uint64) HDC
```
NewHDC returns a HDC that points to addr in the application pool.

#### func (*HDC) Class

```go
func (*HDC) Class() binary.Class
```

#### func (HDC) ElementSize

```go
func (p HDC) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that HDC points to.

#### func (HDC) OnRead

```go
func (p HDC) OnRead(ϟs *gfxapi.State) HDC
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (HDC) OnWrite

```go
func (p HDC) OnWrite(ϟs *gfxapi.State) HDC
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (HDC) Slice

```go
func (p HDC) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type HGLRC

```go
type HGLRC struct {
	binary.Generate
	memory.Pointer
}
```

HGLRC is a pointer to a void element.

#### func  NewHGLRC

```go
func NewHGLRC(addr uint64) HGLRC
```
NewHGLRC returns a HGLRC that points to addr in the application pool.

#### func (*HGLRC) Class

```go
func (*HGLRC) Class() binary.Class
```

#### func (HGLRC) ElementSize

```go
func (p HGLRC) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that HGLRC points to.

#### func (HGLRC) OnRead

```go
func (p HGLRC) OnRead(ϟs *gfxapi.State) HGLRC
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (HGLRC) OnWrite

```go
func (p HGLRC) OnWrite(ϟs *gfxapi.State) HGLRC
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (HGLRC) Slice

```go
func (p HGLRC) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type HGLRCːContextʳᵐ

```go
type HGLRCːContextʳᵐ map[HGLRC](*Context)
```


#### func (HGLRCːContextʳᵐ) Contains

```go
func (m HGLRCːContextʳᵐ) Contains(key HGLRC) bool
```

#### func (HGLRCːContextʳᵐ) Delete

```go
func (m HGLRCːContextʳᵐ) Delete(key HGLRC)
```

#### func (HGLRCːContextʳᵐ) Get

```go
func (m HGLRCːContextʳᵐ) Get(key HGLRC) *Context
```

#### func (HGLRCːContextʳᵐ) Range

```go
func (m HGLRCːContextʳᵐ) Range() [](*Context)
```

#### type HintMode

```go
type HintMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum HintMode
//////////////////////////////////////////////////////////////////////////////

#### func (*HintMode) Parse

```go
func (v *HintMode) Parse(s string) error
```

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

#### func (*HintTarget) Parse

```go
func (v *HintTarget) Parse(s string) error
```

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
	Data      U8ˢ
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
type ImageOES struct {
	binary.Generate
	memory.Pointer
}
```

ImageOES is a pointer to a void element.

#### func  NewImageOES

```go
func NewImageOES(addr uint64) ImageOES
```
NewImageOES returns a ImageOES that points to addr in the application pool.

#### func (*ImageOES) Class

```go
func (*ImageOES) Class() binary.Class
```

#### func (ImageOES) ElementSize

```go
func (p ImageOES) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ImageOES points to.

#### func (ImageOES) OnRead

```go
func (p ImageOES) OnRead(ϟs *gfxapi.State) ImageOES
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (ImageOES) OnWrite

```go
func (p ImageOES) OnWrite(ϟs *gfxapi.State) ImageOES
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (ImageOES) Slice

```go
func (p ImageOES) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type ImageTargetRenderbufferStorage

```go
type ImageTargetRenderbufferStorage uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ImageTargetRenderbufferStorage
//////////////////////////////////////////////////////////////////////////////

#### func (*ImageTargetRenderbufferStorage) Parse

```go
func (v *ImageTargetRenderbufferStorage) Parse(s string) error
```

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

#### func (*ImageTargetTexture) Parse

```go
func (v *ImageTargetTexture) Parse(s string) error
```

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

#### func (*ImageTargetTexture_OES_EGL_image) Parse

```go
func (v *ImageTargetTexture_OES_EGL_image) Parse(s string) error
```

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

#### func (*ImageTargetTexture_OES_EGL_image_external) Parse

```go
func (v *ImageTargetTexture_OES_EGL_image_external) Parse(s string) error
```

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

#### func (*ImageTexelFormat) Parse

```go
func (v *ImageTexelFormat) Parse(s string) error
```

#### func (ImageTexelFormat) String

```go
func (v ImageTexelFormat) String() string
```

#### type IndexedBufferTarget

```go
type IndexedBufferTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum IndexedBufferTarget
//////////////////////////////////////////////////////////////////////////////

#### func (*IndexedBufferTarget) Parse

```go
func (v *IndexedBufferTarget) Parse(s string) error
```

#### func (IndexedBufferTarget) String

```go
func (v IndexedBufferTarget) String() string
```

#### type IndicesPointer

```go
type IndicesPointer struct {
	binary.Generate
	memory.Pointer
}
```

IndicesPointer is a pointer to a void element.

#### func  NewIndicesPointer

```go
func NewIndicesPointer(addr uint64) IndicesPointer
```
NewIndicesPointer returns a IndicesPointer that points to addr in the
application pool.

#### func (*IndicesPointer) Class

```go
func (*IndicesPointer) Class() binary.Class
```

#### func (IndicesPointer) ElementSize

```go
func (p IndicesPointer) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that IndicesPointer points
to.

#### func (IndicesPointer) OnRead

```go
func (p IndicesPointer) OnRead(ϟs *gfxapi.State) IndicesPointer
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (IndicesPointer) OnWrite

```go
func (p IndicesPointer) OnWrite(ϟs *gfxapi.State) IndicesPointer
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (IndicesPointer) Slice

```go
func (p IndicesPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type IndicesType

```go
type IndicesType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum IndicesType
//////////////////////////////////////////////////////////////////////////////

#### func (*IndicesType) Parse

```go
func (v *IndicesType) Parse(s string) error
```

#### func (IndicesType) String

```go
func (v IndicesType) String() string
```

#### type Intˢ

```go
type Intˢ struct {
	binary.Generate
	SliceInfo
}
```

Intˢ is a slice of int64.

#### func  AsIntˢ

```go
func AsIntˢ(s Slice, ϟs *gfxapi.State) Intˢ
```
AsIntˢ returns s cast to a Intˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeIntˢ

```go
func MakeIntˢ(count uint64, ϟs *gfxapi.State) Intˢ
```
MakeIntˢ returns a Intˢ backed by a new memory pool.

#### func (*Intˢ) Class

```go
func (*Intˢ) Class() binary.Class
```

#### func (Intˢ) Clone

```go
func (s Intˢ) Clone(ϟs *gfxapi.State) Intˢ
```
Clone returns a copy of the Intˢ in a new memory pool.

#### func (Intˢ) Copy

```go
func (dst Intˢ) Copy(src Intˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Intˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Intˢ) Decoder

```go
func (s Intˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Intˢ) ElementSize

```go
func (s Intˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Intˢ points to.

#### func (Intˢ) Encoder

```go
func (s Intˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Intˢ) Index

```go
func (s Intˢ) Index(i uint64, ϟs *gfxapi.State) Intᵖ
```
Index returns a Intᵖ to the i'th element in this Intˢ.

#### func (Intˢ) OnRead

```go
func (s Intˢ) OnRead(ϟs *gfxapi.State) Intˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Intˢ) OnWrite

```go
func (s Intˢ) OnWrite(ϟs *gfxapi.State) Intˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Intˢ) Range

```go
func (s Intˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Intˢ) Read

```go
func (s Intˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64
```
Read reads and returns all the int64 elements in this Intˢ.

#### func (Intˢ) ResourceID

```go
func (s Intˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Intˢ) Slice

```go
func (s Intˢ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ
```
Slice returns a sub-slice from the Intˢ using start and end indices.

#### func (Intˢ) String

```go
func (s Intˢ) String() string
```
String returns a string description of the Intˢ slice.

#### func (Intˢ) Write

```go
func (s Intˢ) Write(src []int64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Intᵖ

```go
type Intᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Intᵖ is a pointer to a int64 element.

#### func  NewIntᵖ

```go
func NewIntᵖ(addr uint64) Intᵖ
```
NewIntᵖ returns a Intᵖ that points to addr in the application pool.

#### func (*Intᵖ) Class

```go
func (*Intᵖ) Class() binary.Class
```

#### func (Intᵖ) ElementSize

```go
func (p Intᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Intᵖ points to.

#### func (Intᵖ) OnRead

```go
func (p Intᵖ) OnRead(ϟs *gfxapi.State) Intᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Intᵖ) OnWrite

```go
func (p Intᵖ) OnWrite(ϟs *gfxapi.State) Intᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Intᵖ) Read

```go
func (p Intᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64
```
Read reads and returns the int64 element at the pointer.

#### func (Intᵖ) Slice

```go
func (p Intᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Intˢ
```
Slice returns a new Intˢ from the pointer using start and end indices.

#### func (Intᵖ) Write

```go
func (p Intᵖ) Write(value int64, ϟs *gfxapi.State)
```
Write writes value to the int64 element at the pointer.

#### type MapBufferRangeAccess

```go
type MapBufferRangeAccess uint32
```

//////////////////////////////////////////////////////////////////////////////
enum MapBufferRangeAccess
//////////////////////////////////////////////////////////////////////////////

#### func (*MapBufferRangeAccess) Parse

```go
func (v *MapBufferRangeAccess) Parse(s string) error
```

#### func (MapBufferRangeAccess) String

```go
func (v MapBufferRangeAccess) String() string
```

#### type Mat2f

```go
type Mat2f Vec2fː2ᵃ
```


#### func (*Mat2f) Class

```go
func (*Mat2f) Class() binary.Class
```

#### func (Mat2f) String

```go
func (m Mat2f) String() string
```

#### type Mat2fˢ

```go
type Mat2fˢ struct {
	binary.Generate
	SliceInfo
}
```

Mat2fˢ is a slice of Mat2f.

#### func  AsMat2fˢ

```go
func AsMat2fˢ(s Slice, ϟs *gfxapi.State) Mat2fˢ
```
AsMat2fˢ returns s cast to a Mat2fˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeMat2fˢ

```go
func MakeMat2fˢ(count uint64, ϟs *gfxapi.State) Mat2fˢ
```
MakeMat2fˢ returns a Mat2fˢ backed by a new memory pool.

#### func (*Mat2fˢ) Class

```go
func (*Mat2fˢ) Class() binary.Class
```

#### func (Mat2fˢ) Clone

```go
func (s Mat2fˢ) Clone(ϟs *gfxapi.State) Mat2fˢ
```
Clone returns a copy of the Mat2fˢ in a new memory pool.

#### func (Mat2fˢ) Copy

```go
func (dst Mat2fˢ) Copy(src Mat2fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Mat2fˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Mat2fˢ) Decoder

```go
func (s Mat2fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Mat2fˢ) ElementSize

```go
func (s Mat2fˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Mat2fˢ points to.

#### func (Mat2fˢ) Encoder

```go
func (s Mat2fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Mat2fˢ) Index

```go
func (s Mat2fˢ) Index(i uint64, ϟs *gfxapi.State) Mat2fᵖ
```
Index returns a Mat2fᵖ to the i'th element in this Mat2fˢ.

#### func (Mat2fˢ) OnRead

```go
func (s Mat2fˢ) OnRead(ϟs *gfxapi.State) Mat2fˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Mat2fˢ) OnWrite

```go
func (s Mat2fˢ) OnWrite(ϟs *gfxapi.State) Mat2fˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Mat2fˢ) Range

```go
func (s Mat2fˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Mat2fˢ) Read

```go
func (s Mat2fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Mat2f
```
Read reads and returns all the Mat2f elements in this Mat2fˢ.

#### func (Mat2fˢ) ResourceID

```go
func (s Mat2fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Mat2fˢ) Slice

```go
func (s Mat2fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Mat2fˢ
```
Slice returns a sub-slice from the Mat2fˢ using start and end indices.

#### func (Mat2fˢ) String

```go
func (s Mat2fˢ) String() string
```
String returns a string description of the Mat2fˢ slice.

#### func (Mat2fˢ) Write

```go
func (s Mat2fˢ) Write(src []Mat2f, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Mat2fᵖ

```go
type Mat2fᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Mat2fᵖ is a pointer to a Mat2f element.

#### func  NewMat2fᵖ

```go
func NewMat2fᵖ(addr uint64) Mat2fᵖ
```
NewMat2fᵖ returns a Mat2fᵖ that points to addr in the application pool.

#### func (*Mat2fᵖ) Class

```go
func (*Mat2fᵖ) Class() binary.Class
```

#### func (Mat2fᵖ) ElementSize

```go
func (p Mat2fᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Mat2fᵖ points to.

#### func (Mat2fᵖ) OnRead

```go
func (p Mat2fᵖ) OnRead(ϟs *gfxapi.State) Mat2fᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Mat2fᵖ) OnWrite

```go
func (p Mat2fᵖ) OnWrite(ϟs *gfxapi.State) Mat2fᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Mat2fᵖ) Read

```go
func (p Mat2fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Mat2f
```
Read reads and returns the Mat2f element at the pointer.

#### func (Mat2fᵖ) Slice

```go
func (p Mat2fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Mat2fˢ
```
Slice returns a new Mat2fˢ from the pointer using start and end indices.

#### func (Mat2fᵖ) Write

```go
func (p Mat2fᵖ) Write(value Mat2f, ϟs *gfxapi.State)
```
Write writes value to the Mat2f element at the pointer.

#### type Mat3f

```go
type Mat3f Vec3fː3ᵃ
```


#### func (*Mat3f) Class

```go
func (*Mat3f) Class() binary.Class
```

#### func (Mat3f) String

```go
func (m Mat3f) String() string
```

#### type Mat3fˢ

```go
type Mat3fˢ struct {
	binary.Generate
	SliceInfo
}
```

Mat3fˢ is a slice of Mat3f.

#### func  AsMat3fˢ

```go
func AsMat3fˢ(s Slice, ϟs *gfxapi.State) Mat3fˢ
```
AsMat3fˢ returns s cast to a Mat3fˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeMat3fˢ

```go
func MakeMat3fˢ(count uint64, ϟs *gfxapi.State) Mat3fˢ
```
MakeMat3fˢ returns a Mat3fˢ backed by a new memory pool.

#### func (*Mat3fˢ) Class

```go
func (*Mat3fˢ) Class() binary.Class
```

#### func (Mat3fˢ) Clone

```go
func (s Mat3fˢ) Clone(ϟs *gfxapi.State) Mat3fˢ
```
Clone returns a copy of the Mat3fˢ in a new memory pool.

#### func (Mat3fˢ) Copy

```go
func (dst Mat3fˢ) Copy(src Mat3fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Mat3fˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Mat3fˢ) Decoder

```go
func (s Mat3fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Mat3fˢ) ElementSize

```go
func (s Mat3fˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Mat3fˢ points to.

#### func (Mat3fˢ) Encoder

```go
func (s Mat3fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Mat3fˢ) Index

```go
func (s Mat3fˢ) Index(i uint64, ϟs *gfxapi.State) Mat3fᵖ
```
Index returns a Mat3fᵖ to the i'th element in this Mat3fˢ.

#### func (Mat3fˢ) OnRead

```go
func (s Mat3fˢ) OnRead(ϟs *gfxapi.State) Mat3fˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Mat3fˢ) OnWrite

```go
func (s Mat3fˢ) OnWrite(ϟs *gfxapi.State) Mat3fˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Mat3fˢ) Range

```go
func (s Mat3fˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Mat3fˢ) Read

```go
func (s Mat3fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Mat3f
```
Read reads and returns all the Mat3f elements in this Mat3fˢ.

#### func (Mat3fˢ) ResourceID

```go
func (s Mat3fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Mat3fˢ) Slice

```go
func (s Mat3fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Mat3fˢ
```
Slice returns a sub-slice from the Mat3fˢ using start and end indices.

#### func (Mat3fˢ) String

```go
func (s Mat3fˢ) String() string
```
String returns a string description of the Mat3fˢ slice.

#### func (Mat3fˢ) Write

```go
func (s Mat3fˢ) Write(src []Mat3f, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Mat3fᵖ

```go
type Mat3fᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Mat3fᵖ is a pointer to a Mat3f element.

#### func  NewMat3fᵖ

```go
func NewMat3fᵖ(addr uint64) Mat3fᵖ
```
NewMat3fᵖ returns a Mat3fᵖ that points to addr in the application pool.

#### func (*Mat3fᵖ) Class

```go
func (*Mat3fᵖ) Class() binary.Class
```

#### func (Mat3fᵖ) ElementSize

```go
func (p Mat3fᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Mat3fᵖ points to.

#### func (Mat3fᵖ) OnRead

```go
func (p Mat3fᵖ) OnRead(ϟs *gfxapi.State) Mat3fᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Mat3fᵖ) OnWrite

```go
func (p Mat3fᵖ) OnWrite(ϟs *gfxapi.State) Mat3fᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Mat3fᵖ) Read

```go
func (p Mat3fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Mat3f
```
Read reads and returns the Mat3f element at the pointer.

#### func (Mat3fᵖ) Slice

```go
func (p Mat3fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Mat3fˢ
```
Slice returns a new Mat3fˢ from the pointer using start and end indices.

#### func (Mat3fᵖ) Write

```go
func (p Mat3fᵖ) Write(value Mat3f, ϟs *gfxapi.State)
```
Write writes value to the Mat3f element at the pointer.

#### type Mat4f

```go
type Mat4f Vec4fː4ᵃ
```


#### func (*Mat4f) Class

```go
func (*Mat4f) Class() binary.Class
```

#### func (Mat4f) String

```go
func (m Mat4f) String() string
```

#### type Mat4fˢ

```go
type Mat4fˢ struct {
	binary.Generate
	SliceInfo
}
```

Mat4fˢ is a slice of Mat4f.

#### func  AsMat4fˢ

```go
func AsMat4fˢ(s Slice, ϟs *gfxapi.State) Mat4fˢ
```
AsMat4fˢ returns s cast to a Mat4fˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeMat4fˢ

```go
func MakeMat4fˢ(count uint64, ϟs *gfxapi.State) Mat4fˢ
```
MakeMat4fˢ returns a Mat4fˢ backed by a new memory pool.

#### func (*Mat4fˢ) Class

```go
func (*Mat4fˢ) Class() binary.Class
```

#### func (Mat4fˢ) Clone

```go
func (s Mat4fˢ) Clone(ϟs *gfxapi.State) Mat4fˢ
```
Clone returns a copy of the Mat4fˢ in a new memory pool.

#### func (Mat4fˢ) Copy

```go
func (dst Mat4fˢ) Copy(src Mat4fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Mat4fˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Mat4fˢ) Decoder

```go
func (s Mat4fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Mat4fˢ) ElementSize

```go
func (s Mat4fˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Mat4fˢ points to.

#### func (Mat4fˢ) Encoder

```go
func (s Mat4fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Mat4fˢ) Index

```go
func (s Mat4fˢ) Index(i uint64, ϟs *gfxapi.State) Mat4fᵖ
```
Index returns a Mat4fᵖ to the i'th element in this Mat4fˢ.

#### func (Mat4fˢ) OnRead

```go
func (s Mat4fˢ) OnRead(ϟs *gfxapi.State) Mat4fˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Mat4fˢ) OnWrite

```go
func (s Mat4fˢ) OnWrite(ϟs *gfxapi.State) Mat4fˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Mat4fˢ) Range

```go
func (s Mat4fˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Mat4fˢ) Read

```go
func (s Mat4fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Mat4f
```
Read reads and returns all the Mat4f elements in this Mat4fˢ.

#### func (Mat4fˢ) ResourceID

```go
func (s Mat4fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Mat4fˢ) Slice

```go
func (s Mat4fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Mat4fˢ
```
Slice returns a sub-slice from the Mat4fˢ using start and end indices.

#### func (Mat4fˢ) String

```go
func (s Mat4fˢ) String() string
```
String returns a string description of the Mat4fˢ slice.

#### func (Mat4fˢ) Write

```go
func (s Mat4fˢ) Write(src []Mat4f, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Mat4fᵖ

```go
type Mat4fᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Mat4fᵖ is a pointer to a Mat4f element.

#### func  NewMat4fᵖ

```go
func NewMat4fᵖ(addr uint64) Mat4fᵖ
```
NewMat4fᵖ returns a Mat4fᵖ that points to addr in the application pool.

#### func (*Mat4fᵖ) Class

```go
func (*Mat4fᵖ) Class() binary.Class
```

#### func (Mat4fᵖ) ElementSize

```go
func (p Mat4fᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Mat4fᵖ points to.

#### func (Mat4fᵖ) OnRead

```go
func (p Mat4fᵖ) OnRead(ϟs *gfxapi.State) Mat4fᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Mat4fᵖ) OnWrite

```go
func (p Mat4fᵖ) OnWrite(ϟs *gfxapi.State) Mat4fᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Mat4fᵖ) Read

```go
func (p Mat4fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Mat4f
```
Read reads and returns the Mat4f element at the pointer.

#### func (Mat4fᵖ) Slice

```go
func (p Mat4fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Mat4fˢ
```
Slice returns a new Mat4fˢ from the pointer using start and end indices.

#### func (Mat4fᵖ) Write

```go
func (p Mat4fᵖ) Write(value Mat4f, ϟs *gfxapi.State)
```
Write writes value to the Mat4f element at the pointer.

#### type Objects

```go
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

#### func (*PixelStoreParameter) Parse

```go
func (v *PixelStoreParameter) Parse(s string) error
```

#### func (PixelStoreParameter) String

```go
func (v PixelStoreParameter) String() string
```

#### type PixelStoreParameterːs32ᵐ

```go
type PixelStoreParameterːs32ᵐ map[PixelStoreParameter]int32
```


#### func (PixelStoreParameterːs32ᵐ) Contains

```go
func (m PixelStoreParameterːs32ᵐ) Contains(key PixelStoreParameter) bool
```

#### func (PixelStoreParameterːs32ᵐ) Delete

```go
func (m PixelStoreParameterːs32ᵐ) Delete(key PixelStoreParameter)
```

#### func (PixelStoreParameterːs32ᵐ) Get

```go
func (m PixelStoreParameterːs32ᵐ) Get(key PixelStoreParameter) int32
```

#### func (PixelStoreParameterːs32ᵐ) Range

```go
func (m PixelStoreParameterːs32ᵐ) Range() []int32
```

#### type PrecisionType

```go
type PrecisionType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum PrecisionType
//////////////////////////////////////////////////////////////////////////////

#### func (*PrecisionType) Parse

```go
func (v *PrecisionType) Parse(s string) error
```

#### func (PrecisionType) String

```go
func (v PrecisionType) String() string
```

#### type Program

```go
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

#### type ProgramId

```go
type ProgramId uint32
```


#### type ProgramIdːProgramʳᵐ

```go
type ProgramIdːProgramʳᵐ map[ProgramId](*Program)
```


#### func (ProgramIdːProgramʳᵐ) Contains

```go
func (m ProgramIdːProgramʳᵐ) Contains(key ProgramId) bool
```

#### func (ProgramIdːProgramʳᵐ) Delete

```go
func (m ProgramIdːProgramʳᵐ) Delete(key ProgramId)
```

#### func (ProgramIdːProgramʳᵐ) Get

```go
func (m ProgramIdːProgramʳᵐ) Get(key ProgramId) *Program
```

#### func (ProgramIdːProgramʳᵐ) Range

```go
func (m ProgramIdːProgramʳᵐ) Range() [](*Program)
```

#### type ProgramParameter

```go
type ProgramParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ProgramParameter
//////////////////////////////////////////////////////////////////////////////

#### func (*ProgramParameter) Parse

```go
func (v *ProgramParameter) Parse(s string) error
```

#### func (ProgramParameter) String

```go
func (v ProgramParameter) String() string
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

#### type QueryId

```go
type QueryId uint32
```


#### type QueryIdːQueryʳᵐ

```go
type QueryIdːQueryʳᵐ map[QueryId](*Query)
```


#### func (QueryIdːQueryʳᵐ) Contains

```go
func (m QueryIdːQueryʳᵐ) Contains(key QueryId) bool
```

#### func (QueryIdːQueryʳᵐ) Delete

```go
func (m QueryIdːQueryʳᵐ) Delete(key QueryId)
```

#### func (QueryIdːQueryʳᵐ) Get

```go
func (m QueryIdːQueryʳᵐ) Get(key QueryId) *Query
```

#### func (QueryIdːQueryʳᵐ) Range

```go
func (m QueryIdːQueryʳᵐ) Range() [](*Query)
```

#### type QueryIdˢ

```go
type QueryIdˢ struct {
	binary.Generate
	SliceInfo
}
```

QueryIdˢ is a slice of QueryId.

#### func  AsQueryIdˢ

```go
func AsQueryIdˢ(s Slice, ϟs *gfxapi.State) QueryIdˢ
```
AsQueryIdˢ returns s cast to a QueryIdˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeQueryIdˢ

```go
func MakeQueryIdˢ(count uint64, ϟs *gfxapi.State) QueryIdˢ
```
MakeQueryIdˢ returns a QueryIdˢ backed by a new memory pool.

#### func (*QueryIdˢ) Class

```go
func (*QueryIdˢ) Class() binary.Class
```

#### func (QueryIdˢ) Clone

```go
func (s QueryIdˢ) Clone(ϟs *gfxapi.State) QueryIdˢ
```
Clone returns a copy of the QueryIdˢ in a new memory pool.

#### func (QueryIdˢ) Copy

```go
func (dst QueryIdˢ) Copy(src QueryIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s QueryIdˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (QueryIdˢ) Decoder

```go
func (s QueryIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (QueryIdˢ) ElementSize

```go
func (s QueryIdˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that QueryIdˢ points to.

#### func (QueryIdˢ) Encoder

```go
func (s QueryIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (QueryIdˢ) Index

```go
func (s QueryIdˢ) Index(i uint64, ϟs *gfxapi.State) QueryIdᵖ
```
Index returns a QueryIdᵖ to the i'th element in this QueryIdˢ.

#### func (QueryIdˢ) OnRead

```go
func (s QueryIdˢ) OnRead(ϟs *gfxapi.State) QueryIdˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (QueryIdˢ) OnWrite

```go
func (s QueryIdˢ) OnWrite(ϟs *gfxapi.State) QueryIdˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (QueryIdˢ) Range

```go
func (s QueryIdˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (QueryIdˢ) Read

```go
func (s QueryIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []QueryId
```
Read reads and returns all the QueryId elements in this QueryIdˢ.

#### func (QueryIdˢ) ResourceID

```go
func (s QueryIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (QueryIdˢ) Slice

```go
func (s QueryIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ
```
Slice returns a sub-slice from the QueryIdˢ using start and end indices.

#### func (QueryIdˢ) String

```go
func (s QueryIdˢ) String() string
```
String returns a string description of the QueryIdˢ slice.

#### func (QueryIdˢ) Write

```go
func (s QueryIdˢ) Write(src []QueryId, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type QueryIdᵖ

```go
type QueryIdᵖ struct {
	binary.Generate
	memory.Pointer
}
```

QueryIdᵖ is a pointer to a QueryId element.

#### func  NewQueryIdᵖ

```go
func NewQueryIdᵖ(addr uint64) QueryIdᵖ
```
NewQueryIdᵖ returns a QueryIdᵖ that points to addr in the application pool.

#### func (*QueryIdᵖ) Class

```go
func (*QueryIdᵖ) Class() binary.Class
```

#### func (QueryIdᵖ) ElementSize

```go
func (p QueryIdᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that QueryIdᵖ points to.

#### func (QueryIdᵖ) OnRead

```go
func (p QueryIdᵖ) OnRead(ϟs *gfxapi.State) QueryIdᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (QueryIdᵖ) OnWrite

```go
func (p QueryIdᵖ) OnWrite(ϟs *gfxapi.State) QueryIdᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (QueryIdᵖ) Read

```go
func (p QueryIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) QueryId
```
Read reads and returns the QueryId element at the pointer.

#### func (QueryIdᵖ) Slice

```go
func (p QueryIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ
```
Slice returns a new QueryIdˢ from the pointer using start and end indices.

#### func (QueryIdᵖ) Write

```go
func (p QueryIdᵖ) Write(value QueryId, ϟs *gfxapi.State)
```
Write writes value to the QueryId element at the pointer.

#### type QueryIdᶜᵖ

```go
type QueryIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

QueryIdᶜᵖ is a pointer to a QueryId element.

#### func  NewQueryIdᶜᵖ

```go
func NewQueryIdᶜᵖ(addr uint64) QueryIdᶜᵖ
```
NewQueryIdᶜᵖ returns a QueryIdᶜᵖ that points to addr in the application pool.

#### func (*QueryIdᶜᵖ) Class

```go
func (*QueryIdᶜᵖ) Class() binary.Class
```

#### func (QueryIdᶜᵖ) ElementSize

```go
func (p QueryIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that QueryIdᶜᵖ points to.

#### func (QueryIdᶜᵖ) OnRead

```go
func (p QueryIdᶜᵖ) OnRead(ϟs *gfxapi.State) QueryIdᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (QueryIdᶜᵖ) OnWrite

```go
func (p QueryIdᶜᵖ) OnWrite(ϟs *gfxapi.State) QueryIdᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (QueryIdᶜᵖ) Read

```go
func (p QueryIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) QueryId
```
Read reads and returns the QueryId element at the pointer.

#### func (QueryIdᶜᵖ) Slice

```go
func (p QueryIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) QueryIdˢ
```
Slice returns a new QueryIdˢ from the pointer using start and end indices.

#### func (QueryIdᶜᵖ) Write

```go
func (p QueryIdᶜᵖ) Write(value QueryId, ϟs *gfxapi.State)
```
Write writes value to the QueryId element at the pointer.

#### type QueryObjectParameter

```go
type QueryObjectParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryObjectParameter
//////////////////////////////////////////////////////////////////////////////

#### func (*QueryObjectParameter) Parse

```go
func (v *QueryObjectParameter) Parse(s string) error
```

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

#### type QueryObjectParameter_GLES_3

```go
type QueryObjectParameter_GLES_3 uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryObjectParameter_GLES_3
//////////////////////////////////////////////////////////////////////////////

#### func (*QueryObjectParameter_GLES_3) Parse

```go
func (v *QueryObjectParameter_GLES_3) Parse(s string) error
```

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

#### func (*QueryParameter) Parse

```go
func (v *QueryParameter) Parse(s string) error
```

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

#### func (*QueryParameter_EXT_disjoint_timer_query) Parse

```go
func (v *QueryParameter_EXT_disjoint_timer_query) Parse(s string) error
```

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

#### func (*QueryParameter_GLES_3) Parse

```go
func (v *QueryParameter_GLES_3) Parse(s string) error
```

#### func (QueryParameter_GLES_3) String

```go
func (v QueryParameter_GLES_3) String() string
```

#### type QueryTarget

```go
type QueryTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum QueryTarget
//////////////////////////////////////////////////////////////////////////////

#### func (*QueryTarget) Parse

```go
func (v *QueryTarget) Parse(s string) error
```

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

#### func (*QueryTarget_EXT_disjoint_timer_query) Parse

```go
func (v *QueryTarget_EXT_disjoint_timer_query) Parse(s string) error
```

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

#### func (*QueryTarget_GLES_3) Parse

```go
func (v *QueryTarget_GLES_3) Parse(s string) error
```

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
	Data      U8ˢ
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

#### type RenderbufferFormat

```go
type RenderbufferFormat uint32
```

//////////////////////////////////////////////////////////////////////////////
enum RenderbufferFormat
//////////////////////////////////////////////////////////////////////////////

#### func (*RenderbufferFormat) Parse

```go
func (v *RenderbufferFormat) Parse(s string) error
```

#### func (RenderbufferFormat) String

```go
func (v RenderbufferFormat) String() string
```

#### type RenderbufferId

```go
type RenderbufferId uint32
```


#### type RenderbufferIdːRenderbufferʳᵐ

```go
type RenderbufferIdːRenderbufferʳᵐ map[RenderbufferId](*Renderbuffer)
```


#### func (RenderbufferIdːRenderbufferʳᵐ) Contains

```go
func (m RenderbufferIdːRenderbufferʳᵐ) Contains(key RenderbufferId) bool
```

#### func (RenderbufferIdːRenderbufferʳᵐ) Delete

```go
func (m RenderbufferIdːRenderbufferʳᵐ) Delete(key RenderbufferId)
```

#### func (RenderbufferIdːRenderbufferʳᵐ) Get

```go
func (m RenderbufferIdːRenderbufferʳᵐ) Get(key RenderbufferId) *Renderbuffer
```

#### func (RenderbufferIdːRenderbufferʳᵐ) Range

```go
func (m RenderbufferIdːRenderbufferʳᵐ) Range() [](*Renderbuffer)
```

#### type RenderbufferIdˢ

```go
type RenderbufferIdˢ struct {
	binary.Generate
	SliceInfo
}
```

RenderbufferIdˢ is a slice of RenderbufferId.

#### func  AsRenderbufferIdˢ

```go
func AsRenderbufferIdˢ(s Slice, ϟs *gfxapi.State) RenderbufferIdˢ
```
AsRenderbufferIdˢ returns s cast to a RenderbufferIdˢ. The returned slice length
will be calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeRenderbufferIdˢ

```go
func MakeRenderbufferIdˢ(count uint64, ϟs *gfxapi.State) RenderbufferIdˢ
```
MakeRenderbufferIdˢ returns a RenderbufferIdˢ backed by a new memory pool.

#### func (*RenderbufferIdˢ) Class

```go
func (*RenderbufferIdˢ) Class() binary.Class
```

#### func (RenderbufferIdˢ) Clone

```go
func (s RenderbufferIdˢ) Clone(ϟs *gfxapi.State) RenderbufferIdˢ
```
Clone returns a copy of the RenderbufferIdˢ in a new memory pool.

#### func (RenderbufferIdˢ) Copy

```go
func (dst RenderbufferIdˢ) Copy(src RenderbufferIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s RenderbufferIdˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (RenderbufferIdˢ) Decoder

```go
func (s RenderbufferIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (RenderbufferIdˢ) ElementSize

```go
func (s RenderbufferIdˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that RenderbufferIdˢ points
to.

#### func (RenderbufferIdˢ) Encoder

```go
func (s RenderbufferIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (RenderbufferIdˢ) Index

```go
func (s RenderbufferIdˢ) Index(i uint64, ϟs *gfxapi.State) RenderbufferIdᵖ
```
Index returns a RenderbufferIdᵖ to the i'th element in this RenderbufferIdˢ.

#### func (RenderbufferIdˢ) OnRead

```go
func (s RenderbufferIdˢ) OnRead(ϟs *gfxapi.State) RenderbufferIdˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (RenderbufferIdˢ) OnWrite

```go
func (s RenderbufferIdˢ) OnWrite(ϟs *gfxapi.State) RenderbufferIdˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (RenderbufferIdˢ) Range

```go
func (s RenderbufferIdˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (RenderbufferIdˢ) Read

```go
func (s RenderbufferIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []RenderbufferId
```
Read reads and returns all the RenderbufferId elements in this RenderbufferIdˢ.

#### func (RenderbufferIdˢ) ResourceID

```go
func (s RenderbufferIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (RenderbufferIdˢ) Slice

```go
func (s RenderbufferIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ
```
Slice returns a sub-slice from the RenderbufferIdˢ using start and end indices.

#### func (RenderbufferIdˢ) String

```go
func (s RenderbufferIdˢ) String() string
```
String returns a string description of the RenderbufferIdˢ slice.

#### func (RenderbufferIdˢ) Write

```go
func (s RenderbufferIdˢ) Write(src []RenderbufferId, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type RenderbufferIdᵖ

```go
type RenderbufferIdᵖ struct {
	binary.Generate
	memory.Pointer
}
```

RenderbufferIdᵖ is a pointer to a RenderbufferId element.

#### func  NewRenderbufferIdᵖ

```go
func NewRenderbufferIdᵖ(addr uint64) RenderbufferIdᵖ
```
NewRenderbufferIdᵖ returns a RenderbufferIdᵖ that points to addr in the
application pool.

#### func (*RenderbufferIdᵖ) Class

```go
func (*RenderbufferIdᵖ) Class() binary.Class
```

#### func (RenderbufferIdᵖ) ElementSize

```go
func (p RenderbufferIdᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that RenderbufferIdᵖ points
to.

#### func (RenderbufferIdᵖ) OnRead

```go
func (p RenderbufferIdᵖ) OnRead(ϟs *gfxapi.State) RenderbufferIdᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (RenderbufferIdᵖ) OnWrite

```go
func (p RenderbufferIdᵖ) OnWrite(ϟs *gfxapi.State) RenderbufferIdᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (RenderbufferIdᵖ) Read

```go
func (p RenderbufferIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) RenderbufferId
```
Read reads and returns the RenderbufferId element at the pointer.

#### func (RenderbufferIdᵖ) Slice

```go
func (p RenderbufferIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ
```
Slice returns a new RenderbufferIdˢ from the pointer using start and end
indices.

#### func (RenderbufferIdᵖ) Write

```go
func (p RenderbufferIdᵖ) Write(value RenderbufferId, ϟs *gfxapi.State)
```
Write writes value to the RenderbufferId element at the pointer.

#### type RenderbufferIdᶜᵖ

```go
type RenderbufferIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

RenderbufferIdᶜᵖ is a pointer to a RenderbufferId element.

#### func  NewRenderbufferIdᶜᵖ

```go
func NewRenderbufferIdᶜᵖ(addr uint64) RenderbufferIdᶜᵖ
```
NewRenderbufferIdᶜᵖ returns a RenderbufferIdᶜᵖ that points to addr in the
application pool.

#### func (*RenderbufferIdᶜᵖ) Class

```go
func (*RenderbufferIdᶜᵖ) Class() binary.Class
```

#### func (RenderbufferIdᶜᵖ) ElementSize

```go
func (p RenderbufferIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that RenderbufferIdᶜᵖ points
to.

#### func (RenderbufferIdᶜᵖ) OnRead

```go
func (p RenderbufferIdᶜᵖ) OnRead(ϟs *gfxapi.State) RenderbufferIdᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (RenderbufferIdᶜᵖ) OnWrite

```go
func (p RenderbufferIdᶜᵖ) OnWrite(ϟs *gfxapi.State) RenderbufferIdᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (RenderbufferIdᶜᵖ) Read

```go
func (p RenderbufferIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) RenderbufferId
```
Read reads and returns the RenderbufferId element at the pointer.

#### func (RenderbufferIdᶜᵖ) Slice

```go
func (p RenderbufferIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) RenderbufferIdˢ
```
Slice returns a new RenderbufferIdˢ from the pointer using start and end
indices.

#### func (RenderbufferIdᶜᵖ) Write

```go
func (p RenderbufferIdᶜᵖ) Write(value RenderbufferId, ϟs *gfxapi.State)
```
Write writes value to the RenderbufferId element at the pointer.

#### type RenderbufferParameter

```go
type RenderbufferParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum RenderbufferParameter
//////////////////////////////////////////////////////////////////////////////

#### func (*RenderbufferParameter) Parse

```go
func (v *RenderbufferParameter) Parse(s string) error
```

#### func (RenderbufferParameter) String

```go
func (v RenderbufferParameter) String() string
```

#### type RenderbufferTarget

```go
type RenderbufferTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum RenderbufferTarget
//////////////////////////////////////////////////////////////////////////////

#### func (*RenderbufferTarget) Parse

```go
func (v *RenderbufferTarget) Parse(s string) error
```

#### func (RenderbufferTarget) String

```go
func (v RenderbufferTarget) String() string
```

#### type RenderbufferTargetːRenderbufferIdᵐ

```go
type RenderbufferTargetːRenderbufferIdᵐ map[RenderbufferTarget]RenderbufferId
```


#### func (RenderbufferTargetːRenderbufferIdᵐ) Contains

```go
func (m RenderbufferTargetːRenderbufferIdᵐ) Contains(key RenderbufferTarget) bool
```

#### func (RenderbufferTargetːRenderbufferIdᵐ) Delete

```go
func (m RenderbufferTargetːRenderbufferIdᵐ) Delete(key RenderbufferTarget)
```

#### func (RenderbufferTargetːRenderbufferIdᵐ) Get

```go
func (m RenderbufferTargetːRenderbufferIdᵐ) Get(key RenderbufferTarget) RenderbufferId
```

#### func (RenderbufferTargetːRenderbufferIdᵐ) Range

```go
func (m RenderbufferTargetːRenderbufferIdᵐ) Range() []RenderbufferId
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
func NewReplayBindRenderer(Id uint32) *ReplayBindRenderer
```

#### func (*ReplayBindRenderer) API

```go
func (c *ReplayBindRenderer) API() gfxapi.ID
```

#### func (*ReplayBindRenderer) AddRead

```go
func (a *ReplayBindRenderer) AddRead(rng memory.Range, id binary.ID) *ReplayBindRenderer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The ReplayBindRenderer pointer is returned so that calls can be
chained.

#### func (*ReplayBindRenderer) AddWrite

```go
func (a *ReplayBindRenderer) AddWrite(rng memory.Range, id binary.ID) *ReplayBindRenderer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The ReplayBindRenderer pointer is returned so that calls can be
chained.

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
func (ϟa *ReplayBindRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*ReplayBindRenderer) Observations

```go
func (a *ReplayBindRenderer) Observations() *atom.Observations
```

#### func (*ReplayBindRenderer) Replay

```go
func (ϟa *ReplayBindRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*ReplayBindRenderer) String

```go
func (a *ReplayBindRenderer) String() string
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
func NewReplayCreateRenderer(Id uint32) *ReplayCreateRenderer
```

#### func (*ReplayCreateRenderer) API

```go
func (c *ReplayCreateRenderer) API() gfxapi.ID
```

#### func (*ReplayCreateRenderer) AddRead

```go
func (a *ReplayCreateRenderer) AddRead(rng memory.Range, id binary.ID) *ReplayCreateRenderer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The ReplayCreateRenderer pointer is returned so that calls can be
chained.

#### func (*ReplayCreateRenderer) AddWrite

```go
func (a *ReplayCreateRenderer) AddWrite(rng memory.Range, id binary.ID) *ReplayCreateRenderer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The ReplayCreateRenderer pointer is returned so that calls can be
chained.

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
func (ϟa *ReplayCreateRenderer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*ReplayCreateRenderer) Observations

```go
func (a *ReplayCreateRenderer) Observations() *atom.Observations
```

#### func (*ReplayCreateRenderer) Replay

```go
func (ϟa *ReplayCreateRenderer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*ReplayCreateRenderer) String

```go
func (a *ReplayCreateRenderer) String() string
```

#### type ResetStatus

```go
type ResetStatus uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ResetStatus
//////////////////////////////////////////////////////////////////////////////

#### func (*ResetStatus) Parse

```go
func (v *ResetStatus) Parse(s string) error
```

#### func (ResetStatus) String

```go
func (v ResetStatus) String() string
```

#### type S32ː2ᵃ

```go
type S32ː2ᵃ struct {
	binary.Generate
	Elements [2]int32
}
```


#### func (*S32ː2ᵃ) Class

```go
func (*S32ː2ᵃ) Class() binary.Class
```

#### type S32ː3ᵃ

```go
type S32ː3ᵃ struct {
	binary.Generate
	Elements [3]int32
}
```


#### func (*S32ː3ᵃ) Class

```go
func (*S32ː3ᵃ) Class() binary.Class
```

#### type S32ː4ᵃ

```go
type S32ː4ᵃ struct {
	binary.Generate
	Elements [4]int32
}
```


#### func (*S32ː4ᵃ) Class

```go
func (*S32ː4ᵃ) Class() binary.Class
```

#### type S32ːCubemapLevelᵐ

```go
type S32ːCubemapLevelᵐ map[int32]CubemapLevel
```


#### func (S32ːCubemapLevelᵐ) Contains

```go
func (m S32ːCubemapLevelᵐ) Contains(key int32) bool
```

#### func (S32ːCubemapLevelᵐ) Delete

```go
func (m S32ːCubemapLevelᵐ) Delete(key int32)
```

#### func (S32ːCubemapLevelᵐ) Get

```go
func (m S32ːCubemapLevelᵐ) Get(key int32) CubemapLevel
```

#### func (S32ːCubemapLevelᵐ) Range

```go
func (m S32ːCubemapLevelᵐ) Range() []CubemapLevel
```

#### type S32ːImageᵐ

```go
type S32ːImageᵐ map[int32]Image
```


#### func (S32ːImageᵐ) Contains

```go
func (m S32ːImageᵐ) Contains(key int32) bool
```

#### func (S32ːImageᵐ) Delete

```go
func (m S32ːImageᵐ) Delete(key int32)
```

#### func (S32ːImageᵐ) Get

```go
func (m S32ːImageᵐ) Get(key int32) Image
```

#### func (S32ːImageᵐ) Range

```go
func (m S32ːImageᵐ) Range() []Image
```

#### type S32ːVertexAttributeᵐ

```go
type S32ːVertexAttributeᵐ map[int32]VertexAttribute
```


#### func (S32ːVertexAttributeᵐ) Contains

```go
func (m S32ːVertexAttributeᵐ) Contains(key int32) bool
```

#### func (S32ːVertexAttributeᵐ) Delete

```go
func (m S32ːVertexAttributeᵐ) Delete(key int32)
```

#### func (S32ːVertexAttributeᵐ) Get

```go
func (m S32ːVertexAttributeᵐ) Get(key int32) VertexAttribute
```

#### func (S32ːVertexAttributeᵐ) Range

```go
func (m S32ːVertexAttributeᵐ) Range() []VertexAttribute
```

#### type S32ˢ

```go
type S32ˢ struct {
	binary.Generate
	SliceInfo
}
```

S32ˢ is a slice of int32.

#### func  AsS32ˢ

```go
func AsS32ˢ(s Slice, ϟs *gfxapi.State) S32ˢ
```
AsS32ˢ returns s cast to a S32ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeS32ˢ

```go
func MakeS32ˢ(count uint64, ϟs *gfxapi.State) S32ˢ
```
MakeS32ˢ returns a S32ˢ backed by a new memory pool.

#### func (*S32ˢ) Class

```go
func (*S32ˢ) Class() binary.Class
```

#### func (S32ˢ) Clone

```go
func (s S32ˢ) Clone(ϟs *gfxapi.State) S32ˢ
```
Clone returns a copy of the S32ˢ in a new memory pool.

#### func (S32ˢ) Copy

```go
func (dst S32ˢ) Copy(src S32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S32ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (S32ˢ) Decoder

```go
func (s S32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (S32ˢ) ElementSize

```go
func (s S32ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S32ˢ points to.

#### func (S32ˢ) Encoder

```go
func (s S32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (S32ˢ) Index

```go
func (s S32ˢ) Index(i uint64, ϟs *gfxapi.State) S32ᵖ
```
Index returns a S32ᵖ to the i'th element in this S32ˢ.

#### func (S32ˢ) OnRead

```go
func (s S32ˢ) OnRead(ϟs *gfxapi.State) S32ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (S32ˢ) OnWrite

```go
func (s S32ˢ) OnWrite(ϟs *gfxapi.State) S32ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (S32ˢ) Range

```go
func (s S32ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (S32ˢ) Read

```go
func (s S32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int32
```
Read reads and returns all the int32 elements in this S32ˢ.

#### func (S32ˢ) ResourceID

```go
func (s S32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (S32ˢ) Slice

```go
func (s S32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ
```
Slice returns a sub-slice from the S32ˢ using start and end indices.

#### func (S32ˢ) String

```go
func (s S32ˢ) String() string
```
String returns a string description of the S32ˢ slice.

#### func (S32ˢ) Write

```go
func (s S32ˢ) Write(src []int32, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type S32ᵖ

```go
type S32ᵖ struct {
	binary.Generate
	memory.Pointer
}
```

S32ᵖ is a pointer to a int32 element.

#### func  NewS32ᵖ

```go
func NewS32ᵖ(addr uint64) S32ᵖ
```
NewS32ᵖ returns a S32ᵖ that points to addr in the application pool.

#### func (*S32ᵖ) Class

```go
func (*S32ᵖ) Class() binary.Class
```

#### func (S32ᵖ) ElementSize

```go
func (p S32ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S32ᵖ points to.

#### func (S32ᵖ) OnRead

```go
func (p S32ᵖ) OnRead(ϟs *gfxapi.State) S32ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (S32ᵖ) OnWrite

```go
func (p S32ᵖ) OnWrite(ϟs *gfxapi.State) S32ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (S32ᵖ) Read

```go
func (p S32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int32
```
Read reads and returns the int32 element at the pointer.

#### func (S32ᵖ) Slice

```go
func (p S32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ
```
Slice returns a new S32ˢ from the pointer using start and end indices.

#### func (S32ᵖ) Write

```go
func (p S32ᵖ) Write(value int32, ϟs *gfxapi.State)
```
Write writes value to the int32 element at the pointer.

#### type S32ᶜᵖ

```go
type S32ᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

S32ᶜᵖ is a pointer to a int32 element.

#### func  NewS32ᶜᵖ

```go
func NewS32ᶜᵖ(addr uint64) S32ᶜᵖ
```
NewS32ᶜᵖ returns a S32ᶜᵖ that points to addr in the application pool.

#### func (*S32ᶜᵖ) Class

```go
func (*S32ᶜᵖ) Class() binary.Class
```

#### func (S32ᶜᵖ) ElementSize

```go
func (p S32ᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S32ᶜᵖ points to.

#### func (S32ᶜᵖ) OnRead

```go
func (p S32ᶜᵖ) OnRead(ϟs *gfxapi.State) S32ᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (S32ᶜᵖ) OnWrite

```go
func (p S32ᶜᵖ) OnWrite(ϟs *gfxapi.State) S32ᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (S32ᶜᵖ) Read

```go
func (p S32ᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int32
```
Read reads and returns the int32 element at the pointer.

#### func (S32ᶜᵖ) Slice

```go
func (p S32ᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S32ˢ
```
Slice returns a new S32ˢ from the pointer using start and end indices.

#### func (S32ᶜᵖ) Write

```go
func (p S32ᶜᵖ) Write(value int32, ϟs *gfxapi.State)
```
Write writes value to the int32 element at the pointer.

#### type S64ˢ

```go
type S64ˢ struct {
	binary.Generate
	SliceInfo
}
```

S64ˢ is a slice of int64.

#### func  AsS64ˢ

```go
func AsS64ˢ(s Slice, ϟs *gfxapi.State) S64ˢ
```
AsS64ˢ returns s cast to a S64ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeS64ˢ

```go
func MakeS64ˢ(count uint64, ϟs *gfxapi.State) S64ˢ
```
MakeS64ˢ returns a S64ˢ backed by a new memory pool.

#### func (*S64ˢ) Class

```go
func (*S64ˢ) Class() binary.Class
```

#### func (S64ˢ) Clone

```go
func (s S64ˢ) Clone(ϟs *gfxapi.State) S64ˢ
```
Clone returns a copy of the S64ˢ in a new memory pool.

#### func (S64ˢ) Copy

```go
func (dst S64ˢ) Copy(src S64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s S64ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (S64ˢ) Decoder

```go
func (s S64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (S64ˢ) ElementSize

```go
func (s S64ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S64ˢ points to.

#### func (S64ˢ) Encoder

```go
func (s S64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (S64ˢ) Index

```go
func (s S64ˢ) Index(i uint64, ϟs *gfxapi.State) S64ᵖ
```
Index returns a S64ᵖ to the i'th element in this S64ˢ.

#### func (S64ˢ) OnRead

```go
func (s S64ˢ) OnRead(ϟs *gfxapi.State) S64ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (S64ˢ) OnWrite

```go
func (s S64ˢ) OnWrite(ϟs *gfxapi.State) S64ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (S64ˢ) Range

```go
func (s S64ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (S64ˢ) Read

```go
func (s S64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []int64
```
Read reads and returns all the int64 elements in this S64ˢ.

#### func (S64ˢ) ResourceID

```go
func (s S64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (S64ˢ) Slice

```go
func (s S64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ
```
Slice returns a sub-slice from the S64ˢ using start and end indices.

#### func (S64ˢ) String

```go
func (s S64ˢ) String() string
```
String returns a string description of the S64ˢ slice.

#### func (S64ˢ) Write

```go
func (s S64ˢ) Write(src []int64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type S64ᵖ

```go
type S64ᵖ struct {
	binary.Generate
	memory.Pointer
}
```

S64ᵖ is a pointer to a int64 element.

#### func  NewS64ᵖ

```go
func NewS64ᵖ(addr uint64) S64ᵖ
```
NewS64ᵖ returns a S64ᵖ that points to addr in the application pool.

#### func (*S64ᵖ) Class

```go
func (*S64ᵖ) Class() binary.Class
```

#### func (S64ᵖ) ElementSize

```go
func (p S64ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that S64ᵖ points to.

#### func (S64ᵖ) OnRead

```go
func (p S64ᵖ) OnRead(ϟs *gfxapi.State) S64ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (S64ᵖ) OnWrite

```go
func (p S64ᵖ) OnWrite(ϟs *gfxapi.State) S64ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (S64ᵖ) Read

```go
func (p S64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) int64
```
Read reads and returns the int64 element at the pointer.

#### func (S64ᵖ) Slice

```go
func (p S64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) S64ˢ
```
Slice returns a new S64ˢ from the pointer using start and end indices.

#### func (S64ᵖ) Write

```go
func (p S64ᵖ) Write(value int64, ϟs *gfxapi.State)
```
Write writes value to the int64 element at the pointer.

#### type Shader

```go
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

#### type ShaderAttribType

```go
type ShaderAttribType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderAttribType
//////////////////////////////////////////////////////////////////////////////

#### func (*ShaderAttribType) Parse

```go
func (v *ShaderAttribType) Parse(s string) error
```

#### func (ShaderAttribType) String

```go
func (v ShaderAttribType) String() string
```

#### type ShaderAttribTypeˢ

```go
type ShaderAttribTypeˢ struct {
	binary.Generate
	SliceInfo
}
```

ShaderAttribTypeˢ is a slice of ShaderAttribType.

#### func  AsShaderAttribTypeˢ

```go
func AsShaderAttribTypeˢ(s Slice, ϟs *gfxapi.State) ShaderAttribTypeˢ
```
AsShaderAttribTypeˢ returns s cast to a ShaderAttribTypeˢ. The returned slice
length will be calculated so that the returned slice is no longer (in bytes)
than s.

#### func  MakeShaderAttribTypeˢ

```go
func MakeShaderAttribTypeˢ(count uint64, ϟs *gfxapi.State) ShaderAttribTypeˢ
```
MakeShaderAttribTypeˢ returns a ShaderAttribTypeˢ backed by a new memory pool.

#### func (*ShaderAttribTypeˢ) Class

```go
func (*ShaderAttribTypeˢ) Class() binary.Class
```

#### func (ShaderAttribTypeˢ) Clone

```go
func (s ShaderAttribTypeˢ) Clone(ϟs *gfxapi.State) ShaderAttribTypeˢ
```
Clone returns a copy of the ShaderAttribTypeˢ in a new memory pool.

#### func (ShaderAttribTypeˢ) Copy

```go
func (dst ShaderAttribTypeˢ) Copy(src ShaderAttribTypeˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s ShaderAttribTypeˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (ShaderAttribTypeˢ) Decoder

```go
func (s ShaderAttribTypeˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (ShaderAttribTypeˢ) ElementSize

```go
func (s ShaderAttribTypeˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ShaderAttribTypeˢ
points to.

#### func (ShaderAttribTypeˢ) Encoder

```go
func (s ShaderAttribTypeˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (ShaderAttribTypeˢ) Index

```go
func (s ShaderAttribTypeˢ) Index(i uint64, ϟs *gfxapi.State) ShaderAttribTypeᵖ
```
Index returns a ShaderAttribTypeᵖ to the i'th element in this ShaderAttribTypeˢ.

#### func (ShaderAttribTypeˢ) OnRead

```go
func (s ShaderAttribTypeˢ) OnRead(ϟs *gfxapi.State) ShaderAttribTypeˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (ShaderAttribTypeˢ) OnWrite

```go
func (s ShaderAttribTypeˢ) OnWrite(ϟs *gfxapi.State) ShaderAttribTypeˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (ShaderAttribTypeˢ) Range

```go
func (s ShaderAttribTypeˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (ShaderAttribTypeˢ) Read

```go
func (s ShaderAttribTypeˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []ShaderAttribType
```
Read reads and returns all the ShaderAttribType elements in this
ShaderAttribTypeˢ.

#### func (ShaderAttribTypeˢ) ResourceID

```go
func (s ShaderAttribTypeˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (ShaderAttribTypeˢ) Slice

```go
func (s ShaderAttribTypeˢ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderAttribTypeˢ
```
Slice returns a sub-slice from the ShaderAttribTypeˢ using start and end
indices.

#### func (ShaderAttribTypeˢ) String

```go
func (s ShaderAttribTypeˢ) String() string
```
String returns a string description of the ShaderAttribTypeˢ slice.

#### func (ShaderAttribTypeˢ) Write

```go
func (s ShaderAttribTypeˢ) Write(src []ShaderAttribType, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type ShaderAttribTypeᵖ

```go
type ShaderAttribTypeᵖ struct {
	binary.Generate
	memory.Pointer
}
```

ShaderAttribTypeᵖ is a pointer to a ShaderAttribType element.

#### func  NewShaderAttribTypeᵖ

```go
func NewShaderAttribTypeᵖ(addr uint64) ShaderAttribTypeᵖ
```
NewShaderAttribTypeᵖ returns a ShaderAttribTypeᵖ that points to addr in the
application pool.

#### func (*ShaderAttribTypeᵖ) Class

```go
func (*ShaderAttribTypeᵖ) Class() binary.Class
```

#### func (ShaderAttribTypeᵖ) ElementSize

```go
func (p ShaderAttribTypeᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ShaderAttribTypeᵖ
points to.

#### func (ShaderAttribTypeᵖ) OnRead

```go
func (p ShaderAttribTypeᵖ) OnRead(ϟs *gfxapi.State) ShaderAttribTypeᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (ShaderAttribTypeᵖ) OnWrite

```go
func (p ShaderAttribTypeᵖ) OnWrite(ϟs *gfxapi.State) ShaderAttribTypeᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (ShaderAttribTypeᵖ) Read

```go
func (p ShaderAttribTypeᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderAttribType
```
Read reads and returns the ShaderAttribType element at the pointer.

#### func (ShaderAttribTypeᵖ) Slice

```go
func (p ShaderAttribTypeᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderAttribTypeˢ
```
Slice returns a new ShaderAttribTypeˢ from the pointer using start and end
indices.

#### func (ShaderAttribTypeᵖ) Write

```go
func (p ShaderAttribTypeᵖ) Write(value ShaderAttribType, ϟs *gfxapi.State)
```
Write writes value to the ShaderAttribType element at the pointer.

#### type ShaderId

```go
type ShaderId uint32
```


#### type ShaderIdːShaderʳᵐ

```go
type ShaderIdːShaderʳᵐ map[ShaderId](*Shader)
```


#### func (ShaderIdːShaderʳᵐ) Contains

```go
func (m ShaderIdːShaderʳᵐ) Contains(key ShaderId) bool
```

#### func (ShaderIdːShaderʳᵐ) Delete

```go
func (m ShaderIdːShaderʳᵐ) Delete(key ShaderId)
```

#### func (ShaderIdːShaderʳᵐ) Get

```go
func (m ShaderIdːShaderʳᵐ) Get(key ShaderId) *Shader
```

#### func (ShaderIdːShaderʳᵐ) Range

```go
func (m ShaderIdːShaderʳᵐ) Range() [](*Shader)
```

#### type ShaderIdˢ

```go
type ShaderIdˢ struct {
	binary.Generate
	SliceInfo
}
```

ShaderIdˢ is a slice of ShaderId.

#### func  AsShaderIdˢ

```go
func AsShaderIdˢ(s Slice, ϟs *gfxapi.State) ShaderIdˢ
```
AsShaderIdˢ returns s cast to a ShaderIdˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeShaderIdˢ

```go
func MakeShaderIdˢ(count uint64, ϟs *gfxapi.State) ShaderIdˢ
```
MakeShaderIdˢ returns a ShaderIdˢ backed by a new memory pool.

#### func (*ShaderIdˢ) Class

```go
func (*ShaderIdˢ) Class() binary.Class
```

#### func (ShaderIdˢ) Clone

```go
func (s ShaderIdˢ) Clone(ϟs *gfxapi.State) ShaderIdˢ
```
Clone returns a copy of the ShaderIdˢ in a new memory pool.

#### func (ShaderIdˢ) Copy

```go
func (dst ShaderIdˢ) Copy(src ShaderIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s ShaderIdˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (ShaderIdˢ) Decoder

```go
func (s ShaderIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (ShaderIdˢ) ElementSize

```go
func (s ShaderIdˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ShaderIdˢ points to.

#### func (ShaderIdˢ) Encoder

```go
func (s ShaderIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (ShaderIdˢ) Index

```go
func (s ShaderIdˢ) Index(i uint64, ϟs *gfxapi.State) ShaderIdᵖ
```
Index returns a ShaderIdᵖ to the i'th element in this ShaderIdˢ.

#### func (ShaderIdˢ) OnRead

```go
func (s ShaderIdˢ) OnRead(ϟs *gfxapi.State) ShaderIdˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (ShaderIdˢ) OnWrite

```go
func (s ShaderIdˢ) OnWrite(ϟs *gfxapi.State) ShaderIdˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (ShaderIdˢ) Range

```go
func (s ShaderIdˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (ShaderIdˢ) Read

```go
func (s ShaderIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []ShaderId
```
Read reads and returns all the ShaderId elements in this ShaderIdˢ.

#### func (ShaderIdˢ) ResourceID

```go
func (s ShaderIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (ShaderIdˢ) Slice

```go
func (s ShaderIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ
```
Slice returns a sub-slice from the ShaderIdˢ using start and end indices.

#### func (ShaderIdˢ) String

```go
func (s ShaderIdˢ) String() string
```
String returns a string description of the ShaderIdˢ slice.

#### func (ShaderIdˢ) Write

```go
func (s ShaderIdˢ) Write(src []ShaderId, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type ShaderIdᵖ

```go
type ShaderIdᵖ struct {
	binary.Generate
	memory.Pointer
}
```

ShaderIdᵖ is a pointer to a ShaderId element.

#### func  NewShaderIdᵖ

```go
func NewShaderIdᵖ(addr uint64) ShaderIdᵖ
```
NewShaderIdᵖ returns a ShaderIdᵖ that points to addr in the application pool.

#### func (*ShaderIdᵖ) Class

```go
func (*ShaderIdᵖ) Class() binary.Class
```

#### func (ShaderIdᵖ) ElementSize

```go
func (p ShaderIdᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ShaderIdᵖ points to.

#### func (ShaderIdᵖ) OnRead

```go
func (p ShaderIdᵖ) OnRead(ϟs *gfxapi.State) ShaderIdᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (ShaderIdᵖ) OnWrite

```go
func (p ShaderIdᵖ) OnWrite(ϟs *gfxapi.State) ShaderIdᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (ShaderIdᵖ) Read

```go
func (p ShaderIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderId
```
Read reads and returns the ShaderId element at the pointer.

#### func (ShaderIdᵖ) Slice

```go
func (p ShaderIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ
```
Slice returns a new ShaderIdˢ from the pointer using start and end indices.

#### func (ShaderIdᵖ) Write

```go
func (p ShaderIdᵖ) Write(value ShaderId, ϟs *gfxapi.State)
```
Write writes value to the ShaderId element at the pointer.

#### type ShaderIdᶜᵖ

```go
type ShaderIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

ShaderIdᶜᵖ is a pointer to a ShaderId element.

#### func  NewShaderIdᶜᵖ

```go
func NewShaderIdᶜᵖ(addr uint64) ShaderIdᶜᵖ
```
NewShaderIdᶜᵖ returns a ShaderIdᶜᵖ that points to addr in the application pool.

#### func (*ShaderIdᶜᵖ) Class

```go
func (*ShaderIdᶜᵖ) Class() binary.Class
```

#### func (ShaderIdᶜᵖ) ElementSize

```go
func (p ShaderIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ShaderIdᶜᵖ points to.

#### func (ShaderIdᶜᵖ) OnRead

```go
func (p ShaderIdᶜᵖ) OnRead(ϟs *gfxapi.State) ShaderIdᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (ShaderIdᶜᵖ) OnWrite

```go
func (p ShaderIdᶜᵖ) OnWrite(ϟs *gfxapi.State) ShaderIdᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (ShaderIdᶜᵖ) Read

```go
func (p ShaderIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderId
```
Read reads and returns the ShaderId element at the pointer.

#### func (ShaderIdᶜᵖ) Slice

```go
func (p ShaderIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderIdˢ
```
Slice returns a new ShaderIdˢ from the pointer using start and end indices.

#### func (ShaderIdᶜᵖ) Write

```go
func (p ShaderIdᶜᵖ) Write(value ShaderId, ϟs *gfxapi.State)
```
Write writes value to the ShaderId element at the pointer.

#### type ShaderParameter

```go
type ShaderParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderParameter
//////////////////////////////////////////////////////////////////////////////

#### func (*ShaderParameter) Parse

```go
func (v *ShaderParameter) Parse(s string) error
```

#### func (ShaderParameter) String

```go
func (v ShaderParameter) String() string
```

#### type ShaderType

```go
type ShaderType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderType
//////////////////////////////////////////////////////////////////////////////

#### func (*ShaderType) Parse

```go
func (v *ShaderType) Parse(s string) error
```

#### func (ShaderType) String

```go
func (v ShaderType) String() string
```

#### type ShaderTypeːShaderIdᵐ

```go
type ShaderTypeːShaderIdᵐ map[ShaderType]ShaderId
```


#### func (ShaderTypeːShaderIdᵐ) Contains

```go
func (m ShaderTypeːShaderIdᵐ) Contains(key ShaderType) bool
```

#### func (ShaderTypeːShaderIdᵐ) Delete

```go
func (m ShaderTypeːShaderIdᵐ) Delete(key ShaderType)
```

#### func (ShaderTypeːShaderIdᵐ) Get

```go
func (m ShaderTypeːShaderIdᵐ) Get(key ShaderType) ShaderId
```

#### func (ShaderTypeːShaderIdᵐ) Range

```go
func (m ShaderTypeːShaderIdᵐ) Range() []ShaderId
```

#### type ShaderUniformType

```go
type ShaderUniformType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum ShaderUniformType
//////////////////////////////////////////////////////////////////////////////

#### func (*ShaderUniformType) Parse

```go
func (v *ShaderUniformType) Parse(s string) error
```

#### func (ShaderUniformType) String

```go
func (v ShaderUniformType) String() string
```

#### type ShaderUniformTypeˢ

```go
type ShaderUniformTypeˢ struct {
	binary.Generate
	SliceInfo
}
```

ShaderUniformTypeˢ is a slice of ShaderUniformType.

#### func  AsShaderUniformTypeˢ

```go
func AsShaderUniformTypeˢ(s Slice, ϟs *gfxapi.State) ShaderUniformTypeˢ
```
AsShaderUniformTypeˢ returns s cast to a ShaderUniformTypeˢ. The returned slice
length will be calculated so that the returned slice is no longer (in bytes)
than s.

#### func  MakeShaderUniformTypeˢ

```go
func MakeShaderUniformTypeˢ(count uint64, ϟs *gfxapi.State) ShaderUniformTypeˢ
```
MakeShaderUniformTypeˢ returns a ShaderUniformTypeˢ backed by a new memory pool.

#### func (*ShaderUniformTypeˢ) Class

```go
func (*ShaderUniformTypeˢ) Class() binary.Class
```

#### func (ShaderUniformTypeˢ) Clone

```go
func (s ShaderUniformTypeˢ) Clone(ϟs *gfxapi.State) ShaderUniformTypeˢ
```
Clone returns a copy of the ShaderUniformTypeˢ in a new memory pool.

#### func (ShaderUniformTypeˢ) Copy

```go
func (dst ShaderUniformTypeˢ) Copy(src ShaderUniformTypeˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s ShaderUniformTypeˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (ShaderUniformTypeˢ) Decoder

```go
func (s ShaderUniformTypeˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (ShaderUniformTypeˢ) ElementSize

```go
func (s ShaderUniformTypeˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ShaderUniformTypeˢ
points to.

#### func (ShaderUniformTypeˢ) Encoder

```go
func (s ShaderUniformTypeˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (ShaderUniformTypeˢ) Index

```go
func (s ShaderUniformTypeˢ) Index(i uint64, ϟs *gfxapi.State) ShaderUniformTypeᵖ
```
Index returns a ShaderUniformTypeᵖ to the i'th element in this
ShaderUniformTypeˢ.

#### func (ShaderUniformTypeˢ) OnRead

```go
func (s ShaderUniformTypeˢ) OnRead(ϟs *gfxapi.State) ShaderUniformTypeˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (ShaderUniformTypeˢ) OnWrite

```go
func (s ShaderUniformTypeˢ) OnWrite(ϟs *gfxapi.State) ShaderUniformTypeˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (ShaderUniformTypeˢ) Range

```go
func (s ShaderUniformTypeˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (ShaderUniformTypeˢ) Read

```go
func (s ShaderUniformTypeˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []ShaderUniformType
```
Read reads and returns all the ShaderUniformType elements in this
ShaderUniformTypeˢ.

#### func (ShaderUniformTypeˢ) ResourceID

```go
func (s ShaderUniformTypeˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (ShaderUniformTypeˢ) Slice

```go
func (s ShaderUniformTypeˢ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderUniformTypeˢ
```
Slice returns a sub-slice from the ShaderUniformTypeˢ using start and end
indices.

#### func (ShaderUniformTypeˢ) String

```go
func (s ShaderUniformTypeˢ) String() string
```
String returns a string description of the ShaderUniformTypeˢ slice.

#### func (ShaderUniformTypeˢ) Write

```go
func (s ShaderUniformTypeˢ) Write(src []ShaderUniformType, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type ShaderUniformTypeᵖ

```go
type ShaderUniformTypeᵖ struct {
	binary.Generate
	memory.Pointer
}
```

ShaderUniformTypeᵖ is a pointer to a ShaderUniformType element.

#### func  NewShaderUniformTypeᵖ

```go
func NewShaderUniformTypeᵖ(addr uint64) ShaderUniformTypeᵖ
```
NewShaderUniformTypeᵖ returns a ShaderUniformTypeᵖ that points to addr in the
application pool.

#### func (*ShaderUniformTypeᵖ) Class

```go
func (*ShaderUniformTypeᵖ) Class() binary.Class
```

#### func (ShaderUniformTypeᵖ) ElementSize

```go
func (p ShaderUniformTypeᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that ShaderUniformTypeᵖ
points to.

#### func (ShaderUniformTypeᵖ) OnRead

```go
func (p ShaderUniformTypeᵖ) OnRead(ϟs *gfxapi.State) ShaderUniformTypeᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (ShaderUniformTypeᵖ) OnWrite

```go
func (p ShaderUniformTypeᵖ) OnWrite(ϟs *gfxapi.State) ShaderUniformTypeᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (ShaderUniformTypeᵖ) Read

```go
func (p ShaderUniformTypeᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) ShaderUniformType
```
Read reads and returns the ShaderUniformType element at the pointer.

#### func (ShaderUniformTypeᵖ) Slice

```go
func (p ShaderUniformTypeᵖ) Slice(start, end uint64, ϟs *gfxapi.State) ShaderUniformTypeˢ
```
Slice returns a new ShaderUniformTypeˢ from the pointer using start and end
indices.

#### func (ShaderUniformTypeᵖ) Write

```go
func (p ShaderUniformTypeᵖ) Write(value ShaderUniformType, ϟs *gfxapi.State)
```
Write writes value to the ShaderUniformType element at the pointer.

#### type Slice

```go
type Slice interface {
	// Info returns the SliceInfo of this slice.
	Info() SliceInfo
	// ElementSize returns the size in bytes of a single element in the slice.
	ElementSize(ϟs *gfxapi.State) uint64
}
```

Slice is the interface implemented by all slice types

#### type SliceInfo

```go
type SliceInfo struct {
	binary.Generate
	Root  memory.Pointer // Original pointer this slice derives from.
	Base  uint64         // Address of first element.
	Count uint64         // Number of elements in the slice.
}
```

SliceInfo is the common data between all slice types.

#### func (*SliceInfo) Class

```go
func (*SliceInfo) Class() binary.Class
```

#### func (SliceInfo) Info

```go
func (s SliceInfo) Info() SliceInfo
```
Info returns the SliceInfo. It is used to conform to the Slice interface.

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
func NewStartTimer(Index uint8) *StartTimer
```

#### func (*StartTimer) API

```go
func (c *StartTimer) API() gfxapi.ID
```

#### func (*StartTimer) AddRead

```go
func (a *StartTimer) AddRead(rng memory.Range, id binary.ID) *StartTimer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The StartTimer pointer is returned so that calls can be chained.

#### func (*StartTimer) AddWrite

```go
func (a *StartTimer) AddWrite(rng memory.Range, id binary.ID) *StartTimer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The StartTimer pointer is returned so that calls can be chained.

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
func (ϟa *StartTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*StartTimer) Observations

```go
func (a *StartTimer) Observations() *atom.Observations
```

#### func (*StartTimer) Replay

```go
func (ϟa *StartTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*StartTimer) String

```go
func (a *StartTimer) String() string
```

#### type State

```go
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
```

//////////////////////////////////////////////////////////////////////////////
State
//////////////////////////////////////////////////////////////////////////////

#### func (*State) Class

```go
func (*State) Class() binary.Class
```

#### func (*State) Init

```go
func (g *State) Init()
```

#### type StateVariable

```go
type StateVariable uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StateVariable
//////////////////////////////////////////////////////////////////////////////

#### func (*StateVariable) Parse

```go
func (v *StateVariable) Parse(s string) error
```

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

#### func (*StateVariable_EXT_disjoint_timer_query) Parse

```go
func (v *StateVariable_EXT_disjoint_timer_query) Parse(s string) error
```

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

#### func (*StateVariable_EXT_texture_filter_anisotropic) Parse

```go
func (v *StateVariable_EXT_texture_filter_anisotropic) Parse(s string) error
```

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

#### func (*StateVariable_GLES_2_0) Parse

```go
func (v *StateVariable_GLES_2_0) Parse(s string) error
```

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

#### func (*StateVariable_GLES_3_1) Parse

```go
func (v *StateVariable_GLES_3_1) Parse(s string) error
```

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

#### func (*StencilAction) Parse

```go
func (v *StencilAction) Parse(s string) error
```

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
func NewStopTimer(Index uint8, Result uint64) *StopTimer
```

#### func (*StopTimer) API

```go
func (c *StopTimer) API() gfxapi.ID
```

#### func (*StopTimer) AddRead

```go
func (a *StopTimer) AddRead(rng memory.Range, id binary.ID) *StopTimer
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The StopTimer pointer is returned so that calls can be chained.

#### func (*StopTimer) AddWrite

```go
func (a *StopTimer) AddWrite(rng memory.Range, id binary.ID) *StopTimer
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The StopTimer pointer is returned so that calls can be chained.

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
func (ϟa *StopTimer) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*StopTimer) Observations

```go
func (a *StopTimer) Observations() *atom.Observations
```

#### func (*StopTimer) Replay

```go
func (ϟa *StopTimer) Replay(ϟi atom.ID, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger, ϟb *builder.Builder) error
```

#### func (*StopTimer) String

```go
func (a *StopTimer) String() string
```

#### type StringConstant

```go
type StringConstant uint32
```

//////////////////////////////////////////////////////////////////////////////
enum StringConstant
//////////////////////////////////////////////////////////////////////////////

#### func (*StringConstant) Parse

```go
func (v *StringConstant) Parse(s string) error
```

#### func (StringConstant) String

```go
func (v StringConstant) String() string
```

#### type StringːAttributeLocationᵐ

```go
type StringːAttributeLocationᵐ map[string]AttributeLocation
```


#### func (StringːAttributeLocationᵐ) Contains

```go
func (m StringːAttributeLocationᵐ) Contains(key string) bool
```

#### func (StringːAttributeLocationᵐ) Delete

```go
func (m StringːAttributeLocationᵐ) Delete(key string)
```

#### func (StringːAttributeLocationᵐ) Get

```go
func (m StringːAttributeLocationᵐ) Get(key string) AttributeLocation
```

#### func (StringːAttributeLocationᵐ) Range

```go
func (m StringːAttributeLocationᵐ) Range() []AttributeLocation
```

#### type SyncCondition

```go
type SyncCondition uint32
```

//////////////////////////////////////////////////////////////////////////////
enum SyncCondition
//////////////////////////////////////////////////////////////////////////////

#### func (*SyncCondition) Parse

```go
func (v *SyncCondition) Parse(s string) error
```

#### func (SyncCondition) String

```go
func (v SyncCondition) String() string
```

#### type SyncFlags

```go
type SyncFlags uint32
```

//////////////////////////////////////////////////////////////////////////////
enum SyncFlags
//////////////////////////////////////////////////////////////////////////////

#### func (*SyncFlags) Parse

```go
func (v *SyncFlags) Parse(s string) error
```

#### func (SyncFlags) String

```go
func (v SyncFlags) String() string
```

#### type SyncObject

```go
type SyncObject uint64
```


#### type TestFunction

```go
type TestFunction uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TestFunction
//////////////////////////////////////////////////////////////////////////////

#### func (*TestFunction) Parse

```go
func (v *TestFunction) Parse(s string) error
```

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

#### func (*TexelComponent) Parse

```go
func (v *TexelComponent) Parse(s string) error
```

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

#### func (*TexelFormat) Parse

```go
func (v *TexelFormat) Parse(s string) error
```

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

#### func (*TexelFormat_GLES_1_1) Parse

```go
func (v *TexelFormat_GLES_1_1) Parse(s string) error
```

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

#### func (*TexelFormat_GLES_3_0) Parse

```go
func (v *TexelFormat_GLES_3_0) Parse(s string) error
```

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

#### func (*TexelType) Parse

```go
func (v *TexelType) Parse(s string) error
```

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

#### func (*Texture2DImageTarget) Parse

```go
func (v *Texture2DImageTarget) Parse(s string) error
```

#### func (Texture2DImageTarget) String

```go
func (v Texture2DImageTarget) String() string
```

#### type TextureFilterMode

```go
type TextureFilterMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureFilterMode
//////////////////////////////////////////////////////////////////////////////

#### func (*TextureFilterMode) Parse

```go
func (v *TextureFilterMode) Parse(s string) error
```

#### func (TextureFilterMode) String

```go
func (v TextureFilterMode) String() string
```

#### type TextureId

```go
type TextureId uint32
```


#### type TextureIdːTextureʳᵐ

```go
type TextureIdːTextureʳᵐ map[TextureId](*Texture)
```


#### func (TextureIdːTextureʳᵐ) Contains

```go
func (m TextureIdːTextureʳᵐ) Contains(key TextureId) bool
```

#### func (TextureIdːTextureʳᵐ) Delete

```go
func (m TextureIdːTextureʳᵐ) Delete(key TextureId)
```

#### func (TextureIdːTextureʳᵐ) Get

```go
func (m TextureIdːTextureʳᵐ) Get(key TextureId) *Texture
```

#### func (TextureIdːTextureʳᵐ) Range

```go
func (m TextureIdːTextureʳᵐ) Range() [](*Texture)
```

#### type TextureIdˢ

```go
type TextureIdˢ struct {
	binary.Generate
	SliceInfo
}
```

TextureIdˢ is a slice of TextureId.

#### func  AsTextureIdˢ

```go
func AsTextureIdˢ(s Slice, ϟs *gfxapi.State) TextureIdˢ
```
AsTextureIdˢ returns s cast to a TextureIdˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeTextureIdˢ

```go
func MakeTextureIdˢ(count uint64, ϟs *gfxapi.State) TextureIdˢ
```
MakeTextureIdˢ returns a TextureIdˢ backed by a new memory pool.

#### func (*TextureIdˢ) Class

```go
func (*TextureIdˢ) Class() binary.Class
```

#### func (TextureIdˢ) Clone

```go
func (s TextureIdˢ) Clone(ϟs *gfxapi.State) TextureIdˢ
```
Clone returns a copy of the TextureIdˢ in a new memory pool.

#### func (TextureIdˢ) Copy

```go
func (dst TextureIdˢ) Copy(src TextureIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s TextureIdˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (TextureIdˢ) Decoder

```go
func (s TextureIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (TextureIdˢ) ElementSize

```go
func (s TextureIdˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that TextureIdˢ points to.

#### func (TextureIdˢ) Encoder

```go
func (s TextureIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (TextureIdˢ) Index

```go
func (s TextureIdˢ) Index(i uint64, ϟs *gfxapi.State) TextureIdᵖ
```
Index returns a TextureIdᵖ to the i'th element in this TextureIdˢ.

#### func (TextureIdˢ) OnRead

```go
func (s TextureIdˢ) OnRead(ϟs *gfxapi.State) TextureIdˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (TextureIdˢ) OnWrite

```go
func (s TextureIdˢ) OnWrite(ϟs *gfxapi.State) TextureIdˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (TextureIdˢ) Range

```go
func (s TextureIdˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (TextureIdˢ) Read

```go
func (s TextureIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []TextureId
```
Read reads and returns all the TextureId elements in this TextureIdˢ.

#### func (TextureIdˢ) ResourceID

```go
func (s TextureIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (TextureIdˢ) Slice

```go
func (s TextureIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ
```
Slice returns a sub-slice from the TextureIdˢ using start and end indices.

#### func (TextureIdˢ) String

```go
func (s TextureIdˢ) String() string
```
String returns a string description of the TextureIdˢ slice.

#### func (TextureIdˢ) Write

```go
func (s TextureIdˢ) Write(src []TextureId, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type TextureIdᵖ

```go
type TextureIdᵖ struct {
	binary.Generate
	memory.Pointer
}
```

TextureIdᵖ is a pointer to a TextureId element.

#### func  NewTextureIdᵖ

```go
func NewTextureIdᵖ(addr uint64) TextureIdᵖ
```
NewTextureIdᵖ returns a TextureIdᵖ that points to addr in the application pool.

#### func (*TextureIdᵖ) Class

```go
func (*TextureIdᵖ) Class() binary.Class
```

#### func (TextureIdᵖ) ElementSize

```go
func (p TextureIdᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that TextureIdᵖ points to.

#### func (TextureIdᵖ) OnRead

```go
func (p TextureIdᵖ) OnRead(ϟs *gfxapi.State) TextureIdᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (TextureIdᵖ) OnWrite

```go
func (p TextureIdᵖ) OnWrite(ϟs *gfxapi.State) TextureIdᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (TextureIdᵖ) Read

```go
func (p TextureIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) TextureId
```
Read reads and returns the TextureId element at the pointer.

#### func (TextureIdᵖ) Slice

```go
func (p TextureIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ
```
Slice returns a new TextureIdˢ from the pointer using start and end indices.

#### func (TextureIdᵖ) Write

```go
func (p TextureIdᵖ) Write(value TextureId, ϟs *gfxapi.State)
```
Write writes value to the TextureId element at the pointer.

#### type TextureIdᶜᵖ

```go
type TextureIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

TextureIdᶜᵖ is a pointer to a TextureId element.

#### func  NewTextureIdᶜᵖ

```go
func NewTextureIdᶜᵖ(addr uint64) TextureIdᶜᵖ
```
NewTextureIdᶜᵖ returns a TextureIdᶜᵖ that points to addr in the application
pool.

#### func (*TextureIdᶜᵖ) Class

```go
func (*TextureIdᶜᵖ) Class() binary.Class
```

#### func (TextureIdᶜᵖ) ElementSize

```go
func (p TextureIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that TextureIdᶜᵖ points to.

#### func (TextureIdᶜᵖ) OnRead

```go
func (p TextureIdᶜᵖ) OnRead(ϟs *gfxapi.State) TextureIdᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (TextureIdᶜᵖ) OnWrite

```go
func (p TextureIdᶜᵖ) OnWrite(ϟs *gfxapi.State) TextureIdᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (TextureIdᶜᵖ) Read

```go
func (p TextureIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) TextureId
```
Read reads and returns the TextureId element at the pointer.

#### func (TextureIdᶜᵖ) Slice

```go
func (p TextureIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) TextureIdˢ
```
Slice returns a new TextureIdˢ from the pointer using start and end indices.

#### func (TextureIdᶜᵖ) Write

```go
func (p TextureIdᶜᵖ) Write(value TextureId, ϟs *gfxapi.State)
```
Write writes value to the TextureId element at the pointer.

#### type TextureImageTarget

```go
type TextureImageTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureImageTarget
//////////////////////////////////////////////////////////////////////////////

#### func (*TextureImageTarget) Parse

```go
func (v *TextureImageTarget) Parse(s string) error
```

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

#### func (*TextureKind) Parse

```go
func (v *TextureKind) Parse(s string) error
```

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

#### func (*TextureParameter) Parse

```go
func (v *TextureParameter) Parse(s string) error
```

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

#### func (*TextureParameter_EXT_texture_filter_anisotropic) Parse

```go
func (v *TextureParameter_EXT_texture_filter_anisotropic) Parse(s string) error
```

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

#### func (*TextureParameter_FilterMode) Parse

```go
func (v *TextureParameter_FilterMode) Parse(s string) error
```

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

#### func (*TextureParameter_SwizzleMode) Parse

```go
func (v *TextureParameter_SwizzleMode) Parse(s string) error
```

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

#### func (*TextureParameter_WrapMode) Parse

```go
func (v *TextureParameter_WrapMode) Parse(s string) error
```

#### func (TextureParameter_WrapMode) String

```go
func (v TextureParameter_WrapMode) String() string
```

#### type TexturePointer

```go
type TexturePointer struct {
	binary.Generate
	memory.Pointer
}
```

TexturePointer is a pointer to a void element.

#### func  NewTexturePointer

```go
func NewTexturePointer(addr uint64) TexturePointer
```
NewTexturePointer returns a TexturePointer that points to addr in the
application pool.

#### func (*TexturePointer) Class

```go
func (*TexturePointer) Class() binary.Class
```

#### func (TexturePointer) ElementSize

```go
func (p TexturePointer) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that TexturePointer points
to.

#### func (TexturePointer) OnRead

```go
func (p TexturePointer) OnRead(ϟs *gfxapi.State) TexturePointer
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (TexturePointer) OnWrite

```go
func (p TexturePointer) OnWrite(ϟs *gfxapi.State) TexturePointer
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (TexturePointer) Slice

```go
func (p TexturePointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type TextureTarget

```go
type TextureTarget uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureTarget
//////////////////////////////////////////////////////////////////////////////

#### func (*TextureTarget) Parse

```go
func (v *TextureTarget) Parse(s string) error
```

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

#### func (*TextureTarget_GLES_1_1) Parse

```go
func (v *TextureTarget_GLES_1_1) Parse(s string) error
```

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

#### func (*TextureTarget_GLES_2_0) Parse

```go
func (v *TextureTarget_GLES_2_0) Parse(s string) error
```

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

#### func (*TextureTarget_OES_EGL_image_external) Parse

```go
func (v *TextureTarget_OES_EGL_image_external) Parse(s string) error
```

#### func (TextureTarget_OES_EGL_image_external) String

```go
func (v TextureTarget_OES_EGL_image_external) String() string
```

#### type TextureTargetːTextureIdᵐ

```go
type TextureTargetːTextureIdᵐ map[TextureTarget]TextureId
```


#### func (TextureTargetːTextureIdᵐ) Contains

```go
func (m TextureTargetːTextureIdᵐ) Contains(key TextureTarget) bool
```

#### func (TextureTargetːTextureIdᵐ) Delete

```go
func (m TextureTargetːTextureIdᵐ) Delete(key TextureTarget)
```

#### func (TextureTargetːTextureIdᵐ) Get

```go
func (m TextureTargetːTextureIdᵐ) Get(key TextureTarget) TextureId
```

#### func (TextureTargetːTextureIdᵐ) Range

```go
func (m TextureTargetːTextureIdᵐ) Range() []TextureId
```

#### type TextureUnit

```go
type TextureUnit uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureUnit
//////////////////////////////////////////////////////////////////////////////

#### func (*TextureUnit) Parse

```go
func (v *TextureUnit) Parse(s string) error
```

#### func (TextureUnit) String

```go
func (v TextureUnit) String() string
```

#### type TextureUnitːTextureTargetːTextureIdᵐᵐ

```go
type TextureUnitːTextureTargetːTextureIdᵐᵐ map[TextureUnit]TextureTargetːTextureIdᵐ
```


#### func (TextureUnitːTextureTargetːTextureIdᵐᵐ) Contains

```go
func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Contains(key TextureUnit) bool
```

#### func (TextureUnitːTextureTargetːTextureIdᵐᵐ) Delete

```go
func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Delete(key TextureUnit)
```

#### func (TextureUnitːTextureTargetːTextureIdᵐᵐ) Get

```go
func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Get(key TextureUnit) TextureTargetːTextureIdᵐ
```

#### func (TextureUnitːTextureTargetːTextureIdᵐᵐ) Range

```go
func (m TextureUnitːTextureTargetːTextureIdᵐᵐ) Range() []TextureTargetːTextureIdᵐ
```

#### type TextureWrapMode

```go
type TextureWrapMode uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TextureWrapMode
//////////////////////////////////////////////////////////////////////////////

#### func (*TextureWrapMode) Parse

```go
func (v *TextureWrapMode) Parse(s string) error
```

#### func (TextureWrapMode) String

```go
func (v TextureWrapMode) String() string
```

#### type ThreadID

```go
type ThreadID uint32
```


#### type ThreadIDːContextʳᵐ

```go
type ThreadIDːContextʳᵐ map[ThreadID](*Context)
```


#### func (ThreadIDːContextʳᵐ) Contains

```go
func (m ThreadIDːContextʳᵐ) Contains(key ThreadID) bool
```

#### func (ThreadIDːContextʳᵐ) Delete

```go
func (m ThreadIDːContextʳᵐ) Delete(key ThreadID)
```

#### func (ThreadIDːContextʳᵐ) Get

```go
func (m ThreadIDːContextʳᵐ) Get(key ThreadID) *Context
```

#### func (ThreadIDːContextʳᵐ) Range

```go
func (m ThreadIDːContextʳᵐ) Range() [](*Context)
```

#### type TilePreserveMaskQCOM

```go
type TilePreserveMaskQCOM uint32
```

//////////////////////////////////////////////////////////////////////////////
enum TilePreserveMaskQCOM
//////////////////////////////////////////////////////////////////////////////

#### func (*TilePreserveMaskQCOM) Parse

```go
func (v *TilePreserveMaskQCOM) Parse(s string) error
```

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

#### func (*Type_ARB_half_float_vertex) Parse

```go
func (v *Type_ARB_half_float_vertex) Parse(s string) error
```

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

#### func (*Type_OES_vertex_half_float) Parse

```go
func (v *Type_OES_vertex_half_float) Parse(s string) error
```

#### func (Type_OES_vertex_half_float) String

```go
func (v Type_OES_vertex_half_float) String() string
```

#### type U32ˢ

```go
type U32ˢ struct {
	binary.Generate
	SliceInfo
}
```

U32ˢ is a slice of uint32.

#### func  AsU32ˢ

```go
func AsU32ˢ(s Slice, ϟs *gfxapi.State) U32ˢ
```
AsU32ˢ returns s cast to a U32ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeU32ˢ

```go
func MakeU32ˢ(count uint64, ϟs *gfxapi.State) U32ˢ
```
MakeU32ˢ returns a U32ˢ backed by a new memory pool.

#### func (*U32ˢ) Class

```go
func (*U32ˢ) Class() binary.Class
```

#### func (U32ˢ) Clone

```go
func (s U32ˢ) Clone(ϟs *gfxapi.State) U32ˢ
```
Clone returns a copy of the U32ˢ in a new memory pool.

#### func (U32ˢ) Copy

```go
func (dst U32ˢ) Copy(src U32ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U32ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (U32ˢ) Decoder

```go
func (s U32ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (U32ˢ) ElementSize

```go
func (s U32ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U32ˢ points to.

#### func (U32ˢ) Encoder

```go
func (s U32ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (U32ˢ) Index

```go
func (s U32ˢ) Index(i uint64, ϟs *gfxapi.State) U32ᵖ
```
Index returns a U32ᵖ to the i'th element in this U32ˢ.

#### func (U32ˢ) OnRead

```go
func (s U32ˢ) OnRead(ϟs *gfxapi.State) U32ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (U32ˢ) OnWrite

```go
func (s U32ˢ) OnWrite(ϟs *gfxapi.State) U32ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (U32ˢ) Range

```go
func (s U32ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (U32ˢ) Read

```go
func (s U32ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint32
```
Read reads and returns all the uint32 elements in this U32ˢ.

#### func (U32ˢ) ResourceID

```go
func (s U32ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (U32ˢ) Slice

```go
func (s U32ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ
```
Slice returns a sub-slice from the U32ˢ using start and end indices.

#### func (U32ˢ) String

```go
func (s U32ˢ) String() string
```
String returns a string description of the U32ˢ slice.

#### func (U32ˢ) Write

```go
func (s U32ˢ) Write(src []uint32, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type U32ᵖ

```go
type U32ᵖ struct {
	binary.Generate
	memory.Pointer
}
```

U32ᵖ is a pointer to a uint32 element.

#### func  NewU32ᵖ

```go
func NewU32ᵖ(addr uint64) U32ᵖ
```
NewU32ᵖ returns a U32ᵖ that points to addr in the application pool.

#### func (*U32ᵖ) Class

```go
func (*U32ᵖ) Class() binary.Class
```

#### func (U32ᵖ) ElementSize

```go
func (p U32ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U32ᵖ points to.

#### func (U32ᵖ) OnRead

```go
func (p U32ᵖ) OnRead(ϟs *gfxapi.State) U32ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (U32ᵖ) OnWrite

```go
func (p U32ᵖ) OnWrite(ϟs *gfxapi.State) U32ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (U32ᵖ) Read

```go
func (p U32ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint32
```
Read reads and returns the uint32 element at the pointer.

#### func (U32ᵖ) Slice

```go
func (p U32ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U32ˢ
```
Slice returns a new U32ˢ from the pointer using start and end indices.

#### func (U32ᵖ) Write

```go
func (p U32ᵖ) Write(value uint32, ϟs *gfxapi.State)
```
Write writes value to the uint32 element at the pointer.

#### type U64ˢ

```go
type U64ˢ struct {
	binary.Generate
	SliceInfo
}
```

U64ˢ is a slice of uint64.

#### func  AsU64ˢ

```go
func AsU64ˢ(s Slice, ϟs *gfxapi.State) U64ˢ
```
AsU64ˢ returns s cast to a U64ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeU64ˢ

```go
func MakeU64ˢ(count uint64, ϟs *gfxapi.State) U64ˢ
```
MakeU64ˢ returns a U64ˢ backed by a new memory pool.

#### func (*U64ˢ) Class

```go
func (*U64ˢ) Class() binary.Class
```

#### func (U64ˢ) Clone

```go
func (s U64ˢ) Clone(ϟs *gfxapi.State) U64ˢ
```
Clone returns a copy of the U64ˢ in a new memory pool.

#### func (U64ˢ) Copy

```go
func (dst U64ˢ) Copy(src U64ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U64ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (U64ˢ) Decoder

```go
func (s U64ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (U64ˢ) ElementSize

```go
func (s U64ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U64ˢ points to.

#### func (U64ˢ) Encoder

```go
func (s U64ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (U64ˢ) Index

```go
func (s U64ˢ) Index(i uint64, ϟs *gfxapi.State) U64ᵖ
```
Index returns a U64ᵖ to the i'th element in this U64ˢ.

#### func (U64ˢ) OnRead

```go
func (s U64ˢ) OnRead(ϟs *gfxapi.State) U64ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (U64ˢ) OnWrite

```go
func (s U64ˢ) OnWrite(ϟs *gfxapi.State) U64ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (U64ˢ) Range

```go
func (s U64ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (U64ˢ) Read

```go
func (s U64ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint64
```
Read reads and returns all the uint64 elements in this U64ˢ.

#### func (U64ˢ) ResourceID

```go
func (s U64ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (U64ˢ) Slice

```go
func (s U64ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ
```
Slice returns a sub-slice from the U64ˢ using start and end indices.

#### func (U64ˢ) String

```go
func (s U64ˢ) String() string
```
String returns a string description of the U64ˢ slice.

#### func (U64ˢ) Write

```go
func (s U64ˢ) Write(src []uint64, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type U64ᵖ

```go
type U64ᵖ struct {
	binary.Generate
	memory.Pointer
}
```

U64ᵖ is a pointer to a uint64 element.

#### func  NewU64ᵖ

```go
func NewU64ᵖ(addr uint64) U64ᵖ
```
NewU64ᵖ returns a U64ᵖ that points to addr in the application pool.

#### func (*U64ᵖ) Class

```go
func (*U64ᵖ) Class() binary.Class
```

#### func (U64ᵖ) ElementSize

```go
func (p U64ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U64ᵖ points to.

#### func (U64ᵖ) OnRead

```go
func (p U64ᵖ) OnRead(ϟs *gfxapi.State) U64ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (U64ᵖ) OnWrite

```go
func (p U64ᵖ) OnWrite(ϟs *gfxapi.State) U64ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (U64ᵖ) Read

```go
func (p U64ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint64
```
Read reads and returns the uint64 element at the pointer.

#### func (U64ᵖ) Slice

```go
func (p U64ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U64ˢ
```
Slice returns a new U64ˢ from the pointer using start and end indices.

#### func (U64ᵖ) Write

```go
func (p U64ᵖ) Write(value uint64, ϟs *gfxapi.State)
```
Write writes value to the uint64 element at the pointer.

#### type U8ˢ

```go
type U8ˢ struct {
	binary.Generate
	SliceInfo
}
```

U8ˢ is a slice of uint8.

#### func  AsU8ˢ

```go
func AsU8ˢ(s Slice, ϟs *gfxapi.State) U8ˢ
```
AsU8ˢ returns s cast to a U8ˢ. The returned slice length will be calculated so
that the returned slice is no longer (in bytes) than s.

#### func  MakeU8ˢ

```go
func MakeU8ˢ(count uint64, ϟs *gfxapi.State) U8ˢ
```
MakeU8ˢ returns a U8ˢ backed by a new memory pool.

#### func (*U8ˢ) Class

```go
func (*U8ˢ) Class() binary.Class
```

#### func (U8ˢ) Clone

```go
func (s U8ˢ) Clone(ϟs *gfxapi.State) U8ˢ
```
Clone returns a copy of the U8ˢ in a new memory pool.

#### func (U8ˢ) Copy

```go
func (dst U8ˢ) Copy(src U8ˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s U8ˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (U8ˢ) Decoder

```go
func (s U8ˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (U8ˢ) ElementSize

```go
func (s U8ˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U8ˢ points to.

#### func (U8ˢ) Encoder

```go
func (s U8ˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (U8ˢ) Index

```go
func (s U8ˢ) Index(i uint64, ϟs *gfxapi.State) U8ᵖ
```
Index returns a U8ᵖ to the i'th element in this U8ˢ.

#### func (U8ˢ) OnRead

```go
func (s U8ˢ) OnRead(ϟs *gfxapi.State) U8ˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (U8ˢ) OnWrite

```go
func (s U8ˢ) OnWrite(ϟs *gfxapi.State) U8ˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (U8ˢ) Range

```go
func (s U8ˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (U8ˢ) Read

```go
func (s U8ˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []uint8
```
Read reads and returns all the uint8 elements in this U8ˢ.

#### func (U8ˢ) ResourceID

```go
func (s U8ˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (U8ˢ) Slice

```go
func (s U8ˢ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ
```
Slice returns a sub-slice from the U8ˢ using start and end indices.

#### func (U8ˢ) String

```go
func (s U8ˢ) String() string
```
String returns a string description of the U8ˢ slice.

#### func (U8ˢ) Write

```go
func (s U8ˢ) Write(src []uint8, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type U8ᵖ

```go
type U8ᵖ struct {
	binary.Generate
	memory.Pointer
}
```

U8ᵖ is a pointer to a uint8 element.

#### func  NewU8ᵖ

```go
func NewU8ᵖ(addr uint64) U8ᵖ
```
NewU8ᵖ returns a U8ᵖ that points to addr in the application pool.

#### func (*U8ᵖ) Class

```go
func (*U8ᵖ) Class() binary.Class
```

#### func (U8ᵖ) ElementSize

```go
func (p U8ᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that U8ᵖ points to.

#### func (U8ᵖ) OnRead

```go
func (p U8ᵖ) OnRead(ϟs *gfxapi.State) U8ᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (U8ᵖ) OnWrite

```go
func (p U8ᵖ) OnWrite(ϟs *gfxapi.State) U8ᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (U8ᵖ) Read

```go
func (p U8ᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) uint8
```
Read reads and returns the uint8 element at the pointer.

#### func (U8ᵖ) Slice

```go
func (p U8ᵖ) Slice(start, end uint64, ϟs *gfxapi.State) U8ˢ
```
Slice returns a new U8ˢ from the pointer using start and end indices.

#### func (U8ᵖ) Write

```go
func (p U8ᵖ) Write(value uint8, ϟs *gfxapi.State)
```
Write writes value to the uint8 element at the pointer.

#### type Uniform

```go
type Uniform struct {
	binary.Generate
	CreatedAt atom.ID
	Name      string
	Type      ShaderUniformType
	Value     U8ˢ
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

#### type UniformBlockParameter

```go
type UniformBlockParameter uint32
```

//////////////////////////////////////////////////////////////////////////////
enum UniformBlockParameter
//////////////////////////////////////////////////////////////////////////////

#### func (*UniformBlockParameter) Parse

```go
func (v *UniformBlockParameter) Parse(s string) error
```

#### func (UniformBlockParameter) String

```go
func (v UniformBlockParameter) String() string
```

#### type UniformLocation

```go
type UniformLocation int32
```


#### type UniformLocationːUniformᵐ

```go
type UniformLocationːUniformᵐ map[UniformLocation]Uniform
```


#### func (UniformLocationːUniformᵐ) Contains

```go
func (m UniformLocationːUniformᵐ) Contains(key UniformLocation) bool
```

#### func (UniformLocationːUniformᵐ) Delete

```go
func (m UniformLocationːUniformᵐ) Delete(key UniformLocation)
```

#### func (UniformLocationːUniformᵐ) Get

```go
func (m UniformLocationːUniformᵐ) Get(key UniformLocation) Uniform
```

#### func (UniformLocationːUniformᵐ) Range

```go
func (m UniformLocationːUniformᵐ) Range() []Uniform
```

#### type Vec2f

```go
type Vec2f F32ː2ᵃ
```


#### func (*Vec2f) Class

```go
func (*Vec2f) Class() binary.Class
```

#### func (Vec2f) String

```go
func (v Vec2f) String() string
```

#### type Vec2fː2ᵃ

```go
type Vec2fː2ᵃ struct {
	binary.Generate
	Elements [2]Vec2f
}
```


#### func (*Vec2fː2ᵃ) Class

```go
func (*Vec2fː2ᵃ) Class() binary.Class
```

#### type Vec2fˢ

```go
type Vec2fˢ struct {
	binary.Generate
	SliceInfo
}
```

Vec2fˢ is a slice of Vec2f.

#### func  AsVec2fˢ

```go
func AsVec2fˢ(s Slice, ϟs *gfxapi.State) Vec2fˢ
```
AsVec2fˢ returns s cast to a Vec2fˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeVec2fˢ

```go
func MakeVec2fˢ(count uint64, ϟs *gfxapi.State) Vec2fˢ
```
MakeVec2fˢ returns a Vec2fˢ backed by a new memory pool.

#### func (*Vec2fˢ) Class

```go
func (*Vec2fˢ) Class() binary.Class
```

#### func (Vec2fˢ) Clone

```go
func (s Vec2fˢ) Clone(ϟs *gfxapi.State) Vec2fˢ
```
Clone returns a copy of the Vec2fˢ in a new memory pool.

#### func (Vec2fˢ) Copy

```go
func (dst Vec2fˢ) Copy(src Vec2fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec2fˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Vec2fˢ) Decoder

```go
func (s Vec2fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Vec2fˢ) ElementSize

```go
func (s Vec2fˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec2fˢ points to.

#### func (Vec2fˢ) Encoder

```go
func (s Vec2fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Vec2fˢ) Index

```go
func (s Vec2fˢ) Index(i uint64, ϟs *gfxapi.State) Vec2fᵖ
```
Index returns a Vec2fᵖ to the i'th element in this Vec2fˢ.

#### func (Vec2fˢ) OnRead

```go
func (s Vec2fˢ) OnRead(ϟs *gfxapi.State) Vec2fˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Vec2fˢ) OnWrite

```go
func (s Vec2fˢ) OnWrite(ϟs *gfxapi.State) Vec2fˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Vec2fˢ) Range

```go
func (s Vec2fˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Vec2fˢ) Read

```go
func (s Vec2fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec2f
```
Read reads and returns all the Vec2f elements in this Vec2fˢ.

#### func (Vec2fˢ) ResourceID

```go
func (s Vec2fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Vec2fˢ) Slice

```go
func (s Vec2fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2fˢ
```
Slice returns a sub-slice from the Vec2fˢ using start and end indices.

#### func (Vec2fˢ) String

```go
func (s Vec2fˢ) String() string
```
String returns a string description of the Vec2fˢ slice.

#### func (Vec2fˢ) Write

```go
func (s Vec2fˢ) Write(src []Vec2f, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Vec2fᵖ

```go
type Vec2fᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Vec2fᵖ is a pointer to a Vec2f element.

#### func  NewVec2fᵖ

```go
func NewVec2fᵖ(addr uint64) Vec2fᵖ
```
NewVec2fᵖ returns a Vec2fᵖ that points to addr in the application pool.

#### func (*Vec2fᵖ) Class

```go
func (*Vec2fᵖ) Class() binary.Class
```

#### func (Vec2fᵖ) ElementSize

```go
func (p Vec2fᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec2fᵖ points to.

#### func (Vec2fᵖ) OnRead

```go
func (p Vec2fᵖ) OnRead(ϟs *gfxapi.State) Vec2fᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Vec2fᵖ) OnWrite

```go
func (p Vec2fᵖ) OnWrite(ϟs *gfxapi.State) Vec2fᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Vec2fᵖ) Read

```go
func (p Vec2fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec2f
```
Read reads and returns the Vec2f element at the pointer.

#### func (Vec2fᵖ) Slice

```go
func (p Vec2fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2fˢ
```
Slice returns a new Vec2fˢ from the pointer using start and end indices.

#### func (Vec2fᵖ) Write

```go
func (p Vec2fᵖ) Write(value Vec2f, ϟs *gfxapi.State)
```
Write writes value to the Vec2f element at the pointer.

#### type Vec2i

```go
type Vec2i S32ː2ᵃ
```


#### func (*Vec2i) Class

```go
func (*Vec2i) Class() binary.Class
```

#### func (Vec2i) String

```go
func (v Vec2i) String() string
```

#### type Vec2iˢ

```go
type Vec2iˢ struct {
	binary.Generate
	SliceInfo
}
```

Vec2iˢ is a slice of Vec2i.

#### func  AsVec2iˢ

```go
func AsVec2iˢ(s Slice, ϟs *gfxapi.State) Vec2iˢ
```
AsVec2iˢ returns s cast to a Vec2iˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeVec2iˢ

```go
func MakeVec2iˢ(count uint64, ϟs *gfxapi.State) Vec2iˢ
```
MakeVec2iˢ returns a Vec2iˢ backed by a new memory pool.

#### func (*Vec2iˢ) Class

```go
func (*Vec2iˢ) Class() binary.Class
```

#### func (Vec2iˢ) Clone

```go
func (s Vec2iˢ) Clone(ϟs *gfxapi.State) Vec2iˢ
```
Clone returns a copy of the Vec2iˢ in a new memory pool.

#### func (Vec2iˢ) Copy

```go
func (dst Vec2iˢ) Copy(src Vec2iˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec2iˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Vec2iˢ) Decoder

```go
func (s Vec2iˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Vec2iˢ) ElementSize

```go
func (s Vec2iˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec2iˢ points to.

#### func (Vec2iˢ) Encoder

```go
func (s Vec2iˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Vec2iˢ) Index

```go
func (s Vec2iˢ) Index(i uint64, ϟs *gfxapi.State) Vec2iᵖ
```
Index returns a Vec2iᵖ to the i'th element in this Vec2iˢ.

#### func (Vec2iˢ) OnRead

```go
func (s Vec2iˢ) OnRead(ϟs *gfxapi.State) Vec2iˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Vec2iˢ) OnWrite

```go
func (s Vec2iˢ) OnWrite(ϟs *gfxapi.State) Vec2iˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Vec2iˢ) Range

```go
func (s Vec2iˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Vec2iˢ) Read

```go
func (s Vec2iˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec2i
```
Read reads and returns all the Vec2i elements in this Vec2iˢ.

#### func (Vec2iˢ) ResourceID

```go
func (s Vec2iˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Vec2iˢ) Slice

```go
func (s Vec2iˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2iˢ
```
Slice returns a sub-slice from the Vec2iˢ using start and end indices.

#### func (Vec2iˢ) String

```go
func (s Vec2iˢ) String() string
```
String returns a string description of the Vec2iˢ slice.

#### func (Vec2iˢ) Write

```go
func (s Vec2iˢ) Write(src []Vec2i, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Vec2iᵖ

```go
type Vec2iᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Vec2iᵖ is a pointer to a Vec2i element.

#### func  NewVec2iᵖ

```go
func NewVec2iᵖ(addr uint64) Vec2iᵖ
```
NewVec2iᵖ returns a Vec2iᵖ that points to addr in the application pool.

#### func (*Vec2iᵖ) Class

```go
func (*Vec2iᵖ) Class() binary.Class
```

#### func (Vec2iᵖ) ElementSize

```go
func (p Vec2iᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec2iᵖ points to.

#### func (Vec2iᵖ) OnRead

```go
func (p Vec2iᵖ) OnRead(ϟs *gfxapi.State) Vec2iᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Vec2iᵖ) OnWrite

```go
func (p Vec2iᵖ) OnWrite(ϟs *gfxapi.State) Vec2iᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Vec2iᵖ) Read

```go
func (p Vec2iᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec2i
```
Read reads and returns the Vec2i element at the pointer.

#### func (Vec2iᵖ) Slice

```go
func (p Vec2iᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec2iˢ
```
Slice returns a new Vec2iˢ from the pointer using start and end indices.

#### func (Vec2iᵖ) Write

```go
func (p Vec2iᵖ) Write(value Vec2i, ϟs *gfxapi.State)
```
Write writes value to the Vec2i element at the pointer.

#### type Vec3f

```go
type Vec3f F32ː3ᵃ
```


#### func (*Vec3f) Class

```go
func (*Vec3f) Class() binary.Class
```

#### func (Vec3f) String

```go
func (v Vec3f) String() string
```

#### type Vec3fː3ᵃ

```go
type Vec3fː3ᵃ struct {
	binary.Generate
	Elements [3]Vec3f
}
```


#### func (*Vec3fː3ᵃ) Class

```go
func (*Vec3fː3ᵃ) Class() binary.Class
```

#### type Vec3fˢ

```go
type Vec3fˢ struct {
	binary.Generate
	SliceInfo
}
```

Vec3fˢ is a slice of Vec3f.

#### func  AsVec3fˢ

```go
func AsVec3fˢ(s Slice, ϟs *gfxapi.State) Vec3fˢ
```
AsVec3fˢ returns s cast to a Vec3fˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeVec3fˢ

```go
func MakeVec3fˢ(count uint64, ϟs *gfxapi.State) Vec3fˢ
```
MakeVec3fˢ returns a Vec3fˢ backed by a new memory pool.

#### func (*Vec3fˢ) Class

```go
func (*Vec3fˢ) Class() binary.Class
```

#### func (Vec3fˢ) Clone

```go
func (s Vec3fˢ) Clone(ϟs *gfxapi.State) Vec3fˢ
```
Clone returns a copy of the Vec3fˢ in a new memory pool.

#### func (Vec3fˢ) Copy

```go
func (dst Vec3fˢ) Copy(src Vec3fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec3fˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Vec3fˢ) Decoder

```go
func (s Vec3fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Vec3fˢ) ElementSize

```go
func (s Vec3fˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec3fˢ points to.

#### func (Vec3fˢ) Encoder

```go
func (s Vec3fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Vec3fˢ) Index

```go
func (s Vec3fˢ) Index(i uint64, ϟs *gfxapi.State) Vec3fᵖ
```
Index returns a Vec3fᵖ to the i'th element in this Vec3fˢ.

#### func (Vec3fˢ) OnRead

```go
func (s Vec3fˢ) OnRead(ϟs *gfxapi.State) Vec3fˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Vec3fˢ) OnWrite

```go
func (s Vec3fˢ) OnWrite(ϟs *gfxapi.State) Vec3fˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Vec3fˢ) Range

```go
func (s Vec3fˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Vec3fˢ) Read

```go
func (s Vec3fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec3f
```
Read reads and returns all the Vec3f elements in this Vec3fˢ.

#### func (Vec3fˢ) ResourceID

```go
func (s Vec3fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Vec3fˢ) Slice

```go
func (s Vec3fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3fˢ
```
Slice returns a sub-slice from the Vec3fˢ using start and end indices.

#### func (Vec3fˢ) String

```go
func (s Vec3fˢ) String() string
```
String returns a string description of the Vec3fˢ slice.

#### func (Vec3fˢ) Write

```go
func (s Vec3fˢ) Write(src []Vec3f, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Vec3fᵖ

```go
type Vec3fᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Vec3fᵖ is a pointer to a Vec3f element.

#### func  NewVec3fᵖ

```go
func NewVec3fᵖ(addr uint64) Vec3fᵖ
```
NewVec3fᵖ returns a Vec3fᵖ that points to addr in the application pool.

#### func (*Vec3fᵖ) Class

```go
func (*Vec3fᵖ) Class() binary.Class
```

#### func (Vec3fᵖ) ElementSize

```go
func (p Vec3fᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec3fᵖ points to.

#### func (Vec3fᵖ) OnRead

```go
func (p Vec3fᵖ) OnRead(ϟs *gfxapi.State) Vec3fᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Vec3fᵖ) OnWrite

```go
func (p Vec3fᵖ) OnWrite(ϟs *gfxapi.State) Vec3fᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Vec3fᵖ) Read

```go
func (p Vec3fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec3f
```
Read reads and returns the Vec3f element at the pointer.

#### func (Vec3fᵖ) Slice

```go
func (p Vec3fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3fˢ
```
Slice returns a new Vec3fˢ from the pointer using start and end indices.

#### func (Vec3fᵖ) Write

```go
func (p Vec3fᵖ) Write(value Vec3f, ϟs *gfxapi.State)
```
Write writes value to the Vec3f element at the pointer.

#### type Vec3i

```go
type Vec3i S32ː3ᵃ
```


#### func (*Vec3i) Class

```go
func (*Vec3i) Class() binary.Class
```

#### func (Vec3i) String

```go
func (v Vec3i) String() string
```

#### type Vec3iˢ

```go
type Vec3iˢ struct {
	binary.Generate
	SliceInfo
}
```

Vec3iˢ is a slice of Vec3i.

#### func  AsVec3iˢ

```go
func AsVec3iˢ(s Slice, ϟs *gfxapi.State) Vec3iˢ
```
AsVec3iˢ returns s cast to a Vec3iˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeVec3iˢ

```go
func MakeVec3iˢ(count uint64, ϟs *gfxapi.State) Vec3iˢ
```
MakeVec3iˢ returns a Vec3iˢ backed by a new memory pool.

#### func (*Vec3iˢ) Class

```go
func (*Vec3iˢ) Class() binary.Class
```

#### func (Vec3iˢ) Clone

```go
func (s Vec3iˢ) Clone(ϟs *gfxapi.State) Vec3iˢ
```
Clone returns a copy of the Vec3iˢ in a new memory pool.

#### func (Vec3iˢ) Copy

```go
func (dst Vec3iˢ) Copy(src Vec3iˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec3iˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Vec3iˢ) Decoder

```go
func (s Vec3iˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Vec3iˢ) ElementSize

```go
func (s Vec3iˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec3iˢ points to.

#### func (Vec3iˢ) Encoder

```go
func (s Vec3iˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Vec3iˢ) Index

```go
func (s Vec3iˢ) Index(i uint64, ϟs *gfxapi.State) Vec3iᵖ
```
Index returns a Vec3iᵖ to the i'th element in this Vec3iˢ.

#### func (Vec3iˢ) OnRead

```go
func (s Vec3iˢ) OnRead(ϟs *gfxapi.State) Vec3iˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Vec3iˢ) OnWrite

```go
func (s Vec3iˢ) OnWrite(ϟs *gfxapi.State) Vec3iˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Vec3iˢ) Range

```go
func (s Vec3iˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Vec3iˢ) Read

```go
func (s Vec3iˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec3i
```
Read reads and returns all the Vec3i elements in this Vec3iˢ.

#### func (Vec3iˢ) ResourceID

```go
func (s Vec3iˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Vec3iˢ) Slice

```go
func (s Vec3iˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3iˢ
```
Slice returns a sub-slice from the Vec3iˢ using start and end indices.

#### func (Vec3iˢ) String

```go
func (s Vec3iˢ) String() string
```
String returns a string description of the Vec3iˢ slice.

#### func (Vec3iˢ) Write

```go
func (s Vec3iˢ) Write(src []Vec3i, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Vec3iᵖ

```go
type Vec3iᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Vec3iᵖ is a pointer to a Vec3i element.

#### func  NewVec3iᵖ

```go
func NewVec3iᵖ(addr uint64) Vec3iᵖ
```
NewVec3iᵖ returns a Vec3iᵖ that points to addr in the application pool.

#### func (*Vec3iᵖ) Class

```go
func (*Vec3iᵖ) Class() binary.Class
```

#### func (Vec3iᵖ) ElementSize

```go
func (p Vec3iᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec3iᵖ points to.

#### func (Vec3iᵖ) OnRead

```go
func (p Vec3iᵖ) OnRead(ϟs *gfxapi.State) Vec3iᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Vec3iᵖ) OnWrite

```go
func (p Vec3iᵖ) OnWrite(ϟs *gfxapi.State) Vec3iᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Vec3iᵖ) Read

```go
func (p Vec3iᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec3i
```
Read reads and returns the Vec3i element at the pointer.

#### func (Vec3iᵖ) Slice

```go
func (p Vec3iᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec3iˢ
```
Slice returns a new Vec3iˢ from the pointer using start and end indices.

#### func (Vec3iᵖ) Write

```go
func (p Vec3iᵖ) Write(value Vec3i, ϟs *gfxapi.State)
```
Write writes value to the Vec3i element at the pointer.

#### type Vec4f

```go
type Vec4f F32ː4ᵃ
```


#### func (*Vec4f) Class

```go
func (*Vec4f) Class() binary.Class
```

#### func (Vec4f) String

```go
func (v Vec4f) String() string
```

#### type Vec4fː4ᵃ

```go
type Vec4fː4ᵃ struct {
	binary.Generate
	Elements [4]Vec4f
}
```


#### func (*Vec4fː4ᵃ) Class

```go
func (*Vec4fː4ᵃ) Class() binary.Class
```

#### type Vec4fˢ

```go
type Vec4fˢ struct {
	binary.Generate
	SliceInfo
}
```

Vec4fˢ is a slice of Vec4f.

#### func  AsVec4fˢ

```go
func AsVec4fˢ(s Slice, ϟs *gfxapi.State) Vec4fˢ
```
AsVec4fˢ returns s cast to a Vec4fˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeVec4fˢ

```go
func MakeVec4fˢ(count uint64, ϟs *gfxapi.State) Vec4fˢ
```
MakeVec4fˢ returns a Vec4fˢ backed by a new memory pool.

#### func (*Vec4fˢ) Class

```go
func (*Vec4fˢ) Class() binary.Class
```

#### func (Vec4fˢ) Clone

```go
func (s Vec4fˢ) Clone(ϟs *gfxapi.State) Vec4fˢ
```
Clone returns a copy of the Vec4fˢ in a new memory pool.

#### func (Vec4fˢ) Copy

```go
func (dst Vec4fˢ) Copy(src Vec4fˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec4fˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Vec4fˢ) Decoder

```go
func (s Vec4fˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Vec4fˢ) ElementSize

```go
func (s Vec4fˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec4fˢ points to.

#### func (Vec4fˢ) Encoder

```go
func (s Vec4fˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Vec4fˢ) Index

```go
func (s Vec4fˢ) Index(i uint64, ϟs *gfxapi.State) Vec4fᵖ
```
Index returns a Vec4fᵖ to the i'th element in this Vec4fˢ.

#### func (Vec4fˢ) OnRead

```go
func (s Vec4fˢ) OnRead(ϟs *gfxapi.State) Vec4fˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Vec4fˢ) OnWrite

```go
func (s Vec4fˢ) OnWrite(ϟs *gfxapi.State) Vec4fˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Vec4fˢ) Range

```go
func (s Vec4fˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Vec4fˢ) Read

```go
func (s Vec4fˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec4f
```
Read reads and returns all the Vec4f elements in this Vec4fˢ.

#### func (Vec4fˢ) ResourceID

```go
func (s Vec4fˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Vec4fˢ) Slice

```go
func (s Vec4fˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4fˢ
```
Slice returns a sub-slice from the Vec4fˢ using start and end indices.

#### func (Vec4fˢ) String

```go
func (s Vec4fˢ) String() string
```
String returns a string description of the Vec4fˢ slice.

#### func (Vec4fˢ) Write

```go
func (s Vec4fˢ) Write(src []Vec4f, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Vec4fᵖ

```go
type Vec4fᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Vec4fᵖ is a pointer to a Vec4f element.

#### func  NewVec4fᵖ

```go
func NewVec4fᵖ(addr uint64) Vec4fᵖ
```
NewVec4fᵖ returns a Vec4fᵖ that points to addr in the application pool.

#### func (*Vec4fᵖ) Class

```go
func (*Vec4fᵖ) Class() binary.Class
```

#### func (Vec4fᵖ) ElementSize

```go
func (p Vec4fᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec4fᵖ points to.

#### func (Vec4fᵖ) OnRead

```go
func (p Vec4fᵖ) OnRead(ϟs *gfxapi.State) Vec4fᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Vec4fᵖ) OnWrite

```go
func (p Vec4fᵖ) OnWrite(ϟs *gfxapi.State) Vec4fᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Vec4fᵖ) Read

```go
func (p Vec4fᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec4f
```
Read reads and returns the Vec4f element at the pointer.

#### func (Vec4fᵖ) Slice

```go
func (p Vec4fᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4fˢ
```
Slice returns a new Vec4fˢ from the pointer using start and end indices.

#### func (Vec4fᵖ) Write

```go
func (p Vec4fᵖ) Write(value Vec4f, ϟs *gfxapi.State)
```
Write writes value to the Vec4f element at the pointer.

#### type Vec4i

```go
type Vec4i S32ː4ᵃ
```


#### func (*Vec4i) Class

```go
func (*Vec4i) Class() binary.Class
```

#### func (Vec4i) String

```go
func (v Vec4i) String() string
```

#### type Vec4iˢ

```go
type Vec4iˢ struct {
	binary.Generate
	SliceInfo
}
```

Vec4iˢ is a slice of Vec4i.

#### func  AsVec4iˢ

```go
func AsVec4iˢ(s Slice, ϟs *gfxapi.State) Vec4iˢ
```
AsVec4iˢ returns s cast to a Vec4iˢ. The returned slice length will be
calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeVec4iˢ

```go
func MakeVec4iˢ(count uint64, ϟs *gfxapi.State) Vec4iˢ
```
MakeVec4iˢ returns a Vec4iˢ backed by a new memory pool.

#### func (*Vec4iˢ) Class

```go
func (*Vec4iˢ) Class() binary.Class
```

#### func (Vec4iˢ) Clone

```go
func (s Vec4iˢ) Clone(ϟs *gfxapi.State) Vec4iˢ
```
Clone returns a copy of the Vec4iˢ in a new memory pool.

#### func (Vec4iˢ) Copy

```go
func (dst Vec4iˢ) Copy(src Vec4iˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s Vec4iˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (Vec4iˢ) Decoder

```go
func (s Vec4iˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Vec4iˢ) ElementSize

```go
func (s Vec4iˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec4iˢ points to.

#### func (Vec4iˢ) Encoder

```go
func (s Vec4iˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Vec4iˢ) Index

```go
func (s Vec4iˢ) Index(i uint64, ϟs *gfxapi.State) Vec4iᵖ
```
Index returns a Vec4iᵖ to the i'th element in this Vec4iˢ.

#### func (Vec4iˢ) OnRead

```go
func (s Vec4iˢ) OnRead(ϟs *gfxapi.State) Vec4iˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Vec4iˢ) OnWrite

```go
func (s Vec4iˢ) OnWrite(ϟs *gfxapi.State) Vec4iˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Vec4iˢ) Range

```go
func (s Vec4iˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Vec4iˢ) Read

```go
func (s Vec4iˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []Vec4i
```
Read reads and returns all the Vec4i elements in this Vec4iˢ.

#### func (Vec4iˢ) ResourceID

```go
func (s Vec4iˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Vec4iˢ) Slice

```go
func (s Vec4iˢ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4iˢ
```
Slice returns a sub-slice from the Vec4iˢ using start and end indices.

#### func (Vec4iˢ) String

```go
func (s Vec4iˢ) String() string
```
String returns a string description of the Vec4iˢ slice.

#### func (Vec4iˢ) Write

```go
func (s Vec4iˢ) Write(src []Vec4i, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type Vec4iᵖ

```go
type Vec4iᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Vec4iᵖ is a pointer to a Vec4i element.

#### func  NewVec4iᵖ

```go
func NewVec4iᵖ(addr uint64) Vec4iᵖ
```
NewVec4iᵖ returns a Vec4iᵖ that points to addr in the application pool.

#### func (*Vec4iᵖ) Class

```go
func (*Vec4iᵖ) Class() binary.Class
```

#### func (Vec4iᵖ) ElementSize

```go
func (p Vec4iᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Vec4iᵖ points to.

#### func (Vec4iᵖ) OnRead

```go
func (p Vec4iᵖ) OnRead(ϟs *gfxapi.State) Vec4iᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Vec4iᵖ) OnWrite

```go
func (p Vec4iᵖ) OnWrite(ϟs *gfxapi.State) Vec4iᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Vec4iᵖ) Read

```go
func (p Vec4iᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) Vec4i
```
Read reads and returns the Vec4i element at the pointer.

#### func (Vec4iᵖ) Slice

```go
func (p Vec4iᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Vec4iˢ
```
Slice returns a new Vec4iˢ from the pointer using start and end indices.

#### func (Vec4iᵖ) Write

```go
func (p Vec4iᵖ) Write(value Vec4i, ϟs *gfxapi.State)
```
Write writes value to the Vec4i element at the pointer.

#### type Version

```go
type Version struct {
	IsES  bool
	Major int
	Minor int
}
```

Version represents the GL version major and minor numbers, and whether its
flavour is ES, as opposed to Desktop GL.

#### func  ParseVersion

```go
func ParseVersion(str string) (*Version, error)
```
ParseVersion parses the GL version major, minor and flavour from the output of
glGetString(GL_VERSION).

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

#### type VertexArrayId

```go
type VertexArrayId uint32
```


#### type VertexArrayIdːVertexArrayʳᵐ

```go
type VertexArrayIdːVertexArrayʳᵐ map[VertexArrayId](*VertexArray)
```


#### func (VertexArrayIdːVertexArrayʳᵐ) Contains

```go
func (m VertexArrayIdːVertexArrayʳᵐ) Contains(key VertexArrayId) bool
```

#### func (VertexArrayIdːVertexArrayʳᵐ) Delete

```go
func (m VertexArrayIdːVertexArrayʳᵐ) Delete(key VertexArrayId)
```

#### func (VertexArrayIdːVertexArrayʳᵐ) Get

```go
func (m VertexArrayIdːVertexArrayʳᵐ) Get(key VertexArrayId) *VertexArray
```

#### func (VertexArrayIdːVertexArrayʳᵐ) Range

```go
func (m VertexArrayIdːVertexArrayʳᵐ) Range() [](*VertexArray)
```

#### type VertexArrayIdˢ

```go
type VertexArrayIdˢ struct {
	binary.Generate
	SliceInfo
}
```

VertexArrayIdˢ is a slice of VertexArrayId.

#### func  AsVertexArrayIdˢ

```go
func AsVertexArrayIdˢ(s Slice, ϟs *gfxapi.State) VertexArrayIdˢ
```
AsVertexArrayIdˢ returns s cast to a VertexArrayIdˢ. The returned slice length
will be calculated so that the returned slice is no longer (in bytes) than s.

#### func  MakeVertexArrayIdˢ

```go
func MakeVertexArrayIdˢ(count uint64, ϟs *gfxapi.State) VertexArrayIdˢ
```
MakeVertexArrayIdˢ returns a VertexArrayIdˢ backed by a new memory pool.

#### func (*VertexArrayIdˢ) Class

```go
func (*VertexArrayIdˢ) Class() binary.Class
```

#### func (VertexArrayIdˢ) Clone

```go
func (s VertexArrayIdˢ) Clone(ϟs *gfxapi.State) VertexArrayIdˢ
```
Clone returns a copy of the VertexArrayIdˢ in a new memory pool.

#### func (VertexArrayIdˢ) Copy

```go
func (dst VertexArrayIdˢ) Copy(src VertexArrayIdˢ, ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) (d, s VertexArrayIdˢ)
```
Copy copies elements from src to this slice. The number of elements copied is
the minimum of dst.Count and src.Count. The slices of this and dst to the copied
elements is returned.

#### func (VertexArrayIdˢ) Decoder

```go
func (s VertexArrayIdˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (VertexArrayIdˢ) ElementSize

```go
func (s VertexArrayIdˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that VertexArrayIdˢ points
to.

#### func (VertexArrayIdˢ) Encoder

```go
func (s VertexArrayIdˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (VertexArrayIdˢ) Index

```go
func (s VertexArrayIdˢ) Index(i uint64, ϟs *gfxapi.State) VertexArrayIdᵖ
```
Index returns a VertexArrayIdᵖ to the i'th element in this VertexArrayIdˢ.

#### func (VertexArrayIdˢ) OnRead

```go
func (s VertexArrayIdˢ) OnRead(ϟs *gfxapi.State) VertexArrayIdˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (VertexArrayIdˢ) OnWrite

```go
func (s VertexArrayIdˢ) OnWrite(ϟs *gfxapi.State) VertexArrayIdˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (VertexArrayIdˢ) Range

```go
func (s VertexArrayIdˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (VertexArrayIdˢ) Read

```go
func (s VertexArrayIdˢ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) []VertexArrayId
```
Read reads and returns all the VertexArrayId elements in this VertexArrayIdˢ.

#### func (VertexArrayIdˢ) ResourceID

```go
func (s VertexArrayIdˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (VertexArrayIdˢ) Slice

```go
func (s VertexArrayIdˢ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ
```
Slice returns a sub-slice from the VertexArrayIdˢ using start and end indices.

#### func (VertexArrayIdˢ) String

```go
func (s VertexArrayIdˢ) String() string
```
String returns a string description of the VertexArrayIdˢ slice.

#### func (VertexArrayIdˢ) Write

```go
func (s VertexArrayIdˢ) Write(src []VertexArrayId, ϟs *gfxapi.State) uint64
```
Write copies elements from src to this slice. The number of elements copied is
returned which is the minimum of s.Count and len(src).

#### type VertexArrayIdᵖ

```go
type VertexArrayIdᵖ struct {
	binary.Generate
	memory.Pointer
}
```

VertexArrayIdᵖ is a pointer to a VertexArrayId element.

#### func  NewVertexArrayIdᵖ

```go
func NewVertexArrayIdᵖ(addr uint64) VertexArrayIdᵖ
```
NewVertexArrayIdᵖ returns a VertexArrayIdᵖ that points to addr in the
application pool.

#### func (*VertexArrayIdᵖ) Class

```go
func (*VertexArrayIdᵖ) Class() binary.Class
```

#### func (VertexArrayIdᵖ) ElementSize

```go
func (p VertexArrayIdᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that VertexArrayIdᵖ points
to.

#### func (VertexArrayIdᵖ) OnRead

```go
func (p VertexArrayIdᵖ) OnRead(ϟs *gfxapi.State) VertexArrayIdᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (VertexArrayIdᵖ) OnWrite

```go
func (p VertexArrayIdᵖ) OnWrite(ϟs *gfxapi.State) VertexArrayIdᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (VertexArrayIdᵖ) Read

```go
func (p VertexArrayIdᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) VertexArrayId
```
Read reads and returns the VertexArrayId element at the pointer.

#### func (VertexArrayIdᵖ) Slice

```go
func (p VertexArrayIdᵖ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ
```
Slice returns a new VertexArrayIdˢ from the pointer using start and end indices.

#### func (VertexArrayIdᵖ) Write

```go
func (p VertexArrayIdᵖ) Write(value VertexArrayId, ϟs *gfxapi.State)
```
Write writes value to the VertexArrayId element at the pointer.

#### type VertexArrayIdᶜᵖ

```go
type VertexArrayIdᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

VertexArrayIdᶜᵖ is a pointer to a VertexArrayId element.

#### func  NewVertexArrayIdᶜᵖ

```go
func NewVertexArrayIdᶜᵖ(addr uint64) VertexArrayIdᶜᵖ
```
NewVertexArrayIdᶜᵖ returns a VertexArrayIdᶜᵖ that points to addr in the
application pool.

#### func (*VertexArrayIdᶜᵖ) Class

```go
func (*VertexArrayIdᶜᵖ) Class() binary.Class
```

#### func (VertexArrayIdᶜᵖ) ElementSize

```go
func (p VertexArrayIdᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that VertexArrayIdᶜᵖ points
to.

#### func (VertexArrayIdᶜᵖ) OnRead

```go
func (p VertexArrayIdᶜᵖ) OnRead(ϟs *gfxapi.State) VertexArrayIdᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (VertexArrayIdᶜᵖ) OnWrite

```go
func (p VertexArrayIdᶜᵖ) OnWrite(ϟs *gfxapi.State) VertexArrayIdᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (VertexArrayIdᶜᵖ) Read

```go
func (p VertexArrayIdᶜᵖ) Read(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) VertexArrayId
```
Read reads and returns the VertexArrayId element at the pointer.

#### func (VertexArrayIdᶜᵖ) Slice

```go
func (p VertexArrayIdᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) VertexArrayIdˢ
```
Slice returns a new VertexArrayIdˢ from the pointer using start and end indices.

#### func (VertexArrayIdᶜᵖ) Write

```go
func (p VertexArrayIdᶜᵖ) Write(value VertexArrayId, ϟs *gfxapi.State)
```
Write writes value to the VertexArrayId element at the pointer.

#### type VertexAttribType

```go
type VertexAttribType uint32
```

//////////////////////////////////////////////////////////////////////////////
enum VertexAttribType
//////////////////////////////////////////////////////////////////////////////

#### func (*VertexAttribType) Parse

```go
func (v *VertexAttribType) Parse(s string) error
```

#### func (VertexAttribType) String

```go
func (v VertexAttribType) String() string
```

#### type VertexAttribute

```go
type VertexAttribute struct {
	binary.Generate
	CreatedAt   atom.ID
	Name        Charˢ
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
	Pointer    VertexPointer
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

#### type VertexPointer

```go
type VertexPointer struct {
	binary.Generate
	memory.Pointer
}
```

VertexPointer is a pointer to a void element.

#### func  NewVertexPointer

```go
func NewVertexPointer(addr uint64) VertexPointer
```
NewVertexPointer returns a VertexPointer that points to addr in the application
pool.

#### func (*VertexPointer) Class

```go
func (*VertexPointer) Class() binary.Class
```

#### func (VertexPointer) ElementSize

```go
func (p VertexPointer) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that VertexPointer points
to.

#### func (VertexPointer) OnRead

```go
func (p VertexPointer) OnRead(ϟs *gfxapi.State) VertexPointer
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (VertexPointer) OnWrite

```go
func (p VertexPointer) OnWrite(ϟs *gfxapi.State) VertexPointer
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (VertexPointer) Slice

```go
func (p VertexPointer) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type Voidˢ

```go
type Voidˢ struct {
	binary.Generate
	SliceInfo
}
```

Voidˢ is a slice of void.

#### func  MakeVoidˢ

```go
func MakeVoidˢ(count uint64, ϟs *gfxapi.State) Voidˢ
```
MakeVoidˢ returns a Voidˢ backed by a new memory pool.

#### func (*Voidˢ) Class

```go
func (*Voidˢ) Class() binary.Class
```

#### func (Voidˢ) Clone

```go
func (s Voidˢ) Clone(ϟs *gfxapi.State) Voidˢ
```
Clone returns a copy of the Voidˢ in a new memory pool.

#### func (Voidˢ) Decoder

```go
func (s Voidˢ) Decoder(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.Decoder
```
Decoder returns a memory decoder for the slice.

#### func (Voidˢ) ElementSize

```go
func (s Voidˢ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Voidˢ points to.

#### func (Voidˢ) Encoder

```go
func (s Voidˢ) Encoder(ϟs *gfxapi.State) binary.Encoder
```
Encoder returns a memory encoder for the slice.

#### func (Voidˢ) Index

```go
func (s Voidˢ) Index(i uint64, ϟs *gfxapi.State) Voidᵖ
```
Index returns a Voidᵖ to the i'th element in this Voidˢ.

#### func (Voidˢ) OnRead

```go
func (s Voidˢ) OnRead(ϟs *gfxapi.State) Voidˢ
```
OnRead calls the backing pool's OnRead callback. s is returned so calls can be
chained.

#### func (Voidˢ) OnWrite

```go
func (s Voidˢ) OnWrite(ϟs *gfxapi.State) Voidˢ
```
OnWrite calls the backing pool's OnWrite callback. s is returned so calls can be
chained.

#### func (Voidˢ) Range

```go
func (s Voidˢ) Range(ϟs *gfxapi.State) memory.Range
```
Range returns the memory range this slice represents in the underlying pool.

#### func (Voidˢ) ResourceID

```go
func (s Voidˢ) ResourceID(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) binary.ID
```
ResourceID returns an identifier to a resource representing the data of this
slice.

#### func (Voidˢ) Slice

```go
func (s Voidˢ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a sub-slice from the Voidˢ using start and end indices.

#### func (Voidˢ) String

```go
func (s Voidˢ) String() string
```
String returns a string description of the Voidˢ slice.

#### type Voidᵖ

```go
type Voidᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Voidᵖ is a pointer to a void element.

#### func  NewVoidᵖ

```go
func NewVoidᵖ(addr uint64) Voidᵖ
```
NewVoidᵖ returns a Voidᵖ that points to addr in the application pool.

#### func (*Voidᵖ) Class

```go
func (*Voidᵖ) Class() binary.Class
```

#### func (Voidᵖ) ElementSize

```go
func (p Voidᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Voidᵖ points to.

#### func (Voidᵖ) OnRead

```go
func (p Voidᵖ) OnRead(ϟs *gfxapi.State) Voidᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Voidᵖ) OnWrite

```go
func (p Voidᵖ) OnWrite(ϟs *gfxapi.State) Voidᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Voidᵖ) Slice

```go
func (p Voidᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

#### type Voidᶜᵖ

```go
type Voidᶜᵖ struct {
	binary.Generate
	memory.Pointer
}
```

Voidᶜᵖ is a pointer to a void element.

#### func  NewVoidᶜᵖ

```go
func NewVoidᶜᵖ(addr uint64) Voidᶜᵖ
```
NewVoidᶜᵖ returns a Voidᶜᵖ that points to addr in the application pool.

#### func (*Voidᶜᵖ) Class

```go
func (*Voidᶜᵖ) Class() binary.Class
```

#### func (Voidᶜᵖ) ElementSize

```go
func (p Voidᶜᵖ) ElementSize(ϟs *gfxapi.State) uint64
```
ElementSize returns the size in bytes of an element that Voidᶜᵖ points to.

#### func (Voidᶜᵖ) OnRead

```go
func (p Voidᶜᵖ) OnRead(ϟs *gfxapi.State) Voidᶜᵖ
```
OnRead calls the backing pool's OnRead callback. p is returned so calls can be
chained.

#### func (Voidᶜᵖ) OnWrite

```go
func (p Voidᶜᵖ) OnWrite(ϟs *gfxapi.State) Voidᶜᵖ
```
OnWrite calls the backing pool's OnWrite callback. p is returned so calls can be
chained.

#### func (Voidᶜᵖ) Slice

```go
func (p Voidᶜᵖ) Slice(start, end uint64, ϟs *gfxapi.State) Voidˢ
```
Slice returns a new Voidˢ from the pointer using start and end indices.

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
func NewWglCreateContext(Hdc memory.Pointer, Result memory.Pointer) *WglCreateContext
```

#### func (*WglCreateContext) API

```go
func (c *WglCreateContext) API() gfxapi.ID
```

#### func (*WglCreateContext) AddRead

```go
func (a *WglCreateContext) AddRead(rng memory.Range, id binary.ID) *WglCreateContext
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The WglCreateContext pointer is returned so that calls can be chained.

#### func (*WglCreateContext) AddWrite

```go
func (a *WglCreateContext) AddWrite(rng memory.Range, id binary.ID) *WglCreateContext
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The WglCreateContext pointer is returned so that calls can be chained.

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
func (ϟa *WglCreateContext) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*WglCreateContext) Observations

```go
func (a *WglCreateContext) Observations() *atom.Observations
```

#### func (*WglCreateContext) Replay

```go
func (ω *WglCreateContext) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*WglCreateContext) String

```go
func (a *WglCreateContext) String() string
```

#### type WglCreateContextAttribsARB

```go
type WglCreateContextAttribsARB struct {
	binary.Generate

	Hdc           HDC
	HShareContext HGLRC
	AttribList    Intᵖ
	Result        HGLRC
}
```

//////////////////////////////////////////////////////////////////////////////
WglCreateContextAttribsARB
//////////////////////////////////////////////////////////////////////////////

#### func  NewWglCreateContextAttribsARB

```go
func NewWglCreateContextAttribsARB(Hdc memory.Pointer, HShareContext memory.Pointer, AttribList memory.Pointer, Result memory.Pointer) *WglCreateContextAttribsARB
```

#### func (*WglCreateContextAttribsARB) API

```go
func (c *WglCreateContextAttribsARB) API() gfxapi.ID
```

#### func (*WglCreateContextAttribsARB) AddRead

```go
func (a *WglCreateContextAttribsARB) AddRead(rng memory.Range, id binary.ID) *WglCreateContextAttribsARB
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The WglCreateContextAttribsARB pointer is returned so that calls can be
chained.

#### func (*WglCreateContextAttribsARB) AddWrite

```go
func (a *WglCreateContextAttribsARB) AddWrite(rng memory.Range, id binary.ID) *WglCreateContextAttribsARB
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The WglCreateContextAttribsARB pointer is returned so that calls can be
chained.

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
func (ϟa *WglCreateContextAttribsARB) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*WglCreateContextAttribsARB) Observations

```go
func (a *WglCreateContextAttribsARB) Observations() *atom.Observations
```

#### func (*WglCreateContextAttribsARB) Replay

```go
func (ω *WglCreateContextAttribsARB) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*WglCreateContextAttribsARB) String

```go
func (a *WglCreateContextAttribsARB) String() string
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
func NewWglMakeCurrent(Hdc memory.Pointer, Hglrc memory.Pointer, Result BOOL) *WglMakeCurrent
```

#### func (*WglMakeCurrent) API

```go
func (c *WglMakeCurrent) API() gfxapi.ID
```

#### func (*WglMakeCurrent) AddRead

```go
func (a *WglMakeCurrent) AddRead(rng memory.Range, id binary.ID) *WglMakeCurrent
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The WglMakeCurrent pointer is returned so that calls can be chained.

#### func (*WglMakeCurrent) AddWrite

```go
func (a *WglMakeCurrent) AddWrite(rng memory.Range, id binary.ID) *WglMakeCurrent
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The WglMakeCurrent pointer is returned so that calls can be chained.

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
func (ϟa *WglMakeCurrent) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*WglMakeCurrent) Observations

```go
func (a *WglMakeCurrent) Observations() *atom.Observations
```

#### func (*WglMakeCurrent) Replay

```go
func (ω *WglMakeCurrent) Replay(i atom.ID, s *gfxapi.State, d database.Database, l log.Logger, b *builder.Builder) error
```

#### func (*WglMakeCurrent) String

```go
func (a *WglMakeCurrent) String() string
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
func NewWglSwapBuffers(Hdc memory.Pointer) *WglSwapBuffers
```

#### func (*WglSwapBuffers) API

```go
func (c *WglSwapBuffers) API() gfxapi.ID
```

#### func (*WglSwapBuffers) AddRead

```go
func (a *WglSwapBuffers) AddRead(rng memory.Range, id binary.ID) *WglSwapBuffers
```
AddRead appends a new read observation to the atom of the range rng with the
data id. The WglSwapBuffers pointer is returned so that calls can be chained.

#### func (*WglSwapBuffers) AddWrite

```go
func (a *WglSwapBuffers) AddWrite(rng memory.Range, id binary.ID) *WglSwapBuffers
```
AddWrite appends a new write observation to the atom of the range rng with the
data id. The WglSwapBuffers pointer is returned so that calls can be chained.

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
func (ϟa *WglSwapBuffers) Mutate(ϟs *gfxapi.State, ϟd database.Database, ϟl log.Logger) error
```

#### func (*WglSwapBuffers) Observations

```go
func (a *WglSwapBuffers) Observations() *atom.Observations
```

#### func (*WglSwapBuffers) String

```go
func (a *WglSwapBuffers) String() string
```
