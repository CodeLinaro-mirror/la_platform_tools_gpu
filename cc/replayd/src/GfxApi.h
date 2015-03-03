////////////////////////////////////////////////////////////////////////////////
// Automatically generated file. Do not modify!
////////////////////////////////////////////////////////////////////////////////

#ifndef ANDROID_CAZE_GFXAPI_FUNCTIONS_H
#define ANDROID_CAZE_GFXAPI_FUNCTIONS_H

#include "Target.h"

#include <stdint.h>

namespace android {
namespace caze {

class Interpreter;

namespace gfxapi {

// Register the API functions with the interpreter
extern void Register(Interpreter *interpreter);

// Look-up the API function addresses for the currently bound context.
extern void Initialize();

namespace Ids {
// List of the function ids for the API functions. This list have to be consistent with the
// function ids on the server side because they are part of the communication protocol
static const uint16_t Init = 0;
static const uint16_t StartTimer = 1;
static const uint16_t StopTimer = 2;
static const uint16_t FlushPostBuffer = 3;
static const uint16_t EglCreateContext = 4;
static const uint16_t EglMakeCurrent = 5;
static const uint16_t EglSwapBuffers = 6;
static const uint16_t GlEnableClientState = 7;
static const uint16_t GlDisableClientState = 8;
static const uint16_t GlGetProgramBinaryOES = 9;
static const uint16_t GlProgramBinaryOES = 10;
static const uint16_t GlStartTilingQCOM = 11;
static const uint16_t GlEndTilingQCOM = 12;
static const uint16_t GlDiscardFramebufferEXT = 13;
static const uint16_t GlInsertEventMarkerEXT = 14;
static const uint16_t GlPushGroupMarkerEXT = 15;
static const uint16_t GlPopGroupMarkerEXT = 16;
static const uint16_t GlTexStorage1DEXT = 17;
static const uint16_t GlTexStorage2DEXT = 18;
static const uint16_t GlTexStorage3DEXT = 19;
static const uint16_t GlTextureStorage1DEXT = 20;
static const uint16_t GlTextureStorage2DEXT = 21;
static const uint16_t GlTextureStorage3DEXT = 22;
static const uint16_t GlGenVertexArraysOES = 23;
static const uint16_t GlBindVertexArrayOES = 24;
static const uint16_t GlDeleteVertexArraysOES = 25;
static const uint16_t GlIsVertexArrayOES = 26;
static const uint16_t GlEGLImageTargetTexture2DOES = 27;
static const uint16_t GlEGLImageTargetRenderbufferStorageOES = 28;
static const uint16_t GlGetGraphicsResetStatusEXT = 29;
static const uint16_t GlBindAttribLocation = 30;
static const uint16_t GlBlendFunc = 31;
static const uint16_t GlBlendFuncSeparate = 32;
static const uint16_t GlBlendEquation = 33;
static const uint16_t GlBlendEquationSeparate = 34;
static const uint16_t GlBlendColor = 35;
static const uint16_t GlEnableVertexAttribArray = 36;
static const uint16_t GlDisableVertexAttribArray = 37;
static const uint16_t GlVertexAttribPointer = 38;
static const uint16_t GlGetActiveAttrib = 39;
static const uint16_t GlGetActiveUniform = 40;
static const uint16_t GlGetError = 41;
static const uint16_t GlGetProgramiv = 42;
static const uint16_t GlGetShaderiv = 43;
static const uint16_t GlGetUniformLocation = 44;
static const uint16_t GlGetAttribLocation = 45;
static const uint16_t GlPixelStorei = 46;
static const uint16_t GlTexParameteri = 47;
static const uint16_t GlTexParameterf = 48;
static const uint16_t GlGetTexParameteriv = 49;
static const uint16_t GlGetTexParameterfv = 50;
static const uint16_t GlUniform1i = 51;
static const uint16_t GlUniform2i = 52;
static const uint16_t GlUniform3i = 53;
static const uint16_t GlUniform4i = 54;
static const uint16_t GlUniform1iv = 55;
static const uint16_t GlUniform2iv = 56;
static const uint16_t GlUniform3iv = 57;
static const uint16_t GlUniform4iv = 58;
static const uint16_t GlUniform1f = 59;
static const uint16_t GlUniform2f = 60;
static const uint16_t GlUniform3f = 61;
static const uint16_t GlUniform4f = 62;
static const uint16_t GlUniform1fv = 63;
static const uint16_t GlUniform2fv = 64;
static const uint16_t GlUniform3fv = 65;
static const uint16_t GlUniform4fv = 66;
static const uint16_t GlUniformMatrix2fv = 67;
static const uint16_t GlUniformMatrix3fv = 68;
static const uint16_t GlUniformMatrix4fv = 69;
static const uint16_t GlGetUniformfv = 70;
static const uint16_t GlGetUniformiv = 71;
static const uint16_t GlVertexAttrib1f = 72;
static const uint16_t GlVertexAttrib2f = 73;
static const uint16_t GlVertexAttrib3f = 74;
static const uint16_t GlVertexAttrib4f = 75;
static const uint16_t GlVertexAttrib1fv = 76;
static const uint16_t GlVertexAttrib2fv = 77;
static const uint16_t GlVertexAttrib3fv = 78;
static const uint16_t GlVertexAttrib4fv = 79;
static const uint16_t GlGetShaderPrecisionFormat = 80;
static const uint16_t GlDepthMask = 81;
static const uint16_t GlDepthFunc = 82;
static const uint16_t GlDepthRangef = 83;
static const uint16_t GlColorMask = 84;
static const uint16_t GlStencilMask = 85;
static const uint16_t GlStencilMaskSeparate = 86;
static const uint16_t GlStencilFuncSeparate = 87;
static const uint16_t GlStencilOpSeparate = 88;
static const uint16_t GlFrontFace = 89;
static const uint16_t GlViewport = 90;
static const uint16_t GlScissor = 91;
static const uint16_t GlActiveTexture = 92;
static const uint16_t GlGenTextures = 93;
static const uint16_t GlDeleteTextures = 94;
static const uint16_t GlIsTexture = 95;
static const uint16_t GlBindTexture = 96;
static const uint16_t GlTexImage2D = 97;
static const uint16_t GlTexSubImage2D = 98;
static const uint16_t GlCopyTexImage2D = 99;
static const uint16_t GlCopyTexSubImage2D = 100;
static const uint16_t GlCompressedTexImage2D = 101;
static const uint16_t GlCompressedTexSubImage2D = 102;
static const uint16_t GlGenerateMipmap = 103;
static const uint16_t GlReadPixels = 104;
static const uint16_t GlGenFramebuffers = 105;
static const uint16_t GlBindFramebuffer = 106;
static const uint16_t GlCheckFramebufferStatus = 107;
static const uint16_t GlDeleteFramebuffers = 108;
static const uint16_t GlIsFramebuffer = 109;
static const uint16_t GlGenRenderbuffers = 110;
static const uint16_t GlBindRenderbuffer = 111;
static const uint16_t GlRenderbufferStorage = 112;
static const uint16_t GlDeleteRenderbuffers = 113;
static const uint16_t GlIsRenderbuffer = 114;
static const uint16_t GlGetRenderbufferParameteriv = 115;
static const uint16_t GlGenBuffers = 116;
static const uint16_t GlBindBuffer = 117;
static const uint16_t GlBufferData = 118;
static const uint16_t GlBufferSubData = 119;
static const uint16_t GlDeleteBuffers = 120;
static const uint16_t GlIsBuffer = 121;
static const uint16_t GlGetBufferParameteriv = 122;
static const uint16_t GlCreateShader = 123;
static const uint16_t GlDeleteShader = 124;
static const uint16_t GlShaderSource = 125;
static const uint16_t GlShaderBinary = 126;
static const uint16_t GlGetShaderInfoLog = 127;
static const uint16_t GlGetShaderSource = 128;
static const uint16_t GlReleaseShaderCompiler = 129;
static const uint16_t GlCompileShader = 130;
static const uint16_t GlIsShader = 131;
static const uint16_t GlCreateProgram = 132;
static const uint16_t GlDeleteProgram = 133;
static const uint16_t GlAttachShader = 134;
static const uint16_t GlDetachShader = 135;
static const uint16_t GlGetAttachedShaders = 136;
static const uint16_t GlLinkProgram = 137;
static const uint16_t GlGetProgramInfoLog = 138;
static const uint16_t GlUseProgram = 139;
static const uint16_t GlIsProgram = 140;
static const uint16_t GlValidateProgram = 141;
static const uint16_t GlClearColor = 142;
static const uint16_t GlClearDepthf = 143;
static const uint16_t GlClearStencil = 144;
static const uint16_t GlClear = 145;
static const uint16_t GlCullFace = 146;
static const uint16_t GlPolygonOffset = 147;
static const uint16_t GlLineWidth = 148;
static const uint16_t GlSampleCoverage = 149;
static const uint16_t GlHint = 150;
static const uint16_t GlFramebufferRenderbuffer = 151;
static const uint16_t GlFramebufferTexture2D = 152;
static const uint16_t GlGetFramebufferAttachmentParameteriv = 153;
static const uint16_t GlDrawElements = 154;
static const uint16_t GlDrawArrays = 155;
static const uint16_t GlFlush = 156;
static const uint16_t GlFinish = 157;
static const uint16_t GlGetBooleanv = 158;
static const uint16_t GlGetFloatv = 159;
static const uint16_t GlGetIntegerv = 160;
static const uint16_t GlGetString = 161;
static const uint16_t GlEnable = 162;
static const uint16_t GlDisable = 163;
static const uint16_t GlIsEnabled = 164;
static const uint16_t GlMapBufferRange = 165;
static const uint16_t GlUnmapBuffer = 166;
static const uint16_t GlInvalidateFramebuffer = 167;
static const uint16_t GlRenderbufferStorageMultisample = 168;
static const uint16_t GlBlitFramebuffer = 169;
static const uint16_t GlGenQueries = 170;
static const uint16_t GlBeginQuery = 171;
static const uint16_t GlEndQuery = 172;
static const uint16_t GlDeleteQueries = 173;
static const uint16_t GlIsQuery = 174;
static const uint16_t GlGetQueryiv = 175;
static const uint16_t GlGetQueryObjectuiv = 176;
static const uint16_t GlGenQueriesEXT = 177;
static const uint16_t GlBeginQueryEXT = 178;
static const uint16_t GlEndQueryEXT = 179;
static const uint16_t GlDeleteQueriesEXT = 180;
static const uint16_t GlIsQueryEXT = 181;
static const uint16_t GlQueryCounterEXT = 182;
static const uint16_t GlGetQueryivEXT = 183;
static const uint16_t GlGetQueryObjectivEXT = 184;
static const uint16_t GlGetQueryObjectuivEXT = 185;
static const uint16_t GlGetQueryObjecti64vEXT = 186;
static const uint16_t GlGetQueryObjectui64vEXT = 187;
}  // end of namespace FunctionIds

namespace DrawMode {
static const uint32_t GL_LINE_LOOP = 2;
static const uint32_t GL_LINE_STRIP = 3;
static const uint32_t GL_LINES = 1;
static const uint32_t GL_POINTS = 0;
static const uint32_t GL_TRIANGLE_FAN = 6;
static const uint32_t GL_TRIANGLE_STRIP = 5;
static const uint32_t GL_TRIANGLES = 4;
};  // end of namespace DrawMode

namespace IndicesType {
static const uint32_t GL_UNSIGNED_BYTE = 5121;
static const uint32_t GL_UNSIGNED_SHORT = 5123;
static const uint32_t GL_UNSIGNED_INT = 5125;
};  // end of namespace IndicesType

namespace TextureTarget_GLES_1_1 {
static const uint32_t GL_TEXTURE_2D = 3553;
};  // end of namespace TextureTarget_GLES_1_1

namespace TextureTarget_GLES_2_0 {
static const uint32_t GL_TEXTURE_CUBE_MAP = 34067;
};  // end of namespace TextureTarget_GLES_2_0

namespace TextureTarget_OES_EGL_image_external {
static const uint32_t GL_TEXTURE_EXTERNAL_OES = 36197;
};  // end of namespace TextureTarget_OES_EGL_image_external

namespace TextureTarget {};  // end of namespace TextureTarget

namespace CubeMapImageTarget {
static const uint32_t GL_TEXTURE_CUBE_MAP_NEGATIVE_X = 34070;
static const uint32_t GL_TEXTURE_CUBE_MAP_NEGATIVE_Y = 34072;
static const uint32_t GL_TEXTURE_CUBE_MAP_NEGATIVE_Z = 34074;
static const uint32_t GL_TEXTURE_CUBE_MAP_POSITIVE_X = 34069;
static const uint32_t GL_TEXTURE_CUBE_MAP_POSITIVE_Y = 34071;
static const uint32_t GL_TEXTURE_CUBE_MAP_POSITIVE_Z = 34073;
};  // end of namespace CubeMapImageTarget

namespace Texture2DImageTarget {
static const uint32_t GL_TEXTURE_2D = 3553;
};  // end of namespace Texture2DImageTarget

namespace TextureImageTarget {};  // end of namespace TextureImageTarget

namespace BaseTexelFormat {
static const uint32_t GL_ALPHA = 6406;
static const uint32_t GL_RGB = 6407;
static const uint32_t GL_RGBA = 6408;
};  // end of namespace BaseTexelFormat

namespace TexelFormat_GLES_1_1 {
static const uint32_t GL_LUMINANCE = 6409;
static const uint32_t GL_LUMINANCE_ALPHA = 6410;
};  // end of namespace TexelFormat_GLES_1_1

namespace TexelFormat_GLES_3_0 {
static const uint32_t GL_RED = 6403;
static const uint32_t GL_RED_INTEGER = 36244;
static const uint32_t GL_RG = 33319;
static const uint32_t GL_RG_INTEGER = 33320;
static const uint32_t GL_RGB_INTEGER = 36248;
static const uint32_t GL_RGBA_INTEGER = 36249;
static const uint32_t GL_DEPTH_COMPONENT = 6402;
static const uint32_t GL_DEPTH_STENCIL = 34041;
};  // end of namespace TexelFormat_GLES_3_0

namespace TexelFormat {};  // end of namespace TexelFormat

namespace RenderbufferFormat {
static const uint32_t GL_RGBA4 = 32854;
static const uint32_t GL_RGB5_A1 = 32855;
static const uint32_t GL_RGB565 = 36194;
static const uint32_t GL_RGBA8 = 32856;
static const uint32_t GL_DEPTH_COMPONENT16 = 33189;
static const uint32_t GL_STENCIL_INDEX8 = 36168;
};  // end of namespace RenderbufferFormat

namespace Type_OES_vertex_half_float {
static const uint32_t GL_HALF_FLOAT_OES = 36193;
};  // end of namespace Type_OES_vertex_half_float

namespace CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture {
static const uint32_t GL_ETC1_RGB8_OES = 36196;
};  // end of namespace CompressedTexelFormat_OES_compressed_ETC1_RGB8_texture

namespace CompressedTexelFormat_AMD_compressed_ATC_texture {
static const uint32_t GL_ATC_RGB_AMD = 35986;
static const uint32_t GL_ATC_RGBA_EXPLICIT_ALPHA_AMD = 35987;
static const uint32_t GL_ATC_RGBA_INTERPOLATED_ALPHA_AMD = 34798;
};  // end of namespace CompressedTexelFormat_AMD_compressed_ATC_texture

namespace CompressedTexelFormat {};  // end of namespace CompressedTexelFormat

namespace ImageTexelFormat {};  // end of namespace ImageTexelFormat

namespace TexelType {
static const uint32_t GL_UNSIGNED_BYTE = 5121;
static const uint32_t GL_UNSIGNED_SHORT_4_4_4_4 = 32819;
static const uint32_t GL_UNSIGNED_SHORT_5_5_5_1 = 32820;
static const uint32_t GL_UNSIGNED_SHORT_5_6_5 = 33635;
static const uint32_t GL_FLOAT = 5126;
};  // end of namespace TexelType

namespace FramebufferAttachment {
static const uint32_t GL_COLOR_ATTACHMENT0 = 36064;
static const uint32_t GL_DEPTH_ATTACHMENT = 36096;
static const uint32_t GL_STENCIL_ATTACHMENT = 36128;
};  // end of namespace FramebufferAttachment

namespace FramebufferAttachmentType {
static const uint32_t GL_NONE = 0;
static const uint32_t GL_RENDERBUFFER = 36161;
static const uint32_t GL_TEXTURE = 5890;
};  // end of namespace FramebufferAttachmentType

namespace FramebufferTarget_GLES_2_0 {
static const uint32_t GL_FRAMEBUFFER = 36160;
};  // end of namespace FramebufferTarget_GLES_2_0

namespace FramebufferTarget_GLES_3_1 {
static const uint32_t GL_READ_FRAMEBUFFER = 36008;
static const uint32_t GL_DRAW_FRAMEBUFFER = 36009;
};  // end of namespace FramebufferTarget_GLES_3_1

namespace FramebufferTarget {};  // end of namespace FramebufferTarget

namespace FramebufferAttachmentParameter {
static const uint32_t GL_FRAMEBUFFER_ATTACHMENT_OBJECT_TYPE = 36048;
static const uint32_t GL_FRAMEBUFFER_ATTACHMENT_OBJECT_NAME = 36049;
static const uint32_t GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_LEVEL = 36050;
static const uint32_t GL_FRAMEBUFFER_ATTACHMENT_TEXTURE_CUBE_MAP_FACE = 36051;
};  // end of namespace FramebufferAttachmentParameter

namespace FramebufferStatus {
static const uint32_t GL_FRAMEBUFFER_COMPLETE = 36053;
static const uint32_t GL_FRAMEBUFFER_INCOMPLETE_ATTACHMENT = 36054;
static const uint32_t GL_FRAMEBUFFER_INCOMPLETE_MISSING_ATTACHMENT = 36055;
static const uint32_t GL_FRAMEBUFFER_INCOMPLETE_DIMENSIONS = 36057;
static const uint32_t GL_FRAMEBUFFER_UNSUPPORTED = 36061;
};  // end of namespace FramebufferStatus

namespace RenderbufferTarget {
static const uint32_t GL_RENDERBUFFER = 36161;
};  // end of namespace RenderbufferTarget

namespace RenderbufferParameter {
static const uint32_t GL_RENDERBUFFER_WIDTH = 36162;
static const uint32_t GL_RENDERBUFFER_HEIGHT = 36163;
static const uint32_t GL_RENDERBUFFER_INTERNAL_FORMAT = 36164;
static const uint32_t GL_RENDERBUFFER_RED_SIZE = 36176;
static const uint32_t GL_RENDERBUFFER_GREEN_SIZE = 36177;
static const uint32_t GL_RENDERBUFFER_BLUE_SIZE = 36178;
static const uint32_t GL_RENDERBUFFER_ALPHA_SIZE = 36179;
static const uint32_t GL_RENDERBUFFER_DEPTH_SIZE = 36180;
static const uint32_t GL_RENDERBUFFER_STENCIL_SIZE = 36181;
};  // end of namespace RenderbufferParameter

namespace BufferTarget {
static const uint32_t GL_ARRAY_BUFFER = 34962;
static const uint32_t GL_ELEMENT_ARRAY_BUFFER = 34963;
};  // end of namespace BufferTarget

namespace BufferParameter {
static const uint32_t GL_BUFFER_SIZE = 34660;
static const uint32_t GL_BUFFER_USAGE = 34661;
};  // end of namespace BufferParameter

namespace TextureUnit {
static const uint32_t GL_TEXTURE0 = 33984;
static const uint32_t GL_TEXTURE1 = 33985;
static const uint32_t GL_TEXTURE2 = 33986;
static const uint32_t GL_TEXTURE3 = 33987;
static const uint32_t GL_TEXTURE4 = 33988;
static const uint32_t GL_TEXTURE5 = 33989;
static const uint32_t GL_TEXTURE6 = 33990;
static const uint32_t GL_TEXTURE7 = 33991;
static const uint32_t GL_TEXTURE8 = 33992;
static const uint32_t GL_TEXTURE9 = 33993;
static const uint32_t GL_TEXTURE10 = 33994;
static const uint32_t GL_TEXTURE11 = 33995;
static const uint32_t GL_TEXTURE12 = 33996;
static const uint32_t GL_TEXTURE13 = 33997;
static const uint32_t GL_TEXTURE14 = 33998;
static const uint32_t GL_TEXTURE15 = 33999;
static const uint32_t GL_TEXTURE16 = 34000;
static const uint32_t GL_TEXTURE17 = 34001;
static const uint32_t GL_TEXTURE18 = 34002;
static const uint32_t GL_TEXTURE19 = 34003;
static const uint32_t GL_TEXTURE20 = 34004;
static const uint32_t GL_TEXTURE21 = 34005;
static const uint32_t GL_TEXTURE22 = 34006;
static const uint32_t GL_TEXTURE23 = 34007;
static const uint32_t GL_TEXTURE24 = 34008;
static const uint32_t GL_TEXTURE25 = 34009;
static const uint32_t GL_TEXTURE26 = 34010;
static const uint32_t GL_TEXTURE27 = 34011;
static const uint32_t GL_TEXTURE28 = 34012;
static const uint32_t GL_TEXTURE29 = 34013;
static const uint32_t GL_TEXTURE30 = 34014;
static const uint32_t GL_TEXTURE31 = 34015;
};  // end of namespace TextureUnit

namespace BufferUsage {
static const uint32_t GL_DYNAMIC_DRAW = 35048;
static const uint32_t GL_STATIC_DRAW = 35044;
static const uint32_t GL_STREAM_DRAW = 35040;
};  // end of namespace BufferUsage

namespace ShaderType {
static const uint32_t GL_VERTEX_SHADER = 35633;
static const uint32_t GL_FRAGMENT_SHADER = 35632;
};  // end of namespace ShaderType

namespace StateVariable_GLES_2_0 {
static const uint32_t GL_ACTIVE_TEXTURE = 34016;
static const uint32_t GL_ALIASED_LINE_WIDTH_RANGE = 33902;
static const uint32_t GL_ALIASED_POINT_SIZE_RANGE = 33901;
static const uint32_t GL_ALPHA_BITS = 3413;
static const uint32_t GL_ARRAY_BUFFER_BINDING = 34964;
static const uint32_t GL_BLEND = 3042;
static const uint32_t GL_BLEND_COLOR = 32773;
static const uint32_t GL_BLEND_DST_ALPHA = 32970;
static const uint32_t GL_BLEND_DST_RGB = 32968;
static const uint32_t GL_BLEND_EQUATION_ALPHA = 34877;
static const uint32_t GL_BLEND_EQUATION_RGB = 32777;
static const uint32_t GL_BLEND_SRC_ALPHA = 32971;
static const uint32_t GL_BLEND_SRC_RGB = 32969;
static const uint32_t GL_BLUE_BITS = 3412;
static const uint32_t GL_COLOR_CLEAR_VALUE = 3106;
static const uint32_t GL_COLOR_WRITEMASK = 3107;
static const uint32_t GL_COMPRESSED_TEXTURE_FORMATS = 34467;
static const uint32_t GL_CULL_FACE = 2884;
static const uint32_t GL_CULL_FACE_MODE = 2885;
static const uint32_t GL_CURRENT_PROGRAM = 35725;
static const uint32_t GL_DEPTH_BITS = 3414;
static const uint32_t GL_DEPTH_CLEAR_VALUE = 2931;
static const uint32_t GL_DEPTH_FUNC = 2932;
static const uint32_t GL_DEPTH_RANGE = 2928;
static const uint32_t GL_DEPTH_TEST = 2929;
static const uint32_t GL_DEPTH_WRITEMASK = 2930;
static const uint32_t GL_DITHER = 3024;
static const uint32_t GL_ELEMENT_ARRAY_BUFFER_BINDING = 34965;
static const uint32_t GL_FRAMEBUFFER_BINDING = 36006;
static const uint32_t GL_FRONT_FACE = 2886;
static const uint32_t GL_GENERATE_MIPMAP_HINT = 33170;
static const uint32_t GL_GREEN_BITS = 3411;
static const uint32_t GL_IMPLEMENTATION_COLOR_READ_FORMAT = 35739;
static const uint32_t GL_IMPLEMENTATION_COLOR_READ_TYPE = 35738;
static const uint32_t GL_LINE_WIDTH = 2849;
static const uint32_t GL_MAX_COMBINED_TEXTURE_IMAGE_UNITS = 35661;
static const uint32_t GL_MAX_CUBE_MAP_TEXTURE_SIZE = 34076;
static const uint32_t GL_MAX_FRAGMENT_UNIFORM_VECTORS = 36349;
static const uint32_t GL_MAX_RENDERBUFFER_SIZE = 34024;
static const uint32_t GL_MAX_TEXTURE_IMAGE_UNITS = 34930;
static const uint32_t GL_MAX_TEXTURE_SIZE = 3379;
static const uint32_t GL_MAX_VARYING_VECTORS = 36348;
static const uint32_t GL_MAX_VERTEX_ATTRIBS = 34921;
static const uint32_t GL_MAX_VERTEX_TEXTURE_IMAGE_UNITS = 35660;
static const uint32_t GL_MAX_VERTEX_UNIFORM_VECTORS = 36347;
static const uint32_t GL_MAX_VIEWPORT_DIMS = 3386;
static const uint32_t GL_NUM_COMPRESSED_TEXTURE_FORMATS = 34466;
static const uint32_t GL_NUM_SHADER_BINARY_FORMATS = 36345;
static const uint32_t GL_PACK_ALIGNMENT = 3333;
static const uint32_t GL_POLYGON_OFFSET_FACTOR = 32824;
static const uint32_t GL_POLYGON_OFFSET_FILL = 32823;
static const uint32_t GL_POLYGON_OFFSET_UNITS = 10752;
static const uint32_t GL_RED_BITS = 3410;
static const uint32_t GL_RENDERBUFFER_BINDING = 36007;
static const uint32_t GL_SAMPLE_ALPHA_TO_COVERAGE = 32926;
static const uint32_t GL_SAMPLE_BUFFERS = 32936;
static const uint32_t GL_SAMPLE_COVERAGE = 32928;
static const uint32_t GL_SAMPLE_COVERAGE_INVERT = 32939;
static const uint32_t GL_SAMPLE_COVERAGE_VALUE = 32938;
static const uint32_t GL_SAMPLES = 32937;
static const uint32_t GL_SCISSOR_BOX = 3088;
static const uint32_t GL_SCISSOR_TEST = 3089;
static const uint32_t GL_SHADER_BINARY_FORMATS = 36344;
static const uint32_t GL_SHADER_COMPILER = 36346;
static const uint32_t GL_STENCIL_BACK_FAIL = 34817;
static const uint32_t GL_STENCIL_BACK_FUNC = 34816;
static const uint32_t GL_STENCIL_BACK_PASS_DEPTH_FAIL = 34818;
static const uint32_t GL_STENCIL_BACK_PASS_DEPTH_PASS = 34819;
static const uint32_t GL_STENCIL_BACK_REF = 36003;
static const uint32_t GL_STENCIL_BACK_VALUE_MASK = 36004;
static const uint32_t GL_STENCIL_BACK_WRITEMASK = 36005;
static const uint32_t GL_STENCIL_BITS = 3415;
static const uint32_t GL_STENCIL_CLEAR_VALUE = 2961;
static const uint32_t GL_STENCIL_FAIL = 2964;
static const uint32_t GL_STENCIL_FUNC = 2962;
static const uint32_t GL_STENCIL_PASS_DEPTH_FAIL = 2965;
static const uint32_t GL_STENCIL_PASS_DEPTH_PASS = 2966;
static const uint32_t GL_STENCIL_REF = 2967;
static const uint32_t GL_STENCIL_TEST = 2960;
static const uint32_t GL_STENCIL_VALUE_MASK = 2963;
static const uint32_t GL_STENCIL_WRITEMASK = 2968;
static const uint32_t GL_SUBPIXEL_BITS = 3408;
static const uint32_t GL_TEXTURE_BINDING_2D = 32873;
static const uint32_t GL_TEXTURE_BINDING_CUBE_MAP = 34068;
static const uint32_t GL_UNPACK_ALIGNMENT = 3317;
static const uint32_t GL_VIEWPORT = 2978;
};  // end of namespace StateVariable_GLES_2_0

namespace StateVariable_GLES_3_1 {
static const uint32_t GL_READ_FRAMEBUFFER_BINDING = 36010;
};  // end of namespace StateVariable_GLES_3_1

namespace StateVariable_EXT_texture_filter_anisotropic {
static const uint32_t GL_MAX_TEXTURE_MAX_ANISOTROPY_EXT = 34047;
};  // end of namespace StateVariable_EXT_texture_filter_anisotropic

namespace StateVariable_EXT_disjoint_timer_query {
static const uint32_t GL_GPU_DISJOINT_EXT = 36795;
};  // end of namespace StateVariable_EXT_disjoint_timer_query

namespace StateVariable {};  // end of namespace StateVariable

namespace FaceMode {
static const uint32_t GL_FRONT = 1028;
static const uint32_t GL_BACK = 1029;
static const uint32_t GL_FRONT_AND_BACK = 1032;
};  // end of namespace FaceMode

namespace ArrayType_GLES_1_1 {
static const uint32_t GL_VERTEX_ARRAY = 32884;
static const uint32_t GL_NORMAL_ARRAY = 32885;
static const uint32_t GL_COLOR_ARRAY = 32886;
static const uint32_t GL_TEXTURE_COORD_ARRAY = 32888;
};  // end of namespace ArrayType_GLES_1_1

namespace ArrayType_OES_point_size_array {
static const uint32_t GL_POINT_SIZE_ARRAY_OES = 35740;
};  // end of namespace ArrayType_OES_point_size_array

namespace ArrayType {};  // end of namespace ArrayType

namespace Capability {
static const uint32_t GL_BLEND = 3042;
static const uint32_t GL_CULL_FACE = 2884;
static const uint32_t GL_DEPTH_TEST = 2929;
static const uint32_t GL_DITHER = 3024;
static const uint32_t GL_POLYGON_OFFSET_FILL = 32823;
static const uint32_t GL_SAMPLE_ALPHA_TO_COVERAGE = 32926;
static const uint32_t GL_SAMPLE_COVERAGE = 32928;
static const uint32_t GL_SCISSOR_TEST = 3089;
static const uint32_t GL_STENCIL_TEST = 2960;
};  // end of namespace Capability

namespace StringConstant {
static const uint32_t GL_EXTENSIONS = 7939;
static const uint32_t GL_RENDERER = 7937;
static const uint32_t GL_VENDOR = 7936;
static const uint32_t GL_VERSION = 7938;
};  // end of namespace StringConstant

namespace VertexAttribType {
static const uint32_t GL_BYTE = 5120;
static const uint32_t GL_FIXED = 5132;
static const uint32_t GL_FLOAT = 5126;
static const uint32_t GL_SHORT = 5122;
static const uint32_t GL_UNSIGNED_BYTE = 5121;
static const uint32_t GL_UNSIGNED_SHORT = 5123;
};  // end of namespace VertexAttribType

namespace ShaderAttribType {
static const uint32_t GL_FLOAT = 5126;
static const uint32_t GL_FLOAT_VEC2 = 35664;
static const uint32_t GL_FLOAT_VEC3 = 35665;
static const uint32_t GL_FLOAT_VEC4 = 35666;
static const uint32_t GL_FLOAT_MAT2 = 35674;
static const uint32_t GL_FLOAT_MAT3 = 35675;
static const uint32_t GL_FLOAT_MAT4 = 35676;
};  // end of namespace ShaderAttribType

namespace ShaderUniformType {
static const uint32_t GL_FLOAT = 5126;
static const uint32_t GL_FLOAT_VEC2 = 35664;
static const uint32_t GL_FLOAT_VEC3 = 35665;
static const uint32_t GL_FLOAT_VEC4 = 35666;
static const uint32_t GL_INT = 5124;
static const uint32_t GL_INT_VEC2 = 35667;
static const uint32_t GL_INT_VEC3 = 35668;
static const uint32_t GL_INT_VEC4 = 35669;
static const uint32_t GL_BOOL = 35670;
static const uint32_t GL_BOOL_VEC2 = 35671;
static const uint32_t GL_BOOL_VEC3 = 35672;
static const uint32_t GL_BOOL_VEC4 = 35673;
static const uint32_t GL_FLOAT_MAT2 = 35674;
static const uint32_t GL_FLOAT_MAT3 = 35675;
static const uint32_t GL_FLOAT_MAT4 = 35676;
static const uint32_t GL_SAMPLER_2D = 35678;
static const uint32_t GL_SAMPLER_CUBE = 35680;
};  // end of namespace ShaderUniformType

namespace Error {
static const uint32_t GL_NO_ERROR = 0;
static const uint32_t GL_INVALID_ENUM = 1280;
static const uint32_t GL_INVALID_VALUE = 1281;
static const uint32_t GL_INVALID_OPERATION = 1282;
static const uint32_t GL_INVALID_FRAMEBUFFER_OPERATION = 1286;
static const uint32_t GL_OUT_OF_MEMORY = 1285;
};  // end of namespace Error

namespace HintTarget {
static const uint32_t GL_GENERATE_MIPMAP_HINT = 33170;
};  // end of namespace HintTarget

namespace HintMode {
static const uint32_t GL_DONT_CARE = 4352;
static const uint32_t GL_FASTEST = 4353;
static const uint32_t GL_NICEST = 4354;
};  // end of namespace HintMode

namespace DiscardFramebufferAttachment {
static const uint32_t GL_COLOR_EXT = 6144;
static const uint32_t GL_DEPTH_EXT = 6145;
static const uint32_t GL_STENCIL_EXT = 6146;
};  // end of namespace DiscardFramebufferAttachment

namespace ProgramParameter {
static const uint32_t GL_DELETE_STATUS = 35712;
static const uint32_t GL_LINK_STATUS = 35714;
static const uint32_t GL_VALIDATE_STATUS = 35715;
static const uint32_t GL_INFO_LOG_LENGTH = 35716;
static const uint32_t GL_ATTACHED_SHADERS = 35717;
static const uint32_t GL_ACTIVE_ATTRIBUTES = 35721;
static const uint32_t GL_ACTIVE_ATTRIBUTE_MAX_LENGTH = 35722;
static const uint32_t GL_ACTIVE_UNIFORMS = 35718;
static const uint32_t GL_ACTIVE_UNIFORM_MAX_LENGTH = 35719;
};  // end of namespace ProgramParameter

namespace ShaderParameter {
static const uint32_t GL_SHADER_TYPE = 35663;
static const uint32_t GL_DELETE_STATUS = 35712;
static const uint32_t GL_COMPILE_STATUS = 35713;
static const uint32_t GL_INFO_LOG_LENGTH = 35716;
static const uint32_t GL_SHADER_SOURCE_LENGTH = 35720;
};  // end of namespace ShaderParameter

namespace PixelStoreParameter {
static const uint32_t GL_PACK_ALIGNMENT = 3333;
static const uint32_t GL_UNPACK_ALIGNMENT = 3317;
};  // end of namespace PixelStoreParameter

namespace TextureParameter_FilterMode {
static const uint32_t GL_TEXTURE_MIN_FILTER = 10241;
static const uint32_t GL_TEXTURE_MAG_FILTER = 10240;
};  // end of namespace TextureParameter_FilterMode

namespace TextureParameter_WrapMode {
static const uint32_t GL_TEXTURE_WRAP_S = 10242;
static const uint32_t GL_TEXTURE_WRAP_T = 10243;
};  // end of namespace TextureParameter_WrapMode

namespace TextureParameter_EXT_texture_filter_anisotropic {
static const uint32_t GL_TEXTURE_MAX_ANISOTROPY_EXT = 34046;
};  // end of namespace TextureParameter_EXT_texture_filter_anisotropic

namespace TextureParameter_SwizzleMode {
static const uint32_t GL_TEXTURE_SWIZZLE_R = 36418;
static const uint32_t GL_TEXTURE_SWIZZLE_G = 36419;
static const uint32_t GL_TEXTURE_SWIZZLE_B = 36420;
static const uint32_t GL_TEXTURE_SWIZZLE_A = 36421;
};  // end of namespace TextureParameter_SwizzleMode

namespace TextureParameter {};  // end of namespace TextureParameter

namespace TextureFilterMode {
static const uint32_t GL_NEAREST = 9728;
static const uint32_t GL_LINEAR = 9729;
static const uint32_t GL_NEAREST_MIPMAP_NEAREST = 9984;
static const uint32_t GL_LINEAR_MIPMAP_NEAREST = 9985;
static const uint32_t GL_NEAREST_MIPMAP_LINEAR = 9986;
static const uint32_t GL_LINEAR_MIPMAP_LINEAR = 9987;
};  // end of namespace TextureFilterMode

namespace TextureWrapMode {
static const uint32_t GL_CLAMP_TO_EDGE = 33071;
static const uint32_t GL_MIRRORED_REPEAT = 33648;
static const uint32_t GL_REPEAT = 10497;
};  // end of namespace TextureWrapMode

namespace TexelComponent {
static const uint32_t GL_RED = 6403;
static const uint32_t GL_GREEN = 6404;
static const uint32_t GL_BLUE = 6405;
static const uint32_t GL_ALPHA = 6406;
};  // end of namespace TexelComponent

namespace BlendFactor {
static const uint32_t GL_ZERO = 0;
static const uint32_t GL_ONE = 1;
static const uint32_t GL_SRC_COLOR = 768;
static const uint32_t GL_ONE_MINUS_SRC_COLOR = 769;
static const uint32_t GL_DST_COLOR = 774;
static const uint32_t GL_ONE_MINUS_DST_COLOR = 775;
static const uint32_t GL_SRC_ALPHA = 770;
static const uint32_t GL_ONE_MINUS_SRC_ALPHA = 771;
static const uint32_t GL_DST_ALPHA = 772;
static const uint32_t GL_ONE_MINUS_DST_ALPHA = 773;
static const uint32_t GL_CONSTANT_COLOR = 32769;
static const uint32_t GL_ONE_MINUS_CONSTANT_COLOR = 32770;
static const uint32_t GL_CONSTANT_ALPHA = 32771;
static const uint32_t GL_ONE_MINUS_CONSTANT_ALPHA = 32772;
static const uint32_t GL_SRC_ALPHA_SATURATE = 776;
};  // end of namespace BlendFactor

namespace PrecisionType {
static const uint32_t GL_LOW_FLOAT = 36336;
static const uint32_t GL_MEDIUM_FLOAT = 36337;
static const uint32_t GL_HIGH_FLOAT = 36338;
static const uint32_t GL_LOW_INT = 36339;
static const uint32_t GL_MEDIUM_INT = 36340;
static const uint32_t GL_HIGH_INT = 36341;
};  // end of namespace PrecisionType

namespace TestFunction {
static const uint32_t GL_NEVER = 512;
static const uint32_t GL_LESS = 513;
static const uint32_t GL_EQUAL = 514;
static const uint32_t GL_LEQUAL = 515;
static const uint32_t GL_GREATER = 516;
static const uint32_t GL_NOTEQUAL = 517;
static const uint32_t GL_GEQUAL = 518;
static const uint32_t GL_ALWAYS = 519;
};  // end of namespace TestFunction

namespace StencilAction {
static const uint32_t GL_KEEP = 7680;
static const uint32_t GL_ZERO = 0;
static const uint32_t GL_REPLACE = 7681;
static const uint32_t GL_INCR = 7682;
static const uint32_t GL_INCR_WRAP = 34055;
static const uint32_t GL_DECR = 7683;
static const uint32_t GL_DECR_WRAP = 34056;
static const uint32_t GL_INVERT = 5386;
};  // end of namespace StencilAction

namespace FaceOrientation {
static const uint32_t GL_CW = 2304;
static const uint32_t GL_CCW = 2305;
};  // end of namespace FaceOrientation

namespace BlendEquation {
static const uint32_t GL_FUNC_ADD = 32774;
static const uint32_t GL_FUNC_SUBTRACT = 32778;
static const uint32_t GL_FUNC_REVERSE_SUBTRACT = 32779;
};  // end of namespace BlendEquation

namespace MapBufferTarget {
static const uint32_t GL_ARRAY_BUFFER = 34962;
static const uint32_t GL_COPY_READ_BUFFER = 36662;
static const uint32_t GL_COPY_WRITE_BUFFER = 36663;
static const uint32_t GL_ELEMENT_ARRAY_BUFFER = 34963;
static const uint32_t GL_PIXEL_PACK_BUFFER = 35051;
static const uint32_t GL_PIXEL_UNPACK_BUFFER = 35052;
static const uint32_t GL_TRANSFORM_FEEDBACK_BUFFER = 35982;
static const uint32_t GL_UNIFORM_BUFFER = 35345;
};  // end of namespace MapBufferTarget

namespace ImageTargetTexture_OES_EGL_image {
static const uint32_t GL_TEXTURE_2D = 3553;
};  // end of namespace ImageTargetTexture_OES_EGL_image

namespace ImageTargetTexture_OES_EGL_image_external {
static const uint32_t GL_TEXTURE_EXTERNAL_OES = 36197;
};  // end of namespace ImageTargetTexture_OES_EGL_image_external

namespace ImageTargetTexture {};  // end of namespace ImageTargetTexture

namespace ImageTargetRenderbufferStorage {
static const uint32_t GL_RENDERBUFFER_OES = 36161;
};  // end of namespace ImageTargetRenderbufferStorage

namespace ResetStatus {
static const uint32_t GL_NO_ERROR = 0;
static const uint32_t GL_GUILTY_CONTEXT_RESET_EXT = 33363;
static const uint32_t GL_INNOCENT_CONTEXT_RESET_EXT = 33364;
static const uint32_t GL_UNKNOWN_CONTEXT_RESET_EXT = 33365;
};  // end of namespace ResetStatus

namespace TextureKind {
static const uint32_t UNDEFINED = 0;
static const uint32_t TEXTURE2D = 1;
static const uint32_t CUBEMAP = 2;
};  // end of namespace TextureKind

namespace VertexAttribSize {
static const uint32_t SIZE_1 = 1;
static const uint32_t SIZE_2 = 2;
static const uint32_t SIZE_3 = 3;
static const uint32_t SIZE_4 = 4;
};  // end of namespace VertexAttribSize

namespace QueryParameter_GLES_3 {
static const uint32_t GL_CURRENT_QUERY = 34917;
};  // end of namespace QueryParameter_GLES_3

namespace QueryParameter_EXT_disjoint_timer_query {
static const uint32_t GL_QUERY_COUNTER_BITS_EXT = 34916;
};  // end of namespace QueryParameter_EXT_disjoint_timer_query

namespace QueryParameter {};  // end of namespace QueryParameter

namespace QueryObjectParameter_GLES_3 {
static const uint32_t GL_QUERY_RESULT = 34918;
static const uint32_t GL_QUERY_RESULT_AVAILABLE = 34919;
};  // end of namespace QueryObjectParameter_GLES_3

namespace QueryObjectParameter_EXT_disjoint_timer_query {
};  // end of namespace QueryObjectParameter_EXT_disjoint_timer_query

namespace QueryObjectParameter {};  // end of namespace QueryObjectParameter

namespace QueryTarget_GLES_3 {
static const uint32_t GL_ANY_SAMPLES_PASSED = 35887;
static const uint32_t GL_ANY_SAMPLES_PASSED_CONSERVATIVE = 36202;
static const uint32_t GL_TRANSFORM_FEEDBACK_PRIMITIVES_WRITTEN = 35976;
};  // end of namespace QueryTarget_GLES_3

namespace QueryTarget_EXT_disjoint_timer_query {
static const uint32_t GL_TIME_ELAPSED_EXT = 35007;
static const uint32_t GL_TIMESTAMP_EXT = 36392;
};  // end of namespace QueryTarget_EXT_disjoint_timer_query

namespace QueryTarget {};  // end of namespace QueryTarget

namespace DriverPropertyString {
static const uint32_t EGL_CLIENT_APIS = 0;
static const uint32_t EGL_EXTENSIONS = 1;
static const uint32_t EGL_VENDOR = 2;
static const uint32_t EGL_VERSION = 3;
static const uint32_t GL_EXTENSIONS = 4;
static const uint32_t GL_RENDERER = 5;
static const uint32_t GL_VENDOR = 6;
static const uint32_t GL_VERSION = 7;
};  // end of namespace DriverPropertyString

namespace TilePreserveMaskQCOM {
static const uint32_t GL_COLOR_BUFFER_BIT0_QCOM = 1;
static const uint32_t GL_COLOR_BUFFER_BIT1_QCOM = 2;
static const uint32_t GL_COLOR_BUFFER_BIT2_QCOM = 4;
static const uint32_t GL_COLOR_BUFFER_BIT3_QCOM = 8;
static const uint32_t GL_COLOR_BUFFER_BIT4_QCOM = 16;
static const uint32_t GL_COLOR_BUFFER_BIT5_QCOM = 32;
static const uint32_t GL_COLOR_BUFFER_BIT6_QCOM = 64;
static const uint32_t GL_COLOR_BUFFER_BIT7_QCOM = 128;
static const uint32_t GL_DEPTH_BUFFER_BIT0_QCOM = 256;
static const uint32_t GL_DEPTH_BUFFER_BIT1_QCOM = 512;
static const uint32_t GL_DEPTH_BUFFER_BIT2_QCOM = 1024;
static const uint32_t GL_DEPTH_BUFFER_BIT3_QCOM = 2048;
static const uint32_t GL_DEPTH_BUFFER_BIT4_QCOM = 4096;
static const uint32_t GL_DEPTH_BUFFER_BIT5_QCOM = 8192;
static const uint32_t GL_DEPTH_BUFFER_BIT6_QCOM = 16384;
static const uint32_t GL_DEPTH_BUFFER_BIT7_QCOM = 32768;
static const uint32_t GL_STENCIL_BUFFER_BIT0_QCOM = 65536;
static const uint32_t GL_STENCIL_BUFFER_BIT1_QCOM = 131072;
static const uint32_t GL_STENCIL_BUFFER_BIT2_QCOM = 262144;
static const uint32_t GL_STENCIL_BUFFER_BIT3_QCOM = 524288;
static const uint32_t GL_STENCIL_BUFFER_BIT4_QCOM = 1048576;
static const uint32_t GL_STENCIL_BUFFER_BIT5_QCOM = 2097152;
static const uint32_t GL_STENCIL_BUFFER_BIT6_QCOM = 4194304;
static const uint32_t GL_STENCIL_BUFFER_BIT7_QCOM = 8388608;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT0_QCOM = 16777216;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT1_QCOM = 33554432;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT2_QCOM = 67108864;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT3_QCOM = 134217728;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT4_QCOM = 268435456;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT5_QCOM = 536870912;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT6_QCOM = 1073741824;
static const uint32_t GL_MULTISAMPLE_BUFFER_BIT7_QCOM = 2147483648;
};  // end of namespace TilePreserveMaskQCOM

namespace ClearMask {
static const uint32_t GL_COLOR_BUFFER_BIT = 16384;
static const uint32_t GL_DEPTH_BUFFER_BIT = 256;
static const uint32_t GL_STENCIL_BUFFER_BIT = 1024;
};  // end of namespace ClearMask

namespace MapBufferRangeAccess {
static const uint32_t GL_MAP_READ_BIT = 1;
static const uint32_t GL_MAP_WRITE_BIT = 2;
static const uint32_t GL_MAP_INVALIDATE_RANGE_BIT = 4;
static const uint32_t GL_MAP_INVALIDATE_BUFFER_BIT = 8;
static const uint32_t GL_MAP_FLUSH_EXPLICIT_BIT = 16;
static const uint32_t GL_MAP_UNSYNCHRONIZED_BIT = 32;
};  // end of namespace MapBufferRangeAccess

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
typedef void *ImageOES;

typedef void(STDCALL *PFNEGLCREATECONTEXT)(int32_t *version, int32_t *context);
typedef void(STDCALL *PFNEGLMAKECURRENT)(int32_t context);
typedef void(STDCALL *PFNEGLSWAPBUFFERS)();
typedef void(STDCALL *PFNGLENABLECLIENTSTATE)(uint32_t type);
typedef void(STDCALL *PFNGLDISABLECLIENTSTATE)(uint32_t type);
typedef void(STDCALL *PFNGLGETPROGRAMBINARYOES)(ProgramId program, int32_t buffer_size,
                                                int32_t *bytes_written, uint32_t *binary_format,
                                                void *binary);
typedef void(STDCALL *PFNGLPROGRAMBINARYOES)(ProgramId program, uint32_t binary_format,
                                             void *binary, int32_t binary_size);
typedef void(STDCALL *PFNGLSTARTTILINGQCOM)(int32_t x, int32_t y, int32_t width, int32_t height,
                                            uint32_t preserveMask);
typedef void(STDCALL *PFNGLENDTILINGQCOM)(uint32_t preserve_mask);
typedef void(STDCALL *PFNGLDISCARDFRAMEBUFFEREXT)(uint32_t target, int32_t numAttachments,
                                                  uint32_t *attachments);
typedef void(STDCALL *PFNGLINSERTEVENTMARKEREXT)(int32_t length, const char *marker);
typedef void(STDCALL *PFNGLPUSHGROUPMARKEREXT)(int32_t length, const char *marker);
typedef void(STDCALL *PFNGLPOPGROUPMARKEREXT)();
typedef void(STDCALL *PFNGLTEXSTORAGE1DEXT)(uint32_t target, int32_t levels, uint32_t format,
                                            int32_t width);
typedef void(STDCALL *PFNGLTEXSTORAGE2DEXT)(uint32_t target, int32_t levels, uint32_t format,
                                            int32_t width, int32_t height);
typedef void(STDCALL *PFNGLTEXSTORAGE3DEXT)(uint32_t target, int32_t levels, uint32_t format,
                                            int32_t width, int32_t height, int32_t depth);
typedef void(STDCALL *PFNGLTEXTURESTORAGE1DEXT)(TextureId texture, uint32_t target, int32_t levels,
                                                uint32_t format, int32_t width);
typedef void(STDCALL *PFNGLTEXTURESTORAGE2DEXT)(TextureId texture, uint32_t target, int32_t levels,
                                                uint32_t format, int32_t width, int32_t height);
typedef void(STDCALL *PFNGLTEXTURESTORAGE3DEXT)(TextureId texture, uint32_t target, int32_t levels,
                                                uint32_t format, int32_t width, int32_t height,
                                                int32_t depth);
typedef void(STDCALL *PFNGLGENVERTEXARRAYSOES)(int32_t count, VertexArrayId *arrays);
typedef void(STDCALL *PFNGLBINDVERTEXARRAYOES)(VertexArrayId array);
typedef void(STDCALL *PFNGLDELETEVERTEXARRAYSOES)(int32_t count, VertexArrayId *arrays);
typedef bool(STDCALL *PFNGLISVERTEXARRAYOES)(VertexArrayId array);
typedef void(STDCALL *PFNGLEGLIMAGETARGETTEXTURE2DOES)(uint32_t target, ImageOES image);
typedef void(STDCALL *PFNGLEGLIMAGETARGETRENDERBUFFERSTORAGEOES)(uint32_t target,
                                                                 TexturePointer image);
typedef uint32_t(STDCALL *PFNGLGETGRAPHICSRESETSTATUSEXT)();
typedef void(STDCALL *PFNGLBINDATTRIBLOCATION)(ProgramId program, AttributeLocation location,
                                               const char *name);
typedef void(STDCALL *PFNGLBLENDFUNC)(uint32_t src_factor, uint32_t dst_factor);
typedef void(STDCALL *PFNGLBLENDFUNCSEPARATE)(uint32_t src_factor_rgb, uint32_t dst_factor_rgb,
                                              uint32_t src_factor_alpha, uint32_t dst_factor_alpha);
typedef void(STDCALL *PFNGLBLENDEQUATION)(uint32_t equation);
typedef void(STDCALL *PFNGLBLENDEQUATIONSEPARATE)(uint32_t rgb, uint32_t alpha);
typedef void(STDCALL *PFNGLBLENDCOLOR)(float red, float green, float blue, float alpha);
typedef void(STDCALL *PFNGLENABLEVERTEXATTRIBARRAY)(AttributeLocation location);
typedef void(STDCALL *PFNGLDISABLEVERTEXATTRIBARRAY)(AttributeLocation location);
typedef void(STDCALL *PFNGLVERTEXATTRIBPOINTER)(AttributeLocation location, uint32_t size,
                                                uint32_t type, bool normalized, int32_t stride,
                                                VertexPointer data);
typedef void(STDCALL *PFNGLGETACTIVEATTRIB)(ProgramId program, AttributeLocation location,
                                            int32_t buffer_size, int32_t *buffer_bytes_written,
                                            int32_t *vector_count, uint32_t *type,
                                            const char *name);
typedef void(STDCALL *PFNGLGETACTIVEUNIFORM)(ProgramId program, int32_t location,
                                             int32_t buffer_size, int32_t *buffer_bytes_written,
                                             int32_t *size, uint32_t *type, const char *name);
typedef uint32_t(STDCALL *PFNGLGETERROR)();
typedef void(STDCALL *PFNGLGETPROGRAMIV)(ProgramId program, uint32_t parameter, int32_t *value);
typedef void(STDCALL *PFNGLGETSHADERIV)(ShaderId shader, uint32_t parameter, int32_t *value);
typedef UniformLocation(STDCALL *PFNGLGETUNIFORMLOCATION)(ProgramId program, const char *name);
typedef AttributeLocation(STDCALL *PFNGLGETATTRIBLOCATION)(ProgramId program, const char *name);
typedef void(STDCALL *PFNGLPIXELSTOREI)(uint32_t parameter, int32_t value);
typedef void(STDCALL *PFNGLTEXPARAMETERI)(uint32_t target, uint32_t parameter, int32_t value);
typedef void(STDCALL *PFNGLTEXPARAMETERF)(uint32_t target, uint32_t parameter, float value);
typedef void(STDCALL *PFNGLGETTEXPARAMETERIV)(uint32_t target, uint32_t parameter, int32_t *values);
typedef void(STDCALL *PFNGLGETTEXPARAMETERFV)(uint32_t target, uint32_t parameter, float *values);
typedef void(STDCALL *PFNGLUNIFORM1I)(UniformLocation location, int32_t value);
typedef void(STDCALL *PFNGLUNIFORM2I)(UniformLocation location, int32_t value0, int32_t value1);
typedef void(STDCALL *PFNGLUNIFORM3I)(UniformLocation location, int32_t value0, int32_t value1,
                                      int32_t value2);
typedef void(STDCALL *PFNGLUNIFORM4I)(UniformLocation location, int32_t value0, int32_t value1,
                                      int32_t value2, int32_t value3);
typedef void(STDCALL *PFNGLUNIFORM1IV)(UniformLocation location, int32_t count, int32_t *value);
typedef void(STDCALL *PFNGLUNIFORM2IV)(UniformLocation location, int32_t count, int32_t *value);
typedef void(STDCALL *PFNGLUNIFORM3IV)(UniformLocation location, int32_t count, int32_t *value);
typedef void(STDCALL *PFNGLUNIFORM4IV)(UniformLocation location, int32_t count, int32_t *value);
typedef void(STDCALL *PFNGLUNIFORM1F)(UniformLocation location, float value);
typedef void(STDCALL *PFNGLUNIFORM2F)(UniformLocation location, float value0, float value1);
typedef void(STDCALL *PFNGLUNIFORM3F)(UniformLocation location, float value0, float value1,
                                      float value2);
typedef void(STDCALL *PFNGLUNIFORM4F)(UniformLocation location, float value0, float value1,
                                      float value2, float value3);
typedef void(STDCALL *PFNGLUNIFORM1FV)(UniformLocation location, int32_t count, float *value);
typedef void(STDCALL *PFNGLUNIFORM2FV)(UniformLocation location, int32_t count, float *value);
typedef void(STDCALL *PFNGLUNIFORM3FV)(UniformLocation location, int32_t count, float *value);
typedef void(STDCALL *PFNGLUNIFORM4FV)(UniformLocation location, int32_t count, float *value);
typedef void(STDCALL *PFNGLUNIFORMMATRIX2FV)(UniformLocation location, int32_t count,
                                             bool transpose, float *values);
typedef void(STDCALL *PFNGLUNIFORMMATRIX3FV)(UniformLocation location, int32_t count,
                                             bool transpose, float *values);
typedef void(STDCALL *PFNGLUNIFORMMATRIX4FV)(UniformLocation location, int32_t count,
                                             bool transpose, float *values);
typedef void(STDCALL *PFNGLGETUNIFORMFV)(ProgramId program, UniformLocation location,
                                         float *values);
typedef void(STDCALL *PFNGLGETUNIFORMIV)(ProgramId program, UniformLocation location,
                                         int32_t *values);
typedef void(STDCALL *PFNGLVERTEXATTRIB1F)(AttributeLocation location, float value0);
typedef void(STDCALL *PFNGLVERTEXATTRIB2F)(AttributeLocation location, float value0, float value1);
typedef void(STDCALL *PFNGLVERTEXATTRIB3F)(AttributeLocation location, float value0, float value1,
                                           float value2);
typedef void(STDCALL *PFNGLVERTEXATTRIB4F)(AttributeLocation location, float value0, float value1,
                                           float value2, float value3);
typedef void(STDCALL *PFNGLVERTEXATTRIB1FV)(AttributeLocation location, float *value);
typedef void(STDCALL *PFNGLVERTEXATTRIB2FV)(AttributeLocation location, float *value);
typedef void(STDCALL *PFNGLVERTEXATTRIB3FV)(AttributeLocation location, float *value);
typedef void(STDCALL *PFNGLVERTEXATTRIB4FV)(AttributeLocation location, float *value);
typedef void(STDCALL *PFNGLGETSHADERPRECISIONFORMAT)(uint32_t shader_type, uint32_t precision_type,
                                                     int32_t *range, int32_t *precision);
typedef void(STDCALL *PFNGLDEPTHMASK)(bool enabled);
typedef void(STDCALL *PFNGLDEPTHFUNC)(uint32_t function);
typedef void(STDCALL *PFNGLDEPTHRANGEF)(float near, float far);
typedef void(STDCALL *PFNGLCOLORMASK)(bool red, bool green, bool blue, bool alpha);
typedef void(STDCALL *PFNGLSTENCILMASK)(uint32_t mask);
typedef void(STDCALL *PFNGLSTENCILMASKSEPARATE)(uint32_t face, uint32_t mask);
typedef void(STDCALL *PFNGLSTENCILFUNCSEPARATE)(uint32_t face, uint32_t function,
                                                int32_t reference_value, int32_t mask);
typedef void(STDCALL *PFNGLSTENCILOPSEPARATE)(uint32_t face, uint32_t stencil_fail,
                                              uint32_t stencil_pass_depth_fail,
                                              uint32_t stencil_pass_depth_pass);
typedef void(STDCALL *PFNGLFRONTFACE)(uint32_t orientation);
typedef void(STDCALL *PFNGLVIEWPORT)(int32_t x, int32_t y, int32_t width, int32_t height);
typedef void(STDCALL *PFNGLSCISSOR)(int32_t x, int32_t y, int32_t width, int32_t height);
typedef void(STDCALL *PFNGLACTIVETEXTURE)(uint32_t unit);
typedef void(STDCALL *PFNGLGENTEXTURES)(int32_t count, TextureId *textures);
typedef void(STDCALL *PFNGLDELETETEXTURES)(int32_t count, TextureId *textures);
typedef bool(STDCALL *PFNGLISTEXTURE)(TextureId texture);
typedef void(STDCALL *PFNGLBINDTEXTURE)(uint32_t target, TextureId texture);
typedef void(STDCALL *PFNGLTEXIMAGE2D)(uint32_t target, int32_t level, uint32_t internal_format,
                                       int32_t width, int32_t height, int32_t border,
                                       uint32_t format, uint32_t type, TexturePointer data);
typedef void(STDCALL *PFNGLTEXSUBIMAGE2D)(uint32_t target, int32_t level, int32_t xoffset,
                                          int32_t yoffset, int32_t width, int32_t height,
                                          uint32_t format, uint32_t type, TexturePointer data);
typedef void(STDCALL *PFNGLCOPYTEXIMAGE2D)(uint32_t target, int32_t level, uint32_t format,
                                           int32_t x, int32_t y, int32_t width, int32_t height,
                                           int32_t border);
typedef void(STDCALL *PFNGLCOPYTEXSUBIMAGE2D)(uint32_t target, int32_t level, int32_t xoffset,
                                              int32_t yoffset, int32_t x, int32_t y, int32_t width,
                                              int32_t height);
typedef void(STDCALL *PFNGLCOMPRESSEDTEXIMAGE2D)(uint32_t target, int32_t level, uint32_t format,
                                                 int32_t width, int32_t height, int32_t border,
                                                 int32_t image_size, TexturePointer data);
typedef void(STDCALL *PFNGLCOMPRESSEDTEXSUBIMAGE2D)(uint32_t target, int32_t level, int32_t xoffset,
                                                    int32_t yoffset, int32_t width, int32_t height,
                                                    uint32_t format, int32_t image_size,
                                                    TexturePointer data);
typedef void(STDCALL *PFNGLGENERATEMIPMAP)(uint32_t target);
typedef void(STDCALL *PFNGLREADPIXELS)(int32_t x, int32_t y, int32_t width, int32_t height,
                                       uint32_t format, uint32_t type, void *data);
typedef void(STDCALL *PFNGLGENFRAMEBUFFERS)(int32_t count, FramebufferId *framebuffers);
typedef void(STDCALL *PFNGLBINDFRAMEBUFFER)(uint32_t target, FramebufferId framebuffer);
typedef uint32_t(STDCALL *PFNGLCHECKFRAMEBUFFERSTATUS)(uint32_t target);
typedef void(STDCALL *PFNGLDELETEFRAMEBUFFERS)(int32_t count, FramebufferId *framebuffers);
typedef bool(STDCALL *PFNGLISFRAMEBUFFER)(FramebufferId framebuffer);
typedef void(STDCALL *PFNGLGENRENDERBUFFERS)(int32_t count, RenderbufferId *renderbuffers);
typedef void(STDCALL *PFNGLBINDRENDERBUFFER)(uint32_t target, RenderbufferId renderbuffer);
typedef void(STDCALL *PFNGLRENDERBUFFERSTORAGE)(uint32_t target, uint32_t format, int32_t width,
                                                int32_t height);
typedef void(STDCALL *PFNGLDELETERENDERBUFFERS)(int32_t count, RenderbufferId *renderbuffers);
typedef bool(STDCALL *PFNGLISRENDERBUFFER)(RenderbufferId renderbuffer);
typedef void(STDCALL *PFNGLGETRENDERBUFFERPARAMETERIV)(uint32_t target, uint32_t parameter,
                                                       int32_t *values);
typedef void(STDCALL *PFNGLGENBUFFERS)(int32_t count, BufferId *buffers);
typedef void(STDCALL *PFNGLBINDBUFFER)(uint32_t target, BufferId buffer);
typedef void(STDCALL *PFNGLBUFFERDATA)(uint32_t target, int32_t size, BufferDataPointer data,
                                       uint32_t usage);
typedef void(STDCALL *PFNGLBUFFERSUBDATA)(uint32_t target, int32_t offset, int32_t size,
                                          void *data);
typedef void(STDCALL *PFNGLDELETEBUFFERS)(int32_t count, BufferId *buffers);
typedef bool(STDCALL *PFNGLISBUFFER)(BufferId buffer);
typedef void(STDCALL *PFNGLGETBUFFERPARAMETERIV)(uint32_t target, uint32_t parameter,
                                                 int32_t *value);
typedef ShaderId(STDCALL *PFNGLCREATESHADER)(uint32_t type);
typedef void(STDCALL *PFNGLDELETESHADER)(ShaderId shader);
typedef void(STDCALL *PFNGLSHADERSOURCE)(ShaderId shader, int32_t count, const char **source,
                                         int32_t *length);
typedef void(STDCALL *PFNGLSHADERBINARY)(int32_t count, ShaderId *shaders, uint32_t binary_format,
                                         void *binary, int32_t binary_size);
typedef void(STDCALL *PFNGLGETSHADERINFOLOG)(ShaderId shader, int32_t buffer_length,
                                             int32_t *string_length_written, const char *info);
typedef void(STDCALL *PFNGLGETSHADERSOURCE)(ShaderId shader, int32_t buffer_length,
                                            int32_t *string_length_written, const char *source);
typedef void(STDCALL *PFNGLRELEASESHADERCOMPILER)();
typedef void(STDCALL *PFNGLCOMPILESHADER)(ShaderId shader);
typedef bool(STDCALL *PFNGLISSHADER)(ShaderId shader);
typedef ProgramId(STDCALL *PFNGLCREATEPROGRAM)();
typedef void(STDCALL *PFNGLDELETEPROGRAM)(ProgramId program);
typedef void(STDCALL *PFNGLATTACHSHADER)(ProgramId program, ShaderId shader);
typedef void(STDCALL *PFNGLDETACHSHADER)(ProgramId program, ShaderId shader);
typedef void(STDCALL *PFNGLGETATTACHEDSHADERS)(ProgramId program, int32_t buffer_length,
                                               int32_t *shaders_length_written, ShaderId *shaders);
typedef void(STDCALL *PFNGLLINKPROGRAM)(ProgramId program);
typedef void(STDCALL *PFNGLGETPROGRAMINFOLOG)(ProgramId program, int32_t buffer_length,
                                              int32_t *string_length_written, const char *info);
typedef void(STDCALL *PFNGLUSEPROGRAM)(ProgramId program);
typedef bool(STDCALL *PFNGLISPROGRAM)(ProgramId program);
typedef void(STDCALL *PFNGLVALIDATEPROGRAM)(ProgramId program);
typedef void(STDCALL *PFNGLCLEARCOLOR)(float r, float g, float b, float a);
typedef void(STDCALL *PFNGLCLEARDEPTHF)(float depth);
typedef void(STDCALL *PFNGLCLEARSTENCIL)(int32_t stencil);
typedef void(STDCALL *PFNGLCLEAR)(uint32_t mask);
typedef void(STDCALL *PFNGLCULLFACE)(uint32_t mode);
typedef void(STDCALL *PFNGLPOLYGONOFFSET)(float scale_factor, float units);
typedef void(STDCALL *PFNGLLINEWIDTH)(float width);
typedef void(STDCALL *PFNGLSAMPLECOVERAGE)(float value, bool invert);
typedef void(STDCALL *PFNGLHINT)(uint32_t target, uint32_t mode);
typedef void(STDCALL *PFNGLFRAMEBUFFERRENDERBUFFER)(uint32_t framebuffer_target,
                                                    uint32_t framebuffer_attachment,
                                                    uint32_t renderbuffer_target,
                                                    RenderbufferId renderbuffer);
typedef void(STDCALL *PFNGLFRAMEBUFFERTEXTURE2D)(uint32_t framebuffer_target,
                                                 uint32_t framebuffer_attachment,
                                                 uint32_t texture_target, TextureId texture,
                                                 int32_t level);
typedef void(STDCALL *PFNGLGETFRAMEBUFFERATTACHMENTPARAMETERIV)(uint32_t target,
                                                                uint32_t attachment,
                                                                uint32_t parameter, int32_t *value);
typedef void(STDCALL *PFNGLDRAWELEMENTS)(uint32_t draw_mode, int32_t element_count,
                                         uint32_t indices_type, IndicesPointer indices);
typedef void(STDCALL *PFNGLDRAWARRAYS)(uint32_t draw_mode, int32_t first_index,
                                       int32_t index_count);
typedef void(STDCALL *PFNGLFLUSH)();
typedef void(STDCALL *PFNGLFINISH)();
typedef void(STDCALL *PFNGLGETBOOLEANV)(uint32_t param, bool *values);
typedef void(STDCALL *PFNGLGETFLOATV)(uint32_t param, float *values);
typedef void(STDCALL *PFNGLGETINTEGERV)(uint32_t param, int32_t *values);
typedef const char *(STDCALL *PFNGLGETSTRING)(uint32_t param);
typedef void(STDCALL *PFNGLENABLE)(uint32_t capability);
typedef void(STDCALL *PFNGLDISABLE)(uint32_t capability);
typedef bool(STDCALL *PFNGLISENABLED)(uint32_t capability);
typedef void *(STDCALL *PFNGLMAPBUFFERRANGE)(uint32_t target, int32_t offset, int32_t length,
                                             uint32_t access);
typedef void(STDCALL *PFNGLUNMAPBUFFER)(uint32_t target);
typedef void(STDCALL *PFNGLINVALIDATEFRAMEBUFFER)(uint32_t target, int32_t count,
                                                  uint32_t *attachments);
typedef void(STDCALL *PFNGLRENDERBUFFERSTORAGEMULTISAMPLE)(uint32_t target, int32_t samples,
                                                           uint32_t format, int32_t width,
                                                           int32_t height);
typedef void(STDCALL *PFNGLBLITFRAMEBUFFER)(int32_t srcX0, int32_t srcY0, int32_t srcX1,
                                            int32_t srcY1, int32_t dstX0, int32_t dstY0,
                                            int32_t dstX1, int32_t dstY1, uint32_t mask,
                                            uint32_t filter);
typedef void(STDCALL *PFNGLGENQUERIES)(int32_t count, QueryId *queries);
typedef void(STDCALL *PFNGLBEGINQUERY)(uint32_t target, QueryId query);
typedef void(STDCALL *PFNGLENDQUERY)(uint32_t target);
typedef void(STDCALL *PFNGLDELETEQUERIES)(int32_t count, QueryId *queries);
typedef bool(STDCALL *PFNGLISQUERY)(QueryId query);
typedef void(STDCALL *PFNGLGETQUERYIV)(uint32_t target, uint32_t parameter, int32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTUIV)(QueryId query, uint32_t parameter, uint32_t *value);
typedef void(STDCALL *PFNGLGENQUERIESEXT)(int32_t count, QueryId *queries);
typedef void(STDCALL *PFNGLBEGINQUERYEXT)(uint32_t target, QueryId query);
typedef void(STDCALL *PFNGLENDQUERYEXT)(uint32_t target);
typedef void(STDCALL *PFNGLDELETEQUERIESEXT)(int32_t count, QueryId *queries);
typedef bool(STDCALL *PFNGLISQUERYEXT)(QueryId query);
typedef void(STDCALL *PFNGLQUERYCOUNTEREXT)(QueryId query, uint32_t target);
typedef void(STDCALL *PFNGLGETQUERYIVEXT)(uint32_t target, uint32_t parameter, int32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTIVEXT)(QueryId query, uint32_t parameter, int32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTUIVEXT)(QueryId query, uint32_t parameter,
                                                 uint32_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTI64VEXT)(QueryId query, uint32_t parameter,
                                                  int64_t *value);
typedef void(STDCALL *PFNGLGETQUERYOBJECTUI64VEXT)(QueryId query, uint32_t parameter,
                                                   uint64_t *value);

extern PFNEGLCREATECONTEXT eglCreateContext;
extern PFNEGLMAKECURRENT eglMakeCurrent;
extern PFNEGLSWAPBUFFERS eglSwapBuffers;
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

}  // end of namespace gfxapi
}  // end of namespace caze
}  // end of namespace android

#endif  // ANDROID_CAZE_GFXAPI_FUNCTIONS_H
