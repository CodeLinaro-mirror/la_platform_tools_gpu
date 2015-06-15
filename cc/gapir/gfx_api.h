////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

#ifndef GAPIR_GFX_API_H
#define GAPIR_GFX_API_H

#include <gapic/target.h>

#include <stdint.h>

namespace gapir {

class Interpreter;

namespace gfxapi {

// Register the API functions with the interpreter
extern void Register(Interpreter *interpreter);

// Look-up the API function addresses for the currently bound context.
extern void Initialize();

namespace Ids {
// List of the function ids for the API functions. This list have to be consistent with the
// function ids on the server side because they are part of the communication protocol
static const uint16_t EglInitialize = 0;
static const uint16_t EglCreateContext = 1;
static const uint16_t EglMakeCurrent = 2;
static const uint16_t EglSwapBuffers = 3;
static const uint16_t EglQuerySurface = 4;
static const uint16_t GlXCreateContext = 5;
static const uint16_t GlXCreateNewContext = 6;
static const uint16_t GlXMakeContextCurrent = 7;
static const uint16_t GlXMakeCurrent = 8;
static const uint16_t GlXSwapBuffers = 9;
static const uint16_t GlXQueryDrawable = 10;
static const uint16_t WglCreateContext = 11;
static const uint16_t WglCreateContextAttribsARB = 12;
static const uint16_t WglMakeCurrent = 13;
static const uint16_t WglSwapBuffers = 14;
static const uint16_t CGLCreateContext = 15;
static const uint16_t CGLSetCurrentContext = 16;
static const uint16_t CGLGetSurface = 17;
static const uint16_t CGSGetSurfaceBounds = 18;
static const uint16_t CGLFlushDrawable = 19;
static const uint16_t GlEnableClientState = 20;
static const uint16_t GlDisableClientState = 21;
static const uint16_t GlGetProgramBinaryOES = 22;
static const uint16_t GlProgramBinaryOES = 23;
static const uint16_t GlStartTilingQCOM = 24;
static const uint16_t GlEndTilingQCOM = 25;
static const uint16_t GlDiscardFramebufferEXT = 26;
static const uint16_t GlInsertEventMarkerEXT = 27;
static const uint16_t GlPushGroupMarkerEXT = 28;
static const uint16_t GlPopGroupMarkerEXT = 29;
static const uint16_t GlTexStorage1DEXT = 30;
static const uint16_t GlTexStorage2DEXT = 31;
static const uint16_t GlTexStorage3DEXT = 32;
static const uint16_t GlTextureStorage1DEXT = 33;
static const uint16_t GlTextureStorage2DEXT = 34;
static const uint16_t GlTextureStorage3DEXT = 35;
static const uint16_t GlGenVertexArraysOES = 36;
static const uint16_t GlBindVertexArrayOES = 37;
static const uint16_t GlDeleteVertexArraysOES = 38;
static const uint16_t GlIsVertexArrayOES = 39;
static const uint16_t GlEGLImageTargetTexture2DOES = 40;
static const uint16_t GlEGLImageTargetRenderbufferStorageOES = 41;
static const uint16_t GlGetGraphicsResetStatusEXT = 42;
static const uint16_t GlBindAttribLocation = 43;
static const uint16_t GlBlendFunc = 44;
static const uint16_t GlBlendFuncSeparate = 45;
static const uint16_t GlBlendEquation = 46;
static const uint16_t GlBlendEquationSeparate = 47;
static const uint16_t GlBlendColor = 48;
static const uint16_t GlEnableVertexAttribArray = 49;
static const uint16_t GlDisableVertexAttribArray = 50;
static const uint16_t GlVertexAttribPointer = 51;
static const uint16_t GlGetActiveAttrib = 52;
static const uint16_t GlGetActiveUniform = 53;
static const uint16_t GlGetError = 54;
static const uint16_t GlGetProgramiv = 55;
static const uint16_t GlGetShaderiv = 56;
static const uint16_t GlGetUniformLocation = 57;
static const uint16_t GlGetAttribLocation = 58;
static const uint16_t GlPixelStorei = 59;
static const uint16_t GlTexParameteri = 60;
static const uint16_t GlTexParameterf = 61;
static const uint16_t GlGetTexParameteriv = 62;
static const uint16_t GlGetTexParameterfv = 63;
static const uint16_t GlUniform1i = 64;
static const uint16_t GlUniform2i = 65;
static const uint16_t GlUniform3i = 66;
static const uint16_t GlUniform4i = 67;
static const uint16_t GlUniform1iv = 68;
static const uint16_t GlUniform2iv = 69;
static const uint16_t GlUniform3iv = 70;
static const uint16_t GlUniform4iv = 71;
static const uint16_t GlUniform1f = 72;
static const uint16_t GlUniform2f = 73;
static const uint16_t GlUniform3f = 74;
static const uint16_t GlUniform4f = 75;
static const uint16_t GlUniform1fv = 76;
static const uint16_t GlUniform2fv = 77;
static const uint16_t GlUniform3fv = 78;
static const uint16_t GlUniform4fv = 79;
static const uint16_t GlUniformMatrix2fv = 80;
static const uint16_t GlUniformMatrix3fv = 81;
static const uint16_t GlUniformMatrix4fv = 82;
static const uint16_t GlGetUniformfv = 83;
static const uint16_t GlGetUniformiv = 84;
static const uint16_t GlVertexAttrib1f = 85;
static const uint16_t GlVertexAttrib2f = 86;
static const uint16_t GlVertexAttrib3f = 87;
static const uint16_t GlVertexAttrib4f = 88;
static const uint16_t GlVertexAttrib1fv = 89;
static const uint16_t GlVertexAttrib2fv = 90;
static const uint16_t GlVertexAttrib3fv = 91;
static const uint16_t GlVertexAttrib4fv = 92;
static const uint16_t GlGetShaderPrecisionFormat = 93;
static const uint16_t GlDepthMask = 94;
static const uint16_t GlDepthFunc = 95;
static const uint16_t GlDepthRangef = 96;
static const uint16_t GlColorMask = 97;
static const uint16_t GlStencilMask = 98;
static const uint16_t GlStencilMaskSeparate = 99;
static const uint16_t GlStencilFuncSeparate = 100;
static const uint16_t GlStencilOpSeparate = 101;
static const uint16_t GlFrontFace = 102;
static const uint16_t GlViewport = 103;
static const uint16_t GlScissor = 104;
static const uint16_t GlActiveTexture = 105;
static const uint16_t GlGenTextures = 106;
static const uint16_t GlDeleteTextures = 107;
static const uint16_t GlIsTexture = 108;
static const uint16_t GlBindTexture = 109;
static const uint16_t GlTexImage2D = 110;
static const uint16_t GlTexSubImage2D = 111;
static const uint16_t GlCopyTexImage2D = 112;
static const uint16_t GlCopyTexSubImage2D = 113;
static const uint16_t GlCompressedTexImage2D = 114;
static const uint16_t GlCompressedTexSubImage2D = 115;
static const uint16_t GlGenerateMipmap = 116;
static const uint16_t GlReadPixels = 117;
static const uint16_t GlGenFramebuffers = 118;
static const uint16_t GlBindFramebuffer = 119;
static const uint16_t GlCheckFramebufferStatus = 120;
static const uint16_t GlDeleteFramebuffers = 121;
static const uint16_t GlIsFramebuffer = 122;
static const uint16_t GlGenRenderbuffers = 123;
static const uint16_t GlBindRenderbuffer = 124;
static const uint16_t GlRenderbufferStorage = 125;
static const uint16_t GlDeleteRenderbuffers = 126;
static const uint16_t GlIsRenderbuffer = 127;
static const uint16_t GlGetRenderbufferParameteriv = 128;
static const uint16_t GlGenBuffers = 129;
static const uint16_t GlBindBuffer = 130;
static const uint16_t GlBufferData = 131;
static const uint16_t GlBufferSubData = 132;
static const uint16_t GlDeleteBuffers = 133;
static const uint16_t GlIsBuffer = 134;
static const uint16_t GlGetBufferParameteriv = 135;
static const uint16_t GlCreateShader = 136;
static const uint16_t GlDeleteShader = 137;
static const uint16_t GlShaderSource = 138;
static const uint16_t GlShaderBinary = 139;
static const uint16_t GlGetShaderInfoLog = 140;
static const uint16_t GlGetShaderSource = 141;
static const uint16_t GlReleaseShaderCompiler = 142;
static const uint16_t GlCompileShader = 143;
static const uint16_t GlIsShader = 144;
static const uint16_t GlCreateProgram = 145;
static const uint16_t GlDeleteProgram = 146;
static const uint16_t GlAttachShader = 147;
static const uint16_t GlDetachShader = 148;
static const uint16_t GlGetAttachedShaders = 149;
static const uint16_t GlLinkProgram = 150;
static const uint16_t GlGetProgramInfoLog = 151;
static const uint16_t GlUseProgram = 152;
static const uint16_t GlIsProgram = 153;
static const uint16_t GlValidateProgram = 154;
static const uint16_t GlClearColor = 155;
static const uint16_t GlClearDepthf = 156;
static const uint16_t GlClearStencil = 157;
static const uint16_t GlClear = 158;
static const uint16_t GlCullFace = 159;
static const uint16_t GlPolygonOffset = 160;
static const uint16_t GlLineWidth = 161;
static const uint16_t GlSampleCoverage = 162;
static const uint16_t GlHint = 163;
static const uint16_t GlFramebufferRenderbuffer = 164;
static const uint16_t GlFramebufferTexture2D = 165;
static const uint16_t GlGetFramebufferAttachmentParameteriv = 166;
static const uint16_t GlDrawElements = 167;
static const uint16_t GlDrawArrays = 168;
static const uint16_t GlFlush = 169;
static const uint16_t GlFinish = 170;
static const uint16_t GlGetBooleanv = 171;
static const uint16_t GlGetFloatv = 172;
static const uint16_t GlGetIntegerv = 173;
static const uint16_t GlGetString = 174;
static const uint16_t GlEnable = 175;
static const uint16_t GlDisable = 176;
static const uint16_t GlIsEnabled = 177;
static const uint16_t GlMapBufferRange = 178;
static const uint16_t GlUnmapBuffer = 179;
static const uint16_t GlInvalidateFramebuffer = 180;
static const uint16_t GlRenderbufferStorageMultisample = 181;
static const uint16_t GlBlitFramebuffer = 182;
static const uint16_t GlGenQueries = 183;
static const uint16_t GlBeginQuery = 184;
static const uint16_t GlEndQuery = 185;
static const uint16_t GlDeleteQueries = 186;
static const uint16_t GlIsQuery = 187;
static const uint16_t GlGetQueryiv = 188;
static const uint16_t GlGetQueryObjectuiv = 189;
static const uint16_t GlGetActiveUniformBlockName = 190;
static const uint16_t GlGetActiveUniformBlockiv = 191;
static const uint16_t GlUniformBlockBinding = 192;
static const uint16_t GlGetActiveUniformsiv = 193;
static const uint16_t GlBindBufferBase = 194;
static const uint16_t GlGenVertexArrays = 195;
static const uint16_t GlBindVertexArray = 196;
static const uint16_t GlDeleteVertexArrays = 197;
static const uint16_t GlGenQueriesEXT = 198;
static const uint16_t GlBeginQueryEXT = 199;
static const uint16_t GlEndQueryEXT = 200;
static const uint16_t GlDeleteQueriesEXT = 201;
static const uint16_t GlIsQueryEXT = 202;
static const uint16_t GlQueryCounterEXT = 203;
static const uint16_t GlGetQueryivEXT = 204;
static const uint16_t GlGetQueryObjectivEXT = 205;
static const uint16_t GlGetQueryObjectuivEXT = 206;
static const uint16_t GlGetQueryObjecti64vEXT = 207;
static const uint16_t GlGetQueryObjectui64vEXT = 208;
static const uint16_t Architecture = 209;
static const uint16_t ReplayCreateRenderer = 210;
static const uint16_t ReplayBindRenderer = 211;
static const uint16_t BackbufferInfo = 212;
static const uint16_t StartTimer = 213;
static const uint16_t StopTimer = 214;
static const uint16_t FlushPostBuffer = 215;
}  // namespace FunctionIds

enum class DrawMode : uint32_t {
    GL_LINE_LOOP = 2,
    GL_LINE_STRIP = 3,
    GL_LINES = 1,
    GL_POINTS = 0,
    GL_TRIANGLE_FAN = 6,
    GL_TRIANGLE_STRIP = 5,
    GL_TRIANGLES = 4,
};

enum class IndicesType : uint32_t {
    GL_UNSIGNED_BYTE = 5121,
    GL_UNSIGNED_SHORT = 5123,
    GL_UNSIGNED_INT = 5125,
};

enum class TextureTarget_GLES_1_1 : uint32_t {
    GL_TEXTURE_2D = 3553,
};

enum class TextureTarget_GLES_2_0 : uint32_t {
    GL_TEXTURE_CUBE_MAP = 34067,
};

enum class TextureTarget_OES_EGL_image_external : uint32_t {
    GL_TEXTURE_EXTERNAL_OES = 36197,
};

enum class TextureTarget : uint32_t {};

enum class CubeMapImageTarget : uint32_t {
    GL_TEXTURE_CUBE_MAP_NEGATIVE_X = 34070,
    GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = 34072,
    GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = 34074,
    GL_TEXTURE_CUBE_MAP_POSITIVE_X = 34069,
    GL_TEXTURE_CUBE_MAP_POSITIVE_Y = 34071,
    GL_TEXTURE_CUBE_MAP_POSITIVE_Z = 34073,
};

enum class Texture2DImageTarget : uint32_t {
    GL_TEXTURE_2D = 3553,
};

enum class TextureImageTarget : uint32_t {};

enum class BaseTexelFormat : uint32_t {
    GL_ALPHA = 6406,
    GL_RGB = 6407,
    GL_RGBA = 6408,
};

enum class TexelFormat_GLES_1_1 : uint32_t {
    GL_LUMINANCE = 6409,
    GL_LUMINANCE_ALPHA = 6410,
};

enum class TexelFormat_GLES_3_0 : uint32_t {
    GL_RED = 6403,
    GL_RED_INTEGER = 36244,
    GL_RG = 33319,
    GL_RG_INTEGER = 33320,
    GL_RGB_INTEGER = 36248,
    GL_RGBA_INTEGER = 36249,
    GL_DEPTH_COMPONENT = 6402,
    GL_DEPTH_COMPONENT16 = 33189,
    GL_DEPTH_STENCIL = 34041,
    GL_DEPTH24_STENCIL8 = 35056,
};

enum class TexelFormat : uint32_t {};

enum class RenderbufferFormat : uint32_t {
    GL_RGBA4 = 32854,
    GL_RGB5_A1 = 32855,
    GL_RGB565 = 36194,
    GL_RGBA8 = 32856,
    GL_DEPTH_COMPONENT16 = 33189,
    GL_STENCIL_INDEX8 = 36168,
};

enum class Type_ARB_half_float_vertex : uint32_t {
    GL_ARB_half_float_vertex = 5131,
};

enum class Type_OES_vertex_half_float : uint32_t {
    GL_HALF_FLOAT_OES = 36193,
};

enum class CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture : uint32_t {
    GL_ETC1_RGB8_OES = 36196,
};

enum class CompressedTexelFormat_AMD_compressed_ATC_texture : uint32_t {
    GL_ATC_RGB_AMD = 35986,
    GL_ATC_RGBA_EXPLICIT_ALPHA_AMD = 35987,
    GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = 34798,
};

enum class CompressedTexelFormat : uint32_t {};

enum class ImageTexelFormat : uint32_t {};

enum class TexelType : uint32_t {
    GL_UNSIGNED_BYTE = 5121,
    GL_UNSIGNED_SHORT = 5123,
    GL_UNSIGNED_INT = 5125,
    GL_FLOAT = 5126,
    GL_UNSIGNED_SHORT_4_4_4_4 = 32819,
    GL_UNSIGNED_SHORT_5_5_5_1 = 32820,
    GL_UNSIGNED_SHORT_5_6_5 = 33635,
    GL_UNSIGNED_INT_24_8 = 34042,
};

enum class FramebufferAttachment : uint32_t {
    GL_COLOR_ATTACHMENT0 = 36064,
    GL_DEPTH_ATTACHMENT = 36096,
    GL_STENCIL_ATTACHMENT = 36128,
};

enum class FramebufferAttachmentType : uint32_t {
    GL_NONE = 0,
    GL_RENDERBUFFER = 36161,
    GL_TEXTURE = 5890,
};

enum class FramebufferTarget_GLES_2_0 : uint32_t {
    GL_FRAMEBUFFER = 36160,
};

enum class FramebufferTarget_GLES_3_1 : uint32_t {
    GL_READ_FRAMEBUFFER = 36008,
    GL_DRAW_FRAMEBUFFER = 36009,
};

enum class FramebufferTarget : uint32_t {};

enum class FramebufferAttachmentParameter : uint32_t {
    GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 36048,
    GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 36049,
    GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 36050,
    GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 36051,
};

enum class FramebufferStatus : uint32_t {
    GL_FRAMEBUFFER_COMPLETE = 36053,
    GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 36054,
    GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 36055,
    GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS = 36057,
    GL_FRAMEBUFFER_UNSUPPORTED = 36061,
};

enum class RenderbufferTarget : uint32_t {
    GL_RENDERBUFFER = 36161,
};

enum class RenderbufferParameter : uint32_t {
    GL_RENDERBUFFER_WIDTH = 36162,
    GL_RENDERBUFFER_HEIGHT = 36163,
    GL_RENDERBUFFER_INTERNAL_FORMAT = 36164,
    GL_RENDERBUFFER_RED_SIZE = 36176,
    GL_RENDERBUFFER_GREEN_SIZE = 36177,
    GL_RENDERBUFFER_BLUE_SIZE = 36178,
    GL_RENDERBUFFER_ALPHA_SIZE = 36179,
    GL_RENDERBUFFER_DEPTH_SIZE = 36180,
    GL_RENDERBUFFER_STENCIL_SIZE = 36181,
};

enum class BufferParameter : uint32_t {
    GL_BUFFER_SIZE = 34660,
    GL_BUFFER_USAGE = 34661,
};

enum class TextureUnit : uint32_t {
    GL_TEXTURE0 = 33984,
    GL_TEXTURE1 = 33985,
    GL_TEXTURE2 = 33986,
    GL_TEXTURE3 = 33987,
    GL_TEXTURE4 = 33988,
    GL_TEXTURE5 = 33989,
    GL_TEXTURE6 = 33990,
    GL_TEXTURE7 = 33991,
    GL_TEXTURE8 = 33992,
    GL_TEXTURE9 = 33993,
    GL_TEXTURE10 = 33994,
    GL_TEXTURE11 = 33995,
    GL_TEXTURE12 = 33996,
    GL_TEXTURE13 = 33997,
    GL_TEXTURE14 = 33998,
    GL_TEXTURE15 = 33999,
    GL_TEXTURE16 = 34000,
    GL_TEXTURE17 = 34001,
    GL_TEXTURE18 = 34002,
    GL_TEXTURE19 = 34003,
    GL_TEXTURE20 = 34004,
    GL_TEXTURE21 = 34005,
    GL_TEXTURE22 = 34006,
    GL_TEXTURE23 = 34007,
    GL_TEXTURE24 = 34008,
    GL_TEXTURE25 = 34009,
    GL_TEXTURE26 = 34010,
    GL_TEXTURE27 = 34011,
    GL_TEXTURE28 = 34012,
    GL_TEXTURE29 = 34013,
    GL_TEXTURE30 = 34014,
    GL_TEXTURE31 = 34015,
};

enum class BufferUsage : uint32_t {
    GL_DYNAMIC_DRAW = 35048,
    GL_STATIC_DRAW = 35044,
    GL_STREAM_DRAW = 35040,
};

enum class ShaderType : uint32_t {
    GL_VERTEX_SHADER = 35633,
    GL_FRAGMENT_SHADER = 35632,
};

enum class StateVariable_GLES_2_0 : uint32_t {
    GL_ACTIVE_TEXTURE = 34016,
    GL_ALIASED_LINE_WIDTH_RANGE = 33902,
    GL_ALIASED_POINT_SIZE_RANGE = 33901,
    GL_ALPHA_BITS = 3413,
    GL_ARRAY_BUFFER_BINDING = 34964,
    GL_BLEND = 3042,
    GL_BLEND_COLOR = 32773,
    GL_BLEND_DST_ALPHA = 32970,
    GL_BLEND_DST_RGB = 32968,
    GL_BLEND_EQUATION_ALPHA = 34877,
    GL_BLEND_EQUATION_RGB = 32777,
    GL_BLEND_SRC_ALPHA = 32971,
    GL_BLEND_SRC_RGB = 32969,
    GL_BLUE_BITS = 3412,
    GL_COLOR_CLEAR_VALUE = 3106,
    GL_COLOR_WRITEMASK = 3107,
    GL_COMPRESSED_TEXTURE_FORMATS = 34467,
    GL_CULL_FACE = 2884,
    GL_CULL_FACE_MODE = 2885,
    GL_CURRENT_PROGRAM = 35725,
    GL_DEPTH_BITS = 3414,
    GL_DEPTH_CLEAR_VALUE = 2931,
    GL_DEPTH_FUNC = 2932,
    GL_DEPTH_RANGE = 2928,
    GL_DEPTH_TEST = 2929,
    GL_DEPTH_WRITEMASK = 2930,
    GL_DITHER = 3024,
    GL_ELEMENT_ARRAY_BUFFER_BINDING = 34965,
    GL_FRAMEBUFFER_BINDING = 36006,
    GL_FRONT_FACE = 2886,
    GL_GENERATE_MIPMAP_HINT = 33170,
    GL_GREEN_BITS = 3411,
    GL_IMPLEMENTATION_COLOR_READ_FORMAT = 35739,
    GL_IMPLEMENTATION_COLOR_READ_TYPE = 35738,
    GL_LINE_WIDTH = 2849,
    GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS = 35661,
    GL_MAX_CUBE_MAP_TEXTURE_SIZE = 34076,
    GL_MAX_FRAGMENT_UNIFORM_VECTORS = 36349,
    GL_MAX_RENDERBUFFER_SIZE = 34024,
    GL_MAX_TEXTURE_IMAGE_UNITS = 34930,
    GL_MAX_TEXTURE_SIZE = 3379,
    GL_MAX_VARYING_VECTORS = 36348,
    GL_MAX_VERTEX_ATTRIBS = 34921,
    GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS = 35660,
    GL_MAX_VERTEX_UNIFORM_VECTORS = 36347,
    GL_MAX_VIEWPORT_DIMS = 3386,
    GL_NUM_COMPRESSED_TEXTURE_FORMATS = 34466,
    GL_NUM_SHADER_BINARY_FORMATS = 36345,
    GL_PACK_ALIGNMENT = 3333,
    GL_POLYGON_OFFSET_FACTOR = 32824,
    GL_POLYGON_OFFSET_FILL = 32823,
    GL_POLYGON_OFFSET_UNITS = 10752,
    GL_RED_BITS = 3410,
    GL_RENDERBUFFER_BINDING = 36007,
    GL_SAMPLE_ALPHA_TO_COVERAGE = 32926,
    GL_SAMPLE_BUFFERS = 32936,
    GL_SAMPLE_COVERAGE = 32928,
    GL_SAMPLE_COVERAGE_INVERT = 32939,
    GL_SAMPLE_COVERAGE_VALUE = 32938,
    GL_SAMPLES = 32937,
    GL_SCISSOR_BOX = 3088,
    GL_SCISSOR_TEST = 3089,
    GL_SHADER_BINARY_FORMATS = 36344,
    GL_SHADER_COMPILER = 36346,
    GL_STENCIL_BACK_FAIL = 34817,
    GL_STENCIL_BACK_FUNC = 34816,
    GL_STENCIL_BACK_PASS_DEPTH_FAIL = 34818,
    GL_STENCIL_BACK_PASS_DEPTH_PASS = 34819,
    GL_STENCIL_BACK_REF = 36003,
    GL_STENCIL_BACK_VALUE_MASK = 36004,
    GL_STENCIL_BACK_WRITEMASK = 36005,
    GL_STENCIL_BITS = 3415,
    GL_STENCIL_CLEAR_VALUE = 2961,
    GL_STENCIL_FAIL = 2964,
    GL_STENCIL_FUNC = 2962,
    GL_STENCIL_PASS_DEPTH_FAIL = 2965,
    GL_STENCIL_PASS_DEPTH_PASS = 2966,
    GL_STENCIL_REF = 2967,
    GL_STENCIL_TEST = 2960,
    GL_STENCIL_VALUE_MASK = 2963,
    GL_STENCIL_WRITEMASK = 2968,
    GL_SUBPIXEL_BITS = 3408,
    GL_TEXTURE_BINDING_2D = 32873,
    GL_TEXTURE_BINDING_CUBE_MAP = 34068,
    GL_UNPACK_ALIGNMENT = 3317,
    GL_VIEWPORT = 2978,
};

enum class StateVariable_GLES_3_1 : uint32_t {
    GL_READ_FRAMEBUFFER_BINDING = 36010,
};

enum class StateVariable_EXT_texture_filter_anisotropic : uint32_t {
    GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = 34047,
};

enum class StateVariable_EXT_disjoint_timer_query : uint32_t {
    GL_GPU_DISJOINT_EXT = 36795,
};

enum class StateVariable : uint32_t {};

enum class FaceMode : uint32_t {
    GL_FRONT = 1028,
    GL_BACK = 1029,
    GL_FRONT_AND_BACK = 1032,
};

enum class ArrayType_GLES_1_1 : uint32_t {
    GL_VERTEX_ARRAY = 32884,
    GL_NORMAL_ARRAY = 32885,
    GL_COLOR_ARRAY = 32886,
    GL_TEXTURE_COORD_ARRAY = 32888,
};

enum class ArrayType_OES_point_size_array : uint32_t {
    GL_POINT_SIZE_ARRAY_OES = 35740,
};

enum class ArrayType : uint32_t {};

enum class Capability : uint32_t {
    GL_BLEND = 3042,
    GL_CULL_FACE = 2884,
    GL_DEPTH_TEST = 2929,
    GL_DITHER = 3024,
    GL_POLYGON_OFFSET_FILL = 32823,
    GL_SAMPLE_ALPHA_TO_COVERAGE = 32926,
    GL_SAMPLE_COVERAGE = 32928,
    GL_SCISSOR_TEST = 3089,
    GL_STENCIL_TEST = 2960,
};

enum class StringConstant : uint32_t {
    GL_EXTENSIONS = 7939,
    GL_RENDERER = 7937,
    GL_VENDOR = 7936,
    GL_VERSION = 7938,
};

enum class VertexAttribType : uint32_t {
    GL_BYTE = 5120,
    GL_FIXED = 5132,
    GL_FLOAT = 5126,
    GL_SHORT = 5122,
    GL_UNSIGNED_BYTE = 5121,
    GL_UNSIGNED_SHORT = 5123,
};

enum class ShaderAttribType : uint32_t {
    GL_FLOAT = 5126,
    GL_FLOAT_VEC2 = 35664,
    GL_FLOAT_VEC3 = 35665,
    GL_FLOAT_VEC4 = 35666,
    GL_FLOAT_MAT2 = 35674,
    GL_FLOAT_MAT3 = 35675,
    GL_FLOAT_MAT4 = 35676,
};

enum class ShaderUniformType : uint32_t {
    GL_FLOAT = 5126,
    GL_FLOAT_VEC2 = 35664,
    GL_FLOAT_VEC3 = 35665,
    GL_FLOAT_VEC4 = 35666,
    GL_INT = 5124,
    GL_INT_VEC2 = 35667,
    GL_INT_VEC3 = 35668,
    GL_INT_VEC4 = 35669,
    GL_BOOL = 35670,
    GL_BOOL_VEC2 = 35671,
    GL_BOOL_VEC3 = 35672,
    GL_BOOL_VEC4 = 35673,
    GL_FLOAT_MAT2 = 35674,
    GL_FLOAT_MAT3 = 35675,
    GL_FLOAT_MAT4 = 35676,
    GL_SAMPLER_2D = 35678,
    GL_SAMPLER_CUBE = 35680,
};

enum class Error : uint32_t {
    GL_NO_ERROR = 0,
    GL_INVALID_ENUM = 1280,
    GL_INVALID_VALUE = 1281,
    GL_INVALID_OPERATION = 1282,
    GL_INVALID_FRAMEBUFFER_OPERATION = 1286,
    GL_OUT_OF_MEMORY = 1285,
};

enum class HintTarget : uint32_t {
    GL_GENERATE_MIPMAP_HINT = 33170,
};

enum class HintMode : uint32_t {
    GL_DONT_CARE = 4352,
    GL_FASTEST = 4353,
    GL_NICEST = 4354,
};

enum class DiscardFramebufferAttachment : uint32_t {
    GL_COLOR_EXT = 6144,
    GL_DEPTH_EXT = 6145,
    GL_STENCIL_EXT = 6146,
};

enum class ProgramParameter : uint32_t {
    GL_DELETE_STATUS = 35712,
    GL_LINK_STATUS = 35714,
    GL_VALIDATE_STATUS = 35715,
    GL_INFO_LOG_LENGTH = 35716,
    GL_ATTACHED_SHADERS = 35717,
    GL_ACTIVE_ATTRIBUTES = 35721,
    GL_ACTIVE_ATTRIBUTE_MAX_LENGTH = 35722,
    GL_ACTIVE_UNIFORMS = 35718,
    GL_ACTIVE_UNIFORM_MAX_LENGTH = 35719,
};

enum class ShaderParameter : uint32_t {
    GL_SHADER_TYPE = 35663,
    GL_DELETE_STATUS = 35712,
    GL_COMPILE_STATUS = 35713,
    GL_INFO_LOG_LENGTH = 35716,
    GL_SHADER_SOURCE_LENGTH = 35720,
};

enum class PixelStoreParameter : uint32_t {
    GL_PACK_ALIGNMENT = 3333,
    GL_UNPACK_ALIGNMENT = 3317,
};

enum class TextureParameter_FilterMode : uint32_t {
    GL_TEXTURE_MIN_FILTER = 10241,
    GL_TEXTURE_MAG_FILTER = 10240,
};

enum class TextureParameter_WrapMode : uint32_t {
    GL_TEXTURE_WRAP_S = 10242,
    GL_TEXTURE_WRAP_T = 10243,
};

enum class TextureParameter_EXT_texture_filter_anisotropic : uint32_t {
    GL_TEXTURE_MAX_ANISOTROPY_EXT = 34046,
};

enum class TextureParameter_SwizzleMode : uint32_t {
    GL_TEXTURE_SWIZZLE_R = 36418,
    GL_TEXTURE_SWIZZLE_G = 36419,
    GL_TEXTURE_SWIZZLE_B = 36420,
    GL_TEXTURE_SWIZZLE_A = 36421,
};

enum class TextureParameter : uint32_t {};

enum class TextureFilterMode : uint32_t {
    GL_NEAREST = 9728,
    GL_LINEAR = 9729,
    GL_NEAREST_MIPMAP_NEAREST = 9984,
    GL_LINEAR_MIPMAP_NEAREST = 9985,
    GL_NEAREST_MIPMAP_LINEAR = 9986,
    GL_LINEAR_MIPMAP_LINEAR = 9987,
};

enum class TextureWrapMode : uint32_t {
    GL_CLAMP_TO_EDGE = 33071,
    GL_MIRRORED_REPEAT = 33648,
    GL_REPEAT = 10497,
};

enum class TexelComponent : uint32_t {
    GL_RED = 6403,
    GL_GREEN = 6404,
    GL_BLUE = 6405,
    GL_ALPHA = 6406,
};

enum class BlendFactor : uint32_t {
    GL_ZERO = 0,
    GL_ONE = 1,
    GL_SRC_COLOR = 768,
    GL_ONE_MINUS_SRC_COLOR = 769,
    GL_DST_COLOR = 774,
    GL_ONE_MINUS_DST_COLOR = 775,
    GL_SRC_ALPHA = 770,
    GL_ONE_MINUS_SRC_ALPHA = 771,
    GL_DST_ALPHA = 772,
    GL_ONE_MINUS_DST_ALPHA = 773,
    GL_CONSTANT_COLOR = 32769,
    GL_ONE_MINUS_CONSTANT_COLOR = 32770,
    GL_CONSTANT_ALPHA = 32771,
    GL_ONE_MINUS_CONSTANT_ALPHA = 32772,
    GL_SRC_ALPHA_SATURATE = 776,
};

enum class PrecisionType : uint32_t {
    GL_LOW_FLOAT = 36336,
    GL_MEDIUM_FLOAT = 36337,
    GL_HIGH_FLOAT = 36338,
    GL_LOW_INT = 36339,
    GL_MEDIUM_INT = 36340,
    GL_HIGH_INT = 36341,
};

enum class TestFunction : uint32_t {
    GL_NEVER = 512,
    GL_LESS = 513,
    GL_EQUAL = 514,
    GL_LEQUAL = 515,
    GL_GREATER = 516,
    GL_NOTEQUAL = 517,
    GL_GEQUAL = 518,
    GL_ALWAYS = 519,
};

enum class StencilAction : uint32_t {
    GL_KEEP = 7680,
    GL_ZERO = 0,
    GL_REPLACE = 7681,
    GL_INCR = 7682,
    GL_INCR_WRAP = 34055,
    GL_DECR = 7683,
    GL_DECR_WRAP = 34056,
    GL_INVERT = 5386,
};

enum class FaceOrientation : uint32_t {
    GL_CW = 2304,
    GL_CCW = 2305,
};

enum class BlendEquation : uint32_t {
    GL_FUNC_ADD = 32774,
    GL_FUNC_SUBTRACT = 32778,
    GL_FUNC_REVERSE_SUBTRACT = 32779,
};

enum class BufferTarget : uint32_t {
    GL_ARRAY_BUFFER = 34962,
    GL_COPY_READ_BUFFER = 36662,
    GL_COPY_WRITE_BUFFER = 36663,
    GL_ELEMENT_ARRAY_BUFFER = 34963,
    GL_PIXEL_PACK_BUFFER = 35051,
    GL_PIXEL_UNPACK_BUFFER = 35052,
    GL_TRANSFORM_FEEDBACK_BUFFER = 35982,
    GL_UNIFORM_BUFFER = 35345,
};

enum class ImageTargetTexture_OES_EGL_image : uint32_t {
    GL_TEXTURE_2D = 3553,
};

enum class ImageTargetTexture_OES_EGL_image_external : uint32_t {
    GL_TEXTURE_EXTERNAL_OES = 36197,
};

enum class ImageTargetTexture : uint32_t {};

enum class ImageTargetRenderbufferStorage : uint32_t {
    GL_RENDERBUFFER_OES = 36161,
};

enum class ResetStatus : uint32_t {
    GL_NO_ERROR = 0,
    GL_GUILTY_CONTEXT_RESET_EXT = 33363,
    GL_INNOCENT_CONTEXT_RESET_EXT = 33364,
    GL_UNKNOWN_CONTEXT_RESET_EXT = 33365,
};

enum class TextureKind : uint32_t {
    UNDEFINED = 0,
    TEXTURE2D = 1,
    CUBEMAP = 2,
};

enum class QueryParameter_GLES_3 : uint32_t {
    GL_CURRENT_QUERY = 34917,
};

enum class QueryParameter_EXT_disjoint_timer_query : uint32_t {
    GL_QUERY_COUNTER_BITS_EXT = 34916,
};

enum class QueryParameter : uint32_t {};

enum class QueryObjectParameter_GLES_3 : uint32_t {
    GL_QUERY_RESULT = 34918,
    GL_QUERY_RESULT_AVAILABLE = 34919,
};

enum class QueryObjectParameter_EXT_disjoint_timer_query : uint32_t {};

enum class QueryObjectParameter : uint32_t {};

enum class QueryTarget_GLES_3 : uint32_t {
    GL_ANY_SAMPLES_PASSED = 35887,
    GL_ANY_SAMPLES_PASSED_CONSERVATIVE = 36202,
    GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = 35976,
};

enum class QueryTarget_EXT_disjoint_timer_query : uint32_t {
    GL_TIME_ELAPSED_EXT = 35007,
    GL_TIMESTAMP_EXT = 36392,
};

enum class QueryTarget : uint32_t {};

enum class UniformBlockParameter : uint32_t {
    GL_UNIFORM_BLOCK_BINDING = 35391,
    GL_UNIFORM_BLOCK_DATA_SIZE = 35392,
    GL_UNIFORM_BLOCK_NAME_LENGTH = 35393,
    GL_UNIFORM_BLOCK_ACTIVE_UNIFORMS = 35394,
    GL_UNIFORM_BLOCK_ACTIVE_UNIFORM_INDICES = 35395,
    GL_UNIFORM_BLOCK_REFERENCED_BY_VERTEX_SHADER = 35396,
    GL_UNIFORM_BLOCK_REFERENCED_BY_GEOMETRY_SHADER = 35397,
    GL_UNIFORM_BLOCK_REFERENCED_BY_FRAGMENT_SHADER = 35398,
};

enum class IndexedBufferTarget : uint32_t {
    GL_TRANSFORM_FEEDBACK_BUFFER = 35982,
    GL_UNIFORM_BUFFER = 35345,
};

enum class TilePreserveMaskQCOM : uint32_t {
    GL_COLOR_BUFFER_BIT0_QCOM = 1,
    GL_COLOR_BUFFER_BIT1_QCOM = 2,
    GL_COLOR_BUFFER_BIT2_QCOM = 4,
    GL_COLOR_BUFFER_BIT3_QCOM = 8,
    GL_COLOR_BUFFER_BIT4_QCOM = 16,
    GL_COLOR_BUFFER_BIT5_QCOM = 32,
    GL_COLOR_BUFFER_BIT6_QCOM = 64,
    GL_COLOR_BUFFER_BIT7_QCOM = 128,
    GL_DEPTH_BUFFER_BIT0_QCOM = 256,
    GL_DEPTH_BUFFER_BIT1_QCOM = 512,
    GL_DEPTH_BUFFER_BIT2_QCOM = 1024,
    GL_DEPTH_BUFFER_BIT3_QCOM = 2048,
    GL_DEPTH_BUFFER_BIT4_QCOM = 4096,
    GL_DEPTH_BUFFER_BIT5_QCOM = 8192,
    GL_DEPTH_BUFFER_BIT6_QCOM = 16384,
    GL_DEPTH_BUFFER_BIT7_QCOM = 32768,
    GL_STENCIL_BUFFER_BIT0_QCOM = 65536,
    GL_STENCIL_BUFFER_BIT1_QCOM = 131072,
    GL_STENCIL_BUFFER_BIT2_QCOM = 262144,
    GL_STENCIL_BUFFER_BIT3_QCOM = 524288,
    GL_STENCIL_BUFFER_BIT4_QCOM = 1048576,
    GL_STENCIL_BUFFER_BIT5_QCOM = 2097152,
    GL_STENCIL_BUFFER_BIT6_QCOM = 4194304,
    GL_STENCIL_BUFFER_BIT7_QCOM = 8388608,
    GL_MULTISAMPLE_BUFFER_BIT0_QCOM = 16777216,
    GL_MULTISAMPLE_BUFFER_BIT1_QCOM = 33554432,
    GL_MULTISAMPLE_BUFFER_BIT2_QCOM = 67108864,
    GL_MULTISAMPLE_BUFFER_BIT3_QCOM = 134217728,
    GL_MULTISAMPLE_BUFFER_BIT4_QCOM = 268435456,
    GL_MULTISAMPLE_BUFFER_BIT5_QCOM = 536870912,
    GL_MULTISAMPLE_BUFFER_BIT6_QCOM = 1073741824,
    GL_MULTISAMPLE_BUFFER_BIT7_QCOM = 2147483648,
};

enum class ClearMask : uint32_t {
    GL_COLOR_BUFFER_BIT = 16384,
    GL_DEPTH_BUFFER_BIT = 256,
    GL_STENCIL_BUFFER_BIT = 1024,
};

enum class MapBufferRangeAccess : uint32_t {
    GL_MAP_READ_BIT = 1,
    GL_MAP_WRITE_BIT = 2,
    GL_MAP_INVALIDATE_RANGE_BIT = 4,
    GL_MAP_INVALIDATE_BUFFER_BIT = 8,
    GL_MAP_FLUSH_EXPLICIT_BIT = 16,
    GL_MAP_UNSYNCHRONIZED_BIT = 32,
};

typedef int32_t Vec2i;
typedef int32_t Vec3i;
typedef int32_t Vec4i;
typedef float Vec2f;
typedef float Vec3f;
typedef float Vec4f;
typedef Vec2f Mat2f;
typedef Vec3f Mat3f;
typedef Vec4f Mat4f;
typedef uint32_t RenderbufferId;
typedef uint32_t TextureId;
typedef uint32_t FramebufferId;
typedef uint32_t BufferId;
typedef uint32_t ShaderId;
typedef uint32_t ProgramId;
typedef uint32_t VertexArrayId;
typedef uint32_t QueryId;
typedef int32_t UniformLocation;
typedef int32_t AttributeLocation;
typedef void *IndicesPointer;
typedef void *VertexPointer;
typedef void *TexturePointer;
typedef void *BufferDataPointer;
typedef uint32_t ContextID;
typedef uint32_t ThreadID;
typedef int EGLBoolean;
typedef int EGLint;
typedef void *EGLConfig;
typedef void *EGLContext;
typedef void *EGLDisplay;
typedef void *EGLSurface;
typedef void *GLXContext;
typedef void *GLXDrawable;
typedef int Bool;
typedef void *HGLRC;
typedef void *HDC;
typedef int BOOL;
typedef int CGLError;
typedef void *CGLPixelFormatObj;
typedef void *CGLContextObj;
typedef void *CGSConnectionID;
typedef int32_t CGSWindowID;
typedef int32_t CGSSurfaceID;
typedef void *ImageOES;

typedef int(STDCALL *PFNEGLINITIALIZE)(void *dpy, int *major, int *minor);
typedef void *(STDCALL *PFNEGLCREATECONTEXT)(void *display, void *config, void *share_context,
                                             int *attrib_list);
typedef int(STDCALL *PFNEGLMAKECURRENT)(void *display, void *draw, void *read, void *context);
typedef int(STDCALL *PFNEGLSWAPBUFFERS)(void *display, void *surface);
typedef int(STDCALL *PFNEGLQUERYSURFACE)(void *display, void *surface, int attribute, int *value);
typedef void *(STDCALL *PFNGLXCREATECONTEXT)(void *dpy, void *vis, void *shareList, bool direct);
typedef void *(STDCALL *PFNGLXCREATENEWCONTEXT)(void *display, void *fbconfig, uint32_t type,
                                                void *shared, bool direct);
typedef int(STDCALL *PFNGLXMAKECONTEXTCURRENT)(void *display, void *draw, void *read, void *ctx);
typedef int(STDCALL *PFNGLXMAKECURRENT)(void *display, void *drawable, void *ctx);
typedef void(STDCALL *PFNGLXSWAPBUFFERS)(void *display, void *drawable);
typedef int(STDCALL *PFNGLXQUERYDRAWABLE)(void *display, void *draw, int attribute, int *value);
typedef void *(STDCALL *PFNWGLCREATECONTEXT)(void *hdc);
typedef void *(STDCALL *PFNWGLCREATECONTEXTATTRIBSARB)(void *hdc, void *hShareContext,
                                                       int *attribList);
typedef int(STDCALL *PFNWGLMAKECURRENT)(void *hdc, void *hglrc);
typedef void(STDCALL *PFNWGLSWAPBUFFERS)(void *hdc);
typedef int(STDCALL *PFNCGLCREATECONTEXT)(void *pix, void *share, void **ctx);
typedef int(STDCALL *PFNCGLSETCURRENTCONTEXT)(void *ctx);
typedef int(STDCALL *PFNCGLGETSURFACE)(void *ctx, void **cid, int32_t *wid, int32_t *sid);
typedef int(STDCALL *PFNCGSGETSURFACEBOUNDS)(void *cid, int32_t wid, int32_t sid, double *bounds);
typedef int(STDCALL *PFNCGLFLUSHDRAWABLE)(void *ctx);
typedef void(STDCALL *PFNGLENABLECLIENTSTATE)(ArrayType type);
typedef void(STDCALL *PFNGLDISABLECLIENTSTATE)(ArrayType type);
typedef void(STDCALL *PFNGLGETPROGRAMBINARYOES)(uint32_t program, int32_t buffer_size,
                                                int32_t *bytes_written, uint32_t *binary_format,
                                                void *binary);
typedef void(STDCALL *PFNGLPROGRAMBINARYOES)(uint32_t program, uint32_t binary_format, void *binary,
                                             int32_t binary_size);
typedef void(STDCALL *PFNGLSTARTTILINGQCOM)(int32_t x, int32_t y, int32_t width, int32_t height,
                                            TilePreserveMaskQCOM preserveMask);
typedef void(STDCALL *PFNGLENDTILINGQCOM)(TilePreserveMaskQCOM preserve_mask);
typedef void(STDCALL *PFNGLDISCARDFRAMEBUFFEREXT)(FramebufferTarget target, int32_t numAttachments,
                                                  DiscardFramebufferAttachment *attachments);
typedef void(STDCALL *PFNGLINSERTEVENTMARKEREXT)(int32_t length, char *marker);
typedef void(STDCALL *PFNGLPUSHGROUPMARKEREXT)(int32_t length, char *marker);
typedef void(STDCALL *PFNGLPOPGROUPMARKEREXT)();
typedef void(STDCALL *PFNGLTEXSTORAGE1DEXT)(TextureTarget target, int32_t levels,
                                            TexelFormat format, int32_t width);
typedef void(STDCALL *PFNGLTEXSTORAGE2DEXT)(TextureTarget target, int32_t levels,
                                            TexelFormat format, int32_t width, int32_t height);
typedef void(STDCALL *PFNGLTEXSTORAGE3DEXT)(TextureTarget target, int32_t levels,
                                            TexelFormat format, int32_t width, int32_t height,
                                            int32_t depth);
typedef void(STDCALL *PFNGLTEXTURESTORAGE1DEXT)(uint32_t texture, TextureTarget target,
                                                int32_t levels, TexelFormat format, int32_t width);
typedef void(STDCALL *PFNGLTEXTURESTORAGE2DEXT)(uint32_t texture, TextureTarget target,
                                                int32_t levels, TexelFormat format, int32_t width,
                                                int32_t height);
typedef void(STDCALL *PFNGLTEXTURESTORAGE3DEXT)(uint32_t texture, TextureTarget target,
                                                int32_t levels, TexelFormat format, int32_t width,
                                                int32_t height, int32_t depth);
typedef void(STDCALL *PFNGLGENVERTEXARRAYSOES)(int32_t count, uint32_t *arrays);
typedef void(STDCALL *PFNGLBINDVERTEXARRAYOES)(uint32_t array);
typedef void(STDCALL *PFNGLDELETEVERTEXARRAYSOES)(int32_t count, uint32_t *arrays);
typedef bool(STDCALL *PFNGLISVERTEXARRAYOES)(uint32_t array);
typedef void(STDCALL *PFNGLEGLIMAGETARGETTEXTURE2DOES)(ImageTargetTexture target, void *image);
typedef void(STDCALL *PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOES)(
        ImageTargetRenderbufferStorage target, void *image);
typedef ResetStatus(STDCALL *PFNGLGETGRAPHICSRESETSTATUSEXT)();
typedef void(STDCALL *PFNGLBINDATTRIBLOCATION)(uint32_t program, int32_t location, char *name);
typedef void(STDCALL *PFNGLBLENDFUNC)(BlendFactor src_factor, BlendFactor dst_factor);
typedef void(STDCALL *PFNGLBLENDFUNCSEPARATE)(BlendFactor src_factor_rgb,
                                              BlendFactor dst_factor_rgb,
                                              BlendFactor src_factor_alpha,
                                              BlendFactor dst_factor_alpha);
typedef void(STDCALL *PFNGLBLENDEQUATION)(BlendEquation equation);
typedef void(STDCALL *PFNGLBLENDEQUATIONSEPARATE)(BlendEquation rgb, BlendEquation alpha);
typedef void(STDCALL *PFNGLBLENDCOLOR)(float red, float green, float blue, float alpha);
typedef void(STDCALL *PFNGLENABLEVERTEXATTRIBARRAY)(int32_t location);
typedef void(STDCALL *PFNGLDISABLEVERTEXATTRIBARRAY)(int32_t location);
typedef void(STDCALL *PFNGLVERTEXATTRIBPOINTER)(int32_t location, int32_t size,
                                                VertexAttribType type, bool normalized,
                                                int32_t stride, void *data);
typedef void(STDCALL *PFNGLGETACTIVEATTRIB)(uint32_t program, int32_t location, int32_t buffer_size,
                                            int32_t *buffer_bytes_written, int32_t *vector_count,
                                            ShaderAttribType *type, char *name);
typedef void(STDCALL *PFNGLGETACTIVEUNIFORM)(uint32_t program, int32_t location,
                                             int32_t buffer_size, int32_t *buffer_bytes_written,
                                             int32_t *vector_count, ShaderUniformType *type,
                                             char *name);
typedef Error(STDCALL *PFNGLGETERROR)();
typedef void(STDCALL *PFNGLGETPROGRAMIV)(uint32_t program, ProgramParameter parameter,
                                         int32_t *value);
typedef void(STDCALL *PFNGLGETSHADERIV)(uint32_t shader, ShaderParameter parameter, int32_t *value);
typedef int32_t(STDCALL *PFNGLGETUNIFORMLOCATION)(uint32_t program, char *name);
typedef int32_t(STDCALL *PFNGLGETATTRIBLOCATION)(uint32_t program, char *name);
typedef void(STDCALL *PFNGLPIXELSTOREI)(PixelStoreParameter parameter, int32_t value);
typedef void(STDCALL *PFNGLTEXPARAMETERI)(TextureTarget target, TextureParameter parameter,
                                          int32_t value);
typedef void(STDCALL *PFNGLTEXPARAMETERF)(TextureTarget target, TextureParameter parameter,
                                          float value);
typedef void(STDCALL *PFNGLGETTEXPARAMETERIV)(TextureTarget target, TextureParameter parameter,
                                              int32_t *values);
typedef void(STDCALL *PFNGLGETTEXPARAMETERFV)(TextureTarget target, TextureParameter parameter,
                                              float *values);
typedef void(STDCALL *PFNGLUNIFORM1I)(int32_t location, int32_t value);
typedef void(STDCALL *PFNGLUNIFORM2I)(int32_t location, int32_t value0, int32_t value1);
typedef void(STDCALL *PFNGLUNIFORM3I)(int32_t location, int32_t value0, int32_t value1,
                                      int32_t value2);
typedef void(STDCALL *PFNGLUNIFORM4I)(int32_t location, int32_t value0, int32_t value1,
                                      int32_t value2, int32_t value3);
typedef void(STDCALL *PFNGLUNIFORM1IV)(int32_t location, int32_t count, int32_t *values);
typedef void(STDCALL *PFNGLUNIFORM2IV)(int32_t location, int32_t count, int32_t *values);
typedef void(STDCALL *PFNGLUNIFORM3IV)(int32_t location, int32_t count, int32_t *values);
typedef void(STDCALL *PFNGLUNIFORM4IV)(int32_t location, int32_t count, int32_t *values);
typedef void(STDCALL *PFNGLUNIFORM1F)(int32_t location, float value);
typedef void(STDCALL *PFNGLUNIFORM2F)(int32_t location, float value0, float value1);
typedef void(STDCALL *PFNGLUNIFORM3F)(int32_t location, float value0, float value1, float value2);
typedef void(STDCALL *PFNGLUNIFORM4F)(int32_t location, float value0, float value1, float value2,
                                      float value3);
typedef void(STDCALL *PFNGLUNIFORM1FV)(int32_t location, int32_t count, float *values);
typedef void(STDCALL *PFNGLUNIFORM2FV)(int32_t location, int32_t count, float *values);
typedef void(STDCALL *PFNGLUNIFORM3FV)(int32_t location, int32_t count, float *values);
typedef void(STDCALL *PFNGLUNIFORM4FV)(int32_t location, int32_t count, float *values);
typedef void(STDCALL *PFNGLUNIFORMMATRIX2FV)(int32_t location, int32_t count, bool transpose,
                                             float *values);
typedef void(STDCALL *PFNGLUNIFORMMATRIX3FV)(int32_t location, int32_t count, bool transpose,
                                             float *values);
typedef void(STDCALL *PFNGLUNIFORMMATRIX4FV)(int32_t location, int32_t count, bool transpose,
                                             float *values);
typedef void(STDCALL *PFNGLGETUNIFORMFV)(uint32_t program, int32_t location, float *values);
typedef void(STDCALL *PFNGLGETUNIFORMIV)(uint32_t program, int32_t location, int32_t *values);
typedef void(STDCALL *PFNGLVERTEXATTRIB1F)(int32_t location, float value0);
typedef void(STDCALL *PFNGLVERTEXATTRIB2F)(int32_t location, float value0, float value1);
typedef void(STDCALL *PFNGLVERTEXATTRIB3F)(int32_t location, float value0, float value1,
                                           float value2);
typedef void(STDCALL *PFNGLVERTEXATTRIB4F)(int32_t location, float value0, float value1,
                                           float value2, float value3);
typedef void(STDCALL *PFNGLVERTEXATTRIB1FV)(int32_t location, float *value);
typedef void(STDCALL *PFNGLVERTEXATTRIB2FV)(int32_t location, float *value);
typedef void(STDCALL *PFNGLVERTEXATTRIB3FV)(int32_t location, float *value);
typedef void(STDCALL *PFNGLVERTEXATTRIB4FV)(int32_t location, float *value);
typedef void(STDCALL *PFNGLGETSHADERPRECISIONFORMAT)(ShaderType shader_type,
                                                     PrecisionType precision_type, int32_t *range,
                                                     int32_t *precision);
typedef void(STDCALL *PFNGLDEPTHMASK)(bool enabled);
typedef void(STDCALL *PFNGLDEPTHFUNC)(TestFunction function);
typedef void(STDCALL *PFNGLDEPTHRANGEF)(float near, float far);
typedef void(STDCALL *PFNGLCOLORMASK)(bool red, bool green, bool blue, bool alpha);
typedef void(STDCALL *PFNGLSTENCILMASK)(uint32_t mask);
typedef void(STDCALL *PFNGLSTENCILMASKSEPARATE)(FaceMode face, uint32_t mask);
typedef void(STDCALL *PFNGLSTENCILFUNCSEPARATE)(FaceMode face, TestFunction function,
                                                int32_t reference_value, int32_t mask);
typedef void(STDCALL *PFNGLSTENCILOPSEPARATE)(FaceMode face, StencilAction stencil_fail,
                                              StencilAction stencil_pass_depth_fail,
                                              StencilAction stencil_pass_depth_pass);
typedef void(STDCALL *PFNGLFRONTFACE)(FaceOrientation orientation);
typedef void(STDCALL *PFNGLVIEWPORT)(int32_t x, int32_t y, int32_t width, int32_t height);
typedef void(STDCALL *PFNGLSCISSOR)(int32_t x, int32_t y, int32_t width, int32_t height);
typedef void(STDCALL *PFNGLACTIVETEXTURE)(TextureUnit unit);
typedef void(STDCALL *PFNGLGENTEXTURES)(int32_t count, uint32_t *textures);
typedef void(STDCALL *PFNGLDELETETEXTURES)(int32_t count, uint32_t *textures);
typedef bool(STDCALL *PFNGLISTEXTURE)(uint32_t texture);
typedef void(STDCALL *PFNGLBINDTEXTURE)(TextureTarget target, uint32_t texture);
typedef void(STDCALL *PFNGLTEXIMAGE2D)(TextureImageTarget target, int32_t level,
                                       TexelFormat internal_format, int32_t width, int32_t height,
                                       int32_t border, TexelFormat format, TexelType type,
                                       void *data);
typedef void(STDCALL *PFNGLTEXSUBIMAGE2D)(TextureImageTarget target, int32_t level, int32_t xoffset,
                                          int32_t yoffset, int32_t width, int32_t height,
                                          TexelFormat format, TexelType type, void *data);
typedef void(STDCALL *PFNGLCOPYTEXIMAGE2D)(TextureImageTarget target, int32_t level,
                                           TexelFormat format, int32_t x, int32_t y, int32_t width,
                                           int32_t height, int32_t border);
typedef void(STDCALL *PFNGLCOPYTEXSUBIMAGE2D)(TextureImageTarget target, int32_t level,
                                              int32_t xoffset, int32_t yoffset, int32_t x,
                                              int32_t y, int32_t width, int32_t height);
typedef void(STDCALL *PFNGLCOMPRESSEDTEXIMAGE2D)(TextureImageTarget target, int32_t level,
                                                 CompressedTexelFormat format, int32_t width,
                                                 int32_t height, int32_t border, int32_t image_size,
                                                 void *data);
typedef void(STDCALL *PFNGLCOMPRESSEDTEXSUBIMAGE2D)(TextureImageTarget target, int32_t level,
                                                    int32_t xoffset, int32_t yoffset, int32_t width,
                                                    int32_t height, CompressedTexelFormat format,
                                                    int32_t image_size, void *data);
typedef void(STDCALL *PFNGLGENERATEMIPMAP)(TextureImageTarget target);
typedef void(STDCALL *PFNGLREADPIXELS)(int32_t x, int32_t y, int32_t width, int32_t height,
                                       BaseTexelFormat format, TexelType type, void *data);
typedef void(STDCALL *PFNGLGENFRAMEBUFFERS)(int32_t count, uint32_t *framebuffers);
typedef void(STDCALL *PFNGLBINDFRAMEBUFFER)(FramebufferTarget target, uint32_t framebuffer);
typedef FramebufferStatus(STDCALL *PFNGLCHECKFRAMEBUFFERSTATUS)(FramebufferTarget target);
typedef void(STDCALL *PFNGLDELETEFRAMEBUFFERS)(int32_t count, uint32_t *framebuffers);
typedef bool(STDCALL *PFNGLISFRAMEBUFFER)(uint32_t framebuffer);
typedef void(STDCALL *PFNGLGENRENDERBUFFERS)(int32_t count, uint32_t *renderbuffers);
typedef void(STDCALL *PFNGLBINDRENDERBUFFER)(RenderbufferTarget target, uint32_t renderbuffer);
typedef void(STDCALL *PFNGLRENDERBUFFERSTORAGE)(RenderbufferTarget target,
                                                RenderbufferFormat format, int32_t width,
                                                int32_t height);
typedef void(STDCALL *PFNGLDELETERENDERBUFFERS)(int32_t count, uint32_t *renderbuffers);
typedef bool(STDCALL *PFNGLISRENDERBUFFER)(uint32_t renderbuffer);
typedef void(STDCALL *PFNGLGETRENDERBUFFERPARAMETERIV)(RenderbufferTarget target,
                                                       RenderbufferParameter parameter,
                                                       int32_t *values);
typedef void(STDCALL *PFNGLGENBUFFERS)(int32_t count, uint32_t *buffers);
typedef void(STDCALL *PFNGLBINDBUFFER)(BufferTarget target, uint32_t buffer);
typedef void(STDCALL *PFNGLBUFFERDATA)(BufferTarget target, int32_t size, void *data,
                                       BufferUsage usage);
typedef void(STDCALL *PFNGLBUFFERSUBDATA)(BufferTarget target, int32_t offset, int32_t size,
                                          void *data);
typedef void(STDCALL *PFNGLDELETEBUFFERS)(int32_t count, uint32_t *buffers);
typedef bool(STDCALL *PFNGLISBUFFER)(uint32_t buffer);
typedef void(STDCALL *PFNGLGETBUFFERPARAMETERIV)(BufferTarget target, BufferParameter parameter,
                                                 int32_t *value);
typedef uint32_t(STDCALL *PFNGLCREATESHADER)(ShaderType type);
typedef void(STDCALL *PFNGLDELETESHADER)(uint32_t shader);
typedef void(STDCALL *PFNGLSHADERSOURCE)(uint32_t shader, int32_t count, char **source,
                                         int32_t *length);
typedef void(STDCALL *PFNGLSHADERBINARY)(int32_t count, uint32_t *shaders, uint32_t binary_format,
                                         void *binary, int32_t binary_size);
typedef void(STDCALL *PFNGLGETSHADERINFOLOG)(uint32_t shader, int32_t buffer_length,
                                             int32_t *string_length_written, char *info);
typedef void(STDCALL *PFNGLGETSHADERSOURCE)(uint32_t shader, int32_t buffer_length,
                                            int32_t *string_length_written, char *source);
typedef void(STDCALL *PFNGLRELEASESHADERCOMPILER)();
typedef void(STDCALL *PFNGLCOMPILESHADER)(uint32_t shader);
typedef bool(STDCALL *PFNGLISSHADER)(uint32_t shader);
typedef uint32_t(STDCALL *PFNGLCREATEPROGRAM)();
typedef void(STDCALL *PFNGLDELETEPROGRAM)(uint32_t program);
typedef void(STDCALL *PFNGLATTACHSHADER)(uint32_t program, uint32_t shader);
typedef void(STDCALL *PFNGLDETACHSHADER)(uint32_t program, uint32_t shader);
typedef void(STDCALL *PFNGLGETATTACHEDSHADERS)(uint32_t program, int32_t buffer_length,
                                               int32_t *shaders_length_written, uint32_t *shaders);
typedef void(STDCALL *PFNGLLINKPROGRAM)(uint32_t program);
typedef void(STDCALL *PFNGLGETPROGRAMINFOLOG)(uint32_t program, int32_t buffer_length,
                                              int32_t *string_length_written, char *info);
typedef void(STDCALL *PFNGLUSEPROGRAM)(uint32_t program);
typedef bool(STDCALL *PFNGLISPROGRAM)(uint32_t program);
typedef void(STDCALL *PFNGLVALIDATEPROGRAM)(uint32_t program);
typedef void(STDCALL *PFNGLCLEARCOLOR)(float r, float g, float b, float a);
typedef void(STDCALL *PFNGLCLEARDEPTHF)(float depth);
typedef void(STDCALL *PFNGLCLEARSTENCIL)(int32_t stencil);
typedef void(STDCALL *PFNGLCLEAR)(ClearMask mask);
typedef void(STDCALL *PFNGLCULLFACE)(FaceMode mode);
typedef void(STDCALL *PFNGLPOLYGONOFFSET)(float scale_factor, float units);
typedef void(STDCALL *PFNGLLINEWIDTH)(float width);
typedef void(STDCALL *PFNGLSAMPLECOVERAGE)(float value, bool invert);
typedef void(STDCALL *PFNGLHINT)(HintTarget target, HintMode mode);
typedef void(STDCALL *PFNGLFRAMEBUFFERRENDERBUFFER)(FramebufferTarget framebuffer_target,
                                                    FramebufferAttachment framebuffer_attachment,
                                                    RenderbufferTarget renderbuffer_target,
                                                    uint32_t renderbuffer);
typedef void(STDCALL *PFNGLFRAMEBUFFERTEXTURE2D)(FramebufferTarget framebuffer_target,
                                                 FramebufferAttachment framebuffer_attachment,
                                                 TextureImageTarget texture_target,
                                                 uint32_t texture, int32_t level);
typedef void(STDCALL *PFNGLGETFRAMEBUFFERATTACHMENTPARAMETERIV)(
        FramebufferTarget framebuffer_target, FramebufferAttachment attachment,
        FramebufferAttachmentParameter parameter, int32_t *value);
typedef void(STDCALL *PFNGLDRAWELEMENTS)(DrawMode draw_mode, int32_t element_count,
                                         IndicesType indices_type, void *indices);
typedef void(STDCALL *PFNGLDRAWARRAYS)(DrawMode draw_mode, int32_t first_index,
                                       int32_t index_count);
typedef void(STDCALL *PFNGLFLUSH)();
typedef void(STDCALL *PFNGLFINISH)();
typedef void(STDCALL *PFNGLGETBOOLEANV)(StateVariable param, bool *values);
typedef void(STDCALL *PFNGLGETFLOATV)(StateVariable param, float *values);
typedef void(STDCALL *PFNGLGETINTEGERV)(StateVariable param, int32_t *values);
typedef char *(STDCALL *PFNGLGETSTRING)(StringConstant param);
typedef void(STDCALL *PFNGLENABLE)(Capability capability);
typedef void(STDCALL *PFNGLDISABLE)(Capability capability);
typedef bool(STDCALL *PFNGLISENABLED)(Capability capability);
typedef void *(STDCALL *PFNGLMAPBUFFERRANGE)(BufferTarget target, int32_t offset, int32_t length,
                                             MapBufferRangeAccess access);
typedef void(STDCALL *PFNGLUNMAPBUFFER)(BufferTarget target);
typedef void(STDCALL *PFNGLINVALIDATEFRAMEBUFFER)(FramebufferTarget target, int32_t count,
                                                  FramebufferAttachment *attachments);
typedef void(STDCALL *PFNGLRENDERBUFFERSTORAGEMULTISAMPLE)(RenderbufferTarget target,
                                                           int32_t samples,
                                                           RenderbufferFormat format, int32_t width,
                                                           int32_t height);
typedef void(STDCALL *PFNGLBLITFRAMEBUFFER)(int32_t srcX0, int32_t srcY0, int32_t srcX1,
                                            int32_t srcY1, int32_t dstX0, int32_t dstY0,
                                            int32_t dstX1, int32_t dstY1, ClearMask mask,
                                            TextureFilterMode filter);
typedef void(STDCALL *PFNGLGENQUERIES)(int32_t count, uint32_t *queries);
typedef void(STDCALL *PFNGLBEGINQUERY)(QueryTarget target, uint32_t query);
typedef void(STDCALL *PFNGLENDQUERY)(QueryTarget target);
typedef void(STDCALL *PFNGLDELETEQUERIES)(int32_t count, uint32_t *queries);
typedef bool(STDCALL *PFNGLISQUERY)(uint32_t query);
typedef void(STDCALL *PFNGLGETQUERYIV)(QueryTarget target, QueryParameter parameter,
                                       int32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTUIV)(uint32_t query, QueryObjectParameter parameter,
                                              uint32_t *value);
typedef void(STDCALL *PFNGLGETACTIVEUNIFORMBLOCKNAME)(uint32_t program,
                                                      uint32_t uniform_block_index,
                                                      int32_t buffer_size,
                                                      int32_t *buffer_bytes_written, char *name);
typedef void(STDCALL *PFNGLGETACTIVEUNIFORMBLOCKIV)(uint32_t program, uint32_t uniform_block_index,
                                                    UniformBlockParameter parameter_name,
                                                    int32_t *parameters);
typedef void(STDCALL *PFNGLUNIFORMBLOCKBINDING)(uint32_t program, uint32_t uniform_block_index,
                                                uint32_t uniform_block_binding);
typedef void(STDCALL *PFNGLGETACTIVEUNIFORMSIV)(uint32_t program, uint32_t uniform_count,
                                                uint32_t *uniform_indices,
                                                UniformBlockParameter parameter_name,
                                                int32_t *parameters);
typedef void(STDCALL *PFNGLBINDBUFFERBASE)(IndexedBufferTarget target, uint32_t index,
                                           uint32_t buffer);
typedef void(STDCALL *PFNGLGENVERTEXARRAYS)(int32_t count, uint32_t *arrays);
typedef void(STDCALL *PFNGLBINDVERTEXARRAY)(uint32_t array);
typedef void(STDCALL *PFNGLDELETEVERTEXARRAYS)(uint32_t count, uint32_t *arrays);
typedef void(STDCALL *PFNGLGENQUERIESEXT)(int32_t count, uint32_t *queries);
typedef void(STDCALL *PFNGLBEGINQUERYEXT)(QueryTarget target, uint32_t query);
typedef void(STDCALL *PFNGLENDQUERYEXT)(QueryTarget target);
typedef void(STDCALL *PFNGLDELETEQUERIESEXT)(int32_t count, uint32_t *queries);
typedef bool(STDCALL *PFNGLISQUERYEXT)(uint32_t query);
typedef void(STDCALL *PFNGLQUERYCOUNTEREXT)(uint32_t query, QueryTarget target);
typedef void(STDCALL *PFNGLGETQUERYIVEXT)(QueryTarget target, QueryParameter parameter,
                                          int32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTIVEXT)(uint32_t query, QueryObjectParameter parameter,
                                                int32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTUIVEXT)(uint32_t query, QueryObjectParameter parameter,
                                                 uint32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTI64VEXT)(uint32_t query, QueryObjectParameter parameter,
                                                  int64_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTUI64VEXT)(uint32_t query, QueryObjectParameter parameter,
                                                   uint64_t *value);

extern PFNEGLINITIALIZE eglInitialize;
extern PFNEGLCREATECONTEXT eglCreateContext;
extern PFNEGLMAKECURRENT eglMakeCurrent;
extern PFNEGLSWAPBUFFERS eglSwapBuffers;
extern PFNEGLQUERYSURFACE eglQuerySurface;
extern PFNGLXCREATECONTEXT glXCreateContext;
extern PFNGLXCREATENEWCONTEXT glXCreateNewContext;
extern PFNGLXMAKECONTEXTCURRENT glXMakeContextCurrent;
extern PFNGLXMAKECURRENT glXMakeCurrent;
extern PFNGLXSWAPBUFFERS glXSwapBuffers;
extern PFNGLXQUERYDRAWABLE glXQueryDrawable;
extern PFNWGLCREATECONTEXT wglCreateContext;
extern PFNWGLCREATECONTEXTATTRIBSARB wglCreateContextAttribsARB;
extern PFNWGLMAKECURRENT wglMakeCurrent;
extern PFNWGLSWAPBUFFERS wglSwapBuffers;
extern PFNCGLCREATECONTEXT CGLCreateContext;
extern PFNCGLSETCURRENTCONTEXT CGLSetCurrentContext;
extern PFNCGLGETSURFACE CGLGetSurface;
extern PFNCGSGETSURFACEBOUNDS CGSGetSurfaceBounds;
extern PFNCGLFLUSHDRAWABLE CGLFlushDrawable;
extern PFNGLENABLECLIENTSTATE glEnableClientState;
extern PFNGLDISABLECLIENTSTATE glDisableClientState;
extern PFNGLGETPROGRAMBINARYOES glGetProgramBinaryOES;
extern PFNGLPROGRAMBINARYOES glProgramBinaryOES;
extern PFNGLSTARTTILINGQCOM glStartTilingQCOM;
extern PFNGLENDTILINGQCOM glEndTilingQCOM;
extern PFNGLDISCARDFRAMEBUFFEREXT glDiscardFramebufferEXT;
extern PFNGLINSERTEVENTMARKEREXT glInsertEventMarkerEXT;
extern PFNGLPUSHGROUPMARKEREXT glPushGroupMarkerEXT;
extern PFNGLPOPGROUPMARKEREXT glPopGroupMarkerEXT;
extern PFNGLTEXSTORAGE1DEXT glTexStorage1DEXT;
extern PFNGLTEXSTORAGE2DEXT glTexStorage2DEXT;
extern PFNGLTEXSTORAGE3DEXT glTexStorage3DEXT;
extern PFNGLTEXTURESTORAGE1DEXT glTextureStorage1DEXT;
extern PFNGLTEXTURESTORAGE2DEXT glTextureStorage2DEXT;
extern PFNGLTEXTURESTORAGE3DEXT glTextureStorage3DEXT;
extern PFNGLGENVERTEXARRAYSOES glGenVertexArraysOES;
extern PFNGLBINDVERTEXARRAYOES glBindVertexArrayOES;
extern PFNGLDELETEVERTEXARRAYSOES glDeleteVertexArraysOES;
extern PFNGLISVERTEXARRAYOES glIsVertexArrayOES;
extern PFNGLEGLIMAGETARGETTEXTURE2DOES glEGLImageTargetTexture2DOES;
extern PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOES glEGLImageTargetRenderbufferStorageOES;
extern PFNGLGETGRAPHICSRESETSTATUSEXT glGetGraphicsResetStatusEXT;
extern PFNGLBINDATTRIBLOCATION glBindAttribLocation;
extern PFNGLBLENDFUNC glBlendFunc;
extern PFNGLBLENDFUNCSEPARATE glBlendFuncSeparate;
extern PFNGLBLENDEQUATION glBlendEquation;
extern PFNGLBLENDEQUATIONSEPARATE glBlendEquationSeparate;
extern PFNGLBLENDCOLOR glBlendColor;
extern PFNGLENABLEVERTEXATTRIBARRAY glEnableVertexAttribArray;
extern PFNGLDISABLEVERTEXATTRIBARRAY glDisableVertexAttribArray;
extern PFNGLVERTEXATTRIBPOINTER glVertexAttribPointer;
extern PFNGLGETACTIVEATTRIB glGetActiveAttrib;
extern PFNGLGETACTIVEUNIFORM glGetActiveUniform;
extern PFNGLGETERROR glGetError;
extern PFNGLGETPROGRAMIV glGetProgramiv;
extern PFNGLGETSHADERIV glGetShaderiv;
extern PFNGLGETUNIFORMLOCATION glGetUniformLocation;
extern PFNGLGETATTRIBLOCATION glGetAttribLocation;
extern PFNGLPIXELSTOREI glPixelStorei;
extern PFNGLTEXPARAMETERI glTexParameteri;
extern PFNGLTEXPARAMETERF glTexParameterf;
extern PFNGLGETTEXPARAMETERIV glGetTexParameteriv;
extern PFNGLGETTEXPARAMETERFV glGetTexParameterfv;
extern PFNGLUNIFORM1I glUniform1i;
extern PFNGLUNIFORM2I glUniform2i;
extern PFNGLUNIFORM3I glUniform3i;
extern PFNGLUNIFORM4I glUniform4i;
extern PFNGLUNIFORM1IV glUniform1iv;
extern PFNGLUNIFORM2IV glUniform2iv;
extern PFNGLUNIFORM3IV glUniform3iv;
extern PFNGLUNIFORM4IV glUniform4iv;
extern PFNGLUNIFORM1F glUniform1f;
extern PFNGLUNIFORM2F glUniform2f;
extern PFNGLUNIFORM3F glUniform3f;
extern PFNGLUNIFORM4F glUniform4f;
extern PFNGLUNIFORM1FV glUniform1fv;
extern PFNGLUNIFORM2FV glUniform2fv;
extern PFNGLUNIFORM3FV glUniform3fv;
extern PFNGLUNIFORM4FV glUniform4fv;
extern PFNGLUNIFORMMATRIX2FV glUniformMatrix2fv;
extern PFNGLUNIFORMMATRIX3FV glUniformMatrix3fv;
extern PFNGLUNIFORMMATRIX4FV glUniformMatrix4fv;
extern PFNGLGETUNIFORMFV glGetUniformfv;
extern PFNGLGETUNIFORMIV glGetUniformiv;
extern PFNGLVERTEXATTRIB1F glVertexAttrib1f;
extern PFNGLVERTEXATTRIB2F glVertexAttrib2f;
extern PFNGLVERTEXATTRIB3F glVertexAttrib3f;
extern PFNGLVERTEXATTRIB4F glVertexAttrib4f;
extern PFNGLVERTEXATTRIB1FV glVertexAttrib1fv;
extern PFNGLVERTEXATTRIB2FV glVertexAttrib2fv;
extern PFNGLVERTEXATTRIB3FV glVertexAttrib3fv;
extern PFNGLVERTEXATTRIB4FV glVertexAttrib4fv;
extern PFNGLGETSHADERPRECISIONFORMAT glGetShaderPrecisionFormat;
extern PFNGLDEPTHMASK glDepthMask;
extern PFNGLDEPTHFUNC glDepthFunc;
extern PFNGLDEPTHRANGEF glDepthRangef;
extern PFNGLCOLORMASK glColorMask;
extern PFNGLSTENCILMASK glStencilMask;
extern PFNGLSTENCILMASKSEPARATE glStencilMaskSeparate;
extern PFNGLSTENCILFUNCSEPARATE glStencilFuncSeparate;
extern PFNGLSTENCILOPSEPARATE glStencilOpSeparate;
extern PFNGLFRONTFACE glFrontFace;
extern PFNGLVIEWPORT glViewport;
extern PFNGLSCISSOR glScissor;
extern PFNGLACTIVETEXTURE glActiveTexture;
extern PFNGLGENTEXTURES glGenTextures;
extern PFNGLDELETETEXTURES glDeleteTextures;
extern PFNGLISTEXTURE glIsTexture;
extern PFNGLBINDTEXTURE glBindTexture;
extern PFNGLTEXIMAGE2D glTexImage2D;
extern PFNGLTEXSUBIMAGE2D glTexSubImage2D;
extern PFNGLCOPYTEXIMAGE2D glCopyTexImage2D;
extern PFNGLCOPYTEXSUBIMAGE2D glCopyTexSubImage2D;
extern PFNGLCOMPRESSEDTEXIMAGE2D glCompressedTexImage2D;
extern PFNGLCOMPRESSEDTEXSUBIMAGE2D glCompressedTexSubImage2D;
extern PFNGLGENERATEMIPMAP glGenerateMipmap;
extern PFNGLREADPIXELS glReadPixels;
extern PFNGLGENFRAMEBUFFERS glGenFramebuffers;
extern PFNGLBINDFRAMEBUFFER glBindFramebuffer;
extern PFNGLCHECKFRAMEBUFFERSTATUS glCheckFramebufferStatus;
extern PFNGLDELETEFRAMEBUFFERS glDeleteFramebuffers;
extern PFNGLISFRAMEBUFFER glIsFramebuffer;
extern PFNGLGENRENDERBUFFERS glGenRenderbuffers;
extern PFNGLBINDRENDERBUFFER glBindRenderbuffer;
extern PFNGLRENDERBUFFERSTORAGE glRenderbufferStorage;
extern PFNGLDELETERENDERBUFFERS glDeleteRenderbuffers;
extern PFNGLISRENDERBUFFER glIsRenderbuffer;
extern PFNGLGETRENDERBUFFERPARAMETERIV glGetRenderbufferParameteriv;
extern PFNGLGENBUFFERS glGenBuffers;
extern PFNGLBINDBUFFER glBindBuffer;
extern PFNGLBUFFERDATA glBufferData;
extern PFNGLBUFFERSUBDATA glBufferSubData;
extern PFNGLDELETEBUFFERS glDeleteBuffers;
extern PFNGLISBUFFER glIsBuffer;
extern PFNGLGETBUFFERPARAMETERIV glGetBufferParameteriv;
extern PFNGLCREATESHADER glCreateShader;
extern PFNGLDELETESHADER glDeleteShader;
extern PFNGLSHADERSOURCE glShaderSource;
extern PFNGLSHADERBINARY glShaderBinary;
extern PFNGLGETSHADERINFOLOG glGetShaderInfoLog;
extern PFNGLGETSHADERSOURCE glGetShaderSource;
extern PFNGLRELEASESHADERCOMPILER glReleaseShaderCompiler;
extern PFNGLCOMPILESHADER glCompileShader;
extern PFNGLISSHADER glIsShader;
extern PFNGLCREATEPROGRAM glCreateProgram;
extern PFNGLDELETEPROGRAM glDeleteProgram;
extern PFNGLATTACHSHADER glAttachShader;
extern PFNGLDETACHSHADER glDetachShader;
extern PFNGLGETATTACHEDSHADERS glGetAttachedShaders;
extern PFNGLLINKPROGRAM glLinkProgram;
extern PFNGLGETPROGRAMINFOLOG glGetProgramInfoLog;
extern PFNGLUSEPROGRAM glUseProgram;
extern PFNGLISPROGRAM glIsProgram;
extern PFNGLVALIDATEPROGRAM glValidateProgram;
extern PFNGLCLEARCOLOR glClearColor;
extern PFNGLCLEARDEPTHF glClearDepthf;
extern PFNGLCLEARSTENCIL glClearStencil;
extern PFNGLCLEAR glClear;
extern PFNGLCULLFACE glCullFace;
extern PFNGLPOLYGONOFFSET glPolygonOffset;
extern PFNGLLINEWIDTH glLineWidth;
extern PFNGLSAMPLECOVERAGE glSampleCoverage;
extern PFNGLHINT glHint;
extern PFNGLFRAMEBUFFERRENDERBUFFER glFramebufferRenderbuffer;
extern PFNGLFRAMEBUFFERTEXTURE2D glFramebufferTexture2D;
extern PFNGLGETFRAMEBUFFERATTACHMENTPARAMETERIV glGetFramebufferAttachmentParameteriv;
extern PFNGLDRAWELEMENTS glDrawElements;
extern PFNGLDRAWARRAYS glDrawArrays;
extern PFNGLFLUSH glFlush;
extern PFNGLFINISH glFinish;
extern PFNGLGETBOOLEANV glGetBooleanv;
extern PFNGLGETFLOATV glGetFloatv;
extern PFNGLGETINTEGERV glGetIntegerv;
extern PFNGLGETSTRING glGetString;
extern PFNGLENABLE glEnable;
extern PFNGLDISABLE glDisable;
extern PFNGLISENABLED glIsEnabled;
extern PFNGLMAPBUFFERRANGE glMapBufferRange;
extern PFNGLUNMAPBUFFER glUnmapBuffer;
extern PFNGLINVALIDATEFRAMEBUFFER glInvalidateFramebuffer;
extern PFNGLRENDERBUFFERSTORAGEMULTISAMPLE glRenderbufferStorageMultisample;
extern PFNGLBLITFRAMEBUFFER glBlitFramebuffer;
extern PFNGLGENQUERIES glGenQueries;
extern PFNGLBEGINQUERY glBeginQuery;
extern PFNGLENDQUERY glEndQuery;
extern PFNGLDELETEQUERIES glDeleteQueries;
extern PFNGLISQUERY glIsQuery;
extern PFNGLGETQUERYIV glGetQueryiv;
extern PFNGLGETQUERYOBJECTUIV glGetQueryObjectuiv;
extern PFNGLGETACTIVEUNIFORMBLOCKNAME glGetActiveUniformBlockName;
extern PFNGLGETACTIVEUNIFORMBLOCKIV glGetActiveUniformBlockiv;
extern PFNGLUNIFORMBLOCKBINDING glUniformBlockBinding;
extern PFNGLGETACTIVEUNIFORMSIV glGetActiveUniformsiv;
extern PFNGLBINDBUFFERBASE glBindBufferBase;
extern PFNGLGENVERTEXARRAYS glGenVertexArrays;
extern PFNGLBINDVERTEXARRAY glBindVertexArray;
extern PFNGLDELETEVERTEXARRAYS glDeleteVertexArrays;
extern PFNGLGENQUERIESEXT glGenQueriesEXT;
extern PFNGLBEGINQUERYEXT glBeginQueryEXT;
extern PFNGLENDQUERYEXT glEndQueryEXT;
extern PFNGLDELETEQUERIESEXT glDeleteQueriesEXT;
extern PFNGLISQUERYEXT glIsQueryEXT;
extern PFNGLQUERYCOUNTEREXT glQueryCounterEXT;
extern PFNGLGETQUERYIVEXT glGetQueryivEXT;
extern PFNGLGETQUERYOBJECTIVEXT glGetQueryObjectivEXT;
extern PFNGLGETQUERYOBJECTUIVEXT glGetQueryObjectuivEXT;
extern PFNGLGETQUERYOBJECTI64VEXT glGetQueryObjecti64vEXT;
extern PFNGLGETQUERYOBJECTUI64VEXT glGetQueryObjectui64vEXT;

}  // namespace gfxapi
}  // namespace gapir

#endif  // GAPIR_GFX_API_H
